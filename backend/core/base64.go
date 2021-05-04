package core

import (
	"encoding/base64"
	"errors"
	"unicode"
)

func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Base64Decode(text string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		DebugError(err)
		return "", errors.New("Invalid base64 encoding")
	}
	return string(result), nil
}

func IsBase64EncodedLatin1(text string) bool {
	plain, err := Base64Decode(text)
	if err != nil {
		return false
	}

	for _, character := range plain {
		if character > unicode.MaxLatin1 {
			return false
		}
	}

	return true
}
