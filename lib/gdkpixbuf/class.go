// Code generated - DO NOT EDIT.

package gdkpixbuf

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
)

var pixbufStruct *gi.Struct
var pixbufStruct_Once sync.Once

func pixbufStruct_Set() error {
	var err error
	pixbufStruct_Once.Do(func() {
		pixbufStruct, err = gi.StructNew("GdkPixbuf", "Pixbuf")
	})
	return err
}

type Pixbuf struct {
	native uintptr
}

// UNSUPPORTED : C value 'gdk_pixbuf_new' : parameter 'colorspace' of type 'Colorspace' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_bytes' : parameter 'data' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_data' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_file' : return type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_file_at_scale' : return type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_file_at_size' : return type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_inline' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_resource' : return type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_resource_at_scale' : return type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_stream' : parameter 'stream' of type 'Gio.InputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_stream_at_scale' : parameter 'stream' of type 'Gio.InputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_stream_finish' : parameter 'async_result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_xpm_data' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_add_alpha' : return type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_apply_embedded_orientation' : return type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_composite' : parameter 'dest' of type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_composite_color' : parameter 'dest' of type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_composite_color_simple' : parameter 'interp_type' of type 'InterpType' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_copy' : return type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_copy_area' : parameter 'dest_pixbuf' of type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_copy_options' : parameter 'dest_pixbuf' of type 'Pixbuf' not supported

var pixbufFillFunction *gi.Function
var pixbufFillFunction_Once sync.Once

func pixbufFillFunction_Set() error {
	var err error
	pixbufFillFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufFillFunction, err = pixbufStruct.InvokerNew("fill")
	})
	return err
}

// Fill is a representation of the C type gdk_pixbuf_fill.
func (recv *Pixbuf) Fill(pixel uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(pixel)

	err := pixbufFillFunction_Set()
	if err == nil {
		pixbufFillFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_pixbuf_flip' : return type 'Pixbuf' not supported

var pixbufGetBitsPerSampleFunction *gi.Function
var pixbufGetBitsPerSampleFunction_Once sync.Once

func pixbufGetBitsPerSampleFunction_Set() error {
	var err error
	pixbufGetBitsPerSampleFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetBitsPerSampleFunction, err = pixbufStruct.InvokerNew("get_bits_per_sample")
	})
	return err
}

// GetBitsPerSample is a representation of the C type gdk_pixbuf_get_bits_per_sample.
func (recv *Pixbuf) GetBitsPerSample() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufGetBitsPerSampleFunction_Set()
	if err == nil {
		ret = pixbufGetBitsPerSampleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var pixbufGetByteLengthFunction *gi.Function
var pixbufGetByteLengthFunction_Once sync.Once

func pixbufGetByteLengthFunction_Set() error {
	var err error
	pixbufGetByteLengthFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetByteLengthFunction, err = pixbufStruct.InvokerNew("get_byte_length")
	})
	return err
}

// GetByteLength is a representation of the C type gdk_pixbuf_get_byte_length.
func (recv *Pixbuf) GetByteLength() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufGetByteLengthFunction_Set()
	if err == nil {
		ret = pixbufGetByteLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_get_colorspace' : return type 'Colorspace' not supported

var pixbufGetHasAlphaFunction *gi.Function
var pixbufGetHasAlphaFunction_Once sync.Once

func pixbufGetHasAlphaFunction_Set() error {
	var err error
	pixbufGetHasAlphaFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetHasAlphaFunction, err = pixbufStruct.InvokerNew("get_has_alpha")
	})
	return err
}

// GetHasAlpha is a representation of the C type gdk_pixbuf_get_has_alpha.
func (recv *Pixbuf) GetHasAlpha() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufGetHasAlphaFunction_Set()
	if err == nil {
		ret = pixbufGetHasAlphaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var pixbufGetHeightFunction *gi.Function
var pixbufGetHeightFunction_Once sync.Once

func pixbufGetHeightFunction_Set() error {
	var err error
	pixbufGetHeightFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetHeightFunction, err = pixbufStruct.InvokerNew("get_height")
	})
	return err
}

