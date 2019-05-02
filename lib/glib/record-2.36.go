// This is a generated file - DO NOT EDIT
// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"C"
	"unsafe"
)

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

// GetDataAsBytes is a wrapper around the C function g_variant_get_data_as_bytes.
func (recv *Variant) GetDataAsBytes() *Bytes {
	retC := C.g_variant_get_data_as_bytes((*C.GVariant)(recv.native))
	retGo := BytesNewFromC(unsafe.Pointer(retC))

	return retGo
}
