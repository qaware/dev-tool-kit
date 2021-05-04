package core

import (
	"github.com/go-xmlfmt/xmlfmt"
	"strings"
)

func FormatXml(raw string) string {
	xml := xmlfmt.FormatXML(raw, "", "    ")
	return strings.TrimSpace(strings.ReplaceAll(xml, xmlfmt.NL, "\n"))
}

func IsXml(input string) bool {
	return strings.HasPrefix(input, "<") && strings.HasSuffix(input, ">")
}
