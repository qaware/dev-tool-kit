package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHexEncode(t *testing.T) {
	assert.Equal(t, "54 C3 B6 73 74 20 3F 2F 3D 0A 47 0A", HexEncode("Töst ?/=\nG\n"))
}

func TestHexDecode(t *testing.T) {
	value, err := HexDecode("54 C3 B6 73 74 20 3F 2F 3D 0A 47 0A")
	assert.Nil(t, err)
	assert.Equal(t, "Töst ?/=\nG\n", value)

	value, err = HexDecode("54C3B67374203F2F3D0A470A")
	assert.Nil(t, err)
	assert.Equal(t, "Töst ?/=\nG\n", value)

	value, err = HexDecode("   54C3B 673 \n 74203F2F3\t D0A470A  ")
	assert.Nil(t, err)
	assert.Equal(t, "Töst ?/=\nG\n", value)

	value, err = HexDecode("1")
	assert.NotNil(t, err)
	assert.Empty(t, value)

	value, err = HexDecode("")
	assert.Nil(t, err)
	assert.Empty(t, value)
}
