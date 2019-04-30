// This is a generated file - DO NOT EDIT

package glib

import "unsafe"

// Array is a wrapper around the C record GArray.
type Array struct {
	native unsafe.Pointer
	Data   string
	Len    uint32
}

func ArrayNewFromC(u unsafe.Pointer) *Array {
	if u == nil {
		return nil
	}

	g := &Array{native: u}

	return g
}

func (recv *Array) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AsyncQueue is a wrapper around the C record GAsyncQueue.
type AsyncQueue struct {
	native unsafe.Pointer
}

func AsyncQueueNewFromC(u unsafe.Pointer) *AsyncQueue {
	if u == nil {
		return nil
	}

	g := &AsyncQueue{native: u}

	return g
}

func (recv *AsyncQueue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BookmarkFile is a wrapper around the C record GBookmarkFile.
type BookmarkFile struct {
	native unsafe.Pointer
}

func BookmarkFileNewFromC(u unsafe.Pointer) *BookmarkFile {
	if u == nil {
		return nil
	}

	g := &BookmarkFile{native: u}

	return g
}

func (recv *BookmarkFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GByteArray

// Cond is a wrapper around the C record GCond.
type Cond struct {
	native unsafe.Pointer
	// Private : p
	// Private : i
}

func CondNewFromC(u unsafe.Pointer) *Cond {
	if u == nil {
		return nil
	}

	g := &Cond{native: u}

	return g
}

func (recv *Cond) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Data is a wrapper around the C record GData.
type Data struct {
	native unsafe.Pointer
}

func DataNewFromC(u unsafe.Pointer) *Data {
	if u == nil {
		return nil
	}

	g := &Data{native: u}

	return g
}

func (recv *Data) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func DateNewFromC(u unsafe.Pointer) *Date {
	if u == nil {
		return nil
	}

	g := &Date{native: u}

	return g
}

func (recv *Date) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_date_new : return type :

// Unsupported : g_date_new_dmy : return type :

// Unsupported : g_date_new_julian : return type :

// DebugKey is a wrapper around the C record GDebugKey.
type DebugKey struct {
	native unsafe.Pointer
	Key    string
	Value  uint32
}

func DebugKeyNewFromC(u unsafe.Pointer) *DebugKey {
	if u == nil {
		return nil
	}

	g := &DebugKey{native: u}

	return g
}

func (recv *DebugKey) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dir is a wrapper around the C record GDir.
type Dir struct {
	native unsafe.Pointer
}

func DirNewFromC(u unsafe.Pointer) *Dir {
	if u == nil {
		return nil
	}

	g := &Dir{native: u}

	return g
}

func (recv *Dir) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Error is a wrapper around the C record GError.
type Error struct {
	native  unsafe.Pointer
	Domain  Quark
	Code    int32
	Message string
}

func ErrorNewFromC(u unsafe.Pointer) *Error {
	if u == nil {
		return nil
	}

	g := &Error{native: u}

	return g
}

func (recv *Error) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_error_new : return type :

// Unsupported : g_error_new_literal : return type :

// HashTable is a wrapper around the C record GHashTable.
type HashTable struct {
	native unsafe.Pointer
}

func HashTableNewFromC(u unsafe.Pointer) *HashTable {
	if u == nil {
		return nil
	}

	g := &HashTable{native: u}

	return g
}

func (recv *HashTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HashTableIterNewFromC(u unsafe.Pointer) *HashTableIter {
	if u == nil {
		return nil
	}

	g := &HashTableIter{native: u}

	return g
}

func (recv *HashTableIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HookNewFromC(u unsafe.Pointer) *Hook {
	if u == nil {
		return nil
	}

	g := &Hook{native: u}

	return g
}

func (recv *Hook) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func HookListNewFromC(u unsafe.Pointer) *HookList {
	if u == nil {
		return nil
	}

	g := &HookList{native: u}

	return g
}

func (recv *HookList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func IOChannelNewFromC(u unsafe.Pointer) *IOChannel {
	if u == nil {
		return nil
	}

	g := &IOChannel{native: u}

	return g
}

func (recv *IOChannel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_io_channel_new_file : return type :

// Unsupported : g_io_channel_unix_new : return type :

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

func IOFuncsNewFromC(u unsafe.Pointer) *IOFuncs {
	if u == nil {
		return nil
	}

	g := &IOFuncs{native: u}

	return g
}

func (recv *IOFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// KeyFile is a wrapper around the C record GKeyFile.
type KeyFile struct {
	native unsafe.Pointer
}

func KeyFileNewFromC(u unsafe.Pointer) *KeyFile {
	if u == nil {
		return nil
	}

	g := &KeyFile{native: u}

	return g
}

func (recv *KeyFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// List is a wrapper around the C record GList.
type List struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
}

func ListNewFromC(u unsafe.Pointer) *List {
	if u == nil {
		return nil
	}

	g := &List{native: u}

	return g
}

func (recv *List) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MainContext is a wrapper around the C record GMainContext.
type MainContext struct {
	native unsafe.Pointer
}

func MainContextNewFromC(u unsafe.Pointer) *MainContext {
	if u == nil {
		return nil
	}

	g := &MainContext{native: u}

	return g
}

func (recv *MainContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_main_context_new : return type :

// MainLoop is a wrapper around the C record GMainLoop.
type MainLoop struct {
	native unsafe.Pointer
}

func MainLoopNewFromC(u unsafe.Pointer) *MainLoop {
	if u == nil {
		return nil
	}

	g := &MainLoop{native: u}

	return g
}

func (recv *MainLoop) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_main_loop_new : return type :

// MappedFile is a wrapper around the C record GMappedFile.
type MappedFile struct {
	native unsafe.Pointer
}

func MappedFileNewFromC(u unsafe.Pointer) *MappedFile {
	if u == nil {
		return nil
	}

	g := &MappedFile{native: u}

	return g
}

func (recv *MappedFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MarkupParseContext is a wrapper around the C record GMarkupParseContext.
type MarkupParseContext struct {
	native unsafe.Pointer
}

func MarkupParseContextNewFromC(u unsafe.Pointer) *MarkupParseContext {
	if u == nil {
		return nil
	}

	g := &MarkupParseContext{native: u}

	return g
}

func (recv *MarkupParseContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// MarkupParser is a wrapper around the C record GMarkupParser.
type MarkupParser struct {
	native unsafe.Pointer
	// no type for start_element
	// no type for end_element
	// no type for text
	// no type for passthrough
	// no type for error
}

func MarkupParserNewFromC(u unsafe.Pointer) *MarkupParser {
	if u == nil {
		return nil
	}

	g := &MarkupParser{native: u}

	return g
}

func (recv *MarkupParser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MatchInfo is a wrapper around the C record GMatchInfo.
type MatchInfo struct {
	native unsafe.Pointer
}

func MatchInfoNewFromC(u unsafe.Pointer) *MatchInfo {
	if u == nil {
		return nil
	}

	g := &MatchInfo{native: u}

	return g
}

func (recv *MatchInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func MemVTableNewFromC(u unsafe.Pointer) *MemVTable {
	if u == nil {
		return nil
	}

	g := &MemVTable{native: u}

	return g
}

func (recv *MemVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func NodeNewFromC(u unsafe.Pointer) *Node {
	if u == nil {
		return nil
	}

	g := &Node{native: u}

	return g
}

func (recv *Node) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OptionContext is a wrapper around the C record GOptionContext.
type OptionContext struct {
	native unsafe.Pointer
}

func OptionContextNewFromC(u unsafe.Pointer) *OptionContext {
	if u == nil {
		return nil
	}

	g := &OptionContext{native: u}

	return g
}

func (recv *OptionContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func OptionEntryNewFromC(u unsafe.Pointer) *OptionEntry {
	if u == nil {
		return nil
	}

	g := &OptionEntry{native: u}

	return g
}

func (recv *OptionEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OptionGroup is a wrapper around the C record GOptionGroup.
type OptionGroup struct {
	native unsafe.Pointer
}

func OptionGroupNewFromC(u unsafe.Pointer) *OptionGroup {
	if u == nil {
		return nil
	}

	g := &OptionGroup{native: u}

	return g
}

func (recv *OptionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PatternSpec is a wrapper around the C record GPatternSpec.
type PatternSpec struct {
	native unsafe.Pointer
}

func PatternSpecNewFromC(u unsafe.Pointer) *PatternSpec {
	if u == nil {
		return nil
	}

	g := &PatternSpec{native: u}

	return g
}

func (recv *PatternSpec) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PollFD is a wrapper around the C record GPollFD.
type PollFD struct {
	native  unsafe.Pointer
	Fd      int32
	Events  uint32
	Revents uint32
}

func PollFDNewFromC(u unsafe.Pointer) *PollFD {
	if u == nil {
		return nil
	}

	g := &PollFD{native: u}

	return g
}

func (recv *PollFD) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Private is a wrapper around the C record GPrivate.
type Private struct {
	native unsafe.Pointer
	// Private : p
	// Private : notify
	// Private : future
}

func PrivateNewFromC(u unsafe.Pointer) *Private {
	if u == nil {
		return nil
	}

	g := &Private{native: u}

	return g
}

func (recv *Private) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GPtrArray

// Queue is a wrapper around the C record GQueue.
type Queue struct {
	native unsafe.Pointer
	// head : record
	// tail : record
	Length uint32
}

func QueueNewFromC(u unsafe.Pointer) *Queue {
	if u == nil {
		return nil
	}

	g := &Queue{native: u}

	return g
}

func (recv *Queue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Rand is a wrapper around the C record GRand.
type Rand struct {
	native unsafe.Pointer
}

func RandNewFromC(u unsafe.Pointer) *Rand {
	if u == nil {
		return nil
	}

	g := &Rand{native: u}

	return g
}

func (recv *Rand) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SList is a wrapper around the C record GSList.
type SList struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	// next : record
}

func SListNewFromC(u unsafe.Pointer) *SList {
	if u == nil {
		return nil
	}

	g := &SList{native: u}

	return g
}

func (recv *SList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ScannerNewFromC(u unsafe.Pointer) *Scanner {
	if u == nil {
		return nil
	}

	g := &Scanner{native: u}

	return g
}

func (recv *Scanner) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func ScannerConfigNewFromC(u unsafe.Pointer) *ScannerConfig {
	if u == nil {
		return nil
	}

	g := &ScannerConfig{native: u}

	return g
}

func (recv *ScannerConfig) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Sequence is a wrapper around the C record GSequence.
type Sequence struct {
	native unsafe.Pointer
}

func SequenceNewFromC(u unsafe.Pointer) *Sequence {
	if u == nil {
		return nil
	}

	g := &Sequence{native: u}

	return g
}

func (recv *Sequence) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SequenceIter is a wrapper around the C record GSequenceIter.
type SequenceIter struct {
	native unsafe.Pointer
}

func SequenceIterNewFromC(u unsafe.Pointer) *SequenceIter {
	if u == nil {
		return nil
	}

	g := &SequenceIter{native: u}

	return g
}

func (recv *SequenceIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SourceNewFromC(u unsafe.Pointer) *Source {
	if u == nil {
		return nil
	}

	g := &Source{native: u}

	return g
}

func (recv *Source) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_source_new : return type :

// SourceCallbackFuncs is a wrapper around the C record GSourceCallbackFuncs.
type SourceCallbackFuncs struct {
	native unsafe.Pointer
	// no type for ref
	// no type for unref
	// no type for get
}

func SourceCallbackFuncsNewFromC(u unsafe.Pointer) *SourceCallbackFuncs {
	if u == nil {
		return nil
	}

	g := &SourceCallbackFuncs{native: u}

	return g
}

func (recv *SourceCallbackFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func SourceFuncsNewFromC(u unsafe.Pointer) *SourceFuncs {
	if u == nil {
		return nil
	}

	g := &SourceFuncs{native: u}

	return g
}

func (recv *SourceFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SourcePrivate is a wrapper around the C record GSourcePrivate.
type SourcePrivate struct {
	native unsafe.Pointer
}

func SourcePrivateNewFromC(u unsafe.Pointer) *SourcePrivate {
	if u == nil {
		return nil
	}

	g := &SourcePrivate{native: u}

	return g
}

func (recv *SourcePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatBuf is a wrapper around the C record GStatBuf.
type StatBuf struct {
	native unsafe.Pointer
}

func StatBufNewFromC(u unsafe.Pointer) *StatBuf {
	if u == nil {
		return nil
	}

	g := &StatBuf{native: u}

	return g
}

func (recv *StatBuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// String is a wrapper around the C record GString.
type String struct {
	native       unsafe.Pointer
	Str          string
	Len          uint64
	AllocatedLen uint64
}

func StringNewFromC(u unsafe.Pointer) *String {
	if u == nil {
		return nil
	}

	g := &String{native: u}

	return g
}

func (recv *String) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StringChunk is a wrapper around the C record GStringChunk.
type StringChunk struct {
	native unsafe.Pointer
}

func StringChunkNewFromC(u unsafe.Pointer) *StringChunk {
	if u == nil {
		return nil
	}

	g := &StringChunk{native: u}

	return g
}

func (recv *StringChunk) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TestCase is a wrapper around the C record GTestCase.
type TestCase struct {
	native unsafe.Pointer
}

func TestCaseNewFromC(u unsafe.Pointer) *TestCase {
	if u == nil {
		return nil
	}

	g := &TestCase{native: u}

	return g
}

func (recv *TestCase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func TestConfigNewFromC(u unsafe.Pointer) *TestConfig {
	if u == nil {
		return nil
	}

	g := &TestConfig{native: u}

	return g
}

func (recv *TestConfig) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// TestSuite is a wrapper around the C record GTestSuite.
type TestSuite struct {
	native unsafe.Pointer
}

func TestSuiteNewFromC(u unsafe.Pointer) *TestSuite {
	if u == nil {
		return nil
	}

	g := &TestSuite{native: u}

	return g
}

func (recv *TestSuite) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Thread is a wrapper around the C record GThread.
type Thread struct {
	native unsafe.Pointer
}

func ThreadNewFromC(u unsafe.Pointer) *Thread {
	if u == nil {
		return nil
	}

	g := &Thread{native: u}

	return g
}

func (recv *Thread) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThreadPool is a wrapper around the C record GThreadPool.
type ThreadPool struct {
	native unsafe.Pointer
	// _func : no type generator for Func, GFunc
	// user_data : no type generator for gpointer, gpointer
	Exclusive bool
}

func ThreadPoolNewFromC(u unsafe.Pointer) *ThreadPool {
	if u == nil {
		return nil
	}

	g := &ThreadPool{native: u}

	return g
}

func (recv *ThreadPool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TimeVal is a wrapper around the C record GTimeVal.
type TimeVal struct {
	native unsafe.Pointer
	TvSec  int64
	TvUsec int64
}

func TimeValNewFromC(u unsafe.Pointer) *TimeVal {
	if u == nil {
		return nil
	}

	g := &TimeVal{native: u}

	return g
}

func (recv *TimeVal) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Timer is a wrapper around the C record GTimer.
type Timer struct {
	native unsafe.Pointer
}

func TimerNewFromC(u unsafe.Pointer) *Timer {
	if u == nil {
		return nil
	}

	g := &Timer{native: u}

	return g
}

func (recv *Timer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TrashStack is a wrapper around the C record GTrashStack.
type TrashStack struct {
	native unsafe.Pointer
	// next : record
}

func TrashStackNewFromC(u unsafe.Pointer) *TrashStack {
	if u == nil {
		return nil
	}

	g := &TrashStack{native: u}

	return g
}

func (recv *TrashStack) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Tree is a wrapper around the C record GTree.
type Tree struct {
	native unsafe.Pointer
}

func TreeNewFromC(u unsafe.Pointer) *Tree {
	if u == nil {
		return nil
	}

	g := &Tree{native: u}

	return g
}

func (recv *Tree) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VariantBuilder is a wrapper around the C record GVariantBuilder.
type VariantBuilder struct {
	native unsafe.Pointer
}

func VariantBuilderNewFromC(u unsafe.Pointer) *VariantBuilder {
	if u == nil {
		return nil
	}

	g := &VariantBuilder{native: u}

	return g
}

func (recv *VariantBuilder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VariantIter is a wrapper around the C record GVariantIter.
type VariantIter struct {
	native unsafe.Pointer
	// Private : x
}

func VariantIterNewFromC(u unsafe.Pointer) *VariantIter {
	if u == nil {
		return nil
	}

	g := &VariantIter{native: u}

	return g
}

func (recv *VariantIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VariantType is a wrapper around the C record GVariantType.
type VariantType struct {
	native unsafe.Pointer
}

func VariantTypeNewFromC(u unsafe.Pointer) *VariantType {
	if u == nil {
		return nil
	}

	g := &VariantType{native: u}

	return g
}

func (recv *VariantType) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_variant_type_new_array : return type :

// Unsupported : g_variant_type_new_dict_entry : return type :

// Unsupported : g_variant_type_new_maybe : return type :

// Unsupported : g_variant_type_new_tuple : unsupported parameter items :
