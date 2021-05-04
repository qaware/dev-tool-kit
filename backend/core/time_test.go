package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToTimeString(t *testing.T) {
	text, err := ConvertToTimeString("1578140592")
	assert.Nil(t, err)
	assert.Equal(t, "2020-01-04T13:23:12+01:00", text)

	text, err = ConvertToTimeString("1578140592000")
	assert.True(t, IsInformation(err))
	assert.Equal(t, "2020-01-04T13:23:12+01:00", text)

	text, err = ConvertToTimeString("1578140592123")
	assert.True(t, IsInformation(err))
	assert.Equal(t, "2020-01-04T13:23:12+01:00", text)
}

func TestConvertToTimestamp(t *testing.T) {
	timestamp, err := ConvertToTimestamp("2020-01-04T12:23:00Z")
	assert.Nil(t, err)
	assert.Equal(t, "1578140580", timestamp)

	timestamp, err = ConvertToTimestamp("2020-01-04T13:23:00+01:00")
	assert.Nil(t, err)
	assert.Equal(t, "1578140580", timestamp)

	timestamp, err = ConvertToTimestamp("")
	assert.Nil(t, err)
	assert.Empty(t, timestamp)
}

func TestTimeZone(t *testing.T) {
	text, _ := ConvertToTimeString("1603407373")
	timestamp, err := ConvertToTimestamp(text)
	assert.Nil(t, err)
	assert.Equal(t, "1603407373", timestamp)
}
