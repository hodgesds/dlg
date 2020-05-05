package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSNMPEndpoint(t *testing.T) {
	h, p, err := ParseSNMPEndpoint("foo:123")
	require.NoError(t, err)
	require.NotEmpty(t, h)
	require.NotZero(t, p)

	_, _, err = ParseSNMPEndpoint("foo")
	require.Error(t, err)

	_, _, err = ParseSNMPEndpoint("foo:f")
	require.Error(t, err)

}
