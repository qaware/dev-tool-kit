package core

import (
	"testing"
)

func TestAsciiArt(t *testing.T) {
	_, err := GenerateAsciiArt("42")
	if err == nil {
		t.Error()
	}

	CreateAsciiFont()
	text, err := GenerateAsciiArt("42")
	if err != nil || text != "   __ __ ___ \n  / // /|__ \\\n / // /___/ /\n/__  __/ __/ \n  /_/ /____/ \n             \n" {
		t.Error()
	}
}
