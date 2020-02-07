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
	fn := lib.function("gtk_init")

	assert.NotEqual(t, unsafe.Pointer(nil), fn.fn)
	assert.False(t, handledError)
}

func TestLibraryFunctionNotFound(t *testing.T) {
	handledError := false
	SetErrorHandler(func(err error) {
		handledError = true
	})

	lib := OpenLibrary("Gtk", "libgtk-3.so.0")
	fn := lib.function("bad")

	assert.Nil(t, fn)
	assert.True(t, handledError)
}
