package generate

type Aliases []*Alias

func (aa Aliases) init(ns *Namespace) {
	for _, alias := range aa {
		alias.init(ns)
	}
}

func (aa Aliases) byName(name string) (*Alias, bool) {
	for _, alias := range aa {
		if alias.Name == name {
			return alias, true
		}
	}

	return nil, false
}

//func (aa Aliases) generateSys(f *jen.File, version semver.Version) {
//	f.Comment("aliases")
//
//	for _, a := range aa {
//		a.generateSys(f, version)
//	}
//
//	f.Line()
//}
