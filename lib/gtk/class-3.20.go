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

// Gets the #GtkShortcutsWindow that has been set up with
// a prior call to gtk_application_window_set_help_overlay().
/*

C function : gtk_application_window_get_help_overlay
*/
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

// Associates a shortcuts window with the application window, and
// sets up an action with the name win.show-help-overlay to present
// it.
//
// @window takes resposibility for destroying @help_overlay.
/*

C function : gtk_application_window_set_help_overlay
*/
func (recv *ApplicationWindow) SetHelpOverlay(helpOverlay *ShortcutsWindow) {
	c_help_overlay := (*C.GtkShortcutsWindow)(C.NULL)
	if helpOverlay != nil {
		c_help_overlay = (*C.GtkShortcutsWindow)(helpOverlay.ToC())
	}

	C.gtk_application_window_set_help_overlay((*C.GtkApplicationWindow)(recv.native), c_help_overlay)

	return
}

// #GtkFileChooserNative is an abstraction of a dialog box suitable
// for use with “File/Open” or “File/Save as” commands. By default, this
// just uses a #GtkFileChooserDialog to implement the actual dialog.
// However, on certain platforms, such as Windows and macOS, the native platform
// file chooser is used instead. When the application is running in a
// sandboxed environment without direct filesystem access (such as Flatpak),
// #GtkFileChooserNative may call the proper APIs (portals) to let the user
// choose a file and make it available to the application.
//
// While the API of #GtkFileChooserNative closely mirrors #GtkFileChooserDialog, the main
// difference is that there is no access to any #GtkWindow or #GtkWidget for the dialog.
// This is required, as there may not be one in the case of a platform native dialog.
// Showing, hiding and running the dialog is handled by the #GtkNativeDialog functions.
//
// ## Typical usage ## {#gtkfilechoosernative-typical-usage}
//
// In the simplest of cases, you can the following code to use
// #GtkFileChooserDialog to select a file for opening:
//
// |[
// GtkFileChooserNative *native;
// GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_OPEN;
// gint res;
//
// native = gtk_file_chooser_native_new ("Open File",
// parent_window,
// action,
// "_Open",
// "_Cancel");
//
// res = gtk_native_dialog_run (GTK_NATIVE_DIALOG (native));
// if (res == GTK_RESPONSE_ACCEPT)
// {
// char *filename;
// GtkFileChooser *chooser = GTK_FILE_CHOOSER (native);
// filename = gtk_file_chooser_get_filename (chooser);
// open_file (filename);
// g_free (filename);
// }
//
// g_object_unref (native);
// ]|
//
// To use a dialog for saving, you can use this:
//
// |[
// GtkFileChooserNative *native;
// GtkFileChooser *chooser;
// GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_SAVE;
// gint res;
//
// native = gtk_file_chooser_native_new ("Save File",
// parent_window,
// action,
// "_Save",
// "_Cancel");
// chooser = GTK_FILE_CHOOSER (native);
//
// gtk_file_chooser_set_do_overwrite_confirmation (chooser, TRUE);
//
// if (user_edited_a_new_document)
// gtk_file_chooser_set_current_name (chooser,
// _("Untitled document"));
// else
// gtk_file_chooser_set_filename (chooser,
// existing_filename);
//
// res = gtk_native_dialog_run (GTK_NATIVE_DIALOG (native));
// if (res == GTK_RESPONSE_ACCEPT)
// {
// char *filename;
//
// filename = gtk_file_chooser_get_filename (chooser);
// save_to_file (filename);
// g_free (filename);
// }
//
// g_object_unref (native);
// ]|
//
// For more information on how to best set up a file dialog, see #GtkFileChooserDialog.
//
// ## Response Codes ## {#gtkfilechooserdialognative-responses}
//
// #GtkFileChooserNative inherits from #GtkNativeDialog, which means it
// will return #GTK_RESPONSE_ACCEPT if the user accepted, and
// #GTK_RESPONSE_CANCEL if he pressed cancel. It can also return
// #GTK_RESPONSE_DELETE_EVENT if the window was unexpectedly closed.
//
// ## Differences from #GtkFileChooserDialog ##  {#gtkfilechooserdialognative-differences}
//
// There are a few things in the GtkFileChooser API that are not
// possible to use with #GtkFileChooserNative, as such use would
// prohibit the use of a native dialog.
//
// There is no support for the signals that are emitted when the user
// navigates in the dialog, including:
// * #GtkFileChooser::current-folder-changed
// * #GtkFileChooser::selection-changed
// * #GtkFileChooser::file-activated
// * #GtkFileChooser::confirm-overwrite
//
// You can also not use the methods that directly control user navigation:
// * gtk_file_chooser_unselect_filename()
// * gtk_file_chooser_select_all()
// * gtk_file_chooser_unselect_all()
//
// If you need any of the above you will have to use #GtkFileChooserDialog directly.
//
// No operations that change the the dialog work while the dialog is
// visible. Set all the properties that are required before showing the dialog.
//
// ## Win32 details ## {#gtkfilechooserdialognative-win32}
//
// On windows the IFileDialog implementation (added in Windows Vista) is
// used. It supports many of the features that #GtkFileChooserDialog
// does, but there are some things it does not handle:
//
// * Extra widgets added with gtk_file_chooser_set_extra_widget().
//
// * Use of custom previews by connecting to #GtkFileChooser::update-preview.
//
// * Any #GtkFileFilter added using a mimetype or custom filter.
//
// If any of these features are used the regular #GtkFileChooserDialog
// will be used in place of the native one.
//
// ## Portal details ## {#gtkfilechooserdialognative-portal}
//
// When the org.freedesktop.portal.FileChooser portal is available on the
// session bus, it is used to bring up an out-of-process file chooser. Depending
// on the kind of session the application is running in, this may or may not
// be a GTK+ file chooser. In this situation, the following things are not
// supported and will be silently ignored:
//
// * Extra widgets added with gtk_file_chooser_set_extra_widget().
//
// * Use of custom previews by connecting to #GtkFileChooser::update-preview.
//
// * Any #GtkFileFilter added with a custom filter.
//
// ## macOS details ## {#gtkfilechooserdialognative-macos}
//
// On macOS the NSSavePanel and NSOpenPanel classes are used to provide native
// file chooser dialogs. Some features provided by #GtkFileChooserDialog are
// not supported:
//
// * Extra widgets added with gtk_file_chooser_set_extra_widget(), unless the
// widget is an instance of GtkLabel, in which case the label text will be used
// to set the NSSavePanel message instance property.
//
// * Use of custom previews by connecting to #GtkFileChooser::update-preview.
//
// * Any #GtkFileFilter added with a custom filter.
//
// * Shortcut folders.
/*

C record/class : GtkFileChooserNative
*/
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

