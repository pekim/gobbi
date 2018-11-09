// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void placessidebar_mountHandler(GObject *, GMountOperation *, gpointer);

	static gulong PlacesSidebar_signal_connect_mount(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mount", G_CALLBACK(placessidebar_mountHandler), data);
	}

*/
/*

	void placessidebar_unmountHandler(GObject *, GMountOperation *, gpointer);

	static gulong PlacesSidebar_signal_connect_unmount(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmount", G_CALLBACK(placessidebar_unmountHandler), data);
	}

*/
/*

	void shortcutswindow_closeHandler(GObject *, gpointer);

	static gulong ShortcutsWindow_signal_connect_close(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "close", G_CALLBACK(shortcutswindow_closeHandler), data);
	}

*/
/*

	void shortcutswindow_searchHandler(GObject *, gpointer);

	static gulong ShortcutsWindow_signal_connect_search(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "search", G_CALLBACK(shortcutswindow_searchHandler), data);
	}

*/
import "C"

// GetHelpOverlay is a wrapper around the C function gtk_application_window_get_help_overlay.
func (recv *ApplicationWindow) GetHelpOverlay() *ShortcutsWindow {
	retC := C.gtk_application_window_get_help_overlay((*C.GtkApplicationWindow)(recv.native))
	var retGo (*ShortcutsWindow)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ShortcutsWindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetHelpOverlay is a wrapper around the C function gtk_application_window_set_help_overlay.
func (recv *ApplicationWindow) SetHelpOverlay(helpOverlay *ShortcutsWindow) {
	c_help_overlay := (*C.GtkShortcutsWindow)(C.NULL)
	if helpOverlay != nil {
		c_help_overlay = (*C.GtkShortcutsWindow)(helpOverlay.ToC())
	}

	C.gtk_application_window_set_help_overlay((*C.GtkApplicationWindow)(recv.native), c_help_overlay)

	return
}

// FileChooserNative is a wrapper around the C record GtkFileChooserNative.
type FileChooserNative struct {
	native *C.GtkFileChooserNative
}

func FileChooserNativeNewFromC(u unsafe.Pointer) *FileChooserNative {
	c := (*C.GtkFileChooserNative)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserNative{native: c}

	return g
}

func (recv *FileChooserNative) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// NativeDialog upcasts to *NativeDialog
func (recv *FileChooserNative) NativeDialog() *NativeDialog {
	return NativeDialogNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileChooserNative) Object() *gobject.Object {
	return recv.NativeDialog().Object()
}

// CastToWidget down casts any arbitary Object to FileChooserNative.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileChooserNative.
func CastToFileChooserNative(object *gobject.Object) *FileChooserNative {
	return FileChooserNativeNewFromC(object.ToC())
}

// FileChooserNativeNew is a wrapper around the C function gtk_file_chooser_native_new.
func FileChooserNativeNew(title string, parent *Window, action FileChooserAction, acceptLabel string, cancelLabel string) *FileChooserNative {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_action := (C.GtkFileChooserAction)(action)

	c_accept_label := C.CString(acceptLabel)
	defer C.free(unsafe.Pointer(c_accept_label))

	c_cancel_label := C.CString(cancelLabel)
	defer C.free(unsafe.Pointer(c_cancel_label))

	retC := C.gtk_file_chooser_native_new(c_title, c_parent, c_action, c_accept_label, c_cancel_label)
	retGo := FileChooserNativeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAcceptLabel is a wrapper around the C function gtk_file_chooser_native_get_accept_label.
func (recv *FileChooserNative) GetAcceptLabel() string {
	retC := C.gtk_file_chooser_native_get_accept_label((*C.GtkFileChooserNative)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetCancelLabel is a wrapper around the C function gtk_file_chooser_native_get_cancel_label.
func (recv *FileChooserNative) GetCancelLabel() string {
	retC := C.gtk_file_chooser_native_get_cancel_label((*C.GtkFileChooserNative)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetAcceptLabel is a wrapper around the C function gtk_file_chooser_native_set_accept_label.
func (recv *FileChooserNative) SetAcceptLabel(acceptLabel string) {
	c_accept_label := C.CString(acceptLabel)
	defer C.free(unsafe.Pointer(c_accept_label))

	C.gtk_file_chooser_native_set_accept_label((*C.GtkFileChooserNative)(recv.native), c_accept_label)

	return
}

// SetCancelLabel is a wrapper around the C function gtk_file_chooser_native_set_cancel_label.
func (recv *FileChooserNative) SetCancelLabel(cancelLabel string) {
	c_cancel_label := C.CString(cancelLabel)
	defer C.free(unsafe.Pointer(c_cancel_label))

	C.gtk_file_chooser_native_set_cancel_label((*C.GtkFileChooserNative)(recv.native), c_cancel_label)

	return
}

// NativeDialog is a wrapper around the C record GtkNativeDialog.
type NativeDialog struct {
	native *C.GtkNativeDialog
	// parent_instance : record
}

func NativeDialogNewFromC(u unsafe.Pointer) *NativeDialog {
	c := (*C.GtkNativeDialog)(u)
	if c == nil {
		return nil
	}

	g := &NativeDialog{native: c}

	return g
}

func (recv *NativeDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *NativeDialog) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to NativeDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a NativeDialog.
func CastToNativeDialog(object *gobject.Object) *NativeDialog {
	return NativeDialogNewFromC(object.ToC())
}

// Destroy is a wrapper around the C function gtk_native_dialog_destroy.
func (recv *NativeDialog) Destroy() {
	C.gtk_native_dialog_destroy((*C.GtkNativeDialog)(recv.native))

	return
}

// GetModal is a wrapper around the C function gtk_native_dialog_get_modal.
func (recv *NativeDialog) GetModal() bool {
	retC := C.gtk_native_dialog_get_modal((*C.GtkNativeDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTitle is a wrapper around the C function gtk_native_dialog_get_title.
func (recv *NativeDialog) GetTitle() string {
	retC := C.gtk_native_dialog_get_title((*C.GtkNativeDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTransientFor is a wrapper around the C function gtk_native_dialog_get_transient_for.
func (recv *NativeDialog) GetTransientFor() *Window {
	retC := C.gtk_native_dialog_get_transient_for((*C.GtkNativeDialog)(recv.native))
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetVisible is a wrapper around the C function gtk_native_dialog_get_visible.
func (recv *NativeDialog) GetVisible() bool {
	retC := C.gtk_native_dialog_get_visible((*C.GtkNativeDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Hide is a wrapper around the C function gtk_native_dialog_hide.
func (recv *NativeDialog) Hide() {
	C.gtk_native_dialog_hide((*C.GtkNativeDialog)(recv.native))

	return
}

// Run is a wrapper around the C function gtk_native_dialog_run.
func (recv *NativeDialog) Run() int32 {
	retC := C.gtk_native_dialog_run((*C.GtkNativeDialog)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetModal is a wrapper around the C function gtk_native_dialog_set_modal.
func (recv *NativeDialog) SetModal(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gtk_native_dialog_set_modal((*C.GtkNativeDialog)(recv.native), c_modal)

	return
}

// SetTitle is a wrapper around the C function gtk_native_dialog_set_title.
func (recv *NativeDialog) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_native_dialog_set_title((*C.GtkNativeDialog)(recv.native), c_title)

	return
}

// SetTransientFor is a wrapper around the C function gtk_native_dialog_set_transient_for.
func (recv *NativeDialog) SetTransientFor(parent *Window) {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	C.gtk_native_dialog_set_transient_for((*C.GtkNativeDialog)(recv.native), c_parent)

	return
}

// Show is a wrapper around the C function gtk_native_dialog_show.
func (recv *NativeDialog) Show() {
	C.gtk_native_dialog_show((*C.GtkNativeDialog)(recv.native))

	return
}

// PadController is a wrapper around the C record GtkPadController.
type PadController struct {
	native *C.GtkPadController
}

func PadControllerNewFromC(u unsafe.Pointer) *PadController {
	c := (*C.GtkPadController)(u)
	if c == nil {
		return nil
	}

	g := &PadController{native: c}

	return g
}

func (recv *PadController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventController upcasts to *EventController
func (recv *PadController) EventController() *EventController {
	return EventControllerNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *PadController) Object() *gobject.Object {
	return recv.EventController().Object()
}

// CastToWidget down casts any arbitary Object to PadController.
// Exercise care, as this is a potentially dangerous function if the Object is not a PadController.
func CastToPadController(object *gobject.Object) *PadController {
	return PadControllerNewFromC(object.ToC())
}

type signalPlacesSidebarMountDetail struct {
	callback  PlacesSidebarSignalMountCallback
	handlerID C.gulong
}

var signalPlacesSidebarMountId int
var signalPlacesSidebarMountMap = make(map[int]signalPlacesSidebarMountDetail)
var signalPlacesSidebarMountLock sync.Mutex

// PlacesSidebarSignalMountCallback is a callback function for a 'mount' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalMountCallback func(mountOperation *gio.MountOperation)

/*
ConnectMount connects the callback to the 'mount' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectMount to remove it.
*/
func (recv *PlacesSidebar) ConnectMount(callback PlacesSidebarSignalMountCallback) int {
	signalPlacesSidebarMountLock.Lock()
	defer signalPlacesSidebarMountLock.Unlock()

	signalPlacesSidebarMountId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_mount(instance, C.gpointer(uintptr(signalPlacesSidebarMountId)))

	detail := signalPlacesSidebarMountDetail{callback, handlerID}
	signalPlacesSidebarMountMap[signalPlacesSidebarMountId] = detail

	return signalPlacesSidebarMountId
}

/*
DisconnectMount disconnects a callback from the 'mount' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectMount.
*/
func (recv *PlacesSidebar) DisconnectMount(connectionID int) {
	signalPlacesSidebarMountLock.Lock()
	defer signalPlacesSidebarMountLock.Unlock()

	detail, exists := signalPlacesSidebarMountMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarMountMap, connectionID)
}

//export placessidebar_mountHandler
func placessidebar_mountHandler(_ *C.GObject, c_mount_operation *C.GMountOperation, data C.gpointer) {
	mountOperation := gio.MountOperationNewFromC(unsafe.Pointer(c_mount_operation))

	index := int(uintptr(data))
	callback := signalPlacesSidebarMountMap[index].callback
	callback(mountOperation)
}

// Unsupported signal 'show-other-locations-with-flags' for PlacesSidebar : unsupported parameter open_flags : type PlacesOpenFlags :

type signalPlacesSidebarUnmountDetail struct {
	callback  PlacesSidebarSignalUnmountCallback
	handlerID C.gulong
}

var signalPlacesSidebarUnmountId int
var signalPlacesSidebarUnmountMap = make(map[int]signalPlacesSidebarUnmountDetail)
var signalPlacesSidebarUnmountLock sync.Mutex

// PlacesSidebarSignalUnmountCallback is a callback function for a 'unmount' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalUnmountCallback func(mountOperation *gio.MountOperation)

/*
ConnectUnmount connects the callback to the 'unmount' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectUnmount to remove it.
*/
func (recv *PlacesSidebar) ConnectUnmount(callback PlacesSidebarSignalUnmountCallback) int {
	signalPlacesSidebarUnmountLock.Lock()
	defer signalPlacesSidebarUnmountLock.Unlock()

	signalPlacesSidebarUnmountId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_unmount(instance, C.gpointer(uintptr(signalPlacesSidebarUnmountId)))

	detail := signalPlacesSidebarUnmountDetail{callback, handlerID}
	signalPlacesSidebarUnmountMap[signalPlacesSidebarUnmountId] = detail

	return signalPlacesSidebarUnmountId
}

/*
DisconnectUnmount disconnects a callback from the 'unmount' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectUnmount.
*/
func (recv *PlacesSidebar) DisconnectUnmount(connectionID int) {
	signalPlacesSidebarUnmountLock.Lock()
	defer signalPlacesSidebarUnmountLock.Unlock()

	detail, exists := signalPlacesSidebarUnmountMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarUnmountMap, connectionID)
}

//export placessidebar_unmountHandler
func placessidebar_unmountHandler(_ *C.GObject, c_mount_operation *C.GMountOperation, data C.gpointer) {
	mountOperation := gio.MountOperationNewFromC(unsafe.Pointer(c_mount_operation))

	index := int(uintptr(data))
	callback := signalPlacesSidebarUnmountMap[index].callback
	callback(mountOperation)
}

// GetConstrainTo is a wrapper around the C function gtk_popover_get_constrain_to.
func (recv *Popover) GetConstrainTo() PopoverConstraint {
	retC := C.gtk_popover_get_constrain_to((*C.GtkPopover)(recv.native))
	retGo := (PopoverConstraint)(retC)

	return retGo
}

// SetConstrainTo is a wrapper around the C function gtk_popover_set_constrain_to.
func (recv *Popover) SetConstrainTo(constraint PopoverConstraint) {
	c_constraint := (C.GtkPopoverConstraint)(constraint)

	C.gtk_popover_set_constrain_to((*C.GtkPopover)(recv.native), c_constraint)

	return
}

// ResetProperty is a wrapper around the C function gtk_settings_reset_property.
func (recv *Settings) ResetProperty(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_settings_reset_property((*C.GtkSettings)(recv.native), c_name)

	return
}

// ShortcutLabel is a wrapper around the C record GtkShortcutLabel.
type ShortcutLabel struct {
	native *C.GtkShortcutLabel
}

func ShortcutLabelNewFromC(u unsafe.Pointer) *ShortcutLabel {
	c := (*C.GtkShortcutLabel)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutLabel{native: c}

	return g
}

func (recv *ShortcutLabel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ShortcutLabel) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ShortcutLabel) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutLabel) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutLabel) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutLabel) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutLabel.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutLabel.
func CastToShortcutLabel(object *gobject.Object) *ShortcutLabel {
	return ShortcutLabelNewFromC(object.ToC())
}

// ShortcutsGroup is a wrapper around the C record GtkShortcutsGroup.
type ShortcutsGroup struct {
	native *C.GtkShortcutsGroup
}

func ShortcutsGroupNewFromC(u unsafe.Pointer) *ShortcutsGroup {
	c := (*C.GtkShortcutsGroup)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsGroup{native: c}

	return g
}

func (recv *ShortcutsGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ShortcutsGroup) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ShortcutsGroup) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutsGroup) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutsGroup) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutsGroup) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutsGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutsGroup.
func CastToShortcutsGroup(object *gobject.Object) *ShortcutsGroup {
	return ShortcutsGroupNewFromC(object.ToC())
}

// ShortcutsSection is a wrapper around the C record GtkShortcutsSection.
type ShortcutsSection struct {
	native *C.GtkShortcutsSection
}

func ShortcutsSectionNewFromC(u unsafe.Pointer) *ShortcutsSection {
	c := (*C.GtkShortcutsSection)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsSection{native: c}

	return g
}

func (recv *ShortcutsSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ShortcutsSection) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ShortcutsSection) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutsSection) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutsSection) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutsSection) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutsSection.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutsSection.
func CastToShortcutsSection(object *gobject.Object) *ShortcutsSection {
	return ShortcutsSectionNewFromC(object.ToC())
}

// Unsupported signal 'change-current-page' for ShortcutsSection : unsupported parameter object : type gint :

// ShortcutsShortcut is a wrapper around the C record GtkShortcutsShortcut.
type ShortcutsShortcut struct {
	native *C.GtkShortcutsShortcut
}

func ShortcutsShortcutNewFromC(u unsafe.Pointer) *ShortcutsShortcut {
	c := (*C.GtkShortcutsShortcut)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsShortcut{native: c}

	return g
}

func (recv *ShortcutsShortcut) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ShortcutsShortcut) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ShortcutsShortcut) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutsShortcut) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutsShortcut) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutsShortcut) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutsShortcut.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutsShortcut.
func CastToShortcutsShortcut(object *gobject.Object) *ShortcutsShortcut {
	return ShortcutsShortcutNewFromC(object.ToC())
}

