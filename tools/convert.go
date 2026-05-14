package tools

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

func StringToUUID(s string) (uuid.UUID, error) {
	formattedUUID, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, err
	}

	return formattedUUID, nil
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
