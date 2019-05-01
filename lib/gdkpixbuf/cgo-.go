// This is a generated file - DO NOT EDIT

package gdkpixbuf

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GdkPixdataDumpType

// Blacklisted : GdkPixdataType

// Pixbuf is a wrapper around the C record GdkPixbuf.
type Pixbuf struct {
	native *C.GdkPixbuf
}

func PixbufNewFromC(u unsafe.Pointer) *Pixbuf {
	c := (*C.GdkPixbuf)(u)
	if c == nil {
		return nil
	}

	g := &Pixbuf{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Pixbuf) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Pixbuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufAnimation is a wrapper around the C record GdkPixbufAnimation.
type PixbufAnimation struct {
	native *C.GdkPixbufAnimation
}

func PixbufAnimationNewFromC(u unsafe.Pointer) *PixbufAnimation {
	c := (*C.GdkPixbufAnimation)(u)
	if c == nil {
		return nil
	}

	g := &PixbufAnimation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufAnimation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufAnimation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufAnimationIter is a wrapper around the C record GdkPixbufAnimationIter.
type PixbufAnimationIter struct {
	native *C.GdkPixbufAnimationIter
}

func PixbufAnimationIterNewFromC(u unsafe.Pointer) *PixbufAnimationIter {
	c := (*C.GdkPixbufAnimationIter)(u)
	if c == nil {
		return nil
	}

	g := &PixbufAnimationIter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufAnimationIter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufAnimationIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufLoader is a wrapper around the C record GdkPixbufLoader.
type PixbufLoader struct {
	native *C.GdkPixbufLoader
	// parent_instance : record
	// Private : priv
}

func PixbufLoaderNewFromC(u unsafe.Pointer) *PixbufLoader {
	c := (*C.GdkPixbufLoader)(u)
	if c == nil {
		return nil
	}

	g := &PixbufLoader{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufLoader) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufLoader) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufSimpleAnim is a wrapper around the C record GdkPixbufSimpleAnim.
type PixbufSimpleAnim struct {
	native *C.GdkPixbufSimpleAnim
}

func PixbufSimpleAnimNewFromC(u unsafe.Pointer) *PixbufSimpleAnim {
	c := (*C.GdkPixbufSimpleAnim)(u)
	if c == nil {
		return nil
	}

	g := &PixbufSimpleAnim{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufSimpleAnim) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufSimpleAnim) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : PixbufSimpleAnimIter : no CType

// Blacklisted : PIXBUF_MAGIC_NUMBER

// Blacklisted : PIXDATA_HEADER_LENGTH

// Unsupported : gdk_pixbuf_error_quark : return type :

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
