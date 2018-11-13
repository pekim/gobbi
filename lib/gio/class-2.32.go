// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_app_launch_context_get_environment : no return type

// Arranges for @variable to be set to @value in the child's
// environment when @context is used to launch an application.
/*

C function : g_app_launch_context_setenv
*/
func (recv *AppLaunchContext) Setenv(variable string, value string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_app_launch_context_setenv((*C.GAppLaunchContext)(recv.native), c_variable, c_value)

	return
}

// Arranges for @variable to be unset in the child's environment
// when @context is used to launch an application.
/*

C function : g_app_launch_context_unsetenv
*/
func (recv *AppLaunchContext) Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_app_launch_context_unsetenv((*C.GAppLaunchContext)(recv.native), c_variable)

	return
}

// Immediately quits the application.
//
// Upon return to the mainloop, g_application_run() will return,
// calling only the 'shutdown' function before doing so.
//
// The hold count is ignored.
// Take care if your code has called g_application_hold() on the application and
// is therefore still expecting it to exist.
// (Note that you may have called g_application_hold() indirectly, for example
// through gtk_application_add_window().)
//
// The result of calling g_application_run() again after it returns is
// unspecified.
/*

C function : g_application_quit
*/
func (recv *Application) Quit() {
	C.g_application_quit((*C.GApplication)(recv.native))

	return
}

// Sets or unsets the default application for the process, as returned
// by g_application_get_default().
//
// This function does not take its own reference on @application.  If
// @application is destroyed then the default application will revert
// back to %NULL.
/*

C function : g_application_set_default
*/
func (recv *Application) SetDefault() {
	C.g_application_set_default((*C.GApplication)(recv.native))

	return
}

