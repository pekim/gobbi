// This is a generated file - DO NOT EDIT
// +build glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"C"
	"unsafe"
)

// g_list_copy_deep : unsupported parameter func : no type generator for CopyFunc (GCopyFunc) for param func
// GetBytes is a wrapper around the C function g_mapped_file_get_bytes.
func (recv *MappedFile) GetBytes() *Bytes {
	retC := C.g_mapped_file_get_bytes((*C.GMappedFile)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// g_slist_copy_deep : unsupported parameter func : no type generator for CopyFunc (GCopyFunc) for param func
// FreeToBytes is a wrapper around the C function g_string_free_to_bytes.
func (recv *String) FreeToBytes() *Bytes {
	retC := C.g_string_free_to_bytes((*C.GString)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckFormatString is a wrapper around the C function g_variant_check_format_string.
func (recv *Variant) CheckFormatString(formatString string, copyOnly bool) bool {
	c_format_string := C.CString(formatString)
	defer C.free(unsafe.Pointer(c_format_string))

	c_copy_only :=
		boolToGboolean(copyOnly)

	retC := C.g_variant_check_format_string((*C.GVariant)(recv.native), c_format_string, c_copy_only)
	retGo := retC == C.TRUE

	return retGo
}
