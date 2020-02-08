package gi

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"unsafe"
)

func TestMain(m *testing.M) {
	SetErrorHandler(func(err error) {
		panic(err)
	})

	Require("GLib", "2.0")

	os.Exit(m.Run())
}

func TestFunctionIntegerReturn(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "get_num_processors")

	ret := fn.Invoke([]Arg{}, 0, 0, Arg{typ: ArgType_uint})
	assert.True(t, ret.value.(uint) > 0)
}

func TestFunctionBooleanReturn(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "mem_is_system_malloc")

	ret := fn.Invoke([]Arg{}, 0, 0, Arg{typ: ArgType_boolean})
	assert.True(t, ret.value.(bool))
}

func TestFunctionIntArgPointerReturn(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "malloc")

	in1 := Arg{typ: ArgType_size, value: uint64(8), in: true}
	ret := fn.Invoke([]Arg{in1}, 1, 0, Arg{typ: ArgType_pointer})
	assert.NotNil(t, (*byte)(ret.value.(unsafe.Pointer)))
}

func BenchmarkFuncCall(b *testing.B) {
	fn, _ := Function2InvokerNew("GLib", "get_num_processors")

	for n := 0; n < b.N; n++ {
		fn.Invoke([]Arg{}, 0, 0, Arg{typ: ArgType_uint})
	}
}
