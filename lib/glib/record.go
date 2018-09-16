// This is a generated file - DO NOT EDIT

package glib

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Array is a wrapper around the C record GArray.
type Array struct {
	native *C.GArray
	Data   string
	Len    uint32
}

func arrayNewFromC(c *C.GArray) *Array {
	if c == nil {
		return nil
	}

	return &Array{
		Data: C.GoString(c.data),
		Len:  (uint32)(c.len),
	}
}

// Asyncqueue is a wrapper around the C record GAsyncQueue.
type Asyncqueue struct {
	native *C.GAsyncQueue
}

func asyncqueueNewFromC(c *C.GAsyncQueue) *Asyncqueue {
	if c == nil {
		return nil
	}

	return &Asyncqueue{}
}

// Bookmarkfile is a wrapper around the C record GBookmarkFile.
type Bookmarkfile struct {
	native *C.GBookmarkFile
}

func bookmarkfileNewFromC(c *C.GBookmarkFile) *Bookmarkfile {
	if c == nil {
		return nil
	}

	return &Bookmarkfile{}
}

// Bytearray is a wrapper around the C record GByteArray.
type Bytearray struct {
	native *C.GByteArray
	// data : no type generator for guint8, guint8*
	Len uint32
}

func bytearrayNewFromC(c *C.GByteArray) *Bytearray {
	if c == nil {
		return nil
	}

	return &Bytearray{Len: (uint32)(c.len)}
}

// Cond is a wrapper around the C record GCond.
type Cond struct {
	native *C.GCond
	P      uintptr
	// no type for i
}

func condNewFromC(c *C.GCond) *Cond {
	if c == nil {
		return nil
	}

	return &Cond{P: (uintptr)(c.p)}
}

// Data is a wrapper around the C record GData.
type Data struct {
	native *C.GData
}

func dataNewFromC(c *C.GData) *Data {
	if c == nil {
		return nil
	}

	return &Data{}
}

// Date is a wrapper around the C record GDate.
type Date struct {
	native *C.GDate
	// Bitfield not supported : 32 julian_days
	// Bitfield not supported :  1 julian
	// Bitfield not supported :  1 dmy
	// Bitfield not supported :  6 day
	// Bitfield not supported :  4 month
	// Bitfield not supported : 16 year
}

func dateNewFromC(c *C.GDate) *Date {
	if c == nil {
		return nil
	}

	return &Date{}
}

// Debugkey is a wrapper around the C record GDebugKey.
type Debugkey struct {
	native *C.GDebugKey
	Key    string
	Value  uint32
}

func debugkeyNewFromC(c *C.GDebugKey) *Debugkey {
	if c == nil {
		return nil
	}

	return &Debugkey{
		Key:   C.GoString(c.key),
		Value: (uint32)(c.value),
	}
}

// Dir is a wrapper around the C record GDir.
type Dir struct {
	native *C.GDir
}

func dirNewFromC(c *C.GDir) *Dir {
	if c == nil {
		return nil
	}

	return &Dir{}
}

// Error is a wrapper around the C record GError.
type Error struct {
	native  *C.GError
	Domain  Quark
	Code    int32
	Message string
}

func errorNewFromC(c *C.GError) *Error {
	if c == nil {
		return nil
	}

	return &Error{
		Code:    (int32)(c.code),
		Domain:  (Quark)(c.domain),
		Message: C.GoString(c.message),
	}
}

// Hashtable is a wrapper around the C record GHashTable.
type Hashtable struct {
	native *C.GHashTable
}

func hashtableNewFromC(c *C.GHashTable) *Hashtable {
	if c == nil {
		return nil
	}

	return &Hashtable{}
}

// Hashtableiter is a wrapper around the C record GHashTableIter.
type Hashtableiter struct {
	native *C.GHashTableIter
	Dummy1 uintptr
	Dummy2 uintptr
	Dummy3 uintptr
	Dummy4 int32
	// dummy5 : no type generator for gboolean, gboolean
	Dummy6 uintptr
}

func hashtableiterNewFromC(c *C.GHashTableIter) *Hashtableiter {
	if c == nil {
		return nil
	}

	return &Hashtableiter{
		Dummy1: (uintptr)(c.dummy1),
		Dummy2: (uintptr)(c.dummy2),
		Dummy3: (uintptr)(c.dummy3),
		Dummy4: (int32)(c.dummy4),
		Dummy6: (uintptr)(c.dummy6),
	}
}

