// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
)

var gObjectAccessibleStruct *gi.Struct
var gObjectAccessibleStruct_Once sync.Once

func gObjectAccessibleStruct_Set() error {
	var err error
	gObjectAccessibleStruct_Once.Do(func() {
		gObjectAccessibleStruct, err = gi.StructNew("Atk", "GObjectAccessible")
	})
	return err
}

type GObjectAccessible struct {
	Object
}

// FieldParent returns the C field 'parent'.
func (recv *GObjectAccessible) FieldParent() *Object {
	argValue := gi.FieldGet(gObjectAccessibleStruct, recv.Native, "parent")
	value := &Object{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *GObjectAccessible) SetFieldParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(gObjectAccessibleStruct, recv.Native, "parent", argValue)
}

// UNSUPPORTED : C value 'atk_gobject_accessible_get_object' : return type 'GObject.Object' not supported

// GObjectAccessibleStruct creates an uninitialised GObjectAccessible.
func GObjectAccessibleStruct() *GObjectAccessible {
	err := gObjectAccessibleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GObjectAccessible{}
	structGo.Native = gObjectAccessibleStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeGObjectAccessible)
	return structGo
}
func finalizeGObjectAccessible(obj *GObjectAccessible) {
	gObjectAccessibleStruct.Free(obj.Native)
}

var hyperlinkStruct *gi.Struct
var hyperlinkStruct_Once sync.Once

func hyperlinkStruct_Set() error {
	var err error
	hyperlinkStruct_Once.Do(func() {
		hyperlinkStruct, err = gi.StructNew("Atk", "Hyperlink")
	})
	return err
}

type Hyperlink struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var hyperlinkGetEndIndexFunction *gi.Function
var hyperlinkGetEndIndexFunction_Once sync.Once

func hyperlinkGetEndIndexFunction_Set() error {
	var err error
	hyperlinkGetEndIndexFunction_Once.Do(func() {
		err = hyperlinkStruct_Set()
		if err != nil {
			return
		}
		hyperlinkGetEndIndexFunction, err = hyperlinkStruct.InvokerNew("get_end_index")
	})
	return err
}

// GetEndIndex is a representation of the C type atk_hyperlink_get_end_index.
func (recv *Hyperlink) GetEndIndex() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = hyperlinkStruct_Set()
		if err != nil {
			return
		}
		hyperlinkGetNAnchorsFunction, err = hyperlinkStruct.InvokerNew("get_n_anchors")
	})
	return err
}

// GetNAnchors is a representation of the C type atk_hyperlink_get_n_anchors.
func (recv *Hyperlink) GetNAnchors() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = hyperlinkStruct_Set()
		if err != nil {
			return
		}
		hyperlinkGetObjectFunction, err = hyperlinkStruct.InvokerNew("get_object")
	})
	return err
}

// GetObject is a representation of the C type atk_hyperlink_get_object.
func (recv *Hyperlink) GetObject(i int32) *Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := hyperlinkGetObjectFunction_Set()
	if err == nil {
		ret = hyperlinkGetObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Object{}
	retGo.Native = ret.Pointer()

	return retGo
}

var hyperlinkGetStartIndexFunction *gi.Function
var hyperlinkGetStartIndexFunction_Once sync.Once

func hyperlinkGetStartIndexFunction_Set() error {
	var err error
	hyperlinkGetStartIndexFunction_Once.Do(func() {
		err = hyperlinkStruct_Set()
		if err != nil {
			return
		}
		hyperlinkGetStartIndexFunction, err = hyperlinkStruct.InvokerNew("get_start_index")
	})
	return err
}

// GetStartIndex is a representation of the C type atk_hyperlink_get_start_index.
func (recv *Hyperlink) GetStartIndex() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = hyperlinkStruct_Set()
		if err != nil {
			return
		}
		hyperlinkGetUriFunction, err = hyperlinkStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type atk_hyperlink_get_uri.
func (recv *Hyperlink) GetUri(i int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
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
		err = hyperlinkStruct_Set()
		if err != nil {
			return
		}
		hyperlinkIsInlineFunction, err = hyperlinkStruct.InvokerNew("is_inline")
	})
	return err
}

// IsInline is a representation of the C type atk_hyperlink_is_inline.
func (recv *Hyperlink) IsInline() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = hyperlinkStruct_Set()
		if err != nil {
			return
		}
		hyperlinkIsSelectedLinkFunction, err = hyperlinkStruct.InvokerNew("is_selected_link")
	})
	return err
}

