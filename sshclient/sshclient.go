package sshclient

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Client wraps an SSH client
type Client struct {
	client     *ssh.Client
	sftpClient *sftp.Client
}

// NewClient creates a new SSH client
func NewClient(host, user, keyPath string) (*Client, error) {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return nil, err
	}

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to create SFTP client: %v", err)
	}

	return &Client{
		client:     client,
		sftpClient: sftpClient,
	}, nil
}

// RunCommand runs a command on the remote machine
func (c *Client) RunCommand(cmd string) (string, error) {
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// Close closes the SSH connection
func (c *Client) Close() {
	if err := c.client.Close(); err != nil {
		log.Printf("Error closing SSH client: %v", err)
	}
	if err := c.sftpClient.Close(); err != nil {
		log.Printf("Error closing SFTP client: %v", err)
	}
}

// UploadFile uploads a file to the remote path.
func (c *Client) UploadFile(remotePath string, content []byte) error {
	file, err := c.sftpClient.Create(remotePath)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %v", remotePath, err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("failed to write to file '%s': %v", remotePath, err)
	}

	return nil
}
