// This is a generated file - DO NOT EDIT

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	gboolean callback_accelgroupactivateHandler(GObject *, GtkAccelGroup*, GObject*, guint, GdkModifierType, gpointer);

*/
/*

	void callback_accelmapforeachHandler(GObject *, gpointer, const gchar*, guint, GdkModifierType, gboolean, gpointer);

*/
/*

	gint callback_assistantpagefuncHandler(GObject *, gint, gpointer, gpointer);

*/
/*

	void callback_callbackHandler(GObject *, GtkWidget*, gpointer, gpointer);

*/
/*

	gboolean callback_cellalloccallbackHandler(GObject *, GtkCellRenderer*, const GdkRectangle*, const GdkRectangle*, gpointer, gpointer);

*/
/*

	gboolean callback_cellcallbackHandler(GObject *, GtkCellRenderer*, gpointer, gpointer);

*/
/*

	void callback_celllayoutdatafuncHandler(GObject *, GtkCellLayout*, GtkCellRenderer*, GtkTreeModel*, GtkTreeIter*, gpointer, gpointer);

*/
/*

	void callback_clipboardreceivedfuncHandler(GObject *, GtkClipboard*, GtkSelectionData*, gpointer, gpointer);

*/
/*

	void callback_clipboardtextreceivedfuncHandler(GObject *, GtkClipboard*, const gchar*, gpointer, gpointer);

*/
/*

	gboolean callback_entrycompletionmatchfuncHandler(GObject *, GtkEntryCompletion*, const gchar*, GtkTreeIter*, gpointer, gpointer);

*/
/*

	gboolean callback_filefilterfuncHandler(GObject *, const GtkFileFilterInfo*, gpointer, gpointer);

*/
/*

	gboolean callback_fontfilterfuncHandler(GObject *, const PangoFontFamily*, const PangoFontFace*, gpointer, gpointer);

*/
/*

	void callback_iconviewforeachfuncHandler(GObject *, GtkIconView*, GtkTreePath*, gpointer, gpointer);

*/
/*

	void callback_menudetachfuncHandler(GObject *, GtkWidget*, GtkMenu*, gpointer);

*/
/*

	void callback_menupositionfuncHandler(GObject *, GtkMenu*, gint*, gint*, gboolean*, gpointer, gpointer);

*/
/*

	void callback_pagesetupdonefuncHandler(GObject *, GtkPageSetup*, gpointer, gpointer);

*/
/*

	void callback_printsettingsfuncHandler(GObject *, const gchar*, const gchar*, gpointer, gpointer);

*/
/*

	gboolean callback_rcpropertyparserHandler(GObject *, const GParamSpec*, const GString*, GValue*, gpointer);

*/
/*

	gboolean callback_recentfilterfuncHandler(GObject *, const GtkRecentFilterInfo*, gpointer, gpointer);

*/
/*

	gint callback_recentsortfuncHandler(GObject *, GtkRecentInfo*, GtkRecentInfo*, gpointer, gpointer);

*/
/*

	gboolean callback_stylepropertyparserHandler(GObject *, const gchar*, GValue*, gpointer);

*/
/*

	gboolean callback_textbufferdeserializefuncHandler(GObject *, GtkTextBuffer*, GtkTextBuffer*, GtkTextIter*, guint8*, gsize, gboolean, gpointer, gpointer);

*/
/*

	guint8* callback_textbufferserializefuncHandler(GObject *, GtkTextBuffer*, GtkTextBuffer*, const GtkTextIter*, const GtkTextIter*, gsize*, gpointer, gpointer);

*/
/*

	gboolean callback_textcharpredicateHandler(GObject *, gunichar, gpointer, gpointer);

*/
/*

	void callback_texttagtableforeachHandler(GObject *, GtkTextTag*, gpointer, gpointer);

*/
/*

	void callback_treecelldatafuncHandler(GObject *, GtkTreeViewColumn*, GtkCellRenderer*, GtkTreeModel*, GtkTreeIter*, gpointer, gpointer);

*/
/*

	void callback_treedestroycountfuncHandler(GObject *, GtkTreeView*, GtkTreePath*, gint, gpointer, gpointer);

*/
/*

	gint callback_treeitercomparefuncHandler(GObject *, GtkTreeModel*, GtkTreeIter*, GtkTreeIter*, gpointer, gpointer);

*/
/*

	void callback_treemodelfiltermodifyfuncHandler(GObject *, GtkTreeModel*, GtkTreeIter*, GValue*, gint, gpointer, gpointer);

*/
/*

	gboolean callback_treemodelfiltervisiblefuncHandler(GObject *, GtkTreeModel*, GtkTreeIter*, gpointer, gpointer);

*/
/*

	gboolean callback_treemodelforeachfuncHandler(GObject *, GtkTreeModel*, GtkTreePath*, GtkTreeIter*, gpointer, gpointer);

*/
/*

	void callback_treeselectionforeachfuncHandler(GObject *, GtkTreeModel*, GtkTreePath*, GtkTreeIter*, gpointer, gpointer);

*/
/*

	gboolean callback_treeselectionfuncHandler(GObject *, GtkTreeSelection*, GtkTreeModel*, GtkTreePath*, gboolean, gpointer, gpointer);

*/
/*

	gboolean callback_treeviewcolumndropfuncHandler(GObject *, GtkTreeView*, GtkTreeViewColumn*, GtkTreeViewColumn*, GtkTreeViewColumn*, gpointer, gpointer);

*/
/*

	void callback_treeviewmappingfuncHandler(GObject *, GtkTreeView*, GtkTreePath*, gpointer, gpointer);

*/
/*

	gboolean callback_treeviewrowseparatorfuncHandler(GObject *, GtkTreeModel*, GtkTreeIter*, gpointer, gpointer);

*/
/*

	void callback_treeviewsearchpositionfuncHandler(GObject *, GtkTreeView*, GtkWidget*, gpointer, gpointer);

*/
import "C"

