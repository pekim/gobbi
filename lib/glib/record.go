// This is a generated file - DO NOT EDIT

package glib

import "unsafe"

// Array is a wrapper around the C record GArray.
type Array struct {
	native unsafe.Pointer
	Data   string
	Len    uint32
}

// AsyncQueue is a wrapper around the C record GAsyncQueue.
type AsyncQueue struct {
	native unsafe.Pointer
}

// BookmarkFile is a wrapper around the C record GBookmarkFile.
type BookmarkFile struct {
	native unsafe.Pointer
}

// Blacklisted : GByteArray

// Cond is a wrapper around the C record GCond.
type Cond struct {
	native unsafe.Pointer
	// Private : p
	// Private : i
}

// Data is a wrapper around the C record GData.
type Data struct {
	native unsafe.Pointer
}

// Date is a wrapper around the C record GDate.
type Date struct {
	native unsafe.Pointer
	// Bitfield not supported : 32 julian_days
	// Bitfield not supported :  1 julian
	// Bitfield not supported :  1 dmy
	// Bitfield not supported :  6 day
	// Bitfield not supported :  4 month
	// Bitfield not supported : 16 year
}

// DebugKey is a wrapper around the C record GDebugKey.
type DebugKey struct {
	native unsafe.Pointer
	Key    string
	Value  uint32
}

// Dir is a wrapper around the C record GDir.
type Dir struct {
	native unsafe.Pointer
}

// Error is a wrapper around the C record GError.
type Error struct {
	native  unsafe.Pointer
	Domain  Quark
	Code    int32
	Message string
}

// HashTable is a wrapper around the C record GHashTable.
type HashTable struct {
	native unsafe.Pointer
}

// HashTableIter is a wrapper around the C record GHashTableIter.
type HashTableIter struct {
	native unsafe.Pointer
	// Private : dummy1
	// Private : dummy2
	// Private : dummy3
	// Private : dummy4
	// Private : dummy5
	// Private : dummy6
}

// Hook is a wrapper around the C record GHook.
type Hook struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
	RefCount uint32
	HookId   uint64
	Flags    uint32
	// _func : no type generator for gpointer, gpointer
	// destroy : no type generator for DestroyNotify, GDestroyNotify
}

// HookList is a wrapper around the C record GHookList.
type HookList struct {
	native unsafe.Pointer
	SeqId  uint64
	// Bitfield not supported : 16 hook_size
	// Bitfield not supported :  1 is_setup
	// hooks : record
	// dummy3 : no type generator for gpointer, gpointer
	// finalize_hook : no type generator for HookFinalizeFunc, GHookFinalizeFunc
	// no type for dummy
}

// Blacklisted : GIConv

// IOChannel is a wrapper around the C record GIOChannel.
type IOChannel struct {
	native unsafe.Pointer
	// Private : ref_count
	// Private : funcs
	// Private : encoding
	// Private : read_cd
	// Private : write_cd
	// Private : line_term
	// Private : line_term_len
	// Private : buf_size
	// Private : read_buf
	// Private : encoded_read_buf
	// Private : write_buf
	// Private : partial_write_buf
	// Private : use_buffer
	// Private : do_encode
	// Private : close_on_unref
	// Private : is_readable
	// Private : is_writeable
	// Private : is_seekable
	// Private : reserved1
	// Private : reserved2
}

// IOFuncs is a wrapper around the C record GIOFuncs.
type IOFuncs struct {
	native unsafe.Pointer
	// no type for io_read
	// no type for io_write
	// no type for io_seek
	// no type for io_close
	// no type for io_create_watch
	// no type for io_free
	// no type for io_set_flags
	// no type for io_get_flags
}

// KeyFile is a wrapper around the C record GKeyFile.
type KeyFile struct {
	native unsafe.Pointer
}

// List is a wrapper around the C record GList.
type List struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
}

// MainContext is a wrapper around the C record GMainContext.
type MainContext struct {
	native unsafe.Pointer
}

// MainLoop is a wrapper around the C record GMainLoop.
type MainLoop struct {
	native unsafe.Pointer
}

// MappedFile is a wrapper around the C record GMappedFile.
type MappedFile struct {
	native unsafe.Pointer
}

// MarkupParseContext is a wrapper around the C record GMarkupParseContext.
type MarkupParseContext struct {
	native unsafe.Pointer
}

// MarkupParser is a wrapper around the C record GMarkupParser.
type MarkupParser struct {
	native unsafe.Pointer
	// no type for start_element
	// no type for end_element
	// no type for text
	// no type for passthrough
	// no type for error
}

// MatchInfo is a wrapper around the C record GMatchInfo.
type MatchInfo struct {
	native unsafe.Pointer
}

// MemVTable is a wrapper around the C record GMemVTable.
type MemVTable struct {
	native unsafe.Pointer
	// no type for malloc
	// no type for realloc
	// no type for free
	// no type for calloc
	// no type for try_malloc
	// no type for try_realloc
}

// Node is a wrapper around the C record GNode.
type Node struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
	// parent : record
	// children : record
}

// OptionContext is a wrapper around the C record GOptionContext.
type OptionContext struct {
	native unsafe.Pointer
}

// OptionEntry is a wrapper around the C record GOptionEntry.
type OptionEntry struct {
	native    unsafe.Pointer
	LongName  string
	ShortName rune
	Flags     int32
	Arg       OptionArg
	// arg_data : no type generator for gpointer, gpointer
	Description    string
	ArgDescription string
}

