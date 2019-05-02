// Code generated - DO NOT EDIT.
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void nativedialog_responseHandler(GObject *, gint, gpointer);

	static gulong NativeDialog_signal_connect_response(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "response", G_CALLBACK(nativedialog_responseHandler), data);
	}

*/
/*

	void placessidebar_mountHandler(GObject *, GMountOperation *, gpointer);

	static gulong PlacesSidebar_signal_connect_mount(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mount", G_CALLBACK(placessidebar_mountHandler), data);
	}

*/
/*

	void placessidebar_showOtherLocationsWithFlagsHandler(GObject *, GtkPlacesOpenFlags, gpointer);

	static gulong PlacesSidebar_signal_connect_show_other_locations_with_flags(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-other-locations-with-flags", G_CALLBACK(placessidebar_showOtherLocationsWithFlagsHandler), data);
	}

*/
/*

	void placessidebar_unmountHandler(GObject *, GMountOperation *, gpointer);

	static gulong PlacesSidebar_signal_connect_unmount(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmount", G_CALLBACK(placessidebar_unmountHandler), data);
	}

*/
/*

	gboolean shortcutssection_changeCurrentPageHandler(GObject *, gint, gpointer);

	static gulong ShortcutsSection_signal_connect_change_current_page(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "change-current-page", G_CALLBACK(shortcutssection_changeCurrentPageHandler), data);
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

type StyleContextPrintFlags C.GtkStyleContextPrintFlags

const (
	GTK_STYLE_CONTEXT_PRINT_NONE       StyleContextPrintFlags = 0
	GTK_STYLE_CONTEXT_PRINT_RECURSE    StyleContextPrintFlags = 1
	GTK_STYLE_CONTEXT_PRINT_SHOW_STYLE StyleContextPrintFlags = 2
)

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileChooserNative) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileChooserNative) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileChooserNative with another FileChooserNative, and returns true if they represent the same GObject.
func (recv *FileChooserNative) Equals(other *FileChooserNative) bool {
	return other.ToC() == recv.ToC()
}

// NativeDialog upcasts to *NativeDialog
func (recv *FileChooserNative) NativeDialog() *NativeDialog {
	return NativeDialogNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileChooserNative) Object() *gobject.Object {
	return recv.NativeDialog().Object()
}

// CastToWidget down casts any arbitrary Object to FileChooserNative.
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

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NativeDialog) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NativeDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NativeDialog with another NativeDialog, and returns true if they represent the same GObject.
func (recv *NativeDialog) Equals(other *NativeDialog) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *NativeDialog) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to NativeDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a NativeDialog.
func CastToNativeDialog(object *gobject.Object) *NativeDialog {
	return NativeDialogNewFromC(object.ToC())
}

type signalNativeDialogResponseDetail struct {
	callback  NativeDialogSignalResponseCallback
	handlerID C.gulong
}

var signalNativeDialogResponseId int
var signalNativeDialogResponseMap = make(map[int]signalNativeDialogResponseDetail)
var signalNativeDialogResponseLock sync.RWMutex

// NativeDialogSignalResponseCallback is a callback function for a 'response' signal emitted from a NativeDialog.
type NativeDialogSignalResponseCallback func(responseId int32)

/*
ConnectResponse connects the callback to the 'response' signal for the NativeDialog.

The returned value represents the connection, and may be passed to DisconnectResponse to remove it.
*/
func (recv *NativeDialog) ConnectResponse(callback NativeDialogSignalResponseCallback) int {
	signalNativeDialogResponseLock.Lock()
	defer signalNativeDialogResponseLock.Unlock()

	signalNativeDialogResponseId++
	instance := C.gpointer(recv.native)
	handlerID := C.NativeDialog_signal_connect_response(instance, C.gpointer(uintptr(signalNativeDialogResponseId)))

	detail := signalNativeDialogResponseDetail{callback, handlerID}
	signalNativeDialogResponseMap[signalNativeDialogResponseId] = detail

	return signalNativeDialogResponseId
}

