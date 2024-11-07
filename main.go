package main

import (
	"log"
	"os"
)

func main() {
	key, err := NewKey()
	if err != nil {
		log.Fatalf("new key error: %v", err)
	}

	if len(os.Args) == 2 && os.Args[1] == "server" {
		err = RunServer("localhost:8080", key)
	} else {
		err = RunClient("user", "localhost:8080", key)
	}

	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
