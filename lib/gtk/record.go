// Code generated - DO NOT EDIT.

package gtk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var aboutDialogClassStruct *gi.Struct
var aboutDialogClassStruct_Once sync.Once

func aboutDialogClassStruct_Set() {
	aboutDialogClassStruct_Once.Do(func() {
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
var aboutDialogPrivateStruct_Once sync.Once

func aboutDialogPrivateStruct_Set() {
	aboutDialogPrivateStruct_Once.Do(func() {
		aboutDialogPrivateStruct = gi.StructNew("Gtk", "AboutDialogPrivate")
	})
}

type AboutDialogPrivate struct {
	native uintptr
}

var accelGroupClassStruct *gi.Struct
var accelGroupClassStruct_Once sync.Once

func accelGroupClassStruct_Set() {
	accelGroupClassStruct_Once.Do(func() {
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
var accelGroupEntryStruct_Once sync.Once

func accelGroupEntryStruct_Set() {
	accelGroupEntryStruct_Once.Do(func() {
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
var accelGroupPrivateStruct_Once sync.Once

func accelGroupPrivateStruct_Set() {
	accelGroupPrivateStruct_Once.Do(func() {
		accelGroupPrivateStruct = gi.StructNew("Gtk", "AccelGroupPrivate")
	})
}

type AccelGroupPrivate struct {
	native uintptr
}

var accelKeyStruct *gi.Struct
var accelKeyStruct_Once sync.Once

func accelKeyStruct_Set() {
	accelKeyStruct_Once.Do(func() {
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
var accelLabelClassStruct_Once sync.Once

func accelLabelClassStruct_Set() {
	accelLabelClassStruct_Once.Do(func() {
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
var accelLabelPrivateStruct_Once sync.Once

func accelLabelPrivateStruct_Set() {
	accelLabelPrivateStruct_Once.Do(func() {
		accelLabelPrivateStruct = gi.StructNew("Gtk", "AccelLabelPrivate")
	})
}

type AccelLabelPrivate struct {
	native uintptr
}

var accelMapClassStruct *gi.Struct
var accelMapClassStruct_Once sync.Once

func accelMapClassStruct_Set() {
	accelMapClassStruct_Once.Do(func() {
		accelMapClassStruct = gi.StructNew("Gtk", "AccelMapClass")
	})
}

type AccelMapClass struct {
	native uintptr
}

var accessibleClassStruct *gi.Struct
var accessibleClassStruct_Once sync.Once

func accessibleClassStruct_Set() {
	accessibleClassStruct_Once.Do(func() {
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
var accessiblePrivateStruct_Once sync.Once

func accessiblePrivateStruct_Set() {
	accessiblePrivateStruct_Once.Do(func() {
		accessiblePrivateStruct = gi.StructNew("Gtk", "AccessiblePrivate")
	})
}

type AccessiblePrivate struct {
	native uintptr
}

var actionBarClassStruct *gi.Struct
var actionBarClassStruct_Once sync.Once

func actionBarClassStruct_Set() {
	actionBarClassStruct_Once.Do(func() {
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
var actionBarPrivateStruct_Once sync.Once

func actionBarPrivateStruct_Set() {
	actionBarPrivateStruct_Once.Do(func() {
		actionBarPrivateStruct = gi.StructNew("Gtk", "ActionBarPrivate")
	})
}

type ActionBarPrivate struct {
	native uintptr
}

var actionClassStruct *gi.Struct
var actionClassStruct_Once sync.Once

func actionClassStruct_Set() {
	actionClassStruct_Once.Do(func() {
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
var actionEntryStruct_Once sync.Once

func actionEntryStruct_Set() {
	actionEntryStruct_Once.Do(func() {
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
var actionGroupClassStruct_Once sync.Once

func actionGroupClassStruct_Set() {
	actionGroupClassStruct_Once.Do(func() {
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
var actionGroupPrivateStruct_Once sync.Once

func actionGroupPrivateStruct_Set() {
	actionGroupPrivateStruct_Once.Do(func() {
		actionGroupPrivateStruct = gi.StructNew("Gtk", "ActionGroupPrivate")
	})
}

type ActionGroupPrivate struct {
	native uintptr
}

var actionPrivateStruct *gi.Struct
var actionPrivateStruct_Once sync.Once

func actionPrivateStruct_Set() {
	actionPrivateStruct_Once.Do(func() {
		actionPrivateStruct = gi.StructNew("Gtk", "ActionPrivate")
	})
}

type ActionPrivate struct {
	native uintptr
}

var actionableInterfaceStruct *gi.Struct
var actionableInterfaceStruct_Once sync.Once

func actionableInterfaceStruct_Set() {
	actionableInterfaceStruct_Once.Do(func() {
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
var activatableIfaceStruct_Once sync.Once

func activatableIfaceStruct_Set() {
	activatableIfaceStruct_Once.Do(func() {
		activatableIfaceStruct = gi.StructNew("Gtk", "ActivatableIface")
	})
}

type ActivatableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'update' : missing Type
	// UNSUPPORTED : C value 'sync_action_properties' : missing Type
}

var adjustmentClassStruct *gi.Struct
var adjustmentClassStruct_Once sync.Once

func adjustmentClassStruct_Set() {
	adjustmentClassStruct_Once.Do(func() {
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
var adjustmentPrivateStruct_Once sync.Once

func adjustmentPrivateStruct_Set() {
	adjustmentPrivateStruct_Once.Do(func() {
		adjustmentPrivateStruct = gi.StructNew("Gtk", "AdjustmentPrivate")
	})
}

type AdjustmentPrivate struct {
	native uintptr
}

var alignmentClassStruct *gi.Struct
var alignmentClassStruct_Once sync.Once

func alignmentClassStruct_Set() {
	alignmentClassStruct_Once.Do(func() {
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
var alignmentPrivateStruct_Once sync.Once

func alignmentPrivateStruct_Set() {
	alignmentPrivateStruct_Once.Do(func() {
		alignmentPrivateStruct = gi.StructNew("Gtk", "AlignmentPrivate")
	})
}

type AlignmentPrivate struct {
	native uintptr
}

var appChooserButtonClassStruct *gi.Struct
var appChooserButtonClassStruct_Once sync.Once

func appChooserButtonClassStruct_Set() {
	appChooserButtonClassStruct_Once.Do(func() {
		appChooserButtonClassStruct = gi.StructNew("Gtk", "AppChooserButtonClass")
	})
}

type AppChooserButtonClass struct {
	native      uintptr
	ParentClass *ComboBoxClass
	// UNSUPPORTED : C value 'custom_item_activated' : missing Type
}

var appChooserButtonPrivateStruct *gi.Struct
var appChooserButtonPrivateStruct_Once sync.Once

func appChooserButtonPrivateStruct_Set() {
	appChooserButtonPrivateStruct_Once.Do(func() {
		appChooserButtonPrivateStruct = gi.StructNew("Gtk", "AppChooserButtonPrivate")
	})
}

type AppChooserButtonPrivate struct {
	native uintptr
}

var appChooserDialogClassStruct *gi.Struct
var appChooserDialogClassStruct_Once sync.Once

func appChooserDialogClassStruct_Set() {
	appChooserDialogClassStruct_Once.Do(func() {
		appChooserDialogClassStruct = gi.StructNew("Gtk", "AppChooserDialogClass")
	})
}

type AppChooserDialogClass struct {
	native      uintptr
	ParentClass *DialogClass
}

var appChooserDialogPrivateStruct *gi.Struct
var appChooserDialogPrivateStruct_Once sync.Once

func appChooserDialogPrivateStruct_Set() {
	appChooserDialogPrivateStruct_Once.Do(func() {
		appChooserDialogPrivateStruct = gi.StructNew("Gtk", "AppChooserDialogPrivate")
	})
}

type AppChooserDialogPrivate struct {
	native uintptr
}

var appChooserWidgetClassStruct *gi.Struct
var appChooserWidgetClassStruct_Once sync.Once

func appChooserWidgetClassStruct_Set() {
	appChooserWidgetClassStruct_Once.Do(func() {
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
var appChooserWidgetPrivateStruct_Once sync.Once

func appChooserWidgetPrivateStruct_Set() {
	appChooserWidgetPrivateStruct_Once.Do(func() {
		appChooserWidgetPrivateStruct = gi.StructNew("Gtk", "AppChooserWidgetPrivate")
	})
}

type AppChooserWidgetPrivate struct {
	native uintptr
}

var applicationClassStruct *gi.Struct
var applicationClassStruct_Once sync.Once

func applicationClassStruct_Set() {
	applicationClassStruct_Once.Do(func() {
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
var applicationPrivateStruct_Once sync.Once

func applicationPrivateStruct_Set() {
	applicationPrivateStruct_Once.Do(func() {
		applicationPrivateStruct = gi.StructNew("Gtk", "ApplicationPrivate")
	})
}

type ApplicationPrivate struct {
	native uintptr
}

var applicationWindowClassStruct *gi.Struct
var applicationWindowClassStruct_Once sync.Once

func applicationWindowClassStruct_Set() {
	applicationWindowClassStruct_Once.Do(func() {
		applicationWindowClassStruct = gi.StructNew("Gtk", "ApplicationWindowClass")
	})
}

type ApplicationWindowClass struct {
	native      uintptr
	ParentClass *WindowClass
}

var applicationWindowPrivateStruct *gi.Struct
var applicationWindowPrivateStruct_Once sync.Once

func applicationWindowPrivateStruct_Set() {
	applicationWindowPrivateStruct_Once.Do(func() {
		applicationWindowPrivateStruct = gi.StructNew("Gtk", "ApplicationWindowPrivate")
	})
}

type ApplicationWindowPrivate struct {
	native uintptr
}

var arrowAccessibleClassStruct *gi.Struct
var arrowAccessibleClassStruct_Once sync.Once

func arrowAccessibleClassStruct_Set() {
	arrowAccessibleClassStruct_Once.Do(func() {
		arrowAccessibleClassStruct = gi.StructNew("Gtk", "ArrowAccessibleClass")
	})
}

type ArrowAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var arrowAccessiblePrivateStruct *gi.Struct
var arrowAccessiblePrivateStruct_Once sync.Once

func arrowAccessiblePrivateStruct_Set() {
	arrowAccessiblePrivateStruct_Once.Do(func() {
		arrowAccessiblePrivateStruct = gi.StructNew("Gtk", "ArrowAccessiblePrivate")
	})
}

type ArrowAccessiblePrivate struct {
	native uintptr
}

var arrowClassStruct *gi.Struct
var arrowClassStruct_Once sync.Once

func arrowClassStruct_Set() {
	arrowClassStruct_Once.Do(func() {
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
var arrowPrivateStruct_Once sync.Once

func arrowPrivateStruct_Set() {
	arrowPrivateStruct_Once.Do(func() {
		arrowPrivateStruct = gi.StructNew("Gtk", "ArrowPrivate")
	})
}

type ArrowPrivate struct {
	native uintptr
}

var aspectFrameClassStruct *gi.Struct
var aspectFrameClassStruct_Once sync.Once

func aspectFrameClassStruct_Set() {
	aspectFrameClassStruct_Once.Do(func() {
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
var aspectFramePrivateStruct_Once sync.Once

func aspectFramePrivateStruct_Set() {
	aspectFramePrivateStruct_Once.Do(func() {
		aspectFramePrivateStruct = gi.StructNew("Gtk", "AspectFramePrivate")
	})
}

type AspectFramePrivate struct {
	native uintptr
}

var assistantClassStruct *gi.Struct
var assistantClassStruct_Once sync.Once

func assistantClassStruct_Set() {
	assistantClassStruct_Once.Do(func() {
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
var assistantPrivateStruct_Once sync.Once

func assistantPrivateStruct_Set() {
	assistantPrivateStruct_Once.Do(func() {
		assistantPrivateStruct = gi.StructNew("Gtk", "AssistantPrivate")
	})
}

type AssistantPrivate struct {
	native uintptr
}

var binClassStruct *gi.Struct
var binClassStruct_Once sync.Once

func binClassStruct_Set() {
	binClassStruct_Once.Do(func() {
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
var binPrivateStruct_Once sync.Once

func binPrivateStruct_Set() {
	binPrivateStruct_Once.Do(func() {
		binPrivateStruct = gi.StructNew("Gtk", "BinPrivate")
	})
}

type BinPrivate struct {
	native uintptr
}

var bindingArgStruct *gi.Struct
var bindingArgStruct_Once sync.Once

func bindingArgStruct_Set() {
	bindingArgStruct_Once.Do(func() {
		bindingArgStruct = gi.StructNew("Gtk", "BindingArg")
	})
}

type BindingArg struct {
	native uintptr
	// UNSUPPORTED : C value 'arg_type' : no Go type for 'GType'
}

var bindingEntryStruct *gi.Struct
var bindingEntryStruct_Once sync.Once

func bindingEntryStruct_Set() {
	bindingEntryStruct_Once.Do(func() {
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
var bindingSetStruct_Once sync.Once

func bindingSetStruct_Set() {
	bindingSetStruct_Once.Do(func() {
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
var bindingSignalStruct_Once sync.Once

func bindingSignalStruct_Set() {
	bindingSignalStruct_Once.Do(func() {
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
var booleanCellAccessibleClassStruct_Once sync.Once

func booleanCellAccessibleClassStruct_Set() {
	booleanCellAccessibleClassStruct_Once.Do(func() {
		booleanCellAccessibleClassStruct = gi.StructNew("Gtk", "BooleanCellAccessibleClass")
	})
}

type BooleanCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

var booleanCellAccessiblePrivateStruct *gi.Struct
var booleanCellAccessiblePrivateStruct_Once sync.Once

func booleanCellAccessiblePrivateStruct_Set() {
	booleanCellAccessiblePrivateStruct_Once.Do(func() {
		booleanCellAccessiblePrivateStruct = gi.StructNew("Gtk", "BooleanCellAccessiblePrivate")
	})
}

type BooleanCellAccessiblePrivate struct {
	native uintptr
}

var borderStruct *gi.Struct
var borderStruct_Once sync.Once

func borderStruct_Set() {
	borderStruct_Once.Do(func() {
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
		borderStruct_Set()
		borderNewFunction = borderStruct.InvokerNew("new")
	})
}

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
		borderStruct_Set()
		borderCopyFunction = borderStruct.InvokerNew("copy")
	})
}

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
		borderStruct_Set()
		borderFreeFunction = borderStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type gtk_border_free.
func (recv *Border) Free() {
	borderFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	borderFreeFunction.Invoke(inArgs[:], nil)

}

var boxClassStruct *gi.Struct
var boxClassStruct_Once sync.Once

func boxClassStruct_Set() {
	boxClassStruct_Once.Do(func() {
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
var boxPrivateStruct_Once sync.Once

func boxPrivateStruct_Set() {
	boxPrivateStruct_Once.Do(func() {
		boxPrivateStruct = gi.StructNew("Gtk", "BoxPrivate")
	})
}

type BoxPrivate struct {
	native uintptr
}

var buildableIfaceStruct *gi.Struct
var buildableIfaceStruct_Once sync.Once

func buildableIfaceStruct_Set() {
	buildableIfaceStruct_Once.Do(func() {
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
var builderClassStruct_Once sync.Once

func builderClassStruct_Set() {
	builderClassStruct_Once.Do(func() {
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
var builderPrivateStruct_Once sync.Once

func builderPrivateStruct_Set() {
	builderPrivateStruct_Once.Do(func() {
		builderPrivateStruct = gi.StructNew("Gtk", "BuilderPrivate")
	})
}

type BuilderPrivate struct {
	native uintptr
}

var buttonAccessibleClassStruct *gi.Struct
var buttonAccessibleClassStruct_Once sync.Once

func buttonAccessibleClassStruct_Set() {
	buttonAccessibleClassStruct_Once.Do(func() {
		buttonAccessibleClassStruct = gi.StructNew("Gtk", "ButtonAccessibleClass")
	})
}

type ButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var buttonAccessiblePrivateStruct *gi.Struct
var buttonAccessiblePrivateStruct_Once sync.Once

func buttonAccessiblePrivateStruct_Set() {
	buttonAccessiblePrivateStruct_Once.Do(func() {
		buttonAccessiblePrivateStruct = gi.StructNew("Gtk", "ButtonAccessiblePrivate")
	})
}

type ButtonAccessiblePrivate struct {
	native uintptr
}

var buttonBoxClassStruct *gi.Struct
var buttonBoxClassStruct_Once sync.Once

func buttonBoxClassStruct_Set() {
	buttonBoxClassStruct_Once.Do(func() {
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
var buttonBoxPrivateStruct_Once sync.Once

func buttonBoxPrivateStruct_Set() {
	buttonBoxPrivateStruct_Once.Do(func() {
		buttonBoxPrivateStruct = gi.StructNew("Gtk", "ButtonBoxPrivate")
	})
}

type ButtonBoxPrivate struct {
	native uintptr
}

var buttonClassStruct *gi.Struct
var buttonClassStruct_Once sync.Once

func buttonClassStruct_Set() {
	buttonClassStruct_Once.Do(func() {
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
var buttonPrivateStruct_Once sync.Once

func buttonPrivateStruct_Set() {
	buttonPrivateStruct_Once.Do(func() {
		buttonPrivateStruct = gi.StructNew("Gtk", "ButtonPrivate")
	})
}

type ButtonPrivate struct {
	native uintptr
}

var calendarClassStruct *gi.Struct
var calendarClassStruct_Once sync.Once

func calendarClassStruct_Set() {
	calendarClassStruct_Once.Do(func() {
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
var calendarPrivateStruct_Once sync.Once

func calendarPrivateStruct_Set() {
	calendarPrivateStruct_Once.Do(func() {
		calendarPrivateStruct = gi.StructNew("Gtk", "CalendarPrivate")
	})
}

type CalendarPrivate struct {
	native uintptr
}

var cellAccessibleClassStruct *gi.Struct
var cellAccessibleClassStruct_Once sync.Once

func cellAccessibleClassStruct_Set() {
	cellAccessibleClassStruct_Once.Do(func() {
		cellAccessibleClassStruct = gi.StructNew("Gtk", "CellAccessibleClass")
	})
}

type CellAccessibleClass struct {
	native      uintptr
	ParentClass *AccessibleClass
	// UNSUPPORTED : C value 'update_cache' : missing Type
}

var cellAccessibleParentIfaceStruct *gi.Struct
var cellAccessibleParentIfaceStruct_Once sync.Once

func cellAccessibleParentIfaceStruct_Set() {
	cellAccessibleParentIfaceStruct_Once.Do(func() {
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
var cellAccessiblePrivateStruct_Once sync.Once

func cellAccessiblePrivateStruct_Set() {
	cellAccessiblePrivateStruct_Once.Do(func() {
		cellAccessiblePrivateStruct = gi.StructNew("Gtk", "CellAccessiblePrivate")
	})
}

type CellAccessiblePrivate struct {
	native uintptr
}

var cellAreaBoxClassStruct *gi.Struct
var cellAreaBoxClassStruct_Once sync.Once

func cellAreaBoxClassStruct_Set() {
	cellAreaBoxClassStruct_Once.Do(func() {
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
var cellAreaBoxPrivateStruct_Once sync.Once

func cellAreaBoxPrivateStruct_Set() {
	cellAreaBoxPrivateStruct_Once.Do(func() {
		cellAreaBoxPrivateStruct = gi.StructNew("Gtk", "CellAreaBoxPrivate")
	})
}

type CellAreaBoxPrivate struct {
	native uintptr
}

var cellAreaClassStruct *gi.Struct
var cellAreaClassStruct_Once sync.Once

func cellAreaClassStruct_Set() {
	cellAreaClassStruct_Once.Do(func() {
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
		cellAreaClassStruct_Set()
		cellAreaClassListCellPropertiesFunction = cellAreaClassStruct.InvokerNew("list_cell_properties")
	})
}

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
var cellAreaContextClassStruct_Once sync.Once

func cellAreaContextClassStruct_Set() {
	cellAreaContextClassStruct_Once.Do(func() {
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
var cellAreaContextPrivateStruct_Once sync.Once

func cellAreaContextPrivateStruct_Set() {
	cellAreaContextPrivateStruct_Once.Do(func() {
		cellAreaContextPrivateStruct = gi.StructNew("Gtk", "CellAreaContextPrivate")
	})
}

type CellAreaContextPrivate struct {
	native uintptr
}

var cellAreaPrivateStruct *gi.Struct
var cellAreaPrivateStruct_Once sync.Once

func cellAreaPrivateStruct_Set() {
	cellAreaPrivateStruct_Once.Do(func() {
		cellAreaPrivateStruct = gi.StructNew("Gtk", "CellAreaPrivate")
	})
}

type CellAreaPrivate struct {
	native uintptr
}

var cellEditableIfaceStruct *gi.Struct
var cellEditableIfaceStruct_Once sync.Once

func cellEditableIfaceStruct_Set() {
	cellEditableIfaceStruct_Once.Do(func() {
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
var cellLayoutIfaceStruct_Once sync.Once

func cellLayoutIfaceStruct_Set() {
	cellLayoutIfaceStruct_Once.Do(func() {
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
var cellRendererAccelClassStruct_Once sync.Once

func cellRendererAccelClassStruct_Set() {
	cellRendererAccelClassStruct_Once.Do(func() {
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
var cellRendererAccelPrivateStruct_Once sync.Once

func cellRendererAccelPrivateStruct_Set() {
	cellRendererAccelPrivateStruct_Once.Do(func() {
		cellRendererAccelPrivateStruct = gi.StructNew("Gtk", "CellRendererAccelPrivate")
	})
}

type CellRendererAccelPrivate struct {
	native uintptr
}

var cellRendererClassStruct *gi.Struct
var cellRendererClassStruct_Once sync.Once

func cellRendererClassStruct_Set() {
	cellRendererClassStruct_Once.Do(func() {
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
var cellRendererClassPrivateStruct_Once sync.Once

func cellRendererClassPrivateStruct_Set() {
	cellRendererClassPrivateStruct_Once.Do(func() {
		cellRendererClassPrivateStruct = gi.StructNew("Gtk", "CellRendererClassPrivate")
	})
}

type CellRendererClassPrivate struct {
	native uintptr
}

var cellRendererComboClassStruct *gi.Struct
var cellRendererComboClassStruct_Once sync.Once

func cellRendererComboClassStruct_Set() {
	cellRendererComboClassStruct_Once.Do(func() {
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
var cellRendererComboPrivateStruct_Once sync.Once

func cellRendererComboPrivateStruct_Set() {
	cellRendererComboPrivateStruct_Once.Do(func() {
		cellRendererComboPrivateStruct = gi.StructNew("Gtk", "CellRendererComboPrivate")
	})
}

type CellRendererComboPrivate struct {
	native uintptr
}

var cellRendererPixbufClassStruct *gi.Struct
var cellRendererPixbufClassStruct_Once sync.Once

func cellRendererPixbufClassStruct_Set() {
	cellRendererPixbufClassStruct_Once.Do(func() {
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
var cellRendererPixbufPrivateStruct_Once sync.Once

func cellRendererPixbufPrivateStruct_Set() {
	cellRendererPixbufPrivateStruct_Once.Do(func() {
		cellRendererPixbufPrivateStruct = gi.StructNew("Gtk", "CellRendererPixbufPrivate")
	})
}

type CellRendererPixbufPrivate struct {
	native uintptr
}

var cellRendererPrivateStruct *gi.Struct
var cellRendererPrivateStruct_Once sync.Once

func cellRendererPrivateStruct_Set() {
	cellRendererPrivateStruct_Once.Do(func() {
		cellRendererPrivateStruct = gi.StructNew("Gtk", "CellRendererPrivate")
	})
}

type CellRendererPrivate struct {
	native uintptr
}

var cellRendererProgressClassStruct *gi.Struct
var cellRendererProgressClassStruct_Once sync.Once

func cellRendererProgressClassStruct_Set() {
	cellRendererProgressClassStruct_Once.Do(func() {
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
var cellRendererProgressPrivateStruct_Once sync.Once

func cellRendererProgressPrivateStruct_Set() {
	cellRendererProgressPrivateStruct_Once.Do(func() {
		cellRendererProgressPrivateStruct = gi.StructNew("Gtk", "CellRendererProgressPrivate")
	})
}

type CellRendererProgressPrivate struct {
	native uintptr
}

var cellRendererSpinClassStruct *gi.Struct
var cellRendererSpinClassStruct_Once sync.Once

func cellRendererSpinClassStruct_Set() {
	cellRendererSpinClassStruct_Once.Do(func() {
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
var cellRendererSpinPrivateStruct_Once sync.Once

func cellRendererSpinPrivateStruct_Set() {
	cellRendererSpinPrivateStruct_Once.Do(func() {
		cellRendererSpinPrivateStruct = gi.StructNew("Gtk", "CellRendererSpinPrivate")
	})
}

type CellRendererSpinPrivate struct {
	native uintptr
}

var cellRendererSpinnerClassStruct *gi.Struct
var cellRendererSpinnerClassStruct_Once sync.Once

func cellRendererSpinnerClassStruct_Set() {
	cellRendererSpinnerClassStruct_Once.Do(func() {
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
var cellRendererSpinnerPrivateStruct_Once sync.Once

func cellRendererSpinnerPrivateStruct_Set() {
	cellRendererSpinnerPrivateStruct_Once.Do(func() {
		cellRendererSpinnerPrivateStruct = gi.StructNew("Gtk", "CellRendererSpinnerPrivate")
	})
}

type CellRendererSpinnerPrivate struct {
	native uintptr
}

var cellRendererTextClassStruct *gi.Struct
var cellRendererTextClassStruct_Once sync.Once

func cellRendererTextClassStruct_Set() {
	cellRendererTextClassStruct_Once.Do(func() {
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
var cellRendererTextPrivateStruct_Once sync.Once

func cellRendererTextPrivateStruct_Set() {
	cellRendererTextPrivateStruct_Once.Do(func() {
		cellRendererTextPrivateStruct = gi.StructNew("Gtk", "CellRendererTextPrivate")
	})
}

type CellRendererTextPrivate struct {
	native uintptr
}

var cellRendererToggleClassStruct *gi.Struct
var cellRendererToggleClassStruct_Once sync.Once

func cellRendererToggleClassStruct_Set() {
	cellRendererToggleClassStruct_Once.Do(func() {
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
var cellRendererTogglePrivateStruct_Once sync.Once

func cellRendererTogglePrivateStruct_Set() {
	cellRendererTogglePrivateStruct_Once.Do(func() {
		cellRendererTogglePrivateStruct = gi.StructNew("Gtk", "CellRendererTogglePrivate")
	})
}

type CellRendererTogglePrivate struct {
	native uintptr
}

var cellViewClassStruct *gi.Struct
var cellViewClassStruct_Once sync.Once

func cellViewClassStruct_Set() {
	cellViewClassStruct_Once.Do(func() {
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
var cellViewPrivateStruct_Once sync.Once

func cellViewPrivateStruct_Set() {
	cellViewPrivateStruct_Once.Do(func() {
		cellViewPrivateStruct = gi.StructNew("Gtk", "CellViewPrivate")
	})
}

type CellViewPrivate struct {
	native uintptr
}

var checkButtonClassStruct *gi.Struct
var checkButtonClassStruct_Once sync.Once

func checkButtonClassStruct_Set() {
	checkButtonClassStruct_Once.Do(func() {
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
var checkMenuItemAccessibleClassStruct_Once sync.Once

func checkMenuItemAccessibleClassStruct_Set() {
	checkMenuItemAccessibleClassStruct_Once.Do(func() {
		checkMenuItemAccessibleClassStruct = gi.StructNew("Gtk", "CheckMenuItemAccessibleClass")
	})
}

type CheckMenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *MenuItemAccessibleClass
}

var checkMenuItemAccessiblePrivateStruct *gi.Struct
var checkMenuItemAccessiblePrivateStruct_Once sync.Once

func checkMenuItemAccessiblePrivateStruct_Set() {
	checkMenuItemAccessiblePrivateStruct_Once.Do(func() {
		checkMenuItemAccessiblePrivateStruct = gi.StructNew("Gtk", "CheckMenuItemAccessiblePrivate")
	})
}

type CheckMenuItemAccessiblePrivate struct {
	native uintptr
}

var checkMenuItemClassStruct *gi.Struct
var checkMenuItemClassStruct_Once sync.Once

func checkMenuItemClassStruct_Set() {
	checkMenuItemClassStruct_Once.Do(func() {
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
var checkMenuItemPrivateStruct_Once sync.Once

func checkMenuItemPrivateStruct_Set() {
	checkMenuItemPrivateStruct_Once.Do(func() {
		checkMenuItemPrivateStruct = gi.StructNew("Gtk", "CheckMenuItemPrivate")
	})
}

type CheckMenuItemPrivate struct {
	native uintptr
}

var colorButtonClassStruct *gi.Struct
var colorButtonClassStruct_Once sync.Once

func colorButtonClassStruct_Set() {
	colorButtonClassStruct_Once.Do(func() {
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
var colorButtonPrivateStruct_Once sync.Once

func colorButtonPrivateStruct_Set() {
	colorButtonPrivateStruct_Once.Do(func() {
		colorButtonPrivateStruct = gi.StructNew("Gtk", "ColorButtonPrivate")
	})
}

type ColorButtonPrivate struct {
	native uintptr
}

var colorChooserDialogClassStruct *gi.Struct
var colorChooserDialogClassStruct_Once sync.Once

func colorChooserDialogClassStruct_Set() {
	colorChooserDialogClassStruct_Once.Do(func() {
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
var colorChooserDialogPrivateStruct_Once sync.Once

func colorChooserDialogPrivateStruct_Set() {
	colorChooserDialogPrivateStruct_Once.Do(func() {
		colorChooserDialogPrivateStruct = gi.StructNew("Gtk", "ColorChooserDialogPrivate")
	})
}

type ColorChooserDialogPrivate struct {
	native uintptr
}

var colorChooserInterfaceStruct *gi.Struct
var colorChooserInterfaceStruct_Once sync.Once

func colorChooserInterfaceStruct_Set() {
	colorChooserInterfaceStruct_Once.Do(func() {
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
var colorChooserWidgetClassStruct_Once sync.Once

func colorChooserWidgetClassStruct_Set() {
	colorChooserWidgetClassStruct_Once.Do(func() {
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
var colorChooserWidgetPrivateStruct_Once sync.Once

func colorChooserWidgetPrivateStruct_Set() {
	colorChooserWidgetPrivateStruct_Once.Do(func() {
		colorChooserWidgetPrivateStruct = gi.StructNew("Gtk", "ColorChooserWidgetPrivate")
	})
}

type ColorChooserWidgetPrivate struct {
	native uintptr
}

var colorSelectionClassStruct *gi.Struct
var colorSelectionClassStruct_Once sync.Once

func colorSelectionClassStruct_Set() {
	colorSelectionClassStruct_Once.Do(func() {
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
var colorSelectionDialogClassStruct_Once sync.Once

func colorSelectionDialogClassStruct_Set() {
	colorSelectionDialogClassStruct_Once.Do(func() {
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
var colorSelectionDialogPrivateStruct_Once sync.Once

func colorSelectionDialogPrivateStruct_Set() {
	colorSelectionDialogPrivateStruct_Once.Do(func() {
		colorSelectionDialogPrivateStruct = gi.StructNew("Gtk", "ColorSelectionDialogPrivate")
	})
}

type ColorSelectionDialogPrivate struct {
	native uintptr
}

var colorSelectionPrivateStruct *gi.Struct
var colorSelectionPrivateStruct_Once sync.Once

func colorSelectionPrivateStruct_Set() {
	colorSelectionPrivateStruct_Once.Do(func() {
		colorSelectionPrivateStruct = gi.StructNew("Gtk", "ColorSelectionPrivate")
	})
}

type ColorSelectionPrivate struct {
	native uintptr
}

var comboBoxAccessibleClassStruct *gi.Struct
var comboBoxAccessibleClassStruct_Once sync.Once

func comboBoxAccessibleClassStruct_Set() {
	comboBoxAccessibleClassStruct_Once.Do(func() {
		comboBoxAccessibleClassStruct = gi.StructNew("Gtk", "ComboBoxAccessibleClass")
	})
}

type ComboBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var comboBoxAccessiblePrivateStruct *gi.Struct
var comboBoxAccessiblePrivateStruct_Once sync.Once

func comboBoxAccessiblePrivateStruct_Set() {
	comboBoxAccessiblePrivateStruct_Once.Do(func() {
		comboBoxAccessiblePrivateStruct = gi.StructNew("Gtk", "ComboBoxAccessiblePrivate")
	})
}

type ComboBoxAccessiblePrivate struct {
	native uintptr
}

var comboBoxClassStruct *gi.Struct
var comboBoxClassStruct_Once sync.Once

func comboBoxClassStruct_Set() {
	comboBoxClassStruct_Once.Do(func() {
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
var comboBoxPrivateStruct_Once sync.Once

func comboBoxPrivateStruct_Set() {
	comboBoxPrivateStruct_Once.Do(func() {
		comboBoxPrivateStruct = gi.StructNew("Gtk", "ComboBoxPrivate")
	})
}

type ComboBoxPrivate struct {
	native uintptr
}

var comboBoxTextClassStruct *gi.Struct
var comboBoxTextClassStruct_Once sync.Once

func comboBoxTextClassStruct_Set() {
	comboBoxTextClassStruct_Once.Do(func() {
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
var comboBoxTextPrivateStruct_Once sync.Once

func comboBoxTextPrivateStruct_Set() {
	comboBoxTextPrivateStruct_Once.Do(func() {
		comboBoxTextPrivateStruct = gi.StructNew("Gtk", "ComboBoxTextPrivate")
	})
}

type ComboBoxTextPrivate struct {
	native uintptr
}

var containerAccessibleClassStruct *gi.Struct
var containerAccessibleClassStruct_Once sync.Once

func containerAccessibleClassStruct_Set() {
	containerAccessibleClassStruct_Once.Do(func() {
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
var containerAccessiblePrivateStruct_Once sync.Once

func containerAccessiblePrivateStruct_Set() {
	containerAccessiblePrivateStruct_Once.Do(func() {
		containerAccessiblePrivateStruct = gi.StructNew("Gtk", "ContainerAccessiblePrivate")
	})
}

type ContainerAccessiblePrivate struct {
	native uintptr
}

var containerCellAccessibleClassStruct *gi.Struct
var containerCellAccessibleClassStruct_Once sync.Once

func containerCellAccessibleClassStruct_Set() {
	containerCellAccessibleClassStruct_Once.Do(func() {
		containerCellAccessibleClassStruct = gi.StructNew("Gtk", "ContainerCellAccessibleClass")
	})
}

type ContainerCellAccessibleClass struct {
	native      uintptr
	ParentClass *CellAccessibleClass
}

var containerCellAccessiblePrivateStruct *gi.Struct
var containerCellAccessiblePrivateStruct_Once sync.Once

func containerCellAccessiblePrivateStruct_Set() {
	containerCellAccessiblePrivateStruct_Once.Do(func() {
		containerCellAccessiblePrivateStruct = gi.StructNew("Gtk", "ContainerCellAccessiblePrivate")
	})
}

type ContainerCellAccessiblePrivate struct {
	native uintptr
}

var containerClassStruct *gi.Struct
var containerClassStruct_Once sync.Once

func containerClassStruct_Set() {
	containerClassStruct_Once.Do(func() {
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
		containerClassStruct_Set()
		containerClassHandleBorderWidthFunction = containerClassStruct.InvokerNew("handle_border_width")
	})
}

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
		containerClassStruct_Set()
		containerClassListChildPropertiesFunction = containerClassStruct.InvokerNew("list_child_properties")
	})
}

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
var containerPrivateStruct_Once sync.Once

func containerPrivateStruct_Set() {
	containerPrivateStruct_Once.Do(func() {
		containerPrivateStruct = gi.StructNew("Gtk", "ContainerPrivate")
	})
}

type ContainerPrivate struct {
	native uintptr
}

var cssProviderClassStruct *gi.Struct
var cssProviderClassStruct_Once sync.Once

func cssProviderClassStruct_Set() {
	cssProviderClassStruct_Once.Do(func() {
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
var cssProviderPrivateStruct_Once sync.Once

func cssProviderPrivateStruct_Set() {
	cssProviderPrivateStruct_Once.Do(func() {
		cssProviderPrivateStruct = gi.StructNew("Gtk", "CssProviderPrivate")
	})
}

type CssProviderPrivate struct {
	native uintptr
}

var cssSectionStruct *gi.Struct
var cssSectionStruct_Once sync.Once

func cssSectionStruct_Set() {
	cssSectionStruct_Once.Do(func() {
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
		cssSectionStruct_Set()
		cssSectionGetEndLineFunction = cssSectionStruct.InvokerNew("get_end_line")
	})
}

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
		cssSectionStruct_Set()
		cssSectionGetEndPositionFunction = cssSectionStruct.InvokerNew("get_end_position")
	})
}

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
		cssSectionStruct_Set()
		cssSectionGetParentFunction = cssSectionStruct.InvokerNew("get_parent")
	})
}

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
		cssSectionStruct_Set()
		cssSectionGetStartLineFunction = cssSectionStruct.InvokerNew("get_start_line")
	})
}

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
		cssSectionStruct_Set()
		cssSectionGetStartPositionFunction = cssSectionStruct.InvokerNew("get_start_position")
	})
}

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
		cssSectionStruct_Set()
		cssSectionRefFunction = cssSectionStruct.InvokerNew("ref")
	})
}

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
		cssSectionStruct_Set()
		cssSectionUnrefFunction = cssSectionStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type gtk_css_section_unref.
func (recv *CssSection) Unref() {
	cssSectionUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	cssSectionUnrefFunction.Invoke(inArgs[:], nil)

}

var dialogClassStruct *gi.Struct
var dialogClassStruct_Once sync.Once

func dialogClassStruct_Set() {
	dialogClassStruct_Once.Do(func() {
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
var dialogPrivateStruct_Once sync.Once

func dialogPrivateStruct_Set() {
	dialogPrivateStruct_Once.Do(func() {
		dialogPrivateStruct = gi.StructNew("Gtk", "DialogPrivate")
	})
}

type DialogPrivate struct {
	native uintptr
}

var drawingAreaClassStruct *gi.Struct
var drawingAreaClassStruct_Once sync.Once

func drawingAreaClassStruct_Set() {
	drawingAreaClassStruct_Once.Do(func() {
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
var editableInterfaceStruct_Once sync.Once

func editableInterfaceStruct_Set() {
	editableInterfaceStruct_Once.Do(func() {
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
var entryAccessibleClassStruct_Once sync.Once

func entryAccessibleClassStruct_Set() {
	entryAccessibleClassStruct_Once.Do(func() {
		entryAccessibleClassStruct = gi.StructNew("Gtk", "EntryAccessibleClass")
	})
}

type EntryAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var entryAccessiblePrivateStruct *gi.Struct
var entryAccessiblePrivateStruct_Once sync.Once

func entryAccessiblePrivateStruct_Set() {
	entryAccessiblePrivateStruct_Once.Do(func() {
		entryAccessiblePrivateStruct = gi.StructNew("Gtk", "EntryAccessiblePrivate")
	})
}

type EntryAccessiblePrivate struct {
	native uintptr
}

var entryBufferClassStruct *gi.Struct
var entryBufferClassStruct_Once sync.Once

func entryBufferClassStruct_Set() {
	entryBufferClassStruct_Once.Do(func() {
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
var entryBufferPrivateStruct_Once sync.Once

func entryBufferPrivateStruct_Set() {
	entryBufferPrivateStruct_Once.Do(func() {
		entryBufferPrivateStruct = gi.StructNew("Gtk", "EntryBufferPrivate")
	})
}

type EntryBufferPrivate struct {
	native uintptr
}

var entryClassStruct *gi.Struct
var entryClassStruct_Once sync.Once

func entryClassStruct_Set() {
	entryClassStruct_Once.Do(func() {
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
var entryCompletionClassStruct_Once sync.Once

func entryCompletionClassStruct_Set() {
	entryCompletionClassStruct_Once.Do(func() {
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
var entryCompletionPrivateStruct_Once sync.Once

func entryCompletionPrivateStruct_Set() {
	entryCompletionPrivateStruct_Once.Do(func() {
		entryCompletionPrivateStruct = gi.StructNew("Gtk", "EntryCompletionPrivate")
	})
}

type EntryCompletionPrivate struct {
	native uintptr
}

var entryPrivateStruct *gi.Struct
var entryPrivateStruct_Once sync.Once

func entryPrivateStruct_Set() {
	entryPrivateStruct_Once.Do(func() {
		entryPrivateStruct = gi.StructNew("Gtk", "EntryPrivate")
	})
}

type EntryPrivate struct {
	native uintptr
}

var eventBoxClassStruct *gi.Struct
var eventBoxClassStruct_Once sync.Once

func eventBoxClassStruct_Set() {
	eventBoxClassStruct_Once.Do(func() {
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
var eventBoxPrivateStruct_Once sync.Once

func eventBoxPrivateStruct_Set() {
	eventBoxPrivateStruct_Once.Do(func() {
		eventBoxPrivateStruct = gi.StructNew("Gtk", "EventBoxPrivate")
	})
}

type EventBoxPrivate struct {
	native uintptr
}

var eventControllerClassStruct *gi.Struct
var eventControllerClassStruct_Once sync.Once

func eventControllerClassStruct_Set() {
	eventControllerClassStruct_Once.Do(func() {
		eventControllerClassStruct = gi.StructNew("Gtk", "EventControllerClass")
	})
}

type EventControllerClass struct {
	native uintptr
}

var eventControllerKeyClassStruct *gi.Struct
var eventControllerKeyClassStruct_Once sync.Once

func eventControllerKeyClassStruct_Set() {
	eventControllerKeyClassStruct_Once.Do(func() {
		eventControllerKeyClassStruct = gi.StructNew("Gtk", "EventControllerKeyClass")
	})
}

type EventControllerKeyClass struct {
	native uintptr
}

var eventControllerMotionClassStruct *gi.Struct
var eventControllerMotionClassStruct_Once sync.Once

func eventControllerMotionClassStruct_Set() {
	eventControllerMotionClassStruct_Once.Do(func() {
		eventControllerMotionClassStruct = gi.StructNew("Gtk", "EventControllerMotionClass")
	})
}

type EventControllerMotionClass struct {
	native uintptr
}

var eventControllerScrollClassStruct *gi.Struct
var eventControllerScrollClassStruct_Once sync.Once

func eventControllerScrollClassStruct_Set() {
	eventControllerScrollClassStruct_Once.Do(func() {
		eventControllerScrollClassStruct = gi.StructNew("Gtk", "EventControllerScrollClass")
	})
}

type EventControllerScrollClass struct {
	native uintptr
}

var expanderAccessibleClassStruct *gi.Struct
var expanderAccessibleClassStruct_Once sync.Once

func expanderAccessibleClassStruct_Set() {
	expanderAccessibleClassStruct_Once.Do(func() {
		expanderAccessibleClassStruct = gi.StructNew("Gtk", "ExpanderAccessibleClass")
	})
}

type ExpanderAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var expanderAccessiblePrivateStruct *gi.Struct
var expanderAccessiblePrivateStruct_Once sync.Once

func expanderAccessiblePrivateStruct_Set() {
	expanderAccessiblePrivateStruct_Once.Do(func() {
		expanderAccessiblePrivateStruct = gi.StructNew("Gtk", "ExpanderAccessiblePrivate")
	})
}

type ExpanderAccessiblePrivate struct {
	native uintptr
}

var expanderClassStruct *gi.Struct
var expanderClassStruct_Once sync.Once

func expanderClassStruct_Set() {
	expanderClassStruct_Once.Do(func() {
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
var expanderPrivateStruct_Once sync.Once

func expanderPrivateStruct_Set() {
	expanderPrivateStruct_Once.Do(func() {
		expanderPrivateStruct = gi.StructNew("Gtk", "ExpanderPrivate")
	})
}

type ExpanderPrivate struct {
	native uintptr
}

var fileChooserButtonClassStruct *gi.Struct
var fileChooserButtonClassStruct_Once sync.Once

func fileChooserButtonClassStruct_Set() {
	fileChooserButtonClassStruct_Once.Do(func() {
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
var fileChooserButtonPrivateStruct_Once sync.Once

func fileChooserButtonPrivateStruct_Set() {
	fileChooserButtonPrivateStruct_Once.Do(func() {
		fileChooserButtonPrivateStruct = gi.StructNew("Gtk", "FileChooserButtonPrivate")
	})
}

type FileChooserButtonPrivate struct {
	native uintptr
}

var fileChooserDialogClassStruct *gi.Struct
var fileChooserDialogClassStruct_Once sync.Once

func fileChooserDialogClassStruct_Set() {
	fileChooserDialogClassStruct_Once.Do(func() {
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
var fileChooserDialogPrivateStruct_Once sync.Once

func fileChooserDialogPrivateStruct_Set() {
	fileChooserDialogPrivateStruct_Once.Do(func() {
		fileChooserDialogPrivateStruct = gi.StructNew("Gtk", "FileChooserDialogPrivate")
	})
}

type FileChooserDialogPrivate struct {
	native uintptr
}

var fileChooserNativeClassStruct *gi.Struct
var fileChooserNativeClassStruct_Once sync.Once

func fileChooserNativeClassStruct_Set() {
	fileChooserNativeClassStruct_Once.Do(func() {
		fileChooserNativeClassStruct = gi.StructNew("Gtk", "FileChooserNativeClass")
	})
}

type FileChooserNativeClass struct {
	native      uintptr
	ParentClass *NativeDialogClass
}

var fileChooserWidgetClassStruct *gi.Struct
var fileChooserWidgetClassStruct_Once sync.Once

func fileChooserWidgetClassStruct_Set() {
	fileChooserWidgetClassStruct_Once.Do(func() {
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
var fileChooserWidgetPrivateStruct_Once sync.Once

func fileChooserWidgetPrivateStruct_Set() {
	fileChooserWidgetPrivateStruct_Once.Do(func() {
		fileChooserWidgetPrivateStruct = gi.StructNew("Gtk", "FileChooserWidgetPrivate")
	})
}

type FileChooserWidgetPrivate struct {
	native uintptr
}

var fileFilterInfoStruct *gi.Struct
var fileFilterInfoStruct_Once sync.Once

func fileFilterInfoStruct_Set() {
	fileFilterInfoStruct_Once.Do(func() {
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
var fixedChildStruct_Once sync.Once

func fixedChildStruct_Set() {
	fixedChildStruct_Once.Do(func() {
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
var fixedClassStruct_Once sync.Once

func fixedClassStruct_Set() {
	fixedClassStruct_Once.Do(func() {
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
var fixedPrivateStruct_Once sync.Once

func fixedPrivateStruct_Set() {
	fixedPrivateStruct_Once.Do(func() {
		fixedPrivateStruct = gi.StructNew("Gtk", "FixedPrivate")
	})
}

type FixedPrivate struct {
	native uintptr
}

var flowBoxAccessibleClassStruct *gi.Struct
var flowBoxAccessibleClassStruct_Once sync.Once

func flowBoxAccessibleClassStruct_Set() {
	flowBoxAccessibleClassStruct_Once.Do(func() {
		flowBoxAccessibleClassStruct = gi.StructNew("Gtk", "FlowBoxAccessibleClass")
	})
}

type FlowBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var flowBoxAccessiblePrivateStruct *gi.Struct
var flowBoxAccessiblePrivateStruct_Once sync.Once

func flowBoxAccessiblePrivateStruct_Set() {
	flowBoxAccessiblePrivateStruct_Once.Do(func() {
		flowBoxAccessiblePrivateStruct = gi.StructNew("Gtk", "FlowBoxAccessiblePrivate")
	})
}

type FlowBoxAccessiblePrivate struct {
	native uintptr
}

var flowBoxChildAccessibleClassStruct *gi.Struct
var flowBoxChildAccessibleClassStruct_Once sync.Once

func flowBoxChildAccessibleClassStruct_Set() {
	flowBoxChildAccessibleClassStruct_Once.Do(func() {
		flowBoxChildAccessibleClassStruct = gi.StructNew("Gtk", "FlowBoxChildAccessibleClass")
	})
}

type FlowBoxChildAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var flowBoxChildClassStruct *gi.Struct
var flowBoxChildClassStruct_Once sync.Once

func flowBoxChildClassStruct_Set() {
	flowBoxChildClassStruct_Once.Do(func() {
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
var flowBoxClassStruct_Once sync.Once

func flowBoxClassStruct_Set() {
	flowBoxClassStruct_Once.Do(func() {
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
var fontButtonClassStruct_Once sync.Once

func fontButtonClassStruct_Set() {
	fontButtonClassStruct_Once.Do(func() {
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
var fontButtonPrivateStruct_Once sync.Once

func fontButtonPrivateStruct_Set() {
	fontButtonPrivateStruct_Once.Do(func() {
		fontButtonPrivateStruct = gi.StructNew("Gtk", "FontButtonPrivate")
	})
}

type FontButtonPrivate struct {
	native uintptr
}

var fontChooserDialogClassStruct *gi.Struct
var fontChooserDialogClassStruct_Once sync.Once

func fontChooserDialogClassStruct_Set() {
	fontChooserDialogClassStruct_Once.Do(func() {
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
var fontChooserDialogPrivateStruct_Once sync.Once

func fontChooserDialogPrivateStruct_Set() {
	fontChooserDialogPrivateStruct_Once.Do(func() {
		fontChooserDialogPrivateStruct = gi.StructNew("Gtk", "FontChooserDialogPrivate")
	})
}

type FontChooserDialogPrivate struct {
	native uintptr
}

var fontChooserIfaceStruct *gi.Struct
var fontChooserIfaceStruct_Once sync.Once

func fontChooserIfaceStruct_Set() {
	fontChooserIfaceStruct_Once.Do(func() {
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
var fontChooserWidgetClassStruct_Once sync.Once

func fontChooserWidgetClassStruct_Set() {
	fontChooserWidgetClassStruct_Once.Do(func() {
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
var fontChooserWidgetPrivateStruct_Once sync.Once

func fontChooserWidgetPrivateStruct_Set() {
	fontChooserWidgetPrivateStruct_Once.Do(func() {
		fontChooserWidgetPrivateStruct = gi.StructNew("Gtk", "FontChooserWidgetPrivate")
	})
}

type FontChooserWidgetPrivate struct {
	native uintptr
}

var fontSelectionClassStruct *gi.Struct
var fontSelectionClassStruct_Once sync.Once

func fontSelectionClassStruct_Set() {
	fontSelectionClassStruct_Once.Do(func() {
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
var fontSelectionDialogClassStruct_Once sync.Once

func fontSelectionDialogClassStruct_Set() {
	fontSelectionDialogClassStruct_Once.Do(func() {
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
var fontSelectionDialogPrivateStruct_Once sync.Once

func fontSelectionDialogPrivateStruct_Set() {
	fontSelectionDialogPrivateStruct_Once.Do(func() {
		fontSelectionDialogPrivateStruct = gi.StructNew("Gtk", "FontSelectionDialogPrivate")
	})
}

type FontSelectionDialogPrivate struct {
	native uintptr
}

var fontSelectionPrivateStruct *gi.Struct
var fontSelectionPrivateStruct_Once sync.Once

func fontSelectionPrivateStruct_Set() {
	fontSelectionPrivateStruct_Once.Do(func() {
		fontSelectionPrivateStruct = gi.StructNew("Gtk", "FontSelectionPrivate")
	})
}

type FontSelectionPrivate struct {
	native uintptr
}

var frameAccessibleClassStruct *gi.Struct
var frameAccessibleClassStruct_Once sync.Once

func frameAccessibleClassStruct_Set() {
	frameAccessibleClassStruct_Once.Do(func() {
		frameAccessibleClassStruct = gi.StructNew("Gtk", "FrameAccessibleClass")
	})
}

type FrameAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var frameAccessiblePrivateStruct *gi.Struct
var frameAccessiblePrivateStruct_Once sync.Once

func frameAccessiblePrivateStruct_Set() {
	frameAccessiblePrivateStruct_Once.Do(func() {
		frameAccessiblePrivateStruct = gi.StructNew("Gtk", "FrameAccessiblePrivate")
	})
}

type FrameAccessiblePrivate struct {
	native uintptr
}

var frameClassStruct *gi.Struct
var frameClassStruct_Once sync.Once

func frameClassStruct_Set() {
	frameClassStruct_Once.Do(func() {
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
var framePrivateStruct_Once sync.Once

func framePrivateStruct_Set() {
	framePrivateStruct_Once.Do(func() {
		framePrivateStruct = gi.StructNew("Gtk", "FramePrivate")
	})
}

type FramePrivate struct {
	native uintptr
}

var gLAreaClassStruct *gi.Struct
var gLAreaClassStruct_Once sync.Once

func gLAreaClassStruct_Set() {
	gLAreaClassStruct_Once.Do(func() {
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
var gestureClassStruct_Once sync.Once

func gestureClassStruct_Set() {
	gestureClassStruct_Once.Do(func() {
		gestureClassStruct = gi.StructNew("Gtk", "GestureClass")
	})
}

type GestureClass struct {
	native uintptr
}

var gestureDragClassStruct *gi.Struct
var gestureDragClassStruct_Once sync.Once

func gestureDragClassStruct_Set() {
	gestureDragClassStruct_Once.Do(func() {
		gestureDragClassStruct = gi.StructNew("Gtk", "GestureDragClass")
	})
}

type GestureDragClass struct {
	native uintptr
}

var gestureLongPressClassStruct *gi.Struct
var gestureLongPressClassStruct_Once sync.Once

func gestureLongPressClassStruct_Set() {
	gestureLongPressClassStruct_Once.Do(func() {
		gestureLongPressClassStruct = gi.StructNew("Gtk", "GestureLongPressClass")
	})
}

type GestureLongPressClass struct {
	native uintptr
}

var gestureMultiPressClassStruct *gi.Struct
var gestureMultiPressClassStruct_Once sync.Once

func gestureMultiPressClassStruct_Set() {
	gestureMultiPressClassStruct_Once.Do(func() {
		gestureMultiPressClassStruct = gi.StructNew("Gtk", "GestureMultiPressClass")
	})
}

type GestureMultiPressClass struct {
	native uintptr
}

var gesturePanClassStruct *gi.Struct
var gesturePanClassStruct_Once sync.Once

func gesturePanClassStruct_Set() {
	gesturePanClassStruct_Once.Do(func() {
		gesturePanClassStruct = gi.StructNew("Gtk", "GesturePanClass")
	})
}

type GesturePanClass struct {
	native uintptr
}

var gestureRotateClassStruct *gi.Struct
var gestureRotateClassStruct_Once sync.Once

func gestureRotateClassStruct_Set() {
	gestureRotateClassStruct_Once.Do(func() {
		gestureRotateClassStruct = gi.StructNew("Gtk", "GestureRotateClass")
	})
}

type GestureRotateClass struct {
	native uintptr
}

var gestureSingleClassStruct *gi.Struct
var gestureSingleClassStruct_Once sync.Once

func gestureSingleClassStruct_Set() {
	gestureSingleClassStruct_Once.Do(func() {
		gestureSingleClassStruct = gi.StructNew("Gtk", "GestureSingleClass")
	})
}

type GestureSingleClass struct {
	native uintptr
}

var gestureStylusClassStruct *gi.Struct
var gestureStylusClassStruct_Once sync.Once

func gestureStylusClassStruct_Set() {
	gestureStylusClassStruct_Once.Do(func() {
		gestureStylusClassStruct = gi.StructNew("Gtk", "GestureStylusClass")
	})
}

type GestureStylusClass struct {
	native uintptr
}

var gestureSwipeClassStruct *gi.Struct
var gestureSwipeClassStruct_Once sync.Once

func gestureSwipeClassStruct_Set() {
	gestureSwipeClassStruct_Once.Do(func() {
		gestureSwipeClassStruct = gi.StructNew("Gtk", "GestureSwipeClass")
	})
}

type GestureSwipeClass struct {
	native uintptr
}

var gestureZoomClassStruct *gi.Struct
var gestureZoomClassStruct_Once sync.Once

func gestureZoomClassStruct_Set() {
	gestureZoomClassStruct_Once.Do(func() {
		gestureZoomClassStruct = gi.StructNew("Gtk", "GestureZoomClass")
	})
}

type GestureZoomClass struct {
	native uintptr
}

var gradientStruct *gi.Struct
var gradientStruct_Once sync.Once

func gradientStruct_Set() {
	gradientStruct_Once.Do(func() {
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
		gradientStruct_Set()
		gradientRefFunction = gradientStruct.InvokerNew("ref")
	})
}

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
		gradientStruct_Set()
		gradientToStringFunction = gradientStruct.InvokerNew("to_string")
	})
}

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
		gradientStruct_Set()
		gradientUnrefFunction = gradientStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type gtk_gradient_unref.
func (recv *Gradient) Unref() {
	gradientUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	gradientUnrefFunction.Invoke(inArgs[:], nil)

}

var gridClassStruct *gi.Struct
var gridClassStruct_Once sync.Once

func gridClassStruct_Set() {
	gridClassStruct_Once.Do(func() {
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
var gridPrivateStruct_Once sync.Once

func gridPrivateStruct_Set() {
	gridPrivateStruct_Once.Do(func() {
		gridPrivateStruct = gi.StructNew("Gtk", "GridPrivate")
	})
}

type GridPrivate struct {
	native uintptr
}

var hBoxClassStruct *gi.Struct
var hBoxClassStruct_Once sync.Once

func hBoxClassStruct_Set() {
	hBoxClassStruct_Once.Do(func() {
		hBoxClassStruct = gi.StructNew("Gtk", "HBoxClass")
	})
}

type HBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
}

var hButtonBoxClassStruct *gi.Struct
var hButtonBoxClassStruct_Once sync.Once

func hButtonBoxClassStruct_Set() {
	hButtonBoxClassStruct_Once.Do(func() {
		hButtonBoxClassStruct = gi.StructNew("Gtk", "HButtonBoxClass")
	})
}

type HButtonBoxClass struct {
	native      uintptr
	ParentClass *ButtonBoxClass
}

var hPanedClassStruct *gi.Struct
var hPanedClassStruct_Once sync.Once

func hPanedClassStruct_Set() {
	hPanedClassStruct_Once.Do(func() {
		hPanedClassStruct = gi.StructNew("Gtk", "HPanedClass")
	})
}

type HPanedClass struct {
	native      uintptr
	ParentClass *PanedClass
}

var hSVClassStruct *gi.Struct
var hSVClassStruct_Once sync.Once

func hSVClassStruct_Set() {
	hSVClassStruct_Once.Do(func() {
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
var hSVPrivateStruct_Once sync.Once

func hSVPrivateStruct_Set() {
	hSVPrivateStruct_Once.Do(func() {
		hSVPrivateStruct = gi.StructNew("Gtk", "HSVPrivate")
	})
}

type HSVPrivate struct {
	native uintptr
}

var hScaleClassStruct *gi.Struct
var hScaleClassStruct_Once sync.Once

func hScaleClassStruct_Set() {
	hScaleClassStruct_Once.Do(func() {
		hScaleClassStruct = gi.StructNew("Gtk", "HScaleClass")
	})
}

type HScaleClass struct {
	native      uintptr
	ParentClass *ScaleClass
}

var hScrollbarClassStruct *gi.Struct
var hScrollbarClassStruct_Once sync.Once

func hScrollbarClassStruct_Set() {
	hScrollbarClassStruct_Once.Do(func() {
		hScrollbarClassStruct = gi.StructNew("Gtk", "HScrollbarClass")
	})
}

type HScrollbarClass struct {
	native      uintptr
	ParentClass *ScrollbarClass
}

var hSeparatorClassStruct *gi.Struct
var hSeparatorClassStruct_Once sync.Once

func hSeparatorClassStruct_Set() {
	hSeparatorClassStruct_Once.Do(func() {
		hSeparatorClassStruct = gi.StructNew("Gtk", "HSeparatorClass")
	})
}

type HSeparatorClass struct {
	native      uintptr
	ParentClass *SeparatorClass
}

var handleBoxClassStruct *gi.Struct
var handleBoxClassStruct_Once sync.Once

func handleBoxClassStruct_Set() {
	handleBoxClassStruct_Once.Do(func() {
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
var handleBoxPrivateStruct_Once sync.Once

func handleBoxPrivateStruct_Set() {
	handleBoxPrivateStruct_Once.Do(func() {
		handleBoxPrivateStruct = gi.StructNew("Gtk", "HandleBoxPrivate")
	})
}

type HandleBoxPrivate struct {
	native uintptr
}

var headerBarAccessibleClassStruct *gi.Struct
var headerBarAccessibleClassStruct_Once sync.Once

func headerBarAccessibleClassStruct_Set() {
	headerBarAccessibleClassStruct_Once.Do(func() {
		headerBarAccessibleClassStruct = gi.StructNew("Gtk", "HeaderBarAccessibleClass")
	})
}

type HeaderBarAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var headerBarAccessiblePrivateStruct *gi.Struct
var headerBarAccessiblePrivateStruct_Once sync.Once

func headerBarAccessiblePrivateStruct_Set() {
	headerBarAccessiblePrivateStruct_Once.Do(func() {
		headerBarAccessiblePrivateStruct = gi.StructNew("Gtk", "HeaderBarAccessiblePrivate")
	})
}

type HeaderBarAccessiblePrivate struct {
	native uintptr
}

var headerBarClassStruct *gi.Struct
var headerBarClassStruct_Once sync.Once

func headerBarClassStruct_Set() {
	headerBarClassStruct_Once.Do(func() {
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
var headerBarPrivateStruct_Once sync.Once

func headerBarPrivateStruct_Set() {
	headerBarPrivateStruct_Once.Do(func() {
		headerBarPrivateStruct = gi.StructNew("Gtk", "HeaderBarPrivate")
	})
}

type HeaderBarPrivate struct {
	native uintptr
}

var iMContextClassStruct *gi.Struct
var iMContextClassStruct_Once sync.Once

func iMContextClassStruct_Set() {
	iMContextClassStruct_Once.Do(func() {
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
var iMContextInfoStruct_Once sync.Once

func iMContextInfoStruct_Set() {
	iMContextInfoStruct_Once.Do(func() {
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
var iMContextSimpleClassStruct_Once sync.Once

func iMContextSimpleClassStruct_Set() {
	iMContextSimpleClassStruct_Once.Do(func() {
		iMContextSimpleClassStruct = gi.StructNew("Gtk", "IMContextSimpleClass")
	})
}

type IMContextSimpleClass struct {
	native      uintptr
	ParentClass *IMContextClass
}

var iMContextSimplePrivateStruct *gi.Struct
var iMContextSimplePrivateStruct_Once sync.Once

func iMContextSimplePrivateStruct_Set() {
	iMContextSimplePrivateStruct_Once.Do(func() {
		iMContextSimplePrivateStruct = gi.StructNew("Gtk", "IMContextSimplePrivate")
	})
}

type IMContextSimplePrivate struct {
	native uintptr
}

var iMMulticontextClassStruct *gi.Struct
var iMMulticontextClassStruct_Once sync.Once

func iMMulticontextClassStruct_Set() {
	iMMulticontextClassStruct_Once.Do(func() {
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
var iMMulticontextPrivateStruct_Once sync.Once

func iMMulticontextPrivateStruct_Set() {
	iMMulticontextPrivateStruct_Once.Do(func() {
		iMMulticontextPrivateStruct = gi.StructNew("Gtk", "IMMulticontextPrivate")
	})
}

type IMMulticontextPrivate struct {
	native uintptr
}

var iconFactoryClassStruct *gi.Struct
var iconFactoryClassStruct_Once sync.Once

func iconFactoryClassStruct_Set() {
	iconFactoryClassStruct_Once.Do(func() {
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
var iconFactoryPrivateStruct_Once sync.Once

func iconFactoryPrivateStruct_Set() {
	iconFactoryPrivateStruct_Once.Do(func() {
		iconFactoryPrivateStruct = gi.StructNew("Gtk", "IconFactoryPrivate")
	})
}

type IconFactoryPrivate struct {
	native uintptr
}

var iconInfoClassStruct *gi.Struct
var iconInfoClassStruct_Once sync.Once

func iconInfoClassStruct_Set() {
	iconInfoClassStruct_Once.Do(func() {
		iconInfoClassStruct = gi.StructNew("Gtk", "IconInfoClass")
	})
}

type IconInfoClass struct {
	native uintptr
}

var iconSetStruct *gi.Struct
var iconSetStruct_Once sync.Once

func iconSetStruct_Set() {
	iconSetStruct_Once.Do(func() {
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
		iconSetStruct_Set()
		iconSetNewFunction = iconSetStruct.InvokerNew("new")
	})
}

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
		iconSetStruct_Set()
		iconSetCopyFunction = iconSetStruct.InvokerNew("copy")
	})
}

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
		iconSetStruct_Set()
		iconSetRefFunction = iconSetStruct.InvokerNew("ref")
	})
}

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
		iconSetStruct_Set()
		iconSetUnrefFunction = iconSetStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type gtk_icon_set_unref.
func (recv *IconSet) Unref() {
	iconSetUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	iconSetUnrefFunction.Invoke(inArgs[:], nil)

}

var iconSourceStruct *gi.Struct
var iconSourceStruct_Once sync.Once

func iconSourceStruct_Set() {
	iconSourceStruct_Once.Do(func() {
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
		iconSourceStruct_Set()
		iconSourceNewFunction = iconSourceStruct.InvokerNew("new")
	})
}

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
		iconSourceStruct_Set()
		iconSourceCopyFunction = iconSourceStruct.InvokerNew("copy")
	})
}

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
		iconSourceStruct_Set()
		iconSourceFreeFunction = iconSourceStruct.InvokerNew("free")
	})
}

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
		iconSourceStruct_Set()
		iconSourceGetIconNameFunction = iconSourceStruct.InvokerNew("get_icon_name")
	})
}

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
		iconSourceStruct_Set()
		iconSourceSetIconNameFunction = iconSourceStruct.InvokerNew("set_icon_name")
	})
}

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
var iconThemeClassStruct_Once sync.Once

func iconThemeClassStruct_Set() {
	iconThemeClassStruct_Once.Do(func() {
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
var iconThemePrivateStruct_Once sync.Once

func iconThemePrivateStruct_Set() {
	iconThemePrivateStruct_Once.Do(func() {
		iconThemePrivateStruct = gi.StructNew("Gtk", "IconThemePrivate")
	})
}

type IconThemePrivate struct {
	native uintptr
}

var iconViewAccessibleClassStruct *gi.Struct
var iconViewAccessibleClassStruct_Once sync.Once

func iconViewAccessibleClassStruct_Set() {
	iconViewAccessibleClassStruct_Once.Do(func() {
		iconViewAccessibleClassStruct = gi.StructNew("Gtk", "IconViewAccessibleClass")
	})
}

type IconViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var iconViewAccessiblePrivateStruct *gi.Struct
var iconViewAccessiblePrivateStruct_Once sync.Once

func iconViewAccessiblePrivateStruct_Set() {
	iconViewAccessiblePrivateStruct_Once.Do(func() {
		iconViewAccessiblePrivateStruct = gi.StructNew("Gtk", "IconViewAccessiblePrivate")
	})
}

type IconViewAccessiblePrivate struct {
	native uintptr
}

var iconViewClassStruct *gi.Struct
var iconViewClassStruct_Once sync.Once

func iconViewClassStruct_Set() {
	iconViewClassStruct_Once.Do(func() {
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
var iconViewPrivateStruct_Once sync.Once

func iconViewPrivateStruct_Set() {
	iconViewPrivateStruct_Once.Do(func() {
		iconViewPrivateStruct = gi.StructNew("Gtk", "IconViewPrivate")
	})
}

type IconViewPrivate struct {
	native uintptr
}

var imageAccessibleClassStruct *gi.Struct
var imageAccessibleClassStruct_Once sync.Once

func imageAccessibleClassStruct_Set() {
	imageAccessibleClassStruct_Once.Do(func() {
		imageAccessibleClassStruct = gi.StructNew("Gtk", "ImageAccessibleClass")
	})
}

type ImageAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var imageAccessiblePrivateStruct *gi.Struct
var imageAccessiblePrivateStruct_Once sync.Once

func imageAccessiblePrivateStruct_Set() {
	imageAccessiblePrivateStruct_Once.Do(func() {
		imageAccessiblePrivateStruct = gi.StructNew("Gtk", "ImageAccessiblePrivate")
	})
}

type ImageAccessiblePrivate struct {
	native uintptr
}

var imageCellAccessibleClassStruct *gi.Struct
var imageCellAccessibleClassStruct_Once sync.Once

func imageCellAccessibleClassStruct_Set() {
	imageCellAccessibleClassStruct_Once.Do(func() {
		imageCellAccessibleClassStruct = gi.StructNew("Gtk", "ImageCellAccessibleClass")
	})
}

type ImageCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

var imageCellAccessiblePrivateStruct *gi.Struct
var imageCellAccessiblePrivateStruct_Once sync.Once

func imageCellAccessiblePrivateStruct_Set() {
	imageCellAccessiblePrivateStruct_Once.Do(func() {
		imageCellAccessiblePrivateStruct = gi.StructNew("Gtk", "ImageCellAccessiblePrivate")
	})
}

type ImageCellAccessiblePrivate struct {
	native uintptr
}

var imageClassStruct *gi.Struct
var imageClassStruct_Once sync.Once

func imageClassStruct_Set() {
	imageClassStruct_Once.Do(func() {
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
var imageMenuItemClassStruct_Once sync.Once

func imageMenuItemClassStruct_Set() {
	imageMenuItemClassStruct_Once.Do(func() {
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
var imageMenuItemPrivateStruct_Once sync.Once

func imageMenuItemPrivateStruct_Set() {
	imageMenuItemPrivateStruct_Once.Do(func() {
		imageMenuItemPrivateStruct = gi.StructNew("Gtk", "ImageMenuItemPrivate")
	})
}

type ImageMenuItemPrivate struct {
	native uintptr
}

var imagePrivateStruct *gi.Struct
var imagePrivateStruct_Once sync.Once

func imagePrivateStruct_Set() {
	imagePrivateStruct_Once.Do(func() {
		imagePrivateStruct = gi.StructNew("Gtk", "ImagePrivate")
	})
}

type ImagePrivate struct {
	native uintptr
}

var infoBarClassStruct *gi.Struct
var infoBarClassStruct_Once sync.Once

func infoBarClassStruct_Set() {
	infoBarClassStruct_Once.Do(func() {
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
var infoBarPrivateStruct_Once sync.Once

func infoBarPrivateStruct_Set() {
	infoBarPrivateStruct_Once.Do(func() {
		infoBarPrivateStruct = gi.StructNew("Gtk", "InfoBarPrivate")
	})
}

type InfoBarPrivate struct {
	native uintptr
}

var invisibleClassStruct *gi.Struct
var invisibleClassStruct_Once sync.Once

func invisibleClassStruct_Set() {
	invisibleClassStruct_Once.Do(func() {
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
var invisiblePrivateStruct_Once sync.Once

func invisiblePrivateStruct_Set() {
	invisiblePrivateStruct_Once.Do(func() {
		invisiblePrivateStruct = gi.StructNew("Gtk", "InvisiblePrivate")
	})
}

type InvisiblePrivate struct {
	native uintptr
}

var labelAccessibleClassStruct *gi.Struct
var labelAccessibleClassStruct_Once sync.Once

func labelAccessibleClassStruct_Set() {
	labelAccessibleClassStruct_Once.Do(func() {
		labelAccessibleClassStruct = gi.StructNew("Gtk", "LabelAccessibleClass")
	})
}

type LabelAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var labelAccessiblePrivateStruct *gi.Struct
var labelAccessiblePrivateStruct_Once sync.Once

func labelAccessiblePrivateStruct_Set() {
	labelAccessiblePrivateStruct_Once.Do(func() {
		labelAccessiblePrivateStruct = gi.StructNew("Gtk", "LabelAccessiblePrivate")
	})
}

type LabelAccessiblePrivate struct {
	native uintptr
}

var labelClassStruct *gi.Struct
var labelClassStruct_Once sync.Once

func labelClassStruct_Set() {
	labelClassStruct_Once.Do(func() {
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
var labelPrivateStruct_Once sync.Once

func labelPrivateStruct_Set() {
	labelPrivateStruct_Once.Do(func() {
		labelPrivateStruct = gi.StructNew("Gtk", "LabelPrivate")
	})
}

type LabelPrivate struct {
	native uintptr
}

var labelSelectionInfoStruct *gi.Struct
var labelSelectionInfoStruct_Once sync.Once

func labelSelectionInfoStruct_Set() {
	labelSelectionInfoStruct_Once.Do(func() {
		labelSelectionInfoStruct = gi.StructNew("Gtk", "LabelSelectionInfo")
	})
}

type LabelSelectionInfo struct {
	native uintptr
}

var layoutClassStruct *gi.Struct
var layoutClassStruct_Once sync.Once

func layoutClassStruct_Set() {
	layoutClassStruct_Once.Do(func() {
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
var layoutPrivateStruct_Once sync.Once

func layoutPrivateStruct_Set() {
	layoutPrivateStruct_Once.Do(func() {
		layoutPrivateStruct = gi.StructNew("Gtk", "LayoutPrivate")
	})
}

type LayoutPrivate struct {
	native uintptr
}

var levelBarAccessibleClassStruct *gi.Struct
var levelBarAccessibleClassStruct_Once sync.Once

func levelBarAccessibleClassStruct_Set() {
	levelBarAccessibleClassStruct_Once.Do(func() {
		levelBarAccessibleClassStruct = gi.StructNew("Gtk", "LevelBarAccessibleClass")
	})
}

type LevelBarAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var levelBarAccessiblePrivateStruct *gi.Struct
var levelBarAccessiblePrivateStruct_Once sync.Once

func levelBarAccessiblePrivateStruct_Set() {
	levelBarAccessiblePrivateStruct_Once.Do(func() {
		levelBarAccessiblePrivateStruct = gi.StructNew("Gtk", "LevelBarAccessiblePrivate")
	})
}

type LevelBarAccessiblePrivate struct {
	native uintptr
}

var levelBarClassStruct *gi.Struct
var levelBarClassStruct_Once sync.Once

func levelBarClassStruct_Set() {
	levelBarClassStruct_Once.Do(func() {
		levelBarClassStruct = gi.StructNew("Gtk", "LevelBarClass")
	})
}

type LevelBarClass struct {
	native uintptr
	// UNSUPPORTED : C value 'offset_changed' : missing Type
}

var levelBarPrivateStruct *gi.Struct
var levelBarPrivateStruct_Once sync.Once

func levelBarPrivateStruct_Set() {
	levelBarPrivateStruct_Once.Do(func() {
		levelBarPrivateStruct = gi.StructNew("Gtk", "LevelBarPrivate")
	})
}

type LevelBarPrivate struct {
	native uintptr
}

var linkButtonAccessibleClassStruct *gi.Struct
var linkButtonAccessibleClassStruct_Once sync.Once

func linkButtonAccessibleClassStruct_Set() {
	linkButtonAccessibleClassStruct_Once.Do(func() {
		linkButtonAccessibleClassStruct = gi.StructNew("Gtk", "LinkButtonAccessibleClass")
	})
}

type LinkButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var linkButtonAccessiblePrivateStruct *gi.Struct
var linkButtonAccessiblePrivateStruct_Once sync.Once

func linkButtonAccessiblePrivateStruct_Set() {
	linkButtonAccessiblePrivateStruct_Once.Do(func() {
		linkButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "LinkButtonAccessiblePrivate")
	})
}

type LinkButtonAccessiblePrivate struct {
	native uintptr
}

var linkButtonClassStruct *gi.Struct
var linkButtonClassStruct_Once sync.Once

func linkButtonClassStruct_Set() {
	linkButtonClassStruct_Once.Do(func() {
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
var linkButtonPrivateStruct_Once sync.Once

func linkButtonPrivateStruct_Set() {
	linkButtonPrivateStruct_Once.Do(func() {
		linkButtonPrivateStruct = gi.StructNew("Gtk", "LinkButtonPrivate")
	})
}

type LinkButtonPrivate struct {
	native uintptr
}

var listBoxAccessibleClassStruct *gi.Struct
var listBoxAccessibleClassStruct_Once sync.Once

func listBoxAccessibleClassStruct_Set() {
	listBoxAccessibleClassStruct_Once.Do(func() {
		listBoxAccessibleClassStruct = gi.StructNew("Gtk", "ListBoxAccessibleClass")
	})
}

type ListBoxAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var listBoxAccessiblePrivateStruct *gi.Struct
var listBoxAccessiblePrivateStruct_Once sync.Once

func listBoxAccessiblePrivateStruct_Set() {
	listBoxAccessiblePrivateStruct_Once.Do(func() {
		listBoxAccessiblePrivateStruct = gi.StructNew("Gtk", "ListBoxAccessiblePrivate")
	})
}

type ListBoxAccessiblePrivate struct {
	native uintptr
}

var listBoxClassStruct *gi.Struct
var listBoxClassStruct_Once sync.Once

func listBoxClassStruct_Set() {
	listBoxClassStruct_Once.Do(func() {
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
var listBoxRowAccessibleClassStruct_Once sync.Once

func listBoxRowAccessibleClassStruct_Set() {
	listBoxRowAccessibleClassStruct_Once.Do(func() {
		listBoxRowAccessibleClassStruct = gi.StructNew("Gtk", "ListBoxRowAccessibleClass")
	})
}

type ListBoxRowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var listBoxRowClassStruct *gi.Struct
var listBoxRowClassStruct_Once sync.Once

func listBoxRowClassStruct_Set() {
	listBoxRowClassStruct_Once.Do(func() {
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
var listStoreClassStruct_Once sync.Once

func listStoreClassStruct_Set() {
	listStoreClassStruct_Once.Do(func() {
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
var listStorePrivateStruct_Once sync.Once

func listStorePrivateStruct_Set() {
	listStorePrivateStruct_Once.Do(func() {
		listStorePrivateStruct = gi.StructNew("Gtk", "ListStorePrivate")
	})
}

type ListStorePrivate struct {
	native uintptr
}

var lockButtonAccessibleClassStruct *gi.Struct
var lockButtonAccessibleClassStruct_Once sync.Once

func lockButtonAccessibleClassStruct_Set() {
	lockButtonAccessibleClassStruct_Once.Do(func() {
		lockButtonAccessibleClassStruct = gi.StructNew("Gtk", "LockButtonAccessibleClass")
	})
}

type LockButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var lockButtonAccessiblePrivateStruct *gi.Struct
var lockButtonAccessiblePrivateStruct_Once sync.Once

func lockButtonAccessiblePrivateStruct_Set() {
	lockButtonAccessiblePrivateStruct_Once.Do(func() {
		lockButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "LockButtonAccessiblePrivate")
	})
}

type LockButtonAccessiblePrivate struct {
	native uintptr
}

var lockButtonClassStruct *gi.Struct
var lockButtonClassStruct_Once sync.Once

func lockButtonClassStruct_Set() {
	lockButtonClassStruct_Once.Do(func() {
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
var lockButtonPrivateStruct_Once sync.Once

func lockButtonPrivateStruct_Set() {
	lockButtonPrivateStruct_Once.Do(func() {
		lockButtonPrivateStruct = gi.StructNew("Gtk", "LockButtonPrivate")
	})
}

type LockButtonPrivate struct {
	native uintptr
}

var menuAccessibleClassStruct *gi.Struct
var menuAccessibleClassStruct_Once sync.Once

func menuAccessibleClassStruct_Set() {
	menuAccessibleClassStruct_Once.Do(func() {
		menuAccessibleClassStruct = gi.StructNew("Gtk", "MenuAccessibleClass")
	})
}

type MenuAccessibleClass struct {
	native      uintptr
	ParentClass *MenuShellAccessibleClass
}

var menuAccessiblePrivateStruct *gi.Struct
var menuAccessiblePrivateStruct_Once sync.Once

func menuAccessiblePrivateStruct_Set() {
	menuAccessiblePrivateStruct_Once.Do(func() {
		menuAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuAccessiblePrivate")
	})
}

type MenuAccessiblePrivate struct {
	native uintptr
}

var menuBarClassStruct *gi.Struct
var menuBarClassStruct_Once sync.Once

func menuBarClassStruct_Set() {
	menuBarClassStruct_Once.Do(func() {
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
var menuBarPrivateStruct_Once sync.Once

func menuBarPrivateStruct_Set() {
	menuBarPrivateStruct_Once.Do(func() {
		menuBarPrivateStruct = gi.StructNew("Gtk", "MenuBarPrivate")
	})
}

type MenuBarPrivate struct {
	native uintptr
}

var menuButtonAccessibleClassStruct *gi.Struct
var menuButtonAccessibleClassStruct_Once sync.Once

func menuButtonAccessibleClassStruct_Set() {
	menuButtonAccessibleClassStruct_Once.Do(func() {
		menuButtonAccessibleClassStruct = gi.StructNew("Gtk", "MenuButtonAccessibleClass")
	})
}

type MenuButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ToggleButtonAccessibleClass
}

var menuButtonAccessiblePrivateStruct *gi.Struct
var menuButtonAccessiblePrivateStruct_Once sync.Once

func menuButtonAccessiblePrivateStruct_Set() {
	menuButtonAccessiblePrivateStruct_Once.Do(func() {
		menuButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuButtonAccessiblePrivate")
	})
}

type MenuButtonAccessiblePrivate struct {
	native uintptr
}

var menuButtonClassStruct *gi.Struct
var menuButtonClassStruct_Once sync.Once

func menuButtonClassStruct_Set() {
	menuButtonClassStruct_Once.Do(func() {
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
var menuButtonPrivateStruct_Once sync.Once

func menuButtonPrivateStruct_Set() {
	menuButtonPrivateStruct_Once.Do(func() {
		menuButtonPrivateStruct = gi.StructNew("Gtk", "MenuButtonPrivate")
	})
}

type MenuButtonPrivate struct {
	native uintptr
}

var menuClassStruct *gi.Struct
var menuClassStruct_Once sync.Once

func menuClassStruct_Set() {
	menuClassStruct_Once.Do(func() {
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
var menuItemAccessibleClassStruct_Once sync.Once

func menuItemAccessibleClassStruct_Set() {
	menuItemAccessibleClassStruct_Once.Do(func() {
		menuItemAccessibleClassStruct = gi.StructNew("Gtk", "MenuItemAccessibleClass")
	})
}

type MenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var menuItemAccessiblePrivateStruct *gi.Struct
var menuItemAccessiblePrivateStruct_Once sync.Once

func menuItemAccessiblePrivateStruct_Set() {
	menuItemAccessiblePrivateStruct_Once.Do(func() {
		menuItemAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuItemAccessiblePrivate")
	})
}

type MenuItemAccessiblePrivate struct {
	native uintptr
}

var menuItemClassStruct *gi.Struct
var menuItemClassStruct_Once sync.Once

func menuItemClassStruct_Set() {
	menuItemClassStruct_Once.Do(func() {
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
var menuItemPrivateStruct_Once sync.Once

func menuItemPrivateStruct_Set() {
	menuItemPrivateStruct_Once.Do(func() {
		menuItemPrivateStruct = gi.StructNew("Gtk", "MenuItemPrivate")
	})
}

type MenuItemPrivate struct {
	native uintptr
}

var menuPrivateStruct *gi.Struct
var menuPrivateStruct_Once sync.Once

func menuPrivateStruct_Set() {
	menuPrivateStruct_Once.Do(func() {
		menuPrivateStruct = gi.StructNew("Gtk", "MenuPrivate")
	})
}

type MenuPrivate struct {
	native uintptr
}

var menuShellAccessibleClassStruct *gi.Struct
var menuShellAccessibleClassStruct_Once sync.Once

func menuShellAccessibleClassStruct_Set() {
	menuShellAccessibleClassStruct_Once.Do(func() {
		menuShellAccessibleClassStruct = gi.StructNew("Gtk", "MenuShellAccessibleClass")
	})
}

type MenuShellAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var menuShellAccessiblePrivateStruct *gi.Struct
var menuShellAccessiblePrivateStruct_Once sync.Once

func menuShellAccessiblePrivateStruct_Set() {
	menuShellAccessiblePrivateStruct_Once.Do(func() {
		menuShellAccessiblePrivateStruct = gi.StructNew("Gtk", "MenuShellAccessiblePrivate")
	})
}

type MenuShellAccessiblePrivate struct {
	native uintptr
}

var menuShellClassStruct *gi.Struct
var menuShellClassStruct_Once sync.Once

func menuShellClassStruct_Set() {
	menuShellClassStruct_Once.Do(func() {
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
var menuShellPrivateStruct_Once sync.Once

func menuShellPrivateStruct_Set() {
	menuShellPrivateStruct_Once.Do(func() {
		menuShellPrivateStruct = gi.StructNew("Gtk", "MenuShellPrivate")
	})
}

type MenuShellPrivate struct {
	native uintptr
}

var menuToolButtonClassStruct *gi.Struct
var menuToolButtonClassStruct_Once sync.Once

func menuToolButtonClassStruct_Set() {
	menuToolButtonClassStruct_Once.Do(func() {
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
var menuToolButtonPrivateStruct_Once sync.Once

func menuToolButtonPrivateStruct_Set() {
	menuToolButtonPrivateStruct_Once.Do(func() {
		menuToolButtonPrivateStruct = gi.StructNew("Gtk", "MenuToolButtonPrivate")
	})
}

type MenuToolButtonPrivate struct {
	native uintptr
}

var messageDialogClassStruct *gi.Struct
var messageDialogClassStruct_Once sync.Once

func messageDialogClassStruct_Set() {
	messageDialogClassStruct_Once.Do(func() {
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
var messageDialogPrivateStruct_Once sync.Once

func messageDialogPrivateStruct_Set() {
	messageDialogPrivateStruct_Once.Do(func() {
		messageDialogPrivateStruct = gi.StructNew("Gtk", "MessageDialogPrivate")
	})
}

type MessageDialogPrivate struct {
	native uintptr
}

var miscClassStruct *gi.Struct
var miscClassStruct_Once sync.Once

func miscClassStruct_Set() {
	miscClassStruct_Once.Do(func() {
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
var miscPrivateStruct_Once sync.Once

func miscPrivateStruct_Set() {
	miscPrivateStruct_Once.Do(func() {
		miscPrivateStruct = gi.StructNew("Gtk", "MiscPrivate")
	})
}

type MiscPrivate struct {
	native uintptr
}

var mountOperationClassStruct *gi.Struct
var mountOperationClassStruct_Once sync.Once

func mountOperationClassStruct_Set() {
	mountOperationClassStruct_Once.Do(func() {
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
var mountOperationPrivateStruct_Once sync.Once

func mountOperationPrivateStruct_Set() {
	mountOperationPrivateStruct_Once.Do(func() {
		mountOperationPrivateStruct = gi.StructNew("Gtk", "MountOperationPrivate")
	})
}

type MountOperationPrivate struct {
	native uintptr
}

var nativeDialogClassStruct *gi.Struct
var nativeDialogClassStruct_Once sync.Once

func nativeDialogClassStruct_Set() {
	nativeDialogClassStruct_Once.Do(func() {
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
var notebookAccessibleClassStruct_Once sync.Once

func notebookAccessibleClassStruct_Set() {
	notebookAccessibleClassStruct_Once.Do(func() {
		notebookAccessibleClassStruct = gi.StructNew("Gtk", "NotebookAccessibleClass")
	})
}

type NotebookAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var notebookAccessiblePrivateStruct *gi.Struct
var notebookAccessiblePrivateStruct_Once sync.Once

func notebookAccessiblePrivateStruct_Set() {
	notebookAccessiblePrivateStruct_Once.Do(func() {
		notebookAccessiblePrivateStruct = gi.StructNew("Gtk", "NotebookAccessiblePrivate")
	})
}

type NotebookAccessiblePrivate struct {
	native uintptr
}

var notebookClassStruct *gi.Struct
var notebookClassStruct_Once sync.Once

func notebookClassStruct_Set() {
	notebookClassStruct_Once.Do(func() {
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
var notebookPageAccessibleClassStruct_Once sync.Once

func notebookPageAccessibleClassStruct_Set() {
	notebookPageAccessibleClassStruct_Once.Do(func() {
		notebookPageAccessibleClassStruct = gi.StructNew("Gtk", "NotebookPageAccessibleClass")
	})
}

type NotebookPageAccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Atk.ObjectClass'
}

var notebookPageAccessiblePrivateStruct *gi.Struct
var notebookPageAccessiblePrivateStruct_Once sync.Once

func notebookPageAccessiblePrivateStruct_Set() {
	notebookPageAccessiblePrivateStruct_Once.Do(func() {
		notebookPageAccessiblePrivateStruct = gi.StructNew("Gtk", "NotebookPageAccessiblePrivate")
	})
}

type NotebookPageAccessiblePrivate struct {
	native uintptr
}

var notebookPrivateStruct *gi.Struct
var notebookPrivateStruct_Once sync.Once

func notebookPrivateStruct_Set() {
	notebookPrivateStruct_Once.Do(func() {
		notebookPrivateStruct = gi.StructNew("Gtk", "NotebookPrivate")
	})
}

type NotebookPrivate struct {
	native uintptr
}

var numerableIconClassStruct *gi.Struct
var numerableIconClassStruct_Once sync.Once

func numerableIconClassStruct_Set() {
	numerableIconClassStruct_Once.Do(func() {
		numerableIconClassStruct = gi.StructNew("Gtk", "NumerableIconClass")
	})
}

type NumerableIconClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Gio.EmblemedIconClass'
	// UNSUPPORTED : C value 'padding' : missing Type
}

var numerableIconPrivateStruct *gi.Struct
var numerableIconPrivateStruct_Once sync.Once

func numerableIconPrivateStruct_Set() {
	numerableIconPrivateStruct_Once.Do(func() {
		numerableIconPrivateStruct = gi.StructNew("Gtk", "NumerableIconPrivate")
	})
}

type NumerableIconPrivate struct {
	native uintptr
}

var offscreenWindowClassStruct *gi.Struct
var offscreenWindowClassStruct_Once sync.Once

func offscreenWindowClassStruct_Set() {
	offscreenWindowClassStruct_Once.Do(func() {
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
var orientableIfaceStruct_Once sync.Once

func orientableIfaceStruct_Set() {
	orientableIfaceStruct_Once.Do(func() {
		orientableIfaceStruct = gi.StructNew("Gtk", "OrientableIface")
	})
}

type OrientableIface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'
}

var overlayClassStruct *gi.Struct
var overlayClassStruct_Once sync.Once

func overlayClassStruct_Set() {
	overlayClassStruct_Once.Do(func() {
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
var overlayPrivateStruct_Once sync.Once

func overlayPrivateStruct_Set() {
	overlayPrivateStruct_Once.Do(func() {
		overlayPrivateStruct = gi.StructNew("Gtk", "OverlayPrivate")
	})
}

type OverlayPrivate struct {
	native uintptr
}

var padActionEntryStruct *gi.Struct
var padActionEntryStruct_Once sync.Once

func padActionEntryStruct_Set() {
	padActionEntryStruct_Once.Do(func() {
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
var padControllerClassStruct_Once sync.Once

func padControllerClassStruct_Set() {
	padControllerClassStruct_Once.Do(func() {
		padControllerClassStruct = gi.StructNew("Gtk", "PadControllerClass")
	})
}

type PadControllerClass struct {
	native uintptr
}

var pageRangeStruct *gi.Struct
var pageRangeStruct_Once sync.Once

func pageRangeStruct_Set() {
	pageRangeStruct_Once.Do(func() {
		pageRangeStruct = gi.StructNew("Gtk", "PageRange")
	})
}

type PageRange struct {
	native uintptr
	Start  int32
	End    int32
}

var panedAccessibleClassStruct *gi.Struct
var panedAccessibleClassStruct_Once sync.Once

func panedAccessibleClassStruct_Set() {
	panedAccessibleClassStruct_Once.Do(func() {
		panedAccessibleClassStruct = gi.StructNew("Gtk", "PanedAccessibleClass")
	})
}

type PanedAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var panedAccessiblePrivateStruct *gi.Struct
var panedAccessiblePrivateStruct_Once sync.Once

func panedAccessiblePrivateStruct_Set() {
	panedAccessiblePrivateStruct_Once.Do(func() {
		panedAccessiblePrivateStruct = gi.StructNew("Gtk", "PanedAccessiblePrivate")
	})
}

type PanedAccessiblePrivate struct {
	native uintptr
}

var panedClassStruct *gi.Struct
var panedClassStruct_Once sync.Once

func panedClassStruct_Set() {
	panedClassStruct_Once.Do(func() {
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
var panedPrivateStruct_Once sync.Once

func panedPrivateStruct_Set() {
	panedPrivateStruct_Once.Do(func() {
		panedPrivateStruct = gi.StructNew("Gtk", "PanedPrivate")
	})
}

type PanedPrivate struct {
	native uintptr
}

var paperSizeStruct *gi.Struct
var paperSizeStruct_Once sync.Once

func paperSizeStruct_Set() {
	paperSizeStruct_Once.Do(func() {
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
		paperSizeStruct_Set()
		paperSizeNewFunction = paperSizeStruct.InvokerNew("new")
	})
}

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
		paperSizeStruct_Set()
		paperSizeCopyFunction = paperSizeStruct.InvokerNew("copy")
	})
}

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
		paperSizeStruct_Set()
		paperSizeFreeFunction = paperSizeStruct.InvokerNew("free")
	})
}

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
		paperSizeStruct_Set()
		paperSizeGetDisplayNameFunction = paperSizeStruct.InvokerNew("get_display_name")
	})
}

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
		paperSizeStruct_Set()
		paperSizeGetNameFunction = paperSizeStruct.InvokerNew("get_name")
	})
}

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
		paperSizeStruct_Set()
		paperSizeGetPpdNameFunction = paperSizeStruct.InvokerNew("get_ppd_name")
	})
}

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
var placesSidebarClassStruct_Once sync.Once

func placesSidebarClassStruct_Set() {
	placesSidebarClassStruct_Once.Do(func() {
		placesSidebarClassStruct = gi.StructNew("Gtk", "PlacesSidebarClass")
	})
}

type PlacesSidebarClass struct {
	native uintptr
}

var plugClassStruct *gi.Struct
var plugClassStruct_Once sync.Once

func plugClassStruct_Set() {
	plugClassStruct_Once.Do(func() {
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
var plugPrivateStruct_Once sync.Once

func plugPrivateStruct_Set() {
	plugPrivateStruct_Once.Do(func() {
		plugPrivateStruct = gi.StructNew("Gtk", "PlugPrivate")
	})
}

type PlugPrivate struct {
	native uintptr
}

var popoverAccessibleClassStruct *gi.Struct
var popoverAccessibleClassStruct_Once sync.Once

func popoverAccessibleClassStruct_Set() {
	popoverAccessibleClassStruct_Once.Do(func() {
		popoverAccessibleClassStruct = gi.StructNew("Gtk", "PopoverAccessibleClass")
	})
}

type PopoverAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var popoverClassStruct *gi.Struct
var popoverClassStruct_Once sync.Once

func popoverClassStruct_Set() {
	popoverClassStruct_Once.Do(func() {
		popoverClassStruct = gi.StructNew("Gtk", "PopoverClass")
	})
}

type PopoverClass struct {
	native      uintptr
	ParentClass *BinClass
	// UNSUPPORTED : C value 'closed' : missing Type
}

var popoverMenuClassStruct *gi.Struct
var popoverMenuClassStruct_Once sync.Once

func popoverMenuClassStruct_Set() {
	popoverMenuClassStruct_Once.Do(func() {
		popoverMenuClassStruct = gi.StructNew("Gtk", "PopoverMenuClass")
	})
}

type PopoverMenuClass struct {
	native      uintptr
	ParentClass *PopoverClass
}

var popoverPrivateStruct *gi.Struct
var popoverPrivateStruct_Once sync.Once

func popoverPrivateStruct_Set() {
	popoverPrivateStruct_Once.Do(func() {
		popoverPrivateStruct = gi.StructNew("Gtk", "PopoverPrivate")
	})
}

type PopoverPrivate struct {
	native uintptr
}

var printOperationClassStruct *gi.Struct
var printOperationClassStruct_Once sync.Once

func printOperationClassStruct_Set() {
	printOperationClassStruct_Once.Do(func() {
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
var printOperationPreviewIfaceStruct_Once sync.Once

func printOperationPreviewIfaceStruct_Set() {
	printOperationPreviewIfaceStruct_Once.Do(func() {
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
var printOperationPrivateStruct_Once sync.Once

func printOperationPrivateStruct_Set() {
	printOperationPrivateStruct_Once.Do(func() {
		printOperationPrivateStruct = gi.StructNew("Gtk", "PrintOperationPrivate")
	})
}

type PrintOperationPrivate struct {
	native uintptr
}

var progressBarAccessibleClassStruct *gi.Struct
var progressBarAccessibleClassStruct_Once sync.Once

func progressBarAccessibleClassStruct_Set() {
	progressBarAccessibleClassStruct_Once.Do(func() {
		progressBarAccessibleClassStruct = gi.StructNew("Gtk", "ProgressBarAccessibleClass")
	})
}

type ProgressBarAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var progressBarAccessiblePrivateStruct *gi.Struct
var progressBarAccessiblePrivateStruct_Once sync.Once

func progressBarAccessiblePrivateStruct_Set() {
	progressBarAccessiblePrivateStruct_Once.Do(func() {
		progressBarAccessiblePrivateStruct = gi.StructNew("Gtk", "ProgressBarAccessiblePrivate")
	})
}

type ProgressBarAccessiblePrivate struct {
	native uintptr
}

var progressBarClassStruct *gi.Struct
var progressBarClassStruct_Once sync.Once

func progressBarClassStruct_Set() {
	progressBarClassStruct_Once.Do(func() {
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
var progressBarPrivateStruct_Once sync.Once

func progressBarPrivateStruct_Set() {
	progressBarPrivateStruct_Once.Do(func() {
		progressBarPrivateStruct = gi.StructNew("Gtk", "ProgressBarPrivate")
	})
}

type ProgressBarPrivate struct {
	native uintptr
}

var radioActionClassStruct *gi.Struct
var radioActionClassStruct_Once sync.Once

func radioActionClassStruct_Set() {
	radioActionClassStruct_Once.Do(func() {
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
var radioActionEntryStruct_Once sync.Once

func radioActionEntryStruct_Set() {
	radioActionEntryStruct_Once.Do(func() {
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
var radioActionPrivateStruct_Once sync.Once

func radioActionPrivateStruct_Set() {
	radioActionPrivateStruct_Once.Do(func() {
		radioActionPrivateStruct = gi.StructNew("Gtk", "RadioActionPrivate")
	})
}

type RadioActionPrivate struct {
	native uintptr
}

var radioButtonAccessibleClassStruct *gi.Struct
var radioButtonAccessibleClassStruct_Once sync.Once

func radioButtonAccessibleClassStruct_Set() {
	radioButtonAccessibleClassStruct_Once.Do(func() {
		radioButtonAccessibleClassStruct = gi.StructNew("Gtk", "RadioButtonAccessibleClass")
	})
}

type RadioButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ToggleButtonAccessibleClass
}

var radioButtonAccessiblePrivateStruct *gi.Struct
var radioButtonAccessiblePrivateStruct_Once sync.Once

func radioButtonAccessiblePrivateStruct_Set() {
	radioButtonAccessiblePrivateStruct_Once.Do(func() {
		radioButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "RadioButtonAccessiblePrivate")
	})
}

type RadioButtonAccessiblePrivate struct {
	native uintptr
}

var radioButtonClassStruct *gi.Struct
var radioButtonClassStruct_Once sync.Once

func radioButtonClassStruct_Set() {
	radioButtonClassStruct_Once.Do(func() {
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
var radioButtonPrivateStruct_Once sync.Once

func radioButtonPrivateStruct_Set() {
	radioButtonPrivateStruct_Once.Do(func() {
		radioButtonPrivateStruct = gi.StructNew("Gtk", "RadioButtonPrivate")
	})
}

type RadioButtonPrivate struct {
	native uintptr
}

var radioMenuItemAccessibleClassStruct *gi.Struct
var radioMenuItemAccessibleClassStruct_Once sync.Once

func radioMenuItemAccessibleClassStruct_Set() {
	radioMenuItemAccessibleClassStruct_Once.Do(func() {
		radioMenuItemAccessibleClassStruct = gi.StructNew("Gtk", "RadioMenuItemAccessibleClass")
	})
}

type RadioMenuItemAccessibleClass struct {
	native      uintptr
	ParentClass *CheckMenuItemAccessibleClass
}

var radioMenuItemAccessiblePrivateStruct *gi.Struct
var radioMenuItemAccessiblePrivateStruct_Once sync.Once

func radioMenuItemAccessiblePrivateStruct_Set() {
	radioMenuItemAccessiblePrivateStruct_Once.Do(func() {
		radioMenuItemAccessiblePrivateStruct = gi.StructNew("Gtk", "RadioMenuItemAccessiblePrivate")
	})
}

type RadioMenuItemAccessiblePrivate struct {
	native uintptr
}

var radioMenuItemClassStruct *gi.Struct
var radioMenuItemClassStruct_Once sync.Once

func radioMenuItemClassStruct_Set() {
	radioMenuItemClassStruct_Once.Do(func() {
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
var radioMenuItemPrivateStruct_Once sync.Once

func radioMenuItemPrivateStruct_Set() {
	radioMenuItemPrivateStruct_Once.Do(func() {
		radioMenuItemPrivateStruct = gi.StructNew("Gtk", "RadioMenuItemPrivate")
	})
}

type RadioMenuItemPrivate struct {
	native uintptr
}

var radioToolButtonClassStruct *gi.Struct
var radioToolButtonClassStruct_Once sync.Once

func radioToolButtonClassStruct_Set() {
	radioToolButtonClassStruct_Once.Do(func() {
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
var rangeAccessibleClassStruct_Once sync.Once

func rangeAccessibleClassStruct_Set() {
	rangeAccessibleClassStruct_Once.Do(func() {
		rangeAccessibleClassStruct = gi.StructNew("Gtk", "RangeAccessibleClass")
	})
}

type RangeAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var rangeAccessiblePrivateStruct *gi.Struct
var rangeAccessiblePrivateStruct_Once sync.Once

func rangeAccessiblePrivateStruct_Set() {
	rangeAccessiblePrivateStruct_Once.Do(func() {
		rangeAccessiblePrivateStruct = gi.StructNew("Gtk", "RangeAccessiblePrivate")
	})
}

type RangeAccessiblePrivate struct {
	native uintptr
}

var rangeClassStruct *gi.Struct
var rangeClassStruct_Once sync.Once

func rangeClassStruct_Set() {
	rangeClassStruct_Once.Do(func() {
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
var rangePrivateStruct_Once sync.Once

func rangePrivateStruct_Set() {
	rangePrivateStruct_Once.Do(func() {
		rangePrivateStruct = gi.StructNew("Gtk", "RangePrivate")
	})
}

type RangePrivate struct {
	native uintptr
}

var rcContextStruct *gi.Struct
var rcContextStruct_Once sync.Once

func rcContextStruct_Set() {
	rcContextStruct_Once.Do(func() {
		rcContextStruct = gi.StructNew("Gtk", "RcContext")
	})
}

type RcContext struct {
	native uintptr
}

var rcPropertyStruct *gi.Struct
var rcPropertyStruct_Once sync.Once

func rcPropertyStruct_Set() {
	rcPropertyStruct_Once.Do(func() {
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
var rcStyleClassStruct_Once sync.Once

func rcStyleClassStruct_Set() {
	rcStyleClassStruct_Once.Do(func() {
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
var recentActionClassStruct_Once sync.Once

func recentActionClassStruct_Set() {
	recentActionClassStruct_Once.Do(func() {
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
var recentActionPrivateStruct_Once sync.Once

func recentActionPrivateStruct_Set() {
	recentActionPrivateStruct_Once.Do(func() {
		recentActionPrivateStruct = gi.StructNew("Gtk", "RecentActionPrivate")
	})
}

type RecentActionPrivate struct {
	native uintptr
}

var recentChooserDialogClassStruct *gi.Struct
var recentChooserDialogClassStruct_Once sync.Once

func recentChooserDialogClassStruct_Set() {
	recentChooserDialogClassStruct_Once.Do(func() {
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
var recentChooserDialogPrivateStruct_Once sync.Once

func recentChooserDialogPrivateStruct_Set() {
	recentChooserDialogPrivateStruct_Once.Do(func() {
		recentChooserDialogPrivateStruct = gi.StructNew("Gtk", "RecentChooserDialogPrivate")
	})
}

type RecentChooserDialogPrivate struct {
	native uintptr
}

var recentChooserIfaceStruct *gi.Struct
var recentChooserIfaceStruct_Once sync.Once

func recentChooserIfaceStruct_Set() {
	recentChooserIfaceStruct_Once.Do(func() {
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
var recentChooserMenuClassStruct_Once sync.Once

func recentChooserMenuClassStruct_Set() {
	recentChooserMenuClassStruct_Once.Do(func() {
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
var recentChooserMenuPrivateStruct_Once sync.Once

func recentChooserMenuPrivateStruct_Set() {
	recentChooserMenuPrivateStruct_Once.Do(func() {
		recentChooserMenuPrivateStruct = gi.StructNew("Gtk", "RecentChooserMenuPrivate")
	})
}

type RecentChooserMenuPrivate struct {
	native uintptr
}

var recentChooserWidgetClassStruct *gi.Struct
var recentChooserWidgetClassStruct_Once sync.Once

func recentChooserWidgetClassStruct_Set() {
	recentChooserWidgetClassStruct_Once.Do(func() {
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
var recentChooserWidgetPrivateStruct_Once sync.Once

func recentChooserWidgetPrivateStruct_Set() {
	recentChooserWidgetPrivateStruct_Once.Do(func() {
		recentChooserWidgetPrivateStruct = gi.StructNew("Gtk", "RecentChooserWidgetPrivate")
	})
}

type RecentChooserWidgetPrivate struct {
	native uintptr
}

var recentDataStruct *gi.Struct
var recentDataStruct_Once sync.Once

func recentDataStruct_Set() {
	recentDataStruct_Once.Do(func() {
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
var recentFilterInfoStruct_Once sync.Once

func recentFilterInfoStruct_Set() {
	recentFilterInfoStruct_Once.Do(func() {
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
var recentInfoStruct_Once sync.Once

func recentInfoStruct_Set() {
	recentInfoStruct_Once.Do(func() {
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
		recentInfoStruct_Set()
		recentInfoGetAddedFunction = recentInfoStruct.InvokerNew("get_added")
	})
}

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
		recentInfoStruct_Set()
		recentInfoGetAgeFunction = recentInfoStruct.InvokerNew("get_age")
	})
}

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
		recentInfoStruct_Set()
		recentInfoGetDescriptionFunction = recentInfoStruct.InvokerNew("get_description")
	})
}

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
		recentInfoStruct_Set()
		recentInfoGetDisplayNameFunction = recentInfoStruct.InvokerNew("get_display_name")
	})
}

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
		recentInfoStruct_Set()
		recentInfoGetMimeTypeFunction = recentInfoStruct.InvokerNew("get_mime_type")
	})
}

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
		recentInfoStruct_Set()
		recentInfoGetModifiedFunction = recentInfoStruct.InvokerNew("get_modified")
	})
}

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
		recentInfoStruct_Set()
		recentInfoGetShortNameFunction = recentInfoStruct.InvokerNew("get_short_name")
	})
}

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
		recentInfoStruct_Set()
		recentInfoGetUriFunction = recentInfoStruct.InvokerNew("get_uri")
	})
}

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
		recentInfoStruct_Set()
		recentInfoGetUriDisplayFunction = recentInfoStruct.InvokerNew("get_uri_display")
	})
}

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
		recentInfoStruct_Set()
		recentInfoGetVisitedFunction = recentInfoStruct.InvokerNew("get_visited")
	})
}

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
		recentInfoStruct_Set()
		recentInfoLastApplicationFunction = recentInfoStruct.InvokerNew("last_application")
	})
}

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
		recentInfoStruct_Set()
		recentInfoRefFunction = recentInfoStruct.InvokerNew("ref")
	})
}

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
		recentInfoStruct_Set()
		recentInfoUnrefFunction = recentInfoStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type gtk_recent_info_unref.
func (recv *RecentInfo) Unref() {
	recentInfoUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	recentInfoUnrefFunction.Invoke(inArgs[:], nil)

}

var recentManagerClassStruct *gi.Struct
var recentManagerClassStruct_Once sync.Once

func recentManagerClassStruct_Set() {
	recentManagerClassStruct_Once.Do(func() {
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
var recentManagerPrivateStruct_Once sync.Once

func recentManagerPrivateStruct_Set() {
	recentManagerPrivateStruct_Once.Do(func() {
		recentManagerPrivateStruct = gi.StructNew("Gtk", "RecentManagerPrivate")
	})
}

type RecentManagerPrivate struct {
	native uintptr
}

var rendererCellAccessibleClassStruct *gi.Struct
var rendererCellAccessibleClassStruct_Once sync.Once

func rendererCellAccessibleClassStruct_Set() {
	rendererCellAccessibleClassStruct_Once.Do(func() {
		rendererCellAccessibleClassStruct = gi.StructNew("Gtk", "RendererCellAccessibleClass")
	})
}

type RendererCellAccessibleClass struct {
	native      uintptr
	ParentClass *CellAccessibleClass
}

var rendererCellAccessiblePrivateStruct *gi.Struct
var rendererCellAccessiblePrivateStruct_Once sync.Once

func rendererCellAccessiblePrivateStruct_Set() {
	rendererCellAccessiblePrivateStruct_Once.Do(func() {
		rendererCellAccessiblePrivateStruct = gi.StructNew("Gtk", "RendererCellAccessiblePrivate")
	})
}

type RendererCellAccessiblePrivate struct {
	native uintptr
}

var requestedSizeStruct *gi.Struct
var requestedSizeStruct_Once sync.Once

func requestedSizeStruct_Set() {
	requestedSizeStruct_Once.Do(func() {
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
var requisitionStruct_Once sync.Once

func requisitionStruct_Set() {
	requisitionStruct_Once.Do(func() {
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
		requisitionStruct_Set()
		requisitionNewFunction = requisitionStruct.InvokerNew("new")
	})
}

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
		requisitionStruct_Set()
		requisitionCopyFunction = requisitionStruct.InvokerNew("copy")
	})
}

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
		requisitionStruct_Set()
		requisitionFreeFunction = requisitionStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type gtk_requisition_free.
func (recv *Requisition) Free() {
	requisitionFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	requisitionFreeFunction.Invoke(inArgs[:], nil)

}

var revealerClassStruct *gi.Struct
var revealerClassStruct_Once sync.Once

func revealerClassStruct_Set() {
	revealerClassStruct_Once.Do(func() {
		revealerClassStruct = gi.StructNew("Gtk", "RevealerClass")
	})
}

type RevealerClass struct {
	native      uintptr
	ParentClass *BinClass
}

var scaleAccessibleClassStruct *gi.Struct
var scaleAccessibleClassStruct_Once sync.Once

func scaleAccessibleClassStruct_Set() {
	scaleAccessibleClassStruct_Once.Do(func() {
		scaleAccessibleClassStruct = gi.StructNew("Gtk", "ScaleAccessibleClass")
	})
}

type ScaleAccessibleClass struct {
	native      uintptr
	ParentClass *RangeAccessibleClass
}

var scaleAccessiblePrivateStruct *gi.Struct
var scaleAccessiblePrivateStruct_Once sync.Once

func scaleAccessiblePrivateStruct_Set() {
	scaleAccessiblePrivateStruct_Once.Do(func() {
		scaleAccessiblePrivateStruct = gi.StructNew("Gtk", "ScaleAccessiblePrivate")
	})
}

type ScaleAccessiblePrivate struct {
	native uintptr
}

var scaleButtonAccessibleClassStruct *gi.Struct
var scaleButtonAccessibleClassStruct_Once sync.Once

func scaleButtonAccessibleClassStruct_Set() {
	scaleButtonAccessibleClassStruct_Once.Do(func() {
		scaleButtonAccessibleClassStruct = gi.StructNew("Gtk", "ScaleButtonAccessibleClass")
	})
}

type ScaleButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var scaleButtonAccessiblePrivateStruct *gi.Struct
var scaleButtonAccessiblePrivateStruct_Once sync.Once

func scaleButtonAccessiblePrivateStruct_Set() {
	scaleButtonAccessiblePrivateStruct_Once.Do(func() {
		scaleButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "ScaleButtonAccessiblePrivate")
	})
}

type ScaleButtonAccessiblePrivate struct {
	native uintptr
}

var scaleButtonClassStruct *gi.Struct
var scaleButtonClassStruct_Once sync.Once

func scaleButtonClassStruct_Set() {
	scaleButtonClassStruct_Once.Do(func() {
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
var scaleButtonPrivateStruct_Once sync.Once

func scaleButtonPrivateStruct_Set() {
	scaleButtonPrivateStruct_Once.Do(func() {
		scaleButtonPrivateStruct = gi.StructNew("Gtk", "ScaleButtonPrivate")
	})
}

type ScaleButtonPrivate struct {
	native uintptr
}

var scaleClassStruct *gi.Struct
var scaleClassStruct_Once sync.Once

func scaleClassStruct_Set() {
	scaleClassStruct_Once.Do(func() {
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
var scalePrivateStruct_Once sync.Once

func scalePrivateStruct_Set() {
	scalePrivateStruct_Once.Do(func() {
		scalePrivateStruct = gi.StructNew("Gtk", "ScalePrivate")
	})
}

type ScalePrivate struct {
	native uintptr
}

var scrollableInterfaceStruct *gi.Struct
var scrollableInterfaceStruct_Once sync.Once

func scrollableInterfaceStruct_Set() {
	scrollableInterfaceStruct_Once.Do(func() {
		scrollableInterfaceStruct = gi.StructNew("Gtk", "ScrollableInterface")
	})
}

type ScrollableInterface struct {
	native uintptr
	// UNSUPPORTED : C value 'base_iface' : no Go type for 'GObject.TypeInterface'
	// UNSUPPORTED : C value 'get_border' : missing Type
}

var scrollbarClassStruct *gi.Struct
var scrollbarClassStruct_Once sync.Once

func scrollbarClassStruct_Set() {
	scrollbarClassStruct_Once.Do(func() {
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
var scrolledWindowAccessibleClassStruct_Once sync.Once

func scrolledWindowAccessibleClassStruct_Set() {
	scrolledWindowAccessibleClassStruct_Once.Do(func() {
		scrolledWindowAccessibleClassStruct = gi.StructNew("Gtk", "ScrolledWindowAccessibleClass")
	})
}

type ScrolledWindowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var scrolledWindowAccessiblePrivateStruct *gi.Struct
var scrolledWindowAccessiblePrivateStruct_Once sync.Once

func scrolledWindowAccessiblePrivateStruct_Set() {
	scrolledWindowAccessiblePrivateStruct_Once.Do(func() {
		scrolledWindowAccessiblePrivateStruct = gi.StructNew("Gtk", "ScrolledWindowAccessiblePrivate")
	})
}

type ScrolledWindowAccessiblePrivate struct {
	native uintptr
}

var scrolledWindowClassStruct *gi.Struct
var scrolledWindowClassStruct_Once sync.Once

func scrolledWindowClassStruct_Set() {
	scrolledWindowClassStruct_Once.Do(func() {
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
var scrolledWindowPrivateStruct_Once sync.Once

func scrolledWindowPrivateStruct_Set() {
	scrolledWindowPrivateStruct_Once.Do(func() {
		scrolledWindowPrivateStruct = gi.StructNew("Gtk", "ScrolledWindowPrivate")
	})
}

type ScrolledWindowPrivate struct {
	native uintptr
}

var searchBarClassStruct *gi.Struct
var searchBarClassStruct_Once sync.Once

func searchBarClassStruct_Set() {
	searchBarClassStruct_Once.Do(func() {
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
var searchEntryClassStruct_Once sync.Once

func searchEntryClassStruct_Set() {
	searchEntryClassStruct_Once.Do(func() {
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
var selectionDataStruct_Once sync.Once

func selectionDataStruct_Set() {
	selectionDataStruct_Once.Do(func() {
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
		selectionDataStruct_Set()
		selectionDataCopyFunction = selectionDataStruct.InvokerNew("copy")
	})
}

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
		selectionDataStruct_Set()
		selectionDataFreeFunction = selectionDataStruct.InvokerNew("free")
	})
}

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
		selectionDataStruct_Set()
		selectionDataGetDataFunction = selectionDataStruct.InvokerNew("get_data")
	})
}

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
		selectionDataStruct_Set()
		selectionDataGetDataWithLengthFunction = selectionDataStruct.InvokerNew("get_data_with_length")
	})
}

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
		selectionDataStruct_Set()
		selectionDataGetFormatFunction = selectionDataStruct.InvokerNew("get_format")
	})
}

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
		selectionDataStruct_Set()
		selectionDataGetLengthFunction = selectionDataStruct.InvokerNew("get_length")
	})
}

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
		selectionDataStruct_Set()
		selectionDataGetTextFunction = selectionDataStruct.InvokerNew("get_text")
	})
}

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
		selectionDataStruct_Set()
		selectionDataGetUrisFunction = selectionDataStruct.InvokerNew("get_uris")
	})
}

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
var separatorClassStruct_Once sync.Once

func separatorClassStruct_Set() {
	separatorClassStruct_Once.Do(func() {
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
var separatorMenuItemClassStruct_Once sync.Once

func separatorMenuItemClassStruct_Set() {
	separatorMenuItemClassStruct_Once.Do(func() {
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
var separatorPrivateStruct_Once sync.Once

func separatorPrivateStruct_Set() {
	separatorPrivateStruct_Once.Do(func() {
		separatorPrivateStruct = gi.StructNew("Gtk", "SeparatorPrivate")
	})
}

type SeparatorPrivate struct {
	native uintptr
}

var separatorToolItemClassStruct *gi.Struct
var separatorToolItemClassStruct_Once sync.Once

func separatorToolItemClassStruct_Set() {
	separatorToolItemClassStruct_Once.Do(func() {
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
var separatorToolItemPrivateStruct_Once sync.Once

func separatorToolItemPrivateStruct_Set() {
	separatorToolItemPrivateStruct_Once.Do(func() {
		separatorToolItemPrivateStruct = gi.StructNew("Gtk", "SeparatorToolItemPrivate")
	})
}

type SeparatorToolItemPrivate struct {
	native uintptr
}

var settingsClassStruct *gi.Struct
var settingsClassStruct_Once sync.Once

func settingsClassStruct_Set() {
	settingsClassStruct_Once.Do(func() {
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
var settingsPrivateStruct_Once sync.Once

func settingsPrivateStruct_Set() {
	settingsPrivateStruct_Once.Do(func() {
		settingsPrivateStruct = gi.StructNew("Gtk", "SettingsPrivate")
	})
}

type SettingsPrivate struct {
	native uintptr
}

var settingsValueStruct *gi.Struct
var settingsValueStruct_Once sync.Once

func settingsValueStruct_Set() {
	settingsValueStruct_Once.Do(func() {
		settingsValueStruct = gi.StructNew("Gtk", "SettingsValue")
	})
}

type SettingsValue struct {
	native uintptr
	Origin string
	// UNSUPPORTED : C value 'value' : no Go type for 'GObject.Value'
}

var shortcutLabelClassStruct *gi.Struct
var shortcutLabelClassStruct_Once sync.Once

func shortcutLabelClassStruct_Set() {
	shortcutLabelClassStruct_Once.Do(func() {
		shortcutLabelClassStruct = gi.StructNew("Gtk", "ShortcutLabelClass")
	})
}

type ShortcutLabelClass struct {
	native uintptr
}

var shortcutsGroupClassStruct *gi.Struct
var shortcutsGroupClassStruct_Once sync.Once

func shortcutsGroupClassStruct_Set() {
	shortcutsGroupClassStruct_Once.Do(func() {
		shortcutsGroupClassStruct = gi.StructNew("Gtk", "ShortcutsGroupClass")
	})
}

type ShortcutsGroupClass struct {
	native uintptr
}

var shortcutsSectionClassStruct *gi.Struct
var shortcutsSectionClassStruct_Once sync.Once

func shortcutsSectionClassStruct_Set() {
	shortcutsSectionClassStruct_Once.Do(func() {
		shortcutsSectionClassStruct = gi.StructNew("Gtk", "ShortcutsSectionClass")
	})
}

type ShortcutsSectionClass struct {
	native uintptr
}

var shortcutsShortcutClassStruct *gi.Struct
var shortcutsShortcutClassStruct_Once sync.Once

func shortcutsShortcutClassStruct_Set() {
	shortcutsShortcutClassStruct_Once.Do(func() {
		shortcutsShortcutClassStruct = gi.StructNew("Gtk", "ShortcutsShortcutClass")
	})
}

type ShortcutsShortcutClass struct {
	native uintptr
}

var shortcutsWindowClassStruct *gi.Struct
var shortcutsWindowClassStruct_Once sync.Once

func shortcutsWindowClassStruct_Set() {
	shortcutsWindowClassStruct_Once.Do(func() {
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
var sizeGroupClassStruct_Once sync.Once

func sizeGroupClassStruct_Set() {
	sizeGroupClassStruct_Once.Do(func() {
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
var sizeGroupPrivateStruct_Once sync.Once

func sizeGroupPrivateStruct_Set() {
	sizeGroupPrivateStruct_Once.Do(func() {
		sizeGroupPrivateStruct = gi.StructNew("Gtk", "SizeGroupPrivate")
	})
}

type SizeGroupPrivate struct {
	native uintptr
}

var socketClassStruct *gi.Struct
var socketClassStruct_Once sync.Once

func socketClassStruct_Set() {
	socketClassStruct_Once.Do(func() {
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
var socketPrivateStruct_Once sync.Once

func socketPrivateStruct_Set() {
	socketPrivateStruct_Once.Do(func() {
		socketPrivateStruct = gi.StructNew("Gtk", "SocketPrivate")
	})
}

type SocketPrivate struct {
	native uintptr
}

var spinButtonAccessibleClassStruct *gi.Struct
var spinButtonAccessibleClassStruct_Once sync.Once

func spinButtonAccessibleClassStruct_Set() {
	spinButtonAccessibleClassStruct_Once.Do(func() {
		spinButtonAccessibleClassStruct = gi.StructNew("Gtk", "SpinButtonAccessibleClass")
	})
}

type SpinButtonAccessibleClass struct {
	native      uintptr
	ParentClass *EntryAccessibleClass
}

var spinButtonAccessiblePrivateStruct *gi.Struct
var spinButtonAccessiblePrivateStruct_Once sync.Once

func spinButtonAccessiblePrivateStruct_Set() {
	spinButtonAccessiblePrivateStruct_Once.Do(func() {
		spinButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "SpinButtonAccessiblePrivate")
	})
}

type SpinButtonAccessiblePrivate struct {
	native uintptr
}

var spinButtonClassStruct *gi.Struct
var spinButtonClassStruct_Once sync.Once

func spinButtonClassStruct_Set() {
	spinButtonClassStruct_Once.Do(func() {
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
var spinButtonPrivateStruct_Once sync.Once

func spinButtonPrivateStruct_Set() {
	spinButtonPrivateStruct_Once.Do(func() {
		spinButtonPrivateStruct = gi.StructNew("Gtk", "SpinButtonPrivate")
	})
}

type SpinButtonPrivate struct {
	native uintptr
}

var spinnerAccessibleClassStruct *gi.Struct
var spinnerAccessibleClassStruct_Once sync.Once

func spinnerAccessibleClassStruct_Set() {
	spinnerAccessibleClassStruct_Once.Do(func() {
		spinnerAccessibleClassStruct = gi.StructNew("Gtk", "SpinnerAccessibleClass")
	})
}

type SpinnerAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var spinnerAccessiblePrivateStruct *gi.Struct
var spinnerAccessiblePrivateStruct_Once sync.Once

func spinnerAccessiblePrivateStruct_Set() {
	spinnerAccessiblePrivateStruct_Once.Do(func() {
		spinnerAccessiblePrivateStruct = gi.StructNew("Gtk", "SpinnerAccessiblePrivate")
	})
}

type SpinnerAccessiblePrivate struct {
	native uintptr
}

var spinnerClassStruct *gi.Struct
var spinnerClassStruct_Once sync.Once

func spinnerClassStruct_Set() {
	spinnerClassStruct_Once.Do(func() {
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
var spinnerPrivateStruct_Once sync.Once

func spinnerPrivateStruct_Set() {
	spinnerPrivateStruct_Once.Do(func() {
		spinnerPrivateStruct = gi.StructNew("Gtk", "SpinnerPrivate")
	})
}

type SpinnerPrivate struct {
	native uintptr
}

var stackAccessibleClassStruct *gi.Struct
var stackAccessibleClassStruct_Once sync.Once

func stackAccessibleClassStruct_Set() {
	stackAccessibleClassStruct_Once.Do(func() {
		stackAccessibleClassStruct = gi.StructNew("Gtk", "StackAccessibleClass")
	})
}

type StackAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var stackClassStruct *gi.Struct
var stackClassStruct_Once sync.Once

func stackClassStruct_Set() {
	stackClassStruct_Once.Do(func() {
		stackClassStruct = gi.StructNew("Gtk", "StackClass")
	})
}

type StackClass struct {
	native      uintptr
	ParentClass *ContainerClass
}

var stackSidebarClassStruct *gi.Struct
var stackSidebarClassStruct_Once sync.Once

func stackSidebarClassStruct_Set() {
	stackSidebarClassStruct_Once.Do(func() {
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
var stackSidebarPrivateStruct_Once sync.Once

func stackSidebarPrivateStruct_Set() {
	stackSidebarPrivateStruct_Once.Do(func() {
		stackSidebarPrivateStruct = gi.StructNew("Gtk", "StackSidebarPrivate")
	})
}

type StackSidebarPrivate struct {
	native uintptr
}

var stackSwitcherClassStruct *gi.Struct
var stackSwitcherClassStruct_Once sync.Once

func stackSwitcherClassStruct_Set() {
	stackSwitcherClassStruct_Once.Do(func() {
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
var statusIconClassStruct_Once sync.Once

func statusIconClassStruct_Set() {
	statusIconClassStruct_Once.Do(func() {
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
var statusIconPrivateStruct_Once sync.Once

func statusIconPrivateStruct_Set() {
	statusIconPrivateStruct_Once.Do(func() {
		statusIconPrivateStruct = gi.StructNew("Gtk", "StatusIconPrivate")
	})
}

type StatusIconPrivate struct {
	native uintptr
}

var statusbarAccessibleClassStruct *gi.Struct
var statusbarAccessibleClassStruct_Once sync.Once

func statusbarAccessibleClassStruct_Set() {
	statusbarAccessibleClassStruct_Once.Do(func() {
		statusbarAccessibleClassStruct = gi.StructNew("Gtk", "StatusbarAccessibleClass")
	})
}

type StatusbarAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var statusbarAccessiblePrivateStruct *gi.Struct
var statusbarAccessiblePrivateStruct_Once sync.Once

func statusbarAccessiblePrivateStruct_Set() {
	statusbarAccessiblePrivateStruct_Once.Do(func() {
		statusbarAccessiblePrivateStruct = gi.StructNew("Gtk", "StatusbarAccessiblePrivate")
	})
}

type StatusbarAccessiblePrivate struct {
	native uintptr
}

var statusbarClassStruct *gi.Struct
var statusbarClassStruct_Once sync.Once

func statusbarClassStruct_Set() {
	statusbarClassStruct_Once.Do(func() {
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
var statusbarPrivateStruct_Once sync.Once

func statusbarPrivateStruct_Set() {
	statusbarPrivateStruct_Once.Do(func() {
		statusbarPrivateStruct = gi.StructNew("Gtk", "StatusbarPrivate")
	})
}

type StatusbarPrivate struct {
	native uintptr
}

var stockItemStruct *gi.Struct
var stockItemStruct_Once sync.Once

func stockItemStruct_Set() {
	stockItemStruct_Once.Do(func() {
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
		stockItemStruct_Set()
		stockItemCopyFunction = stockItemStruct.InvokerNew("copy")
	})
}

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
		stockItemStruct_Set()
		stockItemFreeFunction = stockItemStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type gtk_stock_item_free.
func (recv *StockItem) Free() {
	stockItemFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	stockItemFreeFunction.Invoke(inArgs[:], nil)

}

var styleClassStruct *gi.Struct
var styleClassStruct_Once sync.Once

func styleClassStruct_Set() {
	styleClassStruct_Once.Do(func() {
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
var styleContextClassStruct_Once sync.Once

func styleContextClassStruct_Set() {
	styleContextClassStruct_Once.Do(func() {
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
var styleContextPrivateStruct_Once sync.Once

func styleContextPrivateStruct_Set() {
	styleContextPrivateStruct_Once.Do(func() {
		styleContextPrivateStruct = gi.StructNew("Gtk", "StyleContextPrivate")
	})
}

type StyleContextPrivate struct {
	native uintptr
}

var stylePropertiesClassStruct *gi.Struct
var stylePropertiesClassStruct_Once sync.Once

func stylePropertiesClassStruct_Set() {
	stylePropertiesClassStruct_Once.Do(func() {
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
var stylePropertiesPrivateStruct_Once sync.Once

func stylePropertiesPrivateStruct_Set() {
	stylePropertiesPrivateStruct_Once.Do(func() {
		stylePropertiesPrivateStruct = gi.StructNew("Gtk", "StylePropertiesPrivate")
	})
}

type StylePropertiesPrivate struct {
	native uintptr
}

var styleProviderIfaceStruct *gi.Struct
var styleProviderIfaceStruct_Once sync.Once

func styleProviderIfaceStruct_Set() {
	styleProviderIfaceStruct_Once.Do(func() {
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
var switchAccessibleClassStruct_Once sync.Once

func switchAccessibleClassStruct_Set() {
	switchAccessibleClassStruct_Once.Do(func() {
		switchAccessibleClassStruct = gi.StructNew("Gtk", "SwitchAccessibleClass")
	})
}

type SwitchAccessibleClass struct {
	native      uintptr
	ParentClass *WidgetAccessibleClass
}

var switchAccessiblePrivateStruct *gi.Struct
var switchAccessiblePrivateStruct_Once sync.Once

func switchAccessiblePrivateStruct_Set() {
	switchAccessiblePrivateStruct_Once.Do(func() {
		switchAccessiblePrivateStruct = gi.StructNew("Gtk", "SwitchAccessiblePrivate")
	})
}

type SwitchAccessiblePrivate struct {
	native uintptr
}

var switchClassStruct *gi.Struct
var switchClassStruct_Once sync.Once

func switchClassStruct_Set() {
	switchClassStruct_Once.Do(func() {
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
var switchPrivateStruct_Once sync.Once

func switchPrivateStruct_Set() {
	switchPrivateStruct_Once.Do(func() {
		switchPrivateStruct = gi.StructNew("Gtk", "SwitchPrivate")
	})
}

type SwitchPrivate struct {
	native uintptr
}

var symbolicColorStruct *gi.Struct
var symbolicColorStruct_Once sync.Once

func symbolicColorStruct_Set() {
	symbolicColorStruct_Once.Do(func() {
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
		symbolicColorStruct_Set()
		symbolicColorNewNameFunction = symbolicColorStruct.InvokerNew("new_name")
	})
}

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
		symbolicColorStruct_Set()
		symbolicColorNewWin32Function = symbolicColorStruct.InvokerNew("new_win32")
	})
}

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
		symbolicColorStruct_Set()
		symbolicColorRefFunction = symbolicColorStruct.InvokerNew("ref")
	})
}

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
		symbolicColorStruct_Set()
		symbolicColorToStringFunction = symbolicColorStruct.InvokerNew("to_string")
	})
}

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
		symbolicColorStruct_Set()
		symbolicColorUnrefFunction = symbolicColorStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type gtk_symbolic_color_unref.
func (recv *SymbolicColor) Unref() {
	symbolicColorUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	symbolicColorUnrefFunction.Invoke(inArgs[:], nil)

}

var tableChildStruct *gi.Struct
var tableChildStruct_Once sync.Once

func tableChildStruct_Set() {
	tableChildStruct_Once.Do(func() {
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
var tableClassStruct_Once sync.Once

func tableClassStruct_Set() {
	tableClassStruct_Once.Do(func() {
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
var tablePrivateStruct_Once sync.Once

func tablePrivateStruct_Set() {
	tablePrivateStruct_Once.Do(func() {
		tablePrivateStruct = gi.StructNew("Gtk", "TablePrivate")
	})
}

type TablePrivate struct {
	native uintptr
}

var tableRowColStruct *gi.Struct
var tableRowColStruct_Once sync.Once

func tableRowColStruct_Set() {
	tableRowColStruct_Once.Do(func() {
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
var targetEntryStruct_Once sync.Once

func targetEntryStruct_Set() {
	targetEntryStruct_Once.Do(func() {
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
		targetEntryStruct_Set()
		targetEntryNewFunction = targetEntryStruct.InvokerNew("new")
	})
}

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
		targetEntryStruct_Set()
		targetEntryCopyFunction = targetEntryStruct.InvokerNew("copy")
	})
}

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
		targetEntryStruct_Set()
		targetEntryFreeFunction = targetEntryStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type gtk_target_entry_free.
func (recv *TargetEntry) Free() {
	targetEntryFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	targetEntryFreeFunction.Invoke(inArgs[:], nil)

}

var targetListStruct *gi.Struct
var targetListStruct_Once sync.Once

func targetListStruct_Set() {
	targetListStruct_Once.Do(func() {
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
		targetListStruct_Set()
		targetListAddTextTargetsFunction = targetListStruct.InvokerNew("add_text_targets")
	})
}

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
		targetListStruct_Set()
		targetListAddUriTargetsFunction = targetListStruct.InvokerNew("add_uri_targets")
	})
}

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
		targetListStruct_Set()
		targetListRefFunction = targetListStruct.InvokerNew("ref")
	})
}

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
		targetListStruct_Set()
		targetListUnrefFunction = targetListStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type gtk_target_list_unref.
func (recv *TargetList) Unref() {
	targetListUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	targetListUnrefFunction.Invoke(inArgs[:], nil)

}

var targetPairStruct *gi.Struct
var targetPairStruct_Once sync.Once

func targetPairStruct_Set() {
	targetPairStruct_Once.Do(func() {
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
var tearoffMenuItemClassStruct_Once sync.Once

func tearoffMenuItemClassStruct_Set() {
	tearoffMenuItemClassStruct_Once.Do(func() {
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
var tearoffMenuItemPrivateStruct_Once sync.Once

func tearoffMenuItemPrivateStruct_Set() {
	tearoffMenuItemPrivateStruct_Once.Do(func() {
		tearoffMenuItemPrivateStruct = gi.StructNew("Gtk", "TearoffMenuItemPrivate")
	})
}

type TearoffMenuItemPrivate struct {
	native uintptr
}

var textAppearanceStruct *gi.Struct
var textAppearanceStruct_Once sync.Once

func textAppearanceStruct_Set() {
	textAppearanceStruct_Once.Do(func() {
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
var textAttributesStruct_Once sync.Once

func textAttributesStruct_Set() {
	textAttributesStruct_Once.Do(func() {
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
		textAttributesStruct_Set()
		textAttributesNewFunction = textAttributesStruct.InvokerNew("new")
	})
}

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
		textAttributesStruct_Set()
		textAttributesCopyFunction = textAttributesStruct.InvokerNew("copy")
	})
}

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
		textAttributesStruct_Set()
		textAttributesRefFunction = textAttributesStruct.InvokerNew("ref")
	})
}

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
		textAttributesStruct_Set()
		textAttributesUnrefFunction = textAttributesStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type gtk_text_attributes_unref.
func (recv *TextAttributes) Unref() {
	textAttributesUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	textAttributesUnrefFunction.Invoke(inArgs[:], nil)

}

var textBTreeStruct *gi.Struct
var textBTreeStruct_Once sync.Once

func textBTreeStruct_Set() {
	textBTreeStruct_Once.Do(func() {
		textBTreeStruct = gi.StructNew("Gtk", "TextBTree")
	})
}

type TextBTree struct {
	native uintptr
}

var textBufferClassStruct *gi.Struct
var textBufferClassStruct_Once sync.Once

func textBufferClassStruct_Set() {
	textBufferClassStruct_Once.Do(func() {
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
var textBufferPrivateStruct_Once sync.Once

func textBufferPrivateStruct_Set() {
	textBufferPrivateStruct_Once.Do(func() {
		textBufferPrivateStruct = gi.StructNew("Gtk", "TextBufferPrivate")
	})
}

type TextBufferPrivate struct {
	native uintptr
}

var textCellAccessibleClassStruct *gi.Struct
var textCellAccessibleClassStruct_Once sync.Once

func textCellAccessibleClassStruct_Set() {
	textCellAccessibleClassStruct_Once.Do(func() {
		textCellAccessibleClassStruct = gi.StructNew("Gtk", "TextCellAccessibleClass")
	})
}

type TextCellAccessibleClass struct {
	native      uintptr
	ParentClass *RendererCellAccessibleClass
}

var textCellAccessiblePrivateStruct *gi.Struct
var textCellAccessiblePrivateStruct_Once sync.Once

func textCellAccessiblePrivateStruct_Set() {
	textCellAccessiblePrivateStruct_Once.Do(func() {
		textCellAccessiblePrivateStruct = gi.StructNew("Gtk", "TextCellAccessiblePrivate")
	})
}

type TextCellAccessiblePrivate struct {
	native uintptr
}

var textChildAnchorClassStruct *gi.Struct
var textChildAnchorClassStruct_Once sync.Once

func textChildAnchorClassStruct_Set() {
	textChildAnchorClassStruct_Once.Do(func() {
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
var textIterStruct_Once sync.Once

func textIterStruct_Set() {
	textIterStruct_Once.Do(func() {
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
		textIterStruct_Set()
		textIterCopyFunction = textIterStruct.InvokerNew("copy")
	})
}

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
		textIterStruct_Set()
		textIterForwardToEndFunction = textIterStruct.InvokerNew("forward_to_end")
	})
}

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
		textIterStruct_Set()
		textIterFreeFunction = textIterStruct.InvokerNew("free")
	})
}

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
		textIterStruct_Set()
		textIterGetBytesInLineFunction = textIterStruct.InvokerNew("get_bytes_in_line")
	})
}

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
		textIterStruct_Set()
		textIterGetCharsInLineFunction = textIterStruct.InvokerNew("get_chars_in_line")
	})
}

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
		textIterStruct_Set()
		textIterGetLineFunction = textIterStruct.InvokerNew("get_line")
	})
}

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
		textIterStruct_Set()
		textIterGetLineIndexFunction = textIterStruct.InvokerNew("get_line_index")
	})
}

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
		textIterStruct_Set()
		textIterGetLineOffsetFunction = textIterStruct.InvokerNew("get_line_offset")
	})
}

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
		textIterStruct_Set()
		textIterGetOffsetFunction = textIterStruct.InvokerNew("get_offset")
	})
}

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
		textIterStruct_Set()
		textIterGetVisibleLineIndexFunction = textIterStruct.InvokerNew("get_visible_line_index")
	})
}

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
		textIterStruct_Set()
		textIterGetVisibleLineOffsetFunction = textIterStruct.InvokerNew("get_visible_line_offset")
	})
}

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
		textIterStruct_Set()
		textIterSetLineFunction = textIterStruct.InvokerNew("set_line")
	})
}

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
		textIterStruct_Set()
		textIterSetLineIndexFunction = textIterStruct.InvokerNew("set_line_index")
	})
}

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
		textIterStruct_Set()
		textIterSetLineOffsetFunction = textIterStruct.InvokerNew("set_line_offset")
	})
}

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
		textIterStruct_Set()
		textIterSetOffsetFunction = textIterStruct.InvokerNew("set_offset")
	})
}

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
		textIterStruct_Set()
		textIterSetVisibleLineIndexFunction = textIterStruct.InvokerNew("set_visible_line_index")
	})
}

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
		textIterStruct_Set()
		textIterSetVisibleLineOffsetFunction = textIterStruct.InvokerNew("set_visible_line_offset")
	})
}

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
var textMarkClassStruct_Once sync.Once

func textMarkClassStruct_Set() {
	textMarkClassStruct_Once.Do(func() {
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
var textTagClassStruct_Once sync.Once

func textTagClassStruct_Set() {
	textTagClassStruct_Once.Do(func() {
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
var textTagPrivateStruct_Once sync.Once

func textTagPrivateStruct_Set() {
	textTagPrivateStruct_Once.Do(func() {
		textTagPrivateStruct = gi.StructNew("Gtk", "TextTagPrivate")
	})
}

type TextTagPrivate struct {
	native uintptr
}

var textTagTableClassStruct *gi.Struct
var textTagTableClassStruct_Once sync.Once

func textTagTableClassStruct_Set() {
	textTagTableClassStruct_Once.Do(func() {
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
var textTagTablePrivateStruct_Once sync.Once

func textTagTablePrivateStruct_Set() {
	textTagTablePrivateStruct_Once.Do(func() {
		textTagTablePrivateStruct = gi.StructNew("Gtk", "TextTagTablePrivate")
	})
}

type TextTagTablePrivate struct {
	native uintptr
}

var textViewAccessibleClassStruct *gi.Struct
var textViewAccessibleClassStruct_Once sync.Once

func textViewAccessibleClassStruct_Set() {
	textViewAccessibleClassStruct_Once.Do(func() {
		textViewAccessibleClassStruct = gi.StructNew("Gtk", "TextViewAccessibleClass")
	})
}

type TextViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var textViewAccessiblePrivateStruct *gi.Struct
var textViewAccessiblePrivateStruct_Once sync.Once

func textViewAccessiblePrivateStruct_Set() {
	textViewAccessiblePrivateStruct_Once.Do(func() {
		textViewAccessiblePrivateStruct = gi.StructNew("Gtk", "TextViewAccessiblePrivate")
	})
}

type TextViewAccessiblePrivate struct {
	native uintptr
}

var textViewClassStruct *gi.Struct
var textViewClassStruct_Once sync.Once

func textViewClassStruct_Set() {
	textViewClassStruct_Once.Do(func() {
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
var textViewPrivateStruct_Once sync.Once

func textViewPrivateStruct_Set() {
	textViewPrivateStruct_Once.Do(func() {
		textViewPrivateStruct = gi.StructNew("Gtk", "TextViewPrivate")
	})
}

type TextViewPrivate struct {
	native uintptr
}

var themeEngineStruct *gi.Struct
var themeEngineStruct_Once sync.Once

func themeEngineStruct_Set() {
	themeEngineStruct_Once.Do(func() {
		themeEngineStruct = gi.StructNew("Gtk", "ThemeEngine")
	})
}

type ThemeEngine struct {
	native uintptr
}

var themingEngineClassStruct *gi.Struct
var themingEngineClassStruct_Once sync.Once

func themingEngineClassStruct_Set() {
	themingEngineClassStruct_Once.Do(func() {
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
var themingEnginePrivateStruct_Once sync.Once

func themingEnginePrivateStruct_Set() {
	themingEnginePrivateStruct_Once.Do(func() {
		themingEnginePrivateStruct = gi.StructNew("Gtk", "ThemingEnginePrivate")
	})
}

type ThemingEnginePrivate struct {
	native uintptr
}

var toggleActionClassStruct *gi.Struct
var toggleActionClassStruct_Once sync.Once

func toggleActionClassStruct_Set() {
	toggleActionClassStruct_Once.Do(func() {
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
var toggleActionEntryStruct_Once sync.Once

func toggleActionEntryStruct_Set() {
	toggleActionEntryStruct_Once.Do(func() {
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
var toggleActionPrivateStruct_Once sync.Once

func toggleActionPrivateStruct_Set() {
	toggleActionPrivateStruct_Once.Do(func() {
		toggleActionPrivateStruct = gi.StructNew("Gtk", "ToggleActionPrivate")
	})
}

type ToggleActionPrivate struct {
	native uintptr
}

var toggleButtonAccessibleClassStruct *gi.Struct
var toggleButtonAccessibleClassStruct_Once sync.Once

func toggleButtonAccessibleClassStruct_Set() {
	toggleButtonAccessibleClassStruct_Once.Do(func() {
		toggleButtonAccessibleClassStruct = gi.StructNew("Gtk", "ToggleButtonAccessibleClass")
	})
}

type ToggleButtonAccessibleClass struct {
	native      uintptr
	ParentClass *ButtonAccessibleClass
}

var toggleButtonAccessiblePrivateStruct *gi.Struct
var toggleButtonAccessiblePrivateStruct_Once sync.Once

func toggleButtonAccessiblePrivateStruct_Set() {
	toggleButtonAccessiblePrivateStruct_Once.Do(func() {
		toggleButtonAccessiblePrivateStruct = gi.StructNew("Gtk", "ToggleButtonAccessiblePrivate")
	})
}

type ToggleButtonAccessiblePrivate struct {
	native uintptr
}

var toggleButtonClassStruct *gi.Struct
var toggleButtonClassStruct_Once sync.Once

func toggleButtonClassStruct_Set() {
	toggleButtonClassStruct_Once.Do(func() {
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
var toggleButtonPrivateStruct_Once sync.Once

func toggleButtonPrivateStruct_Set() {
	toggleButtonPrivateStruct_Once.Do(func() {
		toggleButtonPrivateStruct = gi.StructNew("Gtk", "ToggleButtonPrivate")
	})
}

type ToggleButtonPrivate struct {
	native uintptr
}

var toggleToolButtonClassStruct *gi.Struct
var toggleToolButtonClassStruct_Once sync.Once

func toggleToolButtonClassStruct_Set() {
	toggleToolButtonClassStruct_Once.Do(func() {
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
var toggleToolButtonPrivateStruct_Once sync.Once

func toggleToolButtonPrivateStruct_Set() {
	toggleToolButtonPrivateStruct_Once.Do(func() {
		toggleToolButtonPrivateStruct = gi.StructNew("Gtk", "ToggleToolButtonPrivate")
	})
}

type ToggleToolButtonPrivate struct {
	native uintptr
}

var toolButtonClassStruct *gi.Struct
var toolButtonClassStruct_Once sync.Once

func toolButtonClassStruct_Set() {
	toolButtonClassStruct_Once.Do(func() {
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
var toolButtonPrivateStruct_Once sync.Once

func toolButtonPrivateStruct_Set() {
	toolButtonPrivateStruct_Once.Do(func() {
		toolButtonPrivateStruct = gi.StructNew("Gtk", "ToolButtonPrivate")
	})
}

type ToolButtonPrivate struct {
	native uintptr
}

var toolItemClassStruct *gi.Struct
var toolItemClassStruct_Once sync.Once

func toolItemClassStruct_Set() {
	toolItemClassStruct_Once.Do(func() {
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
var toolItemGroupClassStruct_Once sync.Once

func toolItemGroupClassStruct_Set() {
	toolItemGroupClassStruct_Once.Do(func() {
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
var toolItemGroupPrivateStruct_Once sync.Once

func toolItemGroupPrivateStruct_Set() {
	toolItemGroupPrivateStruct_Once.Do(func() {
		toolItemGroupPrivateStruct = gi.StructNew("Gtk", "ToolItemGroupPrivate")
	})
}

type ToolItemGroupPrivate struct {
	native uintptr
}

var toolItemPrivateStruct *gi.Struct
var toolItemPrivateStruct_Once sync.Once

func toolItemPrivateStruct_Set() {
	toolItemPrivateStruct_Once.Do(func() {
		toolItemPrivateStruct = gi.StructNew("Gtk", "ToolItemPrivate")
	})
}

type ToolItemPrivate struct {
	native uintptr
}

var toolPaletteClassStruct *gi.Struct
var toolPaletteClassStruct_Once sync.Once

func toolPaletteClassStruct_Set() {
	toolPaletteClassStruct_Once.Do(func() {
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
var toolPalettePrivateStruct_Once sync.Once

func toolPalettePrivateStruct_Set() {
	toolPalettePrivateStruct_Once.Do(func() {
		toolPalettePrivateStruct = gi.StructNew("Gtk", "ToolPalettePrivate")
	})
}

type ToolPalettePrivate struct {
	native uintptr
}

var toolShellIfaceStruct *gi.Struct
var toolShellIfaceStruct_Once sync.Once

func toolShellIfaceStruct_Set() {
	toolShellIfaceStruct_Once.Do(func() {
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
var toolbarClassStruct_Once sync.Once

func toolbarClassStruct_Set() {
	toolbarClassStruct_Once.Do(func() {
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
var toolbarPrivateStruct_Once sync.Once

func toolbarPrivateStruct_Set() {
	toolbarPrivateStruct_Once.Do(func() {
		toolbarPrivateStruct = gi.StructNew("Gtk", "ToolbarPrivate")
	})
}

type ToolbarPrivate struct {
	native uintptr
}

var toplevelAccessibleClassStruct *gi.Struct
var toplevelAccessibleClassStruct_Once sync.Once

func toplevelAccessibleClassStruct_Set() {
	toplevelAccessibleClassStruct_Once.Do(func() {
		toplevelAccessibleClassStruct = gi.StructNew("Gtk", "ToplevelAccessibleClass")
	})
}

type ToplevelAccessibleClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'Atk.ObjectClass'
}

var toplevelAccessiblePrivateStruct *gi.Struct
var toplevelAccessiblePrivateStruct_Once sync.Once

func toplevelAccessiblePrivateStruct_Set() {
	toplevelAccessiblePrivateStruct_Once.Do(func() {
		toplevelAccessiblePrivateStruct = gi.StructNew("Gtk", "ToplevelAccessiblePrivate")
	})
}

type ToplevelAccessiblePrivate struct {
	native uintptr
}

var treeDragDestIfaceStruct *gi.Struct
var treeDragDestIfaceStruct_Once sync.Once

func treeDragDestIfaceStruct_Set() {
	treeDragDestIfaceStruct_Once.Do(func() {
		treeDragDestIfaceStruct = gi.StructNew("Gtk", "TreeDragDestIface")
	})
}

type TreeDragDestIface struct {
	native uintptr
	// UNSUPPORTED : C value 'drag_data_received' : missing Type
	// UNSUPPORTED : C value 'row_drop_possible' : missing Type
}

var treeDragSourceIfaceStruct *gi.Struct
var treeDragSourceIfaceStruct_Once sync.Once

func treeDragSourceIfaceStruct_Set() {
	treeDragSourceIfaceStruct_Once.Do(func() {
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
var treeIterStruct_Once sync.Once

func treeIterStruct_Set() {
	treeIterStruct_Once.Do(func() {
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
		treeIterStruct_Set()
		treeIterCopyFunction = treeIterStruct.InvokerNew("copy")
	})
}

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
		treeIterStruct_Set()
		treeIterFreeFunction = treeIterStruct.InvokerNew("free")
	})
}

// Free is a representation of the C type gtk_tree_iter_free.
func (recv *TreeIter) Free() {
	treeIterFreeFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	treeIterFreeFunction.Invoke(inArgs[:], nil)

}

var treeModelFilterClassStruct *gi.Struct
var treeModelFilterClassStruct_Once sync.Once

func treeModelFilterClassStruct_Set() {
	treeModelFilterClassStruct_Once.Do(func() {
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
var treeModelFilterPrivateStruct_Once sync.Once

func treeModelFilterPrivateStruct_Set() {
	treeModelFilterPrivateStruct_Once.Do(func() {
		treeModelFilterPrivateStruct = gi.StructNew("Gtk", "TreeModelFilterPrivate")
	})
}

type TreeModelFilterPrivate struct {
	native uintptr
}

var treeModelIfaceStruct *gi.Struct
var treeModelIfaceStruct_Once sync.Once

func treeModelIfaceStruct_Set() {
	treeModelIfaceStruct_Once.Do(func() {
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
var treeModelSortClassStruct_Once sync.Once

func treeModelSortClassStruct_Set() {
	treeModelSortClassStruct_Once.Do(func() {
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
var treeModelSortPrivateStruct_Once sync.Once

func treeModelSortPrivateStruct_Set() {
	treeModelSortPrivateStruct_Once.Do(func() {
		treeModelSortPrivateStruct = gi.StructNew("Gtk", "TreeModelSortPrivate")
	})
}

type TreeModelSortPrivate struct {
	native uintptr
}

var treePathStruct *gi.Struct
var treePathStruct_Once sync.Once

func treePathStruct_Set() {
	treePathStruct_Once.Do(func() {
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
		treePathStruct_Set()
		treePathNewFunction = treePathStruct.InvokerNew("new")
	})
}

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
		treePathStruct_Set()
		treePathNewFirstFunction = treePathStruct.InvokerNew("new_first")
	})
}

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
		treePathStruct_Set()
		treePathNewFromStringFunction = treePathStruct.InvokerNew("new_from_string")
	})
}

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
		treePathStruct_Set()
		treePathAppendIndexFunction = treePathStruct.InvokerNew("append_index")
	})
}

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
		treePathStruct_Set()
		treePathCopyFunction = treePathStruct.InvokerNew("copy")
	})
}

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
		treePathStruct_Set()
		treePathDownFunction = treePathStruct.InvokerNew("down")
	})
}

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
		treePathStruct_Set()
		treePathFreeFunction = treePathStruct.InvokerNew("free")
	})
}

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
		treePathStruct_Set()
		treePathGetDepthFunction = treePathStruct.InvokerNew("get_depth")
	})
}

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
		treePathStruct_Set()
		treePathGetIndicesFunction = treePathStruct.InvokerNew("get_indices")
	})
}

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
		treePathStruct_Set()
		treePathGetIndicesWithDepthFunction = treePathStruct.InvokerNew("get_indices_with_depth")
	})
}

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
		treePathStruct_Set()
		treePathNextFunction = treePathStruct.InvokerNew("next")
	})
}

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
		treePathStruct_Set()
		treePathPrependIndexFunction = treePathStruct.InvokerNew("prepend_index")
	})
}

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
		treePathStruct_Set()
		treePathToStringFunction = treePathStruct.InvokerNew("to_string")
	})
}

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
var treeRowReferenceStruct_Once sync.Once

func treeRowReferenceStruct_Set() {
	treeRowReferenceStruct_Once.Do(func() {
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
		treeRowReferenceStruct_Set()
		treeRowReferenceCopyFunction = treeRowReferenceStruct.InvokerNew("copy")
	})
}

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
		treeRowReferenceStruct_Set()
		treeRowReferenceFreeFunction = treeRowReferenceStruct.InvokerNew("free")
	})
}

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
		treeRowReferenceStruct_Set()
		treeRowReferenceGetPathFunction = treeRowReferenceStruct.InvokerNew("get_path")
	})
}

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
var treeSelectionClassStruct_Once sync.Once

func treeSelectionClassStruct_Set() {
	treeSelectionClassStruct_Once.Do(func() {
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
var treeSelectionPrivateStruct_Once sync.Once

func treeSelectionPrivateStruct_Set() {
	treeSelectionPrivateStruct_Once.Do(func() {
		treeSelectionPrivateStruct = gi.StructNew("Gtk", "TreeSelectionPrivate")
	})
}

type TreeSelectionPrivate struct {
	native uintptr
}

var treeSortableIfaceStruct *gi.Struct
var treeSortableIfaceStruct_Once sync.Once

func treeSortableIfaceStruct_Set() {
	treeSortableIfaceStruct_Once.Do(func() {
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
var treeStoreClassStruct_Once sync.Once

func treeStoreClassStruct_Set() {
	treeStoreClassStruct_Once.Do(func() {
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
var treeStorePrivateStruct_Once sync.Once

func treeStorePrivateStruct_Set() {
	treeStorePrivateStruct_Once.Do(func() {
		treeStorePrivateStruct = gi.StructNew("Gtk", "TreeStorePrivate")
	})
}

type TreeStorePrivate struct {
	native uintptr
}

var treeViewAccessibleClassStruct *gi.Struct
var treeViewAccessibleClassStruct_Once sync.Once

func treeViewAccessibleClassStruct_Set() {
	treeViewAccessibleClassStruct_Once.Do(func() {
		treeViewAccessibleClassStruct = gi.StructNew("Gtk", "TreeViewAccessibleClass")
	})
}

type TreeViewAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var treeViewAccessiblePrivateStruct *gi.Struct
var treeViewAccessiblePrivateStruct_Once sync.Once

func treeViewAccessiblePrivateStruct_Set() {
	treeViewAccessiblePrivateStruct_Once.Do(func() {
		treeViewAccessiblePrivateStruct = gi.StructNew("Gtk", "TreeViewAccessiblePrivate")
	})
}

type TreeViewAccessiblePrivate struct {
	native uintptr
}

var treeViewClassStruct *gi.Struct
var treeViewClassStruct_Once sync.Once

func treeViewClassStruct_Set() {
	treeViewClassStruct_Once.Do(func() {
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
var treeViewColumnClassStruct_Once sync.Once

func treeViewColumnClassStruct_Set() {
	treeViewColumnClassStruct_Once.Do(func() {
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
var treeViewColumnPrivateStruct_Once sync.Once

func treeViewColumnPrivateStruct_Set() {
	treeViewColumnPrivateStruct_Once.Do(func() {
		treeViewColumnPrivateStruct = gi.StructNew("Gtk", "TreeViewColumnPrivate")
	})
}

type TreeViewColumnPrivate struct {
	native uintptr
}

var treeViewPrivateStruct *gi.Struct
var treeViewPrivateStruct_Once sync.Once

func treeViewPrivateStruct_Set() {
	treeViewPrivateStruct_Once.Do(func() {
		treeViewPrivateStruct = gi.StructNew("Gtk", "TreeViewPrivate")
	})
}

type TreeViewPrivate struct {
	native uintptr
}

var uIManagerClassStruct *gi.Struct
var uIManagerClassStruct_Once sync.Once

func uIManagerClassStruct_Set() {
	uIManagerClassStruct_Once.Do(func() {
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
var uIManagerPrivateStruct_Once sync.Once

func uIManagerPrivateStruct_Set() {
	uIManagerPrivateStruct_Once.Do(func() {
		uIManagerPrivateStruct = gi.StructNew("Gtk", "UIManagerPrivate")
	})
}

type UIManagerPrivate struct {
	native uintptr
}

var vBoxClassStruct *gi.Struct
var vBoxClassStruct_Once sync.Once

func vBoxClassStruct_Set() {
	vBoxClassStruct_Once.Do(func() {
		vBoxClassStruct = gi.StructNew("Gtk", "VBoxClass")
	})
}

type VBoxClass struct {
	native      uintptr
	ParentClass *BoxClass
}

var vButtonBoxClassStruct *gi.Struct
var vButtonBoxClassStruct_Once sync.Once

func vButtonBoxClassStruct_Set() {
	vButtonBoxClassStruct_Once.Do(func() {
		vButtonBoxClassStruct = gi.StructNew("Gtk", "VButtonBoxClass")
	})
}

type VButtonBoxClass struct {
	native      uintptr
	ParentClass *ButtonBoxClass
}

var vPanedClassStruct *gi.Struct
var vPanedClassStruct_Once sync.Once

func vPanedClassStruct_Set() {
	vPanedClassStruct_Once.Do(func() {
		vPanedClassStruct = gi.StructNew("Gtk", "VPanedClass")
	})
}

type VPanedClass struct {
	native      uintptr
	ParentClass *PanedClass
}

var vScaleClassStruct *gi.Struct
var vScaleClassStruct_Once sync.Once

func vScaleClassStruct_Set() {
	vScaleClassStruct_Once.Do(func() {
		vScaleClassStruct = gi.StructNew("Gtk", "VScaleClass")
	})
}

type VScaleClass struct {
	native      uintptr
	ParentClass *ScaleClass
}

var vScrollbarClassStruct *gi.Struct
var vScrollbarClassStruct_Once sync.Once

func vScrollbarClassStruct_Set() {
	vScrollbarClassStruct_Once.Do(func() {
		vScrollbarClassStruct = gi.StructNew("Gtk", "VScrollbarClass")
	})
}

type VScrollbarClass struct {
	native      uintptr
	ParentClass *ScrollbarClass
}

var vSeparatorClassStruct *gi.Struct
var vSeparatorClassStruct_Once sync.Once

func vSeparatorClassStruct_Set() {
	vSeparatorClassStruct_Once.Do(func() {
		vSeparatorClassStruct = gi.StructNew("Gtk", "VSeparatorClass")
	})
}

type VSeparatorClass struct {
	native      uintptr
	ParentClass *SeparatorClass
}

var viewportClassStruct *gi.Struct
var viewportClassStruct_Once sync.Once

func viewportClassStruct_Set() {
	viewportClassStruct_Once.Do(func() {
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
var viewportPrivateStruct_Once sync.Once

func viewportPrivateStruct_Set() {
	viewportPrivateStruct_Once.Do(func() {
		viewportPrivateStruct = gi.StructNew("Gtk", "ViewportPrivate")
	})
}

type ViewportPrivate struct {
	native uintptr
}

var volumeButtonClassStruct *gi.Struct
var volumeButtonClassStruct_Once sync.Once

func volumeButtonClassStruct_Set() {
	volumeButtonClassStruct_Once.Do(func() {
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
var widgetAccessibleClassStruct_Once sync.Once

func widgetAccessibleClassStruct_Set() {
	widgetAccessibleClassStruct_Once.Do(func() {
		widgetAccessibleClassStruct = gi.StructNew("Gtk", "WidgetAccessibleClass")
	})
}

type WidgetAccessibleClass struct {
	native      uintptr
	ParentClass *AccessibleClass
	// UNSUPPORTED : C value 'notify_gtk' : missing Type
}

var widgetAccessiblePrivateStruct *gi.Struct
var widgetAccessiblePrivateStruct_Once sync.Once

func widgetAccessiblePrivateStruct_Set() {
	widgetAccessiblePrivateStruct_Once.Do(func() {
		widgetAccessiblePrivateStruct = gi.StructNew("Gtk", "WidgetAccessiblePrivate")
	})
}

type WidgetAccessiblePrivate struct {
	native uintptr
}

var widgetClassStruct *gi.Struct
var widgetClassStruct_Once sync.Once

func widgetClassStruct_Set() {
	widgetClassStruct_Once.Do(func() {
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
		widgetClassStruct_Set()
		widgetClassGetCssNameFunction = widgetClassStruct.InvokerNew("get_css_name")
	})
}

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
		widgetClassStruct_Set()
		widgetClassListStylePropertiesFunction = widgetClassStruct.InvokerNew("list_style_properties")
	})
}

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
		widgetClassStruct_Set()
		widgetClassSetCssNameFunction = widgetClassStruct.InvokerNew("set_css_name")
	})
}

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
		widgetClassStruct_Set()
		widgetClassSetTemplateFromResourceFunction = widgetClassStruct.InvokerNew("set_template_from_resource")
	})
}

// SetTemplateFromResource is a representation of the C type gtk_widget_class_set_template_from_resource.
func (recv *WidgetClass) SetTemplateFromResource(resourceName string) {
	widgetClassSetTemplateFromResourceFunction_Set()

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(resourceName)

	widgetClassSetTemplateFromResourceFunction.Invoke(inArgs[:], nil)

}

var widgetClassPrivateStruct *gi.Struct
var widgetClassPrivateStruct_Once sync.Once

func widgetClassPrivateStruct_Set() {
	widgetClassPrivateStruct_Once.Do(func() {
		widgetClassPrivateStruct = gi.StructNew("Gtk", "WidgetClassPrivate")
	})
}

type WidgetClassPrivate struct {
	native uintptr
}

var widgetPathStruct *gi.Struct
var widgetPathStruct_Once sync.Once

func widgetPathStruct_Set() {
	widgetPathStruct_Once.Do(func() {
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
		widgetPathStruct_Set()
		widgetPathNewFunction = widgetPathStruct.InvokerNew("new")
	})
}

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
		widgetPathStruct_Set()
		widgetPathCopyFunction = widgetPathStruct.InvokerNew("copy")
	})
}

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
		widgetPathStruct_Set()
		widgetPathFreeFunction = widgetPathStruct.InvokerNew("free")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterAddClassFunction = widgetPathStruct.InvokerNew("iter_add_class")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterClearClassesFunction = widgetPathStruct.InvokerNew("iter_clear_classes")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterClearRegionsFunction = widgetPathStruct.InvokerNew("iter_clear_regions")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterGetNameFunction = widgetPathStruct.InvokerNew("iter_get_name")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterGetObjectNameFunction = widgetPathStruct.InvokerNew("iter_get_object_name")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterGetSiblingIndexFunction = widgetPathStruct.InvokerNew("iter_get_sibling_index")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterGetSiblingsFunction = widgetPathStruct.InvokerNew("iter_get_siblings")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterRemoveClassFunction = widgetPathStruct.InvokerNew("iter_remove_class")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterRemoveRegionFunction = widgetPathStruct.InvokerNew("iter_remove_region")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterSetNameFunction = widgetPathStruct.InvokerNew("iter_set_name")
	})
}

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
		widgetPathStruct_Set()
		widgetPathIterSetObjectNameFunction = widgetPathStruct.InvokerNew("iter_set_object_name")
	})
}

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
		widgetPathStruct_Set()
		widgetPathLengthFunction = widgetPathStruct.InvokerNew("length")
	})
}

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
		widgetPathStruct_Set()
		widgetPathRefFunction = widgetPathStruct.InvokerNew("ref")
	})
}

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
		widgetPathStruct_Set()
		widgetPathToStringFunction = widgetPathStruct.InvokerNew("to_string")
	})
}

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
		widgetPathStruct_Set()
		widgetPathUnrefFunction = widgetPathStruct.InvokerNew("unref")
	})
}

// Unref is a representation of the C type gtk_widget_path_unref.
func (recv *WidgetPath) Unref() {
	widgetPathUnrefFunction_Set()

	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	widgetPathUnrefFunction.Invoke(inArgs[:], nil)

}

var widgetPrivateStruct *gi.Struct
var widgetPrivateStruct_Once sync.Once

func widgetPrivateStruct_Set() {
	widgetPrivateStruct_Once.Do(func() {
		widgetPrivateStruct = gi.StructNew("Gtk", "WidgetPrivate")
	})
}

type WidgetPrivate struct {
	native uintptr
}

var windowAccessibleClassStruct *gi.Struct
var windowAccessibleClassStruct_Once sync.Once

func windowAccessibleClassStruct_Set() {
	windowAccessibleClassStruct_Once.Do(func() {
		windowAccessibleClassStruct = gi.StructNew("Gtk", "WindowAccessibleClass")
	})
}

type WindowAccessibleClass struct {
	native      uintptr
	ParentClass *ContainerAccessibleClass
}

var windowAccessiblePrivateStruct *gi.Struct
var windowAccessiblePrivateStruct_Once sync.Once

func windowAccessiblePrivateStruct_Set() {
	windowAccessiblePrivateStruct_Once.Do(func() {
		windowAccessiblePrivateStruct = gi.StructNew("Gtk", "WindowAccessiblePrivate")
	})
}

type WindowAccessiblePrivate struct {
	native uintptr
}

var windowClassStruct *gi.Struct
var windowClassStruct_Once sync.Once

func windowClassStruct_Set() {
	windowClassStruct_Once.Do(func() {
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
var windowGeometryInfoStruct_Once sync.Once

func windowGeometryInfoStruct_Set() {
	windowGeometryInfoStruct_Once.Do(func() {
		windowGeometryInfoStruct = gi.StructNew("Gtk", "WindowGeometryInfo")
	})
}

type WindowGeometryInfo struct {
	native uintptr
}

var windowGroupClassStruct *gi.Struct
var windowGroupClassStruct_Once sync.Once

func windowGroupClassStruct_Set() {
	windowGroupClassStruct_Once.Do(func() {
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
var windowGroupPrivateStruct_Once sync.Once

func windowGroupPrivateStruct_Set() {
	windowGroupPrivateStruct_Once.Do(func() {
		windowGroupPrivateStruct = gi.StructNew("Gtk", "WindowGroupPrivate")
	})
}

type WindowGroupPrivate struct {
	native uintptr
}

var windowPrivateStruct *gi.Struct
var windowPrivateStruct_Once sync.Once

func windowPrivateStruct_Set() {
	windowPrivateStruct_Once.Do(func() {
		windowPrivateStruct = gi.StructNew("Gtk", "WindowPrivate")
	})
}

type WindowPrivate struct {
	native uintptr
}
