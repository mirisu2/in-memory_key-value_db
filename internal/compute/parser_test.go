package compute

import (
	"testing"
)

func TestParseValidCommand(t *testing.T) {
	request := "SET key value"
	command, args, err := Parse(request)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if command != "SET" {
		t.Errorf("Expected command 'SET', got '%s'", command)
	}
	if len(args) != 2 || args[0] != "key" || args[1] != "value" {
		t.Errorf("Expected args ['key', 'value'], got %v", args)
	}
}

func TestParseInvalidCommand(t *testing.T) {
	request := "UPDATE key value"
	_, _, err := Parse(request)
	if err == nil {
		t.Error("Expected error for invalid command, got nil")
	}
}

func TestParseNotEnoughArguments(t *testing.T) {
	request := "GET"
	_, _, err := Parse(request)
	if err == nil {
		t.Error("Expected error for not enough arguments, got nil")
	}
}

func TestParseInvalidCharacters(t *testing.T) {
	request := "SET key! value?"
	_, _, err := Parse(request)
	if err == nil {
		t.Error("Expected error for invalid characters in arguments, got nil")
	}
}

func TestParseNoArguments(t *testing.T) {
	request := "SET"
	_, _, err := Parse(request)
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}

func TestParseExtraSpaces(t *testing.T) {
	request := "  SET   key    value  "
	command, args, err := Parse(request)
	if err != nil {
		t.Fatalf("Parse failed with extra spaces: %v", err)
	}
	if command != "SET" {
		t.Errorf("Expected command 'SET', got '%s'", command)
	}
	if len(args) != 2 || args[0] != "key" || args[1] != "value" {
		t.Errorf("Expected args ['key', 'value'], got %v", args)
	}
}
