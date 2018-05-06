// Package address provide some practical functions may used.
package address

import (
	"log"
	"net"
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

// FindPorts find available port in [defaultPort, 65535)
func FindPorts(host string, defaultPort int, reciever func(int, bool) bool) {
	for port := defaultPort; port < 65535; port++ {
		address := host + ":" + strconv.Itoa(port)
		tcpAddr, err := net.ResolveTCPAddr("tcp", address)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = net.DialTCP("tcp", nil, tcpAddr)
		available := err != nil && strings.Contains(err.Error(), "refused")
		if reciever(port, available) {
			return
		}
	}
	reciever(-1, false)
}

// FindPort find the first available port from given default port to 65535
func FindPort(host string, defaultPort int) int {
	var port int
	FindPorts(host, defaultPort, func(p int, ok bool) bool {
		if !ok {
			return false
		}
		port = p
		return true
	})
	return port
}
