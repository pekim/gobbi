// Code generated - DO NOT EDIT.

package gdkpixbuf

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var PixbufFormatStruct *gi.Struct
var PixbufFormatStructOnce sync.Once

func PixbufFormatStructSet() {
	PixbufFormatStructOnce.Do(func() {
		PixbufFormatStruct = gi.StructNew("GdkPixbuf", "PixbufFormat")
	})
}

type PixbufFormat struct {
	native uintptr
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

var PixbufLoaderClassStruct *gi.Struct
var PixbufLoaderClassStructOnce sync.Once

func PixbufLoaderClassStructSet() {
	PixbufLoaderClassStructOnce.Do(func() {
		PixbufLoaderClassStruct = gi.StructNew("GdkPixbuf", "PixbufLoaderClass")
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

var PixbufSimpleAnimClassStruct *gi.Struct
var PixbufSimpleAnimClassStructOnce sync.Once

func PixbufSimpleAnimClassStructSet() {
	PixbufSimpleAnimClassStructOnce.Do(func() {
		PixbufSimpleAnimClassStruct = gi.StructNew("GdkPixbuf", "PixbufSimpleAnimClass")
	})
}

type PixbufSimpleAnimClass struct {
	native uintptr
}
