// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Creates a new #GtkPaperSize object by parsing a
// [PWG 5101.1-2002](ftp://ftp.pwg.org/pub/pwg/candidates/cs-pwgmsn10-20020226-5101.1.pdf)
// paper name.
//
// If @name is %NULL, the default paper size is returned,
// see gtk_paper_size_get_default().
/*

C function : gtk_paper_size_new
*/
func PaperSizeNew(name string) *PaperSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_paper_size_new(c_name)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkPaperSize object with the
// given parameters.
/*

C function : gtk_paper_size_new_custom
*/
func PaperSizeNewCustom(name string, displayName string, width float64, height float64, unit Unit) *PaperSize {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_display_name := C.CString(displayName)
	defer C.free(unsafe.Pointer(c_display_name))

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_new_custom(c_name, c_display_name, c_width, c_height, c_unit)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Creates a new #GtkPaperSize object by using
// PPD information.
//
// If @ppd_name is not a recognized PPD paper name,
// @ppd_display_name, @width and @height are used to
// construct a custom #GtkPaperSize object.
/*

C function : gtk_paper_size_new_from_ppd
*/
func PaperSizeNewFromPpd(ppdName string, ppdDisplayName string, width float64, height float64) *PaperSize {
	c_ppd_name := C.CString(ppdName)
	defer C.free(unsafe.Pointer(c_ppd_name))

	c_ppd_display_name := C.CString(ppdDisplayName)
	defer C.free(unsafe.Pointer(c_ppd_display_name))

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	retC := C.gtk_paper_size_new_from_ppd(c_ppd_name, c_ppd_display_name, c_width, c_height)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copies an existing #GtkPaperSize.
/*

C function : gtk_paper_size_copy
*/
func (recv *PaperSize) Copy() *PaperSize {
	retC := C.gtk_paper_size_copy((*C.GtkPaperSize)(recv.native))
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free the given #GtkPaperSize object.
/*

C function : gtk_paper_size_free
*/
func (recv *PaperSize) Free() {
	C.gtk_paper_size_free((*C.GtkPaperSize)(recv.native))

	return
}

// Gets the default bottom margin for the #GtkPaperSize.
/*

C function : gtk_paper_size_get_default_bottom_margin
*/
func (recv *PaperSize) GetDefaultBottomMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_bottom_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// Gets the default left margin for the #GtkPaperSize.
/*

C function : gtk_paper_size_get_default_left_margin
*/
func (recv *PaperSize) GetDefaultLeftMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_left_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// Gets the default right margin for the #GtkPaperSize.
/*

C function : gtk_paper_size_get_default_right_margin
*/
func (recv *PaperSize) GetDefaultRightMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_right_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// Gets the default top margin for the #GtkPaperSize.
/*

C function : gtk_paper_size_get_default_top_margin
*/
func (recv *PaperSize) GetDefaultTopMargin(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_default_top_margin((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// Gets the human-readable name of the #GtkPaperSize.
/*

C function : gtk_paper_size_get_display_name
*/
func (recv *PaperSize) GetDisplayName() string {
	retC := C.gtk_paper_size_get_display_name((*C.GtkPaperSize)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the paper height of the #GtkPaperSize, in
// units of @unit.
/*

C function : gtk_paper_size_get_height
*/
func (recv *PaperSize) GetHeight(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_height((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// Gets the name of the #GtkPaperSize.
/*

C function : gtk_paper_size_get_name
*/
func (recv *PaperSize) GetName() string {
	retC := C.gtk_paper_size_get_name((*C.GtkPaperSize)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the PPD name of the #GtkPaperSize, which
// may be %NULL.
/*

C function : gtk_paper_size_get_ppd_name
*/
func (recv *PaperSize) GetPpdName() string {
	retC := C.gtk_paper_size_get_ppd_name((*C.GtkPaperSize)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the paper width of the #GtkPaperSize, in
// units of @unit.
/*

C function : gtk_paper_size_get_width
*/
func (recv *PaperSize) GetWidth(unit Unit) float64 {
	c_unit := (C.GtkUnit)(unit)

	retC := C.gtk_paper_size_get_width((*C.GtkPaperSize)(recv.native), c_unit)
	retGo := (float64)(retC)

	return retGo
}

// Compares two #GtkPaperSize objects.
/*

C function : gtk_paper_size_is_equal
*/
func (recv *PaperSize) IsEqual(size2 *PaperSize) bool {
	c_size2 := (*C.GtkPaperSize)(C.NULL)
	if size2 != nil {
		c_size2 = (*C.GtkPaperSize)(size2.ToC())
	}

	retC := C.gtk_paper_size_is_equal((*C.GtkPaperSize)(recv.native), c_size2)
	retGo := retC == C.TRUE

	return retGo
}

// Changes the dimensions of a @size to @width x @height.
/*

C function : gtk_paper_size_set_size
*/
func (recv *PaperSize) SetSize(width float64, height float64, unit Unit) {
	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_unit := (C.GtkUnit)(unit)

	C.gtk_paper_size_set_size((*C.GtkPaperSize)(recv.native), c_width, c_height, c_unit)

	return
}

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

// Creates a #GAppInfo for the specified #GtkRecentInfo
/*

C function : gtk_recent_info_create_app_info
*/
func (recv *RecentInfo) CreateAppInfo(appName string) (*gio.AppInfo, error) {
	c_app_name := C.CString(appName)
	defer C.free(unsafe.Pointer(c_app_name))

	var cThrowableError *C.GError

	retC := C.gtk_recent_info_create_app_info((*C.GtkRecentInfo)(recv.native), c_app_name, &cThrowableError)
	var retGo (*gio.AppInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.AppInfoNewFromC(unsafe.Pointer(retC))
	}

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Checks whether the resource pointed by @info still exists.
// At the moment this check is done only on resources pointing
// to local files.
/*

C function : gtk_recent_info_exists
*/
func (recv *RecentInfo) Exists() bool {
	retC := C.gtk_recent_info_exists((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_recent_info_get_added : no return generator

// Gets the number of days elapsed since the last update
// of the resource pointed by @info.
/*

C function : gtk_recent_info_get_age
*/
func (recv *RecentInfo) GetAge() int32 {
	retC := C.gtk_recent_info_get_age((*C.GtkRecentInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_recent_info_get_application_info : unsupported parameter time_ : no type generator for glong (time_t*) for param time_

// Unsupported : gtk_recent_info_get_applications : no return type

// Gets the (short) description of the resource.
/*

C function : gtk_recent_info_get_description
*/
func (recv *RecentInfo) GetDescription() string {
	retC := C.gtk_recent_info_get_description((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets the name of the resource. If none has been defined, the basename
// of the resource is obtained.
/*

C function : gtk_recent_info_get_display_name
*/
func (recv *RecentInfo) GetDisplayName() string {
	retC := C.gtk_recent_info_get_display_name((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_recent_info_get_groups : no return type

// Retrieves the icon of size @size associated to the resource MIME type.
/*

C function : gtk_recent_info_get_icon
*/
func (recv *RecentInfo) GetIcon(size int32) *gdkpixbuf.Pixbuf {
	c_size := (C.gint)(size)

	retC := C.gtk_recent_info_get_icon((*C.GtkRecentInfo)(recv.native), c_size)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Gets the MIME type of the resource.
/*

C function : gtk_recent_info_get_mime_type
*/
func (recv *RecentInfo) GetMimeType() string {
	retC := C.gtk_recent_info_get_mime_type((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : gtk_recent_info_get_modified : no return generator

// Gets the value of the “private” flag. Resources in the recently used
// list that have this flag set to %TRUE should only be displayed by the
// applications that have registered them.
/*

C function : gtk_recent_info_get_private_hint
*/
func (recv *RecentInfo) GetPrivateHint() bool {
	retC := C.gtk_recent_info_get_private_hint((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Computes a valid UTF-8 string that can be used as the
// name of the item in a menu or list. For example, calling
// this function on an item that refers to
// “file:///foo/bar.txt” will yield “bar.txt”.
/*

C function : gtk_recent_info_get_short_name
*/
func (recv *RecentInfo) GetShortName() string {
	retC := C.gtk_recent_info_get_short_name((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Gets the URI of the resource.
/*

C function : gtk_recent_info_get_uri
*/
func (recv *RecentInfo) GetUri() string {
	retC := C.gtk_recent_info_get_uri((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Gets a displayable version of the resource’s URI. If the resource
// is local, it returns a local path; if the resource is not local,
// it returns the UTF-8 encoded content of gtk_recent_info_get_uri().
/*

C function : gtk_recent_info_get_uri_display
*/
func (recv *RecentInfo) GetUriDisplay() string {
	retC := C.gtk_recent_info_get_uri_display((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_recent_info_get_visited : no return generator

// Checks whether an application registered this resource using @app_name.
/*

C function : gtk_recent_info_has_application
*/
func (recv *RecentInfo) HasApplication(appName string) bool {
	c_app_name := C.CString(appName)
	defer C.free(unsafe.Pointer(c_app_name))

	retC := C.gtk_recent_info_has_application((*C.GtkRecentInfo)(recv.native), c_app_name)
	retGo := retC == C.TRUE

	return retGo
}

// Checks whether @group_name appears inside the groups
// registered for the recently used item @info.
/*

C function : gtk_recent_info_has_group
*/
func (recv *RecentInfo) HasGroup(groupName string) bool {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	retC := C.gtk_recent_info_has_group((*C.GtkRecentInfo)(recv.native), c_group_name)
	retGo := retC == C.TRUE

	return retGo
}

// Checks whether the resource is local or not by looking at the
// scheme of its URI.
/*

C function : gtk_recent_info_is_local
*/
func (recv *RecentInfo) IsLocal() bool {
	retC := C.gtk_recent_info_is_local((*C.GtkRecentInfo)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the name of the last application that have registered the
// recently used resource represented by @info.
/*

C function : gtk_recent_info_last_application
*/
func (recv *RecentInfo) LastApplication() string {
	retC := C.gtk_recent_info_last_application((*C.GtkRecentInfo)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Checks whether two #GtkRecentInfo-struct point to the same
// resource.
/*

C function : gtk_recent_info_match
*/
func (recv *RecentInfo) Match(infoB *RecentInfo) bool {
	c_info_b := (*C.GtkRecentInfo)(C.NULL)
	if infoB != nil {
		c_info_b = (*C.GtkRecentInfo)(infoB.ToC())
	}

	retC := C.gtk_recent_info_match((*C.GtkRecentInfo)(recv.native), c_info_b)
	retGo := retC == C.TRUE

	return retGo
}

// Increases the reference count of @recent_info by one.
/*

C function : gtk_recent_info_ref
*/
func (recv *RecentInfo) Ref() *RecentInfo {
	retC := C.gtk_recent_info_ref((*C.GtkRecentInfo)(recv.native))
	retGo := RecentInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Decreases the reference count of @info by one. If the reference
// count reaches zero, @info is deallocated, and the memory freed.
/*

C function : gtk_recent_info_unref
*/
func (recv *RecentInfo) Unref() {
	C.gtk_recent_info_unref((*C.GtkRecentInfo)(recv.native))

	return
}

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

// Given a #GtkSelectionData object holding a list of targets,
// determines if any of the targets in @targets can be used to
// provide rich text.
/*

C function : gtk_selection_data_targets_include_rich_text
*/
func (recv *SelectionData) TargetsIncludeRichText(buffer *TextBuffer) bool {
	c_buffer := (*C.GtkTextBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkTextBuffer)(buffer.ToC())
	}

	retC := C.gtk_selection_data_targets_include_rich_text((*C.GtkSelectionData)(recv.native), c_buffer)
	retGo := retC == C.TRUE

	return retGo
}

// Given a #GtkSelectionData object holding a list of targets,
// determines if any of the targets in @targets can be used to
// provide a list or URIs.
/*

C function : gtk_selection_data_targets_include_uri
*/
func (recv *SelectionData) TargetsIncludeUri() bool {
	retC := C.gtk_selection_data_targets_include_uri((*C.GtkSelectionData)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Appends the rich text targets registered with
// gtk_text_buffer_register_serialize_format() or
// gtk_text_buffer_register_deserialize_format() to the target list. All
// targets are added with the same @info.
/*

C function : gtk_target_list_add_rich_text_targets
*/
func (recv *TargetList) AddRichTextTargets(info uint32, deserializable bool, buffer *TextBuffer) {
	c_info := (C.guint)(info)

	c_deserializable :=
		boolToGboolean(deserializable)

	c_buffer := (*C.GtkTextBuffer)(C.NULL)
	if buffer != nil {
		c_buffer = (*C.GtkTextBuffer)(buffer.ToC())
	}

	C.gtk_target_list_add_rich_text_targets((*C.GtkTargetList)(recv.native), c_info, c_deserializable, c_buffer)

	return
}
