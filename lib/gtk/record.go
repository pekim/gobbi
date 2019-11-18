// Code generated - DO NOT EDIT.

package gtk

import gi "github.com/pekim/gobbi/internal/gi"

type AboutDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value 'activate_link' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type AboutDialogPrivate struct {
	native uintptr
}

type AccelGroupClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'accel_changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type AccelGroupEntry struct {
	native uintptr
	Key    *AccelKey
	// UNSUPPORTED : C value 'closure' : no Go type for 'GObject.Closure'

	// UNSUPPORTED : C value 'accel_path_quark' : no Go type for 'GLib.Quark'

}

type AccelGroupPrivate struct {
	native uintptr
}

type AccelKey struct {
	native   uintptr
	AccelKey uint32
	// UNSUPPORTED : C value 'accel_mods' : no Go type for 'Gdk.ModifierType'

	AccelFlags uint32
}

type AccelLabelClass struct {
	native         uintptr
	ParentClass    *LabelClass
	SignalQuote1   string
	SignalQuote2   string
	ModNameShift   string
	ModNameControl string
	ModNameAlt     string
	ModSeparator   string
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type AccelLabelPrivate struct {
	native uintptr
}

type AccelMapClass struct {
	native uintptr
}

type AccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Atk.ObjectClass'

	// UNSUPPORTED : C value 'connect_widget_destroyed' : missing Type

	// UNSUPPORTED : C value 'widget_set' : missing Type

	// UNSUPPORTED : C value 'widget_unset' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type AccessiblePrivate struct {
	native uintptr
}

type ActionBarClass struct {
	native uintptr
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ActionBarPrivate struct {
	native uintptr
}

type ActionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value 'create_menu_item' : missing Type

	// UNSUPPORTED : C value 'create_tool_item' : missing Type

	// UNSUPPORTED : C value 'connect_proxy' : missing Type

	// UNSUPPORTED : C value 'disconnect_proxy' : missing Type

	// UNSUPPORTED : C value 'create_menu' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ActionEntry struct {
	native      uintptr
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	// UNSUPPORTED : C value 'callback' : no Go type for 'GObject.Callback'

}

type ActionGroupClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_action' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ActionGroupPrivate struct {
	native uintptr
}

type ActionPrivate struct {
	native uintptr
}

type ActionableInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'get_action_name' : missing Type

	// UNSUPPORTED : C value 'set_action_name' : missing Type

	// UNSUPPORTED : C value 'get_action_target_value' : missing Type

	// UNSUPPORTED : C value 'set_action_target_value' : missing Type

}

type ActivatableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'update' : missing Type

	// UNSUPPORTED : C value 'sync_action_properties' : missing Type

}

type AdjustmentClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'value_changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type AdjustmentPrivate struct {
	native uintptr
}

type AlignmentClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type AlignmentPrivate struct {
	native uintptr
}

type AppChooserButtonClass struct {
	native      uintptr
	ParentClass *ComboBoxClass
	// UNSUPPORTED : C value 'custom_item_activated' : missing Type

}

type AppChooserButtonPrivate struct {
	native uintptr
}

type AppChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
}

type AppChooserDialogPrivate struct {
	native uintptr
}

type AppChooserWidgetClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value 'application_selected' : missing Type

	// UNSUPPORTED : C value 'application_activated' : missing Type

	// UNSUPPORTED : C value 'populate_popup' : missing Type

}

type AppChooserWidgetPrivate struct {
	native uintptr
}

type ApplicationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.ApplicationClass'

	// UNSUPPORTED : C value 'window_added' : missing Type

	// UNSUPPORTED : C value 'window_removed' : missing Type

}

type ApplicationPrivate struct {
	native uintptr
}

type ApplicationWindowClass struct {
	native      uintptr
	ParentClass *WindowClass
}

type ApplicationWindowPrivate struct {
	native uintptr
}

type ArrowAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

type ArrowAccessiblePrivate struct {
	native uintptr
}

type ArrowClass struct {
	native      uintptr
	ParentClass *MiscClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ArrowPrivate struct {
	native uintptr
}

type AspectFrameClass struct {
	native      uintptr
	ParentClass *FrameClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type AspectFramePrivate struct {
	native uintptr
}

type AssistantClass struct {
	native      uintptr
	ParentClass *WindowClass
	// UNSUPPORTED : C value 'prepare' : missing Type

	// UNSUPPORTED : C value 'apply' : missing Type

	// UNSUPPORTED : C value 'close' : missing Type

	// UNSUPPORTED : C value 'cancel' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

}

type AssistantPrivate struct {
	native uintptr
}

type BinClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type BinPrivate struct {
	native uintptr
}

type BindingArg struct {
	native uintptr
	// UNSUPPORTED : C value 'arg_type' : no Go type for 'GType'

}

type BindingEntry struct {
	native uintptr
	Keyval uint32
	// UNSUPPORTED : C value 'modifiers' : no Go type for 'Gdk.ModifierType'

	BindingSet   *BindingSet
	Destroyed    uint32
	InEmission   uint32
	MarksUnbound uint32
	SetNext      *BindingEntry
	HashNext     *BindingEntry
	Signals      *BindingSignal
}

type BindingSet struct {
	native   uintptr
	SetName  string
	Priority int32
	// UNSUPPORTED : C value 'widget_path_pspecs' : no Go type for 'GLib.SList'

	// UNSUPPORTED : C value 'widget_class_pspecs' : no Go type for 'GLib.SList'

	// UNSUPPORTED : C value 'class_branch_pspecs' : no Go type for 'GLib.SList'

	Entries *BindingEntry
	Current *BindingEntry
	Parsed  uint32
}

type BindingSignal struct {
	native     uintptr
	Next       *BindingSignal
	SignalName string
	NArgs      uint32
	// UNSUPPORTED : C value 'args' : missing Type

}

type BooleanCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

type BooleanCellAccessiblePrivate struct {
	native uintptr
}

type Border struct {
	native uintptr
	Left   int16
	Right  int16
	Top    int16
	Bottom int16
}

var newBorderInvoker *gi.Function

// BorderNew is a representation of the C type gtk_border_new.
func BorderNew() *Border {
	if newBorderInvoker == nil {
		newBorderInvoker = gi.StructFunctionInvokerNew("Gtk", "Border", "new")
	}

	ret := newBorderInvoker.Invoke(nil, nil)

	return &Border{native: ret.Pointer()}
}

type BoxClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type BoxPrivate struct {
	native uintptr
}

type BuildableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'set_name' : missing Type

	// UNSUPPORTED : C value 'get_name' : missing Type

	// UNSUPPORTED : C value 'add_child' : missing Type

	// UNSUPPORTED : C value 'set_buildable_property' : missing Type

	// UNSUPPORTED : C value 'construct_child' : missing Type

	// UNSUPPORTED : C value 'custom_tag_start' : missing Type

	// UNSUPPORTED : C value 'custom_tag_end' : missing Type

	// UNSUPPORTED : C value 'custom_finished' : missing Type

	// UNSUPPORTED : C value 'parser_finished' : missing Type

	// UNSUPPORTED : C value 'get_internal_child' : missing Type

}

type BuilderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'get_type_from_name' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type BuilderPrivate struct {
	native uintptr
}

type ButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type ButtonAccessiblePrivate struct {
	native uintptr
}

type ButtonBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ButtonBoxPrivate struct {
	native uintptr
}

type ButtonClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'pressed' : missing Type

	// UNSUPPORTED : C value 'released' : missing Type

	// UNSUPPORTED : C value 'clicked' : missing Type

	// UNSUPPORTED : C value 'enter' : missing Type

	// UNSUPPORTED : C value 'leave' : missing Type

	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ButtonPrivate struct {
	native uintptr
}

type CalendarClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value 'month_changed' : missing Type

	// UNSUPPORTED : C value 'day_selected' : missing Type

	// UNSUPPORTED : C value 'day_selected_double_click' : missing Type

	// UNSUPPORTED : C value 'prev_month' : missing Type

	// UNSUPPORTED : C value 'next_month' : missing Type

	// UNSUPPORTED : C value 'prev_year' : missing Type

	// UNSUPPORTED : C value 'next_year' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CalendarPrivate struct {
	native uintptr
}

type CellAccessibleClass struct {
	native      uintptr
	ParentClass *AccessibleClass
	// UNSUPPORTED : C value 'update_cache' : missing Type

}

type CellAccessibleParentIface struct {
	native uintptr
	// UNSUPPORTED : C value 'parent' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_cell_extents' : missing Type

	// UNSUPPORTED : C value 'get_cell_area' : missing Type

	// UNSUPPORTED : C value 'grab_focus' : missing Type

	// UNSUPPORTED : C value 'get_child_index' : missing Type

	// UNSUPPORTED : C value 'get_renderer_state' : missing Type

	// UNSUPPORTED : C value 'expand_collapse' : missing Type

	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value 'edit' : missing Type

	// UNSUPPORTED : C value 'update_relationset' : missing Type

	// UNSUPPORTED : C value 'get_cell_position' : missing Type

	// UNSUPPORTED : C value 'get_column_header_cells' : missing Type

	// UNSUPPORTED : C value 'get_row_header_cells' : missing Type

}

type CellAccessiblePrivate struct {
	native uintptr
}

