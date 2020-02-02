package gi

import "unsafe"

func incptr(ptr unsafe.Pointer, inc int) unsafe.Pointer {
	newPtr := uintptr(ptr) + uintptr(inc)
	return unsafe.Pointer(newPtr)
}