/*
DisconnectResponse disconnects a callback from the 'response' signal for the NativeDialog.

The connectionID should be a value returned from a call to ConnectResponse.
*/
func (recv *NativeDialog) DisconnectResponse(connectionID int) {
	signalNativeDialogResponseLock.Lock()
	defer signalNativeDialogResponseLock.Unlock()

	detail, exists := signalNativeDialogResponseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalNativeDialogResponseMap, connectionID)
}

//export nativedialog_responseHandler
func nativedialog_responseHandler(_ *C.GObject, c_response_id C.gint, data C.gpointer) {
	signalNativeDialogResponseLock.RLock()
	defer signalNativeDialogResponseLock.RUnlock()

	responseId := int32(c_response_id)

	index := int(uintptr(data))
	callback := signalNativeDialogResponseMap[index].callback
	callback(responseId)
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PadController) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PadController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PadController with another PadController, and returns true if they represent the same GObject.
func (recv *PadController) Equals(other *PadController) bool {
	return other.ToC() == recv.ToC()
}

// EventController upcasts to *EventController
func (recv *PadController) EventController() *EventController {
	return EventControllerNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *PadController) Object() *gobject.Object {
	return recv.EventController().Object()
}

// CastToWidget down casts any arbitrary Object to PadController.
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
var signalPlacesSidebarMountLock sync.RWMutex

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
	signalPlacesSidebarMountLock.RLock()
	defer signalPlacesSidebarMountLock.RUnlock()

	mountOperation := gio.MountOperationNewFromC(unsafe.Pointer(c_mount_operation))

	index := int(uintptr(data))
	callback := signalPlacesSidebarMountMap[index].callback
	callback(mountOperation)
}

type signalPlacesSidebarShowOtherLocationsWithFlagsDetail struct {
	callback  PlacesSidebarSignalShowOtherLocationsWithFlagsCallback
	handlerID C.gulong
}

var signalPlacesSidebarShowOtherLocationsWithFlagsId int
var signalPlacesSidebarShowOtherLocationsWithFlagsMap = make(map[int]signalPlacesSidebarShowOtherLocationsWithFlagsDetail)
var signalPlacesSidebarShowOtherLocationsWithFlagsLock sync.RWMutex

// PlacesSidebarSignalShowOtherLocationsWithFlagsCallback is a callback function for a 'show-other-locations-with-flags' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowOtherLocationsWithFlagsCallback func(openFlags PlacesOpenFlags)

/*
ConnectShowOtherLocationsWithFlags connects the callback to the 'show-other-locations-with-flags' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectShowOtherLocationsWithFlags to remove it.
*/
func (recv *PlacesSidebar) ConnectShowOtherLocationsWithFlags(callback PlacesSidebarSignalShowOtherLocationsWithFlagsCallback) int {
	signalPlacesSidebarShowOtherLocationsWithFlagsLock.Lock()
	defer signalPlacesSidebarShowOtherLocationsWithFlagsLock.Unlock()

	signalPlacesSidebarShowOtherLocationsWithFlagsId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_show_other_locations_with_flags(instance, C.gpointer(uintptr(signalPlacesSidebarShowOtherLocationsWithFlagsId)))

	detail := signalPlacesSidebarShowOtherLocationsWithFlagsDetail{callback, handlerID}
	signalPlacesSidebarShowOtherLocationsWithFlagsMap[signalPlacesSidebarShowOtherLocationsWithFlagsId] = detail

	return signalPlacesSidebarShowOtherLocationsWithFlagsId
}

/*
DisconnectShowOtherLocationsWithFlags disconnects a callback from the 'show-other-locations-with-flags' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectShowOtherLocationsWithFlags.
*/
func (recv *PlacesSidebar) DisconnectShowOtherLocationsWithFlags(connectionID int) {
	signalPlacesSidebarShowOtherLocationsWithFlagsLock.Lock()
	defer signalPlacesSidebarShowOtherLocationsWithFlagsLock.Unlock()

	detail, exists := signalPlacesSidebarShowOtherLocationsWithFlagsMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarShowOtherLocationsWithFlagsMap, connectionID)
}

