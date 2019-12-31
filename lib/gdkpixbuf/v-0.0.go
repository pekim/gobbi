// Code generated - DO NOT EDIT.
// +build !gdkpixbuf_2.2,!gdkpixbuf_2.4,!gdkpixbuf_2.6,!gdkpixbuf_2.8,!gdkpixbuf_2.12,!gdkpixbuf_2.14,!gdkpixbuf_2.18,!gdkpixbuf_2.22,!gdkpixbuf_2.24,!gdkpixbuf_2.26,!gdkpixbuf_2.28,!gdkpixbuf_2.30,!gdkpixbuf_2.32,!gdkpixbuf_2.36,!gdkpixbuf_2.36.8,!gdkpixbuf_2.40

package gdkpixbuf

import "unsafe"

// PIXBUF_FEATURES_H is a representation of the C constant GDK_PIXBUF_FEATURES_H.
const PIXBUF_FEATURES_H = 1

// PIXBUF_MAJOR is a representation of the C constant GDK_PIXBUF_MAJOR.
const PIXBUF_MAJOR = 2

// PIXBUF_MICRO is a representation of the C constant GDK_PIXBUF_MICRO.
const PIXBUF_MICRO = 0

// PIXBUF_MINOR is a representation of the C constant GDK_PIXBUF_MINOR.
const PIXBUF_MINOR = 40

// PIXBUF_VERSION is a representation of the C constant GDK_PIXBUF_VERSION.
const PIXBUF_VERSION = "2.40.0"

// Colorspace is a representation of the C enumeration GdkColorspace.
type Colorspace int

// Colorspace_rgb is a representation of the C enumeration member GDK_COLORSPACE_RGB.
const Colorspace_rgb = Colorspace(0)

// InterpType is a representation of the C enumeration GdkInterpType.
type InterpType int

// InterpType_nearest is a representation of the C enumeration member GDK_INTERP_NEAREST.
const InterpType_nearest = InterpType(0)

// InterpType_tiles is a representation of the C enumeration member GDK_INTERP_TILES.
const InterpType_tiles = InterpType(1)

// InterpType_bilinear is a representation of the C enumeration member GDK_INTERP_BILINEAR.
const InterpType_bilinear = InterpType(2)

// InterpType_hyper is a representation of the C enumeration member GDK_INTERP_HYPER.
const InterpType_hyper = InterpType(3)

// PixbufAlphaMode is a representation of the C enumeration GdkPixbufAlphaMode.
type PixbufAlphaMode int

// PixbufAlphaMode_bilevel is a representation of the C enumeration member GDK_PIXBUF_ALPHA_BILEVEL.
const PixbufAlphaMode_bilevel = PixbufAlphaMode(0)

// PixbufAlphaMode_full is a representation of the C enumeration member GDK_PIXBUF_ALPHA_FULL.
const PixbufAlphaMode_full = PixbufAlphaMode(1)

// PixbufError is a representation of the C enumeration GdkPixbufError.
type PixbufError int

// PixbufError_corrupt_image is a representation of the C enumeration member GDK_PIXBUF_ERROR_CORRUPT_IMAGE.
const PixbufError_corrupt_image = PixbufError(0)

// PixbufError_insufficient_memory is a representation of the C enumeration member GDK_PIXBUF_ERROR_INSUFFICIENT_MEMORY.
const PixbufError_insufficient_memory = PixbufError(1)

// PixbufError_bad_option is a representation of the C enumeration member GDK_PIXBUF_ERROR_BAD_OPTION.
const PixbufError_bad_option = PixbufError(2)

// PixbufError_unknown_type is a representation of the C enumeration member GDK_PIXBUF_ERROR_UNKNOWN_TYPE.
const PixbufError_unknown_type = PixbufError(3)

// PixbufError_unsupported_operation is a representation of the C enumeration member GDK_PIXBUF_ERROR_UNSUPPORTED_OPERATION.
const PixbufError_unsupported_operation = PixbufError(4)

// PixbufError_failed is a representation of the C enumeration member GDK_PIXBUF_ERROR_FAILED.
const PixbufError_failed = PixbufError(5)

// PixbufError_incomplete_animation is a representation of the C enumeration member GDK_PIXBUF_ERROR_INCOMPLETE_ANIMATION.
const PixbufError_incomplete_animation = PixbufError(6)

// PixbufRotation is a representation of the C enumeration GdkPixbufRotation.
type PixbufRotation int

// PixbufRotation_none is a representation of the C enumeration member GDK_PIXBUF_ROTATE_NONE.
const PixbufRotation_none = PixbufRotation(0)

// PixbufRotation_counterclockwise is a representation of the C enumeration member GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE.
const PixbufRotation_counterclockwise = PixbufRotation(90)

// PixbufRotation_upsidedown is a representation of the C enumeration member GDK_PIXBUF_ROTATE_UPSIDEDOWN.
const PixbufRotation_upsidedown = PixbufRotation(180)

// PixbufRotation_clockwise is a representation of the C enumeration member GDK_PIXBUF_ROTATE_CLOCKWISE.
const PixbufRotation_clockwise = PixbufRotation(270)

// PixbufFormat is a representation of the C record GdkPixbufFormat.
type PixbufFormat struct {
	native unsafe.Pointer
}

// PixbufLoaderClass is a representation of the C record GdkPixbufLoaderClass.
type PixbufLoaderClass struct {
	native unsafe.Pointer
}

// PixbufSimpleAnimClass is a representation of the C record GdkPixbufSimpleAnimClass.
type PixbufSimpleAnimClass struct {
	native unsafe.Pointer
}
