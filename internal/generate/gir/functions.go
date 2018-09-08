package gir

type Functions []*Function

func (ff Functions) fixup(ns *Namespace) {
	for _, function := range ff {
		function.fixup(ns)
	}
}

func (ff Functions) versionList() Versions {
	var versions Versions

	for _, f := range ff {
		if f.Version != "" {
			versions = append(versions, VersionNew(f.Version))
		}
	}

	return versions
}
