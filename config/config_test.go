package config

import (
	"io/ioutil"
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
	p := Plan{
		Name:      "test plan",
		Executors: 1,
		Stages: []*Stage{
			{
				Name:   "dhcp4",
				Repeat: 5,
				DHCP4:  &dhcp4.Config{},
			},
			{
				Name:   "dns",
				Repeat: 5,
				DNS:    &dns.Config{},
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
				Concurrent: true,
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
				Name:     "memcache",
				Repeat:   5,
				Memcache: &memcache.Config{},
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
				Name:      "websocket",
				Repeat:    5,
				Websocket: &websocket.Config{},
			},
		},
	}
	require.NoError(t, p.Validate())
	c := &Config{
		Plan: p,
	}
	b, err := yaml.Marshal(c)
	require.NoError(t, err)
	ioutil.WriteFile("example.yaml", b, 0644)
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
