package base58

import (
	"errors"
	"math/big"
)

const (
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	base58   = 58
	// Buffer size calculation factor: log(256)/log(58) ≈ 1.38
	bufferSizeFactor  = 138
	bufferSizeDivisor = 100
)

var alphabetMap = make(map[byte]int)

func init() {
	for i, char := range []byte(alphabet) {
		alphabetMap[char] = i
	}
}

func Encode(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	leading := 0
	for leading < len(data) && data[leading] == 0 {
		leading++
	}

	if leading == len(data) {
		result := ""
		for i := 0; i < leading; i++ {
			result += string(alphabet[0])
		}
		return result
	}

	size := (len(data)-leading)*bufferSizeFactor/bufferSizeDivisor + 1
	encoded := make([]byte, size)

	bigInt := new(big.Int).SetBytes(data[leading:])
	baseInt := big.NewInt(base58)
	zero := big.NewInt(0)

	pos := size - 1
	for bigInt.Cmp(zero) > 0 {
		mod := new(big.Int)
		bigInt.DivMod(bigInt, baseInt, mod)
		encoded[pos] = alphabet[mod.Int64()]
		pos--
	}

	result := string(encoded[pos+1:])

	for i := 0; i < leading; i++ {
		result = string(alphabet[0]) + result
	}

	return result
}

func Decode(s string) ([]byte, error) {
	if s == "" {
		return []byte{}, nil
	}

	leading := 0
	for leading < len(s) && s[leading] == alphabet[0] {
		leading++
	}

	bigInt := big.NewInt(0)
	baseInt := big.NewInt(base58)
	for _, char := range []byte(s[leading:]) {
		value, ok := alphabetMap[char]
		if !ok {
			return nil, errors.New("invalid base58 character")
		}
		bigInt.Mul(bigInt, baseInt)
		bigInt.Add(bigInt, big.NewInt(int64(value)))
	}

	decoded := bigInt.Bytes()

	result := make([]byte, leading+len(decoded))
	copy(result[leading:], decoded)

	return result, nil
}
