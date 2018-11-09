// This is a generated file - DO NOT EDIT
// +build glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// PushFront is a wrapper around the C function g_async_queue_push_front.
func (recv *AsyncQueue) PushFront(item uintptr) {
	c_item := (C.gpointer)(item)

	C.g_async_queue_push_front((*C.GAsyncQueue)(recv.native), c_item)

	return
}

// PushFrontUnlocked is a wrapper around the C function g_async_queue_push_front_unlocked.
func (recv *AsyncQueue) PushFrontUnlocked(item uintptr) {
	c_item := (C.gpointer)(item)

	C.g_async_queue_push_front_unlocked((*C.GAsyncQueue)(recv.native), c_item)

	return
}

// Remove is a wrapper around the C function g_async_queue_remove.
func (recv *AsyncQueue) Remove(item uintptr) bool {
	c_item := (C.gpointer)(item)

	retC := C.g_async_queue_remove((*C.GAsyncQueue)(recv.native), c_item)
	retGo := retC == C.TRUE

	return retGo
}

// RemoveUnlocked is a wrapper around the C function g_async_queue_remove_unlocked.
func (recv *AsyncQueue) RemoveUnlocked(item uintptr) bool {
	c_item := (C.gpointer)(item)

	retC := C.g_async_queue_remove_unlocked((*C.GAsyncQueue)(recv.native), c_item)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_bytes_new_with_free_func : unsupported parameter free_func : no type generator for DestroyNotify (GDestroyNotify) for param free_func

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// MoveToFront is a wrapper around the C function g_thread_pool_move_to_front.
func (recv *ThreadPool) MoveToFront(data uintptr) bool {
	c_data := (C.gpointer)(data)

	retC := C.g_thread_pool_move_to_front((*C.GThreadPool)(recv.native), c_data)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant
