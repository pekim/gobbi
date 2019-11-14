package generate

type ReturnValue struct {
	Namespace *Namespace

	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	//Doc               *Doc   `xml:"doc"`
	Type *Type `xml:"type"`
	//Array             *Array `xml:"array"`
}

func (r *ReturnValue) init(ns *Namespace) {
	r.Namespace = ns

	if r.Type != nil {
		r.Type.init(ns)
	}

	//if r.Array != nil && r.Array.Type != nil {
	//	r.Array.init(ns)
	//}
}
