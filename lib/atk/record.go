// This is a generated file - DO NOT EDIT

package atk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <stdlib.h>
import "C"

/*

C type

AtkActionIface
*/
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

// AtkAttribute is a string name/value pair representing a generic
// attribute. This can be used to expose additional information from
// an accessible object as a whole (see atk_object_get_attributes())
// or an document (see atk_document_get_attributes()). In the case of
// text attributes (see atk_text_get_default_attributes()),
// #AtkTextAttribute enum defines all the possible text attribute
// names. You can use atk_text_attribute_get_name() to get the string
// name from the enum value. See also atk_text_attribute_for_name()
// and atk_text_attribute_get_value() for more information.
//
// A string name/value pair representing a generic attribute.
/*

C type

AtkAttribute
*/
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

/*

C type

AtkComponentIface
*/
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

/*

C type

AtkDocumentIface
*/
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

/*

C type

AtkEditableTextIface
*/
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

/*

C type

AtkGObjectAccessibleClass
*/
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

/*

C type

AtkHyperlinkClass
*/
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

/*

C type

AtkHyperlinkImplIface
*/
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

/*

C type

AtkHypertextIface
*/
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

/*

C type

AtkImageIface
*/
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

/*

C type

AtkImplementor
*/
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

// Gets a reference to an object's #AtkObject implementation, if
// the object implements #AtkObjectIface
/*

C function

atk_implementor_ref_accessible
*/
func (recv *Implementor) RefAccessible() *Object {
	retC := C.atk_implementor_ref_accessible((*C.AtkImplementor)(recv.native))
	retGo := ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Encapsulates information about a key event.
/*

C type

AtkKeyEventStruct
*/
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

// Usage of AtkMisc is deprecated since 2.12 and heavily discouraged.
/*

C type

AtkMiscClass
*/
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

/*

C type

AtkNoOpObjectClass
*/
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

/*

C type

AtkNoOpObjectFactoryClass
*/
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

/*

C type

AtkObjectClass
*/
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

/*

C type

AtkObjectFactoryClass
*/
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

/*

C type

AtkPlugClass
*/
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

// Note: @old_value field of #AtkPropertyValues will not contain a
// valid value. This is a field defined with the purpose of contain
// the previous value of the property, but is not used anymore.
/*

C type

AtkPropertyValues
*/
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

// #AtkRange are used on #AtkValue, in order to represent the full
// range of a given component (for example an slider or a range
// control), or to define each individual subrange this full range is
// splitted if available. See #AtkValue documentation for further
// details.
/*

C type

AtkRange
*/
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

// A data structure for holding a rectangle. Those coordinates are
// relative to the component top-level parent.
/*

C type

AtkRectangle
*/
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

/*

C type

AtkRegistryClass
*/
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

/*

C type

AtkRelationClass
*/
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

/*

C type

AtkRelationSetClass
*/
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

/*

C type

AtkSelectionIface
*/
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

/*

C type

AtkSocketClass
*/
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

/*

C type

AtkStateSetClass
*/
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

/*

C type

AtkStreamableContentIface
*/
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

/*

C type

AtkTableCellIface
*/
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

/*

C type

AtkTableIface
*/
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

/*

C type

AtkTextIface
*/
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

// A structure used to describe a text range.
/*

C type

AtkTextRange
*/
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

// A structure used to store a rectangle used by AtkText.
/*

C type

AtkTextRectangle
*/
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

/*

C type

AtkUtilClass
*/
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

/*

C type

AtkValueIface
*/
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

/*

C type

AtkWindowIface
*/
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
