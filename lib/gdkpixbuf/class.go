// Code generated - DO NOT EDIT.

package gdkpixbuf

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

var pixbufObject *gi.Object
var pixbufObject_Once sync.Once

func pixbufObject_Set() error {
	var err error
	pixbufObject_Once.Do(func() {
		pixbufObject, err = gi.ObjectNew("GdkPixbuf", "Pixbuf")
	})
	return err
}

type Pixbuf struct {
	native unsafe.Pointer
}

func PixbufNewFromNative(native unsafe.Pointer) *Pixbuf {
	instance := &Pixbuf{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *Pixbuf) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

/*
CastToPixbuf down casts any arbitrary Object to Pixbuf.
Exercise care, as this is a potentially dangerous function
if the Object is not a Pixbuf.
*/
func (recv *Pixbuf) CastToPixbuf(object *gobject.Object) *Pixbuf {
	return PixbufNewFromNative(object.Native())
}

func (recv *Pixbuf) Native() unsafe.Pointer {
	return recv.native
}

var pixbufNewFunction *gi.Function
var pixbufNewFunction_Once sync.Once

func pixbufNewFunction_Set() error {
	var err error
	pixbufNewFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufNewFunction, err = pixbufObject.InvokerNew("new")
	})
	return err
}

// PixbufNew is a representation of the C type gdk_pixbuf_new.
func PixbufNew(colorspace Colorspace, hasAlpha bool, bitsPerSample int32, width int32, height int32) *Pixbuf {
	var inArgs [5]gi.Argument
	inArgs[0].SetInt32(int32(colorspace))
	inArgs[1].SetBoolean(hasAlpha)
	inArgs[2].SetInt32(bitsPerSample)
	inArgs[3].SetInt32(width)
	inArgs[4].SetInt32(height)

	var ret gi.Argument

	err := pixbufNewFunction_Set()
	if err == nil {
		ret = pixbufNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_bytes' : parameter 'data' of type 'GLib.Bytes' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_data' : parameter 'data' of type 'nil' not supported

var pixbufNewFromFileFunction *gi.Function
var pixbufNewFromFileFunction_Once sync.Once

func pixbufNewFromFileFunction_Set() error {
	var err error
	pixbufNewFromFileFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufNewFromFileFunction, err = pixbufObject.InvokerNew("new_from_file")
	})
	return err
}

// PixbufNewFromFile is a representation of the C type gdk_pixbuf_new_from_file.
func PixbufNewFromFile(filename string) *Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := pixbufNewFromFileFunction_Set()
	if err == nil {
		ret = pixbufNewFromFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var pixbufNewFromFileAtScaleFunction *gi.Function
var pixbufNewFromFileAtScaleFunction_Once sync.Once

func pixbufNewFromFileAtScaleFunction_Set() error {
	var err error
	pixbufNewFromFileAtScaleFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufNewFromFileAtScaleFunction, err = pixbufObject.InvokerNew("new_from_file_at_scale")
	})
	return err
}

// PixbufNewFromFileAtScale is a representation of the C type gdk_pixbuf_new_from_file_at_scale.
func PixbufNewFromFileAtScale(filename string, width int32, height int32, preserveAspectRatio bool) *Pixbuf {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(filename)
	inArgs[1].SetInt32(width)
	inArgs[2].SetInt32(height)
	inArgs[3].SetBoolean(preserveAspectRatio)

	var ret gi.Argument

	err := pixbufNewFromFileAtScaleFunction_Set()
	if err == nil {
		ret = pixbufNewFromFileAtScaleFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var pixbufNewFromFileAtSizeFunction *gi.Function
var pixbufNewFromFileAtSizeFunction_Once sync.Once

func pixbufNewFromFileAtSizeFunction_Set() error {
	var err error
	pixbufNewFromFileAtSizeFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufNewFromFileAtSizeFunction, err = pixbufObject.InvokerNew("new_from_file_at_size")
	})
	return err
}

// PixbufNewFromFileAtSize is a representation of the C type gdk_pixbuf_new_from_file_at_size.
func PixbufNewFromFileAtSize(filename string, width int32, height int32) *Pixbuf {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(filename)
	inArgs[1].SetInt32(width)
	inArgs[2].SetInt32(height)

	var ret gi.Argument

	err := pixbufNewFromFileAtSizeFunction_Set()
	if err == nil {
		ret = pixbufNewFromFileAtSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_inline' : parameter 'data' of type 'nil' not supported

var pixbufNewFromResourceFunction *gi.Function
var pixbufNewFromResourceFunction_Once sync.Once

func pixbufNewFromResourceFunction_Set() error {
	var err error
	pixbufNewFromResourceFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufNewFromResourceFunction, err = pixbufObject.InvokerNew("new_from_resource")
	})
	return err
}

// PixbufNewFromResource is a representation of the C type gdk_pixbuf_new_from_resource.
func PixbufNewFromResource(resourcePath string) *Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(resourcePath)

	var ret gi.Argument

	err := pixbufNewFromResourceFunction_Set()
	if err == nil {
		ret = pixbufNewFromResourceFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var pixbufNewFromResourceAtScaleFunction *gi.Function
var pixbufNewFromResourceAtScaleFunction_Once sync.Once

func pixbufNewFromResourceAtScaleFunction_Set() error {
	var err error
	pixbufNewFromResourceAtScaleFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufNewFromResourceAtScaleFunction, err = pixbufObject.InvokerNew("new_from_resource_at_scale")
	})
	return err
}

