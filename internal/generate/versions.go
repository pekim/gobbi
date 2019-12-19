package generate

import (
	"fmt"
	"github.com/blang/semver"
	"strings"
)

// noVersion represents the absence of a version.
//
// This is used for entities that do not have a version
// specified in gir file.
// Something that is unrestricted by version.
// In other words supported in all versions of a library.
var noVersion semver.Version = func() semver.Version {
	version, _ := semver.Parse("0.0.0")
	return version
}()

type versions struct {
	versions semver.Versions
}

func versionNew(value string) semver.Version {
	value = strings.TrimSpace(value)
	value = strings.TrimSuffix(value, ".")
	if value == "" {
		return noVersion
	}

	version, err := semver.ParseTolerant(value)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse version '%s' : %s", value, err))
	}

	return version
}

func versionIsNone(version semver.Version) bool {
	return version.EQ(noVersion)
}

func versionString(version semver.Version) string {
	if version.Patch == 0 {
		return strings.TrimSuffix(version.String(), ".0")
	} else {
		return version.String()
	}
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
