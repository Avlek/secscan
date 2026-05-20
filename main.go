package main

import (
	"fmt"
	"os"
	"secscan/report"

	"secscan/scanner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}

	url := os.Args[1]
	portScanner := scanner.PortScanner{}
	report.PrintPorts(portScanner.Scan(url))
	headerScanner := scanner.HeaderScanner{}
	report.PrintHeaders(headerScanner.Scan(url))
}
