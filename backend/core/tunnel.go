// https://gist.github.com/svett/sshtunnel.go

package core

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
)

type Endpoint struct {
	Host string
	Port int
}

func (endpoint *Endpoint) string() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

type SshTunnel struct {
	Local            *Endpoint
	Server           *Endpoint
	Remote           *Endpoint
	Config           *ssh.ClientConfig
	Listener         net.Listener
	LocalConnection  net.Conn
	RemoteConnection net.Conn
}

func (tunnel *SshTunnel) Start() error {
	var err error
	tunnel.Listener, err = net.Listen("tcp", tunnel.Local.string())
	if err != nil {
		DebugError(err)
		return errors.New("Error creating local listener")
	}

	go tunnel.listen()
	return nil
}

func (tunnel *SshTunnel) Stop() error {
	if tunnel.RemoteConnection != nil {
		_ = tunnel.RemoteConnection.Close()
	}
	if tunnel.LocalConnection != nil {
		_ = tunnel.LocalConnection.Close()
	}
	return tunnel.Listener.Close()
}

func (tunnel *SshTunnel) listen() {
	defer tunnel.Listener.Close()

	for {
		var err error
		tunnel.LocalConnection, err = tunnel.Listener.Accept()
		if err != nil {
			DebugError(err)
			return
		}
		go tunnel.forward()
	}
}

func (tunnel *SshTunnel) forward() {
	serverConn, err := ssh.Dial("tcp", tunnel.Server.string(), tunnel.Config)
	if err != nil {
		DebugError(err)
		return
	}

	tunnel.RemoteConnection, err = serverConn.Dial("tcp", tunnel.Remote.string())
	if err != nil {
		DebugError(err)
		return
	}

	copyConn := func(writer, reader net.Conn) {
		defer writer.Close()
		defer reader.Close()

		_, err := io.Copy(writer, reader)
		if err != nil {
			DebugError(err)
		}
	}

	go copyConn(tunnel.LocalConnection, tunnel.RemoteConnection)
	go copyConn(tunnel.RemoteConnection, tunnel.LocalConnection)
}

func NewSshTunnel(sshHost string, sshPort int, username string, password string, remoteHost string, remotePort int, localPort int) *SshTunnel {
	localEndpoint := &Endpoint{
		Host: "localhost",
		Port: localPort,
	}

	serverEndpoint := &Endpoint{
		Host: sshHost,
		Port: sshPort,
	}

	remoteEndpoint := &Endpoint{
		Host: remoteHost,
		Port: remotePort,
	}

	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return &SshTunnel{
		Config: sshConfig,
		Local:  localEndpoint,
		Server: serverEndpoint,
		Remote: remoteEndpoint,
	}
}
