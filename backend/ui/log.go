package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleLogEvent(action string, value string) (string, error) {
	switch action {
	case "load":
		return core.LoadLogFile()
	case "append":
		return core.PutLogMessage(value)
	case "save":
		core.WriteLogFile(value)
		return "", nil
	}
	return "", errors.New("Internal error")
}