// ShortcutsWindow is a wrapper around the C record GtkShortcutsWindow.
type ShortcutsWindow struct {
	native *C.GtkShortcutsWindow
	// window : record
}

func ShortcutsWindowNewFromC(u unsafe.Pointer) *ShortcutsWindow {
	c := (*C.GtkShortcutsWindow)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsWindow{native: c}

	return g
}

func (recv *ShortcutsWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window upcasts to *Window
func (recv *ShortcutsWindow) Window() *Window {
	return WindowNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ShortcutsWindow) Bin() *Bin {
	return recv.Window().Bin()
}

// Container upcasts to *Container
func (recv *ShortcutsWindow) Container() *Container {
	return recv.Window().Container()
}

// Widget upcasts to *Widget
func (recv *ShortcutsWindow) Widget() *Widget {
	return recv.Window().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ShortcutsWindow) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Window().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ShortcutsWindow) Object() *gobject.Object {
	return recv.Window().Object()
}

// CastToWidget down casts any arbitary Object to ShortcutsWindow.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutsWindow.
func CastToShortcutsWindow(object *gobject.Object) *ShortcutsWindow {
	return ShortcutsWindowNewFromC(object.ToC())
}

type signalShortcutsWindowCloseDetail struct {
	callback  ShortcutsWindowSignalCloseCallback
	handlerID C.gulong
}

