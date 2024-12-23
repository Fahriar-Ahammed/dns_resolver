package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Step 1: Take a domain as input from the user
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <domain>")
		return
	}
	domain := os.Args[1]

	// Step 2: Resolve IP addresses
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("Failed to resolve IP for %s: %v\n", domain, err)
		return
	}

	// Step 3: Print the resolved IP addresses
	fmt.Printf("IP addresses for %s:\n", domain)
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}
