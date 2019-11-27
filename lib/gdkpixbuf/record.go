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

var pixbufFormatIsDisabledFunction *gi.Function
var pixbufFormatIsDisabledFunction_Once sync.Once

func pixbufFormatIsDisabledFunction_Set() error {
	var err error
	pixbufFormatIsDisabledFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatIsDisabledFunction, err = pixbufFormatStruct.InvokerNew("is_disabled")
	})
	return err
}

// IsDisabled is a representation of the C type gdk_pixbuf_format_is_disabled.
func (recv *PixbufFormat) IsDisabled() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufFormatIsDisabledFunction_Set()
	if err == nil {
		ret = pixbufFormatIsDisabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var pixbufFormatIsSaveOptionSupportedFunction *gi.Function
var pixbufFormatIsSaveOptionSupportedFunction_Once sync.Once

func pixbufFormatIsSaveOptionSupportedFunction_Set() error {
	var err error
	pixbufFormatIsSaveOptionSupportedFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatIsSaveOptionSupportedFunction, err = pixbufFormatStruct.InvokerNew("is_save_option_supported")
	})
	return err
}

// IsSaveOptionSupported is a representation of the C type gdk_pixbuf_format_is_save_option_supported.
func (recv *PixbufFormat) IsSaveOptionSupported(optionKey string) (bool, error) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(optionKey)

	var ret gi.Argument

	err := pixbufFormatIsSaveOptionSupportedFunction_Set()
	if err == nil {
		ret = pixbufFormatIsSaveOptionSupportedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var pixbufFormatIsScalableFunction *gi.Function
var pixbufFormatIsScalableFunction_Once sync.Once

func pixbufFormatIsScalableFunction_Set() error {
	var err error
	pixbufFormatIsScalableFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatIsScalableFunction, err = pixbufFormatStruct.InvokerNew("is_scalable")
	})
	return err
}

// IsScalable is a representation of the C type gdk_pixbuf_format_is_scalable.
func (recv *PixbufFormat) IsScalable() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufFormatIsScalableFunction_Set()
	if err == nil {
		ret = pixbufFormatIsScalableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var pixbufFormatIsWritableFunction *gi.Function
var pixbufFormatIsWritableFunction_Once sync.Once

func pixbufFormatIsWritableFunction_Set() error {
	var err error
	pixbufFormatIsWritableFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatIsWritableFunction, err = pixbufFormatStruct.InvokerNew("is_writable")
	})
	return err
}

// IsWritable is a representation of the C type gdk_pixbuf_format_is_writable.
func (recv *PixbufFormat) IsWritable() (bool, error) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := pixbufFormatIsWritableFunction_Set()
	if err == nil {
		ret = pixbufFormatIsWritableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo, err
}

var pixbufFormatSetDisabledFunction *gi.Function
var pixbufFormatSetDisabledFunction_Once sync.Once

func pixbufFormatSetDisabledFunction_Set() error {
	var err error
	pixbufFormatSetDisabledFunction_Once.Do(func() {
		err = pixbufFormatStruct_Set()
		if err != nil {
			return
		}
		pixbufFormatSetDisabledFunction, err = pixbufFormatStruct.InvokerNew("set_disabled")
	})
	return err
}

// SetDisabled is a representation of the C type gdk_pixbuf_format_set_disabled.
func (recv *PixbufFormat) SetDisabled(disabled bool) error {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(disabled)

	err := pixbufFormatSetDisabledFunction_Set()
	if err == nil {
		pixbufFormatSetDisabledFunction.Invoke(inArgs[:], nil)
	}

	return err
}

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
