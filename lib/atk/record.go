// This is a generated file - DO NOT EDIT

package atk

import "unsafe"

// ActionIface is a wrapper around the C record AtkActionIface.
type ActionIface struct {
	native unsafe.Pointer
	// parent : record
	// no type for do_action
	// no type for get_n_actions
	// no type for get_description
	// no type for get_name
	// no type for get_keybinding
	// no type for set_description
	// no type for get_localized_name
}

// Attribute is a wrapper around the C record AtkAttribute.
type Attribute struct {
	native unsafe.Pointer
	Name   string
	Value  string
}

// ComponentIface is a wrapper around the C record AtkComponentIface.
type ComponentIface struct {
	native unsafe.Pointer
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

// DocumentIface is a wrapper around the C record AtkDocumentIface.
type DocumentIface struct {
	native unsafe.Pointer
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

// EditableTextIface is a wrapper around the C record AtkEditableTextIface.
type EditableTextIface struct {
	native unsafe.Pointer
	// parent_interface : record
	// no type for set_run_attributes
	// no type for set_text_contents
	// no type for insert_text
	// no type for copy_text
	// no type for cut_text
	// no type for delete_text
	// no type for paste_text
}

// GObjectAccessibleClass is a wrapper around the C record AtkGObjectAccessibleClass.
type GObjectAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
}

// HyperlinkClass is a wrapper around the C record AtkHyperlinkClass.
type HyperlinkClass struct {
	native unsafe.Pointer
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

// HyperlinkImplIface is a wrapper around the C record AtkHyperlinkImplIface.
type HyperlinkImplIface struct {
	native unsafe.Pointer
	// parent : record
	// no type for get_hyperlink
}

// HypertextIface is a wrapper around the C record AtkHypertextIface.
type HypertextIface struct {
	native unsafe.Pointer
	// parent : record
	// no type for get_link
	// no type for get_n_links
	// no type for get_link_index
	// no type for link_selected
}

// ImageIface is a wrapper around the C record AtkImageIface.
type ImageIface struct {
	native unsafe.Pointer
	// parent : record
	// no type for get_image_position
	// no type for get_image_description
	// no type for get_image_size
	// no type for set_image_description
	// no type for get_image_locale
}

// Implementor is a wrapper around the C record AtkImplementor.
type Implementor struct {
	native unsafe.Pointer
}

// KeyEventStruct is a wrapper around the C record AtkKeyEventStruct.
type KeyEventStruct struct {
	native    unsafe.Pointer
	Type      int32
	State     uint32
	Keyval    uint32
	Length    int32
	String    string
	Keycode   uint16
	Timestamp uint32
}

// MiscClass is a wrapper around the C record AtkMiscClass.
type MiscClass struct {
	native unsafe.Pointer
	// parent : record
	// no type for threads_enter
	// no type for threads_leave
	// no type for vfuncs
}

// NoOpObjectClass is a wrapper around the C record AtkNoOpObjectClass.
type NoOpObjectClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// NoOpObjectFactoryClass is a wrapper around the C record AtkNoOpObjectFactoryClass.
type NoOpObjectFactoryClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ObjectClass is a wrapper around the C record AtkObjectClass.
type ObjectClass struct {
	native unsafe.Pointer
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

// ObjectFactoryClass is a wrapper around the C record AtkObjectFactoryClass.
type ObjectFactoryClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for create_accessible
	// no type for invalidate
	// no type for get_accessible_type
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
}

// PlugClass is a wrapper around the C record AtkPlugClass.
type PlugClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_object_id
}

// PropertyValues is a wrapper around the C record AtkPropertyValues.
type PropertyValues struct {
	native       unsafe.Pointer
	PropertyName string
	// old_value : record
	// new_value : record
}

// Range is a wrapper around the C record AtkRange.
type Range struct {
	native unsafe.Pointer
}

// Rectangle is a wrapper around the C record AtkRectangle.
type Rectangle struct {
	native unsafe.Pointer
	X      int32
	Y      int32
	Width  int32
	Height int32
}

// RegistryClass is a wrapper around the C record AtkRegistryClass.
type RegistryClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// RelationClass is a wrapper around the C record AtkRelationClass.
type RelationClass struct {
	native unsafe.Pointer
	// parent : record
}

// RelationSetClass is a wrapper around the C record AtkRelationSetClass.
type RelationSetClass struct {
	native unsafe.Pointer
	// parent : record
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
}

// SelectionIface is a wrapper around the C record AtkSelectionIface.
type SelectionIface struct {
	native unsafe.Pointer
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

// SocketClass is a wrapper around the C record AtkSocketClass.
type SocketClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for embed
}

// StateSetClass is a wrapper around the C record AtkStateSetClass.
type StateSetClass struct {
	native unsafe.Pointer
	// parent : record
}

// StreamableContentIface is a wrapper around the C record AtkStreamableContentIface.
type StreamableContentIface struct {
	native unsafe.Pointer
	// parent : record
	// no type for get_n_mime_types
	// no type for get_mime_type
	// no type for get_stream
	// no type for get_uri
	// pad1 : no type generator for Function, AtkFunction
	// pad2 : no type generator for Function, AtkFunction
	// pad3 : no type generator for Function, AtkFunction
}

// TableCellIface is a wrapper around the C record AtkTableCellIface.
type TableCellIface struct {
	native unsafe.Pointer
	// parent : record
	// no type for get_column_span
	// no type for get_column_header_cells
	// no type for get_position
	// no type for get_row_span
	// no type for get_row_header_cells
	// no type for get_row_column_span
	// no type for get_table
}

// TableIface is a wrapper around the C record AtkTableIface.
type TableIface struct {
	native unsafe.Pointer
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

// TextIface is a wrapper around the C record AtkTextIface.
type TextIface struct {
	native unsafe.Pointer
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

// TextRange is a wrapper around the C record AtkTextRange.
type TextRange struct {
	native unsafe.Pointer
	// bounds : record
	StartOffset int32
	EndOffset   int32
	Content     string
}

// TextRectangle is a wrapper around the C record AtkTextRectangle.
type TextRectangle struct {
	native unsafe.Pointer
	X      int32
	Y      int32
	Width  int32
	Height int32
}

// UtilClass is a wrapper around the C record AtkUtilClass.
type UtilClass struct {
	native unsafe.Pointer
	// parent : record
	// no type for add_global_event_listener
	// no type for remove_global_event_listener
	// no type for add_key_event_listener
	// no type for remove_key_event_listener
	// no type for get_root
	// no type for get_toolkit_name
	// no type for get_toolkit_version
}

// ValueIface is a wrapper around the C record AtkValueIface.
type ValueIface struct {
	native unsafe.Pointer
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

// WindowIface is a wrapper around the C record AtkWindowIface.
type WindowIface struct {
	native unsafe.Pointer
	// parent : record
}
