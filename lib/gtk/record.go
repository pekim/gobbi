// This is a generated file - DO NOT EDIT

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// AboutDialogClass is a wrapper around the C record GtkAboutDialogClass.
type AboutDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate_link
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// AboutDialogPrivate is a wrapper around the C record GtkAboutDialogPrivate.
type AboutDialogPrivate struct {
	native unsafe.Pointer
}

// AccelGroupClass is a wrapper around the C record GtkAccelGroupClass.
type AccelGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for accel_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// AccelGroupEntry is a wrapper around the C record GtkAccelGroupEntry.
type AccelGroupEntry struct {
	native unsafe.Pointer
	// key : record
	// closure : record
	AccelPathQuark glib.Quark
}

// AccelGroupPrivate is a wrapper around the C record GtkAccelGroupPrivate.
type AccelGroupPrivate struct {
	native unsafe.Pointer
}

// AccelKey is a wrapper around the C record GtkAccelKey.
type AccelKey struct {
	native    unsafe.Pointer
	AccelKey  uint32
	AccelMods gdk.ModifierType
	// Bitfield not supported : 16 accel_flags
}

// AccelLabelClass is a wrapper around the C record GtkAccelLabelClass.
type AccelLabelClass struct {
	native unsafe.Pointer
	// parent_class : record
	SignalQuote1   string
	SignalQuote2   string
	ModNameShift   string
	ModNameControl string
	ModNameAlt     string
	ModSeparator   string
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// AccelLabelPrivate is a wrapper around the C record GtkAccelLabelPrivate.
type AccelLabelPrivate struct {
	native unsafe.Pointer
}

// AccelMapClass is a wrapper around the C record GtkAccelMapClass.
type AccelMapClass struct {
	native unsafe.Pointer
}

// AccessibleClass is a wrapper around the C record GtkAccessibleClass.
type AccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for connect_widget_destroyed
	// no type for widget_set
	// no type for widget_unset
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// AccessiblePrivate is a wrapper around the C record GtkAccessiblePrivate.
type AccessiblePrivate struct {
	native unsafe.Pointer
}

// ActionBarClass is a wrapper around the C record GtkActionBarClass.
type ActionBarClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ActionBarPrivate is a wrapper around the C record GtkActionBarPrivate.
type ActionBarPrivate struct {
	native unsafe.Pointer
}

// ActionClass is a wrapper around the C record GtkActionClass.
type ActionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// Private : menu_item_type
	// Private : toolbar_item_type
	// no type for create_menu_item
	// no type for create_tool_item
	// no type for connect_proxy
	// no type for disconnect_proxy
	// no type for create_menu
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ActionEntry is a wrapper around the C record GtkActionEntry.
type ActionEntry struct {
	native      unsafe.Pointer
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	// callback : no type generator for GObject.Callback, GCallback
}

// ActionGroupClass is a wrapper around the C record GtkActionGroupClass.
type ActionGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_action
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ActionGroupPrivate is a wrapper around the C record GtkActionGroupPrivate.
type ActionGroupPrivate struct {
	native unsafe.Pointer
}

// ActionPrivate is a wrapper around the C record GtkActionPrivate.
type ActionPrivate struct {
	native unsafe.Pointer
}

// ActionableInterface is a wrapper around the C record GtkActionableInterface.
type ActionableInterface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for get_action_name
	// no type for set_action_name
	// no type for get_action_target_value
	// no type for set_action_target_value
}

// AdjustmentClass is a wrapper around the C record GtkAdjustmentClass.
type AdjustmentClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for value_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// AdjustmentPrivate is a wrapper around the C record GtkAdjustmentPrivate.
type AdjustmentPrivate struct {
	native unsafe.Pointer
}

// AlignmentClass is a wrapper around the C record GtkAlignmentClass.
type AlignmentClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// AlignmentPrivate is a wrapper around the C record GtkAlignmentPrivate.
type AlignmentPrivate struct {
	native unsafe.Pointer
}

// AppChooserButtonClass is a wrapper around the C record GtkAppChooserButtonClass.
type AppChooserButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for custom_item_activated
	// Private : padding
}

// AppChooserButtonPrivate is a wrapper around the C record GtkAppChooserButtonPrivate.
type AppChooserButtonPrivate struct {
	native unsafe.Pointer
}

// AppChooserDialogClass is a wrapper around the C record GtkAppChooserDialogClass.
type AppChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Private : padding
}

// AppChooserDialogPrivate is a wrapper around the C record GtkAppChooserDialogPrivate.
type AppChooserDialogPrivate struct {
	native unsafe.Pointer
}

// AppChooserWidgetClass is a wrapper around the C record GtkAppChooserWidgetClass.
type AppChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for application_selected
	// no type for application_activated
	// no type for populate_popup
	// Private : padding
}

// AppChooserWidgetPrivate is a wrapper around the C record GtkAppChooserWidgetPrivate.
type AppChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// ApplicationClass is a wrapper around the C record GtkApplicationClass.
type ApplicationClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for window_added
	// no type for window_removed
	// Private : padding
}

// ApplicationPrivate is a wrapper around the C record GtkApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

// ApplicationWindowClass is a wrapper around the C record GtkApplicationWindowClass.
type ApplicationWindowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Private : padding
}

// ApplicationWindowPrivate is a wrapper around the C record GtkApplicationWindowPrivate.
type ApplicationWindowPrivate struct {
	native unsafe.Pointer
}

// ArrowAccessibleClass is a wrapper around the C record GtkArrowAccessibleClass.
type ArrowAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ArrowAccessiblePrivate is a wrapper around the C record GtkArrowAccessiblePrivate.
type ArrowAccessiblePrivate struct {
	native unsafe.Pointer
}

// ArrowClass is a wrapper around the C record GtkArrowClass.
type ArrowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ArrowPrivate is a wrapper around the C record GtkArrowPrivate.
type ArrowPrivate struct {
	native unsafe.Pointer
}

// AspectFrameClass is a wrapper around the C record GtkAspectFrameClass.
type AspectFrameClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// AspectFramePrivate is a wrapper around the C record GtkAspectFramePrivate.
type AspectFramePrivate struct {
	native unsafe.Pointer
}

// AssistantClass is a wrapper around the C record GtkAssistantClass.
type AssistantClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for prepare
	// no type for apply
	// no type for close
	// no type for cancel
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
}

// AssistantPrivate is a wrapper around the C record GtkAssistantPrivate.
type AssistantPrivate struct {
	native unsafe.Pointer
}

// BinClass is a wrapper around the C record GtkBinClass.
type BinClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// BinPrivate is a wrapper around the C record GtkBinPrivate.
type BinPrivate struct {
	native unsafe.Pointer
}

// BindingArg is a wrapper around the C record GtkBindingArg.
type BindingArg struct {
	native  unsafe.Pointer
	ArgType gobject.Type
}

// BindingEntry is a wrapper around the C record GtkBindingEntry.
type BindingEntry struct {
	native    unsafe.Pointer
	Keyval    uint32
	Modifiers gdk.ModifierType
	// binding_set : record
	// Bitfield not supported :  1 destroyed
	// Bitfield not supported :  1 in_emission
	// Bitfield not supported :  1 marks_unbound
	// set_next : record
	// hash_next : record
	// signals : record
}

// BindingSet is a wrapper around the C record GtkBindingSet.
type BindingSet struct {
	native   unsafe.Pointer
	SetName  string
	Priority int32
	// widget_path_pspecs : record
	// widget_class_pspecs : record
	// class_branch_pspecs : record
	// entries : record
	// current : record
	// Bitfield not supported :  1 parsed
}

// BindingSignal is a wrapper around the C record GtkBindingSignal.
type BindingSignal struct {
	native unsafe.Pointer
	// next : record
	SignalName string
	NArgs      uint32
	// no type for args
}

// BooleanCellAccessibleClass is a wrapper around the C record GtkBooleanCellAccessibleClass.
type BooleanCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// BooleanCellAccessiblePrivate is a wrapper around the C record GtkBooleanCellAccessiblePrivate.
type BooleanCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// Border is a wrapper around the C record GtkBorder.
type Border struct {
	native unsafe.Pointer
	Left   int16
	Right  int16
	Top    int16
	Bottom int16
}

// BoxClass is a wrapper around the C record GtkBoxClass.
type BoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// BoxPrivate is a wrapper around the C record GtkBoxPrivate.
type BoxPrivate struct {
	native unsafe.Pointer
}

// BuildableIface is a wrapper around the C record GtkBuildableIface.
type BuildableIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for set_name
	// no type for get_name
	// no type for add_child
	// no type for set_buildable_property
	// no type for construct_child
	// no type for custom_tag_start
	// no type for custom_tag_end
	// no type for custom_finished
	// no type for parser_finished
	// no type for get_internal_child
}

// BuilderClass is a wrapper around the C record GtkBuilderClass.
type BuilderClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_type_from_name
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// BuilderPrivate is a wrapper around the C record GtkBuilderPrivate.
type BuilderPrivate struct {
	native unsafe.Pointer
}

// ButtonAccessibleClass is a wrapper around the C record GtkButtonAccessibleClass.
type ButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ButtonAccessiblePrivate is a wrapper around the C record GtkButtonAccessiblePrivate.
type ButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ButtonBoxClass is a wrapper around the C record GtkButtonBoxClass.
type ButtonBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ButtonBoxPrivate is a wrapper around the C record GtkButtonBoxPrivate.
type ButtonBoxPrivate struct {
	native unsafe.Pointer
}

// ButtonClass is a wrapper around the C record GtkButtonClass.
type ButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for pressed
	// no type for released
	// no type for clicked
	// no type for enter
	// no type for leave
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ButtonPrivate is a wrapper around the C record GtkButtonPrivate.
type ButtonPrivate struct {
	native unsafe.Pointer
}

// CalendarClass is a wrapper around the C record GtkCalendarClass.
type CalendarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for month_changed
	// no type for day_selected
	// no type for day_selected_double_click
	// no type for prev_month
	// no type for next_month
	// no type for prev_year
	// no type for next_year
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CalendarPrivate is a wrapper around the C record GtkCalendarPrivate.
type CalendarPrivate struct {
	native unsafe.Pointer
}

// CellAccessibleClass is a wrapper around the C record GtkCellAccessibleClass.
type CellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for update_cache
}

// CellAccessibleParentIface is a wrapper around the C record GtkCellAccessibleParentIface.
type CellAccessibleParentIface struct {
	native unsafe.Pointer
	// parent : record
	// no type for get_cell_extents
	// no type for get_cell_area
	// no type for grab_focus
	// no type for get_child_index
	// no type for get_renderer_state
	// no type for expand_collapse
	// no type for activate
	// no type for edit
	// no type for update_relationset
}

