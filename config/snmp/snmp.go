package snmp

import (
	"github.com/gosnmp/gosnmp"
)

// Config is used to config SNMP.
type Config struct {
	Addr      string
	Port      uint16
	Transport string
	Community string
	Retries   int
	Version   string
	Oids      []string
	Walk      string
}

// SNMP is a SNMP configuration.
func (c *Config) SNMP() *gosnmp.GoSNMP {
	g := &gosnmp.GoSNMP{
		Target:    c.Addr,
		Port:      c.Port,
		Transport: c.Transport,
		Community: c.Community,
		Retries:   c.Retries,
	}
	switch c.Version {
	case "1":
		g.Version = gosnmp.Version1
	case "2c":
		g.Version = gosnmp.Version2c
	case "3":
		g.Version = gosnmp.Version3
	}
	return g
}
