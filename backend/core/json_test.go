package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatJson(t *testing.T) {
	text, err := FormatJson("_{<_")
	assert.NotNil(t, err)
	assert.Equal(t, "_{<_", text)

	text, err = FormatJson("")
	assert.Nil(t, err)
	assert.Empty(t, text)

	text, err = FormatJson("{\"key\":\"value\",\"bool\":true}")
	assert.Nil(t, err)
	assert.Equal(t, "{\n    \"key\": \"value\",\n    \"bool\": true\n}", text)
}
