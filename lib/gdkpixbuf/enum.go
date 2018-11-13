// This is a generated file - DO NOT EDIT

package gdkpixbuf

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

/*
This enumeration defines the color spaces that are supported by
the gdk-pixbuf library.  Currently only RGB is supported.
*/
type Colorspace C.GdkColorspace

const (
	// Indicates a red/green/blue additive color space.
	GDK_COLORSPACE_RGB Colorspace = 0
)

/*
This enumeration describes the different interpolation modes that
 can be used with the scaling functions. @GDK_INTERP_NEAREST is
 the fastest scaling method, but has horrible quality when
 scaling down. @GDK_INTERP_BILINEAR is the best choice if you
 aren't sure what to choose, it has a good speed/quality balance.

 <note>
	Cubic filtering is missing from the list; hyperbolic
	interpolation is just as fast and results in higher quality.
 </note>
*/
type InterpType C.GdkInterpType

const (
	/*
	   Nearest neighbor sampling; this is the fastest
	    and lowest quality mode. Quality is normally unacceptable when scaling
	    down, but may be OK when scaling up.
	*/
	GDK_INTERP_NEAREST InterpType = 0
	/*
	   This is an accurate simulation of the PostScript
	    image operator without any interpolation enabled.  Each pixel is
	    rendered as a tiny parallelogram of solid color, the edges of which
	    are implemented with antialiasing.  It resembles nearest neighbor for
	    enlargement, and bilinear for reduction.
	*/
	GDK_INTERP_TILES InterpType = 1
	/*
	   Best quality/speed balance; use this mode by
	    default. Bilinear interpolation.  For enlargement, it is
	    equivalent to point-sampling the ideal bilinear-interpolated image.
	    For reduction, it is equivalent to laying down small tiles and
	    integrating over the coverage area.
	*/
	GDK_INTERP_BILINEAR InterpType = 2
	/*
	   This is the slowest and highest quality
	    reconstruction function. It is derived from the hyperbolic filters in
	    Wolberg's "Digital Image Warping", and is formally defined as the
	    hyperbolic-filter sampling the ideal hyperbolic-filter interpolated
	    image (the filter is designed to be idempotent for 1:1 pixel mapping).
	*/
	GDK_INTERP_HYPER InterpType = 3
)

/*
These values can be passed to
gdk_pixbuf_xlib_render_to_drawable_alpha() to control how the alpha
channel of an image should be handled.  This function can create a
bilevel clipping mask (black and white) and use it while painting
the image.  In the future, when the X Window System gets an alpha
channel extension, it will be possible to do full alpha
compositing onto arbitrary drawables.  For now both cases fall
back to a bilevel clipping mask.
*/
type PixbufAlphaMode C.GdkPixbufAlphaMode

const (
	/*
	   A bilevel clipping mask (black and white)
	    will be created and used to draw the image.  Pixels below 0.5 opacity
	    will be considered fully transparent, and all others will be
	    considered fully opaque.
	*/
	GDK_PIXBUF_ALPHA_BILEVEL PixbufAlphaMode = 0
	/*
	   For now falls back to #GDK_PIXBUF_ALPHA_BILEVEL.
	    In the future it will do full alpha compositing.
	*/
	GDK_PIXBUF_ALPHA_FULL PixbufAlphaMode = 1
)

/*
An error code in the #GDK_PIXBUF_ERROR domain. Many gdk-pixbuf
operations can cause errors in this domain, or in the #G_FILE_ERROR
domain.
*/
type PixbufError C.GdkPixbufError

const (
	// An image file was broken somehow.
	GDK_PIXBUF_ERROR_CORRUPT_IMAGE PixbufError = 0
	// Not enough memory.
	GDK_PIXBUF_ERROR_INSUFFICIENT_MEMORY PixbufError = 1
	// A bad option was passed to a pixbuf save module.
	GDK_PIXBUF_ERROR_BAD_OPTION PixbufError = 2
	// Unknown image type.
	GDK_PIXBUF_ERROR_UNKNOWN_TYPE PixbufError = 3
	/*
	   Don't know how to perform the
	    given operation on the type of image at hand.
	*/
	GDK_PIXBUF_ERROR_UNSUPPORTED_OPERATION PixbufError = 4
	// Generic failure code, something went wrong.
	GDK_PIXBUF_ERROR_FAILED PixbufError = 5
	// Only part of the animation was loaded.
	GDK_PIXBUF_ERROR_INCOMPLETE_ANIMATION PixbufError = 6
)

/*
The possible rotations which can be passed to gdk_pixbuf_rotate_simple().
To make them easier to use, their numerical values are the actual degrees.
*/
type PixbufRotation C.GdkPixbufRotation

const (
	// No rotation.
	GDK_PIXBUF_ROTATE_NONE PixbufRotation = 0
	// Rotate by 90 degrees.
	GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = 90
	// Rotate by 180 degrees.
	GDK_PIXBUF_ROTATE_UPSIDEDOWN PixbufRotation = 180
	// Rotate by 270 degrees.
	GDK_PIXBUF_ROTATE_CLOCKWISE PixbufRotation = 270
)
