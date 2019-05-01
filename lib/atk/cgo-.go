// This is a generated file - DO NOT EDIT

package atk

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

// GObjectAccessible is a wrapper around the C record AtkGObjectAccessible.
type GObjectAccessible struct {
	native *C.AtkGObjectAccessible
	// parent : record
}

func GObjectAccessibleNewFromC(u unsafe.Pointer) *GObjectAccessible {
	c := (*C.AtkGObjectAccessible)(u)
	if c == nil {
		return nil
	}

	g := &GObjectAccessible{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GObjectAccessible) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GObjectAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Hyperlink is a wrapper around the C record AtkHyperlink.
type Hyperlink struct {
	native *C.AtkHyperlink
	// parent : record
}

func HyperlinkNewFromC(u unsafe.Pointer) *Hyperlink {
	c := (*C.AtkHyperlink)(u)
	if c == nil {
		return nil
	}

	g := &Hyperlink{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Hyperlink) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Hyperlink) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Misc is a wrapper around the C record AtkMisc.
type Misc struct {
	native *C.AtkMisc
	// parent : record
}

func MiscNewFromC(u unsafe.Pointer) *Misc {
	c := (*C.AtkMisc)(u)
	if c == nil {
		return nil
	}

	g := &Misc{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Misc) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Misc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NoOpObject is a wrapper around the C record AtkNoOpObject.
type NoOpObject struct {
	native *C.AtkNoOpObject
	// parent : record
}

func NoOpObjectNewFromC(u unsafe.Pointer) *NoOpObject {
	c := (*C.AtkNoOpObject)(u)
	if c == nil {
		return nil
	}

	g := &NoOpObject{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NoOpObject) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NoOpObject) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NoOpObjectFactory is a wrapper around the C record AtkNoOpObjectFactory.
type NoOpObjectFactory struct {
	native *C.AtkNoOpObjectFactory
	// parent : record
}

func NoOpObjectFactoryNewFromC(u unsafe.Pointer) *NoOpObjectFactory {
	c := (*C.AtkNoOpObjectFactory)(u)
	if c == nil {
		return nil
	}

	g := &NoOpObjectFactory{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NoOpObjectFactory) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NoOpObjectFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object is a wrapper around the C record AtkObject.
type Object struct {
	native *C.AtkObject
	// parent : record
	Description string
	Name        string
	// accessible_parent : record
	Role Role
	// relation_set : record
	Layer Layer
}

func ObjectNewFromC(u unsafe.Pointer) *Object {
	c := (*C.AtkObject)(u)
	if c == nil {
		return nil
	}

	g := &Object{
		Description: C.GoString(c.description),
		Layer:       (Layer)(c.layer),
		Name:        C.GoString(c.name),
		Role:        (Role)(c.role),
		native:      c,
	}

	return g
}

func (recv *Object) ToC() unsafe.Pointer {
	recv.native.description =
		C.CString(recv.Description)
	recv.native.name =
		C.CString(recv.Name)
	recv.native.role =
		(C.AtkRole)(recv.Role)
	recv.native.layer =
		(C.AtkLayer)(recv.Layer)

	return (unsafe.Pointer)(recv.native)
}

// ObjectFactory is a wrapper around the C record AtkObjectFactory.
type ObjectFactory struct {
	native *C.AtkObjectFactory
	// parent : record
}

func ObjectFactoryNewFromC(u unsafe.Pointer) *ObjectFactory {
	c := (*C.AtkObjectFactory)(u)
	if c == nil {
		return nil
	}

	g := &ObjectFactory{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ObjectFactory) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ObjectFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Plug is a wrapper around the C record AtkPlug.
type Plug struct {
	native *C.AtkPlug
	// parent : record
}

func PlugNewFromC(u unsafe.Pointer) *Plug {
	c := (*C.AtkPlug)(u)
	if c == nil {
		return nil
	}

	g := &Plug{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Plug) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Plug) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Registry is a wrapper around the C record AtkRegistry.
type Registry struct {
	native *C.AtkRegistry
	// parent : record
	// factory_type_registry : record
	// factory_singleton_cache : record
}

func RegistryNewFromC(u unsafe.Pointer) *Registry {
	c := (*C.AtkRegistry)(u)
	if c == nil {
		return nil
	}

	g := &Registry{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Registry) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Registry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Relation is a wrapper around the C record AtkRelation.
type Relation struct {
	native *C.AtkRelation
	// parent : record
	// no type for target
	Relationship RelationType
}

func RelationNewFromC(u unsafe.Pointer) *Relation {
	c := (*C.AtkRelation)(u)
	if c == nil {
		return nil
	}

	g := &Relation{
		Relationship: (RelationType)(c.relationship),
		native:       c,
	}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Relation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Relation) ToC() unsafe.Pointer {
	recv.native.relationship =
		(C.AtkRelationType)(recv.Relationship)

	return (unsafe.Pointer)(recv.native)
}

// RelationSet is a wrapper around the C record AtkRelationSet.
type RelationSet struct {
	native *C.AtkRelationSet
	// parent : record
	// no type for relations
}

func RelationSetNewFromC(u unsafe.Pointer) *RelationSet {
	c := (*C.AtkRelationSet)(u)
	if c == nil {
		return nil
	}

	g := &RelationSet{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *RelationSet) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *RelationSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Socket is a wrapper around the C record AtkSocket.
type Socket struct {
	native *C.AtkSocket
	// parent : record
	// Private : embedded_plug_id
}

func SocketNewFromC(u unsafe.Pointer) *Socket {
	c := (*C.AtkSocket)(u)
	if c == nil {
		return nil
	}

	g := &Socket{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Socket) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Socket) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StateSet is a wrapper around the C record AtkStateSet.
type StateSet struct {
	native *C.AtkStateSet
	// parent : record
}

func StateSetNewFromC(u unsafe.Pointer) *StateSet {
	c := (*C.AtkStateSet)(u)
	if c == nil {
		return nil
	}

	g := &StateSet{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *StateSet) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *StateSet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Util is a wrapper around the C record AtkUtil.
type Util struct {
	native *C.AtkUtil
	// parent : record
}

func UtilNewFromC(u unsafe.Pointer) *Util {
	c := (*C.AtkUtil)(u)
	if c == nil {
		return nil
	}

	g := &Util{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Util) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Util) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : atk_add_focus_tracker : unsupported parameter focus_tracker : no type generator for EventListener (AtkEventListener) for param focus_tracker

// Unsupported : atk_add_global_event_listener : unsupported parameter listener : no type generator for GObject.SignalEmissionHook (GSignalEmissionHook) for param listener

// Unsupported : atk_add_key_event_listener : unsupported parameter listener : no type generator for KeySnoopFunc (AtkKeySnoopFunc) for param listener

// Blacklisted : atk_attribute_set_free

// Unsupported : atk_focus_tracker_init : unsupported parameter init : no type generator for EventListenerInit (AtkEventListenerInit) for param init

// Action is a wrapper around the C record AtkAction.
type Action struct {
	native *C.AtkAction
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.AtkAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Component is a wrapper around the C record AtkComponent.
type Component struct {
	native *C.AtkComponent
}

func ComponentNewFromC(u unsafe.Pointer) *Component {
	c := (*C.AtkComponent)(u)
	if c == nil {
		return nil
	}

	g := &Component{native: c}

	return g
}

func (recv *Component) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Document is a wrapper around the C record AtkDocument.
type Document struct {
	native *C.AtkDocument
}

func DocumentNewFromC(u unsafe.Pointer) *Document {
	c := (*C.AtkDocument)(u)
	if c == nil {
		return nil
	}

	g := &Document{native: c}

	return g
}

func (recv *Document) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EditableText is a wrapper around the C record AtkEditableText.
type EditableText struct {
	native *C.AtkEditableText
}

func EditableTextNewFromC(u unsafe.Pointer) *EditableText {
	c := (*C.AtkEditableText)(u)
	if c == nil {
		return nil
	}

	g := &EditableText{native: c}

	return g
}

func (recv *EditableText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HyperlinkImpl is a wrapper around the C record AtkHyperlinkImpl.
type HyperlinkImpl struct {
	native *C.AtkHyperlinkImpl
}

func HyperlinkImplNewFromC(u unsafe.Pointer) *HyperlinkImpl {
	c := (*C.AtkHyperlinkImpl)(u)
	if c == nil {
		return nil
	}

	g := &HyperlinkImpl{native: c}

	return g
}

func (recv *HyperlinkImpl) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Hypertext is a wrapper around the C record AtkHypertext.
type Hypertext struct {
	native *C.AtkHypertext
}

func HypertextNewFromC(u unsafe.Pointer) *Hypertext {
	c := (*C.AtkHypertext)(u)
	if c == nil {
		return nil
	}

	g := &Hypertext{native: c}

	return g
}

func (recv *Hypertext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Image is a wrapper around the C record AtkImage.
type Image struct {
	native *C.AtkImage
}

func ImageNewFromC(u unsafe.Pointer) *Image {
	c := (*C.AtkImage)(u)
	if c == nil {
		return nil
	}

	g := &Image{native: c}

	return g
}

func (recv *Image) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImplementorIface is a wrapper around the C record AtkImplementorIface.
type ImplementorIface struct {
	native *C.AtkImplementorIface
}

func ImplementorIfaceNewFromC(u unsafe.Pointer) *ImplementorIface {
	c := (*C.AtkImplementorIface)(u)
	if c == nil {
		return nil
	}

	g := &ImplementorIface{native: c}

	return g
}

func (recv *ImplementorIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Selection is a wrapper around the C record AtkSelection.
type Selection struct {
	native *C.AtkSelection
}

func SelectionNewFromC(u unsafe.Pointer) *Selection {
	c := (*C.AtkSelection)(u)
	if c == nil {
		return nil
	}

	g := &Selection{native: c}

	return g
}

func (recv *Selection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StreamableContent is a wrapper around the C record AtkStreamableContent.
type StreamableContent struct {
	native *C.AtkStreamableContent
}

func StreamableContentNewFromC(u unsafe.Pointer) *StreamableContent {
	c := (*C.AtkStreamableContent)(u)
	if c == nil {
		return nil
	}

	g := &StreamableContent{native: c}

	return g
}

func (recv *StreamableContent) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Table is a wrapper around the C record AtkTable.
type Table struct {
	native *C.AtkTable
}

func TableNewFromC(u unsafe.Pointer) *Table {
	c := (*C.AtkTable)(u)
	if c == nil {
		return nil
	}

	g := &Table{native: c}

	return g
}

func (recv *Table) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableCell is a wrapper around the C record AtkTableCell.
type TableCell struct {
	native *C.AtkTableCell
}

func TableCellNewFromC(u unsafe.Pointer) *TableCell {
	c := (*C.AtkTableCell)(u)
	if c == nil {
		return nil
	}

	g := &TableCell{native: c}

	return g
}

func (recv *TableCell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Text is a wrapper around the C record AtkText.
type Text struct {
	native *C.AtkText
}

func TextNewFromC(u unsafe.Pointer) *Text {
	c := (*C.AtkText)(u)
	if c == nil {
		return nil
	}

	g := &Text{native: c}

	return g
}

func (recv *Text) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Value is a wrapper around the C record AtkValue.
type Value struct {
	native *C.AtkValue
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	c := (*C.AtkValue)(u)
	if c == nil {
		return nil
	}

	g := &Value{native: c}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window is a wrapper around the C record AtkWindow.
type Window struct {
	native *C.AtkWindow
}

func WindowNewFromC(u unsafe.Pointer) *Window {
	c := (*C.AtkWindow)(u)
	if c == nil {
		return nil
	}

	g := &Window{native: c}

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ActionIface is a wrapper around the C record AtkActionIface.
type ActionIface struct {
	native *C.AtkActionIface
	// parent : record
	// no type for do_action
	// no type for get_n_actions
	// no type for get_description
	// no type for get_name
	// no type for get_keybinding
	// no type for set_description
	// no type for get_localized_name
}

func ActionIfaceNewFromC(u unsafe.Pointer) *ActionIface {
	c := (*C.AtkActionIface)(u)
	if c == nil {
		return nil
	}

	g := &ActionIface{native: c}

	return g
}

func (recv *ActionIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Attribute is a wrapper around the C record AtkAttribute.
type Attribute struct {
	native *C.AtkAttribute
	Name   string
	Value  string
}

func AttributeNewFromC(u unsafe.Pointer) *Attribute {
	c := (*C.AtkAttribute)(u)
	if c == nil {
		return nil
	}

	g := &Attribute{
		Name:   C.GoString(c.name),
		Value:  C.GoString(c.value),
		native: c,
	}

	return g
}

func (recv *Attribute) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.value =
		C.CString(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// ComponentIface is a wrapper around the C record AtkComponentIface.
type ComponentIface struct {
	native *C.AtkComponentIface
	// parent : record
	// no type for add_focus_handler
	// no type for contains
	// no type for ref_accessible_at_point
	// no type for get_extents
	// no type for get_position
	// no type for get_size
	// no type for grab_focus
	// no type for remove_focus_handler
	// no type for set_extents
	// no type for set_position
	// no type for set_size
	// no type for get_layer
	// no type for get_mdi_zorder
	// no type for bounds_changed
	// no type for get_alpha
}

func ComponentIfaceNewFromC(u unsafe.Pointer) *ComponentIface {
	c := (*C.AtkComponentIface)(u)
	if c == nil {
		return nil
	}

	g := &ComponentIface{native: c}

	return g
}

func (recv *ComponentIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DocumentIface is a wrapper around the C record AtkDocumentIface.
type DocumentIface struct {
	native *C.AtkDocumentIface
	// parent : record
	// no type for get_document_type
	// no type for get_document
	// no type for get_document_locale
	// no type for get_document_attributes
	// no type for get_document_attribute_value
	// no type for set_document_attribute
	// no type for get_current_page_number
	// no type for get_page_count
}

func DocumentIfaceNewFromC(u unsafe.Pointer) *DocumentIface {
	c := (*C.AtkDocumentIface)(u)
	if c == nil {
		return nil
	}

	g := &DocumentIface{native: c}

	return g
}

func (recv *DocumentIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EditableTextIface is a wrapper around the C record AtkEditableTextIface.
type EditableTextIface struct {
	native *C.AtkEditableTextIface
	// parent_interface : record
	// no type for set_run_attributes
	// no type for set_text_contents
	// no type for insert_text
	// no type for copy_text
	// no type for cut_text
	// no type for delete_text
	// no type for paste_text
}

func EditableTextIfaceNewFromC(u unsafe.Pointer) *EditableTextIface {
	c := (*C.AtkEditableTextIface)(u)
	if c == nil {
		return nil
	}

	g := &EditableTextIface{native: c}

	return g
}

func (recv *EditableTextIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GObjectAccessibleClass is a wrapper around the C record AtkGObjectAccessibleClass.
type GObjectAccessibleClass struct {
	native *C.AtkGObjectAccessibleClass
	// parent_class : record
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
}

func GObjectAccessibleClassNewFromC(u unsafe.Pointer) *GObjectAccessibleClass {
	c := (*C.AtkGObjectAccessibleClass)(u)
	if c == nil {
		return nil
	}

	g := &GObjectAccessibleClass{native: c}

	return g
}

func (recv *GObjectAccessibleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HyperlinkClass is a wrapper around the C record AtkHyperlinkClass.
type HyperlinkClass struct {
	native *C.AtkHyperlinkClass
	// parent : record
	// no type for get_uri
	// no type for get_object
	// no type for get_end_index
	// no type for get_start_index
	// no type for is_valid
	// no type for get_n_anchors
	// no type for link_state
	// no type for is_selected_link
	// no type for link_activated
	// pad1 : no type generator for Function, AtkFunction
}

func HyperlinkClassNewFromC(u unsafe.Pointer) *HyperlinkClass {
	c := (*C.AtkHyperlinkClass)(u)
	if c == nil {
		return nil
	}

	g := &HyperlinkClass{native: c}

	return g
}

func (recv *HyperlinkClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HyperlinkImplIface is a wrapper around the C record AtkHyperlinkImplIface.
type HyperlinkImplIface struct {
	native *C.AtkHyperlinkImplIface
	// parent : record
	// no type for get_hyperlink
}

func HyperlinkImplIfaceNewFromC(u unsafe.Pointer) *HyperlinkImplIface {
	c := (*C.AtkHyperlinkImplIface)(u)
	if c == nil {
		return nil
	}

	g := &HyperlinkImplIface{native: c}

	return g
}

func (recv *HyperlinkImplIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// HypertextIface is a wrapper around the C record AtkHypertextIface.
type HypertextIface struct {
	native *C.AtkHypertextIface
	// parent : record
	// no type for get_link
	// no type for get_n_links
	// no type for get_link_index
	// no type for link_selected
}

func HypertextIfaceNewFromC(u unsafe.Pointer) *HypertextIface {
	c := (*C.AtkHypertextIface)(u)
	if c == nil {
		return nil
	}

	g := &HypertextIface{native: c}

	return g
}

func (recv *HypertextIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ImageIface is a wrapper around the C record AtkImageIface.
type ImageIface struct {
	native *C.AtkImageIface
	// parent : record
	// no type for get_image_position
	// no type for get_image_description
	// no type for get_image_size
	// no type for set_image_description
	// no type for get_image_locale
}

func ImageIfaceNewFromC(u unsafe.Pointer) *ImageIface {
	c := (*C.AtkImageIface)(u)
	if c == nil {
		return nil
	}

	g := &ImageIface{native: c}

	return g
}

func (recv *ImageIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Implementor is a wrapper around the C record AtkImplementor.
type Implementor struct {
	native *C.AtkImplementor
}

func ImplementorNewFromC(u unsafe.Pointer) *Implementor {
	c := (*C.AtkImplementor)(u)
	if c == nil {
		return nil
	}

	g := &Implementor{native: c}

	return g
}

func (recv *Implementor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// KeyEventStruct is a wrapper around the C record AtkKeyEventStruct.
type KeyEventStruct struct {
	native    *C.AtkKeyEventStruct
	Type      int32
	State     uint32
	Keyval    uint32
	Length    int32
	String    string
	Keycode   uint16
	Timestamp uint32
}

func KeyEventStructNewFromC(u unsafe.Pointer) *KeyEventStruct {
	c := (*C.AtkKeyEventStruct)(u)
	if c == nil {
		return nil
	}

	g := &KeyEventStruct{
		Keycode:   (uint16)(c.keycode),
		Keyval:    (uint32)(c.keyval),
		Length:    (int32)(c.length),
		State:     (uint32)(c.state),
		String:    C.GoString(c.string),
		Timestamp: (uint32)(c.timestamp),
		Type:      (int32)(c._type),
		native:    c,
	}

	return g
}

func (recv *KeyEventStruct) ToC() unsafe.Pointer {
	recv.native._type =
		(C.gint)(recv.Type)
	recv.native.state =
		(C.guint)(recv.State)
	recv.native.keyval =
		(C.guint)(recv.Keyval)
	recv.native.length =
		(C.gint)(recv.Length)
	recv.native.string =
		C.CString(recv.String)
	recv.native.keycode =
		(C.guint16)(recv.Keycode)
	recv.native.timestamp =
		(C.guint32)(recv.Timestamp)

	return (unsafe.Pointer)(recv.native)
}

// MiscClass is a wrapper around the C record AtkMiscClass.
type MiscClass struct {
	native *C.AtkMiscClass
	// parent : record
	// no type for threads_enter
	// no type for threads_leave
	// no type for vfuncs
}

func MiscClassNewFromC(u unsafe.Pointer) *MiscClass {
	c := (*C.AtkMiscClass)(u)
	if c == nil {
		return nil
	}

	g := &MiscClass{native: c}

	return g
}

func (recv *MiscClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NoOpObjectClass is a wrapper around the C record AtkNoOpObjectClass.
type NoOpObjectClass struct {
	native *C.AtkNoOpObjectClass
	// parent_class : record
}

func NoOpObjectClassNewFromC(u unsafe.Pointer) *NoOpObjectClass {
	c := (*C.AtkNoOpObjectClass)(u)
	if c == nil {
		return nil
	}

	g := &NoOpObjectClass{native: c}

	return g
}

func (recv *NoOpObjectClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NoOpObjectFactoryClass is a wrapper around the C record AtkNoOpObjectFactoryClass.
type NoOpObjectFactoryClass struct {
	native *C.AtkNoOpObjectFactoryClass
	// parent_class : record
}

func NoOpObjectFactoryClassNewFromC(u unsafe.Pointer) *NoOpObjectFactoryClass {
	c := (*C.AtkNoOpObjectFactoryClass)(u)
	if c == nil {
		return nil
	}

	g := &NoOpObjectFactoryClass{native: c}

	return g
}

func (recv *NoOpObjectFactoryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ObjectClass is a wrapper around the C record AtkObjectClass.
type ObjectClass struct {
	native *C.AtkObjectClass
	// parent : record
	// no type for get_name
	// no type for get_description
	// no type for get_parent
	// no type for get_n_children
	// no type for ref_child
	// no type for get_index_in_parent
	// no type for ref_relation_set
	// no type for get_role
	// no type for get_layer
	// no type for get_mdi_zorder
	// no type for ref_state_set
	// no type for set_name
	// no type for set_description
	// no type for set_parent
	// no type for set_role
	// no type for connect_property_change_handler
	// no type for remove_property_change_handler
	// no type for initialize
	// no type for children_changed
	// no type for focus_event
	// no type for property_change
	// no type for state_change
	// no type for visible_data_changed
	// no type for active_descendant_changed
	// no type for get_attributes
	// no type for get_object_locale
	// pad1 : no type generator for Function, AtkFunction
}

func ObjectClassNewFromC(u unsafe.Pointer) *ObjectClass {
	c := (*C.AtkObjectClass)(u)
	if c == nil {
		return nil
	}

	g := &ObjectClass{native: c}

	return g
}

func (recv *ObjectClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ObjectFactoryClass is a wrapper around the C record AtkObjectFactoryClass.
type ObjectFactoryClass struct {
	native *C.AtkObjectFactoryClass
	// parent_class : record
	// no type for create_accessible
	// no type for invalidate
	// no type for get_accessible_type
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
}

func ObjectFactoryClassNewFromC(u unsafe.Pointer) *ObjectFactoryClass {
	c := (*C.AtkObjectFactoryClass)(u)
	if c == nil {
		return nil
	}

	g := &ObjectFactoryClass{native: c}

	return g
}

func (recv *ObjectFactoryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PlugClass is a wrapper around the C record AtkPlugClass.
type PlugClass struct {
	native *C.AtkPlugClass
	// parent_class : record
	// no type for get_object_id
}

func PlugClassNewFromC(u unsafe.Pointer) *PlugClass {
	c := (*C.AtkPlugClass)(u)
	if c == nil {
		return nil
	}

	g := &PlugClass{native: c}

	return g
}

func (recv *PlugClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// PropertyValues is a wrapper around the C record AtkPropertyValues.
type PropertyValues struct {
	native       *C.AtkPropertyValues
	PropertyName string
	// old_value : record
	// new_value : record
}

func PropertyValuesNewFromC(u unsafe.Pointer) *PropertyValues {
	c := (*C.AtkPropertyValues)(u)
	if c == nil {
		return nil
	}

	g := &PropertyValues{
		PropertyName: C.GoString(c.property_name),
		native:       c,
	}

	return g
}

func (recv *PropertyValues) ToC() unsafe.Pointer {
	recv.native.property_name =
		C.CString(recv.PropertyName)

	return (unsafe.Pointer)(recv.native)
}

// Range is a wrapper around the C record AtkRange.
type Range struct {
	native *C.AtkRange
}

func RangeNewFromC(u unsafe.Pointer) *Range {
	c := (*C.AtkRange)(u)
	if c == nil {
		return nil
	}

	g := &Range{native: c}

	return g
}

func (recv *Range) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Rectangle is a wrapper around the C record AtkRectangle.
type Rectangle struct {
	native *C.AtkRectangle
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func RectangleNewFromC(u unsafe.Pointer) *Rectangle {
	c := (*C.AtkRectangle)(u)
	if c == nil {
		return nil
	}

	g := &Rectangle{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *Rectangle) ToC() unsafe.Pointer {
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// RegistryClass is a wrapper around the C record AtkRegistryClass.
type RegistryClass struct {
	native *C.AtkRegistryClass
	// parent_class : record
}

func RegistryClassNewFromC(u unsafe.Pointer) *RegistryClass {
	c := (*C.AtkRegistryClass)(u)
	if c == nil {
		return nil
	}

	g := &RegistryClass{native: c}

	return g
}

func (recv *RegistryClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RelationClass is a wrapper around the C record AtkRelationClass.
type RelationClass struct {
	native *C.AtkRelationClass
	// parent : record
}

func RelationClassNewFromC(u unsafe.Pointer) *RelationClass {
	c := (*C.AtkRelationClass)(u)
	if c == nil {
		return nil
	}

	g := &RelationClass{native: c}

	return g
}

func (recv *RelationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RelationSetClass is a wrapper around the C record AtkRelationSetClass.
type RelationSetClass struct {
	native *C.AtkRelationSetClass
	// parent : record
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
}

func RelationSetClassNewFromC(u unsafe.Pointer) *RelationSetClass {
	c := (*C.AtkRelationSetClass)(u)
	if c == nil {
		return nil
	}

	g := &RelationSetClass{native: c}

	return g
}

func (recv *RelationSetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SelectionIface is a wrapper around the C record AtkSelectionIface.
type SelectionIface struct {
	native *C.AtkSelectionIface
	// parent : record
	// no type for add_selection
	// no type for clear_selection
	// no type for ref_selection
	// no type for get_selection_count
	// no type for is_child_selected
	// no type for remove_selection
	// no type for select_all_selection
	// no type for selection_changed
}

func SelectionIfaceNewFromC(u unsafe.Pointer) *SelectionIface {
	c := (*C.AtkSelectionIface)(u)
	if c == nil {
		return nil
	}

	g := &SelectionIface{native: c}

	return g
}

func (recv *SelectionIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SocketClass is a wrapper around the C record AtkSocketClass.
type SocketClass struct {
	native *C.AtkSocketClass
	// parent_class : record
	// no type for embed
}

func SocketClassNewFromC(u unsafe.Pointer) *SocketClass {
	c := (*C.AtkSocketClass)(u)
	if c == nil {
		return nil
	}

	g := &SocketClass{native: c}

	return g
}

func (recv *SocketClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StateSetClass is a wrapper around the C record AtkStateSetClass.
type StateSetClass struct {
	native *C.AtkStateSetClass
	// parent : record
}

func StateSetClassNewFromC(u unsafe.Pointer) *StateSetClass {
	c := (*C.AtkStateSetClass)(u)
	if c == nil {
		return nil
	}

	g := &StateSetClass{native: c}

	return g
}

func (recv *StateSetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// StreamableContentIface is a wrapper around the C record AtkStreamableContentIface.
type StreamableContentIface struct {
	native *C.AtkStreamableContentIface
	// parent : record
	// no type for get_n_mime_types
	// no type for get_mime_type
	// no type for get_stream
	// no type for get_uri
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
	// pad3 : no type generator for Function, AtkFunction
}

func StreamableContentIfaceNewFromC(u unsafe.Pointer) *StreamableContentIface {
	c := (*C.AtkStreamableContentIface)(u)
	if c == nil {
		return nil
	}

	g := &StreamableContentIface{native: c}

	return g
}

func (recv *StreamableContentIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableCellIface is a wrapper around the C record AtkTableCellIface.
type TableCellIface struct {
	native *C.AtkTableCellIface
	// parent : record
	// no type for get_column_span
	// no type for get_column_header_cells
	// no type for get_position
	// no type for get_row_span
	// no type for get_row_header_cells
	// no type for get_row_column_span
	// no type for get_table
}

func TableCellIfaceNewFromC(u unsafe.Pointer) *TableCellIface {
	c := (*C.AtkTableCellIface)(u)
	if c == nil {
		return nil
	}

	g := &TableCellIface{native: c}

	return g
}

func (recv *TableCellIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TableIface is a wrapper around the C record AtkTableIface.
type TableIface struct {
	native *C.AtkTableIface
	// parent : record
	// no type for ref_at
	// no type for get_index_at
	// no type for get_column_at_index
	// no type for get_row_at_index
	// no type for get_n_columns
	// no type for get_n_rows
	// no type for get_column_extent_at
	// no type for get_row_extent_at
	// no type for get_caption
	// no type for get_column_description
	// no type for get_column_header
	// no type for get_row_description
	// no type for get_row_header
	// no type for get_summary
	// no type for set_caption
	// no type for set_column_description
	// no type for set_column_header
	// no type for set_row_description
	// no type for set_row_header
	// no type for set_summary
	// no type for get_selected_columns
	// no type for get_selected_rows
	// no type for is_column_selected
	// no type for is_row_selected
	// no type for is_selected
	// no type for add_row_selection
	// no type for remove_row_selection
	// no type for add_column_selection
	// no type for remove_column_selection
	// no type for row_inserted
	// no type for column_inserted
	// no type for row_deleted
	// no type for column_deleted
	// no type for row_reordered
	// no type for column_reordered
	// no type for model_changed
}

func TableIfaceNewFromC(u unsafe.Pointer) *TableIface {
	c := (*C.AtkTableIface)(u)
	if c == nil {
		return nil
	}

	g := &TableIface{native: c}

	return g
}

func (recv *TableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextIface is a wrapper around the C record AtkTextIface.
type TextIface struct {
	native *C.AtkTextIface
	// parent : record
	// no type for get_text
	// no type for get_text_after_offset
	// no type for get_text_at_offset
	// no type for get_character_at_offset
	// no type for get_text_before_offset
	// no type for get_caret_offset
	// no type for get_run_attributes
	// no type for get_default_attributes
	// no type for get_character_extents
	// no type for get_character_count
	// no type for get_offset_at_point
	// no type for get_n_selections
	// no type for get_selection
	// no type for add_selection
	// no type for remove_selection
	// no type for set_selection
	// no type for set_caret_offset
	// no type for text_changed
	// no type for text_caret_moved
	// no type for text_selection_changed
	// no type for text_attributes_changed
	// no type for get_range_extents
	// no type for get_bounded_ranges
	// no type for get_string_at_offset
}

func TextIfaceNewFromC(u unsafe.Pointer) *TextIface {
	c := (*C.AtkTextIface)(u)
	if c == nil {
		return nil
	}

	g := &TextIface{native: c}

	return g
}

func (recv *TextIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TextRange is a wrapper around the C record AtkTextRange.
type TextRange struct {
	native *C.AtkTextRange
	// bounds : record
	StartOffset int32
	EndOffset   int32
	Content     string
}

func TextRangeNewFromC(u unsafe.Pointer) *TextRange {
	c := (*C.AtkTextRange)(u)
	if c == nil {
		return nil
	}

	g := &TextRange{
		Content:     C.GoString(c.content),
		EndOffset:   (int32)(c.end_offset),
		StartOffset: (int32)(c.start_offset),
		native:      c,
	}

	return g
}

func (recv *TextRange) ToC() unsafe.Pointer {
	recv.native.start_offset =
		(C.gint)(recv.StartOffset)
	recv.native.end_offset =
		(C.gint)(recv.EndOffset)
	recv.native.content =
		C.CString(recv.Content)

	return (unsafe.Pointer)(recv.native)
}

// TextRectangle is a wrapper around the C record AtkTextRectangle.
type TextRectangle struct {
	native *C.AtkTextRectangle
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func TextRectangleNewFromC(u unsafe.Pointer) *TextRectangle {
	c := (*C.AtkTextRectangle)(u)
	if c == nil {
		return nil
	}

	g := &TextRectangle{
		Height: (int32)(c.height),
		Width:  (int32)(c.width),
		X:      (int32)(c.x),
		Y:      (int32)(c.y),
		native: c,
	}

	return g
}

func (recv *TextRectangle) ToC() unsafe.Pointer {
	recv.native.x =
		(C.gint)(recv.X)
	recv.native.y =
		(C.gint)(recv.Y)
	recv.native.width =
		(C.gint)(recv.Width)
	recv.native.height =
		(C.gint)(recv.Height)

	return (unsafe.Pointer)(recv.native)
}

// UtilClass is a wrapper around the C record AtkUtilClass.
type UtilClass struct {
	native *C.AtkUtilClass
	// parent : record
	// no type for add_global_event_listener
	// no type for remove_global_event_listener
	// no type for add_key_event_listener
	// no type for remove_key_event_listener
	// no type for get_root
	// no type for get_toolkit_name
	// no type for get_toolkit_version
}

func UtilClassNewFromC(u unsafe.Pointer) *UtilClass {
	c := (*C.AtkUtilClass)(u)
	if c == nil {
		return nil
	}

	g := &UtilClass{native: c}

	return g
}

func (recv *UtilClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ValueIface is a wrapper around the C record AtkValueIface.
type ValueIface struct {
	native *C.AtkValueIface
	// parent : record
	// no type for get_current_value
	// no type for get_maximum_value
	// no type for get_minimum_value
	// no type for set_current_value
	// no type for get_minimum_increment
	// no type for get_value_and_text
	// no type for get_range
	// no type for get_increment
	// no type for get_sub_ranges
	// no type for set_value
}

func ValueIfaceNewFromC(u unsafe.Pointer) *ValueIface {
	c := (*C.AtkValueIface)(u)
	if c == nil {
		return nil
	}

	g := &ValueIface{native: c}

	return g
}

func (recv *ValueIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WindowIface is a wrapper around the C record AtkWindowIface.
type WindowIface struct {
	native *C.AtkWindowIface
	// parent : record
}

func WindowIfaceNewFromC(u unsafe.Pointer) *WindowIface {
	c := (*C.AtkWindowIface)(u)
	if c == nil {
		return nil
	}

	g := &WindowIface{native: c}

	return g
}

func (recv *WindowIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
