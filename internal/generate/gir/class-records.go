package gir

type Records []*Record

func (rr Records) init(ns *Namespace) {
	for _, record := range rr {
		record.init(ns)
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