// CellAccessiblePrivate is a wrapper around the C record GtkCellAccessiblePrivate.
type CellAccessiblePrivate struct {
	native unsafe.Pointer
}

// CellAreaBoxClass is a wrapper around the C record GtkCellAreaBoxClass.
type CellAreaBoxClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellAreaBoxPrivate is a wrapper around the C record GtkCellAreaBoxPrivate.
type CellAreaBoxPrivate struct {
	native unsafe.Pointer
}

// CellAreaClass is a wrapper around the C record GtkCellAreaClass.
type CellAreaClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for add
	// no type for remove
	// no type for foreach
	// no type for foreach_alloc
	// no type for event
	// no type for render
	// no type for apply_attributes
	// no type for create_context
	// no type for copy_context
	// no type for get_request_mode
	// no type for get_preferred_width
	// no type for get_preferred_height_for_width
	// no type for get_preferred_height
	// no type for get_preferred_width_for_height
	// no type for set_cell_property
	// no type for get_cell_property
	// no type for focus
	// no type for is_activatable
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// CellAreaContextClass is a wrapper around the C record GtkCellAreaContextClass.
type CellAreaContextClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for allocate
	// no type for reset
	// no type for get_preferred_height_for_width
	// no type for get_preferred_width_for_height
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
}

// CellAreaContextPrivate is a wrapper around the C record GtkCellAreaContextPrivate.
type CellAreaContextPrivate struct {
	native unsafe.Pointer
}

// CellAreaPrivate is a wrapper around the C record GtkCellAreaPrivate.
type CellAreaPrivate struct {
	native unsafe.Pointer
}

// CellEditableIface is a wrapper around the C record GtkCellEditableIface.
type CellEditableIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for editing_done
	// no type for remove_widget
	// no type for start_editing
}

// CellLayoutIface is a wrapper around the C record GtkCellLayoutIface.
type CellLayoutIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for pack_start
	// no type for pack_end
	// no type for clear
	// no type for add_attribute
	// no type for set_cell_data_func
	// no type for clear_attributes
	// no type for reorder
	// no type for get_cells
	// no type for get_area
}

// CellRendererAccelClass is a wrapper around the C record GtkCellRendererAccelClass.
type CellRendererAccelClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for accel_edited
	// no type for accel_cleared
	// no type for _gtk_reserved0
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellRendererAccelPrivate is a wrapper around the C record GtkCellRendererAccelPrivate.
type CellRendererAccelPrivate struct {
	native unsafe.Pointer
}

// CellRendererClass is a wrapper around the C record GtkCellRendererClass.
type CellRendererClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for get_request_mode
	// no type for get_preferred_width
	// no type for get_preferred_height_for_width
	// no type for get_preferred_height
	// no type for get_preferred_width_for_height
	// no type for get_aligned_area
	// no type for get_size
	// no type for render
	// no type for activate
	// no type for start_editing
	// no type for editing_canceled
	// no type for editing_started
	// Private : priv
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellRendererClassPrivate is a wrapper around the C record GtkCellRendererClassPrivate.
type CellRendererClassPrivate struct {
	native unsafe.Pointer
}

// CellRendererComboClass is a wrapper around the C record GtkCellRendererComboClass.
type CellRendererComboClass struct {
	native unsafe.Pointer
	// parent : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellRendererComboPrivate is a wrapper around the C record GtkCellRendererComboPrivate.
type CellRendererComboPrivate struct {
	native unsafe.Pointer
}

// CellRendererPixbufClass is a wrapper around the C record GtkCellRendererPixbufClass.
type CellRendererPixbufClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellRendererPixbufPrivate is a wrapper around the C record GtkCellRendererPixbufPrivate.
type CellRendererPixbufPrivate struct {
	native unsafe.Pointer
}

// CellRendererPrivate is a wrapper around the C record GtkCellRendererPrivate.
type CellRendererPrivate struct {
	native unsafe.Pointer
}

// CellRendererProgressClass is a wrapper around the C record GtkCellRendererProgressClass.
type CellRendererProgressClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellRendererProgressPrivate is a wrapper around the C record GtkCellRendererProgressPrivate.
type CellRendererProgressPrivate struct {
	native unsafe.Pointer
}

// CellRendererSpinClass is a wrapper around the C record GtkCellRendererSpinClass.
type CellRendererSpinClass struct {
	native unsafe.Pointer
	// parent : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellRendererSpinPrivate is a wrapper around the C record GtkCellRendererSpinPrivate.
type CellRendererSpinPrivate struct {
	native unsafe.Pointer
}

// CellRendererSpinnerClass is a wrapper around the C record GtkCellRendererSpinnerClass.
type CellRendererSpinnerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellRendererSpinnerPrivate is a wrapper around the C record GtkCellRendererSpinnerPrivate.
type CellRendererSpinnerPrivate struct {
	native unsafe.Pointer
}

// CellRendererTextClass is a wrapper around the C record GtkCellRendererTextClass.
type CellRendererTextClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for edited
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellRendererTextPrivate is a wrapper around the C record GtkCellRendererTextPrivate.
type CellRendererTextPrivate struct {
	native unsafe.Pointer
}

// CellRendererToggleClass is a wrapper around the C record GtkCellRendererToggleClass.
type CellRendererToggleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellRendererTogglePrivate is a wrapper around the C record GtkCellRendererTogglePrivate.
type CellRendererTogglePrivate struct {
	native unsafe.Pointer
}

// CellViewClass is a wrapper around the C record GtkCellViewClass.
type CellViewClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CellViewPrivate is a wrapper around the C record GtkCellViewPrivate.
type CellViewPrivate struct {
	native unsafe.Pointer
}

// CheckButtonClass is a wrapper around the C record GtkCheckButtonClass.
type CheckButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for draw_indicator
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CheckMenuItemAccessibleClass is a wrapper around the C record GtkCheckMenuItemAccessibleClass.
type CheckMenuItemAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// CheckMenuItemAccessiblePrivate is a wrapper around the C record GtkCheckMenuItemAccessiblePrivate.
type CheckMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

// CheckMenuItemClass is a wrapper around the C record GtkCheckMenuItemClass.
type CheckMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for draw_indicator
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CheckMenuItemPrivate is a wrapper around the C record GtkCheckMenuItemPrivate.
type CheckMenuItemPrivate struct {
	native unsafe.Pointer
}

// ColorButtonClass is a wrapper around the C record GtkColorButtonClass.
type ColorButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for color_set
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ColorButtonPrivate is a wrapper around the C record GtkColorButtonPrivate.
type ColorButtonPrivate struct {
	native unsafe.Pointer
}

// ColorChooserDialogClass is a wrapper around the C record GtkColorChooserDialogClass.
type ColorChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ColorChooserDialogPrivate is a wrapper around the C record GtkColorChooserDialogPrivate.
type ColorChooserDialogPrivate struct {
	native unsafe.Pointer
}

// ColorChooserInterface is a wrapper around the C record GtkColorChooserInterface.
type ColorChooserInterface struct {
	native unsafe.Pointer
	// base_interface : record
	// no type for get_rgba
	// no type for set_rgba
	// no type for add_palette
	// no type for color_activated
	// no type for padding
}

// ColorChooserWidgetClass is a wrapper around the C record GtkColorChooserWidgetClass.
type ColorChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// ColorChooserWidgetPrivate is a wrapper around the C record GtkColorChooserWidgetPrivate.
type ColorChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// ColorSelectionClass is a wrapper around the C record GtkColorSelectionClass.
type ColorSelectionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for color_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ColorSelectionDialogClass is a wrapper around the C record GtkColorSelectionDialogClass.
type ColorSelectionDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ColorSelectionDialogPrivate is a wrapper around the C record GtkColorSelectionDialogPrivate.
type ColorSelectionDialogPrivate struct {
	native unsafe.Pointer
}

// ColorSelectionPrivate is a wrapper around the C record GtkColorSelectionPrivate.
type ColorSelectionPrivate struct {
	native unsafe.Pointer
}

// ComboBoxAccessibleClass is a wrapper around the C record GtkComboBoxAccessibleClass.
type ComboBoxAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ComboBoxAccessiblePrivate is a wrapper around the C record GtkComboBoxAccessiblePrivate.
type ComboBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

// ComboBoxClass is a wrapper around the C record GtkComboBoxClass.
type ComboBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for format_entry_text
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
}

// ComboBoxPrivate is a wrapper around the C record GtkComboBoxPrivate.
type ComboBoxPrivate struct {
	native unsafe.Pointer
}

// ComboBoxTextClass is a wrapper around the C record GtkComboBoxTextClass.
type ComboBoxTextClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ComboBoxTextPrivate is a wrapper around the C record GtkComboBoxTextPrivate.
type ComboBoxTextPrivate struct {
	native unsafe.Pointer
}

// ContainerAccessibleClass is a wrapper around the C record GtkContainerAccessibleClass.
type ContainerAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for add_gtk
	// no type for remove_gtk
}

// ContainerAccessiblePrivate is a wrapper around the C record GtkContainerAccessiblePrivate.
type ContainerAccessiblePrivate struct {
	native unsafe.Pointer
}

// ContainerCellAccessibleClass is a wrapper around the C record GtkContainerCellAccessibleClass.
type ContainerCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ContainerCellAccessiblePrivate is a wrapper around the C record GtkContainerCellAccessiblePrivate.
type ContainerCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ContainerClass is a wrapper around the C record GtkContainerClass.
type ContainerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for add
	// no type for remove
	// no type for check_resize
	// no type for forall
	// no type for set_focus_child
	// no type for child_type
	// no type for composite_name
	// no type for set_child_property
	// no type for get_child_property
	// no type for get_path_for_child
	// Private : _handle_border_width
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// ContainerPrivate is a wrapper around the C record GtkContainerPrivate.
type ContainerPrivate struct {
	native unsafe.Pointer
}

// CssProviderClass is a wrapper around the C record GtkCssProviderClass.
type CssProviderClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for parsing_error
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// CssProviderPrivate is a wrapper around the C record GtkCssProviderPrivate.
type CssProviderPrivate struct {
	native unsafe.Pointer
}

// DialogClass is a wrapper around the C record GtkDialogClass.
type DialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for response
	// no type for close
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// DialogPrivate is a wrapper around the C record GtkDialogPrivate.
type DialogPrivate struct {
	native unsafe.Pointer
}

// DrawingAreaClass is a wrapper around the C record GtkDrawingAreaClass.
type DrawingAreaClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// EditableInterface is a wrapper around the C record GtkEditableInterface.
type EditableInterface struct {
	native unsafe.Pointer
	// base_iface : record
	// no type for insert_text
	// no type for delete_text
	// no type for changed
	// no type for do_insert_text
	// no type for do_delete_text
	// no type for get_chars
	// no type for set_selection_bounds
	// no type for get_selection_bounds
	// no type for set_position
	// no type for get_position
}

