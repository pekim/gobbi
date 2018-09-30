// This is a generated file - DO NOT EDIT
// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_bytes_new : unsupported parameter data : no param type

// Unsupported : g_bytes_new_static : unsupported parameter data : no param type

// Unsupported : g_bytes_new_take : unsupported parameter data : no param type

// Unsupported : g_bytes_new_with_free_func : unsupported parameter data : no param type

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify, GDestroyNotify

// Unsupported : g_node_copy_deep : unsupported parameter copy_func : no type generator for CopyFunc, GCopyFunc

// Once is a wrapper around the C record GOnce.
type Once struct {
	native *C.GOnce
	Status OnceStatus
	Retval uintptr
}

func OnceNewFromC(u unsafe.Pointer) *Once {
	c := (*C.GOnce)(u)
	if c == nil {
		return nil
	}

	g := &Once{
		Retval: (uintptr)(c.retval),
		Status: (OnceStatus)(c.status),
		native: c,
	}

	return g
}

func (recv *Once) ToC() unsafe.Pointer {
	recv.native.status =
		(C.GOnceStatus)(recv.Status)
	recv.native.retval =
		(C.gpointer)(recv.Retval)

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_once_impl : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify, GDestroyNotify

// Copy is a wrapper around the C function g_queue_copy.
func (recv *Queue) Copy() *Queue {
	retC := C.g_queue_copy((*C.GQueue)(recv.native))
	retGo := QueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DeleteLink is a wrapper around the C function g_queue_delete_link.
func (recv *Queue) DeleteLink(link *List) {
	c_link_ := (*C.GList)(link.ToC())

	C.g_queue_delete_link((*C.GQueue)(recv.native), c_link_)

	return
}

// Find is a wrapper around the C function g_queue_find.
func (recv *Queue) Find(data uintptr) *List {
	c_data := (C.gconstpointer)(data)

	retC := C.g_queue_find((*C.GQueue)(recv.native), c_data)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_find_custom : unsupported parameter func : no type generator for CompareFunc, GCompareFunc

// Unsupported : g_queue_foreach : unsupported parameter func : no type generator for Func, GFunc

// GetLength is a wrapper around the C function g_queue_get_length.
func (recv *Queue) GetLength() uint32 {
	retC := C.g_queue_get_length((*C.GQueue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Index is a wrapper around the C function g_queue_index.
func (recv *Queue) Index(data uintptr) int32 {
	c_data := (C.gconstpointer)(data)

	retC := C.g_queue_index((*C.GQueue)(recv.native), c_data)
	retGo := (int32)(retC)

	return retGo
}

// InsertAfter is a wrapper around the C function g_queue_insert_after.
func (recv *Queue) InsertAfter(sibling *List, data uintptr) {
	c_sibling := (*C.GList)(sibling.ToC())

	c_data := (C.gpointer)(data)

	C.g_queue_insert_after((*C.GQueue)(recv.native), c_sibling, c_data)

	return
}

// InsertBefore is a wrapper around the C function g_queue_insert_before.
func (recv *Queue) InsertBefore(sibling *List, data uintptr) {
	c_sibling := (*C.GList)(sibling.ToC())

	c_data := (C.gpointer)(data)

	C.g_queue_insert_before((*C.GQueue)(recv.native), c_sibling, c_data)

	return
}

// Unsupported : g_queue_insert_sorted : unsupported parameter func : no type generator for CompareDataFunc, GCompareDataFunc

// LinkIndex is a wrapper around the C function g_queue_link_index.
func (recv *Queue) LinkIndex(link *List) int32 {
	c_link_ := (*C.GList)(link.ToC())

	retC := C.g_queue_link_index((*C.GQueue)(recv.native), c_link_)
	retGo := (int32)(retC)

	return retGo
}

// PeekHeadLink is a wrapper around the C function g_queue_peek_head_link.
func (recv *Queue) PeekHeadLink() *List {
	retC := C.g_queue_peek_head_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PeekNth is a wrapper around the C function g_queue_peek_nth.
func (recv *Queue) PeekNth(n uint32) uintptr {
	c_n := (C.guint)(n)

	retC := C.g_queue_peek_nth((*C.GQueue)(recv.native), c_n)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PeekNthLink is a wrapper around the C function g_queue_peek_nth_link.
func (recv *Queue) PeekNthLink(n uint32) *List {
	c_n := (C.guint)(n)

	retC := C.g_queue_peek_nth_link((*C.GQueue)(recv.native), c_n)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PeekTailLink is a wrapper around the C function g_queue_peek_tail_link.
func (recv *Queue) PeekTailLink() *List {
	retC := C.g_queue_peek_tail_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PopNth is a wrapper around the C function g_queue_pop_nth.
func (recv *Queue) PopNth(n uint32) uintptr {
	c_n := (C.guint)(n)

	retC := C.g_queue_pop_nth((*C.GQueue)(recv.native), c_n)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// PopNthLink is a wrapper around the C function g_queue_pop_nth_link.
func (recv *Queue) PopNthLink(n uint32) *List {
	c_n := (C.guint)(n)

	retC := C.g_queue_pop_nth_link((*C.GQueue)(recv.native), c_n)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// PushNth is a wrapper around the C function g_queue_push_nth.
func (recv *Queue) PushNth(data uintptr, n int32) {
	c_data := (C.gpointer)(data)

	c_n := (C.gint)(n)

	C.g_queue_push_nth((*C.GQueue)(recv.native), c_data, c_n)

	return
}

// PushNthLink is a wrapper around the C function g_queue_push_nth_link.
func (recv *Queue) PushNthLink(n int32, link *List) {
	c_n := (C.gint)(n)

	c_link_ := (*C.GList)(link.ToC())

	C.g_queue_push_nth_link((*C.GQueue)(recv.native), c_n, c_link_)

	return
}

// Remove is a wrapper around the C function g_queue_remove.
func (recv *Queue) Remove(data uintptr) bool {
	c_data := (C.gconstpointer)(data)

	retC := C.g_queue_remove((*C.GQueue)(recv.native), c_data)
	retGo := retC == C.TRUE

	return retGo
}

// RemoveAll is a wrapper around the C function g_queue_remove_all.
func (recv *Queue) RemoveAll(data uintptr) uint32 {
	c_data := (C.gconstpointer)(data)

	retC := C.g_queue_remove_all((*C.GQueue)(recv.native), c_data)
	retGo := (uint32)(retC)

	return retGo
}

// Reverse is a wrapper around the C function g_queue_reverse.
func (recv *Queue) Reverse() {
	C.g_queue_reverse((*C.GQueue)(recv.native))

	return
}

// Unsupported : g_queue_sort : unsupported parameter compare_func : no type generator for CompareDataFunc, GCompareDataFunc

// Unlink is a wrapper around the C function g_queue_unlink.
func (recv *Queue) Unlink(link *List) {
	c_link_ := (*C.GList)(link.ToC())

	C.g_queue_unlink((*C.GQueue)(recv.native), c_link_)

	return
}

// Copy is a wrapper around the C function g_rand_copy.
func (recv *Rand) Copy() *Rand {
	retC := C.g_rand_copy((*C.GRand)(recv.native))
	retGo := RandNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetSeedArray is a wrapper around the C function g_rand_set_seed_array.
func (recv *Rand) SetSeedArray(seed uint32, seedLength uint32) {
	c_seed := (C.guint32)(seed)

	c_seed_length := (C.guint)(seedLength)

	C.g_rand_set_seed_array((*C.GRand)(recv.native), &c_seed, c_seed_length)

	return
}

// InsertLen is a wrapper around the C function g_string_chunk_insert_len.
func (recv *StringChunk) InsertLen(string string, len int64) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_len := (C.gssize)(len)

	retC := C.g_string_chunk_insert_len((*C.GStringChunk)(recv.native), c_string, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Continue is a wrapper around the C function g_timer_continue.
func (recv *Timer) Continue() {
	C.g_timer_continue((*C.GTimer)(recv.native))

	return
}

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// Unsupported : g_variant_dict_new : unsupported parameter from_asv : Blacklisted record : GVariant
