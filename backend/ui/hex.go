package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleHexEvent(action string, value string) (string, error) {
	switch action {
	case "encode":
		return core.HexEncode(value), nil
	case "decode":
		return core.HexDecode(value)
	}
	return "", errors.New("Internal error")
}