// IsSelectedLink is a representation of the C type atk_hyperlink_is_selected_link.
func (recv *Hyperlink) IsSelectedLink() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = hyperlinkStruct_Set()
		if err != nil {
			return
		}
		hyperlinkIsValidFunction, err = hyperlinkStruct.InvokerNew("is_valid")
	})
	return err
}

// IsValid is a representation of the C type atk_hyperlink_is_valid.
func (recv *Hyperlink) IsValid() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := hyperlinkIsValidFunction_Set()
	if err == nil {
		ret = hyperlinkIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// HyperlinkStruct creates an uninitialised Hyperlink.
func HyperlinkStruct() *Hyperlink {
	err := hyperlinkStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Hyperlink{}
	structGo.Native = hyperlinkStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeHyperlink)
	return structGo
}
func finalizeHyperlink(obj *Hyperlink) {
	hyperlinkStruct.Free(obj.Native)
}

var miscStruct *gi.Struct
var miscStruct_Once sync.Once

func miscStruct_Set() error {
	var err error
	miscStruct_Once.Do(func() {
		miscStruct, err = gi.StructNew("Atk", "Misc")
	})
	return err
}

type Misc struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var miscThreadsEnterFunction *gi.Function
var miscThreadsEnterFunction_Once sync.Once

func miscThreadsEnterFunction_Set() error {
	var err error
	miscThreadsEnterFunction_Once.Do(func() {
		err = miscStruct_Set()
		if err != nil {
			return
		}
		miscThreadsEnterFunction, err = miscStruct.InvokerNew("threads_enter")
	})
	return err
}

// ThreadsEnter is a representation of the C type atk_misc_threads_enter.
func (recv *Misc) ThreadsEnter() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = miscStruct_Set()
		if err != nil {
			return
		}
		miscThreadsLeaveFunction, err = miscStruct.InvokerNew("threads_leave")
	})
	return err
}

// ThreadsLeave is a representation of the C type atk_misc_threads_leave.
func (recv *Misc) ThreadsLeave() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := miscThreadsLeaveFunction_Set()
	if err == nil {
		miscThreadsLeaveFunction.Invoke(inArgs[:], nil)
	}

	return
}

// MiscStruct creates an uninitialised Misc.
func MiscStruct() *Misc {
	err := miscStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Misc{}
	structGo.Native = miscStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMisc)
	return structGo
}
func finalizeMisc(obj *Misc) {
	miscStruct.Free(obj.Native)
}

var noOpObjectStruct *gi.Struct
var noOpObjectStruct_Once sync.Once

func noOpObjectStruct_Set() error {
	var err error
	noOpObjectStruct_Once.Do(func() {
		noOpObjectStruct, err = gi.StructNew("Atk", "NoOpObject")
	})
	return err
}

type NoOpObject struct {
	Object
}

// FieldParent returns the C field 'parent'.
func (recv *NoOpObject) FieldParent() *Object {
	argValue := gi.FieldGet(noOpObjectStruct, recv.Native, "parent")
	value := &Object{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *NoOpObject) SetFieldParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(noOpObjectStruct, recv.Native, "parent", argValue)
}

// UNSUPPORTED : C value 'atk_no_op_object_new' : parameter 'obj' of type 'GObject.Object' not supported

var noOpObjectFactoryStruct *gi.Struct
var noOpObjectFactoryStruct_Once sync.Once

func noOpObjectFactoryStruct_Set() error {
	var err error
	noOpObjectFactoryStruct_Once.Do(func() {
		noOpObjectFactoryStruct, err = gi.StructNew("Atk", "NoOpObjectFactory")
	})
	return err
}

type NoOpObjectFactory struct {
	ObjectFactory
}

// FieldParent returns the C field 'parent'.
func (recv *NoOpObjectFactory) FieldParent() *ObjectFactory {
	argValue := gi.FieldGet(noOpObjectFactoryStruct, recv.Native, "parent")
	value := &ObjectFactory{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *NoOpObjectFactory) SetFieldParent(value *ObjectFactory) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(noOpObjectFactoryStruct, recv.Native, "parent", argValue)
}

var noOpObjectFactoryNewFunction *gi.Function
var noOpObjectFactoryNewFunction_Once sync.Once

