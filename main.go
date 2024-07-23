package main

import (
	"config-helper/dirs"
	"config-helper/facts"
	"config-helper/sshclient"
	"fmt"
	"log"
)

func main() {
	// Replace these variables with your actual values
	user := "vagrant"
	keyPath := "./private_key"
	host := "localhost:2222"

	// Create a new SSH client
	client, err := sshclient.NewClient(user, keyPath, host)
	if err != nil {
		log.Fatalf("Error creating SSH client: %v", err)
	}

	defer client.Client.Close()

	memorySize, err := facts.GetMemorySize(client)
	if err != nil {
		log.Fatalf("Error retrieving memory size: %v", err)
	}

	// Print the result
	fmt.Printf("Memory size: %s", memorySize)

	// Directory details
	dirPath := "/tmp/test"
	owner := "vagrant:vagrant"
	mode := "0755"

	// Ensure directory exists and set ownership on the remote server using the dirs package
	if err := dirs.EnsureDirRemote(client, dirPath, owner, mode); err != nil {
		log.Fatalf("Error ensuring directory on remote server: %v", err)
	}
}