type CellAreaBoxClass struct {
	native uintptr
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellAreaBoxPrivate struct {
	native uintptr
}

type CellAreaClass struct {
	native uintptr
	// UNSUPPORTED : C value 'add' : missing Type

	// UNSUPPORTED : C value 'remove' : missing Type

	// UNSUPPORTED : C value 'foreach' : missing Type

	// UNSUPPORTED : C value 'foreach_alloc' : missing Type

	// UNSUPPORTED : C value 'event' : missing Type

	// UNSUPPORTED : C value 'render' : missing Type

	// UNSUPPORTED : C value 'apply_attributes' : missing Type

	// UNSUPPORTED : C value 'create_context' : missing Type

	// UNSUPPORTED : C value 'copy_context' : missing Type

	// UNSUPPORTED : C value 'get_request_mode' : missing Type

	// UNSUPPORTED : C value 'get_preferred_width' : missing Type

	// UNSUPPORTED : C value 'get_preferred_height_for_width' : missing Type

	// UNSUPPORTED : C value 'get_preferred_height' : missing Type

	// UNSUPPORTED : C value 'get_preferred_width_for_height' : missing Type

	// UNSUPPORTED : C value 'set_cell_property' : missing Type

	// UNSUPPORTED : C value 'get_cell_property' : missing Type

	// UNSUPPORTED : C value 'focus' : missing Type

	// UNSUPPORTED : C value 'is_activatable' : missing Type

	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type CellAreaContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'allocate' : missing Type

	// UNSUPPORTED : C value 'reset' : missing Type

	// UNSUPPORTED : C value 'get_preferred_height_for_width' : missing Type

	// UNSUPPORTED : C value 'get_preferred_width_for_height' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

}

type CellAreaContextPrivate struct {
	native uintptr
}

type CellAreaPrivate struct {
	native uintptr
}

type CellEditableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'editing_done' : missing Type

	// UNSUPPORTED : C value 'remove_widget' : missing Type

	// UNSUPPORTED : C value 'start_editing' : missing Type

}

type CellLayoutIface struct {
	native uintptr
	// UNSUPPORTED : C value 'pack_start' : missing Type

	// UNSUPPORTED : C value 'pack_end' : missing Type

	// UNSUPPORTED : C value 'clear' : missing Type

	// UNSUPPORTED : C value 'add_attribute' : missing Type

	// UNSUPPORTED : C value 'set_cell_data_func' : missing Type

	// UNSUPPORTED : C value 'clear_attributes' : missing Type

	// UNSUPPORTED : C value 'reorder' : missing Type

	// UNSUPPORTED : C value 'get_cells' : missing Type

	// UNSUPPORTED : C value 'get_area' : missing Type

}

type CellRendererAccelClass struct {
	native      uintptr
	ParentClass *CellRendererTextClass
	// UNSUPPORTED : C value 'accel_edited' : missing Type

	// UNSUPPORTED : C value 'accel_cleared' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved0' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellRendererAccelPrivate struct {
	native uintptr
}

type CellRendererClass struct {
	native uintptr
	// UNSUPPORTED : C value 'get_request_mode' : missing Type

	// UNSUPPORTED : C value 'get_preferred_width' : missing Type

	// UNSUPPORTED : C value 'get_preferred_height_for_width' : missing Type

	// UNSUPPORTED : C value 'get_preferred_height' : missing Type

	// UNSUPPORTED : C value 'get_preferred_width_for_height' : missing Type

	// UNSUPPORTED : C value 'get_aligned_area' : missing Type

	// UNSUPPORTED : C value 'get_size' : missing Type

	// UNSUPPORTED : C value 'render' : missing Type

	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value 'start_editing' : missing Type

	// UNSUPPORTED : C value 'editing_canceled' : missing Type

	// UNSUPPORTED : C value 'editing_started' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellRendererClassPrivate struct {
	native uintptr
}

type CellRendererComboClass struct {
	native uintptr
	Parent *CellRendererTextClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellRendererComboPrivate struct {
	native uintptr
}

type CellRendererPixbufClass struct {
	native      uintptr
	ParentClass *CellRendererClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellRendererPixbufPrivate struct {
	native uintptr
}

type CellRendererPrivate struct {
	native uintptr
}

type CellRendererProgressClass struct {
	native      uintptr
	ParentClass *CellRendererClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellRendererProgressPrivate struct {
	native uintptr
}

type CellRendererSpinClass struct {
	native uintptr
	Parent *CellRendererTextClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellRendererSpinPrivate struct {
	native uintptr
}

type CellRendererSpinnerClass struct {
	native      uintptr
	ParentClass *CellRendererClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellRendererSpinnerPrivate struct {
	native uintptr
}

type CellRendererTextClass struct {
	native      uintptr
	ParentClass *CellRendererClass
	// UNSUPPORTED : C value 'edited' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellRendererTextPrivate struct {
	native uintptr
}

type CellRendererToggleClass struct {
	native      uintptr
	ParentClass *CellRendererClass
	// UNSUPPORTED : C value 'toggled' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellRendererTogglePrivate struct {
	native uintptr
}

type CellViewClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CellViewPrivate struct {
	native uintptr
}

type CheckButtonClass struct {
	native      uintptr
	ParentClass *ToggleButtonClass
	// UNSUPPORTED : C value 'draw_indicator' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CheckMenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *MenuItemAccessibleClass
}

type CheckMenuItemAccessiblePrivate struct {
	native uintptr
}

type CheckMenuItemClass struct {
	native      uintptr
	ParentClass *MenuItemClass
	// UNSUPPORTED : C value 'toggled' : missing Type

	// UNSUPPORTED : C value 'draw_indicator' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CheckMenuItemPrivate struct {
	native uintptr
}

type ColorButtonClass struct {
	native      uintptr
	ParentClass *ButtonClass
	// UNSUPPORTED : C value 'color_set' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ColorButtonPrivate struct {
	native uintptr
}

type ColorChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ColorChooserDialogPrivate struct {
	native uintptr
}

type ColorChooserInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_interface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_rgba' : missing Type

	// UNSUPPORTED : C value 'set_rgba' : missing Type

	// UNSUPPORTED : C value 'add_palette' : missing Type

	// UNSUPPORTED : C value 'color_activated' : missing Type

	// UNSUPPORTED : C value 'padding' : missing Type

}

type ColorChooserWidgetClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type ColorChooserWidgetPrivate struct {
	native uintptr
}

type ColorSelectionClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value 'color_changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ColorSelectionDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ColorSelectionDialogPrivate struct {
	native uintptr
}

type ColorSelectionPrivate struct {
	native uintptr
}

type ComboBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type ComboBoxAccessiblePrivate struct {
	native uintptr
}

type ComboBoxClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'format_entry_text' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

}

type ComboBoxPrivate struct {
	native uintptr
}

type ComboBoxTextClass struct {
	native      uintptr
	ParentClass *ComboBoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ComboBoxTextPrivate struct {
	native uintptr
}

type ContainerAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
	// UNSUPPORTED : C value 'add_gtk' : missing Type

	// UNSUPPORTED : C value 'remove_gtk' : missing Type

}

type ContainerAccessiblePrivate struct {
	native uintptr
}

type ContainerCellAccessibleClass struct {
	native      uintptr
	ParentClass *CellAccessibleClass
}

type ContainerCellAccessiblePrivate struct {
	native uintptr
}

type ContainerClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value 'add' : missing Type

	// UNSUPPORTED : C value 'remove' : missing Type

	// UNSUPPORTED : C value 'check_resize' : missing Type

	// UNSUPPORTED : C value 'forall' : missing Type

	// UNSUPPORTED : C value 'set_focus_child' : missing Type

	// UNSUPPORTED : C value 'child_type' : missing Type

	// UNSUPPORTED : C value 'composite_name' : missing Type

	// UNSUPPORTED : C value 'set_child_property' : missing Type

	// UNSUPPORTED : C value 'get_child_property' : missing Type

	// UNSUPPORTED : C value 'get_path_for_child' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type ContainerPrivate struct {
	native uintptr
}

type CssProviderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'parsing_error' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type CssProviderPrivate struct {
	native uintptr
}

type CssSection struct {
	native uintptr
}

type DialogClass struct {
	native      uintptr
	ParentClass *WindowClass
	// UNSUPPORTED : C value 'response' : missing Type

	// UNSUPPORTED : C value 'close' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type DialogPrivate struct {
	native uintptr
}

type DrawingAreaClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type EditableInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'insert_text' : missing Type

	// UNSUPPORTED : C value 'delete_text' : missing Type

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'do_insert_text' : missing Type

	// UNSUPPORTED : C value 'do_delete_text' : missing Type

	// UNSUPPORTED : C value 'get_chars' : missing Type

	// UNSUPPORTED : C value 'set_selection_bounds' : missing Type

	// UNSUPPORTED : C value 'get_selection_bounds' : missing Type

	// UNSUPPORTED : C value 'set_position' : missing Type

	// UNSUPPORTED : C value 'get_position' : missing Type

}

type EntryAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

type EntryAccessiblePrivate struct {
	native uintptr
}

type EntryBufferClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'inserted_text' : missing Type

	// UNSUPPORTED : C value 'deleted_text' : missing Type

	// UNSUPPORTED : C value 'get_text' : missing Type

	// UNSUPPORTED : C value 'get_length' : missing Type

	// UNSUPPORTED : C value 'insert_text' : missing Type

	// UNSUPPORTED : C value 'delete_text' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type EntryBufferPrivate struct {
	native uintptr
}

type EntryClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value 'populate_popup' : missing Type

	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value 'move_cursor' : missing Type

	// UNSUPPORTED : C value 'insert_at_cursor' : missing Type

	// UNSUPPORTED : C value 'delete_from_cursor' : missing Type

	// UNSUPPORTED : C value 'backspace' : missing Type

	// UNSUPPORTED : C value 'cut_clipboard' : missing Type

	// UNSUPPORTED : C value 'copy_clipboard' : missing Type

	// UNSUPPORTED : C value 'paste_clipboard' : missing Type

	// UNSUPPORTED : C value 'toggle_overwrite' : missing Type

	// UNSUPPORTED : C value 'get_text_area_size' : missing Type

	// UNSUPPORTED : C value 'get_frame_size' : missing Type

	// UNSUPPORTED : C value 'insert_emoji' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

}

type EntryCompletionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'match_selected' : missing Type

	// UNSUPPORTED : C value 'action_activated' : missing Type

	// UNSUPPORTED : C value 'insert_prefix' : missing Type

	// UNSUPPORTED : C value 'cursor_on_match' : missing Type

	// UNSUPPORTED : C value 'no_matches' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved0' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

}

type EntryCompletionPrivate struct {
	native uintptr
}