// EntryAccessibleClass is a wrapper around the C record GtkEntryAccessibleClass.
type EntryAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// EntryAccessiblePrivate is a wrapper around the C record GtkEntryAccessiblePrivate.
type EntryAccessiblePrivate struct {
	native unsafe.Pointer
}

// EntryBufferClass is a wrapper around the C record GtkEntryBufferClass.
type EntryBufferClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for inserted_text
	// no type for deleted_text
	// no type for get_text
	// no type for get_length
	// no type for insert_text
	// no type for delete_text
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// EntryBufferPrivate is a wrapper around the C record GtkEntryBufferPrivate.
type EntryBufferPrivate struct {
	native unsafe.Pointer
}

// EntryClass is a wrapper around the C record GtkEntryClass.
type EntryClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for populate_popup
	// no type for activate
	// no type for move_cursor
	// no type for insert_at_cursor
	// no type for delete_from_cursor
	// no type for backspace
	// no type for cut_clipboard
	// no type for copy_clipboard
	// no type for paste_clipboard
	// no type for toggle_overwrite
	// no type for get_text_area_size
	// no type for get_frame_size
	// no type for insert_emoji
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
}

// EntryCompletionClass is a wrapper around the C record GtkEntryCompletionClass.
type EntryCompletionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for match_selected
	// no type for action_activated
	// no type for insert_prefix
	// no type for cursor_on_match
	// no type for no_matches
	// no type for _gtk_reserved0
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
}

// EntryCompletionPrivate is a wrapper around the C record GtkEntryCompletionPrivate.
type EntryCompletionPrivate struct {
	native unsafe.Pointer
}

// EntryPrivate is a wrapper around the C record GtkEntryPrivate.
type EntryPrivate struct {
	native unsafe.Pointer
}

// EventBoxClass is a wrapper around the C record GtkEventBoxClass.
type EventBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// EventBoxPrivate is a wrapper around the C record GtkEventBoxPrivate.
type EventBoxPrivate struct {
	native unsafe.Pointer
}

// EventControllerClass is a wrapper around the C record GtkEventControllerClass.
type EventControllerClass struct {
	native unsafe.Pointer
}

// ExpanderAccessibleClass is a wrapper around the C record GtkExpanderAccessibleClass.
type ExpanderAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ExpanderAccessiblePrivate is a wrapper around the C record GtkExpanderAccessiblePrivate.
type ExpanderAccessiblePrivate struct {
	native unsafe.Pointer
}

// ExpanderClass is a wrapper around the C record GtkExpanderClass.
type ExpanderClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ExpanderPrivate is a wrapper around the C record GtkExpanderPrivate.
type ExpanderPrivate struct {
	native unsafe.Pointer
}

// FileChooserButtonClass is a wrapper around the C record GtkFileChooserButtonClass.
type FileChooserButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for file_set
	// Private : __gtk_reserved1
	// Private : __gtk_reserved2
	// Private : __gtk_reserved3
	// Private : __gtk_reserved4
}

// FileChooserButtonPrivate is a wrapper around the C record GtkFileChooserButtonPrivate.
type FileChooserButtonPrivate struct {
	native unsafe.Pointer
}

// FileChooserDialogClass is a wrapper around the C record GtkFileChooserDialogClass.
type FileChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// FileChooserDialogPrivate is a wrapper around the C record GtkFileChooserDialogPrivate.
type FileChooserDialogPrivate struct {
	native unsafe.Pointer
}

// FileChooserWidgetClass is a wrapper around the C record GtkFileChooserWidgetClass.
type FileChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// FileChooserWidgetPrivate is a wrapper around the C record GtkFileChooserWidgetPrivate.
type FileChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// FileFilterInfo is a wrapper around the C record GtkFileFilterInfo.
type FileFilterInfo struct {
	native      unsafe.Pointer
	Contains    FileFilterFlags
	Filename    string
	Uri         string
	DisplayName string
	MimeType    string
}

// FixedChild is a wrapper around the C record GtkFixedChild.
type FixedChild struct {
	native unsafe.Pointer
	// widget : record
	X int32
	Y int32
}

// FixedClass is a wrapper around the C record GtkFixedClass.
type FixedClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// FixedPrivate is a wrapper around the C record GtkFixedPrivate.
type FixedPrivate struct {
	native unsafe.Pointer
}

// FlowBoxAccessibleClass is a wrapper around the C record GtkFlowBoxAccessibleClass.
type FlowBoxAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// FlowBoxAccessiblePrivate is a wrapper around the C record GtkFlowBoxAccessiblePrivate.
type FlowBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

// FlowBoxChildAccessibleClass is a wrapper around the C record GtkFlowBoxChildAccessibleClass.
type FlowBoxChildAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// FlowBoxChildClass is a wrapper around the C record GtkFlowBoxChildClass.
type FlowBoxChildClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
}

// FlowBoxClass is a wrapper around the C record GtkFlowBoxClass.
type FlowBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for child_activated
	// no type for selected_children_changed
	// no type for activate_cursor_child
	// no type for toggle_cursor_child
	// no type for move_cursor
	// no type for select_all
	// no type for unselect_all
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
}

// FontButtonClass is a wrapper around the C record GtkFontButtonClass.
type FontButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for font_set
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// FontButtonPrivate is a wrapper around the C record GtkFontButtonPrivate.
type FontButtonPrivate struct {
	native unsafe.Pointer
}

// FontChooserDialogClass is a wrapper around the C record GtkFontChooserDialogClass.
type FontChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// FontChooserDialogPrivate is a wrapper around the C record GtkFontChooserDialogPrivate.
type FontChooserDialogPrivate struct {
	native unsafe.Pointer
}

// FontChooserIface is a wrapper around the C record GtkFontChooserIface.
type FontChooserIface struct {
	native unsafe.Pointer
	// base_iface : record
	// no type for get_font_family
	// no type for get_font_face
	// no type for get_font_size
	// no type for set_filter_func
	// no type for font_activated
	// no type for set_font_map
	// no type for get_font_map
	// no type for padding
}

// FontChooserWidgetClass is a wrapper around the C record GtkFontChooserWidgetClass.
type FontChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// FontChooserWidgetPrivate is a wrapper around the C record GtkFontChooserWidgetPrivate.
type FontChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// FontSelectionClass is a wrapper around the C record GtkFontSelectionClass.
type FontSelectionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// FontSelectionDialogClass is a wrapper around the C record GtkFontSelectionDialogClass.
type FontSelectionDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// FontSelectionDialogPrivate is a wrapper around the C record GtkFontSelectionDialogPrivate.
type FontSelectionDialogPrivate struct {
	native unsafe.Pointer
}

// FontSelectionPrivate is a wrapper around the C record GtkFontSelectionPrivate.
type FontSelectionPrivate struct {
	native unsafe.Pointer
}

// FrameAccessibleClass is a wrapper around the C record GtkFrameAccessibleClass.
type FrameAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// FrameAccessiblePrivate is a wrapper around the C record GtkFrameAccessiblePrivate.
type FrameAccessiblePrivate struct {
	native unsafe.Pointer
}

// FrameClass is a wrapper around the C record GtkFrameClass.
type FrameClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for compute_child_allocation
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// FramePrivate is a wrapper around the C record GtkFramePrivate.
type FramePrivate struct {
	native unsafe.Pointer
}

// GestureClass is a wrapper around the C record GtkGestureClass.
type GestureClass struct {
	native unsafe.Pointer
}

// GestureDragClass is a wrapper around the C record GtkGestureDragClass.
type GestureDragClass struct {
	native unsafe.Pointer
}

// GestureLongPressClass is a wrapper around the C record GtkGestureLongPressClass.
type GestureLongPressClass struct {
	native unsafe.Pointer
}

// GestureMultiPressClass is a wrapper around the C record GtkGestureMultiPressClass.
type GestureMultiPressClass struct {
	native unsafe.Pointer
}

// GesturePanClass is a wrapper around the C record GtkGesturePanClass.
type GesturePanClass struct {
	native unsafe.Pointer
}

// GestureRotateClass is a wrapper around the C record GtkGestureRotateClass.
type GestureRotateClass struct {
	native unsafe.Pointer
}

// GestureSingleClass is a wrapper around the C record GtkGestureSingleClass.
type GestureSingleClass struct {
	native unsafe.Pointer
}

// GestureSwipeClass is a wrapper around the C record GtkGestureSwipeClass.
type GestureSwipeClass struct {
	native unsafe.Pointer
}

// GestureZoomClass is a wrapper around the C record GtkGestureZoomClass.
type GestureZoomClass struct {
	native unsafe.Pointer
}

// Gradient is a wrapper around the C record GtkGradient.
type Gradient struct {
	native unsafe.Pointer
}

// GridClass is a wrapper around the C record GtkGridClass.
type GridClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// GridPrivate is a wrapper around the C record GtkGridPrivate.
type GridPrivate struct {
	native unsafe.Pointer
}

// HBoxClass is a wrapper around the C record GtkHBoxClass.
type HBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// HButtonBoxClass is a wrapper around the C record GtkHButtonBoxClass.
type HButtonBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// HPanedClass is a wrapper around the C record GtkHPanedClass.
type HPanedClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// HSVClass is a wrapper around the C record GtkHSVClass.
type HSVClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for move
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// HSVPrivate is a wrapper around the C record GtkHSVPrivate.
type HSVPrivate struct {
	native unsafe.Pointer
}

// HScaleClass is a wrapper around the C record GtkHScaleClass.
type HScaleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// HScrollbarClass is a wrapper around the C record GtkHScrollbarClass.
type HScrollbarClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// HSeparatorClass is a wrapper around the C record GtkHSeparatorClass.
type HSeparatorClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// HandleBoxClass is a wrapper around the C record GtkHandleBoxClass.
type HandleBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for child_attached
	// no type for child_detached
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// HandleBoxPrivate is a wrapper around the C record GtkHandleBoxPrivate.
type HandleBoxPrivate struct {
	native unsafe.Pointer
}

// HeaderBarClass is a wrapper around the C record GtkHeaderBarClass.
type HeaderBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// HeaderBarPrivate is a wrapper around the C record GtkHeaderBarPrivate.
type HeaderBarPrivate struct {
	native unsafe.Pointer
}

// IMContextClass is a wrapper around the C record GtkIMContextClass.
type IMContextClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for preedit_start
	// no type for preedit_end
	// no type for preedit_changed
	// no type for commit
	// no type for retrieve_surrounding
	// no type for delete_surrounding
	// no type for set_client_window
	// no type for get_preedit_string
	// no type for filter_keypress
	// no type for focus_in
	// no type for focus_out
	// no type for reset
	// no type for set_cursor_location
	// no type for set_use_preedit
	// no type for set_surrounding
	// no type for get_surrounding
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
}

