// Code generated - DO NOT EDIT.
// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Close is a wrapper around the C function g_close.
func Close(fd int32) (bool, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_close(c_fd, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetNumProcessors is a wrapper around the C function g_get_num_processors.
func GetNumProcessors() uint32 {
	retC := C.g_get_num_processors()
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_unix_fd_add : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// Unsupported : g_unix_fd_add_full : unsupported parameter function : no type generator for UnixFDSourceFunc (GUnixFDSourceFunc) for param function

// UnixFdSourceNew is a wrapper around the C function g_unix_fd_source_new.
func UnixFdSourceNew(fd int32, condition IOCondition) *Source {
	c_fd := (C.gint)(fd)

	c_condition := (C.GIOCondition)(condition)

	retC := C.g_unix_fd_source_new(c_fd, c_condition)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function g_markup_parse_context_ref.
func (recv *MarkupParseContext) Ref() *MarkupParseContext {
	retC := C.g_markup_parse_context_ref((*C.GMarkupParseContext)(recv.native))
	retGo := MarkupParseContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_markup_parse_context_unref.
func (recv *MarkupParseContext) Unref() {
	C.g_markup_parse_context_unref((*C.GMarkupParseContext)(recv.native))

	return
}

// Unsupported : g_source_add_unix_fd : no return generator

// Unsupported : g_source_modify_unix_fd : unsupported parameter tag : no type generator for gpointer (gpointer) for param tag

// Unsupported : g_source_query_unix_fd : unsupported parameter tag : no type generator for gpointer (gpointer) for param tag

// Unsupported : g_source_remove_unix_fd : unsupported parameter tag : no type generator for gpointer (gpointer) for param tag

// SetReadyTime is a wrapper around the C function g_source_set_ready_time.
func (recv *Source) SetReadyTime(readyTime int64) {
	c_ready_time := (C.gint64)(readyTime)

	C.g_source_set_ready_time((*C.GSource)(recv.native), c_ready_time)

	return
}

// VariantNewFromBytes is a wrapper around the C function g_variant_new_from_bytes.
func VariantNewFromBytes(type_ *VariantType, bytes *Bytes, trusted bool) *Variant {
	c_type := (*C.GVariantType)(C.NULL)
	if type_ != nil {
		c_type = (*C.GVariantType)(type_.ToC())
	}

	c_bytes := (*C.GBytes)(C.NULL)
	if bytes != nil {
		c_bytes = (*C.GBytes)(bytes.ToC())
	}

	c_trusted :=
		boolToGboolean(trusted)

	retC := C.g_variant_new_from_bytes(c_type, c_bytes, c_trusted)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDataAsBytes is a wrapper around the C function g_variant_get_data_as_bytes.
func (recv *Variant) GetDataAsBytes() *Bytes {
	retC := C.g_variant_get_data_as_bytes((*C.GVariant)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}
