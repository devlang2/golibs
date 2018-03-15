package network

import (
	"encoding/binary"
	"net"
)

func IpToInt32(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func GetIpv4CIDRRange(s string) (net.IP, net.IP, error) {
	list, err := GetIpv4CIDRRangeList(s)
	if err != nil {
		return nil, nil, err
	}
	return list[0], list[len(list)-1], nil
}

func GetIpv4CIDRRangeList(s string) ([]net.IP, error) {
	ip, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}

	list := make([]net.IP, 0)
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		list = append(list, net.ParseIP(ip.String()))
	}
	return list, err
}

// https://play.golang.org/p/oJcoGMJngcE
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
