package ui

import (
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleAsciiEvent(value string) (string, error) {
	return core.GenerateAsciiArt(value)
}