// IMContextInfo is a wrapper around the C record GtkIMContextInfo.
type IMContextInfo struct {
	native         unsafe.Pointer
	ContextId      string
	ContextName    string
	Domain         string
	DomainDirname  string
	DefaultLocales string
}

// IMContextSimpleClass is a wrapper around the C record GtkIMContextSimpleClass.
type IMContextSimpleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// IMContextSimplePrivate is a wrapper around the C record GtkIMContextSimplePrivate.
type IMContextSimplePrivate struct {
	native unsafe.Pointer
}

// IMMulticontextClass is a wrapper around the C record GtkIMMulticontextClass.
type IMMulticontextClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// IMMulticontextPrivate is a wrapper around the C record GtkIMMulticontextPrivate.
type IMMulticontextPrivate struct {
	native unsafe.Pointer
}

// IconFactoryClass is a wrapper around the C record GtkIconFactoryClass.
type IconFactoryClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// IconFactoryPrivate is a wrapper around the C record GtkIconFactoryPrivate.
type IconFactoryPrivate struct {
	native unsafe.Pointer
}

// IconInfoClass is a wrapper around the C record GtkIconInfoClass.
type IconInfoClass struct {
	native unsafe.Pointer
}

// IconSet is a wrapper around the C record GtkIconSet.
type IconSet struct {
	native unsafe.Pointer
}

// IconSource is a wrapper around the C record GtkIconSource.
type IconSource struct {
	native unsafe.Pointer
}

// IconThemeClass is a wrapper around the C record GtkIconThemeClass.
type IconThemeClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// IconThemePrivate is a wrapper around the C record GtkIconThemePrivate.
type IconThemePrivate struct {
	native unsafe.Pointer
}

// IconViewAccessibleClass is a wrapper around the C record GtkIconViewAccessibleClass.
type IconViewAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// IconViewAccessiblePrivate is a wrapper around the C record GtkIconViewAccessiblePrivate.
type IconViewAccessiblePrivate struct {
	native unsafe.Pointer
}

// IconViewClass is a wrapper around the C record GtkIconViewClass.
type IconViewClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for item_activated
	// no type for selection_changed
	// no type for select_all
	// no type for unselect_all
	// no type for select_cursor_item
	// no type for toggle_cursor_item
	// no type for move_cursor
	// no type for activate_cursor_item
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// IconViewPrivate is a wrapper around the C record GtkIconViewPrivate.
type IconViewPrivate struct {
	native unsafe.Pointer
}

// ImageAccessibleClass is a wrapper around the C record GtkImageAccessibleClass.
type ImageAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ImageAccessiblePrivate is a wrapper around the C record GtkImageAccessiblePrivate.
type ImageAccessiblePrivate struct {
	native unsafe.Pointer
}

// ImageCellAccessibleClass is a wrapper around the C record GtkImageCellAccessibleClass.
type ImageCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ImageCellAccessiblePrivate is a wrapper around the C record GtkImageCellAccessiblePrivate.
type ImageCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// ImageClass is a wrapper around the C record GtkImageClass.
type ImageClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ImageMenuItemClass is a wrapper around the C record GtkImageMenuItemClass.
type ImageMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ImageMenuItemPrivate is a wrapper around the C record GtkImageMenuItemPrivate.
type ImageMenuItemPrivate struct {
	native unsafe.Pointer
}

// ImagePrivate is a wrapper around the C record GtkImagePrivate.
type ImagePrivate struct {
	native unsafe.Pointer
}

// InfoBarClass is a wrapper around the C record GtkInfoBarClass.
type InfoBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for response
	// no type for close
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// InfoBarPrivate is a wrapper around the C record GtkInfoBarPrivate.
type InfoBarPrivate struct {
	native unsafe.Pointer
}

// InvisibleClass is a wrapper around the C record GtkInvisibleClass.
type InvisibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// InvisiblePrivate is a wrapper around the C record GtkInvisiblePrivate.
type InvisiblePrivate struct {
	native unsafe.Pointer
}

// LabelAccessibleClass is a wrapper around the C record GtkLabelAccessibleClass.
type LabelAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// LabelAccessiblePrivate is a wrapper around the C record GtkLabelAccessiblePrivate.
type LabelAccessiblePrivate struct {
	native unsafe.Pointer
}

// LabelClass is a wrapper around the C record GtkLabelClass.
type LabelClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for move_cursor
	// no type for copy_clipboard
	// no type for populate_popup
	// no type for activate_link
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// LabelPrivate is a wrapper around the C record GtkLabelPrivate.
type LabelPrivate struct {
	native unsafe.Pointer
}

// LabelSelectionInfo is a wrapper around the C record GtkLabelSelectionInfo.
type LabelSelectionInfo struct {
	native unsafe.Pointer
}

// LayoutClass is a wrapper around the C record GtkLayoutClass.
type LayoutClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// LayoutPrivate is a wrapper around the C record GtkLayoutPrivate.
type LayoutPrivate struct {
	native unsafe.Pointer
}

// LevelBarAccessibleClass is a wrapper around the C record GtkLevelBarAccessibleClass.
type LevelBarAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// LevelBarAccessiblePrivate is a wrapper around the C record GtkLevelBarAccessiblePrivate.
type LevelBarAccessiblePrivate struct {
	native unsafe.Pointer
}

// LevelBarClass is a wrapper around the C record GtkLevelBarClass.
type LevelBarClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for offset_changed
	// Private : padding
}

// LevelBarPrivate is a wrapper around the C record GtkLevelBarPrivate.
type LevelBarPrivate struct {
	native unsafe.Pointer
}

// LinkButtonAccessibleClass is a wrapper around the C record GtkLinkButtonAccessibleClass.
type LinkButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// LinkButtonAccessiblePrivate is a wrapper around the C record GtkLinkButtonAccessiblePrivate.
type LinkButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// LinkButtonClass is a wrapper around the C record GtkLinkButtonClass.
type LinkButtonClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for activate_link
	// no type for _gtk_padding1
	// no type for _gtk_padding2
	// no type for _gtk_padding3
	// no type for _gtk_padding4
}

// LinkButtonPrivate is a wrapper around the C record GtkLinkButtonPrivate.
type LinkButtonPrivate struct {
	native unsafe.Pointer
}

// ListBoxAccessibleClass is a wrapper around the C record GtkListBoxAccessibleClass.
type ListBoxAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ListBoxAccessiblePrivate is a wrapper around the C record GtkListBoxAccessiblePrivate.
type ListBoxAccessiblePrivate struct {
	native unsafe.Pointer
}

// ListBoxClass is a wrapper around the C record GtkListBoxClass.
type ListBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for row_selected
	// no type for row_activated
	// no type for activate_cursor_row
	// no type for toggle_cursor_row
	// no type for move_cursor
	// no type for selected_rows_changed
	// no type for select_all
	// no type for unselect_all
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
}

// ListBoxRowAccessibleClass is a wrapper around the C record GtkListBoxRowAccessibleClass.
type ListBoxRowAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ListBoxRowClass is a wrapper around the C record GtkListBoxRowClass.
type ListBoxRowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
}

// ListStoreClass is a wrapper around the C record GtkListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ListStorePrivate is a wrapper around the C record GtkListStorePrivate.
type ListStorePrivate struct {
	native unsafe.Pointer
}

// LockButtonAccessibleClass is a wrapper around the C record GtkLockButtonAccessibleClass.
type LockButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// LockButtonAccessiblePrivate is a wrapper around the C record GtkLockButtonAccessiblePrivate.
type LockButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// LockButtonClass is a wrapper around the C record GtkLockButtonClass.
type LockButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for reserved0
	// no type for reserved1
	// no type for reserved2
	// no type for reserved3
	// no type for reserved4
	// no type for reserved5
	// no type for reserved6
	// no type for reserved7
}

// LockButtonPrivate is a wrapper around the C record GtkLockButtonPrivate.
type LockButtonPrivate struct {
	native unsafe.Pointer
}

// MenuAccessibleClass is a wrapper around the C record GtkMenuAccessibleClass.
type MenuAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// MenuAccessiblePrivate is a wrapper around the C record GtkMenuAccessiblePrivate.
type MenuAccessiblePrivate struct {
	native unsafe.Pointer
}

// MenuBarClass is a wrapper around the C record GtkMenuBarClass.
type MenuBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// MenuBarPrivate is a wrapper around the C record GtkMenuBarPrivate.
type MenuBarPrivate struct {
	native unsafe.Pointer
}

// MenuButtonAccessibleClass is a wrapper around the C record GtkMenuButtonAccessibleClass.
type MenuButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// MenuButtonAccessiblePrivate is a wrapper around the C record GtkMenuButtonAccessiblePrivate.
type MenuButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// MenuButtonClass is a wrapper around the C record GtkMenuButtonClass.
type MenuButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// MenuButtonPrivate is a wrapper around the C record GtkMenuButtonPrivate.
type MenuButtonPrivate struct {
	native unsafe.Pointer
}

// MenuClass is a wrapper around the C record GtkMenuClass.
type MenuClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// MenuItemAccessibleClass is a wrapper around the C record GtkMenuItemAccessibleClass.
type MenuItemAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// MenuItemAccessiblePrivate is a wrapper around the C record GtkMenuItemAccessiblePrivate.
type MenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

// MenuItemClass is a wrapper around the C record GtkMenuItemClass.
type MenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Bitfield not supported :  1 hide_on_activate
	// no type for activate
	// no type for activate_item
	// no type for toggle_size_request
	// no type for toggle_size_allocate
	// no type for set_label
	// no type for get_label
	// no type for _select
	// no type for deselect
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// MenuItemPrivate is a wrapper around the C record GtkMenuItemPrivate.
type MenuItemPrivate struct {
	native unsafe.Pointer
}

// MenuPrivate is a wrapper around the C record GtkMenuPrivate.
type MenuPrivate struct {
	native unsafe.Pointer
}

// MenuShellAccessibleClass is a wrapper around the C record GtkMenuShellAccessibleClass.
type MenuShellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// MenuShellAccessiblePrivate is a wrapper around the C record GtkMenuShellAccessiblePrivate.
type MenuShellAccessiblePrivate struct {
	native unsafe.Pointer
}

// MenuShellClass is a wrapper around the C record GtkMenuShellClass.
type MenuShellClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Bitfield not supported :  1 submenu_placement
	// no type for deactivate
	// no type for selection_done
	// no type for move_current
	// no type for activate_current
	// no type for cancel
	// no type for select_item
	// no type for insert
	// no type for get_popup_delay
	// no type for move_selected
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// MenuShellPrivate is a wrapper around the C record GtkMenuShellPrivate.
type MenuShellPrivate struct {
	native unsafe.Pointer
}

