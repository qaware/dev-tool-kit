package core

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpRequest struct {
	Method         string
	Url            string
	Body           string
	TimeoutSeconds int
	Header         map[string]string
}

type HttpResponse struct {
	Code          string
	Body          []byte
	Failed        bool
	ErrorMessage  string
	ContentLength int
	ContentType   string
	Header        map[string][]string
	Duration      int64
}

var ResponseChannel = make(chan HttpResponse)

func (request HttpRequest) Perform() HttpResponse {
	go request.Send()
	return <-ResponseChannel
}

func GenerateHttpRequest(input string) (*HttpRequest, error) {
	request := &HttpRequest{Method: "GET", Header: make(map[string]string), TimeoutSeconds: 20}

	readBody := false
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "#") {
			continue
		}

		if readBody {
			request.Body += trimmed + "\n"
		} else {
			if strings.Contains(trimmed, "http") {
				if strings.HasPrefix(trimmed, "POST") {
					request.Method = "POST"
				}
				request.Url = strings.TrimSpace(strings.TrimLeft(trimmed, "GETPOST"))
			} else if trimmed == "" {
				readBody = true
			} else if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
				request.Url += trimmed
			} else {
				splitted := strings.Split(trimmed, ":")
				if len(splitted) > 1 {
					key := strings.TrimSpace(splitted[0])
					value := strings.TrimSpace(strings.Join(splitted[1:], ":"))
					request.Header[key] = value
				}
			}
		}
	}

	if request.Url == "" {
		return request, errors.New("Missing URL")
	}

	request.Body = strings.TrimSpace(request.Body)

	value, present := request.Header["Authorization"]
	if present {
		splitValue := strings.Fields(value)
		if len(splitValue) == 3 {
			request.Header["Authorization"] = splitValue[0] + " " + Base64Encode(splitValue[1]+":"+splitValue[2])
		}
	}

	return request, nil
}

func (request HttpRequest) GetCurl() string {
	curl := "curl "
	if request.Method == "POST" {
		curl += "-X POST \"" + request.Url + "\" -d \"" + strings.ReplaceAll(request.Body, "\"", "\\\"") + "\""
	} else {
		curl += "\"" + request.Url + "\""
	}

	for key, value := range request.Header {
		curl += " -H \"" + key + ": " + value + "\""
	}
	return curl
}

func (request HttpRequest) Send() {
	requestObject, err := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Body))
	if err != nil {
		ResponseChannel <- HttpResponse{Failed: true, ErrorMessage: err.Error()}
		return
	}

	for key, value := range request.Header {
		requestObject.Header.Add(key, value)
	}

	client := &http.Client{Timeout: time.Second * time.Duration(request.TimeoutSeconds)}

	start := time.Now()
	responseObject, err := client.Do(requestObject)
	duration := time.Since(start).Milliseconds()
	if err != nil {
		ResponseChannel <- HttpResponse{Failed: true, ErrorMessage: err.Error()}
		return
	}
	defer responseObject.Body.Close()

	body, err := ioutil.ReadAll(responseObject.Body)
	if err != nil {
		ResponseChannel <- HttpResponse{Failed: true, ErrorMessage: err.Error()}
		return
	}

	header := responseObject.Header
	contentType := responseObject.Header.Get("Content-Type")
	response := HttpResponse{Code: responseObject.Status, Body: body, ContentLength: len(body), ContentType: contentType, Header: header, Duration: duration}
	ResponseChannel <- response
}

func (response HttpResponse) IsOk() bool {
	return !response.Failed && strings.HasPrefix(response.Code, "200")
}
