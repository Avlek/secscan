package report

import (
	"fmt"
	"secscan/scanner"
)

var icons = map[scanner.Statuses]string{
	scanner.StatusOk:       "[ OK ]",
	scanner.StatusWarning:  "[ !! ]",
	scanner.StatusCritical: "[CRIT]",
}

func printHeader(header string) {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("    " + header)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func PrintPorts(results []scanner.Result) {
	printHeader("PORTS")

	for i := 0; i < len(results); i++ {
		fmt.Printf("%-6s %-10s %-22s\n", icons[results[i].Status], results[i].Name, results[i].Details)
	}
}

func PrintHeaders(results []scanner.Result) {
	printHeader("HEADERS")

	for i := 0; i < len(results); i++ {
		fmt.Printf("%-6s %s\n", icons[results[i].Status], results[i].Name)
	}
}

func PrintSSL(results []scanner.Result) {
	printHeader("SSL")

	for i := 0; i < len(results); i++ {
		fmt.Printf("%-6s %s: %s\n", icons[results[i].Status], results[i].Name, results[i].Details)
	}
}
