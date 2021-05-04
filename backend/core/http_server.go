package core

import (
	"context"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
)

type HttpServer struct {
	Server              *http.Server
	Port                string
	ResponseStatus      int
	ResponseBody        []byte
	ResponseContentType string
}

func NewHttpServer(body string, status int, port string, receiveCallback func(http.Header, string, string)) *HttpServer {
	responseContentType := "text/plain"
	if IsJson(body) {
		responseContentType = "application/json"
	} else if IsXml(body) {
		responseContentType = "application/xml"
	}

	httpServer := &HttpServer{
		Server:              &http.Server{Addr: ":" + port},
		Port:                port,
		ResponseStatus:      status,
		ResponseBody:        []byte(body),
		ResponseContentType: responseContentType + "; charset=UTF-8",
	}

	handler := func(writer http.ResponseWriter, request *http.Request) {
		body := ""
		defer request.Body.Close()
		readBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			DebugError(err)
		} else {
			body = string(readBytes)
		}
		go receiveCallback(request.Header, request.Method, body)

		writer.Header().Set("Content-Type", httpServer.ResponseContentType)
		writer.WriteHeader(httpServer.ResponseStatus)
		_, err = writer.Write(httpServer.ResponseBody)
		if err != nil {
			DebugError(err)
		}
	}

	httpServer.Server.Handler = http.HandlerFunc(handler)

	return httpServer
}

func (httpServer HttpServer) Start() error {
	DebugInfo("Starting HTTP server at port " + httpServer.Port)

	listener, err := net.Listen("tcp", httpServer.Server.Addr)
	if err != nil {
		DebugError(err)
		return FormatHttpErrorMessage(err.Error())
	}

	httpServer.Port = strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)

	go func() {
		err := httpServer.Server.Serve(listener)
		if err != nil {
			DebugError(err)
		}
	}()

	return &Information{Info: "Listening on port " + httpServer.Port}
}

func (httpServer HttpServer) Stop() error {
	err := httpServer.Server.Shutdown(context.Background())
	if err != nil {
		DebugError(err)
		return errors.New("Error shutting down the HTTP server")
	}
	return nil
}
