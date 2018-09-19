// This is a generated file - DO NOT EDIT

package glib

import "unsafe"

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

	g := &Array{
		Data:   C.GoString(c.data),
		Len:    (uint32)(c.len),
		native: c,
	}

	return g
}

// AsyncQueue is a wrapper around the C record GAsyncQueue.
type AsyncQueue struct {
	native *C.GAsyncQueue
}

func asyncQueueNewFromC(c *C.GAsyncQueue) *AsyncQueue {
	if c == nil {
		return nil
	}

	g := &AsyncQueue{native: c}

	return g
}

// BookmarkFile is a wrapper around the C record GBookmarkFile.
type BookmarkFile struct {
	native *C.GBookmarkFile
}

func bookmarkFileNewFromC(c *C.GBookmarkFile) *BookmarkFile {
	if c == nil {
		return nil
	}

	g := &BookmarkFile{native: c}

	return g
}

// ByteArray is a wrapper around the C record GByteArray.
type ByteArray struct {
	native *C.GByteArray
	// data : no type generator for guint8, guint8*
	Len uint32
}

func byteArrayNewFromC(c *C.GByteArray) *ByteArray {
	if c == nil {
		return nil
	}

	g := &ByteArray{
		Len:    (uint32)(c.len),
		native: c,
	}

	return g
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

	g := &Cond{
		P:      (uintptr)(c.p),
		native: c,
	}

	return g
}

// Data is a wrapper around the C record GData.
type Data struct {
	native *C.GData
}

func dataNewFromC(c *C.GData) *Data {
	if c == nil {
		return nil
	}

	g := &Data{native: c}

	return g
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

	g := &Date{native: c}

	return g
}

// DateNew is a wrapper around the C function g_date_new.
func DateNew() *Date {
	retC := C.g_date_new()
	retGo := dateNewFromC(retC)

	return retGo
}

// DateNewDmy is a wrapper around the C function g_date_new_dmy.
func DateNewDmy(day DateDay, month DateMonth, year DateYear) *Date {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_new_dmy(c_day, c_month, c_year)
	retGo := dateNewFromC(retC)

	return retGo
}

// DateNewJulian is a wrapper around the C function g_date_new_julian.
func DateNewJulian(julianDay uint32) *Date {
	c_julian_day := (C.guint32)(julianDay)

	retC := C.g_date_new_julian(c_julian_day)
	retGo := dateNewFromC(retC)

	return retGo
}

// DebugKey is a wrapper around the C record GDebugKey.
type DebugKey struct {
	native *C.GDebugKey
	Key    string
	Value  uint32
}

