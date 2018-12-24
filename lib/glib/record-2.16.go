// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

// g_async_queue_new_full : unsupported parameter item_free_func : no type generator for DestroyNotify (GDestroyNotify) for param item_free_func
// Checksum is a wrapper around the C record GChecksum.
type Checksum struct {
	native *C.GChecksum
}

func ChecksumNewFromC(u unsafe.Pointer) *Checksum {
	c := (*C.GChecksum)(u)
	if c == nil {
		return nil
	}

	g := &Checksum{native: c}

	return g
}

func (recv *Checksum) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Checksum with another Checksum, and returns true if they represent the same GObject.
func (recv *Checksum) Equals(other *Checksum) bool {
	return other.ToC() == recv.ToC()
}

// ChecksumNew is a wrapper around the C function g_checksum_new.
func ChecksumNew(checksumType ChecksumType) *Checksum {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_new(c_checksum_type)
	retGo := ChecksumNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ChecksumTypeGetLength is a wrapper around the C function g_checksum_type_get_length.
func ChecksumTypeGetLength(checksumType ChecksumType) int64 {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_type_get_length(c_checksum_type)
	retGo := (int64)(retC)

	return retGo
}

// Copy is a wrapper around the C function g_checksum_copy.
func (recv *Checksum) Copy() *Checksum {
	retC := C.g_checksum_copy((*C.GChecksum)(recv.native))
	retGo := ChecksumNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_checksum_free.
func (recv *Checksum) Free() {
	C.g_checksum_free((*C.GChecksum)(recv.native))

	return
}

// GetDigest is a wrapper around the C function g_checksum_get_digest.
func (recv *Checksum) GetDigest(buffer uint8, digestLen uint64) {
	c_buffer := (C.guint8)(buffer)

	c_digest_len := (C.gsize)(digestLen)

	C.g_checksum_get_digest((*C.GChecksum)(recv.native), &c_buffer, &c_digest_len)

	return
}

// GetString is a wrapper around the C function g_checksum_get_string.
func (recv *Checksum) GetString() string {
	retC := C.g_checksum_get_string((*C.GChecksum)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Update is a wrapper around the C function g_checksum_update.
func (recv *Checksum) Update(data []uint8) {
	c_data := &data[0]

	c_length := (C.gssize)(len(data))

	C.g_checksum_update((*C.GChecksum)(recv.native), (*C.guchar)(unsafe.Pointer(c_data)), c_length)

	return
}

// GetHashTable is a wrapper around the C function g_hash_table_iter_get_hash_table.
func (recv *HashTableIter) GetHashTable() *HashTable {
	retC := C.g_hash_table_iter_get_hash_table((*C.GHashTableIter)(recv.native))
	retGo := HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Init is a wrapper around the C function g_hash_table_iter_init.
func (recv *HashTableIter) Init(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_iter_init((*C.GHashTableIter)(recv.native), c_hash_table)

	return
}

// Next is a wrapper around the C function g_hash_table_iter_next.
func (recv *HashTableIter) Next() (bool, uintptr, uintptr) {
	var c_key C.gpointer

	var c_value C.gpointer

	retC := C.g_hash_table_iter_next((*C.GHashTableIter)(recv.native), &c_key, &c_value)
	retGo := retC == C.TRUE

	key := (uintptr)(unsafe.Pointer(&c_key))

	value := (uintptr)(unsafe.Pointer(&c_value))

	return retGo, key, value
}

// Remove is a wrapper around the C function g_hash_table_iter_remove.
func (recv *HashTableIter) Remove() {
	C.g_hash_table_iter_remove((*C.GHashTableIter)(recv.native))

	return
}

// Steal is a wrapper around the C function g_hash_table_iter_steal.
func (recv *HashTableIter) Steal() {
	C.g_hash_table_iter_steal((*C.GHashTableIter)(recv.native))

	return
}

// GetElementStack is a wrapper around the C function g_markup_parse_context_get_element_stack.
func (recv *MarkupParseContext) GetElementStack() *SList {
	retC := C.g_markup_parse_context_get_element_stack((*C.GMarkupParseContext)(recv.native))
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendUriEscaped is a wrapper around the C function g_string_append_uri_escaped.
func (recv *String) AppendUriEscaped(unescaped string, reservedCharsAllowed string, allowUtf8 bool) *String {
	c_unescaped := C.CString(unescaped)
	defer C.free(unsafe.Pointer(c_unescaped))

	c_reserved_chars_allowed := C.CString(reservedCharsAllowed)
	defer C.free(unsafe.Pointer(c_reserved_chars_allowed))

	c_allow_utf8 :=
		boolToGboolean(allowUtf8)

	retC := C.g_string_append_uri_escaped((*C.GString)(recv.native), c_unescaped, c_reserved_chars_allowed, c_allow_utf8)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add is a wrapper around the C function g_test_suite_add.
func (recv *TestSuite) Add(testCase *TestCase) {
	c_test_case := (*C.GTestCase)(C.NULL)
	if testCase != nil {
		c_test_case = (*C.GTestCase)(testCase.ToC())
	}

	C.g_test_suite_add((*C.GTestSuite)(recv.native), c_test_case)

	return
}

// AddSuite is a wrapper around the C function g_test_suite_add_suite.
func (recv *TestSuite) AddSuite(nestedsuite *TestSuite) {
	c_nestedsuite := (*C.GTestSuite)(C.NULL)
	if nestedsuite != nil {
		c_nestedsuite = (*C.GTestSuite)(nestedsuite.ToC())
	}

	C.g_test_suite_add_suite((*C.GTestSuite)(recv.native), c_nestedsuite)

	return
}
