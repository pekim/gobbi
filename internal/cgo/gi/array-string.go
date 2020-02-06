package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func StringArrayToC(strings []string, nilTerminated bool) []unsafe.Pointer {
	if strings == nil || len(strings) == 0 {
		return nil
	}

	cStringsCount := len(strings)
	if nilTerminated {
		cStringsCount++
	}

	cStrings := make([]unsafe.Pointer, cStringsCount, cStringsCount)
	for i, str := range strings {
		cStrings[i] = unsafe.Pointer(C.CString(str))
	}

	if nilTerminated {
		cStrings[cStringsCount-1] = nil
	}

	return cStrings
}

func FreeCStringArray(cArray []unsafe.Pointer, count int) {
	if cArray == nil || count == 0 {
		return
	}

	for _, str := range cArray {
		fmt.Println("free", str, C.GoString((*C.char)(str)))
		C.free(unsafe.Pointer(str))
	}
}
