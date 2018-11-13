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

// Increases the reference count of @context.
/*

C function

g_markup_parse_context_ref
*/
func (recv *MarkupParseContext) Ref() *MarkupParseContext {
	retC := C.g_markup_parse_context_ref((*C.GMarkupParseContext)(recv.native))
	retGo := MarkupParseContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decreases the reference count of @context.  When its reference count
// drops to 0, it is freed.
/*

C function

g_markup_parse_context_unref
*/
func (recv *MarkupParseContext) Unref() {
	C.g_markup_parse_context_unref((*C.GMarkupParseContext)(recv.native))

	return
}

// Monitors @fd for the IO events in @events.
//
// The tag returned by this function can be used to remove or modify the
// monitoring of the fd using g_source_remove_unix_fd() or
// g_source_modify_unix_fd().
//
// It is not necessary to remove the fd before destroying the source; it
// will be cleaned up automatically.
//
// This API is only intended to be used by implementations of #GSource.
// Do not call this API on a #GSource that you did not create.
//
// As the name suggests, this function is not available on Windows.
/*

C function

g_source_add_unix_fd
*/
func (recv *Source) AddUnixFd(fd int32, events IOCondition) uintptr {
	c_fd := (C.gint)(fd)

	c_events := (C.GIOCondition)(events)

	retC := C.g_source_add_unix_fd((*C.GSource)(recv.native), c_fd, c_events)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Updates the event mask to watch for the fd identified by @tag.
//
// @tag is the tag returned from g_source_add_unix_fd().
//
// If you want to remove a fd, don't set its event mask to zero.
// Instead, call g_source_remove_unix_fd().
//
// This API is only intended to be used by implementations of #GSource.
// Do not call this API on a #GSource that you did not create.
//
// As the name suggests, this function is not available on Windows.
/*

C function

g_source_modify_unix_fd
*/
func (recv *Source) ModifyUnixFd(tag uintptr, newEvents IOCondition) {
	c_tag := (C.gpointer)(tag)

	c_new_events := (C.GIOCondition)(newEvents)

	C.g_source_modify_unix_fd((*C.GSource)(recv.native), c_tag, c_new_events)

	return
}

// Queries the events reported for the fd corresponding to @tag on
// @source during the last poll.
//
// The return value of this function is only defined when the function
// is called from the check or dispatch functions for @source.
//
// This API is only intended to be used by implementations of #GSource.
// Do not call this API on a #GSource that you did not create.
//
// As the name suggests, this function is not available on Windows.
/*

C function

g_source_query_unix_fd
*/
func (recv *Source) QueryUnixFd(tag uintptr) IOCondition {
	c_tag := (C.gpointer)(tag)

	retC := C.g_source_query_unix_fd((*C.GSource)(recv.native), c_tag)
	retGo := (IOCondition)(retC)

	return retGo
}

// Reverses the effect of a previous call to g_source_add_unix_fd().
//
// You only need to call this if you want to remove an fd from being
// watched while keeping the same source around.  In the normal case you
// will just want to destroy the source.
//
// This API is only intended to be used by implementations of #GSource.
// Do not call this API on a #GSource that you did not create.
//
// As the name suggests, this function is not available on Windows.
/*

C function

g_source_remove_unix_fd
*/
func (recv *Source) RemoveUnixFd(tag uintptr) {
	c_tag := (C.gpointer)(tag)

	C.g_source_remove_unix_fd((*C.GSource)(recv.native), c_tag)

	return
}

// Sets a #GSource to be dispatched when the given monotonic time is
// reached (or passed).  If the monotonic time is in the past (as it
// always will be if @ready_time is 0) then the source will be
// dispatched immediately.
//
// If @ready_time is -1 then the source is never woken up on the basis
// of the passage of time.
//
// Dispatching the source does not reset the ready time.  You should do
// so yourself, from the source dispatch function.
//
// Note that if you have a pair of sources where the ready time of one
// suggests that it will be delivered first but the priority for the
// other suggests that it would be delivered first, and the ready time
// for both sources is reached during the same main context iteration,
// then the order of dispatch is undefined.
//
// It is a no-op to call this function on a #GSource which has already been
// destroyed with g_source_destroy().
//
// This API is only intended to be used by implementations of #GSource.
// Do not call this API on a #GSource that you did not create.
/*

C function

g_source_set_ready_time
*/
func (recv *Source) SetReadyTime(readyTime int64) {
	c_ready_time := (C.gint64)(readyTime)

	C.g_source_set_ready_time((*C.GSource)(recv.native), c_ready_time)

	return
}
