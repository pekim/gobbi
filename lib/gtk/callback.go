// This is a generated file - DO NOT EDIT

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void callback_accelmapforeachHandler(GObject *, gpointer, gchar*, guint, GdkModifierType, gboolean);

*/
/*

	gint callback_assistantpagefuncHandler(GObject *, gint, gpointer);

*/
/*

	void callback_callbackHandler(GObject *, GtkWidget*, gpointer);

*/
/*

	gboolean callback_cellalloccallbackHandler(GObject *, GtkCellRenderer*, GdkRectangle*, GdkRectangle*, gpointer);

*/
/*

	gboolean callback_cellcallbackHandler(GObject *, GtkCellRenderer*, gpointer);

*/
/*

	void callback_celllayoutdatafuncHandler(GObject *, GtkCellLayout*, GtkCellRenderer*, GtkTreeModel*, GtkTreeIter*, gpointer);

*/
/*

	void callback_clipboardreceivedfuncHandler(GObject *, GtkClipboard*, GtkSelectionData*, gpointer);

*/
/*

	void callback_clipboardtextreceivedfuncHandler(GObject *, GtkClipboard*, gchar*, gpointer);

*/
/*

	gboolean callback_entrycompletionmatchfuncHandler(GObject *, GtkEntryCompletion*, gchar*, GtkTreeIter*, gpointer);

*/
/*

	gboolean callback_filefilterfuncHandler(GObject *, GtkFileFilterInfo*, gpointer);

*/
/*

	gboolean callback_fontfilterfuncHandler(GObject *, PangoFontFamily*, PangoFontFace*, gpointer);

*/
/*

	void callback_iconviewforeachfuncHandler(GObject *, GtkIconView*, GtkTreePath*, gpointer);

*/
/*

	void callback_pagesetupdonefuncHandler(GObject *, GtkPageSetup*, gpointer);

*/
/*

	void callback_printsettingsfuncHandler(GObject *, gchar*, gchar*, gpointer);

*/
/*

	gboolean callback_recentfilterfuncHandler(GObject *, GtkRecentFilterInfo*, gpointer);

*/
/*

	gint callback_recentsortfuncHandler(GObject *, GtkRecentInfo*, GtkRecentInfo*, gpointer);

*/
/*

	gboolean callback_textcharpredicateHandler(GObject *, gunichar, gpointer);

*/
/*

	void callback_texttagtableforeachHandler(GObject *, GtkTextTag*, gpointer);

*/
/*

	void callback_treecelldatafuncHandler(GObject *, GtkTreeViewColumn*, GtkCellRenderer*, GtkTreeModel*, GtkTreeIter*, gpointer);

*/
/*

	void callback_treedestroycountfuncHandler(GObject *, GtkTreeView*, GtkTreePath*, gint, gpointer);

*/
/*

	gint callback_treeitercomparefuncHandler(GObject *, GtkTreeModel*, GtkTreeIter*, GtkTreeIter*, gpointer);

*/
/*

	gboolean callback_treemodelfiltervisiblefuncHandler(GObject *, GtkTreeModel*, GtkTreeIter*, gpointer);

*/
/*

	gboolean callback_treemodelforeachfuncHandler(GObject *, GtkTreeModel*, GtkTreePath*, GtkTreeIter*, gpointer);

*/
/*

	void callback_treeselectionforeachfuncHandler(GObject *, GtkTreeModel*, GtkTreePath*, GtkTreeIter*, gpointer);

*/
/*

	gboolean callback_treeselectionfuncHandler(GObject *, GtkTreeSelection*, GtkTreeModel*, GtkTreePath*, gboolean, gpointer);

*/
/*

	gboolean callback_treeviewcolumndropfuncHandler(GObject *, GtkTreeView*, GtkTreeViewColumn*, GtkTreeViewColumn*, GtkTreeViewColumn*, gpointer);

*/
/*

	void callback_treeviewmappingfuncHandler(GObject *, GtkTreeView*, GtkTreePath*, gpointer);

*/
/*

	gboolean callback_treeviewrowseparatorfuncHandler(GObject *, GtkTreeModel*, GtkTreeIter*, gpointer);

*/
/*

	void callback_treeviewsearchpositionfuncHandler(GObject *, GtkTreeView*, GtkWidget*, gpointer);

*/
import "C"

