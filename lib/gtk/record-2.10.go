// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Blacklisted : gtk_paper_size_new

// Blacklisted : gtk_paper_size_new_custom

// Blacklisted : gtk_paper_size_new_from_ppd

// Blacklisted : gtk_paper_size_get_default

// Blacklisted : gtk_paper_size_copy

// Blacklisted : gtk_paper_size_free

// Blacklisted : gtk_paper_size_get_default_bottom_margin

// Blacklisted : gtk_paper_size_get_default_left_margin

// Blacklisted : gtk_paper_size_get_default_right_margin

// Blacklisted : gtk_paper_size_get_default_top_margin

// Blacklisted : gtk_paper_size_get_display_name

// Blacklisted : gtk_paper_size_get_height

// Blacklisted : gtk_paper_size_get_name

// Blacklisted : gtk_paper_size_get_ppd_name

// Blacklisted : gtk_paper_size_get_width

// Blacklisted : gtk_paper_size_is_equal

// Blacklisted : gtk_paper_size_set_size

// RecentInfo is a wrapper around the C record GtkRecentInfo.
type RecentInfo struct {
	native *C.GtkRecentInfo
}

func RecentInfoNewFromC(u unsafe.Pointer) *RecentInfo {
	c := (*C.GtkRecentInfo)(u)
	if c == nil {
		return nil
	}

	g := &RecentInfo{native: c}

	return g
}

func (recv *RecentInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RecentInfo with another RecentInfo, and returns true if they represent the same GObject.
func (recv *RecentInfo) Equals(other *RecentInfo) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : gtk_recent_info_create_app_info

// Blacklisted : gtk_recent_info_exists

// Blacklisted : gtk_recent_info_get_added

// Blacklisted : gtk_recent_info_get_age

// Blacklisted : gtk_recent_info_get_application_info

// Blacklisted : gtk_recent_info_get_applications

// Blacklisted : gtk_recent_info_get_description

// Blacklisted : gtk_recent_info_get_display_name

// Blacklisted : gtk_recent_info_get_groups

// Blacklisted : gtk_recent_info_get_icon

// Blacklisted : gtk_recent_info_get_mime_type

// Blacklisted : gtk_recent_info_get_modified

// Blacklisted : gtk_recent_info_get_private_hint

// Blacklisted : gtk_recent_info_get_short_name

// Blacklisted : gtk_recent_info_get_uri

// Blacklisted : gtk_recent_info_get_uri_display

// Blacklisted : gtk_recent_info_get_visited

// Blacklisted : gtk_recent_info_has_application

// Blacklisted : gtk_recent_info_has_group

// Blacklisted : gtk_recent_info_is_local

// Blacklisted : gtk_recent_info_last_application

// Blacklisted : gtk_recent_info_match

// Blacklisted : gtk_recent_info_ref

// Blacklisted : gtk_recent_info_unref

// RecentManagerClass is a wrapper around the C record GtkRecentManagerClass.
type RecentManagerClass struct {
	native *C.GtkRecentManagerClass
	// Private : parent_class
	// no type for changed
	// no type for _gtk_recent1
	// no type for _gtk_recent2
	// no type for _gtk_recent3
	// no type for _gtk_recent4
}

func RecentManagerClassNewFromC(u unsafe.Pointer) *RecentManagerClass {
	c := (*C.GtkRecentManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &RecentManagerClass{native: c}

	return g
}

func (recv *RecentManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this RecentManagerClass with another RecentManagerClass, and returns true if they represent the same GObject.
func (recv *RecentManagerClass) Equals(other *RecentManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// Blacklisted : gtk_selection_data_targets_include_rich_text

// Blacklisted : gtk_selection_data_targets_include_uri

// Blacklisted : gtk_target_list_add_rich_text_targets
