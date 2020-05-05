package config

import (
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/hodgesds/dlg/config/dhcp4"
	"github.com/hodgesds/dlg/config/dns"
	"github.com/hodgesds/dlg/config/etcd"
	"github.com/hodgesds/dlg/config/http"
	"github.com/hodgesds/dlg/config/kafka"
	"github.com/hodgesds/dlg/config/ldap"
	"github.com/hodgesds/dlg/config/memcache"
	"github.com/hodgesds/dlg/config/redis"
	"github.com/hodgesds/dlg/config/sql"
	"github.com/hodgesds/dlg/config/ssh"
	"github.com/hodgesds/dlg/config/udp"
	"github.com/hodgesds/dlg/config/websocket"
	"github.com/hodgesds/dlg/util"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestConfig(t *testing.T) {
	macAddr, err := net.ParseMAC("00:00:5e:00:53:01")
	require.NoError(t, err)
	p := Plan{
		Name:      "test",
		Executors: 1,
		Stages: []*Stage{
			{
				Name:   "dhcp4",
				Repeat: 5,
				DHCP4: &dhcp4.Config{
					Iface:  "eth0",
					HwAddr: dhcp4.HwAddr{macAddr},
				},
			},
			{
				Name:   "dns",
				Repeat: 5,
				DNS: &dns.Config{
					Endpoints: []string{"127.0.0.1:53"},
				},
			},
			{
				Name:   "etcd",
				Repeat: 5,
				ETCD: &etcd.Config{
					Endpoints:   []string{"localhost:1234"},
					DialTimeout: 5 * time.Second,
					KV: []*etcd.KV{
						{
							Get: &etcd.Get{
								Key: "foo",
							},
						},
					},
				},
			},
			{
				Name:   "http",
				Repeat: 5,
				HTTP: &http.Config{
					Payload: http.Payload{
						Method: "GET",
						URL:    "http://localhost:8000/",
						Header: map[string][]string{
							"Content-type": []string{
								"application/text",
							},
						},
					},
				},
			},
			{
				Name:       "http children",
				Repeat:     5,
				Concurrent: 2,
				Children: []*Stage{
					{
						Name: "http child",
						HTTP: &http.Config{
							Payload: http.Payload{
								Method: "GET",
								URL:    "http://localhost:8000/",
								Header: map[string][]string{
									"Content-type": []string{
										"application/text",
									},
								},
							},
						},
					},
				},
			},
			{
				Name:   "kafka",
				Repeat: 5,
				Kafka:  &kafka.Config{},
			},
			{
				Name:   "ldap",
				Repeat: 5,
				LDAP:   &ldap.Config{},
			},
			{
				Name:   "memcache",
				Repeat: 5,
				Memcache: &memcache.Config{
					Addrs: []string{"127.0.0.1:11211"},
					Ops: []*memcache.Op{
						{
							Set: &memcache.Set{
								Key:   "foo",
								Value: "baz",
							},
						},
						{
							Get: &memcache.Get{
								Key: "foo",
							},
						},
						{
							Delete: &memcache.Delete{
								Key: "bar",
							},
						},
					},
				},
			},
			{
				Name:   "redis",
				Repeat: 10,
				Redis: &redis.Config{
					Network: "eth1",
					Addr:    "127.0.0.1:1234",
					DB:      1,
					Commands: []*redis.Command{
						{
							Get: &redis.Get{
								Key: "foo",
							},
						},
					},
				},
			},
			{
				Name:   "sql",
				Repeat: 5,
				SQL: &sql.Config{
					MysqlDSN: "user:password@/dbname",
					Payloads: []*sql.Payload{
						{
							Exec: "Select * from users",
						},
					},
				},
			},
			{
				Name:   "ssh",
				Repeat: 5,
				SSH: &ssh.Config{
					Addr:    "127.0.0.1:22",
					User:    "root",
					Cmd:     util.StrPtr("ls /"),
					KeyFile: "/home/foo/.ssh/id_rsa",
				},
			},
			{
				Name:   "udp",
				Repeat: 5,
				UDP:    &udp.Config{},
			},
			{
				Name:   "websocket",
				Repeat: 5,
				Websocket: &websocket.Config{
					URL: "http://127.0.0.1:8888/ws",
					Ops: []*websocket.Op{
						{
							Read: true,
						},
						{
							Write: "foo",
						},
					},
				},
			},
		},
	}
	require.NoError(t, p.Validate())
	b, err := yaml.Marshal(p)
	require.NoError(t, err)
	require.NoError(t, ioutil.WriteFile("example.yaml", b, 0644))
}

func TestMissingStagesError(t *testing.T) {
	p := &Plan{}
	require.Error(t, p.Validate())
}

func TestDuplicateStageNameError(t *testing.T) {
	p := &Plan{
		Stages: []*Stage{
			{
				Name: "test",
				Children: []*Stage{
					{
						Name: "test",
					},
				},
			},
		},
	}
	require.Error(t, p.Validate())
}
