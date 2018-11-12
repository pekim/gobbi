# `gdkpixbuf` Constants

## `PIXBUF_FEATURES_H`



C - `GDK_PIXBUF_FEATURES_H`

## `PIXBUF_MAJOR`

Major version of gdk-pixbuf library, that is the "0" in
"0.8.2" for example.

C - `GDK_PIXBUF_MAJOR`

## `PIXBUF_MICRO`

Micro version of gdk-pixbuf library, that is the "2" in
"0.8.2" for example.

C - `GDK_PIXBUF_MICRO`

## `PIXBUF_MINOR`

Minor version of gdk-pixbuf library, that is the "8" in
"0.8.2" for example.

C - `GDK_PIXBUF_MINOR`

## `PIXBUF_VERSION`

Contains the full version of the gdk-pixbuf header as a string.
This is the version being compiled against; contrast with
&num;gdk_pixbuf_version.

C - `GDK_PIXBUF_VERSION`

ality.
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

