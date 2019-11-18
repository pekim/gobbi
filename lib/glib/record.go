// Code generated - DO NOT EDIT.

package glib

import gi "github.com/pekim/gobbi/internal/gi"

type Array struct {
	native uintptr
	Data   string
	Len    uint32
}

type AsyncQueue struct {
	native uintptr
}

type BookmarkFile struct {
	native uintptr
}

type ByteArray struct {
	native uintptr
	Data   uint8
	Len    uint32
}

type Bytes struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_bytes_new' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_bytes_new_static' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_bytes_new_take' : parameter 'data' has no type

// UNSUPPORTED : C value 'g_bytes_new_with_free_func' : parameter 'data' has no type

type Checksum struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_checksum_new' : parameter 'checksum_type' of type 'ChecksumType' not supported

type Cond struct {
	native uintptr
}

type Data struct {
	native uintptr
}

type Date struct {
	native     uintptr
	JulianDays uint32
	Julian     uint32
	Dmy        uint32
	Day        uint32
	Month      uint32
	Year       uint32
}

var newDateInvoker *gi.Function

// DateNew is a representation of the C type g_date_new.
func DateNew() *Date {
	if newDateInvoker == nil {
		newDateInvoker = gi.FunctionInvokerNew("GLib", "new")
	}

	ret := newDateInvoker.Invoke(nil, nil)

	return &Date{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_date_new_dmy' : parameter 'day' of type 'DateDay' not supported

var newJulianDateInvoker *gi.Function

// DateNewJulian is a representation of the C type g_date_new_julian.
func DateNewJulian(julianDay uint32) *Date {
	if newJulianDateInvoker == nil {
		newJulianDateInvoker = gi.FunctionInvokerNew("GLib", "new_julian")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(julianDay)

	ret := newJulianDateInvoker.Invoke(inArgs[:], nil)

	return &Date{native: ret.Pointer()}
}

type DateTime struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_date_time_new' : parameter 'tz' of type 'TimeZone' not supported

// UNSUPPORTED : C value 'g_date_time_new_from_iso8601' : parameter 'default_tz' of type 'TimeZone' not supported

// UNSUPPORTED : C value 'g_date_time_new_from_timeval_local' : parameter 'tv' of type 'TimeVal' not supported

// UNSUPPORTED : C value 'g_date_time_new_from_timeval_utc' : parameter 'tv' of type 'TimeVal' not supported

var newFromUnixLocalDateTimeInvoker *gi.Function

// DateTimeNewFromUnixLocal is a representation of the C type g_date_time_new_from_unix_local.
func DateTimeNewFromUnixLocal(t int64) *DateTime {
	if newFromUnixLocalDateTimeInvoker == nil {
		newFromUnixLocalDateTimeInvoker = gi.FunctionInvokerNew("GLib", "new_from_unix_local")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(t)

	ret := newFromUnixLocalDateTimeInvoker.Invoke(inArgs[:], nil)

	return &DateTime{native: ret.Pointer()}
}

var newFromUnixUtcDateTimeInvoker *gi.Function

// DateTimeNewFromUnixUtc is a representation of the C type g_date_time_new_from_unix_utc.
func DateTimeNewFromUnixUtc(t int64) *DateTime {
	if newFromUnixUtcDateTimeInvoker == nil {
		newFromUnixUtcDateTimeInvoker = gi.FunctionInvokerNew("GLib", "new_from_unix_utc")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(t)

	ret := newFromUnixUtcDateTimeInvoker.Invoke(inArgs[:], nil)

	return &DateTime{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_date_time_new_local' : parameter 'seconds' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_date_time_new_now' : parameter 'tz' of type 'TimeZone' not supported

var newNowLocalDateTimeInvoker *gi.Function

// DateTimeNewNowLocal is a representation of the C type g_date_time_new_now_local.
func DateTimeNewNowLocal() *DateTime {
	if newNowLocalDateTimeInvoker == nil {
		newNowLocalDateTimeInvoker = gi.FunctionInvokerNew("GLib", "new_now_local")
	}

	ret := newNowLocalDateTimeInvoker.Invoke(nil, nil)

	return &DateTime{native: ret.Pointer()}
}

var newNowUtcDateTimeInvoker *gi.Function

// DateTimeNewNowUtc is a representation of the C type g_date_time_new_now_utc.
func DateTimeNewNowUtc() *DateTime {
	if newNowUtcDateTimeInvoker == nil {
		newNowUtcDateTimeInvoker = gi.FunctionInvokerNew("GLib", "new_now_utc")
	}

	ret := newNowUtcDateTimeInvoker.Invoke(nil, nil)

	return &DateTime{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_date_time_new_utc' : parameter 'seconds' of type 'gdouble' not supported

type DebugKey struct {
	native uintptr
	Key    string
	Value  uint32
}

type Dir struct {
	native uintptr
}

type Error struct {
	native  uintptr
	Domain  Quark
	Code    int32
	Message string
}

// UNSUPPORTED : C value 'g_error_new' : parameter 'domain' of type 'Quark' not supported

// UNSUPPORTED : C value 'g_error_new_literal' : parameter 'domain' of type 'Quark' not supported

// UNSUPPORTED : C value 'g_error_new_valist' : parameter 'domain' of type 'Quark' not supported

type HashTable struct {
	native uintptr
}

type HashTableIter struct {
	native uintptr
}

type Hmac struct {
	native uintptr
}

type Hook struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'

	Next     *Hook
	Prev     *Hook
	RefCount uint32
	HookId   uint64
	Flags    uint32
	// UNSUPPORTED : C value 'func' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'destroy' : no Go type for 'DestroyNotify'

}

type HookList struct {
	native   uintptr
	SeqId    uint64
	HookSize uint32
	IsSetup  uint32
	Hooks    *Hook
	// UNSUPPORTED : C value 'dummy3' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'finalize_hook' : no Go type for 'HookFinalizeFunc'

	// UNSUPPORTED : C value 'dummy' : missing Type

}

type IConv struct {
	native uintptr
}

type IOChannel struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_io_channel_new_file' : parameter 'filename' of type 'filename' not supported

var unixNewIOChannelInvoker *gi.Function

// IOChannelUnixNew is a representation of the C type g_io_channel_unix_new.
func IOChannelUnixNew(fd int32) *IOChannel {
	if unixNewIOChannelInvoker == nil {
		unixNewIOChannelInvoker = gi.FunctionInvokerNew("GLib", "unix_new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(fd)

	ret := unixNewIOChannelInvoker.Invoke(inArgs[:], nil)

	return &IOChannel{native: ret.Pointer()}
}

type IOFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'io_read' : missing Type

	// UNSUPPORTED : C value 'io_write' : missing Type

	// UNSUPPORTED : C value 'io_seek' : missing Type

	// UNSUPPORTED : C value 'io_close' : missing Type

	// UNSUPPORTED : C value 'io_create_watch' : missing Type

	// UNSUPPORTED : C value 'io_free' : missing Type

	// UNSUPPORTED : C value 'io_set_flags' : missing Type

	// UNSUPPORTED : C value 'io_get_flags' : missing Type

}

type KeyFile struct {
	native uintptr
}

var newKeyFileInvoker *gi.Function

// KeyFileNew is a representation of the C type g_key_file_new.
func KeyFileNew() *KeyFile {
	if newKeyFileInvoker == nil {
		newKeyFileInvoker = gi.FunctionInvokerNew("GLib", "new")
	}

	ret := newKeyFileInvoker.Invoke(nil, nil)

	return &KeyFile{native: ret.Pointer()}
}

type List struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'next' : no Go type for 'GLib.List'

	// UNSUPPORTED : C value 'prev' : no Go type for 'GLib.List'

}

type LogField struct {
	native uintptr
	Key    string
	// UNSUPPORTED : C value 'value' : no Go type for 'gpointer'

	Length int32
}

type MainContext struct {
	native uintptr
}

var newMainContextInvoker *gi.Function

// MainContextNew is a representation of the C type g_main_context_new.
func MainContextNew() *MainContext {
	if newMainContextInvoker == nil {
		newMainContextInvoker = gi.FunctionInvokerNew("GLib", "new")
	}

	ret := newMainContextInvoker.Invoke(nil, nil)

	return &MainContext{native: ret.Pointer()}
}

type MainLoop struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_main_loop_new' : parameter 'context' of type 'MainContext' not supported

type MappedFile struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_mapped_file_new' : parameter 'filename' of type 'filename' not supported

// UNSUPPORTED : C value 'g_mapped_file_new_from_fd' : parameter 'writable' of type 'gboolean' not supported

type MarkupParseContext struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_markup_parse_context_new' : parameter 'parser' of type 'MarkupParser' not supported

type MarkupParser struct {
	native uintptr
	// UNSUPPORTED : C value 'start_element' : missing Type

	// UNSUPPORTED : C value 'end_element' : missing Type

	// UNSUPPORTED : C value 'text' : missing Type

	// UNSUPPORTED : C value 'passthrough' : missing Type

	// UNSUPPORTED : C value 'error' : missing Type

}

type MatchInfo struct {
	native uintptr
}

type MemVTable struct {
	native uintptr
	// UNSUPPORTED : C value 'malloc' : missing Type

	// UNSUPPORTED : C value 'realloc' : missing Type

	// UNSUPPORTED : C value 'free' : missing Type

	// UNSUPPORTED : C value 'calloc' : missing Type

	// UNSUPPORTED : C value 'try_malloc' : missing Type

	// UNSUPPORTED : C value 'try_realloc' : missing Type

}

type Node struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'

	Next     *Node
	Prev     *Node
	Parent   *Node
	Children *Node
}

type Once struct {
	native uintptr
	// UNSUPPORTED : C value 'status' : no Go type for 'OnceStatus'

	// UNSUPPORTED : C value 'retval' : no Go type for 'gpointer'

}

type OptionContext struct {
	native uintptr
}

type OptionEntry struct {
	native    uintptr
	LongName  string
	ShortName int8
	Flags     int32
	// UNSUPPORTED : C value 'arg' : no Go type for 'OptionArg'

	// UNSUPPORTED : C value 'arg_data' : no Go type for 'gpointer'

	Description    string
	ArgDescription string
}

type OptionGroup struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_option_group_new' : parameter 'user_data' of type 'gpointer' not supported

type PatternSpec struct {
	native uintptr
}

type PollFD struct {
	native  uintptr
	Fd      int32
	Events  uint16
	Revents uint16
}

type Private struct {
	native uintptr
}

type PtrArray struct {
	native uintptr
	// UNSUPPORTED : C value 'pdata' : no Go type for 'gpointer'

	Len uint32
}

type Queue struct {
	native uintptr
	// UNSUPPORTED : C value 'head' : no Go type for 'GLib.List'

	// UNSUPPORTED : C value 'tail' : no Go type for 'GLib.List'

	Length uint32
}

type RWLock struct {
	native uintptr
}

type Rand struct {
	native uintptr
}

type RecMutex struct {
	native uintptr
}

type Regex struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_regex_new' : parameter 'compile_options' of type 'RegexCompileFlags' not supported

type SList struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'next' : no Go type for 'GLib.SList'

}

type Scanner struct {
	native uintptr
	// UNSUPPORTED : C value 'user_data' : no Go type for 'gpointer'

	MaxParseErrors uint32
	ParseErrors    uint32
	InputName      string
	Qdata          *Data
	Config         *ScannerConfig
	// UNSUPPORTED : C value 'token' : no Go type for 'TokenType'

	// UNSUPPORTED : C value 'value' : no Go type for 'TokenValue'

	Line     uint32
	Position uint32
	// UNSUPPORTED : C value 'next_token' : no Go type for 'TokenType'

	// UNSUPPORTED : C value 'next_value' : no Go type for 'TokenValue'

	NextLine     uint32
	NextPosition uint32
	// UNSUPPORTED : C value 'msg_handler' : no Go type for 'ScannerMsgFunc'

}

type ScannerConfig struct {
	native              uintptr
	CsetSkipCharacters  string
	CsetIdentifierFirst string
	CsetIdentifierNth   string
	CpairCommentSingle  string
	CaseSensitive       uint32
	SkipCommentMulti    uint32
	SkipCommentSingle   uint32
	ScanCommentMulti    uint32
	ScanIdentifier      uint32
	ScanIdentifier1char uint32
	ScanIdentifierNULL  uint32
	ScanSymbols         uint32
	ScanBinary          uint32
	ScanOctal           uint32
	ScanFloat           uint32
	ScanHex             uint32
	ScanHexDollar       uint32
	ScanStringSq        uint32
	ScanStringDq        uint32
	Numbers2Int         uint32
	Int2Float           uint32
	Identifier2String   uint32
	Char2Token          uint32
	Symbol2Token        uint32
	Scope0Fallback      uint32
	StoreInt64          uint32
}

type Sequence struct {
	native uintptr
}

type SequenceIter struct {
	native uintptr
}

type Source struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_source_new' : parameter 'source_funcs' of type 'SourceFuncs' not supported

type SourceCallbackFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'ref' : missing Type

	// UNSUPPORTED : C value 'unref' : missing Type

	// UNSUPPORTED : C value 'get' : missing Type

}

type SourceFuncs struct {
	native uintptr
	// UNSUPPORTED : C value 'prepare' : missing Type

	// UNSUPPORTED : C value 'check' : missing Type

	// UNSUPPORTED : C value 'dispatch' : missing Type

	// UNSUPPORTED : C value 'finalize' : missing Type

}

type SourcePrivate struct {
	native uintptr
}

type StatBuf struct {
	native uintptr
}

type String struct {
	native       uintptr
	Str          string
	Len          uintptr
	AllocatedLen uintptr
}

type StringChunk struct {
	native uintptr
}

type TestCase struct {
	native uintptr
}

type TestConfig struct {
	native uintptr
	// UNSUPPORTED : C value 'test_initialized' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'test_quick' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'test_perf' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'test_verbose' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'test_quiet' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'test_undefined' : no Go type for 'gboolean'

}

type TestLogBuffer struct {
	native uintptr
}

type TestLogMsg struct {
	native uintptr
	// UNSUPPORTED : C value 'log_type' : no Go type for 'TestLogType'

	NStrings uint32
	Strings  string
	NNums    uint32
	// UNSUPPORTED : C value 'nums' : no Go type for 'long double'

}

type TestSuite struct {
	native uintptr
}

type Thread struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_thread_new' : parameter 'func' of type 'ThreadFunc' not supported

// UNSUPPORTED : C value 'g_thread_try_new' : parameter 'func' of type 'ThreadFunc' not supported

type ThreadPool struct {
	native uintptr
	// UNSUPPORTED : C value 'func' : no Go type for 'Func'

	// UNSUPPORTED : C value 'user_data' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'exclusive' : no Go type for 'gboolean'

}

type TimeVal struct {
	native uintptr
	TvSec  int64
	TvUsec int64
}

type TimeZone struct {
	native uintptr
}

var newTimeZoneInvoker *gi.Function

// TimeZoneNew is a representation of the C type g_time_zone_new.
func TimeZoneNew(identifier string) *TimeZone {
	if newTimeZoneInvoker == nil {
		newTimeZoneInvoker = gi.FunctionInvokerNew("GLib", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(identifier)

	ret := newTimeZoneInvoker.Invoke(inArgs[:], nil)

	return &TimeZone{native: ret.Pointer()}
}

var newLocalTimeZoneInvoker *gi.Function

// TimeZoneNewLocal is a representation of the C type g_time_zone_new_local.
func TimeZoneNewLocal() *TimeZone {
	if newLocalTimeZoneInvoker == nil {
		newLocalTimeZoneInvoker = gi.FunctionInvokerNew("GLib", "new_local")
	}

	ret := newLocalTimeZoneInvoker.Invoke(nil, nil)

	return &TimeZone{native: ret.Pointer()}
}

var newOffsetTimeZoneInvoker *gi.Function

// TimeZoneNewOffset is a representation of the C type g_time_zone_new_offset.
func TimeZoneNewOffset(seconds int32) *TimeZone {
	if newOffsetTimeZoneInvoker == nil {
		newOffsetTimeZoneInvoker = gi.FunctionInvokerNew("GLib", "new_offset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(seconds)

	ret := newOffsetTimeZoneInvoker.Invoke(inArgs[:], nil)

	return &TimeZone{native: ret.Pointer()}
}

var newUtcTimeZoneInvoker *gi.Function

// TimeZoneNewUtc is a representation of the C type g_time_zone_new_utc.
func TimeZoneNewUtc() *TimeZone {
	if newUtcTimeZoneInvoker == nil {
		newUtcTimeZoneInvoker = gi.FunctionInvokerNew("GLib", "new_utc")
	}

	ret := newUtcTimeZoneInvoker.Invoke(nil, nil)

	return &TimeZone{native: ret.Pointer()}
}

type Timer struct {
	native uintptr
}

type TrashStack struct {
	native uintptr
	Next   *TrashStack
}

type Tree struct {
	native uintptr
}

type Variant struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_new' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_new_array' : parameter 'child_type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_new_boolean' : parameter 'value' of type 'gboolean' not supported

var newByteVariantInvoker *gi.Function

// VariantNewByte is a representation of the C type g_variant_new_byte.
func VariantNewByte(value uint8) *Variant {
	if newByteVariantInvoker == nil {
		newByteVariantInvoker = gi.FunctionInvokerNew("GLib", "new_byte")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint8(value)

	ret := newByteVariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_variant_new_bytestring' : parameter 'string' has no type

// UNSUPPORTED : C value 'g_variant_new_bytestring_array' : parameter 'strv' has no type

// UNSUPPORTED : C value 'g_variant_new_dict_entry' : parameter 'key' of type 'Variant' not supported

// UNSUPPORTED : C value 'g_variant_new_double' : parameter 'value' of type 'gdouble' not supported

// UNSUPPORTED : C value 'g_variant_new_fixed_array' : parameter 'element_type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_new_from_bytes' : parameter 'type' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_new_from_data' : parameter 'type' of type 'VariantType' not supported

var newHandleVariantInvoker *gi.Function

// VariantNewHandle is a representation of the C type g_variant_new_handle.
func VariantNewHandle(value int32) *Variant {
	if newHandleVariantInvoker == nil {
		newHandleVariantInvoker = gi.FunctionInvokerNew("GLib", "new_handle")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(value)

	ret := newHandleVariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

var newInt16VariantInvoker *gi.Function

// VariantNewInt16 is a representation of the C type g_variant_new_int16.
func VariantNewInt16(value int16) *Variant {
	if newInt16VariantInvoker == nil {
		newInt16VariantInvoker = gi.FunctionInvokerNew("GLib", "new_int16")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt16(value)

	ret := newInt16VariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

var newInt32VariantInvoker *gi.Function

// VariantNewInt32 is a representation of the C type g_variant_new_int32.
func VariantNewInt32(value int32) *Variant {
	if newInt32VariantInvoker == nil {
		newInt32VariantInvoker = gi.FunctionInvokerNew("GLib", "new_int32")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(value)

	ret := newInt32VariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

var newInt64VariantInvoker *gi.Function

// VariantNewInt64 is a representation of the C type g_variant_new_int64.
func VariantNewInt64(value int64) *Variant {
	if newInt64VariantInvoker == nil {
		newInt64VariantInvoker = gi.FunctionInvokerNew("GLib", "new_int64")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(value)

	ret := newInt64VariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_variant_new_maybe' : parameter 'child_type' of type 'VariantType' not supported

var newObjectPathVariantInvoker *gi.Function

// VariantNewObjectPath is a representation of the C type g_variant_new_object_path.
func VariantNewObjectPath(objectPath string) *Variant {
	if newObjectPathVariantInvoker == nil {
		newObjectPathVariantInvoker = gi.FunctionInvokerNew("GLib", "new_object_path")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(objectPath)

	ret := newObjectPathVariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_variant_new_objv' : parameter 'strv' has no type

// UNSUPPORTED : C value 'g_variant_new_parsed' : parameter '...' has no type

// UNSUPPORTED : C value 'g_variant_new_parsed_va' : parameter 'app' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_variant_new_printf' : parameter '...' has no type

var newSignatureVariantInvoker *gi.Function

// VariantNewSignature is a representation of the C type g_variant_new_signature.
func VariantNewSignature(signature string) *Variant {
	if newSignatureVariantInvoker == nil {
		newSignatureVariantInvoker = gi.FunctionInvokerNew("GLib", "new_signature")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(signature)

	ret := newSignatureVariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

var newStringVariantInvoker *gi.Function

// VariantNewString is a representation of the C type g_variant_new_string.
func VariantNewString(string_ string) *Variant {
	if newStringVariantInvoker == nil {
		newStringVariantInvoker = gi.FunctionInvokerNew("GLib", "new_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := newStringVariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_variant_new_strv' : parameter 'strv' has no type

var newTakeStringVariantInvoker *gi.Function

// VariantNewTakeString is a representation of the C type g_variant_new_take_string.
func VariantNewTakeString(string_ string) *Variant {
	if newTakeStringVariantInvoker == nil {
		newTakeStringVariantInvoker = gi.FunctionInvokerNew("GLib", "new_take_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(string_)

	ret := newTakeStringVariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_variant_new_tuple' : parameter 'children' has no type

var newUint16VariantInvoker *gi.Function

// VariantNewUint16 is a representation of the C type g_variant_new_uint16.
func VariantNewUint16(value uint16) *Variant {
	if newUint16VariantInvoker == nil {
		newUint16VariantInvoker = gi.FunctionInvokerNew("GLib", "new_uint16")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint16(value)

	ret := newUint16VariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

var newUint32VariantInvoker *gi.Function

// VariantNewUint32 is a representation of the C type g_variant_new_uint32.
func VariantNewUint32(value uint32) *Variant {
	if newUint32VariantInvoker == nil {
		newUint32VariantInvoker = gi.FunctionInvokerNew("GLib", "new_uint32")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(value)

	ret := newUint32VariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

var newUint64VariantInvoker *gi.Function

// VariantNewUint64 is a representation of the C type g_variant_new_uint64.
func VariantNewUint64(value uint64) *Variant {
	if newUint64VariantInvoker == nil {
		newUint64VariantInvoker = gi.FunctionInvokerNew("GLib", "new_uint64")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetUint64(value)

	ret := newUint64VariantInvoker.Invoke(inArgs[:], nil)

	return &Variant{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_variant_new_va' : parameter 'app' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_variant_new_variant' : parameter 'value' of type 'Variant' not supported

type VariantBuilder struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_builder_new' : parameter 'type' of type 'VariantType' not supported

type VariantDict struct {
	native uintptr
}

// UNSUPPORTED : C value 'g_variant_dict_new' : parameter 'from_asv' of type 'Variant' not supported

type VariantIter struct {
	native uintptr
}

type VariantType struct {
	native uintptr
}

var newVariantTypeInvoker *gi.Function

// VariantTypeNew is a representation of the C type g_variant_type_new.
func VariantTypeNew(typeString string) *VariantType {
	if newVariantTypeInvoker == nil {
		newVariantTypeInvoker = gi.FunctionInvokerNew("GLib", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(typeString)

	ret := newVariantTypeInvoker.Invoke(inArgs[:], nil)

	return &VariantType{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'g_variant_type_new_array' : parameter 'element' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_dict_entry' : parameter 'key' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_maybe' : parameter 'element' of type 'VariantType' not supported

// UNSUPPORTED : C value 'g_variant_type_new_tuple' : parameter 'items' has no type
