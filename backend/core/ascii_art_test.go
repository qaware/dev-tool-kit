package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	_, err := GenerateAsciiArt("42")
	assert.NotNil(t, err)

	CreateAsciiFont()
	text, err := GenerateAsciiArt("42")
	assert.Nil(t, err)
	assert.Equal(t, "   __ __ ___ \n  / // /|__ \\\n / // /___/ /\n/__  __/ __/ \n  /_/ /____/ \n             \n", text)
}
