// This is a generated file - DO NOT EDIT

package gdkpixbuf

import glib "github.com/pekim/gobbi/lib/glib"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

type Colorspace int

const (
	GDK_COLORSPACE_RGB Colorspace = 0
)

type InterpType int

const (
	GDK_INTERP_NEAREST  InterpType = 0
	GDK_INTERP_TILES    InterpType = 1
	GDK_INTERP_BILINEAR InterpType = 2
	GDK_INTERP_HYPER    InterpType = 3
)

type PixbufAlphaMode int

const (
	GDK_PIXBUF_ALPHA_BILEVEL PixbufAlphaMode = 0
	GDK_PIXBUF_ALPHA_FULL    PixbufAlphaMode = 1
)

type PixbufError int

const (
	GDK_PIXBUF_ERROR_CORRUPT_IMAGE         PixbufError = 0
	GDK_PIXBUF_ERROR_INSUFFICIENT_MEMORY   PixbufError = 1
	GDK_PIXBUF_ERROR_BAD_OPTION            PixbufError = 2
	GDK_PIXBUF_ERROR_UNKNOWN_TYPE          PixbufError = 3
	GDK_PIXBUF_ERROR_UNSUPPORTED_OPERATION PixbufError = 4
	GDK_PIXBUF_ERROR_FAILED                PixbufError = 5
	GDK_PIXBUF_ERROR_INCOMPLETE_ANIMATION  PixbufError = 6
)

// PixbufErrorQuark is a wrapper around the C function gdk_pixbuf_error_quark.
func PixbufErrorQuark() glib.Quark {
	retC := C.gdk_pixbuf_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type PixbufRotation int

const (
	GDK_PIXBUF_ROTATE_NONE             PixbufRotation = 0
	GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = 90
	GDK_PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = 180
	GDK_PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = 270
)
