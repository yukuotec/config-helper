// facts/facts.go
package facts

import (
	"config-helper/sshclient"
)

// GetMemorySize retrieves the memory size of the remote server using the provided SSH client.
func GetMemorySize(client *sshclient.Client) (string, error) {
	cmd := "grep MemTotal /proc/meminfo"
	output, err := client.RunCommand(cmd)
	if err != nil {
		return "", err
	}

	return output, nil
}
