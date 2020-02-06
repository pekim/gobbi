package ffi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLibrary(t *testing.T) {
	lib := NewLibrary("libgtk-3.so")
	assert.NotNil(t, lib.handle)
}

func TestNewLibraryFail(t *testing.T) {
	assert.Panics(t, func() {
		NewLibrary("bad")
	})
}
