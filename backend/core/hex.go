package core

import (
	"encoding/hex"
	"errors"
	"regexp"
	"strings"
)

var whitespacePattern = regexp.MustCompile("\\s+")

func HexEncode(text string) string {
	result := hex.EncodeToString([]byte(text))
	return strings.ToUpper(insertSpaces(result))
}

func HexDecode(text string) (string, error) {
	normalized := whitespacePattern.ReplaceAllString(strings.ToLower(strings.TrimSpace(text)), "")

	result, err := hex.DecodeString(normalized)
	if err != nil {
		DebugError(err)
		return "", errors.New("Invalid hex number")
	}
	return string(result), nil
}

func insertSpaces(text string) string {
	buffer := ""
	for i, char := range text {
		buffer += string(char)
		if i%2 == 1 && i != len(text)-1 {
			buffer += " "
		}
	}
	return buffer
}
