package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleGzipEvent(action string, value string) (string, error) {
	switch action {
	case "encode":
		return core.GzipEncode(value)
	case "decode":
		return core.GzipDecode(value)
	}
	return "", errors.New("Internal error")
}
