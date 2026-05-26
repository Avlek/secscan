package scanner

import (
	"crypto/tls"
	"strconv"
	"strings"
	"time"
)

type SSLScanner struct{}

func (scanner *SSLScanner) Scan(host string) []Result {
	var results []Result

	conn, err := tls.Dial("tcp", host+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return []Result{{
			Name:    "SSL",
			Status:  StatusCritical,
			Details: "не удалось подключиться: " + err.Error(),
		}}
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	days := int(cert.NotAfter.Sub(time.Now()).Hours() / 24)
	status := StatusOk
	if days < 30 {
		status = StatusCritical
	} else if days < 90 {
		status = StatusWarning
	}

	results = append(results, Result{
		Name:    "Expires in",
		Status:  status,
		Details: strconv.Itoa(days) + " days",
	})

	organization := "unknown"
	if len(cert.Issuer.Organization) > 0 {
		organization = cert.Issuer.Organization[0]
	}
	results = append(results, Result{
		Name:    "Issuer",
		Status:  StatusWarning,
		Details: organization,
	})

	results = append(results, Result{
		Name:    "DNS Names",
		Status:  StatusOk,
		Details: strings.Join(cert.DNSNames, ", "),
	})

	return results
}
