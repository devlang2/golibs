package network

import (
	"testing"
	"net"
	"errors"
)
func TestGetIpv4CIDRRangeList(t *testing.T) {
	s := "11.0.11.0/24"

	_, _, err := net.ParseCIDR(s)
	if err != nil {
		t.Error(err.Error())
	}

	list, err := GetIpv4CIDRRangeList(s)
	if len(list) != 256 {
		t.Error(errors.New("list error"))
	}

	start, end, err := GetIpv4CIDRRange(s)
	if start.String() != "11.0.11.0" {
		t.Error(errors.New("invalid first IP"))
	}

	if end.String() != "11.0.11.255" {
		t.Error(errors.New("invalid last IP"))
	}
}