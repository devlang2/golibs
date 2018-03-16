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

func GetNetworkHostCount(cidr int) int {
	if cidr == 32 {
		return 1
	} else if cidr == 31 {
		return 2
	}
	return 2<< (31-uint(cidr))
}
