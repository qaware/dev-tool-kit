package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	assert.Equal(t, "VMO2c3Q=", Base64Encode("Töst"))
}

func TestBase64Decode(t *testing.T) {
	value, err := Base64Decode("VMO2c3Q=")
	assert.Nil(t, err)
	assert.Equal(t, "Töst", value)

	value, err = Base64Decode("42")
	assert.NotNil(t, err)
	assert.Empty(t, value)

	value, err = Base64Decode("")
	assert.Nil(t, err)
	assert.Empty(t, value)
}

func TestIsBase64EncodedLatin1(t *testing.T) {
	assert.True(t, IsBase64EncodedLatin1("VMO2c3Q="))

	assert.False(t, IsBase64EncodedLatin1("..."))

	assert.False(t, IsBase64EncodedLatin1("anything"))
}
