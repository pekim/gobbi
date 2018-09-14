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
	Data   int
	Len    int
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
	Data   int
	Len    int
}

// Cond is a wrapper around the C record GCond.
type Cond struct {
	native *C.GCond
	P      int
	I      int
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
	Key    int
	Value  int
}

// Dir is a wrapper around the C record GDir.
type Dir struct {
	native *C.GDir
}

// Error is a wrapper around the C record GError.
type Error struct {
	native  *C.GError
	Domain  int
	Code    int
	Message int
}

// Hashtable is a wrapper around the C record GHashTable.
type Hashtable struct {
	native *C.GHashTable
}

// Hashtableiter is a wrapper around the C record GHashTableIter.
type Hashtableiter struct {
	native *C.GHashTableIter
	Dummy1 int
	Dummy2 int
	Dummy3 int
	Dummy4 int
	Dummy5 int
	Dummy6 int
}

// Hook is a wrapper around the C record GHook.
type Hook struct {
	native   *C.GHook
	Data     int
	Next     int
	Prev     int
	RefCount int
	HookId   int
	Flags    int
	Func     int
	Destroy  int
}

// Hooklist is a wrapper around the C record GHookList.
type Hooklist struct {
	native *C.GHookList
	SeqId  int
	// Bitfield not supported : 16 hook_size
	// Bitfield not supported :  1 is_setup
	Hooks        int
	Dummy3       int
	FinalizeHook int
	Dummy        int
}

// Blacklisted : GIConv

// Iochannel is a wrapper around the C record GIOChannel.
type Iochannel struct {
	native          *C.GIOChannel
	RefCount        int
	Funcs           int
	Encoding        int
	ReadCd          int
	WriteCd         int
	LineTerm        int
	LineTermLen     int
	BufSize         int
	ReadBuf         int
	EncodedReadBuf  int
	WriteBuf        int
	PartialWriteBuf int
	// Bitfield not supported :  1 use_buffer
	// Bitfield not supported :  1 do_encode
	// Bitfield not supported :  1 close_on_unref
	// Bitfield not supported :  1 is_readable
	// Bitfield not supported :  1 is_writeable
	// Bitfield not supported :  1 is_seekable
	Reserved1 int
	Reserved2 int
}

// Iofuncs is a wrapper around the C record GIOFuncs.
type Iofuncs struct {
	native        *C.GIOFuncs
	IoRead        int
	IoWrite       int
	IoSeek        int
	IoClose       int
	IoCreateWatch int
	IoFree        int
	IoSetFlags    int
	IoGetFlags    int
}

// Keyfile is a wrapper around the C record GKeyFile.
type Keyfile struct {
	native *C.GKeyFile
}

// List is a wrapper around the C record GList.
type List struct {
	native *C.GList
	Data   int
	Next   int
	Prev   int
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
	native       *C.GMarkupParser
	StartElement int
	EndElement   int
	Text         int
	Passthrough  int
	Error        int
}

// Matchinfo is a wrapper around the C record GMatchInfo.
type Matchinfo struct {
	native *C.GMatchInfo
}

// Memvtable is a wrapper around the C record GMemVTable.
type Memvtable struct {
	native     *C.GMemVTable
	Malloc     int
	Realloc    int
	Free       int
	Calloc     int
	TryMalloc  int
	TryRealloc int
}

// Node is a wrapper around the C record GNode.
type Node struct {
	native   *C.GNode
	Data     int
	Next     int
	Prev     int
	Parent   int
	Children int
}

// Optioncontext is a wrapper around the C record GOptionContext.
type Optioncontext struct {
	native *C.GOptionContext
}

// Optionentry is a wrapper around the C record GOptionEntry.
type Optionentry struct {
	native         *C.GOptionEntry
	LongName       int
	ShortName      int
	Flags          int
	Arg            int
	ArgData        int
	Description    int
	ArgDescription int
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
	Fd      int
	Events  int
	Revents int
}

// Private is a wrapper around the C record GPrivate.
type Private struct {
	native *C.GPrivate
	P      int
	Notify int
	Future int
}

// Ptrarray is a wrapper around the C record GPtrArray.
type Ptrarray struct {
	native *C.GPtrArray
	Pdata  int
	Len    int
}

// Queue is a wrapper around the C record GQueue.
type Queue struct {
	native *C.GQueue
	Head   int
	Tail   int
	Length int
}

// Rand is a wrapper around the C record GRand.
type Rand struct {
	native *C.GRand
}

// Slist is a wrapper around the C record GSList.
type Slist struct {
	native *C.GSList
	Data   int
	Next   int
}

// Scanner is a wrapper around the C record GScanner.
type Scanner struct {
	native         *C.GScanner
	UserData       int
	MaxParseErrors int
	ParseErrors    int
	InputName      int
	Qdata          int
	Config         int
	Token          int
	Value          int
	Line           int
	Position       int
	NextToken      int
	NextValue      int
	NextLine       int
	NextPosition   int
	SymbolTable    int
	InputFd        int
	Text           int
	TextEnd        int
	Buffer         int
	ScopeId        int
	MsgHandler     int
}

// Scannerconfig is a wrapper around the C record GScannerConfig.
type Scannerconfig struct {
	native              *C.GScannerConfig
	CsetSkipCharacters  int
	CsetIdentifierFirst int
	CsetIdentifierNth   int
	CpairCommentSingle  int
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
	PaddingDummy int
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
	native        *C.GSource
	CallbackData  int
	CallbackFuncs int
	SourceFuncs   int
	RefCount      int
	Context       int
	Priority      int
	Flags         int
	SourceId      int
	PollFds       int
	Prev          int
	Next          int
	Name          int
	Priv          int
}

// Sourcecallbackfuncs is a wrapper around the C record GSourceCallbackFuncs.
type Sourcecallbackfuncs struct {
	native *C.GSourceCallbackFuncs
	Ref    int
	Unref  int
	Get    int
}

// Sourcefuncs is a wrapper around the C record GSourceFuncs.
type Sourcefuncs struct {
	native          *C.GSourceFuncs
	Prepare         int
	Check           int
	Dispatch        int
	Finalize        int
	ClosureCallback int
	ClosureMarshal  int
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
	Str          int
	Len          int
	AllocatedLen int
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
	native          *C.GTestConfig
	TestInitialized int
	TestQuick       int
	TestPerf        int
	TestVerbose     int
	TestQuiet       int
	TestUndefined   int
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
	native    *C.GThreadPool
	Func      int
	UserData  int
	Exclusive int
}

// Timeval is a wrapper around the C record GTimeVal.
type Timeval struct {
	native *C.GTimeVal
	TvSec  int
	TvUsec int
}

// Timer is a wrapper around the C record GTimer.
type Timer struct {
	native *C.GTimer
}

// Trashstack is a wrapper around the C record GTrashStack.
type Trashstack struct {
	native *C.GTrashStack
	Next   int
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
	X      int
}

// Blacklisted : GVariantType
