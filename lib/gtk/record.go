// Code generated - DO NOT EDIT.

package gtk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var AboutDialogClassStruct *gi.Struct
var AboutDialogClassStructOnce sync.Once

func AboutDialogClassStructSet() {
	AboutDialogClassStructOnce.Do(func() {
		AboutDialogClassStruct = gi.StructNew("Gtk", "AboutDialogClass")
	})
}

type AboutDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value 'activate_link' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var AboutDialogPrivateStruct *gi.Struct
var AboutDialogPrivateStructOnce sync.Once

func AboutDialogPrivateStructSet() {
	AboutDialogPrivateStructOnce.Do(func() {
		AboutDialogPrivateStruct = gi.StructNew("Gtk", "AboutDialogPrivate")
	})
}

type AboutDialogPrivate struct {
	native uintptr
}

var AccelGroupClassStruct *gi.Struct
var AccelGroupClassStructOnce sync.Once

func AccelGroupClassStructSet() {
	AccelGroupClassStructOnce.Do(func() {
		AccelGroupClassStruct = gi.StructNew("Gtk", "AccelGroupClass")
	})
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

var AccelGroupEntryStruct *gi.Struct
var AccelGroupEntryStructOnce sync.Once

func AccelGroupEntryStructSet() {
	AccelGroupEntryStructOnce.Do(func() {
		AccelGroupEntryStruct = gi.StructNew("Gtk", "AccelGroupEntry")
	})
}

type AccelGroupEntry struct {
	native uintptr
	Key    *AccelKey
	// UNSUPPORTED : C value 'closure' : no Go type for 'GObject.Closure'
	// UNSUPPORTED : C value 'accel_path_quark' : no Go type for 'GLib.Quark'
}

var AccelGroupPrivateStruct *gi.Struct
var AccelGroupPrivateStructOnce sync.Once

func AccelGroupPrivateStructSet() {
	AccelGroupPrivateStructOnce.Do(func() {
		AccelGroupPrivateStruct = gi.StructNew("Gtk", "AccelGroupPrivate")
	})
}

type AccelGroupPrivate struct {
	native uintptr
}

var AccelKeyStruct *gi.Struct
var AccelKeyStructOnce sync.Once

func AccelKeyStructSet() {
	AccelKeyStructOnce.Do(func() {
		AccelKeyStruct = gi.StructNew("Gtk", "AccelKey")
	})
}

type AccelKey struct {
	native   uintptr
	AccelKey uint32
	// UNSUPPORTED : C value 'accel_mods' : no Go type for 'Gdk.ModifierType'
	AccelFlags uint32
}

var AccelLabelClassStruct *gi.Struct
var AccelLabelClassStructOnce sync.Once

func AccelLabelClassStructSet() {
	AccelLabelClassStructOnce.Do(func() {
		AccelLabelClassStruct = gi.StructNew("Gtk", "AccelLabelClass")
	})
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

var AccelLabelPrivateStruct *gi.Struct
var AccelLabelPrivateStructOnce sync.Once

func AccelLabelPrivateStructSet() {
	AccelLabelPrivateStructOnce.Do(func() {
		AccelLabelPrivateStruct = gi.StructNew("Gtk", "AccelLabelPrivate")
	})
}

type AccelLabelPrivate struct {
	native uintptr
}

var AccelMapClassStruct *gi.Struct
var AccelMapClassStructOnce sync.Once

func AccelMapClassStructSet() {
	AccelMapClassStructOnce.Do(func() {
		AccelMapClassStruct = gi.StructNew("Gtk", "AccelMapClass")
	})
}

type AccelMapClass struct {
	native uintptr
}

var AccessibleClassStruct *gi.Struct
var AccessibleClassStructOnce sync.Once

func AccessibleClassStructSet() {
	AccessibleClassStructOnce.Do(func() {
		AccessibleClassStruct = gi.StructNew("Gtk", "AccessibleClass")
	})
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

var AccessiblePrivateStruct *gi.Struct
var AccessiblePrivateStructOnce sync.Once

func AccessiblePrivateStructSet() {
	AccessiblePrivateStructOnce.Do(func() {
		AccessiblePrivateStruct = gi.StructNew("Gtk", "AccessiblePrivate")
	})
}

type AccessiblePrivate struct {
	native uintptr
}

var ActionBarClassStruct *gi.Struct
var ActionBarClassStructOnce sync.Once

func ActionBarClassStructSet() {
	ActionBarClassStructOnce.Do(func() {
		ActionBarClassStruct = gi.StructNew("Gtk", "ActionBarClass")
	})
}

type ActionBarClass struct {
	native uintptr
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ActionBarPrivateStruct *gi.Struct
var ActionBarPrivateStructOnce sync.Once

func ActionBarPrivateStructSet() {
	ActionBarPrivateStructOnce.Do(func() {
		ActionBarPrivateStruct = gi.StructNew("Gtk", "ActionBarPrivate")
	})
}

type ActionBarPrivate struct {
	native uintptr
}

var ActionClassStruct *gi.Struct
var ActionClassStructOnce sync.Once

func ActionClassStructSet() {
	ActionClassStructOnce.Do(func() {
		ActionClassStruct = gi.StructNew("Gtk", "ActionClass")
	})
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

var ActionEntryStruct *gi.Struct
var ActionEntryStructOnce sync.Once

func ActionEntryStructSet() {
	ActionEntryStructOnce.Do(func() {
		ActionEntryStruct = gi.StructNew("Gtk", "ActionEntry")
	})
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

var ActionGroupClassStruct *gi.Struct
var ActionGroupClassStructOnce sync.Once

func ActionGroupClassStructSet() {
	ActionGroupClassStructOnce.Do(func() {
		ActionGroupClassStruct = gi.StructNew("Gtk", "ActionGroupClass")
	})
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

var ActionGroupPrivateStruct *gi.Struct
var ActionGroupPrivateStructOnce sync.Once

func ActionGroupPrivateStructSet() {
	ActionGroupPrivateStructOnce.Do(func() {
		ActionGroupPrivateStruct = gi.StructNew("Gtk", "ActionGroupPrivate")
	})
}

type ActionGroupPrivate struct {
	native uintptr
}

var ActionPrivateStruct *gi.Struct
var ActionPrivateStructOnce sync.Once

func ActionPrivateStructSet() {
	ActionPrivateStructOnce.Do(func() {
		ActionPrivateStruct = gi.StructNew("Gtk", "ActionPrivate")
	})
}

type ActionPrivate struct {
	native uintptr
}

var ActionableInterfaceStruct *gi.Struct
var ActionableInterfaceStructOnce sync.Once

func ActionableInterfaceStructSet() {
	ActionableInterfaceStructOnce.Do(func() {
		ActionableInterfaceStruct = gi.StructNew("Gtk", "ActionableInterface")
	})
}

type ActionableInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'get_action_name' : missing Type
	// UNSUPPORTED : C value 'set_action_name' : missing Type
	// UNSUPPORTED : C value 'get_action_target_value' : missing Type
	// UNSUPPORTED : C value 'set_action_target_value' : missing Type
}

var ActivatableIfaceStruct *gi.Struct
var ActivatableIfaceStructOnce sync.Once

func ActivatableIfaceStructSet() {
	ActivatableIfaceStructOnce.Do(func() {
		ActivatableIfaceStruct = gi.StructNew("Gtk", "ActivatableIface")
	})
}

type ActivatableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'update' : missing Type
	// UNSUPPORTED : C value 'sync_action_properties' : missing Type
}

var AdjustmentClassStruct *gi.Struct
var AdjustmentClassStructOnce sync.Once

func AdjustmentClassStructSet() {
	AdjustmentClassStructOnce.Do(func() {
		AdjustmentClassStruct = gi.StructNew("Gtk", "AdjustmentClass")
	})
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

var AdjustmentPrivateStruct *gi.Struct
var AdjustmentPrivateStructOnce sync.Once

func AdjustmentPrivateStructSet() {
	AdjustmentPrivateStructOnce.Do(func() {
		AdjustmentPrivateStruct = gi.StructNew("Gtk", "AdjustmentPrivate")
	})
}

type AdjustmentPrivate struct {
	native uintptr
}

var AlignmentClassStruct *gi.Struct
var AlignmentClassStructOnce sync.Once

func AlignmentClassStructSet() {
	AlignmentClassStructOnce.Do(func() {
		AlignmentClassStruct = gi.StructNew("Gtk", "AlignmentClass")
	})
}

type AlignmentClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var AlignmentPrivateStruct *gi.Struct
var AlignmentPrivateStructOnce sync.Once

func AlignmentPrivateStructSet() {
	AlignmentPrivateStructOnce.Do(func() {
		AlignmentPrivateStruct = gi.StructNew("Gtk", "AlignmentPrivate")
	})
}

type AlignmentPrivate struct {
	native uintptr
}

var AppChooserButtonClassStruct *gi.Struct
var AppChooserButtonClassStructOnce sync.Once

func AppChooserButtonClassStructSet() {
	AppChooserButtonClassStructOnce.Do(func() {
		AppChooserButtonClassStruct = gi.StructNew("Gtk", "AppChooserButtonClass")
	})
}

type AppChooserButtonClass struct {
	native      uintptr
	ParentClass *ComboBoxClass
	// UNSUPPORTED : C value 'custom_item_activated' : missing Type
}

var AppChooserButtonPrivateStruct *gi.Struct
var AppChooserButtonPrivateStructOnce sync.Once

func AppChooserButtonPrivateStructSet() {
	AppChooserButtonPrivateStructOnce.Do(func() {
		AppChooserButtonPrivateStruct = gi.StructNew("Gtk", "AppChooserButtonPrivate")
	})
}

type AppChooserButtonPrivate struct {
	native uintptr
}

var AppChooserDialogClassStruct *gi.Struct
var AppChooserDialogClassStructOnce sync.Once

func AppChooserDialogClassStructSet() {
	AppChooserDialogClassStructOnce.Do(func() {
		AppChooserDialogClassStruct = gi.StructNew("Gtk", "AppChooserDialogClass")
	})
}

type AppChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
}

var AppChooserDialogPrivateStruct *gi.Struct
var AppChooserDialogPrivateStructOnce sync.Once

func AppChooserDialogPrivateStructSet() {
	AppChooserDialogPrivateStructOnce.Do(func() {
		AppChooserDialogPrivateStruct = gi.StructNew("Gtk", "AppChooserDialogPrivate")
	})
}

type AppChooserDialogPrivate struct {
	native uintptr
}

var AppChooserWidgetClassStruct *gi.Struct
var AppChooserWidgetClassStructOnce sync.Once

func AppChooserWidgetClassStructSet() {
	AppChooserWidgetClassStructOnce.Do(func() {
		AppChooserWidgetClassStruct = gi.StructNew("Gtk", "AppChooserWidgetClass")
	})
}

type AppChooserWidgetClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value 'application_selected' : missing Type
	// UNSUPPORTED : C value 'application_activated' : missing Type
	// UNSUPPORTED : C value 'populate_popup' : missing Type
}

var AppChooserWidgetPrivateStruct *gi.Struct
var AppChooserWidgetPrivateStructOnce sync.Once

func AppChooserWidgetPrivateStructSet() {
	AppChooserWidgetPrivateStructOnce.Do(func() {
		AppChooserWidgetPrivateStruct = gi.StructNew("Gtk", "AppChooserWidgetPrivate")
	})
}

type AppChooserWidgetPrivate struct {
	native uintptr
}

var ApplicationClassStruct *gi.Struct
var ApplicationClassStructOnce sync.Once

func ApplicationClassStructSet() {
	ApplicationClassStructOnce.Do(func() {
		ApplicationClassStruct = gi.StructNew("Gtk", "ApplicationClass")
	})
}

type ApplicationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.ApplicationClass'
	// UNSUPPORTED : C value 'window_added' : missing Type
	// UNSUPPORTED : C value 'window_removed' : missing Type
}

var ApplicationPrivateStruct *gi.Struct
var ApplicationPrivateStructOnce sync.Once

func ApplicationPrivateStructSet() {
	ApplicationPrivateStructOnce.Do(func() {
		ApplicationPrivateStruct = gi.StructNew("Gtk", "ApplicationPrivate")
	})
}

type ApplicationPrivate struct {
	native uintptr
}

var ApplicationWindowClassStruct *gi.Struct
var ApplicationWindowClassStructOnce sync.Once

func ApplicationWindowClassStructSet() {
	ApplicationWindowClassStructOnce.Do(func() {
		ApplicationWindowClassStruct = gi.StructNew("Gtk", "ApplicationWindowClass")
	})
}

type ApplicationWindowClass struct {
	native      uintptr
	ParentClass *WindowClass
}

var ApplicationWindowPrivateStruct *gi.Struct
var ApplicationWindowPrivateStructOnce sync.Once

func ApplicationWindowPrivateStructSet() {
	ApplicationWindowPrivateStructOnce.Do(func() {
		ApplicationWindowPrivateStruct = gi.StructNew("Gtk", "ApplicationWindowPrivate")
	})
}

type ApplicationWindowPrivate struct {
	native uintptr
}

var ArrowAccessibleClassStruct *gi.Struct
var ArrowAccessibleClassStructOnce sync.Once

func ArrowAccessibleClassStructSet() {
	ArrowAccessibleClassStructOnce.Do(func() {
		ArrowAccessibleClassStruct = gi.StructNew("Gtk", "ArrowAccessibleClass")
	})
}

type ArrowAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var ArrowAccessiblePrivateStruct *gi.Struct
var ArrowAccessiblePrivateStructOnce sync.Once

func ArrowAccessiblePrivateStructSet() {
	ArrowAccessiblePrivateStructOnce.Do(func() {
		ArrowAccessiblePrivateStruct = gi.StructNew("Gtk", "ArrowAccessiblePrivate")
	})
}

type ArrowAccessiblePrivate struct {
	native uintptr
}

var ArrowClassStruct *gi.Struct
var ArrowClassStructOnce sync.Once

func ArrowClassStructSet() {
	ArrowClassStructOnce.Do(func() {
		ArrowClassStruct = gi.StructNew("Gtk", "ArrowClass")
	})
}

type ArrowClass struct {
	native      uintptr
	ParentClass *MiscClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ArrowPrivateStruct *gi.Struct
var ArrowPrivateStructOnce sync.Once

func ArrowPrivateStructSet() {
	ArrowPrivateStructOnce.Do(func() {
		ArrowPrivateStruct = gi.StructNew("Gtk", "ArrowPrivate")
	})
}

type ArrowPrivate struct {
	native uintptr
}

var AspectFrameClassStruct *gi.Struct
var AspectFrameClassStructOnce sync.Once

func AspectFrameClassStructSet() {
	AspectFrameClassStructOnce.Do(func() {
		AspectFrameClassStruct = gi.StructNew("Gtk", "AspectFrameClass")
	})
}

type AspectFrameClass struct {
	native      uintptr
	ParentClass *FrameClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var AspectFramePrivateStruct *gi.Struct
var AspectFramePrivateStructOnce sync.Once

func AspectFramePrivateStructSet() {
	AspectFramePrivateStructOnce.Do(func() {
		AspectFramePrivateStruct = gi.StructNew("Gtk", "AspectFramePrivate")
	})
}

type AspectFramePrivate struct {
	native uintptr
}

var AssistantClassStruct *gi.Struct
var AssistantClassStructOnce sync.Once

func AssistantClassStructSet() {
	AssistantClassStructOnce.Do(func() {
		AssistantClassStruct = gi.StructNew("Gtk", "AssistantClass")
	})
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

var AssistantPrivateStruct *gi.Struct
var AssistantPrivateStructOnce sync.Once

func AssistantPrivateStructSet() {
	AssistantPrivateStructOnce.Do(func() {
		AssistantPrivateStruct = gi.StructNew("Gtk", "AssistantPrivate")
	})
}

type AssistantPrivate struct {
	native uintptr
}

var BinClassStruct *gi.Struct
var BinClassStructOnce sync.Once

func BinClassStructSet() {
	BinClassStructOnce.Do(func() {
		BinClassStruct = gi.StructNew("Gtk", "BinClass")
	})
}

type BinClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var BinPrivateStruct *gi.Struct
var BinPrivateStructOnce sync.Once

func BinPrivateStructSet() {
	BinPrivateStructOnce.Do(func() {
		BinPrivateStruct = gi.StructNew("Gtk", "BinPrivate")
	})
}

type BinPrivate struct {
	native uintptr
}

var BindingArgStruct *gi.Struct
var BindingArgStructOnce sync.Once

func BindingArgStructSet() {
	BindingArgStructOnce.Do(func() {
		BindingArgStruct = gi.StructNew("Gtk", "BindingArg")
	})
}

type BindingArg struct {
	native uintptr
	// UNSUPPORTED : C value 'arg_type' : no Go type for 'GType'
}

var BindingEntryStruct *gi.Struct
var BindingEntryStructOnce sync.Once

func BindingEntryStructSet() {
	BindingEntryStructOnce.Do(func() {
		BindingEntryStruct = gi.StructNew("Gtk", "BindingEntry")
	})
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

var BindingSetStruct *gi.Struct
var BindingSetStructOnce sync.Once

func BindingSetStructSet() {
	BindingSetStructOnce.Do(func() {
		BindingSetStruct = gi.StructNew("Gtk", "BindingSet")
	})
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

// UNSUPPORTED : C value 'gtk_binding_set_activate' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_binding_set_add_path' : parameter 'path_type' of type 'PathType' not supported

var BindingSignalStruct *gi.Struct
var BindingSignalStructOnce sync.Once

func BindingSignalStructSet() {
	BindingSignalStructOnce.Do(func() {
		BindingSignalStruct = gi.StructNew("Gtk", "BindingSignal")
	})
}

type BindingSignal struct {
	native     uintptr
	Next       *BindingSignal
	SignalName string
	NArgs      uint32
	// UNSUPPORTED : C value 'args' : missing Type
}

var BooleanCellAccessibleClassStruct *gi.Struct
var BooleanCellAccessibleClassStructOnce sync.Once

func BooleanCellAccessibleClassStructSet() {
	BooleanCellAccessibleClassStructOnce.Do(func() {
		BooleanCellAccessibleClassStruct = gi.StructNew("Gtk", "BooleanCellAccessibleClass")
	})
}

type BooleanCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

var BooleanCellAccessiblePrivateStruct *gi.Struct
var BooleanCellAccessiblePrivateStructOnce sync.Once

func BooleanCellAccessiblePrivateStructSet() {
	BooleanCellAccessiblePrivateStructOnce.Do(func() {
		BooleanCellAccessiblePrivateStruct = gi.StructNew("Gtk", "BooleanCellAccessiblePrivate")
	})
}

type BooleanCellAccessiblePrivate struct {
	native uintptr
}

var BorderStruct *gi.Struct
var BorderStructOnce sync.Once

func BorderStructSet() {
	BorderStructOnce.Do(func() {
		BorderStruct = gi.StructNew("Gtk", "Border")
	})
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

	retGo := &Border{native: ret.Pointer()}

	return retGo
}

var copyBorderInvoker *gi.Function

// Copy is a representation of the C type gtk_border_copy.
func (recv *Border) Copy() *Border {
	if copyBorderInvoker == nil {
		copyBorderInvoker = gi.StructFunctionInvokerNew("Gtk", "Border", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyBorderInvoker.Invoke(inArgs[:], nil)

	retGo := &Border{native: ret.Pointer()}

	return retGo
}

var freeBorderInvoker *gi.Function

// Free is a representation of the C type gtk_border_free.
func (recv *Border) Free() {
	if freeBorderInvoker == nil {
		freeBorderInvoker = gi.StructFunctionInvokerNew("Gtk", "Border", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeBorderInvoker.Invoke(inArgs[:], nil)

}

var BoxClassStruct *gi.Struct
var BoxClassStructOnce sync.Once

func BoxClassStructSet() {
	BoxClassStructOnce.Do(func() {
		BoxClassStruct = gi.StructNew("Gtk", "BoxClass")
	})
}

type BoxClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var BoxPrivateStruct *gi.Struct
var BoxPrivateStructOnce sync.Once

func BoxPrivateStructSet() {
	BoxPrivateStructOnce.Do(func() {
		BoxPrivateStruct = gi.StructNew("Gtk", "BoxPrivate")
	})
}

type BoxPrivate struct {
	native uintptr
}

var BuildableIfaceStruct *gi.Struct
var BuildableIfaceStructOnce sync.Once

func BuildableIfaceStructSet() {
	BuildableIfaceStructOnce.Do(func() {
		BuildableIfaceStruct = gi.StructNew("Gtk", "BuildableIface")
	})
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

var BuilderClassStruct *gi.Struct
var BuilderClassStructOnce sync.Once

func BuilderClassStructSet() {
	BuilderClassStructOnce.Do(func() {
		BuilderClassStruct = gi.StructNew("Gtk", "BuilderClass")
	})
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

var BuilderPrivateStruct *gi.Struct
var BuilderPrivateStructOnce sync.Once

func BuilderPrivateStructSet() {
	BuilderPrivateStructOnce.Do(func() {
		BuilderPrivateStruct = gi.StructNew("Gtk", "BuilderPrivate")
	})
}

type BuilderPrivate struct {
	native uintptr
}

var ButtonAccessibleClassStruct *gi.Struct
var ButtonAccessibleClassStructOnce sync.Once

func ButtonAccessibleClassStructSet() {
	ButtonAccessibleClassStructOnce.Do(func() {
		ButtonAccessibleClassStruct = gi.StructNew("Gtk", "ButtonAccessibleClass")
	})
}

type ButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var ButtonAccessiblePrivateStruct *gi.Struct
var ButtonAccessiblePrivateStructOnce sync.Once

func ButtonAccessiblePrivateStructSet() {
	ButtonAccessiblePrivateStructOnce.Do(func() {
		ButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "ButtonAccessiblePrivate")
	})
}

type ButtonAccessiblePrivate struct {
	native uintptr
}

var ButtonBoxClassStruct *gi.Struct
var ButtonBoxClassStructOnce sync.Once

func ButtonBoxClassStructSet() {
	ButtonBoxClassStructOnce.Do(func() {
		ButtonBoxClassStruct = gi.StructNew("Gtk", "ButtonBoxClass")
	})
}

type ButtonBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ButtonBoxPrivateStruct *gi.Struct
var ButtonBoxPrivateStructOnce sync.Once

func ButtonBoxPrivateStructSet() {
	ButtonBoxPrivateStructOnce.Do(func() {
		ButtonBoxPrivateStruct = gi.StructNew("Gtk", "ButtonBoxPrivate")
	})
}

type ButtonBoxPrivate struct {
	native uintptr
}

var ButtonClassStruct *gi.Struct
var ButtonClassStructOnce sync.Once

func ButtonClassStructSet() {
	ButtonClassStructOnce.Do(func() {
		ButtonClassStruct = gi.StructNew("Gtk", "ButtonClass")
	})
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

var ButtonPrivateStruct *gi.Struct
var ButtonPrivateStructOnce sync.Once

func ButtonPrivateStructSet() {
	ButtonPrivateStructOnce.Do(func() {
		ButtonPrivateStruct = gi.StructNew("Gtk", "ButtonPrivate")
	})
}

type ButtonPrivate struct {
	native uintptr
}

var CalendarClassStruct *gi.Struct
var CalendarClassStructOnce sync.Once

func CalendarClassStructSet() {
	CalendarClassStructOnce.Do(func() {
		CalendarClassStruct = gi.StructNew("Gtk", "CalendarClass")
	})
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

var CalendarPrivateStruct *gi.Struct
var CalendarPrivateStructOnce sync.Once

func CalendarPrivateStructSet() {
	CalendarPrivateStructOnce.Do(func() {
		CalendarPrivateStruct = gi.StructNew("Gtk", "CalendarPrivate")
	})
}

type CalendarPrivate struct {
	native uintptr
}

var CellAccessibleClassStruct *gi.Struct
var CellAccessibleClassStructOnce sync.Once

func CellAccessibleClassStructSet() {
	CellAccessibleClassStructOnce.Do(func() {
		CellAccessibleClassStruct = gi.StructNew("Gtk", "CellAccessibleClass")
	})
}

type CellAccessibleClass struct {
	native      uintptr
	ParentClass *AccessibleClass
	// UNSUPPORTED : C value 'update_cache' : missing Type
}

var CellAccessibleParentIfaceStruct *gi.Struct
var CellAccessibleParentIfaceStructOnce sync.Once

func CellAccessibleParentIfaceStructSet() {
	CellAccessibleParentIfaceStructOnce.Do(func() {
		CellAccessibleParentIfaceStruct = gi.StructNew("Gtk", "CellAccessibleParentIface")
	})
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

var CellAccessiblePrivateStruct *gi.Struct
var CellAccessiblePrivateStructOnce sync.Once

func CellAccessiblePrivateStructSet() {
	CellAccessiblePrivateStructOnce.Do(func() {
		CellAccessiblePrivateStruct = gi.StructNew("Gtk", "CellAccessiblePrivate")
	})
}

type CellAccessiblePrivate struct {
	native uintptr
}

var CellAreaBoxClassStruct *gi.Struct
var CellAreaBoxClassStructOnce sync.Once

func CellAreaBoxClassStructSet() {
	CellAreaBoxClassStructOnce.Do(func() {
		CellAreaBoxClassStruct = gi.StructNew("Gtk", "CellAreaBoxClass")
	})
}

type CellAreaBoxClass struct {
	native uintptr
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var CellAreaBoxPrivateStruct *gi.Struct
var CellAreaBoxPrivateStructOnce sync.Once

func CellAreaBoxPrivateStructSet() {
	CellAreaBoxPrivateStructOnce.Do(func() {
		CellAreaBoxPrivateStruct = gi.StructNew("Gtk", "CellAreaBoxPrivate")
	})
}

type CellAreaBoxPrivate struct {
	native uintptr
}

var CellAreaClassStruct *gi.Struct
var CellAreaClassStructOnce sync.Once

func CellAreaClassStructSet() {
	CellAreaClassStructOnce.Do(func() {
		CellAreaClassStruct = gi.StructNew("Gtk", "CellAreaClass")
	})
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

// UNSUPPORTED : C value 'gtk_cell_area_class_find_cell_property' : return type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_cell_area_class_install_cell_property' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

var listCellPropertiesCellAreaClassInvoker *gi.Function

// ListCellProperties is a representation of the C type gtk_cell_area_class_list_cell_properties.
func (recv *CellAreaClass) ListCellProperties() uint32 {
	if listCellPropertiesCellAreaClassInvoker == nil {
		listCellPropertiesCellAreaClassInvoker = gi.StructFunctionInvokerNew("Gtk", "CellAreaClass", "list_cell_properties")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	listCellPropertiesCellAreaClassInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint32()

	return out0
}

var CellAreaContextClassStruct *gi.Struct
var CellAreaContextClassStructOnce sync.Once

func CellAreaContextClassStructSet() {
	CellAreaContextClassStructOnce.Do(func() {
		CellAreaContextClassStruct = gi.StructNew("Gtk", "CellAreaContextClass")
	})
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

var CellAreaContextPrivateStruct *gi.Struct
var CellAreaContextPrivateStructOnce sync.Once

func CellAreaContextPrivateStructSet() {
	CellAreaContextPrivateStructOnce.Do(func() {
		CellAreaContextPrivateStruct = gi.StructNew("Gtk", "CellAreaContextPrivate")
	})
}

type CellAreaContextPrivate struct {
	native uintptr
}

var CellAreaPrivateStruct *gi.Struct
var CellAreaPrivateStructOnce sync.Once

func CellAreaPrivateStructSet() {
	CellAreaPrivateStructOnce.Do(func() {
		CellAreaPrivateStruct = gi.StructNew("Gtk", "CellAreaPrivate")
	})
}

type CellAreaPrivate struct {
	native uintptr
}

var CellEditableIfaceStruct *gi.Struct
var CellEditableIfaceStructOnce sync.Once

func CellEditableIfaceStructSet() {
	CellEditableIfaceStructOnce.Do(func() {
		CellEditableIfaceStruct = gi.StructNew("Gtk", "CellEditableIface")
	})
}

type CellEditableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'editing_done' : missing Type
	// UNSUPPORTED : C value 'remove_widget' : missing Type
	// UNSUPPORTED : C value 'start_editing' : missing Type
}

var CellLayoutIfaceStruct *gi.Struct
var CellLayoutIfaceStructOnce sync.Once

func CellLayoutIfaceStructSet() {
	CellLayoutIfaceStructOnce.Do(func() {
		CellLayoutIfaceStruct = gi.StructNew("Gtk", "CellLayoutIface")
	})
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

var CellRendererAccelClassStruct *gi.Struct
var CellRendererAccelClassStructOnce sync.Once

func CellRendererAccelClassStructSet() {
	CellRendererAccelClassStructOnce.Do(func() {
		CellRendererAccelClassStruct = gi.StructNew("Gtk", "CellRendererAccelClass")
	})
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

var CellRendererAccelPrivateStruct *gi.Struct
var CellRendererAccelPrivateStructOnce sync.Once

func CellRendererAccelPrivateStructSet() {
	CellRendererAccelPrivateStructOnce.Do(func() {
		CellRendererAccelPrivateStruct = gi.StructNew("Gtk", "CellRendererAccelPrivate")
	})
}

type CellRendererAccelPrivate struct {
	native uintptr
}

var CellRendererClassStruct *gi.Struct
var CellRendererClassStructOnce sync.Once

func CellRendererClassStructSet() {
	CellRendererClassStructOnce.Do(func() {
		CellRendererClassStruct = gi.StructNew("Gtk", "CellRendererClass")
	})
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

// UNSUPPORTED : C value 'gtk_cell_renderer_class_set_accessible_type' : parameter 'type' of type 'GType' not supported

var CellRendererClassPrivateStruct *gi.Struct
var CellRendererClassPrivateStructOnce sync.Once

func CellRendererClassPrivateStructSet() {
	CellRendererClassPrivateStructOnce.Do(func() {
		CellRendererClassPrivateStruct = gi.StructNew("Gtk", "CellRendererClassPrivate")
	})
}

type CellRendererClassPrivate struct {
	native uintptr
}

var CellRendererComboClassStruct *gi.Struct
var CellRendererComboClassStructOnce sync.Once

func CellRendererComboClassStructSet() {
	CellRendererComboClassStructOnce.Do(func() {
		CellRendererComboClassStruct = gi.StructNew("Gtk", "CellRendererComboClass")
	})
}

type CellRendererComboClass struct {
	native uintptr
	Parent *CellRendererTextClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var CellRendererComboPrivateStruct *gi.Struct
var CellRendererComboPrivateStructOnce sync.Once

func CellRendererComboPrivateStructSet() {
	CellRendererComboPrivateStructOnce.Do(func() {
		CellRendererComboPrivateStruct = gi.StructNew("Gtk", "CellRendererComboPrivate")
	})
}

type CellRendererComboPrivate struct {
	native uintptr
}

var CellRendererPixbufClassStruct *gi.Struct
var CellRendererPixbufClassStructOnce sync.Once

func CellRendererPixbufClassStructSet() {
	CellRendererPixbufClassStructOnce.Do(func() {
		CellRendererPixbufClassStruct = gi.StructNew("Gtk", "CellRendererPixbufClass")
	})
}

type CellRendererPixbufClass struct {
	native      uintptr
	ParentClass *CellRendererClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var CellRendererPixbufPrivateStruct *gi.Struct
var CellRendererPixbufPrivateStructOnce sync.Once

func CellRendererPixbufPrivateStructSet() {
	CellRendererPixbufPrivateStructOnce.Do(func() {
		CellRendererPixbufPrivateStruct = gi.StructNew("Gtk", "CellRendererPixbufPrivate")
	})
}

type CellRendererPixbufPrivate struct {
	native uintptr
}

var CellRendererPrivateStruct *gi.Struct
var CellRendererPrivateStructOnce sync.Once

func CellRendererPrivateStructSet() {
	CellRendererPrivateStructOnce.Do(func() {
		CellRendererPrivateStruct = gi.StructNew("Gtk", "CellRendererPrivate")
	})
}

type CellRendererPrivate struct {
	native uintptr
}

var CellRendererProgressClassStruct *gi.Struct
var CellRendererProgressClassStructOnce sync.Once

func CellRendererProgressClassStructSet() {
	CellRendererProgressClassStructOnce.Do(func() {
		CellRendererProgressClassStruct = gi.StructNew("Gtk", "CellRendererProgressClass")
	})
}

type CellRendererProgressClass struct {
	native      uintptr
	ParentClass *CellRendererClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var CellRendererProgressPrivateStruct *gi.Struct
var CellRendererProgressPrivateStructOnce sync.Once

func CellRendererProgressPrivateStructSet() {
	CellRendererProgressPrivateStructOnce.Do(func() {
		CellRendererProgressPrivateStruct = gi.StructNew("Gtk", "CellRendererProgressPrivate")
	})
}

type CellRendererProgressPrivate struct {
	native uintptr
}

var CellRendererSpinClassStruct *gi.Struct
var CellRendererSpinClassStructOnce sync.Once

func CellRendererSpinClassStructSet() {
	CellRendererSpinClassStructOnce.Do(func() {
		CellRendererSpinClassStruct = gi.StructNew("Gtk", "CellRendererSpinClass")
	})
}

type CellRendererSpinClass struct {
	native uintptr
	Parent *CellRendererTextClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var CellRendererSpinPrivateStruct *gi.Struct
var CellRendererSpinPrivateStructOnce sync.Once

func CellRendererSpinPrivateStructSet() {
	CellRendererSpinPrivateStructOnce.Do(func() {
		CellRendererSpinPrivateStruct = gi.StructNew("Gtk", "CellRendererSpinPrivate")
	})
}

type CellRendererSpinPrivate struct {
	native uintptr
}

var CellRendererSpinnerClassStruct *gi.Struct
var CellRendererSpinnerClassStructOnce sync.Once

func CellRendererSpinnerClassStructSet() {
	CellRendererSpinnerClassStructOnce.Do(func() {
		CellRendererSpinnerClassStruct = gi.StructNew("Gtk", "CellRendererSpinnerClass")
	})
}

type CellRendererSpinnerClass struct {
	native      uintptr
	ParentClass *CellRendererClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var CellRendererSpinnerPrivateStruct *gi.Struct
var CellRendererSpinnerPrivateStructOnce sync.Once

func CellRendererSpinnerPrivateStructSet() {
	CellRendererSpinnerPrivateStructOnce.Do(func() {
		CellRendererSpinnerPrivateStruct = gi.StructNew("Gtk", "CellRendererSpinnerPrivate")
	})
}

type CellRendererSpinnerPrivate struct {
	native uintptr
}

var CellRendererTextClassStruct *gi.Struct
var CellRendererTextClassStructOnce sync.Once

func CellRendererTextClassStructSet() {
	CellRendererTextClassStructOnce.Do(func() {
		CellRendererTextClassStruct = gi.StructNew("Gtk", "CellRendererTextClass")
	})
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

var CellRendererTextPrivateStruct *gi.Struct
var CellRendererTextPrivateStructOnce sync.Once

func CellRendererTextPrivateStructSet() {
	CellRendererTextPrivateStructOnce.Do(func() {
		CellRendererTextPrivateStruct = gi.StructNew("Gtk", "CellRendererTextPrivate")
	})
}

type CellRendererTextPrivate struct {
	native uintptr
}

var CellRendererToggleClassStruct *gi.Struct
var CellRendererToggleClassStructOnce sync.Once

func CellRendererToggleClassStructSet() {
	CellRendererToggleClassStructOnce.Do(func() {
		CellRendererToggleClassStruct = gi.StructNew("Gtk", "CellRendererToggleClass")
	})
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

var CellRendererTogglePrivateStruct *gi.Struct
var CellRendererTogglePrivateStructOnce sync.Once

func CellRendererTogglePrivateStructSet() {
	CellRendererTogglePrivateStructOnce.Do(func() {
		CellRendererTogglePrivateStruct = gi.StructNew("Gtk", "CellRendererTogglePrivate")
	})
}

type CellRendererTogglePrivate struct {
	native uintptr
}

var CellViewClassStruct *gi.Struct
var CellViewClassStructOnce sync.Once

func CellViewClassStructSet() {
	CellViewClassStructOnce.Do(func() {
		CellViewClassStruct = gi.StructNew("Gtk", "CellViewClass")
	})
}

type CellViewClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var CellViewPrivateStruct *gi.Struct
var CellViewPrivateStructOnce sync.Once

func CellViewPrivateStructSet() {
	CellViewPrivateStructOnce.Do(func() {
		CellViewPrivateStruct = gi.StructNew("Gtk", "CellViewPrivate")
	})
}

type CellViewPrivate struct {
	native uintptr
}

var CheckButtonClassStruct *gi.Struct
var CheckButtonClassStructOnce sync.Once

func CheckButtonClassStructSet() {
	CheckButtonClassStructOnce.Do(func() {
		CheckButtonClassStruct = gi.StructNew("Gtk", "CheckButtonClass")
	})
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

var CheckMenuItemAccessibleClassStruct *gi.Struct
var CheckMenuItemAccessibleClassStructOnce sync.Once

func CheckMenuItemAccessibleClassStructSet() {
	CheckMenuItemAccessibleClassStructOnce.Do(func() {
		CheckMenuItemAccessibleClassStruct = gi.StructNew("Gtk", "CheckMenuItemAccessibleClass")
	})
}

type CheckMenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *MenuItemAccessibleClass
}

var CheckMenuItemAccessiblePrivateStruct *gi.Struct
var CheckMenuItemAccessiblePrivateStructOnce sync.Once

func CheckMenuItemAccessiblePrivateStructSet() {
	CheckMenuItemAccessiblePrivateStructOnce.Do(func() {
		CheckMenuItemAccessiblePrivateStruct = gi.StructNew("Gtk", "CheckMenuItemAccessiblePrivate")
	})
}

type CheckMenuItemAccessiblePrivate struct {
	native uintptr
}

var CheckMenuItemClassStruct *gi.Struct
var CheckMenuItemClassStructOnce sync.Once

func CheckMenuItemClassStructSet() {
	CheckMenuItemClassStructOnce.Do(func() {
		CheckMenuItemClassStruct = gi.StructNew("Gtk", "CheckMenuItemClass")
	})
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

var CheckMenuItemPrivateStruct *gi.Struct
var CheckMenuItemPrivateStructOnce sync.Once

func CheckMenuItemPrivateStructSet() {
	CheckMenuItemPrivateStructOnce.Do(func() {
		CheckMenuItemPrivateStruct = gi.StructNew("Gtk", "CheckMenuItemPrivate")
	})
}

type CheckMenuItemPrivate struct {
	native uintptr
}

var ColorButtonClassStruct *gi.Struct
var ColorButtonClassStructOnce sync.Once

func ColorButtonClassStructSet() {
	ColorButtonClassStructOnce.Do(func() {
		ColorButtonClassStruct = gi.StructNew("Gtk", "ColorButtonClass")
	})
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

var ColorButtonPrivateStruct *gi.Struct
var ColorButtonPrivateStructOnce sync.Once

func ColorButtonPrivateStructSet() {
	ColorButtonPrivateStructOnce.Do(func() {
		ColorButtonPrivateStruct = gi.StructNew("Gtk", "ColorButtonPrivate")
	})
}

type ColorButtonPrivate struct {
	native uintptr
}

var ColorChooserDialogClassStruct *gi.Struct
var ColorChooserDialogClassStructOnce sync.Once

func ColorChooserDialogClassStructSet() {
	ColorChooserDialogClassStructOnce.Do(func() {
		ColorChooserDialogClassStruct = gi.StructNew("Gtk", "ColorChooserDialogClass")
	})
}

type ColorChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ColorChooserDialogPrivateStruct *gi.Struct
var ColorChooserDialogPrivateStructOnce sync.Once

func ColorChooserDialogPrivateStructSet() {
	ColorChooserDialogPrivateStructOnce.Do(func() {
		ColorChooserDialogPrivateStruct = gi.StructNew("Gtk", "ColorChooserDialogPrivate")
	})
}

type ColorChooserDialogPrivate struct {
	native uintptr
}

var ColorChooserInterfaceStruct *gi.Struct
var ColorChooserInterfaceStructOnce sync.Once

func ColorChooserInterfaceStructSet() {
	ColorChooserInterfaceStructOnce.Do(func() {
		ColorChooserInterfaceStruct = gi.StructNew("Gtk", "ColorChooserInterface")
	})
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

var ColorChooserWidgetClassStruct *gi.Struct
var ColorChooserWidgetClassStructOnce sync.Once

func ColorChooserWidgetClassStructSet() {
	ColorChooserWidgetClassStructOnce.Do(func() {
		ColorChooserWidgetClassStruct = gi.StructNew("Gtk", "ColorChooserWidgetClass")
	})
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

var ColorChooserWidgetPrivateStruct *gi.Struct
var ColorChooserWidgetPrivateStructOnce sync.Once

func ColorChooserWidgetPrivateStructSet() {
	ColorChooserWidgetPrivateStructOnce.Do(func() {
		ColorChooserWidgetPrivateStruct = gi.StructNew("Gtk", "ColorChooserWidgetPrivate")
	})
}

type ColorChooserWidgetPrivate struct {
	native uintptr
}

var ColorSelectionClassStruct *gi.Struct
var ColorSelectionClassStructOnce sync.Once

func ColorSelectionClassStructSet() {
	ColorSelectionClassStructOnce.Do(func() {
		ColorSelectionClassStruct = gi.StructNew("Gtk", "ColorSelectionClass")
	})
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

var ColorSelectionDialogClassStruct *gi.Struct
var ColorSelectionDialogClassStructOnce sync.Once

func ColorSelectionDialogClassStructSet() {
	ColorSelectionDialogClassStructOnce.Do(func() {
		ColorSelectionDialogClassStruct = gi.StructNew("Gtk", "ColorSelectionDialogClass")
	})
}

type ColorSelectionDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ColorSelectionDialogPrivateStruct *gi.Struct
var ColorSelectionDialogPrivateStructOnce sync.Once

func ColorSelectionDialogPrivateStructSet() {
	ColorSelectionDialogPrivateStructOnce.Do(func() {
		ColorSelectionDialogPrivateStruct = gi.StructNew("Gtk", "ColorSelectionDialogPrivate")
	})
}

type ColorSelectionDialogPrivate struct {
	native uintptr
}

var ColorSelectionPrivateStruct *gi.Struct
var ColorSelectionPrivateStructOnce sync.Once

func ColorSelectionPrivateStructSet() {
	ColorSelectionPrivateStructOnce.Do(func() {
		ColorSelectionPrivateStruct = gi.StructNew("Gtk", "ColorSelectionPrivate")
	})
}

type ColorSelectionPrivate struct {
	native uintptr
}

var ComboBoxAccessibleClassStruct *gi.Struct
var ComboBoxAccessibleClassStructOnce sync.Once

func ComboBoxAccessibleClassStructSet() {
	ComboBoxAccessibleClassStructOnce.Do(func() {
		ComboBoxAccessibleClassStruct = gi.StructNew("Gtk", "ComboBoxAccessibleClass")
	})
}

type ComboBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var ComboBoxAccessiblePrivateStruct *gi.Struct
var ComboBoxAccessiblePrivateStructOnce sync.Once

func ComboBoxAccessiblePrivateStructSet() {
	ComboBoxAccessiblePrivateStructOnce.Do(func() {
		ComboBoxAccessiblePrivateStruct = gi.StructNew("Gtk", "ComboBoxAccessiblePrivate")
	})
}

type ComboBoxAccessiblePrivate struct {
	native uintptr
}

var ComboBoxClassStruct *gi.Struct
var ComboBoxClassStructOnce sync.Once

func ComboBoxClassStructSet() {
	ComboBoxClassStructOnce.Do(func() {
		ComboBoxClassStruct = gi.StructNew("Gtk", "ComboBoxClass")
	})
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

var ComboBoxPrivateStruct *gi.Struct
var ComboBoxPrivateStructOnce sync.Once

func ComboBoxPrivateStructSet() {
	ComboBoxPrivateStructOnce.Do(func() {
		ComboBoxPrivateStruct = gi.StructNew("Gtk", "ComboBoxPrivate")
	})
}

type ComboBoxPrivate struct {
	native uintptr
}

var ComboBoxTextClassStruct *gi.Struct
var ComboBoxTextClassStructOnce sync.Once

func ComboBoxTextClassStructSet() {
	ComboBoxTextClassStructOnce.Do(func() {
		ComboBoxTextClassStruct = gi.StructNew("Gtk", "ComboBoxTextClass")
	})
}

type ComboBoxTextClass struct {
	native      uintptr
	ParentClass *ComboBoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ComboBoxTextPrivateStruct *gi.Struct
var ComboBoxTextPrivateStructOnce sync.Once

func ComboBoxTextPrivateStructSet() {
	ComboBoxTextPrivateStructOnce.Do(func() {
		ComboBoxTextPrivateStruct = gi.StructNew("Gtk", "ComboBoxTextPrivate")
	})
}

type ComboBoxTextPrivate struct {
	native uintptr
}

var ContainerAccessibleClassStruct *gi.Struct
var ContainerAccessibleClassStructOnce sync.Once

func ContainerAccessibleClassStructSet() {
	ContainerAccessibleClassStructOnce.Do(func() {
		ContainerAccessibleClassStruct = gi.StructNew("Gtk", "ContainerAccessibleClass")
	})
}

type ContainerAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
	// UNSUPPORTED : C value 'add_gtk' : missing Type
	// UNSUPPORTED : C value 'remove_gtk' : missing Type
}

var ContainerAccessiblePrivateStruct *gi.Struct
var ContainerAccessiblePrivateStructOnce sync.Once

func ContainerAccessiblePrivateStructSet() {
	ContainerAccessiblePrivateStructOnce.Do(func() {
		ContainerAccessiblePrivateStruct = gi.StructNew("Gtk", "ContainerAccessiblePrivate")
	})
}

type ContainerAccessiblePrivate struct {
	native uintptr
}

var ContainerCellAccessibleClassStruct *gi.Struct
var ContainerCellAccessibleClassStructOnce sync.Once

func ContainerCellAccessibleClassStructSet() {
	ContainerCellAccessibleClassStructOnce.Do(func() {
		ContainerCellAccessibleClassStruct = gi.StructNew("Gtk", "ContainerCellAccessibleClass")
	})
}

type ContainerCellAccessibleClass struct {
	native      uintptr
	ParentClass *CellAccessibleClass
}

var ContainerCellAccessiblePrivateStruct *gi.Struct
var ContainerCellAccessiblePrivateStructOnce sync.Once

func ContainerCellAccessiblePrivateStructSet() {
	ContainerCellAccessiblePrivateStructOnce.Do(func() {
		ContainerCellAccessiblePrivateStruct = gi.StructNew("Gtk", "ContainerCellAccessiblePrivate")
	})
}

type ContainerCellAccessiblePrivate struct {
	native uintptr
}

var ContainerClassStruct *gi.Struct
var ContainerClassStructOnce sync.Once

func ContainerClassStructSet() {
	ContainerClassStructOnce.Do(func() {
		ContainerClassStruct = gi.StructNew("Gtk", "ContainerClass")
	})
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

// UNSUPPORTED : C value 'gtk_container_class_find_child_property' : return type 'GObject.ParamSpec' not supported

var handleBorderWidthContainerClassInvoker *gi.Function

// HandleBorderWidth is a representation of the C type gtk_container_class_handle_border_width.
func (recv *ContainerClass) HandleBorderWidth() {
	if handleBorderWidthContainerClassInvoker == nil {
		handleBorderWidthContainerClassInvoker = gi.StructFunctionInvokerNew("Gtk", "ContainerClass", "handle_border_width")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	handleBorderWidthContainerClassInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_container_class_install_child_properties' : parameter 'pspecs' has no type

// UNSUPPORTED : C value 'gtk_container_class_install_child_property' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

var listChildPropertiesContainerClassInvoker *gi.Function

// ListChildProperties is a representation of the C type gtk_container_class_list_child_properties.
func (recv *ContainerClass) ListChildProperties() uint32 {
	if listChildPropertiesContainerClassInvoker == nil {
		listChildPropertiesContainerClassInvoker = gi.StructFunctionInvokerNew("Gtk", "ContainerClass", "list_child_properties")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	listChildPropertiesContainerClassInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint32()

	return out0
}

var ContainerPrivateStruct *gi.Struct
var ContainerPrivateStructOnce sync.Once

func ContainerPrivateStructSet() {
	ContainerPrivateStructOnce.Do(func() {
		ContainerPrivateStruct = gi.StructNew("Gtk", "ContainerPrivate")
	})
}

type ContainerPrivate struct {
	native uintptr
}

var CssProviderClassStruct *gi.Struct
var CssProviderClassStructOnce sync.Once

func CssProviderClassStructSet() {
	CssProviderClassStructOnce.Do(func() {
		CssProviderClassStruct = gi.StructNew("Gtk", "CssProviderClass")
	})
}

type CssProviderClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'parsing_error' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var CssProviderPrivateStruct *gi.Struct
var CssProviderPrivateStructOnce sync.Once

func CssProviderPrivateStructSet() {
	CssProviderPrivateStructOnce.Do(func() {
		CssProviderPrivateStruct = gi.StructNew("Gtk", "CssProviderPrivate")
	})
}

type CssProviderPrivate struct {
	native uintptr
}

var CssSectionStruct *gi.Struct
var CssSectionStructOnce sync.Once

func CssSectionStructSet() {
	CssSectionStructOnce.Do(func() {
		CssSectionStruct = gi.StructNew("Gtk", "CssSection")
	})
}

type CssSection struct {
	native uintptr
}

var getEndLineCssSectionInvoker *gi.Function

// GetEndLine is a representation of the C type gtk_css_section_get_end_line.
func (recv *CssSection) GetEndLine() uint32 {
	if getEndLineCssSectionInvoker == nil {
		getEndLineCssSectionInvoker = gi.StructFunctionInvokerNew("Gtk", "CssSection", "get_end_line")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getEndLineCssSectionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var getEndPositionCssSectionInvoker *gi.Function

// GetEndPosition is a representation of the C type gtk_css_section_get_end_position.
func (recv *CssSection) GetEndPosition() uint32 {
	if getEndPositionCssSectionInvoker == nil {
		getEndPositionCssSectionInvoker = gi.StructFunctionInvokerNew("Gtk", "CssSection", "get_end_position")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getEndPositionCssSectionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_css_section_get_file' : return type 'Gio.File' not supported

var getParentCssSectionInvoker *gi.Function

// GetParent is a representation of the C type gtk_css_section_get_parent.
func (recv *CssSection) GetParent() *CssSection {
	if getParentCssSectionInvoker == nil {
		getParentCssSectionInvoker = gi.StructFunctionInvokerNew("Gtk", "CssSection", "get_parent")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getParentCssSectionInvoker.Invoke(inArgs[:], nil)

	retGo := &CssSection{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_css_section_get_section_type' : return type 'CssSectionType' not supported

var getStartLineCssSectionInvoker *gi.Function

// GetStartLine is a representation of the C type gtk_css_section_get_start_line.
func (recv *CssSection) GetStartLine() uint32 {
	if getStartLineCssSectionInvoker == nil {
		getStartLineCssSectionInvoker = gi.StructFunctionInvokerNew("Gtk", "CssSection", "get_start_line")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStartLineCssSectionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var getStartPositionCssSectionInvoker *gi.Function

// GetStartPosition is a representation of the C type gtk_css_section_get_start_position.
func (recv *CssSection) GetStartPosition() uint32 {
	if getStartPositionCssSectionInvoker == nil {
		getStartPositionCssSectionInvoker = gi.StructFunctionInvokerNew("Gtk", "CssSection", "get_start_position")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getStartPositionCssSectionInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var refCssSectionInvoker *gi.Function

// Ref is a representation of the C type gtk_css_section_ref.
func (recv *CssSection) Ref() *CssSection {
	if refCssSectionInvoker == nil {
		refCssSectionInvoker = gi.StructFunctionInvokerNew("Gtk", "CssSection", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refCssSectionInvoker.Invoke(inArgs[:], nil)

	retGo := &CssSection{native: ret.Pointer()}

	return retGo
}

var unrefCssSectionInvoker *gi.Function

// Unref is a representation of the C type gtk_css_section_unref.
func (recv *CssSection) Unref() {
	if unrefCssSectionInvoker == nil {
		unrefCssSectionInvoker = gi.StructFunctionInvokerNew("Gtk", "CssSection", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefCssSectionInvoker.Invoke(inArgs[:], nil)

}

var DialogClassStruct *gi.Struct
var DialogClassStructOnce sync.Once

func DialogClassStructSet() {
	DialogClassStructOnce.Do(func() {
		DialogClassStruct = gi.StructNew("Gtk", "DialogClass")
	})
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

var DialogPrivateStruct *gi.Struct
var DialogPrivateStructOnce sync.Once

func DialogPrivateStructSet() {
	DialogPrivateStructOnce.Do(func() {
		DialogPrivateStruct = gi.StructNew("Gtk", "DialogPrivate")
	})
}

type DialogPrivate struct {
	native uintptr
}

var DrawingAreaClassStruct *gi.Struct
var DrawingAreaClassStructOnce sync.Once

func DrawingAreaClassStructSet() {
	DrawingAreaClassStructOnce.Do(func() {
		DrawingAreaClassStruct = gi.StructNew("Gtk", "DrawingAreaClass")
	})
}

type DrawingAreaClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var EditableInterfaceStruct *gi.Struct
var EditableInterfaceStructOnce sync.Once

func EditableInterfaceStructSet() {
	EditableInterfaceStructOnce.Do(func() {
		EditableInterfaceStruct = gi.StructNew("Gtk", "EditableInterface")
	})
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

var EntryAccessibleClassStruct *gi.Struct
var EntryAccessibleClassStructOnce sync.Once

func EntryAccessibleClassStructSet() {
	EntryAccessibleClassStructOnce.Do(func() {
		EntryAccessibleClassStruct = gi.StructNew("Gtk", "EntryAccessibleClass")
	})
}

type EntryAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var EntryAccessiblePrivateStruct *gi.Struct
var EntryAccessiblePrivateStructOnce sync.Once

func EntryAccessiblePrivateStructSet() {
	EntryAccessiblePrivateStructOnce.Do(func() {
		EntryAccessiblePrivateStruct = gi.StructNew("Gtk", "EntryAccessiblePrivate")
	})
}

type EntryAccessiblePrivate struct {
	native uintptr
}

var EntryBufferClassStruct *gi.Struct
var EntryBufferClassStructOnce sync.Once

func EntryBufferClassStructSet() {
	EntryBufferClassStructOnce.Do(func() {
		EntryBufferClassStruct = gi.StructNew("Gtk", "EntryBufferClass")
	})
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

var EntryBufferPrivateStruct *gi.Struct
var EntryBufferPrivateStructOnce sync.Once

func EntryBufferPrivateStructSet() {
	EntryBufferPrivateStructOnce.Do(func() {
		EntryBufferPrivateStruct = gi.StructNew("Gtk", "EntryBufferPrivate")
	})
}

type EntryBufferPrivate struct {
	native uintptr
}

var EntryClassStruct *gi.Struct
var EntryClassStructOnce sync.Once

func EntryClassStructSet() {
	EntryClassStructOnce.Do(func() {
		EntryClassStruct = gi.StructNew("Gtk", "EntryClass")
	})
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

var EntryCompletionClassStruct *gi.Struct
var EntryCompletionClassStructOnce sync.Once

func EntryCompletionClassStructSet() {
	EntryCompletionClassStructOnce.Do(func() {
		EntryCompletionClassStruct = gi.StructNew("Gtk", "EntryCompletionClass")
	})
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

var EntryCompletionPrivateStruct *gi.Struct
var EntryCompletionPrivateStructOnce sync.Once

func EntryCompletionPrivateStructSet() {
	EntryCompletionPrivateStructOnce.Do(func() {
		EntryCompletionPrivateStruct = gi.StructNew("Gtk", "EntryCompletionPrivate")
	})
}

type EntryCompletionPrivate struct {
	native uintptr
}

var EntryPrivateStruct *gi.Struct
var EntryPrivateStructOnce sync.Once

func EntryPrivateStructSet() {
	EntryPrivateStructOnce.Do(func() {
		EntryPrivateStruct = gi.StructNew("Gtk", "EntryPrivate")
	})
}

type EntryPrivate struct {
	native uintptr
}

var EventBoxClassStruct *gi.Struct
var EventBoxClassStructOnce sync.Once

func EventBoxClassStructSet() {
	EventBoxClassStructOnce.Do(func() {
		EventBoxClassStruct = gi.StructNew("Gtk", "EventBoxClass")
	})
}

type EventBoxClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var EventBoxPrivateStruct *gi.Struct
var EventBoxPrivateStructOnce sync.Once

func EventBoxPrivateStructSet() {
	EventBoxPrivateStructOnce.Do(func() {
		EventBoxPrivateStruct = gi.StructNew("Gtk", "EventBoxPrivate")
	})
}

type EventBoxPrivate struct {
	native uintptr
}

var EventControllerClassStruct *gi.Struct
var EventControllerClassStructOnce sync.Once

func EventControllerClassStructSet() {
	EventControllerClassStructOnce.Do(func() {
		EventControllerClassStruct = gi.StructNew("Gtk", "EventControllerClass")
	})
}

type EventControllerClass struct {
	native uintptr
}

var EventControllerKeyClassStruct *gi.Struct
var EventControllerKeyClassStructOnce sync.Once

func EventControllerKeyClassStructSet() {
	EventControllerKeyClassStructOnce.Do(func() {
		EventControllerKeyClassStruct = gi.StructNew("Gtk", "EventControllerKeyClass")
	})
}

type EventControllerKeyClass struct {
	native uintptr
}

var EventControllerMotionClassStruct *gi.Struct
var EventControllerMotionClassStructOnce sync.Once

func EventControllerMotionClassStructSet() {
	EventControllerMotionClassStructOnce.Do(func() {
		EventControllerMotionClassStruct = gi.StructNew("Gtk", "EventControllerMotionClass")
	})
}

type EventControllerMotionClass struct {
	native uintptr
}

var EventControllerScrollClassStruct *gi.Struct
var EventControllerScrollClassStructOnce sync.Once

func EventControllerScrollClassStructSet() {
	EventControllerScrollClassStructOnce.Do(func() {
		EventControllerScrollClassStruct = gi.StructNew("Gtk", "EventControllerScrollClass")
	})
}

type EventControllerScrollClass struct {
	native uintptr
}

var ExpanderAccessibleClassStruct *gi.Struct
var ExpanderAccessibleClassStructOnce sync.Once

func ExpanderAccessibleClassStructSet() {
	ExpanderAccessibleClassStructOnce.Do(func() {
		ExpanderAccessibleClassStruct = gi.StructNew("Gtk", "ExpanderAccessibleClass")
	})
}

type ExpanderAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var ExpanderAccessiblePrivateStruct *gi.Struct
var ExpanderAccessiblePrivateStructOnce sync.Once

func ExpanderAccessiblePrivateStructSet() {
	ExpanderAccessiblePrivateStructOnce.Do(func() {
		ExpanderAccessiblePrivateStruct = gi.StructNew("Gtk", "ExpanderAccessiblePrivate")
	})
}

type ExpanderAccessiblePrivate struct {
	native uintptr
}

var ExpanderClassStruct *gi.Struct
var ExpanderClassStructOnce sync.Once

func ExpanderClassStructSet() {
	ExpanderClassStructOnce.Do(func() {
		ExpanderClassStruct = gi.StructNew("Gtk", "ExpanderClass")
	})
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

var ExpanderPrivateStruct *gi.Struct
var ExpanderPrivateStructOnce sync.Once

func ExpanderPrivateStructSet() {
	ExpanderPrivateStructOnce.Do(func() {
		ExpanderPrivateStruct = gi.StructNew("Gtk", "ExpanderPrivate")
	})
}

type ExpanderPrivate struct {
	native uintptr
}

var FileChooserButtonClassStruct *gi.Struct
var FileChooserButtonClassStructOnce sync.Once

func FileChooserButtonClassStructSet() {
	FileChooserButtonClassStructOnce.Do(func() {
		FileChooserButtonClassStruct = gi.StructNew("Gtk", "FileChooserButtonClass")
	})
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

var FileChooserButtonPrivateStruct *gi.Struct
var FileChooserButtonPrivateStructOnce sync.Once

func FileChooserButtonPrivateStructSet() {
	FileChooserButtonPrivateStructOnce.Do(func() {
		FileChooserButtonPrivateStruct = gi.StructNew("Gtk", "FileChooserButtonPrivate")
	})
}

type FileChooserButtonPrivate struct {
	native uintptr
}

var FileChooserDialogClassStruct *gi.Struct
var FileChooserDialogClassStructOnce sync.Once

func FileChooserDialogClassStructSet() {
	FileChooserDialogClassStructOnce.Do(func() {
		FileChooserDialogClassStruct = gi.StructNew("Gtk", "FileChooserDialogClass")
	})
}

type FileChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var FileChooserDialogPrivateStruct *gi.Struct
var FileChooserDialogPrivateStructOnce sync.Once

func FileChooserDialogPrivateStructSet() {
	FileChooserDialogPrivateStructOnce.Do(func() {
		FileChooserDialogPrivateStruct = gi.StructNew("Gtk", "FileChooserDialogPrivate")
	})
}

type FileChooserDialogPrivate struct {
	native uintptr
}

var FileChooserNativeClassStruct *gi.Struct
var FileChooserNativeClassStructOnce sync.Once

func FileChooserNativeClassStructSet() {
	FileChooserNativeClassStructOnce.Do(func() {
		FileChooserNativeClassStruct = gi.StructNew("Gtk", "FileChooserNativeClass")
	})
}

type FileChooserNativeClass struct {
	native      uintptr
	ParentClass *NativeDialogClass
}

var FileChooserWidgetClassStruct *gi.Struct
var FileChooserWidgetClassStructOnce sync.Once

func FileChooserWidgetClassStructSet() {
	FileChooserWidgetClassStructOnce.Do(func() {
		FileChooserWidgetClassStruct = gi.StructNew("Gtk", "FileChooserWidgetClass")
	})
}

type FileChooserWidgetClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var FileChooserWidgetPrivateStruct *gi.Struct
var FileChooserWidgetPrivateStructOnce sync.Once

func FileChooserWidgetPrivateStructSet() {
	FileChooserWidgetPrivateStructOnce.Do(func() {
		FileChooserWidgetPrivateStruct = gi.StructNew("Gtk", "FileChooserWidgetPrivate")
	})
}

type FileChooserWidgetPrivate struct {
	native uintptr
}

var FileFilterInfoStruct *gi.Struct
var FileFilterInfoStructOnce sync.Once

func FileFilterInfoStructSet() {
	FileFilterInfoStructOnce.Do(func() {
		FileFilterInfoStruct = gi.StructNew("Gtk", "FileFilterInfo")
	})
}

type FileFilterInfo struct {
	native uintptr
	// UNSUPPORTED : C value 'contains' : no Go type for 'FileFilterFlags'
	Filename    string
	Uri         string
	DisplayName string
	MimeType    string
}

var FixedChildStruct *gi.Struct
var FixedChildStructOnce sync.Once

func FixedChildStructSet() {
	FixedChildStructOnce.Do(func() {
		FixedChildStruct = gi.StructNew("Gtk", "FixedChild")
	})
}

type FixedChild struct {
	native uintptr
	// UNSUPPORTED : C value 'widget' : no Go type for 'Widget'
	X int32
	Y int32
}

var FixedClassStruct *gi.Struct
var FixedClassStructOnce sync.Once

func FixedClassStructSet() {
	FixedClassStructOnce.Do(func() {
		FixedClassStruct = gi.StructNew("Gtk", "FixedClass")
	})
}

type FixedClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var FixedPrivateStruct *gi.Struct
var FixedPrivateStructOnce sync.Once

func FixedPrivateStructSet() {
	FixedPrivateStructOnce.Do(func() {
		FixedPrivateStruct = gi.StructNew("Gtk", "FixedPrivate")
	})
}

type FixedPrivate struct {
	native uintptr
}

var FlowBoxAccessibleClassStruct *gi.Struct
var FlowBoxAccessibleClassStructOnce sync.Once

func FlowBoxAccessibleClassStructSet() {
	FlowBoxAccessibleClassStructOnce.Do(func() {
		FlowBoxAccessibleClassStruct = gi.StructNew("Gtk", "FlowBoxAccessibleClass")
	})
}

type FlowBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var FlowBoxAccessiblePrivateStruct *gi.Struct
var FlowBoxAccessiblePrivateStructOnce sync.Once

func FlowBoxAccessiblePrivateStructSet() {
	FlowBoxAccessiblePrivateStructOnce.Do(func() {
		FlowBoxAccessiblePrivateStruct = gi.StructNew("Gtk", "FlowBoxAccessiblePrivate")
	})
}

type FlowBoxAccessiblePrivate struct {
	native uintptr
}

var FlowBoxChildAccessibleClassStruct *gi.Struct
var FlowBoxChildAccessibleClassStructOnce sync.Once

func FlowBoxChildAccessibleClassStructSet() {
	FlowBoxChildAccessibleClassStructOnce.Do(func() {
		FlowBoxChildAccessibleClassStruct = gi.StructNew("Gtk", "FlowBoxChildAccessibleClass")
	})
}

type FlowBoxChildAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var FlowBoxChildClassStruct *gi.Struct
var FlowBoxChildClassStructOnce sync.Once

func FlowBoxChildClassStructSet() {
	FlowBoxChildClassStructOnce.Do(func() {
		FlowBoxChildClassStruct = gi.StructNew("Gtk", "FlowBoxChildClass")
	})
}

type FlowBoxChildClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'activate' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
}

var FlowBoxClassStruct *gi.Struct
var FlowBoxClassStructOnce sync.Once

func FlowBoxClassStructSet() {
	FlowBoxClassStructOnce.Do(func() {
		FlowBoxClassStruct = gi.StructNew("Gtk", "FlowBoxClass")
	})
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

var FontButtonClassStruct *gi.Struct
var FontButtonClassStructOnce sync.Once

func FontButtonClassStructSet() {
	FontButtonClassStructOnce.Do(func() {
		FontButtonClassStruct = gi.StructNew("Gtk", "FontButtonClass")
	})
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

var FontButtonPrivateStruct *gi.Struct
var FontButtonPrivateStructOnce sync.Once

func FontButtonPrivateStructSet() {
	FontButtonPrivateStructOnce.Do(func() {
		FontButtonPrivateStruct = gi.StructNew("Gtk", "FontButtonPrivate")
	})
}

type FontButtonPrivate struct {
	native uintptr
}

var FontChooserDialogClassStruct *gi.Struct
var FontChooserDialogClassStructOnce sync.Once

func FontChooserDialogClassStructSet() {
	FontChooserDialogClassStructOnce.Do(func() {
		FontChooserDialogClassStruct = gi.StructNew("Gtk", "FontChooserDialogClass")
	})
}

type FontChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var FontChooserDialogPrivateStruct *gi.Struct
var FontChooserDialogPrivateStructOnce sync.Once

func FontChooserDialogPrivateStructSet() {
	FontChooserDialogPrivateStructOnce.Do(func() {
		FontChooserDialogPrivateStruct = gi.StructNew("Gtk", "FontChooserDialogPrivate")
	})
}

type FontChooserDialogPrivate struct {
	native uintptr
}

var FontChooserIfaceStruct *gi.Struct
var FontChooserIfaceStructOnce sync.Once

func FontChooserIfaceStructSet() {
	FontChooserIfaceStructOnce.Do(func() {
		FontChooserIfaceStruct = gi.StructNew("Gtk", "FontChooserIface")
	})
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

var FontChooserWidgetClassStruct *gi.Struct
var FontChooserWidgetClassStructOnce sync.Once

func FontChooserWidgetClassStructSet() {
	FontChooserWidgetClassStructOnce.Do(func() {
		FontChooserWidgetClassStruct = gi.StructNew("Gtk", "FontChooserWidgetClass")
	})
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

var FontChooserWidgetPrivateStruct *gi.Struct
var FontChooserWidgetPrivateStructOnce sync.Once

func FontChooserWidgetPrivateStructSet() {
	FontChooserWidgetPrivateStructOnce.Do(func() {
		FontChooserWidgetPrivateStruct = gi.StructNew("Gtk", "FontChooserWidgetPrivate")
	})
}

type FontChooserWidgetPrivate struct {
	native uintptr
}

var FontSelectionClassStruct *gi.Struct
var FontSelectionClassStructOnce sync.Once

func FontSelectionClassStructSet() {
	FontSelectionClassStructOnce.Do(func() {
		FontSelectionClassStruct = gi.StructNew("Gtk", "FontSelectionClass")
	})
}

type FontSelectionClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var FontSelectionDialogClassStruct *gi.Struct
var FontSelectionDialogClassStructOnce sync.Once

func FontSelectionDialogClassStructSet() {
	FontSelectionDialogClassStructOnce.Do(func() {
		FontSelectionDialogClassStruct = gi.StructNew("Gtk", "FontSelectionDialogClass")
	})
}

type FontSelectionDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var FontSelectionDialogPrivateStruct *gi.Struct
var FontSelectionDialogPrivateStructOnce sync.Once

func FontSelectionDialogPrivateStructSet() {
	FontSelectionDialogPrivateStructOnce.Do(func() {
		FontSelectionDialogPrivateStruct = gi.StructNew("Gtk", "FontSelectionDialogPrivate")
	})
}

type FontSelectionDialogPrivate struct {
	native uintptr
}

var FontSelectionPrivateStruct *gi.Struct
var FontSelectionPrivateStructOnce sync.Once

func FontSelectionPrivateStructSet() {
	FontSelectionPrivateStructOnce.Do(func() {
		FontSelectionPrivateStruct = gi.StructNew("Gtk", "FontSelectionPrivate")
	})
}

type FontSelectionPrivate struct {
	native uintptr
}

var FrameAccessibleClassStruct *gi.Struct
var FrameAccessibleClassStructOnce sync.Once

func FrameAccessibleClassStructSet() {
	FrameAccessibleClassStructOnce.Do(func() {
		FrameAccessibleClassStruct = gi.StructNew("Gtk", "FrameAccessibleClass")
	})
}

type FrameAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var FrameAccessiblePrivateStruct *gi.Struct
var FrameAccessiblePrivateStructOnce sync.Once

func FrameAccessiblePrivateStructSet() {
	FrameAccessiblePrivateStructOnce.Do(func() {
		FrameAccessiblePrivateStruct = gi.StructNew("Gtk", "FrameAccessiblePrivate")
	})
}

type FrameAccessiblePrivate struct {
	native uintptr
}

var FrameClassStruct *gi.Struct
var FrameClassStructOnce sync.Once

func FrameClassStructSet() {
	FrameClassStructOnce.Do(func() {
		FrameClassStruct = gi.StructNew("Gtk", "FrameClass")
	})
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

var FramePrivateStruct *gi.Struct
var FramePrivateStructOnce sync.Once

func FramePrivateStructSet() {
	FramePrivateStructOnce.Do(func() {
		FramePrivateStruct = gi.StructNew("Gtk", "FramePrivate")
	})
}

type FramePrivate struct {
	native uintptr
}

var GLAreaClassStruct *gi.Struct
var GLAreaClassStructOnce sync.Once

func GLAreaClassStructSet() {
	GLAreaClassStructOnce.Do(func() {
		GLAreaClassStruct = gi.StructNew("Gtk", "GLAreaClass")
	})
}

type GLAreaClass struct {
	native uintptr
	// UNSUPPORTED : C value 'render' : missing Type
	// UNSUPPORTED : C value 'resize' : missing Type
	// UNSUPPORTED : C value 'create_context' : missing Type
}

var GestureClassStruct *gi.Struct
var GestureClassStructOnce sync.Once

func GestureClassStructSet() {
	GestureClassStructOnce.Do(func() {
		GestureClassStruct = gi.StructNew("Gtk", "GestureClass")
	})
}

type GestureClass struct {
	native uintptr
}

var GestureDragClassStruct *gi.Struct
var GestureDragClassStructOnce sync.Once

func GestureDragClassStructSet() {
	GestureDragClassStructOnce.Do(func() {
		GestureDragClassStruct = gi.StructNew("Gtk", "GestureDragClass")
	})
}

type GestureDragClass struct {
	native uintptr
}

var GestureLongPressClassStruct *gi.Struct
var GestureLongPressClassStructOnce sync.Once

func GestureLongPressClassStructSet() {
	GestureLongPressClassStructOnce.Do(func() {
		GestureLongPressClassStruct = gi.StructNew("Gtk", "GestureLongPressClass")
	})
}

type GestureLongPressClass struct {
	native uintptr
}

var GestureMultiPressClassStruct *gi.Struct
var GestureMultiPressClassStructOnce sync.Once

func GestureMultiPressClassStructSet() {
	GestureMultiPressClassStructOnce.Do(func() {
		GestureMultiPressClassStruct = gi.StructNew("Gtk", "GestureMultiPressClass")
	})
}

type GestureMultiPressClass struct {
	native uintptr
}

var GesturePanClassStruct *gi.Struct
var GesturePanClassStructOnce sync.Once

func GesturePanClassStructSet() {
	GesturePanClassStructOnce.Do(func() {
		GesturePanClassStruct = gi.StructNew("Gtk", "GesturePanClass")
	})
}

type GesturePanClass struct {
	native uintptr
}

var GestureRotateClassStruct *gi.Struct
var GestureRotateClassStructOnce sync.Once

func GestureRotateClassStructSet() {
	GestureRotateClassStructOnce.Do(func() {
		GestureRotateClassStruct = gi.StructNew("Gtk", "GestureRotateClass")
	})
}

type GestureRotateClass struct {
	native uintptr
}

var GestureSingleClassStruct *gi.Struct
var GestureSingleClassStructOnce sync.Once

func GestureSingleClassStructSet() {
	GestureSingleClassStructOnce.Do(func() {
		GestureSingleClassStruct = gi.StructNew("Gtk", "GestureSingleClass")
	})
}

type GestureSingleClass struct {
	native uintptr
}

var GestureStylusClassStruct *gi.Struct
var GestureStylusClassStructOnce sync.Once

func GestureStylusClassStructSet() {
	GestureStylusClassStructOnce.Do(func() {
		GestureStylusClassStruct = gi.StructNew("Gtk", "GestureStylusClass")
	})
}

type GestureStylusClass struct {
	native uintptr
}

var GestureSwipeClassStruct *gi.Struct
var GestureSwipeClassStructOnce sync.Once

func GestureSwipeClassStructSet() {
	GestureSwipeClassStructOnce.Do(func() {
		GestureSwipeClassStruct = gi.StructNew("Gtk", "GestureSwipeClass")
	})
}

type GestureSwipeClass struct {
	native uintptr
}

var GestureZoomClassStruct *gi.Struct
var GestureZoomClassStructOnce sync.Once

func GestureZoomClassStructSet() {
	GestureZoomClassStructOnce.Do(func() {
		GestureZoomClassStruct = gi.StructNew("Gtk", "GestureZoomClass")
	})
}

type GestureZoomClass struct {
	native uintptr
}

var GradientStruct *gi.Struct
var GradientStructOnce sync.Once

func GradientStructSet() {
	GradientStructOnce.Do(func() {
		GradientStruct = gi.StructNew("Gtk", "Gradient")
	})
}

type Gradient struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_gradient_new_linear' : parameter 'x0' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_gradient_new_radial' : parameter 'x0' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_gradient_add_color_stop' : parameter 'offset' of type 'gdouble' not supported

var refGradientInvoker *gi.Function

// Ref is a representation of the C type gtk_gradient_ref.
func (recv *Gradient) Ref() *Gradient {
	if refGradientInvoker == nil {
		refGradientInvoker = gi.StructFunctionInvokerNew("Gtk", "Gradient", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refGradientInvoker.Invoke(inArgs[:], nil)

	retGo := &Gradient{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_gradient_resolve' : parameter 'props' of type 'StyleProperties' not supported

// UNSUPPORTED : C value 'gtk_gradient_resolve_for_context' : parameter 'context' of type 'StyleContext' not supported

var toStringGradientInvoker *gi.Function

// ToString is a representation of the C type gtk_gradient_to_string.
func (recv *Gradient) ToString() string {
	if toStringGradientInvoker == nil {
		toStringGradientInvoker = gi.StructFunctionInvokerNew("Gtk", "Gradient", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringGradientInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var unrefGradientInvoker *gi.Function

// Unref is a representation of the C type gtk_gradient_unref.
func (recv *Gradient) Unref() {
	if unrefGradientInvoker == nil {
		unrefGradientInvoker = gi.StructFunctionInvokerNew("Gtk", "Gradient", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefGradientInvoker.Invoke(inArgs[:], nil)

}

var GridClassStruct *gi.Struct
var GridClassStructOnce sync.Once

func GridClassStructSet() {
	GridClassStructOnce.Do(func() {
		GridClassStruct = gi.StructNew("Gtk", "GridClass")
	})
}

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

var GridPrivateStruct *gi.Struct
var GridPrivateStructOnce sync.Once

func GridPrivateStructSet() {
	GridPrivateStructOnce.Do(func() {
		GridPrivateStruct = gi.StructNew("Gtk", "GridPrivate")
	})
}

type GridPrivate struct {
	native uintptr
}

var HBoxClassStruct *gi.Struct
var HBoxClassStructOnce sync.Once

func HBoxClassStructSet() {
	HBoxClassStructOnce.Do(func() {
		HBoxClassStruct = gi.StructNew("Gtk", "HBoxClass")
	})
}

type HBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
}

var HButtonBoxClassStruct *gi.Struct
var HButtonBoxClassStructOnce sync.Once

func HButtonBoxClassStructSet() {
	HButtonBoxClassStructOnce.Do(func() {
		HButtonBoxClassStruct = gi.StructNew("Gtk", "HButtonBoxClass")
	})
}

type HButtonBoxClass struct {
	native      uintptr
	ParentClass *ButtonBoxClass
}

var HPanedClassStruct *gi.Struct
var HPanedClassStructOnce sync.Once

func HPanedClassStructSet() {
	HPanedClassStructOnce.Do(func() {
		HPanedClassStruct = gi.StructNew("Gtk", "HPanedClass")
	})
}

type HPanedClass struct {
	native      uintptr
	ParentClass *PanedClass
}

var HSVClassStruct *gi.Struct
var HSVClassStructOnce sync.Once

func HSVClassStructSet() {
	HSVClassStructOnce.Do(func() {
		HSVClassStruct = gi.StructNew("Gtk", "HSVClass")
	})
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

var HSVPrivateStruct *gi.Struct
var HSVPrivateStructOnce sync.Once

func HSVPrivateStructSet() {
	HSVPrivateStructOnce.Do(func() {
		HSVPrivateStruct = gi.StructNew("Gtk", "HSVPrivate")
	})
}

type HSVPrivate struct {
	native uintptr
}

var HScaleClassStruct *gi.Struct
var HScaleClassStructOnce sync.Once

func HScaleClassStructSet() {
	HScaleClassStructOnce.Do(func() {
		HScaleClassStruct = gi.StructNew("Gtk", "HScaleClass")
	})
}

type HScaleClass struct {
	native      uintptr
	ParentClass *ScaleClass
}

var HScrollbarClassStruct *gi.Struct
var HScrollbarClassStructOnce sync.Once

func HScrollbarClassStructSet() {
	HScrollbarClassStructOnce.Do(func() {
		HScrollbarClassStruct = gi.StructNew("Gtk", "HScrollbarClass")
	})
}

type HScrollbarClass struct {
	native      uintptr
	ParentClass *ScrollbarClass
}

var HSeparatorClassStruct *gi.Struct
var HSeparatorClassStructOnce sync.Once

func HSeparatorClassStructSet() {
	HSeparatorClassStructOnce.Do(func() {
		HSeparatorClassStruct = gi.StructNew("Gtk", "HSeparatorClass")
	})
}

type HSeparatorClass struct {
	native      uintptr
	ParentClass *SeparatorClass
}

var HandleBoxClassStruct *gi.Struct
var HandleBoxClassStructOnce sync.Once

func HandleBoxClassStructSet() {
	HandleBoxClassStructOnce.Do(func() {
		HandleBoxClassStruct = gi.StructNew("Gtk", "HandleBoxClass")
	})
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

var HandleBoxPrivateStruct *gi.Struct
var HandleBoxPrivateStructOnce sync.Once

func HandleBoxPrivateStructSet() {
	HandleBoxPrivateStructOnce.Do(func() {
		HandleBoxPrivateStruct = gi.StructNew("Gtk", "HandleBoxPrivate")
	})
}

type HandleBoxPrivate struct {
	native uintptr
}

var HeaderBarAccessibleClassStruct *gi.Struct
var HeaderBarAccessibleClassStructOnce sync.Once

func HeaderBarAccessibleClassStructSet() {
	HeaderBarAccessibleClassStructOnce.Do(func() {
		HeaderBarAccessibleClassStruct = gi.StructNew("Gtk", "HeaderBarAccessibleClass")
	})
}

type HeaderBarAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var HeaderBarAccessiblePrivateStruct *gi.Struct
var HeaderBarAccessiblePrivateStructOnce sync.Once

func HeaderBarAccessiblePrivateStructSet() {
	HeaderBarAccessiblePrivateStructOnce.Do(func() {
		HeaderBarAccessiblePrivateStruct = gi.StructNew("Gtk", "HeaderBarAccessiblePrivate")
	})
}

type HeaderBarAccessiblePrivate struct {
	native uintptr
}

var HeaderBarClassStruct *gi.Struct
var HeaderBarClassStructOnce sync.Once

func HeaderBarClassStructSet() {
	HeaderBarClassStructOnce.Do(func() {
		HeaderBarClassStruct = gi.StructNew("Gtk", "HeaderBarClass")
	})
}

type HeaderBarClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var HeaderBarPrivateStruct *gi.Struct
var HeaderBarPrivateStructOnce sync.Once

func HeaderBarPrivateStructSet() {
	HeaderBarPrivateStructOnce.Do(func() {
		HeaderBarPrivateStruct = gi.StructNew("Gtk", "HeaderBarPrivate")
	})
}

type HeaderBarPrivate struct {
	native uintptr
}

var IMContextClassStruct *gi.Struct
var IMContextClassStructOnce sync.Once

func IMContextClassStructSet() {
	IMContextClassStructOnce.Do(func() {
		IMContextClassStruct = gi.StructNew("Gtk", "IMContextClass")
	})
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

var IMContextInfoStruct *gi.Struct
var IMContextInfoStructOnce sync.Once

func IMContextInfoStructSet() {
	IMContextInfoStructOnce.Do(func() {
		IMContextInfoStruct = gi.StructNew("Gtk", "IMContextInfo")
	})
}

type IMContextInfo struct {
	native         uintptr
	ContextId      string
	ContextName    string
	Domain         string
	DomainDirname  string
	DefaultLocales string
}

var IMContextSimpleClassStruct *gi.Struct
var IMContextSimpleClassStructOnce sync.Once

func IMContextSimpleClassStructSet() {
	IMContextSimpleClassStructOnce.Do(func() {
		IMContextSimpleClassStruct = gi.StructNew("Gtk", "IMContextSimpleClass")
	})
}

type IMContextSimpleClass struct {
	native      uintptr
	ParentClass *IMContextClass
}

var IMContextSimplePrivateStruct *gi.Struct
var IMContextSimplePrivateStructOnce sync.Once

func IMContextSimplePrivateStructSet() {
	IMContextSimplePrivateStructOnce.Do(func() {
		IMContextSimplePrivateStruct = gi.StructNew("Gtk", "IMContextSimplePrivate")
	})
}

type IMContextSimplePrivate struct {
	native uintptr
}

var IMMulticontextClassStruct *gi.Struct
var IMMulticontextClassStructOnce sync.Once

func IMMulticontextClassStructSet() {
	IMMulticontextClassStructOnce.Do(func() {
		IMMulticontextClassStruct = gi.StructNew("Gtk", "IMMulticontextClass")
	})
}

type IMMulticontextClass struct {
	native      uintptr
	ParentClass *IMContextClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var IMMulticontextPrivateStruct *gi.Struct
var IMMulticontextPrivateStructOnce sync.Once

func IMMulticontextPrivateStructSet() {
	IMMulticontextPrivateStructOnce.Do(func() {
		IMMulticontextPrivateStruct = gi.StructNew("Gtk", "IMMulticontextPrivate")
	})
}

type IMMulticontextPrivate struct {
	native uintptr
}

var IconFactoryClassStruct *gi.Struct
var IconFactoryClassStructOnce sync.Once

func IconFactoryClassStructSet() {
	IconFactoryClassStructOnce.Do(func() {
		IconFactoryClassStruct = gi.StructNew("Gtk", "IconFactoryClass")
	})
}

type IconFactoryClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var IconFactoryPrivateStruct *gi.Struct
var IconFactoryPrivateStructOnce sync.Once

func IconFactoryPrivateStructSet() {
	IconFactoryPrivateStructOnce.Do(func() {
		IconFactoryPrivateStruct = gi.StructNew("Gtk", "IconFactoryPrivate")
	})
}

type IconFactoryPrivate struct {
	native uintptr
}

var IconInfoClassStruct *gi.Struct
var IconInfoClassStructOnce sync.Once

func IconInfoClassStructSet() {
	IconInfoClassStructOnce.Do(func() {
		IconInfoClassStruct = gi.StructNew("Gtk", "IconInfoClass")
	})
}

type IconInfoClass struct {
	native uintptr
}

var IconSetStruct *gi.Struct
var IconSetStructOnce sync.Once

func IconSetStructSet() {
	IconSetStructOnce.Do(func() {
		IconSetStruct = gi.StructNew("Gtk", "IconSet")
	})
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

	retGo := &IconSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_set_new_from_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_icon_set_add_source' : parameter 'source' of type 'IconSource' not supported

var copyIconSetInvoker *gi.Function

// Copy is a representation of the C type gtk_icon_set_copy.
func (recv *IconSet) Copy() *IconSet {
	if copyIconSetInvoker == nil {
		copyIconSetInvoker = gi.StructFunctionInvokerNew("Gtk", "IconSet", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyIconSetInvoker.Invoke(inArgs[:], nil)

	retGo := &IconSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_set_get_sizes' : parameter 'sizes' has no type

var refIconSetInvoker *gi.Function

// Ref is a representation of the C type gtk_icon_set_ref.
func (recv *IconSet) Ref() *IconSet {
	if refIconSetInvoker == nil {
		refIconSetInvoker = gi.StructFunctionInvokerNew("Gtk", "IconSet", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refIconSetInvoker.Invoke(inArgs[:], nil)

	retGo := &IconSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_set_render_icon' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_icon_set_render_icon_pixbuf' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_icon_set_render_icon_surface' : parameter 'context' of type 'StyleContext' not supported

var unrefIconSetInvoker *gi.Function

// Unref is a representation of the C type gtk_icon_set_unref.
func (recv *IconSet) Unref() {
	if unrefIconSetInvoker == nil {
		unrefIconSetInvoker = gi.StructFunctionInvokerNew("Gtk", "IconSet", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefIconSetInvoker.Invoke(inArgs[:], nil)

}

var IconSourceStruct *gi.Struct
var IconSourceStructOnce sync.Once

func IconSourceStructSet() {
	IconSourceStructOnce.Do(func() {
		IconSourceStruct = gi.StructNew("Gtk", "IconSource")
	})
}

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

	retGo := &IconSource{native: ret.Pointer()}

	return retGo
}

var copyIconSourceInvoker *gi.Function

// Copy is a representation of the C type gtk_icon_source_copy.
func (recv *IconSource) Copy() *IconSource {
	if copyIconSourceInvoker == nil {
		copyIconSourceInvoker = gi.StructFunctionInvokerNew("Gtk", "IconSource", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyIconSourceInvoker.Invoke(inArgs[:], nil)

	retGo := &IconSource{native: ret.Pointer()}

	return retGo
}

var freeIconSourceInvoker *gi.Function

// Free is a representation of the C type gtk_icon_source_free.
func (recv *IconSource) Free() {
	if freeIconSourceInvoker == nil {
		freeIconSourceInvoker = gi.StructFunctionInvokerNew("Gtk", "IconSource", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeIconSourceInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_icon_source_get_direction' : return type 'TextDirection' not supported

// UNSUPPORTED : C value 'gtk_icon_source_get_direction_wildcarded' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_icon_source_get_filename' : return type 'filename' not supported

var getIconNameIconSourceInvoker *gi.Function

// GetIconName is a representation of the C type gtk_icon_source_get_icon_name.
func (recv *IconSource) GetIconName() string {
	if getIconNameIconSourceInvoker == nil {
		getIconNameIconSourceInvoker = gi.StructFunctionInvokerNew("Gtk", "IconSource", "get_icon_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIconNameIconSourceInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_source_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_icon_source_get_size' : return type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_source_get_size_wildcarded' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_icon_source_get_state' : return type 'StateType' not supported

// UNSUPPORTED : C value 'gtk_icon_source_get_state_wildcarded' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_direction' : parameter 'direction' of type 'TextDirection' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_direction_wildcarded' : parameter 'setting' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_filename' : parameter 'filename' of type 'filename' not supported

var setIconNameIconSourceInvoker *gi.Function

// SetIconName is a representation of the C type gtk_icon_source_set_icon_name.
func (recv *IconSource) SetIconName(iconName string) {
	if setIconNameIconSourceInvoker == nil {
		setIconNameIconSourceInvoker = gi.StructFunctionInvokerNew("Gtk", "IconSource", "set_icon_name")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(iconName)

	setIconNameIconSourceInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_icon_source_set_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_size' : parameter 'size' of type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_size_wildcarded' : parameter 'setting' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_state' : parameter 'state' of type 'StateType' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_state_wildcarded' : parameter 'setting' of type 'gboolean' not supported

var IconThemeClassStruct *gi.Struct
var IconThemeClassStructOnce sync.Once

func IconThemeClassStructSet() {
	IconThemeClassStructOnce.Do(func() {
		IconThemeClassStruct = gi.StructNew("Gtk", "IconThemeClass")
	})
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

var IconThemePrivateStruct *gi.Struct
var IconThemePrivateStructOnce sync.Once

func IconThemePrivateStructSet() {
	IconThemePrivateStructOnce.Do(func() {
		IconThemePrivateStruct = gi.StructNew("Gtk", "IconThemePrivate")
	})
}

type IconThemePrivate struct {
	native uintptr
}

var IconViewAccessibleClassStruct *gi.Struct
var IconViewAccessibleClassStructOnce sync.Once

func IconViewAccessibleClassStructSet() {
	IconViewAccessibleClassStructOnce.Do(func() {
		IconViewAccessibleClassStruct = gi.StructNew("Gtk", "IconViewAccessibleClass")
	})
}

type IconViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var IconViewAccessiblePrivateStruct *gi.Struct
var IconViewAccessiblePrivateStructOnce sync.Once

func IconViewAccessiblePrivateStructSet() {
	IconViewAccessiblePrivateStructOnce.Do(func() {
		IconViewAccessiblePrivateStruct = gi.StructNew("Gtk", "IconViewAccessiblePrivate")
	})
}

type IconViewAccessiblePrivate struct {
	native uintptr
}

var IconViewClassStruct *gi.Struct
var IconViewClassStructOnce sync.Once

func IconViewClassStructSet() {
	IconViewClassStructOnce.Do(func() {
		IconViewClassStruct = gi.StructNew("Gtk", "IconViewClass")
	})
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

var IconViewPrivateStruct *gi.Struct
var IconViewPrivateStructOnce sync.Once

func IconViewPrivateStructSet() {
	IconViewPrivateStructOnce.Do(func() {
		IconViewPrivateStruct = gi.StructNew("Gtk", "IconViewPrivate")
	})
}

type IconViewPrivate struct {
	native uintptr
}

var ImageAccessibleClassStruct *gi.Struct
var ImageAccessibleClassStructOnce sync.Once

func ImageAccessibleClassStructSet() {
	ImageAccessibleClassStructOnce.Do(func() {
		ImageAccessibleClassStruct = gi.StructNew("Gtk", "ImageAccessibleClass")
	})
}

type ImageAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var ImageAccessiblePrivateStruct *gi.Struct
var ImageAccessiblePrivateStructOnce sync.Once

func ImageAccessiblePrivateStructSet() {
	ImageAccessiblePrivateStructOnce.Do(func() {
		ImageAccessiblePrivateStruct = gi.StructNew("Gtk", "ImageAccessiblePrivate")
	})
}

type ImageAccessiblePrivate struct {
	native uintptr
}

var ImageCellAccessibleClassStruct *gi.Struct
var ImageCellAccessibleClassStructOnce sync.Once

func ImageCellAccessibleClassStructSet() {
	ImageCellAccessibleClassStructOnce.Do(func() {
		ImageCellAccessibleClassStruct = gi.StructNew("Gtk", "ImageCellAccessibleClass")
	})
}

type ImageCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

var ImageCellAccessiblePrivateStruct *gi.Struct
var ImageCellAccessiblePrivateStructOnce sync.Once

func ImageCellAccessiblePrivateStructSet() {
	ImageCellAccessiblePrivateStructOnce.Do(func() {
		ImageCellAccessiblePrivateStruct = gi.StructNew("Gtk", "ImageCellAccessiblePrivate")
	})
}

type ImageCellAccessiblePrivate struct {
	native uintptr
}

var ImageClassStruct *gi.Struct
var ImageClassStructOnce sync.Once

func ImageClassStructSet() {
	ImageClassStructOnce.Do(func() {
		ImageClassStruct = gi.StructNew("Gtk", "ImageClass")
	})
}

type ImageClass struct {
	native      uintptr
	ParentClass *MiscClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ImageMenuItemClassStruct *gi.Struct
var ImageMenuItemClassStructOnce sync.Once

func ImageMenuItemClassStructSet() {
	ImageMenuItemClassStructOnce.Do(func() {
		ImageMenuItemClassStruct = gi.StructNew("Gtk", "ImageMenuItemClass")
	})
}

type ImageMenuItemClass struct {
	native      uintptr
	ParentClass *MenuItemClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ImageMenuItemPrivateStruct *gi.Struct
var ImageMenuItemPrivateStructOnce sync.Once

func ImageMenuItemPrivateStructSet() {
	ImageMenuItemPrivateStructOnce.Do(func() {
		ImageMenuItemPrivateStruct = gi.StructNew("Gtk", "ImageMenuItemPrivate")
	})
}

type ImageMenuItemPrivate struct {
	native uintptr
}

var ImagePrivateStruct *gi.Struct
var ImagePrivateStructOnce sync.Once

func ImagePrivateStructSet() {
	ImagePrivateStructOnce.Do(func() {
		ImagePrivateStruct = gi.StructNew("Gtk", "ImagePrivate")
	})
}

type ImagePrivate struct {
	native uintptr
}

var InfoBarClassStruct *gi.Struct
var InfoBarClassStructOnce sync.Once

func InfoBarClassStructSet() {
	InfoBarClassStructOnce.Do(func() {
		InfoBarClassStruct = gi.StructNew("Gtk", "InfoBarClass")
	})
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

var InfoBarPrivateStruct *gi.Struct
var InfoBarPrivateStructOnce sync.Once

func InfoBarPrivateStructSet() {
	InfoBarPrivateStructOnce.Do(func() {
		InfoBarPrivateStruct = gi.StructNew("Gtk", "InfoBarPrivate")
	})
}

type InfoBarPrivate struct {
	native uintptr
}

var InvisibleClassStruct *gi.Struct
var InvisibleClassStructOnce sync.Once

func InvisibleClassStructSet() {
	InvisibleClassStructOnce.Do(func() {
		InvisibleClassStruct = gi.StructNew("Gtk", "InvisibleClass")
	})
}

type InvisibleClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var InvisiblePrivateStruct *gi.Struct
var InvisiblePrivateStructOnce sync.Once

func InvisiblePrivateStructSet() {
	InvisiblePrivateStructOnce.Do(func() {
		InvisiblePrivateStruct = gi.StructNew("Gtk", "InvisiblePrivate")
	})
}

type InvisiblePrivate struct {
	native uintptr
}

var LabelAccessibleClassStruct *gi.Struct
var LabelAccessibleClassStructOnce sync.Once

func LabelAccessibleClassStructSet() {
	LabelAccessibleClassStructOnce.Do(func() {
		LabelAccessibleClassStruct = gi.StructNew("Gtk", "LabelAccessibleClass")
	})
}

type LabelAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var LabelAccessiblePrivateStruct *gi.Struct
var LabelAccessiblePrivateStructOnce sync.Once

func LabelAccessiblePrivateStructSet() {
	LabelAccessiblePrivateStructOnce.Do(func() {
		LabelAccessiblePrivateStruct = gi.StructNew("Gtk", "LabelAccessiblePrivate")
	})
}

type LabelAccessiblePrivate struct {
	native uintptr
}

var LabelClassStruct *gi.Struct
var LabelClassStructOnce sync.Once

func LabelClassStructSet() {
	LabelClassStructOnce.Do(func() {
		LabelClassStruct = gi.StructNew("Gtk", "LabelClass")
	})
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

var LabelPrivateStruct *gi.Struct
var LabelPrivateStructOnce sync.Once

func LabelPrivateStructSet() {
	LabelPrivateStructOnce.Do(func() {
		LabelPrivateStruct = gi.StructNew("Gtk", "LabelPrivate")
	})
}

type LabelPrivate struct {
	native uintptr
}

var LabelSelectionInfoStruct *gi.Struct
var LabelSelectionInfoStructOnce sync.Once

func LabelSelectionInfoStructSet() {
	LabelSelectionInfoStructOnce.Do(func() {
		LabelSelectionInfoStruct = gi.StructNew("Gtk", "LabelSelectionInfo")
	})
}

type LabelSelectionInfo struct {
	native uintptr
}

var LayoutClassStruct *gi.Struct
var LayoutClassStructOnce sync.Once

func LayoutClassStructSet() {
	LayoutClassStructOnce.Do(func() {
		LayoutClassStruct = gi.StructNew("Gtk", "LayoutClass")
	})
}

type LayoutClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var LayoutPrivateStruct *gi.Struct
var LayoutPrivateStructOnce sync.Once

func LayoutPrivateStructSet() {
	LayoutPrivateStructOnce.Do(func() {
		LayoutPrivateStruct = gi.StructNew("Gtk", "LayoutPrivate")
	})
}

type LayoutPrivate struct {
	native uintptr
}

var LevelBarAccessibleClassStruct *gi.Struct
var LevelBarAccessibleClassStructOnce sync.Once

func LevelBarAccessibleClassStructSet() {
	LevelBarAccessibleClassStructOnce.Do(func() {
		LevelBarAccessibleClassStruct = gi.StructNew("Gtk", "LevelBarAccessibleClass")
	})
}

type LevelBarAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var LevelBarAccessiblePrivateStruct *gi.Struct
var LevelBarAccessiblePrivateStructOnce sync.Once

func LevelBarAccessiblePrivateStructSet() {
	LevelBarAccessiblePrivateStructOnce.Do(func() {
		LevelBarAccessiblePrivateStruct = gi.StructNew("Gtk", "LevelBarAccessiblePrivate")
	})
}

type LevelBarAccessiblePrivate struct {
	native uintptr
}

var LevelBarClassStruct *gi.Struct
var LevelBarClassStructOnce sync.Once

func LevelBarClassStructSet() {
	LevelBarClassStructOnce.Do(func() {
		LevelBarClassStruct = gi.StructNew("Gtk", "LevelBarClass")
	})
}

type LevelBarClass struct {
	native uintptr
	// UNSUPPORTED : C value 'offset_changed' : missing Type
}

var LevelBarPrivateStruct *gi.Struct
var LevelBarPrivateStructOnce sync.Once

func LevelBarPrivateStructSet() {
	LevelBarPrivateStructOnce.Do(func() {
		LevelBarPrivateStruct = gi.StructNew("Gtk", "LevelBarPrivate")
	})
}

type LevelBarPrivate struct {
	native uintptr
}

var LinkButtonAccessibleClassStruct *gi.Struct
var LinkButtonAccessibleClassStructOnce sync.Once

func LinkButtonAccessibleClassStructSet() {
	LinkButtonAccessibleClassStructOnce.Do(func() {
		LinkButtonAccessibleClassStruct = gi.StructNew("Gtk", "LinkButtonAccessibleClass")
	})
}

type LinkButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var LinkButtonAccessiblePrivateStruct *gi.Struct
var LinkButtonAccessiblePrivateStructOnce sync.Once

func LinkButtonAccessiblePrivateStructSet() {
	LinkButtonAccessiblePrivateStructOnce.Do(func() {
		LinkButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "LinkButtonAccessiblePrivate")
	})
}

type LinkButtonAccessiblePrivate struct {
	native uintptr
}

var LinkButtonClassStruct *gi.Struct
var LinkButtonClassStructOnce sync.Once

func LinkButtonClassStructSet() {
	LinkButtonClassStructOnce.Do(func() {
		LinkButtonClassStruct = gi.StructNew("Gtk", "LinkButtonClass")
	})
}

type LinkButtonClass struct {
	native uintptr
	// UNSUPPORTED : C value 'activate_link' : missing Type
	// UNSUPPORTED : C value '_gtk_padding1' : missing Type
	// UNSUPPORTED : C value '_gtk_padding2' : missing Type
	// UNSUPPORTED : C value '_gtk_padding3' : missing Type
	// UNSUPPORTED : C value '_gtk_padding4' : missing Type
}

var LinkButtonPrivateStruct *gi.Struct
var LinkButtonPrivateStructOnce sync.Once

func LinkButtonPrivateStructSet() {
	LinkButtonPrivateStructOnce.Do(func() {
		LinkButtonPrivateStruct = gi.StructNew("Gtk", "LinkButtonPrivate")
	})
}

type LinkButtonPrivate struct {
	native uintptr
}

var ListBoxAccessibleClassStruct *gi.Struct
var ListBoxAccessibleClassStructOnce sync.Once

func ListBoxAccessibleClassStructSet() {
	ListBoxAccessibleClassStructOnce.Do(func() {
		ListBoxAccessibleClassStruct = gi.StructNew("Gtk", "ListBoxAccessibleClass")
	})
}

type ListBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var ListBoxAccessiblePrivateStruct *gi.Struct
var ListBoxAccessiblePrivateStructOnce sync.Once

func ListBoxAccessiblePrivateStructSet() {
	ListBoxAccessiblePrivateStructOnce.Do(func() {
		ListBoxAccessiblePrivateStruct = gi.StructNew("Gtk", "ListBoxAccessiblePrivate")
	})
}

type ListBoxAccessiblePrivate struct {
	native uintptr
}

var ListBoxClassStruct *gi.Struct
var ListBoxClassStructOnce sync.Once

func ListBoxClassStructSet() {
	ListBoxClassStructOnce.Do(func() {
		ListBoxClassStruct = gi.StructNew("Gtk", "ListBoxClass")
	})
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

var ListBoxRowAccessibleClassStruct *gi.Struct
var ListBoxRowAccessibleClassStructOnce sync.Once

func ListBoxRowAccessibleClassStructSet() {
	ListBoxRowAccessibleClassStructOnce.Do(func() {
		ListBoxRowAccessibleClassStruct = gi.StructNew("Gtk", "ListBoxRowAccessibleClass")
	})
}

type ListBoxRowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var ListBoxRowClassStruct *gi.Struct
var ListBoxRowClassStructOnce sync.Once

func ListBoxRowClassStructSet() {
	ListBoxRowClassStructOnce.Do(func() {
		ListBoxRowClassStruct = gi.StructNew("Gtk", "ListBoxRowClass")
	})
}

type ListBoxRowClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'activate' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
}

var ListStoreClassStruct *gi.Struct
var ListStoreClassStructOnce sync.Once

func ListStoreClassStructSet() {
	ListStoreClassStructOnce.Do(func() {
		ListStoreClassStruct = gi.StructNew("Gtk", "ListStoreClass")
	})
}

type ListStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ListStorePrivateStruct *gi.Struct
var ListStorePrivateStructOnce sync.Once

func ListStorePrivateStructSet() {
	ListStorePrivateStructOnce.Do(func() {
		ListStorePrivateStruct = gi.StructNew("Gtk", "ListStorePrivate")
	})
}

type ListStorePrivate struct {
	native uintptr
}

var LockButtonAccessibleClassStruct *gi.Struct
var LockButtonAccessibleClassStructOnce sync.Once

func LockButtonAccessibleClassStructSet() {
	LockButtonAccessibleClassStructOnce.Do(func() {
		LockButtonAccessibleClassStruct = gi.StructNew("Gtk", "LockButtonAccessibleClass")
	})
}

type LockButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var LockButtonAccessiblePrivateStruct *gi.Struct
var LockButtonAccessiblePrivateStructOnce sync.Once

func LockButtonAccessiblePrivateStructSet() {
	LockButtonAccessiblePrivateStructOnce.Do(func() {
		LockButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "LockButtonAccessiblePrivate")
	})
}

type LockButtonAccessiblePrivate struct {
	native uintptr
}

var LockButtonClassStruct *gi.Struct
var LockButtonClassStructOnce sync.Once

func LockButtonClassStructSet() {
	LockButtonClassStructOnce.Do(func() {
		LockButtonClassStruct = gi.StructNew("Gtk", "LockButtonClass")
	})
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

var LockButtonPrivateStruct *gi.Struct
var LockButtonPrivateStructOnce sync.Once

func LockButtonPrivateStructSet() {
	LockButtonPrivateStructOnce.Do(func() {
		LockButtonPrivateStruct = gi.StructNew("Gtk", "LockButtonPrivate")
	})
}

type LockButtonPrivate struct {
	native uintptr
}

var MenuAccessibleClassStruct *gi.Struct
var MenuAccessibleClassStructOnce sync.Once

func MenuAccessibleClassStructSet() {
	MenuAccessibleClassStructOnce.Do(func() {
		MenuAccessibleClassStruct = gi.StructNew("Gtk", "MenuAccessibleClass")
	})
}

type MenuAccessibleClass struct {
	native      uintptr
	ParentClass *MenuShellAccessibleClass
}

var MenuAccessiblePrivateStruct *gi.Struct
var MenuAccessiblePrivateStructOnce sync.Once

func MenuAccessiblePrivateStructSet() {
	MenuAccessiblePrivateStructOnce.Do(func() {
		MenuAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuAccessiblePrivate")
	})
}

type MenuAccessiblePrivate struct {
	native uintptr
}

var MenuBarClassStruct *gi.Struct
var MenuBarClassStructOnce sync.Once

func MenuBarClassStructSet() {
	MenuBarClassStructOnce.Do(func() {
		MenuBarClassStruct = gi.StructNew("Gtk", "MenuBarClass")
	})
}

type MenuBarClass struct {
	native      uintptr
	ParentClass *MenuShellClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var MenuBarPrivateStruct *gi.Struct
var MenuBarPrivateStructOnce sync.Once

func MenuBarPrivateStructSet() {
	MenuBarPrivateStructOnce.Do(func() {
		MenuBarPrivateStruct = gi.StructNew("Gtk", "MenuBarPrivate")
	})
}

type MenuBarPrivate struct {
	native uintptr
}

var MenuButtonAccessibleClassStruct *gi.Struct
var MenuButtonAccessibleClassStructOnce sync.Once

func MenuButtonAccessibleClassStructSet() {
	MenuButtonAccessibleClassStructOnce.Do(func() {
		MenuButtonAccessibleClassStruct = gi.StructNew("Gtk", "MenuButtonAccessibleClass")
	})
}

type MenuButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ToggleButtonAccessibleClass
}

var MenuButtonAccessiblePrivateStruct *gi.Struct
var MenuButtonAccessiblePrivateStructOnce sync.Once

func MenuButtonAccessiblePrivateStructSet() {
	MenuButtonAccessiblePrivateStructOnce.Do(func() {
		MenuButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuButtonAccessiblePrivate")
	})
}

type MenuButtonAccessiblePrivate struct {
	native uintptr
}

var MenuButtonClassStruct *gi.Struct
var MenuButtonClassStructOnce sync.Once

func MenuButtonClassStructSet() {
	MenuButtonClassStructOnce.Do(func() {
		MenuButtonClassStruct = gi.StructNew("Gtk", "MenuButtonClass")
	})
}

type MenuButtonClass struct {
	native      uintptr
	ParentClass *ToggleButtonClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var MenuButtonPrivateStruct *gi.Struct
var MenuButtonPrivateStructOnce sync.Once

func MenuButtonPrivateStructSet() {
	MenuButtonPrivateStructOnce.Do(func() {
		MenuButtonPrivateStruct = gi.StructNew("Gtk", "MenuButtonPrivate")
	})
}

type MenuButtonPrivate struct {
	native uintptr
}

var MenuClassStruct *gi.Struct
var MenuClassStructOnce sync.Once

func MenuClassStructSet() {
	MenuClassStructOnce.Do(func() {
		MenuClassStruct = gi.StructNew("Gtk", "MenuClass")
	})
}

type MenuClass struct {
	native      uintptr
	ParentClass *MenuShellClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var MenuItemAccessibleClassStruct *gi.Struct
var MenuItemAccessibleClassStructOnce sync.Once

func MenuItemAccessibleClassStructSet() {
	MenuItemAccessibleClassStructOnce.Do(func() {
		MenuItemAccessibleClassStruct = gi.StructNew("Gtk", "MenuItemAccessibleClass")
	})
}

type MenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var MenuItemAccessiblePrivateStruct *gi.Struct
var MenuItemAccessiblePrivateStructOnce sync.Once

func MenuItemAccessiblePrivateStructSet() {
	MenuItemAccessiblePrivateStructOnce.Do(func() {
		MenuItemAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuItemAccessiblePrivate")
	})
}

type MenuItemAccessiblePrivate struct {
	native uintptr
}

var MenuItemClassStruct *gi.Struct
var MenuItemClassStructOnce sync.Once

func MenuItemClassStructSet() {
	MenuItemClassStructOnce.Do(func() {
		MenuItemClassStruct = gi.StructNew("Gtk", "MenuItemClass")
	})
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

var MenuItemPrivateStruct *gi.Struct
var MenuItemPrivateStructOnce sync.Once

func MenuItemPrivateStructSet() {
	MenuItemPrivateStructOnce.Do(func() {
		MenuItemPrivateStruct = gi.StructNew("Gtk", "MenuItemPrivate")
	})
}

type MenuItemPrivate struct {
	native uintptr
}

var MenuPrivateStruct *gi.Struct
var MenuPrivateStructOnce sync.Once

func MenuPrivateStructSet() {
	MenuPrivateStructOnce.Do(func() {
		MenuPrivateStruct = gi.StructNew("Gtk", "MenuPrivate")
	})
}

type MenuPrivate struct {
	native uintptr
}

var MenuShellAccessibleClassStruct *gi.Struct
var MenuShellAccessibleClassStructOnce sync.Once

func MenuShellAccessibleClassStructSet() {
	MenuShellAccessibleClassStructOnce.Do(func() {
		MenuShellAccessibleClassStruct = gi.StructNew("Gtk", "MenuShellAccessibleClass")
	})
}

type MenuShellAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var MenuShellAccessiblePrivateStruct *gi.Struct
var MenuShellAccessiblePrivateStructOnce sync.Once

func MenuShellAccessiblePrivateStructSet() {
	MenuShellAccessiblePrivateStructOnce.Do(func() {
		MenuShellAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuShellAccessiblePrivate")
	})
}

type MenuShellAccessiblePrivate struct {
	native uintptr
}

var MenuShellClassStruct *gi.Struct
var MenuShellClassStructOnce sync.Once

func MenuShellClassStructSet() {
	MenuShellClassStructOnce.Do(func() {
		MenuShellClassStruct = gi.StructNew("Gtk", "MenuShellClass")
	})
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

var MenuShellPrivateStruct *gi.Struct
var MenuShellPrivateStructOnce sync.Once

func MenuShellPrivateStructSet() {
	MenuShellPrivateStructOnce.Do(func() {
		MenuShellPrivateStruct = gi.StructNew("Gtk", "MenuShellPrivate")
	})
}

type MenuShellPrivate struct {
	native uintptr
}

var MenuToolButtonClassStruct *gi.Struct
var MenuToolButtonClassStructOnce sync.Once

func MenuToolButtonClassStructSet() {
	MenuToolButtonClassStructOnce.Do(func() {
		MenuToolButtonClassStruct = gi.StructNew("Gtk", "MenuToolButtonClass")
	})
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

var MenuToolButtonPrivateStruct *gi.Struct
var MenuToolButtonPrivateStructOnce sync.Once

func MenuToolButtonPrivateStructSet() {
	MenuToolButtonPrivateStructOnce.Do(func() {
		MenuToolButtonPrivateStruct = gi.StructNew("Gtk", "MenuToolButtonPrivate")
	})
}

type MenuToolButtonPrivate struct {
	native uintptr
}

var MessageDialogClassStruct *gi.Struct
var MessageDialogClassStructOnce sync.Once

func MessageDialogClassStructSet() {
	MessageDialogClassStructOnce.Do(func() {
		MessageDialogClassStruct = gi.StructNew("Gtk", "MessageDialogClass")
	})
}

type MessageDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var MessageDialogPrivateStruct *gi.Struct
var MessageDialogPrivateStructOnce sync.Once

func MessageDialogPrivateStructSet() {
	MessageDialogPrivateStructOnce.Do(func() {
		MessageDialogPrivateStruct = gi.StructNew("Gtk", "MessageDialogPrivate")
	})
}

type MessageDialogPrivate struct {
	native uintptr
}

var MiscClassStruct *gi.Struct
var MiscClassStructOnce sync.Once

func MiscClassStructSet() {
	MiscClassStructOnce.Do(func() {
		MiscClassStruct = gi.StructNew("Gtk", "MiscClass")
	})
}

type MiscClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var MiscPrivateStruct *gi.Struct
var MiscPrivateStructOnce sync.Once

func MiscPrivateStructSet() {
	MiscPrivateStructOnce.Do(func() {
		MiscPrivateStruct = gi.StructNew("Gtk", "MiscPrivate")
	})
}

type MiscPrivate struct {
	native uintptr
}

var MountOperationClassStruct *gi.Struct
var MountOperationClassStructOnce sync.Once

func MountOperationClassStructSet() {
	MountOperationClassStructOnce.Do(func() {
		MountOperationClassStruct = gi.StructNew("Gtk", "MountOperationClass")
	})
}

type MountOperationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.MountOperationClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var MountOperationPrivateStruct *gi.Struct
var MountOperationPrivateStructOnce sync.Once

func MountOperationPrivateStructSet() {
	MountOperationPrivateStructOnce.Do(func() {
		MountOperationPrivateStruct = gi.StructNew("Gtk", "MountOperationPrivate")
	})
}

type MountOperationPrivate struct {
	native uintptr
}

var NativeDialogClassStruct *gi.Struct
var NativeDialogClassStructOnce sync.Once

func NativeDialogClassStructSet() {
	NativeDialogClassStructOnce.Do(func() {
		NativeDialogClassStruct = gi.StructNew("Gtk", "NativeDialogClass")
	})
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

var NotebookAccessibleClassStruct *gi.Struct
var NotebookAccessibleClassStructOnce sync.Once

func NotebookAccessibleClassStructSet() {
	NotebookAccessibleClassStructOnce.Do(func() {
		NotebookAccessibleClassStruct = gi.StructNew("Gtk", "NotebookAccessibleClass")
	})
}

type NotebookAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var NotebookAccessiblePrivateStruct *gi.Struct
var NotebookAccessiblePrivateStructOnce sync.Once

func NotebookAccessiblePrivateStructSet() {
	NotebookAccessiblePrivateStructOnce.Do(func() {
		NotebookAccessiblePrivateStruct = gi.StructNew("Gtk", "NotebookAccessiblePrivate")
	})
}

type NotebookAccessiblePrivate struct {
	native uintptr
}

var NotebookClassStruct *gi.Struct
var NotebookClassStructOnce sync.Once

func NotebookClassStructSet() {
	NotebookClassStructOnce.Do(func() {
		NotebookClassStruct = gi.StructNew("Gtk", "NotebookClass")
	})
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

var NotebookPageAccessibleClassStruct *gi.Struct
var NotebookPageAccessibleClassStructOnce sync.Once

func NotebookPageAccessibleClassStructSet() {
	NotebookPageAccessibleClassStructOnce.Do(func() {
		NotebookPageAccessibleClassStruct = gi.StructNew("Gtk", "NotebookPageAccessibleClass")
	})
}

type NotebookPageAccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Atk.ObjectClass'
}

var NotebookPageAccessiblePrivateStruct *gi.Struct
var NotebookPageAccessiblePrivateStructOnce sync.Once

func NotebookPageAccessiblePrivateStructSet() {
	NotebookPageAccessiblePrivateStructOnce.Do(func() {
		NotebookPageAccessiblePrivateStruct = gi.StructNew("Gtk", "NotebookPageAccessiblePrivate")
	})
}

type NotebookPageAccessiblePrivate struct {
	native uintptr
}

var NotebookPrivateStruct *gi.Struct
var NotebookPrivateStructOnce sync.Once

func NotebookPrivateStructSet() {
	NotebookPrivateStructOnce.Do(func() {
		NotebookPrivateStruct = gi.StructNew("Gtk", "NotebookPrivate")
	})
}

type NotebookPrivate struct {
	native uintptr
}

var NumerableIconClassStruct *gi.Struct
var NumerableIconClassStructOnce sync.Once

func NumerableIconClassStructSet() {
	NumerableIconClassStructOnce.Do(func() {
		NumerableIconClassStruct = gi.StructNew("Gtk", "NumerableIconClass")
	})
}

type NumerableIconClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.EmblemedIconClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var NumerableIconPrivateStruct *gi.Struct
var NumerableIconPrivateStructOnce sync.Once

func NumerableIconPrivateStructSet() {
	NumerableIconPrivateStructOnce.Do(func() {
		NumerableIconPrivateStruct = gi.StructNew("Gtk", "NumerableIconPrivate")
	})
}

type NumerableIconPrivate struct {
	native uintptr
}

var OffscreenWindowClassStruct *gi.Struct
var OffscreenWindowClassStructOnce sync.Once

func OffscreenWindowClassStructSet() {
	OffscreenWindowClassStructOnce.Do(func() {
		OffscreenWindowClassStruct = gi.StructNew("Gtk", "OffscreenWindowClass")
	})
}

type OffscreenWindowClass struct {
	native      uintptr
	ParentClass *WindowClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var OrientableIfaceStruct *gi.Struct
var OrientableIfaceStructOnce sync.Once

func OrientableIfaceStructSet() {
	OrientableIfaceStructOnce.Do(func() {
		OrientableIfaceStruct = gi.StructNew("Gtk", "OrientableIface")
	})
}

type OrientableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'
}

var OverlayClassStruct *gi.Struct
var OverlayClassStructOnce sync.Once

func OverlayClassStructSet() {
	OverlayClassStructOnce.Do(func() {
		OverlayClassStruct = gi.StructNew("Gtk", "OverlayClass")
	})
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

var OverlayPrivateStruct *gi.Struct
var OverlayPrivateStructOnce sync.Once

func OverlayPrivateStructSet() {
	OverlayPrivateStructOnce.Do(func() {
		OverlayPrivateStruct = gi.StructNew("Gtk", "OverlayPrivate")
	})
}

type OverlayPrivate struct {
	native uintptr
}

var PadActionEntryStruct *gi.Struct
var PadActionEntryStructOnce sync.Once

func PadActionEntryStructSet() {
	PadActionEntryStructOnce.Do(func() {
		PadActionEntryStruct = gi.StructNew("Gtk", "PadActionEntry")
	})
}

type PadActionEntry struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'PadActionType'
	Index      int32
	Mode       int32
	Label      string
	ActionName string
}

var PadControllerClassStruct *gi.Struct
var PadControllerClassStructOnce sync.Once

func PadControllerClassStructSet() {
	PadControllerClassStructOnce.Do(func() {
		PadControllerClassStruct = gi.StructNew("Gtk", "PadControllerClass")
	})
}

type PadControllerClass struct {
	native uintptr
}

var PageRangeStruct *gi.Struct
var PageRangeStructOnce sync.Once

func PageRangeStructSet() {
	PageRangeStructOnce.Do(func() {
		PageRangeStruct = gi.StructNew("Gtk", "PageRange")
	})
}

type PageRange struct {
	native uintptr
	Start  int32
	End    int32
}

var PanedAccessibleClassStruct *gi.Struct
var PanedAccessibleClassStructOnce sync.Once

func PanedAccessibleClassStructSet() {
	PanedAccessibleClassStructOnce.Do(func() {
		PanedAccessibleClassStruct = gi.StructNew("Gtk", "PanedAccessibleClass")
	})
}

type PanedAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var PanedAccessiblePrivateStruct *gi.Struct
var PanedAccessiblePrivateStructOnce sync.Once

func PanedAccessiblePrivateStructSet() {
	PanedAccessiblePrivateStructOnce.Do(func() {
		PanedAccessiblePrivateStruct = gi.StructNew("Gtk", "PanedAccessiblePrivate")
	})
}

type PanedAccessiblePrivate struct {
	native uintptr
}

var PanedClassStruct *gi.Struct
var PanedClassStructOnce sync.Once

func PanedClassStructSet() {
	PanedClassStructOnce.Do(func() {
		PanedClassStruct = gi.StructNew("Gtk", "PanedClass")
	})
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

var PanedPrivateStruct *gi.Struct
var PanedPrivateStructOnce sync.Once

func PanedPrivateStructSet() {
	PanedPrivateStructOnce.Do(func() {
		PanedPrivateStruct = gi.StructNew("Gtk", "PanedPrivate")
	})
}

type PanedPrivate struct {
	native uintptr
}

var PaperSizeStruct *gi.Struct
var PaperSizeStructOnce sync.Once

func PaperSizeStructSet() {
	PaperSizeStructOnce.Do(func() {
		PaperSizeStruct = gi.StructNew("Gtk", "PaperSize")
	})
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

	retGo := &PaperSize{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_new_custom' : parameter 'width' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_gvariant' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_ipp' : parameter 'width' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_key_file' : parameter 'key_file' of type 'GLib.KeyFile' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_ppd' : parameter 'width' of type 'gdouble' not supported

var copyPaperSizeInvoker *gi.Function

// Copy is a representation of the C type gtk_paper_size_copy.
func (recv *PaperSize) Copy() *PaperSize {
	if copyPaperSizeInvoker == nil {
		copyPaperSizeInvoker = gi.StructFunctionInvokerNew("Gtk", "PaperSize", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyPaperSizeInvoker.Invoke(inArgs[:], nil)

	retGo := &PaperSize{native: ret.Pointer()}

	return retGo
}

var freePaperSizeInvoker *gi.Function

// Free is a representation of the C type gtk_paper_size_free.
func (recv *PaperSize) Free() {
	if freePaperSizeInvoker == nil {
		freePaperSizeInvoker = gi.StructFunctionInvokerNew("Gtk", "PaperSize", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freePaperSizeInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_paper_size_get_default_bottom_margin' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_default_left_margin' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_default_right_margin' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_default_top_margin' : parameter 'unit' of type 'Unit' not supported

var getDisplayNamePaperSizeInvoker *gi.Function

// GetDisplayName is a representation of the C type gtk_paper_size_get_display_name.
func (recv *PaperSize) GetDisplayName() string {
	if getDisplayNamePaperSizeInvoker == nil {
		getDisplayNamePaperSizeInvoker = gi.StructFunctionInvokerNew("Gtk", "PaperSize", "get_display_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDisplayNamePaperSizeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_get_height' : parameter 'unit' of type 'Unit' not supported

var getNamePaperSizeInvoker *gi.Function

// GetName is a representation of the C type gtk_paper_size_get_name.
func (recv *PaperSize) GetName() string {
	if getNamePaperSizeInvoker == nil {
		getNamePaperSizeInvoker = gi.StructFunctionInvokerNew("Gtk", "PaperSize", "get_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getNamePaperSizeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getPpdNamePaperSizeInvoker *gi.Function

// GetPpdName is a representation of the C type gtk_paper_size_get_ppd_name.
func (recv *PaperSize) GetPpdName() string {
	if getPpdNamePaperSizeInvoker == nil {
		getPpdNamePaperSizeInvoker = gi.StructFunctionInvokerNew("Gtk", "PaperSize", "get_ppd_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPpdNamePaperSizeInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_get_width' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_is_custom' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_paper_size_is_equal' : parameter 'size2' of type 'PaperSize' not supported

// UNSUPPORTED : C value 'gtk_paper_size_is_ipp' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_paper_size_set_size' : parameter 'width' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_paper_size_to_gvariant' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'gtk_paper_size_to_key_file' : parameter 'key_file' of type 'GLib.KeyFile' not supported

var PlacesSidebarClassStruct *gi.Struct
var PlacesSidebarClassStructOnce sync.Once

func PlacesSidebarClassStructSet() {
	PlacesSidebarClassStructOnce.Do(func() {
		PlacesSidebarClassStruct = gi.StructNew("Gtk", "PlacesSidebarClass")
	})
}

type PlacesSidebarClass struct {
	native uintptr
}

var PlugClassStruct *gi.Struct
var PlugClassStructOnce sync.Once

func PlugClassStructSet() {
	PlugClassStructOnce.Do(func() {
		PlugClassStruct = gi.StructNew("Gtk", "PlugClass")
	})
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

var PlugPrivateStruct *gi.Struct
var PlugPrivateStructOnce sync.Once

func PlugPrivateStructSet() {
	PlugPrivateStructOnce.Do(func() {
		PlugPrivateStruct = gi.StructNew("Gtk", "PlugPrivate")
	})
}

type PlugPrivate struct {
	native uintptr
}

var PopoverAccessibleClassStruct *gi.Struct
var PopoverAccessibleClassStructOnce sync.Once

func PopoverAccessibleClassStructSet() {
	PopoverAccessibleClassStructOnce.Do(func() {
		PopoverAccessibleClassStruct = gi.StructNew("Gtk", "PopoverAccessibleClass")
	})
}

type PopoverAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var PopoverClassStruct *gi.Struct
var PopoverClassStructOnce sync.Once

func PopoverClassStructSet() {
	PopoverClassStructOnce.Do(func() {
		PopoverClassStruct = gi.StructNew("Gtk", "PopoverClass")
	})
}

type PopoverClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'closed' : missing Type
}

var PopoverMenuClassStruct *gi.Struct
var PopoverMenuClassStructOnce sync.Once

func PopoverMenuClassStructSet() {
	PopoverMenuClassStructOnce.Do(func() {
		PopoverMenuClassStruct = gi.StructNew("Gtk", "PopoverMenuClass")
	})
}

type PopoverMenuClass struct {
	native      uintptr
	ParentClass *PopoverClass
}

var PopoverPrivateStruct *gi.Struct
var PopoverPrivateStructOnce sync.Once

func PopoverPrivateStructSet() {
	PopoverPrivateStructOnce.Do(func() {
		PopoverPrivateStruct = gi.StructNew("Gtk", "PopoverPrivate")
	})
}

type PopoverPrivate struct {
	native uintptr
}

var PrintOperationClassStruct *gi.Struct
var PrintOperationClassStructOnce sync.Once

func PrintOperationClassStructSet() {
	PrintOperationClassStructOnce.Do(func() {
		PrintOperationClassStruct = gi.StructNew("Gtk", "PrintOperationClass")
	})
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

var PrintOperationPreviewIfaceStruct *gi.Struct
var PrintOperationPreviewIfaceStructOnce sync.Once

func PrintOperationPreviewIfaceStructSet() {
	PrintOperationPreviewIfaceStructOnce.Do(func() {
		PrintOperationPreviewIfaceStruct = gi.StructNew("Gtk", "PrintOperationPreviewIface")
	})
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

var PrintOperationPrivateStruct *gi.Struct
var PrintOperationPrivateStructOnce sync.Once

func PrintOperationPrivateStructSet() {
	PrintOperationPrivateStructOnce.Do(func() {
		PrintOperationPrivateStruct = gi.StructNew("Gtk", "PrintOperationPrivate")
	})
}

type PrintOperationPrivate struct {
	native uintptr
}

var ProgressBarAccessibleClassStruct *gi.Struct
var ProgressBarAccessibleClassStructOnce sync.Once

func ProgressBarAccessibleClassStructSet() {
	ProgressBarAccessibleClassStructOnce.Do(func() {
		ProgressBarAccessibleClassStruct = gi.StructNew("Gtk", "ProgressBarAccessibleClass")
	})
}

type ProgressBarAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var ProgressBarAccessiblePrivateStruct *gi.Struct
var ProgressBarAccessiblePrivateStructOnce sync.Once

func ProgressBarAccessiblePrivateStructSet() {
	ProgressBarAccessiblePrivateStructOnce.Do(func() {
		ProgressBarAccessiblePrivateStruct = gi.StructNew("Gtk", "ProgressBarAccessiblePrivate")
	})
}

type ProgressBarAccessiblePrivate struct {
	native uintptr
}

var ProgressBarClassStruct *gi.Struct
var ProgressBarClassStructOnce sync.Once

func ProgressBarClassStructSet() {
	ProgressBarClassStructOnce.Do(func() {
		ProgressBarClassStruct = gi.StructNew("Gtk", "ProgressBarClass")
	})
}

type ProgressBarClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ProgressBarPrivateStruct *gi.Struct
var ProgressBarPrivateStructOnce sync.Once

func ProgressBarPrivateStructSet() {
	ProgressBarPrivateStructOnce.Do(func() {
		ProgressBarPrivateStruct = gi.StructNew("Gtk", "ProgressBarPrivate")
	})
}

type ProgressBarPrivate struct {
	native uintptr
}

var RadioActionClassStruct *gi.Struct
var RadioActionClassStructOnce sync.Once

func RadioActionClassStructSet() {
	RadioActionClassStructOnce.Do(func() {
		RadioActionClassStruct = gi.StructNew("Gtk", "RadioActionClass")
	})
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

var RadioActionEntryStruct *gi.Struct
var RadioActionEntryStructOnce sync.Once

func RadioActionEntryStructSet() {
	RadioActionEntryStructOnce.Do(func() {
		RadioActionEntryStruct = gi.StructNew("Gtk", "RadioActionEntry")
	})
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

var RadioActionPrivateStruct *gi.Struct
var RadioActionPrivateStructOnce sync.Once

func RadioActionPrivateStructSet() {
	RadioActionPrivateStructOnce.Do(func() {
		RadioActionPrivateStruct = gi.StructNew("Gtk", "RadioActionPrivate")
	})
}

type RadioActionPrivate struct {
	native uintptr
}

var RadioButtonAccessibleClassStruct *gi.Struct
var RadioButtonAccessibleClassStructOnce sync.Once

func RadioButtonAccessibleClassStructSet() {
	RadioButtonAccessibleClassStructOnce.Do(func() {
		RadioButtonAccessibleClassStruct = gi.StructNew("Gtk", "RadioButtonAccessibleClass")
	})
}

type RadioButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ToggleButtonAccessibleClass
}

var RadioButtonAccessiblePrivateStruct *gi.Struct
var RadioButtonAccessiblePrivateStructOnce sync.Once

func RadioButtonAccessiblePrivateStructSet() {
	RadioButtonAccessiblePrivateStructOnce.Do(func() {
		RadioButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "RadioButtonAccessiblePrivate")
	})
}

type RadioButtonAccessiblePrivate struct {
	native uintptr
}

var RadioButtonClassStruct *gi.Struct
var RadioButtonClassStructOnce sync.Once

func RadioButtonClassStructSet() {
	RadioButtonClassStructOnce.Do(func() {
		RadioButtonClassStruct = gi.StructNew("Gtk", "RadioButtonClass")
	})
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

var RadioButtonPrivateStruct *gi.Struct
var RadioButtonPrivateStructOnce sync.Once

func RadioButtonPrivateStructSet() {
	RadioButtonPrivateStructOnce.Do(func() {
		RadioButtonPrivateStruct = gi.StructNew("Gtk", "RadioButtonPrivate")
	})
}

type RadioButtonPrivate struct {
	native uintptr
}

var RadioMenuItemAccessibleClassStruct *gi.Struct
var RadioMenuItemAccessibleClassStructOnce sync.Once

func RadioMenuItemAccessibleClassStructSet() {
	RadioMenuItemAccessibleClassStructOnce.Do(func() {
		RadioMenuItemAccessibleClassStruct = gi.StructNew("Gtk", "RadioMenuItemAccessibleClass")
	})
}

type RadioMenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *CheckMenuItemAccessibleClass
}

var RadioMenuItemAccessiblePrivateStruct *gi.Struct
var RadioMenuItemAccessiblePrivateStructOnce sync.Once

func RadioMenuItemAccessiblePrivateStructSet() {
	RadioMenuItemAccessiblePrivateStructOnce.Do(func() {
		RadioMenuItemAccessiblePrivateStruct = gi.StructNew("Gtk", "RadioMenuItemAccessiblePrivate")
	})
}

type RadioMenuItemAccessiblePrivate struct {
	native uintptr
}

var RadioMenuItemClassStruct *gi.Struct
var RadioMenuItemClassStructOnce sync.Once

func RadioMenuItemClassStructSet() {
	RadioMenuItemClassStructOnce.Do(func() {
		RadioMenuItemClassStruct = gi.StructNew("Gtk", "RadioMenuItemClass")
	})
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

var RadioMenuItemPrivateStruct *gi.Struct
var RadioMenuItemPrivateStructOnce sync.Once

func RadioMenuItemPrivateStructSet() {
	RadioMenuItemPrivateStructOnce.Do(func() {
		RadioMenuItemPrivateStruct = gi.StructNew("Gtk", "RadioMenuItemPrivate")
	})
}

type RadioMenuItemPrivate struct {
	native uintptr
}

var RadioToolButtonClassStruct *gi.Struct
var RadioToolButtonClassStructOnce sync.Once

func RadioToolButtonClassStructSet() {
	RadioToolButtonClassStructOnce.Do(func() {
		RadioToolButtonClassStruct = gi.StructNew("Gtk", "RadioToolButtonClass")
	})
}

type RadioToolButtonClass struct {
	native      uintptr
	ParentClass *ToggleToolButtonClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var RangeAccessibleClassStruct *gi.Struct
var RangeAccessibleClassStructOnce sync.Once

func RangeAccessibleClassStructSet() {
	RangeAccessibleClassStructOnce.Do(func() {
		RangeAccessibleClassStruct = gi.StructNew("Gtk", "RangeAccessibleClass")
	})
}

type RangeAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var RangeAccessiblePrivateStruct *gi.Struct
var RangeAccessiblePrivateStructOnce sync.Once

func RangeAccessiblePrivateStructSet() {
	RangeAccessiblePrivateStructOnce.Do(func() {
		RangeAccessiblePrivateStruct = gi.StructNew("Gtk", "RangeAccessiblePrivate")
	})
}

type RangeAccessiblePrivate struct {
	native uintptr
}

var RangeClassStruct *gi.Struct
var RangeClassStructOnce sync.Once

func RangeClassStructSet() {
	RangeClassStructOnce.Do(func() {
		RangeClassStruct = gi.StructNew("Gtk", "RangeClass")
	})
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

var RangePrivateStruct *gi.Struct
var RangePrivateStructOnce sync.Once

func RangePrivateStructSet() {
	RangePrivateStructOnce.Do(func() {
		RangePrivateStruct = gi.StructNew("Gtk", "RangePrivate")
	})
}

type RangePrivate struct {
	native uintptr
}

var RcContextStruct *gi.Struct
var RcContextStructOnce sync.Once

func RcContextStructSet() {
	RcContextStructOnce.Do(func() {
		RcContextStruct = gi.StructNew("Gtk", "RcContext")
	})
}

type RcContext struct {
	native uintptr
}

var RcPropertyStruct *gi.Struct
var RcPropertyStructOnce sync.Once

func RcPropertyStructSet() {
	RcPropertyStructOnce.Do(func() {
		RcPropertyStruct = gi.StructNew("Gtk", "RcProperty")
	})
}

type RcProperty struct {
	native uintptr
	// UNSUPPORTED : C value 'type_name' : no Go type for 'GLib.Quark'
	// UNSUPPORTED : C value 'property_name' : no Go type for 'GLib.Quark'
	Origin string
	// UNSUPPORTED : C value 'value' : no Go type for 'GObject.Value'
}

var RcStyleClassStruct *gi.Struct
var RcStyleClassStructOnce sync.Once

func RcStyleClassStructSet() {
	RcStyleClassStructOnce.Do(func() {
		RcStyleClassStruct = gi.StructNew("Gtk", "RcStyleClass")
	})
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

var RecentActionClassStruct *gi.Struct
var RecentActionClassStructOnce sync.Once

func RecentActionClassStructSet() {
	RecentActionClassStructOnce.Do(func() {
		RecentActionClassStruct = gi.StructNew("Gtk", "RecentActionClass")
	})
}

type RecentActionClass struct {
	native      uintptr
	ParentClass *ActionClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var RecentActionPrivateStruct *gi.Struct
var RecentActionPrivateStructOnce sync.Once

func RecentActionPrivateStructSet() {
	RecentActionPrivateStructOnce.Do(func() {
		RecentActionPrivateStruct = gi.StructNew("Gtk", "RecentActionPrivate")
	})
}

type RecentActionPrivate struct {
	native uintptr
}

var RecentChooserDialogClassStruct *gi.Struct
var RecentChooserDialogClassStructOnce sync.Once

func RecentChooserDialogClassStructSet() {
	RecentChooserDialogClassStructOnce.Do(func() {
		RecentChooserDialogClassStruct = gi.StructNew("Gtk", "RecentChooserDialogClass")
	})
}

type RecentChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var RecentChooserDialogPrivateStruct *gi.Struct
var RecentChooserDialogPrivateStructOnce sync.Once

func RecentChooserDialogPrivateStructSet() {
	RecentChooserDialogPrivateStructOnce.Do(func() {
		RecentChooserDialogPrivateStruct = gi.StructNew("Gtk", "RecentChooserDialogPrivate")
	})
}

type RecentChooserDialogPrivate struct {
	native uintptr
}

var RecentChooserIfaceStruct *gi.Struct
var RecentChooserIfaceStructOnce sync.Once

func RecentChooserIfaceStructSet() {
	RecentChooserIfaceStructOnce.Do(func() {
		RecentChooserIfaceStruct = gi.StructNew("Gtk", "RecentChooserIface")
	})
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

var RecentChooserMenuClassStruct *gi.Struct
var RecentChooserMenuClassStructOnce sync.Once

func RecentChooserMenuClassStructSet() {
	RecentChooserMenuClassStructOnce.Do(func() {
		RecentChooserMenuClassStruct = gi.StructNew("Gtk", "RecentChooserMenuClass")
	})
}

type RecentChooserMenuClass struct {
	native      uintptr
	ParentClass *MenuClass
	// UNSUPPORTED : C value 'gtk_recent1' : missing Type
	// UNSUPPORTED : C value 'gtk_recent2' : missing Type
	// UNSUPPORTED : C value 'gtk_recent3' : missing Type
	// UNSUPPORTED : C value 'gtk_recent4' : missing Type
}

var RecentChooserMenuPrivateStruct *gi.Struct
var RecentChooserMenuPrivateStructOnce sync.Once

func RecentChooserMenuPrivateStructSet() {
	RecentChooserMenuPrivateStructOnce.Do(func() {
		RecentChooserMenuPrivateStruct = gi.StructNew("Gtk", "RecentChooserMenuPrivate")
	})
}

type RecentChooserMenuPrivate struct {
	native uintptr
}

var RecentChooserWidgetClassStruct *gi.Struct
var RecentChooserWidgetClassStructOnce sync.Once

func RecentChooserWidgetClassStructSet() {
	RecentChooserWidgetClassStructOnce.Do(func() {
		RecentChooserWidgetClassStruct = gi.StructNew("Gtk", "RecentChooserWidgetClass")
	})
}

type RecentChooserWidgetClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var RecentChooserWidgetPrivateStruct *gi.Struct
var RecentChooserWidgetPrivateStructOnce sync.Once

func RecentChooserWidgetPrivateStructSet() {
	RecentChooserWidgetPrivateStructOnce.Do(func() {
		RecentChooserWidgetPrivateStruct = gi.StructNew("Gtk", "RecentChooserWidgetPrivate")
	})
}

type RecentChooserWidgetPrivate struct {
	native uintptr
}

var RecentDataStruct *gi.Struct
var RecentDataStructOnce sync.Once

func RecentDataStructSet() {
	RecentDataStructOnce.Do(func() {
		RecentDataStruct = gi.StructNew("Gtk", "RecentData")
	})
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

var RecentFilterInfoStruct *gi.Struct
var RecentFilterInfoStructOnce sync.Once

func RecentFilterInfoStructSet() {
	RecentFilterInfoStructOnce.Do(func() {
		RecentFilterInfoStruct = gi.StructNew("Gtk", "RecentFilterInfo")
	})
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

var RecentInfoStruct *gi.Struct
var RecentInfoStructOnce sync.Once

func RecentInfoStructSet() {
	RecentInfoStructOnce.Do(func() {
		RecentInfoStruct = gi.StructNew("Gtk", "RecentInfo")
	})
}

type RecentInfo struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_recent_info_create_app_info' : return type 'Gio.AppInfo' not supported

// UNSUPPORTED : C value 'gtk_recent_info_exists' : return type 'gboolean' not supported

var getAddedRecentInfoInvoker *gi.Function

// GetAdded is a representation of the C type gtk_recent_info_get_added.
func (recv *RecentInfo) GetAdded() int64 {
	if getAddedRecentInfoInvoker == nil {
		getAddedRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_added")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getAddedRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var getAgeRecentInfoInvoker *gi.Function

// GetAge is a representation of the C type gtk_recent_info_get_age.
func (recv *RecentInfo) GetAge() int32 {
	if getAgeRecentInfoInvoker == nil {
		getAgeRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_age")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getAgeRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_get_application_info' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_recent_info_get_applications' : parameter 'length' of type 'gsize' not supported

var getDescriptionRecentInfoInvoker *gi.Function

// GetDescription is a representation of the C type gtk_recent_info_get_description.
func (recv *RecentInfo) GetDescription() string {
	if getDescriptionRecentInfoInvoker == nil {
		getDescriptionRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_description")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDescriptionRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getDisplayNameRecentInfoInvoker *gi.Function

// GetDisplayName is a representation of the C type gtk_recent_info_get_display_name.
func (recv *RecentInfo) GetDisplayName() string {
	if getDisplayNameRecentInfoInvoker == nil {
		getDisplayNameRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_display_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDisplayNameRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_get_gicon' : return type 'Gio.Icon' not supported

// UNSUPPORTED : C value 'gtk_recent_info_get_groups' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'gtk_recent_info_get_icon' : return type 'GdkPixbuf.Pixbuf' not supported

var getMimeTypeRecentInfoInvoker *gi.Function

// GetMimeType is a representation of the C type gtk_recent_info_get_mime_type.
func (recv *RecentInfo) GetMimeType() string {
	if getMimeTypeRecentInfoInvoker == nil {
		getMimeTypeRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_mime_type")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getMimeTypeRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getModifiedRecentInfoInvoker *gi.Function

// GetModified is a representation of the C type gtk_recent_info_get_modified.
func (recv *RecentInfo) GetModified() int64 {
	if getModifiedRecentInfoInvoker == nil {
		getModifiedRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_modified")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getModifiedRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_get_private_hint' : return type 'gboolean' not supported

var getShortNameRecentInfoInvoker *gi.Function

// GetShortName is a representation of the C type gtk_recent_info_get_short_name.
func (recv *RecentInfo) GetShortName() string {
	if getShortNameRecentInfoInvoker == nil {
		getShortNameRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_short_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getShortNameRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var getUriRecentInfoInvoker *gi.Function

// GetUri is a representation of the C type gtk_recent_info_get_uri.
func (recv *RecentInfo) GetUri() string {
	if getUriRecentInfoInvoker == nil {
		getUriRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_uri")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUriRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var getUriDisplayRecentInfoInvoker *gi.Function

// GetUriDisplay is a representation of the C type gtk_recent_info_get_uri_display.
func (recv *RecentInfo) GetUriDisplay() string {
	if getUriDisplayRecentInfoInvoker == nil {
		getUriDisplayRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_uri_display")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getUriDisplayRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var getVisitedRecentInfoInvoker *gi.Function

// GetVisited is a representation of the C type gtk_recent_info_get_visited.
func (recv *RecentInfo) GetVisited() int64 {
	if getVisitedRecentInfoInvoker == nil {
		getVisitedRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "get_visited")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getVisitedRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_has_application' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_recent_info_has_group' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_recent_info_is_local' : return type 'gboolean' not supported

var lastApplicationRecentInfoInvoker *gi.Function

// LastApplication is a representation of the C type gtk_recent_info_last_application.
func (recv *RecentInfo) LastApplication() string {
	if lastApplicationRecentInfoInvoker == nil {
		lastApplicationRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "last_application")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := lastApplicationRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_match' : parameter 'info_b' of type 'RecentInfo' not supported

var refRecentInfoInvoker *gi.Function

// Ref is a representation of the C type gtk_recent_info_ref.
func (recv *RecentInfo) Ref() *RecentInfo {
	if refRecentInfoInvoker == nil {
		refRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refRecentInfoInvoker.Invoke(inArgs[:], nil)

	retGo := &RecentInfo{native: ret.Pointer()}

	return retGo
}

var unrefRecentInfoInvoker *gi.Function

// Unref is a representation of the C type gtk_recent_info_unref.
func (recv *RecentInfo) Unref() {
	if unrefRecentInfoInvoker == nil {
		unrefRecentInfoInvoker = gi.StructFunctionInvokerNew("Gtk", "RecentInfo", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefRecentInfoInvoker.Invoke(inArgs[:], nil)

}

var RecentManagerClassStruct *gi.Struct
var RecentManagerClassStructOnce sync.Once

func RecentManagerClassStructSet() {
	RecentManagerClassStructOnce.Do(func() {
		RecentManagerClassStruct = gi.StructNew("Gtk", "RecentManagerClass")
	})
}

type RecentManagerClass struct {
	native uintptr
	// UNSUPPORTED : C value 'changed' : missing Type
	// UNSUPPORTED : C value '_gtk_recent1' : missing Type
	// UNSUPPORTED : C value '_gtk_recent2' : missing Type
	// UNSUPPORTED : C value '_gtk_recent3' : missing Type
	// UNSUPPORTED : C value '_gtk_recent4' : missing Type
}

var RecentManagerPrivateStruct *gi.Struct
var RecentManagerPrivateStructOnce sync.Once

func RecentManagerPrivateStructSet() {
	RecentManagerPrivateStructOnce.Do(func() {
		RecentManagerPrivateStruct = gi.StructNew("Gtk", "RecentManagerPrivate")
	})
}

type RecentManagerPrivate struct {
	native uintptr
}

var RendererCellAccessibleClassStruct *gi.Struct
var RendererCellAccessibleClassStructOnce sync.Once

func RendererCellAccessibleClassStructSet() {
	RendererCellAccessibleClassStructOnce.Do(func() {
		RendererCellAccessibleClassStruct = gi.StructNew("Gtk", "RendererCellAccessibleClass")
	})
}

type RendererCellAccessibleClass struct {
	native      uintptr
	ParentClass *CellAccessibleClass
}

var RendererCellAccessiblePrivateStruct *gi.Struct
var RendererCellAccessiblePrivateStructOnce sync.Once

func RendererCellAccessiblePrivateStructSet() {
	RendererCellAccessiblePrivateStructOnce.Do(func() {
		RendererCellAccessiblePrivateStruct = gi.StructNew("Gtk", "RendererCellAccessiblePrivate")
	})
}

type RendererCellAccessiblePrivate struct {
	native uintptr
}

var RequestedSizeStruct *gi.Struct
var RequestedSizeStructOnce sync.Once

func RequestedSizeStructSet() {
	RequestedSizeStructOnce.Do(func() {
		RequestedSizeStruct = gi.StructNew("Gtk", "RequestedSize")
	})
}

type RequestedSize struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	MinimumSize int32
	NaturalSize int32
}

var RequisitionStruct *gi.Struct
var RequisitionStructOnce sync.Once

func RequisitionStructSet() {
	RequisitionStructOnce.Do(func() {
		RequisitionStruct = gi.StructNew("Gtk", "Requisition")
	})
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

	retGo := &Requisition{native: ret.Pointer()}

	return retGo
}

var copyRequisitionInvoker *gi.Function

// Copy is a representation of the C type gtk_requisition_copy.
func (recv *Requisition) Copy() *Requisition {
	if copyRequisitionInvoker == nil {
		copyRequisitionInvoker = gi.StructFunctionInvokerNew("Gtk", "Requisition", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyRequisitionInvoker.Invoke(inArgs[:], nil)

	retGo := &Requisition{native: ret.Pointer()}

	return retGo
}

var freeRequisitionInvoker *gi.Function

// Free is a representation of the C type gtk_requisition_free.
func (recv *Requisition) Free() {
	if freeRequisitionInvoker == nil {
		freeRequisitionInvoker = gi.StructFunctionInvokerNew("Gtk", "Requisition", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeRequisitionInvoker.Invoke(inArgs[:], nil)

}

var RevealerClassStruct *gi.Struct
var RevealerClassStructOnce sync.Once

func RevealerClassStructSet() {
	RevealerClassStructOnce.Do(func() {
		RevealerClassStruct = gi.StructNew("Gtk", "RevealerClass")
	})
}

type RevealerClass struct {
	native      uintptr
	ParentClass *BinClass
}

var ScaleAccessibleClassStruct *gi.Struct
var ScaleAccessibleClassStructOnce sync.Once

func ScaleAccessibleClassStructSet() {
	ScaleAccessibleClassStructOnce.Do(func() {
		ScaleAccessibleClassStruct = gi.StructNew("Gtk", "ScaleAccessibleClass")
	})
}

type ScaleAccessibleClass struct {
	native      uintptr
	ParentClass *RangeAccessibleClass
}

var ScaleAccessiblePrivateStruct *gi.Struct
var ScaleAccessiblePrivateStructOnce sync.Once

func ScaleAccessiblePrivateStructSet() {
	ScaleAccessiblePrivateStructOnce.Do(func() {
		ScaleAccessiblePrivateStruct = gi.StructNew("Gtk", "ScaleAccessiblePrivate")
	})
}

type ScaleAccessiblePrivate struct {
	native uintptr
}

var ScaleButtonAccessibleClassStruct *gi.Struct
var ScaleButtonAccessibleClassStructOnce sync.Once

func ScaleButtonAccessibleClassStructSet() {
	ScaleButtonAccessibleClassStructOnce.Do(func() {
		ScaleButtonAccessibleClassStruct = gi.StructNew("Gtk", "ScaleButtonAccessibleClass")
	})
}

type ScaleButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var ScaleButtonAccessiblePrivateStruct *gi.Struct
var ScaleButtonAccessiblePrivateStructOnce sync.Once

func ScaleButtonAccessiblePrivateStructSet() {
	ScaleButtonAccessiblePrivateStructOnce.Do(func() {
		ScaleButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "ScaleButtonAccessiblePrivate")
	})
}

type ScaleButtonAccessiblePrivate struct {
	native uintptr
}

var ScaleButtonClassStruct *gi.Struct
var ScaleButtonClassStructOnce sync.Once

func ScaleButtonClassStructSet() {
	ScaleButtonClassStructOnce.Do(func() {
		ScaleButtonClassStruct = gi.StructNew("Gtk", "ScaleButtonClass")
	})
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

var ScaleButtonPrivateStruct *gi.Struct
var ScaleButtonPrivateStructOnce sync.Once

func ScaleButtonPrivateStructSet() {
	ScaleButtonPrivateStructOnce.Do(func() {
		ScaleButtonPrivateStruct = gi.StructNew("Gtk", "ScaleButtonPrivate")
	})
}

type ScaleButtonPrivate struct {
	native uintptr
}

var ScaleClassStruct *gi.Struct
var ScaleClassStructOnce sync.Once

func ScaleClassStructSet() {
	ScaleClassStructOnce.Do(func() {
		ScaleClassStruct = gi.StructNew("Gtk", "ScaleClass")
	})
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

var ScalePrivateStruct *gi.Struct
var ScalePrivateStructOnce sync.Once

func ScalePrivateStructSet() {
	ScalePrivateStructOnce.Do(func() {
		ScalePrivateStruct = gi.StructNew("Gtk", "ScalePrivate")
	})
}

type ScalePrivate struct {
	native uintptr
}

var ScrollableInterfaceStruct *gi.Struct
var ScrollableInterfaceStructOnce sync.Once

func ScrollableInterfaceStructSet() {
	ScrollableInterfaceStructOnce.Do(func() {
		ScrollableInterfaceStruct = gi.StructNew("Gtk", "ScrollableInterface")
	})
}

type ScrollableInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_border' : missing Type
}

var ScrollbarClassStruct *gi.Struct
var ScrollbarClassStructOnce sync.Once

func ScrollbarClassStructSet() {
	ScrollbarClassStructOnce.Do(func() {
		ScrollbarClassStruct = gi.StructNew("Gtk", "ScrollbarClass")
	})
}

type ScrollbarClass struct {
	native      uintptr
	ParentClass *RangeClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ScrolledWindowAccessibleClassStruct *gi.Struct
var ScrolledWindowAccessibleClassStructOnce sync.Once

func ScrolledWindowAccessibleClassStructSet() {
	ScrolledWindowAccessibleClassStructOnce.Do(func() {
		ScrolledWindowAccessibleClassStruct = gi.StructNew("Gtk", "ScrolledWindowAccessibleClass")
	})
}

type ScrolledWindowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var ScrolledWindowAccessiblePrivateStruct *gi.Struct
var ScrolledWindowAccessiblePrivateStructOnce sync.Once

func ScrolledWindowAccessiblePrivateStructSet() {
	ScrolledWindowAccessiblePrivateStructOnce.Do(func() {
		ScrolledWindowAccessiblePrivateStruct = gi.StructNew("Gtk", "ScrolledWindowAccessiblePrivate")
	})
}

type ScrolledWindowAccessiblePrivate struct {
	native uintptr
}

var ScrolledWindowClassStruct *gi.Struct
var ScrolledWindowClassStructOnce sync.Once

func ScrolledWindowClassStructSet() {
	ScrolledWindowClassStructOnce.Do(func() {
		ScrolledWindowClassStruct = gi.StructNew("Gtk", "ScrolledWindowClass")
	})
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

var ScrolledWindowPrivateStruct *gi.Struct
var ScrolledWindowPrivateStructOnce sync.Once

func ScrolledWindowPrivateStructSet() {
	ScrolledWindowPrivateStructOnce.Do(func() {
		ScrolledWindowPrivateStruct = gi.StructNew("Gtk", "ScrolledWindowPrivate")
	})
}

type ScrolledWindowPrivate struct {
	native uintptr
}

var SearchBarClassStruct *gi.Struct
var SearchBarClassStructOnce sync.Once

func SearchBarClassStructSet() {
	SearchBarClassStructOnce.Do(func() {
		SearchBarClassStruct = gi.StructNew("Gtk", "SearchBarClass")
	})
}

type SearchBarClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var SearchEntryClassStruct *gi.Struct
var SearchEntryClassStructOnce sync.Once

func SearchEntryClassStructSet() {
	SearchEntryClassStructOnce.Do(func() {
		SearchEntryClassStruct = gi.StructNew("Gtk", "SearchEntryClass")
	})
}

type SearchEntryClass struct {
	native      uintptr
	ParentClass *EntryClass
	// UNSUPPORTED : C value 'search_changed' : missing Type
	// UNSUPPORTED : C value 'next_match' : missing Type
	// UNSUPPORTED : C value 'previous_match' : missing Type
	// UNSUPPORTED : C value 'stop_search' : missing Type
}

var SelectionDataStruct *gi.Struct
var SelectionDataStructOnce sync.Once

func SelectionDataStructSet() {
	SelectionDataStructOnce.Do(func() {
		SelectionDataStruct = gi.StructNew("Gtk", "SelectionData")
	})
}

type SelectionData struct {
	native uintptr
}

var copySelectionDataInvoker *gi.Function

// Copy is a representation of the C type gtk_selection_data_copy.
func (recv *SelectionData) Copy() *SelectionData {
	if copySelectionDataInvoker == nil {
		copySelectionDataInvoker = gi.StructFunctionInvokerNew("Gtk", "SelectionData", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copySelectionDataInvoker.Invoke(inArgs[:], nil)

	retGo := &SelectionData{native: ret.Pointer()}

	return retGo
}

var freeSelectionDataInvoker *gi.Function

// Free is a representation of the C type gtk_selection_data_free.
func (recv *SelectionData) Free() {
	if freeSelectionDataInvoker == nil {
		freeSelectionDataInvoker = gi.StructFunctionInvokerNew("Gtk", "SelectionData", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeSelectionDataInvoker.Invoke(inArgs[:], nil)

}

var getDataSelectionDataInvoker *gi.Function

// GetData is a representation of the C type gtk_selection_data_get_data.
func (recv *SelectionData) GetData() {
	if getDataSelectionDataInvoker == nil {
		getDataSelectionDataInvoker = gi.StructFunctionInvokerNew("Gtk", "SelectionData", "get_data")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	getDataSelectionDataInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_selection_data_get_data_type' : return type 'Gdk.Atom' not supported

var getDataWithLengthSelectionDataInvoker *gi.Function

// GetDataWithLength is a representation of the C type gtk_selection_data_get_data_with_length.
func (recv *SelectionData) GetDataWithLength() int32 {
	if getDataWithLengthSelectionDataInvoker == nil {
		getDataWithLengthSelectionDataInvoker = gi.StructFunctionInvokerNew("Gtk", "SelectionData", "get_data_with_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	getDataWithLengthSelectionDataInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'gtk_selection_data_get_display' : return type 'Gdk.Display' not supported

var getFormatSelectionDataInvoker *gi.Function

// GetFormat is a representation of the C type gtk_selection_data_get_format.
func (recv *SelectionData) GetFormat() int32 {
	if getFormatSelectionDataInvoker == nil {
		getFormatSelectionDataInvoker = gi.StructFunctionInvokerNew("Gtk", "SelectionData", "get_format")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getFormatSelectionDataInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getLengthSelectionDataInvoker *gi.Function

// GetLength is a representation of the C type gtk_selection_data_get_length.
func (recv *SelectionData) GetLength() int32 {
	if getLengthSelectionDataInvoker == nil {
		getLengthSelectionDataInvoker = gi.StructFunctionInvokerNew("Gtk", "SelectionData", "get_length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLengthSelectionDataInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_selection_data_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_selection_data_get_selection' : return type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_data_get_target' : return type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_data_get_targets' : parameter 'targets' has no type

var getTextSelectionDataInvoker *gi.Function

// GetText is a representation of the C type gtk_selection_data_get_text.
func (recv *SelectionData) GetText() string {
	if getTextSelectionDataInvoker == nil {
		getTextSelectionDataInvoker = gi.StructFunctionInvokerNew("Gtk", "SelectionData", "get_text")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getTextSelectionDataInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var getUrisSelectionDataInvoker *gi.Function

// GetUris is a representation of the C type gtk_selection_data_get_uris.
func (recv *SelectionData) GetUris() {
	if getUrisSelectionDataInvoker == nil {
		getUrisSelectionDataInvoker = gi.StructFunctionInvokerNew("Gtk", "SelectionData", "get_uris")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	getUrisSelectionDataInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_selection_data_set' : parameter 'type' of type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_data_set_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_selection_data_set_text' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_selection_data_set_uris' : parameter 'uris' has no type

// UNSUPPORTED : C value 'gtk_selection_data_targets_include_image' : parameter 'writable' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_selection_data_targets_include_rich_text' : parameter 'buffer' of type 'TextBuffer' not supported

// UNSUPPORTED : C value 'gtk_selection_data_targets_include_text' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_selection_data_targets_include_uri' : return type 'gboolean' not supported

var SeparatorClassStruct *gi.Struct
var SeparatorClassStructOnce sync.Once

func SeparatorClassStructSet() {
	SeparatorClassStructOnce.Do(func() {
		SeparatorClassStruct = gi.StructNew("Gtk", "SeparatorClass")
	})
}

type SeparatorClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var SeparatorMenuItemClassStruct *gi.Struct
var SeparatorMenuItemClassStructOnce sync.Once

func SeparatorMenuItemClassStructSet() {
	SeparatorMenuItemClassStructOnce.Do(func() {
		SeparatorMenuItemClassStruct = gi.StructNew("Gtk", "SeparatorMenuItemClass")
	})
}

type SeparatorMenuItemClass struct {
	native      uintptr
	ParentClass *MenuItemClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var SeparatorPrivateStruct *gi.Struct
var SeparatorPrivateStructOnce sync.Once

func SeparatorPrivateStructSet() {
	SeparatorPrivateStructOnce.Do(func() {
		SeparatorPrivateStruct = gi.StructNew("Gtk", "SeparatorPrivate")
	})
}

type SeparatorPrivate struct {
	native uintptr
}

var SeparatorToolItemClassStruct *gi.Struct
var SeparatorToolItemClassStructOnce sync.Once

func SeparatorToolItemClassStructSet() {
	SeparatorToolItemClassStructOnce.Do(func() {
		SeparatorToolItemClassStruct = gi.StructNew("Gtk", "SeparatorToolItemClass")
	})
}

type SeparatorToolItemClass struct {
	native      uintptr
	ParentClass *ToolItemClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var SeparatorToolItemPrivateStruct *gi.Struct
var SeparatorToolItemPrivateStructOnce sync.Once

func SeparatorToolItemPrivateStructSet() {
	SeparatorToolItemPrivateStructOnce.Do(func() {
		SeparatorToolItemPrivateStruct = gi.StructNew("Gtk", "SeparatorToolItemPrivate")
	})
}

type SeparatorToolItemPrivate struct {
	native uintptr
}

var SettingsClassStruct *gi.Struct
var SettingsClassStructOnce sync.Once

func SettingsClassStructSet() {
	SettingsClassStructOnce.Do(func() {
		SettingsClassStruct = gi.StructNew("Gtk", "SettingsClass")
	})
}

type SettingsClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var SettingsPrivateStruct *gi.Struct
var SettingsPrivateStructOnce sync.Once

func SettingsPrivateStructSet() {
	SettingsPrivateStructOnce.Do(func() {
		SettingsPrivateStruct = gi.StructNew("Gtk", "SettingsPrivate")
	})
}

type SettingsPrivate struct {
	native uintptr
}

var SettingsValueStruct *gi.Struct
var SettingsValueStructOnce sync.Once

func SettingsValueStructSet() {
	SettingsValueStructOnce.Do(func() {
		SettingsValueStruct = gi.StructNew("Gtk", "SettingsValue")
	})
}

type SettingsValue struct {
	native uintptr
	Origin string
	// UNSUPPORTED : C value 'value' : no Go type for 'GObject.Value'
}

var ShortcutLabelClassStruct *gi.Struct
var ShortcutLabelClassStructOnce sync.Once

func ShortcutLabelClassStructSet() {
	ShortcutLabelClassStructOnce.Do(func() {
		ShortcutLabelClassStruct = gi.StructNew("Gtk", "ShortcutLabelClass")
	})
}

type ShortcutLabelClass struct {
	native uintptr
}

var ShortcutsGroupClassStruct *gi.Struct
var ShortcutsGroupClassStructOnce sync.Once

func ShortcutsGroupClassStructSet() {
	ShortcutsGroupClassStructOnce.Do(func() {
		ShortcutsGroupClassStruct = gi.StructNew("Gtk", "ShortcutsGroupClass")
	})
}

type ShortcutsGroupClass struct {
	native uintptr
}

var ShortcutsSectionClassStruct *gi.Struct
var ShortcutsSectionClassStructOnce sync.Once

func ShortcutsSectionClassStructSet() {
	ShortcutsSectionClassStructOnce.Do(func() {
		ShortcutsSectionClassStruct = gi.StructNew("Gtk", "ShortcutsSectionClass")
	})
}

type ShortcutsSectionClass struct {
	native uintptr
}

var ShortcutsShortcutClassStruct *gi.Struct
var ShortcutsShortcutClassStructOnce sync.Once

func ShortcutsShortcutClassStructSet() {
	ShortcutsShortcutClassStructOnce.Do(func() {
		ShortcutsShortcutClassStruct = gi.StructNew("Gtk", "ShortcutsShortcutClass")
	})
}

type ShortcutsShortcutClass struct {
	native uintptr
}

var ShortcutsWindowClassStruct *gi.Struct
var ShortcutsWindowClassStructOnce sync.Once

func ShortcutsWindowClassStructSet() {
	ShortcutsWindowClassStructOnce.Do(func() {
		ShortcutsWindowClassStruct = gi.StructNew("Gtk", "ShortcutsWindowClass")
	})
}

type ShortcutsWindowClass struct {
	native      uintptr
	ParentClass *WindowClass
	// UNSUPPORTED : C value 'close' : missing Type
	// UNSUPPORTED : C value 'search' : missing Type
}

var SizeGroupClassStruct *gi.Struct
var SizeGroupClassStructOnce sync.Once

func SizeGroupClassStructSet() {
	SizeGroupClassStructOnce.Do(func() {
		SizeGroupClassStruct = gi.StructNew("Gtk", "SizeGroupClass")
	})
}

type SizeGroupClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var SizeGroupPrivateStruct *gi.Struct
var SizeGroupPrivateStructOnce sync.Once

func SizeGroupPrivateStructSet() {
	SizeGroupPrivateStructOnce.Do(func() {
		SizeGroupPrivateStruct = gi.StructNew("Gtk", "SizeGroupPrivate")
	})
}

type SizeGroupPrivate struct {
	native uintptr
}

var SocketClassStruct *gi.Struct
var SocketClassStructOnce sync.Once

func SocketClassStructSet() {
	SocketClassStructOnce.Do(func() {
		SocketClassStruct = gi.StructNew("Gtk", "SocketClass")
	})
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

var SocketPrivateStruct *gi.Struct
var SocketPrivateStructOnce sync.Once

func SocketPrivateStructSet() {
	SocketPrivateStructOnce.Do(func() {
		SocketPrivateStruct = gi.StructNew("Gtk", "SocketPrivate")
	})
}

type SocketPrivate struct {
	native uintptr
}

var SpinButtonAccessibleClassStruct *gi.Struct
var SpinButtonAccessibleClassStructOnce sync.Once

func SpinButtonAccessibleClassStructSet() {
	SpinButtonAccessibleClassStructOnce.Do(func() {
		SpinButtonAccessibleClassStruct = gi.StructNew("Gtk", "SpinButtonAccessibleClass")
	})
}

type SpinButtonAccessibleClass struct {
	native      uintptr
	ParentClass *EntryAccessibleClass
}

var SpinButtonAccessiblePrivateStruct *gi.Struct
var SpinButtonAccessiblePrivateStructOnce sync.Once

func SpinButtonAccessiblePrivateStructSet() {
	SpinButtonAccessiblePrivateStructOnce.Do(func() {
		SpinButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "SpinButtonAccessiblePrivate")
	})
}

type SpinButtonAccessiblePrivate struct {
	native uintptr
}

var SpinButtonClassStruct *gi.Struct
var SpinButtonClassStructOnce sync.Once

func SpinButtonClassStructSet() {
	SpinButtonClassStructOnce.Do(func() {
		SpinButtonClassStruct = gi.StructNew("Gtk", "SpinButtonClass")
	})
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

var SpinButtonPrivateStruct *gi.Struct
var SpinButtonPrivateStructOnce sync.Once

func SpinButtonPrivateStructSet() {
	SpinButtonPrivateStructOnce.Do(func() {
		SpinButtonPrivateStruct = gi.StructNew("Gtk", "SpinButtonPrivate")
	})
}

type SpinButtonPrivate struct {
	native uintptr
}

var SpinnerAccessibleClassStruct *gi.Struct
var SpinnerAccessibleClassStructOnce sync.Once

func SpinnerAccessibleClassStructSet() {
	SpinnerAccessibleClassStructOnce.Do(func() {
		SpinnerAccessibleClassStruct = gi.StructNew("Gtk", "SpinnerAccessibleClass")
	})
}

type SpinnerAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var SpinnerAccessiblePrivateStruct *gi.Struct
var SpinnerAccessiblePrivateStructOnce sync.Once

func SpinnerAccessiblePrivateStructSet() {
	SpinnerAccessiblePrivateStructOnce.Do(func() {
		SpinnerAccessiblePrivateStruct = gi.StructNew("Gtk", "SpinnerAccessiblePrivate")
	})
}

type SpinnerAccessiblePrivate struct {
	native uintptr
}

var SpinnerClassStruct *gi.Struct
var SpinnerClassStructOnce sync.Once

func SpinnerClassStructSet() {
	SpinnerClassStructOnce.Do(func() {
		SpinnerClassStruct = gi.StructNew("Gtk", "SpinnerClass")
	})
}

type SpinnerClass struct {
	native      uintptr
	ParentClass *WidgetClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var SpinnerPrivateStruct *gi.Struct
var SpinnerPrivateStructOnce sync.Once

func SpinnerPrivateStructSet() {
	SpinnerPrivateStructOnce.Do(func() {
		SpinnerPrivateStruct = gi.StructNew("Gtk", "SpinnerPrivate")
	})
}

type SpinnerPrivate struct {
	native uintptr
}

var StackAccessibleClassStruct *gi.Struct
var StackAccessibleClassStructOnce sync.Once

func StackAccessibleClassStructSet() {
	StackAccessibleClassStructOnce.Do(func() {
		StackAccessibleClassStruct = gi.StructNew("Gtk", "StackAccessibleClass")
	})
}

type StackAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var StackClassStruct *gi.Struct
var StackClassStructOnce sync.Once

func StackClassStructSet() {
	StackClassStructOnce.Do(func() {
		StackClassStruct = gi.StructNew("Gtk", "StackClass")
	})
}

type StackClass struct {
	native      uintptr
	ParentClass *ContainerClass
}

var StackSidebarClassStruct *gi.Struct
var StackSidebarClassStructOnce sync.Once

func StackSidebarClassStructSet() {
	StackSidebarClassStructOnce.Do(func() {
		StackSidebarClassStruct = gi.StructNew("Gtk", "StackSidebarClass")
	})
}

type StackSidebarClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var StackSidebarPrivateStruct *gi.Struct
var StackSidebarPrivateStructOnce sync.Once

func StackSidebarPrivateStructSet() {
	StackSidebarPrivateStructOnce.Do(func() {
		StackSidebarPrivateStruct = gi.StructNew("Gtk", "StackSidebarPrivate")
	})
}

type StackSidebarPrivate struct {
	native uintptr
}

var StackSwitcherClassStruct *gi.Struct
var StackSwitcherClassStructOnce sync.Once

func StackSwitcherClassStructSet() {
	StackSwitcherClassStructOnce.Do(func() {
		StackSwitcherClassStruct = gi.StructNew("Gtk", "StackSwitcherClass")
	})
}

type StackSwitcherClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var StatusIconClassStruct *gi.Struct
var StatusIconClassStructOnce sync.Once

func StatusIconClassStructSet() {
	StatusIconClassStructOnce.Do(func() {
		StatusIconClassStruct = gi.StructNew("Gtk", "StatusIconClass")
	})
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

var StatusIconPrivateStruct *gi.Struct
var StatusIconPrivateStructOnce sync.Once

func StatusIconPrivateStructSet() {
	StatusIconPrivateStructOnce.Do(func() {
		StatusIconPrivateStruct = gi.StructNew("Gtk", "StatusIconPrivate")
	})
}

type StatusIconPrivate struct {
	native uintptr
}

var StatusbarAccessibleClassStruct *gi.Struct
var StatusbarAccessibleClassStructOnce sync.Once

func StatusbarAccessibleClassStructSet() {
	StatusbarAccessibleClassStructOnce.Do(func() {
		StatusbarAccessibleClassStruct = gi.StructNew("Gtk", "StatusbarAccessibleClass")
	})
}

type StatusbarAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var StatusbarAccessiblePrivateStruct *gi.Struct
var StatusbarAccessiblePrivateStructOnce sync.Once

func StatusbarAccessiblePrivateStructSet() {
	StatusbarAccessiblePrivateStructOnce.Do(func() {
		StatusbarAccessiblePrivateStruct = gi.StructNew("Gtk", "StatusbarAccessiblePrivate")
	})
}

type StatusbarAccessiblePrivate struct {
	native uintptr
}

var StatusbarClassStruct *gi.Struct
var StatusbarClassStructOnce sync.Once

func StatusbarClassStructSet() {
	StatusbarClassStructOnce.Do(func() {
		StatusbarClassStruct = gi.StructNew("Gtk", "StatusbarClass")
	})
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

var StatusbarPrivateStruct *gi.Struct
var StatusbarPrivateStructOnce sync.Once

func StatusbarPrivateStructSet() {
	StatusbarPrivateStructOnce.Do(func() {
		StatusbarPrivateStruct = gi.StructNew("Gtk", "StatusbarPrivate")
	})
}

type StatusbarPrivate struct {
	native uintptr
}

var StockItemStruct *gi.Struct
var StockItemStructOnce sync.Once

func StockItemStructSet() {
	StockItemStructOnce.Do(func() {
		StockItemStruct = gi.StructNew("Gtk", "StockItem")
	})
}

type StockItem struct {
	native  uintptr
	StockId string
	Label   string
	// UNSUPPORTED : C value 'modifier' : no Go type for 'Gdk.ModifierType'
	Keyval            uint32
	TranslationDomain string
}

var copyStockItemInvoker *gi.Function

// Copy is a representation of the C type gtk_stock_item_copy.
func (recv *StockItem) Copy() *StockItem {
	if copyStockItemInvoker == nil {
		copyStockItemInvoker = gi.StructFunctionInvokerNew("Gtk", "StockItem", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyStockItemInvoker.Invoke(inArgs[:], nil)

	retGo := &StockItem{native: ret.Pointer()}

	return retGo
}

var freeStockItemInvoker *gi.Function

// Free is a representation of the C type gtk_stock_item_free.
func (recv *StockItem) Free() {
	if freeStockItemInvoker == nil {
		freeStockItemInvoker = gi.StructFunctionInvokerNew("Gtk", "StockItem", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeStockItemInvoker.Invoke(inArgs[:], nil)

}

var StyleClassStruct *gi.Struct
var StyleClassStructOnce sync.Once

func StyleClassStructSet() {
	StyleClassStructOnce.Do(func() {
		StyleClassStruct = gi.StructNew("Gtk", "StyleClass")
	})
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

var StyleContextClassStruct *gi.Struct
var StyleContextClassStructOnce sync.Once

func StyleContextClassStructSet() {
	StyleContextClassStructOnce.Do(func() {
		StyleContextClassStruct = gi.StructNew("Gtk", "StyleContextClass")
	})
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

var StyleContextPrivateStruct *gi.Struct
var StyleContextPrivateStructOnce sync.Once

func StyleContextPrivateStructSet() {
	StyleContextPrivateStructOnce.Do(func() {
		StyleContextPrivateStruct = gi.StructNew("Gtk", "StyleContextPrivate")
	})
}

type StyleContextPrivate struct {
	native uintptr
}

var StylePropertiesClassStruct *gi.Struct
var StylePropertiesClassStructOnce sync.Once

func StylePropertiesClassStructSet() {
	StylePropertiesClassStructOnce.Do(func() {
		StylePropertiesClassStruct = gi.StructNew("Gtk", "StylePropertiesClass")
	})
}

type StylePropertiesClass struct {
	native uintptr
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var StylePropertiesPrivateStruct *gi.Struct
var StylePropertiesPrivateStructOnce sync.Once

func StylePropertiesPrivateStructSet() {
	StylePropertiesPrivateStructOnce.Do(func() {
		StylePropertiesPrivateStruct = gi.StructNew("Gtk", "StylePropertiesPrivate")
	})
}

type StylePropertiesPrivate struct {
	native uintptr
}

var StyleProviderIfaceStruct *gi.Struct
var StyleProviderIfaceStructOnce sync.Once

func StyleProviderIfaceStructSet() {
	StyleProviderIfaceStructOnce.Do(func() {
		StyleProviderIfaceStruct = gi.StructNew("Gtk", "StyleProviderIface")
	})
}

type StyleProviderIface struct {
	native uintptr
	// UNSUPPORTED : C value 'get_style' : missing Type
	// UNSUPPORTED : C value 'get_style_property' : missing Type
	// UNSUPPORTED : C value 'get_icon_factory' : missing Type
}

var SwitchAccessibleClassStruct *gi.Struct
var SwitchAccessibleClassStructOnce sync.Once

func SwitchAccessibleClassStructSet() {
	SwitchAccessibleClassStructOnce.Do(func() {
		SwitchAccessibleClassStruct = gi.StructNew("Gtk", "SwitchAccessibleClass")
	})
}

type SwitchAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var SwitchAccessiblePrivateStruct *gi.Struct
var SwitchAccessiblePrivateStructOnce sync.Once

func SwitchAccessiblePrivateStructSet() {
	SwitchAccessiblePrivateStructOnce.Do(func() {
		SwitchAccessiblePrivateStruct = gi.StructNew("Gtk", "SwitchAccessiblePrivate")
	})
}

type SwitchAccessiblePrivate struct {
	native uintptr
}

var SwitchClassStruct *gi.Struct
var SwitchClassStructOnce sync.Once

func SwitchClassStructSet() {
	SwitchClassStructOnce.Do(func() {
		SwitchClassStruct = gi.StructNew("Gtk", "SwitchClass")
	})
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

var SwitchPrivateStruct *gi.Struct
var SwitchPrivateStructOnce sync.Once

func SwitchPrivateStructSet() {
	SwitchPrivateStructOnce.Do(func() {
		SwitchPrivateStruct = gi.StructNew("Gtk", "SwitchPrivate")
	})
}

type SwitchPrivate struct {
	native uintptr
}

var SymbolicColorStruct *gi.Struct
var SymbolicColorStructOnce sync.Once

func SymbolicColorStructSet() {
	SymbolicColorStructOnce.Do(func() {
		SymbolicColorStruct = gi.StructNew("Gtk", "SymbolicColor")
	})
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

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
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

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

var refSymbolicColorInvoker *gi.Function

// Ref is a representation of the C type gtk_symbolic_color_ref.
func (recv *SymbolicColor) Ref() *SymbolicColor {
	if refSymbolicColorInvoker == nil {
		refSymbolicColorInvoker = gi.StructFunctionInvokerNew("Gtk", "SymbolicColor", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refSymbolicColorInvoker.Invoke(inArgs[:], nil)

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_symbolic_color_resolve' : parameter 'props' of type 'StyleProperties' not supported

var toStringSymbolicColorInvoker *gi.Function

// ToString is a representation of the C type gtk_symbolic_color_to_string.
func (recv *SymbolicColor) ToString() string {
	if toStringSymbolicColorInvoker == nil {
		toStringSymbolicColorInvoker = gi.StructFunctionInvokerNew("Gtk", "SymbolicColor", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringSymbolicColorInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var unrefSymbolicColorInvoker *gi.Function

// Unref is a representation of the C type gtk_symbolic_color_unref.
func (recv *SymbolicColor) Unref() {
	if unrefSymbolicColorInvoker == nil {
		unrefSymbolicColorInvoker = gi.StructFunctionInvokerNew("Gtk", "SymbolicColor", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefSymbolicColorInvoker.Invoke(inArgs[:], nil)

}

var TableChildStruct *gi.Struct
var TableChildStructOnce sync.Once

func TableChildStructSet() {
	TableChildStructOnce.Do(func() {
		TableChildStruct = gi.StructNew("Gtk", "TableChild")
	})
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

var TableClassStruct *gi.Struct
var TableClassStructOnce sync.Once

func TableClassStructSet() {
	TableClassStructOnce.Do(func() {
		TableClassStruct = gi.StructNew("Gtk", "TableClass")
	})
}

type TableClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var TablePrivateStruct *gi.Struct
var TablePrivateStructOnce sync.Once

func TablePrivateStructSet() {
	TablePrivateStructOnce.Do(func() {
		TablePrivateStruct = gi.StructNew("Gtk", "TablePrivate")
	})
}

type TablePrivate struct {
	native uintptr
}

var TableRowColStruct *gi.Struct
var TableRowColStructOnce sync.Once

func TableRowColStructSet() {
	TableRowColStructOnce.Do(func() {
		TableRowColStruct = gi.StructNew("Gtk", "TableRowCol")
	})
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

var TargetEntryStruct *gi.Struct
var TargetEntryStructOnce sync.Once

func TargetEntryStructSet() {
	TargetEntryStructOnce.Do(func() {
		TargetEntryStruct = gi.StructNew("Gtk", "TargetEntry")
	})
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

	retGo := &TargetEntry{native: ret.Pointer()}

	return retGo
}

var copyTargetEntryInvoker *gi.Function

// Copy is a representation of the C type gtk_target_entry_copy.
func (recv *TargetEntry) Copy() *TargetEntry {
	if copyTargetEntryInvoker == nil {
		copyTargetEntryInvoker = gi.StructFunctionInvokerNew("Gtk", "TargetEntry", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyTargetEntryInvoker.Invoke(inArgs[:], nil)

	retGo := &TargetEntry{native: ret.Pointer()}

	return retGo
}

var freeTargetEntryInvoker *gi.Function

// Free is a representation of the C type gtk_target_entry_free.
func (recv *TargetEntry) Free() {
	if freeTargetEntryInvoker == nil {
		freeTargetEntryInvoker = gi.StructFunctionInvokerNew("Gtk", "TargetEntry", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeTargetEntryInvoker.Invoke(inArgs[:], nil)

}

var TargetListStruct *gi.Struct
var TargetListStructOnce sync.Once

func TargetListStructSet() {
	TargetListStructOnce.Do(func() {
		TargetListStruct = gi.StructNew("Gtk", "TargetList")
	})
}

type TargetList struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_target_list_new' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_target_list_add' : parameter 'target' of type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_target_list_add_image_targets' : parameter 'writable' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_target_list_add_rich_text_targets' : parameter 'deserializable' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_target_list_add_table' : parameter 'targets' has no type

var addTextTargetsTargetListInvoker *gi.Function

// AddTextTargets is a representation of the C type gtk_target_list_add_text_targets.
func (recv *TargetList) AddTextTargets(info uint32) {
	if addTextTargetsTargetListInvoker == nil {
		addTextTargetsTargetListInvoker = gi.StructFunctionInvokerNew("Gtk", "TargetList", "add_text_targets")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(info)

	addTextTargetsTargetListInvoker.Invoke(inArgs[:], nil)

}

var addUriTargetsTargetListInvoker *gi.Function

// AddUriTargets is a representation of the C type gtk_target_list_add_uri_targets.
func (recv *TargetList) AddUriTargets(info uint32) {
	if addUriTargetsTargetListInvoker == nil {
		addUriTargetsTargetListInvoker = gi.StructFunctionInvokerNew("Gtk", "TargetList", "add_uri_targets")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(info)

	addUriTargetsTargetListInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_target_list_find' : parameter 'target' of type 'Gdk.Atom' not supported

var refTargetListInvoker *gi.Function

// Ref is a representation of the C type gtk_target_list_ref.
func (recv *TargetList) Ref() *TargetList {
	if refTargetListInvoker == nil {
		refTargetListInvoker = gi.StructFunctionInvokerNew("Gtk", "TargetList", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refTargetListInvoker.Invoke(inArgs[:], nil)

	retGo := &TargetList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_target_list_remove' : parameter 'target' of type 'Gdk.Atom' not supported

var unrefTargetListInvoker *gi.Function

// Unref is a representation of the C type gtk_target_list_unref.
func (recv *TargetList) Unref() {
	if unrefTargetListInvoker == nil {
		unrefTargetListInvoker = gi.StructFunctionInvokerNew("Gtk", "TargetList", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefTargetListInvoker.Invoke(inArgs[:], nil)

}

var TargetPairStruct *gi.Struct
var TargetPairStructOnce sync.Once

func TargetPairStructSet() {
	TargetPairStructOnce.Do(func() {
		TargetPairStruct = gi.StructNew("Gtk", "TargetPair")
	})
}

type TargetPair struct {
	native uintptr
	// UNSUPPORTED : C value 'target' : no Go type for 'Gdk.Atom'
	Flags uint32
	Info  uint32
}

var TearoffMenuItemClassStruct *gi.Struct
var TearoffMenuItemClassStructOnce sync.Once

func TearoffMenuItemClassStructSet() {
	TearoffMenuItemClassStructOnce.Do(func() {
		TearoffMenuItemClassStruct = gi.StructNew("Gtk", "TearoffMenuItemClass")
	})
}

type TearoffMenuItemClass struct {
	native      uintptr
	ParentClass *MenuItemClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var TearoffMenuItemPrivateStruct *gi.Struct
var TearoffMenuItemPrivateStructOnce sync.Once

func TearoffMenuItemPrivateStructSet() {
	TearoffMenuItemPrivateStructOnce.Do(func() {
		TearoffMenuItemPrivateStruct = gi.StructNew("Gtk", "TearoffMenuItemPrivate")
	})
}

type TearoffMenuItemPrivate struct {
	native uintptr
}

var TextAppearanceStruct *gi.Struct
var TextAppearanceStructOnce sync.Once

func TextAppearanceStructSet() {
	TextAppearanceStructOnce.Do(func() {
		TextAppearanceStruct = gi.StructNew("Gtk", "TextAppearance")
	})
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

var TextAttributesStruct *gi.Struct
var TextAttributesStructOnce sync.Once

func TextAttributesStructSet() {
	TextAttributesStructOnce.Do(func() {
		TextAttributesStruct = gi.StructNew("Gtk", "TextAttributes")
	})
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

	retGo := &TextAttributes{native: ret.Pointer()}

	return retGo
}

var copyTextAttributesInvoker *gi.Function

// Copy is a representation of the C type gtk_text_attributes_copy.
func (recv *TextAttributes) Copy() *TextAttributes {
	if copyTextAttributesInvoker == nil {
		copyTextAttributesInvoker = gi.StructFunctionInvokerNew("Gtk", "TextAttributes", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyTextAttributesInvoker.Invoke(inArgs[:], nil)

	retGo := &TextAttributes{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_attributes_copy_values' : parameter 'dest' of type 'TextAttributes' not supported

var refTextAttributesInvoker *gi.Function

// Ref is a representation of the C type gtk_text_attributes_ref.
func (recv *TextAttributes) Ref() *TextAttributes {
	if refTextAttributesInvoker == nil {
		refTextAttributesInvoker = gi.StructFunctionInvokerNew("Gtk", "TextAttributes", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refTextAttributesInvoker.Invoke(inArgs[:], nil)

	retGo := &TextAttributes{native: ret.Pointer()}

	return retGo
}

var unrefTextAttributesInvoker *gi.Function

// Unref is a representation of the C type gtk_text_attributes_unref.
func (recv *TextAttributes) Unref() {
	if unrefTextAttributesInvoker == nil {
		unrefTextAttributesInvoker = gi.StructFunctionInvokerNew("Gtk", "TextAttributes", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefTextAttributesInvoker.Invoke(inArgs[:], nil)

}

var TextBTreeStruct *gi.Struct
var TextBTreeStructOnce sync.Once

func TextBTreeStructSet() {
	TextBTreeStructOnce.Do(func() {
		TextBTreeStruct = gi.StructNew("Gtk", "TextBTree")
	})
}

type TextBTree struct {
	native uintptr
}

var TextBufferClassStruct *gi.Struct
var TextBufferClassStructOnce sync.Once

func TextBufferClassStructSet() {
	TextBufferClassStructOnce.Do(func() {
		TextBufferClassStruct = gi.StructNew("Gtk", "TextBufferClass")
	})
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

var TextBufferPrivateStruct *gi.Struct
var TextBufferPrivateStructOnce sync.Once

func TextBufferPrivateStructSet() {
	TextBufferPrivateStructOnce.Do(func() {
		TextBufferPrivateStruct = gi.StructNew("Gtk", "TextBufferPrivate")
	})
}

type TextBufferPrivate struct {
	native uintptr
}

var TextCellAccessibleClassStruct *gi.Struct
var TextCellAccessibleClassStructOnce sync.Once

func TextCellAccessibleClassStructSet() {
	TextCellAccessibleClassStructOnce.Do(func() {
		TextCellAccessibleClassStruct = gi.StructNew("Gtk", "TextCellAccessibleClass")
	})
}

type TextCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

var TextCellAccessiblePrivateStruct *gi.Struct
var TextCellAccessiblePrivateStructOnce sync.Once

func TextCellAccessiblePrivateStructSet() {
	TextCellAccessiblePrivateStructOnce.Do(func() {
		TextCellAccessiblePrivateStruct = gi.StructNew("Gtk", "TextCellAccessiblePrivate")
	})
}

type TextCellAccessiblePrivate struct {
	native uintptr
}

var TextChildAnchorClassStruct *gi.Struct
var TextChildAnchorClassStructOnce sync.Once

func TextChildAnchorClassStructSet() {
	TextChildAnchorClassStructOnce.Do(func() {
		TextChildAnchorClassStruct = gi.StructNew("Gtk", "TextChildAnchorClass")
	})
}

type TextChildAnchorClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var TextIterStruct *gi.Struct
var TextIterStructOnce sync.Once

func TextIterStructSet() {
	TextIterStructOnce.Do(func() {
		TextIterStruct = gi.StructNew("Gtk", "TextIter")
	})
}

type TextIter struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_text_iter_assign' : parameter 'other' of type 'TextIter' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_char' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_chars' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_cursor_position' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_cursor_positions' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_find_char' : parameter 'pred' of type 'TextCharPredicate' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_line' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_lines' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_search' : parameter 'flags' of type 'TextSearchFlags' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_sentence_start' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_sentence_starts' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_to_tag_toggle' : parameter 'tag' of type 'TextTag' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_visible_cursor_position' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_visible_cursor_positions' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_visible_line' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_visible_lines' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_visible_word_start' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_visible_word_starts' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_word_start' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_backward_word_starts' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_begins_tag' : parameter 'tag' of type 'TextTag' not supported

// UNSUPPORTED : C value 'gtk_text_iter_can_insert' : parameter 'default_editability' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_compare' : parameter 'rhs' of type 'TextIter' not supported

var copyTextIterInvoker *gi.Function

// Copy is a representation of the C type gtk_text_iter_copy.
func (recv *TextIter) Copy() *TextIter {
	if copyTextIterInvoker == nil {
		copyTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyTextIterInvoker.Invoke(inArgs[:], nil)

	retGo := &TextIter{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_editable' : parameter 'default_setting' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_ends_line' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_ends_sentence' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_ends_tag' : parameter 'tag' of type 'TextTag' not supported

// UNSUPPORTED : C value 'gtk_text_iter_ends_word' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_equal' : parameter 'rhs' of type 'TextIter' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_char' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_chars' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_cursor_position' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_cursor_positions' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_find_char' : parameter 'pred' of type 'TextCharPredicate' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_line' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_lines' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_search' : parameter 'flags' of type 'TextSearchFlags' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_sentence_end' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_sentence_ends' : return type 'gboolean' not supported

var forwardToEndTextIterInvoker *gi.Function

// ForwardToEnd is a representation of the C type gtk_text_iter_forward_to_end.
func (recv *TextIter) ForwardToEnd() {
	if forwardToEndTextIterInvoker == nil {
		forwardToEndTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "forward_to_end")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	forwardToEndTextIterInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_text_iter_forward_to_line_end' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_to_tag_toggle' : parameter 'tag' of type 'TextTag' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_visible_cursor_position' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_visible_cursor_positions' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_visible_line' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_visible_lines' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_visible_word_end' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_visible_word_ends' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_word_end' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_forward_word_ends' : return type 'gboolean' not supported

var freeTextIterInvoker *gi.Function

// Free is a representation of the C type gtk_text_iter_free.
func (recv *TextIter) Free() {
	if freeTextIterInvoker == nil {
		freeTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeTextIterInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_text_iter_get_attributes' : parameter 'values' of type 'TextAttributes' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_buffer' : return type 'TextBuffer' not supported

var getBytesInLineTextIterInvoker *gi.Function

// GetBytesInLine is a representation of the C type gtk_text_iter_get_bytes_in_line.
func (recv *TextIter) GetBytesInLine() int32 {
	if getBytesInLineTextIterInvoker == nil {
		getBytesInLineTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "get_bytes_in_line")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getBytesInLineTextIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_char' : return type 'gunichar' not supported

var getCharsInLineTextIterInvoker *gi.Function

// GetCharsInLine is a representation of the C type gtk_text_iter_get_chars_in_line.
func (recv *TextIter) GetCharsInLine() int32 {
	if getCharsInLineTextIterInvoker == nil {
		getCharsInLineTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "get_chars_in_line")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getCharsInLineTextIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_child_anchor' : return type 'TextChildAnchor' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_language' : return type 'Pango.Language' not supported

var getLineTextIterInvoker *gi.Function

// GetLine is a representation of the C type gtk_text_iter_get_line.
func (recv *TextIter) GetLine() int32 {
	if getLineTextIterInvoker == nil {
		getLineTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "get_line")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLineTextIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getLineIndexTextIterInvoker *gi.Function

// GetLineIndex is a representation of the C type gtk_text_iter_get_line_index.
func (recv *TextIter) GetLineIndex() int32 {
	if getLineIndexTextIterInvoker == nil {
		getLineIndexTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "get_line_index")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLineIndexTextIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getLineOffsetTextIterInvoker *gi.Function

// GetLineOffset is a representation of the C type gtk_text_iter_get_line_offset.
func (recv *TextIter) GetLineOffset() int32 {
	if getLineOffsetTextIterInvoker == nil {
		getLineOffsetTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "get_line_offset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getLineOffsetTextIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_marks' : return type 'GLib.SList' not supported

var getOffsetTextIterInvoker *gi.Function

// GetOffset is a representation of the C type gtk_text_iter_get_offset.
func (recv *TextIter) GetOffset() int32 {
	if getOffsetTextIterInvoker == nil {
		getOffsetTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "get_offset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getOffsetTextIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_slice' : parameter 'end' of type 'TextIter' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_tags' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_text' : parameter 'end' of type 'TextIter' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_toggled_tags' : parameter 'toggled_on' of type 'gboolean' not supported

var getVisibleLineIndexTextIterInvoker *gi.Function

// GetVisibleLineIndex is a representation of the C type gtk_text_iter_get_visible_line_index.
func (recv *TextIter) GetVisibleLineIndex() int32 {
	if getVisibleLineIndexTextIterInvoker == nil {
		getVisibleLineIndexTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "get_visible_line_index")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getVisibleLineIndexTextIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getVisibleLineOffsetTextIterInvoker *gi.Function

// GetVisibleLineOffset is a representation of the C type gtk_text_iter_get_visible_line_offset.
func (recv *TextIter) GetVisibleLineOffset() int32 {
	if getVisibleLineOffsetTextIterInvoker == nil {
		getVisibleLineOffsetTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "get_visible_line_offset")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getVisibleLineOffsetTextIterInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_visible_slice' : parameter 'end' of type 'TextIter' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_visible_text' : parameter 'end' of type 'TextIter' not supported

// UNSUPPORTED : C value 'gtk_text_iter_has_tag' : parameter 'tag' of type 'TextTag' not supported

// UNSUPPORTED : C value 'gtk_text_iter_in_range' : parameter 'start' of type 'TextIter' not supported

// UNSUPPORTED : C value 'gtk_text_iter_inside_sentence' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_inside_word' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_is_cursor_position' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_is_end' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_is_start' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_order' : parameter 'second' of type 'TextIter' not supported

var setLineTextIterInvoker *gi.Function

// SetLine is a representation of the C type gtk_text_iter_set_line.
func (recv *TextIter) SetLine(lineNumber int32) {
	if setLineTextIterInvoker == nil {
		setLineTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "set_line")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(lineNumber)

	setLineTextIterInvoker.Invoke(inArgs[:], nil)

}

var setLineIndexTextIterInvoker *gi.Function

// SetLineIndex is a representation of the C type gtk_text_iter_set_line_index.
func (recv *TextIter) SetLineIndex(byteOnLine int32) {
	if setLineIndexTextIterInvoker == nil {
		setLineIndexTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "set_line_index")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(byteOnLine)

	setLineIndexTextIterInvoker.Invoke(inArgs[:], nil)

}

var setLineOffsetTextIterInvoker *gi.Function

// SetLineOffset is a representation of the C type gtk_text_iter_set_line_offset.
func (recv *TextIter) SetLineOffset(charOnLine int32) {
	if setLineOffsetTextIterInvoker == nil {
		setLineOffsetTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "set_line_offset")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(charOnLine)

	setLineOffsetTextIterInvoker.Invoke(inArgs[:], nil)

}

var setOffsetTextIterInvoker *gi.Function

// SetOffset is a representation of the C type gtk_text_iter_set_offset.
func (recv *TextIter) SetOffset(charOffset int32) {
	if setOffsetTextIterInvoker == nil {
		setOffsetTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "set_offset")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(charOffset)

	setOffsetTextIterInvoker.Invoke(inArgs[:], nil)

}

var setVisibleLineIndexTextIterInvoker *gi.Function

// SetVisibleLineIndex is a representation of the C type gtk_text_iter_set_visible_line_index.
func (recv *TextIter) SetVisibleLineIndex(byteOnLine int32) {
	if setVisibleLineIndexTextIterInvoker == nil {
		setVisibleLineIndexTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "set_visible_line_index")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(byteOnLine)

	setVisibleLineIndexTextIterInvoker.Invoke(inArgs[:], nil)

}

var setVisibleLineOffsetTextIterInvoker *gi.Function

// SetVisibleLineOffset is a representation of the C type gtk_text_iter_set_visible_line_offset.
func (recv *TextIter) SetVisibleLineOffset(charOnLine int32) {
	if setVisibleLineOffsetTextIterInvoker == nil {
		setVisibleLineOffsetTextIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TextIter", "set_visible_line_offset")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(charOnLine)

	setVisibleLineOffsetTextIterInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_text_iter_starts_line' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_starts_sentence' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_starts_tag' : parameter 'tag' of type 'TextTag' not supported

// UNSUPPORTED : C value 'gtk_text_iter_starts_word' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_toggles_tag' : parameter 'tag' of type 'TextTag' not supported

var TextMarkClassStruct *gi.Struct
var TextMarkClassStructOnce sync.Once

func TextMarkClassStructSet() {
	TextMarkClassStructOnce.Do(func() {
		TextMarkClassStruct = gi.StructNew("Gtk", "TextMarkClass")
	})
}

type TextMarkClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var TextTagClassStruct *gi.Struct
var TextTagClassStructOnce sync.Once

func TextTagClassStructSet() {
	TextTagClassStructOnce.Do(func() {
		TextTagClassStruct = gi.StructNew("Gtk", "TextTagClass")
	})
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

var TextTagPrivateStruct *gi.Struct
var TextTagPrivateStructOnce sync.Once

func TextTagPrivateStructSet() {
	TextTagPrivateStructOnce.Do(func() {
		TextTagPrivateStruct = gi.StructNew("Gtk", "TextTagPrivate")
	})
}

type TextTagPrivate struct {
	native uintptr
}

var TextTagTableClassStruct *gi.Struct
var TextTagTableClassStructOnce sync.Once

func TextTagTableClassStructSet() {
	TextTagTableClassStructOnce.Do(func() {
		TextTagTableClassStruct = gi.StructNew("Gtk", "TextTagTableClass")
	})
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

var TextTagTablePrivateStruct *gi.Struct
var TextTagTablePrivateStructOnce sync.Once

func TextTagTablePrivateStructSet() {
	TextTagTablePrivateStructOnce.Do(func() {
		TextTagTablePrivateStruct = gi.StructNew("Gtk", "TextTagTablePrivate")
	})
}

type TextTagTablePrivate struct {
	native uintptr
}

var TextViewAccessibleClassStruct *gi.Struct
var TextViewAccessibleClassStructOnce sync.Once

func TextViewAccessibleClassStructSet() {
	TextViewAccessibleClassStructOnce.Do(func() {
		TextViewAccessibleClassStruct = gi.StructNew("Gtk", "TextViewAccessibleClass")
	})
}

type TextViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var TextViewAccessiblePrivateStruct *gi.Struct
var TextViewAccessiblePrivateStructOnce sync.Once

func TextViewAccessiblePrivateStructSet() {
	TextViewAccessiblePrivateStructOnce.Do(func() {
		TextViewAccessiblePrivateStruct = gi.StructNew("Gtk", "TextViewAccessiblePrivate")
	})
}

type TextViewAccessiblePrivate struct {
	native uintptr
}

var TextViewClassStruct *gi.Struct
var TextViewClassStructOnce sync.Once

func TextViewClassStructSet() {
	TextViewClassStructOnce.Do(func() {
		TextViewClassStruct = gi.StructNew("Gtk", "TextViewClass")
	})
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

var TextViewPrivateStruct *gi.Struct
var TextViewPrivateStructOnce sync.Once

func TextViewPrivateStructSet() {
	TextViewPrivateStructOnce.Do(func() {
		TextViewPrivateStruct = gi.StructNew("Gtk", "TextViewPrivate")
	})
}

type TextViewPrivate struct {
	native uintptr
}

var ThemeEngineStruct *gi.Struct
var ThemeEngineStructOnce sync.Once

func ThemeEngineStructSet() {
	ThemeEngineStructOnce.Do(func() {
		ThemeEngineStruct = gi.StructNew("Gtk", "ThemeEngine")
	})
}

type ThemeEngine struct {
	native uintptr
}

var ThemingEngineClassStruct *gi.Struct
var ThemingEngineClassStructOnce sync.Once

func ThemingEngineClassStructSet() {
	ThemingEngineClassStructOnce.Do(func() {
		ThemingEngineClassStruct = gi.StructNew("Gtk", "ThemingEngineClass")
	})
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

var ThemingEnginePrivateStruct *gi.Struct
var ThemingEnginePrivateStructOnce sync.Once

func ThemingEnginePrivateStructSet() {
	ThemingEnginePrivateStructOnce.Do(func() {
		ThemingEnginePrivateStruct = gi.StructNew("Gtk", "ThemingEnginePrivate")
	})
}

type ThemingEnginePrivate struct {
	native uintptr
}

var ToggleActionClassStruct *gi.Struct
var ToggleActionClassStructOnce sync.Once

func ToggleActionClassStructSet() {
	ToggleActionClassStructOnce.Do(func() {
		ToggleActionClassStruct = gi.StructNew("Gtk", "ToggleActionClass")
	})
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

var ToggleActionEntryStruct *gi.Struct
var ToggleActionEntryStructOnce sync.Once

func ToggleActionEntryStructSet() {
	ToggleActionEntryStructOnce.Do(func() {
		ToggleActionEntryStruct = gi.StructNew("Gtk", "ToggleActionEntry")
	})
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

var ToggleActionPrivateStruct *gi.Struct
var ToggleActionPrivateStructOnce sync.Once

func ToggleActionPrivateStructSet() {
	ToggleActionPrivateStructOnce.Do(func() {
		ToggleActionPrivateStruct = gi.StructNew("Gtk", "ToggleActionPrivate")
	})
}

type ToggleActionPrivate struct {
	native uintptr
}

var ToggleButtonAccessibleClassStruct *gi.Struct
var ToggleButtonAccessibleClassStructOnce sync.Once

func ToggleButtonAccessibleClassStructSet() {
	ToggleButtonAccessibleClassStructOnce.Do(func() {
		ToggleButtonAccessibleClassStruct = gi.StructNew("Gtk", "ToggleButtonAccessibleClass")
	})
}

type ToggleButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var ToggleButtonAccessiblePrivateStruct *gi.Struct
var ToggleButtonAccessiblePrivateStructOnce sync.Once

func ToggleButtonAccessiblePrivateStructSet() {
	ToggleButtonAccessiblePrivateStructOnce.Do(func() {
		ToggleButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "ToggleButtonAccessiblePrivate")
	})
}

type ToggleButtonAccessiblePrivate struct {
	native uintptr
}

var ToggleButtonClassStruct *gi.Struct
var ToggleButtonClassStructOnce sync.Once

func ToggleButtonClassStructSet() {
	ToggleButtonClassStructOnce.Do(func() {
		ToggleButtonClassStruct = gi.StructNew("Gtk", "ToggleButtonClass")
	})
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

var ToggleButtonPrivateStruct *gi.Struct
var ToggleButtonPrivateStructOnce sync.Once

func ToggleButtonPrivateStructSet() {
	ToggleButtonPrivateStructOnce.Do(func() {
		ToggleButtonPrivateStruct = gi.StructNew("Gtk", "ToggleButtonPrivate")
	})
}

type ToggleButtonPrivate struct {
	native uintptr
}

var ToggleToolButtonClassStruct *gi.Struct
var ToggleToolButtonClassStructOnce sync.Once

func ToggleToolButtonClassStructSet() {
	ToggleToolButtonClassStructOnce.Do(func() {
		ToggleToolButtonClassStruct = gi.StructNew("Gtk", "ToggleToolButtonClass")
	})
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

var ToggleToolButtonPrivateStruct *gi.Struct
var ToggleToolButtonPrivateStructOnce sync.Once

func ToggleToolButtonPrivateStructSet() {
	ToggleToolButtonPrivateStructOnce.Do(func() {
		ToggleToolButtonPrivateStruct = gi.StructNew("Gtk", "ToggleToolButtonPrivate")
	})
}

type ToggleToolButtonPrivate struct {
	native uintptr
}

var ToolButtonClassStruct *gi.Struct
var ToolButtonClassStructOnce sync.Once

func ToolButtonClassStructSet() {
	ToolButtonClassStructOnce.Do(func() {
		ToolButtonClassStruct = gi.StructNew("Gtk", "ToolButtonClass")
	})
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

var ToolButtonPrivateStruct *gi.Struct
var ToolButtonPrivateStructOnce sync.Once

func ToolButtonPrivateStructSet() {
	ToolButtonPrivateStructOnce.Do(func() {
		ToolButtonPrivateStruct = gi.StructNew("Gtk", "ToolButtonPrivate")
	})
}

type ToolButtonPrivate struct {
	native uintptr
}

var ToolItemClassStruct *gi.Struct
var ToolItemClassStructOnce sync.Once

func ToolItemClassStructSet() {
	ToolItemClassStructOnce.Do(func() {
		ToolItemClassStruct = gi.StructNew("Gtk", "ToolItemClass")
	})
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

var ToolItemGroupClassStruct *gi.Struct
var ToolItemGroupClassStructOnce sync.Once

func ToolItemGroupClassStructSet() {
	ToolItemGroupClassStructOnce.Do(func() {
		ToolItemGroupClassStruct = gi.StructNew("Gtk", "ToolItemGroupClass")
	})
}

type ToolItemGroupClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ToolItemGroupPrivateStruct *gi.Struct
var ToolItemGroupPrivateStructOnce sync.Once

func ToolItemGroupPrivateStructSet() {
	ToolItemGroupPrivateStructOnce.Do(func() {
		ToolItemGroupPrivateStruct = gi.StructNew("Gtk", "ToolItemGroupPrivate")
	})
}

type ToolItemGroupPrivate struct {
	native uintptr
}

var ToolItemPrivateStruct *gi.Struct
var ToolItemPrivateStructOnce sync.Once

func ToolItemPrivateStructSet() {
	ToolItemPrivateStructOnce.Do(func() {
		ToolItemPrivateStruct = gi.StructNew("Gtk", "ToolItemPrivate")
	})
}

type ToolItemPrivate struct {
	native uintptr
}

var ToolPaletteClassStruct *gi.Struct
var ToolPaletteClassStructOnce sync.Once

func ToolPaletteClassStructSet() {
	ToolPaletteClassStructOnce.Do(func() {
		ToolPaletteClassStruct = gi.StructNew("Gtk", "ToolPaletteClass")
	})
}

type ToolPaletteClass struct {
	native      uintptr
	ParentClass *ContainerClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ToolPalettePrivateStruct *gi.Struct
var ToolPalettePrivateStructOnce sync.Once

func ToolPalettePrivateStructSet() {
	ToolPalettePrivateStructOnce.Do(func() {
		ToolPalettePrivateStruct = gi.StructNew("Gtk", "ToolPalettePrivate")
	})
}

type ToolPalettePrivate struct {
	native uintptr
}

var ToolShellIfaceStruct *gi.Struct
var ToolShellIfaceStructOnce sync.Once

func ToolShellIfaceStructSet() {
	ToolShellIfaceStructOnce.Do(func() {
		ToolShellIfaceStruct = gi.StructNew("Gtk", "ToolShellIface")
	})
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

var ToolbarClassStruct *gi.Struct
var ToolbarClassStructOnce sync.Once

func ToolbarClassStructSet() {
	ToolbarClassStructOnce.Do(func() {
		ToolbarClassStruct = gi.StructNew("Gtk", "ToolbarClass")
	})
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

var ToolbarPrivateStruct *gi.Struct
var ToolbarPrivateStructOnce sync.Once

func ToolbarPrivateStructSet() {
	ToolbarPrivateStructOnce.Do(func() {
		ToolbarPrivateStruct = gi.StructNew("Gtk", "ToolbarPrivate")
	})
}

type ToolbarPrivate struct {
	native uintptr
}

var ToplevelAccessibleClassStruct *gi.Struct
var ToplevelAccessibleClassStructOnce sync.Once

func ToplevelAccessibleClassStructSet() {
	ToplevelAccessibleClassStructOnce.Do(func() {
		ToplevelAccessibleClassStruct = gi.StructNew("Gtk", "ToplevelAccessibleClass")
	})
}

type ToplevelAccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Atk.ObjectClass'
}

var ToplevelAccessiblePrivateStruct *gi.Struct
var ToplevelAccessiblePrivateStructOnce sync.Once

func ToplevelAccessiblePrivateStructSet() {
	ToplevelAccessiblePrivateStructOnce.Do(func() {
		ToplevelAccessiblePrivateStruct = gi.StructNew("Gtk", "ToplevelAccessiblePrivate")
	})
}

type ToplevelAccessiblePrivate struct {
	native uintptr
}

var TreeDragDestIfaceStruct *gi.Struct
var TreeDragDestIfaceStructOnce sync.Once

func TreeDragDestIfaceStructSet() {
	TreeDragDestIfaceStructOnce.Do(func() {
		TreeDragDestIfaceStruct = gi.StructNew("Gtk", "TreeDragDestIface")
	})
}

type TreeDragDestIface struct {
	native uintptr
	// UNSUPPORTED : C value 'drag_data_received' : missing Type
	// UNSUPPORTED : C value 'row_drop_possible' : missing Type
}

var TreeDragSourceIfaceStruct *gi.Struct
var TreeDragSourceIfaceStructOnce sync.Once

func TreeDragSourceIfaceStructSet() {
	TreeDragSourceIfaceStructOnce.Do(func() {
		TreeDragSourceIfaceStruct = gi.StructNew("Gtk", "TreeDragSourceIface")
	})
}

type TreeDragSourceIface struct {
	native uintptr
	// UNSUPPORTED : C value 'row_draggable' : missing Type
	// UNSUPPORTED : C value 'drag_data_get' : missing Type
	// UNSUPPORTED : C value 'drag_data_delete' : missing Type
}

var TreeIterStruct *gi.Struct
var TreeIterStructOnce sync.Once

func TreeIterStructSet() {
	TreeIterStructOnce.Do(func() {
		TreeIterStruct = gi.StructNew("Gtk", "TreeIter")
	})
}

type TreeIter struct {
	native uintptr
	Stamp  int32
	// UNSUPPORTED : C value 'user_data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'user_data2' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'user_data3' : no Go type for 'gpointer'
}

var copyTreeIterInvoker *gi.Function

// Copy is a representation of the C type gtk_tree_iter_copy.
func (recv *TreeIter) Copy() *TreeIter {
	if copyTreeIterInvoker == nil {
		copyTreeIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TreeIter", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyTreeIterInvoker.Invoke(inArgs[:], nil)

	retGo := &TreeIter{native: ret.Pointer()}

	return retGo
}

var freeTreeIterInvoker *gi.Function

// Free is a representation of the C type gtk_tree_iter_free.
func (recv *TreeIter) Free() {
	if freeTreeIterInvoker == nil {
		freeTreeIterInvoker = gi.StructFunctionInvokerNew("Gtk", "TreeIter", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeTreeIterInvoker.Invoke(inArgs[:], nil)

}

var TreeModelFilterClassStruct *gi.Struct
var TreeModelFilterClassStructOnce sync.Once

func TreeModelFilterClassStructSet() {
	TreeModelFilterClassStructOnce.Do(func() {
		TreeModelFilterClassStruct = gi.StructNew("Gtk", "TreeModelFilterClass")
	})
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

var TreeModelFilterPrivateStruct *gi.Struct
var TreeModelFilterPrivateStructOnce sync.Once

func TreeModelFilterPrivateStructSet() {
	TreeModelFilterPrivateStructOnce.Do(func() {
		TreeModelFilterPrivateStruct = gi.StructNew("Gtk", "TreeModelFilterPrivate")
	})
}

type TreeModelFilterPrivate struct {
	native uintptr
}

var TreeModelIfaceStruct *gi.Struct
var TreeModelIfaceStructOnce sync.Once

func TreeModelIfaceStructSet() {
	TreeModelIfaceStructOnce.Do(func() {
		TreeModelIfaceStruct = gi.StructNew("Gtk", "TreeModelIface")
	})
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

var TreeModelSortClassStruct *gi.Struct
var TreeModelSortClassStructOnce sync.Once

func TreeModelSortClassStructSet() {
	TreeModelSortClassStructOnce.Do(func() {
		TreeModelSortClassStruct = gi.StructNew("Gtk", "TreeModelSortClass")
	})
}

type TreeModelSortClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var TreeModelSortPrivateStruct *gi.Struct
var TreeModelSortPrivateStructOnce sync.Once

func TreeModelSortPrivateStructSet() {
	TreeModelSortPrivateStructOnce.Do(func() {
		TreeModelSortPrivateStruct = gi.StructNew("Gtk", "TreeModelSortPrivate")
	})
}

type TreeModelSortPrivate struct {
	native uintptr
}

var TreePathStruct *gi.Struct
var TreePathStructOnce sync.Once

func TreePathStructSet() {
	TreePathStructOnce.Do(func() {
		TreePathStruct = gi.StructNew("Gtk", "TreePath")
	})
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

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var newFirstTreePathInvoker *gi.Function

// TreePathNewFirst is a representation of the C type gtk_tree_path_new_first.
func TreePathNewFirst() *TreePath {
	if newFirstTreePathInvoker == nil {
		newFirstTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "new_first")
	}

	ret := newFirstTreePathInvoker.Invoke(nil, nil)

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
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

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var appendIndexTreePathInvoker *gi.Function

// AppendIndex is a representation of the C type gtk_tree_path_append_index.
func (recv *TreePath) AppendIndex(index int32) {
	if appendIndexTreePathInvoker == nil {
		appendIndexTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "append_index")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)

	appendIndexTreePathInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_tree_path_compare' : parameter 'b' of type 'TreePath' not supported

var copyTreePathInvoker *gi.Function

// Copy is a representation of the C type gtk_tree_path_copy.
func (recv *TreePath) Copy() *TreePath {
	if copyTreePathInvoker == nil {
		copyTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyTreePathInvoker.Invoke(inArgs[:], nil)

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var downTreePathInvoker *gi.Function

// Down is a representation of the C type gtk_tree_path_down.
func (recv *TreePath) Down() {
	if downTreePathInvoker == nil {
		downTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "down")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	downTreePathInvoker.Invoke(inArgs[:], nil)

}

var freeTreePathInvoker *gi.Function

// Free is a representation of the C type gtk_tree_path_free.
func (recv *TreePath) Free() {
	if freeTreePathInvoker == nil {
		freeTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeTreePathInvoker.Invoke(inArgs[:], nil)

}

var getDepthTreePathInvoker *gi.Function

// GetDepth is a representation of the C type gtk_tree_path_get_depth.
func (recv *TreePath) GetDepth() int32 {
	if getDepthTreePathInvoker == nil {
		getDepthTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "get_depth")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getDepthTreePathInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getIndicesTreePathInvoker *gi.Function

// GetIndices is a representation of the C type gtk_tree_path_get_indices.
func (recv *TreePath) GetIndices() int32 {
	if getIndicesTreePathInvoker == nil {
		getIndicesTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "get_indices")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getIndicesTreePathInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var getIndicesWithDepthTreePathInvoker *gi.Function

// GetIndicesWithDepth is a representation of the C type gtk_tree_path_get_indices_with_depth.
func (recv *TreePath) GetIndicesWithDepth() int32 {
	if getIndicesWithDepthTreePathInvoker == nil {
		getIndicesWithDepthTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "get_indices_with_depth")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	getIndicesWithDepthTreePathInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'gtk_tree_path_is_ancestor' : parameter 'descendant' of type 'TreePath' not supported

// UNSUPPORTED : C value 'gtk_tree_path_is_descendant' : parameter 'ancestor' of type 'TreePath' not supported

var nextTreePathInvoker *gi.Function

// Next is a representation of the C type gtk_tree_path_next.
func (recv *TreePath) Next() {
	if nextTreePathInvoker == nil {
		nextTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "next")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	nextTreePathInvoker.Invoke(inArgs[:], nil)

}

var prependIndexTreePathInvoker *gi.Function

// PrependIndex is a representation of the C type gtk_tree_path_prepend_index.
func (recv *TreePath) PrependIndex(index int32) {
	if prependIndexTreePathInvoker == nil {
		prependIndexTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "prepend_index")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)

	prependIndexTreePathInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_tree_path_prev' : return type 'gboolean' not supported

var toStringTreePathInvoker *gi.Function

// ToString is a representation of the C type gtk_tree_path_to_string.
func (recv *TreePath) ToString() string {
	if toStringTreePathInvoker == nil {
		toStringTreePathInvoker = gi.StructFunctionInvokerNew("Gtk", "TreePath", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringTreePathInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_tree_path_up' : return type 'gboolean' not supported

var TreeRowReferenceStruct *gi.Struct
var TreeRowReferenceStructOnce sync.Once

func TreeRowReferenceStructSet() {
	TreeRowReferenceStructOnce.Do(func() {
		TreeRowReferenceStruct = gi.StructNew("Gtk", "TreeRowReference")
	})
}

type TreeRowReference struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_tree_row_reference_new' : parameter 'model' of type 'TreeModel' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_new_proxy' : parameter 'proxy' of type 'GObject.Object' not supported

var copyTreeRowReferenceInvoker *gi.Function

// Copy is a representation of the C type gtk_tree_row_reference_copy.
func (recv *TreeRowReference) Copy() *TreeRowReference {
	if copyTreeRowReferenceInvoker == nil {
		copyTreeRowReferenceInvoker = gi.StructFunctionInvokerNew("Gtk", "TreeRowReference", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyTreeRowReferenceInvoker.Invoke(inArgs[:], nil)

	retGo := &TreeRowReference{native: ret.Pointer()}

	return retGo
}

var freeTreeRowReferenceInvoker *gi.Function

// Free is a representation of the C type gtk_tree_row_reference_free.
func (recv *TreeRowReference) Free() {
	if freeTreeRowReferenceInvoker == nil {
		freeTreeRowReferenceInvoker = gi.StructFunctionInvokerNew("Gtk", "TreeRowReference", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeTreeRowReferenceInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_tree_row_reference_get_model' : return type 'TreeModel' not supported

var getPathTreeRowReferenceInvoker *gi.Function

// GetPath is a representation of the C type gtk_tree_row_reference_get_path.
func (recv *TreeRowReference) GetPath() *TreePath {
	if getPathTreeRowReferenceInvoker == nil {
		getPathTreeRowReferenceInvoker = gi.StructFunctionInvokerNew("Gtk", "TreeRowReference", "get_path")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getPathTreeRowReferenceInvoker.Invoke(inArgs[:], nil)

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_tree_row_reference_valid' : return type 'gboolean' not supported

var TreeSelectionClassStruct *gi.Struct
var TreeSelectionClassStructOnce sync.Once

func TreeSelectionClassStructSet() {
	TreeSelectionClassStructOnce.Do(func() {
		TreeSelectionClassStruct = gi.StructNew("Gtk", "TreeSelectionClass")
	})
}

type TreeSelectionClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value 'changed' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var TreeSelectionPrivateStruct *gi.Struct
var TreeSelectionPrivateStructOnce sync.Once

func TreeSelectionPrivateStructSet() {
	TreeSelectionPrivateStructOnce.Do(func() {
		TreeSelectionPrivateStruct = gi.StructNew("Gtk", "TreeSelectionPrivate")
	})
}

type TreeSelectionPrivate struct {
	native uintptr
}

var TreeSortableIfaceStruct *gi.Struct
var TreeSortableIfaceStructOnce sync.Once

func TreeSortableIfaceStructSet() {
	TreeSortableIfaceStructOnce.Do(func() {
		TreeSortableIfaceStruct = gi.StructNew("Gtk", "TreeSortableIface")
	})
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

var TreeStoreClassStruct *gi.Struct
var TreeStoreClassStructOnce sync.Once

func TreeStoreClassStructSet() {
	TreeStoreClassStructOnce.Do(func() {
		TreeStoreClassStruct = gi.StructNew("Gtk", "TreeStoreClass")
	})
}

type TreeStoreClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var TreeStorePrivateStruct *gi.Struct
var TreeStorePrivateStructOnce sync.Once

func TreeStorePrivateStructSet() {
	TreeStorePrivateStructOnce.Do(func() {
		TreeStorePrivateStruct = gi.StructNew("Gtk", "TreeStorePrivate")
	})
}

type TreeStorePrivate struct {
	native uintptr
}

var TreeViewAccessibleClassStruct *gi.Struct
var TreeViewAccessibleClassStructOnce sync.Once

func TreeViewAccessibleClassStructSet() {
	TreeViewAccessibleClassStructOnce.Do(func() {
		TreeViewAccessibleClassStruct = gi.StructNew("Gtk", "TreeViewAccessibleClass")
	})
}

type TreeViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var TreeViewAccessiblePrivateStruct *gi.Struct
var TreeViewAccessiblePrivateStructOnce sync.Once

func TreeViewAccessiblePrivateStructSet() {
	TreeViewAccessiblePrivateStructOnce.Do(func() {
		TreeViewAccessiblePrivateStruct = gi.StructNew("Gtk", "TreeViewAccessiblePrivate")
	})
}

type TreeViewAccessiblePrivate struct {
	native uintptr
}

var TreeViewClassStruct *gi.Struct
var TreeViewClassStructOnce sync.Once

func TreeViewClassStructSet() {
	TreeViewClassStructOnce.Do(func() {
		TreeViewClassStruct = gi.StructNew("Gtk", "TreeViewClass")
	})
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

var TreeViewColumnClassStruct *gi.Struct
var TreeViewColumnClassStructOnce sync.Once

func TreeViewColumnClassStructSet() {
	TreeViewColumnClassStructOnce.Do(func() {
		TreeViewColumnClassStruct = gi.StructNew("Gtk", "TreeViewColumnClass")
	})
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

var TreeViewColumnPrivateStruct *gi.Struct
var TreeViewColumnPrivateStructOnce sync.Once

func TreeViewColumnPrivateStructSet() {
	TreeViewColumnPrivateStructOnce.Do(func() {
		TreeViewColumnPrivateStruct = gi.StructNew("Gtk", "TreeViewColumnPrivate")
	})
}

type TreeViewColumnPrivate struct {
	native uintptr
}

var TreeViewPrivateStruct *gi.Struct
var TreeViewPrivateStructOnce sync.Once

func TreeViewPrivateStructSet() {
	TreeViewPrivateStructOnce.Do(func() {
		TreeViewPrivateStruct = gi.StructNew("Gtk", "TreeViewPrivate")
	})
}

type TreeViewPrivate struct {
	native uintptr
}

var UIManagerClassStruct *gi.Struct
var UIManagerClassStructOnce sync.Once

func UIManagerClassStructSet() {
	UIManagerClassStructOnce.Do(func() {
		UIManagerClassStruct = gi.StructNew("Gtk", "UIManagerClass")
	})
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

var UIManagerPrivateStruct *gi.Struct
var UIManagerPrivateStructOnce sync.Once

func UIManagerPrivateStructSet() {
	UIManagerPrivateStructOnce.Do(func() {
		UIManagerPrivateStruct = gi.StructNew("Gtk", "UIManagerPrivate")
	})
}

type UIManagerPrivate struct {
	native uintptr
}

var VBoxClassStruct *gi.Struct
var VBoxClassStructOnce sync.Once

func VBoxClassStructSet() {
	VBoxClassStructOnce.Do(func() {
		VBoxClassStruct = gi.StructNew("Gtk", "VBoxClass")
	})
}

type VBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
}

var VButtonBoxClassStruct *gi.Struct
var VButtonBoxClassStructOnce sync.Once

func VButtonBoxClassStructSet() {
	VButtonBoxClassStructOnce.Do(func() {
		VButtonBoxClassStruct = gi.StructNew("Gtk", "VButtonBoxClass")
	})
}

type VButtonBoxClass struct {
	native      uintptr
	ParentClass *ButtonBoxClass
}

var VPanedClassStruct *gi.Struct
var VPanedClassStructOnce sync.Once

func VPanedClassStructSet() {
	VPanedClassStructOnce.Do(func() {
		VPanedClassStruct = gi.StructNew("Gtk", "VPanedClass")
	})
}

type VPanedClass struct {
	native      uintptr
	ParentClass *PanedClass
}

var VScaleClassStruct *gi.Struct
var VScaleClassStructOnce sync.Once

func VScaleClassStructSet() {
	VScaleClassStructOnce.Do(func() {
		VScaleClassStruct = gi.StructNew("Gtk", "VScaleClass")
	})
}

type VScaleClass struct {
	native      uintptr
	ParentClass *ScaleClass
}

var VScrollbarClassStruct *gi.Struct
var VScrollbarClassStructOnce sync.Once

func VScrollbarClassStructSet() {
	VScrollbarClassStructOnce.Do(func() {
		VScrollbarClassStruct = gi.StructNew("Gtk", "VScrollbarClass")
	})
}

type VScrollbarClass struct {
	native      uintptr
	ParentClass *ScrollbarClass
}

var VSeparatorClassStruct *gi.Struct
var VSeparatorClassStructOnce sync.Once

func VSeparatorClassStructSet() {
	VSeparatorClassStructOnce.Do(func() {
		VSeparatorClassStruct = gi.StructNew("Gtk", "VSeparatorClass")
	})
}

type VSeparatorClass struct {
	native      uintptr
	ParentClass *SeparatorClass
}

var ViewportClassStruct *gi.Struct
var ViewportClassStructOnce sync.Once

func ViewportClassStructSet() {
	ViewportClassStructOnce.Do(func() {
		ViewportClassStruct = gi.StructNew("Gtk", "ViewportClass")
	})
}

type ViewportClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var ViewportPrivateStruct *gi.Struct
var ViewportPrivateStructOnce sync.Once

func ViewportPrivateStructSet() {
	ViewportPrivateStructOnce.Do(func() {
		ViewportPrivateStruct = gi.StructNew("Gtk", "ViewportPrivate")
	})
}

type ViewportPrivate struct {
	native uintptr
}

var VolumeButtonClassStruct *gi.Struct
var VolumeButtonClassStructOnce sync.Once

func VolumeButtonClassStructSet() {
	VolumeButtonClassStructOnce.Do(func() {
		VolumeButtonClassStruct = gi.StructNew("Gtk", "VolumeButtonClass")
	})
}

type VolumeButtonClass struct {
	native      uintptr
	ParentClass *ScaleButtonClass
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var WidgetAccessibleClassStruct *gi.Struct
var WidgetAccessibleClassStructOnce sync.Once

func WidgetAccessibleClassStructSet() {
	WidgetAccessibleClassStructOnce.Do(func() {
		WidgetAccessibleClassStruct = gi.StructNew("Gtk", "WidgetAccessibleClass")
	})
}

type WidgetAccessibleClass struct {
	native      uintptr
	ParentClass *AccessibleClass
	// UNSUPPORTED : C value 'notify_gtk' : missing Type
}

var WidgetAccessiblePrivateStruct *gi.Struct
var WidgetAccessiblePrivateStructOnce sync.Once

func WidgetAccessiblePrivateStructSet() {
	WidgetAccessiblePrivateStructOnce.Do(func() {
		WidgetAccessiblePrivateStruct = gi.StructNew("Gtk", "WidgetAccessiblePrivate")
	})
}

type WidgetAccessiblePrivate struct {
	native uintptr
}

var WidgetClassStruct *gi.Struct
var WidgetClassStructOnce sync.Once

func WidgetClassStructSet() {
	WidgetClassStructOnce.Do(func() {
		WidgetClassStruct = gi.StructNew("Gtk", "WidgetClass")
	})
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

// UNSUPPORTED : C value 'gtk_widget_class_bind_template_callback_full' : parameter 'callback_symbol' of type 'GObject.Callback' not supported

// UNSUPPORTED : C value 'gtk_widget_class_bind_template_child_full' : parameter 'internal_child' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_widget_class_find_style_property' : return type 'GObject.ParamSpec' not supported

var getCssNameWidgetClassInvoker *gi.Function

// GetCssName is a representation of the C type gtk_widget_class_get_css_name.
func (recv *WidgetClass) GetCssName() string {
	if getCssNameWidgetClassInvoker == nil {
		getCssNameWidgetClassInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetClass", "get_css_name")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := getCssNameWidgetClassInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_class_install_style_property' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_widget_class_install_style_property_parser' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

var listStylePropertiesWidgetClassInvoker *gi.Function

// ListStyleProperties is a representation of the C type gtk_widget_class_list_style_properties.
func (recv *WidgetClass) ListStyleProperties() uint32 {
	if listStylePropertiesWidgetClassInvoker == nil {
		listStylePropertiesWidgetClassInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetClass", "list_style_properties")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	listStylePropertiesWidgetClassInvoker.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint32()

	return out0
}

// UNSUPPORTED : C value 'gtk_widget_class_set_accessible_role' : parameter 'role' of type 'Atk.Role' not supported

// UNSUPPORTED : C value 'gtk_widget_class_set_accessible_type' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_class_set_connect_func' : parameter 'connect_func' of type 'BuilderConnectFunc' not supported

var setCssNameWidgetClassInvoker *gi.Function

// SetCssName is a representation of the C type gtk_widget_class_set_css_name.
func (recv *WidgetClass) SetCssName(name string) {
	if setCssNameWidgetClassInvoker == nil {
		setCssNameWidgetClassInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetClass", "set_css_name")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	setCssNameWidgetClassInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_widget_class_set_template' : parameter 'template_bytes' of type 'GLib.Bytes' not supported

var setTemplateFromResourceWidgetClassInvoker *gi.Function

// SetTemplateFromResource is a representation of the C type gtk_widget_class_set_template_from_resource.
func (recv *WidgetClass) SetTemplateFromResource(resourceName string) {
	if setTemplateFromResourceWidgetClassInvoker == nil {
		setTemplateFromResourceWidgetClassInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetClass", "set_template_from_resource")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(resourceName)

	setTemplateFromResourceWidgetClassInvoker.Invoke(inArgs[:], nil)

}

var WidgetClassPrivateStruct *gi.Struct
var WidgetClassPrivateStructOnce sync.Once

func WidgetClassPrivateStructSet() {
	WidgetClassPrivateStructOnce.Do(func() {
		WidgetClassPrivateStruct = gi.StructNew("Gtk", "WidgetClassPrivate")
	})
}

type WidgetClassPrivate struct {
	native uintptr
}

var WidgetPathStruct *gi.Struct
var WidgetPathStructOnce sync.Once

func WidgetPathStructSet() {
	WidgetPathStructOnce.Do(func() {
		WidgetPathStruct = gi.StructNew("Gtk", "WidgetPath")
	})
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

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_append_for_widget' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_widget_path_append_type' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_append_with_siblings' : parameter 'siblings' of type 'WidgetPath' not supported

var copyWidgetPathInvoker *gi.Function

// Copy is a representation of the C type gtk_widget_path_copy.
func (recv *WidgetPath) Copy() *WidgetPath {
	if copyWidgetPathInvoker == nil {
		copyWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "copy")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := copyWidgetPathInvoker.Invoke(inArgs[:], nil)

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

var freeWidgetPathInvoker *gi.Function

// Free is a representation of the C type gtk_widget_path_free.
func (recv *WidgetPath) Free() {
	if freeWidgetPathInvoker == nil {
		freeWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "free")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	freeWidgetPathInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_widget_path_get_object_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_has_parent' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_is_type' : parameter 'type' of type 'GType' not supported

var iterAddClassWidgetPathInvoker *gi.Function

// IterAddClass is a representation of the C type gtk_widget_path_iter_add_class.
func (recv *WidgetPath) IterAddClass(pos int32, name string) {
	if iterAddClassWidgetPathInvoker == nil {
		iterAddClassWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_add_class")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	iterAddClassWidgetPathInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_widget_path_iter_add_region' : parameter 'flags' of type 'RegionFlags' not supported

var iterClearClassesWidgetPathInvoker *gi.Function

// IterClearClasses is a representation of the C type gtk_widget_path_iter_clear_classes.
func (recv *WidgetPath) IterClearClasses(pos int32) {
	if iterClearClassesWidgetPathInvoker == nil {
		iterClearClassesWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_clear_classes")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	iterClearClassesWidgetPathInvoker.Invoke(inArgs[:], nil)

}

var iterClearRegionsWidgetPathInvoker *gi.Function

// IterClearRegions is a representation of the C type gtk_widget_path_iter_clear_regions.
func (recv *WidgetPath) IterClearRegions(pos int32) {
	if iterClearRegionsWidgetPathInvoker == nil {
		iterClearRegionsWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_clear_regions")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	iterClearRegionsWidgetPathInvoker.Invoke(inArgs[:], nil)

}

var iterGetNameWidgetPathInvoker *gi.Function

// IterGetName is a representation of the C type gtk_widget_path_iter_get_name.
func (recv *WidgetPath) IterGetName(pos int32) string {
	if iterGetNameWidgetPathInvoker == nil {
		iterGetNameWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_get_name")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := iterGetNameWidgetPathInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var iterGetObjectNameWidgetPathInvoker *gi.Function

// IterGetObjectName is a representation of the C type gtk_widget_path_iter_get_object_name.
func (recv *WidgetPath) IterGetObjectName(pos int32) string {
	if iterGetObjectNameWidgetPathInvoker == nil {
		iterGetObjectNameWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_get_object_name")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := iterGetObjectNameWidgetPathInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_iter_get_object_type' : return type 'GType' not supported

var iterGetSiblingIndexWidgetPathInvoker *gi.Function

// IterGetSiblingIndex is a representation of the C type gtk_widget_path_iter_get_sibling_index.
func (recv *WidgetPath) IterGetSiblingIndex(pos int32) uint32 {
	if iterGetSiblingIndexWidgetPathInvoker == nil {
		iterGetSiblingIndexWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_get_sibling_index")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := iterGetSiblingIndexWidgetPathInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var iterGetSiblingsWidgetPathInvoker *gi.Function

// IterGetSiblings is a representation of the C type gtk_widget_path_iter_get_siblings.
func (recv *WidgetPath) IterGetSiblings(pos int32) *WidgetPath {
	if iterGetSiblingsWidgetPathInvoker == nil {
		iterGetSiblingsWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_get_siblings")
	}

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := iterGetSiblingsWidgetPathInvoker.Invoke(inArgs[:], nil)

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_iter_get_state' : return type 'StateFlags' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_class' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_name' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_qclass' : parameter 'qname' of type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_qname' : parameter 'qname' of type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_qregion' : parameter 'qname' of type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_region' : parameter 'flags' of type 'RegionFlags' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_list_classes' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_list_regions' : return type 'GLib.SList' not supported

var iterRemoveClassWidgetPathInvoker *gi.Function

// IterRemoveClass is a representation of the C type gtk_widget_path_iter_remove_class.
func (recv *WidgetPath) IterRemoveClass(pos int32, name string) {
	if iterRemoveClassWidgetPathInvoker == nil {
		iterRemoveClassWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_remove_class")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	iterRemoveClassWidgetPathInvoker.Invoke(inArgs[:], nil)

}

var iterRemoveRegionWidgetPathInvoker *gi.Function

// IterRemoveRegion is a representation of the C type gtk_widget_path_iter_remove_region.
func (recv *WidgetPath) IterRemoveRegion(pos int32, name string) {
	if iterRemoveRegionWidgetPathInvoker == nil {
		iterRemoveRegionWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_remove_region")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	iterRemoveRegionWidgetPathInvoker.Invoke(inArgs[:], nil)

}

var iterSetNameWidgetPathInvoker *gi.Function

// IterSetName is a representation of the C type gtk_widget_path_iter_set_name.
func (recv *WidgetPath) IterSetName(pos int32, name string) {
	if iterSetNameWidgetPathInvoker == nil {
		iterSetNameWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_set_name")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	iterSetNameWidgetPathInvoker.Invoke(inArgs[:], nil)

}

var iterSetObjectNameWidgetPathInvoker *gi.Function

// IterSetObjectName is a representation of the C type gtk_widget_path_iter_set_object_name.
func (recv *WidgetPath) IterSetObjectName(pos int32, name string) {
	if iterSetObjectNameWidgetPathInvoker == nil {
		iterSetObjectNameWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "iter_set_object_name")
	}

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	iterSetObjectNameWidgetPathInvoker.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_widget_path_iter_set_object_type' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_set_state' : parameter 'state' of type 'StateFlags' not supported

var lengthWidgetPathInvoker *gi.Function

// Length is a representation of the C type gtk_widget_path_length.
func (recv *WidgetPath) Length() int32 {
	if lengthWidgetPathInvoker == nil {
		lengthWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "length")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := lengthWidgetPathInvoker.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_prepend_type' : parameter 'type' of type 'GType' not supported

var refWidgetPathInvoker *gi.Function

// Ref is a representation of the C type gtk_widget_path_ref.
func (recv *WidgetPath) Ref() *WidgetPath {
	if refWidgetPathInvoker == nil {
		refWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "ref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := refWidgetPathInvoker.Invoke(inArgs[:], nil)

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

var toStringWidgetPathInvoker *gi.Function

// ToString is a representation of the C type gtk_widget_path_to_string.
func (recv *WidgetPath) ToString() string {
	if toStringWidgetPathInvoker == nil {
		toStringWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "to_string")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := toStringWidgetPathInvoker.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var unrefWidgetPathInvoker *gi.Function

// Unref is a representation of the C type gtk_widget_path_unref.
func (recv *WidgetPath) Unref() {
	if unrefWidgetPathInvoker == nil {
		unrefWidgetPathInvoker = gi.StructFunctionInvokerNew("Gtk", "WidgetPath", "unref")
	}

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	unrefWidgetPathInvoker.Invoke(inArgs[:], nil)

}

var WidgetPrivateStruct *gi.Struct
var WidgetPrivateStructOnce sync.Once

func WidgetPrivateStructSet() {
	WidgetPrivateStructOnce.Do(func() {
		WidgetPrivateStruct = gi.StructNew("Gtk", "WidgetPrivate")
	})
}

type WidgetPrivate struct {
	native uintptr
}

var WindowAccessibleClassStruct *gi.Struct
var WindowAccessibleClassStructOnce sync.Once

func WindowAccessibleClassStructSet() {
	WindowAccessibleClassStructOnce.Do(func() {
		WindowAccessibleClassStruct = gi.StructNew("Gtk", "WindowAccessibleClass")
	})
}

type WindowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var WindowAccessiblePrivateStruct *gi.Struct
var WindowAccessiblePrivateStructOnce sync.Once

func WindowAccessiblePrivateStructSet() {
	WindowAccessiblePrivateStructOnce.Do(func() {
		WindowAccessiblePrivateStruct = gi.StructNew("Gtk", "WindowAccessiblePrivate")
	})
}

type WindowAccessiblePrivate struct {
	native uintptr
}

var WindowClassStruct *gi.Struct
var WindowClassStructOnce sync.Once

func WindowClassStructSet() {
	WindowClassStructOnce.Do(func() {
		WindowClassStruct = gi.StructNew("Gtk", "WindowClass")
	})
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

var WindowGeometryInfoStruct *gi.Struct
var WindowGeometryInfoStructOnce sync.Once

func WindowGeometryInfoStructSet() {
	WindowGeometryInfoStructOnce.Do(func() {
		WindowGeometryInfoStruct = gi.StructNew("Gtk", "WindowGeometryInfo")
	})
}

type WindowGeometryInfo struct {
	native uintptr
}

var WindowGroupClassStruct *gi.Struct
var WindowGroupClassStructOnce sync.Once

func WindowGroupClassStructSet() {
	WindowGroupClassStructOnce.Do(func() {
		WindowGroupClassStruct = gi.StructNew("Gtk", "WindowGroupClass")
	})
}

type WindowGroupClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var WindowGroupPrivateStruct *gi.Struct
var WindowGroupPrivateStructOnce sync.Once

func WindowGroupPrivateStructSet() {
	WindowGroupPrivateStructOnce.Do(func() {
		WindowGroupPrivateStruct = gi.StructNew("Gtk", "WindowGroupPrivate")
	})
}

type WindowGroupPrivate struct {
	native uintptr
}

var WindowPrivateStruct *gi.Struct
var WindowPrivateStructOnce sync.Once

func WindowPrivateStructSet() {
	WindowPrivateStructOnce.Do(func() {
		WindowPrivateStruct = gi.StructNew("Gtk", "WindowPrivate")
	})
}

type WindowPrivate struct {
	native uintptr
}
