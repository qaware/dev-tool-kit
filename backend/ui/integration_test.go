package ui

import (
	"github.com/qaware/dev-tool-kit/backend/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIntegrationHttp(t *testing.T) {
	err := handleHttpServerEvent([]string{"{\"foobar\":true}", "202", "8080"}, func(event string, content string, info string) {
		assert.Equal(t, "httpServer.receive", event)
		assert.True(t, strings.Contains(content, "Authorization: Basic bXlVc2VybmFtZTpteVBhc3N3b3Jk"))
		assert.True(t, strings.Contains(info, "GET"))
	})
	assert.NotNil(t, err)
	assert.True(t, core.IsInformation(err))
	assert.True(t, strings.Contains(err.Error(), "8080"))

	ports, err := handlePortsEvent([]string{"localhost", "8080", "8080"})
	assert.Equal(t, "8080", ports)
	assert.NotNil(t, err)
	assert.True(t, core.IsInformation(err))
	assert.Equal(t, "Found 1 open TCP port in range", err.Error())

	response, err := handleHttpClientEvent("send", "GET http://localhost:8080\nAuthorization: Basic myUsername myPassword")
	assert.NotNil(t, err)
	assert.True(t, core.IsInformation(err))
	assert.True(t, strings.HasPrefix(err.Error(), "202 Accepted"))
	assert.True(t, strings.HasPrefix(response, "Content-Length: 15\nContent-Type: application/json; charset=UTF-8"))
	assert.True(t, strings.HasSuffix(response, "{\n    \"foobar\": true\n}"))

	err = handleHttpServerEvent([]string{"", "", ""}, nil)
	assert.Nil(t, err)
}

func TestInvalidInputHttp(t *testing.T) {
	err := handleHttpServerEvent([]string{"", "200", "12345678"}, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid port", err.Error())

	_, err = handleHttpClientEvent("send", "GET http://localhost:9090")
	assert.NotNil(t, err)
	assert.Equal(t, "Connection refused", err.Error())
}
