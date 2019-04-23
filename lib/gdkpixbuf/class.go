// This is a generated file - DO NOT EDIT

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

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

// Equals compares this Pixbuf with another Pixbuf, and returns true if they represent the same GObject.
func (recv *Pixbuf) Equals(other *Pixbuf) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Pixbuf) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Pixbuf.
// Exercise care, as this is a potentially dangerous function if the Object is not a Pixbuf.
func CastToPixbuf(object *gobject.Object) *Pixbuf {
	return PixbufNewFromC(object.ToC())
}

// Blacklisted : gdk_pixbuf_new

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

// Blacklisted : gdk_pixbuf_new_from_inline

// Blacklisted : gdk_pixbuf_new_from_xpm_data

// gdk_pixbuf_from_pixdata : unsupported parameter pixdata : Blacklisted record : GdkPixdata
// Blacklisted : gdk_pixbuf_add_alpha

// Blacklisted : gdk_pixbuf_composite

// Blacklisted : gdk_pixbuf_composite_color

// Blacklisted : gdk_pixbuf_composite_color_simple

// Blacklisted : gdk_pixbuf_copy

// Blacklisted : gdk_pixbuf_copy_area

// Blacklisted : gdk_pixbuf_fill

// Blacklisted : gdk_pixbuf_get_bits_per_sample

// Blacklisted : gdk_pixbuf_get_colorspace

// Blacklisted : gdk_pixbuf_get_has_alpha

// Blacklisted : gdk_pixbuf_get_height

// Blacklisted : gdk_pixbuf_get_n_channels

// Blacklisted : gdk_pixbuf_get_option

// Unsupported : gdk_pixbuf_get_pixels : array return type :

// Blacklisted : gdk_pixbuf_get_rowstride

// Blacklisted : gdk_pixbuf_get_width

// Blacklisted : gdk_pixbuf_new_subpixbuf

// Blacklisted : gdk_pixbuf_ref

// Blacklisted : gdk_pixbuf_saturate_and_pixelate

// Unsupported : gdk_pixbuf_save : unsupported parameter ... : varargs

// Blacklisted : gdk_pixbuf_savev

// Blacklisted : gdk_pixbuf_scale

// Blacklisted : gdk_pixbuf_scale_simple

// Blacklisted : gdk_pixbuf_unref

// Icon returns the Icon interface implemented by Pixbuf
func (recv *Pixbuf) Icon() *gio.Icon {
	return gio.IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by Pixbuf
func (recv *Pixbuf) LoadableIcon() *gio.LoadableIcon {
	return gio.LoadableIconNewFromC(recv.ToC())
}

// Blacklisted : GdkPixbufAnimation

// Blacklisted : GdkPixbufAnimationIter

// Blacklisted : GdkPixbufLoader

// Blacklisted : GdkPixbufSimpleAnim

// Unsupported : PixbufSimpleAnimIter : no CType