// Hook is a wrapper around the C record GHook.
type Hook struct {
	native *C.GHook
	Data   uintptr
	// next : no type generator for Hook, GHook*
	// prev : no type generator for Hook, GHook*
	RefCount uint32
	HookId   uint64
	Flags    uint32
	Func     uintptr
	// destroy : no type generator for DestroyNotify, GDestroyNotify
}

func hookNewFromC(c *C.GHook) *Hook {
	if c == nil {
		return nil
	}

	return &Hook{
		Data:     (uintptr)(c.data),
		Flags:    (uint32)(c.flags),
		Func:     (uintptr)(c._func),
		HookId:   (uint64)(c.hook_id),
		RefCount: (uint32)(c.ref_count),
	}
}

// Hooklist is a wrapper around the C record GHookList.
type Hooklist struct {
	native *C.GHookList
	SeqId  uint64
	// Bitfield not supported : 16 hook_size
	// Bitfield not supported :  1 is_setup
	// hooks : no type generator for Hook, GHook*
	Dummy3 uintptr
	// finalize_hook : no type generator for HookFinalizeFunc, GHookFinalizeFunc
	// no type for dummy
}

func hooklistNewFromC(c *C.GHookList) *Hooklist {
	if c == nil {
		return nil
	}

	return &Hooklist{
		Dummy3: (uintptr)(c.dummy3),
		SeqId:  (uint64)(c.seq_id),
	}
}

// Blacklisted : GIConv

// Iochannel is a wrapper around the C record GIOChannel.
type Iochannel struct {
	native   *C.GIOChannel
	RefCount int32
	// funcs : no type generator for IOFuncs, GIOFuncs*
	Encoding string
	// read_cd : no type generator for IConv, GIConv
	// write_cd : no type generator for IConv, GIConv
	LineTerm    string
	LineTermLen uint32
	BufSize     uint64
	// read_buf : no type generator for String, GString*
	// encoded_read_buf : no type generator for String, GString*
	// write_buf : no type generator for String, GString*
	// no type for partial_write_buf
	// Bitfield not supported :  1 use_buffer
	// Bitfield not supported :  1 do_encode
	// Bitfield not supported :  1 close_on_unref
	// Bitfield not supported :  1 is_readable
	// Bitfield not supported :  1 is_writeable
	// Bitfield not supported :  1 is_seekable
	Reserved1 uintptr
	Reserved2 uintptr
}

func iochannelNewFromC(c *C.GIOChannel) *Iochannel {
	if c == nil {
		return nil
	}

	return &Iochannel{
		BufSize:     (uint64)(c.buf_size),
		Encoding:    C.GoString(c.encoding),
		LineTerm:    C.GoString(c.line_term),
		LineTermLen: (uint32)(c.line_term_len),
		RefCount:    (int32)(c.ref_count),
		Reserved1:   (uintptr)(c.reserved1),
		Reserved2:   (uintptr)(c.reserved2),
	}
}

// Iofuncs is a wrapper around the C record GIOFuncs.
type Iofuncs struct {
	native *C.GIOFuncs
	// no type for io_read
	// no type for io_write
	// no type for io_seek
	// no type for io_close
	// no type for io_create_watch
	// no type for io_free
	// no type for io_set_flags
	// no type for io_get_flags
}

func iofuncsNewFromC(c *C.GIOFuncs) *Iofuncs {
	if c == nil {
		return nil
	}

	return &Iofuncs{}
}

// Keyfile is a wrapper around the C record GKeyFile.
type Keyfile struct {
	native *C.GKeyFile
}

func keyfileNewFromC(c *C.GKeyFile) *Keyfile {
	if c == nil {
		return nil
	}

	return &Keyfile{}
}

// List is a wrapper around the C record GList.
type List struct {
	native *C.GList
	Data   uintptr
	// next : no type generator for GLib.List, GList*
	// prev : no type generator for GLib.List, GList*
}

func listNewFromC(c *C.GList) *List {
	if c == nil {
		return nil
	}

	return &List{Data: (uintptr)(c.data)}
}

// Maincontext is a wrapper around the C record GMainContext.
type Maincontext struct {
	native *C.GMainContext
}

func maincontextNewFromC(c *C.GMainContext) *Maincontext {
	if c == nil {
		return nil
	}

	return &Maincontext{}
}

