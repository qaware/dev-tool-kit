package ui

import "github.com/qaware/dev-tool-kit/backend/core"

func handleJsonEvent(value string) (string, error) {
	return core.FormatJson(value)
}
