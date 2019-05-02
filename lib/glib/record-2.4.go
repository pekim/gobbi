// This is a generated file - DO NOT EDIT
// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"C"
	"unsafe"
)

// g_array_remove_range : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_hash_table_find : unsupported parameter predicate : no type generator for HRFunc (GHRFunc) for param predicate
// Unsupported : g_node_copy_deep : unsupported parameter copy_func : no type generator for CopyFunc (GCopyFunc) for param copy_func

// Equals compares this Once with another Once, and returns true if they represent the same GObject.
func (recv *Once) Equals(other *Once) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : g_once_impl : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Copy is a wrapper around the C function g_queue_copy.
func (recv *Queue) Copy() *Queue {
	retC := C.g_queue_copy((*C.GQueue)(recv.native))
	retGo := QueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DeleteLink is a wrapper around the C function g_queue_delete_link.
func (recv *Queue) DeleteLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_delete_link((*C.GQueue)(recv.native), c_link_)

	return
}

// Unsupported : g_queue_find : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_find_custom : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_foreach : unsupported parameter func : no type generator for Func (GFunc) for param func

// GetLength is a wrapper around the C function g_queue_get_length.
func (recv *Queue) GetLength() uint32 {
	retC := C.g_queue_get_length((*C.GQueue)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_queue_index : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_insert_after : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_queue_insert_before : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : g_queue_insert_sorted : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// LinkIndex is a wrapper around the C function g_queue_link_index.
func (recv *Queue) LinkIndex(link *List) int32 {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

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

// Unsupported : g_queue_peek_nth : no return generator

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

// Unsupported : g_queue_pop_nth : no return generator

// PopNthLink is a wrapper around the C function g_queue_pop_nth_link.
func (recv *Queue) PopNthLink(n uint32) *List {
	c_n := (C.guint)(n)

	retC := C.g_queue_pop_nth_link((*C.GQueue)(recv.native), c_n)
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_queue_push_nth : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// PushNthLink is a wrapper around the C function g_queue_push_nth_link.
func (recv *Queue) PushNthLink(n int32, link *List) {
	c_n := (C.gint)(n)

	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_nth_link((*C.GQueue)(recv.native), c_n, c_link_)

	return
}

// Unsupported : g_queue_remove : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Unsupported : g_queue_remove_all : unsupported parameter data : no type generator for gpointer (gconstpointer) for param data

// Reverse is a wrapper around the C function g_queue_reverse.
func (recv *Queue) Reverse() {
	C.g_queue_reverse((*C.GQueue)(recv.native))

	return
}

// Unsupported : g_queue_sort : unsupported parameter compare_func : no type generator for CompareDataFunc (GCompareDataFunc) for param compare_func

// Unlink is a wrapper around the C function g_queue_unlink.
func (recv *Queue) Unlink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_unlink((*C.GQueue)(recv.native), c_link_)

	return
}

// RandNewWithSeedArray is a wrapper around the C function g_rand_new_with_seed_array.
func RandNewWithSeedArray(seed uint32, seedLength uint32) *Rand {
	c_seed := (C.guint32)(seed)

	c_seed_length := (C.guint)(seedLength)

	retC := C.g_rand_new_with_seed_array(&c_seed, c_seed_length)
	retGo := RandNewFromC(unsafe.Pointer(retC))

	return retGo
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
func (recv *StringChunk) InsertLen(string_ string, len int64) string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_len := (C.gssize)(len)

	retC := C.g_string_chunk_insert_len((*C.GStringChunk)(recv.native), c_string, c_len)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Continue is a wrapper around the C function g_timer_continue.
func (recv *Timer) Continue() {
	C.g_timer_continue((*C.GTimer)(recv.native))

	return
}
