// This is a generated file - DO NOT EDIT

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gio "github.com/pekim/gobbi/lib/gio"
	call "github.com/pekim/gobbi/lib/internal/call"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Allocation is a representation of the C alias GtkAllocation.
type Allocation *gdk.Rectangle

// Stock is a representation of the C alias GtkStock.
type Stock string

type AccelFlags int

const (
	GTK_ACCEL_VISIBLE AccelFlags = 1
	GTK_ACCEL_LOCKED  AccelFlags = 2
	GTK_ACCEL_MASK    AccelFlags = 7
)

type AttachOptions int

const (
	GTK_EXPAND AttachOptions = 1
	GTK_SHRINK AttachOptions = 2
	GTK_FILL   AttachOptions = 4
)

type CalendarDisplayOptions int

const (
	GTK_CALENDAR_SHOW_HEADING      CalendarDisplayOptions = 1
	GTK_CALENDAR_SHOW_DAY_NAMES    CalendarDisplayOptions = 2
	GTK_CALENDAR_NO_MONTH_CHANGE   CalendarDisplayOptions = 4
	GTK_CALENDAR_SHOW_WEEK_NUMBERS CalendarDisplayOptions = 8
	GTK_CALENDAR_SHOW_DETAILS      CalendarDisplayOptions = 32
)

type CellRendererState int

const (
	GTK_CELL_RENDERER_SELECTED    CellRendererState = 1
	GTK_CELL_RENDERER_PRELIT      CellRendererState = 2
	GTK_CELL_RENDERER_INSENSITIVE CellRendererState = 4
	GTK_CELL_RENDERER_SORTED      CellRendererState = 8
	GTK_CELL_RENDERER_FOCUSED     CellRendererState = 16
	GTK_CELL_RENDERER_EXPANDABLE  CellRendererState = 32
	GTK_CELL_RENDERER_EXPANDED    CellRendererState = 64
)

type DebugFlag int

const (
	GTK_DEBUG_MISC           DebugFlag = 1
	GTK_DEBUG_PLUGSOCKET     DebugFlag = 2
	GTK_DEBUG_TEXT           DebugFlag = 4
	GTK_DEBUG_TREE           DebugFlag = 8
	GTK_DEBUG_UPDATES        DebugFlag = 16
	GTK_DEBUG_KEYBINDINGS    DebugFlag = 32
	GTK_DEBUG_MULTIHEAD      DebugFlag = 64
	GTK_DEBUG_MODULES        DebugFlag = 128
	GTK_DEBUG_GEOMETRY       DebugFlag = 256
	GTK_DEBUG_ICONTHEME      DebugFlag = 512
	GTK_DEBUG_PRINTING       DebugFlag = 1024
	GTK_DEBUG_BUILDER        DebugFlag = 2048
	GTK_DEBUG_SIZE_REQUEST   DebugFlag = 4096
	GTK_DEBUG_NO_CSS_CACHE   DebugFlag = 8192
	GTK_DEBUG_BASELINES      DebugFlag = 16384
	GTK_DEBUG_PIXEL_CACHE    DebugFlag = 32768
	GTK_DEBUG_NO_PIXEL_CACHE DebugFlag = 65536
	GTK_DEBUG_INTERACTIVE    DebugFlag = 131072
	GTK_DEBUG_TOUCHSCREEN    DebugFlag = 262144
	GTK_DEBUG_ACTIONS        DebugFlag = 524288
	GTK_DEBUG_RESIZE         DebugFlag = 1048576
	GTK_DEBUG_LAYOUT         DebugFlag = 2097152
)

type DestDefaults int

const (
	GTK_DEST_DEFAULT_MOTION    DestDefaults = 1
	GTK_DEST_DEFAULT_HIGHLIGHT DestDefaults = 2
	GTK_DEST_DEFAULT_DROP      DestDefaults = 4
	GTK_DEST_DEFAULT_ALL       DestDefaults = 7
)

type DialogFlags int

const (
	GTK_DIALOG_MODAL               DialogFlags = 1
	GTK_DIALOG_DESTROY_WITH_PARENT DialogFlags = 2
	GTK_DIALOG_USE_HEADER_BAR      DialogFlags = 4
)

type FileFilterFlags int

const (
	GTK_FILE_FILTER_FILENAME     FileFilterFlags = 1
	GTK_FILE_FILTER_URI          FileFilterFlags = 2
	GTK_FILE_FILTER_DISPLAY_NAME FileFilterFlags = 4
	GTK_FILE_FILTER_MIME_TYPE    FileFilterFlags = 8
)

type IconLookupFlags int

const (
	GTK_ICON_LOOKUP_NO_SVG           IconLookupFlags = 1
	GTK_ICON_LOOKUP_FORCE_SVG        IconLookupFlags = 2
	GTK_ICON_LOOKUP_USE_BUILTIN      IconLookupFlags = 4
	GTK_ICON_LOOKUP_GENERIC_FALLBACK IconLookupFlags = 8
	GTK_ICON_LOOKUP_FORCE_SIZE       IconLookupFlags = 16
	GTK_ICON_LOOKUP_FORCE_REGULAR    IconLookupFlags = 32
	GTK_ICON_LOOKUP_FORCE_SYMBOLIC   IconLookupFlags = 64
	GTK_ICON_LOOKUP_DIR_LTR          IconLookupFlags = 128
	GTK_ICON_LOOKUP_DIR_RTL          IconLookupFlags = 256
)

type JunctionSides int

const (
	GTK_JUNCTION_NONE               JunctionSides = 0
	GTK_JUNCTION_CORNER_TOPLEFT     JunctionSides = 1
	GTK_JUNCTION_CORNER_TOPRIGHT    JunctionSides = 2
	GTK_JUNCTION_CORNER_BOTTOMLEFT  JunctionSides = 4
	GTK_JUNCTION_CORNER_BOTTOMRIGHT JunctionSides = 8
	GTK_JUNCTION_TOP                JunctionSides = 3
	GTK_JUNCTION_BOTTOM             JunctionSides = 12
	GTK_JUNCTION_LEFT               JunctionSides = 5
	GTK_JUNCTION_RIGHT              JunctionSides = 10
)

type PlacesOpenFlags int

const (
	GTK_PLACES_OPEN_NORMAL     PlacesOpenFlags = 1
	GTK_PLACES_OPEN_NEW_TAB    PlacesOpenFlags = 2
	GTK_PLACES_OPEN_NEW_WINDOW PlacesOpenFlags = 4
)

type RcFlags int

const (
	GTK_RC_FG   RcFlags = 1
	GTK_RC_BG   RcFlags = 2
	GTK_RC_TEXT RcFlags = 4
	GTK_RC_BASE RcFlags = 8
)

type RecentFilterFlags int

const (
	GTK_RECENT_FILTER_URI          RecentFilterFlags = 1
	GTK_RECENT_FILTER_DISPLAY_NAME RecentFilterFlags = 2
	GTK_RECENT_FILTER_MIME_TYPE    RecentFilterFlags = 4
	GTK_RECENT_FILTER_APPLICATION  RecentFilterFlags = 8
	GTK_RECENT_FILTER_GROUP        RecentFilterFlags = 16
	GTK_RECENT_FILTER_AGE          RecentFilterFlags = 32
)

type RegionFlags int

const (
	GTK_REGION_EVEN   RegionFlags = 1
	GTK_REGION_ODD    RegionFlags = 2
	GTK_REGION_FIRST  RegionFlags = 4
	GTK_REGION_LAST   RegionFlags = 8
	GTK_REGION_ONLY   RegionFlags = 16
	GTK_REGION_SORTED RegionFlags = 32
)

type StateFlags int

const (
	GTK_STATE_FLAG_NORMAL       StateFlags = 0
	GTK_STATE_FLAG_ACTIVE       StateFlags = 1
	GTK_STATE_FLAG_PRELIGHT     StateFlags = 2
	GTK_STATE_FLAG_SELECTED     StateFlags = 4
	GTK_STATE_FLAG_INSENSITIVE  StateFlags = 8
	GTK_STATE_FLAG_INCONSISTENT StateFlags = 16
	GTK_STATE_FLAG_FOCUSED      StateFlags = 32
	GTK_STATE_FLAG_BACKDROP     StateFlags = 64
	GTK_STATE_FLAG_DIR_LTR      StateFlags = 128
	GTK_STATE_FLAG_DIR_RTL      StateFlags = 256
	GTK_STATE_FLAG_LINK         StateFlags = 512
	GTK_STATE_FLAG_VISITED      StateFlags = 1024
	GTK_STATE_FLAG_CHECKED      StateFlags = 2048
	GTK_STATE_FLAG_DROP_ACTIVE  StateFlags = 4096
)

type TargetFlags int

const (
	GTK_TARGET_SAME_APP     TargetFlags = 1
	GTK_TARGET_SAME_WIDGET  TargetFlags = 2
	GTK_TARGET_OTHER_APP    TargetFlags = 4
	GTK_TARGET_OTHER_WIDGET TargetFlags = 8
)

type TextSearchFlags int

const (
	GTK_TEXT_SEARCH_VISIBLE_ONLY     TextSearchFlags = 1
	GTK_TEXT_SEARCH_TEXT_ONLY        TextSearchFlags = 2
	GTK_TEXT_SEARCH_CASE_INSENSITIVE TextSearchFlags = 4
)

type ToolPaletteDragTargets int

const (
	GTK_TOOL_PALETTE_DRAG_ITEMS  ToolPaletteDragTargets = 1
	GTK_TOOL_PALETTE_DRAG_GROUPS ToolPaletteDragTargets = 2
)

type TreeModelFlags int

const (
	GTK_TREE_MODEL_ITERS_PERSIST TreeModelFlags = 1
	GTK_TREE_MODEL_LIST_ONLY     TreeModelFlags = 2
)

type UIManagerItemType int

const (
	GTK_UI_MANAGER_AUTO              UIManagerItemType = 0
	GTK_UI_MANAGER_MENUBAR           UIManagerItemType = 1
	GTK_UI_MANAGER_MENU              UIManagerItemType = 2
	GTK_UI_MANAGER_TOOLBAR           UIManagerItemType = 4
	GTK_UI_MANAGER_PLACEHOLDER       UIManagerItemType = 8
	GTK_UI_MANAGER_POPUP             UIManagerItemType = 16
	GTK_UI_MANAGER_MENUITEM          UIManagerItemType = 32
	GTK_UI_MANAGER_TOOLITEM          UIManagerItemType = 64
	GTK_UI_MANAGER_SEPARATOR         UIManagerItemType = 128
	GTK_UI_MANAGER_ACCELERATOR       UIManagerItemType = 256
	GTK_UI_MANAGER_POPUP_WITH_ACCELS UIManagerItemType = 512
)