// PixbufNewFromResourceAtScale is a representation of the C type gdk_pixbuf_new_from_resource_at_scale.
func PixbufNewFromResourceAtScale(resourcePath string, width int32, height int32, preserveAspectRatio bool) *Pixbuf {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(resourcePath)
	inArgs[1].SetInt32(width)
	inArgs[2].SetInt32(height)
	inArgs[3].SetBoolean(preserveAspectRatio)

	var ret gi.Argument

	err := pixbufNewFromResourceAtScaleFunction_Set()
	if err == nil {
		ret = pixbufNewFromResourceAtScaleFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_stream' : parameter 'stream' of type 'Gio.InputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_stream_at_scale' : parameter 'stream' of type 'Gio.InputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_stream_finish' : parameter 'async_result' of type 'Gio.AsyncResult' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_new_from_xpm_data' : parameter 'data' of type 'nil' not supported

var pixbufAddAlphaFunction *gi.Function
var pixbufAddAlphaFunction_Once sync.Once

func pixbufAddAlphaFunction_Set() error {
	var err error
	pixbufAddAlphaFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufAddAlphaFunction, err = pixbufObject.InvokerNew("add_alpha")
	})
	return err
}

// AddAlpha is a representation of the C type gdk_pixbuf_add_alpha.
func (recv *Pixbuf) AddAlpha(substituteColor bool, r uint8, g uint8, b uint8) *Pixbuf {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(substituteColor)
	inArgs[2].SetUint8(r)
	inArgs[3].SetUint8(g)
	inArgs[4].SetUint8(b)

	var ret gi.Argument

	err := pixbufAddAlphaFunction_Set()
	if err == nil {
		ret = pixbufAddAlphaFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufApplyEmbeddedOrientationFunction *gi.Function
var pixbufApplyEmbeddedOrientationFunction_Once sync.Once

func pixbufApplyEmbeddedOrientationFunction_Set() error {
	var err error
	pixbufApplyEmbeddedOrientationFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufApplyEmbeddedOrientationFunction, err = pixbufObject.InvokerNew("apply_embedded_orientation")
	})
	return err
}

// ApplyEmbeddedOrientation is a representation of the C type gdk_pixbuf_apply_embedded_orientation.
func (recv *Pixbuf) ApplyEmbeddedOrientation() *Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufApplyEmbeddedOrientationFunction_Set()
	if err == nil {
		ret = pixbufApplyEmbeddedOrientationFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufCompositeFunction *gi.Function
var pixbufCompositeFunction_Once sync.Once

func pixbufCompositeFunction_Set() error {
	var err error
	pixbufCompositeFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufCompositeFunction, err = pixbufObject.InvokerNew("composite")
	})
	return err
}

// Composite is a representation of the C type gdk_pixbuf_composite.
func (recv *Pixbuf) Composite(dest *Pixbuf, destX int32, destY int32, destWidth int32, destHeight int32, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int32) {
	var inArgs [12]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(dest.native)
	inArgs[2].SetInt32(destX)
	inArgs[3].SetInt32(destY)
	inArgs[4].SetInt32(destWidth)
	inArgs[5].SetInt32(destHeight)
	inArgs[6].SetFloat64(offsetX)
	inArgs[7].SetFloat64(offsetY)
	inArgs[8].SetFloat64(scaleX)
	inArgs[9].SetFloat64(scaleY)
	inArgs[10].SetInt32(int32(interpType))
	inArgs[11].SetInt32(overallAlpha)

	err := pixbufCompositeFunction_Set()
	if err == nil {
		pixbufCompositeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pixbufCompositeColorFunction *gi.Function
var pixbufCompositeColorFunction_Once sync.Once

func pixbufCompositeColorFunction_Set() error {
	var err error
	pixbufCompositeColorFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufCompositeColorFunction, err = pixbufObject.InvokerNew("composite_color")
	})
	return err
}

// CompositeColor is a representation of the C type gdk_pixbuf_composite_color.
func (recv *Pixbuf) CompositeColor(dest *Pixbuf, destX int32, destY int32, destWidth int32, destHeight int32, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int32, checkX int32, checkY int32, checkSize int32, color1 uint32, color2 uint32) {
	var inArgs [17]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(dest.native)
	inArgs[2].SetInt32(destX)
	inArgs[3].SetInt32(destY)
	inArgs[4].SetInt32(destWidth)
	inArgs[5].SetInt32(destHeight)
	inArgs[6].SetFloat64(offsetX)
	inArgs[7].SetFloat64(offsetY)
	inArgs[8].SetFloat64(scaleX)
	inArgs[9].SetFloat64(scaleY)
	inArgs[10].SetInt32(int32(interpType))
	inArgs[11].SetInt32(overallAlpha)
	inArgs[12].SetInt32(checkX)
	inArgs[13].SetInt32(checkY)
	inArgs[14].SetInt32(checkSize)
	inArgs[15].SetUint32(color1)
	inArgs[16].SetUint32(color2)

	err := pixbufCompositeColorFunction_Set()
	if err == nil {
		pixbufCompositeColorFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pixbufCompositeColorSimpleFunction *gi.Function
var pixbufCompositeColorSimpleFunction_Once sync.Once

func pixbufCompositeColorSimpleFunction_Set() error {
	var err error
	pixbufCompositeColorSimpleFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufCompositeColorSimpleFunction, err = pixbufObject.InvokerNew("composite_color_simple")
	})
	return err
}

