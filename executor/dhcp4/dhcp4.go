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
	"context"
	"fmt"

	dhcp4config "github.com/hodgesds/dlg/config/dhcp4"
	"github.com/hodgesds/dlg/executor"
	"github.com/insomniacslk/dhcp/dhcpv4/nclient4"
)

type dhcp4Executor struct{}

// New returns a new Executor.
func New() executor.DHCP4 {
	return &dhcp4Executor{}
}

// Execute implements the DHCP4 interface.
func (e *dhcp4Executor) Execute(ctx context.Context, config *dhcp4config.Config) error {
	opts := []nclient4.ClientOpt{}
	if config.Retry != nil {
		opts = append(opts, nclient4.WithRetry(*config.Retry))
	}
	if config.Timeout != nil {
		opts = append(opts, nclient4.WithTimeout(*config.Timeout))
	}
	if config.HwAddr.Addr != nil {
		opts = append(opts, nclient4.WithHWAddr(config.HwAddr.Addr))
	}
	c, err := nclient4.New(config.Iface, opts...)
	if err != nil {
		return err
	}
	o, err := c.DiscoverOffer(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", o)
	return c.Close()
}
