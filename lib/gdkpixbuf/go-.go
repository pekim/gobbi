// This is a generated file - DO NOT EDIT

package gdkpixbuf

import gio "github.com/pekim/gobbi/lib/gio"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

// Blacklisted : GdkPixdataDumpType

// Blacklisted : GdkPixdataType

// Unsupported : gdk_pixbuf_new : return type :

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter destroy_fn : no type generator for PixbufDestroyNotify (GdkPixbufDestroyNotify) for param destroy_fn

// Unsupported : gdk_pixbuf_new_from_file : return type :

// Unsupported : gdk_pixbuf_new_from_inline : return type :

// Unsupported : gdk_pixbuf_new_from_xpm_data : return type :

// Icon returns the Icon interface implemented by Pixbuf
func (recv *Pixbuf) Icon() *gio.Icon {
	return gio.IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by Pixbuf
func (recv *Pixbuf) LoadableIcon() *gio.LoadableIcon {
	return gio.LoadableIconNewFromC(recv.ToC())
}

// Unsupported : gdk_pixbuf_animation_new_from_file : return type :

// Unsupported : gdk_pixbuf_loader_new : return type :

// Unsupported : gdk_pixbuf_loader_new_with_type : return type :

// Unsupported : PixbufSimpleAnimIter : no CType

const PIXBUF_FEATURES_H int32 = 1

// Blacklisted : PIXBUF_MAGIC_NUMBER

const PIXBUF_MAJOR int32 = 2
const PIXBUF_MICRO int32 = 11
const PIXBUF_MINOR int32 = 36
const PIXBUF_VERSION string = "2.36.11"

// Blacklisted : PIXDATA_HEADER_LENGTH

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

// gdk_pixbuf_error_quark : return type :
type PixbufRotation int

const (
	GDK_PIXBUF_ROTATE_NONE             PixbufRotation = 0
	GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = 90
	GDK_PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = 180
	GDK_PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = 270
)

// Unsupported : gdk_pixbuf_error_quark : return type :

// Blacklisted : GdkPixdata
