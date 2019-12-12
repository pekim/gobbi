// Code generated - DO NOT EDIT.

package gdkpixbuf

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
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
	native unsafe.Pointer
}

func PixbufFormatNewFromNative(native unsafe.Pointer) *PixbufFormat {
	err := pixbufFormatStruct_Set()
	if err != nil {
		return nil
	}

	instance := &PixbufFormat{native: native}

	return instance
}

/*
CastToPixbufFormat down casts any arbitrary Object to PixbufFormat.
Exercise care, as this is a potentially dangerous function
if the Object is not a PixbufFormat.
*/
func CastToPixbufFormat(object *gobject.Object) *PixbufFormat {
	return PixbufFormatNewFromNative(object.Native())
}

// Equals compares this PixbufFormat with another PixbufFormat, and returns true if they represent the same Object.
func (recv *PixbufFormat) Equals(other *PixbufFormat) bool {
	return other.Native() == recv.Native()
}

func (recv *PixbufFormat) Native() unsafe.Pointer {
	return recv.native
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
func (recv *PixbufFormat) Copy() *PixbufFormat {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pixbufFormatCopyFunction_Set()
	if err == nil {
		ret = pixbufFormatCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := PixbufFormatNewFromNative(ret.Pointer())

	return retGo
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
func (recv *PixbufFormat) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := pixbufFormatFreeFunction_Set()
	if err == nil {
		pixbufFormatFreeFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *PixbufFormat) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pixbufFormatGetDescriptionFunction_Set()
	if err == nil {
		ret = pixbufFormatGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func (recv *PixbufFormat) GetExtensions() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := pixbufFormatGetExtensionsFunction_Set()
	if err == nil {
		pixbufFormatGetExtensionsFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *PixbufFormat) GetLicense() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pixbufFormatGetLicenseFunction_Set()
	if err == nil {
		ret = pixbufFormatGetLicenseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func (recv *PixbufFormat) GetMimeTypes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := pixbufFormatGetMimeTypesFunction_Set()
	if err == nil {
		pixbufFormatGetMimeTypesFunction.Invoke(inArgs[:], nil)
	}

	return
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
func (recv *PixbufFormat) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pixbufFormatGetNameFunction_Set()
	if err == nil {
		ret = pixbufFormatGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
func (recv *PixbufFormat) IsDisabled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pixbufFormatIsDisabledFunction_Set()
	if err == nil {
		ret = pixbufFormatIsDisabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func (recv *PixbufFormat) IsSaveOptionSupported(optionKey string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(optionKey)

	var ret gi.Argument

	err := pixbufFormatIsSaveOptionSupportedFunction_Set()
	if err == nil {
		ret = pixbufFormatIsSaveOptionSupportedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func (recv *PixbufFormat) IsScalable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pixbufFormatIsScalableFunction_Set()
	if err == nil {
		ret = pixbufFormatIsScalableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func (recv *PixbufFormat) IsWritable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := pixbufFormatIsWritableFunction_Set()
	if err == nil {
		ret = pixbufFormatIsWritableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
func (recv *PixbufFormat) SetDisabled(disabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(disabled)

	err := pixbufFormatSetDisabledFunction_Set()
	if err == nil {
		pixbufFormatSetDisabledFunction.Invoke(inArgs[:], nil)
	}

	return
}

// PixbufFormatStruct creates an uninitialised PixbufFormat.
func PixbufFormatStruct() *PixbufFormat {
	err := pixbufFormatStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PixbufFormatNewFromNative(pixbufFormatStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePixbufFormat)
	return structGo
}
func finalizePixbufFormat(obj *PixbufFormat) {
	pixbufFormatStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func PixbufLoaderClassNewFromNative(native unsafe.Pointer) *PixbufLoaderClass {
	err := pixbufLoaderClassStruct_Set()
	if err != nil {
		return nil
	}

	instance := &PixbufLoaderClass{native: native}

	return instance
}

/*
CastToPixbufLoaderClass down casts any arbitrary Object to PixbufLoaderClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a PixbufLoaderClass.
*/
func CastToPixbufLoaderClass(object *gobject.Object) *PixbufLoaderClass {
	return PixbufLoaderClassNewFromNative(object.Native())
}

// Equals compares this PixbufLoaderClass with another PixbufLoaderClass, and returns true if they represent the same Object.
func (recv *PixbufLoaderClass) Equals(other *PixbufLoaderClass) bool {
	return other.Native() == recv.Native()
}

func (recv *PixbufLoaderClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *PixbufLoaderClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(pixbufLoaderClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *PixbufLoaderClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(pixbufLoaderClassStruct, recv.Native(), "parent_class", argValue)
}

// UNSUPPORTED : C value 'size_prepared' : for field getter : missing Type

// UNSUPPORTED : C value 'size_prepared' : for field setter : missing Type

// UNSUPPORTED : C value 'area_prepared' : for field getter : missing Type

// UNSUPPORTED : C value 'area_prepared' : for field setter : missing Type

// UNSUPPORTED : C value 'area_updated' : for field getter : missing Type

// UNSUPPORTED : C value 'area_updated' : for field setter : missing Type

// UNSUPPORTED : C value 'closed' : for field getter : missing Type

// UNSUPPORTED : C value 'closed' : for field setter : missing Type

// PixbufLoaderClassStruct creates an uninitialised PixbufLoaderClass.
func PixbufLoaderClassStruct() *PixbufLoaderClass {
	err := pixbufLoaderClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PixbufLoaderClassNewFromNative(pixbufLoaderClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePixbufLoaderClass)
	return structGo
}
func finalizePixbufLoaderClass(obj *PixbufLoaderClass) {
	pixbufLoaderClassStruct.Free(obj.Native())
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
	native unsafe.Pointer
}

func PixbufSimpleAnimClassNewFromNative(native unsafe.Pointer) *PixbufSimpleAnimClass {
	err := pixbufSimpleAnimClassStruct_Set()
	if err != nil {
		return nil
	}

	instance := &PixbufSimpleAnimClass{native: native}

	return instance
}

/*
CastToPixbufSimpleAnimClass down casts any arbitrary Object to PixbufSimpleAnimClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a PixbufSimpleAnimClass.
*/
func CastToPixbufSimpleAnimClass(object *gobject.Object) *PixbufSimpleAnimClass {
	return PixbufSimpleAnimClassNewFromNative(object.Native())
}

// Equals compares this PixbufSimpleAnimClass with another PixbufSimpleAnimClass, and returns true if they represent the same Object.
func (recv *PixbufSimpleAnimClass) Equals(other *PixbufSimpleAnimClass) bool {
	return other.Native() == recv.Native()
}

func (recv *PixbufSimpleAnimClass) Native() unsafe.Pointer {
	return recv.native
}

// PixbufSimpleAnimClassStruct creates an uninitialised PixbufSimpleAnimClass.
func PixbufSimpleAnimClassStruct() *PixbufSimpleAnimClass {
	err := pixbufSimpleAnimClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PixbufSimpleAnimClassNewFromNative(pixbufSimpleAnimClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePixbufSimpleAnimClass)
	return structGo
}
func finalizePixbufSimpleAnimClass(obj *PixbufSimpleAnimClass) {
	pixbufSimpleAnimClassStruct.Free(obj.Native())
}