func noOpObjectFactoryNewFunction_Set() error {
	var err error
	noOpObjectFactoryNewFunction_Once.Do(func() {
		err = noOpObjectFactoryStruct_Set()
		if err != nil {
			return
		}
		noOpObjectFactoryNewFunction, err = noOpObjectFactoryStruct.InvokerNew("new")
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

	retGo := &NoOpObjectFactory{}
	retGo.Native = ret.Pointer()

	return retGo
}

var objectStruct *gi.Struct
var objectStruct_Once sync.Once

func objectStruct_Set() error {
	var err error
	objectStruct_Once.Do(func() {
		objectStruct, err = gi.StructNew("Atk", "Object")
	})
	return err
}

type Object struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// FieldDescription returns the C field 'description'.
func (recv *Object) FieldDescription() string {
	argValue := gi.FieldGet(objectStruct, recv.Native, "description")
	value := argValue.String(false)
	return value
}

// SetFieldDescription sets the value of the C field 'description'.
func (recv *Object) SetFieldDescription(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(objectStruct, recv.Native, "description", argValue)
}

// FieldName returns the C field 'name'.
func (recv *Object) FieldName() string {
	argValue := gi.FieldGet(objectStruct, recv.Native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *Object) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(objectStruct, recv.Native, "name", argValue)
}

// FieldAccessibleParent returns the C field 'accessible_parent'.
func (recv *Object) FieldAccessibleParent() *Object {
	argValue := gi.FieldGet(objectStruct, recv.Native, "accessible_parent")
	value := &Object{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldAccessibleParent sets the value of the C field 'accessible_parent'.
func (recv *Object) SetFieldAccessibleParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(objectStruct, recv.Native, "accessible_parent", argValue)
}

// UNSUPPORTED : C value 'role' : for field getter : no Go type for 'Role'

// UNSUPPORTED : C value 'role' : for field setter : no Go type for 'Role'

// FieldRelationSet returns the C field 'relation_set'.
func (recv *Object) FieldRelationSet() *RelationSet {
	argValue := gi.FieldGet(objectStruct, recv.Native, "relation_set")
	value := &RelationSet{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldRelationSet sets the value of the C field 'relation_set'.
func (recv *Object) SetFieldRelationSet(value *RelationSet) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(objectStruct, recv.Native, "relation_set", argValue)
}

// UNSUPPORTED : C value 'layer' : for field getter : no Go type for 'Layer'

// UNSUPPORTED : C value 'layer' : for field setter : no Go type for 'Layer'

// UNSUPPORTED : C value 'atk_object_add_relationship' : parameter 'relationship' of type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_object_connect_property_change_handler' : parameter 'handler' of type 'PropertyChangeHandler' not supported

var objectGetAccessibleIdFunction *gi.Function
var objectGetAccessibleIdFunction_Once sync.Once

func objectGetAccessibleIdFunction_Set() error {
	var err error
	objectGetAccessibleIdFunction_Once.Do(func() {
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectGetAccessibleIdFunction, err = objectStruct.InvokerNew("get_accessible_id")
	})
	return err
}

// GetAccessibleId is a representation of the C type atk_object_get_accessible_id.
func (recv *Object) GetAccessibleId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectGetDescriptionFunction, err = objectStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type atk_object_get_description.
func (recv *Object) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectGetIndexInParentFunction, err = objectStruct.InvokerNew("get_index_in_parent")
	})
	return err
}

// GetIndexInParent is a representation of the C type atk_object_get_index_in_parent.
func (recv *Object) GetIndexInParent() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := objectGetIndexInParentFunction_Set()
	if err == nil {
		ret = objectGetIndexInParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'atk_object_get_layer' : return type 'Layer' not supported

var objectGetMdiZorderFunction *gi.Function
var objectGetMdiZorderFunction_Once sync.Once

func objectGetMdiZorderFunction_Set() error {
	var err error
	objectGetMdiZorderFunction_Once.Do(func() {
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectGetMdiZorderFunction, err = objectStruct.InvokerNew("get_mdi_zorder")
	})
	return err
}

// GetMdiZorder is a representation of the C type atk_object_get_mdi_zorder.
func (recv *Object) GetMdiZorder() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectGetNAccessibleChildrenFunction, err = objectStruct.InvokerNew("get_n_accessible_children")
	})
	return err
}

