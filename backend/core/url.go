package core

import (
	"errors"
	"net/url"
	"strings"
)

func UrlEncode(text string) string {
	return strings.ReplaceAll(url.QueryEscape(text), "+", "%20")
}

func UrlDecode(text string) (string, error) {
	result, err := url.QueryUnescape(text)
	if err != nil {
		DebugError(err)
		return "", errors.New("Invalid URL encoding")
	}
	return result, nil
}