// GetHeight is a representation of the C type gdk_pixbuf_get_height.
func (recv *Pixbuf) GetHeight() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufGetHeightFunction_Set()
	if err == nil {
		ret = pixbufGetHeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var pixbufGetNChannelsFunction *gi.Function
var pixbufGetNChannelsFunction_Once sync.Once

func pixbufGetNChannelsFunction_Set() error {
	var err error
	pixbufGetNChannelsFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetNChannelsFunction, err = pixbufStruct.InvokerNew("get_n_channels")
	})
	return err
}

// GetNChannels is a representation of the C type gdk_pixbuf_get_n_channels.
func (recv *Pixbuf) GetNChannels() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufGetNChannelsFunction_Set()
	if err == nil {
		ret = pixbufGetNChannelsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var pixbufGetOptionFunction *gi.Function
var pixbufGetOptionFunction_Once sync.Once

func pixbufGetOptionFunction_Set() error {
	var err error
	pixbufGetOptionFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetOptionFunction, err = pixbufStruct.InvokerNew("get_option")
	})
	return err
}

// GetOption is a representation of the C type gdk_pixbuf_get_option.
func (recv *Pixbuf) GetOption(key string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := pixbufGetOptionFunction_Set()
	if err == nil {
		ret = pixbufGetOptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_get_options' : return type 'GLib.HashTable' not supported

var pixbufGetPixelsFunction *gi.Function
var pixbufGetPixelsFunction_Once sync.Once

func pixbufGetPixelsFunction_Set() error {
	var err error
	pixbufGetPixelsFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetPixelsFunction, err = pixbufStruct.InvokerNew("get_pixels")
	})
	return err
}

// GetPixels is a representation of the C type gdk_pixbuf_get_pixels.
func (recv *Pixbuf) GetPixels() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := pixbufGetPixelsFunction_Set()
	if err == nil {
		pixbufGetPixelsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pixbufGetPixelsWithLengthFunction *gi.Function
var pixbufGetPixelsWithLengthFunction_Once sync.Once

func pixbufGetPixelsWithLengthFunction_Set() error {
	var err error
	pixbufGetPixelsWithLengthFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetPixelsWithLengthFunction, err = pixbufStruct.InvokerNew("get_pixels_with_length")
	})
	return err
}

// GetPixelsWithLength is a representation of the C type gdk_pixbuf_get_pixels_with_length.
func (recv *Pixbuf) GetPixelsWithLength() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := pixbufGetPixelsWithLengthFunction_Set()
	if err == nil {
		pixbufGetPixelsWithLengthFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
}

var pixbufGetRowstrideFunction *gi.Function
var pixbufGetRowstrideFunction_Once sync.Once

func pixbufGetRowstrideFunction_Set() error {
	var err error
	pixbufGetRowstrideFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetRowstrideFunction, err = pixbufStruct.InvokerNew("get_rowstride")
	})
	return err
}

// GetRowstride is a representation of the C type gdk_pixbuf_get_rowstride.
func (recv *Pixbuf) GetRowstride() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufGetRowstrideFunction_Set()
	if err == nil {
		ret = pixbufGetRowstrideFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var pixbufGetWidthFunction *gi.Function
var pixbufGetWidthFunction_Once sync.Once

func pixbufGetWidthFunction_Set() error {
	var err error
	pixbufGetWidthFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufGetWidthFunction, err = pixbufStruct.InvokerNew("get_width")
	})
	return err
}

// GetWidth is a representation of the C type gdk_pixbuf_get_width.
func (recv *Pixbuf) GetWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufGetWidthFunction_Set()
	if err == nil {
		ret = pixbufGetWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_new_subpixbuf' : return type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_read_pixel_bytes' : return type 'GLib.Bytes' not supported

var pixbufReadPixelsFunction *gi.Function
var pixbufReadPixelsFunction_Once sync.Once

func pixbufReadPixelsFunction_Set() error {
	var err error
	pixbufReadPixelsFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufReadPixelsFunction, err = pixbufStruct.InvokerNew("read_pixels")
	})
	return err
}

