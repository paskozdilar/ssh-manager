package main_test

import (
	"testing"

	sshmanager "github.com/paskozdilar/ssh-manager"
)

func TestMain(t *testing.T) {
	url := "localhost:8080"

	key, err := sshmanager.NewKey()
	if err != nil {
		t.Fatalf("create key: %v", err)
	}

	// TODO: make server cancellable
	go sshmanager.RunServer(url, key)

	t.Run("client", func(t *testing.T) {
		t.Parallel()
		err := sshmanager.RunClient("user", url, key)
		if err != nil {
			t.Fatalf("client: %v", err)
		}
	})
}