var signalShortcutsWindowCloseId int
var signalShortcutsWindowCloseMap = make(map[int]signalShortcutsWindowCloseDetail)
var signalShortcutsWindowCloseLock sync.Mutex

// ShortcutsWindowSignalCloseCallback is a callback function for a 'close' signal emitted from a ShortcutsWindow.
type ShortcutsWindowSignalCloseCallback func()

/*
ConnectClose connects the callback to the 'close' signal for the ShortcutsWindow.

The returned value represents the connection, and may be passed to DisconnectClose to remove it.
*/
func (recv *ShortcutsWindow) ConnectClose(callback ShortcutsWindowSignalCloseCallback) int {
	signalShortcutsWindowCloseLock.Lock()
	defer signalShortcutsWindowCloseLock.Unlock()

	signalShortcutsWindowCloseId++
	instance := C.gpointer(recv.native)
	handlerID := C.ShortcutsWindow_signal_connect_close(instance, C.gpointer(uintptr(signalShortcutsWindowCloseId)))

	detail := signalShortcutsWindowCloseDetail{callback, handlerID}
	signalShortcutsWindowCloseMap[signalShortcutsWindowCloseId] = detail

	return signalShortcutsWindowCloseId
}

/*
DisconnectClose disconnects a callback from the 'close' signal for the ShortcutsWindow.

The connectionID should be a value returned from a call to ConnectClose.
*/
func (recv *ShortcutsWindow) DisconnectClose(connectionID int) {
	signalShortcutsWindowCloseLock.Lock()
	defer signalShortcutsWindowCloseLock.Unlock()

	detail, exists := signalShortcutsWindowCloseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalShortcutsWindowCloseMap, connectionID)
}

