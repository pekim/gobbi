package generate

type Classes []*Class

func (cc Classes) init(ns *Namespace) {
	for _, class := range cc {
		class.init(ns)
	}
}

func (cc Classes) forName(name string) *Class {
	for _, class := range cc {
		if class.Name == name {
			return class
		}
	}

	return nil
}
