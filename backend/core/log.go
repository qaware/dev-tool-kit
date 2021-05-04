package core

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func GetLogFileName() (string, error) {
	tmpDir := os.TempDir()
	tmpFiles, err := filepath.Glob(filepath.Join(tmpDir, getTodaysPrefix()+"*"))
	if err != nil {
		DebugError(err)
		return "", err
	}

	if len(tmpFiles) > 1 {
		err = errors.New("More than one log file found")
		DebugError(err)
		return "", err
	}

	if len(tmpFiles) == 0 {
		return "", nil
	}

	DebugInfo("Log file " + tmpFiles[0] + " found")
	return tmpFiles[0], nil
}

func LoadLogFile() (string, error) {
	content, err := readLogFile()
	if err != nil {
		DebugError(err)
		return "", errors.New("Error loading log file")
	}

	return content, nil
}

func readLogFile() (string, error) {
	DebugInfo("Reading log file")

	tmpFile, err := GetLogFileName()
	if err != nil {
		return "", err
	}

	if len(tmpFile) == 0 {
		DebugInfo("No log file detected")
		return "", nil
	}

	content, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		return "", err
	}

	DebugInfo("Log file read")
	return string(content), nil
}

func WriteLogFile(text string) bool {
	DebugInfo("Writing log file")

	tmpFileName, err := GetLogFileName()
	if err != nil {
		DebugError(err)
		return false
	}

	if tmpFileName == "" {
		tmpFile, err := ioutil.TempFile("", getTodaysPrefix())
		if err != nil {
			DebugError(err)
			return false
		}

		_, err = tmpFile.Write([]byte(text))
		if err != nil {
			DebugError(err)
			return false
		}

		err = tmpFile.Close()
		if err != nil {
			DebugError(err)
			return false
		}
	} else {
		err := ioutil.WriteFile(tmpFileName, []byte(text), 0600)
		if err != nil {
			DebugError(err)
			return false
		}
	}

	DebugInfo("Log file written")
	return true
}

func PutLogMessage(message string) (string, error) {
	content, err := LoadLogFile()
	if err != nil {
		return "", err
	}
	trimmedMessage := strings.TrimSpace(message)
	if trimmedMessage != "" {
		newContent := content
		if len(content) > 0 && !strings.HasSuffix(content, "\n") {
			newContent += "\n"
		}
		newContent += "[" + time.Now().Format("15:04") + "] " + trimmedMessage + "\n"
		go WriteLogFile(newContent)
		return newContent, nil
	}
	return content, nil
}

func getTodaysPrefix() string {
	return "dev-tools-log-" + time.Now().Format("2006-01-02") + "-"
}