type EntryPrivate struct {
	native uintptr
}

type EventBoxClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type EventBoxPrivate struct {
	native uintptr
}

type EventControllerClass struct {
	native uintptr
}

type EventControllerKeyClass struct {
	native uintptr
}

type EventControllerMotionClass struct {
	native uintptr
}

type EventControllerScrollClass struct {
	native uintptr
}

type ExpanderAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type ExpanderAccessiblePrivate struct {
	native uintptr
}

type ExpanderClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ExpanderPrivate struct {
	native uintptr
}

type FileChooserButtonClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value 'file_set' : missing Type

	// UNSUPPORTED : C value '__gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '__gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '__gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '__gtk_reserved4' : missing Type

}

type FileChooserButtonPrivate struct {
	native uintptr
}

type FileChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type FileChooserDialogPrivate struct {
	native uintptr
}

type FileChooserNativeClass struct {
	native      uintptr
	ParentClass *NativeDialogClass
}

type FileChooserWidgetClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type FileChooserWidgetPrivate struct {
	native uintptr
}

type FileFilterInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'contains' : no Go type for 'FileFilterFlags'

	Filename    string
	Uri         string
	DisplayName string
	MimeType    string
}

type FixedChild struct {
	native uintptr
	// UNSUPPORTED : C value 'widget' : no Go type for 'Widget'

	X int32
	Y int32
}

type FixedClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type FixedPrivate struct {
	native uintptr
}

type FlowBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type FlowBoxAccessiblePrivate struct {
	native uintptr
}

type FlowBoxChildAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type FlowBoxChildClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

}

type FlowBoxClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value 'child_activated' : missing Type

	// UNSUPPORTED : C value 'selected_children_changed' : missing Type

	// UNSUPPORTED : C value 'activate_cursor_child' : missing Type

	// UNSUPPORTED : C value 'toggle_cursor_child' : missing Type

	// UNSUPPORTED : C value 'move_cursor' : missing Type

	// UNSUPPORTED : C value 'select_all' : missing Type

	// UNSUPPORTED : C value 'unselect_all' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

}

type FontButtonClass struct {
	native      uintptr
	ParentClass *ButtonClass
	// UNSUPPORTED : C value 'font_set' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type FontButtonPrivate struct {
	native uintptr
}

type FontChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type FontChooserDialogPrivate struct {
	native uintptr
}

type FontChooserIface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_font_family' : missing Type

	// UNSUPPORTED : C value 'get_font_face' : missing Type

	// UNSUPPORTED : C value 'get_font_size' : missing Type

	// UNSUPPORTED : C value 'set_filter_func' : missing Type

	// UNSUPPORTED : C value 'font_activated' : missing Type

	// UNSUPPORTED : C value 'set_font_map' : missing Type

	// UNSUPPORTED : C value 'get_font_map' : missing Type

	// UNSUPPORTED : C value 'padding' : missing Type

}

type FontChooserWidgetClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type FontChooserWidgetPrivate struct {
	native uintptr
}

type FontSelectionClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type FontSelectionDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type FontSelectionDialogPrivate struct {
	native uintptr
}

type FontSelectionPrivate struct {
	native uintptr
}

type FrameAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type FrameAccessiblePrivate struct {
	native uintptr
}

type FrameClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'compute_child_allocation' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type FramePrivate struct {
	native uintptr
}

type GLAreaClass struct {
	native uintptr
	// UNSUPPORTED : C value 'render' : missing Type

	// UNSUPPORTED : C value 'resize' : missing Type

	// UNSUPPORTED : C value 'create_context' : missing Type

}

type GestureClass struct {
	native uintptr
}

type GestureDragClass struct {
	native uintptr
}

type GestureLongPressClass struct {
	native uintptr
}

type GestureMultiPressClass struct {
	native uintptr
}

type GesturePanClass struct {
	native uintptr
}

type GestureRotateClass struct {
	native uintptr
}

type GestureSingleClass struct {
	native uintptr
}

type GestureStylusClass struct {
	native uintptr
}

type GestureSwipeClass struct {
	native uintptr
}

type GestureZoomClass struct {
	native uintptr
}

type Gradient struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_gradient_new_linear' : parameter 'x0' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_gradient_new_radial' : parameter 'x0' of type 'gdouble' not supported

type GridClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type GridPrivate struct {
	native uintptr
}

type HBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
}

type HButtonBoxClass struct {
	native      uintptr
	ParentClass *ButtonBoxClass
}

type HPanedClass struct {
	native      uintptr
	ParentClass *PanedClass
}

type HSVClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'move' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type HSVPrivate struct {
	native uintptr
}

type HScaleClass struct {
	native      uintptr
	ParentClass *ScaleClass
}

type HScrollbarClass struct {
	native      uintptr
	ParentClass *ScrollbarClass
}

type HSeparatorClass struct {
	native      uintptr
	ParentClass *SeparatorClass
}

type HandleBoxClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'child_attached' : missing Type

	// UNSUPPORTED : C value 'child_detached' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type HandleBoxPrivate struct {
	native uintptr
}

type HeaderBarAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type HeaderBarAccessiblePrivate struct {
	native uintptr
}

type HeaderBarClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type HeaderBarPrivate struct {
	native uintptr
}

type IMContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'preedit_start' : missing Type

	// UNSUPPORTED : C value 'preedit_end' : missing Type

	// UNSUPPORTED : C value 'preedit_changed' : missing Type

	// UNSUPPORTED : C value 'commit' : missing Type

	// UNSUPPORTED : C value 'retrieve_surrounding' : missing Type

	// UNSUPPORTED : C value 'delete_surrounding' : missing Type

	// UNSUPPORTED : C value 'set_client_window' : missing Type

	// UNSUPPORTED : C value 'get_preedit_string' : missing Type

	// UNSUPPORTED : C value 'filter_keypress' : missing Type

	// UNSUPPORTED : C value 'focus_in' : missing Type

	// UNSUPPORTED : C value 'focus_out' : missing Type

	// UNSUPPORTED : C value 'reset' : missing Type

	// UNSUPPORTED : C value 'set_cursor_location' : missing Type

	// UNSUPPORTED : C value 'set_use_preedit' : missing Type

	// UNSUPPORTED : C value 'set_surrounding' : missing Type

	// UNSUPPORTED : C value 'get_surrounding' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

}

type IMContextInfo struct {
	native         uintptr
	ContextId      string
	ContextName    string
	Domain         string
	DomainDirname  string
	DefaultLocales string
}

type IMContextSimpleClass struct {
	native      uintptr
	ParentClass *IMContextClass
}

type IMContextSimplePrivate struct {
	native uintptr
}

type IMMulticontextClass struct {
	native      uintptr
	ParentClass *IMContextClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type IMMulticontextPrivate struct {
	native uintptr
}

type IconFactoryClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type IconFactoryPrivate struct {
	native uintptr
}

type IconInfoClass struct {
	native uintptr
}

type IconSet struct {
	native uintptr
}

var newIconSetInvoker *gi.Function

// IconSetNew is a representation of the C type gtk_icon_set_new.
func IconSetNew() *IconSet {
	if newIconSetInvoker == nil {
		newIconSetInvoker = gi.StructFunctionInvokerNew("Gtk", "IconSet", "new")
	}

	ret := newIconSetInvoker.Invoke(nil, nil)

	return &IconSet{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'gtk_icon_set_new_from_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

type IconSource struct {
	native uintptr
}

var newIconSourceInvoker *gi.Function

// IconSourceNew is a representation of the C type gtk_icon_source_new.
func IconSourceNew() *IconSource {
	if newIconSourceInvoker == nil {
		newIconSourceInvoker = gi.StructFunctionInvokerNew("Gtk", "IconSource", "new")
	}

	ret := newIconSourceInvoker.Invoke(nil, nil)

	return &IconSource{native: ret.Pointer()}
}

type IconThemeClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type IconThemePrivate struct {
	native uintptr
}

type IconViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type IconViewAccessiblePrivate struct {
	native uintptr
}

type IconViewClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value 'item_activated' : missing Type

	// UNSUPPORTED : C value 'selection_changed' : missing Type

	// UNSUPPORTED : C value 'select_all' : missing Type

	// UNSUPPORTED : C value 'unselect_all' : missing Type

	// UNSUPPORTED : C value 'select_cursor_item' : missing Type

	// UNSUPPORTED : C value 'toggle_cursor_item' : missing Type

	// UNSUPPORTED : C value 'move_cursor' : missing Type

	// UNSUPPORTED : C value 'activate_cursor_item' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type IconViewPrivate struct {
	native uintptr
}

type ImageAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

type ImageAccessiblePrivate struct {
	native uintptr
}

type ImageCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

type ImageCellAccessiblePrivate struct {
	native uintptr
}

type ImageClass struct {
	native      uintptr
	ParentClass *MiscClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ImageMenuItemClass struct {
	native      uintptr
	ParentClass *MenuItemClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ImageMenuItemPrivate struct {
	native uintptr
}

type ImagePrivate struct {
	native uintptr
}

type InfoBarClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value 'response' : missing Type

	// UNSUPPORTED : C value 'close' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type InfoBarPrivate struct {
	native uintptr
}

type InvisibleClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type InvisiblePrivate struct {
	native uintptr
}

type LabelAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

type LabelAccessiblePrivate struct {
	native uintptr
}

type LabelClass struct {
	native      uintptr
	ParentClass *MiscClass
	// UNSUPPORTED : C value 'move_cursor' : missing Type

	// UNSUPPORTED : C value 'copy_clipboard' : missing Type

	// UNSUPPORTED : C value 'populate_popup' : missing Type

	// UNSUPPORTED : C value 'activate_link' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type LabelPrivate struct {
	native uintptr
}

type LabelSelectionInfo struct {
	native uintptr
}

type LayoutClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type LayoutPrivate struct {
	native uintptr
}

type LevelBarAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

type LevelBarAccessiblePrivate struct {
	native uintptr
}

type LevelBarClass struct {
	native uintptr
	// UNSUPPORTED : C value 'offset_changed' : missing Type

}

type LevelBarPrivate struct {
	native uintptr
}

type LinkButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

type LinkButtonAccessiblePrivate struct {
	native uintptr
}

type LinkButtonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'activate_link' : missing Type

	// UNSUPPORTED : C value '_gtk_padding1' : missing Type

	// UNSUPPORTED : C value '_gtk_padding2' : missing Type

	// UNSUPPORTED : C value '_gtk_padding3' : missing Type

	// UNSUPPORTED : C value '_gtk_padding4' : missing Type

}

type LinkButtonPrivate struct {
	native uintptr
}

type ListBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type ListBoxAccessiblePrivate struct {
	native uintptr
}

type ListBoxClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value 'row_selected' : missing Type

	// UNSUPPORTED : C value 'row_activated' : missing Type

	// UNSUPPORTED : C value 'activate_cursor_row' : missing Type

	// UNSUPPORTED : C value 'toggle_cursor_row' : missing Type

	// UNSUPPORTED : C value 'move_cursor' : missing Type

	// UNSUPPORTED : C value 'selected_rows_changed' : missing Type

	// UNSUPPORTED : C value 'select_all' : missing Type

	// UNSUPPORTED : C value 'unselect_all' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

}

type ListBoxRowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type ListBoxRowClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

}

type ListStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ListStorePrivate struct {
	native uintptr
}

type LockButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

type LockButtonAccessiblePrivate struct {
	native uintptr
}

type LockButtonClass struct {
	native      uintptr
	ParentClass *ButtonClass
	// UNSUPPORTED : C value 'reserved0' : missing Type

	// UNSUPPORTED : C value 'reserved1' : missing Type

	// UNSUPPORTED : C value 'reserved2' : missing Type

	// UNSUPPORTED : C value 'reserved3' : missing Type

	// UNSUPPORTED : C value 'reserved4' : missing Type

	// UNSUPPORTED : C value 'reserved5' : missing Type

	// UNSUPPORTED : C value 'reserved6' : missing Type

	// UNSUPPORTED : C value 'reserved7' : missing Type

}

type LockButtonPrivate struct {
	native uintptr
}

type MenuAccessibleClass struct {
	native      uintptr
	ParentClass *MenuShellAccessibleClass
}

type MenuAccessiblePrivate struct {
	native uintptr
}

type MenuBarClass struct {
	native      uintptr
	ParentClass *MenuShellClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type MenuBarPrivate struct {
	native uintptr
}

type MenuButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ToggleButtonAccessibleClass
}

type MenuButtonAccessiblePrivate struct {
	native uintptr
}

type MenuButtonClass struct {
	native      uintptr
	ParentClass *ToggleButtonClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type MenuButtonPrivate struct {
	native uintptr
}

type MenuClass struct {
	native      uintptr
	ParentClass *MenuShellClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type MenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type MenuItemAccessiblePrivate struct {
	native uintptr
}

type MenuItemClass struct {
	native         uintptr
	ParentClass    *BinClass
	HideOnActivate uint32
	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value 'activate_item' : missing Type

	// UNSUPPORTED : C value 'toggle_size_request' : missing Type

	// UNSUPPORTED : C value 'toggle_size_allocate' : missing Type

	// UNSUPPORTED : C value 'set_label' : missing Type

	// UNSUPPORTED : C value 'get_label' : missing Type

	// UNSUPPORTED : C value 'select' : missing Type

	// UNSUPPORTED : C value 'deselect' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type MenuItemPrivate struct {
	native uintptr
}

type MenuPrivate struct {
	native uintptr
}

type MenuShellAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type MenuShellAccessiblePrivate struct {
	native uintptr
}

type MenuShellClass struct {
	native           uintptr
	ParentClass      *ContainerClass
	SubmenuPlacement uint32
	// UNSUPPORTED : C value 'deactivate' : missing Type

	// UNSUPPORTED : C value 'selection_done' : missing Type

	// UNSUPPORTED : C value 'move_current' : missing Type

	// UNSUPPORTED : C value 'activate_current' : missing Type

	// UNSUPPORTED : C value 'cancel' : missing Type

	// UNSUPPORTED : C value 'select_item' : missing Type

	// UNSUPPORTED : C value 'insert' : missing Type

	// UNSUPPORTED : C value 'get_popup_delay' : missing Type

	// UNSUPPORTED : C value 'move_selected' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type MenuShellPrivate struct {
	native uintptr
}

type MenuToolButtonClass struct {
	native      uintptr
	ParentClass *ToolButtonClass
	// UNSUPPORTED : C value 'show_menu' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type MenuToolButtonPrivate struct {
	native uintptr
}

type MessageDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type MessageDialogPrivate struct {
	native uintptr
}

type MiscClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type MiscPrivate struct {
	native uintptr
}

type MountOperationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.MountOperationClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type MountOperationPrivate struct {
	native uintptr
}

type NativeDialogClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'response' : missing Type

	// UNSUPPORTED : C value 'show' : missing Type

	// UNSUPPORTED : C value 'hide' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type NotebookAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type NotebookAccessiblePrivate struct {
	native uintptr
}

type NotebookClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value 'switch_page' : missing Type

	// UNSUPPORTED : C value 'select_page' : missing Type

	// UNSUPPORTED : C value 'focus_tab' : missing Type

	// UNSUPPORTED : C value 'change_current_page' : missing Type

	// UNSUPPORTED : C value 'move_focus_out' : missing Type

	// UNSUPPORTED : C value 'reorder_tab' : missing Type

	// UNSUPPORTED : C value 'insert_page' : missing Type

	// UNSUPPORTED : C value 'create_window' : missing Type

	// UNSUPPORTED : C value 'page_reordered' : missing Type

	// UNSUPPORTED : C value 'page_removed' : missing Type

	// UNSUPPORTED : C value 'page_added' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type NotebookPageAccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Atk.ObjectClass'

}

type NotebookPageAccessiblePrivate struct {
	native uintptr
}

type NotebookPrivate struct {
	native uintptr
}

type NumerableIconClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.EmblemedIconClass'

	// UNSUPPORTED : C value 'padding' : missing Type

}

type NumerableIconPrivate struct {
	native uintptr
}

type OffscreenWindowClass struct {
	native      uintptr
	ParentClass *WindowClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type OrientableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'

}

type OverlayClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'get_child_position' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type OverlayPrivate struct {
	native uintptr
}

type PadActionEntry struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'PadActionType'

	Index      int32
	Mode       int32
	Label      string
	ActionName string
}

type PadControllerClass struct {
	native uintptr
}

type PageRange struct {
	native uintptr
	Start  int32
	End    int32
}

type PanedAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type PanedAccessiblePrivate struct {
	native uintptr
}

type PanedClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value 'cycle_child_focus' : missing Type

	// UNSUPPORTED : C value 'toggle_handle_focus' : missing Type

	// UNSUPPORTED : C value 'move_handle' : missing Type

	// UNSUPPORTED : C value 'cycle_handle_focus' : missing Type

	// UNSUPPORTED : C value 'accept_position' : missing Type

	// UNSUPPORTED : C value 'cancel_position' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type PanedPrivate struct {
	native uintptr
}

type PaperSize struct {
	native uintptr
}

var newPaperSizeInvoker *gi.Function

// PaperSizeNew is a representation of the C type gtk_paper_size_new.
func PaperSizeNew(name string) *PaperSize {
	if newPaperSizeInvoker == nil {
		newPaperSizeInvoker = gi.StructFunctionInvokerNew("Gtk", "PaperSize", "new")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	ret := newPaperSizeInvoker.Invoke(inArgs[:], nil)

	return &PaperSize{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'gtk_paper_size_new_custom' : parameter 'width' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_gvariant' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_ipp' : parameter 'width' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_key_file' : parameter 'key_file' of type 'GLib.KeyFile' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_ppd' : parameter 'width' of type 'gdouble' not supported

type PlacesSidebarClass struct {
	native uintptr
}

type PlugClass struct {
	native      uintptr
	ParentClass *WindowClass
	// UNSUPPORTED : C value 'embedded' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type PlugPrivate struct {
	native uintptr
}

type PopoverAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type PopoverClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'closed' : missing Type

}

type PopoverMenuClass struct {
	native      uintptr
	ParentClass *PopoverClass
}

type PopoverPrivate struct {
	native uintptr
}

type PrintOperationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'done' : missing Type

	// UNSUPPORTED : C value 'begin_print' : missing Type

	// UNSUPPORTED : C value 'paginate' : missing Type

	// UNSUPPORTED : C value 'request_page_setup' : missing Type

	// UNSUPPORTED : C value 'draw_page' : missing Type

	// UNSUPPORTED : C value 'end_print' : missing Type

	// UNSUPPORTED : C value 'status_changed' : missing Type

	// UNSUPPORTED : C value 'create_custom_widget' : missing Type

	// UNSUPPORTED : C value 'custom_widget_apply' : missing Type

	// UNSUPPORTED : C value 'preview' : missing Type

	// UNSUPPORTED : C value 'update_custom_widget' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type PrintOperationPreviewIface struct {
	native uintptr
	// UNSUPPORTED : C value 'g_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'ready' : missing Type

	// UNSUPPORTED : C value 'got_page_size' : missing Type

	// UNSUPPORTED : C value 'render_page' : missing Type

	// UNSUPPORTED : C value 'is_selected' : missing Type

	// UNSUPPORTED : C value 'end_preview' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type PrintOperationPrivate struct {
	native uintptr
}

type ProgressBarAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

type ProgressBarAccessiblePrivate struct {
	native uintptr
}

type ProgressBarClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ProgressBarPrivate struct {
	native uintptr
}

type RadioActionClass struct {
	native      uintptr
	ParentClass *ToggleActionClass
	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type RadioActionEntry struct {
	native      uintptr
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	Value       int32
}

type RadioActionPrivate struct {
	native uintptr
}

type RadioButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ToggleButtonAccessibleClass
}

type RadioButtonAccessiblePrivate struct {
	native uintptr
}

type RadioButtonClass struct {
	native      uintptr
	ParentClass *CheckButtonClass
	// UNSUPPORTED : C value 'group_changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type RadioButtonPrivate struct {
	native uintptr
}

type RadioMenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *CheckMenuItemAccessibleClass
}