// ReadPixels is a representation of the C type gdk_pixbuf_read_pixels.
func (recv *Pixbuf) ReadPixels() uint8 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufReadPixelsFunction_Set()
	if err == nil {
		ret = pixbufReadPixelsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint8()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_ref' : return type 'Pixbuf' not supported

var pixbufRemoveOptionFunction *gi.Function
var pixbufRemoveOptionFunction_Once sync.Once

func pixbufRemoveOptionFunction_Set() error {
	var err error
	pixbufRemoveOptionFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufRemoveOptionFunction, err = pixbufStruct.InvokerNew("remove_option")
	})
	return err
}

// RemoveOption is a representation of the C type gdk_pixbuf_remove_option.
func (recv *Pixbuf) RemoveOption(key string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(key)

	var ret gi.Argument

	err := pixbufRemoveOptionFunction_Set()
	if err == nil {
		ret = pixbufRemoveOptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_rotate_simple' : parameter 'angle' of type 'PixbufRotation' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_saturate_and_pixelate' : parameter 'dest' of type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_save' : parameter 'error' of type 'GLib.Error' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_save_to_buffer' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_save_to_bufferv' : parameter 'buffer' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_save_to_callback' : parameter 'save_func' of type 'PixbufSaveFunc' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_save_to_callbackv' : parameter 'save_func' of type 'PixbufSaveFunc' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_save_to_stream' : parameter 'stream' of type 'Gio.OutputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_save_to_stream_async' : parameter 'stream' of type 'Gio.OutputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_save_to_streamv' : parameter 'stream' of type 'Gio.OutputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_save_to_streamv_async' : parameter 'stream' of type 'Gio.OutputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_savev' : parameter 'option_keys' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_scale' : parameter 'dest' of type 'Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_scale_simple' : parameter 'interp_type' of type 'InterpType' not supported

var pixbufSetOptionFunction *gi.Function
var pixbufSetOptionFunction_Once sync.Once

func pixbufSetOptionFunction_Set() error {
	var err error
	pixbufSetOptionFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufSetOptionFunction, err = pixbufStruct.InvokerNew("set_option")
	})
	return err
}

