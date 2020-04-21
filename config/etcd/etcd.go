package etcd

import (
	"time"

	"go.etcd.io/etcd/clientv3"
)

// Config is configuration for ETCD.
type Config struct {
	Endpoints            []string       `yaml:"endpoints"`
	DialTimeout          time.Duration  `yaml:"dialTimeout"`
	DialKeepAliveTime    *time.Duration `yaml:"dialKeepAliveTime,omitempty"`
	DialKeepAliveTimeout *time.Duration `yaml:"dialKeepAliveTimeout,omitempty"`
	MaxCallSendMsgSize   *int           `yaml:"maxCallSendMsgSize,omitempty"`
	MaxCallRecvMsgSize   *int           `yaml:"maxCallRecvMsgSize,omitempty"`
	Username             *string        `yaml:"username,omitempty"`
	Password             *string        `yaml:"password,omitempty"`
	RejectOldCluster     *bool          `yaml:"rejectOldCluster,omitempty"`
	KV                   []*KV          `yaml:"kv,omitempty"`
}

// ClientConfig returns an clientv3.ClientConfig, this is a convience method
// for keeping top level configs clean.
func (c *Config) ClientConfig() clientv3.Config {
	cc := clientv3.Config{
		Endpoints:   c.Endpoints,
		DialTimeout: c.DialTimeout,
	}
	if c.DialKeepAliveTime != nil {
		cc.DialKeepAliveTime = *c.DialKeepAliveTime
	}
	if c.DialKeepAliveTimeout != nil {
		cc.DialKeepAliveTimeout = *c.DialKeepAliveTimeout
	}
	if c.MaxCallSendMsgSize != nil {
		cc.MaxCallSendMsgSize = *c.MaxCallSendMsgSize
	}
	if c.MaxCallRecvMsgSize != nil {
		cc.MaxCallRecvMsgSize = *c.MaxCallRecvMsgSize
	}
	if c.Username != nil {
		cc.Username = *c.Username
	}
	if c.Password != nil {
		cc.Password = *c.Password
	}
	if c.RejectOldCluster != nil {
		cc.RejectOldCluster = *c.RejectOldCluster
	}
	return cc
}

// KV is used for configuring ETCD KV operations. Order is always in struct
// order.
type KV struct {
	Compact *Compact `yaml:"compact,omitempty"`
	Delete  *Delete  `yaml:"delete,omitempty"`
	Get     *Get     `yaml:"get,omitempty"`
	Put     *Put     `yaml:"put,omitempty"`
}

// Put is an ETCD put KV op.
type Put struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
	Opts  Opts   `yaml:"opts,omitempty"`
}

// Get is an ETCD get KV op.
type Get struct {
	Key  string `yaml:"key"`
	Opts Opts   `yaml:"opts,omitempty"`
}

// Delete is an ETCD delete KV op.
type Delete struct {
	Key  string `yaml:"key"`
	Opts Opts   `yaml:"opts,omitempty"`
}

// Compact is an ETCD compact KV op.
type Compact struct {
	Rev  int64 `yaml:"rev"`
	Opts Opts  `yaml:"opts,omitempty"`
}

// Opt is used for configuring ETCD options.
type Opt struct {
	// clientv3.Op configs
	IsCountOnly    *bool   `yaml:"isCountOnly"`
	IsDelete       *bool   `yaml:"isDelete"`
	IsGet          *bool   `yaml:"isGet"`
	IsKeysOnly     *bool   `yaml:"isKeysOnly"`
	IsPut          *bool   `yaml:"isPut"`
	IsSerializable *bool   `yaml:"isSerializable"`
	KeyBytes       *bool   `yaml:"keyBytes"`
	MaxCreateRev   *bool   `yaml:"maxCreateRev"`
	MaxModRev      *bool   `yaml:"maxModRev"`
	MinCreateRev   *bool   `yaml:"minCreateRev"`
	MinModRev      *bool   `yaml:"minModRev"`
	RangeBytes     *bool   `yaml:"rangeBytes"`
	Rev            *bool   `yaml:"rev"`
	ValueBytes     *bool   `yaml:"valueBytes"`
	WithKeyBytes   *[]byte `yaml:"withKeyBytes"`
	WithRangeBytes *[]byte `yaml:"withRangeBytes"`
	WithValueBytes *[]byte `yaml:"withValueBytes"`

	// clientv3.OpOption configs
	WithCountOnly      *bool   `yaml:"withCountOnly"`
	WithCreatedNotify  *bool   `yaml:"withCreatedNotify"`
	WithFilterDelete   *bool   `yaml:"withFilterDelete"`
	WithFilterPut      *bool   `yaml:"withFilterPut"`
	WithFirstCreate    *bool   `yaml:"withFirstCreate"`
	WithFirstKey       *bool   `yaml:"withFirstKey"`
	WithFirstRev       *bool   `yaml:"withFirstRev"`
	WithFragment       *bool   `yaml:"withFragment"`
	WithFromKey        *bool   `yaml:"withFromKey"`
	WithIgnoreLease    *bool   `yaml:"withIgnoreLease"`
	WithIgnoreValue    *bool   `yaml:"withIgnoreValue"`
	WithKeysOnly       *bool   `yaml:"withKeysOnly"`
	WithLastCreate     *bool   `yaml:"withLastCreate"`
	WithLastKey        *bool   `yaml:"withLastKey"`
	WithLastRev        *bool   `yaml:"withLastRev"`
	WithLease          *int64  `yaml:"withLease"`
	WithLimit          *int64  `yaml:"withLimit"`
	WithMaxCreateRev   *int64  `yaml:"withMaxCreateRev"`
	WithMaxModRev      *int64  `yaml:"withMaxModRev"`
	WithMinCreateRev   *int64  `yaml:"withMinCreateRev"`
	WithMinModRev      *int64  `yaml:"withMinModRev"`
	WithPrefix         *bool   `yaml:"withPrefix"`
	WithPrevKV         *bool   `yaml:"withPrevKV"`
	WithProgressNotify *bool   `yaml:"withProgressNotify"`
	WithRange          *string `yaml:"withRange"`
	WithRev            *int64  `yaml:"withRev"`
	WithSerializable   *bool   `yaml:"withSerializable"`
}

