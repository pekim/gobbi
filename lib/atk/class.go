// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
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
	return &GObjectAccessible{native: native}
}

// ObjectAtk upcasts to *ObjectAtk
func (recv *GObjectAccessible) ObjectAtk() *Object {
	return ObjectNewFromNative(recv.native)
}

// Object upcasts to *Object
func (recv *GObjectAccessible) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// FieldParent returns the C field 'parent'.
func (recv *GObjectAccessible) FieldParent() *Object {
	argValue := gi.ObjectFieldGet(gObjectAccessibleObject, recv.native, "parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *GObjectAccessible) SetFieldParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.ObjectFieldSet(gObjectAccessibleObject, recv.native, "parent", argValue)
}

// UNSUPPORTED : C value 'atk_gobject_accessible_get_object' : return type 'GObject.Object' not supported

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
	return &Hyperlink{native: native}
}

// Object upcasts to *Object
func (recv *Hyperlink) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := hyperlinkIsValidFunction_Set()
	if err == nil {
		ret = hyperlinkIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
	return &Misc{native: native}
}

// Object upcasts to *Object
func (recv *Misc) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	return &NoOpObject{native: native}
}

// ObjectAtk upcasts to *ObjectAtk
func (recv *NoOpObject) ObjectAtk() *Object {
	return ObjectNewFromNative(recv.native)
}

// Object upcasts to *Object
func (recv *NoOpObject) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// FieldParent returns the C field 'parent'.
func (recv *NoOpObject) FieldParent() *Object {
	argValue := gi.ObjectFieldGet(noOpObjectObject, recv.native, "parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *NoOpObject) SetFieldParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.ObjectFieldSet(noOpObjectObject, recv.native, "parent", argValue)
}

// UNSUPPORTED : C value 'atk_no_op_object_new' : parameter 'obj' of type 'GObject.Object' not supported

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
	return &NoOpObjectFactory{native: native}
}

// ObjectFactory upcasts to *ObjectFactory
func (recv *NoOpObjectFactory) ObjectFactory() *ObjectFactory {
	return ObjectFactoryNewFromNative(recv.native)
}

// Object upcasts to *Object
func (recv *NoOpObjectFactory) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// FieldParent returns the C field 'parent'.
func (recv *NoOpObjectFactory) FieldParent() *ObjectFactory {
	argValue := gi.ObjectFieldGet(noOpObjectFactoryObject, recv.native, "parent")
	value := ObjectFactoryNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *NoOpObjectFactory) SetFieldParent(value *ObjectFactory) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.ObjectFieldSet(noOpObjectFactoryObject, recv.native, "parent", argValue)
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
	return &Object{native: native}
}

// Object upcasts to *Object
func (recv *Object) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldDescription returns the C field 'description'.
func (recv *Object) FieldDescription() string {
	argValue := gi.ObjectFieldGet(objectObject, recv.native, "description")
	value := argValue.String(false)
	return value
}

// SetFieldDescription sets the value of the C field 'description'.
func (recv *Object) SetFieldDescription(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(objectObject, recv.native, "description", argValue)
}

// FieldName returns the C field 'name'.
func (recv *Object) FieldName() string {
	argValue := gi.ObjectFieldGet(objectObject, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *Object) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.ObjectFieldSet(objectObject, recv.native, "name", argValue)
}

// FieldAccessibleParent returns the C field 'accessible_parent'.
func (recv *Object) FieldAccessibleParent() *Object {
	argValue := gi.ObjectFieldGet(objectObject, recv.native, "accessible_parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAccessibleParent sets the value of the C field 'accessible_parent'.
func (recv *Object) SetFieldAccessibleParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.ObjectFieldSet(objectObject, recv.native, "accessible_parent", argValue)
}

// FieldRole returns the C field 'role'.
func (recv *Object) FieldRole() Role {
	argValue := gi.ObjectFieldGet(objectObject, recv.native, "role")
	value := Role(argValue.Int32())
	return value
}

// SetFieldRole sets the value of the C field 'role'.
func (recv *Object) SetFieldRole(value Role) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.ObjectFieldSet(objectObject, recv.native, "role", argValue)
}

// FieldRelationSet returns the C field 'relation_set'.
func (recv *Object) FieldRelationSet() *RelationSet {
	argValue := gi.ObjectFieldGet(objectObject, recv.native, "relation_set")
	value := RelationSetNewFromNative(argValue.Pointer())
	return value
}

// SetFieldRelationSet sets the value of the C field 'relation_set'.
func (recv *Object) SetFieldRelationSet(value *RelationSet) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.ObjectFieldSet(objectObject, recv.native, "relation_set", argValue)
}

