package generate

import "github.com/dave/jennifer/jen"

type Signals []*Signal

func (ss Signals) init(ns *Namespace) {
	for _, signal := range ss {
		signal.init(ns)
	}
}

func (ss Signals) generate(g *jen.Group, version *Version) {
	for _, signal := range ss {
		signal.generate(g, version)
	}
}

//func (rr Records) versionList() Versions {
//	var versions Versions
//
//	for _, r := range rr {
//		if r.Version != "" {
//			versions = append(versions, VersionNew(r.Version))
//		}
//
//		for _, ctor := range r.Constructors {
//			if ctor.Version != "" {
//				versions = append(versions, VersionNew(ctor.Version))
//			}
//		}
//		for _, m := range r.Methods {
//			if m.Version != "" {
//				versions = append(versions, VersionNew(m.Version))
//			}
//		}
//	}
//
//	return versions
//}
//
//func (rr Records) entities() []Generatable {
//	var generatables []Generatable
//
//	for _, record := range rr {
//		generatables = append(generatables, record)
//	}
//
//	return generatables
//}
//
//func (rr Records) forName(name string) *Record {
//	for _, record := range rr {
//		if record.Name == name {
//			return record
//		}
//	}
//
//	return nil
//}
//
//func (rr Records) mergeAddenda(addenda Records) {
//	for _, addendaRecord := range addenda {
//		if record := rr.forName(addendaRecord.Name); record != nil {
//			record.mergeAddenda(addendaRecord)
//		}
//	}
//}
