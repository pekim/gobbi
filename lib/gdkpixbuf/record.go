// Code generated - DO NOT EDIT.

package gdkpixbuf

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var pixbufFormatStruct *gi.Struct
var pixbufFormatStruct_Once sync.Once

func pixbufFormatStruct_Set() error {
	var err error
	pixbufFormatStruct_Once.Do(func() {
		pixbufFormatStruct, err = gi.StructNew("GdkPixbuf", "PixbufFormat")
	})
	return err
}

type PixbufFormat struct {
	native uintptr
}

var pixbufFormatCopyFunction *gi.Function
var pixbufFormatCopyFunction_Once sync.Once

func pixbufFormatCopyFunction_Set() {
	pixbufFormatCopyFunction_Once.Do(func() {
		pixbufFormatStruct_Set()
		pixbufFormatCopyFunction = pixbufFormatStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type gdk_pixbuf_format_copy.
func (recv *PixbufFormat) Copy() *PixbufFormat {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	pixbufFormatCopyFunction_Set()

	ret = pixbufFormatCopyFunction.Invoke(inArgs[:], nil)

	retGo := &PixbufFormat{native: ret.Pointer()}

	return retGo
}

var pixbufFormatFreeFunction *gi.Function
var pixbufFormatFreeFunction_Once sync.Once

func pixbufFormatFreeFunction_Set() {
	pixbufFormatFreeFunction_Once.Do(func() {
		pixbufFormatStruct_Set()
		pixbufFormatFreeFunction = pixbufFormatStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type gdk_pixbuf_format_free.
func (recv *PixbufFormat) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	pixbufFormatFreeFunction_Set()

	pixbufFormatFreeFunction.Invoke(inArgs[:], nil)

}

var pixbufFormatGetDescriptionFunction *gi.Function
var pixbufFormatGetDescriptionFunction_Once sync.Once

func pixbufFormatGetDescriptionFunction_Set() {
	pixbufFormatGetDescriptionFunction_Once.Do(func() {
		pixbufFormatStruct_Set()
		pixbufFormatGetDescriptionFunction = pixbufFormatStruct.InvokerNew("get_description")
	})
}

// GetDescription is a representation of the C type gdk_pixbuf_format_get_description.
func (recv *PixbufFormat) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	pixbufFormatGetDescriptionFunction_Set()

	ret = pixbufFormatGetDescriptionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var pixbufFormatGetExtensionsFunction *gi.Function
var pixbufFormatGetExtensionsFunction_Once sync.Once

func pixbufFormatGetExtensionsFunction_Set() {
	pixbufFormatGetExtensionsFunction_Once.Do(func() {
		pixbufFormatStruct_Set()
		pixbufFormatGetExtensionsFunction = pixbufFormatStruct.InvokerNew("get_extensions")
	})
}

// GetExtensions is a representation of the C type gdk_pixbuf_format_get_extensions.
func (recv *PixbufFormat) GetExtensions() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	pixbufFormatGetExtensionsFunction_Set()

	pixbufFormatGetExtensionsFunction.Invoke(inArgs[:], nil)

}

var pixbufFormatGetLicenseFunction *gi.Function
var pixbufFormatGetLicenseFunction_Once sync.Once

func pixbufFormatGetLicenseFunction_Set() {
	pixbufFormatGetLicenseFunction_Once.Do(func() {
		pixbufFormatStruct_Set()
		pixbufFormatGetLicenseFunction = pixbufFormatStruct.InvokerNew("get_license")
	})
}

// GetLicense is a representation of the C type gdk_pixbuf_format_get_license.
func (recv *PixbufFormat) GetLicense() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	pixbufFormatGetLicenseFunction_Set()

	ret = pixbufFormatGetLicenseFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var pixbufFormatGetMimeTypesFunction *gi.Function
var pixbufFormatGetMimeTypesFunction_Once sync.Once

func pixbufFormatGetMimeTypesFunction_Set() {
	pixbufFormatGetMimeTypesFunction_Once.Do(func() {
		pixbufFormatStruct_Set()
		pixbufFormatGetMimeTypesFunction = pixbufFormatStruct.InvokerNew("get_mime_types")
	})
}

// GetMimeTypes is a representation of the C type gdk_pixbuf_format_get_mime_types.
func (recv *PixbufFormat) GetMimeTypes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	pixbufFormatGetMimeTypesFunction_Set()

	pixbufFormatGetMimeTypesFunction.Invoke(inArgs[:], nil)

}

var pixbufFormatGetNameFunction *gi.Function
var pixbufFormatGetNameFunction_Once sync.Once

func pixbufFormatGetNameFunction_Set() {
	pixbufFormatGetNameFunction_Once.Do(func() {
		pixbufFormatStruct_Set()
		pixbufFormatGetNameFunction = pixbufFormatStruct.InvokerNew("get_name")
	})
}

// GetName is a representation of the C type gdk_pixbuf_format_get_name.
func (recv *PixbufFormat) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	pixbufFormatGetNameFunction_Set()

	ret = pixbufFormatGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gdk_pixbuf_format_is_disabled' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_format_is_save_option_supported' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_format_is_scalable' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_format_is_writable' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_format_set_disabled' : parameter 'disabled' of type 'gboolean' not supported

var pixbufLoaderClassStruct *gi.Struct
var pixbufLoaderClassStruct_Once sync.Once

func pixbufLoaderClassStruct_Set() error {
	var err error
	pixbufLoaderClassStruct_Once.Do(func() {
		pixbufLoaderClassStruct, err = gi.StructNew("GdkPixbuf", "PixbufLoaderClass")
	})
	return err
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
var pixbufSimpleAnimClassStruct_Once sync.Once

func pixbufSimpleAnimClassStruct_Set() error {
	var err error
	pixbufSimpleAnimClassStruct_Once.Do(func() {
		pixbufSimpleAnimClassStruct, err = gi.StructNew("GdkPixbuf", "PixbufSimpleAnimClass")
	})
	return err
}

type PixbufSimpleAnimClass struct {
	native uintptr
}
