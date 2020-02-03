package gi

// #include <girepository.h>
import "C"

import (
	"unsafe"
)

// incrPointer increments a pointer by an increment.
func incrPointer(ptr unsafe.Pointer, incr int) unsafe.Pointer {
	newPtr := uintptr(ptr) + uintptr(incr)
	return unsafe.Pointer(newPtr)
}

// incrPointerPointer increments a pointer-to-a-pointer to the next pointer.
func incrPointerPointer(ptr unsafe.Pointer) unsafe.Pointer {
	return incrPointer(ptr, C.sizeof_gpointer)
}