// ImplementorIface returns the ImplementorIface interface implemented by AboutDialog
func (recv *AboutDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AboutDialog
func (recv *AboutDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_accel_group_new : return type :

// Unsupported : gtk_accel_label_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by AccelLabel
func (recv *AccelLabel) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AccelLabel
func (recv *AccelLabel) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Action
func (recv *Action) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ActionBar
func (recv *ActionBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ActionBar
func (recv *ActionBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ActionGroup
func (recv *ActionGroup) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_adjustment_new : return type :

// Unsupported : gtk_alignment_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Alignment
func (recv *Alignment) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Alignment
func (recv *Alignment) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by AppChooserButton
func (recv *AppChooserButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// AppChooser returns the AppChooser interface implemented by AppChooserButton
func (recv *AppChooserButton) AppChooser() *AppChooser {
	return AppChooserNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AppChooserButton
func (recv *AppChooserButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by AppChooserButton
func (recv *AppChooserButton) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by AppChooserButton
func (recv *AppChooserButton) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by AppChooserDialog
func (recv *AppChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// AppChooser returns the AppChooser interface implemented by AppChooserDialog
func (recv *AppChooserDialog) AppChooser() *AppChooser {
	return AppChooserNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AppChooserDialog
func (recv *AppChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by AppChooserWidget
func (recv *AppChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// AppChooser returns the AppChooser interface implemented by AppChooserWidget
func (recv *AppChooserWidget) AppChooser() *AppChooser {
	return AppChooserNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AppChooserWidget
func (recv *AppChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by AppChooserWidget
func (recv *AppChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ActionGroup returns the ActionGroup interface implemented by Application
func (recv *Application) ActionGroup() *gio.ActionGroup {
	return gio.ActionGroupNewFromC(recv.ToC())
}

// ActionMap returns the ActionMap interface implemented by Application
func (recv *Application) ActionMap() *gio.ActionMap {
	return gio.ActionMapNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ApplicationWindow
func (recv *ApplicationWindow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// ActionGroup returns the ActionGroup interface implemented by ApplicationWindow
func (recv *ApplicationWindow) ActionGroup() *gio.ActionGroup {
	return gio.ActionGroupNewFromC(recv.ToC())
}

// ActionMap returns the ActionMap interface implemented by ApplicationWindow
func (recv *ApplicationWindow) ActionMap() *gio.ActionMap {
	return gio.ActionMapNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ApplicationWindow
func (recv *ApplicationWindow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_arrow_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Arrow
func (recv *Arrow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Arrow
func (recv *Arrow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ArrowAccessible
func (recv *ArrowAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ArrowAccessible
func (recv *ArrowAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Unsupported : gtk_aspect_frame_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by AspectFrame
func (recv *AspectFrame) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by AspectFrame
func (recv *AspectFrame) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Assistant
func (recv *Assistant) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Assistant
func (recv *Assistant) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Bin
func (recv *Bin) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Bin
func (recv *Bin) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by BooleanCellAccessible
func (recv *BooleanCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by BooleanCellAccessible
func (recv *BooleanCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Box
func (recv *Box) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Box
func (recv *Box) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Box
func (recv *Box) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_button_new : return type :

// Unsupported : gtk_button_new_from_stock : return type :

// Unsupported : gtk_button_new_with_label : return type :

// Unsupported : gtk_button_new_with_mnemonic : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Button
func (recv *Button) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by Button
func (recv *Button) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by Button
func (recv *Button) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Button
func (recv *Button) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by ButtonAccessible
func (recv *ButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ButtonAccessible
func (recv *ButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ButtonAccessible
func (recv *ButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ButtonBox
func (recv *ButtonBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ButtonBox
func (recv *ButtonBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ButtonBox
func (recv *ButtonBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_calendar_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Calendar
func (recv *Calendar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Calendar
func (recv *Calendar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by CellAccessible
func (recv *CellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by CellAccessible
func (recv *CellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by CellArea
func (recv *CellArea) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by CellArea
func (recv *CellArea) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by CellAreaBox
func (recv *CellAreaBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by CellAreaBox
func (recv *CellAreaBox) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by CellAreaBox
func (recv *CellAreaBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_cell_renderer_pixbuf_new : return type :

// Orientable returns the Orientable interface implemented by CellRendererProgress
func (recv *CellRendererProgress) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_cell_renderer_text_new : return type :

// Unsupported : gtk_cell_renderer_toggle_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by CellView
func (recv *CellView) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by CellView
func (recv *CellView) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by CellView
func (recv *CellView) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by CellView
func (recv *CellView) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_check_button_new : return type :

// Unsupported : gtk_check_button_new_with_label : return type :

// Unsupported : gtk_check_button_new_with_mnemonic : return type :

// ImplementorIface returns the ImplementorIface interface implemented by CheckButton
func (recv *CheckButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by CheckButton
func (recv *CheckButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by CheckButton
func (recv *CheckButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by CheckButton
func (recv *CheckButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_check_menu_item_new : return type :

// Unsupported : gtk_check_menu_item_new_with_label : return type :

// Unsupported : gtk_check_menu_item_new_with_mnemonic : return type :

// ImplementorIface returns the ImplementorIface interface implemented by CheckMenuItem
func (recv *CheckMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by CheckMenuItem
func (recv *CheckMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by CheckMenuItem
func (recv *CheckMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by CheckMenuItem
func (recv *CheckMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by CheckMenuItemAccessible
func (recv *CheckMenuItemAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by CheckMenuItemAccessible
func (recv *CheckMenuItemAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by CheckMenuItemAccessible
func (recv *CheckMenuItemAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ColorButton
func (recv *ColorButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ColorButton
func (recv *ColorButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ColorButton
func (recv *ColorButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorButton
func (recv *ColorButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ColorChooser returns the ColorChooser interface implemented by ColorButton
func (recv *ColorButton) ColorChooser() *ColorChooser {
	return ColorChooserNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ColorChooserDialog
func (recv *ColorChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorChooserDialog
func (recv *ColorChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ColorChooser returns the ColorChooser interface implemented by ColorChooserDialog
func (recv *ColorChooserDialog) ColorChooser() *ColorChooser {
	return ColorChooserNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ColorChooserWidget
func (recv *ColorChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorChooserWidget
func (recv *ColorChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ColorChooser returns the ColorChooser interface implemented by ColorChooserWidget
func (recv *ColorChooserWidget) ColorChooser() *ColorChooser {
	return ColorChooserNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ColorChooserWidget
func (recv *ColorChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_color_selection_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by ColorSelection
func (recv *ColorSelection) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorSelection
func (recv *ColorSelection) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ColorSelection
func (recv *ColorSelection) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_color_selection_dialog_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by ColorSelectionDialog
func (recv *ColorSelectionDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ColorSelectionDialog
func (recv *ColorSelectionDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_combo_box_new_with_area : return type :

// Unsupported : gtk_combo_box_new_with_area_and_entry : return type :

// ImplementorIface returns the ImplementorIface interface implemented by ComboBox
func (recv *ComboBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ComboBox
func (recv *ComboBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by ComboBox
func (recv *ComboBox) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by ComboBox
func (recv *ComboBox) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by ComboBoxAccessible
func (recv *ComboBoxAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ComboBoxAccessible
func (recv *ComboBoxAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by ComboBoxAccessible
func (recv *ComboBoxAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ComboBoxText
func (recv *ComboBoxText) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ComboBoxText
func (recv *ComboBoxText) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by ComboBoxText
func (recv *ComboBoxText) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by ComboBoxText
func (recv *ComboBoxText) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Container
func (recv *Container) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Container
func (recv *Container) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ContainerAccessible
func (recv *ContainerAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Unsupported : gtk_container_cell_accessible_new : return type :

// Action returns the Action interface implemented by ContainerCellAccessible
func (recv *ContainerCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ContainerCellAccessible
func (recv *ContainerCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Unsupported : gtk_css_provider_new : return type :

// StyleProvider returns the StyleProvider interface implemented by CssProvider
func (recv *CssProvider) StyleProvider() *StyleProvider {
	return StyleProviderNewFromC(recv.ToC())
}

// Unsupported : gtk_dialog_new : return type :

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// ImplementorIface returns the ImplementorIface interface implemented by Dialog
func (recv *Dialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Dialog
func (recv *Dialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_drawing_area_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by DrawingArea
func (recv *DrawingArea) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by DrawingArea
func (recv *DrawingArea) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_entry_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Entry
func (recv *Entry) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Entry
func (recv *Entry) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by Entry
func (recv *Entry) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// Editable returns the Editable interface implemented by Entry
func (recv *Entry) Editable() *Editable {
	return EditableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by EntryAccessible
func (recv *EntryAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by EntryAccessible
func (recv *EntryAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// EditableText returns the EditableText interface implemented by EntryAccessible
func (recv *EntryAccessible) EditableText() *atk.EditableText {
	return atk.EditableTextNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by EntryAccessible
func (recv *EntryAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by EntryCompletion
func (recv *EntryCompletion) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by EntryCompletion
func (recv *EntryCompletion) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Unsupported : EntryIconAccessible : no CType

// Unsupported : gtk_event_box_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by EventBox
func (recv *EventBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by EventBox
func (recv *EventBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Expander
func (recv *Expander) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Expander
func (recv *Expander) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by ExpanderAccessible
func (recv *ExpanderAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ExpanderAccessible
func (recv *ExpanderAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by FileChooserButton
func (recv *FileChooserButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FileChooserButton
func (recv *FileChooserButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FileChooser returns the FileChooser interface implemented by FileChooserButton
func (recv *FileChooserButton) FileChooser() *FileChooser {
	return FileChooserNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FileChooserButton
func (recv *FileChooserButton) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by FileChooserDialog
func (recv *FileChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FileChooserDialog
func (recv *FileChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FileChooser returns the FileChooser interface implemented by FileChooserDialog
func (recv *FileChooserDialog) FileChooser() *FileChooser {
	return FileChooserNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by FileChooserWidget
func (recv *FileChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FileChooserWidget
func (recv *FileChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FileChooser returns the FileChooser interface implemented by FileChooserWidget
func (recv *FileChooserWidget) FileChooser() *FileChooser {
	return FileChooserNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FileChooserWidget
func (recv *FileChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FileFilter
func (recv *FileFilter) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_fixed_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Fixed
func (recv *Fixed) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Fixed
func (recv *Fixed) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by FlowBox
func (recv *FlowBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FlowBox
func (recv *FlowBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FlowBox
func (recv *FlowBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by FlowBoxAccessible
func (recv *FlowBoxAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by FlowBoxAccessible
func (recv *FlowBoxAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by FlowBoxChild
func (recv *FlowBoxChild) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FlowBoxChild
func (recv *FlowBoxChild) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by FlowBoxChildAccessible
func (recv *FlowBoxChildAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by FontButton
func (recv *FontButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by FontButton
func (recv *FontButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by FontButton
func (recv *FontButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontButton
func (recv *FontButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FontChooser returns the FontChooser interface implemented by FontButton
func (recv *FontButton) FontChooser() *FontChooser {
	return FontChooserNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by FontChooserDialog
func (recv *FontChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontChooserDialog
func (recv *FontChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FontChooser returns the FontChooser interface implemented by FontChooserDialog
func (recv *FontChooserDialog) FontChooser() *FontChooser {
	return FontChooserNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by FontChooserWidget
func (recv *FontChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontChooserWidget
func (recv *FontChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// FontChooser returns the FontChooser interface implemented by FontChooserWidget
func (recv *FontChooserWidget) FontChooser() *FontChooser {
	return FontChooserNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FontChooserWidget
func (recv *FontChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_font_selection_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by FontSelection
func (recv *FontSelection) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontSelection
func (recv *FontSelection) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by FontSelection
func (recv *FontSelection) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_font_selection_dialog_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by FontSelectionDialog
func (recv *FontSelectionDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by FontSelectionDialog
func (recv *FontSelectionDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_frame_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Frame
func (recv *Frame) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Frame
func (recv *Frame) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by FrameAccessible
func (recv *FrameAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Unsupported : gtk_grid_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Grid
func (recv *Grid) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Grid
func (recv *Grid) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Grid
func (recv *Grid) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_hbox_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by HBox
func (recv *HBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HBox
func (recv *HBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HBox
func (recv *HBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_hbutton_box_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by HButtonBox
func (recv *HButtonBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HButtonBox
func (recv *HButtonBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HButtonBox
func (recv *HButtonBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_hpaned_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by HPaned
func (recv *HPaned) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HPaned
func (recv *HPaned) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HPaned
func (recv *HPaned) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by HSV
func (recv *HSV) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HSV
func (recv *HSV) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_hscale_new : return type :

// Unsupported : gtk_hscale_new_with_range : return type :

// ImplementorIface returns the ImplementorIface interface implemented by HScale
func (recv *HScale) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HScale
func (recv *HScale) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HScale
func (recv *HScale) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_hscrollbar_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by HScrollbar
func (recv *HScrollbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HScrollbar
func (recv *HScrollbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HScrollbar
func (recv *HScrollbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_hseparator_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by HSeparator
func (recv *HSeparator) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HSeparator
func (recv *HSeparator) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by HSeparator
func (recv *HSeparator) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_handle_box_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by HandleBox
func (recv *HandleBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HandleBox
func (recv *HandleBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by HeaderBar
func (recv *HeaderBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by HeaderBar
func (recv *HeaderBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_im_context_simple_new : return type :

// Unsupported : gtk_im_multicontext_new : return type :

// Unsupported : gtk_icon_factory_new : return type :

// Buildable returns the Buildable interface implemented by IconFactory
func (recv *IconFactory) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by IconView
func (recv *IconView) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by IconView
func (recv *IconView) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by IconView
func (recv *IconView) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by IconView
func (recv *IconView) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by IconViewAccessible
func (recv *IconViewAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by IconViewAccessible
func (recv *IconViewAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// Unsupported : gtk_image_new : return type :

// Unsupported : gtk_image_new_from_animation : return type :

// Unsupported : gtk_image_new_from_file : return type :

// Unsupported : gtk_image_new_from_icon_set : return type :

// Unsupported : gtk_image_new_from_pixbuf : return type :

// Unsupported : gtk_image_new_from_stock : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Image
func (recv *Image) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Image
func (recv *Image) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ImageAccessible
func (recv *ImageAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ImageAccessible
func (recv *ImageAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by ImageCellAccessible
func (recv *ImageCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ImageCellAccessible
func (recv *ImageCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ImageCellAccessible
func (recv *ImageCellAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Unsupported : gtk_image_menu_item_new : return type :

// Unsupported : gtk_image_menu_item_new_from_stock : return type :

// Unsupported : gtk_image_menu_item_new_with_label : return type :

// Unsupported : gtk_image_menu_item_new_with_mnemonic : return type :

// ImplementorIface returns the ImplementorIface interface implemented by ImageMenuItem
func (recv *ImageMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ImageMenuItem
func (recv *ImageMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ImageMenuItem
func (recv *ImageMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ImageMenuItem
func (recv *ImageMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// ImplementorIface returns the ImplementorIface interface implemented by InfoBar
func (recv *InfoBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by InfoBar
func (recv *InfoBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by InfoBar
func (recv *InfoBar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_invisible_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Invisible
func (recv *Invisible) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Invisible
func (recv *Invisible) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_label_new : return type :

// Unsupported : gtk_label_new_with_mnemonic : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Label
func (recv *Label) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Label
func (recv *Label) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by LabelAccessible
func (recv *LabelAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Hypertext returns the Hypertext interface implemented by LabelAccessible
func (recv *LabelAccessible) Hypertext() *atk.Hypertext {
	return atk.HypertextNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by LabelAccessible
func (recv *LabelAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// Unsupported : gtk_layout_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Layout
func (recv *Layout) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Layout
func (recv *Layout) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by Layout
func (recv *Layout) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by LevelBar
func (recv *LevelBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by LevelBar
func (recv *LevelBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by LevelBar
func (recv *LevelBar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by LevelBarAccessible
func (recv *LevelBarAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by LevelBarAccessible
func (recv *LevelBarAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by LinkButton
func (recv *LinkButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by LinkButton
func (recv *LinkButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by LinkButton
func (recv *LinkButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by LinkButton
func (recv *LinkButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by LinkButtonAccessible
func (recv *LinkButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by LinkButtonAccessible
func (recv *LinkButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// HyperlinkImpl returns the HyperlinkImpl interface implemented by LinkButtonAccessible
func (recv *LinkButtonAccessible) HyperlinkImpl() *atk.HyperlinkImpl {
	return atk.HyperlinkImplNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by LinkButtonAccessible
func (recv *LinkButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ListBox
func (recv *ListBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ListBox
func (recv *ListBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ListBoxAccessible
func (recv *ListBoxAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by ListBoxAccessible
func (recv *ListBoxAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ListBoxRow
func (recv *ListBoxRow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ListBoxRow
func (recv *ListBoxRow) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ListBoxRow
func (recv *ListBoxRow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ListBoxRowAccessible
func (recv *ListBoxRowAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : return type :

// Buildable returns the Buildable interface implemented by ListStore
func (recv *ListStore) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// TreeDragDest returns the TreeDragDest interface implemented by ListStore
func (recv *ListStore) TreeDragDest() *TreeDragDest {
	return TreeDragDestNewFromC(recv.ToC())
}

// TreeDragSource returns the TreeDragSource interface implemented by ListStore
func (recv *ListStore) TreeDragSource() *TreeDragSource {
	return TreeDragSourceNewFromC(recv.ToC())
}

// TreeModel returns the TreeModel interface implemented by ListStore
func (recv *ListStore) TreeModel() *TreeModel {
	return TreeModelNewFromC(recv.ToC())
}

// TreeSortable returns the TreeSortable interface implemented by ListStore
func (recv *ListStore) TreeSortable() *TreeSortable {
	return TreeSortableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by LockButton
func (recv *LockButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by LockButton
func (recv *LockButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by LockButton
func (recv *LockButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by LockButton
func (recv *LockButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by LockButtonAccessible
func (recv *LockButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by LockButtonAccessible
func (recv *LockButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by LockButtonAccessible
func (recv *LockButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Unsupported : gtk_menu_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Menu
func (recv *Menu) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Menu
func (recv *Menu) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by MenuAccessible
func (recv *MenuAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by MenuAccessible
func (recv *MenuAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// Unsupported : gtk_menu_bar_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by MenuBar
func (recv *MenuBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuBar
func (recv *MenuBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by MenuButton
func (recv *MenuButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by MenuButton
func (recv *MenuButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by MenuButton
func (recv *MenuButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuButton
func (recv *MenuButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by MenuButtonAccessible
func (recv *MenuButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by MenuButtonAccessible
func (recv *MenuButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by MenuButtonAccessible
func (recv *MenuButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Unsupported : gtk_menu_item_new : return type :

// Unsupported : gtk_menu_item_new_with_label : return type :

// Unsupported : gtk_menu_item_new_with_mnemonic : return type :

// ImplementorIface returns the ImplementorIface interface implemented by MenuItem
func (recv *MenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by MenuItem
func (recv *MenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by MenuItem
func (recv *MenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuItem
func (recv *MenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by MenuItemAccessible
func (recv *MenuItemAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by MenuItemAccessible
func (recv *MenuItemAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by MenuItemAccessible
func (recv *MenuItemAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by MenuShell
func (recv *MenuShell) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuShell
func (recv *MenuShell) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by MenuShellAccessible
func (recv *MenuShellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by MenuShellAccessible
func (recv *MenuShellAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by MenuToolButton
func (recv *MenuToolButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by MenuToolButton
func (recv *MenuToolButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by MenuToolButton
func (recv *MenuToolButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MenuToolButton
func (recv *MenuToolButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_message_dialog_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by MessageDialog
func (recv *MessageDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by MessageDialog
func (recv *MessageDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Misc
func (recv *Misc) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Misc
func (recv *Misc) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ModelButton
func (recv *ModelButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ModelButton
func (recv *ModelButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ModelButton
func (recv *ModelButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ModelButton
func (recv *ModelButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_notebook_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Notebook
func (recv *Notebook) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Notebook
func (recv *Notebook) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by NotebookAccessible
func (recv *NotebookAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by NotebookAccessible
func (recv *NotebookAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// Unsupported : gtk_notebook_page_accessible_new : return type :

// Component returns the Component interface implemented by NotebookPageAccessible
func (recv *NotebookPageAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Icon returns the Icon interface implemented by NumerableIcon
func (recv *NumerableIcon) Icon() *gio.Icon {
	return gio.IconNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by OffscreenWindow
func (recv *OffscreenWindow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by OffscreenWindow
func (recv *OffscreenWindow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Overlay
func (recv *Overlay) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Overlay
func (recv *Overlay) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Paned
func (recv *Paned) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Paned
func (recv *Paned) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Paned
func (recv *Paned) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by PanedAccessible
func (recv *PanedAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by PanedAccessible
func (recv *PanedAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by PlacesSidebar
func (recv *PlacesSidebar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by PlacesSidebar
func (recv *PlacesSidebar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkPlug

// ImplementorIface returns the ImplementorIface interface implemented by Popover
func (recv *Popover) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Popover
func (recv *Popover) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by PopoverAccessible
func (recv *PopoverAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by PopoverMenu
func (recv *PopoverMenu) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by PopoverMenu
func (recv *PopoverMenu) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// PrintOperationPreview returns the PrintOperationPreview interface implemented by PrintOperation
func (recv *PrintOperation) PrintOperationPreview() *PrintOperationPreview {
	return PrintOperationPreviewNewFromC(recv.ToC())
}

// Unsupported : gtk_progress_bar_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by ProgressBar
func (recv *ProgressBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ProgressBar
func (recv *ProgressBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ProgressBar
func (recv *ProgressBar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ProgressBarAccessible
func (recv *ProgressBarAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by ProgressBarAccessible
func (recv *ProgressBarAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RadioAction
func (recv *RadioAction) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_radio_button_new : return type :

// Unsupported : gtk_radio_button_new_from_widget : return type :

// Unsupported : gtk_radio_button_new_with_label : return type :

// Unsupported : gtk_radio_button_new_with_label_from_widget : return type :

// Unsupported : gtk_radio_button_new_with_mnemonic : return type :

// Unsupported : gtk_radio_button_new_with_mnemonic_from_widget : return type :

// ImplementorIface returns the ImplementorIface interface implemented by RadioButton
func (recv *RadioButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by RadioButton
func (recv *RadioButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by RadioButton
func (recv *RadioButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RadioButton
func (recv *RadioButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by RadioButtonAccessible
func (recv *RadioButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by RadioButtonAccessible
func (recv *RadioButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by RadioButtonAccessible
func (recv *RadioButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Unsupported : gtk_radio_menu_item_new : return type :

// Unsupported : gtk_radio_menu_item_new_with_label : return type :

// Unsupported : gtk_radio_menu_item_new_with_mnemonic : return type :

// ImplementorIface returns the ImplementorIface interface implemented by RadioMenuItem
func (recv *RadioMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by RadioMenuItem
func (recv *RadioMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by RadioMenuItem
func (recv *RadioMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RadioMenuItem
func (recv *RadioMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by RadioMenuItemAccessible
func (recv *RadioMenuItemAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by RadioMenuItemAccessible
func (recv *RadioMenuItemAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by RadioMenuItemAccessible
func (recv *RadioMenuItemAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by RadioToolButton
func (recv *RadioToolButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by RadioToolButton
func (recv *RadioToolButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by RadioToolButton
func (recv *RadioToolButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RadioToolButton
func (recv *RadioToolButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Range
func (recv *Range) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Range
func (recv *Range) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Range
func (recv *Range) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by RangeAccessible
func (recv *RangeAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by RangeAccessible
func (recv *RangeAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// Unsupported : gtk_rc_style_new : return type :

// Buildable returns the Buildable interface implemented by RecentAction
func (recv *RecentAction) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RecentChooser returns the RecentChooser interface implemented by RecentAction
func (recv *RecentAction) RecentChooser() *RecentChooser {
	return RecentChooserNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by RecentChooserDialog
func (recv *RecentChooserDialog) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RecentChooserDialog
func (recv *RecentChooserDialog) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RecentChooser returns the RecentChooser interface implemented by RecentChooserDialog
func (recv *RecentChooserDialog) RecentChooser() *RecentChooser {
	return RecentChooserNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by RecentChooserMenu
func (recv *RecentChooserMenu) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by RecentChooserMenu
func (recv *RecentChooserMenu) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RecentChooserMenu
func (recv *RecentChooserMenu) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// RecentChooser returns the RecentChooser interface implemented by RecentChooserMenu
func (recv *RecentChooserMenu) RecentChooser() *RecentChooser {
	return RecentChooserNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by RecentChooserWidget
func (recv *RecentChooserWidget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RecentChooserWidget
func (recv *RecentChooserWidget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by RecentChooserWidget
func (recv *RecentChooserWidget) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// RecentChooser returns the RecentChooser interface implemented by RecentChooserWidget
func (recv *RecentChooserWidget) RecentChooser() *RecentChooser {
	return RecentChooserNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by RecentFilter
func (recv *RecentFilter) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_renderer_cell_accessible_new : return type :

// Action returns the Action interface implemented by RendererCellAccessible
func (recv *RendererCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by RendererCellAccessible
func (recv *RendererCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Revealer
func (recv *Revealer) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Revealer
func (recv *Revealer) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Scale
func (recv *Scale) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Scale
func (recv *Scale) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Scale
func (recv *Scale) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ScaleAccessible
func (recv *ScaleAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by ScaleAccessible
func (recv *ScaleAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ScaleButton
func (recv *ScaleButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ScaleButton
func (recv *ScaleButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ScaleButton
func (recv *ScaleButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ScaleButton
func (recv *ScaleButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ScaleButton
func (recv *ScaleButton) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by ScaleButtonAccessible
func (recv *ScaleButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ScaleButtonAccessible
func (recv *ScaleButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ScaleButtonAccessible
func (recv *ScaleButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by ScaleButtonAccessible
func (recv *ScaleButtonAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Scrollbar
func (recv *Scrollbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Scrollbar
func (recv *Scrollbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Scrollbar
func (recv *Scrollbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_scrolled_window_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by ScrolledWindow
func (recv *ScrolledWindow) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ScrolledWindow
func (recv *ScrolledWindow) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ScrolledWindowAccessible
func (recv *ScrolledWindowAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by SearchBar
func (recv *SearchBar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SearchBar
func (recv *SearchBar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by SearchEntry
func (recv *SearchEntry) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SearchEntry
func (recv *SearchEntry) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by SearchEntry
func (recv *SearchEntry) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// Editable returns the Editable interface implemented by SearchEntry
func (recv *SearchEntry) Editable() *Editable {
	return EditableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Separator
func (recv *Separator) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Separator
func (recv *Separator) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Separator
func (recv *Separator) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_separator_menu_item_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by SeparatorMenuItem
func (recv *SeparatorMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by SeparatorMenuItem
func (recv *SeparatorMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by SeparatorMenuItem
func (recv *SeparatorMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SeparatorMenuItem
func (recv *SeparatorMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by SeparatorToolItem
func (recv *SeparatorToolItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by SeparatorToolItem
func (recv *SeparatorToolItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SeparatorToolItem
func (recv *SeparatorToolItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// StyleProvider returns the StyleProvider interface implemented by Settings
func (recv *Settings) StyleProvider() *StyleProvider {
	return StyleProviderNewFromC(recv.ToC())
}

// Unsupported : gtk_size_group_new : return type :

// Buildable returns the Buildable interface implemented by SizeGroup
func (recv *SizeGroup) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkSocket

// Unsupported : gtk_spin_button_new : return type :

// Unsupported : gtk_spin_button_new_with_range : return type :

// ImplementorIface returns the ImplementorIface interface implemented by SpinButton
func (recv *SpinButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by SpinButton
func (recv *SpinButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellEditable returns the CellEditable interface implemented by SpinButton
func (recv *SpinButton) CellEditable() *CellEditable {
	return CellEditableNewFromC(recv.ToC())
}

// Editable returns the Editable interface implemented by SpinButton
func (recv *SpinButton) Editable() *Editable {
	return EditableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by SpinButton
func (recv *SpinButton) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// EditableText returns the EditableText interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) EditableText() *atk.EditableText {
	return atk.EditableTextNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// Value returns the Value interface implemented by SpinButtonAccessible
func (recv *SpinButtonAccessible) Value() *atk.Value {
	return atk.ValueNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Spinner
func (recv *Spinner) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Spinner
func (recv *Spinner) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by SpinnerAccessible
func (recv *SpinnerAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by SpinnerAccessible
func (recv *SpinnerAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Stack
func (recv *Stack) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Stack
func (recv *Stack) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Blacklisted : GtkStackAccessible

// ImplementorIface returns the ImplementorIface interface implemented by StackSidebar
func (recv *StackSidebar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by StackSidebar
func (recv *StackSidebar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by StackSwitcher
func (recv *StackSwitcher) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by StackSwitcher
func (recv *StackSwitcher) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by StackSwitcher
func (recv *StackSwitcher) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_statusbar_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Statusbar
func (recv *Statusbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Statusbar
func (recv *Statusbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Statusbar
func (recv *Statusbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by StatusbarAccessible
func (recv *StatusbarAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Unsupported : gtk_style_new : return type :

// Unsupported : gtk_style_context_new : return type :

// Unsupported : gtk_style_properties_new : return type :

// StyleProvider returns the StyleProvider interface implemented by StyleProperties
func (recv *StyleProperties) StyleProvider() *StyleProvider {
	return StyleProviderNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by Switch
func (recv *Switch) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by Switch
func (recv *Switch) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by Switch
func (recv *Switch) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Switch
func (recv *Switch) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by SwitchAccessible
func (recv *SwitchAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by SwitchAccessible
func (recv *SwitchAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Unsupported : gtk_table_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Table
func (recv *Table) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Table
func (recv *Table) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_tearoff_menu_item_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by TearoffMenuItem
func (recv *TearoffMenuItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by TearoffMenuItem
func (recv *TearoffMenuItem) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by TearoffMenuItem
func (recv *TearoffMenuItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by TearoffMenuItem
func (recv *TearoffMenuItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_text_buffer_new : return type :

// Action returns the Action interface implemented by TextCellAccessible
func (recv *TextCellAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by TextCellAccessible
func (recv *TextCellAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by TextCellAccessible
func (recv *TextCellAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// Unsupported : gtk_text_child_anchor_new : return type :

// Unsupported : gtk_text_tag_new : return type :

// Unsupported : gtk_text_tag_table_new : return type :

// Buildable returns the Buildable interface implemented by TextTagTable
func (recv *TextTagTable) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_text_view_new : return type :

// Unsupported : gtk_text_view_new_with_buffer : return type :

// ImplementorIface returns the ImplementorIface interface implemented by TextView
func (recv *TextView) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by TextView
func (recv *TextView) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by TextView
func (recv *TextView) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by TextViewAccessible
func (recv *TextViewAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// EditableText returns the EditableText interface implemented by TextViewAccessible
func (recv *TextViewAccessible) EditableText() *atk.EditableText {
	return atk.EditableTextNewFromC(recv.ToC())
}

// StreamableContent returns the StreamableContent interface implemented by TextViewAccessible
func (recv *TextViewAccessible) StreamableContent() *atk.StreamableContent {
	return atk.StreamableContentNewFromC(recv.ToC())
}

// Text returns the Text interface implemented by TextViewAccessible
func (recv *TextViewAccessible) Text() *atk.Text {
	return atk.TextNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToggleAction
func (recv *ToggleAction) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_toggle_button_new : return type :

// Unsupported : gtk_toggle_button_new_with_label : return type :

// Unsupported : gtk_toggle_button_new_with_mnemonic : return type :

// ImplementorIface returns the ImplementorIface interface implemented by ToggleButton
func (recv *ToggleButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ToggleButton
func (recv *ToggleButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ToggleButton
func (recv *ToggleButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToggleButton
func (recv *ToggleButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Action returns the Action interface implemented by ToggleButtonAccessible
func (recv *ToggleButtonAccessible) Action() *atk.Action {
	return atk.ActionNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by ToggleButtonAccessible
func (recv *ToggleButtonAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Image returns the Image interface implemented by ToggleButtonAccessible
func (recv *ToggleButtonAccessible) Image() *atk.Image {
	return atk.ImageNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ToggleToolButton
func (recv *ToggleToolButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ToggleToolButton
func (recv *ToggleToolButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ToggleToolButton
func (recv *ToggleToolButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToggleToolButton
func (recv *ToggleToolButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ToolButton
func (recv *ToolButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by ToolButton
func (recv *ToolButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ToolButton
func (recv *ToolButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToolButton
func (recv *ToolButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ToolItem
func (recv *ToolItem) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by ToolItem
func (recv *ToolItem) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToolItem
func (recv *ToolItem) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ToolItemGroup
func (recv *ToolItemGroup) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToolItemGroup
func (recv *ToolItemGroup) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// ToolShell returns the ToolShell interface implemented by ToolItemGroup
func (recv *ToolItemGroup) ToolShell() *ToolShell {
	return ToolShellNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by ToolPalette
func (recv *ToolPalette) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by ToolPalette
func (recv *ToolPalette) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by ToolPalette
func (recv *ToolPalette) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by ToolPalette
func (recv *ToolPalette) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// Unsupported : gtk_toolbar_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Toolbar
func (recv *Toolbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Toolbar
func (recv *Toolbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by Toolbar
func (recv *Toolbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// ToolShell returns the ToolShell interface implemented by Toolbar
func (recv *Toolbar) ToolShell() *ToolShell {
	return ToolShellNewFromC(recv.ToC())
}

// TreeDragSource returns the TreeDragSource interface implemented by TreeModelFilter
func (recv *TreeModelFilter) TreeDragSource() *TreeDragSource {
	return TreeDragSourceNewFromC(recv.ToC())
}

// TreeModel returns the TreeModel interface implemented by TreeModelFilter
func (recv *TreeModelFilter) TreeModel() *TreeModel {
	return TreeModelNewFromC(recv.ToC())
}

// TreeDragSource returns the TreeDragSource interface implemented by TreeModelSort
func (recv *TreeModelSort) TreeDragSource() *TreeDragSource {
	return TreeDragSourceNewFromC(recv.ToC())
}

// TreeModel returns the TreeModel interface implemented by TreeModelSort
func (recv *TreeModelSort) TreeModel() *TreeModel {
	return TreeModelNewFromC(recv.ToC())
}

// TreeSortable returns the TreeSortable interface implemented by TreeModelSort
func (recv *TreeModelSort) TreeSortable() *TreeSortable {
	return TreeSortableNewFromC(recv.ToC())
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : return type :

// Buildable returns the Buildable interface implemented by TreeStore
func (recv *TreeStore) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// TreeDragDest returns the TreeDragDest interface implemented by TreeStore
func (recv *TreeStore) TreeDragDest() *TreeDragDest {
	return TreeDragDestNewFromC(recv.ToC())
}

// TreeDragSource returns the TreeDragSource interface implemented by TreeStore
func (recv *TreeStore) TreeDragSource() *TreeDragSource {
	return TreeDragSourceNewFromC(recv.ToC())
}

// TreeModel returns the TreeModel interface implemented by TreeStore
func (recv *TreeStore) TreeModel() *TreeModel {
	return TreeModelNewFromC(recv.ToC())
}

// TreeSortable returns the TreeSortable interface implemented by TreeStore
func (recv *TreeStore) TreeSortable() *TreeSortable {
	return TreeSortableNewFromC(recv.ToC())
}

// Unsupported : gtk_tree_view_new : return type :

// Unsupported : gtk_tree_view_new_with_model : return type :

// ImplementorIface returns the ImplementorIface interface implemented by TreeView
func (recv *TreeView) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by TreeView
func (recv *TreeView) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by TreeView
func (recv *TreeView) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by TreeViewAccessible
func (recv *TreeViewAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Selection returns the Selection interface implemented by TreeViewAccessible
func (recv *TreeViewAccessible) Selection() *atk.Selection {
	return atk.SelectionNewFromC(recv.ToC())
}

// Table returns the Table interface implemented by TreeViewAccessible
func (recv *TreeViewAccessible) Table() *atk.Table {
	return atk.TableNewFromC(recv.ToC())
}

// CellAccessibleParent returns the CellAccessibleParent interface implemented by TreeViewAccessible
func (recv *TreeViewAccessible) CellAccessibleParent() *CellAccessibleParent {
	return CellAccessibleParentNewFromC(recv.ToC())
}

// Unsupported : gtk_tree_view_column_new : return type :

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// Buildable returns the Buildable interface implemented by TreeViewColumn
func (recv *TreeViewColumn) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// CellLayout returns the CellLayout interface implemented by TreeViewColumn
func (recv *TreeViewColumn) CellLayout() *CellLayout {
	return CellLayoutNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by UIManager
func (recv *UIManager) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Unsupported : gtk_vbox_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by VBox
func (recv *VBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VBox
func (recv *VBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VBox
func (recv *VBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_vbutton_box_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by VButtonBox
func (recv *VButtonBox) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VButtonBox
func (recv *VButtonBox) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VButtonBox
func (recv *VButtonBox) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_vpaned_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by VPaned
func (recv *VPaned) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VPaned
func (recv *VPaned) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VPaned
func (recv *VPaned) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_vscale_new : return type :

// Unsupported : gtk_vscale_new_with_range : return type :

// ImplementorIface returns the ImplementorIface interface implemented by VScale
func (recv *VScale) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VScale
func (recv *VScale) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VScale
func (recv *VScale) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_vscrollbar_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by VScrollbar
func (recv *VScrollbar) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VScrollbar
func (recv *VScrollbar) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VScrollbar
func (recv *VScrollbar) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_vseparator_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by VSeparator
func (recv *VSeparator) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VSeparator
func (recv *VSeparator) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VSeparator
func (recv *VSeparator) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_viewport_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Viewport
func (recv *Viewport) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Viewport
func (recv *Viewport) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Scrollable returns the Scrollable interface implemented by Viewport
func (recv *Viewport) Scrollable() *Scrollable {
	return ScrollableNewFromC(recv.ToC())
}

// ImplementorIface returns the ImplementorIface interface implemented by VolumeButton
func (recv *VolumeButton) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Actionable returns the Actionable interface implemented by VolumeButton
func (recv *VolumeButton) Actionable() *Actionable {
	return ActionableNewFromC(recv.ToC())
}

// Activatable returns the Activatable interface implemented by VolumeButton
func (recv *VolumeButton) Activatable() *Activatable {
	return ActivatableNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by VolumeButton
func (recv *VolumeButton) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Orientable returns the Orientable interface implemented by VolumeButton
func (recv *VolumeButton) Orientable() *Orientable {
	return OrientableNewFromC(recv.ToC())
}

// Unsupported : gtk_widget_new : unsupported parameter ... : varargs

// ImplementorIface returns the ImplementorIface interface implemented by Widget
func (recv *Widget) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Widget
func (recv *Widget) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by WidgetAccessible
func (recv *WidgetAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Unsupported : gtk_window_new : return type :

// ImplementorIface returns the ImplementorIface interface implemented by Window
func (recv *Window) ImplementorIface() *atk.ImplementorIface {
	return atk.ImplementorIfaceNewFromC(recv.ToC())
}

// Buildable returns the Buildable interface implemented by Window
func (recv *Window) Buildable() *Buildable {
	return BuildableNewFromC(recv.ToC())
}

// Component returns the Component interface implemented by WindowAccessible
func (recv *WindowAccessible) Component() *atk.Component {
	return atk.ComponentNewFromC(recv.ToC())
}

// Window returns the Window interface implemented by WindowAccessible
func (recv *WindowAccessible) Window() *atk.Window {
	return atk.WindowNewFromC(recv.ToC())
}

// Unsupported : gtk_window_group_new : return type :

const BINARY_AGE int32 = 2230
const INPUT_ERROR int32 = -1
const INTERFACE_AGE int32 = 30
const MAJOR_VERSION int32 = 3
const MAX_COMPOSE_LEN int32 = 7
const MICRO_VERSION int32 = 30
const MINOR_VERSION int32 = 22
const PAPER_NAME_A3 string = "iso_a3"
const PAPER_NAME_A4 string = "iso_a4"
const PAPER_NAME_A5 string = "iso_a5"
const PAPER_NAME_B5 string = "iso_b5"
const PAPER_NAME_EXECUTIVE string = "na_executive"
const PAPER_NAME_LEGAL string = "na_legal"
const PAPER_NAME_LETTER string = "na_letter"
const PATH_PRIO_MASK int32 = 15
const PRINT_SETTINGS_COLLATE string = "collate"
const PRINT_SETTINGS_DEFAULT_SOURCE string = "default-source"
const PRINT_SETTINGS_DITHER string = "dither"
const PRINT_SETTINGS_DUPLEX string = "duplex"
const PRINT_SETTINGS_FINISHINGS string = "finishings"
const PRINT_SETTINGS_MEDIA_TYPE string = "media-type"
const PRINT_SETTINGS_NUMBER_UP string = "number-up"
const PRINT_SETTINGS_NUMBER_UP_LAYOUT string = "number-up-layout"
const PRINT_SETTINGS_N_COPIES string = "n-copies"
const PRINT_SETTINGS_ORIENTATION string = "orientation"
const PRINT_SETTINGS_OUTPUT_BIN string = "output-bin"
const PRINT_SETTINGS_OUTPUT_FILE_FORMAT string = "output-file-format"
const PRINT_SETTINGS_OUTPUT_URI string = "output-uri"
const PRINT_SETTINGS_PAGE_RANGES string = "page-ranges"
const PRINT_SETTINGS_PAGE_SET string = "page-set"
const PRINT_SETTINGS_PAPER_FORMAT string = "paper-format"
const PRINT_SETTINGS_PAPER_HEIGHT string = "paper-height"
const PRINT_SETTINGS_PAPER_WIDTH string = "paper-width"
const PRINT_SETTINGS_PRINTER string = "printer"
const PRINT_SETTINGS_PRINTER_LPI string = "printer-lpi"
const PRINT_SETTINGS_PRINT_PAGES string = "print-pages"
const PRINT_SETTINGS_QUALITY string = "quality"
const PRINT_SETTINGS_RESOLUTION string = "resolution"
const PRINT_SETTINGS_RESOLUTION_X string = "resolution-x"
const PRINT_SETTINGS_RESOLUTION_Y string = "resolution-y"
const PRINT_SETTINGS_REVERSE string = "reverse"
const PRINT_SETTINGS_SCALE string = "scale"
const PRINT_SETTINGS_USE_COLOR string = "use-color"
const PRINT_SETTINGS_WIN32_DRIVER_EXTRA string = "win32-driver-extra"
const PRINT_SETTINGS_WIN32_DRIVER_VERSION string = "win32-driver-version"
const PRIORITY_RESIZE int32 = 10

// Blacklisted : STOCK_ADD

// Blacklisted : STOCK_APPLY

// Blacklisted : STOCK_BOLD

// Blacklisted : STOCK_CANCEL

// Blacklisted : STOCK_CDROM

// Blacklisted : STOCK_CLEAR

// Blacklisted : STOCK_CLOSE

// Blacklisted : STOCK_CONVERT

// Blacklisted : STOCK_COPY

// Blacklisted : STOCK_CUT

// Blacklisted : STOCK_DELETE

// Blacklisted : STOCK_DIALOG_ERROR

// Blacklisted : STOCK_DIALOG_INFO

// Blacklisted : STOCK_DIALOG_QUESTION

// Blacklisted : STOCK_DIALOG_WARNING

// Blacklisted : STOCK_DND

// Blacklisted : STOCK_DND_MULTIPLE

// Blacklisted : STOCK_EXECUTE

// Blacklisted : STOCK_FIND

// Blacklisted : STOCK_FIND_AND_REPLACE

// Blacklisted : STOCK_FLOPPY

// Blacklisted : STOCK_GOTO_BOTTOM

// Blacklisted : STOCK_GOTO_FIRST

// Blacklisted : STOCK_GOTO_LAST

// Blacklisted : STOCK_GOTO_TOP

// Blacklisted : STOCK_GO_BACK

// Blacklisted : STOCK_GO_DOWN

// Blacklisted : STOCK_GO_FORWARD

// Blacklisted : STOCK_GO_UP

// Blacklisted : STOCK_HELP

// Blacklisted : STOCK_HOME

// Blacklisted : STOCK_INDEX

// Blacklisted : STOCK_ITALIC

// Blacklisted : STOCK_JUMP_TO

// Blacklisted : STOCK_JUSTIFY_CENTER

// Blacklisted : STOCK_JUSTIFY_FILL

// Blacklisted : STOCK_JUSTIFY_LEFT

// Blacklisted : STOCK_JUSTIFY_RIGHT

// Blacklisted : STOCK_MISSING_IMAGE

// Blacklisted : STOCK_NEW

// Blacklisted : STOCK_NO

// Blacklisted : STOCK_OK

// Blacklisted : STOCK_OPEN

// Blacklisted : STOCK_PASTE

// Blacklisted : STOCK_PREFERENCES

// Blacklisted : STOCK_PRINT

// Blacklisted : STOCK_PRINT_PREVIEW

// Blacklisted : STOCK_PROPERTIES

// Blacklisted : STOCK_QUIT

// Blacklisted : STOCK_REDO

// Blacklisted : STOCK_REFRESH

// Blacklisted : STOCK_REMOVE

// Blacklisted : STOCK_REVERT_TO_SAVED

// Blacklisted : STOCK_SAVE

// Blacklisted : STOCK_SAVE_AS

// Blacklisted : STOCK_SELECT_COLOR

// Blacklisted : STOCK_SELECT_FONT

// Blacklisted : STOCK_SORT_ASCENDING

// Blacklisted : STOCK_SORT_DESCENDING

// Blacklisted : STOCK_SPELL_CHECK

// Blacklisted : STOCK_STOP

// Blacklisted : STOCK_STRIKETHROUGH

// Blacklisted : STOCK_UNDELETE

// Blacklisted : STOCK_UNDERLINE

// Blacklisted : STOCK_UNDO

// Blacklisted : STOCK_YES

// Blacklisted : STOCK_ZOOM_100

// Blacklisted : STOCK_ZOOM_FIT

// Blacklisted : STOCK_ZOOM_IN

// Blacklisted : STOCK_ZOOM_OUT

const STYLE_CLASS_ACCELERATOR string = "accelerator"
const STYLE_CLASS_ARROW string = "arrow"
const STYLE_CLASS_BACKGROUND string = "background"
const STYLE_CLASS_BOTTOM string = "bottom"
const STYLE_CLASS_BUTTON string = "button"
const STYLE_CLASS_CALENDAR string = "calendar"
const STYLE_CLASS_CELL string = "cell"
const STYLE_CLASS_CHECK string = "check"
const STYLE_CLASS_COMBOBOX_ENTRY string = "combobox-entry"
const STYLE_CLASS_CONTEXT_MENU string = "context-menu"
const STYLE_CLASS_CURSOR_HANDLE string = "cursor-handle"
const STYLE_CLASS_DEFAULT string = "default"
const STYLE_CLASS_DIM_LABEL string = "dim-label"
const STYLE_CLASS_DND string = "dnd"
const STYLE_CLASS_DOCK string = "dock"
const STYLE_CLASS_ENTRY string = "entry"
const STYLE_CLASS_ERROR string = "error"
const STYLE_CLASS_EXPANDER string = "expander"
const STYLE_CLASS_FRAME string = "frame"
const STYLE_CLASS_GRIP string = "grip"
const STYLE_CLASS_HEADER string = "header"
const STYLE_CLASS_HIGHLIGHT string = "highlight"
const STYLE_CLASS_HORIZONTAL string = "horizontal"
const STYLE_CLASS_IMAGE string = "image"
const STYLE_CLASS_INFO string = "info"
const STYLE_CLASS_INLINE_TOOLBAR string = "inline-toolbar"
const STYLE_CLASS_INSERTION_CURSOR string = "insertion-cursor"
const STYLE_CLASS_LEFT string = "left"
const STYLE_CLASS_LEVEL_BAR string = "level-bar"
const STYLE_CLASS_LINKED string = "linked"
const STYLE_CLASS_LIST string = "list"
const STYLE_CLASS_LIST_ROW string = "list-row"
const STYLE_CLASS_MARK string = "mark"
const STYLE_CLASS_MENU string = "menu"
const STYLE_CLASS_MENUBAR string = "menubar"
const STYLE_CLASS_MENUITEM string = "menuitem"
const STYLE_CLASS_NOTEBOOK string = "notebook"
const STYLE_CLASS_OSD string = "osd"
const STYLE_CLASS_PANE_SEPARATOR string = "pane-separator"
const STYLE_CLASS_PRIMARY_TOOLBAR string = "primary-toolbar"
const STYLE_CLASS_PROGRESSBAR string = "progressbar"
const STYLE_CLASS_PULSE string = "pulse"
const STYLE_CLASS_QUESTION string = "question"
const STYLE_CLASS_RADIO string = "radio"
const STYLE_CLASS_RAISED string = "raised"
const STYLE_CLASS_READ_ONLY string = "read-only"
const STYLE_CLASS_RIGHT string = "right"
const STYLE_CLASS_RUBBERBAND string = "rubberband"
const STYLE_CLASS_SCALE string = "scale"
const STYLE_CLASS_SCALE_HAS_MARKS_ABOVE string = "scale-has-marks-above"
const STYLE_CLASS_SCALE_HAS_MARKS_BELOW string = "scale-has-marks-below"
const STYLE_CLASS_SCROLLBAR string = "scrollbar"
const STYLE_CLASS_SCROLLBARS_JUNCTION string = "scrollbars-junction"
const STYLE_CLASS_SEPARATOR string = "separator"
const STYLE_CLASS_SIDEBAR string = "sidebar"
const STYLE_CLASS_SLIDER string = "slider"
const STYLE_CLASS_SPINBUTTON string = "spinbutton"
const STYLE_CLASS_SPINNER string = "spinner"
const STYLE_CLASS_TITLEBAR string = "titlebar"
const STYLE_CLASS_TOOLBAR string = "toolbar"
const STYLE_CLASS_TOOLTIP string = "tooltip"
const STYLE_CLASS_TOP string = "top"
const STYLE_CLASS_TROUGH string = "trough"
const STYLE_CLASS_VERTICAL string = "vertical"
const STYLE_CLASS_VIEW string = "view"
const STYLE_CLASS_WARNING string = "warning"
const STYLE_PROPERTY_BACKGROUND_COLOR string = "background-color"
const STYLE_PROPERTY_BACKGROUND_IMAGE string = "background-image"
const STYLE_PROPERTY_BORDER_COLOR string = "border-color"
const STYLE_PROPERTY_BORDER_RADIUS string = "border-radius"
const STYLE_PROPERTY_BORDER_STYLE string = "border-style"
const STYLE_PROPERTY_BORDER_WIDTH string = "border-width"
const STYLE_PROPERTY_COLOR string = "color"
const STYLE_PROPERTY_FONT string = "font"
const STYLE_PROPERTY_MARGIN string = "margin"
const STYLE_PROPERTY_PADDING string = "padding"
const STYLE_PROVIDER_PRIORITY_APPLICATION int32 = 600
const STYLE_PROVIDER_PRIORITY_FALLBACK int32 = 1
const STYLE_PROVIDER_PRIORITY_SETTINGS int32 = 400
const STYLE_PROVIDER_PRIORITY_THEME int32 = 200
const STYLE_PROVIDER_PRIORITY_USER int32 = 800
const STYLE_REGION_COLUMN string = "column"
const STYLE_REGION_COLUMN_HEADER string = "column-header"
const STYLE_REGION_ROW string = "row"
const STYLE_REGION_TAB string = "tab"
const TEXT_VIEW_PRIORITY_VALIDATE int32 = 5
const TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID int32 = -1
const TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID int32 = -2

type Align int

const (
	GTK_ALIGN_FILL     Align = 0
	GTK_ALIGN_START    Align = 1
	GTK_ALIGN_END      Align = 2
	GTK_ALIGN_CENTER   Align = 3
	GTK_ALIGN_BASELINE Align = 4
)

type ArrowPlacement int

const (
	GTK_ARROWS_BOTH  ArrowPlacement = 0
	GTK_ARROWS_START ArrowPlacement = 1
	GTK_ARROWS_END   ArrowPlacement = 2
)

type ArrowType int

const (
	GTK_ARROW_UP    ArrowType = 0
	GTK_ARROW_DOWN  ArrowType = 1
	GTK_ARROW_LEFT  ArrowType = 2
	GTK_ARROW_RIGHT ArrowType = 3
	GTK_ARROW_NONE  ArrowType = 4
)

type AssistantPageType int

const (
	GTK_ASSISTANT_PAGE_CONTENT  AssistantPageType = 0
	GTK_ASSISTANT_PAGE_INTRO    AssistantPageType = 1
	GTK_ASSISTANT_PAGE_CONFIRM  AssistantPageType = 2
	GTK_ASSISTANT_PAGE_SUMMARY  AssistantPageType = 3
	GTK_ASSISTANT_PAGE_PROGRESS AssistantPageType = 4
	GTK_ASSISTANT_PAGE_CUSTOM   AssistantPageType = 5
)

type BorderStyle int

const (
	GTK_BORDER_STYLE_NONE   BorderStyle = 0
	GTK_BORDER_STYLE_SOLID  BorderStyle = 1
	GTK_BORDER_STYLE_INSET  BorderStyle = 2
	GTK_BORDER_STYLE_OUTSET BorderStyle = 3
	GTK_BORDER_STYLE_HIDDEN BorderStyle = 4
	GTK_BORDER_STYLE_DOTTED BorderStyle = 5
	GTK_BORDER_STYLE_DASHED BorderStyle = 6
	GTK_BORDER_STYLE_DOUBLE BorderStyle = 7
	GTK_BORDER_STYLE_GROOVE BorderStyle = 8
	GTK_BORDER_STYLE_RIDGE  BorderStyle = 9
)

type BuilderError int

const (
	GTK_BUILDER_ERROR_INVALID_TYPE_FUNCTION  BuilderError = 0
	GTK_BUILDER_ERROR_UNHANDLED_TAG          BuilderError = 1
	GTK_BUILDER_ERROR_MISSING_ATTRIBUTE      BuilderError = 2
	GTK_BUILDER_ERROR_INVALID_ATTRIBUTE      BuilderError = 3
	GTK_BUILDER_ERROR_INVALID_TAG            BuilderError = 4
	GTK_BUILDER_ERROR_MISSING_PROPERTY_VALUE BuilderError = 5
	GTK_BUILDER_ERROR_INVALID_VALUE          BuilderError = 6
	GTK_BUILDER_ERROR_VERSION_MISMATCH       BuilderError = 7
	GTK_BUILDER_ERROR_DUPLICATE_ID           BuilderError = 8
	GTK_BUILDER_ERROR_OBJECT_TYPE_REFUSED    BuilderError = 9
	GTK_BUILDER_ERROR_TEMPLATE_MISMATCH      BuilderError = 10
	GTK_BUILDER_ERROR_INVALID_PROPERTY       BuilderError = 11
	GTK_BUILDER_ERROR_INVALID_SIGNAL         BuilderError = 12
	GTK_BUILDER_ERROR_INVALID_ID             BuilderError = 13
)

// gtk_builder_error_quark : return type :
type ButtonBoxStyle int

const (
	GTK_BUTTONBOX_SPREAD ButtonBoxStyle = 1
	GTK_BUTTONBOX_EDGE   ButtonBoxStyle = 2
	GTK_BUTTONBOX_START  ButtonBoxStyle = 3
	GTK_BUTTONBOX_END    ButtonBoxStyle = 4
	GTK_BUTTONBOX_CENTER ButtonBoxStyle = 5
	GTK_BUTTONBOX_EXPAND ButtonBoxStyle = 6
)

type ButtonRole int

const (
	GTK_BUTTON_ROLE_NORMAL ButtonRole = 0
	GTK_BUTTON_ROLE_CHECK  ButtonRole = 1
	GTK_BUTTON_ROLE_RADIO  ButtonRole = 2
)

type ButtonsType int

const (
	GTK_BUTTONS_NONE      ButtonsType = 0
	GTK_BUTTONS_OK        ButtonsType = 1
	GTK_BUTTONS_CLOSE     ButtonsType = 2
	GTK_BUTTONS_CANCEL    ButtonsType = 3
	GTK_BUTTONS_YES_NO    ButtonsType = 4
	GTK_BUTTONS_OK_CANCEL ButtonsType = 5
)

type CellRendererAccelMode int

const (
	GTK_CELL_RENDERER_ACCEL_MODE_GTK          CellRendererAccelMode = 0
	GTK_CELL_RENDERER_ACCEL_MODE_OTHER        CellRendererAccelMode = 1
	GTK_CELL_RENDERER_ACCEL_MODE_MODIFIER_TAP CellRendererAccelMode = 2
)

type CellRendererMode int

const (
	GTK_CELL_RENDERER_MODE_INERT       CellRendererMode = 0
	GTK_CELL_RENDERER_MODE_ACTIVATABLE CellRendererMode = 1
	GTK_CELL_RENDERER_MODE_EDITABLE    CellRendererMode = 2
)

type CornerType int

const (
	GTK_CORNER_TOP_LEFT     CornerType = 0
	GTK_CORNER_BOTTOM_LEFT  CornerType = 1
	GTK_CORNER_TOP_RIGHT    CornerType = 2
	GTK_CORNER_BOTTOM_RIGHT CornerType = 3
)

type CssProviderError int

const (
	GTK_CSS_PROVIDER_ERROR_FAILED        CssProviderError = 0
	GTK_CSS_PROVIDER_ERROR_SYNTAX        CssProviderError = 1
	GTK_CSS_PROVIDER_ERROR_IMPORT        CssProviderError = 2
	GTK_CSS_PROVIDER_ERROR_NAME          CssProviderError = 3
	GTK_CSS_PROVIDER_ERROR_DEPRECATED    CssProviderError = 4
	GTK_CSS_PROVIDER_ERROR_UNKNOWN_VALUE CssProviderError = 5
)

// gtk_css_provider_error_quark : return type :
type DeleteType int

const (
	GTK_DELETE_CHARS             DeleteType = 0
	GTK_DELETE_WORD_ENDS         DeleteType = 1
	GTK_DELETE_WORDS             DeleteType = 2
	GTK_DELETE_DISPLAY_LINES     DeleteType = 3
	GTK_DELETE_DISPLAY_LINE_ENDS DeleteType = 4
	GTK_DELETE_PARAGRAPH_ENDS    DeleteType = 5
	GTK_DELETE_PARAGRAPHS        DeleteType = 6
	GTK_DELETE_WHITESPACE        DeleteType = 7
)

type DirectionType int

const (
	GTK_DIR_TAB_FORWARD  DirectionType = 0
	GTK_DIR_TAB_BACKWARD DirectionType = 1
	GTK_DIR_UP           DirectionType = 2
	GTK_DIR_DOWN         DirectionType = 3
	GTK_DIR_LEFT         DirectionType = 4
	GTK_DIR_RIGHT        DirectionType = 5
)

type DragResult int

const (
	GTK_DRAG_RESULT_SUCCESS         DragResult = 0
	GTK_DRAG_RESULT_NO_TARGET       DragResult = 1
	GTK_DRAG_RESULT_USER_CANCELLED  DragResult = 2
	GTK_DRAG_RESULT_TIMEOUT_EXPIRED DragResult = 3
	GTK_DRAG_RESULT_GRAB_BROKEN     DragResult = 4
	GTK_DRAG_RESULT_ERROR           DragResult = 5
)

type ExpanderStyle int

const (
	GTK_EXPANDER_COLLAPSED      ExpanderStyle = 0
	GTK_EXPANDER_SEMI_COLLAPSED ExpanderStyle = 1
	GTK_EXPANDER_SEMI_EXPANDED  ExpanderStyle = 2
	GTK_EXPANDER_EXPANDED       ExpanderStyle = 3
)

type FileChooserAction int

const (
	GTK_FILE_CHOOSER_ACTION_OPEN          FileChooserAction = 0
	GTK_FILE_CHOOSER_ACTION_SAVE          FileChooserAction = 1
	GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER FileChooserAction = 2
	GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER FileChooserAction = 3
)

type FileChooserError int

const (
	GTK_FILE_CHOOSER_ERROR_NONEXISTENT         FileChooserError = 0
	GTK_FILE_CHOOSER_ERROR_BAD_FILENAME        FileChooserError = 1
	GTK_FILE_CHOOSER_ERROR_ALREADY_EXISTS      FileChooserError = 2
	GTK_FILE_CHOOSER_ERROR_INCOMPLETE_HOSTNAME FileChooserError = 3
)

type IMPreeditStyle int

const (
	GTK_IM_PREEDIT_NOTHING  IMPreeditStyle = 0
	GTK_IM_PREEDIT_CALLBACK IMPreeditStyle = 1
	GTK_IM_PREEDIT_NONE     IMPreeditStyle = 2
)

type IMStatusStyle int

const (
	GTK_IM_STATUS_NOTHING  IMStatusStyle = 0
	GTK_IM_STATUS_CALLBACK IMStatusStyle = 1
	GTK_IM_STATUS_NONE     IMStatusStyle = 2
)

type IconSize int

const (
	GTK_ICON_SIZE_INVALID       IconSize = 0
	GTK_ICON_SIZE_MENU          IconSize = 1
	GTK_ICON_SIZE_SMALL_TOOLBAR IconSize = 2
	GTK_ICON_SIZE_LARGE_TOOLBAR IconSize = 3
	GTK_ICON_SIZE_BUTTON        IconSize = 4
	GTK_ICON_SIZE_DND           IconSize = 5
	GTK_ICON_SIZE_DIALOG        IconSize = 6
)

// gtk_icon_size_from_name : return type :
// gtk_icon_size_get_name : return type :
// gtk_icon_size_lookup : return type :
// gtk_icon_size_register : return type :
type IconThemeError int

const (
	GTK_ICON_THEME_NOT_FOUND IconThemeError = 0
	GTK_ICON_THEME_FAILED    IconThemeError = 1
)

// gtk_icon_theme_error_quark : return type :
type IconViewDropPosition int

const (
	GTK_ICON_VIEW_NO_DROP    IconViewDropPosition = 0
	GTK_ICON_VIEW_DROP_INTO  IconViewDropPosition = 1
	GTK_ICON_VIEW_DROP_LEFT  IconViewDropPosition = 2
	GTK_ICON_VIEW_DROP_RIGHT IconViewDropPosition = 3
	GTK_ICON_VIEW_DROP_ABOVE IconViewDropPosition = 4
	GTK_ICON_VIEW_DROP_BELOW IconViewDropPosition = 5
)

type ImageType int

const (
	GTK_IMAGE_EMPTY     ImageType = 0
	GTK_IMAGE_PIXBUF    ImageType = 1
	GTK_IMAGE_STOCK     ImageType = 2
	GTK_IMAGE_ICON_SET  ImageType = 3
	GTK_IMAGE_ANIMATION ImageType = 4
	GTK_IMAGE_ICON_NAME ImageType = 5
	GTK_IMAGE_GICON     ImageType = 6
	GTK_IMAGE_SURFACE   ImageType = 7
)

type Justification int

const (
	GTK_JUSTIFY_LEFT   Justification = 0
	GTK_JUSTIFY_RIGHT  Justification = 1
	GTK_JUSTIFY_CENTER Justification = 2
	GTK_JUSTIFY_FILL   Justification = 3
)

type MenuDirectionType int

const (
	GTK_MENU_DIR_PARENT MenuDirectionType = 0
	GTK_MENU_DIR_CHILD  MenuDirectionType = 1
	GTK_MENU_DIR_NEXT   MenuDirectionType = 2
	GTK_MENU_DIR_PREV   MenuDirectionType = 3
)

type MessageType int

const (
	GTK_MESSAGE_INFO     MessageType = 0
	GTK_MESSAGE_WARNING  MessageType = 1
	GTK_MESSAGE_QUESTION MessageType = 2
	GTK_MESSAGE_ERROR    MessageType = 3
	GTK_MESSAGE_OTHER    MessageType = 4
)

type MovementStep int

const (
	GTK_MOVEMENT_LOGICAL_POSITIONS MovementStep = 0
	GTK_MOVEMENT_VISUAL_POSITIONS  MovementStep = 1
	GTK_MOVEMENT_WORDS             MovementStep = 2
	GTK_MOVEMENT_DISPLAY_LINES     MovementStep = 3
	GTK_MOVEMENT_DISPLAY_LINE_ENDS MovementStep = 4
	GTK_MOVEMENT_PARAGRAPHS        MovementStep = 5
	GTK_MOVEMENT_PARAGRAPH_ENDS    MovementStep = 6
	GTK_MOVEMENT_PAGES             MovementStep = 7
	GTK_MOVEMENT_BUFFER_ENDS       MovementStep = 8
	GTK_MOVEMENT_HORIZONTAL_PAGES  MovementStep = 9
)

type NotebookTab int

const (
	GTK_NOTEBOOK_TAB_FIRST NotebookTab = 0
	GTK_NOTEBOOK_TAB_LAST  NotebookTab = 1
)

type NumberUpLayout int

const (
	GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_TOP_TO_BOTTOM NumberUpLayout = 0
	GTK_NUMBER_UP_LAYOUT_LEFT_TO_RIGHT_BOTTOM_TO_TOP NumberUpLayout = 1
	GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_TOP_TO_BOTTOM NumberUpLayout = 2
	GTK_NUMBER_UP_LAYOUT_RIGHT_TO_LEFT_BOTTOM_TO_TOP NumberUpLayout = 3
	GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_LEFT_TO_RIGHT NumberUpLayout = 4
	GTK_NUMBER_UP_LAYOUT_TOP_TO_BOTTOM_RIGHT_TO_LEFT NumberUpLayout = 5
	GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_LEFT_TO_RIGHT NumberUpLayout = 6
	GTK_NUMBER_UP_LAYOUT_BOTTOM_TO_TOP_RIGHT_TO_LEFT NumberUpLayout = 7
)

type Orientation int

const (
	GTK_ORIENTATION_HORIZONTAL Orientation = 0
	GTK_ORIENTATION_VERTICAL   Orientation = 1
)

type PackDirection int

const (
	GTK_PACK_DIRECTION_LTR PackDirection = 0
	GTK_PACK_DIRECTION_RTL PackDirection = 1
	GTK_PACK_DIRECTION_TTB PackDirection = 2
	GTK_PACK_DIRECTION_BTT PackDirection = 3
)

type PackType int

const (
	GTK_PACK_START PackType = 0
	GTK_PACK_END   PackType = 1
)

type PageOrientation int

const (
	GTK_PAGE_ORIENTATION_PORTRAIT          PageOrientation = 0
	GTK_PAGE_ORIENTATION_LANDSCAPE         PageOrientation = 1
	GTK_PAGE_ORIENTATION_REVERSE_PORTRAIT  PageOrientation = 2
	GTK_PAGE_ORIENTATION_REVERSE_LANDSCAPE PageOrientation = 3
)

type PageSet int

const (
	GTK_PAGE_SET_ALL  PageSet = 0
	GTK_PAGE_SET_EVEN PageSet = 1
	GTK_PAGE_SET_ODD  PageSet = 2
)

type PathPriorityType int

const (
	GTK_PATH_PRIO_LOWEST      PathPriorityType = 0
	GTK_PATH_PRIO_GTK         PathPriorityType = 4
	GTK_PATH_PRIO_APPLICATION PathPriorityType = 8
	GTK_PATH_PRIO_THEME       PathPriorityType = 10
	GTK_PATH_PRIO_RC          PathPriorityType = 12
	GTK_PATH_PRIO_HIGHEST     PathPriorityType = 15
)

type PathType int

const (
	GTK_PATH_WIDGET       PathType = 0
	GTK_PATH_WIDGET_CLASS PathType = 1
	GTK_PATH_CLASS        PathType = 2
)

type PolicyType int

const (
	GTK_POLICY_ALWAYS    PolicyType = 0
	GTK_POLICY_AUTOMATIC PolicyType = 1
	GTK_POLICY_NEVER     PolicyType = 2
	GTK_POLICY_EXTERNAL  PolicyType = 3
)

type PositionType int

const (
	GTK_POS_LEFT   PositionType = 0
	GTK_POS_RIGHT  PositionType = 1
	GTK_POS_TOP    PositionType = 2
	GTK_POS_BOTTOM PositionType = 3
)

type PrintDuplex int

const (
	GTK_PRINT_DUPLEX_SIMPLEX    PrintDuplex = 0
	GTK_PRINT_DUPLEX_HORIZONTAL PrintDuplex = 1
	GTK_PRINT_DUPLEX_VERTICAL   PrintDuplex = 2
)

type PrintError int

const (
	GTK_PRINT_ERROR_GENERAL        PrintError = 0
	GTK_PRINT_ERROR_INTERNAL_ERROR PrintError = 1
	GTK_PRINT_ERROR_NOMEM          PrintError = 2
	GTK_PRINT_ERROR_INVALID_FILE   PrintError = 3
)

type PrintOperationAction int

const (
	GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG PrintOperationAction = 0
	GTK_PRINT_OPERATION_ACTION_PRINT        PrintOperationAction = 1
	GTK_PRINT_OPERATION_ACTION_PREVIEW      PrintOperationAction = 2
	GTK_PRINT_OPERATION_ACTION_EXPORT       PrintOperationAction = 3
)

type PrintOperationResult int

const (
	GTK_PRINT_OPERATION_RESULT_ERROR       PrintOperationResult = 0
	GTK_PRINT_OPERATION_RESULT_APPLY       PrintOperationResult = 1
	GTK_PRINT_OPERATION_RESULT_CANCEL      PrintOperationResult = 2
	GTK_PRINT_OPERATION_RESULT_IN_PROGRESS PrintOperationResult = 3
)

type PrintPages int

const (
	GTK_PRINT_PAGES_ALL       PrintPages = 0
	GTK_PRINT_PAGES_CURRENT   PrintPages = 1
	GTK_PRINT_PAGES_RANGES    PrintPages = 2
	GTK_PRINT_PAGES_SELECTION PrintPages = 3
)

type PrintQuality int

const (
	GTK_PRINT_QUALITY_LOW    PrintQuality = 0
	GTK_PRINT_QUALITY_NORMAL PrintQuality = 1
	GTK_PRINT_QUALITY_HIGH   PrintQuality = 2
	GTK_PRINT_QUALITY_DRAFT  PrintQuality = 3
)

type PrintStatus int

const (
	GTK_PRINT_STATUS_INITIAL          PrintStatus = 0
	GTK_PRINT_STATUS_PREPARING        PrintStatus = 1
	GTK_PRINT_STATUS_GENERATING_DATA  PrintStatus = 2
	GTK_PRINT_STATUS_SENDING_DATA     PrintStatus = 3
	GTK_PRINT_STATUS_PENDING          PrintStatus = 4
	GTK_PRINT_STATUS_PENDING_ISSUE    PrintStatus = 5
	GTK_PRINT_STATUS_PRINTING         PrintStatus = 6
	GTK_PRINT_STATUS_FINISHED         PrintStatus = 7
	GTK_PRINT_STATUS_FINISHED_ABORTED PrintStatus = 8
)

type RcTokenType int

const (
	GTK_RC_TOKEN_INVALID        RcTokenType = 270
	GTK_RC_TOKEN_INCLUDE        RcTokenType = 271
	GTK_RC_TOKEN_NORMAL         RcTokenType = 272
	GTK_RC_TOKEN_ACTIVE         RcTokenType = 273
	GTK_RC_TOKEN_PRELIGHT       RcTokenType = 274
	GTK_RC_TOKEN_SELECTED       RcTokenType = 275
	GTK_RC_TOKEN_INSENSITIVE    RcTokenType = 276
	GTK_RC_TOKEN_FG             RcTokenType = 277
	GTK_RC_TOKEN_BG             RcTokenType = 278
	GTK_RC_TOKEN_TEXT           RcTokenType = 279
	GTK_RC_TOKEN_BASE           RcTokenType = 280
	GTK_RC_TOKEN_XTHICKNESS     RcTokenType = 281
	GTK_RC_TOKEN_YTHICKNESS     RcTokenType = 282
	GTK_RC_TOKEN_FONT           RcTokenType = 283
	GTK_RC_TOKEN_FONTSET        RcTokenType = 284
	GTK_RC_TOKEN_FONT_NAME      RcTokenType = 285
	GTK_RC_TOKEN_BG_PIXMAP      RcTokenType = 286
	GTK_RC_TOKEN_PIXMAP_PATH    RcTokenType = 287
	GTK_RC_TOKEN_STYLE          RcTokenType = 288
	GTK_RC_TOKEN_BINDING        RcTokenType = 289
	GTK_RC_TOKEN_BIND           RcTokenType = 290
	GTK_RC_TOKEN_WIDGET         RcTokenType = 291
	GTK_RC_TOKEN_WIDGET_CLASS   RcTokenType = 292
	GTK_RC_TOKEN_CLASS          RcTokenType = 293
	GTK_RC_TOKEN_LOWEST         RcTokenType = 294
	GTK_RC_TOKEN_GTK            RcTokenType = 295
	GTK_RC_TOKEN_APPLICATION    RcTokenType = 296
	GTK_RC_TOKEN_THEME          RcTokenType = 297
	GTK_RC_TOKEN_RC             RcTokenType = 298
	GTK_RC_TOKEN_HIGHEST        RcTokenType = 299
	GTK_RC_TOKEN_ENGINE         RcTokenType = 300
	GTK_RC_TOKEN_MODULE_PATH    RcTokenType = 301
	GTK_RC_TOKEN_IM_MODULE_PATH RcTokenType = 302
	GTK_RC_TOKEN_IM_MODULE_FILE RcTokenType = 303
	GTK_RC_TOKEN_STOCK          RcTokenType = 304
	GTK_RC_TOKEN_LTR            RcTokenType = 305
	GTK_RC_TOKEN_RTL            RcTokenType = 306
	GTK_RC_TOKEN_COLOR          RcTokenType = 307
	GTK_RC_TOKEN_UNBIND         RcTokenType = 308
	GTK_RC_TOKEN_LAST           RcTokenType = 309
)

type ReliefStyle int

const (
	GTK_RELIEF_NORMAL ReliefStyle = 0
	GTK_RELIEF_HALF   ReliefStyle = 1
	GTK_RELIEF_NONE   ReliefStyle = 2
)

type ResizeMode int

const (
	GTK_RESIZE_PARENT    ResizeMode = 0
	GTK_RESIZE_QUEUE     ResizeMode = 1
	GTK_RESIZE_IMMEDIATE ResizeMode = 2
)

type ResponseType int

const (
	GTK_RESPONSE_NONE         ResponseType = -1
	GTK_RESPONSE_REJECT       ResponseType = -2
	GTK_RESPONSE_ACCEPT       ResponseType = -3
	GTK_RESPONSE_DELETE_EVENT ResponseType = -4
	GTK_RESPONSE_OK           ResponseType = -5
	GTK_RESPONSE_CANCEL       ResponseType = -6
	GTK_RESPONSE_CLOSE        ResponseType = -7
	GTK_RESPONSE_YES          ResponseType = -8
	GTK_RESPONSE_NO           ResponseType = -9
	GTK_RESPONSE_APPLY        ResponseType = -10
	GTK_RESPONSE_HELP         ResponseType = -11
)

type RevealerTransitionType int

const (
	GTK_REVEALER_TRANSITION_TYPE_NONE        RevealerTransitionType = 0
	GTK_REVEALER_TRANSITION_TYPE_CROSSFADE   RevealerTransitionType = 1
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_RIGHT RevealerTransitionType = 2
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_LEFT  RevealerTransitionType = 3
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_UP    RevealerTransitionType = 4
	GTK_REVEALER_TRANSITION_TYPE_SLIDE_DOWN  RevealerTransitionType = 5
)

type ScrollStep int

const (
	GTK_SCROLL_STEPS            ScrollStep = 0
	GTK_SCROLL_PAGES            ScrollStep = 1
	GTK_SCROLL_ENDS             ScrollStep = 2
	GTK_SCROLL_HORIZONTAL_STEPS ScrollStep = 3
	GTK_SCROLL_HORIZONTAL_PAGES ScrollStep = 4
	GTK_SCROLL_HORIZONTAL_ENDS  ScrollStep = 5
)

type ScrollType int

const (
	GTK_SCROLL_NONE          ScrollType = 0
	GTK_SCROLL_JUMP          ScrollType = 1
	GTK_SCROLL_STEP_BACKWARD ScrollType = 2
	GTK_SCROLL_STEP_FORWARD  ScrollType = 3
	GTK_SCROLL_PAGE_BACKWARD ScrollType = 4
	GTK_SCROLL_PAGE_FORWARD  ScrollType = 5
	GTK_SCROLL_STEP_UP       ScrollType = 6
	GTK_SCROLL_STEP_DOWN     ScrollType = 7
	GTK_SCROLL_PAGE_UP       ScrollType = 8
	GTK_SCROLL_PAGE_DOWN     ScrollType = 9
	GTK_SCROLL_STEP_LEFT     ScrollType = 10
	GTK_SCROLL_STEP_RIGHT    ScrollType = 11
	GTK_SCROLL_PAGE_LEFT     ScrollType = 12
	GTK_SCROLL_PAGE_RIGHT    ScrollType = 13
	GTK_SCROLL_START         ScrollType = 14
	GTK_SCROLL_END           ScrollType = 15
)

type ScrollablePolicy int

const (
	GTK_SCROLL_MINIMUM ScrollablePolicy = 0
	GTK_SCROLL_NATURAL ScrollablePolicy = 1
)

type SelectionMode int

const (
	GTK_SELECTION_NONE     SelectionMode = 0
	GTK_SELECTION_SINGLE   SelectionMode = 1
	GTK_SELECTION_BROWSE   SelectionMode = 2
	GTK_SELECTION_MULTIPLE SelectionMode = 3
)

type SensitivityType int

const (
	GTK_SENSITIVITY_AUTO SensitivityType = 0
	GTK_SENSITIVITY_ON   SensitivityType = 1
	GTK_SENSITIVITY_OFF  SensitivityType = 2
)

type ShadowType int

const (
	GTK_SHADOW_NONE       ShadowType = 0
	GTK_SHADOW_IN         ShadowType = 1
	GTK_SHADOW_OUT        ShadowType = 2
	GTK_SHADOW_ETCHED_IN  ShadowType = 3
	GTK_SHADOW_ETCHED_OUT ShadowType = 4
)

type SizeGroupMode int

const (
	GTK_SIZE_GROUP_NONE       SizeGroupMode = 0
	GTK_SIZE_GROUP_HORIZONTAL SizeGroupMode = 1
	GTK_SIZE_GROUP_VERTICAL   SizeGroupMode = 2
	GTK_SIZE_GROUP_BOTH       SizeGroupMode = 3
)

type SizeRequestMode int

const (
	GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH SizeRequestMode = 0
	GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT SizeRequestMode = 1
	GTK_SIZE_REQUEST_CONSTANT_SIZE    SizeRequestMode = 2
)

type SortType int

const (
	GTK_SORT_ASCENDING  SortType = 0
	GTK_SORT_DESCENDING SortType = 1
)

type SpinButtonUpdatePolicy int

const (
	GTK_UPDATE_ALWAYS   SpinButtonUpdatePolicy = 0
	GTK_UPDATE_IF_VALID SpinButtonUpdatePolicy = 1
)

type SpinType int

const (
	GTK_SPIN_STEP_FORWARD  SpinType = 0
	GTK_SPIN_STEP_BACKWARD SpinType = 1
	GTK_SPIN_PAGE_FORWARD  SpinType = 2
	GTK_SPIN_PAGE_BACKWARD SpinType = 3
	GTK_SPIN_HOME          SpinType = 4
	GTK_SPIN_END           SpinType = 5
	GTK_SPIN_USER_DEFINED  SpinType = 6
)

type StackTransitionType int

const (
	GTK_STACK_TRANSITION_TYPE_NONE             StackTransitionType = 0
	GTK_STACK_TRANSITION_TYPE_CROSSFADE        StackTransitionType = 1
	GTK_STACK_TRANSITION_TYPE_SLIDE_RIGHT      StackTransitionType = 2
	GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT       StackTransitionType = 3
	GTK_STACK_TRANSITION_TYPE_SLIDE_UP         StackTransitionType = 4
	GTK_STACK_TRANSITION_TYPE_SLIDE_DOWN       StackTransitionType = 5
	GTK_STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT StackTransitionType = 6
	GTK_STACK_TRANSITION_TYPE_SLIDE_UP_DOWN    StackTransitionType = 7
	GTK_STACK_TRANSITION_TYPE_OVER_UP          StackTransitionType = 8
	GTK_STACK_TRANSITION_TYPE_OVER_DOWN        StackTransitionType = 9
	GTK_STACK_TRANSITION_TYPE_OVER_LEFT        StackTransitionType = 10
	GTK_STACK_TRANSITION_TYPE_OVER_RIGHT       StackTransitionType = 11
	GTK_STACK_TRANSITION_TYPE_UNDER_UP         StackTransitionType = 12
	GTK_STACK_TRANSITION_TYPE_UNDER_DOWN       StackTransitionType = 13
	GTK_STACK_TRANSITION_TYPE_UNDER_LEFT       StackTransitionType = 14
	GTK_STACK_TRANSITION_TYPE_UNDER_RIGHT      StackTransitionType = 15
	GTK_STACK_TRANSITION_TYPE_OVER_UP_DOWN     StackTransitionType = 16
	GTK_STACK_TRANSITION_TYPE_OVER_DOWN_UP     StackTransitionType = 17
	GTK_STACK_TRANSITION_TYPE_OVER_LEFT_RIGHT  StackTransitionType = 18
	GTK_STACK_TRANSITION_TYPE_OVER_RIGHT_LEFT  StackTransitionType = 19
)

type StateType int

const (
	GTK_STATE_NORMAL       StateType = 0
	GTK_STATE_ACTIVE       StateType = 1
	GTK_STATE_PRELIGHT     StateType = 2
	GTK_STATE_SELECTED     StateType = 3
	GTK_STATE_INSENSITIVE  StateType = 4
	GTK_STATE_INCONSISTENT StateType = 5
	GTK_STATE_FOCUSED      StateType = 6
)

type TextBufferTargetInfo int

const (
	GTK_TEXT_BUFFER_TARGET_INFO_BUFFER_CONTENTS TextBufferTargetInfo = -1
	GTK_TEXT_BUFFER_TARGET_INFO_RICH_TEXT       TextBufferTargetInfo = -2
	GTK_TEXT_BUFFER_TARGET_INFO_TEXT            TextBufferTargetInfo = -3
)

type TextDirection int

const (
	GTK_TEXT_DIR_NONE TextDirection = 0
	GTK_TEXT_DIR_LTR  TextDirection = 1
	GTK_TEXT_DIR_RTL  TextDirection = 2
)

type TextViewLayer int

const (
	GTK_TEXT_VIEW_LAYER_BELOW      TextViewLayer = 0
	GTK_TEXT_VIEW_LAYER_ABOVE      TextViewLayer = 1
	GTK_TEXT_VIEW_LAYER_BELOW_TEXT TextViewLayer = 2
	GTK_TEXT_VIEW_LAYER_ABOVE_TEXT TextViewLayer = 3
)

type TextWindowType int

const (
	GTK_TEXT_WINDOW_PRIVATE TextWindowType = 0
	GTK_TEXT_WINDOW_WIDGET  TextWindowType = 1
	GTK_TEXT_WINDOW_TEXT    TextWindowType = 2
	GTK_TEXT_WINDOW_LEFT    TextWindowType = 3
	GTK_TEXT_WINDOW_RIGHT   TextWindowType = 4
	GTK_TEXT_WINDOW_TOP     TextWindowType = 5
	GTK_TEXT_WINDOW_BOTTOM  TextWindowType = 6
)

type ToolbarSpaceStyle int

const (
	GTK_TOOLBAR_SPACE_EMPTY ToolbarSpaceStyle = 0
	GTK_TOOLBAR_SPACE_LINE  ToolbarSpaceStyle = 1
)

type ToolbarStyle int

const (
	GTK_TOOLBAR_ICONS      ToolbarStyle = 0
	GTK_TOOLBAR_TEXT       ToolbarStyle = 1
	GTK_TOOLBAR_BOTH       ToolbarStyle = 2
	GTK_TOOLBAR_BOTH_HORIZ ToolbarStyle = 3
)

type TreeViewColumnSizing int

const (
	GTK_TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = 0
	GTK_TREE_VIEW_COLUMN_AUTOSIZE  TreeViewColumnSizing = 1
	GTK_TREE_VIEW_COLUMN_FIXED     TreeViewColumnSizing = 2
)

type TreeViewDropPosition int

const (
	GTK_TREE_VIEW_DROP_BEFORE         TreeViewDropPosition = 0
	GTK_TREE_VIEW_DROP_AFTER          TreeViewDropPosition = 1
	GTK_TREE_VIEW_DROP_INTO_OR_BEFORE TreeViewDropPosition = 2
	GTK_TREE_VIEW_DROP_INTO_OR_AFTER  TreeViewDropPosition = 3
)

type TreeViewGridLines int

const (
	GTK_TREE_VIEW_GRID_LINES_NONE       TreeViewGridLines = 0
	GTK_TREE_VIEW_GRID_LINES_HORIZONTAL TreeViewGridLines = 1
	GTK_TREE_VIEW_GRID_LINES_VERTICAL   TreeViewGridLines = 2
	GTK_TREE_VIEW_GRID_LINES_BOTH       TreeViewGridLines = 3
)

type Unit int

const (
	GTK_UNIT_NONE   Unit = 0
	GTK_UNIT_POINTS Unit = 1
	GTK_UNIT_INCH   Unit = 2
	GTK_UNIT_MM     Unit = 3
)

type WidgetHelpType int

const (
	GTK_WIDGET_HELP_TOOLTIP    WidgetHelpType = 0
	GTK_WIDGET_HELP_WHATS_THIS WidgetHelpType = 1
)

type WindowPosition int

const (
	GTK_WIN_POS_NONE             WindowPosition = 0
	GTK_WIN_POS_CENTER           WindowPosition = 1
	GTK_WIN_POS_MOUSE            WindowPosition = 2
	GTK_WIN_POS_CENTER_ALWAYS    WindowPosition = 3
	GTK_WIN_POS_CENTER_ON_PARENT WindowPosition = 4
)

type WindowType int

const (
	GTK_WINDOW_TOPLEVEL WindowType = 0
	GTK_WINDOW_POPUP    WindowType = 1
)

type WrapMode int

const (
	GTK_WRAP_NONE      WrapMode = 0
	GTK_WRAP_CHAR      WrapMode = 1
	GTK_WRAP_WORD      WrapMode = 2
	GTK_WRAP_WORD_CHAR WrapMode = 3
)

// Unsupported : gtk_accel_groups_activate : return type :

// Unsupported : gtk_accel_groups_from_object : return type :

// Unsupported : gtk_accelerator_get_default_mod_mask : return type :

// Unsupported : gtk_accelerator_name : return type :

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Unsupported : gtk_accelerator_valid : return type :

// Unsupported : gtk_binding_set_by_class : unsupported parameter object_class : no type generator for gpointer (gpointer) for param object_class

// Unsupported : gtk_binding_set_find : return type :

// Unsupported : gtk_binding_set_new : return type :

// Unsupported : gtk_bindings_activate : return type :

// Unsupported : gtk_builder_error_quark : return type :

// Unsupported : gtk_check_version : return type :

// Unsupported : gtk_css_provider_error_quark : return type :

// DisableSetlocale is a wrapper around the C function gtk_disable_setlocale.
func DisableSetlocale() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(5471, &data)
	return
}

// Unsupported : gtk_drag_get_source_widget : return type :

// Unsupported : gtk_events_pending : return type :

// Unsupported : gtk_false : return type :

// Unsupported : gtk_get_current_event : no return generator

// Unsupported : gtk_get_current_event_device : return type :

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// GetCurrentEventTime is a wrapper around the C function gtk_get_current_event_time.
func GetCurrentEventTime() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(5906, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetDebugFlags is a wrapper around the C function gtk_get_debug_flags.
func GetDebugFlags() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(5907, &data)
	ret := data.Return.Uint32()

	return ret
}

// Unsupported : gtk_get_default_language : return type :

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_grab_get_current : return type :

// Unsupported : gtk_icon_size_from_name : return type :

// Unsupported : gtk_icon_size_get_name : return type :

// Unsupported : gtk_icon_size_lookup : return type :

// Unsupported : gtk_icon_size_register : return type :

// Unsupported : gtk_icon_theme_error_quark : return type :

// Unsupported : gtk_init_check : return type :

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc (GtkKeySnoopFunc) for param snooper

// Main is a wrapper around the C function gtk_main.
func Main() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(6390, &data)
	return
}

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_main_iteration : return type :

// Unsupported : gtk_main_iteration_do : return type :

// MainLevel is a wrapper around the C function gtk_main_level.
func MainLevel() uint32 {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_UINT},
	}
	call.Function(6394, &data)
	ret := data.Return.Uint32()

	return ret
}

// MainQuit is a wrapper around the C function gtk_main_quit.
func MainQuit() {
	data := call.Data{
		Params: []call.Value{},
		Return: call.Value{Type: call.TYPE_VOID},
	}
	call.Function(6395, &data)
	return
}

// Unsupported : gtk_parse_args : return type :

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_rc_find_module_in_path : return type :

// Unsupported : gtk_rc_find_pixmap_in_path : return type :

// Unsupported : gtk_rc_get_default_files : array return type :

// Unsupported : gtk_rc_get_im_module_file : return type :

// Unsupported : gtk_rc_get_im_module_path : return type :

// Unsupported : gtk_rc_get_module_dir : return type :

// Unsupported : gtk_rc_get_style : return type :

// Unsupported : gtk_rc_get_style_by_paths : return type :

// Unsupported : gtk_rc_get_theme_dir : return type :

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// Unsupported : gtk_rc_property_parse_border : return type :

// Unsupported : gtk_rc_property_parse_color : return type :

// Unsupported : gtk_rc_property_parse_enum : return type :

// Unsupported : gtk_rc_property_parse_flags : return type :

// Unsupported : gtk_rc_property_parse_requisition : return type :

// Unsupported : gtk_rc_reparse_all : return type :

// Unsupported : gtk_rc_reparse_all_for_settings : return type :

// Unsupported : gtk_rc_scanner_new : return type :

// Unsupported : gtk_recent_chooser_error_quark : return type :

// Unsupported : gtk_recent_manager_error_quark : return type :

// Unsupported : gtk_selection_add_targets : unsupported parameter targets :

// Unsupported : gtk_selection_convert : return type :

// Unsupported : gtk_selection_owner_set : return type :

// Unsupported : gtk_stock_add : unsupported parameter items :

// Unsupported : gtk_stock_add_static : unsupported parameter items :

// Unsupported : gtk_stock_list_ids : return type :

// Unsupported : gtk_stock_lookup : return type :

// Unsupported : gtk_tree_get_row_drag_data : return type :

// Unsupported : gtk_tree_set_row_drag_data : return type :

// Unsupported : gtk_true : return type :

// Unsupported : gtk_icon_set_new : return type :

// Unsupported : gtk_icon_set_new_from_pixbuf : return type :

// Unsupported : gtk_icon_source_new : return type :

// Blacklisted : GtkStackAccessibleClass

// Unsupported : gtk_target_entry_new : return type :

// Unsupported : gtk_target_list_new : unsupported parameter targets :

// Unsupported : gtk_text_attributes_new : return type :

// Unsupported : gtk_tree_path_new : return type :

// Unsupported : gtk_tree_path_new_first : return type :

// Unsupported : gtk_tree_path_new_from_string : return type :

// Unsupported : gtk_tree_row_reference_new : return type :

// Unsupported : gtk_tree_row_reference_new_proxy : return type :
