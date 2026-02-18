package main

import (
	"encoding/hex"
	"testing"
)

func TestEnhex(t *testing.T) {
	output, err := captureOutput(func() error {
		return enhexCmd.RunE(enhexCmd, []string{"hello"})
	})
	if err != nil {
		t.Fatalf("enhex failed: %v", err)
	}
	expected := hex.EncodeToString([]byte("hello"))
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestDehex(t *testing.T) {
	encoded := hex.EncodeToString([]byte("hello"))
	output, err := captureOutput(func() error {
		return dehexCmd.RunE(dehexCmd, []string{encoded})
	})
	if err != nil {
		t.Fatalf("dehex failed: %v", err)
	}
	if output != "hello" {
		t.Errorf("Expected hello, got %s", output)
	}
}

// Test alias commands
func TestEnb16(t *testing.T) {
	output, err := captureOutput(func() error {
		return enb16Cmd.RunE(enb16Cmd, []string{"hello"})
	})
	if err != nil {
		t.Fatalf("enb16 failed: %v", err)
	}
	expected := hex.EncodeToString([]byte("hello"))
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
