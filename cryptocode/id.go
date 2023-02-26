package cryptocode

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func ID() (string, error) {
	randStr, err := SecureRandomString(50, STANDARD_SET)

	if err != nil {
		return "", err
	}

	fullHash := sha256.Sum256([]byte(randStr + strconv.FormatInt(time.Now().UnixNano(), 10)))

	return hex.EncodeToString(fullHash[:])[:20], nil
}
