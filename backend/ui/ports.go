package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
	"strconv"
	"strings"
)

func handlePortsEvent(values []string) (string, error) {
	hostname := values[0]
	if hostname == "" {
		return "", errors.New("Empty remote host")
	}

	portFrom, err := strconv.Atoi(values[1])
	if err != nil {
		return "", errors.New("Invalid port")
	}

	portTo, err := strconv.Atoi(values[2])
	if err != nil {
		return "", errors.New("Invalid port")
	}

	ports := core.ScanPorts(hostname, portFrom, portTo)
	output := strings.Join(ports, "\n")

	if len(ports) == 0 {
		return output, &core.Information{"No open TCP ports found in range"}
	}

	plural := ""
	if len(ports) > 1 {
		plural = "s"
	}
	return output, &core.Information{"Found " + strconv.Itoa(len(ports)) + " open TCP port" + plural + " in range"}
}
