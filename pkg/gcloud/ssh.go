package gcloud

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"sync"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

type LoggerFunction func(stdout, stderr string)

type SSHSession struct {
	*ssh.Session
	StdoutChan chan string
	StderrChan chan string
}
type SSHClient struct {
	sshClient *ssh.Client

	// allows only one active ssh session per client
	sshSessionMutex sync.Mutex
}

func (c *SSHClient) Close() error {
	return c.sshClient.Close()
}

func (c *SSHClient) openSSHSession() (*SSHSession, error) {
	session, err := c.sshClient.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create ssh session: %w", err)
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		return nil, MaybeMultiError(fmt.Errorf("failed to create stdout pipe: %w", err), session.Close())
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return nil, MaybeMultiError(fmt.Errorf("failed to create stderr pipe: %w", err), session.Close())
	}

	stdoutChan := make(chan string)
	stderrChan := make(chan string)
	go func() {
		logLineScanner := bufio.NewScanner(stdout)
		for logLineScanner.Scan() {
			stdoutChan <- logLineScanner.Text()
		}
		close(stdoutChan)
	}()
	go func() {
		logLineScanner := bufio.NewScanner(stderr)
		for logLineScanner.Scan() {
			stderrChan <- logLineScanner.Text()
		}
		close(stderrChan)
	}()
	return &SSHSession{
		Session:    session,
		StdoutChan: stdoutChan,
		StderrChan: stderrChan,
	}, nil
}

func (c *SSHClient) Run(ctx context.Context, loggerFn LoggerFunction, cmd string) error {
	c.sshSessionMutex.Lock()
	defer c.sshSessionMutex.Unlock()
	session, err := c.openSSHSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var stdioWg sync.WaitGroup
	stdioWg.Add(2)
	go func() {
		for outStr := range session.StdoutChan {
			loggerFn(outStr, "")
		}
		stdioWg.Done()
	}()

	go func() {
		for errStr := range session.StderrChan {
			loggerFn("", errStr)
		}
		stdioWg.Done()
	}()

	if err := session.Start(cmd); err != nil {
		return fmt.Errorf("failed to start command %s: %w", cmd, err)
	}

	waitErrCh := make(chan error, 1)
	go func() {
		waitErrCh <- session.Wait()
		close(waitErrCh)
	}()

	select {
	case <-ctx.Done():
		// send SIGINT to the process
		signalErr := session.Signal(ssh.SIGINT)
		// wait for termination
		waitErr := <-waitErrCh
		// wait for stdio to close
		stdioWg.Wait()
		return MaybeMultiError(ctx.Err(), signalErr, waitErr)
	case err := <-waitErrCh:
		stdioWg.Wait()
		return err
	}
}

func (c *SSHClient) CopyFile(ctx context.Context, data *bytes.Reader, remotePath, permission string) error {
	c.sshSessionMutex.Lock()
	defer c.sshSessionMutex.Unlock()
	scpClient, err := scp.NewClientBySSH(c.sshClient)
	if err != nil {
		return fmt.Errorf("failed to create scp client: %w", err)
	}
	defer scpClient.Close()
	return scpClient.Copy(ctx, data, remotePath, permission, data.Size())
}
