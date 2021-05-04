package core

import (
	"errors"
	"github.com/inconshreveable/go-update"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var version string

func InitUpgrade(currentVersion string) {
	version = currentVersion
}

func GetVersion() string {
	return version
}

func HasNewVersion() bool {
	DebugInfo("Checking version on server")

	request := &HttpRequest{Method: "GET", Url: "https://tobias-melson.de/dev-tool-kit/version", Header: getHeaderForUpdate()}
	response := request.Perform()
	remoteVersion := strings.TrimSpace(string(response.Body))

	DebugInfo("Local version " + version)
	DebugInfo("Remote version " + remoteVersion)

	return response.IsOk() && remoteVersion != version
}

func UpgradeNow() error {
	DebugInfo("Upgrading to new version")

	exe, err := os.Executable()
	if err != nil {
		DebugError(err)
		return errors.New("Error determining the executable")
	}

	exeName := filepath.Base(exe)
	DebugInfo("Downloading executable " + exeName)

	request, err := http.NewRequest("GET", "https://tobias-melson.de/dev-tool-kit/"+exeName, nil)
	if err != nil {
		DebugError(err)
		return errors.New("Error creating the request")
	}

	header := getHeaderForUpdate()

	for key, value := range header {
		request.Header.Add(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		DebugError(err)
		return errors.New("Error downloading the executable")
	}
	if response.StatusCode != 200 {
		return errors.New("Error downloading the executable")
	}

	defer response.Body.Close()

	err = update.Apply(response.Body, update.Options{})
	if err != nil {
		DebugError(err)
		return errors.New("Error upgrading the executable")
	}

	DebugInfo("Upgrade successful")
	return nil
}

func getHeaderForUpdate() map[string]string {
	header := make(map[string]string)
	header["Authorization"] = "Basic ZGV2LXRvb2xzOml0aDZ4ZWVS"
	return header
}
