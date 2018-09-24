// This is a generated file - DO NOT EDIT

package gdkpixbuf

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

func (recv *Pixbuf) toC() *C.GdkPixbuf {

	return recv.native
}

// Unsupported : gdk_pixbuf_new : no return generator

// Unsupported : gdk_pixbuf_new_from_bytes : no return generator

// Unsupported : gdk_pixbuf_new_from_data : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_file : no return generator

// Unsupported : gdk_pixbuf_new_from_file_at_scale : no return generator

// Unsupported : gdk_pixbuf_new_from_file_at_size : no return generator

// Unsupported : gdk_pixbuf_new_from_inline : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_new_from_resource : no return generator

// Unsupported : gdk_pixbuf_new_from_resource_at_scale : no return generator

// Unsupported : gdk_pixbuf_new_from_stream : unsupported parameter stream : no type generator for Gio.InputStream, GInputStream*

// Unsupported : gdk_pixbuf_new_from_stream_at_scale : unsupported parameter stream : no type generator for Gio.InputStream, GInputStream*

// Unsupported : gdk_pixbuf_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// Unsupported : gdk_pixbuf_new_from_xpm_data : unsupported parameter data : no param type

// Unsupported : gdk_pixbuf_add_alpha : no return generator

// Unsupported : gdk_pixbuf_apply_embedded_orientation : no return generator

// Unsupported : gdk_pixbuf_composite : unsupported parameter dest : no type generator for Pixbuf, GdkPixbuf*

// Unsupported : gdk_pixbuf_composite_color : unsupported parameter dest : no type generator for Pixbuf, GdkPixbuf*

// Unsupported : gdk_pixbuf_composite_color_simple : no return generator

// Unsupported : gdk_pixbuf_copy : no return generator

// Unsupported : gdk_pixbuf_copy_area : unsupported parameter dest_pixbuf : no type generator for Pixbuf, GdkPixbuf*

// Unsupported : gdk_pixbuf_copy_options : unsupported parameter dest_pixbuf : no type generator for Pixbuf, GdkPixbuf*

// Unsupported : gdk_pixbuf_fill : no return generator

// Unsupported : gdk_pixbuf_flip : no return generator

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

// Unsupported : gdk_pixbuf_new_subpixbuf : no return generator

// Unsupported : gdk_pixbuf_read_pixels : no return generator

// Unsupported : gdk_pixbuf_ref : no return generator

// Unsupported : gdk_pixbuf_rotate_simple : no return generator

// Unsupported : gdk_pixbuf_saturate_and_pixelate : unsupported parameter dest : no type generator for Pixbuf, GdkPixbuf*

// Unsupported : gdk_pixbuf_save : unsupported parameter error : in string with indirection level of 2

// Unsupported : gdk_pixbuf_save_to_buffer : unsupported parameter buffer : no param type

// Unsupported : gdk_pixbuf_save_to_bufferv : unsupported parameter buffer : no param type

// Unsupported : gdk_pixbuf_save_to_callback : unsupported parameter save_func : no type generator for PixbufSaveFunc, GdkPixbufSaveFunc

// Unsupported : gdk_pixbuf_save_to_callbackv : unsupported parameter save_func : no type generator for PixbufSaveFunc, GdkPixbufSaveFunc

// Unsupported : gdk_pixbuf_save_to_stream : unsupported parameter stream : no type generator for Gio.OutputStream, GOutputStream*

// Unsupported : gdk_pixbuf_save_to_stream_async : unsupported parameter stream : no type generator for Gio.OutputStream, GOutputStream*

// Unsupported : gdk_pixbuf_save_to_streamv : unsupported parameter stream : no type generator for Gio.OutputStream, GOutputStream*

// Unsupported : gdk_pixbuf_save_to_streamv_async : unsupported parameter stream : no type generator for Gio.OutputStream, GOutputStream*

// Unsupported : gdk_pixbuf_savev : unsupported parameter option_keys : no param type

// Unsupported : gdk_pixbuf_scale : unsupported parameter dest : no type generator for Pixbuf, GdkPixbuf*

// Unsupported : gdk_pixbuf_scale_simple : no return generator

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

func (recv *PixbufAnimation) toC() *C.GdkPixbufAnimation {

	return recv.native
}

// Unsupported : gdk_pixbuf_animation_new_from_file : no return generator

// Unsupported : gdk_pixbuf_animation_new_from_resource : no return generator

// Unsupported : gdk_pixbuf_animation_new_from_stream : unsupported parameter stream : no type generator for Gio.InputStream, GInputStream*

// Unsupported : gdk_pixbuf_animation_new_from_stream_finish : unsupported parameter async_result : no type generator for Gio.AsyncResult, GAsyncResult*

// GetHeight is a wrapper around the C function gdk_pixbuf_animation_get_height.
func (recv *PixbufAnimation) GetHeight() int32 {
	retC := C.gdk_pixbuf_animation_get_height((*C.GdkPixbufAnimation)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gdk_pixbuf_animation_get_iter : no return generator

// Unsupported : gdk_pixbuf_animation_get_static_image : no return generator

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

// Unsupported : gdk_pixbuf_animation_ref : no return generator

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

func (recv *PixbufAnimationIter) toC() *C.GdkPixbufAnimationIter {

	return recv.native
}

// Advance is a wrapper around the C function gdk_pixbuf_animation_iter_advance.
func (recv *PixbufAnimationIter) Advance(currentTime *glib.TimeVal) bool {
	c_current_time := currentTime.toC()

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

// Unsupported : gdk_pixbuf_animation_iter_get_pixbuf : no return generator

// OnCurrentlyLoadingFrame is a wrapper around the C function gdk_pixbuf_animation_iter_on_currently_loading_frame.
func (recv *PixbufAnimationIter) OnCurrentlyLoadingFrame() bool {
	retC := C.gdk_pixbuf_animation_iter_on_currently_loading_frame((*C.GdkPixbufAnimationIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PixbufLoader is a wrapper around the C record GdkPixbufLoader.
type PixbufLoader struct {
	native *C.GdkPixbufLoader
	// parent_instance : no type generator for GObject.Object, GObject
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

func (recv *PixbufLoader) toC() *C.GdkPixbufLoader {

	return recv.native
}

// Unsupported : gdk_pixbuf_loader_new : no return generator

// Unsupported : gdk_pixbuf_loader_new_with_mime_type : no return generator

// Unsupported : gdk_pixbuf_loader_new_with_type : no return generator

// Close is a wrapper around the C function gdk_pixbuf_loader_close.
func (recv *PixbufLoader) Close() (bool, error) {
	var cThrowableError *C.GError

	retC := C.gdk_pixbuf_loader_close((*C.GdkPixbufLoader)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : gdk_pixbuf_loader_get_animation : no return generator

// Unsupported : gdk_pixbuf_loader_get_pixbuf : no return generator

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

func (recv *PixbufSimpleAnim) toC() *C.GdkPixbufSimpleAnim {

	return recv.native
}

// Unsupported : gdk_pixbuf_simple_anim_new : no return generator

// Unsupported : gdk_pixbuf_simple_anim_add_frame : unsupported parameter pixbuf : no type generator for Pixbuf, GdkPixbuf*

// Unsupported : gdk_pixbuf_simple_anim_set_loop : no return generator

// Unsupported : PixbufSimpleAnimIter : no CType
