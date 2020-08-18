package remote

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/pkg/sftp"

	"replicate.ai/cli/pkg/console"
)

type Client struct {
	options    *Options
	sftpClient *sftp.Client
}

// TODO: intern connections keyed on options

// NewClient creates a new SSH client
func NewClient(options *Options) (*Client, error) {
	c := &Client{
		options: options,
	}

	var err error
	// FIXME (bfirsh): do we want to create an SFTP client for each SSH client? this could be done lazily in SFTP()
	c.sftpClient, err = makeSFTPClient(options)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// SFTP returns a https://godoc.org/github.com/pkg/sftp#Client,
// capable of issuing filesystem commands remotely.
// For example: client.SFTP().Mkdir("foo")
func (c *Client) SFTP() *sftp.Client {
	return c.sftpClient
}

func (c *Client) WriteFile(data []byte, path string) error {
	remoteFile, err := c.sftpClient.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create remote file at %s, got error: %s", path, err)
	}
	if _, err := remoteFile.Write(data); err != nil {
		return fmt.Errorf("Failed to write remote file at %s, got error: %s", path, err)
	}
	if err := remoteFile.Close(); err != nil {
		return fmt.Errorf("Failed to close remote file at %s, got error: %s", path, err)
	}

	return nil
}

func makeSFTPClient(options *Options) (*sftp.Client, error) {
	args := options.SSHArgs()
	args = append(args, options.Host, "-s", "sftp")
	sshCmd := exec.Command("ssh", args...)
	wr, err := sshCmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	rd, err := sshCmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr, err := sshCmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	// TODO(andreas): what happens if the connection is interrupted?
	// TODO(andreas): handle cmd.Stderr
	if err := sshCmd.Start(); err != nil {
		// TODO(andreas): more actionable error message
		return nil, fmt.Errorf("Failed to establish SSH connection to %s: %w", options.Host, err)
	}
	client, err := sftp.NewClientPipe(rd, wr)
	if err != nil {
		stderrOut, _ := ioutil.ReadAll(stderr)
		// TODO(bfirsh): make a nice error type for this
		if len(stderrOut) > 0 {
			console.Error(strings.TrimSpace(string(stderrOut)))
		}
		return nil, err
	}
	return client, nil
}
