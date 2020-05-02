package util

import (
	"fmt"
	"strings"
)

// ParseMemcacheKV is used for parsing Memcache KV pairs.
func ParseMemcacheKV(s string) (string, string, error) {
	i := strings.Index(s, ":")
	if i < 0 {
		return "", "", fmt.Errorf("invalid kv pair: %q", s)
	}
	return s[:i], s[i:], nil
}
