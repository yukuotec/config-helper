package dirs

import (
    "config-helper/sshclient"
    "fmt"
)

// EnsureDirTask ensures that a directory exists with specified attributes
type EnsureDirTask struct {
    Path  string
    Owner string
    Mode  string
}

// NewEnsureDirTask creates a new EnsureDirTask
func NewEnsureDirTask(params map[string]string) (*EnsureDirTask, error) {
    path, ok := params["path"]
    if !ok {
        return nil, fmt.Errorf("missing parameter: path")
    }
    owner, ok := params["owner"]
    if !ok {
        return nil, fmt.Errorf("missing parameter: owner")
    }
    mode, ok := params["mode"]
    if !ok {
        return nil, fmt.Errorf("missing parameter: mode")
    }
    return &EnsureDirTask{Path: path, Owner: owner, Mode: mode}, nil
}

// Execute executes the ensure directory task
func (t *EnsureDirTask) Execute(client *sshclient.Client) error {
    cmd := fmt.Sprintf("mkdir -p %s && chown %s %s && chmod %s %s", t.Path, t.Owner, t.Path, t.Mode, t.Path)
    if _, err := client.RunCommand(cmd); err != nil {
        return fmt.Errorf("failed to ensure directory: %w", err)
    }
    return nil
}

// Validate validates the parameters of the task
func (t *EnsureDirTask) Validate() error {
    if t.Path == "" || t.Owner == "" || t.Mode == "" {
        return fmt.Errorf("invalid parameters")
    }
    return nil
}

