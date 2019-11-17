// Code generated - DO NOT EDIT.

package atk

type ActionIface struct {
	native uintptr
	// UNSUPPORTED : C value 'do_action' : missing Type

	// UNSUPPORTED : C value 'get_n_actions' : missing Type

	// UNSUPPORTED : C value 'get_description' : missing Type

	// UNSUPPORTED : C value 'get_name' : missing Type

	// UNSUPPORTED : C value 'get_keybinding' : missing Type

	// UNSUPPORTED : C value 'set_description' : missing Type

	// UNSUPPORTED : C value 'get_localized_name' : missing Type

}

type Attribute struct {
	native uintptr
	Name   string
	Value  string
}

type ComponentIface struct {
	native uintptr
	// UNSUPPORTED : C value 'add_focus_handler' : missing Type

	// UNSUPPORTED : C value 'contains' : missing Type

	// UNSUPPORTED : C value 'ref_accessible_at_point' : missing Type

	// UNSUPPORTED : C value 'get_extents' : missing Type

	// UNSUPPORTED : C value 'get_position' : missing Type

	// UNSUPPORTED : C value 'get_size' : missing Type

	// UNSUPPORTED : C value 'grab_focus' : missing Type

	// UNSUPPORTED : C value 'remove_focus_handler' : missing Type

	// UNSUPPORTED : C value 'set_extents' : missing Type

	// UNSUPPORTED : C value 'set_position' : missing Type

	// UNSUPPORTED : C value 'set_size' : missing Type

	// UNSUPPORTED : C value 'get_layer' : missing Type

	// UNSUPPORTED : C value 'get_mdi_zorder' : missing Type

	// UNSUPPORTED : C value 'bounds_changed' : missing Type

	// UNSUPPORTED : C value 'get_alpha' : missing Type

	// UNSUPPORTED : C value 'scroll_to' : missing Type

	// UNSUPPORTED : C value 'scroll_to_point' : missing Type

}

type DocumentIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_document_type' : missing Type

	// UNSUPPORTED : C value 'get_document' : missing Type

	// UNSUPPORTED : C value 'get_document_locale' : missing Type

	// UNSUPPORTED : C value 'get_document_attributes' : missing Type

	// UNSUPPORTED : C value 'get_document_attribute_value' : missing Type

	// UNSUPPORTED : C value 'set_document_attribute' : missing Type

	// UNSUPPORTED : C value 'get_current_page_number' : missing Type

	// UNSUPPORTED : C value 'get_page_count' : missing Type

}

type EditableTextIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_interface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'set_run_attributes' : missing Type

	// UNSUPPORTED : C value 'set_text_contents' : missing Type

	// UNSUPPORTED : C value 'insert_text' : missing Type

	// UNSUPPORTED : C value 'copy_text' : missing Type

	// UNSUPPORTED : C value 'cut_text' : missing Type

	// UNSUPPORTED : C value 'delete_text' : missing Type

	// UNSUPPORTED : C value 'paste_text' : missing Type

}

type GObjectAccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'ObjectClass'

	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'

	// UNSUPPORTED : C value 'pad2' : no Go type for 'Function'

}

type HyperlinkClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_uri' : missing Type

	// UNSUPPORTED : C value 'get_object' : missing Type

	// UNSUPPORTED : C value 'get_end_index' : missing Type

	// UNSUPPORTED : C value 'get_start_index' : missing Type

	// UNSUPPORTED : C value 'is_valid' : missing Type

	// UNSUPPORTED : C value 'get_n_anchors' : missing Type

	// UNSUPPORTED : C value 'link_state' : missing Type

	// UNSUPPORTED : C value 'is_selected_link' : missing Type

	// UNSUPPORTED : C value 'link_activated' : missing Type

	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'

}

type HyperlinkImplIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_hyperlink' : missing Type

}

type HypertextIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_link' : missing Type

	// UNSUPPORTED : C value 'get_n_links' : missing Type

	// UNSUPPORTED : C value 'get_link_index' : missing Type

	// UNSUPPORTED : C value 'link_selected' : missing Type

}

type ImageIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_image_position' : missing Type

	// UNSUPPORTED : C value 'get_image_description' : missing Type

	// UNSUPPORTED : C value 'get_image_size' : missing Type

	// UNSUPPORTED : C value 'set_image_description' : missing Type

	// UNSUPPORTED : C value 'get_image_locale' : missing Type

}

type Implementor struct {
	native uintptr
}

type KeyEventStruct struct {
	native    uintptr
	Type      int32
	State     uint32
	Keyval    uint32
	Length    int32
	String    string
	Keycode   uint16
	Timestamp uint32
}

type MiscClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'threads_enter' : missing Type

	// UNSUPPORTED : C value 'threads_leave' : missing Type

	// UNSUPPORTED : C value 'vfuncs' : missing Type

}

type NoOpObjectClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'ObjectClass'

}

type NoOpObjectFactoryClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'ObjectFactoryClass'

}

type ObjectClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_name' : missing Type

	// UNSUPPORTED : C value 'get_description' : missing Type

	// UNSUPPORTED : C value 'get_parent' : missing Type

	// UNSUPPORTED : C value 'get_n_children' : missing Type

	// UNSUPPORTED : C value 'ref_child' : missing Type

	// UNSUPPORTED : C value 'get_index_in_parent' : missing Type

	// UNSUPPORTED : C value 'ref_relation_set' : missing Type

	// UNSUPPORTED : C value 'get_role' : missing Type

	// UNSUPPORTED : C value 'get_layer' : missing Type

	// UNSUPPORTED : C value 'get_mdi_zorder' : missing Type

	// UNSUPPORTED : C value 'ref_state_set' : missing Type

	// UNSUPPORTED : C value 'set_name' : missing Type

	// UNSUPPORTED : C value 'set_description' : missing Type

	// UNSUPPORTED : C value 'set_parent' : missing Type

	// UNSUPPORTED : C value 'set_role' : missing Type

	// UNSUPPORTED : C value 'connect_property_change_handler' : missing Type

	// UNSUPPORTED : C value 'remove_property_change_handler' : missing Type

	// UNSUPPORTED : C value 'initialize' : missing Type

	// UNSUPPORTED : C value 'children_changed' : missing Type

	// UNSUPPORTED : C value 'focus_event' : missing Type

	// UNSUPPORTED : C value 'property_change' : missing Type

	// UNSUPPORTED : C value 'state_change' : missing Type

	// UNSUPPORTED : C value 'visible_data_changed' : missing Type

	// UNSUPPORTED : C value 'active_descendant_changed' : missing Type

	// UNSUPPORTED : C value 'get_attributes' : missing Type

	// UNSUPPORTED : C value 'get_object_locale' : missing Type

	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'

}

type ObjectFactoryClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'create_accessible' : missing Type

	// UNSUPPORTED : C value 'invalidate' : missing Type

	// UNSUPPORTED : C value 'get_accessible_type' : missing Type

	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'

	// UNSUPPORTED : C value 'pad2' : no Go type for 'Function'

}

type PlugClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'ObjectClass'

	// UNSUPPORTED : C value 'get_object_id' : missing Type

}

type PropertyValues struct {
	native       uintptr
	PropertyName string
	// UNSUPPORTED : C value 'old_value' : no Go type for 'GObject.Value'

	// UNSUPPORTED : C value 'new_value' : no Go type for 'GObject.Value'

}

type Range struct {
	native uintptr
}

// UNSUPPORTED : C value 'atk_range_new' : parameter 'lower_limit' of type 'gdouble' not supported

type Rectangle struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type RegistryClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

}

type RelationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'

}

type RelationSetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'

	// UNSUPPORTED : C value 'pad2' : no Go type for 'Function'

}

type SelectionIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'add_selection' : missing Type

	// UNSUPPORTED : C value 'clear_selection' : missing Type

	// UNSUPPORTED : C value 'ref_selection' : missing Type

	// UNSUPPORTED : C value 'get_selection_count' : missing Type

	// UNSUPPORTED : C value 'is_child_selected' : missing Type

	// UNSUPPORTED : C value 'remove_selection' : missing Type

	// UNSUPPORTED : C value 'select_all_selection' : missing Type

	// UNSUPPORTED : C value 'selection_changed' : missing Type

}

type SocketClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'ObjectClass'

	// UNSUPPORTED : C value 'embed' : missing Type

}

type StateSetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'

}

type StreamableContentIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_n_mime_types' : missing Type

	// UNSUPPORTED : C value 'get_mime_type' : missing Type

	// UNSUPPORTED : C value 'get_stream' : missing Type

	// UNSUPPORTED : C value 'get_uri' : missing Type

	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'

	// UNSUPPORTED : C value 'pad2' : no Go type for 'Function'

	// UNSUPPORTED : C value 'pad3' : no Go type for 'Function'

}

type TableCellIface struct {
	native uintptr
	// UNSUPPORTED : C value 'get_column_span' : missing Type

	// UNSUPPORTED : C value 'get_column_header_cells' : missing Type

	// UNSUPPORTED : C value 'get_position' : missing Type

	// UNSUPPORTED : C value 'get_row_span' : missing Type

	// UNSUPPORTED : C value 'get_row_header_cells' : missing Type

	// UNSUPPORTED : C value 'get_row_column_span' : missing Type

	// UNSUPPORTED : C value 'get_table' : missing Type

}

type TableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'ref_at' : missing Type

	// UNSUPPORTED : C value 'get_index_at' : missing Type

	// UNSUPPORTED : C value 'get_column_at_index' : missing Type

	// UNSUPPORTED : C value 'get_row_at_index' : missing Type

	// UNSUPPORTED : C value 'get_n_columns' : missing Type

	// UNSUPPORTED : C value 'get_n_rows' : missing Type

	// UNSUPPORTED : C value 'get_column_extent_at' : missing Type

	// UNSUPPORTED : C value 'get_row_extent_at' : missing Type

	// UNSUPPORTED : C value 'get_caption' : missing Type

	// UNSUPPORTED : C value 'get_column_description' : missing Type

	// UNSUPPORTED : C value 'get_column_header' : missing Type

	// UNSUPPORTED : C value 'get_row_description' : missing Type

	// UNSUPPORTED : C value 'get_row_header' : missing Type

	// UNSUPPORTED : C value 'get_summary' : missing Type

	// UNSUPPORTED : C value 'set_caption' : missing Type

	// UNSUPPORTED : C value 'set_column_description' : missing Type

	// UNSUPPORTED : C value 'set_column_header' : missing Type

	// UNSUPPORTED : C value 'set_row_description' : missing Type

	// UNSUPPORTED : C value 'set_row_header' : missing Type

	// UNSUPPORTED : C value 'set_summary' : missing Type

	// UNSUPPORTED : C value 'get_selected_columns' : missing Type

	// UNSUPPORTED : C value 'get_selected_rows' : missing Type

	// UNSUPPORTED : C value 'is_column_selected' : missing Type

	// UNSUPPORTED : C value 'is_row_selected' : missing Type

	// UNSUPPORTED : C value 'is_selected' : missing Type

	// UNSUPPORTED : C value 'add_row_selection' : missing Type

	// UNSUPPORTED : C value 'remove_row_selection' : missing Type

	// UNSUPPORTED : C value 'add_column_selection' : missing Type

	// UNSUPPORTED : C value 'remove_column_selection' : missing Type

	// UNSUPPORTED : C value 'row_inserted' : missing Type

	// UNSUPPORTED : C value 'column_inserted' : missing Type

	// UNSUPPORTED : C value 'row_deleted' : missing Type

	// UNSUPPORTED : C value 'column_deleted' : missing Type

	// UNSUPPORTED : C value 'row_reordered' : missing Type

	// UNSUPPORTED : C value 'column_reordered' : missing Type

	// UNSUPPORTED : C value 'model_changed' : missing Type

}

type TextIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_text' : missing Type

	// UNSUPPORTED : C value 'get_text_after_offset' : missing Type

	// UNSUPPORTED : C value 'get_text_at_offset' : missing Type

	// UNSUPPORTED : C value 'get_character_at_offset' : missing Type

	// UNSUPPORTED : C value 'get_text_before_offset' : missing Type

	// UNSUPPORTED : C value 'get_caret_offset' : missing Type

	// UNSUPPORTED : C value 'get_run_attributes' : missing Type

	// UNSUPPORTED : C value 'get_default_attributes' : missing Type

	// UNSUPPORTED : C value 'get_character_extents' : missing Type

	// UNSUPPORTED : C value 'get_character_count' : missing Type

	// UNSUPPORTED : C value 'get_offset_at_point' : missing Type

	// UNSUPPORTED : C value 'get_n_selections' : missing Type

	// UNSUPPORTED : C value 'get_selection' : missing Type

	// UNSUPPORTED : C value 'add_selection' : missing Type

	// UNSUPPORTED : C value 'remove_selection' : missing Type

	// UNSUPPORTED : C value 'set_selection' : missing Type

	// UNSUPPORTED : C value 'set_caret_offset' : missing Type

	// UNSUPPORTED : C value 'text_changed' : missing Type

	// UNSUPPORTED : C value 'text_caret_moved' : missing Type

	// UNSUPPORTED : C value 'text_selection_changed' : missing Type

	// UNSUPPORTED : C value 'text_attributes_changed' : missing Type

	// UNSUPPORTED : C value 'get_range_extents' : missing Type

	// UNSUPPORTED : C value 'get_bounded_ranges' : missing Type

	// UNSUPPORTED : C value 'get_string_at_offset' : missing Type

	// UNSUPPORTED : C value 'scroll_substring_to' : missing Type

	// UNSUPPORTED : C value 'scroll_substring_to_point' : missing Type

}

type TextRange struct {
	native uintptr
	// UNSUPPORTED : C value 'bounds' : no Go type for 'TextRectangle'

	StartOffset int32
	EndOffset   int32
	Content     string
}

type TextRectangle struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type UtilClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'add_global_event_listener' : missing Type

	// UNSUPPORTED : C value 'remove_global_event_listener' : missing Type

	// UNSUPPORTED : C value 'add_key_event_listener' : missing Type

	// UNSUPPORTED : C value 'remove_key_event_listener' : missing Type

	// UNSUPPORTED : C value 'get_root' : missing Type

	// UNSUPPORTED : C value 'get_toolkit_name' : missing Type

	// UNSUPPORTED : C value 'get_toolkit_version' : missing Type

}

type ValueIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_current_value' : missing Type

	// UNSUPPORTED : C value 'get_maximum_value' : missing Type

	// UNSUPPORTED : C value 'get_minimum_value' : missing Type

	// UNSUPPORTED : C value 'set_current_value' : missing Type

	// UNSUPPORTED : C value 'get_minimum_increment' : missing Type

	// UNSUPPORTED : C value 'get_value_and_text' : missing Type

	// UNSUPPORTED : C value 'get_range' : missing Type

	// UNSUPPORTED : C value 'get_increment' : missing Type

	// UNSUPPORTED : C value 'get_sub_ranges' : missing Type

	// UNSUPPORTED : C value 'set_value' : missing Type

}

type WindowIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

}
