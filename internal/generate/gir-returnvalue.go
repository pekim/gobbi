package generate

type ReturnValue struct {
	Namespace *Namespace

	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	//Doc               *Doc   `xml:"doc"`
	Type  *Type  `xml:"type"`
	Array *Array `xml:"array"`
}

func (r *ReturnValue) init(ns *Namespace) {
	r.Namespace = ns
	r.Type.init(ns)
	r.Array.init(ns)
}

func (r *ReturnValue) isVoid() bool {
	if r == nil {
		return true
	}

	if r.Type == nil && r.Array == nil {
		return true
	}

	if r.Type != nil && r.Type.Name == "none" {
		return true
	}

	return false
}

func (r ReturnValue) isSupported() (bool, string) {
	if r.isVoid() {
		return true, ""
	}

	if r.Array != nil {
		return false, "has array return"
	}

	if r.Type.isCallback() {
		return false, "has callback return "
	}

	return true, ""
}
