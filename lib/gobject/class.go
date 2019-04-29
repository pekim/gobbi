// This is a generated file - DO NOT EDIT

package gobject

import "unsafe"

// InitiallyUnowned is a wrapper around the C record GInitiallyUnowned.
type InitiallyUnowned struct {
	native unsafe.Pointer
	// All fields are private
}

// Object is a wrapper around the C record GObject.
type Object struct {
	native unsafe.Pointer
	// All fields are private
}

// ParamSpec is a wrapper around the C record GParamSpec.
type ParamSpec struct {
	native unsafe.Pointer
	// g_type_instance : record
	Name      string
	Flags     ParamFlags
	ValueType Type
	OwnerType Type
	// Private : _nick
	// Private : _blurb
	// Private : qdata
	// Private : ref_count
	// Private : param_id
}

// ParamSpecBoolean is a wrapper around the C record GParamSpecBoolean.
type ParamSpecBoolean struct {
	native unsafe.Pointer
	// parent_instance : record
	DefaultValue bool
}

// ParamSpecBoxed is a wrapper around the C record GParamSpecBoxed.
type ParamSpecBoxed struct {
	native unsafe.Pointer
	// parent_instance : record
}

// ParamSpecChar is a wrapper around the C record GParamSpecChar.
type ParamSpecChar struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      int8
	Maximum      int8
	DefaultValue int8
}

// ParamSpecDouble is a wrapper around the C record GParamSpecDouble.
type ParamSpecDouble struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      float64
	Maximum      float64
	DefaultValue float64
	Epsilon      float64
}

// ParamSpecEnum is a wrapper around the C record GParamSpecEnum.
type ParamSpecEnum struct {
	native unsafe.Pointer
	// parent_instance : record
	// enum_class : record
	DefaultValue int32
}

// ParamSpecFlags is a wrapper around the C record GParamSpecFlags.
type ParamSpecFlags struct {
	native unsafe.Pointer
	// parent_instance : record
	// flags_class : record
	DefaultValue uint32
}

// ParamSpecFloat is a wrapper around the C record GParamSpecFloat.
type ParamSpecFloat struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      float32
	Maximum      float32
	DefaultValue float32
	Epsilon      float32
}

// ParamSpecInt is a wrapper around the C record GParamSpecInt.
type ParamSpecInt struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      int32
	Maximum      int32
	DefaultValue int32
}

// ParamSpecInt64 is a wrapper around the C record GParamSpecInt64.
type ParamSpecInt64 struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      int64
	Maximum      int64
	DefaultValue int64
}

// ParamSpecLong is a wrapper around the C record GParamSpecLong.
type ParamSpecLong struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      int64
	Maximum      int64
	DefaultValue int64
}

// ParamSpecObject is a wrapper around the C record GParamSpecObject.
type ParamSpecObject struct {
	native unsafe.Pointer
	// parent_instance : record
}

// ParamSpecParam is a wrapper around the C record GParamSpecParam.
type ParamSpecParam struct {
	native unsafe.Pointer
	// parent_instance : record
}

// ParamSpecPointer is a wrapper around the C record GParamSpecPointer.
type ParamSpecPointer struct {
	native unsafe.Pointer
	// parent_instance : record
}

// ParamSpecString is a wrapper around the C record GParamSpecString.
type ParamSpecString struct {
	native unsafe.Pointer
	// parent_instance : record
	DefaultValue string
	CsetFirst    string
	CsetNth      string
	Substitutor  rune
	// Bitfield not supported :  1 null_fold_if_empty
	// Bitfield not supported :  1 ensure_non_null
}

// ParamSpecUChar is a wrapper around the C record GParamSpecUChar.
type ParamSpecUChar struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      uint8
	Maximum      uint8
	DefaultValue uint8
}

// ParamSpecUInt is a wrapper around the C record GParamSpecUInt.
type ParamSpecUInt struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      uint32
	Maximum      uint32
	DefaultValue uint32
}

// ParamSpecUInt64 is a wrapper around the C record GParamSpecUInt64.
type ParamSpecUInt64 struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      uint64
	Maximum      uint64
	DefaultValue uint64
}

// ParamSpecULong is a wrapper around the C record GParamSpecULong.
type ParamSpecULong struct {
	native unsafe.Pointer
	// parent_instance : record
	Minimum      uint64
	Maximum      uint64
	DefaultValue uint64
}

// ParamSpecUnichar is a wrapper around the C record GParamSpecUnichar.
type ParamSpecUnichar struct {
	native unsafe.Pointer
	// parent_instance : record
	DefaultValue rune
}

// ParamSpecValueArray is a wrapper around the C record GParamSpecValueArray.
type ParamSpecValueArray struct {
	native unsafe.Pointer
	// parent_instance : record
	// element_spec : record
	FixedNElements uint32
}

// Blacklisted : GTypeModule