// Exports @action_group on @connection at @object_path.
//
// The implemented D-Bus API should be considered private.  It is
// subject to change in the future.
//
// A given object path can only have one action group exported on it.
// If this constraint is violated, the export will fail and 0 will be
// returned (with @error set accordingly).
//
// You can unexport the action group using
// g_dbus_connection_unexport_action_group() with the return value of
// this function.
//
// The thread default main context is taken at the time of this call.
// All incoming action activations and state change requests are
// reported from this context.  Any changes on the action group that
// cause it to emit signals must also come from this same context.
// Since incoming action activations and state change requests are
// rather likely to cause changes on the action group, this effectively
// limits a given action group to being exported from only one main
// context.
/*

C function : g_dbus_connection_export_action_group
*/
func (recv *DBusConnection) ExportActionGroup(objectPath string, actionGroup *ActionGroup) (uint32, error) {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_action_group := (*C.GActionGroup)(actionGroup.ToC())

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_export_action_group((*C.GDBusConnection)(recv.native), c_object_path, c_action_group, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Exports @menu on @connection at @object_path.
//
// The implemented D-Bus API should be considered private.
// It is subject to change in the future.
//
// An object path can only have one menu model exported on it. If this
// constraint is violated, the export will fail and 0 will be
// returned (with @error set accordingly).
//
// You can unexport the menu model using
// g_dbus_connection_unexport_menu_model() with the return value of
// this function.
/*

C function : g_dbus_connection_export_menu_model
*/
func (recv *DBusConnection) ExportMenuModel(objectPath string, menu *MenuModel) (uint32, error) {
	c_object_path := C.CString(objectPath)
	defer C.free(unsafe.Pointer(c_object_path))

	c_menu := (*C.GMenuModel)(C.NULL)
	if menu != nil {
		c_menu = (*C.GMenuModel)(menu.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dbus_connection_export_menu_model((*C.GDBusConnection)(recv.native), c_object_path, c_menu, &cThrowableError)
	retGo := (uint32)(retC)

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Reverses the effect of a previous call to
// g_dbus_connection_export_action_group().
//
// It is an error to call this function with an ID that wasn't returned
// from g_dbus_connection_export_action_group() or to call it with the
// same ID more than once.
/*

C function : g_dbus_connection_unexport_action_group
*/
func (recv *DBusConnection) UnexportActionGroup(exportId uint32) {
	c_export_id := (C.guint)(exportId)

	C.g_dbus_connection_unexport_action_group((*C.GDBusConnection)(recv.native), c_export_id)

	return
}

// Reverses the effect of a previous call to
// g_dbus_connection_export_menu_model().
//
// It is an error to call this function with an ID that wasn't returned
// from g_dbus_connection_export_menu_model() or to call it with the
// same ID more than once.
/*

C function : g_dbus_connection_unexport_menu_model
*/
func (recv *DBusConnection) UnexportMenuModel(exportId uint32) {
	c_export_id := (C.guint)(exportId)

	C.g_dbus_connection_unexport_menu_model((*C.GDBusConnection)(recv.native), c_export_id)

	return
}

// Gets a list of the connections that @interface_ is exported on.
/*

C function : g_dbus_interface_skeleton_get_connections
*/
func (recv *DBusInterfaceSkeleton) GetConnections() *glib.List {
	retC := C.g_dbus_interface_skeleton_get_connections((*C.GDBusInterfaceSkeleton)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks if @interface_ is exported on @connection.
/*

C function : g_dbus_interface_skeleton_has_connection
*/
func (recv *DBusInterfaceSkeleton) HasConnection(connection *DBusConnection) bool {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	retC := C.g_dbus_interface_skeleton_has_connection((*C.GDBusInterfaceSkeleton)(recv.native), c_connection)
	retGo := retC == C.TRUE

	return retGo
}

// Stops exporting @interface_ on @connection.
//
// To stop exporting on all connections the interface is exported on,
// use g_dbus_interface_skeleton_unexport().
/*

C function : g_dbus_interface_skeleton_unexport_from_connection
*/
func (recv *DBusInterfaceSkeleton) UnexportFromConnection(connection *DBusConnection) {
	c_connection := (*C.GDBusConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GDBusConnection)(connection.ToC())
	}

	C.g_dbus_interface_skeleton_unexport_from_connection((*C.GDBusInterfaceSkeleton)(recv.native), c_connection)

	return
}

// Unsupported : g_desktop_app_info_get_keywords : no return type

// InetAddressMask is a wrapper around the C record GInetAddressMask.
type InetAddressMask struct {
	native *C.GInetAddressMask
	// parent_instance : record
	// Private : priv
}

func InetAddressMaskNewFromC(u unsafe.Pointer) *InetAddressMask {
	c := (*C.GInetAddressMask)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMask{native: c}

	return g
}

func (recv *InetAddressMask) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *InetAddressMask) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to InetAddressMask.
// Exercise care, as this is a potentially dangerous function if the Object is not a InetAddressMask.
func CastToInetAddressMask(object *gobject.Object) *InetAddressMask {
	return InetAddressMaskNewFromC(object.ToC())
}

// Creates a new #GInetAddressMask representing all addresses whose
// first @length bits match @addr.
/*

C function : g_inet_address_mask_new
*/
func InetAddressMaskNew(addr *InetAddress, length uint32) (*InetAddressMask, error) {
	c_addr := (*C.GInetAddress)(C.NULL)
	if addr != nil {
		c_addr = (*C.GInetAddress)(addr.ToC())
	}

	c_length := (C.guint)(length)

	var cThrowableError *C.GError

	retC := C.g_inet_address_mask_new(c_addr, c_length, &cThrowableError)
	retGo := InetAddressMaskNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Parses @mask_string as an IP address and (optional) length, and
// creates a new #GInetAddressMask. The length, if present, is
// delimited by a "/". If it is not present, then the length is
// assumed to be the full length of the address.
/*

C function : g_inet_address_mask_new_from_string
*/
func InetAddressMaskNewFromString(maskString string) (*InetAddressMask, error) {
	c_mask_string := C.CString(maskString)
	defer C.free(unsafe.Pointer(c_mask_string))

	var cThrowableError *C.GError

	retC := C.g_inet_address_mask_new_from_string(c_mask_string, &cThrowableError)
	retGo := InetAddressMaskNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Tests if @mask and @mask2 are the same mask.
/*

C function : g_inet_address_mask_equal
*/
func (recv *InetAddressMask) Equal(mask2 *InetAddressMask) bool {
	c_mask2 := (*C.GInetAddressMask)(C.NULL)
	if mask2 != nil {
		c_mask2 = (*C.GInetAddressMask)(mask2.ToC())
	}

	retC := C.g_inet_address_mask_equal((*C.GInetAddressMask)(recv.native), c_mask2)
	retGo := retC == C.TRUE

	return retGo
}

// Gets @mask's base address
/*

C function : g_inet_address_mask_get_address
*/
func (recv *InetAddressMask) GetAddress() *InetAddress {
	retC := C.g_inet_address_mask_get_address((*C.GInetAddressMask)(recv.native))
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the #GSocketFamily of @mask's address
/*

C function : g_inet_address_mask_get_family
*/
func (recv *InetAddressMask) GetFamily() SocketFamily {
	retC := C.g_inet_address_mask_get_family((*C.GInetAddressMask)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// Gets @mask's length
/*

C function : g_inet_address_mask_get_length
*/
func (recv *InetAddressMask) GetLength() uint32 {
	retC := C.g_inet_address_mask_get_length((*C.GInetAddressMask)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Tests if @address falls within the range described by @mask.
/*

C function : g_inet_address_mask_matches
*/
func (recv *InetAddressMask) Matches(address *InetAddress) bool {
	c_address := (*C.GInetAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GInetAddress)(address.ToC())
	}

	retC := C.g_inet_address_mask_matches((*C.GInetAddressMask)(recv.native), c_address)
	retGo := retC == C.TRUE

	return retGo
}

// Converts @mask back to its corresponding string form.
/*

C function : g_inet_address_mask_to_string
*/
func (recv *InetAddressMask) ToString() string {
	retC := C.g_inet_address_mask_to_string((*C.GInetAddressMask)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the `sin6_flowinfo` field from @address,
// which must be an IPv6 address.
/*

C function : g_inet_socket_address_get_flowinfo
*/
func (recv *InetSocketAddress) GetFlowinfo() uint32 {
	retC := C.g_inet_socket_address_get_flowinfo((*C.GInetSocketAddress)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the `sin6_scope_id` field from @address,
// which must be an IPv6 address.
/*

C function : g_inet_socket_address_get_scope_id
*/
func (recv *InetSocketAddress) GetScopeId() uint32 {
	retC := C.g_inet_socket_address_get_scope_id((*C.GInetSocketAddress)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Menu is a wrapper around the C record GMenu.
type Menu struct {
	native *C.GMenu
}

func MenuNewFromC(u unsafe.Pointer) *Menu {
	c := (*C.GMenu)(u)
	if c == nil {
		return nil
	}

	g := &Menu{native: c}

	return g
}

func (recv *Menu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuModel upcasts to *MenuModel
func (recv *Menu) MenuModel() *MenuModel {
	return MenuModelNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Menu) Object() *gobject.Object {
	return recv.MenuModel().Object()
}

// CastToWidget down casts any arbitary Object to Menu.
// Exercise care, as this is a potentially dangerous function if the Object is not a Menu.
func CastToMenu(object *gobject.Object) *Menu {
	return MenuNewFromC(object.ToC())
}

// Creates a new #GMenu.
//
// The new menu has no items.
/*

C function : g_menu_new
*/
func MenuNew() *Menu {
	retC := C.g_menu_new()
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Convenience function for appending a normal menu item to the end of
// @menu.  Combine g_menu_item_new() and g_menu_insert_item() for a more
// flexible alternative.
/*

C function : g_menu_append
*/
func (recv *Menu) Append(label string, detailedAction string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_append((*C.GMenu)(recv.native), c_label, c_detailed_action)

	return
}

// Appends @item to the end of @menu.
//
// See g_menu_insert_item() for more information.
/*

C function : g_menu_append_item
*/
func (recv *Menu) AppendItem(item *MenuItem) {
	c_item := (*C.GMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.GMenuItem)(item.ToC())
	}

	C.g_menu_append_item((*C.GMenu)(recv.native), c_item)

	return
}

// Convenience function for appending a section menu item to the end of
// @menu.  Combine g_menu_item_new_section() and g_menu_insert_item() for a
// more flexible alternative.
/*

C function : g_menu_append_section
*/
func (recv *Menu) AppendSection(label string, section *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(C.NULL)
	if section != nil {
		c_section = (*C.GMenuModel)(section.ToC())
	}

	C.g_menu_append_section((*C.GMenu)(recv.native), c_label, c_section)

	return
}

// Convenience function for appending a submenu menu item to the end of
// @menu.  Combine g_menu_item_new_submenu() and g_menu_insert_item() for a
// more flexible alternative.
/*

C function : g_menu_append_submenu
*/
func (recv *Menu) AppendSubmenu(label string, submenu *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.GMenuModel)(submenu.ToC())
	}

	C.g_menu_append_submenu((*C.GMenu)(recv.native), c_label, c_submenu)

	return
}

// Marks @menu as frozen.
//
// After the menu is frozen, it is an error to attempt to make any
// changes to it.  In effect this means that the #GMenu API must no
// longer be used.
//
// This function causes g_menu_model_is_mutable() to begin returning
// %FALSE, which has some positive performance implications.
/*

C function : g_menu_freeze
*/
func (recv *Menu) Freeze() {
	C.g_menu_freeze((*C.GMenu)(recv.native))

	return
}

// Convenience function for inserting a normal menu item into @menu.
// Combine g_menu_item_new() and g_menu_insert_item() for a more flexible
// alternative.
/*

C function : g_menu_insert
*/
func (recv *Menu) Insert(position int32, label string, detailedAction string) {
	c_position := (C.gint)(position)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_insert((*C.GMenu)(recv.native), c_position, c_label, c_detailed_action)

	return
}

// Inserts @item into @menu.
//
// The "insertion" is actually done by copying all of the attribute and
// link values of @item and using them to form a new item within @menu.
// As such, @item itself is not really inserted, but rather, a menu item
// that is exactly the same as the one presently described by @item.
//
// This means that @item is essentially useless after the insertion
// occurs.  Any changes you make to it are ignored unless it is inserted
// again (at which point its updated values will be copied).
//
// You should probably just free @item once you're done.
//
// There are many convenience functions to take care of common cases.
// See g_menu_insert(), g_menu_insert_section() and
// g_menu_insert_submenu() as well as "prepend" and "append" variants of
// each of these functions.
/*

C function : g_menu_insert_item
*/
func (recv *Menu) InsertItem(position int32, item *MenuItem) {
	c_position := (C.gint)(position)

	c_item := (*C.GMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.GMenuItem)(item.ToC())
	}

	C.g_menu_insert_item((*C.GMenu)(recv.native), c_position, c_item)

	return
}

// Convenience function for inserting a section menu item into @menu.
// Combine g_menu_item_new_section() and g_menu_insert_item() for a more
// flexible alternative.
/*

C function : g_menu_insert_section
*/
func (recv *Menu) InsertSection(position int32, label string, section *MenuModel) {
	c_position := (C.gint)(position)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(C.NULL)
	if section != nil {
		c_section = (*C.GMenuModel)(section.ToC())
	}

	C.g_menu_insert_section((*C.GMenu)(recv.native), c_position, c_label, c_section)

	return
}

// Convenience function for inserting a submenu menu item into @menu.
// Combine g_menu_item_new_submenu() and g_menu_insert_item() for a more
// flexible alternative.
/*

C function : g_menu_insert_submenu
*/
func (recv *Menu) InsertSubmenu(position int32, label string, submenu *MenuModel) {
	c_position := (C.gint)(position)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.GMenuModel)(submenu.ToC())
	}

	C.g_menu_insert_submenu((*C.GMenu)(recv.native), c_position, c_label, c_submenu)

	return
}

// Convenience function for prepending a normal menu item to the start
// of @menu.  Combine g_menu_item_new() and g_menu_insert_item() for a more
// flexible alternative.
/*

C function : g_menu_prepend
*/
func (recv *Menu) Prepend(label string, detailedAction string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_prepend((*C.GMenu)(recv.native), c_label, c_detailed_action)

	return
}

// Prepends @item to the start of @menu.
//
// See g_menu_insert_item() for more information.
/*

C function : g_menu_prepend_item
*/
func (recv *Menu) PrependItem(item *MenuItem) {
	c_item := (*C.GMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.GMenuItem)(item.ToC())
	}

	C.g_menu_prepend_item((*C.GMenu)(recv.native), c_item)

	return
}

// Convenience function for prepending a section menu item to the start
// of @menu.  Combine g_menu_item_new_section() and g_menu_insert_item() for
// a more flexible alternative.
/*

C function : g_menu_prepend_section
*/
func (recv *Menu) PrependSection(label string, section *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(C.NULL)
	if section != nil {
		c_section = (*C.GMenuModel)(section.ToC())
	}

	C.g_menu_prepend_section((*C.GMenu)(recv.native), c_label, c_section)

	return
}

// Convenience function for prepending a submenu menu item to the start
// of @menu.  Combine g_menu_item_new_submenu() and g_menu_insert_item() for
// a more flexible alternative.
/*

C function : g_menu_prepend_submenu
*/
func (recv *Menu) PrependSubmenu(label string, submenu *MenuModel) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.GMenuModel)(submenu.ToC())
	}

	C.g_menu_prepend_submenu((*C.GMenu)(recv.native), c_label, c_submenu)

	return
}

// Removes an item from the menu.
//
// @position gives the index of the item to remove.
//
// It is an error if position is not in range the range from 0 to one
// less than the number of items in the menu.
//
// It is not possible to remove items by identity since items are added
// to the menu simply by copying their links and attributes (ie:
// identity of the item itself is not preserved).
/*

C function : g_menu_remove
*/
func (recv *Menu) Remove(position int32) {
	c_position := (C.gint)(position)

	C.g_menu_remove((*C.GMenu)(recv.native), c_position)

	return
}

// MenuAttributeIter is a wrapper around the C record GMenuAttributeIter.
type MenuAttributeIter struct {
	native *C.GMenuAttributeIter
	// parent_instance : record
	// priv : record
}

func MenuAttributeIterNewFromC(u unsafe.Pointer) *MenuAttributeIter {
	c := (*C.GMenuAttributeIter)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIter{native: c}

	return g
}

func (recv *MenuAttributeIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *MenuAttributeIter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to MenuAttributeIter.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuAttributeIter.
func CastToMenuAttributeIter(object *gobject.Object) *MenuAttributeIter {
	return MenuAttributeIterNewFromC(object.ToC())
}

// Gets the name of the attribute at the current iterator position, as
// a string.
//
// The iterator is not advanced.
/*

C function : g_menu_attribute_iter_get_name
*/
func (recv *MenuAttributeIter) GetName() string {
	retC := C.g_menu_attribute_iter_get_name((*C.GMenuAttributeIter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_menu_attribute_iter_get_next : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_menu_attribute_iter_get_value : return type : Blacklisted record : GVariant

// Attempts to advance the iterator to the next (possibly first)
// attribute.
//
// %TRUE is returned on success, or %FALSE if there are no more
// attributes.
//
// You must call this function when you first acquire the iterator
// to advance it to the first attribute (and determine if the first
// attribute exists at all).
/*

C function : g_menu_attribute_iter_next
*/
func (recv *MenuAttributeIter) Next() bool {
	retC := C.g_menu_attribute_iter_next((*C.GMenuAttributeIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// MenuItem is a wrapper around the C record GMenuItem.
type MenuItem struct {
	native *C.GMenuItem
}

func MenuItemNewFromC(u unsafe.Pointer) *MenuItem {
	c := (*C.GMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &MenuItem{native: c}

	return g
}

func (recv *MenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *MenuItem) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to MenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuItem.
func CastToMenuItem(object *gobject.Object) *MenuItem {
	return MenuItemNewFromC(object.ToC())
}

// Creates a new #GMenuItem.
//
// If @label is non-%NULL it is used to set the "label" attribute of the
// new item.
//
// If @detailed_action is non-%NULL it is used to set the "action" and
// possibly the "target" attribute of the new item.  See
// g_menu_item_set_detailed_action() for more information.
/*

C function : g_menu_item_new
*/
func MenuItemNew(label string, detailedAction string) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	retC := C.g_menu_item_new(c_label, c_detailed_action)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GMenuItem representing a section.
//
// This is a convenience API around g_menu_item_new() and
// g_menu_item_set_section().
//
// The effect of having one menu appear as a section of another is
// exactly as it sounds: the items from @section become a direct part of
// the menu that @menu_item is added to.
//
// Visual separation is typically displayed between two non-empty
// sections.  If @label is non-%NULL then it will be encorporated into
// this visual indication.  This allows for labeled subsections of a
// menu.
//
// As a simple example, consider a typical "Edit" menu from a simple
// program.  It probably contains an "Undo" and "Redo" item, followed by
// a separator, followed by "Cut", "Copy" and "Paste".
//
// This would be accomplished by creating three #GMenu instances.  The
// first would be populated with the "Undo" and "Redo" items, and the
// second with the "Cut", "Copy" and "Paste" items.  The first and
// second menus would then be added as submenus of the third.  In XML
// format, this would look something like the following:
// |[
// <menu id='edit-menu'>
// <section>
// <item label='Undo'/>
// <item label='Redo'/>
// </section>
// <section>
// <item label='Cut'/>
// <item label='Copy'/>
// <item label='Paste'/>
// </section>
// </menu>
// ]|
//
// The following example is exactly equivalent.  It is more illustrative
// of the exact relationship between the menus and items (keeping in
// mind that the 'link' element defines a new menu that is linked to the
// containing one).  The style of the second example is more verbose and
// difficult to read (and therefore not recommended except for the
// purpose of understanding what is really going on).
// |[
// <menu id='edit-menu'>
// <item>
// <link name='section'>
// <item label='Undo'/>
// <item label='Redo'/>
// </link>
// </item>
// <item>
// <link name='section'>
// <item label='Cut'/>
// <item label='Copy'/>
// <item label='Paste'/>
// </link>
// </item>
// </menu>
// ]|
/*

C function : g_menu_item_new_section
*/
func MenuItemNewSection(label string, section *MenuModel) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_section := (*C.GMenuModel)(C.NULL)
	if section != nil {
		c_section = (*C.GMenuModel)(section.ToC())
	}

	retC := C.g_menu_item_new_section(c_label, c_section)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GMenuItem representing a submenu.
//
// This is a convenience API around g_menu_item_new() and
// g_menu_item_set_submenu().
/*

C function : g_menu_item_new_submenu
*/
func MenuItemNewSubmenu(label string, submenu *MenuModel) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.GMenuModel)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.GMenuModel)(submenu.ToC())
	}

	retC := C.g_menu_item_new_submenu(c_label, c_submenu)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_menu_item_set_action_and_target : unsupported parameter ... : varargs

// Unsupported : g_menu_item_set_action_and_target_value : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_menu_item_set_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_set_attribute_value : unsupported parameter value : Blacklisted record : GVariant

// Sets the "action" and possibly the "target" attribute of @menu_item.
//
// The format of @detailed_action is the same format parsed by
// g_action_parse_detailed_name().
//
// See g_menu_item_set_action_and_target() or
// g_menu_item_set_action_and_target_value() for more flexible (but
// slightly less convenient) alternatives.
//
// See also g_menu_item_set_action_and_target_value() for a description of
// the semantics of the action and target attributes.
/*

C function : g_menu_item_set_detailed_action
*/
func (recv *MenuItem) SetDetailedAction(detailedAction string) {
	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_menu_item_set_detailed_action((*C.GMenuItem)(recv.native), c_detailed_action)

	return
}

// Sets or unsets the "label" attribute of @menu_item.
//
// If @label is non-%NULL it is used as the label for the menu item.  If
// it is %NULL then the label attribute is unset.
/*

C function : g_menu_item_set_label
*/
func (recv *MenuItem) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.g_menu_item_set_label((*C.GMenuItem)(recv.native), c_label)

	return
}

// Creates a link from @menu_item to @model if non-%NULL, or unsets it.
//
// Links are used to establish a relationship between a particular menu
// item and another menu.  For example, %G_MENU_LINK_SUBMENU is used to
// associate a submenu with a particular menu item, and %G_MENU_LINK_SECTION
// is used to create a section. Other types of link can be used, but there
// is no guarantee that clients will be able to make sense of them.
// Link types are restricted to lowercase characters, numbers
// and '-'. Furthermore, the names must begin with a lowercase character,
// must not end with a '-', and must not contain consecutive dashes.
/*

C function : g_menu_item_set_link
*/
func (recv *MenuItem) SetLink(link string, model *MenuModel) {
	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	c_model := (*C.GMenuModel)(C.NULL)
	if model != nil {
		c_model = (*C.GMenuModel)(model.ToC())
	}

	C.g_menu_item_set_link((*C.GMenuItem)(recv.native), c_link, c_model)

	return
}

// Sets or unsets the "section" link of @menu_item to @section.
//
// The effect of having one menu appear as a section of another is
// exactly as it sounds: the items from @section become a direct part of
// the menu that @menu_item is added to.  See g_menu_item_new_section()
// for more information about what it means for a menu item to be a
// section.
/*

C function : g_menu_item_set_section
*/
func (recv *MenuItem) SetSection(section *MenuModel) {
	c_section := (*C.GMenuModel)(C.NULL)
	if section != nil {
		c_section = (*C.GMenuModel)(section.ToC())
	}

	C.g_menu_item_set_section((*C.GMenuItem)(recv.native), c_section)

	return
}

// Sets or unsets the "submenu" link of @menu_item to @submenu.
//
// If @submenu is non-%NULL, it is linked to.  If it is %NULL then the
// link is unset.
//
// The effect of having one menu appear as a submenu of another is
// exactly as it sounds.
/*

C function : g_menu_item_set_submenu
*/
func (recv *MenuItem) SetSubmenu(submenu *MenuModel) {
	c_submenu := (*C.GMenuModel)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.GMenuModel)(submenu.ToC())
	}

	C.g_menu_item_set_submenu((*C.GMenuItem)(recv.native), c_submenu)

	return
}

// MenuLinkIter is a wrapper around the C record GMenuLinkIter.
type MenuLinkIter struct {
	native *C.GMenuLinkIter
	// parent_instance : record
	// priv : record
}

func MenuLinkIterNewFromC(u unsafe.Pointer) *MenuLinkIter {
	c := (*C.GMenuLinkIter)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIter{native: c}

	return g
}

func (recv *MenuLinkIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *MenuLinkIter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to MenuLinkIter.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuLinkIter.
func CastToMenuLinkIter(object *gobject.Object) *MenuLinkIter {
	return MenuLinkIterNewFromC(object.ToC())
}

// Gets the name of the link at the current iterator position.
//
// The iterator is not advanced.
/*

C function : g_menu_link_iter_get_name
*/
func (recv *MenuLinkIter) GetName() string {
	retC := C.g_menu_link_iter_get_name((*C.GMenuLinkIter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// This function combines g_menu_link_iter_next() with
// g_menu_link_iter_get_name() and g_menu_link_iter_get_value().
//
// First the iterator is advanced to the next (possibly first) link.
// If that fails, then %FALSE is returned and there are no other effects.
//
// If successful, @out_link and @value are set to the name and #GMenuModel
// of the link that has just been advanced to.  At this point,
// g_menu_link_iter_get_name() and g_menu_link_iter_get_value() will return the
// same values again.
//
// The value returned in @out_link remains valid for as long as the iterator
// remains at the current position.  The value returned in @value must
// be unreffed using g_object_unref() when it is no longer in use.
/*

C function : g_menu_link_iter_get_next
*/
func (recv *MenuLinkIter) GetNext() (bool, string, *MenuModel) {
	var c_out_link *C.gchar

	var c_value *C.GMenuModel

	retC := C.g_menu_link_iter_get_next((*C.GMenuLinkIter)(recv.native), &c_out_link, &c_value)
	retGo := retC == C.TRUE

	outLink := C.GoString(c_out_link)

	value := MenuModelNewFromC(unsafe.Pointer(c_value))

	return retGo, outLink, value
}

// Gets the linked #GMenuModel at the current iterator position.
//
// The iterator is not advanced.
/*

C function : g_menu_link_iter_get_value
*/
func (recv *MenuLinkIter) GetValue() *MenuModel {
	retC := C.g_menu_link_iter_get_value((*C.GMenuLinkIter)(recv.native))
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Attempts to advance the iterator to the next (possibly first)
// link.
//
// %TRUE is returned on success, or %FALSE if there are no more links.
//
// You must call this function when you first acquire the iterator to
// advance it to the first link (and determine if the first link exists
// at all).
/*

C function : g_menu_link_iter_next
*/
func (recv *MenuLinkIter) Next() bool {
	retC := C.g_menu_link_iter_next((*C.GMenuLinkIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// MenuModel is a wrapper around the C record GMenuModel.
type MenuModel struct {
	native *C.GMenuModel
	// parent_instance : record
	// priv : record
}

func MenuModelNewFromC(u unsafe.Pointer) *MenuModel {
	c := (*C.GMenuModel)(u)
	if c == nil {
		return nil
	}

	g := &MenuModel{native: c}

	return g
}

func (recv *MenuModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *MenuModel) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to MenuModel.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuModel.
func CastToMenuModel(object *gobject.Object) *MenuModel {
	return MenuModelNewFromC(object.ToC())
}

// Unsupported signal 'items-changed' for MenuModel : unsupported parameter position : type gint :

// Unsupported : g_menu_model_get_item_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_model_get_item_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// Queries the item at position @item_index in @model for the link
// specified by @link.
//
// If the link exists, the linked #GMenuModel is returned.  If the link
// does not exist, %NULL is returned.
/*

C function : g_menu_model_get_item_link
*/
func (recv *MenuModel) GetItemLink(itemIndex int32, link string) *MenuModel {
	c_item_index := (C.gint)(itemIndex)

	c_link := C.CString(link)
	defer C.free(unsafe.Pointer(c_link))

	retC := C.g_menu_model_get_item_link((*C.GMenuModel)(recv.native), c_item_index, c_link)
	retGo := MenuModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Query the number of items in @model.
/*

C function : g_menu_model_get_n_items
*/
func (recv *MenuModel) GetNItems() int32 {
	retC := C.g_menu_model_get_n_items((*C.GMenuModel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Queries if @model is mutable.
//
// An immutable #GMenuModel will never emit the #GMenuModel::items-changed
// signal. Consumers of the model may make optimisations accordingly.
/*

C function : g_menu_model_is_mutable
*/
func (recv *MenuModel) IsMutable() bool {
	retC := C.g_menu_model_is_mutable((*C.GMenuModel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Requests emission of the #GMenuModel::items-changed signal on @model.
//
// This function should never be called except by #GMenuModel
// subclasses.  Any other calls to this function will very likely lead
// to a violation of the interface of the model.
//
// The implementation should update its internal representation of the
// menu before emitting the signal.  The implementation should further
// expect to receive queries about the new state of the menu (and
// particularly added menu items) while signal handlers are running.
//
// The implementation must dispatch this call directly from a mainloop
// entry and not in response to calls -- particularly those from the
// #GMenuModel API.  Said another way: the menu must not change while
// user code is running without returning to the mainloop.
/*

C function : g_menu_model_items_changed
*/
func (recv *MenuModel) ItemsChanged(position int32, removed int32, added int32) {
	c_position := (C.gint)(position)

	c_removed := (C.gint)(removed)

	c_added := (C.gint)(added)

	C.g_menu_model_items_changed((*C.GMenuModel)(recv.native), c_position, c_removed, c_added)

	return
}

// Creates a #GMenuAttributeIter to iterate over the attributes of
// the item at position @item_index in @model.
//
// You must free the iterator with g_object_unref() when you are done.
/*

C function : g_menu_model_iterate_item_attributes
*/
func (recv *MenuModel) IterateItemAttributes(itemIndex int32) *MenuAttributeIter {
	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_model_iterate_item_attributes((*C.GMenuModel)(recv.native), c_item_index)
	retGo := MenuAttributeIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GMenuLinkIter to iterate over the links of the item at
// position @item_index in @model.
//
// You must free the iterator with g_object_unref() when you are done.
/*

C function : g_menu_model_iterate_item_links
*/
func (recv *MenuModel) IterateItemLinks(itemIndex int32) *MenuLinkIter {
	c_item_index := (C.gint)(itemIndex)

	retC := C.g_menu_model_iterate_item_links((*C.GMenuModel)(recv.native), c_item_index)
	retGo := MenuLinkIterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GSettings object with a given schema, backend and
// path.
//
// It should be extremely rare that you ever want to use this function.
// It is made available for advanced use-cases (such as plugin systems
// that want to provide access to schemas loaded from custom locations,
// etc).
//
// At the most basic level, a #GSettings object is a pure composition of
// 4 things: a #GSettingsSchema, a #GSettingsBackend, a path within that
// backend, and a #GMainContext to which signals are dispatched.
//
// This constructor therefore gives you full control over constructing
// #GSettings instances.  The first 3 parameters are given directly as
// @schema, @backend and @path, and the main context is taken from the
// thread-default (as per g_settings_new()).
//
// If @backend is %NULL then the default backend is used.
//
// If @path is %NULL then the path from the schema is used.  It is an
// error if @path is %NULL and the schema has no path of its own or if
// @path is non-%NULL and not equal to the path that the schema does
// have.
/*

C function : g_settings_new_full
*/
func SettingsNewFull(schema *SettingsSchema, backend *SettingsBackend, path string) *Settings {
	c_schema := (*C.GSettingsSchema)(C.NULL)
	if schema != nil {
		c_schema = (*C.GSettingsSchema)(schema.ToC())
	}

	c_backend := (*C.GSettingsBackend)(C.NULL)
	if backend != nil {
		c_backend = (*C.GSettingsBackend)(backend.ToC())
	}

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_settings_new_full(c_schema, c_backend, c_path)
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a #GAction corresponding to a given #GSettings key.
//
// The action has the same name as the key.
//
// The value of the key becomes the state of the action and the action
// is enabled when the key is writable.  Changing the state of the
// action results in the key being written to.  Changes to the value or
// writability of the key cause appropriate change notifications to be
// emitted for the action.
//
// For boolean-valued keys, action activations take no parameter and
// result in the toggling of the value.  For all other types,
// activations take the new value for the key (which must have the
// correct type).
/*

C function : g_settings_create_action
*/
func (recv *Settings) CreateAction(key string) *Action {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_create_action((*C.GSettings)(recv.native), c_key)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Sets a #GCancellable to check before dispatching results.
//
// This function has one very specific purpose: the provided cancellable
// is checked at the time of g_simple_async_result_propagate_error() If
// it is cancelled, these functions will return an "Operation was
// cancelled" error (%G_IO_ERROR_CANCELLED).
//
// Implementors of cancellable asynchronous functions should use this in
// order to provide a guarantee to their callers that cancelling an
// async operation will reliably result in an error being returned for
// that operation (even if a positive result for the operation has
// already been sent as an idle to the main context to be dispatched).
//
// The checking described above is done regardless of any call to the
// unrelated g_simple_async_result_set_handle_cancellation() function.
/*

C function : g_simple_async_result_set_check_cancellable
*/
func (recv *SimpleAsyncResult) SetCheckCancellable(checkCancellable *Cancellable) {
	c_check_cancellable := (*C.GCancellable)(C.NULL)
	if checkCancellable != nil {
		c_check_cancellable = (*C.GCancellable)(checkCancellable.ToC())
	}

	C.g_simple_async_result_set_check_cancellable((*C.GSimpleAsyncResult)(recv.native), c_check_cancellable)

	return
}

// Waits for up to @timeout microseconds for @condition to become true
// on @socket. If the condition is met, %TRUE is returned.
//
// If @cancellable is cancelled before the condition is met, or if
// @timeout (or the socket's #GSocket:timeout) is reached before the
// condition is met, then %FALSE is returned and @error, if non-%NULL,
// is set to the appropriate value (%G_IO_ERROR_CANCELLED or
// %G_IO_ERROR_TIMED_OUT).
//
// If you don't want a timeout, use g_socket_condition_wait().
// (Alternatively, you can pass -1 for @timeout.)
//
// Note that although @timeout is in microseconds for consistency with
// other GLib APIs, this function actually only has millisecond
// resolution, and the behavior is undefined if @timeout is not an
// exact number of milliseconds.
/*

C function : g_socket_condition_timed_wait
*/
func (recv *Socket) ConditionTimedWait(condition glib.IOCondition, timeout int64, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_timeout := (C.gint64)(timeout)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_condition_timed_wait((*C.GSocket)(recv.native), c_condition, c_timeout, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Get the amount of data pending in the OS input buffer, without blocking.
//
// If @socket is a UDP or SCTP socket, this will return the size of
// just the next packet, even if additional packets are buffered after
// that one.
//
// Note that on Windows, this function is rather inefficient in the
// UDP case, and so if you know any plausible upper bound on the size
// of the incoming packet, it is better to just do a
// g_socket_receive() with a buffer of that size, rather than calling
// g_socket_get_available_bytes() first and then doing a receive of
// exactly the right size.
/*

C function : g_socket_get_available_bytes
*/
func (recv *Socket) GetAvailableBytes() int64 {
	retC := C.g_socket_get_available_bytes((*C.GSocket)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Gets the broadcast setting on @socket; if %TRUE,
// it is possible to send packets to broadcast
// addresses.
/*

C function : g_socket_get_broadcast
*/
func (recv *Socket) GetBroadcast() bool {
	retC := C.g_socket_get_broadcast((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the multicast loopback setting on @socket; if %TRUE (the
// default), outgoing multicast packets will be looped back to
// multicast listeners on the same host.
/*

C function : g_socket_get_multicast_loopback
*/
func (recv *Socket) GetMulticastLoopback() bool {
	retC := C.g_socket_get_multicast_loopback((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the multicast time-to-live setting on @socket; see
// g_socket_set_multicast_ttl() for more details.
/*

C function : g_socket_get_multicast_ttl
*/
func (recv *Socket) GetMulticastTtl() uint32 {
	retC := C.g_socket_get_multicast_ttl((*C.GSocket)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the unicast time-to-live setting on @socket; see
// g_socket_set_ttl() for more details.
/*

C function : g_socket_get_ttl
*/
func (recv *Socket) GetTtl() uint32 {
	retC := C.g_socket_get_ttl((*C.GSocket)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Registers @socket to receive multicast messages sent to @group.
// @socket must be a %G_SOCKET_TYPE_DATAGRAM socket, and must have
// been bound to an appropriate interface and port with
// g_socket_bind().
//
// If @iface is %NULL, the system will automatically pick an interface
// to bind to based on @group.
//
// If @source_specific is %TRUE, source-specific multicast as defined
// in RFC 4604 is used. Note that on older platforms this may fail
// with a %G_IO_ERROR_NOT_SUPPORTED error.
//
// To bind to a given source-specific multicast address, use
// g_socket_join_multicast_group_ssm() instead.
/*

C function : g_socket_join_multicast_group
*/
func (recv *Socket) JoinMulticastGroup(group *InetAddress, sourceSpecific bool, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(C.NULL)
	if group != nil {
		c_group = (*C.GInetAddress)(group.ToC())
	}

	c_source_specific :=
		boolToGboolean(sourceSpecific)

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_join_multicast_group((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Removes @socket from the multicast group defined by @group, @iface,
// and @source_specific (which must all have the same values they had
// when you joined the group).
//
// @socket remains bound to its address and port, and can still receive
// unicast messages after calling this.
//
// To unbind to a given source-specific multicast address, use
// g_socket_leave_multicast_group_ssm() instead.
/*

C function : g_socket_leave_multicast_group
*/
func (recv *Socket) LeaveMulticastGroup(group *InetAddress, sourceSpecific bool, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(C.NULL)
	if group != nil {
		c_group = (*C.GInetAddress)(group.ToC())
	}

	c_source_specific :=
		boolToGboolean(sourceSpecific)

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_leave_multicast_group((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets whether @socket should allow sending to broadcast addresses.
// This is %FALSE by default.
/*

C function : g_socket_set_broadcast
*/
func (recv *Socket) SetBroadcast(broadcast bool) {
	c_broadcast :=
		boolToGboolean(broadcast)

	C.g_socket_set_broadcast((*C.GSocket)(recv.native), c_broadcast)

	return
}

// Sets whether outgoing multicast packets will be received by sockets
// listening on that multicast address on the same host. This is %TRUE
// by default.
/*

C function : g_socket_set_multicast_loopback
*/
func (recv *Socket) SetMulticastLoopback(loopback bool) {
	c_loopback :=
		boolToGboolean(loopback)

	C.g_socket_set_multicast_loopback((*C.GSocket)(recv.native), c_loopback)

	return
}

// Sets the time-to-live for outgoing multicast datagrams on @socket.
// By default, this is 1, meaning that multicast packets will not leave
// the local network.
/*

C function : g_socket_set_multicast_ttl
*/
func (recv *Socket) SetMulticastTtl(ttl uint32) {
	c_ttl := (C.guint)(ttl)

	C.g_socket_set_multicast_ttl((*C.GSocket)(recv.native), c_ttl)

	return
}

// Sets the time-to-live for outgoing unicast packets on @socket.
// By default the platform-specific default value is used.
/*

C function : g_socket_set_ttl
*/
func (recv *Socket) SetTtl(ttl uint32) {
	c_ttl := (C.guint)(ttl)

	C.g_socket_set_ttl((*C.GSocket)(recv.native), c_ttl)

	return
}

// Connect @connection to the specified remote address.
/*

C function : g_socket_connection_connect
*/
func (recv *SocketConnection) Connect(address *SocketAddress, cancellable *Cancellable) (bool, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_connection_connect((*C.GSocketConnection)(recv.native), c_address, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_socket_connection_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Gets the result of a g_socket_connection_connect_async() call.
/*

C function : g_socket_connection_connect_finish
*/
func (recv *SocketConnection) ConnectFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_connection_connect_finish((*C.GSocketConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks if @connection is connected. This is equivalent to calling
// g_socket_is_connected() on @connection's underlying #GSocket.
/*

C function : g_socket_connection_is_connected
*/
func (recv *SocketConnection) IsConnected() bool {
	retC := C.g_socket_connection_is_connected((*C.GSocketConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_unix_connection_receive_credentials_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous receive credentials operation started with
// g_unix_connection_receive_credentials_async().
/*

C function : g_unix_connection_receive_credentials_finish
*/
func (recv *UnixConnection) ReceiveCredentialsFinish(result *AsyncResult) (*Credentials, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_unix_connection_receive_credentials_finish((*C.GUnixConnection)(recv.native), c_result, &cThrowableError)
	retGo := CredentialsNewFromC(unsafe.Pointer(retC))

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Unsupported : g_unix_connection_send_credentials_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Finishes an asynchronous send credentials operation started with
// g_unix_connection_send_credentials_async().
/*

C function : g_unix_connection_send_credentials_finish
*/
func (recv *UnixConnection) SendCredentialsFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_unix_connection_send_credentials_finish((*C.GUnixConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}