var callbackAccelgroupactivateId int
var callbackAccelgroupactivateMap = make(map[int]AccelgroupactivateCallback)
var callbackAccelgroupactivateLock sync.RWMutex

// AccelgroupactivateCallback is a callback function for a 'AccelGroupActivate' callback.
type AccelgroupactivateCallback func(accelGroup *AccelGroup, acceleratable *gobject.Object, keyval uint32, modifier gdk.ModifierType) bool

var callbackAccelmapforeachId int
var callbackAccelmapforeachMap = make(map[int]AccelmapforeachCallback)
var callbackAccelmapforeachLock sync.RWMutex

// AccelmapforeachCallback is a callback function for a 'AccelMapForeach' callback.
type AccelmapforeachCallback func(accelPath string, accelKey uint32, accelMods gdk.ModifierType, changed bool)

var callbackAssistantpagefuncId int
var callbackAssistantpagefuncMap = make(map[int]AssistantpagefuncCallback)
var callbackAssistantpagefuncLock sync.RWMutex

// AssistantpagefuncCallback is a callback function for a 'AssistantPageFunc' callback.
type AssistantpagefuncCallback func(currentPage int32) int32

var callbackCallbackId int
var callbackCallbackMap = make(map[int]CallbackCallback)
var callbackCallbackLock sync.RWMutex

// CallbackCallback is a callback function for a 'Callback' callback.
type CallbackCallback func(widget *Widget)

var callbackCellalloccallbackId int
var callbackCellalloccallbackMap = make(map[int]CellalloccallbackCallback)
var callbackCellalloccallbackLock sync.RWMutex

// CellalloccallbackCallback is a callback function for a 'CellAllocCallback' callback.
type CellalloccallbackCallback func(renderer *CellRenderer, cellArea *gdk.Rectangle, cellBackground *gdk.Rectangle) bool

var callbackCellcallbackId int
var callbackCellcallbackMap = make(map[int]CellcallbackCallback)
var callbackCellcallbackLock sync.RWMutex

// CellcallbackCallback is a callback function for a 'CellCallback' callback.
type CellcallbackCallback func(renderer *CellRenderer) bool

var callbackCelllayoutdatafuncId int
var callbackCelllayoutdatafuncMap = make(map[int]CelllayoutdatafuncCallback)
var callbackCelllayoutdatafuncLock sync.RWMutex

// CelllayoutdatafuncCallback is a callback function for a 'CellLayoutDataFunc' callback.
type CelllayoutdatafuncCallback func(cellLayout *CellLayout, cell *CellRenderer, treeModel *TreeModel, iter *TreeIter)

// Unsupported : callback ClipboardClearFunc : unsupported parameter user_data_or_owner : no type generator for gpointer (gpointer) for param user_data_or_owner

