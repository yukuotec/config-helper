// dirs/dirs.go
package dirs

import (
	"config-helper/sshclient" // Import the sshclient package
	"fmt"
)

// EnsureDirRemote ensures that a directory exists with the specified name and owner on a remote server.
func EnsureDirRemote(client *sshclient.Client, dirPath, owner, mode string) error {
	// Ensure directory exists
	cmd := fmt.Sprintf("mkdir -p %s", dirPath)
	if _, err := client.RunCommand(cmd); err != nil {
		return fmt.Errorf("failed to ensure directory: %w", err)
	}

	// Set the owner of the directory
	cmd = fmt.Sprintf("chown %s %s", owner, dirPath)
	if _, err := client.RunCommand(cmd); err != nil {
		return fmt.Errorf("failed to set owner: %w", err)
	}

	// Set the mode of the directory
	cmd = fmt.Sprintf("chmod %s %s", mode, dirPath)
	if _, err := client.RunCommand(cmd); err != nil {
		return fmt.Errorf("failed to set mode: %w", err)
	}

	fmt.Printf("Directory %s created and ownership set to %s with mode %s\n", dirPath, owner, mode)
	return nil
}
