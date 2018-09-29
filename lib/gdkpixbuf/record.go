// This is a generated file - DO NOT EDIT

package gdkpixbuf

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// PixbufFormat is a wrapper around the C record GdkPixbufFormat.
type PixbufFormat struct {
	native *C.GdkPixbufFormat
}

func PixbufFormatNewFromC(u unsafe.Pointer) *PixbufFormat {
	c := (*C.GdkPixbufFormat)(u)
	if c == nil {
		return nil
	}

	g := &PixbufFormat{native: c}

	return g
}

func (recv *PixbufFormat) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufLoaderClass is a wrapper around the C record GdkPixbufLoaderClass.
type PixbufLoaderClass struct {
	native *C.GdkPixbufLoaderClass
	// parent_class : record
	// no type for size_prepared
	// no type for area_prepared
	// no type for area_updated
	// no type for closed
}

func PixbufLoaderClassNewFromC(u unsafe.Pointer) *PixbufLoaderClass {
	c := (*C.GdkPixbufLoaderClass)(u)
	if c == nil {
		return nil
	}

	g := &PixbufLoaderClass{native: c}

	return g
}

func (recv *PixbufLoaderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufSimpleAnimClass is a wrapper around the C record GdkPixbufSimpleAnimClass.
type PixbufSimpleAnimClass struct {
	native *C.GdkPixbufSimpleAnimClass
}

func PixbufSimpleAnimClassNewFromC(u unsafe.Pointer) *PixbufSimpleAnimClass {
	c := (*C.GdkPixbufSimpleAnimClass)(u)
	if c == nil {
		return nil
	}

	g := &PixbufSimpleAnimClass{native: c}

	return g
}

func (recv *PixbufSimpleAnimClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GdkPixdata
