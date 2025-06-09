package base58

import (
	"crypto/rand"
	"testing"
)

// Test correctness of optimized functions
func TestOptimizedCorrectness(t *testing.T) {
	testCases := [][]byte{
		{},
		{0x00},
		{0x00, 0x00, 0x00},
		[]byte("Hello World"),
		{0x01},
		{0xFF, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA},
		{0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
	}

	for i, data := range testCases {
		t.Run("", func(t *testing.T) {
			// Test encode
			original := Encode(data)
			optimized := EncodeOptimized(data)
			
			if original != optimized {
				t.Errorf("Test %d: Encode mismatch. Original: %q, Optimized: %q", i, original, optimized)
			}

			// Test decode
			originalDecoded, err1 := Decode(original)
			optimizedDecoded, err2 := DecodeOptimized(original)
			
			if err1 != nil || err2 != nil {
				t.Errorf("Test %d: Decode error. Original: %v, Optimized: %v", i, err1, err2)
				return
			}
			
			if len(originalDecoded) != len(optimizedDecoded) {
				t.Errorf("Test %d: Decode length mismatch. Original: %d, Optimized: %d", i, len(originalDecoded), len(optimizedDecoded))
				return
			}
			
			// Check if optimized decode matches original decode
			for j, b := range originalDecoded {
				if b != optimizedDecoded[j] {
					t.Errorf("Test %d: Decode byte mismatch at index %d. Original: %02x, Optimized: %02x", i, j, b, optimizedDecoded[j])
					break
				}
			}
			
			// Check if decoded data matches original input
			if len(originalDecoded) != len(data) {
				t.Errorf("Test %d: Round trip failed. Expected length: %d, Got: %d", i, len(data), len(originalDecoded))
				return
			}
			
			for j, b := range data {
				if b != originalDecoded[j] {
					t.Errorf("Test %d: Round trip failed at index %d. Expected: %02x, Got: %02x", i, j, b, originalDecoded[j])
					break
				}
			}
		})
	}
}

// Fuzz test with random data
func TestOptimizedFuzz(t *testing.T) {
	for i := 0; i < 1000; i++ {
		// Generate random data of random size (1-1024 bytes)
		size := 1 + i%1024
		data := make([]byte, size)
		_, err := rand.Read(data)
		if err != nil {
			t.Fatalf("Failed to generate random data: %v", err)
		}

		// Test encode
		original := Encode(data)
		optimized := EncodeOptimized(data)
		
		if original != optimized {
			t.Errorf("Fuzz %d: Encode mismatch for data %x. Original: %q, Optimized: %q", i, data, original, optimized)
		}

		// Test decode
		originalDecoded, err1 := Decode(original)
		optimizedDecoded, err2 := DecodeOptimized(original)
		
		if err1 != nil || err2 != nil {
			t.Errorf("Fuzz %d: Decode error. Original: %v, Optimized: %v", i, err1, err2)
			continue
		}
		
		// Check if both decoded results match original data
		if len(originalDecoded) != len(data) || len(optimizedDecoded) != len(data) {
			t.Errorf("Fuzz %d: Round trip failed. Original length: %d, Optimized length: %d, Expected: %d", 
				i, len(originalDecoded), len(optimizedDecoded), len(data))
			continue
		}
		
		// Check if optimized decode matches original decode
		for j, b := range originalDecoded {
			if b != optimizedDecoded[j] {
				t.Errorf("Fuzz %d: Decode mismatch at index %d. Original: %02x, Optimized: %02x", i, j, b, optimizedDecoded[j])
				break
			}
		}
		
		// Check if decoded data matches original input
		for j, b := range data {
			if b != originalDecoded[j] {
				t.Errorf("Fuzz %d: Round trip failed at index %d. Expected: %02x, Got: %02x", i, j, b, originalDecoded[j])
				break
			}
		}
	}
}

// Benchmark optimized encode functions
func BenchmarkEncodeOptimized_8B(b *testing.B) {
	data := generateRandomBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EncodeOptimized(data)
	}
}

func BenchmarkEncodeOptimized_32B(b *testing.B) {
	data := generateRandomBytes(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EncodeOptimized(data)
	}
}

func BenchmarkEncodeOptimized_256B(b *testing.B) {
	data := generateRandomBytes(256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EncodeOptimized(data)
	}
}

func BenchmarkEncodeOptimized_1KB(b *testing.B) {
	data := generateRandomBytes(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EncodeOptimized(data)
	}
}

func BenchmarkEncodeOptimized_4KB(b *testing.B) {
	data := generateRandomBytes(4096)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EncodeOptimized(data)
	}
}

// Benchmark optimized decode functions
func BenchmarkDecodeOptimized_32B(b *testing.B) {
	data := generateRandomBytes(32)
	encoded := EncodeOptimized(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecodeOptimized(encoded)
	}
}

func BenchmarkDecodeOptimized_256B(b *testing.B) {
	data := generateRandomBytes(256)
	encoded := EncodeOptimized(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecodeOptimized(encoded)
	}
}

func BenchmarkDecodeOptimized_1KB(b *testing.B) {
	data := generateRandomBytes(1024)
	encoded := EncodeOptimized(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecodeOptimized(encoded)
	}
}

// Memory allocation benchmarks
func BenchmarkEncodeOptimized_32B_Allocs(b *testing.B) {
	data := generateRandomBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EncodeOptimized(data)
	}
}

func BenchmarkDecodeOptimized_32B_Allocs(b *testing.B) {
	data := generateRandomBytes(32)
	encoded := EncodeOptimized(data)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecodeOptimized(encoded)
	}
}

// Comparison benchmarks - Original vs Optimized
func BenchmarkCompare_Encode_32B(b *testing.B) {
	data := generateRandomBytes(32)
	
	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Encode(data)
		}
	})
	
	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = EncodeOptimized(data)
		}
	})
}

func BenchmarkCompare_Encode_1KB(b *testing.B) {
	data := generateRandomBytes(1024)
	
	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Encode(data)
		}
	})
	
	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = EncodeOptimized(data)
		}
	})
}

func BenchmarkCompare_Decode_32B(b *testing.B) {
	data := generateRandomBytes(32)
	encoded := Encode(data)
	
	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = Decode(encoded)
		}
	})
	
	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = DecodeOptimized(encoded)
		}
	})
}

func BenchmarkCompare_Decode_1KB(b *testing.B) {
	data := generateRandomBytes(1024)
	encoded := Encode(data)
	
	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = Decode(encoded)
		}
	})
	
	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = DecodeOptimized(encoded)
		}
	})
}