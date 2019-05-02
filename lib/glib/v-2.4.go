// Code generated - DO NOT EDIT.
// +build glib_2.4 glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import (
	"fmt"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
/*

	static gchar* _g_markup_printf_escaped(const char* format) {
		return g_markup_printf_escaped(format);
    }
*/
import "C"

const GINT16_MODIFIER string = C.G_GINT16_MODIFIER
const GINT32_MODIFIER string = C.G_GINT32_MODIFIER
const GINT64_MODIFIER string = C.G_GINT64_MODIFIER
const MAXINT16 int16 = C.G_MAXINT16
const MAXINT32 int32 = C.G_MAXINT32
const MAXINT8 int8 = C.G_MAXINT8
const MAXUINT16 uint16 = C.G_MAXUINT16
const MAXUINT32 uint32 = C.G_MAXUINT32
const MAXUINT8 uint8 = C.G_MAXUINT8
const MININT16 int16 = C.G_MININT16
const MININT32 int32 = C.G_MININT32
const MININT8 int8 = C.G_MININT8

type OnceStatus C.GOnceStatus

const (
	ONCE_STATUS_NOTCALLED OnceStatus = 0
	ONCE_STATUS_PROGRESS  OnceStatus = 1
	ONCE_STATUS_READY     OnceStatus = 2
)

// Blacklisted : g_atomic_int_add

// Blacklisted : g_atomic_int_compare_and_exchange

// Blacklisted : g_atomic_int_dec_and_test

// Blacklisted : g_atomic_int_exchange_and_add

// Blacklisted : g_atomic_int_get

// Blacklisted : g_atomic_int_inc

// Blacklisted : g_atomic_int_set

// Unsupported : g_atomic_pointer_compare_and_exchange : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_get : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_set : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_child_watch_add : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// Unsupported : g_child_watch_add_full : unsupported parameter function : no type generator for ChildWatchFunc (GChildWatchFunc) for param function

// ChildWatchSourceNew is a wrapper around the C function g_child_watch_source_new.
func ChildWatchSourceNew(pid Pid) *Source {
	c_pid := (C.GPid)(pid)

	retC := C.g_child_watch_source_new(c_pid)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileReadLink is a wrapper around the C function g_file_read_link.
func FileReadLink(filename string) (string, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.g_file_read_link(c_filename, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// MarkupPrintfEscaped is a wrapper around the C function g_markup_printf_escaped.
func MarkupPrintfEscaped(format string, args ...interface{}) string {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_markup_printf_escaped(c_format)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_markup_vprintf_escaped : unsupported parameter args : no type generator for va_list (va_list) for param args

// Setenv is a wrapper around the C function g_setenv.
func Setenv(variable string, value string, overwrite bool) bool {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_overwrite :=
		boolToGboolean(overwrite)

	retC := C.g_setenv(c_variable, c_value, c_overwrite)
	retGo := retC == C.TRUE

	return retGo
}

// StripContext is a wrapper around the C function g_strip_context.
func StripContext(msgid string, msgval string) string {
	c_msgid := C.CString(msgid)
	defer C.free(unsafe.Pointer(c_msgid))

	c_msgval := C.CString(msgval)
	defer C.free(unsafe.Pointer(c_msgval))

	retC := C.g_strip_context(c_msgid, c_msgval)
	retGo := C.GoString(retC)

	return retGo
}

// StrsplitSet is a wrapper around the C function g_strsplit_set.
func StrsplitSet(string_ string, delimiters string, maxTokens int32) []string {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_delimiters := C.CString(delimiters)
	defer C.free(unsafe.Pointer(c_delimiters))

	c_max_tokens := (C.gint)(maxTokens)

	retC := C.g_strsplit_set(c_string, c_delimiters, c_max_tokens)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// UnicharGetMirrorChar is a wrapper around the C function g_unichar_get_mirror_char.
func UnicharGetMirrorChar(ch rune, mirroredCh rune) bool {
	c_ch := (C.gunichar)(ch)

	c_mirrored_ch := (C.gunichar)(mirroredCh)

	retC := C.g_unichar_get_mirror_char(c_ch, &c_mirrored_ch)
	retGo := retC == C.TRUE

	return retGo
}

// Unsetenv is a wrapper around the C function g_unsetenv.
func Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_unsetenv(c_variable)

	return
}

// Unsupported : g_vasprintf : unsupported parameter string : in string with indirection level of 2

// g_array_remove_range : unsupported parameter array : no type generator for gpointer (gpointer) for array param array
// g_hash_table_find : unsupported parameter predicate : no type generator for HRFunc (GHRFunc) for param predicate
// Unsupported : g_node_copy_deep : unsupported parameter copy_func : no type generator for CopyFunc (GCopyFunc) for param copy_func

// Once is a wrapper around the C record GOnce.
type Once struct {
	native *C.GOnce
	Status OnceStatus
	// retval : no type generator for gpointer, volatile gpointer
}

func OnceNewFromC(u unsafe.Pointer) *Once {
	c := (*C.GOnce)(u)
	if c == nil {
		return nil
	}

	g := &Once{
		Status: (OnceStatus)(c.status),
		native: c,
	}

	return g
}

func (recv *Once) ToC() unsafe.Pointer {
	recv.native.status =
		(C.GOnceStatus)(recv.Status)

	return (unsafe.Pointer)(recv.native)
}

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
