package network

import (
	"bytes"
	"net"
	"testing"
)

var ipList = []*struct {
	ip net.IP // see RFC 791 and RFC 4291
}{
	// IPv4 address
	{
		net.IP{192, 0, 2, 1},
	},
	{
		net.IP{10, 10, 10, 10},
	},
}

func TestNetworks(t *testing.T) {
	for _, c := range ipList {
		intIp := IpToInt(c.ip)

		returnedIP := IntToIp(intIp)
		if !bytes.Equal(c.ip, returnedIP) {
			t.Errorf("OriginalIP(%v) = ReturnedIP(%v)", c.ip, returnedIP)
		}
	}
}
