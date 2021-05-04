package ui

import (
	"encoding/json"
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleCustomizeEvent() (string, error) {
	pages := core.ReadCustomPagesXmlFile()
	if pages == nil {
		return "", nil
	}

	raw, err := json.Marshal(pages)
	if err != nil {
		return "", errors.New("Internal error")
	}

	return string(raw), nil
}
