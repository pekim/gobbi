// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

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

// A key in the "filesystem" namespace for getting the number of bytes of used on the
// file system. Corresponding #GFileAttributeType is
// %G_FILE_ATTRIBUTE_TYPE_UINT64.
const FILE_ATTRIBUTE_FILESYSTEM_USED string = C.G_FILE_ATTRIBUTE_FILESYSTEM_USED

// The menu item attribute which holds the action name of the item.  Action
// names are namespaced with an identifier for the action group in which the
// action resides. For example, "win." for window-specific actions and "app."
// for application-wide actions.
//
// See also g_menu_model_get_item_attribute() and g_menu_item_set_attribute().
const MENU_ATTRIBUTE_ACTION string = C.G_MENU_ATTRIBUTE_ACTION

// The menu item attribute which holds the label of the item.
const MENU_ATTRIBUTE_LABEL string = C.G_MENU_ATTRIBUTE_LABEL

// The menu item attribute which holds the target with which the item's action
// will be activated.
//
// See also g_menu_item_set_action_and_target()
const MENU_ATTRIBUTE_TARGET string = C.G_MENU_ATTRIBUTE_TARGET

// The name of the link that associates a menu item with a section.  The linked
// menu will usually be shown in place of the menu item, using the item's label
// as a header.
//
// See also g_menu_item_set_link().
const MENU_LINK_SECTION string = C.G_MENU_LINK_SECTION

// The name of the link that associates a menu item with a submenu.
//
// See also g_menu_item_set_link().
const MENU_LINK_SUBMENU string = C.G_MENU_LINK_SUBMENU
