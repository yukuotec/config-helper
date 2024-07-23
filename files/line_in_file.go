package files

import (
    "config-helper/sshclient"
    "fmt"
)

// LineInFileTask ensures that a line is present in a file
type LineInFileTask struct {
    FilePath string
    Line     string
}

// NewLineInFileTask creates a new LineInFileTask
func NewLineInFileTask(params map[string]string) (*LineInFileTask, error) {
    filePath, ok := params["filePath"]
    if !ok {
        return nil, fmt.Errorf("missing parameter: filePath")
    }
    line, ok := params["line"]
    if !ok {
        return nil, fmt.Errorf("missing parameter: line")
    }
    return &LineInFileTask{FilePath: filePath, Line: line}, nil
}

// Execute executes the line in file task
func (t *LineInFileTask) Execute(client *sshclient.Client) error {
    cmd := fmt.Sprintf("grep -qxF '%s' %s || echo '%s' >> %s", t.Line, t.FilePath, t.Line, t.FilePath)
    if _, err := client.RunCommand(cmd); err != nil {
        return fmt.Errorf("failed to add line to file: %w", err)
    }
    return nil
}

// Validate validates the parameters of the task
func (t *LineInFileTask) Validate() error {
    if t.FilePath == "" || t.Line == "" {
        return fmt.Errorf("invalid parameters")
    }
    return nil
}

