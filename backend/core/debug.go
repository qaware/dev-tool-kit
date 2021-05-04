package core

import (
	"github.com/wailsapp/wails"
)

var logger *wails.CustomLogger

func InitLogger(customLogger *wails.CustomLogger) {
	logger = customLogger
}

func DebugInfo(message string) {
	if logger != nil {
		logger.Info(message)
	}
}

func DebugError(err error) {
	if logger != nil {
		logger.Error(err.Error())
	}
}