// Unsupported : callback AccelGroupActivate : no [user_]data param

var callbackAccelmapforeachId int
var callbackAccelmapforeachMap = make(map[int]AccelmapforeachCallback)
var callbackAccelmapforeachLock sync.RWMutex

// AccelmapforeachCallback is a callback function for a 'AccelMapForeach' callback.
type AccelmapforeachCallback func(accelPath string, accelKey uint32, accelMods gdk.ModifierType, changed bool)

//export callback_accelmapforeachHandler
func callback_accelmapforeachHandler(_ *C.GObject, c_data C.gpointer, c_accel_path *C.gchar, c_accel_key C.guint, c_accel_mods C.guint, c_changed C.gboolean) {
	callbackAccelmapforeachLock.RLock()
	defer callbackAccelmapforeachLock.RUnlock()

	accelPath := C.GoString(c_accel_path)

	accelKey := uint32(c_accel_key)

	accelMods := gdk.ModifierType(c_accel_mods)

	changed := c_changed == C.TRUE

	index := int(uintptr(c_data))
	callback := callbackAccelmapforeachMap[index]
	callback(accelPath, accelKey, accelMods, changed)
}

var callbackAssistantpagefuncId int
var callbackAssistantpagefuncMap = make(map[int]AssistantpagefuncCallback)
var callbackAssistantpagefuncLock sync.RWMutex

// AssistantpagefuncCallback is a callback function for a 'AssistantPageFunc' callback.
type AssistantpagefuncCallback func(currentPage int32) int32

//export callback_assistantpagefuncHandler
func callback_assistantpagefuncHandler(_ *C.GObject, c_current_page C.gint, c_data C.gpointer) C.gint {
	callbackAssistantpagefuncLock.RLock()
	defer callbackAssistantpagefuncLock.RUnlock()

	currentPage := int32(c_current_page)

	index := int(uintptr(c_data))
	callback := callbackAssistantpagefuncMap[index]
	retGo := callback(currentPage)
	retC :=
		(C.gint)(retGo)
	return retC
}

var callbackCallbackId int
var callbackCallbackMap = make(map[int]CallbackCallback)
var callbackCallbackLock sync.RWMutex

// CallbackCallback is a callback function for a 'Callback' callback.
type CallbackCallback func(widget *Widget)

//export callback_callbackHandler
func callback_callbackHandler(_ *C.GObject, c_widget *C.GtkWidget, c_data C.gpointer) {
	callbackCallbackLock.RLock()
	defer callbackCallbackLock.RUnlock()

	widget := WidgetNewFromC(unsafe.Pointer(c_widget))

	index := int(uintptr(c_data))
	callback := callbackCallbackMap[index]
	callback(widget)
}

var callbackCellalloccallbackId int
var callbackCellalloccallbackMap = make(map[int]CellalloccallbackCallback)
var callbackCellalloccallbackLock sync.RWMutex

// CellalloccallbackCallback is a callback function for a 'CellAllocCallback' callback.
type CellalloccallbackCallback func(renderer *CellRenderer, cellArea *gdk.Rectangle, cellBackground *gdk.Rectangle) bool

