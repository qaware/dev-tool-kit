package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiffTexts(t *testing.T) {
	diff, count := DiffTexts("Ein Test", "Kein Test")
	assert.Equal(t, 2, count)
	assert.Equal(t, "<del id=\"diff-0\" style=\"background:#ffe6e6;\">E</del><ins id=\"diff-1\" style=\"background:#e6ffe6;\">Ke</ins><span>in Test</span>", diff)

	diff, count = DiffTexts("Ein Test", "Ein Test")
	assert.Zero(t, count)
	assert.Equal(t, "<span>Ein Test</span>", diff)

	diff, count = DiffTexts("", "")
	assert.Zero(t, count)
	assert.Empty(t, diff)
}