// Mainloop is a wrapper around the C record GMainLoop.
type Mainloop struct {
	native *C.GMainLoop
}

func mainloopNewFromC(c *C.GMainLoop) *Mainloop {
	if c == nil {
		return nil
	}

	return &Mainloop{}
}

// Mappedfile is a wrapper around the C record GMappedFile.
type Mappedfile struct {
	native *C.GMappedFile
}

func mappedfileNewFromC(c *C.GMappedFile) *Mappedfile {
	if c == nil {
		return nil
	}

	return &Mappedfile{}
}

// Markupparsecontext is a wrapper around the C record GMarkupParseContext.
type Markupparsecontext struct {
	native *C.GMarkupParseContext
}

func markupparsecontextNewFromC(c *C.GMarkupParseContext) *Markupparsecontext {
	if c == nil {
		return nil
	}

	return &Markupparsecontext{}
}

// Markupparser is a wrapper around the C record GMarkupParser.
type Markupparser struct {
	native *C.GMarkupParser
	// no type for start_element
	// no type for end_element
	// no type for text
	// no type for passthrough
	// no type for error
}

func markupparserNewFromC(c *C.GMarkupParser) *Markupparser {
	if c == nil {
		return nil
	}

	return &Markupparser{}
}

// Matchinfo is a wrapper around the C record GMatchInfo.
type Matchinfo struct {
	native *C.GMatchInfo
}

func matchinfoNewFromC(c *C.GMatchInfo) *Matchinfo {
	if c == nil {
		return nil
	}

	return &Matchinfo{}
}

// Memvtable is a wrapper around the C record GMemVTable.
type Memvtable struct {
	native *C.GMemVTable
	// no type for malloc
	// no type for realloc
	// no type for free
	// no type for calloc
	// no type for try_malloc
	// no type for try_realloc
}

func memvtableNewFromC(c *C.GMemVTable) *Memvtable {
	if c == nil {
		return nil
	}

	return &Memvtable{}
}

// Node is a wrapper around the C record GNode.
type Node struct {
	native *C.GNode
	Data   uintptr
	// next : no type generator for Node, GNode*
	// prev : no type generator for Node, GNode*
	// parent : no type generator for Node, GNode*
	// children : no type generator for Node, GNode*
}

func nodeNewFromC(c *C.GNode) *Node {
	if c == nil {
		return nil
	}

	return &Node{Data: (uintptr)(c.data)}
}

// Optioncontext is a wrapper around the C record GOptionContext.
type Optioncontext struct {
	native *C.GOptionContext
}

func optioncontextNewFromC(c *C.GOptionContext) *Optioncontext {
	if c == nil {
		return nil
	}

	return &Optioncontext{}
}

// Optionentry is a wrapper around the C record GOptionEntry.
type Optionentry struct {
	native    *C.GOptionEntry
	LongName  string
	ShortName rune
	Flags     int32
	// arg : no type generator for OptionArg, GOptionArg
	ArgData        uintptr
	Description    string
	ArgDescription string
}

func optionentryNewFromC(c *C.GOptionEntry) *Optionentry {
	if c == nil {
		return nil
	}

	return &Optionentry{
		ArgData:        (uintptr)(c.arg_data),
		ArgDescription: C.GoString(c.arg_description),
		Description:    C.GoString(c.description),
		Flags:          (int32)(c.flags),
		LongName:       C.GoString(c.long_name),
		ShortName:      (rune)(c.short_name),
	}
}

// Optiongroup is a wrapper around the C record GOptionGroup.
type Optiongroup struct {
	native *C.GOptionGroup
}

func optiongroupNewFromC(c *C.GOptionGroup) *Optiongroup {
	if c == nil {
		return nil
	}

	return &Optiongroup{}
}

// Patternspec is a wrapper around the C record GPatternSpec.
type Patternspec struct {
	native *C.GPatternSpec
}

func patternspecNewFromC(c *C.GPatternSpec) *Patternspec {
	if c == nil {
		return nil
	}

	return &Patternspec{}
}

// Pollfd is a wrapper around the C record GPollFD.
type Pollfd struct {
	native  *C.GPollFD
	Fd      int32
	Events  uint32
	Revents uint32
}

func pollfdNewFromC(c *C.GPollFD) *Pollfd {
	if c == nil {
		return nil
	}

	return &Pollfd{
		Events:  (uint32)(c.events),
		Fd:      (int32)(c.fd),
		Revents: (uint32)(c.revents),
	}
}

