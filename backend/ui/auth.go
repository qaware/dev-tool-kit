package ui

import (
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleAuthEvent(values []string) string {
	if len(values) < 2 {
		return ""
	}

	user := values[0]
	pass := values[1]
	if len(user) > 0 && len(pass) > 0 {
		return core.Base64Encode(user + ":" + pass)
	} else {
		return ""
	}
}
