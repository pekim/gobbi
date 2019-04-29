// This is a generated file - DO NOT EDIT
// +build glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// Bytes is a wrapper around the C record GBytes.
type Bytes struct {
	native unsafe.Pointer
}

func BytesNewFromC(u unsafe.Pointer) *Bytes {
	if u == nil {
		return nil
	}

	g := &Bytes{native: u}

	return g
}

func (recv *Bytes) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RWLock is a wrapper around the C record GRWLock.
type RWLock struct {
	native unsafe.Pointer
	// Private : p
	// Private : i
}

func RWLockNewFromC(u unsafe.Pointer) *RWLock {
	if u == nil {
		return nil
	}

	g := &RWLock{native: u}

	return g
}

func (recv *RWLock) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecMutex is a wrapper around the C record GRecMutex.
type RecMutex struct {
	native unsafe.Pointer
	// Private : p
	// Private : i
}

func RecMutexNewFromC(u unsafe.Pointer) *RecMutex {
	if u == nil {
		return nil
	}

	g := &RecMutex{native: u}

	return g
}

func (recv *RecMutex) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