func debugKeyNewFromC(c *C.GDebugKey) *DebugKey {
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

// Dir is a wrapper around the C record GDir.
type Dir struct {
	native *C.GDir
}

func dirNewFromC(c *C.GDir) *Dir {
	if c == nil {
		return nil
	}

	g := &Dir{native: c}

	return g
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

	g := &Error{
		Code:    (int32)(c.code),
		Domain:  (Quark)(c.domain),
		Message: C.GoString(c.message),
		native:  c,
	}

	return g
}

// Unsupported : g_error_new : unsupported parameter ... : varargs

// ErrorNewLiteral is a wrapper around the C function g_error_new_literal.
func ErrorNewLiteral(domain Quark, code int32, message string) *Error {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	retC := C.g_error_new_literal(c_domain, c_code, c_message)
	retGo := errorNewFromC(retC)

	return retGo
}

// Unsupported : g_error_new_valist : unsupported parameter args : no type generator for va_list, va_list

// HashTable is a wrapper around the C record GHashTable.
type HashTable struct {
	native *C.GHashTable
}

func hashTableNewFromC(c *C.GHashTable) *HashTable {
	if c == nil {
		return nil
	}

	g := &HashTable{native: c}

	return g
}

// HashTableIter is a wrapper around the C record GHashTableIter.
type HashTableIter struct {
	native *C.GHashTableIter
	Dummy1 uintptr
	Dummy2 uintptr
	Dummy3 uintptr
	Dummy4 int32
	// dummy5 : no type generator for gboolean, gboolean
	Dummy6 uintptr
}

func hashTableIterNewFromC(c *C.GHashTableIter) *HashTableIter {
	if c == nil {
		return nil
	}

	g := &HashTableIter{
		Dummy1: (uintptr)(c.dummy1),
		Dummy2: (uintptr)(c.dummy2),
		Dummy3: (uintptr)(c.dummy3),
		Dummy4: (int32)(c.dummy4),
		Dummy6: (uintptr)(c.dummy6),
		native: c,
	}

	return g
}

// Hook is a wrapper around the C record GHook.
type Hook struct {
	native   *C.GHook
	Data     uintptr
	Next     *Hook
	Prev     *Hook
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

	g := &Hook{
		Data:     (uintptr)(c.data),
		Flags:    (uint32)(c.flags),
		Func:     (uintptr)(c._func),
		HookId:   (uint64)(c.hook_id),
		Next:     hookNewFromC(c.next),
		Prev:     hookNewFromC(c.prev),
		RefCount: (uint32)(c.ref_count),
		native:   c,
	}

	return g
}

// HookList is a wrapper around the C record GHookList.
type HookList struct {
	native *C.GHookList
	SeqId  uint64
	// Bitfield not supported : 16 hook_size
	// Bitfield not supported :  1 is_setup
	Hooks  *Hook
	Dummy3 uintptr
	// finalize_hook : no type generator for HookFinalizeFunc, GHookFinalizeFunc
	// no type for dummy
}

func hookListNewFromC(c *C.GHookList) *HookList {
	if c == nil {
		return nil
	}

	g := &HookList{
		Dummy3: (uintptr)(c.dummy3),
		Hooks:  hookNewFromC(c.hooks),
		SeqId:  (uint64)(c.seq_id),
		native: c,
	}

	return g
}

// Blacklisted : GIConv

// Blacklisted : GIOChannel

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

func iOFuncsNewFromC(c *C.GIOFuncs) *IOFuncs {
	if c == nil {
		return nil
	}

	g := &IOFuncs{native: c}

	return g
}

// KeyFile is a wrapper around the C record GKeyFile.
type KeyFile struct {
	native *C.GKeyFile
}

func keyFileNewFromC(c *C.GKeyFile) *KeyFile {
	if c == nil {
		return nil
	}

	g := &KeyFile{native: c}

	return g
}

// KeyFileNew is a wrapper around the C function g_key_file_new.
func KeyFileNew() *KeyFile {
	retC := C.g_key_file_new()
	retGo := keyFileNewFromC(retC)

	return retGo
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

	g := &List{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

// MainContext is a wrapper around the C record GMainContext.
type MainContext struct {
	native *C.GMainContext
}

func mainContextNewFromC(c *C.GMainContext) *MainContext {
	if c == nil {
		return nil
	}

	g := &MainContext{native: c}

	return g
}

// MainContextNew is a wrapper around the C function g_main_context_new.
func MainContextNew() *MainContext {
	retC := C.g_main_context_new()
	retGo := mainContextNewFromC(retC)

	return retGo
}

// MainLoop is a wrapper around the C record GMainLoop.
type MainLoop struct {
	native *C.GMainLoop
}

func mainLoopNewFromC(c *C.GMainLoop) *MainLoop {
	if c == nil {
		return nil
	}

	g := &MainLoop{native: c}

	return g
}

// Unsupported : g_main_loop_new : unsupported parameter context : record param - coming soon

// MappedFile is a wrapper around the C record GMappedFile.
type MappedFile struct {
	native *C.GMappedFile
}

func mappedFileNewFromC(c *C.GMappedFile) *MappedFile {
	if c == nil {
		return nil
	}

	g := &MappedFile{native: c}

	return g
}

// Unsupported : g_mapped_file_new : unsupported parameter writable : no type generator for gboolean, gboolean

// Unsupported : g_mapped_file_new_from_fd : unsupported parameter writable : no type generator for gboolean, gboolean

// MarkupParseContext is a wrapper around the C record GMarkupParseContext.
type MarkupParseContext struct {
	native *C.GMarkupParseContext
}

func markupParseContextNewFromC(c *C.GMarkupParseContext) *MarkupParseContext {
	if c == nil {
		return nil
	}

	g := &MarkupParseContext{native: c}

	return g
}

// Unsupported : g_markup_parse_context_new : unsupported parameter parser : record param - coming soon

// MarkupParser is a wrapper around the C record GMarkupParser.
type MarkupParser struct {
	native *C.GMarkupParser
	// no type for start_element
	// no type for end_element
	// no type for text
	// no type for passthrough
	// no type for error
}

func markupParserNewFromC(c *C.GMarkupParser) *MarkupParser {
	if c == nil {
		return nil
	}

	g := &MarkupParser{native: c}

	return g
}

// MatchInfo is a wrapper around the C record GMatchInfo.
type MatchInfo struct {
	native *C.GMatchInfo
}

func matchInfoNewFromC(c *C.GMatchInfo) *MatchInfo {
	if c == nil {
		return nil
	}

	g := &MatchInfo{native: c}

	return g
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

func memVTableNewFromC(c *C.GMemVTable) *MemVTable {
	if c == nil {
		return nil
	}

	g := &MemVTable{native: c}

	return g
}

// Node is a wrapper around the C record GNode.
type Node struct {
	native   *C.GNode
	Data     uintptr
	Next     *Node
	Prev     *Node
	Parent   *Node
	Children *Node
}

func nodeNewFromC(c *C.GNode) *Node {
	if c == nil {
		return nil
	}

	g := &Node{
		Children: nodeNewFromC(c.children),
		Data:     (uintptr)(c.data),
		Next:     nodeNewFromC(c.next),
		Parent:   nodeNewFromC(c.parent),
		Prev:     nodeNewFromC(c.prev),
		native:   c,
	}

	return g
}

// OptionContext is a wrapper around the C record GOptionContext.
type OptionContext struct {
	native *C.GOptionContext
}

func optionContextNewFromC(c *C.GOptionContext) *OptionContext {
	if c == nil {
		return nil
	}

	g := &OptionContext{native: c}

	return g
}

// OptionEntry is a wrapper around the C record GOptionEntry.
type OptionEntry struct {
	native         *C.GOptionEntry
	LongName       string
	ShortName      rune
	Flags          int32
	Arg            OptionArg
	ArgData        uintptr
	Description    string
	ArgDescription string
}

func optionEntryNewFromC(c *C.GOptionEntry) *OptionEntry {
	if c == nil {
		return nil
	}

	g := &OptionEntry{
		Arg:            (OptionArg)(c.arg),
		ArgData:        (uintptr)(c.arg_data),
		ArgDescription: C.GoString(c.arg_description),
		Description:    C.GoString(c.description),
		Flags:          (int32)(c.flags),
		LongName:       C.GoString(c.long_name),
		ShortName:      (rune)(c.short_name),
		native:         c,
	}

	return g
}

// OptionGroup is a wrapper around the C record GOptionGroup.
type OptionGroup struct {
	native *C.GOptionGroup
}

func optionGroupNewFromC(c *C.GOptionGroup) *OptionGroup {
	if c == nil {
		return nil
	}

	g := &OptionGroup{native: c}

	return g
}

// Unsupported : g_option_group_new : unsupported parameter destroy : no type generator for DestroyNotify, GDestroyNotify

// PatternSpec is a wrapper around the C record GPatternSpec.
type PatternSpec struct {
	native *C.GPatternSpec
}

func patternSpecNewFromC(c *C.GPatternSpec) *PatternSpec {
	if c == nil {
		return nil
	}

	g := &PatternSpec{native: c}

	return g
}

// PollFD is a wrapper around the C record GPollFD.
type PollFD struct {
	native  *C.GPollFD
	Fd      int32
	Events  uint32
	Revents uint32
}

func pollFDNewFromC(c *C.GPollFD) *PollFD {
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

	g := &Private{
		P:      (uintptr)(c.p),
		native: c,
	}

	return g
}

// PtrArray is a wrapper around the C record GPtrArray.
type PtrArray struct {
	native *C.GPtrArray
	// pdata : no type generator for gpointer, gpointer*
	Len uint32
}

func ptrArrayNewFromC(c *C.GPtrArray) *PtrArray {
	if c == nil {
		return nil
	}

	g := &PtrArray{
		Len:    (uint32)(c.len),
		native: c,
	}

	return g
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

	g := &Queue{
		Length: (uint32)(c.length),
		native: c,
	}

	return g
}

// Rand is a wrapper around the C record GRand.
type Rand struct {
	native *C.GRand
}

func randNewFromC(c *C.GRand) *Rand {
	if c == nil {
		return nil
	}

	g := &Rand{native: c}

	return g
}

// SList is a wrapper around the C record GSList.
type SList struct {
	native *C.GSList
	Data   uintptr
	// next : no type generator for GLib.SList, GSList*
}

func sListNewFromC(c *C.GSList) *SList {
	if c == nil {
		return nil
	}

	g := &SList{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

// Scanner is a wrapper around the C record GScanner.
type Scanner struct {
	native         *C.GScanner
	UserData       uintptr
	MaxParseErrors uint32
	ParseErrors    uint32
	InputName      string
	Qdata          *Data
	Config         *ScannerConfig
	Token          TokenType
	// value : no type generator for TokenValue, GTokenValue
	Line      uint32
	Position  uint32
	NextToken TokenType
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

	g := &Scanner{
		Buffer:         C.GoString(c.buffer),
		Config:         scannerConfigNewFromC(c.config),
		InputFd:        (int32)(c.input_fd),
		InputName:      C.GoString(c.input_name),
		Line:           (uint32)(c.line),
		MaxParseErrors: (uint32)(c.max_parse_errors),
		NextLine:       (uint32)(c.next_line),
		NextPosition:   (uint32)(c.next_position),
		NextToken:      (TokenType)(c.next_token),
		ParseErrors:    (uint32)(c.parse_errors),
		Position:       (uint32)(c.position),
		Qdata:          dataNewFromC(c.qdata),
		ScopeId:        (uint32)(c.scope_id),
		Text:           C.GoString(c.text),
		TextEnd:        C.GoString(c.text_end),
		Token:          (TokenType)(c.token),
		UserData:       (uintptr)(c.user_data),
		native:         c,
	}

	return g
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
	PaddingDummy uint32
}

func scannerConfigNewFromC(c *C.GScannerConfig) *ScannerConfig {
	if c == nil {
		return nil
	}

	g := &ScannerConfig{
		CpairCommentSingle:  C.GoString(c.cpair_comment_single),
		CsetIdentifierFirst: C.GoString(c.cset_identifier_first),
		CsetIdentifierNth:   C.GoString(c.cset_identifier_nth),
		CsetSkipCharacters:  C.GoString(c.cset_skip_characters),
		PaddingDummy:        (uint32)(c.padding_dummy),
		native:              c,
	}

	return g
}

// Sequence is a wrapper around the C record GSequence.
type Sequence struct {
	native *C.GSequence
}

func sequenceNewFromC(c *C.GSequence) *Sequence {
	if c == nil {
		return nil
	}

	g := &Sequence{native: c}

	return g
}

// SequenceIter is a wrapper around the C record GSequenceIter.
type SequenceIter struct {
	native *C.GSequenceIter
}

func sequenceIterNewFromC(c *C.GSequenceIter) *SequenceIter {
	if c == nil {
		return nil
	}

	g := &SequenceIter{native: c}

	return g
}

// Source is a wrapper around the C record GSource.
type Source struct {
	native        *C.GSource
	CallbackData  uintptr
	CallbackFuncs *SourceCallbackFuncs
	SourceFuncs   *SourceFuncs
	RefCount      uint32
	Context       *MainContext
	Priority      int32
	Flags         uint32
	SourceId      uint32
	// poll_fds : no type generator for GLib.SList, GSList*
	Prev *Source
	Next *Source
	Name string
	Priv *SourcePrivate
}

func sourceNewFromC(c *C.GSource) *Source {
	if c == nil {
		return nil
	}

	g := &Source{
		CallbackData:  (uintptr)(c.callback_data),
		CallbackFuncs: sourceCallbackFuncsNewFromC(c.callback_funcs),
		Context:       mainContextNewFromC(c.context),
		Flags:         (uint32)(c.flags),
		Name:          C.GoString(c.name),
		Next:          sourceNewFromC(c.next),
		Prev:          sourceNewFromC(c.prev),
		Priority:      (int32)(c.priority),
		Priv:          sourcePrivateNewFromC(c.priv),
		RefCount:      (uint32)(c.ref_count),
		SourceFuncs:   sourceFuncsNewFromC(c.source_funcs),
		SourceId:      (uint32)(c.source_id),
		native:        c,
	}

	return g
}

// Unsupported : g_source_new : unsupported parameter source_funcs : record param - coming soon

// SourceCallbackFuncs is a wrapper around the C record GSourceCallbackFuncs.
type SourceCallbackFuncs struct {
	native *C.GSourceCallbackFuncs
	// no type for ref
	// no type for unref
	// no type for get
}

func sourceCallbackFuncsNewFromC(c *C.GSourceCallbackFuncs) *SourceCallbackFuncs {
	if c == nil {
		return nil
	}

	g := &SourceCallbackFuncs{native: c}

	return g
}

// SourceFuncs is a wrapper around the C record GSourceFuncs.
type SourceFuncs struct {
	native *C.GSourceFuncs
	// no type for prepare
	// no type for check
	// no type for dispatch
	// no type for finalize
	// closure_callback : no type generator for SourceFunc, GSourceFunc
	// closure_marshal : no type generator for SourceDummyMarshal, GSourceDummyMarshal
}

func sourceFuncsNewFromC(c *C.GSourceFuncs) *SourceFuncs {
	if c == nil {
		return nil
	}

	g := &SourceFuncs{native: c}

	return g
}

// SourcePrivate is a wrapper around the C record GSourcePrivate.
type SourcePrivate struct {
	native *C.GSourcePrivate
}

func sourcePrivateNewFromC(c *C.GSourcePrivate) *SourcePrivate {
	if c == nil {
		return nil
	}

	g := &SourcePrivate{native: c}

	return g
}

// StatBuf is a wrapper around the C record GStatBuf.
type StatBuf struct {
	native *C.GStatBuf
}

func statBufNewFromC(c *C.GStatBuf) *StatBuf {
	if c == nil {
		return nil
	}

	g := &StatBuf{native: c}

	return g
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

	g := &String{
		AllocatedLen: (uint64)(c.allocated_len),
		Len:          (uint64)(c.len),
		Str:          C.GoString(c.str),
		native:       c,
	}

	return g
}

// StringChunk is a wrapper around the C record GStringChunk.
type StringChunk struct {
	native *C.GStringChunk
}

func stringChunkNewFromC(c *C.GStringChunk) *StringChunk {
	if c == nil {
		return nil
	}

	g := &StringChunk{native: c}

	return g
}

// TestCase is a wrapper around the C record GTestCase.
type TestCase struct {
	native *C.GTestCase
}

func testCaseNewFromC(c *C.GTestCase) *TestCase {
	if c == nil {
		return nil
	}

	g := &TestCase{native: c}

	return g
}

// TestConfig is a wrapper around the C record GTestConfig.
type TestConfig struct {
	native *C.GTestConfig
	// test_initialized : no type generator for gboolean, gboolean
	// test_quick : no type generator for gboolean, gboolean
	// test_perf : no type generator for gboolean, gboolean
	// test_verbose : no type generator for gboolean, gboolean
	// test_quiet : no type generator for gboolean, gboolean
	// test_undefined : no type generator for gboolean, gboolean
}

func testConfigNewFromC(c *C.GTestConfig) *TestConfig {
	if c == nil {
		return nil
	}

	g := &TestConfig{native: c}

	return g
}

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// TestSuite is a wrapper around the C record GTestSuite.
type TestSuite struct {
	native *C.GTestSuite
}

func testSuiteNewFromC(c *C.GTestSuite) *TestSuite {
	if c == nil {
		return nil
	}

	g := &TestSuite{native: c}

	return g
}

// Thread is a wrapper around the C record GThread.
type Thread struct {
	native *C.GThread
}

func threadNewFromC(c *C.GThread) *Thread {
	if c == nil {
		return nil
	}

	g := &Thread{native: c}

	return g
}

// Unsupported : g_thread_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// Unsupported : g_thread_try_new : unsupported parameter func : no type generator for ThreadFunc, GThreadFunc

// ThreadPool is a wrapper around the C record GThreadPool.
type ThreadPool struct {
	native *C.GThreadPool
	// _func : no type generator for Func, GFunc
	UserData uintptr
	// exclusive : no type generator for gboolean, gboolean
}

func threadPoolNewFromC(c *C.GThreadPool) *ThreadPool {
	if c == nil {
		return nil
	}

	g := &ThreadPool{
		UserData: (uintptr)(c.user_data),
		native:   c,
	}

	return g
}

// TimeVal is a wrapper around the C record GTimeVal.
type TimeVal struct {
	native *C.GTimeVal
	TvSec  int64
	TvUsec int64
}

func timeValNewFromC(c *C.GTimeVal) *TimeVal {
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

// Timer is a wrapper around the C record GTimer.
type Timer struct {
	native *C.GTimer
}

func timerNewFromC(c *C.GTimer) *Timer {
	if c == nil {
		return nil
	}

	g := &Timer{native: c}

	return g
}

// TrashStack is a wrapper around the C record GTrashStack.
type TrashStack struct {
	native *C.GTrashStack
	Next   *TrashStack
}

func trashStackNewFromC(c *C.GTrashStack) *TrashStack {
	if c == nil {
		return nil
	}

	g := &TrashStack{
		Next:   trashStackNewFromC(c.next),
		native: c,
	}

	return g
}

// Tree is a wrapper around the C record GTree.
type Tree struct {
	native *C.GTree
}

func treeNewFromC(c *C.GTree) *Tree {
	if c == nil {
		return nil
	}

	g := &Tree{native: c}

	return g
}

// VariantBuilder is a wrapper around the C record GVariantBuilder.
type VariantBuilder struct {
	native *C.GVariantBuilder
}

func variantBuilderNewFromC(c *C.GVariantBuilder) *VariantBuilder {
	if c == nil {
		return nil
	}

	g := &VariantBuilder{native: c}

	return g
}

// Unsupported : g_variant_builder_new : unsupported parameter type : Blacklisted record : GVariantType

// VariantIter is a wrapper around the C record GVariantIter.
type VariantIter struct {
	native *C.GVariantIter
	// no type for x
}

func variantIterNewFromC(c *C.GVariantIter) *VariantIter {
	if c == nil {
		return nil
	}

	g := &VariantIter{native: c}

	return g
}

// Blacklisted : GVariantType
