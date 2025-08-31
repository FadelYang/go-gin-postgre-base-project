package tools

import "github.com/google/uuid"

func StringToUUID(s string) (uuid.UUID, error) {
	formattedUUID, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, err
	}

	return formattedUUID, nil
}
