package ui

import (
	"github.com/qaware/dev-tool-kit/backend/core"
	"strings"
	"testing"
)

func TestIntegrationHttp(t *testing.T) {
	err := handleHttpServerEvent([]string{"{\"foobar\":true}", "202", "8080"}, func(event string, content string, info string) {
		if event != "httpServer.receive" {
			t.Error()
		}
		if !strings.Contains(content, "Authorization: Basic bXlVc2VybmFtZTpteVBhc3N3b3Jk") {
			t.Error()
		}
		if !strings.Contains(info, "GET") {
			t.Error()
		}
	})
	if err == nil || !core.IsInformation(err) || !strings.Contains(err.Error(), "8080") {
		t.Error()
	}

	ports, err := handlePortsEvent([]string{"localhost", "8080", "8080"})
	if ports != "8080" || err == nil || !core.IsInformation(err) || err.Error() != "Found 1 open TCP port in range" {
		t.Error()
	}

	response, err := handleHttpClientEvent("send", "GET http://localhost:8080\nAuthorization: Basic myUsername myPassword")
	if err == nil || !core.IsInformation(err) ||
		!strings.HasPrefix(err.Error(), "202 Accepted") ||
		!strings.HasPrefix(response, "Content-Length: 15\nContent-Type: application/json; charset=UTF-8") ||
		!strings.HasSuffix(response, "{\n    \"foobar\": true\n}") {
		t.Error()
	}

	err = handleHttpServerEvent([]string{"", "", ""}, nil)
	if err != nil {
		t.Error()
	}
}

func TestInvalidInputHttp(t *testing.T) {
	err := handleHttpServerEvent([]string{"", "200", "12345678"}, nil)
	if err == nil || err.Error() != "Invalid port" {
		t.Error()
	}

	_, err = handleHttpClientEvent("send", "GET http://localhost:9090")
	if err == nil || err.Error() != "Connection refused" {
		t.Error()
	}
}
