package task

import (
	"config-helper/config"
	"config-helper/dirs"
	"config-helper/files"
	"config-helper/shell"
	"config-helper/sshclient"
	"fmt"
)

// Task is an interface that all specific tasks will implement
type Task interface {
	Execute(client *sshclient.Client) error
	Validate() error
}

// Factory function to create tasks
func NewTask(category string, taskType string, params config.TaskParameters) (Task, error) {
	switch category {
	case "dirs":
		return newDirTask(taskType, params)
	case "files":
		return newFileTask(taskType, params)
	case "networking":
		return newNetworkTask(taskType, params)
	case "shell":
		return newShellTask(taskType, params)
	// Add other categories here
	default:
		return nil, fmt.Errorf("unknown task category: %s", category)
	}
}

func newDirTask(taskType string, params config.TaskParameters) (Task, error) {
	switch taskType {
	case "ensureDir":
		return dirs.NewEnsureDirTask(params)
	// Add other dir tasks here
	default:
		return nil, fmt.Errorf("unknown dir task type: %s", taskType)
	}
}

func newFileTask(taskType string, params config.TaskParameters) (Task, error) {
	switch taskType {
	case "lineInFile":
		return files.NewLineInFileTask(params)
	case "replaceInFile":
		return files.NewReplaceInFileTask(params)
	case "fileUpload":
		return files.NewFileUploadTask(params)
	// Add other file tasks here
	default:
		return nil, fmt.Errorf("unknown file task type: %s", taskType)
	}
}

func newNetworkTask(taskType string, params config.TaskParameters) (Task, error) {
	// Define and implement network tasks similarly
	return nil, fmt.Errorf("network tasks not implemented")
}

func newShellTask(taskType string, params config.TaskParameters) (Task, error) {
	switch taskType {
	case "shellExec":
		return shell.NewShellExecTask(params)
	case "shellExecBatch":
		return shell.NewShellExecBatchTask(params)
	// Add other file tasks here
	default:
		return nil, fmt.Errorf("unknown shell task type: %s", taskType)
	}
}