// GetNAccessibleChildren is a representation of the C type atk_object_get_n_accessible_children.
func (recv *Object) GetNAccessibleChildren() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectGetNameFunction, err = objectStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type atk_object_get_name.
func (recv *Object) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectGetObjectLocaleFunction, err = objectStruct.InvokerNew("get_object_locale")
	})
	return err
}

// GetObjectLocale is a representation of the C type atk_object_get_object_locale.
func (recv *Object) GetObjectLocale() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectGetParentFunction, err = objectStruct.InvokerNew("get_parent")
	})
	return err
}

// GetParent is a representation of the C type atk_object_get_parent.
func (recv *Object) GetParent() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := objectGetParentFunction_Set()
	if err == nil {
		ret = objectGetParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Object{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'atk_object_get_role' : return type 'Role' not supported

// UNSUPPORTED : C value 'atk_object_initialize' : parameter 'data' of type 'gpointer' not supported

var objectNotifyStateChangeFunction *gi.Function
var objectNotifyStateChangeFunction_Once sync.Once

func objectNotifyStateChangeFunction_Set() error {
	var err error
	objectNotifyStateChangeFunction_Once.Do(func() {
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectNotifyStateChangeFunction, err = objectStruct.InvokerNew("notify_state_change")
	})
	return err
}

// NotifyStateChange is a representation of the C type atk_object_notify_state_change.
func (recv *Object) NotifyStateChange(state State, value bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectPeekParentFunction, err = objectStruct.InvokerNew("peek_parent")
	})
	return err
}

// PeekParent is a representation of the C type atk_object_peek_parent.
func (recv *Object) PeekParent() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := objectPeekParentFunction_Set()
	if err == nil {
		ret = objectPeekParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Object{}
	retGo.Native = ret.Pointer()

	return retGo
}

var objectRefAccessibleChildFunction *gi.Function
var objectRefAccessibleChildFunction_Once sync.Once

func objectRefAccessibleChildFunction_Set() error {
	var err error
	objectRefAccessibleChildFunction_Once.Do(func() {
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectRefAccessibleChildFunction, err = objectStruct.InvokerNew("ref_accessible_child")
	})
	return err
}

// RefAccessibleChild is a representation of the C type atk_object_ref_accessible_child.
func (recv *Object) RefAccessibleChild(i int32) *Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := objectRefAccessibleChildFunction_Set()
	if err == nil {
		ret = objectRefAccessibleChildFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Object{}
	retGo.Native = ret.Pointer()

	return retGo
}

var objectRefRelationSetFunction *gi.Function
var objectRefRelationSetFunction_Once sync.Once

func objectRefRelationSetFunction_Set() error {
	var err error
	objectRefRelationSetFunction_Once.Do(func() {
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectRefRelationSetFunction, err = objectStruct.InvokerNew("ref_relation_set")
	})
	return err
}

// RefRelationSet is a representation of the C type atk_object_ref_relation_set.
func (recv *Object) RefRelationSet() *RelationSet {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := objectRefRelationSetFunction_Set()
	if err == nil {
		ret = objectRefRelationSetFunction.Invoke(inArgs[:], nil)
	}

	retGo := &RelationSet{}
	retGo.Native = ret.Pointer()

	return retGo
}

var objectRefStateSetFunction *gi.Function
var objectRefStateSetFunction_Once sync.Once

func objectRefStateSetFunction_Set() error {
	var err error
	objectRefStateSetFunction_Once.Do(func() {
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectRefStateSetFunction, err = objectStruct.InvokerNew("ref_state_set")
	})
	return err
}

// RefStateSet is a representation of the C type atk_object_ref_state_set.
func (recv *Object) RefStateSet() *StateSet {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := objectRefStateSetFunction_Set()
	if err == nil {
		ret = objectRefStateSetFunction.Invoke(inArgs[:], nil)
	}

	retGo := &StateSet{}
	retGo.Native = ret.Pointer()

	return retGo
}

var objectRemovePropertyChangeHandlerFunction *gi.Function
var objectRemovePropertyChangeHandlerFunction_Once sync.Once

func objectRemovePropertyChangeHandlerFunction_Set() error {
	var err error
	objectRemovePropertyChangeHandlerFunction_Once.Do(func() {
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectRemovePropertyChangeHandlerFunction, err = objectStruct.InvokerNew("remove_property_change_handler")
	})
	return err
}

