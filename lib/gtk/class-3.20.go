// This is a generated file - DO NOT EDIT
// +build gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
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

// Blacklisted : gtk_application_window_get_help_overlay

// Blacklisted : gtk_application_window_set_help_overlay

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

// Blacklisted : gtk_file_chooser_native_new

// Blacklisted : gtk_file_chooser_native_get_accept_label

// Blacklisted : gtk_file_chooser_native_get_cancel_label

// Blacklisted : gtk_file_chooser_native_set_accept_label

// Blacklisted : gtk_file_chooser_native_set_cancel_label

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

// Blacklisted : gtk_native_dialog_get_modal

// GetTitle is a wrapper around the C function gtk_native_dialog_get_title.
func (recv *NativeDialog) GetTitle() string {
	retC := C.gtk_native_dialog_get_title((*C.GtkNativeDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Blacklisted : gtk_native_dialog_get_transient_for

// Blacklisted : gtk_native_dialog_get_visible

// Blacklisted : gtk_native_dialog_hide

// Blacklisted : gtk_native_dialog_run

// Blacklisted : gtk_native_dialog_set_modal

// SetTitle is a wrapper around the C function gtk_native_dialog_set_title.
func (recv *NativeDialog) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_native_dialog_set_title((*C.GtkNativeDialog)(recv.native), c_title)

	return
}

// Blacklisted : gtk_native_dialog_set_transient_for

// Blacklisted : gtk_native_dialog_show

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

// Blacklisted : gtk_popover_get_constrain_to

// Blacklisted : gtk_popover_set_constrain_to

// Blacklisted : gtk_settings_reset_property

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

// Blacklisted : gtk_style_context_to_string

// Blacklisted : gtk_text_tag_changed

// Blacklisted : gtk_text_view_reset_cursor_blink

// Blacklisted : gtk_widget_get_allocated_size

// Blacklisted : gtk_widget_get_focus_on_click

// Blacklisted : gtk_widget_queue_allocate

// Blacklisted : gtk_widget_set_focus_on_click
