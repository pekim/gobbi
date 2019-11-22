// Code generated - DO NOT EDIT.

package atk

import (
	gi "github.com/pekim/gobbi/internal/gi"
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

func attributeStruct_Set() error {
	var err error
	attributeStruct_Once.Do(func() {
		attributeStruct, err = gi.StructNew("Atk", "Attribute")
	})
	return err
}

type Attribute struct {
	native uintptr
	Name   string
	Value  string
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

func documentIfaceStruct_Set() error {
	var err error
	documentIfaceStruct_Once.Do(func() {
		documentIfaceStruct, err = gi.StructNew("Atk", "DocumentIface")
	})
	return err
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

func editableTextIfaceStruct_Set() error {
	var err error
	editableTextIfaceStruct_Once.Do(func() {
		editableTextIfaceStruct, err = gi.StructNew("Atk", "EditableTextIface")
	})
	return err
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

func gObjectAccessibleClassStruct_Set() error {
	var err error
	gObjectAccessibleClassStruct_Once.Do(func() {
		gObjectAccessibleClassStruct, err = gi.StructNew("Atk", "GObjectAccessibleClass")
	})
	return err
}

type GObjectAccessibleClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'
	// UNSUPPORTED : C value 'pad2' : no Go type for 'Function'
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

func hyperlinkImplIfaceStruct_Set() error {
	var err error
	hyperlinkImplIfaceStruct_Once.Do(func() {
		hyperlinkImplIfaceStruct, err = gi.StructNew("Atk", "HyperlinkImplIface")
	})
	return err
}

type HyperlinkImplIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_hyperlink' : missing Type
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
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_link' : missing Type
	// UNSUPPORTED : C value 'get_n_links' : missing Type
	// UNSUPPORTED : C value 'get_link_index' : missing Type
	// UNSUPPORTED : C value 'link_selected' : missing Type
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
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_image_position' : missing Type
	// UNSUPPORTED : C value 'get_image_description' : missing Type
	// UNSUPPORTED : C value 'get_image_size' : missing Type
	// UNSUPPORTED : C value 'set_image_description' : missing Type
	// UNSUPPORTED : C value 'get_image_locale' : missing Type
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

func miscClassStruct_Set() error {
	var err error
	miscClassStruct_Once.Do(func() {
		miscClassStruct, err = gi.StructNew("Atk", "MiscClass")
	})
	return err
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

func noOpObjectClassStruct_Set() error {
	var err error
	noOpObjectClassStruct_Once.Do(func() {
		noOpObjectClassStruct, err = gi.StructNew("Atk", "NoOpObjectClass")
	})
	return err
}

type NoOpObjectClass struct {
	native      uintptr
	ParentClass *ObjectClass
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
	native      uintptr
	ParentClass *ObjectFactoryClass
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

func objectFactoryClassStruct_Set() error {
	var err error
	objectFactoryClassStruct_Once.Do(func() {
		objectFactoryClassStruct, err = gi.StructNew("Atk", "ObjectFactoryClass")
	})
	return err
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

func plugClassStruct_Set() error {
	var err error
	plugClassStruct_Once.Do(func() {
		plugClassStruct, err = gi.StructNew("Atk", "PlugClass")
	})
	return err
}

type PlugClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'get_object_id' : missing Type
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
	native       uintptr
	PropertyName string
	// UNSUPPORTED : C value 'old_value' : no Go type for 'GObject.Value'
	// UNSUPPORTED : C value 'new_value' : no Go type for 'GObject.Value'
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

// UNSUPPORTED : C value 'atk_range_new' : parameter 'lower_limit' of type 'gdouble' not supported

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

// UNSUPPORTED : C value 'atk_range_get_lower_limit' : return type 'gdouble' not supported

// UNSUPPORTED : C value 'atk_range_get_upper_limit' : return type 'gdouble' not supported

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
	X      int32
	Y      int32
	Width  int32
	Height int32
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
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
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
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
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
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'pad1' : no Go type for 'Function'
	// UNSUPPORTED : C value 'pad2' : no Go type for 'Function'
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

func socketClassStruct_Set() error {
	var err error
	socketClassStruct_Once.Do(func() {
		socketClassStruct, err = gi.StructNew("Atk", "SocketClass")
	})
	return err
}

type SocketClass struct {
	native      uintptr
	ParentClass *ObjectClass
	// UNSUPPORTED : C value 'embed' : missing Type
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
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.ObjectClass'
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

func tableCellIfaceStruct_Set() error {
	var err error
	tableCellIfaceStruct_Once.Do(func() {
		tableCellIfaceStruct, err = gi.StructNew("Atk", "TableCellIface")
	})
	return err
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

func tableIfaceStruct_Set() error {
	var err error
	tableIfaceStruct_Once.Do(func() {
		tableIfaceStruct, err = gi.StructNew("Atk", "TableIface")
	})
	return err
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

func textIfaceStruct_Set() error {
	var err error
	textIfaceStruct_Once.Do(func() {
		textIfaceStruct, err = gi.StructNew("Atk", "TextIface")
	})
	return err
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

func textRangeStruct_Set() error {
	var err error
	textRangeStruct_Once.Do(func() {
		textRangeStruct, err = gi.StructNew("Atk", "TextRange")
	})
	return err
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

func textRectangleStruct_Set() error {
	var err error
	textRectangleStruct_Once.Do(func() {
		textRectangleStruct, err = gi.StructNew("Atk", "TextRectangle")
	})
	return err
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

func utilClassStruct_Set() error {
	var err error
	utilClassStruct_Once.Do(func() {
		utilClassStruct, err = gi.StructNew("Atk", "UtilClass")
	})
	return err
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

func valueIfaceStruct_Set() error {
	var err error
	valueIfaceStruct_Once.Do(func() {
		valueIfaceStruct, err = gi.StructNew("Atk", "ValueIface")
	})
	return err
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

func windowIfaceStruct_Set() error {
	var err error
	windowIfaceStruct_Once.Do(func() {
		windowIfaceStruct, err = gi.StructNew("Atk", "WindowIface")
	})
	return err
}

type WindowIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'
}
