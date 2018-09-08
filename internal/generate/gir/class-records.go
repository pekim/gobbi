package gir

type Records []*Record

func (rr Records) fixup(ns *Namespace) {
	for _, record := range rr {
		record.fixup(ns)
	}
}

func (rr Records) versionList() Versions {
	var versions Versions

	for _, r := range rr {
		if r.Version != "" {
			versions = append(versions, VersionNew(r.Version))
		}
	}

	return versions
}

type Classes Records

func (cc Classes) fixup(ns *Namespace) {
	for _, class := range cc {
		class.fixup(ns)
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
