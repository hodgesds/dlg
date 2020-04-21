package util

import "time"

// DurPtr returns a pointer for a time.Duration.
func DurPtr(d time.Duration) *time.Duration {
	return &d
}

// StrPtr returns a pointer for a string.
func StrPtr(s string) *string {
	return &s
}

// BoolPtr returns a pointer for a bool.
func BoolPtr(b bool) *bool {
	return &b
}

// IntPtr returns a pointer for a int.
func IntPtr(i int) *int {
	return &i
}

// UintPtr returns a pointer for a uint.
func UintPtr(i uint) *uint {
	return &i
}

// Int8Ptr returns a pointer for a int8.
func Int8Ptr(i int8) *int8 {
	return &i
}

// Uint8Ptr returns a pointer for a uint8.
func Uint8Ptr(i uint8) *uint8 {
	return &i
}

// Int16Ptr returns a pointer for a int16.
func Int16Ptr(i int16) *int16 {
	return &i
}

// Uint16Ptr returns a pointer for a uint16.
func Uint16Ptr(i uint16) *uint16 {
	return &i
}

// Int32Ptr returns a pointer for a int32.
func Int32Ptr(i int32) *int32 {
	return &i
}

// Uint32Ptr returns a pointer for a uint32.
func Uint32Ptr(i uint32) *uint32 {
	return &i
}

// Int64Ptr returns a pointer for a int64.
func Int64Ptr(i int64) *int64 {
	return &i
}

// Uint64Ptr returns a pointer for a uint64.
func Uint64Ptr(i uint64) *uint64 {
	return &i
}

// Float32Ptr returns a pointer for a float32.
func Float32Ptr(f float32) *float32 {
	return &f
}

// Float64Ptr returns a pointer for a float64.
func Float64Ptr(f float64) *float64 {
	return &f
}
