package core

import (
	"errors"
	"testing"
)

func TestIsInformation(t *testing.T) {
	info := &Information{"Foobar"}
	if info.Error() != "Foobar" || !IsInformation(info) {
		t.Error()
	}
	if IsInformation(errors.New("Other error")) {
		t.Error()
	}
}

func TestFormatHttpErrorMessage(t *testing.T) {
	if FormatHttpErrorMessage("Test").Error() != "Test" {
		t.Error()
	}
	if FormatHttpErrorMessage("Foobar: Test").Error() != "Test" {
		t.Error()
	}
	if FormatHttpErrorMessage("foobar: test").Error() != "Test" {
		t.Error()
	}
	if FormatHttpErrorMessage("foobar: t").Error() != "t" {
		t.Error()
	}
	if FormatHttpErrorMessage("foo: bar: test").Error() != "Test" {
		t.Error()
	}
}
