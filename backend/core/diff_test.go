package core

import (
	"testing"
)

func TestDiffTexts(t *testing.T) {
	diff, count := DiffTexts("Ein Test", "Kein Test")
	if count != 2 || diff != "<del id=\"diff-0\" style=\"background:#ffe6e6;\">E</del><ins id=\"diff-1\" style=\"background:#e6ffe6;\">Ke</ins><span>in Test</span>" {
		t.Error()
	}

	diff, count = DiffTexts("Ein Test", "Ein Test")
	if count != 0 || diff != "<span>Ein Test</span>" {
		t.Error()
	}

	diff, count = DiffTexts("", "")
	if count != 0 || diff != "" {
		t.Error()
	}
}
