package etcd

import (
	"testing"
	"time"

	"github.com/hodgesds/dlg/util"
	"github.com/stretchr/testify/require"
)

func TestETCDClientConfig(t *testing.T) {
	c := &Config{
		DialKeepAliveTime:    util.DurPtr(1 * time.Second),
		DialKeepAliveTimeout: util.DurPtr(1 * time.Second),
		MaxCallSendMsgSize:   util.IntPtr(10),
		MaxCallRecvMsgSize:   util.IntPtr(10),
		Username:             util.StrPtr("foo"),
		Password:             util.StrPtr("foo"),
		RejectOldCluster:     util.BoolPtr(false),
	}
	require.NotNil(t, c.ClientConfig())
}
