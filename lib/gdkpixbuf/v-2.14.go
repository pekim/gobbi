// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.14

package gdkpixbuf

import (
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
/*

	static gboolean _gdk_pixbuf_save(GdkPixbuf* pixbuf, const char* filename, const char* type) {
		return gdk_pixbuf_save(pixbuf, filename, type, NULL, NULL);
    }
*/
/*

	static gboolean _gdk_pixbuf_save_to_stream(GdkPixbuf* pixbuf, GOutputStream* stream, const char* type, GCancellable* cancellable) {
		return gdk_pixbuf_save_to_stream(pixbuf, stream, type, cancellable, NULL, NULL);
    }
*/
/*

	void pixbufloader_areaPreparedHandler(GObject *, gpointer);

	static gulong PixbufLoader_signal_connect_area_prepared(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "area-prepared", G_CALLBACK(pixbufloader_areaPreparedHandler), data);
	}

*/
/*

	void pixbufloader_areaUpdatedHandler(GObject *, gint, gint, gint, gint, gpointer);

	static gulong PixbufLoader_signal_connect_area_updated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "area-updated", G_CALLBACK(pixbufloader_areaUpdatedHandler), data);
	}

*/
/*

	void pixbufloader_closedHandler(GObject *, gpointer);

	static gulong PixbufLoader_signal_connect_closed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closed", G_CALLBACK(pixbufloader_closedHandler), data);
	}

*/
/*

	void pixbufloader_sizePreparedHandler(GObject *, gint, gint, gpointer);

	static gulong PixbufLoader_signal_connect_size_prepared(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "size-prepared", G_CALLBACK(pixbufloader_sizePreparedHandler), data);
	}

*/
import "C"

// Blacklisted : GdkPixdataDumpType

// Blacklisted : GdkPixdataType

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

// PixbufNew is a wrapper around the C function gdk_pixbuf_new.
func PixbufNew(colorspace Colorspace, hasAlpha bool, bitsPerSample int32, width int32, height int32) *Pixbuf {
	c_colorspace := (C.GdkColorspace)(colorspace)

	c_has_alpha :=
		boolToGboolean(hasAlpha)

	c_bits_per_sample := (C.int)(bitsPerSample)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	retC := C.gdk_pixbuf_new(c_colorspace, c_has_alpha, c_bits_per_sample, c_width, c_height)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

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

// PixbufNewFromFileAtScale is a wrapper around the C function gdk_pixbuf_new_from_file_at_scale.
func PixbufNewFromFileAtScale(filename string, width int32, height int32, preserveAspectRatio bool) (*Pixbuf, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_preserve_aspect_ratio :=
		boolToGboolean(preserveAspectRatio)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_file_at_scale(c_filename, c_width, c_height, c_preserve_aspect_ratio, &cThrowableError)
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

// PixbufNewFromFileAtSize is a wrapper around the C function gdk_pixbuf_new_from_file_at_size.
func PixbufNewFromFileAtSize(filename string, width int32, height int32) (*Pixbuf, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_file_at_size(c_filename, c_width, c_height, &cThrowableError)
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

// PixbufNewFromInline is a wrapper around the C function gdk_pixbuf_new_from_inline.
func PixbufNewFromInline(data []uint8, copyPixels bool) (*Pixbuf, error) {
	c_data_length := (C.gint)(len(data))

	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (*C.guint8)(unsafe.Pointer(c_data_arrayPtr))

	c_copy_pixels :=
		boolToGboolean(copyPixels)

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_inline(c_data_length, c_data, c_copy_pixels, &cThrowableError)
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

// PixbufNewFromStream is a wrapper around the C function gdk_pixbuf_new_from_stream.
func PixbufNewFromStream(stream *gio.InputStream, cancellable *gio.Cancellable) (*Pixbuf, error) {
	c_stream := (*C.GInputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GInputStream)(stream.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_stream(c_stream, c_cancellable, &cThrowableError)
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

// PixbufNewFromStreamAtScale is a wrapper around the C function gdk_pixbuf_new_from_stream_at_scale.
func PixbufNewFromStreamAtScale(stream *gio.InputStream, width int32, height int32, preserveAspectRatio bool, cancellable *gio.Cancellable) (*Pixbuf, error) {
	c_stream := (*C.GInputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GInputStream)(stream.ToC())
	}

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_preserve_aspect_ratio :=
		boolToGboolean(preserveAspectRatio)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_new_from_stream_at_scale(c_stream, c_width, c_height, c_preserve_aspect_ratio, c_cancellable, &cThrowableError)
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

// PixbufNewFromXpmData is a wrapper around the C function gdk_pixbuf_new_from_xpm_data.
func PixbufNewFromXpmData(data []string) *Pixbuf {
	c_data_array := make([]*C.char, len(data)+1, len(data)+1)
	for i, item := range data {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_data_array[i] = c
	}
	c_data_array[len(data)] = nil
	c_data_arrayPtr := &c_data_array[0]
	c_data := (**C.char)(unsafe.Pointer(c_data_arrayPtr))

	retC := C.gdk_pixbuf_new_from_xpm_data(c_data)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// gdk_pixbuf_from_pixdata : unsupported parameter pixdata : Blacklisted record : GdkPixdata
// PixbufGetFileInfo is a wrapper around the C function gdk_pixbuf_get_file_info.
func PixbufGetFileInfo(filename string) (*PixbufFormat, int32, int32) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var c_width C.gint

	var c_height C.gint

	retC := C.gdk_pixbuf_get_file_info(c_filename, &c_width, &c_height)
	var retGo (*PixbufFormat)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PixbufFormatNewFromC(unsafe.Pointer(retC))
	}

	width := (int32)(c_width)

	height := (int32)(c_height)

	return retGo, width, height
}

// PixbufGetFormats is a wrapper around the C function gdk_pixbuf_get_formats.
func PixbufGetFormats() *glib.SList {
	retC := C.gdk_pixbuf_get_formats()
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddAlpha is a wrapper around the C function gdk_pixbuf_add_alpha.
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

// ApplyEmbeddedOrientation is a wrapper around the C function gdk_pixbuf_apply_embedded_orientation.
func (recv *Pixbuf) ApplyEmbeddedOrientation() *Pixbuf {
	retC := C.gdk_pixbuf_apply_embedded_orientation((*C.GdkPixbuf)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Composite is a wrapper around the C function gdk_pixbuf_composite.
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

// CompositeColor is a wrapper around the C function gdk_pixbuf_composite_color.
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

// CompositeColorSimple is a wrapper around the C function gdk_pixbuf_composite_color_simple.
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

// Copy is a wrapper around the C function gdk_pixbuf_copy.
func (recv *Pixbuf) Copy() *Pixbuf {
	retC := C.gdk_pixbuf_copy((*C.GdkPixbuf)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CopyArea is a wrapper around the C function gdk_pixbuf_copy_area.
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

// Fill is a wrapper around the C function gdk_pixbuf_fill.
func (recv *Pixbuf) Fill(pixel uint32) {
	c_pixel := (C.guint32)(pixel)

	C.gdk_pixbuf_fill((*C.GdkPixbuf)(recv.native), c_pixel)

	return
}

// Flip is a wrapper around the C function gdk_pixbuf_flip.
func (recv *Pixbuf) Flip(horizontal bool) *Pixbuf {
	c_horizontal :=
		boolToGboolean(horizontal)

	retC := C.gdk_pixbuf_flip((*C.GdkPixbuf)(recv.native), c_horizontal)
	var retGo (*Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetBitsPerSample is a wrapper around the C function gdk_pixbuf_get_bits_per_sample.
func (recv *Pixbuf) GetBitsPerSample() int32 {
	retC := C.gdk_pixbuf_get_bits_per_sample((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetColorspace is a wrapper around the C function gdk_pixbuf_get_colorspace.
func (recv *Pixbuf) GetColorspace() Colorspace {
	retC := C.gdk_pixbuf_get_colorspace((*C.GdkPixbuf)(recv.native))
	retGo := (Colorspace)(retC)

	return retGo
}

// GetHasAlpha is a wrapper around the C function gdk_pixbuf_get_has_alpha.
func (recv *Pixbuf) GetHasAlpha() bool {
	retC := C.gdk_pixbuf_get_has_alpha((*C.GdkPixbuf)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHeight is a wrapper around the C function gdk_pixbuf_get_height.
func (recv *Pixbuf) GetHeight() int32 {
	retC := C.gdk_pixbuf_get_height((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNChannels is a wrapper around the C function gdk_pixbuf_get_n_channels.
func (recv *Pixbuf) GetNChannels() int32 {
	retC := C.gdk_pixbuf_get_n_channels((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetOption is a wrapper around the C function gdk_pixbuf_get_option.
func (recv *Pixbuf) GetOption(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.gdk_pixbuf_get_option((*C.GdkPixbuf)(recv.native), c_key)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gdk_pixbuf_get_pixels : array return type :

// GetRowstride is a wrapper around the C function gdk_pixbuf_get_rowstride.
func (recv *Pixbuf) GetRowstride() int32 {
	retC := C.gdk_pixbuf_get_rowstride((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWidth is a wrapper around the C function gdk_pixbuf_get_width.
func (recv *Pixbuf) GetWidth() int32 {
	retC := C.gdk_pixbuf_get_width((*C.GdkPixbuf)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// NewSubpixbuf is a wrapper around the C function gdk_pixbuf_new_subpixbuf.
func (recv *Pixbuf) NewSubpixbuf(srcX int32, srcY int32, width int32, height int32) *Pixbuf {
	c_src_x := (C.int)(srcX)

	c_src_y := (C.int)(srcY)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	retC := C.gdk_pixbuf_new_subpixbuf((*C.GdkPixbuf)(recv.native), c_src_x, c_src_y, c_width, c_height)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function gdk_pixbuf_ref.
func (recv *Pixbuf) Ref() *Pixbuf {
	retC := C.gdk_pixbuf_ref((*C.GdkPixbuf)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RotateSimple is a wrapper around the C function gdk_pixbuf_rotate_simple.
func (recv *Pixbuf) RotateSimple(angle PixbufRotation) *Pixbuf {
	c_angle := (C.GdkPixbufRotation)(angle)

	retC := C.gdk_pixbuf_rotate_simple((*C.GdkPixbuf)(recv.native), c_angle)
	var retGo (*Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SaturateAndPixelate is a wrapper around the C function gdk_pixbuf_saturate_and_pixelate.
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

// Save is a wrapper around the C function gdk_pixbuf_save.
func (recv *Pixbuf) Save(filename string, type_ string, error_ *glib.Error) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_error := (**C.GError)(C.NULL)
	if error_ != nil {
		c_error = (**C.GError)(error_.ToC())
	}

	retC := C._gdk_pixbuf_save((*C.GdkPixbuf)(recv.native), c_filename, c_type, c_error)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gdk_pixbuf_save_to_buffer : unsupported parameter buffer : output array param buffer

// Unsupported : gdk_pixbuf_save_to_bufferv : unsupported parameter buffer : output array param buffer

// Unsupported : gdk_pixbuf_save_to_callback : unsupported parameter save_func : no type generator for PixbufSaveFunc (GdkPixbufSaveFunc) for param save_func

// Unsupported : gdk_pixbuf_save_to_callbackv : unsupported parameter save_func : no type generator for PixbufSaveFunc (GdkPixbufSaveFunc) for param save_func

// SaveToStream is a wrapper around the C function gdk_pixbuf_save_to_stream.
func (recv *Pixbuf) SaveToStream(stream *gio.OutputStream, type_ string, cancellable *gio.Cancellable, error_ *glib.Error) bool {
	c_stream := (*C.GOutputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GOutputStream)(stream.ToC())
	}

	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	c_error := (**C.GError)(C.NULL)
	if error_ != nil {
		c_error = (**C.GError)(error_.ToC())
	}

	retC := C._gdk_pixbuf_save_to_stream((*C.GdkPixbuf)(recv.native), c_stream, c_type, c_cancellable, c_error)
	retGo := retC == C.TRUE

	return retGo
}

// Savev is a wrapper around the C function gdk_pixbuf_savev.
func (recv *Pixbuf) Savev(filename string, type_ string, optionKeys []string, optionValues []string) (bool, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	c_option_keys_array := make([]*C.char, len(optionKeys)+1, len(optionKeys)+1)
	for i, item := range optionKeys {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_option_keys_array[i] = c
	}
	c_option_keys_array[len(optionKeys)] = nil
	c_option_keys_arrayPtr := &c_option_keys_array[0]
	c_option_keys := (**C.char)(unsafe.Pointer(c_option_keys_arrayPtr))

	c_option_values_array := make([]*C.char, len(optionValues)+1, len(optionValues)+1)
	for i, item := range optionValues {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_option_values_array[i] = c
	}
	c_option_values_array[len(optionValues)] = nil
	c_option_values_arrayPtr := &c_option_values_array[0]
	c_option_values := (**C.char)(unsafe.Pointer(c_option_values_arrayPtr))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_savev((*C.GdkPixbuf)(recv.native), c_filename, c_type, c_option_keys, c_option_values, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Scale is a wrapper around the C function gdk_pixbuf_scale.
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

// ScaleSimple is a wrapper around the C function gdk_pixbuf_scale_simple.
func (recv *Pixbuf) ScaleSimple(destWidth int32, destHeight int32, interpType InterpType) *Pixbuf {
	c_dest_width := (C.int)(destWidth)

	c_dest_height := (C.int)(destHeight)

	c_interp_type := (C.GdkInterpType)(interpType)

	retC := C.gdk_pixbuf_scale_simple((*C.GdkPixbuf)(recv.native), c_dest_width, c_dest_height, c_interp_type)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : gdk_pixbuf_set_option

// Unref is a wrapper around the C function gdk_pixbuf_unref.
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

// PixbufAnimation is a wrapper around the C record GdkPixbufAnimation.
type PixbufAnimation struct {
	native *C.GdkPixbufAnimation
}

func PixbufAnimationNewFromC(u unsafe.Pointer) *PixbufAnimation {
	c := (*C.GdkPixbufAnimation)(u)
	if c == nil {
		return nil
	}

	g := &PixbufAnimation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufAnimation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufAnimation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufAnimation with another PixbufAnimation, and returns true if they represent the same GObject.
func (recv *PixbufAnimation) Equals(other *PixbufAnimation) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PixbufAnimation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PixbufAnimation.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufAnimation.
func CastToPixbufAnimation(object *gobject.Object) *PixbufAnimation {
	return PixbufAnimationNewFromC(object.ToC())
}

// PixbufAnimationNewFromFile is a wrapper around the C function gdk_pixbuf_animation_new_from_file.
func PixbufAnimationNewFromFile(filename string) (*PixbufAnimation, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_animation_new_from_file(c_filename, &cThrowableError)
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

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

// GetHeight is a wrapper around the C function gdk_pixbuf_animation_get_height.
func (recv *PixbufAnimation) GetHeight() int32 {
	retC := C.gdk_pixbuf_animation_get_height((*C.GdkPixbufAnimation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetIter is a wrapper around the C function gdk_pixbuf_animation_get_iter.
func (recv *PixbufAnimation) GetIter(startTime *glib.TimeVal) *PixbufAnimationIter {
	c_start_time := (*C.GTimeVal)(C.NULL)
	if startTime != nil {
		c_start_time = (*C.GTimeVal)(startTime.ToC())
	}

	retC := C.gdk_pixbuf_animation_get_iter((*C.GdkPixbufAnimation)(recv.native), c_start_time)
	retGo := PixbufAnimationIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStaticImage is a wrapper around the C function gdk_pixbuf_animation_get_static_image.
func (recv *PixbufAnimation) GetStaticImage() *Pixbuf {
	retC := C.gdk_pixbuf_animation_get_static_image((*C.GdkPixbufAnimation)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidth is a wrapper around the C function gdk_pixbuf_animation_get_width.
func (recv *PixbufAnimation) GetWidth() int32 {
	retC := C.gdk_pixbuf_animation_get_width((*C.GdkPixbufAnimation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// IsStaticImage is a wrapper around the C function gdk_pixbuf_animation_is_static_image.
func (recv *PixbufAnimation) IsStaticImage() bool {
	retC := C.gdk_pixbuf_animation_is_static_image((*C.GdkPixbufAnimation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function gdk_pixbuf_animation_ref.
func (recv *PixbufAnimation) Ref() *PixbufAnimation {
	retC := C.gdk_pixbuf_animation_ref((*C.GdkPixbufAnimation)(recv.native))
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function gdk_pixbuf_animation_unref.
func (recv *PixbufAnimation) Unref() {
	C.gdk_pixbuf_animation_unref((*C.GdkPixbufAnimation)(recv.native))

	return
}

// PixbufAnimationIter is a wrapper around the C record GdkPixbufAnimationIter.
type PixbufAnimationIter struct {
	native *C.GdkPixbufAnimationIter
}

func PixbufAnimationIterNewFromC(u unsafe.Pointer) *PixbufAnimationIter {
	c := (*C.GdkPixbufAnimationIter)(u)
	if c == nil {
		return nil
	}

	g := &PixbufAnimationIter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufAnimationIter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufAnimationIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufAnimationIter with another PixbufAnimationIter, and returns true if they represent the same GObject.
func (recv *PixbufAnimationIter) Equals(other *PixbufAnimationIter) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PixbufAnimationIter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PixbufAnimationIter.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufAnimationIter.
func CastToPixbufAnimationIter(object *gobject.Object) *PixbufAnimationIter {
	return PixbufAnimationIterNewFromC(object.ToC())
}

// Advance is a wrapper around the C function gdk_pixbuf_animation_iter_advance.
func (recv *PixbufAnimationIter) Advance(currentTime *glib.TimeVal) bool {
	c_current_time := (*C.GTimeVal)(C.NULL)
	if currentTime != nil {
		c_current_time = (*C.GTimeVal)(currentTime.ToC())
	}

	retC := C.gdk_pixbuf_animation_iter_advance((*C.GdkPixbufAnimationIter)(recv.native), c_current_time)
	retGo := retC == C.TRUE

	return retGo
}

// GetDelayTime is a wrapper around the C function gdk_pixbuf_animation_iter_get_delay_time.
func (recv *PixbufAnimationIter) GetDelayTime() int32 {
	retC := C.gdk_pixbuf_animation_iter_get_delay_time((*C.GdkPixbufAnimationIter)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPixbuf is a wrapper around the C function gdk_pixbuf_animation_iter_get_pixbuf.
func (recv *PixbufAnimationIter) GetPixbuf() *Pixbuf {
	retC := C.gdk_pixbuf_animation_iter_get_pixbuf((*C.GdkPixbufAnimationIter)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// OnCurrentlyLoadingFrame is a wrapper around the C function gdk_pixbuf_animation_iter_on_currently_loading_frame.
func (recv *PixbufAnimationIter) OnCurrentlyLoadingFrame() bool {
	retC := C.gdk_pixbuf_animation_iter_on_currently_loading_frame((*C.GdkPixbufAnimationIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PixbufLoader is a wrapper around the C record GdkPixbufLoader.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufLoader) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufLoader) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufLoader with another PixbufLoader, and returns true if they represent the same GObject.
func (recv *PixbufLoader) Equals(other *PixbufLoader) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PixbufLoader) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PixbufLoader.
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
var signalPixbufLoaderAreaPreparedLock sync.RWMutex

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
	signalPixbufLoaderAreaPreparedLock.RLock()
	defer signalPixbufLoaderAreaPreparedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalPixbufLoaderAreaPreparedMap[index].callback
	callback()
}

type signalPixbufLoaderAreaUpdatedDetail struct {
	callback  PixbufLoaderSignalAreaUpdatedCallback
	handlerID C.gulong
}

var signalPixbufLoaderAreaUpdatedId int
var signalPixbufLoaderAreaUpdatedMap = make(map[int]signalPixbufLoaderAreaUpdatedDetail)
var signalPixbufLoaderAreaUpdatedLock sync.RWMutex

// PixbufLoaderSignalAreaUpdatedCallback is a callback function for a 'area-updated' signal emitted from a PixbufLoader.
type PixbufLoaderSignalAreaUpdatedCallback func(x int32, y int32, width int32, height int32)

/*
ConnectAreaUpdated connects the callback to the 'area-updated' signal for the PixbufLoader.

The returned value represents the connection, and may be passed to DisconnectAreaUpdated to remove it.
*/
func (recv *PixbufLoader) ConnectAreaUpdated(callback PixbufLoaderSignalAreaUpdatedCallback) int {
	signalPixbufLoaderAreaUpdatedLock.Lock()
	defer signalPixbufLoaderAreaUpdatedLock.Unlock()

	signalPixbufLoaderAreaUpdatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PixbufLoader_signal_connect_area_updated(instance, C.gpointer(uintptr(signalPixbufLoaderAreaUpdatedId)))

	detail := signalPixbufLoaderAreaUpdatedDetail{callback, handlerID}
	signalPixbufLoaderAreaUpdatedMap[signalPixbufLoaderAreaUpdatedId] = detail

	return signalPixbufLoaderAreaUpdatedId
}

/*
DisconnectAreaUpdated disconnects a callback from the 'area-updated' signal for the PixbufLoader.

The connectionID should be a value returned from a call to ConnectAreaUpdated.
*/
func (recv *PixbufLoader) DisconnectAreaUpdated(connectionID int) {
	signalPixbufLoaderAreaUpdatedLock.Lock()
	defer signalPixbufLoaderAreaUpdatedLock.Unlock()

	detail, exists := signalPixbufLoaderAreaUpdatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPixbufLoaderAreaUpdatedMap, connectionID)
}

//export pixbufloader_areaUpdatedHandler
func pixbufloader_areaUpdatedHandler(_ *C.GObject, c_x C.gint, c_y C.gint, c_width C.gint, c_height C.gint, data C.gpointer) {
	signalPixbufLoaderAreaUpdatedLock.RLock()
	defer signalPixbufLoaderAreaUpdatedLock.RUnlock()

	x := int32(c_x)

	y := int32(c_y)

	width := int32(c_width)

	height := int32(c_height)

	index := int(uintptr(data))
	callback := signalPixbufLoaderAreaUpdatedMap[index].callback
	callback(x, y, width, height)
}

type signalPixbufLoaderClosedDetail struct {
	callback  PixbufLoaderSignalClosedCallback
	handlerID C.gulong
}

var signalPixbufLoaderClosedId int
var signalPixbufLoaderClosedMap = make(map[int]signalPixbufLoaderClosedDetail)
var signalPixbufLoaderClosedLock sync.RWMutex

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
	signalPixbufLoaderClosedLock.RLock()
	defer signalPixbufLoaderClosedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalPixbufLoaderClosedMap[index].callback
	callback()
}

type signalPixbufLoaderSizePreparedDetail struct {
	callback  PixbufLoaderSignalSizePreparedCallback
	handlerID C.gulong
}

var signalPixbufLoaderSizePreparedId int
var signalPixbufLoaderSizePreparedMap = make(map[int]signalPixbufLoaderSizePreparedDetail)
var signalPixbufLoaderSizePreparedLock sync.RWMutex

// PixbufLoaderSignalSizePreparedCallback is a callback function for a 'size-prepared' signal emitted from a PixbufLoader.
type PixbufLoaderSignalSizePreparedCallback func(width int32, height int32)

/*
ConnectSizePrepared connects the callback to the 'size-prepared' signal for the PixbufLoader.

The returned value represents the connection, and may be passed to DisconnectSizePrepared to remove it.
*/
func (recv *PixbufLoader) ConnectSizePrepared(callback PixbufLoaderSignalSizePreparedCallback) int {
	signalPixbufLoaderSizePreparedLock.Lock()
	defer signalPixbufLoaderSizePreparedLock.Unlock()

	signalPixbufLoaderSizePreparedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PixbufLoader_signal_connect_size_prepared(instance, C.gpointer(uintptr(signalPixbufLoaderSizePreparedId)))

	detail := signalPixbufLoaderSizePreparedDetail{callback, handlerID}
	signalPixbufLoaderSizePreparedMap[signalPixbufLoaderSizePreparedId] = detail

	return signalPixbufLoaderSizePreparedId
}

/*
DisconnectSizePrepared disconnects a callback from the 'size-prepared' signal for the PixbufLoader.

The connectionID should be a value returned from a call to ConnectSizePrepared.
*/
func (recv *PixbufLoader) DisconnectSizePrepared(connectionID int) {
	signalPixbufLoaderSizePreparedLock.Lock()
	defer signalPixbufLoaderSizePreparedLock.Unlock()

	detail, exists := signalPixbufLoaderSizePreparedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPixbufLoaderSizePreparedMap, connectionID)
}

//export pixbufloader_sizePreparedHandler
func pixbufloader_sizePreparedHandler(_ *C.GObject, c_width C.gint, c_height C.gint, data C.gpointer) {
	signalPixbufLoaderSizePreparedLock.RLock()
	defer signalPixbufLoaderSizePreparedLock.RUnlock()

	width := int32(c_width)

	height := int32(c_height)

	index := int(uintptr(data))
	callback := signalPixbufLoaderSizePreparedMap[index].callback
	callback(width, height)
}

// PixbufLoaderNew is a wrapper around the C function gdk_pixbuf_loader_new.
func PixbufLoaderNew() *PixbufLoader {
	retC := C.gdk_pixbuf_loader_new()
	retGo := PixbufLoaderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// PixbufLoaderNewWithMimeType is a wrapper around the C function gdk_pixbuf_loader_new_with_mime_type.
func PixbufLoaderNewWithMimeType(mimeType string) (*PixbufLoader, error) {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_new_with_mime_type(c_mime_type, &cThrowableError)
	retGo := PixbufLoaderNewFromC(unsafe.Pointer(retC))

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

// PixbufLoaderNewWithType is a wrapper around the C function gdk_pixbuf_loader_new_with_type.
func PixbufLoaderNewWithType(imageType string) (*PixbufLoader, error) {
	c_image_type := C.CString(imageType)
	defer C.free(unsafe.Pointer(c_image_type))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_new_with_type(c_image_type, &cThrowableError)
	retGo := PixbufLoaderNewFromC(unsafe.Pointer(retC))

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

// Close is a wrapper around the C function gdk_pixbuf_loader_close.
func (recv *PixbufLoader) Close() (bool, error) {
	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_close((*C.GdkPixbufLoader)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAnimation is a wrapper around the C function gdk_pixbuf_loader_get_animation.
func (recv *PixbufLoader) GetAnimation() *PixbufAnimation {
	retC := C.gdk_pixbuf_loader_get_animation((*C.GdkPixbufLoader)(recv.native))
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFormat is a wrapper around the C function gdk_pixbuf_loader_get_format.
func (recv *PixbufLoader) GetFormat() *PixbufFormat {
	retC := C.gdk_pixbuf_loader_get_format((*C.GdkPixbufLoader)(recv.native))
	var retGo (*PixbufFormat)
	if retC == nil {
		retGo = nil
	} else {
		retGo = PixbufFormatNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetPixbuf is a wrapper around the C function gdk_pixbuf_loader_get_pixbuf.
func (recv *PixbufLoader) GetPixbuf() *Pixbuf {
	retC := C.gdk_pixbuf_loader_get_pixbuf((*C.GdkPixbufLoader)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetSize is a wrapper around the C function gdk_pixbuf_loader_set_size.
func (recv *PixbufLoader) SetSize(width int32, height int32) {
	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.gdk_pixbuf_loader_set_size((*C.GdkPixbufLoader)(recv.native), c_width, c_height)

	return
}

// Write is a wrapper around the C function gdk_pixbuf_loader_write.
func (recv *PixbufLoader) Write(buf []uint8) (bool, error) {
	c_buf_array := make([]C.guchar, len(buf)+1, len(buf)+1)
	for i, item := range buf {
		c := (C.guchar)(item)
		c_buf_array[i] = c
	}
	c_buf_array[len(buf)] = 0
	c_buf_arrayPtr := &c_buf_array[0]
	c_buf := (*C.guchar)(unsafe.Pointer(c_buf_arrayPtr))

	c_count := (C.gsize)(len(buf))

	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_write((*C.GdkPixbufLoader)(recv.native), c_buf, c_count, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PixbufSimpleAnim is a wrapper around the C record GdkPixbufSimpleAnim.
type PixbufSimpleAnim struct {
	native *C.GdkPixbufSimpleAnim
}

func PixbufSimpleAnimNewFromC(u unsafe.Pointer) *PixbufSimpleAnim {
	c := (*C.GdkPixbufSimpleAnim)(u)
	if c == nil {
		return nil
	}

	g := &PixbufSimpleAnim{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PixbufSimpleAnim) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PixbufSimpleAnim) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufSimpleAnim with another PixbufSimpleAnim, and returns true if they represent the same GObject.
func (recv *PixbufSimpleAnim) Equals(other *PixbufSimpleAnim) bool {
	return other.ToC() == recv.ToC()
}

// PixbufAnimation upcasts to *PixbufAnimation
func (recv *PixbufSimpleAnim) PixbufAnimation() *PixbufAnimation {
	return PixbufAnimationNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *PixbufSimpleAnim) Object() *gobject.Object {
	return recv.PixbufAnimation().Object()
}

// CastToWidget down casts any arbitrary Object to PixbufSimpleAnim.
// Exercise care, as this is a potentially dangerous function if the Object is not a PixbufSimpleAnim.
func CastToPixbufSimpleAnim(object *gobject.Object) *PixbufSimpleAnim {
	return PixbufSimpleAnimNewFromC(object.ToC())
}

// PixbufSimpleAnimNew is a wrapper around the C function gdk_pixbuf_simple_anim_new.
func PixbufSimpleAnimNew(width int32, height int32, rate float32) *PixbufSimpleAnim {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_rate := (C.gfloat)(rate)

	retC := C.gdk_pixbuf_simple_anim_new(c_width, c_height, c_rate)
	retGo := PixbufSimpleAnimNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddFrame is a wrapper around the C function gdk_pixbuf_simple_anim_add_frame.
func (recv *PixbufSimpleAnim) AddFrame(pixbuf *Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gdk_pixbuf_simple_anim_add_frame((*C.GdkPixbufSimpleAnim)(recv.native), c_pixbuf)

	return
}

// Unsupported : PixbufSimpleAnimIter : no CType

const PIXBUF_FEATURES_H int32 = C.GDK_PIXBUF_FEATURES_H

// Blacklisted : PIXBUF_MAGIC_NUMBER

const PIXBUF_MAJOR int32 = C.GDK_PIXBUF_MAJOR
const PIXBUF_MICRO int32 = C.GDK_PIXBUF_MICRO
const PIXBUF_MINOR int32 = C.GDK_PIXBUF_MINOR
const PIXBUF_VERSION string = C.GDK_PIXBUF_VERSION

// Blacklisted : PIXDATA_HEADER_LENGTH

type Colorspace C.GdkColorspace

const (
	GDK_COLORSPACE_RGB Colorspace = 0
)

type InterpType C.GdkInterpType

const (
	GDK_INTERP_NEAREST  InterpType = 0
	GDK_INTERP_TILES    InterpType = 1
	GDK_INTERP_BILINEAR InterpType = 2
	GDK_INTERP_HYPER    InterpType = 3
)

type PixbufAlphaMode C.GdkPixbufAlphaMode

const (
	GDK_PIXBUF_ALPHA_BILEVEL PixbufAlphaMode = 0
	GDK_PIXBUF_ALPHA_FULL    PixbufAlphaMode = 1
)

type PixbufError C.GdkPixbufError

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

type PixbufRotation C.GdkPixbufRotation

const (
	GDK_PIXBUF_ROTATE_NONE             PixbufRotation = 0
	GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = 90
	GDK_PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = 180
	GDK_PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = 270
)

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

func (recv *PixbufFormat) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufFormat with another PixbufFormat, and returns true if they represent the same GObject.
func (recv *PixbufFormat) Equals(other *PixbufFormat) bool {
	return other.ToC() == recv.ToC()
}

// GetDescription is a wrapper around the C function gdk_pixbuf_format_get_description.
func (recv *PixbufFormat) GetDescription() string {
	retC := C.gdk_pixbuf_format_get_description((*C.GdkPixbufFormat)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetExtensions is a wrapper around the C function gdk_pixbuf_format_get_extensions.
func (recv *PixbufFormat) GetExtensions() []string {
	retC := C.gdk_pixbuf_format_get_extensions((*C.GdkPixbufFormat)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetLicense is a wrapper around the C function gdk_pixbuf_format_get_license.
func (recv *PixbufFormat) GetLicense() string {
	retC := C.gdk_pixbuf_format_get_license((*C.GdkPixbufFormat)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetMimeTypes is a wrapper around the C function gdk_pixbuf_format_get_mime_types.
func (recv *PixbufFormat) GetMimeTypes() []string {
	retC := C.gdk_pixbuf_format_get_mime_types((*C.GdkPixbufFormat)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetName is a wrapper around the C function gdk_pixbuf_format_get_name.
func (recv *PixbufFormat) GetName() string {
	retC := C.gdk_pixbuf_format_get_name((*C.GdkPixbufFormat)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// IsDisabled is a wrapper around the C function gdk_pixbuf_format_is_disabled.
func (recv *PixbufFormat) IsDisabled() bool {
	retC := C.gdk_pixbuf_format_is_disabled((*C.GdkPixbufFormat)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsScalable is a wrapper around the C function gdk_pixbuf_format_is_scalable.
func (recv *PixbufFormat) IsScalable() bool {
	retC := C.gdk_pixbuf_format_is_scalable((*C.GdkPixbufFormat)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsWritable is a wrapper around the C function gdk_pixbuf_format_is_writable.
func (recv *PixbufFormat) IsWritable() bool {
	retC := C.gdk_pixbuf_format_is_writable((*C.GdkPixbufFormat)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDisabled is a wrapper around the C function gdk_pixbuf_format_set_disabled.
func (recv *PixbufFormat) SetDisabled(disabled bool) {
	c_disabled :=
		boolToGboolean(disabled)

	C.gdk_pixbuf_format_set_disabled((*C.GdkPixbufFormat)(recv.native), c_disabled)

	return
}

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

func (recv *PixbufLoaderClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufLoaderClass with another PixbufLoaderClass, and returns true if they represent the same GObject.
func (recv *PixbufLoaderClass) Equals(other *PixbufLoaderClass) bool {
	return other.ToC() == recv.ToC()
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

func (recv *PixbufSimpleAnimClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PixbufSimpleAnimClass with another PixbufSimpleAnimClass, and returns true if they represent the same GObject.
func (recv *PixbufSimpleAnimClass) Equals(other *PixbufSimpleAnimClass) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GdkPixdata
