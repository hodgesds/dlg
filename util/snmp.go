package util

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseSNMPEndpoint returns SNMP host and port.
func ParseSNMPEndpoint(s string) (string, uint16, error) {
	ss := strings.Split(s, ":")
	if len(ss) != 2 {
		return "", 0, fmt.Errorf("invalid host port: %q", s)
	}
	p, err := strconv.Atoi(ss[1])
	if err != nil {
		return "", 0, err
	}
	return ss[0], uint16(p), nil
}
