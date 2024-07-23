package shell

import (
	"fmt"
	"log"
	"strings"

	"config-helper/sshclient"
)

// ShellExecTask represents a shell execution task.
type ShellExecTask struct {
	Command string
}

// NewReplaceTask creates a new ReplaceTask
func NewShellExecTask(params map[string]string) (*ShellExecTask, error) {
	command, ok := params["command"]
	if !ok {
		return nil, fmt.Errorf("missing parameter: command")
	}
	return &ShellExecTask{Command: command}, nil
}

// Ensure ShellExecTask implements Task interface.
func (t *ShellExecTask) Validate() error {
	if strings.TrimSpace(t.Command) == "" {
		return fmt.Errorf("command cannot be empty")
	}
	return nil
}

func (t *ShellExecTask) Execute(client *sshclient.Client) error {
	output, err := client.RunCommand(t.Command)
	if err != nil {
		return fmt.Errorf("failed to execute command '%s': %v", t.Command, err)
	}
	log.Printf("Command '%s' output: %s", t.Command, output)
	return nil
}
