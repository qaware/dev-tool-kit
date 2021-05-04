package core

import (
	"os"
	"testing"
)

func deleteLogFile(t *testing.T) {
	filename, err := GetLogFileName()
	if err != nil {
		t.Error(err)
	}
	err = os.Remove(filename)
	if err != nil {
		t.Error(err)
	}
}

func TestWriteAndLoadLogFile(t *testing.T) {
	message := "This is my\nlog message\nö?/#+*."
	if !WriteLogFile(message) {
		t.Error()
	}

	content, err := LoadLogFile()
	if err != nil || content != message {
		t.Error()
	}

	deleteLogFile(t)

	empty, err := LoadLogFile()
	if err != nil || empty != "" {
		t.Error()
	}

	if !WriteLogFile(message) {
		t.Error()
	}

	deleteLogFile(t)
}
