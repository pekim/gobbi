package generate

type Classes Records

func (cc Classes) init(ns *Namespace) {
	for _, class := range cc {
		class.init(ns)
	}
}

func (cc Classes) versionList() Versions {
	var versions Versions

	for _, c := range cc {
		if c.Version != "" {
			versions = append(versions, VersionNew(c.Version))
		}
	}

	return versions
}
