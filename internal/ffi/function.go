package ffi

import "unsafe"

type Function struct {
	fn unsafe.Pointer
}
