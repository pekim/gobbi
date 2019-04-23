// This is a generated file - DO NOT EDIT

package gio

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
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

// Action is a wrapper around the C record GAction.
type Action struct {
	native *C.GAction
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.GAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Action with another Action, and returns true if they represent the same GObject.
func (recv *Action) Equals(other *Action) bool {
	return other.ToC() == recv.ToC()
}

// ActionGroup is a wrapper around the C record GActionGroup.
type ActionGroup struct {
	native *C.GActionGroup
}

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	c := (*C.GActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroup{native: c}

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionGroup with another ActionGroup, and returns true if they represent the same GObject.
func (recv *ActionGroup) Equals(other *ActionGroup) bool {
	return other.ToC() == recv.ToC()
}

// ActionMap is a wrapper around the C record GActionMap.
type ActionMap struct {
	native *C.GActionMap
}

func ActionMapNewFromC(u unsafe.Pointer) *ActionMap {
	c := (*C.GActionMap)(u)
	if c == nil {
		return nil
	}

	g := &ActionMap{native: c}

	return g
}

func (recv *ActionMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ActionMap with another ActionMap, and returns true if they represent the same GObject.
func (recv *ActionMap) Equals(other *ActionMap) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : GAppInfo

// Blacklisted : GAsyncResult

// Blacklisted : GDBusObject

// Blacklisted : GDBusObjectManager

// Blacklisted : GDesktopAppInfoLookup

// Blacklisted : GDrive

// File is a wrapper around the C record GFile.
type File struct {
	native *C.GFile
}

func FileNewFromC(u unsafe.Pointer) *File {
	c := (*C.GFile)(u)
	if c == nil {
		return nil
	}

	g := &File{native: c}

	return g
}

func (recv *File) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this File with another File, and returns true if they represent the same GObject.
func (recv *File) Equals(other *File) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_file_new_for_commandline_arg

// Blacklisted : g_file_new_for_path

// Blacklisted : g_file_new_for_uri

// Blacklisted : g_file_parse_name

// Blacklisted : g_file_append_to

// Unsupported : g_file_append_to_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_append_to_finish

// Unsupported : g_file_copy : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Unsupported : g_file_copy_async : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Blacklisted : g_file_copy_attributes

// Blacklisted : g_file_copy_finish

// Blacklisted : g_file_create

// Unsupported : g_file_create_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_create_finish

// Blacklisted : g_file_delete

// Blacklisted : g_file_dup

// Unsupported : g_file_eject_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_eject_mountable_finish

// Blacklisted : g_file_enumerate_children

// Unsupported : g_file_enumerate_children_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_enumerate_children_finish

// Blacklisted : g_file_equal

// Blacklisted : g_file_find_enclosing_mount

// Unsupported : g_file_find_enclosing_mount_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_find_enclosing_mount_finish

// Blacklisted : g_file_get_basename

// Blacklisted : g_file_get_child

// Blacklisted : g_file_get_child_for_display_name

// Blacklisted : g_file_get_parent

// Blacklisted : g_file_get_parse_name

// Blacklisted : g_file_get_path

// Blacklisted : g_file_get_relative_path

// Blacklisted : g_file_get_uri

// Blacklisted : g_file_get_uri_scheme

// Blacklisted : g_file_has_prefix

// Blacklisted : g_file_has_uri_scheme

// Blacklisted : g_file_hash

// Blacklisted : g_file_is_native

// Unsupported : g_file_load_contents : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_file_load_contents_finish : unsupported parameter contents : output array param contents

// Unsupported : g_file_load_partial_contents_async : unsupported parameter read_more_callback : no type generator for FileReadMoreCallback (GFileReadMoreCallback) for param read_more_callback

// Unsupported : g_file_load_partial_contents_finish : unsupported parameter contents : output array param contents

// Blacklisted : g_file_make_directory

// Blacklisted : g_file_make_symbolic_link

// Blacklisted : g_file_monitor_directory

// Blacklisted : g_file_monitor_file

// Unsupported : g_file_mount_enclosing_volume : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_mount_enclosing_volume_finish

// Unsupported : g_file_mount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_mount_mountable_finish

// Unsupported : g_file_move : unsupported parameter progress_callback : no type generator for FileProgressCallback (GFileProgressCallback) for param progress_callback

// Blacklisted : g_file_query_default_handler

// Blacklisted : g_file_query_exists

// Blacklisted : g_file_query_filesystem_info

// Unsupported : g_file_query_filesystem_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_query_filesystem_info_finish

// Blacklisted : g_file_query_info

// Unsupported : g_file_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_query_info_finish

// Blacklisted : g_file_query_settable_attributes

// Blacklisted : g_file_query_writable_namespaces

// Blacklisted : g_file_read

// Unsupported : g_file_read_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_read_finish

// Blacklisted : g_file_replace

// Unsupported : g_file_replace_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_replace_contents

// Unsupported : g_file_replace_contents_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_replace_contents_finish

// Blacklisted : g_file_replace_finish

// Blacklisted : g_file_resolve_relative_path

// Unsupported : g_file_set_attribute : unsupported parameter value_p : no type generator for gpointer (gpointer) for param value_p

// Blacklisted : g_file_set_attribute_byte_string

// Blacklisted : g_file_set_attribute_int32

// Blacklisted : g_file_set_attribute_int64

// Blacklisted : g_file_set_attribute_string

// Blacklisted : g_file_set_attribute_uint32

// Blacklisted : g_file_set_attribute_uint64

// Unsupported : g_file_set_attributes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_set_attributes_finish

// Blacklisted : g_file_set_attributes_from_info

// Blacklisted : g_file_set_display_name

// Unsupported : g_file_set_display_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_set_display_name_finish

// Blacklisted : g_file_trash

// Unsupported : g_file_unmount_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_file_unmount_mountable_finish

// Blacklisted : GFileDescriptorBased

// Icon is a wrapper around the C record GIcon.
type Icon struct {
	native *C.GIcon
}

func IconNewFromC(u unsafe.Pointer) *Icon {
	c := (*C.GIcon)(u)
	if c == nil {
		return nil
	}

	g := &Icon{native: c}

	return g
}

func (recv *Icon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Icon with another Icon, and returns true if they represent the same GObject.
func (recv *Icon) Equals(other *Icon) bool {
	return other.ToC() == recv.ToC()
}

// g_icon_hash : unsupported parameter icon : no type generator for gpointer (gconstpointer) for param icon
// Blacklisted : g_icon_equal

// Blacklisted : GListModel

// LoadableIcon is a wrapper around the C record GLoadableIcon.
type LoadableIcon struct {
	native *C.GLoadableIcon
}

func LoadableIconNewFromC(u unsafe.Pointer) *LoadableIcon {
	c := (*C.GLoadableIcon)(u)
	if c == nil {
		return nil
	}

	g := &LoadableIcon{native: c}

	return g
}

func (recv *LoadableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this LoadableIcon with another LoadableIcon, and returns true if they represent the same GObject.
func (recv *LoadableIcon) Equals(other *LoadableIcon) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : g_loadable_icon_load

// Unsupported : g_loadable_icon_load_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Blacklisted : g_loadable_icon_load_finish

// Blacklisted : GMount

// Blacklisted : GRemoteActionGroup

// Blacklisted : GSeekable

// Blacklisted : GSocketConnectable

// Blacklisted : GVolume
