package main

import (
	"testing"
)

func TestEnrot13(t *testing.T) {
	output, err := captureOutput(func() error {
		return enrot13Cmd.RunE(enrot13Cmd, []string{"hello"})
	})
	if err != nil {
		t.Fatalf("enrot13 failed: %v", err)
	}
	// rot13(hello) = uryyb
	expected := "uryyb"
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestDerot13(t *testing.T) {
	output, err := captureOutput(func() error {
		return derot13Cmd.RunE(derot13Cmd, []string{"uryyb"})
	})
	if err != nil {
		t.Fatalf("derot13 failed: %v", err)
	}
	if output != "hello" {
		t.Errorf("Expected hello, got %s", output)
	}
}
