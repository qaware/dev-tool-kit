package core

import (
	"net"
	"strconv"
	"time"
)

func ScanPorts(hostname string, fromPort int, toPort int) []string {
	openPorts := make([]string, 0)
	for port := fromPort; port <= toPort; port++ {
		if isPortOpen(hostname, port) {
			openPorts = append(openPorts, strconv.Itoa(port))
		}
	}
	return openPorts
}

func isPortOpen(hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
