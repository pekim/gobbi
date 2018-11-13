// This is a generated file - DO NOT EDIT
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
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

	gboolean combobox_popdownHandler(GObject *, gpointer);

	static gulong ComboBox_signal_connect_popdown(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popdown", G_CALLBACK(combobox_popdownHandler), data);
	}

*/
/*

	void combobox_popupHandler(GObject *, gpointer);

	static gulong ComboBox_signal_connect_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup", G_CALLBACK(combobox_popupHandler), data);
	}

*/
/*

	gboolean entrycompletion_cursorOnMatchHandler(GObject *, GtkTreeModel *, GtkTreeIter *, gpointer);

	static gulong EntryCompletion_signal_connect_cursor_on_match(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cursor-on-match", G_CALLBACK(entrycompletion_cursorOnMatchHandler), data);
	}

*/
/*

	void filechooserbutton_fileSetHandler(GObject *, gpointer);

	static gulong FileChooserButton_signal_connect_file_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "file-set", G_CALLBACK(filechooserbutton_fileSetHandler), data);
	}

*/
/*

	void scalebutton_popdownHandler(GObject *, gpointer);

	static gulong ScaleButton_signal_connect_popdown(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popdown", G_CALLBACK(scalebutton_popdownHandler), data);
	}

*/
/*

	void scalebutton_popupHandler(GObject *, gpointer);

	static gulong ScaleButton_signal_connect_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup", G_CALLBACK(scalebutton_popupHandler), data);
	}

*/
import "C"

