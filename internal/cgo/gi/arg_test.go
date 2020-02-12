package gi

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestArgValueBoolean(t *testing.T) {
	arg1 := Arg{typ: ArgType_boolean, value: true}
	giArg := arg1.getValue()

	arg2 := Arg{typ: ArgType_boolean, value: false}
	arg2.setValue(giArg)

	assert.True(t, arg2.value.(bool))
}

func TestArgValueString(t *testing.T) {
	value := "qwerty"

	arg1 := Arg{typ: ArgType_string, value: value, in: true}
	giArg := arg1.getValue()

	arg2 := Arg{typ: ArgType_string, value: ""}
	arg2.setValue(giArg)

	assert.Equal(t, value, arg2.value.(string))
}

func TestArgValueStringArrayNullTerminated(t *testing.T) {
	value := []string{"one", "two"}

	arg1 := Arg{typ: ArgType_string, value: value, array: true, arrayNullTerminated: true, in: true}
	giArg := arg1.getValue()

	arg2 := Arg{typ: ArgType_string, value: []string{}, array: true, arrayNullTerminated: true}
	arg2.setValue(giArg)

	assert.Equal(t, value, arg2.value.([]string))
}

func TestArgValueSimple(t *testing.T) {
	for _, test := range []struct {
		name      string
		typ       ArgType
		value     interface{}
		initValue interface{}
	}{
		{"int8", ArgType_int8, int8(42), int8(3)},
		{"uint8", ArgType_uint8, uint8(42), uint8(3)},
		{"int16", ArgType_int16, int16(42), int16(3)},
		{"uint16", ArgType_uint16, uint16(42), uint16(3)},
		{"int32", ArgType_int32, int32(42), int32(3)},
		{"uint32", ArgType_uint32, uint32(42), uint32(3)},
		{"int64", ArgType_int64, int64(42), int64(3)},
		{"uint64", ArgType_uint64, uint64(42), uint64(3)},
		{"float", ArgType_float, float32(42), float32(3)},
		{"double", ArgType_double, float64(42), float64(3)},
		{"short", ArgType_int16, int16(42), int16(3)},
		{"ushort", ArgType_uint16, uint16(42), uint16(3)},
		{"int", ArgType_int, int(42), int(3)},
		{"uint", ArgType_uint, uint(42), uint(3)},
		{"long", ArgType_long, int64(42), int64(3)},
		{"ulong", ArgType_ulong, uint64(42), uint64(3)},
		{"ssize", ArgType_ssize, int(42), int(3)},
		{"size", ArgType_size, uint(42), uint(3)},
	} {
		t.Run(test.name, func(t *testing.T) {
			arg1 := Arg{typ: test.typ, value: test.value}
			giArg := arg1.getValue()

			arg2 := Arg{typ: test.typ, value: test.initValue}
			arg2.setValue(giArg)

			assert.Equal(t, test.value, arg2.value)
		})
	}
}

func TestArgValueSimpleArray(t *testing.T) {
	for _, test := range []struct {
		name      string
		typ       ArgType
		value     interface{}
		initValue interface{}
	}{
		{"int8", ArgType_int8, []int8{42, 43}, []int8{1, 2}},
		{"uint8", ArgType_uint8, []uint8{42, 43}, []uint8{1, 2}},
		{"int16", ArgType_int16, []int16{42, 43}, []int16{1, 2}},
		{"uint16", ArgType_uint16, []uint16{42, 43}, []uint16{1, 2}},
		{"int32", ArgType_int32, []int32{42, 43}, []int32{1, 2}},
		{"uint32", ArgType_uint32, []uint32{42, 43}, []uint32{1, 2}},
		{"int64", ArgType_int64, []int64{42, 43}, []int64{1, 2}},
		{"uint64", ArgType_uint64, []uint64{42, 43}, []uint64{1, 2}},
		{"float", ArgType_float, []float32{42, 43}, []float32{1, 2}},
		{"double", ArgType_double, []float64{42, 43}, []float64{1, 2}},
		{"short", ArgType_int16, []int16{42, 43}, []int16{1, 2}},
		{"ushort", ArgType_uint16, []uint16{42, 43}, []uint16{1, 2}},
		{"int", ArgType_int, []int{42, 43}, []int{1, 2}},
		{"uint", ArgType_uint, []uint{42, 43}, []uint{1, 2}},
		{"long", ArgType_long, []int64{42, 43}, []int64{1, 2}},
		{"ulong", ArgType_ulong, []uint64{42, 43}, []uint64{1, 2}},
		{"ssize", ArgType_ssize, []int{42, 43}, []int{1, 2}},
		{"size", ArgType_size, []uint{42, 43}, []uint{1, 2}},
		{"pointer", ArgType_pointer, []unsafe.Pointer{unsafe.Pointer(uintptr(42)), unsafe.Pointer(uintptr(43))},
			[]unsafe.Pointer{unsafe.Pointer(uintptr(1)), unsafe.Pointer(uintptr(2))}},
	} {
		t.Run(test.name, func(t *testing.T) {
			arg1 := Arg{typ: test.typ, array: true, value: test.value}
			giArg := arg1.getValue()

			arg2 := Arg{typ: test.typ, array: true, value: test.initValue, arrayLength: 2}
			arg2.setValue(giArg)

			assert.Equal(t, test.value, arg2.value)
		})
	}
}