//export shortcutswindow_closeHandler
func shortcutswindow_closeHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalShortcutsWindowCloseMap[index].callback
	callback()
}

type signalShortcutsWindowSearchDetail struct {
	callback  ShortcutsWindowSignalSearchCallback
	handlerID C.gulong
}

var signalShortcutsWindowSearchId int
var signalShortcutsWindowSearchMap = make(map[int]signalShortcutsWindowSearchDetail)
var signalShortcutsWindowSearchLock sync.Mutex

// ShortcutsWindowSignalSearchCallback is a callback function for a 'search' signal emitted from a ShortcutsWindow.
type ShortcutsWindowSignalSearchCallback func()

/*
ConnectSearch connects the callback to the 'search' signal for the ShortcutsWindow.

The returned value represents the connection, and may be passed to DisconnectSearch to remove it.
*/
func (recv *ShortcutsWindow) ConnectSearch(callback ShortcutsWindowSignalSearchCallback) int {
	signalShortcutsWindowSearchLock.Lock()
	defer signalShortcutsWindowSearchLock.Unlock()

	signalShortcutsWindowSearchId++
	instance := C.gpointer(recv.native)
	handlerID := C.ShortcutsWindow_signal_connect_search(instance, C.gpointer(uintptr(signalShortcutsWindowSearchId)))

	detail := signalShortcutsWindowSearchDetail{callback, handlerID}
	signalShortcutsWindowSearchMap[signalShortcutsWindowSearchId] = detail

	return signalShortcutsWindowSearchId
}