// Opts is a slice of Opt.
type Opts []*Opt

// Opts returns a set of clientv3.OpOption from a ETCDOpts.
func (e Opts) Opts() []clientv3.OpOption {
	if len(e) == 0 {
		return nil
	}
	opts := []clientv3.OpOption{}
	for _, o := range e {
		if o.WithCountOnly != nil && *o.WithCountOnly {
			opts = append(opts, clientv3.WithCountOnly())
		}
		if o.WithCreatedNotify != nil && *o.WithCreatedNotify {
			opts = append(opts, clientv3.WithCreatedNotify())
		}
		if o.WithFilterDelete != nil && *o.WithFilterDelete {
			opts = append(opts, clientv3.WithFilterDelete())
		}
		if o.WithFilterPut != nil && *o.WithFilterPut {
			opts = append(opts, clientv3.WithFilterPut())
		}
		if o.WithFirstCreate != nil && *o.WithFirstCreate {
			opts = append(opts, clientv3.WithFirstCreate()...)
		}
		if o.WithFirstKey != nil && *o.WithFirstKey {
			opts = append(opts, clientv3.WithFirstKey()...)
		}
		if o.WithFirstRev != nil && *o.WithFirstRev {
			opts = append(opts, clientv3.WithFirstRev()...)
		}
		if o.WithFragment != nil && *o.WithFragment {
			opts = append(opts, clientv3.WithFragment())
		}
		if o.WithFromKey != nil && *o.WithFromKey {
			opts = append(opts, clientv3.WithFromKey())
		}
		if o.WithIgnoreLease != nil && *o.WithIgnoreLease {
			opts = append(opts, clientv3.WithIgnoreLease())
		}
		if o.WithIgnoreValue != nil && *o.WithIgnoreValue {
			opts = append(opts, clientv3.WithIgnoreValue())
		}
		if o.WithKeysOnly != nil && *o.WithKeysOnly {
			opts = append(opts, clientv3.WithKeysOnly())
		}
		if o.WithLastCreate != nil && *o.WithLastCreate {
			opts = append(opts, clientv3.WithLastCreate()...)
		}
		if o.WithLastKey != nil && *o.WithLastKey {
			opts = append(opts, clientv3.WithLastKey()...)
		}
		if o.WithLastRev != nil && *o.WithLastRev {
			opts = append(opts, clientv3.WithLastRev()...)
		}
		if o.WithLease != nil {
			opts = append(opts, clientv3.WithLease(clientv3.LeaseID(*o.WithLease)))
		}
		if o.WithLimit != nil {
			opts = append(opts, clientv3.WithLimit(*o.WithLimit))
		}
		if o.WithMaxCreateRev != nil {
			opts = append(opts, clientv3.WithMaxCreateRev(*o.WithMaxCreateRev))
		}
		if o.WithMaxModRev != nil {
			opts = append(opts, clientv3.WithMaxModRev(*o.WithMaxModRev))
		}
		if o.WithMinCreateRev != nil {
			opts = append(opts, clientv3.WithMinCreateRev(*o.WithMinCreateRev))
		}
		if o.WithMinModRev != nil {
			opts = append(opts, clientv3.WithMinModRev(*o.WithMinModRev))
		}
		if o.WithPrefix != nil && *o.WithPrefix {
			opts = append(opts, clientv3.WithPrefix())
		}
		if o.WithPrevKV != nil && *o.WithPrevKV {
			opts = append(opts, clientv3.WithPrevKV())
		}
		if o.WithProgressNotify != nil && *o.WithProgressNotify {
			opts = append(opts, clientv3.WithProgressNotify())
		}
		if o.WithRange != nil {
			opts = append(opts, clientv3.WithRange(*o.WithRange))
		}
		if o.WithRev != nil {
			opts = append(opts, clientv3.WithRev(*o.WithRev))
		}
		if o.WithSerializable != nil && *o.WithSerializable {
			opts = append(opts, clientv3.WithSerializable())
		}
	}
	return opts
}
