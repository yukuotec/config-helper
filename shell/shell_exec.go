package shell

import (
	"fmt"
	"log"
	"strings"

	"config-helper/config"
	"config-helper/sshclient"
)

// ShellExecTask represents a shell execution task.
type ShellExecTask struct {
	Command string
}

type ShellExecBatchTask struct {
	Commands []string
}

// NewReplaceTask creates a new ReplaceTask
func NewShellExecTask(params config.TaskParameters) (*ShellExecTask, error) {
	command, ok := params["command"].(string)
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

// NewReplaceTask creates a new ReplaceTask
func NewShellExecBatchTask(params config.TaskParameters) (*ShellExecBatchTask, error) {
	cmds, _ := params["commands"].([]interface{})
	commandStrings := make([]string, len(cmds))
	return &ShellExecBatchTask{Commands: commandStrings}, nil
}

// Ensure ShellExecTask implements Task interface.
func (t *ShellExecBatchTask) Validate() error {
	if len(t.Commands) == 0 {
		return fmt.Errorf("commands cannot be empty")
	}
	return nil
}

func (t *ShellExecBatchTask) Execute(client *sshclient.Client) error {
	for _, cmd := range t.Commands {
		output, err := client.RunCommand(cmd)
		if err != nil {
			return fmt.Errorf("failed to execute command '%s': %v", cmd, err)
		}
		log.Printf("Command '%s' output: %s", cmd, output)
	}
	return nil
}