type RadioMenuItemAccessiblePrivate struct {
	native uintptr
}

type RadioMenuItemClass struct {
	native      uintptr
	ParentClass *CheckMenuItemClass
	// UNSUPPORTED : C value 'group_changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type RadioMenuItemPrivate struct {
	native uintptr
}

type RadioToolButtonClass struct {
	native      uintptr
	ParentClass *ToggleToolButtonClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type RangeAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

type RangeAccessiblePrivate struct {
	native uintptr
}

type RangeClass struct {
	native        uintptr
	ParentClass   *WidgetClass
	SliderDetail  string
	StepperDetail string
	// UNSUPPORTED : C value 'value_changed' : missing Type

	// UNSUPPORTED : C value 'adjust_bounds' : missing Type

	// UNSUPPORTED : C value 'move_slider' : missing Type

	// UNSUPPORTED : C value 'get_range_border' : missing Type

	// UNSUPPORTED : C value 'change_value' : missing Type

	// UNSUPPORTED : C value 'get_range_size_request' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

}

type RangePrivate struct {
	native uintptr
}

type RcContext struct {
	native uintptr
}

type RcProperty struct {
	native uintptr
	// UNSUPPORTED : C value 'type_name' : no Go type for 'GLib.Quark'

	// UNSUPPORTED : C value 'property_name' : no Go type for 'GLib.Quark'

	Origin string
	// UNSUPPORTED : C value 'value' : no Go type for 'GObject.Value'

}

type RcStyleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'create_rc_style' : missing Type

	// UNSUPPORTED : C value 'parse' : missing Type

	// UNSUPPORTED : C value 'merge' : missing Type

	// UNSUPPORTED : C value 'create_style' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type RecentActionClass struct {
	native      uintptr
	ParentClass *ActionClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type RecentActionPrivate struct {
	native uintptr
}

type RecentChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type RecentChooserDialogPrivate struct {
	native uintptr
}

type RecentChooserIface struct {
	native uintptr
	// UNSUPPORTED : C value 'set_current_uri' : missing Type

	// UNSUPPORTED : C value 'get_current_uri' : missing Type

	// UNSUPPORTED : C value 'select_uri' : missing Type

	// UNSUPPORTED : C value 'unselect_uri' : missing Type

	// UNSUPPORTED : C value 'select_all' : missing Type

	// UNSUPPORTED : C value 'unselect_all' : missing Type

	// UNSUPPORTED : C value 'get_items' : missing Type

	// UNSUPPORTED : C value 'get_recent_manager' : missing Type

	// UNSUPPORTED : C value 'add_filter' : missing Type

	// UNSUPPORTED : C value 'remove_filter' : missing Type

	// UNSUPPORTED : C value 'list_filters' : missing Type

	// UNSUPPORTED : C value 'set_sort_func' : missing Type

	// UNSUPPORTED : C value 'item_activated' : missing Type

	// UNSUPPORTED : C value 'selection_changed' : missing Type

}

type RecentChooserMenuClass struct {
	native      uintptr
	ParentClass *MenuClass
	// UNSUPPORTED : C value 'gtk_recent1' : missing Type

	// UNSUPPORTED : C value 'gtk_recent2' : missing Type

	// UNSUPPORTED : C value 'gtk_recent3' : missing Type

	// UNSUPPORTED : C value 'gtk_recent4' : missing Type

}

type RecentChooserMenuPrivate struct {
	native uintptr
}

type RecentChooserWidgetClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type RecentChooserWidgetPrivate struct {
	native uintptr
}

type RecentData struct {
	native      uintptr
	DisplayName string
	Description string
	MimeType    string
	AppName     string
	AppExec     string
	// UNSUPPORTED : C value 'groups' : missing Type

	// UNSUPPORTED : C value 'is_private' : no Go type for 'gboolean'

}

type RecentFilterInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'contains' : no Go type for 'RecentFilterFlags'

	Uri         string
	DisplayName string
	MimeType    string
	// UNSUPPORTED : C value 'applications' : missing Type

	// UNSUPPORTED : C value 'groups' : missing Type

	Age int32
}

type RecentInfo struct {
	native uintptr
}

type RecentManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value '_gtk_recent1' : missing Type

	// UNSUPPORTED : C value '_gtk_recent2' : missing Type

	// UNSUPPORTED : C value '_gtk_recent3' : missing Type

	// UNSUPPORTED : C value '_gtk_recent4' : missing Type

}

type RecentManagerPrivate struct {
	native uintptr
}

type RendererCellAccessibleClass struct {
	native      uintptr
	ParentClass *CellAccessibleClass
}

type RendererCellAccessiblePrivate struct {
	native uintptr
}

type RequestedSize struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'

	MinimumSize int32
	NaturalSize int32
}

type Requisition struct {
	native uintptr
	Width  int32
	Height int32
}

var newRequisitionInvoker *gi.Function

// RequisitionNew is a representation of the C type gtk_requisition_new.
func RequisitionNew() *Requisition {
	if newRequisitionInvoker == nil {
		newRequisitionInvoker = gi.StructFunctionInvokerNew("Gtk", "Requisition", "new")
	}

	ret := newRequisitionInvoker.Invoke(nil, nil)

	return &Requisition{native: ret.Pointer()}
}

type RevealerClass struct {
	native      uintptr
	ParentClass *BinClass
}

type ScaleAccessibleClass struct {
	native      uintptr
	ParentClass *RangeAccessibleClass
}

type ScaleAccessiblePrivate struct {
	native uintptr
}

type ScaleButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

type ScaleButtonAccessiblePrivate struct {
	native uintptr
}

type ScaleButtonClass struct {
	native      uintptr
	ParentClass *ButtonClass
	// UNSUPPORTED : C value 'value_changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ScaleButtonPrivate struct {
	native uintptr
}

type ScaleClass struct {
	native      uintptr
	ParentClass *RangeClass
	// UNSUPPORTED : C value 'format_value' : missing Type

	// UNSUPPORTED : C value 'draw_value' : missing Type

	// UNSUPPORTED : C value 'get_layout_offsets' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ScalePrivate struct {
	native uintptr
}

type ScrollableInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'

	// UNSUPPORTED : C value 'get_border' : missing Type

}

type ScrollbarClass struct {
	native      uintptr
	ParentClass *RangeClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ScrolledWindowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type ScrolledWindowAccessiblePrivate struct {
	native uintptr
}

type ScrolledWindowClass struct {
	native           uintptr
	ParentClass      *BinClass
	ScrollbarSpacing int32
	// UNSUPPORTED : C value 'scroll_child' : missing Type

	// UNSUPPORTED : C value 'move_focus_out' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ScrolledWindowPrivate struct {
	native uintptr
}

type SearchBarClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type SearchEntryClass struct {
	native      uintptr
	ParentClass *EntryClass
	// UNSUPPORTED : C value 'search_changed' : missing Type

	// UNSUPPORTED : C value 'next_match' : missing Type

	// UNSUPPORTED : C value 'previous_match' : missing Type

	// UNSUPPORTED : C value 'stop_search' : missing Type

}

type SelectionData struct {
	native uintptr
}

type SeparatorClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type SeparatorMenuItemClass struct {
	native      uintptr
	ParentClass *MenuItemClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type SeparatorPrivate struct {
	native uintptr
}

type SeparatorToolItemClass struct {
	native      uintptr
	ParentClass *ToolItemClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type SeparatorToolItemPrivate struct {
	native uintptr
}

type SettingsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type SettingsPrivate struct {
	native uintptr
}

type SettingsValue struct {
	native uintptr
	Origin string
	// UNSUPPORTED : C value 'value' : no Go type for 'GObject.Value'

}

type ShortcutLabelClass struct {
	native uintptr
}

type ShortcutsGroupClass struct {
	native uintptr
}

type ShortcutsSectionClass struct {
	native uintptr
}

type ShortcutsShortcutClass struct {
	native uintptr
}

type ShortcutsWindowClass struct {
	native      uintptr
	ParentClass *WindowClass
	// UNSUPPORTED : C value 'close' : missing Type

	// UNSUPPORTED : C value 'search' : missing Type

}

type SizeGroupClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type SizeGroupPrivate struct {
	native uintptr
}

type SocketClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value 'plug_added' : missing Type

	// UNSUPPORTED : C value 'plug_removed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type SocketPrivate struct {
	native uintptr
}

type SpinButtonAccessibleClass struct {
	native      uintptr
	ParentClass *EntryAccessibleClass
}

type SpinButtonAccessiblePrivate struct {
	native uintptr
}

type SpinButtonClass struct {
	native      uintptr
	ParentClass *EntryClass
	// UNSUPPORTED : C value 'input' : missing Type

	// UNSUPPORTED : C value 'output' : missing Type

	// UNSUPPORTED : C value 'value_changed' : missing Type

	// UNSUPPORTED : C value 'change_value' : missing Type

	// UNSUPPORTED : C value 'wrapped' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type SpinButtonPrivate struct {
	native uintptr
}

type SpinnerAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

type SpinnerAccessiblePrivate struct {
	native uintptr
}

type SpinnerClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type SpinnerPrivate struct {
	native uintptr
}

type StackAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type StackClass struct {
	native      uintptr
	ParentClass *ContainerClass
}

type StackSidebarClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type StackSidebarPrivate struct {
	native uintptr
}

type StackSwitcherClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type StatusIconClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value 'popup_menu' : missing Type

	// UNSUPPORTED : C value 'size_changed' : missing Type

	// UNSUPPORTED : C value 'button_press_event' : missing Type

	// UNSUPPORTED : C value 'button_release_event' : missing Type

	// UNSUPPORTED : C value 'scroll_event' : missing Type

	// UNSUPPORTED : C value 'query_tooltip' : missing Type

	// UNSUPPORTED : C value '__gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '__gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '__gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '__gtk_reserved4' : missing Type

}

type StatusIconPrivate struct {
	native uintptr
}

type StatusbarAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type StatusbarAccessiblePrivate struct {
	native uintptr
}

type StatusbarClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value 'reserved' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'text_pushed' : missing Type

	// UNSUPPORTED : C value 'text_popped' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type StatusbarPrivate struct {
	native uintptr
}

