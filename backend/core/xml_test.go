package core

import (
	"testing"
)

func TestFormatXml(t *testing.T) {
	if FormatXml("_{<_") != "_{<_" {
		t.Error()
	}
	if FormatXml("") != "" {
		t.Error()
	}
	if FormatXml("<test><a></a><b c=\"d\"/></test>") != "<test>\n    <a>\n    </a>\n    <b c=\"d\"/>\n</test>" {

		t.Error()
	}
}