// Returns the program name displayed in the about dialog.
/*

C function

gtk_about_dialog_get_program_name
*/
func (recv *AboutDialog) GetProgramName() string {
	retC := C.gtk_about_dialog_get_program_name((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Sets the name to display in the about dialog.
// If this is not set, it defaults to g_get_application_name().
/*

C function

gtk_about_dialog_set_program_name
*/
func (recv *AboutDialog) SetProgramName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_about_dialog_set_program_name((*C.GtkAboutDialog)(recv.native), c_name)

	return
}

// If @action provides a #GtkMenu widget as a submenu for the menu
// item or the toolbar item it creates, this function returns an
// instance of that menu.
/*

C function

gtk_action_create_menu
*/
func (recv *Action) CreateMenu() *Widget {
	retC := C.gtk_action_create_menu((*C.GtkAction)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new empty builder object.
//
// This function is only useful if you intend to make multiple calls
// to gtk_builder_add_from_file(), gtk_builder_add_from_resource()
// or gtk_builder_add_from_string() in order to merge multiple UI
// descriptions into a single builder.
//
// Most users will probably want to use gtk_builder_new_from_file(),
// gtk_builder_new_from_resource() or gtk_builder_new_from_string().
/*

C function

gtk_builder_new
*/
func BuilderNew() *Builder {
	retC := C.gtk_builder_new()
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Parses a file containing a [GtkBuilder UI definition][BUILDER-UI]
// and merges it with the current contents of @builder.
//
// Most users will probably want to use gtk_builder_new_from_file().
//
// If an error occurs, 0 will be returned and @error will be assigned a
// #GError from the #GTK_BUILDER_ERROR, #G_MARKUP_ERROR or #G_FILE_ERROR
// domain.
//
// It’s not really reasonable to attempt to handle failures of this
// call. You should not use this function with untrusted files (ie:
// files that are not part of your application). Broken #GtkBuilder
// files can easily crash your program, and it’s possible that memory
// was leaked leading up to the reported failure. The only reasonable
// thing to do when an error is detected is to call g_error().
/*

C function

gtk_builder_add_from_file
*/
func (recv *Builder) AddFromFile(filename string) (uint32, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_from_file((*C.GtkBuilder)(recv.native), c_filename, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Parses a string containing a [GtkBuilder UI definition][BUILDER-UI]
// and merges it with the current contents of @builder.
//
// Most users will probably want to use gtk_builder_new_from_string().
//
// Upon errors 0 will be returned and @error will be assigned a
// #GError from the #GTK_BUILDER_ERROR, #G_MARKUP_ERROR or
// #G_VARIANT_PARSE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this
// call.  The only reasonable thing to do when an error is detected is
// to call g_error().
/*

C function

gtk_builder_add_from_string
*/
func (recv *Builder) AddFromString(buffer string, length uint64) (uint32, error) {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_length := (C.gsize)(length)

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_from_string((*C.GtkBuilder)(recv.native), c_buffer, c_length, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This method is a simpler variation of gtk_builder_connect_signals_full().
// It uses symbols explicitly added to @builder with prior calls to
// gtk_builder_add_callback_symbol(). In the case that symbols are not
// explicitly added; it uses #GModule’s introspective features (by opening the module %NULL)
// to look at the application’s symbol table. From here it tries to match
// the signal handler names given in the interface description with
// symbols in the application and connects the signals. Note that this
// function can only be called once, subsequent calls will do nothing.
//
// Note that unless gtk_builder_add_callback_symbol() is called for
// all signal callbacks which are referenced by the loaded XML, this
// function will require that #GModule be supported on the platform.
//
// If you rely on #GModule support to lookup callbacks in the symbol table,
// the following details should be noted:
//
// When compiling applications for Windows, you must declare signal callbacks
// with #G_MODULE_EXPORT, or they will not be put in the symbol table.
// On Linux and Unices, this is not necessary; applications should instead
// be compiled with the -Wl,--export-dynamic CFLAGS, and linked against
// gmodule-export-2.0.
/*

C function

gtk_builder_connect_signals
*/
func (recv *Builder) ConnectSignals(userData uintptr) {
	c_user_data := (C.gpointer)(userData)

	C.gtk_builder_connect_signals((*C.GtkBuilder)(recv.native), c_user_data)

	return
}

// Unsupported : gtk_builder_connect_signals_full : unsupported parameter func : no type generator for BuilderConnectFunc (GtkBuilderConnectFunc) for param func

// Gets the object named @name. Note that this function does not
// increment the reference count of the returned object.
/*

C function

gtk_builder_get_object
*/
func (recv *Builder) GetObject(name string) *gobject.Object {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_builder_get_object((*C.GtkBuilder)(recv.native), c_name)
	var retGo (*gobject.Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gobject.ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets all objects that have been constructed by @builder. Note that
// this function does not increment the reference counts of the returned
// objects.
/*

C function

gtk_builder_get_objects
*/
func (recv *Builder) GetObjects() *glib.SList {
	retC := C.gtk_builder_get_objects((*C.GtkBuilder)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the translation domain of @builder.
/*

C function

gtk_builder_get_translation_domain
*/
func (recv *Builder) GetTranslationDomain() string {
	retC := C.gtk_builder_get_translation_domain((*C.GtkBuilder)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Looks up a type by name, using the virtual function that
// #GtkBuilder has for that purpose. This is mainly used when
// implementing the #GtkBuildable interface on a type.
/*

C function

gtk_builder_get_type_from_name
*/
func (recv *Builder) GetTypeFromName(typeName string) gobject.Type {
	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	retC := C.gtk_builder_get_type_from_name((*C.GtkBuilder)(recv.native), c_type_name)
	retGo := (gobject.Type)(retC)

	return retGo
}

// Sets the translation domain of @builder.
// See #GtkBuilder:translation-domain.
/*

C function

gtk_builder_set_translation_domain
*/
func (recv *Builder) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.gtk_builder_set_translation_domain((*C.GtkBuilder)(recv.native), c_domain)

	return
}

// Unsupported : gtk_builder_value_from_string : unsupported parameter pspec : Blacklisted record : GParamSpec

// Like gtk_builder_value_from_string(), this function demarshals
// a value from a string, but takes a #GType instead of #GParamSpec.
// This function calls g_value_init() on the @value argument, so it
// need not be initialised beforehand.
//
// Upon errors %FALSE will be returned and @error will be assigned a
// #GError from the #GTK_BUILDER_ERROR domain.
/*

C function

gtk_builder_value_from_string_type
*/
func (recv *Builder) ValueFromStringType(type_ gobject.Type, string string) (bool, *gobject.Value, error) {
	c_type := (C.GType)(type_)

	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	var c_value C.GValue

	var cThrowableError *C.GError

	retC := C.gtk_builder_value_from_string_type((*C.GtkBuilder)(recv.native), c_type, c_string, &c_value, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return retGo, value, goThrowableError
}

// Unsupported signal 'move-active' for ComboBox : unsupported parameter scroll_type : type ScrollType :

type signalComboBoxPopdownDetail struct {
	callback  ComboBoxSignalPopdownCallback
	handlerID C.gulong
}

var signalComboBoxPopdownId int
var signalComboBoxPopdownMap = make(map[int]signalComboBoxPopdownDetail)
var signalComboBoxPopdownLock sync.Mutex

// ComboBoxSignalPopdownCallback is a callback function for a 'popdown' signal emitted from a ComboBox.
type ComboBoxSignalPopdownCallback func() bool

/*
ConnectPopdown connects the callback to the 'popdown' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectPopdown to remove it.
*/
func (recv *ComboBox) ConnectPopdown(callback ComboBoxSignalPopdownCallback) int {
	signalComboBoxPopdownLock.Lock()
	defer signalComboBoxPopdownLock.Unlock()

	signalComboBoxPopdownId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_popdown(instance, C.gpointer(uintptr(signalComboBoxPopdownId)))

	detail := signalComboBoxPopdownDetail{callback, handlerID}
	signalComboBoxPopdownMap[signalComboBoxPopdownId] = detail

	return signalComboBoxPopdownId
}

/*
DisconnectPopdown disconnects a callback from the 'popdown' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectPopdown.
*/
func (recv *ComboBox) DisconnectPopdown(connectionID int) {
	signalComboBoxPopdownLock.Lock()
	defer signalComboBoxPopdownLock.Unlock()

	detail, exists := signalComboBoxPopdownMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxPopdownMap, connectionID)
}

//export combobox_popdownHandler
func combobox_popdownHandler(_ *C.GObject, data C.gpointer) C.gboolean {
	index := int(uintptr(data))
	callback := signalComboBoxPopdownMap[index].callback
	retGo := callback()
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalComboBoxPopupDetail struct {
	callback  ComboBoxSignalPopupCallback
	handlerID C.gulong
}

var signalComboBoxPopupId int
var signalComboBoxPopupMap = make(map[int]signalComboBoxPopupDetail)
var signalComboBoxPopupLock sync.Mutex

// ComboBoxSignalPopupCallback is a callback function for a 'popup' signal emitted from a ComboBox.
type ComboBoxSignalPopupCallback func()

/*
ConnectPopup connects the callback to the 'popup' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectPopup to remove it.
*/
func (recv *ComboBox) ConnectPopup(callback ComboBoxSignalPopupCallback) int {
	signalComboBoxPopupLock.Lock()
	defer signalComboBoxPopupLock.Unlock()

	signalComboBoxPopupId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_popup(instance, C.gpointer(uintptr(signalComboBoxPopupId)))

	detail := signalComboBoxPopupDetail{callback, handlerID}
	signalComboBoxPopupMap[signalComboBoxPopupId] = detail

	return signalComboBoxPopupId
}

/*
DisconnectPopup disconnects a callback from the 'popup' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectPopup.
*/
func (recv *ComboBox) DisconnectPopup(connectionID int) {
	signalComboBoxPopupLock.Lock()
	defer signalComboBoxPopupLock.Unlock()

	detail, exists := signalComboBoxPopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxPopupMap, connectionID)
}

//export combobox_popupHandler
func combobox_popupHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalComboBoxPopupMap[index].callback
	callback()
}

// Retrieves the horizontal cursor adjustment for the entry.
// See gtk_entry_set_cursor_hadjustment().
/*

C function

gtk_entry_get_cursor_hadjustment
*/
func (recv *Entry) GetCursorHadjustment() *Adjustment {
	retC := C.gtk_entry_get_cursor_hadjustment((*C.GtkEntry)(recv.native))
	var retGo (*Adjustment)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AdjustmentNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Hooks up an adjustment to the cursor position in an entry, so that when
// the cursor is moved, the adjustment is scrolled to show that position.
// See gtk_scrolled_window_get_hadjustment() for a typical way of obtaining
// the adjustment.
//
// The adjustment has to be in pixel units and in the same coordinate system
// as the entry.
/*

C function

gtk_entry_set_cursor_hadjustment
*/
func (recv *Entry) SetCursorHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	C.gtk_entry_set_cursor_hadjustment((*C.GtkEntry)(recv.native), c_adjustment)

	return
}

type signalEntryCompletionCursorOnMatchDetail struct {
	callback  EntryCompletionSignalCursorOnMatchCallback
	handlerID C.gulong
}

var signalEntryCompletionCursorOnMatchId int
var signalEntryCompletionCursorOnMatchMap = make(map[int]signalEntryCompletionCursorOnMatchDetail)
var signalEntryCompletionCursorOnMatchLock sync.Mutex

// EntryCompletionSignalCursorOnMatchCallback is a callback function for a 'cursor-on-match' signal emitted from a EntryCompletion.
type EntryCompletionSignalCursorOnMatchCallback func(model *TreeModel, iter *TreeIter) bool

/*
ConnectCursorOnMatch connects the callback to the 'cursor-on-match' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectCursorOnMatch to remove it.
*/
func (recv *EntryCompletion) ConnectCursorOnMatch(callback EntryCompletionSignalCursorOnMatchCallback) int {
	signalEntryCompletionCursorOnMatchLock.Lock()
	defer signalEntryCompletionCursorOnMatchLock.Unlock()

	signalEntryCompletionCursorOnMatchId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_cursor_on_match(instance, C.gpointer(uintptr(signalEntryCompletionCursorOnMatchId)))

	detail := signalEntryCompletionCursorOnMatchDetail{callback, handlerID}
	signalEntryCompletionCursorOnMatchMap[signalEntryCompletionCursorOnMatchId] = detail

	return signalEntryCompletionCursorOnMatchId
}

/*
DisconnectCursorOnMatch disconnects a callback from the 'cursor-on-match' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectCursorOnMatch.
*/
func (recv *EntryCompletion) DisconnectCursorOnMatch(connectionID int) {
	signalEntryCompletionCursorOnMatchLock.Lock()
	defer signalEntryCompletionCursorOnMatchLock.Unlock()

	detail, exists := signalEntryCompletionCursorOnMatchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionCursorOnMatchMap, connectionID)
}

//export entrycompletion_cursorOnMatchHandler
func entrycompletion_cursorOnMatchHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, data C.gpointer) C.gboolean {
	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalEntryCompletionCursorOnMatchMap[index].callback
	retGo := callback(model, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Get the original text entered by the user that triggered
// the completion or %NULL if there’s no completion ongoing.
/*

C function

gtk_entry_completion_get_completion_prefix
*/
func (recv *EntryCompletion) GetCompletionPrefix() string {
	retC := C.gtk_entry_completion_get_completion_prefix((*C.GtkEntryCompletion)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Returns %TRUE if inline-selection mode is turned on.
/*

C function

gtk_entry_completion_get_inline_selection
*/
func (recv *EntryCompletion) GetInlineSelection() bool {
	retC := C.gtk_entry_completion_get_inline_selection((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether it is possible to cycle through the possible completions
// inside the entry.
/*

C function

gtk_entry_completion_set_inline_selection
*/
func (recv *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	c_inline_selection :=
		boolToGboolean(inlineSelection)

	C.gtk_entry_completion_set_inline_selection((*C.GtkEntryCompletion)(recv.native), c_inline_selection)

	return
}

type signalFileChooserButtonFileSetDetail struct {
	callback  FileChooserButtonSignalFileSetCallback
	handlerID C.gulong
}

var signalFileChooserButtonFileSetId int
var signalFileChooserButtonFileSetMap = make(map[int]signalFileChooserButtonFileSetDetail)
var signalFileChooserButtonFileSetLock sync.Mutex

// FileChooserButtonSignalFileSetCallback is a callback function for a 'file-set' signal emitted from a FileChooserButton.
type FileChooserButtonSignalFileSetCallback func()

/*
ConnectFileSet connects the callback to the 'file-set' signal for the FileChooserButton.

The returned value represents the connection, and may be passed to DisconnectFileSet to remove it.
*/
func (recv *FileChooserButton) ConnectFileSet(callback FileChooserButtonSignalFileSetCallback) int {
	signalFileChooserButtonFileSetLock.Lock()
	defer signalFileChooserButtonFileSetLock.Unlock()

	signalFileChooserButtonFileSetId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooserButton_signal_connect_file_set(instance, C.gpointer(uintptr(signalFileChooserButtonFileSetId)))

	detail := signalFileChooserButtonFileSetDetail{callback, handlerID}
	signalFileChooserButtonFileSetMap[signalFileChooserButtonFileSetId] = detail

	return signalFileChooserButtonFileSetId
}

/*
DisconnectFileSet disconnects a callback from the 'file-set' signal for the FileChooserButton.

The connectionID should be a value returned from a call to ConnectFileSet.
*/
func (recv *FileChooserButton) DisconnectFileSet(connectionID int) {
	signalFileChooserButtonFileSetLock.Lock()
	defer signalFileChooserButtonFileSetLock.Unlock()

	detail, exists := signalFileChooserButtonFileSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserButtonFileSetMap, connectionID)
}

//export filechooserbutton_fileSetHandler
func filechooserbutton_fileSetHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalFileChooserButtonFileSetMap[index].callback
	callback()
}

// Unsupported : gtk_icon_theme_choose_icon : unsupported parameter icon_names :

// Gets the list of contexts available within the current
// hierarchy of icon themes.
// See gtk_icon_theme_list_icons() for details about contexts.
/*

C function

gtk_icon_theme_list_contexts
*/
func (recv *IconTheme) ListContexts() *glib.List {
	retC := C.gtk_icon_theme_list_contexts((*C.GtkIconTheme)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Converts widget coordinates to coordinates for the bin_window,
// as expected by e.g. gtk_icon_view_get_path_at_pos().
/*

C function

gtk_icon_view_convert_widget_to_bin_window_coords
*/
func (recv *IconView) ConvertWidgetToBinWindowCoords(wx int32, wy int32) (int32, int32) {
	c_wx := (C.gint)(wx)

	c_wy := (C.gint)(wy)

	var c_bx C.gint

	var c_by C.gint

	C.gtk_icon_view_convert_widget_to_bin_window_coords((*C.GtkIconView)(recv.native), c_wx, c_wy, &c_bx, &c_by)

	bx := (int32)(c_bx)

	by := (int32)(c_by)

	return bx, by
}

// Returns the column of @icon_view’s model which is being used for
// displaying tooltips on @icon_view’s rows.
/*

C function

gtk_icon_view_get_tooltip_column
*/
func (recv *IconView) GetTooltipColumn() int32 {
	retC := C.gtk_icon_view_get_tooltip_column((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// This function is supposed to be used in a #GtkWidget::query-tooltip
// signal handler for #GtkIconView.  The @x, @y and @keyboard_tip values
// which are received in the signal handler, should be passed to this
// function without modification.
//
// The return value indicates whether there is an icon view item at the given
// coordinates (%TRUE) or not (%FALSE) for mouse tooltips. For keyboard
// tooltips the item returned will be the cursor item. When %TRUE, then any of
// @model, @path and @iter which have been provided will be set to point to
// that row and the corresponding model. @x and @y will always be converted
// to be relative to @icon_view’s bin_window if @keyboard_tooltip is %FALSE.
/*

C function

gtk_icon_view_get_tooltip_context
*/
func (recv *IconView) GetTooltipContext(x int32, y int32, keyboardTip bool) (bool, *TreeModel, *TreePath, *TreeIter) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_keyboard_tip :=
		boolToGboolean(keyboardTip)

	var c_model *C.GtkTreeModel

	var c_path *C.GtkTreePath

	var c_iter C.GtkTreeIter

	retC := C.gtk_icon_view_get_tooltip_context((*C.GtkIconView)(recv.native), &c_x, &c_y, c_keyboard_tip, &c_model, &c_path, &c_iter)
	retGo := retC == C.TRUE

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, model, path, iter
}

// Sets the tip area of @tooltip to the area which @cell occupies in
// the item pointed to by @path. See also gtk_tooltip_set_tip_area().
//
// See also gtk_icon_view_set_tooltip_column() for a simpler alternative.
/*

C function

gtk_icon_view_set_tooltip_cell
*/
func (recv *IconView) SetTooltipCell(tooltip *Tooltip, path *TreePath, cell *CellRenderer) {
	c_tooltip := (*C.GtkTooltip)(C.NULL)
	if tooltip != nil {
		c_tooltip = (*C.GtkTooltip)(tooltip.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	C.gtk_icon_view_set_tooltip_cell((*C.GtkIconView)(recv.native), c_tooltip, c_path, c_cell)

	return
}

// If you only plan to have simple (text-only) tooltips on full items, you
// can use this function to have #GtkIconView handle these automatically
// for you. @column should be set to the column in @icon_view’s model
// containing the tooltip texts, or -1 to disable this feature.
//
// When enabled, #GtkWidget:has-tooltip will be set to %TRUE and
// @icon_view will connect a #GtkWidget::query-tooltip signal handler.
//
// Note that the signal handler sets the text with gtk_tooltip_set_markup(),
// so &, <, etc have to be escaped in the text.
/*

C function

gtk_icon_view_set_tooltip_column
*/
func (recv *IconView) SetTooltipColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_icon_view_set_tooltip_column((*C.GtkIconView)(recv.native), c_column)

	return
}

// Sets the tip area of @tooltip to be the area covered by the item at @path.
// See also gtk_icon_view_set_tooltip_column() for a simpler alternative.
// See also gtk_tooltip_set_tip_area().
/*

C function

gtk_icon_view_set_tooltip_item
*/
func (recv *IconView) SetTooltipItem(tooltip *Tooltip, path *TreePath) {
	c_tooltip := (*C.GtkTooltip)(C.NULL)
	if tooltip != nil {
		c_tooltip = (*C.GtkTooltip)(tooltip.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_icon_view_set_tooltip_item((*C.GtkIconView)(recv.native), c_tooltip, c_path)

	return
}

// Unsupported : gtk_list_store_set_valuesv : unsupported parameter values :

// Unsupported signal 'move-selected' for MenuShell : unsupported parameter distance : type gint :

// Sets the tooltip markup text to be used as tooltip for the arrow button
// which pops up the menu.  See gtk_tool_item_set_tooltip_text() for setting
// a tooltip on the whole #GtkMenuToolButton.
/*

C function

gtk_menu_tool_button_set_arrow_tooltip_markup
*/
func (recv *MenuToolButton) SetArrowTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_menu_tool_button_set_arrow_tooltip_markup((*C.GtkMenuToolButton)(recv.native), c_markup)

	return
}

// Sets the tooltip text to be used as tooltip for the arrow button which
// pops up the menu.  See gtk_tool_item_set_tooltip_text() for setting a tooltip
// on the whole #GtkMenuToolButton.
/*

C function

gtk_menu_tool_button_set_arrow_tooltip_text
*/
func (recv *MenuToolButton) SetArrowTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_menu_tool_button_set_arrow_tooltip_text((*C.GtkMenuToolButton)(recv.native), c_text)

	return
}

// Unsupported signal 'create-window' for Notebook : unsupported parameter x : type gint :

// Reads the page setup from the file @file_name. Returns a
// new #GtkPageSetup object with the restored page setup,
// or %NULL if an error occurred. See gtk_page_setup_to_file().
/*

C function

gtk_page_setup_new_from_file
*/
func PageSetupNewFromFile(fileName string) (*PageSetup, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_new_from_file(c_file_name, &cThrowableError)
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads the page setup from the group @group_name in the key file
// @key_file. Returns a new #GtkPageSetup object with the restored
// page setup, or %NULL if an error occurred.
/*

C function

gtk_page_setup_new_from_key_file
*/
func PageSetupNewFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PageSetup, error) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_new_from_key_file(c_key_file, c_group_name, &cThrowableError)
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This function saves the information from @setup to @file_name.
/*

C function

gtk_page_setup_to_file
*/
func (recv *PageSetup) ToFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_to_file((*C.GtkPageSetup)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This function adds the page setup from @setup to @key_file.
/*

C function

gtk_page_setup_to_key_file
*/
func (recv *PageSetup) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_page_setup_to_key_file((*C.GtkPageSetup)(recv.native), c_key_file, c_group_name)

	return
}

// Reads the print settings from @file_name. Returns a new #GtkPrintSettings
// object with the restored settings, or %NULL if an error occurred. If the
// file could not be loaded then error is set to either a #GFileError or
// #GKeyFileError.  See gtk_print_settings_to_file().
/*

C function

gtk_print_settings_new_from_file
*/
func PrintSettingsNewFromFile(fileName string) (*PrintSettings, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_new_from_file(c_file_name, &cThrowableError)
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reads the print settings from the group @group_name in @key_file.  Returns a
// new #GtkPrintSettings object with the restored settings, or %NULL if an
// error occurred. If the file could not be loaded then error is set to either
// a #GFileError or #GKeyFileError.
/*

C function

gtk_print_settings_new_from_key_file
*/
func PrintSettingsNewFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PrintSettings, error) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_new_from_key_file(c_key_file, c_group_name, &cThrowableError)
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This function saves the print settings from @settings to @file_name. If the
// file could not be loaded then error is set to either a #GFileError or
// #GKeyFileError.
/*

C function

gtk_print_settings_to_file
*/
func (recv *PrintSettings) ToFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_to_file((*C.GtkPrintSettings)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// This function adds the print settings from @settings to @key_file.
/*

C function

gtk_print_settings_to_key_file
*/
func (recv *PrintSettings) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_print_settings_to_key_file((*C.GtkPrintSettings)(recv.native), c_key_file, c_group_name)

	return
}

// Gets the current position of the fill level indicator.
/*

C function

gtk_range_get_fill_level
*/
func (recv *Range) GetFillLevel() float64 {
	retC := C.gtk_range_get_fill_level((*C.GtkRange)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Gets whether the range is restricted to the fill level.
/*

C function

gtk_range_get_restrict_to_fill_level
*/
func (recv *Range) GetRestrictToFillLevel() bool {
	retC := C.gtk_range_get_restrict_to_fill_level((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets whether the range displays the fill level graphically.
/*

C function

gtk_range_get_show_fill_level
*/
func (recv *Range) GetShowFillLevel() bool {
	retC := C.gtk_range_get_show_fill_level((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Set the new position of the fill level indicator.
//
// The “fill level” is probably best described by its most prominent
// use case, which is an indicator for the amount of pre-buffering in
// a streaming media player. In that use case, the value of the range
// would indicate the current play position, and the fill level would
// be the position up to which the file/stream has been downloaded.
//
// This amount of prebuffering can be displayed on the range’s trough
// and is themeable separately from the trough. To enable fill level
// display, use gtk_range_set_show_fill_level(). The range defaults
// to not showing the fill level.
//
// Additionally, it’s possible to restrict the range’s slider position
// to values which are smaller than the fill level. This is controller
// by gtk_range_set_restrict_to_fill_level() and is by default
// enabled.
/*

C function

gtk_range_set_fill_level
*/
func (recv *Range) SetFillLevel(fillLevel float64) {
	c_fill_level := (C.gdouble)(fillLevel)

	C.gtk_range_set_fill_level((*C.GtkRange)(recv.native), c_fill_level)

	return
}

// Sets whether the slider is restricted to the fill level. See
// gtk_range_set_fill_level() for a general description of the fill
// level concept.
/*

C function

gtk_range_set_restrict_to_fill_level
*/
func (recv *Range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	c_restrict_to_fill_level :=
		boolToGboolean(restrictToFillLevel)

	C.gtk_range_set_restrict_to_fill_level((*C.GtkRange)(recv.native), c_restrict_to_fill_level)

	return
}

// Sets whether a graphical fill level is show on the trough. See
// gtk_range_set_fill_level() for a general description of the fill
// level concept.
/*

C function

gtk_range_set_show_fill_level
*/
func (recv *Range) SetShowFillLevel(showFillLevel bool) {
	c_show_fill_level :=
		boolToGboolean(showFillLevel)

	C.gtk_range_set_show_fill_level((*C.GtkRange)(recv.native), c_show_fill_level)

	return
}

// Creates a new #GtkRecentAction object. To add the action to
// a #GtkActionGroup and set the accelerator for the action,
// call gtk_action_group_add_action_with_accel().
/*

C function

gtk_recent_action_new
*/
func RecentActionNew(name string, label string, tooltip string, stockId string) *RecentAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_recent_action_new(c_name, c_label, c_tooltip, c_stock_id)
	retGo := RecentActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkRecentAction object. To add the action to
// a #GtkActionGroup and set the accelerator for the action,
// call gtk_action_group_add_action_with_accel().
/*

C function

gtk_recent_action_new_for_manager
*/
func RecentActionNewForManager(name string, label string, tooltip string, stockId string, manager *RecentManager) *RecentAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_manager := (*C.GtkRecentManager)(C.NULL)
	if manager != nil {
		c_manager = (*C.GtkRecentManager)(manager.ToC())
	}

	retC := C.gtk_recent_action_new_for_manager(c_name, c_label, c_tooltip, c_stock_id, c_manager)
	retGo := RecentActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the value set by gtk_recent_chooser_menu_set_show_numbers().
/*

C function

gtk_recent_action_get_show_numbers
*/
func (recv *RecentAction) GetShowNumbers() bool {
	retC := C.gtk_recent_action_get_show_numbers((*C.GtkRecentAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets whether a number should be added to the items shown by the
// widgets representing @action. The numbers are shown to provide
// a unique character for a mnemonic to be used inside the menu item's
// label. Only the first ten items get a number to avoid clashes.
/*

C function

gtk_recent_action_set_show_numbers
*/
func (recv *RecentAction) SetShowNumbers(showNumbers bool) {
	c_show_numbers :=
		boolToGboolean(showNumbers)

	C.gtk_recent_action_set_show_numbers((*C.GtkRecentAction)(recv.native), c_show_numbers)

	return
}

type signalScaleButtonPopdownDetail struct {
	callback  ScaleButtonSignalPopdownCallback
	handlerID C.gulong
}

var signalScaleButtonPopdownId int
var signalScaleButtonPopdownMap = make(map[int]signalScaleButtonPopdownDetail)
var signalScaleButtonPopdownLock sync.Mutex

// ScaleButtonSignalPopdownCallback is a callback function for a 'popdown' signal emitted from a ScaleButton.
type ScaleButtonSignalPopdownCallback func()

/*
ConnectPopdown connects the callback to the 'popdown' signal for the ScaleButton.

The returned value represents the connection, and may be passed to DisconnectPopdown to remove it.
*/
func (recv *ScaleButton) ConnectPopdown(callback ScaleButtonSignalPopdownCallback) int {
	signalScaleButtonPopdownLock.Lock()
	defer signalScaleButtonPopdownLock.Unlock()

	signalScaleButtonPopdownId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScaleButton_signal_connect_popdown(instance, C.gpointer(uintptr(signalScaleButtonPopdownId)))

	detail := signalScaleButtonPopdownDetail{callback, handlerID}
	signalScaleButtonPopdownMap[signalScaleButtonPopdownId] = detail

	return signalScaleButtonPopdownId
}

/*
DisconnectPopdown disconnects a callback from the 'popdown' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectPopdown.
*/
func (recv *ScaleButton) DisconnectPopdown(connectionID int) {
	signalScaleButtonPopdownLock.Lock()
	defer signalScaleButtonPopdownLock.Unlock()

	detail, exists := signalScaleButtonPopdownMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScaleButtonPopdownMap, connectionID)
}

//export scalebutton_popdownHandler
func scalebutton_popdownHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalScaleButtonPopdownMap[index].callback
	callback()
}

type signalScaleButtonPopupDetail struct {
	callback  ScaleButtonSignalPopupCallback
	handlerID C.gulong
}

var signalScaleButtonPopupId int
var signalScaleButtonPopupMap = make(map[int]signalScaleButtonPopupDetail)
var signalScaleButtonPopupLock sync.Mutex

// ScaleButtonSignalPopupCallback is a callback function for a 'popup' signal emitted from a ScaleButton.
type ScaleButtonSignalPopupCallback func()

/*
ConnectPopup connects the callback to the 'popup' signal for the ScaleButton.

The returned value represents the connection, and may be passed to DisconnectPopup to remove it.
*/
func (recv *ScaleButton) ConnectPopup(callback ScaleButtonSignalPopupCallback) int {
	signalScaleButtonPopupLock.Lock()
	defer signalScaleButtonPopupLock.Unlock()

	signalScaleButtonPopupId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScaleButton_signal_connect_popup(instance, C.gpointer(uintptr(signalScaleButtonPopupId)))

	detail := signalScaleButtonPopupDetail{callback, handlerID}
	signalScaleButtonPopupMap[signalScaleButtonPopupId] = detail

	return signalScaleButtonPopupId
}

/*
DisconnectPopup disconnects a callback from the 'popup' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectPopup.
*/
func (recv *ScaleButton) DisconnectPopup(connectionID int) {
	signalScaleButtonPopupLock.Lock()
	defer signalScaleButtonPopupLock.Unlock()

	detail, exists := signalScaleButtonPopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScaleButtonPopupMap, connectionID)
}

//export scalebutton_popupHandler
func scalebutton_popupHandler(_ *C.GObject, data C.gpointer) {
	index := int(uintptr(data))
	callback := signalScaleButtonPopupMap[index].callback
	callback()
}

// Unsupported signal 'value-changed' for ScaleButton : unsupported parameter value : type gdouble :

// Unsupported : gtk_scale_button_new : unsupported parameter icons :

// Gets the #GtkAdjustment associated with the #GtkScaleButton’s scale.
// See gtk_range_get_adjustment() for details.
/*

C function

gtk_scale_button_get_adjustment
*/
func (recv *ScaleButton) GetAdjustment() *Adjustment {
	retC := C.gtk_scale_button_get_adjustment((*C.GtkScaleButton)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the current value of the scale button.
/*

C function

gtk_scale_button_get_value
*/
func (recv *ScaleButton) GetValue() float64 {
	retC := C.gtk_scale_button_get_value((*C.GtkScaleButton)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Sets the #GtkAdjustment to be used as a model
// for the #GtkScaleButton’s scale.
// See gtk_range_set_adjustment() for details.
/*

C function

gtk_scale_button_set_adjustment
*/
func (recv *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	C.gtk_scale_button_set_adjustment((*C.GtkScaleButton)(recv.native), c_adjustment)

	return
}

// Unsupported : gtk_scale_button_set_icons : unsupported parameter icons :

// Sets the current value of the scale; if the value is outside
// the minimum or maximum range values, it will be clamped to fit
// inside them. The scale button emits the #GtkScaleButton::value-changed
// signal if the value changes.
/*

C function

gtk_scale_button_set_value
*/
func (recv *ScaleButton) SetValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_scale_button_set_value((*C.GtkScaleButton)(recv.native), c_value)

	return
}

// Returns the #GdkScreen associated with @status_icon.
/*

C function

gtk_status_icon_get_screen
*/
func (recv *StatusIcon) GetScreen() *gdk.Screen {
	retC := C.gtk_status_icon_get_screen((*C.GtkStatusIcon)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the #GdkScreen where @status_icon is displayed; if
// the icon is already mapped, it will be unmapped, and
// then remapped on the new screen.
/*

C function

gtk_status_icon_set_screen
*/
func (recv *StatusIcon) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_status_icon_set_screen((*C.GtkStatusIcon)(recv.native), c_screen)

	return
}

// Adds the mark at position @where. The mark must not be added to
// another buffer, and if its name is not %NULL then there must not
// be another mark in the buffer with the same name.
//
// Emits the #GtkTextBuffer::mark-set signal as notification of the mark's
// initial placement.
/*

C function

gtk_text_buffer_add_mark
*/
func (recv *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	c_mark := (*C.GtkTextMark)(C.NULL)
	if mark != nil {
		c_mark = (*C.GtkTextMark)(mark.ToC())
	}

	c_where := (*C.GtkTextIter)(C.NULL)
	if where != nil {
		c_where = (*C.GtkTextIter)(where.ToC())
	}

	C.gtk_text_buffer_add_mark((*C.GtkTextBuffer)(recv.native), c_mark, c_where)

	return
}

// Creates a text mark. Add it to a buffer using gtk_text_buffer_add_mark().
// If @name is %NULL, the mark is anonymous; otherwise, the mark can be
// retrieved by name using gtk_text_buffer_get_mark(). If a mark has left
// gravity, and text is inserted at the mark’s current location, the mark
// will be moved to the left of the newly-inserted text. If the mark has
// right gravity (@left_gravity = %FALSE), the mark will end up on the
// right of newly-inserted text. The standard left-to-right cursor is a
// mark with right gravity (when you type, the cursor stays on the right
// side of the text you’re typing).
/*

C function

gtk_text_mark_new
*/
func TextMarkNew(name string, leftGravity bool) *TextMark {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_left_gravity :=
		boolToGboolean(leftGravity)

	retC := C.gtk_text_mark_new(c_name, c_left_gravity)
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets the markup text to be displayed as tooltip on the item.
// See gtk_widget_set_tooltip_markup().
/*

C function

gtk_tool_item_set_tooltip_markup
*/
func (recv *ToolItem) SetTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_tool_item_set_tooltip_markup((*C.GtkToolItem)(recv.native), c_markup)

	return
}

// Sets the text to be displayed as tooltip on the item.
// See gtk_widget_set_tooltip_text().
/*

C function

gtk_tool_item_set_tooltip_text
*/
func (recv *ToolItem) SetTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_tool_item_set_tooltip_text((*C.GtkToolItem)(recv.native), c_text)

	return
}

// Replaces the widget packed into the tooltip with
// @custom_widget. @custom_widget does not get destroyed when the tooltip goes
// away.
// By default a box with a #GtkImage and #GtkLabel is embedded in
// the tooltip, which can be configured using gtk_tooltip_set_markup()
// and gtk_tooltip_set_icon().
/*

C function

gtk_tooltip_set_custom
*/
func (recv *Tooltip) SetCustom(customWidget *Widget) {
	c_custom_widget := (*C.GtkWidget)(C.NULL)
	if customWidget != nil {
		c_custom_widget = (*C.GtkWidget)(customWidget.ToC())
	}

	C.gtk_tooltip_set_custom((*C.GtkTooltip)(recv.native), c_custom_widget)

	return
}

// Sets the icon of the tooltip (which is in front of the text) to be
// @pixbuf.  If @pixbuf is %NULL, the image will be hidden.
/*

C function

gtk_tooltip_set_icon
*/
func (recv *Tooltip) SetIcon(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_tooltip_set_icon((*C.GtkTooltip)(recv.native), c_pixbuf)

	return
}

// Sets the icon of the tooltip (which is in front of the text) to be
// the stock item indicated by @stock_id with the size indicated
// by @size.  If @stock_id is %NULL, the image will be hidden.
/*

C function

gtk_tooltip_set_icon_from_stock
*/
func (recv *Tooltip) SetIconFromStock(stockId string, size IconSize) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_size := (C.GtkIconSize)(size)

	C.gtk_tooltip_set_icon_from_stock((*C.GtkTooltip)(recv.native), c_stock_id, c_size)

	return
}

// Sets the text of the tooltip to be @markup, which is marked up
// with the [Pango text markup language][PangoMarkupFormat].
// If @markup is %NULL, the label will be hidden.
/*

C function

gtk_tooltip_set_markup
*/
func (recv *Tooltip) SetMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_tooltip_set_markup((*C.GtkTooltip)(recv.native), c_markup)

	return
}

// Sets the text of the tooltip to be @text. If @text is %NULL, the label
// will be hidden. See also gtk_tooltip_set_markup().
/*

C function

gtk_tooltip_set_text
*/
func (recv *Tooltip) SetText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_tooltip_set_text((*C.GtkTooltip)(recv.native), c_text)

	return
}

// Unsupported : gtk_tooltip_set_tip_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// Unsupported : gtk_tree_store_set_valuesv : unsupported parameter values :

// Converts bin_window coordinates to coordinates for the
// tree (the full scrollable area of the tree).
/*

C function

gtk_tree_view_convert_bin_window_to_tree_coords
*/
func (recv *TreeView) ConvertBinWindowToTreeCoords(bx int32, by int32) (int32, int32) {
	c_bx := (C.gint)(bx)

	c_by := (C.gint)(by)

	var c_tx C.gint

	var c_ty C.gint

	C.gtk_tree_view_convert_bin_window_to_tree_coords((*C.GtkTreeView)(recv.native), c_bx, c_by, &c_tx, &c_ty)

	tx := (int32)(c_tx)

	ty := (int32)(c_ty)

	return tx, ty
}

// Converts bin_window coordinates (see gtk_tree_view_get_bin_window())
// to widget relative coordinates.
/*

C function

gtk_tree_view_convert_bin_window_to_widget_coords
*/
func (recv *TreeView) ConvertBinWindowToWidgetCoords(bx int32, by int32) (int32, int32) {
	c_bx := (C.gint)(bx)

	c_by := (C.gint)(by)

	var c_wx C.gint

	var c_wy C.gint

	C.gtk_tree_view_convert_bin_window_to_widget_coords((*C.GtkTreeView)(recv.native), c_bx, c_by, &c_wx, &c_wy)

	wx := (int32)(c_wx)

	wy := (int32)(c_wy)

	return wx, wy
}

// Converts tree coordinates (coordinates in full scrollable area of the tree)
// to bin_window coordinates.
/*

C function

gtk_tree_view_convert_tree_to_bin_window_coords
*/
func (recv *TreeView) ConvertTreeToBinWindowCoords(tx int32, ty int32) (int32, int32) {
	c_tx := (C.gint)(tx)

	c_ty := (C.gint)(ty)

	var c_bx C.gint

	var c_by C.gint

	C.gtk_tree_view_convert_tree_to_bin_window_coords((*C.GtkTreeView)(recv.native), c_tx, c_ty, &c_bx, &c_by)

	bx := (int32)(c_bx)

	by := (int32)(c_by)

	return bx, by
}

// Converts tree coordinates (coordinates in full scrollable area of the tree)
// to widget coordinates.
/*

C function

gtk_tree_view_convert_tree_to_widget_coords
*/
func (recv *TreeView) ConvertTreeToWidgetCoords(tx int32, ty int32) (int32, int32) {
	c_tx := (C.gint)(tx)

	c_ty := (C.gint)(ty)

	var c_wx C.gint

	var c_wy C.gint

	C.gtk_tree_view_convert_tree_to_widget_coords((*C.GtkTreeView)(recv.native), c_tx, c_ty, &c_wx, &c_wy)

	wx := (int32)(c_wx)

	wy := (int32)(c_wy)

	return wx, wy
}

// Converts widget coordinates to coordinates for the bin_window
// (see gtk_tree_view_get_bin_window()).
/*

C function

gtk_tree_view_convert_widget_to_bin_window_coords
*/
func (recv *TreeView) ConvertWidgetToBinWindowCoords(wx int32, wy int32) (int32, int32) {
	c_wx := (C.gint)(wx)

	c_wy := (C.gint)(wy)

	var c_bx C.gint

	var c_by C.gint

	C.gtk_tree_view_convert_widget_to_bin_window_coords((*C.GtkTreeView)(recv.native), c_wx, c_wy, &c_bx, &c_by)

	bx := (int32)(c_bx)

	by := (int32)(c_by)

	return bx, by
}

// Converts widget coordinates to coordinates for the
// tree (the full scrollable area of the tree).
/*

C function

gtk_tree_view_convert_widget_to_tree_coords
*/
func (recv *TreeView) ConvertWidgetToTreeCoords(wx int32, wy int32) (int32, int32) {
	c_wx := (C.gint)(wx)

	c_wy := (C.gint)(wy)

	var c_tx C.gint

	var c_ty C.gint

	C.gtk_tree_view_convert_widget_to_tree_coords((*C.GtkTreeView)(recv.native), c_wx, c_wy, &c_tx, &c_ty)

	tx := (int32)(c_tx)

	ty := (int32)(c_ty)

	return tx, ty
}

// Returns the amount, in pixels, of extra indentation for child levels
// in @tree_view.
/*

C function

gtk_tree_view_get_level_indentation
*/
func (recv *TreeView) GetLevelIndentation() int32 {
	retC := C.gtk_tree_view_get_level_indentation((*C.GtkTreeView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns whether or not expanders are drawn in @tree_view.
/*

C function

gtk_tree_view_get_show_expanders
*/
func (recv *TreeView) GetShowExpanders() bool {
	retC := C.gtk_tree_view_get_show_expanders((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the column of @tree_view’s model which is being used for
// displaying tooltips on @tree_view’s rows.
/*

C function

gtk_tree_view_get_tooltip_column
*/
func (recv *TreeView) GetTooltipColumn() int32 {
	retC := C.gtk_tree_view_get_tooltip_column((*C.GtkTreeView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// This function is supposed to be used in a #GtkWidget::query-tooltip
// signal handler for #GtkTreeView.  The @x, @y and @keyboard_tip values
// which are received in the signal handler, should be passed to this
// function without modification.
//
// The return value indicates whether there is a tree view row at the given
// coordinates (%TRUE) or not (%FALSE) for mouse tooltips.  For keyboard
// tooltips the row returned will be the cursor row.  When %TRUE, then any of
// @model, @path and @iter which have been provided will be set to point to
// that row and the corresponding model.  @x and @y will always be converted
// to be relative to @tree_view’s bin_window if @keyboard_tooltip is %FALSE.
/*

C function

gtk_tree_view_get_tooltip_context
*/
func (recv *TreeView) GetTooltipContext(x int32, y int32, keyboardTip bool) (bool, *TreeModel, *TreePath, *TreeIter) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_keyboard_tip :=
		boolToGboolean(keyboardTip)

	var c_model *C.GtkTreeModel

	var c_path *C.GtkTreePath

	var c_iter C.GtkTreeIter

	retC := C.gtk_tree_view_get_tooltip_context((*C.GtkTreeView)(recv.native), &c_x, &c_y, c_keyboard_tip, &c_model, &c_path, &c_iter)
	retGo := retC == C.TRUE

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, model, path, iter
}

// Returns whether a rubber banding operation is currently being done
// in @tree_view.
/*

C function

gtk_tree_view_is_rubber_banding_active
*/
func (recv *TreeView) IsRubberBandingActive() bool {
	retC := C.gtk_tree_view_is_rubber_banding_active((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Sets the amount of extra indentation for child levels to use in @tree_view
// in addition to the default indentation.  The value should be specified in
// pixels, a value of 0 disables this feature and in this case only the default
// indentation will be used.
// This does not have any visible effects for lists.
/*

C function

gtk_tree_view_set_level_indentation
*/
func (recv *TreeView) SetLevelIndentation(indentation int32) {
	c_indentation := (C.gint)(indentation)

	C.gtk_tree_view_set_level_indentation((*C.GtkTreeView)(recv.native), c_indentation)

	return
}

// Sets whether to draw and enable expanders and indent child rows in
// @tree_view.  When disabled there will be no expanders visible in trees
// and there will be no way to expand and collapse rows by default.  Also
// note that hiding the expanders will disable the default indentation.  You
// can set a custom indentation in this case using
// gtk_tree_view_set_level_indentation().
// This does not have any visible effects for lists.
/*

C function

gtk_tree_view_set_show_expanders
*/
func (recv *TreeView) SetShowExpanders(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.gtk_tree_view_set_show_expanders((*C.GtkTreeView)(recv.native), c_enabled)

	return
}

// Sets the tip area of @tooltip to the area @path, @column and @cell have
// in common.  For example if @path is %NULL and @column is set, the tip
// area will be set to the full area covered by @column.  See also
// gtk_tooltip_set_tip_area().
//
// Note that if @path is not specified and @cell is set and part of a column
// containing the expander, the tooltip might not show and hide at the correct
// position.  In such cases @path must be set to the current node under the
// mouse cursor for this function to operate correctly.
//
// See also gtk_tree_view_set_tooltip_column() for a simpler alternative.
/*

C function

gtk_tree_view_set_tooltip_cell
*/
func (recv *TreeView) SetTooltipCell(tooltip *Tooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer) {
	c_tooltip := (*C.GtkTooltip)(C.NULL)
	if tooltip != nil {
		c_tooltip = (*C.GtkTooltip)(tooltip.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_column := (*C.GtkTreeViewColumn)(C.NULL)
	if column != nil {
		c_column = (*C.GtkTreeViewColumn)(column.ToC())
	}

	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	C.gtk_tree_view_set_tooltip_cell((*C.GtkTreeView)(recv.native), c_tooltip, c_path, c_column, c_cell)

	return
}

// If you only plan to have simple (text-only) tooltips on full rows, you
// can use this function to have #GtkTreeView handle these automatically
// for you. @column should be set to the column in @tree_view’s model
// containing the tooltip texts, or -1 to disable this feature.
//
// When enabled, #GtkWidget:has-tooltip will be set to %TRUE and
// @tree_view will connect a #GtkWidget::query-tooltip signal handler.
//
// Note that the signal handler sets the text with gtk_tooltip_set_markup(),
// so &, <, etc have to be escaped in the text.
/*

C function

gtk_tree_view_set_tooltip_column
*/
func (recv *TreeView) SetTooltipColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_tree_view_set_tooltip_column((*C.GtkTreeView)(recv.native), c_column)

	return
}

// Sets the tip area of @tooltip to be the area covered by the row at @path.
// See also gtk_tree_view_set_tooltip_column() for a simpler alternative.
// See also gtk_tooltip_set_tip_area().
/*

C function

gtk_tree_view_set_tooltip_row
*/
func (recv *TreeView) SetTooltipRow(tooltip *Tooltip, path *TreePath) {
	c_tooltip := (*C.GtkTooltip)(C.NULL)
	if tooltip != nil {
		c_tooltip = (*C.GtkTooltip)(tooltip.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_view_set_tooltip_row((*C.GtkTreeView)(recv.native), c_tooltip, c_path)

	return
}

// Returns the #GtkTreeView wherein @tree_column has been inserted.
// If @column is currently not inserted in any tree view, %NULL is
// returned.
/*

C function

gtk_tree_view_column_get_tree_view
*/
func (recv *TreeViewColumn) GetTreeView() *Widget {
	retC := C.gtk_tree_view_column_get_tree_view((*C.GtkTreeViewColumn)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Creates a #GtkVolumeButton, with a range between 0.0 and 1.0, with
// a stepping of 0.02. Volume values can be obtained and modified using
// the functions from #GtkScaleButton.
/*

C function

gtk_volume_button_new
*/
func VolumeButtonNew() *VolumeButton {
	retC := C.gtk_volume_button_new()
	retGo := VolumeButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported signal 'drag-failed' for Widget : unsupported parameter result : type DragResult :

// Unsupported signal 'keynav-failed' for Widget : unsupported parameter direction : type DirectionType :

// Unsupported signal 'query-tooltip' for Widget : unsupported parameter x : type gint :

// Notifies the user about an input-related error on this widget.
// If the #GtkSettings:gtk-error-bell setting is %TRUE, it calls
// gdk_window_beep(), otherwise it does nothing.
//
// Note that the effect of gdk_window_beep() can be configured in many
// ways, depending on the windowing backend and the desktop environment
// or window manager that is used.
/*

C function

gtk_widget_error_bell
*/
func (recv *Widget) ErrorBell() {
	C.gtk_widget_error_bell((*C.GtkWidget)(recv.native))

	return
}

// Returns the current value of the has-tooltip property.  See
// #GtkWidget:has-tooltip for more information.
/*

C function

gtk_widget_get_has_tooltip
*/
func (recv *Widget) GetHasTooltip() bool {
	retC := C.gtk_widget_get_has_tooltip((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the contents of the tooltip for @widget.
/*

C function

gtk_widget_get_tooltip_markup
*/
func (recv *Widget) GetTooltipMarkup() string {
	retC := C.gtk_widget_get_tooltip_markup((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the contents of the tooltip for @widget.
/*

C function

gtk_widget_get_tooltip_text
*/
func (recv *Widget) GetTooltipText() string {
	retC := C.gtk_widget_get_tooltip_text((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Returns the #GtkWindow of the current tooltip. This can be the
// GtkWindow created by default, or the custom tooltip window set
// using gtk_widget_set_tooltip_window().
/*

C function

gtk_widget_get_tooltip_window
*/
func (recv *Widget) GetTooltipWindow() *Window {
	retC := C.gtk_widget_get_tooltip_window((*C.GtkWidget)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// This function should be called whenever keyboard navigation within
// a single widget hits a boundary. The function emits the
// #GtkWidget::keynav-failed signal on the widget and its return
// value should be interpreted in a way similar to the return value of
// gtk_widget_child_focus():
//
// When %TRUE is returned, stay in the widget, the failed keyboard
// navigation is OK and/or there is nowhere we can/should move the
// focus to.
//
// When %FALSE is returned, the caller should continue with keyboard
// navigation outside the widget, e.g. by calling
// gtk_widget_child_focus() on the widget’s toplevel.
//
// The default ::keynav-failed handler returns %FALSE for
// %GTK_DIR_TAB_FORWARD and %GTK_DIR_TAB_BACKWARD. For the other
// values of #GtkDirectionType it returns %TRUE.
//
// Whenever the default handler returns %TRUE, it also calls
// gtk_widget_error_bell() to notify the user of the failed keyboard
// navigation.
//
// A use case for providing an own implementation of ::keynav-failed
// (either by connecting to it or by overriding it) would be a row of
// #GtkEntry widgets where the user should be able to navigate the
// entire row with the cursor keys, as e.g. known from user interfaces
// that require entering license keys.
/*

C function

gtk_widget_keynav_failed
*/
func (recv *Widget) KeynavFailed(direction DirectionType) bool {
	c_direction := (C.GtkDirectionType)(direction)

	retC := C.gtk_widget_keynav_failed((*C.GtkWidget)(recv.native), c_direction)
	retGo := retC == C.TRUE

	return retGo
}

// Sets the cursor color to use in a widget, overriding the #GtkWidget
// cursor-color and secondary-cursor-color
// style properties.
//
// All other style values are left untouched.
// See also gtk_widget_modify_style().
/*

C function

gtk_widget_modify_cursor
*/
func (recv *Widget) ModifyCursor(primary *gdk.Color, secondary *gdk.Color) {
	c_primary := (*C.GdkColor)(C.NULL)
	if primary != nil {
		c_primary = (*C.GdkColor)(primary.ToC())
	}

	c_secondary := (*C.GdkColor)(C.NULL)
	if secondary != nil {
		c_secondary = (*C.GdkColor)(secondary.ToC())
	}

	C.gtk_widget_modify_cursor((*C.GtkWidget)(recv.native), c_primary, c_secondary)

	return
}

// Sets the has-tooltip property on @widget to @has_tooltip.  See
// #GtkWidget:has-tooltip for more information.
/*

C function

gtk_widget_set_has_tooltip
*/
func (recv *Widget) SetHasTooltip(hasTooltip bool) {
	c_has_tooltip :=
		boolToGboolean(hasTooltip)

	C.gtk_widget_set_has_tooltip((*C.GtkWidget)(recv.native), c_has_tooltip)

	return
}

// Sets @markup as the contents of the tooltip, which is marked up with
// the [Pango text markup language][PangoMarkupFormat].
//
// This function will take care of setting #GtkWidget:has-tooltip to %TRUE
// and of the default handler for the #GtkWidget::query-tooltip signal.
//
// See also the #GtkWidget:tooltip-markup property and
// gtk_tooltip_set_markup().
/*

C function

gtk_widget_set_tooltip_markup
*/
func (recv *Widget) SetTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_widget_set_tooltip_markup((*C.GtkWidget)(recv.native), c_markup)

	return
}

// Sets @text as the contents of the tooltip. This function will take
// care of setting #GtkWidget:has-tooltip to %TRUE and of the default
// handler for the #GtkWidget::query-tooltip signal.
//
// See also the #GtkWidget:tooltip-text property and gtk_tooltip_set_text().
/*

C function

gtk_widget_set_tooltip_text
*/
func (recv *Widget) SetTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_widget_set_tooltip_text((*C.GtkWidget)(recv.native), c_text)

	return
}

// Replaces the default window used for displaying
// tooltips with @custom_window. GTK+ will take care of showing and
// hiding @custom_window at the right moment, to behave likewise as
// the default tooltip window. If @custom_window is %NULL, the default
// tooltip window will be used.
/*

C function

gtk_widget_set_tooltip_window
*/
func (recv *Widget) SetTooltipWindow(customWindow *Window) {
	c_custom_window := (*C.GtkWindow)(C.NULL)
	if customWindow != nil {
		c_custom_window = (*C.GtkWindow)(customWindow.ToC())
	}

	C.gtk_widget_set_tooltip_window((*C.GtkWidget)(recv.native), c_custom_window)

	return
}

// Triggers a tooltip query on the display where the toplevel of @widget
// is located. See gtk_tooltip_trigger_tooltip_query() for more
// information.
/*

C function

gtk_widget_trigger_tooltip_query
*/
func (recv *Widget) TriggerTooltipQuery() {
	C.gtk_widget_trigger_tooltip_query((*C.GtkWidget)(recv.native))

	return
}

// Fetches the requested opacity for this window. See
// gtk_window_set_opacity().
/*

C function

gtk_window_get_opacity
*/
func (recv *Window) GetOpacity() float64 {
	retC := C.gtk_window_get_opacity((*C.GtkWindow)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Request the windowing system to make @window partially transparent,
// with opacity 0 being fully transparent and 1 fully opaque. (Values
// of the opacity parameter are clamped to the [0,1] range.) On X11
// this has any effect only on X screens with a compositing manager
// running. See gtk_widget_is_composited(). On Windows it should work
// always.
//
// Note that setting a window’s opacity after the window has been
// shown causes it to flicker once on Windows.
/*

C function

gtk_window_set_opacity
*/
func (recv *Window) SetOpacity(opacity float64) {
	c_opacity := (C.gdouble)(opacity)

	C.gtk_window_set_opacity((*C.GtkWindow)(recv.native), c_opacity)

	return
}

// Startup notification identifiers are used by desktop environment to
// track application startup, to provide user feedback and other
// features. This function changes the corresponding property on the
// underlying GdkWindow. Normally, startup identifier is managed
// automatically and you should only use this function in special cases
// like transferring focus from other processes. You should use this
// function before calling gtk_window_present() or any equivalent
// function generating a window map event.
//
// This function is only useful on X11, not with other GTK+ targets.
/*

C function

gtk_window_set_startup_id
*/
func (recv *Window) SetStartupId(startupId string) {
	c_startup_id := C.CString(startupId)
	defer C.free(unsafe.Pointer(c_startup_id))

	C.gtk_window_set_startup_id((*C.GtkWindow)(recv.native), c_startup_id)

	return
}
