package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
	"strconv"
)

var numberOfDiffs = 0
var currentDiff = -1

func handleDiffEvent(action string, values []string) (string, error) {
	if action == "compare" {
		var diff string
		diff, numberOfDiffs = core.DiffTexts(values[0], values[1])
		return diff, getDiffInfo()
	} else if action == "scroll" {
		if numberOfDiffs == 0 {
			return "", nil
		}

		if values[0] == "0" {
			numberOfDiffs = 0
			currentDiff = -1
			return "", nil
		}

		currentDiff++
		if currentDiff >= numberOfDiffs {
			currentDiff = 0
		}
		return "diff-" + strconv.Itoa(currentDiff), getDiffInfo()
	}
	return "", errors.New("Internal error")
}

func getDiffInfo() error {
	plural := "s"
	if numberOfDiffs == 1 {
		plural = ""
	}
	return &core.Information{"Found " + strconv.Itoa(numberOfDiffs) + " difference" + plural}
}