// Creates a new #GtkFileChooserNative.
/*

C function : gtk_file_chooser_native_new
*/
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

// Retrieves the custom label text for the accept button.
/*

C function : gtk_file_chooser_native_get_accept_label
*/
func (recv *FileChooserNative) GetAcceptLabel() string {
	retC := C.gtk_file_chooser_native_get_accept_label((*C.GtkFileChooserNative)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Retrieves the custom label text for the cancel button.
/*

C function : gtk_file_chooser_native_get_cancel_label
*/
func (recv *FileChooserNative) GetCancelLabel() string {
	retC := C.gtk_file_chooser_native_get_cancel_label((*C.GtkFileChooserNative)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets the custom label text for the accept button.
//
// If characters in @label are preceded by an underscore, they are underlined.
// If you need a literal underscore character in a label, use “__” (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic.
// Pressing Alt and that key activates the button.
/*

C function : gtk_file_chooser_native_set_accept_label
*/
func (recv *FileChooserNative) SetAcceptLabel(acceptLabel string) {
	c_accept_label := C.CString(acceptLabel)
	defer C.free(unsafe.Pointer(c_accept_label))

	C.gtk_file_chooser_native_set_accept_label((*C.GtkFileChooserNative)(recv.native), c_accept_label)

	return
}

// Sets the custom label text for the cancel button.
//
// If characters in @label are preceded by an underscore, they are underlined.
// If you need a literal underscore character in a label, use “__” (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic.
// Pressing Alt and that key activates the button.
/*

C function : gtk_file_chooser_native_set_cancel_label
*/
func (recv *FileChooserNative) SetCancelLabel(cancelLabel string) {
	c_cancel_label := C.CString(cancelLabel)
	defer C.free(unsafe.Pointer(c_cancel_label))

	C.gtk_file_chooser_native_set_cancel_label((*C.GtkFileChooserNative)(recv.native), c_cancel_label)

	return
}

// Native dialogs are platform dialogs that don't use #GtkDialog or
// #GtkWindow. They are used in order to integrate better with a
// platform, by looking the same as other native applications and
// supporting platform specific features.
//
// The #GtkDialog functions cannot be used on such objects, but we
// need a similar API in order to drive them. The #GtkNativeDialog
// object is an API that allows you to do this. It allows you to set
// various common properties on the dialog, as well as show and hide
// it and get a #GtkNativeDialog::response signal when the user finished
// with the dialog.
//
// There is also a gtk_native_dialog_run() helper that makes it easy
// to run any native dialog in a modal way with a recursive mainloop,
// similar to gtk_dialog_run().
/*

C record/class : GtkNativeDialog
*/
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

// Destroys a dialog.
//
// When a dialog is destroyed, it will break any references it holds
// to other objects. If it is visible it will be hidden and any underlying
// window system resources will be destroyed.
//
// Note that this does not release any reference to the object (as opposed to
// destroying a GtkWindow) because there is no reference from the windowing
// system to the #GtkNativeDialog.
/*

C function : gtk_native_dialog_destroy
*/
func (recv *NativeDialog) Destroy() {
	C.gtk_native_dialog_destroy((*C.GtkNativeDialog)(recv.native))

	return
}

// Returns whether the dialog is modal. See gtk_native_dialog_set_modal().
/*

C function : gtk_native_dialog_get_modal
*/
func (recv *NativeDialog) GetModal() bool {
	retC := C.gtk_native_dialog_get_modal((*C.GtkNativeDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the title of the #GtkNativeDialog.
/*

C function : gtk_native_dialog_get_title
*/
func (recv *NativeDialog) GetTitle() string {
	retC := C.gtk_native_dialog_get_title((*C.GtkNativeDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Fetches the transient parent for this window. See
// gtk_native_dialog_set_transient_for().
/*

C function : gtk_native_dialog_get_transient_for
*/
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

// Determines whether the dialog is visible.
/*

C function : gtk_native_dialog_get_visible
*/
func (recv *NativeDialog) GetVisible() bool {
	retC := C.gtk_native_dialog_get_visible((*C.GtkNativeDialog)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Hides the dialog if it is visilbe, aborting any interaction. Once this
// is called the  #GtkNativeDialog::response signal will not be emitted
// until after the next call to gtk_native_dialog_show().
//
// If the dialog is not visible this does nothing.
/*

C function : gtk_native_dialog_hide
*/
func (recv *NativeDialog) Hide() {
	C.gtk_native_dialog_hide((*C.GtkNativeDialog)(recv.native))

	return
}

// Blocks in a recursive main loop until @self emits the
// #GtkNativeDialog::response signal. It then returns the response ID
// from the ::response signal emission.
//
// Before entering the recursive main loop, gtk_native_dialog_run()
// calls gtk_native_dialog_show() on the dialog for you.
//
// After gtk_native_dialog_run() returns, then dialog will be hidden.
//
// Typical usage of this function might be:
// |[<!-- language="C" -->
// gint result = gtk_native_dialog_run (GTK_NATIVE_DIALOG (dialog));
// switch (result)
// {
// case GTK_RESPONSE_ACCEPT:
// do_application_specific_something ();
// break;
// default:
// do_nothing_since_dialog_was_cancelled ();
// break;
// }
// g_object_unref (dialog);
// ]|
//
// Note that even though the recursive main loop gives the effect of a
// modal dialog (it prevents the user from interacting with other
// windows in the same window group while the dialog is run), callbacks
// such as timeouts, IO channel watches, DND drops, etc, will
// be triggered during a gtk_nautilus_dialog_run() call.
/*

C function : gtk_native_dialog_run
*/
func (recv *NativeDialog) Run() int32 {
	retC := C.gtk_native_dialog_run((*C.GtkNativeDialog)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Sets a dialog modal or non-modal. Modal dialogs prevent interaction
// with other windows in the same application. To keep modal dialogs
// on top of main application windows, use
// gtk_native_dialog_set_transient_for() to make the dialog transient for the
// parent; most [window managers][gtk-X11-arch]
// will then disallow lowering the dialog below the parent.
/*

C function : gtk_native_dialog_set_modal
*/
func (recv *NativeDialog) SetModal(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gtk_native_dialog_set_modal((*C.GtkNativeDialog)(recv.native), c_modal)

	return
}

// Sets the title of the #GtkNativeDialog.
/*

C function : gtk_native_dialog_set_title
*/
func (recv *NativeDialog) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_native_dialog_set_title((*C.GtkNativeDialog)(recv.native), c_title)

	return
}

// Dialog windows should be set transient for the main application
// window they were spawned from. This allows
// [window managers][gtk-X11-arch] to e.g. keep the
// dialog on top of the main window, or center the dialog over the
// main window.
//
// Passing %NULL for @parent unsets the current transient window.
/*

C function : gtk_native_dialog_set_transient_for
*/
func (recv *NativeDialog) SetTransientFor(parent *Window) {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	C.gtk_native_dialog_set_transient_for((*C.GtkNativeDialog)(recv.native), c_parent)

	return
}

// Shows the dialog on the display, allowing the user to interact with
// it. When the user accepts the state of the dialog the dialog will
// be automatically hidden and the #GtkNativeDialog::response signal
// will be emitted.
//
// Multiple calls while the dialog is visible will be ignored.
/*

C function : gtk_native_dialog_show
*/
func (recv *NativeDialog) Show() {
	C.gtk_native_dialog_show((*C.GtkNativeDialog)(recv.native))

	return
}

// #GtkPadController is an event controller for the pads found in drawing
// tablets (The collection of buttons and tactile sensors often found around
// the stylus-sensitive area).
//
// These buttons and sensors have no implicit meaning, and by default they
// perform no action, this event controller is provided to map those to
// #GAction objects, thus letting the application give those a more semantic
// meaning.
//
// Buttons and sensors are not constrained to triggering a single action, some
// %GDK_SOURCE_TABLET_PAD devices feature multiple "modes", all these input
// elements have one current mode, which may determine the final action
// being triggered. Pad devices often divide buttons and sensors into groups,
// all elements in a group share the same current mode, but different groups
// may have different modes. See gdk_device_pad_get_n_groups() and
// gdk_device_pad_get_group_n_modes().
//
// Each of the actions that a given button/strip/ring performs for a given
// mode is defined by #GtkPadActionEntry, it contains an action name that
// will be looked up in the given #GActionGroup and activated whenever the
// specified input element and mode are triggered.
//
// A simple example of #GtkPadController usage, assigning button 1 in all
// modes and pad devices to an "invert-selection" action:
// |[
// GtkPadActionEntry *pad_actions[] = {
// { GTK_PAD_ACTION_BUTTON, 1, -1, "Invert selection", "pad-actions.invert-selection" },
// …
// };
//
// …
// action_group = g_simple_action_group_new ();
// action = g_simple_action_new ("pad-actions.invert-selection", NULL);
// g_signal_connect (action, "activate", on_invert_selection_activated, NULL);
// g_action_map_add_action (G_ACTION_MAP (action_group), action);
// …
// pad_controller = gtk_pad_controller_new (window, action_group, NULL);
// ]|
//
// The actions belonging to rings/strips will be activated with a parameter
// of type %G_VARIANT_TYPE_DOUBLE bearing the value of the given axis, it
// is required that those are made stateful and accepting this #GVariantType.
/*

C record/class : GtkPadController
*/
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

// Returns the constraint for placing this popover.
// See gtk_popover_set_constrain_to().
/*

C function : gtk_popover_get_constrain_to
*/
func (recv *Popover) GetConstrainTo() PopoverConstraint {
	retC := C.gtk_popover_get_constrain_to((*C.GtkPopover)(recv.native))
	retGo := (PopoverConstraint)(retC)

	return retGo
}

// Sets a constraint for positioning this popover.
//
// Note that not all platforms support placing popovers freely,
// and may already impose constraints.
/*

C function : gtk_popover_set_constrain_to
*/
func (recv *Popover) SetConstrainTo(constraint PopoverConstraint) {
	c_constraint := (C.GtkPopoverConstraint)(constraint)

	C.gtk_popover_set_constrain_to((*C.GtkPopover)(recv.native), c_constraint)

	return
}

// Undoes the effect of calling g_object_set() to install an
// application-specific value for a setting. After this call,
// the setting will again follow the session-wide value for
// this setting.
/*

C function : gtk_settings_reset_property
*/
func (recv *Settings) ResetProperty(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_settings_reset_property((*C.GtkSettings)(recv.native), c_name)

	return
}

// #GtkShortcutLabel is a widget that represents a single keyboard shortcut or gesture
// in the user interface.
/*

C record/class : GtkShortcutLabel
*/
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

// A GtkShortcutsGroup represents a group of related keyboard shortcuts
// or gestures. The group has a title. It may optionally be associated with
// a view of the application, which can be used to show only relevant shortcuts
// depending on the application context.
//
// This widget is only meant to be used with #GtkShortcutsWindow.
/*

C record/class : GtkShortcutsGroup
*/
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

// A GtkShortcutsSection collects all the keyboard shortcuts and gestures
// for a major application mode. If your application needs multiple sections,
// you should give each section a unique #GtkShortcutsSection:section-name and
// a #GtkShortcutsSection:title that can be shown in the section selector of
// the GtkShortcutsWindow.
//
// The #GtkShortcutsSection:max-height property can be used to influence how
// the groups in the section are distributed over pages and columns.
//
// This widget is only meant to be used with #GtkShortcutsWindow.
/*

C record/class : GtkShortcutsSection
*/
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

// A GtkShortcutsShortcut represents a single keyboard shortcut or gesture
// with a short text. This widget is only meant to be used with #GtkShortcutsWindow.
/*

C record/class : GtkShortcutsShortcut
*/
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

// A GtkShortcutsWindow shows brief information about the keyboard shortcuts
// and gestures of an application. The shortcuts can be grouped, and you can
// have multiple sections in this window, corresponding to the major modes of
// your application.
//
// Additionally, the shortcuts can be filtered by the current view, to avoid
// showing information that is not relevant in the current application context.
//
// The recommended way to construct a GtkShortcutsWindow is with GtkBuilder,
// by populating a #GtkShortcutsWindow with one or more #GtkShortcutsSection
// objects, which contain #GtkShortcutsGroups that in turn contain objects of
// class #GtkShortcutsShortcut.
//
// # A simple example:
//
// ![](gedit-shortcuts.png)
//
// This example has as single section. As you can see, the shortcut groups
// are arranged in columns, and spread across several pages if there are too
// many to find on a single page.
//
// The .ui file for this example can be found [here](https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-gedit.ui).
//
// # An example with multiple views:
//
// ![](clocks-shortcuts.png)
//
// This example shows a #GtkShortcutsWindow that has been configured to show only
// the shortcuts relevant to the "stopwatch" view.
//
// The .ui file for this example can be found [here](https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-clocks.ui).
//
// # An example with multiple sections:
//
// ![](builder-shortcuts.png)
//
// This example shows a #GtkShortcutsWindow with two sections, "Editor Shortcuts"
// and "Terminal Shortcuts".
//
// The .ui file for this example can be found [here](https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-builder.ui).
/*

C record/class : GtkShortcutsWindow
*/
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

// Converts the style context into a string representation.
//
// The string representation always includes information about
// the name, state, id, visibility and style classes of the CSS
// node that is backing @context. Depending on the flags, more
// information may be included.
//
// This function is intended for testing and debugging of the
// CSS implementation in GTK+. There are no guarantees about
// the format of the returned string, it may change.
/*

C function : gtk_style_context_to_string
*/
func (recv *StyleContext) ToString(flags StyleContextPrintFlags) string {
	c_flags := (C.GtkStyleContextPrintFlags)(flags)

	retC := C.gtk_style_context_to_string((*C.GtkStyleContext)(recv.native), c_flags)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Emits the #GtkTextTagTable::tag-changed signal on the #GtkTextTagTable where
// the tag is included.
//
// The signal is already emitted when setting a #GtkTextTag property. This
// function is useful for a #GtkTextTag subclass.
/*

C function : gtk_text_tag_changed
*/
func (recv *TextTag) Changed(sizeChanged bool) {
	c_size_changed :=
		boolToGboolean(sizeChanged)

	C.gtk_text_tag_changed((*C.GtkTextTag)(recv.native), c_size_changed)

	return
}

// Ensures that the cursor is shown (i.e. not in an 'off' blink
// interval) and resets the time that it will stay blinking (or
// visible, in case blinking is disabled).
//
// This function should be called in response to user input
// (e.g. from derived classes that override the textview's
// #GtkWidget::key-press-event handler).
/*

C function : gtk_text_view_reset_cursor_blink
*/
func (recv *TextView) ResetCursorBlink() {
	C.gtk_text_view_reset_cursor_blink((*C.GtkTextView)(recv.native))

	return
}

// Unsupported : gtk_widget_get_allocated_size : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Returns whether the widget should grab focus when it is clicked with the mouse.
// See gtk_widget_set_focus_on_click().
/*

C function : gtk_widget_get_focus_on_click
*/
func (recv *Widget) GetFocusOnClick() bool {
	retC := C.gtk_widget_get_focus_on_click((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// This function is only for use in widget implementations.
//
// Flags the widget for a rerun of the GtkWidgetClass::size_allocate
// function. Use this function instead of gtk_widget_queue_resize()
// when the @widget's size request didn't change but it wants to
// reposition its contents.
//
// An example user of this function is gtk_widget_set_halign().
/*

C function : gtk_widget_queue_allocate
*/
func (recv *Widget) QueueAllocate() {
	C.gtk_widget_queue_allocate((*C.GtkWidget)(recv.native))

	return
}

// Sets whether the widget should grab focus when it is clicked with the mouse.
// Making mouse clicks not grab focus is useful in places like toolbars where
// you don’t want the keyboard focus removed from the main area of the
// application.
/*

C function : gtk_widget_set_focus_on_click
*/
func (recv *Widget) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_widget_set_focus_on_click((*C.GtkWidget)(recv.native), c_focus_on_click)

	return
}
