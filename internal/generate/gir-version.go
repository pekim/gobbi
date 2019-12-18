package generate

import (
	"fmt"
	"github.com/blang/semver"
	"strings"
)

type version struct {
	value   string
	version semver.Version
}

func versionNew(value string) *version {
	value = strings.TrimSuffix(value, ".")
	if value == "" {
		value = "0.0.0"
	}

	ver, err := semver.ParseTolerant(value)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse version '%s' : %s", value, err))
	}

	return &version{
		value:   value,
		version: ver,
	}
}

func (v *version) GTE(other *version) bool {
	return v.version.GTE(other.version)
}

func (v *version) LTE(other *version) bool {
	return v.version.LTE(other.version)
}

func (v *version) LT(other *version) bool {
	return v.version.LT(other.version)
}
