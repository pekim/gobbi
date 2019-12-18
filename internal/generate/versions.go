package generate

import "sort"

type Versions []*version

func (vv Versions) append(v *version) Versions {
	return append(vv, v)
}

func (vv Versions) Len() int {
	return len(vv)
}

func (vv Versions) Less(i, j int) bool {
	return vv[i].LT(vv[j])
}

func (vv Versions) Swap(i, j int) {
	vv[i], vv[j] = vv[j], vv[i]
}

func (vv Versions) sort() {
	sort.Sort(vv)
}

func (vv Versions) dedupe() Versions {
	valuesMap := make(map[string]*version)
	for _, v := range vv {
		valuesMap[v.value] = v
	}

	var newVersions []*version
	for _, v := range valuesMap {
		newVersions = append(newVersions, v)
	}

	return newVersions
}

func (vv Versions) versionStringsGreaterThanOrEqual(version version) []string {
	var eligibleVersions []string

	for _, possibleVersion := range vv {
		if possibleVersion.version.GTE(version.version) {
			eligibleVersions = append(eligibleVersions, possibleVersion.value)
		}
	}

	return eligibleVersions
}

func (vv Versions) String() string {
	str := ""
	for i, v := range vv {
		str += v.value
		if i < len(vv)-1 {
			str += " "
		}
	}

	return str
}
