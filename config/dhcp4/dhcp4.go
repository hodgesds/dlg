package dhcp4

import (
	"net"
	"time"
	//_ "github.com/insomniacslk/dhcp/dhcpv4/nclient4"
)

// Config is used for configuring DHCP4 client.
type Config struct {
	Iface   string
	HwAddr  net.HardwareAddr
	Retry   *int
	Timeout *time.Duration
}
