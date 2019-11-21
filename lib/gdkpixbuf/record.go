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
var pixbufFormatCopyFunctionOnce sync.Once

func pixbufFormatCopyFunctionSet() {
	pixbufFormatCopyFunctionOnce.Do(func() {
		pixbufFormatCopyFunction = gi.FunctionInvokerNew("GdkPixbuf", "copy")
	})
}

var copyPixbufFormatInvoker *gi.Function

// Copy is a representation of the C type gdk_pixbuf_format_copy.
func (recv *PixbufFormat) Copy() *PixbufFormat {
	if copyPixbufFormatInvoker == nil {
		copyPixbufFormatInvoker = gi.StructFunctionInvokerNew("GdkPixbuf", "PixbufFormat", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyPixbufFormatInvoker.Invoke(inArgs[:], nil)

	retGo := &PixbufFormat{native: ret.Pointer()}

	return retGo
}

var pixbufFormatFreeFunction *gi.Function
var pixbufFormatFreeFunctionOnce sync.Once

func pixbufFormatFreeFunctionSet() {
	pixbufFormatFreeFunctionOnce.Do(func() {
		pixbufFormatFreeFunction = gi.FunctionInvokerNew("GdkPixbuf", "free")
	})
}

var freePixbufFormatInvoker *gi.Function

// Free is a representation of the C type gdk_pixbuf_format_free.
func (recv *PixbufFormat) Free() {
	if freePixbufFormatInvoker == nil {
		freePixbufFormatInvoker = gi.StructFunctionInvokerNew("GdkPixbuf", "PixbufFormat", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freePixbufFormatInvoker.Invoke(inArgs[:], nil)

}

var pixbufFormatGetDescriptionFunction *gi.Function
var pixbufFormatGetDescriptionFunctionOnce sync.Once

func pixbufFormatGetDescriptionFunctionSet() {
	pixbufFormatGetDescriptionFunctionOnce.Do(func() {
		pixbufFormatGetDescriptionFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_description")
	})
}

var getDescriptionPixbufFormatInvoker *gi.Function

// GetDescription is a representation of the C type gdk_pixbuf_format_get_description.
func (recv *PixbufFormat) GetDescription() string {
	if getDescriptionPixbufFormatInvoker == nil {
		getDescriptionPixbufFormatInvoker = gi.StructFunctionInvokerNew("GdkPixbuf", "PixbufFormat", "get_description")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDescriptionPixbufFormatInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var pixbufFormatGetExtensionsFunction *gi.Function
var pixbufFormatGetExtensionsFunctionOnce sync.Once

func pixbufFormatGetExtensionsFunctionSet() {
	pixbufFormatGetExtensionsFunctionOnce.Do(func() {
		pixbufFormatGetExtensionsFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_extensions")
	})
}

var getExtensionsPixbufFormatInvoker *gi.Function

// GetExtensions is a representation of the C type gdk_pixbuf_format_get_extensions.
func (recv *PixbufFormat) GetExtensions() {
	if getExtensionsPixbufFormatInvoker == nil {
		getExtensionsPixbufFormatInvoker = gi.StructFunctionInvokerNew("GdkPixbuf", "PixbufFormat", "get_extensions")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	getExtensionsPixbufFormatInvoker.Invoke(inArgs[:], nil)

}

var pixbufFormatGetLicenseFunction *gi.Function
var pixbufFormatGetLicenseFunctionOnce sync.Once

func pixbufFormatGetLicenseFunctionSet() {
	pixbufFormatGetLicenseFunctionOnce.Do(func() {
		pixbufFormatGetLicenseFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_license")
	})
}

var getLicensePixbufFormatInvoker *gi.Function

// GetLicense is a representation of the C type gdk_pixbuf_format_get_license.
func (recv *PixbufFormat) GetLicense() string {
	if getLicensePixbufFormatInvoker == nil {
		getLicensePixbufFormatInvoker = gi.StructFunctionInvokerNew("GdkPixbuf", "PixbufFormat", "get_license")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLicensePixbufFormatInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var pixbufFormatGetMimeTypesFunction *gi.Function
var pixbufFormatGetMimeTypesFunctionOnce sync.Once

func pixbufFormatGetMimeTypesFunctionSet() {
	pixbufFormatGetMimeTypesFunctionOnce.Do(func() {
		pixbufFormatGetMimeTypesFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_mime_types")
	})
}

var getMimeTypesPixbufFormatInvoker *gi.Function

// GetMimeTypes is a representation of the C type gdk_pixbuf_format_get_mime_types.
func (recv *PixbufFormat) GetMimeTypes() {
	if getMimeTypesPixbufFormatInvoker == nil {
		getMimeTypesPixbufFormatInvoker = gi.StructFunctionInvokerNew("GdkPixbuf", "PixbufFormat", "get_mime_types")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	getMimeTypesPixbufFormatInvoker.Invoke(inArgs[:], nil)

}

var pixbufFormatGetNameFunction *gi.Function
var pixbufFormatGetNameFunctionOnce sync.Once

func pixbufFormatGetNameFunctionSet() {
	pixbufFormatGetNameFunctionOnce.Do(func() {
		pixbufFormatGetNameFunction = gi.FunctionInvokerNew("GdkPixbuf", "get_name")
	})
}

var getNamePixbufFormatInvoker *gi.Function

// GetName is a representation of the C type gdk_pixbuf_format_get_name.
func (recv *PixbufFormat) GetName() string {
	if getNamePixbufFormatInvoker == nil {
		getNamePixbufFormatInvoker = gi.StructFunctionInvokerNew("GdkPixbuf", "PixbufFormat", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNamePixbufFormatInvoker.Invoke(inArgs[:], nil)

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
