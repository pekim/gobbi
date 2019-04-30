// This is a generated file - DO NOT EDIT

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

// Array is a wrapper around the C record GArray.
type Array struct {
	native *C.GArray
	Data   string
	Len    uint32
}

func ArrayNewFromC(u unsafe.Pointer) *Array {
	c := (*C.GArray)(u)
	if c == nil {
		return nil
	}

	g := &Array{
		Data:   C.GoString(c.data),
		Len:    (uint32)(c.len),
		native: c,
	}

	return g
}

func (recv *Array) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// AsyncQueue is a wrapper around the C record GAsyncQueue.
type AsyncQueue struct {
	native *C.GAsyncQueue
}

func AsyncQueueNewFromC(u unsafe.Pointer) *AsyncQueue {
	c := (*C.GAsyncQueue)(u)
	if c == nil {
		return nil
	}

	g := &AsyncQueue{native: c}

	return g
}

func (recv *AsyncQueue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// BookmarkFile is a wrapper around the C record GBookmarkFile.
type BookmarkFile struct {
	native *C.GBookmarkFile
}

func BookmarkFileNewFromC(u unsafe.Pointer) *BookmarkFile {
	c := (*C.GBookmarkFile)(u)
	if c == nil {
		return nil
	}

	g := &BookmarkFile{native: c}

	return g
}

func (recv *BookmarkFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GByteArray

// Cond is a wrapper around the C record GCond.
type Cond struct {
	native *C.GCond
	// Private : p
	// Private : i
}

func CondNewFromC(u unsafe.Pointer) *Cond {
	c := (*C.GCond)(u)
	if c == nil {
		return nil
	}

	g := &Cond{native: c}

	return g
}

func (recv *Cond) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Data is a wrapper around the C record GData.
type Data struct {
	native *C.GData
}

func DataNewFromC(u unsafe.Pointer) *Data {
	c := (*C.GData)(u)
	if c == nil {
		return nil
	}

	g := &Data{native: c}

	return g
}

func (recv *Data) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
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

func DateNewFromC(u unsafe.Pointer) *Date {
	c := (*C.GDate)(u)
	if c == nil {
		return nil
	}

	g := &Date{native: c}

	return g
}

func (recv *Date) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DebugKey is a wrapper around the C record GDebugKey.
type DebugKey struct {
	native *C.GDebugKey
	Key    string
	Value  uint32
}

func DebugKeyNewFromC(u unsafe.Pointer) *DebugKey {
	c := (*C.GDebugKey)(u)
	if c == nil {
		return nil
	}

	g := &DebugKey{
		Key:    C.GoString(c.key),
		Value:  (uint32)(c.value),
		native: c,
	}

	return g
}

func (recv *DebugKey) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dir is a wrapper around the C record GDir.
type Dir struct {
	native *C.GDir
}

func DirNewFromC(u unsafe.Pointer) *Dir {
	c := (*C.GDir)(u)
	if c == nil {
		return nil
	}

	g := &Dir{native: c}

	return g
}

func (recv *Dir) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Error is a wrapper around the C record GError.
type Error struct {
	native  *C.GError
	Domain  Quark
	Code    int32
	Message string
}

func ErrorNewFromC(u unsafe.Pointer) *Error {
	c := (*C.GError)(u)
	if c == nil {
		return nil
	}

	g := &Error{
		Code:    (int32)(c.code),
		Domain:  (Quark)(c.domain),
		Message: C.GoString(c.message),
		native:  c,
	}

	return g
}

func (recv *Error) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HashTable is a wrapper around the C record GHashTable.
type HashTable struct {
	native *C.GHashTable
}

func HashTableNewFromC(u unsafe.Pointer) *HashTable {
	c := (*C.GHashTable)(u)
	if c == nil {
		return nil
	}

	g := &HashTable{native: c}

	return g
}

func (recv *HashTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HashTableIter is a wrapper around the C record GHashTableIter.
type HashTableIter struct {
	native *C.GHashTableIter
	// Private : dummy1
	// Private : dummy2
	// Private : dummy3
	// Private : dummy4
	// Private : dummy5
	// Private : dummy6
}

func HashTableIterNewFromC(u unsafe.Pointer) *HashTableIter {
	c := (*C.GHashTableIter)(u)
	if c == nil {
		return nil
	}

	g := &HashTableIter{native: c}

	return g
}

func (recv *HashTableIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Hook is a wrapper around the C record GHook.
type Hook struct {
	native *C.GHook
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
	c := (*C.GHook)(u)
	if c == nil {
		return nil
	}

	g := &Hook{
		Flags:    (uint32)(c.flags),
		HookId:   (uint64)(c.hook_id),
		RefCount: (uint32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *Hook) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HookList is a wrapper around the C record GHookList.
type HookList struct {
	native *C.GHookList
	SeqId  uint64
	// Bitfield not supported : 16 hook_size
	// Bitfield not supported :  1 is_setup
	// hooks : record
	// dummy3 : no type generator for gpointer, gpointer
	// finalize_hook : no type generator for HookFinalizeFunc, GHookFinalizeFunc
	// no type for dummy
}

func HookListNewFromC(u unsafe.Pointer) *HookList {
	c := (*C.GHookList)(u)
	if c == nil {
		return nil
	}

	g := &HookList{
		SeqId:  (uint64)(c.seq_id),
		native: c,
	}

	return g
}

func (recv *HookList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GIConv

// IOChannel is a wrapper around the C record GIOChannel.
type IOChannel struct {
	native *C.GIOChannel
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
	c := (*C.GIOChannel)(u)
	if c == nil {
		return nil
	}

	g := &IOChannel{native: c}

	return g
}

func (recv *IOChannel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IOFuncs is a wrapper around the C record GIOFuncs.
type IOFuncs struct {
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

func IOFuncsNewFromC(u unsafe.Pointer) *IOFuncs {
	c := (*C.GIOFuncs)(u)
	if c == nil {
		return nil
	}

	g := &IOFuncs{native: c}

	return g
}

func (recv *IOFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// KeyFile is a wrapper around the C record GKeyFile.
type KeyFile struct {
	native *C.GKeyFile
}

func KeyFileNewFromC(u unsafe.Pointer) *KeyFile {
	c := (*C.GKeyFile)(u)
	if c == nil {
		return nil
	}

	g := &KeyFile{native: c}

	return g
}

func (recv *KeyFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// List is a wrapper around the C record GList.
type List struct {
	native *C.GList
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
}

func ListNewFromC(u unsafe.Pointer) *List {
	c := (*C.GList)(u)
	if c == nil {
		return nil
	}

	g := &List{native: c}

	return g
}

func (recv *List) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MainContext is a wrapper around the C record GMainContext.
type MainContext struct {
	native *C.GMainContext
}

func MainContextNewFromC(u unsafe.Pointer) *MainContext {
	c := (*C.GMainContext)(u)
	if c == nil {
		return nil
	}

	g := &MainContext{native: c}

	return g
}

func (recv *MainContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MainLoop is a wrapper around the C record GMainLoop.
type MainLoop struct {
	native *C.GMainLoop
}

func MainLoopNewFromC(u unsafe.Pointer) *MainLoop {
	c := (*C.GMainLoop)(u)
	if c == nil {
		return nil
	}

	g := &MainLoop{native: c}

	return g
}

func (recv *MainLoop) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MappedFile is a wrapper around the C record GMappedFile.
type MappedFile struct {
	native *C.GMappedFile
}

func MappedFileNewFromC(u unsafe.Pointer) *MappedFile {
	c := (*C.GMappedFile)(u)
	if c == nil {
		return nil
	}

	g := &MappedFile{native: c}

	return g
}

func (recv *MappedFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MarkupParseContext is a wrapper around the C record GMarkupParseContext.
type MarkupParseContext struct {
	native *C.GMarkupParseContext
}

func MarkupParseContextNewFromC(u unsafe.Pointer) *MarkupParseContext {
	c := (*C.GMarkupParseContext)(u)
	if c == nil {
		return nil
	}

	g := &MarkupParseContext{native: c}

	return g
}

func (recv *MarkupParseContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MarkupParser is a wrapper around the C record GMarkupParser.
type MarkupParser struct {
	native *C.GMarkupParser
	// no type for start_element
	// no type for end_element
	// no type for text
	// no type for passthrough
	// no type for error
}

func MarkupParserNewFromC(u unsafe.Pointer) *MarkupParser {
	c := (*C.GMarkupParser)(u)
	if c == nil {
		return nil
	}

	g := &MarkupParser{native: c}

	return g
}

func (recv *MarkupParser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MatchInfo is a wrapper around the C record GMatchInfo.
type MatchInfo struct {
	native *C.GMatchInfo
}

func MatchInfoNewFromC(u unsafe.Pointer) *MatchInfo {
	c := (*C.GMatchInfo)(u)
	if c == nil {
		return nil
	}

	g := &MatchInfo{native: c}

	return g
}

func (recv *MatchInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MemVTable is a wrapper around the C record GMemVTable.
type MemVTable struct {
	native *C.GMemVTable
	// no type for malloc
	// no type for realloc
	// no type for free
	// no type for calloc
	// no type for try_malloc
	// no type for try_realloc
}

func MemVTableNewFromC(u unsafe.Pointer) *MemVTable {
	c := (*C.GMemVTable)(u)
	if c == nil {
		return nil
	}

	g := &MemVTable{native: c}

	return g
}

func (recv *MemVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Node is a wrapper around the C record GNode.
type Node struct {
	native *C.GNode
	// data : no type generator for gpointer, gpointer
	// next : record
	// prev : record
	// parent : record
	// children : record
}

func NodeNewFromC(u unsafe.Pointer) *Node {
	c := (*C.GNode)(u)
	if c == nil {
		return nil
	}

	g := &Node{native: c}

	return g
}

func (recv *Node) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OptionContext is a wrapper around the C record GOptionContext.
type OptionContext struct {
	native *C.GOptionContext
}

func OptionContextNewFromC(u unsafe.Pointer) *OptionContext {
	c := (*C.GOptionContext)(u)
	if c == nil {
		return nil
	}

	g := &OptionContext{native: c}

	return g
}

func (recv *OptionContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OptionEntry is a wrapper around the C record GOptionEntry.
type OptionEntry struct {
	native    *C.GOptionEntry
	LongName  string
	ShortName rune
	Flags     int32
	Arg       OptionArg
	// arg_data : no type generator for gpointer, gpointer
	Description    string
	ArgDescription string
}

func OptionEntryNewFromC(u unsafe.Pointer) *OptionEntry {
	c := (*C.GOptionEntry)(u)
	if c == nil {
		return nil
	}

	g := &OptionEntry{
		Arg:            (OptionArg)(c.arg),
		ArgDescription: C.GoString(c.arg_description),
		Description:    C.GoString(c.description),
		Flags:          (int32)(c.flags),
		LongName:       C.GoString(c.long_name),
		ShortName:      (rune)(c.short_name),
		native:         c,
	}

	return g
}

func (recv *OptionEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// OptionGroup is a wrapper around the C record GOptionGroup.
type OptionGroup struct {
	native *C.GOptionGroup
}

func OptionGroupNewFromC(u unsafe.Pointer) *OptionGroup {
	c := (*C.GOptionGroup)(u)
	if c == nil {
		return nil
	}

	g := &OptionGroup{native: c}

	return g
}

func (recv *OptionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PatternSpec is a wrapper around the C record GPatternSpec.
type PatternSpec struct {
	native *C.GPatternSpec
}

func PatternSpecNewFromC(u unsafe.Pointer) *PatternSpec {
	c := (*C.GPatternSpec)(u)
	if c == nil {
		return nil
	}

	g := &PatternSpec{native: c}

	return g
}

func (recv *PatternSpec) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PollFD is a wrapper around the C record GPollFD.
type PollFD struct {
	native  *C.GPollFD
	Fd      int32
	Events  uint32
	Revents uint32
}

func PollFDNewFromC(u unsafe.Pointer) *PollFD {
	c := (*C.GPollFD)(u)
	if c == nil {
		return nil
	}

	g := &PollFD{
		Events:  (uint32)(c.events),
		Fd:      (int32)(c.fd),
		Revents: (uint32)(c.revents),
		native:  c,
	}

	return g
}

func (recv *PollFD) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Private is a wrapper around the C record GPrivate.
type Private struct {
	native *C.GPrivate
	// Private : p
	// Private : notify
	// Private : future
}

func PrivateNewFromC(u unsafe.Pointer) *Private {
	c := (*C.GPrivate)(u)
	if c == nil {
		return nil
	}

	g := &Private{native: c}

	return g
}

func (recv *Private) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GPtrArray

// Queue is a wrapper around the C record GQueue.
type Queue struct {
	native *C.GQueue
	// head : record
	// tail : record
	Length uint32
}

func QueueNewFromC(u unsafe.Pointer) *Queue {
	c := (*C.GQueue)(u)
	if c == nil {
		return nil
	}

	g := &Queue{
		Length: (uint32)(c.length),
		native: c,
	}

	return g
}

func (recv *Queue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Rand is a wrapper around the C record GRand.
type Rand struct {
	native *C.GRand
}

func RandNewFromC(u unsafe.Pointer) *Rand {
	c := (*C.GRand)(u)
	if c == nil {
		return nil
	}

	g := &Rand{native: c}

	return g
}

func (recv *Rand) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SList is a wrapper around the C record GSList.
type SList struct {
	native *C.GSList
	// data : no type generator for gpointer, gpointer
	// next : record
}

func SListNewFromC(u unsafe.Pointer) *SList {
	c := (*C.GSList)(u)
	if c == nil {
		return nil
	}

	g := &SList{native: c}

	return g
}

func (recv *SList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Scanner is a wrapper around the C record GScanner.
type Scanner struct {
	native *C.GScanner
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
	c := (*C.GScanner)(u)
	if c == nil {
		return nil
	}

	g := &Scanner{
		InputName:      C.GoString(c.input_name),
		Line:           (uint32)(c.line),
		MaxParseErrors: (uint32)(c.max_parse_errors),
		NextLine:       (uint32)(c.next_line),
		NextPosition:   (uint32)(c.next_position),
		NextToken:      (TokenType)(c.next_token),
		ParseErrors:    (uint32)(c.parse_errors),
		Position:       (uint32)(c.position),
		Token:          (TokenType)(c.token),
		native:         c,
	}

	return g
}

func (recv *Scanner) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScannerConfig is a wrapper around the C record GScannerConfig.
type ScannerConfig struct {
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
	// Private : padding_dummy
}

func ScannerConfigNewFromC(u unsafe.Pointer) *ScannerConfig {
	c := (*C.GScannerConfig)(u)
	if c == nil {
		return nil
	}

	g := &ScannerConfig{
		CpairCommentSingle:  C.GoString(c.cpair_comment_single),
		CsetIdentifierFirst: C.GoString(c.cset_identifier_first),
		CsetIdentifierNth:   C.GoString(c.cset_identifier_nth),
		CsetSkipCharacters:  C.GoString(c.cset_skip_characters),
		native:              c,
	}

	return g
}

func (recv *ScannerConfig) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Sequence is a wrapper around the C record GSequence.
type Sequence struct {
	native *C.GSequence
}

func SequenceNewFromC(u unsafe.Pointer) *Sequence {
	c := (*C.GSequence)(u)
	if c == nil {
		return nil
	}

	g := &Sequence{native: c}

	return g
}

func (recv *Sequence) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SequenceIter is a wrapper around the C record GSequenceIter.
type SequenceIter struct {
	native *C.GSequenceIter
}

func SequenceIterNewFromC(u unsafe.Pointer) *SequenceIter {
	c := (*C.GSequenceIter)(u)
	if c == nil {
		return nil
	}

	g := &SequenceIter{native: c}

	return g
}

func (recv *SequenceIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Source is a wrapper around the C record GSource.
type Source struct {
	native *C.GSource
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
	c := (*C.GSource)(u)
	if c == nil {
		return nil
	}

	g := &Source{native: c}

	return g
}

func (recv *Source) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SourceCallbackFuncs is a wrapper around the C record GSourceCallbackFuncs.
type SourceCallbackFuncs struct {
	native *C.GSourceCallbackFuncs
	// no type for ref
	// no type for unref
	// no type for get
}

func SourceCallbackFuncsNewFromC(u unsafe.Pointer) *SourceCallbackFuncs {
	c := (*C.GSourceCallbackFuncs)(u)
	if c == nil {
		return nil
	}

	g := &SourceCallbackFuncs{native: c}

	return g
}

func (recv *SourceCallbackFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SourceFuncs is a wrapper around the C record GSourceFuncs.
type SourceFuncs struct {
	native *C.GSourceFuncs
	// no type for prepare
	// no type for check
	// no type for dispatch
	// no type for finalize
	// Private : closure_callback
	// Private : closure_marshal
}

func SourceFuncsNewFromC(u unsafe.Pointer) *SourceFuncs {
	c := (*C.GSourceFuncs)(u)
	if c == nil {
		return nil
	}

	g := &SourceFuncs{native: c}

	return g
}

func (recv *SourceFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SourcePrivate is a wrapper around the C record GSourcePrivate.
type SourcePrivate struct {
	native *C.GSourcePrivate
}

func SourcePrivateNewFromC(u unsafe.Pointer) *SourcePrivate {
	c := (*C.GSourcePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SourcePrivate{native: c}

	return g
}

func (recv *SourcePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StatBuf is a wrapper around the C record GStatBuf.
type StatBuf struct {
	native *C.GStatBuf
}

func StatBufNewFromC(u unsafe.Pointer) *StatBuf {
	c := (*C.GStatBuf)(u)
	if c == nil {
		return nil
	}

	g := &StatBuf{native: c}

	return g
}

func (recv *StatBuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// String is a wrapper around the C record GString.
type String struct {
	native       *C.GString
	Str          string
	Len          uint64
	AllocatedLen uint64
}

func StringNewFromC(u unsafe.Pointer) *String {
	c := (*C.GString)(u)
	if c == nil {
		return nil
	}

	g := &String{
		AllocatedLen: (uint64)(c.allocated_len),
		Len:          (uint64)(c.len),
		Str:          C.GoString(c.str),
		native:       c,
	}

	return g
}

func (recv *String) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StringChunk is a wrapper around the C record GStringChunk.
type StringChunk struct {
	native *C.GStringChunk
}

func StringChunkNewFromC(u unsafe.Pointer) *StringChunk {
	c := (*C.GStringChunk)(u)
	if c == nil {
		return nil
	}

	g := &StringChunk{native: c}

	return g
}

func (recv *StringChunk) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TestCase is a wrapper around the C record GTestCase.
type TestCase struct {
	native *C.GTestCase
}

func TestCaseNewFromC(u unsafe.Pointer) *TestCase {
	c := (*C.GTestCase)(u)
	if c == nil {
		return nil
	}

	g := &TestCase{native: c}

	return g
}

func (recv *TestCase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TestConfig is a wrapper around the C record GTestConfig.
type TestConfig struct {
	native          *C.GTestConfig
	TestInitialized bool
	TestQuick       bool
	TestPerf        bool
	TestVerbose     bool
	TestQuiet       bool
	TestUndefined   bool
}

func TestConfigNewFromC(u unsafe.Pointer) *TestConfig {
	c := (*C.GTestConfig)(u)
	if c == nil {
		return nil
	}

	g := &TestConfig{
		TestInitialized: c.test_initialized == C.TRUE,
		TestPerf:        c.test_perf == C.TRUE,
		TestQuick:       c.test_quick == C.TRUE,
		TestQuiet:       c.test_quiet == C.TRUE,
		TestUndefined:   c.test_undefined == C.TRUE,
		TestVerbose:     c.test_verbose == C.TRUE,
		native:          c,
	}

	return g
}

func (recv *TestConfig) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// TestSuite is a wrapper around the C record GTestSuite.
type TestSuite struct {
	native *C.GTestSuite
}

func TestSuiteNewFromC(u unsafe.Pointer) *TestSuite {
	c := (*C.GTestSuite)(u)
	if c == nil {
		return nil
	}

	g := &TestSuite{native: c}

	return g
}

func (recv *TestSuite) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Thread is a wrapper around the C record GThread.
type Thread struct {
	native *C.GThread
}

func ThreadNewFromC(u unsafe.Pointer) *Thread {
	c := (*C.GThread)(u)
	if c == nil {
		return nil
	}

	g := &Thread{native: c}

	return g
}

func (recv *Thread) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ThreadPool is a wrapper around the C record GThreadPool.
type ThreadPool struct {
	native *C.GThreadPool
	// _func : no type generator for Func, GFunc
	// user_data : no type generator for gpointer, gpointer
	Exclusive bool
}

func ThreadPoolNewFromC(u unsafe.Pointer) *ThreadPool {
	c := (*C.GThreadPool)(u)
	if c == nil {
		return nil
	}

	g := &ThreadPool{
		Exclusive: c.exclusive == C.TRUE,
		native:    c,
	}

	return g
}

func (recv *ThreadPool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TimeVal is a wrapper around the C record GTimeVal.
type TimeVal struct {
	native *C.GTimeVal
	TvSec  int64
	TvUsec int64
}

func TimeValNewFromC(u unsafe.Pointer) *TimeVal {
	c := (*C.GTimeVal)(u)
	if c == nil {
		return nil
	}

	g := &TimeVal{
		TvSec:  (int64)(c.tv_sec),
		TvUsec: (int64)(c.tv_usec),
		native: c,
	}

	return g
}

func (recv *TimeVal) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Timer is a wrapper around the C record GTimer.
type Timer struct {
	native *C.GTimer
}

func TimerNewFromC(u unsafe.Pointer) *Timer {
	c := (*C.GTimer)(u)
	if c == nil {
		return nil
	}

	g := &Timer{native: c}

	return g
}

func (recv *Timer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TrashStack is a wrapper around the C record GTrashStack.
type TrashStack struct {
	native *C.GTrashStack
	// next : record
}

func TrashStackNewFromC(u unsafe.Pointer) *TrashStack {
	c := (*C.GTrashStack)(u)
	if c == nil {
		return nil
	}

	g := &TrashStack{native: c}

	return g
}

func (recv *TrashStack) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Tree is a wrapper around the C record GTree.
type Tree struct {
	native *C.GTree
}

func TreeNewFromC(u unsafe.Pointer) *Tree {
	c := (*C.GTree)(u)
	if c == nil {
		return nil
	}

	g := &Tree{native: c}

	return g
}

func (recv *Tree) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VariantBuilder is a wrapper around the C record GVariantBuilder.
type VariantBuilder struct {
	native *C.GVariantBuilder
}

func VariantBuilderNewFromC(u unsafe.Pointer) *VariantBuilder {
	c := (*C.GVariantBuilder)(u)
	if c == nil {
		return nil
	}

	g := &VariantBuilder{native: c}

	return g
}

func (recv *VariantBuilder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VariantIter is a wrapper around the C record GVariantIter.
type VariantIter struct {
	native *C.GVariantIter
	// Private : x
}

func VariantIterNewFromC(u unsafe.Pointer) *VariantIter {
	c := (*C.GVariantIter)(u)
	if c == nil {
		return nil
	}

	g := &VariantIter{native: c}

	return g
}

func (recv *VariantIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// VariantType is a wrapper around the C record GVariantType.
type VariantType struct {
	native *C.GVariantType
}

func VariantTypeNewFromC(u unsafe.Pointer) *VariantType {
	c := (*C.GVariantType)(u)
	if c == nil {
		return nil
	}

	g := &VariantType{native: c}

	return g
}

func (recv *VariantType) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
