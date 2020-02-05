package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

func StringArrayToC(value []string, nilTerminated bool) unsafe.Pointer {
	if value == nil || len(value) == 0 {
		return nil
	}

	cStringsCount := len(value)
	if nilTerminated {
		cStringsCount++
	}

	arraySize := cStringsCount * C.sizeof_gpointer
	cArray := (**C.gchar)(C.malloc(C.ulong(arraySize)))
	cStrings := (*[1 << 28]*C.char)(unsafe.Pointer(cArray))[:cStringsCount:cStringsCount]

	for i, str := range value {
		cStrings[i] = C.CString(str)
	}

	if nilTerminated {
		cStrings[cStringsCount-1] = nil
	}

	return unsafe.Pointer(cArray)
}
