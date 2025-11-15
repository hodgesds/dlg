module github.com/hodgesds/dlg

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

go 1.21

require (
	github.com/ClickHouse/clickhouse-go v1.5.4
	github.com/eclipse/paho.mqtt.golang v1.4.3
	github.com/elastic/go-elasticsearch/v8 v8.11.1
	github.com/gin-gonic/gin v1.9.1
	github.com/gocql/gocql v1.6.0
	github.com/go-ldap/ldap/v3 v3.4.6
	github.com/go-redis/redis/v7 v7.4.1
	github.com/go-sql-driver/mysql v1.7.1
	github.com/google/uuid v1.5.0
	github.com/gorilla/websocket v1.5.1
	github.com/gosnmp/gosnmp v1.37.0
	github.com/influxdata/influxdb-client-go/v2 v2.13.0
	github.com/insomniacslk/dhcp v0.0.0-20231206064809-8c70d406f6d2
	github.com/jonboulle/clockwork v0.4.0
	github.com/lib/pq v1.10.9
	github.com/machinebox/graphql v0.2.2
	github.com/miekg/dns v1.1.57
	github.com/prometheus/client_golang v1.17.0
	github.com/rainycape/memcache v0.0.0-20150622160815-1031fa0ce2f2
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.18.2
	github.com/stretchr/testify v1.8.4
	go.etcd.io/etcd v3.3.27+incompatible
	go.mongodb.org/mongo-driver v1.13.1
	go.uber.org/multierr v1.11.0
	golang.org/x/crypto v0.17.0
	golang.org/x/time v0.5.0
	google.golang.org/grpc v1.60.1
	gopkg.in/yaml.v2 v2.4.0
)