// MenuToolButtonClass is a wrapper around the C record GtkMenuToolButtonClass.
type MenuToolButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for show_menu
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// MenuToolButtonPrivate is a wrapper around the C record GtkMenuToolButtonPrivate.
type MenuToolButtonPrivate struct {
	native unsafe.Pointer
}

// MessageDialogClass is a wrapper around the C record GtkMessageDialogClass.
type MessageDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// MessageDialogPrivate is a wrapper around the C record GtkMessageDialogPrivate.
type MessageDialogPrivate struct {
	native unsafe.Pointer
}

// MiscClass is a wrapper around the C record GtkMiscClass.
type MiscClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// MiscPrivate is a wrapper around the C record GtkMiscPrivate.
type MiscPrivate struct {
	native unsafe.Pointer
}

// MountOperationClass is a wrapper around the C record GtkMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// MountOperationPrivate is a wrapper around the C record GtkMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

// NotebookAccessibleClass is a wrapper around the C record GtkNotebookAccessibleClass.
type NotebookAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// NotebookAccessiblePrivate is a wrapper around the C record GtkNotebookAccessiblePrivate.
type NotebookAccessiblePrivate struct {
	native unsafe.Pointer
}

// NotebookClass is a wrapper around the C record GtkNotebookClass.
type NotebookClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for switch_page
	// no type for select_page
	// no type for focus_tab
	// no type for change_current_page
	// no type for move_focus_out
	// no type for reorder_tab
	// no type for insert_page
	// no type for create_window
	// no type for page_reordered
	// no type for page_removed
	// no type for page_added
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// NotebookPageAccessibleClass is a wrapper around the C record GtkNotebookPageAccessibleClass.
type NotebookPageAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// NotebookPageAccessiblePrivate is a wrapper around the C record GtkNotebookPageAccessiblePrivate.
type NotebookPageAccessiblePrivate struct {
	native unsafe.Pointer
}

// NotebookPrivate is a wrapper around the C record GtkNotebookPrivate.
type NotebookPrivate struct {
	native unsafe.Pointer
}

// NumerableIconClass is a wrapper around the C record GtkNumerableIconClass.
type NumerableIconClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for padding
}

// NumerableIconPrivate is a wrapper around the C record GtkNumerableIconPrivate.
type NumerableIconPrivate struct {
	native unsafe.Pointer
}

// OffscreenWindowClass is a wrapper around the C record GtkOffscreenWindowClass.
type OffscreenWindowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// OrientableIface is a wrapper around the C record GtkOrientableIface.
type OrientableIface struct {
	native unsafe.Pointer
	// base_iface : record
}

// OverlayClass is a wrapper around the C record GtkOverlayClass.
type OverlayClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for get_child_position
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// OverlayPrivate is a wrapper around the C record GtkOverlayPrivate.
type OverlayPrivate struct {
	native unsafe.Pointer
}

// PageRange is a wrapper around the C record GtkPageRange.
type PageRange struct {
	native unsafe.Pointer
	Start  int32
	End    int32
}

// PanedAccessibleClass is a wrapper around the C record GtkPanedAccessibleClass.
type PanedAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// PanedAccessiblePrivate is a wrapper around the C record GtkPanedAccessiblePrivate.
type PanedAccessiblePrivate struct {
	native unsafe.Pointer
}

// PanedClass is a wrapper around the C record GtkPanedClass.
type PanedClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for cycle_child_focus
	// no type for toggle_handle_focus
	// no type for move_handle
	// no type for cycle_handle_focus
	// no type for accept_position
	// no type for cancel_position
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// PanedPrivate is a wrapper around the C record GtkPanedPrivate.
type PanedPrivate struct {
	native unsafe.Pointer
}

// PaperSize is a wrapper around the C record GtkPaperSize.
type PaperSize struct {
	native unsafe.Pointer
}

// PlacesSidebarClass is a wrapper around the C record GtkPlacesSidebarClass.
type PlacesSidebarClass struct {
	native unsafe.Pointer
}

// PlugClass is a wrapper around the C record GtkPlugClass.
type PlugClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for embedded
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// PlugPrivate is a wrapper around the C record GtkPlugPrivate.
type PlugPrivate struct {
	native unsafe.Pointer
}

// PopoverAccessibleClass is a wrapper around the C record GtkPopoverAccessibleClass.
type PopoverAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// PopoverClass is a wrapper around the C record GtkPopoverClass.
type PopoverClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for closed
	// Private : reserved
}

// PopoverMenuClass is a wrapper around the C record GtkPopoverMenuClass.
type PopoverMenuClass struct {
	native unsafe.Pointer
	// parent_class : record
	// Private : reserved
}

// PopoverPrivate is a wrapper around the C record GtkPopoverPrivate.
type PopoverPrivate struct {
	native unsafe.Pointer
}

// PrintOperationClass is a wrapper around the C record GtkPrintOperationClass.
type PrintOperationClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for done
	// no type for begin_print
	// no type for paginate
	// no type for request_page_setup
	// no type for draw_page
	// no type for end_print
	// no type for status_changed
	// no type for create_custom_widget
	// no type for custom_widget_apply
	// no type for preview
	// no type for update_custom_widget
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// PrintOperationPreviewIface is a wrapper around the C record GtkPrintOperationPreviewIface.
type PrintOperationPreviewIface struct {
	native unsafe.Pointer
	// g_iface : record
	// no type for ready
	// no type for got_page_size
	// no type for render_page
	// no type for is_selected
	// no type for end_preview
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// PrintOperationPrivate is a wrapper around the C record GtkPrintOperationPrivate.
type PrintOperationPrivate struct {
	native unsafe.Pointer
}

// ProgressBarAccessibleClass is a wrapper around the C record GtkProgressBarAccessibleClass.
type ProgressBarAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ProgressBarAccessiblePrivate is a wrapper around the C record GtkProgressBarAccessiblePrivate.
type ProgressBarAccessiblePrivate struct {
	native unsafe.Pointer
}

// ProgressBarClass is a wrapper around the C record GtkProgressBarClass.
type ProgressBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ProgressBarPrivate is a wrapper around the C record GtkProgressBarPrivate.
type ProgressBarPrivate struct {
	native unsafe.Pointer
}

// RadioActionClass is a wrapper around the C record GtkRadioActionClass.
type RadioActionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// RadioActionEntry is a wrapper around the C record GtkRadioActionEntry.
type RadioActionEntry struct {
	native      unsafe.Pointer
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	Value       int32
}

// RadioActionPrivate is a wrapper around the C record GtkRadioActionPrivate.
type RadioActionPrivate struct {
	native unsafe.Pointer
}

// RadioButtonAccessibleClass is a wrapper around the C record GtkRadioButtonAccessibleClass.
type RadioButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// RadioButtonAccessiblePrivate is a wrapper around the C record GtkRadioButtonAccessiblePrivate.
type RadioButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// RadioButtonClass is a wrapper around the C record GtkRadioButtonClass.
type RadioButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for group_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// RadioButtonPrivate is a wrapper around the C record GtkRadioButtonPrivate.
type RadioButtonPrivate struct {
	native unsafe.Pointer
}

// RadioMenuItemAccessibleClass is a wrapper around the C record GtkRadioMenuItemAccessibleClass.
type RadioMenuItemAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// RadioMenuItemAccessiblePrivate is a wrapper around the C record GtkRadioMenuItemAccessiblePrivate.
type RadioMenuItemAccessiblePrivate struct {
	native unsafe.Pointer
}

// RadioMenuItemClass is a wrapper around the C record GtkRadioMenuItemClass.
type RadioMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for group_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// RadioMenuItemPrivate is a wrapper around the C record GtkRadioMenuItemPrivate.
type RadioMenuItemPrivate struct {
	native unsafe.Pointer
}

// RadioToolButtonClass is a wrapper around the C record GtkRadioToolButtonClass.
type RadioToolButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// RangeAccessibleClass is a wrapper around the C record GtkRangeAccessibleClass.
type RangeAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// RangeAccessiblePrivate is a wrapper around the C record GtkRangeAccessiblePrivate.
type RangeAccessiblePrivate struct {
	native unsafe.Pointer
}

// RangeClass is a wrapper around the C record GtkRangeClass.
type RangeClass struct {
	native unsafe.Pointer
	// parent_class : record
	SliderDetail  string
	StepperDetail string
	// no type for value_changed
	// no type for adjust_bounds
	// no type for move_slider
	// no type for get_range_border
	// no type for change_value
	// no type for get_range_size_request
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
}

// RangePrivate is a wrapper around the C record GtkRangePrivate.
type RangePrivate struct {
	native unsafe.Pointer
}

// RcContext is a wrapper around the C record GtkRcContext.
type RcContext struct {
	native unsafe.Pointer
}

// RcProperty is a wrapper around the C record GtkRcProperty.
type RcProperty struct {
	native       unsafe.Pointer
	TypeName     glib.Quark
	PropertyName glib.Quark
	Origin       string
	// value : record
}

// RcStyleClass is a wrapper around the C record GtkRcStyleClass.
type RcStyleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for create_rc_style
	// no type for parse
	// no type for merge
	// no type for create_style
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// RecentActionClass is a wrapper around the C record GtkRecentActionClass.
type RecentActionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// RecentActionPrivate is a wrapper around the C record GtkRecentActionPrivate.
type RecentActionPrivate struct {
	native unsafe.Pointer
}

// RecentChooserDialogClass is a wrapper around the C record GtkRecentChooserDialogClass.
type RecentChooserDialogClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// RecentChooserDialogPrivate is a wrapper around the C record GtkRecentChooserDialogPrivate.
type RecentChooserDialogPrivate struct {
	native unsafe.Pointer
}

// RecentChooserIface is a wrapper around the C record GtkRecentChooserIface.
type RecentChooserIface struct {
	native unsafe.Pointer
	// Private : base_iface
	// no type for set_current_uri
	// no type for get_current_uri
	// no type for select_uri
	// no type for unselect_uri
	// no type for select_all
	// no type for unselect_all
	// no type for get_items
	// no type for get_recent_manager
	// no type for add_filter
	// no type for remove_filter
	// no type for list_filters
	// no type for set_sort_func
	// no type for item_activated
	// no type for selection_changed
}

// RecentChooserMenuClass is a wrapper around the C record GtkRecentChooserMenuClass.
type RecentChooserMenuClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for gtk_recent1
	// no type for gtk_recent2
	// no type for gtk_recent3
	// no type for gtk_recent4
}

// RecentChooserMenuPrivate is a wrapper around the C record GtkRecentChooserMenuPrivate.
type RecentChooserMenuPrivate struct {
	native unsafe.Pointer
}

