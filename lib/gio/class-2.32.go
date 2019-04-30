// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// InetAddressMask is a wrapper around the C record GInetAddressMask.
type InetAddressMask struct {
	native unsafe.Pointer
	// parent_instance : record
	// Private : priv
}

func InetAddressMaskNewFromC(u unsafe.Pointer) *InetAddressMask {
	if u == nil {
		return nil
	}

	g := &InetAddressMask{native: u}

	return g
}

func (recv *InetAddressMask) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_inet_address_mask_new : return type :

// Unsupported : g_inet_address_mask_new_from_string : return type :

// Menu is a wrapper around the C record GMenu.
type Menu struct {
	native unsafe.Pointer
}

func MenuNewFromC(u unsafe.Pointer) *Menu {
	if u == nil {
		return nil
	}

	g := &Menu{native: u}

	return g
}

func (recv *Menu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_menu_new : return type :

// MenuAttributeIter is a wrapper around the C record GMenuAttributeIter.
type MenuAttributeIter struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func MenuAttributeIterNewFromC(u unsafe.Pointer) *MenuAttributeIter {
	if u == nil {
		return nil
	}

	g := &MenuAttributeIter{native: u}

	return g
}

func (recv *MenuAttributeIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItem is a wrapper around the C record GMenuItem.
type MenuItem struct {
	native unsafe.Pointer
}

func MenuItemNewFromC(u unsafe.Pointer) *MenuItem {
	if u == nil {
		return nil
	}

	g := &MenuItem{native: u}

	return g
}

func (recv *MenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_menu_item_new : return type :

// Unsupported : g_menu_item_new_section : return type :

// Unsupported : g_menu_item_new_submenu : return type :

// MenuLinkIter is a wrapper around the C record GMenuLinkIter.
type MenuLinkIter struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func MenuLinkIterNewFromC(u unsafe.Pointer) *MenuLinkIter {
	if u == nil {
		return nil
	}

	g := &MenuLinkIter{native: u}

	return g
}

func (recv *MenuLinkIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuModel is a wrapper around the C record GMenuModel.
type MenuModel struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func MenuModelNewFromC(u unsafe.Pointer) *MenuModel {
	if u == nil {
		return nil
	}

	g := &MenuModel{native: u}

	return g
}

func (recv *MenuModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_settings_new_full : return type :
