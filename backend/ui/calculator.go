package ui

import (
	"github.com/qaware/dev-tool-kit/backend/core"
)

func handleCalculatorEvent(action string, value string) (string, error) {
	if action == "calculate" {
		result, err := core.Calculate(value)
		if err != nil {
			return "", err
		}

		return value + " = " + result, nil
	} else {
		return "(0xFF / 255) * 10 + 0b11 - 3^2", nil
	}
}
