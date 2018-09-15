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

// Asyncqueue is a wrapper around the C record GAsyncQueue.
type Asyncqueue struct {
	native *C.GAsyncQueue
}

// Bookmarkfile is a wrapper around the C record GBookmarkFile.
type Bookmarkfile struct {
	native *C.GBookmarkFile
}

// Bytearray is a wrapper around the C record GByteArray.
type Bytearray struct {
	native *C.GByteArray
	// data : no type generator for guint8, guint8*
	Len uint32
}

// Cond is a wrapper around the C record GCond.
type Cond struct {
	native *C.GCond
	P      uintptr
	// no type for i
}

// Data is a wrapper around the C record GData.
type Data struct {
	native *C.GData
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

// Debugkey is a wrapper around the C record GDebugKey.
type Debugkey struct {
	native *C.GDebugKey
	Key    string
	Value  uint32
}

// Dir is a wrapper around the C record GDir.
type Dir struct {
	native *C.GDir
}

// Error is a wrapper around the C record GError.
type Error struct {
	native *C.GError
	// domain : no type generator for Quark, GQuark
	Code    int32
	Message string
}

// Hashtable is a wrapper around the C record GHashTable.
type Hashtable struct {
	native *C.GHashTable
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

// Keyfile is a wrapper around the C record GKeyFile.
type Keyfile struct {
	native *C.GKeyFile
}

// List is a wrapper around the C record GList.
type List struct {
	native *C.GList
	Data   uintptr
	// next : no type generator for GLib.List, GList*
	// prev : no type generator for GLib.List, GList*
}

// Maincontext is a wrapper around the C record GMainContext.
type Maincontext struct {
	native *C.GMainContext
}

// Mainloop is a wrapper around the C record GMainLoop.
type Mainloop struct {
	native *C.GMainLoop
}

// Mappedfile is a wrapper around the C record GMappedFile.
type Mappedfile struct {
	native *C.GMappedFile
}

// Markupparsecontext is a wrapper around the C record GMarkupParseContext.
type Markupparsecontext struct {
	native *C.GMarkupParseContext
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

// Matchinfo is a wrapper around the C record GMatchInfo.
type Matchinfo struct {
	native *C.GMatchInfo
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

// Node is a wrapper around the C record GNode.
type Node struct {
	native *C.GNode
	Data   uintptr
	// next : no type generator for Node, GNode*
	// prev : no type generator for Node, GNode*
	// parent : no type generator for Node, GNode*
	// children : no type generator for Node, GNode*
}

// Optioncontext is a wrapper around the C record GOptionContext.
type Optioncontext struct {
	native *C.GOptionContext
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

// Optiongroup is a wrapper around the C record GOptionGroup.
type Optiongroup struct {
	native *C.GOptionGroup
}

// Patternspec is a wrapper around the C record GPatternSpec.
type Patternspec struct {
	native *C.GPatternSpec
}

// Pollfd is a wrapper around the C record GPollFD.
type Pollfd struct {
	native  *C.GPollFD
	Fd      int32
	Events  uint32
	Revents uint32
}

// Private is a wrapper around the C record GPrivate.
type Private struct {
	native *C.GPrivate
	P      uintptr
	// notify : no type generator for DestroyNotify, GDestroyNotify
	// no type for future
}

// Ptrarray is a wrapper around the C record GPtrArray.
type Ptrarray struct {
	native *C.GPtrArray
	// pdata : no type generator for gpointer, gpointer*
	Len uint32
}

// Queue is a wrapper around the C record GQueue.
type Queue struct {
	native *C.GQueue
	// head : no type generator for GLib.List, GList*
	// tail : no type generator for GLib.List, GList*
	Length uint32
}

// Rand is a wrapper around the C record GRand.
type Rand struct {
	native *C.GRand
}

// Slist is a wrapper around the C record GSList.
type Slist struct {
	native *C.GSList
	Data   uintptr
	// next : no type generator for GLib.SList, GSList*
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

// Sequence is a wrapper around the C record GSequence.
type Sequence struct {
	native *C.GSequence
}

// Sequenceiter is a wrapper around the C record GSequenceIter.
type Sequenceiter struct {
	native *C.GSequenceIter
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

// Sourcecallbackfuncs is a wrapper around the C record GSourceCallbackFuncs.
type Sourcecallbackfuncs struct {
	native *C.GSourceCallbackFuncs
	// no type for ref
	// no type for unref
	// no type for get
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

// Sourceprivate is a wrapper around the C record GSourcePrivate.
type Sourceprivate struct {
	native *C.GSourcePrivate
}

// Statbuf is a wrapper around the C record GStatBuf.
type Statbuf struct {
	native *C.GStatBuf
}

// String is a wrapper around the C record GString.
type String struct {
	native       *C.GString
	Str          string
	Len          uint64
	AllocatedLen uint64
}

// Stringchunk is a wrapper around the C record GStringChunk.
type Stringchunk struct {
	native *C.GStringChunk
}

// Testcase is a wrapper around the C record GTestCase.
type Testcase struct {
	native *C.GTestCase
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

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// Testsuite is a wrapper around the C record GTestSuite.
type Testsuite struct {
	native *C.GTestSuite
}

// Thread is a wrapper around the C record GThread.
type Thread struct {
	native *C.GThread
}

// Threadpool is a wrapper around the C record GThreadPool.
type Threadpool struct {
	native *C.GThreadPool
	// func : no type generator for Func, GFunc
	UserData uintptr
	// exclusive : no type generator for gboolean, gboolean
}

// Timeval is a wrapper around the C record GTimeVal.
type Timeval struct {
	native *C.GTimeVal
	TvSec  int64
	TvUsec int64
}

// Timer is a wrapper around the C record GTimer.
type Timer struct {
	native *C.GTimer
}

// Trashstack is a wrapper around the C record GTrashStack.
type Trashstack struct {
	native *C.GTrashStack
	// next : no type generator for TrashStack, GTrashStack*
}

// Tree is a wrapper around the C record GTree.
type Tree struct {
	native *C.GTree
}

// Variantbuilder is a wrapper around the C record GVariantBuilder.
type Variantbuilder struct {
	native *C.GVariantBuilder
}

// Variantiter is a wrapper around the C record GVariantIter.
type Variantiter struct {
	native *C.GVariantIter
	// no type for x
}

// Blacklisted : GVariantType
