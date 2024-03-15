package goutils

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	base    = 10
	bitSize = 32
)

// Host parser host info response
type Host struct {
	IpAddr string
	Port   int
}

// ParseIPAddrPortFromHost parser ipaddr and port form host.
// support ipv4 127.0.0.1:8001 and ipv6 [::1]:8001
func ParseIPAddrPortFromHost(host string) (*Host, error) {
	posPort := strings.LastIndex(host, ":")
	port, err := strconv.ParseUint(host[posPort+1:], base, bitSize)
	if err != nil {
		return nil, fmt.Errorf("parse uint failed. %v", err)
	}

	var ipAddr net.IP
	if strings.Contains(host, "[") {
		begin := strings.Index(host, "[")
		end := strings.LastIndex(host, "]")
		ipAddr = net.ParseIP(host[begin+1 : end])
	} else {
		ipAddr = net.ParseIP(host[:posPort])
	}

	return &Host{IpAddr: ipAddr.String(), Port: int(port)}, nil
}

// GetAllIpsForEth get all ip address for input eth
func GetAllIpsForEth(ethName string) ([]string, error) {
	var allIps []string
	itf, err := net.InterfaceByName(ethName)
	if err != nil {
		return allIps, err
	}

	addresses, err := itf.Addrs()
	if err != nil {
		return allIps, err
	}

	for _, a := range addresses {
		if ipNet, ok := a.(*net.IPNet); ok && ipNet != nil && ipNet.IP.IsGlobalUnicast() {
			allIps = append(allIps, ipNet.IP.String())
		}
	}

	return allIps, nil
}

// IsIPv6 check whether the IP is IPv6 address.
func IsIPv6(ip net.IP) bool {
	if ip != nil && strings.Contains(ip.String(), ":") {
		return true
	}
	return false
}

// IsIPv4 check ip is ipv4 address
func IsIPv4(ip net.IP) bool {
	if ip != nil && strings.Contains(ip.String(), ".") {
		return true
	}
	return false
}
