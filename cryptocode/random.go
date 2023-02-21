package cryptocode

import (
    "crypto/rand"
    "math/big"
)

const (
	STANDARD_SET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func SecureRandomString(length int, symbols string) (string, error) {
    bytes := make([]byte, length)
    numSymbols := big.NewInt(int64(len(symbols)))

    for i := 0; i < length; i++ {
        randomByte, err := rand.Int(rand.Reader, numSymbols)
        if err != nil {
            return "", err
        }
        bytes[i] = symbols[randomByte.Int64()]
    }
    return string(bytes), nil
}