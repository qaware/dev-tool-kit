package core

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func deleteLogFile(t *testing.T) {
	filename, err := GetLogFileName()
	assert.Nil(t, err)
	err = os.Remove(filename)
	assert.Nil(t, err)
}

func TestWriteAndLoadLogFile(t *testing.T) {
	message := "This is my\nlog message\nö?/#+*."
	assert.True(t, WriteLogFile(message))

	content, err := LoadLogFile()
	assert.Nil(t, err)
	assert.Equal(t, message, content)

	deleteLogFile(t)

	empty, err := LoadLogFile()
	assert.Nil(t, err)
	assert.Empty(t, empty)
	assert.True(t, WriteLogFile(message))

	deleteLogFile(t)
}
