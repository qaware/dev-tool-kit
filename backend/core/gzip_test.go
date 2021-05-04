package core

import (
	"testing"
)

func TestGzipEncode(t *testing.T) {
	value, err := GzipEncode("Töst ?/=")
	if err != nil || value != "H4sIAAAAAAAA/wo5vK24RMFe3xYQAAD//xV+rjoJAAAA" {
		t.Error()
	}

	value, err = GzipEncode("")
	if err != nil || value != "" {
		t.Error()
	}
}

func TestGzipDecode(t *testing.T) {
	value, err := GzipDecode("H4sIAAAAAAAA/wo5vK24RMFe3xYQAAD//xV+rjoJAAAA")
	if err != nil || value != "Töst ?/=" {
		t.Error()
	}

	value, err = GzipDecode("H4sIAF+6TV4AAws5vK24RMFe3xYAFX6uOgkAAAA=")
	if err != nil || value != "Töst ?/=" {
		t.Error()
	}

	value, err = GzipDecode("H4sIAAAAAAAA/ws5vK24RMFe3xYAFX6uOgkAAAA=")
	if err != nil || value != "Töst ?/=" {
		t.Error()
	}

	value, err = GzipDecode("")
	if err != nil || value != "" {
		t.Error()
	}

	value, err = GzipDecode("-")
	if err == nil || value != "" {
		t.Error()
	}
}