/*
DisconnectSearch disconnects a callback from the 'search' signal for the ShortcutsWindow.

The connectionID should be a value returned from a call to ConnectSearch.
*/
func (recv *ShortcutsWindow) DisconnectSearch(connectionID int) {
	signalShortcutsWindowSearchLock.Lock()
	defer signalShortcutsWindowSearchLock.Unlock()

	detail, exists := signalShortcutsWindowSearchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalShortcutsWindowSearchMap, connectionID)
}

//export shortcutswindow_searchHandler
func shortcutswindow_searchHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalShortcutsWindowSearchMap[index].callback
	callback()
}

// ToString is a wrapper around the C function gtk_style_context_to_string.
func (recv *StyleContext) ToString(flags StyleContextPrintFlags) string {
	c_flags := (C.GtkStyleContextPrintFlags)(flags)

	retC := C.gtk_style_context_to_string((*C.GtkStyleContext)(recv.native), c_flags)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Changed is a wrapper around the C function gtk_text_tag_changed.
func (recv *TextTag) Changed(sizeChanged bool) {
	c_size_changed :=
		boolToGboolean(sizeChanged)

	C.gtk_text_tag_changed((*C.GtkTextTag)(recv.native), c_size_changed)

	return
}

// ResetCursorBlink is a wrapper around the C function gtk_text_view_reset_cursor_blink.
func (recv *TextView) ResetCursorBlink() {
	C.gtk_text_view_reset_cursor_blink((*C.GtkTextView)(recv.native))

	return
}

// Unsupported : gtk_widget_get_allocated_size : unsupported parameter allocation : Blacklisted record : GdkRectangle

// GetFocusOnClick is a wrapper around the C function gtk_widget_get_focus_on_click.
func (recv *Widget) GetFocusOnClick() bool {
	retC := C.gtk_widget_get_focus_on_click((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// QueueAllocate is a wrapper around the C function gtk_widget_queue_allocate.
func (recv *Widget) QueueAllocate() {
	C.gtk_widget_queue_allocate((*C.GtkWidget)(recv.native))

	return
}

// SetFocusOnClick is a wrapper around the C function gtk_widget_set_focus_on_click.
func (recv *Widget) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_widget_set_focus_on_click((*C.GtkWidget)(recv.native), c_focus_on_click)

	return
}
