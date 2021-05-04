package core

import (
	"testing"
)

func TestUrlEncode(t *testing.T) {
	if UrlEncode("Töst /:?+") != "T%C3%B6st%20%2F%3A%3F%2B" {
		t.Error()
	}
}

func TestUrlDecode(t *testing.T) {
	value, err := UrlDecode("T%C3%B6st%20%2F%3A%3F%2B")
	if err != nil || value != "Töst /:?+" {
		t.Error()
	}

	value, err = UrlDecode("T%C3%B6st+%2F%3A%3F%2B")
	if err != nil || value != "Töst /:?+" {
		t.Error()
	}

	value, err = UrlDecode("%")
	if err == nil || value != "" {
		t.Error()
	}

	value, err = UrlDecode("")
	if err != nil || value != "" {
		t.Error()
	}
}
