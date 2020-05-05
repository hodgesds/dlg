package executor

import (
	"context"

	"github.com/hodgesds/dlg/config"
	"github.com/hodgesds/dlg/config/dhcp4"
	"github.com/hodgesds/dlg/config/dns"
	"github.com/hodgesds/dlg/config/etcd"
	"github.com/hodgesds/dlg/config/http"
	"github.com/hodgesds/dlg/config/ldap"
	"github.com/hodgesds/dlg/config/memcache"
	"github.com/hodgesds/dlg/config/redis"
	"github.com/hodgesds/dlg/config/snmp"
	"github.com/hodgesds/dlg/config/sql"
	"github.com/hodgesds/dlg/config/ssh"
	"github.com/hodgesds/dlg/config/udp"
	"github.com/hodgesds/dlg/config/websocket"
)

// Plan is used for executing a Plan.
type Plan interface {
	Execute(context.Context, *config.Plan) error
}

// Stage is used for executing a Stage.
type Stage interface {
	Execute(context.Context, *config.Stage) error
}

// DHCP4 is uesd for executing DHCP4.
type DHCP4 interface {
	Execute(context.Context, *dhcp4.Config) error
}

// DNS is used for executing a DNS.
type DNS interface {
	Execute(context.Context, *dns.Config) error
}

// ETCD is used for executing ETCD.
type ETCD interface {
	Execute(context.Context, *etcd.Config) error
}

// HTTP is used for executing HTTP.
type HTTP interface {
	Execute(context.Context, *http.Config) error
}

// LDAP is used for executing LDAP.
type LDAP interface {
	Execute(context.Context, *ldap.Config) error
}

// Memcache is used for executing Memcache.
type Memcache interface {
	Execute(context.Context, *memcache.Config) error
}

// Redis is used for executing Redis.
type Redis interface {
	Execute(context.Context, *redis.Config) error
}

// SQL is used for executing SQL.
type SQL interface {
	Execute(context.Context, *sql.Config) error
}

// SSH is used for executing SSH.
type SSH interface {
	Execute(context.Context, *ssh.Config) error
}

// SNMP is used for executing SNMP.
type SNMP interface {
	Execute(context.Context, *snmp.Config) error
}

// UDP is used for executing UDP.
type UDP interface {
	Execute(context.Context, *udp.Config) error
}

// Websocket is used for websocket.
type Websocket interface {
	Execute(context.Context, *websocket.Config) error
}
