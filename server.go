package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func RunServer(url string, key *Key) error {
	// Configure the SSH server
	config := &ssh.ServerConfig{
		NoClientAuth: true, // Disable client authentication for simplicity
	}
	config.AddHostKey(key.pvt)

	// Start listening for incoming connections
	listener, err := net.Listen("tcp", url)
	if err != nil {
		return fmt.Errorf("tcp listen [%s]: %w", url, err)
	}
	log.Printf("tcp listen: %s\n", url)

	for {
		// Accept an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("tcp accept: %w", err)
		}

		// Perform SSH handshake
		sshConn, chans, reqs, err := ssh.NewServerConn(conn, config)
		if err != nil {
			log.Printf("ssh handshake: %v", err)
			continue
		}
		log.Printf("ssh connection: %s (%s)", sshConn.RemoteAddr(), sshConn.ClientVersion())

		// Discard all global requests
		go ssh.DiscardRequests(reqs)

		// Handle channels
		go handleChannels(chans)
	}
}

func handleChannels(chans <-chan ssh.NewChannel) {
	for newChannel := range chans {
		// Accept the channel
		channel, reqs, err := newChannel.Accept()
		if err != nil {
			log.Printf("ssh connection: %v", err)
			continue
		}

		go ssh.DiscardRequests(reqs)

		// Echo back the data
		go io.Copy(channel, channel)
	}
}
