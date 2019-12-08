// Code generated - DO NOT EDIT.

package webkit2

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	javascriptcore "github.com/pekim/gobbi/lib/javascriptcore"
	"runtime"
	"sync"
	"unsafe"
)

var applicationInfoStruct *gi.Struct
var applicationInfoStruct_Once sync.Once

func applicationInfoStruct_Set() error {
	var err error
	applicationInfoStruct_Once.Do(func() {
		applicationInfoStruct, err = gi.StructNew("WebKit2", "ApplicationInfo")
	})
	return err
}

type ApplicationInfo struct {
	native unsafe.Pointer
}

func ApplicationInfoNewFromNative(native unsafe.Pointer) *ApplicationInfo {
	instance := &ApplicationInfo{native: native}

	return instance
}

/*
CastToApplicationInfo down casts any arbitrary Object to ApplicationInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a ApplicationInfo.
*/
func CastToApplicationInfo(object *gobject.Object) *ApplicationInfo {
	return ApplicationInfoNewFromNative(object.Native())
}

func (recv *ApplicationInfo) Native() unsafe.Pointer {
	return recv.native
}

var applicationInfoNewFunction *gi.Function
var applicationInfoNewFunction_Once sync.Once

func applicationInfoNewFunction_Set() error {
	var err error
	applicationInfoNewFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoNewFunction, err = applicationInfoStruct.InvokerNew("new")
	})
	return err
}

// ApplicationInfoNew is a representation of the C type webkit_application_info_new.
func ApplicationInfoNew() *ApplicationInfo {

	var ret gi.Argument

	err := applicationInfoNewFunction_Set()
	if err == nil {
		ret = applicationInfoNewFunction.Invoke(nil, nil)
	}

	retGo := ApplicationInfoNewFromNative(ret.Pointer())

	return retGo
}

var applicationInfoGetNameFunction *gi.Function
var applicationInfoGetNameFunction_Once sync.Once

func applicationInfoGetNameFunction_Set() error {
	var err error
	applicationInfoGetNameFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoGetNameFunction, err = applicationInfoStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type webkit_application_info_get_name.
func (recv *ApplicationInfo) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationInfoGetNameFunction_Set()
	if err == nil {
		ret = applicationInfoGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var applicationInfoGetVersionFunction *gi.Function
var applicationInfoGetVersionFunction_Once sync.Once

func applicationInfoGetVersionFunction_Set() error {
	var err error
	applicationInfoGetVersionFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoGetVersionFunction, err = applicationInfoStruct.InvokerNew("get_version")
	})
	return err
}

// GetVersion is a representation of the C type webkit_application_info_get_version.
func (recv *ApplicationInfo) GetVersion() (uint64, uint64, uint64) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [3]gi.Argument

	err := applicationInfoGetVersionFunction_Set()
	if err == nil {
		applicationInfoGetVersionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()
	out1 := outArgs[1].Uint64()
	out2 := outArgs[2].Uint64()

	return out0, out1, out2
}

var applicationInfoRefFunction *gi.Function
var applicationInfoRefFunction_Once sync.Once

