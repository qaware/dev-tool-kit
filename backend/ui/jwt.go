package ui

import (
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleJwtEvent(values []string) (string, error) {
	valid, payload, err := core.DecodeJwt(values[0], values[1])
	if err != nil {
		return payload, err
	}

	if valid {
		return payload, &core.Information{"Signature successfully verified"}
	} else {
		return payload, &core.Information{"Invalid signature"}
	}
}
