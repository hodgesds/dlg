package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMemcacheKV(t *testing.T) {
	k, v, err := ParseMemcacheKV("foo:bar")
	require.NoError(t, err)
	require.NotEmpty(t, k)
	require.NotEmpty(t, v)

	_, _, err = ParseMemcacheKV("foo")
	require.Error(t, err)

	k, v, err = ParseMemcacheKV("foo:bar:baz")
	require.NoError(t, err)
	require.NotEmpty(t, k)
	require.Equal(t, "foo", k)
	require.NotEmpty(t, "bar:baz", v)
}
