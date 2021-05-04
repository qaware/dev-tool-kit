package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
	"net/http"
	"strconv"
	"time"
)

var httpServer *core.HttpServer

func handleHttpServerEvent(values []string, callback func(string, string, string)) error {
	if len(values) != 3 {
		return errors.New("Internal error")
	}

	if httpServer == nil {
		body := values[0]
		status := values[1]
		port := values[2]

		if status == "" {
			return errors.New("Missing HTTP status code")
		}
		if port == "" {
			return errors.New("Missing port")
		}

		intStatus, err := strconv.Atoi(status)
		if err != nil || intStatus < 100 || intStatus > 599 {
			return errors.New("Invalid HTTP status code")
		}

		httpServer = core.NewHttpServer(body, intStatus, port, func(header http.Header, method string, body string) {
			callback("httpServer.receive", formatContent(header, body), formatInfo(method))
		})
		return httpServer.Start()
	} else {
		err := httpServer.Stop()
		if err != nil {
			return err
		}
		httpServer = nil
		return nil
	}
}

func formatContent(header http.Header, body string) string {
	contentType := header.Get("Content-Type")
	return getPrettyPrintedHeader(header) + getPrettyPrintedBody(body, contentType)
}

func formatInfo(method string) string {
	return "Received " + method + " request at " + time.Now().Format("15:04:05")
}
