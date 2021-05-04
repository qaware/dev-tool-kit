package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleUrlEvent(action string, value string) (string, error) {
	switch action {
	case "encode":
		return core.UrlEncode(value), nil
	case "decode":
		return core.UrlDecode(value)
	}
	return "", errors.New("Internal error")
}
