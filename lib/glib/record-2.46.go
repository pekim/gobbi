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

// MoveToFront is a wrapper around the C function g_thread_pool_move_to_front.
func (recv *ThreadPool) MoveToFront(data uintptr) bool {
	c_data := (C.gpointer)(data)

	retC := C.g_thread_pool_move_to_front((*C.GThreadPool)(recv.native), c_data)
	retGo := retC == C.TRUE

	return retGo
}
