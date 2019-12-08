package generate

type Interface struct {
	*Record
	VirtualMethods Methods `xml:"virtual-method"`
}
