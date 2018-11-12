# `gdkpixbuf` enums

## `Colorspace`

This enumeration defines the color spaces that are supported by
the gdk-pixbuf library.  Currently only RGB is supported.

- GDK_COLORSPACE_RGB = 0

C - `GdkColorspace`

## `InterpType`

This enumeration describes the different interpolation modes that
 can be used with the scaling functions. @GDK_INTERP_NEAREST is
 the fastest scaling method, but has horrible quality when
 scaling down. @GDK_INTERP_BILINEAR is the best choice if you
 aren't sure what to choose, it has a good speed/quality balance.

 <note>
	Cubic filtering is missing from the list; hyperbolic
	interpolation is just as fast and results in higher quality.
 </note>

- GDK_INTERP_NEAREST = 0
- GDK_INTERP_TILES = 1
- GDK_INTERP_BILINEAR = 2
- GDK_INTERP_HYPER = 3

C - `GdkInterpType`

## `PixbufAlphaMode`

These values can be passed to
gdk_pixbuf_xlib_render_to_drawable_alpha() to control how the alpha
channel of an image should be handled.  This function can create a
bilevel clipping mask (black and white) and use it while painting
the image.  In the future, when the X Window System gets an alpha
channel extension, it will be possible to do full alpha
compositing onto arbitrary drawables.  For now both cases fall
back to a bilevel clipping mask.

- GDK_PIXBUF_ALPHA_BILEVEL = 0
- GDK_PIXBUF_ALPHA_FULL = 1

C - `GdkPixbufAlphaMode`

## `PixbufError`

An error code in the #GDK_PIXBUF_ERROR domain. Many gdk-pixbuf
operations can cause errors in this domain, or in the #G_FILE_ERROR
domain.

- GDK_PIXBUF_ERROR_CORRUPT_IMAGE = 0
- GDK_PIXBUF_ERROR_INSUFFICIENT_MEMORY = 1
- GDK_PIXBUF_ERROR_BAD_OPTION = 2
- GDK_PIXBUF_ERROR_UNKNOWN_TYPE = 3
- GDK_PIXBUF_ERROR_UNSUPPORTED_OPERATION = 4
- GDK_PIXBUF_ERROR_FAILED = 5
- GDK_PIXBUF_ERROR_INCOMPLETE_ANIMATION = 6

C - `GdkPixbufError`

## `PixbufRotation`

The possible rotations which can be passed to gdk_pixbuf_rotate_simple().
To make them easier to use, their numerical values are the actual degrees.

- GDK_PIXBUF_ROTATE_NONE = 0
- GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE = 90
- GDK_PIXBUF_ROTATE_UPSIDEDOWN = 180
- GDK_PIXBUF_ROTATE_CLOCKWISE = 270

C - `GdkPixbufRotation`

E = %!s(int=0)
- GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE = %!s(int=90)
- GDK_PIXBUF_ROTATE_UPSIDEDOWN = %!s(int=180)
- GDK_PIXBUF_ROTATE_CLOCKWISE = %!s(int=270)