// Private is a wrapper around the C record GPrivate.
type Private struct {
	native *C.GPrivate
	P      uintptr
	// notify : no type generator for DestroyNotify, GDestroyNotify
	// no type for future
}

func privateNewFromC(c *C.GPrivate) *Private {
	if c == nil {
		return nil
	}

	return &Private{P: (uintptr)(c.p)}
}

// Ptrarray is a wrapper around the C record GPtrArray.
type Ptrarray struct {
	native *C.GPtrArray
	// pdata : no type generator for gpointer, gpointer*
	Len uint32
}

func ptrarrayNewFromC(c *C.GPtrArray) *Ptrarray {
	if c == nil {
		return nil
	}

	return &Ptrarray{Len: (uint32)(c.len)}
}

// Queue is a wrapper around the C record GQueue.
type Queue struct {
	native *C.GQueue
	// head : no type generator for GLib.List, GList*
	// tail : no type generator for GLib.List, GList*
	Length uint32
}

func queueNewFromC(c *C.GQueue) *Queue {
	if c == nil {
		return nil
	}

	return &Queue{Length: (uint32)(c.length)}
}

// Rand is a wrapper around the C record GRand.
type Rand struct {
	native *C.GRand
}

func randNewFromC(c *C.GRand) *Rand {
	if c == nil {
		return nil
	}

	return &Rand{}
}

// Slist is a wrapper around the C record GSList.
type Slist struct {
	native *C.GSList
	Data   uintptr
	// next : no type generator for GLib.SList, GSList*
}

func slistNewFromC(c *C.GSList) *Slist {
	if c == nil {
		return nil
	}

	return &Slist{Data: (uintptr)(c.data)}
}

// Scanner is a wrapper around the C record GScanner.
type Scanner struct {
	native         *C.GScanner
	UserData       uintptr
	MaxParseErrors uint32
	ParseErrors    uint32
	InputName      string
	// qdata : no type generator for Data, GData*
	// config : no type generator for ScannerConfig, GScannerConfig*
	// token : no type generator for TokenType, GTokenType
	// value : no type generator for TokenValue, GTokenValue
	Line     uint32
	Position uint32
	// next_token : no type generator for TokenType, GTokenType
	// next_value : no type generator for TokenValue, GTokenValue
	NextLine     uint32
	NextPosition uint32
	// symbol_table : no type generator for GLib.HashTable, GHashTable*
	InputFd int32
	Text    string
	TextEnd string
	Buffer  string
	ScopeId uint32
	// msg_handler : no type generator for ScannerMsgFunc, GScannerMsgFunc
}

func scannerNewFromC(c *C.GScanner) *Scanner {
	if c == nil {
		return nil
	}

	return &Scanner{
		Buffer:         C.GoString(c.buffer),
		InputFd:        (int32)(c.input_fd),
		InputName:      C.GoString(c.input_name),
		Line:           (uint32)(c.line),
		MaxParseErrors: (uint32)(c.max_parse_errors),
		NextLine:       (uint32)(c.next_line),
		NextPosition:   (uint32)(c.next_position),
		ParseErrors:    (uint32)(c.parse_errors),
		Position:       (uint32)(c.position),
		ScopeId:        (uint32)(c.scope_id),
		Text:           C.GoString(c.text),
		TextEnd:        C.GoString(c.text_end),
		UserData:       (uintptr)(c.user_data),
	}
}

// Scannerconfig is a wrapper around the C record GScannerConfig.
type Scannerconfig struct {
	native              *C.GScannerConfig
	CsetSkipCharacters  string
	CsetIdentifierFirst string
	CsetIdentifierNth   string
	CpairCommentSingle  string
	// Bitfield not supported :  1 case_sensitive
	// Bitfield not supported :  1 skip_comment_multi
	// Bitfield not supported :  1 skip_comment_single
	// Bitfield not supported :  1 scan_comment_multi
	// Bitfield not supported :  1 scan_identifier
	// Bitfield not supported :  1 scan_identifier_1char
	// Bitfield not supported :  1 scan_identifier_NULL
	// Bitfield not supported :  1 scan_symbols
	// Bitfield not supported :  1 scan_binary
	// Bitfield not supported :  1 scan_octal
	// Bitfield not supported :  1 scan_float
	// Bitfield not supported :  1 scan_hex
	// Bitfield not supported :  1 scan_hex_dollar
	// Bitfield not supported :  1 scan_string_sq
	// Bitfield not supported :  1 scan_string_dq
	// Bitfield not supported :  1 numbers_2_int
	// Bitfield not supported :  1 int_2_float
	// Bitfield not supported :  1 identifier_2_string
	// Bitfield not supported :  1 char_2_token
	// Bitfield not supported :  1 symbol_2_token
	// Bitfield not supported :  1 scope_0_fallback
	// Bitfield not supported :  1 store_int64
	PaddingDummy uint32
}

