package generate

type Interface struct {
	*Record
	VirtualMethods Methods `xml:"virtual-method"`
}

//no copy - virtual-method
//copy    - method
//copy    - glib:signal
