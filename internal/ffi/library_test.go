package ffi

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestNewLibrary(t *testing.T) {
	lib := NewLibrary("libgtk-3.so.0")
	assert.NotNil(t, lib.handle)
}

func TestNewLibraryNotLoaded(t *testing.T) {
	assert.Panics(t, func() {
		NewLibrary("bad")
	})
}

func TestLibraryFunction(t *testing.T) {
	lib := NewLibrary("libgtk-3.so.0")
	fn, err := lib.function("gtk_init")

	assert.NotEqual(t, unsafe.Pointer(nil), fn)
	assert.Nil(t, err)
}

func TestLibraryFunctionNotFound(t *testing.T) {
	lib := NewLibrary("libgtk-3.so.0")
	_, err := lib.function("bad")

	assert.NotNil(t, err)
}
