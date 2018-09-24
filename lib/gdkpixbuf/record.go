// This is a generated file - DO NOT EDIT

package gdkpixbuf

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

func (recv *PixbufFormat) toC() *C.GdkPixbufFormat {

	return recv.native
}

// Unsupported : gdk_pixbuf_format_free : no return generator

// Unsupported : gdk_pixbuf_format_get_extensions : no return type

// Unsupported : gdk_pixbuf_format_get_mime_types : no return type

// Unsupported : gdk_pixbuf_format_is_disabled : no return generator

// Unsupported : gdk_pixbuf_format_is_save_option_supported : no return generator

// Unsupported : gdk_pixbuf_format_is_scalable : no return generator

// Unsupported : gdk_pixbuf_format_is_writable : no return generator

// Unsupported : gdk_pixbuf_format_set_disabled : unsupported parameter disabled : no type generator for gboolean, gboolean

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

func (recv *PixbufLoaderClass) toC() *C.GdkPixbufLoaderClass {

	return recv.native
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

func (recv *PixbufSimpleAnimClass) toC() *C.GdkPixbufSimpleAnimClass {

	return recv.native
}

// Blacklisted : GdkPixdata
