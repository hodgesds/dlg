module github.com/hodgesds/dlg

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0 // indirect

go 1.14

require (
	github.com/ClickHouse/clickhouse-go v1.4.0
	github.com/coreos/etcd v3.3.20+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ldap/ldap/v3 v3.1.8
	github.com/go-redis/redis/v7 v7.2.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/hugelgupf/socketpair v0.0.0-20190730060125-05d35a94e714 // indirect
	github.com/insomniacslk/dhcp v0.0.0-20200420235442-ed3125c2efe7
	github.com/jonboulle/clockwork v0.1.0
	github.com/lib/pq v1.3.0
	github.com/mdlayher/ethernet v0.0.0-20190606142754-0394541c37b7 // indirect
	github.com/mdlayher/raw v0.0.0-20191009151244-50f2db8cc065 // indirect
	github.com/prometheus/client_golang v1.5.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.4.0
	github.com/u-root/u-root v6.0.0+incompatible // indirect
	go.etcd.io/etcd v3.3.20+incompatible
	go.uber.org/multierr v1.5.0
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.0.0-20190510104115-cbcb75029529
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1
	google.golang.org/grpc v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.2.8
	sigs.k8s.io/yaml v1.2.0 // indirect
)
