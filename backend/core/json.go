package core

import (
	"bytes"
	"encoding/json"
	"errors"
)

func FormatJson(raw string) (string, error) {
	if raw == "" {
		return "", nil
	}

	var result bytes.Buffer
	err := json.Indent(&result, []byte(raw), "", "    ")
	if err != nil {
		DebugError(err)
		return raw, errors.New("Invalid JSON")
	}

	return string(result.Bytes()), nil
}

func IsJson(input string) bool {
	return json.Valid([]byte(input))
}