// Unsupported : callback ClipboardGetFunc : unsupported parameter user_data_or_owner : no type generator for gpointer (gpointer) for param user_data_or_owner

var callbackClipboardreceivedfuncId int
var callbackClipboardreceivedfuncMap = make(map[int]ClipboardreceivedfuncCallback)
var callbackClipboardreceivedfuncLock sync.RWMutex

// ClipboardreceivedfuncCallback is a callback function for a 'ClipboardReceivedFunc' callback.
type ClipboardreceivedfuncCallback func(clipboard *Clipboard, selectionData *SelectionData)

var callbackClipboardtextreceivedfuncId int
var callbackClipboardtextreceivedfuncMap = make(map[int]ClipboardtextreceivedfuncCallback)
var callbackClipboardtextreceivedfuncLock sync.RWMutex

// ClipboardtextreceivedfuncCallback is a callback function for a 'ClipboardTextReceivedFunc' callback.
type ClipboardtextreceivedfuncCallback func(clipboard *Clipboard, text string)

// Unsupported : callback ColorSelectionChangePaletteFunc : unsupported parameter colors :

var callbackEntrycompletionmatchfuncId int
var callbackEntrycompletionmatchfuncMap = make(map[int]EntrycompletionmatchfuncCallback)
var callbackEntrycompletionmatchfuncLock sync.RWMutex

// EntrycompletionmatchfuncCallback is a callback function for a 'EntryCompletionMatchFunc' callback.
type EntrycompletionmatchfuncCallback func(completion *EntryCompletion, key string, iter *TreeIter) bool

var callbackFilefilterfuncId int
var callbackFilefilterfuncMap = make(map[int]FilefilterfuncCallback)
var callbackFilefilterfuncLock sync.RWMutex

// FilefilterfuncCallback is a callback function for a 'FileFilterFunc' callback.
type FilefilterfuncCallback func(filterInfo *FileFilterInfo) bool

var callbackFontfilterfuncId int
var callbackFontfilterfuncMap = make(map[int]FontfilterfuncCallback)
var callbackFontfilterfuncLock sync.RWMutex

// FontfilterfuncCallback is a callback function for a 'FontFilterFunc' callback.
type FontfilterfuncCallback func(family *pango.FontFamily, face *pango.FontFace) bool

var callbackIconviewforeachfuncId int
var callbackIconviewforeachfuncMap = make(map[int]IconviewforeachfuncCallback)
var callbackIconviewforeachfuncLock sync.RWMutex

// IconviewforeachfuncCallback is a callback function for a 'IconViewForeachFunc' callback.
type IconviewforeachfuncCallback func(iconView *IconView, path *TreePath)

// Unsupported : callback KeySnoopFunc : unsupported parameter func_data : no type generator for gpointer (gpointer) for param func_data

var callbackMenudetachfuncId int
var callbackMenudetachfuncMap = make(map[int]MenudetachfuncCallback)
var callbackMenudetachfuncLock sync.RWMutex

// MenudetachfuncCallback is a callback function for a 'MenuDetachFunc' callback.
type MenudetachfuncCallback func(attachWidget *Widget, menu *Menu)

var callbackMenupositionfuncId int
var callbackMenupositionfuncMap = make(map[int]MenupositionfuncCallback)
var callbackMenupositionfuncLock sync.RWMutex

// MenupositionfuncCallback is a callback function for a 'MenuPositionFunc' callback.
type MenupositionfuncCallback func(menu *Menu, x int32, y int32)

// Unsupported : callback ModuleInitFunc : unsupported ignore parameter argv

var callbackPagesetupdonefuncId int
var callbackPagesetupdonefuncMap = make(map[int]PagesetupdonefuncCallback)
var callbackPagesetupdonefuncLock sync.RWMutex

// PagesetupdonefuncCallback is a callback function for a 'PageSetupDoneFunc' callback.
type PagesetupdonefuncCallback func(pageSetup *PageSetup)

var callbackPrintsettingsfuncId int
var callbackPrintsettingsfuncMap = make(map[int]PrintsettingsfuncCallback)
var callbackPrintsettingsfuncLock sync.RWMutex

// PrintsettingsfuncCallback is a callback function for a 'PrintSettingsFunc' callback.
type PrintsettingsfuncCallback func(key string, value string)

