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

func pixbufFormatCopyFunction_Set() error {
	var err error
	pixbufFormatCopyFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatCopyFunction, err = pixbufFormatStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gdk_pixbuf_format_copy.
func (recv *PixbufFormat) Copy() (*PixbufFormat, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufFormatCopyFunction_Set()
	if err == nil {
		ret = pixbufFormatCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &PixbufFormat{native: ret.Pointer()}

	return retGo, err
}

var pixbufFormatFreeFunction *gi.Function
var pixbufFormatFreeFunction_Once sync.Once

func pixbufFormatFreeFunction_Set() error {
	var err error
	pixbufFormatFreeFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatFreeFunction, err = pixbufFormatStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gdk_pixbuf_format_free.
func (recv *PixbufFormat) Free() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := pixbufFormatFreeFunction_Set()
	if err == nil {
		pixbufFormatFreeFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var pixbufFormatGetDescriptionFunction *gi.Function
var pixbufFormatGetDescriptionFunction_Once sync.Once

func pixbufFormatGetDescriptionFunction_Set() error {
	var err error
	pixbufFormatGetDescriptionFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatGetDescriptionFunction, err = pixbufFormatStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type gdk_pixbuf_format_get_description.
func (recv *PixbufFormat) GetDescription() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufFormatGetDescriptionFunction_Set()
	if err == nil {
		ret = pixbufFormatGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var pixbufFormatGetExtensionsFunction *gi.Function
var pixbufFormatGetExtensionsFunction_Once sync.Once

func pixbufFormatGetExtensionsFunction_Set() error {
	var err error
	pixbufFormatGetExtensionsFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatGetExtensionsFunction, err = pixbufFormatStruct.InvokerNew("get_extensions")
	})
	return err
}

// GetExtensions is a representation of the C type gdk_pixbuf_format_get_extensions.
func (recv *PixbufFormat) GetExtensions() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := pixbufFormatGetExtensionsFunction_Set()
	if err == nil {
		pixbufFormatGetExtensionsFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var pixbufFormatGetLicenseFunction *gi.Function
var pixbufFormatGetLicenseFunction_Once sync.Once

func pixbufFormatGetLicenseFunction_Set() error {
	var err error
	pixbufFormatGetLicenseFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatGetLicenseFunction, err = pixbufFormatStruct.InvokerNew("get_license")
	})
	return err
}

// GetLicense is a representation of the C type gdk_pixbuf_format_get_license.
func (recv *PixbufFormat) GetLicense() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufFormatGetLicenseFunction_Set()
	if err == nil {
		ret = pixbufFormatGetLicenseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
}

var pixbufFormatGetMimeTypesFunction *gi.Function
var pixbufFormatGetMimeTypesFunction_Once sync.Once

func pixbufFormatGetMimeTypesFunction_Set() error {
	var err error
	pixbufFormatGetMimeTypesFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatGetMimeTypesFunction, err = pixbufFormatStruct.InvokerNew("get_mime_types")
	})
	return err
}

// GetMimeTypes is a representation of the C type gdk_pixbuf_format_get_mime_types.
func (recv *PixbufFormat) GetMimeTypes() error {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := pixbufFormatGetMimeTypesFunction_Set()
	if err == nil {
		pixbufFormatGetMimeTypesFunction.Invoke(inArgs[:], nil)
	}

	return err
}

var pixbufFormatGetNameFunction *gi.Function
var pixbufFormatGetNameFunction_Once sync.Once

func pixbufFormatGetNameFunction_Set() error {
	var err error
	pixbufFormatGetNameFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatGetNameFunction, err = pixbufFormatStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gdk_pixbuf_format_get_name.
func (recv *PixbufFormat) GetName() (string, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufFormatGetNameFunction_Set()
	if err == nil {
		ret = pixbufFormatGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo, err
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
