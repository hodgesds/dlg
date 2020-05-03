package memcache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseOp(t *testing.T) {
	op, err := ParseOp("get:foo")
	require.NoError(t, err)
	require.NotNil(t, op)
}
