package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCLIEncode(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		input    string
		expected string
	}{
		{
			name:     "encode hello world argument",
			args:     []string{"encode", "Hello World"},
			expected: "JxF12TrwUP45BMd\n",
		},
		{
			name:     "encode empty string",
			args:     []string{"encode", ""},
			expected: "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			cmd.Dir = "./"

			var stdout bytes.Buffer
			cmd.Stdout = &stdout

			if tt.input != "" {
				cmd.Stdin = strings.NewReader(tt.input)
			}

			err := cmd.Run()
			if err != nil {
				t.Fatalf("Command failed: %v", err)
			}

			result := stdout.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestCLIDecode(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "decode hello world",
			args:     []string{"decode", "JxF12TrwUP45BMd"},
			expected: "Hello World",
		},
		{
			name:     "decode empty string",
			args:     []string{"decode", ""},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			cmd.Dir = "./"

			var stdout bytes.Buffer
			cmd.Stdout = &stdout

			err := cmd.Run()
			if err != nil {
				t.Fatalf("Command failed: %v", err)
			}

			result := stdout.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestCLIHelp(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"help command", []string{"help"}},
		{"help flag short", []string{"-h"}},
		{"help flag long", []string{"--help"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "main.go"}, tt.args...)...)
			cmd.Dir = "./"

			var stdout bytes.Buffer
			cmd.Stdout = &stdout

			err := cmd.Run()
			if err != nil {
				t.Fatalf("Command failed: %v", err)
			}

			result := stdout.String()
			if !strings.Contains(result, "base58 - Base58 encoding and decoding tool") {
				t.Errorf("Help output should contain tool description")
			}
			if !strings.Contains(result, "Usage:") {
				t.Errorf("Help output should contain usage information")
			}
		})
	}
}

func TestCLIFileInput(t *testing.T) {
	testFile := "test_input.txt"
	testContent := "Hello World"

	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	cmd := exec.Command("go", "run", "main.go", "encode", "-f", testFile)
	cmd.Dir = "./"

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err = cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	result := strings.TrimSpace(stdout.String())
	expected := "Rk9goP9W13SSc4HwybhK5wR"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestCLIInvalidCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "invalid")
	cmd.Dir = "./"

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err == nil {
		t.Fatalf("Expected command to fail")
	}

	result := stderr.String()
	if !strings.Contains(result, "Unknown command") {
		t.Errorf("Should show unknown command error")
	}
}

func TestCLIDecodeError(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "decode", "invalid0character")
	cmd.Dir = "./"

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err == nil {
		t.Fatalf("Expected command to fail")
	}

	result := stderr.String()
	if !strings.Contains(result, "Error:") {
		t.Errorf("Should show decode error")
	}
}
