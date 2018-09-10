package generate

type Array struct {
	Namespace *Namespace

	Type           *Type  `xml:"type"`
	FixedSize      int    `xml:"fixed-size,attr"`
	Length         int    `xml:"length,attr"`
	ZeroTerminated string `xml:"zero-terminated,attr"`
}