// FieldLayer returns the C field 'layer'.
func (recv *Object) FieldLayer() Layer {
	argValue := gi.ObjectFieldGet(objectObject, recv.native, "layer")
	value := Layer(argValue.Int32())
	return value
}

// SetFieldLayer sets the value of the C field 'layer'.
func (recv *Object) SetFieldLayer(value Layer) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.ObjectFieldSet(objectObject, recv.native, "layer", argValue)
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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(relationship))
	inArgs[2].SetPointer(target.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := objectGetRoleFunction_Set()
	if err == nil {
		ret = objectGetRoleFunction.Invoke(inArgs[:], nil)
	}

	retGo := Role(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'atk_object_initialize' : parameter 'data' of type 'gpointer' not supported

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(relationship))
	inArgs[2].SetPointer(target.native)

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(parent.native)

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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(role))

	err := objectSetRoleFunction_Set()
	if err == nil {
		objectSetRoleFunction.Invoke(inArgs[:], nil)
	}

	return
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
	return &ObjectFactory{native: native}
}

// Object upcasts to *Object
func (recv *ObjectFactory) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'atk_object_factory_create_accessible' : parameter 'obj' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'atk_object_factory_get_accessible_type' : return type 'GType' not supported

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
	inArgs[0].SetPointer(recv.native)

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
	return &Plug{native: native}
}

// ObjectAtk upcasts to *ObjectAtk
func (recv *Plug) ObjectAtk() *Object {
	return ObjectNewFromNative(recv.native)
}

// Object upcasts to *Object
func (recv *Plug) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// FieldParent returns the C field 'parent'.
func (recv *Plug) FieldParent() *Object {
	argValue := gi.ObjectFieldGet(plugObject, recv.native, "parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Plug) SetFieldParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.ObjectFieldSet(plugObject, recv.native, "parent", argValue)
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
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := plugGetIdFunction_Set()
	if err == nil {
		ret = plugGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
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
	return &Registry{native: native}
}

// Object upcasts to *Object
func (recv *Registry) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'factory_type_registry' : for field getter : no Go type for 'GLib.HashTable'

// UNSUPPORTED : C value 'factory_type_registry' : for field setter : no Go type for 'GLib.HashTable'

// UNSUPPORTED : C value 'factory_singleton_cache' : for field getter : no Go type for 'GLib.HashTable'

// UNSUPPORTED : C value 'factory_singleton_cache' : for field setter : no Go type for 'GLib.HashTable'

// UNSUPPORTED : C value 'atk_registry_get_factory' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'atk_registry_get_factory_type' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'atk_registry_set_factory_type' : parameter 'type' of type 'GType' not supported

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
	return &Relation{native: native}
}

// Object upcasts to *Object
func (recv *Relation) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'target' : for field getter : missing Type

// UNSUPPORTED : C value 'target' : for field setter : missing Type

// FieldRelationship returns the C field 'relationship'.
func (recv *Relation) FieldRelationship() RelationType {
	argValue := gi.ObjectFieldGet(relationObject, recv.native, "relationship")
	value := RelationType(argValue.Int32())
	return value
}

// SetFieldRelationship sets the value of the C field 'relationship'.
func (recv *Relation) SetFieldRelationship(value RelationType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.ObjectFieldSet(relationObject, recv.native, "relationship", argValue)
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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(target.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(target.native)

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
	return &RelationSet{native: native}
}

// Object upcasts to *Object
func (recv *RelationSet) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(relation.native)

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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(relationship))
	inArgs[2].SetPointer(target.native)

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(relationship))
	inArgs[2].SetPointer(target.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(relation.native)

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
	return &Socket{native: native}
}

// ObjectAtk upcasts to *ObjectAtk
func (recv *Socket) ObjectAtk() *Object {
	return ObjectNewFromNative(recv.native)
}

// Object upcasts to *Object
func (recv *Socket) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// FieldParent returns the C field 'parent'.
func (recv *Socket) FieldParent() *Object {
	argValue := gi.ObjectFieldGet(socketObject, recv.native, "parent")
	value := ObjectNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Socket) SetFieldParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.ObjectFieldSet(socketObject, recv.native, "parent", argValue)
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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := socketIsOccupiedFunction_Set()
	if err == nil {
		ret = socketIsOccupiedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
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
	return &StateSet{native: native}
}

// Object upcasts to *Object
func (recv *StateSet) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(compareSet.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(compareSet.native)

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
	inArgs[0].SetPointer(recv.native)
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
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(compareSet.native)

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
	return &Util{native: native}
}

// Object upcasts to *Object
func (recv *Util) Object() *gobject.Object {
	return gobject.ObjectNewFromNative(recv.native)
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'
