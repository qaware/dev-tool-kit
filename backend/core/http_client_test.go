package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRequest(t *testing.T) {
	request, err := GenerateHttpRequest("// Comment\n  https://httpbin.org  \n /get\n\t?param1=test1  \n   &param2=test2")
	assert.Nil(t, err)
	assert.Equal(t, "GET", request.Method)
	assert.Equal(t, "https://httpbin.org/get?param1=test1&param2=test2", request.Url)

	request, err = GenerateHttpRequest("// Comment\n  GET   https://httpbin.org")
	assert.Nil(t, err)
	assert.Equal(t, "GET", request.Method)
	assert.Equal(t, "https://httpbin.org", request.Url)

	request, err = GenerateHttpRequest("\n// Comment\nPOST https://httpbin.org  \n /get\n ?param1=test1\n   &param2=test2")
	assert.Nil(t, err)
	assert.Equal(t, "POST", request.Method)
	assert.Equal(t, "https://httpbin.org/get?param1=test1&param2=test2", request.Url)

	request, err = GenerateHttpRequest(" GET https://httpbin.org\nAuthorization : Basic TEST   \n")
	assert.Nil(t, err)
	assert.Equal(t, "https://httpbin.org", request.Url)
	assert.Equal(t, "Basic TEST", request.Header["Authorization"])

	request, err = GenerateHttpRequest("GET https://httpbin.org\nAuthorization : Basic TEST:test   \nContent-Type: application/json")
	assert.Nil(t, err)
	assert.Equal(t, "https://httpbin.org", request.Url)
	assert.Equal(t, "Basic TEST:test", request.Header["Authorization"])
	assert.Equal(t, "application/json", request.Header["Content-Type"])

	request, err = GenerateHttpRequest("GET https://httpbin.org\nAuthorization : Basic username password")
	assert.Nil(t, err)
	assert.Equal(t, "https://httpbin.org", request.Url)
	assert.Equal(t, "Basic dXNlcm5hbWU6cGFzc3dvcmQ=", request.Header["Authorization"])

	request, err = GenerateHttpRequest("POST https://httpbin.org\nAuthorization: Basic Test  \n   \n{\"key\" : \n  true}")
	assert.Nil(t, err)
	assert.Equal(t, "POST", request.Method)
	assert.Equal(t, "https://httpbin.org", request.Url)
	assert.Equal(t, "Basic Test", request.Header["Authorization"])
	assert.Equal(t, "{\"key\" :\ntrue}", request.Body)

	request, err = GenerateHttpRequest("foobar")
	assert.NotNil(t, err)
}

func TestGetCurl(t *testing.T) {
	request := HttpRequest{Method: "GET", Url: "http://test"}
	assert.Equal(t, "curl \"http://test\"", request.GetCurl())

	header := make(map[string]string)

	header["Key1"] = "Value1"
	request = HttpRequest{Method: "GET", Url: "http://test", Header: header}
	assert.Equal(t, "curl \"http://test\" -H \"Key1: Value1\"", request.GetCurl())

	header["Key2"] = "Value2"
	request = HttpRequest{Method: "GET", Url: "http://test", Header: header}
	assert.Equal(t, "curl \"http://test\" -H \"Key1: Value1\" -H \"Key2: Value2\"", request.GetCurl())

	request = HttpRequest{Method: "POST", Url: "http://test", Header: header, Body: "{\"Body\"}"}
	assert.Equal(t, "curl -X POST \"http://test\" -d \"{\\\"Body\\\"}\" -H \"Key1: Value1\" -H \"Key2: Value2\"", request.GetCurl())
}