type StockItem struct {
	native  uintptr
	StockId string
	Label   string
	// UNSUPPORTED : C value 'modifier' : no Go type for 'Gdk.ModifierType'

	Keyval            uint32
	TranslationDomain string
}

type StyleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'realize' : missing Type

	// UNSUPPORTED : C value 'unrealize' : missing Type

	// UNSUPPORTED : C value 'copy' : missing Type

	// UNSUPPORTED : C value 'clone' : missing Type

	// UNSUPPORTED : C value 'init_from_rc' : missing Type

	// UNSUPPORTED : C value 'set_background' : missing Type

	// UNSUPPORTED : C value 'render_icon' : missing Type

	// UNSUPPORTED : C value 'draw_hline' : missing Type

	// UNSUPPORTED : C value 'draw_vline' : missing Type

	// UNSUPPORTED : C value 'draw_shadow' : missing Type

	// UNSUPPORTED : C value 'draw_arrow' : missing Type

	// UNSUPPORTED : C value 'draw_diamond' : missing Type

	// UNSUPPORTED : C value 'draw_box' : missing Type

	// UNSUPPORTED : C value 'draw_flat_box' : missing Type

	// UNSUPPORTED : C value 'draw_check' : missing Type

	// UNSUPPORTED : C value 'draw_option' : missing Type

	// UNSUPPORTED : C value 'draw_tab' : missing Type

	// UNSUPPORTED : C value 'draw_shadow_gap' : missing Type

	// UNSUPPORTED : C value 'draw_box_gap' : missing Type

	// UNSUPPORTED : C value 'draw_extension' : missing Type

	// UNSUPPORTED : C value 'draw_focus' : missing Type

	// UNSUPPORTED : C value 'draw_slider' : missing Type

	// UNSUPPORTED : C value 'draw_handle' : missing Type

	// UNSUPPORTED : C value 'draw_expander' : missing Type

	// UNSUPPORTED : C value 'draw_layout' : missing Type

	// UNSUPPORTED : C value 'draw_resize_grip' : missing Type

	// UNSUPPORTED : C value 'draw_spinner' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved9' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved10' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved11' : missing Type

}

type StyleContextClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type StyleContextPrivate struct {
	native uintptr
}

type StylePropertiesClass struct {
	native uintptr
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type StylePropertiesPrivate struct {
	native uintptr
}

type StyleProviderIface struct {
	native uintptr
	// UNSUPPORTED : C value 'get_style' : missing Type

	// UNSUPPORTED : C value 'get_style_property' : missing Type

	// UNSUPPORTED : C value 'get_icon_factory' : missing Type

}

type SwitchAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

type SwitchAccessiblePrivate struct {
	native uintptr
}

type SwitchClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value 'activate' : missing Type

	// UNSUPPORTED : C value 'state_set' : missing Type

	// UNSUPPORTED : C value '_switch_padding_1' : missing Type

	// UNSUPPORTED : C value '_switch_padding_2' : missing Type

	// UNSUPPORTED : C value '_switch_padding_3' : missing Type

	// UNSUPPORTED : C value '_switch_padding_4' : missing Type

	// UNSUPPORTED : C value '_switch_padding_5' : missing Type

}

type SwitchPrivate struct {
	native uintptr
}

type SymbolicColor struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_symbolic_color_new_alpha' : parameter 'color' of type 'SymbolicColor' not supported

// UNSUPPORTED : C value 'gtk_symbolic_color_new_literal' : parameter 'color' of type 'Gdk.RGBA' not supported

// UNSUPPORTED : C value 'gtk_symbolic_color_new_mix' : parameter 'color1' of type 'SymbolicColor' not supported

var newNameSymbolicColorInvoker *gi.Function

// SymbolicColorNewName is a representation of the C type gtk_symbolic_color_new_name.
func SymbolicColorNewName(name string) *SymbolicColor {
	if newNameSymbolicColorInvoker == nil {
		newNameSymbolicColorInvoker = gi.StructFunctionInvokerNew("Gtk", "SymbolicColor", "new_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	ret := newNameSymbolicColorInvoker.Invoke(inArgs[:], nil)

	return &SymbolicColor{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'gtk_symbolic_color_new_shade' : parameter 'color' of type 'SymbolicColor' not supported

var newWin32SymbolicColorInvoker *gi.Function

// SymbolicColorNewWin32 is a representation of the C type gtk_symbolic_color_new_win32.
func SymbolicColorNewWin32(themeClass string, id int32) *SymbolicColor {
	if newWin32SymbolicColorInvoker == nil {
		newWin32SymbolicColorInvoker = gi.StructFunctionInvokerNew("Gtk", "SymbolicColor", "new_win32")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetString(themeClass)
	inArgs[1].SetInt32(id)

	ret := newWin32SymbolicColorInvoker.Invoke(inArgs[:], nil)

	return &SymbolicColor{native: ret.Pointer()}
}

type TableChild struct {
	native uintptr
	// UNSUPPORTED : C value 'widget' : no Go type for 'Widget'

	LeftAttach   uint16
	RightAttach  uint16
	TopAttach    uint16
	BottomAttach uint16
	Xpadding     uint16
	Ypadding     uint16
	Xexpand      uint32
	Yexpand      uint32
	Xshrink      uint32
	Yshrink      uint32
	Xfill        uint32
	Yfill        uint32
}

type TableClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TablePrivate struct {
	native uintptr
}

type TableRowCol struct {
	native      uintptr
	Requisition uint16
	Allocation  uint16
	Spacing     uint16
	NeedExpand  uint32
	NeedShrink  uint32
	Expand      uint32
	Shrink      uint32
	Empty       uint32
}

type TargetEntry struct {
	native uintptr
	Target string
	Flags  uint32
	Info   uint32
}

var newTargetEntryInvoker *gi.Function

// TargetEntryNew is a representation of the C type gtk_target_entry_new.
func TargetEntryNew(target string, flags uint32, info uint32) *TargetEntry {
	if newTargetEntryInvoker == nil {
		newTargetEntryInvoker = gi.StructFunctionInvokerNew("Gtk", "TargetEntry", "new")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetString(target)
	inArgs[1].SetUint32(flags)
	inArgs[2].SetUint32(info)

	ret := newTargetEntryInvoker.Invoke(inArgs[:], nil)

	return &TargetEntry{native: ret.Pointer()}
}

type TargetList struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_target_list_new' : parameter 'targets' has no type

type TargetPair struct {
	native uintptr
	// UNSUPPORTED : C value 'target' : no Go type for 'Gdk.Atom'

	Flags uint32
	Info  uint32
}

type TearoffMenuItemClass struct {
	native      uintptr
	ParentClass *MenuItemClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TearoffMenuItemPrivate struct {
	native uintptr
}

type TextAppearance struct {
	native uintptr
	// UNSUPPORTED : C value 'bg_color' : no Go type for 'Gdk.Color'

	// UNSUPPORTED : C value 'fg_color' : no Go type for 'Gdk.Color'

	Rise            int32
	Underline       uint32
	Strikethrough   uint32
	DrawBg          uint32
	InsideSelection uint32
	IsText          uint32
}

type TextAttributes struct {
	native     uintptr
	Appearance *TextAppearance
	// UNSUPPORTED : C value 'justification' : no Go type for 'Justification'

	// UNSUPPORTED : C value 'direction' : no Go type for 'TextDirection'

	// UNSUPPORTED : C value 'font' : no Go type for 'Pango.FontDescription'

	// UNSUPPORTED : C value 'font_scale' : no Go type for 'gdouble'

	LeftMargin       int32
	RightMargin      int32
	Indent           int32
	PixelsAboveLines int32
	PixelsBelowLines int32
	PixelsInsideWrap int32
	// UNSUPPORTED : C value 'tabs' : no Go type for 'Pango.TabArray'

	// UNSUPPORTED : C value 'wrap_mode' : no Go type for 'WrapMode'

	// UNSUPPORTED : C value 'language' : no Go type for 'Pango.Language'

	Invisible     uint32
	BgFullHeight  uint32
	Editable      uint32
	NoFallback    uint32
	LetterSpacing int32
}

var newTextAttributesInvoker *gi.Function

// TextAttributesNew is a representation of the C type gtk_text_attributes_new.
func TextAttributesNew() *TextAttributes {
	if newTextAttributesInvoker == nil {
		newTextAttributesInvoker = gi.StructFunctionInvokerNew("Gtk", "TextAttributes", "new")
	}

	ret := newTextAttributesInvoker.Invoke(nil, nil)

	return &TextAttributes{native: ret.Pointer()}
}

type TextBTree struct {
	native uintptr
}

type TextBufferClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'insert_text' : missing Type

	// UNSUPPORTED : C value 'insert_pixbuf' : missing Type

	// UNSUPPORTED : C value 'insert_child_anchor' : missing Type

	// UNSUPPORTED : C value 'delete_range' : missing Type

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value 'modified_changed' : missing Type

	// UNSUPPORTED : C value 'mark_set' : missing Type

	// UNSUPPORTED : C value 'mark_deleted' : missing Type

	// UNSUPPORTED : C value 'apply_tag' : missing Type

	// UNSUPPORTED : C value 'remove_tag' : missing Type

	// UNSUPPORTED : C value 'begin_user_action' : missing Type

	// UNSUPPORTED : C value 'end_user_action' : missing Type

	// UNSUPPORTED : C value 'paste_done' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TextBufferPrivate struct {
	native uintptr
}

type TextCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

type TextCellAccessiblePrivate struct {
	native uintptr
}

type TextChildAnchorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TextIter struct {
	native uintptr
}

type TextMarkClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TextTagClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'event' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TextTagPrivate struct {
	native uintptr
}

type TextTagTableClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'tag_changed' : missing Type

	// UNSUPPORTED : C value 'tag_added' : missing Type

	// UNSUPPORTED : C value 'tag_removed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TextTagTablePrivate struct {
	native uintptr
}

type TextViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type TextViewAccessiblePrivate struct {
	native uintptr
}

type TextViewClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value 'populate_popup' : missing Type

	// UNSUPPORTED : C value 'move_cursor' : missing Type

	// UNSUPPORTED : C value 'set_anchor' : missing Type

	// UNSUPPORTED : C value 'insert_at_cursor' : missing Type

	// UNSUPPORTED : C value 'delete_from_cursor' : missing Type

	// UNSUPPORTED : C value 'backspace' : missing Type

	// UNSUPPORTED : C value 'cut_clipboard' : missing Type

	// UNSUPPORTED : C value 'copy_clipboard' : missing Type

	// UNSUPPORTED : C value 'paste_clipboard' : missing Type

	// UNSUPPORTED : C value 'toggle_overwrite' : missing Type

	// UNSUPPORTED : C value 'create_buffer' : missing Type

	// UNSUPPORTED : C value 'draw_layer' : missing Type

	// UNSUPPORTED : C value 'extend_selection' : missing Type

	// UNSUPPORTED : C value 'insert_emoji' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TextViewPrivate struct {
	native uintptr
}

type ThemeEngine struct {
	native uintptr
}

type ThemingEngineClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'render_line' : missing Type

	// UNSUPPORTED : C value 'render_background' : missing Type

	// UNSUPPORTED : C value 'render_frame' : missing Type

	// UNSUPPORTED : C value 'render_frame_gap' : missing Type

	// UNSUPPORTED : C value 'render_extension' : missing Type

	// UNSUPPORTED : C value 'render_check' : missing Type

	// UNSUPPORTED : C value 'render_option' : missing Type

	// UNSUPPORTED : C value 'render_arrow' : missing Type

	// UNSUPPORTED : C value 'render_expander' : missing Type

	// UNSUPPORTED : C value 'render_focus' : missing Type

	// UNSUPPORTED : C value 'render_layout' : missing Type

	// UNSUPPORTED : C value 'render_slider' : missing Type

	// UNSUPPORTED : C value 'render_handle' : missing Type

	// UNSUPPORTED : C value 'render_activity' : missing Type

	// UNSUPPORTED : C value 'render_icon_pixbuf' : missing Type

	// UNSUPPORTED : C value 'render_icon' : missing Type

	// UNSUPPORTED : C value 'render_icon_surface' : missing Type

}

type ThemingEnginePrivate struct {
	native uintptr
}

type ToggleActionClass struct {
	native      uintptr
	ParentClass *ActionClass
	// UNSUPPORTED : C value 'toggled' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ToggleActionEntry struct {
	native      uintptr
	Name        string
	StockId     string
	Label       string
	Accelerator string
	Tooltip     string
	// UNSUPPORTED : C value 'callback' : no Go type for 'GObject.Callback'

	// UNSUPPORTED : C value 'is_active' : no Go type for 'gboolean'

}

type ToggleActionPrivate struct {
	native uintptr
}

type ToggleButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

type ToggleButtonAccessiblePrivate struct {
	native uintptr
}

type ToggleButtonClass struct {
	native      uintptr
	ParentClass *ButtonClass
	// UNSUPPORTED : C value 'toggled' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ToggleButtonPrivate struct {
	native uintptr
}

type ToggleToolButtonClass struct {
	native      uintptr
	ParentClass *ToolButtonClass
	// UNSUPPORTED : C value 'toggled' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ToggleToolButtonPrivate struct {
	native uintptr
}

type ToolButtonClass struct {
	native      uintptr
	ParentClass *ToolItemClass
	// UNSUPPORTED : C value 'button_type' : no Go type for 'GType'

	// UNSUPPORTED : C value 'clicked' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ToolButtonPrivate struct {
	native uintptr
}

type ToolItemClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'create_menu_proxy' : missing Type

	// UNSUPPORTED : C value 'toolbar_reconfigured' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ToolItemGroupClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ToolItemGroupPrivate struct {
	native uintptr
}

type ToolItemPrivate struct {
	native uintptr
}

type ToolPaletteClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ToolPalettePrivate struct {
	native uintptr
}

type ToolShellIface struct {
	native uintptr
	// UNSUPPORTED : C value 'get_icon_size' : missing Type

	// UNSUPPORTED : C value 'get_orientation' : missing Type

	// UNSUPPORTED : C value 'get_style' : missing Type

	// UNSUPPORTED : C value 'get_relief_style' : missing Type

	// UNSUPPORTED : C value 'rebuild_menu' : missing Type

	// UNSUPPORTED : C value 'get_text_orientation' : missing Type

	// UNSUPPORTED : C value 'get_text_alignment' : missing Type

	// UNSUPPORTED : C value 'get_ellipsize_mode' : missing Type

	// UNSUPPORTED : C value 'get_text_size_group' : missing Type

}

type ToolbarClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value 'orientation_changed' : missing Type

	// UNSUPPORTED : C value 'style_changed' : missing Type

	// UNSUPPORTED : C value 'popup_context_menu' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ToolbarPrivate struct {
	native uintptr
}

type ToplevelAccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Atk.ObjectClass'

}

