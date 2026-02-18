package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"testing"
)

// Helper to capture stdout
func captureOutput(f func() error) (string, error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String(), err
}

func TestEnb64(t *testing.T) {
	// Test direct argument
	output, err := captureOutput(func() error {
		return enb64Cmd.RunE(enb64Cmd, []string{"hello"})
	})
	if err != nil {
		t.Fatalf("enb64 failed: %v", err)
	}
	expected := base64.StdEncoding.EncodeToString([]byte("hello"))
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestDeb64(t *testing.T) {
	encoded := base64.StdEncoding.EncodeToString([]byte("hello"))
	output, err := captureOutput(func() error {
		return deb64Cmd.RunE(deb64Cmd, []string{encoded})
	})
	if err != nil {
		t.Fatalf("deb64 failed: %v", err)
	}
	if output != "hello" {
		t.Errorf("Expected hello, got %s", output)
	}
}
