// Code generated - DO NOT EDIT.

package gtk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"sync"
)

var aboutDialogClassStruct *gi.Struct
var aboutDialogClassStruct_Once sync.Once

func aboutDialogClassStruct_Set() error {
	var err error
	aboutDialogClassStruct_Once.Do(func() {
		aboutDialogClassStruct, err = gi.StructNew("Gtk", "AboutDialogClass")
	})
	return err
}

type AboutDialogClass struct {
	native uintptr
}

var aboutDialogPrivateStruct *gi.Struct
var aboutDialogPrivateStruct_Once sync.Once

func aboutDialogPrivateStruct_Set() error {
	var err error
	aboutDialogPrivateStruct_Once.Do(func() {
		aboutDialogPrivateStruct, err = gi.StructNew("Gtk", "AboutDialogPrivate")
	})
	return err
}

type AboutDialogPrivate struct {
	native uintptr
}

var accelGroupClassStruct *gi.Struct
var accelGroupClassStruct_Once sync.Once

func accelGroupClassStruct_Set() error {
	var err error
	accelGroupClassStruct_Once.Do(func() {
		accelGroupClassStruct, err = gi.StructNew("Gtk", "AccelGroupClass")
	})
	return err
}

type AccelGroupClass struct {
	native uintptr
}

var accelGroupEntryStruct *gi.Struct
var accelGroupEntryStruct_Once sync.Once

func accelGroupEntryStruct_Set() error {
	var err error
	accelGroupEntryStruct_Once.Do(func() {
		accelGroupEntryStruct, err = gi.StructNew("Gtk", "AccelGroupEntry")
	})
	return err
}

type AccelGroupEntry struct {
	native uintptr
}

var accelGroupPrivateStruct *gi.Struct
var accelGroupPrivateStruct_Once sync.Once

func accelGroupPrivateStruct_Set() error {
	var err error
	accelGroupPrivateStruct_Once.Do(func() {
		accelGroupPrivateStruct, err = gi.StructNew("Gtk", "AccelGroupPrivate")
	})
	return err
}

type AccelGroupPrivate struct {
	native uintptr
}

var accelKeyStruct *gi.Struct
var accelKeyStruct_Once sync.Once

func accelKeyStruct_Set() error {
	var err error
	accelKeyStruct_Once.Do(func() {
		accelKeyStruct, err = gi.StructNew("Gtk", "AccelKey")
	})
	return err
}

type AccelKey struct {
	native uintptr
}

var accelLabelClassStruct *gi.Struct
var accelLabelClassStruct_Once sync.Once

func accelLabelClassStruct_Set() error {
	var err error
	accelLabelClassStruct_Once.Do(func() {
		accelLabelClassStruct, err = gi.StructNew("Gtk", "AccelLabelClass")
	})
	return err
}

type AccelLabelClass struct {
	native uintptr
}

var accelLabelPrivateStruct *gi.Struct
var accelLabelPrivateStruct_Once sync.Once

func accelLabelPrivateStruct_Set() error {
	var err error
	accelLabelPrivateStruct_Once.Do(func() {
		accelLabelPrivateStruct, err = gi.StructNew("Gtk", "AccelLabelPrivate")
	})
	return err
}

type AccelLabelPrivate struct {
	native uintptr
}

var accelMapClassStruct *gi.Struct
var accelMapClassStruct_Once sync.Once

func accelMapClassStruct_Set() error {
	var err error
	accelMapClassStruct_Once.Do(func() {
		accelMapClassStruct, err = gi.StructNew("Gtk", "AccelMapClass")
	})
	return err
}

type AccelMapClass struct {
	native uintptr
}

var accessibleClassStruct *gi.Struct
var accessibleClassStruct_Once sync.Once

func accessibleClassStruct_Set() error {
	var err error
	accessibleClassStruct_Once.Do(func() {
		accessibleClassStruct, err = gi.StructNew("Gtk", "AccessibleClass")
	})
	return err
}

type AccessibleClass struct {
	native uintptr
}

var accessiblePrivateStruct *gi.Struct
var accessiblePrivateStruct_Once sync.Once

func accessiblePrivateStruct_Set() error {
	var err error
	accessiblePrivateStruct_Once.Do(func() {
		accessiblePrivateStruct, err = gi.StructNew("Gtk", "AccessiblePrivate")
	})
	return err
}

type AccessiblePrivate struct {
	native uintptr
}

var actionBarClassStruct *gi.Struct
var actionBarClassStruct_Once sync.Once

func actionBarClassStruct_Set() error {
	var err error
	actionBarClassStruct_Once.Do(func() {
		actionBarClassStruct, err = gi.StructNew("Gtk", "ActionBarClass")
	})
	return err
}

type ActionBarClass struct {
	native uintptr
}

var actionBarPrivateStruct *gi.Struct
var actionBarPrivateStruct_Once sync.Once

func actionBarPrivateStruct_Set() error {
	var err error
	actionBarPrivateStruct_Once.Do(func() {
		actionBarPrivateStruct, err = gi.StructNew("Gtk", "ActionBarPrivate")
	})
	return err
}

type ActionBarPrivate struct {
	native uintptr
}

var actionClassStruct *gi.Struct
var actionClassStruct_Once sync.Once

func actionClassStruct_Set() error {
	var err error
	actionClassStruct_Once.Do(func() {
		actionClassStruct, err = gi.StructNew("Gtk", "ActionClass")
	})
	return err
}

type ActionClass struct {
	native uintptr
}

var actionEntryStruct *gi.Struct
var actionEntryStruct_Once sync.Once

func actionEntryStruct_Set() error {
	var err error
	actionEntryStruct_Once.Do(func() {
		actionEntryStruct, err = gi.StructNew("Gtk", "ActionEntry")
	})
	return err
}

type ActionEntry struct {
	native uintptr
}

var actionGroupClassStruct *gi.Struct
var actionGroupClassStruct_Once sync.Once

func actionGroupClassStruct_Set() error {
	var err error
	actionGroupClassStruct_Once.Do(func() {
		actionGroupClassStruct, err = gi.StructNew("Gtk", "ActionGroupClass")
	})
	return err
}

type ActionGroupClass struct {
	native uintptr
}

var actionGroupPrivateStruct *gi.Struct
var actionGroupPrivateStruct_Once sync.Once

func actionGroupPrivateStruct_Set() error {
	var err error
	actionGroupPrivateStruct_Once.Do(func() {
		actionGroupPrivateStruct, err = gi.StructNew("Gtk", "ActionGroupPrivate")
	})
	return err
}

type ActionGroupPrivate struct {
	native uintptr
}

var actionPrivateStruct *gi.Struct
var actionPrivateStruct_Once sync.Once

func actionPrivateStruct_Set() error {
	var err error
	actionPrivateStruct_Once.Do(func() {
		actionPrivateStruct, err = gi.StructNew("Gtk", "ActionPrivate")
	})
	return err
}

type ActionPrivate struct {
	native uintptr
}

var actionableInterfaceStruct *gi.Struct
var actionableInterfaceStruct_Once sync.Once

func actionableInterfaceStruct_Set() error {
	var err error
	actionableInterfaceStruct_Once.Do(func() {
		actionableInterfaceStruct, err = gi.StructNew("Gtk", "ActionableInterface")
	})
	return err
}

type ActionableInterface struct {
	native uintptr
}

var activatableIfaceStruct *gi.Struct
var activatableIfaceStruct_Once sync.Once

func activatableIfaceStruct_Set() error {
	var err error
	activatableIfaceStruct_Once.Do(func() {
		activatableIfaceStruct, err = gi.StructNew("Gtk", "ActivatableIface")
	})
	return err
}

type ActivatableIface struct {
	native uintptr
}

var adjustmentClassStruct *gi.Struct
var adjustmentClassStruct_Once sync.Once

func adjustmentClassStruct_Set() error {
	var err error
	adjustmentClassStruct_Once.Do(func() {
		adjustmentClassStruct, err = gi.StructNew("Gtk", "AdjustmentClass")
	})
	return err
}

type AdjustmentClass struct {
	native uintptr
}

var adjustmentPrivateStruct *gi.Struct
var adjustmentPrivateStruct_Once sync.Once

func adjustmentPrivateStruct_Set() error {
	var err error
	adjustmentPrivateStruct_Once.Do(func() {
		adjustmentPrivateStruct, err = gi.StructNew("Gtk", "AdjustmentPrivate")
	})
	return err
}

type AdjustmentPrivate struct {
	native uintptr
}

var alignmentClassStruct *gi.Struct
var alignmentClassStruct_Once sync.Once

func alignmentClassStruct_Set() error {
	var err error
	alignmentClassStruct_Once.Do(func() {
		alignmentClassStruct, err = gi.StructNew("Gtk", "AlignmentClass")
	})
	return err
}

type AlignmentClass struct {
	native uintptr
}

var alignmentPrivateStruct *gi.Struct
var alignmentPrivateStruct_Once sync.Once

func alignmentPrivateStruct_Set() error {
	var err error
	alignmentPrivateStruct_Once.Do(func() {
		alignmentPrivateStruct, err = gi.StructNew("Gtk", "AlignmentPrivate")
	})
	return err
}

type AlignmentPrivate struct {
	native uintptr
}

var appChooserButtonClassStruct *gi.Struct
var appChooserButtonClassStruct_Once sync.Once

func appChooserButtonClassStruct_Set() error {
	var err error
	appChooserButtonClassStruct_Once.Do(func() {
		appChooserButtonClassStruct, err = gi.StructNew("Gtk", "AppChooserButtonClass")
	})
	return err
}

type AppChooserButtonClass struct {
	native uintptr
}

var appChooserButtonPrivateStruct *gi.Struct
var appChooserButtonPrivateStruct_Once sync.Once

func appChooserButtonPrivateStruct_Set() error {
	var err error
	appChooserButtonPrivateStruct_Once.Do(func() {
		appChooserButtonPrivateStruct, err = gi.StructNew("Gtk", "AppChooserButtonPrivate")
	})
	return err
}

type AppChooserButtonPrivate struct {
	native uintptr
}

var appChooserDialogClassStruct *gi.Struct
var appChooserDialogClassStruct_Once sync.Once

func appChooserDialogClassStruct_Set() error {
	var err error
	appChooserDialogClassStruct_Once.Do(func() {
		appChooserDialogClassStruct, err = gi.StructNew("Gtk", "AppChooserDialogClass")
	})
	return err
}

type AppChooserDialogClass struct {
	native uintptr
}

var appChooserDialogPrivateStruct *gi.Struct
var appChooserDialogPrivateStruct_Once sync.Once

func appChooserDialogPrivateStruct_Set() error {
	var err error
	appChooserDialogPrivateStruct_Once.Do(func() {
		appChooserDialogPrivateStruct, err = gi.StructNew("Gtk", "AppChooserDialogPrivate")
	})
	return err
}

type AppChooserDialogPrivate struct {
	native uintptr
}

var appChooserWidgetClassStruct *gi.Struct
var appChooserWidgetClassStruct_Once sync.Once

func appChooserWidgetClassStruct_Set() error {
	var err error
	appChooserWidgetClassStruct_Once.Do(func() {
		appChooserWidgetClassStruct, err = gi.StructNew("Gtk", "AppChooserWidgetClass")
	})
	return err
}

type AppChooserWidgetClass struct {
	native uintptr
}

var appChooserWidgetPrivateStruct *gi.Struct
var appChooserWidgetPrivateStruct_Once sync.Once

func appChooserWidgetPrivateStruct_Set() error {
	var err error
	appChooserWidgetPrivateStruct_Once.Do(func() {
		appChooserWidgetPrivateStruct, err = gi.StructNew("Gtk", "AppChooserWidgetPrivate")
	})
	return err
}

type AppChooserWidgetPrivate struct {
	native uintptr
}

var applicationClassStruct *gi.Struct
var applicationClassStruct_Once sync.Once

func applicationClassStruct_Set() error {
	var err error
	applicationClassStruct_Once.Do(func() {
		applicationClassStruct, err = gi.StructNew("Gtk", "ApplicationClass")
	})
	return err
}

type ApplicationClass struct {
	native uintptr
}

var applicationPrivateStruct *gi.Struct
var applicationPrivateStruct_Once sync.Once

func applicationPrivateStruct_Set() error {
	var err error
	applicationPrivateStruct_Once.Do(func() {
		applicationPrivateStruct, err = gi.StructNew("Gtk", "ApplicationPrivate")
	})
	return err
}

type ApplicationPrivate struct {
	native uintptr
}

var applicationWindowClassStruct *gi.Struct
var applicationWindowClassStruct_Once sync.Once

func applicationWindowClassStruct_Set() error {
	var err error
	applicationWindowClassStruct_Once.Do(func() {
		applicationWindowClassStruct, err = gi.StructNew("Gtk", "ApplicationWindowClass")
	})
	return err
}

type ApplicationWindowClass struct {
	native uintptr
}

var applicationWindowPrivateStruct *gi.Struct
var applicationWindowPrivateStruct_Once sync.Once

func applicationWindowPrivateStruct_Set() error {
	var err error
	applicationWindowPrivateStruct_Once.Do(func() {
		applicationWindowPrivateStruct, err = gi.StructNew("Gtk", "ApplicationWindowPrivate")
	})
	return err
}

type ApplicationWindowPrivate struct {
	native uintptr
}

var arrowAccessibleClassStruct *gi.Struct
var arrowAccessibleClassStruct_Once sync.Once

func arrowAccessibleClassStruct_Set() error {
	var err error
	arrowAccessibleClassStruct_Once.Do(func() {
		arrowAccessibleClassStruct, err = gi.StructNew("Gtk", "ArrowAccessibleClass")
	})
	return err
}

type ArrowAccessibleClass struct {
	native uintptr
}

var arrowAccessiblePrivateStruct *gi.Struct
var arrowAccessiblePrivateStruct_Once sync.Once

func arrowAccessiblePrivateStruct_Set() error {
	var err error
	arrowAccessiblePrivateStruct_Once.Do(func() {
		arrowAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ArrowAccessiblePrivate")
	})
	return err
}

type ArrowAccessiblePrivate struct {
	native uintptr
}

var arrowClassStruct *gi.Struct
var arrowClassStruct_Once sync.Once

func arrowClassStruct_Set() error {
	var err error
	arrowClassStruct_Once.Do(func() {
		arrowClassStruct, err = gi.StructNew("Gtk", "ArrowClass")
	})
	return err
}

type ArrowClass struct {
	native uintptr
}

var arrowPrivateStruct *gi.Struct
var arrowPrivateStruct_Once sync.Once

func arrowPrivateStruct_Set() error {
	var err error
	arrowPrivateStruct_Once.Do(func() {
		arrowPrivateStruct, err = gi.StructNew("Gtk", "ArrowPrivate")
	})
	return err
}

type ArrowPrivate struct {
	native uintptr
}

var aspectFrameClassStruct *gi.Struct
var aspectFrameClassStruct_Once sync.Once

func aspectFrameClassStruct_Set() error {
	var err error
	aspectFrameClassStruct_Once.Do(func() {
		aspectFrameClassStruct, err = gi.StructNew("Gtk", "AspectFrameClass")
	})
	return err
}

type AspectFrameClass struct {
	native uintptr
}

var aspectFramePrivateStruct *gi.Struct
var aspectFramePrivateStruct_Once sync.Once

func aspectFramePrivateStruct_Set() error {
	var err error
	aspectFramePrivateStruct_Once.Do(func() {
		aspectFramePrivateStruct, err = gi.StructNew("Gtk", "AspectFramePrivate")
	})
	return err
}

type AspectFramePrivate struct {
	native uintptr
}

var assistantClassStruct *gi.Struct
var assistantClassStruct_Once sync.Once

func assistantClassStruct_Set() error {
	var err error
	assistantClassStruct_Once.Do(func() {
		assistantClassStruct, err = gi.StructNew("Gtk", "AssistantClass")
	})
	return err
}

type AssistantClass struct {
	native uintptr
}

var assistantPrivateStruct *gi.Struct
var assistantPrivateStruct_Once sync.Once

func assistantPrivateStruct_Set() error {
	var err error
	assistantPrivateStruct_Once.Do(func() {
		assistantPrivateStruct, err = gi.StructNew("Gtk", "AssistantPrivate")
	})
	return err
}

type AssistantPrivate struct {
	native uintptr
}

var binClassStruct *gi.Struct
var binClassStruct_Once sync.Once

func binClassStruct_Set() error {
	var err error
	binClassStruct_Once.Do(func() {
		binClassStruct, err = gi.StructNew("Gtk", "BinClass")
	})
	return err
}

type BinClass struct {
	native uintptr
}

var binPrivateStruct *gi.Struct
var binPrivateStruct_Once sync.Once

func binPrivateStruct_Set() error {
	var err error
	binPrivateStruct_Once.Do(func() {
		binPrivateStruct, err = gi.StructNew("Gtk", "BinPrivate")
	})
	return err
}

type BinPrivate struct {
	native uintptr
}

var bindingArgStruct *gi.Struct
var bindingArgStruct_Once sync.Once

func bindingArgStruct_Set() error {
	var err error
	bindingArgStruct_Once.Do(func() {
		bindingArgStruct, err = gi.StructNew("Gtk", "BindingArg")
	})
	return err
}

type BindingArg struct {
	native uintptr
}

var bindingEntryStruct *gi.Struct
var bindingEntryStruct_Once sync.Once

func bindingEntryStruct_Set() error {
	var err error
	bindingEntryStruct_Once.Do(func() {
		bindingEntryStruct, err = gi.StructNew("Gtk", "BindingEntry")
	})
	return err
}

type BindingEntry struct {
	native uintptr
}

var bindingSetStruct *gi.Struct
var bindingSetStruct_Once sync.Once

func bindingSetStruct_Set() error {
	var err error
	bindingSetStruct_Once.Do(func() {
		bindingSetStruct, err = gi.StructNew("Gtk", "BindingSet")
	})
	return err
}

type BindingSet struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_binding_set_activate' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_binding_set_add_path' : parameter 'path_type' of type 'PathType' not supported

var bindingSignalStruct *gi.Struct
var bindingSignalStruct_Once sync.Once

func bindingSignalStruct_Set() error {
	var err error
	bindingSignalStruct_Once.Do(func() {
		bindingSignalStruct, err = gi.StructNew("Gtk", "BindingSignal")
	})
	return err
}

type BindingSignal struct {
	native uintptr
}

var booleanCellAccessibleClassStruct *gi.Struct
var booleanCellAccessibleClassStruct_Once sync.Once

func booleanCellAccessibleClassStruct_Set() error {
	var err error
	booleanCellAccessibleClassStruct_Once.Do(func() {
		booleanCellAccessibleClassStruct, err = gi.StructNew("Gtk", "BooleanCellAccessibleClass")
	})
	return err
}

type BooleanCellAccessibleClass struct {
	native uintptr
}

var booleanCellAccessiblePrivateStruct *gi.Struct
var booleanCellAccessiblePrivateStruct_Once sync.Once

func booleanCellAccessiblePrivateStruct_Set() error {
	var err error
	booleanCellAccessiblePrivateStruct_Once.Do(func() {
		booleanCellAccessiblePrivateStruct, err = gi.StructNew("Gtk", "BooleanCellAccessiblePrivate")
	})
	return err
}

type BooleanCellAccessiblePrivate struct {
	native uintptr
}

var borderStruct *gi.Struct
var borderStruct_Once sync.Once

func borderStruct_Set() error {
	var err error
	borderStruct_Once.Do(func() {
		borderStruct, err = gi.StructNew("Gtk", "Border")
	})
	return err
}

type Border struct {
	native uintptr
}

var borderNewFunction *gi.Function
var borderNewFunction_Once sync.Once

func borderNewFunction_Set() error {
	var err error
	borderNewFunction_Once.Do(func() {
		err = borderStruct_Set()
		if err != nil {
			return
		}
		borderNewFunction, err = borderStruct.InvokerNew("new")
	})
	return err
}

// BorderNew is a representation of the C type gtk_border_new.
func BorderNew() *Border {

	var ret gi.Argument

	err := borderNewFunction_Set()
	if err == nil {
		ret = borderNewFunction.Invoke(nil, nil)
	}

	retGo := &Border{native: ret.Pointer()}

	return retGo
}

var borderCopyFunction *gi.Function
var borderCopyFunction_Once sync.Once

func borderCopyFunction_Set() error {
	var err error
	borderCopyFunction_Once.Do(func() {
		err = borderStruct_Set()
		if err != nil {
			return
		}
		borderCopyFunction, err = borderStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_border_copy.
func (recv *Border) Copy() *Border {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := borderCopyFunction_Set()
	if err == nil {
		ret = borderCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Border{native: ret.Pointer()}

	return retGo
}

var borderFreeFunction *gi.Function
var borderFreeFunction_Once sync.Once

func borderFreeFunction_Set() error {
	var err error
	borderFreeFunction_Once.Do(func() {
		err = borderStruct_Set()
		if err != nil {
			return
		}
		borderFreeFunction, err = borderStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_border_free.
func (recv *Border) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := borderFreeFunction_Set()
	if err == nil {
		borderFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var boxClassStruct *gi.Struct
var boxClassStruct_Once sync.Once

func boxClassStruct_Set() error {
	var err error
	boxClassStruct_Once.Do(func() {
		boxClassStruct, err = gi.StructNew("Gtk", "BoxClass")
	})
	return err
}

type BoxClass struct {
	native uintptr
}

var boxPrivateStruct *gi.Struct
var boxPrivateStruct_Once sync.Once

func boxPrivateStruct_Set() error {
	var err error
	boxPrivateStruct_Once.Do(func() {
		boxPrivateStruct, err = gi.StructNew("Gtk", "BoxPrivate")
	})
	return err
}

type BoxPrivate struct {
	native uintptr
}

var buildableIfaceStruct *gi.Struct
var buildableIfaceStruct_Once sync.Once

func buildableIfaceStruct_Set() error {
	var err error
	buildableIfaceStruct_Once.Do(func() {
		buildableIfaceStruct, err = gi.StructNew("Gtk", "BuildableIface")
	})
	return err
}

type BuildableIface struct {
	native uintptr
}

var builderClassStruct *gi.Struct
var builderClassStruct_Once sync.Once

func builderClassStruct_Set() error {
	var err error
	builderClassStruct_Once.Do(func() {
		builderClassStruct, err = gi.StructNew("Gtk", "BuilderClass")
	})
	return err
}

type BuilderClass struct {
	native uintptr
}

var builderPrivateStruct *gi.Struct
var builderPrivateStruct_Once sync.Once

func builderPrivateStruct_Set() error {
	var err error
	builderPrivateStruct_Once.Do(func() {
		builderPrivateStruct, err = gi.StructNew("Gtk", "BuilderPrivate")
	})
	return err
}

type BuilderPrivate struct {
	native uintptr
}

var buttonAccessibleClassStruct *gi.Struct
var buttonAccessibleClassStruct_Once sync.Once

func buttonAccessibleClassStruct_Set() error {
	var err error
	buttonAccessibleClassStruct_Once.Do(func() {
		buttonAccessibleClassStruct, err = gi.StructNew("Gtk", "ButtonAccessibleClass")
	})
	return err
}

type ButtonAccessibleClass struct {
	native uintptr
}

var buttonAccessiblePrivateStruct *gi.Struct
var buttonAccessiblePrivateStruct_Once sync.Once

func buttonAccessiblePrivateStruct_Set() error {
	var err error
	buttonAccessiblePrivateStruct_Once.Do(func() {
		buttonAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ButtonAccessiblePrivate")
	})
	return err
}

type ButtonAccessiblePrivate struct {
	native uintptr
}

var buttonBoxClassStruct *gi.Struct
var buttonBoxClassStruct_Once sync.Once

func buttonBoxClassStruct_Set() error {
	var err error
	buttonBoxClassStruct_Once.Do(func() {
		buttonBoxClassStruct, err = gi.StructNew("Gtk", "ButtonBoxClass")
	})
	return err
}

type ButtonBoxClass struct {
	native uintptr
}

var buttonBoxPrivateStruct *gi.Struct
var buttonBoxPrivateStruct_Once sync.Once

func buttonBoxPrivateStruct_Set() error {
	var err error
	buttonBoxPrivateStruct_Once.Do(func() {
		buttonBoxPrivateStruct, err = gi.StructNew("Gtk", "ButtonBoxPrivate")
	})
	return err
}

type ButtonBoxPrivate struct {
	native uintptr
}

var buttonClassStruct *gi.Struct
var buttonClassStruct_Once sync.Once

func buttonClassStruct_Set() error {
	var err error
	buttonClassStruct_Once.Do(func() {
		buttonClassStruct, err = gi.StructNew("Gtk", "ButtonClass")
	})
	return err
}

type ButtonClass struct {
	native uintptr
}

var buttonPrivateStruct *gi.Struct
var buttonPrivateStruct_Once sync.Once

func buttonPrivateStruct_Set() error {
	var err error
	buttonPrivateStruct_Once.Do(func() {
		buttonPrivateStruct, err = gi.StructNew("Gtk", "ButtonPrivate")
	})
	return err
}

type ButtonPrivate struct {
	native uintptr
}

var calendarClassStruct *gi.Struct
var calendarClassStruct_Once sync.Once

func calendarClassStruct_Set() error {
	var err error
	calendarClassStruct_Once.Do(func() {
		calendarClassStruct, err = gi.StructNew("Gtk", "CalendarClass")
	})
	return err
}

type CalendarClass struct {
	native uintptr
}

var calendarPrivateStruct *gi.Struct
var calendarPrivateStruct_Once sync.Once

func calendarPrivateStruct_Set() error {
	var err error
	calendarPrivateStruct_Once.Do(func() {
		calendarPrivateStruct, err = gi.StructNew("Gtk", "CalendarPrivate")
	})
	return err
}

type CalendarPrivate struct {
	native uintptr
}

var cellAccessibleClassStruct *gi.Struct
var cellAccessibleClassStruct_Once sync.Once

func cellAccessibleClassStruct_Set() error {
	var err error
	cellAccessibleClassStruct_Once.Do(func() {
		cellAccessibleClassStruct, err = gi.StructNew("Gtk", "CellAccessibleClass")
	})
	return err
}

type CellAccessibleClass struct {
	native uintptr
}

var cellAccessibleParentIfaceStruct *gi.Struct
var cellAccessibleParentIfaceStruct_Once sync.Once

func cellAccessibleParentIfaceStruct_Set() error {
	var err error
	cellAccessibleParentIfaceStruct_Once.Do(func() {
		cellAccessibleParentIfaceStruct, err = gi.StructNew("Gtk", "CellAccessibleParentIface")
	})
	return err
}

type CellAccessibleParentIface struct {
	native uintptr
}

var cellAccessiblePrivateStruct *gi.Struct
var cellAccessiblePrivateStruct_Once sync.Once

func cellAccessiblePrivateStruct_Set() error {
	var err error
	cellAccessiblePrivateStruct_Once.Do(func() {
		cellAccessiblePrivateStruct, err = gi.StructNew("Gtk", "CellAccessiblePrivate")
	})
	return err
}

type CellAccessiblePrivate struct {
	native uintptr
}

var cellAreaBoxClassStruct *gi.Struct
var cellAreaBoxClassStruct_Once sync.Once

func cellAreaBoxClassStruct_Set() error {
	var err error
	cellAreaBoxClassStruct_Once.Do(func() {
		cellAreaBoxClassStruct, err = gi.StructNew("Gtk", "CellAreaBoxClass")
	})
	return err
}

type CellAreaBoxClass struct {
	native uintptr
}

var cellAreaBoxPrivateStruct *gi.Struct
var cellAreaBoxPrivateStruct_Once sync.Once

func cellAreaBoxPrivateStruct_Set() error {
	var err error
	cellAreaBoxPrivateStruct_Once.Do(func() {
		cellAreaBoxPrivateStruct, err = gi.StructNew("Gtk", "CellAreaBoxPrivate")
	})
	return err
}

type CellAreaBoxPrivate struct {
	native uintptr
}

var cellAreaClassStruct *gi.Struct
var cellAreaClassStruct_Once sync.Once

func cellAreaClassStruct_Set() error {
	var err error
	cellAreaClassStruct_Once.Do(func() {
		cellAreaClassStruct, err = gi.StructNew("Gtk", "CellAreaClass")
	})
	return err
}

type CellAreaClass struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_cell_area_class_find_cell_property' : return type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_cell_area_class_install_cell_property' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

var cellAreaClassListCellPropertiesFunction *gi.Function
var cellAreaClassListCellPropertiesFunction_Once sync.Once

func cellAreaClassListCellPropertiesFunction_Set() error {
	var err error
	cellAreaClassListCellPropertiesFunction_Once.Do(func() {
		err = cellAreaClassStruct_Set()
		if err != nil {
			return
		}
		cellAreaClassListCellPropertiesFunction, err = cellAreaClassStruct.InvokerNew("list_cell_properties")
	})
	return err
}

// ListCellProperties is a representation of the C type gtk_cell_area_class_list_cell_properties.
func (recv *CellAreaClass) ListCellProperties() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := cellAreaClassListCellPropertiesFunction_Set()
	if err == nil {
		cellAreaClassListCellPropertiesFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
}

var cellAreaContextClassStruct *gi.Struct
var cellAreaContextClassStruct_Once sync.Once

func cellAreaContextClassStruct_Set() error {
	var err error
	cellAreaContextClassStruct_Once.Do(func() {
		cellAreaContextClassStruct, err = gi.StructNew("Gtk", "CellAreaContextClass")
	})
	return err
}

type CellAreaContextClass struct {
	native uintptr
}

var cellAreaContextPrivateStruct *gi.Struct
var cellAreaContextPrivateStruct_Once sync.Once

func cellAreaContextPrivateStruct_Set() error {
	var err error
	cellAreaContextPrivateStruct_Once.Do(func() {
		cellAreaContextPrivateStruct, err = gi.StructNew("Gtk", "CellAreaContextPrivate")
	})
	return err
}

type CellAreaContextPrivate struct {
	native uintptr
}

var cellAreaPrivateStruct *gi.Struct
var cellAreaPrivateStruct_Once sync.Once

func cellAreaPrivateStruct_Set() error {
	var err error
	cellAreaPrivateStruct_Once.Do(func() {
		cellAreaPrivateStruct, err = gi.StructNew("Gtk", "CellAreaPrivate")
	})
	return err
}

type CellAreaPrivate struct {
	native uintptr
}

var cellEditableIfaceStruct *gi.Struct
var cellEditableIfaceStruct_Once sync.Once

func cellEditableIfaceStruct_Set() error {
	var err error
	cellEditableIfaceStruct_Once.Do(func() {
		cellEditableIfaceStruct, err = gi.StructNew("Gtk", "CellEditableIface")
	})
	return err
}

type CellEditableIface struct {
	native uintptr
}

var cellLayoutIfaceStruct *gi.Struct
var cellLayoutIfaceStruct_Once sync.Once

func cellLayoutIfaceStruct_Set() error {
	var err error
	cellLayoutIfaceStruct_Once.Do(func() {
		cellLayoutIfaceStruct, err = gi.StructNew("Gtk", "CellLayoutIface")
	})
	return err
}

type CellLayoutIface struct {
	native uintptr
}

var cellRendererAccelClassStruct *gi.Struct
var cellRendererAccelClassStruct_Once sync.Once

func cellRendererAccelClassStruct_Set() error {
	var err error
	cellRendererAccelClassStruct_Once.Do(func() {
		cellRendererAccelClassStruct, err = gi.StructNew("Gtk", "CellRendererAccelClass")
	})
	return err
}

type CellRendererAccelClass struct {
	native uintptr
}

var cellRendererAccelPrivateStruct *gi.Struct
var cellRendererAccelPrivateStruct_Once sync.Once

func cellRendererAccelPrivateStruct_Set() error {
	var err error
	cellRendererAccelPrivateStruct_Once.Do(func() {
		cellRendererAccelPrivateStruct, err = gi.StructNew("Gtk", "CellRendererAccelPrivate")
	})
	return err
}

type CellRendererAccelPrivate struct {
	native uintptr
}

var cellRendererClassStruct *gi.Struct
var cellRendererClassStruct_Once sync.Once

func cellRendererClassStruct_Set() error {
	var err error
	cellRendererClassStruct_Once.Do(func() {
		cellRendererClassStruct, err = gi.StructNew("Gtk", "CellRendererClass")
	})
	return err
}

type CellRendererClass struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_cell_renderer_class_set_accessible_type' : parameter 'type' of type 'GType' not supported

var cellRendererClassPrivateStruct *gi.Struct
var cellRendererClassPrivateStruct_Once sync.Once

func cellRendererClassPrivateStruct_Set() error {
	var err error
	cellRendererClassPrivateStruct_Once.Do(func() {
		cellRendererClassPrivateStruct, err = gi.StructNew("Gtk", "CellRendererClassPrivate")
	})
	return err
}

type CellRendererClassPrivate struct {
	native uintptr
}

var cellRendererComboClassStruct *gi.Struct
var cellRendererComboClassStruct_Once sync.Once

func cellRendererComboClassStruct_Set() error {
	var err error
	cellRendererComboClassStruct_Once.Do(func() {
		cellRendererComboClassStruct, err = gi.StructNew("Gtk", "CellRendererComboClass")
	})
	return err
}

type CellRendererComboClass struct {
	native uintptr
}

var cellRendererComboPrivateStruct *gi.Struct
var cellRendererComboPrivateStruct_Once sync.Once

func cellRendererComboPrivateStruct_Set() error {
	var err error
	cellRendererComboPrivateStruct_Once.Do(func() {
		cellRendererComboPrivateStruct, err = gi.StructNew("Gtk", "CellRendererComboPrivate")
	})
	return err
}

type CellRendererComboPrivate struct {
	native uintptr
}

var cellRendererPixbufClassStruct *gi.Struct
var cellRendererPixbufClassStruct_Once sync.Once

func cellRendererPixbufClassStruct_Set() error {
	var err error
	cellRendererPixbufClassStruct_Once.Do(func() {
		cellRendererPixbufClassStruct, err = gi.StructNew("Gtk", "CellRendererPixbufClass")
	})
	return err
}

type CellRendererPixbufClass struct {
	native uintptr
}

var cellRendererPixbufPrivateStruct *gi.Struct
var cellRendererPixbufPrivateStruct_Once sync.Once

func cellRendererPixbufPrivateStruct_Set() error {
	var err error
	cellRendererPixbufPrivateStruct_Once.Do(func() {
		cellRendererPixbufPrivateStruct, err = gi.StructNew("Gtk", "CellRendererPixbufPrivate")
	})
	return err
}

type CellRendererPixbufPrivate struct {
	native uintptr
}

var cellRendererPrivateStruct *gi.Struct
var cellRendererPrivateStruct_Once sync.Once

func cellRendererPrivateStruct_Set() error {
	var err error
	cellRendererPrivateStruct_Once.Do(func() {
		cellRendererPrivateStruct, err = gi.StructNew("Gtk", "CellRendererPrivate")
	})
	return err
}

type CellRendererPrivate struct {
	native uintptr
}

var cellRendererProgressClassStruct *gi.Struct
var cellRendererProgressClassStruct_Once sync.Once

func cellRendererProgressClassStruct_Set() error {
	var err error
	cellRendererProgressClassStruct_Once.Do(func() {
		cellRendererProgressClassStruct, err = gi.StructNew("Gtk", "CellRendererProgressClass")
	})
	return err
}

type CellRendererProgressClass struct {
	native uintptr
}

var cellRendererProgressPrivateStruct *gi.Struct
var cellRendererProgressPrivateStruct_Once sync.Once

func cellRendererProgressPrivateStruct_Set() error {
	var err error
	cellRendererProgressPrivateStruct_Once.Do(func() {
		cellRendererProgressPrivateStruct, err = gi.StructNew("Gtk", "CellRendererProgressPrivate")
	})
	return err
}

type CellRendererProgressPrivate struct {
	native uintptr
}

var cellRendererSpinClassStruct *gi.Struct
var cellRendererSpinClassStruct_Once sync.Once

func cellRendererSpinClassStruct_Set() error {
	var err error
	cellRendererSpinClassStruct_Once.Do(func() {
		cellRendererSpinClassStruct, err = gi.StructNew("Gtk", "CellRendererSpinClass")
	})
	return err
}

type CellRendererSpinClass struct {
	native uintptr
}

var cellRendererSpinPrivateStruct *gi.Struct
var cellRendererSpinPrivateStruct_Once sync.Once

func cellRendererSpinPrivateStruct_Set() error {
	var err error
	cellRendererSpinPrivateStruct_Once.Do(func() {
		cellRendererSpinPrivateStruct, err = gi.StructNew("Gtk", "CellRendererSpinPrivate")
	})
	return err
}

type CellRendererSpinPrivate struct {
	native uintptr
}

var cellRendererSpinnerClassStruct *gi.Struct
var cellRendererSpinnerClassStruct_Once sync.Once

func cellRendererSpinnerClassStruct_Set() error {
	var err error
	cellRendererSpinnerClassStruct_Once.Do(func() {
		cellRendererSpinnerClassStruct, err = gi.StructNew("Gtk", "CellRendererSpinnerClass")
	})
	return err
}

type CellRendererSpinnerClass struct {
	native uintptr
}

var cellRendererSpinnerPrivateStruct *gi.Struct
var cellRendererSpinnerPrivateStruct_Once sync.Once

func cellRendererSpinnerPrivateStruct_Set() error {
	var err error
	cellRendererSpinnerPrivateStruct_Once.Do(func() {
		cellRendererSpinnerPrivateStruct, err = gi.StructNew("Gtk", "CellRendererSpinnerPrivate")
	})
	return err
}

type CellRendererSpinnerPrivate struct {
	native uintptr
}

var cellRendererTextClassStruct *gi.Struct
var cellRendererTextClassStruct_Once sync.Once

func cellRendererTextClassStruct_Set() error {
	var err error
	cellRendererTextClassStruct_Once.Do(func() {
		cellRendererTextClassStruct, err = gi.StructNew("Gtk", "CellRendererTextClass")
	})
	return err
}

type CellRendererTextClass struct {
	native uintptr
}

var cellRendererTextPrivateStruct *gi.Struct
var cellRendererTextPrivateStruct_Once sync.Once

func cellRendererTextPrivateStruct_Set() error {
	var err error
	cellRendererTextPrivateStruct_Once.Do(func() {
		cellRendererTextPrivateStruct, err = gi.StructNew("Gtk", "CellRendererTextPrivate")
	})
	return err
}

type CellRendererTextPrivate struct {
	native uintptr
}

var cellRendererToggleClassStruct *gi.Struct
var cellRendererToggleClassStruct_Once sync.Once

func cellRendererToggleClassStruct_Set() error {
	var err error
	cellRendererToggleClassStruct_Once.Do(func() {
		cellRendererToggleClassStruct, err = gi.StructNew("Gtk", "CellRendererToggleClass")
	})
	return err
}

type CellRendererToggleClass struct {
	native uintptr
}

var cellRendererTogglePrivateStruct *gi.Struct
var cellRendererTogglePrivateStruct_Once sync.Once

func cellRendererTogglePrivateStruct_Set() error {
	var err error
	cellRendererTogglePrivateStruct_Once.Do(func() {
		cellRendererTogglePrivateStruct, err = gi.StructNew("Gtk", "CellRendererTogglePrivate")
	})
	return err
}

type CellRendererTogglePrivate struct {
	native uintptr
}

var cellViewClassStruct *gi.Struct
var cellViewClassStruct_Once sync.Once

func cellViewClassStruct_Set() error {
	var err error
	cellViewClassStruct_Once.Do(func() {
		cellViewClassStruct, err = gi.StructNew("Gtk", "CellViewClass")
	})
	return err
}

type CellViewClass struct {
	native uintptr
}

var cellViewPrivateStruct *gi.Struct
var cellViewPrivateStruct_Once sync.Once

func cellViewPrivateStruct_Set() error {
	var err error
	cellViewPrivateStruct_Once.Do(func() {
		cellViewPrivateStruct, err = gi.StructNew("Gtk", "CellViewPrivate")
	})
	return err
}

type CellViewPrivate struct {
	native uintptr
}

var checkButtonClassStruct *gi.Struct
var checkButtonClassStruct_Once sync.Once

func checkButtonClassStruct_Set() error {
	var err error
	checkButtonClassStruct_Once.Do(func() {
		checkButtonClassStruct, err = gi.StructNew("Gtk", "CheckButtonClass")
	})
	return err
}

type CheckButtonClass struct {
	native uintptr
}

var checkMenuItemAccessibleClassStruct *gi.Struct
var checkMenuItemAccessibleClassStruct_Once sync.Once

func checkMenuItemAccessibleClassStruct_Set() error {
	var err error
	checkMenuItemAccessibleClassStruct_Once.Do(func() {
		checkMenuItemAccessibleClassStruct, err = gi.StructNew("Gtk", "CheckMenuItemAccessibleClass")
	})
	return err
}

type CheckMenuItemAccessibleClass struct {
	native uintptr
}

var checkMenuItemAccessiblePrivateStruct *gi.Struct
var checkMenuItemAccessiblePrivateStruct_Once sync.Once

func checkMenuItemAccessiblePrivateStruct_Set() error {
	var err error
	checkMenuItemAccessiblePrivateStruct_Once.Do(func() {
		checkMenuItemAccessiblePrivateStruct, err = gi.StructNew("Gtk", "CheckMenuItemAccessiblePrivate")
	})
	return err
}

type CheckMenuItemAccessiblePrivate struct {
	native uintptr
}

var checkMenuItemClassStruct *gi.Struct
var checkMenuItemClassStruct_Once sync.Once

func checkMenuItemClassStruct_Set() error {
	var err error
	checkMenuItemClassStruct_Once.Do(func() {
		checkMenuItemClassStruct, err = gi.StructNew("Gtk", "CheckMenuItemClass")
	})
	return err
}

type CheckMenuItemClass struct {
	native uintptr
}

var checkMenuItemPrivateStruct *gi.Struct
var checkMenuItemPrivateStruct_Once sync.Once

func checkMenuItemPrivateStruct_Set() error {
	var err error
	checkMenuItemPrivateStruct_Once.Do(func() {
		checkMenuItemPrivateStruct, err = gi.StructNew("Gtk", "CheckMenuItemPrivate")
	})
	return err
}

type CheckMenuItemPrivate struct {
	native uintptr
}

var colorButtonClassStruct *gi.Struct
var colorButtonClassStruct_Once sync.Once

func colorButtonClassStruct_Set() error {
	var err error
	colorButtonClassStruct_Once.Do(func() {
		colorButtonClassStruct, err = gi.StructNew("Gtk", "ColorButtonClass")
	})
	return err
}

type ColorButtonClass struct {
	native uintptr
}

var colorButtonPrivateStruct *gi.Struct
var colorButtonPrivateStruct_Once sync.Once

func colorButtonPrivateStruct_Set() error {
	var err error
	colorButtonPrivateStruct_Once.Do(func() {
		colorButtonPrivateStruct, err = gi.StructNew("Gtk", "ColorButtonPrivate")
	})
	return err
}

type ColorButtonPrivate struct {
	native uintptr
}

var colorChooserDialogClassStruct *gi.Struct
var colorChooserDialogClassStruct_Once sync.Once

func colorChooserDialogClassStruct_Set() error {
	var err error
	colorChooserDialogClassStruct_Once.Do(func() {
		colorChooserDialogClassStruct, err = gi.StructNew("Gtk", "ColorChooserDialogClass")
	})
	return err
}

type ColorChooserDialogClass struct {
	native uintptr
}

var colorChooserDialogPrivateStruct *gi.Struct
var colorChooserDialogPrivateStruct_Once sync.Once

func colorChooserDialogPrivateStruct_Set() error {
	var err error
	colorChooserDialogPrivateStruct_Once.Do(func() {
		colorChooserDialogPrivateStruct, err = gi.StructNew("Gtk", "ColorChooserDialogPrivate")
	})
	return err
}

type ColorChooserDialogPrivate struct {
	native uintptr
}

var colorChooserInterfaceStruct *gi.Struct
var colorChooserInterfaceStruct_Once sync.Once

func colorChooserInterfaceStruct_Set() error {
	var err error
	colorChooserInterfaceStruct_Once.Do(func() {
		colorChooserInterfaceStruct, err = gi.StructNew("Gtk", "ColorChooserInterface")
	})
	return err
}

type ColorChooserInterface struct {
	native uintptr
}

var colorChooserWidgetClassStruct *gi.Struct
var colorChooserWidgetClassStruct_Once sync.Once

func colorChooserWidgetClassStruct_Set() error {
	var err error
	colorChooserWidgetClassStruct_Once.Do(func() {
		colorChooserWidgetClassStruct, err = gi.StructNew("Gtk", "ColorChooserWidgetClass")
	})
	return err
}

type ColorChooserWidgetClass struct {
	native uintptr
}

var colorChooserWidgetPrivateStruct *gi.Struct
var colorChooserWidgetPrivateStruct_Once sync.Once

func colorChooserWidgetPrivateStruct_Set() error {
	var err error
	colorChooserWidgetPrivateStruct_Once.Do(func() {
		colorChooserWidgetPrivateStruct, err = gi.StructNew("Gtk", "ColorChooserWidgetPrivate")
	})
	return err
}

type ColorChooserWidgetPrivate struct {
	native uintptr
}

var colorSelectionClassStruct *gi.Struct
var colorSelectionClassStruct_Once sync.Once

func colorSelectionClassStruct_Set() error {
	var err error
	colorSelectionClassStruct_Once.Do(func() {
		colorSelectionClassStruct, err = gi.StructNew("Gtk", "ColorSelectionClass")
	})
	return err
}

type ColorSelectionClass struct {
	native uintptr
}

var colorSelectionDialogClassStruct *gi.Struct
var colorSelectionDialogClassStruct_Once sync.Once

func colorSelectionDialogClassStruct_Set() error {
	var err error
	colorSelectionDialogClassStruct_Once.Do(func() {
		colorSelectionDialogClassStruct, err = gi.StructNew("Gtk", "ColorSelectionDialogClass")
	})
	return err
}

type ColorSelectionDialogClass struct {
	native uintptr
}

var colorSelectionDialogPrivateStruct *gi.Struct
var colorSelectionDialogPrivateStruct_Once sync.Once

func colorSelectionDialogPrivateStruct_Set() error {
	var err error
	colorSelectionDialogPrivateStruct_Once.Do(func() {
		colorSelectionDialogPrivateStruct, err = gi.StructNew("Gtk", "ColorSelectionDialogPrivate")
	})
	return err
}

type ColorSelectionDialogPrivate struct {
	native uintptr
}

var colorSelectionPrivateStruct *gi.Struct
var colorSelectionPrivateStruct_Once sync.Once

func colorSelectionPrivateStruct_Set() error {
	var err error
	colorSelectionPrivateStruct_Once.Do(func() {
		colorSelectionPrivateStruct, err = gi.StructNew("Gtk", "ColorSelectionPrivate")
	})
	return err
}

type ColorSelectionPrivate struct {
	native uintptr
}

var comboBoxAccessibleClassStruct *gi.Struct
var comboBoxAccessibleClassStruct_Once sync.Once

func comboBoxAccessibleClassStruct_Set() error {
	var err error
	comboBoxAccessibleClassStruct_Once.Do(func() {
		comboBoxAccessibleClassStruct, err = gi.StructNew("Gtk", "ComboBoxAccessibleClass")
	})
	return err
}

type ComboBoxAccessibleClass struct {
	native uintptr
}

var comboBoxAccessiblePrivateStruct *gi.Struct
var comboBoxAccessiblePrivateStruct_Once sync.Once

func comboBoxAccessiblePrivateStruct_Set() error {
	var err error
	comboBoxAccessiblePrivateStruct_Once.Do(func() {
		comboBoxAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ComboBoxAccessiblePrivate")
	})
	return err
}

type ComboBoxAccessiblePrivate struct {
	native uintptr
}

var comboBoxClassStruct *gi.Struct
var comboBoxClassStruct_Once sync.Once

func comboBoxClassStruct_Set() error {
	var err error
	comboBoxClassStruct_Once.Do(func() {
		comboBoxClassStruct, err = gi.StructNew("Gtk", "ComboBoxClass")
	})
	return err
}

type ComboBoxClass struct {
	native uintptr
}

var comboBoxPrivateStruct *gi.Struct
var comboBoxPrivateStruct_Once sync.Once

func comboBoxPrivateStruct_Set() error {
	var err error
	comboBoxPrivateStruct_Once.Do(func() {
		comboBoxPrivateStruct, err = gi.StructNew("Gtk", "ComboBoxPrivate")
	})
	return err
}

type ComboBoxPrivate struct {
	native uintptr
}

var comboBoxTextClassStruct *gi.Struct
var comboBoxTextClassStruct_Once sync.Once

func comboBoxTextClassStruct_Set() error {
	var err error
	comboBoxTextClassStruct_Once.Do(func() {
		comboBoxTextClassStruct, err = gi.StructNew("Gtk", "ComboBoxTextClass")
	})
	return err
}

type ComboBoxTextClass struct {
	native uintptr
}

var comboBoxTextPrivateStruct *gi.Struct
var comboBoxTextPrivateStruct_Once sync.Once

func comboBoxTextPrivateStruct_Set() error {
	var err error
	comboBoxTextPrivateStruct_Once.Do(func() {
		comboBoxTextPrivateStruct, err = gi.StructNew("Gtk", "ComboBoxTextPrivate")
	})
	return err
}

type ComboBoxTextPrivate struct {
	native uintptr
}

var containerAccessibleClassStruct *gi.Struct
var containerAccessibleClassStruct_Once sync.Once

func containerAccessibleClassStruct_Set() error {
	var err error
	containerAccessibleClassStruct_Once.Do(func() {
		containerAccessibleClassStruct, err = gi.StructNew("Gtk", "ContainerAccessibleClass")
	})
	return err
}

type ContainerAccessibleClass struct {
	native uintptr
}

var containerAccessiblePrivateStruct *gi.Struct
var containerAccessiblePrivateStruct_Once sync.Once

func containerAccessiblePrivateStruct_Set() error {
	var err error
	containerAccessiblePrivateStruct_Once.Do(func() {
		containerAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ContainerAccessiblePrivate")
	})
	return err
}

type ContainerAccessiblePrivate struct {
	native uintptr
}

var containerCellAccessibleClassStruct *gi.Struct
var containerCellAccessibleClassStruct_Once sync.Once

func containerCellAccessibleClassStruct_Set() error {
	var err error
	containerCellAccessibleClassStruct_Once.Do(func() {
		containerCellAccessibleClassStruct, err = gi.StructNew("Gtk", "ContainerCellAccessibleClass")
	})
	return err
}

type ContainerCellAccessibleClass struct {
	native uintptr
}

var containerCellAccessiblePrivateStruct *gi.Struct
var containerCellAccessiblePrivateStruct_Once sync.Once

func containerCellAccessiblePrivateStruct_Set() error {
	var err error
	containerCellAccessiblePrivateStruct_Once.Do(func() {
		containerCellAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ContainerCellAccessiblePrivate")
	})
	return err
}

type ContainerCellAccessiblePrivate struct {
	native uintptr
}

var containerClassStruct *gi.Struct
var containerClassStruct_Once sync.Once

func containerClassStruct_Set() error {
	var err error
	containerClassStruct_Once.Do(func() {
		containerClassStruct, err = gi.StructNew("Gtk", "ContainerClass")
	})
	return err
}

type ContainerClass struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_container_class_find_child_property' : return type 'GObject.ParamSpec' not supported

var containerClassHandleBorderWidthFunction *gi.Function
var containerClassHandleBorderWidthFunction_Once sync.Once

func containerClassHandleBorderWidthFunction_Set() error {
	var err error
	containerClassHandleBorderWidthFunction_Once.Do(func() {
		err = containerClassStruct_Set()
		if err != nil {
			return
		}
		containerClassHandleBorderWidthFunction, err = containerClassStruct.InvokerNew("handle_border_width")
	})
	return err
}

// HandleBorderWidth is a representation of the C type gtk_container_class_handle_border_width.
func (recv *ContainerClass) HandleBorderWidth() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := containerClassHandleBorderWidthFunction_Set()
	if err == nil {
		containerClassHandleBorderWidthFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_container_class_install_child_properties' : parameter 'pspecs' has no type

// UNSUPPORTED : C value 'gtk_container_class_install_child_property' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

var containerClassListChildPropertiesFunction *gi.Function
var containerClassListChildPropertiesFunction_Once sync.Once

func containerClassListChildPropertiesFunction_Set() error {
	var err error
	containerClassListChildPropertiesFunction_Once.Do(func() {
		err = containerClassStruct_Set()
		if err != nil {
			return
		}
		containerClassListChildPropertiesFunction, err = containerClassStruct.InvokerNew("list_child_properties")
	})
	return err
}

// ListChildProperties is a representation of the C type gtk_container_class_list_child_properties.
func (recv *ContainerClass) ListChildProperties() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := containerClassListChildPropertiesFunction_Set()
	if err == nil {
		containerClassListChildPropertiesFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
}

var containerPrivateStruct *gi.Struct
var containerPrivateStruct_Once sync.Once

func containerPrivateStruct_Set() error {
	var err error
	containerPrivateStruct_Once.Do(func() {
		containerPrivateStruct, err = gi.StructNew("Gtk", "ContainerPrivate")
	})
	return err
}

type ContainerPrivate struct {
	native uintptr
}

var cssProviderClassStruct *gi.Struct
var cssProviderClassStruct_Once sync.Once

func cssProviderClassStruct_Set() error {
	var err error
	cssProviderClassStruct_Once.Do(func() {
		cssProviderClassStruct, err = gi.StructNew("Gtk", "CssProviderClass")
	})
	return err
}

type CssProviderClass struct {
	native uintptr
}

var cssProviderPrivateStruct *gi.Struct
var cssProviderPrivateStruct_Once sync.Once

func cssProviderPrivateStruct_Set() error {
	var err error
	cssProviderPrivateStruct_Once.Do(func() {
		cssProviderPrivateStruct, err = gi.StructNew("Gtk", "CssProviderPrivate")
	})
	return err
}

type CssProviderPrivate struct {
	native uintptr
}

var cssSectionStruct *gi.Struct
var cssSectionStruct_Once sync.Once

func cssSectionStruct_Set() error {
	var err error
	cssSectionStruct_Once.Do(func() {
		cssSectionStruct, err = gi.StructNew("Gtk", "CssSection")
	})
	return err
}

type CssSection struct {
	native uintptr
}

var cssSectionGetEndLineFunction *gi.Function
var cssSectionGetEndLineFunction_Once sync.Once

func cssSectionGetEndLineFunction_Set() error {
	var err error
	cssSectionGetEndLineFunction_Once.Do(func() {
		err = cssSectionStruct_Set()
		if err != nil {
			return
		}
		cssSectionGetEndLineFunction, err = cssSectionStruct.InvokerNew("get_end_line")
	})
	return err
}

// GetEndLine is a representation of the C type gtk_css_section_get_end_line.
func (recv *CssSection) GetEndLine() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cssSectionGetEndLineFunction_Set()
	if err == nil {
		ret = cssSectionGetEndLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var cssSectionGetEndPositionFunction *gi.Function
var cssSectionGetEndPositionFunction_Once sync.Once

func cssSectionGetEndPositionFunction_Set() error {
	var err error
	cssSectionGetEndPositionFunction_Once.Do(func() {
		err = cssSectionStruct_Set()
		if err != nil {
			return
		}
		cssSectionGetEndPositionFunction, err = cssSectionStruct.InvokerNew("get_end_position")
	})
	return err
}

// GetEndPosition is a representation of the C type gtk_css_section_get_end_position.
func (recv *CssSection) GetEndPosition() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cssSectionGetEndPositionFunction_Set()
	if err == nil {
		ret = cssSectionGetEndPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_css_section_get_file' : return type 'Gio.File' not supported

var cssSectionGetParentFunction *gi.Function
var cssSectionGetParentFunction_Once sync.Once

func cssSectionGetParentFunction_Set() error {
	var err error
	cssSectionGetParentFunction_Once.Do(func() {
		err = cssSectionStruct_Set()
		if err != nil {
			return
		}
		cssSectionGetParentFunction, err = cssSectionStruct.InvokerNew("get_parent")
	})
	return err
}

// GetParent is a representation of the C type gtk_css_section_get_parent.
func (recv *CssSection) GetParent() *CssSection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cssSectionGetParentFunction_Set()
	if err == nil {
		ret = cssSectionGetParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := &CssSection{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_css_section_get_section_type' : return type 'CssSectionType' not supported

var cssSectionGetStartLineFunction *gi.Function
var cssSectionGetStartLineFunction_Once sync.Once

func cssSectionGetStartLineFunction_Set() error {
	var err error
	cssSectionGetStartLineFunction_Once.Do(func() {
		err = cssSectionStruct_Set()
		if err != nil {
			return
		}
		cssSectionGetStartLineFunction, err = cssSectionStruct.InvokerNew("get_start_line")
	})
	return err
}

// GetStartLine is a representation of the C type gtk_css_section_get_start_line.
func (recv *CssSection) GetStartLine() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cssSectionGetStartLineFunction_Set()
	if err == nil {
		ret = cssSectionGetStartLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var cssSectionGetStartPositionFunction *gi.Function
var cssSectionGetStartPositionFunction_Once sync.Once

func cssSectionGetStartPositionFunction_Set() error {
	var err error
	cssSectionGetStartPositionFunction_Once.Do(func() {
		err = cssSectionStruct_Set()
		if err != nil {
			return
		}
		cssSectionGetStartPositionFunction, err = cssSectionStruct.InvokerNew("get_start_position")
	})
	return err
}

// GetStartPosition is a representation of the C type gtk_css_section_get_start_position.
func (recv *CssSection) GetStartPosition() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cssSectionGetStartPositionFunction_Set()
	if err == nil {
		ret = cssSectionGetStartPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var cssSectionRefFunction *gi.Function
var cssSectionRefFunction_Once sync.Once

func cssSectionRefFunction_Set() error {
	var err error
	cssSectionRefFunction_Once.Do(func() {
		err = cssSectionStruct_Set()
		if err != nil {
			return
		}
		cssSectionRefFunction, err = cssSectionStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gtk_css_section_ref.
func (recv *CssSection) Ref() *CssSection {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := cssSectionRefFunction_Set()
	if err == nil {
		ret = cssSectionRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &CssSection{native: ret.Pointer()}

	return retGo
}

var cssSectionUnrefFunction *gi.Function
var cssSectionUnrefFunction_Once sync.Once

func cssSectionUnrefFunction_Set() error {
	var err error
	cssSectionUnrefFunction_Once.Do(func() {
		err = cssSectionStruct_Set()
		if err != nil {
			return
		}
		cssSectionUnrefFunction, err = cssSectionStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gtk_css_section_unref.
func (recv *CssSection) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := cssSectionUnrefFunction_Set()
	if err == nil {
		cssSectionUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dialogClassStruct *gi.Struct
var dialogClassStruct_Once sync.Once

func dialogClassStruct_Set() error {
	var err error
	dialogClassStruct_Once.Do(func() {
		dialogClassStruct, err = gi.StructNew("Gtk", "DialogClass")
	})
	return err
}

type DialogClass struct {
	native uintptr
}

var dialogPrivateStruct *gi.Struct
var dialogPrivateStruct_Once sync.Once

func dialogPrivateStruct_Set() error {
	var err error
	dialogPrivateStruct_Once.Do(func() {
		dialogPrivateStruct, err = gi.StructNew("Gtk", "DialogPrivate")
	})
	return err
}

type DialogPrivate struct {
	native uintptr
}

var drawingAreaClassStruct *gi.Struct
var drawingAreaClassStruct_Once sync.Once

func drawingAreaClassStruct_Set() error {
	var err error
	drawingAreaClassStruct_Once.Do(func() {
		drawingAreaClassStruct, err = gi.StructNew("Gtk", "DrawingAreaClass")
	})
	return err
}

type DrawingAreaClass struct {
	native uintptr
}

var editableInterfaceStruct *gi.Struct
var editableInterfaceStruct_Once sync.Once

func editableInterfaceStruct_Set() error {
	var err error
	editableInterfaceStruct_Once.Do(func() {
		editableInterfaceStruct, err = gi.StructNew("Gtk", "EditableInterface")
	})
	return err
}

type EditableInterface struct {
	native uintptr
}

var entryAccessibleClassStruct *gi.Struct
var entryAccessibleClassStruct_Once sync.Once

func entryAccessibleClassStruct_Set() error {
	var err error
	entryAccessibleClassStruct_Once.Do(func() {
		entryAccessibleClassStruct, err = gi.StructNew("Gtk", "EntryAccessibleClass")
	})
	return err
}

type EntryAccessibleClass struct {
	native uintptr
}

var entryAccessiblePrivateStruct *gi.Struct
var entryAccessiblePrivateStruct_Once sync.Once

func entryAccessiblePrivateStruct_Set() error {
	var err error
	entryAccessiblePrivateStruct_Once.Do(func() {
		entryAccessiblePrivateStruct, err = gi.StructNew("Gtk", "EntryAccessiblePrivate")
	})
	return err
}

type EntryAccessiblePrivate struct {
	native uintptr
}

var entryBufferClassStruct *gi.Struct
var entryBufferClassStruct_Once sync.Once

func entryBufferClassStruct_Set() error {
	var err error
	entryBufferClassStruct_Once.Do(func() {
		entryBufferClassStruct, err = gi.StructNew("Gtk", "EntryBufferClass")
	})
	return err
}

type EntryBufferClass struct {
	native uintptr
}

var entryBufferPrivateStruct *gi.Struct
var entryBufferPrivateStruct_Once sync.Once

func entryBufferPrivateStruct_Set() error {
	var err error
	entryBufferPrivateStruct_Once.Do(func() {
		entryBufferPrivateStruct, err = gi.StructNew("Gtk", "EntryBufferPrivate")
	})
	return err
}

type EntryBufferPrivate struct {
	native uintptr
}

var entryClassStruct *gi.Struct
var entryClassStruct_Once sync.Once

func entryClassStruct_Set() error {
	var err error
	entryClassStruct_Once.Do(func() {
		entryClassStruct, err = gi.StructNew("Gtk", "EntryClass")
	})
	return err
}

type EntryClass struct {
	native uintptr
}

var entryCompletionClassStruct *gi.Struct
var entryCompletionClassStruct_Once sync.Once

func entryCompletionClassStruct_Set() error {
	var err error
	entryCompletionClassStruct_Once.Do(func() {
		entryCompletionClassStruct, err = gi.StructNew("Gtk", "EntryCompletionClass")
	})
	return err
}

type EntryCompletionClass struct {
	native uintptr
}

var entryCompletionPrivateStruct *gi.Struct
var entryCompletionPrivateStruct_Once sync.Once

func entryCompletionPrivateStruct_Set() error {
	var err error
	entryCompletionPrivateStruct_Once.Do(func() {
		entryCompletionPrivateStruct, err = gi.StructNew("Gtk", "EntryCompletionPrivate")
	})
	return err
}

type EntryCompletionPrivate struct {
	native uintptr
}

var entryPrivateStruct *gi.Struct
var entryPrivateStruct_Once sync.Once

func entryPrivateStruct_Set() error {
	var err error
	entryPrivateStruct_Once.Do(func() {
		entryPrivateStruct, err = gi.StructNew("Gtk", "EntryPrivate")
	})
	return err
}

type EntryPrivate struct {
	native uintptr
}

var eventBoxClassStruct *gi.Struct
var eventBoxClassStruct_Once sync.Once

func eventBoxClassStruct_Set() error {
	var err error
	eventBoxClassStruct_Once.Do(func() {
		eventBoxClassStruct, err = gi.StructNew("Gtk", "EventBoxClass")
	})
	return err
}

type EventBoxClass struct {
	native uintptr
}

var eventBoxPrivateStruct *gi.Struct
var eventBoxPrivateStruct_Once sync.Once

func eventBoxPrivateStruct_Set() error {
	var err error
	eventBoxPrivateStruct_Once.Do(func() {
		eventBoxPrivateStruct, err = gi.StructNew("Gtk", "EventBoxPrivate")
	})
	return err
}

type EventBoxPrivate struct {
	native uintptr
}

var eventControllerClassStruct *gi.Struct
var eventControllerClassStruct_Once sync.Once

func eventControllerClassStruct_Set() error {
	var err error
	eventControllerClassStruct_Once.Do(func() {
		eventControllerClassStruct, err = gi.StructNew("Gtk", "EventControllerClass")
	})
	return err
}

type EventControllerClass struct {
	native uintptr
}

var eventControllerKeyClassStruct *gi.Struct
var eventControllerKeyClassStruct_Once sync.Once

func eventControllerKeyClassStruct_Set() error {
	var err error
	eventControllerKeyClassStruct_Once.Do(func() {
		eventControllerKeyClassStruct, err = gi.StructNew("Gtk", "EventControllerKeyClass")
	})
	return err
}

type EventControllerKeyClass struct {
	native uintptr
}

var eventControllerMotionClassStruct *gi.Struct
var eventControllerMotionClassStruct_Once sync.Once

func eventControllerMotionClassStruct_Set() error {
	var err error
	eventControllerMotionClassStruct_Once.Do(func() {
		eventControllerMotionClassStruct, err = gi.StructNew("Gtk", "EventControllerMotionClass")
	})
	return err
}

type EventControllerMotionClass struct {
	native uintptr
}

var eventControllerScrollClassStruct *gi.Struct
var eventControllerScrollClassStruct_Once sync.Once

func eventControllerScrollClassStruct_Set() error {
	var err error
	eventControllerScrollClassStruct_Once.Do(func() {
		eventControllerScrollClassStruct, err = gi.StructNew("Gtk", "EventControllerScrollClass")
	})
	return err
}

type EventControllerScrollClass struct {
	native uintptr
}

var expanderAccessibleClassStruct *gi.Struct
var expanderAccessibleClassStruct_Once sync.Once

func expanderAccessibleClassStruct_Set() error {
	var err error
	expanderAccessibleClassStruct_Once.Do(func() {
		expanderAccessibleClassStruct, err = gi.StructNew("Gtk", "ExpanderAccessibleClass")
	})
	return err
}

type ExpanderAccessibleClass struct {
	native uintptr
}

var expanderAccessiblePrivateStruct *gi.Struct
var expanderAccessiblePrivateStruct_Once sync.Once

func expanderAccessiblePrivateStruct_Set() error {
	var err error
	expanderAccessiblePrivateStruct_Once.Do(func() {
		expanderAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ExpanderAccessiblePrivate")
	})
	return err
}

type ExpanderAccessiblePrivate struct {
	native uintptr
}

var expanderClassStruct *gi.Struct
var expanderClassStruct_Once sync.Once

func expanderClassStruct_Set() error {
	var err error
	expanderClassStruct_Once.Do(func() {
		expanderClassStruct, err = gi.StructNew("Gtk", "ExpanderClass")
	})
	return err
}

type ExpanderClass struct {
	native uintptr
}

var expanderPrivateStruct *gi.Struct
var expanderPrivateStruct_Once sync.Once

func expanderPrivateStruct_Set() error {
	var err error
	expanderPrivateStruct_Once.Do(func() {
		expanderPrivateStruct, err = gi.StructNew("Gtk", "ExpanderPrivate")
	})
	return err
}

type ExpanderPrivate struct {
	native uintptr
}

var fileChooserButtonClassStruct *gi.Struct
var fileChooserButtonClassStruct_Once sync.Once

func fileChooserButtonClassStruct_Set() error {
	var err error
	fileChooserButtonClassStruct_Once.Do(func() {
		fileChooserButtonClassStruct, err = gi.StructNew("Gtk", "FileChooserButtonClass")
	})
	return err
}

type FileChooserButtonClass struct {
	native uintptr
}

var fileChooserButtonPrivateStruct *gi.Struct
var fileChooserButtonPrivateStruct_Once sync.Once

func fileChooserButtonPrivateStruct_Set() error {
	var err error
	fileChooserButtonPrivateStruct_Once.Do(func() {
		fileChooserButtonPrivateStruct, err = gi.StructNew("Gtk", "FileChooserButtonPrivate")
	})
	return err
}

type FileChooserButtonPrivate struct {
	native uintptr
}

var fileChooserDialogClassStruct *gi.Struct
var fileChooserDialogClassStruct_Once sync.Once

func fileChooserDialogClassStruct_Set() error {
	var err error
	fileChooserDialogClassStruct_Once.Do(func() {
		fileChooserDialogClassStruct, err = gi.StructNew("Gtk", "FileChooserDialogClass")
	})
	return err
}

type FileChooserDialogClass struct {
	native uintptr
}

var fileChooserDialogPrivateStruct *gi.Struct
var fileChooserDialogPrivateStruct_Once sync.Once

func fileChooserDialogPrivateStruct_Set() error {
	var err error
	fileChooserDialogPrivateStruct_Once.Do(func() {
		fileChooserDialogPrivateStruct, err = gi.StructNew("Gtk", "FileChooserDialogPrivate")
	})
	return err
}

type FileChooserDialogPrivate struct {
	native uintptr
}

var fileChooserNativeClassStruct *gi.Struct
var fileChooserNativeClassStruct_Once sync.Once

func fileChooserNativeClassStruct_Set() error {
	var err error
	fileChooserNativeClassStruct_Once.Do(func() {
		fileChooserNativeClassStruct, err = gi.StructNew("Gtk", "FileChooserNativeClass")
	})
	return err
}

type FileChooserNativeClass struct {
	native uintptr
}

var fileChooserWidgetClassStruct *gi.Struct
var fileChooserWidgetClassStruct_Once sync.Once

func fileChooserWidgetClassStruct_Set() error {
	var err error
	fileChooserWidgetClassStruct_Once.Do(func() {
		fileChooserWidgetClassStruct, err = gi.StructNew("Gtk", "FileChooserWidgetClass")
	})
	return err
}

type FileChooserWidgetClass struct {
	native uintptr
}

var fileChooserWidgetPrivateStruct *gi.Struct
var fileChooserWidgetPrivateStruct_Once sync.Once

func fileChooserWidgetPrivateStruct_Set() error {
	var err error
	fileChooserWidgetPrivateStruct_Once.Do(func() {
		fileChooserWidgetPrivateStruct, err = gi.StructNew("Gtk", "FileChooserWidgetPrivate")
	})
	return err
}

type FileChooserWidgetPrivate struct {
	native uintptr
}

var fileFilterInfoStruct *gi.Struct
var fileFilterInfoStruct_Once sync.Once

func fileFilterInfoStruct_Set() error {
	var err error
	fileFilterInfoStruct_Once.Do(func() {
		fileFilterInfoStruct, err = gi.StructNew("Gtk", "FileFilterInfo")
	})
	return err
}

type FileFilterInfo struct {
	native uintptr
}

var fixedChildStruct *gi.Struct
var fixedChildStruct_Once sync.Once

func fixedChildStruct_Set() error {
	var err error
	fixedChildStruct_Once.Do(func() {
		fixedChildStruct, err = gi.StructNew("Gtk", "FixedChild")
	})
	return err
}

type FixedChild struct {
	native uintptr
}

var fixedClassStruct *gi.Struct
var fixedClassStruct_Once sync.Once

func fixedClassStruct_Set() error {
	var err error
	fixedClassStruct_Once.Do(func() {
		fixedClassStruct, err = gi.StructNew("Gtk", "FixedClass")
	})
	return err
}

type FixedClass struct {
	native uintptr
}

var fixedPrivateStruct *gi.Struct
var fixedPrivateStruct_Once sync.Once

func fixedPrivateStruct_Set() error {
	var err error
	fixedPrivateStruct_Once.Do(func() {
		fixedPrivateStruct, err = gi.StructNew("Gtk", "FixedPrivate")
	})
	return err
}

type FixedPrivate struct {
	native uintptr
}

var flowBoxAccessibleClassStruct *gi.Struct
var flowBoxAccessibleClassStruct_Once sync.Once

func flowBoxAccessibleClassStruct_Set() error {
	var err error
	flowBoxAccessibleClassStruct_Once.Do(func() {
		flowBoxAccessibleClassStruct, err = gi.StructNew("Gtk", "FlowBoxAccessibleClass")
	})
	return err
}

type FlowBoxAccessibleClass struct {
	native uintptr
}

var flowBoxAccessiblePrivateStruct *gi.Struct
var flowBoxAccessiblePrivateStruct_Once sync.Once

func flowBoxAccessiblePrivateStruct_Set() error {
	var err error
	flowBoxAccessiblePrivateStruct_Once.Do(func() {
		flowBoxAccessiblePrivateStruct, err = gi.StructNew("Gtk", "FlowBoxAccessiblePrivate")
	})
	return err
}

type FlowBoxAccessiblePrivate struct {
	native uintptr
}

var flowBoxChildAccessibleClassStruct *gi.Struct
var flowBoxChildAccessibleClassStruct_Once sync.Once

func flowBoxChildAccessibleClassStruct_Set() error {
	var err error
	flowBoxChildAccessibleClassStruct_Once.Do(func() {
		flowBoxChildAccessibleClassStruct, err = gi.StructNew("Gtk", "FlowBoxChildAccessibleClass")
	})
	return err
}

type FlowBoxChildAccessibleClass struct {
	native uintptr
}

var flowBoxChildClassStruct *gi.Struct
var flowBoxChildClassStruct_Once sync.Once

func flowBoxChildClassStruct_Set() error {
	var err error
	flowBoxChildClassStruct_Once.Do(func() {
		flowBoxChildClassStruct, err = gi.StructNew("Gtk", "FlowBoxChildClass")
	})
	return err
}

type FlowBoxChildClass struct {
	native uintptr
}

var flowBoxClassStruct *gi.Struct
var flowBoxClassStruct_Once sync.Once

func flowBoxClassStruct_Set() error {
	var err error
	flowBoxClassStruct_Once.Do(func() {
		flowBoxClassStruct, err = gi.StructNew("Gtk", "FlowBoxClass")
	})
	return err
}

type FlowBoxClass struct {
	native uintptr
}

var fontButtonClassStruct *gi.Struct
var fontButtonClassStruct_Once sync.Once

func fontButtonClassStruct_Set() error {
	var err error
	fontButtonClassStruct_Once.Do(func() {
		fontButtonClassStruct, err = gi.StructNew("Gtk", "FontButtonClass")
	})
	return err
}

type FontButtonClass struct {
	native uintptr
}

var fontButtonPrivateStruct *gi.Struct
var fontButtonPrivateStruct_Once sync.Once

func fontButtonPrivateStruct_Set() error {
	var err error
	fontButtonPrivateStruct_Once.Do(func() {
		fontButtonPrivateStruct, err = gi.StructNew("Gtk", "FontButtonPrivate")
	})
	return err
}

type FontButtonPrivate struct {
	native uintptr
}

var fontChooserDialogClassStruct *gi.Struct
var fontChooserDialogClassStruct_Once sync.Once

func fontChooserDialogClassStruct_Set() error {
	var err error
	fontChooserDialogClassStruct_Once.Do(func() {
		fontChooserDialogClassStruct, err = gi.StructNew("Gtk", "FontChooserDialogClass")
	})
	return err
}

type FontChooserDialogClass struct {
	native uintptr
}

var fontChooserDialogPrivateStruct *gi.Struct
var fontChooserDialogPrivateStruct_Once sync.Once

func fontChooserDialogPrivateStruct_Set() error {
	var err error
	fontChooserDialogPrivateStruct_Once.Do(func() {
		fontChooserDialogPrivateStruct, err = gi.StructNew("Gtk", "FontChooserDialogPrivate")
	})
	return err
}

type FontChooserDialogPrivate struct {
	native uintptr
}

var fontChooserIfaceStruct *gi.Struct
var fontChooserIfaceStruct_Once sync.Once

func fontChooserIfaceStruct_Set() error {
	var err error
	fontChooserIfaceStruct_Once.Do(func() {
		fontChooserIfaceStruct, err = gi.StructNew("Gtk", "FontChooserIface")
	})
	return err
}

type FontChooserIface struct {
	native uintptr
}

var fontChooserWidgetClassStruct *gi.Struct
var fontChooserWidgetClassStruct_Once sync.Once

func fontChooserWidgetClassStruct_Set() error {
	var err error
	fontChooserWidgetClassStruct_Once.Do(func() {
		fontChooserWidgetClassStruct, err = gi.StructNew("Gtk", "FontChooserWidgetClass")
	})
	return err
}

type FontChooserWidgetClass struct {
	native uintptr
}

var fontChooserWidgetPrivateStruct *gi.Struct
var fontChooserWidgetPrivateStruct_Once sync.Once

func fontChooserWidgetPrivateStruct_Set() error {
	var err error
	fontChooserWidgetPrivateStruct_Once.Do(func() {
		fontChooserWidgetPrivateStruct, err = gi.StructNew("Gtk", "FontChooserWidgetPrivate")
	})
	return err
}

type FontChooserWidgetPrivate struct {
	native uintptr
}

var fontSelectionClassStruct *gi.Struct
var fontSelectionClassStruct_Once sync.Once

func fontSelectionClassStruct_Set() error {
	var err error
	fontSelectionClassStruct_Once.Do(func() {
		fontSelectionClassStruct, err = gi.StructNew("Gtk", "FontSelectionClass")
	})
	return err
}

type FontSelectionClass struct {
	native uintptr
}

var fontSelectionDialogClassStruct *gi.Struct
var fontSelectionDialogClassStruct_Once sync.Once

func fontSelectionDialogClassStruct_Set() error {
	var err error
	fontSelectionDialogClassStruct_Once.Do(func() {
		fontSelectionDialogClassStruct, err = gi.StructNew("Gtk", "FontSelectionDialogClass")
	})
	return err
}

type FontSelectionDialogClass struct {
	native uintptr
}

var fontSelectionDialogPrivateStruct *gi.Struct
var fontSelectionDialogPrivateStruct_Once sync.Once

func fontSelectionDialogPrivateStruct_Set() error {
	var err error
	fontSelectionDialogPrivateStruct_Once.Do(func() {
		fontSelectionDialogPrivateStruct, err = gi.StructNew("Gtk", "FontSelectionDialogPrivate")
	})
	return err
}

type FontSelectionDialogPrivate struct {
	native uintptr
}

var fontSelectionPrivateStruct *gi.Struct
var fontSelectionPrivateStruct_Once sync.Once

func fontSelectionPrivateStruct_Set() error {
	var err error
	fontSelectionPrivateStruct_Once.Do(func() {
		fontSelectionPrivateStruct, err = gi.StructNew("Gtk", "FontSelectionPrivate")
	})
	return err
}

type FontSelectionPrivate struct {
	native uintptr
}

var frameAccessibleClassStruct *gi.Struct
var frameAccessibleClassStruct_Once sync.Once

func frameAccessibleClassStruct_Set() error {
	var err error
	frameAccessibleClassStruct_Once.Do(func() {
		frameAccessibleClassStruct, err = gi.StructNew("Gtk", "FrameAccessibleClass")
	})
	return err
}

type FrameAccessibleClass struct {
	native uintptr
}

var frameAccessiblePrivateStruct *gi.Struct
var frameAccessiblePrivateStruct_Once sync.Once

func frameAccessiblePrivateStruct_Set() error {
	var err error
	frameAccessiblePrivateStruct_Once.Do(func() {
		frameAccessiblePrivateStruct, err = gi.StructNew("Gtk", "FrameAccessiblePrivate")
	})
	return err
}

type FrameAccessiblePrivate struct {
	native uintptr
}

var frameClassStruct *gi.Struct
var frameClassStruct_Once sync.Once

func frameClassStruct_Set() error {
	var err error
	frameClassStruct_Once.Do(func() {
		frameClassStruct, err = gi.StructNew("Gtk", "FrameClass")
	})
	return err
}

type FrameClass struct {
	native uintptr
}

var framePrivateStruct *gi.Struct
var framePrivateStruct_Once sync.Once

func framePrivateStruct_Set() error {
	var err error
	framePrivateStruct_Once.Do(func() {
		framePrivateStruct, err = gi.StructNew("Gtk", "FramePrivate")
	})
	return err
}

type FramePrivate struct {
	native uintptr
}

var gLAreaClassStruct *gi.Struct
var gLAreaClassStruct_Once sync.Once

func gLAreaClassStruct_Set() error {
	var err error
	gLAreaClassStruct_Once.Do(func() {
		gLAreaClassStruct, err = gi.StructNew("Gtk", "GLAreaClass")
	})
	return err
}

type GLAreaClass struct {
	native uintptr
}

var gestureClassStruct *gi.Struct
var gestureClassStruct_Once sync.Once

func gestureClassStruct_Set() error {
	var err error
	gestureClassStruct_Once.Do(func() {
		gestureClassStruct, err = gi.StructNew("Gtk", "GestureClass")
	})
	return err
}

type GestureClass struct {
	native uintptr
}

var gestureDragClassStruct *gi.Struct
var gestureDragClassStruct_Once sync.Once

func gestureDragClassStruct_Set() error {
	var err error
	gestureDragClassStruct_Once.Do(func() {
		gestureDragClassStruct, err = gi.StructNew("Gtk", "GestureDragClass")
	})
	return err
}

type GestureDragClass struct {
	native uintptr
}

var gestureLongPressClassStruct *gi.Struct
var gestureLongPressClassStruct_Once sync.Once

func gestureLongPressClassStruct_Set() error {
	var err error
	gestureLongPressClassStruct_Once.Do(func() {
		gestureLongPressClassStruct, err = gi.StructNew("Gtk", "GestureLongPressClass")
	})
	return err
}

type GestureLongPressClass struct {
	native uintptr
}

var gestureMultiPressClassStruct *gi.Struct
var gestureMultiPressClassStruct_Once sync.Once

func gestureMultiPressClassStruct_Set() error {
	var err error
	gestureMultiPressClassStruct_Once.Do(func() {
		gestureMultiPressClassStruct, err = gi.StructNew("Gtk", "GestureMultiPressClass")
	})
	return err
}

type GestureMultiPressClass struct {
	native uintptr
}

var gesturePanClassStruct *gi.Struct
var gesturePanClassStruct_Once sync.Once

func gesturePanClassStruct_Set() error {
	var err error
	gesturePanClassStruct_Once.Do(func() {
		gesturePanClassStruct, err = gi.StructNew("Gtk", "GesturePanClass")
	})
	return err
}

type GesturePanClass struct {
	native uintptr
}

var gestureRotateClassStruct *gi.Struct
var gestureRotateClassStruct_Once sync.Once

func gestureRotateClassStruct_Set() error {
	var err error
	gestureRotateClassStruct_Once.Do(func() {
		gestureRotateClassStruct, err = gi.StructNew("Gtk", "GestureRotateClass")
	})
	return err
}

type GestureRotateClass struct {
	native uintptr
}

var gestureSingleClassStruct *gi.Struct
var gestureSingleClassStruct_Once sync.Once

func gestureSingleClassStruct_Set() error {
	var err error
	gestureSingleClassStruct_Once.Do(func() {
		gestureSingleClassStruct, err = gi.StructNew("Gtk", "GestureSingleClass")
	})
	return err
}

type GestureSingleClass struct {
	native uintptr
}

var gestureStylusClassStruct *gi.Struct
var gestureStylusClassStruct_Once sync.Once

func gestureStylusClassStruct_Set() error {
	var err error
	gestureStylusClassStruct_Once.Do(func() {
		gestureStylusClassStruct, err = gi.StructNew("Gtk", "GestureStylusClass")
	})
	return err
}

type GestureStylusClass struct {
	native uintptr
}

var gestureSwipeClassStruct *gi.Struct
var gestureSwipeClassStruct_Once sync.Once

func gestureSwipeClassStruct_Set() error {
	var err error
	gestureSwipeClassStruct_Once.Do(func() {
		gestureSwipeClassStruct, err = gi.StructNew("Gtk", "GestureSwipeClass")
	})
	return err
}

type GestureSwipeClass struct {
	native uintptr
}

var gestureZoomClassStruct *gi.Struct
var gestureZoomClassStruct_Once sync.Once

func gestureZoomClassStruct_Set() error {
	var err error
	gestureZoomClassStruct_Once.Do(func() {
		gestureZoomClassStruct, err = gi.StructNew("Gtk", "GestureZoomClass")
	})
	return err
}

type GestureZoomClass struct {
	native uintptr
}

var gradientStruct *gi.Struct
var gradientStruct_Once sync.Once

func gradientStruct_Set() error {
	var err error
	gradientStruct_Once.Do(func() {
		gradientStruct, err = gi.StructNew("Gtk", "Gradient")
	})
	return err
}

type Gradient struct {
	native uintptr
}

var gradientNewLinearFunction *gi.Function
var gradientNewLinearFunction_Once sync.Once

func gradientNewLinearFunction_Set() error {
	var err error
	gradientNewLinearFunction_Once.Do(func() {
		err = gradientStruct_Set()
		if err != nil {
			return
		}
		gradientNewLinearFunction, err = gradientStruct.InvokerNew("new_linear")
	})
	return err
}

// GradientNewLinear is a representation of the C type gtk_gradient_new_linear.
func GradientNewLinear(x0 float64, y0 float64, x1 float64, y1 float64) *Gradient {
	var inArgs [4]gi.Argument
	inArgs[0].SetFloat64(x0)
	inArgs[1].SetFloat64(y0)
	inArgs[2].SetFloat64(x1)
	inArgs[3].SetFloat64(y1)

	var ret gi.Argument

	err := gradientNewLinearFunction_Set()
	if err == nil {
		ret = gradientNewLinearFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Gradient{native: ret.Pointer()}

	return retGo
}

var gradientNewRadialFunction *gi.Function
var gradientNewRadialFunction_Once sync.Once

func gradientNewRadialFunction_Set() error {
	var err error
	gradientNewRadialFunction_Once.Do(func() {
		err = gradientStruct_Set()
		if err != nil {
			return
		}
		gradientNewRadialFunction, err = gradientStruct.InvokerNew("new_radial")
	})
	return err
}

// GradientNewRadial is a representation of the C type gtk_gradient_new_radial.
func GradientNewRadial(x0 float64, y0 float64, radius0 float64, x1 float64, y1 float64, radius1 float64) *Gradient {
	var inArgs [6]gi.Argument
	inArgs[0].SetFloat64(x0)
	inArgs[1].SetFloat64(y0)
	inArgs[2].SetFloat64(radius0)
	inArgs[3].SetFloat64(x1)
	inArgs[4].SetFloat64(y1)
	inArgs[5].SetFloat64(radius1)

	var ret gi.Argument

	err := gradientNewRadialFunction_Set()
	if err == nil {
		ret = gradientNewRadialFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Gradient{native: ret.Pointer()}

	return retGo
}

var gradientAddColorStopFunction *gi.Function
var gradientAddColorStopFunction_Once sync.Once

func gradientAddColorStopFunction_Set() error {
	var err error
	gradientAddColorStopFunction_Once.Do(func() {
		err = gradientStruct_Set()
		if err != nil {
			return
		}
		gradientAddColorStopFunction, err = gradientStruct.InvokerNew("add_color_stop")
	})
	return err
}

// AddColorStop is a representation of the C type gtk_gradient_add_color_stop.
func (recv *Gradient) AddColorStop(offset float64, color *SymbolicColor) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(offset)
	inArgs[2].SetPointer(color.native)

	err := gradientAddColorStopFunction_Set()
	if err == nil {
		gradientAddColorStopFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gradientRefFunction *gi.Function
var gradientRefFunction_Once sync.Once

func gradientRefFunction_Set() error {
	var err error
	gradientRefFunction_Once.Do(func() {
		err = gradientStruct_Set()
		if err != nil {
			return
		}
		gradientRefFunction, err = gradientStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gtk_gradient_ref.
func (recv *Gradient) Ref() *Gradient {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := gradientRefFunction_Set()
	if err == nil {
		ret = gradientRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Gradient{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_gradient_resolve' : parameter 'props' of type 'StyleProperties' not supported

// UNSUPPORTED : C value 'gtk_gradient_resolve_for_context' : parameter 'context' of type 'StyleContext' not supported

var gradientToStringFunction *gi.Function
var gradientToStringFunction_Once sync.Once

func gradientToStringFunction_Set() error {
	var err error
	gradientToStringFunction_Once.Do(func() {
		err = gradientStruct_Set()
		if err != nil {
			return
		}
		gradientToStringFunction, err = gradientStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type gtk_gradient_to_string.
func (recv *Gradient) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := gradientToStringFunction_Set()
	if err == nil {
		ret = gradientToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var gradientUnrefFunction *gi.Function
var gradientUnrefFunction_Once sync.Once

func gradientUnrefFunction_Set() error {
	var err error
	gradientUnrefFunction_Once.Do(func() {
		err = gradientStruct_Set()
		if err != nil {
			return
		}
		gradientUnrefFunction, err = gradientStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gtk_gradient_unref.
func (recv *Gradient) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := gradientUnrefFunction_Set()
	if err == nil {
		gradientUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gridClassStruct *gi.Struct
var gridClassStruct_Once sync.Once

func gridClassStruct_Set() error {
	var err error
	gridClassStruct_Once.Do(func() {
		gridClassStruct, err = gi.StructNew("Gtk", "GridClass")
	})
	return err
}

type GridClass struct {
	native uintptr
}

var gridPrivateStruct *gi.Struct
var gridPrivateStruct_Once sync.Once

func gridPrivateStruct_Set() error {
	var err error
	gridPrivateStruct_Once.Do(func() {
		gridPrivateStruct, err = gi.StructNew("Gtk", "GridPrivate")
	})
	return err
}

type GridPrivate struct {
	native uintptr
}

var hBoxClassStruct *gi.Struct
var hBoxClassStruct_Once sync.Once

func hBoxClassStruct_Set() error {
	var err error
	hBoxClassStruct_Once.Do(func() {
		hBoxClassStruct, err = gi.StructNew("Gtk", "HBoxClass")
	})
	return err
}

type HBoxClass struct {
	native uintptr
}

var hButtonBoxClassStruct *gi.Struct
var hButtonBoxClassStruct_Once sync.Once

func hButtonBoxClassStruct_Set() error {
	var err error
	hButtonBoxClassStruct_Once.Do(func() {
		hButtonBoxClassStruct, err = gi.StructNew("Gtk", "HButtonBoxClass")
	})
	return err
}

type HButtonBoxClass struct {
	native uintptr
}

var hPanedClassStruct *gi.Struct
var hPanedClassStruct_Once sync.Once

func hPanedClassStruct_Set() error {
	var err error
	hPanedClassStruct_Once.Do(func() {
		hPanedClassStruct, err = gi.StructNew("Gtk", "HPanedClass")
	})
	return err
}

type HPanedClass struct {
	native uintptr
}

var hSVClassStruct *gi.Struct
var hSVClassStruct_Once sync.Once

func hSVClassStruct_Set() error {
	var err error
	hSVClassStruct_Once.Do(func() {
		hSVClassStruct, err = gi.StructNew("Gtk", "HSVClass")
	})
	return err
}

type HSVClass struct {
	native uintptr
}

var hSVPrivateStruct *gi.Struct
var hSVPrivateStruct_Once sync.Once

func hSVPrivateStruct_Set() error {
	var err error
	hSVPrivateStruct_Once.Do(func() {
		hSVPrivateStruct, err = gi.StructNew("Gtk", "HSVPrivate")
	})
	return err
}

type HSVPrivate struct {
	native uintptr
}

var hScaleClassStruct *gi.Struct
var hScaleClassStruct_Once sync.Once

func hScaleClassStruct_Set() error {
	var err error
	hScaleClassStruct_Once.Do(func() {
		hScaleClassStruct, err = gi.StructNew("Gtk", "HScaleClass")
	})
	return err
}

type HScaleClass struct {
	native uintptr
}

var hScrollbarClassStruct *gi.Struct
var hScrollbarClassStruct_Once sync.Once

func hScrollbarClassStruct_Set() error {
	var err error
	hScrollbarClassStruct_Once.Do(func() {
		hScrollbarClassStruct, err = gi.StructNew("Gtk", "HScrollbarClass")
	})
	return err
}

type HScrollbarClass struct {
	native uintptr
}

var hSeparatorClassStruct *gi.Struct
var hSeparatorClassStruct_Once sync.Once

func hSeparatorClassStruct_Set() error {
	var err error
	hSeparatorClassStruct_Once.Do(func() {
		hSeparatorClassStruct, err = gi.StructNew("Gtk", "HSeparatorClass")
	})
	return err
}

type HSeparatorClass struct {
	native uintptr
}

var handleBoxClassStruct *gi.Struct
var handleBoxClassStruct_Once sync.Once

func handleBoxClassStruct_Set() error {
	var err error
	handleBoxClassStruct_Once.Do(func() {
		handleBoxClassStruct, err = gi.StructNew("Gtk", "HandleBoxClass")
	})
	return err
}

type HandleBoxClass struct {
	native uintptr
}

var handleBoxPrivateStruct *gi.Struct
var handleBoxPrivateStruct_Once sync.Once

func handleBoxPrivateStruct_Set() error {
	var err error
	handleBoxPrivateStruct_Once.Do(func() {
		handleBoxPrivateStruct, err = gi.StructNew("Gtk", "HandleBoxPrivate")
	})
	return err
}

type HandleBoxPrivate struct {
	native uintptr
}

var headerBarAccessibleClassStruct *gi.Struct
var headerBarAccessibleClassStruct_Once sync.Once

func headerBarAccessibleClassStruct_Set() error {
	var err error
	headerBarAccessibleClassStruct_Once.Do(func() {
		headerBarAccessibleClassStruct, err = gi.StructNew("Gtk", "HeaderBarAccessibleClass")
	})
	return err
}

type HeaderBarAccessibleClass struct {
	native uintptr
}

var headerBarAccessiblePrivateStruct *gi.Struct
var headerBarAccessiblePrivateStruct_Once sync.Once

func headerBarAccessiblePrivateStruct_Set() error {
	var err error
	headerBarAccessiblePrivateStruct_Once.Do(func() {
		headerBarAccessiblePrivateStruct, err = gi.StructNew("Gtk", "HeaderBarAccessiblePrivate")
	})
	return err
}

type HeaderBarAccessiblePrivate struct {
	native uintptr
}

var headerBarClassStruct *gi.Struct
var headerBarClassStruct_Once sync.Once

func headerBarClassStruct_Set() error {
	var err error
	headerBarClassStruct_Once.Do(func() {
		headerBarClassStruct, err = gi.StructNew("Gtk", "HeaderBarClass")
	})
	return err
}

type HeaderBarClass struct {
	native uintptr
}

var headerBarPrivateStruct *gi.Struct
var headerBarPrivateStruct_Once sync.Once

func headerBarPrivateStruct_Set() error {
	var err error
	headerBarPrivateStruct_Once.Do(func() {
		headerBarPrivateStruct, err = gi.StructNew("Gtk", "HeaderBarPrivate")
	})
	return err
}

type HeaderBarPrivate struct {
	native uintptr
}

var iMContextClassStruct *gi.Struct
var iMContextClassStruct_Once sync.Once

func iMContextClassStruct_Set() error {
	var err error
	iMContextClassStruct_Once.Do(func() {
		iMContextClassStruct, err = gi.StructNew("Gtk", "IMContextClass")
	})
	return err
}

type IMContextClass struct {
	native uintptr
}

var iMContextInfoStruct *gi.Struct
var iMContextInfoStruct_Once sync.Once

func iMContextInfoStruct_Set() error {
	var err error
	iMContextInfoStruct_Once.Do(func() {
		iMContextInfoStruct, err = gi.StructNew("Gtk", "IMContextInfo")
	})
	return err
}

type IMContextInfo struct {
	native uintptr
}

var iMContextSimpleClassStruct *gi.Struct
var iMContextSimpleClassStruct_Once sync.Once

func iMContextSimpleClassStruct_Set() error {
	var err error
	iMContextSimpleClassStruct_Once.Do(func() {
		iMContextSimpleClassStruct, err = gi.StructNew("Gtk", "IMContextSimpleClass")
	})
	return err
}

type IMContextSimpleClass struct {
	native uintptr
}

var iMContextSimplePrivateStruct *gi.Struct
var iMContextSimplePrivateStruct_Once sync.Once

func iMContextSimplePrivateStruct_Set() error {
	var err error
	iMContextSimplePrivateStruct_Once.Do(func() {
		iMContextSimplePrivateStruct, err = gi.StructNew("Gtk", "IMContextSimplePrivate")
	})
	return err
}

type IMContextSimplePrivate struct {
	native uintptr
}

var iMMulticontextClassStruct *gi.Struct
var iMMulticontextClassStruct_Once sync.Once

func iMMulticontextClassStruct_Set() error {
	var err error
	iMMulticontextClassStruct_Once.Do(func() {
		iMMulticontextClassStruct, err = gi.StructNew("Gtk", "IMMulticontextClass")
	})
	return err
}

type IMMulticontextClass struct {
	native uintptr
}

var iMMulticontextPrivateStruct *gi.Struct
var iMMulticontextPrivateStruct_Once sync.Once

func iMMulticontextPrivateStruct_Set() error {
	var err error
	iMMulticontextPrivateStruct_Once.Do(func() {
		iMMulticontextPrivateStruct, err = gi.StructNew("Gtk", "IMMulticontextPrivate")
	})
	return err
}

type IMMulticontextPrivate struct {
	native uintptr
}

var iconFactoryClassStruct *gi.Struct
var iconFactoryClassStruct_Once sync.Once

func iconFactoryClassStruct_Set() error {
	var err error
	iconFactoryClassStruct_Once.Do(func() {
		iconFactoryClassStruct, err = gi.StructNew("Gtk", "IconFactoryClass")
	})
	return err
}

type IconFactoryClass struct {
	native uintptr
}

var iconFactoryPrivateStruct *gi.Struct
var iconFactoryPrivateStruct_Once sync.Once

func iconFactoryPrivateStruct_Set() error {
	var err error
	iconFactoryPrivateStruct_Once.Do(func() {
		iconFactoryPrivateStruct, err = gi.StructNew("Gtk", "IconFactoryPrivate")
	})
	return err
}

type IconFactoryPrivate struct {
	native uintptr
}

var iconInfoClassStruct *gi.Struct
var iconInfoClassStruct_Once sync.Once

func iconInfoClassStruct_Set() error {
	var err error
	iconInfoClassStruct_Once.Do(func() {
		iconInfoClassStruct, err = gi.StructNew("Gtk", "IconInfoClass")
	})
	return err
}

type IconInfoClass struct {
	native uintptr
}

var iconSetStruct *gi.Struct
var iconSetStruct_Once sync.Once

func iconSetStruct_Set() error {
	var err error
	iconSetStruct_Once.Do(func() {
		iconSetStruct, err = gi.StructNew("Gtk", "IconSet")
	})
	return err
}

type IconSet struct {
	native uintptr
}

var iconSetNewFunction *gi.Function
var iconSetNewFunction_Once sync.Once

func iconSetNewFunction_Set() error {
	var err error
	iconSetNewFunction_Once.Do(func() {
		err = iconSetStruct_Set()
		if err != nil {
			return
		}
		iconSetNewFunction, err = iconSetStruct.InvokerNew("new")
	})
	return err
}

// IconSetNew is a representation of the C type gtk_icon_set_new.
func IconSetNew() *IconSet {

	var ret gi.Argument

	err := iconSetNewFunction_Set()
	if err == nil {
		ret = iconSetNewFunction.Invoke(nil, nil)
	}

	retGo := &IconSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_set_new_from_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

var iconSetAddSourceFunction *gi.Function
var iconSetAddSourceFunction_Once sync.Once

func iconSetAddSourceFunction_Set() error {
	var err error
	iconSetAddSourceFunction_Once.Do(func() {
		err = iconSetStruct_Set()
		if err != nil {
			return
		}
		iconSetAddSourceFunction, err = iconSetStruct.InvokerNew("add_source")
	})
	return err
}

// AddSource is a representation of the C type gtk_icon_set_add_source.
func (recv *IconSet) AddSource(source *IconSource) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(source.native)

	err := iconSetAddSourceFunction_Set()
	if err == nil {
		iconSetAddSourceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iconSetCopyFunction *gi.Function
var iconSetCopyFunction_Once sync.Once

func iconSetCopyFunction_Set() error {
	var err error
	iconSetCopyFunction_Once.Do(func() {
		err = iconSetStruct_Set()
		if err != nil {
			return
		}
		iconSetCopyFunction, err = iconSetStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_icon_set_copy.
func (recv *IconSet) Copy() *IconSet {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iconSetCopyFunction_Set()
	if err == nil {
		ret = iconSetCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IconSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_set_get_sizes' : parameter 'sizes' has no type

var iconSetRefFunction *gi.Function
var iconSetRefFunction_Once sync.Once

func iconSetRefFunction_Set() error {
	var err error
	iconSetRefFunction_Once.Do(func() {
		err = iconSetStruct_Set()
		if err != nil {
			return
		}
		iconSetRefFunction, err = iconSetStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gtk_icon_set_ref.
func (recv *IconSet) Ref() *IconSet {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iconSetRefFunction_Set()
	if err == nil {
		ret = iconSetRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IconSet{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_set_render_icon' : parameter 'style' of type 'Style' not supported

// UNSUPPORTED : C value 'gtk_icon_set_render_icon_pixbuf' : parameter 'context' of type 'StyleContext' not supported

// UNSUPPORTED : C value 'gtk_icon_set_render_icon_surface' : parameter 'context' of type 'StyleContext' not supported

var iconSetUnrefFunction *gi.Function
var iconSetUnrefFunction_Once sync.Once

func iconSetUnrefFunction_Set() error {
	var err error
	iconSetUnrefFunction_Once.Do(func() {
		err = iconSetStruct_Set()
		if err != nil {
			return
		}
		iconSetUnrefFunction, err = iconSetStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gtk_icon_set_unref.
func (recv *IconSet) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := iconSetUnrefFunction_Set()
	if err == nil {
		iconSetUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iconSourceStruct *gi.Struct
var iconSourceStruct_Once sync.Once

func iconSourceStruct_Set() error {
	var err error
	iconSourceStruct_Once.Do(func() {
		iconSourceStruct, err = gi.StructNew("Gtk", "IconSource")
	})
	return err
}

type IconSource struct {
	native uintptr
}

var iconSourceNewFunction *gi.Function
var iconSourceNewFunction_Once sync.Once

func iconSourceNewFunction_Set() error {
	var err error
	iconSourceNewFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceNewFunction, err = iconSourceStruct.InvokerNew("new")
	})
	return err
}

// IconSourceNew is a representation of the C type gtk_icon_source_new.
func IconSourceNew() *IconSource {

	var ret gi.Argument

	err := iconSourceNewFunction_Set()
	if err == nil {
		ret = iconSourceNewFunction.Invoke(nil, nil)
	}

	retGo := &IconSource{native: ret.Pointer()}

	return retGo
}

var iconSourceCopyFunction *gi.Function
var iconSourceCopyFunction_Once sync.Once

func iconSourceCopyFunction_Set() error {
	var err error
	iconSourceCopyFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceCopyFunction, err = iconSourceStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_icon_source_copy.
func (recv *IconSource) Copy() *IconSource {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iconSourceCopyFunction_Set()
	if err == nil {
		ret = iconSourceCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &IconSource{native: ret.Pointer()}

	return retGo
}

var iconSourceFreeFunction *gi.Function
var iconSourceFreeFunction_Once sync.Once

func iconSourceFreeFunction_Set() error {
	var err error
	iconSourceFreeFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceFreeFunction, err = iconSourceStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_icon_source_free.
func (recv *IconSource) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := iconSourceFreeFunction_Set()
	if err == nil {
		iconSourceFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_icon_source_get_direction' : return type 'TextDirection' not supported

var iconSourceGetDirectionWildcardedFunction *gi.Function
var iconSourceGetDirectionWildcardedFunction_Once sync.Once

func iconSourceGetDirectionWildcardedFunction_Set() error {
	var err error
	iconSourceGetDirectionWildcardedFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceGetDirectionWildcardedFunction, err = iconSourceStruct.InvokerNew("get_direction_wildcarded")
	})
	return err
}

// GetDirectionWildcarded is a representation of the C type gtk_icon_source_get_direction_wildcarded.
func (recv *IconSource) GetDirectionWildcarded() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iconSourceGetDirectionWildcardedFunction_Set()
	if err == nil {
		ret = iconSourceGetDirectionWildcardedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var iconSourceGetFilenameFunction *gi.Function
var iconSourceGetFilenameFunction_Once sync.Once

func iconSourceGetFilenameFunction_Set() error {
	var err error
	iconSourceGetFilenameFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceGetFilenameFunction, err = iconSourceStruct.InvokerNew("get_filename")
	})
	return err
}

// GetFilename is a representation of the C type gtk_icon_source_get_filename.
func (recv *IconSource) GetFilename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iconSourceGetFilenameFunction_Set()
	if err == nil {
		ret = iconSourceGetFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var iconSourceGetIconNameFunction *gi.Function
var iconSourceGetIconNameFunction_Once sync.Once

func iconSourceGetIconNameFunction_Set() error {
	var err error
	iconSourceGetIconNameFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceGetIconNameFunction, err = iconSourceStruct.InvokerNew("get_icon_name")
	})
	return err
}

// GetIconName is a representation of the C type gtk_icon_source_get_icon_name.
func (recv *IconSource) GetIconName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iconSourceGetIconNameFunction_Set()
	if err == nil {
		ret = iconSourceGetIconNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_source_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_icon_source_get_size' : return type 'IconSize' not supported

var iconSourceGetSizeWildcardedFunction *gi.Function
var iconSourceGetSizeWildcardedFunction_Once sync.Once

func iconSourceGetSizeWildcardedFunction_Set() error {
	var err error
	iconSourceGetSizeWildcardedFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceGetSizeWildcardedFunction, err = iconSourceStruct.InvokerNew("get_size_wildcarded")
	})
	return err
}

// GetSizeWildcarded is a representation of the C type gtk_icon_source_get_size_wildcarded.
func (recv *IconSource) GetSizeWildcarded() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iconSourceGetSizeWildcardedFunction_Set()
	if err == nil {
		ret = iconSourceGetSizeWildcardedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_source_get_state' : return type 'StateType' not supported

var iconSourceGetStateWildcardedFunction *gi.Function
var iconSourceGetStateWildcardedFunction_Once sync.Once

func iconSourceGetStateWildcardedFunction_Set() error {
	var err error
	iconSourceGetStateWildcardedFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceGetStateWildcardedFunction, err = iconSourceStruct.InvokerNew("get_state_wildcarded")
	})
	return err
}

// GetStateWildcarded is a representation of the C type gtk_icon_source_get_state_wildcarded.
func (recv *IconSource) GetStateWildcarded() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := iconSourceGetStateWildcardedFunction_Set()
	if err == nil {
		ret = iconSourceGetStateWildcardedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_icon_source_set_direction' : parameter 'direction' of type 'TextDirection' not supported

var iconSourceSetDirectionWildcardedFunction *gi.Function
var iconSourceSetDirectionWildcardedFunction_Once sync.Once

func iconSourceSetDirectionWildcardedFunction_Set() error {
	var err error
	iconSourceSetDirectionWildcardedFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceSetDirectionWildcardedFunction, err = iconSourceStruct.InvokerNew("set_direction_wildcarded")
	})
	return err
}

// SetDirectionWildcarded is a representation of the C type gtk_icon_source_set_direction_wildcarded.
func (recv *IconSource) SetDirectionWildcarded(setting bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(setting)

	err := iconSourceSetDirectionWildcardedFunction_Set()
	if err == nil {
		iconSourceSetDirectionWildcardedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iconSourceSetFilenameFunction *gi.Function
var iconSourceSetFilenameFunction_Once sync.Once

func iconSourceSetFilenameFunction_Set() error {
	var err error
	iconSourceSetFilenameFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceSetFilenameFunction, err = iconSourceStruct.InvokerNew("set_filename")
	})
	return err
}

// SetFilename is a representation of the C type gtk_icon_source_set_filename.
func (recv *IconSource) SetFilename(filename string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(filename)

	err := iconSourceSetFilenameFunction_Set()
	if err == nil {
		iconSourceSetFilenameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iconSourceSetIconNameFunction *gi.Function
var iconSourceSetIconNameFunction_Once sync.Once

func iconSourceSetIconNameFunction_Set() error {
	var err error
	iconSourceSetIconNameFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceSetIconNameFunction, err = iconSourceStruct.InvokerNew("set_icon_name")
	})
	return err
}

// SetIconName is a representation of the C type gtk_icon_source_set_icon_name.
func (recv *IconSource) SetIconName(iconName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(iconName)

	err := iconSourceSetIconNameFunction_Set()
	if err == nil {
		iconSourceSetIconNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_icon_source_set_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_icon_source_set_size' : parameter 'size' of type 'IconSize' not supported

var iconSourceSetSizeWildcardedFunction *gi.Function
var iconSourceSetSizeWildcardedFunction_Once sync.Once

func iconSourceSetSizeWildcardedFunction_Set() error {
	var err error
	iconSourceSetSizeWildcardedFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceSetSizeWildcardedFunction, err = iconSourceStruct.InvokerNew("set_size_wildcarded")
	})
	return err
}

// SetSizeWildcarded is a representation of the C type gtk_icon_source_set_size_wildcarded.
func (recv *IconSource) SetSizeWildcarded(setting bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(setting)

	err := iconSourceSetSizeWildcardedFunction_Set()
	if err == nil {
		iconSourceSetSizeWildcardedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_icon_source_set_state' : parameter 'state' of type 'StateType' not supported

var iconSourceSetStateWildcardedFunction *gi.Function
var iconSourceSetStateWildcardedFunction_Once sync.Once

func iconSourceSetStateWildcardedFunction_Set() error {
	var err error
	iconSourceSetStateWildcardedFunction_Once.Do(func() {
		err = iconSourceStruct_Set()
		if err != nil {
			return
		}
		iconSourceSetStateWildcardedFunction, err = iconSourceStruct.InvokerNew("set_state_wildcarded")
	})
	return err
}

// SetStateWildcarded is a representation of the C type gtk_icon_source_set_state_wildcarded.
func (recv *IconSource) SetStateWildcarded(setting bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(setting)

	err := iconSourceSetStateWildcardedFunction_Set()
	if err == nil {
		iconSourceSetStateWildcardedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var iconThemeClassStruct *gi.Struct
var iconThemeClassStruct_Once sync.Once

func iconThemeClassStruct_Set() error {
	var err error
	iconThemeClassStruct_Once.Do(func() {
		iconThemeClassStruct, err = gi.StructNew("Gtk", "IconThemeClass")
	})
	return err
}

type IconThemeClass struct {
	native uintptr
}

var iconThemePrivateStruct *gi.Struct
var iconThemePrivateStruct_Once sync.Once

func iconThemePrivateStruct_Set() error {
	var err error
	iconThemePrivateStruct_Once.Do(func() {
		iconThemePrivateStruct, err = gi.StructNew("Gtk", "IconThemePrivate")
	})
	return err
}

type IconThemePrivate struct {
	native uintptr
}

var iconViewAccessibleClassStruct *gi.Struct
var iconViewAccessibleClassStruct_Once sync.Once

func iconViewAccessibleClassStruct_Set() error {
	var err error
	iconViewAccessibleClassStruct_Once.Do(func() {
		iconViewAccessibleClassStruct, err = gi.StructNew("Gtk", "IconViewAccessibleClass")
	})
	return err
}

type IconViewAccessibleClass struct {
	native uintptr
}

var iconViewAccessiblePrivateStruct *gi.Struct
var iconViewAccessiblePrivateStruct_Once sync.Once

func iconViewAccessiblePrivateStruct_Set() error {
	var err error
	iconViewAccessiblePrivateStruct_Once.Do(func() {
		iconViewAccessiblePrivateStruct, err = gi.StructNew("Gtk", "IconViewAccessiblePrivate")
	})
	return err
}

type IconViewAccessiblePrivate struct {
	native uintptr
}

var iconViewClassStruct *gi.Struct
var iconViewClassStruct_Once sync.Once

func iconViewClassStruct_Set() error {
	var err error
	iconViewClassStruct_Once.Do(func() {
		iconViewClassStruct, err = gi.StructNew("Gtk", "IconViewClass")
	})
	return err
}

type IconViewClass struct {
	native uintptr
}

var iconViewPrivateStruct *gi.Struct
var iconViewPrivateStruct_Once sync.Once

func iconViewPrivateStruct_Set() error {
	var err error
	iconViewPrivateStruct_Once.Do(func() {
		iconViewPrivateStruct, err = gi.StructNew("Gtk", "IconViewPrivate")
	})
	return err
}

type IconViewPrivate struct {
	native uintptr
}

var imageAccessibleClassStruct *gi.Struct
var imageAccessibleClassStruct_Once sync.Once

func imageAccessibleClassStruct_Set() error {
	var err error
	imageAccessibleClassStruct_Once.Do(func() {
		imageAccessibleClassStruct, err = gi.StructNew("Gtk", "ImageAccessibleClass")
	})
	return err
}

type ImageAccessibleClass struct {
	native uintptr
}

var imageAccessiblePrivateStruct *gi.Struct
var imageAccessiblePrivateStruct_Once sync.Once

func imageAccessiblePrivateStruct_Set() error {
	var err error
	imageAccessiblePrivateStruct_Once.Do(func() {
		imageAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ImageAccessiblePrivate")
	})
	return err
}

type ImageAccessiblePrivate struct {
	native uintptr
}

var imageCellAccessibleClassStruct *gi.Struct
var imageCellAccessibleClassStruct_Once sync.Once

func imageCellAccessibleClassStruct_Set() error {
	var err error
	imageCellAccessibleClassStruct_Once.Do(func() {
		imageCellAccessibleClassStruct, err = gi.StructNew("Gtk", "ImageCellAccessibleClass")
	})
	return err
}

type ImageCellAccessibleClass struct {
	native uintptr
}

var imageCellAccessiblePrivateStruct *gi.Struct
var imageCellAccessiblePrivateStruct_Once sync.Once

func imageCellAccessiblePrivateStruct_Set() error {
	var err error
	imageCellAccessiblePrivateStruct_Once.Do(func() {
		imageCellAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ImageCellAccessiblePrivate")
	})
	return err
}

type ImageCellAccessiblePrivate struct {
	native uintptr
}

var imageClassStruct *gi.Struct
var imageClassStruct_Once sync.Once

func imageClassStruct_Set() error {
	var err error
	imageClassStruct_Once.Do(func() {
		imageClassStruct, err = gi.StructNew("Gtk", "ImageClass")
	})
	return err
}

type ImageClass struct {
	native uintptr
}

var imageMenuItemClassStruct *gi.Struct
var imageMenuItemClassStruct_Once sync.Once

func imageMenuItemClassStruct_Set() error {
	var err error
	imageMenuItemClassStruct_Once.Do(func() {
		imageMenuItemClassStruct, err = gi.StructNew("Gtk", "ImageMenuItemClass")
	})
	return err
}

type ImageMenuItemClass struct {
	native uintptr
}

var imageMenuItemPrivateStruct *gi.Struct
var imageMenuItemPrivateStruct_Once sync.Once

func imageMenuItemPrivateStruct_Set() error {
	var err error
	imageMenuItemPrivateStruct_Once.Do(func() {
		imageMenuItemPrivateStruct, err = gi.StructNew("Gtk", "ImageMenuItemPrivate")
	})
	return err
}

type ImageMenuItemPrivate struct {
	native uintptr
}

var imagePrivateStruct *gi.Struct
var imagePrivateStruct_Once sync.Once

func imagePrivateStruct_Set() error {
	var err error
	imagePrivateStruct_Once.Do(func() {
		imagePrivateStruct, err = gi.StructNew("Gtk", "ImagePrivate")
	})
	return err
}

type ImagePrivate struct {
	native uintptr
}

var infoBarClassStruct *gi.Struct
var infoBarClassStruct_Once sync.Once

func infoBarClassStruct_Set() error {
	var err error
	infoBarClassStruct_Once.Do(func() {
		infoBarClassStruct, err = gi.StructNew("Gtk", "InfoBarClass")
	})
	return err
}

type InfoBarClass struct {
	native uintptr
}

var infoBarPrivateStruct *gi.Struct
var infoBarPrivateStruct_Once sync.Once

func infoBarPrivateStruct_Set() error {
	var err error
	infoBarPrivateStruct_Once.Do(func() {
		infoBarPrivateStruct, err = gi.StructNew("Gtk", "InfoBarPrivate")
	})
	return err
}

type InfoBarPrivate struct {
	native uintptr
}

var invisibleClassStruct *gi.Struct
var invisibleClassStruct_Once sync.Once

func invisibleClassStruct_Set() error {
	var err error
	invisibleClassStruct_Once.Do(func() {
		invisibleClassStruct, err = gi.StructNew("Gtk", "InvisibleClass")
	})
	return err
}

type InvisibleClass struct {
	native uintptr
}

var invisiblePrivateStruct *gi.Struct
var invisiblePrivateStruct_Once sync.Once

func invisiblePrivateStruct_Set() error {
	var err error
	invisiblePrivateStruct_Once.Do(func() {
		invisiblePrivateStruct, err = gi.StructNew("Gtk", "InvisiblePrivate")
	})
	return err
}

type InvisiblePrivate struct {
	native uintptr
}

var labelAccessibleClassStruct *gi.Struct
var labelAccessibleClassStruct_Once sync.Once

func labelAccessibleClassStruct_Set() error {
	var err error
	labelAccessibleClassStruct_Once.Do(func() {
		labelAccessibleClassStruct, err = gi.StructNew("Gtk", "LabelAccessibleClass")
	})
	return err
}

type LabelAccessibleClass struct {
	native uintptr
}

var labelAccessiblePrivateStruct *gi.Struct
var labelAccessiblePrivateStruct_Once sync.Once

func labelAccessiblePrivateStruct_Set() error {
	var err error
	labelAccessiblePrivateStruct_Once.Do(func() {
		labelAccessiblePrivateStruct, err = gi.StructNew("Gtk", "LabelAccessiblePrivate")
	})
	return err
}

type LabelAccessiblePrivate struct {
	native uintptr
}

var labelClassStruct *gi.Struct
var labelClassStruct_Once sync.Once

func labelClassStruct_Set() error {
	var err error
	labelClassStruct_Once.Do(func() {
		labelClassStruct, err = gi.StructNew("Gtk", "LabelClass")
	})
	return err
}

type LabelClass struct {
	native uintptr
}

var labelPrivateStruct *gi.Struct
var labelPrivateStruct_Once sync.Once

func labelPrivateStruct_Set() error {
	var err error
	labelPrivateStruct_Once.Do(func() {
		labelPrivateStruct, err = gi.StructNew("Gtk", "LabelPrivate")
	})
	return err
}

type LabelPrivate struct {
	native uintptr
}

var labelSelectionInfoStruct *gi.Struct
var labelSelectionInfoStruct_Once sync.Once

func labelSelectionInfoStruct_Set() error {
	var err error
	labelSelectionInfoStruct_Once.Do(func() {
		labelSelectionInfoStruct, err = gi.StructNew("Gtk", "LabelSelectionInfo")
	})
	return err
}

type LabelSelectionInfo struct {
	native uintptr
}

var layoutClassStruct *gi.Struct
var layoutClassStruct_Once sync.Once

func layoutClassStruct_Set() error {
	var err error
	layoutClassStruct_Once.Do(func() {
		layoutClassStruct, err = gi.StructNew("Gtk", "LayoutClass")
	})
	return err
}

type LayoutClass struct {
	native uintptr
}

var layoutPrivateStruct *gi.Struct
var layoutPrivateStruct_Once sync.Once

func layoutPrivateStruct_Set() error {
	var err error
	layoutPrivateStruct_Once.Do(func() {
		layoutPrivateStruct, err = gi.StructNew("Gtk", "LayoutPrivate")
	})
	return err
}

type LayoutPrivate struct {
	native uintptr
}

var levelBarAccessibleClassStruct *gi.Struct
var levelBarAccessibleClassStruct_Once sync.Once

func levelBarAccessibleClassStruct_Set() error {
	var err error
	levelBarAccessibleClassStruct_Once.Do(func() {
		levelBarAccessibleClassStruct, err = gi.StructNew("Gtk", "LevelBarAccessibleClass")
	})
	return err
}

type LevelBarAccessibleClass struct {
	native uintptr
}

var levelBarAccessiblePrivateStruct *gi.Struct
var levelBarAccessiblePrivateStruct_Once sync.Once

func levelBarAccessiblePrivateStruct_Set() error {
	var err error
	levelBarAccessiblePrivateStruct_Once.Do(func() {
		levelBarAccessiblePrivateStruct, err = gi.StructNew("Gtk", "LevelBarAccessiblePrivate")
	})
	return err
}

type LevelBarAccessiblePrivate struct {
	native uintptr
}

var levelBarClassStruct *gi.Struct
var levelBarClassStruct_Once sync.Once

func levelBarClassStruct_Set() error {
	var err error
	levelBarClassStruct_Once.Do(func() {
		levelBarClassStruct, err = gi.StructNew("Gtk", "LevelBarClass")
	})
	return err
}

type LevelBarClass struct {
	native uintptr
}

var levelBarPrivateStruct *gi.Struct
var levelBarPrivateStruct_Once sync.Once

func levelBarPrivateStruct_Set() error {
	var err error
	levelBarPrivateStruct_Once.Do(func() {
		levelBarPrivateStruct, err = gi.StructNew("Gtk", "LevelBarPrivate")
	})
	return err
}

type LevelBarPrivate struct {
	native uintptr
}

var linkButtonAccessibleClassStruct *gi.Struct
var linkButtonAccessibleClassStruct_Once sync.Once

func linkButtonAccessibleClassStruct_Set() error {
	var err error
	linkButtonAccessibleClassStruct_Once.Do(func() {
		linkButtonAccessibleClassStruct, err = gi.StructNew("Gtk", "LinkButtonAccessibleClass")
	})
	return err
}

type LinkButtonAccessibleClass struct {
	native uintptr
}

var linkButtonAccessiblePrivateStruct *gi.Struct
var linkButtonAccessiblePrivateStruct_Once sync.Once

func linkButtonAccessiblePrivateStruct_Set() error {
	var err error
	linkButtonAccessiblePrivateStruct_Once.Do(func() {
		linkButtonAccessiblePrivateStruct, err = gi.StructNew("Gtk", "LinkButtonAccessiblePrivate")
	})
	return err
}

type LinkButtonAccessiblePrivate struct {
	native uintptr
}

var linkButtonClassStruct *gi.Struct
var linkButtonClassStruct_Once sync.Once

func linkButtonClassStruct_Set() error {
	var err error
	linkButtonClassStruct_Once.Do(func() {
		linkButtonClassStruct, err = gi.StructNew("Gtk", "LinkButtonClass")
	})
	return err
}

type LinkButtonClass struct {
	native uintptr
}

var linkButtonPrivateStruct *gi.Struct
var linkButtonPrivateStruct_Once sync.Once

func linkButtonPrivateStruct_Set() error {
	var err error
	linkButtonPrivateStruct_Once.Do(func() {
		linkButtonPrivateStruct, err = gi.StructNew("Gtk", "LinkButtonPrivate")
	})
	return err
}

type LinkButtonPrivate struct {
	native uintptr
}

var listBoxAccessibleClassStruct *gi.Struct
var listBoxAccessibleClassStruct_Once sync.Once

func listBoxAccessibleClassStruct_Set() error {
	var err error
	listBoxAccessibleClassStruct_Once.Do(func() {
		listBoxAccessibleClassStruct, err = gi.StructNew("Gtk", "ListBoxAccessibleClass")
	})
	return err
}

type ListBoxAccessibleClass struct {
	native uintptr
}

var listBoxAccessiblePrivateStruct *gi.Struct
var listBoxAccessiblePrivateStruct_Once sync.Once

func listBoxAccessiblePrivateStruct_Set() error {
	var err error
	listBoxAccessiblePrivateStruct_Once.Do(func() {
		listBoxAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ListBoxAccessiblePrivate")
	})
	return err
}

type ListBoxAccessiblePrivate struct {
	native uintptr
}

var listBoxClassStruct *gi.Struct
var listBoxClassStruct_Once sync.Once

func listBoxClassStruct_Set() error {
	var err error
	listBoxClassStruct_Once.Do(func() {
		listBoxClassStruct, err = gi.StructNew("Gtk", "ListBoxClass")
	})
	return err
}

type ListBoxClass struct {
	native uintptr
}

var listBoxRowAccessibleClassStruct *gi.Struct
var listBoxRowAccessibleClassStruct_Once sync.Once

func listBoxRowAccessibleClassStruct_Set() error {
	var err error
	listBoxRowAccessibleClassStruct_Once.Do(func() {
		listBoxRowAccessibleClassStruct, err = gi.StructNew("Gtk", "ListBoxRowAccessibleClass")
	})
	return err
}

type ListBoxRowAccessibleClass struct {
	native uintptr
}

var listBoxRowClassStruct *gi.Struct
var listBoxRowClassStruct_Once sync.Once

func listBoxRowClassStruct_Set() error {
	var err error
	listBoxRowClassStruct_Once.Do(func() {
		listBoxRowClassStruct, err = gi.StructNew("Gtk", "ListBoxRowClass")
	})
	return err
}

type ListBoxRowClass struct {
	native uintptr
}

var listStoreClassStruct *gi.Struct
var listStoreClassStruct_Once sync.Once

func listStoreClassStruct_Set() error {
	var err error
	listStoreClassStruct_Once.Do(func() {
		listStoreClassStruct, err = gi.StructNew("Gtk", "ListStoreClass")
	})
	return err
}

type ListStoreClass struct {
	native uintptr
}

var listStorePrivateStruct *gi.Struct
var listStorePrivateStruct_Once sync.Once

func listStorePrivateStruct_Set() error {
	var err error
	listStorePrivateStruct_Once.Do(func() {
		listStorePrivateStruct, err = gi.StructNew("Gtk", "ListStorePrivate")
	})
	return err
}

type ListStorePrivate struct {
	native uintptr
}

var lockButtonAccessibleClassStruct *gi.Struct
var lockButtonAccessibleClassStruct_Once sync.Once

func lockButtonAccessibleClassStruct_Set() error {
	var err error
	lockButtonAccessibleClassStruct_Once.Do(func() {
		lockButtonAccessibleClassStruct, err = gi.StructNew("Gtk", "LockButtonAccessibleClass")
	})
	return err
}

type LockButtonAccessibleClass struct {
	native uintptr
}

var lockButtonAccessiblePrivateStruct *gi.Struct
var lockButtonAccessiblePrivateStruct_Once sync.Once

func lockButtonAccessiblePrivateStruct_Set() error {
	var err error
	lockButtonAccessiblePrivateStruct_Once.Do(func() {
		lockButtonAccessiblePrivateStruct, err = gi.StructNew("Gtk", "LockButtonAccessiblePrivate")
	})
	return err
}

type LockButtonAccessiblePrivate struct {
	native uintptr
}

var lockButtonClassStruct *gi.Struct
var lockButtonClassStruct_Once sync.Once

func lockButtonClassStruct_Set() error {
	var err error
	lockButtonClassStruct_Once.Do(func() {
		lockButtonClassStruct, err = gi.StructNew("Gtk", "LockButtonClass")
	})
	return err
}

type LockButtonClass struct {
	native uintptr
}

var lockButtonPrivateStruct *gi.Struct
var lockButtonPrivateStruct_Once sync.Once

func lockButtonPrivateStruct_Set() error {
	var err error
	lockButtonPrivateStruct_Once.Do(func() {
		lockButtonPrivateStruct, err = gi.StructNew("Gtk", "LockButtonPrivate")
	})
	return err
}

type LockButtonPrivate struct {
	native uintptr
}

var menuAccessibleClassStruct *gi.Struct
var menuAccessibleClassStruct_Once sync.Once

func menuAccessibleClassStruct_Set() error {
	var err error
	menuAccessibleClassStruct_Once.Do(func() {
		menuAccessibleClassStruct, err = gi.StructNew("Gtk", "MenuAccessibleClass")
	})
	return err
}

type MenuAccessibleClass struct {
	native uintptr
}

var menuAccessiblePrivateStruct *gi.Struct
var menuAccessiblePrivateStruct_Once sync.Once

func menuAccessiblePrivateStruct_Set() error {
	var err error
	menuAccessiblePrivateStruct_Once.Do(func() {
		menuAccessiblePrivateStruct, err = gi.StructNew("Gtk", "MenuAccessiblePrivate")
	})
	return err
}

type MenuAccessiblePrivate struct {
	native uintptr
}

var menuBarClassStruct *gi.Struct
var menuBarClassStruct_Once sync.Once

func menuBarClassStruct_Set() error {
	var err error
	menuBarClassStruct_Once.Do(func() {
		menuBarClassStruct, err = gi.StructNew("Gtk", "MenuBarClass")
	})
	return err
}

type MenuBarClass struct {
	native uintptr
}

var menuBarPrivateStruct *gi.Struct
var menuBarPrivateStruct_Once sync.Once

func menuBarPrivateStruct_Set() error {
	var err error
	menuBarPrivateStruct_Once.Do(func() {
		menuBarPrivateStruct, err = gi.StructNew("Gtk", "MenuBarPrivate")
	})
	return err
}

type MenuBarPrivate struct {
	native uintptr
}

var menuButtonAccessibleClassStruct *gi.Struct
var menuButtonAccessibleClassStruct_Once sync.Once

func menuButtonAccessibleClassStruct_Set() error {
	var err error
	menuButtonAccessibleClassStruct_Once.Do(func() {
		menuButtonAccessibleClassStruct, err = gi.StructNew("Gtk", "MenuButtonAccessibleClass")
	})
	return err
}

type MenuButtonAccessibleClass struct {
	native uintptr
}

var menuButtonAccessiblePrivateStruct *gi.Struct
var menuButtonAccessiblePrivateStruct_Once sync.Once

func menuButtonAccessiblePrivateStruct_Set() error {
	var err error
	menuButtonAccessiblePrivateStruct_Once.Do(func() {
		menuButtonAccessiblePrivateStruct, err = gi.StructNew("Gtk", "MenuButtonAccessiblePrivate")
	})
	return err
}

type MenuButtonAccessiblePrivate struct {
	native uintptr
}

var menuButtonClassStruct *gi.Struct
var menuButtonClassStruct_Once sync.Once

func menuButtonClassStruct_Set() error {
	var err error
	menuButtonClassStruct_Once.Do(func() {
		menuButtonClassStruct, err = gi.StructNew("Gtk", "MenuButtonClass")
	})
	return err
}

type MenuButtonClass struct {
	native uintptr
}

var menuButtonPrivateStruct *gi.Struct
var menuButtonPrivateStruct_Once sync.Once

func menuButtonPrivateStruct_Set() error {
	var err error
	menuButtonPrivateStruct_Once.Do(func() {
		menuButtonPrivateStruct, err = gi.StructNew("Gtk", "MenuButtonPrivate")
	})
	return err
}

type MenuButtonPrivate struct {
	native uintptr
}

var menuClassStruct *gi.Struct
var menuClassStruct_Once sync.Once

func menuClassStruct_Set() error {
	var err error
	menuClassStruct_Once.Do(func() {
		menuClassStruct, err = gi.StructNew("Gtk", "MenuClass")
	})
	return err
}

type MenuClass struct {
	native uintptr
}

var menuItemAccessibleClassStruct *gi.Struct
var menuItemAccessibleClassStruct_Once sync.Once

func menuItemAccessibleClassStruct_Set() error {
	var err error
	menuItemAccessibleClassStruct_Once.Do(func() {
		menuItemAccessibleClassStruct, err = gi.StructNew("Gtk", "MenuItemAccessibleClass")
	})
	return err
}

type MenuItemAccessibleClass struct {
	native uintptr
}

var menuItemAccessiblePrivateStruct *gi.Struct
var menuItemAccessiblePrivateStruct_Once sync.Once

func menuItemAccessiblePrivateStruct_Set() error {
	var err error
	menuItemAccessiblePrivateStruct_Once.Do(func() {
		menuItemAccessiblePrivateStruct, err = gi.StructNew("Gtk", "MenuItemAccessiblePrivate")
	})
	return err
}

type MenuItemAccessiblePrivate struct {
	native uintptr
}

var menuItemClassStruct *gi.Struct
var menuItemClassStruct_Once sync.Once

func menuItemClassStruct_Set() error {
	var err error
	menuItemClassStruct_Once.Do(func() {
		menuItemClassStruct, err = gi.StructNew("Gtk", "MenuItemClass")
	})
	return err
}

type MenuItemClass struct {
	native uintptr
}

var menuItemPrivateStruct *gi.Struct
var menuItemPrivateStruct_Once sync.Once

func menuItemPrivateStruct_Set() error {
	var err error
	menuItemPrivateStruct_Once.Do(func() {
		menuItemPrivateStruct, err = gi.StructNew("Gtk", "MenuItemPrivate")
	})
	return err
}

type MenuItemPrivate struct {
	native uintptr
}

var menuPrivateStruct *gi.Struct
var menuPrivateStruct_Once sync.Once

func menuPrivateStruct_Set() error {
	var err error
	menuPrivateStruct_Once.Do(func() {
		menuPrivateStruct, err = gi.StructNew("Gtk", "MenuPrivate")
	})
	return err
}

type MenuPrivate struct {
	native uintptr
}

var menuShellAccessibleClassStruct *gi.Struct
var menuShellAccessibleClassStruct_Once sync.Once

func menuShellAccessibleClassStruct_Set() error {
	var err error
	menuShellAccessibleClassStruct_Once.Do(func() {
		menuShellAccessibleClassStruct, err = gi.StructNew("Gtk", "MenuShellAccessibleClass")
	})
	return err
}

type MenuShellAccessibleClass struct {
	native uintptr
}

var menuShellAccessiblePrivateStruct *gi.Struct
var menuShellAccessiblePrivateStruct_Once sync.Once

func menuShellAccessiblePrivateStruct_Set() error {
	var err error
	menuShellAccessiblePrivateStruct_Once.Do(func() {
		menuShellAccessiblePrivateStruct, err = gi.StructNew("Gtk", "MenuShellAccessiblePrivate")
	})
	return err
}

type MenuShellAccessiblePrivate struct {
	native uintptr
}

var menuShellClassStruct *gi.Struct
var menuShellClassStruct_Once sync.Once

func menuShellClassStruct_Set() error {
	var err error
	menuShellClassStruct_Once.Do(func() {
		menuShellClassStruct, err = gi.StructNew("Gtk", "MenuShellClass")
	})
	return err
}

type MenuShellClass struct {
	native uintptr
}

var menuShellPrivateStruct *gi.Struct
var menuShellPrivateStruct_Once sync.Once

func menuShellPrivateStruct_Set() error {
	var err error
	menuShellPrivateStruct_Once.Do(func() {
		menuShellPrivateStruct, err = gi.StructNew("Gtk", "MenuShellPrivate")
	})
	return err
}

type MenuShellPrivate struct {
	native uintptr
}

var menuToolButtonClassStruct *gi.Struct
var menuToolButtonClassStruct_Once sync.Once

func menuToolButtonClassStruct_Set() error {
	var err error
	menuToolButtonClassStruct_Once.Do(func() {
		menuToolButtonClassStruct, err = gi.StructNew("Gtk", "MenuToolButtonClass")
	})
	return err
}

type MenuToolButtonClass struct {
	native uintptr
}

var menuToolButtonPrivateStruct *gi.Struct
var menuToolButtonPrivateStruct_Once sync.Once

func menuToolButtonPrivateStruct_Set() error {
	var err error
	menuToolButtonPrivateStruct_Once.Do(func() {
		menuToolButtonPrivateStruct, err = gi.StructNew("Gtk", "MenuToolButtonPrivate")
	})
	return err
}

type MenuToolButtonPrivate struct {
	native uintptr
}

var messageDialogClassStruct *gi.Struct
var messageDialogClassStruct_Once sync.Once

func messageDialogClassStruct_Set() error {
	var err error
	messageDialogClassStruct_Once.Do(func() {
		messageDialogClassStruct, err = gi.StructNew("Gtk", "MessageDialogClass")
	})
	return err
}

type MessageDialogClass struct {
	native uintptr
}

var messageDialogPrivateStruct *gi.Struct
var messageDialogPrivateStruct_Once sync.Once

func messageDialogPrivateStruct_Set() error {
	var err error
	messageDialogPrivateStruct_Once.Do(func() {
		messageDialogPrivateStruct, err = gi.StructNew("Gtk", "MessageDialogPrivate")
	})
	return err
}

type MessageDialogPrivate struct {
	native uintptr
}

var miscClassStruct *gi.Struct
var miscClassStruct_Once sync.Once

func miscClassStruct_Set() error {
	var err error
	miscClassStruct_Once.Do(func() {
		miscClassStruct, err = gi.StructNew("Gtk", "MiscClass")
	})
	return err
}

type MiscClass struct {
	native uintptr
}

var miscPrivateStruct *gi.Struct
var miscPrivateStruct_Once sync.Once

func miscPrivateStruct_Set() error {
	var err error
	miscPrivateStruct_Once.Do(func() {
		miscPrivateStruct, err = gi.StructNew("Gtk", "MiscPrivate")
	})
	return err
}

type MiscPrivate struct {
	native uintptr
}

var mountOperationClassStruct *gi.Struct
var mountOperationClassStruct_Once sync.Once

func mountOperationClassStruct_Set() error {
	var err error
	mountOperationClassStruct_Once.Do(func() {
		mountOperationClassStruct, err = gi.StructNew("Gtk", "MountOperationClass")
	})
	return err
}

type MountOperationClass struct {
	native uintptr
}

var mountOperationPrivateStruct *gi.Struct
var mountOperationPrivateStruct_Once sync.Once

func mountOperationPrivateStruct_Set() error {
	var err error
	mountOperationPrivateStruct_Once.Do(func() {
		mountOperationPrivateStruct, err = gi.StructNew("Gtk", "MountOperationPrivate")
	})
	return err
}

type MountOperationPrivate struct {
	native uintptr
}

var nativeDialogClassStruct *gi.Struct
var nativeDialogClassStruct_Once sync.Once

func nativeDialogClassStruct_Set() error {
	var err error
	nativeDialogClassStruct_Once.Do(func() {
		nativeDialogClassStruct, err = gi.StructNew("Gtk", "NativeDialogClass")
	})
	return err
}

type NativeDialogClass struct {
	native uintptr
}

var notebookAccessibleClassStruct *gi.Struct
var notebookAccessibleClassStruct_Once sync.Once

func notebookAccessibleClassStruct_Set() error {
	var err error
	notebookAccessibleClassStruct_Once.Do(func() {
		notebookAccessibleClassStruct, err = gi.StructNew("Gtk", "NotebookAccessibleClass")
	})
	return err
}

type NotebookAccessibleClass struct {
	native uintptr
}

var notebookAccessiblePrivateStruct *gi.Struct
var notebookAccessiblePrivateStruct_Once sync.Once

func notebookAccessiblePrivateStruct_Set() error {
	var err error
	notebookAccessiblePrivateStruct_Once.Do(func() {
		notebookAccessiblePrivateStruct, err = gi.StructNew("Gtk", "NotebookAccessiblePrivate")
	})
	return err
}

type NotebookAccessiblePrivate struct {
	native uintptr
}

var notebookClassStruct *gi.Struct
var notebookClassStruct_Once sync.Once

func notebookClassStruct_Set() error {
	var err error
	notebookClassStruct_Once.Do(func() {
		notebookClassStruct, err = gi.StructNew("Gtk", "NotebookClass")
	})
	return err
}

type NotebookClass struct {
	native uintptr
}

var notebookPageAccessibleClassStruct *gi.Struct
var notebookPageAccessibleClassStruct_Once sync.Once

func notebookPageAccessibleClassStruct_Set() error {
	var err error
	notebookPageAccessibleClassStruct_Once.Do(func() {
		notebookPageAccessibleClassStruct, err = gi.StructNew("Gtk", "NotebookPageAccessibleClass")
	})
	return err
}

type NotebookPageAccessibleClass struct {
	native uintptr
}

var notebookPageAccessiblePrivateStruct *gi.Struct
var notebookPageAccessiblePrivateStruct_Once sync.Once

func notebookPageAccessiblePrivateStruct_Set() error {
	var err error
	notebookPageAccessiblePrivateStruct_Once.Do(func() {
		notebookPageAccessiblePrivateStruct, err = gi.StructNew("Gtk", "NotebookPageAccessiblePrivate")
	})
	return err
}

type NotebookPageAccessiblePrivate struct {
	native uintptr
}

var notebookPrivateStruct *gi.Struct
var notebookPrivateStruct_Once sync.Once

func notebookPrivateStruct_Set() error {
	var err error
	notebookPrivateStruct_Once.Do(func() {
		notebookPrivateStruct, err = gi.StructNew("Gtk", "NotebookPrivate")
	})
	return err
}

type NotebookPrivate struct {
	native uintptr
}

var numerableIconClassStruct *gi.Struct
var numerableIconClassStruct_Once sync.Once

func numerableIconClassStruct_Set() error {
	var err error
	numerableIconClassStruct_Once.Do(func() {
		numerableIconClassStruct, err = gi.StructNew("Gtk", "NumerableIconClass")
	})
	return err
}

type NumerableIconClass struct {
	native uintptr
}

var numerableIconPrivateStruct *gi.Struct
var numerableIconPrivateStruct_Once sync.Once

func numerableIconPrivateStruct_Set() error {
	var err error
	numerableIconPrivateStruct_Once.Do(func() {
		numerableIconPrivateStruct, err = gi.StructNew("Gtk", "NumerableIconPrivate")
	})
	return err
}

type NumerableIconPrivate struct {
	native uintptr
}

var offscreenWindowClassStruct *gi.Struct
var offscreenWindowClassStruct_Once sync.Once

func offscreenWindowClassStruct_Set() error {
	var err error
	offscreenWindowClassStruct_Once.Do(func() {
		offscreenWindowClassStruct, err = gi.StructNew("Gtk", "OffscreenWindowClass")
	})
	return err
}

type OffscreenWindowClass struct {
	native uintptr
}

var orientableIfaceStruct *gi.Struct
var orientableIfaceStruct_Once sync.Once

func orientableIfaceStruct_Set() error {
	var err error
	orientableIfaceStruct_Once.Do(func() {
		orientableIfaceStruct, err = gi.StructNew("Gtk", "OrientableIface")
	})
	return err
}

type OrientableIface struct {
	native uintptr
}

var overlayClassStruct *gi.Struct
var overlayClassStruct_Once sync.Once

func overlayClassStruct_Set() error {
	var err error
	overlayClassStruct_Once.Do(func() {
		overlayClassStruct, err = gi.StructNew("Gtk", "OverlayClass")
	})
	return err
}

type OverlayClass struct {
	native uintptr
}

var overlayPrivateStruct *gi.Struct
var overlayPrivateStruct_Once sync.Once

func overlayPrivateStruct_Set() error {
	var err error
	overlayPrivateStruct_Once.Do(func() {
		overlayPrivateStruct, err = gi.StructNew("Gtk", "OverlayPrivate")
	})
	return err
}

type OverlayPrivate struct {
	native uintptr
}

var padActionEntryStruct *gi.Struct
var padActionEntryStruct_Once sync.Once

func padActionEntryStruct_Set() error {
	var err error
	padActionEntryStruct_Once.Do(func() {
		padActionEntryStruct, err = gi.StructNew("Gtk", "PadActionEntry")
	})
	return err
}

type PadActionEntry struct {
	native uintptr
}

var padControllerClassStruct *gi.Struct
var padControllerClassStruct_Once sync.Once

func padControllerClassStruct_Set() error {
	var err error
	padControllerClassStruct_Once.Do(func() {
		padControllerClassStruct, err = gi.StructNew("Gtk", "PadControllerClass")
	})
	return err
}

type PadControllerClass struct {
	native uintptr
}

var pageRangeStruct *gi.Struct
var pageRangeStruct_Once sync.Once

func pageRangeStruct_Set() error {
	var err error
	pageRangeStruct_Once.Do(func() {
		pageRangeStruct, err = gi.StructNew("Gtk", "PageRange")
	})
	return err
}

type PageRange struct {
	native uintptr
}

var panedAccessibleClassStruct *gi.Struct
var panedAccessibleClassStruct_Once sync.Once

func panedAccessibleClassStruct_Set() error {
	var err error
	panedAccessibleClassStruct_Once.Do(func() {
		panedAccessibleClassStruct, err = gi.StructNew("Gtk", "PanedAccessibleClass")
	})
	return err
}

type PanedAccessibleClass struct {
	native uintptr
}

var panedAccessiblePrivateStruct *gi.Struct
var panedAccessiblePrivateStruct_Once sync.Once

func panedAccessiblePrivateStruct_Set() error {
	var err error
	panedAccessiblePrivateStruct_Once.Do(func() {
		panedAccessiblePrivateStruct, err = gi.StructNew("Gtk", "PanedAccessiblePrivate")
	})
	return err
}

type PanedAccessiblePrivate struct {
	native uintptr
}

var panedClassStruct *gi.Struct
var panedClassStruct_Once sync.Once

func panedClassStruct_Set() error {
	var err error
	panedClassStruct_Once.Do(func() {
		panedClassStruct, err = gi.StructNew("Gtk", "PanedClass")
	})
	return err
}

type PanedClass struct {
	native uintptr
}

var panedPrivateStruct *gi.Struct
var panedPrivateStruct_Once sync.Once

func panedPrivateStruct_Set() error {
	var err error
	panedPrivateStruct_Once.Do(func() {
		panedPrivateStruct, err = gi.StructNew("Gtk", "PanedPrivate")
	})
	return err
}

type PanedPrivate struct {
	native uintptr
}

var paperSizeStruct *gi.Struct
var paperSizeStruct_Once sync.Once

func paperSizeStruct_Set() error {
	var err error
	paperSizeStruct_Once.Do(func() {
		paperSizeStruct, err = gi.StructNew("Gtk", "PaperSize")
	})
	return err
}

type PaperSize struct {
	native uintptr
}

var paperSizeNewFunction *gi.Function
var paperSizeNewFunction_Once sync.Once

func paperSizeNewFunction_Set() error {
	var err error
	paperSizeNewFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeNewFunction, err = paperSizeStruct.InvokerNew("new")
	})
	return err
}

// PaperSizeNew is a representation of the C type gtk_paper_size_new.
func PaperSizeNew(name string) *PaperSize {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := paperSizeNewFunction_Set()
	if err == nil {
		ret = paperSizeNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &PaperSize{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_new_custom' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_new_from_gvariant' : parameter 'variant' of type 'GLib.Variant' not supported

var paperSizeNewFromIppFunction *gi.Function
var paperSizeNewFromIppFunction_Once sync.Once

func paperSizeNewFromIppFunction_Set() error {
	var err error
	paperSizeNewFromIppFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeNewFromIppFunction, err = paperSizeStruct.InvokerNew("new_from_ipp")
	})
	return err
}

// PaperSizeNewFromIpp is a representation of the C type gtk_paper_size_new_from_ipp.
func PaperSizeNewFromIpp(ippName string, width float64, height float64) *PaperSize {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(ippName)
	inArgs[1].SetFloat64(width)
	inArgs[2].SetFloat64(height)

	var ret gi.Argument

	err := paperSizeNewFromIppFunction_Set()
	if err == nil {
		ret = paperSizeNewFromIppFunction.Invoke(inArgs[:], nil)
	}

	retGo := &PaperSize{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_new_from_key_file' : parameter 'key_file' of type 'GLib.KeyFile' not supported

var paperSizeNewFromPpdFunction *gi.Function
var paperSizeNewFromPpdFunction_Once sync.Once

func paperSizeNewFromPpdFunction_Set() error {
	var err error
	paperSizeNewFromPpdFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeNewFromPpdFunction, err = paperSizeStruct.InvokerNew("new_from_ppd")
	})
	return err
}

// PaperSizeNewFromPpd is a representation of the C type gtk_paper_size_new_from_ppd.
func PaperSizeNewFromPpd(ppdName string, ppdDisplayName string, width float64, height float64) *PaperSize {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(ppdName)
	inArgs[1].SetString(ppdDisplayName)
	inArgs[2].SetFloat64(width)
	inArgs[3].SetFloat64(height)

	var ret gi.Argument

	err := paperSizeNewFromPpdFunction_Set()
	if err == nil {
		ret = paperSizeNewFromPpdFunction.Invoke(inArgs[:], nil)
	}

	retGo := &PaperSize{native: ret.Pointer()}

	return retGo
}

var paperSizeCopyFunction *gi.Function
var paperSizeCopyFunction_Once sync.Once

func paperSizeCopyFunction_Set() error {
	var err error
	paperSizeCopyFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeCopyFunction, err = paperSizeStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_paper_size_copy.
func (recv *PaperSize) Copy() *PaperSize {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paperSizeCopyFunction_Set()
	if err == nil {
		ret = paperSizeCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &PaperSize{native: ret.Pointer()}

	return retGo
}

var paperSizeFreeFunction *gi.Function
var paperSizeFreeFunction_Once sync.Once

func paperSizeFreeFunction_Set() error {
	var err error
	paperSizeFreeFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeFreeFunction, err = paperSizeStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_paper_size_free.
func (recv *PaperSize) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := paperSizeFreeFunction_Set()
	if err == nil {
		paperSizeFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_paper_size_get_default_bottom_margin' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_default_left_margin' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_default_right_margin' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_get_default_top_margin' : parameter 'unit' of type 'Unit' not supported

var paperSizeGetDisplayNameFunction *gi.Function
var paperSizeGetDisplayNameFunction_Once sync.Once

func paperSizeGetDisplayNameFunction_Set() error {
	var err error
	paperSizeGetDisplayNameFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeGetDisplayNameFunction, err = paperSizeStruct.InvokerNew("get_display_name")
	})
	return err
}

// GetDisplayName is a representation of the C type gtk_paper_size_get_display_name.
func (recv *PaperSize) GetDisplayName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paperSizeGetDisplayNameFunction_Set()
	if err == nil {
		ret = paperSizeGetDisplayNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_get_height' : parameter 'unit' of type 'Unit' not supported

var paperSizeGetNameFunction *gi.Function
var paperSizeGetNameFunction_Once sync.Once

func paperSizeGetNameFunction_Set() error {
	var err error
	paperSizeGetNameFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeGetNameFunction, err = paperSizeStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gtk_paper_size_get_name.
func (recv *PaperSize) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paperSizeGetNameFunction_Set()
	if err == nil {
		ret = paperSizeGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var paperSizeGetPpdNameFunction *gi.Function
var paperSizeGetPpdNameFunction_Once sync.Once

func paperSizeGetPpdNameFunction_Set() error {
	var err error
	paperSizeGetPpdNameFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeGetPpdNameFunction, err = paperSizeStruct.InvokerNew("get_ppd_name")
	})
	return err
}

// GetPpdName is a representation of the C type gtk_paper_size_get_ppd_name.
func (recv *PaperSize) GetPpdName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paperSizeGetPpdNameFunction_Set()
	if err == nil {
		ret = paperSizeGetPpdNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_get_width' : parameter 'unit' of type 'Unit' not supported

var paperSizeIsCustomFunction *gi.Function
var paperSizeIsCustomFunction_Once sync.Once

func paperSizeIsCustomFunction_Set() error {
	var err error
	paperSizeIsCustomFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeIsCustomFunction, err = paperSizeStruct.InvokerNew("is_custom")
	})
	return err
}

// IsCustom is a representation of the C type gtk_paper_size_is_custom.
func (recv *PaperSize) IsCustom() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paperSizeIsCustomFunction_Set()
	if err == nil {
		ret = paperSizeIsCustomFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var paperSizeIsEqualFunction *gi.Function
var paperSizeIsEqualFunction_Once sync.Once

func paperSizeIsEqualFunction_Set() error {
	var err error
	paperSizeIsEqualFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeIsEqualFunction, err = paperSizeStruct.InvokerNew("is_equal")
	})
	return err
}

// IsEqual is a representation of the C type gtk_paper_size_is_equal.
func (recv *PaperSize) IsEqual(size2 *PaperSize) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(size2.native)

	var ret gi.Argument

	err := paperSizeIsEqualFunction_Set()
	if err == nil {
		ret = paperSizeIsEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var paperSizeIsIppFunction *gi.Function
var paperSizeIsIppFunction_Once sync.Once

func paperSizeIsIppFunction_Set() error {
	var err error
	paperSizeIsIppFunction_Once.Do(func() {
		err = paperSizeStruct_Set()
		if err != nil {
			return
		}
		paperSizeIsIppFunction, err = paperSizeStruct.InvokerNew("is_ipp")
	})
	return err
}

// IsIpp is a representation of the C type gtk_paper_size_is_ipp.
func (recv *PaperSize) IsIpp() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := paperSizeIsIppFunction_Set()
	if err == nil {
		ret = paperSizeIsIppFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_paper_size_set_size' : parameter 'unit' of type 'Unit' not supported

// UNSUPPORTED : C value 'gtk_paper_size_to_gvariant' : return type 'GLib.Variant' not supported

// UNSUPPORTED : C value 'gtk_paper_size_to_key_file' : parameter 'key_file' of type 'GLib.KeyFile' not supported

var placesSidebarClassStruct *gi.Struct
var placesSidebarClassStruct_Once sync.Once

func placesSidebarClassStruct_Set() error {
	var err error
	placesSidebarClassStruct_Once.Do(func() {
		placesSidebarClassStruct, err = gi.StructNew("Gtk", "PlacesSidebarClass")
	})
	return err
}

type PlacesSidebarClass struct {
	native uintptr
}

var plugClassStruct *gi.Struct
var plugClassStruct_Once sync.Once

func plugClassStruct_Set() error {
	var err error
	plugClassStruct_Once.Do(func() {
		plugClassStruct, err = gi.StructNew("Gtk", "PlugClass")
	})
	return err
}

type PlugClass struct {
	native uintptr
}

var plugPrivateStruct *gi.Struct
var plugPrivateStruct_Once sync.Once

func plugPrivateStruct_Set() error {
	var err error
	plugPrivateStruct_Once.Do(func() {
		plugPrivateStruct, err = gi.StructNew("Gtk", "PlugPrivate")
	})
	return err
}

type PlugPrivate struct {
	native uintptr
}

var popoverAccessibleClassStruct *gi.Struct
var popoverAccessibleClassStruct_Once sync.Once

func popoverAccessibleClassStruct_Set() error {
	var err error
	popoverAccessibleClassStruct_Once.Do(func() {
		popoverAccessibleClassStruct, err = gi.StructNew("Gtk", "PopoverAccessibleClass")
	})
	return err
}

type PopoverAccessibleClass struct {
	native uintptr
}

var popoverClassStruct *gi.Struct
var popoverClassStruct_Once sync.Once

func popoverClassStruct_Set() error {
	var err error
	popoverClassStruct_Once.Do(func() {
		popoverClassStruct, err = gi.StructNew("Gtk", "PopoverClass")
	})
	return err
}

type PopoverClass struct {
	native uintptr
}

var popoverMenuClassStruct *gi.Struct
var popoverMenuClassStruct_Once sync.Once

func popoverMenuClassStruct_Set() error {
	var err error
	popoverMenuClassStruct_Once.Do(func() {
		popoverMenuClassStruct, err = gi.StructNew("Gtk", "PopoverMenuClass")
	})
	return err
}

type PopoverMenuClass struct {
	native uintptr
}

var popoverPrivateStruct *gi.Struct
var popoverPrivateStruct_Once sync.Once

func popoverPrivateStruct_Set() error {
	var err error
	popoverPrivateStruct_Once.Do(func() {
		popoverPrivateStruct, err = gi.StructNew("Gtk", "PopoverPrivate")
	})
	return err
}

type PopoverPrivate struct {
	native uintptr
}

var printOperationClassStruct *gi.Struct
var printOperationClassStruct_Once sync.Once

func printOperationClassStruct_Set() error {
	var err error
	printOperationClassStruct_Once.Do(func() {
		printOperationClassStruct, err = gi.StructNew("Gtk", "PrintOperationClass")
	})
	return err
}

type PrintOperationClass struct {
	native uintptr
}

var printOperationPreviewIfaceStruct *gi.Struct
var printOperationPreviewIfaceStruct_Once sync.Once

func printOperationPreviewIfaceStruct_Set() error {
	var err error
	printOperationPreviewIfaceStruct_Once.Do(func() {
		printOperationPreviewIfaceStruct, err = gi.StructNew("Gtk", "PrintOperationPreviewIface")
	})
	return err
}

type PrintOperationPreviewIface struct {
	native uintptr
}

var printOperationPrivateStruct *gi.Struct
var printOperationPrivateStruct_Once sync.Once

func printOperationPrivateStruct_Set() error {
	var err error
	printOperationPrivateStruct_Once.Do(func() {
		printOperationPrivateStruct, err = gi.StructNew("Gtk", "PrintOperationPrivate")
	})
	return err
}

type PrintOperationPrivate struct {
	native uintptr
}

var progressBarAccessibleClassStruct *gi.Struct
var progressBarAccessibleClassStruct_Once sync.Once

func progressBarAccessibleClassStruct_Set() error {
	var err error
	progressBarAccessibleClassStruct_Once.Do(func() {
		progressBarAccessibleClassStruct, err = gi.StructNew("Gtk", "ProgressBarAccessibleClass")
	})
	return err
}

type ProgressBarAccessibleClass struct {
	native uintptr
}

var progressBarAccessiblePrivateStruct *gi.Struct
var progressBarAccessiblePrivateStruct_Once sync.Once

func progressBarAccessiblePrivateStruct_Set() error {
	var err error
	progressBarAccessiblePrivateStruct_Once.Do(func() {
		progressBarAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ProgressBarAccessiblePrivate")
	})
	return err
}

type ProgressBarAccessiblePrivate struct {
	native uintptr
}

var progressBarClassStruct *gi.Struct
var progressBarClassStruct_Once sync.Once

func progressBarClassStruct_Set() error {
	var err error
	progressBarClassStruct_Once.Do(func() {
		progressBarClassStruct, err = gi.StructNew("Gtk", "ProgressBarClass")
	})
	return err
}

type ProgressBarClass struct {
	native uintptr
}

var progressBarPrivateStruct *gi.Struct
var progressBarPrivateStruct_Once sync.Once

func progressBarPrivateStruct_Set() error {
	var err error
	progressBarPrivateStruct_Once.Do(func() {
		progressBarPrivateStruct, err = gi.StructNew("Gtk", "ProgressBarPrivate")
	})
	return err
}

type ProgressBarPrivate struct {
	native uintptr
}

var radioActionClassStruct *gi.Struct
var radioActionClassStruct_Once sync.Once

func radioActionClassStruct_Set() error {
	var err error
	radioActionClassStruct_Once.Do(func() {
		radioActionClassStruct, err = gi.StructNew("Gtk", "RadioActionClass")
	})
	return err
}

type RadioActionClass struct {
	native uintptr
}

var radioActionEntryStruct *gi.Struct
var radioActionEntryStruct_Once sync.Once

func radioActionEntryStruct_Set() error {
	var err error
	radioActionEntryStruct_Once.Do(func() {
		radioActionEntryStruct, err = gi.StructNew("Gtk", "RadioActionEntry")
	})
	return err
}

type RadioActionEntry struct {
	native uintptr
}

var radioActionPrivateStruct *gi.Struct
var radioActionPrivateStruct_Once sync.Once

func radioActionPrivateStruct_Set() error {
	var err error
	radioActionPrivateStruct_Once.Do(func() {
		radioActionPrivateStruct, err = gi.StructNew("Gtk", "RadioActionPrivate")
	})
	return err
}

type RadioActionPrivate struct {
	native uintptr
}

var radioButtonAccessibleClassStruct *gi.Struct
var radioButtonAccessibleClassStruct_Once sync.Once

func radioButtonAccessibleClassStruct_Set() error {
	var err error
	radioButtonAccessibleClassStruct_Once.Do(func() {
		radioButtonAccessibleClassStruct, err = gi.StructNew("Gtk", "RadioButtonAccessibleClass")
	})
	return err
}

type RadioButtonAccessibleClass struct {
	native uintptr
}

var radioButtonAccessiblePrivateStruct *gi.Struct
var radioButtonAccessiblePrivateStruct_Once sync.Once

func radioButtonAccessiblePrivateStruct_Set() error {
	var err error
	radioButtonAccessiblePrivateStruct_Once.Do(func() {
		radioButtonAccessiblePrivateStruct, err = gi.StructNew("Gtk", "RadioButtonAccessiblePrivate")
	})
	return err
}

type RadioButtonAccessiblePrivate struct {
	native uintptr
}

var radioButtonClassStruct *gi.Struct
var radioButtonClassStruct_Once sync.Once

func radioButtonClassStruct_Set() error {
	var err error
	radioButtonClassStruct_Once.Do(func() {
		radioButtonClassStruct, err = gi.StructNew("Gtk", "RadioButtonClass")
	})
	return err
}

type RadioButtonClass struct {
	native uintptr
}

var radioButtonPrivateStruct *gi.Struct
var radioButtonPrivateStruct_Once sync.Once

func radioButtonPrivateStruct_Set() error {
	var err error
	radioButtonPrivateStruct_Once.Do(func() {
		radioButtonPrivateStruct, err = gi.StructNew("Gtk", "RadioButtonPrivate")
	})
	return err
}

type RadioButtonPrivate struct {
	native uintptr
}

var radioMenuItemAccessibleClassStruct *gi.Struct
var radioMenuItemAccessibleClassStruct_Once sync.Once

func radioMenuItemAccessibleClassStruct_Set() error {
	var err error
	radioMenuItemAccessibleClassStruct_Once.Do(func() {
		radioMenuItemAccessibleClassStruct, err = gi.StructNew("Gtk", "RadioMenuItemAccessibleClass")
	})
	return err
}

type RadioMenuItemAccessibleClass struct {
	native uintptr
}

var radioMenuItemAccessiblePrivateStruct *gi.Struct
var radioMenuItemAccessiblePrivateStruct_Once sync.Once

func radioMenuItemAccessiblePrivateStruct_Set() error {
	var err error
	radioMenuItemAccessiblePrivateStruct_Once.Do(func() {
		radioMenuItemAccessiblePrivateStruct, err = gi.StructNew("Gtk", "RadioMenuItemAccessiblePrivate")
	})
	return err
}

type RadioMenuItemAccessiblePrivate struct {
	native uintptr
}

var radioMenuItemClassStruct *gi.Struct
var radioMenuItemClassStruct_Once sync.Once

func radioMenuItemClassStruct_Set() error {
	var err error
	radioMenuItemClassStruct_Once.Do(func() {
		radioMenuItemClassStruct, err = gi.StructNew("Gtk", "RadioMenuItemClass")
	})
	return err
}

type RadioMenuItemClass struct {
	native uintptr
}

var radioMenuItemPrivateStruct *gi.Struct
var radioMenuItemPrivateStruct_Once sync.Once

func radioMenuItemPrivateStruct_Set() error {
	var err error
	radioMenuItemPrivateStruct_Once.Do(func() {
		radioMenuItemPrivateStruct, err = gi.StructNew("Gtk", "RadioMenuItemPrivate")
	})
	return err
}

type RadioMenuItemPrivate struct {
	native uintptr
}

var radioToolButtonClassStruct *gi.Struct
var radioToolButtonClassStruct_Once sync.Once

func radioToolButtonClassStruct_Set() error {
	var err error
	radioToolButtonClassStruct_Once.Do(func() {
		radioToolButtonClassStruct, err = gi.StructNew("Gtk", "RadioToolButtonClass")
	})
	return err
}

type RadioToolButtonClass struct {
	native uintptr
}

var rangeAccessibleClassStruct *gi.Struct
var rangeAccessibleClassStruct_Once sync.Once

func rangeAccessibleClassStruct_Set() error {
	var err error
	rangeAccessibleClassStruct_Once.Do(func() {
		rangeAccessibleClassStruct, err = gi.StructNew("Gtk", "RangeAccessibleClass")
	})
	return err
}

type RangeAccessibleClass struct {
	native uintptr
}

var rangeAccessiblePrivateStruct *gi.Struct
var rangeAccessiblePrivateStruct_Once sync.Once

func rangeAccessiblePrivateStruct_Set() error {
	var err error
	rangeAccessiblePrivateStruct_Once.Do(func() {
		rangeAccessiblePrivateStruct, err = gi.StructNew("Gtk", "RangeAccessiblePrivate")
	})
	return err
}

type RangeAccessiblePrivate struct {
	native uintptr
}

var rangeClassStruct *gi.Struct
var rangeClassStruct_Once sync.Once

func rangeClassStruct_Set() error {
	var err error
	rangeClassStruct_Once.Do(func() {
		rangeClassStruct, err = gi.StructNew("Gtk", "RangeClass")
	})
	return err
}

type RangeClass struct {
	native uintptr
}

var rangePrivateStruct *gi.Struct
var rangePrivateStruct_Once sync.Once

func rangePrivateStruct_Set() error {
	var err error
	rangePrivateStruct_Once.Do(func() {
		rangePrivateStruct, err = gi.StructNew("Gtk", "RangePrivate")
	})
	return err
}

type RangePrivate struct {
	native uintptr
}

var rcContextStruct *gi.Struct
var rcContextStruct_Once sync.Once

func rcContextStruct_Set() error {
	var err error
	rcContextStruct_Once.Do(func() {
		rcContextStruct, err = gi.StructNew("Gtk", "RcContext")
	})
	return err
}

type RcContext struct {
	native uintptr
}

var rcPropertyStruct *gi.Struct
var rcPropertyStruct_Once sync.Once

func rcPropertyStruct_Set() error {
	var err error
	rcPropertyStruct_Once.Do(func() {
		rcPropertyStruct, err = gi.StructNew("Gtk", "RcProperty")
	})
	return err
}

type RcProperty struct {
	native uintptr
}

var rcStyleClassStruct *gi.Struct
var rcStyleClassStruct_Once sync.Once

func rcStyleClassStruct_Set() error {
	var err error
	rcStyleClassStruct_Once.Do(func() {
		rcStyleClassStruct, err = gi.StructNew("Gtk", "RcStyleClass")
	})
	return err
}

type RcStyleClass struct {
	native uintptr
}

var recentActionClassStruct *gi.Struct
var recentActionClassStruct_Once sync.Once

func recentActionClassStruct_Set() error {
	var err error
	recentActionClassStruct_Once.Do(func() {
		recentActionClassStruct, err = gi.StructNew("Gtk", "RecentActionClass")
	})
	return err
}

type RecentActionClass struct {
	native uintptr
}

var recentActionPrivateStruct *gi.Struct
var recentActionPrivateStruct_Once sync.Once

func recentActionPrivateStruct_Set() error {
	var err error
	recentActionPrivateStruct_Once.Do(func() {
		recentActionPrivateStruct, err = gi.StructNew("Gtk", "RecentActionPrivate")
	})
	return err
}

type RecentActionPrivate struct {
	native uintptr
}

var recentChooserDialogClassStruct *gi.Struct
var recentChooserDialogClassStruct_Once sync.Once

func recentChooserDialogClassStruct_Set() error {
	var err error
	recentChooserDialogClassStruct_Once.Do(func() {
		recentChooserDialogClassStruct, err = gi.StructNew("Gtk", "RecentChooserDialogClass")
	})
	return err
}

type RecentChooserDialogClass struct {
	native uintptr
}

var recentChooserDialogPrivateStruct *gi.Struct
var recentChooserDialogPrivateStruct_Once sync.Once

func recentChooserDialogPrivateStruct_Set() error {
	var err error
	recentChooserDialogPrivateStruct_Once.Do(func() {
		recentChooserDialogPrivateStruct, err = gi.StructNew("Gtk", "RecentChooserDialogPrivate")
	})
	return err
}

type RecentChooserDialogPrivate struct {
	native uintptr
}

var recentChooserIfaceStruct *gi.Struct
var recentChooserIfaceStruct_Once sync.Once

func recentChooserIfaceStruct_Set() error {
	var err error
	recentChooserIfaceStruct_Once.Do(func() {
		recentChooserIfaceStruct, err = gi.StructNew("Gtk", "RecentChooserIface")
	})
	return err
}

type RecentChooserIface struct {
	native uintptr
}

var recentChooserMenuClassStruct *gi.Struct
var recentChooserMenuClassStruct_Once sync.Once

func recentChooserMenuClassStruct_Set() error {
	var err error
	recentChooserMenuClassStruct_Once.Do(func() {
		recentChooserMenuClassStruct, err = gi.StructNew("Gtk", "RecentChooserMenuClass")
	})
	return err
}

type RecentChooserMenuClass struct {
	native uintptr
}

var recentChooserMenuPrivateStruct *gi.Struct
var recentChooserMenuPrivateStruct_Once sync.Once

func recentChooserMenuPrivateStruct_Set() error {
	var err error
	recentChooserMenuPrivateStruct_Once.Do(func() {
		recentChooserMenuPrivateStruct, err = gi.StructNew("Gtk", "RecentChooserMenuPrivate")
	})
	return err
}

type RecentChooserMenuPrivate struct {
	native uintptr
}

var recentChooserWidgetClassStruct *gi.Struct
var recentChooserWidgetClassStruct_Once sync.Once

func recentChooserWidgetClassStruct_Set() error {
	var err error
	recentChooserWidgetClassStruct_Once.Do(func() {
		recentChooserWidgetClassStruct, err = gi.StructNew("Gtk", "RecentChooserWidgetClass")
	})
	return err
}

type RecentChooserWidgetClass struct {
	native uintptr
}

var recentChooserWidgetPrivateStruct *gi.Struct
var recentChooserWidgetPrivateStruct_Once sync.Once

func recentChooserWidgetPrivateStruct_Set() error {
	var err error
	recentChooserWidgetPrivateStruct_Once.Do(func() {
		recentChooserWidgetPrivateStruct, err = gi.StructNew("Gtk", "RecentChooserWidgetPrivate")
	})
	return err
}

type RecentChooserWidgetPrivate struct {
	native uintptr
}

var recentDataStruct *gi.Struct
var recentDataStruct_Once sync.Once

func recentDataStruct_Set() error {
	var err error
	recentDataStruct_Once.Do(func() {
		recentDataStruct, err = gi.StructNew("Gtk", "RecentData")
	})
	return err
}

type RecentData struct {
	native uintptr
}

var recentFilterInfoStruct *gi.Struct
var recentFilterInfoStruct_Once sync.Once

func recentFilterInfoStruct_Set() error {
	var err error
	recentFilterInfoStruct_Once.Do(func() {
		recentFilterInfoStruct, err = gi.StructNew("Gtk", "RecentFilterInfo")
	})
	return err
}

type RecentFilterInfo struct {
	native uintptr
}

var recentInfoStruct *gi.Struct
var recentInfoStruct_Once sync.Once

func recentInfoStruct_Set() error {
	var err error
	recentInfoStruct_Once.Do(func() {
		recentInfoStruct, err = gi.StructNew("Gtk", "RecentInfo")
	})
	return err
}

type RecentInfo struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_recent_info_create_app_info' : return type 'Gio.AppInfo' not supported

var recentInfoExistsFunction *gi.Function
var recentInfoExistsFunction_Once sync.Once

func recentInfoExistsFunction_Set() error {
	var err error
	recentInfoExistsFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoExistsFunction, err = recentInfoStruct.InvokerNew("exists")
	})
	return err
}

// Exists is a representation of the C type gtk_recent_info_exists.
func (recv *RecentInfo) Exists() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoExistsFunction_Set()
	if err == nil {
		ret = recentInfoExistsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentInfoGetAddedFunction *gi.Function
var recentInfoGetAddedFunction_Once sync.Once

func recentInfoGetAddedFunction_Set() error {
	var err error
	recentInfoGetAddedFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetAddedFunction, err = recentInfoStruct.InvokerNew("get_added")
	})
	return err
}

// GetAdded is a representation of the C type gtk_recent_info_get_added.
func (recv *RecentInfo) GetAdded() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetAddedFunction_Set()
	if err == nil {
		ret = recentInfoGetAddedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var recentInfoGetAgeFunction *gi.Function
var recentInfoGetAgeFunction_Once sync.Once

func recentInfoGetAgeFunction_Set() error {
	var err error
	recentInfoGetAgeFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetAgeFunction, err = recentInfoStruct.InvokerNew("get_age")
	})
	return err
}

// GetAge is a representation of the C type gtk_recent_info_get_age.
func (recv *RecentInfo) GetAge() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetAgeFunction_Set()
	if err == nil {
		ret = recentInfoGetAgeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var recentInfoGetApplicationInfoFunction *gi.Function
var recentInfoGetApplicationInfoFunction_Once sync.Once

func recentInfoGetApplicationInfoFunction_Set() error {
	var err error
	recentInfoGetApplicationInfoFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetApplicationInfoFunction, err = recentInfoStruct.InvokerNew("get_application_info")
	})
	return err
}

// GetApplicationInfo is a representation of the C type gtk_recent_info_get_application_info.
func (recv *RecentInfo) GetApplicationInfo(appName string) (bool, string, uint32, int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(appName)

	var outArgs [3]gi.Argument
	var ret gi.Argument

	err := recentInfoGetApplicationInfoFunction_Set()
	if err == nil {
		ret = recentInfoGetApplicationInfoFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].String(false)
	out1 := outArgs[1].Uint32()
	out2 := outArgs[2].Int64()

	return retGo, out0, out1, out2
}

var recentInfoGetApplicationsFunction *gi.Function
var recentInfoGetApplicationsFunction_Once sync.Once

func recentInfoGetApplicationsFunction_Set() error {
	var err error
	recentInfoGetApplicationsFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetApplicationsFunction, err = recentInfoStruct.InvokerNew("get_applications")
	})
	return err
}

// GetApplications is a representation of the C type gtk_recent_info_get_applications.
func (recv *RecentInfo) GetApplications() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := recentInfoGetApplicationsFunction_Set()
	if err == nil {
		recentInfoGetApplicationsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()

	return out0
}

var recentInfoGetDescriptionFunction *gi.Function
var recentInfoGetDescriptionFunction_Once sync.Once

func recentInfoGetDescriptionFunction_Set() error {
	var err error
	recentInfoGetDescriptionFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetDescriptionFunction, err = recentInfoStruct.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type gtk_recent_info_get_description.
func (recv *RecentInfo) GetDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetDescriptionFunction_Set()
	if err == nil {
		ret = recentInfoGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var recentInfoGetDisplayNameFunction *gi.Function
var recentInfoGetDisplayNameFunction_Once sync.Once

func recentInfoGetDisplayNameFunction_Set() error {
	var err error
	recentInfoGetDisplayNameFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetDisplayNameFunction, err = recentInfoStruct.InvokerNew("get_display_name")
	})
	return err
}

// GetDisplayName is a representation of the C type gtk_recent_info_get_display_name.
func (recv *RecentInfo) GetDisplayName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetDisplayNameFunction_Set()
	if err == nil {
		ret = recentInfoGetDisplayNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_recent_info_get_gicon' : return type 'Gio.Icon' not supported

var recentInfoGetGroupsFunction *gi.Function
var recentInfoGetGroupsFunction_Once sync.Once

func recentInfoGetGroupsFunction_Set() error {
	var err error
	recentInfoGetGroupsFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetGroupsFunction, err = recentInfoStruct.InvokerNew("get_groups")
	})
	return err
}

// GetGroups is a representation of the C type gtk_recent_info_get_groups.
func (recv *RecentInfo) GetGroups() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := recentInfoGetGroupsFunction_Set()
	if err == nil {
		recentInfoGetGroupsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()

	return out0
}

// UNSUPPORTED : C value 'gtk_recent_info_get_icon' : return type 'GdkPixbuf.Pixbuf' not supported

var recentInfoGetMimeTypeFunction *gi.Function
var recentInfoGetMimeTypeFunction_Once sync.Once

func recentInfoGetMimeTypeFunction_Set() error {
	var err error
	recentInfoGetMimeTypeFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetMimeTypeFunction, err = recentInfoStruct.InvokerNew("get_mime_type")
	})
	return err
}

// GetMimeType is a representation of the C type gtk_recent_info_get_mime_type.
func (recv *RecentInfo) GetMimeType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetMimeTypeFunction_Set()
	if err == nil {
		ret = recentInfoGetMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var recentInfoGetModifiedFunction *gi.Function
var recentInfoGetModifiedFunction_Once sync.Once

func recentInfoGetModifiedFunction_Set() error {
	var err error
	recentInfoGetModifiedFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetModifiedFunction, err = recentInfoStruct.InvokerNew("get_modified")
	})
	return err
}

// GetModified is a representation of the C type gtk_recent_info_get_modified.
func (recv *RecentInfo) GetModified() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetModifiedFunction_Set()
	if err == nil {
		ret = recentInfoGetModifiedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var recentInfoGetPrivateHintFunction *gi.Function
var recentInfoGetPrivateHintFunction_Once sync.Once

func recentInfoGetPrivateHintFunction_Set() error {
	var err error
	recentInfoGetPrivateHintFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetPrivateHintFunction, err = recentInfoStruct.InvokerNew("get_private_hint")
	})
	return err
}

// GetPrivateHint is a representation of the C type gtk_recent_info_get_private_hint.
func (recv *RecentInfo) GetPrivateHint() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetPrivateHintFunction_Set()
	if err == nil {
		ret = recentInfoGetPrivateHintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentInfoGetShortNameFunction *gi.Function
var recentInfoGetShortNameFunction_Once sync.Once

func recentInfoGetShortNameFunction_Set() error {
	var err error
	recentInfoGetShortNameFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetShortNameFunction, err = recentInfoStruct.InvokerNew("get_short_name")
	})
	return err
}

// GetShortName is a representation of the C type gtk_recent_info_get_short_name.
func (recv *RecentInfo) GetShortName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetShortNameFunction_Set()
	if err == nil {
		ret = recentInfoGetShortNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var recentInfoGetUriFunction *gi.Function
var recentInfoGetUriFunction_Once sync.Once

func recentInfoGetUriFunction_Set() error {
	var err error
	recentInfoGetUriFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetUriFunction, err = recentInfoStruct.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type gtk_recent_info_get_uri.
func (recv *RecentInfo) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetUriFunction_Set()
	if err == nil {
		ret = recentInfoGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var recentInfoGetUriDisplayFunction *gi.Function
var recentInfoGetUriDisplayFunction_Once sync.Once

func recentInfoGetUriDisplayFunction_Set() error {
	var err error
	recentInfoGetUriDisplayFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetUriDisplayFunction, err = recentInfoStruct.InvokerNew("get_uri_display")
	})
	return err
}

// GetUriDisplay is a representation of the C type gtk_recent_info_get_uri_display.
func (recv *RecentInfo) GetUriDisplay() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetUriDisplayFunction_Set()
	if err == nil {
		ret = recentInfoGetUriDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var recentInfoGetVisitedFunction *gi.Function
var recentInfoGetVisitedFunction_Once sync.Once

func recentInfoGetVisitedFunction_Set() error {
	var err error
	recentInfoGetVisitedFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoGetVisitedFunction, err = recentInfoStruct.InvokerNew("get_visited")
	})
	return err
}

// GetVisited is a representation of the C type gtk_recent_info_get_visited.
func (recv *RecentInfo) GetVisited() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoGetVisitedFunction_Set()
	if err == nil {
		ret = recentInfoGetVisitedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var recentInfoHasApplicationFunction *gi.Function
var recentInfoHasApplicationFunction_Once sync.Once

func recentInfoHasApplicationFunction_Set() error {
	var err error
	recentInfoHasApplicationFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoHasApplicationFunction, err = recentInfoStruct.InvokerNew("has_application")
	})
	return err
}

// HasApplication is a representation of the C type gtk_recent_info_has_application.
func (recv *RecentInfo) HasApplication(appName string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(appName)

	var ret gi.Argument

	err := recentInfoHasApplicationFunction_Set()
	if err == nil {
		ret = recentInfoHasApplicationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentInfoHasGroupFunction *gi.Function
var recentInfoHasGroupFunction_Once sync.Once

func recentInfoHasGroupFunction_Set() error {
	var err error
	recentInfoHasGroupFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoHasGroupFunction, err = recentInfoStruct.InvokerNew("has_group")
	})
	return err
}

// HasGroup is a representation of the C type gtk_recent_info_has_group.
func (recv *RecentInfo) HasGroup(groupName string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(groupName)

	var ret gi.Argument

	err := recentInfoHasGroupFunction_Set()
	if err == nil {
		ret = recentInfoHasGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentInfoIsLocalFunction *gi.Function
var recentInfoIsLocalFunction_Once sync.Once

func recentInfoIsLocalFunction_Set() error {
	var err error
	recentInfoIsLocalFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoIsLocalFunction, err = recentInfoStruct.InvokerNew("is_local")
	})
	return err
}

// IsLocal is a representation of the C type gtk_recent_info_is_local.
func (recv *RecentInfo) IsLocal() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoIsLocalFunction_Set()
	if err == nil {
		ret = recentInfoIsLocalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentInfoLastApplicationFunction *gi.Function
var recentInfoLastApplicationFunction_Once sync.Once

func recentInfoLastApplicationFunction_Set() error {
	var err error
	recentInfoLastApplicationFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoLastApplicationFunction, err = recentInfoStruct.InvokerNew("last_application")
	})
	return err
}

// LastApplication is a representation of the C type gtk_recent_info_last_application.
func (recv *RecentInfo) LastApplication() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoLastApplicationFunction_Set()
	if err == nil {
		ret = recentInfoLastApplicationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var recentInfoMatchFunction *gi.Function
var recentInfoMatchFunction_Once sync.Once

func recentInfoMatchFunction_Set() error {
	var err error
	recentInfoMatchFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoMatchFunction, err = recentInfoStruct.InvokerNew("match")
	})
	return err
}

// Match is a representation of the C type gtk_recent_info_match.
func (recv *RecentInfo) Match(infoB *RecentInfo) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(infoB.native)

	var ret gi.Argument

	err := recentInfoMatchFunction_Set()
	if err == nil {
		ret = recentInfoMatchFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentInfoRefFunction *gi.Function
var recentInfoRefFunction_Once sync.Once

func recentInfoRefFunction_Set() error {
	var err error
	recentInfoRefFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoRefFunction, err = recentInfoStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gtk_recent_info_ref.
func (recv *RecentInfo) Ref() *RecentInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := recentInfoRefFunction_Set()
	if err == nil {
		ret = recentInfoRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &RecentInfo{native: ret.Pointer()}

	return retGo
}

var recentInfoUnrefFunction *gi.Function
var recentInfoUnrefFunction_Once sync.Once

func recentInfoUnrefFunction_Set() error {
	var err error
	recentInfoUnrefFunction_Once.Do(func() {
		err = recentInfoStruct_Set()
		if err != nil {
			return
		}
		recentInfoUnrefFunction, err = recentInfoStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gtk_recent_info_unref.
func (recv *RecentInfo) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := recentInfoUnrefFunction_Set()
	if err == nil {
		recentInfoUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentManagerClassStruct *gi.Struct
var recentManagerClassStruct_Once sync.Once

func recentManagerClassStruct_Set() error {
	var err error
	recentManagerClassStruct_Once.Do(func() {
		recentManagerClassStruct, err = gi.StructNew("Gtk", "RecentManagerClass")
	})
	return err
}

type RecentManagerClass struct {
	native uintptr
}

var recentManagerPrivateStruct *gi.Struct
var recentManagerPrivateStruct_Once sync.Once

func recentManagerPrivateStruct_Set() error {
	var err error
	recentManagerPrivateStruct_Once.Do(func() {
		recentManagerPrivateStruct, err = gi.StructNew("Gtk", "RecentManagerPrivate")
	})
	return err
}

type RecentManagerPrivate struct {
	native uintptr
}

var rendererCellAccessibleClassStruct *gi.Struct
var rendererCellAccessibleClassStruct_Once sync.Once

func rendererCellAccessibleClassStruct_Set() error {
	var err error
	rendererCellAccessibleClassStruct_Once.Do(func() {
		rendererCellAccessibleClassStruct, err = gi.StructNew("Gtk", "RendererCellAccessibleClass")
	})
	return err
}

type RendererCellAccessibleClass struct {
	native uintptr
}

var rendererCellAccessiblePrivateStruct *gi.Struct
var rendererCellAccessiblePrivateStruct_Once sync.Once

func rendererCellAccessiblePrivateStruct_Set() error {
	var err error
	rendererCellAccessiblePrivateStruct_Once.Do(func() {
		rendererCellAccessiblePrivateStruct, err = gi.StructNew("Gtk", "RendererCellAccessiblePrivate")
	})
	return err
}

type RendererCellAccessiblePrivate struct {
	native uintptr
}

var requestedSizeStruct *gi.Struct
var requestedSizeStruct_Once sync.Once

func requestedSizeStruct_Set() error {
	var err error
	requestedSizeStruct_Once.Do(func() {
		requestedSizeStruct, err = gi.StructNew("Gtk", "RequestedSize")
	})
	return err
}

type RequestedSize struct {
	native uintptr
}

var requisitionStruct *gi.Struct
var requisitionStruct_Once sync.Once

func requisitionStruct_Set() error {
	var err error
	requisitionStruct_Once.Do(func() {
		requisitionStruct, err = gi.StructNew("Gtk", "Requisition")
	})
	return err
}

type Requisition struct {
	native uintptr
}

var requisitionNewFunction *gi.Function
var requisitionNewFunction_Once sync.Once

func requisitionNewFunction_Set() error {
	var err error
	requisitionNewFunction_Once.Do(func() {
		err = requisitionStruct_Set()
		if err != nil {
			return
		}
		requisitionNewFunction, err = requisitionStruct.InvokerNew("new")
	})
	return err
}

// RequisitionNew is a representation of the C type gtk_requisition_new.
func RequisitionNew() *Requisition {

	var ret gi.Argument

	err := requisitionNewFunction_Set()
	if err == nil {
		ret = requisitionNewFunction.Invoke(nil, nil)
	}

	retGo := &Requisition{native: ret.Pointer()}

	return retGo
}

var requisitionCopyFunction *gi.Function
var requisitionCopyFunction_Once sync.Once

func requisitionCopyFunction_Set() error {
	var err error
	requisitionCopyFunction_Once.Do(func() {
		err = requisitionStruct_Set()
		if err != nil {
			return
		}
		requisitionCopyFunction, err = requisitionStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_requisition_copy.
func (recv *Requisition) Copy() *Requisition {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := requisitionCopyFunction_Set()
	if err == nil {
		ret = requisitionCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Requisition{native: ret.Pointer()}

	return retGo
}

var requisitionFreeFunction *gi.Function
var requisitionFreeFunction_Once sync.Once

func requisitionFreeFunction_Set() error {
	var err error
	requisitionFreeFunction_Once.Do(func() {
		err = requisitionStruct_Set()
		if err != nil {
			return
		}
		requisitionFreeFunction, err = requisitionStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_requisition_free.
func (recv *Requisition) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := requisitionFreeFunction_Set()
	if err == nil {
		requisitionFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var revealerClassStruct *gi.Struct
var revealerClassStruct_Once sync.Once

func revealerClassStruct_Set() error {
	var err error
	revealerClassStruct_Once.Do(func() {
		revealerClassStruct, err = gi.StructNew("Gtk", "RevealerClass")
	})
	return err
}

type RevealerClass struct {
	native uintptr
}

var scaleAccessibleClassStruct *gi.Struct
var scaleAccessibleClassStruct_Once sync.Once

func scaleAccessibleClassStruct_Set() error {
	var err error
	scaleAccessibleClassStruct_Once.Do(func() {
		scaleAccessibleClassStruct, err = gi.StructNew("Gtk", "ScaleAccessibleClass")
	})
	return err
}

type ScaleAccessibleClass struct {
	native uintptr
}

var scaleAccessiblePrivateStruct *gi.Struct
var scaleAccessiblePrivateStruct_Once sync.Once

func scaleAccessiblePrivateStruct_Set() error {
	var err error
	scaleAccessiblePrivateStruct_Once.Do(func() {
		scaleAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ScaleAccessiblePrivate")
	})
	return err
}

type ScaleAccessiblePrivate struct {
	native uintptr
}

var scaleButtonAccessibleClassStruct *gi.Struct
var scaleButtonAccessibleClassStruct_Once sync.Once

func scaleButtonAccessibleClassStruct_Set() error {
	var err error
	scaleButtonAccessibleClassStruct_Once.Do(func() {
		scaleButtonAccessibleClassStruct, err = gi.StructNew("Gtk", "ScaleButtonAccessibleClass")
	})
	return err
}

type ScaleButtonAccessibleClass struct {
	native uintptr
}

var scaleButtonAccessiblePrivateStruct *gi.Struct
var scaleButtonAccessiblePrivateStruct_Once sync.Once

func scaleButtonAccessiblePrivateStruct_Set() error {
	var err error
	scaleButtonAccessiblePrivateStruct_Once.Do(func() {
		scaleButtonAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ScaleButtonAccessiblePrivate")
	})
	return err
}

type ScaleButtonAccessiblePrivate struct {
	native uintptr
}

var scaleButtonClassStruct *gi.Struct
var scaleButtonClassStruct_Once sync.Once

func scaleButtonClassStruct_Set() error {
	var err error
	scaleButtonClassStruct_Once.Do(func() {
		scaleButtonClassStruct, err = gi.StructNew("Gtk", "ScaleButtonClass")
	})
	return err
}

type ScaleButtonClass struct {
	native uintptr
}

var scaleButtonPrivateStruct *gi.Struct
var scaleButtonPrivateStruct_Once sync.Once

func scaleButtonPrivateStruct_Set() error {
	var err error
	scaleButtonPrivateStruct_Once.Do(func() {
		scaleButtonPrivateStruct, err = gi.StructNew("Gtk", "ScaleButtonPrivate")
	})
	return err
}

type ScaleButtonPrivate struct {
	native uintptr
}

var scaleClassStruct *gi.Struct
var scaleClassStruct_Once sync.Once

func scaleClassStruct_Set() error {
	var err error
	scaleClassStruct_Once.Do(func() {
		scaleClassStruct, err = gi.StructNew("Gtk", "ScaleClass")
	})
	return err
}

type ScaleClass struct {
	native uintptr
}

var scalePrivateStruct *gi.Struct
var scalePrivateStruct_Once sync.Once

func scalePrivateStruct_Set() error {
	var err error
	scalePrivateStruct_Once.Do(func() {
		scalePrivateStruct, err = gi.StructNew("Gtk", "ScalePrivate")
	})
	return err
}

type ScalePrivate struct {
	native uintptr
}

var scrollableInterfaceStruct *gi.Struct
var scrollableInterfaceStruct_Once sync.Once

func scrollableInterfaceStruct_Set() error {
	var err error
	scrollableInterfaceStruct_Once.Do(func() {
		scrollableInterfaceStruct, err = gi.StructNew("Gtk", "ScrollableInterface")
	})
	return err
}

type ScrollableInterface struct {
	native uintptr
}

var scrollbarClassStruct *gi.Struct
var scrollbarClassStruct_Once sync.Once

func scrollbarClassStruct_Set() error {
	var err error
	scrollbarClassStruct_Once.Do(func() {
		scrollbarClassStruct, err = gi.StructNew("Gtk", "ScrollbarClass")
	})
	return err
}

type ScrollbarClass struct {
	native uintptr
}

var scrolledWindowAccessibleClassStruct *gi.Struct
var scrolledWindowAccessibleClassStruct_Once sync.Once

func scrolledWindowAccessibleClassStruct_Set() error {
	var err error
	scrolledWindowAccessibleClassStruct_Once.Do(func() {
		scrolledWindowAccessibleClassStruct, err = gi.StructNew("Gtk", "ScrolledWindowAccessibleClass")
	})
	return err
}

type ScrolledWindowAccessibleClass struct {
	native uintptr
}

var scrolledWindowAccessiblePrivateStruct *gi.Struct
var scrolledWindowAccessiblePrivateStruct_Once sync.Once

func scrolledWindowAccessiblePrivateStruct_Set() error {
	var err error
	scrolledWindowAccessiblePrivateStruct_Once.Do(func() {
		scrolledWindowAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ScrolledWindowAccessiblePrivate")
	})
	return err
}

type ScrolledWindowAccessiblePrivate struct {
	native uintptr
}

var scrolledWindowClassStruct *gi.Struct
var scrolledWindowClassStruct_Once sync.Once

func scrolledWindowClassStruct_Set() error {
	var err error
	scrolledWindowClassStruct_Once.Do(func() {
		scrolledWindowClassStruct, err = gi.StructNew("Gtk", "ScrolledWindowClass")
	})
	return err
}

type ScrolledWindowClass struct {
	native uintptr
}

var scrolledWindowPrivateStruct *gi.Struct
var scrolledWindowPrivateStruct_Once sync.Once

func scrolledWindowPrivateStruct_Set() error {
	var err error
	scrolledWindowPrivateStruct_Once.Do(func() {
		scrolledWindowPrivateStruct, err = gi.StructNew("Gtk", "ScrolledWindowPrivate")
	})
	return err
}

type ScrolledWindowPrivate struct {
	native uintptr
}

var searchBarClassStruct *gi.Struct
var searchBarClassStruct_Once sync.Once

func searchBarClassStruct_Set() error {
	var err error
	searchBarClassStruct_Once.Do(func() {
		searchBarClassStruct, err = gi.StructNew("Gtk", "SearchBarClass")
	})
	return err
}

type SearchBarClass struct {
	native uintptr
}

var searchEntryClassStruct *gi.Struct
var searchEntryClassStruct_Once sync.Once

func searchEntryClassStruct_Set() error {
	var err error
	searchEntryClassStruct_Once.Do(func() {
		searchEntryClassStruct, err = gi.StructNew("Gtk", "SearchEntryClass")
	})
	return err
}

type SearchEntryClass struct {
	native uintptr
}

var selectionDataStruct *gi.Struct
var selectionDataStruct_Once sync.Once

func selectionDataStruct_Set() error {
	var err error
	selectionDataStruct_Once.Do(func() {
		selectionDataStruct, err = gi.StructNew("Gtk", "SelectionData")
	})
	return err
}

type SelectionData struct {
	native uintptr
}

var selectionDataCopyFunction *gi.Function
var selectionDataCopyFunction_Once sync.Once

func selectionDataCopyFunction_Set() error {
	var err error
	selectionDataCopyFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataCopyFunction, err = selectionDataStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_selection_data_copy.
func (recv *SelectionData) Copy() *SelectionData {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := selectionDataCopyFunction_Set()
	if err == nil {
		ret = selectionDataCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SelectionData{native: ret.Pointer()}

	return retGo
}

var selectionDataFreeFunction *gi.Function
var selectionDataFreeFunction_Once sync.Once

func selectionDataFreeFunction_Set() error {
	var err error
	selectionDataFreeFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataFreeFunction, err = selectionDataStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_selection_data_free.
func (recv *SelectionData) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := selectionDataFreeFunction_Set()
	if err == nil {
		selectionDataFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var selectionDataGetDataFunction *gi.Function
var selectionDataGetDataFunction_Once sync.Once

func selectionDataGetDataFunction_Set() error {
	var err error
	selectionDataGetDataFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataGetDataFunction, err = selectionDataStruct.InvokerNew("get_data")
	})
	return err
}

// GetData is a representation of the C type gtk_selection_data_get_data.
func (recv *SelectionData) GetData() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := selectionDataGetDataFunction_Set()
	if err == nil {
		selectionDataGetDataFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_selection_data_get_data_type' : return type 'Gdk.Atom' not supported

var selectionDataGetDataWithLengthFunction *gi.Function
var selectionDataGetDataWithLengthFunction_Once sync.Once

func selectionDataGetDataWithLengthFunction_Set() error {
	var err error
	selectionDataGetDataWithLengthFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataGetDataWithLengthFunction, err = selectionDataStruct.InvokerNew("get_data_with_length")
	})
	return err
}

// GetDataWithLength is a representation of the C type gtk_selection_data_get_data_with_length.
func (recv *SelectionData) GetDataWithLength() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := selectionDataGetDataWithLengthFunction_Set()
	if err == nil {
		selectionDataGetDataWithLengthFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

// UNSUPPORTED : C value 'gtk_selection_data_get_display' : return type 'Gdk.Display' not supported

var selectionDataGetFormatFunction *gi.Function
var selectionDataGetFormatFunction_Once sync.Once

func selectionDataGetFormatFunction_Set() error {
	var err error
	selectionDataGetFormatFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataGetFormatFunction, err = selectionDataStruct.InvokerNew("get_format")
	})
	return err
}

// GetFormat is a representation of the C type gtk_selection_data_get_format.
func (recv *SelectionData) GetFormat() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := selectionDataGetFormatFunction_Set()
	if err == nil {
		ret = selectionDataGetFormatFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var selectionDataGetLengthFunction *gi.Function
var selectionDataGetLengthFunction_Once sync.Once

func selectionDataGetLengthFunction_Set() error {
	var err error
	selectionDataGetLengthFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataGetLengthFunction, err = selectionDataStruct.InvokerNew("get_length")
	})
	return err
}

// GetLength is a representation of the C type gtk_selection_data_get_length.
func (recv *SelectionData) GetLength() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := selectionDataGetLengthFunction_Set()
	if err == nil {
		ret = selectionDataGetLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_selection_data_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gtk_selection_data_get_selection' : return type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_data_get_target' : return type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_data_get_targets' : parameter 'targets' has no type

var selectionDataGetTextFunction *gi.Function
var selectionDataGetTextFunction_Once sync.Once

func selectionDataGetTextFunction_Set() error {
	var err error
	selectionDataGetTextFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataGetTextFunction, err = selectionDataStruct.InvokerNew("get_text")
	})
	return err
}

// GetText is a representation of the C type gtk_selection_data_get_text.
func (recv *SelectionData) GetText() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := selectionDataGetTextFunction_Set()
	if err == nil {
		ret = selectionDataGetTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var selectionDataGetUrisFunction *gi.Function
var selectionDataGetUrisFunction_Once sync.Once

func selectionDataGetUrisFunction_Set() error {
	var err error
	selectionDataGetUrisFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataGetUrisFunction, err = selectionDataStruct.InvokerNew("get_uris")
	})
	return err
}

// GetUris is a representation of the C type gtk_selection_data_get_uris.
func (recv *SelectionData) GetUris() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := selectionDataGetUrisFunction_Set()
	if err == nil {
		selectionDataGetUrisFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_selection_data_set' : parameter 'type' of type 'Gdk.Atom' not supported

// UNSUPPORTED : C value 'gtk_selection_data_set_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

var selectionDataSetTextFunction *gi.Function
var selectionDataSetTextFunction_Once sync.Once

func selectionDataSetTextFunction_Set() error {
	var err error
	selectionDataSetTextFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataSetTextFunction, err = selectionDataStruct.InvokerNew("set_text")
	})
	return err
}

// SetText is a representation of the C type gtk_selection_data_set_text.
func (recv *SelectionData) SetText(str string, len int32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(str)
	inArgs[2].SetInt32(len)

	var ret gi.Argument

	err := selectionDataSetTextFunction_Set()
	if err == nil {
		ret = selectionDataSetTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_selection_data_set_uris' : parameter 'uris' has no type

var selectionDataTargetsIncludeImageFunction *gi.Function
var selectionDataTargetsIncludeImageFunction_Once sync.Once

func selectionDataTargetsIncludeImageFunction_Set() error {
	var err error
	selectionDataTargetsIncludeImageFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataTargetsIncludeImageFunction, err = selectionDataStruct.InvokerNew("targets_include_image")
	})
	return err
}

// TargetsIncludeImage is a representation of the C type gtk_selection_data_targets_include_image.
func (recv *SelectionData) TargetsIncludeImage(writable bool) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(writable)

	var ret gi.Argument

	err := selectionDataTargetsIncludeImageFunction_Set()
	if err == nil {
		ret = selectionDataTargetsIncludeImageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_selection_data_targets_include_rich_text' : parameter 'buffer' of type 'TextBuffer' not supported

var selectionDataTargetsIncludeTextFunction *gi.Function
var selectionDataTargetsIncludeTextFunction_Once sync.Once

func selectionDataTargetsIncludeTextFunction_Set() error {
	var err error
	selectionDataTargetsIncludeTextFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataTargetsIncludeTextFunction, err = selectionDataStruct.InvokerNew("targets_include_text")
	})
	return err
}

// TargetsIncludeText is a representation of the C type gtk_selection_data_targets_include_text.
func (recv *SelectionData) TargetsIncludeText() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := selectionDataTargetsIncludeTextFunction_Set()
	if err == nil {
		ret = selectionDataTargetsIncludeTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var selectionDataTargetsIncludeUriFunction *gi.Function
var selectionDataTargetsIncludeUriFunction_Once sync.Once

func selectionDataTargetsIncludeUriFunction_Set() error {
	var err error
	selectionDataTargetsIncludeUriFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataTargetsIncludeUriFunction, err = selectionDataStruct.InvokerNew("targets_include_uri")
	})
	return err
}

// TargetsIncludeUri is a representation of the C type gtk_selection_data_targets_include_uri.
func (recv *SelectionData) TargetsIncludeUri() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := selectionDataTargetsIncludeUriFunction_Set()
	if err == nil {
		ret = selectionDataTargetsIncludeUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var separatorClassStruct *gi.Struct
var separatorClassStruct_Once sync.Once

func separatorClassStruct_Set() error {
	var err error
	separatorClassStruct_Once.Do(func() {
		separatorClassStruct, err = gi.StructNew("Gtk", "SeparatorClass")
	})
	return err
}

type SeparatorClass struct {
	native uintptr
}

var separatorMenuItemClassStruct *gi.Struct
var separatorMenuItemClassStruct_Once sync.Once

func separatorMenuItemClassStruct_Set() error {
	var err error
	separatorMenuItemClassStruct_Once.Do(func() {
		separatorMenuItemClassStruct, err = gi.StructNew("Gtk", "SeparatorMenuItemClass")
	})
	return err
}

type SeparatorMenuItemClass struct {
	native uintptr
}

var separatorPrivateStruct *gi.Struct
var separatorPrivateStruct_Once sync.Once

func separatorPrivateStruct_Set() error {
	var err error
	separatorPrivateStruct_Once.Do(func() {
		separatorPrivateStruct, err = gi.StructNew("Gtk", "SeparatorPrivate")
	})
	return err
}

type SeparatorPrivate struct {
	native uintptr
}

var separatorToolItemClassStruct *gi.Struct
var separatorToolItemClassStruct_Once sync.Once

func separatorToolItemClassStruct_Set() error {
	var err error
	separatorToolItemClassStruct_Once.Do(func() {
		separatorToolItemClassStruct, err = gi.StructNew("Gtk", "SeparatorToolItemClass")
	})
	return err
}

type SeparatorToolItemClass struct {
	native uintptr
}

var separatorToolItemPrivateStruct *gi.Struct
var separatorToolItemPrivateStruct_Once sync.Once

func separatorToolItemPrivateStruct_Set() error {
	var err error
	separatorToolItemPrivateStruct_Once.Do(func() {
		separatorToolItemPrivateStruct, err = gi.StructNew("Gtk", "SeparatorToolItemPrivate")
	})
	return err
}

type SeparatorToolItemPrivate struct {
	native uintptr
}

var settingsClassStruct *gi.Struct
var settingsClassStruct_Once sync.Once

func settingsClassStruct_Set() error {
	var err error
	settingsClassStruct_Once.Do(func() {
		settingsClassStruct, err = gi.StructNew("Gtk", "SettingsClass")
	})
	return err
}

type SettingsClass struct {
	native uintptr
}

var settingsPrivateStruct *gi.Struct
var settingsPrivateStruct_Once sync.Once

func settingsPrivateStruct_Set() error {
	var err error
	settingsPrivateStruct_Once.Do(func() {
		settingsPrivateStruct, err = gi.StructNew("Gtk", "SettingsPrivate")
	})
	return err
}

type SettingsPrivate struct {
	native uintptr
}

var settingsValueStruct *gi.Struct
var settingsValueStruct_Once sync.Once

func settingsValueStruct_Set() error {
	var err error
	settingsValueStruct_Once.Do(func() {
		settingsValueStruct, err = gi.StructNew("Gtk", "SettingsValue")
	})
	return err
}

type SettingsValue struct {
	native uintptr
}

var shortcutLabelClassStruct *gi.Struct
var shortcutLabelClassStruct_Once sync.Once

func shortcutLabelClassStruct_Set() error {
	var err error
	shortcutLabelClassStruct_Once.Do(func() {
		shortcutLabelClassStruct, err = gi.StructNew("Gtk", "ShortcutLabelClass")
	})
	return err
}

type ShortcutLabelClass struct {
	native uintptr
}

var shortcutsGroupClassStruct *gi.Struct
var shortcutsGroupClassStruct_Once sync.Once

func shortcutsGroupClassStruct_Set() error {
	var err error
	shortcutsGroupClassStruct_Once.Do(func() {
		shortcutsGroupClassStruct, err = gi.StructNew("Gtk", "ShortcutsGroupClass")
	})
	return err
}

type ShortcutsGroupClass struct {
	native uintptr
}

var shortcutsSectionClassStruct *gi.Struct
var shortcutsSectionClassStruct_Once sync.Once

func shortcutsSectionClassStruct_Set() error {
	var err error
	shortcutsSectionClassStruct_Once.Do(func() {
		shortcutsSectionClassStruct, err = gi.StructNew("Gtk", "ShortcutsSectionClass")
	})
	return err
}

type ShortcutsSectionClass struct {
	native uintptr
}

var shortcutsShortcutClassStruct *gi.Struct
var shortcutsShortcutClassStruct_Once sync.Once

func shortcutsShortcutClassStruct_Set() error {
	var err error
	shortcutsShortcutClassStruct_Once.Do(func() {
		shortcutsShortcutClassStruct, err = gi.StructNew("Gtk", "ShortcutsShortcutClass")
	})
	return err
}

type ShortcutsShortcutClass struct {
	native uintptr
}

var shortcutsWindowClassStruct *gi.Struct
var shortcutsWindowClassStruct_Once sync.Once

func shortcutsWindowClassStruct_Set() error {
	var err error
	shortcutsWindowClassStruct_Once.Do(func() {
		shortcutsWindowClassStruct, err = gi.StructNew("Gtk", "ShortcutsWindowClass")
	})
	return err
}

type ShortcutsWindowClass struct {
	native uintptr
}

var sizeGroupClassStruct *gi.Struct
var sizeGroupClassStruct_Once sync.Once

func sizeGroupClassStruct_Set() error {
	var err error
	sizeGroupClassStruct_Once.Do(func() {
		sizeGroupClassStruct, err = gi.StructNew("Gtk", "SizeGroupClass")
	})
	return err
}

type SizeGroupClass struct {
	native uintptr
}

var sizeGroupPrivateStruct *gi.Struct
var sizeGroupPrivateStruct_Once sync.Once

func sizeGroupPrivateStruct_Set() error {
	var err error
	sizeGroupPrivateStruct_Once.Do(func() {
		sizeGroupPrivateStruct, err = gi.StructNew("Gtk", "SizeGroupPrivate")
	})
	return err
}

type SizeGroupPrivate struct {
	native uintptr
}

var socketClassStruct *gi.Struct
var socketClassStruct_Once sync.Once

func socketClassStruct_Set() error {
	var err error
	socketClassStruct_Once.Do(func() {
		socketClassStruct, err = gi.StructNew("Gtk", "SocketClass")
	})
	return err
}

type SocketClass struct {
	native uintptr
}

var socketPrivateStruct *gi.Struct
var socketPrivateStruct_Once sync.Once

func socketPrivateStruct_Set() error {
	var err error
	socketPrivateStruct_Once.Do(func() {
		socketPrivateStruct, err = gi.StructNew("Gtk", "SocketPrivate")
	})
	return err
}

type SocketPrivate struct {
	native uintptr
}

var spinButtonAccessibleClassStruct *gi.Struct
var spinButtonAccessibleClassStruct_Once sync.Once

func spinButtonAccessibleClassStruct_Set() error {
	var err error
	spinButtonAccessibleClassStruct_Once.Do(func() {
		spinButtonAccessibleClassStruct, err = gi.StructNew("Gtk", "SpinButtonAccessibleClass")
	})
	return err
}

type SpinButtonAccessibleClass struct {
	native uintptr
}

var spinButtonAccessiblePrivateStruct *gi.Struct
var spinButtonAccessiblePrivateStruct_Once sync.Once

func spinButtonAccessiblePrivateStruct_Set() error {
	var err error
	spinButtonAccessiblePrivateStruct_Once.Do(func() {
		spinButtonAccessiblePrivateStruct, err = gi.StructNew("Gtk", "SpinButtonAccessiblePrivate")
	})
	return err
}

type SpinButtonAccessiblePrivate struct {
	native uintptr
}

var spinButtonClassStruct *gi.Struct
var spinButtonClassStruct_Once sync.Once

func spinButtonClassStruct_Set() error {
	var err error
	spinButtonClassStruct_Once.Do(func() {
		spinButtonClassStruct, err = gi.StructNew("Gtk", "SpinButtonClass")
	})
	return err
}

type SpinButtonClass struct {
	native uintptr
}

var spinButtonPrivateStruct *gi.Struct
var spinButtonPrivateStruct_Once sync.Once

func spinButtonPrivateStruct_Set() error {
	var err error
	spinButtonPrivateStruct_Once.Do(func() {
		spinButtonPrivateStruct, err = gi.StructNew("Gtk", "SpinButtonPrivate")
	})
	return err
}

type SpinButtonPrivate struct {
	native uintptr
}

var spinnerAccessibleClassStruct *gi.Struct
var spinnerAccessibleClassStruct_Once sync.Once

func spinnerAccessibleClassStruct_Set() error {
	var err error
	spinnerAccessibleClassStruct_Once.Do(func() {
		spinnerAccessibleClassStruct, err = gi.StructNew("Gtk", "SpinnerAccessibleClass")
	})
	return err
}

type SpinnerAccessibleClass struct {
	native uintptr
}

var spinnerAccessiblePrivateStruct *gi.Struct
var spinnerAccessiblePrivateStruct_Once sync.Once

func spinnerAccessiblePrivateStruct_Set() error {
	var err error
	spinnerAccessiblePrivateStruct_Once.Do(func() {
		spinnerAccessiblePrivateStruct, err = gi.StructNew("Gtk", "SpinnerAccessiblePrivate")
	})
	return err
}

type SpinnerAccessiblePrivate struct {
	native uintptr
}

var spinnerClassStruct *gi.Struct
var spinnerClassStruct_Once sync.Once

func spinnerClassStruct_Set() error {
	var err error
	spinnerClassStruct_Once.Do(func() {
		spinnerClassStruct, err = gi.StructNew("Gtk", "SpinnerClass")
	})
	return err
}

type SpinnerClass struct {
	native uintptr
}

var spinnerPrivateStruct *gi.Struct
var spinnerPrivateStruct_Once sync.Once

func spinnerPrivateStruct_Set() error {
	var err error
	spinnerPrivateStruct_Once.Do(func() {
		spinnerPrivateStruct, err = gi.StructNew("Gtk", "SpinnerPrivate")
	})
	return err
}

type SpinnerPrivate struct {
	native uintptr
}

var stackAccessibleClassStruct *gi.Struct
var stackAccessibleClassStruct_Once sync.Once

func stackAccessibleClassStruct_Set() error {
	var err error
	stackAccessibleClassStruct_Once.Do(func() {
		stackAccessibleClassStruct, err = gi.StructNew("Gtk", "StackAccessibleClass")
	})
	return err
}

type StackAccessibleClass struct {
	native uintptr
}

var stackClassStruct *gi.Struct
var stackClassStruct_Once sync.Once

func stackClassStruct_Set() error {
	var err error
	stackClassStruct_Once.Do(func() {
		stackClassStruct, err = gi.StructNew("Gtk", "StackClass")
	})
	return err
}

type StackClass struct {
	native uintptr
}

var stackSidebarClassStruct *gi.Struct
var stackSidebarClassStruct_Once sync.Once

func stackSidebarClassStruct_Set() error {
	var err error
	stackSidebarClassStruct_Once.Do(func() {
		stackSidebarClassStruct, err = gi.StructNew("Gtk", "StackSidebarClass")
	})
	return err
}

type StackSidebarClass struct {
	native uintptr
}

var stackSidebarPrivateStruct *gi.Struct
var stackSidebarPrivateStruct_Once sync.Once

func stackSidebarPrivateStruct_Set() error {
	var err error
	stackSidebarPrivateStruct_Once.Do(func() {
		stackSidebarPrivateStruct, err = gi.StructNew("Gtk", "StackSidebarPrivate")
	})
	return err
}

type StackSidebarPrivate struct {
	native uintptr
}

var stackSwitcherClassStruct *gi.Struct
var stackSwitcherClassStruct_Once sync.Once

func stackSwitcherClassStruct_Set() error {
	var err error
	stackSwitcherClassStruct_Once.Do(func() {
		stackSwitcherClassStruct, err = gi.StructNew("Gtk", "StackSwitcherClass")
	})
	return err
}

type StackSwitcherClass struct {
	native uintptr
}

var statusIconClassStruct *gi.Struct
var statusIconClassStruct_Once sync.Once

func statusIconClassStruct_Set() error {
	var err error
	statusIconClassStruct_Once.Do(func() {
		statusIconClassStruct, err = gi.StructNew("Gtk", "StatusIconClass")
	})
	return err
}

type StatusIconClass struct {
	native uintptr
}

var statusIconPrivateStruct *gi.Struct
var statusIconPrivateStruct_Once sync.Once

func statusIconPrivateStruct_Set() error {
	var err error
	statusIconPrivateStruct_Once.Do(func() {
		statusIconPrivateStruct, err = gi.StructNew("Gtk", "StatusIconPrivate")
	})
	return err
}

type StatusIconPrivate struct {
	native uintptr
}

var statusbarAccessibleClassStruct *gi.Struct
var statusbarAccessibleClassStruct_Once sync.Once

func statusbarAccessibleClassStruct_Set() error {
	var err error
	statusbarAccessibleClassStruct_Once.Do(func() {
		statusbarAccessibleClassStruct, err = gi.StructNew("Gtk", "StatusbarAccessibleClass")
	})
	return err
}

type StatusbarAccessibleClass struct {
	native uintptr
}

var statusbarAccessiblePrivateStruct *gi.Struct
var statusbarAccessiblePrivateStruct_Once sync.Once

func statusbarAccessiblePrivateStruct_Set() error {
	var err error
	statusbarAccessiblePrivateStruct_Once.Do(func() {
		statusbarAccessiblePrivateStruct, err = gi.StructNew("Gtk", "StatusbarAccessiblePrivate")
	})
	return err
}

type StatusbarAccessiblePrivate struct {
	native uintptr
}

var statusbarClassStruct *gi.Struct
var statusbarClassStruct_Once sync.Once

func statusbarClassStruct_Set() error {
	var err error
	statusbarClassStruct_Once.Do(func() {
		statusbarClassStruct, err = gi.StructNew("Gtk", "StatusbarClass")
	})
	return err
}

type StatusbarClass struct {
	native uintptr
}

var statusbarPrivateStruct *gi.Struct
var statusbarPrivateStruct_Once sync.Once

func statusbarPrivateStruct_Set() error {
	var err error
	statusbarPrivateStruct_Once.Do(func() {
		statusbarPrivateStruct, err = gi.StructNew("Gtk", "StatusbarPrivate")
	})
	return err
}

type StatusbarPrivate struct {
	native uintptr
}

var stockItemStruct *gi.Struct
var stockItemStruct_Once sync.Once

func stockItemStruct_Set() error {
	var err error
	stockItemStruct_Once.Do(func() {
		stockItemStruct, err = gi.StructNew("Gtk", "StockItem")
	})
	return err
}

type StockItem struct {
	native uintptr
}

var stockItemCopyFunction *gi.Function
var stockItemCopyFunction_Once sync.Once

func stockItemCopyFunction_Set() error {
	var err error
	stockItemCopyFunction_Once.Do(func() {
		err = stockItemStruct_Set()
		if err != nil {
			return
		}
		stockItemCopyFunction, err = stockItemStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_stock_item_copy.
func (recv *StockItem) Copy() *StockItem {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := stockItemCopyFunction_Set()
	if err == nil {
		ret = stockItemCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &StockItem{native: ret.Pointer()}

	return retGo
}

var stockItemFreeFunction *gi.Function
var stockItemFreeFunction_Once sync.Once

func stockItemFreeFunction_Set() error {
	var err error
	stockItemFreeFunction_Once.Do(func() {
		err = stockItemStruct_Set()
		if err != nil {
			return
		}
		stockItemFreeFunction, err = stockItemStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_stock_item_free.
func (recv *StockItem) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := stockItemFreeFunction_Set()
	if err == nil {
		stockItemFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var styleClassStruct *gi.Struct
var styleClassStruct_Once sync.Once

func styleClassStruct_Set() error {
	var err error
	styleClassStruct_Once.Do(func() {
		styleClassStruct, err = gi.StructNew("Gtk", "StyleClass")
	})
	return err
}

type StyleClass struct {
	native uintptr
}

var styleContextClassStruct *gi.Struct
var styleContextClassStruct_Once sync.Once

func styleContextClassStruct_Set() error {
	var err error
	styleContextClassStruct_Once.Do(func() {
		styleContextClassStruct, err = gi.StructNew("Gtk", "StyleContextClass")
	})
	return err
}

type StyleContextClass struct {
	native uintptr
}

var styleContextPrivateStruct *gi.Struct
var styleContextPrivateStruct_Once sync.Once

func styleContextPrivateStruct_Set() error {
	var err error
	styleContextPrivateStruct_Once.Do(func() {
		styleContextPrivateStruct, err = gi.StructNew("Gtk", "StyleContextPrivate")
	})
	return err
}

type StyleContextPrivate struct {
	native uintptr
}

var stylePropertiesClassStruct *gi.Struct
var stylePropertiesClassStruct_Once sync.Once

func stylePropertiesClassStruct_Set() error {
	var err error
	stylePropertiesClassStruct_Once.Do(func() {
		stylePropertiesClassStruct, err = gi.StructNew("Gtk", "StylePropertiesClass")
	})
	return err
}

type StylePropertiesClass struct {
	native uintptr
}

var stylePropertiesPrivateStruct *gi.Struct
var stylePropertiesPrivateStruct_Once sync.Once

func stylePropertiesPrivateStruct_Set() error {
	var err error
	stylePropertiesPrivateStruct_Once.Do(func() {
		stylePropertiesPrivateStruct, err = gi.StructNew("Gtk", "StylePropertiesPrivate")
	})
	return err
}

type StylePropertiesPrivate struct {
	native uintptr
}

var styleProviderIfaceStruct *gi.Struct
var styleProviderIfaceStruct_Once sync.Once

func styleProviderIfaceStruct_Set() error {
	var err error
	styleProviderIfaceStruct_Once.Do(func() {
		styleProviderIfaceStruct, err = gi.StructNew("Gtk", "StyleProviderIface")
	})
	return err
}

type StyleProviderIface struct {
	native uintptr
}

var switchAccessibleClassStruct *gi.Struct
var switchAccessibleClassStruct_Once sync.Once

func switchAccessibleClassStruct_Set() error {
	var err error
	switchAccessibleClassStruct_Once.Do(func() {
		switchAccessibleClassStruct, err = gi.StructNew("Gtk", "SwitchAccessibleClass")
	})
	return err
}

type SwitchAccessibleClass struct {
	native uintptr
}

var switchAccessiblePrivateStruct *gi.Struct
var switchAccessiblePrivateStruct_Once sync.Once

func switchAccessiblePrivateStruct_Set() error {
	var err error
	switchAccessiblePrivateStruct_Once.Do(func() {
		switchAccessiblePrivateStruct, err = gi.StructNew("Gtk", "SwitchAccessiblePrivate")
	})
	return err
}

type SwitchAccessiblePrivate struct {
	native uintptr
}

var switchClassStruct *gi.Struct
var switchClassStruct_Once sync.Once

func switchClassStruct_Set() error {
	var err error
	switchClassStruct_Once.Do(func() {
		switchClassStruct, err = gi.StructNew("Gtk", "SwitchClass")
	})
	return err
}

type SwitchClass struct {
	native uintptr
}

var switchPrivateStruct *gi.Struct
var switchPrivateStruct_Once sync.Once

func switchPrivateStruct_Set() error {
	var err error
	switchPrivateStruct_Once.Do(func() {
		switchPrivateStruct, err = gi.StructNew("Gtk", "SwitchPrivate")
	})
	return err
}

type SwitchPrivate struct {
	native uintptr
}

var symbolicColorStruct *gi.Struct
var symbolicColorStruct_Once sync.Once

func symbolicColorStruct_Set() error {
	var err error
	symbolicColorStruct_Once.Do(func() {
		symbolicColorStruct, err = gi.StructNew("Gtk", "SymbolicColor")
	})
	return err
}

type SymbolicColor struct {
	native uintptr
}

var symbolicColorNewAlphaFunction *gi.Function
var symbolicColorNewAlphaFunction_Once sync.Once

func symbolicColorNewAlphaFunction_Set() error {
	var err error
	symbolicColorNewAlphaFunction_Once.Do(func() {
		err = symbolicColorStruct_Set()
		if err != nil {
			return
		}
		symbolicColorNewAlphaFunction, err = symbolicColorStruct.InvokerNew("new_alpha")
	})
	return err
}

// SymbolicColorNewAlpha is a representation of the C type gtk_symbolic_color_new_alpha.
func SymbolicColorNewAlpha(color *SymbolicColor, factor float64) *SymbolicColor {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(color.native)
	inArgs[1].SetFloat64(factor)

	var ret gi.Argument

	err := symbolicColorNewAlphaFunction_Set()
	if err == nil {
		ret = symbolicColorNewAlphaFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_symbolic_color_new_literal' : parameter 'color' of type 'Gdk.RGBA' not supported

var symbolicColorNewMixFunction *gi.Function
var symbolicColorNewMixFunction_Once sync.Once

func symbolicColorNewMixFunction_Set() error {
	var err error
	symbolicColorNewMixFunction_Once.Do(func() {
		err = symbolicColorStruct_Set()
		if err != nil {
			return
		}
		symbolicColorNewMixFunction, err = symbolicColorStruct.InvokerNew("new_mix")
	})
	return err
}

// SymbolicColorNewMix is a representation of the C type gtk_symbolic_color_new_mix.
func SymbolicColorNewMix(color1 *SymbolicColor, color2 *SymbolicColor, factor float64) *SymbolicColor {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(color1.native)
	inArgs[1].SetPointer(color2.native)
	inArgs[2].SetFloat64(factor)

	var ret gi.Argument

	err := symbolicColorNewMixFunction_Set()
	if err == nil {
		ret = symbolicColorNewMixFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

var symbolicColorNewNameFunction *gi.Function
var symbolicColorNewNameFunction_Once sync.Once

func symbolicColorNewNameFunction_Set() error {
	var err error
	symbolicColorNewNameFunction_Once.Do(func() {
		err = symbolicColorStruct_Set()
		if err != nil {
			return
		}
		symbolicColorNewNameFunction, err = symbolicColorStruct.InvokerNew("new_name")
	})
	return err
}

// SymbolicColorNewName is a representation of the C type gtk_symbolic_color_new_name.
func SymbolicColorNewName(name string) *SymbolicColor {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := symbolicColorNewNameFunction_Set()
	if err == nil {
		ret = symbolicColorNewNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

var symbolicColorNewShadeFunction *gi.Function
var symbolicColorNewShadeFunction_Once sync.Once

func symbolicColorNewShadeFunction_Set() error {
	var err error
	symbolicColorNewShadeFunction_Once.Do(func() {
		err = symbolicColorStruct_Set()
		if err != nil {
			return
		}
		symbolicColorNewShadeFunction, err = symbolicColorStruct.InvokerNew("new_shade")
	})
	return err
}

// SymbolicColorNewShade is a representation of the C type gtk_symbolic_color_new_shade.
func SymbolicColorNewShade(color *SymbolicColor, factor float64) *SymbolicColor {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(color.native)
	inArgs[1].SetFloat64(factor)

	var ret gi.Argument

	err := symbolicColorNewShadeFunction_Set()
	if err == nil {
		ret = symbolicColorNewShadeFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

var symbolicColorNewWin32Function *gi.Function
var symbolicColorNewWin32Function_Once sync.Once

func symbolicColorNewWin32Function_Set() error {
	var err error
	symbolicColorNewWin32Function_Once.Do(func() {
		err = symbolicColorStruct_Set()
		if err != nil {
			return
		}
		symbolicColorNewWin32Function, err = symbolicColorStruct.InvokerNew("new_win32")
	})
	return err
}

// SymbolicColorNewWin32 is a representation of the C type gtk_symbolic_color_new_win32.
func SymbolicColorNewWin32(themeClass string, id int32) *SymbolicColor {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(themeClass)
	inArgs[1].SetInt32(id)

	var ret gi.Argument

	err := symbolicColorNewWin32Function_Set()
	if err == nil {
		ret = symbolicColorNewWin32Function.Invoke(inArgs[:], nil)
	}

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

var symbolicColorRefFunction *gi.Function
var symbolicColorRefFunction_Once sync.Once

func symbolicColorRefFunction_Set() error {
	var err error
	symbolicColorRefFunction_Once.Do(func() {
		err = symbolicColorStruct_Set()
		if err != nil {
			return
		}
		symbolicColorRefFunction, err = symbolicColorStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gtk_symbolic_color_ref.
func (recv *SymbolicColor) Ref() *SymbolicColor {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := symbolicColorRefFunction_Set()
	if err == nil {
		ret = symbolicColorRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &SymbolicColor{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_symbolic_color_resolve' : parameter 'props' of type 'StyleProperties' not supported

var symbolicColorToStringFunction *gi.Function
var symbolicColorToStringFunction_Once sync.Once

func symbolicColorToStringFunction_Set() error {
	var err error
	symbolicColorToStringFunction_Once.Do(func() {
		err = symbolicColorStruct_Set()
		if err != nil {
			return
		}
		symbolicColorToStringFunction, err = symbolicColorStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type gtk_symbolic_color_to_string.
func (recv *SymbolicColor) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := symbolicColorToStringFunction_Set()
	if err == nil {
		ret = symbolicColorToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var symbolicColorUnrefFunction *gi.Function
var symbolicColorUnrefFunction_Once sync.Once

func symbolicColorUnrefFunction_Set() error {
	var err error
	symbolicColorUnrefFunction_Once.Do(func() {
		err = symbolicColorStruct_Set()
		if err != nil {
			return
		}
		symbolicColorUnrefFunction, err = symbolicColorStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gtk_symbolic_color_unref.
func (recv *SymbolicColor) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := symbolicColorUnrefFunction_Set()
	if err == nil {
		symbolicColorUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tableChildStruct *gi.Struct
var tableChildStruct_Once sync.Once

func tableChildStruct_Set() error {
	var err error
	tableChildStruct_Once.Do(func() {
		tableChildStruct, err = gi.StructNew("Gtk", "TableChild")
	})
	return err
}

type TableChild struct {
	native uintptr
}

var tableClassStruct *gi.Struct
var tableClassStruct_Once sync.Once

func tableClassStruct_Set() error {
	var err error
	tableClassStruct_Once.Do(func() {
		tableClassStruct, err = gi.StructNew("Gtk", "TableClass")
	})
	return err
}

type TableClass struct {
	native uintptr
}

var tablePrivateStruct *gi.Struct
var tablePrivateStruct_Once sync.Once

func tablePrivateStruct_Set() error {
	var err error
	tablePrivateStruct_Once.Do(func() {
		tablePrivateStruct, err = gi.StructNew("Gtk", "TablePrivate")
	})
	return err
}

type TablePrivate struct {
	native uintptr
}

var tableRowColStruct *gi.Struct
var tableRowColStruct_Once sync.Once

func tableRowColStruct_Set() error {
	var err error
	tableRowColStruct_Once.Do(func() {
		tableRowColStruct, err = gi.StructNew("Gtk", "TableRowCol")
	})
	return err
}

type TableRowCol struct {
	native uintptr
}

var targetEntryStruct *gi.Struct
var targetEntryStruct_Once sync.Once

func targetEntryStruct_Set() error {
	var err error
	targetEntryStruct_Once.Do(func() {
		targetEntryStruct, err = gi.StructNew("Gtk", "TargetEntry")
	})
	return err
}

type TargetEntry struct {
	native uintptr
}

var targetEntryNewFunction *gi.Function
var targetEntryNewFunction_Once sync.Once

func targetEntryNewFunction_Set() error {
	var err error
	targetEntryNewFunction_Once.Do(func() {
		err = targetEntryStruct_Set()
		if err != nil {
			return
		}
		targetEntryNewFunction, err = targetEntryStruct.InvokerNew("new")
	})
	return err
}

// TargetEntryNew is a representation of the C type gtk_target_entry_new.
func TargetEntryNew(target string, flags uint32, info uint32) *TargetEntry {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(target)
	inArgs[1].SetUint32(flags)
	inArgs[2].SetUint32(info)

	var ret gi.Argument

	err := targetEntryNewFunction_Set()
	if err == nil {
		ret = targetEntryNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TargetEntry{native: ret.Pointer()}

	return retGo
}

var targetEntryCopyFunction *gi.Function
var targetEntryCopyFunction_Once sync.Once

func targetEntryCopyFunction_Set() error {
	var err error
	targetEntryCopyFunction_Once.Do(func() {
		err = targetEntryStruct_Set()
		if err != nil {
			return
		}
		targetEntryCopyFunction, err = targetEntryStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_target_entry_copy.
func (recv *TargetEntry) Copy() *TargetEntry {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := targetEntryCopyFunction_Set()
	if err == nil {
		ret = targetEntryCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TargetEntry{native: ret.Pointer()}

	return retGo
}

var targetEntryFreeFunction *gi.Function
var targetEntryFreeFunction_Once sync.Once

func targetEntryFreeFunction_Set() error {
	var err error
	targetEntryFreeFunction_Once.Do(func() {
		err = targetEntryStruct_Set()
		if err != nil {
			return
		}
		targetEntryFreeFunction, err = targetEntryStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_target_entry_free.
func (recv *TargetEntry) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := targetEntryFreeFunction_Set()
	if err == nil {
		targetEntryFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var targetListStruct *gi.Struct
var targetListStruct_Once sync.Once

func targetListStruct_Set() error {
	var err error
	targetListStruct_Once.Do(func() {
		targetListStruct, err = gi.StructNew("Gtk", "TargetList")
	})
	return err
}

type TargetList struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_target_list_new' : parameter 'targets' has no type

// UNSUPPORTED : C value 'gtk_target_list_add' : parameter 'target' of type 'Gdk.Atom' not supported

var targetListAddImageTargetsFunction *gi.Function
var targetListAddImageTargetsFunction_Once sync.Once

func targetListAddImageTargetsFunction_Set() error {
	var err error
	targetListAddImageTargetsFunction_Once.Do(func() {
		err = targetListStruct_Set()
		if err != nil {
			return
		}
		targetListAddImageTargetsFunction, err = targetListStruct.InvokerNew("add_image_targets")
	})
	return err
}

// AddImageTargets is a representation of the C type gtk_target_list_add_image_targets.
func (recv *TargetList) AddImageTargets(info uint32, writable bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(info)
	inArgs[2].SetBoolean(writable)

	err := targetListAddImageTargetsFunction_Set()
	if err == nil {
		targetListAddImageTargetsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_target_list_add_rich_text_targets' : parameter 'buffer' of type 'TextBuffer' not supported

// UNSUPPORTED : C value 'gtk_target_list_add_table' : parameter 'targets' has no type

var targetListAddTextTargetsFunction *gi.Function
var targetListAddTextTargetsFunction_Once sync.Once

func targetListAddTextTargetsFunction_Set() error {
	var err error
	targetListAddTextTargetsFunction_Once.Do(func() {
		err = targetListStruct_Set()
		if err != nil {
			return
		}
		targetListAddTextTargetsFunction, err = targetListStruct.InvokerNew("add_text_targets")
	})
	return err
}

// AddTextTargets is a representation of the C type gtk_target_list_add_text_targets.
func (recv *TargetList) AddTextTargets(info uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(info)

	err := targetListAddTextTargetsFunction_Set()
	if err == nil {
		targetListAddTextTargetsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var targetListAddUriTargetsFunction *gi.Function
var targetListAddUriTargetsFunction_Once sync.Once

func targetListAddUriTargetsFunction_Set() error {
	var err error
	targetListAddUriTargetsFunction_Once.Do(func() {
		err = targetListStruct_Set()
		if err != nil {
			return
		}
		targetListAddUriTargetsFunction, err = targetListStruct.InvokerNew("add_uri_targets")
	})
	return err
}

// AddUriTargets is a representation of the C type gtk_target_list_add_uri_targets.
func (recv *TargetList) AddUriTargets(info uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(info)

	err := targetListAddUriTargetsFunction_Set()
	if err == nil {
		targetListAddUriTargetsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_target_list_find' : parameter 'target' of type 'Gdk.Atom' not supported

var targetListRefFunction *gi.Function
var targetListRefFunction_Once sync.Once

func targetListRefFunction_Set() error {
	var err error
	targetListRefFunction_Once.Do(func() {
		err = targetListStruct_Set()
		if err != nil {
			return
		}
		targetListRefFunction, err = targetListStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gtk_target_list_ref.
func (recv *TargetList) Ref() *TargetList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := targetListRefFunction_Set()
	if err == nil {
		ret = targetListRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TargetList{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_target_list_remove' : parameter 'target' of type 'Gdk.Atom' not supported

var targetListUnrefFunction *gi.Function
var targetListUnrefFunction_Once sync.Once

func targetListUnrefFunction_Set() error {
	var err error
	targetListUnrefFunction_Once.Do(func() {
		err = targetListStruct_Set()
		if err != nil {
			return
		}
		targetListUnrefFunction, err = targetListStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gtk_target_list_unref.
func (recv *TargetList) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := targetListUnrefFunction_Set()
	if err == nil {
		targetListUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var targetPairStruct *gi.Struct
var targetPairStruct_Once sync.Once

func targetPairStruct_Set() error {
	var err error
	targetPairStruct_Once.Do(func() {
		targetPairStruct, err = gi.StructNew("Gtk", "TargetPair")
	})
	return err
}

type TargetPair struct {
	native uintptr
}

var tearoffMenuItemClassStruct *gi.Struct
var tearoffMenuItemClassStruct_Once sync.Once

func tearoffMenuItemClassStruct_Set() error {
	var err error
	tearoffMenuItemClassStruct_Once.Do(func() {
		tearoffMenuItemClassStruct, err = gi.StructNew("Gtk", "TearoffMenuItemClass")
	})
	return err
}

type TearoffMenuItemClass struct {
	native uintptr
}

var tearoffMenuItemPrivateStruct *gi.Struct
var tearoffMenuItemPrivateStruct_Once sync.Once

func tearoffMenuItemPrivateStruct_Set() error {
	var err error
	tearoffMenuItemPrivateStruct_Once.Do(func() {
		tearoffMenuItemPrivateStruct, err = gi.StructNew("Gtk", "TearoffMenuItemPrivate")
	})
	return err
}

type TearoffMenuItemPrivate struct {
	native uintptr
}

var textAppearanceStruct *gi.Struct
var textAppearanceStruct_Once sync.Once

func textAppearanceStruct_Set() error {
	var err error
	textAppearanceStruct_Once.Do(func() {
		textAppearanceStruct, err = gi.StructNew("Gtk", "TextAppearance")
	})
	return err
}

type TextAppearance struct {
	native uintptr
}

var textAttributesStruct *gi.Struct
var textAttributesStruct_Once sync.Once

func textAttributesStruct_Set() error {
	var err error
	textAttributesStruct_Once.Do(func() {
		textAttributesStruct, err = gi.StructNew("Gtk", "TextAttributes")
	})
	return err
}

type TextAttributes struct {
	native uintptr
}

var textAttributesNewFunction *gi.Function
var textAttributesNewFunction_Once sync.Once

func textAttributesNewFunction_Set() error {
	var err error
	textAttributesNewFunction_Once.Do(func() {
		err = textAttributesStruct_Set()
		if err != nil {
			return
		}
		textAttributesNewFunction, err = textAttributesStruct.InvokerNew("new")
	})
	return err
}

// TextAttributesNew is a representation of the C type gtk_text_attributes_new.
func TextAttributesNew() *TextAttributes {

	var ret gi.Argument

	err := textAttributesNewFunction_Set()
	if err == nil {
		ret = textAttributesNewFunction.Invoke(nil, nil)
	}

	retGo := &TextAttributes{native: ret.Pointer()}

	return retGo
}

var textAttributesCopyFunction *gi.Function
var textAttributesCopyFunction_Once sync.Once

func textAttributesCopyFunction_Set() error {
	var err error
	textAttributesCopyFunction_Once.Do(func() {
		err = textAttributesStruct_Set()
		if err != nil {
			return
		}
		textAttributesCopyFunction, err = textAttributesStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_text_attributes_copy.
func (recv *TextAttributes) Copy() *TextAttributes {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textAttributesCopyFunction_Set()
	if err == nil {
		ret = textAttributesCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TextAttributes{native: ret.Pointer()}

	return retGo
}

var textAttributesCopyValuesFunction *gi.Function
var textAttributesCopyValuesFunction_Once sync.Once

func textAttributesCopyValuesFunction_Set() error {
	var err error
	textAttributesCopyValuesFunction_Once.Do(func() {
		err = textAttributesStruct_Set()
		if err != nil {
			return
		}
		textAttributesCopyValuesFunction, err = textAttributesStruct.InvokerNew("copy_values")
	})
	return err
}

// CopyValues is a representation of the C type gtk_text_attributes_copy_values.
func (recv *TextAttributes) CopyValues(dest *TextAttributes) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(dest.native)

	err := textAttributesCopyValuesFunction_Set()
	if err == nil {
		textAttributesCopyValuesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textAttributesRefFunction *gi.Function
var textAttributesRefFunction_Once sync.Once

func textAttributesRefFunction_Set() error {
	var err error
	textAttributesRefFunction_Once.Do(func() {
		err = textAttributesStruct_Set()
		if err != nil {
			return
		}
		textAttributesRefFunction, err = textAttributesStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gtk_text_attributes_ref.
func (recv *TextAttributes) Ref() *TextAttributes {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textAttributesRefFunction_Set()
	if err == nil {
		ret = textAttributesRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TextAttributes{native: ret.Pointer()}

	return retGo
}

var textAttributesUnrefFunction *gi.Function
var textAttributesUnrefFunction_Once sync.Once

func textAttributesUnrefFunction_Set() error {
	var err error
	textAttributesUnrefFunction_Once.Do(func() {
		err = textAttributesStruct_Set()
		if err != nil {
			return
		}
		textAttributesUnrefFunction, err = textAttributesStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gtk_text_attributes_unref.
func (recv *TextAttributes) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := textAttributesUnrefFunction_Set()
	if err == nil {
		textAttributesUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textBTreeStruct *gi.Struct
var textBTreeStruct_Once sync.Once

func textBTreeStruct_Set() error {
	var err error
	textBTreeStruct_Once.Do(func() {
		textBTreeStruct, err = gi.StructNew("Gtk", "TextBTree")
	})
	return err
}

type TextBTree struct {
	native uintptr
}

var textBufferClassStruct *gi.Struct
var textBufferClassStruct_Once sync.Once

func textBufferClassStruct_Set() error {
	var err error
	textBufferClassStruct_Once.Do(func() {
		textBufferClassStruct, err = gi.StructNew("Gtk", "TextBufferClass")
	})
	return err
}

type TextBufferClass struct {
	native uintptr
}

var textBufferPrivateStruct *gi.Struct
var textBufferPrivateStruct_Once sync.Once

func textBufferPrivateStruct_Set() error {
	var err error
	textBufferPrivateStruct_Once.Do(func() {
		textBufferPrivateStruct, err = gi.StructNew("Gtk", "TextBufferPrivate")
	})
	return err
}

type TextBufferPrivate struct {
	native uintptr
}

var textCellAccessibleClassStruct *gi.Struct
var textCellAccessibleClassStruct_Once sync.Once

func textCellAccessibleClassStruct_Set() error {
	var err error
	textCellAccessibleClassStruct_Once.Do(func() {
		textCellAccessibleClassStruct, err = gi.StructNew("Gtk", "TextCellAccessibleClass")
	})
	return err
}

type TextCellAccessibleClass struct {
	native uintptr
}

var textCellAccessiblePrivateStruct *gi.Struct
var textCellAccessiblePrivateStruct_Once sync.Once

func textCellAccessiblePrivateStruct_Set() error {
	var err error
	textCellAccessiblePrivateStruct_Once.Do(func() {
		textCellAccessiblePrivateStruct, err = gi.StructNew("Gtk", "TextCellAccessiblePrivate")
	})
	return err
}

type TextCellAccessiblePrivate struct {
	native uintptr
}

var textChildAnchorClassStruct *gi.Struct
var textChildAnchorClassStruct_Once sync.Once

func textChildAnchorClassStruct_Set() error {
	var err error
	textChildAnchorClassStruct_Once.Do(func() {
		textChildAnchorClassStruct, err = gi.StructNew("Gtk", "TextChildAnchorClass")
	})
	return err
}

type TextChildAnchorClass struct {
	native uintptr
}

var textIterStruct *gi.Struct
var textIterStruct_Once sync.Once

func textIterStruct_Set() error {
	var err error
	textIterStruct_Once.Do(func() {
		textIterStruct, err = gi.StructNew("Gtk", "TextIter")
	})
	return err
}

type TextIter struct {
	native uintptr
}

var textIterAssignFunction *gi.Function
var textIterAssignFunction_Once sync.Once

func textIterAssignFunction_Set() error {
	var err error
	textIterAssignFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterAssignFunction, err = textIterStruct.InvokerNew("assign")
	})
	return err
}

// Assign is a representation of the C type gtk_text_iter_assign.
func (recv *TextIter) Assign(other *TextIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(other.native)

	err := textIterAssignFunction_Set()
	if err == nil {
		textIterAssignFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterBackwardCharFunction *gi.Function
var textIterBackwardCharFunction_Once sync.Once

func textIterBackwardCharFunction_Set() error {
	var err error
	textIterBackwardCharFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardCharFunction, err = textIterStruct.InvokerNew("backward_char")
	})
	return err
}

// BackwardChar is a representation of the C type gtk_text_iter_backward_char.
func (recv *TextIter) BackwardChar() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterBackwardCharFunction_Set()
	if err == nil {
		ret = textIterBackwardCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardCharsFunction *gi.Function
var textIterBackwardCharsFunction_Once sync.Once

func textIterBackwardCharsFunction_Set() error {
	var err error
	textIterBackwardCharsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardCharsFunction, err = textIterStruct.InvokerNew("backward_chars")
	})
	return err
}

// BackwardChars is a representation of the C type gtk_text_iter_backward_chars.
func (recv *TextIter) BackwardChars(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterBackwardCharsFunction_Set()
	if err == nil {
		ret = textIterBackwardCharsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardCursorPositionFunction *gi.Function
var textIterBackwardCursorPositionFunction_Once sync.Once

func textIterBackwardCursorPositionFunction_Set() error {
	var err error
	textIterBackwardCursorPositionFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardCursorPositionFunction, err = textIterStruct.InvokerNew("backward_cursor_position")
	})
	return err
}

// BackwardCursorPosition is a representation of the C type gtk_text_iter_backward_cursor_position.
func (recv *TextIter) BackwardCursorPosition() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterBackwardCursorPositionFunction_Set()
	if err == nil {
		ret = textIterBackwardCursorPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardCursorPositionsFunction *gi.Function
var textIterBackwardCursorPositionsFunction_Once sync.Once

func textIterBackwardCursorPositionsFunction_Set() error {
	var err error
	textIterBackwardCursorPositionsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardCursorPositionsFunction, err = textIterStruct.InvokerNew("backward_cursor_positions")
	})
	return err
}

// BackwardCursorPositions is a representation of the C type gtk_text_iter_backward_cursor_positions.
func (recv *TextIter) BackwardCursorPositions(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterBackwardCursorPositionsFunction_Set()
	if err == nil {
		ret = textIterBackwardCursorPositionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_backward_find_char' : parameter 'pred' of type 'TextCharPredicate' not supported

var textIterBackwardLineFunction *gi.Function
var textIterBackwardLineFunction_Once sync.Once

func textIterBackwardLineFunction_Set() error {
	var err error
	textIterBackwardLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardLineFunction, err = textIterStruct.InvokerNew("backward_line")
	})
	return err
}

// BackwardLine is a representation of the C type gtk_text_iter_backward_line.
func (recv *TextIter) BackwardLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterBackwardLineFunction_Set()
	if err == nil {
		ret = textIterBackwardLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardLinesFunction *gi.Function
var textIterBackwardLinesFunction_Once sync.Once

func textIterBackwardLinesFunction_Set() error {
	var err error
	textIterBackwardLinesFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardLinesFunction, err = textIterStruct.InvokerNew("backward_lines")
	})
	return err
}

// BackwardLines is a representation of the C type gtk_text_iter_backward_lines.
func (recv *TextIter) BackwardLines(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterBackwardLinesFunction_Set()
	if err == nil {
		ret = textIterBackwardLinesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_backward_search' : parameter 'flags' of type 'TextSearchFlags' not supported

var textIterBackwardSentenceStartFunction *gi.Function
var textIterBackwardSentenceStartFunction_Once sync.Once

func textIterBackwardSentenceStartFunction_Set() error {
	var err error
	textIterBackwardSentenceStartFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardSentenceStartFunction, err = textIterStruct.InvokerNew("backward_sentence_start")
	})
	return err
}

// BackwardSentenceStart is a representation of the C type gtk_text_iter_backward_sentence_start.
func (recv *TextIter) BackwardSentenceStart() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterBackwardSentenceStartFunction_Set()
	if err == nil {
		ret = textIterBackwardSentenceStartFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardSentenceStartsFunction *gi.Function
var textIterBackwardSentenceStartsFunction_Once sync.Once

func textIterBackwardSentenceStartsFunction_Set() error {
	var err error
	textIterBackwardSentenceStartsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardSentenceStartsFunction, err = textIterStruct.InvokerNew("backward_sentence_starts")
	})
	return err
}

// BackwardSentenceStarts is a representation of the C type gtk_text_iter_backward_sentence_starts.
func (recv *TextIter) BackwardSentenceStarts(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterBackwardSentenceStartsFunction_Set()
	if err == nil {
		ret = textIterBackwardSentenceStartsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_backward_to_tag_toggle' : parameter 'tag' of type 'TextTag' not supported

var textIterBackwardVisibleCursorPositionFunction *gi.Function
var textIterBackwardVisibleCursorPositionFunction_Once sync.Once

func textIterBackwardVisibleCursorPositionFunction_Set() error {
	var err error
	textIterBackwardVisibleCursorPositionFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardVisibleCursorPositionFunction, err = textIterStruct.InvokerNew("backward_visible_cursor_position")
	})
	return err
}

// BackwardVisibleCursorPosition is a representation of the C type gtk_text_iter_backward_visible_cursor_position.
func (recv *TextIter) BackwardVisibleCursorPosition() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterBackwardVisibleCursorPositionFunction_Set()
	if err == nil {
		ret = textIterBackwardVisibleCursorPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardVisibleCursorPositionsFunction *gi.Function
var textIterBackwardVisibleCursorPositionsFunction_Once sync.Once

func textIterBackwardVisibleCursorPositionsFunction_Set() error {
	var err error
	textIterBackwardVisibleCursorPositionsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardVisibleCursorPositionsFunction, err = textIterStruct.InvokerNew("backward_visible_cursor_positions")
	})
	return err
}

// BackwardVisibleCursorPositions is a representation of the C type gtk_text_iter_backward_visible_cursor_positions.
func (recv *TextIter) BackwardVisibleCursorPositions(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterBackwardVisibleCursorPositionsFunction_Set()
	if err == nil {
		ret = textIterBackwardVisibleCursorPositionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardVisibleLineFunction *gi.Function
var textIterBackwardVisibleLineFunction_Once sync.Once

func textIterBackwardVisibleLineFunction_Set() error {
	var err error
	textIterBackwardVisibleLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardVisibleLineFunction, err = textIterStruct.InvokerNew("backward_visible_line")
	})
	return err
}

// BackwardVisibleLine is a representation of the C type gtk_text_iter_backward_visible_line.
func (recv *TextIter) BackwardVisibleLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterBackwardVisibleLineFunction_Set()
	if err == nil {
		ret = textIterBackwardVisibleLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardVisibleLinesFunction *gi.Function
var textIterBackwardVisibleLinesFunction_Once sync.Once

func textIterBackwardVisibleLinesFunction_Set() error {
	var err error
	textIterBackwardVisibleLinesFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardVisibleLinesFunction, err = textIterStruct.InvokerNew("backward_visible_lines")
	})
	return err
}

// BackwardVisibleLines is a representation of the C type gtk_text_iter_backward_visible_lines.
func (recv *TextIter) BackwardVisibleLines(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterBackwardVisibleLinesFunction_Set()
	if err == nil {
		ret = textIterBackwardVisibleLinesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardVisibleWordStartFunction *gi.Function
var textIterBackwardVisibleWordStartFunction_Once sync.Once

func textIterBackwardVisibleWordStartFunction_Set() error {
	var err error
	textIterBackwardVisibleWordStartFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardVisibleWordStartFunction, err = textIterStruct.InvokerNew("backward_visible_word_start")
	})
	return err
}

// BackwardVisibleWordStart is a representation of the C type gtk_text_iter_backward_visible_word_start.
func (recv *TextIter) BackwardVisibleWordStart() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterBackwardVisibleWordStartFunction_Set()
	if err == nil {
		ret = textIterBackwardVisibleWordStartFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardVisibleWordStartsFunction *gi.Function
var textIterBackwardVisibleWordStartsFunction_Once sync.Once

func textIterBackwardVisibleWordStartsFunction_Set() error {
	var err error
	textIterBackwardVisibleWordStartsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardVisibleWordStartsFunction, err = textIterStruct.InvokerNew("backward_visible_word_starts")
	})
	return err
}

// BackwardVisibleWordStarts is a representation of the C type gtk_text_iter_backward_visible_word_starts.
func (recv *TextIter) BackwardVisibleWordStarts(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterBackwardVisibleWordStartsFunction_Set()
	if err == nil {
		ret = textIterBackwardVisibleWordStartsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardWordStartFunction *gi.Function
var textIterBackwardWordStartFunction_Once sync.Once

func textIterBackwardWordStartFunction_Set() error {
	var err error
	textIterBackwardWordStartFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardWordStartFunction, err = textIterStruct.InvokerNew("backward_word_start")
	})
	return err
}

// BackwardWordStart is a representation of the C type gtk_text_iter_backward_word_start.
func (recv *TextIter) BackwardWordStart() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterBackwardWordStartFunction_Set()
	if err == nil {
		ret = textIterBackwardWordStartFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterBackwardWordStartsFunction *gi.Function
var textIterBackwardWordStartsFunction_Once sync.Once

func textIterBackwardWordStartsFunction_Set() error {
	var err error
	textIterBackwardWordStartsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardWordStartsFunction, err = textIterStruct.InvokerNew("backward_word_starts")
	})
	return err
}

// BackwardWordStarts is a representation of the C type gtk_text_iter_backward_word_starts.
func (recv *TextIter) BackwardWordStarts(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterBackwardWordStartsFunction_Set()
	if err == nil {
		ret = textIterBackwardWordStartsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_begins_tag' : parameter 'tag' of type 'TextTag' not supported

var textIterCanInsertFunction *gi.Function
var textIterCanInsertFunction_Once sync.Once

func textIterCanInsertFunction_Set() error {
	var err error
	textIterCanInsertFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterCanInsertFunction, err = textIterStruct.InvokerNew("can_insert")
	})
	return err
}

// CanInsert is a representation of the C type gtk_text_iter_can_insert.
func (recv *TextIter) CanInsert(defaultEditability bool) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(defaultEditability)

	var ret gi.Argument

	err := textIterCanInsertFunction_Set()
	if err == nil {
		ret = textIterCanInsertFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterCompareFunction *gi.Function
var textIterCompareFunction_Once sync.Once

func textIterCompareFunction_Set() error {
	var err error
	textIterCompareFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterCompareFunction, err = textIterStruct.InvokerNew("compare")
	})
	return err
}

// Compare is a representation of the C type gtk_text_iter_compare.
func (recv *TextIter) Compare(rhs *TextIter) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(rhs.native)

	var ret gi.Argument

	err := textIterCompareFunction_Set()
	if err == nil {
		ret = textIterCompareFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var textIterCopyFunction *gi.Function
var textIterCopyFunction_Once sync.Once

func textIterCopyFunction_Set() error {
	var err error
	textIterCopyFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterCopyFunction, err = textIterStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_text_iter_copy.
func (recv *TextIter) Copy() *TextIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterCopyFunction_Set()
	if err == nil {
		ret = textIterCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TextIter{native: ret.Pointer()}

	return retGo
}

var textIterEditableFunction *gi.Function
var textIterEditableFunction_Once sync.Once

func textIterEditableFunction_Set() error {
	var err error
	textIterEditableFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterEditableFunction, err = textIterStruct.InvokerNew("editable")
	})
	return err
}

// Editable is a representation of the C type gtk_text_iter_editable.
func (recv *TextIter) Editable(defaultSetting bool) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetBoolean(defaultSetting)

	var ret gi.Argument

	err := textIterEditableFunction_Set()
	if err == nil {
		ret = textIterEditableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterEndsLineFunction *gi.Function
var textIterEndsLineFunction_Once sync.Once

func textIterEndsLineFunction_Set() error {
	var err error
	textIterEndsLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterEndsLineFunction, err = textIterStruct.InvokerNew("ends_line")
	})
	return err
}

// EndsLine is a representation of the C type gtk_text_iter_ends_line.
func (recv *TextIter) EndsLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterEndsLineFunction_Set()
	if err == nil {
		ret = textIterEndsLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterEndsSentenceFunction *gi.Function
var textIterEndsSentenceFunction_Once sync.Once

func textIterEndsSentenceFunction_Set() error {
	var err error
	textIterEndsSentenceFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterEndsSentenceFunction, err = textIterStruct.InvokerNew("ends_sentence")
	})
	return err
}

// EndsSentence is a representation of the C type gtk_text_iter_ends_sentence.
func (recv *TextIter) EndsSentence() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterEndsSentenceFunction_Set()
	if err == nil {
		ret = textIterEndsSentenceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_ends_tag' : parameter 'tag' of type 'TextTag' not supported

var textIterEndsWordFunction *gi.Function
var textIterEndsWordFunction_Once sync.Once

func textIterEndsWordFunction_Set() error {
	var err error
	textIterEndsWordFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterEndsWordFunction, err = textIterStruct.InvokerNew("ends_word")
	})
	return err
}

// EndsWord is a representation of the C type gtk_text_iter_ends_word.
func (recv *TextIter) EndsWord() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterEndsWordFunction_Set()
	if err == nil {
		ret = textIterEndsWordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterEqualFunction *gi.Function
var textIterEqualFunction_Once sync.Once

func textIterEqualFunction_Set() error {
	var err error
	textIterEqualFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterEqualFunction, err = textIterStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type gtk_text_iter_equal.
func (recv *TextIter) Equal(rhs *TextIter) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(rhs.native)

	var ret gi.Argument

	err := textIterEqualFunction_Set()
	if err == nil {
		ret = textIterEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardCharFunction *gi.Function
var textIterForwardCharFunction_Once sync.Once

func textIterForwardCharFunction_Set() error {
	var err error
	textIterForwardCharFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardCharFunction, err = textIterStruct.InvokerNew("forward_char")
	})
	return err
}

// ForwardChar is a representation of the C type gtk_text_iter_forward_char.
func (recv *TextIter) ForwardChar() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterForwardCharFunction_Set()
	if err == nil {
		ret = textIterForwardCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardCharsFunction *gi.Function
var textIterForwardCharsFunction_Once sync.Once

func textIterForwardCharsFunction_Set() error {
	var err error
	textIterForwardCharsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardCharsFunction, err = textIterStruct.InvokerNew("forward_chars")
	})
	return err
}

// ForwardChars is a representation of the C type gtk_text_iter_forward_chars.
func (recv *TextIter) ForwardChars(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterForwardCharsFunction_Set()
	if err == nil {
		ret = textIterForwardCharsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardCursorPositionFunction *gi.Function
var textIterForwardCursorPositionFunction_Once sync.Once

func textIterForwardCursorPositionFunction_Set() error {
	var err error
	textIterForwardCursorPositionFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardCursorPositionFunction, err = textIterStruct.InvokerNew("forward_cursor_position")
	})
	return err
}

// ForwardCursorPosition is a representation of the C type gtk_text_iter_forward_cursor_position.
func (recv *TextIter) ForwardCursorPosition() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterForwardCursorPositionFunction_Set()
	if err == nil {
		ret = textIterForwardCursorPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardCursorPositionsFunction *gi.Function
var textIterForwardCursorPositionsFunction_Once sync.Once

func textIterForwardCursorPositionsFunction_Set() error {
	var err error
	textIterForwardCursorPositionsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardCursorPositionsFunction, err = textIterStruct.InvokerNew("forward_cursor_positions")
	})
	return err
}

// ForwardCursorPositions is a representation of the C type gtk_text_iter_forward_cursor_positions.
func (recv *TextIter) ForwardCursorPositions(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterForwardCursorPositionsFunction_Set()
	if err == nil {
		ret = textIterForwardCursorPositionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_forward_find_char' : parameter 'pred' of type 'TextCharPredicate' not supported

var textIterForwardLineFunction *gi.Function
var textIterForwardLineFunction_Once sync.Once

func textIterForwardLineFunction_Set() error {
	var err error
	textIterForwardLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardLineFunction, err = textIterStruct.InvokerNew("forward_line")
	})
	return err
}

// ForwardLine is a representation of the C type gtk_text_iter_forward_line.
func (recv *TextIter) ForwardLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterForwardLineFunction_Set()
	if err == nil {
		ret = textIterForwardLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardLinesFunction *gi.Function
var textIterForwardLinesFunction_Once sync.Once

func textIterForwardLinesFunction_Set() error {
	var err error
	textIterForwardLinesFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardLinesFunction, err = textIterStruct.InvokerNew("forward_lines")
	})
	return err
}

// ForwardLines is a representation of the C type gtk_text_iter_forward_lines.
func (recv *TextIter) ForwardLines(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterForwardLinesFunction_Set()
	if err == nil {
		ret = textIterForwardLinesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_forward_search' : parameter 'flags' of type 'TextSearchFlags' not supported

var textIterForwardSentenceEndFunction *gi.Function
var textIterForwardSentenceEndFunction_Once sync.Once

func textIterForwardSentenceEndFunction_Set() error {
	var err error
	textIterForwardSentenceEndFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardSentenceEndFunction, err = textIterStruct.InvokerNew("forward_sentence_end")
	})
	return err
}

// ForwardSentenceEnd is a representation of the C type gtk_text_iter_forward_sentence_end.
func (recv *TextIter) ForwardSentenceEnd() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterForwardSentenceEndFunction_Set()
	if err == nil {
		ret = textIterForwardSentenceEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardSentenceEndsFunction *gi.Function
var textIterForwardSentenceEndsFunction_Once sync.Once

func textIterForwardSentenceEndsFunction_Set() error {
	var err error
	textIterForwardSentenceEndsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardSentenceEndsFunction, err = textIterStruct.InvokerNew("forward_sentence_ends")
	})
	return err
}

// ForwardSentenceEnds is a representation of the C type gtk_text_iter_forward_sentence_ends.
func (recv *TextIter) ForwardSentenceEnds(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterForwardSentenceEndsFunction_Set()
	if err == nil {
		ret = textIterForwardSentenceEndsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardToEndFunction *gi.Function
var textIterForwardToEndFunction_Once sync.Once

func textIterForwardToEndFunction_Set() error {
	var err error
	textIterForwardToEndFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardToEndFunction, err = textIterStruct.InvokerNew("forward_to_end")
	})
	return err
}

// ForwardToEnd is a representation of the C type gtk_text_iter_forward_to_end.
func (recv *TextIter) ForwardToEnd() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := textIterForwardToEndFunction_Set()
	if err == nil {
		textIterForwardToEndFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterForwardToLineEndFunction *gi.Function
var textIterForwardToLineEndFunction_Once sync.Once

func textIterForwardToLineEndFunction_Set() error {
	var err error
	textIterForwardToLineEndFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardToLineEndFunction, err = textIterStruct.InvokerNew("forward_to_line_end")
	})
	return err
}

// ForwardToLineEnd is a representation of the C type gtk_text_iter_forward_to_line_end.
func (recv *TextIter) ForwardToLineEnd() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterForwardToLineEndFunction_Set()
	if err == nil {
		ret = textIterForwardToLineEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_forward_to_tag_toggle' : parameter 'tag' of type 'TextTag' not supported

var textIterForwardVisibleCursorPositionFunction *gi.Function
var textIterForwardVisibleCursorPositionFunction_Once sync.Once

func textIterForwardVisibleCursorPositionFunction_Set() error {
	var err error
	textIterForwardVisibleCursorPositionFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardVisibleCursorPositionFunction, err = textIterStruct.InvokerNew("forward_visible_cursor_position")
	})
	return err
}

// ForwardVisibleCursorPosition is a representation of the C type gtk_text_iter_forward_visible_cursor_position.
func (recv *TextIter) ForwardVisibleCursorPosition() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterForwardVisibleCursorPositionFunction_Set()
	if err == nil {
		ret = textIterForwardVisibleCursorPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardVisibleCursorPositionsFunction *gi.Function
var textIterForwardVisibleCursorPositionsFunction_Once sync.Once

func textIterForwardVisibleCursorPositionsFunction_Set() error {
	var err error
	textIterForwardVisibleCursorPositionsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardVisibleCursorPositionsFunction, err = textIterStruct.InvokerNew("forward_visible_cursor_positions")
	})
	return err
}

// ForwardVisibleCursorPositions is a representation of the C type gtk_text_iter_forward_visible_cursor_positions.
func (recv *TextIter) ForwardVisibleCursorPositions(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterForwardVisibleCursorPositionsFunction_Set()
	if err == nil {
		ret = textIterForwardVisibleCursorPositionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardVisibleLineFunction *gi.Function
var textIterForwardVisibleLineFunction_Once sync.Once

func textIterForwardVisibleLineFunction_Set() error {
	var err error
	textIterForwardVisibleLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardVisibleLineFunction, err = textIterStruct.InvokerNew("forward_visible_line")
	})
	return err
}

// ForwardVisibleLine is a representation of the C type gtk_text_iter_forward_visible_line.
func (recv *TextIter) ForwardVisibleLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterForwardVisibleLineFunction_Set()
	if err == nil {
		ret = textIterForwardVisibleLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardVisibleLinesFunction *gi.Function
var textIterForwardVisibleLinesFunction_Once sync.Once

func textIterForwardVisibleLinesFunction_Set() error {
	var err error
	textIterForwardVisibleLinesFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardVisibleLinesFunction, err = textIterStruct.InvokerNew("forward_visible_lines")
	})
	return err
}

// ForwardVisibleLines is a representation of the C type gtk_text_iter_forward_visible_lines.
func (recv *TextIter) ForwardVisibleLines(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterForwardVisibleLinesFunction_Set()
	if err == nil {
		ret = textIterForwardVisibleLinesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardVisibleWordEndFunction *gi.Function
var textIterForwardVisibleWordEndFunction_Once sync.Once

func textIterForwardVisibleWordEndFunction_Set() error {
	var err error
	textIterForwardVisibleWordEndFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardVisibleWordEndFunction, err = textIterStruct.InvokerNew("forward_visible_word_end")
	})
	return err
}

// ForwardVisibleWordEnd is a representation of the C type gtk_text_iter_forward_visible_word_end.
func (recv *TextIter) ForwardVisibleWordEnd() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterForwardVisibleWordEndFunction_Set()
	if err == nil {
		ret = textIterForwardVisibleWordEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardVisibleWordEndsFunction *gi.Function
var textIterForwardVisibleWordEndsFunction_Once sync.Once

func textIterForwardVisibleWordEndsFunction_Set() error {
	var err error
	textIterForwardVisibleWordEndsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardVisibleWordEndsFunction, err = textIterStruct.InvokerNew("forward_visible_word_ends")
	})
	return err
}

// ForwardVisibleWordEnds is a representation of the C type gtk_text_iter_forward_visible_word_ends.
func (recv *TextIter) ForwardVisibleWordEnds(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterForwardVisibleWordEndsFunction_Set()
	if err == nil {
		ret = textIterForwardVisibleWordEndsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardWordEndFunction *gi.Function
var textIterForwardWordEndFunction_Once sync.Once

func textIterForwardWordEndFunction_Set() error {
	var err error
	textIterForwardWordEndFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardWordEndFunction, err = textIterStruct.InvokerNew("forward_word_end")
	})
	return err
}

// ForwardWordEnd is a representation of the C type gtk_text_iter_forward_word_end.
func (recv *TextIter) ForwardWordEnd() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterForwardWordEndFunction_Set()
	if err == nil {
		ret = textIterForwardWordEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterForwardWordEndsFunction *gi.Function
var textIterForwardWordEndsFunction_Once sync.Once

func textIterForwardWordEndsFunction_Set() error {
	var err error
	textIterForwardWordEndsFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardWordEndsFunction, err = textIterStruct.InvokerNew("forward_word_ends")
	})
	return err
}

// ForwardWordEnds is a representation of the C type gtk_text_iter_forward_word_ends.
func (recv *TextIter) ForwardWordEnds(count int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(count)

	var ret gi.Argument

	err := textIterForwardWordEndsFunction_Set()
	if err == nil {
		ret = textIterForwardWordEndsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterFreeFunction *gi.Function
var textIterFreeFunction_Once sync.Once

func textIterFreeFunction_Set() error {
	var err error
	textIterFreeFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterFreeFunction, err = textIterStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_text_iter_free.
func (recv *TextIter) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := textIterFreeFunction_Set()
	if err == nil {
		textIterFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterGetAttributesFunction *gi.Function
var textIterGetAttributesFunction_Once sync.Once

func textIterGetAttributesFunction_Set() error {
	var err error
	textIterGetAttributesFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetAttributesFunction, err = textIterStruct.InvokerNew("get_attributes")
	})
	return err
}

// GetAttributes is a representation of the C type gtk_text_iter_get_attributes.
func (recv *TextIter) GetAttributes() (bool, *TextAttributes) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := textIterGetAttributesFunction_Set()
	if err == nil {
		ret = textIterGetAttributesFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := &TextAttributes{native: outArgs[0].Pointer()}

	return retGo, out0
}

// UNSUPPORTED : C value 'gtk_text_iter_get_buffer' : return type 'TextBuffer' not supported

var textIterGetBytesInLineFunction *gi.Function
var textIterGetBytesInLineFunction_Once sync.Once

func textIterGetBytesInLineFunction_Set() error {
	var err error
	textIterGetBytesInLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetBytesInLineFunction, err = textIterStruct.InvokerNew("get_bytes_in_line")
	})
	return err
}

// GetBytesInLine is a representation of the C type gtk_text_iter_get_bytes_in_line.
func (recv *TextIter) GetBytesInLine() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetBytesInLineFunction_Set()
	if err == nil {
		ret = textIterGetBytesInLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_char' : return type 'gunichar' not supported

var textIterGetCharsInLineFunction *gi.Function
var textIterGetCharsInLineFunction_Once sync.Once

func textIterGetCharsInLineFunction_Set() error {
	var err error
	textIterGetCharsInLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetCharsInLineFunction, err = textIterStruct.InvokerNew("get_chars_in_line")
	})
	return err
}

// GetCharsInLine is a representation of the C type gtk_text_iter_get_chars_in_line.
func (recv *TextIter) GetCharsInLine() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetCharsInLineFunction_Set()
	if err == nil {
		ret = textIterGetCharsInLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_child_anchor' : return type 'TextChildAnchor' not supported

// UNSUPPORTED : C value 'gtk_text_iter_get_language' : return type 'Pango.Language' not supported

var textIterGetLineFunction *gi.Function
var textIterGetLineFunction_Once sync.Once

func textIterGetLineFunction_Set() error {
	var err error
	textIterGetLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetLineFunction, err = textIterStruct.InvokerNew("get_line")
	})
	return err
}

// GetLine is a representation of the C type gtk_text_iter_get_line.
func (recv *TextIter) GetLine() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetLineFunction_Set()
	if err == nil {
		ret = textIterGetLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var textIterGetLineIndexFunction *gi.Function
var textIterGetLineIndexFunction_Once sync.Once

func textIterGetLineIndexFunction_Set() error {
	var err error
	textIterGetLineIndexFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetLineIndexFunction, err = textIterStruct.InvokerNew("get_line_index")
	})
	return err
}

// GetLineIndex is a representation of the C type gtk_text_iter_get_line_index.
func (recv *TextIter) GetLineIndex() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetLineIndexFunction_Set()
	if err == nil {
		ret = textIterGetLineIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var textIterGetLineOffsetFunction *gi.Function
var textIterGetLineOffsetFunction_Once sync.Once

func textIterGetLineOffsetFunction_Set() error {
	var err error
	textIterGetLineOffsetFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetLineOffsetFunction, err = textIterStruct.InvokerNew("get_line_offset")
	})
	return err
}

// GetLineOffset is a representation of the C type gtk_text_iter_get_line_offset.
func (recv *TextIter) GetLineOffset() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetLineOffsetFunction_Set()
	if err == nil {
		ret = textIterGetLineOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_marks' : return type 'GLib.SList' not supported

var textIterGetOffsetFunction *gi.Function
var textIterGetOffsetFunction_Once sync.Once

func textIterGetOffsetFunction_Set() error {
	var err error
	textIterGetOffsetFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetOffsetFunction, err = textIterStruct.InvokerNew("get_offset")
	})
	return err
}

// GetOffset is a representation of the C type gtk_text_iter_get_offset.
func (recv *TextIter) GetOffset() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetOffsetFunction_Set()
	if err == nil {
		ret = textIterGetOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_pixbuf' : return type 'GdkPixbuf.Pixbuf' not supported

var textIterGetSliceFunction *gi.Function
var textIterGetSliceFunction_Once sync.Once

func textIterGetSliceFunction_Set() error {
	var err error
	textIterGetSliceFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetSliceFunction, err = textIterStruct.InvokerNew("get_slice")
	})
	return err
}

// GetSlice is a representation of the C type gtk_text_iter_get_slice.
func (recv *TextIter) GetSlice(end *TextIter) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(end.native)

	var ret gi.Argument

	err := textIterGetSliceFunction_Set()
	if err == nil {
		ret = textIterGetSliceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_tags' : return type 'GLib.SList' not supported

var textIterGetTextFunction *gi.Function
var textIterGetTextFunction_Once sync.Once

func textIterGetTextFunction_Set() error {
	var err error
	textIterGetTextFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetTextFunction, err = textIterStruct.InvokerNew("get_text")
	})
	return err
}

// GetText is a representation of the C type gtk_text_iter_get_text.
func (recv *TextIter) GetText(end *TextIter) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(end.native)

	var ret gi.Argument

	err := textIterGetTextFunction_Set()
	if err == nil {
		ret = textIterGetTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_get_toggled_tags' : return type 'GLib.SList' not supported

var textIterGetVisibleLineIndexFunction *gi.Function
var textIterGetVisibleLineIndexFunction_Once sync.Once

func textIterGetVisibleLineIndexFunction_Set() error {
	var err error
	textIterGetVisibleLineIndexFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetVisibleLineIndexFunction, err = textIterStruct.InvokerNew("get_visible_line_index")
	})
	return err
}

// GetVisibleLineIndex is a representation of the C type gtk_text_iter_get_visible_line_index.
func (recv *TextIter) GetVisibleLineIndex() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetVisibleLineIndexFunction_Set()
	if err == nil {
		ret = textIterGetVisibleLineIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var textIterGetVisibleLineOffsetFunction *gi.Function
var textIterGetVisibleLineOffsetFunction_Once sync.Once

func textIterGetVisibleLineOffsetFunction_Set() error {
	var err error
	textIterGetVisibleLineOffsetFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetVisibleLineOffsetFunction, err = textIterStruct.InvokerNew("get_visible_line_offset")
	})
	return err
}

// GetVisibleLineOffset is a representation of the C type gtk_text_iter_get_visible_line_offset.
func (recv *TextIter) GetVisibleLineOffset() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetVisibleLineOffsetFunction_Set()
	if err == nil {
		ret = textIterGetVisibleLineOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var textIterGetVisibleSliceFunction *gi.Function
var textIterGetVisibleSliceFunction_Once sync.Once

func textIterGetVisibleSliceFunction_Set() error {
	var err error
	textIterGetVisibleSliceFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetVisibleSliceFunction, err = textIterStruct.InvokerNew("get_visible_slice")
	})
	return err
}

// GetVisibleSlice is a representation of the C type gtk_text_iter_get_visible_slice.
func (recv *TextIter) GetVisibleSlice(end *TextIter) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(end.native)

	var ret gi.Argument

	err := textIterGetVisibleSliceFunction_Set()
	if err == nil {
		ret = textIterGetVisibleSliceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var textIterGetVisibleTextFunction *gi.Function
var textIterGetVisibleTextFunction_Once sync.Once

func textIterGetVisibleTextFunction_Set() error {
	var err error
	textIterGetVisibleTextFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetVisibleTextFunction, err = textIterStruct.InvokerNew("get_visible_text")
	})
	return err
}

// GetVisibleText is a representation of the C type gtk_text_iter_get_visible_text.
func (recv *TextIter) GetVisibleText(end *TextIter) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(end.native)

	var ret gi.Argument

	err := textIterGetVisibleTextFunction_Set()
	if err == nil {
		ret = textIterGetVisibleTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_has_tag' : parameter 'tag' of type 'TextTag' not supported

var textIterInRangeFunction *gi.Function
var textIterInRangeFunction_Once sync.Once

func textIterInRangeFunction_Set() error {
	var err error
	textIterInRangeFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterInRangeFunction, err = textIterStruct.InvokerNew("in_range")
	})
	return err
}

// InRange is a representation of the C type gtk_text_iter_in_range.
func (recv *TextIter) InRange(start *TextIter, end *TextIter) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(start.native)
	inArgs[2].SetPointer(end.native)

	var ret gi.Argument

	err := textIterInRangeFunction_Set()
	if err == nil {
		ret = textIterInRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterInsideSentenceFunction *gi.Function
var textIterInsideSentenceFunction_Once sync.Once

func textIterInsideSentenceFunction_Set() error {
	var err error
	textIterInsideSentenceFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterInsideSentenceFunction, err = textIterStruct.InvokerNew("inside_sentence")
	})
	return err
}

// InsideSentence is a representation of the C type gtk_text_iter_inside_sentence.
func (recv *TextIter) InsideSentence() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterInsideSentenceFunction_Set()
	if err == nil {
		ret = textIterInsideSentenceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterInsideWordFunction *gi.Function
var textIterInsideWordFunction_Once sync.Once

func textIterInsideWordFunction_Set() error {
	var err error
	textIterInsideWordFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterInsideWordFunction, err = textIterStruct.InvokerNew("inside_word")
	})
	return err
}

// InsideWord is a representation of the C type gtk_text_iter_inside_word.
func (recv *TextIter) InsideWord() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterInsideWordFunction_Set()
	if err == nil {
		ret = textIterInsideWordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterIsCursorPositionFunction *gi.Function
var textIterIsCursorPositionFunction_Once sync.Once

func textIterIsCursorPositionFunction_Set() error {
	var err error
	textIterIsCursorPositionFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterIsCursorPositionFunction, err = textIterStruct.InvokerNew("is_cursor_position")
	})
	return err
}

// IsCursorPosition is a representation of the C type gtk_text_iter_is_cursor_position.
func (recv *TextIter) IsCursorPosition() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterIsCursorPositionFunction_Set()
	if err == nil {
		ret = textIterIsCursorPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterIsEndFunction *gi.Function
var textIterIsEndFunction_Once sync.Once

func textIterIsEndFunction_Set() error {
	var err error
	textIterIsEndFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterIsEndFunction, err = textIterStruct.InvokerNew("is_end")
	})
	return err
}

// IsEnd is a representation of the C type gtk_text_iter_is_end.
func (recv *TextIter) IsEnd() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterIsEndFunction_Set()
	if err == nil {
		ret = textIterIsEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterIsStartFunction *gi.Function
var textIterIsStartFunction_Once sync.Once

func textIterIsStartFunction_Set() error {
	var err error
	textIterIsStartFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterIsStartFunction, err = textIterStruct.InvokerNew("is_start")
	})
	return err
}

// IsStart is a representation of the C type gtk_text_iter_is_start.
func (recv *TextIter) IsStart() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterIsStartFunction_Set()
	if err == nil {
		ret = textIterIsStartFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterOrderFunction *gi.Function
var textIterOrderFunction_Once sync.Once

func textIterOrderFunction_Set() error {
	var err error
	textIterOrderFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterOrderFunction, err = textIterStruct.InvokerNew("order")
	})
	return err
}

// Order is a representation of the C type gtk_text_iter_order.
func (recv *TextIter) Order(second *TextIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(second.native)

	err := textIterOrderFunction_Set()
	if err == nil {
		textIterOrderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterSetLineFunction *gi.Function
var textIterSetLineFunction_Once sync.Once

func textIterSetLineFunction_Set() error {
	var err error
	textIterSetLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterSetLineFunction, err = textIterStruct.InvokerNew("set_line")
	})
	return err
}

// SetLine is a representation of the C type gtk_text_iter_set_line.
func (recv *TextIter) SetLine(lineNumber int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(lineNumber)

	err := textIterSetLineFunction_Set()
	if err == nil {
		textIterSetLineFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterSetLineIndexFunction *gi.Function
var textIterSetLineIndexFunction_Once sync.Once

func textIterSetLineIndexFunction_Set() error {
	var err error
	textIterSetLineIndexFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterSetLineIndexFunction, err = textIterStruct.InvokerNew("set_line_index")
	})
	return err
}

// SetLineIndex is a representation of the C type gtk_text_iter_set_line_index.
func (recv *TextIter) SetLineIndex(byteOnLine int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(byteOnLine)

	err := textIterSetLineIndexFunction_Set()
	if err == nil {
		textIterSetLineIndexFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterSetLineOffsetFunction *gi.Function
var textIterSetLineOffsetFunction_Once sync.Once

func textIterSetLineOffsetFunction_Set() error {
	var err error
	textIterSetLineOffsetFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterSetLineOffsetFunction, err = textIterStruct.InvokerNew("set_line_offset")
	})
	return err
}

// SetLineOffset is a representation of the C type gtk_text_iter_set_line_offset.
func (recv *TextIter) SetLineOffset(charOnLine int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(charOnLine)

	err := textIterSetLineOffsetFunction_Set()
	if err == nil {
		textIterSetLineOffsetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterSetOffsetFunction *gi.Function
var textIterSetOffsetFunction_Once sync.Once

func textIterSetOffsetFunction_Set() error {
	var err error
	textIterSetOffsetFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterSetOffsetFunction, err = textIterStruct.InvokerNew("set_offset")
	})
	return err
}

// SetOffset is a representation of the C type gtk_text_iter_set_offset.
func (recv *TextIter) SetOffset(charOffset int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(charOffset)

	err := textIterSetOffsetFunction_Set()
	if err == nil {
		textIterSetOffsetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterSetVisibleLineIndexFunction *gi.Function
var textIterSetVisibleLineIndexFunction_Once sync.Once

func textIterSetVisibleLineIndexFunction_Set() error {
	var err error
	textIterSetVisibleLineIndexFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterSetVisibleLineIndexFunction, err = textIterStruct.InvokerNew("set_visible_line_index")
	})
	return err
}

// SetVisibleLineIndex is a representation of the C type gtk_text_iter_set_visible_line_index.
func (recv *TextIter) SetVisibleLineIndex(byteOnLine int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(byteOnLine)

	err := textIterSetVisibleLineIndexFunction_Set()
	if err == nil {
		textIterSetVisibleLineIndexFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterSetVisibleLineOffsetFunction *gi.Function
var textIterSetVisibleLineOffsetFunction_Once sync.Once

func textIterSetVisibleLineOffsetFunction_Set() error {
	var err error
	textIterSetVisibleLineOffsetFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterSetVisibleLineOffsetFunction, err = textIterStruct.InvokerNew("set_visible_line_offset")
	})
	return err
}

// SetVisibleLineOffset is a representation of the C type gtk_text_iter_set_visible_line_offset.
func (recv *TextIter) SetVisibleLineOffset(charOnLine int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(charOnLine)

	err := textIterSetVisibleLineOffsetFunction_Set()
	if err == nil {
		textIterSetVisibleLineOffsetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textIterStartsLineFunction *gi.Function
var textIterStartsLineFunction_Once sync.Once

func textIterStartsLineFunction_Set() error {
	var err error
	textIterStartsLineFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterStartsLineFunction, err = textIterStruct.InvokerNew("starts_line")
	})
	return err
}

// StartsLine is a representation of the C type gtk_text_iter_starts_line.
func (recv *TextIter) StartsLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterStartsLineFunction_Set()
	if err == nil {
		ret = textIterStartsLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textIterStartsSentenceFunction *gi.Function
var textIterStartsSentenceFunction_Once sync.Once

func textIterStartsSentenceFunction_Set() error {
	var err error
	textIterStartsSentenceFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterStartsSentenceFunction, err = textIterStruct.InvokerNew("starts_sentence")
	})
	return err
}

// StartsSentence is a representation of the C type gtk_text_iter_starts_sentence.
func (recv *TextIter) StartsSentence() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterStartsSentenceFunction_Set()
	if err == nil {
		ret = textIterStartsSentenceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_starts_tag' : parameter 'tag' of type 'TextTag' not supported

var textIterStartsWordFunction *gi.Function
var textIterStartsWordFunction_Once sync.Once

func textIterStartsWordFunction_Set() error {
	var err error
	textIterStartsWordFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterStartsWordFunction, err = textIterStruct.InvokerNew("starts_word")
	})
	return err
}

// StartsWord is a representation of the C type gtk_text_iter_starts_word.
func (recv *TextIter) StartsWord() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterStartsWordFunction_Set()
	if err == nil {
		ret = textIterStartsWordFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_text_iter_toggles_tag' : parameter 'tag' of type 'TextTag' not supported

var textMarkClassStruct *gi.Struct
var textMarkClassStruct_Once sync.Once

func textMarkClassStruct_Set() error {
	var err error
	textMarkClassStruct_Once.Do(func() {
		textMarkClassStruct, err = gi.StructNew("Gtk", "TextMarkClass")
	})
	return err
}

type TextMarkClass struct {
	native uintptr
}

var textTagClassStruct *gi.Struct
var textTagClassStruct_Once sync.Once

func textTagClassStruct_Set() error {
	var err error
	textTagClassStruct_Once.Do(func() {
		textTagClassStruct, err = gi.StructNew("Gtk", "TextTagClass")
	})
	return err
}

type TextTagClass struct {
	native uintptr
}

var textTagPrivateStruct *gi.Struct
var textTagPrivateStruct_Once sync.Once

func textTagPrivateStruct_Set() error {
	var err error
	textTagPrivateStruct_Once.Do(func() {
		textTagPrivateStruct, err = gi.StructNew("Gtk", "TextTagPrivate")
	})
	return err
}

type TextTagPrivate struct {
	native uintptr
}

var textTagTableClassStruct *gi.Struct
var textTagTableClassStruct_Once sync.Once

func textTagTableClassStruct_Set() error {
	var err error
	textTagTableClassStruct_Once.Do(func() {
		textTagTableClassStruct, err = gi.StructNew("Gtk", "TextTagTableClass")
	})
	return err
}

type TextTagTableClass struct {
	native uintptr
}

var textTagTablePrivateStruct *gi.Struct
var textTagTablePrivateStruct_Once sync.Once

func textTagTablePrivateStruct_Set() error {
	var err error
	textTagTablePrivateStruct_Once.Do(func() {
		textTagTablePrivateStruct, err = gi.StructNew("Gtk", "TextTagTablePrivate")
	})
	return err
}

type TextTagTablePrivate struct {
	native uintptr
}

var textViewAccessibleClassStruct *gi.Struct
var textViewAccessibleClassStruct_Once sync.Once

func textViewAccessibleClassStruct_Set() error {
	var err error
	textViewAccessibleClassStruct_Once.Do(func() {
		textViewAccessibleClassStruct, err = gi.StructNew("Gtk", "TextViewAccessibleClass")
	})
	return err
}

type TextViewAccessibleClass struct {
	native uintptr
}

var textViewAccessiblePrivateStruct *gi.Struct
var textViewAccessiblePrivateStruct_Once sync.Once

func textViewAccessiblePrivateStruct_Set() error {
	var err error
	textViewAccessiblePrivateStruct_Once.Do(func() {
		textViewAccessiblePrivateStruct, err = gi.StructNew("Gtk", "TextViewAccessiblePrivate")
	})
	return err
}

type TextViewAccessiblePrivate struct {
	native uintptr
}

var textViewClassStruct *gi.Struct
var textViewClassStruct_Once sync.Once

func textViewClassStruct_Set() error {
	var err error
	textViewClassStruct_Once.Do(func() {
		textViewClassStruct, err = gi.StructNew("Gtk", "TextViewClass")
	})
	return err
}

type TextViewClass struct {
	native uintptr
}

var textViewPrivateStruct *gi.Struct
var textViewPrivateStruct_Once sync.Once

func textViewPrivateStruct_Set() error {
	var err error
	textViewPrivateStruct_Once.Do(func() {
		textViewPrivateStruct, err = gi.StructNew("Gtk", "TextViewPrivate")
	})
	return err
}

type TextViewPrivate struct {
	native uintptr
}

var themeEngineStruct *gi.Struct
var themeEngineStruct_Once sync.Once

func themeEngineStruct_Set() error {
	var err error
	themeEngineStruct_Once.Do(func() {
		themeEngineStruct, err = gi.StructNew("Gtk", "ThemeEngine")
	})
	return err
}

type ThemeEngine struct {
	native uintptr
}

var themingEngineClassStruct *gi.Struct
var themingEngineClassStruct_Once sync.Once

func themingEngineClassStruct_Set() error {
	var err error
	themingEngineClassStruct_Once.Do(func() {
		themingEngineClassStruct, err = gi.StructNew("Gtk", "ThemingEngineClass")
	})
	return err
}

type ThemingEngineClass struct {
	native uintptr
}

var themingEnginePrivateStruct *gi.Struct
var themingEnginePrivateStruct_Once sync.Once

func themingEnginePrivateStruct_Set() error {
	var err error
	themingEnginePrivateStruct_Once.Do(func() {
		themingEnginePrivateStruct, err = gi.StructNew("Gtk", "ThemingEnginePrivate")
	})
	return err
}

type ThemingEnginePrivate struct {
	native uintptr
}

var toggleActionClassStruct *gi.Struct
var toggleActionClassStruct_Once sync.Once

func toggleActionClassStruct_Set() error {
	var err error
	toggleActionClassStruct_Once.Do(func() {
		toggleActionClassStruct, err = gi.StructNew("Gtk", "ToggleActionClass")
	})
	return err
}

type ToggleActionClass struct {
	native uintptr
}

var toggleActionEntryStruct *gi.Struct
var toggleActionEntryStruct_Once sync.Once

func toggleActionEntryStruct_Set() error {
	var err error
	toggleActionEntryStruct_Once.Do(func() {
		toggleActionEntryStruct, err = gi.StructNew("Gtk", "ToggleActionEntry")
	})
	return err
}

type ToggleActionEntry struct {
	native uintptr
}

var toggleActionPrivateStruct *gi.Struct
var toggleActionPrivateStruct_Once sync.Once

func toggleActionPrivateStruct_Set() error {
	var err error
	toggleActionPrivateStruct_Once.Do(func() {
		toggleActionPrivateStruct, err = gi.StructNew("Gtk", "ToggleActionPrivate")
	})
	return err
}

type ToggleActionPrivate struct {
	native uintptr
}

var toggleButtonAccessibleClassStruct *gi.Struct
var toggleButtonAccessibleClassStruct_Once sync.Once

func toggleButtonAccessibleClassStruct_Set() error {
	var err error
	toggleButtonAccessibleClassStruct_Once.Do(func() {
		toggleButtonAccessibleClassStruct, err = gi.StructNew("Gtk", "ToggleButtonAccessibleClass")
	})
	return err
}

type ToggleButtonAccessibleClass struct {
	native uintptr
}

var toggleButtonAccessiblePrivateStruct *gi.Struct
var toggleButtonAccessiblePrivateStruct_Once sync.Once

func toggleButtonAccessiblePrivateStruct_Set() error {
	var err error
	toggleButtonAccessiblePrivateStruct_Once.Do(func() {
		toggleButtonAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ToggleButtonAccessiblePrivate")
	})
	return err
}

type ToggleButtonAccessiblePrivate struct {
	native uintptr
}

var toggleButtonClassStruct *gi.Struct
var toggleButtonClassStruct_Once sync.Once

func toggleButtonClassStruct_Set() error {
	var err error
	toggleButtonClassStruct_Once.Do(func() {
		toggleButtonClassStruct, err = gi.StructNew("Gtk", "ToggleButtonClass")
	})
	return err
}

type ToggleButtonClass struct {
	native uintptr
}

var toggleButtonPrivateStruct *gi.Struct
var toggleButtonPrivateStruct_Once sync.Once

func toggleButtonPrivateStruct_Set() error {
	var err error
	toggleButtonPrivateStruct_Once.Do(func() {
		toggleButtonPrivateStruct, err = gi.StructNew("Gtk", "ToggleButtonPrivate")
	})
	return err
}

type ToggleButtonPrivate struct {
	native uintptr
}

var toggleToolButtonClassStruct *gi.Struct
var toggleToolButtonClassStruct_Once sync.Once

func toggleToolButtonClassStruct_Set() error {
	var err error
	toggleToolButtonClassStruct_Once.Do(func() {
		toggleToolButtonClassStruct, err = gi.StructNew("Gtk", "ToggleToolButtonClass")
	})
	return err
}

type ToggleToolButtonClass struct {
	native uintptr
}

var toggleToolButtonPrivateStruct *gi.Struct
var toggleToolButtonPrivateStruct_Once sync.Once

func toggleToolButtonPrivateStruct_Set() error {
	var err error
	toggleToolButtonPrivateStruct_Once.Do(func() {
		toggleToolButtonPrivateStruct, err = gi.StructNew("Gtk", "ToggleToolButtonPrivate")
	})
	return err
}

type ToggleToolButtonPrivate struct {
	native uintptr
}

var toolButtonClassStruct *gi.Struct
var toolButtonClassStruct_Once sync.Once

func toolButtonClassStruct_Set() error {
	var err error
	toolButtonClassStruct_Once.Do(func() {
		toolButtonClassStruct, err = gi.StructNew("Gtk", "ToolButtonClass")
	})
	return err
}

type ToolButtonClass struct {
	native uintptr
}

var toolButtonPrivateStruct *gi.Struct
var toolButtonPrivateStruct_Once sync.Once

func toolButtonPrivateStruct_Set() error {
	var err error
	toolButtonPrivateStruct_Once.Do(func() {
		toolButtonPrivateStruct, err = gi.StructNew("Gtk", "ToolButtonPrivate")
	})
	return err
}

type ToolButtonPrivate struct {
	native uintptr
}

var toolItemClassStruct *gi.Struct
var toolItemClassStruct_Once sync.Once

func toolItemClassStruct_Set() error {
	var err error
	toolItemClassStruct_Once.Do(func() {
		toolItemClassStruct, err = gi.StructNew("Gtk", "ToolItemClass")
	})
	return err
}

type ToolItemClass struct {
	native uintptr
}

var toolItemGroupClassStruct *gi.Struct
var toolItemGroupClassStruct_Once sync.Once

func toolItemGroupClassStruct_Set() error {
	var err error
	toolItemGroupClassStruct_Once.Do(func() {
		toolItemGroupClassStruct, err = gi.StructNew("Gtk", "ToolItemGroupClass")
	})
	return err
}

type ToolItemGroupClass struct {
	native uintptr
}

var toolItemGroupPrivateStruct *gi.Struct
var toolItemGroupPrivateStruct_Once sync.Once

func toolItemGroupPrivateStruct_Set() error {
	var err error
	toolItemGroupPrivateStruct_Once.Do(func() {
		toolItemGroupPrivateStruct, err = gi.StructNew("Gtk", "ToolItemGroupPrivate")
	})
	return err
}

type ToolItemGroupPrivate struct {
	native uintptr
}

var toolItemPrivateStruct *gi.Struct
var toolItemPrivateStruct_Once sync.Once

func toolItemPrivateStruct_Set() error {
	var err error
	toolItemPrivateStruct_Once.Do(func() {
		toolItemPrivateStruct, err = gi.StructNew("Gtk", "ToolItemPrivate")
	})
	return err
}

type ToolItemPrivate struct {
	native uintptr
}

var toolPaletteClassStruct *gi.Struct
var toolPaletteClassStruct_Once sync.Once

func toolPaletteClassStruct_Set() error {
	var err error
	toolPaletteClassStruct_Once.Do(func() {
		toolPaletteClassStruct, err = gi.StructNew("Gtk", "ToolPaletteClass")
	})
	return err
}

type ToolPaletteClass struct {
	native uintptr
}

var toolPalettePrivateStruct *gi.Struct
var toolPalettePrivateStruct_Once sync.Once

func toolPalettePrivateStruct_Set() error {
	var err error
	toolPalettePrivateStruct_Once.Do(func() {
		toolPalettePrivateStruct, err = gi.StructNew("Gtk", "ToolPalettePrivate")
	})
	return err
}

type ToolPalettePrivate struct {
	native uintptr
}

var toolShellIfaceStruct *gi.Struct
var toolShellIfaceStruct_Once sync.Once

func toolShellIfaceStruct_Set() error {
	var err error
	toolShellIfaceStruct_Once.Do(func() {
		toolShellIfaceStruct, err = gi.StructNew("Gtk", "ToolShellIface")
	})
	return err
}

type ToolShellIface struct {
	native uintptr
}

var toolbarClassStruct *gi.Struct
var toolbarClassStruct_Once sync.Once

func toolbarClassStruct_Set() error {
	var err error
	toolbarClassStruct_Once.Do(func() {
		toolbarClassStruct, err = gi.StructNew("Gtk", "ToolbarClass")
	})
	return err
}

type ToolbarClass struct {
	native uintptr
}

var toolbarPrivateStruct *gi.Struct
var toolbarPrivateStruct_Once sync.Once

func toolbarPrivateStruct_Set() error {
	var err error
	toolbarPrivateStruct_Once.Do(func() {
		toolbarPrivateStruct, err = gi.StructNew("Gtk", "ToolbarPrivate")
	})
	return err
}

type ToolbarPrivate struct {
	native uintptr
}

var toplevelAccessibleClassStruct *gi.Struct
var toplevelAccessibleClassStruct_Once sync.Once

func toplevelAccessibleClassStruct_Set() error {
	var err error
	toplevelAccessibleClassStruct_Once.Do(func() {
		toplevelAccessibleClassStruct, err = gi.StructNew("Gtk", "ToplevelAccessibleClass")
	})
	return err
}

type ToplevelAccessibleClass struct {
	native uintptr
}

var toplevelAccessiblePrivateStruct *gi.Struct
var toplevelAccessiblePrivateStruct_Once sync.Once

func toplevelAccessiblePrivateStruct_Set() error {
	var err error
	toplevelAccessiblePrivateStruct_Once.Do(func() {
		toplevelAccessiblePrivateStruct, err = gi.StructNew("Gtk", "ToplevelAccessiblePrivate")
	})
	return err
}

type ToplevelAccessiblePrivate struct {
	native uintptr
}

var treeDragDestIfaceStruct *gi.Struct
var treeDragDestIfaceStruct_Once sync.Once

func treeDragDestIfaceStruct_Set() error {
	var err error
	treeDragDestIfaceStruct_Once.Do(func() {
		treeDragDestIfaceStruct, err = gi.StructNew("Gtk", "TreeDragDestIface")
	})
	return err
}

type TreeDragDestIface struct {
	native uintptr
}

var treeDragSourceIfaceStruct *gi.Struct
var treeDragSourceIfaceStruct_Once sync.Once

func treeDragSourceIfaceStruct_Set() error {
	var err error
	treeDragSourceIfaceStruct_Once.Do(func() {
		treeDragSourceIfaceStruct, err = gi.StructNew("Gtk", "TreeDragSourceIface")
	})
	return err
}

type TreeDragSourceIface struct {
	native uintptr
}

var treeIterStruct *gi.Struct
var treeIterStruct_Once sync.Once

func treeIterStruct_Set() error {
	var err error
	treeIterStruct_Once.Do(func() {
		treeIterStruct, err = gi.StructNew("Gtk", "TreeIter")
	})
	return err
}

type TreeIter struct {
	native uintptr
}

var treeIterCopyFunction *gi.Function
var treeIterCopyFunction_Once sync.Once

func treeIterCopyFunction_Set() error {
	var err error
	treeIterCopyFunction_Once.Do(func() {
		err = treeIterStruct_Set()
		if err != nil {
			return
		}
		treeIterCopyFunction, err = treeIterStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_tree_iter_copy.
func (recv *TreeIter) Copy() *TreeIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treeIterCopyFunction_Set()
	if err == nil {
		ret = treeIterCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TreeIter{native: ret.Pointer()}

	return retGo
}

var treeIterFreeFunction *gi.Function
var treeIterFreeFunction_Once sync.Once

func treeIterFreeFunction_Set() error {
	var err error
	treeIterFreeFunction_Once.Do(func() {
		err = treeIterStruct_Set()
		if err != nil {
			return
		}
		treeIterFreeFunction, err = treeIterStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_tree_iter_free.
func (recv *TreeIter) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := treeIterFreeFunction_Set()
	if err == nil {
		treeIterFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treeModelFilterClassStruct *gi.Struct
var treeModelFilterClassStruct_Once sync.Once

func treeModelFilterClassStruct_Set() error {
	var err error
	treeModelFilterClassStruct_Once.Do(func() {
		treeModelFilterClassStruct, err = gi.StructNew("Gtk", "TreeModelFilterClass")
	})
	return err
}

type TreeModelFilterClass struct {
	native uintptr
}

var treeModelFilterPrivateStruct *gi.Struct
var treeModelFilterPrivateStruct_Once sync.Once

func treeModelFilterPrivateStruct_Set() error {
	var err error
	treeModelFilterPrivateStruct_Once.Do(func() {
		treeModelFilterPrivateStruct, err = gi.StructNew("Gtk", "TreeModelFilterPrivate")
	})
	return err
}

type TreeModelFilterPrivate struct {
	native uintptr
}

var treeModelIfaceStruct *gi.Struct
var treeModelIfaceStruct_Once sync.Once

func treeModelIfaceStruct_Set() error {
	var err error
	treeModelIfaceStruct_Once.Do(func() {
		treeModelIfaceStruct, err = gi.StructNew("Gtk", "TreeModelIface")
	})
	return err
}

type TreeModelIface struct {
	native uintptr
}

var treeModelSortClassStruct *gi.Struct
var treeModelSortClassStruct_Once sync.Once

func treeModelSortClassStruct_Set() error {
	var err error
	treeModelSortClassStruct_Once.Do(func() {
		treeModelSortClassStruct, err = gi.StructNew("Gtk", "TreeModelSortClass")
	})
	return err
}

type TreeModelSortClass struct {
	native uintptr
}

var treeModelSortPrivateStruct *gi.Struct
var treeModelSortPrivateStruct_Once sync.Once

func treeModelSortPrivateStruct_Set() error {
	var err error
	treeModelSortPrivateStruct_Once.Do(func() {
		treeModelSortPrivateStruct, err = gi.StructNew("Gtk", "TreeModelSortPrivate")
	})
	return err
}

type TreeModelSortPrivate struct {
	native uintptr
}

var treePathStruct *gi.Struct
var treePathStruct_Once sync.Once

func treePathStruct_Set() error {
	var err error
	treePathStruct_Once.Do(func() {
		treePathStruct, err = gi.StructNew("Gtk", "TreePath")
	})
	return err
}

type TreePath struct {
	native uintptr
}

var treePathNewFunction *gi.Function
var treePathNewFunction_Once sync.Once

func treePathNewFunction_Set() error {
	var err error
	treePathNewFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathNewFunction, err = treePathStruct.InvokerNew("new")
	})
	return err
}

// TreePathNew is a representation of the C type gtk_tree_path_new.
func TreePathNew() *TreePath {

	var ret gi.Argument

	err := treePathNewFunction_Set()
	if err == nil {
		ret = treePathNewFunction.Invoke(nil, nil)
	}

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var treePathNewFirstFunction *gi.Function
var treePathNewFirstFunction_Once sync.Once

func treePathNewFirstFunction_Set() error {
	var err error
	treePathNewFirstFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathNewFirstFunction, err = treePathStruct.InvokerNew("new_first")
	})
	return err
}

// TreePathNewFirst is a representation of the C type gtk_tree_path_new_first.
func TreePathNewFirst() *TreePath {

	var ret gi.Argument

	err := treePathNewFirstFunction_Set()
	if err == nil {
		ret = treePathNewFirstFunction.Invoke(nil, nil)
	}

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_tree_path_new_from_indices' : parameter '...' has no type

// UNSUPPORTED : C value 'gtk_tree_path_new_from_indicesv' : parameter 'indices' has no type

var treePathNewFromStringFunction *gi.Function
var treePathNewFromStringFunction_Once sync.Once

func treePathNewFromStringFunction_Set() error {
	var err error
	treePathNewFromStringFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathNewFromStringFunction, err = treePathStruct.InvokerNew("new_from_string")
	})
	return err
}

// TreePathNewFromString is a representation of the C type gtk_tree_path_new_from_string.
func TreePathNewFromString(path string) *TreePath {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(path)

	var ret gi.Argument

	err := treePathNewFromStringFunction_Set()
	if err == nil {
		ret = treePathNewFromStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var treePathAppendIndexFunction *gi.Function
var treePathAppendIndexFunction_Once sync.Once

func treePathAppendIndexFunction_Set() error {
	var err error
	treePathAppendIndexFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathAppendIndexFunction, err = treePathStruct.InvokerNew("append_index")
	})
	return err
}

// AppendIndex is a representation of the C type gtk_tree_path_append_index.
func (recv *TreePath) AppendIndex(index int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)

	err := treePathAppendIndexFunction_Set()
	if err == nil {
		treePathAppendIndexFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treePathCompareFunction *gi.Function
var treePathCompareFunction_Once sync.Once

func treePathCompareFunction_Set() error {
	var err error
	treePathCompareFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathCompareFunction, err = treePathStruct.InvokerNew("compare")
	})
	return err
}

// Compare is a representation of the C type gtk_tree_path_compare.
func (recv *TreePath) Compare(b *TreePath) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(b.native)

	var ret gi.Argument

	err := treePathCompareFunction_Set()
	if err == nil {
		ret = treePathCompareFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var treePathCopyFunction *gi.Function
var treePathCopyFunction_Once sync.Once

func treePathCopyFunction_Set() error {
	var err error
	treePathCopyFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathCopyFunction, err = treePathStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_tree_path_copy.
func (recv *TreePath) Copy() *TreePath {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treePathCopyFunction_Set()
	if err == nil {
		ret = treePathCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var treePathDownFunction *gi.Function
var treePathDownFunction_Once sync.Once

func treePathDownFunction_Set() error {
	var err error
	treePathDownFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathDownFunction, err = treePathStruct.InvokerNew("down")
	})
	return err
}

// Down is a representation of the C type gtk_tree_path_down.
func (recv *TreePath) Down() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := treePathDownFunction_Set()
	if err == nil {
		treePathDownFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treePathFreeFunction *gi.Function
var treePathFreeFunction_Once sync.Once

func treePathFreeFunction_Set() error {
	var err error
	treePathFreeFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathFreeFunction, err = treePathStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_tree_path_free.
func (recv *TreePath) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := treePathFreeFunction_Set()
	if err == nil {
		treePathFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treePathGetDepthFunction *gi.Function
var treePathGetDepthFunction_Once sync.Once

func treePathGetDepthFunction_Set() error {
	var err error
	treePathGetDepthFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathGetDepthFunction, err = treePathStruct.InvokerNew("get_depth")
	})
	return err
}

// GetDepth is a representation of the C type gtk_tree_path_get_depth.
func (recv *TreePath) GetDepth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treePathGetDepthFunction_Set()
	if err == nil {
		ret = treePathGetDepthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var treePathGetIndicesFunction *gi.Function
var treePathGetIndicesFunction_Once sync.Once

func treePathGetIndicesFunction_Set() error {
	var err error
	treePathGetIndicesFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathGetIndicesFunction, err = treePathStruct.InvokerNew("get_indices")
	})
	return err
}

// GetIndices is a representation of the C type gtk_tree_path_get_indices.
func (recv *TreePath) GetIndices() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treePathGetIndicesFunction_Set()
	if err == nil {
		ret = treePathGetIndicesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var treePathGetIndicesWithDepthFunction *gi.Function
var treePathGetIndicesWithDepthFunction_Once sync.Once

func treePathGetIndicesWithDepthFunction_Set() error {
	var err error
	treePathGetIndicesWithDepthFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathGetIndicesWithDepthFunction, err = treePathStruct.InvokerNew("get_indices_with_depth")
	})
	return err
}

// GetIndicesWithDepth is a representation of the C type gtk_tree_path_get_indices_with_depth.
func (recv *TreePath) GetIndicesWithDepth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := treePathGetIndicesWithDepthFunction_Set()
	if err == nil {
		treePathGetIndicesWithDepthFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var treePathIsAncestorFunction *gi.Function
var treePathIsAncestorFunction_Once sync.Once

func treePathIsAncestorFunction_Set() error {
	var err error
	treePathIsAncestorFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathIsAncestorFunction, err = treePathStruct.InvokerNew("is_ancestor")
	})
	return err
}

// IsAncestor is a representation of the C type gtk_tree_path_is_ancestor.
func (recv *TreePath) IsAncestor(descendant *TreePath) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(descendant.native)

	var ret gi.Argument

	err := treePathIsAncestorFunction_Set()
	if err == nil {
		ret = treePathIsAncestorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treePathIsDescendantFunction *gi.Function
var treePathIsDescendantFunction_Once sync.Once

func treePathIsDescendantFunction_Set() error {
	var err error
	treePathIsDescendantFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathIsDescendantFunction, err = treePathStruct.InvokerNew("is_descendant")
	})
	return err
}

// IsDescendant is a representation of the C type gtk_tree_path_is_descendant.
func (recv *TreePath) IsDescendant(ancestor *TreePath) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(ancestor.native)

	var ret gi.Argument

	err := treePathIsDescendantFunction_Set()
	if err == nil {
		ret = treePathIsDescendantFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treePathNextFunction *gi.Function
var treePathNextFunction_Once sync.Once

func treePathNextFunction_Set() error {
	var err error
	treePathNextFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathNextFunction, err = treePathStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type gtk_tree_path_next.
func (recv *TreePath) Next() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := treePathNextFunction_Set()
	if err == nil {
		treePathNextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treePathPrependIndexFunction *gi.Function
var treePathPrependIndexFunction_Once sync.Once

func treePathPrependIndexFunction_Set() error {
	var err error
	treePathPrependIndexFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathPrependIndexFunction, err = treePathStruct.InvokerNew("prepend_index")
	})
	return err
}

// PrependIndex is a representation of the C type gtk_tree_path_prepend_index.
func (recv *TreePath) PrependIndex(index int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)

	err := treePathPrependIndexFunction_Set()
	if err == nil {
		treePathPrependIndexFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treePathPrevFunction *gi.Function
var treePathPrevFunction_Once sync.Once

func treePathPrevFunction_Set() error {
	var err error
	treePathPrevFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathPrevFunction, err = treePathStruct.InvokerNew("prev")
	})
	return err
}

// Prev is a representation of the C type gtk_tree_path_prev.
func (recv *TreePath) Prev() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treePathPrevFunction_Set()
	if err == nil {
		ret = treePathPrevFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treePathToStringFunction *gi.Function
var treePathToStringFunction_Once sync.Once

func treePathToStringFunction_Set() error {
	var err error
	treePathToStringFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathToStringFunction, err = treePathStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type gtk_tree_path_to_string.
func (recv *TreePath) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treePathToStringFunction_Set()
	if err == nil {
		ret = treePathToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var treePathUpFunction *gi.Function
var treePathUpFunction_Once sync.Once

func treePathUpFunction_Set() error {
	var err error
	treePathUpFunction_Once.Do(func() {
		err = treePathStruct_Set()
		if err != nil {
			return
		}
		treePathUpFunction, err = treePathStruct.InvokerNew("up")
	})
	return err
}

// Up is a representation of the C type gtk_tree_path_up.
func (recv *TreePath) Up() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treePathUpFunction_Set()
	if err == nil {
		ret = treePathUpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeRowReferenceStruct *gi.Struct
var treeRowReferenceStruct_Once sync.Once

func treeRowReferenceStruct_Set() error {
	var err error
	treeRowReferenceStruct_Once.Do(func() {
		treeRowReferenceStruct, err = gi.StructNew("Gtk", "TreeRowReference")
	})
	return err
}

type TreeRowReference struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_tree_row_reference_new' : parameter 'model' of type 'TreeModel' not supported

// UNSUPPORTED : C value 'gtk_tree_row_reference_new_proxy' : parameter 'proxy' of type 'GObject.Object' not supported

var treeRowReferenceCopyFunction *gi.Function
var treeRowReferenceCopyFunction_Once sync.Once

func treeRowReferenceCopyFunction_Set() error {
	var err error
	treeRowReferenceCopyFunction_Once.Do(func() {
		err = treeRowReferenceStruct_Set()
		if err != nil {
			return
		}
		treeRowReferenceCopyFunction, err = treeRowReferenceStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_tree_row_reference_copy.
func (recv *TreeRowReference) Copy() *TreeRowReference {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treeRowReferenceCopyFunction_Set()
	if err == nil {
		ret = treeRowReferenceCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TreeRowReference{native: ret.Pointer()}

	return retGo
}

var treeRowReferenceFreeFunction *gi.Function
var treeRowReferenceFreeFunction_Once sync.Once

func treeRowReferenceFreeFunction_Set() error {
	var err error
	treeRowReferenceFreeFunction_Once.Do(func() {
		err = treeRowReferenceStruct_Set()
		if err != nil {
			return
		}
		treeRowReferenceFreeFunction, err = treeRowReferenceStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_tree_row_reference_free.
func (recv *TreeRowReference) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := treeRowReferenceFreeFunction_Set()
	if err == nil {
		treeRowReferenceFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_tree_row_reference_get_model' : return type 'TreeModel' not supported

var treeRowReferenceGetPathFunction *gi.Function
var treeRowReferenceGetPathFunction_Once sync.Once

func treeRowReferenceGetPathFunction_Set() error {
	var err error
	treeRowReferenceGetPathFunction_Once.Do(func() {
		err = treeRowReferenceStruct_Set()
		if err != nil {
			return
		}
		treeRowReferenceGetPathFunction, err = treeRowReferenceStruct.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type gtk_tree_row_reference_get_path.
func (recv *TreeRowReference) GetPath() *TreePath {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treeRowReferenceGetPathFunction_Set()
	if err == nil {
		ret = treeRowReferenceGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TreePath{native: ret.Pointer()}

	return retGo
}

var treeRowReferenceValidFunction *gi.Function
var treeRowReferenceValidFunction_Once sync.Once

func treeRowReferenceValidFunction_Set() error {
	var err error
	treeRowReferenceValidFunction_Once.Do(func() {
		err = treeRowReferenceStruct_Set()
		if err != nil {
			return
		}
		treeRowReferenceValidFunction, err = treeRowReferenceStruct.InvokerNew("valid")
	})
	return err
}

// Valid is a representation of the C type gtk_tree_row_reference_valid.
func (recv *TreeRowReference) Valid() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := treeRowReferenceValidFunction_Set()
	if err == nil {
		ret = treeRowReferenceValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeSelectionClassStruct *gi.Struct
var treeSelectionClassStruct_Once sync.Once

func treeSelectionClassStruct_Set() error {
	var err error
	treeSelectionClassStruct_Once.Do(func() {
		treeSelectionClassStruct, err = gi.StructNew("Gtk", "TreeSelectionClass")
	})
	return err
}

type TreeSelectionClass struct {
	native uintptr
}

var treeSelectionPrivateStruct *gi.Struct
var treeSelectionPrivateStruct_Once sync.Once

func treeSelectionPrivateStruct_Set() error {
	var err error
	treeSelectionPrivateStruct_Once.Do(func() {
		treeSelectionPrivateStruct, err = gi.StructNew("Gtk", "TreeSelectionPrivate")
	})
	return err
}

type TreeSelectionPrivate struct {
	native uintptr
}

var treeSortableIfaceStruct *gi.Struct
var treeSortableIfaceStruct_Once sync.Once

func treeSortableIfaceStruct_Set() error {
	var err error
	treeSortableIfaceStruct_Once.Do(func() {
		treeSortableIfaceStruct, err = gi.StructNew("Gtk", "TreeSortableIface")
	})
	return err
}

type TreeSortableIface struct {
	native uintptr
}

var treeStoreClassStruct *gi.Struct
var treeStoreClassStruct_Once sync.Once

func treeStoreClassStruct_Set() error {
	var err error
	treeStoreClassStruct_Once.Do(func() {
		treeStoreClassStruct, err = gi.StructNew("Gtk", "TreeStoreClass")
	})
	return err
}

type TreeStoreClass struct {
	native uintptr
}

var treeStorePrivateStruct *gi.Struct
var treeStorePrivateStruct_Once sync.Once

func treeStorePrivateStruct_Set() error {
	var err error
	treeStorePrivateStruct_Once.Do(func() {
		treeStorePrivateStruct, err = gi.StructNew("Gtk", "TreeStorePrivate")
	})
	return err
}

type TreeStorePrivate struct {
	native uintptr
}

var treeViewAccessibleClassStruct *gi.Struct
var treeViewAccessibleClassStruct_Once sync.Once

func treeViewAccessibleClassStruct_Set() error {
	var err error
	treeViewAccessibleClassStruct_Once.Do(func() {
		treeViewAccessibleClassStruct, err = gi.StructNew("Gtk", "TreeViewAccessibleClass")
	})
	return err
}

type TreeViewAccessibleClass struct {
	native uintptr
}

var treeViewAccessiblePrivateStruct *gi.Struct
var treeViewAccessiblePrivateStruct_Once sync.Once

func treeViewAccessiblePrivateStruct_Set() error {
	var err error
	treeViewAccessiblePrivateStruct_Once.Do(func() {
		treeViewAccessiblePrivateStruct, err = gi.StructNew("Gtk", "TreeViewAccessiblePrivate")
	})
	return err
}

type TreeViewAccessiblePrivate struct {
	native uintptr
}

var treeViewClassStruct *gi.Struct
var treeViewClassStruct_Once sync.Once

func treeViewClassStruct_Set() error {
	var err error
	treeViewClassStruct_Once.Do(func() {
		treeViewClassStruct, err = gi.StructNew("Gtk", "TreeViewClass")
	})
	return err
}

type TreeViewClass struct {
	native uintptr
}

var treeViewColumnClassStruct *gi.Struct
var treeViewColumnClassStruct_Once sync.Once

func treeViewColumnClassStruct_Set() error {
	var err error
	treeViewColumnClassStruct_Once.Do(func() {
		treeViewColumnClassStruct, err = gi.StructNew("Gtk", "TreeViewColumnClass")
	})
	return err
}

type TreeViewColumnClass struct {
	native uintptr
}

var treeViewColumnPrivateStruct *gi.Struct
var treeViewColumnPrivateStruct_Once sync.Once

func treeViewColumnPrivateStruct_Set() error {
	var err error
	treeViewColumnPrivateStruct_Once.Do(func() {
		treeViewColumnPrivateStruct, err = gi.StructNew("Gtk", "TreeViewColumnPrivate")
	})
	return err
}

type TreeViewColumnPrivate struct {
	native uintptr
}

var treeViewPrivateStruct *gi.Struct
var treeViewPrivateStruct_Once sync.Once

func treeViewPrivateStruct_Set() error {
	var err error
	treeViewPrivateStruct_Once.Do(func() {
		treeViewPrivateStruct, err = gi.StructNew("Gtk", "TreeViewPrivate")
	})
	return err
}

type TreeViewPrivate struct {
	native uintptr
}

var uIManagerClassStruct *gi.Struct
var uIManagerClassStruct_Once sync.Once

func uIManagerClassStruct_Set() error {
	var err error
	uIManagerClassStruct_Once.Do(func() {
		uIManagerClassStruct, err = gi.StructNew("Gtk", "UIManagerClass")
	})
	return err
}

type UIManagerClass struct {
	native uintptr
}

var uIManagerPrivateStruct *gi.Struct
var uIManagerPrivateStruct_Once sync.Once

func uIManagerPrivateStruct_Set() error {
	var err error
	uIManagerPrivateStruct_Once.Do(func() {
		uIManagerPrivateStruct, err = gi.StructNew("Gtk", "UIManagerPrivate")
	})
	return err
}

type UIManagerPrivate struct {
	native uintptr
}

var vBoxClassStruct *gi.Struct
var vBoxClassStruct_Once sync.Once

func vBoxClassStruct_Set() error {
	var err error
	vBoxClassStruct_Once.Do(func() {
		vBoxClassStruct, err = gi.StructNew("Gtk", "VBoxClass")
	})
	return err
}

type VBoxClass struct {
	native uintptr
}

var vButtonBoxClassStruct *gi.Struct
var vButtonBoxClassStruct_Once sync.Once

func vButtonBoxClassStruct_Set() error {
	var err error
	vButtonBoxClassStruct_Once.Do(func() {
		vButtonBoxClassStruct, err = gi.StructNew("Gtk", "VButtonBoxClass")
	})
	return err
}

type VButtonBoxClass struct {
	native uintptr
}

var vPanedClassStruct *gi.Struct
var vPanedClassStruct_Once sync.Once

func vPanedClassStruct_Set() error {
	var err error
	vPanedClassStruct_Once.Do(func() {
		vPanedClassStruct, err = gi.StructNew("Gtk", "VPanedClass")
	})
	return err
}

type VPanedClass struct {
	native uintptr
}

var vScaleClassStruct *gi.Struct
var vScaleClassStruct_Once sync.Once

func vScaleClassStruct_Set() error {
	var err error
	vScaleClassStruct_Once.Do(func() {
		vScaleClassStruct, err = gi.StructNew("Gtk", "VScaleClass")
	})
	return err
}

type VScaleClass struct {
	native uintptr
}

var vScrollbarClassStruct *gi.Struct
var vScrollbarClassStruct_Once sync.Once

func vScrollbarClassStruct_Set() error {
	var err error
	vScrollbarClassStruct_Once.Do(func() {
		vScrollbarClassStruct, err = gi.StructNew("Gtk", "VScrollbarClass")
	})
	return err
}

type VScrollbarClass struct {
	native uintptr
}

var vSeparatorClassStruct *gi.Struct
var vSeparatorClassStruct_Once sync.Once

func vSeparatorClassStruct_Set() error {
	var err error
	vSeparatorClassStruct_Once.Do(func() {
		vSeparatorClassStruct, err = gi.StructNew("Gtk", "VSeparatorClass")
	})
	return err
}

type VSeparatorClass struct {
	native uintptr
}

var viewportClassStruct *gi.Struct
var viewportClassStruct_Once sync.Once

func viewportClassStruct_Set() error {
	var err error
	viewportClassStruct_Once.Do(func() {
		viewportClassStruct, err = gi.StructNew("Gtk", "ViewportClass")
	})
	return err
}

type ViewportClass struct {
	native uintptr
}

var viewportPrivateStruct *gi.Struct
var viewportPrivateStruct_Once sync.Once

func viewportPrivateStruct_Set() error {
	var err error
	viewportPrivateStruct_Once.Do(func() {
		viewportPrivateStruct, err = gi.StructNew("Gtk", "ViewportPrivate")
	})
	return err
}

type ViewportPrivate struct {
	native uintptr
}

var volumeButtonClassStruct *gi.Struct
var volumeButtonClassStruct_Once sync.Once

func volumeButtonClassStruct_Set() error {
	var err error
	volumeButtonClassStruct_Once.Do(func() {
		volumeButtonClassStruct, err = gi.StructNew("Gtk", "VolumeButtonClass")
	})
	return err
}

type VolumeButtonClass struct {
	native uintptr
}

var widgetAccessibleClassStruct *gi.Struct
var widgetAccessibleClassStruct_Once sync.Once

func widgetAccessibleClassStruct_Set() error {
	var err error
	widgetAccessibleClassStruct_Once.Do(func() {
		widgetAccessibleClassStruct, err = gi.StructNew("Gtk", "WidgetAccessibleClass")
	})
	return err
}

type WidgetAccessibleClass struct {
	native uintptr
}

var widgetAccessiblePrivateStruct *gi.Struct
var widgetAccessiblePrivateStruct_Once sync.Once

func widgetAccessiblePrivateStruct_Set() error {
	var err error
	widgetAccessiblePrivateStruct_Once.Do(func() {
		widgetAccessiblePrivateStruct, err = gi.StructNew("Gtk", "WidgetAccessiblePrivate")
	})
	return err
}

type WidgetAccessiblePrivate struct {
	native uintptr
}

var widgetClassStruct *gi.Struct
var widgetClassStruct_Once sync.Once

func widgetClassStruct_Set() error {
	var err error
	widgetClassStruct_Once.Do(func() {
		widgetClassStruct, err = gi.StructNew("Gtk", "WidgetClass")
	})
	return err
}

type WidgetClass struct {
	native uintptr
}

// UNSUPPORTED : C value 'gtk_widget_class_bind_template_callback_full' : parameter 'callback_symbol' of type 'GObject.Callback' not supported

var widgetClassBindTemplateChildFullFunction *gi.Function
var widgetClassBindTemplateChildFullFunction_Once sync.Once

func widgetClassBindTemplateChildFullFunction_Set() error {
	var err error
	widgetClassBindTemplateChildFullFunction_Once.Do(func() {
		err = widgetClassStruct_Set()
		if err != nil {
			return
		}
		widgetClassBindTemplateChildFullFunction, err = widgetClassStruct.InvokerNew("bind_template_child_full")
	})
	return err
}

// BindTemplateChildFull is a representation of the C type gtk_widget_class_bind_template_child_full.
func (recv *WidgetClass) BindTemplateChildFull(name string, internalChild bool, structOffset int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)
	inArgs[2].SetBoolean(internalChild)
	inArgs[3].SetInt32(structOffset)

	err := widgetClassBindTemplateChildFullFunction_Set()
	if err == nil {
		widgetClassBindTemplateChildFullFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_widget_class_find_style_property' : return type 'GObject.ParamSpec' not supported

var widgetClassGetCssNameFunction *gi.Function
var widgetClassGetCssNameFunction_Once sync.Once

func widgetClassGetCssNameFunction_Set() error {
	var err error
	widgetClassGetCssNameFunction_Once.Do(func() {
		err = widgetClassStruct_Set()
		if err != nil {
			return
		}
		widgetClassGetCssNameFunction, err = widgetClassStruct.InvokerNew("get_css_name")
	})
	return err
}

// GetCssName is a representation of the C type gtk_widget_class_get_css_name.
func (recv *WidgetClass) GetCssName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := widgetClassGetCssNameFunction_Set()
	if err == nil {
		ret = widgetClassGetCssNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_class_install_style_property' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

// UNSUPPORTED : C value 'gtk_widget_class_install_style_property_parser' : parameter 'pspec' of type 'GObject.ParamSpec' not supported

var widgetClassListStylePropertiesFunction *gi.Function
var widgetClassListStylePropertiesFunction_Once sync.Once

func widgetClassListStylePropertiesFunction_Set() error {
	var err error
	widgetClassListStylePropertiesFunction_Once.Do(func() {
		err = widgetClassStruct_Set()
		if err != nil {
			return
		}
		widgetClassListStylePropertiesFunction, err = widgetClassStruct.InvokerNew("list_style_properties")
	})
	return err
}

// ListStyleProperties is a representation of the C type gtk_widget_class_list_style_properties.
func (recv *WidgetClass) ListStyleProperties() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := widgetClassListStylePropertiesFunction_Set()
	if err == nil {
		widgetClassListStylePropertiesFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
}

// UNSUPPORTED : C value 'gtk_widget_class_set_accessible_role' : parameter 'role' of type 'Atk.Role' not supported

// UNSUPPORTED : C value 'gtk_widget_class_set_accessible_type' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_class_set_connect_func' : parameter 'connect_func' of type 'BuilderConnectFunc' not supported

var widgetClassSetCssNameFunction *gi.Function
var widgetClassSetCssNameFunction_Once sync.Once

func widgetClassSetCssNameFunction_Set() error {
	var err error
	widgetClassSetCssNameFunction_Once.Do(func() {
		err = widgetClassStruct_Set()
		if err != nil {
			return
		}
		widgetClassSetCssNameFunction, err = widgetClassStruct.InvokerNew("set_css_name")
	})
	return err
}

// SetCssName is a representation of the C type gtk_widget_class_set_css_name.
func (recv *WidgetClass) SetCssName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(name)

	err := widgetClassSetCssNameFunction_Set()
	if err == nil {
		widgetClassSetCssNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_widget_class_set_template' : parameter 'template_bytes' of type 'GLib.Bytes' not supported

var widgetClassSetTemplateFromResourceFunction *gi.Function
var widgetClassSetTemplateFromResourceFunction_Once sync.Once

func widgetClassSetTemplateFromResourceFunction_Set() error {
	var err error
	widgetClassSetTemplateFromResourceFunction_Once.Do(func() {
		err = widgetClassStruct_Set()
		if err != nil {
			return
		}
		widgetClassSetTemplateFromResourceFunction, err = widgetClassStruct.InvokerNew("set_template_from_resource")
	})
	return err
}

// SetTemplateFromResource is a representation of the C type gtk_widget_class_set_template_from_resource.
func (recv *WidgetClass) SetTemplateFromResource(resourceName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(resourceName)

	err := widgetClassSetTemplateFromResourceFunction_Set()
	if err == nil {
		widgetClassSetTemplateFromResourceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var widgetClassPrivateStruct *gi.Struct
var widgetClassPrivateStruct_Once sync.Once

func widgetClassPrivateStruct_Set() error {
	var err error
	widgetClassPrivateStruct_Once.Do(func() {
		widgetClassPrivateStruct, err = gi.StructNew("Gtk", "WidgetClassPrivate")
	})
	return err
}

type WidgetClassPrivate struct {
	native uintptr
}

var widgetPathStruct *gi.Struct
var widgetPathStruct_Once sync.Once

func widgetPathStruct_Set() error {
	var err error
	widgetPathStruct_Once.Do(func() {
		widgetPathStruct, err = gi.StructNew("Gtk", "WidgetPath")
	})
	return err
}

type WidgetPath struct {
	native uintptr
}

var widgetPathNewFunction *gi.Function
var widgetPathNewFunction_Once sync.Once

func widgetPathNewFunction_Set() error {
	var err error
	widgetPathNewFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathNewFunction, err = widgetPathStruct.InvokerNew("new")
	})
	return err
}

// WidgetPathNew is a representation of the C type gtk_widget_path_new.
func WidgetPathNew() *WidgetPath {

	var ret gi.Argument

	err := widgetPathNewFunction_Set()
	if err == nil {
		ret = widgetPathNewFunction.Invoke(nil, nil)
	}

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_append_for_widget' : parameter 'widget' of type 'Widget' not supported

// UNSUPPORTED : C value 'gtk_widget_path_append_type' : parameter 'type' of type 'GType' not supported

var widgetPathAppendWithSiblingsFunction *gi.Function
var widgetPathAppendWithSiblingsFunction_Once sync.Once

func widgetPathAppendWithSiblingsFunction_Set() error {
	var err error
	widgetPathAppendWithSiblingsFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathAppendWithSiblingsFunction, err = widgetPathStruct.InvokerNew("append_with_siblings")
	})
	return err
}

// AppendWithSiblings is a representation of the C type gtk_widget_path_append_with_siblings.
func (recv *WidgetPath) AppendWithSiblings(siblings *WidgetPath, siblingIndex uint32) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(siblings.native)
	inArgs[2].SetUint32(siblingIndex)

	var ret gi.Argument

	err := widgetPathAppendWithSiblingsFunction_Set()
	if err == nil {
		ret = widgetPathAppendWithSiblingsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var widgetPathCopyFunction *gi.Function
var widgetPathCopyFunction_Once sync.Once

func widgetPathCopyFunction_Set() error {
	var err error
	widgetPathCopyFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathCopyFunction, err = widgetPathStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type gtk_widget_path_copy.
func (recv *WidgetPath) Copy() *WidgetPath {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := widgetPathCopyFunction_Set()
	if err == nil {
		ret = widgetPathCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

var widgetPathFreeFunction *gi.Function
var widgetPathFreeFunction_Once sync.Once

func widgetPathFreeFunction_Set() error {
	var err error
	widgetPathFreeFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathFreeFunction, err = widgetPathStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type gtk_widget_path_free.
func (recv *WidgetPath) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := widgetPathFreeFunction_Set()
	if err == nil {
		widgetPathFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_widget_path_get_object_type' : return type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_has_parent' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_is_type' : parameter 'type' of type 'GType' not supported

var widgetPathIterAddClassFunction *gi.Function
var widgetPathIterAddClassFunction_Once sync.Once

func widgetPathIterAddClassFunction_Set() error {
	var err error
	widgetPathIterAddClassFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterAddClassFunction, err = widgetPathStruct.InvokerNew("iter_add_class")
	})
	return err
}

// IterAddClass is a representation of the C type gtk_widget_path_iter_add_class.
func (recv *WidgetPath) IterAddClass(pos int32, name string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	err := widgetPathIterAddClassFunction_Set()
	if err == nil {
		widgetPathIterAddClassFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_widget_path_iter_add_region' : parameter 'flags' of type 'RegionFlags' not supported

var widgetPathIterClearClassesFunction *gi.Function
var widgetPathIterClearClassesFunction_Once sync.Once

func widgetPathIterClearClassesFunction_Set() error {
	var err error
	widgetPathIterClearClassesFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterClearClassesFunction, err = widgetPathStruct.InvokerNew("iter_clear_classes")
	})
	return err
}

// IterClearClasses is a representation of the C type gtk_widget_path_iter_clear_classes.
func (recv *WidgetPath) IterClearClasses(pos int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	err := widgetPathIterClearClassesFunction_Set()
	if err == nil {
		widgetPathIterClearClassesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var widgetPathIterClearRegionsFunction *gi.Function
var widgetPathIterClearRegionsFunction_Once sync.Once

func widgetPathIterClearRegionsFunction_Set() error {
	var err error
	widgetPathIterClearRegionsFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterClearRegionsFunction, err = widgetPathStruct.InvokerNew("iter_clear_regions")
	})
	return err
}

// IterClearRegions is a representation of the C type gtk_widget_path_iter_clear_regions.
func (recv *WidgetPath) IterClearRegions(pos int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	err := widgetPathIterClearRegionsFunction_Set()
	if err == nil {
		widgetPathIterClearRegionsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var widgetPathIterGetNameFunction *gi.Function
var widgetPathIterGetNameFunction_Once sync.Once

func widgetPathIterGetNameFunction_Set() error {
	var err error
	widgetPathIterGetNameFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterGetNameFunction, err = widgetPathStruct.InvokerNew("iter_get_name")
	})
	return err
}

// IterGetName is a representation of the C type gtk_widget_path_iter_get_name.
func (recv *WidgetPath) IterGetName(pos int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	var ret gi.Argument

	err := widgetPathIterGetNameFunction_Set()
	if err == nil {
		ret = widgetPathIterGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var widgetPathIterGetObjectNameFunction *gi.Function
var widgetPathIterGetObjectNameFunction_Once sync.Once

func widgetPathIterGetObjectNameFunction_Set() error {
	var err error
	widgetPathIterGetObjectNameFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterGetObjectNameFunction, err = widgetPathStruct.InvokerNew("iter_get_object_name")
	})
	return err
}

// IterGetObjectName is a representation of the C type gtk_widget_path_iter_get_object_name.
func (recv *WidgetPath) IterGetObjectName(pos int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	var ret gi.Argument

	err := widgetPathIterGetObjectNameFunction_Set()
	if err == nil {
		ret = widgetPathIterGetObjectNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_iter_get_object_type' : return type 'GType' not supported

var widgetPathIterGetSiblingIndexFunction *gi.Function
var widgetPathIterGetSiblingIndexFunction_Once sync.Once

func widgetPathIterGetSiblingIndexFunction_Set() error {
	var err error
	widgetPathIterGetSiblingIndexFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterGetSiblingIndexFunction, err = widgetPathStruct.InvokerNew("iter_get_sibling_index")
	})
	return err
}

// IterGetSiblingIndex is a representation of the C type gtk_widget_path_iter_get_sibling_index.
func (recv *WidgetPath) IterGetSiblingIndex(pos int32) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	var ret gi.Argument

	err := widgetPathIterGetSiblingIndexFunction_Set()
	if err == nil {
		ret = widgetPathIterGetSiblingIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var widgetPathIterGetSiblingsFunction *gi.Function
var widgetPathIterGetSiblingsFunction_Once sync.Once

func widgetPathIterGetSiblingsFunction_Set() error {
	var err error
	widgetPathIterGetSiblingsFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterGetSiblingsFunction, err = widgetPathStruct.InvokerNew("iter_get_siblings")
	})
	return err
}

// IterGetSiblings is a representation of the C type gtk_widget_path_iter_get_siblings.
func (recv *WidgetPath) IterGetSiblings(pos int32) *WidgetPath {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)

	var ret gi.Argument

	err := widgetPathIterGetSiblingsFunction_Set()
	if err == nil {
		ret = widgetPathIterGetSiblingsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_iter_get_state' : return type 'StateFlags' not supported

var widgetPathIterHasClassFunction *gi.Function
var widgetPathIterHasClassFunction_Once sync.Once

func widgetPathIterHasClassFunction_Set() error {
	var err error
	widgetPathIterHasClassFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterHasClassFunction, err = widgetPathStruct.InvokerNew("iter_has_class")
	})
	return err
}

// IterHasClass is a representation of the C type gtk_widget_path_iter_has_class.
func (recv *WidgetPath) IterHasClass(pos int32, name string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	var ret gi.Argument

	err := widgetPathIterHasClassFunction_Set()
	if err == nil {
		ret = widgetPathIterHasClassFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var widgetPathIterHasNameFunction *gi.Function
var widgetPathIterHasNameFunction_Once sync.Once

func widgetPathIterHasNameFunction_Set() error {
	var err error
	widgetPathIterHasNameFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterHasNameFunction, err = widgetPathStruct.InvokerNew("iter_has_name")
	})
	return err
}

// IterHasName is a representation of the C type gtk_widget_path_iter_has_name.
func (recv *WidgetPath) IterHasName(pos int32, name string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	var ret gi.Argument

	err := widgetPathIterHasNameFunction_Set()
	if err == nil {
		ret = widgetPathIterHasNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_qclass' : parameter 'qname' of type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_qname' : parameter 'qname' of type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_qregion' : parameter 'qname' of type 'GLib.Quark' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_region' : parameter 'flags' of type 'RegionFlags' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_list_classes' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_list_regions' : return type 'GLib.SList' not supported

var widgetPathIterRemoveClassFunction *gi.Function
var widgetPathIterRemoveClassFunction_Once sync.Once

func widgetPathIterRemoveClassFunction_Set() error {
	var err error
	widgetPathIterRemoveClassFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterRemoveClassFunction, err = widgetPathStruct.InvokerNew("iter_remove_class")
	})
	return err
}

// IterRemoveClass is a representation of the C type gtk_widget_path_iter_remove_class.
func (recv *WidgetPath) IterRemoveClass(pos int32, name string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	err := widgetPathIterRemoveClassFunction_Set()
	if err == nil {
		widgetPathIterRemoveClassFunction.Invoke(inArgs[:], nil)
	}

	return
}

var widgetPathIterRemoveRegionFunction *gi.Function
var widgetPathIterRemoveRegionFunction_Once sync.Once

func widgetPathIterRemoveRegionFunction_Set() error {
	var err error
	widgetPathIterRemoveRegionFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterRemoveRegionFunction, err = widgetPathStruct.InvokerNew("iter_remove_region")
	})
	return err
}

// IterRemoveRegion is a representation of the C type gtk_widget_path_iter_remove_region.
func (recv *WidgetPath) IterRemoveRegion(pos int32, name string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	err := widgetPathIterRemoveRegionFunction_Set()
	if err == nil {
		widgetPathIterRemoveRegionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var widgetPathIterSetNameFunction *gi.Function
var widgetPathIterSetNameFunction_Once sync.Once

func widgetPathIterSetNameFunction_Set() error {
	var err error
	widgetPathIterSetNameFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterSetNameFunction, err = widgetPathStruct.InvokerNew("iter_set_name")
	})
	return err
}

// IterSetName is a representation of the C type gtk_widget_path_iter_set_name.
func (recv *WidgetPath) IterSetName(pos int32, name string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	err := widgetPathIterSetNameFunction_Set()
	if err == nil {
		widgetPathIterSetNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var widgetPathIterSetObjectNameFunction *gi.Function
var widgetPathIterSetObjectNameFunction_Once sync.Once

func widgetPathIterSetObjectNameFunction_Set() error {
	var err error
	widgetPathIterSetObjectNameFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterSetObjectNameFunction, err = widgetPathStruct.InvokerNew("iter_set_object_name")
	})
	return err
}

// IterSetObjectName is a representation of the C type gtk_widget_path_iter_set_object_name.
func (recv *WidgetPath) IterSetObjectName(pos int32, name string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetString(name)

	err := widgetPathIterSetObjectNameFunction_Set()
	if err == nil {
		widgetPathIterSetObjectNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_widget_path_iter_set_object_type' : parameter 'type' of type 'GType' not supported

// UNSUPPORTED : C value 'gtk_widget_path_iter_set_state' : parameter 'state' of type 'StateFlags' not supported

var widgetPathLengthFunction *gi.Function
var widgetPathLengthFunction_Once sync.Once

func widgetPathLengthFunction_Set() error {
	var err error
	widgetPathLengthFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathLengthFunction, err = widgetPathStruct.InvokerNew("length")
	})
	return err
}

// Length is a representation of the C type gtk_widget_path_length.
func (recv *WidgetPath) Length() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := widgetPathLengthFunction_Set()
	if err == nil {
		ret = widgetPathLengthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_prepend_type' : parameter 'type' of type 'GType' not supported

var widgetPathRefFunction *gi.Function
var widgetPathRefFunction_Once sync.Once

func widgetPathRefFunction_Set() error {
	var err error
	widgetPathRefFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathRefFunction, err = widgetPathStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gtk_widget_path_ref.
func (recv *WidgetPath) Ref() *WidgetPath {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := widgetPathRefFunction_Set()
	if err == nil {
		ret = widgetPathRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &WidgetPath{native: ret.Pointer()}

	return retGo
}

var widgetPathToStringFunction *gi.Function
var widgetPathToStringFunction_Once sync.Once

func widgetPathToStringFunction_Set() error {
	var err error
	widgetPathToStringFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathToStringFunction, err = widgetPathStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type gtk_widget_path_to_string.
func (recv *WidgetPath) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := widgetPathToStringFunction_Set()
	if err == nil {
		ret = widgetPathToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var widgetPathUnrefFunction *gi.Function
var widgetPathUnrefFunction_Once sync.Once

func widgetPathUnrefFunction_Set() error {
	var err error
	widgetPathUnrefFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathUnrefFunction, err = widgetPathStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gtk_widget_path_unref.
func (recv *WidgetPath) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := widgetPathUnrefFunction_Set()
	if err == nil {
		widgetPathUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var widgetPrivateStruct *gi.Struct
var widgetPrivateStruct_Once sync.Once

func widgetPrivateStruct_Set() error {
	var err error
	widgetPrivateStruct_Once.Do(func() {
		widgetPrivateStruct, err = gi.StructNew("Gtk", "WidgetPrivate")
	})
	return err
}

type WidgetPrivate struct {
	native uintptr
}

var windowAccessibleClassStruct *gi.Struct
var windowAccessibleClassStruct_Once sync.Once

func windowAccessibleClassStruct_Set() error {
	var err error
	windowAccessibleClassStruct_Once.Do(func() {
		windowAccessibleClassStruct, err = gi.StructNew("Gtk", "WindowAccessibleClass")
	})
	return err
}

type WindowAccessibleClass struct {
	native uintptr
}

var windowAccessiblePrivateStruct *gi.Struct
var windowAccessiblePrivateStruct_Once sync.Once

func windowAccessiblePrivateStruct_Set() error {
	var err error
	windowAccessiblePrivateStruct_Once.Do(func() {
		windowAccessiblePrivateStruct, err = gi.StructNew("Gtk", "WindowAccessiblePrivate")
	})
	return err
}

type WindowAccessiblePrivate struct {
	native uintptr
}

var windowClassStruct *gi.Struct
var windowClassStruct_Once sync.Once

func windowClassStruct_Set() error {
	var err error
	windowClassStruct_Once.Do(func() {
		windowClassStruct, err = gi.StructNew("Gtk", "WindowClass")
	})
	return err
}

type WindowClass struct {
	native uintptr
}

var windowGeometryInfoStruct *gi.Struct
var windowGeometryInfoStruct_Once sync.Once

func windowGeometryInfoStruct_Set() error {
	var err error
	windowGeometryInfoStruct_Once.Do(func() {
		windowGeometryInfoStruct, err = gi.StructNew("Gtk", "WindowGeometryInfo")
	})
	return err
}

type WindowGeometryInfo struct {
	native uintptr
}

var windowGroupClassStruct *gi.Struct
var windowGroupClassStruct_Once sync.Once

func windowGroupClassStruct_Set() error {
	var err error
	windowGroupClassStruct_Once.Do(func() {
		windowGroupClassStruct, err = gi.StructNew("Gtk", "WindowGroupClass")
	})
	return err
}

type WindowGroupClass struct {
	native uintptr
}

var windowGroupPrivateStruct *gi.Struct
var windowGroupPrivateStruct_Once sync.Once

func windowGroupPrivateStruct_Set() error {
	var err error
	windowGroupPrivateStruct_Once.Do(func() {
		windowGroupPrivateStruct, err = gi.StructNew("Gtk", "WindowGroupPrivate")
	})
	return err
}

type WindowGroupPrivate struct {
	native uintptr
}

var windowPrivateStruct *gi.Struct
var windowPrivateStruct_Once sync.Once

func windowPrivateStruct_Set() error {
	var err error
	windowPrivateStruct_Once.Do(func() {
		windowPrivateStruct, err = gi.StructNew("Gtk", "WindowPrivate")
	})
	return err
}

type WindowPrivate struct {
	native uintptr
}
