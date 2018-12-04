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

// DirMakeTmp is a wrapper around the C function g_dir_make_tmp.
func DirMakeTmp(tmpl string) (string, error) {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	var cThrowableError *C.GError

	retC := C.g_dir_make_tmp(c_tmpl, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Replace is a wrapper around the C function g_hash_table_iter_replace.
func (recv *HashTableIter) Replace(value uintptr) {
	c_value := (C.gpointer)(value)

	C.g_hash_table_iter_replace((*C.GHashTableIter)(recv.native), c_value)

	return
}

// Hmac is a wrapper around the C record GHmac.
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

// Equals compares this Hmac with another Hmac, and returns true if they represent the same GObject.
func (recv *Hmac) Equals(other *Hmac) bool {
	return other.ToC() == recv.ToC()
}

// HmacNew is a wrapper around the C function g_hmac_new.
func HmacNew(digestType ChecksumType, key []uint8) *Hmac {
	c_digest_type := (C.GChecksumType)(digestType)

	c_key := &key[0]

	c_key_len := (C.gsize)(len(key))

	retC := C.g_hmac_new(c_digest_type, (*C.guchar)(unsafe.Pointer(c_key)), c_key_len)
	retGo := HmacNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_hmac_copy.
func (recv *Hmac) Copy() *Hmac {
	retC := C.g_hmac_copy((*C.GHmac)(recv.native))
	retGo := HmacNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDigest is a wrapper around the C function g_hmac_get_digest.
func (recv *Hmac) GetDigest(buffer uint8, digestLen uint64) {
	c_buffer := (C.guint8)(buffer)

	c_digest_len := (C.gsize)(digestLen)

	C.g_hmac_get_digest((*C.GHmac)(recv.native), &c_buffer, &c_digest_len)

	return
}

// GetString is a wrapper around the C function g_hmac_get_string.
func (recv *Hmac) GetString() string {
	retC := C.g_hmac_get_string((*C.GHmac)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Ref is a wrapper around the C function g_hmac_ref.
func (recv *Hmac) Ref() *Hmac {
	retC := C.g_hmac_ref((*C.GHmac)(recv.native))
	retGo := HmacNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_hmac_unref.
func (recv *Hmac) Unref() {
	C.g_hmac_unref((*C.GHmac)(recv.native))

	return
}

// Update is a wrapper around the C function g_hmac_update.
func (recv *Hmac) Update(data []uint8) {
	c_data := &data[0]

	c_length := (C.gssize)(len(data))

	C.g_hmac_update((*C.GHmac)(recv.native), (*C.guchar)(unsafe.Pointer(c_data)), c_length)

	return
}

// Ref is a wrapper around the C function g_match_info_ref.
func (recv *MatchInfo) Ref() *MatchInfo {
	retC := C.g_match_info_ref((*C.GMatchInfo)(recv.native))
	retGo := MatchInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_match_info_unref.
func (recv *MatchInfo) Unref() {
	C.g_match_info_unref((*C.GMatchInfo)(recv.native))

	return
}

// RegexEscapeNul is a wrapper around the C function g_regex_escape_nul.
func RegexEscapeNul(string string, length int32) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gint)(length)

	retC := C.g_regex_escape_nul(c_string, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}