func applicationInfoRefFunction_Set() error {
	var err error
	applicationInfoRefFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoRefFunction, err = applicationInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_application_info_ref.
func (recv *ApplicationInfo) Ref() *ApplicationInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := applicationInfoRefFunction_Set()
	if err == nil {
		ret = applicationInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := ApplicationInfoNewFromNative(ret.Pointer())

	return retGo
}

var applicationInfoSetNameFunction *gi.Function
var applicationInfoSetNameFunction_Once sync.Once

func applicationInfoSetNameFunction_Set() error {
	var err error
	applicationInfoSetNameFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoSetNameFunction, err = applicationInfoStruct.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type webkit_application_info_set_name.
func (recv *ApplicationInfo) SetName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	err := applicationInfoSetNameFunction_Set()
	if err == nil {
		applicationInfoSetNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationInfoSetVersionFunction *gi.Function
var applicationInfoSetVersionFunction_Once sync.Once

func applicationInfoSetVersionFunction_Set() error {
	var err error
	applicationInfoSetVersionFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoSetVersionFunction, err = applicationInfoStruct.InvokerNew("set_version")
	})
	return err
}

// SetVersion is a representation of the C type webkit_application_info_set_version.
func (recv *ApplicationInfo) SetVersion(major uint64, minor uint64, micro uint64) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(major)
	inArgs[2].SetUint64(minor)
	inArgs[3].SetUint64(micro)

	err := applicationInfoSetVersionFunction_Set()
	if err == nil {
		applicationInfoSetVersionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var applicationInfoUnrefFunction *gi.Function
var applicationInfoUnrefFunction_Once sync.Once

func applicationInfoUnrefFunction_Set() error {
	var err error
	applicationInfoUnrefFunction_Once.Do(func() {
		err = applicationInfoStruct_Set()
		if err != nil {
			return
		}
		applicationInfoUnrefFunction, err = applicationInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_application_info_unref.
func (recv *ApplicationInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := applicationInfoUnrefFunction_Set()
	if err == nil {
		applicationInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var authenticationRequestClassStruct *gi.Struct
var authenticationRequestClassStruct_Once sync.Once

func authenticationRequestClassStruct_Set() error {
	var err error
	authenticationRequestClassStruct_Once.Do(func() {
		authenticationRequestClassStruct, err = gi.StructNew("WebKit2", "AuthenticationRequestClass")
	})
	return err
}

type AuthenticationRequestClass struct {
	native unsafe.Pointer
}

func AuthenticationRequestClassNewFromNative(native unsafe.Pointer) *AuthenticationRequestClass {
	instance := &AuthenticationRequestClass{native: native}

	return instance
}

/*
CastToAuthenticationRequestClass down casts any arbitrary Object to AuthenticationRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthenticationRequestClass.
*/
func CastToAuthenticationRequestClass(object *gobject.Object) *AuthenticationRequestClass {
	return AuthenticationRequestClassNewFromNative(object.Native())
}

func (recv *AuthenticationRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// AuthenticationRequestClassStruct creates an uninitialised AuthenticationRequestClass.
func AuthenticationRequestClassStruct() *AuthenticationRequestClass {
	err := authenticationRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AuthenticationRequestClassNewFromNative(authenticationRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAuthenticationRequestClass)
	return structGo
}
func finalizeAuthenticationRequestClass(obj *AuthenticationRequestClass) {
	authenticationRequestClassStruct.Free(obj.Native())
}

var authenticationRequestPrivateStruct *gi.Struct
var authenticationRequestPrivateStruct_Once sync.Once

func authenticationRequestPrivateStruct_Set() error {
	var err error
	authenticationRequestPrivateStruct_Once.Do(func() {
		authenticationRequestPrivateStruct, err = gi.StructNew("WebKit2", "AuthenticationRequestPrivate")
	})
	return err
}

type AuthenticationRequestPrivate struct {
	native unsafe.Pointer
}

func AuthenticationRequestPrivateNewFromNative(native unsafe.Pointer) *AuthenticationRequestPrivate {
	instance := &AuthenticationRequestPrivate{native: native}

	return instance
}

/*
CastToAuthenticationRequestPrivate down casts any arbitrary Object to AuthenticationRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a AuthenticationRequestPrivate.
*/
func CastToAuthenticationRequestPrivate(object *gobject.Object) *AuthenticationRequestPrivate {
	return AuthenticationRequestPrivateNewFromNative(object.Native())
}

func (recv *AuthenticationRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// AuthenticationRequestPrivateStruct creates an uninitialised AuthenticationRequestPrivate.
func AuthenticationRequestPrivateStruct() *AuthenticationRequestPrivate {
	err := authenticationRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AuthenticationRequestPrivateNewFromNative(authenticationRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAuthenticationRequestPrivate)
	return structGo
}
func finalizeAuthenticationRequestPrivate(obj *AuthenticationRequestPrivate) {
	authenticationRequestPrivateStruct.Free(obj.Native())
}

var automationSessionClassStruct *gi.Struct
var automationSessionClassStruct_Once sync.Once

func automationSessionClassStruct_Set() error {
	var err error
	automationSessionClassStruct_Once.Do(func() {
		automationSessionClassStruct, err = gi.StructNew("WebKit2", "AutomationSessionClass")
	})
	return err
}

type AutomationSessionClass struct {
	native unsafe.Pointer
}

func AutomationSessionClassNewFromNative(native unsafe.Pointer) *AutomationSessionClass {
	instance := &AutomationSessionClass{native: native}

	return instance
}

/*
CastToAutomationSessionClass down casts any arbitrary Object to AutomationSessionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a AutomationSessionClass.
*/
func CastToAutomationSessionClass(object *gobject.Object) *AutomationSessionClass {
	return AutomationSessionClassNewFromNative(object.Native())
}

func (recv *AutomationSessionClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// AutomationSessionClassStruct creates an uninitialised AutomationSessionClass.
func AutomationSessionClassStruct() *AutomationSessionClass {
	err := automationSessionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AutomationSessionClassNewFromNative(automationSessionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAutomationSessionClass)
	return structGo
}
func finalizeAutomationSessionClass(obj *AutomationSessionClass) {
	automationSessionClassStruct.Free(obj.Native())
}

var automationSessionPrivateStruct *gi.Struct
var automationSessionPrivateStruct_Once sync.Once

func automationSessionPrivateStruct_Set() error {
	var err error
	automationSessionPrivateStruct_Once.Do(func() {
		automationSessionPrivateStruct, err = gi.StructNew("WebKit2", "AutomationSessionPrivate")
	})
	return err
}

type AutomationSessionPrivate struct {
	native unsafe.Pointer
}

func AutomationSessionPrivateNewFromNative(native unsafe.Pointer) *AutomationSessionPrivate {
	instance := &AutomationSessionPrivate{native: native}

	return instance
}

/*
CastToAutomationSessionPrivate down casts any arbitrary Object to AutomationSessionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a AutomationSessionPrivate.
*/
func CastToAutomationSessionPrivate(object *gobject.Object) *AutomationSessionPrivate {
	return AutomationSessionPrivateNewFromNative(object.Native())
}

func (recv *AutomationSessionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// AutomationSessionPrivateStruct creates an uninitialised AutomationSessionPrivate.
func AutomationSessionPrivateStruct() *AutomationSessionPrivate {
	err := automationSessionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AutomationSessionPrivateNewFromNative(automationSessionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAutomationSessionPrivate)
	return structGo
}
func finalizeAutomationSessionPrivate(obj *AutomationSessionPrivate) {
	automationSessionPrivateStruct.Free(obj.Native())
}

var backForwardListClassStruct *gi.Struct
var backForwardListClassStruct_Once sync.Once

func backForwardListClassStruct_Set() error {
	var err error
	backForwardListClassStruct_Once.Do(func() {
		backForwardListClassStruct, err = gi.StructNew("WebKit2", "BackForwardListClass")
	})
	return err
}

type BackForwardListClass struct {
	native unsafe.Pointer
}

func BackForwardListClassNewFromNative(native unsafe.Pointer) *BackForwardListClass {
	instance := &BackForwardListClass{native: native}

	return instance
}

/*
CastToBackForwardListClass down casts any arbitrary Object to BackForwardListClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a BackForwardListClass.
*/
func CastToBackForwardListClass(object *gobject.Object) *BackForwardListClass {
	return BackForwardListClassNewFromNative(object.Native())
}

func (recv *BackForwardListClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// BackForwardListClassStruct creates an uninitialised BackForwardListClass.
func BackForwardListClassStruct() *BackForwardListClass {
	err := backForwardListClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BackForwardListClassNewFromNative(backForwardListClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBackForwardListClass)
	return structGo
}
func finalizeBackForwardListClass(obj *BackForwardListClass) {
	backForwardListClassStruct.Free(obj.Native())
}

var backForwardListItemClassStruct *gi.Struct
var backForwardListItemClassStruct_Once sync.Once

func backForwardListItemClassStruct_Set() error {
	var err error
	backForwardListItemClassStruct_Once.Do(func() {
		backForwardListItemClassStruct, err = gi.StructNew("WebKit2", "BackForwardListItemClass")
	})
	return err
}

type BackForwardListItemClass struct {
	native unsafe.Pointer
}

func BackForwardListItemClassNewFromNative(native unsafe.Pointer) *BackForwardListItemClass {
	instance := &BackForwardListItemClass{native: native}

	return instance
}

/*
CastToBackForwardListItemClass down casts any arbitrary Object to BackForwardListItemClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a BackForwardListItemClass.
*/
func CastToBackForwardListItemClass(object *gobject.Object) *BackForwardListItemClass {
	return BackForwardListItemClassNewFromNative(object.Native())
}

func (recv *BackForwardListItemClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// BackForwardListItemClassStruct creates an uninitialised BackForwardListItemClass.
func BackForwardListItemClassStruct() *BackForwardListItemClass {
	err := backForwardListItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BackForwardListItemClassNewFromNative(backForwardListItemClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBackForwardListItemClass)
	return structGo
}
func finalizeBackForwardListItemClass(obj *BackForwardListItemClass) {
	backForwardListItemClassStruct.Free(obj.Native())
}

var backForwardListItemPrivateStruct *gi.Struct
var backForwardListItemPrivateStruct_Once sync.Once

func backForwardListItemPrivateStruct_Set() error {
	var err error
	backForwardListItemPrivateStruct_Once.Do(func() {
		backForwardListItemPrivateStruct, err = gi.StructNew("WebKit2", "BackForwardListItemPrivate")
	})
	return err
}

type BackForwardListItemPrivate struct {
	native unsafe.Pointer
}

func BackForwardListItemPrivateNewFromNative(native unsafe.Pointer) *BackForwardListItemPrivate {
	instance := &BackForwardListItemPrivate{native: native}

	return instance
}

/*
CastToBackForwardListItemPrivate down casts any arbitrary Object to BackForwardListItemPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a BackForwardListItemPrivate.
*/
func CastToBackForwardListItemPrivate(object *gobject.Object) *BackForwardListItemPrivate {
	return BackForwardListItemPrivateNewFromNative(object.Native())
}

func (recv *BackForwardListItemPrivate) Native() unsafe.Pointer {
	return recv.native
}

// BackForwardListItemPrivateStruct creates an uninitialised BackForwardListItemPrivate.
func BackForwardListItemPrivateStruct() *BackForwardListItemPrivate {
	err := backForwardListItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BackForwardListItemPrivateNewFromNative(backForwardListItemPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBackForwardListItemPrivate)
	return structGo
}
func finalizeBackForwardListItemPrivate(obj *BackForwardListItemPrivate) {
	backForwardListItemPrivateStruct.Free(obj.Native())
}

var backForwardListPrivateStruct *gi.Struct
var backForwardListPrivateStruct_Once sync.Once

func backForwardListPrivateStruct_Set() error {
	var err error
	backForwardListPrivateStruct_Once.Do(func() {
		backForwardListPrivateStruct, err = gi.StructNew("WebKit2", "BackForwardListPrivate")
	})
	return err
}

type BackForwardListPrivate struct {
	native unsafe.Pointer
}

func BackForwardListPrivateNewFromNative(native unsafe.Pointer) *BackForwardListPrivate {
	instance := &BackForwardListPrivate{native: native}

	return instance
}

/*
CastToBackForwardListPrivate down casts any arbitrary Object to BackForwardListPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a BackForwardListPrivate.
*/
func CastToBackForwardListPrivate(object *gobject.Object) *BackForwardListPrivate {
	return BackForwardListPrivateNewFromNative(object.Native())
}

func (recv *BackForwardListPrivate) Native() unsafe.Pointer {
	return recv.native
}

// BackForwardListPrivateStruct creates an uninitialised BackForwardListPrivate.
func BackForwardListPrivateStruct() *BackForwardListPrivate {
	err := backForwardListPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := BackForwardListPrivateNewFromNative(backForwardListPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeBackForwardListPrivate)
	return structGo
}
func finalizeBackForwardListPrivate(obj *BackForwardListPrivate) {
	backForwardListPrivateStruct.Free(obj.Native())
}

var colorChooserRequestClassStruct *gi.Struct
var colorChooserRequestClassStruct_Once sync.Once

func colorChooserRequestClassStruct_Set() error {
	var err error
	colorChooserRequestClassStruct_Once.Do(func() {
		colorChooserRequestClassStruct, err = gi.StructNew("WebKit2", "ColorChooserRequestClass")
	})
	return err
}

type ColorChooserRequestClass struct {
	native unsafe.Pointer
}

func ColorChooserRequestClassNewFromNative(native unsafe.Pointer) *ColorChooserRequestClass {
	instance := &ColorChooserRequestClass{native: native}

	return instance
}

/*
CastToColorChooserRequestClass down casts any arbitrary Object to ColorChooserRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ColorChooserRequestClass.
*/
func CastToColorChooserRequestClass(object *gobject.Object) *ColorChooserRequestClass {
	return ColorChooserRequestClassNewFromNative(object.Native())
}

func (recv *ColorChooserRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// ColorChooserRequestClassStruct creates an uninitialised ColorChooserRequestClass.
func ColorChooserRequestClassStruct() *ColorChooserRequestClass {
	err := colorChooserRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ColorChooserRequestClassNewFromNative(colorChooserRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeColorChooserRequestClass)
	return structGo
}
func finalizeColorChooserRequestClass(obj *ColorChooserRequestClass) {
	colorChooserRequestClassStruct.Free(obj.Native())
}

var colorChooserRequestPrivateStruct *gi.Struct
var colorChooserRequestPrivateStruct_Once sync.Once

func colorChooserRequestPrivateStruct_Set() error {
	var err error
	colorChooserRequestPrivateStruct_Once.Do(func() {
		colorChooserRequestPrivateStruct, err = gi.StructNew("WebKit2", "ColorChooserRequestPrivate")
	})
	return err
}

type ColorChooserRequestPrivate struct {
	native unsafe.Pointer
}

func ColorChooserRequestPrivateNewFromNative(native unsafe.Pointer) *ColorChooserRequestPrivate {
	instance := &ColorChooserRequestPrivate{native: native}

	return instance
}

/*
CastToColorChooserRequestPrivate down casts any arbitrary Object to ColorChooserRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ColorChooserRequestPrivate.
*/
func CastToColorChooserRequestPrivate(object *gobject.Object) *ColorChooserRequestPrivate {
	return ColorChooserRequestPrivateNewFromNative(object.Native())
}

func (recv *ColorChooserRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ColorChooserRequestPrivateStruct creates an uninitialised ColorChooserRequestPrivate.
func ColorChooserRequestPrivateStruct() *ColorChooserRequestPrivate {
	err := colorChooserRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ColorChooserRequestPrivateNewFromNative(colorChooserRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeColorChooserRequestPrivate)
	return structGo
}
func finalizeColorChooserRequestPrivate(obj *ColorChooserRequestPrivate) {
	colorChooserRequestPrivateStruct.Free(obj.Native())
}

var contextMenuClassStruct *gi.Struct
var contextMenuClassStruct_Once sync.Once

func contextMenuClassStruct_Set() error {
	var err error
	contextMenuClassStruct_Once.Do(func() {
		contextMenuClassStruct, err = gi.StructNew("WebKit2", "ContextMenuClass")
	})
	return err
}

type ContextMenuClass struct {
	native unsafe.Pointer
}

func ContextMenuClassNewFromNative(native unsafe.Pointer) *ContextMenuClass {
	instance := &ContextMenuClass{native: native}

	return instance
}

/*
CastToContextMenuClass down casts any arbitrary Object to ContextMenuClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContextMenuClass.
*/
func CastToContextMenuClass(object *gobject.Object) *ContextMenuClass {
	return ContextMenuClassNewFromNative(object.Native())
}

func (recv *ContextMenuClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// ContextMenuClassStruct creates an uninitialised ContextMenuClass.
func ContextMenuClassStruct() *ContextMenuClass {
	err := contextMenuClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContextMenuClassNewFromNative(contextMenuClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContextMenuClass)
	return structGo
}
func finalizeContextMenuClass(obj *ContextMenuClass) {
	contextMenuClassStruct.Free(obj.Native())
}

var contextMenuItemClassStruct *gi.Struct
var contextMenuItemClassStruct_Once sync.Once

func contextMenuItemClassStruct_Set() error {
	var err error
	contextMenuItemClassStruct_Once.Do(func() {
		contextMenuItemClassStruct, err = gi.StructNew("WebKit2", "ContextMenuItemClass")
	})
	return err
}

type ContextMenuItemClass struct {
	native unsafe.Pointer
}

func ContextMenuItemClassNewFromNative(native unsafe.Pointer) *ContextMenuItemClass {
	instance := &ContextMenuItemClass{native: native}

	return instance
}

/*
CastToContextMenuItemClass down casts any arbitrary Object to ContextMenuItemClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContextMenuItemClass.
*/
func CastToContextMenuItemClass(object *gobject.Object) *ContextMenuItemClass {
	return ContextMenuItemClassNewFromNative(object.Native())
}

func (recv *ContextMenuItemClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// ContextMenuItemClassStruct creates an uninitialised ContextMenuItemClass.
func ContextMenuItemClassStruct() *ContextMenuItemClass {
	err := contextMenuItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContextMenuItemClassNewFromNative(contextMenuItemClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContextMenuItemClass)
	return structGo
}
func finalizeContextMenuItemClass(obj *ContextMenuItemClass) {
	contextMenuItemClassStruct.Free(obj.Native())
}

var contextMenuItemPrivateStruct *gi.Struct
var contextMenuItemPrivateStruct_Once sync.Once

func contextMenuItemPrivateStruct_Set() error {
	var err error
	contextMenuItemPrivateStruct_Once.Do(func() {
		contextMenuItemPrivateStruct, err = gi.StructNew("WebKit2", "ContextMenuItemPrivate")
	})
	return err
}

type ContextMenuItemPrivate struct {
	native unsafe.Pointer
}

func ContextMenuItemPrivateNewFromNative(native unsafe.Pointer) *ContextMenuItemPrivate {
	instance := &ContextMenuItemPrivate{native: native}

	return instance
}

/*
CastToContextMenuItemPrivate down casts any arbitrary Object to ContextMenuItemPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContextMenuItemPrivate.
*/
func CastToContextMenuItemPrivate(object *gobject.Object) *ContextMenuItemPrivate {
	return ContextMenuItemPrivateNewFromNative(object.Native())
}

func (recv *ContextMenuItemPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ContextMenuItemPrivateStruct creates an uninitialised ContextMenuItemPrivate.
func ContextMenuItemPrivateStruct() *ContextMenuItemPrivate {
	err := contextMenuItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContextMenuItemPrivateNewFromNative(contextMenuItemPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContextMenuItemPrivate)
	return structGo
}
func finalizeContextMenuItemPrivate(obj *ContextMenuItemPrivate) {
	contextMenuItemPrivateStruct.Free(obj.Native())
}

var contextMenuPrivateStruct *gi.Struct
var contextMenuPrivateStruct_Once sync.Once

func contextMenuPrivateStruct_Set() error {
	var err error
	contextMenuPrivateStruct_Once.Do(func() {
		contextMenuPrivateStruct, err = gi.StructNew("WebKit2", "ContextMenuPrivate")
	})
	return err
}

type ContextMenuPrivate struct {
	native unsafe.Pointer
}

func ContextMenuPrivateNewFromNative(native unsafe.Pointer) *ContextMenuPrivate {
	instance := &ContextMenuPrivate{native: native}

	return instance
}

/*
CastToContextMenuPrivate down casts any arbitrary Object to ContextMenuPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ContextMenuPrivate.
*/
func CastToContextMenuPrivate(object *gobject.Object) *ContextMenuPrivate {
	return ContextMenuPrivateNewFromNative(object.Native())
}

func (recv *ContextMenuPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ContextMenuPrivateStruct creates an uninitialised ContextMenuPrivate.
func ContextMenuPrivateStruct() *ContextMenuPrivate {
	err := contextMenuPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContextMenuPrivateNewFromNative(contextMenuPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContextMenuPrivate)
	return structGo
}
func finalizeContextMenuPrivate(obj *ContextMenuPrivate) {
	contextMenuPrivateStruct.Free(obj.Native())
}

var cookieManagerClassStruct *gi.Struct
var cookieManagerClassStruct_Once sync.Once

func cookieManagerClassStruct_Set() error {
	var err error
	cookieManagerClassStruct_Once.Do(func() {
		cookieManagerClassStruct, err = gi.StructNew("WebKit2", "CookieManagerClass")
	})
	return err
}

type CookieManagerClass struct {
	native unsafe.Pointer
}

func CookieManagerClassNewFromNative(native unsafe.Pointer) *CookieManagerClass {
	instance := &CookieManagerClass{native: native}

	return instance
}

/*
CastToCookieManagerClass down casts any arbitrary Object to CookieManagerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a CookieManagerClass.
*/
func CastToCookieManagerClass(object *gobject.Object) *CookieManagerClass {
	return CookieManagerClassNewFromNative(object.Native())
}

func (recv *CookieManagerClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// CookieManagerClassStruct creates an uninitialised CookieManagerClass.
func CookieManagerClassStruct() *CookieManagerClass {
	err := cookieManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CookieManagerClassNewFromNative(cookieManagerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCookieManagerClass)
	return structGo
}
func finalizeCookieManagerClass(obj *CookieManagerClass) {
	cookieManagerClassStruct.Free(obj.Native())
}

var cookieManagerPrivateStruct *gi.Struct
var cookieManagerPrivateStruct_Once sync.Once

func cookieManagerPrivateStruct_Set() error {
	var err error
	cookieManagerPrivateStruct_Once.Do(func() {
		cookieManagerPrivateStruct, err = gi.StructNew("WebKit2", "CookieManagerPrivate")
	})
	return err
}

type CookieManagerPrivate struct {
	native unsafe.Pointer
}

func CookieManagerPrivateNewFromNative(native unsafe.Pointer) *CookieManagerPrivate {
	instance := &CookieManagerPrivate{native: native}

	return instance
}

/*
CastToCookieManagerPrivate down casts any arbitrary Object to CookieManagerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a CookieManagerPrivate.
*/
func CastToCookieManagerPrivate(object *gobject.Object) *CookieManagerPrivate {
	return CookieManagerPrivateNewFromNative(object.Native())
}

func (recv *CookieManagerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// CookieManagerPrivateStruct creates an uninitialised CookieManagerPrivate.
func CookieManagerPrivateStruct() *CookieManagerPrivate {
	err := cookieManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CookieManagerPrivateNewFromNative(cookieManagerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCookieManagerPrivate)
	return structGo
}
func finalizeCookieManagerPrivate(obj *CookieManagerPrivate) {
	cookieManagerPrivateStruct.Free(obj.Native())
}

var credentialStruct *gi.Struct
var credentialStruct_Once sync.Once

func credentialStruct_Set() error {
	var err error
	credentialStruct_Once.Do(func() {
		credentialStruct, err = gi.StructNew("WebKit2", "Credential")
	})
	return err
}

type Credential struct {
	native unsafe.Pointer
}

func CredentialNewFromNative(native unsafe.Pointer) *Credential {
	instance := &Credential{native: native}

	return instance
}

/*
CastToCredential down casts any arbitrary Object to Credential.
Exercise care, as this is a potentially dangerous function
if the Object is not a Credential.
*/
func CastToCredential(object *gobject.Object) *Credential {
	return CredentialNewFromNative(object.Native())
}

func (recv *Credential) Native() unsafe.Pointer {
	return recv.native
}

var credentialNewFunction *gi.Function
var credentialNewFunction_Once sync.Once

func credentialNewFunction_Set() error {
	var err error
	credentialNewFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialNewFunction, err = credentialStruct.InvokerNew("new")
	})
	return err
}

// CredentialNew is a representation of the C type webkit_credential_new.
func CredentialNew(username string, password string, persistence CredentialPersistence) *Credential {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(username)
	inArgs[1].SetString(password)
	inArgs[2].SetInt32(int32(persistence))

	var ret gi.Argument

	err := credentialNewFunction_Set()
	if err == nil {
		ret = credentialNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := CredentialNewFromNative(ret.Pointer())

	return retGo
}

var credentialCopyFunction *gi.Function
var credentialCopyFunction_Once sync.Once

func credentialCopyFunction_Set() error {
	var err error
	credentialCopyFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialCopyFunction, err = credentialStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_credential_copy.
func (recv *Credential) Copy() *Credential {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := credentialCopyFunction_Set()
	if err == nil {
		ret = credentialCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := CredentialNewFromNative(ret.Pointer())

	return retGo
}

var credentialFreeFunction *gi.Function
var credentialFreeFunction_Once sync.Once

func credentialFreeFunction_Set() error {
	var err error
	credentialFreeFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialFreeFunction, err = credentialStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_credential_free.
func (recv *Credential) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := credentialFreeFunction_Set()
	if err == nil {
		credentialFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var credentialGetPasswordFunction *gi.Function
var credentialGetPasswordFunction_Once sync.Once

func credentialGetPasswordFunction_Set() error {
	var err error
	credentialGetPasswordFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialGetPasswordFunction, err = credentialStruct.InvokerNew("get_password")
	})
	return err
}

// GetPassword is a representation of the C type webkit_credential_get_password.
func (recv *Credential) GetPassword() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := credentialGetPasswordFunction_Set()
	if err == nil {
		ret = credentialGetPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var credentialGetPersistenceFunction *gi.Function
var credentialGetPersistenceFunction_Once sync.Once

func credentialGetPersistenceFunction_Set() error {
	var err error
	credentialGetPersistenceFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialGetPersistenceFunction, err = credentialStruct.InvokerNew("get_persistence")
	})
	return err
}

// GetPersistence is a representation of the C type webkit_credential_get_persistence.
func (recv *Credential) GetPersistence() CredentialPersistence {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := credentialGetPersistenceFunction_Set()
	if err == nil {
		ret = credentialGetPersistenceFunction.Invoke(inArgs[:], nil)
	}

	retGo := CredentialPersistence(ret.Int32())

	return retGo
}

var credentialGetUsernameFunction *gi.Function
var credentialGetUsernameFunction_Once sync.Once

func credentialGetUsernameFunction_Set() error {
	var err error
	credentialGetUsernameFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialGetUsernameFunction, err = credentialStruct.InvokerNew("get_username")
	})
	return err
}

// GetUsername is a representation of the C type webkit_credential_get_username.
func (recv *Credential) GetUsername() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := credentialGetUsernameFunction_Set()
	if err == nil {
		ret = credentialGetUsernameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var credentialHasPasswordFunction *gi.Function
var credentialHasPasswordFunction_Once sync.Once

func credentialHasPasswordFunction_Set() error {
	var err error
	credentialHasPasswordFunction_Once.Do(func() {
		err = credentialStruct_Set()
		if err != nil {
			return
		}
		credentialHasPasswordFunction, err = credentialStruct.InvokerNew("has_password")
	})
	return err
}

// HasPassword is a representation of the C type webkit_credential_has_password.
func (recv *Credential) HasPassword() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := credentialHasPasswordFunction_Set()
	if err == nil {
		ret = credentialHasPasswordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var deviceInfoPermissionRequestClassStruct *gi.Struct
var deviceInfoPermissionRequestClassStruct_Once sync.Once

func deviceInfoPermissionRequestClassStruct_Set() error {
	var err error
	deviceInfoPermissionRequestClassStruct_Once.Do(func() {
		deviceInfoPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "DeviceInfoPermissionRequestClass")
	})
	return err
}

type DeviceInfoPermissionRequestClass struct {
	native unsafe.Pointer
}

func DeviceInfoPermissionRequestClassNewFromNative(native unsafe.Pointer) *DeviceInfoPermissionRequestClass {
	instance := &DeviceInfoPermissionRequestClass{native: native}

	return instance
}

/*
CastToDeviceInfoPermissionRequestClass down casts any arbitrary Object to DeviceInfoPermissionRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DeviceInfoPermissionRequestClass.
*/
func CastToDeviceInfoPermissionRequestClass(object *gobject.Object) *DeviceInfoPermissionRequestClass {
	return DeviceInfoPermissionRequestClassNewFromNative(object.Native())
}

func (recv *DeviceInfoPermissionRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// DeviceInfoPermissionRequestClassStruct creates an uninitialised DeviceInfoPermissionRequestClass.
func DeviceInfoPermissionRequestClassStruct() *DeviceInfoPermissionRequestClass {
	err := deviceInfoPermissionRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DeviceInfoPermissionRequestClassNewFromNative(deviceInfoPermissionRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDeviceInfoPermissionRequestClass)
	return structGo
}
func finalizeDeviceInfoPermissionRequestClass(obj *DeviceInfoPermissionRequestClass) {
	deviceInfoPermissionRequestClassStruct.Free(obj.Native())
}

var deviceInfoPermissionRequestPrivateStruct *gi.Struct
var deviceInfoPermissionRequestPrivateStruct_Once sync.Once

func deviceInfoPermissionRequestPrivateStruct_Set() error {
	var err error
	deviceInfoPermissionRequestPrivateStruct_Once.Do(func() {
		deviceInfoPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "DeviceInfoPermissionRequestPrivate")
	})
	return err
}

type DeviceInfoPermissionRequestPrivate struct {
	native unsafe.Pointer
}

func DeviceInfoPermissionRequestPrivateNewFromNative(native unsafe.Pointer) *DeviceInfoPermissionRequestPrivate {
	instance := &DeviceInfoPermissionRequestPrivate{native: native}

	return instance
}

/*
CastToDeviceInfoPermissionRequestPrivate down casts any arbitrary Object to DeviceInfoPermissionRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DeviceInfoPermissionRequestPrivate.
*/
func CastToDeviceInfoPermissionRequestPrivate(object *gobject.Object) *DeviceInfoPermissionRequestPrivate {
	return DeviceInfoPermissionRequestPrivateNewFromNative(object.Native())
}

func (recv *DeviceInfoPermissionRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DeviceInfoPermissionRequestPrivateStruct creates an uninitialised DeviceInfoPermissionRequestPrivate.
func DeviceInfoPermissionRequestPrivateStruct() *DeviceInfoPermissionRequestPrivate {
	err := deviceInfoPermissionRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DeviceInfoPermissionRequestPrivateNewFromNative(deviceInfoPermissionRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDeviceInfoPermissionRequestPrivate)
	return structGo
}
func finalizeDeviceInfoPermissionRequestPrivate(obj *DeviceInfoPermissionRequestPrivate) {
	deviceInfoPermissionRequestPrivateStruct.Free(obj.Native())
}

var downloadClassStruct *gi.Struct
var downloadClassStruct_Once sync.Once

func downloadClassStruct_Set() error {
	var err error
	downloadClassStruct_Once.Do(func() {
		downloadClassStruct, err = gi.StructNew("WebKit2", "DownloadClass")
	})
	return err
}

type DownloadClass struct {
	native unsafe.Pointer
}

func DownloadClassNewFromNative(native unsafe.Pointer) *DownloadClass {
	instance := &DownloadClass{native: native}

	return instance
}

/*
CastToDownloadClass down casts any arbitrary Object to DownloadClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a DownloadClass.
*/
func CastToDownloadClass(object *gobject.Object) *DownloadClass {
	return DownloadClassNewFromNative(object.Native())
}

func (recv *DownloadClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'decide_destination' : for field getter : missing Type

// UNSUPPORTED : C value 'decide_destination' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// DownloadClassStruct creates an uninitialised DownloadClass.
func DownloadClassStruct() *DownloadClass {
	err := downloadClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DownloadClassNewFromNative(downloadClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDownloadClass)
	return structGo
}
func finalizeDownloadClass(obj *DownloadClass) {
	downloadClassStruct.Free(obj.Native())
}

var downloadPrivateStruct *gi.Struct
var downloadPrivateStruct_Once sync.Once

func downloadPrivateStruct_Set() error {
	var err error
	downloadPrivateStruct_Once.Do(func() {
		downloadPrivateStruct, err = gi.StructNew("WebKit2", "DownloadPrivate")
	})
	return err
}

type DownloadPrivate struct {
	native unsafe.Pointer
}

func DownloadPrivateNewFromNative(native unsafe.Pointer) *DownloadPrivate {
	instance := &DownloadPrivate{native: native}

	return instance
}

/*
CastToDownloadPrivate down casts any arbitrary Object to DownloadPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a DownloadPrivate.
*/
func CastToDownloadPrivate(object *gobject.Object) *DownloadPrivate {
	return DownloadPrivateNewFromNative(object.Native())
}

func (recv *DownloadPrivate) Native() unsafe.Pointer {
	return recv.native
}

// DownloadPrivateStruct creates an uninitialised DownloadPrivate.
func DownloadPrivateStruct() *DownloadPrivate {
	err := downloadPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DownloadPrivateNewFromNative(downloadPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDownloadPrivate)
	return structGo
}
func finalizeDownloadPrivate(obj *DownloadPrivate) {
	downloadPrivateStruct.Free(obj.Native())
}

var editorStateClassStruct *gi.Struct
var editorStateClassStruct_Once sync.Once

func editorStateClassStruct_Set() error {
	var err error
	editorStateClassStruct_Once.Do(func() {
		editorStateClassStruct, err = gi.StructNew("WebKit2", "EditorStateClass")
	})
	return err
}

type EditorStateClass struct {
	native unsafe.Pointer
}

func EditorStateClassNewFromNative(native unsafe.Pointer) *EditorStateClass {
	instance := &EditorStateClass{native: native}

	return instance
}

/*
CastToEditorStateClass down casts any arbitrary Object to EditorStateClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a EditorStateClass.
*/
func CastToEditorStateClass(object *gobject.Object) *EditorStateClass {
	return EditorStateClassNewFromNative(object.Native())
}

func (recv *EditorStateClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// EditorStateClassStruct creates an uninitialised EditorStateClass.
func EditorStateClassStruct() *EditorStateClass {
	err := editorStateClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EditorStateClassNewFromNative(editorStateClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEditorStateClass)
	return structGo
}
func finalizeEditorStateClass(obj *EditorStateClass) {
	editorStateClassStruct.Free(obj.Native())
}

var editorStatePrivateStruct *gi.Struct
var editorStatePrivateStruct_Once sync.Once

func editorStatePrivateStruct_Set() error {
	var err error
	editorStatePrivateStruct_Once.Do(func() {
		editorStatePrivateStruct, err = gi.StructNew("WebKit2", "EditorStatePrivate")
	})
	return err
}

type EditorStatePrivate struct {
	native unsafe.Pointer
}

func EditorStatePrivateNewFromNative(native unsafe.Pointer) *EditorStatePrivate {
	instance := &EditorStatePrivate{native: native}

	return instance
}

/*
CastToEditorStatePrivate down casts any arbitrary Object to EditorStatePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a EditorStatePrivate.
*/
func CastToEditorStatePrivate(object *gobject.Object) *EditorStatePrivate {
	return EditorStatePrivateNewFromNative(object.Native())
}

func (recv *EditorStatePrivate) Native() unsafe.Pointer {
	return recv.native
}

// EditorStatePrivateStruct creates an uninitialised EditorStatePrivate.
func EditorStatePrivateStruct() *EditorStatePrivate {
	err := editorStatePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EditorStatePrivateNewFromNative(editorStatePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEditorStatePrivate)
	return structGo
}
func finalizeEditorStatePrivate(obj *EditorStatePrivate) {
	editorStatePrivateStruct.Free(obj.Native())
}

var faviconDatabaseClassStruct *gi.Struct
var faviconDatabaseClassStruct_Once sync.Once

func faviconDatabaseClassStruct_Set() error {
	var err error
	faviconDatabaseClassStruct_Once.Do(func() {
		faviconDatabaseClassStruct, err = gi.StructNew("WebKit2", "FaviconDatabaseClass")
	})
	return err
}

type FaviconDatabaseClass struct {
	native unsafe.Pointer
}

func FaviconDatabaseClassNewFromNative(native unsafe.Pointer) *FaviconDatabaseClass {
	instance := &FaviconDatabaseClass{native: native}

	return instance
}

/*
CastToFaviconDatabaseClass down casts any arbitrary Object to FaviconDatabaseClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FaviconDatabaseClass.
*/
func CastToFaviconDatabaseClass(object *gobject.Object) *FaviconDatabaseClass {
	return FaviconDatabaseClassNewFromNative(object.Native())
}

func (recv *FaviconDatabaseClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// FaviconDatabaseClassStruct creates an uninitialised FaviconDatabaseClass.
func FaviconDatabaseClassStruct() *FaviconDatabaseClass {
	err := faviconDatabaseClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FaviconDatabaseClassNewFromNative(faviconDatabaseClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFaviconDatabaseClass)
	return structGo
}
func finalizeFaviconDatabaseClass(obj *FaviconDatabaseClass) {
	faviconDatabaseClassStruct.Free(obj.Native())
}

var faviconDatabasePrivateStruct *gi.Struct
var faviconDatabasePrivateStruct_Once sync.Once

func faviconDatabasePrivateStruct_Set() error {
	var err error
	faviconDatabasePrivateStruct_Once.Do(func() {
		faviconDatabasePrivateStruct, err = gi.StructNew("WebKit2", "FaviconDatabasePrivate")
	})
	return err
}

type FaviconDatabasePrivate struct {
	native unsafe.Pointer
}

func FaviconDatabasePrivateNewFromNative(native unsafe.Pointer) *FaviconDatabasePrivate {
	instance := &FaviconDatabasePrivate{native: native}

	return instance
}

/*
CastToFaviconDatabasePrivate down casts any arbitrary Object to FaviconDatabasePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FaviconDatabasePrivate.
*/
func CastToFaviconDatabasePrivate(object *gobject.Object) *FaviconDatabasePrivate {
	return FaviconDatabasePrivateNewFromNative(object.Native())
}

func (recv *FaviconDatabasePrivate) Native() unsafe.Pointer {
	return recv.native
}

// FaviconDatabasePrivateStruct creates an uninitialised FaviconDatabasePrivate.
func FaviconDatabasePrivateStruct() *FaviconDatabasePrivate {
	err := faviconDatabasePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FaviconDatabasePrivateNewFromNative(faviconDatabasePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFaviconDatabasePrivate)
	return structGo
}
func finalizeFaviconDatabasePrivate(obj *FaviconDatabasePrivate) {
	faviconDatabasePrivateStruct.Free(obj.Native())
}

var fileChooserRequestClassStruct *gi.Struct
var fileChooserRequestClassStruct_Once sync.Once

func fileChooserRequestClassStruct_Set() error {
	var err error
	fileChooserRequestClassStruct_Once.Do(func() {
		fileChooserRequestClassStruct, err = gi.StructNew("WebKit2", "FileChooserRequestClass")
	})
	return err
}

type FileChooserRequestClass struct {
	native unsafe.Pointer
}

func FileChooserRequestClassNewFromNative(native unsafe.Pointer) *FileChooserRequestClass {
	instance := &FileChooserRequestClass{native: native}

	return instance
}

/*
CastToFileChooserRequestClass down casts any arbitrary Object to FileChooserRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileChooserRequestClass.
*/
func CastToFileChooserRequestClass(object *gobject.Object) *FileChooserRequestClass {
	return FileChooserRequestClassNewFromNative(object.Native())
}

func (recv *FileChooserRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// FileChooserRequestClassStruct creates an uninitialised FileChooserRequestClass.
func FileChooserRequestClassStruct() *FileChooserRequestClass {
	err := fileChooserRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileChooserRequestClassNewFromNative(fileChooserRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileChooserRequestClass)
	return structGo
}
func finalizeFileChooserRequestClass(obj *FileChooserRequestClass) {
	fileChooserRequestClassStruct.Free(obj.Native())
}

var fileChooserRequestPrivateStruct *gi.Struct
var fileChooserRequestPrivateStruct_Once sync.Once

func fileChooserRequestPrivateStruct_Set() error {
	var err error
	fileChooserRequestPrivateStruct_Once.Do(func() {
		fileChooserRequestPrivateStruct, err = gi.StructNew("WebKit2", "FileChooserRequestPrivate")
	})
	return err
}

type FileChooserRequestPrivate struct {
	native unsafe.Pointer
}

func FileChooserRequestPrivateNewFromNative(native unsafe.Pointer) *FileChooserRequestPrivate {
	instance := &FileChooserRequestPrivate{native: native}

	return instance
}

/*
CastToFileChooserRequestPrivate down casts any arbitrary Object to FileChooserRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileChooserRequestPrivate.
*/
func CastToFileChooserRequestPrivate(object *gobject.Object) *FileChooserRequestPrivate {
	return FileChooserRequestPrivateNewFromNative(object.Native())
}

func (recv *FileChooserRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// FileChooserRequestPrivateStruct creates an uninitialised FileChooserRequestPrivate.
func FileChooserRequestPrivateStruct() *FileChooserRequestPrivate {
	err := fileChooserRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FileChooserRequestPrivateNewFromNative(fileChooserRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFileChooserRequestPrivate)
	return structGo
}
func finalizeFileChooserRequestPrivate(obj *FileChooserRequestPrivate) {
	fileChooserRequestPrivateStruct.Free(obj.Native())
}

var findControllerClassStruct *gi.Struct
var findControllerClassStruct_Once sync.Once

func findControllerClassStruct_Set() error {
	var err error
	findControllerClassStruct_Once.Do(func() {
		findControllerClassStruct, err = gi.StructNew("WebKit2", "FindControllerClass")
	})
	return err
}

type FindControllerClass struct {
	native unsafe.Pointer
}

func FindControllerClassNewFromNative(native unsafe.Pointer) *FindControllerClass {
	instance := &FindControllerClass{native: native}

	return instance
}

/*
CastToFindControllerClass down casts any arbitrary Object to FindControllerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FindControllerClass.
*/
func CastToFindControllerClass(object *gobject.Object) *FindControllerClass {
	return FindControllerClassNewFromNative(object.Native())
}

func (recv *FindControllerClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// FindControllerClassStruct creates an uninitialised FindControllerClass.
func FindControllerClassStruct() *FindControllerClass {
	err := findControllerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FindControllerClassNewFromNative(findControllerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFindControllerClass)
	return structGo
}
func finalizeFindControllerClass(obj *FindControllerClass) {
	findControllerClassStruct.Free(obj.Native())
}

var findControllerPrivateStruct *gi.Struct
var findControllerPrivateStruct_Once sync.Once

func findControllerPrivateStruct_Set() error {
	var err error
	findControllerPrivateStruct_Once.Do(func() {
		findControllerPrivateStruct, err = gi.StructNew("WebKit2", "FindControllerPrivate")
	})
	return err
}

type FindControllerPrivate struct {
	native unsafe.Pointer
}

func FindControllerPrivateNewFromNative(native unsafe.Pointer) *FindControllerPrivate {
	instance := &FindControllerPrivate{native: native}

	return instance
}

/*
CastToFindControllerPrivate down casts any arbitrary Object to FindControllerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FindControllerPrivate.
*/
func CastToFindControllerPrivate(object *gobject.Object) *FindControllerPrivate {
	return FindControllerPrivateNewFromNative(object.Native())
}

func (recv *FindControllerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// FindControllerPrivateStruct creates an uninitialised FindControllerPrivate.
func FindControllerPrivateStruct() *FindControllerPrivate {
	err := findControllerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FindControllerPrivateNewFromNative(findControllerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFindControllerPrivate)
	return structGo
}
func finalizeFindControllerPrivate(obj *FindControllerPrivate) {
	findControllerPrivateStruct.Free(obj.Native())
}

var formSubmissionRequestClassStruct *gi.Struct
var formSubmissionRequestClassStruct_Once sync.Once

func formSubmissionRequestClassStruct_Set() error {
	var err error
	formSubmissionRequestClassStruct_Once.Do(func() {
		formSubmissionRequestClassStruct, err = gi.StructNew("WebKit2", "FormSubmissionRequestClass")
	})
	return err
}

type FormSubmissionRequestClass struct {
	native unsafe.Pointer
}

func FormSubmissionRequestClassNewFromNative(native unsafe.Pointer) *FormSubmissionRequestClass {
	instance := &FormSubmissionRequestClass{native: native}

	return instance
}

/*
CastToFormSubmissionRequestClass down casts any arbitrary Object to FormSubmissionRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a FormSubmissionRequestClass.
*/
func CastToFormSubmissionRequestClass(object *gobject.Object) *FormSubmissionRequestClass {
	return FormSubmissionRequestClassNewFromNative(object.Native())
}

func (recv *FormSubmissionRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// FormSubmissionRequestClassStruct creates an uninitialised FormSubmissionRequestClass.
func FormSubmissionRequestClassStruct() *FormSubmissionRequestClass {
	err := formSubmissionRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FormSubmissionRequestClassNewFromNative(formSubmissionRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFormSubmissionRequestClass)
	return structGo
}
func finalizeFormSubmissionRequestClass(obj *FormSubmissionRequestClass) {
	formSubmissionRequestClassStruct.Free(obj.Native())
}

var formSubmissionRequestPrivateStruct *gi.Struct
var formSubmissionRequestPrivateStruct_Once sync.Once

func formSubmissionRequestPrivateStruct_Set() error {
	var err error
	formSubmissionRequestPrivateStruct_Once.Do(func() {
		formSubmissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "FormSubmissionRequestPrivate")
	})
	return err
}

type FormSubmissionRequestPrivate struct {
	native unsafe.Pointer
}

func FormSubmissionRequestPrivateNewFromNative(native unsafe.Pointer) *FormSubmissionRequestPrivate {
	instance := &FormSubmissionRequestPrivate{native: native}

	return instance
}

/*
CastToFormSubmissionRequestPrivate down casts any arbitrary Object to FormSubmissionRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a FormSubmissionRequestPrivate.
*/
func CastToFormSubmissionRequestPrivate(object *gobject.Object) *FormSubmissionRequestPrivate {
	return FormSubmissionRequestPrivateNewFromNative(object.Native())
}

func (recv *FormSubmissionRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// FormSubmissionRequestPrivateStruct creates an uninitialised FormSubmissionRequestPrivate.
func FormSubmissionRequestPrivateStruct() *FormSubmissionRequestPrivate {
	err := formSubmissionRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FormSubmissionRequestPrivateNewFromNative(formSubmissionRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFormSubmissionRequestPrivate)
	return structGo
}
func finalizeFormSubmissionRequestPrivate(obj *FormSubmissionRequestPrivate) {
	formSubmissionRequestPrivateStruct.Free(obj.Native())
}

var geolocationManagerClassStruct *gi.Struct
var geolocationManagerClassStruct_Once sync.Once

func geolocationManagerClassStruct_Set() error {
	var err error
	geolocationManagerClassStruct_Once.Do(func() {
		geolocationManagerClassStruct, err = gi.StructNew("WebKit2", "GeolocationManagerClass")
	})
	return err
}

type GeolocationManagerClass struct {
	native unsafe.Pointer
}

func GeolocationManagerClassNewFromNative(native unsafe.Pointer) *GeolocationManagerClass {
	instance := &GeolocationManagerClass{native: native}

	return instance
}

/*
CastToGeolocationManagerClass down casts any arbitrary Object to GeolocationManagerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a GeolocationManagerClass.
*/
func CastToGeolocationManagerClass(object *gobject.Object) *GeolocationManagerClass {
	return GeolocationManagerClassNewFromNative(object.Native())
}

func (recv *GeolocationManagerClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// GeolocationManagerClassStruct creates an uninitialised GeolocationManagerClass.
func GeolocationManagerClassStruct() *GeolocationManagerClass {
	err := geolocationManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GeolocationManagerClassNewFromNative(geolocationManagerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGeolocationManagerClass)
	return structGo
}
func finalizeGeolocationManagerClass(obj *GeolocationManagerClass) {
	geolocationManagerClassStruct.Free(obj.Native())
}

var geolocationManagerPrivateStruct *gi.Struct
var geolocationManagerPrivateStruct_Once sync.Once

func geolocationManagerPrivateStruct_Set() error {
	var err error
	geolocationManagerPrivateStruct_Once.Do(func() {
		geolocationManagerPrivateStruct, err = gi.StructNew("WebKit2", "GeolocationManagerPrivate")
	})
	return err
}

type GeolocationManagerPrivate struct {
	native unsafe.Pointer
}

func GeolocationManagerPrivateNewFromNative(native unsafe.Pointer) *GeolocationManagerPrivate {
	instance := &GeolocationManagerPrivate{native: native}

	return instance
}

/*
CastToGeolocationManagerPrivate down casts any arbitrary Object to GeolocationManagerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a GeolocationManagerPrivate.
*/
func CastToGeolocationManagerPrivate(object *gobject.Object) *GeolocationManagerPrivate {
	return GeolocationManagerPrivateNewFromNative(object.Native())
}

func (recv *GeolocationManagerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// GeolocationManagerPrivateStruct creates an uninitialised GeolocationManagerPrivate.
func GeolocationManagerPrivateStruct() *GeolocationManagerPrivate {
	err := geolocationManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GeolocationManagerPrivateNewFromNative(geolocationManagerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGeolocationManagerPrivate)
	return structGo
}
func finalizeGeolocationManagerPrivate(obj *GeolocationManagerPrivate) {
	geolocationManagerPrivateStruct.Free(obj.Native())
}

var geolocationPermissionRequestClassStruct *gi.Struct
var geolocationPermissionRequestClassStruct_Once sync.Once

func geolocationPermissionRequestClassStruct_Set() error {
	var err error
	geolocationPermissionRequestClassStruct_Once.Do(func() {
		geolocationPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "GeolocationPermissionRequestClass")
	})
	return err
}

type GeolocationPermissionRequestClass struct {
	native unsafe.Pointer
}

func GeolocationPermissionRequestClassNewFromNative(native unsafe.Pointer) *GeolocationPermissionRequestClass {
	instance := &GeolocationPermissionRequestClass{native: native}

	return instance
}

/*
CastToGeolocationPermissionRequestClass down casts any arbitrary Object to GeolocationPermissionRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a GeolocationPermissionRequestClass.
*/
func CastToGeolocationPermissionRequestClass(object *gobject.Object) *GeolocationPermissionRequestClass {
	return GeolocationPermissionRequestClassNewFromNative(object.Native())
}

func (recv *GeolocationPermissionRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// GeolocationPermissionRequestClassStruct creates an uninitialised GeolocationPermissionRequestClass.
func GeolocationPermissionRequestClassStruct() *GeolocationPermissionRequestClass {
	err := geolocationPermissionRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GeolocationPermissionRequestClassNewFromNative(geolocationPermissionRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGeolocationPermissionRequestClass)
	return structGo
}
func finalizeGeolocationPermissionRequestClass(obj *GeolocationPermissionRequestClass) {
	geolocationPermissionRequestClassStruct.Free(obj.Native())
}

var geolocationPermissionRequestPrivateStruct *gi.Struct
var geolocationPermissionRequestPrivateStruct_Once sync.Once

func geolocationPermissionRequestPrivateStruct_Set() error {
	var err error
	geolocationPermissionRequestPrivateStruct_Once.Do(func() {
		geolocationPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "GeolocationPermissionRequestPrivate")
	})
	return err
}

type GeolocationPermissionRequestPrivate struct {
	native unsafe.Pointer
}

func GeolocationPermissionRequestPrivateNewFromNative(native unsafe.Pointer) *GeolocationPermissionRequestPrivate {
	instance := &GeolocationPermissionRequestPrivate{native: native}

	return instance
}

/*
CastToGeolocationPermissionRequestPrivate down casts any arbitrary Object to GeolocationPermissionRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a GeolocationPermissionRequestPrivate.
*/
func CastToGeolocationPermissionRequestPrivate(object *gobject.Object) *GeolocationPermissionRequestPrivate {
	return GeolocationPermissionRequestPrivateNewFromNative(object.Native())
}

func (recv *GeolocationPermissionRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// GeolocationPermissionRequestPrivateStruct creates an uninitialised GeolocationPermissionRequestPrivate.
func GeolocationPermissionRequestPrivateStruct() *GeolocationPermissionRequestPrivate {
	err := geolocationPermissionRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GeolocationPermissionRequestPrivateNewFromNative(geolocationPermissionRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGeolocationPermissionRequestPrivate)
	return structGo
}
func finalizeGeolocationPermissionRequestPrivate(obj *GeolocationPermissionRequestPrivate) {
	geolocationPermissionRequestPrivateStruct.Free(obj.Native())
}

var geolocationPositionStruct *gi.Struct
var geolocationPositionStruct_Once sync.Once

func geolocationPositionStruct_Set() error {
	var err error
	geolocationPositionStruct_Once.Do(func() {
		geolocationPositionStruct, err = gi.StructNew("WebKit2", "GeolocationPosition")
	})
	return err
}

type GeolocationPosition struct {
	native unsafe.Pointer
}

func GeolocationPositionNewFromNative(native unsafe.Pointer) *GeolocationPosition {
	instance := &GeolocationPosition{native: native}

	return instance
}

/*
CastToGeolocationPosition down casts any arbitrary Object to GeolocationPosition.
Exercise care, as this is a potentially dangerous function
if the Object is not a GeolocationPosition.
*/
func CastToGeolocationPosition(object *gobject.Object) *GeolocationPosition {
	return GeolocationPositionNewFromNative(object.Native())
}

func (recv *GeolocationPosition) Native() unsafe.Pointer {
	return recv.native
}

var geolocationPositionNewFunction *gi.Function
var geolocationPositionNewFunction_Once sync.Once

func geolocationPositionNewFunction_Set() error {
	var err error
	geolocationPositionNewFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionNewFunction, err = geolocationPositionStruct.InvokerNew("new")
	})
	return err
}

// GeolocationPositionNew is a representation of the C type webkit_geolocation_position_new.
func GeolocationPositionNew(latitude float64, longitude float64, accuracy float64) *GeolocationPosition {
	var inArgs [3]gi.Argument
	inArgs[0].SetFloat64(latitude)
	inArgs[1].SetFloat64(longitude)
	inArgs[2].SetFloat64(accuracy)

	var ret gi.Argument

	err := geolocationPositionNewFunction_Set()
	if err == nil {
		ret = geolocationPositionNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := GeolocationPositionNewFromNative(ret.Pointer())

	return retGo
}

var geolocationPositionCopyFunction *gi.Function
var geolocationPositionCopyFunction_Once sync.Once

func geolocationPositionCopyFunction_Set() error {
	var err error
	geolocationPositionCopyFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionCopyFunction, err = geolocationPositionStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_geolocation_position_copy.
func (recv *GeolocationPosition) Copy() *GeolocationPosition {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := geolocationPositionCopyFunction_Set()
	if err == nil {
		ret = geolocationPositionCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := GeolocationPositionNewFromNative(ret.Pointer())

	return retGo
}

var geolocationPositionFreeFunction *gi.Function
var geolocationPositionFreeFunction_Once sync.Once

func geolocationPositionFreeFunction_Set() error {
	var err error
	geolocationPositionFreeFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionFreeFunction, err = geolocationPositionStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_geolocation_position_free.
func (recv *GeolocationPosition) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := geolocationPositionFreeFunction_Set()
	if err == nil {
		geolocationPositionFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var geolocationPositionSetAltitudeFunction *gi.Function
var geolocationPositionSetAltitudeFunction_Once sync.Once

func geolocationPositionSetAltitudeFunction_Set() error {
	var err error
	geolocationPositionSetAltitudeFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionSetAltitudeFunction, err = geolocationPositionStruct.InvokerNew("set_altitude")
	})
	return err
}

// SetAltitude is a representation of the C type webkit_geolocation_position_set_altitude.
func (recv *GeolocationPosition) SetAltitude(altitude float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetFloat64(altitude)

	err := geolocationPositionSetAltitudeFunction_Set()
	if err == nil {
		geolocationPositionSetAltitudeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var geolocationPositionSetAltitudeAccuracyFunction *gi.Function
var geolocationPositionSetAltitudeAccuracyFunction_Once sync.Once

func geolocationPositionSetAltitudeAccuracyFunction_Set() error {
	var err error
	geolocationPositionSetAltitudeAccuracyFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionSetAltitudeAccuracyFunction, err = geolocationPositionStruct.InvokerNew("set_altitude_accuracy")
	})
	return err
}

// SetAltitudeAccuracy is a representation of the C type webkit_geolocation_position_set_altitude_accuracy.
func (recv *GeolocationPosition) SetAltitudeAccuracy(altitudeAccuracy float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetFloat64(altitudeAccuracy)

	err := geolocationPositionSetAltitudeAccuracyFunction_Set()
	if err == nil {
		geolocationPositionSetAltitudeAccuracyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var geolocationPositionSetHeadingFunction *gi.Function
var geolocationPositionSetHeadingFunction_Once sync.Once

func geolocationPositionSetHeadingFunction_Set() error {
	var err error
	geolocationPositionSetHeadingFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionSetHeadingFunction, err = geolocationPositionStruct.InvokerNew("set_heading")
	})
	return err
}

// SetHeading is a representation of the C type webkit_geolocation_position_set_heading.
func (recv *GeolocationPosition) SetHeading(heading float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetFloat64(heading)

	err := geolocationPositionSetHeadingFunction_Set()
	if err == nil {
		geolocationPositionSetHeadingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var geolocationPositionSetSpeedFunction *gi.Function
var geolocationPositionSetSpeedFunction_Once sync.Once

func geolocationPositionSetSpeedFunction_Set() error {
	var err error
	geolocationPositionSetSpeedFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionSetSpeedFunction, err = geolocationPositionStruct.InvokerNew("set_speed")
	})
	return err
}

// SetSpeed is a representation of the C type webkit_geolocation_position_set_speed.
func (recv *GeolocationPosition) SetSpeed(speed float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetFloat64(speed)

	err := geolocationPositionSetSpeedFunction_Set()
	if err == nil {
		geolocationPositionSetSpeedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var geolocationPositionSetTimestampFunction *gi.Function
var geolocationPositionSetTimestampFunction_Once sync.Once

func geolocationPositionSetTimestampFunction_Set() error {
	var err error
	geolocationPositionSetTimestampFunction_Once.Do(func() {
		err = geolocationPositionStruct_Set()
		if err != nil {
			return
		}
		geolocationPositionSetTimestampFunction, err = geolocationPositionStruct.InvokerNew("set_timestamp")
	})
	return err
}

// SetTimestamp is a representation of the C type webkit_geolocation_position_set_timestamp.
func (recv *GeolocationPosition) SetTimestamp(timestamp uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(timestamp)

	err := geolocationPositionSetTimestampFunction_Set()
	if err == nil {
		geolocationPositionSetTimestampFunction.Invoke(inArgs[:], nil)
	}

	return
}

var hitTestResultClassStruct *gi.Struct
var hitTestResultClassStruct_Once sync.Once

func hitTestResultClassStruct_Set() error {
	var err error
	hitTestResultClassStruct_Once.Do(func() {
		hitTestResultClassStruct, err = gi.StructNew("WebKit2", "HitTestResultClass")
	})
	return err
}

type HitTestResultClass struct {
	native unsafe.Pointer
}

func HitTestResultClassNewFromNative(native unsafe.Pointer) *HitTestResultClass {
	instance := &HitTestResultClass{native: native}

	return instance
}

/*
CastToHitTestResultClass down casts any arbitrary Object to HitTestResultClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a HitTestResultClass.
*/
func CastToHitTestResultClass(object *gobject.Object) *HitTestResultClass {
	return HitTestResultClassNewFromNative(object.Native())
}

func (recv *HitTestResultClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// HitTestResultClassStruct creates an uninitialised HitTestResultClass.
func HitTestResultClassStruct() *HitTestResultClass {
	err := hitTestResultClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := HitTestResultClassNewFromNative(hitTestResultClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeHitTestResultClass)
	return structGo
}
func finalizeHitTestResultClass(obj *HitTestResultClass) {
	hitTestResultClassStruct.Free(obj.Native())
}

var hitTestResultPrivateStruct *gi.Struct
var hitTestResultPrivateStruct_Once sync.Once

func hitTestResultPrivateStruct_Set() error {
	var err error
	hitTestResultPrivateStruct_Once.Do(func() {
		hitTestResultPrivateStruct, err = gi.StructNew("WebKit2", "HitTestResultPrivate")
	})
	return err
}

type HitTestResultPrivate struct {
	native unsafe.Pointer
}

func HitTestResultPrivateNewFromNative(native unsafe.Pointer) *HitTestResultPrivate {
	instance := &HitTestResultPrivate{native: native}

	return instance
}

/*
CastToHitTestResultPrivate down casts any arbitrary Object to HitTestResultPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a HitTestResultPrivate.
*/
func CastToHitTestResultPrivate(object *gobject.Object) *HitTestResultPrivate {
	return HitTestResultPrivateNewFromNative(object.Native())
}

func (recv *HitTestResultPrivate) Native() unsafe.Pointer {
	return recv.native
}

// HitTestResultPrivateStruct creates an uninitialised HitTestResultPrivate.
func HitTestResultPrivateStruct() *HitTestResultPrivate {
	err := hitTestResultPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := HitTestResultPrivateNewFromNative(hitTestResultPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeHitTestResultPrivate)
	return structGo
}
func finalizeHitTestResultPrivate(obj *HitTestResultPrivate) {
	hitTestResultPrivateStruct.Free(obj.Native())
}

var installMissingMediaPluginsPermissionRequestClassStruct *gi.Struct
var installMissingMediaPluginsPermissionRequestClassStruct_Once sync.Once

func installMissingMediaPluginsPermissionRequestClassStruct_Set() error {
	var err error
	installMissingMediaPluginsPermissionRequestClassStruct_Once.Do(func() {
		installMissingMediaPluginsPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "InstallMissingMediaPluginsPermissionRequestClass")
	})
	return err
}

type InstallMissingMediaPluginsPermissionRequestClass struct {
	native unsafe.Pointer
}

func InstallMissingMediaPluginsPermissionRequestClassNewFromNative(native unsafe.Pointer) *InstallMissingMediaPluginsPermissionRequestClass {
	instance := &InstallMissingMediaPluginsPermissionRequestClass{native: native}

	return instance
}

/*
CastToInstallMissingMediaPluginsPermissionRequestClass down casts any arbitrary Object to InstallMissingMediaPluginsPermissionRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a InstallMissingMediaPluginsPermissionRequestClass.
*/
func CastToInstallMissingMediaPluginsPermissionRequestClass(object *gobject.Object) *InstallMissingMediaPluginsPermissionRequestClass {
	return InstallMissingMediaPluginsPermissionRequestClassNewFromNative(object.Native())
}

func (recv *InstallMissingMediaPluginsPermissionRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// InstallMissingMediaPluginsPermissionRequestClassStruct creates an uninitialised InstallMissingMediaPluginsPermissionRequestClass.
func InstallMissingMediaPluginsPermissionRequestClassStruct() *InstallMissingMediaPluginsPermissionRequestClass {
	err := installMissingMediaPluginsPermissionRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InstallMissingMediaPluginsPermissionRequestClassNewFromNative(installMissingMediaPluginsPermissionRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInstallMissingMediaPluginsPermissionRequestClass)
	return structGo
}
func finalizeInstallMissingMediaPluginsPermissionRequestClass(obj *InstallMissingMediaPluginsPermissionRequestClass) {
	installMissingMediaPluginsPermissionRequestClassStruct.Free(obj.Native())
}

var installMissingMediaPluginsPermissionRequestPrivateStruct *gi.Struct
var installMissingMediaPluginsPermissionRequestPrivateStruct_Once sync.Once

func installMissingMediaPluginsPermissionRequestPrivateStruct_Set() error {
	var err error
	installMissingMediaPluginsPermissionRequestPrivateStruct_Once.Do(func() {
		installMissingMediaPluginsPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "InstallMissingMediaPluginsPermissionRequestPrivate")
	})
	return err
}

type InstallMissingMediaPluginsPermissionRequestPrivate struct {
	native unsafe.Pointer
}

func InstallMissingMediaPluginsPermissionRequestPrivateNewFromNative(native unsafe.Pointer) *InstallMissingMediaPluginsPermissionRequestPrivate {
	instance := &InstallMissingMediaPluginsPermissionRequestPrivate{native: native}

	return instance
}

/*
CastToInstallMissingMediaPluginsPermissionRequestPrivate down casts any arbitrary Object to InstallMissingMediaPluginsPermissionRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a InstallMissingMediaPluginsPermissionRequestPrivate.
*/
func CastToInstallMissingMediaPluginsPermissionRequestPrivate(object *gobject.Object) *InstallMissingMediaPluginsPermissionRequestPrivate {
	return InstallMissingMediaPluginsPermissionRequestPrivateNewFromNative(object.Native())
}

func (recv *InstallMissingMediaPluginsPermissionRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// InstallMissingMediaPluginsPermissionRequestPrivateStruct creates an uninitialised InstallMissingMediaPluginsPermissionRequestPrivate.
func InstallMissingMediaPluginsPermissionRequestPrivateStruct() *InstallMissingMediaPluginsPermissionRequestPrivate {
	err := installMissingMediaPluginsPermissionRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := InstallMissingMediaPluginsPermissionRequestPrivateNewFromNative(installMissingMediaPluginsPermissionRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeInstallMissingMediaPluginsPermissionRequestPrivate)
	return structGo
}
func finalizeInstallMissingMediaPluginsPermissionRequestPrivate(obj *InstallMissingMediaPluginsPermissionRequestPrivate) {
	installMissingMediaPluginsPermissionRequestPrivateStruct.Free(obj.Native())
}

var javascriptResultStruct *gi.Struct
var javascriptResultStruct_Once sync.Once

func javascriptResultStruct_Set() error {
	var err error
	javascriptResultStruct_Once.Do(func() {
		javascriptResultStruct, err = gi.StructNew("WebKit2", "JavascriptResult")
	})
	return err
}

type JavascriptResult struct {
	native unsafe.Pointer
}

func JavascriptResultNewFromNative(native unsafe.Pointer) *JavascriptResult {
	instance := &JavascriptResult{native: native}

	return instance
}

/*
CastToJavascriptResult down casts any arbitrary Object to JavascriptResult.
Exercise care, as this is a potentially dangerous function
if the Object is not a JavascriptResult.
*/
func CastToJavascriptResult(object *gobject.Object) *JavascriptResult {
	return JavascriptResultNewFromNative(object.Native())
}

func (recv *JavascriptResult) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'webkit_javascript_result_get_global_context' : return type 'JavaScriptCore.GlobalContextRef' not supported

var javascriptResultGetJsValueFunction *gi.Function
var javascriptResultGetJsValueFunction_Once sync.Once

func javascriptResultGetJsValueFunction_Set() error {
	var err error
	javascriptResultGetJsValueFunction_Once.Do(func() {
		err = javascriptResultStruct_Set()
		if err != nil {
			return
		}
		javascriptResultGetJsValueFunction, err = javascriptResultStruct.InvokerNew("get_js_value")
	})
	return err
}

// GetJsValue is a representation of the C type webkit_javascript_result_get_js_value.
func (recv *JavascriptResult) GetJsValue() *javascriptcore.Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := javascriptResultGetJsValueFunction_Set()
	if err == nil {
		ret = javascriptResultGetJsValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := javascriptcore.ValueNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'webkit_javascript_result_get_value' : return type 'JavaScriptCore.ValueRef' not supported

var javascriptResultRefFunction *gi.Function
var javascriptResultRefFunction_Once sync.Once

func javascriptResultRefFunction_Set() error {
	var err error
	javascriptResultRefFunction_Once.Do(func() {
		err = javascriptResultStruct_Set()
		if err != nil {
			return
		}
		javascriptResultRefFunction, err = javascriptResultStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_javascript_result_ref.
func (recv *JavascriptResult) Ref() *JavascriptResult {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := javascriptResultRefFunction_Set()
	if err == nil {
		ret = javascriptResultRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := JavascriptResultNewFromNative(ret.Pointer())

	return retGo
}

var javascriptResultUnrefFunction *gi.Function
var javascriptResultUnrefFunction_Once sync.Once

func javascriptResultUnrefFunction_Set() error {
	var err error
	javascriptResultUnrefFunction_Once.Do(func() {
		err = javascriptResultStruct_Set()
		if err != nil {
			return
		}
		javascriptResultUnrefFunction, err = javascriptResultStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_javascript_result_unref.
func (recv *JavascriptResult) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := javascriptResultUnrefFunction_Set()
	if err == nil {
		javascriptResultUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// JavascriptResultStruct creates an uninitialised JavascriptResult.
func JavascriptResultStruct() *JavascriptResult {
	err := javascriptResultStruct_Set()
	if err != nil {
		return nil
	}

	structGo := JavascriptResultNewFromNative(javascriptResultStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeJavascriptResult)
	return structGo
}
func finalizeJavascriptResult(obj *JavascriptResult) {
	javascriptResultStruct.Free(obj.Native())
}

var mimeInfoStruct *gi.Struct
var mimeInfoStruct_Once sync.Once

func mimeInfoStruct_Set() error {
	var err error
	mimeInfoStruct_Once.Do(func() {
		mimeInfoStruct, err = gi.StructNew("WebKit2", "MimeInfo")
	})
	return err
}

type MimeInfo struct {
	native unsafe.Pointer
}

func MimeInfoNewFromNative(native unsafe.Pointer) *MimeInfo {
	instance := &MimeInfo{native: native}

	return instance
}

/*
CastToMimeInfo down casts any arbitrary Object to MimeInfo.
Exercise care, as this is a potentially dangerous function
if the Object is not a MimeInfo.
*/
func CastToMimeInfo(object *gobject.Object) *MimeInfo {
	return MimeInfoNewFromNative(object.Native())
}

func (recv *MimeInfo) Native() unsafe.Pointer {
	return recv.native
}

var mimeInfoGetDescriptionFunction *gi.Function
var mimeInfoGetDescriptionFunction_Once sync.Once

func mimeInfoGetDescriptionFunction_Set() error {
	var err error
	mimeInfoGetDescriptionFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoGetDescriptionFunction, err = mimeInfoStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type webkit_mime_info_get_description.
func (recv *MimeInfo) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mimeInfoGetDescriptionFunction_Set()
	if err == nil {
		ret = mimeInfoGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var mimeInfoGetExtensionsFunction *gi.Function
var mimeInfoGetExtensionsFunction_Once sync.Once

func mimeInfoGetExtensionsFunction_Set() error {
	var err error
	mimeInfoGetExtensionsFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoGetExtensionsFunction, err = mimeInfoStruct.InvokerNew("get_extensions")
	})
	return err
}

// GetExtensions is a representation of the C type webkit_mime_info_get_extensions.
func (recv *MimeInfo) GetExtensions() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := mimeInfoGetExtensionsFunction_Set()
	if err == nil {
		mimeInfoGetExtensionsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var mimeInfoGetMimeTypeFunction *gi.Function
var mimeInfoGetMimeTypeFunction_Once sync.Once

func mimeInfoGetMimeTypeFunction_Set() error {
	var err error
	mimeInfoGetMimeTypeFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoGetMimeTypeFunction, err = mimeInfoStruct.InvokerNew("get_mime_type")
	})
	return err
}

// GetMimeType is a representation of the C type webkit_mime_info_get_mime_type.
func (recv *MimeInfo) GetMimeType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mimeInfoGetMimeTypeFunction_Set()
	if err == nil {
		ret = mimeInfoGetMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var mimeInfoRefFunction *gi.Function
var mimeInfoRefFunction_Once sync.Once

func mimeInfoRefFunction_Set() error {
	var err error
	mimeInfoRefFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoRefFunction, err = mimeInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_mime_info_ref.
func (recv *MimeInfo) Ref() *MimeInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := mimeInfoRefFunction_Set()
	if err == nil {
		ret = mimeInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := MimeInfoNewFromNative(ret.Pointer())

	return retGo
}

var mimeInfoUnrefFunction *gi.Function
var mimeInfoUnrefFunction_Once sync.Once

func mimeInfoUnrefFunction_Set() error {
	var err error
	mimeInfoUnrefFunction_Once.Do(func() {
		err = mimeInfoStruct_Set()
		if err != nil {
			return
		}
		mimeInfoUnrefFunction, err = mimeInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_mime_info_unref.
func (recv *MimeInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := mimeInfoUnrefFunction_Set()
	if err == nil {
		mimeInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// MimeInfoStruct creates an uninitialised MimeInfo.
func MimeInfoStruct() *MimeInfo {
	err := mimeInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MimeInfoNewFromNative(mimeInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMimeInfo)
	return structGo
}
func finalizeMimeInfo(obj *MimeInfo) {
	mimeInfoStruct.Free(obj.Native())
}

var navigationActionStruct *gi.Struct
var navigationActionStruct_Once sync.Once

func navigationActionStruct_Set() error {
	var err error
	navigationActionStruct_Once.Do(func() {
		navigationActionStruct, err = gi.StructNew("WebKit2", "NavigationAction")
	})
	return err
}

type NavigationAction struct {
	native unsafe.Pointer
}

func NavigationActionNewFromNative(native unsafe.Pointer) *NavigationAction {
	instance := &NavigationAction{native: native}

	return instance
}

/*
CastToNavigationAction down casts any arbitrary Object to NavigationAction.
Exercise care, as this is a potentially dangerous function
if the Object is not a NavigationAction.
*/
func CastToNavigationAction(object *gobject.Object) *NavigationAction {
	return NavigationActionNewFromNative(object.Native())
}

func (recv *NavigationAction) Native() unsafe.Pointer {
	return recv.native
}

var navigationActionCopyFunction *gi.Function
var navigationActionCopyFunction_Once sync.Once

func navigationActionCopyFunction_Set() error {
	var err error
	navigationActionCopyFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionCopyFunction, err = navigationActionStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_navigation_action_copy.
func (recv *NavigationAction) Copy() *NavigationAction {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := navigationActionCopyFunction_Set()
	if err == nil {
		ret = navigationActionCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := NavigationActionNewFromNative(ret.Pointer())

	return retGo
}

var navigationActionFreeFunction *gi.Function
var navigationActionFreeFunction_Once sync.Once

func navigationActionFreeFunction_Set() error {
	var err error
	navigationActionFreeFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionFreeFunction, err = navigationActionStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_navigation_action_free.
func (recv *NavigationAction) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := navigationActionFreeFunction_Set()
	if err == nil {
		navigationActionFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var navigationActionGetModifiersFunction *gi.Function
var navigationActionGetModifiersFunction_Once sync.Once

func navigationActionGetModifiersFunction_Set() error {
	var err error
	navigationActionGetModifiersFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionGetModifiersFunction, err = navigationActionStruct.InvokerNew("get_modifiers")
	})
	return err
}

// GetModifiers is a representation of the C type webkit_navigation_action_get_modifiers.
func (recv *NavigationAction) GetModifiers() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := navigationActionGetModifiersFunction_Set()
	if err == nil {
		ret = navigationActionGetModifiersFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var navigationActionGetMouseButtonFunction *gi.Function
var navigationActionGetMouseButtonFunction_Once sync.Once

func navigationActionGetMouseButtonFunction_Set() error {
	var err error
	navigationActionGetMouseButtonFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionGetMouseButtonFunction, err = navigationActionStruct.InvokerNew("get_mouse_button")
	})
	return err
}

// GetMouseButton is a representation of the C type webkit_navigation_action_get_mouse_button.
func (recv *NavigationAction) GetMouseButton() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := navigationActionGetMouseButtonFunction_Set()
	if err == nil {
		ret = navigationActionGetMouseButtonFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var navigationActionGetNavigationTypeFunction *gi.Function
var navigationActionGetNavigationTypeFunction_Once sync.Once

func navigationActionGetNavigationTypeFunction_Set() error {
	var err error
	navigationActionGetNavigationTypeFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionGetNavigationTypeFunction, err = navigationActionStruct.InvokerNew("get_navigation_type")
	})
	return err
}

// GetNavigationType is a representation of the C type webkit_navigation_action_get_navigation_type.
func (recv *NavigationAction) GetNavigationType() NavigationType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := navigationActionGetNavigationTypeFunction_Set()
	if err == nil {
		ret = navigationActionGetNavigationTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := NavigationType(ret.Int32())

	return retGo
}

var navigationActionGetRequestFunction *gi.Function
var navigationActionGetRequestFunction_Once sync.Once

func navigationActionGetRequestFunction_Set() error {
	var err error
	navigationActionGetRequestFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionGetRequestFunction, err = navigationActionStruct.InvokerNew("get_request")
	})
	return err
}

// GetRequest is a representation of the C type webkit_navigation_action_get_request.
func (recv *NavigationAction) GetRequest() *URIRequest {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := navigationActionGetRequestFunction_Set()
	if err == nil {
		ret = navigationActionGetRequestFunction.Invoke(inArgs[:], nil)
	}

	retGo := URIRequestNewFromNative(ret.Pointer())

	return retGo
}

var navigationActionIsRedirectFunction *gi.Function
var navigationActionIsRedirectFunction_Once sync.Once

func navigationActionIsRedirectFunction_Set() error {
	var err error
	navigationActionIsRedirectFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionIsRedirectFunction, err = navigationActionStruct.InvokerNew("is_redirect")
	})
	return err
}

// IsRedirect is a representation of the C type webkit_navigation_action_is_redirect.
func (recv *NavigationAction) IsRedirect() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := navigationActionIsRedirectFunction_Set()
	if err == nil {
		ret = navigationActionIsRedirectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var navigationActionIsUserGestureFunction *gi.Function
var navigationActionIsUserGestureFunction_Once sync.Once

func navigationActionIsUserGestureFunction_Set() error {
	var err error
	navigationActionIsUserGestureFunction_Once.Do(func() {
		err = navigationActionStruct_Set()
		if err != nil {
			return
		}
		navigationActionIsUserGestureFunction, err = navigationActionStruct.InvokerNew("is_user_gesture")
	})
	return err
}

// IsUserGesture is a representation of the C type webkit_navigation_action_is_user_gesture.
func (recv *NavigationAction) IsUserGesture() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := navigationActionIsUserGestureFunction_Set()
	if err == nil {
		ret = navigationActionIsUserGestureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// NavigationActionStruct creates an uninitialised NavigationAction.
func NavigationActionStruct() *NavigationAction {
	err := navigationActionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NavigationActionNewFromNative(navigationActionStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNavigationAction)
	return structGo
}
func finalizeNavigationAction(obj *NavigationAction) {
	navigationActionStruct.Free(obj.Native())
}

var navigationPolicyDecisionClassStruct *gi.Struct
var navigationPolicyDecisionClassStruct_Once sync.Once

func navigationPolicyDecisionClassStruct_Set() error {
	var err error
	navigationPolicyDecisionClassStruct_Once.Do(func() {
		navigationPolicyDecisionClassStruct, err = gi.StructNew("WebKit2", "NavigationPolicyDecisionClass")
	})
	return err
}

type NavigationPolicyDecisionClass struct {
	native unsafe.Pointer
}

func NavigationPolicyDecisionClassNewFromNative(native unsafe.Pointer) *NavigationPolicyDecisionClass {
	instance := &NavigationPolicyDecisionClass{native: native}

	return instance
}

/*
CastToNavigationPolicyDecisionClass down casts any arbitrary Object to NavigationPolicyDecisionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a NavigationPolicyDecisionClass.
*/
func CastToNavigationPolicyDecisionClass(object *gobject.Object) *NavigationPolicyDecisionClass {
	return NavigationPolicyDecisionClassNewFromNative(object.Native())
}

func (recv *NavigationPolicyDecisionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *NavigationPolicyDecisionClass) FieldParentClass() *PolicyDecisionClass {
	argValue := gi.StructFieldGet(navigationPolicyDecisionClassStruct, recv.Native(), "parent_class")
	value := PolicyDecisionClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *NavigationPolicyDecisionClass) SetFieldParentClass(value *PolicyDecisionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(navigationPolicyDecisionClassStruct, recv.Native(), "parent_class", argValue)
}

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// NavigationPolicyDecisionClassStruct creates an uninitialised NavigationPolicyDecisionClass.
func NavigationPolicyDecisionClassStruct() *NavigationPolicyDecisionClass {
	err := navigationPolicyDecisionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NavigationPolicyDecisionClassNewFromNative(navigationPolicyDecisionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNavigationPolicyDecisionClass)
	return structGo
}
func finalizeNavigationPolicyDecisionClass(obj *NavigationPolicyDecisionClass) {
	navigationPolicyDecisionClassStruct.Free(obj.Native())
}

var navigationPolicyDecisionPrivateStruct *gi.Struct
var navigationPolicyDecisionPrivateStruct_Once sync.Once

func navigationPolicyDecisionPrivateStruct_Set() error {
	var err error
	navigationPolicyDecisionPrivateStruct_Once.Do(func() {
		navigationPolicyDecisionPrivateStruct, err = gi.StructNew("WebKit2", "NavigationPolicyDecisionPrivate")
	})
	return err
}

type NavigationPolicyDecisionPrivate struct {
	native unsafe.Pointer
}

func NavigationPolicyDecisionPrivateNewFromNative(native unsafe.Pointer) *NavigationPolicyDecisionPrivate {
	instance := &NavigationPolicyDecisionPrivate{native: native}

	return instance
}

/*
CastToNavigationPolicyDecisionPrivate down casts any arbitrary Object to NavigationPolicyDecisionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a NavigationPolicyDecisionPrivate.
*/
func CastToNavigationPolicyDecisionPrivate(object *gobject.Object) *NavigationPolicyDecisionPrivate {
	return NavigationPolicyDecisionPrivateNewFromNative(object.Native())
}

func (recv *NavigationPolicyDecisionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// NavigationPolicyDecisionPrivateStruct creates an uninitialised NavigationPolicyDecisionPrivate.
func NavigationPolicyDecisionPrivateStruct() *NavigationPolicyDecisionPrivate {
	err := navigationPolicyDecisionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NavigationPolicyDecisionPrivateNewFromNative(navigationPolicyDecisionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNavigationPolicyDecisionPrivate)
	return structGo
}
func finalizeNavigationPolicyDecisionPrivate(obj *NavigationPolicyDecisionPrivate) {
	navigationPolicyDecisionPrivateStruct.Free(obj.Native())
}

var networkProxySettingsStruct *gi.Struct
var networkProxySettingsStruct_Once sync.Once

func networkProxySettingsStruct_Set() error {
	var err error
	networkProxySettingsStruct_Once.Do(func() {
		networkProxySettingsStruct, err = gi.StructNew("WebKit2", "NetworkProxySettings")
	})
	return err
}

type NetworkProxySettings struct {
	native unsafe.Pointer
}

func NetworkProxySettingsNewFromNative(native unsafe.Pointer) *NetworkProxySettings {
	instance := &NetworkProxySettings{native: native}

	return instance
}

/*
CastToNetworkProxySettings down casts any arbitrary Object to NetworkProxySettings.
Exercise care, as this is a potentially dangerous function
if the Object is not a NetworkProxySettings.
*/
func CastToNetworkProxySettings(object *gobject.Object) *NetworkProxySettings {
	return NetworkProxySettingsNewFromNative(object.Native())
}

func (recv *NetworkProxySettings) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'webkit_network_proxy_settings_new' : parameter 'ignore_hosts' of type 'nil' not supported

var networkProxySettingsAddProxyForSchemeFunction *gi.Function
var networkProxySettingsAddProxyForSchemeFunction_Once sync.Once

func networkProxySettingsAddProxyForSchemeFunction_Set() error {
	var err error
	networkProxySettingsAddProxyForSchemeFunction_Once.Do(func() {
		err = networkProxySettingsStruct_Set()
		if err != nil {
			return
		}
		networkProxySettingsAddProxyForSchemeFunction, err = networkProxySettingsStruct.InvokerNew("add_proxy_for_scheme")
	})
	return err
}

// AddProxyForScheme is a representation of the C type webkit_network_proxy_settings_add_proxy_for_scheme.
func (recv *NetworkProxySettings) AddProxyForScheme(scheme string, proxyUri string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(scheme)
	inArgs[2].SetString(proxyUri)

	err := networkProxySettingsAddProxyForSchemeFunction_Set()
	if err == nil {
		networkProxySettingsAddProxyForSchemeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var networkProxySettingsCopyFunction *gi.Function
var networkProxySettingsCopyFunction_Once sync.Once

func networkProxySettingsCopyFunction_Set() error {
	var err error
	networkProxySettingsCopyFunction_Once.Do(func() {
		err = networkProxySettingsStruct_Set()
		if err != nil {
			return
		}
		networkProxySettingsCopyFunction, err = networkProxySettingsStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_network_proxy_settings_copy.
func (recv *NetworkProxySettings) Copy() *NetworkProxySettings {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := networkProxySettingsCopyFunction_Set()
	if err == nil {
		ret = networkProxySettingsCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := NetworkProxySettingsNewFromNative(ret.Pointer())

	return retGo
}

var networkProxySettingsFreeFunction *gi.Function
var networkProxySettingsFreeFunction_Once sync.Once

func networkProxySettingsFreeFunction_Set() error {
	var err error
	networkProxySettingsFreeFunction_Once.Do(func() {
		err = networkProxySettingsStruct_Set()
		if err != nil {
			return
		}
		networkProxySettingsFreeFunction, err = networkProxySettingsStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_network_proxy_settings_free.
func (recv *NetworkProxySettings) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := networkProxySettingsFreeFunction_Set()
	if err == nil {
		networkProxySettingsFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var notificationClassStruct *gi.Struct
var notificationClassStruct_Once sync.Once

func notificationClassStruct_Set() error {
	var err error
	notificationClassStruct_Once.Do(func() {
		notificationClassStruct, err = gi.StructNew("WebKit2", "NotificationClass")
	})
	return err
}

type NotificationClass struct {
	native unsafe.Pointer
}

func NotificationClassNewFromNative(native unsafe.Pointer) *NotificationClass {
	instance := &NotificationClass{native: native}

	return instance
}

/*
CastToNotificationClass down casts any arbitrary Object to NotificationClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a NotificationClass.
*/
func CastToNotificationClass(object *gobject.Object) *NotificationClass {
	return NotificationClassNewFromNative(object.Native())
}

func (recv *NotificationClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved5' : for field setter : missing Type

// NotificationClassStruct creates an uninitialised NotificationClass.
func NotificationClassStruct() *NotificationClass {
	err := notificationClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NotificationClassNewFromNative(notificationClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNotificationClass)
	return structGo
}
func finalizeNotificationClass(obj *NotificationClass) {
	notificationClassStruct.Free(obj.Native())
}

var notificationPermissionRequestClassStruct *gi.Struct
var notificationPermissionRequestClassStruct_Once sync.Once

func notificationPermissionRequestClassStruct_Set() error {
	var err error
	notificationPermissionRequestClassStruct_Once.Do(func() {
		notificationPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "NotificationPermissionRequestClass")
	})
	return err
}

type NotificationPermissionRequestClass struct {
	native unsafe.Pointer
}

func NotificationPermissionRequestClassNewFromNative(native unsafe.Pointer) *NotificationPermissionRequestClass {
	instance := &NotificationPermissionRequestClass{native: native}

	return instance
}

/*
CastToNotificationPermissionRequestClass down casts any arbitrary Object to NotificationPermissionRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a NotificationPermissionRequestClass.
*/
func CastToNotificationPermissionRequestClass(object *gobject.Object) *NotificationPermissionRequestClass {
	return NotificationPermissionRequestClassNewFromNative(object.Native())
}

func (recv *NotificationPermissionRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// NotificationPermissionRequestClassStruct creates an uninitialised NotificationPermissionRequestClass.
func NotificationPermissionRequestClassStruct() *NotificationPermissionRequestClass {
	err := notificationPermissionRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NotificationPermissionRequestClassNewFromNative(notificationPermissionRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNotificationPermissionRequestClass)
	return structGo
}
func finalizeNotificationPermissionRequestClass(obj *NotificationPermissionRequestClass) {
	notificationPermissionRequestClassStruct.Free(obj.Native())
}

var notificationPermissionRequestPrivateStruct *gi.Struct
var notificationPermissionRequestPrivateStruct_Once sync.Once

func notificationPermissionRequestPrivateStruct_Set() error {
	var err error
	notificationPermissionRequestPrivateStruct_Once.Do(func() {
		notificationPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "NotificationPermissionRequestPrivate")
	})
	return err
}

type NotificationPermissionRequestPrivate struct {
	native unsafe.Pointer
}

func NotificationPermissionRequestPrivateNewFromNative(native unsafe.Pointer) *NotificationPermissionRequestPrivate {
	instance := &NotificationPermissionRequestPrivate{native: native}

	return instance
}

/*
CastToNotificationPermissionRequestPrivate down casts any arbitrary Object to NotificationPermissionRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a NotificationPermissionRequestPrivate.
*/
func CastToNotificationPermissionRequestPrivate(object *gobject.Object) *NotificationPermissionRequestPrivate {
	return NotificationPermissionRequestPrivateNewFromNative(object.Native())
}

func (recv *NotificationPermissionRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// NotificationPermissionRequestPrivateStruct creates an uninitialised NotificationPermissionRequestPrivate.
func NotificationPermissionRequestPrivateStruct() *NotificationPermissionRequestPrivate {
	err := notificationPermissionRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NotificationPermissionRequestPrivateNewFromNative(notificationPermissionRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNotificationPermissionRequestPrivate)
	return structGo
}
func finalizeNotificationPermissionRequestPrivate(obj *NotificationPermissionRequestPrivate) {
	notificationPermissionRequestPrivateStruct.Free(obj.Native())
}

var notificationPrivateStruct *gi.Struct
var notificationPrivateStruct_Once sync.Once

func notificationPrivateStruct_Set() error {
	var err error
	notificationPrivateStruct_Once.Do(func() {
		notificationPrivateStruct, err = gi.StructNew("WebKit2", "NotificationPrivate")
	})
	return err
}

type NotificationPrivate struct {
	native unsafe.Pointer
}

func NotificationPrivateNewFromNative(native unsafe.Pointer) *NotificationPrivate {
	instance := &NotificationPrivate{native: native}

	return instance
}

/*
CastToNotificationPrivate down casts any arbitrary Object to NotificationPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a NotificationPrivate.
*/
func CastToNotificationPrivate(object *gobject.Object) *NotificationPrivate {
	return NotificationPrivateNewFromNative(object.Native())
}

func (recv *NotificationPrivate) Native() unsafe.Pointer {
	return recv.native
}

// NotificationPrivateStruct creates an uninitialised NotificationPrivate.
func NotificationPrivateStruct() *NotificationPrivate {
	err := notificationPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NotificationPrivateNewFromNative(notificationPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNotificationPrivate)
	return structGo
}
func finalizeNotificationPrivate(obj *NotificationPrivate) {
	notificationPrivateStruct.Free(obj.Native())
}

var optionMenuClassStruct *gi.Struct
var optionMenuClassStruct_Once sync.Once

func optionMenuClassStruct_Set() error {
	var err error
	optionMenuClassStruct_Once.Do(func() {
		optionMenuClassStruct, err = gi.StructNew("WebKit2", "OptionMenuClass")
	})
	return err
}

type OptionMenuClass struct {
	native unsafe.Pointer
}

func OptionMenuClassNewFromNative(native unsafe.Pointer) *OptionMenuClass {
	instance := &OptionMenuClass{native: native}

	return instance
}

/*
CastToOptionMenuClass down casts any arbitrary Object to OptionMenuClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a OptionMenuClass.
*/
func CastToOptionMenuClass(object *gobject.Object) *OptionMenuClass {
	return OptionMenuClassNewFromNative(object.Native())
}

func (recv *OptionMenuClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// OptionMenuClassStruct creates an uninitialised OptionMenuClass.
func OptionMenuClassStruct() *OptionMenuClass {
	err := optionMenuClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := OptionMenuClassNewFromNative(optionMenuClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeOptionMenuClass)
	return structGo
}
func finalizeOptionMenuClass(obj *OptionMenuClass) {
	optionMenuClassStruct.Free(obj.Native())
}

var optionMenuItemStruct *gi.Struct
var optionMenuItemStruct_Once sync.Once

func optionMenuItemStruct_Set() error {
	var err error
	optionMenuItemStruct_Once.Do(func() {
		optionMenuItemStruct, err = gi.StructNew("WebKit2", "OptionMenuItem")
	})
	return err
}

type OptionMenuItem struct {
	native unsafe.Pointer
}

func OptionMenuItemNewFromNative(native unsafe.Pointer) *OptionMenuItem {
	instance := &OptionMenuItem{native: native}

	return instance
}

/*
CastToOptionMenuItem down casts any arbitrary Object to OptionMenuItem.
Exercise care, as this is a potentially dangerous function
if the Object is not a OptionMenuItem.
*/
func CastToOptionMenuItem(object *gobject.Object) *OptionMenuItem {
	return OptionMenuItemNewFromNative(object.Native())
}

func (recv *OptionMenuItem) Native() unsafe.Pointer {
	return recv.native
}

var optionMenuItemCopyFunction *gi.Function
var optionMenuItemCopyFunction_Once sync.Once

func optionMenuItemCopyFunction_Set() error {
	var err error
	optionMenuItemCopyFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemCopyFunction, err = optionMenuItemStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type webkit_option_menu_item_copy.
func (recv *OptionMenuItem) Copy() *OptionMenuItem {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := optionMenuItemCopyFunction_Set()
	if err == nil {
		ret = optionMenuItemCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := OptionMenuItemNewFromNative(ret.Pointer())

	return retGo
}

var optionMenuItemFreeFunction *gi.Function
var optionMenuItemFreeFunction_Once sync.Once

func optionMenuItemFreeFunction_Set() error {
	var err error
	optionMenuItemFreeFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemFreeFunction, err = optionMenuItemStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type webkit_option_menu_item_free.
func (recv *OptionMenuItem) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := optionMenuItemFreeFunction_Set()
	if err == nil {
		optionMenuItemFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var optionMenuItemGetLabelFunction *gi.Function
var optionMenuItemGetLabelFunction_Once sync.Once

func optionMenuItemGetLabelFunction_Set() error {
	var err error
	optionMenuItemGetLabelFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemGetLabelFunction, err = optionMenuItemStruct.InvokerNew("get_label")
	})
	return err
}

// GetLabel is a representation of the C type webkit_option_menu_item_get_label.
func (recv *OptionMenuItem) GetLabel() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := optionMenuItemGetLabelFunction_Set()
	if err == nil {
		ret = optionMenuItemGetLabelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var optionMenuItemGetTooltipFunction *gi.Function
var optionMenuItemGetTooltipFunction_Once sync.Once

func optionMenuItemGetTooltipFunction_Set() error {
	var err error
	optionMenuItemGetTooltipFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemGetTooltipFunction, err = optionMenuItemStruct.InvokerNew("get_tooltip")
	})
	return err
}

// GetTooltip is a representation of the C type webkit_option_menu_item_get_tooltip.
func (recv *OptionMenuItem) GetTooltip() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := optionMenuItemGetTooltipFunction_Set()
	if err == nil {
		ret = optionMenuItemGetTooltipFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var optionMenuItemIsEnabledFunction *gi.Function
var optionMenuItemIsEnabledFunction_Once sync.Once

func optionMenuItemIsEnabledFunction_Set() error {
	var err error
	optionMenuItemIsEnabledFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemIsEnabledFunction, err = optionMenuItemStruct.InvokerNew("is_enabled")
	})
	return err
}

// IsEnabled is a representation of the C type webkit_option_menu_item_is_enabled.
func (recv *OptionMenuItem) IsEnabled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := optionMenuItemIsEnabledFunction_Set()
	if err == nil {
		ret = optionMenuItemIsEnabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var optionMenuItemIsGroupChildFunction *gi.Function
var optionMenuItemIsGroupChildFunction_Once sync.Once

func optionMenuItemIsGroupChildFunction_Set() error {
	var err error
	optionMenuItemIsGroupChildFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemIsGroupChildFunction, err = optionMenuItemStruct.InvokerNew("is_group_child")
	})
	return err
}

// IsGroupChild is a representation of the C type webkit_option_menu_item_is_group_child.
func (recv *OptionMenuItem) IsGroupChild() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := optionMenuItemIsGroupChildFunction_Set()
	if err == nil {
		ret = optionMenuItemIsGroupChildFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var optionMenuItemIsGroupLabelFunction *gi.Function
var optionMenuItemIsGroupLabelFunction_Once sync.Once

func optionMenuItemIsGroupLabelFunction_Set() error {
	var err error
	optionMenuItemIsGroupLabelFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemIsGroupLabelFunction, err = optionMenuItemStruct.InvokerNew("is_group_label")
	})
	return err
}

// IsGroupLabel is a representation of the C type webkit_option_menu_item_is_group_label.
func (recv *OptionMenuItem) IsGroupLabel() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := optionMenuItemIsGroupLabelFunction_Set()
	if err == nil {
		ret = optionMenuItemIsGroupLabelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var optionMenuItemIsSelectedFunction *gi.Function
var optionMenuItemIsSelectedFunction_Once sync.Once

func optionMenuItemIsSelectedFunction_Set() error {
	var err error
	optionMenuItemIsSelectedFunction_Once.Do(func() {
		err = optionMenuItemStruct_Set()
		if err != nil {
			return
		}
		optionMenuItemIsSelectedFunction, err = optionMenuItemStruct.InvokerNew("is_selected")
	})
	return err
}

// IsSelected is a representation of the C type webkit_option_menu_item_is_selected.
func (recv *OptionMenuItem) IsSelected() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := optionMenuItemIsSelectedFunction_Set()
	if err == nil {
		ret = optionMenuItemIsSelectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// OptionMenuItemStruct creates an uninitialised OptionMenuItem.
func OptionMenuItemStruct() *OptionMenuItem {
	err := optionMenuItemStruct_Set()
	if err != nil {
		return nil
	}

	structGo := OptionMenuItemNewFromNative(optionMenuItemStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeOptionMenuItem)
	return structGo
}
func finalizeOptionMenuItem(obj *OptionMenuItem) {
	optionMenuItemStruct.Free(obj.Native())
}

var optionMenuPrivateStruct *gi.Struct
var optionMenuPrivateStruct_Once sync.Once

func optionMenuPrivateStruct_Set() error {
	var err error
	optionMenuPrivateStruct_Once.Do(func() {
		optionMenuPrivateStruct, err = gi.StructNew("WebKit2", "OptionMenuPrivate")
	})
	return err
}

type OptionMenuPrivate struct {
	native unsafe.Pointer
}

func OptionMenuPrivateNewFromNative(native unsafe.Pointer) *OptionMenuPrivate {
	instance := &OptionMenuPrivate{native: native}

	return instance
}

/*
CastToOptionMenuPrivate down casts any arbitrary Object to OptionMenuPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a OptionMenuPrivate.
*/
func CastToOptionMenuPrivate(object *gobject.Object) *OptionMenuPrivate {
	return OptionMenuPrivateNewFromNative(object.Native())
}

func (recv *OptionMenuPrivate) Native() unsafe.Pointer {
	return recv.native
}

// OptionMenuPrivateStruct creates an uninitialised OptionMenuPrivate.
func OptionMenuPrivateStruct() *OptionMenuPrivate {
	err := optionMenuPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := OptionMenuPrivateNewFromNative(optionMenuPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeOptionMenuPrivate)
	return structGo
}
func finalizeOptionMenuPrivate(obj *OptionMenuPrivate) {
	optionMenuPrivateStruct.Free(obj.Native())
}

var permissionRequestIfaceStruct *gi.Struct
var permissionRequestIfaceStruct_Once sync.Once

func permissionRequestIfaceStruct_Set() error {
	var err error
	permissionRequestIfaceStruct_Once.Do(func() {
		permissionRequestIfaceStruct, err = gi.StructNew("WebKit2", "PermissionRequestIface")
	})
	return err
}

type PermissionRequestIface struct {
	native unsafe.Pointer
}

func PermissionRequestIfaceNewFromNative(native unsafe.Pointer) *PermissionRequestIface {
	instance := &PermissionRequestIface{native: native}

	return instance
}

/*
CastToPermissionRequestIface down casts any arbitrary Object to PermissionRequestIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a PermissionRequestIface.
*/
func CastToPermissionRequestIface(object *gobject.Object) *PermissionRequestIface {
	return PermissionRequestIfaceNewFromNative(object.Native())
}

func (recv *PermissionRequestIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_interface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'parent_interface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'allow' : for field getter : missing Type

// UNSUPPORTED : C value 'allow' : for field setter : missing Type

// UNSUPPORTED : C value 'deny' : for field getter : missing Type

// UNSUPPORTED : C value 'deny' : for field setter : missing Type

// PermissionRequestIfaceStruct creates an uninitialised PermissionRequestIface.
func PermissionRequestIfaceStruct() *PermissionRequestIface {
	err := permissionRequestIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PermissionRequestIfaceNewFromNative(permissionRequestIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePermissionRequestIface)
	return structGo
}
func finalizePermissionRequestIface(obj *PermissionRequestIface) {
	permissionRequestIfaceStruct.Free(obj.Native())
}

var pluginClassStruct *gi.Struct
var pluginClassStruct_Once sync.Once

func pluginClassStruct_Set() error {
	var err error
	pluginClassStruct_Once.Do(func() {
		pluginClassStruct, err = gi.StructNew("WebKit2", "PluginClass")
	})
	return err
}

type PluginClass struct {
	native unsafe.Pointer
}

func PluginClassNewFromNative(native unsafe.Pointer) *PluginClass {
	instance := &PluginClass{native: native}

	return instance
}

/*
CastToPluginClass down casts any arbitrary Object to PluginClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a PluginClass.
*/
func CastToPluginClass(object *gobject.Object) *PluginClass {
	return PluginClassNewFromNative(object.Native())
}

func (recv *PluginClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// PluginClassStruct creates an uninitialised PluginClass.
func PluginClassStruct() *PluginClass {
	err := pluginClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PluginClassNewFromNative(pluginClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePluginClass)
	return structGo
}
func finalizePluginClass(obj *PluginClass) {
	pluginClassStruct.Free(obj.Native())
}

var pluginPrivateStruct *gi.Struct
var pluginPrivateStruct_Once sync.Once

func pluginPrivateStruct_Set() error {
	var err error
	pluginPrivateStruct_Once.Do(func() {
		pluginPrivateStruct, err = gi.StructNew("WebKit2", "PluginPrivate")
	})
	return err
}

type PluginPrivate struct {
	native unsafe.Pointer
}

func PluginPrivateNewFromNative(native unsafe.Pointer) *PluginPrivate {
	instance := &PluginPrivate{native: native}

	return instance
}

/*
CastToPluginPrivate down casts any arbitrary Object to PluginPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a PluginPrivate.
*/
func CastToPluginPrivate(object *gobject.Object) *PluginPrivate {
	return PluginPrivateNewFromNative(object.Native())
}

func (recv *PluginPrivate) Native() unsafe.Pointer {
	return recv.native
}

// PluginPrivateStruct creates an uninitialised PluginPrivate.
func PluginPrivateStruct() *PluginPrivate {
	err := pluginPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PluginPrivateNewFromNative(pluginPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePluginPrivate)
	return structGo
}
func finalizePluginPrivate(obj *PluginPrivate) {
	pluginPrivateStruct.Free(obj.Native())
}

var policyDecisionClassStruct *gi.Struct
var policyDecisionClassStruct_Once sync.Once

func policyDecisionClassStruct_Set() error {
	var err error
	policyDecisionClassStruct_Once.Do(func() {
		policyDecisionClassStruct, err = gi.StructNew("WebKit2", "PolicyDecisionClass")
	})
	return err
}

type PolicyDecisionClass struct {
	native unsafe.Pointer
}

func PolicyDecisionClassNewFromNative(native unsafe.Pointer) *PolicyDecisionClass {
	instance := &PolicyDecisionClass{native: native}

	return instance
}

/*
CastToPolicyDecisionClass down casts any arbitrary Object to PolicyDecisionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a PolicyDecisionClass.
*/
func CastToPolicyDecisionClass(object *gobject.Object) *PolicyDecisionClass {
	return PolicyDecisionClassNewFromNative(object.Native())
}

func (recv *PolicyDecisionClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// PolicyDecisionClassStruct creates an uninitialised PolicyDecisionClass.
func PolicyDecisionClassStruct() *PolicyDecisionClass {
	err := policyDecisionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PolicyDecisionClassNewFromNative(policyDecisionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePolicyDecisionClass)
	return structGo
}
func finalizePolicyDecisionClass(obj *PolicyDecisionClass) {
	policyDecisionClassStruct.Free(obj.Native())
}

var policyDecisionPrivateStruct *gi.Struct
var policyDecisionPrivateStruct_Once sync.Once

func policyDecisionPrivateStruct_Set() error {
	var err error
	policyDecisionPrivateStruct_Once.Do(func() {
		policyDecisionPrivateStruct, err = gi.StructNew("WebKit2", "PolicyDecisionPrivate")
	})
	return err
}

type PolicyDecisionPrivate struct {
	native unsafe.Pointer
}

func PolicyDecisionPrivateNewFromNative(native unsafe.Pointer) *PolicyDecisionPrivate {
	instance := &PolicyDecisionPrivate{native: native}

	return instance
}

/*
CastToPolicyDecisionPrivate down casts any arbitrary Object to PolicyDecisionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a PolicyDecisionPrivate.
*/
func CastToPolicyDecisionPrivate(object *gobject.Object) *PolicyDecisionPrivate {
	return PolicyDecisionPrivateNewFromNative(object.Native())
}

func (recv *PolicyDecisionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// PolicyDecisionPrivateStruct creates an uninitialised PolicyDecisionPrivate.
func PolicyDecisionPrivateStruct() *PolicyDecisionPrivate {
	err := policyDecisionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PolicyDecisionPrivateNewFromNative(policyDecisionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePolicyDecisionPrivate)
	return structGo
}
func finalizePolicyDecisionPrivate(obj *PolicyDecisionPrivate) {
	policyDecisionPrivateStruct.Free(obj.Native())
}

var printCustomWidgetClassStruct *gi.Struct
var printCustomWidgetClassStruct_Once sync.Once

func printCustomWidgetClassStruct_Set() error {
	var err error
	printCustomWidgetClassStruct_Once.Do(func() {
		printCustomWidgetClassStruct, err = gi.StructNew("WebKit2", "PrintCustomWidgetClass")
	})
	return err
}

type PrintCustomWidgetClass struct {
	native unsafe.Pointer
}

func PrintCustomWidgetClassNewFromNative(native unsafe.Pointer) *PrintCustomWidgetClass {
	instance := &PrintCustomWidgetClass{native: native}

	return instance
}

/*
CastToPrintCustomWidgetClass down casts any arbitrary Object to PrintCustomWidgetClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a PrintCustomWidgetClass.
*/
func CastToPrintCustomWidgetClass(object *gobject.Object) *PrintCustomWidgetClass {
	return PrintCustomWidgetClassNewFromNative(object.Native())
}

func (recv *PrintCustomWidgetClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'apply' : for field getter : missing Type

// UNSUPPORTED : C value 'apply' : for field setter : missing Type

// UNSUPPORTED : C value 'update' : for field getter : missing Type

// UNSUPPORTED : C value 'update' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// PrintCustomWidgetClassStruct creates an uninitialised PrintCustomWidgetClass.
func PrintCustomWidgetClassStruct() *PrintCustomWidgetClass {
	err := printCustomWidgetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PrintCustomWidgetClassNewFromNative(printCustomWidgetClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePrintCustomWidgetClass)
	return structGo
}
func finalizePrintCustomWidgetClass(obj *PrintCustomWidgetClass) {
	printCustomWidgetClassStruct.Free(obj.Native())
}

var printCustomWidgetPrivateStruct *gi.Struct
var printCustomWidgetPrivateStruct_Once sync.Once

func printCustomWidgetPrivateStruct_Set() error {
	var err error
	printCustomWidgetPrivateStruct_Once.Do(func() {
		printCustomWidgetPrivateStruct, err = gi.StructNew("WebKit2", "PrintCustomWidgetPrivate")
	})
	return err
}

type PrintCustomWidgetPrivate struct {
	native unsafe.Pointer
}

func PrintCustomWidgetPrivateNewFromNative(native unsafe.Pointer) *PrintCustomWidgetPrivate {
	instance := &PrintCustomWidgetPrivate{native: native}

	return instance
}

/*
CastToPrintCustomWidgetPrivate down casts any arbitrary Object to PrintCustomWidgetPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a PrintCustomWidgetPrivate.
*/
func CastToPrintCustomWidgetPrivate(object *gobject.Object) *PrintCustomWidgetPrivate {
	return PrintCustomWidgetPrivateNewFromNative(object.Native())
}

func (recv *PrintCustomWidgetPrivate) Native() unsafe.Pointer {
	return recv.native
}

// PrintCustomWidgetPrivateStruct creates an uninitialised PrintCustomWidgetPrivate.
func PrintCustomWidgetPrivateStruct() *PrintCustomWidgetPrivate {
	err := printCustomWidgetPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PrintCustomWidgetPrivateNewFromNative(printCustomWidgetPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePrintCustomWidgetPrivate)
	return structGo
}
func finalizePrintCustomWidgetPrivate(obj *PrintCustomWidgetPrivate) {
	printCustomWidgetPrivateStruct.Free(obj.Native())
}

var printOperationClassStruct *gi.Struct
var printOperationClassStruct_Once sync.Once

func printOperationClassStruct_Set() error {
	var err error
	printOperationClassStruct_Once.Do(func() {
		printOperationClassStruct, err = gi.StructNew("WebKit2", "PrintOperationClass")
	})
	return err
}

type PrintOperationClass struct {
	native unsafe.Pointer
}

func PrintOperationClassNewFromNative(native unsafe.Pointer) *PrintOperationClass {
	instance := &PrintOperationClass{native: native}

	return instance
}

/*
CastToPrintOperationClass down casts any arbitrary Object to PrintOperationClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a PrintOperationClass.
*/
func CastToPrintOperationClass(object *gobject.Object) *PrintOperationClass {
	return PrintOperationClassNewFromNative(object.Native())
}

func (recv *PrintOperationClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// PrintOperationClassStruct creates an uninitialised PrintOperationClass.
func PrintOperationClassStruct() *PrintOperationClass {
	err := printOperationClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PrintOperationClassNewFromNative(printOperationClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePrintOperationClass)
	return structGo
}
func finalizePrintOperationClass(obj *PrintOperationClass) {
	printOperationClassStruct.Free(obj.Native())
}

var printOperationPrivateStruct *gi.Struct
var printOperationPrivateStruct_Once sync.Once

func printOperationPrivateStruct_Set() error {
	var err error
	printOperationPrivateStruct_Once.Do(func() {
		printOperationPrivateStruct, err = gi.StructNew("WebKit2", "PrintOperationPrivate")
	})
	return err
}

type PrintOperationPrivate struct {
	native unsafe.Pointer
}

func PrintOperationPrivateNewFromNative(native unsafe.Pointer) *PrintOperationPrivate {
	instance := &PrintOperationPrivate{native: native}

	return instance
}

/*
CastToPrintOperationPrivate down casts any arbitrary Object to PrintOperationPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a PrintOperationPrivate.
*/
func CastToPrintOperationPrivate(object *gobject.Object) *PrintOperationPrivate {
	return PrintOperationPrivateNewFromNative(object.Native())
}

func (recv *PrintOperationPrivate) Native() unsafe.Pointer {
	return recv.native
}

// PrintOperationPrivateStruct creates an uninitialised PrintOperationPrivate.
func PrintOperationPrivateStruct() *PrintOperationPrivate {
	err := printOperationPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PrintOperationPrivateNewFromNative(printOperationPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePrintOperationPrivate)
	return structGo
}
func finalizePrintOperationPrivate(obj *PrintOperationPrivate) {
	printOperationPrivateStruct.Free(obj.Native())
}

var responsePolicyDecisionClassStruct *gi.Struct
var responsePolicyDecisionClassStruct_Once sync.Once

func responsePolicyDecisionClassStruct_Set() error {
	var err error
	responsePolicyDecisionClassStruct_Once.Do(func() {
		responsePolicyDecisionClassStruct, err = gi.StructNew("WebKit2", "ResponsePolicyDecisionClass")
	})
	return err
}

type ResponsePolicyDecisionClass struct {
	native unsafe.Pointer
}

func ResponsePolicyDecisionClassNewFromNative(native unsafe.Pointer) *ResponsePolicyDecisionClass {
	instance := &ResponsePolicyDecisionClass{native: native}

	return instance
}

/*
CastToResponsePolicyDecisionClass down casts any arbitrary Object to ResponsePolicyDecisionClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ResponsePolicyDecisionClass.
*/
func CastToResponsePolicyDecisionClass(object *gobject.Object) *ResponsePolicyDecisionClass {
	return ResponsePolicyDecisionClassNewFromNative(object.Native())
}

func (recv *ResponsePolicyDecisionClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ResponsePolicyDecisionClass) FieldParentClass() *PolicyDecisionClass {
	argValue := gi.StructFieldGet(responsePolicyDecisionClassStruct, recv.Native(), "parent_class")
	value := PolicyDecisionClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ResponsePolicyDecisionClass) SetFieldParentClass(value *PolicyDecisionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(responsePolicyDecisionClassStruct, recv.Native(), "parent_class", argValue)
}

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// ResponsePolicyDecisionClassStruct creates an uninitialised ResponsePolicyDecisionClass.
func ResponsePolicyDecisionClassStruct() *ResponsePolicyDecisionClass {
	err := responsePolicyDecisionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ResponsePolicyDecisionClassNewFromNative(responsePolicyDecisionClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeResponsePolicyDecisionClass)
	return structGo
}
func finalizeResponsePolicyDecisionClass(obj *ResponsePolicyDecisionClass) {
	responsePolicyDecisionClassStruct.Free(obj.Native())
}

var responsePolicyDecisionPrivateStruct *gi.Struct
var responsePolicyDecisionPrivateStruct_Once sync.Once

func responsePolicyDecisionPrivateStruct_Set() error {
	var err error
	responsePolicyDecisionPrivateStruct_Once.Do(func() {
		responsePolicyDecisionPrivateStruct, err = gi.StructNew("WebKit2", "ResponsePolicyDecisionPrivate")
	})
	return err
}

type ResponsePolicyDecisionPrivate struct {
	native unsafe.Pointer
}

func ResponsePolicyDecisionPrivateNewFromNative(native unsafe.Pointer) *ResponsePolicyDecisionPrivate {
	instance := &ResponsePolicyDecisionPrivate{native: native}

	return instance
}

/*
CastToResponsePolicyDecisionPrivate down casts any arbitrary Object to ResponsePolicyDecisionPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a ResponsePolicyDecisionPrivate.
*/
func CastToResponsePolicyDecisionPrivate(object *gobject.Object) *ResponsePolicyDecisionPrivate {
	return ResponsePolicyDecisionPrivateNewFromNative(object.Native())
}

func (recv *ResponsePolicyDecisionPrivate) Native() unsafe.Pointer {
	return recv.native
}

// ResponsePolicyDecisionPrivateStruct creates an uninitialised ResponsePolicyDecisionPrivate.
func ResponsePolicyDecisionPrivateStruct() *ResponsePolicyDecisionPrivate {
	err := responsePolicyDecisionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ResponsePolicyDecisionPrivateNewFromNative(responsePolicyDecisionPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeResponsePolicyDecisionPrivate)
	return structGo
}
func finalizeResponsePolicyDecisionPrivate(obj *ResponsePolicyDecisionPrivate) {
	responsePolicyDecisionPrivateStruct.Free(obj.Native())
}

var scriptDialogStruct *gi.Struct
var scriptDialogStruct_Once sync.Once

func scriptDialogStruct_Set() error {
	var err error
	scriptDialogStruct_Once.Do(func() {
		scriptDialogStruct, err = gi.StructNew("WebKit2", "ScriptDialog")
	})
	return err
}

type ScriptDialog struct {
	native unsafe.Pointer
}

func ScriptDialogNewFromNative(native unsafe.Pointer) *ScriptDialog {
	instance := &ScriptDialog{native: native}

	return instance
}

/*
CastToScriptDialog down casts any arbitrary Object to ScriptDialog.
Exercise care, as this is a potentially dangerous function
if the Object is not a ScriptDialog.
*/
func CastToScriptDialog(object *gobject.Object) *ScriptDialog {
	return ScriptDialogNewFromNative(object.Native())
}

func (recv *ScriptDialog) Native() unsafe.Pointer {
	return recv.native
}

var scriptDialogCloseFunction *gi.Function
var scriptDialogCloseFunction_Once sync.Once

func scriptDialogCloseFunction_Set() error {
	var err error
	scriptDialogCloseFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogCloseFunction, err = scriptDialogStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type webkit_script_dialog_close.
func (recv *ScriptDialog) Close() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := scriptDialogCloseFunction_Set()
	if err == nil {
		scriptDialogCloseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var scriptDialogConfirmSetConfirmedFunction *gi.Function
var scriptDialogConfirmSetConfirmedFunction_Once sync.Once

func scriptDialogConfirmSetConfirmedFunction_Set() error {
	var err error
	scriptDialogConfirmSetConfirmedFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogConfirmSetConfirmedFunction, err = scriptDialogStruct.InvokerNew("confirm_set_confirmed")
	})
	return err
}

// ConfirmSetConfirmed is a representation of the C type webkit_script_dialog_confirm_set_confirmed.
func (recv *ScriptDialog) ConfirmSetConfirmed(confirmed bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(confirmed)

	err := scriptDialogConfirmSetConfirmedFunction_Set()
	if err == nil {
		scriptDialogConfirmSetConfirmedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var scriptDialogGetDialogTypeFunction *gi.Function
var scriptDialogGetDialogTypeFunction_Once sync.Once

func scriptDialogGetDialogTypeFunction_Set() error {
	var err error
	scriptDialogGetDialogTypeFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogGetDialogTypeFunction, err = scriptDialogStruct.InvokerNew("get_dialog_type")
	})
	return err
}

// GetDialogType is a representation of the C type webkit_script_dialog_get_dialog_type.
func (recv *ScriptDialog) GetDialogType() ScriptDialogType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := scriptDialogGetDialogTypeFunction_Set()
	if err == nil {
		ret = scriptDialogGetDialogTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ScriptDialogType(ret.Int32())

	return retGo
}

var scriptDialogGetMessageFunction *gi.Function
var scriptDialogGetMessageFunction_Once sync.Once

func scriptDialogGetMessageFunction_Set() error {
	var err error
	scriptDialogGetMessageFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogGetMessageFunction, err = scriptDialogStruct.InvokerNew("get_message")
	})
	return err
}

// GetMessage is a representation of the C type webkit_script_dialog_get_message.
func (recv *ScriptDialog) GetMessage() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := scriptDialogGetMessageFunction_Set()
	if err == nil {
		ret = scriptDialogGetMessageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var scriptDialogPromptGetDefaultTextFunction *gi.Function
var scriptDialogPromptGetDefaultTextFunction_Once sync.Once

func scriptDialogPromptGetDefaultTextFunction_Set() error {
	var err error
	scriptDialogPromptGetDefaultTextFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogPromptGetDefaultTextFunction, err = scriptDialogStruct.InvokerNew("prompt_get_default_text")
	})
	return err
}

// PromptGetDefaultText is a representation of the C type webkit_script_dialog_prompt_get_default_text.
func (recv *ScriptDialog) PromptGetDefaultText() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := scriptDialogPromptGetDefaultTextFunction_Set()
	if err == nil {
		ret = scriptDialogPromptGetDefaultTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var scriptDialogPromptSetTextFunction *gi.Function
var scriptDialogPromptSetTextFunction_Once sync.Once

func scriptDialogPromptSetTextFunction_Set() error {
	var err error
	scriptDialogPromptSetTextFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogPromptSetTextFunction, err = scriptDialogStruct.InvokerNew("prompt_set_text")
	})
	return err
}

// PromptSetText is a representation of the C type webkit_script_dialog_prompt_set_text.
func (recv *ScriptDialog) PromptSetText(text string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(text)

	err := scriptDialogPromptSetTextFunction_Set()
	if err == nil {
		scriptDialogPromptSetTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var scriptDialogRefFunction *gi.Function
var scriptDialogRefFunction_Once sync.Once

func scriptDialogRefFunction_Set() error {
	var err error
	scriptDialogRefFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogRefFunction, err = scriptDialogStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_script_dialog_ref.
func (recv *ScriptDialog) Ref() *ScriptDialog {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := scriptDialogRefFunction_Set()
	if err == nil {
		ret = scriptDialogRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := ScriptDialogNewFromNative(ret.Pointer())

	return retGo
}

var scriptDialogUnrefFunction *gi.Function
var scriptDialogUnrefFunction_Once sync.Once

func scriptDialogUnrefFunction_Set() error {
	var err error
	scriptDialogUnrefFunction_Once.Do(func() {
		err = scriptDialogStruct_Set()
		if err != nil {
			return
		}
		scriptDialogUnrefFunction, err = scriptDialogStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_script_dialog_unref.
func (recv *ScriptDialog) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := scriptDialogUnrefFunction_Set()
	if err == nil {
		scriptDialogUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// ScriptDialogStruct creates an uninitialised ScriptDialog.
func ScriptDialogStruct() *ScriptDialog {
	err := scriptDialogStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ScriptDialogNewFromNative(scriptDialogStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeScriptDialog)
	return structGo
}
func finalizeScriptDialog(obj *ScriptDialog) {
	scriptDialogStruct.Free(obj.Native())
}

var securityManagerClassStruct *gi.Struct
var securityManagerClassStruct_Once sync.Once

func securityManagerClassStruct_Set() error {
	var err error
	securityManagerClassStruct_Once.Do(func() {
		securityManagerClassStruct, err = gi.StructNew("WebKit2", "SecurityManagerClass")
	})
	return err
}

type SecurityManagerClass struct {
	native unsafe.Pointer
}

func SecurityManagerClassNewFromNative(native unsafe.Pointer) *SecurityManagerClass {
	instance := &SecurityManagerClass{native: native}

	return instance
}

/*
CastToSecurityManagerClass down casts any arbitrary Object to SecurityManagerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SecurityManagerClass.
*/
func CastToSecurityManagerClass(object *gobject.Object) *SecurityManagerClass {
	return SecurityManagerClassNewFromNative(object.Native())
}

func (recv *SecurityManagerClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// SecurityManagerClassStruct creates an uninitialised SecurityManagerClass.
func SecurityManagerClassStruct() *SecurityManagerClass {
	err := securityManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SecurityManagerClassNewFromNative(securityManagerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSecurityManagerClass)
	return structGo
}
func finalizeSecurityManagerClass(obj *SecurityManagerClass) {
	securityManagerClassStruct.Free(obj.Native())
}

var securityManagerPrivateStruct *gi.Struct
var securityManagerPrivateStruct_Once sync.Once

func securityManagerPrivateStruct_Set() error {
	var err error
	securityManagerPrivateStruct_Once.Do(func() {
		securityManagerPrivateStruct, err = gi.StructNew("WebKit2", "SecurityManagerPrivate")
	})
	return err
}

type SecurityManagerPrivate struct {
	native unsafe.Pointer
}

func SecurityManagerPrivateNewFromNative(native unsafe.Pointer) *SecurityManagerPrivate {
	instance := &SecurityManagerPrivate{native: native}

	return instance
}

/*
CastToSecurityManagerPrivate down casts any arbitrary Object to SecurityManagerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SecurityManagerPrivate.
*/
func CastToSecurityManagerPrivate(object *gobject.Object) *SecurityManagerPrivate {
	return SecurityManagerPrivateNewFromNative(object.Native())
}

func (recv *SecurityManagerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SecurityManagerPrivateStruct creates an uninitialised SecurityManagerPrivate.
func SecurityManagerPrivateStruct() *SecurityManagerPrivate {
	err := securityManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SecurityManagerPrivateNewFromNative(securityManagerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSecurityManagerPrivate)
	return structGo
}
func finalizeSecurityManagerPrivate(obj *SecurityManagerPrivate) {
	securityManagerPrivateStruct.Free(obj.Native())
}

var securityOriginStruct *gi.Struct
var securityOriginStruct_Once sync.Once

func securityOriginStruct_Set() error {
	var err error
	securityOriginStruct_Once.Do(func() {
		securityOriginStruct, err = gi.StructNew("WebKit2", "SecurityOrigin")
	})
	return err
}

type SecurityOrigin struct {
	native unsafe.Pointer
}

func SecurityOriginNewFromNative(native unsafe.Pointer) *SecurityOrigin {
	instance := &SecurityOrigin{native: native}

	return instance
}

/*
CastToSecurityOrigin down casts any arbitrary Object to SecurityOrigin.
Exercise care, as this is a potentially dangerous function
if the Object is not a SecurityOrigin.
*/
func CastToSecurityOrigin(object *gobject.Object) *SecurityOrigin {
	return SecurityOriginNewFromNative(object.Native())
}

func (recv *SecurityOrigin) Native() unsafe.Pointer {
	return recv.native
}

var securityOriginNewFunction *gi.Function
var securityOriginNewFunction_Once sync.Once

func securityOriginNewFunction_Set() error {
	var err error
	securityOriginNewFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginNewFunction, err = securityOriginStruct.InvokerNew("new")
	})
	return err
}

// SecurityOriginNew is a representation of the C type webkit_security_origin_new.
func SecurityOriginNew(protocol string, host string, port uint16) *SecurityOrigin {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(protocol)
	inArgs[1].SetString(host)
	inArgs[2].SetUint16(port)

	var ret gi.Argument

	err := securityOriginNewFunction_Set()
	if err == nil {
		ret = securityOriginNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := SecurityOriginNewFromNative(ret.Pointer())

	return retGo
}

var securityOriginNewForUriFunction *gi.Function
var securityOriginNewForUriFunction_Once sync.Once

func securityOriginNewForUriFunction_Set() error {
	var err error
	securityOriginNewForUriFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginNewForUriFunction, err = securityOriginStruct.InvokerNew("new_for_uri")
	})
	return err
}

// SecurityOriginNewForUri is a representation of the C type webkit_security_origin_new_for_uri.
func SecurityOriginNewForUri(uri string) *SecurityOrigin {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(uri)

	var ret gi.Argument

	err := securityOriginNewForUriFunction_Set()
	if err == nil {
		ret = securityOriginNewForUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := SecurityOriginNewFromNative(ret.Pointer())

	return retGo
}

var securityOriginGetHostFunction *gi.Function
var securityOriginGetHostFunction_Once sync.Once

func securityOriginGetHostFunction_Set() error {
	var err error
	securityOriginGetHostFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginGetHostFunction, err = securityOriginStruct.InvokerNew("get_host")
	})
	return err
}

// GetHost is a representation of the C type webkit_security_origin_get_host.
func (recv *SecurityOrigin) GetHost() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := securityOriginGetHostFunction_Set()
	if err == nil {
		ret = securityOriginGetHostFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var securityOriginGetPortFunction *gi.Function
var securityOriginGetPortFunction_Once sync.Once

func securityOriginGetPortFunction_Set() error {
	var err error
	securityOriginGetPortFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginGetPortFunction, err = securityOriginStruct.InvokerNew("get_port")
	})
	return err
}

// GetPort is a representation of the C type webkit_security_origin_get_port.
func (recv *SecurityOrigin) GetPort() uint16 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := securityOriginGetPortFunction_Set()
	if err == nil {
		ret = securityOriginGetPortFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint16()

	return retGo
}

var securityOriginGetProtocolFunction *gi.Function
var securityOriginGetProtocolFunction_Once sync.Once

func securityOriginGetProtocolFunction_Set() error {
	var err error
	securityOriginGetProtocolFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginGetProtocolFunction, err = securityOriginStruct.InvokerNew("get_protocol")
	})
	return err
}

// GetProtocol is a representation of the C type webkit_security_origin_get_protocol.
func (recv *SecurityOrigin) GetProtocol() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := securityOriginGetProtocolFunction_Set()
	if err == nil {
		ret = securityOriginGetProtocolFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var securityOriginIsOpaqueFunction *gi.Function
var securityOriginIsOpaqueFunction_Once sync.Once

func securityOriginIsOpaqueFunction_Set() error {
	var err error
	securityOriginIsOpaqueFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginIsOpaqueFunction, err = securityOriginStruct.InvokerNew("is_opaque")
	})
	return err
}

// IsOpaque is a representation of the C type webkit_security_origin_is_opaque.
func (recv *SecurityOrigin) IsOpaque() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := securityOriginIsOpaqueFunction_Set()
	if err == nil {
		ret = securityOriginIsOpaqueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var securityOriginRefFunction *gi.Function
var securityOriginRefFunction_Once sync.Once

func securityOriginRefFunction_Set() error {
	var err error
	securityOriginRefFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginRefFunction, err = securityOriginStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_security_origin_ref.
func (recv *SecurityOrigin) Ref() *SecurityOrigin {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := securityOriginRefFunction_Set()
	if err == nil {
		ret = securityOriginRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := SecurityOriginNewFromNative(ret.Pointer())

	return retGo
}

var securityOriginToStringFunction *gi.Function
var securityOriginToStringFunction_Once sync.Once

func securityOriginToStringFunction_Set() error {
	var err error
	securityOriginToStringFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginToStringFunction, err = securityOriginStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type webkit_security_origin_to_string.
func (recv *SecurityOrigin) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := securityOriginToStringFunction_Set()
	if err == nil {
		ret = securityOriginToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var securityOriginUnrefFunction *gi.Function
var securityOriginUnrefFunction_Once sync.Once

func securityOriginUnrefFunction_Set() error {
	var err error
	securityOriginUnrefFunction_Once.Do(func() {
		err = securityOriginStruct_Set()
		if err != nil {
			return
		}
		securityOriginUnrefFunction, err = securityOriginStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_security_origin_unref.
func (recv *SecurityOrigin) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := securityOriginUnrefFunction_Set()
	if err == nil {
		securityOriginUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var settingsClassStruct *gi.Struct
var settingsClassStruct_Once sync.Once

func settingsClassStruct_Set() error {
	var err error
	settingsClassStruct_Once.Do(func() {
		settingsClassStruct, err = gi.StructNew("WebKit2", "SettingsClass")
	})
	return err
}

type SettingsClass struct {
	native unsafe.Pointer
}

func SettingsClassNewFromNative(native unsafe.Pointer) *SettingsClass {
	instance := &SettingsClass{native: native}

	return instance
}

/*
CastToSettingsClass down casts any arbitrary Object to SettingsClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsClass.
*/
func CastToSettingsClass(object *gobject.Object) *SettingsClass {
	return SettingsClassNewFromNative(object.Native())
}

func (recv *SettingsClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// SettingsClassStruct creates an uninitialised SettingsClass.
func SettingsClassStruct() *SettingsClass {
	err := settingsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SettingsClassNewFromNative(settingsClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSettingsClass)
	return structGo
}
func finalizeSettingsClass(obj *SettingsClass) {
	settingsClassStruct.Free(obj.Native())
}

var settingsPrivateStruct *gi.Struct
var settingsPrivateStruct_Once sync.Once

func settingsPrivateStruct_Set() error {
	var err error
	settingsPrivateStruct_Once.Do(func() {
		settingsPrivateStruct, err = gi.StructNew("WebKit2", "SettingsPrivate")
	})
	return err
}

type SettingsPrivate struct {
	native unsafe.Pointer
}

func SettingsPrivateNewFromNative(native unsafe.Pointer) *SettingsPrivate {
	instance := &SettingsPrivate{native: native}

	return instance
}

/*
CastToSettingsPrivate down casts any arbitrary Object to SettingsPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a SettingsPrivate.
*/
func CastToSettingsPrivate(object *gobject.Object) *SettingsPrivate {
	return SettingsPrivateNewFromNative(object.Native())
}

func (recv *SettingsPrivate) Native() unsafe.Pointer {
	return recv.native
}

// SettingsPrivateStruct creates an uninitialised SettingsPrivate.
func SettingsPrivateStruct() *SettingsPrivate {
	err := settingsPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SettingsPrivateNewFromNative(settingsPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSettingsPrivate)
	return structGo
}
func finalizeSettingsPrivate(obj *SettingsPrivate) {
	settingsPrivateStruct.Free(obj.Native())
}

var uRIRequestClassStruct *gi.Struct
var uRIRequestClassStruct_Once sync.Once

func uRIRequestClassStruct_Set() error {
	var err error
	uRIRequestClassStruct_Once.Do(func() {
		uRIRequestClassStruct, err = gi.StructNew("WebKit2", "URIRequestClass")
	})
	return err
}

type URIRequestClass struct {
	native unsafe.Pointer
}

func URIRequestClassNewFromNative(native unsafe.Pointer) *URIRequestClass {
	instance := &URIRequestClass{native: native}

	return instance
}

/*
CastToURIRequestClass down casts any arbitrary Object to URIRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a URIRequestClass.
*/
func CastToURIRequestClass(object *gobject.Object) *URIRequestClass {
	return URIRequestClassNewFromNative(object.Native())
}

func (recv *URIRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// URIRequestClassStruct creates an uninitialised URIRequestClass.
func URIRequestClassStruct() *URIRequestClass {
	err := uRIRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := URIRequestClassNewFromNative(uRIRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeURIRequestClass)
	return structGo
}
func finalizeURIRequestClass(obj *URIRequestClass) {
	uRIRequestClassStruct.Free(obj.Native())
}

var uRIRequestPrivateStruct *gi.Struct
var uRIRequestPrivateStruct_Once sync.Once

func uRIRequestPrivateStruct_Set() error {
	var err error
	uRIRequestPrivateStruct_Once.Do(func() {
		uRIRequestPrivateStruct, err = gi.StructNew("WebKit2", "URIRequestPrivate")
	})
	return err
}

type URIRequestPrivate struct {
	native unsafe.Pointer
}

func URIRequestPrivateNewFromNative(native unsafe.Pointer) *URIRequestPrivate {
	instance := &URIRequestPrivate{native: native}

	return instance
}

/*
CastToURIRequestPrivate down casts any arbitrary Object to URIRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a URIRequestPrivate.
*/
func CastToURIRequestPrivate(object *gobject.Object) *URIRequestPrivate {
	return URIRequestPrivateNewFromNative(object.Native())
}

func (recv *URIRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// URIRequestPrivateStruct creates an uninitialised URIRequestPrivate.
func URIRequestPrivateStruct() *URIRequestPrivate {
	err := uRIRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := URIRequestPrivateNewFromNative(uRIRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeURIRequestPrivate)
	return structGo
}
func finalizeURIRequestPrivate(obj *URIRequestPrivate) {
	uRIRequestPrivateStruct.Free(obj.Native())
}

var uRIResponseClassStruct *gi.Struct
var uRIResponseClassStruct_Once sync.Once

func uRIResponseClassStruct_Set() error {
	var err error
	uRIResponseClassStruct_Once.Do(func() {
		uRIResponseClassStruct, err = gi.StructNew("WebKit2", "URIResponseClass")
	})
	return err
}

type URIResponseClass struct {
	native unsafe.Pointer
}

func URIResponseClassNewFromNative(native unsafe.Pointer) *URIResponseClass {
	instance := &URIResponseClass{native: native}

	return instance
}

/*
CastToURIResponseClass down casts any arbitrary Object to URIResponseClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a URIResponseClass.
*/
func CastToURIResponseClass(object *gobject.Object) *URIResponseClass {
	return URIResponseClassNewFromNative(object.Native())
}

func (recv *URIResponseClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// URIResponseClassStruct creates an uninitialised URIResponseClass.
func URIResponseClassStruct() *URIResponseClass {
	err := uRIResponseClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := URIResponseClassNewFromNative(uRIResponseClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeURIResponseClass)
	return structGo
}
func finalizeURIResponseClass(obj *URIResponseClass) {
	uRIResponseClassStruct.Free(obj.Native())
}

var uRIResponsePrivateStruct *gi.Struct
var uRIResponsePrivateStruct_Once sync.Once

func uRIResponsePrivateStruct_Set() error {
	var err error
	uRIResponsePrivateStruct_Once.Do(func() {
		uRIResponsePrivateStruct, err = gi.StructNew("WebKit2", "URIResponsePrivate")
	})
	return err
}

type URIResponsePrivate struct {
	native unsafe.Pointer
}

func URIResponsePrivateNewFromNative(native unsafe.Pointer) *URIResponsePrivate {
	instance := &URIResponsePrivate{native: native}

	return instance
}

/*
CastToURIResponsePrivate down casts any arbitrary Object to URIResponsePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a URIResponsePrivate.
*/
func CastToURIResponsePrivate(object *gobject.Object) *URIResponsePrivate {
	return URIResponsePrivateNewFromNative(object.Native())
}

func (recv *URIResponsePrivate) Native() unsafe.Pointer {
	return recv.native
}

// URIResponsePrivateStruct creates an uninitialised URIResponsePrivate.
func URIResponsePrivateStruct() *URIResponsePrivate {
	err := uRIResponsePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := URIResponsePrivateNewFromNative(uRIResponsePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeURIResponsePrivate)
	return structGo
}
func finalizeURIResponsePrivate(obj *URIResponsePrivate) {
	uRIResponsePrivateStruct.Free(obj.Native())
}

var uRISchemeRequestClassStruct *gi.Struct
var uRISchemeRequestClassStruct_Once sync.Once

func uRISchemeRequestClassStruct_Set() error {
	var err error
	uRISchemeRequestClassStruct_Once.Do(func() {
		uRISchemeRequestClassStruct, err = gi.StructNew("WebKit2", "URISchemeRequestClass")
	})
	return err
}

type URISchemeRequestClass struct {
	native unsafe.Pointer
}

func URISchemeRequestClassNewFromNative(native unsafe.Pointer) *URISchemeRequestClass {
	instance := &URISchemeRequestClass{native: native}

	return instance
}

/*
CastToURISchemeRequestClass down casts any arbitrary Object to URISchemeRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a URISchemeRequestClass.
*/
func CastToURISchemeRequestClass(object *gobject.Object) *URISchemeRequestClass {
	return URISchemeRequestClassNewFromNative(object.Native())
}

func (recv *URISchemeRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// URISchemeRequestClassStruct creates an uninitialised URISchemeRequestClass.
func URISchemeRequestClassStruct() *URISchemeRequestClass {
	err := uRISchemeRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := URISchemeRequestClassNewFromNative(uRISchemeRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeURISchemeRequestClass)
	return structGo
}
func finalizeURISchemeRequestClass(obj *URISchemeRequestClass) {
	uRISchemeRequestClassStruct.Free(obj.Native())
}

var uRISchemeRequestPrivateStruct *gi.Struct
var uRISchemeRequestPrivateStruct_Once sync.Once

func uRISchemeRequestPrivateStruct_Set() error {
	var err error
	uRISchemeRequestPrivateStruct_Once.Do(func() {
		uRISchemeRequestPrivateStruct, err = gi.StructNew("WebKit2", "URISchemeRequestPrivate")
	})
	return err
}

type URISchemeRequestPrivate struct {
	native unsafe.Pointer
}

func URISchemeRequestPrivateNewFromNative(native unsafe.Pointer) *URISchemeRequestPrivate {
	instance := &URISchemeRequestPrivate{native: native}

	return instance
}

/*
CastToURISchemeRequestPrivate down casts any arbitrary Object to URISchemeRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a URISchemeRequestPrivate.
*/
func CastToURISchemeRequestPrivate(object *gobject.Object) *URISchemeRequestPrivate {
	return URISchemeRequestPrivateNewFromNative(object.Native())
}

func (recv *URISchemeRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// URISchemeRequestPrivateStruct creates an uninitialised URISchemeRequestPrivate.
func URISchemeRequestPrivateStruct() *URISchemeRequestPrivate {
	err := uRISchemeRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := URISchemeRequestPrivateNewFromNative(uRISchemeRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeURISchemeRequestPrivate)
	return structGo
}
func finalizeURISchemeRequestPrivate(obj *URISchemeRequestPrivate) {
	uRISchemeRequestPrivateStruct.Free(obj.Native())
}

var userContentFilterStruct *gi.Struct
var userContentFilterStruct_Once sync.Once

func userContentFilterStruct_Set() error {
	var err error
	userContentFilterStruct_Once.Do(func() {
		userContentFilterStruct, err = gi.StructNew("WebKit2", "UserContentFilter")
	})
	return err
}

type UserContentFilter struct {
	native unsafe.Pointer
}

func UserContentFilterNewFromNative(native unsafe.Pointer) *UserContentFilter {
	instance := &UserContentFilter{native: native}

	return instance
}

/*
CastToUserContentFilter down casts any arbitrary Object to UserContentFilter.
Exercise care, as this is a potentially dangerous function
if the Object is not a UserContentFilter.
*/
func CastToUserContentFilter(object *gobject.Object) *UserContentFilter {
	return UserContentFilterNewFromNative(object.Native())
}

func (recv *UserContentFilter) Native() unsafe.Pointer {
	return recv.native
}

var userContentFilterGetIdentifierFunction *gi.Function
var userContentFilterGetIdentifierFunction_Once sync.Once

func userContentFilterGetIdentifierFunction_Set() error {
	var err error
	userContentFilterGetIdentifierFunction_Once.Do(func() {
		err = userContentFilterStruct_Set()
		if err != nil {
			return
		}
		userContentFilterGetIdentifierFunction, err = userContentFilterStruct.InvokerNew("get_identifier")
	})
	return err
}

// GetIdentifier is a representation of the C type webkit_user_content_filter_get_identifier.
func (recv *UserContentFilter) GetIdentifier() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := userContentFilterGetIdentifierFunction_Set()
	if err == nil {
		ret = userContentFilterGetIdentifierFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var userContentFilterRefFunction *gi.Function
var userContentFilterRefFunction_Once sync.Once

func userContentFilterRefFunction_Set() error {
	var err error
	userContentFilterRefFunction_Once.Do(func() {
		err = userContentFilterStruct_Set()
		if err != nil {
			return
		}
		userContentFilterRefFunction, err = userContentFilterStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_user_content_filter_ref.
func (recv *UserContentFilter) Ref() *UserContentFilter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := userContentFilterRefFunction_Set()
	if err == nil {
		ret = userContentFilterRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := UserContentFilterNewFromNative(ret.Pointer())

	return retGo
}

var userContentFilterUnrefFunction *gi.Function
var userContentFilterUnrefFunction_Once sync.Once

func userContentFilterUnrefFunction_Set() error {
	var err error
	userContentFilterUnrefFunction_Once.Do(func() {
		err = userContentFilterStruct_Set()
		if err != nil {
			return
		}
		userContentFilterUnrefFunction, err = userContentFilterStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_user_content_filter_unref.
func (recv *UserContentFilter) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := userContentFilterUnrefFunction_Set()
	if err == nil {
		userContentFilterUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UserContentFilterStruct creates an uninitialised UserContentFilter.
func UserContentFilterStruct() *UserContentFilter {
	err := userContentFilterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UserContentFilterNewFromNative(userContentFilterStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUserContentFilter)
	return structGo
}
func finalizeUserContentFilter(obj *UserContentFilter) {
	userContentFilterStruct.Free(obj.Native())
}

var userContentFilterStoreClassStruct *gi.Struct
var userContentFilterStoreClassStruct_Once sync.Once

func userContentFilterStoreClassStruct_Set() error {
	var err error
	userContentFilterStoreClassStruct_Once.Do(func() {
		userContentFilterStoreClassStruct, err = gi.StructNew("WebKit2", "UserContentFilterStoreClass")
	})
	return err
}

type UserContentFilterStoreClass struct {
	native unsafe.Pointer
}

func UserContentFilterStoreClassNewFromNative(native unsafe.Pointer) *UserContentFilterStoreClass {
	instance := &UserContentFilterStoreClass{native: native}

	return instance
}

/*
CastToUserContentFilterStoreClass down casts any arbitrary Object to UserContentFilterStoreClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UserContentFilterStoreClass.
*/
func CastToUserContentFilterStoreClass(object *gobject.Object) *UserContentFilterStoreClass {
	return UserContentFilterStoreClassNewFromNative(object.Native())
}

func (recv *UserContentFilterStoreClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// UserContentFilterStoreClassStruct creates an uninitialised UserContentFilterStoreClass.
func UserContentFilterStoreClassStruct() *UserContentFilterStoreClass {
	err := userContentFilterStoreClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UserContentFilterStoreClassNewFromNative(userContentFilterStoreClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUserContentFilterStoreClass)
	return structGo
}
func finalizeUserContentFilterStoreClass(obj *UserContentFilterStoreClass) {
	userContentFilterStoreClassStruct.Free(obj.Native())
}

var userContentFilterStorePrivateStruct *gi.Struct
var userContentFilterStorePrivateStruct_Once sync.Once

func userContentFilterStorePrivateStruct_Set() error {
	var err error
	userContentFilterStorePrivateStruct_Once.Do(func() {
		userContentFilterStorePrivateStruct, err = gi.StructNew("WebKit2", "UserContentFilterStorePrivate")
	})
	return err
}

type UserContentFilterStorePrivate struct {
	native unsafe.Pointer
}

func UserContentFilterStorePrivateNewFromNative(native unsafe.Pointer) *UserContentFilterStorePrivate {
	instance := &UserContentFilterStorePrivate{native: native}

	return instance
}

/*
CastToUserContentFilterStorePrivate down casts any arbitrary Object to UserContentFilterStorePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UserContentFilterStorePrivate.
*/
func CastToUserContentFilterStorePrivate(object *gobject.Object) *UserContentFilterStorePrivate {
	return UserContentFilterStorePrivateNewFromNative(object.Native())
}

func (recv *UserContentFilterStorePrivate) Native() unsafe.Pointer {
	return recv.native
}

// UserContentFilterStorePrivateStruct creates an uninitialised UserContentFilterStorePrivate.
func UserContentFilterStorePrivateStruct() *UserContentFilterStorePrivate {
	err := userContentFilterStorePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UserContentFilterStorePrivateNewFromNative(userContentFilterStorePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUserContentFilterStorePrivate)
	return structGo
}
func finalizeUserContentFilterStorePrivate(obj *UserContentFilterStorePrivate) {
	userContentFilterStorePrivateStruct.Free(obj.Native())
}

var userContentManagerClassStruct *gi.Struct
var userContentManagerClassStruct_Once sync.Once

func userContentManagerClassStruct_Set() error {
	var err error
	userContentManagerClassStruct_Once.Do(func() {
		userContentManagerClassStruct, err = gi.StructNew("WebKit2", "UserContentManagerClass")
	})
	return err
}

type UserContentManagerClass struct {
	native unsafe.Pointer
}

func UserContentManagerClassNewFromNative(native unsafe.Pointer) *UserContentManagerClass {
	instance := &UserContentManagerClass{native: native}

	return instance
}

/*
CastToUserContentManagerClass down casts any arbitrary Object to UserContentManagerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UserContentManagerClass.
*/
func CastToUserContentManagerClass(object *gobject.Object) *UserContentManagerClass {
	return UserContentManagerClassNewFromNative(object.Native())
}

func (recv *UserContentManagerClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// UserContentManagerClassStruct creates an uninitialised UserContentManagerClass.
func UserContentManagerClassStruct() *UserContentManagerClass {
	err := userContentManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UserContentManagerClassNewFromNative(userContentManagerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUserContentManagerClass)
	return structGo
}
func finalizeUserContentManagerClass(obj *UserContentManagerClass) {
	userContentManagerClassStruct.Free(obj.Native())
}

var userContentManagerPrivateStruct *gi.Struct
var userContentManagerPrivateStruct_Once sync.Once

func userContentManagerPrivateStruct_Set() error {
	var err error
	userContentManagerPrivateStruct_Once.Do(func() {
		userContentManagerPrivateStruct, err = gi.StructNew("WebKit2", "UserContentManagerPrivate")
	})
	return err
}

type UserContentManagerPrivate struct {
	native unsafe.Pointer
}

func UserContentManagerPrivateNewFromNative(native unsafe.Pointer) *UserContentManagerPrivate {
	instance := &UserContentManagerPrivate{native: native}

	return instance
}

/*
CastToUserContentManagerPrivate down casts any arbitrary Object to UserContentManagerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UserContentManagerPrivate.
*/
func CastToUserContentManagerPrivate(object *gobject.Object) *UserContentManagerPrivate {
	return UserContentManagerPrivateNewFromNative(object.Native())
}

func (recv *UserContentManagerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// UserContentManagerPrivateStruct creates an uninitialised UserContentManagerPrivate.
func UserContentManagerPrivateStruct() *UserContentManagerPrivate {
	err := userContentManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UserContentManagerPrivateNewFromNative(userContentManagerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUserContentManagerPrivate)
	return structGo
}
func finalizeUserContentManagerPrivate(obj *UserContentManagerPrivate) {
	userContentManagerPrivateStruct.Free(obj.Native())
}

var userMediaPermissionRequestClassStruct *gi.Struct
var userMediaPermissionRequestClassStruct_Once sync.Once

func userMediaPermissionRequestClassStruct_Set() error {
	var err error
	userMediaPermissionRequestClassStruct_Once.Do(func() {
		userMediaPermissionRequestClassStruct, err = gi.StructNew("WebKit2", "UserMediaPermissionRequestClass")
	})
	return err
}

type UserMediaPermissionRequestClass struct {
	native unsafe.Pointer
}

func UserMediaPermissionRequestClassNewFromNative(native unsafe.Pointer) *UserMediaPermissionRequestClass {
	instance := &UserMediaPermissionRequestClass{native: native}

	return instance
}

/*
CastToUserMediaPermissionRequestClass down casts any arbitrary Object to UserMediaPermissionRequestClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UserMediaPermissionRequestClass.
*/
func CastToUserMediaPermissionRequestClass(object *gobject.Object) *UserMediaPermissionRequestClass {
	return UserMediaPermissionRequestClassNewFromNative(object.Native())
}

func (recv *UserMediaPermissionRequestClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// UserMediaPermissionRequestClassStruct creates an uninitialised UserMediaPermissionRequestClass.
func UserMediaPermissionRequestClassStruct() *UserMediaPermissionRequestClass {
	err := userMediaPermissionRequestClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UserMediaPermissionRequestClassNewFromNative(userMediaPermissionRequestClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUserMediaPermissionRequestClass)
	return structGo
}
func finalizeUserMediaPermissionRequestClass(obj *UserMediaPermissionRequestClass) {
	userMediaPermissionRequestClassStruct.Free(obj.Native())
}

var userMediaPermissionRequestPrivateStruct *gi.Struct
var userMediaPermissionRequestPrivateStruct_Once sync.Once

func userMediaPermissionRequestPrivateStruct_Set() error {
	var err error
	userMediaPermissionRequestPrivateStruct_Once.Do(func() {
		userMediaPermissionRequestPrivateStruct, err = gi.StructNew("WebKit2", "UserMediaPermissionRequestPrivate")
	})
	return err
}

type UserMediaPermissionRequestPrivate struct {
	native unsafe.Pointer
}

func UserMediaPermissionRequestPrivateNewFromNative(native unsafe.Pointer) *UserMediaPermissionRequestPrivate {
	instance := &UserMediaPermissionRequestPrivate{native: native}

	return instance
}

/*
CastToUserMediaPermissionRequestPrivate down casts any arbitrary Object to UserMediaPermissionRequestPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a UserMediaPermissionRequestPrivate.
*/
func CastToUserMediaPermissionRequestPrivate(object *gobject.Object) *UserMediaPermissionRequestPrivate {
	return UserMediaPermissionRequestPrivateNewFromNative(object.Native())
}

func (recv *UserMediaPermissionRequestPrivate) Native() unsafe.Pointer {
	return recv.native
}

// UserMediaPermissionRequestPrivateStruct creates an uninitialised UserMediaPermissionRequestPrivate.
func UserMediaPermissionRequestPrivateStruct() *UserMediaPermissionRequestPrivate {
	err := userMediaPermissionRequestPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UserMediaPermissionRequestPrivateNewFromNative(userMediaPermissionRequestPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUserMediaPermissionRequestPrivate)
	return structGo
}
func finalizeUserMediaPermissionRequestPrivate(obj *UserMediaPermissionRequestPrivate) {
	userMediaPermissionRequestPrivateStruct.Free(obj.Native())
}

var userScriptStruct *gi.Struct
var userScriptStruct_Once sync.Once

func userScriptStruct_Set() error {
	var err error
	userScriptStruct_Once.Do(func() {
		userScriptStruct, err = gi.StructNew("WebKit2", "UserScript")
	})
	return err
}

type UserScript struct {
	native unsafe.Pointer
}

func UserScriptNewFromNative(native unsafe.Pointer) *UserScript {
	instance := &UserScript{native: native}

	return instance
}

/*
CastToUserScript down casts any arbitrary Object to UserScript.
Exercise care, as this is a potentially dangerous function
if the Object is not a UserScript.
*/
func CastToUserScript(object *gobject.Object) *UserScript {
	return UserScriptNewFromNative(object.Native())
}

func (recv *UserScript) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'webkit_user_script_new' : parameter 'whitelist' of type 'nil' not supported

// UNSUPPORTED : C value 'webkit_user_script_new_for_world' : parameter 'whitelist' of type 'nil' not supported

var userScriptRefFunction *gi.Function
var userScriptRefFunction_Once sync.Once

func userScriptRefFunction_Set() error {
	var err error
	userScriptRefFunction_Once.Do(func() {
		err = userScriptStruct_Set()
		if err != nil {
			return
		}
		userScriptRefFunction, err = userScriptStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_user_script_ref.
func (recv *UserScript) Ref() *UserScript {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := userScriptRefFunction_Set()
	if err == nil {
		ret = userScriptRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := UserScriptNewFromNative(ret.Pointer())

	return retGo
}

var userScriptUnrefFunction *gi.Function
var userScriptUnrefFunction_Once sync.Once

func userScriptUnrefFunction_Set() error {
	var err error
	userScriptUnrefFunction_Once.Do(func() {
		err = userScriptStruct_Set()
		if err != nil {
			return
		}
		userScriptUnrefFunction, err = userScriptStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_user_script_unref.
func (recv *UserScript) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := userScriptUnrefFunction_Set()
	if err == nil {
		userScriptUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var userStyleSheetStruct *gi.Struct
var userStyleSheetStruct_Once sync.Once

func userStyleSheetStruct_Set() error {
	var err error
	userStyleSheetStruct_Once.Do(func() {
		userStyleSheetStruct, err = gi.StructNew("WebKit2", "UserStyleSheet")
	})
	return err
}

type UserStyleSheet struct {
	native unsafe.Pointer
}

func UserStyleSheetNewFromNative(native unsafe.Pointer) *UserStyleSheet {
	instance := &UserStyleSheet{native: native}

	return instance
}

/*
CastToUserStyleSheet down casts any arbitrary Object to UserStyleSheet.
Exercise care, as this is a potentially dangerous function
if the Object is not a UserStyleSheet.
*/
func CastToUserStyleSheet(object *gobject.Object) *UserStyleSheet {
	return UserStyleSheetNewFromNative(object.Native())
}

func (recv *UserStyleSheet) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'webkit_user_style_sheet_new' : parameter 'whitelist' of type 'nil' not supported

// UNSUPPORTED : C value 'webkit_user_style_sheet_new_for_world' : parameter 'whitelist' of type 'nil' not supported

var userStyleSheetRefFunction *gi.Function
var userStyleSheetRefFunction_Once sync.Once

func userStyleSheetRefFunction_Set() error {
	var err error
	userStyleSheetRefFunction_Once.Do(func() {
		err = userStyleSheetStruct_Set()
		if err != nil {
			return
		}
		userStyleSheetRefFunction, err = userStyleSheetStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_user_style_sheet_ref.
func (recv *UserStyleSheet) Ref() *UserStyleSheet {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := userStyleSheetRefFunction_Set()
	if err == nil {
		ret = userStyleSheetRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := UserStyleSheetNewFromNative(ret.Pointer())

	return retGo
}

var userStyleSheetUnrefFunction *gi.Function
var userStyleSheetUnrefFunction_Once sync.Once

func userStyleSheetUnrefFunction_Set() error {
	var err error
	userStyleSheetUnrefFunction_Once.Do(func() {
		err = userStyleSheetStruct_Set()
		if err != nil {
			return
		}
		userStyleSheetUnrefFunction, err = userStyleSheetStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_user_style_sheet_unref.
func (recv *UserStyleSheet) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := userStyleSheetUnrefFunction_Set()
	if err == nil {
		userStyleSheetUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var webContextClassStruct *gi.Struct
var webContextClassStruct_Once sync.Once

func webContextClassStruct_Set() error {
	var err error
	webContextClassStruct_Once.Do(func() {
		webContextClassStruct, err = gi.StructNew("WebKit2", "WebContextClass")
	})
	return err
}

type WebContextClass struct {
	native unsafe.Pointer
}

func WebContextClassNewFromNative(native unsafe.Pointer) *WebContextClass {
	instance := &WebContextClass{native: native}

	return instance
}

/*
CastToWebContextClass down casts any arbitrary Object to WebContextClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebContextClass.
*/
func CastToWebContextClass(object *gobject.Object) *WebContextClass {
	return WebContextClassNewFromNative(object.Native())
}

func (recv *WebContextClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'download_started' : for field getter : missing Type

// UNSUPPORTED : C value 'download_started' : for field setter : missing Type

// UNSUPPORTED : C value 'initialize_web_extensions' : for field getter : missing Type

// UNSUPPORTED : C value 'initialize_web_extensions' : for field setter : missing Type

// UNSUPPORTED : C value 'initialize_notification_permissions' : for field getter : missing Type

// UNSUPPORTED : C value 'initialize_notification_permissions' : for field setter : missing Type

// UNSUPPORTED : C value 'automation_started' : for field getter : missing Type

// UNSUPPORTED : C value 'automation_started' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// WebContextClassStruct creates an uninitialised WebContextClass.
func WebContextClassStruct() *WebContextClass {
	err := webContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebContextClassNewFromNative(webContextClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebContextClass)
	return structGo
}
func finalizeWebContextClass(obj *WebContextClass) {
	webContextClassStruct.Free(obj.Native())
}

var webContextPrivateStruct *gi.Struct
var webContextPrivateStruct_Once sync.Once

func webContextPrivateStruct_Set() error {
	var err error
	webContextPrivateStruct_Once.Do(func() {
		webContextPrivateStruct, err = gi.StructNew("WebKit2", "WebContextPrivate")
	})
	return err
}

type WebContextPrivate struct {
	native unsafe.Pointer
}

func WebContextPrivateNewFromNative(native unsafe.Pointer) *WebContextPrivate {
	instance := &WebContextPrivate{native: native}

	return instance
}

/*
CastToWebContextPrivate down casts any arbitrary Object to WebContextPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebContextPrivate.
*/
func CastToWebContextPrivate(object *gobject.Object) *WebContextPrivate {
	return WebContextPrivateNewFromNative(object.Native())
}

func (recv *WebContextPrivate) Native() unsafe.Pointer {
	return recv.native
}

// WebContextPrivateStruct creates an uninitialised WebContextPrivate.
func WebContextPrivateStruct() *WebContextPrivate {
	err := webContextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebContextPrivateNewFromNative(webContextPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebContextPrivate)
	return structGo
}
func finalizeWebContextPrivate(obj *WebContextPrivate) {
	webContextPrivateStruct.Free(obj.Native())
}

var webInspectorClassStruct *gi.Struct
var webInspectorClassStruct_Once sync.Once

func webInspectorClassStruct_Set() error {
	var err error
	webInspectorClassStruct_Once.Do(func() {
		webInspectorClassStruct, err = gi.StructNew("WebKit2", "WebInspectorClass")
	})
	return err
}

type WebInspectorClass struct {
	native unsafe.Pointer
}

func WebInspectorClassNewFromNative(native unsafe.Pointer) *WebInspectorClass {
	instance := &WebInspectorClass{native: native}

	return instance
}

/*
CastToWebInspectorClass down casts any arbitrary Object to WebInspectorClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebInspectorClass.
*/
func CastToWebInspectorClass(object *gobject.Object) *WebInspectorClass {
	return WebInspectorClassNewFromNative(object.Native())
}

func (recv *WebInspectorClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// WebInspectorClassStruct creates an uninitialised WebInspectorClass.
func WebInspectorClassStruct() *WebInspectorClass {
	err := webInspectorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebInspectorClassNewFromNative(webInspectorClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebInspectorClass)
	return structGo
}
func finalizeWebInspectorClass(obj *WebInspectorClass) {
	webInspectorClassStruct.Free(obj.Native())
}

var webInspectorPrivateStruct *gi.Struct
var webInspectorPrivateStruct_Once sync.Once

func webInspectorPrivateStruct_Set() error {
	var err error
	webInspectorPrivateStruct_Once.Do(func() {
		webInspectorPrivateStruct, err = gi.StructNew("WebKit2", "WebInspectorPrivate")
	})
	return err
}

type WebInspectorPrivate struct {
	native unsafe.Pointer
}

func WebInspectorPrivateNewFromNative(native unsafe.Pointer) *WebInspectorPrivate {
	instance := &WebInspectorPrivate{native: native}

	return instance
}

/*
CastToWebInspectorPrivate down casts any arbitrary Object to WebInspectorPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebInspectorPrivate.
*/
func CastToWebInspectorPrivate(object *gobject.Object) *WebInspectorPrivate {
	return WebInspectorPrivateNewFromNative(object.Native())
}

func (recv *WebInspectorPrivate) Native() unsafe.Pointer {
	return recv.native
}

// WebInspectorPrivateStruct creates an uninitialised WebInspectorPrivate.
func WebInspectorPrivateStruct() *WebInspectorPrivate {
	err := webInspectorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebInspectorPrivateNewFromNative(webInspectorPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebInspectorPrivate)
	return structGo
}
func finalizeWebInspectorPrivate(obj *WebInspectorPrivate) {
	webInspectorPrivateStruct.Free(obj.Native())
}

var webResourceClassStruct *gi.Struct
var webResourceClassStruct_Once sync.Once

func webResourceClassStruct_Set() error {
	var err error
	webResourceClassStruct_Once.Do(func() {
		webResourceClassStruct, err = gi.StructNew("WebKit2", "WebResourceClass")
	})
	return err
}

type WebResourceClass struct {
	native unsafe.Pointer
}

func WebResourceClassNewFromNative(native unsafe.Pointer) *WebResourceClass {
	instance := &WebResourceClass{native: native}

	return instance
}

/*
CastToWebResourceClass down casts any arbitrary Object to WebResourceClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebResourceClass.
*/
func CastToWebResourceClass(object *gobject.Object) *WebResourceClass {
	return WebResourceClassNewFromNative(object.Native())
}

func (recv *WebResourceClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// WebResourceClassStruct creates an uninitialised WebResourceClass.
func WebResourceClassStruct() *WebResourceClass {
	err := webResourceClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebResourceClassNewFromNative(webResourceClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebResourceClass)
	return structGo
}
func finalizeWebResourceClass(obj *WebResourceClass) {
	webResourceClassStruct.Free(obj.Native())
}

var webResourcePrivateStruct *gi.Struct
var webResourcePrivateStruct_Once sync.Once

func webResourcePrivateStruct_Set() error {
	var err error
	webResourcePrivateStruct_Once.Do(func() {
		webResourcePrivateStruct, err = gi.StructNew("WebKit2", "WebResourcePrivate")
	})
	return err
}

type WebResourcePrivate struct {
	native unsafe.Pointer
}

func WebResourcePrivateNewFromNative(native unsafe.Pointer) *WebResourcePrivate {
	instance := &WebResourcePrivate{native: native}

	return instance
}

/*
CastToWebResourcePrivate down casts any arbitrary Object to WebResourcePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebResourcePrivate.
*/
func CastToWebResourcePrivate(object *gobject.Object) *WebResourcePrivate {
	return WebResourcePrivateNewFromNative(object.Native())
}

func (recv *WebResourcePrivate) Native() unsafe.Pointer {
	return recv.native
}

// WebResourcePrivateStruct creates an uninitialised WebResourcePrivate.
func WebResourcePrivateStruct() *WebResourcePrivate {
	err := webResourcePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebResourcePrivateNewFromNative(webResourcePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebResourcePrivate)
	return structGo
}
func finalizeWebResourcePrivate(obj *WebResourcePrivate) {
	webResourcePrivateStruct.Free(obj.Native())
}

var webViewBaseClassStruct *gi.Struct
var webViewBaseClassStruct_Once sync.Once

func webViewBaseClassStruct_Set() error {
	var err error
	webViewBaseClassStruct_Once.Do(func() {
		webViewBaseClassStruct, err = gi.StructNew("WebKit2", "WebViewBaseClass")
	})
	return err
}

type WebViewBaseClass struct {
	native unsafe.Pointer
}

func WebViewBaseClassNewFromNative(native unsafe.Pointer) *WebViewBaseClass {
	instance := &WebViewBaseClass{native: native}

	return instance
}

/*
CastToWebViewBaseClass down casts any arbitrary Object to WebViewBaseClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebViewBaseClass.
*/
func CastToWebViewBaseClass(object *gobject.Object) *WebViewBaseClass {
	return WebViewBaseClassNewFromNative(object.Native())
}

func (recv *WebViewBaseClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parentClass' : for field getter : no Go type for 'Gtk.ContainerClass'

// UNSUPPORTED : C value 'parentClass' : for field setter : no Go type for 'Gtk.ContainerClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// WebViewBaseClassStruct creates an uninitialised WebViewBaseClass.
func WebViewBaseClassStruct() *WebViewBaseClass {
	err := webViewBaseClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebViewBaseClassNewFromNative(webViewBaseClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebViewBaseClass)
	return structGo
}
func finalizeWebViewBaseClass(obj *WebViewBaseClass) {
	webViewBaseClassStruct.Free(obj.Native())
}

var webViewBasePrivateStruct *gi.Struct
var webViewBasePrivateStruct_Once sync.Once

func webViewBasePrivateStruct_Set() error {
	var err error
	webViewBasePrivateStruct_Once.Do(func() {
		webViewBasePrivateStruct, err = gi.StructNew("WebKit2", "WebViewBasePrivate")
	})
	return err
}

type WebViewBasePrivate struct {
	native unsafe.Pointer
}

func WebViewBasePrivateNewFromNative(native unsafe.Pointer) *WebViewBasePrivate {
	instance := &WebViewBasePrivate{native: native}

	return instance
}

/*
CastToWebViewBasePrivate down casts any arbitrary Object to WebViewBasePrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebViewBasePrivate.
*/
func CastToWebViewBasePrivate(object *gobject.Object) *WebViewBasePrivate {
	return WebViewBasePrivateNewFromNative(object.Native())
}

func (recv *WebViewBasePrivate) Native() unsafe.Pointer {
	return recv.native
}

// WebViewBasePrivateStruct creates an uninitialised WebViewBasePrivate.
func WebViewBasePrivateStruct() *WebViewBasePrivate {
	err := webViewBasePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebViewBasePrivateNewFromNative(webViewBasePrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebViewBasePrivate)
	return structGo
}
func finalizeWebViewBasePrivate(obj *WebViewBasePrivate) {
	webViewBasePrivateStruct.Free(obj.Native())
}

var webViewClassStruct *gi.Struct
var webViewClassStruct_Once sync.Once

func webViewClassStruct_Set() error {
	var err error
	webViewClassStruct_Once.Do(func() {
		webViewClassStruct, err = gi.StructNew("WebKit2", "WebViewClass")
	})
	return err
}

type WebViewClass struct {
	native unsafe.Pointer
}

func WebViewClassNewFromNative(native unsafe.Pointer) *WebViewClass {
	instance := &WebViewClass{native: native}

	return instance
}

/*
CastToWebViewClass down casts any arbitrary Object to WebViewClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebViewClass.
*/
func CastToWebViewClass(object *gobject.Object) *WebViewClass {
	return WebViewClassNewFromNative(object.Native())
}

func (recv *WebViewClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *WebViewClass) FieldParent() *WebViewBaseClass {
	argValue := gi.StructFieldGet(webViewClassStruct, recv.Native(), "parent")
	value := WebViewBaseClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *WebViewClass) SetFieldParent(value *WebViewBaseClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(webViewClassStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'load_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'load_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'load_failed' : for field getter : missing Type

// UNSUPPORTED : C value 'load_failed' : for field setter : missing Type

// UNSUPPORTED : C value 'create' : for field getter : missing Type

// UNSUPPORTED : C value 'create' : for field setter : missing Type

// UNSUPPORTED : C value 'ready_to_show' : for field getter : missing Type

// UNSUPPORTED : C value 'ready_to_show' : for field setter : missing Type

// UNSUPPORTED : C value 'run_as_modal' : for field getter : missing Type

// UNSUPPORTED : C value 'run_as_modal' : for field setter : missing Type

// UNSUPPORTED : C value 'close' : for field getter : missing Type

// UNSUPPORTED : C value 'close' : for field setter : missing Type

// UNSUPPORTED : C value 'script_dialog' : for field getter : missing Type

// UNSUPPORTED : C value 'script_dialog' : for field setter : missing Type

// UNSUPPORTED : C value 'decide_policy' : for field getter : missing Type

// UNSUPPORTED : C value 'decide_policy' : for field setter : missing Type

// UNSUPPORTED : C value 'permission_request' : for field getter : missing Type

// UNSUPPORTED : C value 'permission_request' : for field setter : missing Type

// UNSUPPORTED : C value 'mouse_target_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'mouse_target_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'print' : for field getter : missing Type

// UNSUPPORTED : C value 'print' : for field setter : missing Type

// UNSUPPORTED : C value 'resource_load_started' : for field getter : missing Type

// UNSUPPORTED : C value 'resource_load_started' : for field setter : missing Type

// UNSUPPORTED : C value 'enter_fullscreen' : for field getter : missing Type

// UNSUPPORTED : C value 'enter_fullscreen' : for field setter : missing Type

// UNSUPPORTED : C value 'leave_fullscreen' : for field getter : missing Type

// UNSUPPORTED : C value 'leave_fullscreen' : for field setter : missing Type

// UNSUPPORTED : C value 'run_file_chooser' : for field getter : missing Type

// UNSUPPORTED : C value 'run_file_chooser' : for field setter : missing Type

// UNSUPPORTED : C value 'context_menu' : for field getter : missing Type

// UNSUPPORTED : C value 'context_menu' : for field setter : missing Type

// UNSUPPORTED : C value 'context_menu_dismissed' : for field getter : missing Type

// UNSUPPORTED : C value 'context_menu_dismissed' : for field setter : missing Type

// UNSUPPORTED : C value 'submit_form' : for field getter : missing Type

// UNSUPPORTED : C value 'submit_form' : for field setter : missing Type

// UNSUPPORTED : C value 'insecure_content_detected' : for field getter : missing Type

// UNSUPPORTED : C value 'insecure_content_detected' : for field setter : missing Type

// UNSUPPORTED : C value 'web_process_crashed' : for field getter : missing Type

// UNSUPPORTED : C value 'web_process_crashed' : for field setter : missing Type

// UNSUPPORTED : C value 'authenticate' : for field getter : missing Type

// UNSUPPORTED : C value 'authenticate' : for field setter : missing Type

// UNSUPPORTED : C value 'load_failed_with_tls_errors' : for field getter : missing Type

// UNSUPPORTED : C value 'load_failed_with_tls_errors' : for field setter : missing Type

// UNSUPPORTED : C value 'show_notification' : for field getter : missing Type

// UNSUPPORTED : C value 'show_notification' : for field setter : missing Type

// UNSUPPORTED : C value 'run_color_chooser' : for field getter : missing Type

// UNSUPPORTED : C value 'run_color_chooser' : for field setter : missing Type

// UNSUPPORTED : C value 'show_option_menu' : for field getter : missing Type

// UNSUPPORTED : C value 'show_option_menu' : for field setter : missing Type

// UNSUPPORTED : C value 'web_process_terminated' : for field getter : missing Type

// UNSUPPORTED : C value 'web_process_terminated' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// WebViewClassStruct creates an uninitialised WebViewClass.
func WebViewClassStruct() *WebViewClass {
	err := webViewClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebViewClassNewFromNative(webViewClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebViewClass)
	return structGo
}
func finalizeWebViewClass(obj *WebViewClass) {
	webViewClassStruct.Free(obj.Native())
}

var webViewPrivateStruct *gi.Struct
var webViewPrivateStruct_Once sync.Once

func webViewPrivateStruct_Set() error {
	var err error
	webViewPrivateStruct_Once.Do(func() {
		webViewPrivateStruct, err = gi.StructNew("WebKit2", "WebViewPrivate")
	})
	return err
}

type WebViewPrivate struct {
	native unsafe.Pointer
}

func WebViewPrivateNewFromNative(native unsafe.Pointer) *WebViewPrivate {
	instance := &WebViewPrivate{native: native}

	return instance
}

/*
CastToWebViewPrivate down casts any arbitrary Object to WebViewPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebViewPrivate.
*/
func CastToWebViewPrivate(object *gobject.Object) *WebViewPrivate {
	return WebViewPrivateNewFromNative(object.Native())
}

func (recv *WebViewPrivate) Native() unsafe.Pointer {
	return recv.native
}

// WebViewPrivateStruct creates an uninitialised WebViewPrivate.
func WebViewPrivateStruct() *WebViewPrivate {
	err := webViewPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebViewPrivateNewFromNative(webViewPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebViewPrivate)
	return structGo
}
func finalizeWebViewPrivate(obj *WebViewPrivate) {
	webViewPrivateStruct.Free(obj.Native())
}

var webViewSessionStateStruct *gi.Struct
var webViewSessionStateStruct_Once sync.Once

func webViewSessionStateStruct_Set() error {
	var err error
	webViewSessionStateStruct_Once.Do(func() {
		webViewSessionStateStruct, err = gi.StructNew("WebKit2", "WebViewSessionState")
	})
	return err
}

type WebViewSessionState struct {
	native unsafe.Pointer
}

func WebViewSessionStateNewFromNative(native unsafe.Pointer) *WebViewSessionState {
	instance := &WebViewSessionState{native: native}

	return instance
}

/*
CastToWebViewSessionState down casts any arbitrary Object to WebViewSessionState.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebViewSessionState.
*/
func CastToWebViewSessionState(object *gobject.Object) *WebViewSessionState {
	return WebViewSessionStateNewFromNative(object.Native())
}

func (recv *WebViewSessionState) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'webkit_web_view_session_state_new' : parameter 'data' of type 'GLib.Bytes' not supported

var webViewSessionStateRefFunction *gi.Function
var webViewSessionStateRefFunction_Once sync.Once

func webViewSessionStateRefFunction_Set() error {
	var err error
	webViewSessionStateRefFunction_Once.Do(func() {
		err = webViewSessionStateStruct_Set()
		if err != nil {
			return
		}
		webViewSessionStateRefFunction, err = webViewSessionStateStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_web_view_session_state_ref.
func (recv *WebViewSessionState) Ref() *WebViewSessionState {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := webViewSessionStateRefFunction_Set()
	if err == nil {
		ret = webViewSessionStateRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := WebViewSessionStateNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'webkit_web_view_session_state_serialize' : return type 'GLib.Bytes' not supported

var webViewSessionStateUnrefFunction *gi.Function
var webViewSessionStateUnrefFunction_Once sync.Once

func webViewSessionStateUnrefFunction_Set() error {
	var err error
	webViewSessionStateUnrefFunction_Once.Do(func() {
		err = webViewSessionStateStruct_Set()
		if err != nil {
			return
		}
		webViewSessionStateUnrefFunction, err = webViewSessionStateStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_web_view_session_state_unref.
func (recv *WebViewSessionState) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := webViewSessionStateUnrefFunction_Set()
	if err == nil {
		webViewSessionStateUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var websiteDataStruct *gi.Struct
var websiteDataStruct_Once sync.Once

func websiteDataStruct_Set() error {
	var err error
	websiteDataStruct_Once.Do(func() {
		websiteDataStruct, err = gi.StructNew("WebKit2", "WebsiteData")
	})
	return err
}

type WebsiteData struct {
	native unsafe.Pointer
}

func WebsiteDataNewFromNative(native unsafe.Pointer) *WebsiteData {
	instance := &WebsiteData{native: native}

	return instance
}

/*
CastToWebsiteData down casts any arbitrary Object to WebsiteData.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsiteData.
*/
func CastToWebsiteData(object *gobject.Object) *WebsiteData {
	return WebsiteDataNewFromNative(object.Native())
}

func (recv *WebsiteData) Native() unsafe.Pointer {
	return recv.native
}

var websiteDataGetNameFunction *gi.Function
var websiteDataGetNameFunction_Once sync.Once

func websiteDataGetNameFunction_Set() error {
	var err error
	websiteDataGetNameFunction_Once.Do(func() {
		err = websiteDataStruct_Set()
		if err != nil {
			return
		}
		websiteDataGetNameFunction, err = websiteDataStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type webkit_website_data_get_name.
func (recv *WebsiteData) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websiteDataGetNameFunction_Set()
	if err == nil {
		ret = websiteDataGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'webkit_website_data_get_size' : parameter 'types' of type 'WebsiteDataTypes' not supported

// UNSUPPORTED : C value 'webkit_website_data_get_types' : return type 'WebsiteDataTypes' not supported

var websiteDataRefFunction *gi.Function
var websiteDataRefFunction_Once sync.Once

func websiteDataRefFunction_Set() error {
	var err error
	websiteDataRefFunction_Once.Do(func() {
		err = websiteDataStruct_Set()
		if err != nil {
			return
		}
		websiteDataRefFunction, err = websiteDataStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type webkit_website_data_ref.
func (recv *WebsiteData) Ref() *WebsiteData {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := websiteDataRefFunction_Set()
	if err == nil {
		ret = websiteDataRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := WebsiteDataNewFromNative(ret.Pointer())

	return retGo
}

var websiteDataUnrefFunction *gi.Function
var websiteDataUnrefFunction_Once sync.Once

func websiteDataUnrefFunction_Set() error {
	var err error
	websiteDataUnrefFunction_Once.Do(func() {
		err = websiteDataStruct_Set()
		if err != nil {
			return
		}
		websiteDataUnrefFunction, err = websiteDataStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type webkit_website_data_unref.
func (recv *WebsiteData) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := websiteDataUnrefFunction_Set()
	if err == nil {
		websiteDataUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// WebsiteDataStruct creates an uninitialised WebsiteData.
func WebsiteDataStruct() *WebsiteData {
	err := websiteDataStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebsiteDataNewFromNative(websiteDataStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebsiteData)
	return structGo
}
func finalizeWebsiteData(obj *WebsiteData) {
	websiteDataStruct.Free(obj.Native())
}

var websiteDataManagerClassStruct *gi.Struct
var websiteDataManagerClassStruct_Once sync.Once

func websiteDataManagerClassStruct_Set() error {
	var err error
	websiteDataManagerClassStruct_Once.Do(func() {
		websiteDataManagerClassStruct, err = gi.StructNew("WebKit2", "WebsiteDataManagerClass")
	})
	return err
}

type WebsiteDataManagerClass struct {
	native unsafe.Pointer
}

func WebsiteDataManagerClassNewFromNative(native unsafe.Pointer) *WebsiteDataManagerClass {
	instance := &WebsiteDataManagerClass{native: native}

	return instance
}

/*
CastToWebsiteDataManagerClass down casts any arbitrary Object to WebsiteDataManagerClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsiteDataManagerClass.
*/
func CastToWebsiteDataManagerClass(object *gobject.Object) *WebsiteDataManagerClass {
	return WebsiteDataManagerClassNewFromNative(object.Native())
}

func (recv *WebsiteDataManagerClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// WebsiteDataManagerClassStruct creates an uninitialised WebsiteDataManagerClass.
func WebsiteDataManagerClassStruct() *WebsiteDataManagerClass {
	err := websiteDataManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebsiteDataManagerClassNewFromNative(websiteDataManagerClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebsiteDataManagerClass)
	return structGo
}
func finalizeWebsiteDataManagerClass(obj *WebsiteDataManagerClass) {
	websiteDataManagerClassStruct.Free(obj.Native())
}

var websiteDataManagerPrivateStruct *gi.Struct
var websiteDataManagerPrivateStruct_Once sync.Once

func websiteDataManagerPrivateStruct_Set() error {
	var err error
	websiteDataManagerPrivateStruct_Once.Do(func() {
		websiteDataManagerPrivateStruct, err = gi.StructNew("WebKit2", "WebsiteDataManagerPrivate")
	})
	return err
}

type WebsiteDataManagerPrivate struct {
	native unsafe.Pointer
}

func WebsiteDataManagerPrivateNewFromNative(native unsafe.Pointer) *WebsiteDataManagerPrivate {
	instance := &WebsiteDataManagerPrivate{native: native}

	return instance
}

/*
CastToWebsiteDataManagerPrivate down casts any arbitrary Object to WebsiteDataManagerPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a WebsiteDataManagerPrivate.
*/
func CastToWebsiteDataManagerPrivate(object *gobject.Object) *WebsiteDataManagerPrivate {
	return WebsiteDataManagerPrivateNewFromNative(object.Native())
}

func (recv *WebsiteDataManagerPrivate) Native() unsafe.Pointer {
	return recv.native
}

// WebsiteDataManagerPrivateStruct creates an uninitialised WebsiteDataManagerPrivate.
func WebsiteDataManagerPrivateStruct() *WebsiteDataManagerPrivate {
	err := websiteDataManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WebsiteDataManagerPrivateNewFromNative(websiteDataManagerPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWebsiteDataManagerPrivate)
	return structGo
}
func finalizeWebsiteDataManagerPrivate(obj *WebsiteDataManagerPrivate) {
	websiteDataManagerPrivateStruct.Free(obj.Native())
}

var windowPropertiesClassStruct *gi.Struct
var windowPropertiesClassStruct_Once sync.Once

func windowPropertiesClassStruct_Set() error {
	var err error
	windowPropertiesClassStruct_Once.Do(func() {
		windowPropertiesClassStruct, err = gi.StructNew("WebKit2", "WindowPropertiesClass")
	})
	return err
}

type WindowPropertiesClass struct {
	native unsafe.Pointer
}

func WindowPropertiesClassNewFromNative(native unsafe.Pointer) *WindowPropertiesClass {
	instance := &WindowPropertiesClass{native: native}

	return instance
}

/*
CastToWindowPropertiesClass down casts any arbitrary Object to WindowPropertiesClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a WindowPropertiesClass.
*/
func CastToWindowPropertiesClass(object *gobject.Object) *WindowPropertiesClass {
	return WindowPropertiesClassNewFromNative(object.Native())
}

func (recv *WindowPropertiesClass) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_webkit_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_webkit_reserved3' : for field setter : missing Type

// WindowPropertiesClassStruct creates an uninitialised WindowPropertiesClass.
func WindowPropertiesClassStruct() *WindowPropertiesClass {
	err := windowPropertiesClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WindowPropertiesClassNewFromNative(windowPropertiesClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWindowPropertiesClass)
	return structGo
}
func finalizeWindowPropertiesClass(obj *WindowPropertiesClass) {
	windowPropertiesClassStruct.Free(obj.Native())
}

var windowPropertiesPrivateStruct *gi.Struct
var windowPropertiesPrivateStruct_Once sync.Once

func windowPropertiesPrivateStruct_Set() error {
	var err error
	windowPropertiesPrivateStruct_Once.Do(func() {
		windowPropertiesPrivateStruct, err = gi.StructNew("WebKit2", "WindowPropertiesPrivate")
	})
	return err
}

type WindowPropertiesPrivate struct {
	native unsafe.Pointer
}

func WindowPropertiesPrivateNewFromNative(native unsafe.Pointer) *WindowPropertiesPrivate {
	instance := &WindowPropertiesPrivate{native: native}

	return instance
}

/*
CastToWindowPropertiesPrivate down casts any arbitrary Object to WindowPropertiesPrivate.
Exercise care, as this is a potentially dangerous function
if the Object is not a WindowPropertiesPrivate.
*/
func CastToWindowPropertiesPrivate(object *gobject.Object) *WindowPropertiesPrivate {
	return WindowPropertiesPrivateNewFromNative(object.Native())
}

func (recv *WindowPropertiesPrivate) Native() unsafe.Pointer {
	return recv.native
}

// WindowPropertiesPrivateStruct creates an uninitialised WindowPropertiesPrivate.
func WindowPropertiesPrivateStruct() *WindowPropertiesPrivate {
	err := windowPropertiesPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WindowPropertiesPrivateNewFromNative(windowPropertiesPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWindowPropertiesPrivate)
	return structGo
}
func finalizeWindowPropertiesPrivate(obj *WindowPropertiesPrivate) {
	windowPropertiesPrivateStruct.Free(obj.Native())
}
