// Code generated - DO NOT EDIT.

package gdkpixbuf

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var pixbufFormatStruct *gi.Struct
var pixbufFormatStructOnce sync.Once

func pixbufFormatStructSet() {
	pixbufFormatStructOnce.Do(func() {
		pixbufFormatStruct = gi.StructNew("GdkPixbuf", "PixbufFormat")
	})
}

type PixbufFormat struct {
	native uintptr
}

var pixbufFormatCopyFunction *gi.Function
var pixbufFormatCopyFunction_Once sync.Once

func pixbufFormatCopyFunction_Set() {
	pixbufFormatCopyFunction_Once.Do(func() {
		pixbufFormatCopyFunction = gi.FunctionInvokerNew("GdkPixbuf", "copy")
	})
}

var copyPixbufFormatInvoker *gi.Function

// Copy is a representation of the C type gdk_pixbuf_format_copy.
func (recv *PixbufFormat) Copy() *PixbufFormat {
	pixbufFormatCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := pixbufFormatCopyFunction.Invoke(inArgs[:], nil)

	retGo := &PixbufFormat{native: ret.Pointer()}

	return retGo
}

var pixbufFormatFreeFunction *gi.Function
var pixbufFormatFreeFunction_Once sync.Once

func pixbufFormatFreeFunction_Set() {
	pixbufFormatFreeFunction_Once.Do(func() {
		pixbufFormatFreeFunction = gi.FunctionInvokerNew("GdkPixbuf", "free")
	})
}

var freePixbufFormatInvoker *gi.Function

// Free is a representation of the C type gdk_pixbuf_format_free.
func (recv *PixbufFormat) Free() {
	pixbufFormatFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	pixbufFormatFreeFunction.Invoke(inArgs[:], nil)

}

var pixbufFormatGetDescriptionFunction *gi.Function
var pixbufFormatGetDescriptionFunction_Once sync.Once

func pixbufFormatGetDescriptionFunction_Set() {
	pixbufFormatGetDescriptionFunction_Once.Do(func() {
		pixbufFormatGetDescriptionFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_description")
	})
}

var getDescriptionPixbufFormatInvoker *gi.Function

// GetDescription is a representation of the C type gdk_pixbuf_format_get_description.
func (recv *PixbufFormat) GetDescription() string {
	pixbufFormatGetDescriptionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := pixbufFormatGetDescriptionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var pixbufFormatGetExtensionsFunction *gi.Function
var pixbufFormatGetExtensionsFunction_Once sync.Once

func pixbufFormatGetExtensionsFunction_Set() {
	pixbufFormatGetExtensionsFunction_Once.Do(func() {
		pixbufFormatGetExtensionsFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_extensions")
	})
}

var getExtensionsPixbufFormatInvoker *gi.Function

// GetExtensions is a representation of the C type gdk_pixbuf_format_get_extensions.
func (recv *PixbufFormat) GetExtensions() {
	pixbufFormatGetExtensionsFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	pixbufFormatGetExtensionsFunction.Invoke(inArgs[:], nil)

}

var pixbufFormatGetLicenseFunction *gi.Function
var pixbufFormatGetLicenseFunction_Once sync.Once

func pixbufFormatGetLicenseFunction_Set() {
	pixbufFormatGetLicenseFunction_Once.Do(func() {
		pixbufFormatGetLicenseFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_license")
	})
}

var getLicensePixbufFormatInvoker *gi.Function

// GetLicense is a representation of the C type gdk_pixbuf_format_get_license.
func (recv *PixbufFormat) GetLicense() string {
	pixbufFormatGetLicenseFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := pixbufFormatGetLicenseFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var pixbufFormatGetMimeTypesFunction *gi.Function
var pixbufFormatGetMimeTypesFunction_Once sync.Once

func pixbufFormatGetMimeTypesFunction_Set() {
	pixbufFormatGetMimeTypesFunction_Once.Do(func() {
		pixbufFormatGetMimeTypesFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_mime_types")
	})
}

var getMimeTypesPixbufFormatInvoker *gi.Function

// GetMimeTypes is a representation of the C type gdk_pixbuf_format_get_mime_types.
func (recv *PixbufFormat) GetMimeTypes() {
	pixbufFormatGetMimeTypesFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	pixbufFormatGetMimeTypesFunction.Invoke(inArgs[:], nil)

}

var pixbufFormatGetNameFunction *gi.Function
var pixbufFormatGetNameFunction_Once sync.Once

func pixbufFormatGetNameFunction_Set() {
	pixbufFormatGetNameFunction_Once.Do(func() {
		pixbufFormatGetNameFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_name")
	})
}

var getNamePixbufFormatInvoker *gi.Function

// GetName is a representation of the C type gdk_pixbuf_format_get_name.
func (recv *PixbufFormat) GetName() string {
	pixbufFormatGetNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := pixbufFormatGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_format_is_disabled' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_format_is_save_option_supported' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_format_is_scalable' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_format_is_writable' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_format_set_disabled' : parameter 'disabled' of type 'gboolean' not supported

var pixbufLoaderClassStruct *gi.Struct
var pixbufLoaderClassStructOnce sync.Once

func pixbufLoaderClassStructSet() {
	pixbufLoaderClassStructOnce.Do(func() {
		pixbufLoaderClassStruct = gi.StructNew("GdkPixbuf", "PixbufLoaderClass")
	})
}

type PixbufLoaderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'size_prepared' : missing Type
	// UNSUPPORTED : C value 'area_prepared' : missing Type
	// UNSUPPORTED : C value 'area_updated' : missing Type
	// UNSUPPORTED : C value 'closed' : missing Type
}

var pixbufSimpleAnimClassStruct *gi.Struct
var pixbufSimpleAnimClassStructOnce sync.Once

func pixbufSimpleAnimClassStructSet() {
	pixbufSimpleAnimClassStructOnce.Do(func() {
		pixbufSimpleAnimClassStruct = gi.StructNew("GdkPixbuf", "PixbufSimpleAnimClass")
	})
}

type PixbufSimpleAnimClass struct {
	native uintptr
}
