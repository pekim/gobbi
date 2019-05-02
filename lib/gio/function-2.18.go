// This is a generated file - DO NOT EDIT
// +build gio_2.18 gio_2.20 gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"C"
	"unsafe"
)

// ContentTypeFromMimeType is a wrapper around the C function g_content_type_from_mime_type.
func ContentTypeFromMimeType(mimeType string) string {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	retC := C.g_content_type_from_mime_type(c_mime_type)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ContentTypeGuessForTree is a wrapper around the C function g_content_type_guess_for_tree.
func ContentTypeGuessForTree(root *File) []string {
	c_root := (*C.GFile)(root.ToC())

	retC := C.g_content_type_guess_for_tree(c_root)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}
