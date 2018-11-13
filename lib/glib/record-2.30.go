// This is a generated file - DO NOT EDIT
// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Replaces the value currently pointed to by the iterator
// from its associated #GHashTable. Can only be called after
// g_hash_table_iter_next() returned %TRUE.
//
// If you supplied a @value_destroy_func when creating the
// #GHashTable, the old value is freed using that function.
/*

C function : g_hash_table_iter_replace
*/
func (recv *HashTableIter) Replace(value uintptr) {
	c_value := (C.gpointer)(value)

	C.g_hash_table_iter_replace((*C.GHashTableIter)(recv.native), c_value)

	return
}

// An opaque structure representing a HMAC operation.
// To create a new GHmac, use g_hmac_new(). To free
// a GHmac, use g_hmac_unref().
/*

C record/class : GHmac
*/
type Hmac struct {
	native *C.GHmac
}

func HmacNewFromC(u unsafe.Pointer) *Hmac {
	c := (*C.GHmac)(u)
	if c == nil {
		return nil
	}

	g := &Hmac{native: c}

	return g
}

func (recv *Hmac) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Copies a #GHmac. If @hmac has been closed, by calling
// g_hmac_get_string() or g_hmac_get_digest(), the copied
// HMAC will be closed as well.
/*

C function : g_hmac_copy
*/
func (recv *Hmac) Copy() *Hmac {
	retC := C.g_hmac_copy((*C.GHmac)(recv.native))
	retGo := HmacNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the digest from @checksum as a raw binary array and places it
// into @buffer. The size of the digest depends on the type of checksum.
//
// Once this function has been called, the #GHmac is closed and can
// no longer be updated with g_checksum_update().
/*

C function : g_hmac_get_digest
*/
func (recv *Hmac) GetDigest(buffer uint8, digestLen uint64) {
	c_buffer := (C.guint8)(buffer)

	c_digest_len := (C.gsize)(digestLen)

	C.g_hmac_get_digest((*C.GHmac)(recv.native), &c_buffer, &c_digest_len)

	return
}

// Gets the HMAC as an hexadecimal string.
//
// Once this function has been called the #GHmac can no longer be
// updated with g_hmac_update().
//
// The hexadecimal characters will be lower case.
/*

C function : g_hmac_get_string
*/
func (recv *Hmac) GetString() string {
	retC := C.g_hmac_get_string((*C.GHmac)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Atomically increments the reference count of @hmac by one.
//
// This function is MT-safe and may be called from any thread.
/*

C function : g_hmac_ref
*/
func (recv *Hmac) Ref() *Hmac {
	retC := C.g_hmac_ref((*C.GHmac)(recv.native))
	retGo := HmacNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Atomically decrements the reference count of @hmac by one.
//
// If the reference count drops to 0, all keys and values will be
// destroyed, and all memory allocated by the hash table is released.
// This function is MT-safe and may be called from any thread.
// Frees the memory allocated for @hmac.
/*

C function : g_hmac_unref
*/
func (recv *Hmac) Unref() {
	C.g_hmac_unref((*C.GHmac)(recv.native))

	return
}

// Feeds @data into an existing #GHmac.
//
// The HMAC must still be open, that is g_hmac_get_string() or
// g_hmac_get_digest() must not have been called on @hmac.
/*

C function : g_hmac_update
*/
func (recv *Hmac) Update(data []uint8) {
	c_data := &data[0]

	c_length := (C.gssize)(len(data))

	C.g_hmac_update((*C.GHmac)(recv.native), (*C.guchar)(unsafe.Pointer(c_data)), c_length)

	return
}

// Increases reference count of @match_info by 1.
/*

C function : g_match_info_ref
*/
func (recv *MatchInfo) Ref() *MatchInfo {
	retC := C.g_match_info_ref((*C.GMatchInfo)(recv.native))
	retGo := MatchInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decreases reference count of @match_info by 1. When reference count drops
// to zero, it frees all the memory associated with the match_info structure.
/*

C function : g_match_info_unref
*/
func (recv *MatchInfo) Unref() {
	C.g_match_info_unref((*C.GMatchInfo)(recv.native))

	return
}
