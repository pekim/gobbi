package generate

type Callbacks []*Callback

func (cc Callbacks) init(ns *Namespace) {
	for _, callback := range cc {
		callback.init(ns, nil, "")
	}
}

func (cc Callbacks) versionList() Versions {
	var versions Versions

	for _, c := range cc {
		if c.Version != "" {
			versions = append(versions, VersionNew(c.Version))
		}

	}

	return versions
}

func (cc Callbacks) entities() []Generatable {
	var generatables []Generatable

	for _, callback := range cc {
		generatables = append(generatables, callback)
	}

	return generatables
}

func (cc Callbacks) forName(name string) *Callback {
	for _, callback := range cc {
		if callback.Name == name {
			return callback
		}
	}

	return nil
}
