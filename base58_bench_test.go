package base58

import (
	"crypto/rand"
	"testing"
)

// Benchmark data generators
func generateRandomBytes(size int) []byte {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}

// Encode benchmarks for various data sizes
func BenchmarkEncode_8B(b *testing.B) {
	data := generateRandomBytes(8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkEncode_32B(b *testing.B) {
	data := generateRandomBytes(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkEncode_64B(b *testing.B) {
	data := generateRandomBytes(64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkEncode_256B(b *testing.B) {
	data := generateRandomBytes(256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkEncode_1KB(b *testing.B) {
	data := generateRandomBytes(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkEncode_4KB(b *testing.B) {
	data := generateRandomBytes(4096)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkEncode_16KB(b *testing.B) {
	data := generateRandomBytes(16384)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

// Decode benchmarks for various string sizes
func BenchmarkDecode_8B(b *testing.B) {
	data := generateRandomBytes(8)
	encoded := Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

func BenchmarkDecode_32B(b *testing.B) {
	data := generateRandomBytes(32)
	encoded := Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

func BenchmarkDecode_64B(b *testing.B) {
	data := generateRandomBytes(64)
	encoded := Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

func BenchmarkDecode_256B(b *testing.B) {
	data := generateRandomBytes(256)
	encoded := Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

func BenchmarkDecode_1KB(b *testing.B) {
	data := generateRandomBytes(1024)
	encoded := Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

func BenchmarkDecode_4KB(b *testing.B) {
	data := generateRandomBytes(4096)
	encoded := Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

func BenchmarkDecode_16KB(b *testing.B) {
	data := generateRandomBytes(16384)
	encoded := Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

// Round trip benchmarks
func BenchmarkRoundTrip_32B(b *testing.B) {
	data := generateRandomBytes(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded := Encode(data)
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

func BenchmarkRoundTrip_256B(b *testing.B) {
	data := generateRandomBytes(256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded := Encode(data)
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

func BenchmarkRoundTrip_1KB(b *testing.B) {
	data := generateRandomBytes(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded := Encode(data)
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

// Special case benchmarks
func BenchmarkEncode_EmptyData(b *testing.B) {
	data := []byte{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkEncode_AllZeros(b *testing.B) {
	data := make([]byte, 32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkEncode_AllOnes(b *testing.B) {
	data := make([]byte, 32)
	for i := range data {
		data[i] = 0xFF
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkDecode_EmptyString(b *testing.B) {
	encoded := ""
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

func BenchmarkDecode_AllOnes(b *testing.B) {
	encoded := "111111111111111111111111111111111111111111111"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}

// Memory allocation benchmarks
func BenchmarkEncode_32B_Allocs(b *testing.B) {
	data := generateRandomBytes(32)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encode(data)
	}
}

func BenchmarkDecode_32B_Allocs(b *testing.B) {
	data := generateRandomBytes(32)
	encoded := Encode(data)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded) //nolint:errcheck
	}
}