//export callback_cellalloccallbackHandler
func callback_cellalloccallbackHandler(_ *C.GObject, c_renderer *C.GtkCellRenderer, c_cell_area *C.GdkRectangle, c_cell_background *C.GdkRectangle, c_data C.gpointer) C.gboolean {
	callbackCellalloccallbackLock.RLock()
	defer callbackCellalloccallbackLock.RUnlock()

	renderer := CellRendererNewFromC(unsafe.Pointer(c_renderer))

	cellArea := gdk.RectangleNewFromC(unsafe.Pointer(c_cell_area))

	cellBackground := gdk.RectangleNewFromC(unsafe.Pointer(c_cell_background))

	index := int(uintptr(c_data))
	callback := callbackCellalloccallbackMap[index]
	retGo := callback(renderer, cellArea, cellBackground)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackCellcallbackId int
var callbackCellcallbackMap = make(map[int]CellcallbackCallback)
var callbackCellcallbackLock sync.RWMutex

// CellcallbackCallback is a callback function for a 'CellCallback' callback.
type CellcallbackCallback func(renderer *CellRenderer) bool

//export callback_cellcallbackHandler
func callback_cellcallbackHandler(_ *C.GObject, c_renderer *C.GtkCellRenderer, c_data C.gpointer) C.gboolean {
	callbackCellcallbackLock.RLock()
	defer callbackCellcallbackLock.RUnlock()

	renderer := CellRendererNewFromC(unsafe.Pointer(c_renderer))

	index := int(uintptr(c_data))
	callback := callbackCellcallbackMap[index]
	retGo := callback(renderer)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackCelllayoutdatafuncId int
var callbackCelllayoutdatafuncMap = make(map[int]CelllayoutdatafuncCallback)
var callbackCelllayoutdatafuncLock sync.RWMutex

// CelllayoutdatafuncCallback is a callback function for a 'CellLayoutDataFunc' callback.
type CelllayoutdatafuncCallback func(cellLayout *CellLayout, cell *CellRenderer, treeModel *TreeModel, iter *TreeIter)

//export callback_celllayoutdatafuncHandler
func callback_celllayoutdatafuncHandler(_ *C.GObject, c_cell_layout *C.GtkCellLayout, c_cell *C.GtkCellRenderer, c_tree_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, c_data C.gpointer) {
	callbackCelllayoutdatafuncLock.RLock()
	defer callbackCelllayoutdatafuncLock.RUnlock()

	cellLayout := CellLayoutNewFromC(unsafe.Pointer(c_cell_layout))

	cell := CellRendererNewFromC(unsafe.Pointer(c_cell))

	treeModel := TreeModelNewFromC(unsafe.Pointer(c_tree_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(c_data))
	callback := callbackCelllayoutdatafuncMap[index]
	callback(cellLayout, cell, treeModel, iter)
}

// Unsupported : callback ClipboardClearFunc : no [user_]data param

// Unsupported : callback ClipboardGetFunc : no [user_]data param

var callbackClipboardreceivedfuncId int
var callbackClipboardreceivedfuncMap = make(map[int]ClipboardreceivedfuncCallback)
var callbackClipboardreceivedfuncLock sync.RWMutex

// ClipboardreceivedfuncCallback is a callback function for a 'ClipboardReceivedFunc' callback.
type ClipboardreceivedfuncCallback func(clipboard *Clipboard, selectionData *SelectionData)

//export callback_clipboardreceivedfuncHandler
func callback_clipboardreceivedfuncHandler(_ *C.GObject, c_clipboard *C.GtkClipboard, c_selection_data *C.GtkSelectionData, c_data C.gpointer) {
	callbackClipboardreceivedfuncLock.RLock()
	defer callbackClipboardreceivedfuncLock.RUnlock()

	clipboard := ClipboardNewFromC(unsafe.Pointer(c_clipboard))

	selectionData := SelectionDataNewFromC(unsafe.Pointer(c_selection_data))

	index := int(uintptr(c_data))
	callback := callbackClipboardreceivedfuncMap[index]
	callback(clipboard, selectionData)
}

var callbackClipboardtextreceivedfuncId int
var callbackClipboardtextreceivedfuncMap = make(map[int]ClipboardtextreceivedfuncCallback)
var callbackClipboardtextreceivedfuncLock sync.RWMutex

// ClipboardtextreceivedfuncCallback is a callback function for a 'ClipboardTextReceivedFunc' callback.
type ClipboardtextreceivedfuncCallback func(clipboard *Clipboard, text string)

//export callback_clipboardtextreceivedfuncHandler
func callback_clipboardtextreceivedfuncHandler(_ *C.GObject, c_clipboard *C.GtkClipboard, c_text *C.gchar, c_data C.gpointer) {
	callbackClipboardtextreceivedfuncLock.RLock()
	defer callbackClipboardtextreceivedfuncLock.RUnlock()

	clipboard := ClipboardNewFromC(unsafe.Pointer(c_clipboard))

	text := C.GoString(c_text)

	index := int(uintptr(c_data))
	callback := callbackClipboardtextreceivedfuncMap[index]
	callback(clipboard, text)
}

// Unsupported : callback ColorSelectionChangePaletteFunc : no [user_]data param

var callbackEntrycompletionmatchfuncId int
var callbackEntrycompletionmatchfuncMap = make(map[int]EntrycompletionmatchfuncCallback)
var callbackEntrycompletionmatchfuncLock sync.RWMutex

// EntrycompletionmatchfuncCallback is a callback function for a 'EntryCompletionMatchFunc' callback.
type EntrycompletionmatchfuncCallback func(completion *EntryCompletion, key string, iter *TreeIter) bool

//export callback_entrycompletionmatchfuncHandler
func callback_entrycompletionmatchfuncHandler(_ *C.GObject, c_completion *C.GtkEntryCompletion, c_key *C.gchar, c_iter *C.GtkTreeIter, c_user_data C.gpointer) C.gboolean {
	callbackEntrycompletionmatchfuncLock.RLock()
	defer callbackEntrycompletionmatchfuncLock.RUnlock()

	completion := EntryCompletionNewFromC(unsafe.Pointer(c_completion))

	key := C.GoString(c_key)

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(c_user_data))
	callback := callbackEntrycompletionmatchfuncMap[index]
	retGo := callback(completion, key, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackFilefilterfuncId int
var callbackFilefilterfuncMap = make(map[int]FilefilterfuncCallback)
var callbackFilefilterfuncLock sync.RWMutex

// FilefilterfuncCallback is a callback function for a 'FileFilterFunc' callback.
type FilefilterfuncCallback func(filterInfo *FileFilterInfo) bool

//export callback_filefilterfuncHandler
func callback_filefilterfuncHandler(_ *C.GObject, c_filter_info *C.GtkFileFilterInfo, c_data C.gpointer) C.gboolean {
	callbackFilefilterfuncLock.RLock()
	defer callbackFilefilterfuncLock.RUnlock()

	filterInfo := FileFilterInfoNewFromC(unsafe.Pointer(c_filter_info))

	index := int(uintptr(c_data))
	callback := callbackFilefilterfuncMap[index]
	retGo := callback(filterInfo)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackFontfilterfuncId int
var callbackFontfilterfuncMap = make(map[int]FontfilterfuncCallback)
var callbackFontfilterfuncLock sync.RWMutex

// FontfilterfuncCallback is a callback function for a 'FontFilterFunc' callback.
type FontfilterfuncCallback func(family *pango.FontFamily, face *pango.FontFace) bool

//export callback_fontfilterfuncHandler
func callback_fontfilterfuncHandler(_ *C.GObject, c_family *C.PangoFontFamily, c_face *C.PangoFontFace, c_data C.gpointer) C.gboolean {
	callbackFontfilterfuncLock.RLock()
	defer callbackFontfilterfuncLock.RUnlock()

	family := pango.FontFamilyNewFromC(unsafe.Pointer(c_family))

	face := pango.FontFaceNewFromC(unsafe.Pointer(c_face))

	index := int(uintptr(c_data))
	callback := callbackFontfilterfuncMap[index]
	retGo := callback(family, face)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackIconviewforeachfuncId int
var callbackIconviewforeachfuncMap = make(map[int]IconviewforeachfuncCallback)
var callbackIconviewforeachfuncLock sync.RWMutex

// IconviewforeachfuncCallback is a callback function for a 'IconViewForeachFunc' callback.
type IconviewforeachfuncCallback func(iconView *IconView, path *TreePath)

//export callback_iconviewforeachfuncHandler
func callback_iconviewforeachfuncHandler(_ *C.GObject, c_icon_view *C.GtkIconView, c_path *C.GtkTreePath, c_data C.gpointer) {
	callbackIconviewforeachfuncLock.RLock()
	defer callbackIconviewforeachfuncLock.RUnlock()

	iconView := IconViewNewFromC(unsafe.Pointer(c_icon_view))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	index := int(uintptr(c_data))
	callback := callbackIconviewforeachfuncMap[index]
	callback(iconView, path)
}

// Unsupported : callback KeySnoopFunc : no [user_]data param

// Unsupported : callback MenuDetachFunc : no [user_]data param

// Unsupported : callback MenuPositionFunc : unsupported parameter x, direction inout

// Unsupported : callback ModuleInitFunc : no [user_]data param

var callbackPagesetupdonefuncId int
var callbackPagesetupdonefuncMap = make(map[int]PagesetupdonefuncCallback)
var callbackPagesetupdonefuncLock sync.RWMutex

// PagesetupdonefuncCallback is a callback function for a 'PageSetupDoneFunc' callback.
type PagesetupdonefuncCallback func(pageSetup *PageSetup)

//export callback_pagesetupdonefuncHandler
func callback_pagesetupdonefuncHandler(_ *C.GObject, c_page_setup *C.GtkPageSetup, c_data C.gpointer) {
	callbackPagesetupdonefuncLock.RLock()
	defer callbackPagesetupdonefuncLock.RUnlock()

	pageSetup := PageSetupNewFromC(unsafe.Pointer(c_page_setup))

	index := int(uintptr(c_data))
	callback := callbackPagesetupdonefuncMap[index]
	callback(pageSetup)
}

var callbackPrintsettingsfuncId int
var callbackPrintsettingsfuncMap = make(map[int]PrintsettingsfuncCallback)
var callbackPrintsettingsfuncLock sync.RWMutex

// PrintsettingsfuncCallback is a callback function for a 'PrintSettingsFunc' callback.
type PrintsettingsfuncCallback func(key string, value string)

//export callback_printsettingsfuncHandler
func callback_printsettingsfuncHandler(_ *C.GObject, c_key *C.gchar, c_value *C.gchar, c_user_data C.gpointer) {
	callbackPrintsettingsfuncLock.RLock()
	defer callbackPrintsettingsfuncLock.RUnlock()

	key := C.GoString(c_key)

	value := C.GoString(c_value)

	index := int(uintptr(c_user_data))
	callback := callbackPrintsettingsfuncMap[index]
	callback(key, value)
}

// Unsupported : callback RcPropertyParser : no [user_]data param

var callbackRecentfilterfuncId int
var callbackRecentfilterfuncMap = make(map[int]RecentfilterfuncCallback)
var callbackRecentfilterfuncLock sync.RWMutex

// RecentfilterfuncCallback is a callback function for a 'RecentFilterFunc' callback.
type RecentfilterfuncCallback func(filterInfo *RecentFilterInfo) bool

//export callback_recentfilterfuncHandler
func callback_recentfilterfuncHandler(_ *C.GObject, c_filter_info *C.GtkRecentFilterInfo, c_user_data C.gpointer) C.gboolean {
	callbackRecentfilterfuncLock.RLock()
	defer callbackRecentfilterfuncLock.RUnlock()

	filterInfo := RecentFilterInfoNewFromC(unsafe.Pointer(c_filter_info))

	index := int(uintptr(c_user_data))
	callback := callbackRecentfilterfuncMap[index]
	retGo := callback(filterInfo)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackRecentsortfuncId int
var callbackRecentsortfuncMap = make(map[int]RecentsortfuncCallback)
var callbackRecentsortfuncLock sync.RWMutex

// RecentsortfuncCallback is a callback function for a 'RecentSortFunc' callback.
type RecentsortfuncCallback func(a *RecentInfo, b *RecentInfo) int32

//export callback_recentsortfuncHandler
func callback_recentsortfuncHandler(_ *C.GObject, c_a *C.GtkRecentInfo, c_b *C.GtkRecentInfo, c_user_data C.gpointer) C.gint {
	callbackRecentsortfuncLock.RLock()
	defer callbackRecentsortfuncLock.RUnlock()

	a := RecentInfoNewFromC(unsafe.Pointer(c_a))

	b := RecentInfoNewFromC(unsafe.Pointer(c_b))

	index := int(uintptr(c_user_data))
	callback := callbackRecentsortfuncMap[index]
	retGo := callback(a, b)
	retC :=
		(C.gint)(retGo)
	return retC
}

// Unsupported : callback StylePropertyParser : no [user_]data param

// Unsupported : callback TextBufferDeserializeFunc : unsupported parameter data, array type interface not found

// Unsupported : callback TextBufferSerializeFunc : unsupported parameter length, gsize*

var callbackTextcharpredicateId int
var callbackTextcharpredicateMap = make(map[int]TextcharpredicateCallback)
var callbackTextcharpredicateLock sync.RWMutex

// TextcharpredicateCallback is a callback function for a 'TextCharPredicate' callback.
type TextcharpredicateCallback func(ch rune) bool

//export callback_textcharpredicateHandler
func callback_textcharpredicateHandler(_ *C.GObject, c_ch C.gunichar, c_user_data C.gpointer) C.gboolean {
	callbackTextcharpredicateLock.RLock()
	defer callbackTextcharpredicateLock.RUnlock()

	ch := rune(c_ch)

	index := int(uintptr(c_user_data))
	callback := callbackTextcharpredicateMap[index]
	retGo := callback(ch)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackTexttagtableforeachId int
var callbackTexttagtableforeachMap = make(map[int]TexttagtableforeachCallback)
var callbackTexttagtableforeachLock sync.RWMutex

// TexttagtableforeachCallback is a callback function for a 'TextTagTableForeach' callback.
type TexttagtableforeachCallback func(tag *TextTag)

//export callback_texttagtableforeachHandler
func callback_texttagtableforeachHandler(_ *C.GObject, c_tag *C.GtkTextTag, c_data C.gpointer) {
	callbackTexttagtableforeachLock.RLock()
	defer callbackTexttagtableforeachLock.RUnlock()

	tag := TextTagNewFromC(unsafe.Pointer(c_tag))

	index := int(uintptr(c_data))
	callback := callbackTexttagtableforeachMap[index]
	callback(tag)
}

// Unsupported : callback TranslateFunc : no [user_]data param

var callbackTreecelldatafuncId int
var callbackTreecelldatafuncMap = make(map[int]TreecelldatafuncCallback)
var callbackTreecelldatafuncLock sync.RWMutex

// TreecelldatafuncCallback is a callback function for a 'TreeCellDataFunc' callback.
type TreecelldatafuncCallback func(treeColumn *TreeViewColumn, cell *CellRenderer, treeModel *TreeModel, iter *TreeIter)

//export callback_treecelldatafuncHandler
func callback_treecelldatafuncHandler(_ *C.GObject, c_tree_column *C.GtkTreeViewColumn, c_cell *C.GtkCellRenderer, c_tree_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, c_data C.gpointer) {
	callbackTreecelldatafuncLock.RLock()
	defer callbackTreecelldatafuncLock.RUnlock()

	treeColumn := TreeViewColumnNewFromC(unsafe.Pointer(c_tree_column))

	cell := CellRendererNewFromC(unsafe.Pointer(c_cell))

	treeModel := TreeModelNewFromC(unsafe.Pointer(c_tree_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(c_data))
	callback := callbackTreecelldatafuncMap[index]
	callback(treeColumn, cell, treeModel, iter)
}

var callbackTreedestroycountfuncId int
var callbackTreedestroycountfuncMap = make(map[int]TreedestroycountfuncCallback)
var callbackTreedestroycountfuncLock sync.RWMutex

// TreedestroycountfuncCallback is a callback function for a 'TreeDestroyCountFunc' callback.
type TreedestroycountfuncCallback func(treeView *TreeView, path *TreePath, children int32)

//export callback_treedestroycountfuncHandler
func callback_treedestroycountfuncHandler(_ *C.GObject, c_tree_view *C.GtkTreeView, c_path *C.GtkTreePath, c_children C.gint, c_user_data C.gpointer) {
	callbackTreedestroycountfuncLock.RLock()
	defer callbackTreedestroycountfuncLock.RUnlock()

	treeView := TreeViewNewFromC(unsafe.Pointer(c_tree_view))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	children := int32(c_children)

	index := int(uintptr(c_user_data))
	callback := callbackTreedestroycountfuncMap[index]
	callback(treeView, path, children)
}

var callbackTreeitercomparefuncId int
var callbackTreeitercomparefuncMap = make(map[int]TreeitercomparefuncCallback)
var callbackTreeitercomparefuncLock sync.RWMutex

// TreeitercomparefuncCallback is a callback function for a 'TreeIterCompareFunc' callback.
type TreeitercomparefuncCallback func(model *TreeModel, a *TreeIter, b *TreeIter) int32

//export callback_treeitercomparefuncHandler
func callback_treeitercomparefuncHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_a *C.GtkTreeIter, c_b *C.GtkTreeIter, c_user_data C.gpointer) C.gint {
	callbackTreeitercomparefuncLock.RLock()
	defer callbackTreeitercomparefuncLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	a := TreeIterNewFromC(unsafe.Pointer(c_a))

	b := TreeIterNewFromC(unsafe.Pointer(c_b))

	index := int(uintptr(c_user_data))
	callback := callbackTreeitercomparefuncMap[index]
	retGo := callback(model, a, b)
	retC :=
		(C.gint)(retGo)
	return retC
}

// Unsupported : callback TreeModelFilterModifyFunc : unsupported parameter value, direction out

var callbackTreemodelfiltervisiblefuncId int
var callbackTreemodelfiltervisiblefuncMap = make(map[int]TreemodelfiltervisiblefuncCallback)
var callbackTreemodelfiltervisiblefuncLock sync.RWMutex

// TreemodelfiltervisiblefuncCallback is a callback function for a 'TreeModelFilterVisibleFunc' callback.
type TreemodelfiltervisiblefuncCallback func(model *TreeModel, iter *TreeIter) bool

//export callback_treemodelfiltervisiblefuncHandler
func callback_treemodelfiltervisiblefuncHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, c_data C.gpointer) C.gboolean {
	callbackTreemodelfiltervisiblefuncLock.RLock()
	defer callbackTreemodelfiltervisiblefuncLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(c_data))
	callback := callbackTreemodelfiltervisiblefuncMap[index]
	retGo := callback(model, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackTreemodelforeachfuncId int
var callbackTreemodelforeachfuncMap = make(map[int]TreemodelforeachfuncCallback)
var callbackTreemodelforeachfuncLock sync.RWMutex

// TreemodelforeachfuncCallback is a callback function for a 'TreeModelForeachFunc' callback.
type TreemodelforeachfuncCallback func(model *TreeModel, path *TreePath, iter *TreeIter) bool

//export callback_treemodelforeachfuncHandler
func callback_treemodelforeachfuncHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_path *C.GtkTreePath, c_iter *C.GtkTreeIter, c_data C.gpointer) C.gboolean {
	callbackTreemodelforeachfuncLock.RLock()
	defer callbackTreemodelforeachfuncLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(c_data))
	callback := callbackTreemodelforeachfuncMap[index]
	retGo := callback(model, path, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackTreeselectionforeachfuncId int
var callbackTreeselectionforeachfuncMap = make(map[int]TreeselectionforeachfuncCallback)
var callbackTreeselectionforeachfuncLock sync.RWMutex

// TreeselectionforeachfuncCallback is a callback function for a 'TreeSelectionForeachFunc' callback.
type TreeselectionforeachfuncCallback func(model *TreeModel, path *TreePath, iter *TreeIter)

//export callback_treeselectionforeachfuncHandler
func callback_treeselectionforeachfuncHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_path *C.GtkTreePath, c_iter *C.GtkTreeIter, c_data C.gpointer) {
	callbackTreeselectionforeachfuncLock.RLock()
	defer callbackTreeselectionforeachfuncLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(c_data))
	callback := callbackTreeselectionforeachfuncMap[index]
	callback(model, path, iter)
}

var callbackTreeselectionfuncId int
var callbackTreeselectionfuncMap = make(map[int]TreeselectionfuncCallback)
var callbackTreeselectionfuncLock sync.RWMutex

// TreeselectionfuncCallback is a callback function for a 'TreeSelectionFunc' callback.
type TreeselectionfuncCallback func(selection *TreeSelection, model *TreeModel, path *TreePath, pathCurrentlySelected bool) bool

//export callback_treeselectionfuncHandler
func callback_treeselectionfuncHandler(_ *C.GObject, c_selection *C.GtkTreeSelection, c_model *C.GtkTreeModel, c_path *C.GtkTreePath, c_path_currently_selected C.gboolean, c_data C.gpointer) C.gboolean {
	callbackTreeselectionfuncLock.RLock()
	defer callbackTreeselectionfuncLock.RUnlock()

	selection := TreeSelectionNewFromC(unsafe.Pointer(c_selection))

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	pathCurrentlySelected := c_path_currently_selected == C.TRUE

	index := int(uintptr(c_data))
	callback := callbackTreeselectionfuncMap[index]
	retGo := callback(selection, model, path, pathCurrentlySelected)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackTreeviewcolumndropfuncId int
var callbackTreeviewcolumndropfuncMap = make(map[int]TreeviewcolumndropfuncCallback)
var callbackTreeviewcolumndropfuncLock sync.RWMutex

// TreeviewcolumndropfuncCallback is a callback function for a 'TreeViewColumnDropFunc' callback.
type TreeviewcolumndropfuncCallback func(treeView *TreeView, column *TreeViewColumn, prevColumn *TreeViewColumn, nextColumn *TreeViewColumn) bool

//export callback_treeviewcolumndropfuncHandler
func callback_treeviewcolumndropfuncHandler(_ *C.GObject, c_tree_view *C.GtkTreeView, c_column *C.GtkTreeViewColumn, c_prev_column *C.GtkTreeViewColumn, c_next_column *C.GtkTreeViewColumn, c_data C.gpointer) C.gboolean {
	callbackTreeviewcolumndropfuncLock.RLock()
	defer callbackTreeviewcolumndropfuncLock.RUnlock()

	treeView := TreeViewNewFromC(unsafe.Pointer(c_tree_view))

	column := TreeViewColumnNewFromC(unsafe.Pointer(c_column))

	prevColumn := TreeViewColumnNewFromC(unsafe.Pointer(c_prev_column))

	nextColumn := TreeViewColumnNewFromC(unsafe.Pointer(c_next_column))

	index := int(uintptr(c_data))
	callback := callbackTreeviewcolumndropfuncMap[index]
	retGo := callback(treeView, column, prevColumn, nextColumn)
	retC :=
		boolToGboolean(retGo)
	return retC
}

var callbackTreeviewmappingfuncId int
var callbackTreeviewmappingfuncMap = make(map[int]TreeviewmappingfuncCallback)
var callbackTreeviewmappingfuncLock sync.RWMutex

// TreeviewmappingfuncCallback is a callback function for a 'TreeViewMappingFunc' callback.
type TreeviewmappingfuncCallback func(treeView *TreeView, path *TreePath)

//export callback_treeviewmappingfuncHandler
func callback_treeviewmappingfuncHandler(_ *C.GObject, c_tree_view *C.GtkTreeView, c_path *C.GtkTreePath, c_user_data C.gpointer) {
	callbackTreeviewmappingfuncLock.RLock()
	defer callbackTreeviewmappingfuncLock.RUnlock()

	treeView := TreeViewNewFromC(unsafe.Pointer(c_tree_view))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	index := int(uintptr(c_user_data))
	callback := callbackTreeviewmappingfuncMap[index]
	callback(treeView, path)
}

var callbackTreeviewrowseparatorfuncId int
var callbackTreeviewrowseparatorfuncMap = make(map[int]TreeviewrowseparatorfuncCallback)
var callbackTreeviewrowseparatorfuncLock sync.RWMutex

// TreeviewrowseparatorfuncCallback is a callback function for a 'TreeViewRowSeparatorFunc' callback.
type TreeviewrowseparatorfuncCallback func(model *TreeModel, iter *TreeIter) bool

//export callback_treeviewrowseparatorfuncHandler
func callback_treeviewrowseparatorfuncHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, c_data C.gpointer) C.gboolean {
	callbackTreeviewrowseparatorfuncLock.RLock()
	defer callbackTreeviewrowseparatorfuncLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(c_data))
	callback := callbackTreeviewrowseparatorfuncMap[index]
	retGo := callback(model, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Unsupported : callback TreeViewSearchEqualFunc : no [user_]data param

var callbackTreeviewsearchpositionfuncId int
var callbackTreeviewsearchpositionfuncMap = make(map[int]TreeviewsearchpositionfuncCallback)
var callbackTreeviewsearchpositionfuncLock sync.RWMutex

// TreeviewsearchpositionfuncCallback is a callback function for a 'TreeViewSearchPositionFunc' callback.
type TreeviewsearchpositionfuncCallback func(treeView *TreeView, searchDialog *Widget)

//export callback_treeviewsearchpositionfuncHandler
func callback_treeviewsearchpositionfuncHandler(_ *C.GObject, c_tree_view *C.GtkTreeView, c_search_dialog *C.GtkWidget, c_user_data C.gpointer) {
	callbackTreeviewsearchpositionfuncLock.RLock()
	defer callbackTreeviewsearchpositionfuncLock.RUnlock()

	treeView := TreeViewNewFromC(unsafe.Pointer(c_tree_view))

	searchDialog := WidgetNewFromC(unsafe.Pointer(c_search_dialog))

	index := int(uintptr(c_user_data))
	callback := callbackTreeviewsearchpositionfuncMap[index]
	callback(treeView, searchDialog)
}
