package generate

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

func (rr Records) entities() []Generatable {
	var generatables []Generatable

	for _, record := range rr {
		generatables = append(generatables, record)
	}

	return generatables
}

func (rr Records) forName(name string) *Record {
	for _, record := range rr {
		if record.Name == name {
			return record
		}
	}

	return nil
}

func (rr Records) mergeAddenda(addenda Records) {
	for _, addendaRecord := range addenda {
		if record := rr.forName(addendaRecord.Name); record != nil {
			record.mergeAddenda(addendaRecord)
		}
	}
}
