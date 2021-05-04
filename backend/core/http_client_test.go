package core

import (
	"testing"
)

func TestGenerateRequest(t *testing.T) {
	request, err := GenerateHttpRequest("// Comment\n  https://httpbin.org  \n /get\n\t?param1=test1  \n   &param2=test2")
	if err != nil || request.Method != "GET" || request.Url != "https://httpbin.org/get?param1=test1&param2=test2" {
		t.Error()
	}

	request, err = GenerateHttpRequest("// Comment\n  GET   https://httpbin.org")
	if err != nil || request.Method != "GET" || request.Url != "https://httpbin.org" {
		t.Error()
	}

	request, err = GenerateHttpRequest("\n// Comment\nPOST https://httpbin.org  \n /get\n ?param1=test1\n   &param2=test2")
	if err != nil || request.Method != "POST" || request.Url != "https://httpbin.org/get?param1=test1&param2=test2" {
		t.Error()
	}

	request, err = GenerateHttpRequest(" GET https://httpbin.org\nAuthorization : Basic TEST   \n")
	if err != nil || request.Url != "https://httpbin.org" || request.Header["Authorization"] != "Basic TEST" {
		t.Error()
	}

	request, err = GenerateHttpRequest("GET https://httpbin.org\nAuthorization : Basic TEST:test   \nContent-Type: application/json")
	if err != nil || request.Url != "https://httpbin.org" || request.Header["Authorization"] != "Basic TEST:test" || request.Header["Content-Type"] != "application/json" {
		t.Error()
	}

	request, err = GenerateHttpRequest("GET https://httpbin.org\nAuthorization : Basic username password")
	if err != nil || request.Url != "https://httpbin.org" || request.Header["Authorization"] != "Basic dXNlcm5hbWU6cGFzc3dvcmQ=" {
		t.Error()
	}

	request, err = GenerateHttpRequest("POST https://httpbin.org\nAuthorization: Basic Test  \n   \n{\"key\" : \n  true}")
	if err != nil || request.Method != "POST" || request.Url != "https://httpbin.org" || request.Header["Authorization"] != "Basic Test" || request.Body != "{\"key\" :\ntrue}" {
		t.Error()
	}

	request, err = GenerateHttpRequest("foobar")
	if err == nil {
		t.Error()
	}
}

func TestGetCurl(t *testing.T) {
	request := HttpRequest{Method: "GET", Url: "http://test"}
	if request.GetCurl() != "curl \"http://test\"" {
		t.Error()
	}

	header := make(map[string]string)

	header["Key1"] = "Value1"
	request = HttpRequest{Method: "GET", Url: "http://test", Header: header}
	if request.GetCurl() != "curl \"http://test\" -H \"Key1: Value1\"" {
		t.Error()
	}

	header["Key2"] = "Value2"
	request = HttpRequest{Method: "GET", Url: "http://test", Header: header}
	if request.GetCurl() != "curl \"http://test\" -H \"Key1: Value1\" -H \"Key2: Value2\"" {
		t.Error()
	}

	request = HttpRequest{Method: "POST", Url: "http://test", Header: header, Body: "{\"Body\"}"}
	if request.GetCurl() != "curl -X POST \"http://test\" -d \"{\\\"Body\\\"}\" -H \"Key1: Value1\" -H \"Key2: Value2\"" {
		t.Error()
	}
}