func scannerconfigNewFromC(c *C.GScannerConfig) *Scannerconfig {
	if c == nil {
		return nil
	}

	return &Scannerconfig{
		CpairCommentSingle:  C.GoString(c.cpair_comment_single),
		CsetIdentifierFirst: C.GoString(c.cset_identifier_first),
		CsetIdentifierNth:   C.GoString(c.cset_identifier_nth),
		CsetSkipCharacters:  C.GoString(c.cset_skip_characters),
		PaddingDummy:        (uint32)(c.padding_dummy),
	}
}

// Sequence is a wrapper around the C record GSequence.
type Sequence struct {
	native *C.GSequence
}

func sequenceNewFromC(c *C.GSequence) *Sequence {
	if c == nil {
		return nil
	}

	return &Sequence{}
}

// Sequenceiter is a wrapper around the C record GSequenceIter.
type Sequenceiter struct {
	native *C.GSequenceIter
}

func sequenceiterNewFromC(c *C.GSequenceIter) *Sequenceiter {
	if c == nil {
		return nil
	}

	return &Sequenceiter{}
}

// Source is a wrapper around the C record GSource.
type Source struct {
	native       *C.GSource
	CallbackData uintptr
	// callback_funcs : no type generator for SourceCallbackFuncs, GSourceCallbackFuncs*
	// source_funcs : no type generator for SourceFuncs, const GSourceFuncs*
	RefCount uint32
	// context : no type generator for MainContext, GMainContext*
	Priority int32
	Flags    uint32
	SourceId uint32
	// poll_fds : no type generator for GLib.SList, GSList*
	// prev : no type generator for Source, GSource*
	// next : no type generator for Source, GSource*
	Name string
	// priv : no type generator for SourcePrivate, GSourcePrivate*
}

func sourceNewFromC(c *C.GSource) *Source {
	if c == nil {
		return nil
	}

	return &Source{
		CallbackData: (uintptr)(c.callback_data),
		Flags:        (uint32)(c.flags),
		Name:         C.GoString(c.name),
		Priority:     (int32)(c.priority),
		RefCount:     (uint32)(c.ref_count),
		SourceId:     (uint32)(c.source_id),
	}
}

// Sourcecallbackfuncs is a wrapper around the C record GSourceCallbackFuncs.
type Sourcecallbackfuncs struct {
	native *C.GSourceCallbackFuncs
	// no type for ref
	// no type for unref
	// no type for get
}

func sourcecallbackfuncsNewFromC(c *C.GSourceCallbackFuncs) *Sourcecallbackfuncs {
	if c == nil {
		return nil
	}

	return &Sourcecallbackfuncs{}
}

// Sourcefuncs is a wrapper around the C record GSourceFuncs.
type Sourcefuncs struct {
	native *C.GSourceFuncs
	// no type for prepare
	// no type for check
	// no type for dispatch
	// no type for finalize
	// closure_callback : no type generator for SourceFunc, GSourceFunc
	// closure_marshal : no type generator for SourceDummyMarshal, GSourceDummyMarshal
}

func sourcefuncsNewFromC(c *C.GSourceFuncs) *Sourcefuncs {
	if c == nil {
		return nil
	}

	return &Sourcefuncs{}
}

// Sourceprivate is a wrapper around the C record GSourcePrivate.
type Sourceprivate struct {
	native *C.GSourcePrivate
}

func sourceprivateNewFromC(c *C.GSourcePrivate) *Sourceprivate {
	if c == nil {
		return nil
	}

	return &Sourceprivate{}
}

// Statbuf is a wrapper around the C record GStatBuf.
type Statbuf struct {
	native *C.GStatBuf
}

func statbufNewFromC(c *C.GStatBuf) *Statbuf {
	if c == nil {
		return nil
	}

	return &Statbuf{}
}

