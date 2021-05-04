package ui

import "github.com/qaware/dev-tool-kit/backend/core"

func handleUuidEvent() (string, error) {
	return core.GenerateUuids()
}
