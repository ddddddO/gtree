package main

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func init() {
	go func() {
		fmt.Println("launch DB container(expose 15432 port)")
		fmt.Println("※rm DB container after test..")

		cmd := exec.Command("docker", "run", "-p", "15432:5432", "postgres_pgdb")
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(10 * time.Second) // 微妙だ..
}

const (
	validDSN   = "host=localhost user=postgres port=15432"
	invalidDSN = "host=localhost user=postgres port=5432"
)

func TestConnect(t *testing.T) {
	// TDD
	tests := []struct {
		dsn       string
		wantError bool
	}{
		{validDSN, false},
		{invalidDSN, true},
	}
	for _, tt := range tests {
		_, err := connect(tt.dsn)
		if tt.wantError != (err != nil) {
			t.Fatalf("wantError: %v pattern", tt.wantError)
		}
	}
}

func TestNotify(t *testing.T) {
	conn, err := connect(validDSN)
	if err != nil {
		t.Fatalf("failed to connect DB")
	}

	tests := []struct {
		wantError bool
	}{
		{false},
	}
	for _, tt := range tests {
		err = notify(conn)
		if tt.wantError != (err != nil) {
			t.Fatalf("wantError: %v pattern", tt.wantError)
		}
	}
}
