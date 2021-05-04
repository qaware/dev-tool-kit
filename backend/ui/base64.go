package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleBase64Event(action string, value string) (string, error) {
	switch action {
	case "encode":
		return core.Base64Encode(value), nil
	case "decode":
		return core.Base64Decode(value)
	}
	return "", errors.New("Internal error")
}