// SetOption is a representation of the C type gdk_pixbuf_set_option.
func (recv *Pixbuf) SetOption(key string, value string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(key)
	inArgs[2].SetString(value)

	var ret gi.Argument

	err := pixbufSetOptionFunction_Set()
	if err == nil {
		ret = pixbufSetOptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var pixbufUnrefFunction *gi.Function
var pixbufUnrefFunction_Once sync.Once

func pixbufUnrefFunction_Set() error {
	var err error
	pixbufUnrefFunction_Once.Do(func() {
		err = pixbufStruct_Set()
		if err != nil {
			return
		}
		pixbufUnrefFunction, err = pixbufStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gdk_pixbuf_unref.
func (recv *Pixbuf) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := pixbufUnrefFunction_Set()
	if err == nil {
		pixbufUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pixbufAnimationStruct *gi.Struct
var pixbufAnimationStruct_Once sync.Once

func pixbufAnimationStruct_Set() error {
	var err error
	pixbufAnimationStruct_Once.Do(func() {
		pixbufAnimationStruct, err = gi.StructNew("GdkPixbuf", "PixbufAnimation")
	})
	return err
}

type PixbufAnimation struct {
	native uintptr
}

// UNSUPPORTED : C value 'gdk_pixbuf_animation_new_from_file' : return type 'PixbufAnimation' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_animation_new_from_resource' : return type 'PixbufAnimation' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_animation_new_from_stream' : parameter 'stream' of type 'Gio.InputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_animation_new_from_stream_finish' : parameter 'async_result' of type 'Gio.AsyncResult' not supported

var pixbufAnimationGetHeightFunction *gi.Function
var pixbufAnimationGetHeightFunction_Once sync.Once

func pixbufAnimationGetHeightFunction_Set() error {
	var err error
	pixbufAnimationGetHeightFunction_Once.Do(func() {
		err = pixbufAnimationStruct_Set()
		if err != nil {
			return
		}
		pixbufAnimationGetHeightFunction, err = pixbufAnimationStruct.InvokerNew("get_height")
	})
	return err
}

// GetHeight is a representation of the C type gdk_pixbuf_animation_get_height.
func (recv *PixbufAnimation) GetHeight() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufAnimationGetHeightFunction_Set()
	if err == nil {
		ret = pixbufAnimationGetHeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_animation_get_iter' : parameter 'start_time' of type 'GLib.TimeVal' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_animation_get_static_image' : return type 'Pixbuf' not supported

var pixbufAnimationGetWidthFunction *gi.Function
var pixbufAnimationGetWidthFunction_Once sync.Once

func pixbufAnimationGetWidthFunction_Set() error {
	var err error
	pixbufAnimationGetWidthFunction_Once.Do(func() {
		err = pixbufAnimationStruct_Set()
		if err != nil {
			return
		}
		pixbufAnimationGetWidthFunction, err = pixbufAnimationStruct.InvokerNew("get_width")
	})
	return err
}

// GetWidth is a representation of the C type gdk_pixbuf_animation_get_width.
func (recv *PixbufAnimation) GetWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufAnimationGetWidthFunction_Set()
	if err == nil {
		ret = pixbufAnimationGetWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var pixbufAnimationIsStaticImageFunction *gi.Function
var pixbufAnimationIsStaticImageFunction_Once sync.Once

func pixbufAnimationIsStaticImageFunction_Set() error {
	var err error
	pixbufAnimationIsStaticImageFunction_Once.Do(func() {
		err = pixbufAnimationStruct_Set()
		if err != nil {
			return
		}
		pixbufAnimationIsStaticImageFunction, err = pixbufAnimationStruct.InvokerNew("is_static_image")
	})
	return err
}

// IsStaticImage is a representation of the C type gdk_pixbuf_animation_is_static_image.
func (recv *PixbufAnimation) IsStaticImage() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufAnimationIsStaticImageFunction_Set()
	if err == nil {
		ret = pixbufAnimationIsStaticImageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_animation_ref' : return type 'PixbufAnimation' not supported

var pixbufAnimationUnrefFunction *gi.Function
var pixbufAnimationUnrefFunction_Once sync.Once

func pixbufAnimationUnrefFunction_Set() error {
	var err error
	pixbufAnimationUnrefFunction_Once.Do(func() {
		err = pixbufAnimationStruct_Set()
		if err != nil {
			return
		}
		pixbufAnimationUnrefFunction, err = pixbufAnimationStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gdk_pixbuf_animation_unref.
func (recv *PixbufAnimation) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := pixbufAnimationUnrefFunction_Set()
	if err == nil {
		pixbufAnimationUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pixbufAnimationIterStruct *gi.Struct
var pixbufAnimationIterStruct_Once sync.Once

func pixbufAnimationIterStruct_Set() error {
	var err error
	pixbufAnimationIterStruct_Once.Do(func() {
		pixbufAnimationIterStruct, err = gi.StructNew("GdkPixbuf", "PixbufAnimationIter")
	})
	return err
}

type PixbufAnimationIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'gdk_pixbuf_animation_iter_advance' : parameter 'current_time' of type 'GLib.TimeVal' not supported

var pixbufAnimationIterGetDelayTimeFunction *gi.Function
var pixbufAnimationIterGetDelayTimeFunction_Once sync.Once

func pixbufAnimationIterGetDelayTimeFunction_Set() error {
	var err error
	pixbufAnimationIterGetDelayTimeFunction_Once.Do(func() {
		err = pixbufAnimationIterStruct_Set()
		if err != nil {
			return
		}
		pixbufAnimationIterGetDelayTimeFunction, err = pixbufAnimationIterStruct.InvokerNew("get_delay_time")
	})
	return err
}

// GetDelayTime is a representation of the C type gdk_pixbuf_animation_iter_get_delay_time.
func (recv *PixbufAnimationIter) GetDelayTime() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufAnimationIterGetDelayTimeFunction_Set()
	if err == nil {
		ret = pixbufAnimationIterGetDelayTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_animation_iter_get_pixbuf' : return type 'Pixbuf' not supported

var pixbufAnimationIterOnCurrentlyLoadingFrameFunction *gi.Function
var pixbufAnimationIterOnCurrentlyLoadingFrameFunction_Once sync.Once

func pixbufAnimationIterOnCurrentlyLoadingFrameFunction_Set() error {
	var err error
	pixbufAnimationIterOnCurrentlyLoadingFrameFunction_Once.Do(func() {
		err = pixbufAnimationIterStruct_Set()
		if err != nil {
			return
		}
		pixbufAnimationIterOnCurrentlyLoadingFrameFunction, err = pixbufAnimationIterStruct.InvokerNew("on_currently_loading_frame")
	})
	return err
}

// OnCurrentlyLoadingFrame is a representation of the C type gdk_pixbuf_animation_iter_on_currently_loading_frame.
func (recv *PixbufAnimationIter) OnCurrentlyLoadingFrame() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufAnimationIterOnCurrentlyLoadingFrameFunction_Set()
	if err == nil {
		ret = pixbufAnimationIterOnCurrentlyLoadingFrameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// PixbufAnimationIterStruct creates an uninitialised PixbufAnimationIter.
func PixbufAnimationIterStruct() *PixbufAnimationIter {
	err := pixbufAnimationIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PixbufAnimationIter{native: pixbufAnimationIterStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePixbufAnimationIter)
	return structGo
}
func finalizePixbufAnimationIter(obj *PixbufAnimationIter) {
	pixbufAnimationIterStruct.Free(obj.native)
}

var pixbufLoaderStruct *gi.Struct
var pixbufLoaderStruct_Once sync.Once

func pixbufLoaderStruct_Set() error {
	var err error
	pixbufLoaderStruct_Once.Do(func() {
		pixbufLoaderStruct, err = gi.StructNew("GdkPixbuf", "PixbufLoader")
	})
	return err
}

type PixbufLoader struct {
	native uintptr
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'gdk_pixbuf_loader_new' : return type 'PixbufLoader' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_loader_new_with_mime_type' : return type 'PixbufLoader' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_loader_new_with_type' : return type 'PixbufLoader' not supported

var pixbufLoaderCloseFunction *gi.Function
var pixbufLoaderCloseFunction_Once sync.Once

func pixbufLoaderCloseFunction_Set() error {
	var err error
	pixbufLoaderCloseFunction_Once.Do(func() {
		err = pixbufLoaderStruct_Set()
		if err != nil {
			return
		}
		pixbufLoaderCloseFunction, err = pixbufLoaderStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type gdk_pixbuf_loader_close.
func (recv *PixbufLoader) Close() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufLoaderCloseFunction_Set()
	if err == nil {
		ret = pixbufLoaderCloseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_loader_get_animation' : return type 'PixbufAnimation' not supported

var pixbufLoaderGetFormatFunction *gi.Function
var pixbufLoaderGetFormatFunction_Once sync.Once

func pixbufLoaderGetFormatFunction_Set() error {
	var err error
	pixbufLoaderGetFormatFunction_Once.Do(func() {
		err = pixbufLoaderStruct_Set()
		if err != nil {
			return
		}
		pixbufLoaderGetFormatFunction, err = pixbufLoaderStruct.InvokerNew("get_format")
	})
	return err
}

// GetFormat is a representation of the C type gdk_pixbuf_loader_get_format.
func (recv *PixbufLoader) GetFormat() *PixbufFormat {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufLoaderGetFormatFunction_Set()
	if err == nil {
		ret = pixbufLoaderGetFormatFunction.Invoke(inArgs[:], nil)
	}

	retGo := &PixbufFormat{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_loader_get_pixbuf' : return type 'Pixbuf' not supported

var pixbufLoaderSetSizeFunction *gi.Function
var pixbufLoaderSetSizeFunction_Once sync.Once

func pixbufLoaderSetSizeFunction_Set() error {
	var err error
	pixbufLoaderSetSizeFunction_Once.Do(func() {
		err = pixbufLoaderStruct_Set()
		if err != nil {
			return
		}
		pixbufLoaderSetSizeFunction, err = pixbufLoaderStruct.InvokerNew("set_size")
	})
	return err
}

// SetSize is a representation of the C type gdk_pixbuf_loader_set_size.
func (recv *PixbufLoader) SetSize(width int32, height int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(width)
	inArgs[2].SetInt32(height)

	err := pixbufLoaderSetSizeFunction_Set()
	if err == nil {
		pixbufLoaderSetSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_pixbuf_loader_write' : parameter 'buf' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_loader_write_bytes' : parameter 'buffer' of type 'GLib.Bytes' not supported

var pixbufSimpleAnimStruct *gi.Struct
var pixbufSimpleAnimStruct_Once sync.Once

func pixbufSimpleAnimStruct_Set() error {
	var err error
	pixbufSimpleAnimStruct_Once.Do(func() {
		pixbufSimpleAnimStruct, err = gi.StructNew("GdkPixbuf", "PixbufSimpleAnim")
	})
	return err
}

type PixbufSimpleAnim struct {
	native uintptr
}

// UNSUPPORTED : C value 'gdk_pixbuf_simple_anim_new' : return type 'PixbufSimpleAnim' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_simple_anim_add_frame' : parameter 'pixbuf' of type 'Pixbuf' not supported

var pixbufSimpleAnimGetLoopFunction *gi.Function
var pixbufSimpleAnimGetLoopFunction_Once sync.Once

func pixbufSimpleAnimGetLoopFunction_Set() error {
	var err error
	pixbufSimpleAnimGetLoopFunction_Once.Do(func() {
		err = pixbufSimpleAnimStruct_Set()
		if err != nil {
			return
		}
		pixbufSimpleAnimGetLoopFunction, err = pixbufSimpleAnimStruct.InvokerNew("get_loop")
	})
	return err
}

// GetLoop is a representation of the C type gdk_pixbuf_simple_anim_get_loop.
func (recv *PixbufSimpleAnim) GetLoop() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufSimpleAnimGetLoopFunction_Set()
	if err == nil {
		ret = pixbufSimpleAnimGetLoopFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var pixbufSimpleAnimSetLoopFunction *gi.Function
var pixbufSimpleAnimSetLoopFunction_Once sync.Once

func pixbufSimpleAnimSetLoopFunction_Set() error {
	var err error
	pixbufSimpleAnimSetLoopFunction_Once.Do(func() {
		err = pixbufSimpleAnimStruct_Set()
		if err != nil {
			return
		}
		pixbufSimpleAnimSetLoopFunction, err = pixbufSimpleAnimStruct.InvokerNew("set_loop")
	})
	return err
}

// SetLoop is a representation of the C type gdk_pixbuf_simple_anim_set_loop.
func (recv *PixbufSimpleAnim) SetLoop(loop bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(loop)

	err := pixbufSimpleAnimSetLoopFunction_Set()
	if err == nil {
		pixbufSimpleAnimSetLoopFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pixbufSimpleAnimIterStruct *gi.Struct
var pixbufSimpleAnimIterStruct_Once sync.Once

func pixbufSimpleAnimIterStruct_Set() error {
	var err error
	pixbufSimpleAnimIterStruct_Once.Do(func() {
		pixbufSimpleAnimIterStruct, err = gi.StructNew("GdkPixbuf", "PixbufSimpleAnimIter")
	})
	return err
}

type PixbufSimpleAnimIter struct {
	native uintptr
}

// PixbufSimpleAnimIterStruct creates an uninitialised PixbufSimpleAnimIter.
func PixbufSimpleAnimIterStruct() *PixbufSimpleAnimIter {
	err := pixbufSimpleAnimIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PixbufSimpleAnimIter{native: pixbufSimpleAnimIterStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePixbufSimpleAnimIter)
	return structGo
}
func finalizePixbufSimpleAnimIter(obj *PixbufSimpleAnimIter) {
	pixbufSimpleAnimIterStruct.Free(obj.native)
}
