// Package address provide some practical functions may used.
package address

import (
	"net"
	"net/http"
	"strconv"
	"strings"
)

// ScanIPv4 scan all the available local ipv4 address.
func ScanIPv4() (map[string]net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := map[string]net.IP{}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil && ipnet.IP.IsGlobalUnicast() {
					ips[iface.Name] = ipnet.IP
				}
			}
		}
	}

	return ips, nil
}

// FindPort find the first available port from given default port to 65535
func FindPort(host string, defaultPort int) int {
	for port := defaultPort; port < 65535; port++ {
		address := "http://" + host + ":" + strconv.Itoa(port)
		_, err := http.Head(address)
		if err != nil && strings.Contains(err.Error(), "refused") {
			return port
		}
	}
	return -1
}
