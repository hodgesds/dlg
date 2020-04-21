package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDurPtr(t *testing.T) {
	require.NotNil(t, DurPtr(1*time.Second))
}

func TestStrPtr(t *testing.T) {
	require.NotNil(t, StrPtr("foo"))
}

func TestBoolPtr(t *testing.T) {
	require.NotNil(t, BoolPtr(true))
}

func TestIntPtr(t *testing.T) {
	require.NotNil(t, IntPtr(1))
}

func TestUintPtr(t *testing.T) {
	require.NotNil(t, UintPtr(1))
}

func TestInt8Ptr(t *testing.T) {
	require.NotNil(t, Int8Ptr(1))
}

func TestUint8Ptr(t *testing.T) {
	require.NotNil(t, Uint8Ptr(1))
}

func TestInt16Ptr(t *testing.T) {
	require.NotNil(t, Int16Ptr(1))
}

func TestUint16Ptr(t *testing.T) {
	require.NotNil(t, Uint16Ptr(1))
}

func TestInt32Ptr(t *testing.T) {
	require.NotNil(t, Int32Ptr(1))
}

func TestUint32Ptr(t *testing.T) {
	require.NotNil(t, Uint32Ptr(1))
}

func TestInt64Ptr(t *testing.T) {
	require.NotNil(t, Int64Ptr(1))
}

func TestUint64Ptr(t *testing.T) {
	require.NotNil(t, Uint64Ptr(1))
}

func TestFloat32Ptr(t *testing.T) {
	require.NotNil(t, Float32Ptr(1))
}

func TestFloat64Ptr(t *testing.T) {
	require.NotNil(t, Float64Ptr(1))
}
