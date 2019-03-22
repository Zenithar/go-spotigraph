package version

import (
	"encoding/json"
	"fmt"
)

// Build information. Populated at build-time.
var (
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion string
)

// Map provides the iterable version information.
var Map = map[string]string{
	"version":   Version,
	"revision":  Revision,
	"branch":    Branch,
	"buildUser": BuildUser,
	"buildDate": BuildDate,
	"goVersion": GoVersion,
}

// Full returns full composed version string
func Full() string {
	return fmt.Sprintf("%s [%s] (Go: %s, User: %s, Date: %s)", Version, Branch, GoVersion, BuildUser, BuildDate)
}

// JSON returns json representation of build info
func JSON() string {
	payload, err := json.Marshal(Map)
	if err != nil {
		panic(err)
	}
	return string(payload)
}
