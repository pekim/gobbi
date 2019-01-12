// This is a generated file - DO NOT EDIT

package fontconfig

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <stdlib.h>
import "C"

// Pattern is a wrapper around the C record FcPattern.
type Pattern struct {
	native *C.FcPattern
}

func PatternNewFromC(u unsafe.Pointer) *Pattern {
	c := (*C.FcPattern)(u)
	if c == nil {
		return nil
	}

	g := &Pattern{native: c}

	return g
}

func (recv *Pattern) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Pattern with another Pattern, and returns true if they represent the same GObject.
func (recv *Pattern) Equals(other *Pattern) bool {
	return other.ToC() == recv.ToC()
}

// CharSet is a wrapper around the C record FcCharSet.
type CharSet struct {
	native *C.FcCharSet
}

func CharSetNewFromC(u unsafe.Pointer) *CharSet {
	c := (*C.FcCharSet)(u)
	if c == nil {
		return nil
	}

	g := &CharSet{native: c}

	return g
}

func (recv *CharSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CharSet with another CharSet, and returns true if they represent the same GObject.
func (recv *CharSet) Equals(other *CharSet) bool {
	return other.ToC() == recv.ToC()
}
