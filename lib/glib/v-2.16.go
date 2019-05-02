// This is a generated file - DO NOT EDIT
// +build glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

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

	static void _g_prefix_error(GError** err, const gchar* format) {
		return g_prefix_error(err, format);
    }
*/
/*

	static void _g_propagate_prefixed_error(GError** dest, GError* src, const gchar* format) {
		return g_propagate_prefixed_error(dest, src, format);
    }
*/
/*

	static void _g_test_maximized_result(double maximized_quantity, const char* format) {
		return g_test_maximized_result(maximized_quantity, format);
    }
*/
/*

	static void _g_test_message(const char* format) {
		return g_test_message(format);
    }
*/
/*

	static void _g_test_minimized_result(double minimized_quantity, const char* format) {
		return g_test_minimized_result(minimized_quantity, format);
    }
*/
import "C"

type ChecksumType C.GChecksumType

const (
	CHECKSUM_MD5    ChecksumType = 0
	CHECKSUM_SHA1   ChecksumType = 1
	CHECKSUM_SHA256 ChecksumType = 2
	CHECKSUM_SHA512 ChecksumType = 3
	CHECKSUM_SHA384 ChecksumType = 4
)

// ComputeChecksumForData is a wrapper around the C function g_compute_checksum_for_data.
func ComputeChecksumForData(checksumType ChecksumType, data []uint8) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gsize)(len(data))

	retC := C.g_compute_checksum_for_data(c_checksum_type, c_data, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// ComputeChecksumForString is a wrapper around the C function g_compute_checksum_for_string.
func ComputeChecksumForString(checksumType ChecksumType, str string) string {
	c_checksum_type := (C.GChecksumType)(checksumType)

	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	c_length := (C.gssize)(len(str))

	retC := C.g_compute_checksum_for_string(c_checksum_type, c_str, c_length)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Dpgettext is a wrapper around the C function g_dpgettext.
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) string {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_msgctxtid := C.CString(msgctxtid)
	defer C.free(unsafe.Pointer(c_msgctxtid))

	c_msgidoffset := (C.gsize)(msgidoffset)

	retC := C.g_dpgettext(c_domain, c_msgctxtid, c_msgidoffset)
	retGo := C.GoString(retC)

	return retGo
}

// FormatSizeForDisplay is a wrapper around the C function g_format_size_for_display.
func FormatSizeForDisplay(size int64) string {
	c_size := (C.goffset)(size)

	retC := C.g_format_size_for_display(c_size)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_markup_collect_attributes : unsupported parameter attribute_names : in string with indirection level of 2

// PrefixError is a wrapper around the C function g_prefix_error.
func PrefixError(err *Error, format string, args ...interface{}) {
	c_err := (**C.GError)(C.NULL)
	if err != nil {
		c_err = (**C.GError)(err.ToC())
	}

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_prefix_error(c_err, c_format)

	return
}

// PropagatePrefixedError is a wrapper around the C function g_propagate_prefixed_error.
func PropagatePrefixedError(dest *Error, src *Error, format string, args ...interface{}) {
	c_dest := (**C.GError)(C.NULL)
	if dest != nil {
		c_dest = (**C.GError)(dest.ToC())
	}

	c_src := (*C.GError)(C.NULL)
	if src != nil {
		c_src = (*C.GError)(src.ToC())
	}

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_propagate_prefixed_error(c_dest, c_src, c_format)

	return
}

// Strcmp0 is a wrapper around the C function g_strcmp0.
func Strcmp0(str1 string, str2 string) int32 {
	c_str1 := C.CString(str1)
	defer C.free(unsafe.Pointer(c_str1))

	c_str2 := C.CString(str2)
	defer C.free(unsafe.Pointer(c_str2))

	retC := C.g_strcmp0(c_str1, c_str2)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_test_add_func : unsupported parameter test_func : no type generator for TestFunc (GTestFunc) for param test_func

// TestBug is a wrapper around the C function g_test_bug.
func TestBug(bugUriSnippet string) {
	c_bug_uri_snippet := C.CString(bugUriSnippet)
	defer C.free(unsafe.Pointer(c_bug_uri_snippet))

	C.g_test_bug(c_bug_uri_snippet)

	return
}

// TestBugBase is a wrapper around the C function g_test_bug_base.
func TestBugBase(uriPattern string) {
	c_uri_pattern := C.CString(uriPattern)
	defer C.free(unsafe.Pointer(c_uri_pattern))

	C.g_test_bug_base(c_uri_pattern)

	return
}

// TestCreateSuite is a wrapper around the C function g_test_create_suite.
func TestCreateSuite(suiteName string) *TestSuite {
	c_suite_name := C.CString(suiteName)
	defer C.free(unsafe.Pointer(c_suite_name))

	retC := C.g_test_create_suite(c_suite_name)
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TestGetRoot is a wrapper around the C function g_test_get_root.
func TestGetRoot() *TestSuite {
	retC := C.g_test_get_root()
	retGo := TestSuiteNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_test_init : unsupported parameter ... : varargs

// TestMaximizedResult is a wrapper around the C function g_test_maximized_result.
func TestMaximizedResult(maximizedQuantity float64, format string, args ...interface{}) {
	c_maximized_quantity := (C.double)(maximizedQuantity)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_test_maximized_result(c_maximized_quantity, c_format)

	return
}

// TestMessage is a wrapper around the C function g_test_message.
func TestMessage(format string, args ...interface{}) {
	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_test_message(c_format)

	return
}

// TestMinimizedResult is a wrapper around the C function g_test_minimized_result.
func TestMinimizedResult(minimizedQuantity float64, format string, args ...interface{}) {
	c_minimized_quantity := (C.double)(minimizedQuantity)

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	C._g_test_minimized_result(c_minimized_quantity, c_format)

	return
}

// Unsupported : g_test_queue_destroy : unsupported parameter destroy_func : no type generator for DestroyNotify (GDestroyNotify) for param destroy_func

// Unsupported : g_test_queue_free : unsupported parameter gfree_pointer : no type generator for gpointer (gpointer) for param gfree_pointer

// TestRandDouble is a wrapper around the C function g_test_rand_double.
func TestRandDouble() float64 {
	retC := C.g_test_rand_double()
	retGo := (float64)(retC)

	return retGo
}

// TestRandDoubleRange is a wrapper around the C function g_test_rand_double_range.
func TestRandDoubleRange(rangeStart float64, rangeEnd float64) float64 {
	c_range_start := (C.double)(rangeStart)

	c_range_end := (C.double)(rangeEnd)

	retC := C.g_test_rand_double_range(c_range_start, c_range_end)
	retGo := (float64)(retC)

	return retGo
}

// TestRandInt is a wrapper around the C function g_test_rand_int.
func TestRandInt() int32 {
	retC := C.g_test_rand_int()
	retGo := (int32)(retC)

	return retGo
}

// TestRandIntRange is a wrapper around the C function g_test_rand_int_range.
func TestRandIntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_test_rand_int_range(c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// TestRun is a wrapper around the C function g_test_run.
func TestRun() int32 {
	retC := C.g_test_run()
	retGo := (int32)(retC)

	return retGo
}

// TestRunSuite is a wrapper around the C function g_test_run_suite.
func TestRunSuite(suite *TestSuite) int32 {
	c_suite := (*C.GTestSuite)(C.NULL)
	if suite != nil {
		c_suite = (*C.GTestSuite)(suite.ToC())
	}

	retC := C.g_test_run_suite(c_suite)
	retGo := (int32)(retC)

	return retGo
}

// TestTimerElapsed is a wrapper around the C function g_test_timer_elapsed.
func TestTimerElapsed() float64 {
	retC := C.g_test_timer_elapsed()
	retGo := (float64)(retC)

	return retGo
}

// TestTimerLast is a wrapper around the C function g_test_timer_last.
func TestTimerLast() float64 {
	retC := C.g_test_timer_last()
	retGo := (float64)(retC)

	return retGo
}

// TestTimerStart is a wrapper around the C function g_test_timer_start.
func TestTimerStart() {
	C.g_test_timer_start()

	return
}

// TestTrapFork is a wrapper around the C function g_test_trap_fork.
func TestTrapFork(usecTimeout uint64, testTrapFlags TestTrapFlags) bool {
	c_usec_timeout := (C.guint64)(usecTimeout)

	c_test_trap_flags := (C.GTestTrapFlags)(testTrapFlags)

	retC := C.g_test_trap_fork(c_usec_timeout, c_test_trap_flags)
	retGo := retC == C.TRUE

	return retGo
}

// TestTrapHasPassed is a wrapper around the C function g_test_trap_has_passed.
func TestTrapHasPassed() bool {
	retC := C.g_test_trap_has_passed()
	retGo := retC == C.TRUE

	return retGo
}

// TestTrapReachedTimeout is a wrapper around the C function g_test_trap_reached_timeout.
func TestTrapReachedTimeout() bool {
	retC := C.g_test_trap_reached_timeout()
	retGo := retC == C.TRUE

	return retGo
}

// UriEscapeString is a wrapper around the C function g_uri_escape_string.
func UriEscapeString(unescaped string, reservedCharsAllowed string, allowUtf8 bool) string {
	c_unescaped := C.CString(unescaped)
	defer C.free(unsafe.Pointer(c_unescaped))

	c_reserved_chars_allowed := C.CString(reservedCharsAllowed)
	defer C.free(unsafe.Pointer(c_reserved_chars_allowed))

	c_allow_utf8 :=
		boolToGboolean(allowUtf8)

	retC := C.g_uri_escape_string(c_unescaped, c_reserved_chars_allowed, c_allow_utf8)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UriParseScheme is a wrapper around the C function g_uri_parse_scheme.
func UriParseScheme(uri string) string {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.g_uri_parse_scheme(c_uri)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UriUnescapeSegment is a wrapper around the C function g_uri_unescape_segment.
func UriUnescapeSegment(escapedString string, escapedStringEnd string, illegalCharacters string) string {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_escaped_string_end := C.CString(escapedStringEnd)
	defer C.free(unsafe.Pointer(c_escaped_string_end))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	retC := C.g_uri_unescape_segment(c_escaped_string, c_escaped_string_end, c_illegal_characters)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// UriUnescapeString is a wrapper around the C function g_uri_unescape_string.
func UriUnescapeString(escapedString string, illegalCharacters string) string {
	c_escaped_string := C.CString(escapedString)
	defer C.free(unsafe.Pointer(c_escaped_string))

	c_illegal_characters := C.CString(illegalCharacters)
	defer C.free(unsafe.Pointer(c_illegal_characters))

	retC := C.g_uri_unescape_string(c_escaped_string, c_illegal_characters)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

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
	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guchar)(unsafe.Pointer(c_data_arrayPtr))

	c_length := (C.gssize)(len(data))

	C.g_checksum_update((*C.GChecksum)(recv.native), c_data, c_length)

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

// Unsupported : g_hash_table_iter_next : unsupported parameter key : no type generator for gpointer (gpointer*) for param key

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
