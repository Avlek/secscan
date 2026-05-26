package scanner

import (
	"fmt"
	"net"
	"sort"
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

var portStatus = map[int]Statuses{

	80:  StatusOk,
	443: StatusOk,
	25:  StatusOk,
	465: StatusOk,
	587: StatusOk,
	993: StatusOk,
	995: StatusOk,
	53:  StatusOk,
	123: StatusOk,

	22:   StatusWarning,
	8080: StatusWarning,
	8443: StatusWarning,
	8888: StatusWarning,
	3000: StatusWarning,
	5000: StatusWarning,
	9000: StatusWarning,
	8001: StatusWarning,
	21:   StatusWarning,
	23:   StatusWarning,
	161:  StatusWarning,

	3306:  StatusCritical,
	5432:  StatusCritical,
	27017: StatusCritical,
	27018: StatusCritical,
	6379:  StatusCritical,
	1433:  StatusCritical,
	1521:  StatusCritical,
	5984:  StatusCritical,
	9200:  StatusCritical,
	9300:  StatusCritical,
	2375:  StatusCritical,
	2376:  StatusCritical,
	6443:  StatusCritical,
	3389:  StatusCritical,
	5900:  StatusCritical,
	5901:  StatusCritical,
	445:   StatusCritical,
	139:   StatusCritical,
	137:   StatusCritical,
	11211: StatusCritical,
	9092:  StatusCritical,
	2181:  StatusCritical,
	4369:  StatusCritical,
	5672:  StatusCritical,
	8500:  StatusCritical,
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
				status := StatusWarning
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

	sort.Slice(results, func(i, j int) bool {
		port1, _ := strconv.Atoi(results[i].Name)
		port2, _ := strconv.Atoi(results[j].Name)
		return statusPriority[results[i].Status] < statusPriority[results[j].Status] ||
			statusPriority[results[i].Status] == statusPriority[results[j].Status] && port1 < port2
	})

	return results
}