var callbackRcpropertyparserId int
var callbackRcpropertyparserMap = make(map[int]RcpropertyparserCallback)
var callbackRcpropertyparserLock sync.RWMutex

// RcpropertyparserCallback is a callback function for a 'RcPropertyParser' callback.
type RcpropertyparserCallback func(pspec *gobject.ParamSpec, rcString *glib.String, propertyValue *gobject.Value) bool

var callbackRecentfilterfuncId int
var callbackRecentfilterfuncMap = make(map[int]RecentfilterfuncCallback)
var callbackRecentfilterfuncLock sync.RWMutex

// RecentfilterfuncCallback is a callback function for a 'RecentFilterFunc' callback.
type RecentfilterfuncCallback func(filterInfo *RecentFilterInfo) bool

var callbackRecentsortfuncId int
var callbackRecentsortfuncMap = make(map[int]RecentsortfuncCallback)
var callbackRecentsortfuncLock sync.RWMutex

// RecentsortfuncCallback is a callback function for a 'RecentSortFunc' callback.
type RecentsortfuncCallback func(a *RecentInfo, b *RecentInfo) int32

var callbackStylepropertyparserId int
var callbackStylepropertyparserMap = make(map[int]StylepropertyparserCallback)
var callbackStylepropertyparserLock sync.RWMutex

// StylepropertyparserCallback is a callback function for a 'StylePropertyParser' callback.
type StylepropertyparserCallback func(string_ string, value *gobject.Value) bool

var callbackTextbufferdeserializefuncId int
var callbackTextbufferdeserializefuncMap = make(map[int]TextbufferdeserializefuncCallback)
var callbackTextbufferdeserializefuncLock sync.RWMutex

// TextbufferdeserializefuncCallback is a callback function for a 'TextBufferDeserializeFunc' callback.
type TextbufferdeserializefuncCallback func(registerBuffer *TextBuffer, contentBuffer *TextBuffer, iter *TextIter, createTags bool) bool

var callbackTextbufferserializefuncId int
var callbackTextbufferserializefuncMap = make(map[int]TextbufferserializefuncCallback)
var callbackTextbufferserializefuncLock sync.RWMutex

// TextbufferserializefuncCallback is a callback function for a 'TextBufferSerializeFunc' callback.
type TextbufferserializefuncCallback func(registerBuffer *TextBuffer, contentBuffer *TextBuffer, start *TextIter, end *TextIter, length uint64) uint8

var callbackTextcharpredicateId int
var callbackTextcharpredicateMap = make(map[int]TextcharpredicateCallback)
var callbackTextcharpredicateLock sync.RWMutex

// TextcharpredicateCallback is a callback function for a 'TextCharPredicate' callback.
type TextcharpredicateCallback func(ch rune) bool

var callbackTexttagtableforeachId int
var callbackTexttagtableforeachMap = make(map[int]TexttagtableforeachCallback)
var callbackTexttagtableforeachLock sync.RWMutex

// TexttagtableforeachCallback is a callback function for a 'TextTagTableForeach' callback.
type TexttagtableforeachCallback func(tag *TextTag)

// Unsupported : callback TranslateFunc : unsupported parameter func_data : no type generator for gpointer (gpointer) for param func_data

var callbackTreecelldatafuncId int
var callbackTreecelldatafuncMap = make(map[int]TreecelldatafuncCallback)
var callbackTreecelldatafuncLock sync.RWMutex

// TreecelldatafuncCallback is a callback function for a 'TreeCellDataFunc' callback.
type TreecelldatafuncCallback func(treeColumn *TreeViewColumn, cell *CellRenderer, treeModel *TreeModel, iter *TreeIter)

var callbackTreedestroycountfuncId int
var callbackTreedestroycountfuncMap = make(map[int]TreedestroycountfuncCallback)
var callbackTreedestroycountfuncLock sync.RWMutex

// TreedestroycountfuncCallback is a callback function for a 'TreeDestroyCountFunc' callback.
type TreedestroycountfuncCallback func(treeView *TreeView, path *TreePath, children int32)

var callbackTreeitercomparefuncId int
var callbackTreeitercomparefuncMap = make(map[int]TreeitercomparefuncCallback)
var callbackTreeitercomparefuncLock sync.RWMutex

