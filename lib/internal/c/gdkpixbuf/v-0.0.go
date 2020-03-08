// Code generated - DO NOT EDIT.
// +build !gdkpixbuf_2.2,!gdkpixbuf_2.4,!gdkpixbuf_2.6,!gdkpixbuf_2.8,!gdkpixbuf_2.12,!gdkpixbuf_2.14,!gdkpixbuf_2.18,!gdkpixbuf_2.22,!gdkpixbuf_2.24,!gdkpixbuf_2.26,!gdkpixbuf_2.28,!gdkpixbuf_2.30,!gdkpixbuf_2.32,!gdkpixbuf_2.36,!gdkpixbuf_2.36.8,!gdkpixbuf_2.40

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

type PixbufFormat C.GdkPixbufFormat
type PixbufLoaderClass C.GdkPixbufLoaderClass
type PixbufSimpleAnimClass C.GdkPixbufSimpleAnimClass

// UNSUPPORTED : gdk_pixbuf_format_get_extensions : no array length

// UNSUPPORTED : gdk_pixbuf_format_get_mime_types : no array length

func Fn_gdk_pixbuf_new(colorspace int, hasAlpha bool, bitsPerSample int, width int, height int) unsafe.Pointer {
	c_colorspace := (C.GdkColorspace)(colorspace)

	c_hasAlpha := toCBool(hasAlpha)

	c_bitsPerSample := (C.int)(bitsPerSample)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	ret := C.gdk_pixbuf_new(c_colorspace, c_hasAlpha, c_bitsPerSample, c_width, c_height)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_data : parameter 'destroy_fn' is callback

func Fn_gdk_pixbuf_new_from_file(filename string, error unsafe.Pointer) unsafe.Pointer {
	c_filename := (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_file(c_filename, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_inline(dataLength int, data []uint8, copyPixels bool, error unsafe.Pointer) unsafe.Pointer {
	c_dataLength := (C.gint)(dataLength)

	c_data := (*C.guint8)(unsafe.Pointer(&data[0]))

	c_copyPixels := toCBool(copyPixels)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_inline(c_dataLength, c_data, c_copyPixels, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_xpm_data(data []string) unsafe.Pointer {
	dataLen := len(data)
	c_dataArray := C.malloc((C.ulong)(dataLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_dataArray))
	dataSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_dataArray))[:dataLen:dataLen]
	for datai, dataString := range data {
		dataSlice[datai] = (*C.gchar)(C.CString(dataString))
		defer C.free(unsafe.Pointer(dataSlice[datai]))
	}
	c_data := &dataSlice[0]

	ret := C.gdk_pixbuf_new_from_xpm_data(c_data)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_add_alpha(pixbuf unsafe.Pointer, substituteColor bool, r uint8, g uint8, b uint8) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_substituteColor := toCBool(substituteColor)

	c_r := (C.guchar)(r)

	c_g := (C.guchar)(g)

	c_b := (C.guchar)(b)

	ret := C.gdk_pixbuf_add_alpha(c_pixbuf, c_substituteColor, c_r, c_g, c_b)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_composite(src unsafe.Pointer, dest unsafe.Pointer, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType int, overallAlpha int) {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_dest := (*C.GdkPixbuf)(unsafe.Pointer(dest))

	c_destX := (C.int)(destX)

	c_destY := (C.int)(destY)

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_offsetX := (C.double)(offsetX)

	c_offsetY := (C.double)(offsetY)

	c_scaleX := (C.double)(scaleX)

	c_scaleY := (C.double)(scaleY)

	c_interpType := (C.GdkInterpType)(interpType)

	c_overallAlpha := (C.int)(overallAlpha)

	C.gdk_pixbuf_composite(c_src, c_dest, c_destX, c_destY, c_destWidth, c_destHeight, c_offsetX, c_offsetY, c_scaleX, c_scaleY, c_interpType, c_overallAlpha)
}

func Fn_gdk_pixbuf_composite_color(src unsafe.Pointer, dest unsafe.Pointer, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType int, overallAlpha int, checkX int, checkY int, checkSize int, color1 uint32, color2 uint32) {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_dest := (*C.GdkPixbuf)(unsafe.Pointer(dest))

	c_destX := (C.int)(destX)

	c_destY := (C.int)(destY)

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_offsetX := (C.double)(offsetX)

	c_offsetY := (C.double)(offsetY)

	c_scaleX := (C.double)(scaleX)

	c_scaleY := (C.double)(scaleY)

	c_interpType := (C.GdkInterpType)(interpType)

	c_overallAlpha := (C.int)(overallAlpha)

	c_checkX := (C.int)(checkX)

	c_checkY := (C.int)(checkY)

	c_checkSize := (C.int)(checkSize)

	c_color1 := (C.guint32)(color1)

	c_color2 := (C.guint32)(color2)

	C.gdk_pixbuf_composite_color(c_src, c_dest, c_destX, c_destY, c_destWidth, c_destHeight, c_offsetX, c_offsetY, c_scaleX, c_scaleY, c_interpType, c_overallAlpha, c_checkX, c_checkY, c_checkSize, c_color1, c_color2)
}

func Fn_gdk_pixbuf_composite_color_simple(src unsafe.Pointer, destWidth int, destHeight int, interpType int, overallAlpha int, checkSize int, color1 uint32, color2 uint32) unsafe.Pointer {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_interpType := (C.GdkInterpType)(interpType)

	c_overallAlpha := (C.int)(overallAlpha)

	c_checkSize := (C.int)(checkSize)

	c_color1 := (C.guint32)(color1)

	c_color2 := (C.guint32)(color2)

	ret := C.gdk_pixbuf_composite_color_simple(c_src, c_destWidth, c_destHeight, c_interpType, c_overallAlpha, c_checkSize, c_color1, c_color2)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_copy(pixbuf unsafe.Pointer) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_copy(c_pixbuf)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_copy_area(srcPixbuf unsafe.Pointer, srcX int, srcY int, width int, height int, destPixbuf unsafe.Pointer, destX int, destY int) {
	c_srcPixbuf := (*C.GdkPixbuf)(unsafe.Pointer(srcPixbuf))

	c_srcX := (C.int)(srcX)

	c_srcY := (C.int)(srcY)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_destPixbuf := (*C.GdkPixbuf)(unsafe.Pointer(destPixbuf))

	c_destX := (C.int)(destX)

	c_destY := (C.int)(destY)

	C.gdk_pixbuf_copy_area(c_srcPixbuf, c_srcX, c_srcY, c_width, c_height, c_destPixbuf, c_destX, c_destY)
}

func Fn_gdk_pixbuf_fill(pixbuf unsafe.Pointer, pixel uint32) {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_pixel := (C.guint32)(pixel)

	C.gdk_pixbuf_fill(c_pixbuf, c_pixel)
}

func Fn_gdk_pixbuf_get_bits_per_sample(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_bits_per_sample(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_colorspace(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_colorspace(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_has_alpha(pixbuf unsafe.Pointer) bool {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_has_alpha(c_pixbuf)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_get_height(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_height(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_n_channels(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_n_channels(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_option(pixbuf unsafe.Pointer, key string) string {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.gdk_pixbuf_get_option(c_pixbuf, c_key)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_pixels : no array length

func Fn_gdk_pixbuf_get_rowstride(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_rowstride(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_width(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_width(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_new_subpixbuf(srcPixbuf unsafe.Pointer, srcX int, srcY int, width int, height int) unsafe.Pointer {
	c_srcPixbuf := (*C.GdkPixbuf)(unsafe.Pointer(srcPixbuf))

	c_srcX := (C.int)(srcX)

	c_srcY := (C.int)(srcY)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	ret := C.gdk_pixbuf_new_subpixbuf(c_srcPixbuf, c_srcX, c_srcY, c_width, c_height)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_ref(pixbuf unsafe.Pointer) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_ref(c_pixbuf)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_saturate_and_pixelate(src unsafe.Pointer, dest unsafe.Pointer, saturation float32, pixelate bool) {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_dest := (*C.GdkPixbuf)(unsafe.Pointer(dest))

	c_saturation := (C.gfloat)(saturation)

	c_pixelate := toCBool(pixelate)

	C.gdk_pixbuf_saturate_and_pixelate(c_src, c_dest, c_saturation, c_pixelate)
}

// UNSUPPORTED : gdk_pixbuf_save : parameter 'error' is non array with indirect count > 1

// UNSUPPORTED : gdk_pixbuf_save_to_buffer : blacklisted

// UNSUPPORTED : gdk_pixbuf_save_to_bufferv : parameter 'buffer' is non array with indirect count > 1

// UNSUPPORTED : gdk_pixbuf_save_to_callback : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_callbackv : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_stream : parameter 'error' is non array with indirect count > 1

// UNSUPPORTED : gdk_pixbuf_save_to_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_streamv_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_savev(pixbuf unsafe.Pointer, filename string, type_ string, optionKeys []string, optionValues []string, error unsafe.Pointer) bool {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_filename := (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	c_type_ := (*C.char)(C.CString(type_))
	defer C.free(unsafe.Pointer(c_type_))

	optionKeysLen := len(optionKeys)
	c_optionKeysArray := C.malloc((C.ulong)(optionKeysLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_optionKeysArray))
	optionKeysSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_optionKeysArray))[:optionKeysLen:optionKeysLen]
	for optionKeysi, optionKeysString := range optionKeys {
		optionKeysSlice[optionKeysi] = (*C.gchar)(C.CString(optionKeysString))
		defer C.free(unsafe.Pointer(optionKeysSlice[optionKeysi]))
	}
	c_optionKeys := &optionKeysSlice[0]

	optionValuesLen := len(optionValues)
	c_optionValuesArray := C.malloc((C.ulong)(optionValuesLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_optionValuesArray))
	optionValuesSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_optionValuesArray))[:optionValuesLen:optionValuesLen]
	for optionValuesi, optionValuesString := range optionValues {
		optionValuesSlice[optionValuesi] = (*C.gchar)(C.CString(optionValuesString))
		defer C.free(unsafe.Pointer(optionValuesSlice[optionValuesi]))
	}
	c_optionValues := &optionValuesSlice[0]

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_savev(c_pixbuf, c_filename, c_type_, c_optionKeys, c_optionValues, cError)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_scale(src unsafe.Pointer, dest unsafe.Pointer, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType int) {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_dest := (*C.GdkPixbuf)(unsafe.Pointer(dest))

	c_destX := (C.int)(destX)

	c_destY := (C.int)(destY)

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_offsetX := (C.double)(offsetX)

	c_offsetY := (C.double)(offsetY)

	c_scaleX := (C.double)(scaleX)

	c_scaleY := (C.double)(scaleY)

	c_interpType := (C.GdkInterpType)(interpType)

	C.gdk_pixbuf_scale(c_src, c_dest, c_destX, c_destY, c_destWidth, c_destHeight, c_offsetX, c_offsetY, c_scaleX, c_scaleY, c_interpType)
}

func Fn_gdk_pixbuf_scale_simple(src unsafe.Pointer, destWidth int, destHeight int, interpType int) unsafe.Pointer {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_interpType := (C.GdkInterpType)(interpType)

	ret := C.gdk_pixbuf_scale_simple(c_src, c_destWidth, c_destHeight, c_interpType)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_unref(pixbuf unsafe.Pointer) {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	C.gdk_pixbuf_unref(c_pixbuf)
}

// UNSUPPORTED : gdk_pixbuf_get_file_info_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_at_scale_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_animation_new_from_file(filename string, error unsafe.Pointer) unsafe.Pointer {
	c_filename := (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_animation_new_from_file(c_filename, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_get_height(animation unsafe.Pointer) int {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_get_height(c_animation)

	return (int)(ret)
}

func Fn_gdk_pixbuf_animation_get_iter(animation unsafe.Pointer, startTime unsafe.Pointer) unsafe.Pointer {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	c_startTime := (*C.GTimeVal)(unsafe.Pointer(startTime))

	ret := C.gdk_pixbuf_animation_get_iter(c_animation, c_startTime)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_get_static_image(animation unsafe.Pointer) unsafe.Pointer {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_get_static_image(c_animation)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_get_width(animation unsafe.Pointer) int {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_get_width(c_animation)

	return (int)(ret)
}

func Fn_gdk_pixbuf_animation_is_static_image(animation unsafe.Pointer) bool {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_is_static_image(c_animation)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_animation_ref(animation unsafe.Pointer) unsafe.Pointer {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_ref(c_animation)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_unref(animation unsafe.Pointer) {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	C.gdk_pixbuf_animation_unref(c_animation)
}

// UNSUPPORTED : gdk_pixbuf_animation_new_from_stream_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_animation_iter_advance(iter unsafe.Pointer, currentTime unsafe.Pointer) bool {
	c_iter := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter))

	c_currentTime := (*C.GTimeVal)(unsafe.Pointer(currentTime))

	ret := C.gdk_pixbuf_animation_iter_advance(c_iter, c_currentTime)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_animation_iter_get_delay_time(iter unsafe.Pointer) int {
	c_iter := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter))

	ret := C.gdk_pixbuf_animation_iter_get_delay_time(c_iter)

	return (int)(ret)
}

func Fn_gdk_pixbuf_animation_iter_get_pixbuf(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter))

	ret := C.gdk_pixbuf_animation_iter_get_pixbuf(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_iter_on_currently_loading_frame(iter unsafe.Pointer) bool {
	c_iter := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter))

	ret := C.gdk_pixbuf_animation_iter_on_currently_loading_frame(c_iter)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_loader_new() unsafe.Pointer {
	ret := C.gdk_pixbuf_loader_new()

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_new_with_type(imageType string, error unsafe.Pointer) unsafe.Pointer {
	c_imageType := (*C.char)(C.CString(imageType))
	defer C.free(unsafe.Pointer(c_imageType))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_new_with_type(c_imageType, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_close(loader unsafe.Pointer, error unsafe.Pointer) bool {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_close(c_loader, cError)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_loader_get_animation(loader unsafe.Pointer) unsafe.Pointer {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	ret := C.gdk_pixbuf_loader_get_animation(c_loader)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_get_pixbuf(loader unsafe.Pointer) unsafe.Pointer {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	ret := C.gdk_pixbuf_loader_get_pixbuf(c_loader)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_write(loader unsafe.Pointer, buf []uint8, count uint64, error unsafe.Pointer) bool {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	c_buf := (*C.guchar)(unsafe.Pointer(&buf[0]))

	c_count := (C.gsize)(count)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_write(c_loader, c_buf, c_count, cError)

	return toGoBool(ret)
}
