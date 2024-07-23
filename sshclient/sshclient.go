// sshclient/sshclient.go
package sshclient

import (
    "fmt"
    "golang.org/x/crypto/ssh"
    "io/ioutil"
)

// Client manages SSH connections.
type Client struct {
    *ssh.Client
    Config *ssh.ClientConfig
}

// NewClient creates a new SSH client.
func NewClient(user, keyPath, host string) (*Client, error) {
    // Load the private key
    key, err := ioutil.ReadFile(keyPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read private key: %w", err)
    }

    // Parse the private key
    signer, err := ssh.ParsePrivateKey(key)
    if err != nil {
        return nil, fmt.Errorf("failed to parse private key: %w", err)
    }

    // Create SSH client configuration
    config := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(signer),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    // Connect to the SSH server
    client, err := ssh.Dial("tcp", host, config)
    if err != nil {
        return nil, fmt.Errorf("failed to dial: %w", err)
    }

    return &Client{Client: client, Config: config}, nil
}

// RunCommand executes a command on the remote server.
func (c *Client) RunCommand(cmd string) (string, error) {
    session, err := c.NewSession()
    if err != nil {
        return "", fmt.Errorf("failed to create session: %w", err)
    }
    defer session.Close()

    output, err := session.CombinedOutput(cmd)
    if err != nil {
        return "", fmt.Errorf("failed to run command: %w", err)
    }

    return string(output), nil
}

