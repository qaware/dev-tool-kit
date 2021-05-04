package core

import (
	"errors"
	"math"
	"strconv"
	"time"
)

func ConvertToTimestamp(timeString string) (string, error) {
	if timeString == "" {
		return "", nil
	}

	dateTime, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		DebugError(err)
		return "", errors.New("Invalid ISO 8601 format")
	}

	return strconv.FormatInt(dateTime.Unix(), 10), nil
}

func ConvertToTimeString(timestamp string) (string, error) {
	if timestamp == "" {
		return "", nil
	}

	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		DebugError(err)
		return "", errors.New("Invalid timestamp")
	}

	isMilliseconds := false
	if i >= 1000000000000 {
		i = int64(math.Round(float64(i) / 1000.))
		isMilliseconds = true
	}

	result := time.Unix(i, 0).In(time.Local).Format(time.RFC3339)

	if isMilliseconds {
		return result, &Information{"Timestamp is assumed to be in milliseconds"}
	}
	return result, nil
}

func GetNow() string {
	return time.Now().Format(time.RFC3339)
}
