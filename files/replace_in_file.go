package files

import (
	"config-helper/config"
	"config-helper/sshclient"
	"fmt"
)

// ReplaceTask replaces a pattern in a file with a new string
type ReplaceInFileTask struct {
	FilePath   string
	OldPattern string
	NewPattern string
}

// NewReplaceTask creates a new ReplaceTask
func NewReplaceInFileTask(params config.TaskParameters) (*ReplaceInFileTask, error) {
	filePath, ok := params["filePath"].(string)
	if !ok {
		return nil, fmt.Errorf("missing parameter: filePath")
	}
	oldPattern, ok := params["oldPattern"].(string)
	if !ok {
		return nil, fmt.Errorf("missing parameter: oldPattern")
	}
	newPattern, ok := params["newPattern"].(string)
	if !ok {
		return nil, fmt.Errorf("missing parameter: newPattern")
	}
	return &ReplaceInFileTask{FilePath: filePath, OldPattern: oldPattern, NewPattern: newPattern}, nil
}

// Execute executes the replace task
func (t *ReplaceInFileTask) Execute(client *sshclient.Client) error {
	cmd := fmt.Sprintf("sed -i 's/%s/%s/g' %s", t.OldPattern, t.NewPattern, t.FilePath)
	if _, err := client.RunCommand(cmd); err != nil {
		return fmt.Errorf("failed to replace pattern in file: %w", err)
	}
	return nil
}

// Validate validates the parameters of the task
func (t *ReplaceInFileTask) Validate() error {
	if t.FilePath == "" || t.OldPattern == "" || t.NewPattern == "" {
		return fmt.Errorf("invalid parameters")
	}
	return nil
}
