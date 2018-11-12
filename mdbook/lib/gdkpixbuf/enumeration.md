# `gdkpixbuf` Enums

## `Colorspace`

This enumeration defines the color spaces that are supported by
the gdk-pixbuf library.  Currently only RGB is supported.

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

C - `GdkPixbufAlphaMode`

## `PixbufError`

An error code in the #GDK_PIXBUF_ERROR domain. Many gdk-pixbuf
operations can cause errors in this domain, or in the #G_FILE_ERROR
domain.

C - `GdkPixbufError`

## `PixbufRotation`

The possible rotations which can be passed to gdk_pixbuf_rotate_simple().
To make them easier to use, their numerical values are the actual degrees.

C - `GdkPixbufRotation`