// RecentChooserWidgetClass is a wrapper around the C record GtkRecentChooserWidgetClass.
type RecentChooserWidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// RecentChooserWidgetPrivate is a wrapper around the C record GtkRecentChooserWidgetPrivate.
type RecentChooserWidgetPrivate struct {
	native unsafe.Pointer
}

// RecentData is a wrapper around the C record GtkRecentData.
type RecentData struct {
	native      unsafe.Pointer
	DisplayName string
	Description string
	MimeType    string
	AppName     string
	AppExec     string
	// no type for groups
	IsPrivate bool
}

// RecentFilterInfo is a wrapper around the C record GtkRecentFilterInfo.
type RecentFilterInfo struct {
	native      unsafe.Pointer
	Contains    RecentFilterFlags
	Uri         string
	DisplayName string
	MimeType    string
	// no type for applications
	// no type for groups
	Age int32
}

// RecentManagerPrivate is a wrapper around the C record GtkRecentManagerPrivate.
type RecentManagerPrivate struct {
	native unsafe.Pointer
}

// RendererCellAccessibleClass is a wrapper around the C record GtkRendererCellAccessibleClass.
type RendererCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// RendererCellAccessiblePrivate is a wrapper around the C record GtkRendererCellAccessiblePrivate.
type RendererCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// RequestedSize is a wrapper around the C record GtkRequestedSize.
type RequestedSize struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	MinimumSize int32
	NaturalSize int32
}

// Requisition is a wrapper around the C record GtkRequisition.
type Requisition struct {
	native unsafe.Pointer
	Width  int32
	Height int32
}

// RevealerClass is a wrapper around the C record GtkRevealerClass.
type RevealerClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ScaleAccessibleClass is a wrapper around the C record GtkScaleAccessibleClass.
type ScaleAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ScaleAccessiblePrivate is a wrapper around the C record GtkScaleAccessiblePrivate.
type ScaleAccessiblePrivate struct {
	native unsafe.Pointer
}

// ScaleButtonAccessibleClass is a wrapper around the C record GtkScaleButtonAccessibleClass.
type ScaleButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ScaleButtonAccessiblePrivate is a wrapper around the C record GtkScaleButtonAccessiblePrivate.
type ScaleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ScaleButtonClass is a wrapper around the C record GtkScaleButtonClass.
type ScaleButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for value_changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ScaleButtonPrivate is a wrapper around the C record GtkScaleButtonPrivate.
type ScaleButtonPrivate struct {
	native unsafe.Pointer
}

// ScaleClass is a wrapper around the C record GtkScaleClass.
type ScaleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for format_value
	// no type for draw_value
	// no type for get_layout_offsets
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ScalePrivate is a wrapper around the C record GtkScalePrivate.
type ScalePrivate struct {
	native unsafe.Pointer
}

// ScrollableInterface is a wrapper around the C record GtkScrollableInterface.
type ScrollableInterface struct {
	native unsafe.Pointer
	// base_iface : record
	// no type for get_border
}

// ScrollbarClass is a wrapper around the C record GtkScrollbarClass.
type ScrollbarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ScrolledWindowAccessibleClass is a wrapper around the C record GtkScrolledWindowAccessibleClass.
type ScrolledWindowAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ScrolledWindowAccessiblePrivate is a wrapper around the C record GtkScrolledWindowAccessiblePrivate.
type ScrolledWindowAccessiblePrivate struct {
	native unsafe.Pointer
}

// ScrolledWindowClass is a wrapper around the C record GtkScrolledWindowClass.
type ScrolledWindowClass struct {
	native unsafe.Pointer
	// parent_class : record
	ScrollbarSpacing int32
	// no type for scroll_child
	// no type for move_focus_out
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ScrolledWindowPrivate is a wrapper around the C record GtkScrolledWindowPrivate.
type ScrolledWindowPrivate struct {
	native unsafe.Pointer
}

// SearchBarClass is a wrapper around the C record GtkSearchBarClass.
type SearchBarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// SearchEntryClass is a wrapper around the C record GtkSearchEntryClass.
type SearchEntryClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for search_changed
	// no type for next_match
	// no type for previous_match
	// no type for stop_search
}

// SelectionData is a wrapper around the C record GtkSelectionData.
type SelectionData struct {
	native unsafe.Pointer
}

// SeparatorClass is a wrapper around the C record GtkSeparatorClass.
type SeparatorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// SeparatorMenuItemClass is a wrapper around the C record GtkSeparatorMenuItemClass.
type SeparatorMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// SeparatorPrivate is a wrapper around the C record GtkSeparatorPrivate.
type SeparatorPrivate struct {
	native unsafe.Pointer
}

// SeparatorToolItemClass is a wrapper around the C record GtkSeparatorToolItemClass.
type SeparatorToolItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// SeparatorToolItemPrivate is a wrapper around the C record GtkSeparatorToolItemPrivate.
type SeparatorToolItemPrivate struct {
	native unsafe.Pointer
}

// SettingsClass is a wrapper around the C record GtkSettingsClass.
type SettingsClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// SettingsPrivate is a wrapper around the C record GtkSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

// SettingsValue is a wrapper around the C record GtkSettingsValue.
type SettingsValue struct {
	native unsafe.Pointer
	Origin string
	// value : record
}

// SizeGroupClass is a wrapper around the C record GtkSizeGroupClass.
type SizeGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// SizeGroupPrivate is a wrapper around the C record GtkSizeGroupPrivate.
type SizeGroupPrivate struct {
	native unsafe.Pointer
}

// SocketClass is a wrapper around the C record GtkSocketClass.
type SocketClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for plug_added
	// no type for plug_removed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// SocketPrivate is a wrapper around the C record GtkSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

// SpinButtonAccessibleClass is a wrapper around the C record GtkSpinButtonAccessibleClass.
type SpinButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// SpinButtonAccessiblePrivate is a wrapper around the C record GtkSpinButtonAccessiblePrivate.
type SpinButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// SpinButtonClass is a wrapper around the C record GtkSpinButtonClass.
type SpinButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for input
	// no type for output
	// no type for value_changed
	// no type for change_value
	// no type for wrapped
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// SpinButtonPrivate is a wrapper around the C record GtkSpinButtonPrivate.
type SpinButtonPrivate struct {
	native unsafe.Pointer
}

// SpinnerAccessibleClass is a wrapper around the C record GtkSpinnerAccessibleClass.
type SpinnerAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// SpinnerAccessiblePrivate is a wrapper around the C record GtkSpinnerAccessiblePrivate.
type SpinnerAccessiblePrivate struct {
	native unsafe.Pointer
}

// SpinnerClass is a wrapper around the C record GtkSpinnerClass.
type SpinnerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// SpinnerPrivate is a wrapper around the C record GtkSpinnerPrivate.
type SpinnerPrivate struct {
	native unsafe.Pointer
}

// Blacklisted : GtkStackAccessibleClass

// StackClass is a wrapper around the C record GtkStackClass.
type StackClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// StackSidebarClass is a wrapper around the C record GtkStackSidebarClass.
type StackSidebarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// StackSidebarPrivate is a wrapper around the C record GtkStackSidebarPrivate.
type StackSidebarPrivate struct {
	native unsafe.Pointer
}

// StackSwitcherClass is a wrapper around the C record GtkStackSwitcherClass.
type StackSwitcherClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// StatusIconClass is a wrapper around the C record GtkStatusIconClass.
type StatusIconClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for popup_menu
	// no type for size_changed
	// no type for button_press_event
	// no type for button_release_event
	// no type for scroll_event
	// no type for query_tooltip
	// __gtk_reserved1 : no type generator for gpointer, void*
	// __gtk_reserved2 : no type generator for gpointer, void*
	// __gtk_reserved3 : no type generator for gpointer, void*
	// __gtk_reserved4 : no type generator for gpointer, void*
}

// StatusIconPrivate is a wrapper around the C record GtkStatusIconPrivate.
type StatusIconPrivate struct {
	native unsafe.Pointer
}

// StatusbarAccessibleClass is a wrapper around the C record GtkStatusbarAccessibleClass.
type StatusbarAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// StatusbarAccessiblePrivate is a wrapper around the C record GtkStatusbarAccessiblePrivate.
type StatusbarAccessiblePrivate struct {
	native unsafe.Pointer
}

// StatusbarClass is a wrapper around the C record GtkStatusbarClass.
type StatusbarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// reserved : no type generator for gpointer, gpointer
	// no type for text_pushed
	// no type for text_popped
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// StatusbarPrivate is a wrapper around the C record GtkStatusbarPrivate.
type StatusbarPrivate struct {
	native unsafe.Pointer
}

// StockItem is a wrapper around the C record GtkStockItem.
type StockItem struct {
	native            unsafe.Pointer
	StockId           string
	Label             string
	Modifier          gdk.ModifierType
	Keyval            uint32
	TranslationDomain string
}

// StyleClass is a wrapper around the C record GtkStyleClass.
type StyleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for realize
	// no type for unrealize
	// no type for copy
	// no type for clone
	// no type for init_from_rc
	// no type for set_background
	// no type for render_icon
	// no type for draw_hline
	// no type for draw_vline
	// no type for draw_shadow
	// no type for draw_arrow
	// no type for draw_diamond
	// no type for draw_box
	// no type for draw_flat_box
	// no type for draw_check
	// no type for draw_option
	// no type for draw_tab
	// no type for draw_shadow_gap
	// no type for draw_box_gap
	// no type for draw_extension
	// no type for draw_focus
	// no type for draw_slider
	// no type for draw_handle
	// no type for draw_expander
	// no type for draw_layout
	// no type for draw_resize_grip
	// no type for draw_spinner
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
	// no type for _gtk_reserved9
	// no type for _gtk_reserved10
	// no type for _gtk_reserved11
}

// StyleContextClass is a wrapper around the C record GtkStyleContextClass.
type StyleContextClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// StyleContextPrivate is a wrapper around the C record GtkStyleContextPrivate.
type StyleContextPrivate struct {
	native unsafe.Pointer
}

// StylePropertiesClass is a wrapper around the C record GtkStylePropertiesClass.
type StylePropertiesClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// StylePropertiesPrivate is a wrapper around the C record GtkStylePropertiesPrivate.
type StylePropertiesPrivate struct {
	native unsafe.Pointer
}

// StyleProviderIface is a wrapper around the C record GtkStyleProviderIface.
type StyleProviderIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for get_style
	// no type for get_style_property
	// no type for get_icon_factory
}

// SwitchAccessibleClass is a wrapper around the C record GtkSwitchAccessibleClass.
type SwitchAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// SwitchAccessiblePrivate is a wrapper around the C record GtkSwitchAccessiblePrivate.
type SwitchAccessiblePrivate struct {
	native unsafe.Pointer
}

