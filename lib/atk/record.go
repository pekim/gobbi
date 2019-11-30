// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
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
	native uintptr
}

// UNSUPPORTED : C value 'do_action' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_actions' : for field getter : missing Type

// UNSUPPORTED : C value 'get_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_keybinding' : for field getter : missing Type

// UNSUPPORTED : C value 'set_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_localized_name' : for field getter : missing Type

// ActionIfaceStruct creates an uninitialised ActionIface.
func ActionIfaceStruct() *ActionIface {
	err := actionIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionIface{native: actionIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActionIface)
	return structGo
}
func finalizeActionIface(obj *ActionIface) {
	actionIfaceStruct.Free(obj.native)
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
	native uintptr
}

// Name returns the C field 'name'.
func (recv *Attribute) Name() string {
	argValue := gi.FieldGet(attributeStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// Value returns the C field 'value'.
func (recv *Attribute) Value() string {
	argValue := gi.FieldGet(attributeStruct, recv.native, "value")
	value := argValue.String(false)
	return value
}

// AttributeStruct creates an uninitialised Attribute.
func AttributeStruct() *Attribute {
	err := attributeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Attribute{native: attributeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAttribute)
	return structGo
}
func finalizeAttribute(obj *Attribute) {
	attributeStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'add_focus_handler' : for field getter : missing Type

// UNSUPPORTED : C value 'contains' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_accessible_at_point' : for field getter : missing Type

// UNSUPPORTED : C value 'get_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'get_position' : for field getter : missing Type

// UNSUPPORTED : C value 'get_size' : for field getter : missing Type

// UNSUPPORTED : C value 'grab_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_focus_handler' : for field getter : missing Type

// UNSUPPORTED : C value 'set_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'set_position' : for field getter : missing Type

// UNSUPPORTED : C value 'set_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_layer' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mdi_zorder' : for field getter : missing Type

// UNSUPPORTED : C value 'bounds_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'get_alpha' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_to' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_to_point' : for field getter : missing Type

// ComponentIfaceStruct creates an uninitialised ComponentIface.
func ComponentIfaceStruct() *ComponentIface {
	err := componentIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ComponentIface{native: componentIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeComponentIface)
	return structGo
}
func finalizeComponentIface(obj *ComponentIface) {
	componentIfaceStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_document_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_document' : for field getter : missing Type

// UNSUPPORTED : C value 'get_document_locale' : for field getter : missing Type

// UNSUPPORTED : C value 'get_document_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_document_attribute_value' : for field getter : missing Type

// UNSUPPORTED : C value 'set_document_attribute' : for field getter : missing Type

// UNSUPPORTED : C value 'get_current_page_number' : for field getter : missing Type

// UNSUPPORTED : C value 'get_page_count' : for field getter : missing Type

// DocumentIfaceStruct creates an uninitialised DocumentIface.
func DocumentIfaceStruct() *DocumentIface {
	err := documentIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DocumentIface{native: documentIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeDocumentIface)
	return structGo
}
func finalizeDocumentIface(obj *DocumentIface) {
	documentIfaceStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_interface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'set_run_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'set_text_contents' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_text' : for field getter : missing Type

// UNSUPPORTED : C value 'copy_text' : for field getter : missing Type

// UNSUPPORTED : C value 'cut_text' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_text' : for field getter : missing Type

// UNSUPPORTED : C value 'paste_text' : for field getter : missing Type

// EditableTextIfaceStruct creates an uninitialised EditableTextIface.
func EditableTextIfaceStruct() *EditableTextIface {
	err := editableTextIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EditableTextIface{native: editableTextIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEditableTextIface)
	return structGo
}
func finalizeEditableTextIface(obj *EditableTextIface) {
	editableTextIfaceStruct.Free(obj.native)
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
	native uintptr
}

// ParentClass returns the C field 'parent_class'.
func (recv *GObjectAccessibleClass) ParentClass() *ObjectClass {
	argValue := gi.FieldGet(gObjectAccessibleClassStruct, recv.native, "parent_class")
	value := &ObjectClass{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field getter : no Go type for 'Function'

// GObjectAccessibleClassStruct creates an uninitialised GObjectAccessibleClass.
func GObjectAccessibleClassStruct() *GObjectAccessibleClass {
	err := gObjectAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GObjectAccessibleClass{native: gObjectAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGObjectAccessibleClass)
	return structGo
}
func finalizeGObjectAccessibleClass(obj *GObjectAccessibleClass) {
	gObjectAccessibleClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'get_object' : for field getter : missing Type

// UNSUPPORTED : C value 'get_end_index' : for field getter : missing Type

// UNSUPPORTED : C value 'get_start_index' : for field getter : missing Type

// UNSUPPORTED : C value 'is_valid' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_anchors' : for field getter : missing Type

// UNSUPPORTED : C value 'link_state' : for field getter : missing Type

// UNSUPPORTED : C value 'is_selected_link' : for field getter : missing Type

// UNSUPPORTED : C value 'link_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// HyperlinkClassStruct creates an uninitialised HyperlinkClass.
func HyperlinkClassStruct() *HyperlinkClass {
	err := hyperlinkClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HyperlinkClass{native: hyperlinkClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHyperlinkClass)
	return structGo
}
func finalizeHyperlinkClass(obj *HyperlinkClass) {
	hyperlinkClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_hyperlink' : for field getter : missing Type

// HyperlinkImplIfaceStruct creates an uninitialised HyperlinkImplIface.
func HyperlinkImplIfaceStruct() *HyperlinkImplIface {
	err := hyperlinkImplIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HyperlinkImplIface{native: hyperlinkImplIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHyperlinkImplIface)
	return structGo
}
func finalizeHyperlinkImplIface(obj *HyperlinkImplIface) {
	hyperlinkImplIfaceStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_link' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_links' : for field getter : missing Type

// UNSUPPORTED : C value 'get_link_index' : for field getter : missing Type

// UNSUPPORTED : C value 'link_selected' : for field getter : missing Type

// HypertextIfaceStruct creates an uninitialised HypertextIface.
func HypertextIfaceStruct() *HypertextIface {
	err := hypertextIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HypertextIface{native: hypertextIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHypertextIface)
	return structGo
}
func finalizeHypertextIface(obj *HypertextIface) {
	hypertextIfaceStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_image_position' : for field getter : missing Type

// UNSUPPORTED : C value 'get_image_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_image_size' : for field getter : missing Type

// UNSUPPORTED : C value 'set_image_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_image_locale' : for field getter : missing Type

// ImageIfaceStruct creates an uninitialised ImageIface.
func ImageIfaceStruct() *ImageIface {
	err := imageIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ImageIface{native: imageIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImageIface)
	return structGo
}
func finalizeImageIface(obj *ImageIface) {
	imageIfaceStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'atk_implementor_ref_accessible' : return type 'Object' not supported

// ImplementorStruct creates an uninitialised Implementor.
func ImplementorStruct() *Implementor {
	err := implementorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Implementor{native: implementorStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImplementor)
	return structGo
}
func finalizeImplementor(obj *Implementor) {
	implementorStruct.Free(obj.native)
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
	native uintptr
}

// Type returns the C field 'type'.
func (recv *KeyEventStruct) Type() int32 {
	argValue := gi.FieldGet(keyEventStructStruct, recv.native, "type")
	value := argValue.Int32()
	return value
}

// State returns the C field 'state'.
func (recv *KeyEventStruct) State() uint32 {
	argValue := gi.FieldGet(keyEventStructStruct, recv.native, "state")
	value := argValue.Uint32()
	return value
}

// Keyval returns the C field 'keyval'.
func (recv *KeyEventStruct) Keyval() uint32 {
	argValue := gi.FieldGet(keyEventStructStruct, recv.native, "keyval")
	value := argValue.Uint32()
	return value
}

// Length returns the C field 'length'.
func (recv *KeyEventStruct) Length() int32 {
	argValue := gi.FieldGet(keyEventStructStruct, recv.native, "length")
	value := argValue.Int32()
	return value
}

// String returns the C field 'string'.
func (recv *KeyEventStruct) String() string {
	argValue := gi.FieldGet(keyEventStructStruct, recv.native, "string")
	value := argValue.String(false)
	return value
}

// Keycode returns the C field 'keycode'.
func (recv *KeyEventStruct) Keycode() uint16 {
	argValue := gi.FieldGet(keyEventStructStruct, recv.native, "keycode")
	value := argValue.Uint16()
	return value
}

// Timestamp returns the C field 'timestamp'.
func (recv *KeyEventStruct) Timestamp() uint32 {
	argValue := gi.FieldGet(keyEventStructStruct, recv.native, "timestamp")
	value := argValue.Uint32()
	return value
}

// KeyEventStructStruct creates an uninitialised KeyEventStruct.
func KeyEventStructStruct() *KeyEventStruct {
	err := keyEventStructStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &KeyEventStruct{native: keyEventStructStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeKeyEventStruct)
	return structGo
}
func finalizeKeyEventStruct(obj *KeyEventStruct) {
	keyEventStructStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'threads_enter' : for field getter : missing Type

// UNSUPPORTED : C value 'threads_leave' : for field getter : missing Type

// UNSUPPORTED : C value 'vfuncs' : for field getter : missing Type

// MiscClassStruct creates an uninitialised MiscClass.
func MiscClassStruct() *MiscClass {
	err := miscClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MiscClass{native: miscClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMiscClass)
	return structGo
}
func finalizeMiscClass(obj *MiscClass) {
	miscClassStruct.Free(obj.native)
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
	native uintptr
}

// ParentClass returns the C field 'parent_class'.
func (recv *NoOpObjectClass) ParentClass() *ObjectClass {
	argValue := gi.FieldGet(noOpObjectClassStruct, recv.native, "parent_class")
	value := &ObjectClass{native: argValue.Pointer()}
	return value
}

// NoOpObjectClassStruct creates an uninitialised NoOpObjectClass.
func NoOpObjectClassStruct() *NoOpObjectClass {
	err := noOpObjectClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NoOpObjectClass{native: noOpObjectClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNoOpObjectClass)
	return structGo
}
func finalizeNoOpObjectClass(obj *NoOpObjectClass) {
	noOpObjectClassStruct.Free(obj.native)
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
	native uintptr
}

// ParentClass returns the C field 'parent_class'.
func (recv *NoOpObjectFactoryClass) ParentClass() *ObjectFactoryClass {
	argValue := gi.FieldGet(noOpObjectFactoryClassStruct, recv.native, "parent_class")
	value := &ObjectFactoryClass{native: argValue.Pointer()}
	return value
}

// NoOpObjectFactoryClassStruct creates an uninitialised NoOpObjectFactoryClass.
func NoOpObjectFactoryClassStruct() *NoOpObjectFactoryClass {
	err := noOpObjectFactoryClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NoOpObjectFactoryClass{native: noOpObjectFactoryClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNoOpObjectFactoryClass)
	return structGo
}
func finalizeNoOpObjectFactoryClass(obj *NoOpObjectFactoryClass) {
	noOpObjectFactoryClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_parent' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_children' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_child' : for field getter : missing Type

// UNSUPPORTED : C value 'get_index_in_parent' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_relation_set' : for field getter : missing Type

// UNSUPPORTED : C value 'get_role' : for field getter : missing Type

// UNSUPPORTED : C value 'get_layer' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mdi_zorder' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_state_set' : for field getter : missing Type

// UNSUPPORTED : C value 'set_name' : for field getter : missing Type

// UNSUPPORTED : C value 'set_description' : for field getter : missing Type

// UNSUPPORTED : C value 'set_parent' : for field getter : missing Type

// UNSUPPORTED : C value 'set_role' : for field getter : missing Type

// UNSUPPORTED : C value 'connect_property_change_handler' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_property_change_handler' : for field getter : missing Type

// UNSUPPORTED : C value 'initialize' : for field getter : missing Type

// UNSUPPORTED : C value 'children_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'focus_event' : for field getter : missing Type

// UNSUPPORTED : C value 'property_change' : for field getter : missing Type

// UNSUPPORTED : C value 'state_change' : for field getter : missing Type

// UNSUPPORTED : C value 'visible_data_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'active_descendant_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'get_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_object_locale' : for field getter : missing Type

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// ObjectClassStruct creates an uninitialised ObjectClass.
func ObjectClassStruct() *ObjectClass {
	err := objectClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ObjectClass{native: objectClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeObjectClass)
	return structGo
}
func finalizeObjectClass(obj *ObjectClass) {
	objectClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'create_accessible' : for field getter : missing Type

// UNSUPPORTED : C value 'invalidate' : for field getter : missing Type

// UNSUPPORTED : C value 'get_accessible_type' : for field getter : missing Type

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field getter : no Go type for 'Function'

// ObjectFactoryClassStruct creates an uninitialised ObjectFactoryClass.
func ObjectFactoryClassStruct() *ObjectFactoryClass {
	err := objectFactoryClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ObjectFactoryClass{native: objectFactoryClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeObjectFactoryClass)
	return structGo
}
func finalizeObjectFactoryClass(obj *ObjectFactoryClass) {
	objectFactoryClassStruct.Free(obj.native)
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
	native uintptr
}

// ParentClass returns the C field 'parent_class'.
func (recv *PlugClass) ParentClass() *ObjectClass {
	argValue := gi.FieldGet(plugClassStruct, recv.native, "parent_class")
	value := &ObjectClass{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'get_object_id' : for field getter : missing Type

// PlugClassStruct creates an uninitialised PlugClass.
func PlugClassStruct() *PlugClass {
	err := plugClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PlugClass{native: plugClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePlugClass)
	return structGo
}
func finalizePlugClass(obj *PlugClass) {
	plugClassStruct.Free(obj.native)
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
	native uintptr
}

// PropertyName returns the C field 'property_name'.
func (recv *PropertyValues) PropertyName() string {
	argValue := gi.FieldGet(propertyValuesStruct, recv.native, "property_name")
	value := argValue.String(false)
	return value
}

// UNSUPPORTED : C value 'old_value' : for field getter : no Go type for 'GObject.Value'

// UNSUPPORTED : C value 'new_value' : for field getter : no Go type for 'GObject.Value'

// PropertyValuesStruct creates an uninitialised PropertyValues.
func PropertyValuesStruct() *PropertyValues {
	err := propertyValuesStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PropertyValues{native: propertyValuesStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePropertyValues)
	return structGo
}
func finalizePropertyValues(obj *PropertyValues) {
	propertyValuesStruct.Free(obj.native)
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
	native uintptr
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

	retGo := &Range{native: ret.Pointer()}

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
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := rangeCopyFunction_Set()
	if err == nil {
		ret = rangeCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Range{native: ret.Pointer()}

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	inArgs[0].SetPointer(recv.native)

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
	native uintptr
}

// X returns the C field 'x'.
func (recv *Rectangle) X() int32 {
	argValue := gi.FieldGet(rectangleStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// Y returns the C field 'y'.
func (recv *Rectangle) Y() int32 {
	argValue := gi.FieldGet(rectangleStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// Width returns the C field 'width'.
func (recv *Rectangle) Width() int32 {
	argValue := gi.FieldGet(rectangleStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// Height returns the C field 'height'.
func (recv *Rectangle) Height() int32 {
	argValue := gi.FieldGet(rectangleStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// RectangleStruct creates an uninitialised Rectangle.
func RectangleStruct() *Rectangle {
	err := rectangleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Rectangle{native: rectangleStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRectangle)
	return structGo
}
func finalizeRectangle(obj *Rectangle) {
	rectangleStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// RegistryClassStruct creates an uninitialised RegistryClass.
func RegistryClassStruct() *RegistryClass {
	err := registryClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RegistryClass{native: registryClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRegistryClass)
	return structGo
}
func finalizeRegistryClass(obj *RegistryClass) {
	registryClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// RelationClassStruct creates an uninitialised RelationClass.
func RelationClassStruct() *RelationClass {
	err := relationClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RelationClass{native: relationClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRelationClass)
	return structGo
}
func finalizeRelationClass(obj *RelationClass) {
	relationClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field getter : no Go type for 'Function'

// RelationSetClassStruct creates an uninitialised RelationSetClass.
func RelationSetClassStruct() *RelationSetClass {
	err := relationSetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RelationSetClass{native: relationSetClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRelationSetClass)
	return structGo
}
func finalizeRelationSetClass(obj *RelationSetClass) {
	relationSetClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'add_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'clear_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'get_selection_count' : for field getter : missing Type

// UNSUPPORTED : C value 'is_child_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'select_all_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_changed' : for field getter : missing Type

// SelectionIfaceStruct creates an uninitialised SelectionIface.
func SelectionIfaceStruct() *SelectionIface {
	err := selectionIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SelectionIface{native: selectionIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSelectionIface)
	return structGo
}
func finalizeSelectionIface(obj *SelectionIface) {
	selectionIfaceStruct.Free(obj.native)
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
	native uintptr
}

// ParentClass returns the C field 'parent_class'.
func (recv *SocketClass) ParentClass() *ObjectClass {
	argValue := gi.FieldGet(socketClassStruct, recv.native, "parent_class")
	value := &ObjectClass{native: argValue.Pointer()}
	return value
}

// UNSUPPORTED : C value 'embed' : for field getter : missing Type

// SocketClassStruct creates an uninitialised SocketClass.
func SocketClassStruct() *SocketClass {
	err := socketClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketClass{native: socketClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSocketClass)
	return structGo
}
func finalizeSocketClass(obj *SocketClass) {
	socketClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// StateSetClassStruct creates an uninitialised StateSetClass.
func StateSetClassStruct() *StateSetClass {
	err := stateSetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StateSetClass{native: stateSetClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStateSetClass)
	return structGo
}
func finalizeStateSetClass(obj *StateSetClass) {
	stateSetClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_n_mime_types' : for field getter : missing Type

// UNSUPPORTED : C value 'get_mime_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_stream' : for field getter : missing Type

// UNSUPPORTED : C value 'get_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'pad1' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad2' : for field getter : no Go type for 'Function'

// UNSUPPORTED : C value 'pad3' : for field getter : no Go type for 'Function'

// StreamableContentIfaceStruct creates an uninitialised StreamableContentIface.
func StreamableContentIfaceStruct() *StreamableContentIface {
	err := streamableContentIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StreamableContentIface{native: streamableContentIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStreamableContentIface)
	return structGo
}
func finalizeStreamableContentIface(obj *StreamableContentIface) {
	streamableContentIfaceStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'get_column_span' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_header_cells' : for field getter : missing Type

// UNSUPPORTED : C value 'get_position' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_span' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_header_cells' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_column_span' : for field getter : missing Type

// UNSUPPORTED : C value 'get_table' : for field getter : missing Type

// TableCellIfaceStruct creates an uninitialised TableCellIface.
func TableCellIfaceStruct() *TableCellIface {
	err := tableCellIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TableCellIface{native: tableCellIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTableCellIface)
	return structGo
}
func finalizeTableCellIface(obj *TableCellIface) {
	tableCellIfaceStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'ref_at' : for field getter : missing Type

// UNSUPPORTED : C value 'get_index_at' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_at_index' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_at_index' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_columns' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_rows' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_extent_at' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_extent_at' : for field getter : missing Type

// UNSUPPORTED : C value 'get_caption' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_header' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_description' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_header' : for field getter : missing Type

// UNSUPPORTED : C value 'get_summary' : for field getter : missing Type

// UNSUPPORTED : C value 'set_caption' : for field getter : missing Type

// UNSUPPORTED : C value 'set_column_description' : for field getter : missing Type

// UNSUPPORTED : C value 'set_column_header' : for field getter : missing Type

// UNSUPPORTED : C value 'set_row_description' : for field getter : missing Type

// UNSUPPORTED : C value 'set_row_header' : for field getter : missing Type

// UNSUPPORTED : C value 'set_summary' : for field getter : missing Type

// UNSUPPORTED : C value 'get_selected_columns' : for field getter : missing Type

// UNSUPPORTED : C value 'get_selected_rows' : for field getter : missing Type

// UNSUPPORTED : C value 'is_column_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'is_row_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'is_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'add_row_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_row_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'add_column_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_column_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'row_inserted' : for field getter : missing Type

// UNSUPPORTED : C value 'column_inserted' : for field getter : missing Type

// UNSUPPORTED : C value 'row_deleted' : for field getter : missing Type

// UNSUPPORTED : C value 'column_deleted' : for field getter : missing Type

// UNSUPPORTED : C value 'row_reordered' : for field getter : missing Type

// UNSUPPORTED : C value 'column_reordered' : for field getter : missing Type

// UNSUPPORTED : C value 'model_changed' : for field getter : missing Type

// TableIfaceStruct creates an uninitialised TableIface.
func TableIfaceStruct() *TableIface {
	err := tableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TableIface{native: tableIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTableIface)
	return structGo
}
func finalizeTableIface(obj *TableIface) {
	tableIfaceStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_text' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_after_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_at_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_character_at_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_before_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_caret_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'get_run_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_default_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'get_character_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'get_character_count' : for field getter : missing Type

// UNSUPPORTED : C value 'get_offset_at_point' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_selections' : for field getter : missing Type

// UNSUPPORTED : C value 'get_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'add_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'set_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'set_caret_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'text_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'text_caret_moved' : for field getter : missing Type

// UNSUPPORTED : C value 'text_selection_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'text_attributes_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'get_range_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'get_bounded_ranges' : for field getter : missing Type

// UNSUPPORTED : C value 'get_string_at_offset' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_substring_to' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_substring_to_point' : for field getter : missing Type

// TextIfaceStruct creates an uninitialised TextIface.
func TextIfaceStruct() *TextIface {
	err := textIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextIface{native: textIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextIface)
	return structGo
}
func finalizeTextIface(obj *TextIface) {
	textIfaceStruct.Free(obj.native)
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
	native uintptr
}

// Bounds returns the C field 'bounds'.
func (recv *TextRange) Bounds() *TextRectangle {
	argValue := gi.FieldGet(textRangeStruct, recv.native, "bounds")
	value := &TextRectangle{native: argValue.Pointer()}
	return value
}

// StartOffset returns the C field 'start_offset'.
func (recv *TextRange) StartOffset() int32 {
	argValue := gi.FieldGet(textRangeStruct, recv.native, "start_offset")
	value := argValue.Int32()
	return value
}

// EndOffset returns the C field 'end_offset'.
func (recv *TextRange) EndOffset() int32 {
	argValue := gi.FieldGet(textRangeStruct, recv.native, "end_offset")
	value := argValue.Int32()
	return value
}

// Content returns the C field 'content'.
func (recv *TextRange) Content() string {
	argValue := gi.FieldGet(textRangeStruct, recv.native, "content")
	value := argValue.String(false)
	return value
}

// TextRangeStruct creates an uninitialised TextRange.
func TextRangeStruct() *TextRange {
	err := textRangeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextRange{native: textRangeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextRange)
	return structGo
}
func finalizeTextRange(obj *TextRange) {
	textRangeStruct.Free(obj.native)
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
	native uintptr
}

// X returns the C field 'x'.
func (recv *TextRectangle) X() int32 {
	argValue := gi.FieldGet(textRectangleStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// Y returns the C field 'y'.
func (recv *TextRectangle) Y() int32 {
	argValue := gi.FieldGet(textRectangleStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// Width returns the C field 'width'.
func (recv *TextRectangle) Width() int32 {
	argValue := gi.FieldGet(textRectangleStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// Height returns the C field 'height'.
func (recv *TextRectangle) Height() int32 {
	argValue := gi.FieldGet(textRectangleStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// TextRectangleStruct creates an uninitialised TextRectangle.
func TextRectangleStruct() *TextRectangle {
	err := textRectangleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextRectangle{native: textRectangleStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextRectangle)
	return structGo
}
func finalizeTextRectangle(obj *TextRectangle) {
	textRectangleStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'add_global_event_listener' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_global_event_listener' : for field getter : missing Type

// UNSUPPORTED : C value 'add_key_event_listener' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_key_event_listener' : for field getter : missing Type

// UNSUPPORTED : C value 'get_root' : for field getter : missing Type

// UNSUPPORTED : C value 'get_toolkit_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_toolkit_version' : for field getter : missing Type

// UtilClassStruct creates an uninitialised UtilClass.
func UtilClassStruct() *UtilClass {
	err := utilClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UtilClass{native: utilClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeUtilClass)
	return structGo
}
func finalizeUtilClass(obj *UtilClass) {
	utilClassStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_current_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_maximum_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_minimum_value' : for field getter : missing Type

// UNSUPPORTED : C value 'set_current_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_minimum_increment' : for field getter : missing Type

// UNSUPPORTED : C value 'get_value_and_text' : for field getter : missing Type

// UNSUPPORTED : C value 'get_range' : for field getter : missing Type

// UNSUPPORTED : C value 'get_increment' : for field getter : missing Type

// UNSUPPORTED : C value 'get_sub_ranges' : for field getter : missing Type

// UNSUPPORTED : C value 'set_value' : for field getter : missing Type

// ValueIfaceStruct creates an uninitialised ValueIface.
func ValueIfaceStruct() *ValueIface {
	err := valueIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ValueIface{native: valueIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeValueIface)
	return structGo
}
func finalizeValueIface(obj *ValueIface) {
	valueIfaceStruct.Free(obj.native)
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
	native uintptr
}

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// WindowIfaceStruct creates an uninitialised WindowIface.
func WindowIfaceStruct() *WindowIface {
	err := windowIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowIface{native: windowIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowIface)
	return structGo
}
func finalizeWindowIface(obj *WindowIface) {
	windowIfaceStruct.Free(obj.native)
}
