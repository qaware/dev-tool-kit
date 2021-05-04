package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGzipEncode(t *testing.T) {
	value, err := GzipEncode("Töst ?/=")
	assert.Nil(t, err)
	assert.Equal(t, "H4sIAAAAAAAA/wo5vK24RMFe3xYQAAD//xV+rjoJAAAA", value)

	value, err = GzipEncode("")
	assert.Nil(t, err)
	assert.Empty(t, value)
}

func TestGzipDecode(t *testing.T) {
	value, err := GzipDecode("H4sIAAAAAAAA/wo5vK24RMFe3xYQAAD//xV+rjoJAAAA")
	assert.Nil(t, err)
	assert.Equal(t, "Töst ?/=", value)

	value, err = GzipDecode("H4sIAF+6TV4AAws5vK24RMFe3xYAFX6uOgkAAAA=")
	assert.Nil(t, err)
	assert.Equal(t, "Töst ?/=", value)

	value, err = GzipDecode("H4sIAAAAAAAA/ws5vK24RMFe3xYAFX6uOgkAAAA=")
	assert.Nil(t, err)
	assert.Equal(t, "Töst ?/=", value)

	value, err = GzipDecode("")
	assert.Nil(t, err)
	assert.Empty(t, value)

	value, err = GzipDecode("-")
	assert.NotNil(t, err)
	assert.Empty(t, value)
}
