// This is a generated file - DO NOT EDIT
// +build glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_bytes_new : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_bytes_new_static : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_bytes_new_take : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_bytes_new_with_free_func : unsupported parameter data : no type generator for guint8 () for array param data

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

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

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// AddUnixFd is a wrapper around the C function g_source_add_unix_fd.
func (recv *Source) AddUnixFd(fd int32, events IOCondition) uintptr {
	c_fd := (C.gint)(fd)

	c_events := (C.GIOCondition)(events)

	retC := C.g_source_add_unix_fd((*C.GSource)(recv.native), c_fd, c_events)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// ModifyUnixFd is a wrapper around the C function g_source_modify_unix_fd.
func (recv *Source) ModifyUnixFd(tag uintptr, newEvents IOCondition) {
	c_tag := (C.gpointer)(tag)

	c_new_events := (C.GIOCondition)(newEvents)

	C.g_source_modify_unix_fd((*C.GSource)(recv.native), c_tag, c_new_events)

	return
}

// QueryUnixFd is a wrapper around the C function g_source_query_unix_fd.
func (recv *Source) QueryUnixFd(tag uintptr) IOCondition {
	c_tag := (C.gpointer)(tag)

	retC := C.g_source_query_unix_fd((*C.GSource)(recv.native), c_tag)
	retGo := (IOCondition)(retC)

	return retGo
}

// RemoveUnixFd is a wrapper around the C function g_source_remove_unix_fd.
func (recv *Source) RemoveUnixFd(tag uintptr) {
	c_tag := (C.gpointer)(tag)

	C.g_source_remove_unix_fd((*C.GSource)(recv.native), c_tag)

	return
}

// SetReadyTime is a wrapper around the C function g_source_set_ready_time.
func (recv *Source) SetReadyTime(readyTime int64) {
	c_ready_time := (C.gint64)(readyTime)

	C.g_source_set_ready_time((*C.GSource)(recv.native), c_ready_time)

	return
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType
