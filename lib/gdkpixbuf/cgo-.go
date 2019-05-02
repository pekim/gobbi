// This is a generated file - DO NOT EDIT

package gdkpixbuf

import (
	glib "github.com/pekim/gobbi/lib/glib"
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

// PixbufNew is a wrapper around the C function gdk_pixbuf_new.
func PixbufNew(colorspace Colorspace, hasAlpha bool, bitsPerSample int32, width int32, height int32) *Pixbuf {
	c_colorspace := (C.GdkColorspace)(colorspace)

	c_has_alpha :=
		boolToGboolean(hasAlpha)

	c_bits_per_sample := (C.int)(bitsPerSample)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	retC := C.gdk_pixbuf_new(c_colorspace, c_has_alpha, c_bits_per_sample, c_width, c_height)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter destroy_fn : no type generator for PixbufDestroyNotify (GdkPixbufDestroyNotify) for param destroy_fn

// PixbufNewFromFile is a wrapper around the C function gdk_pixbuf_new_from_file.
func PixbufNewFromFile(filename string) (*Pixbuf, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_file(c_filename, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PixbufNewFromInline is a wrapper around the C function gdk_pixbuf_new_from_inline.
func PixbufNewFromInline(data []uint8, copyPixels bool) (*Pixbuf, error) {
	c_data_length := (C.gint)(len(data))

	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guint8)(unsafe.Pointer(c_data_arrayPtr))

	c_copy_pixels :=
		boolToGboolean(copyPixels)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_inline(c_data_length, c_data, c_copy_pixels, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PixbufNewFromXpmData is a wrapper around the C function gdk_pixbuf_new_from_xpm_data.
func PixbufNewFromXpmData(data []string) *Pixbuf {
	c_data_array := make([]*C.char, len(data)+1, len(data)+1)
	for i, item := range data {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_data_array[i] = c
	}
	c_data_array[len(data)] = nil
	c_data_arrayPtr := &c_data_array[0]
	c_data := (**C.char)(unsafe.Pointer(c_data_arrayPtr))

	retC := C.gdk_pixbuf_new_from_xpm_data(c_data)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
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

// PixbufAnimationNewFromFile is a wrapper around the C function gdk_pixbuf_animation_new_from_file.
func PixbufAnimationNewFromFile(filename string) (*PixbufAnimation, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_animation_new_from_file(c_filename, &cThrowableError)
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
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

// PixbufLoaderNew is a wrapper around the C function gdk_pixbuf_loader_new.
func PixbufLoaderNew() *PixbufLoader {
	retC := C.gdk_pixbuf_loader_new()
	retGo := PixbufLoaderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// PixbufLoaderNewWithType is a wrapper around the C function gdk_pixbuf_loader_new_with_type.
func PixbufLoaderNewWithType(imageType string) (*PixbufLoader, error) {
	c_image_type := C.CString(imageType)
	defer C.free(unsafe.Pointer(c_image_type))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_new_with_type(c_image_type, &cThrowableError)
	retGo := PixbufLoaderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
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