// RemovePropertyChangeHandler is a representation of the C type atk_object_remove_property_change_handler.
func (recv *Object) RemovePropertyChangeHandler(handlerId uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(handlerId)

	err := objectRemovePropertyChangeHandlerFunction_Set()
	if err == nil {
		objectRemovePropertyChangeHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'atk_object_remove_relationship' : parameter 'relationship' of type 'RelationType' not supported

var objectSetAccessibleIdFunction *gi.Function
var objectSetAccessibleIdFunction_Once sync.Once

func objectSetAccessibleIdFunction_Set() error {
	var err error
	objectSetAccessibleIdFunction_Once.Do(func() {
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectSetAccessibleIdFunction, err = objectStruct.InvokerNew("set_accessible_id")
	})
	return err
}

// SetAccessibleId is a representation of the C type atk_object_set_accessible_id.
func (recv *Object) SetAccessibleId(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectSetDescriptionFunction, err = objectStruct.InvokerNew("set_description")
	})
	return err
}

// SetDescription is a representation of the C type atk_object_set_description.
func (recv *Object) SetDescription(description string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectSetNameFunction, err = objectStruct.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type atk_object_set_name.
func (recv *Object) SetName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
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
		err = objectStruct_Set()
		if err != nil {
			return
		}
		objectSetParentFunction, err = objectStruct.InvokerNew("set_parent")
	})
	return err
}

// SetParent is a representation of the C type atk_object_set_parent.
func (recv *Object) SetParent(parent *Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(parent.Native)

	err := objectSetParentFunction_Set()
	if err == nil {
		objectSetParentFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'atk_object_set_role' : parameter 'role' of type 'Role' not supported

// ObjectStruct creates an uninitialised Object.
func ObjectStruct() *Object {
	err := objectStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Object{}
	structGo.Native = objectStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeObject)
	return structGo
}
func finalizeObject(obj *Object) {
	objectStruct.Free(obj.Native)
}

var objectFactoryStruct *gi.Struct
var objectFactoryStruct_Once sync.Once

func objectFactoryStruct_Set() error {
	var err error
	objectFactoryStruct_Once.Do(func() {
		objectFactoryStruct, err = gi.StructNew("Atk", "ObjectFactory")
	})
	return err
}

type ObjectFactory struct {
	gobject.Object
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
		err = objectFactoryStruct_Set()
		if err != nil {
			return
		}
		objectFactoryInvalidateFunction, err = objectFactoryStruct.InvokerNew("invalidate")
	})
	return err
}

// Invalidate is a representation of the C type atk_object_factory_invalidate.
func (recv *ObjectFactory) Invalidate() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := objectFactoryInvalidateFunction_Set()
	if err == nil {
		objectFactoryInvalidateFunction.Invoke(inArgs[:], nil)
	}

	return
}

// ObjectFactoryStruct creates an uninitialised ObjectFactory.
func ObjectFactoryStruct() *ObjectFactory {
	err := objectFactoryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ObjectFactory{}
	structGo.Native = objectFactoryStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeObjectFactory)
	return structGo
}
func finalizeObjectFactory(obj *ObjectFactory) {
	objectFactoryStruct.Free(obj.Native)
}

var plugStruct *gi.Struct
var plugStruct_Once sync.Once

func plugStruct_Set() error {
	var err error
	plugStruct_Once.Do(func() {
		plugStruct, err = gi.StructNew("Atk", "Plug")
	})
	return err
}

type Plug struct {
	Object
}

// FieldParent returns the C field 'parent'.
func (recv *Plug) FieldParent() *Object {
	argValue := gi.FieldGet(plugStruct, recv.Native, "parent")
	value := &Object{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Plug) SetFieldParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(plugStruct, recv.Native, "parent", argValue)
}

var plugNewFunction *gi.Function
var plugNewFunction_Once sync.Once

func plugNewFunction_Set() error {
	var err error
	plugNewFunction_Once.Do(func() {
		err = plugStruct_Set()
		if err != nil {
			return
		}
		plugNewFunction, err = plugStruct.InvokerNew("new")
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

	retGo := &Plug{}
	retGo.Native = ret.Pointer()

	return retGo
}

var plugGetIdFunction *gi.Function
var plugGetIdFunction_Once sync.Once

func plugGetIdFunction_Set() error {
	var err error
	plugGetIdFunction_Once.Do(func() {
		err = plugStruct_Set()
		if err != nil {
			return
		}
		plugGetIdFunction, err = plugStruct.InvokerNew("get_id")
	})
	return err
}

// GetId is a representation of the C type atk_plug_get_id.
func (recv *Plug) GetId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := plugGetIdFunction_Set()
	if err == nil {
		ret = plugGetIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var registryStruct *gi.Struct
var registryStruct_Once sync.Once

func registryStruct_Set() error {
	var err error
	registryStruct_Once.Do(func() {
		registryStruct, err = gi.StructNew("Atk", "Registry")
	})
	return err
}

type Registry struct {
	gobject.Object
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

// RegistryStruct creates an uninitialised Registry.
func RegistryStruct() *Registry {
	err := registryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Registry{}
	structGo.Native = registryStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeRegistry)
	return structGo
}
func finalizeRegistry(obj *Registry) {
	registryStruct.Free(obj.Native)
}

var relationStruct *gi.Struct
var relationStruct_Once sync.Once

func relationStruct_Set() error {
	var err error
	relationStruct_Once.Do(func() {
		relationStruct, err = gi.StructNew("Atk", "Relation")
	})
	return err
}

type Relation struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'target' : for field getter : missing Type

// UNSUPPORTED : C value 'target' : for field setter : missing Type

// UNSUPPORTED : C value 'relationship' : for field getter : no Go type for 'RelationType'

// UNSUPPORTED : C value 'relationship' : for field setter : no Go type for 'RelationType'

// UNSUPPORTED : C value 'atk_relation_new' : parameter 'targets' of type 'nil' not supported

var relationAddTargetFunction *gi.Function
var relationAddTargetFunction_Once sync.Once

func relationAddTargetFunction_Set() error {
	var err error
	relationAddTargetFunction_Once.Do(func() {
		err = relationStruct_Set()
		if err != nil {
			return
		}
		relationAddTargetFunction, err = relationStruct.InvokerNew("add_target")
	})
	return err
}

// AddTarget is a representation of the C type atk_relation_add_target.
func (recv *Relation) AddTarget(target *Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(target.Native)

	err := relationAddTargetFunction_Set()
	if err == nil {
		relationAddTargetFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'atk_relation_get_relation_type' : return type 'RelationType' not supported

var relationGetTargetFunction *gi.Function
var relationGetTargetFunction_Once sync.Once

func relationGetTargetFunction_Set() error {
	var err error
	relationGetTargetFunction_Once.Do(func() {
		err = relationStruct_Set()
		if err != nil {
			return
		}
		relationGetTargetFunction, err = relationStruct.InvokerNew("get_target")
	})
	return err
}

// GetTarget is a representation of the C type atk_relation_get_target.
func (recv *Relation) GetTarget() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = relationStruct_Set()
		if err != nil {
			return
		}
		relationRemoveTargetFunction, err = relationStruct.InvokerNew("remove_target")
	})
	return err
}

// RemoveTarget is a representation of the C type atk_relation_remove_target.
func (recv *Relation) RemoveTarget(target *Object) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(target.Native)

	var ret gi.Argument

	err := relationRemoveTargetFunction_Set()
	if err == nil {
		ret = relationRemoveTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var relationSetStruct *gi.Struct
var relationSetStruct_Once sync.Once

func relationSetStruct_Set() error {
	var err error
	relationSetStruct_Once.Do(func() {
		relationSetStruct, err = gi.StructNew("Atk", "RelationSet")
	})
	return err
}

type RelationSet struct {
	gobject.Object
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
		err = relationSetStruct_Set()
		if err != nil {
			return
		}
		relationSetNewFunction, err = relationSetStruct.InvokerNew("new")
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

	retGo := &RelationSet{}
	retGo.Native = ret.Pointer()

	return retGo
}

var relationSetAddFunction *gi.Function
var relationSetAddFunction_Once sync.Once

func relationSetAddFunction_Set() error {
	var err error
	relationSetAddFunction_Once.Do(func() {
		err = relationSetStruct_Set()
		if err != nil {
			return
		}
		relationSetAddFunction, err = relationSetStruct.InvokerNew("add")
	})
	return err
}

// Add is a representation of the C type atk_relation_set_add.
func (recv *RelationSet) Add(relation *Relation) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(relation.Native)

	err := relationSetAddFunction_Set()
	if err == nil {
		relationSetAddFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'atk_relation_set_add_relation_by_type' : parameter 'relationship' of type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_relation_set_contains' : parameter 'relationship' of type 'RelationType' not supported

// UNSUPPORTED : C value 'atk_relation_set_contains_target' : parameter 'relationship' of type 'RelationType' not supported

var relationSetGetNRelationsFunction *gi.Function
var relationSetGetNRelationsFunction_Once sync.Once

func relationSetGetNRelationsFunction_Set() error {
	var err error
	relationSetGetNRelationsFunction_Once.Do(func() {
		err = relationSetStruct_Set()
		if err != nil {
			return
		}
		relationSetGetNRelationsFunction, err = relationSetStruct.InvokerNew("get_n_relations")
	})
	return err
}

// GetNRelations is a representation of the C type atk_relation_set_get_n_relations.
func (recv *RelationSet) GetNRelations() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = relationSetStruct_Set()
		if err != nil {
			return
		}
		relationSetGetRelationFunction, err = relationSetStruct.InvokerNew("get_relation")
	})
	return err
}

// GetRelation is a representation of the C type atk_relation_set_get_relation.
func (recv *RelationSet) GetRelation(i int32) *Relation {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := relationSetGetRelationFunction_Set()
	if err == nil {
		ret = relationSetGetRelationFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Relation{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'atk_relation_set_get_relation_by_type' : parameter 'relationship' of type 'RelationType' not supported

var relationSetRemoveFunction *gi.Function
var relationSetRemoveFunction_Once sync.Once

func relationSetRemoveFunction_Set() error {
	var err error
	relationSetRemoveFunction_Once.Do(func() {
		err = relationSetStruct_Set()
		if err != nil {
			return
		}
		relationSetRemoveFunction, err = relationSetStruct.InvokerNew("remove")
	})
	return err
}

// Remove is a representation of the C type atk_relation_set_remove.
func (recv *RelationSet) Remove(relation *Relation) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(relation.Native)

	err := relationSetRemoveFunction_Set()
	if err == nil {
		relationSetRemoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var socketStruct *gi.Struct
var socketStruct_Once sync.Once

func socketStruct_Set() error {
	var err error
	socketStruct_Once.Do(func() {
		socketStruct, err = gi.StructNew("Atk", "Socket")
	})
	return err
}

type Socket struct {
	Object
}

// FieldParent returns the C field 'parent'.
func (recv *Socket) FieldParent() *Object {
	argValue := gi.FieldGet(socketStruct, recv.Native, "parent")
	value := &Object{}
	value.Native = argValue.Pointer()
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *Socket) SetFieldParent(value *Object) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native)
	gi.FieldSet(socketStruct, recv.Native, "parent", argValue)
}

var socketNewFunction *gi.Function
var socketNewFunction_Once sync.Once

func socketNewFunction_Set() error {
	var err error
	socketNewFunction_Once.Do(func() {
		err = socketStruct_Set()
		if err != nil {
			return
		}
		socketNewFunction, err = socketStruct.InvokerNew("new")
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

	retGo := &Socket{}
	retGo.Native = ret.Pointer()

	return retGo
}

var socketEmbedFunction *gi.Function
var socketEmbedFunction_Once sync.Once

func socketEmbedFunction_Set() error {
	var err error
	socketEmbedFunction_Once.Do(func() {
		err = socketStruct_Set()
		if err != nil {
			return
		}
		socketEmbedFunction, err = socketStruct.InvokerNew("embed")
	})
	return err
}

// Embed is a representation of the C type atk_socket_embed.
func (recv *Socket) Embed(plugId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
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
		err = socketStruct_Set()
		if err != nil {
			return
		}
		socketIsOccupiedFunction, err = socketStruct.InvokerNew("is_occupied")
	})
	return err
}

// IsOccupied is a representation of the C type atk_socket_is_occupied.
func (recv *Socket) IsOccupied() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := socketIsOccupiedFunction_Set()
	if err == nil {
		ret = socketIsOccupiedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var stateSetStruct *gi.Struct
var stateSetStruct_Once sync.Once

func stateSetStruct_Set() error {
	var err error
	stateSetStruct_Once.Do(func() {
		stateSetStruct, err = gi.StructNew("Atk", "StateSet")
	})
	return err
}

type StateSet struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

var stateSetNewFunction *gi.Function
var stateSetNewFunction_Once sync.Once

func stateSetNewFunction_Set() error {
	var err error
	stateSetNewFunction_Once.Do(func() {
		err = stateSetStruct_Set()
		if err != nil {
			return
		}
		stateSetNewFunction, err = stateSetStruct.InvokerNew("new")
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

	retGo := &StateSet{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'atk_state_set_add_state' : parameter 'type' of type 'StateType' not supported

// UNSUPPORTED : C value 'atk_state_set_add_states' : parameter 'types' of type 'nil' not supported

var stateSetAndSetsFunction *gi.Function
var stateSetAndSetsFunction_Once sync.Once

func stateSetAndSetsFunction_Set() error {
	var err error
	stateSetAndSetsFunction_Once.Do(func() {
		err = stateSetStruct_Set()
		if err != nil {
			return
		}
		stateSetAndSetsFunction, err = stateSetStruct.InvokerNew("and_sets")
	})
	return err
}

// AndSets is a representation of the C type atk_state_set_and_sets.
func (recv *StateSet) AndSets(compareSet *StateSet) *StateSet {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(compareSet.Native)

	var ret gi.Argument

	err := stateSetAndSetsFunction_Set()
	if err == nil {
		ret = stateSetAndSetsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &StateSet{}
	retGo.Native = ret.Pointer()

	return retGo
}

var stateSetClearStatesFunction *gi.Function
var stateSetClearStatesFunction_Once sync.Once

func stateSetClearStatesFunction_Set() error {
	var err error
	stateSetClearStatesFunction_Once.Do(func() {
		err = stateSetStruct_Set()
		if err != nil {
			return
		}
		stateSetClearStatesFunction, err = stateSetStruct.InvokerNew("clear_states")
	})
	return err
}

// ClearStates is a representation of the C type atk_state_set_clear_states.
func (recv *StateSet) ClearStates() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := stateSetClearStatesFunction_Set()
	if err == nil {
		stateSetClearStatesFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'atk_state_set_contains_state' : parameter 'type' of type 'StateType' not supported

// UNSUPPORTED : C value 'atk_state_set_contains_states' : parameter 'types' of type 'nil' not supported

var stateSetIsEmptyFunction *gi.Function
var stateSetIsEmptyFunction_Once sync.Once

func stateSetIsEmptyFunction_Set() error {
	var err error
	stateSetIsEmptyFunction_Once.Do(func() {
		err = stateSetStruct_Set()
		if err != nil {
			return
		}
		stateSetIsEmptyFunction, err = stateSetStruct.InvokerNew("is_empty")
	})
	return err
}

// IsEmpty is a representation of the C type atk_state_set_is_empty.
func (recv *StateSet) IsEmpty() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

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
		err = stateSetStruct_Set()
		if err != nil {
			return
		}
		stateSetOrSetsFunction, err = stateSetStruct.InvokerNew("or_sets")
	})
	return err
}

// OrSets is a representation of the C type atk_state_set_or_sets.
func (recv *StateSet) OrSets(compareSet *StateSet) *StateSet {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(compareSet.Native)

	var ret gi.Argument

	err := stateSetOrSetsFunction_Set()
	if err == nil {
		ret = stateSetOrSetsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &StateSet{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'atk_state_set_remove_state' : parameter 'type' of type 'StateType' not supported

var stateSetXorSetsFunction *gi.Function
var stateSetXorSetsFunction_Once sync.Once

func stateSetXorSetsFunction_Set() error {
	var err error
	stateSetXorSetsFunction_Once.Do(func() {
		err = stateSetStruct_Set()
		if err != nil {
			return
		}
		stateSetXorSetsFunction, err = stateSetStruct.InvokerNew("xor_sets")
	})
	return err
}

// XorSets is a representation of the C type atk_state_set_xor_sets.
func (recv *StateSet) XorSets(compareSet *StateSet) *StateSet {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(compareSet.Native)

	var ret gi.Argument

	err := stateSetXorSetsFunction_Set()
	if err == nil {
		ret = stateSetXorSetsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &StateSet{}
	retGo.Native = ret.Pointer()

	return retGo
}

var utilStruct *gi.Struct
var utilStruct_Once sync.Once

func utilStruct_Set() error {
	var err error
	utilStruct_Once.Do(func() {
		utilStruct, err = gi.StructNew("Atk", "Util")
	})
	return err
}

type Util struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.Object'

// UtilStruct creates an uninitialised Util.
func UtilStruct() *Util {
	err := utilStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Util{}
	structGo.Native = utilStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeUtil)
	return structGo
}
func finalizeUtil(obj *Util) {
	utilStruct.Free(obj.Native)
}
