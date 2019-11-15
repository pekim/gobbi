package gi

// #include <girepository.h>
import "C"
import "unsafe"

type Argument C.GIArgument

func (a Argument) Uint32() uint32 {
	return (uint32)(*(*C.guint)(unsafe.Pointer(&a)))
}
