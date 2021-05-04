package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatXml(t *testing.T) {
	assert.Equal(t, "_{<_", FormatXml("_{<_"))
	assert.Empty(t, FormatXml(""))
	assert.Equal(t, "<test>\n    <a>\n    </a>\n    <b c=\"d\"/>\n</test>", FormatXml("<test><a></a><b c=\"d\"/></test>"))
}
