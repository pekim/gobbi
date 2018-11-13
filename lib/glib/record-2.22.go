// This is a generated file - DO NOT EDIT
// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Pops @context off the thread-default context stack (verifying that
// it was on the top of the stack).
/*

C function : g_main_context_pop_thread_default
*/
func (recv *MainContext) PopThreadDefault() {
	C.g_main_context_pop_thread_default((*C.GMainContext)(recv.native))

	return
}

// Acquires @context and sets it as the thread-default context for the
// current thread. This will cause certain asynchronous operations
// (such as most [gio][gio]-based I/O) which are
// started in this thread to run under @context and deliver their
// results to its main loop, rather than running under the global
// default context in the main thread. Note that calling this function
// changes the context returned by g_main_context_get_thread_default(),
// not the one returned by g_main_context_default(), so it does not affect
// the context used by functions like g_idle_add().
//
// Normally you would call this function shortly after creating a new
// thread, passing it a #GMainContext which will be run by a
// #GMainLoop in that thread, to set a new default context for all
// async operations in that thread. In this case you may not need to
// ever call g_main_context_pop_thread_default(), assuming you want the
// new #GMainContext to be the default for the whole lifecycle of the
// thread.
//
// If you don't have control over how the new thread was created (e.g.
// in the new thread isn't newly created, or if the thread life
// cycle is managed by a #GThreadPool), it is always suggested to wrap
// the logic that needs to use the new #GMainContext inside a
// g_main_context_push_thread_default() / g_main_context_pop_thread_default()
// pair, otherwise threads that are re-used will end up never explicitly
// releasing the #GMainContext reference they hold.
//
// In some cases you may want to schedule a single operation in a
// non-default context, or temporarily use a non-default context in
// the main thread. In that case, you can wrap the call to the
// asynchronous operation inside a
// g_main_context_push_thread_default() /
// g_main_context_pop_thread_default() pair, but it is up to you to
// ensure that no other asynchronous operations accidentally get
// started while the non-default context is active.
//
// Beware that libraries that predate this function may not correctly
// handle being used from a thread with a thread-default context. Eg,
// see g_file_supports_thread_contexts().
/*

C function : g_main_context_push_thread_default
*/
func (recv *MainContext) PushThreadDefault() {
	C.g_main_context_push_thread_default((*C.GMainContext)(recv.native))

	return
}

// Increments the reference count of @file by one.  It is safe to call
// this function from any thread.
/*

C function : g_mapped_file_ref
*/
func (recv *MappedFile) Ref() *MappedFile {
	retC := C.g_mapped_file_ref((*C.GMappedFile)(recv.native))
	retGo := MappedFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Increments the reference count of @tree by one.
//
// It is safe to call this function from any thread.
/*

C function : g_tree_ref
*/
func (recv *Tree) Ref() *Tree {
	retC := C.g_tree_ref((*C.GTree)(recv.native))
	retGo := TreeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decrements the reference count of @tree by one.
// If the reference count drops to 0, all keys and values will
// be destroyed (if destroy functions were specified) and all
// memory allocated by @tree will be released.
//
// It is safe to call this function from any thread.
/*

C function : g_tree_unref
*/
func (recv *Tree) Unref() {
	C.g_tree_unref((*C.GTree)(recv.native))

	return
}
