package core

import (
	"testing"
)

func TestConvertToTimeString(t *testing.T) {
	text, err := ConvertToTimeString("1578140592")
	if err != nil || text != "2020-01-04T13:23:12+01:00" {
		t.Error()
	}

	text, err = ConvertToTimeString("1578140592000")
	if !IsInformation(err) || text != "2020-01-04T13:23:12+01:00" {
		t.Error()
	}

	text, err = ConvertToTimeString("1578140592123")
	if !IsInformation(err) || text != "2020-01-04T13:23:12+01:00" {
		t.Error()
	}
}

func TestConvertToTimestamp(t *testing.T) {
	timestamp, err := ConvertToTimestamp("2020-01-04T12:23:00Z")
	if err != nil || timestamp != "1578140580" {
		t.Error()
	}

	timestamp, err = ConvertToTimestamp("2020-01-04T13:23:00+01:00")
	if err != nil || timestamp != "1578140580" {
		t.Error()
	}

	timestamp, err = ConvertToTimestamp("")
	if err != nil || timestamp != "" {
		t.Error()
	}
}

func TestTimeZone(t *testing.T) {
	text, _ := ConvertToTimeString("1603407373")
	timestamp, err := ConvertToTimestamp(text)
	if err != nil || timestamp != "1603407373" {
		t.Error()
	}
}
