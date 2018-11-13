// This is a generated file - DO NOT EDIT

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
/*

	void pixbufloader_areaPreparedHandler(GObject *, gpointer);

	static gulong PixbufLoader_signal_connect_area_prepared(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "area-prepared", G_CALLBACK(pixbufloader_areaPreparedHandler), data);
	}

*/
/*

	void pixbufloader_closedHandler(GObject *, gpointer);

	static gulong PixbufLoader_signal_connect_closed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closed", G_CALLBACK(pixbufloader_closedHandler), data);
	}

*/
import "C"

// This is the main structure in the gdk-pixbuf library.  It is
// used to represent images.  It contains information about the
// image's pixel data, its color space, bits per sample, width and
// height, and the rowstride (the number of bytes between the start of
// one row and the start of the next).
/*

C record/class : GdkPixbuf
*/
type Pixbuf struct {
	native *C.GdkPixbuf
}

func PixbufNewFromC(u unsafe.Pointer) *Pixbuf {
	c := (*C.GdkPixbuf)(u)
	if c == nil {
		return nil
	}

	g := &Pixbuf{native: c}

	return g
}

func (recv *Pixbuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Pixbuf) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Pixbuf.
// Exercise care, as this is a potentially dangerous function if the Object is not a Pixbuf.
func CastToPixbuf(object *gobject.Object) *Pixbuf {
	return PixbufNewFromC(object.ToC())
}

