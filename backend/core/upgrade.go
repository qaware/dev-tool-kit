package core

import (
	"encoding/json"
	"errors"
	"github.com/inconshreveable/go-update"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var version string

type GitHub struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadUrl string `json:"browser_download_url"`
	} `json:"assets"`
}

func InitUpgrade(currentVersion string) {
	version = currentVersion
}

func GetVersion() string {
	return version
}

func HasNewVersion() bool {
	DebugInfo("Checking version on server")

	remoteVersion, err := getLatestVersion()
	if err != nil {
		DebugError(err)
		return false
	}

	DebugInfo("Local version " + version)
	DebugInfo("Remote version " + remoteVersion)

	return remoteVersion != version
}

func UpgradeNow() error {
	DebugInfo("Upgrading to latest version")

	remoteVersion, err := getLatestVersion()
	if err != nil {
		DebugError(err)
		return err
	}

	exe, err := os.Executable()
	if err != nil {
		DebugError(err)
		return errors.New("Error determining the executable")
	}

	exeName := filepath.Base(exe)
	DebugInfo("Downloading executable " + exeName)

	url := "https://github.com/qaware/dev-tool-kit/releases/download/v" + remoteVersion + "/" + exeName
	DebugInfo("Downloading " + url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		DebugError(err)
		return errors.New("Error creating the request")
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

func getLatestVersion() (string, error) {
	request := &HttpRequest{Method: "GET", Url: "https://api.github.com/repos/qaware/dev-tool-kit/releases/latest"}
	response := request.Perform()

	if response.Failed {
		return "", errors.New(response.ErrorMessage)
	}
	if !response.IsOk() {
		return "", errors.New("Response code from GitHub:" + response.Code)
	}

	var githubResponse GitHub
	err := json.Unmarshal(response.Body, &githubResponse)
	if err != nil {
		return "", err
	}

	return strings.TrimLeft(githubResponse.TagName, "v"), nil
}
