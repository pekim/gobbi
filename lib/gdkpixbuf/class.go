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

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// PixbufNewFromInline is a wrapper around the C function gdk_pixbuf_new_from_inline.
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

// Unsupported : gdk_pixbuf_get_pixels : no return type

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

// Unsupported : gdk_pixbuf_save : unsupported parameter ... : varargs

// Unsupported : gdk_pixbuf_savev : unsupported parameter option_keys :

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

// PixbufAnimationNewFromFile is a wrapper around the C function gdk_pixbuf_animation_new_from_file.
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

type signalPixbufLoaderAreaUpdatedDetail struct {
	callback  PixbufLoaderSignalAreaUpdatedCallback
	handlerID C.gulong
}

var signalPixbufLoaderAreaUpdatedId int
var signalPixbufLoaderAreaUpdatedMap = make(map[int]signalPixbufLoaderAreaUpdatedDetail)
var signalPixbufLoaderAreaUpdatedLock sync.Mutex

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

type signalPixbufLoaderSizePreparedDetail struct {
	callback  PixbufLoaderSignalSizePreparedCallback
	handlerID C.gulong
}

var signalPixbufLoaderSizePreparedId int
var signalPixbufLoaderSizePreparedMap = make(map[int]signalPixbufLoaderSizePreparedDetail)
var signalPixbufLoaderSizePreparedLock sync.Mutex

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

	return retGo
}

// PixbufLoaderNewWithType is a wrapper around the C function gdk_pixbuf_loader_new_with_type.
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

// Close is a wrapper around the C function gdk_pixbuf_loader_close.
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

// GetAnimation is a wrapper around the C function gdk_pixbuf_loader_get_animation.
func (recv *PixbufLoader) GetAnimation() *PixbufAnimation {
	retC := C.gdk_pixbuf_loader_get_animation((*C.GdkPixbufLoader)(recv.native))
	retGo := PixbufAnimationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPixbuf is a wrapper around the C function gdk_pixbuf_loader_get_pixbuf.
func (recv *PixbufLoader) GetPixbuf() *Pixbuf {
	retC := C.gdk_pixbuf_loader_get_pixbuf((*C.GdkPixbufLoader)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Write is a wrapper around the C function gdk_pixbuf_loader_write.
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
