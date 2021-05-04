package core

import (
	"errors"
	"reflect"
	"strings"
)

type Information struct {
	Info string
}

func (info *Information) Error() string {
	return info.Info
}

func IsInformation(obj error) bool {
	DebugInfo(reflect.TypeOf(obj).String())
	return strings.Contains(reflect.TypeOf(obj).String(), "Information")
}

func FormatHttpErrorMessage(message string) error {
	allParts := strings.Split(message, ": ")
	lastPart := allParts[len(allParts)-1]

	if len(lastPart) < 2 {
		return errors.New(lastPart)
	}
	return errors.New(strings.ToUpper(lastPart[0:1]) + lastPart[1:])
}
