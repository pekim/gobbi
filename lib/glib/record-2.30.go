// This is a generated file - DO NOT EDIT
// +build glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_hash_table_iter_replace : unsupported parameter value : no type generator for gpointer (gpointer) for param value

// Equals compares this Hmac with another Hmac, and returns true if they represent the same GObject.
func (recv *Hmac) Equals(other *Hmac) bool {
	return other.ToC() == recv.ToC()
}

// HmacNew is a wrapper around the C function g_hmac_new.
func HmacNew(digestType ChecksumType, key []uint8) *Hmac {
	c_digest_type := (C.GChecksumType)(digestType)

	c_key_array := make([]C.guchar, len(key)+1, len(key)+1)
	for i, item := range key {
		c := (C.guchar)(item)
		c_key_array[i] = c
	}
	c_key_array[len(key)] = 0
	c_key_arrayPtr := &c_key_array[0]
	c_key := (*C.guchar)(unsafe.Pointer(c_key_arrayPtr))

	c_key_len := (C.gsize)(len(key))

	retC := C.g_hmac_new(c_digest_type, c_key, c_key_len)
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
	c_data_array := make([]C.guchar, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guchar)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gssize)(len(data))

	C.g_hmac_update((*C.GHmac)(recv.native), c_data, c_length)

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

// VariantNewObjv is a wrapper around the C function g_variant_new_objv.
func VariantNewObjv(strv []string) *Variant {
	c_strv_array := make([]*C.gchar, len(strv)+1, len(strv)+1)
	for i, item := range strv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_strv_array[i] = c
	}
	c_strv_array[len(strv)] = nil
	c_strv_arrayPtr := &c_strv_array[0]
	c_strv := (**C.gchar)(unsafe.Pointer(c_strv_arrayPtr))

	c_length := (C.gssize)(len(strv))

	retC := C.g_variant_new_objv(c_strv, c_length)
	retGo := VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DupObjv is a wrapper around the C function g_variant_dup_objv.
func (recv *Variant) DupObjv() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_dup_objv((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// GetObjv is a wrapper around the C function g_variant_get_objv.
func (recv *Variant) GetObjv() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_variant_get_objv((*C.GVariant)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	length := (uint64)(c_length)

	return retGo, length
}
