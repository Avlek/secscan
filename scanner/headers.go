package scanner

import (
	"net/http"
	"sort"
	"time"
)

var securityHeaders = map[string]bool{
	"Strict-Transport-Security":         true,
	"Content-Security-Policy":           true,
	"X-Frame-Options":                   true,
	"X-Content-Type-Options":            true,
	"Permissions-Policy":                true,
	"Referrer-Policy":                   false,
	"X-XSS-Protection":                  false,
	"Cross-Origin-Opener-Policy":        false,
	"Cross-Origin-Embedder-Policy":      false,
	"Cross-Origin-Resource-Policy":      false,
	"Cache-Control":                     false,
	"Clear-Site-Data":                   false,
	"X-Permitted-Cross-Domain-Policies": false,
	"Expect-CT":                         false,
	"X-DNS-Prefetch-Control":            false,
}

type HeaderScanner struct{}

func (scanner *HeaderScanner) Scan(host string) []Result {
	var results []Result

	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get("https://" + host)
	if err != nil {
		resp, err = client.Get("http://" + host)
		if err != nil {
			return []Result{
				{
					Name:    "HTTP",
					Status:  StatusCritical,
					Details: err.Error(),
				},
			}
		}
	}
	defer resp.Body.Close()

	for k := range securityHeaders {
		if len(resp.Header.Get(k)) == 0 {
			status := StatusWarning
			if securityHeaders[k] {
				status = StatusCritical
			}
			results = append(results, Result{
				Name:    k,
				Status:  status,
				Details: "Missing",
			})
		} else {
			results = append(results, Result{
				Name:    k,
				Status:  StatusOk,
				Details: "Ok",
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Status < results[j].Status
	})

	return results
}
