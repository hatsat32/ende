package main

import (
	"html"
	"net/url"
	"testing"
)

func TestEnhtml(t *testing.T) {
	output, err := captureOutput(func() error {
		return enhtmlCmd.RunE(enhtmlCmd, []string{"<foo>"})
	})
	if err != nil {
		t.Fatalf("enhtml failed: %v", err)
	}
	expected := html.EscapeString("<foo>")
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestDehtml(t *testing.T) {
	encoded := html.EscapeString("<foo>")
	output, err := captureOutput(func() error {
		return dehtmlCmd.RunE(dehtmlCmd, []string{encoded})
	})
	if err != nil {
		t.Fatalf("dehtml failed: %v", err)
	}
	if output != "<foo>" {
		t.Errorf("Expected <foo>, got %s", output)
	}
}

func TestEnurl(t *testing.T) {
	output, err := captureOutput(func() error {
		return enurlCmd.RunE(enurlCmd, []string{"foo bar"})
	})
	if err != nil {
		t.Fatalf("enurl failed: %v", err)
	}
	expected := url.QueryEscape("foo bar")
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}

func TestDeurl(t *testing.T) {
	encoded := url.QueryEscape("foo bar")
	output, err := captureOutput(func() error {
		return deurlCmd.RunE(deurlCmd, []string{encoded})
	})
	if err != nil {
		t.Fatalf("deurl failed: %v", err)
	}
	if output != "foo bar" {
		t.Errorf("Expected foo bar, got %s", output)
	}
}
