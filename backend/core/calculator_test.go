package core

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	result, err := Calculate("")
	if err != nil || result != "" {
		t.Error()
	}

	result, err = Calculate("5 / 2")
	if err != nil || result != "2.5" {
		t.Error()
	}

	result, err = Calculate("40 / 2")
	if err != nil || result != "20" {
		t.Error()
	}

	result, err = Calculate("0x42")
	if err != nil || result != "66" {
		t.Error()
	}

	result, err = Calculate("0x4242")
	if err != nil || result != "16962" {
		t.Error()
	}

	result, err = Calculate("0b1001001")
	if err != nil || result != "73" {
		t.Error()
	}

	result, err = Calculate("(0xa2+0b1001001-1)*0xA")
	if err != nil || result != "2340" {
		t.Error()
	}

	result, err = Calculate("11_2 == 0b11")
	if err != nil || result != "1" {
		t.Error()
	}

	result, err = Calculate("Ab5_16 == 0xab5")
	if err != nil || result != "1" {
		t.Error()
	}

	result, err = Calculate("AG10K_36 == 17543972")
	if err != nil || result != "1" {
		t.Error()
	}

	result, err = Calculate("invalid")
	if err == nil {
		t.Error()
	}
}