// SwitchClass is a wrapper around the C record GtkSwitchClass.
type SwitchClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for activate
	// no type for state_set
	// no type for _switch_padding_1
	// no type for _switch_padding_2
	// no type for _switch_padding_3
	// no type for _switch_padding_4
	// no type for _switch_padding_5
}

// SwitchPrivate is a wrapper around the C record GtkSwitchPrivate.
type SwitchPrivate struct {
	native unsafe.Pointer
}

// SymbolicColor is a wrapper around the C record GtkSymbolicColor.
type SymbolicColor struct {
	native unsafe.Pointer
}

// TableChild is a wrapper around the C record GtkTableChild.
type TableChild struct {
	native unsafe.Pointer
	// widget : record
	LeftAttach   uint16
	RightAttach  uint16
	TopAttach    uint16
	BottomAttach uint16
	Xpadding     uint16
	Ypadding     uint16
	// Bitfield not supported :  1 xexpand
	// Bitfield not supported :  1 yexpand
	// Bitfield not supported :  1 xshrink
	// Bitfield not supported :  1 yshrink
	// Bitfield not supported :  1 xfill
	// Bitfield not supported :  1 yfill
}

// TableClass is a wrapper around the C record GtkTableClass.
type TableClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TablePrivate is a wrapper around the C record GtkTablePrivate.
type TablePrivate struct {
	native unsafe.Pointer
}

// TableRowCol is a wrapper around the C record GtkTableRowCol.
type TableRowCol struct {
	native      unsafe.Pointer
	Requisition uint16
	Allocation  uint16
	Spacing     uint16
	// Bitfield not supported :  1 need_expand
	// Bitfield not supported :  1 need_shrink
	// Bitfield not supported :  1 expand
	// Bitfield not supported :  1 shrink
	// Bitfield not supported :  1 empty
}

// TargetEntry is a wrapper around the C record GtkTargetEntry.
type TargetEntry struct {
	native unsafe.Pointer
	Target string
	Flags  uint32
	Info   uint32
}

// TargetList is a wrapper around the C record GtkTargetList.
type TargetList struct {
	native unsafe.Pointer
}

// TargetPair is a wrapper around the C record GtkTargetPair.
type TargetPair struct {
	native unsafe.Pointer
	// target : record
	Flags uint32
	Info  uint32
}

// TearoffMenuItemClass is a wrapper around the C record GtkTearoffMenuItemClass.
type TearoffMenuItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TearoffMenuItemPrivate is a wrapper around the C record GtkTearoffMenuItemPrivate.
type TearoffMenuItemPrivate struct {
	native unsafe.Pointer
}

// TextAppearance is a wrapper around the C record GtkTextAppearance.
type TextAppearance struct {
	native unsafe.Pointer
	// bg_color : record
	// fg_color : record
	Rise int32
	// Bitfield not supported :  4 underline
	// Bitfield not supported :  1 strikethrough
	// Bitfield not supported :  1 draw_bg
	// Bitfield not supported :  1 inside_selection
	// Bitfield not supported :  1 is_text
}

// TextAttributes is a wrapper around the C record GtkTextAttributes.
type TextAttributes struct {
	native unsafe.Pointer
	// Private : refcount
	// appearance : record
	Justification Justification
	Direction     TextDirection
	// font : record
	FontScale        float64
	LeftMargin       int32
	RightMargin      int32
	Indent           int32
	PixelsAboveLines int32
	PixelsBelowLines int32
	PixelsInsideWrap int32
	// tabs : record
	WrapMode WrapMode
	// language : record
	// Private : pg_bg_color
	// Bitfield not supported :  1 invisible
	// Bitfield not supported :  1 bg_full_height
	// Bitfield not supported :  1 editable
	// Bitfield not supported :  1 no_fallback
	// Private : pg_bg_rgba
	LetterSpacing int32
}

// TextBTree is a wrapper around the C record GtkTextBTree.
type TextBTree struct {
	native unsafe.Pointer
}

// TextBufferClass is a wrapper around the C record GtkTextBufferClass.
type TextBufferClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for insert_text
	// no type for insert_pixbuf
	// no type for insert_child_anchor
	// no type for delete_range
	// no type for changed
	// no type for modified_changed
	// no type for mark_set
	// no type for mark_deleted
	// no type for apply_tag
	// no type for remove_tag
	// no type for begin_user_action
	// no type for end_user_action
	// no type for paste_done
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TextBufferPrivate is a wrapper around the C record GtkTextBufferPrivate.
type TextBufferPrivate struct {
	native unsafe.Pointer
}

// TextCellAccessibleClass is a wrapper around the C record GtkTextCellAccessibleClass.
type TextCellAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// TextCellAccessiblePrivate is a wrapper around the C record GtkTextCellAccessiblePrivate.
type TextCellAccessiblePrivate struct {
	native unsafe.Pointer
}

// TextChildAnchorClass is a wrapper around the C record GtkTextChildAnchorClass.
type TextChildAnchorClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TextIter is a wrapper around the C record GtkTextIter.
type TextIter struct {
	native unsafe.Pointer
	// Private : dummy1
	// Private : dummy2
	// Private : dummy3
	// Private : dummy4
	// Private : dummy5
	// Private : dummy6
	// Private : dummy7
	// Private : dummy8
	// Private : dummy9
	// Private : dummy10
	// Private : dummy11
	// Private : dummy12
	// Private : dummy13
	// Private : dummy14
}

// TextMarkClass is a wrapper around the C record GtkTextMarkClass.
type TextMarkClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TextTagClass is a wrapper around the C record GtkTextTagClass.
type TextTagClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for event
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TextTagPrivate is a wrapper around the C record GtkTextTagPrivate.
type TextTagPrivate struct {
	native unsafe.Pointer
}

// TextTagTableClass is a wrapper around the C record GtkTextTagTableClass.
type TextTagTableClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for tag_changed
	// no type for tag_added
	// no type for tag_removed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TextTagTablePrivate is a wrapper around the C record GtkTextTagTablePrivate.
type TextTagTablePrivate struct {
	native unsafe.Pointer
}

// TextViewAccessibleClass is a wrapper around the C record GtkTextViewAccessibleClass.
type TextViewAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// TextViewAccessiblePrivate is a wrapper around the C record GtkTextViewAccessiblePrivate.
type TextViewAccessiblePrivate struct {
	native unsafe.Pointer
}

// TextViewClass is a wrapper around the C record GtkTextViewClass.
type TextViewClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for populate_popup
	// no type for move_cursor
	// no type for set_anchor
	// no type for insert_at_cursor
	// no type for delete_from_cursor
	// no type for backspace
	// no type for cut_clipboard
	// no type for copy_clipboard
	// no type for paste_clipboard
	// no type for toggle_overwrite
	// no type for create_buffer
	// no type for draw_layer
	// no type for extend_selection
	// no type for insert_emoji
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TextViewPrivate is a wrapper around the C record GtkTextViewPrivate.
type TextViewPrivate struct {
	native unsafe.Pointer
}

// ThemeEngine is a wrapper around the C record GtkThemeEngine.
type ThemeEngine struct {
	native unsafe.Pointer
}

// ThemingEngineClass is a wrapper around the C record GtkThemingEngineClass.
type ThemingEngineClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for render_line
	// no type for render_background
	// no type for render_frame
	// no type for render_frame_gap
	// no type for render_extension
	// no type for render_check
	// no type for render_option
	// no type for render_arrow
	// no type for render_expander
	// no type for render_focus
	// no type for render_layout
	// no type for render_slider
	// no type for render_handle
	// no type for render_activity
	// no type for render_icon_pixbuf
	// no type for render_icon
	// no type for render_icon_surface
	// Private : padding
}

// ThemingEnginePrivate is a wrapper around the C record GtkThemingEnginePrivate.
type ThemingEnginePrivate struct {
	native unsafe.Pointer
}

// ToggleActionClass is a wrapper around the C record GtkToggleActionClass.
type ToggleActionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ToggleActionEntry is a wrapper around the C record GtkToggleActionEntry.
type ToggleActionEntry struct {
	native      unsafe.Pointer
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	// callback : no type generator for GObject.Callback, GCallback
	IsActive bool
}

// ToggleActionPrivate is a wrapper around the C record GtkToggleActionPrivate.
type ToggleActionPrivate struct {
	native unsafe.Pointer
}

// ToggleButtonAccessibleClass is a wrapper around the C record GtkToggleButtonAccessibleClass.
type ToggleButtonAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ToggleButtonAccessiblePrivate is a wrapper around the C record GtkToggleButtonAccessiblePrivate.
type ToggleButtonAccessiblePrivate struct {
	native unsafe.Pointer
}

// ToggleButtonClass is a wrapper around the C record GtkToggleButtonClass.
type ToggleButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ToggleButtonPrivate is a wrapper around the C record GtkToggleButtonPrivate.
type ToggleButtonPrivate struct {
	native unsafe.Pointer
}

// ToggleToolButtonClass is a wrapper around the C record GtkToggleToolButtonClass.
type ToggleToolButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for toggled
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ToggleToolButtonPrivate is a wrapper around the C record GtkToggleToolButtonPrivate.
type ToggleToolButtonPrivate struct {
	native unsafe.Pointer
}

// ToolButtonClass is a wrapper around the C record GtkToolButtonClass.
type ToolButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	ButtonType gobject.Type
	// no type for clicked
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ToolButtonPrivate is a wrapper around the C record GtkToolButtonPrivate.
type ToolButtonPrivate struct {
	native unsafe.Pointer
}

// ToolItemClass is a wrapper around the C record GtkToolItemClass.
type ToolItemClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for create_menu_proxy
	// no type for toolbar_reconfigured
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ToolItemGroupClass is a wrapper around the C record GtkToolItemGroupClass.
type ToolItemGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ToolItemGroupPrivate is a wrapper around the C record GtkToolItemGroupPrivate.
type ToolItemGroupPrivate struct {
	native unsafe.Pointer
}

// ToolItemPrivate is a wrapper around the C record GtkToolItemPrivate.
type ToolItemPrivate struct {
	native unsafe.Pointer
}

// ToolPaletteClass is a wrapper around the C record GtkToolPaletteClass.
type ToolPaletteClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ToolPalettePrivate is a wrapper around the C record GtkToolPalettePrivate.
type ToolPalettePrivate struct {
	native unsafe.Pointer
}

// ToolShellIface is a wrapper around the C record GtkToolShellIface.
type ToolShellIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for get_icon_size
	// no type for get_orientation
	// no type for get_style
	// no type for get_relief_style
	// no type for rebuild_menu
	// no type for get_text_orientation
	// no type for get_text_alignment
	// no type for get_ellipsize_mode
	// no type for get_text_size_group
}

// ToolbarClass is a wrapper around the C record GtkToolbarClass.
type ToolbarClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for orientation_changed
	// no type for style_changed
	// no type for popup_context_menu
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ToolbarPrivate is a wrapper around the C record GtkToolbarPrivate.
type ToolbarPrivate struct {
	native unsafe.Pointer
}

