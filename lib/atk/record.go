// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var actionIfaceStruct *gi.Struct
var actionIfaceStruct_Once sync.Once

func actionIfaceStruct_Set() {
	actionIfaceStruct_Once.Do(func() {
		actionIfaceStruct = gi.StructNew("Atk", "ActionIface")
	})
}

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

var attributeStruct *gi.Struct
var attributeStruct_Once sync.Once

func attributeStruct_Set() {
	attributeStruct_Once.Do(func() {
		attributeStruct = gi.StructNew("Atk", "Attribute")
	})
}

type Attribute struct {
	native uintptr
	Name   string
	Value  string
}

var componentIfaceStruct *gi.Struct
var componentIfaceStruct_Once sync.Once

func componentIfaceStruct_Set() {
	componentIfaceStruct_Once.Do(func() {
		componentIfaceStruct = gi.StructNew("Atk", "ComponentIface")
	})
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

var documentIfaceStruct *gi.Struct
var documentIfaceStruct_Once sync.Once

func documentIfaceStruct_Set() {
	documentIfaceStruct_Once.Do(func() {
		documentIfaceStruct = gi.StructNew("Atk", "DocumentIface")
	})
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

var editableTextIfaceStruct *gi.Struct
var editableTextIfaceStruct_Once sync.Once

func editableTextIfaceStruct_Set() {
	editableTextIfaceStruct_Once.Do(func() {
		editableTextIfaceStruct = gi.StructNew("Atk", "EditableTextIface")
	})
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

var gObjectAccessibleClassStruct *gi.Struct
var gObjectAccessibleClassStruct_Once sync.Once

func gObjectAccessibleClassStruct_Set() {
	gObjectAccessibleClassStruct_Once.Do(func() {
		gObjectAccessibleClassStruct = gi.StructNew("Atk", "GObjectAccessibleClass")
	})
}

type GObjectAccessibleClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'
	// UNSUPPORTED : C value 'pad2' : no Go type for 'Function'
}

var hyperlinkClassStruct *gi.Struct
var hyperlinkClassStruct_Once sync.Once

func hyperlinkClassStruct_Set() {
	hyperlinkClassStruct_Once.Do(func() {
		hyperlinkClassStruct = gi.StructNew("Atk", "HyperlinkClass")
	})
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

var hyperlinkImplIfaceStruct *gi.Struct
var hyperlinkImplIfaceStruct_Once sync.Once

func hyperlinkImplIfaceStruct_Set() {
	hyperlinkImplIfaceStruct_Once.Do(func() {
		hyperlinkImplIfaceStruct = gi.StructNew("Atk", "HyperlinkImplIface")
	})
}

type HyperlinkImplIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_hyperlink' : missing Type
}

var hypertextIfaceStruct *gi.Struct
var hypertextIfaceStruct_Once sync.Once

func hypertextIfaceStruct_Set() {
	hypertextIfaceStruct_Once.Do(func() {
		hypertextIfaceStruct = gi.StructNew("Atk", "HypertextIface")
	})
}

type HypertextIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_link' : missing Type
	// UNSUPPORTED : C value 'get_n_links' : missing Type
	// UNSUPPORTED : C value 'get_link_index' : missing Type
	// UNSUPPORTED : C value 'link_selected' : missing Type
}

var imageIfaceStruct *gi.Struct
var imageIfaceStruct_Once sync.Once

func imageIfaceStruct_Set() {
	imageIfaceStruct_Once.Do(func() {
		imageIfaceStruct = gi.StructNew("Atk", "ImageIface")
	})
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

var implementorStruct *gi.Struct
var implementorStruct_Once sync.Once

func implementorStruct_Set() {
	implementorStruct_Once.Do(func() {
		implementorStruct = gi.StructNew("Atk", "Implementor")
	})
}

type Implementor struct {
	native uintptr
}

// UNSUPPORTED : C value 'atk_implementor_ref_accessible' : return type 'Object' not supported

var keyEventStructStruct *gi.Struct
var keyEventStructStruct_Once sync.Once

func keyEventStructStruct_Set() {
	keyEventStructStruct_Once.Do(func() {
		keyEventStructStruct = gi.StructNew("Atk", "KeyEventStruct")
	})
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

var miscClassStruct *gi.Struct
var miscClassStruct_Once sync.Once

func miscClassStruct_Set() {
	miscClassStruct_Once.Do(func() {
		miscClassStruct = gi.StructNew("Atk", "MiscClass")
	})
}

type MiscClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'threads_enter' : missing Type
	// UNSUPPORTED : C value 'threads_leave' : missing Type
	// UNSUPPORTED : C value 'vfuncs' : missing Type
}

var noOpObjectClassStruct *gi.Struct
var noOpObjectClassStruct_Once sync.Once

func noOpObjectClassStruct_Set() {
	noOpObjectClassStruct_Once.Do(func() {
		noOpObjectClassStruct = gi.StructNew("Atk", "NoOpObjectClass")
	})
}

type NoOpObjectClass struct {
	native      uintptr
	ParentClass *ObjectClass
}

var noOpObjectFactoryClassStruct *gi.Struct
var noOpObjectFactoryClassStruct_Once sync.Once

func noOpObjectFactoryClassStruct_Set() {
	noOpObjectFactoryClassStruct_Once.Do(func() {
		noOpObjectFactoryClassStruct = gi.StructNew("Atk", "NoOpObjectFactoryClass")
	})
}

type NoOpObjectFactoryClass struct {
	native      uintptr
	ParentClass *ObjectFactoryClass
}

var objectClassStruct *gi.Struct
var objectClassStruct_Once sync.Once

func objectClassStruct_Set() {
	objectClassStruct_Once.Do(func() {
		objectClassStruct = gi.StructNew("Atk", "ObjectClass")
	})
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

var objectFactoryClassStruct *gi.Struct
var objectFactoryClassStruct_Once sync.Once

func objectFactoryClassStruct_Set() {
	objectFactoryClassStruct_Once.Do(func() {
		objectFactoryClassStruct = gi.StructNew("Atk", "ObjectFactoryClass")
	})
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

var plugClassStruct *gi.Struct
var plugClassStruct_Once sync.Once

func plugClassStruct_Set() {
	plugClassStruct_Once.Do(func() {
		plugClassStruct = gi.StructNew("Atk", "PlugClass")
	})
}

type PlugClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'get_object_id' : missing Type
}

var propertyValuesStruct *gi.Struct
var propertyValuesStruct_Once sync.Once

func propertyValuesStruct_Set() {
	propertyValuesStruct_Once.Do(func() {
		propertyValuesStruct = gi.StructNew("Atk", "PropertyValues")
	})
}

type PropertyValues struct {
	native       uintptr
	PropertyName string
	// UNSUPPORTED : C value 'old_value' : no Go type for 'GObject.Value'
	// UNSUPPORTED : C value 'new_value' : no Go type for 'GObject.Value'
}

var rangeStruct *gi.Struct
var rangeStruct_Once sync.Once

func rangeStruct_Set() {
	rangeStruct_Once.Do(func() {
		rangeStruct = gi.StructNew("Atk", "Range")
	})
}

type Range struct {
	native uintptr
}

// UNSUPPORTED : C value 'atk_range_new' : parameter 'lower_limit' of type 'gdouble' not supported

var rangeCopyFunction *gi.Function
var rangeCopyFunction_Once sync.Once

func rangeCopyFunction_Set() {
	rangeCopyFunction_Once.Do(func() {
		rangeStruct_Set()
		rangeCopyFunction = rangeStruct.InvokerNew("copy")
	})
}

// Copy is a representation of the C type atk_range_copy.
func (recv *Range) Copy() *Range {
	rangeCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := rangeCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Range{native: ret.Pointer()}

	return retGo
}

var rangeFreeFunction *gi.Function
var rangeFreeFunction_Once sync.Once

func rangeFreeFunction_Set() {
	rangeFreeFunction_Once.Do(func() {
		rangeStruct_Set()
		rangeFreeFunction = rangeStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type atk_range_free.
func (recv *Range) Free() {
	rangeFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	rangeFreeFunction.Invoke(inArgs[:], nil)

}

var rangeGetDescriptionFunction *gi.Function
var rangeGetDescriptionFunction_Once sync.Once

func rangeGetDescriptionFunction_Set() {
	rangeGetDescriptionFunction_Once.Do(func() {
		rangeStruct_Set()
		rangeGetDescriptionFunction = rangeStruct.InvokerNew("get_description")
	})
}

// GetDescription is a representation of the C type atk_range_get_description.
func (recv *Range) GetDescription() string {
	rangeGetDescriptionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := rangeGetDescriptionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'atk_range_get_lower_limit' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'atk_range_get_upper_limit' : return type 'gdouble' not supported

var rectangleStruct *gi.Struct
var rectangleStruct_Once sync.Once

func rectangleStruct_Set() {
	rectangleStruct_Once.Do(func() {
		rectangleStruct = gi.StructNew("Atk", "Rectangle")
	})
}

type Rectangle struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

var registryClassStruct *gi.Struct
var registryClassStruct_Once sync.Once

func registryClassStruct_Set() {
	registryClassStruct_Once.Do(func() {
		registryClassStruct = gi.StructNew("Atk", "RegistryClass")
	})
}

type RegistryClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var relationClassStruct *gi.Struct
var relationClassStruct_Once sync.Once

func relationClassStruct_Set() {
	relationClassStruct_Once.Do(func() {
		relationClassStruct = gi.StructNew("Atk", "RelationClass")
	})
}

type RelationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
}

var relationSetClassStruct *gi.Struct
var relationSetClassStruct_Once sync.Once

func relationSetClassStruct_Set() {
	relationSetClassStruct_Once.Do(func() {
		relationSetClassStruct = gi.StructNew("Atk", "RelationSetClass")
	})
}

type RelationSetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'
	// UNSUPPORTED : C value 'pad2' : no Go type for 'Function'
}

var selectionIfaceStruct *gi.Struct
var selectionIfaceStruct_Once sync.Once

func selectionIfaceStruct_Set() {
	selectionIfaceStruct_Once.Do(func() {
		selectionIfaceStruct = gi.StructNew("Atk", "SelectionIface")
	})
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

var socketClassStruct *gi.Struct
var socketClassStruct_Once sync.Once

func socketClassStruct_Set() {
	socketClassStruct_Once.Do(func() {
		socketClassStruct = gi.StructNew("Atk", "SocketClass")
	})
}

type SocketClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'embed' : missing Type
}

var stateSetClassStruct *gi.Struct
var stateSetClassStruct_Once sync.Once

func stateSetClassStruct_Set() {
	stateSetClassStruct_Once.Do(func() {
		stateSetClassStruct = gi.StructNew("Atk", "StateSetClass")
	})
}

type StateSetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
}

var streamableContentIfaceStruct *gi.Struct
var streamableContentIfaceStruct_Once sync.Once

func streamableContentIfaceStruct_Set() {
	streamableContentIfaceStruct_Once.Do(func() {
		streamableContentIfaceStruct = gi.StructNew("Atk", "StreamableContentIface")
	})
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

var tableCellIfaceStruct *gi.Struct
var tableCellIfaceStruct_Once sync.Once

func tableCellIfaceStruct_Set() {
	tableCellIfaceStruct_Once.Do(func() {
		tableCellIfaceStruct = gi.StructNew("Atk", "TableCellIface")
	})
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

var tableIfaceStruct *gi.Struct
var tableIfaceStruct_Once sync.Once

func tableIfaceStruct_Set() {
	tableIfaceStruct_Once.Do(func() {
		tableIfaceStruct = gi.StructNew("Atk", "TableIface")
	})
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

var textIfaceStruct *gi.Struct
var textIfaceStruct_Once sync.Once

func textIfaceStruct_Set() {
	textIfaceStruct_Once.Do(func() {
		textIfaceStruct = gi.StructNew("Atk", "TextIface")
	})
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

var textRangeStruct *gi.Struct
var textRangeStruct_Once sync.Once

func textRangeStruct_Set() {
	textRangeStruct_Once.Do(func() {
		textRangeStruct = gi.StructNew("Atk", "TextRange")
	})
}

type TextRange struct {
	native      uintptr
	Bounds      *TextRectangle
	StartOffset int32
	EndOffset   int32
	Content     string
}

var textRectangleStruct *gi.Struct
var textRectangleStruct_Once sync.Once

func textRectangleStruct_Set() {
	textRectangleStruct_Once.Do(func() {
		textRectangleStruct = gi.StructNew("Atk", "TextRectangle")
	})
}

type TextRectangle struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

var utilClassStruct *gi.Struct
var utilClassStruct_Once sync.Once

func utilClassStruct_Set() {
	utilClassStruct_Once.Do(func() {
		utilClassStruct = gi.StructNew("Atk", "UtilClass")
	})
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

var valueIfaceStruct *gi.Struct
var valueIfaceStruct_Once sync.Once

func valueIfaceStruct_Set() {
	valueIfaceStruct_Once.Do(func() {
		valueIfaceStruct = gi.StructNew("Atk", "ValueIface")
	})
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

var windowIfaceStruct *gi.Struct
var windowIfaceStruct_Once sync.Once

func windowIfaceStruct_Set() {
	windowIfaceStruct_Once.Do(func() {
		windowIfaceStruct = gi.StructNew("Atk", "WindowIface")
	})
}

type WindowIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
}
