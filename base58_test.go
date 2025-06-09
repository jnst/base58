package base58

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty data",
			input:    []byte{},
			expected: "",
		},
		{
			name:     "single zero byte",
			input:    []byte{0x00},
			expected: "1",
		},
		{
			name:     "multiple zero bytes",
			input:    []byte{0x00, 0x00, 0x00},
			expected: "111",
		},
		{
			name:     "hello world",
			input:    []byte("Hello World"),
			expected: "JxF12TrwUP45BMd",
		},
		{
			name:     "binary data",
			input:    []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
			expected: "17bWpTW",
		},
		{
			name:     "single byte 1",
			input:    []byte{0x01},
			expected: "2",
		},
		{
			name:     "bitcoin address example",
			input:    []byte{0x00, 0x01, 0x09, 0x66, 0x77, 0x60, 0x06, 0x95, 0x3D, 0x55, 0x67, 0x43, 0x9E, 0x5E, 0x39, 0xF8, 0x6A, 0x0D, 0x27, 0x3B, 0xEE},
			expected: "1qb3y62fmEEVTPySXPQ77WXok6H",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Encode(tt.input)
			if result != tt.expected {
				t.Errorf("Encode(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
		hasError bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []byte{},
			hasError: false,
		},
		{
			name:     "single 1",
			input:    "1",
			expected: []byte{0x00},
			hasError: false,
		},
		{
			name:     "multiple 1s",
			input:    "111",
			expected: []byte{0x00, 0x00, 0x00},
			hasError: false,
		},
		{
			name:     "hello world",
			input:    "JxF12TrwUP45BMd",
			expected: []byte("Hello World"),
			hasError: false,
		},
		{
			name:     "binary data",
			input:    "17bWpTW",
			expected: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
			hasError: false,
		},
		{
			name:     "single 2",
			input:    "2",
			expected: []byte{0x01},
			hasError: false,
		},
		{
			name:     "bitcoin address example",
			input:    "1qb3y62fmEEVTPySXPQ77WXok6H",
			expected: []byte{0x00, 0x01, 0x09, 0x66, 0x77, 0x60, 0x06, 0x95, 0x3D, 0x55, 0x67, 0x43, 0x9E, 0x5E, 0x39, 0xF8, 0x6A, 0x0D, 0x27, 0x3B, 0xEE},
			hasError: false,
		},
		{
			name:     "invalid character 0",
			input:    "1230",
			expected: nil,
			hasError: true,
		},
		{
			name:     "invalid character O",
			input:    "123O",
			expected: nil,
			hasError: true,
		},
		{
			name:     "invalid character I",
			input:    "123I",
			expected: nil,
			hasError: true,
		},
		{
			name:     "invalid character l",
			input:    "123l",
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Decode(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("Decode(%q) expected error, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("Decode(%q) unexpected error: %v", tt.input, err)
				return
			}

			if len(result) != len(tt.expected) {
				t.Errorf("Decode(%q) length = %d, want %d", tt.input, len(result), len(tt.expected))
				return
			}

			for i, b := range result {
				if b != tt.expected[i] {
					t.Errorf("Decode(%q) = %v, want %v", tt.input, result, tt.expected)
					break
				}
			}
		})
	}
}

func TestEncodeDecodeRoundTrip(t *testing.T) {
	tests := [][]byte{
		{},
		{0x00},
		{0x01},
		{0x00, 0x00, 0x00},
		[]byte("Hello World"),
		{0xFF, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA},
		{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
	}

	for _, original := range tests {
		t.Run("", func(t *testing.T) {
			encoded := Encode(original)
			decoded, err := Decode(encoded)

			if err != nil {
				t.Errorf("Decode error: %v", err)
				return
			}

			if len(decoded) != len(original) {
				t.Errorf("Round trip length mismatch: %d != %d", len(decoded), len(original))
				return
			}

			for i, b := range decoded {
				if b != original[i] {
					t.Errorf("Round trip mismatch at index %d: %02x != %02x", i, b, original[i])
					break
				}
			}
		})
	}
}