// OptionGroup is a wrapper around the C record GOptionGroup.
type OptionGroup struct {
	native unsafe.Pointer
}

// PatternSpec is a wrapper around the C record GPatternSpec.
type PatternSpec struct {
	native unsafe.Pointer
}

// PollFD is a wrapper around the C record GPollFD.
type PollFD struct {
	native  unsafe.Pointer
	Fd      int32
	Events  uint32
	Revents uint32
}

// Private is a wrapper around the C record GPrivate.
type Private struct {
	native unsafe.Pointer
	// Private : p
	// Private : notify
	// Private : future
}

// Blacklisted : GPtrArray

// Queue is a wrapper around the C record GQueue.
type Queue struct {
	native unsafe.Pointer
	// head : record
	// tail : record
	Length uint32
}

// Rand is a wrapper around the C record GRand.
type Rand struct {
	native unsafe.Pointer
}

// SList is a wrapper around the C record GSList.
type SList struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	// next : record
}

// Scanner is a wrapper around the C record GScanner.
type Scanner struct {
	native unsafe.Pointer
	// user_data : no type generator for gpointer, gpointer
	MaxParseErrors uint32
	ParseErrors    uint32
	InputName      string
	// qdata : record
	// config : record
	Token TokenType
	// value : no type generator for TokenValue, GTokenValue
	Line      uint32
	Position  uint32
	NextToken TokenType
	// next_value : no type generator for TokenValue, GTokenValue
	NextLine     uint32
	NextPosition uint32
	// Private : symbol_table
	// Private : input_fd
	// Private : text
	// Private : text_end
	// Private : buffer
	// Private : scope_id
	// msg_handler : no type generator for ScannerMsgFunc, GScannerMsgFunc
}

// ScannerConfig is a wrapper around the C record GScannerConfig.
type ScannerConfig struct {
	native              unsafe.Pointer
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
	// Private : padding_dummy
}

// Sequence is a wrapper around the C record GSequence.
type Sequence struct {
	native unsafe.Pointer
}

// SequenceIter is a wrapper around the C record GSequenceIter.
type SequenceIter struct {
	native unsafe.Pointer
}

// Source is a wrapper around the C record GSource.
type Source struct {
	native unsafe.Pointer
	// Private : callback_data
	// Private : callback_funcs
	// Private : source_funcs
	// Private : ref_count
	// Private : context
	// Private : priority
	// Private : flags
	// Private : source_id
	// Private : poll_fds
	// Private : prev
	// Private : next
	// Private : name
	// Private : priv
}

// SourceCallbackFuncs is a wrapper around the C record GSourceCallbackFuncs.
type SourceCallbackFuncs struct {
	native unsafe.Pointer
	// no type for ref
	// no type for unref
	// no type for get
}

// SourceFuncs is a wrapper around the C record GSourceFuncs.
type SourceFuncs struct {
	native unsafe.Pointer
	// no type for prepare
	// no type for check
	// no type for dispatch
	// no type for finalize
	// Private : closure_callback
	// Private : closure_marshal
}

// SourcePrivate is a wrapper around the C record GSourcePrivate.
type SourcePrivate struct {
	native unsafe.Pointer
}

// StatBuf is a wrapper around the C record GStatBuf.
type StatBuf struct {
	native unsafe.Pointer
}

// String is a wrapper around the C record GString.
type String struct {
	native       unsafe.Pointer
	Str          string
	Len          uint64
	AllocatedLen uint64
}

// StringChunk is a wrapper around the C record GStringChunk.
type StringChunk struct {
	native unsafe.Pointer
}

// TestCase is a wrapper around the C record GTestCase.
type TestCase struct {
	native unsafe.Pointer
}

// TestConfig is a wrapper around the C record GTestConfig.
type TestConfig struct {
	native          unsafe.Pointer
	TestInitialized bool
	TestQuick       bool
	TestPerf        bool
	TestVerbose     bool
	TestQuiet       bool
	TestUndefined   bool
}

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// TestSuite is a wrapper around the C record GTestSuite.
type TestSuite struct {
	native unsafe.Pointer
}

// Thread is a wrapper around the C record GThread.
type Thread struct {
	native unsafe.Pointer
}

// ThreadPool is a wrapper around the C record GThreadPool.
type ThreadPool struct {
	native unsafe.Pointer
	// _func : no type generator for Func, GFunc
	// user_data : no type generator for gpointer, gpointer
	Exclusive bool
}

// TimeVal is a wrapper around the C record GTimeVal.
type TimeVal struct {
	native unsafe.Pointer
	TvSec  int64
	TvUsec int64
}

// Timer is a wrapper around the C record GTimer.
type Timer struct {
	native unsafe.Pointer
}

// TrashStack is a wrapper around the C record GTrashStack.
type TrashStack struct {
	native unsafe.Pointer
	// next : record
}

// Tree is a wrapper around the C record GTree.
type Tree struct {
	native unsafe.Pointer
}

// VariantBuilder is a wrapper around the C record GVariantBuilder.
type VariantBuilder struct {
	native unsafe.Pointer
}

// VariantIter is a wrapper around the C record GVariantIter.
type VariantIter struct {
	native unsafe.Pointer
	// Private : x
}

// VariantType is a wrapper around the C record GVariantType.
type VariantType struct {
	native unsafe.Pointer
}
