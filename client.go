package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func RunClient(user, url string, key *Key) error {
	// SSH client configuration
	config := &ssh.ClientConfig{
		// User: "your_username",
		// Auth: []ssh.AuthMethod{
		// ssh.Password("your_password"), // Placeholder for simplicity
		// },
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // For testing purposes only
	}

	// Connect to the SSH server
	conn, err := ssh.Dial("tcp", url, config)
	if err != nil {
		return fmt.Errorf("SSH dial %s: %w", url, err)
	}
	defer conn.Close()

	chSSH, chReq, err := conn.OpenChannel("stream", nil)
	if err != nil {
		return fmt.Errorf("SSH open channel: %w", err)
	}
	defer chSSH.Close()

	go ssh.DiscardRequests(chReq)

	chSSH.Write([]byte("TEST\n"))
	var buf []byte
	_, err = chSSH.Read(buf)
	if err != nil {
		return fmt.Errorf("SSH channel read: %w", err)
	}
	if string(buf) != "TEST\n" {
		return fmt.Errorf("SSH channel read: unexpected output '%s'", buf)
	}

	return nil
}
