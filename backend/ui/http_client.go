package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
	"sort"
	"strconv"
	"strings"
)

func handleHttpClientEvent(action string, value string) (string, error) {
	switch action {
	case "send":
		request, err := core.GenerateHttpRequest(value)
		if err != nil {
			return "", err
		}

		response := request.Perform()

		if response.Failed {
			return "", core.FormatHttpErrorMessage(response.ErrorMessage)
		}

		body := getPrettyPrintedBody(string(response.Body), response.ContentType)
		header := getPrettyPrintedHeader(response.Header)
		duration := strconv.FormatInt(response.Duration, 10) + "ms"

		return header + body, &core.Information{response.Code + " \u279C " + duration}

	case "curl":
		request, err := core.GenerateHttpRequest(value)
		if err != nil {
			return "", err
		}
		return request.GetCurl(), nil

	case "example":
		return generateHttpExample(), nil
	}
	return "", errors.New("Internal error")
}

func generateHttpExample() string {
	return "POST https://httpbin.org/post?my-parameter=true\nAuthorization: Basic myUsername myPassword\nContent-Type: application/json\n\n{\"my-json\": [true, false]}"
}

func getPrettyPrintedBody(body string, contentType string) string {
	if strings.Contains(strings.ToLower(contentType), "json") {
		json, err := core.FormatJson(body)
		if err == nil {
			return json
		}
	} else if strings.Contains(strings.ToLower(contentType), "xml") {
		return core.FormatXml(body)
	}
	if strings.TrimSpace(body) == "" {
		return "<Empty body>"
	}
	return body
}

func getPrettyPrintedHeader(header map[string][]string) string {
	headerLines := make([]string, 0, len(header))
	widthHeaderLines := 0
	for key, value := range header {
		line := key + ": " + strings.Join(value, "; ")
		headerLines = append(headerLines, line)
		if len(line) > widthHeaderLines {
			widthHeaderLines = len(line)
		}
	}

	sort.Strings(headerLines)

	result := ""
	if len(headerLines) > 0 {
		bar := strings.Repeat("=", widthHeaderLines)
		result = strings.Join(headerLines, "\n") + "\n" + bar + "\n"
	}
	return result
}
