// Code generated - DO NOT EDIT.
// +build !webkit2_2.2,!webkit2_2.4,!webkit2_2.6,!webkit2_2.8,!webkit2_2.10,!webkit2_2.12,!webkit2_2.14,!webkit2_2.16,!webkit2_2.18,!webkit2_2.20,!webkit2_2.22,!webkit2_2.24,!webkit2_2.26

package webkit2

import (
	atk "github.com/pekim/gobbi/lib/atk"
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	gtk "github.com/pekim/gobbi/lib/gtk"
	soup "github.com/pekim/gobbi/lib/soup"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <webkit2/webkit2.h>
// #include <stdlib.h>
/*

	void cookiemanager_changedHandler(GObject *, gpointer);

	static gulong CookieManager_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(cookiemanager_changedHandler), data);
	}

*/
/*

	void download_createdDestinationHandler(GObject *, gchar*, gpointer);

	static gulong Download_signal_connect_created_destination(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "created-destination", G_CALLBACK(download_createdDestinationHandler), data);
	}

*/
/*

	gboolean download_decideDestinationHandler(GObject *, gchar*, gpointer);

	static gulong Download_signal_connect_decide_destination(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "decide-destination", G_CALLBACK(download_decideDestinationHandler), data);
	}

*/
/*

	void download_failedHandler(GObject *, GError *, gpointer);

	static gulong Download_signal_connect_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "failed", G_CALLBACK(download_failedHandler), data);
	}

*/
/*

	void download_finishedHandler(GObject *, gpointer);

	static gulong Download_signal_connect_finished(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "finished", G_CALLBACK(download_finishedHandler), data);
	}

*/
/*

	void download_receivedDataHandler(GObject *, guint64, gpointer);

	static gulong Download_signal_connect_received_data(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "received-data", G_CALLBACK(download_receivedDataHandler), data);
	}

*/
/*

	void favicondatabase_faviconChangedHandler(GObject *, gchar*, gchar*, gpointer);

	static gulong FaviconDatabase_signal_connect_favicon_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "favicon-changed", G_CALLBACK(favicondatabase_faviconChangedHandler), data);
	}

*/
/*

	void findcontroller_countedMatchesHandler(GObject *, guint, gpointer);

	static gulong FindController_signal_connect_counted_matches(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "counted-matches", G_CALLBACK(findcontroller_countedMatchesHandler), data);
	}

*/
/*

	void findcontroller_failedToFindTextHandler(GObject *, gpointer);

	static gulong FindController_signal_connect_failed_to_find_text(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "failed-to-find-text", G_CALLBACK(findcontroller_failedToFindTextHandler), data);
	}

*/
/*

	void findcontroller_foundTextHandler(GObject *, guint, gpointer);

	static gulong FindController_signal_connect_found_text(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "found-text", G_CALLBACK(findcontroller_foundTextHandler), data);
	}

*/
/*

	void printoperation_failedHandler(GObject *, GError *, gpointer);

	static gulong PrintOperation_signal_connect_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "failed", G_CALLBACK(printoperation_failedHandler), data);
	}

*/
/*

	void printoperation_finishedHandler(GObject *, gpointer);

	static gulong PrintOperation_signal_connect_finished(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "finished", G_CALLBACK(printoperation_finishedHandler), data);
	}

*/
/*

	void webcontext_downloadStartedHandler(GObject *, WebKitDownload *, gpointer);

	static gulong WebContext_signal_connect_download_started(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "download-started", G_CALLBACK(webcontext_downloadStartedHandler), data);
	}

*/
/*

	gboolean webinspector_attachHandler(GObject *, gpointer);

	static gulong WebInspector_signal_connect_attach(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "attach", G_CALLBACK(webinspector_attachHandler), data);
	}

*/
/*

	gboolean webinspector_bringToFrontHandler(GObject *, gpointer);

	static gulong WebInspector_signal_connect_bring_to_front(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "bring-to-front", G_CALLBACK(webinspector_bringToFrontHandler), data);
	}

*/
/*

	void webinspector_closedHandler(GObject *, gpointer);

	static gulong WebInspector_signal_connect_closed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closed", G_CALLBACK(webinspector_closedHandler), data);
	}

*/
/*

	gboolean webinspector_detachHandler(GObject *, gpointer);

	static gulong WebInspector_signal_connect_detach(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "detach", G_CALLBACK(webinspector_detachHandler), data);
	}

*/
/*

	gboolean webinspector_openWindowHandler(GObject *, gpointer);

	static gulong WebInspector_signal_connect_open_window(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "open-window", G_CALLBACK(webinspector_openWindowHandler), data);
	}

*/
/*

	void webresource_failedHandler(GObject *, GError *, gpointer);

	static gulong WebResource_signal_connect_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "failed", G_CALLBACK(webresource_failedHandler), data);
	}

*/
/*

	void webresource_finishedHandler(GObject *, gpointer);

	static gulong WebResource_signal_connect_finished(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "finished", G_CALLBACK(webresource_finishedHandler), data);
	}

*/
/*

	void webresource_receivedDataHandler(GObject *, guint64, gpointer);

	static gulong WebResource_signal_connect_received_data(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "received-data", G_CALLBACK(webresource_receivedDataHandler), data);
	}

*/
/*

	void webresource_sentRequestHandler(GObject *, WebKitURIRequest *, WebKitURIResponse *, gpointer);

	static gulong WebResource_signal_connect_sent_request(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "sent-request", G_CALLBACK(webresource_sentRequestHandler), data);
	}

*/
/*

	void webview_closeHandler(GObject *, gpointer);

	static gulong WebView_signal_connect_close(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "close", G_CALLBACK(webview_closeHandler), data);
	}

*/
/*

	void webview_contextMenuDismissedHandler(GObject *, gpointer);

	static gulong WebView_signal_connect_context_menu_dismissed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "context-menu-dismissed", G_CALLBACK(webview_contextMenuDismissedHandler), data);
	}

*/
/*

	GtkWidget * webview_createHandler(GObject *, WebKitNavigationAction *, gpointer);

	static gulong WebView_signal_connect_create(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "create", G_CALLBACK(webview_createHandler), data);
	}

*/
/*

	gboolean webview_decidePolicyHandler(GObject *, WebKitPolicyDecision *, WebKitPolicyDecisionType, gpointer);

	static gulong WebView_signal_connect_decide_policy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "decide-policy", G_CALLBACK(webview_decidePolicyHandler), data);
	}

*/
/*

	gboolean webview_enterFullscreenHandler(GObject *, gpointer);

	static gulong WebView_signal_connect_enter_fullscreen(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "enter-fullscreen", G_CALLBACK(webview_enterFullscreenHandler), data);
	}

*/
/*

	void webview_insecureContentDetectedHandler(GObject *, WebKitInsecureContentEvent, gpointer);

	static gulong WebView_signal_connect_insecure_content_detected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insecure-content-detected", G_CALLBACK(webview_insecureContentDetectedHandler), data);
	}

*/
/*

	gboolean webview_leaveFullscreenHandler(GObject *, gpointer);

	static gulong WebView_signal_connect_leave_fullscreen(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "leave-fullscreen", G_CALLBACK(webview_leaveFullscreenHandler), data);
	}

*/
/*

	void webview_loadChangedHandler(GObject *, WebKitLoadEvent, gpointer);

	static gulong WebView_signal_connect_load_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "load-changed", G_CALLBACK(webview_loadChangedHandler), data);
	}

*/
/*

	gboolean webview_loadFailedHandler(GObject *, WebKitLoadEvent, gchar*, GError *, gpointer);

	static gulong WebView_signal_connect_load_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "load-failed", G_CALLBACK(webview_loadFailedHandler), data);
	}

*/
/*

	void webview_mouseTargetChangedHandler(GObject *, WebKitHitTestResult *, guint, gpointer);

	static gulong WebView_signal_connect_mouse_target_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mouse-target-changed", G_CALLBACK(webview_mouseTargetChangedHandler), data);
	}

*/
/*

	gboolean webview_permissionRequestHandler(GObject *, WebKitPermissionRequest *, gpointer);

	static gulong WebView_signal_connect_permission_request(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "permission-request", G_CALLBACK(webview_permissionRequestHandler), data);
	}

*/
/*

	gboolean webview_printHandler(GObject *, WebKitPrintOperation *, gpointer);

	static gulong WebView_signal_connect_print(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "print", G_CALLBACK(webview_printHandler), data);
	}

*/
/*

	void webview_readyToShowHandler(GObject *, gpointer);

	static gulong WebView_signal_connect_ready_to_show(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "ready-to-show", G_CALLBACK(webview_readyToShowHandler), data);
	}

*/
/*

	void webview_resourceLoadStartedHandler(GObject *, WebKitWebResource *, WebKitURIRequest *, gpointer);

	static gulong WebView_signal_connect_resource_load_started(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "resource-load-started", G_CALLBACK(webview_resourceLoadStartedHandler), data);
	}

*/
/*

	void webview_runAsModalHandler(GObject *, gpointer);

	static gulong WebView_signal_connect_run_as_modal(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "run-as-modal", G_CALLBACK(webview_runAsModalHandler), data);
	}

*/
/*

	gboolean webview_runFileChooserHandler(GObject *, WebKitFileChooserRequest *, gpointer);

	static gulong WebView_signal_connect_run_file_chooser(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "run-file-chooser", G_CALLBACK(webview_runFileChooserHandler), data);
	}

*/
/*

	gboolean webview_scriptDialogHandler(GObject *, WebKitScriptDialog *, gpointer);

	static gulong WebView_signal_connect_script_dialog(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "script-dialog", G_CALLBACK(webview_scriptDialogHandler), data);
	}

*/
/*

	void webview_submitFormHandler(GObject *, WebKitFormSubmissionRequest *, gpointer);

	static gulong WebView_signal_connect_submit_form(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "submit-form", G_CALLBACK(webview_submitFormHandler), data);
	}

*/
/*

	gboolean webview_webProcessCrashedHandler(GObject *, gpointer);

	static gulong WebView_signal_connect_web_process_crashed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "web-process-crashed", G_CALLBACK(webview_webProcessCrashedHandler), data);
	}

*/
import "C"

type FindOptions C.WebKitFindOptions

const (
	WEBKIT_FIND_OPTIONS_NONE                               FindOptions = 0
	WEBKIT_FIND_OPTIONS_CASE_INSENSITIVE                   FindOptions = 1
	WEBKIT_FIND_OPTIONS_AT_WORD_STARTS                     FindOptions = 2
	WEBKIT_FIND_OPTIONS_TREAT_MEDIAL_CAPITAL_AS_WORD_START FindOptions = 4
	WEBKIT_FIND_OPTIONS_BACKWARDS                          FindOptions = 8
	WEBKIT_FIND_OPTIONS_WRAP_AROUND                        FindOptions = 16
)

type HitTestResultContext C.WebKitHitTestResultContext

const (
	WEBKIT_HIT_TEST_RESULT_CONTEXT_DOCUMENT  HitTestResultContext = 2
	WEBKIT_HIT_TEST_RESULT_CONTEXT_LINK      HitTestResultContext = 4
	WEBKIT_HIT_TEST_RESULT_CONTEXT_IMAGE     HitTestResultContext = 8
	WEBKIT_HIT_TEST_RESULT_CONTEXT_MEDIA     HitTestResultContext = 16
	WEBKIT_HIT_TEST_RESULT_CONTEXT_EDITABLE  HitTestResultContext = 32
	WEBKIT_HIT_TEST_RESULT_CONTEXT_SCROLLBAR HitTestResultContext = 64
	WEBKIT_HIT_TEST_RESULT_CONTEXT_SELECTION HitTestResultContext = 128
)

type SnapshotOptions C.WebKitSnapshotOptions

const (
	WEBKIT_SNAPSHOT_OPTIONS_NONE                           SnapshotOptions = 0
	WEBKIT_SNAPSHOT_OPTIONS_INCLUDE_SELECTION_HIGHLIGHTING SnapshotOptions = 1
	WEBKIT_SNAPSHOT_OPTIONS_TRANSPARENT_BACKGROUND         SnapshotOptions = 2
)

// AuthenticationRequest is a wrapper around the C record WebKitAuthenticationRequest.
type AuthenticationRequest struct {
	native *C.WebKitAuthenticationRequest
	// parent : record
	// Private : priv
}

func AuthenticationRequestNewFromC(u unsafe.Pointer) *AuthenticationRequest {
	c := (*C.WebKitAuthenticationRequest)(u)
	if c == nil {
		return nil
	}

	g := &AuthenticationRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AuthenticationRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AuthenticationRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthenticationRequest with another AuthenticationRequest, and returns true if they represent the same GObject.
func (recv *AuthenticationRequest) Equals(other *AuthenticationRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *AuthenticationRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to AuthenticationRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a AuthenticationRequest.
func CastToAuthenticationRequest(object *gobject.Object) *AuthenticationRequest {
	return AuthenticationRequestNewFromC(object.ToC())
}

// AutomationSession is a wrapper around the C record WebKitAutomationSession.
type AutomationSession struct {
	native *C.WebKitAutomationSession
	// parent : record
	// priv : record
}

func AutomationSessionNewFromC(u unsafe.Pointer) *AutomationSession {
	c := (*C.WebKitAutomationSession)(u)
	if c == nil {
		return nil
	}

	g := &AutomationSession{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AutomationSession) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AutomationSession) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AutomationSession with another AutomationSession, and returns true if they represent the same GObject.
func (recv *AutomationSession) Equals(other *AutomationSession) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *AutomationSession) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to AutomationSession.
// Exercise care, as this is a potentially dangerous function if the Object is not a AutomationSession.
func CastToAutomationSession(object *gobject.Object) *AutomationSession {
	return AutomationSessionNewFromC(object.ToC())
}

// BackForwardList is a wrapper around the C record WebKitBackForwardList.
type BackForwardList struct {
	native *C.WebKitBackForwardList
	// parent : record
	// priv : record
}

func BackForwardListNewFromC(u unsafe.Pointer) *BackForwardList {
	c := (*C.WebKitBackForwardList)(u)
	if c == nil {
		return nil
	}

	g := &BackForwardList{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *BackForwardList) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *BackForwardList) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BackForwardList with another BackForwardList, and returns true if they represent the same GObject.
func (recv *BackForwardList) Equals(other *BackForwardList) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *BackForwardList) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to BackForwardList.
// Exercise care, as this is a potentially dangerous function if the Object is not a BackForwardList.
func CastToBackForwardList(object *gobject.Object) *BackForwardList {
	return BackForwardListNewFromC(object.ToC())
}

// Unsupported signal 'changed' for BackForwardList : param items_removed : gpointer

// GetBackItem is a wrapper around the C function webkit_back_forward_list_get_back_item.
func (recv *BackForwardList) GetBackItem() *BackForwardListItem {
	retC := C.webkit_back_forward_list_get_back_item((*C.WebKitBackForwardList)(recv.native))
	var retGo (*BackForwardListItem)
	if retC == nil {
		retGo = nil
	} else {
		retGo = BackForwardListItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetBackList is a wrapper around the C function webkit_back_forward_list_get_back_list.
func (recv *BackForwardList) GetBackList() *glib.List {
	retC := C.webkit_back_forward_list_get_back_list((*C.WebKitBackForwardList)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBackListWithLimit is a wrapper around the C function webkit_back_forward_list_get_back_list_with_limit.
func (recv *BackForwardList) GetBackListWithLimit(limit uint32) *glib.List {
	c_limit := (C.guint)(limit)

	retC := C.webkit_back_forward_list_get_back_list_with_limit((*C.WebKitBackForwardList)(recv.native), c_limit)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentItem is a wrapper around the C function webkit_back_forward_list_get_current_item.
func (recv *BackForwardList) GetCurrentItem() *BackForwardListItem {
	retC := C.webkit_back_forward_list_get_current_item((*C.WebKitBackForwardList)(recv.native))
	var retGo (*BackForwardListItem)
	if retC == nil {
		retGo = nil
	} else {
		retGo = BackForwardListItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetForwardItem is a wrapper around the C function webkit_back_forward_list_get_forward_item.
func (recv *BackForwardList) GetForwardItem() *BackForwardListItem {
	retC := C.webkit_back_forward_list_get_forward_item((*C.WebKitBackForwardList)(recv.native))
	var retGo (*BackForwardListItem)
	if retC == nil {
		retGo = nil
	} else {
		retGo = BackForwardListItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetForwardList is a wrapper around the C function webkit_back_forward_list_get_forward_list.
func (recv *BackForwardList) GetForwardList() *glib.List {
	retC := C.webkit_back_forward_list_get_forward_list((*C.WebKitBackForwardList)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetForwardListWithLimit is a wrapper around the C function webkit_back_forward_list_get_forward_list_with_limit.
func (recv *BackForwardList) GetForwardListWithLimit(limit uint32) *glib.List {
	c_limit := (C.guint)(limit)

	retC := C.webkit_back_forward_list_get_forward_list_with_limit((*C.WebKitBackForwardList)(recv.native), c_limit)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLength is a wrapper around the C function webkit_back_forward_list_get_length.
func (recv *BackForwardList) GetLength() uint32 {
	retC := C.webkit_back_forward_list_get_length((*C.WebKitBackForwardList)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetNthItem is a wrapper around the C function webkit_back_forward_list_get_nth_item.
func (recv *BackForwardList) GetNthItem(index int32) *BackForwardListItem {
	c_index := (C.gint)(index)

	retC := C.webkit_back_forward_list_get_nth_item((*C.WebKitBackForwardList)(recv.native), c_index)
	var retGo (*BackForwardListItem)
	if retC == nil {
		retGo = nil
	} else {
		retGo = BackForwardListItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// BackForwardListItem is a wrapper around the C record WebKitBackForwardListItem.
type BackForwardListItem struct {
	native *C.WebKitBackForwardListItem
	// parent : record
	// priv : record
}

func BackForwardListItemNewFromC(u unsafe.Pointer) *BackForwardListItem {
	c := (*C.WebKitBackForwardListItem)(u)
	if c == nil {
		return nil
	}

	g := &BackForwardListItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *BackForwardListItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *BackForwardListItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BackForwardListItem with another BackForwardListItem, and returns true if they represent the same GObject.
func (recv *BackForwardListItem) Equals(other *BackForwardListItem) bool {
	return other.ToC() == recv.ToC()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *BackForwardListItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *BackForwardListItem) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitrary Object to BackForwardListItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a BackForwardListItem.
func CastToBackForwardListItem(object *gobject.Object) *BackForwardListItem {
	return BackForwardListItemNewFromC(object.ToC())
}

// GetOriginalUri is a wrapper around the C function webkit_back_forward_list_item_get_original_uri.
func (recv *BackForwardListItem) GetOriginalUri() string {
	retC := C.webkit_back_forward_list_item_get_original_uri((*C.WebKitBackForwardListItem)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTitle is a wrapper around the C function webkit_back_forward_list_item_get_title.
func (recv *BackForwardListItem) GetTitle() string {
	retC := C.webkit_back_forward_list_item_get_title((*C.WebKitBackForwardListItem)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUri is a wrapper around the C function webkit_back_forward_list_item_get_uri.
func (recv *BackForwardListItem) GetUri() string {
	retC := C.webkit_back_forward_list_item_get_uri((*C.WebKitBackForwardListItem)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// ColorChooserRequest is a wrapper around the C record WebKitColorChooserRequest.
type ColorChooserRequest struct {
	native *C.WebKitColorChooserRequest
	// parent : record
	// Private : priv
}

func ColorChooserRequestNewFromC(u unsafe.Pointer) *ColorChooserRequest {
	c := (*C.WebKitColorChooserRequest)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ColorChooserRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ColorChooserRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ColorChooserRequest with another ColorChooserRequest, and returns true if they represent the same GObject.
func (recv *ColorChooserRequest) Equals(other *ColorChooserRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ColorChooserRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ColorChooserRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a ColorChooserRequest.
func CastToColorChooserRequest(object *gobject.Object) *ColorChooserRequest {
	return ColorChooserRequestNewFromC(object.ToC())
}

// ContextMenu is a wrapper around the C record WebKitContextMenu.
type ContextMenu struct {
	native *C.WebKitContextMenu
	// parent : record
	// priv : record
}

func ContextMenuNewFromC(u unsafe.Pointer) *ContextMenu {
	c := (*C.WebKitContextMenu)(u)
	if c == nil {
		return nil
	}

	g := &ContextMenu{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ContextMenu) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ContextMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContextMenu with another ContextMenu, and returns true if they represent the same GObject.
func (recv *ContextMenu) Equals(other *ContextMenu) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *ContextMenu) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to ContextMenu.
// Exercise care, as this is a potentially dangerous function if the Object is not a ContextMenu.
func CastToContextMenu(object *gobject.Object) *ContextMenu {
	return ContextMenuNewFromC(object.ToC())
}

// ContextMenuNew is a wrapper around the C function webkit_context_menu_new.
func ContextMenuNew() *ContextMenu {
	retC := C.webkit_context_menu_new()
	retGo := ContextMenuNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ContextMenuNewWithItems is a wrapper around the C function webkit_context_menu_new_with_items.
func ContextMenuNewWithItems(items *glib.List) *ContextMenu {
	c_items := (*C.GList)(C.NULL)
	if items != nil {
		c_items = (*C.GList)(items.ToC())
	}

	retC := C.webkit_context_menu_new_with_items(c_items)
	retGo := ContextMenuNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Append is a wrapper around the C function webkit_context_menu_append.
func (recv *ContextMenu) Append(item *ContextMenuItem) {
	c_item := (*C.WebKitContextMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.WebKitContextMenuItem)(item.ToC())
	}

	C.webkit_context_menu_append((*C.WebKitContextMenu)(recv.native), c_item)

	return
}

// First is a wrapper around the C function webkit_context_menu_first.
func (recv *ContextMenu) First() *ContextMenuItem {
	retC := C.webkit_context_menu_first((*C.WebKitContextMenu)(recv.native))
	retGo := ContextMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetItemAtPosition is a wrapper around the C function webkit_context_menu_get_item_at_position.
func (recv *ContextMenu) GetItemAtPosition(position uint32) *ContextMenuItem {
	c_position := (C.guint)(position)

	retC := C.webkit_context_menu_get_item_at_position((*C.WebKitContextMenu)(recv.native), c_position)
	retGo := ContextMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetItems is a wrapper around the C function webkit_context_menu_get_items.
func (recv *ContextMenu) GetItems() *glib.List {
	retC := C.webkit_context_menu_get_items((*C.WebKitContextMenu)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetNItems is a wrapper around the C function webkit_context_menu_get_n_items.
func (recv *ContextMenu) GetNItems() uint32 {
	retC := C.webkit_context_menu_get_n_items((*C.WebKitContextMenu)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Insert is a wrapper around the C function webkit_context_menu_insert.
func (recv *ContextMenu) Insert(item *ContextMenuItem, position int32) {
	c_item := (*C.WebKitContextMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.WebKitContextMenuItem)(item.ToC())
	}

	c_position := (C.gint)(position)

	C.webkit_context_menu_insert((*C.WebKitContextMenu)(recv.native), c_item, c_position)

	return
}

// Last is a wrapper around the C function webkit_context_menu_last.
func (recv *ContextMenu) Last() *ContextMenuItem {
	retC := C.webkit_context_menu_last((*C.WebKitContextMenu)(recv.native))
	retGo := ContextMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MoveItem is a wrapper around the C function webkit_context_menu_move_item.
func (recv *ContextMenu) MoveItem(item *ContextMenuItem, position int32) {
	c_item := (*C.WebKitContextMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.WebKitContextMenuItem)(item.ToC())
	}

	c_position := (C.gint)(position)

	C.webkit_context_menu_move_item((*C.WebKitContextMenu)(recv.native), c_item, c_position)

	return
}

// Prepend is a wrapper around the C function webkit_context_menu_prepend.
func (recv *ContextMenu) Prepend(item *ContextMenuItem) {
	c_item := (*C.WebKitContextMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.WebKitContextMenuItem)(item.ToC())
	}

	C.webkit_context_menu_prepend((*C.WebKitContextMenu)(recv.native), c_item)

	return
}

// Remove is a wrapper around the C function webkit_context_menu_remove.
func (recv *ContextMenu) Remove(item *ContextMenuItem) {
	c_item := (*C.WebKitContextMenuItem)(C.NULL)
	if item != nil {
		c_item = (*C.WebKitContextMenuItem)(item.ToC())
	}

	C.webkit_context_menu_remove((*C.WebKitContextMenu)(recv.native), c_item)

	return
}

// RemoveAll is a wrapper around the C function webkit_context_menu_remove_all.
func (recv *ContextMenu) RemoveAll() {
	C.webkit_context_menu_remove_all((*C.WebKitContextMenu)(recv.native))

	return
}

// ContextMenuItem is a wrapper around the C record WebKitContextMenuItem.
type ContextMenuItem struct {
	native *C.WebKitContextMenuItem
	// parent : record
	// priv : record
}

func ContextMenuItemNewFromC(u unsafe.Pointer) *ContextMenuItem {
	c := (*C.WebKitContextMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &ContextMenuItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ContextMenuItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ContextMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContextMenuItem with another ContextMenuItem, and returns true if they represent the same GObject.
func (recv *ContextMenuItem) Equals(other *ContextMenuItem) bool {
	return other.ToC() == recv.ToC()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ContextMenuItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *ContextMenuItem) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitrary Object to ContextMenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a ContextMenuItem.
func CastToContextMenuItem(object *gobject.Object) *ContextMenuItem {
	return ContextMenuItemNewFromC(object.ToC())
}

// ContextMenuItemNew is a wrapper around the C function webkit_context_menu_item_new.
func ContextMenuItemNew(action *gtk.Action) *ContextMenuItem {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	retC := C.webkit_context_menu_item_new(c_action)
	retGo := ContextMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ContextMenuItemNewFromStockAction is a wrapper around the C function webkit_context_menu_item_new_from_stock_action.
func ContextMenuItemNewFromStockAction(action ContextMenuAction) *ContextMenuItem {
	c_action := (C.WebKitContextMenuAction)(action)

	retC := C.webkit_context_menu_item_new_from_stock_action(c_action)
	retGo := ContextMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ContextMenuItemNewFromStockActionWithLabel is a wrapper around the C function webkit_context_menu_item_new_from_stock_action_with_label.
func ContextMenuItemNewFromStockActionWithLabel(action ContextMenuAction, label string) *ContextMenuItem {
	c_action := (C.WebKitContextMenuAction)(action)

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.webkit_context_menu_item_new_from_stock_action_with_label(c_action, c_label)
	retGo := ContextMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ContextMenuItemNewSeparator is a wrapper around the C function webkit_context_menu_item_new_separator.
func ContextMenuItemNewSeparator() *ContextMenuItem {
	retC := C.webkit_context_menu_item_new_separator()
	retGo := ContextMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ContextMenuItemNewWithSubmenu is a wrapper around the C function webkit_context_menu_item_new_with_submenu.
func ContextMenuItemNewWithSubmenu(label string, submenu *ContextMenu) *ContextMenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_submenu := (*C.WebKitContextMenu)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.WebKitContextMenu)(submenu.ToC())
	}

	retC := C.webkit_context_menu_item_new_with_submenu(c_label, c_submenu)
	retGo := ContextMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAction is a wrapper around the C function webkit_context_menu_item_get_action.
func (recv *ContextMenuItem) GetAction() *gtk.Action {
	retC := C.webkit_context_menu_item_get_action((*C.WebKitContextMenuItem)(recv.native))
	retGo := gtk.ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStockAction is a wrapper around the C function webkit_context_menu_item_get_stock_action.
func (recv *ContextMenuItem) GetStockAction() ContextMenuAction {
	retC := C.webkit_context_menu_item_get_stock_action((*C.WebKitContextMenuItem)(recv.native))
	retGo := (ContextMenuAction)(retC)

	return retGo
}

// GetSubmenu is a wrapper around the C function webkit_context_menu_item_get_submenu.
func (recv *ContextMenuItem) GetSubmenu() *ContextMenu {
	retC := C.webkit_context_menu_item_get_submenu((*C.WebKitContextMenuItem)(recv.native))
	retGo := ContextMenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsSeparator is a wrapper around the C function webkit_context_menu_item_is_separator.
func (recv *ContextMenuItem) IsSeparator() bool {
	retC := C.webkit_context_menu_item_is_separator((*C.WebKitContextMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetSubmenu is a wrapper around the C function webkit_context_menu_item_set_submenu.
func (recv *ContextMenuItem) SetSubmenu(submenu *ContextMenu) {
	c_submenu := (*C.WebKitContextMenu)(C.NULL)
	if submenu != nil {
		c_submenu = (*C.WebKitContextMenu)(submenu.ToC())
	}

	C.webkit_context_menu_item_set_submenu((*C.WebKitContextMenuItem)(recv.native), c_submenu)

	return
}

// CookieManager is a wrapper around the C record WebKitCookieManager.
type CookieManager struct {
	native *C.WebKitCookieManager
	// parent : record
	// priv : record
}

func CookieManagerNewFromC(u unsafe.Pointer) *CookieManager {
	c := (*C.WebKitCookieManager)(u)
	if c == nil {
		return nil
	}

	g := &CookieManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *CookieManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *CookieManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CookieManager with another CookieManager, and returns true if they represent the same GObject.
func (recv *CookieManager) Equals(other *CookieManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *CookieManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to CookieManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a CookieManager.
func CastToCookieManager(object *gobject.Object) *CookieManager {
	return CookieManagerNewFromC(object.ToC())
}

type signalCookieManagerChangedDetail struct {
	callback  CookieManagerSignalChangedCallback
	handlerID C.gulong
}

var signalCookieManagerChangedId int
var signalCookieManagerChangedMap = make(map[int]signalCookieManagerChangedDetail)
var signalCookieManagerChangedLock sync.RWMutex

// CookieManagerSignalChangedCallback is a callback function for a 'changed' signal emitted from a CookieManager.
type CookieManagerSignalChangedCallback func(targetObject *CookieManager)

/*
ConnectChanged connects the callback to the 'changed' signal for the CookieManager.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *CookieManager) ConnectChanged(callback CookieManagerSignalChangedCallback) int {
	signalCookieManagerChangedLock.Lock()
	defer signalCookieManagerChangedLock.Unlock()

	signalCookieManagerChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CookieManager_signal_connect_changed(instance, C.gpointer(uintptr(signalCookieManagerChangedId)))

	detail := signalCookieManagerChangedDetail{callback, handlerID}
	signalCookieManagerChangedMap[signalCookieManagerChangedId] = detail

	return signalCookieManagerChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the CookieManager.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *CookieManager) DisconnectChanged(connectionID int) {
	signalCookieManagerChangedLock.Lock()
	defer signalCookieManagerChangedLock.Unlock()

	detail, exists := signalCookieManagerChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCookieManagerChangedMap, connectionID)
}

//export cookiemanager_changedHandler
func cookiemanager_changedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalCookieManagerChangedLock.RLock()
	defer signalCookieManagerChangedLock.RUnlock()

	targetObject := CookieManagerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalCookieManagerChangedMap[index].callback
	callback(targetObject)
}

// DeleteAllCookies is a wrapper around the C function webkit_cookie_manager_delete_all_cookies.
func (recv *CookieManager) DeleteAllCookies() {
	C.webkit_cookie_manager_delete_all_cookies((*C.WebKitCookieManager)(recv.native))

	return
}

// DeleteCookiesForDomain is a wrapper around the C function webkit_cookie_manager_delete_cookies_for_domain.
func (recv *CookieManager) DeleteCookiesForDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.webkit_cookie_manager_delete_cookies_for_domain((*C.WebKitCookieManager)(recv.native), c_domain)

	return
}

// Unsupported : webkit_cookie_manager_get_accept_policy : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// GetAcceptPolicyFinish is a wrapper around the C function webkit_cookie_manager_get_accept_policy_finish.
func (recv *CookieManager) GetAcceptPolicyFinish(result *gio.AsyncResult) (CookieAcceptPolicy, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_cookie_manager_get_accept_policy_finish((*C.WebKitCookieManager)(recv.native), c_result, &cThrowableError)
	retGo := (CookieAcceptPolicy)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : webkit_cookie_manager_get_domains_with_cookies : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// GetDomainsWithCookiesFinish is a wrapper around the C function webkit_cookie_manager_get_domains_with_cookies_finish.
func (recv *CookieManager) GetDomainsWithCookiesFinish(result *gio.AsyncResult) ([]string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_cookie_manager_get_domains_with_cookies_finish((*C.WebKitCookieManager)(recv.native), c_result, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetAcceptPolicy is a wrapper around the C function webkit_cookie_manager_set_accept_policy.
func (recv *CookieManager) SetAcceptPolicy(policy CookieAcceptPolicy) {
	c_policy := (C.WebKitCookieAcceptPolicy)(policy)

	C.webkit_cookie_manager_set_accept_policy((*C.WebKitCookieManager)(recv.native), c_policy)

	return
}

// SetPersistentStorage is a wrapper around the C function webkit_cookie_manager_set_persistent_storage.
func (recv *CookieManager) SetPersistentStorage(filename string, storage CookiePersistentStorage) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_storage := (C.WebKitCookiePersistentStorage)(storage)

	C.webkit_cookie_manager_set_persistent_storage((*C.WebKitCookieManager)(recv.native), c_filename, c_storage)

	return
}

// DeviceInfoPermissionRequest is a wrapper around the C record WebKitDeviceInfoPermissionRequest.
type DeviceInfoPermissionRequest struct {
	native *C.WebKitDeviceInfoPermissionRequest
	// parent : record
	// Private : priv
}

func DeviceInfoPermissionRequestNewFromC(u unsafe.Pointer) *DeviceInfoPermissionRequest {
	c := (*C.WebKitDeviceInfoPermissionRequest)(u)
	if c == nil {
		return nil
	}

	g := &DeviceInfoPermissionRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *DeviceInfoPermissionRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *DeviceInfoPermissionRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DeviceInfoPermissionRequest with another DeviceInfoPermissionRequest, and returns true if they represent the same GObject.
func (recv *DeviceInfoPermissionRequest) Equals(other *DeviceInfoPermissionRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *DeviceInfoPermissionRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to DeviceInfoPermissionRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a DeviceInfoPermissionRequest.
func CastToDeviceInfoPermissionRequest(object *gobject.Object) *DeviceInfoPermissionRequest {
	return DeviceInfoPermissionRequestNewFromC(object.ToC())
}

// PermissionRequest returns the PermissionRequest interface implemented by DeviceInfoPermissionRequest
func (recv *DeviceInfoPermissionRequest) PermissionRequest() *PermissionRequest {
	return PermissionRequestNewFromC(recv.ToC())
}

// Download is a wrapper around the C record WebKitDownload.
type Download struct {
	native *C.WebKitDownload
	// parent : record
	// priv : record
}

func DownloadNewFromC(u unsafe.Pointer) *Download {
	c := (*C.WebKitDownload)(u)
	if c == nil {
		return nil
	}

	g := &Download{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Download) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Download) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Download with another Download, and returns true if they represent the same GObject.
func (recv *Download) Equals(other *Download) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Download) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Download.
// Exercise care, as this is a potentially dangerous function if the Object is not a Download.
func CastToDownload(object *gobject.Object) *Download {
	return DownloadNewFromC(object.ToC())
}

type signalDownloadCreatedDestinationDetail struct {
	callback  DownloadSignalCreatedDestinationCallback
	handlerID C.gulong
}

var signalDownloadCreatedDestinationId int
var signalDownloadCreatedDestinationMap = make(map[int]signalDownloadCreatedDestinationDetail)
var signalDownloadCreatedDestinationLock sync.RWMutex

// DownloadSignalCreatedDestinationCallback is a callback function for a 'created-destination' signal emitted from a Download.
type DownloadSignalCreatedDestinationCallback func(targetObject *Download, destination string)

/*
ConnectCreatedDestination connects the callback to the 'created-destination' signal for the Download.

The returned value represents the connection, and may be passed to DisconnectCreatedDestination to remove it.
*/
func (recv *Download) ConnectCreatedDestination(callback DownloadSignalCreatedDestinationCallback) int {
	signalDownloadCreatedDestinationLock.Lock()
	defer signalDownloadCreatedDestinationLock.Unlock()

	signalDownloadCreatedDestinationId++
	instance := C.gpointer(recv.native)
	handlerID := C.Download_signal_connect_created_destination(instance, C.gpointer(uintptr(signalDownloadCreatedDestinationId)))

	detail := signalDownloadCreatedDestinationDetail{callback, handlerID}
	signalDownloadCreatedDestinationMap[signalDownloadCreatedDestinationId] = detail

	return signalDownloadCreatedDestinationId
}

/*
DisconnectCreatedDestination disconnects a callback from the 'created-destination' signal for the Download.

The connectionID should be a value returned from a call to ConnectCreatedDestination.
*/
func (recv *Download) DisconnectCreatedDestination(connectionID int) {
	signalDownloadCreatedDestinationLock.Lock()
	defer signalDownloadCreatedDestinationLock.Unlock()

	detail, exists := signalDownloadCreatedDestinationMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDownloadCreatedDestinationMap, connectionID)
}

//export download_createdDestinationHandler
func download_createdDestinationHandler(c_targetObject *C.GObject, c_destination *C.gchar, data C.gpointer) {
	signalDownloadCreatedDestinationLock.RLock()
	defer signalDownloadCreatedDestinationLock.RUnlock()

	destination := C.GoString(c_destination)

	targetObject := DownloadNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalDownloadCreatedDestinationMap[index].callback
	callback(targetObject, destination)
}

type signalDownloadDecideDestinationDetail struct {
	callback  DownloadSignalDecideDestinationCallback
	handlerID C.gulong
}

var signalDownloadDecideDestinationId int
var signalDownloadDecideDestinationMap = make(map[int]signalDownloadDecideDestinationDetail)
var signalDownloadDecideDestinationLock sync.RWMutex

// DownloadSignalDecideDestinationCallback is a callback function for a 'decide-destination' signal emitted from a Download.
type DownloadSignalDecideDestinationCallback func(targetObject *Download, suggestedFilename string) bool

/*
ConnectDecideDestination connects the callback to the 'decide-destination' signal for the Download.

The returned value represents the connection, and may be passed to DisconnectDecideDestination to remove it.
*/
func (recv *Download) ConnectDecideDestination(callback DownloadSignalDecideDestinationCallback) int {
	signalDownloadDecideDestinationLock.Lock()
	defer signalDownloadDecideDestinationLock.Unlock()

	signalDownloadDecideDestinationId++
	instance := C.gpointer(recv.native)
	handlerID := C.Download_signal_connect_decide_destination(instance, C.gpointer(uintptr(signalDownloadDecideDestinationId)))

	detail := signalDownloadDecideDestinationDetail{callback, handlerID}
	signalDownloadDecideDestinationMap[signalDownloadDecideDestinationId] = detail

	return signalDownloadDecideDestinationId
}

/*
DisconnectDecideDestination disconnects a callback from the 'decide-destination' signal for the Download.

The connectionID should be a value returned from a call to ConnectDecideDestination.
*/
func (recv *Download) DisconnectDecideDestination(connectionID int) {
	signalDownloadDecideDestinationLock.Lock()
	defer signalDownloadDecideDestinationLock.Unlock()

	detail, exists := signalDownloadDecideDestinationMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDownloadDecideDestinationMap, connectionID)
}

//export download_decideDestinationHandler
func download_decideDestinationHandler(c_targetObject *C.GObject, c_suggested_filename *C.gchar, data C.gpointer) C.gboolean {
	signalDownloadDecideDestinationLock.RLock()
	defer signalDownloadDecideDestinationLock.RUnlock()

	suggestedFilename := C.GoString(c_suggested_filename)

	targetObject := DownloadNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalDownloadDecideDestinationMap[index].callback
	retGo := callback(targetObject, suggestedFilename)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalDownloadFailedDetail struct {
	callback  DownloadSignalFailedCallback
	handlerID C.gulong
}

var signalDownloadFailedId int
var signalDownloadFailedMap = make(map[int]signalDownloadFailedDetail)
var signalDownloadFailedLock sync.RWMutex

// DownloadSignalFailedCallback is a callback function for a 'failed' signal emitted from a Download.
type DownloadSignalFailedCallback func(targetObject *Download, error *glib.Error)

/*
ConnectFailed connects the callback to the 'failed' signal for the Download.

The returned value represents the connection, and may be passed to DisconnectFailed to remove it.
*/
func (recv *Download) ConnectFailed(callback DownloadSignalFailedCallback) int {
	signalDownloadFailedLock.Lock()
	defer signalDownloadFailedLock.Unlock()

	signalDownloadFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Download_signal_connect_failed(instance, C.gpointer(uintptr(signalDownloadFailedId)))

	detail := signalDownloadFailedDetail{callback, handlerID}
	signalDownloadFailedMap[signalDownloadFailedId] = detail

	return signalDownloadFailedId
}

/*
DisconnectFailed disconnects a callback from the 'failed' signal for the Download.

The connectionID should be a value returned from a call to ConnectFailed.
*/
func (recv *Download) DisconnectFailed(connectionID int) {
	signalDownloadFailedLock.Lock()
	defer signalDownloadFailedLock.Unlock()

	detail, exists := signalDownloadFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDownloadFailedMap, connectionID)
}

//export download_failedHandler
func download_failedHandler(c_targetObject *C.GObject, c_error *C.GError, data C.gpointer) {
	signalDownloadFailedLock.RLock()
	defer signalDownloadFailedLock.RUnlock()

	error := glib.ErrorNewFromC(unsafe.Pointer(c_error))

	targetObject := DownloadNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalDownloadFailedMap[index].callback
	callback(targetObject, error)
}

type signalDownloadFinishedDetail struct {
	callback  DownloadSignalFinishedCallback
	handlerID C.gulong
}

var signalDownloadFinishedId int
var signalDownloadFinishedMap = make(map[int]signalDownloadFinishedDetail)
var signalDownloadFinishedLock sync.RWMutex

// DownloadSignalFinishedCallback is a callback function for a 'finished' signal emitted from a Download.
type DownloadSignalFinishedCallback func(targetObject *Download)

/*
ConnectFinished connects the callback to the 'finished' signal for the Download.

The returned value represents the connection, and may be passed to DisconnectFinished to remove it.
*/
func (recv *Download) ConnectFinished(callback DownloadSignalFinishedCallback) int {
	signalDownloadFinishedLock.Lock()
	defer signalDownloadFinishedLock.Unlock()

	signalDownloadFinishedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Download_signal_connect_finished(instance, C.gpointer(uintptr(signalDownloadFinishedId)))

	detail := signalDownloadFinishedDetail{callback, handlerID}
	signalDownloadFinishedMap[signalDownloadFinishedId] = detail

	return signalDownloadFinishedId
}

/*
DisconnectFinished disconnects a callback from the 'finished' signal for the Download.

The connectionID should be a value returned from a call to ConnectFinished.
*/
func (recv *Download) DisconnectFinished(connectionID int) {
	signalDownloadFinishedLock.Lock()
	defer signalDownloadFinishedLock.Unlock()

	detail, exists := signalDownloadFinishedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDownloadFinishedMap, connectionID)
}

//export download_finishedHandler
func download_finishedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalDownloadFinishedLock.RLock()
	defer signalDownloadFinishedLock.RUnlock()

	targetObject := DownloadNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalDownloadFinishedMap[index].callback
	callback(targetObject)
}

type signalDownloadReceivedDataDetail struct {
	callback  DownloadSignalReceivedDataCallback
	handlerID C.gulong
}

var signalDownloadReceivedDataId int
var signalDownloadReceivedDataMap = make(map[int]signalDownloadReceivedDataDetail)
var signalDownloadReceivedDataLock sync.RWMutex

// DownloadSignalReceivedDataCallback is a callback function for a 'received-data' signal emitted from a Download.
type DownloadSignalReceivedDataCallback func(targetObject *Download, dataLength uint64)

/*
ConnectReceivedData connects the callback to the 'received-data' signal for the Download.

The returned value represents the connection, and may be passed to DisconnectReceivedData to remove it.
*/
func (recv *Download) ConnectReceivedData(callback DownloadSignalReceivedDataCallback) int {
	signalDownloadReceivedDataLock.Lock()
	defer signalDownloadReceivedDataLock.Unlock()

	signalDownloadReceivedDataId++
	instance := C.gpointer(recv.native)
	handlerID := C.Download_signal_connect_received_data(instance, C.gpointer(uintptr(signalDownloadReceivedDataId)))

	detail := signalDownloadReceivedDataDetail{callback, handlerID}
	signalDownloadReceivedDataMap[signalDownloadReceivedDataId] = detail

	return signalDownloadReceivedDataId
}

/*
DisconnectReceivedData disconnects a callback from the 'received-data' signal for the Download.

The connectionID should be a value returned from a call to ConnectReceivedData.
*/
func (recv *Download) DisconnectReceivedData(connectionID int) {
	signalDownloadReceivedDataLock.Lock()
	defer signalDownloadReceivedDataLock.Unlock()

	detail, exists := signalDownloadReceivedDataMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDownloadReceivedDataMap, connectionID)
}

//export download_receivedDataHandler
func download_receivedDataHandler(c_targetObject *C.GObject, c_data_length C.guint64, data C.gpointer) {
	signalDownloadReceivedDataLock.RLock()
	defer signalDownloadReceivedDataLock.RUnlock()

	dataLength := uint64(c_data_length)

	targetObject := DownloadNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalDownloadReceivedDataMap[index].callback
	callback(targetObject, dataLength)
}

// Cancel is a wrapper around the C function webkit_download_cancel.
func (recv *Download) Cancel() {
	C.webkit_download_cancel((*C.WebKitDownload)(recv.native))

	return
}

// GetDestination is a wrapper around the C function webkit_download_get_destination.
func (recv *Download) GetDestination() string {
	retC := C.webkit_download_get_destination((*C.WebKitDownload)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetElapsedTime is a wrapper around the C function webkit_download_get_elapsed_time.
func (recv *Download) GetElapsedTime() float64 {
	retC := C.webkit_download_get_elapsed_time((*C.WebKitDownload)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetEstimatedProgress is a wrapper around the C function webkit_download_get_estimated_progress.
func (recv *Download) GetEstimatedProgress() float64 {
	retC := C.webkit_download_get_estimated_progress((*C.WebKitDownload)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetReceivedDataLength is a wrapper around the C function webkit_download_get_received_data_length.
func (recv *Download) GetReceivedDataLength() uint64 {
	retC := C.webkit_download_get_received_data_length((*C.WebKitDownload)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetRequest is a wrapper around the C function webkit_download_get_request.
func (recv *Download) GetRequest() *URIRequest {
	retC := C.webkit_download_get_request((*C.WebKitDownload)(recv.native))
	retGo := URIRequestNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetResponse is a wrapper around the C function webkit_download_get_response.
func (recv *Download) GetResponse() *URIResponse {
	retC := C.webkit_download_get_response((*C.WebKitDownload)(recv.native))
	retGo := URIResponseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWebView is a wrapper around the C function webkit_download_get_web_view.
func (recv *Download) GetWebView() *WebView {
	retC := C.webkit_download_get_web_view((*C.WebKitDownload)(recv.native))
	retGo := WebViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetDestination is a wrapper around the C function webkit_download_set_destination.
func (recv *Download) SetDestination(uri string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.webkit_download_set_destination((*C.WebKitDownload)(recv.native), c_uri)

	return
}

// EditorState is a wrapper around the C record WebKitEditorState.
type EditorState struct {
	native *C.WebKitEditorState
	// parent : record
	// priv : record
}

func EditorStateNewFromC(u unsafe.Pointer) *EditorState {
	c := (*C.WebKitEditorState)(u)
	if c == nil {
		return nil
	}

	g := &EditorState{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *EditorState) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *EditorState) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EditorState with another EditorState, and returns true if they represent the same GObject.
func (recv *EditorState) Equals(other *EditorState) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *EditorState) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to EditorState.
// Exercise care, as this is a potentially dangerous function if the Object is not a EditorState.
func CastToEditorState(object *gobject.Object) *EditorState {
	return EditorStateNewFromC(object.ToC())
}

// FaviconDatabase is a wrapper around the C record WebKitFaviconDatabase.
type FaviconDatabase struct {
	native *C.WebKitFaviconDatabase
	// parent : record
	// priv : record
}

func FaviconDatabaseNewFromC(u unsafe.Pointer) *FaviconDatabase {
	c := (*C.WebKitFaviconDatabase)(u)
	if c == nil {
		return nil
	}

	g := &FaviconDatabase{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FaviconDatabase) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FaviconDatabase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FaviconDatabase with another FaviconDatabase, and returns true if they represent the same GObject.
func (recv *FaviconDatabase) Equals(other *FaviconDatabase) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FaviconDatabase) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FaviconDatabase.
// Exercise care, as this is a potentially dangerous function if the Object is not a FaviconDatabase.
func CastToFaviconDatabase(object *gobject.Object) *FaviconDatabase {
	return FaviconDatabaseNewFromC(object.ToC())
}

type signalFaviconDatabaseFaviconChangedDetail struct {
	callback  FaviconDatabaseSignalFaviconChangedCallback
	handlerID C.gulong
}

var signalFaviconDatabaseFaviconChangedId int
var signalFaviconDatabaseFaviconChangedMap = make(map[int]signalFaviconDatabaseFaviconChangedDetail)
var signalFaviconDatabaseFaviconChangedLock sync.RWMutex

// FaviconDatabaseSignalFaviconChangedCallback is a callback function for a 'favicon-changed' signal emitted from a FaviconDatabase.
type FaviconDatabaseSignalFaviconChangedCallback func(targetObject *FaviconDatabase, pageUri string, faviconUri string)

/*
ConnectFaviconChanged connects the callback to the 'favicon-changed' signal for the FaviconDatabase.

The returned value represents the connection, and may be passed to DisconnectFaviconChanged to remove it.
*/
func (recv *FaviconDatabase) ConnectFaviconChanged(callback FaviconDatabaseSignalFaviconChangedCallback) int {
	signalFaviconDatabaseFaviconChangedLock.Lock()
	defer signalFaviconDatabaseFaviconChangedLock.Unlock()

	signalFaviconDatabaseFaviconChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.FaviconDatabase_signal_connect_favicon_changed(instance, C.gpointer(uintptr(signalFaviconDatabaseFaviconChangedId)))

	detail := signalFaviconDatabaseFaviconChangedDetail{callback, handlerID}
	signalFaviconDatabaseFaviconChangedMap[signalFaviconDatabaseFaviconChangedId] = detail

	return signalFaviconDatabaseFaviconChangedId
}

/*
DisconnectFaviconChanged disconnects a callback from the 'favicon-changed' signal for the FaviconDatabase.

The connectionID should be a value returned from a call to ConnectFaviconChanged.
*/
func (recv *FaviconDatabase) DisconnectFaviconChanged(connectionID int) {
	signalFaviconDatabaseFaviconChangedLock.Lock()
	defer signalFaviconDatabaseFaviconChangedLock.Unlock()

	detail, exists := signalFaviconDatabaseFaviconChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFaviconDatabaseFaviconChangedMap, connectionID)
}

//export favicondatabase_faviconChangedHandler
func favicondatabase_faviconChangedHandler(c_targetObject *C.GObject, c_page_uri *C.gchar, c_favicon_uri *C.gchar, data C.gpointer) {
	signalFaviconDatabaseFaviconChangedLock.RLock()
	defer signalFaviconDatabaseFaviconChangedLock.RUnlock()

	pageUri := C.GoString(c_page_uri)

	faviconUri := C.GoString(c_favicon_uri)

	targetObject := FaviconDatabaseNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalFaviconDatabaseFaviconChangedMap[index].callback
	callback(targetObject, pageUri, faviconUri)
}

// Clear is a wrapper around the C function webkit_favicon_database_clear.
func (recv *FaviconDatabase) Clear() {
	C.webkit_favicon_database_clear((*C.WebKitFaviconDatabase)(recv.native))

	return
}

// Unsupported : webkit_favicon_database_get_favicon : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// GetFaviconFinish is a wrapper around the C function webkit_favicon_database_get_favicon_finish.
func (recv *FaviconDatabase) GetFaviconFinish(result *gio.AsyncResult) (*cairo.Surface, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_favicon_database_get_favicon_finish((*C.WebKitFaviconDatabase)(recv.native), c_result, &cThrowableError)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetFaviconUri is a wrapper around the C function webkit_favicon_database_get_favicon_uri.
func (recv *FaviconDatabase) GetFaviconUri(pageUri string) string {
	c_page_uri := C.CString(pageUri)
	defer C.free(unsafe.Pointer(c_page_uri))

	retC := C.webkit_favicon_database_get_favicon_uri((*C.WebKitFaviconDatabase)(recv.native), c_page_uri)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FileChooserRequest is a wrapper around the C record WebKitFileChooserRequest.
type FileChooserRequest struct {
	native *C.WebKitFileChooserRequest
	// parent : record
	// Private : priv
}

func FileChooserRequestNewFromC(u unsafe.Pointer) *FileChooserRequest {
	c := (*C.WebKitFileChooserRequest)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FileChooserRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FileChooserRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileChooserRequest with another FileChooserRequest, and returns true if they represent the same GObject.
func (recv *FileChooserRequest) Equals(other *FileChooserRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FileChooserRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FileChooserRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileChooserRequest.
func CastToFileChooserRequest(object *gobject.Object) *FileChooserRequest {
	return FileChooserRequestNewFromC(object.ToC())
}

// Cancel is a wrapper around the C function webkit_file_chooser_request_cancel.
func (recv *FileChooserRequest) Cancel() {
	C.webkit_file_chooser_request_cancel((*C.WebKitFileChooserRequest)(recv.native))

	return
}

// GetMimeTypes is a wrapper around the C function webkit_file_chooser_request_get_mime_types.
func (recv *FileChooserRequest) GetMimeTypes() []string {
	retC := C.webkit_file_chooser_request_get_mime_types((*C.WebKitFileChooserRequest)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetMimeTypesFilter is a wrapper around the C function webkit_file_chooser_request_get_mime_types_filter.
func (recv *FileChooserRequest) GetMimeTypesFilter() *gtk.FileFilter {
	retC := C.webkit_file_chooser_request_get_mime_types_filter((*C.WebKitFileChooserRequest)(recv.native))
	retGo := gtk.FileFilterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectMultiple is a wrapper around the C function webkit_file_chooser_request_get_select_multiple.
func (recv *FileChooserRequest) GetSelectMultiple() bool {
	retC := C.webkit_file_chooser_request_get_select_multiple((*C.WebKitFileChooserRequest)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSelectedFiles is a wrapper around the C function webkit_file_chooser_request_get_selected_files.
func (recv *FileChooserRequest) GetSelectedFiles() []string {
	retC := C.webkit_file_chooser_request_get_selected_files((*C.WebKitFileChooserRequest)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// SelectFiles is a wrapper around the C function webkit_file_chooser_request_select_files.
func (recv *FileChooserRequest) SelectFiles(files []string) {
	c_files_array := make([]*C.gchar, len(files)+1, len(files)+1)
	for i, item := range files {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_files_array[i] = c
	}
	c_files_array[len(files)] = nil
	c_files_arrayPtr := &c_files_array[0]
	c_files := (**C.gchar)(unsafe.Pointer(c_files_arrayPtr))

	C.webkit_file_chooser_request_select_files((*C.WebKitFileChooserRequest)(recv.native), c_files)

	return
}

// FindController is a wrapper around the C record WebKitFindController.
type FindController struct {
	native *C.WebKitFindController
	// parent : record
	// Private : priv
}

func FindControllerNewFromC(u unsafe.Pointer) *FindController {
	c := (*C.WebKitFindController)(u)
	if c == nil {
		return nil
	}

	g := &FindController{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FindController) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FindController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FindController with another FindController, and returns true if they represent the same GObject.
func (recv *FindController) Equals(other *FindController) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FindController) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FindController.
// Exercise care, as this is a potentially dangerous function if the Object is not a FindController.
func CastToFindController(object *gobject.Object) *FindController {
	return FindControllerNewFromC(object.ToC())
}

type signalFindControllerCountedMatchesDetail struct {
	callback  FindControllerSignalCountedMatchesCallback
	handlerID C.gulong
}

var signalFindControllerCountedMatchesId int
var signalFindControllerCountedMatchesMap = make(map[int]signalFindControllerCountedMatchesDetail)
var signalFindControllerCountedMatchesLock sync.RWMutex

// FindControllerSignalCountedMatchesCallback is a callback function for a 'counted-matches' signal emitted from a FindController.
type FindControllerSignalCountedMatchesCallback func(targetObject *FindController, matchCount uint32)

/*
ConnectCountedMatches connects the callback to the 'counted-matches' signal for the FindController.

The returned value represents the connection, and may be passed to DisconnectCountedMatches to remove it.
*/
func (recv *FindController) ConnectCountedMatches(callback FindControllerSignalCountedMatchesCallback) int {
	signalFindControllerCountedMatchesLock.Lock()
	defer signalFindControllerCountedMatchesLock.Unlock()

	signalFindControllerCountedMatchesId++
	instance := C.gpointer(recv.native)
	handlerID := C.FindController_signal_connect_counted_matches(instance, C.gpointer(uintptr(signalFindControllerCountedMatchesId)))

	detail := signalFindControllerCountedMatchesDetail{callback, handlerID}
	signalFindControllerCountedMatchesMap[signalFindControllerCountedMatchesId] = detail

	return signalFindControllerCountedMatchesId
}

/*
DisconnectCountedMatches disconnects a callback from the 'counted-matches' signal for the FindController.

The connectionID should be a value returned from a call to ConnectCountedMatches.
*/
func (recv *FindController) DisconnectCountedMatches(connectionID int) {
	signalFindControllerCountedMatchesLock.Lock()
	defer signalFindControllerCountedMatchesLock.Unlock()

	detail, exists := signalFindControllerCountedMatchesMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFindControllerCountedMatchesMap, connectionID)
}

//export findcontroller_countedMatchesHandler
func findcontroller_countedMatchesHandler(c_targetObject *C.GObject, c_match_count C.guint, data C.gpointer) {
	signalFindControllerCountedMatchesLock.RLock()
	defer signalFindControllerCountedMatchesLock.RUnlock()

	matchCount := uint32(c_match_count)

	targetObject := FindControllerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalFindControllerCountedMatchesMap[index].callback
	callback(targetObject, matchCount)
}

type signalFindControllerFailedToFindTextDetail struct {
	callback  FindControllerSignalFailedToFindTextCallback
	handlerID C.gulong
}

var signalFindControllerFailedToFindTextId int
var signalFindControllerFailedToFindTextMap = make(map[int]signalFindControllerFailedToFindTextDetail)
var signalFindControllerFailedToFindTextLock sync.RWMutex

// FindControllerSignalFailedToFindTextCallback is a callback function for a 'failed-to-find-text' signal emitted from a FindController.
type FindControllerSignalFailedToFindTextCallback func(targetObject *FindController)

/*
ConnectFailedToFindText connects the callback to the 'failed-to-find-text' signal for the FindController.

The returned value represents the connection, and may be passed to DisconnectFailedToFindText to remove it.
*/
func (recv *FindController) ConnectFailedToFindText(callback FindControllerSignalFailedToFindTextCallback) int {
	signalFindControllerFailedToFindTextLock.Lock()
	defer signalFindControllerFailedToFindTextLock.Unlock()

	signalFindControllerFailedToFindTextId++
	instance := C.gpointer(recv.native)
	handlerID := C.FindController_signal_connect_failed_to_find_text(instance, C.gpointer(uintptr(signalFindControllerFailedToFindTextId)))

	detail := signalFindControllerFailedToFindTextDetail{callback, handlerID}
	signalFindControllerFailedToFindTextMap[signalFindControllerFailedToFindTextId] = detail

	return signalFindControllerFailedToFindTextId
}

/*
DisconnectFailedToFindText disconnects a callback from the 'failed-to-find-text' signal for the FindController.

The connectionID should be a value returned from a call to ConnectFailedToFindText.
*/
func (recv *FindController) DisconnectFailedToFindText(connectionID int) {
	signalFindControllerFailedToFindTextLock.Lock()
	defer signalFindControllerFailedToFindTextLock.Unlock()

	detail, exists := signalFindControllerFailedToFindTextMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFindControllerFailedToFindTextMap, connectionID)
}

//export findcontroller_failedToFindTextHandler
func findcontroller_failedToFindTextHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalFindControllerFailedToFindTextLock.RLock()
	defer signalFindControllerFailedToFindTextLock.RUnlock()

	targetObject := FindControllerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalFindControllerFailedToFindTextMap[index].callback
	callback(targetObject)
}

type signalFindControllerFoundTextDetail struct {
	callback  FindControllerSignalFoundTextCallback
	handlerID C.gulong
}

var signalFindControllerFoundTextId int
var signalFindControllerFoundTextMap = make(map[int]signalFindControllerFoundTextDetail)
var signalFindControllerFoundTextLock sync.RWMutex

// FindControllerSignalFoundTextCallback is a callback function for a 'found-text' signal emitted from a FindController.
type FindControllerSignalFoundTextCallback func(targetObject *FindController, matchCount uint32)

/*
ConnectFoundText connects the callback to the 'found-text' signal for the FindController.

The returned value represents the connection, and may be passed to DisconnectFoundText to remove it.
*/
func (recv *FindController) ConnectFoundText(callback FindControllerSignalFoundTextCallback) int {
	signalFindControllerFoundTextLock.Lock()
	defer signalFindControllerFoundTextLock.Unlock()

	signalFindControllerFoundTextId++
	instance := C.gpointer(recv.native)
	handlerID := C.FindController_signal_connect_found_text(instance, C.gpointer(uintptr(signalFindControllerFoundTextId)))

	detail := signalFindControllerFoundTextDetail{callback, handlerID}
	signalFindControllerFoundTextMap[signalFindControllerFoundTextId] = detail

	return signalFindControllerFoundTextId
}

/*
DisconnectFoundText disconnects a callback from the 'found-text' signal for the FindController.

The connectionID should be a value returned from a call to ConnectFoundText.
*/
func (recv *FindController) DisconnectFoundText(connectionID int) {
	signalFindControllerFoundTextLock.Lock()
	defer signalFindControllerFoundTextLock.Unlock()

	detail, exists := signalFindControllerFoundTextMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFindControllerFoundTextMap, connectionID)
}

//export findcontroller_foundTextHandler
func findcontroller_foundTextHandler(c_targetObject *C.GObject, c_match_count C.guint, data C.gpointer) {
	signalFindControllerFoundTextLock.RLock()
	defer signalFindControllerFoundTextLock.RUnlock()

	matchCount := uint32(c_match_count)

	targetObject := FindControllerNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalFindControllerFoundTextMap[index].callback
	callback(targetObject, matchCount)
}

// CountMatches is a wrapper around the C function webkit_find_controller_count_matches.
func (recv *FindController) CountMatches(searchText string, findOptions uint32, maxMatchCount uint32) {
	c_search_text := C.CString(searchText)
	defer C.free(unsafe.Pointer(c_search_text))

	c_find_options := (C.guint32)(findOptions)

	c_max_match_count := (C.guint)(maxMatchCount)

	C.webkit_find_controller_count_matches((*C.WebKitFindController)(recv.native), c_search_text, c_find_options, c_max_match_count)

	return
}

// GetMaxMatchCount is a wrapper around the C function webkit_find_controller_get_max_match_count.
func (recv *FindController) GetMaxMatchCount() uint32 {
	retC := C.webkit_find_controller_get_max_match_count((*C.WebKitFindController)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetOptions is a wrapper around the C function webkit_find_controller_get_options.
func (recv *FindController) GetOptions() uint32 {
	retC := C.webkit_find_controller_get_options((*C.WebKitFindController)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSearchText is a wrapper around the C function webkit_find_controller_get_search_text.
func (recv *FindController) GetSearchText() string {
	retC := C.webkit_find_controller_get_search_text((*C.WebKitFindController)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWebView is a wrapper around the C function webkit_find_controller_get_web_view.
func (recv *FindController) GetWebView() *WebView {
	retC := C.webkit_find_controller_get_web_view((*C.WebKitFindController)(recv.native))
	retGo := WebViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Search is a wrapper around the C function webkit_find_controller_search.
func (recv *FindController) Search(searchText string, findOptions uint32, maxMatchCount uint32) {
	c_search_text := C.CString(searchText)
	defer C.free(unsafe.Pointer(c_search_text))

	c_find_options := (C.guint32)(findOptions)

	c_max_match_count := (C.guint)(maxMatchCount)

	C.webkit_find_controller_search((*C.WebKitFindController)(recv.native), c_search_text, c_find_options, c_max_match_count)

	return
}

// SearchFinish is a wrapper around the C function webkit_find_controller_search_finish.
func (recv *FindController) SearchFinish() {
	C.webkit_find_controller_search_finish((*C.WebKitFindController)(recv.native))

	return
}

// SearchNext is a wrapper around the C function webkit_find_controller_search_next.
func (recv *FindController) SearchNext() {
	C.webkit_find_controller_search_next((*C.WebKitFindController)(recv.native))

	return
}

// SearchPrevious is a wrapper around the C function webkit_find_controller_search_previous.
func (recv *FindController) SearchPrevious() {
	C.webkit_find_controller_search_previous((*C.WebKitFindController)(recv.native))

	return
}

// FormSubmissionRequest is a wrapper around the C record WebKitFormSubmissionRequest.
type FormSubmissionRequest struct {
	native *C.WebKitFormSubmissionRequest
	// parent : record
	// Private : priv
}

func FormSubmissionRequestNewFromC(u unsafe.Pointer) *FormSubmissionRequest {
	c := (*C.WebKitFormSubmissionRequest)(u)
	if c == nil {
		return nil
	}

	g := &FormSubmissionRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *FormSubmissionRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *FormSubmissionRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FormSubmissionRequest with another FormSubmissionRequest, and returns true if they represent the same GObject.
func (recv *FormSubmissionRequest) Equals(other *FormSubmissionRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *FormSubmissionRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to FormSubmissionRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a FormSubmissionRequest.
func CastToFormSubmissionRequest(object *gobject.Object) *FormSubmissionRequest {
	return FormSubmissionRequestNewFromC(object.ToC())
}

// GetTextFields is a wrapper around the C function webkit_form_submission_request_get_text_fields.
func (recv *FormSubmissionRequest) GetTextFields() *glib.HashTable {
	retC := C.webkit_form_submission_request_get_text_fields((*C.WebKitFormSubmissionRequest)(recv.native))
	var retGo (*glib.HashTable)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.HashTableNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Submit is a wrapper around the C function webkit_form_submission_request_submit.
func (recv *FormSubmissionRequest) Submit() {
	C.webkit_form_submission_request_submit((*C.WebKitFormSubmissionRequest)(recv.native))

	return
}

// GeolocationManager is a wrapper around the C record WebKitGeolocationManager.
type GeolocationManager struct {
	native *C.WebKitGeolocationManager
	// parent : record
	// Private : priv
}

func GeolocationManagerNewFromC(u unsafe.Pointer) *GeolocationManager {
	c := (*C.WebKitGeolocationManager)(u)
	if c == nil {
		return nil
	}

	g := &GeolocationManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GeolocationManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GeolocationManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GeolocationManager with another GeolocationManager, and returns true if they represent the same GObject.
func (recv *GeolocationManager) Equals(other *GeolocationManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *GeolocationManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to GeolocationManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a GeolocationManager.
func CastToGeolocationManager(object *gobject.Object) *GeolocationManager {
	return GeolocationManagerNewFromC(object.ToC())
}

// GeolocationPermissionRequest is a wrapper around the C record WebKitGeolocationPermissionRequest.
type GeolocationPermissionRequest struct {
	native *C.WebKitGeolocationPermissionRequest
	// parent : record
	// Private : priv
}

func GeolocationPermissionRequestNewFromC(u unsafe.Pointer) *GeolocationPermissionRequest {
	c := (*C.WebKitGeolocationPermissionRequest)(u)
	if c == nil {
		return nil
	}

	g := &GeolocationPermissionRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *GeolocationPermissionRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *GeolocationPermissionRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GeolocationPermissionRequest with another GeolocationPermissionRequest, and returns true if they represent the same GObject.
func (recv *GeolocationPermissionRequest) Equals(other *GeolocationPermissionRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *GeolocationPermissionRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to GeolocationPermissionRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a GeolocationPermissionRequest.
func CastToGeolocationPermissionRequest(object *gobject.Object) *GeolocationPermissionRequest {
	return GeolocationPermissionRequestNewFromC(object.ToC())
}

// PermissionRequest returns the PermissionRequest interface implemented by GeolocationPermissionRequest
func (recv *GeolocationPermissionRequest) PermissionRequest() *PermissionRequest {
	return PermissionRequestNewFromC(recv.ToC())
}

// HitTestResult is a wrapper around the C record WebKitHitTestResult.
type HitTestResult struct {
	native *C.WebKitHitTestResult
	// parent : record
	// priv : record
}

func HitTestResultNewFromC(u unsafe.Pointer) *HitTestResult {
	c := (*C.WebKitHitTestResult)(u)
	if c == nil {
		return nil
	}

	g := &HitTestResult{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *HitTestResult) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *HitTestResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HitTestResult with another HitTestResult, and returns true if they represent the same GObject.
func (recv *HitTestResult) Equals(other *HitTestResult) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *HitTestResult) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to HitTestResult.
// Exercise care, as this is a potentially dangerous function if the Object is not a HitTestResult.
func CastToHitTestResult(object *gobject.Object) *HitTestResult {
	return HitTestResultNewFromC(object.ToC())
}

// ContextIsEditable is a wrapper around the C function webkit_hit_test_result_context_is_editable.
func (recv *HitTestResult) ContextIsEditable() bool {
	retC := C.webkit_hit_test_result_context_is_editable((*C.WebKitHitTestResult)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ContextIsImage is a wrapper around the C function webkit_hit_test_result_context_is_image.
func (recv *HitTestResult) ContextIsImage() bool {
	retC := C.webkit_hit_test_result_context_is_image((*C.WebKitHitTestResult)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ContextIsLink is a wrapper around the C function webkit_hit_test_result_context_is_link.
func (recv *HitTestResult) ContextIsLink() bool {
	retC := C.webkit_hit_test_result_context_is_link((*C.WebKitHitTestResult)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ContextIsMedia is a wrapper around the C function webkit_hit_test_result_context_is_media.
func (recv *HitTestResult) ContextIsMedia() bool {
	retC := C.webkit_hit_test_result_context_is_media((*C.WebKitHitTestResult)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ContextIsScrollbar is a wrapper around the C function webkit_hit_test_result_context_is_scrollbar.
func (recv *HitTestResult) ContextIsScrollbar() bool {
	retC := C.webkit_hit_test_result_context_is_scrollbar((*C.WebKitHitTestResult)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetContext is a wrapper around the C function webkit_hit_test_result_get_context.
func (recv *HitTestResult) GetContext() uint32 {
	retC := C.webkit_hit_test_result_get_context((*C.WebKitHitTestResult)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetImageUri is a wrapper around the C function webkit_hit_test_result_get_image_uri.
func (recv *HitTestResult) GetImageUri() string {
	retC := C.webkit_hit_test_result_get_image_uri((*C.WebKitHitTestResult)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLinkLabel is a wrapper around the C function webkit_hit_test_result_get_link_label.
func (recv *HitTestResult) GetLinkLabel() string {
	retC := C.webkit_hit_test_result_get_link_label((*C.WebKitHitTestResult)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLinkTitle is a wrapper around the C function webkit_hit_test_result_get_link_title.
func (recv *HitTestResult) GetLinkTitle() string {
	retC := C.webkit_hit_test_result_get_link_title((*C.WebKitHitTestResult)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLinkUri is a wrapper around the C function webkit_hit_test_result_get_link_uri.
func (recv *HitTestResult) GetLinkUri() string {
	retC := C.webkit_hit_test_result_get_link_uri((*C.WebKitHitTestResult)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMediaUri is a wrapper around the C function webkit_hit_test_result_get_media_uri.
func (recv *HitTestResult) GetMediaUri() string {
	retC := C.webkit_hit_test_result_get_media_uri((*C.WebKitHitTestResult)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// InstallMissingMediaPluginsPermissionRequest is a wrapper around the C record WebKitInstallMissingMediaPluginsPermissionRequest.
type InstallMissingMediaPluginsPermissionRequest struct {
	native *C.WebKitInstallMissingMediaPluginsPermissionRequest
	// parent : record
	// priv : record
}

func InstallMissingMediaPluginsPermissionRequestNewFromC(u unsafe.Pointer) *InstallMissingMediaPluginsPermissionRequest {
	c := (*C.WebKitInstallMissingMediaPluginsPermissionRequest)(u)
	if c == nil {
		return nil
	}

	g := &InstallMissingMediaPluginsPermissionRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InstallMissingMediaPluginsPermissionRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InstallMissingMediaPluginsPermissionRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InstallMissingMediaPluginsPermissionRequest with another InstallMissingMediaPluginsPermissionRequest, and returns true if they represent the same GObject.
func (recv *InstallMissingMediaPluginsPermissionRequest) Equals(other *InstallMissingMediaPluginsPermissionRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *InstallMissingMediaPluginsPermissionRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to InstallMissingMediaPluginsPermissionRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a InstallMissingMediaPluginsPermissionRequest.
func CastToInstallMissingMediaPluginsPermissionRequest(object *gobject.Object) *InstallMissingMediaPluginsPermissionRequest {
	return InstallMissingMediaPluginsPermissionRequestNewFromC(object.ToC())
}

// PermissionRequest returns the PermissionRequest interface implemented by InstallMissingMediaPluginsPermissionRequest
func (recv *InstallMissingMediaPluginsPermissionRequest) PermissionRequest() *PermissionRequest {
	return PermissionRequestNewFromC(recv.ToC())
}

// NavigationPolicyDecision is a wrapper around the C record WebKitNavigationPolicyDecision.
type NavigationPolicyDecision struct {
	native *C.WebKitNavigationPolicyDecision
	// parent : record
	// Private : priv
}

func NavigationPolicyDecisionNewFromC(u unsafe.Pointer) *NavigationPolicyDecision {
	c := (*C.WebKitNavigationPolicyDecision)(u)
	if c == nil {
		return nil
	}

	g := &NavigationPolicyDecision{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NavigationPolicyDecision) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NavigationPolicyDecision) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NavigationPolicyDecision with another NavigationPolicyDecision, and returns true if they represent the same GObject.
func (recv *NavigationPolicyDecision) Equals(other *NavigationPolicyDecision) bool {
	return other.ToC() == recv.ToC()
}

// PolicyDecision upcasts to *PolicyDecision
func (recv *NavigationPolicyDecision) PolicyDecision() *PolicyDecision {
	return PolicyDecisionNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *NavigationPolicyDecision) Object() *gobject.Object {
	return recv.PolicyDecision().Object()
}

// CastToWidget down casts any arbitrary Object to NavigationPolicyDecision.
// Exercise care, as this is a potentially dangerous function if the Object is not a NavigationPolicyDecision.
func CastToNavigationPolicyDecision(object *gobject.Object) *NavigationPolicyDecision {
	return NavigationPolicyDecisionNewFromC(object.ToC())
}

// GetFrameName is a wrapper around the C function webkit_navigation_policy_decision_get_frame_name.
func (recv *NavigationPolicyDecision) GetFrameName() string {
	retC := C.webkit_navigation_policy_decision_get_frame_name((*C.WebKitNavigationPolicyDecision)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetModifiers is a wrapper around the C function webkit_navigation_policy_decision_get_modifiers.
func (recv *NavigationPolicyDecision) GetModifiers() uint32 {
	retC := C.webkit_navigation_policy_decision_get_modifiers((*C.WebKitNavigationPolicyDecision)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMouseButton is a wrapper around the C function webkit_navigation_policy_decision_get_mouse_button.
func (recv *NavigationPolicyDecision) GetMouseButton() uint32 {
	retC := C.webkit_navigation_policy_decision_get_mouse_button((*C.WebKitNavigationPolicyDecision)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetNavigationType is a wrapper around the C function webkit_navigation_policy_decision_get_navigation_type.
func (recv *NavigationPolicyDecision) GetNavigationType() NavigationType {
	retC := C.webkit_navigation_policy_decision_get_navigation_type((*C.WebKitNavigationPolicyDecision)(recv.native))
	retGo := (NavigationType)(retC)

	return retGo
}

// GetRequest is a wrapper around the C function webkit_navigation_policy_decision_get_request.
func (recv *NavigationPolicyDecision) GetRequest() *URIRequest {
	retC := C.webkit_navigation_policy_decision_get_request((*C.WebKitNavigationPolicyDecision)(recv.native))
	retGo := URIRequestNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Notification is a wrapper around the C record WebKitNotification.
type Notification struct {
	native *C.WebKitNotification
	// parent : record
	// priv : record
}

func NotificationNewFromC(u unsafe.Pointer) *Notification {
	c := (*C.WebKitNotification)(u)
	if c == nil {
		return nil
	}

	g := &Notification{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Notification) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Notification) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Notification with another Notification, and returns true if they represent the same GObject.
func (recv *Notification) Equals(other *Notification) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Notification) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Notification.
// Exercise care, as this is a potentially dangerous function if the Object is not a Notification.
func CastToNotification(object *gobject.Object) *Notification {
	return NotificationNewFromC(object.ToC())
}

// NotificationPermissionRequest is a wrapper around the C record WebKitNotificationPermissionRequest.
type NotificationPermissionRequest struct {
	native *C.WebKitNotificationPermissionRequest
	// parent : record
	// Private : priv
}

func NotificationPermissionRequestNewFromC(u unsafe.Pointer) *NotificationPermissionRequest {
	c := (*C.WebKitNotificationPermissionRequest)(u)
	if c == nil {
		return nil
	}

	g := &NotificationPermissionRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *NotificationPermissionRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *NotificationPermissionRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NotificationPermissionRequest with another NotificationPermissionRequest, and returns true if they represent the same GObject.
func (recv *NotificationPermissionRequest) Equals(other *NotificationPermissionRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *NotificationPermissionRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to NotificationPermissionRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a NotificationPermissionRequest.
func CastToNotificationPermissionRequest(object *gobject.Object) *NotificationPermissionRequest {
	return NotificationPermissionRequestNewFromC(object.ToC())
}

// PermissionRequest returns the PermissionRequest interface implemented by NotificationPermissionRequest
func (recv *NotificationPermissionRequest) PermissionRequest() *PermissionRequest {
	return PermissionRequestNewFromC(recv.ToC())
}

// OptionMenu is a wrapper around the C record WebKitOptionMenu.
type OptionMenu struct {
	native *C.WebKitOptionMenu
	// parent : record
	// priv : record
}

func OptionMenuNewFromC(u unsafe.Pointer) *OptionMenu {
	c := (*C.WebKitOptionMenu)(u)
	if c == nil {
		return nil
	}

	g := &OptionMenu{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *OptionMenu) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *OptionMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionMenu with another OptionMenu, and returns true if they represent the same GObject.
func (recv *OptionMenu) Equals(other *OptionMenu) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *OptionMenu) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to OptionMenu.
// Exercise care, as this is a potentially dangerous function if the Object is not a OptionMenu.
func CastToOptionMenu(object *gobject.Object) *OptionMenu {
	return OptionMenuNewFromC(object.ToC())
}

// Plugin is a wrapper around the C record WebKitPlugin.
type Plugin struct {
	native *C.WebKitPlugin
	// parent : record
	// priv : record
}

func PluginNewFromC(u unsafe.Pointer) *Plugin {
	c := (*C.WebKitPlugin)(u)
	if c == nil {
		return nil
	}

	g := &Plugin{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Plugin) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Plugin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Plugin with another Plugin, and returns true if they represent the same GObject.
func (recv *Plugin) Equals(other *Plugin) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Plugin) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Plugin.
// Exercise care, as this is a potentially dangerous function if the Object is not a Plugin.
func CastToPlugin(object *gobject.Object) *Plugin {
	return PluginNewFromC(object.ToC())
}

// GetDescription is a wrapper around the C function webkit_plugin_get_description.
func (recv *Plugin) GetDescription() string {
	retC := C.webkit_plugin_get_description((*C.WebKitPlugin)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMimeInfoList is a wrapper around the C function webkit_plugin_get_mime_info_list.
func (recv *Plugin) GetMimeInfoList() *glib.List {
	retC := C.webkit_plugin_get_mime_info_list((*C.WebKitPlugin)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function webkit_plugin_get_name.
func (recv *Plugin) GetName() string {
	retC := C.webkit_plugin_get_name((*C.WebKitPlugin)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPath is a wrapper around the C function webkit_plugin_get_path.
func (recv *Plugin) GetPath() string {
	retC := C.webkit_plugin_get_path((*C.WebKitPlugin)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// PolicyDecision is a wrapper around the C record WebKitPolicyDecision.
type PolicyDecision struct {
	native *C.WebKitPolicyDecision
	// parent : record
	// Private : priv
}

func PolicyDecisionNewFromC(u unsafe.Pointer) *PolicyDecision {
	c := (*C.WebKitPolicyDecision)(u)
	if c == nil {
		return nil
	}

	g := &PolicyDecision{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PolicyDecision) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PolicyDecision) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PolicyDecision with another PolicyDecision, and returns true if they represent the same GObject.
func (recv *PolicyDecision) Equals(other *PolicyDecision) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PolicyDecision) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PolicyDecision.
// Exercise care, as this is a potentially dangerous function if the Object is not a PolicyDecision.
func CastToPolicyDecision(object *gobject.Object) *PolicyDecision {
	return PolicyDecisionNewFromC(object.ToC())
}

// Download is a wrapper around the C function webkit_policy_decision_download.
func (recv *PolicyDecision) Download() {
	C.webkit_policy_decision_download((*C.WebKitPolicyDecision)(recv.native))

	return
}

// Ignore is a wrapper around the C function webkit_policy_decision_ignore.
func (recv *PolicyDecision) Ignore() {
	C.webkit_policy_decision_ignore((*C.WebKitPolicyDecision)(recv.native))

	return
}

// Use is a wrapper around the C function webkit_policy_decision_use.
func (recv *PolicyDecision) Use() {
	C.webkit_policy_decision_use((*C.WebKitPolicyDecision)(recv.native))

	return
}

// PrintCustomWidget is a wrapper around the C record WebKitPrintCustomWidget.
type PrintCustomWidget struct {
	native *C.WebKitPrintCustomWidget
	// parent : record
	// priv : record
}

func PrintCustomWidgetNewFromC(u unsafe.Pointer) *PrintCustomWidget {
	c := (*C.WebKitPrintCustomWidget)(u)
	if c == nil {
		return nil
	}

	g := &PrintCustomWidget{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PrintCustomWidget) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PrintCustomWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PrintCustomWidget with another PrintCustomWidget, and returns true if they represent the same GObject.
func (recv *PrintCustomWidget) Equals(other *PrintCustomWidget) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PrintCustomWidget) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PrintCustomWidget.
// Exercise care, as this is a potentially dangerous function if the Object is not a PrintCustomWidget.
func CastToPrintCustomWidget(object *gobject.Object) *PrintCustomWidget {
	return PrintCustomWidgetNewFromC(object.ToC())
}

// PrintOperation is a wrapper around the C record WebKitPrintOperation.
type PrintOperation struct {
	native *C.WebKitPrintOperation
	// parent : record
	// priv : record
}

func PrintOperationNewFromC(u unsafe.Pointer) *PrintOperation {
	c := (*C.WebKitPrintOperation)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperation{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *PrintOperation) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *PrintOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PrintOperation with another PrintOperation, and returns true if they represent the same GObject.
func (recv *PrintOperation) Equals(other *PrintOperation) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *PrintOperation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to PrintOperation.
// Exercise care, as this is a potentially dangerous function if the Object is not a PrintOperation.
func CastToPrintOperation(object *gobject.Object) *PrintOperation {
	return PrintOperationNewFromC(object.ToC())
}

// Blacklisted signal create-custom-widget

type signalPrintOperationFailedDetail struct {
	callback  PrintOperationSignalFailedCallback
	handlerID C.gulong
}

var signalPrintOperationFailedId int
var signalPrintOperationFailedMap = make(map[int]signalPrintOperationFailedDetail)
var signalPrintOperationFailedLock sync.RWMutex

// PrintOperationSignalFailedCallback is a callback function for a 'failed' signal emitted from a PrintOperation.
type PrintOperationSignalFailedCallback func(targetObject *PrintOperation, error *glib.Error)

/*
ConnectFailed connects the callback to the 'failed' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectFailed to remove it.
*/
func (recv *PrintOperation) ConnectFailed(callback PrintOperationSignalFailedCallback) int {
	signalPrintOperationFailedLock.Lock()
	defer signalPrintOperationFailedLock.Unlock()

	signalPrintOperationFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_failed(instance, C.gpointer(uintptr(signalPrintOperationFailedId)))

	detail := signalPrintOperationFailedDetail{callback, handlerID}
	signalPrintOperationFailedMap[signalPrintOperationFailedId] = detail

	return signalPrintOperationFailedId
}

/*
DisconnectFailed disconnects a callback from the 'failed' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectFailed.
*/
func (recv *PrintOperation) DisconnectFailed(connectionID int) {
	signalPrintOperationFailedLock.Lock()
	defer signalPrintOperationFailedLock.Unlock()

	detail, exists := signalPrintOperationFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationFailedMap, connectionID)
}

//export printoperation_failedHandler
func printoperation_failedHandler(c_targetObject *C.GObject, c_error *C.GError, data C.gpointer) {
	signalPrintOperationFailedLock.RLock()
	defer signalPrintOperationFailedLock.RUnlock()

	error := glib.ErrorNewFromC(unsafe.Pointer(c_error))

	targetObject := PrintOperationNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalPrintOperationFailedMap[index].callback
	callback(targetObject, error)
}

type signalPrintOperationFinishedDetail struct {
	callback  PrintOperationSignalFinishedCallback
	handlerID C.gulong
}

var signalPrintOperationFinishedId int
var signalPrintOperationFinishedMap = make(map[int]signalPrintOperationFinishedDetail)
var signalPrintOperationFinishedLock sync.RWMutex

// PrintOperationSignalFinishedCallback is a callback function for a 'finished' signal emitted from a PrintOperation.
type PrintOperationSignalFinishedCallback func(targetObject *PrintOperation)

/*
ConnectFinished connects the callback to the 'finished' signal for the PrintOperation.

The returned value represents the connection, and may be passed to DisconnectFinished to remove it.
*/
func (recv *PrintOperation) ConnectFinished(callback PrintOperationSignalFinishedCallback) int {
	signalPrintOperationFinishedLock.Lock()
	defer signalPrintOperationFinishedLock.Unlock()

	signalPrintOperationFinishedId++
	instance := C.gpointer(recv.native)
	handlerID := C.PrintOperation_signal_connect_finished(instance, C.gpointer(uintptr(signalPrintOperationFinishedId)))

	detail := signalPrintOperationFinishedDetail{callback, handlerID}
	signalPrintOperationFinishedMap[signalPrintOperationFinishedId] = detail

	return signalPrintOperationFinishedId
}

/*
DisconnectFinished disconnects a callback from the 'finished' signal for the PrintOperation.

The connectionID should be a value returned from a call to ConnectFinished.
*/
func (recv *PrintOperation) DisconnectFinished(connectionID int) {
	signalPrintOperationFinishedLock.Lock()
	defer signalPrintOperationFinishedLock.Unlock()

	detail, exists := signalPrintOperationFinishedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPrintOperationFinishedMap, connectionID)
}

//export printoperation_finishedHandler
func printoperation_finishedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalPrintOperationFinishedLock.RLock()
	defer signalPrintOperationFinishedLock.RUnlock()

	targetObject := PrintOperationNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalPrintOperationFinishedMap[index].callback
	callback(targetObject)
}

// PrintOperationNew is a wrapper around the C function webkit_print_operation_new.
func PrintOperationNew(webView *WebView) *PrintOperation {
	c_web_view := (*C.WebKitWebView)(C.NULL)
	if webView != nil {
		c_web_view = (*C.WebKitWebView)(webView.ToC())
	}

	retC := C.webkit_print_operation_new(c_web_view)
	retGo := PrintOperationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetPageSetup is a wrapper around the C function webkit_print_operation_get_page_setup.
func (recv *PrintOperation) GetPageSetup() *gtk.PageSetup {
	retC := C.webkit_print_operation_get_page_setup((*C.WebKitPrintOperation)(recv.native))
	retGo := gtk.PageSetupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPrintSettings is a wrapper around the C function webkit_print_operation_get_print_settings.
func (recv *PrintOperation) GetPrintSettings() *gtk.PrintSettings {
	retC := C.webkit_print_operation_get_print_settings((*C.WebKitPrintOperation)(recv.native))
	retGo := gtk.PrintSettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Print is a wrapper around the C function webkit_print_operation_print.
func (recv *PrintOperation) Print() {
	C.webkit_print_operation_print((*C.WebKitPrintOperation)(recv.native))

	return
}

// RunDialog is a wrapper around the C function webkit_print_operation_run_dialog.
func (recv *PrintOperation) RunDialog(parent *gtk.Window) PrintOperationResponse {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	retC := C.webkit_print_operation_run_dialog((*C.WebKitPrintOperation)(recv.native), c_parent)
	retGo := (PrintOperationResponse)(retC)

	return retGo
}

// SetPageSetup is a wrapper around the C function webkit_print_operation_set_page_setup.
func (recv *PrintOperation) SetPageSetup(pageSetup *gtk.PageSetup) {
	c_page_setup := (*C.GtkPageSetup)(C.NULL)
	if pageSetup != nil {
		c_page_setup = (*C.GtkPageSetup)(pageSetup.ToC())
	}

	C.webkit_print_operation_set_page_setup((*C.WebKitPrintOperation)(recv.native), c_page_setup)

	return
}

// SetPrintSettings is a wrapper around the C function webkit_print_operation_set_print_settings.
func (recv *PrintOperation) SetPrintSettings(printSettings *gtk.PrintSettings) {
	c_print_settings := (*C.GtkPrintSettings)(C.NULL)
	if printSettings != nil {
		c_print_settings = (*C.GtkPrintSettings)(printSettings.ToC())
	}

	C.webkit_print_operation_set_print_settings((*C.WebKitPrintOperation)(recv.native), c_print_settings)

	return
}

// ResponsePolicyDecision is a wrapper around the C record WebKitResponsePolicyDecision.
type ResponsePolicyDecision struct {
	native *C.WebKitResponsePolicyDecision
	// parent : record
	// Private : priv
}

func ResponsePolicyDecisionNewFromC(u unsafe.Pointer) *ResponsePolicyDecision {
	c := (*C.WebKitResponsePolicyDecision)(u)
	if c == nil {
		return nil
	}

	g := &ResponsePolicyDecision{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ResponsePolicyDecision) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ResponsePolicyDecision) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ResponsePolicyDecision with another ResponsePolicyDecision, and returns true if they represent the same GObject.
func (recv *ResponsePolicyDecision) Equals(other *ResponsePolicyDecision) bool {
	return other.ToC() == recv.ToC()
}

// PolicyDecision upcasts to *PolicyDecision
func (recv *ResponsePolicyDecision) PolicyDecision() *PolicyDecision {
	return PolicyDecisionNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *ResponsePolicyDecision) Object() *gobject.Object {
	return recv.PolicyDecision().Object()
}

// CastToWidget down casts any arbitrary Object to ResponsePolicyDecision.
// Exercise care, as this is a potentially dangerous function if the Object is not a ResponsePolicyDecision.
func CastToResponsePolicyDecision(object *gobject.Object) *ResponsePolicyDecision {
	return ResponsePolicyDecisionNewFromC(object.ToC())
}

// GetRequest is a wrapper around the C function webkit_response_policy_decision_get_request.
func (recv *ResponsePolicyDecision) GetRequest() *URIRequest {
	retC := C.webkit_response_policy_decision_get_request((*C.WebKitResponsePolicyDecision)(recv.native))
	retGo := URIRequestNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetResponse is a wrapper around the C function webkit_response_policy_decision_get_response.
func (recv *ResponsePolicyDecision) GetResponse() *URIResponse {
	retC := C.webkit_response_policy_decision_get_response((*C.WebKitResponsePolicyDecision)(recv.native))
	retGo := URIResponseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SecurityManager is a wrapper around the C record WebKitSecurityManager.
type SecurityManager struct {
	native *C.WebKitSecurityManager
	// parent : record
	// priv : record
}

func SecurityManagerNewFromC(u unsafe.Pointer) *SecurityManager {
	c := (*C.WebKitSecurityManager)(u)
	if c == nil {
		return nil
	}

	g := &SecurityManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SecurityManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SecurityManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SecurityManager with another SecurityManager, and returns true if they represent the same GObject.
func (recv *SecurityManager) Equals(other *SecurityManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SecurityManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SecurityManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a SecurityManager.
func CastToSecurityManager(object *gobject.Object) *SecurityManager {
	return SecurityManagerNewFromC(object.ToC())
}

// RegisterUriSchemeAsCorsEnabled is a wrapper around the C function webkit_security_manager_register_uri_scheme_as_cors_enabled.
func (recv *SecurityManager) RegisterUriSchemeAsCorsEnabled(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.webkit_security_manager_register_uri_scheme_as_cors_enabled((*C.WebKitSecurityManager)(recv.native), c_scheme)

	return
}

// RegisterUriSchemeAsDisplayIsolated is a wrapper around the C function webkit_security_manager_register_uri_scheme_as_display_isolated.
func (recv *SecurityManager) RegisterUriSchemeAsDisplayIsolated(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.webkit_security_manager_register_uri_scheme_as_display_isolated((*C.WebKitSecurityManager)(recv.native), c_scheme)

	return
}

// RegisterUriSchemeAsEmptyDocument is a wrapper around the C function webkit_security_manager_register_uri_scheme_as_empty_document.
func (recv *SecurityManager) RegisterUriSchemeAsEmptyDocument(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.webkit_security_manager_register_uri_scheme_as_empty_document((*C.WebKitSecurityManager)(recv.native), c_scheme)

	return
}

// RegisterUriSchemeAsLocal is a wrapper around the C function webkit_security_manager_register_uri_scheme_as_local.
func (recv *SecurityManager) RegisterUriSchemeAsLocal(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.webkit_security_manager_register_uri_scheme_as_local((*C.WebKitSecurityManager)(recv.native), c_scheme)

	return
}

// RegisterUriSchemeAsNoAccess is a wrapper around the C function webkit_security_manager_register_uri_scheme_as_no_access.
func (recv *SecurityManager) RegisterUriSchemeAsNoAccess(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.webkit_security_manager_register_uri_scheme_as_no_access((*C.WebKitSecurityManager)(recv.native), c_scheme)

	return
}

// RegisterUriSchemeAsSecure is a wrapper around the C function webkit_security_manager_register_uri_scheme_as_secure.
func (recv *SecurityManager) RegisterUriSchemeAsSecure(scheme string) {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	C.webkit_security_manager_register_uri_scheme_as_secure((*C.WebKitSecurityManager)(recv.native), c_scheme)

	return
}

// UriSchemeIsCorsEnabled is a wrapper around the C function webkit_security_manager_uri_scheme_is_cors_enabled.
func (recv *SecurityManager) UriSchemeIsCorsEnabled(scheme string) bool {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	retC := C.webkit_security_manager_uri_scheme_is_cors_enabled((*C.WebKitSecurityManager)(recv.native), c_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// UriSchemeIsDisplayIsolated is a wrapper around the C function webkit_security_manager_uri_scheme_is_display_isolated.
func (recv *SecurityManager) UriSchemeIsDisplayIsolated(scheme string) bool {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	retC := C.webkit_security_manager_uri_scheme_is_display_isolated((*C.WebKitSecurityManager)(recv.native), c_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// UriSchemeIsEmptyDocument is a wrapper around the C function webkit_security_manager_uri_scheme_is_empty_document.
func (recv *SecurityManager) UriSchemeIsEmptyDocument(scheme string) bool {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	retC := C.webkit_security_manager_uri_scheme_is_empty_document((*C.WebKitSecurityManager)(recv.native), c_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// UriSchemeIsLocal is a wrapper around the C function webkit_security_manager_uri_scheme_is_local.
func (recv *SecurityManager) UriSchemeIsLocal(scheme string) bool {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	retC := C.webkit_security_manager_uri_scheme_is_local((*C.WebKitSecurityManager)(recv.native), c_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// UriSchemeIsNoAccess is a wrapper around the C function webkit_security_manager_uri_scheme_is_no_access.
func (recv *SecurityManager) UriSchemeIsNoAccess(scheme string) bool {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	retC := C.webkit_security_manager_uri_scheme_is_no_access((*C.WebKitSecurityManager)(recv.native), c_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// UriSchemeIsSecure is a wrapper around the C function webkit_security_manager_uri_scheme_is_secure.
func (recv *SecurityManager) UriSchemeIsSecure(scheme string) bool {
	c_scheme := C.CString(scheme)
	defer C.free(unsafe.Pointer(c_scheme))

	retC := C.webkit_security_manager_uri_scheme_is_secure((*C.WebKitSecurityManager)(recv.native), c_scheme)
	retGo := retC == C.TRUE

	return retGo
}

// Settings is a wrapper around the C record WebKitSettings.
type Settings struct {
	native *C.WebKitSettings
	// parent_instance : record
	// priv : record
}

func SettingsNewFromC(u unsafe.Pointer) *Settings {
	c := (*C.WebKitSettings)(u)
	if c == nil {
		return nil
	}

	g := &Settings{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Settings) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Settings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Settings with another Settings, and returns true if they represent the same GObject.
func (recv *Settings) Equals(other *Settings) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Settings) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Settings.
// Exercise care, as this is a potentially dangerous function if the Object is not a Settings.
func CastToSettings(object *gobject.Object) *Settings {
	return SettingsNewFromC(object.ToC())
}

// SettingsNew is a wrapper around the C function webkit_settings_new.
func SettingsNew() *Settings {
	retC := C.webkit_settings_new()
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : webkit_settings_new_with_settings : unsupported parameter ... : varargs

// GetAllowModalDialogs is a wrapper around the C function webkit_settings_get_allow_modal_dialogs.
func (recv *Settings) GetAllowModalDialogs() bool {
	retC := C.webkit_settings_get_allow_modal_dialogs((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetAutoLoadImages is a wrapper around the C function webkit_settings_get_auto_load_images.
func (recv *Settings) GetAutoLoadImages() bool {
	retC := C.webkit_settings_get_auto_load_images((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCursiveFontFamily is a wrapper around the C function webkit_settings_get_cursive_font_family.
func (recv *Settings) GetCursiveFontFamily() string {
	retC := C.webkit_settings_get_cursive_font_family((*C.WebKitSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDefaultCharset is a wrapper around the C function webkit_settings_get_default_charset.
func (recv *Settings) GetDefaultCharset() string {
	retC := C.webkit_settings_get_default_charset((*C.WebKitSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDefaultFontFamily is a wrapper around the C function webkit_settings_get_default_font_family.
func (recv *Settings) GetDefaultFontFamily() string {
	retC := C.webkit_settings_get_default_font_family((*C.WebKitSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetDefaultFontSize is a wrapper around the C function webkit_settings_get_default_font_size.
func (recv *Settings) GetDefaultFontSize() uint32 {
	retC := C.webkit_settings_get_default_font_size((*C.WebKitSettings)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetDefaultMonospaceFontSize is a wrapper around the C function webkit_settings_get_default_monospace_font_size.
func (recv *Settings) GetDefaultMonospaceFontSize() uint32 {
	retC := C.webkit_settings_get_default_monospace_font_size((*C.WebKitSettings)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetDrawCompositingIndicators is a wrapper around the C function webkit_settings_get_draw_compositing_indicators.
func (recv *Settings) GetDrawCompositingIndicators() bool {
	retC := C.webkit_settings_get_draw_compositing_indicators((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableCaretBrowsing is a wrapper around the C function webkit_settings_get_enable_caret_browsing.
func (recv *Settings) GetEnableCaretBrowsing() bool {
	retC := C.webkit_settings_get_enable_caret_browsing((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableDeveloperExtras is a wrapper around the C function webkit_settings_get_enable_developer_extras.
func (recv *Settings) GetEnableDeveloperExtras() bool {
	retC := C.webkit_settings_get_enable_developer_extras((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableDnsPrefetching is a wrapper around the C function webkit_settings_get_enable_dns_prefetching.
func (recv *Settings) GetEnableDnsPrefetching() bool {
	retC := C.webkit_settings_get_enable_dns_prefetching((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableFrameFlattening is a wrapper around the C function webkit_settings_get_enable_frame_flattening.
func (recv *Settings) GetEnableFrameFlattening() bool {
	retC := C.webkit_settings_get_enable_frame_flattening((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableFullscreen is a wrapper around the C function webkit_settings_get_enable_fullscreen.
func (recv *Settings) GetEnableFullscreen() bool {
	retC := C.webkit_settings_get_enable_fullscreen((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableHtml5Database is a wrapper around the C function webkit_settings_get_enable_html5_database.
func (recv *Settings) GetEnableHtml5Database() bool {
	retC := C.webkit_settings_get_enable_html5_database((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableHtml5LocalStorage is a wrapper around the C function webkit_settings_get_enable_html5_local_storage.
func (recv *Settings) GetEnableHtml5LocalStorage() bool {
	retC := C.webkit_settings_get_enable_html5_local_storage((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableHyperlinkAuditing is a wrapper around the C function webkit_settings_get_enable_hyperlink_auditing.
func (recv *Settings) GetEnableHyperlinkAuditing() bool {
	retC := C.webkit_settings_get_enable_hyperlink_auditing((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableJava is a wrapper around the C function webkit_settings_get_enable_java.
func (recv *Settings) GetEnableJava() bool {
	retC := C.webkit_settings_get_enable_java((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableJavascript is a wrapper around the C function webkit_settings_get_enable_javascript.
func (recv *Settings) GetEnableJavascript() bool {
	retC := C.webkit_settings_get_enable_javascript((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableOfflineWebApplicationCache is a wrapper around the C function webkit_settings_get_enable_offline_web_application_cache.
func (recv *Settings) GetEnableOfflineWebApplicationCache() bool {
	retC := C.webkit_settings_get_enable_offline_web_application_cache((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnablePageCache is a wrapper around the C function webkit_settings_get_enable_page_cache.
func (recv *Settings) GetEnablePageCache() bool {
	retC := C.webkit_settings_get_enable_page_cache((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnablePlugins is a wrapper around the C function webkit_settings_get_enable_plugins.
func (recv *Settings) GetEnablePlugins() bool {
	retC := C.webkit_settings_get_enable_plugins((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnablePrivateBrowsing is a wrapper around the C function webkit_settings_get_enable_private_browsing.
func (recv *Settings) GetEnablePrivateBrowsing() bool {
	retC := C.webkit_settings_get_enable_private_browsing((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableResizableTextAreas is a wrapper around the C function webkit_settings_get_enable_resizable_text_areas.
func (recv *Settings) GetEnableResizableTextAreas() bool {
	retC := C.webkit_settings_get_enable_resizable_text_areas((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableSiteSpecificQuirks is a wrapper around the C function webkit_settings_get_enable_site_specific_quirks.
func (recv *Settings) GetEnableSiteSpecificQuirks() bool {
	retC := C.webkit_settings_get_enable_site_specific_quirks((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableSmoothScrolling is a wrapper around the C function webkit_settings_get_enable_smooth_scrolling.
func (recv *Settings) GetEnableSmoothScrolling() bool {
	retC := C.webkit_settings_get_enable_smooth_scrolling((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableTabsToLinks is a wrapper around the C function webkit_settings_get_enable_tabs_to_links.
func (recv *Settings) GetEnableTabsToLinks() bool {
	retC := C.webkit_settings_get_enable_tabs_to_links((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableWebaudio is a wrapper around the C function webkit_settings_get_enable_webaudio.
func (recv *Settings) GetEnableWebaudio() bool {
	retC := C.webkit_settings_get_enable_webaudio((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableWebgl is a wrapper around the C function webkit_settings_get_enable_webgl.
func (recv *Settings) GetEnableWebgl() bool {
	retC := C.webkit_settings_get_enable_webgl((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetEnableXssAuditor is a wrapper around the C function webkit_settings_get_enable_xss_auditor.
func (recv *Settings) GetEnableXssAuditor() bool {
	retC := C.webkit_settings_get_enable_xss_auditor((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFantasyFontFamily is a wrapper around the C function webkit_settings_get_fantasy_font_family.
func (recv *Settings) GetFantasyFontFamily() string {
	retC := C.webkit_settings_get_fantasy_font_family((*C.WebKitSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetJavascriptCanAccessClipboard is a wrapper around the C function webkit_settings_get_javascript_can_access_clipboard.
func (recv *Settings) GetJavascriptCanAccessClipboard() bool {
	retC := C.webkit_settings_get_javascript_can_access_clipboard((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetJavascriptCanOpenWindowsAutomatically is a wrapper around the C function webkit_settings_get_javascript_can_open_windows_automatically.
func (recv *Settings) GetJavascriptCanOpenWindowsAutomatically() bool {
	retC := C.webkit_settings_get_javascript_can_open_windows_automatically((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetLoadIconsIgnoringImageLoadSetting is a wrapper around the C function webkit_settings_get_load_icons_ignoring_image_load_setting.
func (recv *Settings) GetLoadIconsIgnoringImageLoadSetting() bool {
	retC := C.webkit_settings_get_load_icons_ignoring_image_load_setting((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMediaPlaybackAllowsInline is a wrapper around the C function webkit_settings_get_media_playback_allows_inline.
func (recv *Settings) GetMediaPlaybackAllowsInline() bool {
	retC := C.webkit_settings_get_media_playback_allows_inline((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMediaPlaybackRequiresUserGesture is a wrapper around the C function webkit_settings_get_media_playback_requires_user_gesture.
func (recv *Settings) GetMediaPlaybackRequiresUserGesture() bool {
	retC := C.webkit_settings_get_media_playback_requires_user_gesture((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMinimumFontSize is a wrapper around the C function webkit_settings_get_minimum_font_size.
func (recv *Settings) GetMinimumFontSize() uint32 {
	retC := C.webkit_settings_get_minimum_font_size((*C.WebKitSettings)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMonospaceFontFamily is a wrapper around the C function webkit_settings_get_monospace_font_family.
func (recv *Settings) GetMonospaceFontFamily() string {
	retC := C.webkit_settings_get_monospace_font_family((*C.WebKitSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPictographFontFamily is a wrapper around the C function webkit_settings_get_pictograph_font_family.
func (recv *Settings) GetPictographFontFamily() string {
	retC := C.webkit_settings_get_pictograph_font_family((*C.WebKitSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPrintBackgrounds is a wrapper around the C function webkit_settings_get_print_backgrounds.
func (recv *Settings) GetPrintBackgrounds() bool {
	retC := C.webkit_settings_get_print_backgrounds((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSansSerifFontFamily is a wrapper around the C function webkit_settings_get_sans_serif_font_family.
func (recv *Settings) GetSansSerifFontFamily() string {
	retC := C.webkit_settings_get_sans_serif_font_family((*C.WebKitSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSerifFontFamily is a wrapper around the C function webkit_settings_get_serif_font_family.
func (recv *Settings) GetSerifFontFamily() string {
	retC := C.webkit_settings_get_serif_font_family((*C.WebKitSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUserAgent is a wrapper around the C function webkit_settings_get_user_agent.
func (recv *Settings) GetUserAgent() string {
	retC := C.webkit_settings_get_user_agent((*C.WebKitSettings)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetZoomTextOnly is a wrapper around the C function webkit_settings_get_zoom_text_only.
func (recv *Settings) GetZoomTextOnly() bool {
	retC := C.webkit_settings_get_zoom_text_only((*C.WebKitSettings)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAllowModalDialogs is a wrapper around the C function webkit_settings_set_allow_modal_dialogs.
func (recv *Settings) SetAllowModalDialogs(allowed bool) {
	c_allowed :=
		boolToGboolean(allowed)

	C.webkit_settings_set_allow_modal_dialogs((*C.WebKitSettings)(recv.native), c_allowed)

	return
}

// SetAutoLoadImages is a wrapper around the C function webkit_settings_set_auto_load_images.
func (recv *Settings) SetAutoLoadImages(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_auto_load_images((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetCursiveFontFamily is a wrapper around the C function webkit_settings_set_cursive_font_family.
func (recv *Settings) SetCursiveFontFamily(cursiveFontFamily string) {
	c_cursive_font_family := C.CString(cursiveFontFamily)
	defer C.free(unsafe.Pointer(c_cursive_font_family))

	C.webkit_settings_set_cursive_font_family((*C.WebKitSettings)(recv.native), c_cursive_font_family)

	return
}

// SetDefaultCharset is a wrapper around the C function webkit_settings_set_default_charset.
func (recv *Settings) SetDefaultCharset(defaultCharset string) {
	c_default_charset := C.CString(defaultCharset)
	defer C.free(unsafe.Pointer(c_default_charset))

	C.webkit_settings_set_default_charset((*C.WebKitSettings)(recv.native), c_default_charset)

	return
}

// SetDefaultFontFamily is a wrapper around the C function webkit_settings_set_default_font_family.
func (recv *Settings) SetDefaultFontFamily(defaultFontFamily string) {
	c_default_font_family := C.CString(defaultFontFamily)
	defer C.free(unsafe.Pointer(c_default_font_family))

	C.webkit_settings_set_default_font_family((*C.WebKitSettings)(recv.native), c_default_font_family)

	return
}

// SetDefaultFontSize is a wrapper around the C function webkit_settings_set_default_font_size.
func (recv *Settings) SetDefaultFontSize(fontSize uint32) {
	c_font_size := (C.guint32)(fontSize)

	C.webkit_settings_set_default_font_size((*C.WebKitSettings)(recv.native), c_font_size)

	return
}

// SetDefaultMonospaceFontSize is a wrapper around the C function webkit_settings_set_default_monospace_font_size.
func (recv *Settings) SetDefaultMonospaceFontSize(fontSize uint32) {
	c_font_size := (C.guint32)(fontSize)

	C.webkit_settings_set_default_monospace_font_size((*C.WebKitSettings)(recv.native), c_font_size)

	return
}

// SetDrawCompositingIndicators is a wrapper around the C function webkit_settings_set_draw_compositing_indicators.
func (recv *Settings) SetDrawCompositingIndicators(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_draw_compositing_indicators((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableCaretBrowsing is a wrapper around the C function webkit_settings_set_enable_caret_browsing.
func (recv *Settings) SetEnableCaretBrowsing(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_caret_browsing((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableDeveloperExtras is a wrapper around the C function webkit_settings_set_enable_developer_extras.
func (recv *Settings) SetEnableDeveloperExtras(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_developer_extras((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableDnsPrefetching is a wrapper around the C function webkit_settings_set_enable_dns_prefetching.
func (recv *Settings) SetEnableDnsPrefetching(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_dns_prefetching((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableFrameFlattening is a wrapper around the C function webkit_settings_set_enable_frame_flattening.
func (recv *Settings) SetEnableFrameFlattening(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_frame_flattening((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableFullscreen is a wrapper around the C function webkit_settings_set_enable_fullscreen.
func (recv *Settings) SetEnableFullscreen(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_fullscreen((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableHtml5Database is a wrapper around the C function webkit_settings_set_enable_html5_database.
func (recv *Settings) SetEnableHtml5Database(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_html5_database((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableHtml5LocalStorage is a wrapper around the C function webkit_settings_set_enable_html5_local_storage.
func (recv *Settings) SetEnableHtml5LocalStorage(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_html5_local_storage((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableHyperlinkAuditing is a wrapper around the C function webkit_settings_set_enable_hyperlink_auditing.
func (recv *Settings) SetEnableHyperlinkAuditing(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_hyperlink_auditing((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableJava is a wrapper around the C function webkit_settings_set_enable_java.
func (recv *Settings) SetEnableJava(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_java((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableJavascript is a wrapper around the C function webkit_settings_set_enable_javascript.
func (recv *Settings) SetEnableJavascript(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_javascript((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableOfflineWebApplicationCache is a wrapper around the C function webkit_settings_set_enable_offline_web_application_cache.
func (recv *Settings) SetEnableOfflineWebApplicationCache(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_offline_web_application_cache((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnablePageCache is a wrapper around the C function webkit_settings_set_enable_page_cache.
func (recv *Settings) SetEnablePageCache(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_page_cache((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnablePlugins is a wrapper around the C function webkit_settings_set_enable_plugins.
func (recv *Settings) SetEnablePlugins(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_plugins((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnablePrivateBrowsing is a wrapper around the C function webkit_settings_set_enable_private_browsing.
func (recv *Settings) SetEnablePrivateBrowsing(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_private_browsing((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableResizableTextAreas is a wrapper around the C function webkit_settings_set_enable_resizable_text_areas.
func (recv *Settings) SetEnableResizableTextAreas(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_resizable_text_areas((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableSiteSpecificQuirks is a wrapper around the C function webkit_settings_set_enable_site_specific_quirks.
func (recv *Settings) SetEnableSiteSpecificQuirks(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_site_specific_quirks((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableSmoothScrolling is a wrapper around the C function webkit_settings_set_enable_smooth_scrolling.
func (recv *Settings) SetEnableSmoothScrolling(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_smooth_scrolling((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableTabsToLinks is a wrapper around the C function webkit_settings_set_enable_tabs_to_links.
func (recv *Settings) SetEnableTabsToLinks(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_tabs_to_links((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableWebaudio is a wrapper around the C function webkit_settings_set_enable_webaudio.
func (recv *Settings) SetEnableWebaudio(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_webaudio((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableWebgl is a wrapper around the C function webkit_settings_set_enable_webgl.
func (recv *Settings) SetEnableWebgl(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_webgl((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetEnableXssAuditor is a wrapper around the C function webkit_settings_set_enable_xss_auditor.
func (recv *Settings) SetEnableXssAuditor(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_enable_xss_auditor((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetFantasyFontFamily is a wrapper around the C function webkit_settings_set_fantasy_font_family.
func (recv *Settings) SetFantasyFontFamily(fantasyFontFamily string) {
	c_fantasy_font_family := C.CString(fantasyFontFamily)
	defer C.free(unsafe.Pointer(c_fantasy_font_family))

	C.webkit_settings_set_fantasy_font_family((*C.WebKitSettings)(recv.native), c_fantasy_font_family)

	return
}

// SetJavascriptCanAccessClipboard is a wrapper around the C function webkit_settings_set_javascript_can_access_clipboard.
func (recv *Settings) SetJavascriptCanAccessClipboard(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_javascript_can_access_clipboard((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetJavascriptCanOpenWindowsAutomatically is a wrapper around the C function webkit_settings_set_javascript_can_open_windows_automatically.
func (recv *Settings) SetJavascriptCanOpenWindowsAutomatically(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_javascript_can_open_windows_automatically((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetLoadIconsIgnoringImageLoadSetting is a wrapper around the C function webkit_settings_set_load_icons_ignoring_image_load_setting.
func (recv *Settings) SetLoadIconsIgnoringImageLoadSetting(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_load_icons_ignoring_image_load_setting((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetMediaPlaybackAllowsInline is a wrapper around the C function webkit_settings_set_media_playback_allows_inline.
func (recv *Settings) SetMediaPlaybackAllowsInline(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_media_playback_allows_inline((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetMediaPlaybackRequiresUserGesture is a wrapper around the C function webkit_settings_set_media_playback_requires_user_gesture.
func (recv *Settings) SetMediaPlaybackRequiresUserGesture(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_settings_set_media_playback_requires_user_gesture((*C.WebKitSettings)(recv.native), c_enabled)

	return
}

// SetMinimumFontSize is a wrapper around the C function webkit_settings_set_minimum_font_size.
func (recv *Settings) SetMinimumFontSize(fontSize uint32) {
	c_font_size := (C.guint32)(fontSize)

	C.webkit_settings_set_minimum_font_size((*C.WebKitSettings)(recv.native), c_font_size)

	return
}

// SetMonospaceFontFamily is a wrapper around the C function webkit_settings_set_monospace_font_family.
func (recv *Settings) SetMonospaceFontFamily(monospaceFontFamily string) {
	c_monospace_font_family := C.CString(monospaceFontFamily)
	defer C.free(unsafe.Pointer(c_monospace_font_family))

	C.webkit_settings_set_monospace_font_family((*C.WebKitSettings)(recv.native), c_monospace_font_family)

	return
}

// SetPictographFontFamily is a wrapper around the C function webkit_settings_set_pictograph_font_family.
func (recv *Settings) SetPictographFontFamily(pictographFontFamily string) {
	c_pictograph_font_family := C.CString(pictographFontFamily)
	defer C.free(unsafe.Pointer(c_pictograph_font_family))

	C.webkit_settings_set_pictograph_font_family((*C.WebKitSettings)(recv.native), c_pictograph_font_family)

	return
}

// SetPrintBackgrounds is a wrapper around the C function webkit_settings_set_print_backgrounds.
func (recv *Settings) SetPrintBackgrounds(printBackgrounds bool) {
	c_print_backgrounds :=
		boolToGboolean(printBackgrounds)

	C.webkit_settings_set_print_backgrounds((*C.WebKitSettings)(recv.native), c_print_backgrounds)

	return
}

// SetSansSerifFontFamily is a wrapper around the C function webkit_settings_set_sans_serif_font_family.
func (recv *Settings) SetSansSerifFontFamily(sansSerifFontFamily string) {
	c_sans_serif_font_family := C.CString(sansSerifFontFamily)
	defer C.free(unsafe.Pointer(c_sans_serif_font_family))

	C.webkit_settings_set_sans_serif_font_family((*C.WebKitSettings)(recv.native), c_sans_serif_font_family)

	return
}

// SetSerifFontFamily is a wrapper around the C function webkit_settings_set_serif_font_family.
func (recv *Settings) SetSerifFontFamily(serifFontFamily string) {
	c_serif_font_family := C.CString(serifFontFamily)
	defer C.free(unsafe.Pointer(c_serif_font_family))

	C.webkit_settings_set_serif_font_family((*C.WebKitSettings)(recv.native), c_serif_font_family)

	return
}

// SetUserAgent is a wrapper around the C function webkit_settings_set_user_agent.
func (recv *Settings) SetUserAgent(userAgent string) {
	c_user_agent := C.CString(userAgent)
	defer C.free(unsafe.Pointer(c_user_agent))

	C.webkit_settings_set_user_agent((*C.WebKitSettings)(recv.native), c_user_agent)

	return
}

// SetUserAgentWithApplicationDetails is a wrapper around the C function webkit_settings_set_user_agent_with_application_details.
func (recv *Settings) SetUserAgentWithApplicationDetails(applicationName string, applicationVersion string) {
	c_application_name := C.CString(applicationName)
	defer C.free(unsafe.Pointer(c_application_name))

	c_application_version := C.CString(applicationVersion)
	defer C.free(unsafe.Pointer(c_application_version))

	C.webkit_settings_set_user_agent_with_application_details((*C.WebKitSettings)(recv.native), c_application_name, c_application_version)

	return
}

// SetZoomTextOnly is a wrapper around the C function webkit_settings_set_zoom_text_only.
func (recv *Settings) SetZoomTextOnly(zoomTextOnly bool) {
	c_zoom_text_only :=
		boolToGboolean(zoomTextOnly)

	C.webkit_settings_set_zoom_text_only((*C.WebKitSettings)(recv.native), c_zoom_text_only)

	return
}

// URIRequest is a wrapper around the C record WebKitURIRequest.
type URIRequest struct {
	native *C.WebKitURIRequest
	// parent : record
	// Private : priv
}

func URIRequestNewFromC(u unsafe.Pointer) *URIRequest {
	c := (*C.WebKitURIRequest)(u)
	if c == nil {
		return nil
	}

	g := &URIRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *URIRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *URIRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URIRequest with another URIRequest, and returns true if they represent the same GObject.
func (recv *URIRequest) Equals(other *URIRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *URIRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to URIRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a URIRequest.
func CastToURIRequest(object *gobject.Object) *URIRequest {
	return URIRequestNewFromC(object.ToC())
}

// URIRequestNew is a wrapper around the C function webkit_uri_request_new.
func URIRequestNew(uri string) *URIRequest {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.webkit_uri_request_new(c_uri)
	retGo := URIRequestNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetHttpHeaders is a wrapper around the C function webkit_uri_request_get_http_headers.
func (recv *URIRequest) GetHttpHeaders() *soup.MessageHeaders {
	retC := C.webkit_uri_request_get_http_headers((*C.WebKitURIRequest)(recv.native))
	retGo := soup.MessageHeadersNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function webkit_uri_request_get_uri.
func (recv *URIRequest) GetUri() string {
	retC := C.webkit_uri_request_get_uri((*C.WebKitURIRequest)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetUri is a wrapper around the C function webkit_uri_request_set_uri.
func (recv *URIRequest) SetUri(uri string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.webkit_uri_request_set_uri((*C.WebKitURIRequest)(recv.native), c_uri)

	return
}

// URIResponse is a wrapper around the C record WebKitURIResponse.
type URIResponse struct {
	native *C.WebKitURIResponse
	// parent : record
	// Private : priv
}

func URIResponseNewFromC(u unsafe.Pointer) *URIResponse {
	c := (*C.WebKitURIResponse)(u)
	if c == nil {
		return nil
	}

	g := &URIResponse{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *URIResponse) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *URIResponse) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URIResponse with another URIResponse, and returns true if they represent the same GObject.
func (recv *URIResponse) Equals(other *URIResponse) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *URIResponse) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to URIResponse.
// Exercise care, as this is a potentially dangerous function if the Object is not a URIResponse.
func CastToURIResponse(object *gobject.Object) *URIResponse {
	return URIResponseNewFromC(object.ToC())
}

// GetContentLength is a wrapper around the C function webkit_uri_response_get_content_length.
func (recv *URIResponse) GetContentLength() uint64 {
	retC := C.webkit_uri_response_get_content_length((*C.WebKitURIResponse)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetMimeType is a wrapper around the C function webkit_uri_response_get_mime_type.
func (recv *URIResponse) GetMimeType() string {
	retC := C.webkit_uri_response_get_mime_type((*C.WebKitURIResponse)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStatusCode is a wrapper around the C function webkit_uri_response_get_status_code.
func (recv *URIResponse) GetStatusCode() uint32 {
	retC := C.webkit_uri_response_get_status_code((*C.WebKitURIResponse)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetSuggestedFilename is a wrapper around the C function webkit_uri_response_get_suggested_filename.
func (recv *URIResponse) GetSuggestedFilename() string {
	retC := C.webkit_uri_response_get_suggested_filename((*C.WebKitURIResponse)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUri is a wrapper around the C function webkit_uri_response_get_uri.
func (recv *URIResponse) GetUri() string {
	retC := C.webkit_uri_response_get_uri((*C.WebKitURIResponse)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// URISchemeRequest is a wrapper around the C record WebKitURISchemeRequest.
type URISchemeRequest struct {
	native *C.WebKitURISchemeRequest
	// parent : record
	// priv : record
}

func URISchemeRequestNewFromC(u unsafe.Pointer) *URISchemeRequest {
	c := (*C.WebKitURISchemeRequest)(u)
	if c == nil {
		return nil
	}

	g := &URISchemeRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *URISchemeRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *URISchemeRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URISchemeRequest with another URISchemeRequest, and returns true if they represent the same GObject.
func (recv *URISchemeRequest) Equals(other *URISchemeRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *URISchemeRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to URISchemeRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a URISchemeRequest.
func CastToURISchemeRequest(object *gobject.Object) *URISchemeRequest {
	return URISchemeRequestNewFromC(object.ToC())
}

// Finish is a wrapper around the C function webkit_uri_scheme_request_finish.
func (recv *URISchemeRequest) Finish(stream *gio.InputStream, streamLength int64, mimeType string) {
	c_stream := (*C.GInputStream)(C.NULL)
	if stream != nil {
		c_stream = (*C.GInputStream)(stream.ToC())
	}

	c_stream_length := (C.gint64)(streamLength)

	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.webkit_uri_scheme_request_finish((*C.WebKitURISchemeRequest)(recv.native), c_stream, c_stream_length, c_mime_type)

	return
}

// GetPath is a wrapper around the C function webkit_uri_scheme_request_get_path.
func (recv *URISchemeRequest) GetPath() string {
	retC := C.webkit_uri_scheme_request_get_path((*C.WebKitURISchemeRequest)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetScheme is a wrapper around the C function webkit_uri_scheme_request_get_scheme.
func (recv *URISchemeRequest) GetScheme() string {
	retC := C.webkit_uri_scheme_request_get_scheme((*C.WebKitURISchemeRequest)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUri is a wrapper around the C function webkit_uri_scheme_request_get_uri.
func (recv *URISchemeRequest) GetUri() string {
	retC := C.webkit_uri_scheme_request_get_uri((*C.WebKitURISchemeRequest)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWebView is a wrapper around the C function webkit_uri_scheme_request_get_web_view.
func (recv *URISchemeRequest) GetWebView() *WebView {
	retC := C.webkit_uri_scheme_request_get_web_view((*C.WebKitURISchemeRequest)(recv.native))
	retGo := WebViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// UserContentFilterStore is a wrapper around the C record WebKitUserContentFilterStore.
type UserContentFilterStore struct {
	native *C.WebKitUserContentFilterStore
	// parent : record
	// Private : priv
}

func UserContentFilterStoreNewFromC(u unsafe.Pointer) *UserContentFilterStore {
	c := (*C.WebKitUserContentFilterStore)(u)
	if c == nil {
		return nil
	}

	g := &UserContentFilterStore{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UserContentFilterStore) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UserContentFilterStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserContentFilterStore with another UserContentFilterStore, and returns true if they represent the same GObject.
func (recv *UserContentFilterStore) Equals(other *UserContentFilterStore) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *UserContentFilterStore) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to UserContentFilterStore.
// Exercise care, as this is a potentially dangerous function if the Object is not a UserContentFilterStore.
func CastToUserContentFilterStore(object *gobject.Object) *UserContentFilterStore {
	return UserContentFilterStoreNewFromC(object.ToC())
}

// UserContentManager is a wrapper around the C record WebKitUserContentManager.
type UserContentManager struct {
	native *C.WebKitUserContentManager
	// parent : record
	// Private : priv
}

func UserContentManagerNewFromC(u unsafe.Pointer) *UserContentManager {
	c := (*C.WebKitUserContentManager)(u)
	if c == nil {
		return nil
	}

	g := &UserContentManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UserContentManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UserContentManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserContentManager with another UserContentManager, and returns true if they represent the same GObject.
func (recv *UserContentManager) Equals(other *UserContentManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *UserContentManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to UserContentManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a UserContentManager.
func CastToUserContentManager(object *gobject.Object) *UserContentManager {
	return UserContentManagerNewFromC(object.ToC())
}

// RemoveFilter is a wrapper around the C function webkit_user_content_manager_remove_filter.
func (recv *UserContentManager) RemoveFilter(filter *UserContentFilter) {
	c_filter := (*C.WebKitUserContentFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.WebKitUserContentFilter)(filter.ToC())
	}

	C.webkit_user_content_manager_remove_filter((*C.WebKitUserContentManager)(recv.native), c_filter)

	return
}

// UserMediaPermissionRequest is a wrapper around the C record WebKitUserMediaPermissionRequest.
type UserMediaPermissionRequest struct {
	native *C.WebKitUserMediaPermissionRequest
	// parent : record
	// Private : priv
}

func UserMediaPermissionRequestNewFromC(u unsafe.Pointer) *UserMediaPermissionRequest {
	c := (*C.WebKitUserMediaPermissionRequest)(u)
	if c == nil {
		return nil
	}

	g := &UserMediaPermissionRequest{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *UserMediaPermissionRequest) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *UserMediaPermissionRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserMediaPermissionRequest with another UserMediaPermissionRequest, and returns true if they represent the same GObject.
func (recv *UserMediaPermissionRequest) Equals(other *UserMediaPermissionRequest) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *UserMediaPermissionRequest) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to UserMediaPermissionRequest.
// Exercise care, as this is a potentially dangerous function if the Object is not a UserMediaPermissionRequest.
func CastToUserMediaPermissionRequest(object *gobject.Object) *UserMediaPermissionRequest {
	return UserMediaPermissionRequestNewFromC(object.ToC())
}

// PermissionRequest returns the PermissionRequest interface implemented by UserMediaPermissionRequest
func (recv *UserMediaPermissionRequest) PermissionRequest() *PermissionRequest {
	return PermissionRequestNewFromC(recv.ToC())
}

// WebContext is a wrapper around the C record WebKitWebContext.
type WebContext struct {
	native *C.WebKitWebContext
	// parent : record
	// Private : priv
}

func WebContextNewFromC(u unsafe.Pointer) *WebContext {
	c := (*C.WebKitWebContext)(u)
	if c == nil {
		return nil
	}

	g := &WebContext{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WebContext) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WebContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebContext with another WebContext, and returns true if they represent the same GObject.
func (recv *WebContext) Equals(other *WebContext) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *WebContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to WebContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a WebContext.
func CastToWebContext(object *gobject.Object) *WebContext {
	return WebContextNewFromC(object.ToC())
}

type signalWebContextDownloadStartedDetail struct {
	callback  WebContextSignalDownloadStartedCallback
	handlerID C.gulong
}

var signalWebContextDownloadStartedId int
var signalWebContextDownloadStartedMap = make(map[int]signalWebContextDownloadStartedDetail)
var signalWebContextDownloadStartedLock sync.RWMutex

// WebContextSignalDownloadStartedCallback is a callback function for a 'download-started' signal emitted from a WebContext.
type WebContextSignalDownloadStartedCallback func(targetObject *WebContext, download *Download)

/*
ConnectDownloadStarted connects the callback to the 'download-started' signal for the WebContext.

The returned value represents the connection, and may be passed to DisconnectDownloadStarted to remove it.
*/
func (recv *WebContext) ConnectDownloadStarted(callback WebContextSignalDownloadStartedCallback) int {
	signalWebContextDownloadStartedLock.Lock()
	defer signalWebContextDownloadStartedLock.Unlock()

	signalWebContextDownloadStartedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebContext_signal_connect_download_started(instance, C.gpointer(uintptr(signalWebContextDownloadStartedId)))

	detail := signalWebContextDownloadStartedDetail{callback, handlerID}
	signalWebContextDownloadStartedMap[signalWebContextDownloadStartedId] = detail

	return signalWebContextDownloadStartedId
}

/*
DisconnectDownloadStarted disconnects a callback from the 'download-started' signal for the WebContext.

The connectionID should be a value returned from a call to ConnectDownloadStarted.
*/
func (recv *WebContext) DisconnectDownloadStarted(connectionID int) {
	signalWebContextDownloadStartedLock.Lock()
	defer signalWebContextDownloadStartedLock.Unlock()

	detail, exists := signalWebContextDownloadStartedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebContextDownloadStartedMap, connectionID)
}

//export webcontext_downloadStartedHandler
func webcontext_downloadStartedHandler(c_targetObject *C.GObject, c_download *C.WebKitDownload, data C.gpointer) {
	signalWebContextDownloadStartedLock.RLock()
	defer signalWebContextDownloadStartedLock.RUnlock()

	download := DownloadNewFromC(unsafe.Pointer(c_download))

	targetObject := WebContextNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebContextDownloadStartedMap[index].callback
	callback(targetObject, download)
}

// WebContextGetDefault is a wrapper around the C function webkit_web_context_get_default.
func WebContextGetDefault() *WebContext {
	retC := C.webkit_web_context_get_default()
	retGo := WebContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ClearCache is a wrapper around the C function webkit_web_context_clear_cache.
func (recv *WebContext) ClearCache() {
	C.webkit_web_context_clear_cache((*C.WebKitWebContext)(recv.native))

	return
}

// DownloadUri is a wrapper around the C function webkit_web_context_download_uri.
func (recv *WebContext) DownloadUri(uri string) *Download {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.webkit_web_context_download_uri((*C.WebKitWebContext)(recv.native), c_uri)
	retGo := DownloadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCacheModel is a wrapper around the C function webkit_web_context_get_cache_model.
func (recv *WebContext) GetCacheModel() CacheModel {
	retC := C.webkit_web_context_get_cache_model((*C.WebKitWebContext)(recv.native))
	retGo := (CacheModel)(retC)

	return retGo
}

// GetCookieManager is a wrapper around the C function webkit_web_context_get_cookie_manager.
func (recv *WebContext) GetCookieManager() *CookieManager {
	retC := C.webkit_web_context_get_cookie_manager((*C.WebKitWebContext)(recv.native))
	retGo := CookieManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFaviconDatabase is a wrapper around the C function webkit_web_context_get_favicon_database.
func (recv *WebContext) GetFaviconDatabase() *FaviconDatabase {
	retC := C.webkit_web_context_get_favicon_database((*C.WebKitWebContext)(recv.native))
	retGo := FaviconDatabaseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFaviconDatabaseDirectory is a wrapper around the C function webkit_web_context_get_favicon_database_directory.
func (recv *WebContext) GetFaviconDatabaseDirectory() string {
	retC := C.webkit_web_context_get_favicon_database_directory((*C.WebKitWebContext)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : webkit_web_context_get_plugins : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// GetPluginsFinish is a wrapper around the C function webkit_web_context_get_plugins_finish.
func (recv *WebContext) GetPluginsFinish(result *gio.AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_web_context_get_plugins_finish((*C.WebKitWebContext)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetSecurityManager is a wrapper around the C function webkit_web_context_get_security_manager.
func (recv *WebContext) GetSecurityManager() *SecurityManager {
	retC := C.webkit_web_context_get_security_manager((*C.WebKitWebContext)(recv.native))
	retGo := SecurityManagerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSpellCheckingEnabled is a wrapper around the C function webkit_web_context_get_spell_checking_enabled.
func (recv *WebContext) GetSpellCheckingEnabled() bool {
	retC := C.webkit_web_context_get_spell_checking_enabled((*C.WebKitWebContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSpellCheckingLanguages is a wrapper around the C function webkit_web_context_get_spell_checking_languages.
func (recv *WebContext) GetSpellCheckingLanguages() []string {
	retC := C.webkit_web_context_get_spell_checking_languages((*C.WebKitWebContext)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetTlsErrorsPolicy is a wrapper around the C function webkit_web_context_get_tls_errors_policy.
func (recv *WebContext) GetTlsErrorsPolicy() TLSErrorsPolicy {
	retC := C.webkit_web_context_get_tls_errors_policy((*C.WebKitWebContext)(recv.native))
	retGo := (TLSErrorsPolicy)(retC)

	return retGo
}

// PrefetchDns is a wrapper around the C function webkit_web_context_prefetch_dns.
func (recv *WebContext) PrefetchDns(hostname string) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	C.webkit_web_context_prefetch_dns((*C.WebKitWebContext)(recv.native), c_hostname)

	return
}

// Unsupported : webkit_web_context_register_uri_scheme : unsupported parameter callback : no type generator for URISchemeRequestCallback (WebKitURISchemeRequestCallback) for param callback

// SetAdditionalPluginsDirectory is a wrapper around the C function webkit_web_context_set_additional_plugins_directory.
func (recv *WebContext) SetAdditionalPluginsDirectory(directory string) {
	c_directory := C.CString(directory)
	defer C.free(unsafe.Pointer(c_directory))

	C.webkit_web_context_set_additional_plugins_directory((*C.WebKitWebContext)(recv.native), c_directory)

	return
}

// SetCacheModel is a wrapper around the C function webkit_web_context_set_cache_model.
func (recv *WebContext) SetCacheModel(cacheModel CacheModel) {
	c_cache_model := (C.WebKitCacheModel)(cacheModel)

	C.webkit_web_context_set_cache_model((*C.WebKitWebContext)(recv.native), c_cache_model)

	return
}

// SetDiskCacheDirectory is a wrapper around the C function webkit_web_context_set_disk_cache_directory.
func (recv *WebContext) SetDiskCacheDirectory(directory string) {
	c_directory := C.CString(directory)
	defer C.free(unsafe.Pointer(c_directory))

	C.webkit_web_context_set_disk_cache_directory((*C.WebKitWebContext)(recv.native), c_directory)

	return
}

// SetFaviconDatabaseDirectory is a wrapper around the C function webkit_web_context_set_favicon_database_directory.
func (recv *WebContext) SetFaviconDatabaseDirectory(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.webkit_web_context_set_favicon_database_directory((*C.WebKitWebContext)(recv.native), c_path)

	return
}

// SetPreferredLanguages is a wrapper around the C function webkit_web_context_set_preferred_languages.
func (recv *WebContext) SetPreferredLanguages(languages []string) {
	c_languages_array := make([]*C.gchar, len(languages)+1, len(languages)+1)
	for i, item := range languages {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_languages_array[i] = c
	}
	c_languages_array[len(languages)] = nil
	c_languages_arrayPtr := &c_languages_array[0]
	c_languages := (**C.gchar)(unsafe.Pointer(c_languages_arrayPtr))

	C.webkit_web_context_set_preferred_languages((*C.WebKitWebContext)(recv.native), c_languages)

	return
}

// SetSpellCheckingEnabled is a wrapper around the C function webkit_web_context_set_spell_checking_enabled.
func (recv *WebContext) SetSpellCheckingEnabled(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.webkit_web_context_set_spell_checking_enabled((*C.WebKitWebContext)(recv.native), c_enabled)

	return
}

// SetSpellCheckingLanguages is a wrapper around the C function webkit_web_context_set_spell_checking_languages.
func (recv *WebContext) SetSpellCheckingLanguages(languages []string) {
	c_languages_array := make([]*C.gchar, len(languages)+1, len(languages)+1)
	for i, item := range languages {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_languages_array[i] = c
	}
	c_languages_array[len(languages)] = nil
	c_languages_arrayPtr := &c_languages_array[0]
	c_languages := (**C.gchar)(unsafe.Pointer(c_languages_arrayPtr))

	C.webkit_web_context_set_spell_checking_languages((*C.WebKitWebContext)(recv.native), c_languages)

	return
}

// SetTlsErrorsPolicy is a wrapper around the C function webkit_web_context_set_tls_errors_policy.
func (recv *WebContext) SetTlsErrorsPolicy(policy TLSErrorsPolicy) {
	c_policy := (C.WebKitTLSErrorsPolicy)(policy)

	C.webkit_web_context_set_tls_errors_policy((*C.WebKitWebContext)(recv.native), c_policy)

	return
}

// SetWebExtensionsDirectory is a wrapper around the C function webkit_web_context_set_web_extensions_directory.
func (recv *WebContext) SetWebExtensionsDirectory(directory string) {
	c_directory := C.CString(directory)
	defer C.free(unsafe.Pointer(c_directory))

	C.webkit_web_context_set_web_extensions_directory((*C.WebKitWebContext)(recv.native), c_directory)

	return
}

// WebInspector is a wrapper around the C record WebKitWebInspector.
type WebInspector struct {
	native *C.WebKitWebInspector
	// parent : record
	// priv : record
}

func WebInspectorNewFromC(u unsafe.Pointer) *WebInspector {
	c := (*C.WebKitWebInspector)(u)
	if c == nil {
		return nil
	}

	g := &WebInspector{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WebInspector) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WebInspector) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebInspector with another WebInspector, and returns true if they represent the same GObject.
func (recv *WebInspector) Equals(other *WebInspector) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *WebInspector) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to WebInspector.
// Exercise care, as this is a potentially dangerous function if the Object is not a WebInspector.
func CastToWebInspector(object *gobject.Object) *WebInspector {
	return WebInspectorNewFromC(object.ToC())
}

type signalWebInspectorAttachDetail struct {
	callback  WebInspectorSignalAttachCallback
	handlerID C.gulong
}

var signalWebInspectorAttachId int
var signalWebInspectorAttachMap = make(map[int]signalWebInspectorAttachDetail)
var signalWebInspectorAttachLock sync.RWMutex

// WebInspectorSignalAttachCallback is a callback function for a 'attach' signal emitted from a WebInspector.
type WebInspectorSignalAttachCallback func(targetObject *WebInspector) bool

/*
ConnectAttach connects the callback to the 'attach' signal for the WebInspector.

The returned value represents the connection, and may be passed to DisconnectAttach to remove it.
*/
func (recv *WebInspector) ConnectAttach(callback WebInspectorSignalAttachCallback) int {
	signalWebInspectorAttachLock.Lock()
	defer signalWebInspectorAttachLock.Unlock()

	signalWebInspectorAttachId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebInspector_signal_connect_attach(instance, C.gpointer(uintptr(signalWebInspectorAttachId)))

	detail := signalWebInspectorAttachDetail{callback, handlerID}
	signalWebInspectorAttachMap[signalWebInspectorAttachId] = detail

	return signalWebInspectorAttachId
}

/*
DisconnectAttach disconnects a callback from the 'attach' signal for the WebInspector.

The connectionID should be a value returned from a call to ConnectAttach.
*/
func (recv *WebInspector) DisconnectAttach(connectionID int) {
	signalWebInspectorAttachLock.Lock()
	defer signalWebInspectorAttachLock.Unlock()

	detail, exists := signalWebInspectorAttachMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebInspectorAttachMap, connectionID)
}

//export webinspector_attachHandler
func webinspector_attachHandler(c_targetObject *C.GObject, data C.gpointer) C.gboolean {
	signalWebInspectorAttachLock.RLock()
	defer signalWebInspectorAttachLock.RUnlock()

	targetObject := WebInspectorNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebInspectorAttachMap[index].callback
	retGo := callback(targetObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebInspectorBringToFrontDetail struct {
	callback  WebInspectorSignalBringToFrontCallback
	handlerID C.gulong
}

var signalWebInspectorBringToFrontId int
var signalWebInspectorBringToFrontMap = make(map[int]signalWebInspectorBringToFrontDetail)
var signalWebInspectorBringToFrontLock sync.RWMutex

// WebInspectorSignalBringToFrontCallback is a callback function for a 'bring-to-front' signal emitted from a WebInspector.
type WebInspectorSignalBringToFrontCallback func(targetObject *WebInspector) bool

/*
ConnectBringToFront connects the callback to the 'bring-to-front' signal for the WebInspector.

The returned value represents the connection, and may be passed to DisconnectBringToFront to remove it.
*/
func (recv *WebInspector) ConnectBringToFront(callback WebInspectorSignalBringToFrontCallback) int {
	signalWebInspectorBringToFrontLock.Lock()
	defer signalWebInspectorBringToFrontLock.Unlock()

	signalWebInspectorBringToFrontId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebInspector_signal_connect_bring_to_front(instance, C.gpointer(uintptr(signalWebInspectorBringToFrontId)))

	detail := signalWebInspectorBringToFrontDetail{callback, handlerID}
	signalWebInspectorBringToFrontMap[signalWebInspectorBringToFrontId] = detail

	return signalWebInspectorBringToFrontId
}

/*
DisconnectBringToFront disconnects a callback from the 'bring-to-front' signal for the WebInspector.

The connectionID should be a value returned from a call to ConnectBringToFront.
*/
func (recv *WebInspector) DisconnectBringToFront(connectionID int) {
	signalWebInspectorBringToFrontLock.Lock()
	defer signalWebInspectorBringToFrontLock.Unlock()

	detail, exists := signalWebInspectorBringToFrontMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebInspectorBringToFrontMap, connectionID)
}

//export webinspector_bringToFrontHandler
func webinspector_bringToFrontHandler(c_targetObject *C.GObject, data C.gpointer) C.gboolean {
	signalWebInspectorBringToFrontLock.RLock()
	defer signalWebInspectorBringToFrontLock.RUnlock()

	targetObject := WebInspectorNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebInspectorBringToFrontMap[index].callback
	retGo := callback(targetObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebInspectorClosedDetail struct {
	callback  WebInspectorSignalClosedCallback
	handlerID C.gulong
}

var signalWebInspectorClosedId int
var signalWebInspectorClosedMap = make(map[int]signalWebInspectorClosedDetail)
var signalWebInspectorClosedLock sync.RWMutex

// WebInspectorSignalClosedCallback is a callback function for a 'closed' signal emitted from a WebInspector.
type WebInspectorSignalClosedCallback func(targetObject *WebInspector)

/*
ConnectClosed connects the callback to the 'closed' signal for the WebInspector.

The returned value represents the connection, and may be passed to DisconnectClosed to remove it.
*/
func (recv *WebInspector) ConnectClosed(callback WebInspectorSignalClosedCallback) int {
	signalWebInspectorClosedLock.Lock()
	defer signalWebInspectorClosedLock.Unlock()

	signalWebInspectorClosedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebInspector_signal_connect_closed(instance, C.gpointer(uintptr(signalWebInspectorClosedId)))

	detail := signalWebInspectorClosedDetail{callback, handlerID}
	signalWebInspectorClosedMap[signalWebInspectorClosedId] = detail

	return signalWebInspectorClosedId
}

/*
DisconnectClosed disconnects a callback from the 'closed' signal for the WebInspector.

The connectionID should be a value returned from a call to ConnectClosed.
*/
func (recv *WebInspector) DisconnectClosed(connectionID int) {
	signalWebInspectorClosedLock.Lock()
	defer signalWebInspectorClosedLock.Unlock()

	detail, exists := signalWebInspectorClosedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebInspectorClosedMap, connectionID)
}

//export webinspector_closedHandler
func webinspector_closedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalWebInspectorClosedLock.RLock()
	defer signalWebInspectorClosedLock.RUnlock()

	targetObject := WebInspectorNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebInspectorClosedMap[index].callback
	callback(targetObject)
}

type signalWebInspectorDetachDetail struct {
	callback  WebInspectorSignalDetachCallback
	handlerID C.gulong
}

var signalWebInspectorDetachId int
var signalWebInspectorDetachMap = make(map[int]signalWebInspectorDetachDetail)
var signalWebInspectorDetachLock sync.RWMutex

// WebInspectorSignalDetachCallback is a callback function for a 'detach' signal emitted from a WebInspector.
type WebInspectorSignalDetachCallback func(targetObject *WebInspector) bool

/*
ConnectDetach connects the callback to the 'detach' signal for the WebInspector.

The returned value represents the connection, and may be passed to DisconnectDetach to remove it.
*/
func (recv *WebInspector) ConnectDetach(callback WebInspectorSignalDetachCallback) int {
	signalWebInspectorDetachLock.Lock()
	defer signalWebInspectorDetachLock.Unlock()

	signalWebInspectorDetachId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebInspector_signal_connect_detach(instance, C.gpointer(uintptr(signalWebInspectorDetachId)))

	detail := signalWebInspectorDetachDetail{callback, handlerID}
	signalWebInspectorDetachMap[signalWebInspectorDetachId] = detail

	return signalWebInspectorDetachId
}

/*
DisconnectDetach disconnects a callback from the 'detach' signal for the WebInspector.

The connectionID should be a value returned from a call to ConnectDetach.
*/
func (recv *WebInspector) DisconnectDetach(connectionID int) {
	signalWebInspectorDetachLock.Lock()
	defer signalWebInspectorDetachLock.Unlock()

	detail, exists := signalWebInspectorDetachMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebInspectorDetachMap, connectionID)
}

//export webinspector_detachHandler
func webinspector_detachHandler(c_targetObject *C.GObject, data C.gpointer) C.gboolean {
	signalWebInspectorDetachLock.RLock()
	defer signalWebInspectorDetachLock.RUnlock()

	targetObject := WebInspectorNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebInspectorDetachMap[index].callback
	retGo := callback(targetObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebInspectorOpenWindowDetail struct {
	callback  WebInspectorSignalOpenWindowCallback
	handlerID C.gulong
}

var signalWebInspectorOpenWindowId int
var signalWebInspectorOpenWindowMap = make(map[int]signalWebInspectorOpenWindowDetail)
var signalWebInspectorOpenWindowLock sync.RWMutex

// WebInspectorSignalOpenWindowCallback is a callback function for a 'open-window' signal emitted from a WebInspector.
type WebInspectorSignalOpenWindowCallback func(targetObject *WebInspector) bool

/*
ConnectOpenWindow connects the callback to the 'open-window' signal for the WebInspector.

The returned value represents the connection, and may be passed to DisconnectOpenWindow to remove it.
*/
func (recv *WebInspector) ConnectOpenWindow(callback WebInspectorSignalOpenWindowCallback) int {
	signalWebInspectorOpenWindowLock.Lock()
	defer signalWebInspectorOpenWindowLock.Unlock()

	signalWebInspectorOpenWindowId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebInspector_signal_connect_open_window(instance, C.gpointer(uintptr(signalWebInspectorOpenWindowId)))

	detail := signalWebInspectorOpenWindowDetail{callback, handlerID}
	signalWebInspectorOpenWindowMap[signalWebInspectorOpenWindowId] = detail

	return signalWebInspectorOpenWindowId
}

/*
DisconnectOpenWindow disconnects a callback from the 'open-window' signal for the WebInspector.

The connectionID should be a value returned from a call to ConnectOpenWindow.
*/
func (recv *WebInspector) DisconnectOpenWindow(connectionID int) {
	signalWebInspectorOpenWindowLock.Lock()
	defer signalWebInspectorOpenWindowLock.Unlock()

	detail, exists := signalWebInspectorOpenWindowMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebInspectorOpenWindowMap, connectionID)
}

//export webinspector_openWindowHandler
func webinspector_openWindowHandler(c_targetObject *C.GObject, data C.gpointer) C.gboolean {
	signalWebInspectorOpenWindowLock.RLock()
	defer signalWebInspectorOpenWindowLock.RUnlock()

	targetObject := WebInspectorNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebInspectorOpenWindowMap[index].callback
	retGo := callback(targetObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Attach is a wrapper around the C function webkit_web_inspector_attach.
func (recv *WebInspector) Attach() {
	C.webkit_web_inspector_attach((*C.WebKitWebInspector)(recv.native))

	return
}

// Close is a wrapper around the C function webkit_web_inspector_close.
func (recv *WebInspector) Close() {
	C.webkit_web_inspector_close((*C.WebKitWebInspector)(recv.native))

	return
}

// Detach is a wrapper around the C function webkit_web_inspector_detach.
func (recv *WebInspector) Detach() {
	C.webkit_web_inspector_detach((*C.WebKitWebInspector)(recv.native))

	return
}

// GetAttachedHeight is a wrapper around the C function webkit_web_inspector_get_attached_height.
func (recv *WebInspector) GetAttachedHeight() uint32 {
	retC := C.webkit_web_inspector_get_attached_height((*C.WebKitWebInspector)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetInspectedUri is a wrapper around the C function webkit_web_inspector_get_inspected_uri.
func (recv *WebInspector) GetInspectedUri() string {
	retC := C.webkit_web_inspector_get_inspected_uri((*C.WebKitWebInspector)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWebView is a wrapper around the C function webkit_web_inspector_get_web_view.
func (recv *WebInspector) GetWebView() *WebViewBase {
	retC := C.webkit_web_inspector_get_web_view((*C.WebKitWebInspector)(recv.native))
	retGo := WebViewBaseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsAttached is a wrapper around the C function webkit_web_inspector_is_attached.
func (recv *WebInspector) IsAttached() bool {
	retC := C.webkit_web_inspector_is_attached((*C.WebKitWebInspector)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Show is a wrapper around the C function webkit_web_inspector_show.
func (recv *WebInspector) Show() {
	C.webkit_web_inspector_show((*C.WebKitWebInspector)(recv.native))

	return
}

// WebResource is a wrapper around the C record WebKitWebResource.
type WebResource struct {
	native *C.WebKitWebResource
	// parent : record
	// priv : record
}

func WebResourceNewFromC(u unsafe.Pointer) *WebResource {
	c := (*C.WebKitWebResource)(u)
	if c == nil {
		return nil
	}

	g := &WebResource{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WebResource) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WebResource) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebResource with another WebResource, and returns true if they represent the same GObject.
func (recv *WebResource) Equals(other *WebResource) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *WebResource) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to WebResource.
// Exercise care, as this is a potentially dangerous function if the Object is not a WebResource.
func CastToWebResource(object *gobject.Object) *WebResource {
	return WebResourceNewFromC(object.ToC())
}

type signalWebResourceFailedDetail struct {
	callback  WebResourceSignalFailedCallback
	handlerID C.gulong
}

var signalWebResourceFailedId int
var signalWebResourceFailedMap = make(map[int]signalWebResourceFailedDetail)
var signalWebResourceFailedLock sync.RWMutex

// WebResourceSignalFailedCallback is a callback function for a 'failed' signal emitted from a WebResource.
type WebResourceSignalFailedCallback func(targetObject *WebResource, error *glib.Error)

/*
ConnectFailed connects the callback to the 'failed' signal for the WebResource.

The returned value represents the connection, and may be passed to DisconnectFailed to remove it.
*/
func (recv *WebResource) ConnectFailed(callback WebResourceSignalFailedCallback) int {
	signalWebResourceFailedLock.Lock()
	defer signalWebResourceFailedLock.Unlock()

	signalWebResourceFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebResource_signal_connect_failed(instance, C.gpointer(uintptr(signalWebResourceFailedId)))

	detail := signalWebResourceFailedDetail{callback, handlerID}
	signalWebResourceFailedMap[signalWebResourceFailedId] = detail

	return signalWebResourceFailedId
}

/*
DisconnectFailed disconnects a callback from the 'failed' signal for the WebResource.

The connectionID should be a value returned from a call to ConnectFailed.
*/
func (recv *WebResource) DisconnectFailed(connectionID int) {
	signalWebResourceFailedLock.Lock()
	defer signalWebResourceFailedLock.Unlock()

	detail, exists := signalWebResourceFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebResourceFailedMap, connectionID)
}

//export webresource_failedHandler
func webresource_failedHandler(c_targetObject *C.GObject, c_error *C.GError, data C.gpointer) {
	signalWebResourceFailedLock.RLock()
	defer signalWebResourceFailedLock.RUnlock()

	error := glib.ErrorNewFromC(unsafe.Pointer(c_error))

	targetObject := WebResourceNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebResourceFailedMap[index].callback
	callback(targetObject, error)
}

type signalWebResourceFinishedDetail struct {
	callback  WebResourceSignalFinishedCallback
	handlerID C.gulong
}

var signalWebResourceFinishedId int
var signalWebResourceFinishedMap = make(map[int]signalWebResourceFinishedDetail)
var signalWebResourceFinishedLock sync.RWMutex

// WebResourceSignalFinishedCallback is a callback function for a 'finished' signal emitted from a WebResource.
type WebResourceSignalFinishedCallback func(targetObject *WebResource)

/*
ConnectFinished connects the callback to the 'finished' signal for the WebResource.

The returned value represents the connection, and may be passed to DisconnectFinished to remove it.
*/
func (recv *WebResource) ConnectFinished(callback WebResourceSignalFinishedCallback) int {
	signalWebResourceFinishedLock.Lock()
	defer signalWebResourceFinishedLock.Unlock()

	signalWebResourceFinishedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebResource_signal_connect_finished(instance, C.gpointer(uintptr(signalWebResourceFinishedId)))

	detail := signalWebResourceFinishedDetail{callback, handlerID}
	signalWebResourceFinishedMap[signalWebResourceFinishedId] = detail

	return signalWebResourceFinishedId
}

/*
DisconnectFinished disconnects a callback from the 'finished' signal for the WebResource.

The connectionID should be a value returned from a call to ConnectFinished.
*/
func (recv *WebResource) DisconnectFinished(connectionID int) {
	signalWebResourceFinishedLock.Lock()
	defer signalWebResourceFinishedLock.Unlock()

	detail, exists := signalWebResourceFinishedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebResourceFinishedMap, connectionID)
}

//export webresource_finishedHandler
func webresource_finishedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalWebResourceFinishedLock.RLock()
	defer signalWebResourceFinishedLock.RUnlock()

	targetObject := WebResourceNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebResourceFinishedMap[index].callback
	callback(targetObject)
}

type signalWebResourceReceivedDataDetail struct {
	callback  WebResourceSignalReceivedDataCallback
	handlerID C.gulong
}

var signalWebResourceReceivedDataId int
var signalWebResourceReceivedDataMap = make(map[int]signalWebResourceReceivedDataDetail)
var signalWebResourceReceivedDataLock sync.RWMutex

// WebResourceSignalReceivedDataCallback is a callback function for a 'received-data' signal emitted from a WebResource.
type WebResourceSignalReceivedDataCallback func(targetObject *WebResource, dataLength uint64)

/*
ConnectReceivedData connects the callback to the 'received-data' signal for the WebResource.

The returned value represents the connection, and may be passed to DisconnectReceivedData to remove it.
*/
func (recv *WebResource) ConnectReceivedData(callback WebResourceSignalReceivedDataCallback) int {
	signalWebResourceReceivedDataLock.Lock()
	defer signalWebResourceReceivedDataLock.Unlock()

	signalWebResourceReceivedDataId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebResource_signal_connect_received_data(instance, C.gpointer(uintptr(signalWebResourceReceivedDataId)))

	detail := signalWebResourceReceivedDataDetail{callback, handlerID}
	signalWebResourceReceivedDataMap[signalWebResourceReceivedDataId] = detail

	return signalWebResourceReceivedDataId
}

/*
DisconnectReceivedData disconnects a callback from the 'received-data' signal for the WebResource.

The connectionID should be a value returned from a call to ConnectReceivedData.
*/
func (recv *WebResource) DisconnectReceivedData(connectionID int) {
	signalWebResourceReceivedDataLock.Lock()
	defer signalWebResourceReceivedDataLock.Unlock()

	detail, exists := signalWebResourceReceivedDataMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebResourceReceivedDataMap, connectionID)
}

//export webresource_receivedDataHandler
func webresource_receivedDataHandler(c_targetObject *C.GObject, c_data_length C.guint64, data C.gpointer) {
	signalWebResourceReceivedDataLock.RLock()
	defer signalWebResourceReceivedDataLock.RUnlock()

	dataLength := uint64(c_data_length)

	targetObject := WebResourceNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebResourceReceivedDataMap[index].callback
	callback(targetObject, dataLength)
}

type signalWebResourceSentRequestDetail struct {
	callback  WebResourceSignalSentRequestCallback
	handlerID C.gulong
}

var signalWebResourceSentRequestId int
var signalWebResourceSentRequestMap = make(map[int]signalWebResourceSentRequestDetail)
var signalWebResourceSentRequestLock sync.RWMutex

// WebResourceSignalSentRequestCallback is a callback function for a 'sent-request' signal emitted from a WebResource.
type WebResourceSignalSentRequestCallback func(targetObject *WebResource, request *URIRequest, redirectedResponse *URIResponse)

/*
ConnectSentRequest connects the callback to the 'sent-request' signal for the WebResource.

The returned value represents the connection, and may be passed to DisconnectSentRequest to remove it.
*/
func (recv *WebResource) ConnectSentRequest(callback WebResourceSignalSentRequestCallback) int {
	signalWebResourceSentRequestLock.Lock()
	defer signalWebResourceSentRequestLock.Unlock()

	signalWebResourceSentRequestId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebResource_signal_connect_sent_request(instance, C.gpointer(uintptr(signalWebResourceSentRequestId)))

	detail := signalWebResourceSentRequestDetail{callback, handlerID}
	signalWebResourceSentRequestMap[signalWebResourceSentRequestId] = detail

	return signalWebResourceSentRequestId
}

/*
DisconnectSentRequest disconnects a callback from the 'sent-request' signal for the WebResource.

The connectionID should be a value returned from a call to ConnectSentRequest.
*/
func (recv *WebResource) DisconnectSentRequest(connectionID int) {
	signalWebResourceSentRequestLock.Lock()
	defer signalWebResourceSentRequestLock.Unlock()

	detail, exists := signalWebResourceSentRequestMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebResourceSentRequestMap, connectionID)
}

//export webresource_sentRequestHandler
func webresource_sentRequestHandler(c_targetObject *C.GObject, c_request *C.WebKitURIRequest, c_redirected_response *C.WebKitURIResponse, data C.gpointer) {
	signalWebResourceSentRequestLock.RLock()
	defer signalWebResourceSentRequestLock.RUnlock()

	request := URIRequestNewFromC(unsafe.Pointer(c_request))

	redirectedResponse := URIResponseNewFromC(unsafe.Pointer(c_redirected_response))

	targetObject := WebResourceNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebResourceSentRequestMap[index].callback
	callback(targetObject, request, redirectedResponse)
}

// Unsupported : webkit_web_resource_get_data : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : webkit_web_resource_get_data_finish : array return type :

// GetResponse is a wrapper around the C function webkit_web_resource_get_response.
func (recv *WebResource) GetResponse() *URIResponse {
	retC := C.webkit_web_resource_get_response((*C.WebKitWebResource)(recv.native))
	retGo := URIResponseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUri is a wrapper around the C function webkit_web_resource_get_uri.
func (recv *WebResource) GetUri() string {
	retC := C.webkit_web_resource_get_uri((*C.WebKitWebResource)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// WebView is a wrapper around the C record WebKitWebView.
type WebView struct {
	native *C.WebKitWebView
	// parent : record
	// Private : priv
}

func WebViewNewFromC(u unsafe.Pointer) *WebView {
	c := (*C.WebKitWebView)(u)
	if c == nil {
		return nil
	}

	g := &WebView{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WebView) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WebView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebView with another WebView, and returns true if they represent the same GObject.
func (recv *WebView) Equals(other *WebView) bool {
	return other.ToC() == recv.ToC()
}

// WebViewBase upcasts to *WebViewBase
func (recv *WebView) WebViewBase() *WebViewBase {
	return WebViewBaseNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *WebView) Container() *gtk.Container {
	return recv.WebViewBase().Container()
}

// Widget upcasts to *Widget
func (recv *WebView) Widget() *gtk.Widget {
	return recv.WebViewBase().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *WebView) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.WebViewBase().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *WebView) Object() *gobject.Object {
	return recv.WebViewBase().Object()
}

// CastToWidget down casts any arbitrary Object to WebView.
// Exercise care, as this is a potentially dangerous function if the Object is not a WebView.
func CastToWebView(object *gobject.Object) *WebView {
	return WebViewNewFromC(object.ToC())
}

type signalWebViewCloseDetail struct {
	callback  WebViewSignalCloseCallback
	handlerID C.gulong
}

var signalWebViewCloseId int
var signalWebViewCloseMap = make(map[int]signalWebViewCloseDetail)
var signalWebViewCloseLock sync.RWMutex

// WebViewSignalCloseCallback is a callback function for a 'close' signal emitted from a WebView.
type WebViewSignalCloseCallback func(targetObject *WebView)

/*
ConnectClose connects the callback to the 'close' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectClose to remove it.
*/
func (recv *WebView) ConnectClose(callback WebViewSignalCloseCallback) int {
	signalWebViewCloseLock.Lock()
	defer signalWebViewCloseLock.Unlock()

	signalWebViewCloseId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_close(instance, C.gpointer(uintptr(signalWebViewCloseId)))

	detail := signalWebViewCloseDetail{callback, handlerID}
	signalWebViewCloseMap[signalWebViewCloseId] = detail

	return signalWebViewCloseId
}

/*
DisconnectClose disconnects a callback from the 'close' signal for the WebView.

The connectionID should be a value returned from a call to ConnectClose.
*/
func (recv *WebView) DisconnectClose(connectionID int) {
	signalWebViewCloseLock.Lock()
	defer signalWebViewCloseLock.Unlock()

	detail, exists := signalWebViewCloseMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewCloseMap, connectionID)
}

//export webview_closeHandler
func webview_closeHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalWebViewCloseLock.RLock()
	defer signalWebViewCloseLock.RUnlock()

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewCloseMap[index].callback
	callback(targetObject)
}

// Blacklisted signal context-menu

type signalWebViewContextMenuDismissedDetail struct {
	callback  WebViewSignalContextMenuDismissedCallback
	handlerID C.gulong
}

var signalWebViewContextMenuDismissedId int
var signalWebViewContextMenuDismissedMap = make(map[int]signalWebViewContextMenuDismissedDetail)
var signalWebViewContextMenuDismissedLock sync.RWMutex

// WebViewSignalContextMenuDismissedCallback is a callback function for a 'context-menu-dismissed' signal emitted from a WebView.
type WebViewSignalContextMenuDismissedCallback func(targetObject *WebView)

/*
ConnectContextMenuDismissed connects the callback to the 'context-menu-dismissed' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectContextMenuDismissed to remove it.
*/
func (recv *WebView) ConnectContextMenuDismissed(callback WebViewSignalContextMenuDismissedCallback) int {
	signalWebViewContextMenuDismissedLock.Lock()
	defer signalWebViewContextMenuDismissedLock.Unlock()

	signalWebViewContextMenuDismissedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_context_menu_dismissed(instance, C.gpointer(uintptr(signalWebViewContextMenuDismissedId)))

	detail := signalWebViewContextMenuDismissedDetail{callback, handlerID}
	signalWebViewContextMenuDismissedMap[signalWebViewContextMenuDismissedId] = detail

	return signalWebViewContextMenuDismissedId
}

/*
DisconnectContextMenuDismissed disconnects a callback from the 'context-menu-dismissed' signal for the WebView.

The connectionID should be a value returned from a call to ConnectContextMenuDismissed.
*/
func (recv *WebView) DisconnectContextMenuDismissed(connectionID int) {
	signalWebViewContextMenuDismissedLock.Lock()
	defer signalWebViewContextMenuDismissedLock.Unlock()

	detail, exists := signalWebViewContextMenuDismissedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewContextMenuDismissedMap, connectionID)
}

//export webview_contextMenuDismissedHandler
func webview_contextMenuDismissedHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalWebViewContextMenuDismissedLock.RLock()
	defer signalWebViewContextMenuDismissedLock.RUnlock()

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewContextMenuDismissedMap[index].callback
	callback(targetObject)
}

type signalWebViewCreateDetail struct {
	callback  WebViewSignalCreateCallback
	handlerID C.gulong
}

var signalWebViewCreateId int
var signalWebViewCreateMap = make(map[int]signalWebViewCreateDetail)
var signalWebViewCreateLock sync.RWMutex

// WebViewSignalCreateCallback is a callback function for a 'create' signal emitted from a WebView.
type WebViewSignalCreateCallback func(targetObject *WebView, navigationAction *NavigationAction) gtk.Widget

/*
ConnectCreate connects the callback to the 'create' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectCreate to remove it.
*/
func (recv *WebView) ConnectCreate(callback WebViewSignalCreateCallback) int {
	signalWebViewCreateLock.Lock()
	defer signalWebViewCreateLock.Unlock()

	signalWebViewCreateId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_create(instance, C.gpointer(uintptr(signalWebViewCreateId)))

	detail := signalWebViewCreateDetail{callback, handlerID}
	signalWebViewCreateMap[signalWebViewCreateId] = detail

	return signalWebViewCreateId
}

/*
DisconnectCreate disconnects a callback from the 'create' signal for the WebView.

The connectionID should be a value returned from a call to ConnectCreate.
*/
func (recv *WebView) DisconnectCreate(connectionID int) {
	signalWebViewCreateLock.Lock()
	defer signalWebViewCreateLock.Unlock()

	detail, exists := signalWebViewCreateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewCreateMap, connectionID)
}

//export webview_createHandler
func webview_createHandler(c_targetObject *C.GObject, c_navigation_action *C.WebKitNavigationAction, data C.gpointer) *C.GtkWidget {
	signalWebViewCreateLock.RLock()
	defer signalWebViewCreateLock.RUnlock()

	navigationAction := NavigationActionNewFromC(unsafe.Pointer(c_navigation_action))

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewCreateMap[index].callback
	retGo := callback(targetObject, navigationAction)
	retC :=
		(*C.GtkWidget)(retGo.ToC())
	return retC
}

type signalWebViewDecidePolicyDetail struct {
	callback  WebViewSignalDecidePolicyCallback
	handlerID C.gulong
}

var signalWebViewDecidePolicyId int
var signalWebViewDecidePolicyMap = make(map[int]signalWebViewDecidePolicyDetail)
var signalWebViewDecidePolicyLock sync.RWMutex

// WebViewSignalDecidePolicyCallback is a callback function for a 'decide-policy' signal emitted from a WebView.
type WebViewSignalDecidePolicyCallback func(targetObject *WebView, decision *PolicyDecision, decisionType PolicyDecisionType) bool

/*
ConnectDecidePolicy connects the callback to the 'decide-policy' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectDecidePolicy to remove it.
*/
func (recv *WebView) ConnectDecidePolicy(callback WebViewSignalDecidePolicyCallback) int {
	signalWebViewDecidePolicyLock.Lock()
	defer signalWebViewDecidePolicyLock.Unlock()

	signalWebViewDecidePolicyId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_decide_policy(instance, C.gpointer(uintptr(signalWebViewDecidePolicyId)))

	detail := signalWebViewDecidePolicyDetail{callback, handlerID}
	signalWebViewDecidePolicyMap[signalWebViewDecidePolicyId] = detail

	return signalWebViewDecidePolicyId
}

/*
DisconnectDecidePolicy disconnects a callback from the 'decide-policy' signal for the WebView.

The connectionID should be a value returned from a call to ConnectDecidePolicy.
*/
func (recv *WebView) DisconnectDecidePolicy(connectionID int) {
	signalWebViewDecidePolicyLock.Lock()
	defer signalWebViewDecidePolicyLock.Unlock()

	detail, exists := signalWebViewDecidePolicyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewDecidePolicyMap, connectionID)
}

//export webview_decidePolicyHandler
func webview_decidePolicyHandler(c_targetObject *C.GObject, c_decision *C.WebKitPolicyDecision, c_decision_type C.WebKitPolicyDecisionType, data C.gpointer) C.gboolean {
	signalWebViewDecidePolicyLock.RLock()
	defer signalWebViewDecidePolicyLock.RUnlock()

	decision := PolicyDecisionNewFromC(unsafe.Pointer(c_decision))

	decisionType := PolicyDecisionType(c_decision_type)

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewDecidePolicyMap[index].callback
	retGo := callback(targetObject, decision, decisionType)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebViewEnterFullscreenDetail struct {
	callback  WebViewSignalEnterFullscreenCallback
	handlerID C.gulong
}

var signalWebViewEnterFullscreenId int
var signalWebViewEnterFullscreenMap = make(map[int]signalWebViewEnterFullscreenDetail)
var signalWebViewEnterFullscreenLock sync.RWMutex

// WebViewSignalEnterFullscreenCallback is a callback function for a 'enter-fullscreen' signal emitted from a WebView.
type WebViewSignalEnterFullscreenCallback func(targetObject *WebView) bool

/*
ConnectEnterFullscreen connects the callback to the 'enter-fullscreen' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectEnterFullscreen to remove it.
*/
func (recv *WebView) ConnectEnterFullscreen(callback WebViewSignalEnterFullscreenCallback) int {
	signalWebViewEnterFullscreenLock.Lock()
	defer signalWebViewEnterFullscreenLock.Unlock()

	signalWebViewEnterFullscreenId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_enter_fullscreen(instance, C.gpointer(uintptr(signalWebViewEnterFullscreenId)))

	detail := signalWebViewEnterFullscreenDetail{callback, handlerID}
	signalWebViewEnterFullscreenMap[signalWebViewEnterFullscreenId] = detail

	return signalWebViewEnterFullscreenId
}

/*
DisconnectEnterFullscreen disconnects a callback from the 'enter-fullscreen' signal for the WebView.

The connectionID should be a value returned from a call to ConnectEnterFullscreen.
*/
func (recv *WebView) DisconnectEnterFullscreen(connectionID int) {
	signalWebViewEnterFullscreenLock.Lock()
	defer signalWebViewEnterFullscreenLock.Unlock()

	detail, exists := signalWebViewEnterFullscreenMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewEnterFullscreenMap, connectionID)
}

//export webview_enterFullscreenHandler
func webview_enterFullscreenHandler(c_targetObject *C.GObject, data C.gpointer) C.gboolean {
	signalWebViewEnterFullscreenLock.RLock()
	defer signalWebViewEnterFullscreenLock.RUnlock()

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewEnterFullscreenMap[index].callback
	retGo := callback(targetObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebViewInsecureContentDetectedDetail struct {
	callback  WebViewSignalInsecureContentDetectedCallback
	handlerID C.gulong
}

var signalWebViewInsecureContentDetectedId int
var signalWebViewInsecureContentDetectedMap = make(map[int]signalWebViewInsecureContentDetectedDetail)
var signalWebViewInsecureContentDetectedLock sync.RWMutex

// WebViewSignalInsecureContentDetectedCallback is a callback function for a 'insecure-content-detected' signal emitted from a WebView.
type WebViewSignalInsecureContentDetectedCallback func(targetObject *WebView, event InsecureContentEvent)

/*
ConnectInsecureContentDetected connects the callback to the 'insecure-content-detected' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectInsecureContentDetected to remove it.
*/
func (recv *WebView) ConnectInsecureContentDetected(callback WebViewSignalInsecureContentDetectedCallback) int {
	signalWebViewInsecureContentDetectedLock.Lock()
	defer signalWebViewInsecureContentDetectedLock.Unlock()

	signalWebViewInsecureContentDetectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_insecure_content_detected(instance, C.gpointer(uintptr(signalWebViewInsecureContentDetectedId)))

	detail := signalWebViewInsecureContentDetectedDetail{callback, handlerID}
	signalWebViewInsecureContentDetectedMap[signalWebViewInsecureContentDetectedId] = detail

	return signalWebViewInsecureContentDetectedId
}

/*
DisconnectInsecureContentDetected disconnects a callback from the 'insecure-content-detected' signal for the WebView.

The connectionID should be a value returned from a call to ConnectInsecureContentDetected.
*/
func (recv *WebView) DisconnectInsecureContentDetected(connectionID int) {
	signalWebViewInsecureContentDetectedLock.Lock()
	defer signalWebViewInsecureContentDetectedLock.Unlock()

	detail, exists := signalWebViewInsecureContentDetectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewInsecureContentDetectedMap, connectionID)
}

//export webview_insecureContentDetectedHandler
func webview_insecureContentDetectedHandler(c_targetObject *C.GObject, c_event C.WebKitInsecureContentEvent, data C.gpointer) {
	signalWebViewInsecureContentDetectedLock.RLock()
	defer signalWebViewInsecureContentDetectedLock.RUnlock()

	event := InsecureContentEvent(c_event)

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewInsecureContentDetectedMap[index].callback
	callback(targetObject, event)
}

type signalWebViewLeaveFullscreenDetail struct {
	callback  WebViewSignalLeaveFullscreenCallback
	handlerID C.gulong
}

var signalWebViewLeaveFullscreenId int
var signalWebViewLeaveFullscreenMap = make(map[int]signalWebViewLeaveFullscreenDetail)
var signalWebViewLeaveFullscreenLock sync.RWMutex

// WebViewSignalLeaveFullscreenCallback is a callback function for a 'leave-fullscreen' signal emitted from a WebView.
type WebViewSignalLeaveFullscreenCallback func(targetObject *WebView) bool

/*
ConnectLeaveFullscreen connects the callback to the 'leave-fullscreen' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectLeaveFullscreen to remove it.
*/
func (recv *WebView) ConnectLeaveFullscreen(callback WebViewSignalLeaveFullscreenCallback) int {
	signalWebViewLeaveFullscreenLock.Lock()
	defer signalWebViewLeaveFullscreenLock.Unlock()

	signalWebViewLeaveFullscreenId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_leave_fullscreen(instance, C.gpointer(uintptr(signalWebViewLeaveFullscreenId)))

	detail := signalWebViewLeaveFullscreenDetail{callback, handlerID}
	signalWebViewLeaveFullscreenMap[signalWebViewLeaveFullscreenId] = detail

	return signalWebViewLeaveFullscreenId
}

/*
DisconnectLeaveFullscreen disconnects a callback from the 'leave-fullscreen' signal for the WebView.

The connectionID should be a value returned from a call to ConnectLeaveFullscreen.
*/
func (recv *WebView) DisconnectLeaveFullscreen(connectionID int) {
	signalWebViewLeaveFullscreenLock.Lock()
	defer signalWebViewLeaveFullscreenLock.Unlock()

	detail, exists := signalWebViewLeaveFullscreenMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewLeaveFullscreenMap, connectionID)
}

//export webview_leaveFullscreenHandler
func webview_leaveFullscreenHandler(c_targetObject *C.GObject, data C.gpointer) C.gboolean {
	signalWebViewLeaveFullscreenLock.RLock()
	defer signalWebViewLeaveFullscreenLock.RUnlock()

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewLeaveFullscreenMap[index].callback
	retGo := callback(targetObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebViewLoadChangedDetail struct {
	callback  WebViewSignalLoadChangedCallback
	handlerID C.gulong
}

var signalWebViewLoadChangedId int
var signalWebViewLoadChangedMap = make(map[int]signalWebViewLoadChangedDetail)
var signalWebViewLoadChangedLock sync.RWMutex

// WebViewSignalLoadChangedCallback is a callback function for a 'load-changed' signal emitted from a WebView.
type WebViewSignalLoadChangedCallback func(targetObject *WebView, loadEvent LoadEvent)

/*
ConnectLoadChanged connects the callback to the 'load-changed' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectLoadChanged to remove it.
*/
func (recv *WebView) ConnectLoadChanged(callback WebViewSignalLoadChangedCallback) int {
	signalWebViewLoadChangedLock.Lock()
	defer signalWebViewLoadChangedLock.Unlock()

	signalWebViewLoadChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_load_changed(instance, C.gpointer(uintptr(signalWebViewLoadChangedId)))

	detail := signalWebViewLoadChangedDetail{callback, handlerID}
	signalWebViewLoadChangedMap[signalWebViewLoadChangedId] = detail

	return signalWebViewLoadChangedId
}

/*
DisconnectLoadChanged disconnects a callback from the 'load-changed' signal for the WebView.

The connectionID should be a value returned from a call to ConnectLoadChanged.
*/
func (recv *WebView) DisconnectLoadChanged(connectionID int) {
	signalWebViewLoadChangedLock.Lock()
	defer signalWebViewLoadChangedLock.Unlock()

	detail, exists := signalWebViewLoadChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewLoadChangedMap, connectionID)
}

//export webview_loadChangedHandler
func webview_loadChangedHandler(c_targetObject *C.GObject, c_load_event C.WebKitLoadEvent, data C.gpointer) {
	signalWebViewLoadChangedLock.RLock()
	defer signalWebViewLoadChangedLock.RUnlock()

	loadEvent := LoadEvent(c_load_event)

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewLoadChangedMap[index].callback
	callback(targetObject, loadEvent)
}

type signalWebViewLoadFailedDetail struct {
	callback  WebViewSignalLoadFailedCallback
	handlerID C.gulong
}

var signalWebViewLoadFailedId int
var signalWebViewLoadFailedMap = make(map[int]signalWebViewLoadFailedDetail)
var signalWebViewLoadFailedLock sync.RWMutex

// WebViewSignalLoadFailedCallback is a callback function for a 'load-failed' signal emitted from a WebView.
type WebViewSignalLoadFailedCallback func(targetObject *WebView, loadEvent LoadEvent, failingUri string, error *glib.Error) bool

/*
ConnectLoadFailed connects the callback to the 'load-failed' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectLoadFailed to remove it.
*/
func (recv *WebView) ConnectLoadFailed(callback WebViewSignalLoadFailedCallback) int {
	signalWebViewLoadFailedLock.Lock()
	defer signalWebViewLoadFailedLock.Unlock()

	signalWebViewLoadFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_load_failed(instance, C.gpointer(uintptr(signalWebViewLoadFailedId)))

	detail := signalWebViewLoadFailedDetail{callback, handlerID}
	signalWebViewLoadFailedMap[signalWebViewLoadFailedId] = detail

	return signalWebViewLoadFailedId
}

/*
DisconnectLoadFailed disconnects a callback from the 'load-failed' signal for the WebView.

The connectionID should be a value returned from a call to ConnectLoadFailed.
*/
func (recv *WebView) DisconnectLoadFailed(connectionID int) {
	signalWebViewLoadFailedLock.Lock()
	defer signalWebViewLoadFailedLock.Unlock()

	detail, exists := signalWebViewLoadFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewLoadFailedMap, connectionID)
}

//export webview_loadFailedHandler
func webview_loadFailedHandler(c_targetObject *C.GObject, c_load_event C.WebKitLoadEvent, c_failing_uri *C.gchar, c_error *C.GError, data C.gpointer) C.gboolean {
	signalWebViewLoadFailedLock.RLock()
	defer signalWebViewLoadFailedLock.RUnlock()

	loadEvent := LoadEvent(c_load_event)

	failingUri := C.GoString(c_failing_uri)

	error := glib.ErrorNewFromC(unsafe.Pointer(c_error))

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewLoadFailedMap[index].callback
	retGo := callback(targetObject, loadEvent, failingUri, error)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebViewMouseTargetChangedDetail struct {
	callback  WebViewSignalMouseTargetChangedCallback
	handlerID C.gulong
}

var signalWebViewMouseTargetChangedId int
var signalWebViewMouseTargetChangedMap = make(map[int]signalWebViewMouseTargetChangedDetail)
var signalWebViewMouseTargetChangedLock sync.RWMutex

// WebViewSignalMouseTargetChangedCallback is a callback function for a 'mouse-target-changed' signal emitted from a WebView.
type WebViewSignalMouseTargetChangedCallback func(targetObject *WebView, hitTestResult *HitTestResult, modifiers uint32)

/*
ConnectMouseTargetChanged connects the callback to the 'mouse-target-changed' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectMouseTargetChanged to remove it.
*/
func (recv *WebView) ConnectMouseTargetChanged(callback WebViewSignalMouseTargetChangedCallback) int {
	signalWebViewMouseTargetChangedLock.Lock()
	defer signalWebViewMouseTargetChangedLock.Unlock()

	signalWebViewMouseTargetChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_mouse_target_changed(instance, C.gpointer(uintptr(signalWebViewMouseTargetChangedId)))

	detail := signalWebViewMouseTargetChangedDetail{callback, handlerID}
	signalWebViewMouseTargetChangedMap[signalWebViewMouseTargetChangedId] = detail

	return signalWebViewMouseTargetChangedId
}

/*
DisconnectMouseTargetChanged disconnects a callback from the 'mouse-target-changed' signal for the WebView.

The connectionID should be a value returned from a call to ConnectMouseTargetChanged.
*/
func (recv *WebView) DisconnectMouseTargetChanged(connectionID int) {
	signalWebViewMouseTargetChangedLock.Lock()
	defer signalWebViewMouseTargetChangedLock.Unlock()

	detail, exists := signalWebViewMouseTargetChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewMouseTargetChangedMap, connectionID)
}

//export webview_mouseTargetChangedHandler
func webview_mouseTargetChangedHandler(c_targetObject *C.GObject, c_hit_test_result *C.WebKitHitTestResult, c_modifiers C.guint, data C.gpointer) {
	signalWebViewMouseTargetChangedLock.RLock()
	defer signalWebViewMouseTargetChangedLock.RUnlock()

	hitTestResult := HitTestResultNewFromC(unsafe.Pointer(c_hit_test_result))

	modifiers := uint32(c_modifiers)

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewMouseTargetChangedMap[index].callback
	callback(targetObject, hitTestResult, modifiers)
}

type signalWebViewPermissionRequestDetail struct {
	callback  WebViewSignalPermissionRequestCallback
	handlerID C.gulong
}

var signalWebViewPermissionRequestId int
var signalWebViewPermissionRequestMap = make(map[int]signalWebViewPermissionRequestDetail)
var signalWebViewPermissionRequestLock sync.RWMutex

// WebViewSignalPermissionRequestCallback is a callback function for a 'permission-request' signal emitted from a WebView.
type WebViewSignalPermissionRequestCallback func(targetObject *WebView, request *PermissionRequest) bool

/*
ConnectPermissionRequest connects the callback to the 'permission-request' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectPermissionRequest to remove it.
*/
func (recv *WebView) ConnectPermissionRequest(callback WebViewSignalPermissionRequestCallback) int {
	signalWebViewPermissionRequestLock.Lock()
	defer signalWebViewPermissionRequestLock.Unlock()

	signalWebViewPermissionRequestId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_permission_request(instance, C.gpointer(uintptr(signalWebViewPermissionRequestId)))

	detail := signalWebViewPermissionRequestDetail{callback, handlerID}
	signalWebViewPermissionRequestMap[signalWebViewPermissionRequestId] = detail

	return signalWebViewPermissionRequestId
}

/*
DisconnectPermissionRequest disconnects a callback from the 'permission-request' signal for the WebView.

The connectionID should be a value returned from a call to ConnectPermissionRequest.
*/
func (recv *WebView) DisconnectPermissionRequest(connectionID int) {
	signalWebViewPermissionRequestLock.Lock()
	defer signalWebViewPermissionRequestLock.Unlock()

	detail, exists := signalWebViewPermissionRequestMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewPermissionRequestMap, connectionID)
}

//export webview_permissionRequestHandler
func webview_permissionRequestHandler(c_targetObject *C.GObject, c_request *C.WebKitPermissionRequest, data C.gpointer) C.gboolean {
	signalWebViewPermissionRequestLock.RLock()
	defer signalWebViewPermissionRequestLock.RUnlock()

	request := PermissionRequestNewFromC(unsafe.Pointer(c_request))

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewPermissionRequestMap[index].callback
	retGo := callback(targetObject, request)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebViewPrintDetail struct {
	callback  WebViewSignalPrintCallback
	handlerID C.gulong
}

var signalWebViewPrintId int
var signalWebViewPrintMap = make(map[int]signalWebViewPrintDetail)
var signalWebViewPrintLock sync.RWMutex

// WebViewSignalPrintCallback is a callback function for a 'print' signal emitted from a WebView.
type WebViewSignalPrintCallback func(targetObject *WebView, printOperation *PrintOperation) bool

/*
ConnectPrint connects the callback to the 'print' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectPrint to remove it.
*/
func (recv *WebView) ConnectPrint(callback WebViewSignalPrintCallback) int {
	signalWebViewPrintLock.Lock()
	defer signalWebViewPrintLock.Unlock()

	signalWebViewPrintId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_print(instance, C.gpointer(uintptr(signalWebViewPrintId)))

	detail := signalWebViewPrintDetail{callback, handlerID}
	signalWebViewPrintMap[signalWebViewPrintId] = detail

	return signalWebViewPrintId
}

/*
DisconnectPrint disconnects a callback from the 'print' signal for the WebView.

The connectionID should be a value returned from a call to ConnectPrint.
*/
func (recv *WebView) DisconnectPrint(connectionID int) {
	signalWebViewPrintLock.Lock()
	defer signalWebViewPrintLock.Unlock()

	detail, exists := signalWebViewPrintMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewPrintMap, connectionID)
}

//export webview_printHandler
func webview_printHandler(c_targetObject *C.GObject, c_print_operation *C.WebKitPrintOperation, data C.gpointer) C.gboolean {
	signalWebViewPrintLock.RLock()
	defer signalWebViewPrintLock.RUnlock()

	printOperation := PrintOperationNewFromC(unsafe.Pointer(c_print_operation))

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewPrintMap[index].callback
	retGo := callback(targetObject, printOperation)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebViewReadyToShowDetail struct {
	callback  WebViewSignalReadyToShowCallback
	handlerID C.gulong
}

var signalWebViewReadyToShowId int
var signalWebViewReadyToShowMap = make(map[int]signalWebViewReadyToShowDetail)
var signalWebViewReadyToShowLock sync.RWMutex

// WebViewSignalReadyToShowCallback is a callback function for a 'ready-to-show' signal emitted from a WebView.
type WebViewSignalReadyToShowCallback func(targetObject *WebView)

/*
ConnectReadyToShow connects the callback to the 'ready-to-show' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectReadyToShow to remove it.
*/
func (recv *WebView) ConnectReadyToShow(callback WebViewSignalReadyToShowCallback) int {
	signalWebViewReadyToShowLock.Lock()
	defer signalWebViewReadyToShowLock.Unlock()

	signalWebViewReadyToShowId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_ready_to_show(instance, C.gpointer(uintptr(signalWebViewReadyToShowId)))

	detail := signalWebViewReadyToShowDetail{callback, handlerID}
	signalWebViewReadyToShowMap[signalWebViewReadyToShowId] = detail

	return signalWebViewReadyToShowId
}

/*
DisconnectReadyToShow disconnects a callback from the 'ready-to-show' signal for the WebView.

The connectionID should be a value returned from a call to ConnectReadyToShow.
*/
func (recv *WebView) DisconnectReadyToShow(connectionID int) {
	signalWebViewReadyToShowLock.Lock()
	defer signalWebViewReadyToShowLock.Unlock()

	detail, exists := signalWebViewReadyToShowMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewReadyToShowMap, connectionID)
}

//export webview_readyToShowHandler
func webview_readyToShowHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalWebViewReadyToShowLock.RLock()
	defer signalWebViewReadyToShowLock.RUnlock()

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewReadyToShowMap[index].callback
	callback(targetObject)
}

type signalWebViewResourceLoadStartedDetail struct {
	callback  WebViewSignalResourceLoadStartedCallback
	handlerID C.gulong
}

var signalWebViewResourceLoadStartedId int
var signalWebViewResourceLoadStartedMap = make(map[int]signalWebViewResourceLoadStartedDetail)
var signalWebViewResourceLoadStartedLock sync.RWMutex

// WebViewSignalResourceLoadStartedCallback is a callback function for a 'resource-load-started' signal emitted from a WebView.
type WebViewSignalResourceLoadStartedCallback func(targetObject *WebView, resource *WebResource, request *URIRequest)

/*
ConnectResourceLoadStarted connects the callback to the 'resource-load-started' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectResourceLoadStarted to remove it.
*/
func (recv *WebView) ConnectResourceLoadStarted(callback WebViewSignalResourceLoadStartedCallback) int {
	signalWebViewResourceLoadStartedLock.Lock()
	defer signalWebViewResourceLoadStartedLock.Unlock()

	signalWebViewResourceLoadStartedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_resource_load_started(instance, C.gpointer(uintptr(signalWebViewResourceLoadStartedId)))

	detail := signalWebViewResourceLoadStartedDetail{callback, handlerID}
	signalWebViewResourceLoadStartedMap[signalWebViewResourceLoadStartedId] = detail

	return signalWebViewResourceLoadStartedId
}

/*
DisconnectResourceLoadStarted disconnects a callback from the 'resource-load-started' signal for the WebView.

The connectionID should be a value returned from a call to ConnectResourceLoadStarted.
*/
func (recv *WebView) DisconnectResourceLoadStarted(connectionID int) {
	signalWebViewResourceLoadStartedLock.Lock()
	defer signalWebViewResourceLoadStartedLock.Unlock()

	detail, exists := signalWebViewResourceLoadStartedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewResourceLoadStartedMap, connectionID)
}

//export webview_resourceLoadStartedHandler
func webview_resourceLoadStartedHandler(c_targetObject *C.GObject, c_resource *C.WebKitWebResource, c_request *C.WebKitURIRequest, data C.gpointer) {
	signalWebViewResourceLoadStartedLock.RLock()
	defer signalWebViewResourceLoadStartedLock.RUnlock()

	resource := WebResourceNewFromC(unsafe.Pointer(c_resource))

	request := URIRequestNewFromC(unsafe.Pointer(c_request))

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewResourceLoadStartedMap[index].callback
	callback(targetObject, resource, request)
}

type signalWebViewRunAsModalDetail struct {
	callback  WebViewSignalRunAsModalCallback
	handlerID C.gulong
}

var signalWebViewRunAsModalId int
var signalWebViewRunAsModalMap = make(map[int]signalWebViewRunAsModalDetail)
var signalWebViewRunAsModalLock sync.RWMutex

// WebViewSignalRunAsModalCallback is a callback function for a 'run-as-modal' signal emitted from a WebView.
type WebViewSignalRunAsModalCallback func(targetObject *WebView)

/*
ConnectRunAsModal connects the callback to the 'run-as-modal' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectRunAsModal to remove it.
*/
func (recv *WebView) ConnectRunAsModal(callback WebViewSignalRunAsModalCallback) int {
	signalWebViewRunAsModalLock.Lock()
	defer signalWebViewRunAsModalLock.Unlock()

	signalWebViewRunAsModalId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_run_as_modal(instance, C.gpointer(uintptr(signalWebViewRunAsModalId)))

	detail := signalWebViewRunAsModalDetail{callback, handlerID}
	signalWebViewRunAsModalMap[signalWebViewRunAsModalId] = detail

	return signalWebViewRunAsModalId
}

/*
DisconnectRunAsModal disconnects a callback from the 'run-as-modal' signal for the WebView.

The connectionID should be a value returned from a call to ConnectRunAsModal.
*/
func (recv *WebView) DisconnectRunAsModal(connectionID int) {
	signalWebViewRunAsModalLock.Lock()
	defer signalWebViewRunAsModalLock.Unlock()

	detail, exists := signalWebViewRunAsModalMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewRunAsModalMap, connectionID)
}

//export webview_runAsModalHandler
func webview_runAsModalHandler(c_targetObject *C.GObject, data C.gpointer) {
	signalWebViewRunAsModalLock.RLock()
	defer signalWebViewRunAsModalLock.RUnlock()

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewRunAsModalMap[index].callback
	callback(targetObject)
}

type signalWebViewRunFileChooserDetail struct {
	callback  WebViewSignalRunFileChooserCallback
	handlerID C.gulong
}

var signalWebViewRunFileChooserId int
var signalWebViewRunFileChooserMap = make(map[int]signalWebViewRunFileChooserDetail)
var signalWebViewRunFileChooserLock sync.RWMutex

// WebViewSignalRunFileChooserCallback is a callback function for a 'run-file-chooser' signal emitted from a WebView.
type WebViewSignalRunFileChooserCallback func(targetObject *WebView, request *FileChooserRequest) bool

/*
ConnectRunFileChooser connects the callback to the 'run-file-chooser' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectRunFileChooser to remove it.
*/
func (recv *WebView) ConnectRunFileChooser(callback WebViewSignalRunFileChooserCallback) int {
	signalWebViewRunFileChooserLock.Lock()
	defer signalWebViewRunFileChooserLock.Unlock()

	signalWebViewRunFileChooserId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_run_file_chooser(instance, C.gpointer(uintptr(signalWebViewRunFileChooserId)))

	detail := signalWebViewRunFileChooserDetail{callback, handlerID}
	signalWebViewRunFileChooserMap[signalWebViewRunFileChooserId] = detail

	return signalWebViewRunFileChooserId
}

/*
DisconnectRunFileChooser disconnects a callback from the 'run-file-chooser' signal for the WebView.

The connectionID should be a value returned from a call to ConnectRunFileChooser.
*/
func (recv *WebView) DisconnectRunFileChooser(connectionID int) {
	signalWebViewRunFileChooserLock.Lock()
	defer signalWebViewRunFileChooserLock.Unlock()

	detail, exists := signalWebViewRunFileChooserMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewRunFileChooserMap, connectionID)
}

//export webview_runFileChooserHandler
func webview_runFileChooserHandler(c_targetObject *C.GObject, c_request *C.WebKitFileChooserRequest, data C.gpointer) C.gboolean {
	signalWebViewRunFileChooserLock.RLock()
	defer signalWebViewRunFileChooserLock.RUnlock()

	request := FileChooserRequestNewFromC(unsafe.Pointer(c_request))

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewRunFileChooserMap[index].callback
	retGo := callback(targetObject, request)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWebViewScriptDialogDetail struct {
	callback  WebViewSignalScriptDialogCallback
	handlerID C.gulong
}

var signalWebViewScriptDialogId int
var signalWebViewScriptDialogMap = make(map[int]signalWebViewScriptDialogDetail)
var signalWebViewScriptDialogLock sync.RWMutex

// WebViewSignalScriptDialogCallback is a callback function for a 'script-dialog' signal emitted from a WebView.
type WebViewSignalScriptDialogCallback func(targetObject *WebView, dialog *ScriptDialog) bool

/*
ConnectScriptDialog connects the callback to the 'script-dialog' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectScriptDialog to remove it.
*/
func (recv *WebView) ConnectScriptDialog(callback WebViewSignalScriptDialogCallback) int {
	signalWebViewScriptDialogLock.Lock()
	defer signalWebViewScriptDialogLock.Unlock()

	signalWebViewScriptDialogId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_script_dialog(instance, C.gpointer(uintptr(signalWebViewScriptDialogId)))

	detail := signalWebViewScriptDialogDetail{callback, handlerID}
	signalWebViewScriptDialogMap[signalWebViewScriptDialogId] = detail

	return signalWebViewScriptDialogId
}

/*
DisconnectScriptDialog disconnects a callback from the 'script-dialog' signal for the WebView.

The connectionID should be a value returned from a call to ConnectScriptDialog.
*/
func (recv *WebView) DisconnectScriptDialog(connectionID int) {
	signalWebViewScriptDialogLock.Lock()
	defer signalWebViewScriptDialogLock.Unlock()

	detail, exists := signalWebViewScriptDialogMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewScriptDialogMap, connectionID)
}

//export webview_scriptDialogHandler
func webview_scriptDialogHandler(c_targetObject *C.GObject, c_dialog *C.WebKitScriptDialog, data C.gpointer) C.gboolean {
	signalWebViewScriptDialogLock.RLock()
	defer signalWebViewScriptDialogLock.RUnlock()

	dialog := ScriptDialogNewFromC(unsafe.Pointer(c_dialog))

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewScriptDialogMap[index].callback
	retGo := callback(targetObject, dialog)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Blacklisted signal show-option-menu

type signalWebViewSubmitFormDetail struct {
	callback  WebViewSignalSubmitFormCallback
	handlerID C.gulong
}

var signalWebViewSubmitFormId int
var signalWebViewSubmitFormMap = make(map[int]signalWebViewSubmitFormDetail)
var signalWebViewSubmitFormLock sync.RWMutex

// WebViewSignalSubmitFormCallback is a callback function for a 'submit-form' signal emitted from a WebView.
type WebViewSignalSubmitFormCallback func(targetObject *WebView, request *FormSubmissionRequest)

/*
ConnectSubmitForm connects the callback to the 'submit-form' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectSubmitForm to remove it.
*/
func (recv *WebView) ConnectSubmitForm(callback WebViewSignalSubmitFormCallback) int {
	signalWebViewSubmitFormLock.Lock()
	defer signalWebViewSubmitFormLock.Unlock()

	signalWebViewSubmitFormId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_submit_form(instance, C.gpointer(uintptr(signalWebViewSubmitFormId)))

	detail := signalWebViewSubmitFormDetail{callback, handlerID}
	signalWebViewSubmitFormMap[signalWebViewSubmitFormId] = detail

	return signalWebViewSubmitFormId
}

/*
DisconnectSubmitForm disconnects a callback from the 'submit-form' signal for the WebView.

The connectionID should be a value returned from a call to ConnectSubmitForm.
*/
func (recv *WebView) DisconnectSubmitForm(connectionID int) {
	signalWebViewSubmitFormLock.Lock()
	defer signalWebViewSubmitFormLock.Unlock()

	detail, exists := signalWebViewSubmitFormMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewSubmitFormMap, connectionID)
}

//export webview_submitFormHandler
func webview_submitFormHandler(c_targetObject *C.GObject, c_request *C.WebKitFormSubmissionRequest, data C.gpointer) {
	signalWebViewSubmitFormLock.RLock()
	defer signalWebViewSubmitFormLock.RUnlock()

	request := FormSubmissionRequestNewFromC(unsafe.Pointer(c_request))

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewSubmitFormMap[index].callback
	callback(targetObject, request)
}

type signalWebViewWebProcessCrashedDetail struct {
	callback  WebViewSignalWebProcessCrashedCallback
	handlerID C.gulong
}

var signalWebViewWebProcessCrashedId int
var signalWebViewWebProcessCrashedMap = make(map[int]signalWebViewWebProcessCrashedDetail)
var signalWebViewWebProcessCrashedLock sync.RWMutex

// WebViewSignalWebProcessCrashedCallback is a callback function for a 'web-process-crashed' signal emitted from a WebView.
type WebViewSignalWebProcessCrashedCallback func(targetObject *WebView) bool

/*
ConnectWebProcessCrashed connects the callback to the 'web-process-crashed' signal for the WebView.

The returned value represents the connection, and may be passed to DisconnectWebProcessCrashed to remove it.
*/
func (recv *WebView) ConnectWebProcessCrashed(callback WebViewSignalWebProcessCrashedCallback) int {
	signalWebViewWebProcessCrashedLock.Lock()
	defer signalWebViewWebProcessCrashedLock.Unlock()

	signalWebViewWebProcessCrashedId++
	instance := C.gpointer(recv.native)
	handlerID := C.WebView_signal_connect_web_process_crashed(instance, C.gpointer(uintptr(signalWebViewWebProcessCrashedId)))

	detail := signalWebViewWebProcessCrashedDetail{callback, handlerID}
	signalWebViewWebProcessCrashedMap[signalWebViewWebProcessCrashedId] = detail

	return signalWebViewWebProcessCrashedId
}

/*
DisconnectWebProcessCrashed disconnects a callback from the 'web-process-crashed' signal for the WebView.

The connectionID should be a value returned from a call to ConnectWebProcessCrashed.
*/
func (recv *WebView) DisconnectWebProcessCrashed(connectionID int) {
	signalWebViewWebProcessCrashedLock.Lock()
	defer signalWebViewWebProcessCrashedLock.Unlock()

	detail, exists := signalWebViewWebProcessCrashedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWebViewWebProcessCrashedMap, connectionID)
}

//export webview_webProcessCrashedHandler
func webview_webProcessCrashedHandler(c_targetObject *C.GObject, data C.gpointer) C.gboolean {
	signalWebViewWebProcessCrashedLock.RLock()
	defer signalWebViewWebProcessCrashedLock.RUnlock()

	targetObject := WebViewNewFromC((unsafe.Pointer)(c_targetObject))

	index := int(uintptr(data))
	callback := signalWebViewWebProcessCrashedMap[index].callback
	retGo := callback(targetObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// WebViewNew is a wrapper around the C function webkit_web_view_new.
func WebViewNew() *WebView {
	retC := C.webkit_web_view_new()
	retGo := WebViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// WebViewNewWithContext is a wrapper around the C function webkit_web_view_new_with_context.
func WebViewNewWithContext(context *WebContext) *WebView {
	c_context := (*C.WebKitWebContext)(C.NULL)
	if context != nil {
		c_context = (*C.WebKitWebContext)(context.ToC())
	}

	retC := C.webkit_web_view_new_with_context(c_context)
	retGo := WebViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : webkit_web_view_can_execute_editing_command : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CanExecuteEditingCommandFinish is a wrapper around the C function webkit_web_view_can_execute_editing_command_finish.
func (recv *WebView) CanExecuteEditingCommandFinish(result *gio.AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_web_view_can_execute_editing_command_finish((*C.WebKitWebView)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CanGoBack is a wrapper around the C function webkit_web_view_can_go_back.
func (recv *WebView) CanGoBack() bool {
	retC := C.webkit_web_view_can_go_back((*C.WebKitWebView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanGoForward is a wrapper around the C function webkit_web_view_can_go_forward.
func (recv *WebView) CanGoForward() bool {
	retC := C.webkit_web_view_can_go_forward((*C.WebKitWebView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanShowMimeType is a wrapper around the C function webkit_web_view_can_show_mime_type.
func (recv *WebView) CanShowMimeType(mimeType string) bool {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	retC := C.webkit_web_view_can_show_mime_type((*C.WebKitWebView)(recv.native), c_mime_type)
	retGo := retC == C.TRUE

	return retGo
}

// DownloadUri is a wrapper around the C function webkit_web_view_download_uri.
func (recv *WebView) DownloadUri(uri string) *Download {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.webkit_web_view_download_uri((*C.WebKitWebView)(recv.native), c_uri)
	retGo := DownloadNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ExecuteEditingCommand is a wrapper around the C function webkit_web_view_execute_editing_command.
func (recv *WebView) ExecuteEditingCommand(command string) {
	c_command := C.CString(command)
	defer C.free(unsafe.Pointer(c_command))

	C.webkit_web_view_execute_editing_command((*C.WebKitWebView)(recv.native), c_command)

	return
}

// GetBackForwardList is a wrapper around the C function webkit_web_view_get_back_forward_list.
func (recv *WebView) GetBackForwardList() *BackForwardList {
	retC := C.webkit_web_view_get_back_forward_list((*C.WebKitWebView)(recv.native))
	retGo := BackForwardListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContext is a wrapper around the C function webkit_web_view_get_context.
func (recv *WebView) GetContext() *WebContext {
	retC := C.webkit_web_view_get_context((*C.WebKitWebView)(recv.native))
	retGo := WebContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCustomCharset is a wrapper around the C function webkit_web_view_get_custom_charset.
func (recv *WebView) GetCustomCharset() string {
	retC := C.webkit_web_view_get_custom_charset((*C.WebKitWebView)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetEstimatedLoadProgress is a wrapper around the C function webkit_web_view_get_estimated_load_progress.
func (recv *WebView) GetEstimatedLoadProgress() float64 {
	retC := C.webkit_web_view_get_estimated_load_progress((*C.WebKitWebView)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetFavicon is a wrapper around the C function webkit_web_view_get_favicon.
func (recv *WebView) GetFavicon() *cairo.Surface {
	retC := C.webkit_web_view_get_favicon((*C.WebKitWebView)(recv.native))
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFindController is a wrapper around the C function webkit_web_view_get_find_controller.
func (recv *WebView) GetFindController() *FindController {
	retC := C.webkit_web_view_get_find_controller((*C.WebKitWebView)(recv.native))
	retGo := FindControllerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInspector is a wrapper around the C function webkit_web_view_get_inspector.
func (recv *WebView) GetInspector() *WebInspector {
	retC := C.webkit_web_view_get_inspector((*C.WebKitWebView)(recv.native))
	retGo := WebInspectorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : webkit_web_view_get_javascript_global_context : no return generator

// GetMainResource is a wrapper around the C function webkit_web_view_get_main_resource.
func (recv *WebView) GetMainResource() *WebResource {
	retC := C.webkit_web_view_get_main_resource((*C.WebKitWebView)(recv.native))
	retGo := WebResourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPageId is a wrapper around the C function webkit_web_view_get_page_id.
func (recv *WebView) GetPageId() uint64 {
	retC := C.webkit_web_view_get_page_id((*C.WebKitWebView)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// GetSettings is a wrapper around the C function webkit_web_view_get_settings.
func (recv *WebView) GetSettings() *Settings {
	retC := C.webkit_web_view_get_settings((*C.WebKitWebView)(recv.native))
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : webkit_web_view_get_snapshot : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// GetSnapshotFinish is a wrapper around the C function webkit_web_view_get_snapshot_finish.
func (recv *WebView) GetSnapshotFinish(result *gio.AsyncResult) (*cairo.Surface, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_web_view_get_snapshot_finish((*C.WebKitWebView)(recv.native), c_result, &cThrowableError)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetTitle is a wrapper around the C function webkit_web_view_get_title.
func (recv *WebView) GetTitle() string {
	retC := C.webkit_web_view_get_title((*C.WebKitWebView)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTlsInfo is a wrapper around the C function webkit_web_view_get_tls_info.
func (recv *WebView) GetTlsInfo() (bool, *gio.TlsCertificate, gio.TlsCertificateFlags) {
	var c_certificate *C.GTlsCertificate

	var c_errors C.GTlsCertificateFlags

	retC := C.webkit_web_view_get_tls_info((*C.WebKitWebView)(recv.native), &c_certificate, &c_errors)
	retGo := retC == C.TRUE

	certificate := gio.TlsCertificateNewFromC(unsafe.Pointer(c_certificate))

	errors := (gio.TlsCertificateFlags)(c_errors)

	return retGo, certificate, errors
}

// GetUri is a wrapper around the C function webkit_web_view_get_uri.
func (recv *WebView) GetUri() string {
	retC := C.webkit_web_view_get_uri((*C.WebKitWebView)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetWindowProperties is a wrapper around the C function webkit_web_view_get_window_properties.
func (recv *WebView) GetWindowProperties() *WindowProperties {
	retC := C.webkit_web_view_get_window_properties((*C.WebKitWebView)(recv.native))
	retGo := WindowPropertiesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetZoomLevel is a wrapper around the C function webkit_web_view_get_zoom_level.
func (recv *WebView) GetZoomLevel() float64 {
	retC := C.webkit_web_view_get_zoom_level((*C.WebKitWebView)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GoBack is a wrapper around the C function webkit_web_view_go_back.
func (recv *WebView) GoBack() {
	C.webkit_web_view_go_back((*C.WebKitWebView)(recv.native))

	return
}

// GoForward is a wrapper around the C function webkit_web_view_go_forward.
func (recv *WebView) GoForward() {
	C.webkit_web_view_go_forward((*C.WebKitWebView)(recv.native))

	return
}

// GoToBackForwardListItem is a wrapper around the C function webkit_web_view_go_to_back_forward_list_item.
func (recv *WebView) GoToBackForwardListItem(listItem *BackForwardListItem) {
	c_list_item := (*C.WebKitBackForwardListItem)(C.NULL)
	if listItem != nil {
		c_list_item = (*C.WebKitBackForwardListItem)(listItem.ToC())
	}

	C.webkit_web_view_go_to_back_forward_list_item((*C.WebKitWebView)(recv.native), c_list_item)

	return
}

// IsEditable is a wrapper around the C function webkit_web_view_is_editable.
func (recv *WebView) IsEditable() bool {
	retC := C.webkit_web_view_is_editable((*C.WebKitWebView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsLoading is a wrapper around the C function webkit_web_view_is_loading.
func (recv *WebView) IsLoading() bool {
	retC := C.webkit_web_view_is_loading((*C.WebKitWebView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// LoadAlternateHtml is a wrapper around the C function webkit_web_view_load_alternate_html.
func (recv *WebView) LoadAlternateHtml(content string, contentUri string, baseUri string) {
	c_content := C.CString(content)
	defer C.free(unsafe.Pointer(c_content))

	c_content_uri := C.CString(contentUri)
	defer C.free(unsafe.Pointer(c_content_uri))

	c_base_uri := C.CString(baseUri)
	defer C.free(unsafe.Pointer(c_base_uri))

	C.webkit_web_view_load_alternate_html((*C.WebKitWebView)(recv.native), c_content, c_content_uri, c_base_uri)

	return
}

// LoadHtml is a wrapper around the C function webkit_web_view_load_html.
func (recv *WebView) LoadHtml(content string, baseUri string) {
	c_content := C.CString(content)
	defer C.free(unsafe.Pointer(c_content))

	c_base_uri := C.CString(baseUri)
	defer C.free(unsafe.Pointer(c_base_uri))

	C.webkit_web_view_load_html((*C.WebKitWebView)(recv.native), c_content, c_base_uri)

	return
}

// LoadPlainText is a wrapper around the C function webkit_web_view_load_plain_text.
func (recv *WebView) LoadPlainText(plainText string) {
	c_plain_text := C.CString(plainText)
	defer C.free(unsafe.Pointer(c_plain_text))

	C.webkit_web_view_load_plain_text((*C.WebKitWebView)(recv.native), c_plain_text)

	return
}

// LoadRequest is a wrapper around the C function webkit_web_view_load_request.
func (recv *WebView) LoadRequest(request *URIRequest) {
	c_request := (*C.WebKitURIRequest)(C.NULL)
	if request != nil {
		c_request = (*C.WebKitURIRequest)(request.ToC())
	}

	C.webkit_web_view_load_request((*C.WebKitWebView)(recv.native), c_request)

	return
}

// LoadUri is a wrapper around the C function webkit_web_view_load_uri.
func (recv *WebView) LoadUri(uri string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.webkit_web_view_load_uri((*C.WebKitWebView)(recv.native), c_uri)

	return
}

// Reload is a wrapper around the C function webkit_web_view_reload.
func (recv *WebView) Reload() {
	C.webkit_web_view_reload((*C.WebKitWebView)(recv.native))

	return
}

// ReloadBypassCache is a wrapper around the C function webkit_web_view_reload_bypass_cache.
func (recv *WebView) ReloadBypassCache() {
	C.webkit_web_view_reload_bypass_cache((*C.WebKitWebView)(recv.native))

	return
}

// Unsupported : webkit_web_view_run_javascript : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// RunJavascriptFinish is a wrapper around the C function webkit_web_view_run_javascript_finish.
func (recv *WebView) RunJavascriptFinish(result *gio.AsyncResult) (*JavascriptResult, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_web_view_run_javascript_finish((*C.WebKitWebView)(recv.native), c_result, &cThrowableError)
	retGo := JavascriptResultNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : webkit_web_view_run_javascript_from_gresource : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// RunJavascriptFromGresourceFinish is a wrapper around the C function webkit_web_view_run_javascript_from_gresource_finish.
func (recv *WebView) RunJavascriptFromGresourceFinish(result *gio.AsyncResult) (*JavascriptResult, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_web_view_run_javascript_from_gresource_finish((*C.WebKitWebView)(recv.native), c_result, &cThrowableError)
	retGo := JavascriptResultNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : webkit_web_view_save : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SaveFinish is a wrapper around the C function webkit_web_view_save_finish.
func (recv *WebView) SaveFinish(result *gio.AsyncResult) (*gio.InputStream, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_web_view_save_finish((*C.WebKitWebView)(recv.native), c_result, &cThrowableError)
	retGo := gio.InputStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : webkit_web_view_save_to_file : unsupported parameter callback : no type generator for Gio.AsyncReadyCallback (GAsyncReadyCallback) for param callback

// SaveToFileFinish is a wrapper around the C function webkit_web_view_save_to_file_finish.
func (recv *WebView) SaveToFileFinish(result *gio.AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.webkit_web_view_save_to_file_finish((*C.WebKitWebView)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetCustomCharset is a wrapper around the C function webkit_web_view_set_custom_charset.
func (recv *WebView) SetCustomCharset(charset string) {
	c_charset := C.CString(charset)
	defer C.free(unsafe.Pointer(c_charset))

	C.webkit_web_view_set_custom_charset((*C.WebKitWebView)(recv.native), c_charset)

	return
}

// SetSettings is a wrapper around the C function webkit_web_view_set_settings.
func (recv *WebView) SetSettings(settings *Settings) {
	c_settings := (*C.WebKitSettings)(C.NULL)
	if settings != nil {
		c_settings = (*C.WebKitSettings)(settings.ToC())
	}

	C.webkit_web_view_set_settings((*C.WebKitWebView)(recv.native), c_settings)

	return
}

// SetZoomLevel is a wrapper around the C function webkit_web_view_set_zoom_level.
func (recv *WebView) SetZoomLevel(zoomLevel float64) {
	c_zoom_level := (C.gdouble)(zoomLevel)

	C.webkit_web_view_set_zoom_level((*C.WebKitWebView)(recv.native), c_zoom_level)

	return
}

// StopLoading is a wrapper around the C function webkit_web_view_stop_loading.
func (recv *WebView) StopLoading() {
	C.webkit_web_view_stop_loading((*C.WebKitWebView)(recv.native))

	return
}

// ImplementorIface returns the ImplementorIface interface implemented by WebView
func (recv *WebView) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by WebView
func (recv *WebView) Buildable() *gtk.Buildable {
	return gtk.BuildableNewFromC(recv.ToC())
}

// WebViewBase is a wrapper around the C record WebKitWebViewBase.
type WebViewBase struct {
	native *C.WebKitWebViewBase
	// parentInstance : record
	// Private : priv
}

func WebViewBaseNewFromC(u unsafe.Pointer) *WebViewBase {
	c := (*C.WebKitWebViewBase)(u)
	if c == nil {
		return nil
	}

	g := &WebViewBase{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WebViewBase) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WebViewBase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebViewBase with another WebViewBase, and returns true if they represent the same GObject.
func (recv *WebViewBase) Equals(other *WebViewBase) bool {
	return other.ToC() == recv.ToC()
}

// Container upcasts to *Container
func (recv *WebViewBase) Container() *gtk.Container {
	return gtk.ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *WebViewBase) Widget() *gtk.Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *WebViewBase) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *WebViewBase) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitrary Object to WebViewBase.
// Exercise care, as this is a potentially dangerous function if the Object is not a WebViewBase.
func CastToWebViewBase(object *gobject.Object) *WebViewBase {
	return WebViewBaseNewFromC(object.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by WebViewBase
func (recv *WebViewBase) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by WebViewBase
func (recv *WebViewBase) Buildable() *gtk.Buildable {
	return gtk.BuildableNewFromC(recv.ToC())
}

// WebsiteDataManager is a wrapper around the C record WebKitWebsiteDataManager.
type WebsiteDataManager struct {
	native *C.WebKitWebsiteDataManager
	// parent : record
	// priv : record
}

func WebsiteDataManagerNewFromC(u unsafe.Pointer) *WebsiteDataManager {
	c := (*C.WebKitWebsiteDataManager)(u)
	if c == nil {
		return nil
	}

	g := &WebsiteDataManager{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WebsiteDataManager) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WebsiteDataManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebsiteDataManager with another WebsiteDataManager, and returns true if they represent the same GObject.
func (recv *WebsiteDataManager) Equals(other *WebsiteDataManager) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *WebsiteDataManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to WebsiteDataManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a WebsiteDataManager.
func CastToWebsiteDataManager(object *gobject.Object) *WebsiteDataManager {
	return WebsiteDataManagerNewFromC(object.ToC())
}

// WindowProperties is a wrapper around the C record WebKitWindowProperties.
type WindowProperties struct {
	native *C.WebKitWindowProperties
	// parent : record
	// Private : priv
}

func WindowPropertiesNewFromC(u unsafe.Pointer) *WindowProperties {
	c := (*C.WebKitWindowProperties)(u)
	if c == nil {
		return nil
	}

	g := &WindowProperties{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *WindowProperties) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *WindowProperties) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WindowProperties with another WindowProperties, and returns true if they represent the same GObject.
func (recv *WindowProperties) Equals(other *WindowProperties) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *WindowProperties) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to WindowProperties.
// Exercise care, as this is a potentially dangerous function if the Object is not a WindowProperties.
func CastToWindowProperties(object *gobject.Object) *WindowProperties {
	return WindowPropertiesNewFromC(object.ToC())
}

// GetFullscreen is a wrapper around the C function webkit_window_properties_get_fullscreen.
func (recv *WindowProperties) GetFullscreen() bool {
	retC := C.webkit_window_properties_get_fullscreen((*C.WebKitWindowProperties)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetGeometry is a wrapper around the C function webkit_window_properties_get_geometry.
func (recv *WindowProperties) GetGeometry() *gdk.Rectangle {
	var c_geometry C.GdkRectangle

	C.webkit_window_properties_get_geometry((*C.WebKitWindowProperties)(recv.native), &c_geometry)

	geometry := gdk.RectangleNewFromC(unsafe.Pointer(&c_geometry))

	return geometry
}

// GetLocationbarVisible is a wrapper around the C function webkit_window_properties_get_locationbar_visible.
func (recv *WindowProperties) GetLocationbarVisible() bool {
	retC := C.webkit_window_properties_get_locationbar_visible((*C.WebKitWindowProperties)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMenubarVisible is a wrapper around the C function webkit_window_properties_get_menubar_visible.
func (recv *WindowProperties) GetMenubarVisible() bool {
	retC := C.webkit_window_properties_get_menubar_visible((*C.WebKitWindowProperties)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetResizable is a wrapper around the C function webkit_window_properties_get_resizable.
func (recv *WindowProperties) GetResizable() bool {
	retC := C.webkit_window_properties_get_resizable((*C.WebKitWindowProperties)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetScrollbarsVisible is a wrapper around the C function webkit_window_properties_get_scrollbars_visible.
func (recv *WindowProperties) GetScrollbarsVisible() bool {
	retC := C.webkit_window_properties_get_scrollbars_visible((*C.WebKitWindowProperties)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetStatusbarVisible is a wrapper around the C function webkit_window_properties_get_statusbar_visible.
func (recv *WindowProperties) GetStatusbarVisible() bool {
	retC := C.webkit_window_properties_get_statusbar_visible((*C.WebKitWindowProperties)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetToolbarVisible is a wrapper around the C function webkit_window_properties_get_toolbar_visible.
func (recv *WindowProperties) GetToolbarVisible() bool {
	retC := C.webkit_window_properties_get_toolbar_visible((*C.WebKitWindowProperties)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

const EDITING_COMMAND_COPY string = C.WEBKIT_EDITING_COMMAND_COPY
const EDITING_COMMAND_CUT string = C.WEBKIT_EDITING_COMMAND_CUT
const EDITING_COMMAND_PASTE string = C.WEBKIT_EDITING_COMMAND_PASTE
const EDITING_COMMAND_REDO string = C.WEBKIT_EDITING_COMMAND_REDO
const EDITING_COMMAND_SELECT_ALL string = C.WEBKIT_EDITING_COMMAND_SELECT_ALL
const EDITING_COMMAND_UNDO string = C.WEBKIT_EDITING_COMMAND_UNDO
const MAJOR_VERSION int32 = C.WEBKIT_MAJOR_VERSION
const MICRO_VERSION int32 = C.WEBKIT_MICRO_VERSION
const MINOR_VERSION int32 = C.WEBKIT_MINOR_VERSION

type CacheModel C.WebKitCacheModel

const (
	WEBKIT_CACHE_MODEL_DOCUMENT_VIEWER  CacheModel = 0
	WEBKIT_CACHE_MODEL_WEB_BROWSER      CacheModel = 1
	WEBKIT_CACHE_MODEL_DOCUMENT_BROWSER CacheModel = 2
)

type ContextMenuAction C.WebKitContextMenuAction

const (
	WEBKIT_CONTEXT_MENU_ACTION_NO_ACTION                    ContextMenuAction = 0
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_LINK                    ContextMenuAction = 1
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_LINK_IN_NEW_WINDOW      ContextMenuAction = 2
	WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_LINK_TO_DISK        ContextMenuAction = 3
	WEBKIT_CONTEXT_MENU_ACTION_COPY_LINK_TO_CLIPBOARD       ContextMenuAction = 4
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_IMAGE_IN_NEW_WINDOW     ContextMenuAction = 5
	WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_IMAGE_TO_DISK       ContextMenuAction = 6
	WEBKIT_CONTEXT_MENU_ACTION_COPY_IMAGE_TO_CLIPBOARD      ContextMenuAction = 7
	WEBKIT_CONTEXT_MENU_ACTION_COPY_IMAGE_URL_TO_CLIPBOARD  ContextMenuAction = 8
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_FRAME_IN_NEW_WINDOW     ContextMenuAction = 9
	WEBKIT_CONTEXT_MENU_ACTION_GO_BACK                      ContextMenuAction = 10
	WEBKIT_CONTEXT_MENU_ACTION_GO_FORWARD                   ContextMenuAction = 11
	WEBKIT_CONTEXT_MENU_ACTION_STOP                         ContextMenuAction = 12
	WEBKIT_CONTEXT_MENU_ACTION_RELOAD                       ContextMenuAction = 13
	WEBKIT_CONTEXT_MENU_ACTION_COPY                         ContextMenuAction = 14
	WEBKIT_CONTEXT_MENU_ACTION_CUT                          ContextMenuAction = 15
	WEBKIT_CONTEXT_MENU_ACTION_PASTE                        ContextMenuAction = 16
	WEBKIT_CONTEXT_MENU_ACTION_DELETE                       ContextMenuAction = 17
	WEBKIT_CONTEXT_MENU_ACTION_SELECT_ALL                   ContextMenuAction = 18
	WEBKIT_CONTEXT_MENU_ACTION_INPUT_METHODS                ContextMenuAction = 19
	WEBKIT_CONTEXT_MENU_ACTION_UNICODE                      ContextMenuAction = 20
	WEBKIT_CONTEXT_MENU_ACTION_SPELLING_GUESS               ContextMenuAction = 21
	WEBKIT_CONTEXT_MENU_ACTION_NO_GUESSES_FOUND             ContextMenuAction = 22
	WEBKIT_CONTEXT_MENU_ACTION_IGNORE_SPELLING              ContextMenuAction = 23
	WEBKIT_CONTEXT_MENU_ACTION_LEARN_SPELLING               ContextMenuAction = 24
	WEBKIT_CONTEXT_MENU_ACTION_IGNORE_GRAMMAR               ContextMenuAction = 25
	WEBKIT_CONTEXT_MENU_ACTION_FONT_MENU                    ContextMenuAction = 26
	WEBKIT_CONTEXT_MENU_ACTION_BOLD                         ContextMenuAction = 27
	WEBKIT_CONTEXT_MENU_ACTION_ITALIC                       ContextMenuAction = 28
	WEBKIT_CONTEXT_MENU_ACTION_UNDERLINE                    ContextMenuAction = 29
	WEBKIT_CONTEXT_MENU_ACTION_OUTLINE                      ContextMenuAction = 30
	WEBKIT_CONTEXT_MENU_ACTION_INSPECT_ELEMENT              ContextMenuAction = 31
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_VIDEO_IN_NEW_WINDOW     ContextMenuAction = 32
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_AUDIO_IN_NEW_WINDOW     ContextMenuAction = 33
	WEBKIT_CONTEXT_MENU_ACTION_COPY_VIDEO_LINK_TO_CLIPBOARD ContextMenuAction = 34
	WEBKIT_CONTEXT_MENU_ACTION_COPY_AUDIO_LINK_TO_CLIPBOARD ContextMenuAction = 35
	WEBKIT_CONTEXT_MENU_ACTION_TOGGLE_MEDIA_CONTROLS        ContextMenuAction = 36
	WEBKIT_CONTEXT_MENU_ACTION_TOGGLE_MEDIA_LOOP            ContextMenuAction = 37
	WEBKIT_CONTEXT_MENU_ACTION_ENTER_VIDEO_FULLSCREEN       ContextMenuAction = 38
	WEBKIT_CONTEXT_MENU_ACTION_MEDIA_PLAY                   ContextMenuAction = 39
	WEBKIT_CONTEXT_MENU_ACTION_MEDIA_PAUSE                  ContextMenuAction = 40
	WEBKIT_CONTEXT_MENU_ACTION_MEDIA_MUTE                   ContextMenuAction = 41
	WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_VIDEO_TO_DISK       ContextMenuAction = 42
	WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_AUDIO_TO_DISK       ContextMenuAction = 43
	WEBKIT_CONTEXT_MENU_ACTION_INSERT_EMOJI                 ContextMenuAction = 44
	WEBKIT_CONTEXT_MENU_ACTION_CUSTOM                       ContextMenuAction = 10000
)

type CookieAcceptPolicy C.WebKitCookieAcceptPolicy

const (
	WEBKIT_COOKIE_POLICY_ACCEPT_ALWAYS         CookieAcceptPolicy = 0
	WEBKIT_COOKIE_POLICY_ACCEPT_NEVER          CookieAcceptPolicy = 1
	WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY CookieAcceptPolicy = 2
)

type CookiePersistentStorage C.WebKitCookiePersistentStorage

const (
	WEBKIT_COOKIE_PERSISTENT_STORAGE_TEXT   CookiePersistentStorage = 0
	WEBKIT_COOKIE_PERSISTENT_STORAGE_SQLITE CookiePersistentStorage = 1
)

type DownloadError C.WebKitDownloadError

const (
	WEBKIT_DOWNLOAD_ERROR_NETWORK           DownloadError = 499
	WEBKIT_DOWNLOAD_ERROR_CANCELLED_BY_USER DownloadError = 400
	WEBKIT_DOWNLOAD_ERROR_DESTINATION       DownloadError = 401
)

// DownloadErrorQuark is a wrapper around the C function webkit_download_error_quark.
func DownloadErrorQuark() glib.Quark {
	retC := C.webkit_download_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type FaviconDatabaseError C.WebKitFaviconDatabaseError

const (
	WEBKIT_FAVICON_DATABASE_ERROR_NOT_INITIALIZED   FaviconDatabaseError = 0
	WEBKIT_FAVICON_DATABASE_ERROR_FAVICON_NOT_FOUND FaviconDatabaseError = 1
	WEBKIT_FAVICON_DATABASE_ERROR_FAVICON_UNKNOWN   FaviconDatabaseError = 2
)

// FaviconDatabaseErrorQuark is a wrapper around the C function webkit_favicon_database_error_quark.
func FaviconDatabaseErrorQuark() glib.Quark {
	retC := C.webkit_favicon_database_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type InsecureContentEvent C.WebKitInsecureContentEvent

const (
	WEBKIT_INSECURE_CONTENT_RUN       InsecureContentEvent = 0
	WEBKIT_INSECURE_CONTENT_DISPLAYED InsecureContentEvent = 1
)

type JavascriptError C.WebKitJavascriptError

const (
	WEBKIT_JAVASCRIPT_ERROR_SCRIPT_FAILED JavascriptError = 699
)

// JavascriptErrorQuark is a wrapper around the C function webkit_javascript_error_quark.
func JavascriptErrorQuark() glib.Quark {
	retC := C.webkit_javascript_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type LoadEvent C.WebKitLoadEvent

const (
	WEBKIT_LOAD_STARTED    LoadEvent = 0
	WEBKIT_LOAD_REDIRECTED LoadEvent = 1
	WEBKIT_LOAD_COMMITTED  LoadEvent = 2
	WEBKIT_LOAD_FINISHED   LoadEvent = 3
)

type NavigationType C.WebKitNavigationType

const (
	WEBKIT_NAVIGATION_TYPE_LINK_CLICKED     NavigationType = 0
	WEBKIT_NAVIGATION_TYPE_FORM_SUBMITTED   NavigationType = 1
	WEBKIT_NAVIGATION_TYPE_BACK_FORWARD     NavigationType = 2
	WEBKIT_NAVIGATION_TYPE_RELOAD           NavigationType = 3
	WEBKIT_NAVIGATION_TYPE_FORM_RESUBMITTED NavigationType = 4
	WEBKIT_NAVIGATION_TYPE_OTHER            NavigationType = 5
)

type NetworkError C.WebKitNetworkError

const (
	WEBKIT_NETWORK_ERROR_FAILED              NetworkError = 399
	WEBKIT_NETWORK_ERROR_TRANSPORT           NetworkError = 300
	WEBKIT_NETWORK_ERROR_UNKNOWN_PROTOCOL    NetworkError = 301
	WEBKIT_NETWORK_ERROR_CANCELLED           NetworkError = 302
	WEBKIT_NETWORK_ERROR_FILE_DOES_NOT_EXIST NetworkError = 303
)

// NetworkErrorQuark is a wrapper around the C function webkit_network_error_quark.
func NetworkErrorQuark() glib.Quark {
	retC := C.webkit_network_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type PluginError C.WebKitPluginError

const (
	WEBKIT_PLUGIN_ERROR_FAILED               PluginError = 299
	WEBKIT_PLUGIN_ERROR_CANNOT_FIND_PLUGIN   PluginError = 200
	WEBKIT_PLUGIN_ERROR_CANNOT_LOAD_PLUGIN   PluginError = 201
	WEBKIT_PLUGIN_ERROR_JAVA_UNAVAILABLE     PluginError = 202
	WEBKIT_PLUGIN_ERROR_CONNECTION_CANCELLED PluginError = 203
	WEBKIT_PLUGIN_ERROR_WILL_HANDLE_LOAD     PluginError = 204
)

// PluginErrorQuark is a wrapper around the C function webkit_plugin_error_quark.
func PluginErrorQuark() glib.Quark {
	retC := C.webkit_plugin_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type PolicyDecisionType C.WebKitPolicyDecisionType

const (
	WEBKIT_POLICY_DECISION_TYPE_NAVIGATION_ACTION PolicyDecisionType = 0
	WEBKIT_POLICY_DECISION_TYPE_NEW_WINDOW_ACTION PolicyDecisionType = 1
	WEBKIT_POLICY_DECISION_TYPE_RESPONSE          PolicyDecisionType = 2
)

type PolicyError C.WebKitPolicyError

const (
	WEBKIT_POLICY_ERROR_FAILED                                  PolicyError = 199
	WEBKIT_POLICY_ERROR_CANNOT_SHOW_MIME_TYPE                   PolicyError = 100
	WEBKIT_POLICY_ERROR_CANNOT_SHOW_URI                         PolicyError = 101
	WEBKIT_POLICY_ERROR_FRAME_LOAD_INTERRUPTED_BY_POLICY_CHANGE PolicyError = 102
	WEBKIT_POLICY_ERROR_CANNOT_USE_RESTRICTED_PORT              PolicyError = 103
)

// PolicyErrorQuark is a wrapper around the C function webkit_policy_error_quark.
func PolicyErrorQuark() glib.Quark {
	retC := C.webkit_policy_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type PrintError C.WebKitPrintError

const (
	WEBKIT_PRINT_ERROR_GENERAL            PrintError = 599
	WEBKIT_PRINT_ERROR_PRINTER_NOT_FOUND  PrintError = 500
	WEBKIT_PRINT_ERROR_INVALID_PAGE_RANGE PrintError = 501
)

// PrintErrorQuark is a wrapper around the C function webkit_print_error_quark.
func PrintErrorQuark() glib.Quark {
	retC := C.webkit_print_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type PrintOperationResponse C.WebKitPrintOperationResponse

const (
	WEBKIT_PRINT_OPERATION_RESPONSE_PRINT  PrintOperationResponse = 0
	WEBKIT_PRINT_OPERATION_RESPONSE_CANCEL PrintOperationResponse = 1
)

type SaveMode C.WebKitSaveMode

const (
	WEBKIT_SAVE_MODE_MHTML SaveMode = 0
)

type ScriptDialogType C.WebKitScriptDialogType

const (
	WEBKIT_SCRIPT_DIALOG_ALERT                 ScriptDialogType = 0
	WEBKIT_SCRIPT_DIALOG_CONFIRM               ScriptDialogType = 1
	WEBKIT_SCRIPT_DIALOG_PROMPT                ScriptDialogType = 2
	WEBKIT_SCRIPT_DIALOG_BEFORE_UNLOAD_CONFIRM ScriptDialogType = 3
)

type SnapshotError C.WebKitSnapshotError

const (
	WEBKIT_SNAPSHOT_ERROR_FAILED_TO_CREATE SnapshotError = 799
)

// SnapshotErrorQuark is a wrapper around the C function webkit_snapshot_error_quark.
func SnapshotErrorQuark() glib.Quark {
	retC := C.webkit_snapshot_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type SnapshotRegion C.WebKitSnapshotRegion

const (
	WEBKIT_SNAPSHOT_REGION_VISIBLE       SnapshotRegion = 0
	WEBKIT_SNAPSHOT_REGION_FULL_DOCUMENT SnapshotRegion = 1
)

type TLSErrorsPolicy C.WebKitTLSErrorsPolicy

const (
	WEBKIT_TLS_ERRORS_POLICY_IGNORE TLSErrorsPolicy = 0
	WEBKIT_TLS_ERRORS_POLICY_FAIL   TLSErrorsPolicy = 1
)

// GetMajorVersion is a wrapper around the C function webkit_get_major_version.
func GetMajorVersion() uint32 {
	retC := C.webkit_get_major_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMicroVersion is a wrapper around the C function webkit_get_micro_version.
func GetMicroVersion() uint32 {
	retC := C.webkit_get_micro_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMinorVersion is a wrapper around the C function webkit_get_minor_version.
func GetMinorVersion() uint32 {
	retC := C.webkit_get_minor_version()
	retGo := (uint32)(retC)

	return retGo
}

// PermissionRequest is a wrapper around the C record WebKitPermissionRequest.
type PermissionRequest struct {
	native *C.WebKitPermissionRequest
}

func PermissionRequestNewFromC(u unsafe.Pointer) *PermissionRequest {
	c := (*C.WebKitPermissionRequest)(u)
	if c == nil {
		return nil
	}

	g := &PermissionRequest{native: c}

	return g
}

func (recv *PermissionRequest) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PermissionRequest with another PermissionRequest, and returns true if they represent the same GObject.
func (recv *PermissionRequest) Equals(other *PermissionRequest) bool {
	return other.ToC() == recv.ToC()
}

// Allow is a wrapper around the C function webkit_permission_request_allow.
func (recv *PermissionRequest) Allow() {
	C.webkit_permission_request_allow((*C.WebKitPermissionRequest)(recv.native))

	return
}

// Deny is a wrapper around the C function webkit_permission_request_deny.
func (recv *PermissionRequest) Deny() {
	C.webkit_permission_request_deny((*C.WebKitPermissionRequest)(recv.native))

	return
}

// ApplicationInfo is a wrapper around the C record WebKitApplicationInfo.
type ApplicationInfo struct {
	native *C.WebKitApplicationInfo
}

func ApplicationInfoNewFromC(u unsafe.Pointer) *ApplicationInfo {
	c := (*C.WebKitApplicationInfo)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationInfo{native: c}

	return g
}

func (recv *ApplicationInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ApplicationInfo with another ApplicationInfo, and returns true if they represent the same GObject.
func (recv *ApplicationInfo) Equals(other *ApplicationInfo) bool {
	return other.ToC() == recv.ToC()
}

// AuthenticationRequestClass is a wrapper around the C record WebKitAuthenticationRequestClass.
type AuthenticationRequestClass struct {
	native *C.WebKitAuthenticationRequestClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func AuthenticationRequestClassNewFromC(u unsafe.Pointer) *AuthenticationRequestClass {
	c := (*C.WebKitAuthenticationRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &AuthenticationRequestClass{native: c}

	return g
}

func (recv *AuthenticationRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthenticationRequestClass with another AuthenticationRequestClass, and returns true if they represent the same GObject.
func (recv *AuthenticationRequestClass) Equals(other *AuthenticationRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// AuthenticationRequestPrivate is a wrapper around the C record WebKitAuthenticationRequestPrivate.
type AuthenticationRequestPrivate struct {
	native *C.WebKitAuthenticationRequestPrivate
}

func AuthenticationRequestPrivateNewFromC(u unsafe.Pointer) *AuthenticationRequestPrivate {
	c := (*C.WebKitAuthenticationRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AuthenticationRequestPrivate{native: c}

	return g
}

func (recv *AuthenticationRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AuthenticationRequestPrivate with another AuthenticationRequestPrivate, and returns true if they represent the same GObject.
func (recv *AuthenticationRequestPrivate) Equals(other *AuthenticationRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// AutomationSessionClass is a wrapper around the C record WebKitAutomationSessionClass.
type AutomationSessionClass struct {
	native *C.WebKitAutomationSessionClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func AutomationSessionClassNewFromC(u unsafe.Pointer) *AutomationSessionClass {
	c := (*C.WebKitAutomationSessionClass)(u)
	if c == nil {
		return nil
	}

	g := &AutomationSessionClass{native: c}

	return g
}

func (recv *AutomationSessionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AutomationSessionClass with another AutomationSessionClass, and returns true if they represent the same GObject.
func (recv *AutomationSessionClass) Equals(other *AutomationSessionClass) bool {
	return other.ToC() == recv.ToC()
}

// AutomationSessionPrivate is a wrapper around the C record WebKitAutomationSessionPrivate.
type AutomationSessionPrivate struct {
	native *C.WebKitAutomationSessionPrivate
}

func AutomationSessionPrivateNewFromC(u unsafe.Pointer) *AutomationSessionPrivate {
	c := (*C.WebKitAutomationSessionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &AutomationSessionPrivate{native: c}

	return g
}

func (recv *AutomationSessionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AutomationSessionPrivate with another AutomationSessionPrivate, and returns true if they represent the same GObject.
func (recv *AutomationSessionPrivate) Equals(other *AutomationSessionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// BackForwardListClass is a wrapper around the C record WebKitBackForwardListClass.
type BackForwardListClass struct {
	native *C.WebKitBackForwardListClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func BackForwardListClassNewFromC(u unsafe.Pointer) *BackForwardListClass {
	c := (*C.WebKitBackForwardListClass)(u)
	if c == nil {
		return nil
	}

	g := &BackForwardListClass{native: c}

	return g
}

func (recv *BackForwardListClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BackForwardListClass with another BackForwardListClass, and returns true if they represent the same GObject.
func (recv *BackForwardListClass) Equals(other *BackForwardListClass) bool {
	return other.ToC() == recv.ToC()
}

// BackForwardListItemClass is a wrapper around the C record WebKitBackForwardListItemClass.
type BackForwardListItemClass struct {
	native *C.WebKitBackForwardListItemClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func BackForwardListItemClassNewFromC(u unsafe.Pointer) *BackForwardListItemClass {
	c := (*C.WebKitBackForwardListItemClass)(u)
	if c == nil {
		return nil
	}

	g := &BackForwardListItemClass{native: c}

	return g
}

func (recv *BackForwardListItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BackForwardListItemClass with another BackForwardListItemClass, and returns true if they represent the same GObject.
func (recv *BackForwardListItemClass) Equals(other *BackForwardListItemClass) bool {
	return other.ToC() == recv.ToC()
}

// BackForwardListItemPrivate is a wrapper around the C record WebKitBackForwardListItemPrivate.
type BackForwardListItemPrivate struct {
	native *C.WebKitBackForwardListItemPrivate
}

func BackForwardListItemPrivateNewFromC(u unsafe.Pointer) *BackForwardListItemPrivate {
	c := (*C.WebKitBackForwardListItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BackForwardListItemPrivate{native: c}

	return g
}

func (recv *BackForwardListItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BackForwardListItemPrivate with another BackForwardListItemPrivate, and returns true if they represent the same GObject.
func (recv *BackForwardListItemPrivate) Equals(other *BackForwardListItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// BackForwardListPrivate is a wrapper around the C record WebKitBackForwardListPrivate.
type BackForwardListPrivate struct {
	native *C.WebKitBackForwardListPrivate
}

func BackForwardListPrivateNewFromC(u unsafe.Pointer) *BackForwardListPrivate {
	c := (*C.WebKitBackForwardListPrivate)(u)
	if c == nil {
		return nil
	}

	g := &BackForwardListPrivate{native: c}

	return g
}

func (recv *BackForwardListPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this BackForwardListPrivate with another BackForwardListPrivate, and returns true if they represent the same GObject.
func (recv *BackForwardListPrivate) Equals(other *BackForwardListPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ColorChooserRequestClass is a wrapper around the C record WebKitColorChooserRequestClass.
type ColorChooserRequestClass struct {
	native *C.WebKitColorChooserRequestClass
	// parent_class : record
}

func ColorChooserRequestClassNewFromC(u unsafe.Pointer) *ColorChooserRequestClass {
	c := (*C.WebKitColorChooserRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserRequestClass{native: c}

	return g
}

func (recv *ColorChooserRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ColorChooserRequestClass with another ColorChooserRequestClass, and returns true if they represent the same GObject.
func (recv *ColorChooserRequestClass) Equals(other *ColorChooserRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// ColorChooserRequestPrivate is a wrapper around the C record WebKitColorChooserRequestPrivate.
type ColorChooserRequestPrivate struct {
	native *C.WebKitColorChooserRequestPrivate
}

func ColorChooserRequestPrivateNewFromC(u unsafe.Pointer) *ColorChooserRequestPrivate {
	c := (*C.WebKitColorChooserRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserRequestPrivate{native: c}

	return g
}

func (recv *ColorChooserRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ColorChooserRequestPrivate with another ColorChooserRequestPrivate, and returns true if they represent the same GObject.
func (recv *ColorChooserRequestPrivate) Equals(other *ColorChooserRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ContextMenuClass is a wrapper around the C record WebKitContextMenuClass.
type ContextMenuClass struct {
	native *C.WebKitContextMenuClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func ContextMenuClassNewFromC(u unsafe.Pointer) *ContextMenuClass {
	c := (*C.WebKitContextMenuClass)(u)
	if c == nil {
		return nil
	}

	g := &ContextMenuClass{native: c}

	return g
}

func (recv *ContextMenuClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContextMenuClass with another ContextMenuClass, and returns true if they represent the same GObject.
func (recv *ContextMenuClass) Equals(other *ContextMenuClass) bool {
	return other.ToC() == recv.ToC()
}

// ContextMenuItemClass is a wrapper around the C record WebKitContextMenuItemClass.
type ContextMenuItemClass struct {
	native *C.WebKitContextMenuItemClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func ContextMenuItemClassNewFromC(u unsafe.Pointer) *ContextMenuItemClass {
	c := (*C.WebKitContextMenuItemClass)(u)
	if c == nil {
		return nil
	}

	g := &ContextMenuItemClass{native: c}

	return g
}

func (recv *ContextMenuItemClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContextMenuItemClass with another ContextMenuItemClass, and returns true if they represent the same GObject.
func (recv *ContextMenuItemClass) Equals(other *ContextMenuItemClass) bool {
	return other.ToC() == recv.ToC()
}

// ContextMenuItemPrivate is a wrapper around the C record WebKitContextMenuItemPrivate.
type ContextMenuItemPrivate struct {
	native *C.WebKitContextMenuItemPrivate
}

func ContextMenuItemPrivateNewFromC(u unsafe.Pointer) *ContextMenuItemPrivate {
	c := (*C.WebKitContextMenuItemPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ContextMenuItemPrivate{native: c}

	return g
}

func (recv *ContextMenuItemPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContextMenuItemPrivate with another ContextMenuItemPrivate, and returns true if they represent the same GObject.
func (recv *ContextMenuItemPrivate) Equals(other *ContextMenuItemPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ContextMenuPrivate is a wrapper around the C record WebKitContextMenuPrivate.
type ContextMenuPrivate struct {
	native *C.WebKitContextMenuPrivate
}

func ContextMenuPrivateNewFromC(u unsafe.Pointer) *ContextMenuPrivate {
	c := (*C.WebKitContextMenuPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ContextMenuPrivate{native: c}

	return g
}

func (recv *ContextMenuPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ContextMenuPrivate with another ContextMenuPrivate, and returns true if they represent the same GObject.
func (recv *ContextMenuPrivate) Equals(other *ContextMenuPrivate) bool {
	return other.ToC() == recv.ToC()
}

// CookieManagerClass is a wrapper around the C record WebKitCookieManagerClass.
type CookieManagerClass struct {
	native *C.WebKitCookieManagerClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func CookieManagerClassNewFromC(u unsafe.Pointer) *CookieManagerClass {
	c := (*C.WebKitCookieManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &CookieManagerClass{native: c}

	return g
}

func (recv *CookieManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CookieManagerClass with another CookieManagerClass, and returns true if they represent the same GObject.
func (recv *CookieManagerClass) Equals(other *CookieManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// CookieManagerPrivate is a wrapper around the C record WebKitCookieManagerPrivate.
type CookieManagerPrivate struct {
	native *C.WebKitCookieManagerPrivate
}

func CookieManagerPrivateNewFromC(u unsafe.Pointer) *CookieManagerPrivate {
	c := (*C.WebKitCookieManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &CookieManagerPrivate{native: c}

	return g
}

func (recv *CookieManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this CookieManagerPrivate with another CookieManagerPrivate, and returns true if they represent the same GObject.
func (recv *CookieManagerPrivate) Equals(other *CookieManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// Credential is a wrapper around the C record WebKitCredential.
type Credential struct {
	native *C.WebKitCredential
}

func CredentialNewFromC(u unsafe.Pointer) *Credential {
	c := (*C.WebKitCredential)(u)
	if c == nil {
		return nil
	}

	g := &Credential{native: c}

	return g
}

func (recv *Credential) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Credential with another Credential, and returns true if they represent the same GObject.
func (recv *Credential) Equals(other *Credential) bool {
	return other.ToC() == recv.ToC()
}

// DeviceInfoPermissionRequestClass is a wrapper around the C record WebKitDeviceInfoPermissionRequestClass.
type DeviceInfoPermissionRequestClass struct {
	native *C.WebKitDeviceInfoPermissionRequestClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func DeviceInfoPermissionRequestClassNewFromC(u unsafe.Pointer) *DeviceInfoPermissionRequestClass {
	c := (*C.WebKitDeviceInfoPermissionRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &DeviceInfoPermissionRequestClass{native: c}

	return g
}

func (recv *DeviceInfoPermissionRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DeviceInfoPermissionRequestClass with another DeviceInfoPermissionRequestClass, and returns true if they represent the same GObject.
func (recv *DeviceInfoPermissionRequestClass) Equals(other *DeviceInfoPermissionRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// DeviceInfoPermissionRequestPrivate is a wrapper around the C record WebKitDeviceInfoPermissionRequestPrivate.
type DeviceInfoPermissionRequestPrivate struct {
	native *C.WebKitDeviceInfoPermissionRequestPrivate
}

func DeviceInfoPermissionRequestPrivateNewFromC(u unsafe.Pointer) *DeviceInfoPermissionRequestPrivate {
	c := (*C.WebKitDeviceInfoPermissionRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DeviceInfoPermissionRequestPrivate{native: c}

	return g
}

func (recv *DeviceInfoPermissionRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DeviceInfoPermissionRequestPrivate with another DeviceInfoPermissionRequestPrivate, and returns true if they represent the same GObject.
func (recv *DeviceInfoPermissionRequestPrivate) Equals(other *DeviceInfoPermissionRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// DownloadClass is a wrapper around the C record WebKitDownloadClass.
type DownloadClass struct {
	native *C.WebKitDownloadClass
	// parent_class : record
	// no type for decide_destination
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func DownloadClassNewFromC(u unsafe.Pointer) *DownloadClass {
	c := (*C.WebKitDownloadClass)(u)
	if c == nil {
		return nil
	}

	g := &DownloadClass{native: c}

	return g
}

func (recv *DownloadClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DownloadClass with another DownloadClass, and returns true if they represent the same GObject.
func (recv *DownloadClass) Equals(other *DownloadClass) bool {
	return other.ToC() == recv.ToC()
}

// DownloadPrivate is a wrapper around the C record WebKitDownloadPrivate.
type DownloadPrivate struct {
	native *C.WebKitDownloadPrivate
}

func DownloadPrivateNewFromC(u unsafe.Pointer) *DownloadPrivate {
	c := (*C.WebKitDownloadPrivate)(u)
	if c == nil {
		return nil
	}

	g := &DownloadPrivate{native: c}

	return g
}

func (recv *DownloadPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this DownloadPrivate with another DownloadPrivate, and returns true if they represent the same GObject.
func (recv *DownloadPrivate) Equals(other *DownloadPrivate) bool {
	return other.ToC() == recv.ToC()
}

// EditorStateClass is a wrapper around the C record WebKitEditorStateClass.
type EditorStateClass struct {
	native *C.WebKitEditorStateClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func EditorStateClassNewFromC(u unsafe.Pointer) *EditorStateClass {
	c := (*C.WebKitEditorStateClass)(u)
	if c == nil {
		return nil
	}

	g := &EditorStateClass{native: c}

	return g
}

func (recv *EditorStateClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EditorStateClass with another EditorStateClass, and returns true if they represent the same GObject.
func (recv *EditorStateClass) Equals(other *EditorStateClass) bool {
	return other.ToC() == recv.ToC()
}

// EditorStatePrivate is a wrapper around the C record WebKitEditorStatePrivate.
type EditorStatePrivate struct {
	native *C.WebKitEditorStatePrivate
}

func EditorStatePrivateNewFromC(u unsafe.Pointer) *EditorStatePrivate {
	c := (*C.WebKitEditorStatePrivate)(u)
	if c == nil {
		return nil
	}

	g := &EditorStatePrivate{native: c}

	return g
}

func (recv *EditorStatePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this EditorStatePrivate with another EditorStatePrivate, and returns true if they represent the same GObject.
func (recv *EditorStatePrivate) Equals(other *EditorStatePrivate) bool {
	return other.ToC() == recv.ToC()
}

// FaviconDatabaseClass is a wrapper around the C record WebKitFaviconDatabaseClass.
type FaviconDatabaseClass struct {
	native *C.WebKitFaviconDatabaseClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func FaviconDatabaseClassNewFromC(u unsafe.Pointer) *FaviconDatabaseClass {
	c := (*C.WebKitFaviconDatabaseClass)(u)
	if c == nil {
		return nil
	}

	g := &FaviconDatabaseClass{native: c}

	return g
}

func (recv *FaviconDatabaseClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FaviconDatabaseClass with another FaviconDatabaseClass, and returns true if they represent the same GObject.
func (recv *FaviconDatabaseClass) Equals(other *FaviconDatabaseClass) bool {
	return other.ToC() == recv.ToC()
}

// FaviconDatabasePrivate is a wrapper around the C record WebKitFaviconDatabasePrivate.
type FaviconDatabasePrivate struct {
	native *C.WebKitFaviconDatabasePrivate
}

func FaviconDatabasePrivateNewFromC(u unsafe.Pointer) *FaviconDatabasePrivate {
	c := (*C.WebKitFaviconDatabasePrivate)(u)
	if c == nil {
		return nil
	}

	g := &FaviconDatabasePrivate{native: c}

	return g
}

func (recv *FaviconDatabasePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FaviconDatabasePrivate with another FaviconDatabasePrivate, and returns true if they represent the same GObject.
func (recv *FaviconDatabasePrivate) Equals(other *FaviconDatabasePrivate) bool {
	return other.ToC() == recv.ToC()
}

// FileChooserRequestClass is a wrapper around the C record WebKitFileChooserRequestClass.
type FileChooserRequestClass struct {
	native *C.WebKitFileChooserRequestClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func FileChooserRequestClassNewFromC(u unsafe.Pointer) *FileChooserRequestClass {
	c := (*C.WebKitFileChooserRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserRequestClass{native: c}

	return g
}

func (recv *FileChooserRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileChooserRequestClass with another FileChooserRequestClass, and returns true if they represent the same GObject.
func (recv *FileChooserRequestClass) Equals(other *FileChooserRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// FileChooserRequestPrivate is a wrapper around the C record WebKitFileChooserRequestPrivate.
type FileChooserRequestPrivate struct {
	native *C.WebKitFileChooserRequestPrivate
}

func FileChooserRequestPrivateNewFromC(u unsafe.Pointer) *FileChooserRequestPrivate {
	c := (*C.WebKitFileChooserRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserRequestPrivate{native: c}

	return g
}

func (recv *FileChooserRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FileChooserRequestPrivate with another FileChooserRequestPrivate, and returns true if they represent the same GObject.
func (recv *FileChooserRequestPrivate) Equals(other *FileChooserRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FindControllerClass is a wrapper around the C record WebKitFindControllerClass.
type FindControllerClass struct {
	native *C.WebKitFindControllerClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func FindControllerClassNewFromC(u unsafe.Pointer) *FindControllerClass {
	c := (*C.WebKitFindControllerClass)(u)
	if c == nil {
		return nil
	}

	g := &FindControllerClass{native: c}

	return g
}

func (recv *FindControllerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FindControllerClass with another FindControllerClass, and returns true if they represent the same GObject.
func (recv *FindControllerClass) Equals(other *FindControllerClass) bool {
	return other.ToC() == recv.ToC()
}

// FindControllerPrivate is a wrapper around the C record WebKitFindControllerPrivate.
type FindControllerPrivate struct {
	native *C.WebKitFindControllerPrivate
}

func FindControllerPrivateNewFromC(u unsafe.Pointer) *FindControllerPrivate {
	c := (*C.WebKitFindControllerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FindControllerPrivate{native: c}

	return g
}

func (recv *FindControllerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FindControllerPrivate with another FindControllerPrivate, and returns true if they represent the same GObject.
func (recv *FindControllerPrivate) Equals(other *FindControllerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// FormSubmissionRequestClass is a wrapper around the C record WebKitFormSubmissionRequestClass.
type FormSubmissionRequestClass struct {
	native *C.WebKitFormSubmissionRequestClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func FormSubmissionRequestClassNewFromC(u unsafe.Pointer) *FormSubmissionRequestClass {
	c := (*C.WebKitFormSubmissionRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &FormSubmissionRequestClass{native: c}

	return g
}

func (recv *FormSubmissionRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FormSubmissionRequestClass with another FormSubmissionRequestClass, and returns true if they represent the same GObject.
func (recv *FormSubmissionRequestClass) Equals(other *FormSubmissionRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// FormSubmissionRequestPrivate is a wrapper around the C record WebKitFormSubmissionRequestPrivate.
type FormSubmissionRequestPrivate struct {
	native *C.WebKitFormSubmissionRequestPrivate
}

func FormSubmissionRequestPrivateNewFromC(u unsafe.Pointer) *FormSubmissionRequestPrivate {
	c := (*C.WebKitFormSubmissionRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &FormSubmissionRequestPrivate{native: c}

	return g
}

func (recv *FormSubmissionRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this FormSubmissionRequestPrivate with another FormSubmissionRequestPrivate, and returns true if they represent the same GObject.
func (recv *FormSubmissionRequestPrivate) Equals(other *FormSubmissionRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// GeolocationManagerClass is a wrapper around the C record WebKitGeolocationManagerClass.
type GeolocationManagerClass struct {
	native *C.WebKitGeolocationManagerClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func GeolocationManagerClassNewFromC(u unsafe.Pointer) *GeolocationManagerClass {
	c := (*C.WebKitGeolocationManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &GeolocationManagerClass{native: c}

	return g
}

func (recv *GeolocationManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GeolocationManagerClass with another GeolocationManagerClass, and returns true if they represent the same GObject.
func (recv *GeolocationManagerClass) Equals(other *GeolocationManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// GeolocationManagerPrivate is a wrapper around the C record WebKitGeolocationManagerPrivate.
type GeolocationManagerPrivate struct {
	native *C.WebKitGeolocationManagerPrivate
}

func GeolocationManagerPrivateNewFromC(u unsafe.Pointer) *GeolocationManagerPrivate {
	c := (*C.WebKitGeolocationManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &GeolocationManagerPrivate{native: c}

	return g
}

func (recv *GeolocationManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GeolocationManagerPrivate with another GeolocationManagerPrivate, and returns true if they represent the same GObject.
func (recv *GeolocationManagerPrivate) Equals(other *GeolocationManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// GeolocationPermissionRequestClass is a wrapper around the C record WebKitGeolocationPermissionRequestClass.
type GeolocationPermissionRequestClass struct {
	native *C.WebKitGeolocationPermissionRequestClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func GeolocationPermissionRequestClassNewFromC(u unsafe.Pointer) *GeolocationPermissionRequestClass {
	c := (*C.WebKitGeolocationPermissionRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &GeolocationPermissionRequestClass{native: c}

	return g
}

func (recv *GeolocationPermissionRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GeolocationPermissionRequestClass with another GeolocationPermissionRequestClass, and returns true if they represent the same GObject.
func (recv *GeolocationPermissionRequestClass) Equals(other *GeolocationPermissionRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// GeolocationPermissionRequestPrivate is a wrapper around the C record WebKitGeolocationPermissionRequestPrivate.
type GeolocationPermissionRequestPrivate struct {
	native *C.WebKitGeolocationPermissionRequestPrivate
}

func GeolocationPermissionRequestPrivateNewFromC(u unsafe.Pointer) *GeolocationPermissionRequestPrivate {
	c := (*C.WebKitGeolocationPermissionRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &GeolocationPermissionRequestPrivate{native: c}

	return g
}

func (recv *GeolocationPermissionRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this GeolocationPermissionRequestPrivate with another GeolocationPermissionRequestPrivate, and returns true if they represent the same GObject.
func (recv *GeolocationPermissionRequestPrivate) Equals(other *GeolocationPermissionRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// HitTestResultClass is a wrapper around the C record WebKitHitTestResultClass.
type HitTestResultClass struct {
	native *C.WebKitHitTestResultClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func HitTestResultClassNewFromC(u unsafe.Pointer) *HitTestResultClass {
	c := (*C.WebKitHitTestResultClass)(u)
	if c == nil {
		return nil
	}

	g := &HitTestResultClass{native: c}

	return g
}

func (recv *HitTestResultClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HitTestResultClass with another HitTestResultClass, and returns true if they represent the same GObject.
func (recv *HitTestResultClass) Equals(other *HitTestResultClass) bool {
	return other.ToC() == recv.ToC()
}

// HitTestResultPrivate is a wrapper around the C record WebKitHitTestResultPrivate.
type HitTestResultPrivate struct {
	native *C.WebKitHitTestResultPrivate
}

func HitTestResultPrivateNewFromC(u unsafe.Pointer) *HitTestResultPrivate {
	c := (*C.WebKitHitTestResultPrivate)(u)
	if c == nil {
		return nil
	}

	g := &HitTestResultPrivate{native: c}

	return g
}

func (recv *HitTestResultPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this HitTestResultPrivate with another HitTestResultPrivate, and returns true if they represent the same GObject.
func (recv *HitTestResultPrivate) Equals(other *HitTestResultPrivate) bool {
	return other.ToC() == recv.ToC()
}

// InstallMissingMediaPluginsPermissionRequestClass is a wrapper around the C record WebKitInstallMissingMediaPluginsPermissionRequestClass.
type InstallMissingMediaPluginsPermissionRequestClass struct {
	native *C.WebKitInstallMissingMediaPluginsPermissionRequestClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func InstallMissingMediaPluginsPermissionRequestClassNewFromC(u unsafe.Pointer) *InstallMissingMediaPluginsPermissionRequestClass {
	c := (*C.WebKitInstallMissingMediaPluginsPermissionRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &InstallMissingMediaPluginsPermissionRequestClass{native: c}

	return g
}

func (recv *InstallMissingMediaPluginsPermissionRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InstallMissingMediaPluginsPermissionRequestClass with another InstallMissingMediaPluginsPermissionRequestClass, and returns true if they represent the same GObject.
func (recv *InstallMissingMediaPluginsPermissionRequestClass) Equals(other *InstallMissingMediaPluginsPermissionRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// InstallMissingMediaPluginsPermissionRequestPrivate is a wrapper around the C record WebKitInstallMissingMediaPluginsPermissionRequestPrivate.
type InstallMissingMediaPluginsPermissionRequestPrivate struct {
	native *C.WebKitInstallMissingMediaPluginsPermissionRequestPrivate
}

func InstallMissingMediaPluginsPermissionRequestPrivateNewFromC(u unsafe.Pointer) *InstallMissingMediaPluginsPermissionRequestPrivate {
	c := (*C.WebKitInstallMissingMediaPluginsPermissionRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &InstallMissingMediaPluginsPermissionRequestPrivate{native: c}

	return g
}

func (recv *InstallMissingMediaPluginsPermissionRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InstallMissingMediaPluginsPermissionRequestPrivate with another InstallMissingMediaPluginsPermissionRequestPrivate, and returns true if they represent the same GObject.
func (recv *InstallMissingMediaPluginsPermissionRequestPrivate) Equals(other *InstallMissingMediaPluginsPermissionRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// JavascriptResult is a wrapper around the C record WebKitJavascriptResult.
type JavascriptResult struct {
	native *C.WebKitJavascriptResult
}

func JavascriptResultNewFromC(u unsafe.Pointer) *JavascriptResult {
	c := (*C.WebKitJavascriptResult)(u)
	if c == nil {
		return nil
	}

	g := &JavascriptResult{native: c}

	return g
}

func (recv *JavascriptResult) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this JavascriptResult with another JavascriptResult, and returns true if they represent the same GObject.
func (recv *JavascriptResult) Equals(other *JavascriptResult) bool {
	return other.ToC() == recv.ToC()
}

// Unsupported : webkit_javascript_result_get_global_context : no return generator

// Unsupported : webkit_javascript_result_get_value : no return generator

// Ref is a wrapper around the C function webkit_javascript_result_ref.
func (recv *JavascriptResult) Ref() *JavascriptResult {
	retC := C.webkit_javascript_result_ref((*C.WebKitJavascriptResult)(recv.native))
	retGo := JavascriptResultNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function webkit_javascript_result_unref.
func (recv *JavascriptResult) Unref() {
	C.webkit_javascript_result_unref((*C.WebKitJavascriptResult)(recv.native))

	return
}

// MimeInfo is a wrapper around the C record WebKitMimeInfo.
type MimeInfo struct {
	native *C.WebKitMimeInfo
}

func MimeInfoNewFromC(u unsafe.Pointer) *MimeInfo {
	c := (*C.WebKitMimeInfo)(u)
	if c == nil {
		return nil
	}

	g := &MimeInfo{native: c}

	return g
}

func (recv *MimeInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this MimeInfo with another MimeInfo, and returns true if they represent the same GObject.
func (recv *MimeInfo) Equals(other *MimeInfo) bool {
	return other.ToC() == recv.ToC()
}

// GetDescription is a wrapper around the C function webkit_mime_info_get_description.
func (recv *MimeInfo) GetDescription() string {
	retC := C.webkit_mime_info_get_description((*C.WebKitMimeInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetExtensions is a wrapper around the C function webkit_mime_info_get_extensions.
func (recv *MimeInfo) GetExtensions() []string {
	retC := C.webkit_mime_info_get_extensions((*C.WebKitMimeInfo)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetMimeType is a wrapper around the C function webkit_mime_info_get_mime_type.
func (recv *MimeInfo) GetMimeType() string {
	retC := C.webkit_mime_info_get_mime_type((*C.WebKitMimeInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Ref is a wrapper around the C function webkit_mime_info_ref.
func (recv *MimeInfo) Ref() *MimeInfo {
	retC := C.webkit_mime_info_ref((*C.WebKitMimeInfo)(recv.native))
	retGo := MimeInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function webkit_mime_info_unref.
func (recv *MimeInfo) Unref() {
	C.webkit_mime_info_unref((*C.WebKitMimeInfo)(recv.native))

	return
}

// NavigationAction is a wrapper around the C record WebKitNavigationAction.
type NavigationAction struct {
	native *C.WebKitNavigationAction
}

func NavigationActionNewFromC(u unsafe.Pointer) *NavigationAction {
	c := (*C.WebKitNavigationAction)(u)
	if c == nil {
		return nil
	}

	g := &NavigationAction{native: c}

	return g
}

func (recv *NavigationAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NavigationAction with another NavigationAction, and returns true if they represent the same GObject.
func (recv *NavigationAction) Equals(other *NavigationAction) bool {
	return other.ToC() == recv.ToC()
}

// NavigationPolicyDecisionClass is a wrapper around the C record WebKitNavigationPolicyDecisionClass.
type NavigationPolicyDecisionClass struct {
	native *C.WebKitNavigationPolicyDecisionClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func NavigationPolicyDecisionClassNewFromC(u unsafe.Pointer) *NavigationPolicyDecisionClass {
	c := (*C.WebKitNavigationPolicyDecisionClass)(u)
	if c == nil {
		return nil
	}

	g := &NavigationPolicyDecisionClass{native: c}

	return g
}

func (recv *NavigationPolicyDecisionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NavigationPolicyDecisionClass with another NavigationPolicyDecisionClass, and returns true if they represent the same GObject.
func (recv *NavigationPolicyDecisionClass) Equals(other *NavigationPolicyDecisionClass) bool {
	return other.ToC() == recv.ToC()
}

// NavigationPolicyDecisionPrivate is a wrapper around the C record WebKitNavigationPolicyDecisionPrivate.
type NavigationPolicyDecisionPrivate struct {
	native *C.WebKitNavigationPolicyDecisionPrivate
}

func NavigationPolicyDecisionPrivateNewFromC(u unsafe.Pointer) *NavigationPolicyDecisionPrivate {
	c := (*C.WebKitNavigationPolicyDecisionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &NavigationPolicyDecisionPrivate{native: c}

	return g
}

func (recv *NavigationPolicyDecisionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NavigationPolicyDecisionPrivate with another NavigationPolicyDecisionPrivate, and returns true if they represent the same GObject.
func (recv *NavigationPolicyDecisionPrivate) Equals(other *NavigationPolicyDecisionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// NetworkProxySettings is a wrapper around the C record WebKitNetworkProxySettings.
type NetworkProxySettings struct {
	native *C.WebKitNetworkProxySettings
}

func NetworkProxySettingsNewFromC(u unsafe.Pointer) *NetworkProxySettings {
	c := (*C.WebKitNetworkProxySettings)(u)
	if c == nil {
		return nil
	}

	g := &NetworkProxySettings{native: c}

	return g
}

func (recv *NetworkProxySettings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NetworkProxySettings with another NetworkProxySettings, and returns true if they represent the same GObject.
func (recv *NetworkProxySettings) Equals(other *NetworkProxySettings) bool {
	return other.ToC() == recv.ToC()
}

// NotificationClass is a wrapper around the C record WebKitNotificationClass.
type NotificationClass struct {
	native *C.WebKitNotificationClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
	// no type for _webkit_reserved4
	// no type for _webkit_reserved5
}

func NotificationClassNewFromC(u unsafe.Pointer) *NotificationClass {
	c := (*C.WebKitNotificationClass)(u)
	if c == nil {
		return nil
	}

	g := &NotificationClass{native: c}

	return g
}

func (recv *NotificationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NotificationClass with another NotificationClass, and returns true if they represent the same GObject.
func (recv *NotificationClass) Equals(other *NotificationClass) bool {
	return other.ToC() == recv.ToC()
}

// NotificationPermissionRequestClass is a wrapper around the C record WebKitNotificationPermissionRequestClass.
type NotificationPermissionRequestClass struct {
	native *C.WebKitNotificationPermissionRequestClass
	// parent_class : record
}

func NotificationPermissionRequestClassNewFromC(u unsafe.Pointer) *NotificationPermissionRequestClass {
	c := (*C.WebKitNotificationPermissionRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &NotificationPermissionRequestClass{native: c}

	return g
}

func (recv *NotificationPermissionRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NotificationPermissionRequestClass with another NotificationPermissionRequestClass, and returns true if they represent the same GObject.
func (recv *NotificationPermissionRequestClass) Equals(other *NotificationPermissionRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// NotificationPermissionRequestPrivate is a wrapper around the C record WebKitNotificationPermissionRequestPrivate.
type NotificationPermissionRequestPrivate struct {
	native *C.WebKitNotificationPermissionRequestPrivate
}

func NotificationPermissionRequestPrivateNewFromC(u unsafe.Pointer) *NotificationPermissionRequestPrivate {
	c := (*C.WebKitNotificationPermissionRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &NotificationPermissionRequestPrivate{native: c}

	return g
}

func (recv *NotificationPermissionRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NotificationPermissionRequestPrivate with another NotificationPermissionRequestPrivate, and returns true if they represent the same GObject.
func (recv *NotificationPermissionRequestPrivate) Equals(other *NotificationPermissionRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// NotificationPrivate is a wrapper around the C record WebKitNotificationPrivate.
type NotificationPrivate struct {
	native *C.WebKitNotificationPrivate
}

func NotificationPrivateNewFromC(u unsafe.Pointer) *NotificationPrivate {
	c := (*C.WebKitNotificationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &NotificationPrivate{native: c}

	return g
}

func (recv *NotificationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this NotificationPrivate with another NotificationPrivate, and returns true if they represent the same GObject.
func (recv *NotificationPrivate) Equals(other *NotificationPrivate) bool {
	return other.ToC() == recv.ToC()
}

// OptionMenuClass is a wrapper around the C record WebKitOptionMenuClass.
type OptionMenuClass struct {
	native *C.WebKitOptionMenuClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func OptionMenuClassNewFromC(u unsafe.Pointer) *OptionMenuClass {
	c := (*C.WebKitOptionMenuClass)(u)
	if c == nil {
		return nil
	}

	g := &OptionMenuClass{native: c}

	return g
}

func (recv *OptionMenuClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionMenuClass with another OptionMenuClass, and returns true if they represent the same GObject.
func (recv *OptionMenuClass) Equals(other *OptionMenuClass) bool {
	return other.ToC() == recv.ToC()
}

// OptionMenuItem is a wrapper around the C record WebKitOptionMenuItem.
type OptionMenuItem struct {
	native *C.WebKitOptionMenuItem
}

func OptionMenuItemNewFromC(u unsafe.Pointer) *OptionMenuItem {
	c := (*C.WebKitOptionMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &OptionMenuItem{native: c}

	return g
}

func (recv *OptionMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionMenuItem with another OptionMenuItem, and returns true if they represent the same GObject.
func (recv *OptionMenuItem) Equals(other *OptionMenuItem) bool {
	return other.ToC() == recv.ToC()
}

// OptionMenuPrivate is a wrapper around the C record WebKitOptionMenuPrivate.
type OptionMenuPrivate struct {
	native *C.WebKitOptionMenuPrivate
}

func OptionMenuPrivateNewFromC(u unsafe.Pointer) *OptionMenuPrivate {
	c := (*C.WebKitOptionMenuPrivate)(u)
	if c == nil {
		return nil
	}

	g := &OptionMenuPrivate{native: c}

	return g
}

func (recv *OptionMenuPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OptionMenuPrivate with another OptionMenuPrivate, and returns true if they represent the same GObject.
func (recv *OptionMenuPrivate) Equals(other *OptionMenuPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PermissionRequestIface is a wrapper around the C record WebKitPermissionRequestIface.
type PermissionRequestIface struct {
	native *C.WebKitPermissionRequestIface
	// parent_interface : record
	// no type for allow
	// no type for deny
}

func PermissionRequestIfaceNewFromC(u unsafe.Pointer) *PermissionRequestIface {
	c := (*C.WebKitPermissionRequestIface)(u)
	if c == nil {
		return nil
	}

	g := &PermissionRequestIface{native: c}

	return g
}

func (recv *PermissionRequestIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PermissionRequestIface with another PermissionRequestIface, and returns true if they represent the same GObject.
func (recv *PermissionRequestIface) Equals(other *PermissionRequestIface) bool {
	return other.ToC() == recv.ToC()
}

// PluginClass is a wrapper around the C record WebKitPluginClass.
type PluginClass struct {
	native *C.WebKitPluginClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func PluginClassNewFromC(u unsafe.Pointer) *PluginClass {
	c := (*C.WebKitPluginClass)(u)
	if c == nil {
		return nil
	}

	g := &PluginClass{native: c}

	return g
}

func (recv *PluginClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PluginClass with another PluginClass, and returns true if they represent the same GObject.
func (recv *PluginClass) Equals(other *PluginClass) bool {
	return other.ToC() == recv.ToC()
}

// PluginPrivate is a wrapper around the C record WebKitPluginPrivate.
type PluginPrivate struct {
	native *C.WebKitPluginPrivate
}

func PluginPrivateNewFromC(u unsafe.Pointer) *PluginPrivate {
	c := (*C.WebKitPluginPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PluginPrivate{native: c}

	return g
}

func (recv *PluginPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PluginPrivate with another PluginPrivate, and returns true if they represent the same GObject.
func (recv *PluginPrivate) Equals(other *PluginPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PolicyDecisionClass is a wrapper around the C record WebKitPolicyDecisionClass.
type PolicyDecisionClass struct {
	native *C.WebKitPolicyDecisionClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func PolicyDecisionClassNewFromC(u unsafe.Pointer) *PolicyDecisionClass {
	c := (*C.WebKitPolicyDecisionClass)(u)
	if c == nil {
		return nil
	}

	g := &PolicyDecisionClass{native: c}

	return g
}

func (recv *PolicyDecisionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PolicyDecisionClass with another PolicyDecisionClass, and returns true if they represent the same GObject.
func (recv *PolicyDecisionClass) Equals(other *PolicyDecisionClass) bool {
	return other.ToC() == recv.ToC()
}

// PolicyDecisionPrivate is a wrapper around the C record WebKitPolicyDecisionPrivate.
type PolicyDecisionPrivate struct {
	native *C.WebKitPolicyDecisionPrivate
}

func PolicyDecisionPrivateNewFromC(u unsafe.Pointer) *PolicyDecisionPrivate {
	c := (*C.WebKitPolicyDecisionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PolicyDecisionPrivate{native: c}

	return g
}

func (recv *PolicyDecisionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PolicyDecisionPrivate with another PolicyDecisionPrivate, and returns true if they represent the same GObject.
func (recv *PolicyDecisionPrivate) Equals(other *PolicyDecisionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PrintCustomWidgetClass is a wrapper around the C record WebKitPrintCustomWidgetClass.
type PrintCustomWidgetClass struct {
	native *C.WebKitPrintCustomWidgetClass
	// parent_class : record
	// no type for apply
	// no type for update
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func PrintCustomWidgetClassNewFromC(u unsafe.Pointer) *PrintCustomWidgetClass {
	c := (*C.WebKitPrintCustomWidgetClass)(u)
	if c == nil {
		return nil
	}

	g := &PrintCustomWidgetClass{native: c}

	return g
}

func (recv *PrintCustomWidgetClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PrintCustomWidgetClass with another PrintCustomWidgetClass, and returns true if they represent the same GObject.
func (recv *PrintCustomWidgetClass) Equals(other *PrintCustomWidgetClass) bool {
	return other.ToC() == recv.ToC()
}

// PrintCustomWidgetPrivate is a wrapper around the C record WebKitPrintCustomWidgetPrivate.
type PrintCustomWidgetPrivate struct {
	native *C.WebKitPrintCustomWidgetPrivate
}

func PrintCustomWidgetPrivateNewFromC(u unsafe.Pointer) *PrintCustomWidgetPrivate {
	c := (*C.WebKitPrintCustomWidgetPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PrintCustomWidgetPrivate{native: c}

	return g
}

func (recv *PrintCustomWidgetPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PrintCustomWidgetPrivate with another PrintCustomWidgetPrivate, and returns true if they represent the same GObject.
func (recv *PrintCustomWidgetPrivate) Equals(other *PrintCustomWidgetPrivate) bool {
	return other.ToC() == recv.ToC()
}

// PrintOperationClass is a wrapper around the C record WebKitPrintOperationClass.
type PrintOperationClass struct {
	native *C.WebKitPrintOperationClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func PrintOperationClassNewFromC(u unsafe.Pointer) *PrintOperationClass {
	c := (*C.WebKitPrintOperationClass)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperationClass{native: c}

	return g
}

func (recv *PrintOperationClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PrintOperationClass with another PrintOperationClass, and returns true if they represent the same GObject.
func (recv *PrintOperationClass) Equals(other *PrintOperationClass) bool {
	return other.ToC() == recv.ToC()
}

// PrintOperationPrivate is a wrapper around the C record WebKitPrintOperationPrivate.
type PrintOperationPrivate struct {
	native *C.WebKitPrintOperationPrivate
}

func PrintOperationPrivateNewFromC(u unsafe.Pointer) *PrintOperationPrivate {
	c := (*C.WebKitPrintOperationPrivate)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperationPrivate{native: c}

	return g
}

func (recv *PrintOperationPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this PrintOperationPrivate with another PrintOperationPrivate, and returns true if they represent the same GObject.
func (recv *PrintOperationPrivate) Equals(other *PrintOperationPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ResponsePolicyDecisionClass is a wrapper around the C record WebKitResponsePolicyDecisionClass.
type ResponsePolicyDecisionClass struct {
	native *C.WebKitResponsePolicyDecisionClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func ResponsePolicyDecisionClassNewFromC(u unsafe.Pointer) *ResponsePolicyDecisionClass {
	c := (*C.WebKitResponsePolicyDecisionClass)(u)
	if c == nil {
		return nil
	}

	g := &ResponsePolicyDecisionClass{native: c}

	return g
}

func (recv *ResponsePolicyDecisionClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ResponsePolicyDecisionClass with another ResponsePolicyDecisionClass, and returns true if they represent the same GObject.
func (recv *ResponsePolicyDecisionClass) Equals(other *ResponsePolicyDecisionClass) bool {
	return other.ToC() == recv.ToC()
}

// ResponsePolicyDecisionPrivate is a wrapper around the C record WebKitResponsePolicyDecisionPrivate.
type ResponsePolicyDecisionPrivate struct {
	native *C.WebKitResponsePolicyDecisionPrivate
}

func ResponsePolicyDecisionPrivateNewFromC(u unsafe.Pointer) *ResponsePolicyDecisionPrivate {
	c := (*C.WebKitResponsePolicyDecisionPrivate)(u)
	if c == nil {
		return nil
	}

	g := &ResponsePolicyDecisionPrivate{native: c}

	return g
}

func (recv *ResponsePolicyDecisionPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ResponsePolicyDecisionPrivate with another ResponsePolicyDecisionPrivate, and returns true if they represent the same GObject.
func (recv *ResponsePolicyDecisionPrivate) Equals(other *ResponsePolicyDecisionPrivate) bool {
	return other.ToC() == recv.ToC()
}

// ScriptDialog is a wrapper around the C record WebKitScriptDialog.
type ScriptDialog struct {
	native *C.WebKitScriptDialog
}

func ScriptDialogNewFromC(u unsafe.Pointer) *ScriptDialog {
	c := (*C.WebKitScriptDialog)(u)
	if c == nil {
		return nil
	}

	g := &ScriptDialog{native: c}

	return g
}

func (recv *ScriptDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ScriptDialog with another ScriptDialog, and returns true if they represent the same GObject.
func (recv *ScriptDialog) Equals(other *ScriptDialog) bool {
	return other.ToC() == recv.ToC()
}

// ConfirmSetConfirmed is a wrapper around the C function webkit_script_dialog_confirm_set_confirmed.
func (recv *ScriptDialog) ConfirmSetConfirmed(confirmed bool) {
	c_confirmed :=
		boolToGboolean(confirmed)

	C.webkit_script_dialog_confirm_set_confirmed((*C.WebKitScriptDialog)(recv.native), c_confirmed)

	return
}

// GetDialogType is a wrapper around the C function webkit_script_dialog_get_dialog_type.
func (recv *ScriptDialog) GetDialogType() ScriptDialogType {
	retC := C.webkit_script_dialog_get_dialog_type((*C.WebKitScriptDialog)(recv.native))
	retGo := (ScriptDialogType)(retC)

	return retGo
}

// GetMessage is a wrapper around the C function webkit_script_dialog_get_message.
func (recv *ScriptDialog) GetMessage() string {
	retC := C.webkit_script_dialog_get_message((*C.WebKitScriptDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// PromptGetDefaultText is a wrapper around the C function webkit_script_dialog_prompt_get_default_text.
func (recv *ScriptDialog) PromptGetDefaultText() string {
	retC := C.webkit_script_dialog_prompt_get_default_text((*C.WebKitScriptDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// PromptSetText is a wrapper around the C function webkit_script_dialog_prompt_set_text.
func (recv *ScriptDialog) PromptSetText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.webkit_script_dialog_prompt_set_text((*C.WebKitScriptDialog)(recv.native), c_text)

	return
}

// SecurityManagerClass is a wrapper around the C record WebKitSecurityManagerClass.
type SecurityManagerClass struct {
	native *C.WebKitSecurityManagerClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func SecurityManagerClassNewFromC(u unsafe.Pointer) *SecurityManagerClass {
	c := (*C.WebKitSecurityManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &SecurityManagerClass{native: c}

	return g
}

func (recv *SecurityManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SecurityManagerClass with another SecurityManagerClass, and returns true if they represent the same GObject.
func (recv *SecurityManagerClass) Equals(other *SecurityManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// SecurityManagerPrivate is a wrapper around the C record WebKitSecurityManagerPrivate.
type SecurityManagerPrivate struct {
	native *C.WebKitSecurityManagerPrivate
}

func SecurityManagerPrivateNewFromC(u unsafe.Pointer) *SecurityManagerPrivate {
	c := (*C.WebKitSecurityManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SecurityManagerPrivate{native: c}

	return g
}

func (recv *SecurityManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SecurityManagerPrivate with another SecurityManagerPrivate, and returns true if they represent the same GObject.
func (recv *SecurityManagerPrivate) Equals(other *SecurityManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// SecurityOrigin is a wrapper around the C record WebKitSecurityOrigin.
type SecurityOrigin struct {
	native *C.WebKitSecurityOrigin
}

func SecurityOriginNewFromC(u unsafe.Pointer) *SecurityOrigin {
	c := (*C.WebKitSecurityOrigin)(u)
	if c == nil {
		return nil
	}

	g := &SecurityOrigin{native: c}

	return g
}

func (recv *SecurityOrigin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SecurityOrigin with another SecurityOrigin, and returns true if they represent the same GObject.
func (recv *SecurityOrigin) Equals(other *SecurityOrigin) bool {
	return other.ToC() == recv.ToC()
}

// SettingsClass is a wrapper around the C record WebKitSettingsClass.
type SettingsClass struct {
	native *C.WebKitSettingsClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func SettingsClassNewFromC(u unsafe.Pointer) *SettingsClass {
	c := (*C.WebKitSettingsClass)(u)
	if c == nil {
		return nil
	}

	g := &SettingsClass{native: c}

	return g
}

func (recv *SettingsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SettingsClass with another SettingsClass, and returns true if they represent the same GObject.
func (recv *SettingsClass) Equals(other *SettingsClass) bool {
	return other.ToC() == recv.ToC()
}

// SettingsPrivate is a wrapper around the C record WebKitSettingsPrivate.
type SettingsPrivate struct {
	native *C.WebKitSettingsPrivate
}

func SettingsPrivateNewFromC(u unsafe.Pointer) *SettingsPrivate {
	c := (*C.WebKitSettingsPrivate)(u)
	if c == nil {
		return nil
	}

	g := &SettingsPrivate{native: c}

	return g
}

func (recv *SettingsPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SettingsPrivate with another SettingsPrivate, and returns true if they represent the same GObject.
func (recv *SettingsPrivate) Equals(other *SettingsPrivate) bool {
	return other.ToC() == recv.ToC()
}

// URIRequestClass is a wrapper around the C record WebKitURIRequestClass.
type URIRequestClass struct {
	native *C.WebKitURIRequestClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func URIRequestClassNewFromC(u unsafe.Pointer) *URIRequestClass {
	c := (*C.WebKitURIRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &URIRequestClass{native: c}

	return g
}

func (recv *URIRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URIRequestClass with another URIRequestClass, and returns true if they represent the same GObject.
func (recv *URIRequestClass) Equals(other *URIRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// URIRequestPrivate is a wrapper around the C record WebKitURIRequestPrivate.
type URIRequestPrivate struct {
	native *C.WebKitURIRequestPrivate
}

func URIRequestPrivateNewFromC(u unsafe.Pointer) *URIRequestPrivate {
	c := (*C.WebKitURIRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &URIRequestPrivate{native: c}

	return g
}

func (recv *URIRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URIRequestPrivate with another URIRequestPrivate, and returns true if they represent the same GObject.
func (recv *URIRequestPrivate) Equals(other *URIRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// URIResponseClass is a wrapper around the C record WebKitURIResponseClass.
type URIResponseClass struct {
	native *C.WebKitURIResponseClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func URIResponseClassNewFromC(u unsafe.Pointer) *URIResponseClass {
	c := (*C.WebKitURIResponseClass)(u)
	if c == nil {
		return nil
	}

	g := &URIResponseClass{native: c}

	return g
}

func (recv *URIResponseClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URIResponseClass with another URIResponseClass, and returns true if they represent the same GObject.
func (recv *URIResponseClass) Equals(other *URIResponseClass) bool {
	return other.ToC() == recv.ToC()
}

// URIResponsePrivate is a wrapper around the C record WebKitURIResponsePrivate.
type URIResponsePrivate struct {
	native *C.WebKitURIResponsePrivate
}

func URIResponsePrivateNewFromC(u unsafe.Pointer) *URIResponsePrivate {
	c := (*C.WebKitURIResponsePrivate)(u)
	if c == nil {
		return nil
	}

	g := &URIResponsePrivate{native: c}

	return g
}

func (recv *URIResponsePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URIResponsePrivate with another URIResponsePrivate, and returns true if they represent the same GObject.
func (recv *URIResponsePrivate) Equals(other *URIResponsePrivate) bool {
	return other.ToC() == recv.ToC()
}

// URISchemeRequestClass is a wrapper around the C record WebKitURISchemeRequestClass.
type URISchemeRequestClass struct {
	native *C.WebKitURISchemeRequestClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func URISchemeRequestClassNewFromC(u unsafe.Pointer) *URISchemeRequestClass {
	c := (*C.WebKitURISchemeRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &URISchemeRequestClass{native: c}

	return g
}

func (recv *URISchemeRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URISchemeRequestClass with another URISchemeRequestClass, and returns true if they represent the same GObject.
func (recv *URISchemeRequestClass) Equals(other *URISchemeRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// URISchemeRequestPrivate is a wrapper around the C record WebKitURISchemeRequestPrivate.
type URISchemeRequestPrivate struct {
	native *C.WebKitURISchemeRequestPrivate
}

func URISchemeRequestPrivateNewFromC(u unsafe.Pointer) *URISchemeRequestPrivate {
	c := (*C.WebKitURISchemeRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &URISchemeRequestPrivate{native: c}

	return g
}

func (recv *URISchemeRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this URISchemeRequestPrivate with another URISchemeRequestPrivate, and returns true if they represent the same GObject.
func (recv *URISchemeRequestPrivate) Equals(other *URISchemeRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// UserContentFilter is a wrapper around the C record WebKitUserContentFilter.
type UserContentFilter struct {
	native *C.WebKitUserContentFilter
}

func UserContentFilterNewFromC(u unsafe.Pointer) *UserContentFilter {
	c := (*C.WebKitUserContentFilter)(u)
	if c == nil {
		return nil
	}

	g := &UserContentFilter{native: c}

	return g
}

func (recv *UserContentFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserContentFilter with another UserContentFilter, and returns true if they represent the same GObject.
func (recv *UserContentFilter) Equals(other *UserContentFilter) bool {
	return other.ToC() == recv.ToC()
}

// UserContentFilterStoreClass is a wrapper around the C record WebKitUserContentFilterStoreClass.
type UserContentFilterStoreClass struct {
	native *C.WebKitUserContentFilterStoreClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func UserContentFilterStoreClassNewFromC(u unsafe.Pointer) *UserContentFilterStoreClass {
	c := (*C.WebKitUserContentFilterStoreClass)(u)
	if c == nil {
		return nil
	}

	g := &UserContentFilterStoreClass{native: c}

	return g
}

func (recv *UserContentFilterStoreClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserContentFilterStoreClass with another UserContentFilterStoreClass, and returns true if they represent the same GObject.
func (recv *UserContentFilterStoreClass) Equals(other *UserContentFilterStoreClass) bool {
	return other.ToC() == recv.ToC()
}

// UserContentFilterStorePrivate is a wrapper around the C record WebKitUserContentFilterStorePrivate.
type UserContentFilterStorePrivate struct {
	native *C.WebKitUserContentFilterStorePrivate
}

func UserContentFilterStorePrivateNewFromC(u unsafe.Pointer) *UserContentFilterStorePrivate {
	c := (*C.WebKitUserContentFilterStorePrivate)(u)
	if c == nil {
		return nil
	}

	g := &UserContentFilterStorePrivate{native: c}

	return g
}

func (recv *UserContentFilterStorePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserContentFilterStorePrivate with another UserContentFilterStorePrivate, and returns true if they represent the same GObject.
func (recv *UserContentFilterStorePrivate) Equals(other *UserContentFilterStorePrivate) bool {
	return other.ToC() == recv.ToC()
}

// UserContentManagerClass is a wrapper around the C record WebKitUserContentManagerClass.
type UserContentManagerClass struct {
	native *C.WebKitUserContentManagerClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func UserContentManagerClassNewFromC(u unsafe.Pointer) *UserContentManagerClass {
	c := (*C.WebKitUserContentManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &UserContentManagerClass{native: c}

	return g
}

func (recv *UserContentManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserContentManagerClass with another UserContentManagerClass, and returns true if they represent the same GObject.
func (recv *UserContentManagerClass) Equals(other *UserContentManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// UserContentManagerPrivate is a wrapper around the C record WebKitUserContentManagerPrivate.
type UserContentManagerPrivate struct {
	native *C.WebKitUserContentManagerPrivate
}

func UserContentManagerPrivateNewFromC(u unsafe.Pointer) *UserContentManagerPrivate {
	c := (*C.WebKitUserContentManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UserContentManagerPrivate{native: c}

	return g
}

func (recv *UserContentManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserContentManagerPrivate with another UserContentManagerPrivate, and returns true if they represent the same GObject.
func (recv *UserContentManagerPrivate) Equals(other *UserContentManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// UserMediaPermissionRequestClass is a wrapper around the C record WebKitUserMediaPermissionRequestClass.
type UserMediaPermissionRequestClass struct {
	native *C.WebKitUserMediaPermissionRequestClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func UserMediaPermissionRequestClassNewFromC(u unsafe.Pointer) *UserMediaPermissionRequestClass {
	c := (*C.WebKitUserMediaPermissionRequestClass)(u)
	if c == nil {
		return nil
	}

	g := &UserMediaPermissionRequestClass{native: c}

	return g
}

func (recv *UserMediaPermissionRequestClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserMediaPermissionRequestClass with another UserMediaPermissionRequestClass, and returns true if they represent the same GObject.
func (recv *UserMediaPermissionRequestClass) Equals(other *UserMediaPermissionRequestClass) bool {
	return other.ToC() == recv.ToC()
}

// UserMediaPermissionRequestPrivate is a wrapper around the C record WebKitUserMediaPermissionRequestPrivate.
type UserMediaPermissionRequestPrivate struct {
	native *C.WebKitUserMediaPermissionRequestPrivate
}

func UserMediaPermissionRequestPrivateNewFromC(u unsafe.Pointer) *UserMediaPermissionRequestPrivate {
	c := (*C.WebKitUserMediaPermissionRequestPrivate)(u)
	if c == nil {
		return nil
	}

	g := &UserMediaPermissionRequestPrivate{native: c}

	return g
}

func (recv *UserMediaPermissionRequestPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserMediaPermissionRequestPrivate with another UserMediaPermissionRequestPrivate, and returns true if they represent the same GObject.
func (recv *UserMediaPermissionRequestPrivate) Equals(other *UserMediaPermissionRequestPrivate) bool {
	return other.ToC() == recv.ToC()
}

// UserScript is a wrapper around the C record WebKitUserScript.
type UserScript struct {
	native *C.WebKitUserScript
}

func UserScriptNewFromC(u unsafe.Pointer) *UserScript {
	c := (*C.WebKitUserScript)(u)
	if c == nil {
		return nil
	}

	g := &UserScript{native: c}

	return g
}

func (recv *UserScript) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserScript with another UserScript, and returns true if they represent the same GObject.
func (recv *UserScript) Equals(other *UserScript) bool {
	return other.ToC() == recv.ToC()
}

// UserStyleSheet is a wrapper around the C record WebKitUserStyleSheet.
type UserStyleSheet struct {
	native *C.WebKitUserStyleSheet
}

func UserStyleSheetNewFromC(u unsafe.Pointer) *UserStyleSheet {
	c := (*C.WebKitUserStyleSheet)(u)
	if c == nil {
		return nil
	}

	g := &UserStyleSheet{native: c}

	return g
}

func (recv *UserStyleSheet) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this UserStyleSheet with another UserStyleSheet, and returns true if they represent the same GObject.
func (recv *UserStyleSheet) Equals(other *UserStyleSheet) bool {
	return other.ToC() == recv.ToC()
}

// WebContextClass is a wrapper around the C record WebKitWebContextClass.
type WebContextClass struct {
	native *C.WebKitWebContextClass
	// parent : record
	// no type for download_started
	// no type for initialize_web_extensions
	// no type for initialize_notification_permissions
	// no type for automation_started
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func WebContextClassNewFromC(u unsafe.Pointer) *WebContextClass {
	c := (*C.WebKitWebContextClass)(u)
	if c == nil {
		return nil
	}

	g := &WebContextClass{native: c}

	return g
}

func (recv *WebContextClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebContextClass with another WebContextClass, and returns true if they represent the same GObject.
func (recv *WebContextClass) Equals(other *WebContextClass) bool {
	return other.ToC() == recv.ToC()
}

// WebContextPrivate is a wrapper around the C record WebKitWebContextPrivate.
type WebContextPrivate struct {
	native *C.WebKitWebContextPrivate
}

func WebContextPrivateNewFromC(u unsafe.Pointer) *WebContextPrivate {
	c := (*C.WebKitWebContextPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WebContextPrivate{native: c}

	return g
}

func (recv *WebContextPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebContextPrivate with another WebContextPrivate, and returns true if they represent the same GObject.
func (recv *WebContextPrivate) Equals(other *WebContextPrivate) bool {
	return other.ToC() == recv.ToC()
}

// WebInspectorClass is a wrapper around the C record WebKitWebInspectorClass.
type WebInspectorClass struct {
	native *C.WebKitWebInspectorClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func WebInspectorClassNewFromC(u unsafe.Pointer) *WebInspectorClass {
	c := (*C.WebKitWebInspectorClass)(u)
	if c == nil {
		return nil
	}

	g := &WebInspectorClass{native: c}

	return g
}

func (recv *WebInspectorClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebInspectorClass with another WebInspectorClass, and returns true if they represent the same GObject.
func (recv *WebInspectorClass) Equals(other *WebInspectorClass) bool {
	return other.ToC() == recv.ToC()
}

// WebInspectorPrivate is a wrapper around the C record WebKitWebInspectorPrivate.
type WebInspectorPrivate struct {
	native *C.WebKitWebInspectorPrivate
}

func WebInspectorPrivateNewFromC(u unsafe.Pointer) *WebInspectorPrivate {
	c := (*C.WebKitWebInspectorPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WebInspectorPrivate{native: c}

	return g
}

func (recv *WebInspectorPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebInspectorPrivate with another WebInspectorPrivate, and returns true if they represent the same GObject.
func (recv *WebInspectorPrivate) Equals(other *WebInspectorPrivate) bool {
	return other.ToC() == recv.ToC()
}

// WebResourceClass is a wrapper around the C record WebKitWebResourceClass.
type WebResourceClass struct {
	native *C.WebKitWebResourceClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func WebResourceClassNewFromC(u unsafe.Pointer) *WebResourceClass {
	c := (*C.WebKitWebResourceClass)(u)
	if c == nil {
		return nil
	}

	g := &WebResourceClass{native: c}

	return g
}

func (recv *WebResourceClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebResourceClass with another WebResourceClass, and returns true if they represent the same GObject.
func (recv *WebResourceClass) Equals(other *WebResourceClass) bool {
	return other.ToC() == recv.ToC()
}

// WebResourcePrivate is a wrapper around the C record WebKitWebResourcePrivate.
type WebResourcePrivate struct {
	native *C.WebKitWebResourcePrivate
}

func WebResourcePrivateNewFromC(u unsafe.Pointer) *WebResourcePrivate {
	c := (*C.WebKitWebResourcePrivate)(u)
	if c == nil {
		return nil
	}

	g := &WebResourcePrivate{native: c}

	return g
}

func (recv *WebResourcePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebResourcePrivate with another WebResourcePrivate, and returns true if they represent the same GObject.
func (recv *WebResourcePrivate) Equals(other *WebResourcePrivate) bool {
	return other.ToC() == recv.ToC()
}

// WebViewBaseClass is a wrapper around the C record WebKitWebViewBaseClass.
type WebViewBaseClass struct {
	native *C.WebKitWebViewBaseClass
	// parentClass : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func WebViewBaseClassNewFromC(u unsafe.Pointer) *WebViewBaseClass {
	c := (*C.WebKitWebViewBaseClass)(u)
	if c == nil {
		return nil
	}

	g := &WebViewBaseClass{native: c}

	return g
}

func (recv *WebViewBaseClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebViewBaseClass with another WebViewBaseClass, and returns true if they represent the same GObject.
func (recv *WebViewBaseClass) Equals(other *WebViewBaseClass) bool {
	return other.ToC() == recv.ToC()
}

// WebViewBasePrivate is a wrapper around the C record WebKitWebViewBasePrivate.
type WebViewBasePrivate struct {
	native *C.WebKitWebViewBasePrivate
}

func WebViewBasePrivateNewFromC(u unsafe.Pointer) *WebViewBasePrivate {
	c := (*C.WebKitWebViewBasePrivate)(u)
	if c == nil {
		return nil
	}

	g := &WebViewBasePrivate{native: c}

	return g
}

func (recv *WebViewBasePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebViewBasePrivate with another WebViewBasePrivate, and returns true if they represent the same GObject.
func (recv *WebViewBasePrivate) Equals(other *WebViewBasePrivate) bool {
	return other.ToC() == recv.ToC()
}

// WebViewClass is a wrapper around the C record WebKitWebViewClass.
type WebViewClass struct {
	native *C.WebKitWebViewClass
	// parent : record
	// no type for load_changed
	// no type for load_failed
	// no type for create
	// no type for ready_to_show
	// no type for run_as_modal
	// no type for close
	// no type for script_dialog
	// no type for decide_policy
	// no type for permission_request
	// no type for mouse_target_changed
	// no type for print
	// no type for resource_load_started
	// no type for enter_fullscreen
	// no type for leave_fullscreen
	// no type for run_file_chooser
	// no type for context_menu
	// no type for context_menu_dismissed
	// no type for submit_form
	// no type for insecure_content_detected
	// no type for web_process_crashed
	// no type for authenticate
	// no type for load_failed_with_tls_errors
	// no type for show_notification
	// no type for run_color_chooser
	// no type for show_option_menu
	// no type for web_process_terminated
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
}

func WebViewClassNewFromC(u unsafe.Pointer) *WebViewClass {
	c := (*C.WebKitWebViewClass)(u)
	if c == nil {
		return nil
	}

	g := &WebViewClass{native: c}

	return g
}

func (recv *WebViewClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebViewClass with another WebViewClass, and returns true if they represent the same GObject.
func (recv *WebViewClass) Equals(other *WebViewClass) bool {
	return other.ToC() == recv.ToC()
}

// WebViewPrivate is a wrapper around the C record WebKitWebViewPrivate.
type WebViewPrivate struct {
	native *C.WebKitWebViewPrivate
}

func WebViewPrivateNewFromC(u unsafe.Pointer) *WebViewPrivate {
	c := (*C.WebKitWebViewPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WebViewPrivate{native: c}

	return g
}

func (recv *WebViewPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebViewPrivate with another WebViewPrivate, and returns true if they represent the same GObject.
func (recv *WebViewPrivate) Equals(other *WebViewPrivate) bool {
	return other.ToC() == recv.ToC()
}

// WebViewSessionState is a wrapper around the C record WebKitWebViewSessionState.
type WebViewSessionState struct {
	native *C.WebKitWebViewSessionState
}

func WebViewSessionStateNewFromC(u unsafe.Pointer) *WebViewSessionState {
	c := (*C.WebKitWebViewSessionState)(u)
	if c == nil {
		return nil
	}

	g := &WebViewSessionState{native: c}

	return g
}

func (recv *WebViewSessionState) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebViewSessionState with another WebViewSessionState, and returns true if they represent the same GObject.
func (recv *WebViewSessionState) Equals(other *WebViewSessionState) bool {
	return other.ToC() == recv.ToC()
}

// WebsiteData is a wrapper around the C record WebKitWebsiteData.
type WebsiteData struct {
	native *C.WebKitWebsiteData
}

func WebsiteDataNewFromC(u unsafe.Pointer) *WebsiteData {
	c := (*C.WebKitWebsiteData)(u)
	if c == nil {
		return nil
	}

	g := &WebsiteData{native: c}

	return g
}

func (recv *WebsiteData) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebsiteData with another WebsiteData, and returns true if they represent the same GObject.
func (recv *WebsiteData) Equals(other *WebsiteData) bool {
	return other.ToC() == recv.ToC()
}

// WebsiteDataManagerClass is a wrapper around the C record WebKitWebsiteDataManagerClass.
type WebsiteDataManagerClass struct {
	native *C.WebKitWebsiteDataManagerClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func WebsiteDataManagerClassNewFromC(u unsafe.Pointer) *WebsiteDataManagerClass {
	c := (*C.WebKitWebsiteDataManagerClass)(u)
	if c == nil {
		return nil
	}

	g := &WebsiteDataManagerClass{native: c}

	return g
}

func (recv *WebsiteDataManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebsiteDataManagerClass with another WebsiteDataManagerClass, and returns true if they represent the same GObject.
func (recv *WebsiteDataManagerClass) Equals(other *WebsiteDataManagerClass) bool {
	return other.ToC() == recv.ToC()
}

// WebsiteDataManagerPrivate is a wrapper around the C record WebKitWebsiteDataManagerPrivate.
type WebsiteDataManagerPrivate struct {
	native *C.WebKitWebsiteDataManagerPrivate
}

func WebsiteDataManagerPrivateNewFromC(u unsafe.Pointer) *WebsiteDataManagerPrivate {
	c := (*C.WebKitWebsiteDataManagerPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WebsiteDataManagerPrivate{native: c}

	return g
}

func (recv *WebsiteDataManagerPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WebsiteDataManagerPrivate with another WebsiteDataManagerPrivate, and returns true if they represent the same GObject.
func (recv *WebsiteDataManagerPrivate) Equals(other *WebsiteDataManagerPrivate) bool {
	return other.ToC() == recv.ToC()
}

// WindowPropertiesClass is a wrapper around the C record WebKitWindowPropertiesClass.
type WindowPropertiesClass struct {
	native *C.WebKitWindowPropertiesClass
	// parent_class : record
	// no type for _webkit_reserved0
	// no type for _webkit_reserved1
	// no type for _webkit_reserved2
	// no type for _webkit_reserved3
}

func WindowPropertiesClassNewFromC(u unsafe.Pointer) *WindowPropertiesClass {
	c := (*C.WebKitWindowPropertiesClass)(u)
	if c == nil {
		return nil
	}

	g := &WindowPropertiesClass{native: c}

	return g
}

func (recv *WindowPropertiesClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WindowPropertiesClass with another WindowPropertiesClass, and returns true if they represent the same GObject.
func (recv *WindowPropertiesClass) Equals(other *WindowPropertiesClass) bool {
	return other.ToC() == recv.ToC()
}

// WindowPropertiesPrivate is a wrapper around the C record WebKitWindowPropertiesPrivate.
type WindowPropertiesPrivate struct {
	native *C.WebKitWindowPropertiesPrivate
}

func WindowPropertiesPrivateNewFromC(u unsafe.Pointer) *WindowPropertiesPrivate {
	c := (*C.WebKitWindowPropertiesPrivate)(u)
	if c == nil {
		return nil
	}

	g := &WindowPropertiesPrivate{native: c}

	return g
}

func (recv *WindowPropertiesPrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this WindowPropertiesPrivate with another WindowPropertiesPrivate, and returns true if they represent the same GObject.
func (recv *WindowPropertiesPrivate) Equals(other *WindowPropertiesPrivate) bool {
	return other.ToC() == recv.ToC()
}
