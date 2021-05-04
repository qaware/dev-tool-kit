package core

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInformation(t *testing.T) {
	info := &Information{"Foobar"}
	assert.Equal(t, "Foobar", info.Error())
	assert.True(t, IsInformation(info))
	assert.False(t, IsInformation(errors.New("Other error")))
}

func TestFormatHttpErrorMessage(t *testing.T) {
	assert.Equal(t, "Test", FormatHttpErrorMessage("Test").Error())
	assert.Equal(t, "Test", FormatHttpErrorMessage("Foobar: Test").Error())
	assert.Equal(t, "Test", FormatHttpErrorMessage("foobar: test").Error())
	assert.Equal(t, "t", FormatHttpErrorMessage("foobar: t").Error())
	assert.Equal(t, "Test", FormatHttpErrorMessage("foo: bar: test").Error())
}
