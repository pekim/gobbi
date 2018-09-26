// This is a generated file - DO NOT EDIT

package gdkpixbuf

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
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

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter data : no param type

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

// Unsupported : gdk_pixbuf_new_from_inline : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data : no param type

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

// Unsupported : gdk_pixbuf_composite : no return generator

// Unsupported : gdk_pixbuf_composite_color : no return generator

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

// Unsupported : gdk_pixbuf_copy_area : no return generator

// Unsupported : gdk_pixbuf_fill : no return generator

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

// Unsupported : gdk_pixbuf_get_pixels_with_length : unsupported parameter length : no type generator for guint, guint*

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

// Unsupported : gdk_pixbuf_read_pixels : no return generator

// Ref is a wrapper around the C function gdk_pixbuf_ref.
func (recv *Pixbuf) Ref() *Pixbuf {
	retC := C.gdk_pixbuf_ref((*C.GdkPixbuf)(recv.native))
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_pixbuf_saturate_and_pixelate : no return generator

// Unsupported : gdk_pixbuf_save : unsupported parameter error : record with indirection level of 2

// Unsupported : gdk_pixbuf_save_to_buffer : unsupported parameter buffer : no param type

// Unsupported : gdk_pixbuf_save_to_bufferv : unsupported parameter buffer : no param type

// Unsupported : gdk_pixbuf_save_to_callback : unsupported parameter save_func : no type generator for PixbufSaveFunc, GdkPixbufSaveFunc

// Unsupported : gdk_pixbuf_save_to_callbackv : unsupported parameter save_func : no type generator for PixbufSaveFunc, GdkPixbufSaveFunc

// Unsupported : gdk_pixbuf_save_to_stream : unsupported parameter error : record with indirection level of 2

// Unsupported : gdk_pixbuf_save_to_stream_async : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback, GAsyncReadyCallback

// Unsupported : gdk_pixbuf_save_to_streamv : unsupported parameter option_keys : no param type

// Unsupported : gdk_pixbuf_save_to_streamv_async : unsupported parameter option_keys : no param type

// Unsupported : gdk_pixbuf_savev : unsupported parameter option_keys : no param type

// Unsupported : gdk_pixbuf_scale : no return generator

// ScaleSimple is a wrapper around the C function gdk_pixbuf_scale_simple.
func (recv *Pixbuf) ScaleSimple(destWidth int32, destHeight int32, interpType InterpType) *Pixbuf {
	c_dest_width := (C.int)(destWidth)

	c_dest_height := (C.int)(destHeight)

	c_interp_type := (C.GdkInterpType)(interpType)

	retC := C.gdk_pixbuf_scale_simple((*C.GdkPixbuf)(recv.native), c_dest_width, c_dest_height, c_interp_type)
	retGo := PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gdk_pixbuf_unref : no return generator

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

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// GetHeight is a wrapper around the C function gdk_pixbuf_animation_get_height.
func (recv *PixbufAnimation) GetHeight() int32 {
	retC := C.gdk_pixbuf_animation_get_height((*C.GdkPixbufAnimation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetIter is a wrapper around the C function gdk_pixbuf_animation_get_iter.
func (recv *PixbufAnimation) GetIter(startTime *glib.TimeVal) *PixbufAnimationIter {
	c_start_time := (*C.GTimeVal)(startTime.ToC())

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

// Unsupported : gdk_pixbuf_animation_unref : no return generator

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

// Advance is a wrapper around the C function gdk_pixbuf_animation_iter_advance.
func (recv *PixbufAnimationIter) Advance(currentTime *glib.TimeVal) bool {
	c_current_time := (*C.GTimeVal)(currentTime.ToC())

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

// Unsupported : gdk_pixbuf_loader_set_size : no return generator

// Unsupported : gdk_pixbuf_loader_write : unsupported parameter buf : no param type

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

// Unsupported : gdk_pixbuf_simple_anim_add_frame : no return generator

// Unsupported : gdk_pixbuf_simple_anim_set_loop : no return generator

// Unsupported : PixbufSimpleAnimIter : no CType
