package v1

import "strconv"

func toUint32(value string, fallback uint32) uint32 {
	v, err := strconv.ParseUint(value, 10, 32)
	if v <= 0 || err != nil {
		return fallback
	}
	return uint32(v)
}
