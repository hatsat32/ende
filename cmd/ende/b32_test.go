package main

import (
	"encoding/base32"
	"testing"
)

func TestEnb32(t *testing.T) {
	output, err := captureOutput(func() error {
		return enb32Cmd.RunE(enb32Cmd, []string{"hello"})
	})
	if err != nil {
		t.Fatalf("enb32 failed: %v", err)
	}
	expected := base32.StdEncoding.EncodeToString([]byte("hello"))
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestDeb32(t *testing.T) {
	encoded := base32.StdEncoding.EncodeToString([]byte("hello"))
	output, err := captureOutput(func() error {
		return deb32Cmd.RunE(deb32Cmd, []string{encoded})
	})
	if err != nil {
		t.Fatalf("deb32 failed: %v", err)
	}
	if output != "hello" {
		t.Errorf("Expected hello, got %s", output)
	}
}
