package utils

import "math/big"

const (
	base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// EncodeBase62 encodes a uint64 number into Base62 string
func EncodeBase62(number uint64) string {
	if number == 0 {
		return "0"
	}

	base := big.NewInt(62)
	var chars []byte

	for number > 0 {
		rem := new(big.Int).Mod(big.NewInt(int64(number)), base)
		chars = append([]byte{base62Alphabet[rem.Int64()]}, chars...)
		number = number / 62
	}

	return string(chars)
}
