package gi

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
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

	args := []*Arg{}
	retArg := &Arg{typ: ArgType_uint}
	fn.Invoke(args, 0, 0, retArg)
	assert.True(t, retArg.value.(uint) > 0)
}

func TestFunctionBooleanReturn(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "mem_is_system_malloc")

	args := []*Arg{}
	retArg := &Arg{typ: ArgType_boolean}
	fn.Invoke(args, 0, 0, retArg)
	assert.True(t, retArg.value.(bool))
}

func TestFunctionIntegerArgPointerReturn(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "malloc")

	in1 := &Arg{typ: ArgType_size, value: uint(8), in: true}
	args := []*Arg{in1}
	retArg := &Arg{typ: ArgType_pointer}
	fn.Invoke(args, 1, 0, retArg)
	assert.NotNil(t, (*byte)(retArg.value.(unsafe.Pointer)))
}

func TestFunctionStringArg(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "ascii_strup")

	in1 := &Arg{typ: ArgType_string, value: "QWerty", in: true}
	in2 := &Arg{typ: ArgType_ssize, value: -1, in: true}
	args := []*Arg{in1, in2}
	ret := &Arg{typ: ArgType_string}
	fn.Invoke(args, 2, 0, ret)

	assert.Equal(t, "QWERTY", ret.value.(string))
}

func TestFunctionOutStringArg(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "ascii_strtoll")

	in1 := &Arg{typ: ArgType_string, value: "123abc", in: true}
	out2 := &Arg{typ: ArgType_string, out: true}
	in3 := &Arg{typ: ArgType_uint, value: uint(10), in: true}
	args := []*Arg{in1, out2, in3}
	ret := &Arg{typ: ArgType_int64}
	fn.Invoke(args, 2, 1, ret)

	assert.Equal(t, int64(123), ret.value.(int64))
	assert.Equal(t, "abc", out2.value.(string))
}

func TestFunctionReturnStringArrayArg(t *testing.T) {
	fn, _ := Function2InvokerNew("GLib", "get_environ")

	args := []*Arg{}
	ret := &Arg{
		typ:                 ArgType_string,
		array:               true,
		arrayNullTerminated: true,
		transferOwnership:   TransferOwnershipFull,
	}
	fn.Invoke(args, 0, 0, ret)

	envVars := ret.value.([]string)
	assert.True(t, len(envVars) > 0)
	env := envVars[0]
	assert.True(t, strings.ContainsRune(env, '='))
}

func TestFunctionPointerArg(t *testing.T) {
	// malloc some memory, ready to free it
	malloc, _ := Function2InvokerNew("GLib", "malloc")
	in1 := &Arg{typ: ArgType_size, value: uint(8), in: true}
	args := []*Arg{in1}
	retArg := &Arg{typ: ArgType_pointer}
	malloc.Invoke(args, 1, 0, retArg)

	// free is the function with a pointer arg
	free, _ := Function2InvokerNew("GLib", "free")
	in1 = &Arg{typ: ArgType_pointer, value: retArg.value.(unsafe.Pointer), in: true}
	args = []*Arg{in1}
	retArg = &Arg{}
	free.Invoke(args, 1, 0, retArg)
}

func BenchmarkFuncCall(b *testing.B) {
	fn, _ := Function2InvokerNew("GLib", "get_num_processors")

	for n := 0; n < b.N; n++ {
		fn.Invoke([]*Arg{}, 0, 0, &Arg{typ: ArgType_uint})
	}
}
