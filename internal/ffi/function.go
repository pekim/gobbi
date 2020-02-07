package ffi

import "unsafe"

type Function struct {
	library *Library
	name    string
	fn      unsafe.Pointer
}
