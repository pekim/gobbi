package ffi

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestNewLibrary(t *testing.T) {
	lib := OpenLibrary("Gtk", "libgtk-3.so.0")
	assert.NotNil(t, lib.handle)
}

func TestNewLibraryNotLoaded(t *testing.T) {
	assert.Panics(t, func() {
		OpenLibrary("Bad", "bad")
	})
}

func TestLibraryFunction(t *testing.T) {
	handledError := false
	SetErrorHandler(func(err error) {
		handledError = true
	})

	lib := OpenLibrary("Gtk", "libgtk-3.so.0")
	fn := lib.Function("gtk_init", []Type{}, Type_void)

	assert.NotEqual(t, unsafe.Pointer(nil), fn.fn)
	assert.False(t, handledError)
}

func TestLibraryFunctionNotFound(t *testing.T) {
	handledError := false
	SetErrorHandler(func(err error) {
		handledError = true
	})

	lib := OpenLibrary("Gtk", "libgtk-3.so.0")
	fn := lib.Function("bad", []Type{}, Type_void)

	assert.Nil(t, fn)
	assert.True(t, handledError)
}

func TestLibraryFunctionCall(t *testing.T) {
	SetErrorHandler(func(err error) {
		panic(err)
	})

	lib := OpenLibrary("GLib", "libglib-2.0.so.0")
	fn := lib.Function("g_utf8_strdown", []Type{Type_pointer, Type_slong}, Type_void)
	fn.Call()
}
