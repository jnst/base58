package base58

import (
	"errors"
	"math/big"
	"strings"
	"sync"
)

const (
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	base58   = 58
	// Buffer size calculation: log(256)/log(58) ≈ 1.3658
	bufferSizeMultiplier = 1366
	bufferSizeDivisor    = 1000
	bufferSizeExtra      = 2
)

var alphabetMap = make(map[byte]int)

// Pool for reusing big.Int objects to reduce allocations
var bigIntPool = sync.Pool{
	New: func() any {
		return new(big.Int)
	},
}

// Pool for reusing strings.Builder objects
var stringBuilderPool = sync.Pool{
	New: func() any {
		return &strings.Builder{}
	},
}

func init() {
	for i, char := range []byte(alphabet) {
		alphabetMap[char] = i
	}
}

// getBigInt gets a big.Int from the pool
func getBigInt() *big.Int {
	bi := bigIntPool.Get().(*big.Int) //nolint:errcheck
	return bi
}

// putBigInt returns a big.Int to the pool after resetting it
func putBigInt(bi *big.Int) {
	bi.SetInt64(0) // Reset to zero
	bigIntPool.Put(bi)
}

// getStringBuilder gets a strings.Builder from the pool
func getStringBuilder() *strings.Builder {
	sb := stringBuilderPool.Get().(*strings.Builder) //nolint:errcheck
	return sb
}

// putStringBuilder returns a strings.Builder to the pool after resetting it
func putStringBuilder(sb *strings.Builder) {
	sb.Reset()
	stringBuilderPool.Put(sb)
}

// calculateOptimalBufferSize calculates a more accurate buffer size
func calculateOptimalBufferSize(dataLen int) int {
	if dataLen == 0 {
		return 0
	}
	// More precise calculation: log(256)/log(58) ≈ 1.3658
	return (dataLen*bufferSizeMultiplier)/bufferSizeDivisor + bufferSizeExtra
}

// Encode encodes byte data to Base58 string using optimized implementation
func Encode(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	// Count leading zeros
	leading := 0
	for leading < len(data) && data[leading] == 0 {
		leading++
	}

	// Special case: all zeros
	if leading == len(data) {
		sb := getStringBuilder()
		defer putStringBuilder(sb)

		sb.Grow(leading)
		for i := 0; i < leading; i++ {
			sb.WriteByte(alphabet[0])
		}
		return sb.String()
	}

	// Calculate optimal buffer size
	size := calculateOptimalBufferSize(len(data) - leading)
	encoded := make([]byte, size)

	// Get big.Int objects from pool
	bigInt := getBigInt()
	baseInt := getBigInt()
	zero := getBigInt()
	mod := getBigInt()

	defer func() {
		putBigInt(bigInt)
		putBigInt(baseInt)
		putBigInt(zero)
		putBigInt(mod)
	}()

	bigInt.SetBytes(data[leading:])
	baseInt.SetInt64(base58)
	zero.SetInt64(0)

	pos := size - 1
	for bigInt.Cmp(zero) > 0 {
		bigInt.DivMod(bigInt, baseInt, mod)
		encoded[pos] = alphabet[mod.Int64()]
		pos--
	}

	// Build result string efficiently
	sb := getStringBuilder()
	defer putStringBuilder(sb)

	// Pre-allocate capacity
	resultLen := leading + (size - pos - 1)
	sb.Grow(resultLen)

	// Add leading zeros
	for i := 0; i < leading; i++ {
		sb.WriteByte(alphabet[0])
	}

	// Add encoded part
	sb.Write(encoded[pos+1:])

	return sb.String()
}

// Decode decodes Base58 string to byte data using optimized implementation
func Decode(s string) ([]byte, error) {
	if s == "" {
		return []byte{}, nil
	}

	// Count leading '1's
	leading := 0
	for leading < len(s) && s[leading] == alphabet[0] {
		leading++
	}

	// Get big.Int objects from pool
	bigInt := getBigInt()
	baseInt := getBigInt()
	temp := getBigInt()

	defer func() {
		putBigInt(bigInt)
		putBigInt(baseInt)
		putBigInt(temp)
	}()

	bigInt.SetInt64(0)
	baseInt.SetInt64(base58)

	// Process non-leading characters
	for _, char := range []byte(s[leading:]) {
		value, ok := alphabetMap[char]
		if !ok {
			return nil, errors.New("invalid base58 character")
		}
		bigInt.Mul(bigInt, baseInt)
		temp.SetInt64(int64(value))
		bigInt.Add(bigInt, temp)
	}

	decoded := bigInt.Bytes()

	// Prepare result with leading zeros
	result := make([]byte, leading+len(decoded))
	copy(result[leading:], decoded)

	return result, nil
}
