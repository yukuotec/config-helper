package main

import (
	"config-helper/config"
	"config-helper/sshclient"
	"config-helper/task"
	"flag"
	"log"
)

func main() {
	// Define command-line flags
	configFilePath := flag.String("f", "config.yaml", "Path to the configuration file")
	flag.Parse()

	// Load configuration from the provided file
	cfg, err := config.LoadConfig(*configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config from %s: %v", *configFilePath, err)
	}

	// Create SSH client
	sshClient, err := sshclient.NewClient(cfg.Host.Host, cfg.Host.User, cfg.Host.KeyPath)
	if err != nil {
		log.Fatalf("Failed to create SSH client: %v", err)
	}
	defer sshClient.Close()

	// Retrieve facts
	for _, cmd := range cfg.Facts.Commands {
		output, err := sshClient.RunCommand(cmd)
		if err != nil {
			log.Fatalf("Failed to run command '%s': %v", cmd, err)
		}
		log.Printf("Command '%s' output: %s", cmd, output)
	}

	// Execute tasks
	for _, taskConfig := range cfg.Tasks {
		t, err := task.NewTask(taskConfig.Category, taskConfig.Type, taskConfig.Parameters)
		if err != nil {
			log.Fatalf("Failed to create task: %v", err)
		}
		if err := t.Validate(); err != nil {
			log.Fatalf("Invalid task parameters: %v", err)
		}
		if err := t.Execute(sshClient); err != nil {
			log.Printf("Task execution failed: %v", err)
		} else {
			log.Printf("Tasks execution succeed: %T", t)
		}
	}
}
