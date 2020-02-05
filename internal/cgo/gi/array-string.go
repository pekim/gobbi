package gi

// #include <girepository.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
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

	fmt.Println("satc", cArray)
	return unsafe.Pointer(cArray)
}

func FreeCStringArray(array unsafe.Pointer, count int) {
	fmt.Println("fcsa", array)
	cArray := (**C.gchar)(array)
	fmt.Println(cArray)

	cStrings := (*[1 << 28]*C.char)(unsafe.Pointer(cArray))[:count:count]

	for _, str := range cStrings {
		C.free(unsafe.Pointer(str))
	}

	C.free(array)
}
