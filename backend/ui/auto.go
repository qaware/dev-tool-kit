package ui

import "github.com/qaware/dev-tool-kit/backend/core"

func handleAutoEvent(value string) (string, error) {
	_, err := core.GzipDecode(value)
	if err == nil {
		return "pageGzip", nil
	}

	_, err = core.HexDecode(value)
	if err == nil {
		return "pageHex", nil
	}

	if core.IsBase64EncodedLatin1(value) {
		return "pageBase64", nil
	}

	_, err = core.ConvertToTimestamp(value)
	if err == nil {
		return "pageTime", nil
	}

	_, _, err = core.DecodeJwt(value, "")
	if err == nil || core.IsInformation(err) {
		return "pageJwt", nil
	}

	if core.IsJson(value) {
		return "pageJson", nil
	}

	if core.IsXml(value) {
		return "pageXml", nil
	}

	return "", &core.Information{"Unrecognized input"}
}