//export placessidebar_showOtherLocationsWithFlagsHandler
func placessidebar_showOtherLocationsWithFlagsHandler(_ *C.GObject, c_open_flags C.GtkPlacesOpenFlags, data C.gpointer) {
	signalPlacesSidebarShowOtherLocationsWithFlagsLock.RLock()
	defer signalPlacesSidebarShowOtherLocationsWithFlagsLock.RUnlock()

	openFlags := PlacesOpenFlags(c_open_flags)

	index := int(uintptr(data))
	callback := signalPlacesSidebarShowOtherLocationsWithFlagsMap[index].callback
	callback(openFlags)
}

type signalPlacesSidebarUnmountDetail struct {
	callback  PlacesSidebarSignalUnmountCallback
	handlerID C.gulong
}

var signalPlacesSidebarUnmountId int
var signalPlacesSidebarUnmountMap = make(map[int]signalPlacesSidebarUnmountDetail)
var signalPlacesSidebarUnmountLock sync.RWMutex

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
	signalPlacesSidebarUnmountLock.RLock()
	defer signalPlacesSidebarUnmountLock.RUnlock()

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ShortcutLabel) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ShortcutLabel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutLabel with another ShortcutLabel, and returns true if they represent the same GObject.
func (recv *ShortcutLabel) Equals(other *ShortcutLabel) bool {
	return other.ToC() == recv.ToC()
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

// CastToWidget down casts any arbitrary Object to ShortcutLabel.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ShortcutsGroup) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ShortcutsGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsGroup with another ShortcutsGroup, and returns true if they represent the same GObject.
func (recv *ShortcutsGroup) Equals(other *ShortcutsGroup) bool {
	return other.ToC() == recv.ToC()
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

// CastToWidget down casts any arbitrary Object to ShortcutsGroup.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ShortcutsSection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ShortcutsSection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsSection with another ShortcutsSection, and returns true if they represent the same GObject.
func (recv *ShortcutsSection) Equals(other *ShortcutsSection) bool {
	return other.ToC() == recv.ToC()
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

// CastToWidget down casts any arbitrary Object to ShortcutsSection.
// Exercise care, as this is a potentially dangerous function if the Object is not a ShortcutsSection.
func CastToShortcutsSection(object *gobject.Object) *ShortcutsSection {
	return ShortcutsSectionNewFromC(object.ToC())
}

type signalShortcutsSectionChangeCurrentPageDetail struct {
	callback  ShortcutsSectionSignalChangeCurrentPageCallback
	handlerID C.gulong
}

var signalShortcutsSectionChangeCurrentPageId int
var signalShortcutsSectionChangeCurrentPageMap = make(map[int]signalShortcutsSectionChangeCurrentPageDetail)
var signalShortcutsSectionChangeCurrentPageLock sync.RWMutex

// ShortcutsSectionSignalChangeCurrentPageCallback is a callback function for a 'change-current-page' signal emitted from a ShortcutsSection.
type ShortcutsSectionSignalChangeCurrentPageCallback func(object int32) bool

/*
ConnectChangeCurrentPage connects the callback to the 'change-current-page' signal for the ShortcutsSection.

The returned value represents the connection, and may be passed to DisconnectChangeCurrentPage to remove it.
*/
func (recv *ShortcutsSection) ConnectChangeCurrentPage(callback ShortcutsSectionSignalChangeCurrentPageCallback) int {
	signalShortcutsSectionChangeCurrentPageLock.Lock()
	defer signalShortcutsSectionChangeCurrentPageLock.Unlock()

	signalShortcutsSectionChangeCurrentPageId++
	instance := C.gpointer(recv.native)
	handlerID := C.ShortcutsSection_signal_connect_change_current_page(instance, C.gpointer(uintptr(signalShortcutsSectionChangeCurrentPageId)))

	detail := signalShortcutsSectionChangeCurrentPageDetail{callback, handlerID}
	signalShortcutsSectionChangeCurrentPageMap[signalShortcutsSectionChangeCurrentPageId] = detail

	return signalShortcutsSectionChangeCurrentPageId
}

/*
DisconnectChangeCurrentPage disconnects a callback from the 'change-current-page' signal for the ShortcutsSection.

The connectionID should be a value returned from a call to ConnectChangeCurrentPage.
*/
func (recv *ShortcutsSection) DisconnectChangeCurrentPage(connectionID int) {
	signalShortcutsSectionChangeCurrentPageLock.Lock()
	defer signalShortcutsSectionChangeCurrentPageLock.Unlock()

	detail, exists := signalShortcutsSectionChangeCurrentPageMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalShortcutsSectionChangeCurrentPageMap, connectionID)
}

//export shortcutssection_changeCurrentPageHandler
func shortcutssection_changeCurrentPageHandler(_ *C.GObject, c_object C.gint, data C.gpointer) C.gboolean {
	signalShortcutsSectionChangeCurrentPageLock.RLock()
	defer signalShortcutsSectionChangeCurrentPageLock.RUnlock()

	object := int32(c_object)

	index := int(uintptr(data))
	callback := signalShortcutsSectionChangeCurrentPageMap[index].callback
	retGo := callback(object)
	retC :=
		boolToGboolean(retGo)
	return retC
}

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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ShortcutsShortcut) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ShortcutsShortcut) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsShortcut with another ShortcutsShortcut, and returns true if they represent the same GObject.
func (recv *ShortcutsShortcut) Equals(other *ShortcutsShortcut) bool {
	return other.ToC() == recv.ToC()
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

// CastToWidget down casts any arbitrary Object to ShortcutsShortcut.
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

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ShortcutsWindow) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ShortcutsWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsWindow with another ShortcutsWindow, and returns true if they represent the same GObject.
func (recv *ShortcutsWindow) Equals(other *ShortcutsWindow) bool {
	return other.ToC() == recv.ToC()
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

// CastToWidget down casts any arbitrary Object to ShortcutsWindow.
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
var signalShortcutsWindowCloseLock sync.RWMutex

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
	signalShortcutsWindowCloseLock.RLock()
	defer signalShortcutsWindowCloseLock.RUnlock()

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
var signalShortcutsWindowSearchLock sync.RWMutex

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
	signalShortcutsWindowSearchLock.RLock()
	defer signalShortcutsWindowSearchLock.RUnlock()

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

// GetAllocatedSize is a wrapper around the C function gtk_widget_get_allocated_size.
func (recv *Widget) GetAllocatedSize() (*gdk.Rectangle, int32) {
	var c_allocation C.GdkRectangle

	var c_baseline C.int

	C.gtk_widget_get_allocated_size((*C.GtkWidget)(recv.native), &c_allocation, &c_baseline)

	allocation := gdk.RectangleNewFromC(unsafe.Pointer(&c_allocation))

	baseline := (int32)(c_baseline)

	return allocation, baseline
}

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

const LEVEL_BAR_OFFSET_FULL string = C.GTK_LEVEL_BAR_OFFSET_FULL

type PopoverConstraint C.GtkPopoverConstraint

const (
	GTK_POPOVER_CONSTRAINT_NONE   PopoverConstraint = 0
	GTK_POPOVER_CONSTRAINT_WINDOW PopoverConstraint = 1
)

type ShortcutType C.GtkShortcutType

const (
	GTK_SHORTCUT_ACCELERATOR                     ShortcutType = 0
	GTK_SHORTCUT_GESTURE_PINCH                   ShortcutType = 1
	GTK_SHORTCUT_GESTURE_STRETCH                 ShortcutType = 2
	GTK_SHORTCUT_GESTURE_ROTATE_CLOCKWISE        ShortcutType = 3
	GTK_SHORTCUT_GESTURE_ROTATE_COUNTERCLOCKWISE ShortcutType = 4
	GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_LEFT   ShortcutType = 5
	GTK_SHORTCUT_GESTURE_TWO_FINGER_SWIPE_RIGHT  ShortcutType = 6
	GTK_SHORTCUT_GESTURE                         ShortcutType = 7
)

// RenderBackgroundGetClip is a wrapper around the C function gtk_render_background_get_clip.
func RenderBackgroundGetClip(context *StyleContext, x float64, y float64, width float64, height float64) *gdk.Rectangle {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	var c_out_clip C.GdkRectangle

	C.gtk_render_background_get_clip(c_context, c_x, c_y, c_width, c_height, &c_out_clip)

	outClip := gdk.RectangleNewFromC(unsafe.Pointer(&c_out_clip))

	return outClip
}

// FileChooserNativeClass is a wrapper around the C record GtkFileChooserNativeClass.
type FileChooserNativeClass struct {
	native *C.GtkFileChooserNativeClass
	// parent_class : record
}

func FileChooserNativeClassNewFromC(u unsafe.Pointer) *FileChooserNativeClass {
	c := (*C.GtkFileChooserNativeClass)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserNativeClass{native: c}

	return g
}

func (recv *FileChooserNativeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileChooserNativeClass with another FileChooserNativeClass, and returns true if they represent the same GObject.
func (recv *FileChooserNativeClass) Equals(other *FileChooserNativeClass) bool {
	return other.ToC() == recv.ToC()
}

// NativeDialogClass is a wrapper around the C record GtkNativeDialogClass.
type NativeDialogClass struct {
	native *C.GtkNativeDialogClass
	// parent_class : record
	// no type for response
	// no type for show
	// no type for hide
	// no type for _gtk_reserved1
	// no type for _gtk_reserved2
	// no type for _gtk_reserved3
	// no type for _gtk_reserved4
}

func NativeDialogClassNewFromC(u unsafe.Pointer) *NativeDialogClass {
	c := (*C.GtkNativeDialogClass)(u)
	if c == nil {
		return nil
	}

	g := &NativeDialogClass{native: c}

	return g
}

func (recv *NativeDialogClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NativeDialogClass with another NativeDialogClass, and returns true if they represent the same GObject.
func (recv *NativeDialogClass) Equals(other *NativeDialogClass) bool {
	return other.ToC() == recv.ToC()
}

// PadActionEntry is a wrapper around the C record GtkPadActionEntry.
type PadActionEntry struct {
	native     *C.GtkPadActionEntry
	Type       PadActionType
	Index      int32
	Mode       int32
	Label      string
	ActionName string
}

func PadActionEntryNewFromC(u unsafe.Pointer) *PadActionEntry {
	c := (*C.GtkPadActionEntry)(u)
	if c == nil {
		return nil
	}

	g := &PadActionEntry{
		ActionName: C.GoString(c.action_name),
		Index:      (int32)(c.index),
		Label:      C.GoString(c.label),
		Mode:       (int32)(c.mode),
		Type:       (PadActionType)(c._type),
		native:     c,
	}

	return g
}

func (recv *PadActionEntry) ToC() unsafe.Pointer {
	recv.native._type =
		(C.GtkPadActionType)(recv.Type)
	recv.native.index =
		(C.gint)(recv.Index)
	recv.native.mode =
		(C.gint)(recv.Mode)
	recv.native.label =
		C.CString(recv.Label)
	recv.native.action_name =
		C.CString(recv.ActionName)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PadActionEntry with another PadActionEntry, and returns true if they represent the same GObject.
func (recv *PadActionEntry) Equals(other *PadActionEntry) bool {
	return other.ToC() == recv.ToC()
}

// PadControllerClass is a wrapper around the C record GtkPadControllerClass.
type PadControllerClass struct {
	native *C.GtkPadControllerClass
}

func PadControllerClassNewFromC(u unsafe.Pointer) *PadControllerClass {
	c := (*C.GtkPadControllerClass)(u)
	if c == nil {
		return nil
	}

	g := &PadControllerClass{native: c}

	return g
}

func (recv *PadControllerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PadControllerClass with another PadControllerClass, and returns true if they represent the same GObject.
func (recv *PadControllerClass) Equals(other *PadControllerClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutLabelClass is a wrapper around the C record GtkShortcutLabelClass.
type ShortcutLabelClass struct {
	native *C.GtkShortcutLabelClass
}

func ShortcutLabelClassNewFromC(u unsafe.Pointer) *ShortcutLabelClass {
	c := (*C.GtkShortcutLabelClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutLabelClass{native: c}

	return g
}

func (recv *ShortcutLabelClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutLabelClass with another ShortcutLabelClass, and returns true if they represent the same GObject.
func (recv *ShortcutLabelClass) Equals(other *ShortcutLabelClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutsGroupClass is a wrapper around the C record GtkShortcutsGroupClass.
type ShortcutsGroupClass struct {
	native *C.GtkShortcutsGroupClass
}

func ShortcutsGroupClassNewFromC(u unsafe.Pointer) *ShortcutsGroupClass {
	c := (*C.GtkShortcutsGroupClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsGroupClass{native: c}

	return g
}

func (recv *ShortcutsGroupClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsGroupClass with another ShortcutsGroupClass, and returns true if they represent the same GObject.
func (recv *ShortcutsGroupClass) Equals(other *ShortcutsGroupClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutsSectionClass is a wrapper around the C record GtkShortcutsSectionClass.
type ShortcutsSectionClass struct {
	native *C.GtkShortcutsSectionClass
}

func ShortcutsSectionClassNewFromC(u unsafe.Pointer) *ShortcutsSectionClass {
	c := (*C.GtkShortcutsSectionClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsSectionClass{native: c}

	return g
}

func (recv *ShortcutsSectionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsSectionClass with another ShortcutsSectionClass, and returns true if they represent the same GObject.
func (recv *ShortcutsSectionClass) Equals(other *ShortcutsSectionClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutsShortcutClass is a wrapper around the C record GtkShortcutsShortcutClass.
type ShortcutsShortcutClass struct {
	native *C.GtkShortcutsShortcutClass
}

func ShortcutsShortcutClassNewFromC(u unsafe.Pointer) *ShortcutsShortcutClass {
	c := (*C.GtkShortcutsShortcutClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsShortcutClass{native: c}

	return g
}

func (recv *ShortcutsShortcutClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsShortcutClass with another ShortcutsShortcutClass, and returns true if they represent the same GObject.
func (recv *ShortcutsShortcutClass) Equals(other *ShortcutsShortcutClass) bool {
	return other.ToC() == recv.ToC()
}

// ShortcutsWindowClass is a wrapper around the C record GtkShortcutsWindowClass.
type ShortcutsWindowClass struct {
	native *C.GtkShortcutsWindowClass
	// parent_class : record
	// no type for close
	// no type for search
}

func ShortcutsWindowClassNewFromC(u unsafe.Pointer) *ShortcutsWindowClass {
	c := (*C.GtkShortcutsWindowClass)(u)
	if c == nil {
		return nil
	}

	g := &ShortcutsWindowClass{native: c}

	return g
}

func (recv *ShortcutsWindowClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ShortcutsWindowClass with another ShortcutsWindowClass, and returns true if they represent the same GObject.
func (recv *ShortcutsWindowClass) Equals(other *ShortcutsWindowClass) bool {
	return other.ToC() == recv.ToC()
}

// StartsTag is a wrapper around the C function gtk_text_iter_starts_tag.
func (recv *TextIter) StartsTag(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(C.NULL)
	if tag != nil {
		c_tag = (*C.GtkTextTag)(tag.ToC())
	}

	retC := C.gtk_text_iter_starts_tag((*C.GtkTextIter)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// GetCssName is a wrapper around the C function gtk_widget_class_get_css_name.
func (recv *WidgetClass) GetCssName() string {
	retC := C.gtk_widget_class_get_css_name((*C.GtkWidgetClass)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetCssName is a wrapper around the C function gtk_widget_class_set_css_name.
func (recv *WidgetClass) SetCssName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_class_set_css_name((*C.GtkWidgetClass)(recv.native), c_name)

	return
}

// IterGetObjectName is a wrapper around the C function gtk_widget_path_iter_get_object_name.
func (recv *WidgetPath) IterGetObjectName(pos int32) string {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_object_name((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := C.GoString(retC)

	return retGo
}

// IterSetObjectName is a wrapper around the C function gtk_widget_path_iter_set_object_name.
func (recv *WidgetPath) IterSetObjectName(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_set_object_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}
