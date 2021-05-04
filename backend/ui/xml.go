package ui

import "github.com/qaware/dev-tool-kit/backend/core"

func handleXmlEvent(value string) string {
	return core.FormatXml(value)
}
