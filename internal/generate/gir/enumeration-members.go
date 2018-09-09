package gir

type Members []*Member

func (mm Members) forName(name string) *Member {
	for _, member := range mm {
		if member.Name == name {
			return member
		}
	}

	return nil
}

func (mm Members) mergeAddenda(addenda Members) {
	for _, addendaMember := range addenda {
		if member := mm.forName(addendaMember.Name); member != nil {
			member.mergeAddenda(addendaMember)
		}
	}
}
