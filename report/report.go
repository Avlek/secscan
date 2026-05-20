package report

import (
	"fmt"
	"secscan/scanner"
)

func printHeader(header string) {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("    " + header)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func PrintPorts(results []scanner.Result) {
	printHeader("PORTS")

	for i := 0; i < len(results); i++ {
		fmt.Printf("%6s %-22s %-12s\n", results[i].Name, results[i].Details, fmt.Sprintf("(%s)", results[i].Status))
	}
}
