// Copyright © 2020 Daniel Hodges <hodges.daniel.scott@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dhcp4

import (
	"net"
	"time"
)

// Config is used for configuring DHCP4 client.
type Config struct {
	Iface   string         `yaml:"iface"`
	HwAddr  HwAddr         `yaml:"hwAddr,inline"`
	Retry   *int           `yaml:"retry,omitempty"`
	Timeout *time.Duration `yaml:"timeout,omitempty"`
}

// HwAddr is a hardware address.
type HwAddr struct {
	Addr net.HardwareAddr
}

// UnmarshalYAML implements the yaml unmarshal interface.
func (a *HwAddr) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return err
	}
	newA, err := net.ParseMAC(buf)
	if err != nil {
		return err
	}
	a.Addr = newA
	return nil
}

// MarshalYAML implements the yaml marshal interface.
func (a *HwAddr) MarshalYAML() (interface{}, error) {
	return a.Addr.String(), nil
}
