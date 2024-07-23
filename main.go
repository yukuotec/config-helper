package main

import (
	"config-helper/facts"
	"fmt"
	"log"
)

func main() {
	// Replace these variables with your actual values
	user := "vagrant"
	keyPath := "./private_key"
	host := "localhost:2222"

	memorySize, err := facts.GetMemorySize(user, keyPath, host)
	if err != nil {
		log.Fatalf("Error retrieving memory size: %v", err)
	}

	// Print the result
	fmt.Printf("Memory size: %s", memorySize)
}