// ToplevelAccessibleClass is a wrapper around the C record GtkToplevelAccessibleClass.
type ToplevelAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ToplevelAccessiblePrivate is a wrapper around the C record GtkToplevelAccessiblePrivate.
type ToplevelAccessiblePrivate struct {
	native unsafe.Pointer
}

// TreeDragDestIface is a wrapper around the C record GtkTreeDragDestIface.
type TreeDragDestIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for drag_data_received
	// no type for row_drop_possible
}

// TreeDragSourceIface is a wrapper around the C record GtkTreeDragSourceIface.
type TreeDragSourceIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for row_draggable
	// no type for drag_data_get
	// no type for drag_data_delete
}

// TreeIter is a wrapper around the C record GtkTreeIter.
type TreeIter struct {
	native unsafe.Pointer
	Stamp  int32
	// user_data : no type generator for gpointer, gpointer
	// user_data2 : no type generator for gpointer, gpointer
	// user_data3 : no type generator for gpointer, gpointer
}

// TreeModelFilterClass is a wrapper around the C record GtkTreeModelFilterClass.
type TreeModelFilterClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for visible
	// no type for modify
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TreeModelFilterPrivate is a wrapper around the C record GtkTreeModelFilterPrivate.
type TreeModelFilterPrivate struct {
	native unsafe.Pointer
}

// TreeModelIface is a wrapper around the C record GtkTreeModelIface.
type TreeModelIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for row_changed
	// no type for row_inserted
	// no type for row_has_child_toggled
	// no type for row_deleted
	// no type for rows_reordered
	// no type for get_flags
	// no type for get_n_columns
	// no type for get_column_type
	// no type for get_iter
	// no type for get_path
	// no type for get_value
	// no type for iter_next
	// no type for iter_previous
	// no type for iter_children
	// no type for iter_has_child
	// no type for iter_n_children
	// no type for iter_nth_child
	// no type for iter_parent
	// no type for ref_node
	// no type for unref_node
}

// TreeModelSortClass is a wrapper around the C record GtkTreeModelSortClass.
type TreeModelSortClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TreeModelSortPrivate is a wrapper around the C record GtkTreeModelSortPrivate.
type TreeModelSortPrivate struct {
	native unsafe.Pointer
}

// TreePath is a wrapper around the C record GtkTreePath.
type TreePath struct {
	native unsafe.Pointer
}

// TreeRowReference is a wrapper around the C record GtkTreeRowReference.
type TreeRowReference struct {
	native unsafe.Pointer
}

// TreeSelectionClass is a wrapper around the C record GtkTreeSelectionClass.
type TreeSelectionClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for changed
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TreeSelectionPrivate is a wrapper around the C record GtkTreeSelectionPrivate.
type TreeSelectionPrivate struct {
	native unsafe.Pointer
}

// TreeSortableIface is a wrapper around the C record GtkTreeSortableIface.
type TreeSortableIface struct {
	native unsafe.Pointer
	// Private : g_iface
	// no type for sort_column_changed
	// no type for get_sort_column_id
	// no type for set_sort_column_id
	// no type for set_sort_func
	// no type for set_default_sort_func
	// no type for has_default_sort_func
}

// TreeStoreClass is a wrapper around the C record GtkTreeStoreClass.
type TreeStoreClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TreeStorePrivate is a wrapper around the C record GtkTreeStorePrivate.
type TreeStorePrivate struct {
	native unsafe.Pointer
}

// TreeViewAccessibleClass is a wrapper around the C record GtkTreeViewAccessibleClass.
type TreeViewAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// TreeViewAccessiblePrivate is a wrapper around the C record GtkTreeViewAccessiblePrivate.
type TreeViewAccessiblePrivate struct {
	native unsafe.Pointer
}

// TreeViewClass is a wrapper around the C record GtkTreeViewClass.
type TreeViewClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for row_activated
	// no type for test_expand_row
	// no type for test_collapse_row
	// no type for row_expanded
	// no type for row_collapsed
	// no type for columns_changed
	// no type for cursor_changed
	// no type for move_cursor
	// no type for select_all
	// no type for unselect_all
	// no type for select_cursor_row
	// no type for toggle_cursor_row
	// no type for expand_collapse_cursor_row
	// no type for select_cursor_parent
	// no type for start_interactive_search
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
	// no type for _gtk_reserved5
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
	// no type for _gtk_reserved8
}

// TreeViewColumnClass is a wrapper around the C record GtkTreeViewColumnClass.
type TreeViewColumnClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for clicked
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// TreeViewColumnPrivate is a wrapper around the C record GtkTreeViewColumnPrivate.
type TreeViewColumnPrivate struct {
	native unsafe.Pointer
}

// TreeViewPrivate is a wrapper around the C record GtkTreeViewPrivate.
type TreeViewPrivate struct {
	native unsafe.Pointer
}

// UIManagerClass is a wrapper around the C record GtkUIManagerClass.
type UIManagerClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for add_widget
	// no type for actions_changed
	// no type for connect_proxy
	// no type for disconnect_proxy
	// no type for pre_activate
	// no type for post_activate
	// no type for get_widget
	// no type for get_action
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// UIManagerPrivate is a wrapper around the C record GtkUIManagerPrivate.
type UIManagerPrivate struct {
	native unsafe.Pointer
}

// VBoxClass is a wrapper around the C record GtkVBoxClass.
type VBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// VButtonBoxClass is a wrapper around the C record GtkVButtonBoxClass.
type VButtonBoxClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// VPanedClass is a wrapper around the C record GtkVPanedClass.
type VPanedClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// VScaleClass is a wrapper around the C record GtkVScaleClass.
type VScaleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// VScrollbarClass is a wrapper around the C record GtkVScrollbarClass.
type VScrollbarClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// VSeparatorClass is a wrapper around the C record GtkVSeparatorClass.
type VSeparatorClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// ViewportClass is a wrapper around the C record GtkViewportClass.
type ViewportClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// ViewportPrivate is a wrapper around the C record GtkViewportPrivate.
type ViewportPrivate struct {
	native unsafe.Pointer
}

// VolumeButtonClass is a wrapper around the C record GtkVolumeButtonClass.
type VolumeButtonClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// WidgetAccessibleClass is a wrapper around the C record GtkWidgetAccessibleClass.
type WidgetAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for notify_gtk
}

// WidgetAccessiblePrivate is a wrapper around the C record GtkWidgetAccessiblePrivate.
type WidgetAccessiblePrivate struct {
	native unsafe.Pointer
}

// WidgetClass is a wrapper around the C record GtkWidgetClass.
type WidgetClass struct {
	native unsafe.Pointer
	// parent_class : record
	ActivateSignal uint32
	// no type for dispatch_child_properties_changed
	// no type for destroy
	// no type for show
	// no type for show_all
	// no type for hide
	// no type for map
	// no type for unmap
	// no type for realize
	// no type for unrealize
	// no type for size_allocate
	// no type for state_changed
	// no type for state_flags_changed
	// no type for parent_set
	// no type for hierarchy_changed
	// no type for style_set
	// no type for direction_changed
	// no type for grab_notify
	// no type for child_notify
	// no type for draw
	// no type for get_request_mode
	// no type for get_preferred_height
	// no type for get_preferred_width_for_height
	// no type for get_preferred_width
	// no type for get_preferred_height_for_width
	// no type for mnemonic_activate
	// no type for grab_focus
	// no type for focus
	// no type for move_focus
	// no type for keynav_failed
	// no type for event
	// no type for button_press_event
	// no type for button_release_event
	// no type for scroll_event
	// no type for motion_notify_event
	// no type for delete_event
	// no type for destroy_event
	// no type for key_press_event
	// no type for key_release_event
	// no type for enter_notify_event
	// no type for leave_notify_event
	// no type for configure_event
	// no type for focus_in_event
	// no type for focus_out_event
	// no type for map_event
	// no type for unmap_event
	// no type for property_notify_event
	// no type for selection_clear_event
	// no type for selection_request_event
	// no type for selection_notify_event
	// no type for proximity_in_event
	// no type for proximity_out_event
	// no type for visibility_notify_event
	// no type for window_state_event
	// no type for damage_event
	// no type for grab_broken_event
	// no type for selection_get
	// no type for selection_received
	// no type for drag_begin
	// no type for drag_end
	// no type for drag_data_get
	// no type for drag_data_delete
	// no type for drag_leave
	// no type for drag_motion
	// no type for drag_drop
	// no type for drag_data_received
	// no type for drag_failed
	// no type for popup_menu
	// no type for show_help
	// no type for get_accessible
	// no type for screen_changed
	// no type for can_activate_accel
	// no type for composited_changed
	// no type for query_tooltip
	// no type for compute_expand
	// no type for adjust_size_request
	// no type for adjust_size_allocation
	// no type for style_updated
	// no type for touch_event
	// no type for get_preferred_height_and_baseline_for_width
	// no type for adjust_baseline_request
	// no type for adjust_baseline_allocation
	// no type for queue_draw_region
	// Private : priv
	// no type for _gtk_reserved6
	// no type for _gtk_reserved7
}

// WidgetClassPrivate is a wrapper around the C record GtkWidgetClassPrivate.
type WidgetClassPrivate struct {
	native unsafe.Pointer
}

// WidgetPath is a wrapper around the C record GtkWidgetPath.
type WidgetPath struct {
	native unsafe.Pointer
}

// WidgetPrivate is a wrapper around the C record GtkWidgetPrivate.
type WidgetPrivate struct {
	native unsafe.Pointer
}

// WindowAccessibleClass is a wrapper around the C record GtkWindowAccessibleClass.
type WindowAccessibleClass struct {
	native unsafe.Pointer
	// parent_class : record
}

// WindowAccessiblePrivate is a wrapper around the C record GtkWindowAccessiblePrivate.
type WindowAccessiblePrivate struct {
	native unsafe.Pointer
}

// WindowClass is a wrapper around the C record GtkWindowClass.
type WindowClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for set_focus
	// no type for activate_focus
	// no type for activate_default
	// no type for keys_changed
	// no type for enable_debugging
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
}

// WindowGeometryInfo is a wrapper around the C record GtkWindowGeometryInfo.
type WindowGeometryInfo struct {
	native unsafe.Pointer
}

// WindowGroupClass is a wrapper around the C record GtkWindowGroupClass.
type WindowGroupClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

// WindowGroupPrivate is a wrapper around the C record GtkWindowGroupPrivate.
type WindowGroupPrivate struct {
	native unsafe.Pointer
}

// WindowPrivate is a wrapper around the C record GtkWindowPrivate.
type WindowPrivate struct {
	native unsafe.Pointer
}
