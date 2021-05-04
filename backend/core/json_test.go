package core

import (
	"testing"
)

func TestFormatJson(t *testing.T) {
	text, err := FormatJson("_{<_")
	if err == nil || text != "_{<_" {
		t.Error()
	}

	text, err = FormatJson("")
	if err != nil || text != "" {
		t.Error()
	}

	text, err = FormatJson("{\"key\":\"value\",\"bool\":true}")
	if err != nil || text != "{\n    \"key\": \"value\",\n    \"bool\": true\n}" {
		t.Error()
	}
}
