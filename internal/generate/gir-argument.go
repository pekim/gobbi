package generate

type Argument struct {
	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          bool   `xml:"nullable,attr"`
	Type              *Type  `xml:"type"`
	Array             *Array `xml:"array"`
}

func (a Argument) transferNone() bool {
	return a.TransferOwnership == "" || a.TransferOwnership == "none"
}

func (a Argument) transferContainer() bool {
	return a.TransferOwnership == "container"
}

func (a Argument) transferFull() bool {
	return a.TransferOwnership == "full"
}
