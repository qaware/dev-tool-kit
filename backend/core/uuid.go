package core

import (
	"errors"
	"github.com/google/uuid"
)

func GenerateUuids() (string, error) {
	text := ""
	n := 0
	for n < 10 {
		id, err := generateUuid()
		if err != nil {
			DebugError(err)
			return "", errors.New("Error generating a UUID")
		}
		text += id + "\n"
		n++
	}
	return text, nil
}

func generateUuid() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