// TreeitercomparefuncCallback is a callback function for a 'TreeIterCompareFunc' callback.
type TreeitercomparefuncCallback func(model *TreeModel, a *TreeIter, b *TreeIter) int32

var callbackTreemodelfiltermodifyfuncId int
var callbackTreemodelfiltermodifyfuncMap = make(map[int]TreemodelfiltermodifyfuncCallback)
var callbackTreemodelfiltermodifyfuncLock sync.RWMutex

// TreemodelfiltermodifyfuncCallback is a callback function for a 'TreeModelFilterModifyFunc' callback.
type TreemodelfiltermodifyfuncCallback func(model *TreeModel, iter *TreeIter, column int32)

var callbackTreemodelfiltervisiblefuncId int
var callbackTreemodelfiltervisiblefuncMap = make(map[int]TreemodelfiltervisiblefuncCallback)
var callbackTreemodelfiltervisiblefuncLock sync.RWMutex

// TreemodelfiltervisiblefuncCallback is a callback function for a 'TreeModelFilterVisibleFunc' callback.
type TreemodelfiltervisiblefuncCallback func(model *TreeModel, iter *TreeIter) bool

var callbackTreemodelforeachfuncId int
var callbackTreemodelforeachfuncMap = make(map[int]TreemodelforeachfuncCallback)
var callbackTreemodelforeachfuncLock sync.RWMutex

// TreemodelforeachfuncCallback is a callback function for a 'TreeModelForeachFunc' callback.
type TreemodelforeachfuncCallback func(model *TreeModel, path *TreePath, iter *TreeIter) bool

var callbackTreeselectionforeachfuncId int
var callbackTreeselectionforeachfuncMap = make(map[int]TreeselectionforeachfuncCallback)
var callbackTreeselectionforeachfuncLock sync.RWMutex

// TreeselectionforeachfuncCallback is a callback function for a 'TreeSelectionForeachFunc' callback.
type TreeselectionforeachfuncCallback func(model *TreeModel, path *TreePath, iter *TreeIter)

var callbackTreeselectionfuncId int
var callbackTreeselectionfuncMap = make(map[int]TreeselectionfuncCallback)
var callbackTreeselectionfuncLock sync.RWMutex

// TreeselectionfuncCallback is a callback function for a 'TreeSelectionFunc' callback.
type TreeselectionfuncCallback func(selection *TreeSelection, model *TreeModel, path *TreePath, pathCurrentlySelected bool) bool

var callbackTreeviewcolumndropfuncId int
var callbackTreeviewcolumndropfuncMap = make(map[int]TreeviewcolumndropfuncCallback)
var callbackTreeviewcolumndropfuncLock sync.RWMutex

// TreeviewcolumndropfuncCallback is a callback function for a 'TreeViewColumnDropFunc' callback.
type TreeviewcolumndropfuncCallback func(treeView *TreeView, column *TreeViewColumn, prevColumn *TreeViewColumn, nextColumn *TreeViewColumn) bool

var callbackTreeviewmappingfuncId int
var callbackTreeviewmappingfuncMap = make(map[int]TreeviewmappingfuncCallback)
var callbackTreeviewmappingfuncLock sync.RWMutex

// TreeviewmappingfuncCallback is a callback function for a 'TreeViewMappingFunc' callback.
type TreeviewmappingfuncCallback func(treeView *TreeView, path *TreePath)

var callbackTreeviewrowseparatorfuncId int
var callbackTreeviewrowseparatorfuncMap = make(map[int]TreeviewrowseparatorfuncCallback)
var callbackTreeviewrowseparatorfuncLock sync.RWMutex

// TreeviewrowseparatorfuncCallback is a callback function for a 'TreeViewRowSeparatorFunc' callback.
type TreeviewrowseparatorfuncCallback func(model *TreeModel, iter *TreeIter) bool

// Unsupported : callback TreeViewSearchEqualFunc : unsupported parameter search_data : no type generator for gpointer (gpointer) for param search_data

var callbackTreeviewsearchpositionfuncId int
var callbackTreeviewsearchpositionfuncMap = make(map[int]TreeviewsearchpositionfuncCallback)
var callbackTreeviewsearchpositionfuncLock sync.RWMutex

// TreeviewsearchpositionfuncCallback is a callback function for a 'TreeViewSearchPositionFunc' callback.
type TreeviewsearchpositionfuncCallback func(treeView *TreeView, searchDialog *Widget)
