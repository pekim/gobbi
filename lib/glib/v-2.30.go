// Code generated - DO NOT EDIT.
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

// AtomicIntAnd is a wrapper around the C function g_atomic_int_and.
func AtomicIntAnd(atomic uint32, val uint32) uint32 {
	c_atomic := (C.guint)(atomic)

	c_val := (C.guint)(val)

	retC := C.g_atomic_int_and(&c_atomic, c_val)
	retGo := (uint32)(retC)

	return retGo
}

// AtomicIntOr is a wrapper around the C function g_atomic_int_or.
func AtomicIntOr(atomic uint32, val uint32) uint32 {
	c_atomic := (C.guint)(atomic)

	c_val := (C.guint)(val)

	retC := C.g_atomic_int_or(&c_atomic, c_val)
	retGo := (uint32)(retC)

	return retGo
}

// AtomicIntXor is a wrapper around the C function g_atomic_int_xor.
func AtomicIntXor(atomic uint32, val uint32) uint32 {
	c_atomic := (C.guint)(atomic)

	c_val := (C.guint)(val)

	retC := C.g_atomic_int_xor(&c_atomic, c_val)
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_atomic_pointer_add : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_and : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_or : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// Unsupported : g_atomic_pointer_xor : unsupported parameter atomic : no type generator for gpointer (void*) for param atomic

// ComputeHmacForData is a wrapper around the C function g_compute_hmac_for_data.
func ComputeHmacForData(digestType ChecksumType, key []uint8, data []uint8) string {
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

	c_data_array := make([]C.guchar, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guchar)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	retC := C.g_compute_hmac_for_data(c_digest_type, c_key, c_key_len, c_data, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ComputeHmacForString is a wrapper around the C function g_compute_hmac_for_string.
func ComputeHmacForString(digestType ChecksumType, key []uint8, str string) string {
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

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_length := (C.gssize)(len(str))

	retC := C.g_compute_hmac_for_string(c_digest_type, c_key, c_key_len, c_str, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FormatSize is a wrapper around the C function g_format_size.
func FormatSize(size uint64) string {
	c_size := (C.guint64)(size)

	retC := C.g_format_size(c_size)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FormatSizeFull is a wrapper around the C function g_format_size_full.
func FormatSizeFull(size uint64, flags FormatSizeFlags) string {
	c_size := (C.guint64)(size)

	c_flags := (C.GFormatSizeFlags)(flags)

	retC := C.g_format_size_full(c_size, c_flags)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Mkdtemp is a wrapper around the C function g_mkdtemp.
func Mkdtemp(tmpl string) string {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	retC := C.g_mkdtemp(c_tmpl)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// MkdtempFull is a wrapper around the C function g_mkdtemp_full.
func MkdtempFull(tmpl string, mode int32) string {
	c_tmpl := C.CString(tmpl)
	defer C.free(unsafe.Pointer(c_tmpl))

	c_mode := (C.gint)(mode)

	retC := C.g_mkdtemp_full(c_tmpl, c_mode)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_pointer_bit_lock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_trylock : unsupported parameter address : no type generator for gpointer (void*) for param address

// Unsupported : g_pointer_bit_unlock : unsupported parameter address : no type generator for gpointer (void*) for param address

// TestFail is a wrapper around the C function g_test_fail.
func TestFail() {
	C.g_test_fail()

	return
}

// UnicharCompose is a wrapper around the C function g_unichar_compose.
func UnicharCompose(a rune, b rune, ch rune) bool {
	c_a := (C.gunichar)(a)

	c_b := (C.gunichar)(b)

	c_ch := (C.gunichar)(ch)

	retC := C.g_unichar_compose(c_a, c_b, &c_ch)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharDecompose is a wrapper around the C function g_unichar_decompose.
func UnicharDecompose(ch rune, a rune, b rune) bool {
	c_ch := (C.gunichar)(ch)

	c_a := (C.gunichar)(a)

	c_b := (C.gunichar)(b)

	retC := C.g_unichar_decompose(c_ch, &c_a, &c_b)
	retGo := retC == C.TRUE

	return retGo
}

// UnicharFullyDecompose is a wrapper around the C function g_unichar_fully_decompose.
func UnicharFullyDecompose(ch rune, compat bool, result rune, resultLen uint64) uint64 {
	c_ch := (C.gunichar)(ch)

	c_compat :=
		boolToGboolean(compat)

	c_result := (C.gunichar)(result)

	c_result_len := (C.gsize)(resultLen)

	retC := C.g_unichar_fully_decompose(c_ch, c_compat, &c_result, c_result_len)
	retGo := (uint64)(retC)

	return retGo
}

// UnicodeScriptFromIso15924 is a wrapper around the C function g_unicode_script_from_iso15924.
func UnicodeScriptFromIso15924(iso15924 uint32) UnicodeScript {
	c_iso15924 := (C.guint32)(iso15924)

	retC := C.g_unicode_script_from_iso15924(c_iso15924)
	retGo := (UnicodeScript)(retC)

	return retGo
}

// UnicodeScriptToIso15924 is a wrapper around the C function g_unicode_script_to_iso15924.
func UnicodeScriptToIso15924(script UnicodeScript) uint32 {
	c_script := (C.GUnicodeScript)(script)

	retC := C.g_unicode_script_to_iso15924(c_script)
	retGo := (uint32)(retC)

	return retGo
}

// UnixOpenPipe is a wrapper around the C function g_unix_open_pipe.
func UnixOpenPipe(fds int32, flags int32) (bool, error) {
	c_fds := (C.gint)(fds)

	c_flags := (C.gint)(flags)

	var cThrowableError *C.GError

	retC := C.g_unix_open_pipe(&c_fds, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnixSetFdNonblocking is a wrapper around the C function g_unix_set_fd_nonblocking.
func UnixSetFdNonblocking(fd int32, nonblock bool) (bool, error) {
	c_fd := (C.gint)(fd)

	c_nonblock :=
		boolToGboolean(nonblock)

	var cThrowableError *C.GError

	retC := C.g_unix_set_fd_nonblocking(c_fd, c_nonblock, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_unix_signal_add : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// Unsupported : g_unix_signal_add_full : unsupported parameter handler : no type generator for SourceFunc (GSourceFunc) for param handler

// UnixSignalSourceNew is a wrapper around the C function g_unix_signal_source_new.
func UnixSignalSourceNew(signum int32) *Source {
	c_signum := (C.gint)(signum)

	retC := C.g_unix_signal_source_new(c_signum)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Utf8Substring is a wrapper around the C function g_utf8_substring.
func Utf8Substring(str string, startPos int64, endPos int64) string {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_start_pos := (C.glong)(startPos)

	c_end_pos := (C.glong)(endPos)

	retC := C.g_utf8_substring(c_str, c_start_pos, c_end_pos)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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
