package core

import (
	"testing"
)

func TestHexEncode(t *testing.T) {
	if HexEncode("Töst ?/=\nG\n") != "54 C3 B6 73 74 20 3F 2F 3D 0A 47 0A" {
		t.Error()
	}
}

func TestHexDecode(t *testing.T) {
	value, err := HexDecode("54 C3 B6 73 74 20 3F 2F 3D 0A 47 0A")
	if err != nil || value != "Töst ?/=\nG\n" {
		t.Error()
	}

	value, err = HexDecode("54C3B67374203F2F3D0A470A")
	if err != nil || value != "Töst ?/=\nG\n" {
		t.Error()
	}

	value, err = HexDecode("   54C3B 673 \n 74203F2F3\t D0A470A  ")
	if err != nil || value != "Töst ?/=\nG\n" {
		t.Error()
	}

	value, err = HexDecode("1")
	if err == nil || value != "" {
		t.Error()
	}

	value, err = HexDecode("")
	if err != nil || value != "" {
		t.Error()
	}
}
