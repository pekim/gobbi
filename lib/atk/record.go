// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var actionIfaceStruct *gi.Struct
var actionIfaceStructOnce sync.Once

func actionIfaceStructSet() {
	actionIfaceStructOnce.Do(func() {
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
var attributeStructOnce sync.Once

func attributeStructSet() {
	attributeStructOnce.Do(func() {
		attributeStruct = gi.StructNew("Atk", "Attribute")
	})
}

type Attribute struct {
	native uintptr
	Name   string
	Value  string
}

var componentIfaceStruct *gi.Struct
var componentIfaceStructOnce sync.Once

func componentIfaceStructSet() {
	componentIfaceStructOnce.Do(func() {
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
var documentIfaceStructOnce sync.Once

func documentIfaceStructSet() {
	documentIfaceStructOnce.Do(func() {
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
var editableTextIfaceStructOnce sync.Once

func editableTextIfaceStructSet() {
	editableTextIfaceStructOnce.Do(func() {
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
var gObjectAccessibleClassStructOnce sync.Once

func gObjectAccessibleClassStructSet() {
	gObjectAccessibleClassStructOnce.Do(func() {
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
var hyperlinkClassStructOnce sync.Once

func hyperlinkClassStructSet() {
	hyperlinkClassStructOnce.Do(func() {
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
var hyperlinkImplIfaceStructOnce sync.Once

func hyperlinkImplIfaceStructSet() {
	hyperlinkImplIfaceStructOnce.Do(func() {
		hyperlinkImplIfaceStruct = gi.StructNew("Atk", "HyperlinkImplIface")
	})
}

type HyperlinkImplIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_hyperlink' : missing Type
}

var hypertextIfaceStruct *gi.Struct
var hypertextIfaceStructOnce sync.Once

func hypertextIfaceStructSet() {
	hypertextIfaceStructOnce.Do(func() {
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
var imageIfaceStructOnce sync.Once

func imageIfaceStructSet() {
	imageIfaceStructOnce.Do(func() {
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
var implementorStructOnce sync.Once

func implementorStructSet() {
	implementorStructOnce.Do(func() {
		implementorStruct = gi.StructNew("Atk", "Implementor")
	})
}

type Implementor struct {
	native uintptr
}

// UNSUPPORTED : C value 'atk_implementor_ref_accessible' : return type 'Object' not supported

var keyEventStructStruct *gi.Struct
var keyEventStructStructOnce sync.Once

func keyEventStructStructSet() {
	keyEventStructStructOnce.Do(func() {
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
var miscClassStructOnce sync.Once

func miscClassStructSet() {
	miscClassStructOnce.Do(func() {
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
var noOpObjectClassStructOnce sync.Once

func noOpObjectClassStructSet() {
	noOpObjectClassStructOnce.Do(func() {
		noOpObjectClassStruct = gi.StructNew("Atk", "NoOpObjectClass")
	})
}

type NoOpObjectClass struct {
	native      uintptr
	ParentClass *ObjectClass
}

var noOpObjectFactoryClassStruct *gi.Struct
var noOpObjectFactoryClassStructOnce sync.Once

func noOpObjectFactoryClassStructSet() {
	noOpObjectFactoryClassStructOnce.Do(func() {
		noOpObjectFactoryClassStruct = gi.StructNew("Atk", "NoOpObjectFactoryClass")
	})
}

type NoOpObjectFactoryClass struct {
	native      uintptr
	ParentClass *ObjectFactoryClass
}

var objectClassStruct *gi.Struct
var objectClassStructOnce sync.Once

func objectClassStructSet() {
	objectClassStructOnce.Do(func() {
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
var objectFactoryClassStructOnce sync.Once

func objectFactoryClassStructSet() {
	objectFactoryClassStructOnce.Do(func() {
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
var plugClassStructOnce sync.Once

func plugClassStructSet() {
	plugClassStructOnce.Do(func() {
		plugClassStruct = gi.StructNew("Atk", "PlugClass")
	})
}

type PlugClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'get_object_id' : missing Type
}

var propertyValuesStruct *gi.Struct
var propertyValuesStructOnce sync.Once

func propertyValuesStructSet() {
	propertyValuesStructOnce.Do(func() {
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
var rangeStructOnce sync.Once

func rangeStructSet() {
	rangeStructOnce.Do(func() {
		rangeStruct = gi.StructNew("Atk", "Range")
	})
}

type Range struct {
	native uintptr
}

// UNSUPPORTED : C value 'atk_range_new' : parameter 'lower_limit' of type 'gdouble' not supported

var rangeCopyFunction *gi.Function
var rangeCopyFunctionOnce sync.Once

func rangeCopyFunctionSet() {
	rangeCopyFunctionOnce.Do(func() {
		rangeCopyFunction = gi.FunctionInvokerNew("Atk", "copy")
	})
}

var copyRangeInvoker *gi.Function

// Copy is a representation of the C type atk_range_copy.
func (recv *Range) Copy() *Range {
	if copyRangeInvoker == nil {
		copyRangeInvoker = gi.StructFunctionInvokerNew("Atk", "Range", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyRangeInvoker.Invoke(inArgs[:], nil)

	retGo := &Range{native: ret.Pointer()}

	return retGo
}

var rangeFreeFunction *gi.Function
var rangeFreeFunctionOnce sync.Once

func rangeFreeFunctionSet() {
	rangeFreeFunctionOnce.Do(func() {
		rangeFreeFunction = gi.FunctionInvokerNew("Atk", "free")
	})
}

var freeRangeInvoker *gi.Function

// Free is a representation of the C type atk_range_free.
func (recv *Range) Free() {
	if freeRangeInvoker == nil {
		freeRangeInvoker = gi.StructFunctionInvokerNew("Atk", "Range", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeRangeInvoker.Invoke(inArgs[:], nil)

}

var rangeGetDescriptionFunction *gi.Function
var rangeGetDescriptionFunctionOnce sync.Once

func rangeGetDescriptionFunctionSet() {
	rangeGetDescriptionFunctionOnce.Do(func() {
		rangeGetDescriptionFunction = gi.FunctionInvokerNew("Atk", "get_description")
	})
}

var getDescriptionRangeInvoker *gi.Function

// GetDescription is a representation of the C type atk_range_get_description.
func (recv *Range) GetDescription() string {
	if getDescriptionRangeInvoker == nil {
		getDescriptionRangeInvoker = gi.StructFunctionInvokerNew("Atk", "Range", "get_description")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDescriptionRangeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'atk_range_get_lower_limit' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'atk_range_get_upper_limit' : return type 'gdouble' not supported

var rectangleStruct *gi.Struct
var rectangleStructOnce sync.Once

func rectangleStructSet() {
	rectangleStructOnce.Do(func() {
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
var registryClassStructOnce sync.Once

func registryClassStructSet() {
	registryClassStructOnce.Do(func() {
		registryClassStruct = gi.StructNew("Atk", "RegistryClass")
	})
}

type RegistryClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
}

var relationClassStruct *gi.Struct
var relationClassStructOnce sync.Once

func relationClassStructSet() {
	relationClassStructOnce.Do(func() {
		relationClassStruct = gi.StructNew("Atk", "RelationClass")
	})
}

type RelationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
}

var relationSetClassStruct *gi.Struct
var relationSetClassStructOnce sync.Once

func relationSetClassStructSet() {
	relationSetClassStructOnce.Do(func() {
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
var selectionIfaceStructOnce sync.Once

func selectionIfaceStructSet() {
	selectionIfaceStructOnce.Do(func() {
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
var socketClassStructOnce sync.Once

func socketClassStructSet() {
	socketClassStructOnce.Do(func() {
		socketClassStruct = gi.StructNew("Atk", "SocketClass")
	})
}

type SocketClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'embed' : missing Type
}

var stateSetClassStruct *gi.Struct
var stateSetClassStructOnce sync.Once

func stateSetClassStructSet() {
	stateSetClassStructOnce.Do(func() {
		stateSetClassStruct = gi.StructNew("Atk", "StateSetClass")
	})
}

type StateSetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
}

var streamableContentIfaceStruct *gi.Struct
var streamableContentIfaceStructOnce sync.Once

func streamableContentIfaceStructSet() {
	streamableContentIfaceStructOnce.Do(func() {
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
var tableCellIfaceStructOnce sync.Once

func tableCellIfaceStructSet() {
	tableCellIfaceStructOnce.Do(func() {
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
var tableIfaceStructOnce sync.Once

func tableIfaceStructSet() {
	tableIfaceStructOnce.Do(func() {
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
var textIfaceStructOnce sync.Once

func textIfaceStructSet() {
	textIfaceStructOnce.Do(func() {
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
var textRangeStructOnce sync.Once

func textRangeStructSet() {
	textRangeStructOnce.Do(func() {
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
var textRectangleStructOnce sync.Once

func textRectangleStructSet() {
	textRectangleStructOnce.Do(func() {
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
var utilClassStructOnce sync.Once

func utilClassStructSet() {
	utilClassStructOnce.Do(func() {
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
var valueIfaceStructOnce sync.Once

func valueIfaceStructSet() {
	valueIfaceStructOnce.Do(func() {
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
var windowIfaceStructOnce sync.Once

func windowIfaceStructSet() {
	windowIfaceStructOnce.Do(func() {
		windowIfaceStruct = gi.StructNew("Atk", "WindowIface")
	})
}

type WindowIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
}