// String is a wrapper around the C record GString.
type String struct {
	native       *C.GString
	Str          string
	Len          uint64
	AllocatedLen uint64
}

func stringNewFromC(c *C.GString) *String {
	if c == nil {
		return nil
	}

	return &String{
		AllocatedLen: (uint64)(c.allocated_len),
		Len:          (uint64)(c.len),
		Str:          C.GoString(c.str),
	}
}

// Stringchunk is a wrapper around the C record GStringChunk.
type Stringchunk struct {
	native *C.GStringChunk
}

func stringchunkNewFromC(c *C.GStringChunk) *Stringchunk {
	if c == nil {
		return nil
	}

	return &Stringchunk{}
}

// Testcase is a wrapper around the C record GTestCase.
type Testcase struct {
	native *C.GTestCase
}

func testcaseNewFromC(c *C.GTestCase) *Testcase {
	if c == nil {
		return nil
	}

	return &Testcase{}
}

// Testconfig is a wrapper around the C record GTestConfig.
type Testconfig struct {
	native *C.GTestConfig
	// test_initialized : no type generator for gboolean, gboolean
	// test_quick : no type generator for gboolean, gboolean
	// test_perf : no type generator for gboolean, gboolean
	// test_verbose : no type generator for gboolean, gboolean
	// test_quiet : no type generator for gboolean, gboolean
	// test_undefined : no type generator for gboolean, gboolean
}

func testconfigNewFromC(c *C.GTestConfig) *Testconfig {
	if c == nil {
		return nil
	}

	return &Testconfig{}
}

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// Testsuite is a wrapper around the C record GTestSuite.
type Testsuite struct {
	native *C.GTestSuite
}

func testsuiteNewFromC(c *C.GTestSuite) *Testsuite {
	if c == nil {
		return nil
	}

	return &Testsuite{}
}

// Thread is a wrapper around the C record GThread.
type Thread struct {
	native *C.GThread
}

func threadNewFromC(c *C.GThread) *Thread {
	if c == nil {
		return nil
	}

	return &Thread{}
}

// Threadpool is a wrapper around the C record GThreadPool.
type Threadpool struct {
	native *C.GThreadPool
	// _func : no type generator for Func, GFunc
	UserData uintptr
	// exclusive : no type generator for gboolean, gboolean
}

func threadpoolNewFromC(c *C.GThreadPool) *Threadpool {
	if c == nil {
		return nil
	}

	return &Threadpool{UserData: (uintptr)(c.user_data)}
}

// Timeval is a wrapper around the C record GTimeVal.
type Timeval struct {
	native *C.GTimeVal
	TvSec  int64
	TvUsec int64
}

func timevalNewFromC(c *C.GTimeVal) *Timeval {
	if c == nil {
		return nil
	}

	return &Timeval{
		TvSec:  (int64)(c.tv_sec),
		TvUsec: (int64)(c.tv_usec),
	}
}

// Timer is a wrapper around the C record GTimer.
type Timer struct {
	native *C.GTimer
}

func timerNewFromC(c *C.GTimer) *Timer {
	if c == nil {
		return nil
	}

	return &Timer{}
}

// Trashstack is a wrapper around the C record GTrashStack.
type Trashstack struct {
	native *C.GTrashStack
	// next : no type generator for TrashStack, GTrashStack*
}

func trashstackNewFromC(c *C.GTrashStack) *Trashstack {
	if c == nil {
		return nil
	}

	return &Trashstack{}
}

// Tree is a wrapper around the C record GTree.
type Tree struct {
	native *C.GTree
}

func treeNewFromC(c *C.GTree) *Tree {
	if c == nil {
		return nil
	}

	return &Tree{}
}

// Variantbuilder is a wrapper around the C record GVariantBuilder.
type Variantbuilder struct {
	native *C.GVariantBuilder
}

func variantbuilderNewFromC(c *C.GVariantBuilder) *Variantbuilder {
	if c == nil {
		return nil
	}

	return &Variantbuilder{}
}

// Variantiter is a wrapper around the C record GVariantIter.
type Variantiter struct {
	native *C.GVariantIter
	// no type for x
}

func variantiterNewFromC(c *C.GVariantIter) *Variantiter {
	if c == nil {
		return nil
	}

	return &Variantiter{}
}

// Blacklisted : GVariantType
