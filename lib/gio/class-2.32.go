// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

// InetAddressMask is a wrapper around the C record GInetAddressMask.
type InetAddressMask struct {
	native *C.GInetAddressMask
	// parent_instance : no type generator for GObject.Object, GObject
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

func (recv *InetAddressMask) toC() *C.GInetAddressMask {

	return recv.native
}

// Unsupported : g_inet_address_mask_new : unsupported parameter addr : no type generator for InetAddress, GInetAddress*

// Unsupported : g_inet_address_mask_new_from_string : no return generator

// Unsupported : g_inet_address_mask_equal : unsupported parameter mask2 : no type generator for InetAddressMask, GInetAddressMask*

// Unsupported : g_inet_address_mask_get_address : no return generator

// GetFamily is a wrapper around the C function g_inet_address_mask_get_family.
func (recv *InetAddressMask) GetFamily() SocketFamily {
	retC := C.g_inet_address_mask_get_family((*C.GInetAddressMask)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetLength is a wrapper around the C function g_inet_address_mask_get_length.
func (recv *InetAddressMask) GetLength() uint32 {
	retC := C.g_inet_address_mask_get_length((*C.GInetAddressMask)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Unsupported : g_inet_address_mask_matches : unsupported parameter address : no type generator for InetAddress, GInetAddress*

// ToString is a wrapper around the C function g_inet_address_mask_to_string.
func (recv *InetAddressMask) ToString() string {
	retC := C.g_inet_address_mask_to_string((*C.GInetAddressMask)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

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

func (recv *Menu) toC() *C.GMenu {

	return recv.native
}

// Unsupported : g_menu_new : no return generator

// Unsupported : g_menu_append : no return generator

// Unsupported : g_menu_append_item : unsupported parameter item : no type generator for MenuItem, GMenuItem*

// Unsupported : g_menu_append_section : unsupported parameter section : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_append_submenu : unsupported parameter submenu : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_freeze : no return generator

// Unsupported : g_menu_insert : no return generator

// Unsupported : g_menu_insert_item : unsupported parameter item : no type generator for MenuItem, GMenuItem*

// Unsupported : g_menu_insert_section : unsupported parameter section : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_insert_submenu : unsupported parameter submenu : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_prepend : no return generator

// Unsupported : g_menu_prepend_item : unsupported parameter item : no type generator for MenuItem, GMenuItem*

// Unsupported : g_menu_prepend_section : unsupported parameter section : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_prepend_submenu : unsupported parameter submenu : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_remove : no return generator

// Unsupported : g_menu_remove_all : no return generator

// MenuAttributeIter is a wrapper around the C record GMenuAttributeIter.
type MenuAttributeIter struct {
	native *C.GMenuAttributeIter
	// parent_instance : no type generator for GObject.Object, GObject
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

func (recv *MenuAttributeIter) toC() *C.GMenuAttributeIter {

	return recv.native
}

// GetName is a wrapper around the C function g_menu_attribute_iter_get_name.
func (recv *MenuAttributeIter) GetName() string {
	retC := C.g_menu_attribute_iter_get_name((*C.GMenuAttributeIter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_menu_attribute_iter_get_next : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_menu_attribute_iter_get_value : return type : Blacklisted record : GVariant

// Unsupported : g_menu_attribute_iter_next : no return generator

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

func (recv *MenuItem) toC() *C.GMenuItem {

	return recv.native
}

// Unsupported : g_menu_item_new : no return generator

// Unsupported : g_menu_item_new_from_model : unsupported parameter model : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_item_new_section : unsupported parameter section : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_item_new_submenu : unsupported parameter submenu : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_item_get_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_get_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// Unsupported : g_menu_item_get_link : no return generator

// Unsupported : g_menu_item_set_action_and_target : unsupported parameter ... : varargs

// Unsupported : g_menu_item_set_action_and_target_value : unsupported parameter target_value : Blacklisted record : GVariant

// Unsupported : g_menu_item_set_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_item_set_attribute_value : unsupported parameter value : Blacklisted record : GVariant

// Unsupported : g_menu_item_set_detailed_action : no return generator

// Unsupported : g_menu_item_set_icon : unsupported parameter icon : no type generator for Icon, GIcon*

// Unsupported : g_menu_item_set_label : no return generator

// Unsupported : g_menu_item_set_link : unsupported parameter model : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_item_set_section : unsupported parameter section : no type generator for MenuModel, GMenuModel*

// Unsupported : g_menu_item_set_submenu : unsupported parameter submenu : no type generator for MenuModel, GMenuModel*

// MenuLinkIter is a wrapper around the C record GMenuLinkIter.
type MenuLinkIter struct {
	native *C.GMenuLinkIter
	// parent_instance : no type generator for GObject.Object, GObject
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

func (recv *MenuLinkIter) toC() *C.GMenuLinkIter {

	return recv.native
}

// GetName is a wrapper around the C function g_menu_link_iter_get_name.
func (recv *MenuLinkIter) GetName() string {
	retC := C.g_menu_link_iter_get_name((*C.GMenuLinkIter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_menu_link_iter_get_next : unsupported parameter value : no type generator for MenuModel, GMenuModel**

// Unsupported : g_menu_link_iter_get_value : no return generator

// Unsupported : g_menu_link_iter_next : no return generator

// MenuModel is a wrapper around the C record GMenuModel.
type MenuModel struct {
	native *C.GMenuModel
	// parent_instance : no type generator for GObject.Object, GObject
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

func (recv *MenuModel) toC() *C.GMenuModel {

	return recv.native
}

// Unsupported : g_menu_model_get_item_attribute : unsupported parameter ... : varargs

// Unsupported : g_menu_model_get_item_attribute_value : unsupported parameter expected_type : Blacklisted record : GVariantType

// Unsupported : g_menu_model_get_item_link : no return generator

// GetNItems is a wrapper around the C function g_menu_model_get_n_items.
func (recv *MenuModel) GetNItems() int32 {
	retC := C.g_menu_model_get_n_items((*C.GMenuModel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_menu_model_is_mutable : no return generator

// Unsupported : g_menu_model_items_changed : no return generator

// Unsupported : g_menu_model_iterate_item_attributes : no return generator

// Unsupported : g_menu_model_iterate_item_links : no return generator
