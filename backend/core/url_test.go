package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	assert.Equal(t, "T%C3%B6st%20%2F%3A%3F%2B", UrlEncode("Töst /:?+"))
}

func TestUrlDecode(t *testing.T) {
	value, err := UrlDecode("T%C3%B6st%20%2F%3A%3F%2B")
	assert.Nil(t, err)
	assert.Equal(t, "Töst /:?+", value)

	value, err = UrlDecode("T%C3%B6st+%2F%3A%3F%2B")
	assert.Nil(t, err)
	assert.Equal(t, "Töst /:?+", value)

	value, err = UrlDecode("%")
	assert.NotNil(t, err)
	assert.Empty(t, value)

	value, err = UrlDecode("")
	assert.Nil(t, err)
	assert.Empty(t, value)
}
