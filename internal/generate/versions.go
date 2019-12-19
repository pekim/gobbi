package generate

import (
	"fmt"
	"github.com/blang/semver"
	"strings"
)

type versions struct {
	versions semver.Versions
}

func versionNew(value string) semver.Version {
	value = strings.TrimSpace(value)
	value = strings.TrimSuffix(value, ".")
	if value == "" {
		// default to 0 to allow for easy version comparison
		value = "0.0.0"
	}

	version, err := semver.ParseTolerant(value)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse version '%s' : %s", value, err))
	}

	return version
}

func (vv *versions) add(version semver.Version) {
	vv.versions = append(vv.versions, version)
}

func (vv *versions) len() int {
	return vv.versions.Len()
}

func (vv *versions) dedupe() {
	valuesMap := make(map[string]semver.Version)
	for _, v := range vv.versions {
		valuesMap[v.String()] = v
	}

	var newVersions []semver.Version
	for _, v := range valuesMap {
		newVersions = append(newVersions, v)
	}

	vv.versions = newVersions
}

func (vv *versions) sort() {
	vv.dedupe()
	semver.Sort(vv.versions)
}

func (vv versions) versionStringsGreaterThanOrEqual(version semver.Version) []string {
	var eligibleVersions []string

	for _, possibleVersion := range vv.versions {
		if possibleVersion.GTE(version) {
			eligibleVersions = append(eligibleVersions, possibleVersion.String())
		}
	}

	return eligibleVersions
}

func (vv *versions) String() string {
	str := ""
	for i, v := range vv.versions {
		str += v.String()
		if i < len(vv.versions)-1 {
			str += " "
		}
	}

	return str
}
