package common

import (
	"encoding/hex"
	"crypto/rand"
)

func GenUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// RFC 4122
	uuid[8] = 0x80	// variant bits see page 5
	uuid[4] = 0x40	// version 4 pseudo random, see page 7

	return hex.EncodeToString(uuid), nil
}
