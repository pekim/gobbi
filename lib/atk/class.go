// Code generated - DO NOT EDIT.

package atk

import (
	callback "github.com/pekim/gobbi/internal/cgo/callback"
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

var gObjectAccessibleObject *gi.Object
var gObjectAccessibleObject_Once sync.Once

func gObjectAccessibleObject_Set() error {
	var err error
	gObjectAccessibleObject_Once.Do(func() {
		gObjectAccessibleObject, err = gi.ObjectNew("Atk", "GObjectAccessible")
	})
	return err
}

type GObjectAccessible struct {
	native unsafe.Pointer
}

func GObjectAccessibleNewFromNative(native unsafe.Pointer) *GObjectAccessible {
	instance := &GObjectAccessible{native: native}

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

// ObjectAtk upcasts to *ObjectAtk
func (recv *GObjectAccessible) ObjectAtk() *Object {
	return ObjectNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *GObjectAccessible) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToGObjectAccessible down casts any arbitrary Object to GObjectAccessible.
Exercise care, as this is a potentially dangerous function
if the Object is not a GObjectAccessible.
*/
func CastToGObjectAccessible(object *gobject.Object) *GObjectAccessible {
	return GObjectAccessibleNewFromNative(object.Native())
}

// Equals compares this GObjectAccessible with another GObjectAccessible, and returns true if they represent the same Object.
func (recv *GObjectAccessible) Equals(other *GObjectAccessible) bool {
	return other.Native() == recv.Native()
}

func (recv *GObjectAccessible) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *GObjectAccessible) FieldParent() *Object {
	var nilValue *Object
	err := gObjectAccessibleObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(gObjectAccessibleObject, recv.Native(), "parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *GObjectAccessible) SetFieldParent(value *Object) {
	err := gObjectAccessibleObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(gObjectAccessibleObject, recv.Native(), "parent", argValue)
}

var gObjectAccessibleGetObjectFunction *gi.Function
var gObjectAccessibleGetObjectFunction_Once sync.Once

func gObjectAccessibleGetObjectFunction_Set() error {
	var err error
	gObjectAccessibleGetObjectFunction_Once.Do(func() {
		err = gObjectAccessibleObject_Set()
		if err != nil {
			return
		}
		gObjectAccessibleGetObjectFunction, err = gObjectAccessibleObject.InvokerNew("get_object")
	})
	return err
}

// GetObject is a representation of the C type atk_gobject_accessible_get_object.
func (recv *GObjectAccessible) GetObject() *gobject.Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := gObjectAccessibleGetObjectFunction_Set()
	if err == nil {
		ret = gObjectAccessibleGetObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := gobject.ObjectNewFromNative(ret.Pointer())

	return retGo
}

var hyperlinkObject *gi.Object
var hyperlinkObject_Once sync.Once

func hyperlinkObject_Set() error {
	var err error
	hyperlinkObject_Once.Do(func() {
		hyperlinkObject, err = gi.ObjectNew("Atk", "Hyperlink")
	})
	return err
}

type Hyperlink struct {
	native unsafe.Pointer
}

func HyperlinkNewFromNative(native unsafe.Pointer) *Hyperlink {
	instance := &Hyperlink{native: native}

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
func (recv *Hyperlink) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToHyperlink down casts any arbitrary Object to Hyperlink.
Exercise care, as this is a potentially dangerous function
if the Object is not a Hyperlink.
*/
func CastToHyperlink(object *gobject.Object) *Hyperlink {
	return HyperlinkNewFromNative(object.Native())
}

// Equals compares this Hyperlink with another Hyperlink, and returns true if they represent the same Object.
func (recv *Hyperlink) Equals(other *Hyperlink) bool {
	return other.Native() == recv.Native()
}

func (recv *Hyperlink) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Hyperlink) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := hyperlinkObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(hyperlinkObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Hyperlink) SetFieldParent(value *gobject.Object) {
	err := hyperlinkObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(hyperlinkObject, recv.Native(), "parent", argValue)
}

var hyperlinkGetEndIndexFunction *gi.Function
var hyperlinkGetEndIndexFunction_Once sync.Once

func hyperlinkGetEndIndexFunction_Set() error {
	var err error
	hyperlinkGetEndIndexFunction_Once.Do(func() {
		err = hyperlinkObject_Set()
		if err != nil {
			return
		}
		hyperlinkGetEndIndexFunction, err = hyperlinkObject.InvokerNew("get_end_index")
	})
	return err
}

// GetEndIndex is a representation of the C type atk_hyperlink_get_end_index.
func (recv *Hyperlink) GetEndIndex() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hyperlinkGetEndIndexFunction_Set()
	if err == nil {
		ret = hyperlinkGetEndIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var hyperlinkGetNAnchorsFunction *gi.Function
var hyperlinkGetNAnchorsFunction_Once sync.Once

func hyperlinkGetNAnchorsFunction_Set() error {
	var err error
	hyperlinkGetNAnchorsFunction_Once.Do(func() {
		err = hyperlinkObject_Set()
		if err != nil {
			return
		}
		hyperlinkGetNAnchorsFunction, err = hyperlinkObject.InvokerNew("get_n_anchors")
	})
	return err
}

// GetNAnchors is a representation of the C type atk_hyperlink_get_n_anchors.
func (recv *Hyperlink) GetNAnchors() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hyperlinkGetNAnchorsFunction_Set()
	if err == nil {
		ret = hyperlinkGetNAnchorsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var hyperlinkGetObjectFunction *gi.Function
var hyperlinkGetObjectFunction_Once sync.Once

func hyperlinkGetObjectFunction_Set() error {
	var err error
	hyperlinkGetObjectFunction_Once.Do(func() {
		err = hyperlinkObject_Set()
		if err != nil {
			return
		}
		hyperlinkGetObjectFunction, err = hyperlinkObject.InvokerNew("get_object")
	})
	return err
}

// GetObject is a representation of the C type atk_hyperlink_get_object.
func (recv *Hyperlink) GetObject(i int32) *Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := hyperlinkGetObjectFunction_Set()
	if err == nil {
		ret = hyperlinkGetObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var hyperlinkGetStartIndexFunction *gi.Function
var hyperlinkGetStartIndexFunction_Once sync.Once

func hyperlinkGetStartIndexFunction_Set() error {
	var err error
	hyperlinkGetStartIndexFunction_Once.Do(func() {
		err = hyperlinkObject_Set()
		if err != nil {
			return
		}
		hyperlinkGetStartIndexFunction, err = hyperlinkObject.InvokerNew("get_start_index")
	})
	return err
}

// GetStartIndex is a representation of the C type atk_hyperlink_get_start_index.
func (recv *Hyperlink) GetStartIndex() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hyperlinkGetStartIndexFunction_Set()
	if err == nil {
		ret = hyperlinkGetStartIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var hyperlinkGetUriFunction *gi.Function
var hyperlinkGetUriFunction_Once sync.Once

func hyperlinkGetUriFunction_Set() error {
	var err error
	hyperlinkGetUriFunction_Once.Do(func() {
		err = hyperlinkObject_Set()
		if err != nil {
			return
		}
		hyperlinkGetUriFunction, err = hyperlinkObject.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type atk_hyperlink_get_uri.
func (recv *Hyperlink) GetUri(i int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := hyperlinkGetUriFunction_Set()
	if err == nil {
		ret = hyperlinkGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var hyperlinkIsInlineFunction *gi.Function
var hyperlinkIsInlineFunction_Once sync.Once

func hyperlinkIsInlineFunction_Set() error {
	var err error
	hyperlinkIsInlineFunction_Once.Do(func() {
		err = hyperlinkObject_Set()
		if err != nil {
			return
		}
		hyperlinkIsInlineFunction, err = hyperlinkObject.InvokerNew("is_inline")
	})
	return err
}

// IsInline is a representation of the C type atk_hyperlink_is_inline.
func (recv *Hyperlink) IsInline() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hyperlinkIsInlineFunction_Set()
	if err == nil {
		ret = hyperlinkIsInlineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hyperlinkIsSelectedLinkFunction *gi.Function
var hyperlinkIsSelectedLinkFunction_Once sync.Once

func hyperlinkIsSelectedLinkFunction_Set() error {
	var err error
	hyperlinkIsSelectedLinkFunction_Once.Do(func() {
		err = hyperlinkObject_Set()
		if err != nil {
			return
		}
		hyperlinkIsSelectedLinkFunction, err = hyperlinkObject.InvokerNew("is_selected_link")
	})
	return err
}

// IsSelectedLink is a representation of the C type atk_hyperlink_is_selected_link.
func (recv *Hyperlink) IsSelectedLink() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hyperlinkIsSelectedLinkFunction_Set()
	if err == nil {
		ret = hyperlinkIsSelectedLinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var hyperlinkIsValidFunction *gi.Function
var hyperlinkIsValidFunction_Once sync.Once

func hyperlinkIsValidFunction_Set() error {
	var err error
	hyperlinkIsValidFunction_Once.Do(func() {
		err = hyperlinkObject_Set()
		if err != nil {
			return
		}
		hyperlinkIsValidFunction, err = hyperlinkObject.InvokerNew("is_valid")
	})
	return err
}

// IsValid is a representation of the C type atk_hyperlink_is_valid.
func (recv *Hyperlink) IsValid() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hyperlinkIsValidFunction_Set()
	if err == nil {
		ret = hyperlinkIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

/*
ConnectLinkActivated connects a callback to the 'link-activated' signal of the Hyperlink.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Hyperlink) ConnectLinkActivated(handler func(instance *Hyperlink)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := HyperlinkNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "link-activated", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Hyperlink) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

// Action returns the Action interface implemented by Hyperlink
func (recv *Hyperlink) Action() *Action {
	return ActionNewFromNative(recv.Native())
}

var miscObject *gi.Object
var miscObject_Once sync.Once

func miscObject_Set() error {
	var err error
	miscObject_Once.Do(func() {
		miscObject, err = gi.ObjectNew("Atk", "Misc")
	})
	return err
}

type Misc struct {
	native unsafe.Pointer
}

func MiscNewFromNative(native unsafe.Pointer) *Misc {
	instance := &Misc{native: native}

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
func (recv *Misc) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToMisc down casts any arbitrary Object to Misc.
Exercise care, as this is a potentially dangerous function
if the Object is not a Misc.
*/
func CastToMisc(object *gobject.Object) *Misc {
	return MiscNewFromNative(object.Native())
}

// Equals compares this Misc with another Misc, and returns true if they represent the same Object.
func (recv *Misc) Equals(other *Misc) bool {
	return other.Native() == recv.Native()
}

func (recv *Misc) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Misc) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := miscObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(miscObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Misc) SetFieldParent(value *gobject.Object) {
	err := miscObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(miscObject, recv.Native(), "parent", argValue)
}

var miscThreadsEnterFunction *gi.Function
var miscThreadsEnterFunction_Once sync.Once

func miscThreadsEnterFunction_Set() error {
	var err error
	miscThreadsEnterFunction_Once.Do(func() {
		err = miscObject_Set()
		if err != nil {
			return
		}
		miscThreadsEnterFunction, err = miscObject.InvokerNew("threads_enter")
	})
	return err
}

// ThreadsEnter is a representation of the C type atk_misc_threads_enter.
func (recv *Misc) ThreadsEnter() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := miscThreadsEnterFunction_Set()
	if err == nil {
		miscThreadsEnterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var miscThreadsLeaveFunction *gi.Function
var miscThreadsLeaveFunction_Once sync.Once

func miscThreadsLeaveFunction_Set() error {
	var err error
	miscThreadsLeaveFunction_Once.Do(func() {
		err = miscObject_Set()
		if err != nil {
			return
		}
		miscThreadsLeaveFunction, err = miscObject.InvokerNew("threads_leave")
	})
	return err
}

// ThreadsLeave is a representation of the C type atk_misc_threads_leave.
func (recv *Misc) ThreadsLeave() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := miscThreadsLeaveFunction_Set()
	if err == nil {
		miscThreadsLeaveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var noOpObjectObject *gi.Object
var noOpObjectObject_Once sync.Once

func noOpObjectObject_Set() error {
	var err error
	noOpObjectObject_Once.Do(func() {
		noOpObjectObject, err = gi.ObjectNew("Atk", "NoOpObject")
	})
	return err
}

type NoOpObject struct {
	native unsafe.Pointer
}

func NoOpObjectNewFromNative(native unsafe.Pointer) *NoOpObject {
	instance := &NoOpObject{native: native}

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

// ObjectAtk upcasts to *ObjectAtk
func (recv *NoOpObject) ObjectAtk() *Object {
	return ObjectNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *NoOpObject) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToNoOpObject down casts any arbitrary Object to NoOpObject.
Exercise care, as this is a potentially dangerous function
if the Object is not a NoOpObject.
*/
func CastToNoOpObject(object *gobject.Object) *NoOpObject {
	return NoOpObjectNewFromNative(object.Native())
}

// Equals compares this NoOpObject with another NoOpObject, and returns true if they represent the same Object.
func (recv *NoOpObject) Equals(other *NoOpObject) bool {
	return other.Native() == recv.Native()
}

func (recv *NoOpObject) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *NoOpObject) FieldParent() *Object {
	var nilValue *Object
	err := noOpObjectObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(noOpObjectObject, recv.Native(), "parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *NoOpObject) SetFieldParent(value *Object) {
	err := noOpObjectObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(noOpObjectObject, recv.Native(), "parent", argValue)
}

var noOpObjectNewFunction *gi.Function
var noOpObjectNewFunction_Once sync.Once

func noOpObjectNewFunction_Set() error {
	var err error
	noOpObjectNewFunction_Once.Do(func() {
		err = noOpObjectObject_Set()
		if err != nil {
			return
		}
		noOpObjectNewFunction, err = noOpObjectObject.InvokerNew("new")
	})
	return err
}

// NoOpObjectNew is a representation of the C type atk_no_op_object_new.
func NoOpObjectNew(obj *gobject.Object) *NoOpObject {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(obj.Native())

	var ret gi.Argument

	err := noOpObjectNewFunction_Set()
	if err == nil {
		ret = noOpObjectNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := NoOpObjectNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

// Action returns the Action interface implemented by NoOpObject
func (recv *NoOpObject) Action() *Action {
	return ActionNewFromNative(recv.Native())
}

// Component returns the Component interface implemented by NoOpObject
func (recv *NoOpObject) Component() *Component {
	return ComponentNewFromNative(recv.Native())
}

// Document returns the Document interface implemented by NoOpObject
func (recv *NoOpObject) Document() *Document {
	return DocumentNewFromNative(recv.Native())
}

// EditableText returns the EditableText interface implemented by NoOpObject
func (recv *NoOpObject) EditableText() *EditableText {
	return EditableTextNewFromNative(recv.Native())
}

// Hypertext returns the Hypertext interface implemented by NoOpObject
func (recv *NoOpObject) Hypertext() *Hypertext {
	return HypertextNewFromNative(recv.Native())
}

// Image returns the Image interface implemented by NoOpObject
func (recv *NoOpObject) Image() *Image {
	return ImageNewFromNative(recv.Native())
}

// Selection returns the Selection interface implemented by NoOpObject
func (recv *NoOpObject) Selection() *Selection {
	return SelectionNewFromNative(recv.Native())
}

// Table returns the Table interface implemented by NoOpObject
func (recv *NoOpObject) Table() *Table {
	return TableNewFromNative(recv.Native())
}

// TableCell returns the TableCell interface implemented by NoOpObject
func (recv *NoOpObject) TableCell() *TableCell {
	return TableCellNewFromNative(recv.Native())
}

// Text returns the Text interface implemented by NoOpObject
func (recv *NoOpObject) Text() *Text {
	return TextNewFromNative(recv.Native())
}

// Value returns the Value interface implemented by NoOpObject
func (recv *NoOpObject) Value() *Value {
	return ValueNewFromNative(recv.Native())
}

// Window returns the Window interface implemented by NoOpObject
func (recv *NoOpObject) Window() *Window {
	return WindowNewFromNative(recv.Native())
}

var noOpObjectFactoryObject *gi.Object
var noOpObjectFactoryObject_Once sync.Once

func noOpObjectFactoryObject_Set() error {
	var err error
	noOpObjectFactoryObject_Once.Do(func() {
		noOpObjectFactoryObject, err = gi.ObjectNew("Atk", "NoOpObjectFactory")
	})
	return err
}

type NoOpObjectFactory struct {
	native unsafe.Pointer
}

func NoOpObjectFactoryNewFromNative(native unsafe.Pointer) *NoOpObjectFactory {
	instance := &NoOpObjectFactory{native: native}

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

// ObjectFactory upcasts to *ObjectFactory
func (recv *NoOpObjectFactory) ObjectFactory() *ObjectFactory {
	return ObjectFactoryNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *NoOpObjectFactory) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToNoOpObjectFactory down casts any arbitrary Object to NoOpObjectFactory.
Exercise care, as this is a potentially dangerous function
if the Object is not a NoOpObjectFactory.
*/
func CastToNoOpObjectFactory(object *gobject.Object) *NoOpObjectFactory {
	return NoOpObjectFactoryNewFromNative(object.Native())
}

// Equals compares this NoOpObjectFactory with another NoOpObjectFactory, and returns true if they represent the same Object.
func (recv *NoOpObjectFactory) Equals(other *NoOpObjectFactory) bool {
	return other.Native() == recv.Native()
}

func (recv *NoOpObjectFactory) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *NoOpObjectFactory) FieldParent() *ObjectFactory {
	var nilValue *ObjectFactory
	err := noOpObjectFactoryObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(noOpObjectFactoryObject, recv.Native(), "parent")
	value := ObjectFactoryNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *NoOpObjectFactory) SetFieldParent(value *ObjectFactory) {
	err := noOpObjectFactoryObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(noOpObjectFactoryObject, recv.Native(), "parent", argValue)
}

var noOpObjectFactoryNewFunction *gi.Function
var noOpObjectFactoryNewFunction_Once sync.Once

func noOpObjectFactoryNewFunction_Set() error {
	var err error
	noOpObjectFactoryNewFunction_Once.Do(func() {
		err = noOpObjectFactoryObject_Set()
		if err != nil {
			return
		}
		noOpObjectFactoryNewFunction, err = noOpObjectFactoryObject.InvokerNew("new")
	})
	return err
}

// NoOpObjectFactoryNew is a representation of the C type atk_no_op_object_factory_new.
func NoOpObjectFactoryNew() *NoOpObjectFactory {

	var ret gi.Argument

	err := noOpObjectFactoryNewFunction_Set()
	if err == nil {
		ret = noOpObjectFactoryNewFunction.Invoke(nil, nil)
	}

	retGo := NoOpObjectFactoryNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var objectObject *gi.Object
var objectObject_Once sync.Once

func objectObject_Set() error {
	var err error
	objectObject_Once.Do(func() {
		objectObject, err = gi.ObjectNew("Atk", "Object")
	})
	return err
}

type Object struct {
	native unsafe.Pointer
}

func ObjectNewFromNative(native unsafe.Pointer) *Object {
	instance := &Object{native: native}

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
func (recv *Object) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToObject down casts any arbitrary Object to Object.
Exercise care, as this is a potentially dangerous function
if the Object is not a Object.
*/
func CastToObject(object *gobject.Object) *Object {
	return ObjectNewFromNative(object.Native())
}

// Equals compares this Object with another Object, and returns true if they represent the same Object.
func (recv *Object) Equals(other *Object) bool {
	return other.Native() == recv.Native()
}

func (recv *Object) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Object) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := objectObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(objectObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Object) SetFieldParent(value *gobject.Object) {
	err := objectObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(objectObject, recv.Native(), "parent", argValue)
}

// FieldDescription returns the C field 'description'.
func (recv *Object) FieldDescription() string {
	var nilValue string
	err := objectObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(objectObject, recv.Native(), "description")
	value := argValue.String(false)
	return value
}

// SetFieldDescription sets the value of the C field 'description'.
func (recv *Object) SetFieldDescription(value string) {
	err := objectObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(objectObject, recv.Native(), "description", argValue)
}

// FieldName returns the C field 'name'.
func (recv *Object) FieldName() string {
	var nilValue string
	err := objectObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(objectObject, recv.Native(), "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *Object) SetFieldName(value string) {
	err := objectObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(objectObject, recv.Native(), "name", argValue)
}

// FieldAccessibleParent returns the C field 'accessible_parent'.
func (recv *Object) FieldAccessibleParent() *Object {
	var nilValue *Object
	err := objectObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(objectObject, recv.Native(), "accessible_parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAccessibleParent sets the value of the C field 'accessible_parent'.
func (recv *Object) SetFieldAccessibleParent(value *Object) {
	err := objectObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(objectObject, recv.Native(), "accessible_parent", argValue)
}

// FieldRole returns the C field 'role'.
func (recv *Object) FieldRole() Role {
	var nilValue Role
	err := objectObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(objectObject, recv.Native(), "role")
	value := Role(argValue.Int32())
	return value
}

// SetFieldRole sets the value of the C field 'role'.
func (recv *Object) SetFieldRole(value Role) {
	err := objectObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.ObjectFieldSet(objectObject, recv.Native(), "role", argValue)
}

// FieldRelationSet returns the C field 'relation_set'.
func (recv *Object) FieldRelationSet() *RelationSet {
	var nilValue *RelationSet
	err := objectObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(objectObject, recv.Native(), "relation_set")
	value := RelationSetNewFromNative(argValue.Pointer())
	return value
}

// SetFieldRelationSet sets the value of the C field 'relation_set'.
func (recv *Object) SetFieldRelationSet(value *RelationSet) {
	err := objectObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(objectObject, recv.Native(), "relation_set", argValue)
}

// FieldLayer returns the C field 'layer'.
func (recv *Object) FieldLayer() Layer {
	var nilValue Layer
	err := objectObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(objectObject, recv.Native(), "layer")
	value := Layer(argValue.Int32())
	return value
}

// SetFieldLayer sets the value of the C field 'layer'.
func (recv *Object) SetFieldLayer(value Layer) {
	err := objectObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.ObjectFieldSet(objectObject, recv.Native(), "layer", argValue)
}

var objectAddRelationshipFunction *gi.Function
var objectAddRelationshipFunction_Once sync.Once

func objectAddRelationshipFunction_Set() error {
	var err error
	objectAddRelationshipFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectAddRelationshipFunction, err = objectObject.InvokerNew("add_relationship")
	})
	return err
}

// AddRelationship is a representation of the C type atk_object_add_relationship.
func (recv *Object) AddRelationship(relationship RelationType, target *Object) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(relationship))
	inArgs[2].SetPointer(target.Native())

	var ret gi.Argument

	err := objectAddRelationshipFunction_Set()
	if err == nil {
		ret = objectAddRelationshipFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'atk_object_connect_property_change_handler' : parameter 'handler' of type 'PropertyChangeHandler' not supported

var objectGetAccessibleIdFunction *gi.Function
var objectGetAccessibleIdFunction_Once sync.Once

func objectGetAccessibleIdFunction_Set() error {
	var err error
	objectGetAccessibleIdFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetAccessibleIdFunction, err = objectObject.InvokerNew("get_accessible_id")
	})
	return err
}

// GetAccessibleId is a representation of the C type atk_object_get_accessible_id.
func (recv *Object) GetAccessibleId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetAccessibleIdFunction_Set()
	if err == nil {
		ret = objectGetAccessibleIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'atk_object_get_attributes' : return type 'AttributeSet' not supported

var objectGetDescriptionFunction *gi.Function
var objectGetDescriptionFunction_Once sync.Once

func objectGetDescriptionFunction_Set() error {
	var err error
	objectGetDescriptionFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetDescriptionFunction, err = objectObject.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type atk_object_get_description.
func (recv *Object) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetDescriptionFunction_Set()
	if err == nil {
		ret = objectGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var objectGetIndexInParentFunction *gi.Function
var objectGetIndexInParentFunction_Once sync.Once

func objectGetIndexInParentFunction_Set() error {
	var err error
	objectGetIndexInParentFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetIndexInParentFunction, err = objectObject.InvokerNew("get_index_in_parent")
	})
	return err
}

// GetIndexInParent is a representation of the C type atk_object_get_index_in_parent.
func (recv *Object) GetIndexInParent() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetIndexInParentFunction_Set()
	if err == nil {
		ret = objectGetIndexInParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var objectGetLayerFunction *gi.Function
var objectGetLayerFunction_Once sync.Once

func objectGetLayerFunction_Set() error {
	var err error
	objectGetLayerFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetLayerFunction, err = objectObject.InvokerNew("get_layer")
	})
	return err
}

// GetLayer is a representation of the C type atk_object_get_layer.
func (recv *Object) GetLayer() Layer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetLayerFunction_Set()
	if err == nil {
		ret = objectGetLayerFunction.Invoke(inArgs[:], nil)
	}

	retGo := Layer(ret.Int32())

	return retGo
}

var objectGetMdiZorderFunction *gi.Function
var objectGetMdiZorderFunction_Once sync.Once

func objectGetMdiZorderFunction_Set() error {
	var err error
	objectGetMdiZorderFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetMdiZorderFunction, err = objectObject.InvokerNew("get_mdi_zorder")
	})
	return err
}

// GetMdiZorder is a representation of the C type atk_object_get_mdi_zorder.
func (recv *Object) GetMdiZorder() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetMdiZorderFunction_Set()
	if err == nil {
		ret = objectGetMdiZorderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var objectGetNAccessibleChildrenFunction *gi.Function
var objectGetNAccessibleChildrenFunction_Once sync.Once

func objectGetNAccessibleChildrenFunction_Set() error {
	var err error
	objectGetNAccessibleChildrenFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetNAccessibleChildrenFunction, err = objectObject.InvokerNew("get_n_accessible_children")
	})
	return err
}

// GetNAccessibleChildren is a representation of the C type atk_object_get_n_accessible_children.
func (recv *Object) GetNAccessibleChildren() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetNAccessibleChildrenFunction_Set()
	if err == nil {
		ret = objectGetNAccessibleChildrenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var objectGetNameFunction *gi.Function
var objectGetNameFunction_Once sync.Once

func objectGetNameFunction_Set() error {
	var err error
	objectGetNameFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetNameFunction, err = objectObject.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type atk_object_get_name.
func (recv *Object) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetNameFunction_Set()
	if err == nil {
		ret = objectGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var objectGetObjectLocaleFunction *gi.Function
var objectGetObjectLocaleFunction_Once sync.Once

func objectGetObjectLocaleFunction_Set() error {
	var err error
	objectGetObjectLocaleFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetObjectLocaleFunction, err = objectObject.InvokerNew("get_object_locale")
	})
	return err
}

// GetObjectLocale is a representation of the C type atk_object_get_object_locale.
func (recv *Object) GetObjectLocale() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetObjectLocaleFunction_Set()
	if err == nil {
		ret = objectGetObjectLocaleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var objectGetParentFunction *gi.Function
var objectGetParentFunction_Once sync.Once

func objectGetParentFunction_Set() error {
	var err error
	objectGetParentFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetParentFunction, err = objectObject.InvokerNew("get_parent")
	})
	return err
}

// GetParent is a representation of the C type atk_object_get_parent.
func (recv *Object) GetParent() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetParentFunction_Set()
	if err == nil {
		ret = objectGetParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var objectGetRoleFunction *gi.Function
var objectGetRoleFunction_Once sync.Once

func objectGetRoleFunction_Set() error {
	var err error
	objectGetRoleFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectGetRoleFunction, err = objectObject.InvokerNew("get_role")
	})
	return err
}

// GetRole is a representation of the C type atk_object_get_role.
func (recv *Object) GetRole() Role {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectGetRoleFunction_Set()
	if err == nil {
		ret = objectGetRoleFunction.Invoke(inArgs[:], nil)
	}

	retGo := Role(ret.Int32())

	return retGo
}

var objectInitializeFunction *gi.Function
var objectInitializeFunction_Once sync.Once

func objectInitializeFunction_Set() error {
	var err error
	objectInitializeFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectInitializeFunction, err = objectObject.InvokerNew("initialize")
	})
	return err
}

// Initialize is a representation of the C type atk_object_initialize.
func (recv *Object) Initialize(data unsafe.Pointer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(data)

	err := objectInitializeFunction_Set()
	if err == nil {
		objectInitializeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectNotifyStateChangeFunction *gi.Function
var objectNotifyStateChangeFunction_Once sync.Once

func objectNotifyStateChangeFunction_Set() error {
	var err error
	objectNotifyStateChangeFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectNotifyStateChangeFunction, err = objectObject.InvokerNew("notify_state_change")
	})
	return err
}

// NotifyStateChange is a representation of the C type atk_object_notify_state_change.
func (recv *Object) NotifyStateChange(state State, value bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint64(uint64(state))
	inArgs[2].SetBoolean(value)

	err := objectNotifyStateChangeFunction_Set()
	if err == nil {
		objectNotifyStateChangeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectPeekParentFunction *gi.Function
var objectPeekParentFunction_Once sync.Once

func objectPeekParentFunction_Set() error {
	var err error
	objectPeekParentFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectPeekParentFunction, err = objectObject.InvokerNew("peek_parent")
	})
	return err
}

// PeekParent is a representation of the C type atk_object_peek_parent.
func (recv *Object) PeekParent() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectPeekParentFunction_Set()
	if err == nil {
		ret = objectPeekParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var objectRefAccessibleChildFunction *gi.Function
var objectRefAccessibleChildFunction_Once sync.Once

func objectRefAccessibleChildFunction_Set() error {
	var err error
	objectRefAccessibleChildFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectRefAccessibleChildFunction, err = objectObject.InvokerNew("ref_accessible_child")
	})
	return err
}

// RefAccessibleChild is a representation of the C type atk_object_ref_accessible_child.
func (recv *Object) RefAccessibleChild(i int32) *Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := objectRefAccessibleChildFunction_Set()
	if err == nil {
		ret = objectRefAccessibleChildFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var objectRefRelationSetFunction *gi.Function
var objectRefRelationSetFunction_Once sync.Once

func objectRefRelationSetFunction_Set() error {
	var err error
	objectRefRelationSetFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectRefRelationSetFunction, err = objectObject.InvokerNew("ref_relation_set")
	})
	return err
}

// RefRelationSet is a representation of the C type atk_object_ref_relation_set.
func (recv *Object) RefRelationSet() *RelationSet {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectRefRelationSetFunction_Set()
	if err == nil {
		ret = objectRefRelationSetFunction.Invoke(inArgs[:], nil)
	}

	retGo := RelationSetNewFromNative(ret.Pointer())

	return retGo
}

var objectRefStateSetFunction *gi.Function
var objectRefStateSetFunction_Once sync.Once

func objectRefStateSetFunction_Set() error {
	var err error
	objectRefStateSetFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectRefStateSetFunction, err = objectObject.InvokerNew("ref_state_set")
	})
	return err
}

// RefStateSet is a representation of the C type atk_object_ref_state_set.
func (recv *Object) RefStateSet() *StateSet {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectRefStateSetFunction_Set()
	if err == nil {
		ret = objectRefStateSetFunction.Invoke(inArgs[:], nil)
	}

	retGo := StateSetNewFromNative(ret.Pointer())

	return retGo
}

var objectRemovePropertyChangeHandlerFunction *gi.Function
var objectRemovePropertyChangeHandlerFunction_Once sync.Once

func objectRemovePropertyChangeHandlerFunction_Set() error {
	var err error
	objectRemovePropertyChangeHandlerFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectRemovePropertyChangeHandlerFunction, err = objectObject.InvokerNew("remove_property_change_handler")
	})
	return err
}

// RemovePropertyChangeHandler is a representation of the C type atk_object_remove_property_change_handler.
func (recv *Object) RemovePropertyChangeHandler(handlerId uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(handlerId)

	err := objectRemovePropertyChangeHandlerFunction_Set()
	if err == nil {
		objectRemovePropertyChangeHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectRemoveRelationshipFunction *gi.Function
var objectRemoveRelationshipFunction_Once sync.Once

func objectRemoveRelationshipFunction_Set() error {
	var err error
	objectRemoveRelationshipFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectRemoveRelationshipFunction, err = objectObject.InvokerNew("remove_relationship")
	})
	return err
}

// RemoveRelationship is a representation of the C type atk_object_remove_relationship.
func (recv *Object) RemoveRelationship(relationship RelationType, target *Object) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(relationship))
	inArgs[2].SetPointer(target.Native())

	var ret gi.Argument

	err := objectRemoveRelationshipFunction_Set()
	if err == nil {
		ret = objectRemoveRelationshipFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var objectSetAccessibleIdFunction *gi.Function
var objectSetAccessibleIdFunction_Once sync.Once

func objectSetAccessibleIdFunction_Set() error {
	var err error
	objectSetAccessibleIdFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectSetAccessibleIdFunction, err = objectObject.InvokerNew("set_accessible_id")
	})
	return err
}

// SetAccessibleId is a representation of the C type atk_object_set_accessible_id.
func (recv *Object) SetAccessibleId(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	err := objectSetAccessibleIdFunction_Set()
	if err == nil {
		objectSetAccessibleIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectSetDescriptionFunction *gi.Function
var objectSetDescriptionFunction_Once sync.Once

func objectSetDescriptionFunction_Set() error {
	var err error
	objectSetDescriptionFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectSetDescriptionFunction, err = objectObject.InvokerNew("set_description")
	})
	return err
}

// SetDescription is a representation of the C type atk_object_set_description.
func (recv *Object) SetDescription(description string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(description)

	err := objectSetDescriptionFunction_Set()
	if err == nil {
		objectSetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectSetNameFunction *gi.Function
var objectSetNameFunction_Once sync.Once

func objectSetNameFunction_Set() error {
	var err error
	objectSetNameFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectSetNameFunction, err = objectObject.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type atk_object_set_name.
func (recv *Object) SetName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	err := objectSetNameFunction_Set()
	if err == nil {
		objectSetNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectSetParentFunction *gi.Function
var objectSetParentFunction_Once sync.Once

func objectSetParentFunction_Set() error {
	var err error
	objectSetParentFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectSetParentFunction, err = objectObject.InvokerNew("set_parent")
	})
	return err
}

// SetParent is a representation of the C type atk_object_set_parent.
func (recv *Object) SetParent(parent *Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(parent.Native())

	err := objectSetParentFunction_Set()
	if err == nil {
		objectSetParentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var objectSetRoleFunction *gi.Function
var objectSetRoleFunction_Once sync.Once

func objectSetRoleFunction_Set() error {
	var err error
	objectSetRoleFunction_Once.Do(func() {
		err = objectObject_Set()
		if err != nil {
			return
		}
		objectSetRoleFunction, err = objectObject.InvokerNew("set_role")
	})
	return err
}

// SetRole is a representation of the C type atk_object_set_role.
func (recv *Object) SetRole(role Role) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(role))

	err := objectSetRoleFunction_Set()
	if err == nil {
		objectSetRoleFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectActiveDescendantChanged connects a callback to the 'active-descendant-changed' signal of the Object.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Object) ConnectActiveDescendantChanged(handler func(instance *Object, arg1 *Object)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ObjectNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := ObjectNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "active-descendant-changed", marshal)
}

/*
ConnectChildrenChanged connects a callback to the 'children-changed' signal of the Object.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Object) ConnectChildrenChanged(handler func(instance *Object, arg1 uint32, arg2 *Object)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ObjectNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetUint()

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := ObjectNewFromNative(object2.GetObject().Native())

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "children-changed", marshal)
}

/*
ConnectFocusEvent connects a callback to the 'focus-event' signal of the Object.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Object) ConnectFocusEvent(handler func(instance *Object, arg1 bool)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ObjectNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetBoolean()

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "focus-event", marshal)
}

/*
ConnectPropertyChange connects a callback to the 'property-change' signal of the Object.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Object) ConnectPropertyChange(handler func(instance *Object, arg1 *PropertyValues)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ObjectNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := PropertyValuesNewFromNative(object1.GetPointer())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "property-change", marshal)
}

/*
ConnectStateChange connects a callback to the 'state-change' signal of the Object.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Object) ConnectStateChange(handler func(instance *Object, arg1 string, arg2 bool)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ObjectNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetString()

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := object2.GetBoolean()

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "state-change", marshal)
}

/*
ConnectVisibleDataChanged connects a callback to the 'visible-data-changed' signal of the Object.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Object) ConnectVisibleDataChanged(handler func(instance *Object)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ObjectNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "visible-data-changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Object) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var objectFactoryObject *gi.Object
var objectFactoryObject_Once sync.Once

func objectFactoryObject_Set() error {
	var err error
	objectFactoryObject_Once.Do(func() {
		objectFactoryObject, err = gi.ObjectNew("Atk", "ObjectFactory")
	})
	return err
}

type ObjectFactory struct {
	native unsafe.Pointer
}

func ObjectFactoryNewFromNative(native unsafe.Pointer) *ObjectFactory {
	instance := &ObjectFactory{native: native}

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
func (recv *ObjectFactory) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToObjectFactory down casts any arbitrary Object to ObjectFactory.
Exercise care, as this is a potentially dangerous function
if the Object is not a ObjectFactory.
*/
func CastToObjectFactory(object *gobject.Object) *ObjectFactory {
	return ObjectFactoryNewFromNative(object.Native())
}

// Equals compares this ObjectFactory with another ObjectFactory, and returns true if they represent the same Object.
func (recv *ObjectFactory) Equals(other *ObjectFactory) bool {
	return other.Native() == recv.Native()
}

func (recv *ObjectFactory) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *ObjectFactory) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := objectFactoryObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(objectFactoryObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *ObjectFactory) SetFieldParent(value *gobject.Object) {
	err := objectFactoryObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(objectFactoryObject, recv.Native(), "parent", argValue)
}

var objectFactoryCreateAccessibleFunction *gi.Function
var objectFactoryCreateAccessibleFunction_Once sync.Once

func objectFactoryCreateAccessibleFunction_Set() error {
	var err error
	objectFactoryCreateAccessibleFunction_Once.Do(func() {
		err = objectFactoryObject_Set()
		if err != nil {
			return
		}
		objectFactoryCreateAccessibleFunction, err = objectFactoryObject.InvokerNew("create_accessible")
	})
	return err
}

// CreateAccessible is a representation of the C type atk_object_factory_create_accessible.
func (recv *ObjectFactory) CreateAccessible(obj *gobject.Object) *Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(obj.Native())

	var ret gi.Argument

	err := objectFactoryCreateAccessibleFunction_Set()
	if err == nil {
		ret = objectFactoryCreateAccessibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var objectFactoryGetAccessibleTypeFunction *gi.Function
var objectFactoryGetAccessibleTypeFunction_Once sync.Once

func objectFactoryGetAccessibleTypeFunction_Set() error {
	var err error
	objectFactoryGetAccessibleTypeFunction_Once.Do(func() {
		err = objectFactoryObject_Set()
		if err != nil {
			return
		}
		objectFactoryGetAccessibleTypeFunction, err = objectFactoryObject.InvokerNew("get_accessible_type")
	})
	return err
}

// GetAccessibleType is a representation of the C type atk_object_factory_get_accessible_type.
func (recv *ObjectFactory) GetAccessibleType() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := objectFactoryGetAccessibleTypeFunction_Set()
	if err == nil {
		ret = objectFactoryGetAccessibleTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var objectFactoryInvalidateFunction *gi.Function
var objectFactoryInvalidateFunction_Once sync.Once

func objectFactoryInvalidateFunction_Set() error {
	var err error
	objectFactoryInvalidateFunction_Once.Do(func() {
		err = objectFactoryObject_Set()
		if err != nil {
			return
		}
		objectFactoryInvalidateFunction, err = objectFactoryObject.InvokerNew("invalidate")
	})
	return err
}

// Invalidate is a representation of the C type atk_object_factory_invalidate.
func (recv *ObjectFactory) Invalidate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := objectFactoryInvalidateFunction_Set()
	if err == nil {
		objectFactoryInvalidateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var plugObject *gi.Object
var plugObject_Once sync.Once

func plugObject_Set() error {
	var err error
	plugObject_Once.Do(func() {
		plugObject, err = gi.ObjectNew("Atk", "Plug")
	})
	return err
}

type Plug struct {
	native unsafe.Pointer
}

func PlugNewFromNative(native unsafe.Pointer) *Plug {
	instance := &Plug{native: native}

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

// ObjectAtk upcasts to *ObjectAtk
func (recv *Plug) ObjectAtk() *Object {
	return ObjectNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *Plug) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToPlug down casts any arbitrary Object to Plug.
Exercise care, as this is a potentially dangerous function
if the Object is not a Plug.
*/
func CastToPlug(object *gobject.Object) *Plug {
	return PlugNewFromNative(object.Native())
}

// Equals compares this Plug with another Plug, and returns true if they represent the same Object.
func (recv *Plug) Equals(other *Plug) bool {
	return other.Native() == recv.Native()
}

func (recv *Plug) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Plug) FieldParent() *Object {
	var nilValue *Object
	err := plugObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(plugObject, recv.Native(), "parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Plug) SetFieldParent(value *Object) {
	err := plugObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(plugObject, recv.Native(), "parent", argValue)
}

var plugNewFunction *gi.Function
var plugNewFunction_Once sync.Once

func plugNewFunction_Set() error {
	var err error
	plugNewFunction_Once.Do(func() {
		err = plugObject_Set()
		if err != nil {
			return
		}
		plugNewFunction, err = plugObject.InvokerNew("new")
	})
	return err
}

// PlugNew is a representation of the C type atk_plug_new.
func PlugNew() *Plug {

	var ret gi.Argument

	err := plugNewFunction_Set()
	if err == nil {
		ret = plugNewFunction.Invoke(nil, nil)
	}

	retGo := PlugNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var plugGetIdFunction *gi.Function
var plugGetIdFunction_Once sync.Once

func plugGetIdFunction_Set() error {
	var err error
	plugGetIdFunction_Once.Do(func() {
		err = plugObject_Set()
		if err != nil {
			return
		}
		plugGetIdFunction, err = plugObject.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type atk_plug_get_id.
func (recv *Plug) GetId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := plugGetIdFunction_Set()
	if err == nil {
		ret = plugGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// Component returns the Component interface implemented by Plug
func (recv *Plug) Component() *Component {
	return ComponentNewFromNative(recv.Native())
}

var registryObject *gi.Object
var registryObject_Once sync.Once

func registryObject_Set() error {
	var err error
	registryObject_Once.Do(func() {
		registryObject, err = gi.ObjectNew("Atk", "Registry")
	})
	return err
}

type Registry struct {
	native unsafe.Pointer
}

func RegistryNewFromNative(native unsafe.Pointer) *Registry {
	instance := &Registry{native: native}

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
func (recv *Registry) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRegistry down casts any arbitrary Object to Registry.
Exercise care, as this is a potentially dangerous function
if the Object is not a Registry.
*/
func CastToRegistry(object *gobject.Object) *Registry {
	return RegistryNewFromNative(object.Native())
}

// Equals compares this Registry with another Registry, and returns true if they represent the same Object.
func (recv *Registry) Equals(other *Registry) bool {
	return other.Native() == recv.Native()
}

func (recv *Registry) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Registry) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := registryObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(registryObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Registry) SetFieldParent(value *gobject.Object) {
	err := registryObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(registryObject, recv.Native(), "parent", argValue)
}

// FieldFactoryTypeRegistry returns the C field 'factory_type_registry'.
func (recv *Registry) FieldFactoryTypeRegistry() *glib.HashTable {
	var nilValue *glib.HashTable
	err := registryObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(registryObject, recv.Native(), "factory_type_registry")
	value := glib.HashTableNewFromNative(argValue.Pointer())
	return value
}

// SetFieldFactoryTypeRegistry sets the value of the C field 'factory_type_registry'.
func (recv *Registry) SetFieldFactoryTypeRegistry(value *glib.HashTable) {
	err := registryObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(registryObject, recv.Native(), "factory_type_registry", argValue)
}

// FieldFactorySingletonCache returns the C field 'factory_singleton_cache'.
func (recv *Registry) FieldFactorySingletonCache() *glib.HashTable {
	var nilValue *glib.HashTable
	err := registryObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(registryObject, recv.Native(), "factory_singleton_cache")
	value := glib.HashTableNewFromNative(argValue.Pointer())
	return value
}

// SetFieldFactorySingletonCache sets the value of the C field 'factory_singleton_cache'.
func (recv *Registry) SetFieldFactorySingletonCache(value *glib.HashTable) {
	err := registryObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(registryObject, recv.Native(), "factory_singleton_cache", argValue)
}

var registryGetFactoryFunction *gi.Function
var registryGetFactoryFunction_Once sync.Once

func registryGetFactoryFunction_Set() error {
	var err error
	registryGetFactoryFunction_Once.Do(func() {
		err = registryObject_Set()
		if err != nil {
			return
		}
		registryGetFactoryFunction, err = registryObject.InvokerNew("get_factory")
	})
	return err
}

// GetFactory is a representation of the C type atk_registry_get_factory.
func (recv *Registry) GetFactory(type_ int64) *ObjectFactory {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(type_)

	var ret gi.Argument

	err := registryGetFactoryFunction_Set()
	if err == nil {
		ret = registryGetFactoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectFactoryNewFromNative(ret.Pointer())

	return retGo
}

var registryGetFactoryTypeFunction *gi.Function
var registryGetFactoryTypeFunction_Once sync.Once

func registryGetFactoryTypeFunction_Set() error {
	var err error
	registryGetFactoryTypeFunction_Once.Do(func() {
		err = registryObject_Set()
		if err != nil {
			return
		}
		registryGetFactoryTypeFunction, err = registryObject.InvokerNew("get_factory_type")
	})
	return err
}

// GetFactoryType is a representation of the C type atk_registry_get_factory_type.
func (recv *Registry) GetFactoryType(type_ int64) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(type_)

	var ret gi.Argument

	err := registryGetFactoryTypeFunction_Set()
	if err == nil {
		ret = registryGetFactoryTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var registrySetFactoryTypeFunction *gi.Function
var registrySetFactoryTypeFunction_Once sync.Once

func registrySetFactoryTypeFunction_Set() error {
	var err error
	registrySetFactoryTypeFunction_Once.Do(func() {
		err = registryObject_Set()
		if err != nil {
			return
		}
		registrySetFactoryTypeFunction, err = registryObject.InvokerNew("set_factory_type")
	})
	return err
}

// SetFactoryType is a representation of the C type atk_registry_set_factory_type.
func (recv *Registry) SetFactoryType(type_ int64, factoryType int64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt64(type_)
	inArgs[2].SetInt64(factoryType)

	err := registrySetFactoryTypeFunction_Set()
	if err == nil {
		registrySetFactoryTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var relationObject *gi.Object
var relationObject_Once sync.Once

func relationObject_Set() error {
	var err error
	relationObject_Once.Do(func() {
		relationObject, err = gi.ObjectNew("Atk", "Relation")
	})
	return err
}

type Relation struct {
	native unsafe.Pointer
}

func RelationNewFromNative(native unsafe.Pointer) *Relation {
	instance := &Relation{native: native}

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
func (recv *Relation) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRelation down casts any arbitrary Object to Relation.
Exercise care, as this is a potentially dangerous function
if the Object is not a Relation.
*/
func CastToRelation(object *gobject.Object) *Relation {
	return RelationNewFromNative(object.Native())
}

// Equals compares this Relation with another Relation, and returns true if they represent the same Object.
func (recv *Relation) Equals(other *Relation) bool {
	return other.Native() == recv.Native()
}

func (recv *Relation) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Relation) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := relationObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(relationObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Relation) SetFieldParent(value *gobject.Object) {
	err := relationObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(relationObject, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'target' : for field getter : missing Type

// UNSUPPORTED : C value 'target' : for field setter : missing Type

// FieldRelationship returns the C field 'relationship'.
func (recv *Relation) FieldRelationship() RelationType {
	var nilValue RelationType
	err := relationObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(relationObject, recv.Native(), "relationship")
	value := RelationType(argValue.Int32())
	return value
}

// SetFieldRelationship sets the value of the C field 'relationship'.
func (recv *Relation) SetFieldRelationship(value RelationType) {
	err := relationObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.ObjectFieldSet(relationObject, recv.Native(), "relationship", argValue)
}

// UNSUPPORTED : C value 'atk_relation_new' : parameter 'targets' of type 'nil' not supported

var relationAddTargetFunction *gi.Function
var relationAddTargetFunction_Once sync.Once

func relationAddTargetFunction_Set() error {
	var err error
	relationAddTargetFunction_Once.Do(func() {
		err = relationObject_Set()
		if err != nil {
			return
		}
		relationAddTargetFunction, err = relationObject.InvokerNew("add_target")
	})
	return err
}

// AddTarget is a representation of the C type atk_relation_add_target.
func (recv *Relation) AddTarget(target *Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(target.Native())

	err := relationAddTargetFunction_Set()
	if err == nil {
		relationAddTargetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var relationGetRelationTypeFunction *gi.Function
var relationGetRelationTypeFunction_Once sync.Once

func relationGetRelationTypeFunction_Set() error {
	var err error
	relationGetRelationTypeFunction_Once.Do(func() {
		err = relationObject_Set()
		if err != nil {
			return
		}
		relationGetRelationTypeFunction, err = relationObject.InvokerNew("get_relation_type")
	})
	return err
}

// GetRelationType is a representation of the C type atk_relation_get_relation_type.
func (recv *Relation) GetRelationType() RelationType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := relationGetRelationTypeFunction_Set()
	if err == nil {
		ret = relationGetRelationTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := RelationType(ret.Int32())

	return retGo
}

var relationGetTargetFunction *gi.Function
var relationGetTargetFunction_Once sync.Once

func relationGetTargetFunction_Set() error {
	var err error
	relationGetTargetFunction_Once.Do(func() {
		err = relationObject_Set()
		if err != nil {
			return
		}
		relationGetTargetFunction, err = relationObject.InvokerNew("get_target")
	})
	return err
}

// GetTarget is a representation of the C type atk_relation_get_target.
func (recv *Relation) GetTarget() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := relationGetTargetFunction_Set()
	if err == nil {
		relationGetTargetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var relationRemoveTargetFunction *gi.Function
var relationRemoveTargetFunction_Once sync.Once

func relationRemoveTargetFunction_Set() error {
	var err error
	relationRemoveTargetFunction_Once.Do(func() {
		err = relationObject_Set()
		if err != nil {
			return
		}
		relationRemoveTargetFunction, err = relationObject.InvokerNew("remove_target")
	})
	return err
}

// RemoveTarget is a representation of the C type atk_relation_remove_target.
func (recv *Relation) RemoveTarget(target *Object) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(target.Native())

	var ret gi.Argument

	err := relationRemoveTargetFunction_Set()
	if err == nil {
		ret = relationRemoveTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var relationSetObject *gi.Object
var relationSetObject_Once sync.Once

func relationSetObject_Set() error {
	var err error
	relationSetObject_Once.Do(func() {
		relationSetObject, err = gi.ObjectNew("Atk", "RelationSet")
	})
	return err
}

type RelationSet struct {
	native unsafe.Pointer
}

func RelationSetNewFromNative(native unsafe.Pointer) *RelationSet {
	instance := &RelationSet{native: native}

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
func (recv *RelationSet) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToRelationSet down casts any arbitrary Object to RelationSet.
Exercise care, as this is a potentially dangerous function
if the Object is not a RelationSet.
*/
func CastToRelationSet(object *gobject.Object) *RelationSet {
	return RelationSetNewFromNative(object.Native())
}

// Equals compares this RelationSet with another RelationSet, and returns true if they represent the same Object.
func (recv *RelationSet) Equals(other *RelationSet) bool {
	return other.Native() == recv.Native()
}

func (recv *RelationSet) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RelationSet) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := relationSetObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(relationSetObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RelationSet) SetFieldParent(value *gobject.Object) {
	err := relationSetObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(relationSetObject, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'relations' : for field getter : missing Type

// UNSUPPORTED : C value 'relations' : for field setter : missing Type

var relationSetNewFunction *gi.Function
var relationSetNewFunction_Once sync.Once

func relationSetNewFunction_Set() error {
	var err error
	relationSetNewFunction_Once.Do(func() {
		err = relationSetObject_Set()
		if err != nil {
			return
		}
		relationSetNewFunction, err = relationSetObject.InvokerNew("new")
	})
	return err
}

// RelationSetNew is a representation of the C type atk_relation_set_new.
func RelationSetNew() *RelationSet {

	var ret gi.Argument

	err := relationSetNewFunction_Set()
	if err == nil {
		ret = relationSetNewFunction.Invoke(nil, nil)
	}

	retGo := RelationSetNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var relationSetAddFunction *gi.Function
var relationSetAddFunction_Once sync.Once

func relationSetAddFunction_Set() error {
	var err error
	relationSetAddFunction_Once.Do(func() {
		err = relationSetObject_Set()
		if err != nil {
			return
		}
		relationSetAddFunction, err = relationSetObject.InvokerNew("add")
	})
	return err
}

// Add is a representation of the C type atk_relation_set_add.
func (recv *RelationSet) Add(relation *Relation) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(relation.Native())

	err := relationSetAddFunction_Set()
	if err == nil {
		relationSetAddFunction.Invoke(inArgs[:], nil)
	}

	return
}

var relationSetAddRelationByTypeFunction *gi.Function
var relationSetAddRelationByTypeFunction_Once sync.Once

func relationSetAddRelationByTypeFunction_Set() error {
	var err error
	relationSetAddRelationByTypeFunction_Once.Do(func() {
		err = relationSetObject_Set()
		if err != nil {
			return
		}
		relationSetAddRelationByTypeFunction, err = relationSetObject.InvokerNew("add_relation_by_type")
	})
	return err
}

// AddRelationByType is a representation of the C type atk_relation_set_add_relation_by_type.
func (recv *RelationSet) AddRelationByType(relationship RelationType, target *Object) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(relationship))
	inArgs[2].SetPointer(target.Native())

	err := relationSetAddRelationByTypeFunction_Set()
	if err == nil {
		relationSetAddRelationByTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var relationSetContainsFunction *gi.Function
var relationSetContainsFunction_Once sync.Once

func relationSetContainsFunction_Set() error {
	var err error
	relationSetContainsFunction_Once.Do(func() {
		err = relationSetObject_Set()
		if err != nil {
			return
		}
		relationSetContainsFunction, err = relationSetObject.InvokerNew("contains")
	})
	return err
}

// Contains is a representation of the C type atk_relation_set_contains.
func (recv *RelationSet) Contains(relationship RelationType) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(relationship))

	var ret gi.Argument

	err := relationSetContainsFunction_Set()
	if err == nil {
		ret = relationSetContainsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var relationSetContainsTargetFunction *gi.Function
var relationSetContainsTargetFunction_Once sync.Once

func relationSetContainsTargetFunction_Set() error {
	var err error
	relationSetContainsTargetFunction_Once.Do(func() {
		err = relationSetObject_Set()
		if err != nil {
			return
		}
		relationSetContainsTargetFunction, err = relationSetObject.InvokerNew("contains_target")
	})
	return err
}

// ContainsTarget is a representation of the C type atk_relation_set_contains_target.
func (recv *RelationSet) ContainsTarget(relationship RelationType, target *Object) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(relationship))
	inArgs[2].SetPointer(target.Native())

	var ret gi.Argument

	err := relationSetContainsTargetFunction_Set()
	if err == nil {
		ret = relationSetContainsTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var relationSetGetNRelationsFunction *gi.Function
var relationSetGetNRelationsFunction_Once sync.Once

func relationSetGetNRelationsFunction_Set() error {
	var err error
	relationSetGetNRelationsFunction_Once.Do(func() {
		err = relationSetObject_Set()
		if err != nil {
			return
		}
		relationSetGetNRelationsFunction, err = relationSetObject.InvokerNew("get_n_relations")
	})
	return err
}

// GetNRelations is a representation of the C type atk_relation_set_get_n_relations.
func (recv *RelationSet) GetNRelations() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := relationSetGetNRelationsFunction_Set()
	if err == nil {
		ret = relationSetGetNRelationsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var relationSetGetRelationFunction *gi.Function
var relationSetGetRelationFunction_Once sync.Once

func relationSetGetRelationFunction_Set() error {
	var err error
	relationSetGetRelationFunction_Once.Do(func() {
		err = relationSetObject_Set()
		if err != nil {
			return
		}
		relationSetGetRelationFunction, err = relationSetObject.InvokerNew("get_relation")
	})
	return err
}

// GetRelation is a representation of the C type atk_relation_set_get_relation.
func (recv *RelationSet) GetRelation(i int32) *Relation {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := relationSetGetRelationFunction_Set()
	if err == nil {
		ret = relationSetGetRelationFunction.Invoke(inArgs[:], nil)
	}

	retGo := RelationNewFromNative(ret.Pointer())

	return retGo
}

var relationSetGetRelationByTypeFunction *gi.Function
var relationSetGetRelationByTypeFunction_Once sync.Once

func relationSetGetRelationByTypeFunction_Set() error {
	var err error
	relationSetGetRelationByTypeFunction_Once.Do(func() {
		err = relationSetObject_Set()
		if err != nil {
			return
		}
		relationSetGetRelationByTypeFunction, err = relationSetObject.InvokerNew("get_relation_by_type")
	})
	return err
}

// GetRelationByType is a representation of the C type atk_relation_set_get_relation_by_type.
func (recv *RelationSet) GetRelationByType(relationship RelationType) *Relation {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(relationship))

	var ret gi.Argument

	err := relationSetGetRelationByTypeFunction_Set()
	if err == nil {
		ret = relationSetGetRelationByTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := RelationNewFromNative(ret.Pointer())

	return retGo
}

var relationSetRemoveFunction *gi.Function
var relationSetRemoveFunction_Once sync.Once

func relationSetRemoveFunction_Set() error {
	var err error
	relationSetRemoveFunction_Once.Do(func() {
		err = relationSetObject_Set()
		if err != nil {
			return
		}
		relationSetRemoveFunction, err = relationSetObject.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type atk_relation_set_remove.
func (recv *RelationSet) Remove(relation *Relation) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(relation.Native())

	err := relationSetRemoveFunction_Set()
	if err == nil {
		relationSetRemoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketObject *gi.Object
var socketObject_Once sync.Once

func socketObject_Set() error {
	var err error
	socketObject_Once.Do(func() {
		socketObject, err = gi.ObjectNew("Atk", "Socket")
	})
	return err
}

type Socket struct {
	native unsafe.Pointer
}

func SocketNewFromNative(native unsafe.Pointer) *Socket {
	instance := &Socket{native: native}

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

// ObjectAtk upcasts to *ObjectAtk
func (recv *Socket) ObjectAtk() *Object {
	return ObjectNewFromNative(recv.Native())
}

// Object upcasts to *Object
func (recv *Socket) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToSocket down casts any arbitrary Object to Socket.
Exercise care, as this is a potentially dangerous function
if the Object is not a Socket.
*/
func CastToSocket(object *gobject.Object) *Socket {
	return SocketNewFromNative(object.Native())
}

// Equals compares this Socket with another Socket, and returns true if they represent the same Object.
func (recv *Socket) Equals(other *Socket) bool {
	return other.Native() == recv.Native()
}

func (recv *Socket) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Socket) FieldParent() *Object {
	var nilValue *Object
	err := socketObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(socketObject, recv.Native(), "parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Socket) SetFieldParent(value *Object) {
	err := socketObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(socketObject, recv.Native(), "parent", argValue)
}

var socketNewFunction *gi.Function
var socketNewFunction_Once sync.Once

func socketNewFunction_Set() error {
	var err error
	socketNewFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketNewFunction, err = socketObject.InvokerNew("new")
	})
	return err
}

// SocketNew is a representation of the C type atk_socket_new.
func SocketNew() *Socket {

	var ret gi.Argument

	err := socketNewFunction_Set()
	if err == nil {
		ret = socketNewFunction.Invoke(nil, nil)
	}

	retGo := SocketNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var socketEmbedFunction *gi.Function
var socketEmbedFunction_Once sync.Once

func socketEmbedFunction_Set() error {
	var err error
	socketEmbedFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketEmbedFunction, err = socketObject.InvokerNew("embed")
	})
	return err
}

// Embed is a representation of the C type atk_socket_embed.
func (recv *Socket) Embed(plugId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(plugId)

	err := socketEmbedFunction_Set()
	if err == nil {
		socketEmbedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketIsOccupiedFunction *gi.Function
var socketIsOccupiedFunction_Once sync.Once

func socketIsOccupiedFunction_Set() error {
	var err error
	socketIsOccupiedFunction_Once.Do(func() {
		err = socketObject_Set()
		if err != nil {
			return
		}
		socketIsOccupiedFunction, err = socketObject.InvokerNew("is_occupied")
	})
	return err
}

// IsOccupied is a representation of the C type atk_socket_is_occupied.
func (recv *Socket) IsOccupied() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := socketIsOccupiedFunction_Set()
	if err == nil {
		ret = socketIsOccupiedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// Component returns the Component interface implemented by Socket
func (recv *Socket) Component() *Component {
	return ComponentNewFromNative(recv.Native())
}

var stateSetObject *gi.Object
var stateSetObject_Once sync.Once

func stateSetObject_Set() error {
	var err error
	stateSetObject_Once.Do(func() {
		stateSetObject, err = gi.ObjectNew("Atk", "StateSet")
	})
	return err
}

type StateSet struct {
	native unsafe.Pointer
}

func StateSetNewFromNative(native unsafe.Pointer) *StateSet {
	instance := &StateSet{native: native}

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
func (recv *StateSet) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToStateSet down casts any arbitrary Object to StateSet.
Exercise care, as this is a potentially dangerous function
if the Object is not a StateSet.
*/
func CastToStateSet(object *gobject.Object) *StateSet {
	return StateSetNewFromNative(object.Native())
}

// Equals compares this StateSet with another StateSet, and returns true if they represent the same Object.
func (recv *StateSet) Equals(other *StateSet) bool {
	return other.Native() == recv.Native()
}

func (recv *StateSet) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *StateSet) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := stateSetObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(stateSetObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *StateSet) SetFieldParent(value *gobject.Object) {
	err := stateSetObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(stateSetObject, recv.Native(), "parent", argValue)
}

var stateSetNewFunction *gi.Function
var stateSetNewFunction_Once sync.Once

func stateSetNewFunction_Set() error {
	var err error
	stateSetNewFunction_Once.Do(func() {
		err = stateSetObject_Set()
		if err != nil {
			return
		}
		stateSetNewFunction, err = stateSetObject.InvokerNew("new")
	})
	return err
}

// StateSetNew is a representation of the C type atk_state_set_new.
func StateSetNew() *StateSet {

	var ret gi.Argument

	err := stateSetNewFunction_Set()
	if err == nil {
		ret = stateSetNewFunction.Invoke(nil, nil)
	}

	retGo := StateSetNewFromNative(ret.Pointer())
	object := retGo.Object()
	object.RefSink()

	return retGo
}

var stateSetAddStateFunction *gi.Function
var stateSetAddStateFunction_Once sync.Once

func stateSetAddStateFunction_Set() error {
	var err error
	stateSetAddStateFunction_Once.Do(func() {
		err = stateSetObject_Set()
		if err != nil {
			return
		}
		stateSetAddStateFunction, err = stateSetObject.InvokerNew("add_state")
	})
	return err
}

// AddState is a representation of the C type atk_state_set_add_state.
func (recv *StateSet) AddState(type_ StateType) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(type_))

	var ret gi.Argument

	err := stateSetAddStateFunction_Set()
	if err == nil {
		ret = stateSetAddStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'atk_state_set_add_states' : parameter 'types' of type 'nil' not supported

var stateSetAndSetsFunction *gi.Function
var stateSetAndSetsFunction_Once sync.Once

func stateSetAndSetsFunction_Set() error {
	var err error
	stateSetAndSetsFunction_Once.Do(func() {
		err = stateSetObject_Set()
		if err != nil {
			return
		}
		stateSetAndSetsFunction, err = stateSetObject.InvokerNew("and_sets")
	})
	return err
}

// AndSets is a representation of the C type atk_state_set_and_sets.
func (recv *StateSet) AndSets(compareSet *StateSet) *StateSet {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(compareSet.Native())

	var ret gi.Argument

	err := stateSetAndSetsFunction_Set()
	if err == nil {
		ret = stateSetAndSetsFunction.Invoke(inArgs[:], nil)
	}

	retGo := StateSetNewFromNative(ret.Pointer())

	return retGo
}

var stateSetClearStatesFunction *gi.Function
var stateSetClearStatesFunction_Once sync.Once

func stateSetClearStatesFunction_Set() error {
	var err error
	stateSetClearStatesFunction_Once.Do(func() {
		err = stateSetObject_Set()
		if err != nil {
			return
		}
		stateSetClearStatesFunction, err = stateSetObject.InvokerNew("clear_states")
	})
	return err
}

// ClearStates is a representation of the C type atk_state_set_clear_states.
func (recv *StateSet) ClearStates() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := stateSetClearStatesFunction_Set()
	if err == nil {
		stateSetClearStatesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var stateSetContainsStateFunction *gi.Function
var stateSetContainsStateFunction_Once sync.Once

func stateSetContainsStateFunction_Set() error {
	var err error
	stateSetContainsStateFunction_Once.Do(func() {
		err = stateSetObject_Set()
		if err != nil {
			return
		}
		stateSetContainsStateFunction, err = stateSetObject.InvokerNew("contains_state")
	})
	return err
}

// ContainsState is a representation of the C type atk_state_set_contains_state.
func (recv *StateSet) ContainsState(type_ StateType) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(type_))

	var ret gi.Argument

	err := stateSetContainsStateFunction_Set()
	if err == nil {
		ret = stateSetContainsStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'atk_state_set_contains_states' : parameter 'types' of type 'nil' not supported

var stateSetIsEmptyFunction *gi.Function
var stateSetIsEmptyFunction_Once sync.Once

func stateSetIsEmptyFunction_Set() error {
	var err error
	stateSetIsEmptyFunction_Once.Do(func() {
		err = stateSetObject_Set()
		if err != nil {
			return
		}
		stateSetIsEmptyFunction, err = stateSetObject.InvokerNew("is_empty")
	})
	return err
}

// IsEmpty is a representation of the C type atk_state_set_is_empty.
func (recv *StateSet) IsEmpty() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := stateSetIsEmptyFunction_Set()
	if err == nil {
		ret = stateSetIsEmptyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var stateSetOrSetsFunction *gi.Function
var stateSetOrSetsFunction_Once sync.Once

func stateSetOrSetsFunction_Set() error {
	var err error
	stateSetOrSetsFunction_Once.Do(func() {
		err = stateSetObject_Set()
		if err != nil {
			return
		}
		stateSetOrSetsFunction, err = stateSetObject.InvokerNew("or_sets")
	})
	return err
}

// OrSets is a representation of the C type atk_state_set_or_sets.
func (recv *StateSet) OrSets(compareSet *StateSet) *StateSet {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(compareSet.Native())

	var ret gi.Argument

	err := stateSetOrSetsFunction_Set()
	if err == nil {
		ret = stateSetOrSetsFunction.Invoke(inArgs[:], nil)
	}

	retGo := StateSetNewFromNative(ret.Pointer())

	return retGo
}

var stateSetRemoveStateFunction *gi.Function
var stateSetRemoveStateFunction_Once sync.Once

func stateSetRemoveStateFunction_Set() error {
	var err error
	stateSetRemoveStateFunction_Once.Do(func() {
		err = stateSetObject_Set()
		if err != nil {
			return
		}
		stateSetRemoveStateFunction, err = stateSetObject.InvokerNew("remove_state")
	})
	return err
}

// RemoveState is a representation of the C type atk_state_set_remove_state.
func (recv *StateSet) RemoveState(type_ StateType) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(type_))

	var ret gi.Argument

	err := stateSetRemoveStateFunction_Set()
	if err == nil {
		ret = stateSetRemoveStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var stateSetXorSetsFunction *gi.Function
var stateSetXorSetsFunction_Once sync.Once

func stateSetXorSetsFunction_Set() error {
	var err error
	stateSetXorSetsFunction_Once.Do(func() {
		err = stateSetObject_Set()
		if err != nil {
			return
		}
		stateSetXorSetsFunction, err = stateSetObject.InvokerNew("xor_sets")
	})
	return err
}

// XorSets is a representation of the C type atk_state_set_xor_sets.
func (recv *StateSet) XorSets(compareSet *StateSet) *StateSet {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(compareSet.Native())

	var ret gi.Argument

	err := stateSetXorSetsFunction_Set()
	if err == nil {
		ret = stateSetXorSetsFunction.Invoke(inArgs[:], nil)
	}

	retGo := StateSetNewFromNative(ret.Pointer())

	return retGo
}

var utilObject *gi.Object
var utilObject_Once sync.Once

func utilObject_Set() error {
	var err error
	utilObject_Once.Do(func() {
		utilObject, err = gi.ObjectNew("Atk", "Util")
	})
	return err
}

type Util struct {
	native unsafe.Pointer
}

func UtilNewFromNative(native unsafe.Pointer) *Util {
	instance := &Util{native: native}

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
func (recv *Util) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.Native())
}

/*
CastToUtil down casts any arbitrary Object to Util.
Exercise care, as this is a potentially dangerous function
if the Object is not a Util.
*/
func CastToUtil(object *gobject.Object) *Util {
	return UtilNewFromNative(object.Native())
}

// Equals compares this Util with another Util, and returns true if they represent the same Object.
func (recv *Util) Equals(other *Util) bool {
	return other.Native() == recv.Native()
}

func (recv *Util) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *Util) FieldParent() *gobject.Object {
	var nilValue *gobject.Object
	err := utilObject_Set()
	if err != nil {
		return nilValue
	}

	argValue := gi.ObjectFieldGet(utilObject, recv.Native(), "parent")
	value := gobject.ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Util) SetFieldParent(value *gobject.Object) {
	err := utilObject_Set()
	if err != nil {
		return
	}

	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.ObjectFieldSet(utilObject, recv.Native(), "parent", argValue)
}
