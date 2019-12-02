// Code generated - DO NOT EDIT.

package gtk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	"runtime"
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *AboutDialogClass) FieldParentClass() *DialogClass {
	argValue := gi.FieldGet(aboutDialogClassStruct, recv.native, "parent_class")
	value := &DialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AboutDialogClass) SetFieldParentClass(value *DialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(aboutDialogClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'activate_link' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_link' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// AboutDialogClassStruct creates an uninitialised AboutDialogClass.
func AboutDialogClassStruct() *AboutDialogClass {
	err := aboutDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AboutDialogClass{native: aboutDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAboutDialogClass)
	return structGo
}
func finalizeAboutDialogClass(obj *AboutDialogClass) {
	aboutDialogClassStruct.Free(obj.native)
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

// AboutDialogPrivateStruct creates an uninitialised AboutDialogPrivate.
func AboutDialogPrivateStruct() *AboutDialogPrivate {
	err := aboutDialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AboutDialogPrivate{native: aboutDialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAboutDialogPrivate)
	return structGo
}
func finalizeAboutDialogPrivate(obj *AboutDialogPrivate) {
	aboutDialogPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'accel_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'accel_changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// AccelGroupClassStruct creates an uninitialised AccelGroupClass.
func AccelGroupClassStruct() *AccelGroupClass {
	err := accelGroupClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AccelGroupClass{native: accelGroupClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAccelGroupClass)
	return structGo
}
func finalizeAccelGroupClass(obj *AccelGroupClass) {
	accelGroupClassStruct.Free(obj.native)
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

// FieldKey returns the C field 'key'.
func (recv *AccelGroupEntry) FieldKey() *AccelKey {
	argValue := gi.FieldGet(accelGroupEntryStruct, recv.native, "key")
	value := &AccelKey{native: argValue.Pointer()}
	return value
}

// SetFieldKey sets the value of the C field 'key'.
func (recv *AccelGroupEntry) SetFieldKey(value *AccelKey) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(accelGroupEntryStruct, recv.native, "key", argValue)
}

// UNSUPPORTED : C value 'closure' : for field getter : no Go type for 'GObject.Closure'

// UNSUPPORTED : C value 'closure' : for field setter : no Go type for 'GObject.Closure'

// FieldAccelPathQuark returns the C field 'accel_path_quark'.
func (recv *AccelGroupEntry) FieldAccelPathQuark() glib.Quark {
	argValue := gi.FieldGet(accelGroupEntryStruct, recv.native, "accel_path_quark")
	value := glib.Quark(argValue.Uint32())
	return value
}

// SetFieldAccelPathQuark sets the value of the C field 'accel_path_quark'.
func (recv *AccelGroupEntry) SetFieldAccelPathQuark(value glib.Quark) {
	var argValue gi.Argument
	argValue.SetUint32(uint32(value))
	gi.FieldSet(accelGroupEntryStruct, recv.native, "accel_path_quark", argValue)
}

// AccelGroupEntryStruct creates an uninitialised AccelGroupEntry.
func AccelGroupEntryStruct() *AccelGroupEntry {
	err := accelGroupEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AccelGroupEntry{native: accelGroupEntryStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAccelGroupEntry)
	return structGo
}
func finalizeAccelGroupEntry(obj *AccelGroupEntry) {
	accelGroupEntryStruct.Free(obj.native)
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

// AccelGroupPrivateStruct creates an uninitialised AccelGroupPrivate.
func AccelGroupPrivateStruct() *AccelGroupPrivate {
	err := accelGroupPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AccelGroupPrivate{native: accelGroupPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAccelGroupPrivate)
	return structGo
}
func finalizeAccelGroupPrivate(obj *AccelGroupPrivate) {
	accelGroupPrivateStruct.Free(obj.native)
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

// FieldAccelKey returns the C field 'accel_key'.
func (recv *AccelKey) FieldAccelKey() uint32 {
	argValue := gi.FieldGet(accelKeyStruct, recv.native, "accel_key")
	value := argValue.Uint32()
	return value
}

// SetFieldAccelKey sets the value of the C field 'accel_key'.
func (recv *AccelKey) SetFieldAccelKey(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(accelKeyStruct, recv.native, "accel_key", argValue)
}

// UNSUPPORTED : C value 'accel_mods' : for field getter : no Go type for 'Gdk.ModifierType'

// UNSUPPORTED : C value 'accel_mods' : for field setter : no Go type for 'Gdk.ModifierType'

// FieldAccelFlags returns the C field 'accel_flags'.
func (recv *AccelKey) FieldAccelFlags() uint32 {
	argValue := gi.FieldGet(accelKeyStruct, recv.native, "accel_flags")
	value := argValue.Uint32()
	return value
}

// SetFieldAccelFlags sets the value of the C field 'accel_flags'.
func (recv *AccelKey) SetFieldAccelFlags(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(accelKeyStruct, recv.native, "accel_flags", argValue)
}

// AccelKeyStruct creates an uninitialised AccelKey.
func AccelKeyStruct() *AccelKey {
	err := accelKeyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AccelKey{native: accelKeyStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAccelKey)
	return structGo
}
func finalizeAccelKey(obj *AccelKey) {
	accelKeyStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *AccelLabelClass) FieldParentClass() *LabelClass {
	argValue := gi.FieldGet(accelLabelClassStruct, recv.native, "parent_class")
	value := &LabelClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AccelLabelClass) SetFieldParentClass(value *LabelClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(accelLabelClassStruct, recv.native, "parent_class", argValue)
}

// FieldSignalQuote1 returns the C field 'signal_quote1'.
func (recv *AccelLabelClass) FieldSignalQuote1() string {
	argValue := gi.FieldGet(accelLabelClassStruct, recv.native, "signal_quote1")
	value := argValue.String(false)
	return value
}

// SetFieldSignalQuote1 sets the value of the C field 'signal_quote1'.
func (recv *AccelLabelClass) SetFieldSignalQuote1(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(accelLabelClassStruct, recv.native, "signal_quote1", argValue)
}

// FieldSignalQuote2 returns the C field 'signal_quote2'.
func (recv *AccelLabelClass) FieldSignalQuote2() string {
	argValue := gi.FieldGet(accelLabelClassStruct, recv.native, "signal_quote2")
	value := argValue.String(false)
	return value
}

// SetFieldSignalQuote2 sets the value of the C field 'signal_quote2'.
func (recv *AccelLabelClass) SetFieldSignalQuote2(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(accelLabelClassStruct, recv.native, "signal_quote2", argValue)
}

// FieldModNameShift returns the C field 'mod_name_shift'.
func (recv *AccelLabelClass) FieldModNameShift() string {
	argValue := gi.FieldGet(accelLabelClassStruct, recv.native, "mod_name_shift")
	value := argValue.String(false)
	return value
}

// SetFieldModNameShift sets the value of the C field 'mod_name_shift'.
func (recv *AccelLabelClass) SetFieldModNameShift(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(accelLabelClassStruct, recv.native, "mod_name_shift", argValue)
}

// FieldModNameControl returns the C field 'mod_name_control'.
func (recv *AccelLabelClass) FieldModNameControl() string {
	argValue := gi.FieldGet(accelLabelClassStruct, recv.native, "mod_name_control")
	value := argValue.String(false)
	return value
}

// SetFieldModNameControl sets the value of the C field 'mod_name_control'.
func (recv *AccelLabelClass) SetFieldModNameControl(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(accelLabelClassStruct, recv.native, "mod_name_control", argValue)
}

// FieldModNameAlt returns the C field 'mod_name_alt'.
func (recv *AccelLabelClass) FieldModNameAlt() string {
	argValue := gi.FieldGet(accelLabelClassStruct, recv.native, "mod_name_alt")
	value := argValue.String(false)
	return value
}

// SetFieldModNameAlt sets the value of the C field 'mod_name_alt'.
func (recv *AccelLabelClass) SetFieldModNameAlt(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(accelLabelClassStruct, recv.native, "mod_name_alt", argValue)
}

// FieldModSeparator returns the C field 'mod_separator'.
func (recv *AccelLabelClass) FieldModSeparator() string {
	argValue := gi.FieldGet(accelLabelClassStruct, recv.native, "mod_separator")
	value := argValue.String(false)
	return value
}

// SetFieldModSeparator sets the value of the C field 'mod_separator'.
func (recv *AccelLabelClass) SetFieldModSeparator(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(accelLabelClassStruct, recv.native, "mod_separator", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// AccelLabelClassStruct creates an uninitialised AccelLabelClass.
func AccelLabelClassStruct() *AccelLabelClass {
	err := accelLabelClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AccelLabelClass{native: accelLabelClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAccelLabelClass)
	return structGo
}
func finalizeAccelLabelClass(obj *AccelLabelClass) {
	accelLabelClassStruct.Free(obj.native)
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

// AccelLabelPrivateStruct creates an uninitialised AccelLabelPrivate.
func AccelLabelPrivateStruct() *AccelLabelPrivate {
	err := accelLabelPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AccelLabelPrivate{native: accelLabelPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAccelLabelPrivate)
	return structGo
}
func finalizeAccelLabelPrivate(obj *AccelLabelPrivate) {
	accelLabelPrivateStruct.Free(obj.native)
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

// AccelMapClassStruct creates an uninitialised AccelMapClass.
func AccelMapClassStruct() *AccelMapClass {
	err := accelMapClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AccelMapClass{native: accelMapClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAccelMapClass)
	return structGo
}
func finalizeAccelMapClass(obj *AccelMapClass) {
	accelMapClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Atk.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Atk.ObjectClass'

// UNSUPPORTED : C value 'connect_widget_destroyed' : for field getter : missing Type

// UNSUPPORTED : C value 'connect_widget_destroyed' : for field setter : missing Type

// UNSUPPORTED : C value 'widget_set' : for field getter : missing Type

// UNSUPPORTED : C value 'widget_set' : for field setter : missing Type

// UNSUPPORTED : C value 'widget_unset' : for field getter : missing Type

// UNSUPPORTED : C value 'widget_unset' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// AccessibleClassStruct creates an uninitialised AccessibleClass.
func AccessibleClassStruct() *AccessibleClass {
	err := accessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AccessibleClass{native: accessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAccessibleClass)
	return structGo
}
func finalizeAccessibleClass(obj *AccessibleClass) {
	accessibleClassStruct.Free(obj.native)
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

// AccessiblePrivateStruct creates an uninitialised AccessiblePrivate.
func AccessiblePrivateStruct() *AccessiblePrivate {
	err := accessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AccessiblePrivate{native: accessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAccessiblePrivate)
	return structGo
}
func finalizeAccessiblePrivate(obj *AccessiblePrivate) {
	accessiblePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ActionBarClassStruct creates an uninitialised ActionBarClass.
func ActionBarClassStruct() *ActionBarClass {
	err := actionBarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionBarClass{native: actionBarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActionBarClass)
	return structGo
}
func finalizeActionBarClass(obj *ActionBarClass) {
	actionBarClassStruct.Free(obj.native)
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

// ActionBarPrivateStruct creates an uninitialised ActionBarPrivate.
func ActionBarPrivateStruct() *ActionBarPrivate {
	err := actionBarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionBarPrivate{native: actionBarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActionBarPrivate)
	return structGo
}
func finalizeActionBarPrivate(obj *ActionBarPrivate) {
	actionBarPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value 'create_menu_item' : for field getter : missing Type

// UNSUPPORTED : C value 'create_menu_item' : for field setter : missing Type

// UNSUPPORTED : C value 'create_tool_item' : for field getter : missing Type

// UNSUPPORTED : C value 'create_tool_item' : for field setter : missing Type

// UNSUPPORTED : C value 'connect_proxy' : for field getter : missing Type

// UNSUPPORTED : C value 'connect_proxy' : for field setter : missing Type

// UNSUPPORTED : C value 'disconnect_proxy' : for field getter : missing Type

// UNSUPPORTED : C value 'disconnect_proxy' : for field setter : missing Type

// UNSUPPORTED : C value 'create_menu' : for field getter : missing Type

// UNSUPPORTED : C value 'create_menu' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ActionClassStruct creates an uninitialised ActionClass.
func ActionClassStruct() *ActionClass {
	err := actionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionClass{native: actionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActionClass)
	return structGo
}
func finalizeActionClass(obj *ActionClass) {
	actionClassStruct.Free(obj.native)
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

// FieldName returns the C field 'name'.
func (recv *ActionEntry) FieldName() string {
	argValue := gi.FieldGet(actionEntryStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *ActionEntry) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(actionEntryStruct, recv.native, "name", argValue)
}

// FieldStockId returns the C field 'stock_id'.
func (recv *ActionEntry) FieldStockId() string {
	argValue := gi.FieldGet(actionEntryStruct, recv.native, "stock_id")
	value := argValue.String(false)
	return value
}

// SetFieldStockId sets the value of the C field 'stock_id'.
func (recv *ActionEntry) SetFieldStockId(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(actionEntryStruct, recv.native, "stock_id", argValue)
}

// FieldLabel returns the C field 'label'.
func (recv *ActionEntry) FieldLabel() string {
	argValue := gi.FieldGet(actionEntryStruct, recv.native, "label")
	value := argValue.String(false)
	return value
}

// SetFieldLabel sets the value of the C field 'label'.
func (recv *ActionEntry) SetFieldLabel(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(actionEntryStruct, recv.native, "label", argValue)
}

// FieldAccelerator returns the C field 'accelerator'.
func (recv *ActionEntry) FieldAccelerator() string {
	argValue := gi.FieldGet(actionEntryStruct, recv.native, "accelerator")
	value := argValue.String(false)
	return value
}

// SetFieldAccelerator sets the value of the C field 'accelerator'.
func (recv *ActionEntry) SetFieldAccelerator(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(actionEntryStruct, recv.native, "accelerator", argValue)
}

// FieldTooltip returns the C field 'tooltip'.
func (recv *ActionEntry) FieldTooltip() string {
	argValue := gi.FieldGet(actionEntryStruct, recv.native, "tooltip")
	value := argValue.String(false)
	return value
}

// SetFieldTooltip sets the value of the C field 'tooltip'.
func (recv *ActionEntry) SetFieldTooltip(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(actionEntryStruct, recv.native, "tooltip", argValue)
}

// UNSUPPORTED : C value 'callback' : for field getter : no Go type for 'GObject.Callback'

// UNSUPPORTED : C value 'callback' : for field setter : no Go type for 'GObject.Callback'

// ActionEntryStruct creates an uninitialised ActionEntry.
func ActionEntryStruct() *ActionEntry {
	err := actionEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionEntry{native: actionEntryStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActionEntry)
	return structGo
}
func finalizeActionEntry(obj *ActionEntry) {
	actionEntryStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_action' : for field getter : missing Type

// UNSUPPORTED : C value 'get_action' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ActionGroupClassStruct creates an uninitialised ActionGroupClass.
func ActionGroupClassStruct() *ActionGroupClass {
	err := actionGroupClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionGroupClass{native: actionGroupClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActionGroupClass)
	return structGo
}
func finalizeActionGroupClass(obj *ActionGroupClass) {
	actionGroupClassStruct.Free(obj.native)
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

// ActionGroupPrivateStruct creates an uninitialised ActionGroupPrivate.
func ActionGroupPrivateStruct() *ActionGroupPrivate {
	err := actionGroupPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionGroupPrivate{native: actionGroupPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActionGroupPrivate)
	return structGo
}
func finalizeActionGroupPrivate(obj *ActionGroupPrivate) {
	actionGroupPrivateStruct.Free(obj.native)
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

// ActionPrivateStruct creates an uninitialised ActionPrivate.
func ActionPrivateStruct() *ActionPrivate {
	err := actionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionPrivate{native: actionPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActionPrivate)
	return structGo
}
func finalizeActionPrivate(obj *ActionPrivate) {
	actionPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'get_action_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_action_name' : for field setter : missing Type

// UNSUPPORTED : C value 'set_action_name' : for field getter : missing Type

// UNSUPPORTED : C value 'set_action_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_action_target_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_action_target_value' : for field setter : missing Type

// UNSUPPORTED : C value 'set_action_target_value' : for field getter : missing Type

// UNSUPPORTED : C value 'set_action_target_value' : for field setter : missing Type

// ActionableInterfaceStruct creates an uninitialised ActionableInterface.
func ActionableInterfaceStruct() *ActionableInterface {
	err := actionableInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActionableInterface{native: actionableInterfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActionableInterface)
	return structGo
}
func finalizeActionableInterface(obj *ActionableInterface) {
	actionableInterfaceStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'update' : for field getter : missing Type

// UNSUPPORTED : C value 'update' : for field setter : missing Type

// UNSUPPORTED : C value 'sync_action_properties' : for field getter : missing Type

// UNSUPPORTED : C value 'sync_action_properties' : for field setter : missing Type

// ActivatableIfaceStruct creates an uninitialised ActivatableIface.
func ActivatableIfaceStruct() *ActivatableIface {
	err := activatableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ActivatableIface{native: activatableIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeActivatableIface)
	return structGo
}
func finalizeActivatableIface(obj *ActivatableIface) {
	activatableIfaceStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'value_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'value_changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// AdjustmentClassStruct creates an uninitialised AdjustmentClass.
func AdjustmentClassStruct() *AdjustmentClass {
	err := adjustmentClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AdjustmentClass{native: adjustmentClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAdjustmentClass)
	return structGo
}
func finalizeAdjustmentClass(obj *AdjustmentClass) {
	adjustmentClassStruct.Free(obj.native)
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

// AdjustmentPrivateStruct creates an uninitialised AdjustmentPrivate.
func AdjustmentPrivateStruct() *AdjustmentPrivate {
	err := adjustmentPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AdjustmentPrivate{native: adjustmentPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAdjustmentPrivate)
	return structGo
}
func finalizeAdjustmentPrivate(obj *AdjustmentPrivate) {
	adjustmentPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *AlignmentClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(alignmentClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AlignmentClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(alignmentClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// AlignmentClassStruct creates an uninitialised AlignmentClass.
func AlignmentClassStruct() *AlignmentClass {
	err := alignmentClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AlignmentClass{native: alignmentClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAlignmentClass)
	return structGo
}
func finalizeAlignmentClass(obj *AlignmentClass) {
	alignmentClassStruct.Free(obj.native)
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

// AlignmentPrivateStruct creates an uninitialised AlignmentPrivate.
func AlignmentPrivateStruct() *AlignmentPrivate {
	err := alignmentPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AlignmentPrivate{native: alignmentPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAlignmentPrivate)
	return structGo
}
func finalizeAlignmentPrivate(obj *AlignmentPrivate) {
	alignmentPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *AppChooserButtonClass) FieldParentClass() *ComboBoxClass {
	argValue := gi.FieldGet(appChooserButtonClassStruct, recv.native, "parent_class")
	value := &ComboBoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AppChooserButtonClass) SetFieldParentClass(value *ComboBoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(appChooserButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'custom_item_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'custom_item_activated' : for field setter : missing Type

// AppChooserButtonClassStruct creates an uninitialised AppChooserButtonClass.
func AppChooserButtonClassStruct() *AppChooserButtonClass {
	err := appChooserButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AppChooserButtonClass{native: appChooserButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAppChooserButtonClass)
	return structGo
}
func finalizeAppChooserButtonClass(obj *AppChooserButtonClass) {
	appChooserButtonClassStruct.Free(obj.native)
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

// AppChooserButtonPrivateStruct creates an uninitialised AppChooserButtonPrivate.
func AppChooserButtonPrivateStruct() *AppChooserButtonPrivate {
	err := appChooserButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AppChooserButtonPrivate{native: appChooserButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAppChooserButtonPrivate)
	return structGo
}
func finalizeAppChooserButtonPrivate(obj *AppChooserButtonPrivate) {
	appChooserButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *AppChooserDialogClass) FieldParentClass() *DialogClass {
	argValue := gi.FieldGet(appChooserDialogClassStruct, recv.native, "parent_class")
	value := &DialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AppChooserDialogClass) SetFieldParentClass(value *DialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(appChooserDialogClassStruct, recv.native, "parent_class", argValue)
}

// AppChooserDialogClassStruct creates an uninitialised AppChooserDialogClass.
func AppChooserDialogClassStruct() *AppChooserDialogClass {
	err := appChooserDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AppChooserDialogClass{native: appChooserDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAppChooserDialogClass)
	return structGo
}
func finalizeAppChooserDialogClass(obj *AppChooserDialogClass) {
	appChooserDialogClassStruct.Free(obj.native)
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

// AppChooserDialogPrivateStruct creates an uninitialised AppChooserDialogPrivate.
func AppChooserDialogPrivateStruct() *AppChooserDialogPrivate {
	err := appChooserDialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AppChooserDialogPrivate{native: appChooserDialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAppChooserDialogPrivate)
	return structGo
}
func finalizeAppChooserDialogPrivate(obj *AppChooserDialogPrivate) {
	appChooserDialogPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *AppChooserWidgetClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(appChooserWidgetClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AppChooserWidgetClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(appChooserWidgetClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'application_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'application_selected' : for field setter : missing Type

// UNSUPPORTED : C value 'application_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'application_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'populate_popup' : for field getter : missing Type

// UNSUPPORTED : C value 'populate_popup' : for field setter : missing Type

// AppChooserWidgetClassStruct creates an uninitialised AppChooserWidgetClass.
func AppChooserWidgetClassStruct() *AppChooserWidgetClass {
	err := appChooserWidgetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AppChooserWidgetClass{native: appChooserWidgetClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAppChooserWidgetClass)
	return structGo
}
func finalizeAppChooserWidgetClass(obj *AppChooserWidgetClass) {
	appChooserWidgetClassStruct.Free(obj.native)
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

// AppChooserWidgetPrivateStruct creates an uninitialised AppChooserWidgetPrivate.
func AppChooserWidgetPrivateStruct() *AppChooserWidgetPrivate {
	err := appChooserWidgetPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AppChooserWidgetPrivate{native: appChooserWidgetPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAppChooserWidgetPrivate)
	return structGo
}
func finalizeAppChooserWidgetPrivate(obj *AppChooserWidgetPrivate) {
	appChooserWidgetPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gio.ApplicationClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Gio.ApplicationClass'

// UNSUPPORTED : C value 'window_added' : for field getter : missing Type

// UNSUPPORTED : C value 'window_added' : for field setter : missing Type

// UNSUPPORTED : C value 'window_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'window_removed' : for field setter : missing Type

// ApplicationClassStruct creates an uninitialised ApplicationClass.
func ApplicationClassStruct() *ApplicationClass {
	err := applicationClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ApplicationClass{native: applicationClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeApplicationClass)
	return structGo
}
func finalizeApplicationClass(obj *ApplicationClass) {
	applicationClassStruct.Free(obj.native)
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

// ApplicationPrivateStruct creates an uninitialised ApplicationPrivate.
func ApplicationPrivateStruct() *ApplicationPrivate {
	err := applicationPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ApplicationPrivate{native: applicationPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeApplicationPrivate)
	return structGo
}
func finalizeApplicationPrivate(obj *ApplicationPrivate) {
	applicationPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ApplicationWindowClass) FieldParentClass() *WindowClass {
	argValue := gi.FieldGet(applicationWindowClassStruct, recv.native, "parent_class")
	value := &WindowClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ApplicationWindowClass) SetFieldParentClass(value *WindowClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(applicationWindowClassStruct, recv.native, "parent_class", argValue)
}

// ApplicationWindowClassStruct creates an uninitialised ApplicationWindowClass.
func ApplicationWindowClassStruct() *ApplicationWindowClass {
	err := applicationWindowClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ApplicationWindowClass{native: applicationWindowClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeApplicationWindowClass)
	return structGo
}
func finalizeApplicationWindowClass(obj *ApplicationWindowClass) {
	applicationWindowClassStruct.Free(obj.native)
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

// ApplicationWindowPrivateStruct creates an uninitialised ApplicationWindowPrivate.
func ApplicationWindowPrivateStruct() *ApplicationWindowPrivate {
	err := applicationWindowPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ApplicationWindowPrivate{native: applicationWindowPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeApplicationWindowPrivate)
	return structGo
}
func finalizeApplicationWindowPrivate(obj *ApplicationWindowPrivate) {
	applicationWindowPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ArrowAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(arrowAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ArrowAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(arrowAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ArrowAccessibleClassStruct creates an uninitialised ArrowAccessibleClass.
func ArrowAccessibleClassStruct() *ArrowAccessibleClass {
	err := arrowAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ArrowAccessibleClass{native: arrowAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeArrowAccessibleClass)
	return structGo
}
func finalizeArrowAccessibleClass(obj *ArrowAccessibleClass) {
	arrowAccessibleClassStruct.Free(obj.native)
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

// ArrowAccessiblePrivateStruct creates an uninitialised ArrowAccessiblePrivate.
func ArrowAccessiblePrivateStruct() *ArrowAccessiblePrivate {
	err := arrowAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ArrowAccessiblePrivate{native: arrowAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeArrowAccessiblePrivate)
	return structGo
}
func finalizeArrowAccessiblePrivate(obj *ArrowAccessiblePrivate) {
	arrowAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ArrowClass) FieldParentClass() *MiscClass {
	argValue := gi.FieldGet(arrowClassStruct, recv.native, "parent_class")
	value := &MiscClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ArrowClass) SetFieldParentClass(value *MiscClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(arrowClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ArrowClassStruct creates an uninitialised ArrowClass.
func ArrowClassStruct() *ArrowClass {
	err := arrowClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ArrowClass{native: arrowClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeArrowClass)
	return structGo
}
func finalizeArrowClass(obj *ArrowClass) {
	arrowClassStruct.Free(obj.native)
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

// ArrowPrivateStruct creates an uninitialised ArrowPrivate.
func ArrowPrivateStruct() *ArrowPrivate {
	err := arrowPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ArrowPrivate{native: arrowPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeArrowPrivate)
	return structGo
}
func finalizeArrowPrivate(obj *ArrowPrivate) {
	arrowPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *AspectFrameClass) FieldParentClass() *FrameClass {
	argValue := gi.FieldGet(aspectFrameClassStruct, recv.native, "parent_class")
	value := &FrameClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AspectFrameClass) SetFieldParentClass(value *FrameClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(aspectFrameClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// AspectFrameClassStruct creates an uninitialised AspectFrameClass.
func AspectFrameClassStruct() *AspectFrameClass {
	err := aspectFrameClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AspectFrameClass{native: aspectFrameClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAspectFrameClass)
	return structGo
}
func finalizeAspectFrameClass(obj *AspectFrameClass) {
	aspectFrameClassStruct.Free(obj.native)
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

// AspectFramePrivateStruct creates an uninitialised AspectFramePrivate.
func AspectFramePrivateStruct() *AspectFramePrivate {
	err := aspectFramePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AspectFramePrivate{native: aspectFramePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAspectFramePrivate)
	return structGo
}
func finalizeAspectFramePrivate(obj *AspectFramePrivate) {
	aspectFramePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *AssistantClass) FieldParentClass() *WindowClass {
	argValue := gi.FieldGet(assistantClassStruct, recv.native, "parent_class")
	value := &WindowClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *AssistantClass) SetFieldParentClass(value *WindowClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(assistantClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'prepare' : for field getter : missing Type

// UNSUPPORTED : C value 'prepare' : for field setter : missing Type

// UNSUPPORTED : C value 'apply' : for field getter : missing Type

// UNSUPPORTED : C value 'apply' : for field setter : missing Type

// UNSUPPORTED : C value 'close' : for field getter : missing Type

// UNSUPPORTED : C value 'close' : for field setter : missing Type

// UNSUPPORTED : C value 'cancel' : for field getter : missing Type

// UNSUPPORTED : C value 'cancel' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// AssistantClassStruct creates an uninitialised AssistantClass.
func AssistantClassStruct() *AssistantClass {
	err := assistantClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AssistantClass{native: assistantClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAssistantClass)
	return structGo
}
func finalizeAssistantClass(obj *AssistantClass) {
	assistantClassStruct.Free(obj.native)
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

// AssistantPrivateStruct creates an uninitialised AssistantPrivate.
func AssistantPrivateStruct() *AssistantPrivate {
	err := assistantPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &AssistantPrivate{native: assistantPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeAssistantPrivate)
	return structGo
}
func finalizeAssistantPrivate(obj *AssistantPrivate) {
	assistantPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *BinClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(binClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *BinClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(binClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// BinClassStruct creates an uninitialised BinClass.
func BinClassStruct() *BinClass {
	err := binClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BinClass{native: binClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBinClass)
	return structGo
}
func finalizeBinClass(obj *BinClass) {
	binClassStruct.Free(obj.native)
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

// BinPrivateStruct creates an uninitialised BinPrivate.
func BinPrivateStruct() *BinPrivate {
	err := binPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BinPrivate{native: binPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBinPrivate)
	return structGo
}
func finalizeBinPrivate(obj *BinPrivate) {
	binPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'arg_type' : for field getter : no Go type for 'GType'

// UNSUPPORTED : C value 'arg_type' : for field setter : no Go type for 'GType'

// BindingArgStruct creates an uninitialised BindingArg.
func BindingArgStruct() *BindingArg {
	err := bindingArgStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BindingArg{native: bindingArgStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBindingArg)
	return structGo
}
func finalizeBindingArg(obj *BindingArg) {
	bindingArgStruct.Free(obj.native)
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

// FieldKeyval returns the C field 'keyval'.
func (recv *BindingEntry) FieldKeyval() uint32 {
	argValue := gi.FieldGet(bindingEntryStruct, recv.native, "keyval")
	value := argValue.Uint32()
	return value
}

// SetFieldKeyval sets the value of the C field 'keyval'.
func (recv *BindingEntry) SetFieldKeyval(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(bindingEntryStruct, recv.native, "keyval", argValue)
}

// UNSUPPORTED : C value 'modifiers' : for field getter : no Go type for 'Gdk.ModifierType'

// UNSUPPORTED : C value 'modifiers' : for field setter : no Go type for 'Gdk.ModifierType'

// FieldBindingSet returns the C field 'binding_set'.
func (recv *BindingEntry) FieldBindingSet() *BindingSet {
	argValue := gi.FieldGet(bindingEntryStruct, recv.native, "binding_set")
	value := &BindingSet{native: argValue.Pointer()}
	return value
}

// SetFieldBindingSet sets the value of the C field 'binding_set'.
func (recv *BindingEntry) SetFieldBindingSet(value *BindingSet) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(bindingEntryStruct, recv.native, "binding_set", argValue)
}

// FieldDestroyed returns the C field 'destroyed'.
func (recv *BindingEntry) FieldDestroyed() uint32 {
	argValue := gi.FieldGet(bindingEntryStruct, recv.native, "destroyed")
	value := argValue.Uint32()
	return value
}

// SetFieldDestroyed sets the value of the C field 'destroyed'.
func (recv *BindingEntry) SetFieldDestroyed(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(bindingEntryStruct, recv.native, "destroyed", argValue)
}

// FieldInEmission returns the C field 'in_emission'.
func (recv *BindingEntry) FieldInEmission() uint32 {
	argValue := gi.FieldGet(bindingEntryStruct, recv.native, "in_emission")
	value := argValue.Uint32()
	return value
}

// SetFieldInEmission sets the value of the C field 'in_emission'.
func (recv *BindingEntry) SetFieldInEmission(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(bindingEntryStruct, recv.native, "in_emission", argValue)
}

// FieldMarksUnbound returns the C field 'marks_unbound'.
func (recv *BindingEntry) FieldMarksUnbound() uint32 {
	argValue := gi.FieldGet(bindingEntryStruct, recv.native, "marks_unbound")
	value := argValue.Uint32()
	return value
}

// SetFieldMarksUnbound sets the value of the C field 'marks_unbound'.
func (recv *BindingEntry) SetFieldMarksUnbound(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(bindingEntryStruct, recv.native, "marks_unbound", argValue)
}

// FieldSetNext returns the C field 'set_next'.
func (recv *BindingEntry) FieldSetNext() *BindingEntry {
	argValue := gi.FieldGet(bindingEntryStruct, recv.native, "set_next")
	value := &BindingEntry{native: argValue.Pointer()}
	return value
}

// SetFieldSetNext sets the value of the C field 'set_next'.
func (recv *BindingEntry) SetFieldSetNext(value *BindingEntry) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(bindingEntryStruct, recv.native, "set_next", argValue)
}

// FieldHashNext returns the C field 'hash_next'.
func (recv *BindingEntry) FieldHashNext() *BindingEntry {
	argValue := gi.FieldGet(bindingEntryStruct, recv.native, "hash_next")
	value := &BindingEntry{native: argValue.Pointer()}
	return value
}

// SetFieldHashNext sets the value of the C field 'hash_next'.
func (recv *BindingEntry) SetFieldHashNext(value *BindingEntry) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(bindingEntryStruct, recv.native, "hash_next", argValue)
}

// FieldSignals returns the C field 'signals'.
func (recv *BindingEntry) FieldSignals() *BindingSignal {
	argValue := gi.FieldGet(bindingEntryStruct, recv.native, "signals")
	value := &BindingSignal{native: argValue.Pointer()}
	return value
}

// SetFieldSignals sets the value of the C field 'signals'.
func (recv *BindingEntry) SetFieldSignals(value *BindingSignal) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(bindingEntryStruct, recv.native, "signals", argValue)
}

// BindingEntryStruct creates an uninitialised BindingEntry.
func BindingEntryStruct() *BindingEntry {
	err := bindingEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BindingEntry{native: bindingEntryStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBindingEntry)
	return structGo
}
func finalizeBindingEntry(obj *BindingEntry) {
	bindingEntryStruct.Free(obj.native)
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

// FieldSetName returns the C field 'set_name'.
func (recv *BindingSet) FieldSetName() string {
	argValue := gi.FieldGet(bindingSetStruct, recv.native, "set_name")
	value := argValue.String(false)
	return value
}

// SetFieldSetName sets the value of the C field 'set_name'.
func (recv *BindingSet) SetFieldSetName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(bindingSetStruct, recv.native, "set_name", argValue)
}

// FieldPriority returns the C field 'priority'.
func (recv *BindingSet) FieldPriority() int32 {
	argValue := gi.FieldGet(bindingSetStruct, recv.native, "priority")
	value := argValue.Int32()
	return value
}

// SetFieldPriority sets the value of the C field 'priority'.
func (recv *BindingSet) SetFieldPriority(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(bindingSetStruct, recv.native, "priority", argValue)
}

// UNSUPPORTED : C value 'widget_path_pspecs' : for field getter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'widget_path_pspecs' : for field setter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'widget_class_pspecs' : for field getter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'widget_class_pspecs' : for field setter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'class_branch_pspecs' : for field getter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'class_branch_pspecs' : for field setter : no Go type for 'GLib.SList'

// FieldEntries returns the C field 'entries'.
func (recv *BindingSet) FieldEntries() *BindingEntry {
	argValue := gi.FieldGet(bindingSetStruct, recv.native, "entries")
	value := &BindingEntry{native: argValue.Pointer()}
	return value
}

// SetFieldEntries sets the value of the C field 'entries'.
func (recv *BindingSet) SetFieldEntries(value *BindingEntry) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(bindingSetStruct, recv.native, "entries", argValue)
}

// FieldCurrent returns the C field 'current'.
func (recv *BindingSet) FieldCurrent() *BindingEntry {
	argValue := gi.FieldGet(bindingSetStruct, recv.native, "current")
	value := &BindingEntry{native: argValue.Pointer()}
	return value
}

// SetFieldCurrent sets the value of the C field 'current'.
func (recv *BindingSet) SetFieldCurrent(value *BindingEntry) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(bindingSetStruct, recv.native, "current", argValue)
}

// FieldParsed returns the C field 'parsed'.
func (recv *BindingSet) FieldParsed() uint32 {
	argValue := gi.FieldGet(bindingSetStruct, recv.native, "parsed")
	value := argValue.Uint32()
	return value
}

// SetFieldParsed sets the value of the C field 'parsed'.
func (recv *BindingSet) SetFieldParsed(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(bindingSetStruct, recv.native, "parsed", argValue)
}

// UNSUPPORTED : C value 'gtk_binding_set_activate' : parameter 'modifiers' of type 'Gdk.ModifierType' not supported

// UNSUPPORTED : C value 'gtk_binding_set_add_path' : parameter 'path_type' of type 'PathType' not supported

// BindingSetStruct creates an uninitialised BindingSet.
func BindingSetStruct() *BindingSet {
	err := bindingSetStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BindingSet{native: bindingSetStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBindingSet)
	return structGo
}
func finalizeBindingSet(obj *BindingSet) {
	bindingSetStruct.Free(obj.native)
}

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

// FieldNext returns the C field 'next'.
func (recv *BindingSignal) FieldNext() *BindingSignal {
	argValue := gi.FieldGet(bindingSignalStruct, recv.native, "next")
	value := &BindingSignal{native: argValue.Pointer()}
	return value
}

// SetFieldNext sets the value of the C field 'next'.
func (recv *BindingSignal) SetFieldNext(value *BindingSignal) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(bindingSignalStruct, recv.native, "next", argValue)
}

// FieldSignalName returns the C field 'signal_name'.
func (recv *BindingSignal) FieldSignalName() string {
	argValue := gi.FieldGet(bindingSignalStruct, recv.native, "signal_name")
	value := argValue.String(false)
	return value
}

// SetFieldSignalName sets the value of the C field 'signal_name'.
func (recv *BindingSignal) SetFieldSignalName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(bindingSignalStruct, recv.native, "signal_name", argValue)
}

// FieldNArgs returns the C field 'n_args'.
func (recv *BindingSignal) FieldNArgs() uint32 {
	argValue := gi.FieldGet(bindingSignalStruct, recv.native, "n_args")
	value := argValue.Uint32()
	return value
}

// SetFieldNArgs sets the value of the C field 'n_args'.
func (recv *BindingSignal) SetFieldNArgs(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(bindingSignalStruct, recv.native, "n_args", argValue)
}

// UNSUPPORTED : C value 'args' : for field getter : missing Type

// UNSUPPORTED : C value 'args' : for field setter : missing Type

// BindingSignalStruct creates an uninitialised BindingSignal.
func BindingSignalStruct() *BindingSignal {
	err := bindingSignalStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BindingSignal{native: bindingSignalStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBindingSignal)
	return structGo
}
func finalizeBindingSignal(obj *BindingSignal) {
	bindingSignalStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *BooleanCellAccessibleClass) FieldParentClass() *RendererCellAccessibleClass {
	argValue := gi.FieldGet(booleanCellAccessibleClassStruct, recv.native, "parent_class")
	value := &RendererCellAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *BooleanCellAccessibleClass) SetFieldParentClass(value *RendererCellAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(booleanCellAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// BooleanCellAccessibleClassStruct creates an uninitialised BooleanCellAccessibleClass.
func BooleanCellAccessibleClassStruct() *BooleanCellAccessibleClass {
	err := booleanCellAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BooleanCellAccessibleClass{native: booleanCellAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBooleanCellAccessibleClass)
	return structGo
}
func finalizeBooleanCellAccessibleClass(obj *BooleanCellAccessibleClass) {
	booleanCellAccessibleClassStruct.Free(obj.native)
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

// BooleanCellAccessiblePrivateStruct creates an uninitialised BooleanCellAccessiblePrivate.
func BooleanCellAccessiblePrivateStruct() *BooleanCellAccessiblePrivate {
	err := booleanCellAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BooleanCellAccessiblePrivate{native: booleanCellAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBooleanCellAccessiblePrivate)
	return structGo
}
func finalizeBooleanCellAccessiblePrivate(obj *BooleanCellAccessiblePrivate) {
	booleanCellAccessiblePrivateStruct.Free(obj.native)
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

// FieldLeft returns the C field 'left'.
func (recv *Border) FieldLeft() int16 {
	argValue := gi.FieldGet(borderStruct, recv.native, "left")
	value := argValue.Int16()
	return value
}

// SetFieldLeft sets the value of the C field 'left'.
func (recv *Border) SetFieldLeft(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.FieldSet(borderStruct, recv.native, "left", argValue)
}

// FieldRight returns the C field 'right'.
func (recv *Border) FieldRight() int16 {
	argValue := gi.FieldGet(borderStruct, recv.native, "right")
	value := argValue.Int16()
	return value
}

// SetFieldRight sets the value of the C field 'right'.
func (recv *Border) SetFieldRight(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.FieldSet(borderStruct, recv.native, "right", argValue)
}

// FieldTop returns the C field 'top'.
func (recv *Border) FieldTop() int16 {
	argValue := gi.FieldGet(borderStruct, recv.native, "top")
	value := argValue.Int16()
	return value
}

// SetFieldTop sets the value of the C field 'top'.
func (recv *Border) SetFieldTop(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.FieldSet(borderStruct, recv.native, "top", argValue)
}

// FieldBottom returns the C field 'bottom'.
func (recv *Border) FieldBottom() int16 {
	argValue := gi.FieldGet(borderStruct, recv.native, "bottom")
	value := argValue.Int16()
	return value
}

// SetFieldBottom sets the value of the C field 'bottom'.
func (recv *Border) SetFieldBottom(value int16) {
	var argValue gi.Argument
	argValue.SetInt16(value)
	gi.FieldSet(borderStruct, recv.native, "bottom", argValue)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *BoxClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(boxClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *BoxClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(boxClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// BoxClassStruct creates an uninitialised BoxClass.
func BoxClassStruct() *BoxClass {
	err := boxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BoxClass{native: boxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBoxClass)
	return structGo
}
func finalizeBoxClass(obj *BoxClass) {
	boxClassStruct.Free(obj.native)
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

// BoxPrivateStruct creates an uninitialised BoxPrivate.
func BoxPrivateStruct() *BoxPrivate {
	err := boxPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BoxPrivate{native: boxPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBoxPrivate)
	return structGo
}
func finalizeBoxPrivate(obj *BoxPrivate) {
	boxPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'set_name' : for field getter : missing Type

// UNSUPPORTED : C value 'set_name' : for field setter : missing Type

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'add_child' : for field getter : missing Type

// UNSUPPORTED : C value 'add_child' : for field setter : missing Type

// UNSUPPORTED : C value 'set_buildable_property' : for field getter : missing Type

// UNSUPPORTED : C value 'set_buildable_property' : for field setter : missing Type

// UNSUPPORTED : C value 'construct_child' : for field getter : missing Type

// UNSUPPORTED : C value 'construct_child' : for field setter : missing Type

// UNSUPPORTED : C value 'custom_tag_start' : for field getter : missing Type

// UNSUPPORTED : C value 'custom_tag_start' : for field setter : missing Type

// UNSUPPORTED : C value 'custom_tag_end' : for field getter : missing Type

// UNSUPPORTED : C value 'custom_tag_end' : for field setter : missing Type

// UNSUPPORTED : C value 'custom_finished' : for field getter : missing Type

// UNSUPPORTED : C value 'custom_finished' : for field setter : missing Type

// UNSUPPORTED : C value 'parser_finished' : for field getter : missing Type

// UNSUPPORTED : C value 'parser_finished' : for field setter : missing Type

// UNSUPPORTED : C value 'get_internal_child' : for field getter : missing Type

// UNSUPPORTED : C value 'get_internal_child' : for field setter : missing Type

// BuildableIfaceStruct creates an uninitialised BuildableIface.
func BuildableIfaceStruct() *BuildableIface {
	err := buildableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BuildableIface{native: buildableIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBuildableIface)
	return structGo
}
func finalizeBuildableIface(obj *BuildableIface) {
	buildableIfaceStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_type_from_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_type_from_name' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// BuilderClassStruct creates an uninitialised BuilderClass.
func BuilderClassStruct() *BuilderClass {
	err := builderClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BuilderClass{native: builderClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBuilderClass)
	return structGo
}
func finalizeBuilderClass(obj *BuilderClass) {
	builderClassStruct.Free(obj.native)
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

// BuilderPrivateStruct creates an uninitialised BuilderPrivate.
func BuilderPrivateStruct() *BuilderPrivate {
	err := builderPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &BuilderPrivate{native: builderPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeBuilderPrivate)
	return structGo
}
func finalizeBuilderPrivate(obj *BuilderPrivate) {
	builderPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ButtonAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(buttonAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ButtonAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(buttonAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ButtonAccessibleClassStruct creates an uninitialised ButtonAccessibleClass.
func ButtonAccessibleClassStruct() *ButtonAccessibleClass {
	err := buttonAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ButtonAccessibleClass{native: buttonAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeButtonAccessibleClass)
	return structGo
}
func finalizeButtonAccessibleClass(obj *ButtonAccessibleClass) {
	buttonAccessibleClassStruct.Free(obj.native)
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

// ButtonAccessiblePrivateStruct creates an uninitialised ButtonAccessiblePrivate.
func ButtonAccessiblePrivateStruct() *ButtonAccessiblePrivate {
	err := buttonAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ButtonAccessiblePrivate{native: buttonAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeButtonAccessiblePrivate)
	return structGo
}
func finalizeButtonAccessiblePrivate(obj *ButtonAccessiblePrivate) {
	buttonAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ButtonBoxClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(buttonBoxClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ButtonBoxClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(buttonBoxClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ButtonBoxClassStruct creates an uninitialised ButtonBoxClass.
func ButtonBoxClassStruct() *ButtonBoxClass {
	err := buttonBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ButtonBoxClass{native: buttonBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeButtonBoxClass)
	return structGo
}
func finalizeButtonBoxClass(obj *ButtonBoxClass) {
	buttonBoxClassStruct.Free(obj.native)
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

// ButtonBoxPrivateStruct creates an uninitialised ButtonBoxPrivate.
func ButtonBoxPrivateStruct() *ButtonBoxPrivate {
	err := buttonBoxPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ButtonBoxPrivate{native: buttonBoxPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeButtonBoxPrivate)
	return structGo
}
func finalizeButtonBoxPrivate(obj *ButtonBoxPrivate) {
	buttonBoxPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ButtonClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(buttonClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ButtonClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(buttonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'pressed' : for field getter : missing Type

// UNSUPPORTED : C value 'pressed' : for field setter : missing Type

// UNSUPPORTED : C value 'released' : for field getter : missing Type

// UNSUPPORTED : C value 'released' : for field setter : missing Type

// UNSUPPORTED : C value 'clicked' : for field getter : missing Type

// UNSUPPORTED : C value 'clicked' : for field setter : missing Type

// UNSUPPORTED : C value 'enter' : for field getter : missing Type

// UNSUPPORTED : C value 'enter' : for field setter : missing Type

// UNSUPPORTED : C value 'leave' : for field getter : missing Type

// UNSUPPORTED : C value 'leave' : for field setter : missing Type

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ButtonClassStruct creates an uninitialised ButtonClass.
func ButtonClassStruct() *ButtonClass {
	err := buttonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ButtonClass{native: buttonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeButtonClass)
	return structGo
}
func finalizeButtonClass(obj *ButtonClass) {
	buttonClassStruct.Free(obj.native)
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

// ButtonPrivateStruct creates an uninitialised ButtonPrivate.
func ButtonPrivateStruct() *ButtonPrivate {
	err := buttonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ButtonPrivate{native: buttonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeButtonPrivate)
	return structGo
}
func finalizeButtonPrivate(obj *ButtonPrivate) {
	buttonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CalendarClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(calendarClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CalendarClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(calendarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'month_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'month_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'day_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'day_selected' : for field setter : missing Type

// UNSUPPORTED : C value 'day_selected_double_click' : for field getter : missing Type

// UNSUPPORTED : C value 'day_selected_double_click' : for field setter : missing Type

// UNSUPPORTED : C value 'prev_month' : for field getter : missing Type

// UNSUPPORTED : C value 'prev_month' : for field setter : missing Type

// UNSUPPORTED : C value 'next_month' : for field getter : missing Type

// UNSUPPORTED : C value 'next_month' : for field setter : missing Type

// UNSUPPORTED : C value 'prev_year' : for field getter : missing Type

// UNSUPPORTED : C value 'prev_year' : for field setter : missing Type

// UNSUPPORTED : C value 'next_year' : for field getter : missing Type

// UNSUPPORTED : C value 'next_year' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CalendarClassStruct creates an uninitialised CalendarClass.
func CalendarClassStruct() *CalendarClass {
	err := calendarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CalendarClass{native: calendarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCalendarClass)
	return structGo
}
func finalizeCalendarClass(obj *CalendarClass) {
	calendarClassStruct.Free(obj.native)
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

// CalendarPrivateStruct creates an uninitialised CalendarPrivate.
func CalendarPrivateStruct() *CalendarPrivate {
	err := calendarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CalendarPrivate{native: calendarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCalendarPrivate)
	return structGo
}
func finalizeCalendarPrivate(obj *CalendarPrivate) {
	calendarPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CellAccessibleClass) FieldParentClass() *AccessibleClass {
	argValue := gi.FieldGet(cellAccessibleClassStruct, recv.native, "parent_class")
	value := &AccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CellAccessibleClass) SetFieldParentClass(value *AccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'update_cache' : for field getter : missing Type

// UNSUPPORTED : C value 'update_cache' : for field setter : missing Type

// CellAccessibleClassStruct creates an uninitialised CellAccessibleClass.
func CellAccessibleClassStruct() *CellAccessibleClass {
	err := cellAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellAccessibleClass{native: cellAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellAccessibleClass)
	return structGo
}
func finalizeCellAccessibleClass(obj *CellAccessibleClass) {
	cellAccessibleClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'parent' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_cell_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'get_cell_extents' : for field setter : missing Type

// UNSUPPORTED : C value 'get_cell_area' : for field getter : missing Type

// UNSUPPORTED : C value 'get_cell_area' : for field setter : missing Type

// UNSUPPORTED : C value 'grab_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'grab_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'get_child_index' : for field getter : missing Type

// UNSUPPORTED : C value 'get_child_index' : for field setter : missing Type

// UNSUPPORTED : C value 'get_renderer_state' : for field getter : missing Type

// UNSUPPORTED : C value 'get_renderer_state' : for field setter : missing Type

// UNSUPPORTED : C value 'expand_collapse' : for field getter : missing Type

// UNSUPPORTED : C value 'expand_collapse' : for field setter : missing Type

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value 'edit' : for field getter : missing Type

// UNSUPPORTED : C value 'edit' : for field setter : missing Type

// UNSUPPORTED : C value 'update_relationset' : for field getter : missing Type

// UNSUPPORTED : C value 'update_relationset' : for field setter : missing Type

// UNSUPPORTED : C value 'get_cell_position' : for field getter : missing Type

// UNSUPPORTED : C value 'get_cell_position' : for field setter : missing Type

// UNSUPPORTED : C value 'get_column_header_cells' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_header_cells' : for field setter : missing Type

// UNSUPPORTED : C value 'get_row_header_cells' : for field getter : missing Type

// UNSUPPORTED : C value 'get_row_header_cells' : for field setter : missing Type

// CellAccessibleParentIfaceStruct creates an uninitialised CellAccessibleParentIface.
func CellAccessibleParentIfaceStruct() *CellAccessibleParentIface {
	err := cellAccessibleParentIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellAccessibleParentIface{native: cellAccessibleParentIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellAccessibleParentIface)
	return structGo
}
func finalizeCellAccessibleParentIface(obj *CellAccessibleParentIface) {
	cellAccessibleParentIfaceStruct.Free(obj.native)
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

// CellAccessiblePrivateStruct creates an uninitialised CellAccessiblePrivate.
func CellAccessiblePrivateStruct() *CellAccessiblePrivate {
	err := cellAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellAccessiblePrivate{native: cellAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellAccessiblePrivate)
	return structGo
}
func finalizeCellAccessiblePrivate(obj *CellAccessiblePrivate) {
	cellAccessiblePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellAreaBoxClassStruct creates an uninitialised CellAreaBoxClass.
func CellAreaBoxClassStruct() *CellAreaBoxClass {
	err := cellAreaBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellAreaBoxClass{native: cellAreaBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellAreaBoxClass)
	return structGo
}
func finalizeCellAreaBoxClass(obj *CellAreaBoxClass) {
	cellAreaBoxClassStruct.Free(obj.native)
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

// CellAreaBoxPrivateStruct creates an uninitialised CellAreaBoxPrivate.
func CellAreaBoxPrivateStruct() *CellAreaBoxPrivate {
	err := cellAreaBoxPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellAreaBoxPrivate{native: cellAreaBoxPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellAreaBoxPrivate)
	return structGo
}
func finalizeCellAreaBoxPrivate(obj *CellAreaBoxPrivate) {
	cellAreaBoxPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'add' : for field getter : missing Type

// UNSUPPORTED : C value 'add' : for field setter : missing Type

// UNSUPPORTED : C value 'remove' : for field getter : missing Type

// UNSUPPORTED : C value 'remove' : for field setter : missing Type

// UNSUPPORTED : C value 'foreach' : for field getter : missing Type

// UNSUPPORTED : C value 'foreach' : for field setter : missing Type

// UNSUPPORTED : C value 'foreach_alloc' : for field getter : missing Type

// UNSUPPORTED : C value 'foreach_alloc' : for field setter : missing Type

// UNSUPPORTED : C value 'event' : for field getter : missing Type

// UNSUPPORTED : C value 'event' : for field setter : missing Type

// UNSUPPORTED : C value 'render' : for field getter : missing Type

// UNSUPPORTED : C value 'render' : for field setter : missing Type

// UNSUPPORTED : C value 'apply_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'apply_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'create_context' : for field getter : missing Type

// UNSUPPORTED : C value 'create_context' : for field setter : missing Type

// UNSUPPORTED : C value 'copy_context' : for field getter : missing Type

// UNSUPPORTED : C value 'copy_context' : for field setter : missing Type

// UNSUPPORTED : C value 'get_request_mode' : for field getter : missing Type

// UNSUPPORTED : C value 'get_request_mode' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_width' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_width' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_for_width' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_for_width' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_height' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_height' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_width_for_height' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_width_for_height' : for field setter : missing Type

// UNSUPPORTED : C value 'set_cell_property' : for field getter : missing Type

// UNSUPPORTED : C value 'set_cell_property' : for field setter : missing Type

// UNSUPPORTED : C value 'get_cell_property' : for field getter : missing Type

// UNSUPPORTED : C value 'get_cell_property' : for field setter : missing Type

// UNSUPPORTED : C value 'focus' : for field getter : missing Type

// UNSUPPORTED : C value 'focus' : for field setter : missing Type

// UNSUPPORTED : C value 'is_activatable' : for field getter : missing Type

// UNSUPPORTED : C value 'is_activatable' : for field setter : missing Type

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

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

// CellAreaClassStruct creates an uninitialised CellAreaClass.
func CellAreaClassStruct() *CellAreaClass {
	err := cellAreaClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellAreaClass{native: cellAreaClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellAreaClass)
	return structGo
}
func finalizeCellAreaClass(obj *CellAreaClass) {
	cellAreaClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'allocate' : for field getter : missing Type

// UNSUPPORTED : C value 'allocate' : for field setter : missing Type

// UNSUPPORTED : C value 'reset' : for field getter : missing Type

// UNSUPPORTED : C value 'reset' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_for_width' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_for_width' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_width_for_height' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_width_for_height' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// CellAreaContextClassStruct creates an uninitialised CellAreaContextClass.
func CellAreaContextClassStruct() *CellAreaContextClass {
	err := cellAreaContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellAreaContextClass{native: cellAreaContextClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellAreaContextClass)
	return structGo
}
func finalizeCellAreaContextClass(obj *CellAreaContextClass) {
	cellAreaContextClassStruct.Free(obj.native)
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

// CellAreaContextPrivateStruct creates an uninitialised CellAreaContextPrivate.
func CellAreaContextPrivateStruct() *CellAreaContextPrivate {
	err := cellAreaContextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellAreaContextPrivate{native: cellAreaContextPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellAreaContextPrivate)
	return structGo
}
func finalizeCellAreaContextPrivate(obj *CellAreaContextPrivate) {
	cellAreaContextPrivateStruct.Free(obj.native)
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

// CellAreaPrivateStruct creates an uninitialised CellAreaPrivate.
func CellAreaPrivateStruct() *CellAreaPrivate {
	err := cellAreaPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellAreaPrivate{native: cellAreaPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellAreaPrivate)
	return structGo
}
func finalizeCellAreaPrivate(obj *CellAreaPrivate) {
	cellAreaPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'editing_done' : for field getter : missing Type

// UNSUPPORTED : C value 'editing_done' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_widget' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_widget' : for field setter : missing Type

// UNSUPPORTED : C value 'start_editing' : for field getter : missing Type

// UNSUPPORTED : C value 'start_editing' : for field setter : missing Type

// CellEditableIfaceStruct creates an uninitialised CellEditableIface.
func CellEditableIfaceStruct() *CellEditableIface {
	err := cellEditableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellEditableIface{native: cellEditableIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellEditableIface)
	return structGo
}
func finalizeCellEditableIface(obj *CellEditableIface) {
	cellEditableIfaceStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'pack_start' : for field getter : missing Type

// UNSUPPORTED : C value 'pack_start' : for field setter : missing Type

// UNSUPPORTED : C value 'pack_end' : for field getter : missing Type

// UNSUPPORTED : C value 'pack_end' : for field setter : missing Type

// UNSUPPORTED : C value 'clear' : for field getter : missing Type

// UNSUPPORTED : C value 'clear' : for field setter : missing Type

// UNSUPPORTED : C value 'add_attribute' : for field getter : missing Type

// UNSUPPORTED : C value 'add_attribute' : for field setter : missing Type

// UNSUPPORTED : C value 'set_cell_data_func' : for field getter : missing Type

// UNSUPPORTED : C value 'set_cell_data_func' : for field setter : missing Type

// UNSUPPORTED : C value 'clear_attributes' : for field getter : missing Type

// UNSUPPORTED : C value 'clear_attributes' : for field setter : missing Type

// UNSUPPORTED : C value 'reorder' : for field getter : missing Type

// UNSUPPORTED : C value 'reorder' : for field setter : missing Type

// UNSUPPORTED : C value 'get_cells' : for field getter : missing Type

// UNSUPPORTED : C value 'get_cells' : for field setter : missing Type

// UNSUPPORTED : C value 'get_area' : for field getter : missing Type

// UNSUPPORTED : C value 'get_area' : for field setter : missing Type

// CellLayoutIfaceStruct creates an uninitialised CellLayoutIface.
func CellLayoutIfaceStruct() *CellLayoutIface {
	err := cellLayoutIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellLayoutIface{native: cellLayoutIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellLayoutIface)
	return structGo
}
func finalizeCellLayoutIface(obj *CellLayoutIface) {
	cellLayoutIfaceStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CellRendererAccelClass) FieldParentClass() *CellRendererTextClass {
	argValue := gi.FieldGet(cellRendererAccelClassStruct, recv.native, "parent_class")
	value := &CellRendererTextClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CellRendererAccelClass) SetFieldParentClass(value *CellRendererTextClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellRendererAccelClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'accel_edited' : for field getter : missing Type

// UNSUPPORTED : C value 'accel_edited' : for field setter : missing Type

// UNSUPPORTED : C value 'accel_cleared' : for field getter : missing Type

// UNSUPPORTED : C value 'accel_cleared' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellRendererAccelClassStruct creates an uninitialised CellRendererAccelClass.
func CellRendererAccelClassStruct() *CellRendererAccelClass {
	err := cellRendererAccelClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererAccelClass{native: cellRendererAccelClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererAccelClass)
	return structGo
}
func finalizeCellRendererAccelClass(obj *CellRendererAccelClass) {
	cellRendererAccelClassStruct.Free(obj.native)
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

// CellRendererAccelPrivateStruct creates an uninitialised CellRendererAccelPrivate.
func CellRendererAccelPrivateStruct() *CellRendererAccelPrivate {
	err := cellRendererAccelPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererAccelPrivate{native: cellRendererAccelPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererAccelPrivate)
	return structGo
}
func finalizeCellRendererAccelPrivate(obj *CellRendererAccelPrivate) {
	cellRendererAccelPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'get_request_mode' : for field getter : missing Type

// UNSUPPORTED : C value 'get_request_mode' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_width' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_width' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_for_width' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_for_width' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_height' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_height' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_width_for_height' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_width_for_height' : for field setter : missing Type

// UNSUPPORTED : C value 'get_aligned_area' : for field getter : missing Type

// UNSUPPORTED : C value 'get_aligned_area' : for field setter : missing Type

// UNSUPPORTED : C value 'get_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_size' : for field setter : missing Type

// UNSUPPORTED : C value 'render' : for field getter : missing Type

// UNSUPPORTED : C value 'render' : for field setter : missing Type

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value 'start_editing' : for field getter : missing Type

// UNSUPPORTED : C value 'start_editing' : for field setter : missing Type

// UNSUPPORTED : C value 'editing_canceled' : for field getter : missing Type

// UNSUPPORTED : C value 'editing_canceled' : for field setter : missing Type

// UNSUPPORTED : C value 'editing_started' : for field getter : missing Type

// UNSUPPORTED : C value 'editing_started' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value 'gtk_cell_renderer_class_set_accessible_type' : parameter 'type' of type 'GType' not supported

// CellRendererClassStruct creates an uninitialised CellRendererClass.
func CellRendererClassStruct() *CellRendererClass {
	err := cellRendererClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererClass{native: cellRendererClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererClass)
	return structGo
}
func finalizeCellRendererClass(obj *CellRendererClass) {
	cellRendererClassStruct.Free(obj.native)
}

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

// CellRendererClassPrivateStruct creates an uninitialised CellRendererClassPrivate.
func CellRendererClassPrivateStruct() *CellRendererClassPrivate {
	err := cellRendererClassPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererClassPrivate{native: cellRendererClassPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererClassPrivate)
	return structGo
}
func finalizeCellRendererClassPrivate(obj *CellRendererClassPrivate) {
	cellRendererClassPrivateStruct.Free(obj.native)
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

// FieldParent returns the C field 'parent'.
func (recv *CellRendererComboClass) FieldParent() *CellRendererTextClass {
	argValue := gi.FieldGet(cellRendererComboClassStruct, recv.native, "parent")
	value := &CellRendererTextClass{native: argValue.Pointer()}
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CellRendererComboClass) SetFieldParent(value *CellRendererTextClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellRendererComboClassStruct, recv.native, "parent", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellRendererComboClassStruct creates an uninitialised CellRendererComboClass.
func CellRendererComboClassStruct() *CellRendererComboClass {
	err := cellRendererComboClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererComboClass{native: cellRendererComboClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererComboClass)
	return structGo
}
func finalizeCellRendererComboClass(obj *CellRendererComboClass) {
	cellRendererComboClassStruct.Free(obj.native)
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

// CellRendererComboPrivateStruct creates an uninitialised CellRendererComboPrivate.
func CellRendererComboPrivateStruct() *CellRendererComboPrivate {
	err := cellRendererComboPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererComboPrivate{native: cellRendererComboPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererComboPrivate)
	return structGo
}
func finalizeCellRendererComboPrivate(obj *CellRendererComboPrivate) {
	cellRendererComboPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CellRendererPixbufClass) FieldParentClass() *CellRendererClass {
	argValue := gi.FieldGet(cellRendererPixbufClassStruct, recv.native, "parent_class")
	value := &CellRendererClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CellRendererPixbufClass) SetFieldParentClass(value *CellRendererClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellRendererPixbufClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellRendererPixbufClassStruct creates an uninitialised CellRendererPixbufClass.
func CellRendererPixbufClassStruct() *CellRendererPixbufClass {
	err := cellRendererPixbufClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererPixbufClass{native: cellRendererPixbufClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererPixbufClass)
	return structGo
}
func finalizeCellRendererPixbufClass(obj *CellRendererPixbufClass) {
	cellRendererPixbufClassStruct.Free(obj.native)
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

// CellRendererPixbufPrivateStruct creates an uninitialised CellRendererPixbufPrivate.
func CellRendererPixbufPrivateStruct() *CellRendererPixbufPrivate {
	err := cellRendererPixbufPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererPixbufPrivate{native: cellRendererPixbufPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererPixbufPrivate)
	return structGo
}
func finalizeCellRendererPixbufPrivate(obj *CellRendererPixbufPrivate) {
	cellRendererPixbufPrivateStruct.Free(obj.native)
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

// CellRendererPrivateStruct creates an uninitialised CellRendererPrivate.
func CellRendererPrivateStruct() *CellRendererPrivate {
	err := cellRendererPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererPrivate{native: cellRendererPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererPrivate)
	return structGo
}
func finalizeCellRendererPrivate(obj *CellRendererPrivate) {
	cellRendererPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CellRendererProgressClass) FieldParentClass() *CellRendererClass {
	argValue := gi.FieldGet(cellRendererProgressClassStruct, recv.native, "parent_class")
	value := &CellRendererClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CellRendererProgressClass) SetFieldParentClass(value *CellRendererClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellRendererProgressClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellRendererProgressClassStruct creates an uninitialised CellRendererProgressClass.
func CellRendererProgressClassStruct() *CellRendererProgressClass {
	err := cellRendererProgressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererProgressClass{native: cellRendererProgressClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererProgressClass)
	return structGo
}
func finalizeCellRendererProgressClass(obj *CellRendererProgressClass) {
	cellRendererProgressClassStruct.Free(obj.native)
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

// CellRendererProgressPrivateStruct creates an uninitialised CellRendererProgressPrivate.
func CellRendererProgressPrivateStruct() *CellRendererProgressPrivate {
	err := cellRendererProgressPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererProgressPrivate{native: cellRendererProgressPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererProgressPrivate)
	return structGo
}
func finalizeCellRendererProgressPrivate(obj *CellRendererProgressPrivate) {
	cellRendererProgressPrivateStruct.Free(obj.native)
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

// FieldParent returns the C field 'parent'.
func (recv *CellRendererSpinClass) FieldParent() *CellRendererTextClass {
	argValue := gi.FieldGet(cellRendererSpinClassStruct, recv.native, "parent")
	value := &CellRendererTextClass{native: argValue.Pointer()}
	return value
}

// SetFieldParent sets the value of the C field 'parent'.
func (recv *CellRendererSpinClass) SetFieldParent(value *CellRendererTextClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellRendererSpinClassStruct, recv.native, "parent", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellRendererSpinClassStruct creates an uninitialised CellRendererSpinClass.
func CellRendererSpinClassStruct() *CellRendererSpinClass {
	err := cellRendererSpinClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererSpinClass{native: cellRendererSpinClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererSpinClass)
	return structGo
}
func finalizeCellRendererSpinClass(obj *CellRendererSpinClass) {
	cellRendererSpinClassStruct.Free(obj.native)
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

// CellRendererSpinPrivateStruct creates an uninitialised CellRendererSpinPrivate.
func CellRendererSpinPrivateStruct() *CellRendererSpinPrivate {
	err := cellRendererSpinPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererSpinPrivate{native: cellRendererSpinPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererSpinPrivate)
	return structGo
}
func finalizeCellRendererSpinPrivate(obj *CellRendererSpinPrivate) {
	cellRendererSpinPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CellRendererSpinnerClass) FieldParentClass() *CellRendererClass {
	argValue := gi.FieldGet(cellRendererSpinnerClassStruct, recv.native, "parent_class")
	value := &CellRendererClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CellRendererSpinnerClass) SetFieldParentClass(value *CellRendererClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellRendererSpinnerClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellRendererSpinnerClassStruct creates an uninitialised CellRendererSpinnerClass.
func CellRendererSpinnerClassStruct() *CellRendererSpinnerClass {
	err := cellRendererSpinnerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererSpinnerClass{native: cellRendererSpinnerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererSpinnerClass)
	return structGo
}
func finalizeCellRendererSpinnerClass(obj *CellRendererSpinnerClass) {
	cellRendererSpinnerClassStruct.Free(obj.native)
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

// CellRendererSpinnerPrivateStruct creates an uninitialised CellRendererSpinnerPrivate.
func CellRendererSpinnerPrivateStruct() *CellRendererSpinnerPrivate {
	err := cellRendererSpinnerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererSpinnerPrivate{native: cellRendererSpinnerPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererSpinnerPrivate)
	return structGo
}
func finalizeCellRendererSpinnerPrivate(obj *CellRendererSpinnerPrivate) {
	cellRendererSpinnerPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CellRendererTextClass) FieldParentClass() *CellRendererClass {
	argValue := gi.FieldGet(cellRendererTextClassStruct, recv.native, "parent_class")
	value := &CellRendererClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CellRendererTextClass) SetFieldParentClass(value *CellRendererClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellRendererTextClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'edited' : for field getter : missing Type

// UNSUPPORTED : C value 'edited' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellRendererTextClassStruct creates an uninitialised CellRendererTextClass.
func CellRendererTextClassStruct() *CellRendererTextClass {
	err := cellRendererTextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererTextClass{native: cellRendererTextClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererTextClass)
	return structGo
}
func finalizeCellRendererTextClass(obj *CellRendererTextClass) {
	cellRendererTextClassStruct.Free(obj.native)
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

// CellRendererTextPrivateStruct creates an uninitialised CellRendererTextPrivate.
func CellRendererTextPrivateStruct() *CellRendererTextPrivate {
	err := cellRendererTextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererTextPrivate{native: cellRendererTextPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererTextPrivate)
	return structGo
}
func finalizeCellRendererTextPrivate(obj *CellRendererTextPrivate) {
	cellRendererTextPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CellRendererToggleClass) FieldParentClass() *CellRendererClass {
	argValue := gi.FieldGet(cellRendererToggleClassStruct, recv.native, "parent_class")
	value := &CellRendererClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CellRendererToggleClass) SetFieldParentClass(value *CellRendererClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellRendererToggleClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'toggled' : for field getter : missing Type

// UNSUPPORTED : C value 'toggled' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellRendererToggleClassStruct creates an uninitialised CellRendererToggleClass.
func CellRendererToggleClassStruct() *CellRendererToggleClass {
	err := cellRendererToggleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererToggleClass{native: cellRendererToggleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererToggleClass)
	return structGo
}
func finalizeCellRendererToggleClass(obj *CellRendererToggleClass) {
	cellRendererToggleClassStruct.Free(obj.native)
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

// CellRendererTogglePrivateStruct creates an uninitialised CellRendererTogglePrivate.
func CellRendererTogglePrivateStruct() *CellRendererTogglePrivate {
	err := cellRendererTogglePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellRendererTogglePrivate{native: cellRendererTogglePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellRendererTogglePrivate)
	return structGo
}
func finalizeCellRendererTogglePrivate(obj *CellRendererTogglePrivate) {
	cellRendererTogglePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CellViewClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(cellViewClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CellViewClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(cellViewClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CellViewClassStruct creates an uninitialised CellViewClass.
func CellViewClassStruct() *CellViewClass {
	err := cellViewClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellViewClass{native: cellViewClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellViewClass)
	return structGo
}
func finalizeCellViewClass(obj *CellViewClass) {
	cellViewClassStruct.Free(obj.native)
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

// CellViewPrivateStruct creates an uninitialised CellViewPrivate.
func CellViewPrivateStruct() *CellViewPrivate {
	err := cellViewPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CellViewPrivate{native: cellViewPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCellViewPrivate)
	return structGo
}
func finalizeCellViewPrivate(obj *CellViewPrivate) {
	cellViewPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CheckButtonClass) FieldParentClass() *ToggleButtonClass {
	argValue := gi.FieldGet(checkButtonClassStruct, recv.native, "parent_class")
	value := &ToggleButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CheckButtonClass) SetFieldParentClass(value *ToggleButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(checkButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'draw_indicator' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_indicator' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CheckButtonClassStruct creates an uninitialised CheckButtonClass.
func CheckButtonClassStruct() *CheckButtonClass {
	err := checkButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CheckButtonClass{native: checkButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCheckButtonClass)
	return structGo
}
func finalizeCheckButtonClass(obj *CheckButtonClass) {
	checkButtonClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CheckMenuItemAccessibleClass) FieldParentClass() *MenuItemAccessibleClass {
	argValue := gi.FieldGet(checkMenuItemAccessibleClassStruct, recv.native, "parent_class")
	value := &MenuItemAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CheckMenuItemAccessibleClass) SetFieldParentClass(value *MenuItemAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(checkMenuItemAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// CheckMenuItemAccessibleClassStruct creates an uninitialised CheckMenuItemAccessibleClass.
func CheckMenuItemAccessibleClassStruct() *CheckMenuItemAccessibleClass {
	err := checkMenuItemAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CheckMenuItemAccessibleClass{native: checkMenuItemAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCheckMenuItemAccessibleClass)
	return structGo
}
func finalizeCheckMenuItemAccessibleClass(obj *CheckMenuItemAccessibleClass) {
	checkMenuItemAccessibleClassStruct.Free(obj.native)
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

// CheckMenuItemAccessiblePrivateStruct creates an uninitialised CheckMenuItemAccessiblePrivate.
func CheckMenuItemAccessiblePrivateStruct() *CheckMenuItemAccessiblePrivate {
	err := checkMenuItemAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CheckMenuItemAccessiblePrivate{native: checkMenuItemAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCheckMenuItemAccessiblePrivate)
	return structGo
}
func finalizeCheckMenuItemAccessiblePrivate(obj *CheckMenuItemAccessiblePrivate) {
	checkMenuItemAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *CheckMenuItemClass) FieldParentClass() *MenuItemClass {
	argValue := gi.FieldGet(checkMenuItemClassStruct, recv.native, "parent_class")
	value := &MenuItemClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *CheckMenuItemClass) SetFieldParentClass(value *MenuItemClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(checkMenuItemClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'toggled' : for field getter : missing Type

// UNSUPPORTED : C value 'toggled' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_indicator' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_indicator' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CheckMenuItemClassStruct creates an uninitialised CheckMenuItemClass.
func CheckMenuItemClassStruct() *CheckMenuItemClass {
	err := checkMenuItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CheckMenuItemClass{native: checkMenuItemClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCheckMenuItemClass)
	return structGo
}
func finalizeCheckMenuItemClass(obj *CheckMenuItemClass) {
	checkMenuItemClassStruct.Free(obj.native)
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

// CheckMenuItemPrivateStruct creates an uninitialised CheckMenuItemPrivate.
func CheckMenuItemPrivateStruct() *CheckMenuItemPrivate {
	err := checkMenuItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CheckMenuItemPrivate{native: checkMenuItemPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCheckMenuItemPrivate)
	return structGo
}
func finalizeCheckMenuItemPrivate(obj *CheckMenuItemPrivate) {
	checkMenuItemPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ColorButtonClass) FieldParentClass() *ButtonClass {
	argValue := gi.FieldGet(colorButtonClassStruct, recv.native, "parent_class")
	value := &ButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ColorButtonClass) SetFieldParentClass(value *ButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(colorButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'color_set' : for field getter : missing Type

// UNSUPPORTED : C value 'color_set' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ColorButtonClassStruct creates an uninitialised ColorButtonClass.
func ColorButtonClassStruct() *ColorButtonClass {
	err := colorButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorButtonClass{native: colorButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorButtonClass)
	return structGo
}
func finalizeColorButtonClass(obj *ColorButtonClass) {
	colorButtonClassStruct.Free(obj.native)
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

// ColorButtonPrivateStruct creates an uninitialised ColorButtonPrivate.
func ColorButtonPrivateStruct() *ColorButtonPrivate {
	err := colorButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorButtonPrivate{native: colorButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorButtonPrivate)
	return structGo
}
func finalizeColorButtonPrivate(obj *ColorButtonPrivate) {
	colorButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ColorChooserDialogClass) FieldParentClass() *DialogClass {
	argValue := gi.FieldGet(colorChooserDialogClassStruct, recv.native, "parent_class")
	value := &DialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ColorChooserDialogClass) SetFieldParentClass(value *DialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(colorChooserDialogClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ColorChooserDialogClassStruct creates an uninitialised ColorChooserDialogClass.
func ColorChooserDialogClassStruct() *ColorChooserDialogClass {
	err := colorChooserDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorChooserDialogClass{native: colorChooserDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorChooserDialogClass)
	return structGo
}
func finalizeColorChooserDialogClass(obj *ColorChooserDialogClass) {
	colorChooserDialogClassStruct.Free(obj.native)
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

// ColorChooserDialogPrivateStruct creates an uninitialised ColorChooserDialogPrivate.
func ColorChooserDialogPrivateStruct() *ColorChooserDialogPrivate {
	err := colorChooserDialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorChooserDialogPrivate{native: colorChooserDialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorChooserDialogPrivate)
	return structGo
}
func finalizeColorChooserDialogPrivate(obj *ColorChooserDialogPrivate) {
	colorChooserDialogPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'base_interface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'base_interface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_rgba' : for field getter : missing Type

// UNSUPPORTED : C value 'get_rgba' : for field setter : missing Type

// UNSUPPORTED : C value 'set_rgba' : for field getter : missing Type

// UNSUPPORTED : C value 'set_rgba' : for field setter : missing Type

// UNSUPPORTED : C value 'add_palette' : for field getter : missing Type

// UNSUPPORTED : C value 'add_palette' : for field setter : missing Type

// UNSUPPORTED : C value 'color_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'color_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// ColorChooserInterfaceStruct creates an uninitialised ColorChooserInterface.
func ColorChooserInterfaceStruct() *ColorChooserInterface {
	err := colorChooserInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorChooserInterface{native: colorChooserInterfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorChooserInterface)
	return structGo
}
func finalizeColorChooserInterface(obj *ColorChooserInterface) {
	colorChooserInterfaceStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ColorChooserWidgetClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(colorChooserWidgetClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ColorChooserWidgetClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(colorChooserWidgetClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// ColorChooserWidgetClassStruct creates an uninitialised ColorChooserWidgetClass.
func ColorChooserWidgetClassStruct() *ColorChooserWidgetClass {
	err := colorChooserWidgetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorChooserWidgetClass{native: colorChooserWidgetClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorChooserWidgetClass)
	return structGo
}
func finalizeColorChooserWidgetClass(obj *ColorChooserWidgetClass) {
	colorChooserWidgetClassStruct.Free(obj.native)
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

// ColorChooserWidgetPrivateStruct creates an uninitialised ColorChooserWidgetPrivate.
func ColorChooserWidgetPrivateStruct() *ColorChooserWidgetPrivate {
	err := colorChooserWidgetPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorChooserWidgetPrivate{native: colorChooserWidgetPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorChooserWidgetPrivate)
	return structGo
}
func finalizeColorChooserWidgetPrivate(obj *ColorChooserWidgetPrivate) {
	colorChooserWidgetPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ColorSelectionClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(colorSelectionClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ColorSelectionClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(colorSelectionClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'color_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'color_changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ColorSelectionClassStruct creates an uninitialised ColorSelectionClass.
func ColorSelectionClassStruct() *ColorSelectionClass {
	err := colorSelectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorSelectionClass{native: colorSelectionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorSelectionClass)
	return structGo
}
func finalizeColorSelectionClass(obj *ColorSelectionClass) {
	colorSelectionClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ColorSelectionDialogClass) FieldParentClass() *DialogClass {
	argValue := gi.FieldGet(colorSelectionDialogClassStruct, recv.native, "parent_class")
	value := &DialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ColorSelectionDialogClass) SetFieldParentClass(value *DialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(colorSelectionDialogClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ColorSelectionDialogClassStruct creates an uninitialised ColorSelectionDialogClass.
func ColorSelectionDialogClassStruct() *ColorSelectionDialogClass {
	err := colorSelectionDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorSelectionDialogClass{native: colorSelectionDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorSelectionDialogClass)
	return structGo
}
func finalizeColorSelectionDialogClass(obj *ColorSelectionDialogClass) {
	colorSelectionDialogClassStruct.Free(obj.native)
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

// ColorSelectionDialogPrivateStruct creates an uninitialised ColorSelectionDialogPrivate.
func ColorSelectionDialogPrivateStruct() *ColorSelectionDialogPrivate {
	err := colorSelectionDialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorSelectionDialogPrivate{native: colorSelectionDialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorSelectionDialogPrivate)
	return structGo
}
func finalizeColorSelectionDialogPrivate(obj *ColorSelectionDialogPrivate) {
	colorSelectionDialogPrivateStruct.Free(obj.native)
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

// ColorSelectionPrivateStruct creates an uninitialised ColorSelectionPrivate.
func ColorSelectionPrivateStruct() *ColorSelectionPrivate {
	err := colorSelectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ColorSelectionPrivate{native: colorSelectionPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeColorSelectionPrivate)
	return structGo
}
func finalizeColorSelectionPrivate(obj *ColorSelectionPrivate) {
	colorSelectionPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ComboBoxAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(comboBoxAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ComboBoxAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(comboBoxAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ComboBoxAccessibleClassStruct creates an uninitialised ComboBoxAccessibleClass.
func ComboBoxAccessibleClassStruct() *ComboBoxAccessibleClass {
	err := comboBoxAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ComboBoxAccessibleClass{native: comboBoxAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeComboBoxAccessibleClass)
	return structGo
}
func finalizeComboBoxAccessibleClass(obj *ComboBoxAccessibleClass) {
	comboBoxAccessibleClassStruct.Free(obj.native)
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

// ComboBoxAccessiblePrivateStruct creates an uninitialised ComboBoxAccessiblePrivate.
func ComboBoxAccessiblePrivateStruct() *ComboBoxAccessiblePrivate {
	err := comboBoxAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ComboBoxAccessiblePrivate{native: comboBoxAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeComboBoxAccessiblePrivate)
	return structGo
}
func finalizeComboBoxAccessiblePrivate(obj *ComboBoxAccessiblePrivate) {
	comboBoxAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ComboBoxClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(comboBoxClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ComboBoxClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(comboBoxClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'format_entry_text' : for field getter : missing Type

// UNSUPPORTED : C value 'format_entry_text' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// ComboBoxClassStruct creates an uninitialised ComboBoxClass.
func ComboBoxClassStruct() *ComboBoxClass {
	err := comboBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ComboBoxClass{native: comboBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeComboBoxClass)
	return structGo
}
func finalizeComboBoxClass(obj *ComboBoxClass) {
	comboBoxClassStruct.Free(obj.native)
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

// ComboBoxPrivateStruct creates an uninitialised ComboBoxPrivate.
func ComboBoxPrivateStruct() *ComboBoxPrivate {
	err := comboBoxPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ComboBoxPrivate{native: comboBoxPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeComboBoxPrivate)
	return structGo
}
func finalizeComboBoxPrivate(obj *ComboBoxPrivate) {
	comboBoxPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ComboBoxTextClass) FieldParentClass() *ComboBoxClass {
	argValue := gi.FieldGet(comboBoxTextClassStruct, recv.native, "parent_class")
	value := &ComboBoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ComboBoxTextClass) SetFieldParentClass(value *ComboBoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(comboBoxTextClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ComboBoxTextClassStruct creates an uninitialised ComboBoxTextClass.
func ComboBoxTextClassStruct() *ComboBoxTextClass {
	err := comboBoxTextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ComboBoxTextClass{native: comboBoxTextClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeComboBoxTextClass)
	return structGo
}
func finalizeComboBoxTextClass(obj *ComboBoxTextClass) {
	comboBoxTextClassStruct.Free(obj.native)
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

// ComboBoxTextPrivateStruct creates an uninitialised ComboBoxTextPrivate.
func ComboBoxTextPrivateStruct() *ComboBoxTextPrivate {
	err := comboBoxTextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ComboBoxTextPrivate{native: comboBoxTextPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeComboBoxTextPrivate)
	return structGo
}
func finalizeComboBoxTextPrivate(obj *ComboBoxTextPrivate) {
	comboBoxTextPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ContainerAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(containerAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ContainerAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(containerAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'add_gtk' : for field getter : missing Type

// UNSUPPORTED : C value 'add_gtk' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_gtk' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_gtk' : for field setter : missing Type

// ContainerAccessibleClassStruct creates an uninitialised ContainerAccessibleClass.
func ContainerAccessibleClassStruct() *ContainerAccessibleClass {
	err := containerAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContainerAccessibleClass{native: containerAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContainerAccessibleClass)
	return structGo
}
func finalizeContainerAccessibleClass(obj *ContainerAccessibleClass) {
	containerAccessibleClassStruct.Free(obj.native)
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

// ContainerAccessiblePrivateStruct creates an uninitialised ContainerAccessiblePrivate.
func ContainerAccessiblePrivateStruct() *ContainerAccessiblePrivate {
	err := containerAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContainerAccessiblePrivate{native: containerAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContainerAccessiblePrivate)
	return structGo
}
func finalizeContainerAccessiblePrivate(obj *ContainerAccessiblePrivate) {
	containerAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ContainerCellAccessibleClass) FieldParentClass() *CellAccessibleClass {
	argValue := gi.FieldGet(containerCellAccessibleClassStruct, recv.native, "parent_class")
	value := &CellAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ContainerCellAccessibleClass) SetFieldParentClass(value *CellAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(containerCellAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ContainerCellAccessibleClassStruct creates an uninitialised ContainerCellAccessibleClass.
func ContainerCellAccessibleClassStruct() *ContainerCellAccessibleClass {
	err := containerCellAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContainerCellAccessibleClass{native: containerCellAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContainerCellAccessibleClass)
	return structGo
}
func finalizeContainerCellAccessibleClass(obj *ContainerCellAccessibleClass) {
	containerCellAccessibleClassStruct.Free(obj.native)
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

// ContainerCellAccessiblePrivateStruct creates an uninitialised ContainerCellAccessiblePrivate.
func ContainerCellAccessiblePrivateStruct() *ContainerCellAccessiblePrivate {
	err := containerCellAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContainerCellAccessiblePrivate{native: containerCellAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContainerCellAccessiblePrivate)
	return structGo
}
func finalizeContainerCellAccessiblePrivate(obj *ContainerCellAccessiblePrivate) {
	containerCellAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ContainerClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(containerClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ContainerClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(containerClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'add' : for field getter : missing Type

// UNSUPPORTED : C value 'add' : for field setter : missing Type

// UNSUPPORTED : C value 'remove' : for field getter : missing Type

// UNSUPPORTED : C value 'remove' : for field setter : missing Type

// UNSUPPORTED : C value 'check_resize' : for field getter : missing Type

// UNSUPPORTED : C value 'check_resize' : for field setter : missing Type

// UNSUPPORTED : C value 'forall' : for field getter : missing Type

// UNSUPPORTED : C value 'forall' : for field setter : missing Type

// UNSUPPORTED : C value 'set_focus_child' : for field getter : missing Type

// UNSUPPORTED : C value 'set_focus_child' : for field setter : missing Type

// UNSUPPORTED : C value 'child_type' : for field getter : missing Type

// UNSUPPORTED : C value 'child_type' : for field setter : missing Type

// UNSUPPORTED : C value 'composite_name' : for field getter : missing Type

// UNSUPPORTED : C value 'composite_name' : for field setter : missing Type

// UNSUPPORTED : C value 'set_child_property' : for field getter : missing Type

// UNSUPPORTED : C value 'set_child_property' : for field setter : missing Type

// UNSUPPORTED : C value 'get_child_property' : for field getter : missing Type

// UNSUPPORTED : C value 'get_child_property' : for field setter : missing Type

// UNSUPPORTED : C value 'get_path_for_child' : for field getter : missing Type

// UNSUPPORTED : C value 'get_path_for_child' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

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

// UNSUPPORTED : C value 'gtk_container_class_install_child_properties' : parameter 'pspecs' of type 'nil' not supported

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

// ContainerClassStruct creates an uninitialised ContainerClass.
func ContainerClassStruct() *ContainerClass {
	err := containerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContainerClass{native: containerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContainerClass)
	return structGo
}
func finalizeContainerClass(obj *ContainerClass) {
	containerClassStruct.Free(obj.native)
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

// ContainerPrivateStruct creates an uninitialised ContainerPrivate.
func ContainerPrivateStruct() *ContainerPrivate {
	err := containerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ContainerPrivate{native: containerPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeContainerPrivate)
	return structGo
}
func finalizeContainerPrivate(obj *ContainerPrivate) {
	containerPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parsing_error' : for field getter : missing Type

// UNSUPPORTED : C value 'parsing_error' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// CssProviderClassStruct creates an uninitialised CssProviderClass.
func CssProviderClassStruct() *CssProviderClass {
	err := cssProviderClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CssProviderClass{native: cssProviderClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCssProviderClass)
	return structGo
}
func finalizeCssProviderClass(obj *CssProviderClass) {
	cssProviderClassStruct.Free(obj.native)
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

// CssProviderPrivateStruct creates an uninitialised CssProviderPrivate.
func CssProviderPrivateStruct() *CssProviderPrivate {
	err := cssProviderPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CssProviderPrivate{native: cssProviderPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCssProviderPrivate)
	return structGo
}
func finalizeCssProviderPrivate(obj *CssProviderPrivate) {
	cssProviderPrivateStruct.Free(obj.native)
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

// CssSectionStruct creates an uninitialised CssSection.
func CssSectionStruct() *CssSection {
	err := cssSectionStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &CssSection{native: cssSectionStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeCssSection)
	return structGo
}
func finalizeCssSection(obj *CssSection) {
	cssSectionStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *DialogClass) FieldParentClass() *WindowClass {
	argValue := gi.FieldGet(dialogClassStruct, recv.native, "parent_class")
	value := &WindowClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *DialogClass) SetFieldParentClass(value *WindowClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(dialogClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'response' : for field getter : missing Type

// UNSUPPORTED : C value 'response' : for field setter : missing Type

// UNSUPPORTED : C value 'close' : for field getter : missing Type

// UNSUPPORTED : C value 'close' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// DialogClassStruct creates an uninitialised DialogClass.
func DialogClassStruct() *DialogClass {
	err := dialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DialogClass{native: dialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeDialogClass)
	return structGo
}
func finalizeDialogClass(obj *DialogClass) {
	dialogClassStruct.Free(obj.native)
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

// DialogPrivateStruct creates an uninitialised DialogPrivate.
func DialogPrivateStruct() *DialogPrivate {
	err := dialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DialogPrivate{native: dialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeDialogPrivate)
	return structGo
}
func finalizeDialogPrivate(obj *DialogPrivate) {
	dialogPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *DrawingAreaClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(drawingAreaClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *DrawingAreaClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(drawingAreaClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// DrawingAreaClassStruct creates an uninitialised DrawingAreaClass.
func DrawingAreaClassStruct() *DrawingAreaClass {
	err := drawingAreaClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DrawingAreaClass{native: drawingAreaClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeDrawingAreaClass)
	return structGo
}
func finalizeDrawingAreaClass(obj *DrawingAreaClass) {
	drawingAreaClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'base_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'base_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'insert_text' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_text' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_text' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_text' : for field setter : missing Type

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'do_insert_text' : for field getter : missing Type

// UNSUPPORTED : C value 'do_insert_text' : for field setter : missing Type

// UNSUPPORTED : C value 'do_delete_text' : for field getter : missing Type

// UNSUPPORTED : C value 'do_delete_text' : for field setter : missing Type

// UNSUPPORTED : C value 'get_chars' : for field getter : missing Type

// UNSUPPORTED : C value 'get_chars' : for field setter : missing Type

// UNSUPPORTED : C value 'set_selection_bounds' : for field getter : missing Type

// UNSUPPORTED : C value 'set_selection_bounds' : for field setter : missing Type

// UNSUPPORTED : C value 'get_selection_bounds' : for field getter : missing Type

// UNSUPPORTED : C value 'get_selection_bounds' : for field setter : missing Type

// UNSUPPORTED : C value 'set_position' : for field getter : missing Type

// UNSUPPORTED : C value 'set_position' : for field setter : missing Type

// UNSUPPORTED : C value 'get_position' : for field getter : missing Type

// UNSUPPORTED : C value 'get_position' : for field setter : missing Type

// EditableInterfaceStruct creates an uninitialised EditableInterface.
func EditableInterfaceStruct() *EditableInterface {
	err := editableInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EditableInterface{native: editableInterfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEditableInterface)
	return structGo
}
func finalizeEditableInterface(obj *EditableInterface) {
	editableInterfaceStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *EntryAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(entryAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *EntryAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(entryAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// EntryAccessibleClassStruct creates an uninitialised EntryAccessibleClass.
func EntryAccessibleClassStruct() *EntryAccessibleClass {
	err := entryAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EntryAccessibleClass{native: entryAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEntryAccessibleClass)
	return structGo
}
func finalizeEntryAccessibleClass(obj *EntryAccessibleClass) {
	entryAccessibleClassStruct.Free(obj.native)
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

// EntryAccessiblePrivateStruct creates an uninitialised EntryAccessiblePrivate.
func EntryAccessiblePrivateStruct() *EntryAccessiblePrivate {
	err := entryAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EntryAccessiblePrivate{native: entryAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEntryAccessiblePrivate)
	return structGo
}
func finalizeEntryAccessiblePrivate(obj *EntryAccessiblePrivate) {
	entryAccessiblePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'inserted_text' : for field getter : missing Type

// UNSUPPORTED : C value 'inserted_text' : for field setter : missing Type

// UNSUPPORTED : C value 'deleted_text' : for field getter : missing Type

// UNSUPPORTED : C value 'deleted_text' : for field setter : missing Type

// UNSUPPORTED : C value 'get_text' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text' : for field setter : missing Type

// UNSUPPORTED : C value 'get_length' : for field getter : missing Type

// UNSUPPORTED : C value 'get_length' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_text' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_text' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_text' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_text' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// EntryBufferClassStruct creates an uninitialised EntryBufferClass.
func EntryBufferClassStruct() *EntryBufferClass {
	err := entryBufferClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EntryBufferClass{native: entryBufferClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEntryBufferClass)
	return structGo
}
func finalizeEntryBufferClass(obj *EntryBufferClass) {
	entryBufferClassStruct.Free(obj.native)
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

// EntryBufferPrivateStruct creates an uninitialised EntryBufferPrivate.
func EntryBufferPrivateStruct() *EntryBufferPrivate {
	err := entryBufferPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EntryBufferPrivate{native: entryBufferPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEntryBufferPrivate)
	return structGo
}
func finalizeEntryBufferPrivate(obj *EntryBufferPrivate) {
	entryBufferPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *EntryClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(entryClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *EntryClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(entryClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'populate_popup' : for field getter : missing Type

// UNSUPPORTED : C value 'populate_popup' : for field setter : missing Type

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_at_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_at_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_from_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_from_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'backspace' : for field getter : missing Type

// UNSUPPORTED : C value 'backspace' : for field setter : missing Type

// UNSUPPORTED : C value 'cut_clipboard' : for field getter : missing Type

// UNSUPPORTED : C value 'cut_clipboard' : for field setter : missing Type

// UNSUPPORTED : C value 'copy_clipboard' : for field getter : missing Type

// UNSUPPORTED : C value 'copy_clipboard' : for field setter : missing Type

// UNSUPPORTED : C value 'paste_clipboard' : for field getter : missing Type

// UNSUPPORTED : C value 'paste_clipboard' : for field setter : missing Type

// UNSUPPORTED : C value 'toggle_overwrite' : for field getter : missing Type

// UNSUPPORTED : C value 'toggle_overwrite' : for field setter : missing Type

// UNSUPPORTED : C value 'get_text_area_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_area_size' : for field setter : missing Type

// UNSUPPORTED : C value 'get_frame_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_frame_size' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_emoji' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_emoji' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// EntryClassStruct creates an uninitialised EntryClass.
func EntryClassStruct() *EntryClass {
	err := entryClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EntryClass{native: entryClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEntryClass)
	return structGo
}
func finalizeEntryClass(obj *EntryClass) {
	entryClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'match_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'match_selected' : for field setter : missing Type

// UNSUPPORTED : C value 'action_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'action_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_prefix' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_prefix' : for field setter : missing Type

// UNSUPPORTED : C value 'cursor_on_match' : for field getter : missing Type

// UNSUPPORTED : C value 'cursor_on_match' : for field setter : missing Type

// UNSUPPORTED : C value 'no_matches' : for field getter : missing Type

// UNSUPPORTED : C value 'no_matches' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved0' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved0' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// EntryCompletionClassStruct creates an uninitialised EntryCompletionClass.
func EntryCompletionClassStruct() *EntryCompletionClass {
	err := entryCompletionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EntryCompletionClass{native: entryCompletionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEntryCompletionClass)
	return structGo
}
func finalizeEntryCompletionClass(obj *EntryCompletionClass) {
	entryCompletionClassStruct.Free(obj.native)
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

// EntryCompletionPrivateStruct creates an uninitialised EntryCompletionPrivate.
func EntryCompletionPrivateStruct() *EntryCompletionPrivate {
	err := entryCompletionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EntryCompletionPrivate{native: entryCompletionPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEntryCompletionPrivate)
	return structGo
}
func finalizeEntryCompletionPrivate(obj *EntryCompletionPrivate) {
	entryCompletionPrivateStruct.Free(obj.native)
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

// EntryPrivateStruct creates an uninitialised EntryPrivate.
func EntryPrivateStruct() *EntryPrivate {
	err := entryPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EntryPrivate{native: entryPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEntryPrivate)
	return structGo
}
func finalizeEntryPrivate(obj *EntryPrivate) {
	entryPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *EventBoxClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(eventBoxClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *EventBoxClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(eventBoxClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// EventBoxClassStruct creates an uninitialised EventBoxClass.
func EventBoxClassStruct() *EventBoxClass {
	err := eventBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventBoxClass{native: eventBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventBoxClass)
	return structGo
}
func finalizeEventBoxClass(obj *EventBoxClass) {
	eventBoxClassStruct.Free(obj.native)
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

// EventBoxPrivateStruct creates an uninitialised EventBoxPrivate.
func EventBoxPrivateStruct() *EventBoxPrivate {
	err := eventBoxPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventBoxPrivate{native: eventBoxPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventBoxPrivate)
	return structGo
}
func finalizeEventBoxPrivate(obj *EventBoxPrivate) {
	eventBoxPrivateStruct.Free(obj.native)
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

// EventControllerClassStruct creates an uninitialised EventControllerClass.
func EventControllerClassStruct() *EventControllerClass {
	err := eventControllerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventControllerClass{native: eventControllerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventControllerClass)
	return structGo
}
func finalizeEventControllerClass(obj *EventControllerClass) {
	eventControllerClassStruct.Free(obj.native)
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

// EventControllerKeyClassStruct creates an uninitialised EventControllerKeyClass.
func EventControllerKeyClassStruct() *EventControllerKeyClass {
	err := eventControllerKeyClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventControllerKeyClass{native: eventControllerKeyClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventControllerKeyClass)
	return structGo
}
func finalizeEventControllerKeyClass(obj *EventControllerKeyClass) {
	eventControllerKeyClassStruct.Free(obj.native)
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

// EventControllerMotionClassStruct creates an uninitialised EventControllerMotionClass.
func EventControllerMotionClassStruct() *EventControllerMotionClass {
	err := eventControllerMotionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventControllerMotionClass{native: eventControllerMotionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventControllerMotionClass)
	return structGo
}
func finalizeEventControllerMotionClass(obj *EventControllerMotionClass) {
	eventControllerMotionClassStruct.Free(obj.native)
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

// EventControllerScrollClassStruct creates an uninitialised EventControllerScrollClass.
func EventControllerScrollClassStruct() *EventControllerScrollClass {
	err := eventControllerScrollClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &EventControllerScrollClass{native: eventControllerScrollClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeEventControllerScrollClass)
	return structGo
}
func finalizeEventControllerScrollClass(obj *EventControllerScrollClass) {
	eventControllerScrollClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ExpanderAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(expanderAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ExpanderAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(expanderAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ExpanderAccessibleClassStruct creates an uninitialised ExpanderAccessibleClass.
func ExpanderAccessibleClassStruct() *ExpanderAccessibleClass {
	err := expanderAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ExpanderAccessibleClass{native: expanderAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeExpanderAccessibleClass)
	return structGo
}
func finalizeExpanderAccessibleClass(obj *ExpanderAccessibleClass) {
	expanderAccessibleClassStruct.Free(obj.native)
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

// ExpanderAccessiblePrivateStruct creates an uninitialised ExpanderAccessiblePrivate.
func ExpanderAccessiblePrivateStruct() *ExpanderAccessiblePrivate {
	err := expanderAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ExpanderAccessiblePrivate{native: expanderAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeExpanderAccessiblePrivate)
	return structGo
}
func finalizeExpanderAccessiblePrivate(obj *ExpanderAccessiblePrivate) {
	expanderAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ExpanderClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(expanderClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ExpanderClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(expanderClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ExpanderClassStruct creates an uninitialised ExpanderClass.
func ExpanderClassStruct() *ExpanderClass {
	err := expanderClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ExpanderClass{native: expanderClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeExpanderClass)
	return structGo
}
func finalizeExpanderClass(obj *ExpanderClass) {
	expanderClassStruct.Free(obj.native)
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

// ExpanderPrivateStruct creates an uninitialised ExpanderPrivate.
func ExpanderPrivateStruct() *ExpanderPrivate {
	err := expanderPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ExpanderPrivate{native: expanderPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeExpanderPrivate)
	return structGo
}
func finalizeExpanderPrivate(obj *ExpanderPrivate) {
	expanderPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileChooserButtonClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(fileChooserButtonClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileChooserButtonClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fileChooserButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'file_set' : for field getter : missing Type

// UNSUPPORTED : C value 'file_set' : for field setter : missing Type

// UNSUPPORTED : C value '__gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '__gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '__gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '__gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '__gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '__gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '__gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '__gtk_reserved4' : for field setter : missing Type

// FileChooserButtonClassStruct creates an uninitialised FileChooserButtonClass.
func FileChooserButtonClassStruct() *FileChooserButtonClass {
	err := fileChooserButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileChooserButtonClass{native: fileChooserButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFileChooserButtonClass)
	return structGo
}
func finalizeFileChooserButtonClass(obj *FileChooserButtonClass) {
	fileChooserButtonClassStruct.Free(obj.native)
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

// FileChooserButtonPrivateStruct creates an uninitialised FileChooserButtonPrivate.
func FileChooserButtonPrivateStruct() *FileChooserButtonPrivate {
	err := fileChooserButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileChooserButtonPrivate{native: fileChooserButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFileChooserButtonPrivate)
	return structGo
}
func finalizeFileChooserButtonPrivate(obj *FileChooserButtonPrivate) {
	fileChooserButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileChooserDialogClass) FieldParentClass() *DialogClass {
	argValue := gi.FieldGet(fileChooserDialogClassStruct, recv.native, "parent_class")
	value := &DialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileChooserDialogClass) SetFieldParentClass(value *DialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fileChooserDialogClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// FileChooserDialogClassStruct creates an uninitialised FileChooserDialogClass.
func FileChooserDialogClassStruct() *FileChooserDialogClass {
	err := fileChooserDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileChooserDialogClass{native: fileChooserDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFileChooserDialogClass)
	return structGo
}
func finalizeFileChooserDialogClass(obj *FileChooserDialogClass) {
	fileChooserDialogClassStruct.Free(obj.native)
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

// FileChooserDialogPrivateStruct creates an uninitialised FileChooserDialogPrivate.
func FileChooserDialogPrivateStruct() *FileChooserDialogPrivate {
	err := fileChooserDialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileChooserDialogPrivate{native: fileChooserDialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFileChooserDialogPrivate)
	return structGo
}
func finalizeFileChooserDialogPrivate(obj *FileChooserDialogPrivate) {
	fileChooserDialogPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileChooserNativeClass) FieldParentClass() *NativeDialogClass {
	argValue := gi.FieldGet(fileChooserNativeClassStruct, recv.native, "parent_class")
	value := &NativeDialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileChooserNativeClass) SetFieldParentClass(value *NativeDialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fileChooserNativeClassStruct, recv.native, "parent_class", argValue)
}

// FileChooserNativeClassStruct creates an uninitialised FileChooserNativeClass.
func FileChooserNativeClassStruct() *FileChooserNativeClass {
	err := fileChooserNativeClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileChooserNativeClass{native: fileChooserNativeClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFileChooserNativeClass)
	return structGo
}
func finalizeFileChooserNativeClass(obj *FileChooserNativeClass) {
	fileChooserNativeClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FileChooserWidgetClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(fileChooserWidgetClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FileChooserWidgetClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fileChooserWidgetClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// FileChooserWidgetClassStruct creates an uninitialised FileChooserWidgetClass.
func FileChooserWidgetClassStruct() *FileChooserWidgetClass {
	err := fileChooserWidgetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileChooserWidgetClass{native: fileChooserWidgetClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFileChooserWidgetClass)
	return structGo
}
func finalizeFileChooserWidgetClass(obj *FileChooserWidgetClass) {
	fileChooserWidgetClassStruct.Free(obj.native)
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

// FileChooserWidgetPrivateStruct creates an uninitialised FileChooserWidgetPrivate.
func FileChooserWidgetPrivateStruct() *FileChooserWidgetPrivate {
	err := fileChooserWidgetPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileChooserWidgetPrivate{native: fileChooserWidgetPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFileChooserWidgetPrivate)
	return structGo
}
func finalizeFileChooserWidgetPrivate(obj *FileChooserWidgetPrivate) {
	fileChooserWidgetPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'contains' : for field getter : no Go type for 'FileFilterFlags'

// UNSUPPORTED : C value 'contains' : for field setter : no Go type for 'FileFilterFlags'

// FieldFilename returns the C field 'filename'.
func (recv *FileFilterInfo) FieldFilename() string {
	argValue := gi.FieldGet(fileFilterInfoStruct, recv.native, "filename")
	value := argValue.String(false)
	return value
}

// SetFieldFilename sets the value of the C field 'filename'.
func (recv *FileFilterInfo) SetFieldFilename(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(fileFilterInfoStruct, recv.native, "filename", argValue)
}

// FieldUri returns the C field 'uri'.
func (recv *FileFilterInfo) FieldUri() string {
	argValue := gi.FieldGet(fileFilterInfoStruct, recv.native, "uri")
	value := argValue.String(false)
	return value
}

// SetFieldUri sets the value of the C field 'uri'.
func (recv *FileFilterInfo) SetFieldUri(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(fileFilterInfoStruct, recv.native, "uri", argValue)
}

// FieldDisplayName returns the C field 'display_name'.
func (recv *FileFilterInfo) FieldDisplayName() string {
	argValue := gi.FieldGet(fileFilterInfoStruct, recv.native, "display_name")
	value := argValue.String(false)
	return value
}

// SetFieldDisplayName sets the value of the C field 'display_name'.
func (recv *FileFilterInfo) SetFieldDisplayName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(fileFilterInfoStruct, recv.native, "display_name", argValue)
}

// FieldMimeType returns the C field 'mime_type'.
func (recv *FileFilterInfo) FieldMimeType() string {
	argValue := gi.FieldGet(fileFilterInfoStruct, recv.native, "mime_type")
	value := argValue.String(false)
	return value
}

// SetFieldMimeType sets the value of the C field 'mime_type'.
func (recv *FileFilterInfo) SetFieldMimeType(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(fileFilterInfoStruct, recv.native, "mime_type", argValue)
}

// FileFilterInfoStruct creates an uninitialised FileFilterInfo.
func FileFilterInfoStruct() *FileFilterInfo {
	err := fileFilterInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FileFilterInfo{native: fileFilterInfoStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFileFilterInfo)
	return structGo
}
func finalizeFileFilterInfo(obj *FileFilterInfo) {
	fileFilterInfoStruct.Free(obj.native)
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

// FieldWidget returns the C field 'widget'.
func (recv *FixedChild) FieldWidget() *Widget {
	argValue := gi.FieldGet(fixedChildStruct, recv.native, "widget")
	value := &Widget{native: argValue.Pointer()}
	return value
}

// SetFieldWidget sets the value of the C field 'widget'.
func (recv *FixedChild) SetFieldWidget(value *Widget) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fixedChildStruct, recv.native, "widget", argValue)
}

// FieldX returns the C field 'x'.
func (recv *FixedChild) FieldX() int32 {
	argValue := gi.FieldGet(fixedChildStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *FixedChild) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(fixedChildStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *FixedChild) FieldY() int32 {
	argValue := gi.FieldGet(fixedChildStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *FixedChild) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(fixedChildStruct, recv.native, "y", argValue)
}

// FixedChildStruct creates an uninitialised FixedChild.
func FixedChildStruct() *FixedChild {
	err := fixedChildStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FixedChild{native: fixedChildStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFixedChild)
	return structGo
}
func finalizeFixedChild(obj *FixedChild) {
	fixedChildStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FixedClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(fixedClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FixedClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fixedClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// FixedClassStruct creates an uninitialised FixedClass.
func FixedClassStruct() *FixedClass {
	err := fixedClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FixedClass{native: fixedClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFixedClass)
	return structGo
}
func finalizeFixedClass(obj *FixedClass) {
	fixedClassStruct.Free(obj.native)
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

// FixedPrivateStruct creates an uninitialised FixedPrivate.
func FixedPrivateStruct() *FixedPrivate {
	err := fixedPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FixedPrivate{native: fixedPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFixedPrivate)
	return structGo
}
func finalizeFixedPrivate(obj *FixedPrivate) {
	fixedPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FlowBoxAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(flowBoxAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FlowBoxAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(flowBoxAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// FlowBoxAccessibleClassStruct creates an uninitialised FlowBoxAccessibleClass.
func FlowBoxAccessibleClassStruct() *FlowBoxAccessibleClass {
	err := flowBoxAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FlowBoxAccessibleClass{native: flowBoxAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFlowBoxAccessibleClass)
	return structGo
}
func finalizeFlowBoxAccessibleClass(obj *FlowBoxAccessibleClass) {
	flowBoxAccessibleClassStruct.Free(obj.native)
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

// FlowBoxAccessiblePrivateStruct creates an uninitialised FlowBoxAccessiblePrivate.
func FlowBoxAccessiblePrivateStruct() *FlowBoxAccessiblePrivate {
	err := flowBoxAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FlowBoxAccessiblePrivate{native: flowBoxAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFlowBoxAccessiblePrivate)
	return structGo
}
func finalizeFlowBoxAccessiblePrivate(obj *FlowBoxAccessiblePrivate) {
	flowBoxAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FlowBoxChildAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(flowBoxChildAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FlowBoxChildAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(flowBoxChildAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// FlowBoxChildAccessibleClassStruct creates an uninitialised FlowBoxChildAccessibleClass.
func FlowBoxChildAccessibleClassStruct() *FlowBoxChildAccessibleClass {
	err := flowBoxChildAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FlowBoxChildAccessibleClass{native: flowBoxChildAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFlowBoxChildAccessibleClass)
	return structGo
}
func finalizeFlowBoxChildAccessibleClass(obj *FlowBoxChildAccessibleClass) {
	flowBoxChildAccessibleClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FlowBoxChildClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(flowBoxChildClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FlowBoxChildClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(flowBoxChildClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// FlowBoxChildClassStruct creates an uninitialised FlowBoxChildClass.
func FlowBoxChildClassStruct() *FlowBoxChildClass {
	err := flowBoxChildClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FlowBoxChildClass{native: flowBoxChildClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFlowBoxChildClass)
	return structGo
}
func finalizeFlowBoxChildClass(obj *FlowBoxChildClass) {
	flowBoxChildClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FlowBoxClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(flowBoxClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FlowBoxClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(flowBoxClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'child_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'child_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'selected_children_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'selected_children_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_cursor_child' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_cursor_child' : for field setter : missing Type

// UNSUPPORTED : C value 'toggle_cursor_child' : for field getter : missing Type

// UNSUPPORTED : C value 'toggle_cursor_child' : for field setter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'select_all' : for field getter : missing Type

// UNSUPPORTED : C value 'select_all' : for field setter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field getter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// FlowBoxClassStruct creates an uninitialised FlowBoxClass.
func FlowBoxClassStruct() *FlowBoxClass {
	err := flowBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FlowBoxClass{native: flowBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFlowBoxClass)
	return structGo
}
func finalizeFlowBoxClass(obj *FlowBoxClass) {
	flowBoxClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FontButtonClass) FieldParentClass() *ButtonClass {
	argValue := gi.FieldGet(fontButtonClassStruct, recv.native, "parent_class")
	value := &ButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FontButtonClass) SetFieldParentClass(value *ButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fontButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'font_set' : for field getter : missing Type

// UNSUPPORTED : C value 'font_set' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// FontButtonClassStruct creates an uninitialised FontButtonClass.
func FontButtonClassStruct() *FontButtonClass {
	err := fontButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontButtonClass{native: fontButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontButtonClass)
	return structGo
}
func finalizeFontButtonClass(obj *FontButtonClass) {
	fontButtonClassStruct.Free(obj.native)
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

// FontButtonPrivateStruct creates an uninitialised FontButtonPrivate.
func FontButtonPrivateStruct() *FontButtonPrivate {
	err := fontButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontButtonPrivate{native: fontButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontButtonPrivate)
	return structGo
}
func finalizeFontButtonPrivate(obj *FontButtonPrivate) {
	fontButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FontChooserDialogClass) FieldParentClass() *DialogClass {
	argValue := gi.FieldGet(fontChooserDialogClassStruct, recv.native, "parent_class")
	value := &DialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FontChooserDialogClass) SetFieldParentClass(value *DialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fontChooserDialogClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// FontChooserDialogClassStruct creates an uninitialised FontChooserDialogClass.
func FontChooserDialogClassStruct() *FontChooserDialogClass {
	err := fontChooserDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontChooserDialogClass{native: fontChooserDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontChooserDialogClass)
	return structGo
}
func finalizeFontChooserDialogClass(obj *FontChooserDialogClass) {
	fontChooserDialogClassStruct.Free(obj.native)
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

// FontChooserDialogPrivateStruct creates an uninitialised FontChooserDialogPrivate.
func FontChooserDialogPrivateStruct() *FontChooserDialogPrivate {
	err := fontChooserDialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontChooserDialogPrivate{native: fontChooserDialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontChooserDialogPrivate)
	return structGo
}
func finalizeFontChooserDialogPrivate(obj *FontChooserDialogPrivate) {
	fontChooserDialogPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'base_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'base_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_font_family' : for field getter : missing Type

// UNSUPPORTED : C value 'get_font_family' : for field setter : missing Type

// UNSUPPORTED : C value 'get_font_face' : for field getter : missing Type

// UNSUPPORTED : C value 'get_font_face' : for field setter : missing Type

// UNSUPPORTED : C value 'get_font_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_font_size' : for field setter : missing Type

// UNSUPPORTED : C value 'set_filter_func' : for field getter : missing Type

// UNSUPPORTED : C value 'set_filter_func' : for field setter : missing Type

// UNSUPPORTED : C value 'font_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'font_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'set_font_map' : for field getter : missing Type

// UNSUPPORTED : C value 'set_font_map' : for field setter : missing Type

// UNSUPPORTED : C value 'get_font_map' : for field getter : missing Type

// UNSUPPORTED : C value 'get_font_map' : for field setter : missing Type

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// FontChooserIfaceStruct creates an uninitialised FontChooserIface.
func FontChooserIfaceStruct() *FontChooserIface {
	err := fontChooserIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontChooserIface{native: fontChooserIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontChooserIface)
	return structGo
}
func finalizeFontChooserIface(obj *FontChooserIface) {
	fontChooserIfaceStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FontChooserWidgetClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(fontChooserWidgetClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FontChooserWidgetClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fontChooserWidgetClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// FontChooserWidgetClassStruct creates an uninitialised FontChooserWidgetClass.
func FontChooserWidgetClassStruct() *FontChooserWidgetClass {
	err := fontChooserWidgetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontChooserWidgetClass{native: fontChooserWidgetClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontChooserWidgetClass)
	return structGo
}
func finalizeFontChooserWidgetClass(obj *FontChooserWidgetClass) {
	fontChooserWidgetClassStruct.Free(obj.native)
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

// FontChooserWidgetPrivateStruct creates an uninitialised FontChooserWidgetPrivate.
func FontChooserWidgetPrivateStruct() *FontChooserWidgetPrivate {
	err := fontChooserWidgetPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontChooserWidgetPrivate{native: fontChooserWidgetPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontChooserWidgetPrivate)
	return structGo
}
func finalizeFontChooserWidgetPrivate(obj *FontChooserWidgetPrivate) {
	fontChooserWidgetPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FontSelectionClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(fontSelectionClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FontSelectionClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fontSelectionClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// FontSelectionClassStruct creates an uninitialised FontSelectionClass.
func FontSelectionClassStruct() *FontSelectionClass {
	err := fontSelectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontSelectionClass{native: fontSelectionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontSelectionClass)
	return structGo
}
func finalizeFontSelectionClass(obj *FontSelectionClass) {
	fontSelectionClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FontSelectionDialogClass) FieldParentClass() *DialogClass {
	argValue := gi.FieldGet(fontSelectionDialogClassStruct, recv.native, "parent_class")
	value := &DialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FontSelectionDialogClass) SetFieldParentClass(value *DialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(fontSelectionDialogClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// FontSelectionDialogClassStruct creates an uninitialised FontSelectionDialogClass.
func FontSelectionDialogClassStruct() *FontSelectionDialogClass {
	err := fontSelectionDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontSelectionDialogClass{native: fontSelectionDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontSelectionDialogClass)
	return structGo
}
func finalizeFontSelectionDialogClass(obj *FontSelectionDialogClass) {
	fontSelectionDialogClassStruct.Free(obj.native)
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

// FontSelectionDialogPrivateStruct creates an uninitialised FontSelectionDialogPrivate.
func FontSelectionDialogPrivateStruct() *FontSelectionDialogPrivate {
	err := fontSelectionDialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontSelectionDialogPrivate{native: fontSelectionDialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontSelectionDialogPrivate)
	return structGo
}
func finalizeFontSelectionDialogPrivate(obj *FontSelectionDialogPrivate) {
	fontSelectionDialogPrivateStruct.Free(obj.native)
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

// FontSelectionPrivateStruct creates an uninitialised FontSelectionPrivate.
func FontSelectionPrivateStruct() *FontSelectionPrivate {
	err := fontSelectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FontSelectionPrivate{native: fontSelectionPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFontSelectionPrivate)
	return structGo
}
func finalizeFontSelectionPrivate(obj *FontSelectionPrivate) {
	fontSelectionPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FrameAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(frameAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FrameAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(frameAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// FrameAccessibleClassStruct creates an uninitialised FrameAccessibleClass.
func FrameAccessibleClassStruct() *FrameAccessibleClass {
	err := frameAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FrameAccessibleClass{native: frameAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFrameAccessibleClass)
	return structGo
}
func finalizeFrameAccessibleClass(obj *FrameAccessibleClass) {
	frameAccessibleClassStruct.Free(obj.native)
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

// FrameAccessiblePrivateStruct creates an uninitialised FrameAccessiblePrivate.
func FrameAccessiblePrivateStruct() *FrameAccessiblePrivate {
	err := frameAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FrameAccessiblePrivate{native: frameAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFrameAccessiblePrivate)
	return structGo
}
func finalizeFrameAccessiblePrivate(obj *FrameAccessiblePrivate) {
	frameAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *FrameClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(frameClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *FrameClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(frameClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'compute_child_allocation' : for field getter : missing Type

// UNSUPPORTED : C value 'compute_child_allocation' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// FrameClassStruct creates an uninitialised FrameClass.
func FrameClassStruct() *FrameClass {
	err := frameClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FrameClass{native: frameClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFrameClass)
	return structGo
}
func finalizeFrameClass(obj *FrameClass) {
	frameClassStruct.Free(obj.native)
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

// FramePrivateStruct creates an uninitialised FramePrivate.
func FramePrivateStruct() *FramePrivate {
	err := framePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FramePrivate{native: framePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeFramePrivate)
	return structGo
}
func finalizeFramePrivate(obj *FramePrivate) {
	framePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'render' : for field getter : missing Type

// UNSUPPORTED : C value 'render' : for field setter : missing Type

// UNSUPPORTED : C value 'resize' : for field getter : missing Type

// UNSUPPORTED : C value 'resize' : for field setter : missing Type

// UNSUPPORTED : C value 'create_context' : for field getter : missing Type

// UNSUPPORTED : C value 'create_context' : for field setter : missing Type

// GLAreaClassStruct creates an uninitialised GLAreaClass.
func GLAreaClassStruct() *GLAreaClass {
	err := gLAreaClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GLAreaClass{native: gLAreaClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGLAreaClass)
	return structGo
}
func finalizeGLAreaClass(obj *GLAreaClass) {
	gLAreaClassStruct.Free(obj.native)
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

// GestureClassStruct creates an uninitialised GestureClass.
func GestureClassStruct() *GestureClass {
	err := gestureClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GestureClass{native: gestureClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGestureClass)
	return structGo
}
func finalizeGestureClass(obj *GestureClass) {
	gestureClassStruct.Free(obj.native)
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

// GestureDragClassStruct creates an uninitialised GestureDragClass.
func GestureDragClassStruct() *GestureDragClass {
	err := gestureDragClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GestureDragClass{native: gestureDragClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGestureDragClass)
	return structGo
}
func finalizeGestureDragClass(obj *GestureDragClass) {
	gestureDragClassStruct.Free(obj.native)
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

// GestureLongPressClassStruct creates an uninitialised GestureLongPressClass.
func GestureLongPressClassStruct() *GestureLongPressClass {
	err := gestureLongPressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GestureLongPressClass{native: gestureLongPressClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGestureLongPressClass)
	return structGo
}
func finalizeGestureLongPressClass(obj *GestureLongPressClass) {
	gestureLongPressClassStruct.Free(obj.native)
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

// GestureMultiPressClassStruct creates an uninitialised GestureMultiPressClass.
func GestureMultiPressClassStruct() *GestureMultiPressClass {
	err := gestureMultiPressClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GestureMultiPressClass{native: gestureMultiPressClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGestureMultiPressClass)
	return structGo
}
func finalizeGestureMultiPressClass(obj *GestureMultiPressClass) {
	gestureMultiPressClassStruct.Free(obj.native)
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

// GesturePanClassStruct creates an uninitialised GesturePanClass.
func GesturePanClassStruct() *GesturePanClass {
	err := gesturePanClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GesturePanClass{native: gesturePanClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGesturePanClass)
	return structGo
}
func finalizeGesturePanClass(obj *GesturePanClass) {
	gesturePanClassStruct.Free(obj.native)
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

// GestureRotateClassStruct creates an uninitialised GestureRotateClass.
func GestureRotateClassStruct() *GestureRotateClass {
	err := gestureRotateClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GestureRotateClass{native: gestureRotateClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGestureRotateClass)
	return structGo
}
func finalizeGestureRotateClass(obj *GestureRotateClass) {
	gestureRotateClassStruct.Free(obj.native)
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

// GestureSingleClassStruct creates an uninitialised GestureSingleClass.
func GestureSingleClassStruct() *GestureSingleClass {
	err := gestureSingleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GestureSingleClass{native: gestureSingleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGestureSingleClass)
	return structGo
}
func finalizeGestureSingleClass(obj *GestureSingleClass) {
	gestureSingleClassStruct.Free(obj.native)
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

// GestureStylusClassStruct creates an uninitialised GestureStylusClass.
func GestureStylusClassStruct() *GestureStylusClass {
	err := gestureStylusClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GestureStylusClass{native: gestureStylusClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGestureStylusClass)
	return structGo
}
func finalizeGestureStylusClass(obj *GestureStylusClass) {
	gestureStylusClassStruct.Free(obj.native)
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

// GestureSwipeClassStruct creates an uninitialised GestureSwipeClass.
func GestureSwipeClassStruct() *GestureSwipeClass {
	err := gestureSwipeClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GestureSwipeClass{native: gestureSwipeClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGestureSwipeClass)
	return structGo
}
func finalizeGestureSwipeClass(obj *GestureSwipeClass) {
	gestureSwipeClassStruct.Free(obj.native)
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

// GestureZoomClassStruct creates an uninitialised GestureZoomClass.
func GestureZoomClassStruct() *GestureZoomClass {
	err := gestureZoomClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GestureZoomClass{native: gestureZoomClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGestureZoomClass)
	return structGo
}
func finalizeGestureZoomClass(obj *GestureZoomClass) {
	gestureZoomClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'gtk_gradient_resolve' : parameter 'resolved_gradient' of type 'cairo.Pattern' not supported

// UNSUPPORTED : C value 'gtk_gradient_resolve_for_context' : return type 'cairo.Pattern' not supported

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

// FieldParentClass returns the C field 'parent_class'.
func (recv *GridClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(gridClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *GridClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(gridClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// GridClassStruct creates an uninitialised GridClass.
func GridClassStruct() *GridClass {
	err := gridClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GridClass{native: gridClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGridClass)
	return structGo
}
func finalizeGridClass(obj *GridClass) {
	gridClassStruct.Free(obj.native)
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

// GridPrivateStruct creates an uninitialised GridPrivate.
func GridPrivateStruct() *GridPrivate {
	err := gridPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GridPrivate{native: gridPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeGridPrivate)
	return structGo
}
func finalizeGridPrivate(obj *GridPrivate) {
	gridPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HBoxClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(hBoxClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HBoxClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hBoxClassStruct, recv.native, "parent_class", argValue)
}

// HBoxClassStruct creates an uninitialised HBoxClass.
func HBoxClassStruct() *HBoxClass {
	err := hBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HBoxClass{native: hBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHBoxClass)
	return structGo
}
func finalizeHBoxClass(obj *HBoxClass) {
	hBoxClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HButtonBoxClass) FieldParentClass() *ButtonBoxClass {
	argValue := gi.FieldGet(hButtonBoxClassStruct, recv.native, "parent_class")
	value := &ButtonBoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HButtonBoxClass) SetFieldParentClass(value *ButtonBoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hButtonBoxClassStruct, recv.native, "parent_class", argValue)
}

// HButtonBoxClassStruct creates an uninitialised HButtonBoxClass.
func HButtonBoxClassStruct() *HButtonBoxClass {
	err := hButtonBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HButtonBoxClass{native: hButtonBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHButtonBoxClass)
	return structGo
}
func finalizeHButtonBoxClass(obj *HButtonBoxClass) {
	hButtonBoxClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HPanedClass) FieldParentClass() *PanedClass {
	argValue := gi.FieldGet(hPanedClassStruct, recv.native, "parent_class")
	value := &PanedClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HPanedClass) SetFieldParentClass(value *PanedClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hPanedClassStruct, recv.native, "parent_class", argValue)
}

// HPanedClassStruct creates an uninitialised HPanedClass.
func HPanedClassStruct() *HPanedClass {
	err := hPanedClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HPanedClass{native: hPanedClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHPanedClass)
	return structGo
}
func finalizeHPanedClass(obj *HPanedClass) {
	hPanedClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HSVClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(hSVClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HSVClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hSVClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'move' : for field getter : missing Type

// UNSUPPORTED : C value 'move' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// HSVClassStruct creates an uninitialised HSVClass.
func HSVClassStruct() *HSVClass {
	err := hSVClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HSVClass{native: hSVClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHSVClass)
	return structGo
}
func finalizeHSVClass(obj *HSVClass) {
	hSVClassStruct.Free(obj.native)
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

// HSVPrivateStruct creates an uninitialised HSVPrivate.
func HSVPrivateStruct() *HSVPrivate {
	err := hSVPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HSVPrivate{native: hSVPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHSVPrivate)
	return structGo
}
func finalizeHSVPrivate(obj *HSVPrivate) {
	hSVPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HScaleClass) FieldParentClass() *ScaleClass {
	argValue := gi.FieldGet(hScaleClassStruct, recv.native, "parent_class")
	value := &ScaleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HScaleClass) SetFieldParentClass(value *ScaleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hScaleClassStruct, recv.native, "parent_class", argValue)
}

// HScaleClassStruct creates an uninitialised HScaleClass.
func HScaleClassStruct() *HScaleClass {
	err := hScaleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HScaleClass{native: hScaleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHScaleClass)
	return structGo
}
func finalizeHScaleClass(obj *HScaleClass) {
	hScaleClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HScrollbarClass) FieldParentClass() *ScrollbarClass {
	argValue := gi.FieldGet(hScrollbarClassStruct, recv.native, "parent_class")
	value := &ScrollbarClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HScrollbarClass) SetFieldParentClass(value *ScrollbarClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hScrollbarClassStruct, recv.native, "parent_class", argValue)
}

// HScrollbarClassStruct creates an uninitialised HScrollbarClass.
func HScrollbarClassStruct() *HScrollbarClass {
	err := hScrollbarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HScrollbarClass{native: hScrollbarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHScrollbarClass)
	return structGo
}
func finalizeHScrollbarClass(obj *HScrollbarClass) {
	hScrollbarClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HSeparatorClass) FieldParentClass() *SeparatorClass {
	argValue := gi.FieldGet(hSeparatorClassStruct, recv.native, "parent_class")
	value := &SeparatorClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HSeparatorClass) SetFieldParentClass(value *SeparatorClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(hSeparatorClassStruct, recv.native, "parent_class", argValue)
}

// HSeparatorClassStruct creates an uninitialised HSeparatorClass.
func HSeparatorClassStruct() *HSeparatorClass {
	err := hSeparatorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HSeparatorClass{native: hSeparatorClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHSeparatorClass)
	return structGo
}
func finalizeHSeparatorClass(obj *HSeparatorClass) {
	hSeparatorClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HandleBoxClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(handleBoxClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HandleBoxClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(handleBoxClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'child_attached' : for field getter : missing Type

// UNSUPPORTED : C value 'child_attached' : for field setter : missing Type

// UNSUPPORTED : C value 'child_detached' : for field getter : missing Type

// UNSUPPORTED : C value 'child_detached' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// HandleBoxClassStruct creates an uninitialised HandleBoxClass.
func HandleBoxClassStruct() *HandleBoxClass {
	err := handleBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HandleBoxClass{native: handleBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHandleBoxClass)
	return structGo
}
func finalizeHandleBoxClass(obj *HandleBoxClass) {
	handleBoxClassStruct.Free(obj.native)
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

// HandleBoxPrivateStruct creates an uninitialised HandleBoxPrivate.
func HandleBoxPrivateStruct() *HandleBoxPrivate {
	err := handleBoxPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HandleBoxPrivate{native: handleBoxPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHandleBoxPrivate)
	return structGo
}
func finalizeHandleBoxPrivate(obj *HandleBoxPrivate) {
	handleBoxPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HeaderBarAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(headerBarAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HeaderBarAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(headerBarAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// HeaderBarAccessibleClassStruct creates an uninitialised HeaderBarAccessibleClass.
func HeaderBarAccessibleClassStruct() *HeaderBarAccessibleClass {
	err := headerBarAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HeaderBarAccessibleClass{native: headerBarAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHeaderBarAccessibleClass)
	return structGo
}
func finalizeHeaderBarAccessibleClass(obj *HeaderBarAccessibleClass) {
	headerBarAccessibleClassStruct.Free(obj.native)
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

// HeaderBarAccessiblePrivateStruct creates an uninitialised HeaderBarAccessiblePrivate.
func HeaderBarAccessiblePrivateStruct() *HeaderBarAccessiblePrivate {
	err := headerBarAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HeaderBarAccessiblePrivate{native: headerBarAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHeaderBarAccessiblePrivate)
	return structGo
}
func finalizeHeaderBarAccessiblePrivate(obj *HeaderBarAccessiblePrivate) {
	headerBarAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *HeaderBarClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(headerBarClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *HeaderBarClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(headerBarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// HeaderBarClassStruct creates an uninitialised HeaderBarClass.
func HeaderBarClassStruct() *HeaderBarClass {
	err := headerBarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HeaderBarClass{native: headerBarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHeaderBarClass)
	return structGo
}
func finalizeHeaderBarClass(obj *HeaderBarClass) {
	headerBarClassStruct.Free(obj.native)
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

// HeaderBarPrivateStruct creates an uninitialised HeaderBarPrivate.
func HeaderBarPrivateStruct() *HeaderBarPrivate {
	err := headerBarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &HeaderBarPrivate{native: headerBarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeHeaderBarPrivate)
	return structGo
}
func finalizeHeaderBarPrivate(obj *HeaderBarPrivate) {
	headerBarPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'preedit_start' : for field getter : missing Type

// UNSUPPORTED : C value 'preedit_start' : for field setter : missing Type

// UNSUPPORTED : C value 'preedit_end' : for field getter : missing Type

// UNSUPPORTED : C value 'preedit_end' : for field setter : missing Type

// UNSUPPORTED : C value 'preedit_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'preedit_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'commit' : for field getter : missing Type

// UNSUPPORTED : C value 'commit' : for field setter : missing Type

// UNSUPPORTED : C value 'retrieve_surrounding' : for field getter : missing Type

// UNSUPPORTED : C value 'retrieve_surrounding' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_surrounding' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_surrounding' : for field setter : missing Type

// UNSUPPORTED : C value 'set_client_window' : for field getter : missing Type

// UNSUPPORTED : C value 'set_client_window' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preedit_string' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preedit_string' : for field setter : missing Type

// UNSUPPORTED : C value 'filter_keypress' : for field getter : missing Type

// UNSUPPORTED : C value 'filter_keypress' : for field setter : missing Type

// UNSUPPORTED : C value 'focus_in' : for field getter : missing Type

// UNSUPPORTED : C value 'focus_in' : for field setter : missing Type

// UNSUPPORTED : C value 'focus_out' : for field getter : missing Type

// UNSUPPORTED : C value 'focus_out' : for field setter : missing Type

// UNSUPPORTED : C value 'reset' : for field getter : missing Type

// UNSUPPORTED : C value 'reset' : for field setter : missing Type

// UNSUPPORTED : C value 'set_cursor_location' : for field getter : missing Type

// UNSUPPORTED : C value 'set_cursor_location' : for field setter : missing Type

// UNSUPPORTED : C value 'set_use_preedit' : for field getter : missing Type

// UNSUPPORTED : C value 'set_use_preedit' : for field setter : missing Type

// UNSUPPORTED : C value 'set_surrounding' : for field getter : missing Type

// UNSUPPORTED : C value 'set_surrounding' : for field setter : missing Type

// UNSUPPORTED : C value 'get_surrounding' : for field getter : missing Type

// UNSUPPORTED : C value 'get_surrounding' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// IMContextClassStruct creates an uninitialised IMContextClass.
func IMContextClassStruct() *IMContextClass {
	err := iMContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IMContextClass{native: iMContextClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIMContextClass)
	return structGo
}
func finalizeIMContextClass(obj *IMContextClass) {
	iMContextClassStruct.Free(obj.native)
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

// FieldContextId returns the C field 'context_id'.
func (recv *IMContextInfo) FieldContextId() string {
	argValue := gi.FieldGet(iMContextInfoStruct, recv.native, "context_id")
	value := argValue.String(false)
	return value
}

// SetFieldContextId sets the value of the C field 'context_id'.
func (recv *IMContextInfo) SetFieldContextId(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(iMContextInfoStruct, recv.native, "context_id", argValue)
}

// FieldContextName returns the C field 'context_name'.
func (recv *IMContextInfo) FieldContextName() string {
	argValue := gi.FieldGet(iMContextInfoStruct, recv.native, "context_name")
	value := argValue.String(false)
	return value
}

// SetFieldContextName sets the value of the C field 'context_name'.
func (recv *IMContextInfo) SetFieldContextName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(iMContextInfoStruct, recv.native, "context_name", argValue)
}

// FieldDomain returns the C field 'domain'.
func (recv *IMContextInfo) FieldDomain() string {
	argValue := gi.FieldGet(iMContextInfoStruct, recv.native, "domain")
	value := argValue.String(false)
	return value
}

// SetFieldDomain sets the value of the C field 'domain'.
func (recv *IMContextInfo) SetFieldDomain(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(iMContextInfoStruct, recv.native, "domain", argValue)
}

// FieldDomainDirname returns the C field 'domain_dirname'.
func (recv *IMContextInfo) FieldDomainDirname() string {
	argValue := gi.FieldGet(iMContextInfoStruct, recv.native, "domain_dirname")
	value := argValue.String(false)
	return value
}

// SetFieldDomainDirname sets the value of the C field 'domain_dirname'.
func (recv *IMContextInfo) SetFieldDomainDirname(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(iMContextInfoStruct, recv.native, "domain_dirname", argValue)
}

// FieldDefaultLocales returns the C field 'default_locales'.
func (recv *IMContextInfo) FieldDefaultLocales() string {
	argValue := gi.FieldGet(iMContextInfoStruct, recv.native, "default_locales")
	value := argValue.String(false)
	return value
}

// SetFieldDefaultLocales sets the value of the C field 'default_locales'.
func (recv *IMContextInfo) SetFieldDefaultLocales(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(iMContextInfoStruct, recv.native, "default_locales", argValue)
}

// IMContextInfoStruct creates an uninitialised IMContextInfo.
func IMContextInfoStruct() *IMContextInfo {
	err := iMContextInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IMContextInfo{native: iMContextInfoStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIMContextInfo)
	return structGo
}
func finalizeIMContextInfo(obj *IMContextInfo) {
	iMContextInfoStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *IMContextSimpleClass) FieldParentClass() *IMContextClass {
	argValue := gi.FieldGet(iMContextSimpleClassStruct, recv.native, "parent_class")
	value := &IMContextClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *IMContextSimpleClass) SetFieldParentClass(value *IMContextClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(iMContextSimpleClassStruct, recv.native, "parent_class", argValue)
}

// IMContextSimpleClassStruct creates an uninitialised IMContextSimpleClass.
func IMContextSimpleClassStruct() *IMContextSimpleClass {
	err := iMContextSimpleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IMContextSimpleClass{native: iMContextSimpleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIMContextSimpleClass)
	return structGo
}
func finalizeIMContextSimpleClass(obj *IMContextSimpleClass) {
	iMContextSimpleClassStruct.Free(obj.native)
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

// IMContextSimplePrivateStruct creates an uninitialised IMContextSimplePrivate.
func IMContextSimplePrivateStruct() *IMContextSimplePrivate {
	err := iMContextSimplePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IMContextSimplePrivate{native: iMContextSimplePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIMContextSimplePrivate)
	return structGo
}
func finalizeIMContextSimplePrivate(obj *IMContextSimplePrivate) {
	iMContextSimplePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *IMMulticontextClass) FieldParentClass() *IMContextClass {
	argValue := gi.FieldGet(iMMulticontextClassStruct, recv.native, "parent_class")
	value := &IMContextClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *IMMulticontextClass) SetFieldParentClass(value *IMContextClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(iMMulticontextClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// IMMulticontextClassStruct creates an uninitialised IMMulticontextClass.
func IMMulticontextClassStruct() *IMMulticontextClass {
	err := iMMulticontextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IMMulticontextClass{native: iMMulticontextClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIMMulticontextClass)
	return structGo
}
func finalizeIMMulticontextClass(obj *IMMulticontextClass) {
	iMMulticontextClassStruct.Free(obj.native)
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

// IMMulticontextPrivateStruct creates an uninitialised IMMulticontextPrivate.
func IMMulticontextPrivateStruct() *IMMulticontextPrivate {
	err := iMMulticontextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IMMulticontextPrivate{native: iMMulticontextPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIMMulticontextPrivate)
	return structGo
}
func finalizeIMMulticontextPrivate(obj *IMMulticontextPrivate) {
	iMMulticontextPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// IconFactoryClassStruct creates an uninitialised IconFactoryClass.
func IconFactoryClassStruct() *IconFactoryClass {
	err := iconFactoryClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IconFactoryClass{native: iconFactoryClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIconFactoryClass)
	return structGo
}
func finalizeIconFactoryClass(obj *IconFactoryClass) {
	iconFactoryClassStruct.Free(obj.native)
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

// IconFactoryPrivateStruct creates an uninitialised IconFactoryPrivate.
func IconFactoryPrivateStruct() *IconFactoryPrivate {
	err := iconFactoryPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IconFactoryPrivate{native: iconFactoryPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIconFactoryPrivate)
	return structGo
}
func finalizeIconFactoryPrivate(obj *IconFactoryPrivate) {
	iconFactoryPrivateStruct.Free(obj.native)
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

// IconInfoClassStruct creates an uninitialised IconInfoClass.
func IconInfoClassStruct() *IconInfoClass {
	err := iconInfoClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IconInfoClass{native: iconInfoClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIconInfoClass)
	return structGo
}
func finalizeIconInfoClass(obj *IconInfoClass) {
	iconInfoClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'gtk_icon_set_get_sizes' : parameter 'sizes' of type 'nil' not supported

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

// UNSUPPORTED : C value 'gtk_icon_set_render_icon' : parameter 'direction' of type 'TextDirection' not supported

// UNSUPPORTED : C value 'gtk_icon_set_render_icon_pixbuf' : parameter 'size' of type 'IconSize' not supported

// UNSUPPORTED : C value 'gtk_icon_set_render_icon_surface' : parameter 'size' of type 'IconSize' not supported

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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// IconThemeClassStruct creates an uninitialised IconThemeClass.
func IconThemeClassStruct() *IconThemeClass {
	err := iconThemeClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IconThemeClass{native: iconThemeClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIconThemeClass)
	return structGo
}
func finalizeIconThemeClass(obj *IconThemeClass) {
	iconThemeClassStruct.Free(obj.native)
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

// IconThemePrivateStruct creates an uninitialised IconThemePrivate.
func IconThemePrivateStruct() *IconThemePrivate {
	err := iconThemePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IconThemePrivate{native: iconThemePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIconThemePrivate)
	return structGo
}
func finalizeIconThemePrivate(obj *IconThemePrivate) {
	iconThemePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *IconViewAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(iconViewAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *IconViewAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(iconViewAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// IconViewAccessibleClassStruct creates an uninitialised IconViewAccessibleClass.
func IconViewAccessibleClassStruct() *IconViewAccessibleClass {
	err := iconViewAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IconViewAccessibleClass{native: iconViewAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIconViewAccessibleClass)
	return structGo
}
func finalizeIconViewAccessibleClass(obj *IconViewAccessibleClass) {
	iconViewAccessibleClassStruct.Free(obj.native)
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

// IconViewAccessiblePrivateStruct creates an uninitialised IconViewAccessiblePrivate.
func IconViewAccessiblePrivateStruct() *IconViewAccessiblePrivate {
	err := iconViewAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IconViewAccessiblePrivate{native: iconViewAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIconViewAccessiblePrivate)
	return structGo
}
func finalizeIconViewAccessiblePrivate(obj *IconViewAccessiblePrivate) {
	iconViewAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *IconViewClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(iconViewClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *IconViewClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(iconViewClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'item_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'item_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'selection_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'select_all' : for field getter : missing Type

// UNSUPPORTED : C value 'select_all' : for field setter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field getter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field setter : missing Type

// UNSUPPORTED : C value 'select_cursor_item' : for field getter : missing Type

// UNSUPPORTED : C value 'select_cursor_item' : for field setter : missing Type

// UNSUPPORTED : C value 'toggle_cursor_item' : for field getter : missing Type

// UNSUPPORTED : C value 'toggle_cursor_item' : for field setter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_cursor_item' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_cursor_item' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// IconViewClassStruct creates an uninitialised IconViewClass.
func IconViewClassStruct() *IconViewClass {
	err := iconViewClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IconViewClass{native: iconViewClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIconViewClass)
	return structGo
}
func finalizeIconViewClass(obj *IconViewClass) {
	iconViewClassStruct.Free(obj.native)
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

// IconViewPrivateStruct creates an uninitialised IconViewPrivate.
func IconViewPrivateStruct() *IconViewPrivate {
	err := iconViewPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &IconViewPrivate{native: iconViewPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeIconViewPrivate)
	return structGo
}
func finalizeIconViewPrivate(obj *IconViewPrivate) {
	iconViewPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ImageAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(imageAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ImageAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(imageAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ImageAccessibleClassStruct creates an uninitialised ImageAccessibleClass.
func ImageAccessibleClassStruct() *ImageAccessibleClass {
	err := imageAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ImageAccessibleClass{native: imageAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImageAccessibleClass)
	return structGo
}
func finalizeImageAccessibleClass(obj *ImageAccessibleClass) {
	imageAccessibleClassStruct.Free(obj.native)
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

// ImageAccessiblePrivateStruct creates an uninitialised ImageAccessiblePrivate.
func ImageAccessiblePrivateStruct() *ImageAccessiblePrivate {
	err := imageAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ImageAccessiblePrivate{native: imageAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImageAccessiblePrivate)
	return structGo
}
func finalizeImageAccessiblePrivate(obj *ImageAccessiblePrivate) {
	imageAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ImageCellAccessibleClass) FieldParentClass() *RendererCellAccessibleClass {
	argValue := gi.FieldGet(imageCellAccessibleClassStruct, recv.native, "parent_class")
	value := &RendererCellAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ImageCellAccessibleClass) SetFieldParentClass(value *RendererCellAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(imageCellAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ImageCellAccessibleClassStruct creates an uninitialised ImageCellAccessibleClass.
func ImageCellAccessibleClassStruct() *ImageCellAccessibleClass {
	err := imageCellAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ImageCellAccessibleClass{native: imageCellAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImageCellAccessibleClass)
	return structGo
}
func finalizeImageCellAccessibleClass(obj *ImageCellAccessibleClass) {
	imageCellAccessibleClassStruct.Free(obj.native)
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

// ImageCellAccessiblePrivateStruct creates an uninitialised ImageCellAccessiblePrivate.
func ImageCellAccessiblePrivateStruct() *ImageCellAccessiblePrivate {
	err := imageCellAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ImageCellAccessiblePrivate{native: imageCellAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImageCellAccessiblePrivate)
	return structGo
}
func finalizeImageCellAccessiblePrivate(obj *ImageCellAccessiblePrivate) {
	imageCellAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ImageClass) FieldParentClass() *MiscClass {
	argValue := gi.FieldGet(imageClassStruct, recv.native, "parent_class")
	value := &MiscClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ImageClass) SetFieldParentClass(value *MiscClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(imageClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ImageClassStruct creates an uninitialised ImageClass.
func ImageClassStruct() *ImageClass {
	err := imageClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ImageClass{native: imageClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImageClass)
	return structGo
}
func finalizeImageClass(obj *ImageClass) {
	imageClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ImageMenuItemClass) FieldParentClass() *MenuItemClass {
	argValue := gi.FieldGet(imageMenuItemClassStruct, recv.native, "parent_class")
	value := &MenuItemClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ImageMenuItemClass) SetFieldParentClass(value *MenuItemClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(imageMenuItemClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ImageMenuItemClassStruct creates an uninitialised ImageMenuItemClass.
func ImageMenuItemClassStruct() *ImageMenuItemClass {
	err := imageMenuItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ImageMenuItemClass{native: imageMenuItemClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImageMenuItemClass)
	return structGo
}
func finalizeImageMenuItemClass(obj *ImageMenuItemClass) {
	imageMenuItemClassStruct.Free(obj.native)
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

// ImageMenuItemPrivateStruct creates an uninitialised ImageMenuItemPrivate.
func ImageMenuItemPrivateStruct() *ImageMenuItemPrivate {
	err := imageMenuItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ImageMenuItemPrivate{native: imageMenuItemPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImageMenuItemPrivate)
	return structGo
}
func finalizeImageMenuItemPrivate(obj *ImageMenuItemPrivate) {
	imageMenuItemPrivateStruct.Free(obj.native)
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

// ImagePrivateStruct creates an uninitialised ImagePrivate.
func ImagePrivateStruct() *ImagePrivate {
	err := imagePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ImagePrivate{native: imagePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeImagePrivate)
	return structGo
}
func finalizeImagePrivate(obj *ImagePrivate) {
	imagePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *InfoBarClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(infoBarClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *InfoBarClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(infoBarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'response' : for field getter : missing Type

// UNSUPPORTED : C value 'response' : for field setter : missing Type

// UNSUPPORTED : C value 'close' : for field getter : missing Type

// UNSUPPORTED : C value 'close' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// InfoBarClassStruct creates an uninitialised InfoBarClass.
func InfoBarClassStruct() *InfoBarClass {
	err := infoBarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InfoBarClass{native: infoBarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeInfoBarClass)
	return structGo
}
func finalizeInfoBarClass(obj *InfoBarClass) {
	infoBarClassStruct.Free(obj.native)
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

// InfoBarPrivateStruct creates an uninitialised InfoBarPrivate.
func InfoBarPrivateStruct() *InfoBarPrivate {
	err := infoBarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InfoBarPrivate{native: infoBarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeInfoBarPrivate)
	return structGo
}
func finalizeInfoBarPrivate(obj *InfoBarPrivate) {
	infoBarPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *InvisibleClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(invisibleClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *InvisibleClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(invisibleClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// InvisibleClassStruct creates an uninitialised InvisibleClass.
func InvisibleClassStruct() *InvisibleClass {
	err := invisibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InvisibleClass{native: invisibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeInvisibleClass)
	return structGo
}
func finalizeInvisibleClass(obj *InvisibleClass) {
	invisibleClassStruct.Free(obj.native)
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

// InvisiblePrivateStruct creates an uninitialised InvisiblePrivate.
func InvisiblePrivateStruct() *InvisiblePrivate {
	err := invisiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &InvisiblePrivate{native: invisiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeInvisiblePrivate)
	return structGo
}
func finalizeInvisiblePrivate(obj *InvisiblePrivate) {
	invisiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *LabelAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(labelAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *LabelAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(labelAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// LabelAccessibleClassStruct creates an uninitialised LabelAccessibleClass.
func LabelAccessibleClassStruct() *LabelAccessibleClass {
	err := labelAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LabelAccessibleClass{native: labelAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLabelAccessibleClass)
	return structGo
}
func finalizeLabelAccessibleClass(obj *LabelAccessibleClass) {
	labelAccessibleClassStruct.Free(obj.native)
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

// LabelAccessiblePrivateStruct creates an uninitialised LabelAccessiblePrivate.
func LabelAccessiblePrivateStruct() *LabelAccessiblePrivate {
	err := labelAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LabelAccessiblePrivate{native: labelAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLabelAccessiblePrivate)
	return structGo
}
func finalizeLabelAccessiblePrivate(obj *LabelAccessiblePrivate) {
	labelAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *LabelClass) FieldParentClass() *MiscClass {
	argValue := gi.FieldGet(labelClassStruct, recv.native, "parent_class")
	value := &MiscClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *LabelClass) SetFieldParentClass(value *MiscClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(labelClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'move_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'copy_clipboard' : for field getter : missing Type

// UNSUPPORTED : C value 'copy_clipboard' : for field setter : missing Type

// UNSUPPORTED : C value 'populate_popup' : for field getter : missing Type

// UNSUPPORTED : C value 'populate_popup' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_link' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_link' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// LabelClassStruct creates an uninitialised LabelClass.
func LabelClassStruct() *LabelClass {
	err := labelClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LabelClass{native: labelClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLabelClass)
	return structGo
}
func finalizeLabelClass(obj *LabelClass) {
	labelClassStruct.Free(obj.native)
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

// LabelPrivateStruct creates an uninitialised LabelPrivate.
func LabelPrivateStruct() *LabelPrivate {
	err := labelPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LabelPrivate{native: labelPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLabelPrivate)
	return structGo
}
func finalizeLabelPrivate(obj *LabelPrivate) {
	labelPrivateStruct.Free(obj.native)
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

// LabelSelectionInfoStruct creates an uninitialised LabelSelectionInfo.
func LabelSelectionInfoStruct() *LabelSelectionInfo {
	err := labelSelectionInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LabelSelectionInfo{native: labelSelectionInfoStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLabelSelectionInfo)
	return structGo
}
func finalizeLabelSelectionInfo(obj *LabelSelectionInfo) {
	labelSelectionInfoStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *LayoutClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(layoutClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *LayoutClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(layoutClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// LayoutClassStruct creates an uninitialised LayoutClass.
func LayoutClassStruct() *LayoutClass {
	err := layoutClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LayoutClass{native: layoutClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLayoutClass)
	return structGo
}
func finalizeLayoutClass(obj *LayoutClass) {
	layoutClassStruct.Free(obj.native)
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

// LayoutPrivateStruct creates an uninitialised LayoutPrivate.
func LayoutPrivateStruct() *LayoutPrivate {
	err := layoutPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LayoutPrivate{native: layoutPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLayoutPrivate)
	return structGo
}
func finalizeLayoutPrivate(obj *LayoutPrivate) {
	layoutPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *LevelBarAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(levelBarAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *LevelBarAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(levelBarAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// LevelBarAccessibleClassStruct creates an uninitialised LevelBarAccessibleClass.
func LevelBarAccessibleClassStruct() *LevelBarAccessibleClass {
	err := levelBarAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LevelBarAccessibleClass{native: levelBarAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLevelBarAccessibleClass)
	return structGo
}
func finalizeLevelBarAccessibleClass(obj *LevelBarAccessibleClass) {
	levelBarAccessibleClassStruct.Free(obj.native)
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

// LevelBarAccessiblePrivateStruct creates an uninitialised LevelBarAccessiblePrivate.
func LevelBarAccessiblePrivateStruct() *LevelBarAccessiblePrivate {
	err := levelBarAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LevelBarAccessiblePrivate{native: levelBarAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLevelBarAccessiblePrivate)
	return structGo
}
func finalizeLevelBarAccessiblePrivate(obj *LevelBarAccessiblePrivate) {
	levelBarAccessiblePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'offset_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'offset_changed' : for field setter : missing Type

// LevelBarClassStruct creates an uninitialised LevelBarClass.
func LevelBarClassStruct() *LevelBarClass {
	err := levelBarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LevelBarClass{native: levelBarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLevelBarClass)
	return structGo
}
func finalizeLevelBarClass(obj *LevelBarClass) {
	levelBarClassStruct.Free(obj.native)
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

// LevelBarPrivateStruct creates an uninitialised LevelBarPrivate.
func LevelBarPrivateStruct() *LevelBarPrivate {
	err := levelBarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LevelBarPrivate{native: levelBarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLevelBarPrivate)
	return structGo
}
func finalizeLevelBarPrivate(obj *LevelBarPrivate) {
	levelBarPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *LinkButtonAccessibleClass) FieldParentClass() *ButtonAccessibleClass {
	argValue := gi.FieldGet(linkButtonAccessibleClassStruct, recv.native, "parent_class")
	value := &ButtonAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *LinkButtonAccessibleClass) SetFieldParentClass(value *ButtonAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(linkButtonAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// LinkButtonAccessibleClassStruct creates an uninitialised LinkButtonAccessibleClass.
func LinkButtonAccessibleClassStruct() *LinkButtonAccessibleClass {
	err := linkButtonAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LinkButtonAccessibleClass{native: linkButtonAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLinkButtonAccessibleClass)
	return structGo
}
func finalizeLinkButtonAccessibleClass(obj *LinkButtonAccessibleClass) {
	linkButtonAccessibleClassStruct.Free(obj.native)
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

// LinkButtonAccessiblePrivateStruct creates an uninitialised LinkButtonAccessiblePrivate.
func LinkButtonAccessiblePrivateStruct() *LinkButtonAccessiblePrivate {
	err := linkButtonAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LinkButtonAccessiblePrivate{native: linkButtonAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLinkButtonAccessiblePrivate)
	return structGo
}
func finalizeLinkButtonAccessiblePrivate(obj *LinkButtonAccessiblePrivate) {
	linkButtonAccessiblePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'activate_link' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_link' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_padding1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_padding1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_padding2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_padding2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_padding3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_padding3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_padding4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_padding4' : for field setter : missing Type

// LinkButtonClassStruct creates an uninitialised LinkButtonClass.
func LinkButtonClassStruct() *LinkButtonClass {
	err := linkButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LinkButtonClass{native: linkButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLinkButtonClass)
	return structGo
}
func finalizeLinkButtonClass(obj *LinkButtonClass) {
	linkButtonClassStruct.Free(obj.native)
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

// LinkButtonPrivateStruct creates an uninitialised LinkButtonPrivate.
func LinkButtonPrivateStruct() *LinkButtonPrivate {
	err := linkButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LinkButtonPrivate{native: linkButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLinkButtonPrivate)
	return structGo
}
func finalizeLinkButtonPrivate(obj *LinkButtonPrivate) {
	linkButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ListBoxAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(listBoxAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ListBoxAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(listBoxAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ListBoxAccessibleClassStruct creates an uninitialised ListBoxAccessibleClass.
func ListBoxAccessibleClassStruct() *ListBoxAccessibleClass {
	err := listBoxAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ListBoxAccessibleClass{native: listBoxAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeListBoxAccessibleClass)
	return structGo
}
func finalizeListBoxAccessibleClass(obj *ListBoxAccessibleClass) {
	listBoxAccessibleClassStruct.Free(obj.native)
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

// ListBoxAccessiblePrivateStruct creates an uninitialised ListBoxAccessiblePrivate.
func ListBoxAccessiblePrivateStruct() *ListBoxAccessiblePrivate {
	err := listBoxAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ListBoxAccessiblePrivate{native: listBoxAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeListBoxAccessiblePrivate)
	return structGo
}
func finalizeListBoxAccessiblePrivate(obj *ListBoxAccessiblePrivate) {
	listBoxAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ListBoxClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(listBoxClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ListBoxClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(listBoxClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'row_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'row_selected' : for field setter : missing Type

// UNSUPPORTED : C value 'row_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'row_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_cursor_row' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_cursor_row' : for field setter : missing Type

// UNSUPPORTED : C value 'toggle_cursor_row' : for field getter : missing Type

// UNSUPPORTED : C value 'toggle_cursor_row' : for field setter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'selected_rows_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'selected_rows_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'select_all' : for field getter : missing Type

// UNSUPPORTED : C value 'select_all' : for field setter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field getter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// ListBoxClassStruct creates an uninitialised ListBoxClass.
func ListBoxClassStruct() *ListBoxClass {
	err := listBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ListBoxClass{native: listBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeListBoxClass)
	return structGo
}
func finalizeListBoxClass(obj *ListBoxClass) {
	listBoxClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ListBoxRowAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(listBoxRowAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ListBoxRowAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(listBoxRowAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ListBoxRowAccessibleClassStruct creates an uninitialised ListBoxRowAccessibleClass.
func ListBoxRowAccessibleClassStruct() *ListBoxRowAccessibleClass {
	err := listBoxRowAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ListBoxRowAccessibleClass{native: listBoxRowAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeListBoxRowAccessibleClass)
	return structGo
}
func finalizeListBoxRowAccessibleClass(obj *ListBoxRowAccessibleClass) {
	listBoxRowAccessibleClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ListBoxRowClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(listBoxRowClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ListBoxRowClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(listBoxRowClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// ListBoxRowClassStruct creates an uninitialised ListBoxRowClass.
func ListBoxRowClassStruct() *ListBoxRowClass {
	err := listBoxRowClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ListBoxRowClass{native: listBoxRowClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeListBoxRowClass)
	return structGo
}
func finalizeListBoxRowClass(obj *ListBoxRowClass) {
	listBoxRowClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ListStoreClassStruct creates an uninitialised ListStoreClass.
func ListStoreClassStruct() *ListStoreClass {
	err := listStoreClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ListStoreClass{native: listStoreClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeListStoreClass)
	return structGo
}
func finalizeListStoreClass(obj *ListStoreClass) {
	listStoreClassStruct.Free(obj.native)
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

// ListStorePrivateStruct creates an uninitialised ListStorePrivate.
func ListStorePrivateStruct() *ListStorePrivate {
	err := listStorePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ListStorePrivate{native: listStorePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeListStorePrivate)
	return structGo
}
func finalizeListStorePrivate(obj *ListStorePrivate) {
	listStorePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *LockButtonAccessibleClass) FieldParentClass() *ButtonAccessibleClass {
	argValue := gi.FieldGet(lockButtonAccessibleClassStruct, recv.native, "parent_class")
	value := &ButtonAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *LockButtonAccessibleClass) SetFieldParentClass(value *ButtonAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(lockButtonAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// LockButtonAccessibleClassStruct creates an uninitialised LockButtonAccessibleClass.
func LockButtonAccessibleClassStruct() *LockButtonAccessibleClass {
	err := lockButtonAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LockButtonAccessibleClass{native: lockButtonAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLockButtonAccessibleClass)
	return structGo
}
func finalizeLockButtonAccessibleClass(obj *LockButtonAccessibleClass) {
	lockButtonAccessibleClassStruct.Free(obj.native)
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

// LockButtonAccessiblePrivateStruct creates an uninitialised LockButtonAccessiblePrivate.
func LockButtonAccessiblePrivateStruct() *LockButtonAccessiblePrivate {
	err := lockButtonAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LockButtonAccessiblePrivate{native: lockButtonAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLockButtonAccessiblePrivate)
	return structGo
}
func finalizeLockButtonAccessiblePrivate(obj *LockButtonAccessiblePrivate) {
	lockButtonAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *LockButtonClass) FieldParentClass() *ButtonClass {
	argValue := gi.FieldGet(lockButtonClassStruct, recv.native, "parent_class")
	value := &ButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *LockButtonClass) SetFieldParentClass(value *ButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(lockButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'reserved0' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved0' : for field setter : missing Type

// UNSUPPORTED : C value 'reserved1' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved1' : for field setter : missing Type

// UNSUPPORTED : C value 'reserved2' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved2' : for field setter : missing Type

// UNSUPPORTED : C value 'reserved3' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved3' : for field setter : missing Type

// UNSUPPORTED : C value 'reserved4' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved4' : for field setter : missing Type

// UNSUPPORTED : C value 'reserved5' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved5' : for field setter : missing Type

// UNSUPPORTED : C value 'reserved6' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved6' : for field setter : missing Type

// UNSUPPORTED : C value 'reserved7' : for field getter : missing Type

// UNSUPPORTED : C value 'reserved7' : for field setter : missing Type

// LockButtonClassStruct creates an uninitialised LockButtonClass.
func LockButtonClassStruct() *LockButtonClass {
	err := lockButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LockButtonClass{native: lockButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLockButtonClass)
	return structGo
}
func finalizeLockButtonClass(obj *LockButtonClass) {
	lockButtonClassStruct.Free(obj.native)
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

// LockButtonPrivateStruct creates an uninitialised LockButtonPrivate.
func LockButtonPrivateStruct() *LockButtonPrivate {
	err := lockButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &LockButtonPrivate{native: lockButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeLockButtonPrivate)
	return structGo
}
func finalizeLockButtonPrivate(obj *LockButtonPrivate) {
	lockButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuAccessibleClass) FieldParentClass() *MenuShellAccessibleClass {
	argValue := gi.FieldGet(menuAccessibleClassStruct, recv.native, "parent_class")
	value := &MenuShellAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuAccessibleClass) SetFieldParentClass(value *MenuShellAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// MenuAccessibleClassStruct creates an uninitialised MenuAccessibleClass.
func MenuAccessibleClassStruct() *MenuAccessibleClass {
	err := menuAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuAccessibleClass{native: menuAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuAccessibleClass)
	return structGo
}
func finalizeMenuAccessibleClass(obj *MenuAccessibleClass) {
	menuAccessibleClassStruct.Free(obj.native)
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

// MenuAccessiblePrivateStruct creates an uninitialised MenuAccessiblePrivate.
func MenuAccessiblePrivateStruct() *MenuAccessiblePrivate {
	err := menuAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuAccessiblePrivate{native: menuAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuAccessiblePrivate)
	return structGo
}
func finalizeMenuAccessiblePrivate(obj *MenuAccessiblePrivate) {
	menuAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuBarClass) FieldParentClass() *MenuShellClass {
	argValue := gi.FieldGet(menuBarClassStruct, recv.native, "parent_class")
	value := &MenuShellClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuBarClass) SetFieldParentClass(value *MenuShellClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuBarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// MenuBarClassStruct creates an uninitialised MenuBarClass.
func MenuBarClassStruct() *MenuBarClass {
	err := menuBarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuBarClass{native: menuBarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuBarClass)
	return structGo
}
func finalizeMenuBarClass(obj *MenuBarClass) {
	menuBarClassStruct.Free(obj.native)
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

// MenuBarPrivateStruct creates an uninitialised MenuBarPrivate.
func MenuBarPrivateStruct() *MenuBarPrivate {
	err := menuBarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuBarPrivate{native: menuBarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuBarPrivate)
	return structGo
}
func finalizeMenuBarPrivate(obj *MenuBarPrivate) {
	menuBarPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuButtonAccessibleClass) FieldParentClass() *ToggleButtonAccessibleClass {
	argValue := gi.FieldGet(menuButtonAccessibleClassStruct, recv.native, "parent_class")
	value := &ToggleButtonAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuButtonAccessibleClass) SetFieldParentClass(value *ToggleButtonAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuButtonAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// MenuButtonAccessibleClassStruct creates an uninitialised MenuButtonAccessibleClass.
func MenuButtonAccessibleClassStruct() *MenuButtonAccessibleClass {
	err := menuButtonAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuButtonAccessibleClass{native: menuButtonAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuButtonAccessibleClass)
	return structGo
}
func finalizeMenuButtonAccessibleClass(obj *MenuButtonAccessibleClass) {
	menuButtonAccessibleClassStruct.Free(obj.native)
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

// MenuButtonAccessiblePrivateStruct creates an uninitialised MenuButtonAccessiblePrivate.
func MenuButtonAccessiblePrivateStruct() *MenuButtonAccessiblePrivate {
	err := menuButtonAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuButtonAccessiblePrivate{native: menuButtonAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuButtonAccessiblePrivate)
	return structGo
}
func finalizeMenuButtonAccessiblePrivate(obj *MenuButtonAccessiblePrivate) {
	menuButtonAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuButtonClass) FieldParentClass() *ToggleButtonClass {
	argValue := gi.FieldGet(menuButtonClassStruct, recv.native, "parent_class")
	value := &ToggleButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuButtonClass) SetFieldParentClass(value *ToggleButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// MenuButtonClassStruct creates an uninitialised MenuButtonClass.
func MenuButtonClassStruct() *MenuButtonClass {
	err := menuButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuButtonClass{native: menuButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuButtonClass)
	return structGo
}
func finalizeMenuButtonClass(obj *MenuButtonClass) {
	menuButtonClassStruct.Free(obj.native)
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

// MenuButtonPrivateStruct creates an uninitialised MenuButtonPrivate.
func MenuButtonPrivateStruct() *MenuButtonPrivate {
	err := menuButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuButtonPrivate{native: menuButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuButtonPrivate)
	return structGo
}
func finalizeMenuButtonPrivate(obj *MenuButtonPrivate) {
	menuButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuClass) FieldParentClass() *MenuShellClass {
	argValue := gi.FieldGet(menuClassStruct, recv.native, "parent_class")
	value := &MenuShellClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuClass) SetFieldParentClass(value *MenuShellClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// MenuClassStruct creates an uninitialised MenuClass.
func MenuClassStruct() *MenuClass {
	err := menuClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuClass{native: menuClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuClass)
	return structGo
}
func finalizeMenuClass(obj *MenuClass) {
	menuClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuItemAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(menuItemAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuItemAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuItemAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// MenuItemAccessibleClassStruct creates an uninitialised MenuItemAccessibleClass.
func MenuItemAccessibleClassStruct() *MenuItemAccessibleClass {
	err := menuItemAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuItemAccessibleClass{native: menuItemAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuItemAccessibleClass)
	return structGo
}
func finalizeMenuItemAccessibleClass(obj *MenuItemAccessibleClass) {
	menuItemAccessibleClassStruct.Free(obj.native)
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

// MenuItemAccessiblePrivateStruct creates an uninitialised MenuItemAccessiblePrivate.
func MenuItemAccessiblePrivateStruct() *MenuItemAccessiblePrivate {
	err := menuItemAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuItemAccessiblePrivate{native: menuItemAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuItemAccessiblePrivate)
	return structGo
}
func finalizeMenuItemAccessiblePrivate(obj *MenuItemAccessiblePrivate) {
	menuItemAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuItemClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(menuItemClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuItemClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuItemClassStruct, recv.native, "parent_class", argValue)
}

// FieldHideOnActivate returns the C field 'hide_on_activate'.
func (recv *MenuItemClass) FieldHideOnActivate() uint32 {
	argValue := gi.FieldGet(menuItemClassStruct, recv.native, "hide_on_activate")
	value := argValue.Uint32()
	return value
}

// SetFieldHideOnActivate sets the value of the C field 'hide_on_activate'.
func (recv *MenuItemClass) SetFieldHideOnActivate(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(menuItemClassStruct, recv.native, "hide_on_activate", argValue)
}

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_item' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_item' : for field setter : missing Type

// UNSUPPORTED : C value 'toggle_size_request' : for field getter : missing Type

// UNSUPPORTED : C value 'toggle_size_request' : for field setter : missing Type

// UNSUPPORTED : C value 'toggle_size_allocate' : for field getter : missing Type

// UNSUPPORTED : C value 'toggle_size_allocate' : for field setter : missing Type

// UNSUPPORTED : C value 'set_label' : for field getter : missing Type

// UNSUPPORTED : C value 'set_label' : for field setter : missing Type

// UNSUPPORTED : C value 'get_label' : for field getter : missing Type

// UNSUPPORTED : C value 'get_label' : for field setter : missing Type

// UNSUPPORTED : C value 'select' : for field getter : missing Type

// UNSUPPORTED : C value 'select' : for field setter : missing Type

// UNSUPPORTED : C value 'deselect' : for field getter : missing Type

// UNSUPPORTED : C value 'deselect' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// MenuItemClassStruct creates an uninitialised MenuItemClass.
func MenuItemClassStruct() *MenuItemClass {
	err := menuItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuItemClass{native: menuItemClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuItemClass)
	return structGo
}
func finalizeMenuItemClass(obj *MenuItemClass) {
	menuItemClassStruct.Free(obj.native)
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

// MenuItemPrivateStruct creates an uninitialised MenuItemPrivate.
func MenuItemPrivateStruct() *MenuItemPrivate {
	err := menuItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuItemPrivate{native: menuItemPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuItemPrivate)
	return structGo
}
func finalizeMenuItemPrivate(obj *MenuItemPrivate) {
	menuItemPrivateStruct.Free(obj.native)
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

// MenuPrivateStruct creates an uninitialised MenuPrivate.
func MenuPrivateStruct() *MenuPrivate {
	err := menuPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuPrivate{native: menuPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuPrivate)
	return structGo
}
func finalizeMenuPrivate(obj *MenuPrivate) {
	menuPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuShellAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(menuShellAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuShellAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuShellAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// MenuShellAccessibleClassStruct creates an uninitialised MenuShellAccessibleClass.
func MenuShellAccessibleClassStruct() *MenuShellAccessibleClass {
	err := menuShellAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuShellAccessibleClass{native: menuShellAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuShellAccessibleClass)
	return structGo
}
func finalizeMenuShellAccessibleClass(obj *MenuShellAccessibleClass) {
	menuShellAccessibleClassStruct.Free(obj.native)
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

// MenuShellAccessiblePrivateStruct creates an uninitialised MenuShellAccessiblePrivate.
func MenuShellAccessiblePrivateStruct() *MenuShellAccessiblePrivate {
	err := menuShellAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuShellAccessiblePrivate{native: menuShellAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuShellAccessiblePrivate)
	return structGo
}
func finalizeMenuShellAccessiblePrivate(obj *MenuShellAccessiblePrivate) {
	menuShellAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuShellClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(menuShellClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuShellClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuShellClassStruct, recv.native, "parent_class", argValue)
}

// FieldSubmenuPlacement returns the C field 'submenu_placement'.
func (recv *MenuShellClass) FieldSubmenuPlacement() uint32 {
	argValue := gi.FieldGet(menuShellClassStruct, recv.native, "submenu_placement")
	value := argValue.Uint32()
	return value
}

// SetFieldSubmenuPlacement sets the value of the C field 'submenu_placement'.
func (recv *MenuShellClass) SetFieldSubmenuPlacement(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(menuShellClassStruct, recv.native, "submenu_placement", argValue)
}

// UNSUPPORTED : C value 'deactivate' : for field getter : missing Type

// UNSUPPORTED : C value 'deactivate' : for field setter : missing Type

// UNSUPPORTED : C value 'selection_done' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_done' : for field setter : missing Type

// UNSUPPORTED : C value 'move_current' : for field getter : missing Type

// UNSUPPORTED : C value 'move_current' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_current' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_current' : for field setter : missing Type

// UNSUPPORTED : C value 'cancel' : for field getter : missing Type

// UNSUPPORTED : C value 'cancel' : for field setter : missing Type

// UNSUPPORTED : C value 'select_item' : for field getter : missing Type

// UNSUPPORTED : C value 'select_item' : for field setter : missing Type

// UNSUPPORTED : C value 'insert' : for field getter : missing Type

// UNSUPPORTED : C value 'insert' : for field setter : missing Type

// UNSUPPORTED : C value 'get_popup_delay' : for field getter : missing Type

// UNSUPPORTED : C value 'get_popup_delay' : for field setter : missing Type

// UNSUPPORTED : C value 'move_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'move_selected' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// MenuShellClassStruct creates an uninitialised MenuShellClass.
func MenuShellClassStruct() *MenuShellClass {
	err := menuShellClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuShellClass{native: menuShellClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuShellClass)
	return structGo
}
func finalizeMenuShellClass(obj *MenuShellClass) {
	menuShellClassStruct.Free(obj.native)
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

// MenuShellPrivateStruct creates an uninitialised MenuShellPrivate.
func MenuShellPrivateStruct() *MenuShellPrivate {
	err := menuShellPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuShellPrivate{native: menuShellPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuShellPrivate)
	return structGo
}
func finalizeMenuShellPrivate(obj *MenuShellPrivate) {
	menuShellPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MenuToolButtonClass) FieldParentClass() *ToolButtonClass {
	argValue := gi.FieldGet(menuToolButtonClassStruct, recv.native, "parent_class")
	value := &ToolButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MenuToolButtonClass) SetFieldParentClass(value *ToolButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(menuToolButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'show_menu' : for field getter : missing Type

// UNSUPPORTED : C value 'show_menu' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// MenuToolButtonClassStruct creates an uninitialised MenuToolButtonClass.
func MenuToolButtonClassStruct() *MenuToolButtonClass {
	err := menuToolButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuToolButtonClass{native: menuToolButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuToolButtonClass)
	return structGo
}
func finalizeMenuToolButtonClass(obj *MenuToolButtonClass) {
	menuToolButtonClassStruct.Free(obj.native)
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

// MenuToolButtonPrivateStruct creates an uninitialised MenuToolButtonPrivate.
func MenuToolButtonPrivateStruct() *MenuToolButtonPrivate {
	err := menuToolButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MenuToolButtonPrivate{native: menuToolButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMenuToolButtonPrivate)
	return structGo
}
func finalizeMenuToolButtonPrivate(obj *MenuToolButtonPrivate) {
	menuToolButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MessageDialogClass) FieldParentClass() *DialogClass {
	argValue := gi.FieldGet(messageDialogClassStruct, recv.native, "parent_class")
	value := &DialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MessageDialogClass) SetFieldParentClass(value *DialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(messageDialogClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// MessageDialogClassStruct creates an uninitialised MessageDialogClass.
func MessageDialogClassStruct() *MessageDialogClass {
	err := messageDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MessageDialogClass{native: messageDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMessageDialogClass)
	return structGo
}
func finalizeMessageDialogClass(obj *MessageDialogClass) {
	messageDialogClassStruct.Free(obj.native)
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

// MessageDialogPrivateStruct creates an uninitialised MessageDialogPrivate.
func MessageDialogPrivateStruct() *MessageDialogPrivate {
	err := messageDialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MessageDialogPrivate{native: messageDialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMessageDialogPrivate)
	return structGo
}
func finalizeMessageDialogPrivate(obj *MessageDialogPrivate) {
	messageDialogPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *MiscClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(miscClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *MiscClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(miscClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// MiscClassStruct creates an uninitialised MiscClass.
func MiscClassStruct() *MiscClass {
	err := miscClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MiscClass{native: miscClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMiscClass)
	return structGo
}
func finalizeMiscClass(obj *MiscClass) {
	miscClassStruct.Free(obj.native)
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

// MiscPrivateStruct creates an uninitialised MiscPrivate.
func MiscPrivateStruct() *MiscPrivate {
	err := miscPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MiscPrivate{native: miscPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMiscPrivate)
	return structGo
}
func finalizeMiscPrivate(obj *MiscPrivate) {
	miscPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gio.MountOperationClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Gio.MountOperationClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// MountOperationClassStruct creates an uninitialised MountOperationClass.
func MountOperationClassStruct() *MountOperationClass {
	err := mountOperationClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MountOperationClass{native: mountOperationClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMountOperationClass)
	return structGo
}
func finalizeMountOperationClass(obj *MountOperationClass) {
	mountOperationClassStruct.Free(obj.native)
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

// MountOperationPrivateStruct creates an uninitialised MountOperationPrivate.
func MountOperationPrivateStruct() *MountOperationPrivate {
	err := mountOperationPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &MountOperationPrivate{native: mountOperationPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeMountOperationPrivate)
	return structGo
}
func finalizeMountOperationPrivate(obj *MountOperationPrivate) {
	mountOperationPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'response' : for field getter : missing Type

// UNSUPPORTED : C value 'response' : for field setter : missing Type

// UNSUPPORTED : C value 'show' : for field getter : missing Type

// UNSUPPORTED : C value 'show' : for field setter : missing Type

// UNSUPPORTED : C value 'hide' : for field getter : missing Type

// UNSUPPORTED : C value 'hide' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// NativeDialogClassStruct creates an uninitialised NativeDialogClass.
func NativeDialogClassStruct() *NativeDialogClass {
	err := nativeDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NativeDialogClass{native: nativeDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNativeDialogClass)
	return structGo
}
func finalizeNativeDialogClass(obj *NativeDialogClass) {
	nativeDialogClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *NotebookAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(notebookAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *NotebookAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(notebookAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// NotebookAccessibleClassStruct creates an uninitialised NotebookAccessibleClass.
func NotebookAccessibleClassStruct() *NotebookAccessibleClass {
	err := notebookAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NotebookAccessibleClass{native: notebookAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNotebookAccessibleClass)
	return structGo
}
func finalizeNotebookAccessibleClass(obj *NotebookAccessibleClass) {
	notebookAccessibleClassStruct.Free(obj.native)
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

// NotebookAccessiblePrivateStruct creates an uninitialised NotebookAccessiblePrivate.
func NotebookAccessiblePrivateStruct() *NotebookAccessiblePrivate {
	err := notebookAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NotebookAccessiblePrivate{native: notebookAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNotebookAccessiblePrivate)
	return structGo
}
func finalizeNotebookAccessiblePrivate(obj *NotebookAccessiblePrivate) {
	notebookAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *NotebookClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(notebookClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *NotebookClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(notebookClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'switch_page' : for field getter : missing Type

// UNSUPPORTED : C value 'switch_page' : for field setter : missing Type

// UNSUPPORTED : C value 'select_page' : for field getter : missing Type

// UNSUPPORTED : C value 'select_page' : for field setter : missing Type

// UNSUPPORTED : C value 'focus_tab' : for field getter : missing Type

// UNSUPPORTED : C value 'focus_tab' : for field setter : missing Type

// UNSUPPORTED : C value 'change_current_page' : for field getter : missing Type

// UNSUPPORTED : C value 'change_current_page' : for field setter : missing Type

// UNSUPPORTED : C value 'move_focus_out' : for field getter : missing Type

// UNSUPPORTED : C value 'move_focus_out' : for field setter : missing Type

// UNSUPPORTED : C value 'reorder_tab' : for field getter : missing Type

// UNSUPPORTED : C value 'reorder_tab' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_page' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_page' : for field setter : missing Type

// UNSUPPORTED : C value 'create_window' : for field getter : missing Type

// UNSUPPORTED : C value 'create_window' : for field setter : missing Type

// UNSUPPORTED : C value 'page_reordered' : for field getter : missing Type

// UNSUPPORTED : C value 'page_reordered' : for field setter : missing Type

// UNSUPPORTED : C value 'page_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'page_removed' : for field setter : missing Type

// UNSUPPORTED : C value 'page_added' : for field getter : missing Type

// UNSUPPORTED : C value 'page_added' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// NotebookClassStruct creates an uninitialised NotebookClass.
func NotebookClassStruct() *NotebookClass {
	err := notebookClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NotebookClass{native: notebookClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNotebookClass)
	return structGo
}
func finalizeNotebookClass(obj *NotebookClass) {
	notebookClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Atk.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Atk.ObjectClass'

// NotebookPageAccessibleClassStruct creates an uninitialised NotebookPageAccessibleClass.
func NotebookPageAccessibleClassStruct() *NotebookPageAccessibleClass {
	err := notebookPageAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NotebookPageAccessibleClass{native: notebookPageAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNotebookPageAccessibleClass)
	return structGo
}
func finalizeNotebookPageAccessibleClass(obj *NotebookPageAccessibleClass) {
	notebookPageAccessibleClassStruct.Free(obj.native)
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

// NotebookPageAccessiblePrivateStruct creates an uninitialised NotebookPageAccessiblePrivate.
func NotebookPageAccessiblePrivateStruct() *NotebookPageAccessiblePrivate {
	err := notebookPageAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NotebookPageAccessiblePrivate{native: notebookPageAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNotebookPageAccessiblePrivate)
	return structGo
}
func finalizeNotebookPageAccessiblePrivate(obj *NotebookPageAccessiblePrivate) {
	notebookPageAccessiblePrivateStruct.Free(obj.native)
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

// NotebookPrivateStruct creates an uninitialised NotebookPrivate.
func NotebookPrivateStruct() *NotebookPrivate {
	err := notebookPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NotebookPrivate{native: notebookPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNotebookPrivate)
	return structGo
}
func finalizeNotebookPrivate(obj *NotebookPrivate) {
	notebookPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Gio.EmblemedIconClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Gio.EmblemedIconClass'

// UNSUPPORTED : C value 'padding' : for field getter : missing Type

// UNSUPPORTED : C value 'padding' : for field setter : missing Type

// NumerableIconClassStruct creates an uninitialised NumerableIconClass.
func NumerableIconClassStruct() *NumerableIconClass {
	err := numerableIconClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NumerableIconClass{native: numerableIconClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNumerableIconClass)
	return structGo
}
func finalizeNumerableIconClass(obj *NumerableIconClass) {
	numerableIconClassStruct.Free(obj.native)
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

// NumerableIconPrivateStruct creates an uninitialised NumerableIconPrivate.
func NumerableIconPrivateStruct() *NumerableIconPrivate {
	err := numerableIconPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &NumerableIconPrivate{native: numerableIconPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeNumerableIconPrivate)
	return structGo
}
func finalizeNumerableIconPrivate(obj *NumerableIconPrivate) {
	numerableIconPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *OffscreenWindowClass) FieldParentClass() *WindowClass {
	argValue := gi.FieldGet(offscreenWindowClassStruct, recv.native, "parent_class")
	value := &WindowClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *OffscreenWindowClass) SetFieldParentClass(value *WindowClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(offscreenWindowClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// OffscreenWindowClassStruct creates an uninitialised OffscreenWindowClass.
func OffscreenWindowClassStruct() *OffscreenWindowClass {
	err := offscreenWindowClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &OffscreenWindowClass{native: offscreenWindowClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeOffscreenWindowClass)
	return structGo
}
func finalizeOffscreenWindowClass(obj *OffscreenWindowClass) {
	offscreenWindowClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'base_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'base_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// OrientableIfaceStruct creates an uninitialised OrientableIface.
func OrientableIfaceStruct() *OrientableIface {
	err := orientableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &OrientableIface{native: orientableIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeOrientableIface)
	return structGo
}
func finalizeOrientableIface(obj *OrientableIface) {
	orientableIfaceStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *OverlayClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(overlayClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *OverlayClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(overlayClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'get_child_position' : for field getter : missing Type

// UNSUPPORTED : C value 'get_child_position' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// OverlayClassStruct creates an uninitialised OverlayClass.
func OverlayClassStruct() *OverlayClass {
	err := overlayClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &OverlayClass{native: overlayClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeOverlayClass)
	return structGo
}
func finalizeOverlayClass(obj *OverlayClass) {
	overlayClassStruct.Free(obj.native)
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

// OverlayPrivateStruct creates an uninitialised OverlayPrivate.
func OverlayPrivateStruct() *OverlayPrivate {
	err := overlayPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &OverlayPrivate{native: overlayPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeOverlayPrivate)
	return structGo
}
func finalizeOverlayPrivate(obj *OverlayPrivate) {
	overlayPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'type' : for field getter : no Go type for 'PadActionType'

// UNSUPPORTED : C value 'type' : for field setter : no Go type for 'PadActionType'

// FieldIndex returns the C field 'index'.
func (recv *PadActionEntry) FieldIndex() int32 {
	argValue := gi.FieldGet(padActionEntryStruct, recv.native, "index")
	value := argValue.Int32()
	return value
}

// SetFieldIndex sets the value of the C field 'index'.
func (recv *PadActionEntry) SetFieldIndex(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(padActionEntryStruct, recv.native, "index", argValue)
}

// FieldMode returns the C field 'mode'.
func (recv *PadActionEntry) FieldMode() int32 {
	argValue := gi.FieldGet(padActionEntryStruct, recv.native, "mode")
	value := argValue.Int32()
	return value
}

// SetFieldMode sets the value of the C field 'mode'.
func (recv *PadActionEntry) SetFieldMode(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(padActionEntryStruct, recv.native, "mode", argValue)
}

// FieldLabel returns the C field 'label'.
func (recv *PadActionEntry) FieldLabel() string {
	argValue := gi.FieldGet(padActionEntryStruct, recv.native, "label")
	value := argValue.String(false)
	return value
}

// SetFieldLabel sets the value of the C field 'label'.
func (recv *PadActionEntry) SetFieldLabel(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(padActionEntryStruct, recv.native, "label", argValue)
}

// FieldActionName returns the C field 'action_name'.
func (recv *PadActionEntry) FieldActionName() string {
	argValue := gi.FieldGet(padActionEntryStruct, recv.native, "action_name")
	value := argValue.String(false)
	return value
}

// SetFieldActionName sets the value of the C field 'action_name'.
func (recv *PadActionEntry) SetFieldActionName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(padActionEntryStruct, recv.native, "action_name", argValue)
}

// PadActionEntryStruct creates an uninitialised PadActionEntry.
func PadActionEntryStruct() *PadActionEntry {
	err := padActionEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PadActionEntry{native: padActionEntryStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePadActionEntry)
	return structGo
}
func finalizePadActionEntry(obj *PadActionEntry) {
	padActionEntryStruct.Free(obj.native)
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

// PadControllerClassStruct creates an uninitialised PadControllerClass.
func PadControllerClassStruct() *PadControllerClass {
	err := padControllerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PadControllerClass{native: padControllerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePadControllerClass)
	return structGo
}
func finalizePadControllerClass(obj *PadControllerClass) {
	padControllerClassStruct.Free(obj.native)
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

// FieldStart returns the C field 'start'.
func (recv *PageRange) FieldStart() int32 {
	argValue := gi.FieldGet(pageRangeStruct, recv.native, "start")
	value := argValue.Int32()
	return value
}

// SetFieldStart sets the value of the C field 'start'.
func (recv *PageRange) SetFieldStart(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(pageRangeStruct, recv.native, "start", argValue)
}

// FieldEnd returns the C field 'end'.
func (recv *PageRange) FieldEnd() int32 {
	argValue := gi.FieldGet(pageRangeStruct, recv.native, "end")
	value := argValue.Int32()
	return value
}

// SetFieldEnd sets the value of the C field 'end'.
func (recv *PageRange) SetFieldEnd(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(pageRangeStruct, recv.native, "end", argValue)
}

// PageRangeStruct creates an uninitialised PageRange.
func PageRangeStruct() *PageRange {
	err := pageRangeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PageRange{native: pageRangeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePageRange)
	return structGo
}
func finalizePageRange(obj *PageRange) {
	pageRangeStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *PanedAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(panedAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *PanedAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(panedAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// PanedAccessibleClassStruct creates an uninitialised PanedAccessibleClass.
func PanedAccessibleClassStruct() *PanedAccessibleClass {
	err := panedAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PanedAccessibleClass{native: panedAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePanedAccessibleClass)
	return structGo
}
func finalizePanedAccessibleClass(obj *PanedAccessibleClass) {
	panedAccessibleClassStruct.Free(obj.native)
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

// PanedAccessiblePrivateStruct creates an uninitialised PanedAccessiblePrivate.
func PanedAccessiblePrivateStruct() *PanedAccessiblePrivate {
	err := panedAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PanedAccessiblePrivate{native: panedAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePanedAccessiblePrivate)
	return structGo
}
func finalizePanedAccessiblePrivate(obj *PanedAccessiblePrivate) {
	panedAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *PanedClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(panedClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *PanedClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(panedClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'cycle_child_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'cycle_child_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'toggle_handle_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'toggle_handle_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'move_handle' : for field getter : missing Type

// UNSUPPORTED : C value 'move_handle' : for field setter : missing Type

// UNSUPPORTED : C value 'cycle_handle_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'cycle_handle_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'accept_position' : for field getter : missing Type

// UNSUPPORTED : C value 'accept_position' : for field setter : missing Type

// UNSUPPORTED : C value 'cancel_position' : for field getter : missing Type

// UNSUPPORTED : C value 'cancel_position' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// PanedClassStruct creates an uninitialised PanedClass.
func PanedClassStruct() *PanedClass {
	err := panedClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PanedClass{native: panedClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePanedClass)
	return structGo
}
func finalizePanedClass(obj *PanedClass) {
	panedClassStruct.Free(obj.native)
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

// PanedPrivateStruct creates an uninitialised PanedPrivate.
func PanedPrivateStruct() *PanedPrivate {
	err := panedPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PanedPrivate{native: panedPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePanedPrivate)
	return structGo
}
func finalizePanedPrivate(obj *PanedPrivate) {
	panedPrivateStruct.Free(obj.native)
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

// PlacesSidebarClassStruct creates an uninitialised PlacesSidebarClass.
func PlacesSidebarClassStruct() *PlacesSidebarClass {
	err := placesSidebarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PlacesSidebarClass{native: placesSidebarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePlacesSidebarClass)
	return structGo
}
func finalizePlacesSidebarClass(obj *PlacesSidebarClass) {
	placesSidebarClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *PlugClass) FieldParentClass() *WindowClass {
	argValue := gi.FieldGet(plugClassStruct, recv.native, "parent_class")
	value := &WindowClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *PlugClass) SetFieldParentClass(value *WindowClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(plugClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'embedded' : for field getter : missing Type

// UNSUPPORTED : C value 'embedded' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// PlugClassStruct creates an uninitialised PlugClass.
func PlugClassStruct() *PlugClass {
	err := plugClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PlugClass{native: plugClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePlugClass)
	return structGo
}
func finalizePlugClass(obj *PlugClass) {
	plugClassStruct.Free(obj.native)
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

// PlugPrivateStruct creates an uninitialised PlugPrivate.
func PlugPrivateStruct() *PlugPrivate {
	err := plugPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PlugPrivate{native: plugPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePlugPrivate)
	return structGo
}
func finalizePlugPrivate(obj *PlugPrivate) {
	plugPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *PopoverAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(popoverAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *PopoverAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(popoverAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// PopoverAccessibleClassStruct creates an uninitialised PopoverAccessibleClass.
func PopoverAccessibleClassStruct() *PopoverAccessibleClass {
	err := popoverAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PopoverAccessibleClass{native: popoverAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePopoverAccessibleClass)
	return structGo
}
func finalizePopoverAccessibleClass(obj *PopoverAccessibleClass) {
	popoverAccessibleClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *PopoverClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(popoverClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *PopoverClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(popoverClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'closed' : for field getter : missing Type

// UNSUPPORTED : C value 'closed' : for field setter : missing Type

// PopoverClassStruct creates an uninitialised PopoverClass.
func PopoverClassStruct() *PopoverClass {
	err := popoverClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PopoverClass{native: popoverClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePopoverClass)
	return structGo
}
func finalizePopoverClass(obj *PopoverClass) {
	popoverClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *PopoverMenuClass) FieldParentClass() *PopoverClass {
	argValue := gi.FieldGet(popoverMenuClassStruct, recv.native, "parent_class")
	value := &PopoverClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *PopoverMenuClass) SetFieldParentClass(value *PopoverClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(popoverMenuClassStruct, recv.native, "parent_class", argValue)
}

// PopoverMenuClassStruct creates an uninitialised PopoverMenuClass.
func PopoverMenuClassStruct() *PopoverMenuClass {
	err := popoverMenuClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PopoverMenuClass{native: popoverMenuClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePopoverMenuClass)
	return structGo
}
func finalizePopoverMenuClass(obj *PopoverMenuClass) {
	popoverMenuClassStruct.Free(obj.native)
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

// PopoverPrivateStruct creates an uninitialised PopoverPrivate.
func PopoverPrivateStruct() *PopoverPrivate {
	err := popoverPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PopoverPrivate{native: popoverPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePopoverPrivate)
	return structGo
}
func finalizePopoverPrivate(obj *PopoverPrivate) {
	popoverPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'done' : for field getter : missing Type

// UNSUPPORTED : C value 'done' : for field setter : missing Type

// UNSUPPORTED : C value 'begin_print' : for field getter : missing Type

// UNSUPPORTED : C value 'begin_print' : for field setter : missing Type

// UNSUPPORTED : C value 'paginate' : for field getter : missing Type

// UNSUPPORTED : C value 'paginate' : for field setter : missing Type

// UNSUPPORTED : C value 'request_page_setup' : for field getter : missing Type

// UNSUPPORTED : C value 'request_page_setup' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_page' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_page' : for field setter : missing Type

// UNSUPPORTED : C value 'end_print' : for field getter : missing Type

// UNSUPPORTED : C value 'end_print' : for field setter : missing Type

// UNSUPPORTED : C value 'status_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'status_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'create_custom_widget' : for field getter : missing Type

// UNSUPPORTED : C value 'create_custom_widget' : for field setter : missing Type

// UNSUPPORTED : C value 'custom_widget_apply' : for field getter : missing Type

// UNSUPPORTED : C value 'custom_widget_apply' : for field setter : missing Type

// UNSUPPORTED : C value 'preview' : for field getter : missing Type

// UNSUPPORTED : C value 'preview' : for field setter : missing Type

// UNSUPPORTED : C value 'update_custom_widget' : for field getter : missing Type

// UNSUPPORTED : C value 'update_custom_widget' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// PrintOperationClassStruct creates an uninitialised PrintOperationClass.
func PrintOperationClassStruct() *PrintOperationClass {
	err := printOperationClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PrintOperationClass{native: printOperationClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePrintOperationClass)
	return structGo
}
func finalizePrintOperationClass(obj *PrintOperationClass) {
	printOperationClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'g_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'g_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'ready' : for field getter : missing Type

// UNSUPPORTED : C value 'ready' : for field setter : missing Type

// UNSUPPORTED : C value 'got_page_size' : for field getter : missing Type

// UNSUPPORTED : C value 'got_page_size' : for field setter : missing Type

// UNSUPPORTED : C value 'render_page' : for field getter : missing Type

// UNSUPPORTED : C value 'render_page' : for field setter : missing Type

// UNSUPPORTED : C value 'is_selected' : for field getter : missing Type

// UNSUPPORTED : C value 'is_selected' : for field setter : missing Type

// UNSUPPORTED : C value 'end_preview' : for field getter : missing Type

// UNSUPPORTED : C value 'end_preview' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// PrintOperationPreviewIfaceStruct creates an uninitialised PrintOperationPreviewIface.
func PrintOperationPreviewIfaceStruct() *PrintOperationPreviewIface {
	err := printOperationPreviewIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PrintOperationPreviewIface{native: printOperationPreviewIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePrintOperationPreviewIface)
	return structGo
}
func finalizePrintOperationPreviewIface(obj *PrintOperationPreviewIface) {
	printOperationPreviewIfaceStruct.Free(obj.native)
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

// PrintOperationPrivateStruct creates an uninitialised PrintOperationPrivate.
func PrintOperationPrivateStruct() *PrintOperationPrivate {
	err := printOperationPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &PrintOperationPrivate{native: printOperationPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizePrintOperationPrivate)
	return structGo
}
func finalizePrintOperationPrivate(obj *PrintOperationPrivate) {
	printOperationPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ProgressBarAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(progressBarAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ProgressBarAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(progressBarAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ProgressBarAccessibleClassStruct creates an uninitialised ProgressBarAccessibleClass.
func ProgressBarAccessibleClassStruct() *ProgressBarAccessibleClass {
	err := progressBarAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProgressBarAccessibleClass{native: progressBarAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeProgressBarAccessibleClass)
	return structGo
}
func finalizeProgressBarAccessibleClass(obj *ProgressBarAccessibleClass) {
	progressBarAccessibleClassStruct.Free(obj.native)
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

// ProgressBarAccessiblePrivateStruct creates an uninitialised ProgressBarAccessiblePrivate.
func ProgressBarAccessiblePrivateStruct() *ProgressBarAccessiblePrivate {
	err := progressBarAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProgressBarAccessiblePrivate{native: progressBarAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeProgressBarAccessiblePrivate)
	return structGo
}
func finalizeProgressBarAccessiblePrivate(obj *ProgressBarAccessiblePrivate) {
	progressBarAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ProgressBarClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(progressBarClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ProgressBarClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(progressBarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ProgressBarClassStruct creates an uninitialised ProgressBarClass.
func ProgressBarClassStruct() *ProgressBarClass {
	err := progressBarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProgressBarClass{native: progressBarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeProgressBarClass)
	return structGo
}
func finalizeProgressBarClass(obj *ProgressBarClass) {
	progressBarClassStruct.Free(obj.native)
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

// ProgressBarPrivateStruct creates an uninitialised ProgressBarPrivate.
func ProgressBarPrivateStruct() *ProgressBarPrivate {
	err := progressBarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ProgressBarPrivate{native: progressBarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeProgressBarPrivate)
	return structGo
}
func finalizeProgressBarPrivate(obj *ProgressBarPrivate) {
	progressBarPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RadioActionClass) FieldParentClass() *ToggleActionClass {
	argValue := gi.FieldGet(radioActionClassStruct, recv.native, "parent_class")
	value := &ToggleActionClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RadioActionClass) SetFieldParentClass(value *ToggleActionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(radioActionClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// RadioActionClassStruct creates an uninitialised RadioActionClass.
func RadioActionClassStruct() *RadioActionClass {
	err := radioActionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioActionClass{native: radioActionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioActionClass)
	return structGo
}
func finalizeRadioActionClass(obj *RadioActionClass) {
	radioActionClassStruct.Free(obj.native)
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

// FieldName returns the C field 'name'.
func (recv *RadioActionEntry) FieldName() string {
	argValue := gi.FieldGet(radioActionEntryStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *RadioActionEntry) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(radioActionEntryStruct, recv.native, "name", argValue)
}

// FieldStockId returns the C field 'stock_id'.
func (recv *RadioActionEntry) FieldStockId() string {
	argValue := gi.FieldGet(radioActionEntryStruct, recv.native, "stock_id")
	value := argValue.String(false)
	return value
}

// SetFieldStockId sets the value of the C field 'stock_id'.
func (recv *RadioActionEntry) SetFieldStockId(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(radioActionEntryStruct, recv.native, "stock_id", argValue)
}

// FieldLabel returns the C field 'label'.
func (recv *RadioActionEntry) FieldLabel() string {
	argValue := gi.FieldGet(radioActionEntryStruct, recv.native, "label")
	value := argValue.String(false)
	return value
}

// SetFieldLabel sets the value of the C field 'label'.
func (recv *RadioActionEntry) SetFieldLabel(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(radioActionEntryStruct, recv.native, "label", argValue)
}

// FieldAccelerator returns the C field 'accelerator'.
func (recv *RadioActionEntry) FieldAccelerator() string {
	argValue := gi.FieldGet(radioActionEntryStruct, recv.native, "accelerator")
	value := argValue.String(false)
	return value
}

// SetFieldAccelerator sets the value of the C field 'accelerator'.
func (recv *RadioActionEntry) SetFieldAccelerator(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(radioActionEntryStruct, recv.native, "accelerator", argValue)
}

// FieldTooltip returns the C field 'tooltip'.
func (recv *RadioActionEntry) FieldTooltip() string {
	argValue := gi.FieldGet(radioActionEntryStruct, recv.native, "tooltip")
	value := argValue.String(false)
	return value
}

// SetFieldTooltip sets the value of the C field 'tooltip'.
func (recv *RadioActionEntry) SetFieldTooltip(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(radioActionEntryStruct, recv.native, "tooltip", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *RadioActionEntry) FieldValue() int32 {
	argValue := gi.FieldGet(radioActionEntryStruct, recv.native, "value")
	value := argValue.Int32()
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *RadioActionEntry) SetFieldValue(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(radioActionEntryStruct, recv.native, "value", argValue)
}

// RadioActionEntryStruct creates an uninitialised RadioActionEntry.
func RadioActionEntryStruct() *RadioActionEntry {
	err := radioActionEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioActionEntry{native: radioActionEntryStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioActionEntry)
	return structGo
}
func finalizeRadioActionEntry(obj *RadioActionEntry) {
	radioActionEntryStruct.Free(obj.native)
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

// RadioActionPrivateStruct creates an uninitialised RadioActionPrivate.
func RadioActionPrivateStruct() *RadioActionPrivate {
	err := radioActionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioActionPrivate{native: radioActionPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioActionPrivate)
	return structGo
}
func finalizeRadioActionPrivate(obj *RadioActionPrivate) {
	radioActionPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RadioButtonAccessibleClass) FieldParentClass() *ToggleButtonAccessibleClass {
	argValue := gi.FieldGet(radioButtonAccessibleClassStruct, recv.native, "parent_class")
	value := &ToggleButtonAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RadioButtonAccessibleClass) SetFieldParentClass(value *ToggleButtonAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(radioButtonAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// RadioButtonAccessibleClassStruct creates an uninitialised RadioButtonAccessibleClass.
func RadioButtonAccessibleClassStruct() *RadioButtonAccessibleClass {
	err := radioButtonAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioButtonAccessibleClass{native: radioButtonAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioButtonAccessibleClass)
	return structGo
}
func finalizeRadioButtonAccessibleClass(obj *RadioButtonAccessibleClass) {
	radioButtonAccessibleClassStruct.Free(obj.native)
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

// RadioButtonAccessiblePrivateStruct creates an uninitialised RadioButtonAccessiblePrivate.
func RadioButtonAccessiblePrivateStruct() *RadioButtonAccessiblePrivate {
	err := radioButtonAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioButtonAccessiblePrivate{native: radioButtonAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioButtonAccessiblePrivate)
	return structGo
}
func finalizeRadioButtonAccessiblePrivate(obj *RadioButtonAccessiblePrivate) {
	radioButtonAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RadioButtonClass) FieldParentClass() *CheckButtonClass {
	argValue := gi.FieldGet(radioButtonClassStruct, recv.native, "parent_class")
	value := &CheckButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RadioButtonClass) SetFieldParentClass(value *CheckButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(radioButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'group_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'group_changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// RadioButtonClassStruct creates an uninitialised RadioButtonClass.
func RadioButtonClassStruct() *RadioButtonClass {
	err := radioButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioButtonClass{native: radioButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioButtonClass)
	return structGo
}
func finalizeRadioButtonClass(obj *RadioButtonClass) {
	radioButtonClassStruct.Free(obj.native)
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

// RadioButtonPrivateStruct creates an uninitialised RadioButtonPrivate.
func RadioButtonPrivateStruct() *RadioButtonPrivate {
	err := radioButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioButtonPrivate{native: radioButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioButtonPrivate)
	return structGo
}
func finalizeRadioButtonPrivate(obj *RadioButtonPrivate) {
	radioButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RadioMenuItemAccessibleClass) FieldParentClass() *CheckMenuItemAccessibleClass {
	argValue := gi.FieldGet(radioMenuItemAccessibleClassStruct, recv.native, "parent_class")
	value := &CheckMenuItemAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RadioMenuItemAccessibleClass) SetFieldParentClass(value *CheckMenuItemAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(radioMenuItemAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// RadioMenuItemAccessibleClassStruct creates an uninitialised RadioMenuItemAccessibleClass.
func RadioMenuItemAccessibleClassStruct() *RadioMenuItemAccessibleClass {
	err := radioMenuItemAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioMenuItemAccessibleClass{native: radioMenuItemAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioMenuItemAccessibleClass)
	return structGo
}
func finalizeRadioMenuItemAccessibleClass(obj *RadioMenuItemAccessibleClass) {
	radioMenuItemAccessibleClassStruct.Free(obj.native)
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

// RadioMenuItemAccessiblePrivateStruct creates an uninitialised RadioMenuItemAccessiblePrivate.
func RadioMenuItemAccessiblePrivateStruct() *RadioMenuItemAccessiblePrivate {
	err := radioMenuItemAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioMenuItemAccessiblePrivate{native: radioMenuItemAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioMenuItemAccessiblePrivate)
	return structGo
}
func finalizeRadioMenuItemAccessiblePrivate(obj *RadioMenuItemAccessiblePrivate) {
	radioMenuItemAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RadioMenuItemClass) FieldParentClass() *CheckMenuItemClass {
	argValue := gi.FieldGet(radioMenuItemClassStruct, recv.native, "parent_class")
	value := &CheckMenuItemClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RadioMenuItemClass) SetFieldParentClass(value *CheckMenuItemClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(radioMenuItemClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'group_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'group_changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// RadioMenuItemClassStruct creates an uninitialised RadioMenuItemClass.
func RadioMenuItemClassStruct() *RadioMenuItemClass {
	err := radioMenuItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioMenuItemClass{native: radioMenuItemClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioMenuItemClass)
	return structGo
}
func finalizeRadioMenuItemClass(obj *RadioMenuItemClass) {
	radioMenuItemClassStruct.Free(obj.native)
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

// RadioMenuItemPrivateStruct creates an uninitialised RadioMenuItemPrivate.
func RadioMenuItemPrivateStruct() *RadioMenuItemPrivate {
	err := radioMenuItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioMenuItemPrivate{native: radioMenuItemPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioMenuItemPrivate)
	return structGo
}
func finalizeRadioMenuItemPrivate(obj *RadioMenuItemPrivate) {
	radioMenuItemPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RadioToolButtonClass) FieldParentClass() *ToggleToolButtonClass {
	argValue := gi.FieldGet(radioToolButtonClassStruct, recv.native, "parent_class")
	value := &ToggleToolButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RadioToolButtonClass) SetFieldParentClass(value *ToggleToolButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(radioToolButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// RadioToolButtonClassStruct creates an uninitialised RadioToolButtonClass.
func RadioToolButtonClassStruct() *RadioToolButtonClass {
	err := radioToolButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RadioToolButtonClass{native: radioToolButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRadioToolButtonClass)
	return structGo
}
func finalizeRadioToolButtonClass(obj *RadioToolButtonClass) {
	radioToolButtonClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RangeAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(rangeAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RangeAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(rangeAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// RangeAccessibleClassStruct creates an uninitialised RangeAccessibleClass.
func RangeAccessibleClassStruct() *RangeAccessibleClass {
	err := rangeAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RangeAccessibleClass{native: rangeAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRangeAccessibleClass)
	return structGo
}
func finalizeRangeAccessibleClass(obj *RangeAccessibleClass) {
	rangeAccessibleClassStruct.Free(obj.native)
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

// RangeAccessiblePrivateStruct creates an uninitialised RangeAccessiblePrivate.
func RangeAccessiblePrivateStruct() *RangeAccessiblePrivate {
	err := rangeAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RangeAccessiblePrivate{native: rangeAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRangeAccessiblePrivate)
	return structGo
}
func finalizeRangeAccessiblePrivate(obj *RangeAccessiblePrivate) {
	rangeAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RangeClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(rangeClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RangeClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(rangeClassStruct, recv.native, "parent_class", argValue)
}

// FieldSliderDetail returns the C field 'slider_detail'.
func (recv *RangeClass) FieldSliderDetail() string {
	argValue := gi.FieldGet(rangeClassStruct, recv.native, "slider_detail")
	value := argValue.String(false)
	return value
}

// SetFieldSliderDetail sets the value of the C field 'slider_detail'.
func (recv *RangeClass) SetFieldSliderDetail(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(rangeClassStruct, recv.native, "slider_detail", argValue)
}

// FieldStepperDetail returns the C field 'stepper_detail'.
func (recv *RangeClass) FieldStepperDetail() string {
	argValue := gi.FieldGet(rangeClassStruct, recv.native, "stepper_detail")
	value := argValue.String(false)
	return value
}

// SetFieldStepperDetail sets the value of the C field 'stepper_detail'.
func (recv *RangeClass) SetFieldStepperDetail(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(rangeClassStruct, recv.native, "stepper_detail", argValue)
}

// UNSUPPORTED : C value 'value_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'value_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'adjust_bounds' : for field getter : missing Type

// UNSUPPORTED : C value 'adjust_bounds' : for field setter : missing Type

// UNSUPPORTED : C value 'move_slider' : for field getter : missing Type

// UNSUPPORTED : C value 'move_slider' : for field setter : missing Type

// UNSUPPORTED : C value 'get_range_border' : for field getter : missing Type

// UNSUPPORTED : C value 'get_range_border' : for field setter : missing Type

// UNSUPPORTED : C value 'change_value' : for field getter : missing Type

// UNSUPPORTED : C value 'change_value' : for field setter : missing Type

// UNSUPPORTED : C value 'get_range_size_request' : for field getter : missing Type

// UNSUPPORTED : C value 'get_range_size_request' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// RangeClassStruct creates an uninitialised RangeClass.
func RangeClassStruct() *RangeClass {
	err := rangeClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RangeClass{native: rangeClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRangeClass)
	return structGo
}
func finalizeRangeClass(obj *RangeClass) {
	rangeClassStruct.Free(obj.native)
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

// RangePrivateStruct creates an uninitialised RangePrivate.
func RangePrivateStruct() *RangePrivate {
	err := rangePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RangePrivate{native: rangePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRangePrivate)
	return structGo
}
func finalizeRangePrivate(obj *RangePrivate) {
	rangePrivateStruct.Free(obj.native)
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

// RcContextStruct creates an uninitialised RcContext.
func RcContextStruct() *RcContext {
	err := rcContextStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RcContext{native: rcContextStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRcContext)
	return structGo
}
func finalizeRcContext(obj *RcContext) {
	rcContextStruct.Free(obj.native)
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

// FieldTypeName returns the C field 'type_name'.
func (recv *RcProperty) FieldTypeName() glib.Quark {
	argValue := gi.FieldGet(rcPropertyStruct, recv.native, "type_name")
	value := glib.Quark(argValue.Uint32())
	return value
}

// SetFieldTypeName sets the value of the C field 'type_name'.
func (recv *RcProperty) SetFieldTypeName(value glib.Quark) {
	var argValue gi.Argument
	argValue.SetUint32(uint32(value))
	gi.FieldSet(rcPropertyStruct, recv.native, "type_name", argValue)
}

// FieldPropertyName returns the C field 'property_name'.
func (recv *RcProperty) FieldPropertyName() glib.Quark {
	argValue := gi.FieldGet(rcPropertyStruct, recv.native, "property_name")
	value := glib.Quark(argValue.Uint32())
	return value
}

// SetFieldPropertyName sets the value of the C field 'property_name'.
func (recv *RcProperty) SetFieldPropertyName(value glib.Quark) {
	var argValue gi.Argument
	argValue.SetUint32(uint32(value))
	gi.FieldSet(rcPropertyStruct, recv.native, "property_name", argValue)
}

// FieldOrigin returns the C field 'origin'.
func (recv *RcProperty) FieldOrigin() string {
	argValue := gi.FieldGet(rcPropertyStruct, recv.native, "origin")
	value := argValue.String(false)
	return value
}

// SetFieldOrigin sets the value of the C field 'origin'.
func (recv *RcProperty) SetFieldOrigin(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(rcPropertyStruct, recv.native, "origin", argValue)
}

// UNSUPPORTED : C value 'value' : for field getter : no Go type for 'GObject.Value'

// UNSUPPORTED : C value 'value' : for field setter : no Go type for 'GObject.Value'

// RcPropertyStruct creates an uninitialised RcProperty.
func RcPropertyStruct() *RcProperty {
	err := rcPropertyStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RcProperty{native: rcPropertyStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRcProperty)
	return structGo
}
func finalizeRcProperty(obj *RcProperty) {
	rcPropertyStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'create_rc_style' : for field getter : missing Type

// UNSUPPORTED : C value 'create_rc_style' : for field setter : missing Type

// UNSUPPORTED : C value 'parse' : for field getter : missing Type

// UNSUPPORTED : C value 'parse' : for field setter : missing Type

// UNSUPPORTED : C value 'merge' : for field getter : missing Type

// UNSUPPORTED : C value 'merge' : for field setter : missing Type

// UNSUPPORTED : C value 'create_style' : for field getter : missing Type

// UNSUPPORTED : C value 'create_style' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// RcStyleClassStruct creates an uninitialised RcStyleClass.
func RcStyleClassStruct() *RcStyleClass {
	err := rcStyleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RcStyleClass{native: rcStyleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRcStyleClass)
	return structGo
}
func finalizeRcStyleClass(obj *RcStyleClass) {
	rcStyleClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RecentActionClass) FieldParentClass() *ActionClass {
	argValue := gi.FieldGet(recentActionClassStruct, recv.native, "parent_class")
	value := &ActionClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RecentActionClass) SetFieldParentClass(value *ActionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(recentActionClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// RecentActionClassStruct creates an uninitialised RecentActionClass.
func RecentActionClassStruct() *RecentActionClass {
	err := recentActionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentActionClass{native: recentActionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentActionClass)
	return structGo
}
func finalizeRecentActionClass(obj *RecentActionClass) {
	recentActionClassStruct.Free(obj.native)
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

// RecentActionPrivateStruct creates an uninitialised RecentActionPrivate.
func RecentActionPrivateStruct() *RecentActionPrivate {
	err := recentActionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentActionPrivate{native: recentActionPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentActionPrivate)
	return structGo
}
func finalizeRecentActionPrivate(obj *RecentActionPrivate) {
	recentActionPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RecentChooserDialogClass) FieldParentClass() *DialogClass {
	argValue := gi.FieldGet(recentChooserDialogClassStruct, recv.native, "parent_class")
	value := &DialogClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RecentChooserDialogClass) SetFieldParentClass(value *DialogClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(recentChooserDialogClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// RecentChooserDialogClassStruct creates an uninitialised RecentChooserDialogClass.
func RecentChooserDialogClassStruct() *RecentChooserDialogClass {
	err := recentChooserDialogClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentChooserDialogClass{native: recentChooserDialogClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentChooserDialogClass)
	return structGo
}
func finalizeRecentChooserDialogClass(obj *RecentChooserDialogClass) {
	recentChooserDialogClassStruct.Free(obj.native)
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

// RecentChooserDialogPrivateStruct creates an uninitialised RecentChooserDialogPrivate.
func RecentChooserDialogPrivateStruct() *RecentChooserDialogPrivate {
	err := recentChooserDialogPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentChooserDialogPrivate{native: recentChooserDialogPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentChooserDialogPrivate)
	return structGo
}
func finalizeRecentChooserDialogPrivate(obj *RecentChooserDialogPrivate) {
	recentChooserDialogPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'set_current_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'set_current_uri' : for field setter : missing Type

// UNSUPPORTED : C value 'get_current_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'get_current_uri' : for field setter : missing Type

// UNSUPPORTED : C value 'select_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'select_uri' : for field setter : missing Type

// UNSUPPORTED : C value 'unselect_uri' : for field getter : missing Type

// UNSUPPORTED : C value 'unselect_uri' : for field setter : missing Type

// UNSUPPORTED : C value 'select_all' : for field getter : missing Type

// UNSUPPORTED : C value 'select_all' : for field setter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field getter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field setter : missing Type

// UNSUPPORTED : C value 'get_items' : for field getter : missing Type

// UNSUPPORTED : C value 'get_items' : for field setter : missing Type

// UNSUPPORTED : C value 'get_recent_manager' : for field getter : missing Type

// UNSUPPORTED : C value 'get_recent_manager' : for field setter : missing Type

// UNSUPPORTED : C value 'add_filter' : for field getter : missing Type

// UNSUPPORTED : C value 'add_filter' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_filter' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_filter' : for field setter : missing Type

// UNSUPPORTED : C value 'list_filters' : for field getter : missing Type

// UNSUPPORTED : C value 'list_filters' : for field setter : missing Type

// UNSUPPORTED : C value 'set_sort_func' : for field getter : missing Type

// UNSUPPORTED : C value 'set_sort_func' : for field setter : missing Type

// UNSUPPORTED : C value 'item_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'item_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'selection_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_changed' : for field setter : missing Type

// RecentChooserIfaceStruct creates an uninitialised RecentChooserIface.
func RecentChooserIfaceStruct() *RecentChooserIface {
	err := recentChooserIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentChooserIface{native: recentChooserIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentChooserIface)
	return structGo
}
func finalizeRecentChooserIface(obj *RecentChooserIface) {
	recentChooserIfaceStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RecentChooserMenuClass) FieldParentClass() *MenuClass {
	argValue := gi.FieldGet(recentChooserMenuClassStruct, recv.native, "parent_class")
	value := &MenuClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RecentChooserMenuClass) SetFieldParentClass(value *MenuClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(recentChooserMenuClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'gtk_recent1' : for field getter : missing Type

// UNSUPPORTED : C value 'gtk_recent1' : for field setter : missing Type

// UNSUPPORTED : C value 'gtk_recent2' : for field getter : missing Type

// UNSUPPORTED : C value 'gtk_recent2' : for field setter : missing Type

// UNSUPPORTED : C value 'gtk_recent3' : for field getter : missing Type

// UNSUPPORTED : C value 'gtk_recent3' : for field setter : missing Type

// UNSUPPORTED : C value 'gtk_recent4' : for field getter : missing Type

// UNSUPPORTED : C value 'gtk_recent4' : for field setter : missing Type

// RecentChooserMenuClassStruct creates an uninitialised RecentChooserMenuClass.
func RecentChooserMenuClassStruct() *RecentChooserMenuClass {
	err := recentChooserMenuClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentChooserMenuClass{native: recentChooserMenuClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentChooserMenuClass)
	return structGo
}
func finalizeRecentChooserMenuClass(obj *RecentChooserMenuClass) {
	recentChooserMenuClassStruct.Free(obj.native)
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

// RecentChooserMenuPrivateStruct creates an uninitialised RecentChooserMenuPrivate.
func RecentChooserMenuPrivateStruct() *RecentChooserMenuPrivate {
	err := recentChooserMenuPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentChooserMenuPrivate{native: recentChooserMenuPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentChooserMenuPrivate)
	return structGo
}
func finalizeRecentChooserMenuPrivate(obj *RecentChooserMenuPrivate) {
	recentChooserMenuPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RecentChooserWidgetClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(recentChooserWidgetClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RecentChooserWidgetClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(recentChooserWidgetClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// RecentChooserWidgetClassStruct creates an uninitialised RecentChooserWidgetClass.
func RecentChooserWidgetClassStruct() *RecentChooserWidgetClass {
	err := recentChooserWidgetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentChooserWidgetClass{native: recentChooserWidgetClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentChooserWidgetClass)
	return structGo
}
func finalizeRecentChooserWidgetClass(obj *RecentChooserWidgetClass) {
	recentChooserWidgetClassStruct.Free(obj.native)
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

// RecentChooserWidgetPrivateStruct creates an uninitialised RecentChooserWidgetPrivate.
func RecentChooserWidgetPrivateStruct() *RecentChooserWidgetPrivate {
	err := recentChooserWidgetPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentChooserWidgetPrivate{native: recentChooserWidgetPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentChooserWidgetPrivate)
	return structGo
}
func finalizeRecentChooserWidgetPrivate(obj *RecentChooserWidgetPrivate) {
	recentChooserWidgetPrivateStruct.Free(obj.native)
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

// FieldDisplayName returns the C field 'display_name'.
func (recv *RecentData) FieldDisplayName() string {
	argValue := gi.FieldGet(recentDataStruct, recv.native, "display_name")
	value := argValue.String(false)
	return value
}

// SetFieldDisplayName sets the value of the C field 'display_name'.
func (recv *RecentData) SetFieldDisplayName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(recentDataStruct, recv.native, "display_name", argValue)
}

// FieldDescription returns the C field 'description'.
func (recv *RecentData) FieldDescription() string {
	argValue := gi.FieldGet(recentDataStruct, recv.native, "description")
	value := argValue.String(false)
	return value
}

// SetFieldDescription sets the value of the C field 'description'.
func (recv *RecentData) SetFieldDescription(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(recentDataStruct, recv.native, "description", argValue)
}

// FieldMimeType returns the C field 'mime_type'.
func (recv *RecentData) FieldMimeType() string {
	argValue := gi.FieldGet(recentDataStruct, recv.native, "mime_type")
	value := argValue.String(false)
	return value
}

// SetFieldMimeType sets the value of the C field 'mime_type'.
func (recv *RecentData) SetFieldMimeType(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(recentDataStruct, recv.native, "mime_type", argValue)
}

// FieldAppName returns the C field 'app_name'.
func (recv *RecentData) FieldAppName() string {
	argValue := gi.FieldGet(recentDataStruct, recv.native, "app_name")
	value := argValue.String(false)
	return value
}

// SetFieldAppName sets the value of the C field 'app_name'.
func (recv *RecentData) SetFieldAppName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(recentDataStruct, recv.native, "app_name", argValue)
}

// FieldAppExec returns the C field 'app_exec'.
func (recv *RecentData) FieldAppExec() string {
	argValue := gi.FieldGet(recentDataStruct, recv.native, "app_exec")
	value := argValue.String(false)
	return value
}

// SetFieldAppExec sets the value of the C field 'app_exec'.
func (recv *RecentData) SetFieldAppExec(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(recentDataStruct, recv.native, "app_exec", argValue)
}

// UNSUPPORTED : C value 'groups' : for field getter : missing Type

// UNSUPPORTED : C value 'groups' : for field setter : missing Type

// FieldIsPrivate returns the C field 'is_private'.
func (recv *RecentData) FieldIsPrivate() bool {
	argValue := gi.FieldGet(recentDataStruct, recv.native, "is_private")
	value := argValue.Boolean()
	return value
}

// SetFieldIsPrivate sets the value of the C field 'is_private'.
func (recv *RecentData) SetFieldIsPrivate(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.FieldSet(recentDataStruct, recv.native, "is_private", argValue)
}

// RecentDataStruct creates an uninitialised RecentData.
func RecentDataStruct() *RecentData {
	err := recentDataStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentData{native: recentDataStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentData)
	return structGo
}
func finalizeRecentData(obj *RecentData) {
	recentDataStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'contains' : for field getter : no Go type for 'RecentFilterFlags'

// UNSUPPORTED : C value 'contains' : for field setter : no Go type for 'RecentFilterFlags'

// FieldUri returns the C field 'uri'.
func (recv *RecentFilterInfo) FieldUri() string {
	argValue := gi.FieldGet(recentFilterInfoStruct, recv.native, "uri")
	value := argValue.String(false)
	return value
}

// SetFieldUri sets the value of the C field 'uri'.
func (recv *RecentFilterInfo) SetFieldUri(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(recentFilterInfoStruct, recv.native, "uri", argValue)
}

// FieldDisplayName returns the C field 'display_name'.
func (recv *RecentFilterInfo) FieldDisplayName() string {
	argValue := gi.FieldGet(recentFilterInfoStruct, recv.native, "display_name")
	value := argValue.String(false)
	return value
}

// SetFieldDisplayName sets the value of the C field 'display_name'.
func (recv *RecentFilterInfo) SetFieldDisplayName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(recentFilterInfoStruct, recv.native, "display_name", argValue)
}

// FieldMimeType returns the C field 'mime_type'.
func (recv *RecentFilterInfo) FieldMimeType() string {
	argValue := gi.FieldGet(recentFilterInfoStruct, recv.native, "mime_type")
	value := argValue.String(false)
	return value
}

// SetFieldMimeType sets the value of the C field 'mime_type'.
func (recv *RecentFilterInfo) SetFieldMimeType(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(recentFilterInfoStruct, recv.native, "mime_type", argValue)
}

// UNSUPPORTED : C value 'applications' : for field getter : missing Type

// UNSUPPORTED : C value 'applications' : for field setter : missing Type

// UNSUPPORTED : C value 'groups' : for field getter : missing Type

// UNSUPPORTED : C value 'groups' : for field setter : missing Type

// FieldAge returns the C field 'age'.
func (recv *RecentFilterInfo) FieldAge() int32 {
	argValue := gi.FieldGet(recentFilterInfoStruct, recv.native, "age")
	value := argValue.Int32()
	return value
}

// SetFieldAge sets the value of the C field 'age'.
func (recv *RecentFilterInfo) SetFieldAge(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(recentFilterInfoStruct, recv.native, "age", argValue)
}

// RecentFilterInfoStruct creates an uninitialised RecentFilterInfo.
func RecentFilterInfoStruct() *RecentFilterInfo {
	err := recentFilterInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentFilterInfo{native: recentFilterInfoStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentFilterInfo)
	return structGo
}
func finalizeRecentFilterInfo(obj *RecentFilterInfo) {
	recentFilterInfoStruct.Free(obj.native)
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

// RecentInfoStruct creates an uninitialised RecentInfo.
func RecentInfoStruct() *RecentInfo {
	err := recentInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentInfo{native: recentInfoStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentInfo)
	return structGo
}
func finalizeRecentInfo(obj *RecentInfo) {
	recentInfoStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_recent1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_recent1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_recent2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_recent2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_recent3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_recent3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_recent4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_recent4' : for field setter : missing Type

// RecentManagerClassStruct creates an uninitialised RecentManagerClass.
func RecentManagerClassStruct() *RecentManagerClass {
	err := recentManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentManagerClass{native: recentManagerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentManagerClass)
	return structGo
}
func finalizeRecentManagerClass(obj *RecentManagerClass) {
	recentManagerClassStruct.Free(obj.native)
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

// RecentManagerPrivateStruct creates an uninitialised RecentManagerPrivate.
func RecentManagerPrivateStruct() *RecentManagerPrivate {
	err := recentManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RecentManagerPrivate{native: recentManagerPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRecentManagerPrivate)
	return structGo
}
func finalizeRecentManagerPrivate(obj *RecentManagerPrivate) {
	recentManagerPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RendererCellAccessibleClass) FieldParentClass() *CellAccessibleClass {
	argValue := gi.FieldGet(rendererCellAccessibleClassStruct, recv.native, "parent_class")
	value := &CellAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RendererCellAccessibleClass) SetFieldParentClass(value *CellAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(rendererCellAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// RendererCellAccessibleClassStruct creates an uninitialised RendererCellAccessibleClass.
func RendererCellAccessibleClassStruct() *RendererCellAccessibleClass {
	err := rendererCellAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RendererCellAccessibleClass{native: rendererCellAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRendererCellAccessibleClass)
	return structGo
}
func finalizeRendererCellAccessibleClass(obj *RendererCellAccessibleClass) {
	rendererCellAccessibleClassStruct.Free(obj.native)
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

// RendererCellAccessiblePrivateStruct creates an uninitialised RendererCellAccessiblePrivate.
func RendererCellAccessiblePrivateStruct() *RendererCellAccessiblePrivate {
	err := rendererCellAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RendererCellAccessiblePrivate{native: rendererCellAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRendererCellAccessiblePrivate)
	return structGo
}
func finalizeRendererCellAccessiblePrivate(obj *RendererCellAccessiblePrivate) {
	rendererCellAccessiblePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'data' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'data' : for field setter : no Go type for 'gpointer'

// FieldMinimumSize returns the C field 'minimum_size'.
func (recv *RequestedSize) FieldMinimumSize() int32 {
	argValue := gi.FieldGet(requestedSizeStruct, recv.native, "minimum_size")
	value := argValue.Int32()
	return value
}

// SetFieldMinimumSize sets the value of the C field 'minimum_size'.
func (recv *RequestedSize) SetFieldMinimumSize(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(requestedSizeStruct, recv.native, "minimum_size", argValue)
}

// FieldNaturalSize returns the C field 'natural_size'.
func (recv *RequestedSize) FieldNaturalSize() int32 {
	argValue := gi.FieldGet(requestedSizeStruct, recv.native, "natural_size")
	value := argValue.Int32()
	return value
}

// SetFieldNaturalSize sets the value of the C field 'natural_size'.
func (recv *RequestedSize) SetFieldNaturalSize(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(requestedSizeStruct, recv.native, "natural_size", argValue)
}

// RequestedSizeStruct creates an uninitialised RequestedSize.
func RequestedSizeStruct() *RequestedSize {
	err := requestedSizeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RequestedSize{native: requestedSizeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRequestedSize)
	return structGo
}
func finalizeRequestedSize(obj *RequestedSize) {
	requestedSizeStruct.Free(obj.native)
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

// FieldWidth returns the C field 'width'.
func (recv *Requisition) FieldWidth() int32 {
	argValue := gi.FieldGet(requisitionStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *Requisition) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(requisitionStruct, recv.native, "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *Requisition) FieldHeight() int32 {
	argValue := gi.FieldGet(requisitionStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *Requisition) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(requisitionStruct, recv.native, "height", argValue)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *RevealerClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(revealerClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *RevealerClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(revealerClassStruct, recv.native, "parent_class", argValue)
}

// RevealerClassStruct creates an uninitialised RevealerClass.
func RevealerClassStruct() *RevealerClass {
	err := revealerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &RevealerClass{native: revealerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeRevealerClass)
	return structGo
}
func finalizeRevealerClass(obj *RevealerClass) {
	revealerClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ScaleAccessibleClass) FieldParentClass() *RangeAccessibleClass {
	argValue := gi.FieldGet(scaleAccessibleClassStruct, recv.native, "parent_class")
	value := &RangeAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ScaleAccessibleClass) SetFieldParentClass(value *RangeAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(scaleAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ScaleAccessibleClassStruct creates an uninitialised ScaleAccessibleClass.
func ScaleAccessibleClassStruct() *ScaleAccessibleClass {
	err := scaleAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScaleAccessibleClass{native: scaleAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScaleAccessibleClass)
	return structGo
}
func finalizeScaleAccessibleClass(obj *ScaleAccessibleClass) {
	scaleAccessibleClassStruct.Free(obj.native)
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

// ScaleAccessiblePrivateStruct creates an uninitialised ScaleAccessiblePrivate.
func ScaleAccessiblePrivateStruct() *ScaleAccessiblePrivate {
	err := scaleAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScaleAccessiblePrivate{native: scaleAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScaleAccessiblePrivate)
	return structGo
}
func finalizeScaleAccessiblePrivate(obj *ScaleAccessiblePrivate) {
	scaleAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ScaleButtonAccessibleClass) FieldParentClass() *ButtonAccessibleClass {
	argValue := gi.FieldGet(scaleButtonAccessibleClassStruct, recv.native, "parent_class")
	value := &ButtonAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ScaleButtonAccessibleClass) SetFieldParentClass(value *ButtonAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(scaleButtonAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ScaleButtonAccessibleClassStruct creates an uninitialised ScaleButtonAccessibleClass.
func ScaleButtonAccessibleClassStruct() *ScaleButtonAccessibleClass {
	err := scaleButtonAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScaleButtonAccessibleClass{native: scaleButtonAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScaleButtonAccessibleClass)
	return structGo
}
func finalizeScaleButtonAccessibleClass(obj *ScaleButtonAccessibleClass) {
	scaleButtonAccessibleClassStruct.Free(obj.native)
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

// ScaleButtonAccessiblePrivateStruct creates an uninitialised ScaleButtonAccessiblePrivate.
func ScaleButtonAccessiblePrivateStruct() *ScaleButtonAccessiblePrivate {
	err := scaleButtonAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScaleButtonAccessiblePrivate{native: scaleButtonAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScaleButtonAccessiblePrivate)
	return structGo
}
func finalizeScaleButtonAccessiblePrivate(obj *ScaleButtonAccessiblePrivate) {
	scaleButtonAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ScaleButtonClass) FieldParentClass() *ButtonClass {
	argValue := gi.FieldGet(scaleButtonClassStruct, recv.native, "parent_class")
	value := &ButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ScaleButtonClass) SetFieldParentClass(value *ButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(scaleButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'value_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'value_changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ScaleButtonClassStruct creates an uninitialised ScaleButtonClass.
func ScaleButtonClassStruct() *ScaleButtonClass {
	err := scaleButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScaleButtonClass{native: scaleButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScaleButtonClass)
	return structGo
}
func finalizeScaleButtonClass(obj *ScaleButtonClass) {
	scaleButtonClassStruct.Free(obj.native)
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

// ScaleButtonPrivateStruct creates an uninitialised ScaleButtonPrivate.
func ScaleButtonPrivateStruct() *ScaleButtonPrivate {
	err := scaleButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScaleButtonPrivate{native: scaleButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScaleButtonPrivate)
	return structGo
}
func finalizeScaleButtonPrivate(obj *ScaleButtonPrivate) {
	scaleButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ScaleClass) FieldParentClass() *RangeClass {
	argValue := gi.FieldGet(scaleClassStruct, recv.native, "parent_class")
	value := &RangeClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ScaleClass) SetFieldParentClass(value *RangeClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(scaleClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'format_value' : for field getter : missing Type

// UNSUPPORTED : C value 'format_value' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_value' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_value' : for field setter : missing Type

// UNSUPPORTED : C value 'get_layout_offsets' : for field getter : missing Type

// UNSUPPORTED : C value 'get_layout_offsets' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ScaleClassStruct creates an uninitialised ScaleClass.
func ScaleClassStruct() *ScaleClass {
	err := scaleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScaleClass{native: scaleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScaleClass)
	return structGo
}
func finalizeScaleClass(obj *ScaleClass) {
	scaleClassStruct.Free(obj.native)
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

// ScalePrivateStruct creates an uninitialised ScalePrivate.
func ScalePrivateStruct() *ScalePrivate {
	err := scalePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScalePrivate{native: scalePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScalePrivate)
	return structGo
}
func finalizeScalePrivate(obj *ScalePrivate) {
	scalePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'base_iface' : for field getter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'base_iface' : for field setter : no Go type for 'GObject.TypeInterface'

// UNSUPPORTED : C value 'get_border' : for field getter : missing Type

// UNSUPPORTED : C value 'get_border' : for field setter : missing Type

// ScrollableInterfaceStruct creates an uninitialised ScrollableInterface.
func ScrollableInterfaceStruct() *ScrollableInterface {
	err := scrollableInterfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScrollableInterface{native: scrollableInterfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScrollableInterface)
	return structGo
}
func finalizeScrollableInterface(obj *ScrollableInterface) {
	scrollableInterfaceStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ScrollbarClass) FieldParentClass() *RangeClass {
	argValue := gi.FieldGet(scrollbarClassStruct, recv.native, "parent_class")
	value := &RangeClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ScrollbarClass) SetFieldParentClass(value *RangeClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(scrollbarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ScrollbarClassStruct creates an uninitialised ScrollbarClass.
func ScrollbarClassStruct() *ScrollbarClass {
	err := scrollbarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScrollbarClass{native: scrollbarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScrollbarClass)
	return structGo
}
func finalizeScrollbarClass(obj *ScrollbarClass) {
	scrollbarClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ScrolledWindowAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(scrolledWindowAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ScrolledWindowAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(scrolledWindowAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ScrolledWindowAccessibleClassStruct creates an uninitialised ScrolledWindowAccessibleClass.
func ScrolledWindowAccessibleClassStruct() *ScrolledWindowAccessibleClass {
	err := scrolledWindowAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScrolledWindowAccessibleClass{native: scrolledWindowAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScrolledWindowAccessibleClass)
	return structGo
}
func finalizeScrolledWindowAccessibleClass(obj *ScrolledWindowAccessibleClass) {
	scrolledWindowAccessibleClassStruct.Free(obj.native)
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

// ScrolledWindowAccessiblePrivateStruct creates an uninitialised ScrolledWindowAccessiblePrivate.
func ScrolledWindowAccessiblePrivateStruct() *ScrolledWindowAccessiblePrivate {
	err := scrolledWindowAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScrolledWindowAccessiblePrivate{native: scrolledWindowAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScrolledWindowAccessiblePrivate)
	return structGo
}
func finalizeScrolledWindowAccessiblePrivate(obj *ScrolledWindowAccessiblePrivate) {
	scrolledWindowAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ScrolledWindowClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(scrolledWindowClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ScrolledWindowClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(scrolledWindowClassStruct, recv.native, "parent_class", argValue)
}

// FieldScrollbarSpacing returns the C field 'scrollbar_spacing'.
func (recv *ScrolledWindowClass) FieldScrollbarSpacing() int32 {
	argValue := gi.FieldGet(scrolledWindowClassStruct, recv.native, "scrollbar_spacing")
	value := argValue.Int32()
	return value
}

// SetFieldScrollbarSpacing sets the value of the C field 'scrollbar_spacing'.
func (recv *ScrolledWindowClass) SetFieldScrollbarSpacing(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(scrolledWindowClassStruct, recv.native, "scrollbar_spacing", argValue)
}

// UNSUPPORTED : C value 'scroll_child' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_child' : for field setter : missing Type

// UNSUPPORTED : C value 'move_focus_out' : for field getter : missing Type

// UNSUPPORTED : C value 'move_focus_out' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ScrolledWindowClassStruct creates an uninitialised ScrolledWindowClass.
func ScrolledWindowClassStruct() *ScrolledWindowClass {
	err := scrolledWindowClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScrolledWindowClass{native: scrolledWindowClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScrolledWindowClass)
	return structGo
}
func finalizeScrolledWindowClass(obj *ScrolledWindowClass) {
	scrolledWindowClassStruct.Free(obj.native)
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

// ScrolledWindowPrivateStruct creates an uninitialised ScrolledWindowPrivate.
func ScrolledWindowPrivateStruct() *ScrolledWindowPrivate {
	err := scrolledWindowPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ScrolledWindowPrivate{native: scrolledWindowPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeScrolledWindowPrivate)
	return structGo
}
func finalizeScrolledWindowPrivate(obj *ScrolledWindowPrivate) {
	scrolledWindowPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SearchBarClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(searchBarClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SearchBarClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(searchBarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// SearchBarClassStruct creates an uninitialised SearchBarClass.
func SearchBarClassStruct() *SearchBarClass {
	err := searchBarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SearchBarClass{native: searchBarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSearchBarClass)
	return structGo
}
func finalizeSearchBarClass(obj *SearchBarClass) {
	searchBarClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SearchEntryClass) FieldParentClass() *EntryClass {
	argValue := gi.FieldGet(searchEntryClassStruct, recv.native, "parent_class")
	value := &EntryClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SearchEntryClass) SetFieldParentClass(value *EntryClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(searchEntryClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'search_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'search_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'next_match' : for field getter : missing Type

// UNSUPPORTED : C value 'next_match' : for field setter : missing Type

// UNSUPPORTED : C value 'previous_match' : for field getter : missing Type

// UNSUPPORTED : C value 'previous_match' : for field setter : missing Type

// UNSUPPORTED : C value 'stop_search' : for field getter : missing Type

// UNSUPPORTED : C value 'stop_search' : for field setter : missing Type

// SearchEntryClassStruct creates an uninitialised SearchEntryClass.
func SearchEntryClassStruct() *SearchEntryClass {
	err := searchEntryClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SearchEntryClass{native: searchEntryClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSearchEntryClass)
	return structGo
}
func finalizeSearchEntryClass(obj *SearchEntryClass) {
	searchEntryClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'gtk_selection_data_get_targets' : parameter 'targets' of type 'nil' not supported

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

// UNSUPPORTED : C value 'gtk_selection_data_set_uris' : parameter 'uris' of type 'nil' not supported

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

var selectionDataTargetsIncludeRichTextFunction *gi.Function
var selectionDataTargetsIncludeRichTextFunction_Once sync.Once

func selectionDataTargetsIncludeRichTextFunction_Set() error {
	var err error
	selectionDataTargetsIncludeRichTextFunction_Once.Do(func() {
		err = selectionDataStruct_Set()
		if err != nil {
			return
		}
		selectionDataTargetsIncludeRichTextFunction, err = selectionDataStruct.InvokerNew("targets_include_rich_text")
	})
	return err
}

// TargetsIncludeRichText is a representation of the C type gtk_selection_data_targets_include_rich_text.
func (recv *SelectionData) TargetsIncludeRichText(buffer *TextBuffer) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(buffer.native)

	var ret gi.Argument

	err := selectionDataTargetsIncludeRichTextFunction_Set()
	if err == nil {
		ret = selectionDataTargetsIncludeRichTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

// SelectionDataStruct creates an uninitialised SelectionData.
func SelectionDataStruct() *SelectionData {
	err := selectionDataStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SelectionData{native: selectionDataStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSelectionData)
	return structGo
}
func finalizeSelectionData(obj *SelectionData) {
	selectionDataStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SeparatorClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(separatorClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SeparatorClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(separatorClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// SeparatorClassStruct creates an uninitialised SeparatorClass.
func SeparatorClassStruct() *SeparatorClass {
	err := separatorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SeparatorClass{native: separatorClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSeparatorClass)
	return structGo
}
func finalizeSeparatorClass(obj *SeparatorClass) {
	separatorClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SeparatorMenuItemClass) FieldParentClass() *MenuItemClass {
	argValue := gi.FieldGet(separatorMenuItemClassStruct, recv.native, "parent_class")
	value := &MenuItemClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SeparatorMenuItemClass) SetFieldParentClass(value *MenuItemClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(separatorMenuItemClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// SeparatorMenuItemClassStruct creates an uninitialised SeparatorMenuItemClass.
func SeparatorMenuItemClassStruct() *SeparatorMenuItemClass {
	err := separatorMenuItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SeparatorMenuItemClass{native: separatorMenuItemClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSeparatorMenuItemClass)
	return structGo
}
func finalizeSeparatorMenuItemClass(obj *SeparatorMenuItemClass) {
	separatorMenuItemClassStruct.Free(obj.native)
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

// SeparatorPrivateStruct creates an uninitialised SeparatorPrivate.
func SeparatorPrivateStruct() *SeparatorPrivate {
	err := separatorPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SeparatorPrivate{native: separatorPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSeparatorPrivate)
	return structGo
}
func finalizeSeparatorPrivate(obj *SeparatorPrivate) {
	separatorPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SeparatorToolItemClass) FieldParentClass() *ToolItemClass {
	argValue := gi.FieldGet(separatorToolItemClassStruct, recv.native, "parent_class")
	value := &ToolItemClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SeparatorToolItemClass) SetFieldParentClass(value *ToolItemClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(separatorToolItemClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// SeparatorToolItemClassStruct creates an uninitialised SeparatorToolItemClass.
func SeparatorToolItemClassStruct() *SeparatorToolItemClass {
	err := separatorToolItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SeparatorToolItemClass{native: separatorToolItemClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSeparatorToolItemClass)
	return structGo
}
func finalizeSeparatorToolItemClass(obj *SeparatorToolItemClass) {
	separatorToolItemClassStruct.Free(obj.native)
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

// SeparatorToolItemPrivateStruct creates an uninitialised SeparatorToolItemPrivate.
func SeparatorToolItemPrivateStruct() *SeparatorToolItemPrivate {
	err := separatorToolItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SeparatorToolItemPrivate{native: separatorToolItemPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSeparatorToolItemPrivate)
	return structGo
}
func finalizeSeparatorToolItemPrivate(obj *SeparatorToolItemPrivate) {
	separatorToolItemPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// SettingsClassStruct creates an uninitialised SettingsClass.
func SettingsClassStruct() *SettingsClass {
	err := settingsClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SettingsClass{native: settingsClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSettingsClass)
	return structGo
}
func finalizeSettingsClass(obj *SettingsClass) {
	settingsClassStruct.Free(obj.native)
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

// SettingsPrivateStruct creates an uninitialised SettingsPrivate.
func SettingsPrivateStruct() *SettingsPrivate {
	err := settingsPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SettingsPrivate{native: settingsPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSettingsPrivate)
	return structGo
}
func finalizeSettingsPrivate(obj *SettingsPrivate) {
	settingsPrivateStruct.Free(obj.native)
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

// FieldOrigin returns the C field 'origin'.
func (recv *SettingsValue) FieldOrigin() string {
	argValue := gi.FieldGet(settingsValueStruct, recv.native, "origin")
	value := argValue.String(false)
	return value
}

// SetFieldOrigin sets the value of the C field 'origin'.
func (recv *SettingsValue) SetFieldOrigin(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(settingsValueStruct, recv.native, "origin", argValue)
}

// UNSUPPORTED : C value 'value' : for field getter : no Go type for 'GObject.Value'

// UNSUPPORTED : C value 'value' : for field setter : no Go type for 'GObject.Value'

// SettingsValueStruct creates an uninitialised SettingsValue.
func SettingsValueStruct() *SettingsValue {
	err := settingsValueStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SettingsValue{native: settingsValueStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSettingsValue)
	return structGo
}
func finalizeSettingsValue(obj *SettingsValue) {
	settingsValueStruct.Free(obj.native)
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

// ShortcutLabelClassStruct creates an uninitialised ShortcutLabelClass.
func ShortcutLabelClassStruct() *ShortcutLabelClass {
	err := shortcutLabelClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ShortcutLabelClass{native: shortcutLabelClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeShortcutLabelClass)
	return structGo
}
func finalizeShortcutLabelClass(obj *ShortcutLabelClass) {
	shortcutLabelClassStruct.Free(obj.native)
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

// ShortcutsGroupClassStruct creates an uninitialised ShortcutsGroupClass.
func ShortcutsGroupClassStruct() *ShortcutsGroupClass {
	err := shortcutsGroupClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ShortcutsGroupClass{native: shortcutsGroupClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeShortcutsGroupClass)
	return structGo
}
func finalizeShortcutsGroupClass(obj *ShortcutsGroupClass) {
	shortcutsGroupClassStruct.Free(obj.native)
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

// ShortcutsSectionClassStruct creates an uninitialised ShortcutsSectionClass.
func ShortcutsSectionClassStruct() *ShortcutsSectionClass {
	err := shortcutsSectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ShortcutsSectionClass{native: shortcutsSectionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeShortcutsSectionClass)
	return structGo
}
func finalizeShortcutsSectionClass(obj *ShortcutsSectionClass) {
	shortcutsSectionClassStruct.Free(obj.native)
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

// ShortcutsShortcutClassStruct creates an uninitialised ShortcutsShortcutClass.
func ShortcutsShortcutClassStruct() *ShortcutsShortcutClass {
	err := shortcutsShortcutClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ShortcutsShortcutClass{native: shortcutsShortcutClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeShortcutsShortcutClass)
	return structGo
}
func finalizeShortcutsShortcutClass(obj *ShortcutsShortcutClass) {
	shortcutsShortcutClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ShortcutsWindowClass) FieldParentClass() *WindowClass {
	argValue := gi.FieldGet(shortcutsWindowClassStruct, recv.native, "parent_class")
	value := &WindowClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ShortcutsWindowClass) SetFieldParentClass(value *WindowClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(shortcutsWindowClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'close' : for field getter : missing Type

// UNSUPPORTED : C value 'close' : for field setter : missing Type

// UNSUPPORTED : C value 'search' : for field getter : missing Type

// UNSUPPORTED : C value 'search' : for field setter : missing Type

// ShortcutsWindowClassStruct creates an uninitialised ShortcutsWindowClass.
func ShortcutsWindowClassStruct() *ShortcutsWindowClass {
	err := shortcutsWindowClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ShortcutsWindowClass{native: shortcutsWindowClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeShortcutsWindowClass)
	return structGo
}
func finalizeShortcutsWindowClass(obj *ShortcutsWindowClass) {
	shortcutsWindowClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// SizeGroupClassStruct creates an uninitialised SizeGroupClass.
func SizeGroupClassStruct() *SizeGroupClass {
	err := sizeGroupClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SizeGroupClass{native: sizeGroupClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSizeGroupClass)
	return structGo
}
func finalizeSizeGroupClass(obj *SizeGroupClass) {
	sizeGroupClassStruct.Free(obj.native)
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

// SizeGroupPrivateStruct creates an uninitialised SizeGroupPrivate.
func SizeGroupPrivateStruct() *SizeGroupPrivate {
	err := sizeGroupPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SizeGroupPrivate{native: sizeGroupPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSizeGroupPrivate)
	return structGo
}
func finalizeSizeGroupPrivate(obj *SizeGroupPrivate) {
	sizeGroupPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SocketClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(socketClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SocketClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(socketClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'plug_added' : for field getter : missing Type

// UNSUPPORTED : C value 'plug_added' : for field setter : missing Type

// UNSUPPORTED : C value 'plug_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'plug_removed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// SocketClassStruct creates an uninitialised SocketClass.
func SocketClassStruct() *SocketClass {
	err := socketClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketClass{native: socketClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSocketClass)
	return structGo
}
func finalizeSocketClass(obj *SocketClass) {
	socketClassStruct.Free(obj.native)
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

// SocketPrivateStruct creates an uninitialised SocketPrivate.
func SocketPrivateStruct() *SocketPrivate {
	err := socketPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SocketPrivate{native: socketPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSocketPrivate)
	return structGo
}
func finalizeSocketPrivate(obj *SocketPrivate) {
	socketPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SpinButtonAccessibleClass) FieldParentClass() *EntryAccessibleClass {
	argValue := gi.FieldGet(spinButtonAccessibleClassStruct, recv.native, "parent_class")
	value := &EntryAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SpinButtonAccessibleClass) SetFieldParentClass(value *EntryAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(spinButtonAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// SpinButtonAccessibleClassStruct creates an uninitialised SpinButtonAccessibleClass.
func SpinButtonAccessibleClassStruct() *SpinButtonAccessibleClass {
	err := spinButtonAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpinButtonAccessibleClass{native: spinButtonAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSpinButtonAccessibleClass)
	return structGo
}
func finalizeSpinButtonAccessibleClass(obj *SpinButtonAccessibleClass) {
	spinButtonAccessibleClassStruct.Free(obj.native)
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

// SpinButtonAccessiblePrivateStruct creates an uninitialised SpinButtonAccessiblePrivate.
func SpinButtonAccessiblePrivateStruct() *SpinButtonAccessiblePrivate {
	err := spinButtonAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpinButtonAccessiblePrivate{native: spinButtonAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSpinButtonAccessiblePrivate)
	return structGo
}
func finalizeSpinButtonAccessiblePrivate(obj *SpinButtonAccessiblePrivate) {
	spinButtonAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SpinButtonClass) FieldParentClass() *EntryClass {
	argValue := gi.FieldGet(spinButtonClassStruct, recv.native, "parent_class")
	value := &EntryClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SpinButtonClass) SetFieldParentClass(value *EntryClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(spinButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'input' : for field getter : missing Type

// UNSUPPORTED : C value 'input' : for field setter : missing Type

// UNSUPPORTED : C value 'output' : for field getter : missing Type

// UNSUPPORTED : C value 'output' : for field setter : missing Type

// UNSUPPORTED : C value 'value_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'value_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'change_value' : for field getter : missing Type

// UNSUPPORTED : C value 'change_value' : for field setter : missing Type

// UNSUPPORTED : C value 'wrapped' : for field getter : missing Type

// UNSUPPORTED : C value 'wrapped' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// SpinButtonClassStruct creates an uninitialised SpinButtonClass.
func SpinButtonClassStruct() *SpinButtonClass {
	err := spinButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpinButtonClass{native: spinButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSpinButtonClass)
	return structGo
}
func finalizeSpinButtonClass(obj *SpinButtonClass) {
	spinButtonClassStruct.Free(obj.native)
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

// SpinButtonPrivateStruct creates an uninitialised SpinButtonPrivate.
func SpinButtonPrivateStruct() *SpinButtonPrivate {
	err := spinButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpinButtonPrivate{native: spinButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSpinButtonPrivate)
	return structGo
}
func finalizeSpinButtonPrivate(obj *SpinButtonPrivate) {
	spinButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SpinnerAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(spinnerAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SpinnerAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(spinnerAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// SpinnerAccessibleClassStruct creates an uninitialised SpinnerAccessibleClass.
func SpinnerAccessibleClassStruct() *SpinnerAccessibleClass {
	err := spinnerAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpinnerAccessibleClass{native: spinnerAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSpinnerAccessibleClass)
	return structGo
}
func finalizeSpinnerAccessibleClass(obj *SpinnerAccessibleClass) {
	spinnerAccessibleClassStruct.Free(obj.native)
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

// SpinnerAccessiblePrivateStruct creates an uninitialised SpinnerAccessiblePrivate.
func SpinnerAccessiblePrivateStruct() *SpinnerAccessiblePrivate {
	err := spinnerAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpinnerAccessiblePrivate{native: spinnerAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSpinnerAccessiblePrivate)
	return structGo
}
func finalizeSpinnerAccessiblePrivate(obj *SpinnerAccessiblePrivate) {
	spinnerAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SpinnerClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(spinnerClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SpinnerClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(spinnerClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// SpinnerClassStruct creates an uninitialised SpinnerClass.
func SpinnerClassStruct() *SpinnerClass {
	err := spinnerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpinnerClass{native: spinnerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSpinnerClass)
	return structGo
}
func finalizeSpinnerClass(obj *SpinnerClass) {
	spinnerClassStruct.Free(obj.native)
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

// SpinnerPrivateStruct creates an uninitialised SpinnerPrivate.
func SpinnerPrivateStruct() *SpinnerPrivate {
	err := spinnerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SpinnerPrivate{native: spinnerPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSpinnerPrivate)
	return structGo
}
func finalizeSpinnerPrivate(obj *SpinnerPrivate) {
	spinnerPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *StackAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(stackAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *StackAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(stackAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// StackAccessibleClassStruct creates an uninitialised StackAccessibleClass.
func StackAccessibleClassStruct() *StackAccessibleClass {
	err := stackAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StackAccessibleClass{native: stackAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStackAccessibleClass)
	return structGo
}
func finalizeStackAccessibleClass(obj *StackAccessibleClass) {
	stackAccessibleClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *StackClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(stackClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *StackClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(stackClassStruct, recv.native, "parent_class", argValue)
}

// StackClassStruct creates an uninitialised StackClass.
func StackClassStruct() *StackClass {
	err := stackClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StackClass{native: stackClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStackClass)
	return structGo
}
func finalizeStackClass(obj *StackClass) {
	stackClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *StackSidebarClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(stackSidebarClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *StackSidebarClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(stackSidebarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// StackSidebarClassStruct creates an uninitialised StackSidebarClass.
func StackSidebarClassStruct() *StackSidebarClass {
	err := stackSidebarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StackSidebarClass{native: stackSidebarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStackSidebarClass)
	return structGo
}
func finalizeStackSidebarClass(obj *StackSidebarClass) {
	stackSidebarClassStruct.Free(obj.native)
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

// StackSidebarPrivateStruct creates an uninitialised StackSidebarPrivate.
func StackSidebarPrivateStruct() *StackSidebarPrivate {
	err := stackSidebarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StackSidebarPrivate{native: stackSidebarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStackSidebarPrivate)
	return structGo
}
func finalizeStackSidebarPrivate(obj *StackSidebarPrivate) {
	stackSidebarPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *StackSwitcherClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(stackSwitcherClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *StackSwitcherClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(stackSwitcherClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// StackSwitcherClassStruct creates an uninitialised StackSwitcherClass.
func StackSwitcherClassStruct() *StackSwitcherClass {
	err := stackSwitcherClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StackSwitcherClass{native: stackSwitcherClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStackSwitcherClass)
	return structGo
}
func finalizeStackSwitcherClass(obj *StackSwitcherClass) {
	stackSwitcherClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value 'popup_menu' : for field getter : missing Type

// UNSUPPORTED : C value 'popup_menu' : for field setter : missing Type

// UNSUPPORTED : C value 'size_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'size_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'button_press_event' : for field getter : missing Type

// UNSUPPORTED : C value 'button_press_event' : for field setter : missing Type

// UNSUPPORTED : C value 'button_release_event' : for field getter : missing Type

// UNSUPPORTED : C value 'button_release_event' : for field setter : missing Type

// UNSUPPORTED : C value 'scroll_event' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_event' : for field setter : missing Type

// UNSUPPORTED : C value 'query_tooltip' : for field getter : missing Type

// UNSUPPORTED : C value 'query_tooltip' : for field setter : missing Type

// UNSUPPORTED : C value '__gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '__gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '__gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '__gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '__gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '__gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '__gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '__gtk_reserved4' : for field setter : missing Type

// StatusIconClassStruct creates an uninitialised StatusIconClass.
func StatusIconClassStruct() *StatusIconClass {
	err := statusIconClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StatusIconClass{native: statusIconClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStatusIconClass)
	return structGo
}
func finalizeStatusIconClass(obj *StatusIconClass) {
	statusIconClassStruct.Free(obj.native)
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

// StatusIconPrivateStruct creates an uninitialised StatusIconPrivate.
func StatusIconPrivateStruct() *StatusIconPrivate {
	err := statusIconPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StatusIconPrivate{native: statusIconPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStatusIconPrivate)
	return structGo
}
func finalizeStatusIconPrivate(obj *StatusIconPrivate) {
	statusIconPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *StatusbarAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(statusbarAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *StatusbarAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(statusbarAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// StatusbarAccessibleClassStruct creates an uninitialised StatusbarAccessibleClass.
func StatusbarAccessibleClassStruct() *StatusbarAccessibleClass {
	err := statusbarAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StatusbarAccessibleClass{native: statusbarAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStatusbarAccessibleClass)
	return structGo
}
func finalizeStatusbarAccessibleClass(obj *StatusbarAccessibleClass) {
	statusbarAccessibleClassStruct.Free(obj.native)
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

// StatusbarAccessiblePrivateStruct creates an uninitialised StatusbarAccessiblePrivate.
func StatusbarAccessiblePrivateStruct() *StatusbarAccessiblePrivate {
	err := statusbarAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StatusbarAccessiblePrivate{native: statusbarAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStatusbarAccessiblePrivate)
	return structGo
}
func finalizeStatusbarAccessiblePrivate(obj *StatusbarAccessiblePrivate) {
	statusbarAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *StatusbarClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(statusbarClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *StatusbarClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(statusbarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'reserved' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'reserved' : for field setter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'text_pushed' : for field getter : missing Type

// UNSUPPORTED : C value 'text_pushed' : for field setter : missing Type

// UNSUPPORTED : C value 'text_popped' : for field getter : missing Type

// UNSUPPORTED : C value 'text_popped' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// StatusbarClassStruct creates an uninitialised StatusbarClass.
func StatusbarClassStruct() *StatusbarClass {
	err := statusbarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StatusbarClass{native: statusbarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStatusbarClass)
	return structGo
}
func finalizeStatusbarClass(obj *StatusbarClass) {
	statusbarClassStruct.Free(obj.native)
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

// StatusbarPrivateStruct creates an uninitialised StatusbarPrivate.
func StatusbarPrivateStruct() *StatusbarPrivate {
	err := statusbarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StatusbarPrivate{native: statusbarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStatusbarPrivate)
	return structGo
}
func finalizeStatusbarPrivate(obj *StatusbarPrivate) {
	statusbarPrivateStruct.Free(obj.native)
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

// FieldStockId returns the C field 'stock_id'.
func (recv *StockItem) FieldStockId() string {
	argValue := gi.FieldGet(stockItemStruct, recv.native, "stock_id")
	value := argValue.String(false)
	return value
}

// SetFieldStockId sets the value of the C field 'stock_id'.
func (recv *StockItem) SetFieldStockId(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(stockItemStruct, recv.native, "stock_id", argValue)
}

// FieldLabel returns the C field 'label'.
func (recv *StockItem) FieldLabel() string {
	argValue := gi.FieldGet(stockItemStruct, recv.native, "label")
	value := argValue.String(false)
	return value
}

// SetFieldLabel sets the value of the C field 'label'.
func (recv *StockItem) SetFieldLabel(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(stockItemStruct, recv.native, "label", argValue)
}

// UNSUPPORTED : C value 'modifier' : for field getter : no Go type for 'Gdk.ModifierType'

// UNSUPPORTED : C value 'modifier' : for field setter : no Go type for 'Gdk.ModifierType'

// FieldKeyval returns the C field 'keyval'.
func (recv *StockItem) FieldKeyval() uint32 {
	argValue := gi.FieldGet(stockItemStruct, recv.native, "keyval")
	value := argValue.Uint32()
	return value
}

// SetFieldKeyval sets the value of the C field 'keyval'.
func (recv *StockItem) SetFieldKeyval(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(stockItemStruct, recv.native, "keyval", argValue)
}

// FieldTranslationDomain returns the C field 'translation_domain'.
func (recv *StockItem) FieldTranslationDomain() string {
	argValue := gi.FieldGet(stockItemStruct, recv.native, "translation_domain")
	value := argValue.String(false)
	return value
}

// SetFieldTranslationDomain sets the value of the C field 'translation_domain'.
func (recv *StockItem) SetFieldTranslationDomain(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(stockItemStruct, recv.native, "translation_domain", argValue)
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

// StockItemStruct creates an uninitialised StockItem.
func StockItemStruct() *StockItem {
	err := stockItemStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StockItem{native: stockItemStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStockItem)
	return structGo
}
func finalizeStockItem(obj *StockItem) {
	stockItemStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'realize' : for field getter : missing Type

// UNSUPPORTED : C value 'realize' : for field setter : missing Type

// UNSUPPORTED : C value 'unrealize' : for field getter : missing Type

// UNSUPPORTED : C value 'unrealize' : for field setter : missing Type

// UNSUPPORTED : C value 'copy' : for field getter : missing Type

// UNSUPPORTED : C value 'copy' : for field setter : missing Type

// UNSUPPORTED : C value 'clone' : for field getter : missing Type

// UNSUPPORTED : C value 'clone' : for field setter : missing Type

// UNSUPPORTED : C value 'init_from_rc' : for field getter : missing Type

// UNSUPPORTED : C value 'init_from_rc' : for field setter : missing Type

// UNSUPPORTED : C value 'set_background' : for field getter : missing Type

// UNSUPPORTED : C value 'set_background' : for field setter : missing Type

// UNSUPPORTED : C value 'render_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'render_icon' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_hline' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_hline' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_vline' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_vline' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_shadow' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_shadow' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_arrow' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_arrow' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_diamond' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_diamond' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_box' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_box' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_flat_box' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_flat_box' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_check' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_check' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_option' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_option' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_tab' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_tab' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_shadow_gap' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_shadow_gap' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_box_gap' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_box_gap' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_extension' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_extension' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_slider' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_slider' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_handle' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_handle' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_expander' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_expander' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_layout' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_layout' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_resize_grip' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_resize_grip' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_spinner' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_spinner' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved9' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved9' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved10' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved10' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved11' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved11' : for field setter : missing Type

// StyleClassStruct creates an uninitialised StyleClass.
func StyleClassStruct() *StyleClass {
	err := styleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleClass{native: styleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStyleClass)
	return structGo
}
func finalizeStyleClass(obj *StyleClass) {
	styleClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// StyleContextClassStruct creates an uninitialised StyleContextClass.
func StyleContextClassStruct() *StyleContextClass {
	err := styleContextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleContextClass{native: styleContextClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStyleContextClass)
	return structGo
}
func finalizeStyleContextClass(obj *StyleContextClass) {
	styleContextClassStruct.Free(obj.native)
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

// StyleContextPrivateStruct creates an uninitialised StyleContextPrivate.
func StyleContextPrivateStruct() *StyleContextPrivate {
	err := styleContextPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleContextPrivate{native: styleContextPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStyleContextPrivate)
	return structGo
}
func finalizeStyleContextPrivate(obj *StyleContextPrivate) {
	styleContextPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// StylePropertiesClassStruct creates an uninitialised StylePropertiesClass.
func StylePropertiesClassStruct() *StylePropertiesClass {
	err := stylePropertiesClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StylePropertiesClass{native: stylePropertiesClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStylePropertiesClass)
	return structGo
}
func finalizeStylePropertiesClass(obj *StylePropertiesClass) {
	stylePropertiesClassStruct.Free(obj.native)
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

// StylePropertiesPrivateStruct creates an uninitialised StylePropertiesPrivate.
func StylePropertiesPrivateStruct() *StylePropertiesPrivate {
	err := stylePropertiesPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StylePropertiesPrivate{native: stylePropertiesPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStylePropertiesPrivate)
	return structGo
}
func finalizeStylePropertiesPrivate(obj *StylePropertiesPrivate) {
	stylePropertiesPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'get_style' : for field getter : missing Type

// UNSUPPORTED : C value 'get_style' : for field setter : missing Type

// UNSUPPORTED : C value 'get_style_property' : for field getter : missing Type

// UNSUPPORTED : C value 'get_style_property' : for field setter : missing Type

// UNSUPPORTED : C value 'get_icon_factory' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon_factory' : for field setter : missing Type

// StyleProviderIfaceStruct creates an uninitialised StyleProviderIface.
func StyleProviderIfaceStruct() *StyleProviderIface {
	err := styleProviderIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &StyleProviderIface{native: styleProviderIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeStyleProviderIface)
	return structGo
}
func finalizeStyleProviderIface(obj *StyleProviderIface) {
	styleProviderIfaceStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SwitchAccessibleClass) FieldParentClass() *WidgetAccessibleClass {
	argValue := gi.FieldGet(switchAccessibleClassStruct, recv.native, "parent_class")
	value := &WidgetAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SwitchAccessibleClass) SetFieldParentClass(value *WidgetAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(switchAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// SwitchAccessibleClassStruct creates an uninitialised SwitchAccessibleClass.
func SwitchAccessibleClassStruct() *SwitchAccessibleClass {
	err := switchAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SwitchAccessibleClass{native: switchAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSwitchAccessibleClass)
	return structGo
}
func finalizeSwitchAccessibleClass(obj *SwitchAccessibleClass) {
	switchAccessibleClassStruct.Free(obj.native)
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

// SwitchAccessiblePrivateStruct creates an uninitialised SwitchAccessiblePrivate.
func SwitchAccessiblePrivateStruct() *SwitchAccessiblePrivate {
	err := switchAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SwitchAccessiblePrivate{native: switchAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSwitchAccessiblePrivate)
	return structGo
}
func finalizeSwitchAccessiblePrivate(obj *SwitchAccessiblePrivate) {
	switchAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *SwitchClass) FieldParentClass() *WidgetClass {
	argValue := gi.FieldGet(switchClassStruct, recv.native, "parent_class")
	value := &WidgetClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *SwitchClass) SetFieldParentClass(value *WidgetClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(switchClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'activate' : for field getter : missing Type

// UNSUPPORTED : C value 'activate' : for field setter : missing Type

// UNSUPPORTED : C value 'state_set' : for field getter : missing Type

// UNSUPPORTED : C value 'state_set' : for field setter : missing Type

// UNSUPPORTED : C value '_switch_padding_1' : for field getter : missing Type

// UNSUPPORTED : C value '_switch_padding_1' : for field setter : missing Type

// UNSUPPORTED : C value '_switch_padding_2' : for field getter : missing Type

// UNSUPPORTED : C value '_switch_padding_2' : for field setter : missing Type

// UNSUPPORTED : C value '_switch_padding_3' : for field getter : missing Type

// UNSUPPORTED : C value '_switch_padding_3' : for field setter : missing Type

// UNSUPPORTED : C value '_switch_padding_4' : for field getter : missing Type

// UNSUPPORTED : C value '_switch_padding_4' : for field setter : missing Type

// UNSUPPORTED : C value '_switch_padding_5' : for field getter : missing Type

// UNSUPPORTED : C value '_switch_padding_5' : for field setter : missing Type

// SwitchClassStruct creates an uninitialised SwitchClass.
func SwitchClassStruct() *SwitchClass {
	err := switchClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SwitchClass{native: switchClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSwitchClass)
	return structGo
}
func finalizeSwitchClass(obj *SwitchClass) {
	switchClassStruct.Free(obj.native)
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

// SwitchPrivateStruct creates an uninitialised SwitchPrivate.
func SwitchPrivateStruct() *SwitchPrivate {
	err := switchPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &SwitchPrivate{native: switchPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeSwitchPrivate)
	return structGo
}
func finalizeSwitchPrivate(obj *SwitchPrivate) {
	switchPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'gtk_symbolic_color_resolve' : parameter 'resolved_color' of type 'Gdk.RGBA' not supported

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

// FieldWidget returns the C field 'widget'.
func (recv *TableChild) FieldWidget() *Widget {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "widget")
	value := &Widget{native: argValue.Pointer()}
	return value
}

// SetFieldWidget sets the value of the C field 'widget'.
func (recv *TableChild) SetFieldWidget(value *Widget) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(tableChildStruct, recv.native, "widget", argValue)
}

// FieldLeftAttach returns the C field 'left_attach'.
func (recv *TableChild) FieldLeftAttach() uint16 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "left_attach")
	value := argValue.Uint16()
	return value
}

// SetFieldLeftAttach sets the value of the C field 'left_attach'.
func (recv *TableChild) SetFieldLeftAttach(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.FieldSet(tableChildStruct, recv.native, "left_attach", argValue)
}

// FieldRightAttach returns the C field 'right_attach'.
func (recv *TableChild) FieldRightAttach() uint16 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "right_attach")
	value := argValue.Uint16()
	return value
}

// SetFieldRightAttach sets the value of the C field 'right_attach'.
func (recv *TableChild) SetFieldRightAttach(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.FieldSet(tableChildStruct, recv.native, "right_attach", argValue)
}

// FieldTopAttach returns the C field 'top_attach'.
func (recv *TableChild) FieldTopAttach() uint16 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "top_attach")
	value := argValue.Uint16()
	return value
}

// SetFieldTopAttach sets the value of the C field 'top_attach'.
func (recv *TableChild) SetFieldTopAttach(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.FieldSet(tableChildStruct, recv.native, "top_attach", argValue)
}

// FieldBottomAttach returns the C field 'bottom_attach'.
func (recv *TableChild) FieldBottomAttach() uint16 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "bottom_attach")
	value := argValue.Uint16()
	return value
}

// SetFieldBottomAttach sets the value of the C field 'bottom_attach'.
func (recv *TableChild) SetFieldBottomAttach(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.FieldSet(tableChildStruct, recv.native, "bottom_attach", argValue)
}

// FieldXpadding returns the C field 'xpadding'.
func (recv *TableChild) FieldXpadding() uint16 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "xpadding")
	value := argValue.Uint16()
	return value
}

// SetFieldXpadding sets the value of the C field 'xpadding'.
func (recv *TableChild) SetFieldXpadding(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.FieldSet(tableChildStruct, recv.native, "xpadding", argValue)
}

// FieldYpadding returns the C field 'ypadding'.
func (recv *TableChild) FieldYpadding() uint16 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "ypadding")
	value := argValue.Uint16()
	return value
}

// SetFieldYpadding sets the value of the C field 'ypadding'.
func (recv *TableChild) SetFieldYpadding(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.FieldSet(tableChildStruct, recv.native, "ypadding", argValue)
}

// FieldXexpand returns the C field 'xexpand'.
func (recv *TableChild) FieldXexpand() uint32 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "xexpand")
	value := argValue.Uint32()
	return value
}

// SetFieldXexpand sets the value of the C field 'xexpand'.
func (recv *TableChild) SetFieldXexpand(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableChildStruct, recv.native, "xexpand", argValue)
}

// FieldYexpand returns the C field 'yexpand'.
func (recv *TableChild) FieldYexpand() uint32 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "yexpand")
	value := argValue.Uint32()
	return value
}

// SetFieldYexpand sets the value of the C field 'yexpand'.
func (recv *TableChild) SetFieldYexpand(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableChildStruct, recv.native, "yexpand", argValue)
}

// FieldXshrink returns the C field 'xshrink'.
func (recv *TableChild) FieldXshrink() uint32 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "xshrink")
	value := argValue.Uint32()
	return value
}

// SetFieldXshrink sets the value of the C field 'xshrink'.
func (recv *TableChild) SetFieldXshrink(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableChildStruct, recv.native, "xshrink", argValue)
}

// FieldYshrink returns the C field 'yshrink'.
func (recv *TableChild) FieldYshrink() uint32 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "yshrink")
	value := argValue.Uint32()
	return value
}

// SetFieldYshrink sets the value of the C field 'yshrink'.
func (recv *TableChild) SetFieldYshrink(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableChildStruct, recv.native, "yshrink", argValue)
}

// FieldXfill returns the C field 'xfill'.
func (recv *TableChild) FieldXfill() uint32 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "xfill")
	value := argValue.Uint32()
	return value
}

// SetFieldXfill sets the value of the C field 'xfill'.
func (recv *TableChild) SetFieldXfill(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableChildStruct, recv.native, "xfill", argValue)
}

// FieldYfill returns the C field 'yfill'.
func (recv *TableChild) FieldYfill() uint32 {
	argValue := gi.FieldGet(tableChildStruct, recv.native, "yfill")
	value := argValue.Uint32()
	return value
}

// SetFieldYfill sets the value of the C field 'yfill'.
func (recv *TableChild) SetFieldYfill(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableChildStruct, recv.native, "yfill", argValue)
}

// TableChildStruct creates an uninitialised TableChild.
func TableChildStruct() *TableChild {
	err := tableChildStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TableChild{native: tableChildStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTableChild)
	return structGo
}
func finalizeTableChild(obj *TableChild) {
	tableChildStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *TableClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(tableClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TableClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(tableClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TableClassStruct creates an uninitialised TableClass.
func TableClassStruct() *TableClass {
	err := tableClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TableClass{native: tableClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTableClass)
	return structGo
}
func finalizeTableClass(obj *TableClass) {
	tableClassStruct.Free(obj.native)
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

// TablePrivateStruct creates an uninitialised TablePrivate.
func TablePrivateStruct() *TablePrivate {
	err := tablePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TablePrivate{native: tablePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTablePrivate)
	return structGo
}
func finalizeTablePrivate(obj *TablePrivate) {
	tablePrivateStruct.Free(obj.native)
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

// FieldRequisition returns the C field 'requisition'.
func (recv *TableRowCol) FieldRequisition() uint16 {
	argValue := gi.FieldGet(tableRowColStruct, recv.native, "requisition")
	value := argValue.Uint16()
	return value
}

// SetFieldRequisition sets the value of the C field 'requisition'.
func (recv *TableRowCol) SetFieldRequisition(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.FieldSet(tableRowColStruct, recv.native, "requisition", argValue)
}

// FieldAllocation returns the C field 'allocation'.
func (recv *TableRowCol) FieldAllocation() uint16 {
	argValue := gi.FieldGet(tableRowColStruct, recv.native, "allocation")
	value := argValue.Uint16()
	return value
}

// SetFieldAllocation sets the value of the C field 'allocation'.
func (recv *TableRowCol) SetFieldAllocation(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.FieldSet(tableRowColStruct, recv.native, "allocation", argValue)
}

// FieldSpacing returns the C field 'spacing'.
func (recv *TableRowCol) FieldSpacing() uint16 {
	argValue := gi.FieldGet(tableRowColStruct, recv.native, "spacing")
	value := argValue.Uint16()
	return value
}

// SetFieldSpacing sets the value of the C field 'spacing'.
func (recv *TableRowCol) SetFieldSpacing(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.FieldSet(tableRowColStruct, recv.native, "spacing", argValue)
}

// FieldNeedExpand returns the C field 'need_expand'.
func (recv *TableRowCol) FieldNeedExpand() uint32 {
	argValue := gi.FieldGet(tableRowColStruct, recv.native, "need_expand")
	value := argValue.Uint32()
	return value
}

// SetFieldNeedExpand sets the value of the C field 'need_expand'.
func (recv *TableRowCol) SetFieldNeedExpand(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableRowColStruct, recv.native, "need_expand", argValue)
}

// FieldNeedShrink returns the C field 'need_shrink'.
func (recv *TableRowCol) FieldNeedShrink() uint32 {
	argValue := gi.FieldGet(tableRowColStruct, recv.native, "need_shrink")
	value := argValue.Uint32()
	return value
}

// SetFieldNeedShrink sets the value of the C field 'need_shrink'.
func (recv *TableRowCol) SetFieldNeedShrink(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableRowColStruct, recv.native, "need_shrink", argValue)
}

// FieldExpand returns the C field 'expand'.
func (recv *TableRowCol) FieldExpand() uint32 {
	argValue := gi.FieldGet(tableRowColStruct, recv.native, "expand")
	value := argValue.Uint32()
	return value
}

// SetFieldExpand sets the value of the C field 'expand'.
func (recv *TableRowCol) SetFieldExpand(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableRowColStruct, recv.native, "expand", argValue)
}

// FieldShrink returns the C field 'shrink'.
func (recv *TableRowCol) FieldShrink() uint32 {
	argValue := gi.FieldGet(tableRowColStruct, recv.native, "shrink")
	value := argValue.Uint32()
	return value
}

// SetFieldShrink sets the value of the C field 'shrink'.
func (recv *TableRowCol) SetFieldShrink(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableRowColStruct, recv.native, "shrink", argValue)
}

// FieldEmpty returns the C field 'empty'.
func (recv *TableRowCol) FieldEmpty() uint32 {
	argValue := gi.FieldGet(tableRowColStruct, recv.native, "empty")
	value := argValue.Uint32()
	return value
}

// SetFieldEmpty sets the value of the C field 'empty'.
func (recv *TableRowCol) SetFieldEmpty(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(tableRowColStruct, recv.native, "empty", argValue)
}

// TableRowColStruct creates an uninitialised TableRowCol.
func TableRowColStruct() *TableRowCol {
	err := tableRowColStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TableRowCol{native: tableRowColStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTableRowCol)
	return structGo
}
func finalizeTableRowCol(obj *TableRowCol) {
	tableRowColStruct.Free(obj.native)
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

// FieldTarget returns the C field 'target'.
func (recv *TargetEntry) FieldTarget() string {
	argValue := gi.FieldGet(targetEntryStruct, recv.native, "target")
	value := argValue.String(false)
	return value
}

// SetFieldTarget sets the value of the C field 'target'.
func (recv *TargetEntry) SetFieldTarget(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(targetEntryStruct, recv.native, "target", argValue)
}

// FieldFlags returns the C field 'flags'.
func (recv *TargetEntry) FieldFlags() uint32 {
	argValue := gi.FieldGet(targetEntryStruct, recv.native, "flags")
	value := argValue.Uint32()
	return value
}

// SetFieldFlags sets the value of the C field 'flags'.
func (recv *TargetEntry) SetFieldFlags(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(targetEntryStruct, recv.native, "flags", argValue)
}

// FieldInfo returns the C field 'info'.
func (recv *TargetEntry) FieldInfo() uint32 {
	argValue := gi.FieldGet(targetEntryStruct, recv.native, "info")
	value := argValue.Uint32()
	return value
}

// SetFieldInfo sets the value of the C field 'info'.
func (recv *TargetEntry) SetFieldInfo(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(targetEntryStruct, recv.native, "info", argValue)
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

// UNSUPPORTED : C value 'gtk_target_list_new' : parameter 'targets' of type 'nil' not supported

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

var targetListAddRichTextTargetsFunction *gi.Function
var targetListAddRichTextTargetsFunction_Once sync.Once

func targetListAddRichTextTargetsFunction_Set() error {
	var err error
	targetListAddRichTextTargetsFunction_Once.Do(func() {
		err = targetListStruct_Set()
		if err != nil {
			return
		}
		targetListAddRichTextTargetsFunction, err = targetListStruct.InvokerNew("add_rich_text_targets")
	})
	return err
}

// AddRichTextTargets is a representation of the C type gtk_target_list_add_rich_text_targets.
func (recv *TargetList) AddRichTextTargets(info uint32, deserializable bool, buffer *TextBuffer) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetUint32(info)
	inArgs[2].SetBoolean(deserializable)
	inArgs[3].SetPointer(buffer.native)

	err := targetListAddRichTextTargetsFunction_Set()
	if err == nil {
		targetListAddRichTextTargetsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_target_list_add_table' : parameter 'targets' of type 'nil' not supported

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

// UNSUPPORTED : C value 'target' : for field getter : no Go type for 'Gdk.Atom'

// UNSUPPORTED : C value 'target' : for field setter : no Go type for 'Gdk.Atom'

// FieldFlags returns the C field 'flags'.
func (recv *TargetPair) FieldFlags() uint32 {
	argValue := gi.FieldGet(targetPairStruct, recv.native, "flags")
	value := argValue.Uint32()
	return value
}

// SetFieldFlags sets the value of the C field 'flags'.
func (recv *TargetPair) SetFieldFlags(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(targetPairStruct, recv.native, "flags", argValue)
}

// FieldInfo returns the C field 'info'.
func (recv *TargetPair) FieldInfo() uint32 {
	argValue := gi.FieldGet(targetPairStruct, recv.native, "info")
	value := argValue.Uint32()
	return value
}

// SetFieldInfo sets the value of the C field 'info'.
func (recv *TargetPair) SetFieldInfo(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(targetPairStruct, recv.native, "info", argValue)
}

// TargetPairStruct creates an uninitialised TargetPair.
func TargetPairStruct() *TargetPair {
	err := targetPairStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TargetPair{native: targetPairStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTargetPair)
	return structGo
}
func finalizeTargetPair(obj *TargetPair) {
	targetPairStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *TearoffMenuItemClass) FieldParentClass() *MenuItemClass {
	argValue := gi.FieldGet(tearoffMenuItemClassStruct, recv.native, "parent_class")
	value := &MenuItemClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TearoffMenuItemClass) SetFieldParentClass(value *MenuItemClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(tearoffMenuItemClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TearoffMenuItemClassStruct creates an uninitialised TearoffMenuItemClass.
func TearoffMenuItemClassStruct() *TearoffMenuItemClass {
	err := tearoffMenuItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TearoffMenuItemClass{native: tearoffMenuItemClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTearoffMenuItemClass)
	return structGo
}
func finalizeTearoffMenuItemClass(obj *TearoffMenuItemClass) {
	tearoffMenuItemClassStruct.Free(obj.native)
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

// TearoffMenuItemPrivateStruct creates an uninitialised TearoffMenuItemPrivate.
func TearoffMenuItemPrivateStruct() *TearoffMenuItemPrivate {
	err := tearoffMenuItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TearoffMenuItemPrivate{native: tearoffMenuItemPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTearoffMenuItemPrivate)
	return structGo
}
func finalizeTearoffMenuItemPrivate(obj *TearoffMenuItemPrivate) {
	tearoffMenuItemPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'bg_color' : for field getter : no Go type for 'Gdk.Color'

// UNSUPPORTED : C value 'bg_color' : for field setter : no Go type for 'Gdk.Color'

// UNSUPPORTED : C value 'fg_color' : for field getter : no Go type for 'Gdk.Color'

// UNSUPPORTED : C value 'fg_color' : for field setter : no Go type for 'Gdk.Color'

// FieldRise returns the C field 'rise'.
func (recv *TextAppearance) FieldRise() int32 {
	argValue := gi.FieldGet(textAppearanceStruct, recv.native, "rise")
	value := argValue.Int32()
	return value
}

// SetFieldRise sets the value of the C field 'rise'.
func (recv *TextAppearance) SetFieldRise(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(textAppearanceStruct, recv.native, "rise", argValue)
}

// FieldUnderline returns the C field 'underline'.
func (recv *TextAppearance) FieldUnderline() uint32 {
	argValue := gi.FieldGet(textAppearanceStruct, recv.native, "underline")
	value := argValue.Uint32()
	return value
}

// SetFieldUnderline sets the value of the C field 'underline'.
func (recv *TextAppearance) SetFieldUnderline(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(textAppearanceStruct, recv.native, "underline", argValue)
}

// FieldStrikethrough returns the C field 'strikethrough'.
func (recv *TextAppearance) FieldStrikethrough() uint32 {
	argValue := gi.FieldGet(textAppearanceStruct, recv.native, "strikethrough")
	value := argValue.Uint32()
	return value
}

// SetFieldStrikethrough sets the value of the C field 'strikethrough'.
func (recv *TextAppearance) SetFieldStrikethrough(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(textAppearanceStruct, recv.native, "strikethrough", argValue)
}

// FieldDrawBg returns the C field 'draw_bg'.
func (recv *TextAppearance) FieldDrawBg() uint32 {
	argValue := gi.FieldGet(textAppearanceStruct, recv.native, "draw_bg")
	value := argValue.Uint32()
	return value
}

// SetFieldDrawBg sets the value of the C field 'draw_bg'.
func (recv *TextAppearance) SetFieldDrawBg(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(textAppearanceStruct, recv.native, "draw_bg", argValue)
}

// FieldInsideSelection returns the C field 'inside_selection'.
func (recv *TextAppearance) FieldInsideSelection() uint32 {
	argValue := gi.FieldGet(textAppearanceStruct, recv.native, "inside_selection")
	value := argValue.Uint32()
	return value
}

// SetFieldInsideSelection sets the value of the C field 'inside_selection'.
func (recv *TextAppearance) SetFieldInsideSelection(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(textAppearanceStruct, recv.native, "inside_selection", argValue)
}

// FieldIsText returns the C field 'is_text'.
func (recv *TextAppearance) FieldIsText() uint32 {
	argValue := gi.FieldGet(textAppearanceStruct, recv.native, "is_text")
	value := argValue.Uint32()
	return value
}

// SetFieldIsText sets the value of the C field 'is_text'.
func (recv *TextAppearance) SetFieldIsText(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(textAppearanceStruct, recv.native, "is_text", argValue)
}

// TextAppearanceStruct creates an uninitialised TextAppearance.
func TextAppearanceStruct() *TextAppearance {
	err := textAppearanceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextAppearance{native: textAppearanceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextAppearance)
	return structGo
}
func finalizeTextAppearance(obj *TextAppearance) {
	textAppearanceStruct.Free(obj.native)
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

// FieldAppearance returns the C field 'appearance'.
func (recv *TextAttributes) FieldAppearance() *TextAppearance {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "appearance")
	value := &TextAppearance{native: argValue.Pointer()}
	return value
}

// SetFieldAppearance sets the value of the C field 'appearance'.
func (recv *TextAttributes) SetFieldAppearance(value *TextAppearance) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(textAttributesStruct, recv.native, "appearance", argValue)
}

// UNSUPPORTED : C value 'justification' : for field getter : no Go type for 'Justification'

// UNSUPPORTED : C value 'justification' : for field setter : no Go type for 'Justification'

// UNSUPPORTED : C value 'direction' : for field getter : no Go type for 'TextDirection'

// UNSUPPORTED : C value 'direction' : for field setter : no Go type for 'TextDirection'

// UNSUPPORTED : C value 'font' : for field getter : no Go type for 'Pango.FontDescription'

// UNSUPPORTED : C value 'font' : for field setter : no Go type for 'Pango.FontDescription'

// FieldFontScale returns the C field 'font_scale'.
func (recv *TextAttributes) FieldFontScale() float64 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "font_scale")
	value := argValue.Float64()
	return value
}

// SetFieldFontScale sets the value of the C field 'font_scale'.
func (recv *TextAttributes) SetFieldFontScale(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.FieldSet(textAttributesStruct, recv.native, "font_scale", argValue)
}

// FieldLeftMargin returns the C field 'left_margin'.
func (recv *TextAttributes) FieldLeftMargin() int32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "left_margin")
	value := argValue.Int32()
	return value
}

// SetFieldLeftMargin sets the value of the C field 'left_margin'.
func (recv *TextAttributes) SetFieldLeftMargin(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "left_margin", argValue)
}

// FieldRightMargin returns the C field 'right_margin'.
func (recv *TextAttributes) FieldRightMargin() int32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "right_margin")
	value := argValue.Int32()
	return value
}

// SetFieldRightMargin sets the value of the C field 'right_margin'.
func (recv *TextAttributes) SetFieldRightMargin(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "right_margin", argValue)
}

// FieldIndent returns the C field 'indent'.
func (recv *TextAttributes) FieldIndent() int32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "indent")
	value := argValue.Int32()
	return value
}

// SetFieldIndent sets the value of the C field 'indent'.
func (recv *TextAttributes) SetFieldIndent(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "indent", argValue)
}

// FieldPixelsAboveLines returns the C field 'pixels_above_lines'.
func (recv *TextAttributes) FieldPixelsAboveLines() int32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "pixels_above_lines")
	value := argValue.Int32()
	return value
}

// SetFieldPixelsAboveLines sets the value of the C field 'pixels_above_lines'.
func (recv *TextAttributes) SetFieldPixelsAboveLines(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "pixels_above_lines", argValue)
}

// FieldPixelsBelowLines returns the C field 'pixels_below_lines'.
func (recv *TextAttributes) FieldPixelsBelowLines() int32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "pixels_below_lines")
	value := argValue.Int32()
	return value
}

// SetFieldPixelsBelowLines sets the value of the C field 'pixels_below_lines'.
func (recv *TextAttributes) SetFieldPixelsBelowLines(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "pixels_below_lines", argValue)
}

// FieldPixelsInsideWrap returns the C field 'pixels_inside_wrap'.
func (recv *TextAttributes) FieldPixelsInsideWrap() int32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "pixels_inside_wrap")
	value := argValue.Int32()
	return value
}

// SetFieldPixelsInsideWrap sets the value of the C field 'pixels_inside_wrap'.
func (recv *TextAttributes) SetFieldPixelsInsideWrap(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "pixels_inside_wrap", argValue)
}

// UNSUPPORTED : C value 'tabs' : for field getter : no Go type for 'Pango.TabArray'

// UNSUPPORTED : C value 'tabs' : for field setter : no Go type for 'Pango.TabArray'

// UNSUPPORTED : C value 'wrap_mode' : for field getter : no Go type for 'WrapMode'

// UNSUPPORTED : C value 'wrap_mode' : for field setter : no Go type for 'WrapMode'

// UNSUPPORTED : C value 'language' : for field getter : no Go type for 'Pango.Language'

// UNSUPPORTED : C value 'language' : for field setter : no Go type for 'Pango.Language'

// FieldInvisible returns the C field 'invisible'.
func (recv *TextAttributes) FieldInvisible() uint32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "invisible")
	value := argValue.Uint32()
	return value
}

// SetFieldInvisible sets the value of the C field 'invisible'.
func (recv *TextAttributes) SetFieldInvisible(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "invisible", argValue)
}

// FieldBgFullHeight returns the C field 'bg_full_height'.
func (recv *TextAttributes) FieldBgFullHeight() uint32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "bg_full_height")
	value := argValue.Uint32()
	return value
}

// SetFieldBgFullHeight sets the value of the C field 'bg_full_height'.
func (recv *TextAttributes) SetFieldBgFullHeight(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "bg_full_height", argValue)
}

// FieldEditable returns the C field 'editable'.
func (recv *TextAttributes) FieldEditable() uint32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "editable")
	value := argValue.Uint32()
	return value
}

// SetFieldEditable sets the value of the C field 'editable'.
func (recv *TextAttributes) SetFieldEditable(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "editable", argValue)
}

// FieldNoFallback returns the C field 'no_fallback'.
func (recv *TextAttributes) FieldNoFallback() uint32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "no_fallback")
	value := argValue.Uint32()
	return value
}

// SetFieldNoFallback sets the value of the C field 'no_fallback'.
func (recv *TextAttributes) SetFieldNoFallback(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "no_fallback", argValue)
}

// FieldLetterSpacing returns the C field 'letter_spacing'.
func (recv *TextAttributes) FieldLetterSpacing() int32 {
	argValue := gi.FieldGet(textAttributesStruct, recv.native, "letter_spacing")
	value := argValue.Int32()
	return value
}

// SetFieldLetterSpacing sets the value of the C field 'letter_spacing'.
func (recv *TextAttributes) SetFieldLetterSpacing(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(textAttributesStruct, recv.native, "letter_spacing", argValue)
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

// TextBTreeStruct creates an uninitialised TextBTree.
func TextBTreeStruct() *TextBTree {
	err := textBTreeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextBTree{native: textBTreeStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextBTree)
	return structGo
}
func finalizeTextBTree(obj *TextBTree) {
	textBTreeStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'insert_text' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_text' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_pixbuf' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_pixbuf' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_child_anchor' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_child_anchor' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_range' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_range' : for field setter : missing Type

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value 'modified_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'modified_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'mark_set' : for field getter : missing Type

// UNSUPPORTED : C value 'mark_set' : for field setter : missing Type

// UNSUPPORTED : C value 'mark_deleted' : for field getter : missing Type

// UNSUPPORTED : C value 'mark_deleted' : for field setter : missing Type

// UNSUPPORTED : C value 'apply_tag' : for field getter : missing Type

// UNSUPPORTED : C value 'apply_tag' : for field setter : missing Type

// UNSUPPORTED : C value 'remove_tag' : for field getter : missing Type

// UNSUPPORTED : C value 'remove_tag' : for field setter : missing Type

// UNSUPPORTED : C value 'begin_user_action' : for field getter : missing Type

// UNSUPPORTED : C value 'begin_user_action' : for field setter : missing Type

// UNSUPPORTED : C value 'end_user_action' : for field getter : missing Type

// UNSUPPORTED : C value 'end_user_action' : for field setter : missing Type

// UNSUPPORTED : C value 'paste_done' : for field getter : missing Type

// UNSUPPORTED : C value 'paste_done' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TextBufferClassStruct creates an uninitialised TextBufferClass.
func TextBufferClassStruct() *TextBufferClass {
	err := textBufferClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextBufferClass{native: textBufferClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextBufferClass)
	return structGo
}
func finalizeTextBufferClass(obj *TextBufferClass) {
	textBufferClassStruct.Free(obj.native)
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

// TextBufferPrivateStruct creates an uninitialised TextBufferPrivate.
func TextBufferPrivateStruct() *TextBufferPrivate {
	err := textBufferPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextBufferPrivate{native: textBufferPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextBufferPrivate)
	return structGo
}
func finalizeTextBufferPrivate(obj *TextBufferPrivate) {
	textBufferPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *TextCellAccessibleClass) FieldParentClass() *RendererCellAccessibleClass {
	argValue := gi.FieldGet(textCellAccessibleClassStruct, recv.native, "parent_class")
	value := &RendererCellAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TextCellAccessibleClass) SetFieldParentClass(value *RendererCellAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(textCellAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// TextCellAccessibleClassStruct creates an uninitialised TextCellAccessibleClass.
func TextCellAccessibleClassStruct() *TextCellAccessibleClass {
	err := textCellAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextCellAccessibleClass{native: textCellAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextCellAccessibleClass)
	return structGo
}
func finalizeTextCellAccessibleClass(obj *TextCellAccessibleClass) {
	textCellAccessibleClassStruct.Free(obj.native)
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

// TextCellAccessiblePrivateStruct creates an uninitialised TextCellAccessiblePrivate.
func TextCellAccessiblePrivateStruct() *TextCellAccessiblePrivate {
	err := textCellAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextCellAccessiblePrivate{native: textCellAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextCellAccessiblePrivate)
	return structGo
}
func finalizeTextCellAccessiblePrivate(obj *TextCellAccessiblePrivate) {
	textCellAccessiblePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TextChildAnchorClassStruct creates an uninitialised TextChildAnchorClass.
func TextChildAnchorClassStruct() *TextChildAnchorClass {
	err := textChildAnchorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextChildAnchorClass{native: textChildAnchorClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextChildAnchorClass)
	return structGo
}
func finalizeTextChildAnchorClass(obj *TextChildAnchorClass) {
	textChildAnchorClassStruct.Free(obj.native)
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

var textIterBackwardToTagToggleFunction *gi.Function
var textIterBackwardToTagToggleFunction_Once sync.Once

func textIterBackwardToTagToggleFunction_Set() error {
	var err error
	textIterBackwardToTagToggleFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBackwardToTagToggleFunction, err = textIterStruct.InvokerNew("backward_to_tag_toggle")
	})
	return err
}

// BackwardToTagToggle is a representation of the C type gtk_text_iter_backward_to_tag_toggle.
func (recv *TextIter) BackwardToTagToggle(tag *TextTag) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(tag.native)

	var ret gi.Argument

	err := textIterBackwardToTagToggleFunction_Set()
	if err == nil {
		ret = textIterBackwardToTagToggleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var textIterBeginsTagFunction *gi.Function
var textIterBeginsTagFunction_Once sync.Once

func textIterBeginsTagFunction_Set() error {
	var err error
	textIterBeginsTagFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterBeginsTagFunction, err = textIterStruct.InvokerNew("begins_tag")
	})
	return err
}

// BeginsTag is a representation of the C type gtk_text_iter_begins_tag.
func (recv *TextIter) BeginsTag(tag *TextTag) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(tag.native)

	var ret gi.Argument

	err := textIterBeginsTagFunction_Set()
	if err == nil {
		ret = textIterBeginsTagFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var textIterEndsTagFunction *gi.Function
var textIterEndsTagFunction_Once sync.Once

func textIterEndsTagFunction_Set() error {
	var err error
	textIterEndsTagFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterEndsTagFunction, err = textIterStruct.InvokerNew("ends_tag")
	})
	return err
}

// EndsTag is a representation of the C type gtk_text_iter_ends_tag.
func (recv *TextIter) EndsTag(tag *TextTag) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(tag.native)

	var ret gi.Argument

	err := textIterEndsTagFunction_Set()
	if err == nil {
		ret = textIterEndsTagFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var textIterForwardToTagToggleFunction *gi.Function
var textIterForwardToTagToggleFunction_Once sync.Once

func textIterForwardToTagToggleFunction_Set() error {
	var err error
	textIterForwardToTagToggleFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterForwardToTagToggleFunction, err = textIterStruct.InvokerNew("forward_to_tag_toggle")
	})
	return err
}

// ForwardToTagToggle is a representation of the C type gtk_text_iter_forward_to_tag_toggle.
func (recv *TextIter) ForwardToTagToggle(tag *TextTag) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(tag.native)

	var ret gi.Argument

	err := textIterForwardToTagToggleFunction_Set()
	if err == nil {
		ret = textIterForwardToTagToggleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var textIterGetBufferFunction *gi.Function
var textIterGetBufferFunction_Once sync.Once

func textIterGetBufferFunction_Set() error {
	var err error
	textIterGetBufferFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetBufferFunction, err = textIterStruct.InvokerNew("get_buffer")
	})
	return err
}

// GetBuffer is a representation of the C type gtk_text_iter_get_buffer.
func (recv *TextIter) GetBuffer() *TextBuffer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetBufferFunction_Set()
	if err == nil {
		ret = textIterGetBufferFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TextBuffer{native: ret.Pointer()}

	return retGo
}

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

var textIterGetChildAnchorFunction *gi.Function
var textIterGetChildAnchorFunction_Once sync.Once

func textIterGetChildAnchorFunction_Set() error {
	var err error
	textIterGetChildAnchorFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterGetChildAnchorFunction, err = textIterStruct.InvokerNew("get_child_anchor")
	})
	return err
}

// GetChildAnchor is a representation of the C type gtk_text_iter_get_child_anchor.
func (recv *TextIter) GetChildAnchor() *TextChildAnchor {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := textIterGetChildAnchorFunction_Set()
	if err == nil {
		ret = textIterGetChildAnchorFunction.Invoke(inArgs[:], nil)
	}

	retGo := &TextChildAnchor{native: ret.Pointer()}

	return retGo
}

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

var textIterHasTagFunction *gi.Function
var textIterHasTagFunction_Once sync.Once

func textIterHasTagFunction_Set() error {
	var err error
	textIterHasTagFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterHasTagFunction, err = textIterStruct.InvokerNew("has_tag")
	})
	return err
}

// HasTag is a representation of the C type gtk_text_iter_has_tag.
func (recv *TextIter) HasTag(tag *TextTag) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(tag.native)

	var ret gi.Argument

	err := textIterHasTagFunction_Set()
	if err == nil {
		ret = textIterHasTagFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var textIterStartsTagFunction *gi.Function
var textIterStartsTagFunction_Once sync.Once

func textIterStartsTagFunction_Set() error {
	var err error
	textIterStartsTagFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterStartsTagFunction, err = textIterStruct.InvokerNew("starts_tag")
	})
	return err
}

// StartsTag is a representation of the C type gtk_text_iter_starts_tag.
func (recv *TextIter) StartsTag(tag *TextTag) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(tag.native)

	var ret gi.Argument

	err := textIterStartsTagFunction_Set()
	if err == nil {
		ret = textIterStartsTagFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

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

var textIterTogglesTagFunction *gi.Function
var textIterTogglesTagFunction_Once sync.Once

func textIterTogglesTagFunction_Set() error {
	var err error
	textIterTogglesTagFunction_Once.Do(func() {
		err = textIterStruct_Set()
		if err != nil {
			return
		}
		textIterTogglesTagFunction, err = textIterStruct.InvokerNew("toggles_tag")
	})
	return err
}

// TogglesTag is a representation of the C type gtk_text_iter_toggles_tag.
func (recv *TextIter) TogglesTag(tag *TextTag) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(tag.native)

	var ret gi.Argument

	err := textIterTogglesTagFunction_Set()
	if err == nil {
		ret = textIterTogglesTagFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// TextIterStruct creates an uninitialised TextIter.
func TextIterStruct() *TextIter {
	err := textIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextIter{native: textIterStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextIter)
	return structGo
}
func finalizeTextIter(obj *TextIter) {
	textIterStruct.Free(obj.native)
}

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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TextMarkClassStruct creates an uninitialised TextMarkClass.
func TextMarkClassStruct() *TextMarkClass {
	err := textMarkClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextMarkClass{native: textMarkClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextMarkClass)
	return structGo
}
func finalizeTextMarkClass(obj *TextMarkClass) {
	textMarkClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'event' : for field getter : missing Type

// UNSUPPORTED : C value 'event' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TextTagClassStruct creates an uninitialised TextTagClass.
func TextTagClassStruct() *TextTagClass {
	err := textTagClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextTagClass{native: textTagClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextTagClass)
	return structGo
}
func finalizeTextTagClass(obj *TextTagClass) {
	textTagClassStruct.Free(obj.native)
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

// TextTagPrivateStruct creates an uninitialised TextTagPrivate.
func TextTagPrivateStruct() *TextTagPrivate {
	err := textTagPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextTagPrivate{native: textTagPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextTagPrivate)
	return structGo
}
func finalizeTextTagPrivate(obj *TextTagPrivate) {
	textTagPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'tag_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'tag_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'tag_added' : for field getter : missing Type

// UNSUPPORTED : C value 'tag_added' : for field setter : missing Type

// UNSUPPORTED : C value 'tag_removed' : for field getter : missing Type

// UNSUPPORTED : C value 'tag_removed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TextTagTableClassStruct creates an uninitialised TextTagTableClass.
func TextTagTableClassStruct() *TextTagTableClass {
	err := textTagTableClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextTagTableClass{native: textTagTableClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextTagTableClass)
	return structGo
}
func finalizeTextTagTableClass(obj *TextTagTableClass) {
	textTagTableClassStruct.Free(obj.native)
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

// TextTagTablePrivateStruct creates an uninitialised TextTagTablePrivate.
func TextTagTablePrivateStruct() *TextTagTablePrivate {
	err := textTagTablePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextTagTablePrivate{native: textTagTablePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextTagTablePrivate)
	return structGo
}
func finalizeTextTagTablePrivate(obj *TextTagTablePrivate) {
	textTagTablePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *TextViewAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(textViewAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TextViewAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(textViewAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// TextViewAccessibleClassStruct creates an uninitialised TextViewAccessibleClass.
func TextViewAccessibleClassStruct() *TextViewAccessibleClass {
	err := textViewAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextViewAccessibleClass{native: textViewAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextViewAccessibleClass)
	return structGo
}
func finalizeTextViewAccessibleClass(obj *TextViewAccessibleClass) {
	textViewAccessibleClassStruct.Free(obj.native)
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

// TextViewAccessiblePrivateStruct creates an uninitialised TextViewAccessiblePrivate.
func TextViewAccessiblePrivateStruct() *TextViewAccessiblePrivate {
	err := textViewAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextViewAccessiblePrivate{native: textViewAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextViewAccessiblePrivate)
	return structGo
}
func finalizeTextViewAccessiblePrivate(obj *TextViewAccessiblePrivate) {
	textViewAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *TextViewClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(textViewClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TextViewClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(textViewClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'populate_popup' : for field getter : missing Type

// UNSUPPORTED : C value 'populate_popup' : for field setter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'set_anchor' : for field getter : missing Type

// UNSUPPORTED : C value 'set_anchor' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_at_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_at_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_from_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_from_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'backspace' : for field getter : missing Type

// UNSUPPORTED : C value 'backspace' : for field setter : missing Type

// UNSUPPORTED : C value 'cut_clipboard' : for field getter : missing Type

// UNSUPPORTED : C value 'cut_clipboard' : for field setter : missing Type

// UNSUPPORTED : C value 'copy_clipboard' : for field getter : missing Type

// UNSUPPORTED : C value 'copy_clipboard' : for field setter : missing Type

// UNSUPPORTED : C value 'paste_clipboard' : for field getter : missing Type

// UNSUPPORTED : C value 'paste_clipboard' : for field setter : missing Type

// UNSUPPORTED : C value 'toggle_overwrite' : for field getter : missing Type

// UNSUPPORTED : C value 'toggle_overwrite' : for field setter : missing Type

// UNSUPPORTED : C value 'create_buffer' : for field getter : missing Type

// UNSUPPORTED : C value 'create_buffer' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_layer' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_layer' : for field setter : missing Type

// UNSUPPORTED : C value 'extend_selection' : for field getter : missing Type

// UNSUPPORTED : C value 'extend_selection' : for field setter : missing Type

// UNSUPPORTED : C value 'insert_emoji' : for field getter : missing Type

// UNSUPPORTED : C value 'insert_emoji' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TextViewClassStruct creates an uninitialised TextViewClass.
func TextViewClassStruct() *TextViewClass {
	err := textViewClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextViewClass{native: textViewClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextViewClass)
	return structGo
}
func finalizeTextViewClass(obj *TextViewClass) {
	textViewClassStruct.Free(obj.native)
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

// TextViewPrivateStruct creates an uninitialised TextViewPrivate.
func TextViewPrivateStruct() *TextViewPrivate {
	err := textViewPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TextViewPrivate{native: textViewPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTextViewPrivate)
	return structGo
}
func finalizeTextViewPrivate(obj *TextViewPrivate) {
	textViewPrivateStruct.Free(obj.native)
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

// ThemeEngineStruct creates an uninitialised ThemeEngine.
func ThemeEngineStruct() *ThemeEngine {
	err := themeEngineStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ThemeEngine{native: themeEngineStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeThemeEngine)
	return structGo
}
func finalizeThemeEngine(obj *ThemeEngine) {
	themeEngineStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'render_line' : for field getter : missing Type

// UNSUPPORTED : C value 'render_line' : for field setter : missing Type

// UNSUPPORTED : C value 'render_background' : for field getter : missing Type

// UNSUPPORTED : C value 'render_background' : for field setter : missing Type

// UNSUPPORTED : C value 'render_frame' : for field getter : missing Type

// UNSUPPORTED : C value 'render_frame' : for field setter : missing Type

// UNSUPPORTED : C value 'render_frame_gap' : for field getter : missing Type

// UNSUPPORTED : C value 'render_frame_gap' : for field setter : missing Type

// UNSUPPORTED : C value 'render_extension' : for field getter : missing Type

// UNSUPPORTED : C value 'render_extension' : for field setter : missing Type

// UNSUPPORTED : C value 'render_check' : for field getter : missing Type

// UNSUPPORTED : C value 'render_check' : for field setter : missing Type

// UNSUPPORTED : C value 'render_option' : for field getter : missing Type

// UNSUPPORTED : C value 'render_option' : for field setter : missing Type

// UNSUPPORTED : C value 'render_arrow' : for field getter : missing Type

// UNSUPPORTED : C value 'render_arrow' : for field setter : missing Type

// UNSUPPORTED : C value 'render_expander' : for field getter : missing Type

// UNSUPPORTED : C value 'render_expander' : for field setter : missing Type

// UNSUPPORTED : C value 'render_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'render_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'render_layout' : for field getter : missing Type

// UNSUPPORTED : C value 'render_layout' : for field setter : missing Type

// UNSUPPORTED : C value 'render_slider' : for field getter : missing Type

// UNSUPPORTED : C value 'render_slider' : for field setter : missing Type

// UNSUPPORTED : C value 'render_handle' : for field getter : missing Type

// UNSUPPORTED : C value 'render_handle' : for field setter : missing Type

// UNSUPPORTED : C value 'render_activity' : for field getter : missing Type

// UNSUPPORTED : C value 'render_activity' : for field setter : missing Type

// UNSUPPORTED : C value 'render_icon_pixbuf' : for field getter : missing Type

// UNSUPPORTED : C value 'render_icon_pixbuf' : for field setter : missing Type

// UNSUPPORTED : C value 'render_icon' : for field getter : missing Type

// UNSUPPORTED : C value 'render_icon' : for field setter : missing Type

// UNSUPPORTED : C value 'render_icon_surface' : for field getter : missing Type

// UNSUPPORTED : C value 'render_icon_surface' : for field setter : missing Type

// ThemingEngineClassStruct creates an uninitialised ThemingEngineClass.
func ThemingEngineClassStruct() *ThemingEngineClass {
	err := themingEngineClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ThemingEngineClass{native: themingEngineClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeThemingEngineClass)
	return structGo
}
func finalizeThemingEngineClass(obj *ThemingEngineClass) {
	themingEngineClassStruct.Free(obj.native)
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

// ThemingEnginePrivateStruct creates an uninitialised ThemingEnginePrivate.
func ThemingEnginePrivateStruct() *ThemingEnginePrivate {
	err := themingEnginePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ThemingEnginePrivate{native: themingEnginePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeThemingEnginePrivate)
	return structGo
}
func finalizeThemingEnginePrivate(obj *ThemingEnginePrivate) {
	themingEnginePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ToggleActionClass) FieldParentClass() *ActionClass {
	argValue := gi.FieldGet(toggleActionClassStruct, recv.native, "parent_class")
	value := &ActionClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ToggleActionClass) SetFieldParentClass(value *ActionClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(toggleActionClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'toggled' : for field getter : missing Type

// UNSUPPORTED : C value 'toggled' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ToggleActionClassStruct creates an uninitialised ToggleActionClass.
func ToggleActionClassStruct() *ToggleActionClass {
	err := toggleActionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToggleActionClass{native: toggleActionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToggleActionClass)
	return structGo
}
func finalizeToggleActionClass(obj *ToggleActionClass) {
	toggleActionClassStruct.Free(obj.native)
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

// FieldName returns the C field 'name'.
func (recv *ToggleActionEntry) FieldName() string {
	argValue := gi.FieldGet(toggleActionEntryStruct, recv.native, "name")
	value := argValue.String(false)
	return value
}

// SetFieldName sets the value of the C field 'name'.
func (recv *ToggleActionEntry) SetFieldName(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(toggleActionEntryStruct, recv.native, "name", argValue)
}

// FieldStockId returns the C field 'stock_id'.
func (recv *ToggleActionEntry) FieldStockId() string {
	argValue := gi.FieldGet(toggleActionEntryStruct, recv.native, "stock_id")
	value := argValue.String(false)
	return value
}

// SetFieldStockId sets the value of the C field 'stock_id'.
func (recv *ToggleActionEntry) SetFieldStockId(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(toggleActionEntryStruct, recv.native, "stock_id", argValue)
}

// FieldLabel returns the C field 'label'.
func (recv *ToggleActionEntry) FieldLabel() string {
	argValue := gi.FieldGet(toggleActionEntryStruct, recv.native, "label")
	value := argValue.String(false)
	return value
}

// SetFieldLabel sets the value of the C field 'label'.
func (recv *ToggleActionEntry) SetFieldLabel(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(toggleActionEntryStruct, recv.native, "label", argValue)
}

// FieldAccelerator returns the C field 'accelerator'.
func (recv *ToggleActionEntry) FieldAccelerator() string {
	argValue := gi.FieldGet(toggleActionEntryStruct, recv.native, "accelerator")
	value := argValue.String(false)
	return value
}

// SetFieldAccelerator sets the value of the C field 'accelerator'.
func (recv *ToggleActionEntry) SetFieldAccelerator(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(toggleActionEntryStruct, recv.native, "accelerator", argValue)
}

// FieldTooltip returns the C field 'tooltip'.
func (recv *ToggleActionEntry) FieldTooltip() string {
	argValue := gi.FieldGet(toggleActionEntryStruct, recv.native, "tooltip")
	value := argValue.String(false)
	return value
}

// SetFieldTooltip sets the value of the C field 'tooltip'.
func (recv *ToggleActionEntry) SetFieldTooltip(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.FieldSet(toggleActionEntryStruct, recv.native, "tooltip", argValue)
}

// UNSUPPORTED : C value 'callback' : for field getter : no Go type for 'GObject.Callback'

// UNSUPPORTED : C value 'callback' : for field setter : no Go type for 'GObject.Callback'

// FieldIsActive returns the C field 'is_active'.
func (recv *ToggleActionEntry) FieldIsActive() bool {
	argValue := gi.FieldGet(toggleActionEntryStruct, recv.native, "is_active")
	value := argValue.Boolean()
	return value
}

// SetFieldIsActive sets the value of the C field 'is_active'.
func (recv *ToggleActionEntry) SetFieldIsActive(value bool) {
	var argValue gi.Argument
	argValue.SetBoolean(value)
	gi.FieldSet(toggleActionEntryStruct, recv.native, "is_active", argValue)
}

// ToggleActionEntryStruct creates an uninitialised ToggleActionEntry.
func ToggleActionEntryStruct() *ToggleActionEntry {
	err := toggleActionEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToggleActionEntry{native: toggleActionEntryStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToggleActionEntry)
	return structGo
}
func finalizeToggleActionEntry(obj *ToggleActionEntry) {
	toggleActionEntryStruct.Free(obj.native)
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

// ToggleActionPrivateStruct creates an uninitialised ToggleActionPrivate.
func ToggleActionPrivateStruct() *ToggleActionPrivate {
	err := toggleActionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToggleActionPrivate{native: toggleActionPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToggleActionPrivate)
	return structGo
}
func finalizeToggleActionPrivate(obj *ToggleActionPrivate) {
	toggleActionPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ToggleButtonAccessibleClass) FieldParentClass() *ButtonAccessibleClass {
	argValue := gi.FieldGet(toggleButtonAccessibleClassStruct, recv.native, "parent_class")
	value := &ButtonAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ToggleButtonAccessibleClass) SetFieldParentClass(value *ButtonAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(toggleButtonAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// ToggleButtonAccessibleClassStruct creates an uninitialised ToggleButtonAccessibleClass.
func ToggleButtonAccessibleClassStruct() *ToggleButtonAccessibleClass {
	err := toggleButtonAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToggleButtonAccessibleClass{native: toggleButtonAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToggleButtonAccessibleClass)
	return structGo
}
func finalizeToggleButtonAccessibleClass(obj *ToggleButtonAccessibleClass) {
	toggleButtonAccessibleClassStruct.Free(obj.native)
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

// ToggleButtonAccessiblePrivateStruct creates an uninitialised ToggleButtonAccessiblePrivate.
func ToggleButtonAccessiblePrivateStruct() *ToggleButtonAccessiblePrivate {
	err := toggleButtonAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToggleButtonAccessiblePrivate{native: toggleButtonAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToggleButtonAccessiblePrivate)
	return structGo
}
func finalizeToggleButtonAccessiblePrivate(obj *ToggleButtonAccessiblePrivate) {
	toggleButtonAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ToggleButtonClass) FieldParentClass() *ButtonClass {
	argValue := gi.FieldGet(toggleButtonClassStruct, recv.native, "parent_class")
	value := &ButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ToggleButtonClass) SetFieldParentClass(value *ButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(toggleButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'toggled' : for field getter : missing Type

// UNSUPPORTED : C value 'toggled' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ToggleButtonClassStruct creates an uninitialised ToggleButtonClass.
func ToggleButtonClassStruct() *ToggleButtonClass {
	err := toggleButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToggleButtonClass{native: toggleButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToggleButtonClass)
	return structGo
}
func finalizeToggleButtonClass(obj *ToggleButtonClass) {
	toggleButtonClassStruct.Free(obj.native)
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

// ToggleButtonPrivateStruct creates an uninitialised ToggleButtonPrivate.
func ToggleButtonPrivateStruct() *ToggleButtonPrivate {
	err := toggleButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToggleButtonPrivate{native: toggleButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToggleButtonPrivate)
	return structGo
}
func finalizeToggleButtonPrivate(obj *ToggleButtonPrivate) {
	toggleButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ToggleToolButtonClass) FieldParentClass() *ToolButtonClass {
	argValue := gi.FieldGet(toggleToolButtonClassStruct, recv.native, "parent_class")
	value := &ToolButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ToggleToolButtonClass) SetFieldParentClass(value *ToolButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(toggleToolButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'toggled' : for field getter : missing Type

// UNSUPPORTED : C value 'toggled' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ToggleToolButtonClassStruct creates an uninitialised ToggleToolButtonClass.
func ToggleToolButtonClassStruct() *ToggleToolButtonClass {
	err := toggleToolButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToggleToolButtonClass{native: toggleToolButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToggleToolButtonClass)
	return structGo
}
func finalizeToggleToolButtonClass(obj *ToggleToolButtonClass) {
	toggleToolButtonClassStruct.Free(obj.native)
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

// ToggleToolButtonPrivateStruct creates an uninitialised ToggleToolButtonPrivate.
func ToggleToolButtonPrivateStruct() *ToggleToolButtonPrivate {
	err := toggleToolButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToggleToolButtonPrivate{native: toggleToolButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToggleToolButtonPrivate)
	return structGo
}
func finalizeToggleToolButtonPrivate(obj *ToggleToolButtonPrivate) {
	toggleToolButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ToolButtonClass) FieldParentClass() *ToolItemClass {
	argValue := gi.FieldGet(toolButtonClassStruct, recv.native, "parent_class")
	value := &ToolItemClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ToolButtonClass) SetFieldParentClass(value *ToolItemClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(toolButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'button_type' : for field getter : no Go type for 'GType'

// UNSUPPORTED : C value 'button_type' : for field setter : no Go type for 'GType'

// UNSUPPORTED : C value 'clicked' : for field getter : missing Type

// UNSUPPORTED : C value 'clicked' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ToolButtonClassStruct creates an uninitialised ToolButtonClass.
func ToolButtonClassStruct() *ToolButtonClass {
	err := toolButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolButtonClass{native: toolButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolButtonClass)
	return structGo
}
func finalizeToolButtonClass(obj *ToolButtonClass) {
	toolButtonClassStruct.Free(obj.native)
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

// ToolButtonPrivateStruct creates an uninitialised ToolButtonPrivate.
func ToolButtonPrivateStruct() *ToolButtonPrivate {
	err := toolButtonPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolButtonPrivate{native: toolButtonPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolButtonPrivate)
	return structGo
}
func finalizeToolButtonPrivate(obj *ToolButtonPrivate) {
	toolButtonPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ToolItemClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(toolItemClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ToolItemClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(toolItemClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'create_menu_proxy' : for field getter : missing Type

// UNSUPPORTED : C value 'create_menu_proxy' : for field setter : missing Type

// UNSUPPORTED : C value 'toolbar_reconfigured' : for field getter : missing Type

// UNSUPPORTED : C value 'toolbar_reconfigured' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ToolItemClassStruct creates an uninitialised ToolItemClass.
func ToolItemClassStruct() *ToolItemClass {
	err := toolItemClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolItemClass{native: toolItemClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolItemClass)
	return structGo
}
func finalizeToolItemClass(obj *ToolItemClass) {
	toolItemClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ToolItemGroupClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(toolItemGroupClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ToolItemGroupClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(toolItemGroupClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ToolItemGroupClassStruct creates an uninitialised ToolItemGroupClass.
func ToolItemGroupClassStruct() *ToolItemGroupClass {
	err := toolItemGroupClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolItemGroupClass{native: toolItemGroupClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolItemGroupClass)
	return structGo
}
func finalizeToolItemGroupClass(obj *ToolItemGroupClass) {
	toolItemGroupClassStruct.Free(obj.native)
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

// ToolItemGroupPrivateStruct creates an uninitialised ToolItemGroupPrivate.
func ToolItemGroupPrivateStruct() *ToolItemGroupPrivate {
	err := toolItemGroupPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolItemGroupPrivate{native: toolItemGroupPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolItemGroupPrivate)
	return structGo
}
func finalizeToolItemGroupPrivate(obj *ToolItemGroupPrivate) {
	toolItemGroupPrivateStruct.Free(obj.native)
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

// ToolItemPrivateStruct creates an uninitialised ToolItemPrivate.
func ToolItemPrivateStruct() *ToolItemPrivate {
	err := toolItemPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolItemPrivate{native: toolItemPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolItemPrivate)
	return structGo
}
func finalizeToolItemPrivate(obj *ToolItemPrivate) {
	toolItemPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ToolPaletteClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(toolPaletteClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ToolPaletteClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(toolPaletteClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ToolPaletteClassStruct creates an uninitialised ToolPaletteClass.
func ToolPaletteClassStruct() *ToolPaletteClass {
	err := toolPaletteClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolPaletteClass{native: toolPaletteClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolPaletteClass)
	return structGo
}
func finalizeToolPaletteClass(obj *ToolPaletteClass) {
	toolPaletteClassStruct.Free(obj.native)
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

// ToolPalettePrivateStruct creates an uninitialised ToolPalettePrivate.
func ToolPalettePrivateStruct() *ToolPalettePrivate {
	err := toolPalettePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolPalettePrivate{native: toolPalettePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolPalettePrivate)
	return structGo
}
func finalizeToolPalettePrivate(obj *ToolPalettePrivate) {
	toolPalettePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'get_icon_size' : for field getter : missing Type

// UNSUPPORTED : C value 'get_icon_size' : for field setter : missing Type

// UNSUPPORTED : C value 'get_orientation' : for field getter : missing Type

// UNSUPPORTED : C value 'get_orientation' : for field setter : missing Type

// UNSUPPORTED : C value 'get_style' : for field getter : missing Type

// UNSUPPORTED : C value 'get_style' : for field setter : missing Type

// UNSUPPORTED : C value 'get_relief_style' : for field getter : missing Type

// UNSUPPORTED : C value 'get_relief_style' : for field setter : missing Type

// UNSUPPORTED : C value 'rebuild_menu' : for field getter : missing Type

// UNSUPPORTED : C value 'rebuild_menu' : for field setter : missing Type

// UNSUPPORTED : C value 'get_text_orientation' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_orientation' : for field setter : missing Type

// UNSUPPORTED : C value 'get_text_alignment' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_alignment' : for field setter : missing Type

// UNSUPPORTED : C value 'get_ellipsize_mode' : for field getter : missing Type

// UNSUPPORTED : C value 'get_ellipsize_mode' : for field setter : missing Type

// UNSUPPORTED : C value 'get_text_size_group' : for field getter : missing Type

// UNSUPPORTED : C value 'get_text_size_group' : for field setter : missing Type

// ToolShellIfaceStruct creates an uninitialised ToolShellIface.
func ToolShellIfaceStruct() *ToolShellIface {
	err := toolShellIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolShellIface{native: toolShellIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolShellIface)
	return structGo
}
func finalizeToolShellIface(obj *ToolShellIface) {
	toolShellIfaceStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ToolbarClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(toolbarClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ToolbarClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(toolbarClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'orientation_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'orientation_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'style_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'style_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'popup_context_menu' : for field getter : missing Type

// UNSUPPORTED : C value 'popup_context_menu' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ToolbarClassStruct creates an uninitialised ToolbarClass.
func ToolbarClassStruct() *ToolbarClass {
	err := toolbarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolbarClass{native: toolbarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolbarClass)
	return structGo
}
func finalizeToolbarClass(obj *ToolbarClass) {
	toolbarClassStruct.Free(obj.native)
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

// ToolbarPrivateStruct creates an uninitialised ToolbarPrivate.
func ToolbarPrivateStruct() *ToolbarPrivate {
	err := toolbarPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToolbarPrivate{native: toolbarPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToolbarPrivate)
	return structGo
}
func finalizeToolbarPrivate(obj *ToolbarPrivate) {
	toolbarPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'Atk.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'Atk.ObjectClass'

// ToplevelAccessibleClassStruct creates an uninitialised ToplevelAccessibleClass.
func ToplevelAccessibleClassStruct() *ToplevelAccessibleClass {
	err := toplevelAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToplevelAccessibleClass{native: toplevelAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToplevelAccessibleClass)
	return structGo
}
func finalizeToplevelAccessibleClass(obj *ToplevelAccessibleClass) {
	toplevelAccessibleClassStruct.Free(obj.native)
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

// ToplevelAccessiblePrivateStruct creates an uninitialised ToplevelAccessiblePrivate.
func ToplevelAccessiblePrivateStruct() *ToplevelAccessiblePrivate {
	err := toplevelAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ToplevelAccessiblePrivate{native: toplevelAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeToplevelAccessiblePrivate)
	return structGo
}
func finalizeToplevelAccessiblePrivate(obj *ToplevelAccessiblePrivate) {
	toplevelAccessiblePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'drag_data_received' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_data_received' : for field setter : missing Type

// UNSUPPORTED : C value 'row_drop_possible' : for field getter : missing Type

// UNSUPPORTED : C value 'row_drop_possible' : for field setter : missing Type

// TreeDragDestIfaceStruct creates an uninitialised TreeDragDestIface.
func TreeDragDestIfaceStruct() *TreeDragDestIface {
	err := treeDragDestIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeDragDestIface{native: treeDragDestIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeDragDestIface)
	return structGo
}
func finalizeTreeDragDestIface(obj *TreeDragDestIface) {
	treeDragDestIfaceStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'row_draggable' : for field getter : missing Type

// UNSUPPORTED : C value 'row_draggable' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_data_get' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_data_get' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_data_delete' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_data_delete' : for field setter : missing Type

// TreeDragSourceIfaceStruct creates an uninitialised TreeDragSourceIface.
func TreeDragSourceIfaceStruct() *TreeDragSourceIface {
	err := treeDragSourceIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeDragSourceIface{native: treeDragSourceIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeDragSourceIface)
	return structGo
}
func finalizeTreeDragSourceIface(obj *TreeDragSourceIface) {
	treeDragSourceIfaceStruct.Free(obj.native)
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

// FieldStamp returns the C field 'stamp'.
func (recv *TreeIter) FieldStamp() int32 {
	argValue := gi.FieldGet(treeIterStruct, recv.native, "stamp")
	value := argValue.Int32()
	return value
}

// SetFieldStamp sets the value of the C field 'stamp'.
func (recv *TreeIter) SetFieldStamp(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.FieldSet(treeIterStruct, recv.native, "stamp", argValue)
}

// UNSUPPORTED : C value 'user_data' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'user_data' : for field setter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'user_data2' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'user_data2' : for field setter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'user_data3' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'user_data3' : for field setter : no Go type for 'gpointer'

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

// TreeIterStruct creates an uninitialised TreeIter.
func TreeIterStruct() *TreeIter {
	err := treeIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeIter{native: treeIterStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeIter)
	return structGo
}
func finalizeTreeIter(obj *TreeIter) {
	treeIterStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'visible' : for field getter : missing Type

// UNSUPPORTED : C value 'visible' : for field setter : missing Type

// UNSUPPORTED : C value 'modify' : for field getter : missing Type

// UNSUPPORTED : C value 'modify' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TreeModelFilterClassStruct creates an uninitialised TreeModelFilterClass.
func TreeModelFilterClassStruct() *TreeModelFilterClass {
	err := treeModelFilterClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeModelFilterClass{native: treeModelFilterClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeModelFilterClass)
	return structGo
}
func finalizeTreeModelFilterClass(obj *TreeModelFilterClass) {
	treeModelFilterClassStruct.Free(obj.native)
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

// TreeModelFilterPrivateStruct creates an uninitialised TreeModelFilterPrivate.
func TreeModelFilterPrivateStruct() *TreeModelFilterPrivate {
	err := treeModelFilterPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeModelFilterPrivate{native: treeModelFilterPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeModelFilterPrivate)
	return structGo
}
func finalizeTreeModelFilterPrivate(obj *TreeModelFilterPrivate) {
	treeModelFilterPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'row_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'row_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'row_inserted' : for field getter : missing Type

// UNSUPPORTED : C value 'row_inserted' : for field setter : missing Type

// UNSUPPORTED : C value 'row_has_child_toggled' : for field getter : missing Type

// UNSUPPORTED : C value 'row_has_child_toggled' : for field setter : missing Type

// UNSUPPORTED : C value 'row_deleted' : for field getter : missing Type

// UNSUPPORTED : C value 'row_deleted' : for field setter : missing Type

// UNSUPPORTED : C value 'rows_reordered' : for field getter : missing Type

// UNSUPPORTED : C value 'rows_reordered' : for field setter : missing Type

// UNSUPPORTED : C value 'get_flags' : for field getter : missing Type

// UNSUPPORTED : C value 'get_flags' : for field setter : missing Type

// UNSUPPORTED : C value 'get_n_columns' : for field getter : missing Type

// UNSUPPORTED : C value 'get_n_columns' : for field setter : missing Type

// UNSUPPORTED : C value 'get_column_type' : for field getter : missing Type

// UNSUPPORTED : C value 'get_column_type' : for field setter : missing Type

// UNSUPPORTED : C value 'get_iter' : for field getter : missing Type

// UNSUPPORTED : C value 'get_iter' : for field setter : missing Type

// UNSUPPORTED : C value 'get_path' : for field getter : missing Type

// UNSUPPORTED : C value 'get_path' : for field setter : missing Type

// UNSUPPORTED : C value 'get_value' : for field getter : missing Type

// UNSUPPORTED : C value 'get_value' : for field setter : missing Type

// UNSUPPORTED : C value 'iter_next' : for field getter : missing Type

// UNSUPPORTED : C value 'iter_next' : for field setter : missing Type

// UNSUPPORTED : C value 'iter_previous' : for field getter : missing Type

// UNSUPPORTED : C value 'iter_previous' : for field setter : missing Type

// UNSUPPORTED : C value 'iter_children' : for field getter : missing Type

// UNSUPPORTED : C value 'iter_children' : for field setter : missing Type

// UNSUPPORTED : C value 'iter_has_child' : for field getter : missing Type

// UNSUPPORTED : C value 'iter_has_child' : for field setter : missing Type

// UNSUPPORTED : C value 'iter_n_children' : for field getter : missing Type

// UNSUPPORTED : C value 'iter_n_children' : for field setter : missing Type

// UNSUPPORTED : C value 'iter_nth_child' : for field getter : missing Type

// UNSUPPORTED : C value 'iter_nth_child' : for field setter : missing Type

// UNSUPPORTED : C value 'iter_parent' : for field getter : missing Type

// UNSUPPORTED : C value 'iter_parent' : for field setter : missing Type

// UNSUPPORTED : C value 'ref_node' : for field getter : missing Type

// UNSUPPORTED : C value 'ref_node' : for field setter : missing Type

// UNSUPPORTED : C value 'unref_node' : for field getter : missing Type

// UNSUPPORTED : C value 'unref_node' : for field setter : missing Type

// TreeModelIfaceStruct creates an uninitialised TreeModelIface.
func TreeModelIfaceStruct() *TreeModelIface {
	err := treeModelIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeModelIface{native: treeModelIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeModelIface)
	return structGo
}
func finalizeTreeModelIface(obj *TreeModelIface) {
	treeModelIfaceStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TreeModelSortClassStruct creates an uninitialised TreeModelSortClass.
func TreeModelSortClassStruct() *TreeModelSortClass {
	err := treeModelSortClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeModelSortClass{native: treeModelSortClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeModelSortClass)
	return structGo
}
func finalizeTreeModelSortClass(obj *TreeModelSortClass) {
	treeModelSortClassStruct.Free(obj.native)
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

// TreeModelSortPrivateStruct creates an uninitialised TreeModelSortPrivate.
func TreeModelSortPrivateStruct() *TreeModelSortPrivate {
	err := treeModelSortPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeModelSortPrivate{native: treeModelSortPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeModelSortPrivate)
	return structGo
}
func finalizeTreeModelSortPrivate(obj *TreeModelSortPrivate) {
	treeModelSortPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'gtk_tree_path_new_from_indices' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_tree_path_new_from_indicesv' : parameter 'indices' of type 'nil' not supported

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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TreeSelectionClassStruct creates an uninitialised TreeSelectionClass.
func TreeSelectionClassStruct() *TreeSelectionClass {
	err := treeSelectionClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeSelectionClass{native: treeSelectionClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeSelectionClass)
	return structGo
}
func finalizeTreeSelectionClass(obj *TreeSelectionClass) {
	treeSelectionClassStruct.Free(obj.native)
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

// TreeSelectionPrivateStruct creates an uninitialised TreeSelectionPrivate.
func TreeSelectionPrivateStruct() *TreeSelectionPrivate {
	err := treeSelectionPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeSelectionPrivate{native: treeSelectionPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeSelectionPrivate)
	return structGo
}
func finalizeTreeSelectionPrivate(obj *TreeSelectionPrivate) {
	treeSelectionPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'sort_column_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'sort_column_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'get_sort_column_id' : for field getter : missing Type

// UNSUPPORTED : C value 'get_sort_column_id' : for field setter : missing Type

// UNSUPPORTED : C value 'set_sort_column_id' : for field getter : missing Type

// UNSUPPORTED : C value 'set_sort_column_id' : for field setter : missing Type

// UNSUPPORTED : C value 'set_sort_func' : for field getter : missing Type

// UNSUPPORTED : C value 'set_sort_func' : for field setter : missing Type

// UNSUPPORTED : C value 'set_default_sort_func' : for field getter : missing Type

// UNSUPPORTED : C value 'set_default_sort_func' : for field setter : missing Type

// UNSUPPORTED : C value 'has_default_sort_func' : for field getter : missing Type

// UNSUPPORTED : C value 'has_default_sort_func' : for field setter : missing Type

// TreeSortableIfaceStruct creates an uninitialised TreeSortableIface.
func TreeSortableIfaceStruct() *TreeSortableIface {
	err := treeSortableIfaceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeSortableIface{native: treeSortableIfaceStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeSortableIface)
	return structGo
}
func finalizeTreeSortableIface(obj *TreeSortableIface) {
	treeSortableIfaceStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TreeStoreClassStruct creates an uninitialised TreeStoreClass.
func TreeStoreClassStruct() *TreeStoreClass {
	err := treeStoreClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeStoreClass{native: treeStoreClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeStoreClass)
	return structGo
}
func finalizeTreeStoreClass(obj *TreeStoreClass) {
	treeStoreClassStruct.Free(obj.native)
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

// TreeStorePrivateStruct creates an uninitialised TreeStorePrivate.
func TreeStorePrivateStruct() *TreeStorePrivate {
	err := treeStorePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeStorePrivate{native: treeStorePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeStorePrivate)
	return structGo
}
func finalizeTreeStorePrivate(obj *TreeStorePrivate) {
	treeStorePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *TreeViewAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(treeViewAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TreeViewAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(treeViewAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// TreeViewAccessibleClassStruct creates an uninitialised TreeViewAccessibleClass.
func TreeViewAccessibleClassStruct() *TreeViewAccessibleClass {
	err := treeViewAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeViewAccessibleClass{native: treeViewAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeViewAccessibleClass)
	return structGo
}
func finalizeTreeViewAccessibleClass(obj *TreeViewAccessibleClass) {
	treeViewAccessibleClassStruct.Free(obj.native)
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

// TreeViewAccessiblePrivateStruct creates an uninitialised TreeViewAccessiblePrivate.
func TreeViewAccessiblePrivateStruct() *TreeViewAccessiblePrivate {
	err := treeViewAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeViewAccessiblePrivate{native: treeViewAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeViewAccessiblePrivate)
	return structGo
}
func finalizeTreeViewAccessiblePrivate(obj *TreeViewAccessiblePrivate) {
	treeViewAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *TreeViewClass) FieldParentClass() *ContainerClass {
	argValue := gi.FieldGet(treeViewClassStruct, recv.native, "parent_class")
	value := &ContainerClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *TreeViewClass) SetFieldParentClass(value *ContainerClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(treeViewClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'row_activated' : for field getter : missing Type

// UNSUPPORTED : C value 'row_activated' : for field setter : missing Type

// UNSUPPORTED : C value 'test_expand_row' : for field getter : missing Type

// UNSUPPORTED : C value 'test_expand_row' : for field setter : missing Type

// UNSUPPORTED : C value 'test_collapse_row' : for field getter : missing Type

// UNSUPPORTED : C value 'test_collapse_row' : for field setter : missing Type

// UNSUPPORTED : C value 'row_expanded' : for field getter : missing Type

// UNSUPPORTED : C value 'row_expanded' : for field setter : missing Type

// UNSUPPORTED : C value 'row_collapsed' : for field getter : missing Type

// UNSUPPORTED : C value 'row_collapsed' : for field setter : missing Type

// UNSUPPORTED : C value 'columns_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'columns_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'cursor_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'cursor_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field getter : missing Type

// UNSUPPORTED : C value 'move_cursor' : for field setter : missing Type

// UNSUPPORTED : C value 'select_all' : for field getter : missing Type

// UNSUPPORTED : C value 'select_all' : for field setter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field getter : missing Type

// UNSUPPORTED : C value 'unselect_all' : for field setter : missing Type

// UNSUPPORTED : C value 'select_cursor_row' : for field getter : missing Type

// UNSUPPORTED : C value 'select_cursor_row' : for field setter : missing Type

// UNSUPPORTED : C value 'toggle_cursor_row' : for field getter : missing Type

// UNSUPPORTED : C value 'toggle_cursor_row' : for field setter : missing Type

// UNSUPPORTED : C value 'expand_collapse_cursor_row' : for field getter : missing Type

// UNSUPPORTED : C value 'expand_collapse_cursor_row' : for field setter : missing Type

// UNSUPPORTED : C value 'select_cursor_parent' : for field getter : missing Type

// UNSUPPORTED : C value 'select_cursor_parent' : for field setter : missing Type

// UNSUPPORTED : C value 'start_interactive_search' : for field getter : missing Type

// UNSUPPORTED : C value 'start_interactive_search' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved5' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved8' : for field setter : missing Type

// TreeViewClassStruct creates an uninitialised TreeViewClass.
func TreeViewClassStruct() *TreeViewClass {
	err := treeViewClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeViewClass{native: treeViewClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeViewClass)
	return structGo
}
func finalizeTreeViewClass(obj *TreeViewClass) {
	treeViewClassStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'clicked' : for field getter : missing Type

// UNSUPPORTED : C value 'clicked' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// TreeViewColumnClassStruct creates an uninitialised TreeViewColumnClass.
func TreeViewColumnClassStruct() *TreeViewColumnClass {
	err := treeViewColumnClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeViewColumnClass{native: treeViewColumnClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeViewColumnClass)
	return structGo
}
func finalizeTreeViewColumnClass(obj *TreeViewColumnClass) {
	treeViewColumnClassStruct.Free(obj.native)
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

// TreeViewColumnPrivateStruct creates an uninitialised TreeViewColumnPrivate.
func TreeViewColumnPrivateStruct() *TreeViewColumnPrivate {
	err := treeViewColumnPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeViewColumnPrivate{native: treeViewColumnPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeViewColumnPrivate)
	return structGo
}
func finalizeTreeViewColumnPrivate(obj *TreeViewColumnPrivate) {
	treeViewColumnPrivateStruct.Free(obj.native)
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

// TreeViewPrivateStruct creates an uninitialised TreeViewPrivate.
func TreeViewPrivateStruct() *TreeViewPrivate {
	err := treeViewPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &TreeViewPrivate{native: treeViewPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeTreeViewPrivate)
	return structGo
}
func finalizeTreeViewPrivate(obj *TreeViewPrivate) {
	treeViewPrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'add_widget' : for field getter : missing Type

// UNSUPPORTED : C value 'add_widget' : for field setter : missing Type

// UNSUPPORTED : C value 'actions_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'actions_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'connect_proxy' : for field getter : missing Type

// UNSUPPORTED : C value 'connect_proxy' : for field setter : missing Type

// UNSUPPORTED : C value 'disconnect_proxy' : for field getter : missing Type

// UNSUPPORTED : C value 'disconnect_proxy' : for field setter : missing Type

// UNSUPPORTED : C value 'pre_activate' : for field getter : missing Type

// UNSUPPORTED : C value 'pre_activate' : for field setter : missing Type

// UNSUPPORTED : C value 'post_activate' : for field getter : missing Type

// UNSUPPORTED : C value 'post_activate' : for field setter : missing Type

// UNSUPPORTED : C value 'get_widget' : for field getter : missing Type

// UNSUPPORTED : C value 'get_widget' : for field setter : missing Type

// UNSUPPORTED : C value 'get_action' : for field getter : missing Type

// UNSUPPORTED : C value 'get_action' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// UIManagerClassStruct creates an uninitialised UIManagerClass.
func UIManagerClassStruct() *UIManagerClass {
	err := uIManagerClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UIManagerClass{native: uIManagerClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeUIManagerClass)
	return structGo
}
func finalizeUIManagerClass(obj *UIManagerClass) {
	uIManagerClassStruct.Free(obj.native)
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

// UIManagerPrivateStruct creates an uninitialised UIManagerPrivate.
func UIManagerPrivateStruct() *UIManagerPrivate {
	err := uIManagerPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &UIManagerPrivate{native: uIManagerPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeUIManagerPrivate)
	return structGo
}
func finalizeUIManagerPrivate(obj *UIManagerPrivate) {
	uIManagerPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *VBoxClass) FieldParentClass() *BoxClass {
	argValue := gi.FieldGet(vBoxClassStruct, recv.native, "parent_class")
	value := &BoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *VBoxClass) SetFieldParentClass(value *BoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(vBoxClassStruct, recv.native, "parent_class", argValue)
}

// VBoxClassStruct creates an uninitialised VBoxClass.
func VBoxClassStruct() *VBoxClass {
	err := vBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &VBoxClass{native: vBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeVBoxClass)
	return structGo
}
func finalizeVBoxClass(obj *VBoxClass) {
	vBoxClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *VButtonBoxClass) FieldParentClass() *ButtonBoxClass {
	argValue := gi.FieldGet(vButtonBoxClassStruct, recv.native, "parent_class")
	value := &ButtonBoxClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *VButtonBoxClass) SetFieldParentClass(value *ButtonBoxClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(vButtonBoxClassStruct, recv.native, "parent_class", argValue)
}

// VButtonBoxClassStruct creates an uninitialised VButtonBoxClass.
func VButtonBoxClassStruct() *VButtonBoxClass {
	err := vButtonBoxClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &VButtonBoxClass{native: vButtonBoxClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeVButtonBoxClass)
	return structGo
}
func finalizeVButtonBoxClass(obj *VButtonBoxClass) {
	vButtonBoxClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *VPanedClass) FieldParentClass() *PanedClass {
	argValue := gi.FieldGet(vPanedClassStruct, recv.native, "parent_class")
	value := &PanedClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *VPanedClass) SetFieldParentClass(value *PanedClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(vPanedClassStruct, recv.native, "parent_class", argValue)
}

// VPanedClassStruct creates an uninitialised VPanedClass.
func VPanedClassStruct() *VPanedClass {
	err := vPanedClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &VPanedClass{native: vPanedClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeVPanedClass)
	return structGo
}
func finalizeVPanedClass(obj *VPanedClass) {
	vPanedClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *VScaleClass) FieldParentClass() *ScaleClass {
	argValue := gi.FieldGet(vScaleClassStruct, recv.native, "parent_class")
	value := &ScaleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *VScaleClass) SetFieldParentClass(value *ScaleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(vScaleClassStruct, recv.native, "parent_class", argValue)
}

// VScaleClassStruct creates an uninitialised VScaleClass.
func VScaleClassStruct() *VScaleClass {
	err := vScaleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &VScaleClass{native: vScaleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeVScaleClass)
	return structGo
}
func finalizeVScaleClass(obj *VScaleClass) {
	vScaleClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *VScrollbarClass) FieldParentClass() *ScrollbarClass {
	argValue := gi.FieldGet(vScrollbarClassStruct, recv.native, "parent_class")
	value := &ScrollbarClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *VScrollbarClass) SetFieldParentClass(value *ScrollbarClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(vScrollbarClassStruct, recv.native, "parent_class", argValue)
}

// VScrollbarClassStruct creates an uninitialised VScrollbarClass.
func VScrollbarClassStruct() *VScrollbarClass {
	err := vScrollbarClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &VScrollbarClass{native: vScrollbarClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeVScrollbarClass)
	return structGo
}
func finalizeVScrollbarClass(obj *VScrollbarClass) {
	vScrollbarClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *VSeparatorClass) FieldParentClass() *SeparatorClass {
	argValue := gi.FieldGet(vSeparatorClassStruct, recv.native, "parent_class")
	value := &SeparatorClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *VSeparatorClass) SetFieldParentClass(value *SeparatorClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(vSeparatorClassStruct, recv.native, "parent_class", argValue)
}

// VSeparatorClassStruct creates an uninitialised VSeparatorClass.
func VSeparatorClassStruct() *VSeparatorClass {
	err := vSeparatorClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &VSeparatorClass{native: vSeparatorClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeVSeparatorClass)
	return structGo
}
func finalizeVSeparatorClass(obj *VSeparatorClass) {
	vSeparatorClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *ViewportClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(viewportClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *ViewportClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(viewportClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// ViewportClassStruct creates an uninitialised ViewportClass.
func ViewportClassStruct() *ViewportClass {
	err := viewportClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ViewportClass{native: viewportClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeViewportClass)
	return structGo
}
func finalizeViewportClass(obj *ViewportClass) {
	viewportClassStruct.Free(obj.native)
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

// ViewportPrivateStruct creates an uninitialised ViewportPrivate.
func ViewportPrivateStruct() *ViewportPrivate {
	err := viewportPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &ViewportPrivate{native: viewportPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeViewportPrivate)
	return structGo
}
func finalizeViewportPrivate(obj *ViewportPrivate) {
	viewportPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *VolumeButtonClass) FieldParentClass() *ScaleButtonClass {
	argValue := gi.FieldGet(volumeButtonClassStruct, recv.native, "parent_class")
	value := &ScaleButtonClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *VolumeButtonClass) SetFieldParentClass(value *ScaleButtonClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(volumeButtonClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// VolumeButtonClassStruct creates an uninitialised VolumeButtonClass.
func VolumeButtonClassStruct() *VolumeButtonClass {
	err := volumeButtonClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &VolumeButtonClass{native: volumeButtonClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeVolumeButtonClass)
	return structGo
}
func finalizeVolumeButtonClass(obj *VolumeButtonClass) {
	volumeButtonClassStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *WidgetAccessibleClass) FieldParentClass() *AccessibleClass {
	argValue := gi.FieldGet(widgetAccessibleClassStruct, recv.native, "parent_class")
	value := &AccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *WidgetAccessibleClass) SetFieldParentClass(value *AccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(widgetAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'notify_gtk' : for field getter : missing Type

// UNSUPPORTED : C value 'notify_gtk' : for field setter : missing Type

// WidgetAccessibleClassStruct creates an uninitialised WidgetAccessibleClass.
func WidgetAccessibleClassStruct() *WidgetAccessibleClass {
	err := widgetAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WidgetAccessibleClass{native: widgetAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWidgetAccessibleClass)
	return structGo
}
func finalizeWidgetAccessibleClass(obj *WidgetAccessibleClass) {
	widgetAccessibleClassStruct.Free(obj.native)
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

// WidgetAccessiblePrivateStruct creates an uninitialised WidgetAccessiblePrivate.
func WidgetAccessiblePrivateStruct() *WidgetAccessiblePrivate {
	err := widgetAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WidgetAccessiblePrivate{native: widgetAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWidgetAccessiblePrivate)
	return structGo
}
func finalizeWidgetAccessiblePrivate(obj *WidgetAccessiblePrivate) {
	widgetAccessiblePrivateStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.InitiallyUnownedClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.InitiallyUnownedClass'

// FieldActivateSignal returns the C field 'activate_signal'.
func (recv *WidgetClass) FieldActivateSignal() uint32 {
	argValue := gi.FieldGet(widgetClassStruct, recv.native, "activate_signal")
	value := argValue.Uint32()
	return value
}

// SetFieldActivateSignal sets the value of the C field 'activate_signal'.
func (recv *WidgetClass) SetFieldActivateSignal(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.FieldSet(widgetClassStruct, recv.native, "activate_signal", argValue)
}

// UNSUPPORTED : C value 'dispatch_child_properties_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'dispatch_child_properties_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'destroy' : for field getter : missing Type

// UNSUPPORTED : C value 'destroy' : for field setter : missing Type

// UNSUPPORTED : C value 'show' : for field getter : missing Type

// UNSUPPORTED : C value 'show' : for field setter : missing Type

// UNSUPPORTED : C value 'show_all' : for field getter : missing Type

// UNSUPPORTED : C value 'show_all' : for field setter : missing Type

// UNSUPPORTED : C value 'hide' : for field getter : missing Type

// UNSUPPORTED : C value 'hide' : for field setter : missing Type

// UNSUPPORTED : C value 'map' : for field getter : missing Type

// UNSUPPORTED : C value 'map' : for field setter : missing Type

// UNSUPPORTED : C value 'unmap' : for field getter : missing Type

// UNSUPPORTED : C value 'unmap' : for field setter : missing Type

// UNSUPPORTED : C value 'realize' : for field getter : missing Type

// UNSUPPORTED : C value 'realize' : for field setter : missing Type

// UNSUPPORTED : C value 'unrealize' : for field getter : missing Type

// UNSUPPORTED : C value 'unrealize' : for field setter : missing Type

// UNSUPPORTED : C value 'size_allocate' : for field getter : missing Type

// UNSUPPORTED : C value 'size_allocate' : for field setter : missing Type

// UNSUPPORTED : C value 'state_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'state_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'state_flags_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'state_flags_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'parent_set' : for field getter : missing Type

// UNSUPPORTED : C value 'parent_set' : for field setter : missing Type

// UNSUPPORTED : C value 'hierarchy_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'hierarchy_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'style_set' : for field getter : missing Type

// UNSUPPORTED : C value 'style_set' : for field setter : missing Type

// UNSUPPORTED : C value 'direction_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'direction_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'grab_notify' : for field getter : missing Type

// UNSUPPORTED : C value 'grab_notify' : for field setter : missing Type

// UNSUPPORTED : C value 'child_notify' : for field getter : missing Type

// UNSUPPORTED : C value 'child_notify' : for field setter : missing Type

// UNSUPPORTED : C value 'draw' : for field getter : missing Type

// UNSUPPORTED : C value 'draw' : for field setter : missing Type

// UNSUPPORTED : C value 'get_request_mode' : for field getter : missing Type

// UNSUPPORTED : C value 'get_request_mode' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_height' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_height' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_width_for_height' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_width_for_height' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_width' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_width' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_for_width' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_for_width' : for field setter : missing Type

// UNSUPPORTED : C value 'mnemonic_activate' : for field getter : missing Type

// UNSUPPORTED : C value 'mnemonic_activate' : for field setter : missing Type

// UNSUPPORTED : C value 'grab_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'grab_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'focus' : for field getter : missing Type

// UNSUPPORTED : C value 'focus' : for field setter : missing Type

// UNSUPPORTED : C value 'move_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'move_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'keynav_failed' : for field getter : missing Type

// UNSUPPORTED : C value 'keynav_failed' : for field setter : missing Type

// UNSUPPORTED : C value 'event' : for field getter : missing Type

// UNSUPPORTED : C value 'event' : for field setter : missing Type

// UNSUPPORTED : C value 'button_press_event' : for field getter : missing Type

// UNSUPPORTED : C value 'button_press_event' : for field setter : missing Type

// UNSUPPORTED : C value 'button_release_event' : for field getter : missing Type

// UNSUPPORTED : C value 'button_release_event' : for field setter : missing Type

// UNSUPPORTED : C value 'scroll_event' : for field getter : missing Type

// UNSUPPORTED : C value 'scroll_event' : for field setter : missing Type

// UNSUPPORTED : C value 'motion_notify_event' : for field getter : missing Type

// UNSUPPORTED : C value 'motion_notify_event' : for field setter : missing Type

// UNSUPPORTED : C value 'delete_event' : for field getter : missing Type

// UNSUPPORTED : C value 'delete_event' : for field setter : missing Type

// UNSUPPORTED : C value 'destroy_event' : for field getter : missing Type

// UNSUPPORTED : C value 'destroy_event' : for field setter : missing Type

// UNSUPPORTED : C value 'key_press_event' : for field getter : missing Type

// UNSUPPORTED : C value 'key_press_event' : for field setter : missing Type

// UNSUPPORTED : C value 'key_release_event' : for field getter : missing Type

// UNSUPPORTED : C value 'key_release_event' : for field setter : missing Type

// UNSUPPORTED : C value 'enter_notify_event' : for field getter : missing Type

// UNSUPPORTED : C value 'enter_notify_event' : for field setter : missing Type

// UNSUPPORTED : C value 'leave_notify_event' : for field getter : missing Type

// UNSUPPORTED : C value 'leave_notify_event' : for field setter : missing Type

// UNSUPPORTED : C value 'configure_event' : for field getter : missing Type

// UNSUPPORTED : C value 'configure_event' : for field setter : missing Type

// UNSUPPORTED : C value 'focus_in_event' : for field getter : missing Type

// UNSUPPORTED : C value 'focus_in_event' : for field setter : missing Type

// UNSUPPORTED : C value 'focus_out_event' : for field getter : missing Type

// UNSUPPORTED : C value 'focus_out_event' : for field setter : missing Type

// UNSUPPORTED : C value 'map_event' : for field getter : missing Type

// UNSUPPORTED : C value 'map_event' : for field setter : missing Type

// UNSUPPORTED : C value 'unmap_event' : for field getter : missing Type

// UNSUPPORTED : C value 'unmap_event' : for field setter : missing Type

// UNSUPPORTED : C value 'property_notify_event' : for field getter : missing Type

// UNSUPPORTED : C value 'property_notify_event' : for field setter : missing Type

// UNSUPPORTED : C value 'selection_clear_event' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_clear_event' : for field setter : missing Type

// UNSUPPORTED : C value 'selection_request_event' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_request_event' : for field setter : missing Type

// UNSUPPORTED : C value 'selection_notify_event' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_notify_event' : for field setter : missing Type

// UNSUPPORTED : C value 'proximity_in_event' : for field getter : missing Type

// UNSUPPORTED : C value 'proximity_in_event' : for field setter : missing Type

// UNSUPPORTED : C value 'proximity_out_event' : for field getter : missing Type

// UNSUPPORTED : C value 'proximity_out_event' : for field setter : missing Type

// UNSUPPORTED : C value 'visibility_notify_event' : for field getter : missing Type

// UNSUPPORTED : C value 'visibility_notify_event' : for field setter : missing Type

// UNSUPPORTED : C value 'window_state_event' : for field getter : missing Type

// UNSUPPORTED : C value 'window_state_event' : for field setter : missing Type

// UNSUPPORTED : C value 'damage_event' : for field getter : missing Type

// UNSUPPORTED : C value 'damage_event' : for field setter : missing Type

// UNSUPPORTED : C value 'grab_broken_event' : for field getter : missing Type

// UNSUPPORTED : C value 'grab_broken_event' : for field setter : missing Type

// UNSUPPORTED : C value 'selection_get' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_get' : for field setter : missing Type

// UNSUPPORTED : C value 'selection_received' : for field getter : missing Type

// UNSUPPORTED : C value 'selection_received' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_begin' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_begin' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_end' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_end' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_data_get' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_data_get' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_data_delete' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_data_delete' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_leave' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_leave' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_motion' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_motion' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_drop' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_drop' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_data_received' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_data_received' : for field setter : missing Type

// UNSUPPORTED : C value 'drag_failed' : for field getter : missing Type

// UNSUPPORTED : C value 'drag_failed' : for field setter : missing Type

// UNSUPPORTED : C value 'popup_menu' : for field getter : missing Type

// UNSUPPORTED : C value 'popup_menu' : for field setter : missing Type

// UNSUPPORTED : C value 'show_help' : for field getter : missing Type

// UNSUPPORTED : C value 'show_help' : for field setter : missing Type

// UNSUPPORTED : C value 'get_accessible' : for field getter : missing Type

// UNSUPPORTED : C value 'get_accessible' : for field setter : missing Type

// UNSUPPORTED : C value 'screen_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'screen_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'can_activate_accel' : for field getter : missing Type

// UNSUPPORTED : C value 'can_activate_accel' : for field setter : missing Type

// UNSUPPORTED : C value 'composited_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'composited_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'query_tooltip' : for field getter : missing Type

// UNSUPPORTED : C value 'query_tooltip' : for field setter : missing Type

// UNSUPPORTED : C value 'compute_expand' : for field getter : missing Type

// UNSUPPORTED : C value 'compute_expand' : for field setter : missing Type

// UNSUPPORTED : C value 'adjust_size_request' : for field getter : missing Type

// UNSUPPORTED : C value 'adjust_size_request' : for field setter : missing Type

// UNSUPPORTED : C value 'adjust_size_allocation' : for field getter : missing Type

// UNSUPPORTED : C value 'adjust_size_allocation' : for field setter : missing Type

// UNSUPPORTED : C value 'style_updated' : for field getter : missing Type

// UNSUPPORTED : C value 'style_updated' : for field setter : missing Type

// UNSUPPORTED : C value 'touch_event' : for field getter : missing Type

// UNSUPPORTED : C value 'touch_event' : for field setter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_and_baseline_for_width' : for field getter : missing Type

// UNSUPPORTED : C value 'get_preferred_height_and_baseline_for_width' : for field setter : missing Type

// UNSUPPORTED : C value 'adjust_baseline_request' : for field getter : missing Type

// UNSUPPORTED : C value 'adjust_baseline_request' : for field setter : missing Type

// UNSUPPORTED : C value 'adjust_baseline_allocation' : for field getter : missing Type

// UNSUPPORTED : C value 'adjust_baseline_allocation' : for field setter : missing Type

// UNSUPPORTED : C value 'queue_draw_region' : for field getter : missing Type

// UNSUPPORTED : C value 'queue_draw_region' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved6' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved7' : for field setter : missing Type

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

// WidgetClassStruct creates an uninitialised WidgetClass.
func WidgetClassStruct() *WidgetClass {
	err := widgetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WidgetClass{native: widgetClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWidgetClass)
	return structGo
}
func finalizeWidgetClass(obj *WidgetClass) {
	widgetClassStruct.Free(obj.native)
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

// WidgetClassPrivateStruct creates an uninitialised WidgetClassPrivate.
func WidgetClassPrivateStruct() *WidgetClassPrivate {
	err := widgetClassPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WidgetClassPrivate{native: widgetClassPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWidgetClassPrivate)
	return structGo
}
func finalizeWidgetClassPrivate(obj *WidgetClassPrivate) {
	widgetClassPrivateStruct.Free(obj.native)
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

var widgetPathAppendForWidgetFunction *gi.Function
var widgetPathAppendForWidgetFunction_Once sync.Once

func widgetPathAppendForWidgetFunction_Set() error {
	var err error
	widgetPathAppendForWidgetFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathAppendForWidgetFunction, err = widgetPathStruct.InvokerNew("append_for_widget")
	})
	return err
}

// AppendForWidget is a representation of the C type gtk_widget_path_append_for_widget.
func (recv *WidgetPath) AppendForWidget(widget *Widget) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(widget.native)

	var ret gi.Argument

	err := widgetPathAppendForWidgetFunction_Set()
	if err == nil {
		ret = widgetPathAppendForWidgetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

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

var widgetPathIterHasQclassFunction *gi.Function
var widgetPathIterHasQclassFunction_Once sync.Once

func widgetPathIterHasQclassFunction_Set() error {
	var err error
	widgetPathIterHasQclassFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterHasQclassFunction, err = widgetPathStruct.InvokerNew("iter_has_qclass")
	})
	return err
}

// IterHasQclass is a representation of the C type gtk_widget_path_iter_has_qclass.
func (recv *WidgetPath) IterHasQclass(pos int32, qname glib.Quark) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetUint32(uint32(qname))

	var ret gi.Argument

	err := widgetPathIterHasQclassFunction_Set()
	if err == nil {
		ret = widgetPathIterHasQclassFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var widgetPathIterHasQnameFunction *gi.Function
var widgetPathIterHasQnameFunction_Once sync.Once

func widgetPathIterHasQnameFunction_Set() error {
	var err error
	widgetPathIterHasQnameFunction_Once.Do(func() {
		err = widgetPathStruct_Set()
		if err != nil {
			return
		}
		widgetPathIterHasQnameFunction, err = widgetPathStruct.InvokerNew("iter_has_qname")
	})
	return err
}

// IterHasQname is a representation of the C type gtk_widget_path_iter_has_qname.
func (recv *WidgetPath) IterHasQname(pos int32, qname glib.Quark) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(pos)
	inArgs[2].SetUint32(uint32(qname))

	var ret gi.Argument

	err := widgetPathIterHasQnameFunction_Set()
	if err == nil {
		ret = widgetPathIterHasQnameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_widget_path_iter_has_qregion' : parameter 'flags' of type 'RegionFlags' not supported

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

// WidgetPrivateStruct creates an uninitialised WidgetPrivate.
func WidgetPrivateStruct() *WidgetPrivate {
	err := widgetPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WidgetPrivate{native: widgetPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWidgetPrivate)
	return structGo
}
func finalizeWidgetPrivate(obj *WidgetPrivate) {
	widgetPrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *WindowAccessibleClass) FieldParentClass() *ContainerAccessibleClass {
	argValue := gi.FieldGet(windowAccessibleClassStruct, recv.native, "parent_class")
	value := &ContainerAccessibleClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *WindowAccessibleClass) SetFieldParentClass(value *ContainerAccessibleClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(windowAccessibleClassStruct, recv.native, "parent_class", argValue)
}

// WindowAccessibleClassStruct creates an uninitialised WindowAccessibleClass.
func WindowAccessibleClassStruct() *WindowAccessibleClass {
	err := windowAccessibleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowAccessibleClass{native: windowAccessibleClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowAccessibleClass)
	return structGo
}
func finalizeWindowAccessibleClass(obj *WindowAccessibleClass) {
	windowAccessibleClassStruct.Free(obj.native)
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

// WindowAccessiblePrivateStruct creates an uninitialised WindowAccessiblePrivate.
func WindowAccessiblePrivateStruct() *WindowAccessiblePrivate {
	err := windowAccessiblePrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowAccessiblePrivate{native: windowAccessiblePrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowAccessiblePrivate)
	return structGo
}
func finalizeWindowAccessiblePrivate(obj *WindowAccessiblePrivate) {
	windowAccessiblePrivateStruct.Free(obj.native)
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

// FieldParentClass returns the C field 'parent_class'.
func (recv *WindowClass) FieldParentClass() *BinClass {
	argValue := gi.FieldGet(windowClassStruct, recv.native, "parent_class")
	value := &BinClass{native: argValue.Pointer()}
	return value
}

// SetFieldParentClass sets the value of the C field 'parent_class'.
func (recv *WindowClass) SetFieldParentClass(value *BinClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.FieldSet(windowClassStruct, recv.native, "parent_class", argValue)
}

// UNSUPPORTED : C value 'set_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'set_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_focus' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_focus' : for field setter : missing Type

// UNSUPPORTED : C value 'activate_default' : for field getter : missing Type

// UNSUPPORTED : C value 'activate_default' : for field setter : missing Type

// UNSUPPORTED : C value 'keys_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'keys_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'enable_debugging' : for field getter : missing Type

// UNSUPPORTED : C value 'enable_debugging' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// WindowClassStruct creates an uninitialised WindowClass.
func WindowClassStruct() *WindowClass {
	err := windowClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowClass{native: windowClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowClass)
	return structGo
}
func finalizeWindowClass(obj *WindowClass) {
	windowClassStruct.Free(obj.native)
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

// WindowGeometryInfoStruct creates an uninitialised WindowGeometryInfo.
func WindowGeometryInfoStruct() *WindowGeometryInfo {
	err := windowGeometryInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowGeometryInfo{native: windowGeometryInfoStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowGeometryInfo)
	return structGo
}
func finalizeWindowGeometryInfo(obj *WindowGeometryInfo) {
	windowGeometryInfoStruct.Free(obj.native)
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

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value '_gtk_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_gtk_reserved4' : for field setter : missing Type

// WindowGroupClassStruct creates an uninitialised WindowGroupClass.
func WindowGroupClassStruct() *WindowGroupClass {
	err := windowGroupClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowGroupClass{native: windowGroupClassStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowGroupClass)
	return structGo
}
func finalizeWindowGroupClass(obj *WindowGroupClass) {
	windowGroupClassStruct.Free(obj.native)
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

// WindowGroupPrivateStruct creates an uninitialised WindowGroupPrivate.
func WindowGroupPrivateStruct() *WindowGroupPrivate {
	err := windowGroupPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowGroupPrivate{native: windowGroupPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowGroupPrivate)
	return structGo
}
func finalizeWindowGroupPrivate(obj *WindowGroupPrivate) {
	windowGroupPrivateStruct.Free(obj.native)
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

// WindowPrivateStruct creates an uninitialised WindowPrivate.
func WindowPrivateStruct() *WindowPrivate {
	err := windowPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &WindowPrivate{native: windowPrivateStruct.Alloc()}
	runtime.SetFinalizer(structGo, finalizeWindowPrivate)
	return structGo
}
func finalizeWindowPrivate(obj *WindowPrivate) {
	windowPrivateStruct.Free(obj.native)
}
