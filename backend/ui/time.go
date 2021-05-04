package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleTimeEvent(action string, value string) (string, error) {
	switch action {
	case "load":
		return core.GetNow(), nil
	case "convertDateTime":
		return core.ConvertToTimestamp(value)
	case "convertTimestamp":
		return core.ConvertToTimeString(value)
	}
	return "", errors.New("Internal error")
}