// CompositeColorSimple is a representation of the C type gdk_pixbuf_composite_color_simple.
func (recv *Pixbuf) CompositeColorSimple(destWidth int32, destHeight int32, interpType InterpType, overallAlpha int32, checkSize int32, color1 uint32, color2 uint32) *Pixbuf {
	var inArgs [8]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(destWidth)
	inArgs[2].SetInt32(destHeight)
	inArgs[3].SetInt32(int32(interpType))
	inArgs[4].SetInt32(overallAlpha)
	inArgs[5].SetInt32(checkSize)
	inArgs[6].SetUint32(color1)
	inArgs[7].SetUint32(color2)

	var ret gi.Argument

	err := pixbufCompositeColorSimpleFunction_Set()
	if err == nil {
		ret = pixbufCompositeColorSimpleFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufCopyFunction *gi.Function
var pixbufCopyFunction_Once sync.Once

func pixbufCopyFunction_Set() error {
	var err error
	pixbufCopyFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufCopyFunction, err = pixbufObject.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gdk_pixbuf_copy.
func (recv *Pixbuf) Copy() *Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufCopyFunction_Set()
	if err == nil {
		ret = pixbufCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufCopyAreaFunction *gi.Function
var pixbufCopyAreaFunction_Once sync.Once

func pixbufCopyAreaFunction_Set() error {
	var err error
	pixbufCopyAreaFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufCopyAreaFunction, err = pixbufObject.InvokerNew("copy_area")
	})
	return err
}

// CopyArea is a representation of the C type gdk_pixbuf_copy_area.
func (recv *Pixbuf) CopyArea(srcX int32, srcY int32, width int32, height int32, destPixbuf *Pixbuf, destX int32, destY int32) {
	var inArgs [8]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(srcX)
	inArgs[2].SetInt32(srcY)
	inArgs[3].SetInt32(width)
	inArgs[4].SetInt32(height)
	inArgs[5].SetPointer(destPixbuf.native)
	inArgs[6].SetInt32(destX)
	inArgs[7].SetInt32(destY)

	err := pixbufCopyAreaFunction_Set()
	if err == nil {
		pixbufCopyAreaFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pixbufCopyOptionsFunction *gi.Function
var pixbufCopyOptionsFunction_Once sync.Once

func pixbufCopyOptionsFunction_Set() error {
	var err error
	pixbufCopyOptionsFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufCopyOptionsFunction, err = pixbufObject.InvokerNew("copy_options")
	})
	return err
}

// CopyOptions is a representation of the C type gdk_pixbuf_copy_options.
func (recv *Pixbuf) CopyOptions(destPixbuf *Pixbuf) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(destPixbuf.native)

	var ret gi.Argument

	err := pixbufCopyOptionsFunction_Set()
	if err == nil {
		ret = pixbufCopyOptionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var pixbufFillFunction *gi.Function
var pixbufFillFunction_Once sync.Once

func pixbufFillFunction_Set() error {
	var err error
	pixbufFillFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufFillFunction, err = pixbufObject.InvokerNew("fill")
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

var pixbufFlipFunction *gi.Function
var pixbufFlipFunction_Once sync.Once

func pixbufFlipFunction_Set() error {
	var err error
	pixbufFlipFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufFlipFunction, err = pixbufObject.InvokerNew("flip")
	})
	return err
}

// Flip is a representation of the C type gdk_pixbuf_flip.
func (recv *Pixbuf) Flip(horizontal bool) *Pixbuf {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(horizontal)

	var ret gi.Argument

	err := pixbufFlipFunction_Set()
	if err == nil {
		ret = pixbufFlipFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufGetBitsPerSampleFunction *gi.Function
var pixbufGetBitsPerSampleFunction_Once sync.Once

func pixbufGetBitsPerSampleFunction_Set() error {
	var err error
	pixbufGetBitsPerSampleFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetBitsPerSampleFunction, err = pixbufObject.InvokerNew("get_bits_per_sample")
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
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetByteLengthFunction, err = pixbufObject.InvokerNew("get_byte_length")
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

var pixbufGetColorspaceFunction *gi.Function
var pixbufGetColorspaceFunction_Once sync.Once

func pixbufGetColorspaceFunction_Set() error {
	var err error
	pixbufGetColorspaceFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetColorspaceFunction, err = pixbufObject.InvokerNew("get_colorspace")
	})
	return err
}

// GetColorspace is a representation of the C type gdk_pixbuf_get_colorspace.
func (recv *Pixbuf) GetColorspace() Colorspace {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufGetColorspaceFunction_Set()
	if err == nil {
		ret = pixbufGetColorspaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := Colorspace(ret.Int32())

	return retGo
}

var pixbufGetHasAlphaFunction *gi.Function
var pixbufGetHasAlphaFunction_Once sync.Once

func pixbufGetHasAlphaFunction_Set() error {
	var err error
	pixbufGetHasAlphaFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetHasAlphaFunction, err = pixbufObject.InvokerNew("get_has_alpha")
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
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetHeightFunction, err = pixbufObject.InvokerNew("get_height")
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
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetNChannelsFunction, err = pixbufObject.InvokerNew("get_n_channels")
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
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetOptionFunction, err = pixbufObject.InvokerNew("get_option")
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
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetPixelsFunction, err = pixbufObject.InvokerNew("get_pixels")
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
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetPixelsWithLengthFunction, err = pixbufObject.InvokerNew("get_pixels_with_length")
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
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetRowstrideFunction, err = pixbufObject.InvokerNew("get_rowstride")
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
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufGetWidthFunction, err = pixbufObject.InvokerNew("get_width")
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

var pixbufNewSubpixbufFunction *gi.Function
var pixbufNewSubpixbufFunction_Once sync.Once

func pixbufNewSubpixbufFunction_Set() error {
	var err error
	pixbufNewSubpixbufFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufNewSubpixbufFunction, err = pixbufObject.InvokerNew("new_subpixbuf")
	})
	return err
}

// NewSubpixbuf is a representation of the C type gdk_pixbuf_new_subpixbuf.
func (recv *Pixbuf) NewSubpixbuf(srcX int32, srcY int32, width int32, height int32) *Pixbuf {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(srcX)
	inArgs[2].SetInt32(srcY)
	inArgs[3].SetInt32(width)
	inArgs[4].SetInt32(height)

	var ret gi.Argument

	err := pixbufNewSubpixbufFunction_Set()
	if err == nil {
		ret = pixbufNewSubpixbufFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_read_pixel_bytes' : return type 'GLib.Bytes' not supported

var pixbufReadPixelsFunction *gi.Function
var pixbufReadPixelsFunction_Once sync.Once

func pixbufReadPixelsFunction_Set() error {
	var err error
	pixbufReadPixelsFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufReadPixelsFunction, err = pixbufObject.InvokerNew("read_pixels")
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

var pixbufRefFunction *gi.Function
var pixbufRefFunction_Once sync.Once

func pixbufRefFunction_Set() error {
	var err error
	pixbufRefFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufRefFunction, err = pixbufObject.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gdk_pixbuf_ref.
func (recv *Pixbuf) Ref() *Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufRefFunction_Set()
	if err == nil {
		ret = pixbufRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufRemoveOptionFunction *gi.Function
var pixbufRemoveOptionFunction_Once sync.Once

func pixbufRemoveOptionFunction_Set() error {
	var err error
	pixbufRemoveOptionFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufRemoveOptionFunction, err = pixbufObject.InvokerNew("remove_option")
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

var pixbufRotateSimpleFunction *gi.Function
var pixbufRotateSimpleFunction_Once sync.Once

func pixbufRotateSimpleFunction_Set() error {
	var err error
	pixbufRotateSimpleFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufRotateSimpleFunction, err = pixbufObject.InvokerNew("rotate_simple")
	})
	return err
}

// RotateSimple is a representation of the C type gdk_pixbuf_rotate_simple.
func (recv *Pixbuf) RotateSimple(angle PixbufRotation) *Pixbuf {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(angle))

	var ret gi.Argument

	err := pixbufRotateSimpleFunction_Set()
	if err == nil {
		ret = pixbufRotateSimpleFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufSaturateAndPixelateFunction *gi.Function
var pixbufSaturateAndPixelateFunction_Once sync.Once

func pixbufSaturateAndPixelateFunction_Set() error {
	var err error
	pixbufSaturateAndPixelateFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufSaturateAndPixelateFunction, err = pixbufObject.InvokerNew("saturate_and_pixelate")
	})
	return err
}

// SaturateAndPixelate is a representation of the C type gdk_pixbuf_saturate_and_pixelate.
func (recv *Pixbuf) SaturateAndPixelate(dest *Pixbuf, saturation float32, pixelate bool) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(dest.native)
	inArgs[2].SetFloat32(saturation)
	inArgs[3].SetBoolean(pixelate)

	err := pixbufSaturateAndPixelateFunction_Set()
	if err == nil {
		pixbufSaturateAndPixelateFunction.Invoke(inArgs[:], nil)
	}

	return
}

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

var pixbufScaleFunction *gi.Function
var pixbufScaleFunction_Once sync.Once

func pixbufScaleFunction_Set() error {
	var err error
	pixbufScaleFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufScaleFunction, err = pixbufObject.InvokerNew("scale")
	})
	return err
}

// Scale is a representation of the C type gdk_pixbuf_scale.
func (recv *Pixbuf) Scale(dest *Pixbuf, destX int32, destY int32, destWidth int32, destHeight int32, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType) {
	var inArgs [11]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(dest.native)
	inArgs[2].SetInt32(destX)
	inArgs[3].SetInt32(destY)
	inArgs[4].SetInt32(destWidth)
	inArgs[5].SetInt32(destHeight)
	inArgs[6].SetFloat64(offsetX)
	inArgs[7].SetFloat64(offsetY)
	inArgs[8].SetFloat64(scaleX)
	inArgs[9].SetFloat64(scaleY)
	inArgs[10].SetInt32(int32(interpType))

	err := pixbufScaleFunction_Set()
	if err == nil {
		pixbufScaleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pixbufScaleSimpleFunction *gi.Function
var pixbufScaleSimpleFunction_Once sync.Once

func pixbufScaleSimpleFunction_Set() error {
	var err error
	pixbufScaleSimpleFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufScaleSimpleFunction, err = pixbufObject.InvokerNew("scale_simple")
	})
	return err
}

// ScaleSimple is a representation of the C type gdk_pixbuf_scale_simple.
func (recv *Pixbuf) ScaleSimple(destWidth int32, destHeight int32, interpType InterpType) *Pixbuf {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(destWidth)
	inArgs[2].SetInt32(destHeight)
	inArgs[3].SetInt32(int32(interpType))

	var ret gi.Argument

	err := pixbufScaleSimpleFunction_Set()
	if err == nil {
		ret = pixbufScaleSimpleFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufSetOptionFunction *gi.Function
var pixbufSetOptionFunction_Once sync.Once

func pixbufSetOptionFunction_Set() error {
	var err error
	pixbufSetOptionFunction_Once.Do(func() {
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufSetOptionFunction, err = pixbufObject.InvokerNew("set_option")
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
		err = pixbufObject_Set()
		if err != nil {
			return
		}
		pixbufUnrefFunction, err = pixbufObject.InvokerNew("unref")
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

var pixbufAnimationObject *gi.Object
var pixbufAnimationObject_Once sync.Once

func pixbufAnimationObject_Set() error {
	var err error
	pixbufAnimationObject_Once.Do(func() {
		pixbufAnimationObject, err = gi.ObjectNew("GdkPixbuf", "PixbufAnimation")
	})
	return err
}

type PixbufAnimation struct {
	native unsafe.Pointer
}

func PixbufAnimationNewFromNative(native unsafe.Pointer) *PixbufAnimation {
	instance := &PixbufAnimation{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *PixbufAnimation) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

/*
CastToPixbufAnimation down casts any arbitrary Object to PixbufAnimation.
Exercise care, as this is a potentially dangerous function
if the Object is not a PixbufAnimation.
*/
func (recv *PixbufAnimation) CastToPixbufAnimation(object *gobject.Object) *PixbufAnimation {
	return PixbufAnimationNewFromNative(object.Native())
}

func (recv *PixbufAnimation) Native() unsafe.Pointer {
	return recv.native
}

var pixbufAnimationNewFromFileFunction *gi.Function
var pixbufAnimationNewFromFileFunction_Once sync.Once

func pixbufAnimationNewFromFileFunction_Set() error {
	var err error
	pixbufAnimationNewFromFileFunction_Once.Do(func() {
		err = pixbufAnimationObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationNewFromFileFunction, err = pixbufAnimationObject.InvokerNew("new_from_file")
	})
	return err
}

// PixbufAnimationNewFromFile is a representation of the C type gdk_pixbuf_animation_new_from_file.
func PixbufAnimationNewFromFile(filename string) *PixbufAnimation {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(filename)

	var ret gi.Argument

	err := pixbufAnimationNewFromFileFunction_Set()
	if err == nil {
		ret = pixbufAnimationNewFromFileFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufAnimationNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var pixbufAnimationNewFromResourceFunction *gi.Function
var pixbufAnimationNewFromResourceFunction_Once sync.Once

func pixbufAnimationNewFromResourceFunction_Set() error {
	var err error
	pixbufAnimationNewFromResourceFunction_Once.Do(func() {
		err = pixbufAnimationObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationNewFromResourceFunction, err = pixbufAnimationObject.InvokerNew("new_from_resource")
	})
	return err
}

// PixbufAnimationNewFromResource is a representation of the C type gdk_pixbuf_animation_new_from_resource.
func PixbufAnimationNewFromResource(resourcePath string) *PixbufAnimation {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(resourcePath)

	var ret gi.Argument

	err := pixbufAnimationNewFromResourceFunction_Set()
	if err == nil {
		ret = pixbufAnimationNewFromResourceFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufAnimationNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_animation_new_from_stream' : parameter 'stream' of type 'Gio.InputStream' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_animation_new_from_stream_finish' : parameter 'async_result' of type 'Gio.AsyncResult' not supported

var pixbufAnimationGetHeightFunction *gi.Function
var pixbufAnimationGetHeightFunction_Once sync.Once

func pixbufAnimationGetHeightFunction_Set() error {
	var err error
	pixbufAnimationGetHeightFunction_Once.Do(func() {
		err = pixbufAnimationObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationGetHeightFunction, err = pixbufAnimationObject.InvokerNew("get_height")
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

var pixbufAnimationGetStaticImageFunction *gi.Function
var pixbufAnimationGetStaticImageFunction_Once sync.Once

func pixbufAnimationGetStaticImageFunction_Set() error {
	var err error
	pixbufAnimationGetStaticImageFunction_Once.Do(func() {
		err = pixbufAnimationObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationGetStaticImageFunction, err = pixbufAnimationObject.InvokerNew("get_static_image")
	})
	return err
}

// GetStaticImage is a representation of the C type gdk_pixbuf_animation_get_static_image.
func (recv *PixbufAnimation) GetStaticImage() *Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufAnimationGetStaticImageFunction_Set()
	if err == nil {
		ret = pixbufAnimationGetStaticImageFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufAnimationGetWidthFunction *gi.Function
var pixbufAnimationGetWidthFunction_Once sync.Once

func pixbufAnimationGetWidthFunction_Set() error {
	var err error
	pixbufAnimationGetWidthFunction_Once.Do(func() {
		err = pixbufAnimationObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationGetWidthFunction, err = pixbufAnimationObject.InvokerNew("get_width")
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
		err = pixbufAnimationObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationIsStaticImageFunction, err = pixbufAnimationObject.InvokerNew("is_static_image")
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

var pixbufAnimationRefFunction *gi.Function
var pixbufAnimationRefFunction_Once sync.Once

func pixbufAnimationRefFunction_Set() error {
	var err error
	pixbufAnimationRefFunction_Once.Do(func() {
		err = pixbufAnimationObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationRefFunction, err = pixbufAnimationObject.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gdk_pixbuf_animation_ref.
func (recv *PixbufAnimation) Ref() *PixbufAnimation {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufAnimationRefFunction_Set()
	if err == nil {
		ret = pixbufAnimationRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufAnimationNewFromNative(ret.Pointer())

	return retGo
}

var pixbufAnimationUnrefFunction *gi.Function
var pixbufAnimationUnrefFunction_Once sync.Once

func pixbufAnimationUnrefFunction_Set() error {
	var err error
	pixbufAnimationUnrefFunction_Once.Do(func() {
		err = pixbufAnimationObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationUnrefFunction, err = pixbufAnimationObject.InvokerNew("unref")
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

var pixbufAnimationIterObject *gi.Object
var pixbufAnimationIterObject_Once sync.Once

func pixbufAnimationIterObject_Set() error {
	var err error
	pixbufAnimationIterObject_Once.Do(func() {
		pixbufAnimationIterObject, err = gi.ObjectNew("GdkPixbuf", "PixbufAnimationIter")
	})
	return err
}

type PixbufAnimationIter struct {
	native unsafe.Pointer
}

func PixbufAnimationIterNewFromNative(native unsafe.Pointer) *PixbufAnimationIter {
	instance := &PixbufAnimationIter{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *PixbufAnimationIter) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

/*
CastToPixbufAnimationIter down casts any arbitrary Object to PixbufAnimationIter.
Exercise care, as this is a potentially dangerous function
if the Object is not a PixbufAnimationIter.
*/
func (recv *PixbufAnimationIter) CastToPixbufAnimationIter(object *gobject.Object) *PixbufAnimationIter {
	return PixbufAnimationIterNewFromNative(object.Native())
}

func (recv *PixbufAnimationIter) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'gdk_pixbuf_animation_iter_advance' : parameter 'current_time' of type 'GLib.TimeVal' not supported

var pixbufAnimationIterGetDelayTimeFunction *gi.Function
var pixbufAnimationIterGetDelayTimeFunction_Once sync.Once

func pixbufAnimationIterGetDelayTimeFunction_Set() error {
	var err error
	pixbufAnimationIterGetDelayTimeFunction_Once.Do(func() {
		err = pixbufAnimationIterObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationIterGetDelayTimeFunction, err = pixbufAnimationIterObject.InvokerNew("get_delay_time")
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

var pixbufAnimationIterGetPixbufFunction *gi.Function
var pixbufAnimationIterGetPixbufFunction_Once sync.Once

func pixbufAnimationIterGetPixbufFunction_Set() error {
	var err error
	pixbufAnimationIterGetPixbufFunction_Once.Do(func() {
		err = pixbufAnimationIterObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationIterGetPixbufFunction, err = pixbufAnimationIterObject.InvokerNew("get_pixbuf")
	})
	return err
}

// GetPixbuf is a representation of the C type gdk_pixbuf_animation_iter_get_pixbuf.
func (recv *PixbufAnimationIter) GetPixbuf() *Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufAnimationIterGetPixbufFunction_Set()
	if err == nil {
		ret = pixbufAnimationIterGetPixbufFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufAnimationIterOnCurrentlyLoadingFrameFunction *gi.Function
var pixbufAnimationIterOnCurrentlyLoadingFrameFunction_Once sync.Once

func pixbufAnimationIterOnCurrentlyLoadingFrameFunction_Set() error {
	var err error
	pixbufAnimationIterOnCurrentlyLoadingFrameFunction_Once.Do(func() {
		err = pixbufAnimationIterObject_Set()
		if err != nil {
			return
		}
		pixbufAnimationIterOnCurrentlyLoadingFrameFunction, err = pixbufAnimationIterObject.InvokerNew("on_currently_loading_frame")
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

var pixbufLoaderObject *gi.Object
var pixbufLoaderObject_Once sync.Once

func pixbufLoaderObject_Set() error {
	var err error
	pixbufLoaderObject_Once.Do(func() {
		pixbufLoaderObject, err = gi.ObjectNew("GdkPixbuf", "PixbufLoader")
	})
	return err
}

type PixbufLoader struct {
	native unsafe.Pointer
}

func PixbufLoaderNewFromNative(native unsafe.Pointer) *PixbufLoader {
	instance := &PixbufLoader{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// Object upcasts to *Object
func (recv *PixbufLoader) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

/*
CastToPixbufLoader down casts any arbitrary Object to PixbufLoader.
Exercise care, as this is a potentially dangerous function
if the Object is not a PixbufLoader.
*/
func (recv *PixbufLoader) CastToPixbufLoader(object *gobject.Object) *PixbufLoader {
	return PixbufLoaderNewFromNative(object.Native())
}

func (recv *PixbufLoader) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

var pixbufLoaderNewFunction *gi.Function
var pixbufLoaderNewFunction_Once sync.Once

func pixbufLoaderNewFunction_Set() error {
	var err error
	pixbufLoaderNewFunction_Once.Do(func() {
		err = pixbufLoaderObject_Set()
		if err != nil {
			return
		}
		pixbufLoaderNewFunction, err = pixbufLoaderObject.InvokerNew("new")
	})
	return err
}

// PixbufLoaderNew is a representation of the C type gdk_pixbuf_loader_new.
func PixbufLoaderNew() *PixbufLoader {

	var ret gi.Argument

	err := pixbufLoaderNewFunction_Set()
	if err == nil {
		ret = pixbufLoaderNewFunction.Invoke(nil, nil)
	}

	retGo := PixbufLoaderNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var pixbufLoaderNewWithMimeTypeFunction *gi.Function
var pixbufLoaderNewWithMimeTypeFunction_Once sync.Once

func pixbufLoaderNewWithMimeTypeFunction_Set() error {
	var err error
	pixbufLoaderNewWithMimeTypeFunction_Once.Do(func() {
		err = pixbufLoaderObject_Set()
		if err != nil {
			return
		}
		pixbufLoaderNewWithMimeTypeFunction, err = pixbufLoaderObject.InvokerNew("new_with_mime_type")
	})
	return err
}

// PixbufLoaderNewWithMimeType is a representation of the C type gdk_pixbuf_loader_new_with_mime_type.
func PixbufLoaderNewWithMimeType(mimeType string) *PixbufLoader {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(mimeType)

	var ret gi.Argument

	err := pixbufLoaderNewWithMimeTypeFunction_Set()
	if err == nil {
		ret = pixbufLoaderNewWithMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufLoaderNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var pixbufLoaderNewWithTypeFunction *gi.Function
var pixbufLoaderNewWithTypeFunction_Once sync.Once

func pixbufLoaderNewWithTypeFunction_Set() error {
	var err error
	pixbufLoaderNewWithTypeFunction_Once.Do(func() {
		err = pixbufLoaderObject_Set()
		if err != nil {
			return
		}
		pixbufLoaderNewWithTypeFunction, err = pixbufLoaderObject.InvokerNew("new_with_type")
	})
	return err
}

// PixbufLoaderNewWithType is a representation of the C type gdk_pixbuf_loader_new_with_type.
func PixbufLoaderNewWithType(imageType string) *PixbufLoader {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(imageType)

	var ret gi.Argument

	err := pixbufLoaderNewWithTypeFunction_Set()
	if err == nil {
		ret = pixbufLoaderNewWithTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufLoaderNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var pixbufLoaderCloseFunction *gi.Function
var pixbufLoaderCloseFunction_Once sync.Once

func pixbufLoaderCloseFunction_Set() error {
	var err error
	pixbufLoaderCloseFunction_Once.Do(func() {
		err = pixbufLoaderObject_Set()
		if err != nil {
			return
		}
		pixbufLoaderCloseFunction, err = pixbufLoaderObject.InvokerNew("close")
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

var pixbufLoaderGetAnimationFunction *gi.Function
var pixbufLoaderGetAnimationFunction_Once sync.Once

func pixbufLoaderGetAnimationFunction_Set() error {
	var err error
	pixbufLoaderGetAnimationFunction_Once.Do(func() {
		err = pixbufLoaderObject_Set()
		if err != nil {
			return
		}
		pixbufLoaderGetAnimationFunction, err = pixbufLoaderObject.InvokerNew("get_animation")
	})
	return err
}

// GetAnimation is a representation of the C type gdk_pixbuf_loader_get_animation.
func (recv *PixbufLoader) GetAnimation() *PixbufAnimation {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufLoaderGetAnimationFunction_Set()
	if err == nil {
		ret = pixbufLoaderGetAnimationFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufAnimationNewFromNative(ret.Pointer())

	return retGo
}

var pixbufLoaderGetFormatFunction *gi.Function
var pixbufLoaderGetFormatFunction_Once sync.Once

func pixbufLoaderGetFormatFunction_Set() error {
	var err error
	pixbufLoaderGetFormatFunction_Once.Do(func() {
		err = pixbufLoaderObject_Set()
		if err != nil {
			return
		}
		pixbufLoaderGetFormatFunction, err = pixbufLoaderObject.InvokerNew("get_format")
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

	retGo := PixbufFormatNewFromNative(ret.Pointer())

	return retGo
}

var pixbufLoaderGetPixbufFunction *gi.Function
var pixbufLoaderGetPixbufFunction_Once sync.Once

func pixbufLoaderGetPixbufFunction_Set() error {
	var err error
	pixbufLoaderGetPixbufFunction_Once.Do(func() {
		err = pixbufLoaderObject_Set()
		if err != nil {
			return
		}
		pixbufLoaderGetPixbufFunction, err = pixbufLoaderObject.InvokerNew("get_pixbuf")
	})
	return err
}

// GetPixbuf is a representation of the C type gdk_pixbuf_loader_get_pixbuf.
func (recv *PixbufLoader) GetPixbuf() *Pixbuf {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufLoaderGetPixbufFunction_Set()
	if err == nil {
		ret = pixbufLoaderGetPixbufFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufNewFromNative(ret.Pointer())

	return retGo
}

var pixbufLoaderSetSizeFunction *gi.Function
var pixbufLoaderSetSizeFunction_Once sync.Once

func pixbufLoaderSetSizeFunction_Set() error {
	var err error
	pixbufLoaderSetSizeFunction_Once.Do(func() {
		err = pixbufLoaderObject_Set()
		if err != nil {
			return
		}
		pixbufLoaderSetSizeFunction, err = pixbufLoaderObject.InvokerNew("set_size")
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

var pixbufSimpleAnimObject *gi.Object
var pixbufSimpleAnimObject_Once sync.Once

func pixbufSimpleAnimObject_Set() error {
	var err error
	pixbufSimpleAnimObject_Once.Do(func() {
		pixbufSimpleAnimObject, err = gi.ObjectNew("GdkPixbuf", "PixbufSimpleAnim")
	})
	return err
}

type PixbufSimpleAnim struct {
	native unsafe.Pointer
}

func PixbufSimpleAnimNewFromNative(native unsafe.Pointer) *PixbufSimpleAnim {
	instance := &PixbufSimpleAnim{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// PixbufAnimation upcasts to *PixbufAnimation
func (recv *PixbufSimpleAnim) PixbufAnimation() *PixbufAnimation {
	return PixbufAnimationNewFromNative(recv.native)
}

// Object upcasts to *Object
func (recv *PixbufSimpleAnim) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

/*
CastToPixbufSimpleAnim down casts any arbitrary Object to PixbufSimpleAnim.
Exercise care, as this is a potentially dangerous function
if the Object is not a PixbufSimpleAnim.
*/
func (recv *PixbufSimpleAnim) CastToPixbufSimpleAnim(object *gobject.Object) *PixbufSimpleAnim {
	return PixbufSimpleAnimNewFromNative(object.Native())
}

func (recv *PixbufSimpleAnim) Native() unsafe.Pointer {
	return recv.native
}

var pixbufSimpleAnimNewFunction *gi.Function
var pixbufSimpleAnimNewFunction_Once sync.Once

func pixbufSimpleAnimNewFunction_Set() error {
	var err error
	pixbufSimpleAnimNewFunction_Once.Do(func() {
		err = pixbufSimpleAnimObject_Set()
		if err != nil {
			return
		}
		pixbufSimpleAnimNewFunction, err = pixbufSimpleAnimObject.InvokerNew("new")
	})
	return err
}

// PixbufSimpleAnimNew is a representation of the C type gdk_pixbuf_simple_anim_new.
func PixbufSimpleAnimNew(width int32, height int32, rate float32) *PixbufSimpleAnim {
	var inArgs [3]gi.Argument
	inArgs[0].SetInt32(width)
	inArgs[1].SetInt32(height)
	inArgs[2].SetFloat32(rate)

	var ret gi.Argument

	err := pixbufSimpleAnimNewFunction_Set()
	if err == nil {
		ret = pixbufSimpleAnimNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufSimpleAnimNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var pixbufSimpleAnimAddFrameFunction *gi.Function
var pixbufSimpleAnimAddFrameFunction_Once sync.Once

func pixbufSimpleAnimAddFrameFunction_Set() error {
	var err error
	pixbufSimpleAnimAddFrameFunction_Once.Do(func() {
		err = pixbufSimpleAnimObject_Set()
		if err != nil {
			return
		}
		pixbufSimpleAnimAddFrameFunction, err = pixbufSimpleAnimObject.InvokerNew("add_frame")
	})
	return err
}

// AddFrame is a representation of the C type gdk_pixbuf_simple_anim_add_frame.
func (recv *PixbufSimpleAnim) AddFrame(pixbuf *Pixbuf) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(pixbuf.native)

	err := pixbufSimpleAnimAddFrameFunction_Set()
	if err == nil {
		pixbufSimpleAnimAddFrameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var pixbufSimpleAnimGetLoopFunction *gi.Function
var pixbufSimpleAnimGetLoopFunction_Once sync.Once

func pixbufSimpleAnimGetLoopFunction_Set() error {
	var err error
	pixbufSimpleAnimGetLoopFunction_Once.Do(func() {
		err = pixbufSimpleAnimObject_Set()
		if err != nil {
			return
		}
		pixbufSimpleAnimGetLoopFunction, err = pixbufSimpleAnimObject.InvokerNew("get_loop")
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
		err = pixbufSimpleAnimObject_Set()
		if err != nil {
			return
		}
		pixbufSimpleAnimSetLoopFunction, err = pixbufSimpleAnimObject.InvokerNew("set_loop")
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

var pixbufSimpleAnimIterObject *gi.Object
var pixbufSimpleAnimIterObject_Once sync.Once

func pixbufSimpleAnimIterObject_Set() error {
	var err error
	pixbufSimpleAnimIterObject_Once.Do(func() {
		pixbufSimpleAnimIterObject, err = gi.ObjectNew("GdkPixbuf", "PixbufSimpleAnimIter")
	})
	return err
}

type PixbufSimpleAnimIter struct {
	native unsafe.Pointer
}

func PixbufSimpleAnimIterNewFromNative(native unsafe.Pointer) *PixbufSimpleAnimIter {
	instance := &PixbufSimpleAnimIter{native: native}

	object := instance.Object()
	if object.IsFloating() {
		object.RefSink()
	} else {
		object.Ref()
	}

	runtime.SetFinalizer(object, func(o *gobject.Object) {
		o.Unref()
	})

	return instance
}

// PixbufAnimationIter upcasts to *PixbufAnimationIter
func (recv *PixbufSimpleAnimIter) PixbufAnimationIter() *PixbufAnimationIter {
	return PixbufAnimationIterNewFromNative(recv.native)
}

// Object upcasts to *Object
func (recv *PixbufSimpleAnimIter) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

/*
CastToPixbufSimpleAnimIter down casts any arbitrary Object to PixbufSimpleAnimIter.
Exercise care, as this is a potentially dangerous function
if the Object is not a PixbufSimpleAnimIter.
*/
func (recv *PixbufSimpleAnimIter) CastToPixbufSimpleAnimIter(object *gobject.Object) *PixbufSimpleAnimIter {
	return PixbufSimpleAnimIterNewFromNative(object.Native())
}

func (recv *PixbufSimpleAnimIter) Native() unsafe.Pointer {
	return recv.native
}
