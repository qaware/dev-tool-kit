package ui

import (
	"errors"
	"github.com/qaware/dev-tool-kit/backend/core"
	"strconv"
)

var sshTunnel *core.SshTunnel

func handleSshTunnel(values []string) error {
	if len(values) != 7 {
		return errors.New("Internal error")
	}

	if sshTunnel == nil {
		sshHost := values[0]
		sshPortString := values[1]
		username := values[2]
		password := values[3]
		remoteHost := values[4]
		remotePortString := values[5]
		localPortString := values[6]

		if sshHost == "" {
			return errors.New("Missing SSH host")
		}
		if sshPortString == "" {
			sshPortString = "22"
		}
		if remoteHost == "" {
			return errors.New("Missing remote host")
		}
		if remotePortString == "" {
			return errors.New("Missing remote port")
		}
		if localPortString == "" {
			return errors.New("Missing local port")
		}

		sshPort, err := strconv.Atoi(sshPortString)
		if err != nil {
			return errors.New("Invalid SSH port")
		}
		remotePort, err := strconv.Atoi(remotePortString)
		if err != nil {
			return errors.New("Invalid remote port")
		}
		localPort, err := strconv.Atoi(localPortString)
		if err != nil {
			return errors.New("Invalid local port")
		}

		sshTunnel = core.NewSshTunnel(sshHost, sshPort, username, password, remoteHost, remotePort, localPort)
		return sshTunnel.Start()
	} else {
		err := sshTunnel.Stop()
		if err != nil {
			return err
		}
		sshTunnel = nil
		return nil
	}
}
