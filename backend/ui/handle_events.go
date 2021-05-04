package ui

import (
	"encoding/json"
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
	"github.com/wailsapp/wails"
	"strings"
)

type Bus struct {
	log     *wails.CustomLogger
	runtime *wails.Runtime
}

type ToBackendEvent struct {
	Source string   `json:"source"`
	Action string   `json:"action"`
	Values []string `json:"values"`
}

type ToFrontendEvent struct {
	Value string `json:"value"`
	Info  string `json:"info"`
}

func (bus *Bus) WailsInit(runtime *wails.Runtime) error {
	bus.runtime = runtime
	bus.log = runtime.Log.New("Backend")
	core.InitLogger(bus.log)
	return nil
}

func (bus *Bus) SendEvent(eventJson string) (string, error) {
	event := ToBackendEvent{}
	err := json.Unmarshal([]byte(eventJson), &event)
	if err != nil {
		bus.log.Error(err.Error())
		return bus.toFrontendEvent("", ""), errors.New("Internal error")
	}

	action := strings.TrimSpace(event.Action)
	singleValue := ""
	if len(event.Values) > 0 {
		singleValue = event.Values[0]
	}

	output := ""
	err = nil

	switch event.Source {
	case "customize":
		output, err = handleCustomizeEvent()
	case "upgrade":
		output, err = handleUpgradeEvent(action)
	case "log":
		output, err = handleLogEvent(action, singleValue)
	case "auto":
		output, err = handleAutoEvent(singleValue)
	case "ascii":
		output, err = handleAsciiEvent(singleValue)
	case "auth":
		output = handleAuthEvent(event.Values)
	case "calculator":
		output, err = handleCalculatorEvent(action, singleValue)
	case "base64":
		output, err = handleBase64Event(action, singleValue)
	case "diff":
		output, err = handleDiffEvent(action, event.Values)
	case "gzip":
		output, err = handleGzipEvent(action, singleValue)
	case "hex":
		output, err = handleHexEvent(action, singleValue)
	case "httpClient":
		output, err = handleHttpClientEvent(action, singleValue)
	case "httpServer":
		err = handleHttpServerEvent(event.Values, bus.emitToFrontendEvent)
	case "json":
		output, err = handleJsonEvent(singleValue)
	case "jwt":
		output, err = handleJwtEvent(event.Values)
	case "ports":
		output, err = handlePortsEvent(event.Values)
	case "tunnel":
		err = handleSshTunnel(event.Values)
	case "time":
		output, err = handleTimeEvent(action, singleValue)
	case "url":
		output, err = handleUrlEvent(action, singleValue)
	case "uuid":
		output, err = handleUuidEvent()
	case "xml":
		output = handleXmlEvent(singleValue)
	}

	if err == nil {
		return bus.toFrontendEvent(output, ""), nil
	} else if core.IsInformation(err) {
		return bus.toFrontendEvent(output, err.Error()), nil
	}
	return bus.toFrontendEvent(output, ""), err
}

func (bus *Bus) toFrontendEvent(value string, info string) string {
	raw, err := json.Marshal(ToFrontendEvent{
		Value: value,
		Info:  info,
	})

	if err != nil {
		bus.log.Error(err.Error())
		return ""
	}

	return string(raw)
}

func (bus *Bus) emitToFrontendEvent(event string, value string, info string) {
	bus.runtime.Events.Emit(event, bus.toFrontendEvent(value, info))
}