type ToplevelAccessiblePrivate struct {
	native uintptr
}

type TreeDragDestIface struct {
	native uintptr
	// UNSUPPORTED : C value 'drag_data_received' : missing Type

	// UNSUPPORTED : C value 'row_drop_possible' : missing Type

}

type TreeDragSourceIface struct {
	native uintptr
	// UNSUPPORTED : C value 'row_draggable' : missing Type

	// UNSUPPORTED : C value 'drag_data_get' : missing Type

	// UNSUPPORTED : C value 'drag_data_delete' : missing Type

}

type TreeIter struct {
	native uintptr
	Stamp  int32
	// UNSUPPORTED : C value 'user_data' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'user_data2' : no Go type for 'gpointer'

	// UNSUPPORTED : C value 'user_data3' : no Go type for 'gpointer'

}

type TreeModelFilterClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'visible' : missing Type

	// UNSUPPORTED : C value 'modify' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TreeModelFilterPrivate struct {
	native uintptr
}

type TreeModelIface struct {
	native uintptr
	// UNSUPPORTED : C value 'row_changed' : missing Type

	// UNSUPPORTED : C value 'row_inserted' : missing Type

	// UNSUPPORTED : C value 'row_has_child_toggled' : missing Type

	// UNSUPPORTED : C value 'row_deleted' : missing Type

	// UNSUPPORTED : C value 'rows_reordered' : missing Type

	// UNSUPPORTED : C value 'get_flags' : missing Type

	// UNSUPPORTED : C value 'get_n_columns' : missing Type

	// UNSUPPORTED : C value 'get_column_type' : missing Type

	// UNSUPPORTED : C value 'get_iter' : missing Type

	// UNSUPPORTED : C value 'get_path' : missing Type

	// UNSUPPORTED : C value 'get_value' : missing Type

	// UNSUPPORTED : C value 'iter_next' : missing Type

	// UNSUPPORTED : C value 'iter_previous' : missing Type

	// UNSUPPORTED : C value 'iter_children' : missing Type

	// UNSUPPORTED : C value 'iter_has_child' : missing Type

	// UNSUPPORTED : C value 'iter_n_children' : missing Type

	// UNSUPPORTED : C value 'iter_nth_child' : missing Type

	// UNSUPPORTED : C value 'iter_parent' : missing Type

	// UNSUPPORTED : C value 'ref_node' : missing Type

	// UNSUPPORTED : C value 'unref_node' : missing Type

}

type TreeModelSortClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TreeModelSortPrivate struct {
	native uintptr
}

type TreePath struct {
	native uintptr
}

var newTreePathInvoker *gi.Function

// TreePathNew is a representation of the C type gtk_tree_path_new.
func TreePathNew() *TreePath {
	if newTreePathInvoker == nil {
		newTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "new")
	}

	ret := newTreePathInvoker.Invoke(nil, nil)

	return &TreePath{native: ret.Pointer()}
}

var newFirstTreePathInvoker *gi.Function

// TreePathNewFirst is a representation of the C type gtk_tree_path_new_first.
func TreePathNewFirst() *TreePath {
	if newFirstTreePathInvoker == nil {
		newFirstTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "new_first")
	}

	ret := newFirstTreePathInvoker.Invoke(nil, nil)

	return &TreePath{native: ret.Pointer()}
}

// UNSUPPORTED : C value 'gtk_tree_path_new_from_indices' : parameter '...' has no type

// UNSUPPORTED : C value 'gtk_tree_path_new_from_indicesv' : parameter 'indices' has no type

var newFromStringTreePathInvoker *gi.Function

// TreePathNewFromString is a representation of the C type gtk_tree_path_new_from_string.
func TreePathNewFromString(path string) *TreePath {
	if newFromStringTreePathInvoker == nil {
		newFromStringTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "new_from_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetString(path)

	ret := newFromStringTreePathInvoker.Invoke(inArgs[:], nil)

	return &TreePath{native: ret.Pointer()}
}

type TreeRowReference struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_tree_row_reference_new' : parameter 'model' of type 'TreeModel' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_new_proxy' : parameter 'proxy' of type 'GObject.Object' not supported

type TreeSelectionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'changed' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TreeSelectionPrivate struct {
	native uintptr
}

type TreeSortableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'sort_column_changed' : missing Type

	// UNSUPPORTED : C value 'get_sort_column_id' : missing Type

	// UNSUPPORTED : C value 'set_sort_column_id' : missing Type

	// UNSUPPORTED : C value 'set_sort_func' : missing Type

	// UNSUPPORTED : C value 'set_default_sort_func' : missing Type

	// UNSUPPORTED : C value 'has_default_sort_func' : missing Type

}

type TreeStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TreeStorePrivate struct {
	native uintptr
}

type TreeViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type TreeViewAccessiblePrivate struct {
	native uintptr
}

type TreeViewClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value 'row_activated' : missing Type

	// UNSUPPORTED : C value 'test_expand_row' : missing Type

	// UNSUPPORTED : C value 'test_collapse_row' : missing Type

	// UNSUPPORTED : C value 'row_expanded' : missing Type

	// UNSUPPORTED : C value 'row_collapsed' : missing Type

	// UNSUPPORTED : C value 'columns_changed' : missing Type

	// UNSUPPORTED : C value 'cursor_changed' : missing Type

	// UNSUPPORTED : C value 'move_cursor' : missing Type

	// UNSUPPORTED : C value 'select_all' : missing Type

	// UNSUPPORTED : C value 'unselect_all' : missing Type

	// UNSUPPORTED : C value 'select_cursor_row' : missing Type

	// UNSUPPORTED : C value 'toggle_cursor_row' : missing Type

	// UNSUPPORTED : C value 'expand_collapse_cursor_row' : missing Type

	// UNSUPPORTED : C value 'select_cursor_parent' : missing Type

	// UNSUPPORTED : C value 'start_interactive_search' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved8' : missing Type

}

type TreeViewColumnClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'

	// UNSUPPORTED : C value 'clicked' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type TreeViewColumnPrivate struct {
	native uintptr
}

type TreeViewPrivate struct {
	native uintptr
}

type UIManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'add_widget' : missing Type

	// UNSUPPORTED : C value 'actions_changed' : missing Type

	// UNSUPPORTED : C value 'connect_proxy' : missing Type

	// UNSUPPORTED : C value 'disconnect_proxy' : missing Type

	// UNSUPPORTED : C value 'pre_activate' : missing Type

	// UNSUPPORTED : C value 'post_activate' : missing Type

	// UNSUPPORTED : C value 'get_widget' : missing Type

	// UNSUPPORTED : C value 'get_action' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type UIManagerPrivate struct {
	native uintptr
}

type VBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
}

type VButtonBoxClass struct {
	native      uintptr
	ParentClass *ButtonBoxClass
}

type VPanedClass struct {
	native      uintptr
	ParentClass *PanedClass
}

type VScaleClass struct {
	native      uintptr
	ParentClass *ScaleClass
}

type VScrollbarClass struct {
	native      uintptr
	ParentClass *ScrollbarClass
}

type VSeparatorClass struct {
	native      uintptr
	ParentClass *SeparatorClass
}

type ViewportClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type ViewportPrivate struct {
	native uintptr
}

type VolumeButtonClass struct {
	native      uintptr
	ParentClass *ScaleButtonClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type WidgetAccessibleClass struct {
	native      uintptr
	ParentClass *AccessibleClass
	// UNSUPPORTED : C value 'notify_gtk' : missing Type

}

type WidgetAccessiblePrivate struct {
	native uintptr
}

type WidgetClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.InitiallyUnownedClass'

	ActivateSignal uint32
	// UNSUPPORTED : C value 'dispatch_child_properties_changed' : missing Type

	// UNSUPPORTED : C value 'destroy' : missing Type

	// UNSUPPORTED : C value 'show' : missing Type

	// UNSUPPORTED : C value 'show_all' : missing Type

	// UNSUPPORTED : C value 'hide' : missing Type

	// UNSUPPORTED : C value 'map' : missing Type

	// UNSUPPORTED : C value 'unmap' : missing Type

	// UNSUPPORTED : C value 'realize' : missing Type

	// UNSUPPORTED : C value 'unrealize' : missing Type

	// UNSUPPORTED : C value 'size_allocate' : missing Type

	// UNSUPPORTED : C value 'state_changed' : missing Type

	// UNSUPPORTED : C value 'state_flags_changed' : missing Type

	// UNSUPPORTED : C value 'parent_set' : missing Type

	// UNSUPPORTED : C value 'hierarchy_changed' : missing Type

	// UNSUPPORTED : C value 'style_set' : missing Type

	// UNSUPPORTED : C value 'direction_changed' : missing Type

	// UNSUPPORTED : C value 'grab_notify' : missing Type

	// UNSUPPORTED : C value 'child_notify' : missing Type

	// UNSUPPORTED : C value 'draw' : missing Type

	// UNSUPPORTED : C value 'get_request_mode' : missing Type

	// UNSUPPORTED : C value 'get_preferred_height' : missing Type

	// UNSUPPORTED : C value 'get_preferred_width_for_height' : missing Type

	// UNSUPPORTED : C value 'get_preferred_width' : missing Type

	// UNSUPPORTED : C value 'get_preferred_height_for_width' : missing Type

	// UNSUPPORTED : C value 'mnemonic_activate' : missing Type

	// UNSUPPORTED : C value 'grab_focus' : missing Type

	// UNSUPPORTED : C value 'focus' : missing Type

	// UNSUPPORTED : C value 'move_focus' : missing Type

	// UNSUPPORTED : C value 'keynav_failed' : missing Type

	// UNSUPPORTED : C value 'event' : missing Type

	// UNSUPPORTED : C value 'button_press_event' : missing Type

	// UNSUPPORTED : C value 'button_release_event' : missing Type

	// UNSUPPORTED : C value 'scroll_event' : missing Type

	// UNSUPPORTED : C value 'motion_notify_event' : missing Type

	// UNSUPPORTED : C value 'delete_event' : missing Type

	// UNSUPPORTED : C value 'destroy_event' : missing Type

	// UNSUPPORTED : C value 'key_press_event' : missing Type

	// UNSUPPORTED : C value 'key_release_event' : missing Type

	// UNSUPPORTED : C value 'enter_notify_event' : missing Type

	// UNSUPPORTED : C value 'leave_notify_event' : missing Type

	// UNSUPPORTED : C value 'configure_event' : missing Type

	// UNSUPPORTED : C value 'focus_in_event' : missing Type

	// UNSUPPORTED : C value 'focus_out_event' : missing Type

	// UNSUPPORTED : C value 'map_event' : missing Type

	// UNSUPPORTED : C value 'unmap_event' : missing Type

	// UNSUPPORTED : C value 'property_notify_event' : missing Type

	// UNSUPPORTED : C value 'selection_clear_event' : missing Type

	// UNSUPPORTED : C value 'selection_request_event' : missing Type

	// UNSUPPORTED : C value 'selection_notify_event' : missing Type

	// UNSUPPORTED : C value 'proximity_in_event' : missing Type

	// UNSUPPORTED : C value 'proximity_out_event' : missing Type

	// UNSUPPORTED : C value 'visibility_notify_event' : missing Type

	// UNSUPPORTED : C value 'window_state_event' : missing Type

	// UNSUPPORTED : C value 'damage_event' : missing Type

	// UNSUPPORTED : C value 'grab_broken_event' : missing Type

	// UNSUPPORTED : C value 'selection_get' : missing Type

	// UNSUPPORTED : C value 'selection_received' : missing Type

	// UNSUPPORTED : C value 'drag_begin' : missing Type

	// UNSUPPORTED : C value 'drag_end' : missing Type

	// UNSUPPORTED : C value 'drag_data_get' : missing Type

	// UNSUPPORTED : C value 'drag_data_delete' : missing Type

	// UNSUPPORTED : C value 'drag_leave' : missing Type

	// UNSUPPORTED : C value 'drag_motion' : missing Type

	// UNSUPPORTED : C value 'drag_drop' : missing Type

	// UNSUPPORTED : C value 'drag_data_received' : missing Type

	// UNSUPPORTED : C value 'drag_failed' : missing Type

	// UNSUPPORTED : C value 'popup_menu' : missing Type

	// UNSUPPORTED : C value 'show_help' : missing Type

	// UNSUPPORTED : C value 'get_accessible' : missing Type

	// UNSUPPORTED : C value 'screen_changed' : missing Type

	// UNSUPPORTED : C value 'can_activate_accel' : missing Type

	// UNSUPPORTED : C value 'composited_changed' : missing Type

	// UNSUPPORTED : C value 'query_tooltip' : missing Type

	// UNSUPPORTED : C value 'compute_expand' : missing Type

	// UNSUPPORTED : C value 'adjust_size_request' : missing Type

	// UNSUPPORTED : C value 'adjust_size_allocation' : missing Type

	// UNSUPPORTED : C value 'style_updated' : missing Type

	// UNSUPPORTED : C value 'touch_event' : missing Type

	// UNSUPPORTED : C value 'get_preferred_height_and_baseline_for_width' : missing Type

	// UNSUPPORTED : C value 'adjust_baseline_request' : missing Type

	// UNSUPPORTED : C value 'adjust_baseline_allocation' : missing Type

	// UNSUPPORTED : C value 'queue_draw_region' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved7' : missing Type

}

type WidgetClassPrivate struct {
	native uintptr
}

type WidgetPath struct {
	native uintptr
}

var newWidgetPathInvoker *gi.Function

// WidgetPathNew is a representation of the C type gtk_widget_path_new.
func WidgetPathNew() *WidgetPath {
	if newWidgetPathInvoker == nil {
		newWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "new")
	}

	ret := newWidgetPathInvoker.Invoke(nil, nil)

	return &WidgetPath{native: ret.Pointer()}
}

type WidgetPrivate struct {
	native uintptr
}

type WindowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

type WindowAccessiblePrivate struct {
	native uintptr
}

type WindowClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'set_focus' : missing Type

	// UNSUPPORTED : C value 'activate_focus' : missing Type

	// UNSUPPORTED : C value 'activate_default' : missing Type

	// UNSUPPORTED : C value 'keys_changed' : missing Type

	// UNSUPPORTED : C value 'enable_debugging' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

}

type WindowGeometryInfo struct {
	native uintptr
}

type WindowGroupClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type

}

type WindowGroupPrivate struct {
	native uintptr
}

type WindowPrivate struct {
	native uintptr
}
