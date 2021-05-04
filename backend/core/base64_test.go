package core

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
	if Base64Encode("Töst") != "VMO2c3Q=" {
		t.Error()
	}
}

func TestBase64Decode(t *testing.T) {
	value, err := Base64Decode("VMO2c3Q=")
	if err != nil || value != "Töst" {
		t.Error()
	}

	value, err = Base64Decode("42")
	if err == nil || value != "" {
		t.Error()
	}

	value, err = Base64Decode("")
	if err != nil || value != "" {
		t.Error()
	}
}

func TestIsBase64EncodedLatin1(t *testing.T) {
	if !IsBase64EncodedLatin1("VMO2c3Q=") {
		t.Error()
	}

	if IsBase64EncodedLatin1("...") {
		t.Error()
	}

	if IsBase64EncodedLatin1("anything") {
		t.Error()
	}
}
