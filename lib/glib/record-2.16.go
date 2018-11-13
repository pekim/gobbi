// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// An opaque structure representing a checksumming operation.
// To create a new GChecksum, use g_checksum_new(). To free
// a GChecksum, use g_checksum_free().
/*

C type

GChecksum
*/
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

// Creates a new #GChecksum, using the checksum algorithm @checksum_type.
// If the @checksum_type is not known, %NULL is returned.
// A #GChecksum can be used to compute the checksum, or digest, of an
// arbitrary binary blob, using different hashing algorithms.
//
// A #GChecksum works by feeding a binary blob through g_checksum_update()
// until there is data to be checked; the digest can then be extracted
// using g_checksum_get_string(), which will return the checksum as a
// hexadecimal string; or g_checksum_get_digest(), which will return a
// vector of raw bytes. Once either g_checksum_get_string() or
// g_checksum_get_digest() have been called on a #GChecksum, the checksum
// will be closed and it won't be possible to call g_checksum_update()
// on it anymore.
/*

C function

g_checksum_new
*/
func ChecksumNew(checksumType ChecksumType) *Checksum {
	c_checksum_type := (C.GChecksumType)(checksumType)

	retC := C.g_checksum_new(c_checksum_type)
	retGo := ChecksumNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copies a #GChecksum. If @checksum has been closed, by calling
// g_checksum_get_string() or g_checksum_get_digest(), the copied
// checksum will be closed as well.
/*

C function

g_checksum_copy
*/
func (recv *Checksum) Copy() *Checksum {
	retC := C.g_checksum_copy((*C.GChecksum)(recv.native))
	retGo := ChecksumNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees the memory allocated for @checksum.
/*

C function

g_checksum_free
*/
func (recv *Checksum) Free() {
	C.g_checksum_free((*C.GChecksum)(recv.native))

	return
}

// Gets the digest from @checksum as a raw binary vector and places it
// into @buffer. The size of the digest depends on the type of checksum.
//
// Once this function has been called, the #GChecksum is closed and can
// no longer be updated with g_checksum_update().
/*

C function

g_checksum_get_digest
*/
func (recv *Checksum) GetDigest(buffer uint8, digestLen uint64) {
	c_buffer := (C.guint8)(buffer)

	c_digest_len := (C.gsize)(digestLen)

	C.g_checksum_get_digest((*C.GChecksum)(recv.native), &c_buffer, &c_digest_len)

	return
}

// Gets the digest as an hexadecimal string.
//
// Once this function has been called the #GChecksum can no longer be
// updated with g_checksum_update().
//
// The hexadecimal characters will be lower case.
/*

C function

g_checksum_get_string
*/
func (recv *Checksum) GetString() string {
	retC := C.g_checksum_get_string((*C.GChecksum)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Feeds @data into an existing #GChecksum. The checksum must still be
// open, that is g_checksum_get_string() or g_checksum_get_digest() must
// not have been called on @checksum.
/*

C function

g_checksum_update
*/
func (recv *Checksum) Update(data []uint8) {
	c_data := &data[0]

	c_length := (C.gssize)(len(data))

	C.g_checksum_update((*C.GChecksum)(recv.native), (*C.guchar)(unsafe.Pointer(c_data)), c_length)

	return
}

// Returns the #GHashTable associated with @iter.
/*

C function

g_hash_table_iter_get_hash_table
*/
func (recv *HashTableIter) GetHashTable() *HashTable {
	retC := C.g_hash_table_iter_get_hash_table((*C.GHashTableIter)(recv.native))
	retGo := HashTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Initializes a key/value pair iterator and associates it with
// @hash_table. Modifying the hash table after calling this function
// invalidates the returned iterator.
// |[<!-- language="C" -->
// GHashTableIter iter;
// gpointer key, value;
//
// g_hash_table_iter_init (&iter, hash_table);
// while (g_hash_table_iter_next (&iter, &key, &value))
// {
// do something with key and value
// }
// ]|
/*

C function

g_hash_table_iter_init
*/
func (recv *HashTableIter) Init(hashTable *HashTable) {
	c_hash_table := (*C.GHashTable)(C.NULL)
	if hashTable != nil {
		c_hash_table = (*C.GHashTable)(hashTable.ToC())
	}

	C.g_hash_table_iter_init((*C.GHashTableIter)(recv.native), c_hash_table)

	return
}

// Advances @iter and retrieves the key and/or value that are now
// pointed to as a result of this advancement. If %FALSE is returned,
// @key and @value are not set, and the iterator becomes invalid.
/*

C function

g_hash_table_iter_next
*/
func (recv *HashTableIter) Next() (bool, uintptr, uintptr) {
	var c_key C.gpointer

	var c_value C.gpointer

	retC := C.g_hash_table_iter_next((*C.GHashTableIter)(recv.native), &c_key, &c_value)
	retGo := retC == C.TRUE

	key := (uintptr)(unsafe.Pointer(&c_key))

	value := (uintptr)(unsafe.Pointer(&c_value))

	return retGo, key, value
}

// Removes the key/value pair currently pointed to by the iterator
// from its associated #GHashTable. Can only be called after
// g_hash_table_iter_next() returned %TRUE, and cannot be called
// more than once for the same key/value pair.
//
// If the #GHashTable was created using g_hash_table_new_full(),
// the key and value are freed using the supplied destroy functions,
// otherwise you have to make sure that any dynamically allocated
// values are freed yourself.
//
// It is safe to continue iterating the #GHashTable afterward:
// |[<!-- language="C" -->
// while (g_hash_table_iter_next (&iter, &key, &value))
// {
// if (condition)
// g_hash_table_iter_remove (&iter);
// }
// ]|
/*

C function

g_hash_table_iter_remove
*/
func (recv *HashTableIter) Remove() {
	C.g_hash_table_iter_remove((*C.GHashTableIter)(recv.native))

	return
}

// Removes the key/value pair currently pointed to by the
// iterator from its associated #GHashTable, without calling
// the key and value destroy functions. Can only be called
// after g_hash_table_iter_next() returned %TRUE, and cannot
// be called more than once for the same key/value pair.
/*

C function

g_hash_table_iter_steal
*/
func (recv *HashTableIter) Steal() {
	C.g_hash_table_iter_steal((*C.GHashTableIter)(recv.native))

	return
}

// Retrieves the element stack from the internal state of the parser.
//
// The returned #GSList is a list of strings where the first item is
// the currently open tag (as would be returned by
// g_markup_parse_context_get_element()) and the next item is its
// immediate parent.
//
// This function is intended to be used in the start_element and
// end_element handlers where g_markup_parse_context_get_element()
// would merely return the name of the element that is being
// processed.
/*

C function

g_markup_parse_context_get_element_stack
*/
func (recv *MarkupParseContext) GetElementStack() *SList {
	retC := C.g_markup_parse_context_get_element_stack((*C.GMarkupParseContext)(recv.native))
	retGo := SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Appends @unescaped to @string, escaped any characters that
// are reserved in URIs using URI-style escape sequences.
/*

C function

g_string_append_uri_escaped
*/
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

// Adds @test_case to @suite.
/*

C function

g_test_suite_add
*/
func (recv *TestSuite) Add(testCase *TestCase) {
	c_test_case := (*C.GTestCase)(C.NULL)
	if testCase != nil {
		c_test_case = (*C.GTestCase)(testCase.ToC())
	}

	C.g_test_suite_add((*C.GTestSuite)(recv.native), c_test_case)

	return
}

// Adds @nestedsuite to @suite.
/*

C function

g_test_suite_add_suite
*/
func (recv *TestSuite) AddSuite(nestedsuite *TestSuite) {
	c_nestedsuite := (*C.GTestSuite)(C.NULL)
	if nestedsuite != nil {
		c_nestedsuite = (*C.GTestSuite)(nestedsuite.ToC())
	}

	C.g_test_suite_add_suite((*C.GTestSuite)(recv.native), c_nestedsuite)

	return
}
