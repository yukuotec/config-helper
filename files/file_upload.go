package files

import (
	"fmt"
	"io/ioutil"
	"log"

	"config-helper/sshclient"
)

// FileUploadTask represents a file upload task.
type FileUploadTask struct {
	LocalPath  string
	RemotePath string
}

// NewFileUploadTask creates a new FileUploadTask
func NewFileUploadTask(params map[string]string) (*FileUploadTask, error) {
	localPath, ok := params["localPath"]
	if !ok {
		return nil, fmt.Errorf("missing parameter: localPath")
	}
	remotePath, ok := params["remotePath"]
	if !ok {
		return nil, fmt.Errorf("missing parameter: remotePath")
	}
	return &FileUploadTask{LocalPath: localPath, RemotePath: remotePath}, nil
}

// Ensure FileUploadTask implements Task interface.
func (t *FileUploadTask) Validate() error {
	if t.LocalPath == "" || t.RemotePath == "" {
		return fmt.Errorf("both localPath and remotePath must be specified")
	}
	return nil
}

func (t *FileUploadTask) Execute(client *sshclient.Client) error {
	fileContent, err := ioutil.ReadFile(t.LocalPath)
	if err != nil {
		return fmt.Errorf("failed to read file '%s': %v", t.LocalPath, err)
	}

	err = client.UploadFile(t.RemotePath, fileContent)
	if err != nil {
		return fmt.Errorf("failed to upload file to '%s': %v", t.RemotePath, err)
	}

	log.Printf("File '%s' uploaded to '%s'", t.LocalPath, t.RemotePath)
	return nil
}
