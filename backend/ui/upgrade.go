package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
	"strconv"
)

func handleUpgradeEvent(action string) (string, error) {
	switch action {
	case "getVersion":
		return "Version " + core.GetVersion(), nil
	case "checkUpgrade":
		return strconv.FormatBool(core.HasNewVersion()), nil
	case "upgradeNow":
		err := core.UpgradeNow()
		if err != nil {
			return "", err
		}
		return "Upgrade successful, please restart", nil
	}
	return "", errors.New("Internal error")
}
