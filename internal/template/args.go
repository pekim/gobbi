package template

// #include <stdlib.h>
/*

static inline gchar* get_string(gchar** strings, int n) {
	return strings[n];
}

static inline void set_string(gchar** strings, int n, gchar* str) {
	strings[n] = str;
}

*/
import "C"
import (
	"unsafe"
)

/*
argsIn converts a string slice in to argc/argv pointers
suitable for use with some library functions.

The caller must be sure to free the returned argv at some point.
*/
func argsIn(args []string) (C.int, **C.char) {
	argc := (C.gint)(len(args))

	argv := (**C.gchar)(C.malloc(C.sizeof_gpointer * (C.ulong)(argc)))

	for i, arg := range args {
		cstr := C.CString(arg)
		C.set_string(argv, C.int(i), (*C.gchar)(cstr))
	}

	return argc,
		(**C.char)(unsafe.Pointer(argv))
}

func argsOut(argc C.int, argv **C.char) []string {
	args := make([]string, argc)
	for i := 0; i < int(argc); i++ {
		cstr := C.get_string(argv, C.int(i))
		args[i] = C.GoString((*C.char)(cstr))
		C.free(unsafe.Pointer(cstr))
	}

	return args
}