// Creates a new #GdkPixbuf structure and allocates a buffer for it.  The
// buffer has an optimal rowstride.  Note that the buffer is not cleared;
// you will have to fill it completely yourself.
/*

C function : gdk_pixbuf_new
*/
func PixbufNew(colorspace Colorspace, hasAlpha bool, bitsPerSample int32, width int32, height int32) *Pixbuf {
	c_colorspace := (C.GdkColorspace)(colorspace)

	c_has_alpha :=
		boolToGboolean(hasAlpha)

	c_bits_per_sample := (C.int)(bitsPerSample)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	retC := C.gdk_pixbuf_new(c_colorspace, c_has_alpha, c_bits_per_sample, c_width, c_height)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter destroy_fn : no type generator for PixbufDestroyNotify (GdkPixbufDestroyNotify) for param destroy_fn

// Creates a new pixbuf by loading an image from a file.  The file format is
// detected automatically. If %NULL is returned, then @error will be set.
// Possible errors are in the #GDK_PIXBUF_ERROR and #G_FILE_ERROR domains.
/*

C function : gdk_pixbuf_new_from_file
*/
func PixbufNewFromFile(filename string) (*Pixbuf, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_file(c_filename, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Create a #GdkPixbuf from a flat representation that is suitable for
// storing as inline data in a program. This is useful if you want to
// ship a program with images, but don't want to depend on any
// external files.
//
// gdk-pixbuf ships with a program called [gdk-pixbuf-csource][gdk-pixbuf-csource],
// which allows for conversion of #GdkPixbufs into such a inline representation.
// In almost all cases, you should pass the `--raw` option to
// `gdk-pixbuf-csource`. A sample invocation would be:
//
// |[
// gdk-pixbuf-csource --raw --name=myimage_inline myimage.png
// ]|
//
// For the typical case where the inline pixbuf is read-only static data,
// you don't need to copy the pixel data unless you intend to write to
// it, so you can pass %FALSE for @copy_pixels.  (If you pass `--rle` to
// `gdk-pixbuf-csource`, a copy will be made even if @copy_pixels is %FALSE,
// so using this option is generally a bad idea.)
//
// If you create a pixbuf from const inline data compiled into your
// program, it's probably safe to ignore errors and disable length checks,
// since things will always succeed:
// |[
// pixbuf = gdk_pixbuf_new_from_inline (-1, myimage_inline, FALSE, NULL);
// ]|
//
// For non-const inline data, you could get out of memory. For untrusted
// inline data located at runtime, you could have corrupt inline data in
// addition.
/*

C function : gdk_pixbuf_new_from_inline
*/
func PixbufNewFromInline(data []uint8, copyPixels bool) (*Pixbuf, error) {
	c_data_length := (C.gint)(len(data))

	c_data := &data[0]

	c_copy_pixels :=
		boolToGboolean(copyPixels)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_inline(c_data_length, (*C.guint8)(unsafe.Pointer(c_data)), c_copy_pixels, &cThrowableError)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data :

// Takes an existing pixbuf and adds an alpha channel to it.
// If the existing pixbuf already had an alpha channel, the channel
// values are copied from the original; otherwise, the alpha channel
// is initialized to 255 (full opacity).
//
// If @substitute_color is %TRUE, then the color specified by (@r, @g, @b) will be
// assigned zero opacity. That is, if you pass (255, 255, 255) for the
// substitute color, all white pixels will become fully transparent.
/*

C function : gdk_pixbuf_add_alpha
*/
func (recv *Pixbuf) AddAlpha(substituteColor bool, r uint8, g uint8, b uint8) *Pixbuf {
	c_substitute_color :=
		boolToGboolean(substituteColor)

	c_r := (C.guchar)(r)

	c_g := (C.guchar)(g)

	c_b := (C.guchar)(b)

	retC := C.gdk_pixbuf_add_alpha((*C.GdkPixbuf)(recv.native), c_substitute_color, c_r, c_g, c_b)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a transformation of the source image @src by scaling by
// @scale_x and @scale_y then translating by @offset_x and @offset_y.
// This gives an image in the coordinates of the destination pixbuf.
// The rectangle (@dest_x, @dest_y, @dest_width, @dest_height)
// is then alpha blended onto the corresponding rectangle of the
// original destination image.
//
// When the destination rectangle contains parts not in the source
// image, the data at the edges of the source image is replicated
// to infinity.
//
// ![](composite.png)
/*

C function : gdk_pixbuf_composite
*/
func (recv *Pixbuf) Composite(dest *Pixbuf, destX int32, destY int32, destWidth int32, destHeight int32, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int32) {
	c_dest := (*C.GdkPixbuf)(C.NULL)
	if dest != nil {
		c_dest = (*C.GdkPixbuf)(dest.ToC())
	}

	c_dest_x := (C.int)(destX)

	c_dest_y := (C.int)(destY)

	c_dest_width := (C.int)(destWidth)

	c_dest_height := (C.int)(destHeight)

	c_offset_x := (C.double)(offsetX)

	c_offset_y := (C.double)(offsetY)

	c_scale_x := (C.double)(scaleX)

	c_scale_y := (C.double)(scaleY)

	c_interp_type := (C.GdkInterpType)(interpType)

	c_overall_alpha := (C.int)(overallAlpha)

	C.gdk_pixbuf_composite((*C.GdkPixbuf)(recv.native), c_dest, c_dest_x, c_dest_y, c_dest_width, c_dest_height, c_offset_x, c_offset_y, c_scale_x, c_scale_y, c_interp_type, c_overall_alpha)

	return
}

// Creates a transformation of the source image @src by scaling by
// @scale_x and @scale_y then translating by @offset_x and @offset_y,
// then alpha blends the rectangle (@dest_x ,@dest_y, @dest_width,
// @dest_height) of the resulting image with a checkboard of the
// colors @color1 and @color2 and renders it onto the destination
// image.
//
// If the source image has no alpha channel, and @overall_alpha is 255, a fast
// path is used which omits the alpha blending and just performs the scaling.
//
// See gdk_pixbuf_composite_color_simple() for a simpler variant of this
// function suitable for many tasks.
/*

C function : gdk_pixbuf_composite_color
*/
func (recv *Pixbuf) CompositeColor(dest *Pixbuf, destX int32, destY int32, destWidth int32, destHeight int32, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int32, checkX int32, checkY int32, checkSize int32, color1 uint32, color2 uint32) {
	c_dest := (*C.GdkPixbuf)(C.NULL)
	if dest != nil {
		c_dest = (*C.GdkPixbuf)(dest.ToC())
	}

	c_dest_x := (C.int)(destX)

	c_dest_y := (C.int)(destY)

	c_dest_width := (C.int)(destWidth)

	c_dest_height := (C.int)(destHeight)

	c_offset_x := (C.double)(offsetX)

	c_offset_y := (C.double)(offsetY)

	c_scale_x := (C.double)(scaleX)

	c_scale_y := (C.double)(scaleY)

	c_interp_type := (C.GdkInterpType)(interpType)

	c_overall_alpha := (C.int)(overallAlpha)

	c_check_x := (C.int)(checkX)

	c_check_y := (C.int)(checkY)

	c_check_size := (C.int)(checkSize)

	c_color1 := (C.guint32)(color1)

	c_color2 := (C.guint32)(color2)

	C.gdk_pixbuf_composite_color((*C.GdkPixbuf)(recv.native), c_dest, c_dest_x, c_dest_y, c_dest_width, c_dest_height, c_offset_x, c_offset_y, c_scale_x, c_scale_y, c_interp_type, c_overall_alpha, c_check_x, c_check_y, c_check_size, c_color1, c_color2)

	return
}

// Creates a new #GdkPixbuf by scaling @src to @dest_width x
// @dest_height and alpha blending the result with a checkboard of colors
// @color1 and @color2.
/*

C function : gdk_pixbuf_composite_color_simple
*/
func (recv *Pixbuf) CompositeColorSimple(destWidth int32, destHeight int32, interpType InterpType, overallAlpha int32, checkSize int32, color1 uint32, color2 uint32) *Pixbuf {
	c_dest_width := (C.int)(destWidth)

	c_dest_height := (C.int)(destHeight)

	c_interp_type := (C.GdkInterpType)(interpType)

	c_overall_alpha := (C.int)(overallAlpha)

	c_check_size := (C.int)(checkSize)

	c_color1 := (C.guint32)(color1)

	c_color2 := (C.guint32)(color2)

	retC := C.gdk_pixbuf_composite_color_simple((*C.GdkPixbuf)(recv.native), c_dest_width, c_dest_height, c_interp_type, c_overall_alpha, c_check_size, c_color1, c_color2)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GdkPixbuf with a copy of the information in the specified
// @pixbuf. Note that this does not copy the options set on the original #GdkPixbuf,
// use gdk_pixbuf_copy_options() for this.
/*

C function : gdk_pixbuf_copy
*/
func (recv *Pixbuf) Copy() *Pixbuf {
	retC := C.gdk_pixbuf_copy((*C.GdkPixbuf)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copies a rectangular area from @src_pixbuf to @dest_pixbuf.  Conversion of
// pixbuf formats is done automatically.
//
// If the source rectangle overlaps the destination rectangle on the
// same pixbuf, it will be overwritten during the copy operation.
// Therefore, you can not use this function to scroll a pixbuf.
/*

C function : gdk_pixbuf_copy_area
*/
func (recv *Pixbuf) CopyArea(srcX int32, srcY int32, width int32, height int32, destPixbuf *Pixbuf, destX int32, destY int32) {
	c_src_x := (C.int)(srcX)

	c_src_y := (C.int)(srcY)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_dest_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if destPixbuf != nil {
		c_dest_pixbuf = (*C.GdkPixbuf)(destPixbuf.ToC())
	}

	c_dest_x := (C.int)(destX)

	c_dest_y := (C.int)(destY)

	C.gdk_pixbuf_copy_area((*C.GdkPixbuf)(recv.native), c_src_x, c_src_y, c_width, c_height, c_dest_pixbuf, c_dest_x, c_dest_y)

	return
}

// Clears a pixbuf to the given RGBA value, converting the RGBA value into
// the pixbuf's pixel format. The alpha will be ignored if the pixbuf
// doesn't have an alpha channel.
/*

C function : gdk_pixbuf_fill
*/
func (recv *Pixbuf) Fill(pixel uint32) {
	c_pixel := (C.guint32)(pixel)

	C.gdk_pixbuf_fill((*C.GdkPixbuf)(recv.native), c_pixel)

	return
}

// Queries the number of bits per color sample in a pixbuf.
/*

C function : gdk_pixbuf_get_bits_per_sample
*/
func (recv *Pixbuf) GetBitsPerSample() int32 {
	retC := C.gdk_pixbuf_get_bits_per_sample((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Queries the color space of a pixbuf.
/*

C function : gdk_pixbuf_get_colorspace
*/
func (recv *Pixbuf) GetColorspace() Colorspace {
	retC := C.gdk_pixbuf_get_colorspace((*C.GdkPixbuf)(recv.native))
	retGo := (Colorspace)(retC)

	return retGo
}

// Queries whether a pixbuf has an alpha channel (opacity information).
/*

C function : gdk_pixbuf_get_has_alpha
*/
func (recv *Pixbuf) GetHasAlpha() bool {
	retC := C.gdk_pixbuf_get_has_alpha((*C.GdkPixbuf)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Queries the height of a pixbuf.
/*

C function : gdk_pixbuf_get_height
*/
func (recv *Pixbuf) GetHeight() int32 {
	retC := C.gdk_pixbuf_get_height((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Queries the number of channels of a pixbuf.
/*

C function : gdk_pixbuf_get_n_channels
*/
func (recv *Pixbuf) GetNChannels() int32 {
	retC := C.gdk_pixbuf_get_n_channels((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Looks up @key in the list of options that may have been attached to the
// @pixbuf when it was loaded, or that may have been attached by another
// function using gdk_pixbuf_set_option().
//
// For instance, the ANI loader provides "Title" and "Artist" options.
// The ICO, XBM, and XPM loaders provide "x_hot" and "y_hot" hot-spot
// options for cursor definitions. The PNG loader provides the tEXt ancillary
// chunk key/value pairs as options. Since 2.12, the TIFF and JPEG loaders
// return an "orientation" option string that corresponds to the embedded
// TIFF/Exif orientation tag (if present). Since 2.32, the TIFF loader sets
// the "multipage" option string to "yes" when a multi-page TIFF is loaded.
// Since 2.32 the JPEG and PNG loaders set "x-dpi" and "y-dpi" if the file
// contains image density information in dots per inch.
// Since 2.36.6, the JPEG loader sets the "comment" option with the comment
// EXIF tag.
/*

C function : gdk_pixbuf_get_option
*/
func (recv *Pixbuf) GetOption(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gdk_pixbuf_get_option((*C.GdkPixbuf)(recv.native), c_key)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gdk_pixbuf_get_pixels : no return type

// Queries the rowstride of a pixbuf, which is the number of bytes between
// the start of a row and the start of the next row.
/*

C function : gdk_pixbuf_get_rowstride
*/
func (recv *Pixbuf) GetRowstride() int32 {
	retC := C.gdk_pixbuf_get_rowstride((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Queries the width of a pixbuf.
/*

C function : gdk_pixbuf_get_width
*/
func (recv *Pixbuf) GetWidth() int32 {
	retC := C.gdk_pixbuf_get_width((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Creates a new pixbuf which represents a sub-region of @src_pixbuf.
// The new pixbuf shares its pixels with the original pixbuf, so
// writing to one affects both.  The new pixbuf holds a reference to
// @src_pixbuf, so @src_pixbuf will not be finalized until the new
// pixbuf is finalized.
//
// Note that if @src_pixbuf is read-only, this function will force it
// to be mutable.
/*

C function : gdk_pixbuf_new_subpixbuf
*/
func (recv *Pixbuf) NewSubpixbuf(srcX int32, srcY int32, width int32, height int32) *Pixbuf {
	c_src_x := (C.int)(srcX)

	c_src_y := (C.int)(srcY)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	retC := C.gdk_pixbuf_new_subpixbuf((*C.GdkPixbuf)(recv.native), c_src_x, c_src_y, c_width, c_height)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a reference to a pixbuf.
/*

C function : gdk_pixbuf_ref
*/
func (recv *Pixbuf) Ref() *Pixbuf {
	retC := C.gdk_pixbuf_ref((*C.GdkPixbuf)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Modifies saturation and optionally pixelates @src, placing the result in
// @dest. @src and @dest may be the same pixbuf with no ill effects.  If
// @saturation is 1.0 then saturation is not changed. If it's less than 1.0,
// saturation is reduced (the image turns toward grayscale); if greater than
// 1.0, saturation is increased (the image gets more vivid colors). If @pixelate
// is %TRUE, then pixels are faded in a checkerboard pattern to create a
// pixelated image. @src and @dest must have the same image format, size, and
// rowstride.
/*

C function : gdk_pixbuf_saturate_and_pixelate
*/
func (recv *Pixbuf) SaturateAndPixelate(dest *Pixbuf, saturation float32, pixelate bool) {
	c_dest := (*C.GdkPixbuf)(C.NULL)
	if dest != nil {
		c_dest = (*C.GdkPixbuf)(dest.ToC())
	}

	c_saturation := (C.gfloat)(saturation)

	c_pixelate :=
		boolToGboolean(pixelate)

	C.gdk_pixbuf_saturate_and_pixelate((*C.GdkPixbuf)(recv.native), c_dest, c_saturation, c_pixelate)

	return
}

// Unsupported : gdk_pixbuf_save : unsupported parameter ... : varargs

// Unsupported : gdk_pixbuf_savev : unsupported parameter option_keys :

// Creates a transformation of the source image @src by scaling by
// @scale_x and @scale_y then translating by @offset_x and @offset_y,
// then renders the rectangle (@dest_x, @dest_y, @dest_width,
// @dest_height) of the resulting image onto the destination image
// replacing the previous contents.
//
// Try to use gdk_pixbuf_scale_simple() first, this function is
// the industrial-strength power tool you can fall back to if
// gdk_pixbuf_scale_simple() isn't powerful enough.
//
// If the source rectangle overlaps the destination rectangle on the
// same pixbuf, it will be overwritten during the scaling which
// results in rendering artifacts.
/*

C function : gdk_pixbuf_scale
*/
func (recv *Pixbuf) Scale(dest *Pixbuf, destX int32, destY int32, destWidth int32, destHeight int32, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType) {
	c_dest := (*C.GdkPixbuf)(C.NULL)
	if dest != nil {
		c_dest = (*C.GdkPixbuf)(dest.ToC())
	}

	c_dest_x := (C.int)(destX)

	c_dest_y := (C.int)(destY)

	c_dest_width := (C.int)(destWidth)

	c_dest_height := (C.int)(destHeight)

	c_offset_x := (C.double)(offsetX)

	c_offset_y := (C.double)(offsetY)

	c_scale_x := (C.double)(scaleX)

	c_scale_y := (C.double)(scaleY)

	c_interp_type := (C.GdkInterpType)(interpType)

	C.gdk_pixbuf_scale((*C.GdkPixbuf)(recv.native), c_dest, c_dest_x, c_dest_y, c_dest_width, c_dest_height, c_offset_x, c_offset_y, c_scale_x, c_scale_y, c_interp_type)

	return
}

// Create a new #GdkPixbuf containing a copy of @src scaled to
// @dest_width x @dest_height. Leaves @src unaffected.  @interp_type
// should be #GDK_INTERP_NEAREST if you want maximum speed (but when
// scaling down #GDK_INTERP_NEAREST is usually unusably ugly).  The
// default @interp_type should be #GDK_INTERP_BILINEAR which offers
// reasonable quality and speed.
//
// You can scale a sub-portion of @src by creating a sub-pixbuf
// pointing into @src; see gdk_pixbuf_new_subpixbuf().
//
// If @dest_width and @dest_height are equal to the @src width and height, a
// copy of @src is returned, avoiding any scaling.
//
// For more complicated scaling/alpha blending see gdk_pixbuf_scale()
// and gdk_pixbuf_composite().
/*

C function : gdk_pixbuf_scale_simple
*/
func (recv *Pixbuf) ScaleSimple(destWidth int32, destHeight int32, interpType InterpType) *Pixbuf {
	c_dest_width := (C.int)(destWidth)

	c_dest_height := (C.int)(destHeight)

	c_interp_type := (C.GdkInterpType)(interpType)

	retC := C.gdk_pixbuf_scale_simple((*C.GdkPixbuf)(recv.native), c_dest_width, c_dest_height, c_interp_type)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes a reference from a pixbuf.
/*

C function : gdk_pixbuf_unref
*/
func (recv *Pixbuf) Unref() {
	C.gdk_pixbuf_unref((*C.GdkPixbuf)(recv.native))

	return
}

// Icon returns the Icon interface implemented by Pixbuf
func (recv *Pixbuf) Icon() *gio.Icon {
	return gio.IconNewFromC(recv.ToC())
}

// LoadableIcon returns the LoadableIcon interface implemented by Pixbuf
func (recv *Pixbuf) LoadableIcon() *gio.LoadableIcon {
	return gio.LoadableIconNewFromC(recv.ToC())
}

// An opaque struct representing an animation.
/*

C record/class : GdkPixbufAnimation
*/
type PixbufAnimation struct {
	native *C.GdkPixbufAnimation
}

func PixbufAnimationNewFromC(u unsafe.Pointer) *PixbufAnimation {
	c := (*C.GdkPixbufAnimation)(u)
	if c == nil {
		return nil
	}

	g := &PixbufAnimation{native: c}

	return g
}

func (recv *PixbufAnimation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *PixbufAnimation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to PixbufAnimation.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufAnimation.
func CastToPixbufAnimation(object *gobject.Object) *PixbufAnimation {
	return PixbufAnimationNewFromC(object.ToC())
}

// Creates a new animation by loading it from a file. The file format is
// detected automatically. If the file's format does not support multi-frame
// images, then an animation with a single frame will be created. Possible errors
// are in the #GDK_PIXBUF_ERROR and #G_FILE_ERROR domains.
/*

C function : gdk_pixbuf_animation_new_from_file
*/
func PixbufAnimationNewFromFile(filename string) (*PixbufAnimation, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_animation_new_from_file(c_filename, &cThrowableError)
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Queries the height of the bounding box of a pixbuf animation.
/*

C function : gdk_pixbuf_animation_get_height
*/
func (recv *PixbufAnimation) GetHeight() int32 {
	retC := C.gdk_pixbuf_animation_get_height((*C.GdkPixbufAnimation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Get an iterator for displaying an animation. The iterator provides
// the frames that should be displayed at a given time. It should be
// freed after use with g_object_unref().
//
// @start_time would normally come from g_get_current_time(), and marks
// the beginning of animation playback. After creating an iterator, you
// should immediately display the pixbuf returned by
// gdk_pixbuf_animation_iter_get_pixbuf(). Then, you should install
// a timeout (with g_timeout_add()) or by some other mechanism ensure
// that you'll update the image after
// gdk_pixbuf_animation_iter_get_delay_time() milliseconds. Each time
// the image is updated, you should reinstall the timeout with the new,
// possibly-changed delay time.
//
// As a shortcut, if @start_time is %NULL, the result of
// g_get_current_time() will be used automatically.
//
// To update the image (i.e. possibly change the result of
// gdk_pixbuf_animation_iter_get_pixbuf() to a new frame of the animation),
// call gdk_pixbuf_animation_iter_advance().
//
// If you're using #GdkPixbufLoader, in addition to updating the image
// after the delay time, you should also update it whenever you
// receive the area_updated signal and
// gdk_pixbuf_animation_iter_on_currently_loading_frame() returns
// %TRUE. In this case, the frame currently being fed into the loader
// has received new data, so needs to be refreshed. The delay time for
// a frame may also be modified after an area_updated signal, for
// example if the delay time for a frame is encoded in the data after
// the frame itself. So your timeout should be reinstalled after any
// area_updated signal.
//
// A delay time of -1 is possible, indicating "infinite."
/*

C function : gdk_pixbuf_animation_get_iter
*/
func (recv *PixbufAnimation) GetIter(startTime *glib.TimeVal) *PixbufAnimationIter {
	c_start_time := (*C.GTimeVal)(C.NULL)
	if startTime != nil {
		c_start_time = (*C.GTimeVal)(startTime.ToC())
	}

	retC := C.gdk_pixbuf_animation_get_iter((*C.GdkPixbufAnimation)(recv.native), c_start_time)
	retGo := PixbufAnimationIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// If an animation is really just a plain image (has only one frame),
// this function returns that image. If the animation is an animation,
// this function returns a reasonable thing to display as a static
// unanimated image, which might be the first frame, or something more
// sophisticated. If an animation hasn't loaded any frames yet, this
// function will return %NULL.
/*

C function : gdk_pixbuf_animation_get_static_image
*/
func (recv *PixbufAnimation) GetStaticImage() *Pixbuf {
	retC := C.gdk_pixbuf_animation_get_static_image((*C.GdkPixbufAnimation)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Queries the width of the bounding box of a pixbuf animation.
/*

C function : gdk_pixbuf_animation_get_width
*/
func (recv *PixbufAnimation) GetWidth() int32 {
	retC := C.gdk_pixbuf_animation_get_width((*C.GdkPixbufAnimation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// If you load a file with gdk_pixbuf_animation_new_from_file() and it
// turns out to be a plain, unanimated image, then this function will
// return %TRUE. Use gdk_pixbuf_animation_get_static_image() to retrieve
// the image.
/*

C function : gdk_pixbuf_animation_is_static_image
*/
func (recv *PixbufAnimation) IsStaticImage() bool {
	retC := C.gdk_pixbuf_animation_is_static_image((*C.GdkPixbufAnimation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Adds a reference to an animation.
/*

C function : gdk_pixbuf_animation_ref
*/
func (recv *PixbufAnimation) Ref() *PixbufAnimation {
	retC := C.gdk_pixbuf_animation_ref((*C.GdkPixbufAnimation)(recv.native))
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes a reference from an animation.
/*

C function : gdk_pixbuf_animation_unref
*/
func (recv *PixbufAnimation) Unref() {
	C.gdk_pixbuf_animation_unref((*C.GdkPixbufAnimation)(recv.native))

	return
}

// An opaque struct representing an iterator which points to a
// certain position in an animation.
/*

C record/class : GdkPixbufAnimationIter
*/
type PixbufAnimationIter struct {
	native *C.GdkPixbufAnimationIter
}

func PixbufAnimationIterNewFromC(u unsafe.Pointer) *PixbufAnimationIter {
	c := (*C.GdkPixbufAnimationIter)(u)
	if c == nil {
		return nil
	}

	g := &PixbufAnimationIter{native: c}

	return g
}

func (recv *PixbufAnimationIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *PixbufAnimationIter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to PixbufAnimationIter.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufAnimationIter.
func CastToPixbufAnimationIter(object *gobject.Object) *PixbufAnimationIter {
	return PixbufAnimationIterNewFromC(object.ToC())
}

// Possibly advances an animation to a new frame. Chooses the frame based
// on the start time passed to gdk_pixbuf_animation_get_iter().
//
// @current_time would normally come from g_get_current_time(), and
// must be greater than or equal to the time passed to
// gdk_pixbuf_animation_get_iter(), and must increase or remain
// unchanged each time gdk_pixbuf_animation_iter_get_pixbuf() is
// called. That is, you can't go backward in time; animations only
// play forward.
//
// As a shortcut, pass %NULL for the current time and g_get_current_time()
// will be invoked on your behalf. So you only need to explicitly pass
// @current_time if you're doing something odd like playing the animation
// at double speed.
//
// If this function returns %FALSE, there's no need to update the animation
// display, assuming the display had been rendered prior to advancing;
// if %TRUE, you need to call gdk_pixbuf_animation_iter_get_pixbuf()
// and update the display with the new pixbuf.
/*

C function : gdk_pixbuf_animation_iter_advance
*/
func (recv *PixbufAnimationIter) Advance(currentTime *glib.TimeVal) bool {
	c_current_time := (*C.GTimeVal)(C.NULL)
	if currentTime != nil {
		c_current_time = (*C.GTimeVal)(currentTime.ToC())
	}

	retC := C.gdk_pixbuf_animation_iter_advance((*C.GdkPixbufAnimationIter)(recv.native), c_current_time)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the number of milliseconds the current pixbuf should be displayed,
// or -1 if the current pixbuf should be displayed forever. g_timeout_add()
// conveniently takes a timeout in milliseconds, so you can use a timeout
// to schedule the next update.
//
// Note that some formats, like GIF, might clamp the timeout values in the
// image file to avoid updates that are just too quick. The minimum timeout
// for GIF images is currently 20 milliseconds.
/*

C function : gdk_pixbuf_animation_iter_get_delay_time
*/
func (recv *PixbufAnimationIter) GetDelayTime() int32 {
	retC := C.gdk_pixbuf_animation_iter_get_delay_time((*C.GdkPixbufAnimationIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the current pixbuf which should be displayed; the pixbuf might not
// be the same size as the animation itself
// (gdk_pixbuf_animation_get_width(), gdk_pixbuf_animation_get_height()).
// This pixbuf should be displayed for
// gdk_pixbuf_animation_iter_get_delay_time() milliseconds. The caller
// of this function does not own a reference to the returned pixbuf;
// the returned pixbuf will become invalid when the iterator advances
// to the next frame, which may happen anytime you call
// gdk_pixbuf_animation_iter_advance(). Copy the pixbuf to keep it
// (don't just add a reference), as it may get recycled as you advance
// the iterator.
/*

C function : gdk_pixbuf_animation_iter_get_pixbuf
*/
func (recv *PixbufAnimationIter) GetPixbuf() *Pixbuf {
	retC := C.gdk_pixbuf_animation_iter_get_pixbuf((*C.GdkPixbufAnimationIter)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Used to determine how to respond to the area_updated signal on
// #GdkPixbufLoader when loading an animation. area_updated is emitted
// for an area of the frame currently streaming in to the loader. So if
// you're on the currently loading frame, you need to redraw the screen for
// the updated area.
/*

C function : gdk_pixbuf_animation_iter_on_currently_loading_frame
*/
func (recv *PixbufAnimationIter) OnCurrentlyLoadingFrame() bool {
	retC := C.gdk_pixbuf_animation_iter_on_currently_loading_frame((*C.GdkPixbufAnimationIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// The GdkPixbufLoader struct contains only private
// fields.
/*

C record/class : GdkPixbufLoader
*/
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

	return g
}

func (recv *PixbufLoader) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *PixbufLoader) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to PixbufLoader.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufLoader.
func CastToPixbufLoader(object *gobject.Object) *PixbufLoader {
	return PixbufLoaderNewFromC(object.ToC())
}

type signalPixbufLoaderAreaPreparedDetail struct {
	callback  PixbufLoaderSignalAreaPreparedCallback
	handlerID C.gulong
}

var signalPixbufLoaderAreaPreparedId int
var signalPixbufLoaderAreaPreparedMap = make(map[int]signalPixbufLoaderAreaPreparedDetail)
var signalPixbufLoaderAreaPreparedLock sync.Mutex

// PixbufLoaderSignalAreaPreparedCallback is a callback function for a 'area-prepared' signal emitted from a PixbufLoader.
type PixbufLoaderSignalAreaPreparedCallback func()

/*
ConnectAreaPrepared connects the callback to the 'area-prepared' signal for the PixbufLoader.

The returned value represents the connection, and may be passed to DisconnectAreaPrepared to remove it.
*/
func (recv *PixbufLoader) ConnectAreaPrepared(callback PixbufLoaderSignalAreaPreparedCallback) int {
	signalPixbufLoaderAreaPreparedLock.Lock()
	defer signalPixbufLoaderAreaPreparedLock.Unlock()

	signalPixbufLoaderAreaPreparedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PixbufLoader_signal_connect_area_prepared(instance, C.gpointer(uintptr(signalPixbufLoaderAreaPreparedId)))

	detail := signalPixbufLoaderAreaPreparedDetail{callback, handlerID}
	signalPixbufLoaderAreaPreparedMap[signalPixbufLoaderAreaPreparedId] = detail

	return signalPixbufLoaderAreaPreparedId
}

/*
DisconnectAreaPrepared disconnects a callback from the 'area-prepared' signal for the PixbufLoader.

The connectionID should be a value returned from a call to ConnectAreaPrepared.
*/
func (recv *PixbufLoader) DisconnectAreaPrepared(connectionID int) {
	signalPixbufLoaderAreaPreparedLock.Lock()
	defer signalPixbufLoaderAreaPreparedLock.Unlock()

	detail, exists := signalPixbufLoaderAreaPreparedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPixbufLoaderAreaPreparedMap, connectionID)
}

//export pixbufloader_areaPreparedHandler
func pixbufloader_areaPreparedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalPixbufLoaderAreaPreparedMap[index].callback
	callback()
}

// Unsupported signal 'area-updated' for PixbufLoader : unsupported parameter x : type gint :

type signalPixbufLoaderClosedDetail struct {
	callback  PixbufLoaderSignalClosedCallback
	handlerID C.gulong
}

var signalPixbufLoaderClosedId int
var signalPixbufLoaderClosedMap = make(map[int]signalPixbufLoaderClosedDetail)
var signalPixbufLoaderClosedLock sync.Mutex

// PixbufLoaderSignalClosedCallback is a callback function for a 'closed' signal emitted from a PixbufLoader.
type PixbufLoaderSignalClosedCallback func()

/*
ConnectClosed connects the callback to the 'closed' signal for the PixbufLoader.

The returned value represents the connection, and may be passed to DisconnectClosed to remove it.
*/
func (recv *PixbufLoader) ConnectClosed(callback PixbufLoaderSignalClosedCallback) int {
	signalPixbufLoaderClosedLock.Lock()
	defer signalPixbufLoaderClosedLock.Unlock()

	signalPixbufLoaderClosedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PixbufLoader_signal_connect_closed(instance, C.gpointer(uintptr(signalPixbufLoaderClosedId)))

	detail := signalPixbufLoaderClosedDetail{callback, handlerID}
	signalPixbufLoaderClosedMap[signalPixbufLoaderClosedId] = detail

	return signalPixbufLoaderClosedId
}

/*
DisconnectClosed disconnects a callback from the 'closed' signal for the PixbufLoader.

The connectionID should be a value returned from a call to ConnectClosed.
*/
func (recv *PixbufLoader) DisconnectClosed(connectionID int) {
	signalPixbufLoaderClosedLock.Lock()
	defer signalPixbufLoaderClosedLock.Unlock()

	detail, exists := signalPixbufLoaderClosedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPixbufLoaderClosedMap, connectionID)
}

//export pixbufloader_closedHandler
func pixbufloader_closedHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalPixbufLoaderClosedMap[index].callback
	callback()
}

// Unsupported signal 'size-prepared' for PixbufLoader : unsupported parameter width : type gint :

// Creates a new pixbuf loader object.
/*

C function : gdk_pixbuf_loader_new
*/
func PixbufLoaderNew() *PixbufLoader {
	retC := C.gdk_pixbuf_loader_new()
	retGo := PixbufLoaderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new pixbuf loader object that always attempts to parse
// image data as if it were an image of type @image_type, instead of
// identifying the type automatically. Useful if you want an error if
// the image isn't the expected type, for loading image formats
// that can't be reliably identified by looking at the data, or if
// the user manually forces a specific type.
//
// The list of supported image formats depends on what image loaders
// are installed, but typically "png", "jpeg", "gif", "tiff" and
// "xpm" are among the supported formats. To obtain the full list of
// supported image formats, call gdk_pixbuf_format_get_name() on each
// of the #GdkPixbufFormat structs returned by gdk_pixbuf_get_formats().
/*

C function : gdk_pixbuf_loader_new_with_type
*/
func PixbufLoaderNewWithType(imageType string) (*PixbufLoader, error) {
	c_image_type := C.CString(imageType)
	defer C.free(unsafe.Pointer(c_image_type))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_new_with_type(c_image_type, &cThrowableError)
	retGo := PixbufLoaderNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Informs a pixbuf loader that no further writes with
// gdk_pixbuf_loader_write() will occur, so that it can free its
// internal loading structures. Also, tries to parse any data that
// hasn't yet been parsed; if the remaining data is partial or
// corrupt, an error will be returned.  If %FALSE is returned, @error
// will be set to an error from the #GDK_PIXBUF_ERROR or #G_FILE_ERROR
// domains. If you're just cancelling a load rather than expecting it
// to be finished, passing %NULL for @error to ignore it is
// reasonable.
//
// Remember that this does not unref the loader, so if you plan not to
// use it anymore, please g_object_unref() it.
/*

C function : gdk_pixbuf_loader_close
*/
func (recv *PixbufLoader) Close() (bool, error) {
	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_close((*C.GdkPixbufLoader)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Queries the #GdkPixbufAnimation that a pixbuf loader is currently creating.
// In general it only makes sense to call this function after the "area-prepared"
// signal has been emitted by the loader. If the loader doesn't have enough
// bytes yet (hasn't emitted the "area-prepared" signal) this function will
// return %NULL.
/*

C function : gdk_pixbuf_loader_get_animation
*/
func (recv *PixbufLoader) GetAnimation() *PixbufAnimation {
	retC := C.gdk_pixbuf_loader_get_animation((*C.GdkPixbufLoader)(recv.native))
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Queries the #GdkPixbuf that a pixbuf loader is currently creating.
// In general it only makes sense to call this function after the
// "area-prepared" signal has been emitted by the loader; this means
// that enough data has been read to know the size of the image that
// will be allocated.  If the loader has not received enough data via
// gdk_pixbuf_loader_write(), then this function returns %NULL.  The
// returned pixbuf will be the same in all future calls to the loader,
// so simply calling g_object_ref() should be sufficient to continue
// using it.  Additionally, if the loader is an animation, it will
// return the "static image" of the animation
// (see gdk_pixbuf_animation_get_static_image()).
/*

C function : gdk_pixbuf_loader_get_pixbuf
*/
func (recv *PixbufLoader) GetPixbuf() *Pixbuf {
	retC := C.gdk_pixbuf_loader_get_pixbuf((*C.GdkPixbufLoader)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This will cause a pixbuf loader to parse the next @count bytes of
// an image.  It will return %TRUE if the data was loaded successfully,
// and %FALSE if an error occurred.  In the latter case, the loader
// will be closed, and will not accept further writes. If %FALSE is
// returned, @error will be set to an error from the #GDK_PIXBUF_ERROR
// or #G_FILE_ERROR domains.
/*

C function : gdk_pixbuf_loader_write
*/
func (recv *PixbufLoader) Write(buf []uint8) (bool, error) {
	c_buf := &buf[0]

	c_count := (C.gsize)(len(buf))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_write((*C.GdkPixbufLoader)(recv.native), (*C.guchar)(unsafe.Pointer(c_buf)), c_count, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// An opaque struct representing a simple animation.
/*

C record/class : GdkPixbufSimpleAnim
*/
type PixbufSimpleAnim struct {
	native *C.GdkPixbufSimpleAnim
}

func PixbufSimpleAnimNewFromC(u unsafe.Pointer) *PixbufSimpleAnim {
	c := (*C.GdkPixbufSimpleAnim)(u)
	if c == nil {
		return nil
	}

	g := &PixbufSimpleAnim{native: c}

	return g
}

func (recv *PixbufSimpleAnim) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PixbufAnimation upcasts to *PixbufAnimation
func (recv *PixbufSimpleAnim) PixbufAnimation() *PixbufAnimation {
	return PixbufAnimationNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *PixbufSimpleAnim) Object() *gobject.Object {
	return recv.PixbufAnimation().Object()
}

// CastToWidget down casts any arbitary Object to PixbufSimpleAnim.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufSimpleAnim.
func CastToPixbufSimpleAnim(object *gobject.Object) *PixbufSimpleAnim {
	return PixbufSimpleAnimNewFromC(object.ToC())
}

// Unsupported : PixbufSimpleAnimIter : no CType
