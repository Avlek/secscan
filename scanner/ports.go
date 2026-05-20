package scanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var targetPorts = map[int]string{
	80:   "HTTP",
	443:  "HTTPS",
	8080: "HTTP-Alt",
	8443: "HTTPS-Alt",
	8888: "HTTP-Alt",
	3000: "Dev-Server",
	5000: "Dev-Server",
	9000: "Dev-Server",

	22:   "SSH",
	23:   "Telnet",
	3389: "RDP",
	5900: "VNC",
	5901: "VNC-Alt",

	3306:  "MySQL",
	5432:  "PostgreSQL",
	27017: "MongoDB",
	27018: "MongoDB-Alt",
	6379:  "Redis",
	1433:  "MSSQL",
	1521:  "Oracle",
	5984:  "CouchDB",
	9200:  "Elasticsearch",
	9300:  "Elasticsearch-Cluster",

	25:  "SMTP",
	110: "POP3",
	143: "IMAP",
	465: "SMTPS",
	587: "SMTP-Submission",
	993: "IMAPS",
	995: "POP3S",

	21:  "FTP",
	20:  "FTP-Data",
	69:  "TFTP",
	445: "SMB",
	139: "NetBIOS",
	137: "NetBIOS-NS",

	53:  "DNS",
	161: "SNMP",
	162: "SNMP-Trap",
	123: "NTP",

	2375: "Docker",
	2376: "Docker-TLS",
	6443: "Kubernetes",
	8001: "Kubernetes-API",
	2181: "Zookeeper",
	9092: "Kafka",
	4369: "RabbitMQ",
	5672: "AMQP",

	9090: "Prometheus",
	3100: "Loki",
	9411: "Zipkin",

	11211: "Memcached",
	8500:  "Consul",
	8600:  "Consul-DNS",
}

var portStatus = map[int]string{

	80:  "ok",
	443: "ok",
	25:  "ok",
	465: "ok",
	587: "ok",
	993: "ok",
	995: "ok",
	53:  "ok",
	123: "ok",

	22:   "warning",
	8080: "warning",
	8443: "warning",
	8888: "warning",
	3000: "warning",
	5000: "warning",
	9000: "warning",
	8001: "warning",
	21:   "warning",
	23:   "warning",
	161:  "warning",

	3306:  "critical",
	5432:  "critical",
	27017: "critical",
	27018: "critical",
	6379:  "critical",
	1433:  "critical",
	1521:  "critical",
	5984:  "critical",
	9200:  "critical",
	9300:  "critical",
	2375:  "critical",
	2376:  "critical",
	6443:  "critical",
	3389:  "critical",
	5900:  "critical",
	5901:  "critical",
	445:   "critical",
	139:   "critical",
	137:   "critical",
	11211: "critical",
	9092:  "critical",
	2181:  "critical",
	4369:  "critical",
	5672:  "critical",
	8500:  "critical",
}

type PortScanner struct{}

func (scanner *PortScanner) Scan(host string) []Result {
	var results []Result
	var mu sync.Mutex

	timeout := 2 * time.Second
	wg := sync.WaitGroup{}
	for k := range targetPorts {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
			if err == nil {
				status := "warning"
				if _, ok := portStatus[port]; ok {
					status = portStatus[port]
				}
				mu.Lock()
				results = append(results, Result{
					Name:    strconv.Itoa(port),
					Status:  status,
					Details: targetPorts[port],
				})
				mu.Unlock()
				conn.Close()
			}
		}(k)
	}
	wg.Wait()

	return results
}
