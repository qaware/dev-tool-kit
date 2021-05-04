package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	result, err := Calculate("")
	assert.Nil(t, err)
	assert.Equal(t, "", result)

	result, err = Calculate("5 / 2")
	assert.Nil(t, err)
	assert.Equal(t, "2.5", result)

	result, err = Calculate("40 / 2")
	assert.Nil(t, err)
	assert.Equal(t, "20", result)

	result, err = Calculate("0x42")
	assert.Nil(t, err)
	assert.Equal(t, "66", result)

	result, err = Calculate("0x4242")
	assert.Nil(t, err)
	assert.Equal(t, "16962", result)

	result, err = Calculate("0b1001001")
	assert.Nil(t, err)
	assert.Equal(t, "73", result)

	result, err = Calculate("(0xa2+0b1001001-1)*0xA")
	assert.Nil(t, err)
	assert.Equal(t, "2340", result)

	result, err = Calculate("11_2 == 0b11")
	assert.Nil(t, err)
	assert.Equal(t, "1", result)

	result, err = Calculate("Ab5_16 == 0xab5")
	assert.Nil(t, err)
	assert.Equal(t, "1", result)

	result, err = Calculate("AG10K_36 == 17543972")
	assert.Nil(t, err)
	assert.Equal(t, "1", result)

	result, err = Calculate("invalid")
	assert.NotNil(t, err)
}
