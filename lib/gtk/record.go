// Code generated - DO NOT EDIT.

package gtk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var aboutDialogClassStruct *gi.Struct
var aboutDialogClassStructOnce sync.Once

func aboutDialogClassStructSet() {
	aboutDialogClassStructOnce.Do(func() {
		aboutDialogClassStruct = gi.StructNew("Gtk", "AboutDialogClass")
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

var aboutDialogPrivateStruct *gi.Struct
var aboutDialogPrivateStructOnce sync.Once

func aboutDialogPrivateStructSet() {
	aboutDialogPrivateStructOnce.Do(func() {
		aboutDialogPrivateStruct = gi.StructNew("Gtk", "AboutDialogPrivate")
	})
}

type AboutDialogPrivate struct {
	native uintptr
}

var accelGroupClassStruct *gi.Struct
var accelGroupClassStructOnce sync.Once

func accelGroupClassStructSet() {
	accelGroupClassStructOnce.Do(func() {
		accelGroupClassStruct = gi.StructNew("Gtk", "AccelGroupClass")
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

var accelGroupEntryStruct *gi.Struct
var accelGroupEntryStructOnce sync.Once

func accelGroupEntryStructSet() {
	accelGroupEntryStructOnce.Do(func() {
		accelGroupEntryStruct = gi.StructNew("Gtk", "AccelGroupEntry")
	})
}

type AccelGroupEntry struct {
	native uintptr
	Key    *AccelKey
	// UNSUPPORTED : C value 'closure' : no Go type for 'GObject.Closure'
	// UNSUPPORTED : C value 'accel_path_quark' : no Go type for 'GLib.Quark'
}

var accelGroupPrivateStruct *gi.Struct
var accelGroupPrivateStructOnce sync.Once

func accelGroupPrivateStructSet() {
	accelGroupPrivateStructOnce.Do(func() {
		accelGroupPrivateStruct = gi.StructNew("Gtk", "AccelGroupPrivate")
	})
}

type AccelGroupPrivate struct {
	native uintptr
}

var accelKeyStruct *gi.Struct
var accelKeyStructOnce sync.Once

func accelKeyStructSet() {
	accelKeyStructOnce.Do(func() {
		accelKeyStruct = gi.StructNew("Gtk", "AccelKey")
	})
}

type AccelKey struct {
	native   uintptr
	AccelKey uint32
	// UNSUPPORTED : C value 'accel_mods' : no Go type for 'Gdk.ModifierType'
	AccelFlags uint32
}

var accelLabelClassStruct *gi.Struct
var accelLabelClassStructOnce sync.Once

func accelLabelClassStructSet() {
	accelLabelClassStructOnce.Do(func() {
		accelLabelClassStruct = gi.StructNew("Gtk", "AccelLabelClass")
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

var accelLabelPrivateStruct *gi.Struct
var accelLabelPrivateStructOnce sync.Once

func accelLabelPrivateStructSet() {
	accelLabelPrivateStructOnce.Do(func() {
		accelLabelPrivateStruct = gi.StructNew("Gtk", "AccelLabelPrivate")
	})
}

type AccelLabelPrivate struct {
	native uintptr
}

var accelMapClassStruct *gi.Struct
var accelMapClassStructOnce sync.Once

func accelMapClassStructSet() {
	accelMapClassStructOnce.Do(func() {
		accelMapClassStruct = gi.StructNew("Gtk", "AccelMapClass")
	})
}

type AccelMapClass struct {
	native uintptr
}

var accessibleClassStruct *gi.Struct
var accessibleClassStructOnce sync.Once

func accessibleClassStructSet() {
	accessibleClassStructOnce.Do(func() {
		accessibleClassStruct = gi.StructNew("Gtk", "AccessibleClass")
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

var accessiblePrivateStruct *gi.Struct
var accessiblePrivateStructOnce sync.Once

func accessiblePrivateStructSet() {
	accessiblePrivateStructOnce.Do(func() {
		accessiblePrivateStruct = gi.StructNew("Gtk", "AccessiblePrivate")
	})
}

type AccessiblePrivate struct {
	native uintptr
}

var actionBarClassStruct *gi.Struct
var actionBarClassStructOnce sync.Once

func actionBarClassStructSet() {
	actionBarClassStructOnce.Do(func() {
		actionBarClassStruct = gi.StructNew("Gtk", "ActionBarClass")
	})
}

type ActionBarClass struct {
	native uintptr
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var actionBarPrivateStruct *gi.Struct
var actionBarPrivateStructOnce sync.Once

func actionBarPrivateStructSet() {
	actionBarPrivateStructOnce.Do(func() {
		actionBarPrivateStruct = gi.StructNew("Gtk", "ActionBarPrivate")
	})
}

type ActionBarPrivate struct {
	native uintptr
}

var actionClassStruct *gi.Struct
var actionClassStructOnce sync.Once

func actionClassStructSet() {
	actionClassStructOnce.Do(func() {
		actionClassStruct = gi.StructNew("Gtk", "ActionClass")
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

var actionEntryStruct *gi.Struct
var actionEntryStructOnce sync.Once

func actionEntryStructSet() {
	actionEntryStructOnce.Do(func() {
		actionEntryStruct = gi.StructNew("Gtk", "ActionEntry")
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

var actionGroupClassStruct *gi.Struct
var actionGroupClassStructOnce sync.Once

func actionGroupClassStructSet() {
	actionGroupClassStructOnce.Do(func() {
		actionGroupClassStruct = gi.StructNew("Gtk", "ActionGroupClass")
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

var actionGroupPrivateStruct *gi.Struct
var actionGroupPrivateStructOnce sync.Once

func actionGroupPrivateStructSet() {
	actionGroupPrivateStructOnce.Do(func() {
		actionGroupPrivateStruct = gi.StructNew("Gtk", "ActionGroupPrivate")
	})
}

type ActionGroupPrivate struct {
	native uintptr
}

var actionPrivateStruct *gi.Struct
var actionPrivateStructOnce sync.Once

func actionPrivateStructSet() {
	actionPrivateStructOnce.Do(func() {
		actionPrivateStruct = gi.StructNew("Gtk", "ActionPrivate")
	})
}

type ActionPrivate struct {
	native uintptr
}

var actionableInterfaceStruct *gi.Struct
var actionableInterfaceStructOnce sync.Once

func actionableInterfaceStructSet() {
	actionableInterfaceStructOnce.Do(func() {
		actionableInterfaceStruct = gi.StructNew("Gtk", "ActionableInterface")
	})
}

type ActionableInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'get_action_name' : missing Type
	// UNSUPPORTED : C value 'set_action_name' : missing Type
	// UNSUPPORTED : C value 'get_action_target_value' : missing Type
	// UNSUPPORTED : C value 'set_action_target_value' : missing Type
}

var activatableIfaceStruct *gi.Struct
var activatableIfaceStructOnce sync.Once

func activatableIfaceStructSet() {
	activatableIfaceStructOnce.Do(func() {
		activatableIfaceStruct = gi.StructNew("Gtk", "ActivatableIface")
	})
}

type ActivatableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'update' : missing Type
	// UNSUPPORTED : C value 'sync_action_properties' : missing Type
}

var adjustmentClassStruct *gi.Struct
var adjustmentClassStructOnce sync.Once

func adjustmentClassStructSet() {
	adjustmentClassStructOnce.Do(func() {
		adjustmentClassStruct = gi.StructNew("Gtk", "AdjustmentClass")
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

var adjustmentPrivateStruct *gi.Struct
var adjustmentPrivateStructOnce sync.Once

func adjustmentPrivateStructSet() {
	adjustmentPrivateStructOnce.Do(func() {
		adjustmentPrivateStruct = gi.StructNew("Gtk", "AdjustmentPrivate")
	})
}

type AdjustmentPrivate struct {
	native uintptr
}

var alignmentClassStruct *gi.Struct
var alignmentClassStructOnce sync.Once

func alignmentClassStructSet() {
	alignmentClassStructOnce.Do(func() {
		alignmentClassStruct = gi.StructNew("Gtk", "AlignmentClass")
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

var alignmentPrivateStruct *gi.Struct
var alignmentPrivateStructOnce sync.Once

func alignmentPrivateStructSet() {
	alignmentPrivateStructOnce.Do(func() {
		alignmentPrivateStruct = gi.StructNew("Gtk", "AlignmentPrivate")
	})
}

type AlignmentPrivate struct {
	native uintptr
}

var appChooserButtonClassStruct *gi.Struct
var appChooserButtonClassStructOnce sync.Once

func appChooserButtonClassStructSet() {
	appChooserButtonClassStructOnce.Do(func() {
		appChooserButtonClassStruct = gi.StructNew("Gtk", "AppChooserButtonClass")
	})
}

type AppChooserButtonClass struct {
	native      uintptr
	ParentClass *ComboBoxClass
	// UNSUPPORTED : C value 'custom_item_activated' : missing Type
}

var appChooserButtonPrivateStruct *gi.Struct
var appChooserButtonPrivateStructOnce sync.Once

func appChooserButtonPrivateStructSet() {
	appChooserButtonPrivateStructOnce.Do(func() {
		appChooserButtonPrivateStruct = gi.StructNew("Gtk", "AppChooserButtonPrivate")
	})
}

type AppChooserButtonPrivate struct {
	native uintptr
}

var appChooserDialogClassStruct *gi.Struct
var appChooserDialogClassStructOnce sync.Once

func appChooserDialogClassStructSet() {
	appChooserDialogClassStructOnce.Do(func() {
		appChooserDialogClassStruct = gi.StructNew("Gtk", "AppChooserDialogClass")
	})
}

type AppChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
}

var appChooserDialogPrivateStruct *gi.Struct
var appChooserDialogPrivateStructOnce sync.Once

func appChooserDialogPrivateStructSet() {
	appChooserDialogPrivateStructOnce.Do(func() {
		appChooserDialogPrivateStruct = gi.StructNew("Gtk", "AppChooserDialogPrivate")
	})
}

type AppChooserDialogPrivate struct {
	native uintptr
}

var appChooserWidgetClassStruct *gi.Struct
var appChooserWidgetClassStructOnce sync.Once

func appChooserWidgetClassStructSet() {
	appChooserWidgetClassStructOnce.Do(func() {
		appChooserWidgetClassStruct = gi.StructNew("Gtk", "AppChooserWidgetClass")
	})
}

type AppChooserWidgetClass struct {
	native      uintptr
	ParentClass *BoxClass
	// UNSUPPORTED : C value 'application_selected' : missing Type
	// UNSUPPORTED : C value 'application_activated' : missing Type
	// UNSUPPORTED : C value 'populate_popup' : missing Type
}

var appChooserWidgetPrivateStruct *gi.Struct
var appChooserWidgetPrivateStructOnce sync.Once

func appChooserWidgetPrivateStructSet() {
	appChooserWidgetPrivateStructOnce.Do(func() {
		appChooserWidgetPrivateStruct = gi.StructNew("Gtk", "AppChooserWidgetPrivate")
	})
}

type AppChooserWidgetPrivate struct {
	native uintptr
}

var applicationClassStruct *gi.Struct
var applicationClassStructOnce sync.Once

func applicationClassStructSet() {
	applicationClassStructOnce.Do(func() {
		applicationClassStruct = gi.StructNew("Gtk", "ApplicationClass")
	})
}

type ApplicationClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.ApplicationClass'
	// UNSUPPORTED : C value 'window_added' : missing Type
	// UNSUPPORTED : C value 'window_removed' : missing Type
}

var applicationPrivateStruct *gi.Struct
var applicationPrivateStructOnce sync.Once

func applicationPrivateStructSet() {
	applicationPrivateStructOnce.Do(func() {
		applicationPrivateStruct = gi.StructNew("Gtk", "ApplicationPrivate")
	})
}

type ApplicationPrivate struct {
	native uintptr
}

var applicationWindowClassStruct *gi.Struct
var applicationWindowClassStructOnce sync.Once

func applicationWindowClassStructSet() {
	applicationWindowClassStructOnce.Do(func() {
		applicationWindowClassStruct = gi.StructNew("Gtk", "ApplicationWindowClass")
	})
}

type ApplicationWindowClass struct {
	native      uintptr
	ParentClass *WindowClass
}

var applicationWindowPrivateStruct *gi.Struct
var applicationWindowPrivateStructOnce sync.Once

func applicationWindowPrivateStructSet() {
	applicationWindowPrivateStructOnce.Do(func() {
		applicationWindowPrivateStruct = gi.StructNew("Gtk", "ApplicationWindowPrivate")
	})
}

type ApplicationWindowPrivate struct {
	native uintptr
}

var arrowAccessibleClassStruct *gi.Struct
var arrowAccessibleClassStructOnce sync.Once

func arrowAccessibleClassStructSet() {
	arrowAccessibleClassStructOnce.Do(func() {
		arrowAccessibleClassStruct = gi.StructNew("Gtk", "ArrowAccessibleClass")
	})
}

type ArrowAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var arrowAccessiblePrivateStruct *gi.Struct
var arrowAccessiblePrivateStructOnce sync.Once

func arrowAccessiblePrivateStructSet() {
	arrowAccessiblePrivateStructOnce.Do(func() {
		arrowAccessiblePrivateStruct = gi.StructNew("Gtk", "ArrowAccessiblePrivate")
	})
}

type ArrowAccessiblePrivate struct {
	native uintptr
}

var arrowClassStruct *gi.Struct
var arrowClassStructOnce sync.Once

func arrowClassStructSet() {
	arrowClassStructOnce.Do(func() {
		arrowClassStruct = gi.StructNew("Gtk", "ArrowClass")
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

var arrowPrivateStruct *gi.Struct
var arrowPrivateStructOnce sync.Once

func arrowPrivateStructSet() {
	arrowPrivateStructOnce.Do(func() {
		arrowPrivateStruct = gi.StructNew("Gtk", "ArrowPrivate")
	})
}

type ArrowPrivate struct {
	native uintptr
}

var aspectFrameClassStruct *gi.Struct
var aspectFrameClassStructOnce sync.Once

func aspectFrameClassStructSet() {
	aspectFrameClassStructOnce.Do(func() {
		aspectFrameClassStruct = gi.StructNew("Gtk", "AspectFrameClass")
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

var aspectFramePrivateStruct *gi.Struct
var aspectFramePrivateStructOnce sync.Once

func aspectFramePrivateStructSet() {
	aspectFramePrivateStructOnce.Do(func() {
		aspectFramePrivateStruct = gi.StructNew("Gtk", "AspectFramePrivate")
	})
}

type AspectFramePrivate struct {
	native uintptr
}

var assistantClassStruct *gi.Struct
var assistantClassStructOnce sync.Once

func assistantClassStructSet() {
	assistantClassStructOnce.Do(func() {
		assistantClassStruct = gi.StructNew("Gtk", "AssistantClass")
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

var assistantPrivateStruct *gi.Struct
var assistantPrivateStructOnce sync.Once

func assistantPrivateStructSet() {
	assistantPrivateStructOnce.Do(func() {
		assistantPrivateStruct = gi.StructNew("Gtk", "AssistantPrivate")
	})
}

type AssistantPrivate struct {
	native uintptr
}

var binClassStruct *gi.Struct
var binClassStructOnce sync.Once

func binClassStructSet() {
	binClassStructOnce.Do(func() {
		binClassStruct = gi.StructNew("Gtk", "BinClass")
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

var binPrivateStruct *gi.Struct
var binPrivateStructOnce sync.Once

func binPrivateStructSet() {
	binPrivateStructOnce.Do(func() {
		binPrivateStruct = gi.StructNew("Gtk", "BinPrivate")
	})
}

type BinPrivate struct {
	native uintptr
}

var bindingArgStruct *gi.Struct
var bindingArgStructOnce sync.Once

func bindingArgStructSet() {
	bindingArgStructOnce.Do(func() {
		bindingArgStruct = gi.StructNew("Gtk", "BindingArg")
	})
}

type BindingArg struct {
	native uintptr
	// UNSUPPORTED : C value 'arg_type' : no Go type for 'GType'
}

var bindingEntryStruct *gi.Struct
var bindingEntryStructOnce sync.Once

func bindingEntryStructSet() {
	bindingEntryStructOnce.Do(func() {
		bindingEntryStruct = gi.StructNew("Gtk", "BindingEntry")
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

var bindingSetStruct *gi.Struct
var bindingSetStructOnce sync.Once

func bindingSetStructSet() {
	bindingSetStructOnce.Do(func() {
		bindingSetStruct = gi.StructNew("Gtk", "BindingSet")
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

var bindingSignalStruct *gi.Struct
var bindingSignalStructOnce sync.Once

func bindingSignalStructSet() {
	bindingSignalStructOnce.Do(func() {
		bindingSignalStruct = gi.StructNew("Gtk", "BindingSignal")
	})
}

type BindingSignal struct {
	native     uintptr
	Next       *BindingSignal
	SignalName string
	NArgs      uint32
	// UNSUPPORTED : C value 'args' : missing Type
}

var booleanCellAccessibleClassStruct *gi.Struct
var booleanCellAccessibleClassStructOnce sync.Once

func booleanCellAccessibleClassStructSet() {
	booleanCellAccessibleClassStructOnce.Do(func() {
		booleanCellAccessibleClassStruct = gi.StructNew("Gtk", "BooleanCellAccessibleClass")
	})
}

type BooleanCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

var booleanCellAccessiblePrivateStruct *gi.Struct
var booleanCellAccessiblePrivateStructOnce sync.Once

func booleanCellAccessiblePrivateStructSet() {
	booleanCellAccessiblePrivateStructOnce.Do(func() {
		booleanCellAccessiblePrivateStruct = gi.StructNew("Gtk", "BooleanCellAccessiblePrivate")
	})
}

type BooleanCellAccessiblePrivate struct {
	native uintptr
}

var borderStruct *gi.Struct
var borderStructOnce sync.Once

func borderStructSet() {
	borderStructOnce.Do(func() {
		borderStruct = gi.StructNew("Gtk", "Border")
	})
}

type Border struct {
	native uintptr
	Left   int16
	Right  int16
	Top    int16
	Bottom int16
}

var borderNewFunction *gi.Function
var borderNewFunction_Once sync.Once

func borderNewFunction_Set() {
	borderNewFunction_Once.Do(func() {
		borderNewFunction = gi.FunctionInvokerNew("Gtk", "new")
	})
}

var newBorderInvoker *gi.Function

// BorderNew is a representation of the C type gtk_border_new.
func BorderNew() *Border {
	borderNewFunction_Set()

	ret := borderNewFunction.Invoke(nil, nil)

	retGo := &Border{native: ret.Pointer()}

	return retGo
}

var borderCopyFunction *gi.Function
var borderCopyFunction_Once sync.Once

func borderCopyFunction_Set() {
	borderCopyFunction_Once.Do(func() {
		borderCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyBorderInvoker *gi.Function

// Copy is a representation of the C type gtk_border_copy.
func (recv *Border) Copy() *Border {
	borderCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := borderCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Border{native: ret.Pointer()}

	return retGo
}

var borderFreeFunction *gi.Function
var borderFreeFunction_Once sync.Once

func borderFreeFunction_Set() {
	borderFreeFunction_Once.Do(func() {
		borderFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeBorderInvoker *gi.Function

// Free is a representation of the C type gtk_border_free.
func (recv *Border) Free() {
	borderFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	borderFreeFunction.Invoke(inArgs[:], nil)

}

var boxClassStruct *gi.Struct
var boxClassStructOnce sync.Once

func boxClassStructSet() {
	boxClassStructOnce.Do(func() {
		boxClassStruct = gi.StructNew("Gtk", "BoxClass")
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

var boxPrivateStruct *gi.Struct
var boxPrivateStructOnce sync.Once

func boxPrivateStructSet() {
	boxPrivateStructOnce.Do(func() {
		boxPrivateStruct = gi.StructNew("Gtk", "BoxPrivate")
	})
}

type BoxPrivate struct {
	native uintptr
}

var buildableIfaceStruct *gi.Struct
var buildableIfaceStructOnce sync.Once

func buildableIfaceStructSet() {
	buildableIfaceStructOnce.Do(func() {
		buildableIfaceStruct = gi.StructNew("Gtk", "BuildableIface")
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

var builderClassStruct *gi.Struct
var builderClassStructOnce sync.Once

func builderClassStructSet() {
	builderClassStructOnce.Do(func() {
		builderClassStruct = gi.StructNew("Gtk", "BuilderClass")
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

var builderPrivateStruct *gi.Struct
var builderPrivateStructOnce sync.Once

func builderPrivateStructSet() {
	builderPrivateStructOnce.Do(func() {
		builderPrivateStruct = gi.StructNew("Gtk", "BuilderPrivate")
	})
}

type BuilderPrivate struct {
	native uintptr
}

var buttonAccessibleClassStruct *gi.Struct
var buttonAccessibleClassStructOnce sync.Once

func buttonAccessibleClassStructSet() {
	buttonAccessibleClassStructOnce.Do(func() {
		buttonAccessibleClassStruct = gi.StructNew("Gtk", "ButtonAccessibleClass")
	})
}

type ButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var buttonAccessiblePrivateStruct *gi.Struct
var buttonAccessiblePrivateStructOnce sync.Once

func buttonAccessiblePrivateStructSet() {
	buttonAccessiblePrivateStructOnce.Do(func() {
		buttonAccessiblePrivateStruct = gi.StructNew("Gtk", "ButtonAccessiblePrivate")
	})
}

type ButtonAccessiblePrivate struct {
	native uintptr
}

var buttonBoxClassStruct *gi.Struct
var buttonBoxClassStructOnce sync.Once

func buttonBoxClassStructSet() {
	buttonBoxClassStructOnce.Do(func() {
		buttonBoxClassStruct = gi.StructNew("Gtk", "ButtonBoxClass")
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

var buttonBoxPrivateStruct *gi.Struct
var buttonBoxPrivateStructOnce sync.Once

func buttonBoxPrivateStructSet() {
	buttonBoxPrivateStructOnce.Do(func() {
		buttonBoxPrivateStruct = gi.StructNew("Gtk", "ButtonBoxPrivate")
	})
}

type ButtonBoxPrivate struct {
	native uintptr
}

var buttonClassStruct *gi.Struct
var buttonClassStructOnce sync.Once

func buttonClassStructSet() {
	buttonClassStructOnce.Do(func() {
		buttonClassStruct = gi.StructNew("Gtk", "ButtonClass")
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

var buttonPrivateStruct *gi.Struct
var buttonPrivateStructOnce sync.Once

func buttonPrivateStructSet() {
	buttonPrivateStructOnce.Do(func() {
		buttonPrivateStruct = gi.StructNew("Gtk", "ButtonPrivate")
	})
}

type ButtonPrivate struct {
	native uintptr
}

var calendarClassStruct *gi.Struct
var calendarClassStructOnce sync.Once

func calendarClassStructSet() {
	calendarClassStructOnce.Do(func() {
		calendarClassStruct = gi.StructNew("Gtk", "CalendarClass")
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

var calendarPrivateStruct *gi.Struct
var calendarPrivateStructOnce sync.Once

func calendarPrivateStructSet() {
	calendarPrivateStructOnce.Do(func() {
		calendarPrivateStruct = gi.StructNew("Gtk", "CalendarPrivate")
	})
}

type CalendarPrivate struct {
	native uintptr
}

var cellAccessibleClassStruct *gi.Struct
var cellAccessibleClassStructOnce sync.Once

func cellAccessibleClassStructSet() {
	cellAccessibleClassStructOnce.Do(func() {
		cellAccessibleClassStruct = gi.StructNew("Gtk", "CellAccessibleClass")
	})
}

type CellAccessibleClass struct {
	native      uintptr
	ParentClass *AccessibleClass
	// UNSUPPORTED : C value 'update_cache' : missing Type
}

var cellAccessibleParentIfaceStruct *gi.Struct
var cellAccessibleParentIfaceStructOnce sync.Once

func cellAccessibleParentIfaceStructSet() {
	cellAccessibleParentIfaceStructOnce.Do(func() {
		cellAccessibleParentIfaceStruct = gi.StructNew("Gtk", "CellAccessibleParentIface")
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

var cellAccessiblePrivateStruct *gi.Struct
var cellAccessiblePrivateStructOnce sync.Once

func cellAccessiblePrivateStructSet() {
	cellAccessiblePrivateStructOnce.Do(func() {
		cellAccessiblePrivateStruct = gi.StructNew("Gtk", "CellAccessiblePrivate")
	})
}

type CellAccessiblePrivate struct {
	native uintptr
}

var cellAreaBoxClassStruct *gi.Struct
var cellAreaBoxClassStructOnce sync.Once

func cellAreaBoxClassStructSet() {
	cellAreaBoxClassStructOnce.Do(func() {
		cellAreaBoxClassStruct = gi.StructNew("Gtk", "CellAreaBoxClass")
	})
}

type CellAreaBoxClass struct {
	native uintptr
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var cellAreaBoxPrivateStruct *gi.Struct
var cellAreaBoxPrivateStructOnce sync.Once

func cellAreaBoxPrivateStructSet() {
	cellAreaBoxPrivateStructOnce.Do(func() {
		cellAreaBoxPrivateStruct = gi.StructNew("Gtk", "CellAreaBoxPrivate")
	})
}

type CellAreaBoxPrivate struct {
	native uintptr
}

var cellAreaClassStruct *gi.Struct
var cellAreaClassStructOnce sync.Once

func cellAreaClassStructSet() {
	cellAreaClassStructOnce.Do(func() {
		cellAreaClassStruct = gi.StructNew("Gtk", "CellAreaClass")
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

var cellAreaClassListCellPropertiesFunction *gi.Function
var cellAreaClassListCellPropertiesFunction_Once sync.Once

func cellAreaClassListCellPropertiesFunction_Set() {
	cellAreaClassListCellPropertiesFunction_Once.Do(func() {
		cellAreaClassListCellPropertiesFunction = gi.FunctionInvokerNew("Gtk", "list_cell_properties")
	})
}

var listCellPropertiesCellAreaClassInvoker *gi.Function

// ListCellProperties is a representation of the C type gtk_cell_area_class_list_cell_properties.
func (recv *CellAreaClass) ListCellProperties() uint32 {
	cellAreaClassListCellPropertiesFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	cellAreaClassListCellPropertiesFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint32()

	return out0
}

var cellAreaContextClassStruct *gi.Struct
var cellAreaContextClassStructOnce sync.Once

func cellAreaContextClassStructSet() {
	cellAreaContextClassStructOnce.Do(func() {
		cellAreaContextClassStruct = gi.StructNew("Gtk", "CellAreaContextClass")
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

var cellAreaContextPrivateStruct *gi.Struct
var cellAreaContextPrivateStructOnce sync.Once

func cellAreaContextPrivateStructSet() {
	cellAreaContextPrivateStructOnce.Do(func() {
		cellAreaContextPrivateStruct = gi.StructNew("Gtk", "CellAreaContextPrivate")
	})
}

type CellAreaContextPrivate struct {
	native uintptr
}

var cellAreaPrivateStruct *gi.Struct
var cellAreaPrivateStructOnce sync.Once

func cellAreaPrivateStructSet() {
	cellAreaPrivateStructOnce.Do(func() {
		cellAreaPrivateStruct = gi.StructNew("Gtk", "CellAreaPrivate")
	})
}

type CellAreaPrivate struct {
	native uintptr
}

var cellEditableIfaceStruct *gi.Struct
var cellEditableIfaceStructOnce sync.Once

func cellEditableIfaceStructSet() {
	cellEditableIfaceStructOnce.Do(func() {
		cellEditableIfaceStruct = gi.StructNew("Gtk", "CellEditableIface")
	})
}

type CellEditableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'editing_done' : missing Type
	// UNSUPPORTED : C value 'remove_widget' : missing Type
	// UNSUPPORTED : C value 'start_editing' : missing Type
}

var cellLayoutIfaceStruct *gi.Struct
var cellLayoutIfaceStructOnce sync.Once

func cellLayoutIfaceStructSet() {
	cellLayoutIfaceStructOnce.Do(func() {
		cellLayoutIfaceStruct = gi.StructNew("Gtk", "CellLayoutIface")
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

var cellRendererAccelClassStruct *gi.Struct
var cellRendererAccelClassStructOnce sync.Once

func cellRendererAccelClassStructSet() {
	cellRendererAccelClassStructOnce.Do(func() {
		cellRendererAccelClassStruct = gi.StructNew("Gtk", "CellRendererAccelClass")
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

var cellRendererAccelPrivateStruct *gi.Struct
var cellRendererAccelPrivateStructOnce sync.Once

func cellRendererAccelPrivateStructSet() {
	cellRendererAccelPrivateStructOnce.Do(func() {
		cellRendererAccelPrivateStruct = gi.StructNew("Gtk", "CellRendererAccelPrivate")
	})
}

type CellRendererAccelPrivate struct {
	native uintptr
}

var cellRendererClassStruct *gi.Struct
var cellRendererClassStructOnce sync.Once

func cellRendererClassStructSet() {
	cellRendererClassStructOnce.Do(func() {
		cellRendererClassStruct = gi.StructNew("Gtk", "CellRendererClass")
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

var cellRendererClassPrivateStruct *gi.Struct
var cellRendererClassPrivateStructOnce sync.Once

func cellRendererClassPrivateStructSet() {
	cellRendererClassPrivateStructOnce.Do(func() {
		cellRendererClassPrivateStruct = gi.StructNew("Gtk", "CellRendererClassPrivate")
	})
}

type CellRendererClassPrivate struct {
	native uintptr
}

var cellRendererComboClassStruct *gi.Struct
var cellRendererComboClassStructOnce sync.Once

func cellRendererComboClassStructSet() {
	cellRendererComboClassStructOnce.Do(func() {
		cellRendererComboClassStruct = gi.StructNew("Gtk", "CellRendererComboClass")
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

var cellRendererComboPrivateStruct *gi.Struct
var cellRendererComboPrivateStructOnce sync.Once

func cellRendererComboPrivateStructSet() {
	cellRendererComboPrivateStructOnce.Do(func() {
		cellRendererComboPrivateStruct = gi.StructNew("Gtk", "CellRendererComboPrivate")
	})
}

type CellRendererComboPrivate struct {
	native uintptr
}

var cellRendererPixbufClassStruct *gi.Struct
var cellRendererPixbufClassStructOnce sync.Once

func cellRendererPixbufClassStructSet() {
	cellRendererPixbufClassStructOnce.Do(func() {
		cellRendererPixbufClassStruct = gi.StructNew("Gtk", "CellRendererPixbufClass")
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

var cellRendererPixbufPrivateStruct *gi.Struct
var cellRendererPixbufPrivateStructOnce sync.Once

func cellRendererPixbufPrivateStructSet() {
	cellRendererPixbufPrivateStructOnce.Do(func() {
		cellRendererPixbufPrivateStruct = gi.StructNew("Gtk", "CellRendererPixbufPrivate")
	})
}

type CellRendererPixbufPrivate struct {
	native uintptr
}

var cellRendererPrivateStruct *gi.Struct
var cellRendererPrivateStructOnce sync.Once

func cellRendererPrivateStructSet() {
	cellRendererPrivateStructOnce.Do(func() {
		cellRendererPrivateStruct = gi.StructNew("Gtk", "CellRendererPrivate")
	})
}

type CellRendererPrivate struct {
	native uintptr
}

var cellRendererProgressClassStruct *gi.Struct
var cellRendererProgressClassStructOnce sync.Once

func cellRendererProgressClassStructSet() {
	cellRendererProgressClassStructOnce.Do(func() {
		cellRendererProgressClassStruct = gi.StructNew("Gtk", "CellRendererProgressClass")
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

var cellRendererProgressPrivateStruct *gi.Struct
var cellRendererProgressPrivateStructOnce sync.Once

func cellRendererProgressPrivateStructSet() {
	cellRendererProgressPrivateStructOnce.Do(func() {
		cellRendererProgressPrivateStruct = gi.StructNew("Gtk", "CellRendererProgressPrivate")
	})
}

type CellRendererProgressPrivate struct {
	native uintptr
}

var cellRendererSpinClassStruct *gi.Struct
var cellRendererSpinClassStructOnce sync.Once

func cellRendererSpinClassStructSet() {
	cellRendererSpinClassStructOnce.Do(func() {
		cellRendererSpinClassStruct = gi.StructNew("Gtk", "CellRendererSpinClass")
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

var cellRendererSpinPrivateStruct *gi.Struct
var cellRendererSpinPrivateStructOnce sync.Once

func cellRendererSpinPrivateStructSet() {
	cellRendererSpinPrivateStructOnce.Do(func() {
		cellRendererSpinPrivateStruct = gi.StructNew("Gtk", "CellRendererSpinPrivate")
	})
}

type CellRendererSpinPrivate struct {
	native uintptr
}

var cellRendererSpinnerClassStruct *gi.Struct
var cellRendererSpinnerClassStructOnce sync.Once

func cellRendererSpinnerClassStructSet() {
	cellRendererSpinnerClassStructOnce.Do(func() {
		cellRendererSpinnerClassStruct = gi.StructNew("Gtk", "CellRendererSpinnerClass")
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

var cellRendererSpinnerPrivateStruct *gi.Struct
var cellRendererSpinnerPrivateStructOnce sync.Once

func cellRendererSpinnerPrivateStructSet() {
	cellRendererSpinnerPrivateStructOnce.Do(func() {
		cellRendererSpinnerPrivateStruct = gi.StructNew("Gtk", "CellRendererSpinnerPrivate")
	})
}

type CellRendererSpinnerPrivate struct {
	native uintptr
}

var cellRendererTextClassStruct *gi.Struct
var cellRendererTextClassStructOnce sync.Once

func cellRendererTextClassStructSet() {
	cellRendererTextClassStructOnce.Do(func() {
		cellRendererTextClassStruct = gi.StructNew("Gtk", "CellRendererTextClass")
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

var cellRendererTextPrivateStruct *gi.Struct
var cellRendererTextPrivateStructOnce sync.Once

func cellRendererTextPrivateStructSet() {
	cellRendererTextPrivateStructOnce.Do(func() {
		cellRendererTextPrivateStruct = gi.StructNew("Gtk", "CellRendererTextPrivate")
	})
}

type CellRendererTextPrivate struct {
	native uintptr
}

var cellRendererToggleClassStruct *gi.Struct
var cellRendererToggleClassStructOnce sync.Once

func cellRendererToggleClassStructSet() {
	cellRendererToggleClassStructOnce.Do(func() {
		cellRendererToggleClassStruct = gi.StructNew("Gtk", "CellRendererToggleClass")
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

var cellRendererTogglePrivateStruct *gi.Struct
var cellRendererTogglePrivateStructOnce sync.Once

func cellRendererTogglePrivateStructSet() {
	cellRendererTogglePrivateStructOnce.Do(func() {
		cellRendererTogglePrivateStruct = gi.StructNew("Gtk", "CellRendererTogglePrivate")
	})
}

type CellRendererTogglePrivate struct {
	native uintptr
}

var cellViewClassStruct *gi.Struct
var cellViewClassStructOnce sync.Once

func cellViewClassStructSet() {
	cellViewClassStructOnce.Do(func() {
		cellViewClassStruct = gi.StructNew("Gtk", "CellViewClass")
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

var cellViewPrivateStruct *gi.Struct
var cellViewPrivateStructOnce sync.Once

func cellViewPrivateStructSet() {
	cellViewPrivateStructOnce.Do(func() {
		cellViewPrivateStruct = gi.StructNew("Gtk", "CellViewPrivate")
	})
}

type CellViewPrivate struct {
	native uintptr
}

var checkButtonClassStruct *gi.Struct
var checkButtonClassStructOnce sync.Once

func checkButtonClassStructSet() {
	checkButtonClassStructOnce.Do(func() {
		checkButtonClassStruct = gi.StructNew("Gtk", "CheckButtonClass")
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

var checkMenuItemAccessibleClassStruct *gi.Struct
var checkMenuItemAccessibleClassStructOnce sync.Once

func checkMenuItemAccessibleClassStructSet() {
	checkMenuItemAccessibleClassStructOnce.Do(func() {
		checkMenuItemAccessibleClassStruct = gi.StructNew("Gtk", "CheckMenuItemAccessibleClass")
	})
}

type CheckMenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *MenuItemAccessibleClass
}

var checkMenuItemAccessiblePrivateStruct *gi.Struct
var checkMenuItemAccessiblePrivateStructOnce sync.Once

func checkMenuItemAccessiblePrivateStructSet() {
	checkMenuItemAccessiblePrivateStructOnce.Do(func() {
		checkMenuItemAccessiblePrivateStruct = gi.StructNew("Gtk", "CheckMenuItemAccessiblePrivate")
	})
}

type CheckMenuItemAccessiblePrivate struct {
	native uintptr
}

var checkMenuItemClassStruct *gi.Struct
var checkMenuItemClassStructOnce sync.Once

func checkMenuItemClassStructSet() {
	checkMenuItemClassStructOnce.Do(func() {
		checkMenuItemClassStruct = gi.StructNew("Gtk", "CheckMenuItemClass")
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

var checkMenuItemPrivateStruct *gi.Struct
var checkMenuItemPrivateStructOnce sync.Once

func checkMenuItemPrivateStructSet() {
	checkMenuItemPrivateStructOnce.Do(func() {
		checkMenuItemPrivateStruct = gi.StructNew("Gtk", "CheckMenuItemPrivate")
	})
}

type CheckMenuItemPrivate struct {
	native uintptr
}

var colorButtonClassStruct *gi.Struct
var colorButtonClassStructOnce sync.Once

func colorButtonClassStructSet() {
	colorButtonClassStructOnce.Do(func() {
		colorButtonClassStruct = gi.StructNew("Gtk", "ColorButtonClass")
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

var colorButtonPrivateStruct *gi.Struct
var colorButtonPrivateStructOnce sync.Once

func colorButtonPrivateStructSet() {
	colorButtonPrivateStructOnce.Do(func() {
		colorButtonPrivateStruct = gi.StructNew("Gtk", "ColorButtonPrivate")
	})
}

type ColorButtonPrivate struct {
	native uintptr
}

var colorChooserDialogClassStruct *gi.Struct
var colorChooserDialogClassStructOnce sync.Once

func colorChooserDialogClassStructSet() {
	colorChooserDialogClassStructOnce.Do(func() {
		colorChooserDialogClassStruct = gi.StructNew("Gtk", "ColorChooserDialogClass")
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

var colorChooserDialogPrivateStruct *gi.Struct
var colorChooserDialogPrivateStructOnce sync.Once

func colorChooserDialogPrivateStructSet() {
	colorChooserDialogPrivateStructOnce.Do(func() {
		colorChooserDialogPrivateStruct = gi.StructNew("Gtk", "ColorChooserDialogPrivate")
	})
}

type ColorChooserDialogPrivate struct {
	native uintptr
}

var colorChooserInterfaceStruct *gi.Struct
var colorChooserInterfaceStructOnce sync.Once

func colorChooserInterfaceStructSet() {
	colorChooserInterfaceStructOnce.Do(func() {
		colorChooserInterfaceStruct = gi.StructNew("Gtk", "ColorChooserInterface")
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

var colorChooserWidgetClassStruct *gi.Struct
var colorChooserWidgetClassStructOnce sync.Once

func colorChooserWidgetClassStructSet() {
	colorChooserWidgetClassStructOnce.Do(func() {
		colorChooserWidgetClassStruct = gi.StructNew("Gtk", "ColorChooserWidgetClass")
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

var colorChooserWidgetPrivateStruct *gi.Struct
var colorChooserWidgetPrivateStructOnce sync.Once

func colorChooserWidgetPrivateStructSet() {
	colorChooserWidgetPrivateStructOnce.Do(func() {
		colorChooserWidgetPrivateStruct = gi.StructNew("Gtk", "ColorChooserWidgetPrivate")
	})
}

type ColorChooserWidgetPrivate struct {
	native uintptr
}

var colorSelectionClassStruct *gi.Struct
var colorSelectionClassStructOnce sync.Once

func colorSelectionClassStructSet() {
	colorSelectionClassStructOnce.Do(func() {
		colorSelectionClassStruct = gi.StructNew("Gtk", "ColorSelectionClass")
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

var colorSelectionDialogClassStruct *gi.Struct
var colorSelectionDialogClassStructOnce sync.Once

func colorSelectionDialogClassStructSet() {
	colorSelectionDialogClassStructOnce.Do(func() {
		colorSelectionDialogClassStruct = gi.StructNew("Gtk", "ColorSelectionDialogClass")
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

var colorSelectionDialogPrivateStruct *gi.Struct
var colorSelectionDialogPrivateStructOnce sync.Once

func colorSelectionDialogPrivateStructSet() {
	colorSelectionDialogPrivateStructOnce.Do(func() {
		colorSelectionDialogPrivateStruct = gi.StructNew("Gtk", "ColorSelectionDialogPrivate")
	})
}

type ColorSelectionDialogPrivate struct {
	native uintptr
}

var colorSelectionPrivateStruct *gi.Struct
var colorSelectionPrivateStructOnce sync.Once

func colorSelectionPrivateStructSet() {
	colorSelectionPrivateStructOnce.Do(func() {
		colorSelectionPrivateStruct = gi.StructNew("Gtk", "ColorSelectionPrivate")
	})
}

type ColorSelectionPrivate struct {
	native uintptr
}

var comboBoxAccessibleClassStruct *gi.Struct
var comboBoxAccessibleClassStructOnce sync.Once

func comboBoxAccessibleClassStructSet() {
	comboBoxAccessibleClassStructOnce.Do(func() {
		comboBoxAccessibleClassStruct = gi.StructNew("Gtk", "ComboBoxAccessibleClass")
	})
}

type ComboBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var comboBoxAccessiblePrivateStruct *gi.Struct
var comboBoxAccessiblePrivateStructOnce sync.Once

func comboBoxAccessiblePrivateStructSet() {
	comboBoxAccessiblePrivateStructOnce.Do(func() {
		comboBoxAccessiblePrivateStruct = gi.StructNew("Gtk", "ComboBoxAccessiblePrivate")
	})
}

type ComboBoxAccessiblePrivate struct {
	native uintptr
}

var comboBoxClassStruct *gi.Struct
var comboBoxClassStructOnce sync.Once

func comboBoxClassStructSet() {
	comboBoxClassStructOnce.Do(func() {
		comboBoxClassStruct = gi.StructNew("Gtk", "ComboBoxClass")
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

var comboBoxPrivateStruct *gi.Struct
var comboBoxPrivateStructOnce sync.Once

func comboBoxPrivateStructSet() {
	comboBoxPrivateStructOnce.Do(func() {
		comboBoxPrivateStruct = gi.StructNew("Gtk", "ComboBoxPrivate")
	})
}

type ComboBoxPrivate struct {
	native uintptr
}

var comboBoxTextClassStruct *gi.Struct
var comboBoxTextClassStructOnce sync.Once

func comboBoxTextClassStructSet() {
	comboBoxTextClassStructOnce.Do(func() {
		comboBoxTextClassStruct = gi.StructNew("Gtk", "ComboBoxTextClass")
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

var comboBoxTextPrivateStruct *gi.Struct
var comboBoxTextPrivateStructOnce sync.Once

func comboBoxTextPrivateStructSet() {
	comboBoxTextPrivateStructOnce.Do(func() {
		comboBoxTextPrivateStruct = gi.StructNew("Gtk", "ComboBoxTextPrivate")
	})
}

type ComboBoxTextPrivate struct {
	native uintptr
}

var containerAccessibleClassStruct *gi.Struct
var containerAccessibleClassStructOnce sync.Once

func containerAccessibleClassStructSet() {
	containerAccessibleClassStructOnce.Do(func() {
		containerAccessibleClassStruct = gi.StructNew("Gtk", "ContainerAccessibleClass")
	})
}

type ContainerAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
	// UNSUPPORTED : C value 'add_gtk' : missing Type
	// UNSUPPORTED : C value 'remove_gtk' : missing Type
}

var containerAccessiblePrivateStruct *gi.Struct
var containerAccessiblePrivateStructOnce sync.Once

func containerAccessiblePrivateStructSet() {
	containerAccessiblePrivateStructOnce.Do(func() {
		containerAccessiblePrivateStruct = gi.StructNew("Gtk", "ContainerAccessiblePrivate")
	})
}

type ContainerAccessiblePrivate struct {
	native uintptr
}

var containerCellAccessibleClassStruct *gi.Struct
var containerCellAccessibleClassStructOnce sync.Once

func containerCellAccessibleClassStructSet() {
	containerCellAccessibleClassStructOnce.Do(func() {
		containerCellAccessibleClassStruct = gi.StructNew("Gtk", "ContainerCellAccessibleClass")
	})
}

type ContainerCellAccessibleClass struct {
	native      uintptr
	ParentClass *CellAccessibleClass
}

var containerCellAccessiblePrivateStruct *gi.Struct
var containerCellAccessiblePrivateStructOnce sync.Once

func containerCellAccessiblePrivateStructSet() {
	containerCellAccessiblePrivateStructOnce.Do(func() {
		containerCellAccessiblePrivateStruct = gi.StructNew("Gtk", "ContainerCellAccessiblePrivate")
	})
}

type ContainerCellAccessiblePrivate struct {
	native uintptr
}

var containerClassStruct *gi.Struct
var containerClassStructOnce sync.Once

func containerClassStructSet() {
	containerClassStructOnce.Do(func() {
		containerClassStruct = gi.StructNew("Gtk", "ContainerClass")
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

var containerClassHandleBorderWidthFunction *gi.Function
var containerClassHandleBorderWidthFunction_Once sync.Once

func containerClassHandleBorderWidthFunction_Set() {
	containerClassHandleBorderWidthFunction_Once.Do(func() {
		containerClassHandleBorderWidthFunction = gi.FunctionInvokerNew("Gtk", "handle_border_width")
	})
}

var handleBorderWidthContainerClassInvoker *gi.Function

// HandleBorderWidth is a representation of the C type gtk_container_class_handle_border_width.
func (recv *ContainerClass) HandleBorderWidth() {
	containerClassHandleBorderWidthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	containerClassHandleBorderWidthFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_container_class_install_child_properties' : parameter 'pspecs' has no type

// UNSUPPORTED : C value 'gtk_container_class_install_child_property' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

var containerClassListChildPropertiesFunction *gi.Function
var containerClassListChildPropertiesFunction_Once sync.Once

func containerClassListChildPropertiesFunction_Set() {
	containerClassListChildPropertiesFunction_Once.Do(func() {
		containerClassListChildPropertiesFunction = gi.FunctionInvokerNew("Gtk", "list_child_properties")
	})
}

var listChildPropertiesContainerClassInvoker *gi.Function

// ListChildProperties is a representation of the C type gtk_container_class_list_child_properties.
func (recv *ContainerClass) ListChildProperties() uint32 {
	containerClassListChildPropertiesFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	containerClassListChildPropertiesFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint32()

	return out0
}

var containerPrivateStruct *gi.Struct
var containerPrivateStructOnce sync.Once

func containerPrivateStructSet() {
	containerPrivateStructOnce.Do(func() {
		containerPrivateStruct = gi.StructNew("Gtk", "ContainerPrivate")
	})
}

type ContainerPrivate struct {
	native uintptr
}

var cssProviderClassStruct *gi.Struct
var cssProviderClassStructOnce sync.Once

func cssProviderClassStructSet() {
	cssProviderClassStructOnce.Do(func() {
		cssProviderClassStruct = gi.StructNew("Gtk", "CssProviderClass")
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

var cssProviderPrivateStruct *gi.Struct
var cssProviderPrivateStructOnce sync.Once

func cssProviderPrivateStructSet() {
	cssProviderPrivateStructOnce.Do(func() {
		cssProviderPrivateStruct = gi.StructNew("Gtk", "CssProviderPrivate")
	})
}

type CssProviderPrivate struct {
	native uintptr
}

var cssSectionStruct *gi.Struct
var cssSectionStructOnce sync.Once

func cssSectionStructSet() {
	cssSectionStructOnce.Do(func() {
		cssSectionStruct = gi.StructNew("Gtk", "CssSection")
	})
}

type CssSection struct {
	native uintptr
}

var cssSectionGetEndLineFunction *gi.Function
var cssSectionGetEndLineFunction_Once sync.Once

func cssSectionGetEndLineFunction_Set() {
	cssSectionGetEndLineFunction_Once.Do(func() {
		cssSectionGetEndLineFunction = gi.FunctionInvokerNew("Gtk", "get_end_line")
	})
}

var getEndLineCssSectionInvoker *gi.Function

// GetEndLine is a representation of the C type gtk_css_section_get_end_line.
func (recv *CssSection) GetEndLine() uint32 {
	cssSectionGetEndLineFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cssSectionGetEndLineFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var cssSectionGetEndPositionFunction *gi.Function
var cssSectionGetEndPositionFunction_Once sync.Once

func cssSectionGetEndPositionFunction_Set() {
	cssSectionGetEndPositionFunction_Once.Do(func() {
		cssSectionGetEndPositionFunction = gi.FunctionInvokerNew("Gtk", "get_end_position")
	})
}

var getEndPositionCssSectionInvoker *gi.Function

// GetEndPosition is a representation of the C type gtk_css_section_get_end_position.
func (recv *CssSection) GetEndPosition() uint32 {
	cssSectionGetEndPositionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cssSectionGetEndPositionFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_css_section_get_file' : return type 'Gio.File' not supported

var cssSectionGetParentFunction *gi.Function
var cssSectionGetParentFunction_Once sync.Once

func cssSectionGetParentFunction_Set() {
	cssSectionGetParentFunction_Once.Do(func() {
		cssSectionGetParentFunction = gi.FunctionInvokerNew("Gtk", "get_parent")
	})
}

var getParentCssSectionInvoker *gi.Function

// GetParent is a representation of the C type gtk_css_section_get_parent.
func (recv *CssSection) GetParent() *CssSection {
	cssSectionGetParentFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cssSectionGetParentFunction.Invoke(inArgs[:], nil)

	retGo := &CssSection{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_css_section_get_section_type' : return type 'CssSectionType' not supported

var cssSectionGetStartLineFunction *gi.Function
var cssSectionGetStartLineFunction_Once sync.Once

func cssSectionGetStartLineFunction_Set() {
	cssSectionGetStartLineFunction_Once.Do(func() {
		cssSectionGetStartLineFunction = gi.FunctionInvokerNew("Gtk", "get_start_line")
	})
}

var getStartLineCssSectionInvoker *gi.Function

// GetStartLine is a representation of the C type gtk_css_section_get_start_line.
func (recv *CssSection) GetStartLine() uint32 {
	cssSectionGetStartLineFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cssSectionGetStartLineFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var cssSectionGetStartPositionFunction *gi.Function
var cssSectionGetStartPositionFunction_Once sync.Once

func cssSectionGetStartPositionFunction_Set() {
	cssSectionGetStartPositionFunction_Once.Do(func() {
		cssSectionGetStartPositionFunction = gi.FunctionInvokerNew("Gtk", "get_start_position")
	})
}

var getStartPositionCssSectionInvoker *gi.Function

// GetStartPosition is a representation of the C type gtk_css_section_get_start_position.
func (recv *CssSection) GetStartPosition() uint32 {
	cssSectionGetStartPositionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cssSectionGetStartPositionFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var cssSectionRefFunction *gi.Function
var cssSectionRefFunction_Once sync.Once

func cssSectionRefFunction_Set() {
	cssSectionRefFunction_Once.Do(func() {
		cssSectionRefFunction = gi.FunctionInvokerNew("Gtk", "ref")
	})
}

var refCssSectionInvoker *gi.Function

// Ref is a representation of the C type gtk_css_section_ref.
func (recv *CssSection) Ref() *CssSection {
	cssSectionRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := cssSectionRefFunction.Invoke(inArgs[:], nil)

	retGo := &CssSection{native: ret.Pointer()}

	return retGo
}

var cssSectionUnrefFunction *gi.Function
var cssSectionUnrefFunction_Once sync.Once

func cssSectionUnrefFunction_Set() {
	cssSectionUnrefFunction_Once.Do(func() {
		cssSectionUnrefFunction = gi.FunctionInvokerNew("Gtk", "unref")
	})
}

var unrefCssSectionInvoker *gi.Function

// Unref is a representation of the C type gtk_css_section_unref.
func (recv *CssSection) Unref() {
	cssSectionUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	cssSectionUnrefFunction.Invoke(inArgs[:], nil)

}

var dialogClassStruct *gi.Struct
var dialogClassStructOnce sync.Once

func dialogClassStructSet() {
	dialogClassStructOnce.Do(func() {
		dialogClassStruct = gi.StructNew("Gtk", "DialogClass")
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

var dialogPrivateStruct *gi.Struct
var dialogPrivateStructOnce sync.Once

func dialogPrivateStructSet() {
	dialogPrivateStructOnce.Do(func() {
		dialogPrivateStruct = gi.StructNew("Gtk", "DialogPrivate")
	})
}

type DialogPrivate struct {
	native uintptr
}

var drawingAreaClassStruct *gi.Struct
var drawingAreaClassStructOnce sync.Once

func drawingAreaClassStructSet() {
	drawingAreaClassStructOnce.Do(func() {
		drawingAreaClassStruct = gi.StructNew("Gtk", "DrawingAreaClass")
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

var editableInterfaceStruct *gi.Struct
var editableInterfaceStructOnce sync.Once

func editableInterfaceStructSet() {
	editableInterfaceStructOnce.Do(func() {
		editableInterfaceStruct = gi.StructNew("Gtk", "EditableInterface")
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

var entryAccessibleClassStruct *gi.Struct
var entryAccessibleClassStructOnce sync.Once

func entryAccessibleClassStructSet() {
	entryAccessibleClassStructOnce.Do(func() {
		entryAccessibleClassStruct = gi.StructNew("Gtk", "EntryAccessibleClass")
	})
}

type EntryAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var entryAccessiblePrivateStruct *gi.Struct
var entryAccessiblePrivateStructOnce sync.Once

func entryAccessiblePrivateStructSet() {
	entryAccessiblePrivateStructOnce.Do(func() {
		entryAccessiblePrivateStruct = gi.StructNew("Gtk", "EntryAccessiblePrivate")
	})
}

type EntryAccessiblePrivate struct {
	native uintptr
}

var entryBufferClassStruct *gi.Struct
var entryBufferClassStructOnce sync.Once

func entryBufferClassStructSet() {
	entryBufferClassStructOnce.Do(func() {
		entryBufferClassStruct = gi.StructNew("Gtk", "EntryBufferClass")
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

var entryBufferPrivateStruct *gi.Struct
var entryBufferPrivateStructOnce sync.Once

func entryBufferPrivateStructSet() {
	entryBufferPrivateStructOnce.Do(func() {
		entryBufferPrivateStruct = gi.StructNew("Gtk", "EntryBufferPrivate")
	})
}

type EntryBufferPrivate struct {
	native uintptr
}

var entryClassStruct *gi.Struct
var entryClassStructOnce sync.Once

func entryClassStructSet() {
	entryClassStructOnce.Do(func() {
		entryClassStruct = gi.StructNew("Gtk", "EntryClass")
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

var entryCompletionClassStruct *gi.Struct
var entryCompletionClassStructOnce sync.Once

func entryCompletionClassStructSet() {
	entryCompletionClassStructOnce.Do(func() {
		entryCompletionClassStruct = gi.StructNew("Gtk", "EntryCompletionClass")
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

var entryCompletionPrivateStruct *gi.Struct
var entryCompletionPrivateStructOnce sync.Once

func entryCompletionPrivateStructSet() {
	entryCompletionPrivateStructOnce.Do(func() {
		entryCompletionPrivateStruct = gi.StructNew("Gtk", "EntryCompletionPrivate")
	})
}

type EntryCompletionPrivate struct {
	native uintptr
}

var entryPrivateStruct *gi.Struct
var entryPrivateStructOnce sync.Once

func entryPrivateStructSet() {
	entryPrivateStructOnce.Do(func() {
		entryPrivateStruct = gi.StructNew("Gtk", "EntryPrivate")
	})
}

type EntryPrivate struct {
	native uintptr
}

var eventBoxClassStruct *gi.Struct
var eventBoxClassStructOnce sync.Once

func eventBoxClassStructSet() {
	eventBoxClassStructOnce.Do(func() {
		eventBoxClassStruct = gi.StructNew("Gtk", "EventBoxClass")
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

var eventBoxPrivateStruct *gi.Struct
var eventBoxPrivateStructOnce sync.Once

func eventBoxPrivateStructSet() {
	eventBoxPrivateStructOnce.Do(func() {
		eventBoxPrivateStruct = gi.StructNew("Gtk", "EventBoxPrivate")
	})
}

type EventBoxPrivate struct {
	native uintptr
}

var eventControllerClassStruct *gi.Struct
var eventControllerClassStructOnce sync.Once

func eventControllerClassStructSet() {
	eventControllerClassStructOnce.Do(func() {
		eventControllerClassStruct = gi.StructNew("Gtk", "EventControllerClass")
	})
}

type EventControllerClass struct {
	native uintptr
}

var eventControllerKeyClassStruct *gi.Struct
var eventControllerKeyClassStructOnce sync.Once

func eventControllerKeyClassStructSet() {
	eventControllerKeyClassStructOnce.Do(func() {
		eventControllerKeyClassStruct = gi.StructNew("Gtk", "EventControllerKeyClass")
	})
}

type EventControllerKeyClass struct {
	native uintptr
}

var eventControllerMotionClassStruct *gi.Struct
var eventControllerMotionClassStructOnce sync.Once

func eventControllerMotionClassStructSet() {
	eventControllerMotionClassStructOnce.Do(func() {
		eventControllerMotionClassStruct = gi.StructNew("Gtk", "EventControllerMotionClass")
	})
}

type EventControllerMotionClass struct {
	native uintptr
}

var eventControllerScrollClassStruct *gi.Struct
var eventControllerScrollClassStructOnce sync.Once

func eventControllerScrollClassStructSet() {
	eventControllerScrollClassStructOnce.Do(func() {
		eventControllerScrollClassStruct = gi.StructNew("Gtk", "EventControllerScrollClass")
	})
}

type EventControllerScrollClass struct {
	native uintptr
}

var expanderAccessibleClassStruct *gi.Struct
var expanderAccessibleClassStructOnce sync.Once

func expanderAccessibleClassStructSet() {
	expanderAccessibleClassStructOnce.Do(func() {
		expanderAccessibleClassStruct = gi.StructNew("Gtk", "ExpanderAccessibleClass")
	})
}

type ExpanderAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var expanderAccessiblePrivateStruct *gi.Struct
var expanderAccessiblePrivateStructOnce sync.Once

func expanderAccessiblePrivateStructSet() {
	expanderAccessiblePrivateStructOnce.Do(func() {
		expanderAccessiblePrivateStruct = gi.StructNew("Gtk", "ExpanderAccessiblePrivate")
	})
}

type ExpanderAccessiblePrivate struct {
	native uintptr
}

var expanderClassStruct *gi.Struct
var expanderClassStructOnce sync.Once

func expanderClassStructSet() {
	expanderClassStructOnce.Do(func() {
		expanderClassStruct = gi.StructNew("Gtk", "ExpanderClass")
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

var expanderPrivateStruct *gi.Struct
var expanderPrivateStructOnce sync.Once

func expanderPrivateStructSet() {
	expanderPrivateStructOnce.Do(func() {
		expanderPrivateStruct = gi.StructNew("Gtk", "ExpanderPrivate")
	})
}

type ExpanderPrivate struct {
	native uintptr
}

var fileChooserButtonClassStruct *gi.Struct
var fileChooserButtonClassStructOnce sync.Once

func fileChooserButtonClassStructSet() {
	fileChooserButtonClassStructOnce.Do(func() {
		fileChooserButtonClassStruct = gi.StructNew("Gtk", "FileChooserButtonClass")
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

var fileChooserButtonPrivateStruct *gi.Struct
var fileChooserButtonPrivateStructOnce sync.Once

func fileChooserButtonPrivateStructSet() {
	fileChooserButtonPrivateStructOnce.Do(func() {
		fileChooserButtonPrivateStruct = gi.StructNew("Gtk", "FileChooserButtonPrivate")
	})
}

type FileChooserButtonPrivate struct {
	native uintptr
}

var fileChooserDialogClassStruct *gi.Struct
var fileChooserDialogClassStructOnce sync.Once

func fileChooserDialogClassStructSet() {
	fileChooserDialogClassStructOnce.Do(func() {
		fileChooserDialogClassStruct = gi.StructNew("Gtk", "FileChooserDialogClass")
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

var fileChooserDialogPrivateStruct *gi.Struct
var fileChooserDialogPrivateStructOnce sync.Once

func fileChooserDialogPrivateStructSet() {
	fileChooserDialogPrivateStructOnce.Do(func() {
		fileChooserDialogPrivateStruct = gi.StructNew("Gtk", "FileChooserDialogPrivate")
	})
}

type FileChooserDialogPrivate struct {
	native uintptr
}

var fileChooserNativeClassStruct *gi.Struct
var fileChooserNativeClassStructOnce sync.Once

func fileChooserNativeClassStructSet() {
	fileChooserNativeClassStructOnce.Do(func() {
		fileChooserNativeClassStruct = gi.StructNew("Gtk", "FileChooserNativeClass")
	})
}

type FileChooserNativeClass struct {
	native      uintptr
	ParentClass *NativeDialogClass
}

var fileChooserWidgetClassStruct *gi.Struct
var fileChooserWidgetClassStructOnce sync.Once

func fileChooserWidgetClassStructSet() {
	fileChooserWidgetClassStructOnce.Do(func() {
		fileChooserWidgetClassStruct = gi.StructNew("Gtk", "FileChooserWidgetClass")
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

var fileChooserWidgetPrivateStruct *gi.Struct
var fileChooserWidgetPrivateStructOnce sync.Once

func fileChooserWidgetPrivateStructSet() {
	fileChooserWidgetPrivateStructOnce.Do(func() {
		fileChooserWidgetPrivateStruct = gi.StructNew("Gtk", "FileChooserWidgetPrivate")
	})
}

type FileChooserWidgetPrivate struct {
	native uintptr
}

var fileFilterInfoStruct *gi.Struct
var fileFilterInfoStructOnce sync.Once

func fileFilterInfoStructSet() {
	fileFilterInfoStructOnce.Do(func() {
		fileFilterInfoStruct = gi.StructNew("Gtk", "FileFilterInfo")
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

var fixedChildStruct *gi.Struct
var fixedChildStructOnce sync.Once

func fixedChildStructSet() {
	fixedChildStructOnce.Do(func() {
		fixedChildStruct = gi.StructNew("Gtk", "FixedChild")
	})
}

type FixedChild struct {
	native uintptr
	// UNSUPPORTED : C value 'widget' : no Go type for 'Widget'
	X int32
	Y int32
}

var fixedClassStruct *gi.Struct
var fixedClassStructOnce sync.Once

func fixedClassStructSet() {
	fixedClassStructOnce.Do(func() {
		fixedClassStruct = gi.StructNew("Gtk", "FixedClass")
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

var fixedPrivateStruct *gi.Struct
var fixedPrivateStructOnce sync.Once

func fixedPrivateStructSet() {
	fixedPrivateStructOnce.Do(func() {
		fixedPrivateStruct = gi.StructNew("Gtk", "FixedPrivate")
	})
}

type FixedPrivate struct {
	native uintptr
}

var flowBoxAccessibleClassStruct *gi.Struct
var flowBoxAccessibleClassStructOnce sync.Once

func flowBoxAccessibleClassStructSet() {
	flowBoxAccessibleClassStructOnce.Do(func() {
		flowBoxAccessibleClassStruct = gi.StructNew("Gtk", "FlowBoxAccessibleClass")
	})
}

type FlowBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var flowBoxAccessiblePrivateStruct *gi.Struct
var flowBoxAccessiblePrivateStructOnce sync.Once

func flowBoxAccessiblePrivateStructSet() {
	flowBoxAccessiblePrivateStructOnce.Do(func() {
		flowBoxAccessiblePrivateStruct = gi.StructNew("Gtk", "FlowBoxAccessiblePrivate")
	})
}

type FlowBoxAccessiblePrivate struct {
	native uintptr
}

var flowBoxChildAccessibleClassStruct *gi.Struct
var flowBoxChildAccessibleClassStructOnce sync.Once

func flowBoxChildAccessibleClassStructSet() {
	flowBoxChildAccessibleClassStructOnce.Do(func() {
		flowBoxChildAccessibleClassStruct = gi.StructNew("Gtk", "FlowBoxChildAccessibleClass")
	})
}

type FlowBoxChildAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var flowBoxChildClassStruct *gi.Struct
var flowBoxChildClassStructOnce sync.Once

func flowBoxChildClassStructSet() {
	flowBoxChildClassStructOnce.Do(func() {
		flowBoxChildClassStruct = gi.StructNew("Gtk", "FlowBoxChildClass")
	})
}

type FlowBoxChildClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'activate' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
}

var flowBoxClassStruct *gi.Struct
var flowBoxClassStructOnce sync.Once

func flowBoxClassStructSet() {
	flowBoxClassStructOnce.Do(func() {
		flowBoxClassStruct = gi.StructNew("Gtk", "FlowBoxClass")
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

var fontButtonClassStruct *gi.Struct
var fontButtonClassStructOnce sync.Once

func fontButtonClassStructSet() {
	fontButtonClassStructOnce.Do(func() {
		fontButtonClassStruct = gi.StructNew("Gtk", "FontButtonClass")
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

var fontButtonPrivateStruct *gi.Struct
var fontButtonPrivateStructOnce sync.Once

func fontButtonPrivateStructSet() {
	fontButtonPrivateStructOnce.Do(func() {
		fontButtonPrivateStruct = gi.StructNew("Gtk", "FontButtonPrivate")
	})
}

type FontButtonPrivate struct {
	native uintptr
}

var fontChooserDialogClassStruct *gi.Struct
var fontChooserDialogClassStructOnce sync.Once

func fontChooserDialogClassStructSet() {
	fontChooserDialogClassStructOnce.Do(func() {
		fontChooserDialogClassStruct = gi.StructNew("Gtk", "FontChooserDialogClass")
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

var fontChooserDialogPrivateStruct *gi.Struct
var fontChooserDialogPrivateStructOnce sync.Once

func fontChooserDialogPrivateStructSet() {
	fontChooserDialogPrivateStructOnce.Do(func() {
		fontChooserDialogPrivateStruct = gi.StructNew("Gtk", "FontChooserDialogPrivate")
	})
}

type FontChooserDialogPrivate struct {
	native uintptr
}

var fontChooserIfaceStruct *gi.Struct
var fontChooserIfaceStructOnce sync.Once

func fontChooserIfaceStructSet() {
	fontChooserIfaceStructOnce.Do(func() {
		fontChooserIfaceStruct = gi.StructNew("Gtk", "FontChooserIface")
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

var fontChooserWidgetClassStruct *gi.Struct
var fontChooserWidgetClassStructOnce sync.Once

func fontChooserWidgetClassStructSet() {
	fontChooserWidgetClassStructOnce.Do(func() {
		fontChooserWidgetClassStruct = gi.StructNew("Gtk", "FontChooserWidgetClass")
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

var fontChooserWidgetPrivateStruct *gi.Struct
var fontChooserWidgetPrivateStructOnce sync.Once

func fontChooserWidgetPrivateStructSet() {
	fontChooserWidgetPrivateStructOnce.Do(func() {
		fontChooserWidgetPrivateStruct = gi.StructNew("Gtk", "FontChooserWidgetPrivate")
	})
}

type FontChooserWidgetPrivate struct {
	native uintptr
}

var fontSelectionClassStruct *gi.Struct
var fontSelectionClassStructOnce sync.Once

func fontSelectionClassStructSet() {
	fontSelectionClassStructOnce.Do(func() {
		fontSelectionClassStruct = gi.StructNew("Gtk", "FontSelectionClass")
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

var fontSelectionDialogClassStruct *gi.Struct
var fontSelectionDialogClassStructOnce sync.Once

func fontSelectionDialogClassStructSet() {
	fontSelectionDialogClassStructOnce.Do(func() {
		fontSelectionDialogClassStruct = gi.StructNew("Gtk", "FontSelectionDialogClass")
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

var fontSelectionDialogPrivateStruct *gi.Struct
var fontSelectionDialogPrivateStructOnce sync.Once

func fontSelectionDialogPrivateStructSet() {
	fontSelectionDialogPrivateStructOnce.Do(func() {
		fontSelectionDialogPrivateStruct = gi.StructNew("Gtk", "FontSelectionDialogPrivate")
	})
}

type FontSelectionDialogPrivate struct {
	native uintptr
}

var fontSelectionPrivateStruct *gi.Struct
var fontSelectionPrivateStructOnce sync.Once

func fontSelectionPrivateStructSet() {
	fontSelectionPrivateStructOnce.Do(func() {
		fontSelectionPrivateStruct = gi.StructNew("Gtk", "FontSelectionPrivate")
	})
}

type FontSelectionPrivate struct {
	native uintptr
}

var frameAccessibleClassStruct *gi.Struct
var frameAccessibleClassStructOnce sync.Once

func frameAccessibleClassStructSet() {
	frameAccessibleClassStructOnce.Do(func() {
		frameAccessibleClassStruct = gi.StructNew("Gtk", "FrameAccessibleClass")
	})
}

type FrameAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var frameAccessiblePrivateStruct *gi.Struct
var frameAccessiblePrivateStructOnce sync.Once

func frameAccessiblePrivateStructSet() {
	frameAccessiblePrivateStructOnce.Do(func() {
		frameAccessiblePrivateStruct = gi.StructNew("Gtk", "FrameAccessiblePrivate")
	})
}

type FrameAccessiblePrivate struct {
	native uintptr
}

var frameClassStruct *gi.Struct
var frameClassStructOnce sync.Once

func frameClassStructSet() {
	frameClassStructOnce.Do(func() {
		frameClassStruct = gi.StructNew("Gtk", "FrameClass")
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

var framePrivateStruct *gi.Struct
var framePrivateStructOnce sync.Once

func framePrivateStructSet() {
	framePrivateStructOnce.Do(func() {
		framePrivateStruct = gi.StructNew("Gtk", "FramePrivate")
	})
}

type FramePrivate struct {
	native uintptr
}

var gLAreaClassStruct *gi.Struct
var gLAreaClassStructOnce sync.Once

func gLAreaClassStructSet() {
	gLAreaClassStructOnce.Do(func() {
		gLAreaClassStruct = gi.StructNew("Gtk", "GLAreaClass")
	})
}

type GLAreaClass struct {
	native uintptr
	// UNSUPPORTED : C value 'render' : missing Type
	// UNSUPPORTED : C value 'resize' : missing Type
	// UNSUPPORTED : C value 'create_context' : missing Type
}

var gestureClassStruct *gi.Struct
var gestureClassStructOnce sync.Once

func gestureClassStructSet() {
	gestureClassStructOnce.Do(func() {
		gestureClassStruct = gi.StructNew("Gtk", "GestureClass")
	})
}

type GestureClass struct {
	native uintptr
}

var gestureDragClassStruct *gi.Struct
var gestureDragClassStructOnce sync.Once

func gestureDragClassStructSet() {
	gestureDragClassStructOnce.Do(func() {
		gestureDragClassStruct = gi.StructNew("Gtk", "GestureDragClass")
	})
}

type GestureDragClass struct {
	native uintptr
}

var gestureLongPressClassStruct *gi.Struct
var gestureLongPressClassStructOnce sync.Once

func gestureLongPressClassStructSet() {
	gestureLongPressClassStructOnce.Do(func() {
		gestureLongPressClassStruct = gi.StructNew("Gtk", "GestureLongPressClass")
	})
}

type GestureLongPressClass struct {
	native uintptr
}

var gestureMultiPressClassStruct *gi.Struct
var gestureMultiPressClassStructOnce sync.Once

func gestureMultiPressClassStructSet() {
	gestureMultiPressClassStructOnce.Do(func() {
		gestureMultiPressClassStruct = gi.StructNew("Gtk", "GestureMultiPressClass")
	})
}

type GestureMultiPressClass struct {
	native uintptr
}

var gesturePanClassStruct *gi.Struct
var gesturePanClassStructOnce sync.Once

func gesturePanClassStructSet() {
	gesturePanClassStructOnce.Do(func() {
		gesturePanClassStruct = gi.StructNew("Gtk", "GesturePanClass")
	})
}

type GesturePanClass struct {
	native uintptr
}

var gestureRotateClassStruct *gi.Struct
var gestureRotateClassStructOnce sync.Once

func gestureRotateClassStructSet() {
	gestureRotateClassStructOnce.Do(func() {
		gestureRotateClassStruct = gi.StructNew("Gtk", "GestureRotateClass")
	})
}

type GestureRotateClass struct {
	native uintptr
}

var gestureSingleClassStruct *gi.Struct
var gestureSingleClassStructOnce sync.Once

func gestureSingleClassStructSet() {
	gestureSingleClassStructOnce.Do(func() {
		gestureSingleClassStruct = gi.StructNew("Gtk", "GestureSingleClass")
	})
}

type GestureSingleClass struct {
	native uintptr
}

var gestureStylusClassStruct *gi.Struct
var gestureStylusClassStructOnce sync.Once

func gestureStylusClassStructSet() {
	gestureStylusClassStructOnce.Do(func() {
		gestureStylusClassStruct = gi.StructNew("Gtk", "GestureStylusClass")
	})
}

type GestureStylusClass struct {
	native uintptr
}

var gestureSwipeClassStruct *gi.Struct
var gestureSwipeClassStructOnce sync.Once

func gestureSwipeClassStructSet() {
	gestureSwipeClassStructOnce.Do(func() {
		gestureSwipeClassStruct = gi.StructNew("Gtk", "GestureSwipeClass")
	})
}

type GestureSwipeClass struct {
	native uintptr
}

var gestureZoomClassStruct *gi.Struct
var gestureZoomClassStructOnce sync.Once

func gestureZoomClassStructSet() {
	gestureZoomClassStructOnce.Do(func() {
		gestureZoomClassStruct = gi.StructNew("Gtk", "GestureZoomClass")
	})
}

type GestureZoomClass struct {
	native uintptr
}

var gradientStruct *gi.Struct
var gradientStructOnce sync.Once

func gradientStructSet() {
	gradientStructOnce.Do(func() {
		gradientStruct = gi.StructNew("Gtk", "Gradient")
	})
}

type Gradient struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_gradient_new_linear' : parameter 'x0' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_gradient_new_radial' : parameter 'x0' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_gradient_add_color_stop' : parameter 'offset' of type 'gdouble' not supported

var gradientRefFunction *gi.Function
var gradientRefFunction_Once sync.Once

func gradientRefFunction_Set() {
	gradientRefFunction_Once.Do(func() {
		gradientRefFunction = gi.FunctionInvokerNew("Gtk", "ref")
	})
}

var refGradientInvoker *gi.Function

// Ref is a representation of the C type gtk_gradient_ref.
func (recv *Gradient) Ref() *Gradient {
	gradientRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := gradientRefFunction.Invoke(inArgs[:], nil)

	retGo := &Gradient{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_gradient_resolve' : parameter 'props' of type 'StyleProperties' not supported

// UNSUPPORTED : C value 'gtk_gradient_resolve_for_context' : parameter 'context' of type 'StyleContext' not supported

var gradientToStringFunction *gi.Function
var gradientToStringFunction_Once sync.Once

func gradientToStringFunction_Set() {
	gradientToStringFunction_Once.Do(func() {
		gradientToStringFunction = gi.FunctionInvokerNew("Gtk", "to_string")
	})
}

var toStringGradientInvoker *gi.Function

// ToString is a representation of the C type gtk_gradient_to_string.
func (recv *Gradient) ToString() string {
	gradientToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := gradientToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var gradientUnrefFunction *gi.Function
var gradientUnrefFunction_Once sync.Once

func gradientUnrefFunction_Set() {
	gradientUnrefFunction_Once.Do(func() {
		gradientUnrefFunction = gi.FunctionInvokerNew("Gtk", "unref")
	})
}

var unrefGradientInvoker *gi.Function

// Unref is a representation of the C type gtk_gradient_unref.
func (recv *Gradient) Unref() {
	gradientUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	gradientUnrefFunction.Invoke(inArgs[:], nil)

}

var gridClassStruct *gi.Struct
var gridClassStructOnce sync.Once

func gridClassStructSet() {
	gridClassStructOnce.Do(func() {
		gridClassStruct = gi.StructNew("Gtk", "GridClass")
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

var gridPrivateStruct *gi.Struct
var gridPrivateStructOnce sync.Once

func gridPrivateStructSet() {
	gridPrivateStructOnce.Do(func() {
		gridPrivateStruct = gi.StructNew("Gtk", "GridPrivate")
	})
}

type GridPrivate struct {
	native uintptr
}

var hBoxClassStruct *gi.Struct
var hBoxClassStructOnce sync.Once

func hBoxClassStructSet() {
	hBoxClassStructOnce.Do(func() {
		hBoxClassStruct = gi.StructNew("Gtk", "HBoxClass")
	})
}

type HBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
}

var hButtonBoxClassStruct *gi.Struct
var hButtonBoxClassStructOnce sync.Once

func hButtonBoxClassStructSet() {
	hButtonBoxClassStructOnce.Do(func() {
		hButtonBoxClassStruct = gi.StructNew("Gtk", "HButtonBoxClass")
	})
}

type HButtonBoxClass struct {
	native      uintptr
	ParentClass *ButtonBoxClass
}

var hPanedClassStruct *gi.Struct
var hPanedClassStructOnce sync.Once

func hPanedClassStructSet() {
	hPanedClassStructOnce.Do(func() {
		hPanedClassStruct = gi.StructNew("Gtk", "HPanedClass")
	})
}

type HPanedClass struct {
	native      uintptr
	ParentClass *PanedClass
}

var hSVClassStruct *gi.Struct
var hSVClassStructOnce sync.Once

func hSVClassStructSet() {
	hSVClassStructOnce.Do(func() {
		hSVClassStruct = gi.StructNew("Gtk", "HSVClass")
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

var hSVPrivateStruct *gi.Struct
var hSVPrivateStructOnce sync.Once

func hSVPrivateStructSet() {
	hSVPrivateStructOnce.Do(func() {
		hSVPrivateStruct = gi.StructNew("Gtk", "HSVPrivate")
	})
}

type HSVPrivate struct {
	native uintptr
}

var hScaleClassStruct *gi.Struct
var hScaleClassStructOnce sync.Once

func hScaleClassStructSet() {
	hScaleClassStructOnce.Do(func() {
		hScaleClassStruct = gi.StructNew("Gtk", "HScaleClass")
	})
}

type HScaleClass struct {
	native      uintptr
	ParentClass *ScaleClass
}

var hScrollbarClassStruct *gi.Struct
var hScrollbarClassStructOnce sync.Once

func hScrollbarClassStructSet() {
	hScrollbarClassStructOnce.Do(func() {
		hScrollbarClassStruct = gi.StructNew("Gtk", "HScrollbarClass")
	})
}

type HScrollbarClass struct {
	native      uintptr
	ParentClass *ScrollbarClass
}

var hSeparatorClassStruct *gi.Struct
var hSeparatorClassStructOnce sync.Once

func hSeparatorClassStructSet() {
	hSeparatorClassStructOnce.Do(func() {
		hSeparatorClassStruct = gi.StructNew("Gtk", "HSeparatorClass")
	})
}

type HSeparatorClass struct {
	native      uintptr
	ParentClass *SeparatorClass
}

var handleBoxClassStruct *gi.Struct
var handleBoxClassStructOnce sync.Once

func handleBoxClassStructSet() {
	handleBoxClassStructOnce.Do(func() {
		handleBoxClassStruct = gi.StructNew("Gtk", "HandleBoxClass")
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

var handleBoxPrivateStruct *gi.Struct
var handleBoxPrivateStructOnce sync.Once

func handleBoxPrivateStructSet() {
	handleBoxPrivateStructOnce.Do(func() {
		handleBoxPrivateStruct = gi.StructNew("Gtk", "HandleBoxPrivate")
	})
}

type HandleBoxPrivate struct {
	native uintptr
}

var headerBarAccessibleClassStruct *gi.Struct
var headerBarAccessibleClassStructOnce sync.Once

func headerBarAccessibleClassStructSet() {
	headerBarAccessibleClassStructOnce.Do(func() {
		headerBarAccessibleClassStruct = gi.StructNew("Gtk", "HeaderBarAccessibleClass")
	})
}

type HeaderBarAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var headerBarAccessiblePrivateStruct *gi.Struct
var headerBarAccessiblePrivateStructOnce sync.Once

func headerBarAccessiblePrivateStructSet() {
	headerBarAccessiblePrivateStructOnce.Do(func() {
		headerBarAccessiblePrivateStruct = gi.StructNew("Gtk", "HeaderBarAccessiblePrivate")
	})
}

type HeaderBarAccessiblePrivate struct {
	native uintptr
}

var headerBarClassStruct *gi.Struct
var headerBarClassStructOnce sync.Once

func headerBarClassStructSet() {
	headerBarClassStructOnce.Do(func() {
		headerBarClassStruct = gi.StructNew("Gtk", "HeaderBarClass")
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

var headerBarPrivateStruct *gi.Struct
var headerBarPrivateStructOnce sync.Once

func headerBarPrivateStructSet() {
	headerBarPrivateStructOnce.Do(func() {
		headerBarPrivateStruct = gi.StructNew("Gtk", "HeaderBarPrivate")
	})
}

type HeaderBarPrivate struct {
	native uintptr
}

var iMContextClassStruct *gi.Struct
var iMContextClassStructOnce sync.Once

func iMContextClassStructSet() {
	iMContextClassStructOnce.Do(func() {
		iMContextClassStruct = gi.StructNew("Gtk", "IMContextClass")
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

var iMContextInfoStruct *gi.Struct
var iMContextInfoStructOnce sync.Once

func iMContextInfoStructSet() {
	iMContextInfoStructOnce.Do(func() {
		iMContextInfoStruct = gi.StructNew("Gtk", "IMContextInfo")
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

var iMContextSimpleClassStruct *gi.Struct
var iMContextSimpleClassStructOnce sync.Once

func iMContextSimpleClassStructSet() {
	iMContextSimpleClassStructOnce.Do(func() {
		iMContextSimpleClassStruct = gi.StructNew("Gtk", "IMContextSimpleClass")
	})
}

type IMContextSimpleClass struct {
	native      uintptr
	ParentClass *IMContextClass
}

var iMContextSimplePrivateStruct *gi.Struct
var iMContextSimplePrivateStructOnce sync.Once

func iMContextSimplePrivateStructSet() {
	iMContextSimplePrivateStructOnce.Do(func() {
		iMContextSimplePrivateStruct = gi.StructNew("Gtk", "IMContextSimplePrivate")
	})
}

type IMContextSimplePrivate struct {
	native uintptr
}

var iMMulticontextClassStruct *gi.Struct
var iMMulticontextClassStructOnce sync.Once

func iMMulticontextClassStructSet() {
	iMMulticontextClassStructOnce.Do(func() {
		iMMulticontextClassStruct = gi.StructNew("Gtk", "IMMulticontextClass")
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

var iMMulticontextPrivateStruct *gi.Struct
var iMMulticontextPrivateStructOnce sync.Once

func iMMulticontextPrivateStructSet() {
	iMMulticontextPrivateStructOnce.Do(func() {
		iMMulticontextPrivateStruct = gi.StructNew("Gtk", "IMMulticontextPrivate")
	})
}

type IMMulticontextPrivate struct {
	native uintptr
}

var iconFactoryClassStruct *gi.Struct
var iconFactoryClassStructOnce sync.Once

func iconFactoryClassStructSet() {
	iconFactoryClassStructOnce.Do(func() {
		iconFactoryClassStruct = gi.StructNew("Gtk", "IconFactoryClass")
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

var iconFactoryPrivateStruct *gi.Struct
var iconFactoryPrivateStructOnce sync.Once

func iconFactoryPrivateStructSet() {
	iconFactoryPrivateStructOnce.Do(func() {
		iconFactoryPrivateStruct = gi.StructNew("Gtk", "IconFactoryPrivate")
	})
}

type IconFactoryPrivate struct {
	native uintptr
}

var iconInfoClassStruct *gi.Struct
var iconInfoClassStructOnce sync.Once

func iconInfoClassStructSet() {
	iconInfoClassStructOnce.Do(func() {
		iconInfoClassStruct = gi.StructNew("Gtk", "IconInfoClass")
	})
}

type IconInfoClass struct {
	native uintptr
}

var iconSetStruct *gi.Struct
var iconSetStructOnce sync.Once

func iconSetStructSet() {
	iconSetStructOnce.Do(func() {
		iconSetStruct = gi.StructNew("Gtk", "IconSet")
	})
}

type IconSet struct {
	native uintptr
}

var iconSetNewFunction *gi.Function
var iconSetNewFunction_Once sync.Once

func iconSetNewFunction_Set() {
	iconSetNewFunction_Once.Do(func() {
		iconSetNewFunction = gi.FunctionInvokerNew("Gtk", "new")
	})
}

var newIconSetInvoker *gi.Function

// IconSetNew is a representation of the C type gtk_icon_set_new.
func IconSetNew() *IconSet {
	iconSetNewFunction_Set()

	ret := iconSetNewFunction.Invoke(nil, nil)

	retGo := &IconSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_set_new_from_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_icon_set_add_source' : parameter 'source' of type 'IconSource' not supported

var iconSetCopyFunction *gi.Function
var iconSetCopyFunction_Once sync.Once

func iconSetCopyFunction_Set() {
	iconSetCopyFunction_Once.Do(func() {
		iconSetCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyIconSetInvoker *gi.Function

// Copy is a representation of the C type gtk_icon_set_copy.
func (recv *IconSet) Copy() *IconSet {
	iconSetCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iconSetCopyFunction.Invoke(inArgs[:], nil)

	retGo := &IconSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_set_get_sizes' : parameter 'sizes' has no type

var iconSetRefFunction *gi.Function
var iconSetRefFunction_Once sync.Once

func iconSetRefFunction_Set() {
	iconSetRefFunction_Once.Do(func() {
		iconSetRefFunction = gi.FunctionInvokerNew("Gtk", "ref")
	})
}

var refIconSetInvoker *gi.Function

// Ref is a representation of the C type gtk_icon_set_ref.
func (recv *IconSet) Ref() *IconSet {
	iconSetRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iconSetRefFunction.Invoke(inArgs[:], nil)

	retGo := &IconSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_set_render_icon' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_icon_set_render_icon_pixbuf' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_icon_set_render_icon_surface' : parameter 'context' of type 'StyleContext' not supported

var iconSetUnrefFunction *gi.Function
var iconSetUnrefFunction_Once sync.Once

func iconSetUnrefFunction_Set() {
	iconSetUnrefFunction_Once.Do(func() {
		iconSetUnrefFunction = gi.FunctionInvokerNew("Gtk", "unref")
	})
}

var unrefIconSetInvoker *gi.Function

// Unref is a representation of the C type gtk_icon_set_unref.
func (recv *IconSet) Unref() {
	iconSetUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	iconSetUnrefFunction.Invoke(inArgs[:], nil)

}

var iconSourceStruct *gi.Struct
var iconSourceStructOnce sync.Once

func iconSourceStructSet() {
	iconSourceStructOnce.Do(func() {
		iconSourceStruct = gi.StructNew("Gtk", "IconSource")
	})
}

type IconSource struct {
	native uintptr
}

var iconSourceNewFunction *gi.Function
var iconSourceNewFunction_Once sync.Once

func iconSourceNewFunction_Set() {
	iconSourceNewFunction_Once.Do(func() {
		iconSourceNewFunction = gi.FunctionInvokerNew("Gtk", "new")
	})
}

var newIconSourceInvoker *gi.Function

// IconSourceNew is a representation of the C type gtk_icon_source_new.
func IconSourceNew() *IconSource {
	iconSourceNewFunction_Set()

	ret := iconSourceNewFunction.Invoke(nil, nil)

	retGo := &IconSource{native: ret.Pointer()}

	return retGo
}

var iconSourceCopyFunction *gi.Function
var iconSourceCopyFunction_Once sync.Once

func iconSourceCopyFunction_Set() {
	iconSourceCopyFunction_Once.Do(func() {
		iconSourceCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyIconSourceInvoker *gi.Function

// Copy is a representation of the C type gtk_icon_source_copy.
func (recv *IconSource) Copy() *IconSource {
	iconSourceCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iconSourceCopyFunction.Invoke(inArgs[:], nil)

	retGo := &IconSource{native: ret.Pointer()}

	return retGo
}

var iconSourceFreeFunction *gi.Function
var iconSourceFreeFunction_Once sync.Once

func iconSourceFreeFunction_Set() {
	iconSourceFreeFunction_Once.Do(func() {
		iconSourceFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeIconSourceInvoker *gi.Function

// Free is a representation of the C type gtk_icon_source_free.
func (recv *IconSource) Free() {
	iconSourceFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	iconSourceFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_icon_source_get_direction' : return type 'TextDirection' not supported

// UNSUPPORTED : C value 'gtk_icon_source_get_direction_wildcarded' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_icon_source_get_filename' : return type 'filename' not supported

var iconSourceGetIconNameFunction *gi.Function
var iconSourceGetIconNameFunction_Once sync.Once

func iconSourceGetIconNameFunction_Set() {
	iconSourceGetIconNameFunction_Once.Do(func() {
		iconSourceGetIconNameFunction = gi.FunctionInvokerNew("Gtk", "get_icon_name")
	})
}

var getIconNameIconSourceInvoker *gi.Function

// GetIconName is a representation of the C type gtk_icon_source_get_icon_name.
func (recv *IconSource) GetIconName() string {
	iconSourceGetIconNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := iconSourceGetIconNameFunction.Invoke(inArgs[:], nil)

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

var iconSourceSetIconNameFunction *gi.Function
var iconSourceSetIconNameFunction_Once sync.Once

func iconSourceSetIconNameFunction_Set() {
	iconSourceSetIconNameFunction_Once.Do(func() {
		iconSourceSetIconNameFunction = gi.FunctionInvokerNew("Gtk", "set_icon_name")
	})
}

var setIconNameIconSourceInvoker *gi.Function

// SetIconName is a representation of the C type gtk_icon_source_set_icon_name.
func (recv *IconSource) SetIconName(iconName string) {
	iconSourceSetIconNameFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(iconName)

	iconSourceSetIconNameFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_icon_source_set_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_size' : parameter 'size' of type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_size_wildcarded' : parameter 'setting' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_state' : parameter 'state' of type 'StateType' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_state_wildcarded' : parameter 'setting' of type 'gboolean' not supported

var iconThemeClassStruct *gi.Struct
var iconThemeClassStructOnce sync.Once

func iconThemeClassStructSet() {
	iconThemeClassStructOnce.Do(func() {
		iconThemeClassStruct = gi.StructNew("Gtk", "IconThemeClass")
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

var iconThemePrivateStruct *gi.Struct
var iconThemePrivateStructOnce sync.Once

func iconThemePrivateStructSet() {
	iconThemePrivateStructOnce.Do(func() {
		iconThemePrivateStruct = gi.StructNew("Gtk", "IconThemePrivate")
	})
}

type IconThemePrivate struct {
	native uintptr
}

var iconViewAccessibleClassStruct *gi.Struct
var iconViewAccessibleClassStructOnce sync.Once

func iconViewAccessibleClassStructSet() {
	iconViewAccessibleClassStructOnce.Do(func() {
		iconViewAccessibleClassStruct = gi.StructNew("Gtk", "IconViewAccessibleClass")
	})
}

type IconViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var iconViewAccessiblePrivateStruct *gi.Struct
var iconViewAccessiblePrivateStructOnce sync.Once

func iconViewAccessiblePrivateStructSet() {
	iconViewAccessiblePrivateStructOnce.Do(func() {
		iconViewAccessiblePrivateStruct = gi.StructNew("Gtk", "IconViewAccessiblePrivate")
	})
}

type IconViewAccessiblePrivate struct {
	native uintptr
}

var iconViewClassStruct *gi.Struct
var iconViewClassStructOnce sync.Once

func iconViewClassStructSet() {
	iconViewClassStructOnce.Do(func() {
		iconViewClassStruct = gi.StructNew("Gtk", "IconViewClass")
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

var iconViewPrivateStruct *gi.Struct
var iconViewPrivateStructOnce sync.Once

func iconViewPrivateStructSet() {
	iconViewPrivateStructOnce.Do(func() {
		iconViewPrivateStruct = gi.StructNew("Gtk", "IconViewPrivate")
	})
}

type IconViewPrivate struct {
	native uintptr
}

var imageAccessibleClassStruct *gi.Struct
var imageAccessibleClassStructOnce sync.Once

func imageAccessibleClassStructSet() {
	imageAccessibleClassStructOnce.Do(func() {
		imageAccessibleClassStruct = gi.StructNew("Gtk", "ImageAccessibleClass")
	})
}

type ImageAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var imageAccessiblePrivateStruct *gi.Struct
var imageAccessiblePrivateStructOnce sync.Once

func imageAccessiblePrivateStructSet() {
	imageAccessiblePrivateStructOnce.Do(func() {
		imageAccessiblePrivateStruct = gi.StructNew("Gtk", "ImageAccessiblePrivate")
	})
}

type ImageAccessiblePrivate struct {
	native uintptr
}

var imageCellAccessibleClassStruct *gi.Struct
var imageCellAccessibleClassStructOnce sync.Once

func imageCellAccessibleClassStructSet() {
	imageCellAccessibleClassStructOnce.Do(func() {
		imageCellAccessibleClassStruct = gi.StructNew("Gtk", "ImageCellAccessibleClass")
	})
}

type ImageCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

var imageCellAccessiblePrivateStruct *gi.Struct
var imageCellAccessiblePrivateStructOnce sync.Once

func imageCellAccessiblePrivateStructSet() {
	imageCellAccessiblePrivateStructOnce.Do(func() {
		imageCellAccessiblePrivateStruct = gi.StructNew("Gtk", "ImageCellAccessiblePrivate")
	})
}

type ImageCellAccessiblePrivate struct {
	native uintptr
}

var imageClassStruct *gi.Struct
var imageClassStructOnce sync.Once

func imageClassStructSet() {
	imageClassStructOnce.Do(func() {
		imageClassStruct = gi.StructNew("Gtk", "ImageClass")
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

var imageMenuItemClassStruct *gi.Struct
var imageMenuItemClassStructOnce sync.Once

func imageMenuItemClassStructSet() {
	imageMenuItemClassStructOnce.Do(func() {
		imageMenuItemClassStruct = gi.StructNew("Gtk", "ImageMenuItemClass")
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

var imageMenuItemPrivateStruct *gi.Struct
var imageMenuItemPrivateStructOnce sync.Once

func imageMenuItemPrivateStructSet() {
	imageMenuItemPrivateStructOnce.Do(func() {
		imageMenuItemPrivateStruct = gi.StructNew("Gtk", "ImageMenuItemPrivate")
	})
}

type ImageMenuItemPrivate struct {
	native uintptr
}

var imagePrivateStruct *gi.Struct
var imagePrivateStructOnce sync.Once

func imagePrivateStructSet() {
	imagePrivateStructOnce.Do(func() {
		imagePrivateStruct = gi.StructNew("Gtk", "ImagePrivate")
	})
}

type ImagePrivate struct {
	native uintptr
}

var infoBarClassStruct *gi.Struct
var infoBarClassStructOnce sync.Once

func infoBarClassStructSet() {
	infoBarClassStructOnce.Do(func() {
		infoBarClassStruct = gi.StructNew("Gtk", "InfoBarClass")
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

var infoBarPrivateStruct *gi.Struct
var infoBarPrivateStructOnce sync.Once

func infoBarPrivateStructSet() {
	infoBarPrivateStructOnce.Do(func() {
		infoBarPrivateStruct = gi.StructNew("Gtk", "InfoBarPrivate")
	})
}

type InfoBarPrivate struct {
	native uintptr
}

var invisibleClassStruct *gi.Struct
var invisibleClassStructOnce sync.Once

func invisibleClassStructSet() {
	invisibleClassStructOnce.Do(func() {
		invisibleClassStruct = gi.StructNew("Gtk", "InvisibleClass")
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

var invisiblePrivateStruct *gi.Struct
var invisiblePrivateStructOnce sync.Once

func invisiblePrivateStructSet() {
	invisiblePrivateStructOnce.Do(func() {
		invisiblePrivateStruct = gi.StructNew("Gtk", "InvisiblePrivate")
	})
}

type InvisiblePrivate struct {
	native uintptr
}

var labelAccessibleClassStruct *gi.Struct
var labelAccessibleClassStructOnce sync.Once

func labelAccessibleClassStructSet() {
	labelAccessibleClassStructOnce.Do(func() {
		labelAccessibleClassStruct = gi.StructNew("Gtk", "LabelAccessibleClass")
	})
}

type LabelAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var labelAccessiblePrivateStruct *gi.Struct
var labelAccessiblePrivateStructOnce sync.Once

func labelAccessiblePrivateStructSet() {
	labelAccessiblePrivateStructOnce.Do(func() {
		labelAccessiblePrivateStruct = gi.StructNew("Gtk", "LabelAccessiblePrivate")
	})
}

type LabelAccessiblePrivate struct {
	native uintptr
}

var labelClassStruct *gi.Struct
var labelClassStructOnce sync.Once

func labelClassStructSet() {
	labelClassStructOnce.Do(func() {
		labelClassStruct = gi.StructNew("Gtk", "LabelClass")
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

var labelPrivateStruct *gi.Struct
var labelPrivateStructOnce sync.Once

func labelPrivateStructSet() {
	labelPrivateStructOnce.Do(func() {
		labelPrivateStruct = gi.StructNew("Gtk", "LabelPrivate")
	})
}

type LabelPrivate struct {
	native uintptr
}

var labelSelectionInfoStruct *gi.Struct
var labelSelectionInfoStructOnce sync.Once

func labelSelectionInfoStructSet() {
	labelSelectionInfoStructOnce.Do(func() {
		labelSelectionInfoStruct = gi.StructNew("Gtk", "LabelSelectionInfo")
	})
}

type LabelSelectionInfo struct {
	native uintptr
}

var layoutClassStruct *gi.Struct
var layoutClassStructOnce sync.Once

func layoutClassStructSet() {
	layoutClassStructOnce.Do(func() {
		layoutClassStruct = gi.StructNew("Gtk", "LayoutClass")
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

var layoutPrivateStruct *gi.Struct
var layoutPrivateStructOnce sync.Once

func layoutPrivateStructSet() {
	layoutPrivateStructOnce.Do(func() {
		layoutPrivateStruct = gi.StructNew("Gtk", "LayoutPrivate")
	})
}

type LayoutPrivate struct {
	native uintptr
}

var levelBarAccessibleClassStruct *gi.Struct
var levelBarAccessibleClassStructOnce sync.Once

func levelBarAccessibleClassStructSet() {
	levelBarAccessibleClassStructOnce.Do(func() {
		levelBarAccessibleClassStruct = gi.StructNew("Gtk", "LevelBarAccessibleClass")
	})
}

type LevelBarAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var levelBarAccessiblePrivateStruct *gi.Struct
var levelBarAccessiblePrivateStructOnce sync.Once

func levelBarAccessiblePrivateStructSet() {
	levelBarAccessiblePrivateStructOnce.Do(func() {
		levelBarAccessiblePrivateStruct = gi.StructNew("Gtk", "LevelBarAccessiblePrivate")
	})
}

type LevelBarAccessiblePrivate struct {
	native uintptr
}

var levelBarClassStruct *gi.Struct
var levelBarClassStructOnce sync.Once

func levelBarClassStructSet() {
	levelBarClassStructOnce.Do(func() {
		levelBarClassStruct = gi.StructNew("Gtk", "LevelBarClass")
	})
}

type LevelBarClass struct {
	native uintptr
	// UNSUPPORTED : C value 'offset_changed' : missing Type
}

var levelBarPrivateStruct *gi.Struct
var levelBarPrivateStructOnce sync.Once

func levelBarPrivateStructSet() {
	levelBarPrivateStructOnce.Do(func() {
		levelBarPrivateStruct = gi.StructNew("Gtk", "LevelBarPrivate")
	})
}

type LevelBarPrivate struct {
	native uintptr
}

var linkButtonAccessibleClassStruct *gi.Struct
var linkButtonAccessibleClassStructOnce sync.Once

func linkButtonAccessibleClassStructSet() {
	linkButtonAccessibleClassStructOnce.Do(func() {
		linkButtonAccessibleClassStruct = gi.StructNew("Gtk", "LinkButtonAccessibleClass")
	})
}

type LinkButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var linkButtonAccessiblePrivateStruct *gi.Struct
var linkButtonAccessiblePrivateStructOnce sync.Once

func linkButtonAccessiblePrivateStructSet() {
	linkButtonAccessiblePrivateStructOnce.Do(func() {
		linkButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "LinkButtonAccessiblePrivate")
	})
}

type LinkButtonAccessiblePrivate struct {
	native uintptr
}

var linkButtonClassStruct *gi.Struct
var linkButtonClassStructOnce sync.Once

func linkButtonClassStructSet() {
	linkButtonClassStructOnce.Do(func() {
		linkButtonClassStruct = gi.StructNew("Gtk", "LinkButtonClass")
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

var linkButtonPrivateStruct *gi.Struct
var linkButtonPrivateStructOnce sync.Once

func linkButtonPrivateStructSet() {
	linkButtonPrivateStructOnce.Do(func() {
		linkButtonPrivateStruct = gi.StructNew("Gtk", "LinkButtonPrivate")
	})
}

type LinkButtonPrivate struct {
	native uintptr
}

var listBoxAccessibleClassStruct *gi.Struct
var listBoxAccessibleClassStructOnce sync.Once

func listBoxAccessibleClassStructSet() {
	listBoxAccessibleClassStructOnce.Do(func() {
		listBoxAccessibleClassStruct = gi.StructNew("Gtk", "ListBoxAccessibleClass")
	})
}

type ListBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var listBoxAccessiblePrivateStruct *gi.Struct
var listBoxAccessiblePrivateStructOnce sync.Once

func listBoxAccessiblePrivateStructSet() {
	listBoxAccessiblePrivateStructOnce.Do(func() {
		listBoxAccessiblePrivateStruct = gi.StructNew("Gtk", "ListBoxAccessiblePrivate")
	})
}

type ListBoxAccessiblePrivate struct {
	native uintptr
}

var listBoxClassStruct *gi.Struct
var listBoxClassStructOnce sync.Once

func listBoxClassStructSet() {
	listBoxClassStructOnce.Do(func() {
		listBoxClassStruct = gi.StructNew("Gtk", "ListBoxClass")
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

var listBoxRowAccessibleClassStruct *gi.Struct
var listBoxRowAccessibleClassStructOnce sync.Once

func listBoxRowAccessibleClassStructSet() {
	listBoxRowAccessibleClassStructOnce.Do(func() {
		listBoxRowAccessibleClassStruct = gi.StructNew("Gtk", "ListBoxRowAccessibleClass")
	})
}

type ListBoxRowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var listBoxRowClassStruct *gi.Struct
var listBoxRowClassStructOnce sync.Once

func listBoxRowClassStructSet() {
	listBoxRowClassStructOnce.Do(func() {
		listBoxRowClassStruct = gi.StructNew("Gtk", "ListBoxRowClass")
	})
}

type ListBoxRowClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'activate' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
}

var listStoreClassStruct *gi.Struct
var listStoreClassStructOnce sync.Once

func listStoreClassStructSet() {
	listStoreClassStructOnce.Do(func() {
		listStoreClassStruct = gi.StructNew("Gtk", "ListStoreClass")
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

var listStorePrivateStruct *gi.Struct
var listStorePrivateStructOnce sync.Once

func listStorePrivateStructSet() {
	listStorePrivateStructOnce.Do(func() {
		listStorePrivateStruct = gi.StructNew("Gtk", "ListStorePrivate")
	})
}

type ListStorePrivate struct {
	native uintptr
}

var lockButtonAccessibleClassStruct *gi.Struct
var lockButtonAccessibleClassStructOnce sync.Once

func lockButtonAccessibleClassStructSet() {
	lockButtonAccessibleClassStructOnce.Do(func() {
		lockButtonAccessibleClassStruct = gi.StructNew("Gtk", "LockButtonAccessibleClass")
	})
}

type LockButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var lockButtonAccessiblePrivateStruct *gi.Struct
var lockButtonAccessiblePrivateStructOnce sync.Once

func lockButtonAccessiblePrivateStructSet() {
	lockButtonAccessiblePrivateStructOnce.Do(func() {
		lockButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "LockButtonAccessiblePrivate")
	})
}

type LockButtonAccessiblePrivate struct {
	native uintptr
}

var lockButtonClassStruct *gi.Struct
var lockButtonClassStructOnce sync.Once

func lockButtonClassStructSet() {
	lockButtonClassStructOnce.Do(func() {
		lockButtonClassStruct = gi.StructNew("Gtk", "LockButtonClass")
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

var lockButtonPrivateStruct *gi.Struct
var lockButtonPrivateStructOnce sync.Once

func lockButtonPrivateStructSet() {
	lockButtonPrivateStructOnce.Do(func() {
		lockButtonPrivateStruct = gi.StructNew("Gtk", "LockButtonPrivate")
	})
}

type LockButtonPrivate struct {
	native uintptr
}

var menuAccessibleClassStruct *gi.Struct
var menuAccessibleClassStructOnce sync.Once

func menuAccessibleClassStructSet() {
	menuAccessibleClassStructOnce.Do(func() {
		menuAccessibleClassStruct = gi.StructNew("Gtk", "MenuAccessibleClass")
	})
}

type MenuAccessibleClass struct {
	native      uintptr
	ParentClass *MenuShellAccessibleClass
}

var menuAccessiblePrivateStruct *gi.Struct
var menuAccessiblePrivateStructOnce sync.Once

func menuAccessiblePrivateStructSet() {
	menuAccessiblePrivateStructOnce.Do(func() {
		menuAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuAccessiblePrivate")
	})
}

type MenuAccessiblePrivate struct {
	native uintptr
}

var menuBarClassStruct *gi.Struct
var menuBarClassStructOnce sync.Once

func menuBarClassStructSet() {
	menuBarClassStructOnce.Do(func() {
		menuBarClassStruct = gi.StructNew("Gtk", "MenuBarClass")
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

var menuBarPrivateStruct *gi.Struct
var menuBarPrivateStructOnce sync.Once

func menuBarPrivateStructSet() {
	menuBarPrivateStructOnce.Do(func() {
		menuBarPrivateStruct = gi.StructNew("Gtk", "MenuBarPrivate")
	})
}

type MenuBarPrivate struct {
	native uintptr
}

var menuButtonAccessibleClassStruct *gi.Struct
var menuButtonAccessibleClassStructOnce sync.Once

func menuButtonAccessibleClassStructSet() {
	menuButtonAccessibleClassStructOnce.Do(func() {
		menuButtonAccessibleClassStruct = gi.StructNew("Gtk", "MenuButtonAccessibleClass")
	})
}

type MenuButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ToggleButtonAccessibleClass
}

var menuButtonAccessiblePrivateStruct *gi.Struct
var menuButtonAccessiblePrivateStructOnce sync.Once

func menuButtonAccessiblePrivateStructSet() {
	menuButtonAccessiblePrivateStructOnce.Do(func() {
		menuButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuButtonAccessiblePrivate")
	})
}

type MenuButtonAccessiblePrivate struct {
	native uintptr
}

var menuButtonClassStruct *gi.Struct
var menuButtonClassStructOnce sync.Once

func menuButtonClassStructSet() {
	menuButtonClassStructOnce.Do(func() {
		menuButtonClassStruct = gi.StructNew("Gtk", "MenuButtonClass")
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

var menuButtonPrivateStruct *gi.Struct
var menuButtonPrivateStructOnce sync.Once

func menuButtonPrivateStructSet() {
	menuButtonPrivateStructOnce.Do(func() {
		menuButtonPrivateStruct = gi.StructNew("Gtk", "MenuButtonPrivate")
	})
}

type MenuButtonPrivate struct {
	native uintptr
}

var menuClassStruct *gi.Struct
var menuClassStructOnce sync.Once

func menuClassStructSet() {
	menuClassStructOnce.Do(func() {
		menuClassStruct = gi.StructNew("Gtk", "MenuClass")
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

var menuItemAccessibleClassStruct *gi.Struct
var menuItemAccessibleClassStructOnce sync.Once

func menuItemAccessibleClassStructSet() {
	menuItemAccessibleClassStructOnce.Do(func() {
		menuItemAccessibleClassStruct = gi.StructNew("Gtk", "MenuItemAccessibleClass")
	})
}

type MenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var menuItemAccessiblePrivateStruct *gi.Struct
var menuItemAccessiblePrivateStructOnce sync.Once

func menuItemAccessiblePrivateStructSet() {
	menuItemAccessiblePrivateStructOnce.Do(func() {
		menuItemAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuItemAccessiblePrivate")
	})
}

type MenuItemAccessiblePrivate struct {
	native uintptr
}

var menuItemClassStruct *gi.Struct
var menuItemClassStructOnce sync.Once

func menuItemClassStructSet() {
	menuItemClassStructOnce.Do(func() {
		menuItemClassStruct = gi.StructNew("Gtk", "MenuItemClass")
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

var menuItemPrivateStruct *gi.Struct
var menuItemPrivateStructOnce sync.Once

func menuItemPrivateStructSet() {
	menuItemPrivateStructOnce.Do(func() {
		menuItemPrivateStruct = gi.StructNew("Gtk", "MenuItemPrivate")
	})
}

type MenuItemPrivate struct {
	native uintptr
}

var menuPrivateStruct *gi.Struct
var menuPrivateStructOnce sync.Once

func menuPrivateStructSet() {
	menuPrivateStructOnce.Do(func() {
		menuPrivateStruct = gi.StructNew("Gtk", "MenuPrivate")
	})
}

type MenuPrivate struct {
	native uintptr
}

var menuShellAccessibleClassStruct *gi.Struct
var menuShellAccessibleClassStructOnce sync.Once

func menuShellAccessibleClassStructSet() {
	menuShellAccessibleClassStructOnce.Do(func() {
		menuShellAccessibleClassStruct = gi.StructNew("Gtk", "MenuShellAccessibleClass")
	})
}

type MenuShellAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var menuShellAccessiblePrivateStruct *gi.Struct
var menuShellAccessiblePrivateStructOnce sync.Once

func menuShellAccessiblePrivateStructSet() {
	menuShellAccessiblePrivateStructOnce.Do(func() {
		menuShellAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuShellAccessiblePrivate")
	})
}

type MenuShellAccessiblePrivate struct {
	native uintptr
}

var menuShellClassStruct *gi.Struct
var menuShellClassStructOnce sync.Once

func menuShellClassStructSet() {
	menuShellClassStructOnce.Do(func() {
		menuShellClassStruct = gi.StructNew("Gtk", "MenuShellClass")
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

var menuShellPrivateStruct *gi.Struct
var menuShellPrivateStructOnce sync.Once

func menuShellPrivateStructSet() {
	menuShellPrivateStructOnce.Do(func() {
		menuShellPrivateStruct = gi.StructNew("Gtk", "MenuShellPrivate")
	})
}

type MenuShellPrivate struct {
	native uintptr
}

var menuToolButtonClassStruct *gi.Struct
var menuToolButtonClassStructOnce sync.Once

func menuToolButtonClassStructSet() {
	menuToolButtonClassStructOnce.Do(func() {
		menuToolButtonClassStruct = gi.StructNew("Gtk", "MenuToolButtonClass")
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

var menuToolButtonPrivateStruct *gi.Struct
var menuToolButtonPrivateStructOnce sync.Once

func menuToolButtonPrivateStructSet() {
	menuToolButtonPrivateStructOnce.Do(func() {
		menuToolButtonPrivateStruct = gi.StructNew("Gtk", "MenuToolButtonPrivate")
	})
}

type MenuToolButtonPrivate struct {
	native uintptr
}

var messageDialogClassStruct *gi.Struct
var messageDialogClassStructOnce sync.Once

func messageDialogClassStructSet() {
	messageDialogClassStructOnce.Do(func() {
		messageDialogClassStruct = gi.StructNew("Gtk", "MessageDialogClass")
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

var messageDialogPrivateStruct *gi.Struct
var messageDialogPrivateStructOnce sync.Once

func messageDialogPrivateStructSet() {
	messageDialogPrivateStructOnce.Do(func() {
		messageDialogPrivateStruct = gi.StructNew("Gtk", "MessageDialogPrivate")
	})
}

type MessageDialogPrivate struct {
	native uintptr
}

var miscClassStruct *gi.Struct
var miscClassStructOnce sync.Once

func miscClassStructSet() {
	miscClassStructOnce.Do(func() {
		miscClassStruct = gi.StructNew("Gtk", "MiscClass")
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

var miscPrivateStruct *gi.Struct
var miscPrivateStructOnce sync.Once

func miscPrivateStructSet() {
	miscPrivateStructOnce.Do(func() {
		miscPrivateStruct = gi.StructNew("Gtk", "MiscPrivate")
	})
}

type MiscPrivate struct {
	native uintptr
}

var mountOperationClassStruct *gi.Struct
var mountOperationClassStructOnce sync.Once

func mountOperationClassStructSet() {
	mountOperationClassStructOnce.Do(func() {
		mountOperationClassStruct = gi.StructNew("Gtk", "MountOperationClass")
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

var mountOperationPrivateStruct *gi.Struct
var mountOperationPrivateStructOnce sync.Once

func mountOperationPrivateStructSet() {
	mountOperationPrivateStructOnce.Do(func() {
		mountOperationPrivateStruct = gi.StructNew("Gtk", "MountOperationPrivate")
	})
}

type MountOperationPrivate struct {
	native uintptr
}

var nativeDialogClassStruct *gi.Struct
var nativeDialogClassStructOnce sync.Once

func nativeDialogClassStructSet() {
	nativeDialogClassStructOnce.Do(func() {
		nativeDialogClassStruct = gi.StructNew("Gtk", "NativeDialogClass")
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

var notebookAccessibleClassStruct *gi.Struct
var notebookAccessibleClassStructOnce sync.Once

func notebookAccessibleClassStructSet() {
	notebookAccessibleClassStructOnce.Do(func() {
		notebookAccessibleClassStruct = gi.StructNew("Gtk", "NotebookAccessibleClass")
	})
}

type NotebookAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var notebookAccessiblePrivateStruct *gi.Struct
var notebookAccessiblePrivateStructOnce sync.Once

func notebookAccessiblePrivateStructSet() {
	notebookAccessiblePrivateStructOnce.Do(func() {
		notebookAccessiblePrivateStruct = gi.StructNew("Gtk", "NotebookAccessiblePrivate")
	})
}

type NotebookAccessiblePrivate struct {
	native uintptr
}

var notebookClassStruct *gi.Struct
var notebookClassStructOnce sync.Once

func notebookClassStructSet() {
	notebookClassStructOnce.Do(func() {
		notebookClassStruct = gi.StructNew("Gtk", "NotebookClass")
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

var notebookPageAccessibleClassStruct *gi.Struct
var notebookPageAccessibleClassStructOnce sync.Once

func notebookPageAccessibleClassStructSet() {
	notebookPageAccessibleClassStructOnce.Do(func() {
		notebookPageAccessibleClassStruct = gi.StructNew("Gtk", "NotebookPageAccessibleClass")
	})
}

type NotebookPageAccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Atk.ObjectClass'
}

var notebookPageAccessiblePrivateStruct *gi.Struct
var notebookPageAccessiblePrivateStructOnce sync.Once

func notebookPageAccessiblePrivateStructSet() {
	notebookPageAccessiblePrivateStructOnce.Do(func() {
		notebookPageAccessiblePrivateStruct = gi.StructNew("Gtk", "NotebookPageAccessiblePrivate")
	})
}

type NotebookPageAccessiblePrivate struct {
	native uintptr
}

var notebookPrivateStruct *gi.Struct
var notebookPrivateStructOnce sync.Once

func notebookPrivateStructSet() {
	notebookPrivateStructOnce.Do(func() {
		notebookPrivateStruct = gi.StructNew("Gtk", "NotebookPrivate")
	})
}

type NotebookPrivate struct {
	native uintptr
}

var numerableIconClassStruct *gi.Struct
var numerableIconClassStructOnce sync.Once

func numerableIconClassStructSet() {
	numerableIconClassStructOnce.Do(func() {
		numerableIconClassStruct = gi.StructNew("Gtk", "NumerableIconClass")
	})
}

type NumerableIconClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.EmblemedIconClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var numerableIconPrivateStruct *gi.Struct
var numerableIconPrivateStructOnce sync.Once

func numerableIconPrivateStructSet() {
	numerableIconPrivateStructOnce.Do(func() {
		numerableIconPrivateStruct = gi.StructNew("Gtk", "NumerableIconPrivate")
	})
}

type NumerableIconPrivate struct {
	native uintptr
}

var offscreenWindowClassStruct *gi.Struct
var offscreenWindowClassStructOnce sync.Once

func offscreenWindowClassStructSet() {
	offscreenWindowClassStructOnce.Do(func() {
		offscreenWindowClassStruct = gi.StructNew("Gtk", "OffscreenWindowClass")
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

var orientableIfaceStruct *gi.Struct
var orientableIfaceStructOnce sync.Once

func orientableIfaceStructSet() {
	orientableIfaceStructOnce.Do(func() {
		orientableIfaceStruct = gi.StructNew("Gtk", "OrientableIface")
	})
}

type OrientableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'
}

var overlayClassStruct *gi.Struct
var overlayClassStructOnce sync.Once

func overlayClassStructSet() {
	overlayClassStructOnce.Do(func() {
		overlayClassStruct = gi.StructNew("Gtk", "OverlayClass")
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

var overlayPrivateStruct *gi.Struct
var overlayPrivateStructOnce sync.Once

func overlayPrivateStructSet() {
	overlayPrivateStructOnce.Do(func() {
		overlayPrivateStruct = gi.StructNew("Gtk", "OverlayPrivate")
	})
}

type OverlayPrivate struct {
	native uintptr
}

var padActionEntryStruct *gi.Struct
var padActionEntryStructOnce sync.Once

func padActionEntryStructSet() {
	padActionEntryStructOnce.Do(func() {
		padActionEntryStruct = gi.StructNew("Gtk", "PadActionEntry")
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

var padControllerClassStruct *gi.Struct
var padControllerClassStructOnce sync.Once

func padControllerClassStructSet() {
	padControllerClassStructOnce.Do(func() {
		padControllerClassStruct = gi.StructNew("Gtk", "PadControllerClass")
	})
}

type PadControllerClass struct {
	native uintptr
}

var pageRangeStruct *gi.Struct
var pageRangeStructOnce sync.Once

func pageRangeStructSet() {
	pageRangeStructOnce.Do(func() {
		pageRangeStruct = gi.StructNew("Gtk", "PageRange")
	})
}

type PageRange struct {
	native uintptr
	Start  int32
	End    int32
}

var panedAccessibleClassStruct *gi.Struct
var panedAccessibleClassStructOnce sync.Once

func panedAccessibleClassStructSet() {
	panedAccessibleClassStructOnce.Do(func() {
		panedAccessibleClassStruct = gi.StructNew("Gtk", "PanedAccessibleClass")
	})
}

type PanedAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var panedAccessiblePrivateStruct *gi.Struct
var panedAccessiblePrivateStructOnce sync.Once

func panedAccessiblePrivateStructSet() {
	panedAccessiblePrivateStructOnce.Do(func() {
		panedAccessiblePrivateStruct = gi.StructNew("Gtk", "PanedAccessiblePrivate")
	})
}

type PanedAccessiblePrivate struct {
	native uintptr
}

var panedClassStruct *gi.Struct
var panedClassStructOnce sync.Once

func panedClassStructSet() {
	panedClassStructOnce.Do(func() {
		panedClassStruct = gi.StructNew("Gtk", "PanedClass")
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

var panedPrivateStruct *gi.Struct
var panedPrivateStructOnce sync.Once

func panedPrivateStructSet() {
	panedPrivateStructOnce.Do(func() {
		panedPrivateStruct = gi.StructNew("Gtk", "PanedPrivate")
	})
}

type PanedPrivate struct {
	native uintptr
}

var paperSizeStruct *gi.Struct
var paperSizeStructOnce sync.Once

func paperSizeStructSet() {
	paperSizeStructOnce.Do(func() {
		paperSizeStruct = gi.StructNew("Gtk", "PaperSize")
	})
}

type PaperSize struct {
	native uintptr
}

var paperSizeNewFunction *gi.Function
var paperSizeNewFunction_Once sync.Once

func paperSizeNewFunction_Set() {
	paperSizeNewFunction_Once.Do(func() {
		paperSizeNewFunction = gi.FunctionInvokerNew("Gtk", "new")
	})
}

var newPaperSizeInvoker *gi.Function

// PaperSizeNew is a representation of the C type gtk_paper_size_new.
func PaperSizeNew(name string) *PaperSize {
	paperSizeNewFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	ret := paperSizeNewFunction.Invoke(inArgs[:], nil)

	retGo := &PaperSize{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_new_custom' : parameter 'width' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_gvariant' : parameter 'variant' of type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_ipp' : parameter 'width' of type 'gdouble' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_key_file' : parameter 'key_file' of type 'GLib.KeyFile' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_ppd' : parameter 'width' of type 'gdouble' not supported

var paperSizeCopyFunction *gi.Function
var paperSizeCopyFunction_Once sync.Once

func paperSizeCopyFunction_Set() {
	paperSizeCopyFunction_Once.Do(func() {
		paperSizeCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyPaperSizeInvoker *gi.Function

// Copy is a representation of the C type gtk_paper_size_copy.
func (recv *PaperSize) Copy() *PaperSize {
	paperSizeCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := paperSizeCopyFunction.Invoke(inArgs[:], nil)

	retGo := &PaperSize{native: ret.Pointer()}

	return retGo
}

var paperSizeFreeFunction *gi.Function
var paperSizeFreeFunction_Once sync.Once

func paperSizeFreeFunction_Set() {
	paperSizeFreeFunction_Once.Do(func() {
		paperSizeFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freePaperSizeInvoker *gi.Function

// Free is a representation of the C type gtk_paper_size_free.
func (recv *PaperSize) Free() {
	paperSizeFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	paperSizeFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_paper_size_get_default_bottom_margin' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_default_left_margin' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_default_right_margin' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_default_top_margin' : parameter 'unit' of type 'Unit' not supported

var paperSizeGetDisplayNameFunction *gi.Function
var paperSizeGetDisplayNameFunction_Once sync.Once

func paperSizeGetDisplayNameFunction_Set() {
	paperSizeGetDisplayNameFunction_Once.Do(func() {
		paperSizeGetDisplayNameFunction = gi.FunctionInvokerNew("Gtk", "get_display_name")
	})
}

var getDisplayNamePaperSizeInvoker *gi.Function

// GetDisplayName is a representation of the C type gtk_paper_size_get_display_name.
func (recv *PaperSize) GetDisplayName() string {
	paperSizeGetDisplayNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := paperSizeGetDisplayNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_get_height' : parameter 'unit' of type 'Unit' not supported

var paperSizeGetNameFunction *gi.Function
var paperSizeGetNameFunction_Once sync.Once

func paperSizeGetNameFunction_Set() {
	paperSizeGetNameFunction_Once.Do(func() {
		paperSizeGetNameFunction = gi.FunctionInvokerNew("Gtk", "get_name")
	})
}

var getNamePaperSizeInvoker *gi.Function

// GetName is a representation of the C type gtk_paper_size_get_name.
func (recv *PaperSize) GetName() string {
	paperSizeGetNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := paperSizeGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var paperSizeGetPpdNameFunction *gi.Function
var paperSizeGetPpdNameFunction_Once sync.Once

func paperSizeGetPpdNameFunction_Set() {
	paperSizeGetPpdNameFunction_Once.Do(func() {
		paperSizeGetPpdNameFunction = gi.FunctionInvokerNew("Gtk", "get_ppd_name")
	})
}

var getPpdNamePaperSizeInvoker *gi.Function

// GetPpdName is a representation of the C type gtk_paper_size_get_ppd_name.
func (recv *PaperSize) GetPpdName() string {
	paperSizeGetPpdNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := paperSizeGetPpdNameFunction.Invoke(inArgs[:], nil)

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

var placesSidebarClassStruct *gi.Struct
var placesSidebarClassStructOnce sync.Once

func placesSidebarClassStructSet() {
	placesSidebarClassStructOnce.Do(func() {
		placesSidebarClassStruct = gi.StructNew("Gtk", "PlacesSidebarClass")
	})
}

type PlacesSidebarClass struct {
	native uintptr
}

var plugClassStruct *gi.Struct
var plugClassStructOnce sync.Once

func plugClassStructSet() {
	plugClassStructOnce.Do(func() {
		plugClassStruct = gi.StructNew("Gtk", "PlugClass")
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

var plugPrivateStruct *gi.Struct
var plugPrivateStructOnce sync.Once

func plugPrivateStructSet() {
	plugPrivateStructOnce.Do(func() {
		plugPrivateStruct = gi.StructNew("Gtk", "PlugPrivate")
	})
}

type PlugPrivate struct {
	native uintptr
}

var popoverAccessibleClassStruct *gi.Struct
var popoverAccessibleClassStructOnce sync.Once

func popoverAccessibleClassStructSet() {
	popoverAccessibleClassStructOnce.Do(func() {
		popoverAccessibleClassStruct = gi.StructNew("Gtk", "PopoverAccessibleClass")
	})
}

type PopoverAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var popoverClassStruct *gi.Struct
var popoverClassStructOnce sync.Once

func popoverClassStructSet() {
	popoverClassStructOnce.Do(func() {
		popoverClassStruct = gi.StructNew("Gtk", "PopoverClass")
	})
}

type PopoverClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'closed' : missing Type
}

var popoverMenuClassStruct *gi.Struct
var popoverMenuClassStructOnce sync.Once

func popoverMenuClassStructSet() {
	popoverMenuClassStructOnce.Do(func() {
		popoverMenuClassStruct = gi.StructNew("Gtk", "PopoverMenuClass")
	})
}

type PopoverMenuClass struct {
	native      uintptr
	ParentClass *PopoverClass
}

var popoverPrivateStruct *gi.Struct
var popoverPrivateStructOnce sync.Once

func popoverPrivateStructSet() {
	popoverPrivateStructOnce.Do(func() {
		popoverPrivateStruct = gi.StructNew("Gtk", "PopoverPrivate")
	})
}

type PopoverPrivate struct {
	native uintptr
}

var printOperationClassStruct *gi.Struct
var printOperationClassStructOnce sync.Once

func printOperationClassStructSet() {
	printOperationClassStructOnce.Do(func() {
		printOperationClassStruct = gi.StructNew("Gtk", "PrintOperationClass")
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

var printOperationPreviewIfaceStruct *gi.Struct
var printOperationPreviewIfaceStructOnce sync.Once

func printOperationPreviewIfaceStructSet() {
	printOperationPreviewIfaceStructOnce.Do(func() {
		printOperationPreviewIfaceStruct = gi.StructNew("Gtk", "PrintOperationPreviewIface")
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

var printOperationPrivateStruct *gi.Struct
var printOperationPrivateStructOnce sync.Once

func printOperationPrivateStructSet() {
	printOperationPrivateStructOnce.Do(func() {
		printOperationPrivateStruct = gi.StructNew("Gtk", "PrintOperationPrivate")
	})
}

type PrintOperationPrivate struct {
	native uintptr
}

var progressBarAccessibleClassStruct *gi.Struct
var progressBarAccessibleClassStructOnce sync.Once

func progressBarAccessibleClassStructSet() {
	progressBarAccessibleClassStructOnce.Do(func() {
		progressBarAccessibleClassStruct = gi.StructNew("Gtk", "ProgressBarAccessibleClass")
	})
}

type ProgressBarAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var progressBarAccessiblePrivateStruct *gi.Struct
var progressBarAccessiblePrivateStructOnce sync.Once

func progressBarAccessiblePrivateStructSet() {
	progressBarAccessiblePrivateStructOnce.Do(func() {
		progressBarAccessiblePrivateStruct = gi.StructNew("Gtk", "ProgressBarAccessiblePrivate")
	})
}

type ProgressBarAccessiblePrivate struct {
	native uintptr
}

var progressBarClassStruct *gi.Struct
var progressBarClassStructOnce sync.Once

func progressBarClassStructSet() {
	progressBarClassStructOnce.Do(func() {
		progressBarClassStruct = gi.StructNew("Gtk", "ProgressBarClass")
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

var progressBarPrivateStruct *gi.Struct
var progressBarPrivateStructOnce sync.Once

func progressBarPrivateStructSet() {
	progressBarPrivateStructOnce.Do(func() {
		progressBarPrivateStruct = gi.StructNew("Gtk", "ProgressBarPrivate")
	})
}

type ProgressBarPrivate struct {
	native uintptr
}

var radioActionClassStruct *gi.Struct
var radioActionClassStructOnce sync.Once

func radioActionClassStructSet() {
	radioActionClassStructOnce.Do(func() {
		radioActionClassStruct = gi.StructNew("Gtk", "RadioActionClass")
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

var radioActionEntryStruct *gi.Struct
var radioActionEntryStructOnce sync.Once

func radioActionEntryStructSet() {
	radioActionEntryStructOnce.Do(func() {
		radioActionEntryStruct = gi.StructNew("Gtk", "RadioActionEntry")
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

var radioActionPrivateStruct *gi.Struct
var radioActionPrivateStructOnce sync.Once

func radioActionPrivateStructSet() {
	radioActionPrivateStructOnce.Do(func() {
		radioActionPrivateStruct = gi.StructNew("Gtk", "RadioActionPrivate")
	})
}

type RadioActionPrivate struct {
	native uintptr
}

var radioButtonAccessibleClassStruct *gi.Struct
var radioButtonAccessibleClassStructOnce sync.Once

func radioButtonAccessibleClassStructSet() {
	radioButtonAccessibleClassStructOnce.Do(func() {
		radioButtonAccessibleClassStruct = gi.StructNew("Gtk", "RadioButtonAccessibleClass")
	})
}

type RadioButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ToggleButtonAccessibleClass
}

var radioButtonAccessiblePrivateStruct *gi.Struct
var radioButtonAccessiblePrivateStructOnce sync.Once

func radioButtonAccessiblePrivateStructSet() {
	radioButtonAccessiblePrivateStructOnce.Do(func() {
		radioButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "RadioButtonAccessiblePrivate")
	})
}

type RadioButtonAccessiblePrivate struct {
	native uintptr
}

var radioButtonClassStruct *gi.Struct
var radioButtonClassStructOnce sync.Once

func radioButtonClassStructSet() {
	radioButtonClassStructOnce.Do(func() {
		radioButtonClassStruct = gi.StructNew("Gtk", "RadioButtonClass")
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

var radioButtonPrivateStruct *gi.Struct
var radioButtonPrivateStructOnce sync.Once

func radioButtonPrivateStructSet() {
	radioButtonPrivateStructOnce.Do(func() {
		radioButtonPrivateStruct = gi.StructNew("Gtk", "RadioButtonPrivate")
	})
}

type RadioButtonPrivate struct {
	native uintptr
}

var radioMenuItemAccessibleClassStruct *gi.Struct
var radioMenuItemAccessibleClassStructOnce sync.Once

func radioMenuItemAccessibleClassStructSet() {
	radioMenuItemAccessibleClassStructOnce.Do(func() {
		radioMenuItemAccessibleClassStruct = gi.StructNew("Gtk", "RadioMenuItemAccessibleClass")
	})
}

type RadioMenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *CheckMenuItemAccessibleClass
}

var radioMenuItemAccessiblePrivateStruct *gi.Struct
var radioMenuItemAccessiblePrivateStructOnce sync.Once

func radioMenuItemAccessiblePrivateStructSet() {
	radioMenuItemAccessiblePrivateStructOnce.Do(func() {
		radioMenuItemAccessiblePrivateStruct = gi.StructNew("Gtk", "RadioMenuItemAccessiblePrivate")
	})
}

type RadioMenuItemAccessiblePrivate struct {
	native uintptr
}

var radioMenuItemClassStruct *gi.Struct
var radioMenuItemClassStructOnce sync.Once

func radioMenuItemClassStructSet() {
	radioMenuItemClassStructOnce.Do(func() {
		radioMenuItemClassStruct = gi.StructNew("Gtk", "RadioMenuItemClass")
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

var radioMenuItemPrivateStruct *gi.Struct
var radioMenuItemPrivateStructOnce sync.Once

func radioMenuItemPrivateStructSet() {
	radioMenuItemPrivateStructOnce.Do(func() {
		radioMenuItemPrivateStruct = gi.StructNew("Gtk", "RadioMenuItemPrivate")
	})
}

type RadioMenuItemPrivate struct {
	native uintptr
}

var radioToolButtonClassStruct *gi.Struct
var radioToolButtonClassStructOnce sync.Once

func radioToolButtonClassStructSet() {
	radioToolButtonClassStructOnce.Do(func() {
		radioToolButtonClassStruct = gi.StructNew("Gtk", "RadioToolButtonClass")
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

var rangeAccessibleClassStruct *gi.Struct
var rangeAccessibleClassStructOnce sync.Once

func rangeAccessibleClassStructSet() {
	rangeAccessibleClassStructOnce.Do(func() {
		rangeAccessibleClassStruct = gi.StructNew("Gtk", "RangeAccessibleClass")
	})
}

type RangeAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var rangeAccessiblePrivateStruct *gi.Struct
var rangeAccessiblePrivateStructOnce sync.Once

func rangeAccessiblePrivateStructSet() {
	rangeAccessiblePrivateStructOnce.Do(func() {
		rangeAccessiblePrivateStruct = gi.StructNew("Gtk", "RangeAccessiblePrivate")
	})
}

type RangeAccessiblePrivate struct {
	native uintptr
}

var rangeClassStruct *gi.Struct
var rangeClassStructOnce sync.Once

func rangeClassStructSet() {
	rangeClassStructOnce.Do(func() {
		rangeClassStruct = gi.StructNew("Gtk", "RangeClass")
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

var rangePrivateStruct *gi.Struct
var rangePrivateStructOnce sync.Once

func rangePrivateStructSet() {
	rangePrivateStructOnce.Do(func() {
		rangePrivateStruct = gi.StructNew("Gtk", "RangePrivate")
	})
}

type RangePrivate struct {
	native uintptr
}

var rcContextStruct *gi.Struct
var rcContextStructOnce sync.Once

func rcContextStructSet() {
	rcContextStructOnce.Do(func() {
		rcContextStruct = gi.StructNew("Gtk", "RcContext")
	})
}

type RcContext struct {
	native uintptr
}

var rcPropertyStruct *gi.Struct
var rcPropertyStructOnce sync.Once

func rcPropertyStructSet() {
	rcPropertyStructOnce.Do(func() {
		rcPropertyStruct = gi.StructNew("Gtk", "RcProperty")
	})
}

type RcProperty struct {
	native uintptr
	// UNSUPPORTED : C value 'type_name' : no Go type for 'GLib.Quark'
	// UNSUPPORTED : C value 'property_name' : no Go type for 'GLib.Quark'
	Origin string
	// UNSUPPORTED : C value 'value' : no Go type for 'GObject.Value'
}

var rcStyleClassStruct *gi.Struct
var rcStyleClassStructOnce sync.Once

func rcStyleClassStructSet() {
	rcStyleClassStructOnce.Do(func() {
		rcStyleClassStruct = gi.StructNew("Gtk", "RcStyleClass")
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

var recentActionClassStruct *gi.Struct
var recentActionClassStructOnce sync.Once

func recentActionClassStructSet() {
	recentActionClassStructOnce.Do(func() {
		recentActionClassStruct = gi.StructNew("Gtk", "RecentActionClass")
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

var recentActionPrivateStruct *gi.Struct
var recentActionPrivateStructOnce sync.Once

func recentActionPrivateStructSet() {
	recentActionPrivateStructOnce.Do(func() {
		recentActionPrivateStruct = gi.StructNew("Gtk", "RecentActionPrivate")
	})
}

type RecentActionPrivate struct {
	native uintptr
}

var recentChooserDialogClassStruct *gi.Struct
var recentChooserDialogClassStructOnce sync.Once

func recentChooserDialogClassStructSet() {
	recentChooserDialogClassStructOnce.Do(func() {
		recentChooserDialogClassStruct = gi.StructNew("Gtk", "RecentChooserDialogClass")
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

var recentChooserDialogPrivateStruct *gi.Struct
var recentChooserDialogPrivateStructOnce sync.Once

func recentChooserDialogPrivateStructSet() {
	recentChooserDialogPrivateStructOnce.Do(func() {
		recentChooserDialogPrivateStruct = gi.StructNew("Gtk", "RecentChooserDialogPrivate")
	})
}

type RecentChooserDialogPrivate struct {
	native uintptr
}

var recentChooserIfaceStruct *gi.Struct
var recentChooserIfaceStructOnce sync.Once

func recentChooserIfaceStructSet() {
	recentChooserIfaceStructOnce.Do(func() {
		recentChooserIfaceStruct = gi.StructNew("Gtk", "RecentChooserIface")
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

var recentChooserMenuClassStruct *gi.Struct
var recentChooserMenuClassStructOnce sync.Once

func recentChooserMenuClassStructSet() {
	recentChooserMenuClassStructOnce.Do(func() {
		recentChooserMenuClassStruct = gi.StructNew("Gtk", "RecentChooserMenuClass")
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

var recentChooserMenuPrivateStruct *gi.Struct
var recentChooserMenuPrivateStructOnce sync.Once

func recentChooserMenuPrivateStructSet() {
	recentChooserMenuPrivateStructOnce.Do(func() {
		recentChooserMenuPrivateStruct = gi.StructNew("Gtk", "RecentChooserMenuPrivate")
	})
}

type RecentChooserMenuPrivate struct {
	native uintptr
}

var recentChooserWidgetClassStruct *gi.Struct
var recentChooserWidgetClassStructOnce sync.Once

func recentChooserWidgetClassStructSet() {
	recentChooserWidgetClassStructOnce.Do(func() {
		recentChooserWidgetClassStruct = gi.StructNew("Gtk", "RecentChooserWidgetClass")
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

var recentChooserWidgetPrivateStruct *gi.Struct
var recentChooserWidgetPrivateStructOnce sync.Once

func recentChooserWidgetPrivateStructSet() {
	recentChooserWidgetPrivateStructOnce.Do(func() {
		recentChooserWidgetPrivateStruct = gi.StructNew("Gtk", "RecentChooserWidgetPrivate")
	})
}

type RecentChooserWidgetPrivate struct {
	native uintptr
}

var recentDataStruct *gi.Struct
var recentDataStructOnce sync.Once

func recentDataStructSet() {
	recentDataStructOnce.Do(func() {
		recentDataStruct = gi.StructNew("Gtk", "RecentData")
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

var recentFilterInfoStruct *gi.Struct
var recentFilterInfoStructOnce sync.Once

func recentFilterInfoStructSet() {
	recentFilterInfoStructOnce.Do(func() {
		recentFilterInfoStruct = gi.StructNew("Gtk", "RecentFilterInfo")
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

var recentInfoStruct *gi.Struct
var recentInfoStructOnce sync.Once

func recentInfoStructSet() {
	recentInfoStructOnce.Do(func() {
		recentInfoStruct = gi.StructNew("Gtk", "RecentInfo")
	})
}

type RecentInfo struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_recent_info_create_app_info' : return type 'Gio.AppInfo' not supported

// UNSUPPORTED : C value 'gtk_recent_info_exists' : return type 'gboolean' not supported

var recentInfoGetAddedFunction *gi.Function
var recentInfoGetAddedFunction_Once sync.Once

func recentInfoGetAddedFunction_Set() {
	recentInfoGetAddedFunction_Once.Do(func() {
		recentInfoGetAddedFunction = gi.FunctionInvokerNew("Gtk", "get_added")
	})
}

var getAddedRecentInfoInvoker *gi.Function

// GetAdded is a representation of the C type gtk_recent_info_get_added.
func (recv *RecentInfo) GetAdded() int64 {
	recentInfoGetAddedFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetAddedFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

var recentInfoGetAgeFunction *gi.Function
var recentInfoGetAgeFunction_Once sync.Once

func recentInfoGetAgeFunction_Set() {
	recentInfoGetAgeFunction_Once.Do(func() {
		recentInfoGetAgeFunction = gi.FunctionInvokerNew("Gtk", "get_age")
	})
}

var getAgeRecentInfoInvoker *gi.Function

// GetAge is a representation of the C type gtk_recent_info_get_age.
func (recv *RecentInfo) GetAge() int32 {
	recentInfoGetAgeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetAgeFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_get_application_info' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_recent_info_get_applications' : parameter 'length' of type 'gsize' not supported

var recentInfoGetDescriptionFunction *gi.Function
var recentInfoGetDescriptionFunction_Once sync.Once

func recentInfoGetDescriptionFunction_Set() {
	recentInfoGetDescriptionFunction_Once.Do(func() {
		recentInfoGetDescriptionFunction = gi.FunctionInvokerNew("Gtk", "get_description")
	})
}

var getDescriptionRecentInfoInvoker *gi.Function

// GetDescription is a representation of the C type gtk_recent_info_get_description.
func (recv *RecentInfo) GetDescription() string {
	recentInfoGetDescriptionFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetDescriptionFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var recentInfoGetDisplayNameFunction *gi.Function
var recentInfoGetDisplayNameFunction_Once sync.Once

func recentInfoGetDisplayNameFunction_Set() {
	recentInfoGetDisplayNameFunction_Once.Do(func() {
		recentInfoGetDisplayNameFunction = gi.FunctionInvokerNew("Gtk", "get_display_name")
	})
}

var getDisplayNameRecentInfoInvoker *gi.Function

// GetDisplayName is a representation of the C type gtk_recent_info_get_display_name.
func (recv *RecentInfo) GetDisplayName() string {
	recentInfoGetDisplayNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetDisplayNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_get_gicon' : return type 'Gio.Icon' not supported

// UNSUPPORTED : C value 'gtk_recent_info_get_groups' : parameter 'length' of type 'gsize' not supported

// UNSUPPORTED : C value 'gtk_recent_info_get_icon' : return type 'GdkPixbuf.Pixbuf' not supported

var recentInfoGetMimeTypeFunction *gi.Function
var recentInfoGetMimeTypeFunction_Once sync.Once

func recentInfoGetMimeTypeFunction_Set() {
	recentInfoGetMimeTypeFunction_Once.Do(func() {
		recentInfoGetMimeTypeFunction = gi.FunctionInvokerNew("Gtk", "get_mime_type")
	})
}

var getMimeTypeRecentInfoInvoker *gi.Function

// GetMimeType is a representation of the C type gtk_recent_info_get_mime_type.
func (recv *RecentInfo) GetMimeType() string {
	recentInfoGetMimeTypeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetMimeTypeFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var recentInfoGetModifiedFunction *gi.Function
var recentInfoGetModifiedFunction_Once sync.Once

func recentInfoGetModifiedFunction_Set() {
	recentInfoGetModifiedFunction_Once.Do(func() {
		recentInfoGetModifiedFunction = gi.FunctionInvokerNew("Gtk", "get_modified")
	})
}

var getModifiedRecentInfoInvoker *gi.Function

// GetModified is a representation of the C type gtk_recent_info_get_modified.
func (recv *RecentInfo) GetModified() int64 {
	recentInfoGetModifiedFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetModifiedFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_get_private_hint' : return type 'gboolean' not supported

var recentInfoGetShortNameFunction *gi.Function
var recentInfoGetShortNameFunction_Once sync.Once

func recentInfoGetShortNameFunction_Set() {
	recentInfoGetShortNameFunction_Once.Do(func() {
		recentInfoGetShortNameFunction = gi.FunctionInvokerNew("Gtk", "get_short_name")
	})
}

var getShortNameRecentInfoInvoker *gi.Function

// GetShortName is a representation of the C type gtk_recent_info_get_short_name.
func (recv *RecentInfo) GetShortName() string {
	recentInfoGetShortNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetShortNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var recentInfoGetUriFunction *gi.Function
var recentInfoGetUriFunction_Once sync.Once

func recentInfoGetUriFunction_Set() {
	recentInfoGetUriFunction_Once.Do(func() {
		recentInfoGetUriFunction = gi.FunctionInvokerNew("Gtk", "get_uri")
	})
}

var getUriRecentInfoInvoker *gi.Function

// GetUri is a representation of the C type gtk_recent_info_get_uri.
func (recv *RecentInfo) GetUri() string {
	recentInfoGetUriFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetUriFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var recentInfoGetUriDisplayFunction *gi.Function
var recentInfoGetUriDisplayFunction_Once sync.Once

func recentInfoGetUriDisplayFunction_Set() {
	recentInfoGetUriDisplayFunction_Once.Do(func() {
		recentInfoGetUriDisplayFunction = gi.FunctionInvokerNew("Gtk", "get_uri_display")
	})
}

var getUriDisplayRecentInfoInvoker *gi.Function

// GetUriDisplay is a representation of the C type gtk_recent_info_get_uri_display.
func (recv *RecentInfo) GetUriDisplay() string {
	recentInfoGetUriDisplayFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetUriDisplayFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var recentInfoGetVisitedFunction *gi.Function
var recentInfoGetVisitedFunction_Once sync.Once

func recentInfoGetVisitedFunction_Set() {
	recentInfoGetVisitedFunction_Once.Do(func() {
		recentInfoGetVisitedFunction = gi.FunctionInvokerNew("Gtk", "get_visited")
	})
}

var getVisitedRecentInfoInvoker *gi.Function

// GetVisited is a representation of the C type gtk_recent_info_get_visited.
func (recv *RecentInfo) GetVisited() int64 {
	recentInfoGetVisitedFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoGetVisitedFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_has_application' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_recent_info_has_group' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_recent_info_is_local' : return type 'gboolean' not supported

var recentInfoLastApplicationFunction *gi.Function
var recentInfoLastApplicationFunction_Once sync.Once

func recentInfoLastApplicationFunction_Set() {
	recentInfoLastApplicationFunction_Once.Do(func() {
		recentInfoLastApplicationFunction = gi.FunctionInvokerNew("Gtk", "last_application")
	})
}

var lastApplicationRecentInfoInvoker *gi.Function

// LastApplication is a representation of the C type gtk_recent_info_last_application.
func (recv *RecentInfo) LastApplication() string {
	recentInfoLastApplicationFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoLastApplicationFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_match' : parameter 'info_b' of type 'RecentInfo' not supported

var recentInfoRefFunction *gi.Function
var recentInfoRefFunction_Once sync.Once

func recentInfoRefFunction_Set() {
	recentInfoRefFunction_Once.Do(func() {
		recentInfoRefFunction = gi.FunctionInvokerNew("Gtk", "ref")
	})
}

var refRecentInfoInvoker *gi.Function

// Ref is a representation of the C type gtk_recent_info_ref.
func (recv *RecentInfo) Ref() *RecentInfo {
	recentInfoRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := recentInfoRefFunction.Invoke(inArgs[:], nil)

	retGo := &RecentInfo{native: ret.Pointer()}

	return retGo
}

var recentInfoUnrefFunction *gi.Function
var recentInfoUnrefFunction_Once sync.Once

func recentInfoUnrefFunction_Set() {
	recentInfoUnrefFunction_Once.Do(func() {
		recentInfoUnrefFunction = gi.FunctionInvokerNew("Gtk", "unref")
	})
}

var unrefRecentInfoInvoker *gi.Function

// Unref is a representation of the C type gtk_recent_info_unref.
func (recv *RecentInfo) Unref() {
	recentInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	recentInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var recentManagerClassStruct *gi.Struct
var recentManagerClassStructOnce sync.Once

func recentManagerClassStructSet() {
	recentManagerClassStructOnce.Do(func() {
		recentManagerClassStruct = gi.StructNew("Gtk", "RecentManagerClass")
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

var recentManagerPrivateStruct *gi.Struct
var recentManagerPrivateStructOnce sync.Once

func recentManagerPrivateStructSet() {
	recentManagerPrivateStructOnce.Do(func() {
		recentManagerPrivateStruct = gi.StructNew("Gtk", "RecentManagerPrivate")
	})
}

type RecentManagerPrivate struct {
	native uintptr
}

var rendererCellAccessibleClassStruct *gi.Struct
var rendererCellAccessibleClassStructOnce sync.Once

func rendererCellAccessibleClassStructSet() {
	rendererCellAccessibleClassStructOnce.Do(func() {
		rendererCellAccessibleClassStruct = gi.StructNew("Gtk", "RendererCellAccessibleClass")
	})
}

type RendererCellAccessibleClass struct {
	native      uintptr
	ParentClass *CellAccessibleClass
}

var rendererCellAccessiblePrivateStruct *gi.Struct
var rendererCellAccessiblePrivateStructOnce sync.Once

func rendererCellAccessiblePrivateStructSet() {
	rendererCellAccessiblePrivateStructOnce.Do(func() {
		rendererCellAccessiblePrivateStruct = gi.StructNew("Gtk", "RendererCellAccessiblePrivate")
	})
}

type RendererCellAccessiblePrivate struct {
	native uintptr
}

var requestedSizeStruct *gi.Struct
var requestedSizeStructOnce sync.Once

func requestedSizeStructSet() {
	requestedSizeStructOnce.Do(func() {
		requestedSizeStruct = gi.StructNew("Gtk", "RequestedSize")
	})
}

type RequestedSize struct {
	native uintptr
	// UNSUPPORTED : C value 'data' : no Go type for 'gpointer'
	MinimumSize int32
	NaturalSize int32
}

var requisitionStruct *gi.Struct
var requisitionStructOnce sync.Once

func requisitionStructSet() {
	requisitionStructOnce.Do(func() {
		requisitionStruct = gi.StructNew("Gtk", "Requisition")
	})
}

type Requisition struct {
	native uintptr
	Width  int32
	Height int32
}

var requisitionNewFunction *gi.Function
var requisitionNewFunction_Once sync.Once

func requisitionNewFunction_Set() {
	requisitionNewFunction_Once.Do(func() {
		requisitionNewFunction = gi.FunctionInvokerNew("Gtk", "new")
	})
}

var newRequisitionInvoker *gi.Function

// RequisitionNew is a representation of the C type gtk_requisition_new.
func RequisitionNew() *Requisition {
	requisitionNewFunction_Set()

	ret := requisitionNewFunction.Invoke(nil, nil)

	retGo := &Requisition{native: ret.Pointer()}

	return retGo
}

var requisitionCopyFunction *gi.Function
var requisitionCopyFunction_Once sync.Once

func requisitionCopyFunction_Set() {
	requisitionCopyFunction_Once.Do(func() {
		requisitionCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyRequisitionInvoker *gi.Function

// Copy is a representation of the C type gtk_requisition_copy.
func (recv *Requisition) Copy() *Requisition {
	requisitionCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := requisitionCopyFunction.Invoke(inArgs[:], nil)

	retGo := &Requisition{native: ret.Pointer()}

	return retGo
}

var requisitionFreeFunction *gi.Function
var requisitionFreeFunction_Once sync.Once

func requisitionFreeFunction_Set() {
	requisitionFreeFunction_Once.Do(func() {
		requisitionFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeRequisitionInvoker *gi.Function

// Free is a representation of the C type gtk_requisition_free.
func (recv *Requisition) Free() {
	requisitionFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	requisitionFreeFunction.Invoke(inArgs[:], nil)

}

var revealerClassStruct *gi.Struct
var revealerClassStructOnce sync.Once

func revealerClassStructSet() {
	revealerClassStructOnce.Do(func() {
		revealerClassStruct = gi.StructNew("Gtk", "RevealerClass")
	})
}

type RevealerClass struct {
	native      uintptr
	ParentClass *BinClass
}

var scaleAccessibleClassStruct *gi.Struct
var scaleAccessibleClassStructOnce sync.Once

func scaleAccessibleClassStructSet() {
	scaleAccessibleClassStructOnce.Do(func() {
		scaleAccessibleClassStruct = gi.StructNew("Gtk", "ScaleAccessibleClass")
	})
}

type ScaleAccessibleClass struct {
	native      uintptr
	ParentClass *RangeAccessibleClass
}

var scaleAccessiblePrivateStruct *gi.Struct
var scaleAccessiblePrivateStructOnce sync.Once

func scaleAccessiblePrivateStructSet() {
	scaleAccessiblePrivateStructOnce.Do(func() {
		scaleAccessiblePrivateStruct = gi.StructNew("Gtk", "ScaleAccessiblePrivate")
	})
}

type ScaleAccessiblePrivate struct {
	native uintptr
}

var scaleButtonAccessibleClassStruct *gi.Struct
var scaleButtonAccessibleClassStructOnce sync.Once

func scaleButtonAccessibleClassStructSet() {
	scaleButtonAccessibleClassStructOnce.Do(func() {
		scaleButtonAccessibleClassStruct = gi.StructNew("Gtk", "ScaleButtonAccessibleClass")
	})
}

type ScaleButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var scaleButtonAccessiblePrivateStruct *gi.Struct
var scaleButtonAccessiblePrivateStructOnce sync.Once

func scaleButtonAccessiblePrivateStructSet() {
	scaleButtonAccessiblePrivateStructOnce.Do(func() {
		scaleButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "ScaleButtonAccessiblePrivate")
	})
}

type ScaleButtonAccessiblePrivate struct {
	native uintptr
}

var scaleButtonClassStruct *gi.Struct
var scaleButtonClassStructOnce sync.Once

func scaleButtonClassStructSet() {
	scaleButtonClassStructOnce.Do(func() {
		scaleButtonClassStruct = gi.StructNew("Gtk", "ScaleButtonClass")
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

var scaleButtonPrivateStruct *gi.Struct
var scaleButtonPrivateStructOnce sync.Once

func scaleButtonPrivateStructSet() {
	scaleButtonPrivateStructOnce.Do(func() {
		scaleButtonPrivateStruct = gi.StructNew("Gtk", "ScaleButtonPrivate")
	})
}

type ScaleButtonPrivate struct {
	native uintptr
}

var scaleClassStruct *gi.Struct
var scaleClassStructOnce sync.Once

func scaleClassStructSet() {
	scaleClassStructOnce.Do(func() {
		scaleClassStruct = gi.StructNew("Gtk", "ScaleClass")
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

var scalePrivateStruct *gi.Struct
var scalePrivateStructOnce sync.Once

func scalePrivateStructSet() {
	scalePrivateStructOnce.Do(func() {
		scalePrivateStruct = gi.StructNew("Gtk", "ScalePrivate")
	})
}

type ScalePrivate struct {
	native uintptr
}

var scrollableInterfaceStruct *gi.Struct
var scrollableInterfaceStructOnce sync.Once

func scrollableInterfaceStructSet() {
	scrollableInterfaceStructOnce.Do(func() {
		scrollableInterfaceStruct = gi.StructNew("Gtk", "ScrollableInterface")
	})
}

type ScrollableInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_border' : missing Type
}

var scrollbarClassStruct *gi.Struct
var scrollbarClassStructOnce sync.Once

func scrollbarClassStructSet() {
	scrollbarClassStructOnce.Do(func() {
		scrollbarClassStruct = gi.StructNew("Gtk", "ScrollbarClass")
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

var scrolledWindowAccessibleClassStruct *gi.Struct
var scrolledWindowAccessibleClassStructOnce sync.Once

func scrolledWindowAccessibleClassStructSet() {
	scrolledWindowAccessibleClassStructOnce.Do(func() {
		scrolledWindowAccessibleClassStruct = gi.StructNew("Gtk", "ScrolledWindowAccessibleClass")
	})
}

type ScrolledWindowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var scrolledWindowAccessiblePrivateStruct *gi.Struct
var scrolledWindowAccessiblePrivateStructOnce sync.Once

func scrolledWindowAccessiblePrivateStructSet() {
	scrolledWindowAccessiblePrivateStructOnce.Do(func() {
		scrolledWindowAccessiblePrivateStruct = gi.StructNew("Gtk", "ScrolledWindowAccessiblePrivate")
	})
}

type ScrolledWindowAccessiblePrivate struct {
	native uintptr
}

var scrolledWindowClassStruct *gi.Struct
var scrolledWindowClassStructOnce sync.Once

func scrolledWindowClassStructSet() {
	scrolledWindowClassStructOnce.Do(func() {
		scrolledWindowClassStruct = gi.StructNew("Gtk", "ScrolledWindowClass")
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

var scrolledWindowPrivateStruct *gi.Struct
var scrolledWindowPrivateStructOnce sync.Once

func scrolledWindowPrivateStructSet() {
	scrolledWindowPrivateStructOnce.Do(func() {
		scrolledWindowPrivateStruct = gi.StructNew("Gtk", "ScrolledWindowPrivate")
	})
}

type ScrolledWindowPrivate struct {
	native uintptr
}

var searchBarClassStruct *gi.Struct
var searchBarClassStructOnce sync.Once

func searchBarClassStructSet() {
	searchBarClassStructOnce.Do(func() {
		searchBarClassStruct = gi.StructNew("Gtk", "SearchBarClass")
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

var searchEntryClassStruct *gi.Struct
var searchEntryClassStructOnce sync.Once

func searchEntryClassStructSet() {
	searchEntryClassStructOnce.Do(func() {
		searchEntryClassStruct = gi.StructNew("Gtk", "SearchEntryClass")
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

var selectionDataStruct *gi.Struct
var selectionDataStructOnce sync.Once

func selectionDataStructSet() {
	selectionDataStructOnce.Do(func() {
		selectionDataStruct = gi.StructNew("Gtk", "SelectionData")
	})
}

type SelectionData struct {
	native uintptr
}

var selectionDataCopyFunction *gi.Function
var selectionDataCopyFunction_Once sync.Once

func selectionDataCopyFunction_Set() {
	selectionDataCopyFunction_Once.Do(func() {
		selectionDataCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copySelectionDataInvoker *gi.Function

// Copy is a representation of the C type gtk_selection_data_copy.
func (recv *SelectionData) Copy() *SelectionData {
	selectionDataCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := selectionDataCopyFunction.Invoke(inArgs[:], nil)

	retGo := &SelectionData{native: ret.Pointer()}

	return retGo
}

var selectionDataFreeFunction *gi.Function
var selectionDataFreeFunction_Once sync.Once

func selectionDataFreeFunction_Set() {
	selectionDataFreeFunction_Once.Do(func() {
		selectionDataFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeSelectionDataInvoker *gi.Function

// Free is a representation of the C type gtk_selection_data_free.
func (recv *SelectionData) Free() {
	selectionDataFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	selectionDataFreeFunction.Invoke(inArgs[:], nil)

}

var selectionDataGetDataFunction *gi.Function
var selectionDataGetDataFunction_Once sync.Once

func selectionDataGetDataFunction_Set() {
	selectionDataGetDataFunction_Once.Do(func() {
		selectionDataGetDataFunction = gi.FunctionInvokerNew("Gtk", "get_data")
	})
}

var getDataSelectionDataInvoker *gi.Function

// GetData is a representation of the C type gtk_selection_data_get_data.
func (recv *SelectionData) GetData() {
	selectionDataGetDataFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	selectionDataGetDataFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_selection_data_get_data_type' : return type 'Gdk.Atom' not supported

var selectionDataGetDataWithLengthFunction *gi.Function
var selectionDataGetDataWithLengthFunction_Once sync.Once

func selectionDataGetDataWithLengthFunction_Set() {
	selectionDataGetDataWithLengthFunction_Once.Do(func() {
		selectionDataGetDataWithLengthFunction = gi.FunctionInvokerNew("Gtk", "get_data_with_length")
	})
}

var getDataWithLengthSelectionDataInvoker *gi.Function

// GetDataWithLength is a representation of the C type gtk_selection_data_get_data_with_length.
func (recv *SelectionData) GetDataWithLength() int32 {
	selectionDataGetDataWithLengthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	selectionDataGetDataWithLengthFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'gtk_selection_data_get_display' : return type 'Gdk.Display' not supported

var selectionDataGetFormatFunction *gi.Function
var selectionDataGetFormatFunction_Once sync.Once

func selectionDataGetFormatFunction_Set() {
	selectionDataGetFormatFunction_Once.Do(func() {
		selectionDataGetFormatFunction = gi.FunctionInvokerNew("Gtk", "get_format")
	})
}

var getFormatSelectionDataInvoker *gi.Function

// GetFormat is a representation of the C type gtk_selection_data_get_format.
func (recv *SelectionData) GetFormat() int32 {
	selectionDataGetFormatFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := selectionDataGetFormatFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var selectionDataGetLengthFunction *gi.Function
var selectionDataGetLengthFunction_Once sync.Once

func selectionDataGetLengthFunction_Set() {
	selectionDataGetLengthFunction_Once.Do(func() {
		selectionDataGetLengthFunction = gi.FunctionInvokerNew("Gtk", "get_length")
	})
}

var getLengthSelectionDataInvoker *gi.Function

// GetLength is a representation of the C type gtk_selection_data_get_length.
func (recv *SelectionData) GetLength() int32 {
	selectionDataGetLengthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := selectionDataGetLengthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_selection_data_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_selection_data_get_selection' : return type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_data_get_target' : return type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_data_get_targets' : parameter 'targets' has no type

var selectionDataGetTextFunction *gi.Function
var selectionDataGetTextFunction_Once sync.Once

func selectionDataGetTextFunction_Set() {
	selectionDataGetTextFunction_Once.Do(func() {
		selectionDataGetTextFunction = gi.FunctionInvokerNew("Gtk", "get_text")
	})
}

var getTextSelectionDataInvoker *gi.Function

// GetText is a representation of the C type gtk_selection_data_get_text.
func (recv *SelectionData) GetText() string {
	selectionDataGetTextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := selectionDataGetTextFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var selectionDataGetUrisFunction *gi.Function
var selectionDataGetUrisFunction_Once sync.Once

func selectionDataGetUrisFunction_Set() {
	selectionDataGetUrisFunction_Once.Do(func() {
		selectionDataGetUrisFunction = gi.FunctionInvokerNew("Gtk", "get_uris")
	})
}

var getUrisSelectionDataInvoker *gi.Function

// GetUris is a representation of the C type gtk_selection_data_get_uris.
func (recv *SelectionData) GetUris() {
	selectionDataGetUrisFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	selectionDataGetUrisFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_selection_data_set' : parameter 'type' of type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_data_set_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_selection_data_set_text' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_selection_data_set_uris' : parameter 'uris' has no type

// UNSUPPORTED : C value 'gtk_selection_data_targets_include_image' : parameter 'writable' of type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_selection_data_targets_include_rich_text' : parameter 'buffer' of type 'TextBuffer' not supported

// UNSUPPORTED : C value 'gtk_selection_data_targets_include_text' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_selection_data_targets_include_uri' : return type 'gboolean' not supported

var separatorClassStruct *gi.Struct
var separatorClassStructOnce sync.Once

func separatorClassStructSet() {
	separatorClassStructOnce.Do(func() {
		separatorClassStruct = gi.StructNew("Gtk", "SeparatorClass")
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

var separatorMenuItemClassStruct *gi.Struct
var separatorMenuItemClassStructOnce sync.Once

func separatorMenuItemClassStructSet() {
	separatorMenuItemClassStructOnce.Do(func() {
		separatorMenuItemClassStruct = gi.StructNew("Gtk", "SeparatorMenuItemClass")
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

var separatorPrivateStruct *gi.Struct
var separatorPrivateStructOnce sync.Once

func separatorPrivateStructSet() {
	separatorPrivateStructOnce.Do(func() {
		separatorPrivateStruct = gi.StructNew("Gtk", "SeparatorPrivate")
	})
}

type SeparatorPrivate struct {
	native uintptr
}

var separatorToolItemClassStruct *gi.Struct
var separatorToolItemClassStructOnce sync.Once

func separatorToolItemClassStructSet() {
	separatorToolItemClassStructOnce.Do(func() {
		separatorToolItemClassStruct = gi.StructNew("Gtk", "SeparatorToolItemClass")
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

var separatorToolItemPrivateStruct *gi.Struct
var separatorToolItemPrivateStructOnce sync.Once

func separatorToolItemPrivateStructSet() {
	separatorToolItemPrivateStructOnce.Do(func() {
		separatorToolItemPrivateStruct = gi.StructNew("Gtk", "SeparatorToolItemPrivate")
	})
}

type SeparatorToolItemPrivate struct {
	native uintptr
}

var settingsClassStruct *gi.Struct
var settingsClassStructOnce sync.Once

func settingsClassStructSet() {
	settingsClassStructOnce.Do(func() {
		settingsClassStruct = gi.StructNew("Gtk", "SettingsClass")
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

var settingsPrivateStruct *gi.Struct
var settingsPrivateStructOnce sync.Once

func settingsPrivateStructSet() {
	settingsPrivateStructOnce.Do(func() {
		settingsPrivateStruct = gi.StructNew("Gtk", "SettingsPrivate")
	})
}

type SettingsPrivate struct {
	native uintptr
}

var settingsValueStruct *gi.Struct
var settingsValueStructOnce sync.Once

func settingsValueStructSet() {
	settingsValueStructOnce.Do(func() {
		settingsValueStruct = gi.StructNew("Gtk", "SettingsValue")
	})
}

type SettingsValue struct {
	native uintptr
	Origin string
	// UNSUPPORTED : C value 'value' : no Go type for 'GObject.Value'
}

var shortcutLabelClassStruct *gi.Struct
var shortcutLabelClassStructOnce sync.Once

func shortcutLabelClassStructSet() {
	shortcutLabelClassStructOnce.Do(func() {
		shortcutLabelClassStruct = gi.StructNew("Gtk", "ShortcutLabelClass")
	})
}

type ShortcutLabelClass struct {
	native uintptr
}

var shortcutsGroupClassStruct *gi.Struct
var shortcutsGroupClassStructOnce sync.Once

func shortcutsGroupClassStructSet() {
	shortcutsGroupClassStructOnce.Do(func() {
		shortcutsGroupClassStruct = gi.StructNew("Gtk", "ShortcutsGroupClass")
	})
}

type ShortcutsGroupClass struct {
	native uintptr
}

var shortcutsSectionClassStruct *gi.Struct
var shortcutsSectionClassStructOnce sync.Once

func shortcutsSectionClassStructSet() {
	shortcutsSectionClassStructOnce.Do(func() {
		shortcutsSectionClassStruct = gi.StructNew("Gtk", "ShortcutsSectionClass")
	})
}

type ShortcutsSectionClass struct {
	native uintptr
}

var shortcutsShortcutClassStruct *gi.Struct
var shortcutsShortcutClassStructOnce sync.Once

func shortcutsShortcutClassStructSet() {
	shortcutsShortcutClassStructOnce.Do(func() {
		shortcutsShortcutClassStruct = gi.StructNew("Gtk", "ShortcutsShortcutClass")
	})
}

type ShortcutsShortcutClass struct {
	native uintptr
}

var shortcutsWindowClassStruct *gi.Struct
var shortcutsWindowClassStructOnce sync.Once

func shortcutsWindowClassStructSet() {
	shortcutsWindowClassStructOnce.Do(func() {
		shortcutsWindowClassStruct = gi.StructNew("Gtk", "ShortcutsWindowClass")
	})
}

type ShortcutsWindowClass struct {
	native      uintptr
	ParentClass *WindowClass
	// UNSUPPORTED : C value 'close' : missing Type
	// UNSUPPORTED : C value 'search' : missing Type
}

var sizeGroupClassStruct *gi.Struct
var sizeGroupClassStructOnce sync.Once

func sizeGroupClassStructSet() {
	sizeGroupClassStructOnce.Do(func() {
		sizeGroupClassStruct = gi.StructNew("Gtk", "SizeGroupClass")
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

var sizeGroupPrivateStruct *gi.Struct
var sizeGroupPrivateStructOnce sync.Once

func sizeGroupPrivateStructSet() {
	sizeGroupPrivateStructOnce.Do(func() {
		sizeGroupPrivateStruct = gi.StructNew("Gtk", "SizeGroupPrivate")
	})
}

type SizeGroupPrivate struct {
	native uintptr
}

var socketClassStruct *gi.Struct
var socketClassStructOnce sync.Once

func socketClassStructSet() {
	socketClassStructOnce.Do(func() {
		socketClassStruct = gi.StructNew("Gtk", "SocketClass")
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

var socketPrivateStruct *gi.Struct
var socketPrivateStructOnce sync.Once

func socketPrivateStructSet() {
	socketPrivateStructOnce.Do(func() {
		socketPrivateStruct = gi.StructNew("Gtk", "SocketPrivate")
	})
}

type SocketPrivate struct {
	native uintptr
}

var spinButtonAccessibleClassStruct *gi.Struct
var spinButtonAccessibleClassStructOnce sync.Once

func spinButtonAccessibleClassStructSet() {
	spinButtonAccessibleClassStructOnce.Do(func() {
		spinButtonAccessibleClassStruct = gi.StructNew("Gtk", "SpinButtonAccessibleClass")
	})
}

type SpinButtonAccessibleClass struct {
	native      uintptr
	ParentClass *EntryAccessibleClass
}

var spinButtonAccessiblePrivateStruct *gi.Struct
var spinButtonAccessiblePrivateStructOnce sync.Once

func spinButtonAccessiblePrivateStructSet() {
	spinButtonAccessiblePrivateStructOnce.Do(func() {
		spinButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "SpinButtonAccessiblePrivate")
	})
}

type SpinButtonAccessiblePrivate struct {
	native uintptr
}

var spinButtonClassStruct *gi.Struct
var spinButtonClassStructOnce sync.Once

func spinButtonClassStructSet() {
	spinButtonClassStructOnce.Do(func() {
		spinButtonClassStruct = gi.StructNew("Gtk", "SpinButtonClass")
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

var spinButtonPrivateStruct *gi.Struct
var spinButtonPrivateStructOnce sync.Once

func spinButtonPrivateStructSet() {
	spinButtonPrivateStructOnce.Do(func() {
		spinButtonPrivateStruct = gi.StructNew("Gtk", "SpinButtonPrivate")
	})
}

type SpinButtonPrivate struct {
	native uintptr
}

var spinnerAccessibleClassStruct *gi.Struct
var spinnerAccessibleClassStructOnce sync.Once

func spinnerAccessibleClassStructSet() {
	spinnerAccessibleClassStructOnce.Do(func() {
		spinnerAccessibleClassStruct = gi.StructNew("Gtk", "SpinnerAccessibleClass")
	})
}

type SpinnerAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var spinnerAccessiblePrivateStruct *gi.Struct
var spinnerAccessiblePrivateStructOnce sync.Once

func spinnerAccessiblePrivateStructSet() {
	spinnerAccessiblePrivateStructOnce.Do(func() {
		spinnerAccessiblePrivateStruct = gi.StructNew("Gtk", "SpinnerAccessiblePrivate")
	})
}

type SpinnerAccessiblePrivate struct {
	native uintptr
}

var spinnerClassStruct *gi.Struct
var spinnerClassStructOnce sync.Once

func spinnerClassStructSet() {
	spinnerClassStructOnce.Do(func() {
		spinnerClassStruct = gi.StructNew("Gtk", "SpinnerClass")
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

var spinnerPrivateStruct *gi.Struct
var spinnerPrivateStructOnce sync.Once

func spinnerPrivateStructSet() {
	spinnerPrivateStructOnce.Do(func() {
		spinnerPrivateStruct = gi.StructNew("Gtk", "SpinnerPrivate")
	})
}

type SpinnerPrivate struct {
	native uintptr
}

var stackAccessibleClassStruct *gi.Struct
var stackAccessibleClassStructOnce sync.Once

func stackAccessibleClassStructSet() {
	stackAccessibleClassStructOnce.Do(func() {
		stackAccessibleClassStruct = gi.StructNew("Gtk", "StackAccessibleClass")
	})
}

type StackAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var stackClassStruct *gi.Struct
var stackClassStructOnce sync.Once

func stackClassStructSet() {
	stackClassStructOnce.Do(func() {
		stackClassStruct = gi.StructNew("Gtk", "StackClass")
	})
}

type StackClass struct {
	native      uintptr
	ParentClass *ContainerClass
}

var stackSidebarClassStruct *gi.Struct
var stackSidebarClassStructOnce sync.Once

func stackSidebarClassStructSet() {
	stackSidebarClassStructOnce.Do(func() {
		stackSidebarClassStruct = gi.StructNew("Gtk", "StackSidebarClass")
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

var stackSidebarPrivateStruct *gi.Struct
var stackSidebarPrivateStructOnce sync.Once

func stackSidebarPrivateStructSet() {
	stackSidebarPrivateStructOnce.Do(func() {
		stackSidebarPrivateStruct = gi.StructNew("Gtk", "StackSidebarPrivate")
	})
}

type StackSidebarPrivate struct {
	native uintptr
}

var stackSwitcherClassStruct *gi.Struct
var stackSwitcherClassStructOnce sync.Once

func stackSwitcherClassStructSet() {
	stackSwitcherClassStructOnce.Do(func() {
		stackSwitcherClassStruct = gi.StructNew("Gtk", "StackSwitcherClass")
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

var statusIconClassStruct *gi.Struct
var statusIconClassStructOnce sync.Once

func statusIconClassStructSet() {
	statusIconClassStructOnce.Do(func() {
		statusIconClassStruct = gi.StructNew("Gtk", "StatusIconClass")
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

var statusIconPrivateStruct *gi.Struct
var statusIconPrivateStructOnce sync.Once

func statusIconPrivateStructSet() {
	statusIconPrivateStructOnce.Do(func() {
		statusIconPrivateStruct = gi.StructNew("Gtk", "StatusIconPrivate")
	})
}

type StatusIconPrivate struct {
	native uintptr
}

var statusbarAccessibleClassStruct *gi.Struct
var statusbarAccessibleClassStructOnce sync.Once

func statusbarAccessibleClassStructSet() {
	statusbarAccessibleClassStructOnce.Do(func() {
		statusbarAccessibleClassStruct = gi.StructNew("Gtk", "StatusbarAccessibleClass")
	})
}

type StatusbarAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var statusbarAccessiblePrivateStruct *gi.Struct
var statusbarAccessiblePrivateStructOnce sync.Once

func statusbarAccessiblePrivateStructSet() {
	statusbarAccessiblePrivateStructOnce.Do(func() {
		statusbarAccessiblePrivateStruct = gi.StructNew("Gtk", "StatusbarAccessiblePrivate")
	})
}

type StatusbarAccessiblePrivate struct {
	native uintptr
}

var statusbarClassStruct *gi.Struct
var statusbarClassStructOnce sync.Once

func statusbarClassStructSet() {
	statusbarClassStructOnce.Do(func() {
		statusbarClassStruct = gi.StructNew("Gtk", "StatusbarClass")
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

var statusbarPrivateStruct *gi.Struct
var statusbarPrivateStructOnce sync.Once

func statusbarPrivateStructSet() {
	statusbarPrivateStructOnce.Do(func() {
		statusbarPrivateStruct = gi.StructNew("Gtk", "StatusbarPrivate")
	})
}

type StatusbarPrivate struct {
	native uintptr
}

var stockItemStruct *gi.Struct
var stockItemStructOnce sync.Once

func stockItemStructSet() {
	stockItemStructOnce.Do(func() {
		stockItemStruct = gi.StructNew("Gtk", "StockItem")
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

var stockItemCopyFunction *gi.Function
var stockItemCopyFunction_Once sync.Once

func stockItemCopyFunction_Set() {
	stockItemCopyFunction_Once.Do(func() {
		stockItemCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyStockItemInvoker *gi.Function

// Copy is a representation of the C type gtk_stock_item_copy.
func (recv *StockItem) Copy() *StockItem {
	stockItemCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := stockItemCopyFunction.Invoke(inArgs[:], nil)

	retGo := &StockItem{native: ret.Pointer()}

	return retGo
}

var stockItemFreeFunction *gi.Function
var stockItemFreeFunction_Once sync.Once

func stockItemFreeFunction_Set() {
	stockItemFreeFunction_Once.Do(func() {
		stockItemFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeStockItemInvoker *gi.Function

// Free is a representation of the C type gtk_stock_item_free.
func (recv *StockItem) Free() {
	stockItemFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	stockItemFreeFunction.Invoke(inArgs[:], nil)

}

var styleClassStruct *gi.Struct
var styleClassStructOnce sync.Once

func styleClassStructSet() {
	styleClassStructOnce.Do(func() {
		styleClassStruct = gi.StructNew("Gtk", "StyleClass")
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

var styleContextClassStruct *gi.Struct
var styleContextClassStructOnce sync.Once

func styleContextClassStructSet() {
	styleContextClassStructOnce.Do(func() {
		styleContextClassStruct = gi.StructNew("Gtk", "StyleContextClass")
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

var styleContextPrivateStruct *gi.Struct
var styleContextPrivateStructOnce sync.Once

func styleContextPrivateStructSet() {
	styleContextPrivateStructOnce.Do(func() {
		styleContextPrivateStruct = gi.StructNew("Gtk", "StyleContextPrivate")
	})
}

type StyleContextPrivate struct {
	native uintptr
}

var stylePropertiesClassStruct *gi.Struct
var stylePropertiesClassStructOnce sync.Once

func stylePropertiesClassStructSet() {
	stylePropertiesClassStructOnce.Do(func() {
		stylePropertiesClassStruct = gi.StructNew("Gtk", "StylePropertiesClass")
	})
}

type StylePropertiesClass struct {
	native uintptr
	// UNSUPPORTED : C value '_gtk_reserved1' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved2' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved3' : missing Type
	// UNSUPPORTED : C value '_gtk_reserved4' : missing Type
}

var stylePropertiesPrivateStruct *gi.Struct
var stylePropertiesPrivateStructOnce sync.Once

func stylePropertiesPrivateStructSet() {
	stylePropertiesPrivateStructOnce.Do(func() {
		stylePropertiesPrivateStruct = gi.StructNew("Gtk", "StylePropertiesPrivate")
	})
}

type StylePropertiesPrivate struct {
	native uintptr
}

var styleProviderIfaceStruct *gi.Struct
var styleProviderIfaceStructOnce sync.Once

func styleProviderIfaceStructSet() {
	styleProviderIfaceStructOnce.Do(func() {
		styleProviderIfaceStruct = gi.StructNew("Gtk", "StyleProviderIface")
	})
}

type StyleProviderIface struct {
	native uintptr
	// UNSUPPORTED : C value 'get_style' : missing Type
	// UNSUPPORTED : C value 'get_style_property' : missing Type
	// UNSUPPORTED : C value 'get_icon_factory' : missing Type
}

var switchAccessibleClassStruct *gi.Struct
var switchAccessibleClassStructOnce sync.Once

func switchAccessibleClassStructSet() {
	switchAccessibleClassStructOnce.Do(func() {
		switchAccessibleClassStruct = gi.StructNew("Gtk", "SwitchAccessibleClass")
	})
}

type SwitchAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var switchAccessiblePrivateStruct *gi.Struct
var switchAccessiblePrivateStructOnce sync.Once

func switchAccessiblePrivateStructSet() {
	switchAccessiblePrivateStructOnce.Do(func() {
		switchAccessiblePrivateStruct = gi.StructNew("Gtk", "SwitchAccessiblePrivate")
	})
}

type SwitchAccessiblePrivate struct {
	native uintptr
}

var switchClassStruct *gi.Struct
var switchClassStructOnce sync.Once

func switchClassStructSet() {
	switchClassStructOnce.Do(func() {
		switchClassStruct = gi.StructNew("Gtk", "SwitchClass")
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

var switchPrivateStruct *gi.Struct
var switchPrivateStructOnce sync.Once

func switchPrivateStructSet() {
	switchPrivateStructOnce.Do(func() {
		switchPrivateStruct = gi.StructNew("Gtk", "SwitchPrivate")
	})
}

type SwitchPrivate struct {
	native uintptr
}

var symbolicColorStruct *gi.Struct
var symbolicColorStructOnce sync.Once

func symbolicColorStructSet() {
	symbolicColorStructOnce.Do(func() {
		symbolicColorStruct = gi.StructNew("Gtk", "SymbolicColor")
	})
}

type SymbolicColor struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_symbolic_color_new_alpha' : parameter 'color' of type 'SymbolicColor' not supported

// UNSUPPORTED : C value 'gtk_symbolic_color_new_literal' : parameter 'color' of type 'Gdk.RGBA' not supported

// UNSUPPORTED : C value 'gtk_symbolic_color_new_mix' : parameter 'color1' of type 'SymbolicColor' not supported

var symbolicColorNewNameFunction *gi.Function
var symbolicColorNewNameFunction_Once sync.Once

func symbolicColorNewNameFunction_Set() {
	symbolicColorNewNameFunction_Once.Do(func() {
		symbolicColorNewNameFunction = gi.FunctionInvokerNew("Gtk", "new_name")
	})
}

var newNameSymbolicColorInvoker *gi.Function

// SymbolicColorNewName is a representation of the C type gtk_symbolic_color_new_name.
func SymbolicColorNewName(name string) *SymbolicColor {
	symbolicColorNewNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	ret := symbolicColorNewNameFunction.Invoke(inArgs[:], nil)

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_symbolic_color_new_shade' : parameter 'color' of type 'SymbolicColor' not supported

var symbolicColorNewWin32Function *gi.Function
var symbolicColorNewWin32Function_Once sync.Once

func symbolicColorNewWin32Function_Set() {
	symbolicColorNewWin32Function_Once.Do(func() {
		symbolicColorNewWin32Function = gi.FunctionInvokerNew("Gtk", "new_win32")
	})
}

var newWin32SymbolicColorInvoker *gi.Function

// SymbolicColorNewWin32 is a representation of the C type gtk_symbolic_color_new_win32.
func SymbolicColorNewWin32(themeClass string, id int32) *SymbolicColor {
	symbolicColorNewWin32Function_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetString(themeClass)
	inArgs[1].SetInt32(id)

	ret := symbolicColorNewWin32Function.Invoke(inArgs[:], nil)

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

var symbolicColorRefFunction *gi.Function
var symbolicColorRefFunction_Once sync.Once

func symbolicColorRefFunction_Set() {
	symbolicColorRefFunction_Once.Do(func() {
		symbolicColorRefFunction = gi.FunctionInvokerNew("Gtk", "ref")
	})
}

var refSymbolicColorInvoker *gi.Function

// Ref is a representation of the C type gtk_symbolic_color_ref.
func (recv *SymbolicColor) Ref() *SymbolicColor {
	symbolicColorRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := symbolicColorRefFunction.Invoke(inArgs[:], nil)

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_symbolic_color_resolve' : parameter 'props' of type 'StyleProperties' not supported

var symbolicColorToStringFunction *gi.Function
var symbolicColorToStringFunction_Once sync.Once

func symbolicColorToStringFunction_Set() {
	symbolicColorToStringFunction_Once.Do(func() {
		symbolicColorToStringFunction = gi.FunctionInvokerNew("Gtk", "to_string")
	})
}

var toStringSymbolicColorInvoker *gi.Function

// ToString is a representation of the C type gtk_symbolic_color_to_string.
func (recv *SymbolicColor) ToString() string {
	symbolicColorToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := symbolicColorToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var symbolicColorUnrefFunction *gi.Function
var symbolicColorUnrefFunction_Once sync.Once

func symbolicColorUnrefFunction_Set() {
	symbolicColorUnrefFunction_Once.Do(func() {
		symbolicColorUnrefFunction = gi.FunctionInvokerNew("Gtk", "unref")
	})
}

var unrefSymbolicColorInvoker *gi.Function

// Unref is a representation of the C type gtk_symbolic_color_unref.
func (recv *SymbolicColor) Unref() {
	symbolicColorUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	symbolicColorUnrefFunction.Invoke(inArgs[:], nil)

}

var tableChildStruct *gi.Struct
var tableChildStructOnce sync.Once

func tableChildStructSet() {
	tableChildStructOnce.Do(func() {
		tableChildStruct = gi.StructNew("Gtk", "TableChild")
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

var tableClassStruct *gi.Struct
var tableClassStructOnce sync.Once

func tableClassStructSet() {
	tableClassStructOnce.Do(func() {
		tableClassStruct = gi.StructNew("Gtk", "TableClass")
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

var tablePrivateStruct *gi.Struct
var tablePrivateStructOnce sync.Once

func tablePrivateStructSet() {
	tablePrivateStructOnce.Do(func() {
		tablePrivateStruct = gi.StructNew("Gtk", "TablePrivate")
	})
}

type TablePrivate struct {
	native uintptr
}

var tableRowColStruct *gi.Struct
var tableRowColStructOnce sync.Once

func tableRowColStructSet() {
	tableRowColStructOnce.Do(func() {
		tableRowColStruct = gi.StructNew("Gtk", "TableRowCol")
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

var targetEntryStruct *gi.Struct
var targetEntryStructOnce sync.Once

func targetEntryStructSet() {
	targetEntryStructOnce.Do(func() {
		targetEntryStruct = gi.StructNew("Gtk", "TargetEntry")
	})
}

type TargetEntry struct {
	native uintptr
	Target string
	Flags  uint32
	Info   uint32
}

var targetEntryNewFunction *gi.Function
var targetEntryNewFunction_Once sync.Once

func targetEntryNewFunction_Set() {
	targetEntryNewFunction_Once.Do(func() {
		targetEntryNewFunction = gi.FunctionInvokerNew("Gtk", "new")
	})
}

var newTargetEntryInvoker *gi.Function

// TargetEntryNew is a representation of the C type gtk_target_entry_new.
func TargetEntryNew(target string, flags uint32, info uint32) *TargetEntry {
	targetEntryNewFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetString(target)
	inArgs[1].SetUint32(flags)
	inArgs[2].SetUint32(info)

	ret := targetEntryNewFunction.Invoke(inArgs[:], nil)

	retGo := &TargetEntry{native: ret.Pointer()}

	return retGo
}

var targetEntryCopyFunction *gi.Function
var targetEntryCopyFunction_Once sync.Once

func targetEntryCopyFunction_Set() {
	targetEntryCopyFunction_Once.Do(func() {
		targetEntryCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyTargetEntryInvoker *gi.Function

// Copy is a representation of the C type gtk_target_entry_copy.
func (recv *TargetEntry) Copy() *TargetEntry {
	targetEntryCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := targetEntryCopyFunction.Invoke(inArgs[:], nil)

	retGo := &TargetEntry{native: ret.Pointer()}

	return retGo
}

var targetEntryFreeFunction *gi.Function
var targetEntryFreeFunction_Once sync.Once

func targetEntryFreeFunction_Set() {
	targetEntryFreeFunction_Once.Do(func() {
		targetEntryFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeTargetEntryInvoker *gi.Function

// Free is a representation of the C type gtk_target_entry_free.
func (recv *TargetEntry) Free() {
	targetEntryFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	targetEntryFreeFunction.Invoke(inArgs[:], nil)

}

var targetListStruct *gi.Struct
var targetListStructOnce sync.Once

func targetListStructSet() {
	targetListStructOnce.Do(func() {
		targetListStruct = gi.StructNew("Gtk", "TargetList")
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

var targetListAddTextTargetsFunction *gi.Function
var targetListAddTextTargetsFunction_Once sync.Once

func targetListAddTextTargetsFunction_Set() {
	targetListAddTextTargetsFunction_Once.Do(func() {
		targetListAddTextTargetsFunction = gi.FunctionInvokerNew("Gtk", "add_text_targets")
	})
}

var addTextTargetsTargetListInvoker *gi.Function

// AddTextTargets is a representation of the C type gtk_target_list_add_text_targets.
func (recv *TargetList) AddTextTargets(info uint32) {
	targetListAddTextTargetsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(info)

	targetListAddTextTargetsFunction.Invoke(inArgs[:], nil)

}

var targetListAddUriTargetsFunction *gi.Function
var targetListAddUriTargetsFunction_Once sync.Once

func targetListAddUriTargetsFunction_Set() {
	targetListAddUriTargetsFunction_Once.Do(func() {
		targetListAddUriTargetsFunction = gi.FunctionInvokerNew("Gtk", "add_uri_targets")
	})
}

var addUriTargetsTargetListInvoker *gi.Function

// AddUriTargets is a representation of the C type gtk_target_list_add_uri_targets.
func (recv *TargetList) AddUriTargets(info uint32) {
	targetListAddUriTargetsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(info)

	targetListAddUriTargetsFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_target_list_find' : parameter 'target' of type 'Gdk.Atom' not supported

var targetListRefFunction *gi.Function
var targetListRefFunction_Once sync.Once

func targetListRefFunction_Set() {
	targetListRefFunction_Once.Do(func() {
		targetListRefFunction = gi.FunctionInvokerNew("Gtk", "ref")
	})
}

var refTargetListInvoker *gi.Function

// Ref is a representation of the C type gtk_target_list_ref.
func (recv *TargetList) Ref() *TargetList {
	targetListRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := targetListRefFunction.Invoke(inArgs[:], nil)

	retGo := &TargetList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_target_list_remove' : parameter 'target' of type 'Gdk.Atom' not supported

var targetListUnrefFunction *gi.Function
var targetListUnrefFunction_Once sync.Once

func targetListUnrefFunction_Set() {
	targetListUnrefFunction_Once.Do(func() {
		targetListUnrefFunction = gi.FunctionInvokerNew("Gtk", "unref")
	})
}

var unrefTargetListInvoker *gi.Function

// Unref is a representation of the C type gtk_target_list_unref.
func (recv *TargetList) Unref() {
	targetListUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	targetListUnrefFunction.Invoke(inArgs[:], nil)

}

var targetPairStruct *gi.Struct
var targetPairStructOnce sync.Once

func targetPairStructSet() {
	targetPairStructOnce.Do(func() {
		targetPairStruct = gi.StructNew("Gtk", "TargetPair")
	})
}

type TargetPair struct {
	native uintptr
	// UNSUPPORTED : C value 'target' : no Go type for 'Gdk.Atom'
	Flags uint32
	Info  uint32
}

var tearoffMenuItemClassStruct *gi.Struct
var tearoffMenuItemClassStructOnce sync.Once

func tearoffMenuItemClassStructSet() {
	tearoffMenuItemClassStructOnce.Do(func() {
		tearoffMenuItemClassStruct = gi.StructNew("Gtk", "TearoffMenuItemClass")
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

var tearoffMenuItemPrivateStruct *gi.Struct
var tearoffMenuItemPrivateStructOnce sync.Once

func tearoffMenuItemPrivateStructSet() {
	tearoffMenuItemPrivateStructOnce.Do(func() {
		tearoffMenuItemPrivateStruct = gi.StructNew("Gtk", "TearoffMenuItemPrivate")
	})
}

type TearoffMenuItemPrivate struct {
	native uintptr
}

var textAppearanceStruct *gi.Struct
var textAppearanceStructOnce sync.Once

func textAppearanceStructSet() {
	textAppearanceStructOnce.Do(func() {
		textAppearanceStruct = gi.StructNew("Gtk", "TextAppearance")
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

var textAttributesStruct *gi.Struct
var textAttributesStructOnce sync.Once

func textAttributesStructSet() {
	textAttributesStructOnce.Do(func() {
		textAttributesStruct = gi.StructNew("Gtk", "TextAttributes")
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

var textAttributesNewFunction *gi.Function
var textAttributesNewFunction_Once sync.Once

func textAttributesNewFunction_Set() {
	textAttributesNewFunction_Once.Do(func() {
		textAttributesNewFunction = gi.FunctionInvokerNew("Gtk", "new")
	})
}

var newTextAttributesInvoker *gi.Function

// TextAttributesNew is a representation of the C type gtk_text_attributes_new.
func TextAttributesNew() *TextAttributes {
	textAttributesNewFunction_Set()

	ret := textAttributesNewFunction.Invoke(nil, nil)

	retGo := &TextAttributes{native: ret.Pointer()}

	return retGo
}

var textAttributesCopyFunction *gi.Function
var textAttributesCopyFunction_Once sync.Once

func textAttributesCopyFunction_Set() {
	textAttributesCopyFunction_Once.Do(func() {
		textAttributesCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyTextAttributesInvoker *gi.Function

// Copy is a representation of the C type gtk_text_attributes_copy.
func (recv *TextAttributes) Copy() *TextAttributes {
	textAttributesCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textAttributesCopyFunction.Invoke(inArgs[:], nil)

	retGo := &TextAttributes{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_attributes_copy_values' : parameter 'dest' of type 'TextAttributes' not supported

var textAttributesRefFunction *gi.Function
var textAttributesRefFunction_Once sync.Once

func textAttributesRefFunction_Set() {
	textAttributesRefFunction_Once.Do(func() {
		textAttributesRefFunction = gi.FunctionInvokerNew("Gtk", "ref")
	})
}

var refTextAttributesInvoker *gi.Function

// Ref is a representation of the C type gtk_text_attributes_ref.
func (recv *TextAttributes) Ref() *TextAttributes {
	textAttributesRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textAttributesRefFunction.Invoke(inArgs[:], nil)

	retGo := &TextAttributes{native: ret.Pointer()}

	return retGo
}

var textAttributesUnrefFunction *gi.Function
var textAttributesUnrefFunction_Once sync.Once

func textAttributesUnrefFunction_Set() {
	textAttributesUnrefFunction_Once.Do(func() {
		textAttributesUnrefFunction = gi.FunctionInvokerNew("Gtk", "unref")
	})
}

var unrefTextAttributesInvoker *gi.Function

// Unref is a representation of the C type gtk_text_attributes_unref.
func (recv *TextAttributes) Unref() {
	textAttributesUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	textAttributesUnrefFunction.Invoke(inArgs[:], nil)

}

var textBTreeStruct *gi.Struct
var textBTreeStructOnce sync.Once

func textBTreeStructSet() {
	textBTreeStructOnce.Do(func() {
		textBTreeStruct = gi.StructNew("Gtk", "TextBTree")
	})
}

type TextBTree struct {
	native uintptr
}

var textBufferClassStruct *gi.Struct
var textBufferClassStructOnce sync.Once

func textBufferClassStructSet() {
	textBufferClassStructOnce.Do(func() {
		textBufferClassStruct = gi.StructNew("Gtk", "TextBufferClass")
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

var textBufferPrivateStruct *gi.Struct
var textBufferPrivateStructOnce sync.Once

func textBufferPrivateStructSet() {
	textBufferPrivateStructOnce.Do(func() {
		textBufferPrivateStruct = gi.StructNew("Gtk", "TextBufferPrivate")
	})
}

type TextBufferPrivate struct {
	native uintptr
}

var textCellAccessibleClassStruct *gi.Struct
var textCellAccessibleClassStructOnce sync.Once

func textCellAccessibleClassStructSet() {
	textCellAccessibleClassStructOnce.Do(func() {
		textCellAccessibleClassStruct = gi.StructNew("Gtk", "TextCellAccessibleClass")
	})
}

type TextCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

var textCellAccessiblePrivateStruct *gi.Struct
var textCellAccessiblePrivateStructOnce sync.Once

func textCellAccessiblePrivateStructSet() {
	textCellAccessiblePrivateStructOnce.Do(func() {
		textCellAccessiblePrivateStruct = gi.StructNew("Gtk", "TextCellAccessiblePrivate")
	})
}

type TextCellAccessiblePrivate struct {
	native uintptr
}

var textChildAnchorClassStruct *gi.Struct
var textChildAnchorClassStructOnce sync.Once

func textChildAnchorClassStructSet() {
	textChildAnchorClassStructOnce.Do(func() {
		textChildAnchorClassStruct = gi.StructNew("Gtk", "TextChildAnchorClass")
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

var textIterStruct *gi.Struct
var textIterStructOnce sync.Once

func textIterStructSet() {
	textIterStructOnce.Do(func() {
		textIterStruct = gi.StructNew("Gtk", "TextIter")
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

var textIterCopyFunction *gi.Function
var textIterCopyFunction_Once sync.Once

func textIterCopyFunction_Set() {
	textIterCopyFunction_Once.Do(func() {
		textIterCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyTextIterInvoker *gi.Function

// Copy is a representation of the C type gtk_text_iter_copy.
func (recv *TextIter) Copy() *TextIter {
	textIterCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textIterCopyFunction.Invoke(inArgs[:], nil)

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

var textIterForwardToEndFunction *gi.Function
var textIterForwardToEndFunction_Once sync.Once

func textIterForwardToEndFunction_Set() {
	textIterForwardToEndFunction_Once.Do(func() {
		textIterForwardToEndFunction = gi.FunctionInvokerNew("Gtk", "forward_to_end")
	})
}

var forwardToEndTextIterInvoker *gi.Function

// ForwardToEnd is a representation of the C type gtk_text_iter_forward_to_end.
func (recv *TextIter) ForwardToEnd() {
	textIterForwardToEndFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	textIterForwardToEndFunction.Invoke(inArgs[:], nil)

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

var textIterFreeFunction *gi.Function
var textIterFreeFunction_Once sync.Once

func textIterFreeFunction_Set() {
	textIterFreeFunction_Once.Do(func() {
		textIterFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeTextIterInvoker *gi.Function

// Free is a representation of the C type gtk_text_iter_free.
func (recv *TextIter) Free() {
	textIterFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	textIterFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_text_iter_get_attributes' : parameter 'values' of type 'TextAttributes' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_buffer' : return type 'TextBuffer' not supported

var textIterGetBytesInLineFunction *gi.Function
var textIterGetBytesInLineFunction_Once sync.Once

func textIterGetBytesInLineFunction_Set() {
	textIterGetBytesInLineFunction_Once.Do(func() {
		textIterGetBytesInLineFunction = gi.FunctionInvokerNew("Gtk", "get_bytes_in_line")
	})
}

var getBytesInLineTextIterInvoker *gi.Function

// GetBytesInLine is a representation of the C type gtk_text_iter_get_bytes_in_line.
func (recv *TextIter) GetBytesInLine() int32 {
	textIterGetBytesInLineFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textIterGetBytesInLineFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_char' : return type 'gunichar' not supported

var textIterGetCharsInLineFunction *gi.Function
var textIterGetCharsInLineFunction_Once sync.Once

func textIterGetCharsInLineFunction_Set() {
	textIterGetCharsInLineFunction_Once.Do(func() {
		textIterGetCharsInLineFunction = gi.FunctionInvokerNew("Gtk", "get_chars_in_line")
	})
}

var getCharsInLineTextIterInvoker *gi.Function

// GetCharsInLine is a representation of the C type gtk_text_iter_get_chars_in_line.
func (recv *TextIter) GetCharsInLine() int32 {
	textIterGetCharsInLineFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textIterGetCharsInLineFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_child_anchor' : return type 'TextChildAnchor' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_language' : return type 'Pango.Language' not supported

var textIterGetLineFunction *gi.Function
var textIterGetLineFunction_Once sync.Once

func textIterGetLineFunction_Set() {
	textIterGetLineFunction_Once.Do(func() {
		textIterGetLineFunction = gi.FunctionInvokerNew("Gtk", "get_line")
	})
}

var getLineTextIterInvoker *gi.Function

// GetLine is a representation of the C type gtk_text_iter_get_line.
func (recv *TextIter) GetLine() int32 {
	textIterGetLineFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textIterGetLineFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var textIterGetLineIndexFunction *gi.Function
var textIterGetLineIndexFunction_Once sync.Once

func textIterGetLineIndexFunction_Set() {
	textIterGetLineIndexFunction_Once.Do(func() {
		textIterGetLineIndexFunction = gi.FunctionInvokerNew("Gtk", "get_line_index")
	})
}

var getLineIndexTextIterInvoker *gi.Function

// GetLineIndex is a representation of the C type gtk_text_iter_get_line_index.
func (recv *TextIter) GetLineIndex() int32 {
	textIterGetLineIndexFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textIterGetLineIndexFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var textIterGetLineOffsetFunction *gi.Function
var textIterGetLineOffsetFunction_Once sync.Once

func textIterGetLineOffsetFunction_Set() {
	textIterGetLineOffsetFunction_Once.Do(func() {
		textIterGetLineOffsetFunction = gi.FunctionInvokerNew("Gtk", "get_line_offset")
	})
}

var getLineOffsetTextIterInvoker *gi.Function

// GetLineOffset is a representation of the C type gtk_text_iter_get_line_offset.
func (recv *TextIter) GetLineOffset() int32 {
	textIterGetLineOffsetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textIterGetLineOffsetFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_marks' : return type 'GLib.SList' not supported

var textIterGetOffsetFunction *gi.Function
var textIterGetOffsetFunction_Once sync.Once

func textIterGetOffsetFunction_Set() {
	textIterGetOffsetFunction_Once.Do(func() {
		textIterGetOffsetFunction = gi.FunctionInvokerNew("Gtk", "get_offset")
	})
}

var getOffsetTextIterInvoker *gi.Function

// GetOffset is a representation of the C type gtk_text_iter_get_offset.
func (recv *TextIter) GetOffset() int32 {
	textIterGetOffsetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textIterGetOffsetFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_slice' : parameter 'end' of type 'TextIter' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_tags' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_text' : parameter 'end' of type 'TextIter' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_toggled_tags' : parameter 'toggled_on' of type 'gboolean' not supported

var textIterGetVisibleLineIndexFunction *gi.Function
var textIterGetVisibleLineIndexFunction_Once sync.Once

func textIterGetVisibleLineIndexFunction_Set() {
	textIterGetVisibleLineIndexFunction_Once.Do(func() {
		textIterGetVisibleLineIndexFunction = gi.FunctionInvokerNew("Gtk", "get_visible_line_index")
	})
}

var getVisibleLineIndexTextIterInvoker *gi.Function

// GetVisibleLineIndex is a representation of the C type gtk_text_iter_get_visible_line_index.
func (recv *TextIter) GetVisibleLineIndex() int32 {
	textIterGetVisibleLineIndexFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textIterGetVisibleLineIndexFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var textIterGetVisibleLineOffsetFunction *gi.Function
var textIterGetVisibleLineOffsetFunction_Once sync.Once

func textIterGetVisibleLineOffsetFunction_Set() {
	textIterGetVisibleLineOffsetFunction_Once.Do(func() {
		textIterGetVisibleLineOffsetFunction = gi.FunctionInvokerNew("Gtk", "get_visible_line_offset")
	})
}

var getVisibleLineOffsetTextIterInvoker *gi.Function

// GetVisibleLineOffset is a representation of the C type gtk_text_iter_get_visible_line_offset.
func (recv *TextIter) GetVisibleLineOffset() int32 {
	textIterGetVisibleLineOffsetFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := textIterGetVisibleLineOffsetFunction.Invoke(inArgs[:], nil)

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

var textIterSetLineFunction *gi.Function
var textIterSetLineFunction_Once sync.Once

func textIterSetLineFunction_Set() {
	textIterSetLineFunction_Once.Do(func() {
		textIterSetLineFunction = gi.FunctionInvokerNew("Gtk", "set_line")
	})
}

var setLineTextIterInvoker *gi.Function

// SetLine is a representation of the C type gtk_text_iter_set_line.
func (recv *TextIter) SetLine(lineNumber int32) {
	textIterSetLineFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(lineNumber)

	textIterSetLineFunction.Invoke(inArgs[:], nil)

}

var textIterSetLineIndexFunction *gi.Function
var textIterSetLineIndexFunction_Once sync.Once

func textIterSetLineIndexFunction_Set() {
	textIterSetLineIndexFunction_Once.Do(func() {
		textIterSetLineIndexFunction = gi.FunctionInvokerNew("Gtk", "set_line_index")
	})
}

var setLineIndexTextIterInvoker *gi.Function

// SetLineIndex is a representation of the C type gtk_text_iter_set_line_index.
func (recv *TextIter) SetLineIndex(byteOnLine int32) {
	textIterSetLineIndexFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(byteOnLine)

	textIterSetLineIndexFunction.Invoke(inArgs[:], nil)

}

var textIterSetLineOffsetFunction *gi.Function
var textIterSetLineOffsetFunction_Once sync.Once

func textIterSetLineOffsetFunction_Set() {
	textIterSetLineOffsetFunction_Once.Do(func() {
		textIterSetLineOffsetFunction = gi.FunctionInvokerNew("Gtk", "set_line_offset")
	})
}

var setLineOffsetTextIterInvoker *gi.Function

// SetLineOffset is a representation of the C type gtk_text_iter_set_line_offset.
func (recv *TextIter) SetLineOffset(charOnLine int32) {
	textIterSetLineOffsetFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(charOnLine)

	textIterSetLineOffsetFunction.Invoke(inArgs[:], nil)

}

var textIterSetOffsetFunction *gi.Function
var textIterSetOffsetFunction_Once sync.Once

func textIterSetOffsetFunction_Set() {
	textIterSetOffsetFunction_Once.Do(func() {
		textIterSetOffsetFunction = gi.FunctionInvokerNew("Gtk", "set_offset")
	})
}

var setOffsetTextIterInvoker *gi.Function

// SetOffset is a representation of the C type gtk_text_iter_set_offset.
func (recv *TextIter) SetOffset(charOffset int32) {
	textIterSetOffsetFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(charOffset)

	textIterSetOffsetFunction.Invoke(inArgs[:], nil)

}

var textIterSetVisibleLineIndexFunction *gi.Function
var textIterSetVisibleLineIndexFunction_Once sync.Once

func textIterSetVisibleLineIndexFunction_Set() {
	textIterSetVisibleLineIndexFunction_Once.Do(func() {
		textIterSetVisibleLineIndexFunction = gi.FunctionInvokerNew("Gtk", "set_visible_line_index")
	})
}

var setVisibleLineIndexTextIterInvoker *gi.Function

// SetVisibleLineIndex is a representation of the C type gtk_text_iter_set_visible_line_index.
func (recv *TextIter) SetVisibleLineIndex(byteOnLine int32) {
	textIterSetVisibleLineIndexFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(byteOnLine)

	textIterSetVisibleLineIndexFunction.Invoke(inArgs[:], nil)

}

var textIterSetVisibleLineOffsetFunction *gi.Function
var textIterSetVisibleLineOffsetFunction_Once sync.Once

func textIterSetVisibleLineOffsetFunction_Set() {
	textIterSetVisibleLineOffsetFunction_Once.Do(func() {
		textIterSetVisibleLineOffsetFunction = gi.FunctionInvokerNew("Gtk", "set_visible_line_offset")
	})
}

var setVisibleLineOffsetTextIterInvoker *gi.Function

// SetVisibleLineOffset is a representation of the C type gtk_text_iter_set_visible_line_offset.
func (recv *TextIter) SetVisibleLineOffset(charOnLine int32) {
	textIterSetVisibleLineOffsetFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(charOnLine)

	textIterSetVisibleLineOffsetFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_text_iter_starts_line' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_starts_sentence' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_starts_tag' : parameter 'tag' of type 'TextTag' not supported

// UNSUPPORTED : C value 'gtk_text_iter_starts_word' : return type 'gboolean' not supported

// UNSUPPORTED : C value 'gtk_text_iter_toggles_tag' : parameter 'tag' of type 'TextTag' not supported

var textMarkClassStruct *gi.Struct
var textMarkClassStructOnce sync.Once

func textMarkClassStructSet() {
	textMarkClassStructOnce.Do(func() {
		textMarkClassStruct = gi.StructNew("Gtk", "TextMarkClass")
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

var textTagClassStruct *gi.Struct
var textTagClassStructOnce sync.Once

func textTagClassStructSet() {
	textTagClassStructOnce.Do(func() {
		textTagClassStruct = gi.StructNew("Gtk", "TextTagClass")
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

var textTagPrivateStruct *gi.Struct
var textTagPrivateStructOnce sync.Once

func textTagPrivateStructSet() {
	textTagPrivateStructOnce.Do(func() {
		textTagPrivateStruct = gi.StructNew("Gtk", "TextTagPrivate")
	})
}

type TextTagPrivate struct {
	native uintptr
}

var textTagTableClassStruct *gi.Struct
var textTagTableClassStructOnce sync.Once

func textTagTableClassStructSet() {
	textTagTableClassStructOnce.Do(func() {
		textTagTableClassStruct = gi.StructNew("Gtk", "TextTagTableClass")
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

var textTagTablePrivateStruct *gi.Struct
var textTagTablePrivateStructOnce sync.Once

func textTagTablePrivateStructSet() {
	textTagTablePrivateStructOnce.Do(func() {
		textTagTablePrivateStruct = gi.StructNew("Gtk", "TextTagTablePrivate")
	})
}

type TextTagTablePrivate struct {
	native uintptr
}

var textViewAccessibleClassStruct *gi.Struct
var textViewAccessibleClassStructOnce sync.Once

func textViewAccessibleClassStructSet() {
	textViewAccessibleClassStructOnce.Do(func() {
		textViewAccessibleClassStruct = gi.StructNew("Gtk", "TextViewAccessibleClass")
	})
}

type TextViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var textViewAccessiblePrivateStruct *gi.Struct
var textViewAccessiblePrivateStructOnce sync.Once

func textViewAccessiblePrivateStructSet() {
	textViewAccessiblePrivateStructOnce.Do(func() {
		textViewAccessiblePrivateStruct = gi.StructNew("Gtk", "TextViewAccessiblePrivate")
	})
}

type TextViewAccessiblePrivate struct {
	native uintptr
}

var textViewClassStruct *gi.Struct
var textViewClassStructOnce sync.Once

func textViewClassStructSet() {
	textViewClassStructOnce.Do(func() {
		textViewClassStruct = gi.StructNew("Gtk", "TextViewClass")
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

var textViewPrivateStruct *gi.Struct
var textViewPrivateStructOnce sync.Once

func textViewPrivateStructSet() {
	textViewPrivateStructOnce.Do(func() {
		textViewPrivateStruct = gi.StructNew("Gtk", "TextViewPrivate")
	})
}

type TextViewPrivate struct {
	native uintptr
}

var themeEngineStruct *gi.Struct
var themeEngineStructOnce sync.Once

func themeEngineStructSet() {
	themeEngineStructOnce.Do(func() {
		themeEngineStruct = gi.StructNew("Gtk", "ThemeEngine")
	})
}

type ThemeEngine struct {
	native uintptr
}

var themingEngineClassStruct *gi.Struct
var themingEngineClassStructOnce sync.Once

func themingEngineClassStructSet() {
	themingEngineClassStructOnce.Do(func() {
		themingEngineClassStruct = gi.StructNew("Gtk", "ThemingEngineClass")
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

var themingEnginePrivateStruct *gi.Struct
var themingEnginePrivateStructOnce sync.Once

func themingEnginePrivateStructSet() {
	themingEnginePrivateStructOnce.Do(func() {
		themingEnginePrivateStruct = gi.StructNew("Gtk", "ThemingEnginePrivate")
	})
}

type ThemingEnginePrivate struct {
	native uintptr
}

var toggleActionClassStruct *gi.Struct
var toggleActionClassStructOnce sync.Once

func toggleActionClassStructSet() {
	toggleActionClassStructOnce.Do(func() {
		toggleActionClassStruct = gi.StructNew("Gtk", "ToggleActionClass")
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

var toggleActionEntryStruct *gi.Struct
var toggleActionEntryStructOnce sync.Once

func toggleActionEntryStructSet() {
	toggleActionEntryStructOnce.Do(func() {
		toggleActionEntryStruct = gi.StructNew("Gtk", "ToggleActionEntry")
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

var toggleActionPrivateStruct *gi.Struct
var toggleActionPrivateStructOnce sync.Once

func toggleActionPrivateStructSet() {
	toggleActionPrivateStructOnce.Do(func() {
		toggleActionPrivateStruct = gi.StructNew("Gtk", "ToggleActionPrivate")
	})
}

type ToggleActionPrivate struct {
	native uintptr
}

var toggleButtonAccessibleClassStruct *gi.Struct
var toggleButtonAccessibleClassStructOnce sync.Once

func toggleButtonAccessibleClassStructSet() {
	toggleButtonAccessibleClassStructOnce.Do(func() {
		toggleButtonAccessibleClassStruct = gi.StructNew("Gtk", "ToggleButtonAccessibleClass")
	})
}

type ToggleButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var toggleButtonAccessiblePrivateStruct *gi.Struct
var toggleButtonAccessiblePrivateStructOnce sync.Once

func toggleButtonAccessiblePrivateStructSet() {
	toggleButtonAccessiblePrivateStructOnce.Do(func() {
		toggleButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "ToggleButtonAccessiblePrivate")
	})
}

type ToggleButtonAccessiblePrivate struct {
	native uintptr
}

var toggleButtonClassStruct *gi.Struct
var toggleButtonClassStructOnce sync.Once

func toggleButtonClassStructSet() {
	toggleButtonClassStructOnce.Do(func() {
		toggleButtonClassStruct = gi.StructNew("Gtk", "ToggleButtonClass")
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

var toggleButtonPrivateStruct *gi.Struct
var toggleButtonPrivateStructOnce sync.Once

func toggleButtonPrivateStructSet() {
	toggleButtonPrivateStructOnce.Do(func() {
		toggleButtonPrivateStruct = gi.StructNew("Gtk", "ToggleButtonPrivate")
	})
}

type ToggleButtonPrivate struct {
	native uintptr
}

var toggleToolButtonClassStruct *gi.Struct
var toggleToolButtonClassStructOnce sync.Once

func toggleToolButtonClassStructSet() {
	toggleToolButtonClassStructOnce.Do(func() {
		toggleToolButtonClassStruct = gi.StructNew("Gtk", "ToggleToolButtonClass")
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

var toggleToolButtonPrivateStruct *gi.Struct
var toggleToolButtonPrivateStructOnce sync.Once

func toggleToolButtonPrivateStructSet() {
	toggleToolButtonPrivateStructOnce.Do(func() {
		toggleToolButtonPrivateStruct = gi.StructNew("Gtk", "ToggleToolButtonPrivate")
	})
}

type ToggleToolButtonPrivate struct {
	native uintptr
}

var toolButtonClassStruct *gi.Struct
var toolButtonClassStructOnce sync.Once

func toolButtonClassStructSet() {
	toolButtonClassStructOnce.Do(func() {
		toolButtonClassStruct = gi.StructNew("Gtk", "ToolButtonClass")
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

var toolButtonPrivateStruct *gi.Struct
var toolButtonPrivateStructOnce sync.Once

func toolButtonPrivateStructSet() {
	toolButtonPrivateStructOnce.Do(func() {
		toolButtonPrivateStruct = gi.StructNew("Gtk", "ToolButtonPrivate")
	})
}

type ToolButtonPrivate struct {
	native uintptr
}

var toolItemClassStruct *gi.Struct
var toolItemClassStructOnce sync.Once

func toolItemClassStructSet() {
	toolItemClassStructOnce.Do(func() {
		toolItemClassStruct = gi.StructNew("Gtk", "ToolItemClass")
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

var toolItemGroupClassStruct *gi.Struct
var toolItemGroupClassStructOnce sync.Once

func toolItemGroupClassStructSet() {
	toolItemGroupClassStructOnce.Do(func() {
		toolItemGroupClassStruct = gi.StructNew("Gtk", "ToolItemGroupClass")
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

var toolItemGroupPrivateStruct *gi.Struct
var toolItemGroupPrivateStructOnce sync.Once

func toolItemGroupPrivateStructSet() {
	toolItemGroupPrivateStructOnce.Do(func() {
		toolItemGroupPrivateStruct = gi.StructNew("Gtk", "ToolItemGroupPrivate")
	})
}

type ToolItemGroupPrivate struct {
	native uintptr
}

var toolItemPrivateStruct *gi.Struct
var toolItemPrivateStructOnce sync.Once

func toolItemPrivateStructSet() {
	toolItemPrivateStructOnce.Do(func() {
		toolItemPrivateStruct = gi.StructNew("Gtk", "ToolItemPrivate")
	})
}

type ToolItemPrivate struct {
	native uintptr
}

var toolPaletteClassStruct *gi.Struct
var toolPaletteClassStructOnce sync.Once

func toolPaletteClassStructSet() {
	toolPaletteClassStructOnce.Do(func() {
		toolPaletteClassStruct = gi.StructNew("Gtk", "ToolPaletteClass")
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

var toolPalettePrivateStruct *gi.Struct
var toolPalettePrivateStructOnce sync.Once

func toolPalettePrivateStructSet() {
	toolPalettePrivateStructOnce.Do(func() {
		toolPalettePrivateStruct = gi.StructNew("Gtk", "ToolPalettePrivate")
	})
}

type ToolPalettePrivate struct {
	native uintptr
}

var toolShellIfaceStruct *gi.Struct
var toolShellIfaceStructOnce sync.Once

func toolShellIfaceStructSet() {
	toolShellIfaceStructOnce.Do(func() {
		toolShellIfaceStruct = gi.StructNew("Gtk", "ToolShellIface")
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

var toolbarClassStruct *gi.Struct
var toolbarClassStructOnce sync.Once

func toolbarClassStructSet() {
	toolbarClassStructOnce.Do(func() {
		toolbarClassStruct = gi.StructNew("Gtk", "ToolbarClass")
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

var toolbarPrivateStruct *gi.Struct
var toolbarPrivateStructOnce sync.Once

func toolbarPrivateStructSet() {
	toolbarPrivateStructOnce.Do(func() {
		toolbarPrivateStruct = gi.StructNew("Gtk", "ToolbarPrivate")
	})
}

type ToolbarPrivate struct {
	native uintptr
}

var toplevelAccessibleClassStruct *gi.Struct
var toplevelAccessibleClassStructOnce sync.Once

func toplevelAccessibleClassStructSet() {
	toplevelAccessibleClassStructOnce.Do(func() {
		toplevelAccessibleClassStruct = gi.StructNew("Gtk", "ToplevelAccessibleClass")
	})
}

type ToplevelAccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Atk.ObjectClass'
}

var toplevelAccessiblePrivateStruct *gi.Struct
var toplevelAccessiblePrivateStructOnce sync.Once

func toplevelAccessiblePrivateStructSet() {
	toplevelAccessiblePrivateStructOnce.Do(func() {
		toplevelAccessiblePrivateStruct = gi.StructNew("Gtk", "ToplevelAccessiblePrivate")
	})
}

type ToplevelAccessiblePrivate struct {
	native uintptr
}

var treeDragDestIfaceStruct *gi.Struct
var treeDragDestIfaceStructOnce sync.Once

func treeDragDestIfaceStructSet() {
	treeDragDestIfaceStructOnce.Do(func() {
		treeDragDestIfaceStruct = gi.StructNew("Gtk", "TreeDragDestIface")
	})
}

type TreeDragDestIface struct {
	native uintptr
	// UNSUPPORTED : C value 'drag_data_received' : missing Type
	// UNSUPPORTED : C value 'row_drop_possible' : missing Type
}

var treeDragSourceIfaceStruct *gi.Struct
var treeDragSourceIfaceStructOnce sync.Once

func treeDragSourceIfaceStructSet() {
	treeDragSourceIfaceStructOnce.Do(func() {
		treeDragSourceIfaceStruct = gi.StructNew("Gtk", "TreeDragSourceIface")
	})
}

type TreeDragSourceIface struct {
	native uintptr
	// UNSUPPORTED : C value 'row_draggable' : missing Type
	// UNSUPPORTED : C value 'drag_data_get' : missing Type
	// UNSUPPORTED : C value 'drag_data_delete' : missing Type
}

var treeIterStruct *gi.Struct
var treeIterStructOnce sync.Once

func treeIterStructSet() {
	treeIterStructOnce.Do(func() {
		treeIterStruct = gi.StructNew("Gtk", "TreeIter")
	})
}

type TreeIter struct {
	native uintptr
	Stamp  int32
	// UNSUPPORTED : C value 'user_data' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'user_data2' : no Go type for 'gpointer'
	// UNSUPPORTED : C value 'user_data3' : no Go type for 'gpointer'
}

var treeIterCopyFunction *gi.Function
var treeIterCopyFunction_Once sync.Once

func treeIterCopyFunction_Set() {
	treeIterCopyFunction_Once.Do(func() {
		treeIterCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyTreeIterInvoker *gi.Function

// Copy is a representation of the C type gtk_tree_iter_copy.
func (recv *TreeIter) Copy() *TreeIter {
	treeIterCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treeIterCopyFunction.Invoke(inArgs[:], nil)

	retGo := &TreeIter{native: ret.Pointer()}

	return retGo
}

var treeIterFreeFunction *gi.Function
var treeIterFreeFunction_Once sync.Once

func treeIterFreeFunction_Set() {
	treeIterFreeFunction_Once.Do(func() {
		treeIterFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeTreeIterInvoker *gi.Function

// Free is a representation of the C type gtk_tree_iter_free.
func (recv *TreeIter) Free() {
	treeIterFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	treeIterFreeFunction.Invoke(inArgs[:], nil)

}

var treeModelFilterClassStruct *gi.Struct
var treeModelFilterClassStructOnce sync.Once

func treeModelFilterClassStructSet() {
	treeModelFilterClassStructOnce.Do(func() {
		treeModelFilterClassStruct = gi.StructNew("Gtk", "TreeModelFilterClass")
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

var treeModelFilterPrivateStruct *gi.Struct
var treeModelFilterPrivateStructOnce sync.Once

func treeModelFilterPrivateStructSet() {
	treeModelFilterPrivateStructOnce.Do(func() {
		treeModelFilterPrivateStruct = gi.StructNew("Gtk", "TreeModelFilterPrivate")
	})
}

type TreeModelFilterPrivate struct {
	native uintptr
}

var treeModelIfaceStruct *gi.Struct
var treeModelIfaceStructOnce sync.Once

func treeModelIfaceStructSet() {
	treeModelIfaceStructOnce.Do(func() {
		treeModelIfaceStruct = gi.StructNew("Gtk", "TreeModelIface")
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

var treeModelSortClassStruct *gi.Struct
var treeModelSortClassStructOnce sync.Once

func treeModelSortClassStructSet() {
	treeModelSortClassStructOnce.Do(func() {
		treeModelSortClassStruct = gi.StructNew("Gtk", "TreeModelSortClass")
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

var treeModelSortPrivateStruct *gi.Struct
var treeModelSortPrivateStructOnce sync.Once

func treeModelSortPrivateStructSet() {
	treeModelSortPrivateStructOnce.Do(func() {
		treeModelSortPrivateStruct = gi.StructNew("Gtk", "TreeModelSortPrivate")
	})
}

type TreeModelSortPrivate struct {
	native uintptr
}

var treePathStruct *gi.Struct
var treePathStructOnce sync.Once

func treePathStructSet() {
	treePathStructOnce.Do(func() {
		treePathStruct = gi.StructNew("Gtk", "TreePath")
	})
}

type TreePath struct {
	native uintptr
}

var treePathNewFunction *gi.Function
var treePathNewFunction_Once sync.Once

func treePathNewFunction_Set() {
	treePathNewFunction_Once.Do(func() {
		treePathNewFunction = gi.FunctionInvokerNew("Gtk", "new")
	})
}

var newTreePathInvoker *gi.Function

// TreePathNew is a representation of the C type gtk_tree_path_new.
func TreePathNew() *TreePath {
	treePathNewFunction_Set()

	ret := treePathNewFunction.Invoke(nil, nil)

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var treePathNewFirstFunction *gi.Function
var treePathNewFirstFunction_Once sync.Once

func treePathNewFirstFunction_Set() {
	treePathNewFirstFunction_Once.Do(func() {
		treePathNewFirstFunction = gi.FunctionInvokerNew("Gtk", "new_first")
	})
}

var newFirstTreePathInvoker *gi.Function

// TreePathNewFirst is a representation of the C type gtk_tree_path_new_first.
func TreePathNewFirst() *TreePath {
	treePathNewFirstFunction_Set()

	ret := treePathNewFirstFunction.Invoke(nil, nil)

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_tree_path_new_from_indices' : parameter '...' has no type

// UNSUPPORTED : C value 'gtk_tree_path_new_from_indicesv' : parameter 'indices' has no type

var treePathNewFromStringFunction *gi.Function
var treePathNewFromStringFunction_Once sync.Once

func treePathNewFromStringFunction_Set() {
	treePathNewFromStringFunction_Once.Do(func() {
		treePathNewFromStringFunction = gi.FunctionInvokerNew("Gtk", "new_from_string")
	})
}

var newFromStringTreePathInvoker *gi.Function

// TreePathNewFromString is a representation of the C type gtk_tree_path_new_from_string.
func TreePathNewFromString(path string) *TreePath {
	treePathNewFromStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetString(path)

	ret := treePathNewFromStringFunction.Invoke(inArgs[:], nil)

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var treePathAppendIndexFunction *gi.Function
var treePathAppendIndexFunction_Once sync.Once

func treePathAppendIndexFunction_Set() {
	treePathAppendIndexFunction_Once.Do(func() {
		treePathAppendIndexFunction = gi.FunctionInvokerNew("Gtk", "append_index")
	})
}

var appendIndexTreePathInvoker *gi.Function

// AppendIndex is a representation of the C type gtk_tree_path_append_index.
func (recv *TreePath) AppendIndex(index int32) {
	treePathAppendIndexFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)

	treePathAppendIndexFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_tree_path_compare' : parameter 'b' of type 'TreePath' not supported

var treePathCopyFunction *gi.Function
var treePathCopyFunction_Once sync.Once

func treePathCopyFunction_Set() {
	treePathCopyFunction_Once.Do(func() {
		treePathCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyTreePathInvoker *gi.Function

// Copy is a representation of the C type gtk_tree_path_copy.
func (recv *TreePath) Copy() *TreePath {
	treePathCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treePathCopyFunction.Invoke(inArgs[:], nil)

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var treePathDownFunction *gi.Function
var treePathDownFunction_Once sync.Once

func treePathDownFunction_Set() {
	treePathDownFunction_Once.Do(func() {
		treePathDownFunction = gi.FunctionInvokerNew("Gtk", "down")
	})
}

var downTreePathInvoker *gi.Function

// Down is a representation of the C type gtk_tree_path_down.
func (recv *TreePath) Down() {
	treePathDownFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	treePathDownFunction.Invoke(inArgs[:], nil)

}

var treePathFreeFunction *gi.Function
var treePathFreeFunction_Once sync.Once

func treePathFreeFunction_Set() {
	treePathFreeFunction_Once.Do(func() {
		treePathFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeTreePathInvoker *gi.Function

// Free is a representation of the C type gtk_tree_path_free.
func (recv *TreePath) Free() {
	treePathFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	treePathFreeFunction.Invoke(inArgs[:], nil)

}

var treePathGetDepthFunction *gi.Function
var treePathGetDepthFunction_Once sync.Once

func treePathGetDepthFunction_Set() {
	treePathGetDepthFunction_Once.Do(func() {
		treePathGetDepthFunction = gi.FunctionInvokerNew("Gtk", "get_depth")
	})
}

var getDepthTreePathInvoker *gi.Function

// GetDepth is a representation of the C type gtk_tree_path_get_depth.
func (recv *TreePath) GetDepth() int32 {
	treePathGetDepthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treePathGetDepthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var treePathGetIndicesFunction *gi.Function
var treePathGetIndicesFunction_Once sync.Once

func treePathGetIndicesFunction_Set() {
	treePathGetIndicesFunction_Once.Do(func() {
		treePathGetIndicesFunction = gi.FunctionInvokerNew("Gtk", "get_indices")
	})
}

var getIndicesTreePathInvoker *gi.Function

// GetIndices is a representation of the C type gtk_tree_path_get_indices.
func (recv *TreePath) GetIndices() int32 {
	treePathGetIndicesFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treePathGetIndicesFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

var treePathGetIndicesWithDepthFunction *gi.Function
var treePathGetIndicesWithDepthFunction_Once sync.Once

func treePathGetIndicesWithDepthFunction_Set() {
	treePathGetIndicesWithDepthFunction_Once.Do(func() {
		treePathGetIndicesWithDepthFunction = gi.FunctionInvokerNew("Gtk", "get_indices_with_depth")
	})
}

var getIndicesWithDepthTreePathInvoker *gi.Function

// GetIndicesWithDepth is a representation of the C type gtk_tree_path_get_indices_with_depth.
func (recv *TreePath) GetIndicesWithDepth() int32 {
	treePathGetIndicesWithDepthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	treePathGetIndicesWithDepthFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'gtk_tree_path_is_ancestor' : parameter 'descendant' of type 'TreePath' not supported

// UNSUPPORTED : C value 'gtk_tree_path_is_descendant' : parameter 'ancestor' of type 'TreePath' not supported

var treePathNextFunction *gi.Function
var treePathNextFunction_Once sync.Once

func treePathNextFunction_Set() {
	treePathNextFunction_Once.Do(func() {
		treePathNextFunction = gi.FunctionInvokerNew("Gtk", "next")
	})
}

var nextTreePathInvoker *gi.Function

// Next is a representation of the C type gtk_tree_path_next.
func (recv *TreePath) Next() {
	treePathNextFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	treePathNextFunction.Invoke(inArgs[:], nil)

}

var treePathPrependIndexFunction *gi.Function
var treePathPrependIndexFunction_Once sync.Once

func treePathPrependIndexFunction_Set() {
	treePathPrependIndexFunction_Once.Do(func() {
		treePathPrependIndexFunction = gi.FunctionInvokerNew("Gtk", "prepend_index")
	})
}

var prependIndexTreePathInvoker *gi.Function

// PrependIndex is a representation of the C type gtk_tree_path_prepend_index.
func (recv *TreePath) PrependIndex(index int32) {
	treePathPrependIndexFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)

	treePathPrependIndexFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_tree_path_prev' : return type 'gboolean' not supported

var treePathToStringFunction *gi.Function
var treePathToStringFunction_Once sync.Once

func treePathToStringFunction_Set() {
	treePathToStringFunction_Once.Do(func() {
		treePathToStringFunction = gi.FunctionInvokerNew("Gtk", "to_string")
	})
}

var toStringTreePathInvoker *gi.Function

// ToString is a representation of the C type gtk_tree_path_to_string.
func (recv *TreePath) ToString() string {
	treePathToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treePathToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_tree_path_up' : return type 'gboolean' not supported

var treeRowReferenceStruct *gi.Struct
var treeRowReferenceStructOnce sync.Once

func treeRowReferenceStructSet() {
	treeRowReferenceStructOnce.Do(func() {
		treeRowReferenceStruct = gi.StructNew("Gtk", "TreeRowReference")
	})
}

type TreeRowReference struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_tree_row_reference_new' : parameter 'model' of type 'TreeModel' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_new_proxy' : parameter 'proxy' of type 'GObject.Object' not supported

var treeRowReferenceCopyFunction *gi.Function
var treeRowReferenceCopyFunction_Once sync.Once

func treeRowReferenceCopyFunction_Set() {
	treeRowReferenceCopyFunction_Once.Do(func() {
		treeRowReferenceCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyTreeRowReferenceInvoker *gi.Function

// Copy is a representation of the C type gtk_tree_row_reference_copy.
func (recv *TreeRowReference) Copy() *TreeRowReference {
	treeRowReferenceCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treeRowReferenceCopyFunction.Invoke(inArgs[:], nil)

	retGo := &TreeRowReference{native: ret.Pointer()}

	return retGo
}

var treeRowReferenceFreeFunction *gi.Function
var treeRowReferenceFreeFunction_Once sync.Once

func treeRowReferenceFreeFunction_Set() {
	treeRowReferenceFreeFunction_Once.Do(func() {
		treeRowReferenceFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeTreeRowReferenceInvoker *gi.Function

// Free is a representation of the C type gtk_tree_row_reference_free.
func (recv *TreeRowReference) Free() {
	treeRowReferenceFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	treeRowReferenceFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_tree_row_reference_get_model' : return type 'TreeModel' not supported

var treeRowReferenceGetPathFunction *gi.Function
var treeRowReferenceGetPathFunction_Once sync.Once

func treeRowReferenceGetPathFunction_Set() {
	treeRowReferenceGetPathFunction_Once.Do(func() {
		treeRowReferenceGetPathFunction = gi.FunctionInvokerNew("Gtk", "get_path")
	})
}

var getPathTreeRowReferenceInvoker *gi.Function

// GetPath is a representation of the C type gtk_tree_row_reference_get_path.
func (recv *TreeRowReference) GetPath() *TreePath {
	treeRowReferenceGetPathFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := treeRowReferenceGetPathFunction.Invoke(inArgs[:], nil)

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_tree_row_reference_valid' : return type 'gboolean' not supported

var treeSelectionClassStruct *gi.Struct
var treeSelectionClassStructOnce sync.Once

func treeSelectionClassStructSet() {
	treeSelectionClassStructOnce.Do(func() {
		treeSelectionClassStruct = gi.StructNew("Gtk", "TreeSelectionClass")
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

var treeSelectionPrivateStruct *gi.Struct
var treeSelectionPrivateStructOnce sync.Once

func treeSelectionPrivateStructSet() {
	treeSelectionPrivateStructOnce.Do(func() {
		treeSelectionPrivateStruct = gi.StructNew("Gtk", "TreeSelectionPrivate")
	})
}

type TreeSelectionPrivate struct {
	native uintptr
}

var treeSortableIfaceStruct *gi.Struct
var treeSortableIfaceStructOnce sync.Once

func treeSortableIfaceStructSet() {
	treeSortableIfaceStructOnce.Do(func() {
		treeSortableIfaceStruct = gi.StructNew("Gtk", "TreeSortableIface")
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

var treeStoreClassStruct *gi.Struct
var treeStoreClassStructOnce sync.Once

func treeStoreClassStructSet() {
	treeStoreClassStructOnce.Do(func() {
		treeStoreClassStruct = gi.StructNew("Gtk", "TreeStoreClass")
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

var treeStorePrivateStruct *gi.Struct
var treeStorePrivateStructOnce sync.Once

func treeStorePrivateStructSet() {
	treeStorePrivateStructOnce.Do(func() {
		treeStorePrivateStruct = gi.StructNew("Gtk", "TreeStorePrivate")
	})
}

type TreeStorePrivate struct {
	native uintptr
}

var treeViewAccessibleClassStruct *gi.Struct
var treeViewAccessibleClassStructOnce sync.Once

func treeViewAccessibleClassStructSet() {
	treeViewAccessibleClassStructOnce.Do(func() {
		treeViewAccessibleClassStruct = gi.StructNew("Gtk", "TreeViewAccessibleClass")
	})
}

type TreeViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var treeViewAccessiblePrivateStruct *gi.Struct
var treeViewAccessiblePrivateStructOnce sync.Once

func treeViewAccessiblePrivateStructSet() {
	treeViewAccessiblePrivateStructOnce.Do(func() {
		treeViewAccessiblePrivateStruct = gi.StructNew("Gtk", "TreeViewAccessiblePrivate")
	})
}

type TreeViewAccessiblePrivate struct {
	native uintptr
}

var treeViewClassStruct *gi.Struct
var treeViewClassStructOnce sync.Once

func treeViewClassStructSet() {
	treeViewClassStructOnce.Do(func() {
		treeViewClassStruct = gi.StructNew("Gtk", "TreeViewClass")
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

var treeViewColumnClassStruct *gi.Struct
var treeViewColumnClassStructOnce sync.Once

func treeViewColumnClassStructSet() {
	treeViewColumnClassStructOnce.Do(func() {
		treeViewColumnClassStruct = gi.StructNew("Gtk", "TreeViewColumnClass")
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

var treeViewColumnPrivateStruct *gi.Struct
var treeViewColumnPrivateStructOnce sync.Once

func treeViewColumnPrivateStructSet() {
	treeViewColumnPrivateStructOnce.Do(func() {
		treeViewColumnPrivateStruct = gi.StructNew("Gtk", "TreeViewColumnPrivate")
	})
}

type TreeViewColumnPrivate struct {
	native uintptr
}

var treeViewPrivateStruct *gi.Struct
var treeViewPrivateStructOnce sync.Once

func treeViewPrivateStructSet() {
	treeViewPrivateStructOnce.Do(func() {
		treeViewPrivateStruct = gi.StructNew("Gtk", "TreeViewPrivate")
	})
}

type TreeViewPrivate struct {
	native uintptr
}

var uIManagerClassStruct *gi.Struct
var uIManagerClassStructOnce sync.Once

func uIManagerClassStructSet() {
	uIManagerClassStructOnce.Do(func() {
		uIManagerClassStruct = gi.StructNew("Gtk", "UIManagerClass")
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

var uIManagerPrivateStruct *gi.Struct
var uIManagerPrivateStructOnce sync.Once

func uIManagerPrivateStructSet() {
	uIManagerPrivateStructOnce.Do(func() {
		uIManagerPrivateStruct = gi.StructNew("Gtk", "UIManagerPrivate")
	})
}

type UIManagerPrivate struct {
	native uintptr
}

var vBoxClassStruct *gi.Struct
var vBoxClassStructOnce sync.Once

func vBoxClassStructSet() {
	vBoxClassStructOnce.Do(func() {
		vBoxClassStruct = gi.StructNew("Gtk", "VBoxClass")
	})
}

type VBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
}

var vButtonBoxClassStruct *gi.Struct
var vButtonBoxClassStructOnce sync.Once

func vButtonBoxClassStructSet() {
	vButtonBoxClassStructOnce.Do(func() {
		vButtonBoxClassStruct = gi.StructNew("Gtk", "VButtonBoxClass")
	})
}

type VButtonBoxClass struct {
	native      uintptr
	ParentClass *ButtonBoxClass
}

var vPanedClassStruct *gi.Struct
var vPanedClassStructOnce sync.Once

func vPanedClassStructSet() {
	vPanedClassStructOnce.Do(func() {
		vPanedClassStruct = gi.StructNew("Gtk", "VPanedClass")
	})
}

type VPanedClass struct {
	native      uintptr
	ParentClass *PanedClass
}

var vScaleClassStruct *gi.Struct
var vScaleClassStructOnce sync.Once

func vScaleClassStructSet() {
	vScaleClassStructOnce.Do(func() {
		vScaleClassStruct = gi.StructNew("Gtk", "VScaleClass")
	})
}

type VScaleClass struct {
	native      uintptr
	ParentClass *ScaleClass
}

var vScrollbarClassStruct *gi.Struct
var vScrollbarClassStructOnce sync.Once

func vScrollbarClassStructSet() {
	vScrollbarClassStructOnce.Do(func() {
		vScrollbarClassStruct = gi.StructNew("Gtk", "VScrollbarClass")
	})
}

type VScrollbarClass struct {
	native      uintptr
	ParentClass *ScrollbarClass
}

var vSeparatorClassStruct *gi.Struct
var vSeparatorClassStructOnce sync.Once

func vSeparatorClassStructSet() {
	vSeparatorClassStructOnce.Do(func() {
		vSeparatorClassStruct = gi.StructNew("Gtk", "VSeparatorClass")
	})
}

type VSeparatorClass struct {
	native      uintptr
	ParentClass *SeparatorClass
}

var viewportClassStruct *gi.Struct
var viewportClassStructOnce sync.Once

func viewportClassStructSet() {
	viewportClassStructOnce.Do(func() {
		viewportClassStruct = gi.StructNew("Gtk", "ViewportClass")
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

var viewportPrivateStruct *gi.Struct
var viewportPrivateStructOnce sync.Once

func viewportPrivateStructSet() {
	viewportPrivateStructOnce.Do(func() {
		viewportPrivateStruct = gi.StructNew("Gtk", "ViewportPrivate")
	})
}

type ViewportPrivate struct {
	native uintptr
}

var volumeButtonClassStruct *gi.Struct
var volumeButtonClassStructOnce sync.Once

func volumeButtonClassStructSet() {
	volumeButtonClassStructOnce.Do(func() {
		volumeButtonClassStruct = gi.StructNew("Gtk", "VolumeButtonClass")
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

var widgetAccessibleClassStruct *gi.Struct
var widgetAccessibleClassStructOnce sync.Once

func widgetAccessibleClassStructSet() {
	widgetAccessibleClassStructOnce.Do(func() {
		widgetAccessibleClassStruct = gi.StructNew("Gtk", "WidgetAccessibleClass")
	})
}

type WidgetAccessibleClass struct {
	native      uintptr
	ParentClass *AccessibleClass
	// UNSUPPORTED : C value 'notify_gtk' : missing Type
}

var widgetAccessiblePrivateStruct *gi.Struct
var widgetAccessiblePrivateStructOnce sync.Once

func widgetAccessiblePrivateStructSet() {
	widgetAccessiblePrivateStructOnce.Do(func() {
		widgetAccessiblePrivateStruct = gi.StructNew("Gtk", "WidgetAccessiblePrivate")
	})
}

type WidgetAccessiblePrivate struct {
	native uintptr
}

var widgetClassStruct *gi.Struct
var widgetClassStructOnce sync.Once

func widgetClassStructSet() {
	widgetClassStructOnce.Do(func() {
		widgetClassStruct = gi.StructNew("Gtk", "WidgetClass")
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

var widgetClassGetCssNameFunction *gi.Function
var widgetClassGetCssNameFunction_Once sync.Once

func widgetClassGetCssNameFunction_Set() {
	widgetClassGetCssNameFunction_Once.Do(func() {
		widgetClassGetCssNameFunction = gi.FunctionInvokerNew("Gtk", "get_css_name")
	})
}

var getCssNameWidgetClassInvoker *gi.Function

// GetCssName is a representation of the C type gtk_widget_class_get_css_name.
func (recv *WidgetClass) GetCssName() string {
	widgetClassGetCssNameFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := widgetClassGetCssNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_class_install_style_property' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_widget_class_install_style_property_parser' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

var widgetClassListStylePropertiesFunction *gi.Function
var widgetClassListStylePropertiesFunction_Once sync.Once

func widgetClassListStylePropertiesFunction_Set() {
	widgetClassListStylePropertiesFunction_Once.Do(func() {
		widgetClassListStylePropertiesFunction = gi.FunctionInvokerNew("Gtk", "list_style_properties")
	})
}

var listStylePropertiesWidgetClassInvoker *gi.Function

// ListStyleProperties is a representation of the C type gtk_widget_class_list_style_properties.
func (recv *WidgetClass) ListStyleProperties() uint32 {
	widgetClassListStylePropertiesFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	widgetClassListStylePropertiesFunction.Invoke(inArgs[:], outArgs[:])

	out0 := outArgs[0].Uint32()

	return out0
}

// UNSUPPORTED : C value 'gtk_widget_class_set_accessible_role' : parameter 'role' of type 'Atk.Role' not supported

// UNSUPPORTED : C value 'gtk_widget_class_set_accessible_type' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_class_set_connect_func' : parameter 'connect_func' of type 'BuilderConnectFunc' not supported

var widgetClassSetCssNameFunction *gi.Function
var widgetClassSetCssNameFunction_Once sync.Once

func widgetClassSetCssNameFunction_Set() {
	widgetClassSetCssNameFunction_Once.Do(func() {
		widgetClassSetCssNameFunction = gi.FunctionInvokerNew("Gtk", "set_css_name")
	})
}

var setCssNameWidgetClassInvoker *gi.Function

// SetCssName is a representation of the C type gtk_widget_class_set_css_name.
func (recv *WidgetClass) SetCssName(name string) {
	widgetClassSetCssNameFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	widgetClassSetCssNameFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_widget_class_set_template' : parameter 'template_bytes' of type 'GLib.Bytes' not supported

var widgetClassSetTemplateFromResourceFunction *gi.Function
var widgetClassSetTemplateFromResourceFunction_Once sync.Once

func widgetClassSetTemplateFromResourceFunction_Set() {
	widgetClassSetTemplateFromResourceFunction_Once.Do(func() {
		widgetClassSetTemplateFromResourceFunction = gi.FunctionInvokerNew("Gtk", "set_template_from_resource")
	})
}

var setTemplateFromResourceWidgetClassInvoker *gi.Function

// SetTemplateFromResource is a representation of the C type gtk_widget_class_set_template_from_resource.
func (recv *WidgetClass) SetTemplateFromResource(resourceName string) {
	widgetClassSetTemplateFromResourceFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(resourceName)

	widgetClassSetTemplateFromResourceFunction.Invoke(inArgs[:], nil)

}

var widgetClassPrivateStruct *gi.Struct
var widgetClassPrivateStructOnce sync.Once

func widgetClassPrivateStructSet() {
	widgetClassPrivateStructOnce.Do(func() {
		widgetClassPrivateStruct = gi.StructNew("Gtk", "WidgetClassPrivate")
	})
}

type WidgetClassPrivate struct {
	native uintptr
}

var widgetPathStruct *gi.Struct
var widgetPathStructOnce sync.Once

func widgetPathStructSet() {
	widgetPathStructOnce.Do(func() {
		widgetPathStruct = gi.StructNew("Gtk", "WidgetPath")
	})
}

type WidgetPath struct {
	native uintptr
}

var widgetPathNewFunction *gi.Function
var widgetPathNewFunction_Once sync.Once

func widgetPathNewFunction_Set() {
	widgetPathNewFunction_Once.Do(func() {
		widgetPathNewFunction = gi.FunctionInvokerNew("Gtk", "new")
	})
}

var newWidgetPathInvoker *gi.Function

// WidgetPathNew is a representation of the C type gtk_widget_path_new.
func WidgetPathNew() *WidgetPath {
	widgetPathNewFunction_Set()

	ret := widgetPathNewFunction.Invoke(nil, nil)

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_append_for_widget' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_widget_path_append_type' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_append_with_siblings' : parameter 'siblings' of type 'WidgetPath' not supported

var widgetPathCopyFunction *gi.Function
var widgetPathCopyFunction_Once sync.Once

func widgetPathCopyFunction_Set() {
	widgetPathCopyFunction_Once.Do(func() {
		widgetPathCopyFunction = gi.FunctionInvokerNew("Gtk", "copy")
	})
}

var copyWidgetPathInvoker *gi.Function

// Copy is a representation of the C type gtk_widget_path_copy.
func (recv *WidgetPath) Copy() *WidgetPath {
	widgetPathCopyFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := widgetPathCopyFunction.Invoke(inArgs[:], nil)

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

var widgetPathFreeFunction *gi.Function
var widgetPathFreeFunction_Once sync.Once

func widgetPathFreeFunction_Set() {
	widgetPathFreeFunction_Once.Do(func() {
		widgetPathFreeFunction = gi.FunctionInvokerNew("Gtk", "free")
	})
}

var freeWidgetPathInvoker *gi.Function

// Free is a representation of the C type gtk_widget_path_free.
func (recv *WidgetPath) Free() {
	widgetPathFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	widgetPathFreeFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_widget_path_get_object_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_has_parent' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_is_type' : parameter 'type' of type 'GType' not supported

var widgetPathIterAddClassFunction *gi.Function
var widgetPathIterAddClassFunction_Once sync.Once

func widgetPathIterAddClassFunction_Set() {
	widgetPathIterAddClassFunction_Once.Do(func() {
		widgetPathIterAddClassFunction = gi.FunctionInvokerNew("Gtk", "iter_add_class")
	})
}

var iterAddClassWidgetPathInvoker *gi.Function

// IterAddClass is a representation of the C type gtk_widget_path_iter_add_class.
func (recv *WidgetPath) IterAddClass(pos int32, name string) {
	widgetPathIterAddClassFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	widgetPathIterAddClassFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_widget_path_iter_add_region' : parameter 'flags' of type 'RegionFlags' not supported

var widgetPathIterClearClassesFunction *gi.Function
var widgetPathIterClearClassesFunction_Once sync.Once

func widgetPathIterClearClassesFunction_Set() {
	widgetPathIterClearClassesFunction_Once.Do(func() {
		widgetPathIterClearClassesFunction = gi.FunctionInvokerNew("Gtk", "iter_clear_classes")
	})
}

var iterClearClassesWidgetPathInvoker *gi.Function

// IterClearClasses is a representation of the C type gtk_widget_path_iter_clear_classes.
func (recv *WidgetPath) IterClearClasses(pos int32) {
	widgetPathIterClearClassesFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	widgetPathIterClearClassesFunction.Invoke(inArgs[:], nil)

}

var widgetPathIterClearRegionsFunction *gi.Function
var widgetPathIterClearRegionsFunction_Once sync.Once

func widgetPathIterClearRegionsFunction_Set() {
	widgetPathIterClearRegionsFunction_Once.Do(func() {
		widgetPathIterClearRegionsFunction = gi.FunctionInvokerNew("Gtk", "iter_clear_regions")
	})
}

var iterClearRegionsWidgetPathInvoker *gi.Function

// IterClearRegions is a representation of the C type gtk_widget_path_iter_clear_regions.
func (recv *WidgetPath) IterClearRegions(pos int32) {
	widgetPathIterClearRegionsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	widgetPathIterClearRegionsFunction.Invoke(inArgs[:], nil)

}

var widgetPathIterGetNameFunction *gi.Function
var widgetPathIterGetNameFunction_Once sync.Once

func widgetPathIterGetNameFunction_Set() {
	widgetPathIterGetNameFunction_Once.Do(func() {
		widgetPathIterGetNameFunction = gi.FunctionInvokerNew("Gtk", "iter_get_name")
	})
}

var iterGetNameWidgetPathInvoker *gi.Function

// IterGetName is a representation of the C type gtk_widget_path_iter_get_name.
func (recv *WidgetPath) IterGetName(pos int32) string {
	widgetPathIterGetNameFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := widgetPathIterGetNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

var widgetPathIterGetObjectNameFunction *gi.Function
var widgetPathIterGetObjectNameFunction_Once sync.Once

func widgetPathIterGetObjectNameFunction_Set() {
	widgetPathIterGetObjectNameFunction_Once.Do(func() {
		widgetPathIterGetObjectNameFunction = gi.FunctionInvokerNew("Gtk", "iter_get_object_name")
	})
}

var iterGetObjectNameWidgetPathInvoker *gi.Function

// IterGetObjectName is a representation of the C type gtk_widget_path_iter_get_object_name.
func (recv *WidgetPath) IterGetObjectName(pos int32) string {
	widgetPathIterGetObjectNameFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := widgetPathIterGetObjectNameFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_iter_get_object_type' : return type 'GType' not supported

var widgetPathIterGetSiblingIndexFunction *gi.Function
var widgetPathIterGetSiblingIndexFunction_Once sync.Once

func widgetPathIterGetSiblingIndexFunction_Set() {
	widgetPathIterGetSiblingIndexFunction_Once.Do(func() {
		widgetPathIterGetSiblingIndexFunction = gi.FunctionInvokerNew("Gtk", "iter_get_sibling_index")
	})
}

var iterGetSiblingIndexWidgetPathInvoker *gi.Function

// IterGetSiblingIndex is a representation of the C type gtk_widget_path_iter_get_sibling_index.
func (recv *WidgetPath) IterGetSiblingIndex(pos int32) uint32 {
	widgetPathIterGetSiblingIndexFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := widgetPathIterGetSiblingIndexFunction.Invoke(inArgs[:], nil)

	retGo := ret.Uint32()

	return retGo
}

var widgetPathIterGetSiblingsFunction *gi.Function
var widgetPathIterGetSiblingsFunction_Once sync.Once

func widgetPathIterGetSiblingsFunction_Set() {
	widgetPathIterGetSiblingsFunction_Once.Do(func() {
		widgetPathIterGetSiblingsFunction = gi.FunctionInvokerNew("Gtk", "iter_get_siblings")
	})
}

var iterGetSiblingsWidgetPathInvoker *gi.Function

// IterGetSiblings is a representation of the C type gtk_widget_path_iter_get_siblings.
func (recv *WidgetPath) IterGetSiblings(pos int32) *WidgetPath {
	widgetPathIterGetSiblingsFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	ret := widgetPathIterGetSiblingsFunction.Invoke(inArgs[:], nil)

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

var widgetPathIterRemoveClassFunction *gi.Function
var widgetPathIterRemoveClassFunction_Once sync.Once

func widgetPathIterRemoveClassFunction_Set() {
	widgetPathIterRemoveClassFunction_Once.Do(func() {
		widgetPathIterRemoveClassFunction = gi.FunctionInvokerNew("Gtk", "iter_remove_class")
	})
}

var iterRemoveClassWidgetPathInvoker *gi.Function

// IterRemoveClass is a representation of the C type gtk_widget_path_iter_remove_class.
func (recv *WidgetPath) IterRemoveClass(pos int32, name string) {
	widgetPathIterRemoveClassFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	widgetPathIterRemoveClassFunction.Invoke(inArgs[:], nil)

}

var widgetPathIterRemoveRegionFunction *gi.Function
var widgetPathIterRemoveRegionFunction_Once sync.Once

func widgetPathIterRemoveRegionFunction_Set() {
	widgetPathIterRemoveRegionFunction_Once.Do(func() {
		widgetPathIterRemoveRegionFunction = gi.FunctionInvokerNew("Gtk", "iter_remove_region")
	})
}

var iterRemoveRegionWidgetPathInvoker *gi.Function

// IterRemoveRegion is a representation of the C type gtk_widget_path_iter_remove_region.
func (recv *WidgetPath) IterRemoveRegion(pos int32, name string) {
	widgetPathIterRemoveRegionFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	widgetPathIterRemoveRegionFunction.Invoke(inArgs[:], nil)

}

var widgetPathIterSetNameFunction *gi.Function
var widgetPathIterSetNameFunction_Once sync.Once

func widgetPathIterSetNameFunction_Set() {
	widgetPathIterSetNameFunction_Once.Do(func() {
		widgetPathIterSetNameFunction = gi.FunctionInvokerNew("Gtk", "iter_set_name")
	})
}

var iterSetNameWidgetPathInvoker *gi.Function

// IterSetName is a representation of the C type gtk_widget_path_iter_set_name.
func (recv *WidgetPath) IterSetName(pos int32, name string) {
	widgetPathIterSetNameFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	widgetPathIterSetNameFunction.Invoke(inArgs[:], nil)

}

var widgetPathIterSetObjectNameFunction *gi.Function
var widgetPathIterSetObjectNameFunction_Once sync.Once

func widgetPathIterSetObjectNameFunction_Set() {
	widgetPathIterSetObjectNameFunction_Once.Do(func() {
		widgetPathIterSetObjectNameFunction = gi.FunctionInvokerNew("Gtk", "iter_set_object_name")
	})
}

var iterSetObjectNameWidgetPathInvoker *gi.Function

// IterSetObjectName is a representation of the C type gtk_widget_path_iter_set_object_name.
func (recv *WidgetPath) IterSetObjectName(pos int32, name string) {
	widgetPathIterSetObjectNameFunction_Set()

	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	widgetPathIterSetObjectNameFunction.Invoke(inArgs[:], nil)

}

// UNSUPPORTED : C value 'gtk_widget_path_iter_set_object_type' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_set_state' : parameter 'state' of type 'StateFlags' not supported

var widgetPathLengthFunction *gi.Function
var widgetPathLengthFunction_Once sync.Once

func widgetPathLengthFunction_Set() {
	widgetPathLengthFunction_Once.Do(func() {
		widgetPathLengthFunction = gi.FunctionInvokerNew("Gtk", "length")
	})
}

var lengthWidgetPathInvoker *gi.Function

// Length is a representation of the C type gtk_widget_path_length.
func (recv *WidgetPath) Length() int32 {
	widgetPathLengthFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := widgetPathLengthFunction.Invoke(inArgs[:], nil)

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_prepend_type' : parameter 'type' of type 'GType' not supported

var widgetPathRefFunction *gi.Function
var widgetPathRefFunction_Once sync.Once

func widgetPathRefFunction_Set() {
	widgetPathRefFunction_Once.Do(func() {
		widgetPathRefFunction = gi.FunctionInvokerNew("Gtk", "ref")
	})
}

var refWidgetPathInvoker *gi.Function

// Ref is a representation of the C type gtk_widget_path_ref.
func (recv *WidgetPath) Ref() *WidgetPath {
	widgetPathRefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := widgetPathRefFunction.Invoke(inArgs[:], nil)

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

var widgetPathToStringFunction *gi.Function
var widgetPathToStringFunction_Once sync.Once

func widgetPathToStringFunction_Set() {
	widgetPathToStringFunction_Once.Do(func() {
		widgetPathToStringFunction = gi.FunctionInvokerNew("Gtk", "to_string")
	})
}

var toStringWidgetPathInvoker *gi.Function

// ToString is a representation of the C type gtk_widget_path_to_string.
func (recv *WidgetPath) ToString() string {
	widgetPathToStringFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	ret := widgetPathToStringFunction.Invoke(inArgs[:], nil)

	retGo := ret.String(true)

	return retGo
}

var widgetPathUnrefFunction *gi.Function
var widgetPathUnrefFunction_Once sync.Once

func widgetPathUnrefFunction_Set() {
	widgetPathUnrefFunction_Once.Do(func() {
		widgetPathUnrefFunction = gi.FunctionInvokerNew("Gtk", "unref")
	})
}

var unrefWidgetPathInvoker *gi.Function

// Unref is a representation of the C type gtk_widget_path_unref.
func (recv *WidgetPath) Unref() {
	widgetPathUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	widgetPathUnrefFunction.Invoke(inArgs[:], nil)

}

var widgetPrivateStruct *gi.Struct
var widgetPrivateStructOnce sync.Once

func widgetPrivateStructSet() {
	widgetPrivateStructOnce.Do(func() {
		widgetPrivateStruct = gi.StructNew("Gtk", "WidgetPrivate")
	})
}

type WidgetPrivate struct {
	native uintptr
}

var windowAccessibleClassStruct *gi.Struct
var windowAccessibleClassStructOnce sync.Once

func windowAccessibleClassStructSet() {
	windowAccessibleClassStructOnce.Do(func() {
		windowAccessibleClassStruct = gi.StructNew("Gtk", "WindowAccessibleClass")
	})
}

type WindowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var windowAccessiblePrivateStruct *gi.Struct
var windowAccessiblePrivateStructOnce sync.Once

func windowAccessiblePrivateStructSet() {
	windowAccessiblePrivateStructOnce.Do(func() {
		windowAccessiblePrivateStruct = gi.StructNew("Gtk", "WindowAccessiblePrivate")
	})
}

type WindowAccessiblePrivate struct {
	native uintptr
}

var windowClassStruct *gi.Struct
var windowClassStructOnce sync.Once

func windowClassStructSet() {
	windowClassStructOnce.Do(func() {
		windowClassStruct = gi.StructNew("Gtk", "WindowClass")
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

var windowGeometryInfoStruct *gi.Struct
var windowGeometryInfoStructOnce sync.Once

func windowGeometryInfoStructSet() {
	windowGeometryInfoStructOnce.Do(func() {
		windowGeometryInfoStruct = gi.StructNew("Gtk", "WindowGeometryInfo")
	})
}

type WindowGeometryInfo struct {
	native uintptr
}

var windowGroupClassStruct *gi.Struct
var windowGroupClassStructOnce sync.Once

func windowGroupClassStructSet() {
	windowGroupClassStructOnce.Do(func() {
		windowGroupClassStruct = gi.StructNew("Gtk", "WindowGroupClass")
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

var windowGroupPrivateStruct *gi.Struct
var windowGroupPrivateStructOnce sync.Once

func windowGroupPrivateStructSet() {
	windowGroupPrivateStructOnce.Do(func() {
		windowGroupPrivateStruct = gi.StructNew("Gtk", "WindowGroupPrivate")
	})
}

type WindowGroupPrivate struct {
	native uintptr
}

var windowPrivateStruct *gi.Struct
var windowPrivateStructOnce sync.Once

func windowPrivateStructSet() {
	windowPrivateStructOnce.Do(func() {
		windowPrivateStruct = gi.StructNew("Gtk", "WindowPrivate")
	})
}

type WindowPrivate struct {
	native uintptr
}
