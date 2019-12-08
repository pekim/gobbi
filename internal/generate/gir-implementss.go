package generate

type Implementss []*Implements

func (ii Implementss) byName(name string) *Implements {
	for _, implements := range ii {
		if implements.Name == name {
			return implements
		}
	}

	return nil
}

type Implements struct {
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`
}
