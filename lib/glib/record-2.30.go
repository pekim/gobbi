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

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

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

// Unsupported : g_hmac_update : unsupported parameter data : no type generator for guint8 (guchar) for array param data

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

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

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify (GDestroyNotify) for param destroy

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc (GThreadFunc) for param func

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType
