// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

var actionIfaceStruct *gi.Struct
var actionIfaceStruct_Once sync.Once

func actionIfaceStruct_Set() error {
	var err error
	actionIfaceStruct_Once.Do(func() {
		actionIfaceStruct, err = gi.StructNew("Atk", "ActionIface")
	})
	return err
}

type ActionIface struct {
	native unsafe.Pointer
}

func ActionIfaceNewFromNative(native unsafe.Pointer) *ActionIface {
	instance := &ActionIface{native: native}

	return instance
}

/*
CastToActionIface down casts any arbitrary Object to ActionIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ActionIface.
*/
func CastToActionIface(object *gobject.Object) *ActionIface {
	return ActionIfaceNewFromNative(object.Native())
}

// Equals compares this ActionIface with another ActionIface, and returns true if they represent the same Object.
func (recv *ActionIface) Equals(other *ActionIface) bool {
	return other.Native() == recv.Native()
}

func (recv *ActionIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'do_action' : for field getter : missing Type

// UNSUPPORTED : C value 'do_action' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_actions' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_actions' : for field setter : missing Type

// UNSUPPORTED : C value 'get_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_description' : for field setter : missing Type

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_keybinding' : for field getter : missing Type

// UNSUPPORTED : C value 'get_keybinding' : for field setter : missing Type

// UNSUPPORTED : C value 'set_description' : for field getter : missing Type

// UNSUPPORTED : C value 'set_description' : for field setter : missing Type

// UNSUPPORTED : C value 'get_localized_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_localized_name' : for field setter : missing Type

// ActionIfaceStruct creates an uninitialised ActionIface.
func ActionIfaceStruct() *ActionIface {
	err := actionIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ActionIfaceNewFromNative(actionIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeActionIface)
	return structGo
}
func finalizeActionIface(obj *ActionIface) {
	actionIfaceStruct.Free(obj.Native())
}

var attributeStruct *gi.Struct
var attributeStruct_Once sync.Once

func attributeStruct_Set() error {
	var err error
	attributeStruct_Once.Do(func() {
		attributeStruct, err = gi.StructNew("Atk", "Attribute")
	})
	return err
}

type Attribute struct {
	native unsafe.Pointer
}

func AttributeNewFromNative(native unsafe.Pointer) *Attribute {
	instance := &Attribute{native: native}

	return instance
}

/*
CastToAttribute down casts any arbitrary Object to Attribute.
Exercise care, as this is a potentially dangerous function
if the Object is not a Attribute.
*/
func CastToAttribute(object *gobject.Object) *Attribute {
	return AttributeNewFromNative(object.Native())
}

// Equals compares this Attribute with another Attribute, and returns true if they represent the same Object.
func (recv *Attribute) Equals(other *Attribute) bool {
	return other.Native() == recv.Native()
}

func (recv *Attribute) Native() unsafe.Pointer {
	return recv.native
}

// FieldName returns the C field 'name'.
func (recv *Attribute) FieldName() string {
	argValue := gi.StructFieldGet(attributeStruct, recv.Native(), "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *Attribute) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(attributeStruct, recv.Native(), "name", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *Attribute) FieldValue() string {
	argValue := gi.StructFieldGet(attributeStruct, recv.Native(), "value")
	value := argValue.String(false)
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *Attribute) SetFieldValue(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(attributeStruct, recv.Native(), "value", argValue)
}

// AttributeStruct creates an uninitialised Attribute.
func AttributeStruct() *Attribute {
	err := attributeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttributeNewFromNative(attributeStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttribute)
	return structGo
}
func finalizeAttribute(obj *Attribute) {
	attributeStruct.Free(obj.Native())
}

var componentIfaceStruct *gi.Struct
var componentIfaceStruct_Once sync.Once

func componentIfaceStruct_Set() error {
	var err error
	componentIfaceStruct_Once.Do(func() {
		componentIfaceStruct, err = gi.StructNew("Atk", "ComponentIface")
	})
	return err
}

type ComponentIface struct {
	native unsafe.Pointer
}

func ComponentIfaceNewFromNative(native unsafe.Pointer) *ComponentIface {
	instance := &ComponentIface{native: native}

	return instance
}

/*
CastToComponentIface down casts any arbitrary Object to ComponentIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ComponentIface.
*/
func CastToComponentIface(object *gobject.Object) *ComponentIface {
	return ComponentIfaceNewFromNative(object.Native())
}

// Equals compares this ComponentIface with another ComponentIface, and returns true if they represent the same Object.
func (recv *ComponentIface) Equals(other *ComponentIface) bool {
	return other.Native() == recv.Native()
}

func (recv *ComponentIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'add_focus_handler' : for field getter : missing Type

// UNSUPPORTED : C value 'add_focus_handler' : for field setter : missing Type

// UNSUPPORTED : C value 'contains' : for field getter : missing Type

// UNSUPPORTED : C value 'contains' : for field setter : missing Type

// UNSUPPORTED : C value 'ref_accessible_at_point' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_accessible_at_point' : for field setter : missing Type

// UNSUPPORTED : C value 'get_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'get_extents' : for field setter : missing Type

// UNSUPPORTED : C value 'get_position' : for field getter : missing Type

// UNSUPPORTED : C value 'get_position' : for field setter : missing Type

// UNSUPPORTED : C value 'get_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_size' : for field setter : missing Type

// UNSUPPORTED : C value 'grab_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'grab_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_focus_handler' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_focus_handler' : for field setter : missing Type

// UNSUPPORTED : C value 'set_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'set_extents' : for field setter : missing Type

// UNSUPPORTED : C value 'set_position' : for field getter : missing Type

// UNSUPPORTED : C value 'set_position' : for field setter : missing Type

// UNSUPPORTED : C value 'set_size' : for field getter : missing Type

// UNSUPPORTED : C value 'set_size' : for field setter : missing Type

// UNSUPPORTED : C value 'get_layer' : for field getter : missing Type

// UNSUPPORTED : C value 'get_layer' : for field setter : missing Type

// UNSUPPORTED : C value 'get_mdi_zorder' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mdi_zorder' : for field setter : missing Type

// UNSUPPORTED : C value 'bounds_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'bounds_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'get_alpha' : for field getter : missing Type

// UNSUPPORTED : C value 'get_alpha' : for field setter : missing Type

// UNSUPPORTED : C value 'scroll_to' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_to' : for field setter : missing Type

// UNSUPPORTED : C value 'scroll_to_point' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_to_point' : for field setter : missing Type

// ComponentIfaceStruct creates an uninitialised ComponentIface.
func ComponentIfaceStruct() *ComponentIface {
	err := componentIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ComponentIfaceNewFromNative(componentIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeComponentIface)
	return structGo
}
func finalizeComponentIface(obj *ComponentIface) {
	componentIfaceStruct.Free(obj.Native())
}

var documentIfaceStruct *gi.Struct
var documentIfaceStruct_Once sync.Once

func documentIfaceStruct_Set() error {
	var err error
	documentIfaceStruct_Once.Do(func() {
		documentIfaceStruct, err = gi.StructNew("Atk", "DocumentIface")
	})
	return err
}

type DocumentIface struct {
	native unsafe.Pointer
}

func DocumentIfaceNewFromNative(native unsafe.Pointer) *DocumentIface {
	instance := &DocumentIface{native: native}

	return instance
}

/*
CastToDocumentIface down casts any arbitrary Object to DocumentIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a DocumentIface.
*/
func CastToDocumentIface(object *gobject.Object) *DocumentIface {
	return DocumentIfaceNewFromNative(object.Native())
}

// Equals compares this DocumentIface with another DocumentIface, and returns true if they represent the same Object.
func (recv *DocumentIface) Equals(other *DocumentIface) bool {
	return other.Native() == recv.Native()
}

func (recv *DocumentIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *DocumentIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(documentIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *DocumentIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(documentIfaceStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'get_document_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_document_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_document' : for field getter : missing Type

// UNSUPPORTED : C value 'get_document' : for field setter : missing Type

// UNSUPPORTED : C value 'get_document_locale' : for field getter : missing Type

// UNSUPPORTED : C value 'get_document_locale' : for field setter : missing Type

// UNSUPPORTED : C value 'get_document_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_document_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'get_document_attribute_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_document_attribute_value' : for field setter : missing Type

// UNSUPPORTED : C value 'set_document_attribute' : for field getter : missing Type

// UNSUPPORTED : C value 'set_document_attribute' : for field setter : missing Type

// UNSUPPORTED : C value 'get_current_page_number' : for field getter : missing Type

// UNSUPPORTED : C value 'get_current_page_number' : for field setter : missing Type

// UNSUPPORTED : C value 'get_page_count' : for field getter : missing Type

// UNSUPPORTED : C value 'get_page_count' : for field setter : missing Type

// DocumentIfaceStruct creates an uninitialised DocumentIface.
func DocumentIfaceStruct() *DocumentIface {
	err := documentIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := DocumentIfaceNewFromNative(documentIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeDocumentIface)
	return structGo
}
func finalizeDocumentIface(obj *DocumentIface) {
	documentIfaceStruct.Free(obj.Native())
}

var editableTextIfaceStruct *gi.Struct
var editableTextIfaceStruct_Once sync.Once

func editableTextIfaceStruct_Set() error {
	var err error
	editableTextIfaceStruct_Once.Do(func() {
		editableTextIfaceStruct, err = gi.StructNew("Atk", "EditableTextIface")
	})
	return err
}

type EditableTextIface struct {
	native unsafe.Pointer
}

func EditableTextIfaceNewFromNative(native unsafe.Pointer) *EditableTextIface {
	instance := &EditableTextIface{native: native}

	return instance
}

/*
CastToEditableTextIface down casts any arbitrary Object to EditableTextIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a EditableTextIface.
*/
func CastToEditableTextIface(object *gobject.Object) *EditableTextIface {
	return EditableTextIfaceNewFromNative(object.Native())
}

// Equals compares this EditableTextIface with another EditableTextIface, and returns true if they represent the same Object.
func (recv *EditableTextIface) Equals(other *EditableTextIface) bool {
	return other.Native() == recv.Native()
}

func (recv *EditableTextIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentInterface returns the C field 'parent_interface'.
func (recv *EditableTextIface) FieldParentInterface() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(editableTextIfaceStruct, recv.Native(), "parent_interface")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentInterface sets the value of the C field 'parent_interface'.
func (recv *EditableTextIface) SetFieldParentInterface(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(editableTextIfaceStruct, recv.Native(), "parent_interface", argValue)
}

// UNSUPPORTED : C value 'set_run_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'set_run_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'set_text_contents' : for field getter : missing Type

// UNSUPPORTED : C value 'set_text_contents' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_text' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_text' : for field setter : missing Type

// UNSUPPORTED : C value 'copy_text' : for field getter : missing Type

// UNSUPPORTED : C value 'copy_text' : for field setter : missing Type

// UNSUPPORTED : C value 'cut_text' : for field getter : missing Type

// UNSUPPORTED : C value 'cut_text' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_text' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_text' : for field setter : missing Type

// UNSUPPORTED : C value 'paste_text' : for field getter : missing Type

// UNSUPPORTED : C value 'paste_text' : for field setter : missing Type

// EditableTextIfaceStruct creates an uninitialised EditableTextIface.
func EditableTextIfaceStruct() *EditableTextIface {
	err := editableTextIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EditableTextIfaceNewFromNative(editableTextIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEditableTextIface)
	return structGo
}
func finalizeEditableTextIface(obj *EditableTextIface) {
	editableTextIfaceStruct.Free(obj.Native())
}

var gObjectAccessibleClassStruct *gi.Struct
var gObjectAccessibleClassStruct_Once sync.Once

func gObjectAccessibleClassStruct_Set() error {
	var err error
	gObjectAccessibleClassStruct_Once.Do(func() {
		gObjectAccessibleClassStruct, err = gi.StructNew("Atk", "GObjectAccessibleClass")
	})
	return err
}

type GObjectAccessibleClass struct {
	native unsafe.Pointer
}

func GObjectAccessibleClassNewFromNative(native unsafe.Pointer) *GObjectAccessibleClass {
	instance := &GObjectAccessibleClass{native: native}

	return instance
}

/*
CastToGObjectAccessibleClass down casts any arbitrary Object to GObjectAccessibleClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a GObjectAccessibleClass.
*/
func CastToGObjectAccessibleClass(object *gobject.Object) *GObjectAccessibleClass {
	return GObjectAccessibleClassNewFromNative(object.Native())
}

// Equals compares this GObjectAccessibleClass with another GObjectAccessibleClass, and returns true if they represent the same Object.
func (recv *GObjectAccessibleClass) Equals(other *GObjectAccessibleClass) bool {
	return other.Native() == recv.Native()
}

func (recv *GObjectAccessibleClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *GObjectAccessibleClass) FieldParentClass() *ObjectClass {
	argValue := gi.StructFieldGet(gObjectAccessibleClassStruct, recv.Native(), "parent_class")
	value := ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *GObjectAccessibleClass) SetFieldParentClass(value *ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(gObjectAccessibleClassStruct, recv.Native(), "parent_class", argValue)
}

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad1' : for field setter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field setter : no Go type for 'Function'

// GObjectAccessibleClassStruct creates an uninitialised GObjectAccessibleClass.
func GObjectAccessibleClassStruct() *GObjectAccessibleClass {
	err := gObjectAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GObjectAccessibleClassNewFromNative(gObjectAccessibleClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGObjectAccessibleClass)
	return structGo
}
func finalizeGObjectAccessibleClass(obj *GObjectAccessibleClass) {
	gObjectAccessibleClassStruct.Free(obj.Native())
}

var hyperlinkClassStruct *gi.Struct
var hyperlinkClassStruct_Once sync.Once

func hyperlinkClassStruct_Set() error {
	var err error
	hyperlinkClassStruct_Once.Do(func() {
		hyperlinkClassStruct, err = gi.StructNew("Atk", "HyperlinkClass")
	})
	return err
}

type HyperlinkClass struct {
	native unsafe.Pointer
}

func HyperlinkClassNewFromNative(native unsafe.Pointer) *HyperlinkClass {
	instance := &HyperlinkClass{native: native}

	return instance
}

/*
CastToHyperlinkClass down casts any arbitrary Object to HyperlinkClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a HyperlinkClass.
*/
func CastToHyperlinkClass(object *gobject.Object) *HyperlinkClass {
	return HyperlinkClassNewFromNative(object.Native())
}

// Equals compares this HyperlinkClass with another HyperlinkClass, and returns true if they represent the same Object.
func (recv *HyperlinkClass) Equals(other *HyperlinkClass) bool {
	return other.Native() == recv.Native()
}

func (recv *HyperlinkClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *HyperlinkClass) FieldParent() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(hyperlinkClassStruct, recv.Native(), "parent")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *HyperlinkClass) SetFieldParent(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(hyperlinkClassStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'get_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'get_uri' : for field setter : missing Type

// UNSUPPORTED : C value 'get_object' : for field getter : missing Type

// UNSUPPORTED : C value 'get_object' : for field setter : missing Type

// UNSUPPORTED : C value 'get_end_index' : for field getter : missing Type

// UNSUPPORTED : C value 'get_end_index' : for field setter : missing Type

// UNSUPPORTED : C value 'get_start_index' : for field getter : missing Type

// UNSUPPORTED : C value 'get_start_index' : for field setter : missing Type

// UNSUPPORTED : C value 'is_valid' : for field getter : missing Type

// UNSUPPORTED : C value 'is_valid' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_anchors' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_anchors' : for field setter : missing Type

// UNSUPPORTED : C value 'link_state' : for field getter : missing Type

// UNSUPPORTED : C value 'link_state' : for field setter : missing Type

// UNSUPPORTED : C value 'is_selected_link' : for field getter : missing Type

// UNSUPPORTED : C value 'is_selected_link' : for field setter : missing Type

// UNSUPPORTED : C value 'link_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'link_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad1' : for field setter : no Go type for 'Function'

// HyperlinkClassStruct creates an uninitialised HyperlinkClass.
func HyperlinkClassStruct() *HyperlinkClass {
	err := hyperlinkClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := HyperlinkClassNewFromNative(hyperlinkClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeHyperlinkClass)
	return structGo
}
func finalizeHyperlinkClass(obj *HyperlinkClass) {
	hyperlinkClassStruct.Free(obj.Native())
}

var hyperlinkImplIfaceStruct *gi.Struct
var hyperlinkImplIfaceStruct_Once sync.Once

func hyperlinkImplIfaceStruct_Set() error {
	var err error
	hyperlinkImplIfaceStruct_Once.Do(func() {
		hyperlinkImplIfaceStruct, err = gi.StructNew("Atk", "HyperlinkImplIface")
	})
	return err
}

type HyperlinkImplIface struct {
	native unsafe.Pointer
}

func HyperlinkImplIfaceNewFromNative(native unsafe.Pointer) *HyperlinkImplIface {
	instance := &HyperlinkImplIface{native: native}

	return instance
}

/*
CastToHyperlinkImplIface down casts any arbitrary Object to HyperlinkImplIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a HyperlinkImplIface.
*/
func CastToHyperlinkImplIface(object *gobject.Object) *HyperlinkImplIface {
	return HyperlinkImplIfaceNewFromNative(object.Native())
}

// Equals compares this HyperlinkImplIface with another HyperlinkImplIface, and returns true if they represent the same Object.
func (recv *HyperlinkImplIface) Equals(other *HyperlinkImplIface) bool {
	return other.Native() == recv.Native()
}

func (recv *HyperlinkImplIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *HyperlinkImplIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(hyperlinkImplIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *HyperlinkImplIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(hyperlinkImplIfaceStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'get_hyperlink' : for field getter : missing Type

// UNSUPPORTED : C value 'get_hyperlink' : for field setter : missing Type

// HyperlinkImplIfaceStruct creates an uninitialised HyperlinkImplIface.
func HyperlinkImplIfaceStruct() *HyperlinkImplIface {
	err := hyperlinkImplIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := HyperlinkImplIfaceNewFromNative(hyperlinkImplIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeHyperlinkImplIface)
	return structGo
}
func finalizeHyperlinkImplIface(obj *HyperlinkImplIface) {
	hyperlinkImplIfaceStruct.Free(obj.Native())
}

var hypertextIfaceStruct *gi.Struct
var hypertextIfaceStruct_Once sync.Once

func hypertextIfaceStruct_Set() error {
	var err error
	hypertextIfaceStruct_Once.Do(func() {
		hypertextIfaceStruct, err = gi.StructNew("Atk", "HypertextIface")
	})
	return err
}

type HypertextIface struct {
	native unsafe.Pointer
}

func HypertextIfaceNewFromNative(native unsafe.Pointer) *HypertextIface {
	instance := &HypertextIface{native: native}

	return instance
}

/*
CastToHypertextIface down casts any arbitrary Object to HypertextIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a HypertextIface.
*/
func CastToHypertextIface(object *gobject.Object) *HypertextIface {
	return HypertextIfaceNewFromNative(object.Native())
}

// Equals compares this HypertextIface with another HypertextIface, and returns true if they represent the same Object.
func (recv *HypertextIface) Equals(other *HypertextIface) bool {
	return other.Native() == recv.Native()
}

func (recv *HypertextIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *HypertextIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(hypertextIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *HypertextIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(hypertextIfaceStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'get_link' : for field getter : missing Type

// UNSUPPORTED : C value 'get_link' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_links' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_links' : for field setter : missing Type

// UNSUPPORTED : C value 'get_link_index' : for field getter : missing Type

// UNSUPPORTED : C value 'get_link_index' : for field setter : missing Type

// UNSUPPORTED : C value 'link_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'link_selected' : for field setter : missing Type

// HypertextIfaceStruct creates an uninitialised HypertextIface.
func HypertextIfaceStruct() *HypertextIface {
	err := hypertextIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := HypertextIfaceNewFromNative(hypertextIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeHypertextIface)
	return structGo
}
func finalizeHypertextIface(obj *HypertextIface) {
	hypertextIfaceStruct.Free(obj.Native())
}

var imageIfaceStruct *gi.Struct
var imageIfaceStruct_Once sync.Once

func imageIfaceStruct_Set() error {
	var err error
	imageIfaceStruct_Once.Do(func() {
		imageIfaceStruct, err = gi.StructNew("Atk", "ImageIface")
	})
	return err
}

type ImageIface struct {
	native unsafe.Pointer
}

func ImageIfaceNewFromNative(native unsafe.Pointer) *ImageIface {
	instance := &ImageIface{native: native}

	return instance
}

/*
CastToImageIface down casts any arbitrary Object to ImageIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ImageIface.
*/
func CastToImageIface(object *gobject.Object) *ImageIface {
	return ImageIfaceNewFromNative(object.Native())
}

// Equals compares this ImageIface with another ImageIface, and returns true if they represent the same Object.
func (recv *ImageIface) Equals(other *ImageIface) bool {
	return other.Native() == recv.Native()
}

func (recv *ImageIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *ImageIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(imageIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *ImageIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(imageIfaceStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'get_image_position' : for field getter : missing Type

// UNSUPPORTED : C value 'get_image_position' : for field setter : missing Type

// UNSUPPORTED : C value 'get_image_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_image_description' : for field setter : missing Type

// UNSUPPORTED : C value 'get_image_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_image_size' : for field setter : missing Type

// UNSUPPORTED : C value 'set_image_description' : for field getter : missing Type

// UNSUPPORTED : C value 'set_image_description' : for field setter : missing Type

// UNSUPPORTED : C value 'get_image_locale' : for field getter : missing Type

// UNSUPPORTED : C value 'get_image_locale' : for field setter : missing Type

// ImageIfaceStruct creates an uninitialised ImageIface.
func ImageIfaceStruct() *ImageIface {
	err := imageIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ImageIfaceNewFromNative(imageIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeImageIface)
	return structGo
}
func finalizeImageIface(obj *ImageIface) {
	imageIfaceStruct.Free(obj.Native())
}

var implementorStruct *gi.Struct
var implementorStruct_Once sync.Once

func implementorStruct_Set() error {
	var err error
	implementorStruct_Once.Do(func() {
		implementorStruct, err = gi.StructNew("Atk", "Implementor")
	})
	return err
}

type Implementor struct {
	native unsafe.Pointer
}

func ImplementorNewFromNative(native unsafe.Pointer) *Implementor {
	instance := &Implementor{native: native}

	return instance
}

/*
CastToImplementor down casts any arbitrary Object to Implementor.
Exercise care, as this is a potentially dangerous function
if the Object is not a Implementor.
*/
func CastToImplementor(object *gobject.Object) *Implementor {
	return ImplementorNewFromNative(object.Native())
}

// Equals compares this Implementor with another Implementor, and returns true if they represent the same Object.
func (recv *Implementor) Equals(other *Implementor) bool {
	return other.Native() == recv.Native()
}

func (recv *Implementor) Native() unsafe.Pointer {
	return recv.native
}

var implementorRefAccessibleFunction *gi.Function
var implementorRefAccessibleFunction_Once sync.Once

func implementorRefAccessibleFunction_Set() error {
	var err error
	implementorRefAccessibleFunction_Once.Do(func() {
		err = implementorStruct_Set()
		if err != nil {
			return
		}
		implementorRefAccessibleFunction, err = implementorStruct.InvokerNew("ref_accessible")
	})
	return err
}

// RefAccessible is a representation of the C type atk_implementor_ref_accessible.
func (recv *Implementor) RefAccessible() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := implementorRefAccessibleFunction_Set()
	if err == nil {
		ret = implementorRefAccessibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

// ImplementorStruct creates an uninitialised Implementor.
func ImplementorStruct() *Implementor {
	err := implementorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ImplementorNewFromNative(implementorStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeImplementor)
	return structGo
}
func finalizeImplementor(obj *Implementor) {
	implementorStruct.Free(obj.Native())
}

var keyEventStructStruct *gi.Struct
var keyEventStructStruct_Once sync.Once

func keyEventStructStruct_Set() error {
	var err error
	keyEventStructStruct_Once.Do(func() {
		keyEventStructStruct, err = gi.StructNew("Atk", "KeyEventStruct")
	})
	return err
}

type KeyEventStruct struct {
	native unsafe.Pointer
}

func KeyEventStructNewFromNative(native unsafe.Pointer) *KeyEventStruct {
	instance := &KeyEventStruct{native: native}

	return instance
}

/*
CastToKeyEventStruct down casts any arbitrary Object to KeyEventStruct.
Exercise care, as this is a potentially dangerous function
if the Object is not a KeyEventStruct.
*/
func CastToKeyEventStruct(object *gobject.Object) *KeyEventStruct {
	return KeyEventStructNewFromNative(object.Native())
}

// Equals compares this KeyEventStruct with another KeyEventStruct, and returns true if they represent the same Object.
func (recv *KeyEventStruct) Equals(other *KeyEventStruct) bool {
	return other.Native() == recv.Native()
}

func (recv *KeyEventStruct) Native() unsafe.Pointer {
	return recv.native
}

// FieldType returns the C field 'type'.
func (recv *KeyEventStruct) FieldType() int32 {
	argValue := gi.StructFieldGet(keyEventStructStruct, recv.Native(), "type")
	value := argValue.Int32()
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *KeyEventStruct) SetFieldType(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(keyEventStructStruct, recv.Native(), "type", argValue)
}

// FieldState returns the C field 'state'.
func (recv *KeyEventStruct) FieldState() uint32 {
	argValue := gi.StructFieldGet(keyEventStructStruct, recv.Native(), "state")
	value := argValue.Uint32()
	return value
}

// SetFieldState sets the value of the C field 'state'.
func (recv *KeyEventStruct) SetFieldState(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(keyEventStructStruct, recv.Native(), "state", argValue)
}

// FieldKeyval returns the C field 'keyval'.
func (recv *KeyEventStruct) FieldKeyval() uint32 {
	argValue := gi.StructFieldGet(keyEventStructStruct, recv.Native(), "keyval")
	value := argValue.Uint32()
	return value
}

// SetFieldKeyval sets the value of the C field 'keyval'.
func (recv *KeyEventStruct) SetFieldKeyval(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(keyEventStructStruct, recv.Native(), "keyval", argValue)
}

// FieldLength returns the C field 'length'.
func (recv *KeyEventStruct) FieldLength() int32 {
	argValue := gi.StructFieldGet(keyEventStructStruct, recv.Native(), "length")
	value := argValue.Int32()
	return value
}

// SetFieldLength sets the value of the C field 'length'.
func (recv *KeyEventStruct) SetFieldLength(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(keyEventStructStruct, recv.Native(), "length", argValue)
}

// FieldString returns the C field 'string'.
func (recv *KeyEventStruct) FieldString() string {
	argValue := gi.StructFieldGet(keyEventStructStruct, recv.Native(), "string")
	value := argValue.String(false)
	return value
}

// SetFieldString sets the value of the C field 'string'.
func (recv *KeyEventStruct) SetFieldString(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(keyEventStructStruct, recv.Native(), "string", argValue)
}

// FieldKeycode returns the C field 'keycode'.
func (recv *KeyEventStruct) FieldKeycode() uint16 {
	argValue := gi.StructFieldGet(keyEventStructStruct, recv.Native(), "keycode")
	value := argValue.Uint16()
	return value
}

// SetFieldKeycode sets the value of the C field 'keycode'.
func (recv *KeyEventStruct) SetFieldKeycode(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(keyEventStructStruct, recv.Native(), "keycode", argValue)
}

// FieldTimestamp returns the C field 'timestamp'.
func (recv *KeyEventStruct) FieldTimestamp() uint32 {
	argValue := gi.StructFieldGet(keyEventStructStruct, recv.Native(), "timestamp")
	value := argValue.Uint32()
	return value
}

// SetFieldTimestamp sets the value of the C field 'timestamp'.
func (recv *KeyEventStruct) SetFieldTimestamp(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(keyEventStructStruct, recv.Native(), "timestamp", argValue)
}

// KeyEventStructStruct creates an uninitialised KeyEventStruct.
func KeyEventStructStruct() *KeyEventStruct {
	err := keyEventStructStruct_Set()
	if err != nil {
		return nil
	}

	structGo := KeyEventStructNewFromNative(keyEventStructStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeKeyEventStruct)
	return structGo
}
func finalizeKeyEventStruct(obj *KeyEventStruct) {
	keyEventStructStruct.Free(obj.Native())
}

var miscClassStruct *gi.Struct
var miscClassStruct_Once sync.Once

func miscClassStruct_Set() error {
	var err error
	miscClassStruct_Once.Do(func() {
		miscClassStruct, err = gi.StructNew("Atk", "MiscClass")
	})
	return err
}

type MiscClass struct {
	native unsafe.Pointer
}

func MiscClassNewFromNative(native unsafe.Pointer) *MiscClass {
	instance := &MiscClass{native: native}

	return instance
}

/*
CastToMiscClass down casts any arbitrary Object to MiscClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a MiscClass.
*/
func CastToMiscClass(object *gobject.Object) *MiscClass {
	return MiscClassNewFromNative(object.Native())
}

// Equals compares this MiscClass with another MiscClass, and returns true if they represent the same Object.
func (recv *MiscClass) Equals(other *MiscClass) bool {
	return other.Native() == recv.Native()
}

func (recv *MiscClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *MiscClass) FieldParent() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(miscClassStruct, recv.Native(), "parent")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *MiscClass) SetFieldParent(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(miscClassStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'threads_enter' : for field getter : missing Type

// UNSUPPORTED : C value 'threads_enter' : for field setter : missing Type

// UNSUPPORTED : C value 'threads_leave' : for field getter : missing Type

// UNSUPPORTED : C value 'threads_leave' : for field setter : missing Type

// UNSUPPORTED : C value 'vfuncs' : for field getter : missing Type

// UNSUPPORTED : C value 'vfuncs' : for field setter : missing Type

// MiscClassStruct creates an uninitialised MiscClass.
func MiscClassStruct() *MiscClass {
	err := miscClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MiscClassNewFromNative(miscClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMiscClass)
	return structGo
}
func finalizeMiscClass(obj *MiscClass) {
	miscClassStruct.Free(obj.Native())
}

var noOpObjectClassStruct *gi.Struct
var noOpObjectClassStruct_Once sync.Once

func noOpObjectClassStruct_Set() error {
	var err error
	noOpObjectClassStruct_Once.Do(func() {
		noOpObjectClassStruct, err = gi.StructNew("Atk", "NoOpObjectClass")
	})
	return err
}

type NoOpObjectClass struct {
	native unsafe.Pointer
}

func NoOpObjectClassNewFromNative(native unsafe.Pointer) *NoOpObjectClass {
	instance := &NoOpObjectClass{native: native}

	return instance
}

/*
CastToNoOpObjectClass down casts any arbitrary Object to NoOpObjectClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a NoOpObjectClass.
*/
func CastToNoOpObjectClass(object *gobject.Object) *NoOpObjectClass {
	return NoOpObjectClassNewFromNative(object.Native())
}

// Equals compares this NoOpObjectClass with another NoOpObjectClass, and returns true if they represent the same Object.
func (recv *NoOpObjectClass) Equals(other *NoOpObjectClass) bool {
	return other.Native() == recv.Native()
}

func (recv *NoOpObjectClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *NoOpObjectClass) FieldParentClass() *ObjectClass {
	argValue := gi.StructFieldGet(noOpObjectClassStruct, recv.Native(), "parent_class")
	value := ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *NoOpObjectClass) SetFieldParentClass(value *ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(noOpObjectClassStruct, recv.Native(), "parent_class", argValue)
}

// NoOpObjectClassStruct creates an uninitialised NoOpObjectClass.
func NoOpObjectClassStruct() *NoOpObjectClass {
	err := noOpObjectClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NoOpObjectClassNewFromNative(noOpObjectClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNoOpObjectClass)
	return structGo
}
func finalizeNoOpObjectClass(obj *NoOpObjectClass) {
	noOpObjectClassStruct.Free(obj.Native())
}

var noOpObjectFactoryClassStruct *gi.Struct
var noOpObjectFactoryClassStruct_Once sync.Once

func noOpObjectFactoryClassStruct_Set() error {
	var err error
	noOpObjectFactoryClassStruct_Once.Do(func() {
		noOpObjectFactoryClassStruct, err = gi.StructNew("Atk", "NoOpObjectFactoryClass")
	})
	return err
}

type NoOpObjectFactoryClass struct {
	native unsafe.Pointer
}

func NoOpObjectFactoryClassNewFromNative(native unsafe.Pointer) *NoOpObjectFactoryClass {
	instance := &NoOpObjectFactoryClass{native: native}

	return instance
}

/*
CastToNoOpObjectFactoryClass down casts any arbitrary Object to NoOpObjectFactoryClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a NoOpObjectFactoryClass.
*/
func CastToNoOpObjectFactoryClass(object *gobject.Object) *NoOpObjectFactoryClass {
	return NoOpObjectFactoryClassNewFromNative(object.Native())
}

// Equals compares this NoOpObjectFactoryClass with another NoOpObjectFactoryClass, and returns true if they represent the same Object.
func (recv *NoOpObjectFactoryClass) Equals(other *NoOpObjectFactoryClass) bool {
	return other.Native() == recv.Native()
}

func (recv *NoOpObjectFactoryClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *NoOpObjectFactoryClass) FieldParentClass() *ObjectFactoryClass {
	argValue := gi.StructFieldGet(noOpObjectFactoryClassStruct, recv.Native(), "parent_class")
	value := ObjectFactoryClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *NoOpObjectFactoryClass) SetFieldParentClass(value *ObjectFactoryClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(noOpObjectFactoryClassStruct, recv.Native(), "parent_class", argValue)
}

// NoOpObjectFactoryClassStruct creates an uninitialised NoOpObjectFactoryClass.
func NoOpObjectFactoryClassStruct() *NoOpObjectFactoryClass {
	err := noOpObjectFactoryClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := NoOpObjectFactoryClassNewFromNative(noOpObjectFactoryClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeNoOpObjectFactoryClass)
	return structGo
}
func finalizeNoOpObjectFactoryClass(obj *NoOpObjectFactoryClass) {
	noOpObjectFactoryClassStruct.Free(obj.Native())
}

var objectClassStruct *gi.Struct
var objectClassStruct_Once sync.Once

func objectClassStruct_Set() error {
	var err error
	objectClassStruct_Once.Do(func() {
		objectClassStruct, err = gi.StructNew("Atk", "ObjectClass")
	})
	return err
}

type ObjectClass struct {
	native unsafe.Pointer
}

func ObjectClassNewFromNative(native unsafe.Pointer) *ObjectClass {
	instance := &ObjectClass{native: native}

	return instance
}

/*
CastToObjectClass down casts any arbitrary Object to ObjectClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ObjectClass.
*/
func CastToObjectClass(object *gobject.Object) *ObjectClass {
	return ObjectClassNewFromNative(object.Native())
}

// Equals compares this ObjectClass with another ObjectClass, and returns true if they represent the same Object.
func (recv *ObjectClass) Equals(other *ObjectClass) bool {
	return other.Native() == recv.Native()
}

func (recv *ObjectClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *ObjectClass) FieldParent() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(objectClassStruct, recv.Native(), "parent")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *ObjectClass) SetFieldParent(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(objectClassStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_description' : for field setter : missing Type

// UNSUPPORTED : C value 'get_parent' : for field getter : missing Type

// UNSUPPORTED : C value 'get_parent' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_children' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_children' : for field setter : missing Type

// UNSUPPORTED : C value 'ref_child' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_child' : for field setter : missing Type

// UNSUPPORTED : C value 'get_index_in_parent' : for field getter : missing Type

// UNSUPPORTED : C value 'get_index_in_parent' : for field setter : missing Type

// UNSUPPORTED : C value 'ref_relation_set' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_relation_set' : for field setter : missing Type

// UNSUPPORTED : C value 'get_role' : for field getter : missing Type

// UNSUPPORTED : C value 'get_role' : for field setter : missing Type

// UNSUPPORTED : C value 'get_layer' : for field getter : missing Type

// UNSUPPORTED : C value 'get_layer' : for field setter : missing Type

// UNSUPPORTED : C value 'get_mdi_zorder' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mdi_zorder' : for field setter : missing Type

// UNSUPPORTED : C value 'ref_state_set' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_state_set' : for field setter : missing Type

// UNSUPPORTED : C value 'set_name' : for field getter : missing Type

// UNSUPPORTED : C value 'set_name' : for field setter : missing Type

// UNSUPPORTED : C value 'set_description' : for field getter : missing Type

// UNSUPPORTED : C value 'set_description' : for field setter : missing Type

// UNSUPPORTED : C value 'set_parent' : for field getter : missing Type

// UNSUPPORTED : C value 'set_parent' : for field setter : missing Type

// UNSUPPORTED : C value 'set_role' : for field getter : missing Type

// UNSUPPORTED : C value 'set_role' : for field setter : missing Type

// UNSUPPORTED : C value 'connect_property_change_handler' : for field getter : missing Type

// UNSUPPORTED : C value 'connect_property_change_handler' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_property_change_handler' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_property_change_handler' : for field setter : missing Type

// UNSUPPORTED : C value 'initialize' : for field getter : missing Type

// UNSUPPORTED : C value 'initialize' : for field setter : missing Type

// UNSUPPORTED : C value 'children_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'children_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'focus_event' : for field getter : missing Type

// UNSUPPORTED : C value 'focus_event' : for field setter : missing Type

// UNSUPPORTED : C value 'property_change' : for field getter : missing Type

// UNSUPPORTED : C value 'property_change' : for field setter : missing Type

// UNSUPPORTED : C value 'state_change' : for field getter : missing Type

// UNSUPPORTED : C value 'state_change' : for field setter : missing Type

// UNSUPPORTED : C value 'visible_data_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'visible_data_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'active_descendant_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'active_descendant_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'get_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'get_object_locale' : for field getter : missing Type

// UNSUPPORTED : C value 'get_object_locale' : for field setter : missing Type

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad1' : for field setter : no Go type for 'Function'

// ObjectClassStruct creates an uninitialised ObjectClass.
func ObjectClassStruct() *ObjectClass {
	err := objectClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ObjectClassNewFromNative(objectClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeObjectClass)
	return structGo
}
func finalizeObjectClass(obj *ObjectClass) {
	objectClassStruct.Free(obj.Native())
}

var objectFactoryClassStruct *gi.Struct
var objectFactoryClassStruct_Once sync.Once

func objectFactoryClassStruct_Set() error {
	var err error
	objectFactoryClassStruct_Once.Do(func() {
		objectFactoryClassStruct, err = gi.StructNew("Atk", "ObjectFactoryClass")
	})
	return err
}

type ObjectFactoryClass struct {
	native unsafe.Pointer
}

func ObjectFactoryClassNewFromNative(native unsafe.Pointer) *ObjectFactoryClass {
	instance := &ObjectFactoryClass{native: native}

	return instance
}

/*
CastToObjectFactoryClass down casts any arbitrary Object to ObjectFactoryClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a ObjectFactoryClass.
*/
func CastToObjectFactoryClass(object *gobject.Object) *ObjectFactoryClass {
	return ObjectFactoryClassNewFromNative(object.Native())
}

// Equals compares this ObjectFactoryClass with another ObjectFactoryClass, and returns true if they represent the same Object.
func (recv *ObjectFactoryClass) Equals(other *ObjectFactoryClass) bool {
	return other.Native() == recv.Native()
}

func (recv *ObjectFactoryClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *ObjectFactoryClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(objectFactoryClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ObjectFactoryClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(objectFactoryClassStruct, recv.Native(), "parent_class", argValue)
}

// UNSUPPORTED : C value 'create_accessible' : for field getter : missing Type

// UNSUPPORTED : C value 'create_accessible' : for field setter : missing Type

// UNSUPPORTED : C value 'invalidate' : for field getter : missing Type

// UNSUPPORTED : C value 'invalidate' : for field setter : missing Type

// UNSUPPORTED : C value 'get_accessible_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_accessible_type' : for field setter : missing Type

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad1' : for field setter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field setter : no Go type for 'Function'

// ObjectFactoryClassStruct creates an uninitialised ObjectFactoryClass.
func ObjectFactoryClassStruct() *ObjectFactoryClass {
	err := objectFactoryClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ObjectFactoryClassNewFromNative(objectFactoryClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeObjectFactoryClass)
	return structGo
}
func finalizeObjectFactoryClass(obj *ObjectFactoryClass) {
	objectFactoryClassStruct.Free(obj.Native())
}

var plugClassStruct *gi.Struct
var plugClassStruct_Once sync.Once

func plugClassStruct_Set() error {
	var err error
	plugClassStruct_Once.Do(func() {
		plugClassStruct, err = gi.StructNew("Atk", "PlugClass")
	})
	return err
}

type PlugClass struct {
	native unsafe.Pointer
}

func PlugClassNewFromNative(native unsafe.Pointer) *PlugClass {
	instance := &PlugClass{native: native}

	return instance
}

/*
CastToPlugClass down casts any arbitrary Object to PlugClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a PlugClass.
*/
func CastToPlugClass(object *gobject.Object) *PlugClass {
	return PlugClassNewFromNative(object.Native())
}

// Equals compares this PlugClass with another PlugClass, and returns true if they represent the same Object.
func (recv *PlugClass) Equals(other *PlugClass) bool {
	return other.Native() == recv.Native()
}

func (recv *PlugClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *PlugClass) FieldParentClass() *ObjectClass {
	argValue := gi.StructFieldGet(plugClassStruct, recv.Native(), "parent_class")
	value := ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *PlugClass) SetFieldParentClass(value *ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(plugClassStruct, recv.Native(), "parent_class", argValue)
}

// UNSUPPORTED : C value 'get_object_id' : for field getter : missing Type

// UNSUPPORTED : C value 'get_object_id' : for field setter : missing Type

// PlugClassStruct creates an uninitialised PlugClass.
func PlugClassStruct() *PlugClass {
	err := plugClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PlugClassNewFromNative(plugClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePlugClass)
	return structGo
}
func finalizePlugClass(obj *PlugClass) {
	plugClassStruct.Free(obj.Native())
}

var propertyValuesStruct *gi.Struct
var propertyValuesStruct_Once sync.Once

func propertyValuesStruct_Set() error {
	var err error
	propertyValuesStruct_Once.Do(func() {
		propertyValuesStruct, err = gi.StructNew("Atk", "PropertyValues")
	})
	return err
}

type PropertyValues struct {
	native unsafe.Pointer
}

func PropertyValuesNewFromNative(native unsafe.Pointer) *PropertyValues {
	instance := &PropertyValues{native: native}

	return instance
}

/*
CastToPropertyValues down casts any arbitrary Object to PropertyValues.
Exercise care, as this is a potentially dangerous function
if the Object is not a PropertyValues.
*/
func CastToPropertyValues(object *gobject.Object) *PropertyValues {
	return PropertyValuesNewFromNative(object.Native())
}

// Equals compares this PropertyValues with another PropertyValues, and returns true if they represent the same Object.
func (recv *PropertyValues) Equals(other *PropertyValues) bool {
	return other.Native() == recv.Native()
}

func (recv *PropertyValues) Native() unsafe.Pointer {
	return recv.native
}

// FieldPropertyName returns the C field 'property_name'.
func (recv *PropertyValues) FieldPropertyName() string {
	argValue := gi.StructFieldGet(propertyValuesStruct, recv.Native(), "property_name")
	value := argValue.String(false)
	return value
}

// SetFieldPropertyName sets the value of the C field 'property_name'.
func (recv *PropertyValues) SetFieldPropertyName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(propertyValuesStruct, recv.Native(), "property_name", argValue)
}

// FieldOldValue returns the C field 'old_value'.
func (recv *PropertyValues) FieldOldValue() *gobject.Value {
	argValue := gi.StructFieldGet(propertyValuesStruct, recv.Native(), "old_value")
	value := gobject.ValueNewFromNative(argValue.Pointer())
	return value
}

// SetFieldOldValue sets the value of the C field 'old_value'.
func (recv *PropertyValues) SetFieldOldValue(value *gobject.Value) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(propertyValuesStruct, recv.Native(), "old_value", argValue)
}

// FieldNewValue returns the C field 'new_value'.
func (recv *PropertyValues) FieldNewValue() *gobject.Value {
	argValue := gi.StructFieldGet(propertyValuesStruct, recv.Native(), "new_value")
	value := gobject.ValueNewFromNative(argValue.Pointer())
	return value
}

// SetFieldNewValue sets the value of the C field 'new_value'.
func (recv *PropertyValues) SetFieldNewValue(value *gobject.Value) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(propertyValuesStruct, recv.Native(), "new_value", argValue)
}

// PropertyValuesStruct creates an uninitialised PropertyValues.
func PropertyValuesStruct() *PropertyValues {
	err := propertyValuesStruct_Set()
	if err != nil {
		return nil
	}

	structGo := PropertyValuesNewFromNative(propertyValuesStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizePropertyValues)
	return structGo
}
func finalizePropertyValues(obj *PropertyValues) {
	propertyValuesStruct.Free(obj.Native())
}

var rangeStruct *gi.Struct
var rangeStruct_Once sync.Once

func rangeStruct_Set() error {
	var err error
	rangeStruct_Once.Do(func() {
		rangeStruct, err = gi.StructNew("Atk", "Range")
	})
	return err
}

type Range struct {
	native unsafe.Pointer
}

func RangeNewFromNative(native unsafe.Pointer) *Range {
	instance := &Range{native: native}

	return instance
}

/*
CastToRange down casts any arbitrary Object to Range.
Exercise care, as this is a potentially dangerous function
if the Object is not a Range.
*/
func CastToRange(object *gobject.Object) *Range {
	return RangeNewFromNative(object.Native())
}

// Equals compares this Range with another Range, and returns true if they represent the same Object.
func (recv *Range) Equals(other *Range) bool {
	return other.Native() == recv.Native()
}

func (recv *Range) Native() unsafe.Pointer {
	return recv.native
}

var rangeNewFunction *gi.Function
var rangeNewFunction_Once sync.Once

func rangeNewFunction_Set() error {
	var err error
	rangeNewFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeNewFunction, err = rangeStruct.InvokerNew("new")
	})
	return err
}

// RangeNew is a representation of the C type atk_range_new.
func RangeNew(lowerLimit float64, upperLimit float64, description string) *Range {
	var inArgs [3]gi.Argument
	inArgs[0].SetFloat64(lowerLimit)
	inArgs[1].SetFloat64(upperLimit)
	inArgs[2].SetString(description)

	var ret gi.Argument

	err := rangeNewFunction_Set()
	if err == nil {
		ret = rangeNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := RangeNewFromNative(ret.Pointer())

	return retGo
}

var rangeCopyFunction *gi.Function
var rangeCopyFunction_Once sync.Once

func rangeCopyFunction_Set() error {
	var err error
	rangeCopyFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeCopyFunction, err = rangeStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type atk_range_copy.
func (recv *Range) Copy() *Range {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := rangeCopyFunction_Set()
	if err == nil {
		ret = rangeCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := RangeNewFromNative(ret.Pointer())

	return retGo
}

var rangeFreeFunction *gi.Function
var rangeFreeFunction_Once sync.Once

func rangeFreeFunction_Set() error {
	var err error
	rangeFreeFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeFreeFunction, err = rangeStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type atk_range_free.
func (recv *Range) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := rangeFreeFunction_Set()
	if err == nil {
		rangeFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var rangeGetDescriptionFunction *gi.Function
var rangeGetDescriptionFunction_Once sync.Once

func rangeGetDescriptionFunction_Set() error {
	var err error
	rangeGetDescriptionFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeGetDescriptionFunction, err = rangeStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type atk_range_get_description.
func (recv *Range) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := rangeGetDescriptionFunction_Set()
	if err == nil {
		ret = rangeGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var rangeGetLowerLimitFunction *gi.Function
var rangeGetLowerLimitFunction_Once sync.Once

func rangeGetLowerLimitFunction_Set() error {
	var err error
	rangeGetLowerLimitFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeGetLowerLimitFunction, err = rangeStruct.InvokerNew("get_lower_limit")
	})
	return err
}

// GetLowerLimit is a representation of the C type atk_range_get_lower_limit.
func (recv *Range) GetLowerLimit() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := rangeGetLowerLimitFunction_Set()
	if err == nil {
		ret = rangeGetLowerLimitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var rangeGetUpperLimitFunction *gi.Function
var rangeGetUpperLimitFunction_Once sync.Once

func rangeGetUpperLimitFunction_Set() error {
	var err error
	rangeGetUpperLimitFunction_Once.Do(func() {
		err = rangeStruct_Set()
		if err != nil {
			return
		}
		rangeGetUpperLimitFunction, err = rangeStruct.InvokerNew("get_upper_limit")
	})
	return err
}

// GetUpperLimit is a representation of the C type atk_range_get_upper_limit.
func (recv *Range) GetUpperLimit() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := rangeGetUpperLimitFunction_Set()
	if err == nil {
		ret = rangeGetUpperLimitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var rectangleStruct *gi.Struct
var rectangleStruct_Once sync.Once

func rectangleStruct_Set() error {
	var err error
	rectangleStruct_Once.Do(func() {
		rectangleStruct, err = gi.StructNew("Atk", "Rectangle")
	})
	return err
}

type Rectangle struct {
	native unsafe.Pointer
}

func RectangleNewFromNative(native unsafe.Pointer) *Rectangle {
	instance := &Rectangle{native: native}

	return instance
}

/*
CastToRectangle down casts any arbitrary Object to Rectangle.
Exercise care, as this is a potentially dangerous function
if the Object is not a Rectangle.
*/
func CastToRectangle(object *gobject.Object) *Rectangle {
	return RectangleNewFromNative(object.Native())
}

// Equals compares this Rectangle with another Rectangle, and returns true if they represent the same Object.
func (recv *Rectangle) Equals(other *Rectangle) bool {
	return other.Native() == recv.Native()
}

func (recv *Rectangle) Native() unsafe.Pointer {
	return recv.native
}

// FieldX returns the C field 'x'.
func (recv *Rectangle) FieldX() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.Native(), "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *Rectangle) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *Rectangle) FieldY() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.Native(), "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *Rectangle) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.Native(), "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *Rectangle) FieldWidth() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.Native(), "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *Rectangle) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.Native(), "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *Rectangle) FieldHeight() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.Native(), "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *Rectangle) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.Native(), "height", argValue)
}

// RectangleStruct creates an uninitialised Rectangle.
func RectangleStruct() *Rectangle {
	err := rectangleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RectangleNewFromNative(rectangleStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRectangle)
	return structGo
}
func finalizeRectangle(obj *Rectangle) {
	rectangleStruct.Free(obj.Native())
}

var registryClassStruct *gi.Struct
var registryClassStruct_Once sync.Once

func registryClassStruct_Set() error {
	var err error
	registryClassStruct_Once.Do(func() {
		registryClassStruct, err = gi.StructNew("Atk", "RegistryClass")
	})
	return err
}

type RegistryClass struct {
	native unsafe.Pointer
}

func RegistryClassNewFromNative(native unsafe.Pointer) *RegistryClass {
	instance := &RegistryClass{native: native}

	return instance
}

/*
CastToRegistryClass down casts any arbitrary Object to RegistryClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a RegistryClass.
*/
func CastToRegistryClass(object *gobject.Object) *RegistryClass {
	return RegistryClassNewFromNative(object.Native())
}

// Equals compares this RegistryClass with another RegistryClass, and returns true if they represent the same Object.
func (recv *RegistryClass) Equals(other *RegistryClass) bool {
	return other.Native() == recv.Native()
}

func (recv *RegistryClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *RegistryClass) FieldParentClass() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(registryClassStruct, recv.Native(), "parent_class")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RegistryClass) SetFieldParentClass(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(registryClassStruct, recv.Native(), "parent_class", argValue)
}

// RegistryClassStruct creates an uninitialised RegistryClass.
func RegistryClassStruct() *RegistryClass {
	err := registryClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RegistryClassNewFromNative(registryClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRegistryClass)
	return structGo
}
func finalizeRegistryClass(obj *RegistryClass) {
	registryClassStruct.Free(obj.Native())
}

var relationClassStruct *gi.Struct
var relationClassStruct_Once sync.Once

func relationClassStruct_Set() error {
	var err error
	relationClassStruct_Once.Do(func() {
		relationClassStruct, err = gi.StructNew("Atk", "RelationClass")
	})
	return err
}

type RelationClass struct {
	native unsafe.Pointer
}

func RelationClassNewFromNative(native unsafe.Pointer) *RelationClass {
	instance := &RelationClass{native: native}

	return instance
}

/*
CastToRelationClass down casts any arbitrary Object to RelationClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a RelationClass.
*/
func CastToRelationClass(object *gobject.Object) *RelationClass {
	return RelationClassNewFromNative(object.Native())
}

// Equals compares this RelationClass with another RelationClass, and returns true if they represent the same Object.
func (recv *RelationClass) Equals(other *RelationClass) bool {
	return other.Native() == recv.Native()
}

func (recv *RelationClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RelationClass) FieldParent() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(relationClassStruct, recv.Native(), "parent")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RelationClass) SetFieldParent(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(relationClassStruct, recv.Native(), "parent", argValue)
}

// RelationClassStruct creates an uninitialised RelationClass.
func RelationClassStruct() *RelationClass {
	err := relationClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RelationClassNewFromNative(relationClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRelationClass)
	return structGo
}
func finalizeRelationClass(obj *RelationClass) {
	relationClassStruct.Free(obj.Native())
}

var relationSetClassStruct *gi.Struct
var relationSetClassStruct_Once sync.Once

func relationSetClassStruct_Set() error {
	var err error
	relationSetClassStruct_Once.Do(func() {
		relationSetClassStruct, err = gi.StructNew("Atk", "RelationSetClass")
	})
	return err
}

type RelationSetClass struct {
	native unsafe.Pointer
}

func RelationSetClassNewFromNative(native unsafe.Pointer) *RelationSetClass {
	instance := &RelationSetClass{native: native}

	return instance
}

/*
CastToRelationSetClass down casts any arbitrary Object to RelationSetClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a RelationSetClass.
*/
func CastToRelationSetClass(object *gobject.Object) *RelationSetClass {
	return RelationSetClassNewFromNative(object.Native())
}

// Equals compares this RelationSetClass with another RelationSetClass, and returns true if they represent the same Object.
func (recv *RelationSetClass) Equals(other *RelationSetClass) bool {
	return other.Native() == recv.Native()
}

func (recv *RelationSetClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *RelationSetClass) FieldParent() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(relationSetClassStruct, recv.Native(), "parent")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *RelationSetClass) SetFieldParent(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(relationSetClassStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad1' : for field setter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field setter : no Go type for 'Function'

// RelationSetClassStruct creates an uninitialised RelationSetClass.
func RelationSetClassStruct() *RelationSetClass {
	err := relationSetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RelationSetClassNewFromNative(relationSetClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRelationSetClass)
	return structGo
}
func finalizeRelationSetClass(obj *RelationSetClass) {
	relationSetClassStruct.Free(obj.Native())
}

var selectionIfaceStruct *gi.Struct
var selectionIfaceStruct_Once sync.Once

func selectionIfaceStruct_Set() error {
	var err error
	selectionIfaceStruct_Once.Do(func() {
		selectionIfaceStruct, err = gi.StructNew("Atk", "SelectionIface")
	})
	return err
}

type SelectionIface struct {
	native unsafe.Pointer
}

func SelectionIfaceNewFromNative(native unsafe.Pointer) *SelectionIface {
	instance := &SelectionIface{native: native}

	return instance
}

/*
CastToSelectionIface down casts any arbitrary Object to SelectionIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a SelectionIface.
*/
func CastToSelectionIface(object *gobject.Object) *SelectionIface {
	return SelectionIfaceNewFromNative(object.Native())
}

// Equals compares this SelectionIface with another SelectionIface, and returns true if they represent the same Object.
func (recv *SelectionIface) Equals(other *SelectionIface) bool {
	return other.Native() == recv.Native()
}

func (recv *SelectionIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *SelectionIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(selectionIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *SelectionIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(selectionIfaceStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'add_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'add_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'clear_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'clear_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'ref_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'get_selection_count' : for field getter : missing Type

// UNSUPPORTED : C value 'get_selection_count' : for field setter : missing Type

// UNSUPPORTED : C value 'is_child_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'is_child_selected' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'select_all_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'select_all_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'selection_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_changed' : for field setter : missing Type

// SelectionIfaceStruct creates an uninitialised SelectionIface.
func SelectionIfaceStruct() *SelectionIface {
	err := selectionIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SelectionIfaceNewFromNative(selectionIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSelectionIface)
	return structGo
}
func finalizeSelectionIface(obj *SelectionIface) {
	selectionIfaceStruct.Free(obj.Native())
}

var socketClassStruct *gi.Struct
var socketClassStruct_Once sync.Once

func socketClassStruct_Set() error {
	var err error
	socketClassStruct_Once.Do(func() {
		socketClassStruct, err = gi.StructNew("Atk", "SocketClass")
	})
	return err
}

type SocketClass struct {
	native unsafe.Pointer
}

func SocketClassNewFromNative(native unsafe.Pointer) *SocketClass {
	instance := &SocketClass{native: native}

	return instance
}

/*
CastToSocketClass down casts any arbitrary Object to SocketClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a SocketClass.
*/
func CastToSocketClass(object *gobject.Object) *SocketClass {
	return SocketClassNewFromNative(object.Native())
}

// Equals compares this SocketClass with another SocketClass, and returns true if they represent the same Object.
func (recv *SocketClass) Equals(other *SocketClass) bool {
	return other.Native() == recv.Native()
}

func (recv *SocketClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParentClass returns the C field 'parent_class'.
func (recv *SocketClass) FieldParentClass() *ObjectClass {
	argValue := gi.StructFieldGet(socketClassStruct, recv.Native(), "parent_class")
	value := ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SocketClass) SetFieldParentClass(value *ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(socketClassStruct, recv.Native(), "parent_class", argValue)
}

// UNSUPPORTED : C value 'embed' : for field getter : missing Type

// UNSUPPORTED : C value 'embed' : for field setter : missing Type

// SocketClassStruct creates an uninitialised SocketClass.
func SocketClassStruct() *SocketClass {
	err := socketClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := SocketClassNewFromNative(socketClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeSocketClass)
	return structGo
}
func finalizeSocketClass(obj *SocketClass) {
	socketClassStruct.Free(obj.Native())
}

var stateSetClassStruct *gi.Struct
var stateSetClassStruct_Once sync.Once

func stateSetClassStruct_Set() error {
	var err error
	stateSetClassStruct_Once.Do(func() {
		stateSetClassStruct, err = gi.StructNew("Atk", "StateSetClass")
	})
	return err
}

type StateSetClass struct {
	native unsafe.Pointer
}

func StateSetClassNewFromNative(native unsafe.Pointer) *StateSetClass {
	instance := &StateSetClass{native: native}

	return instance
}

/*
CastToStateSetClass down casts any arbitrary Object to StateSetClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a StateSetClass.
*/
func CastToStateSetClass(object *gobject.Object) *StateSetClass {
	return StateSetClassNewFromNative(object.Native())
}

// Equals compares this StateSetClass with another StateSetClass, and returns true if they represent the same Object.
func (recv *StateSetClass) Equals(other *StateSetClass) bool {
	return other.Native() == recv.Native()
}

func (recv *StateSetClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *StateSetClass) FieldParent() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(stateSetClassStruct, recv.Native(), "parent")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *StateSetClass) SetFieldParent(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(stateSetClassStruct, recv.Native(), "parent", argValue)
}

// StateSetClassStruct creates an uninitialised StateSetClass.
func StateSetClassStruct() *StateSetClass {
	err := stateSetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StateSetClassNewFromNative(stateSetClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStateSetClass)
	return structGo
}
func finalizeStateSetClass(obj *StateSetClass) {
	stateSetClassStruct.Free(obj.Native())
}

var streamableContentIfaceStruct *gi.Struct
var streamableContentIfaceStruct_Once sync.Once

func streamableContentIfaceStruct_Set() error {
	var err error
	streamableContentIfaceStruct_Once.Do(func() {
		streamableContentIfaceStruct, err = gi.StructNew("Atk", "StreamableContentIface")
	})
	return err
}

type StreamableContentIface struct {
	native unsafe.Pointer
}

func StreamableContentIfaceNewFromNative(native unsafe.Pointer) *StreamableContentIface {
	instance := &StreamableContentIface{native: native}

	return instance
}

/*
CastToStreamableContentIface down casts any arbitrary Object to StreamableContentIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a StreamableContentIface.
*/
func CastToStreamableContentIface(object *gobject.Object) *StreamableContentIface {
	return StreamableContentIfaceNewFromNative(object.Native())
}

// Equals compares this StreamableContentIface with another StreamableContentIface, and returns true if they represent the same Object.
func (recv *StreamableContentIface) Equals(other *StreamableContentIface) bool {
	return other.Native() == recv.Native()
}

func (recv *StreamableContentIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *StreamableContentIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(streamableContentIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *StreamableContentIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(streamableContentIfaceStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'get_n_mime_types' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_mime_types' : for field setter : missing Type

// UNSUPPORTED : C value 'get_mime_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mime_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_stream' : for field getter : missing Type

// UNSUPPORTED : C value 'get_stream' : for field setter : missing Type

// UNSUPPORTED : C value 'get_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'get_uri' : for field setter : missing Type

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad1' : for field setter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field setter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad3' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad3' : for field setter : no Go type for 'Function'

// StreamableContentIfaceStruct creates an uninitialised StreamableContentIface.
func StreamableContentIfaceStruct() *StreamableContentIface {
	err := streamableContentIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := StreamableContentIfaceNewFromNative(streamableContentIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeStreamableContentIface)
	return structGo
}
func finalizeStreamableContentIface(obj *StreamableContentIface) {
	streamableContentIfaceStruct.Free(obj.Native())
}

var tableCellIfaceStruct *gi.Struct
var tableCellIfaceStruct_Once sync.Once

func tableCellIfaceStruct_Set() error {
	var err error
	tableCellIfaceStruct_Once.Do(func() {
		tableCellIfaceStruct, err = gi.StructNew("Atk", "TableCellIface")
	})
	return err
}

type TableCellIface struct {
	native unsafe.Pointer
}

func TableCellIfaceNewFromNative(native unsafe.Pointer) *TableCellIface {
	instance := &TableCellIface{native: native}

	return instance
}

/*
CastToTableCellIface down casts any arbitrary Object to TableCellIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a TableCellIface.
*/
func CastToTableCellIface(object *gobject.Object) *TableCellIface {
	return TableCellIfaceNewFromNative(object.Native())
}

// Equals compares this TableCellIface with another TableCellIface, and returns true if they represent the same Object.
func (recv *TableCellIface) Equals(other *TableCellIface) bool {
	return other.Native() == recv.Native()
}

func (recv *TableCellIface) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'get_column_span' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_span' : for field setter : missing Type

// UNSUPPORTED : C value 'get_column_header_cells' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_header_cells' : for field setter : missing Type

// UNSUPPORTED : C value 'get_position' : for field getter : missing Type

// UNSUPPORTED : C value 'get_position' : for field setter : missing Type

// UNSUPPORTED : C value 'get_row_span' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_span' : for field setter : missing Type

// UNSUPPORTED : C value 'get_row_header_cells' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_header_cells' : for field setter : missing Type

// UNSUPPORTED : C value 'get_row_column_span' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_column_span' : for field setter : missing Type

// UNSUPPORTED : C value 'get_table' : for field getter : missing Type

// UNSUPPORTED : C value 'get_table' : for field setter : missing Type

// TableCellIfaceStruct creates an uninitialised TableCellIface.
func TableCellIfaceStruct() *TableCellIface {
	err := tableCellIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TableCellIfaceNewFromNative(tableCellIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTableCellIface)
	return structGo
}
func finalizeTableCellIface(obj *TableCellIface) {
	tableCellIfaceStruct.Free(obj.Native())
}

var tableIfaceStruct *gi.Struct
var tableIfaceStruct_Once sync.Once

func tableIfaceStruct_Set() error {
	var err error
	tableIfaceStruct_Once.Do(func() {
		tableIfaceStruct, err = gi.StructNew("Atk", "TableIface")
	})
	return err
}

type TableIface struct {
	native unsafe.Pointer
}

func TableIfaceNewFromNative(native unsafe.Pointer) *TableIface {
	instance := &TableIface{native: native}

	return instance
}

/*
CastToTableIface down casts any arbitrary Object to TableIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a TableIface.
*/
func CastToTableIface(object *gobject.Object) *TableIface {
	return TableIfaceNewFromNative(object.Native())
}

// Equals compares this TableIface with another TableIface, and returns true if they represent the same Object.
func (recv *TableIface) Equals(other *TableIface) bool {
	return other.Native() == recv.Native()
}

func (recv *TableIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *TableIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(tableIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *TableIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(tableIfaceStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'ref_at' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_at' : for field setter : missing Type

// UNSUPPORTED : C value 'get_index_at' : for field getter : missing Type

// UNSUPPORTED : C value 'get_index_at' : for field setter : missing Type

// UNSUPPORTED : C value 'get_column_at_index' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_at_index' : for field setter : missing Type

// UNSUPPORTED : C value 'get_row_at_index' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_at_index' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_columns' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_columns' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_rows' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_rows' : for field setter : missing Type

// UNSUPPORTED : C value 'get_column_extent_at' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_extent_at' : for field setter : missing Type

// UNSUPPORTED : C value 'get_row_extent_at' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_extent_at' : for field setter : missing Type

// UNSUPPORTED : C value 'get_caption' : for field getter : missing Type

// UNSUPPORTED : C value 'get_caption' : for field setter : missing Type

// UNSUPPORTED : C value 'get_column_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_description' : for field setter : missing Type

// UNSUPPORTED : C value 'get_column_header' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_header' : for field setter : missing Type

// UNSUPPORTED : C value 'get_row_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_description' : for field setter : missing Type

// UNSUPPORTED : C value 'get_row_header' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_header' : for field setter : missing Type

// UNSUPPORTED : C value 'get_summary' : for field getter : missing Type

// UNSUPPORTED : C value 'get_summary' : for field setter : missing Type

// UNSUPPORTED : C value 'set_caption' : for field getter : missing Type

// UNSUPPORTED : C value 'set_caption' : for field setter : missing Type

// UNSUPPORTED : C value 'set_column_description' : for field getter : missing Type

// UNSUPPORTED : C value 'set_column_description' : for field setter : missing Type

// UNSUPPORTED : C value 'set_column_header' : for field getter : missing Type

// UNSUPPORTED : C value 'set_column_header' : for field setter : missing Type

// UNSUPPORTED : C value 'set_row_description' : for field getter : missing Type

// UNSUPPORTED : C value 'set_row_description' : for field setter : missing Type

// UNSUPPORTED : C value 'set_row_header' : for field getter : missing Type

// UNSUPPORTED : C value 'set_row_header' : for field setter : missing Type

// UNSUPPORTED : C value 'set_summary' : for field getter : missing Type

// UNSUPPORTED : C value 'set_summary' : for field setter : missing Type

// UNSUPPORTED : C value 'get_selected_columns' : for field getter : missing Type

// UNSUPPORTED : C value 'get_selected_columns' : for field setter : missing Type

// UNSUPPORTED : C value 'get_selected_rows' : for field getter : missing Type

// UNSUPPORTED : C value 'get_selected_rows' : for field setter : missing Type

// UNSUPPORTED : C value 'is_column_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'is_column_selected' : for field setter : missing Type

// UNSUPPORTED : C value 'is_row_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'is_row_selected' : for field setter : missing Type

// UNSUPPORTED : C value 'is_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'is_selected' : for field setter : missing Type

// UNSUPPORTED : C value 'add_row_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'add_row_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_row_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_row_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'add_column_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'add_column_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_column_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_column_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'row_inserted' : for field getter : missing Type

// UNSUPPORTED : C value 'row_inserted' : for field setter : missing Type

// UNSUPPORTED : C value 'column_inserted' : for field getter : missing Type

// UNSUPPORTED : C value 'column_inserted' : for field setter : missing Type

// UNSUPPORTED : C value 'row_deleted' : for field getter : missing Type

// UNSUPPORTED : C value 'row_deleted' : for field setter : missing Type

// UNSUPPORTED : C value 'column_deleted' : for field getter : missing Type

// UNSUPPORTED : C value 'column_deleted' : for field setter : missing Type

// UNSUPPORTED : C value 'row_reordered' : for field getter : missing Type

// UNSUPPORTED : C value 'row_reordered' : for field setter : missing Type

// UNSUPPORTED : C value 'column_reordered' : for field getter : missing Type

// UNSUPPORTED : C value 'column_reordered' : for field setter : missing Type

// UNSUPPORTED : C value 'model_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'model_changed' : for field setter : missing Type

// TableIfaceStruct creates an uninitialised TableIface.
func TableIfaceStruct() *TableIface {
	err := tableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TableIfaceNewFromNative(tableIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTableIface)
	return structGo
}
func finalizeTableIface(obj *TableIface) {
	tableIfaceStruct.Free(obj.Native())
}

var textIfaceStruct *gi.Struct
var textIfaceStruct_Once sync.Once

func textIfaceStruct_Set() error {
	var err error
	textIfaceStruct_Once.Do(func() {
		textIfaceStruct, err = gi.StructNew("Atk", "TextIface")
	})
	return err
}

type TextIface struct {
	native unsafe.Pointer
}

func TextIfaceNewFromNative(native unsafe.Pointer) *TextIface {
	instance := &TextIface{native: native}

	return instance
}

/*
CastToTextIface down casts any arbitrary Object to TextIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a TextIface.
*/
func CastToTextIface(object *gobject.Object) *TextIface {
	return TextIfaceNewFromNative(object.Native())
}

// Equals compares this TextIface with another TextIface, and returns true if they represent the same Object.
func (recv *TextIface) Equals(other *TextIface) bool {
	return other.Native() == recv.Native()
}

func (recv *TextIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *TextIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(textIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *TextIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(textIfaceStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'get_text' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text' : for field setter : missing Type

// UNSUPPORTED : C value 'get_text_after_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_after_offset' : for field setter : missing Type

// UNSUPPORTED : C value 'get_text_at_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_at_offset' : for field setter : missing Type

// UNSUPPORTED : C value 'get_character_at_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_character_at_offset' : for field setter : missing Type

// UNSUPPORTED : C value 'get_text_before_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_before_offset' : for field setter : missing Type

// UNSUPPORTED : C value 'get_caret_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_caret_offset' : for field setter : missing Type

// UNSUPPORTED : C value 'get_run_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_run_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'get_default_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_default_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'get_character_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'get_character_extents' : for field setter : missing Type

// UNSUPPORTED : C value 'get_character_count' : for field getter : missing Type

// UNSUPPORTED : C value 'get_character_count' : for field setter : missing Type

// UNSUPPORTED : C value 'get_offset_at_point' : for field getter : missing Type

// UNSUPPORTED : C value 'get_offset_at_point' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_selections' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_selections' : for field setter : missing Type

// UNSUPPORTED : C value 'get_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'get_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'add_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'add_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'set_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'set_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'set_caret_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'set_caret_offset' : for field setter : missing Type

// UNSUPPORTED : C value 'text_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'text_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'text_caret_moved' : for field getter : missing Type

// UNSUPPORTED : C value 'text_caret_moved' : for field setter : missing Type

// UNSUPPORTED : C value 'text_selection_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'text_selection_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'text_attributes_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'text_attributes_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'get_range_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'get_range_extents' : for field setter : missing Type

// UNSUPPORTED : C value 'get_bounded_ranges' : for field getter : missing Type

// UNSUPPORTED : C value 'get_bounded_ranges' : for field setter : missing Type

// UNSUPPORTED : C value 'get_string_at_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_string_at_offset' : for field setter : missing Type

// UNSUPPORTED : C value 'scroll_substring_to' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_substring_to' : for field setter : missing Type

// UNSUPPORTED : C value 'scroll_substring_to_point' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_substring_to_point' : for field setter : missing Type

// TextIfaceStruct creates an uninitialised TextIface.
func TextIfaceStruct() *TextIface {
	err := textIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TextIfaceNewFromNative(textIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTextIface)
	return structGo
}
func finalizeTextIface(obj *TextIface) {
	textIfaceStruct.Free(obj.Native())
}

var textRangeStruct *gi.Struct
var textRangeStruct_Once sync.Once

func textRangeStruct_Set() error {
	var err error
	textRangeStruct_Once.Do(func() {
		textRangeStruct, err = gi.StructNew("Atk", "TextRange")
	})
	return err
}

type TextRange struct {
	native unsafe.Pointer
}

func TextRangeNewFromNative(native unsafe.Pointer) *TextRange {
	instance := &TextRange{native: native}

	return instance
}

/*
CastToTextRange down casts any arbitrary Object to TextRange.
Exercise care, as this is a potentially dangerous function
if the Object is not a TextRange.
*/
func CastToTextRange(object *gobject.Object) *TextRange {
	return TextRangeNewFromNative(object.Native())
}

// Equals compares this TextRange with another TextRange, and returns true if they represent the same Object.
func (recv *TextRange) Equals(other *TextRange) bool {
	return other.Native() == recv.Native()
}

func (recv *TextRange) Native() unsafe.Pointer {
	return recv.native
}

// FieldBounds returns the C field 'bounds'.
func (recv *TextRange) FieldBounds() *TextRectangle {
	argValue := gi.StructFieldGet(textRangeStruct, recv.Native(), "bounds")
	value := TextRectangleNewFromNative(argValue.Pointer())
	return value
}

// SetFieldBounds sets the value of the C field 'bounds'.
func (recv *TextRange) SetFieldBounds(value *TextRectangle) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(textRangeStruct, recv.Native(), "bounds", argValue)
}

// FieldStartOffset returns the C field 'start_offset'.
func (recv *TextRange) FieldStartOffset() int32 {
	argValue := gi.StructFieldGet(textRangeStruct, recv.Native(), "start_offset")
	value := argValue.Int32()
	return value
}

// SetFieldStartOffset sets the value of the C field 'start_offset'.
func (recv *TextRange) SetFieldStartOffset(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(textRangeStruct, recv.Native(), "start_offset", argValue)
}

// FieldEndOffset returns the C field 'end_offset'.
func (recv *TextRange) FieldEndOffset() int32 {
	argValue := gi.StructFieldGet(textRangeStruct, recv.Native(), "end_offset")
	value := argValue.Int32()
	return value
}

// SetFieldEndOffset sets the value of the C field 'end_offset'.
func (recv *TextRange) SetFieldEndOffset(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(textRangeStruct, recv.Native(), "end_offset", argValue)
}

// FieldContent returns the C field 'content'.
func (recv *TextRange) FieldContent() string {
	argValue := gi.StructFieldGet(textRangeStruct, recv.Native(), "content")
	value := argValue.String(false)
	return value
}

// SetFieldContent sets the value of the C field 'content'.
func (recv *TextRange) SetFieldContent(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(textRangeStruct, recv.Native(), "content", argValue)
}

// TextRangeStruct creates an uninitialised TextRange.
func TextRangeStruct() *TextRange {
	err := textRangeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TextRangeNewFromNative(textRangeStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTextRange)
	return structGo
}
func finalizeTextRange(obj *TextRange) {
	textRangeStruct.Free(obj.Native())
}

var textRectangleStruct *gi.Struct
var textRectangleStruct_Once sync.Once

func textRectangleStruct_Set() error {
	var err error
	textRectangleStruct_Once.Do(func() {
		textRectangleStruct, err = gi.StructNew("Atk", "TextRectangle")
	})
	return err
}

type TextRectangle struct {
	native unsafe.Pointer
}

func TextRectangleNewFromNative(native unsafe.Pointer) *TextRectangle {
	instance := &TextRectangle{native: native}

	return instance
}

/*
CastToTextRectangle down casts any arbitrary Object to TextRectangle.
Exercise care, as this is a potentially dangerous function
if the Object is not a TextRectangle.
*/
func CastToTextRectangle(object *gobject.Object) *TextRectangle {
	return TextRectangleNewFromNative(object.Native())
}

// Equals compares this TextRectangle with another TextRectangle, and returns true if they represent the same Object.
func (recv *TextRectangle) Equals(other *TextRectangle) bool {
	return other.Native() == recv.Native()
}

func (recv *TextRectangle) Native() unsafe.Pointer {
	return recv.native
}

// FieldX returns the C field 'x'.
func (recv *TextRectangle) FieldX() int32 {
	argValue := gi.StructFieldGet(textRectangleStruct, recv.Native(), "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *TextRectangle) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(textRectangleStruct, recv.Native(), "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *TextRectangle) FieldY() int32 {
	argValue := gi.StructFieldGet(textRectangleStruct, recv.Native(), "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *TextRectangle) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(textRectangleStruct, recv.Native(), "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *TextRectangle) FieldWidth() int32 {
	argValue := gi.StructFieldGet(textRectangleStruct, recv.Native(), "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *TextRectangle) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(textRectangleStruct, recv.Native(), "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *TextRectangle) FieldHeight() int32 {
	argValue := gi.StructFieldGet(textRectangleStruct, recv.Native(), "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *TextRectangle) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(textRectangleStruct, recv.Native(), "height", argValue)
}

// TextRectangleStruct creates an uninitialised TextRectangle.
func TextRectangleStruct() *TextRectangle {
	err := textRectangleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := TextRectangleNewFromNative(textRectangleStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeTextRectangle)
	return structGo
}
func finalizeTextRectangle(obj *TextRectangle) {
	textRectangleStruct.Free(obj.Native())
}

var utilClassStruct *gi.Struct
var utilClassStruct_Once sync.Once

func utilClassStruct_Set() error {
	var err error
	utilClassStruct_Once.Do(func() {
		utilClassStruct, err = gi.StructNew("Atk", "UtilClass")
	})
	return err
}

type UtilClass struct {
	native unsafe.Pointer
}

func UtilClassNewFromNative(native unsafe.Pointer) *UtilClass {
	instance := &UtilClass{native: native}

	return instance
}

/*
CastToUtilClass down casts any arbitrary Object to UtilClass.
Exercise care, as this is a potentially dangerous function
if the Object is not a UtilClass.
*/
func CastToUtilClass(object *gobject.Object) *UtilClass {
	return UtilClassNewFromNative(object.Native())
}

// Equals compares this UtilClass with another UtilClass, and returns true if they represent the same Object.
func (recv *UtilClass) Equals(other *UtilClass) bool {
	return other.Native() == recv.Native()
}

func (recv *UtilClass) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *UtilClass) FieldParent() *gobject.ObjectClass {
	argValue := gi.StructFieldGet(utilClassStruct, recv.Native(), "parent")
	value := gobject.ObjectClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *UtilClass) SetFieldParent(value *gobject.ObjectClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(utilClassStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'add_global_event_listener' : for field getter : missing Type

// UNSUPPORTED : C value 'add_global_event_listener' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_global_event_listener' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_global_event_listener' : for field setter : missing Type

// UNSUPPORTED : C value 'add_key_event_listener' : for field getter : missing Type

// UNSUPPORTED : C value 'add_key_event_listener' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_key_event_listener' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_key_event_listener' : for field setter : missing Type

// UNSUPPORTED : C value 'get_root' : for field getter : missing Type

// UNSUPPORTED : C value 'get_root' : for field setter : missing Type

// UNSUPPORTED : C value 'get_toolkit_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_toolkit_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_toolkit_version' : for field getter : missing Type

// UNSUPPORTED : C value 'get_toolkit_version' : for field setter : missing Type

// UtilClassStruct creates an uninitialised UtilClass.
func UtilClassStruct() *UtilClass {
	err := utilClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := UtilClassNewFromNative(utilClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeUtilClass)
	return structGo
}
func finalizeUtilClass(obj *UtilClass) {
	utilClassStruct.Free(obj.Native())
}

var valueIfaceStruct *gi.Struct
var valueIfaceStruct_Once sync.Once

func valueIfaceStruct_Set() error {
	var err error
	valueIfaceStruct_Once.Do(func() {
		valueIfaceStruct, err = gi.StructNew("Atk", "ValueIface")
	})
	return err
}

type ValueIface struct {
	native unsafe.Pointer
}

func ValueIfaceNewFromNative(native unsafe.Pointer) *ValueIface {
	instance := &ValueIface{native: native}

	return instance
}

/*
CastToValueIface down casts any arbitrary Object to ValueIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ValueIface.
*/
func CastToValueIface(object *gobject.Object) *ValueIface {
	return ValueIfaceNewFromNative(object.Native())
}

// Equals compares this ValueIface with another ValueIface, and returns true if they represent the same Object.
func (recv *ValueIface) Equals(other *ValueIface) bool {
	return other.Native() == recv.Native()
}

func (recv *ValueIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *ValueIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(valueIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *ValueIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(valueIfaceStruct, recv.Native(), "parent", argValue)
}

// UNSUPPORTED : C value 'get_current_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_current_value' : for field setter : missing Type

// UNSUPPORTED : C value 'get_maximum_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_maximum_value' : for field setter : missing Type

// UNSUPPORTED : C value 'get_minimum_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_minimum_value' : for field setter : missing Type

// UNSUPPORTED : C value 'set_current_value' : for field getter : missing Type

// UNSUPPORTED : C value 'set_current_value' : for field setter : missing Type

// UNSUPPORTED : C value 'get_minimum_increment' : for field getter : missing Type

// UNSUPPORTED : C value 'get_minimum_increment' : for field setter : missing Type

// UNSUPPORTED : C value 'get_value_and_text' : for field getter : missing Type

// UNSUPPORTED : C value 'get_value_and_text' : for field setter : missing Type

// UNSUPPORTED : C value 'get_range' : for field getter : missing Type

// UNSUPPORTED : C value 'get_range' : for field setter : missing Type

// UNSUPPORTED : C value 'get_increment' : for field getter : missing Type

// UNSUPPORTED : C value 'get_increment' : for field setter : missing Type

// UNSUPPORTED : C value 'get_sub_ranges' : for field getter : missing Type

// UNSUPPORTED : C value 'get_sub_ranges' : for field setter : missing Type

// UNSUPPORTED : C value 'set_value' : for field getter : missing Type

// UNSUPPORTED : C value 'set_value' : for field setter : missing Type

// ValueIfaceStruct creates an uninitialised ValueIface.
func ValueIfaceStruct() *ValueIface {
	err := valueIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ValueIfaceNewFromNative(valueIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeValueIface)
	return structGo
}
func finalizeValueIface(obj *ValueIface) {
	valueIfaceStruct.Free(obj.Native())
}

var windowIfaceStruct *gi.Struct
var windowIfaceStruct_Once sync.Once

func windowIfaceStruct_Set() error {
	var err error
	windowIfaceStruct_Once.Do(func() {
		windowIfaceStruct, err = gi.StructNew("Atk", "WindowIface")
	})
	return err
}

type WindowIface struct {
	native unsafe.Pointer
}

func WindowIfaceNewFromNative(native unsafe.Pointer) *WindowIface {
	instance := &WindowIface{native: native}

	return instance
}

/*
CastToWindowIface down casts any arbitrary Object to WindowIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a WindowIface.
*/
func CastToWindowIface(object *gobject.Object) *WindowIface {
	return WindowIfaceNewFromNative(object.Native())
}

// Equals compares this WindowIface with another WindowIface, and returns true if they represent the same Object.
func (recv *WindowIface) Equals(other *WindowIface) bool {
	return other.Native() == recv.Native()
}

func (recv *WindowIface) Native() unsafe.Pointer {
	return recv.native
}

// FieldParent returns the C field 'parent'.
func (recv *WindowIface) FieldParent() *gobject.TypeInterface {
	argValue := gi.StructFieldGet(windowIfaceStruct, recv.Native(), "parent")
	value := gobject.TypeInterfaceNewFromNative(argValue.Pointer())
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *WindowIface) SetFieldParent(value *gobject.TypeInterface) {
	var argValue gi.Argument
	argValue.SetPointer(value.Native())
	gi.StructFieldSet(windowIfaceStruct, recv.Native(), "parent", argValue)
}

// WindowIfaceStruct creates an uninitialised WindowIface.
func WindowIfaceStruct() *WindowIface {
	err := windowIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := WindowIfaceNewFromNative(windowIfaceStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeWindowIface)
	return structGo
}
func finalizeWindowIface(obj *WindowIface) {
	windowIfaceStruct.Free(obj.Native())
}
