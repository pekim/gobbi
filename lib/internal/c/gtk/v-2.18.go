// Code generated - DO NOT EDIT.
// +build gtk_2.18

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	"unsafe"
)

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.TRUE
}

type AboutDialogClass C.GtkAboutDialogClass
type AboutDialogPrivate C.GtkAboutDialogPrivate
type AccelGroupClass C.GtkAccelGroupClass
type AccelGroupEntry C.GtkAccelGroupEntry
type AccelGroupPrivate C.GtkAccelGroupPrivate
type AccelKey C.GtkAccelKey
type AccelLabelClass C.GtkAccelLabelClass
type AccelLabelPrivate C.GtkAccelLabelPrivate
type AccelMapClass C.GtkAccelMapClass
type AccessibleClass C.GtkAccessibleClass
type AccessiblePrivate C.GtkAccessiblePrivate
type ActionBarClass C.GtkActionBarClass
type ActionBarPrivate C.GtkActionBarPrivate
type ActionClass C.GtkActionClass
type ActionEntry C.GtkActionEntry
type ActionGroupClass C.GtkActionGroupClass
type ActionGroupPrivate C.GtkActionGroupPrivate
type ActionPrivate C.GtkActionPrivate
type ActionableInterface C.GtkActionableInterface
type ActivatableIface C.GtkActivatableIface
type AdjustmentClass C.GtkAdjustmentClass
type AdjustmentPrivate C.GtkAdjustmentPrivate
type AlignmentClass C.GtkAlignmentClass
type AlignmentPrivate C.GtkAlignmentPrivate
type AppChooserButtonClass C.GtkAppChooserButtonClass
type AppChooserButtonPrivate C.GtkAppChooserButtonPrivate
type AppChooserDialogClass C.GtkAppChooserDialogClass
type AppChooserDialogPrivate C.GtkAppChooserDialogPrivate
type AppChooserWidgetClass C.GtkAppChooserWidgetClass
type AppChooserWidgetPrivate C.GtkAppChooserWidgetPrivate
type ApplicationClass C.GtkApplicationClass
type ApplicationPrivate C.GtkApplicationPrivate
type ApplicationWindowClass C.GtkApplicationWindowClass
type ApplicationWindowPrivate C.GtkApplicationWindowPrivate
type ArrowAccessibleClass C.GtkArrowAccessibleClass
type ArrowAccessiblePrivate C.GtkArrowAccessiblePrivate
type ArrowClass C.GtkArrowClass
type ArrowPrivate C.GtkArrowPrivate
type AspectFrameClass C.GtkAspectFrameClass
type AspectFramePrivate C.GtkAspectFramePrivate
type AssistantClass C.GtkAssistantClass
type AssistantPrivate C.GtkAssistantPrivate
type BinClass C.GtkBinClass
type BinPrivate C.GtkBinPrivate
type BindingArg C.GtkBindingArg
type BindingEntry C.GtkBindingEntry
type BindingSet C.GtkBindingSet
type BindingSignal C.GtkBindingSignal
type BooleanCellAccessibleClass C.GtkBooleanCellAccessibleClass
type BooleanCellAccessiblePrivate C.GtkBooleanCellAccessiblePrivate
type Border C.GtkBorder
type BoxClass C.GtkBoxClass
type BoxPrivate C.GtkBoxPrivate
type BuildableIface C.GtkBuildableIface
type BuilderClass C.GtkBuilderClass
type BuilderPrivate C.GtkBuilderPrivate
type ButtonAccessibleClass C.GtkButtonAccessibleClass
type ButtonAccessiblePrivate C.GtkButtonAccessiblePrivate
type ButtonBoxClass C.GtkButtonBoxClass
type ButtonBoxPrivate C.GtkButtonBoxPrivate
type ButtonClass C.GtkButtonClass
type ButtonPrivate C.GtkButtonPrivate
type CalendarClass C.GtkCalendarClass
type CalendarPrivate C.GtkCalendarPrivate
type CellAccessibleClass C.GtkCellAccessibleClass
type CellAccessibleParentIface C.GtkCellAccessibleParentIface
type CellAccessiblePrivate C.GtkCellAccessiblePrivate
type CellAreaBoxClass C.GtkCellAreaBoxClass
type CellAreaBoxPrivate C.GtkCellAreaBoxPrivate
type CellAreaClass C.GtkCellAreaClass
type CellAreaContextClass C.GtkCellAreaContextClass
type CellAreaContextPrivate C.GtkCellAreaContextPrivate
type CellAreaPrivate C.GtkCellAreaPrivate
type CellEditableIface C.GtkCellEditableIface
type CellLayoutIface C.GtkCellLayoutIface
type CellRendererAccelClass C.GtkCellRendererAccelClass
type CellRendererAccelPrivate C.GtkCellRendererAccelPrivate
type CellRendererClass C.GtkCellRendererClass
type CellRendererClassPrivate C.GtkCellRendererClassPrivate
type CellRendererComboClass C.GtkCellRendererComboClass
type CellRendererComboPrivate C.GtkCellRendererComboPrivate
type CellRendererPixbufClass C.GtkCellRendererPixbufClass
type CellRendererPixbufPrivate C.GtkCellRendererPixbufPrivate
type CellRendererPrivate C.GtkCellRendererPrivate
type CellRendererProgressClass C.GtkCellRendererProgressClass
type CellRendererProgressPrivate C.GtkCellRendererProgressPrivate
type CellRendererSpinClass C.GtkCellRendererSpinClass
type CellRendererSpinPrivate C.GtkCellRendererSpinPrivate
type CellRendererSpinnerClass C.GtkCellRendererSpinnerClass
type CellRendererSpinnerPrivate C.GtkCellRendererSpinnerPrivate
type CellRendererTextClass C.GtkCellRendererTextClass
type CellRendererTextPrivate C.GtkCellRendererTextPrivate
type CellRendererToggleClass C.GtkCellRendererToggleClass
type CellRendererTogglePrivate C.GtkCellRendererTogglePrivate
type CellViewClass C.GtkCellViewClass
type CellViewPrivate C.GtkCellViewPrivate
type CheckButtonClass C.GtkCheckButtonClass
type CheckMenuItemAccessibleClass C.GtkCheckMenuItemAccessibleClass
type CheckMenuItemAccessiblePrivate C.GtkCheckMenuItemAccessiblePrivate
type CheckMenuItemClass C.GtkCheckMenuItemClass
type CheckMenuItemPrivate C.GtkCheckMenuItemPrivate
type ColorButtonClass C.GtkColorButtonClass
type ColorButtonPrivate C.GtkColorButtonPrivate
type ColorChooserDialogClass C.GtkColorChooserDialogClass
type ColorChooserDialogPrivate C.GtkColorChooserDialogPrivate
type ColorChooserInterface C.GtkColorChooserInterface
type ColorChooserWidgetClass C.GtkColorChooserWidgetClass
type ColorChooserWidgetPrivate C.GtkColorChooserWidgetPrivate
type ColorSelectionClass C.GtkColorSelectionClass
type ColorSelectionDialogClass C.GtkColorSelectionDialogClass
type ColorSelectionDialogPrivate C.GtkColorSelectionDialogPrivate
type ColorSelectionPrivate C.GtkColorSelectionPrivate
type ComboBoxAccessibleClass C.GtkComboBoxAccessibleClass
type ComboBoxAccessiblePrivate C.GtkComboBoxAccessiblePrivate
type ComboBoxClass C.GtkComboBoxClass
type ComboBoxPrivate C.GtkComboBoxPrivate
type ComboBoxTextClass C.GtkComboBoxTextClass
type ComboBoxTextPrivate C.GtkComboBoxTextPrivate
type ContainerAccessibleClass C.GtkContainerAccessibleClass
type ContainerAccessiblePrivate C.GtkContainerAccessiblePrivate
type ContainerCellAccessibleClass C.GtkContainerCellAccessibleClass
type ContainerCellAccessiblePrivate C.GtkContainerCellAccessiblePrivate
type ContainerClass C.GtkContainerClass
type ContainerPrivate C.GtkContainerPrivate
type CssProviderClass C.GtkCssProviderClass
type CssProviderPrivate C.GtkCssProviderPrivate
type DialogClass C.GtkDialogClass
type DialogPrivate C.GtkDialogPrivate
type DrawingAreaClass C.GtkDrawingAreaClass
type EditableInterface C.GtkEditableInterface
type EntryAccessibleClass C.GtkEntryAccessibleClass
type EntryAccessiblePrivate C.GtkEntryAccessiblePrivate
type EntryBufferClass C.GtkEntryBufferClass
type EntryBufferPrivate C.GtkEntryBufferPrivate
type EntryClass C.GtkEntryClass
type EntryCompletionClass C.GtkEntryCompletionClass
type EntryCompletionPrivate C.GtkEntryCompletionPrivate
type EntryPrivate C.GtkEntryPrivate
type EventBoxClass C.GtkEventBoxClass
type EventBoxPrivate C.GtkEventBoxPrivate
type EventControllerClass C.GtkEventControllerClass

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
type ExpanderAccessibleClass C.GtkExpanderAccessibleClass
type ExpanderAccessiblePrivate C.GtkExpanderAccessiblePrivate
type ExpanderClass C.GtkExpanderClass
type ExpanderPrivate C.GtkExpanderPrivate
type FileChooserButtonClass C.GtkFileChooserButtonClass
type FileChooserButtonPrivate C.GtkFileChooserButtonPrivate
type FileChooserDialogClass C.GtkFileChooserDialogClass
type FileChooserDialogPrivate C.GtkFileChooserDialogPrivate
type FileChooserNativeClass C.GtkFileChooserNativeClass
type FileChooserWidgetClass C.GtkFileChooserWidgetClass
type FileChooserWidgetPrivate C.GtkFileChooserWidgetPrivate
type FileFilterInfo C.GtkFileFilterInfo
type FixedChild C.GtkFixedChild
type FixedClass C.GtkFixedClass
type FixedPrivate C.GtkFixedPrivate
type FlowBoxAccessibleClass C.GtkFlowBoxAccessibleClass
type FlowBoxAccessiblePrivate C.GtkFlowBoxAccessiblePrivate
type FlowBoxChildAccessibleClass C.GtkFlowBoxChildAccessibleClass
type FlowBoxChildClass C.GtkFlowBoxChildClass
type FlowBoxClass C.GtkFlowBoxClass
type FontButtonClass C.GtkFontButtonClass
type FontButtonPrivate C.GtkFontButtonPrivate
type FontChooserDialogClass C.GtkFontChooserDialogClass
type FontChooserDialogPrivate C.GtkFontChooserDialogPrivate
type FontChooserIface C.GtkFontChooserIface
type FontChooserWidgetClass C.GtkFontChooserWidgetClass
type FontChooserWidgetPrivate C.GtkFontChooserWidgetPrivate
type FontSelectionClass C.GtkFontSelectionClass
type FontSelectionDialogClass C.GtkFontSelectionDialogClass
type FontSelectionDialogPrivate C.GtkFontSelectionDialogPrivate
type FontSelectionPrivate C.GtkFontSelectionPrivate
type FrameAccessibleClass C.GtkFrameAccessibleClass
type FrameAccessiblePrivate C.GtkFrameAccessiblePrivate
type FrameClass C.GtkFrameClass
type FramePrivate C.GtkFramePrivate
type GestureClass C.GtkGestureClass
type GestureDragClass C.GtkGestureDragClass
type GestureLongPressClass C.GtkGestureLongPressClass
type GestureMultiPressClass C.GtkGestureMultiPressClass
type GesturePanClass C.GtkGesturePanClass
type GestureRotateClass C.GtkGestureRotateClass
type GestureSingleClass C.GtkGestureSingleClass

// UNSUPPORTED : GestureStylusClass : blacklisted
type GestureSwipeClass C.GtkGestureSwipeClass
type GestureZoomClass C.GtkGestureZoomClass
type Gradient C.GtkGradient
type GridClass C.GtkGridClass
type GridPrivate C.GtkGridPrivate
type HBoxClass C.GtkHBoxClass
type HButtonBoxClass C.GtkHButtonBoxClass
type HPanedClass C.GtkHPanedClass
type HSVClass C.GtkHSVClass
type HSVPrivate C.GtkHSVPrivate
type HScaleClass C.GtkHScaleClass
type HScrollbarClass C.GtkHScrollbarClass
type HSeparatorClass C.GtkHSeparatorClass
type HandleBoxClass C.GtkHandleBoxClass
type HandleBoxPrivate C.GtkHandleBoxPrivate

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
type HeaderBarClass C.GtkHeaderBarClass
type HeaderBarPrivate C.GtkHeaderBarPrivate
type IMContextClass C.GtkIMContextClass
type IMContextInfo C.GtkIMContextInfo
type IMContextSimpleClass C.GtkIMContextSimpleClass
type IMContextSimplePrivate C.GtkIMContextSimplePrivate
type IMMulticontextClass C.GtkIMMulticontextClass
type IMMulticontextPrivate C.GtkIMMulticontextPrivate
type IconFactoryClass C.GtkIconFactoryClass
type IconFactoryPrivate C.GtkIconFactoryPrivate
type IconInfoClass C.GtkIconInfoClass
type IconSet C.GtkIconSet
type IconSource C.GtkIconSource
type IconThemeClass C.GtkIconThemeClass
type IconThemePrivate C.GtkIconThemePrivate
type IconViewAccessibleClass C.GtkIconViewAccessibleClass
type IconViewAccessiblePrivate C.GtkIconViewAccessiblePrivate
type IconViewClass C.GtkIconViewClass
type IconViewPrivate C.GtkIconViewPrivate
type ImageAccessibleClass C.GtkImageAccessibleClass
type ImageAccessiblePrivate C.GtkImageAccessiblePrivate
type ImageCellAccessibleClass C.GtkImageCellAccessibleClass
type ImageCellAccessiblePrivate C.GtkImageCellAccessiblePrivate
type ImageClass C.GtkImageClass
type ImageMenuItemClass C.GtkImageMenuItemClass
type ImageMenuItemPrivate C.GtkImageMenuItemPrivate
type ImagePrivate C.GtkImagePrivate
type InfoBarClass C.GtkInfoBarClass
type InfoBarPrivate C.GtkInfoBarPrivate
type InvisibleClass C.GtkInvisibleClass
type InvisiblePrivate C.GtkInvisiblePrivate
type LabelAccessibleClass C.GtkLabelAccessibleClass
type LabelAccessiblePrivate C.GtkLabelAccessiblePrivate
type LabelClass C.GtkLabelClass
type LabelPrivate C.GtkLabelPrivate
type LabelSelectionInfo C.GtkLabelSelectionInfo
type LayoutClass C.GtkLayoutClass
type LayoutPrivate C.GtkLayoutPrivate
type LevelBarAccessibleClass C.GtkLevelBarAccessibleClass
type LevelBarAccessiblePrivate C.GtkLevelBarAccessiblePrivate
type LevelBarClass C.GtkLevelBarClass
type LevelBarPrivate C.GtkLevelBarPrivate
type LinkButtonAccessibleClass C.GtkLinkButtonAccessibleClass
type LinkButtonAccessiblePrivate C.GtkLinkButtonAccessiblePrivate
type LinkButtonClass C.GtkLinkButtonClass
type LinkButtonPrivate C.GtkLinkButtonPrivate
type ListBoxAccessibleClass C.GtkListBoxAccessibleClass
type ListBoxAccessiblePrivate C.GtkListBoxAccessiblePrivate
type ListBoxClass C.GtkListBoxClass
type ListBoxRowAccessibleClass C.GtkListBoxRowAccessibleClass
type ListBoxRowClass C.GtkListBoxRowClass
type ListStoreClass C.GtkListStoreClass
type ListStorePrivate C.GtkListStorePrivate
type LockButtonAccessibleClass C.GtkLockButtonAccessibleClass
type LockButtonAccessiblePrivate C.GtkLockButtonAccessiblePrivate
type LockButtonClass C.GtkLockButtonClass
type LockButtonPrivate C.GtkLockButtonPrivate
type MenuAccessibleClass C.GtkMenuAccessibleClass
type MenuAccessiblePrivate C.GtkMenuAccessiblePrivate
type MenuBarClass C.GtkMenuBarClass
type MenuBarPrivate C.GtkMenuBarPrivate
type MenuButtonAccessibleClass C.GtkMenuButtonAccessibleClass
type MenuButtonAccessiblePrivate C.GtkMenuButtonAccessiblePrivate
type MenuButtonClass C.GtkMenuButtonClass
type MenuButtonPrivate C.GtkMenuButtonPrivate
type MenuClass C.GtkMenuClass
type MenuItemAccessibleClass C.GtkMenuItemAccessibleClass
type MenuItemAccessiblePrivate C.GtkMenuItemAccessiblePrivate
type MenuItemClass C.GtkMenuItemClass
type MenuItemPrivate C.GtkMenuItemPrivate
type MenuPrivate C.GtkMenuPrivate
type MenuShellAccessibleClass C.GtkMenuShellAccessibleClass
type MenuShellAccessiblePrivate C.GtkMenuShellAccessiblePrivate
type MenuShellClass C.GtkMenuShellClass
type MenuShellPrivate C.GtkMenuShellPrivate
type MenuToolButtonClass C.GtkMenuToolButtonClass
type MenuToolButtonPrivate C.GtkMenuToolButtonPrivate
type MessageDialogClass C.GtkMessageDialogClass
type MessageDialogPrivate C.GtkMessageDialogPrivate
type MiscClass C.GtkMiscClass
type MiscPrivate C.GtkMiscPrivate
type MountOperationClass C.GtkMountOperationClass
type MountOperationPrivate C.GtkMountOperationPrivate
type NativeDialogClass C.GtkNativeDialogClass
type NotebookAccessibleClass C.GtkNotebookAccessibleClass
type NotebookAccessiblePrivate C.GtkNotebookAccessiblePrivate
type NotebookClass C.GtkNotebookClass
type NotebookPageAccessibleClass C.GtkNotebookPageAccessibleClass
type NotebookPageAccessiblePrivate C.GtkNotebookPageAccessiblePrivate
type NotebookPrivate C.GtkNotebookPrivate
type NumerableIconClass C.GtkNumerableIconClass
type NumerableIconPrivate C.GtkNumerableIconPrivate
type OffscreenWindowClass C.GtkOffscreenWindowClass
type OrientableIface C.GtkOrientableIface
type OverlayClass C.GtkOverlayClass
type OverlayPrivate C.GtkOverlayPrivate
type PadActionEntry C.GtkPadActionEntry
type PadControllerClass C.GtkPadControllerClass
type PageRange C.GtkPageRange
type PanedAccessibleClass C.GtkPanedAccessibleClass
type PanedAccessiblePrivate C.GtkPanedAccessiblePrivate
type PanedClass C.GtkPanedClass
type PanedPrivate C.GtkPanedPrivate
type PaperSize C.GtkPaperSize
type PlacesSidebarClass C.GtkPlacesSidebarClass
type PlugClass C.GtkPlugClass
type PlugPrivate C.GtkPlugPrivate
type PopoverAccessibleClass C.GtkPopoverAccessibleClass
type PopoverClass C.GtkPopoverClass
type PopoverMenuClass C.GtkPopoverMenuClass
type PopoverPrivate C.GtkPopoverPrivate
type PrintOperationClass C.GtkPrintOperationClass
type PrintOperationPreviewIface C.GtkPrintOperationPreviewIface
type PrintOperationPrivate C.GtkPrintOperationPrivate
type ProgressBarAccessibleClass C.GtkProgressBarAccessibleClass
type ProgressBarAccessiblePrivate C.GtkProgressBarAccessiblePrivate
type ProgressBarClass C.GtkProgressBarClass
type ProgressBarPrivate C.GtkProgressBarPrivate
type RadioActionClass C.GtkRadioActionClass
type RadioActionEntry C.GtkRadioActionEntry
type RadioActionPrivate C.GtkRadioActionPrivate
type RadioButtonAccessibleClass C.GtkRadioButtonAccessibleClass
type RadioButtonAccessiblePrivate C.GtkRadioButtonAccessiblePrivate
type RadioButtonClass C.GtkRadioButtonClass
type RadioButtonPrivate C.GtkRadioButtonPrivate
type RadioMenuItemAccessibleClass C.GtkRadioMenuItemAccessibleClass
type RadioMenuItemAccessiblePrivate C.GtkRadioMenuItemAccessiblePrivate
type RadioMenuItemClass C.GtkRadioMenuItemClass
type RadioMenuItemPrivate C.GtkRadioMenuItemPrivate
type RadioToolButtonClass C.GtkRadioToolButtonClass
type RangeAccessibleClass C.GtkRangeAccessibleClass
type RangeAccessiblePrivate C.GtkRangeAccessiblePrivate
type RangeClass C.GtkRangeClass
type RangePrivate C.GtkRangePrivate
type RcContext C.GtkRcContext
type RcProperty C.GtkRcProperty
type RcStyleClass C.GtkRcStyleClass
type RecentActionClass C.GtkRecentActionClass
type RecentActionPrivate C.GtkRecentActionPrivate
type RecentChooserDialogClass C.GtkRecentChooserDialogClass
type RecentChooserDialogPrivate C.GtkRecentChooserDialogPrivate
type RecentChooserIface C.GtkRecentChooserIface
type RecentChooserMenuClass C.GtkRecentChooserMenuClass
type RecentChooserMenuPrivate C.GtkRecentChooserMenuPrivate
type RecentChooserWidgetClass C.GtkRecentChooserWidgetClass
type RecentChooserWidgetPrivate C.GtkRecentChooserWidgetPrivate
type RecentData C.GtkRecentData
type RecentFilterInfo C.GtkRecentFilterInfo
type RecentInfo C.GtkRecentInfo
type RecentManagerClass C.GtkRecentManagerClass
type RecentManagerPrivate C.GtkRecentManagerPrivate
type RendererCellAccessibleClass C.GtkRendererCellAccessibleClass
type RendererCellAccessiblePrivate C.GtkRendererCellAccessiblePrivate
type RequestedSize C.GtkRequestedSize
type Requisition C.GtkRequisition
type RevealerClass C.GtkRevealerClass
type ScaleAccessibleClass C.GtkScaleAccessibleClass
type ScaleAccessiblePrivate C.GtkScaleAccessiblePrivate
type ScaleButtonAccessibleClass C.GtkScaleButtonAccessibleClass
type ScaleButtonAccessiblePrivate C.GtkScaleButtonAccessiblePrivate
type ScaleButtonClass C.GtkScaleButtonClass
type ScaleButtonPrivate C.GtkScaleButtonPrivate
type ScaleClass C.GtkScaleClass
type ScalePrivate C.GtkScalePrivate
type ScrollableInterface C.GtkScrollableInterface
type ScrollbarClass C.GtkScrollbarClass
type ScrolledWindowAccessibleClass C.GtkScrolledWindowAccessibleClass
type ScrolledWindowAccessiblePrivate C.GtkScrolledWindowAccessiblePrivate
type ScrolledWindowClass C.GtkScrolledWindowClass
type ScrolledWindowPrivate C.GtkScrolledWindowPrivate
type SearchBarClass C.GtkSearchBarClass
type SearchEntryClass C.GtkSearchEntryClass
type SelectionData C.GtkSelectionData
type SeparatorClass C.GtkSeparatorClass
type SeparatorMenuItemClass C.GtkSeparatorMenuItemClass
type SeparatorPrivate C.GtkSeparatorPrivate
type SeparatorToolItemClass C.GtkSeparatorToolItemClass
type SeparatorToolItemPrivate C.GtkSeparatorToolItemPrivate
type SettingsClass C.GtkSettingsClass
type SettingsPrivate C.GtkSettingsPrivate
type SettingsValue C.GtkSettingsValue
type ShortcutLabelClass C.GtkShortcutLabelClass
type ShortcutsGroupClass C.GtkShortcutsGroupClass
type ShortcutsSectionClass C.GtkShortcutsSectionClass
type ShortcutsShortcutClass C.GtkShortcutsShortcutClass
type ShortcutsWindowClass C.GtkShortcutsWindowClass
type SizeGroupClass C.GtkSizeGroupClass
type SizeGroupPrivate C.GtkSizeGroupPrivate
type SocketClass C.GtkSocketClass
type SocketPrivate C.GtkSocketPrivate
type SpinButtonAccessibleClass C.GtkSpinButtonAccessibleClass
type SpinButtonAccessiblePrivate C.GtkSpinButtonAccessiblePrivate
type SpinButtonClass C.GtkSpinButtonClass
type SpinButtonPrivate C.GtkSpinButtonPrivate
type SpinnerAccessibleClass C.GtkSpinnerAccessibleClass
type SpinnerAccessiblePrivate C.GtkSpinnerAccessiblePrivate
type SpinnerClass C.GtkSpinnerClass
type SpinnerPrivate C.GtkSpinnerPrivate

// UNSUPPORTED : StackAccessibleClass : blacklisted
type StackClass C.GtkStackClass
type StackSidebarClass C.GtkStackSidebarClass
type StackSidebarPrivate C.GtkStackSidebarPrivate
type StackSwitcherClass C.GtkStackSwitcherClass
type StatusIconClass C.GtkStatusIconClass
type StatusIconPrivate C.GtkStatusIconPrivate
type StatusbarAccessibleClass C.GtkStatusbarAccessibleClass
type StatusbarAccessiblePrivate C.GtkStatusbarAccessiblePrivate
type StatusbarClass C.GtkStatusbarClass
type StatusbarPrivate C.GtkStatusbarPrivate
type StockItem C.GtkStockItem
type StyleClass C.GtkStyleClass
type StyleContextClass C.GtkStyleContextClass
type StyleContextPrivate C.GtkStyleContextPrivate
type StylePropertiesClass C.GtkStylePropertiesClass
type StylePropertiesPrivate C.GtkStylePropertiesPrivate
type StyleProviderIface C.GtkStyleProviderIface
type SwitchAccessibleClass C.GtkSwitchAccessibleClass
type SwitchAccessiblePrivate C.GtkSwitchAccessiblePrivate
type SwitchClass C.GtkSwitchClass
type SwitchPrivate C.GtkSwitchPrivate
type SymbolicColor C.GtkSymbolicColor
type TableChild C.GtkTableChild
type TableClass C.GtkTableClass
type TablePrivate C.GtkTablePrivate
type TableRowCol C.GtkTableRowCol
type TargetEntry C.GtkTargetEntry
type TargetList C.GtkTargetList
type TargetPair C.GtkTargetPair
type TearoffMenuItemClass C.GtkTearoffMenuItemClass
type TearoffMenuItemPrivate C.GtkTearoffMenuItemPrivate
type TextAppearance C.GtkTextAppearance
type TextAttributes C.GtkTextAttributes
type TextBTree C.GtkTextBTree
type TextBufferClass C.GtkTextBufferClass
type TextBufferPrivate C.GtkTextBufferPrivate
type TextCellAccessibleClass C.GtkTextCellAccessibleClass
type TextCellAccessiblePrivate C.GtkTextCellAccessiblePrivate
type TextChildAnchorClass C.GtkTextChildAnchorClass
type TextIter C.GtkTextIter
type TextMarkClass C.GtkTextMarkClass
type TextTagClass C.GtkTextTagClass
type TextTagPrivate C.GtkTextTagPrivate
type TextTagTableClass C.GtkTextTagTableClass
type TextTagTablePrivate C.GtkTextTagTablePrivate
type TextViewAccessibleClass C.GtkTextViewAccessibleClass
type TextViewAccessiblePrivate C.GtkTextViewAccessiblePrivate
type TextViewClass C.GtkTextViewClass
type TextViewPrivate C.GtkTextViewPrivate
type ThemeEngine C.GtkThemeEngine
type ThemingEngineClass C.GtkThemingEngineClass
type ThemingEnginePrivate C.GtkThemingEnginePrivate
type ToggleActionClass C.GtkToggleActionClass
type ToggleActionEntry C.GtkToggleActionEntry
type ToggleActionPrivate C.GtkToggleActionPrivate
type ToggleButtonAccessibleClass C.GtkToggleButtonAccessibleClass
type ToggleButtonAccessiblePrivate C.GtkToggleButtonAccessiblePrivate
type ToggleButtonClass C.GtkToggleButtonClass
type ToggleButtonPrivate C.GtkToggleButtonPrivate
type ToggleToolButtonClass C.GtkToggleToolButtonClass
type ToggleToolButtonPrivate C.GtkToggleToolButtonPrivate
type ToolButtonClass C.GtkToolButtonClass
type ToolButtonPrivate C.GtkToolButtonPrivate
type ToolItemClass C.GtkToolItemClass
type ToolItemGroupClass C.GtkToolItemGroupClass
type ToolItemGroupPrivate C.GtkToolItemGroupPrivate
type ToolItemPrivate C.GtkToolItemPrivate
type ToolPaletteClass C.GtkToolPaletteClass
type ToolPalettePrivate C.GtkToolPalettePrivate
type ToolShellIface C.GtkToolShellIface
type ToolbarClass C.GtkToolbarClass
type ToolbarPrivate C.GtkToolbarPrivate
type ToplevelAccessibleClass C.GtkToplevelAccessibleClass
type ToplevelAccessiblePrivate C.GtkToplevelAccessiblePrivate
type TreeDragDestIface C.GtkTreeDragDestIface
type TreeDragSourceIface C.GtkTreeDragSourceIface
type TreeIter C.GtkTreeIter
type TreeModelFilterClass C.GtkTreeModelFilterClass
type TreeModelFilterPrivate C.GtkTreeModelFilterPrivate
type TreeModelIface C.GtkTreeModelIface
type TreeModelSortClass C.GtkTreeModelSortClass
type TreeModelSortPrivate C.GtkTreeModelSortPrivate
type TreePath C.GtkTreePath
type TreeRowReference C.GtkTreeRowReference
type TreeSelectionClass C.GtkTreeSelectionClass
type TreeSelectionPrivate C.GtkTreeSelectionPrivate
type TreeSortableIface C.GtkTreeSortableIface
type TreeStoreClass C.GtkTreeStoreClass
type TreeStorePrivate C.GtkTreeStorePrivate
type TreeViewAccessibleClass C.GtkTreeViewAccessibleClass
type TreeViewAccessiblePrivate C.GtkTreeViewAccessiblePrivate
type TreeViewClass C.GtkTreeViewClass
type TreeViewColumnClass C.GtkTreeViewColumnClass
type TreeViewColumnPrivate C.GtkTreeViewColumnPrivate
type TreeViewPrivate C.GtkTreeViewPrivate
type UIManagerClass C.GtkUIManagerClass
type UIManagerPrivate C.GtkUIManagerPrivate
type VBoxClass C.GtkVBoxClass
type VButtonBoxClass C.GtkVButtonBoxClass
type VPanedClass C.GtkVPanedClass
type VScaleClass C.GtkVScaleClass
type VScrollbarClass C.GtkVScrollbarClass
type VSeparatorClass C.GtkVSeparatorClass
type ViewportClass C.GtkViewportClass
type ViewportPrivate C.GtkViewportPrivate
type VolumeButtonClass C.GtkVolumeButtonClass
type WidgetAccessibleClass C.GtkWidgetAccessibleClass
type WidgetAccessiblePrivate C.GtkWidgetAccessiblePrivate
type WidgetClass C.GtkWidgetClass
type WidgetClassPrivate C.GtkWidgetClassPrivate
type WidgetPath C.GtkWidgetPath
type WidgetPrivate C.GtkWidgetPrivate
type WindowAccessibleClass C.GtkWindowAccessibleClass
type WindowAccessiblePrivate C.GtkWindowAccessiblePrivate
type WindowClass C.GtkWindowClass
type WindowGeometryInfo C.GtkWindowGeometryInfo
type WindowGroupClass C.GtkWindowGroupClass
type WindowGroupPrivate C.GtkWindowGroupPrivate
type WindowPrivate C.GtkWindowPrivate

func Fn_gtk_accel_groups_activate(param0 unsafe.Pointer, param1 uint, param2 int) {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	C.gtk_accel_groups_activate(cValue0, cValue1, cValue2)
}

func Fn_gtk_accel_groups_from_object(param0 unsafe.Pointer) {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	C.gtk_accel_groups_from_object(cValue0)
}

func Fn_gtk_accelerator_get_default_mod_mask() {
	C.gtk_accelerator_get_default_mod_mask()
}

func Fn_gtk_accelerator_get_label(param0 uint, param1 int) {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	C.gtk_accelerator_get_label(cValue0, cValue1)
}

func Fn_gtk_accelerator_name(param0 uint, param1 int) {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	C.gtk_accelerator_name(cValue0, cValue1)
}

func Fn_gtk_accelerator_parse(param0 string, param1 *uint, param2 *int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkModifierType)(unsafe.Pointer(param2))

	C.gtk_accelerator_parse(cValue0, cValue1, cValue2)
}

func Fn_gtk_accelerator_set_default_mod_mask(param0 int) {
	cValue0 := (C.GdkModifierType)(param0)

	C.gtk_accelerator_set_default_mod_mask(cValue0)
}

func Fn_gtk_accelerator_valid(param0 uint, param1 int) {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	C.gtk_accelerator_valid(cValue0, cValue1)
}

func Fn_gtk_alternative_dialog_button_order(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_alternative_dialog_button_order(cValue0)
}

func Fn_gtk_binding_entry_add_signall(param0 unsafe.Pointer, param1 uint, param2 int, param3 string, param4 unsafe.Pointer) {
	cValue0 := (*C.GtkBindingSet)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.GSList)(unsafe.Pointer(param4))

	C.gtk_binding_entry_add_signall(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_binding_entry_remove(param0 unsafe.Pointer, param1 uint, param2 int) {
	cValue0 := (*C.GtkBindingSet)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	C.gtk_binding_entry_remove(cValue0, cValue1, cValue2)
}

func Fn_gtk_binding_entry_skip(param0 unsafe.Pointer, param1 uint, param2 int) {
	cValue0 := (*C.GtkBindingSet)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	C.gtk_binding_entry_skip(cValue0, cValue1, cValue2)
}

func Fn_gtk_binding_set_by_class(param0 unsafe.Pointer) {
	cValue0 := (C.gpointer)(param0)

	C.gtk_binding_set_by_class(cValue0)
}

func Fn_gtk_binding_set_find(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_binding_set_find(cValue0)
}

func Fn_gtk_binding_set_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_binding_set_new(cValue0)
}

func Fn_gtk_bindings_activate(param0 unsafe.Pointer, param1 uint, param2 int) {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	C.gtk_bindings_activate(cValue0, cValue1, cValue2)
}

func Fn_gtk_bindings_activate_event(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkEventKey)(unsafe.Pointer(param1))

	C.gtk_bindings_activate_event(cValue0, cValue1)
}

func Fn_gtk_builder_error_quark() {
	C.gtk_builder_error_quark()
}

func Fn_gtk_check_version(param0 uint, param1 uint, param2 uint) {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	C.gtk_check_version(cValue0, cValue1, cValue2)
}

func Fn_gtk_css_provider_error_quark() {
	C.gtk_css_provider_error_quark()
}

func Fn_gtk_disable_setlocale() {
	C.gtk_disable_setlocale()
}

func Fn_gtk_distribute_natural_allocation(param0 int, param1 uint, param2 unsafe.Pointer) {
	cValue0 := (C.gint)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (*C.GtkRequestedSize)(unsafe.Pointer(param2))

	C.gtk_distribute_natural_allocation(cValue0, cValue1, cValue2)
}

func Fn_gtk_drag_finish(param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint32) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.guint32)(param3)

	C.gtk_drag_finish(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_drag_get_source_widget(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	C.gtk_drag_get_source_widget(cValue0)
}

func Fn_gtk_drag_set_icon_default(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	C.gtk_drag_set_icon_default(cValue0)
}

func Fn_gtk_drag_set_icon_name(param0 unsafe.Pointer, param1 string, param2 int, param3 int) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_drag_set_icon_name(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_drag_set_icon_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_drag_set_icon_pixbuf(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_drag_set_icon_stock(param0 unsafe.Pointer, param1 string, param2 int, param3 int) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_drag_set_icon_stock(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_drag_set_icon_surface(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_surface_t)(unsafe.Pointer(param1))

	C.gtk_drag_set_icon_surface(cValue0, cValue1)
}

func Fn_gtk_drag_set_icon_widget(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_drag_set_icon_widget(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_events_pending() {
	C.gtk_events_pending()
}

func Fn_gtk_false() {
	C.gtk_false()
}

func Fn_gtk_file_chooser_error_quark() {
	C.gtk_file_chooser_error_quark()
}

func Fn_gtk_get_current_event() {
	C.gtk_get_current_event()
}

func Fn_gtk_get_current_event_device() {
	C.gtk_get_current_event_device()
}

func Fn_gtk_get_current_event_state(param0 *int) {
	cValue0 := (*C.GdkModifierType)(unsafe.Pointer(param0))

	C.gtk_get_current_event_state(cValue0)
}

func Fn_gtk_get_current_event_time() {
	C.gtk_get_current_event_time()
}

func Fn_gtk_get_debug_flags() {
	C.gtk_get_debug_flags()
}

func Fn_gtk_get_default_language() {
	C.gtk_get_default_language()
}

func Fn_gtk_get_event_widget(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	C.gtk_get_event_widget(cValue0)
}

func Fn_gtk_get_option_group(param0 bool) {
	cValue0 := toCBool(param0)

	C.gtk_get_option_group(cValue0)
}

func Fn_gtk_grab_get_current() {
	C.gtk_grab_get_current()
}

func Fn_gtk_icon_size_from_name(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_size_from_name(cValue0)
}

func Fn_gtk_icon_size_get_name(param0 int) {
	cValue0 := (C.GtkIconSize)(param0)

	C.gtk_icon_size_get_name(cValue0)
}

func Fn_gtk_icon_size_lookup(param0 int, param1 *int, param2 *int) {
	cValue0 := (C.GtkIconSize)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_icon_size_lookup(cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_size_lookup_for_settings(param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_icon_size_lookup_for_settings(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_icon_size_register(param0 string, param1 int, param2 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_icon_size_register(cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_size_register_alias(param0 string, param1 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_icon_size_register_alias(cValue0, cValue1)
}

func Fn_gtk_icon_theme_error_quark() {
	C.gtk_icon_theme_error_quark()
}

func Fn_gtk_init(param0 *int, param1 *[]string) {
	cValue0 := (*C.int)(unsafe.Pointer(param0))

	var cValue1ArrayPointer **C.gchar
	cValue1 := &cValue1ArrayPointer
	param1Indirected := *param1
	param1IndirectedLen := len(param1Indirected)
	cValue1Array := C.malloc((C.ulong)(param1IndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1IndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1IndirectedLen:param1IndirectedLen]
	for param1Indirectedi, param1IndirectedString := range param1Indirected {
		param1IndirectedSlice[param1Indirectedi] = (*C.gchar)(C.CString(param1IndirectedString))
		defer C.free(unsafe.Pointer(param1IndirectedSlice[param1Indirectedi]))
	}
	if len(param1IndirectedSlice) > 0 {
		cValue1ArrayPointer = &param1IndirectedSlice[0]
	}

	C.gtk_init(cValue0, cValue1)

	param1OutLen := int(*cValue0)
	param1Out := make([]string, param1OutLen, param1OutLen)
	if param1OutLen > 0 {
		param1OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1ArrayPointer))[:param1OutLen:param1OutLen]
		for param1Outi, param1OutCString := range param1OutCSlice {
			param1Out[param1Outi] = C.GoString(param1OutCString)
		}
	}
	*param1 = param1Out
}

func Fn_gtk_init_check(param0 *int, param1 *[]string) {
	cValue0 := (*C.int)(unsafe.Pointer(param0))

	var cValue1ArrayPointer **C.gchar
	cValue1 := &cValue1ArrayPointer
	param1Indirected := *param1
	param1IndirectedLen := len(param1Indirected)
	cValue1Array := C.malloc((C.ulong)(param1IndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1IndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1IndirectedLen:param1IndirectedLen]
	for param1Indirectedi, param1IndirectedString := range param1Indirected {
		param1IndirectedSlice[param1Indirectedi] = (*C.gchar)(C.CString(param1IndirectedString))
		defer C.free(unsafe.Pointer(param1IndirectedSlice[param1Indirectedi]))
	}
	if len(param1IndirectedSlice) > 0 {
		cValue1ArrayPointer = &param1IndirectedSlice[0]
	}

	C.gtk_init_check(cValue0, cValue1)

	param1OutLen := int(*cValue0)
	param1Out := make([]string, param1OutLen, param1OutLen)
	if param1OutLen > 0 {
		param1OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1ArrayPointer))[:param1OutLen:param1OutLen]
		for param1Outi, param1OutCString := range param1OutCSlice {
			param1Out[param1Outi] = C.GoString(param1OutCString)
		}
	}
	*param1 = param1Out
}

func Fn_gtk_init_with_args(param0 *int, param1 *[]string, param2 string, param3 []glib.OptionEntry, param4 string, error unsafe.Pointer) {
	// has non-string array param
}

// UNSUPPORTED : key_snooper_install : has callback

func Fn_gtk_key_snooper_remove(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gtk_key_snooper_remove(cValue0)
}

func Fn_gtk_main() {
	C.gtk_main()
}

func Fn_gtk_main_do_event(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	C.gtk_main_do_event(cValue0)
}

func Fn_gtk_main_iteration() {
	C.gtk_main_iteration()
}

func Fn_gtk_main_iteration_do(param0 bool) {
	cValue0 := toCBool(param0)

	C.gtk_main_iteration_do(cValue0)
}

func Fn_gtk_main_level() {
	C.gtk_main_level()
}

func Fn_gtk_main_quit() {
	C.gtk_main_quit()
}

func Fn_gtk_paint_arrow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 bool, param8 int, param9 int, param10 int, param11 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.GtkArrowType)(param6)

	cValue7 := toCBool(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	cValue10 := (C.gint)(param10)

	cValue11 := (C.gint)(param11)

	C.gtk_paint_arrow(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10, cValue11)
}

func Fn_gtk_paint_box(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	C.gtk_paint_box(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gtk_paint_box_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int, param11 int, param12 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	cValue10 := (C.GtkPositionType)(param10)

	cValue11 := (C.gint)(param11)

	cValue12 := (C.gint)(param12)

	C.gtk_paint_box_gap(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10, cValue11, cValue12)
}

func Fn_gtk_paint_check(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	C.gtk_paint_check(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gtk_paint_diamond(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	C.gtk_paint_diamond(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gtk_paint_expander(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (*C.GtkWidget)(unsafe.Pointer(param3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	cValue5 := (C.gint)(param5)

	cValue6 := (C.gint)(param6)

	cValue7 := (C.GtkExpanderStyle)(param7)

	C.gtk_paint_expander(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7)
}

func Fn_gtk_paint_extension(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	cValue10 := (C.GtkPositionType)(param10)

	C.gtk_paint_extension(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10)
}

func Fn_gtk_paint_flat_box(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	C.gtk_paint_flat_box(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gtk_paint_focus(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int, param8 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (*C.GtkWidget)(unsafe.Pointer(param3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	cValue5 := (C.gint)(param5)

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	C.gtk_paint_focus(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8)
}

func Fn_gtk_paint_handle(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	cValue10 := (C.GtkOrientation)(param10)

	C.gtk_paint_handle(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10)
}

func Fn_gtk_paint_hline(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (*C.GtkWidget)(unsafe.Pointer(param3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	cValue5 := (C.gint)(param5)

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	C.gtk_paint_hline(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7)
}

func Fn_gtk_paint_layout(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 bool, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 unsafe.Pointer) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := toCBool(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (*C.PangoLayout)(unsafe.Pointer(param8))

	C.gtk_paint_layout(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8)
}

func Fn_gtk_paint_option(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	C.gtk_paint_option(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gtk_paint_resize_grip(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int, param8 int, param9 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (*C.GtkWidget)(unsafe.Pointer(param3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	cValue5 := (C.GdkWindowEdge)(param5)

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	C.gtk_paint_resize_grip(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gtk_paint_shadow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	C.gtk_paint_shadow(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gtk_paint_shadow_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int, param11 int, param12 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	cValue10 := (C.GtkPositionType)(param10)

	cValue11 := (C.gint)(param11)

	cValue12 := (C.gint)(param12)

	C.gtk_paint_shadow_gap(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10, cValue11, cValue12)
}

func Fn_gtk_paint_slider(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	cValue10 := (C.GtkOrientation)(param10)

	C.gtk_paint_slider(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9, cValue10)
}

func Fn_gtk_paint_spinner(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 uint, param6 int, param7 int, param8 int, param9 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (*C.GtkWidget)(unsafe.Pointer(param3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	cValue5 := (C.guint)(param5)

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	C.gtk_paint_spinner(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gtk_paint_tab(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkShadowType)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	cValue8 := (C.gint)(param8)

	cValue9 := (C.gint)(param9)

	C.gtk_paint_tab(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8, cValue9)
}

func Fn_gtk_paint_vline(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int) {
	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (*C.GtkWidget)(unsafe.Pointer(param3))

	cValue4 := (*C.gchar)(C.CString(param4))
	defer C.free(unsafe.Pointer(cValue4))

	cValue5 := (C.gint)(param5)

	cValue6 := (C.gint)(param6)

	cValue7 := (C.gint)(param7)

	C.gtk_paint_vline(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7)
}

func Fn_gtk_paper_size_get_default() {
	C.gtk_paper_size_get_default()
}

func Fn_gtk_paper_size_get_paper_sizes(param0 bool) {
	cValue0 := toCBool(param0)

	C.gtk_paper_size_get_paper_sizes(cValue0)
}

func Fn_gtk_parse_args(param0 *int, param1 *[]string) {
	cValue0 := (*C.int)(unsafe.Pointer(param0))

	var cValue1ArrayPointer **C.gchar
	cValue1 := &cValue1ArrayPointer
	param1Indirected := *param1
	param1IndirectedLen := len(param1Indirected)
	cValue1Array := C.malloc((C.ulong)(param1IndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1IndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1IndirectedLen:param1IndirectedLen]
	for param1Indirectedi, param1IndirectedString := range param1Indirected {
		param1IndirectedSlice[param1Indirectedi] = (*C.gchar)(C.CString(param1IndirectedString))
		defer C.free(unsafe.Pointer(param1IndirectedSlice[param1Indirectedi]))
	}
	if len(param1IndirectedSlice) > 0 {
		cValue1ArrayPointer = &param1IndirectedSlice[0]
	}

	C.gtk_parse_args(cValue0, cValue1)

	param1OutLen := int(*cValue0)
	param1Out := make([]string, param1OutLen, param1OutLen)
	if param1OutLen > 0 {
		param1OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1ArrayPointer))[:param1OutLen:param1OutLen]
		for param1Outi, param1OutCString := range param1OutCSlice {
			param1Out[param1Outi] = C.GoString(param1OutCString)
		}
	}
	*param1 = param1Out
}

func Fn_gtk_print_error_quark() {
	C.gtk_print_error_quark()
}

func Fn_gtk_print_run_page_setup_dialog(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkPageSetup)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkPrintSettings)(unsafe.Pointer(param2))

	C.gtk_print_run_page_setup_dialog(cValue0, cValue1, cValue2)
}

// UNSUPPORTED : print_run_page_setup_dialog_async : has callback

func Fn_gtk_propagate_event(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkEvent)(unsafe.Pointer(param1))

	C.gtk_propagate_event(cValue0, cValue1)
}

func Fn_gtk_rc_add_default_file(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_rc_add_default_file(cValue0)
}

func Fn_gtk_rc_find_module_in_path(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_rc_find_module_in_path(cValue0)
}

func Fn_gtk_rc_find_pixmap_in_path(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	cValue1 := (*C.GScanner)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_rc_find_pixmap_in_path(cValue0, cValue1, cValue2)
}

func Fn_gtk_rc_get_default_files() {
	C.gtk_rc_get_default_files()
}

func Fn_gtk_rc_get_im_module_file() {
	C.gtk_rc_get_im_module_file()
}

func Fn_gtk_rc_get_im_module_path() {
	C.gtk_rc_get_im_module_path()
}

func Fn_gtk_rc_get_module_dir() {
	C.gtk_rc_get_module_dir()
}

func Fn_gtk_rc_get_style(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_rc_get_style(cValue0)
}

func Fn_gtk_rc_get_style_by_paths(param0 unsafe.Pointer, param1 string, param2 string, param3 uint64) {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.char)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GType)(param3)

	C.gtk_rc_get_style_by_paths(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_rc_get_theme_dir() {
	C.gtk_rc_get_theme_dir()
}

func Fn_gtk_rc_parse(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_rc_parse(cValue0)
}

func Fn_gtk_rc_parse_color(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gtk_rc_parse_color(cValue0, cValue1)
}

func Fn_gtk_rc_parse_color_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRcStyle)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkColor)(unsafe.Pointer(param2))

	C.gtk_rc_parse_color_full(cValue0, cValue1, cValue2)
}

func Fn_gtk_rc_parse_priority(param0 unsafe.Pointer, param1 *int) {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkPathPriorityType)(unsafe.Pointer(param1))

	C.gtk_rc_parse_priority(cValue0, cValue1)
}

func Fn_gtk_rc_parse_state(param0 unsafe.Pointer, param1 *int) {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkStateType)(unsafe.Pointer(param1))

	C.gtk_rc_parse_state(cValue0, cValue1)
}

func Fn_gtk_rc_parse_string(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_rc_parse_string(cValue0)
}

func Fn_gtk_rc_property_parse_border(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_rc_property_parse_border(cValue0, cValue1, cValue2)
}

func Fn_gtk_rc_property_parse_color(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_rc_property_parse_color(cValue0, cValue1, cValue2)
}

func Fn_gtk_rc_property_parse_enum(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_rc_property_parse_enum(cValue0, cValue1, cValue2)
}

func Fn_gtk_rc_property_parse_flags(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_rc_property_parse_flags(cValue0, cValue1, cValue2)
}

func Fn_gtk_rc_property_parse_requisition(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_rc_property_parse_requisition(cValue0, cValue1, cValue2)
}

func Fn_gtk_rc_reparse_all() {
	C.gtk_rc_reparse_all()
}

func Fn_gtk_rc_reparse_all_for_settings(param0 unsafe.Pointer, param1 bool) {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_rc_reparse_all_for_settings(cValue0, cValue1)
}

func Fn_gtk_rc_reset_styles(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	C.gtk_rc_reset_styles(cValue0)
}

func Fn_gtk_rc_scanner_new() {
	C.gtk_rc_scanner_new()
}

func Fn_gtk_rc_set_default_files(param0 []string) {
	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	C.gtk_rc_set_default_files(cValue0)
}

func Fn_gtk_recent_chooser_error_quark() {
	C.gtk_recent_chooser_error_quark()
}

func Fn_gtk_recent_manager_error_quark() {
	C.gtk_recent_manager_error_quark()
}

func Fn_gtk_rgb_to_hsv(param0 float64, param1 float64, param2 float64, param3 *float64, param4 *float64, param5 *float64) {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (*C.gdouble)(unsafe.Pointer(param3))

	cValue4 := (*C.gdouble)(unsafe.Pointer(param4))

	cValue5 := (*C.gdouble)(unsafe.Pointer(param5))

	C.gtk_rgb_to_hsv(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_selection_add_target(param0 unsafe.Pointer, param1 gdk.Atom, param2 gdk.Atom, param3 uint) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	cValue2 := (C.GdkAtom)(unsafe.Pointer(param2))

	cValue3 := (C.guint)(param3)

	C.gtk_selection_add_target(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_selection_add_targets(param0 unsafe.Pointer, param1 gdk.Atom, param2 []TargetEntry, param3 uint) {
	// has non-string array param
}

func Fn_gtk_selection_clear_targets(param0 unsafe.Pointer, param1 gdk.Atom) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	C.gtk_selection_clear_targets(cValue0, cValue1)
}

func Fn_gtk_selection_convert(param0 unsafe.Pointer, param1 gdk.Atom, param2 gdk.Atom, param3 uint32) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	cValue2 := (C.GdkAtom)(unsafe.Pointer(param2))

	cValue3 := (C.guint32)(param3)

	C.gtk_selection_convert(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_selection_owner_set(param0 unsafe.Pointer, param1 gdk.Atom, param2 uint32) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	cValue2 := (C.guint32)(param2)

	C.gtk_selection_owner_set(cValue0, cValue1, cValue2)
}

func Fn_gtk_selection_owner_set_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 gdk.Atom, param3 uint32) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.GdkAtom)(unsafe.Pointer(param2))

	cValue3 := (C.guint32)(param3)

	C.gtk_selection_owner_set_for_display(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_selection_remove_all(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_selection_remove_all(cValue0)
}

func Fn_gtk_set_debug_flags(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gtk_set_debug_flags(cValue0)
}

// UNSUPPORTED : show_about_dialog : has varargs

func Fn_gtk_show_uri(param0 unsafe.Pointer, param1 string, param2 uint32, error unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.guint32)(param2)

	cError := (**C.GError)(error)

	C.gtk_show_uri(cValue0, cValue1, cValue2, cError)
}

func Fn_gtk_stock_add(param0 []StockItem, param1 uint) {
	// has non-string array param
}

func Fn_gtk_stock_add_static(param0 []StockItem, param1 uint) {
	// has non-string array param
}

func Fn_gtk_stock_list_ids() {
	C.gtk_stock_list_ids()
}

func Fn_gtk_stock_lookup(param0 string, param1 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkStockItem)(unsafe.Pointer(param1))

	C.gtk_stock_lookup(cValue0, cValue1)
}

// UNSUPPORTED : stock_set_translate_func : has callback

func Fn_gtk_target_table_free(param0 []TargetEntry, param1 int) {
	// has non-string array param
}

func Fn_gtk_target_table_new_from_list(param0 unsafe.Pointer, param1 *int) {
	cValue0 := (*C.GtkTargetList)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_target_table_new_from_list(cValue0, cValue1)
}

func Fn_gtk_targets_include_image(param0 []gdk.Atom, param1 int, param2 bool) {
	// has non-string array param
}

func Fn_gtk_targets_include_rich_text(param0 []gdk.Atom, param1 int, param2 unsafe.Pointer) {
	// has non-string array param
}

func Fn_gtk_targets_include_text(param0 []gdk.Atom, param1 int) {
	// has non-string array param
}

func Fn_gtk_targets_include_uri(param0 []gdk.Atom, param1 int) {
	// has non-string array param
}

func Fn_gtk_test_create_simple_window(param0 string, param1 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_test_create_simple_window(cValue0, cValue1)
}

// UNSUPPORTED : test_create_widget : has varargs

// UNSUPPORTED : test_display_button_window : has varargs

func Fn_gtk_test_find_label(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_test_find_label(cValue0, cValue1)
}

func Fn_gtk_test_find_sibling(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	C.gtk_test_find_sibling(cValue0, cValue1)
}

func Fn_gtk_test_find_widget(param0 unsafe.Pointer, param1 string, param2 uint64) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GType)(param2)

	C.gtk_test_find_widget(cValue0, cValue1, cValue2)
}

// UNSUPPORTED : test_init : has varargs

func Fn_gtk_test_list_all_types(param0 *uint) {
	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	C.gtk_test_list_all_types(cValue0)
}

func Fn_gtk_test_register_all_types() {
	C.gtk_test_register_all_types()
}

func Fn_gtk_test_slider_get_value(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_test_slider_get_value(cValue0)
}

func Fn_gtk_test_slider_set_perc(param0 unsafe.Pointer, param1 float64) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.double)(param1)

	C.gtk_test_slider_set_perc(cValue0, cValue1)
}

func Fn_gtk_test_spin_button_click(param0 unsafe.Pointer, param1 uint, param2 bool) {
	cValue0 := (*C.GtkSpinButton)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := toCBool(param2)

	C.gtk_test_spin_button_click(cValue0, cValue1, cValue2)
}

func Fn_gtk_test_text_get(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_test_text_get(cValue0)
}

func Fn_gtk_test_text_set(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_test_text_set(cValue0, cValue1)
}

func Fn_gtk_test_widget_click(param0 unsafe.Pointer, param1 uint, param2 int) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	C.gtk_test_widget_click(cValue0, cValue1, cValue2)
}

func Fn_gtk_test_widget_send_key(param0 unsafe.Pointer, param1 uint, param2 int) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	C.gtk_test_widget_send_key(cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_get_row_drag_data(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer) {
	cValue0 := (*C.GtkSelectionData)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreeModel)(unsafe.Pointer(param1))

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	C.gtk_tree_get_row_drag_data(cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_row_reference_deleted(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_tree_row_reference_deleted(cValue0, cValue1)
}

func Fn_gtk_tree_row_reference_inserted(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_tree_row_reference_inserted(cValue0, cValue1)
}

func Fn_gtk_tree_row_reference_reordered(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 []int) {
	// has non-string array param
}

func Fn_gtk_tree_set_row_drag_data(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.GtkSelectionData)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeModel)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTreePath)(unsafe.Pointer(param2))

	C.gtk_tree_set_row_drag_data(cValue0, cValue1, cValue2)
}

func Fn_gtk_true() {
	C.gtk_true()
}

func Fn_gtk_about_dialog_new() {
	C.gtk_about_dialog_new()
}

func Fn_gtk_about_dialog_get_artists(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_artists(cValueInstance)
}

func Fn_gtk_about_dialog_get_authors(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_authors(cValueInstance)
}

func Fn_gtk_about_dialog_get_comments(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_comments(cValueInstance)
}

func Fn_gtk_about_dialog_get_copyright(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_copyright(cValueInstance)
}

func Fn_gtk_about_dialog_get_documenters(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_documenters(cValueInstance)
}

func Fn_gtk_about_dialog_get_license(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_license(cValueInstance)
}

func Fn_gtk_about_dialog_get_logo(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_logo(cValueInstance)
}

func Fn_gtk_about_dialog_get_logo_icon_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_logo_icon_name(cValueInstance)
}

func Fn_gtk_about_dialog_get_program_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_program_name(cValueInstance)
}

func Fn_gtk_about_dialog_get_translator_credits(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_translator_credits(cValueInstance)
}

func Fn_gtk_about_dialog_get_version(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_version(cValueInstance)
}

func Fn_gtk_about_dialog_get_website(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_website(cValueInstance)
}

func Fn_gtk_about_dialog_get_website_label(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_website_label(cValueInstance)
}

func Fn_gtk_about_dialog_get_wrap_license(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	C.gtk_about_dialog_get_wrap_license(cValueInstance)
}

func Fn_gtk_about_dialog_set_artists(paramInstance unsafe.Pointer, param0 []string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	C.gtk_about_dialog_set_artists(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_authors(paramInstance unsafe.Pointer, param0 []string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	C.gtk_about_dialog_set_authors(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_comments(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_comments(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_copyright(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_copyright(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_documenters(paramInstance unsafe.Pointer, param0 []string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	C.gtk_about_dialog_set_documenters(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_license(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_license(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_logo(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_about_dialog_set_logo(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_logo_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_logo_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_program_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_program_name(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_translator_credits(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_translator_credits(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_version(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_version(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_website(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_website(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_website_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_about_dialog_set_website_label(cValueInstance, cValue0)
}

func Fn_gtk_about_dialog_set_wrap_license(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_about_dialog_set_wrap_license(cValueInstance, cValue0)
}

func Fn_gtk_accel_group_new() {
	C.gtk_accel_group_new()
}

func Fn_gtk_accel_group_activate(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer, param2 uint, param3 int) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (C.GdkModifierType)(param3)

	C.gtk_accel_group_activate(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_accel_group_connect(paramInstance unsafe.Pointer, param0 uint, param1 int, param2 int, param3 unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	cValue2 := (C.GtkAccelFlags)(param2)

	cValue3 := (*C.GClosure)(unsafe.Pointer(param3))

	C.gtk_accel_group_connect(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_accel_group_connect_by_path(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GClosure)(unsafe.Pointer(param1))

	C.gtk_accel_group_connect_by_path(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_accel_group_disconnect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	C.gtk_accel_group_disconnect(cValueInstance, cValue0)
}

func Fn_gtk_accel_group_disconnect_key(paramInstance unsafe.Pointer, param0 uint, param1 int) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	C.gtk_accel_group_disconnect_key(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : find : has callback

func Fn_gtk_accel_group_get_is_locked(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	C.gtk_accel_group_get_is_locked(cValueInstance)
}

func Fn_gtk_accel_group_get_modifier_mask(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	C.gtk_accel_group_get_modifier_mask(cValueInstance)
}

func Fn_gtk_accel_group_lock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	C.gtk_accel_group_lock(cValueInstance)
}

func Fn_gtk_accel_group_query(paramInstance unsafe.Pointer, param0 uint, param1 int, param2 *uint) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	cValue2 := (*C.guint)(unsafe.Pointer(param2))

	C.gtk_accel_group_query(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_accel_group_unlock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	C.gtk_accel_group_unlock(cValueInstance)
}

func Fn_gtk_accel_group_from_accel_closure(param0 unsafe.Pointer) {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	C.gtk_accel_group_from_accel_closure(cValue0)
}

func Fn_gtk_accel_label_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_accel_label_new(cValue0)
}

func Fn_gtk_accel_label_get_accel_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	C.gtk_accel_label_get_accel_widget(cValueInstance)
}

func Fn_gtk_accel_label_get_accel_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	C.gtk_accel_label_get_accel_width(cValueInstance)
}

func Fn_gtk_accel_label_refetch(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	C.gtk_accel_label_refetch(cValueInstance)
}

func Fn_gtk_accel_label_set_accel_closure(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	C.gtk_accel_label_set_accel_closure(cValueInstance, cValue0)
}

func Fn_gtk_accel_label_set_accel_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_accel_label_set_accel_widget(cValueInstance, cValue0)
}

func Fn_gtk_accel_map_add_entry(param0 string, param1 uint, param2 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	C.gtk_accel_map_add_entry(cValue0, cValue1, cValue2)
}

func Fn_gtk_accel_map_add_filter(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_accel_map_add_filter(cValue0)
}

func Fn_gtk_accel_map_change_entry(param0 string, param1 uint, param2 int, param3 bool) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	cValue3 := toCBool(param3)

	C.gtk_accel_map_change_entry(cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : foreach : has callback

// UNSUPPORTED : foreach_unfiltered : has callback

func Fn_gtk_accel_map_get() {
	C.gtk_accel_map_get()
}

func Fn_gtk_accel_map_load(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_accel_map_load(cValue0)
}

func Fn_gtk_accel_map_load_fd(param0 int) {
	cValue0 := (C.gint)(param0)

	C.gtk_accel_map_load_fd(cValue0)
}

func Fn_gtk_accel_map_load_scanner(param0 unsafe.Pointer) {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	C.gtk_accel_map_load_scanner(cValue0)
}

func Fn_gtk_accel_map_lock_path(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_accel_map_lock_path(cValue0)
}

func Fn_gtk_accel_map_lookup_entry(param0 string, param1 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkAccelKey)(unsafe.Pointer(param1))

	C.gtk_accel_map_lookup_entry(cValue0, cValue1)
}

func Fn_gtk_accel_map_save(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_accel_map_save(cValue0)
}

func Fn_gtk_accel_map_save_fd(param0 int) {
	cValue0 := (C.gint)(param0)

	C.gtk_accel_map_save_fd(cValue0)
}

func Fn_gtk_accel_map_unlock_path(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_accel_map_unlock_path(cValue0)
}

func Fn_gtk_accessible_connect_widget_destroyed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccessible)(unsafe.Pointer(paramInstance))

	C.gtk_accessible_connect_widget_destroyed(cValueInstance)
}

func Fn_gtk_action_new(param0 string, param1 string, param2 string, param3 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	C.gtk_action_new(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_action_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_activate(cValueInstance)
}

func Fn_gtk_action_block_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_block_activate(cValueInstance)
}

func Fn_gtk_action_connect_accelerator(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_connect_accelerator(cValueInstance)
}

func Fn_gtk_action_create_icon(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkIconSize)(param0)

	C.gtk_action_create_icon(cValueInstance, cValue0)
}

func Fn_gtk_action_create_menu(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_create_menu(cValueInstance)
}

func Fn_gtk_action_create_menu_item(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_create_menu_item(cValueInstance)
}

func Fn_gtk_action_create_tool_item(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_create_tool_item(cValueInstance)
}

func Fn_gtk_action_disconnect_accelerator(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_disconnect_accelerator(cValueInstance)
}

func Fn_gtk_action_get_accel_closure(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_accel_closure(cValueInstance)
}

func Fn_gtk_action_get_accel_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_accel_path(cValueInstance)
}

func Fn_gtk_action_get_gicon(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_gicon(cValueInstance)
}

func Fn_gtk_action_get_icon_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_icon_name(cValueInstance)
}

func Fn_gtk_action_get_is_important(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_is_important(cValueInstance)
}

func Fn_gtk_action_get_label(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_label(cValueInstance)
}

func Fn_gtk_action_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_name(cValueInstance)
}

func Fn_gtk_action_get_proxies(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_proxies(cValueInstance)
}

func Fn_gtk_action_get_sensitive(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_sensitive(cValueInstance)
}

func Fn_gtk_action_get_short_label(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_short_label(cValueInstance)
}

func Fn_gtk_action_get_stock_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_stock_id(cValueInstance)
}

func Fn_gtk_action_get_tooltip(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_tooltip(cValueInstance)
}

func Fn_gtk_action_get_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_visible(cValueInstance)
}

func Fn_gtk_action_get_visible_horizontal(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_visible_horizontal(cValueInstance)
}

func Fn_gtk_action_get_visible_vertical(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_get_visible_vertical(cValueInstance)
}

func Fn_gtk_action_is_sensitive(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_is_sensitive(cValueInstance)
}

func Fn_gtk_action_is_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_is_visible(cValueInstance)
}

func Fn_gtk_action_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	C.gtk_action_set_accel_group(cValueInstance, cValue0)
}

func Fn_gtk_action_set_accel_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_accel_path(cValueInstance, cValue0)
}

func Fn_gtk_action_set_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.gtk_action_set_gicon(cValueInstance, cValue0)
}

func Fn_gtk_action_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_action_set_is_important(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_is_important(cValueInstance, cValue0)
}

func Fn_gtk_action_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_label(cValueInstance, cValue0)
}

func Fn_gtk_action_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_sensitive(cValueInstance, cValue0)
}

func Fn_gtk_action_set_short_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_short_label(cValueInstance, cValue0)
}

func Fn_gtk_action_set_stock_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_stock_id(cValueInstance, cValue0)
}

func Fn_gtk_action_set_tooltip(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_set_tooltip(cValueInstance, cValue0)
}

func Fn_gtk_action_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_action_set_visible_horizontal(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_visible_horizontal(cValueInstance, cValue0)
}

func Fn_gtk_action_set_visible_vertical(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_visible_vertical(cValueInstance, cValue0)
}

func Fn_gtk_action_unblock_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_unblock_activate(cValueInstance)
}

func Fn_gtk_action_group_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_group_new(cValue0)
}

func Fn_gtk_action_group_add_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_action_group_add_action(cValueInstance, cValue0)
}

func Fn_gtk_action_group_add_action_with_accel(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_action_group_add_action_with_accel(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_action_group_add_actions(paramInstance unsafe.Pointer, param0 []ActionEntry, param1 uint, param2 unsafe.Pointer) {
	// has non-string array param
}

// UNSUPPORTED : add_actions_full : has callback

// UNSUPPORTED : add_radio_actions : has callback

// UNSUPPORTED : add_radio_actions_full : has callback

func Fn_gtk_action_group_add_toggle_actions(paramInstance unsafe.Pointer, param0 []ToggleActionEntry, param1 uint, param2 unsafe.Pointer) {
	// has non-string array param
}

// UNSUPPORTED : add_toggle_actions_full : has callback

func Fn_gtk_action_group_get_action(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_group_get_action(cValueInstance, cValue0)
}

func Fn_gtk_action_group_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	C.gtk_action_group_get_name(cValueInstance)
}

func Fn_gtk_action_group_get_sensitive(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	C.gtk_action_group_get_sensitive(cValueInstance)
}

func Fn_gtk_action_group_get_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	C.gtk_action_group_get_visible(cValueInstance)
}

func Fn_gtk_action_group_list_actions(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	C.gtk_action_group_list_actions(cValueInstance)
}

func Fn_gtk_action_group_remove_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_action_group_remove_action(cValueInstance, cValue0)
}

func Fn_gtk_action_group_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_group_set_sensitive(cValueInstance, cValue0)
}

// UNSUPPORTED : set_translate_func : has callback

func Fn_gtk_action_group_set_translation_domain(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_group_set_translation_domain(cValueInstance, cValue0)
}

func Fn_gtk_action_group_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_group_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_action_group_translate_string(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_action_group_translate_string(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_new(param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_adjustment_new(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_adjustment_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	C.gtk_adjustment_changed(cValueInstance)
}

func Fn_gtk_adjustment_clamp_page(paramInstance unsafe.Pointer, param0 float64, param1 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	C.gtk_adjustment_clamp_page(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_adjustment_configure(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_adjustment_configure(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_adjustment_get_lower(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	C.gtk_adjustment_get_lower(cValueInstance)
}

func Fn_gtk_adjustment_get_page_increment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	C.gtk_adjustment_get_page_increment(cValueInstance)
}

func Fn_gtk_adjustment_get_page_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	C.gtk_adjustment_get_page_size(cValueInstance)
}

func Fn_gtk_adjustment_get_step_increment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	C.gtk_adjustment_get_step_increment(cValueInstance)
}

func Fn_gtk_adjustment_get_upper(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	C.gtk_adjustment_get_upper(cValueInstance)
}

func Fn_gtk_adjustment_get_value(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	C.gtk_adjustment_get_value(cValueInstance)
}

func Fn_gtk_adjustment_set_lower(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_lower(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_set_page_increment(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_page_increment(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_set_page_size(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_page_size(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_set_step_increment(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_step_increment(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_set_upper(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_upper(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_set_value(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_adjustment_set_value(cValueInstance, cValue0)
}

func Fn_gtk_adjustment_value_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	C.gtk_adjustment_value_changed(cValueInstance)
}

func Fn_gtk_alignment_new(param0 float32, param1 float32, param2 float32, param3 float32) {
	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	cValue2 := (C.gfloat)(param2)

	cValue3 := (C.gfloat)(param3)

	C.gtk_alignment_new(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_alignment_get_padding(paramInstance unsafe.Pointer, param0 *uint, param1 *uint, param2 *uint, param3 *uint) {
	cValueInstance := (*C.GtkAlignment)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	cValue2 := (*C.guint)(unsafe.Pointer(param2))

	cValue3 := (*C.guint)(unsafe.Pointer(param3))

	C.gtk_alignment_get_padding(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_alignment_set(paramInstance unsafe.Pointer, param0 float32, param1 float32, param2 float32, param3 float32) {
	cValueInstance := (*C.GtkAlignment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	cValue2 := (C.gfloat)(param2)

	cValue3 := (C.gfloat)(param3)

	C.gtk_alignment_set(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_alignment_set_padding(paramInstance unsafe.Pointer, param0 uint, param1 uint, param2 uint, param3 uint) {
	cValueInstance := (*C.GtkAlignment)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.guint)(param3)

	C.gtk_alignment_set_padding(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_app_chooser_button_get_heading(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	C.gtk_app_chooser_button_get_heading(cValueInstance)
}

func Fn_gtk_app_chooser_button_set_heading(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_app_chooser_button_set_heading(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_dialog_get_heading(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAppChooserDialog)(unsafe.Pointer(paramInstance))

	C.gtk_app_chooser_dialog_get_heading(cValueInstance)
}

func Fn_gtk_app_chooser_dialog_set_heading(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAppChooserDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_app_chooser_dialog_set_heading(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_default_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_app_chooser_widget_set_default_text(cValueInstance, cValue0)
}

func Fn_gtk_arrow_new(param0 int, param1 int) {
	cValue0 := (C.GtkArrowType)(param0)

	cValue1 := (C.GtkShadowType)(param1)

	C.gtk_arrow_new(cValue0, cValue1)
}

func Fn_gtk_arrow_set(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkArrow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkArrowType)(param0)

	cValue1 := (C.GtkShadowType)(param1)

	C.gtk_arrow_set(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_aspect_frame_new(param0 string, param1 float32, param2 float32, param3 float32, param4 bool) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gfloat)(param1)

	cValue2 := (C.gfloat)(param2)

	cValue3 := (C.gfloat)(param3)

	cValue4 := toCBool(param4)

	C.gtk_aspect_frame_new(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_aspect_frame_set(paramInstance unsafe.Pointer, param0 float32, param1 float32, param2 float32, param3 bool) {
	cValueInstance := (*C.GtkAspectFrame)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	cValue2 := (C.gfloat)(param2)

	cValue3 := toCBool(param3)

	C.gtk_aspect_frame_set(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_assistant_new() {
	C.gtk_assistant_new()
}

func Fn_gtk_assistant_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_add_action_widget(cValueInstance, cValue0)
}

func Fn_gtk_assistant_append_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_append_page(cValueInstance, cValue0)
}

func Fn_gtk_assistant_get_current_page(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	C.gtk_assistant_get_current_page(cValueInstance)
}

func Fn_gtk_assistant_get_n_pages(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	C.gtk_assistant_get_n_pages(cValueInstance)
}

func Fn_gtk_assistant_get_nth_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_assistant_get_nth_page(cValueInstance, cValue0)
}

func Fn_gtk_assistant_get_page_complete(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_get_page_complete(cValueInstance, cValue0)
}

func Fn_gtk_assistant_get_page_header_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_get_page_header_image(cValueInstance, cValue0)
}

func Fn_gtk_assistant_get_page_side_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_get_page_side_image(cValueInstance, cValue0)
}

func Fn_gtk_assistant_get_page_title(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_get_page_title(cValueInstance, cValue0)
}

func Fn_gtk_assistant_get_page_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_get_page_type(cValueInstance, cValue0)
}

func Fn_gtk_assistant_insert_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_assistant_insert_page(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_prepend_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_prepend_page(cValueInstance, cValue0)
}

func Fn_gtk_assistant_remove_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_remove_action_widget(cValueInstance, cValue0)
}

func Fn_gtk_assistant_set_current_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_assistant_set_current_page(cValueInstance, cValue0)
}

// UNSUPPORTED : set_forward_page_func : has callback

func Fn_gtk_assistant_set_page_complete(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_assistant_set_page_complete(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_set_page_header_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	C.gtk_assistant_set_page_header_image(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_set_page_side_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	C.gtk_assistant_set_page_side_image(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_set_page_title(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_assistant_set_page_title(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_set_page_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkAssistantPageType)(param1)

	C.gtk_assistant_set_page_type(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_assistant_update_buttons_state(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	C.gtk_assistant_update_buttons_state(cValueInstance)
}

func Fn_gtk_bin_get_child(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkBin)(unsafe.Pointer(paramInstance))

	C.gtk_bin_get_child(cValueInstance)
}

func Fn_gtk_box_get_homogeneous(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	C.gtk_box_get_homogeneous(cValueInstance)
}

func Fn_gtk_box_get_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	C.gtk_box_get_spacing(cValueInstance)
}

func Fn_gtk_box_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.guint)(param3)

	C.gtk_box_pack_end(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_box_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.guint)(param3)

	C.gtk_box_pack_start(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_box_query_child_packing(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *bool, param2 *bool, param3 *uint, param4 *int) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gboolean)(unsafe.Pointer(param1))

	cValue2 := (*C.gboolean)(unsafe.Pointer(param2))

	cValue3 := (*C.guint)(unsafe.Pointer(param3))

	cValue4 := (*C.GtkPackType)(unsafe.Pointer(param4))

	C.gtk_box_query_child_packing(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_box_reorder_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_box_reorder_child(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_box_set_child_packing(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint, param4 int) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.guint)(param3)

	cValue4 := (C.GtkPackType)(param4)

	C.gtk_box_set_child_packing(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_box_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_box_set_homogeneous(cValueInstance, cValue0)
}

func Fn_gtk_box_set_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_box_set_spacing(cValueInstance, cValue0)
}

func Fn_gtk_builder_new() {
	C.gtk_builder_new()
}

// UNSUPPORTED : add_callback_symbol : has callback

// UNSUPPORTED : add_callback_symbols : has varargs

func Fn_gtk_builder_add_from_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_builder_add_from_file(cValueInstance, cValue0, cError)
}

func Fn_gtk_builder_add_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64, error unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gsize)(param1)

	cError := (**C.GError)(error)

	C.gtk_builder_add_from_string(cValueInstance, cValue0, cValue1, cError)
}

func Fn_gtk_builder_add_objects_from_file(paramInstance unsafe.Pointer, param0 string, param1 []string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	param1Len := len(param1)
	cValue1Array := C.malloc((C.ulong)(param1Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue1Array))
	param1Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1Array))[:param1Len:param1Len]
	for param1i, param1String := range param1 {
		param1Slice[param1i] = (*C.gchar)(C.CString(param1String))
		defer C.free(unsafe.Pointer(param1Slice[param1i]))
	}
	cValue1 := &param1Slice[0]

	cError := (**C.GError)(error)

	C.gtk_builder_add_objects_from_file(cValueInstance, cValue0, cValue1, cError)
}

func Fn_gtk_builder_add_objects_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64, param2 []string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gsize)(param1)

	param2Len := len(param2)
	cValue2Array := C.malloc((C.ulong)(param2Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue2Array))
	param2Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue2Array))[:param2Len:param2Len]
	for param2i, param2String := range param2 {
		param2Slice[param2i] = (*C.gchar)(C.CString(param2String))
		defer C.free(unsafe.Pointer(param2Slice[param2i]))
	}
	cValue2 := &param2Slice[0]

	cError := (**C.GError)(error)

	C.gtk_builder_add_objects_from_string(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_gtk_builder_connect_signals(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.gtk_builder_connect_signals(cValueInstance, cValue0)
}

// UNSUPPORTED : connect_signals_full : has callback

func Fn_gtk_builder_extend_with_template(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint64, param2 string, param3 uint64, error unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gsize)(param3)

	cError := (**C.GError)(error)

	C.gtk_builder_extend_with_template(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)
}

func Fn_gtk_builder_get_object(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_builder_get_object(cValueInstance, cValue0)
}

func Fn_gtk_builder_get_objects(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	C.gtk_builder_get_objects(cValueInstance)
}

func Fn_gtk_builder_get_translation_domain(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	C.gtk_builder_get_translation_domain(cValueInstance)
}

func Fn_gtk_builder_get_type_from_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_builder_get_type_from_name(cValueInstance, cValue0)
}

func Fn_gtk_builder_set_translation_domain(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_builder_set_translation_domain(cValueInstance, cValue0)
}

func Fn_gtk_builder_value_from_string(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	C.gtk_builder_value_from_string(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_gtk_builder_value_from_string_type(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	C.gtk_builder_value_from_string_type(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_gtk_button_new() {
	C.gtk_button_new()
}

func Fn_gtk_button_new_from_stock(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_button_new_from_stock(cValue0)
}

func Fn_gtk_button_new_with_label(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_button_new_with_label(cValue0)
}

func Fn_gtk_button_new_with_mnemonic(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_button_new_with_mnemonic(cValue0)
}

func Fn_gtk_button_clicked(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_clicked(cValueInstance)
}

func Fn_gtk_button_enter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_enter(cValueInstance)
}

func Fn_gtk_button_get_alignment(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gfloat)(unsafe.Pointer(param0))

	cValue1 := (*C.gfloat)(unsafe.Pointer(param1))

	C.gtk_button_get_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_button_get_focus_on_click(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_get_focus_on_click(cValueInstance)
}

func Fn_gtk_button_get_image(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_get_image(cValueInstance)
}

func Fn_gtk_button_get_image_position(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_get_image_position(cValueInstance)
}

func Fn_gtk_button_get_label(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_get_label(cValueInstance)
}

func Fn_gtk_button_get_relief(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_get_relief(cValueInstance)
}

func Fn_gtk_button_get_use_stock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_get_use_stock(cValueInstance)
}

func Fn_gtk_button_get_use_underline(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_get_use_underline(cValueInstance)
}

func Fn_gtk_button_leave(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_leave(cValueInstance)
}

func Fn_gtk_button_pressed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_pressed(cValueInstance)
}

func Fn_gtk_button_released(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	C.gtk_button_released(cValueInstance)
}

func Fn_gtk_button_set_alignment(paramInstance unsafe.Pointer, param0 float32, param1 float32) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	C.gtk_button_set_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_button_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_button_set_focus_on_click(cValueInstance, cValue0)
}

func Fn_gtk_button_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_button_set_image(cValueInstance, cValue0)
}

func Fn_gtk_button_set_image_position(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPositionType)(param0)

	C.gtk_button_set_image_position(cValueInstance, cValue0)
}

func Fn_gtk_button_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_button_set_label(cValueInstance, cValue0)
}

func Fn_gtk_button_set_relief(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkReliefStyle)(param0)

	C.gtk_button_set_relief(cValueInstance, cValue0)
}

func Fn_gtk_button_set_use_stock(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_button_set_use_stock(cValueInstance, cValue0)
}

func Fn_gtk_button_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_button_set_use_underline(cValueInstance, cValue0)
}

func Fn_gtk_button_box_get_child_secondary(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkButtonBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_button_box_get_child_secondary(cValueInstance, cValue0)
}

func Fn_gtk_button_box_get_layout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkButtonBox)(unsafe.Pointer(paramInstance))

	C.gtk_button_box_get_layout(cValueInstance)
}

func Fn_gtk_button_box_set_child_secondary(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkButtonBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_button_box_set_child_secondary(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_button_box_set_layout(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkButtonBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkButtonBoxStyle)(param0)

	C.gtk_button_box_set_layout(cValueInstance, cValue0)
}

func Fn_gtk_calendar_new() {
	C.gtk_calendar_new()
}

func Fn_gtk_calendar_clear_marks(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	C.gtk_calendar_clear_marks(cValueInstance)
}

func Fn_gtk_calendar_get_date(paramInstance unsafe.Pointer, param0 *uint, param1 *uint, param2 *uint) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	cValue2 := (*C.guint)(unsafe.Pointer(param2))

	C.gtk_calendar_get_date(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_calendar_get_detail_height_rows(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	C.gtk_calendar_get_detail_height_rows(cValueInstance)
}

func Fn_gtk_calendar_get_detail_width_chars(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	C.gtk_calendar_get_detail_width_chars(cValueInstance)
}

func Fn_gtk_calendar_get_display_options(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	C.gtk_calendar_get_display_options(cValueInstance)
}

func Fn_gtk_calendar_mark_day(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_calendar_mark_day(cValueInstance, cValue0)
}

func Fn_gtk_calendar_select_day(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_calendar_select_day(cValueInstance, cValue0)
}

func Fn_gtk_calendar_select_month(paramInstance unsafe.Pointer, param0 uint, param1 uint) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	C.gtk_calendar_select_month(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : set_detail_func : has callback

func Fn_gtk_calendar_set_detail_height_rows(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_calendar_set_detail_height_rows(cValueInstance, cValue0)
}

func Fn_gtk_calendar_set_detail_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_calendar_set_detail_width_chars(cValueInstance, cValue0)
}

func Fn_gtk_calendar_set_display_options(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkCalendarDisplayOptions)(param0)

	C.gtk_calendar_set_display_options(cValueInstance, cValue0)
}

func Fn_gtk_calendar_unmark_day(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_calendar_unmark_day(cValueInstance, cValue0)
}

// UNSUPPORTED : add_with_properties : has varargs

// UNSUPPORTED : cell_get : has varargs

// UNSUPPORTED : cell_get_valist : has va_list

// UNSUPPORTED : cell_set : has varargs

// UNSUPPORTED : cell_set_valist : has va_list

// UNSUPPORTED : foreach : has callback

// UNSUPPORTED : foreach_alloc : has callback

func Fn_gtk_cell_area_context_allocate(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_cell_area_context_allocate(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_reset(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	C.gtk_cell_area_context_reset(cValueInstance)
}

func Fn_gtk_cell_renderer_activate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (*C.GdkRectangle)(unsafe.Pointer(param4))

	cValue5 := (C.GtkCellRendererState)(param5)

	C.gtk_cell_renderer_activate(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_cell_renderer_get_alignment(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gfloat)(unsafe.Pointer(param0))

	cValue1 := (*C.gfloat)(unsafe.Pointer(param1))

	C.gtk_cell_renderer_get_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_get_fixed_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_cell_renderer_get_fixed_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_get_padding(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_cell_renderer_get_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_get_sensitive(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	C.gtk_cell_renderer_get_sensitive(cValueInstance)
}

func Fn_gtk_cell_renderer_get_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int, param4 *int, param5 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (*C.gint)(unsafe.Pointer(param5))

	C.gtk_cell_renderer_get_size(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_cell_renderer_get_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	C.gtk_cell_renderer_get_visible(cValueInstance)
}

func Fn_gtk_cell_renderer_render(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (C.GtkCellRendererState)(param4)

	C.gtk_cell_renderer_render(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_cell_renderer_set_alignment(paramInstance unsafe.Pointer, param0 float32, param1 float32) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	C.gtk_cell_renderer_set_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_set_fixed_size(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_cell_renderer_set_fixed_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_set_padding(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_cell_renderer_set_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_renderer_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_set_sensitive(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_start_editing(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (*C.GdkRectangle)(unsafe.Pointer(param4))

	cValue5 := (C.GtkCellRendererState)(param5)

	C.gtk_cell_renderer_start_editing(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_cell_renderer_stop_editing(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_stop_editing(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_accel_new() {
	C.gtk_cell_renderer_accel_new()
}

func Fn_gtk_cell_renderer_combo_new() {
	C.gtk_cell_renderer_combo_new()
}

func Fn_gtk_cell_renderer_pixbuf_new() {
	C.gtk_cell_renderer_pixbuf_new()
}

func Fn_gtk_cell_renderer_progress_new() {
	C.gtk_cell_renderer_progress_new()
}

func Fn_gtk_cell_renderer_spin_new() {
	C.gtk_cell_renderer_spin_new()
}

func Fn_gtk_cell_renderer_text_new() {
	C.gtk_cell_renderer_text_new()
}

func Fn_gtk_cell_renderer_text_set_fixed_height_from_font(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCellRendererText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_cell_renderer_text_set_fixed_height_from_font(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_toggle_new() {
	C.gtk_cell_renderer_toggle_new()
}

func Fn_gtk_cell_renderer_toggle_get_activatable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	C.gtk_cell_renderer_toggle_get_activatable(cValueInstance)
}

func Fn_gtk_cell_renderer_toggle_get_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	C.gtk_cell_renderer_toggle_get_active(cValueInstance)
}

func Fn_gtk_cell_renderer_toggle_get_radio(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	C.gtk_cell_renderer_toggle_get_radio(cValueInstance)
}

func Fn_gtk_cell_renderer_toggle_set_activatable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_toggle_set_activatable(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_toggle_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_toggle_set_active(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_toggle_set_radio(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_toggle_set_radio(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_new() {
	C.gtk_cell_view_new()
}

func Fn_gtk_cell_view_new_with_context(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellAreaContext)(unsafe.Pointer(param1))

	C.gtk_cell_view_new_with_context(cValue0, cValue1)
}

func Fn_gtk_cell_view_new_with_markup(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_cell_view_new_with_markup(cValue0)
}

func Fn_gtk_cell_view_new_with_pixbuf(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_cell_view_new_with_pixbuf(cValue0)
}

func Fn_gtk_cell_view_new_with_text(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_cell_view_new_with_text(cValue0)
}

func Fn_gtk_cell_view_get_displayed_row(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	C.gtk_cell_view_get_displayed_row(cValueInstance)
}

func Fn_gtk_cell_view_get_model(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	C.gtk_cell_view_get_model(cValueInstance)
}

func Fn_gtk_cell_view_get_size_of_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRequisition)(unsafe.Pointer(param1))

	C.gtk_cell_view_get_size_of_row(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_view_set_background_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_background_color(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_displayed_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_displayed_row(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_model(cValueInstance, cValue0)
}

func Fn_gtk_check_button_new() {
	C.gtk_check_button_new()
}

func Fn_gtk_check_button_new_with_label(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_check_button_new_with_label(cValue0)
}

func Fn_gtk_check_button_new_with_mnemonic(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_check_button_new_with_mnemonic(cValue0)
}

func Fn_gtk_check_menu_item_new() {
	C.gtk_check_menu_item_new()
}

func Fn_gtk_check_menu_item_new_with_label(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_check_menu_item_new_with_label(cValue0)
}

func Fn_gtk_check_menu_item_new_with_mnemonic(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_check_menu_item_new_with_mnemonic(cValue0)
}

func Fn_gtk_check_menu_item_get_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_check_menu_item_get_active(cValueInstance)
}

func Fn_gtk_check_menu_item_get_draw_as_radio(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_check_menu_item_get_draw_as_radio(cValueInstance)
}

func Fn_gtk_check_menu_item_get_inconsistent(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_check_menu_item_get_inconsistent(cValueInstance)
}

func Fn_gtk_check_menu_item_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_check_menu_item_set_active(cValueInstance, cValue0)
}

func Fn_gtk_check_menu_item_set_draw_as_radio(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_check_menu_item_set_draw_as_radio(cValueInstance, cValue0)
}

func Fn_gtk_check_menu_item_set_inconsistent(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_check_menu_item_set_inconsistent(cValueInstance, cValue0)
}

func Fn_gtk_check_menu_item_toggled(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_check_menu_item_toggled(cValueInstance)
}

func Fn_gtk_clipboard_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_clear(cValueInstance)
}

func Fn_gtk_clipboard_get_display(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_get_display(cValueInstance)
}

func Fn_gtk_clipboard_get_owner(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_get_owner(cValueInstance)
}

// UNSUPPORTED : request_contents : has callback

// UNSUPPORTED : request_image : has callback

// UNSUPPORTED : request_rich_text : has callback

// UNSUPPORTED : request_targets : has callback

// UNSUPPORTED : request_text : has callback

// UNSUPPORTED : request_uris : has callback

func Fn_gtk_clipboard_set_can_store(paramInstance unsafe.Pointer, param0 []TargetEntry, param1 int) {
	// has non-string array param
}

func Fn_gtk_clipboard_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_clipboard_set_image(cValueInstance, cValue0)
}

func Fn_gtk_clipboard_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_clipboard_set_text(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : set_with_data : has callback

// UNSUPPORTED : set_with_owner : has callback

func Fn_gtk_clipboard_store(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_store(cValueInstance)
}

func Fn_gtk_clipboard_wait_for_contents(paramInstance unsafe.Pointer, param0 gdk.Atom) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	C.gtk_clipboard_wait_for_contents(cValueInstance, cValue0)
}

func Fn_gtk_clipboard_wait_for_image(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_wait_for_image(cValueInstance)
}

func Fn_gtk_clipboard_wait_for_rich_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *uint64) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkAtom)(unsafe.Pointer(param1))

	cValue2 := (*C.gsize)(unsafe.Pointer(param2))

	C.gtk_clipboard_wait_for_rich_text(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_clipboard_wait_for_targets(paramInstance unsafe.Pointer, param0 []unsafe.Pointer, param1 *int) {
	// has non-string array param
}

func Fn_gtk_clipboard_wait_for_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_wait_for_text(cValueInstance)
}

func Fn_gtk_clipboard_wait_for_uris(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_wait_for_uris(cValueInstance)
}

func Fn_gtk_clipboard_wait_is_image_available(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_wait_is_image_available(cValueInstance)
}

func Fn_gtk_clipboard_wait_is_rich_text_available(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	C.gtk_clipboard_wait_is_rich_text_available(cValueInstance, cValue0)
}

func Fn_gtk_clipboard_wait_is_target_available(paramInstance unsafe.Pointer, param0 gdk.Atom) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	C.gtk_clipboard_wait_is_target_available(cValueInstance, cValue0)
}

func Fn_gtk_clipboard_wait_is_text_available(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_wait_is_text_available(cValueInstance)
}

func Fn_gtk_clipboard_wait_is_uris_available(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_wait_is_uris_available(cValueInstance)
}

func Fn_gtk_clipboard_get(param0 gdk.Atom) {
	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	C.gtk_clipboard_get(cValue0)
}

func Fn_gtk_clipboard_get_for_display(param0 unsafe.Pointer, param1 gdk.Atom) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	C.gtk_clipboard_get_for_display(cValue0, cValue1)
}

func Fn_gtk_color_button_new() {
	C.gtk_color_button_new()
}

func Fn_gtk_color_button_new_with_color(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_button_new_with_color(cValue0)
}

func Fn_gtk_color_button_get_alpha(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	C.gtk_color_button_get_alpha(cValueInstance)
}

func Fn_gtk_color_button_get_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_button_get_color(cValueInstance, cValue0)
}

func Fn_gtk_color_button_get_title(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	C.gtk_color_button_get_title(cValueInstance)
}

func Fn_gtk_color_button_get_use_alpha(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	C.gtk_color_button_get_use_alpha(cValueInstance)
}

func Fn_gtk_color_button_set_alpha(paramInstance unsafe.Pointer, param0 uint16) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint16)(param0)

	C.gtk_color_button_set_alpha(cValueInstance, cValue0)
}

func Fn_gtk_color_button_set_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_button_set_color(cValueInstance, cValue0)
}

func Fn_gtk_color_button_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_color_button_set_title(cValueInstance, cValue0)
}

func Fn_gtk_color_button_set_use_alpha(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_color_button_set_use_alpha(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_new() {
	C.gtk_color_selection_new()
}

func Fn_gtk_color_selection_get_current_alpha(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	C.gtk_color_selection_get_current_alpha(cValueInstance)
}

func Fn_gtk_color_selection_get_current_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_selection_get_current_color(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_get_has_opacity_control(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	C.gtk_color_selection_get_has_opacity_control(cValueInstance)
}

func Fn_gtk_color_selection_get_has_palette(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	C.gtk_color_selection_get_has_palette(cValueInstance)
}

func Fn_gtk_color_selection_get_previous_alpha(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	C.gtk_color_selection_get_previous_alpha(cValueInstance)
}

func Fn_gtk_color_selection_get_previous_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_selection_get_previous_color(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_is_adjusting(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	C.gtk_color_selection_is_adjusting(cValueInstance)
}

func Fn_gtk_color_selection_set_current_alpha(paramInstance unsafe.Pointer, param0 uint16) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint16)(param0)

	C.gtk_color_selection_set_current_alpha(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_set_current_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_selection_set_current_color(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_set_has_opacity_control(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_color_selection_set_has_opacity_control(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_set_has_palette(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_color_selection_set_has_palette(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_set_previous_alpha(paramInstance unsafe.Pointer, param0 uint16) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint16)(param0)

	C.gtk_color_selection_set_previous_alpha(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_set_previous_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_selection_set_previous_color(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_palette_from_string(param0 string, param1 []unsafe.Pointer, param2 *int) {
	// has non-string array param
}

func Fn_gtk_color_selection_palette_to_string(param0 []gdk.Color, param1 int) {
	// has non-string array param
}

// UNSUPPORTED : set_change_palette_with_screen_hook : has callback

func Fn_gtk_color_selection_dialog_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_color_selection_dialog_new(cValue0)
}

func Fn_gtk_color_selection_dialog_get_color_selection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelectionDialog)(unsafe.Pointer(paramInstance))

	C.gtk_color_selection_dialog_get_color_selection(cValueInstance)
}

func Fn_gtk_combo_box_new() {
	C.gtk_combo_box_new()
}

func Fn_gtk_combo_box_new_with_area(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	C.gtk_combo_box_new_with_area(cValue0)
}

func Fn_gtk_combo_box_new_with_area_and_entry(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	C.gtk_combo_box_new_with_area_and_entry(cValue0)
}

func Fn_gtk_combo_box_new_with_model(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_combo_box_new_with_model(cValue0)
}

func Fn_gtk_combo_box_get_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_active(cValueInstance)
}

func Fn_gtk_combo_box_get_active_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_combo_box_get_active_iter(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_get_add_tearoffs(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_add_tearoffs(cValueInstance)
}

func Fn_gtk_combo_box_get_button_sensitivity(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_button_sensitivity(cValueInstance)
}

func Fn_gtk_combo_box_get_column_span_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_column_span_column(cValueInstance)
}

func Fn_gtk_combo_box_get_focus_on_click(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_focus_on_click(cValueInstance)
}

func Fn_gtk_combo_box_get_model(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_model(cValueInstance)
}

func Fn_gtk_combo_box_get_popup_accessible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_popup_accessible(cValueInstance)
}

func Fn_gtk_combo_box_get_row_separator_func(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_row_separator_func(cValueInstance)
}

func Fn_gtk_combo_box_get_row_span_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_row_span_column(cValueInstance)
}

func Fn_gtk_combo_box_get_title(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_title(cValueInstance)
}

func Fn_gtk_combo_box_get_wrap_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_get_wrap_width(cValueInstance)
}

func Fn_gtk_combo_box_popdown(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_popdown(cValueInstance)
}

func Fn_gtk_combo_box_popup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_popup(cValueInstance)
}

func Fn_gtk_combo_box_set_active(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_active(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_active_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_combo_box_set_active_iter(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_add_tearoffs(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_combo_box_set_add_tearoffs(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_button_sensitivity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSensitivityType)(param0)

	C.gtk_combo_box_set_button_sensitivity(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_column_span_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_column_span_column(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_combo_box_set_focus_on_click(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_combo_box_set_model(cValueInstance, cValue0)
}

// UNSUPPORTED : set_row_separator_func : has callback

func Fn_gtk_combo_box_set_row_span_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_row_span_column(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_combo_box_set_title(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_wrap_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_wrap_width(cValueInstance, cValue0)
}

func Fn_gtk_container_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_container_add(cValueInstance, cValue0)
}

// UNSUPPORTED : add_with_properties : has varargs

func Fn_gtk_container_check_resize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_check_resize(cValueInstance)
}

// UNSUPPORTED : child_get : has varargs

func Fn_gtk_container_child_get_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_container_child_get_property(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : child_get_valist : has va_list

// UNSUPPORTED : child_set : has varargs

func Fn_gtk_container_child_set_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_container_child_set_property(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : child_set_valist : has va_list

func Fn_gtk_container_child_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_child_type(cValueInstance)
}

// UNSUPPORTED : forall : has callback

// UNSUPPORTED : foreach : has callback

func Fn_gtk_container_get_border_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_get_border_width(cValueInstance)
}

func Fn_gtk_container_get_children(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_get_children(cValueInstance)
}

func Fn_gtk_container_get_focus_chain(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GList)(unsafe.Pointer(param0))

	C.gtk_container_get_focus_chain(cValueInstance, cValue0)
}

func Fn_gtk_container_get_focus_child(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_get_focus_child(cValueInstance)
}

func Fn_gtk_container_get_focus_hadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_get_focus_hadjustment(cValueInstance)
}

func Fn_gtk_container_get_focus_vadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_get_focus_vadjustment(cValueInstance)
}

func Fn_gtk_container_get_path_for_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_container_get_path_for_child(cValueInstance, cValue0)
}

func Fn_gtk_container_get_resize_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_get_resize_mode(cValueInstance)
}

func Fn_gtk_container_propagate_draw(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	C.gtk_container_propagate_draw(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_container_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_container_remove(cValueInstance, cValue0)
}

func Fn_gtk_container_resize_children(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_resize_children(cValueInstance)
}

func Fn_gtk_container_set_border_width(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_container_set_border_width(cValueInstance, cValue0)
}

func Fn_gtk_container_set_focus_chain(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.gtk_container_set_focus_chain(cValueInstance, cValue0)
}

func Fn_gtk_container_set_focus_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_container_set_focus_child(cValueInstance, cValue0)
}

func Fn_gtk_container_set_focus_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_container_set_focus_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_container_set_focus_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_container_set_focus_vadjustment(cValueInstance, cValue0)
}

func Fn_gtk_container_set_reallocate_redraws(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_container_set_reallocate_redraws(cValueInstance, cValue0)
}

func Fn_gtk_container_set_resize_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkResizeMode)(param0)

	C.gtk_container_set_resize_mode(cValueInstance, cValue0)
}

func Fn_gtk_container_unset_focus_chain(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_unset_focus_chain(cValueInstance)
}

func Fn_gtk_container_cell_accessible_new() {
	C.gtk_container_cell_accessible_new()
}

func Fn_gtk_container_cell_accessible_add_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainerCellAccessible)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	C.gtk_container_cell_accessible_add_child(cValueInstance, cValue0)
}

func Fn_gtk_container_cell_accessible_get_children(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainerCellAccessible)(unsafe.Pointer(paramInstance))

	C.gtk_container_cell_accessible_get_children(cValueInstance)
}

func Fn_gtk_container_cell_accessible_remove_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainerCellAccessible)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	C.gtk_container_cell_accessible_remove_child(cValueInstance, cValue0)
}

func Fn_gtk_css_provider_new() {
	C.gtk_css_provider_new()
}

func Fn_gtk_css_provider_load_from_data(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64, error unsafe.Pointer) {
	// has non-string array param
}

func Fn_gtk_css_provider_load_from_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GtkCssProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	C.gtk_css_provider_load_from_file(cValueInstance, cValue0, cError)
}

func Fn_gtk_css_provider_load_from_path(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkCssProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_css_provider_load_from_path(cValueInstance, cValue0, cError)
}

func Fn_gtk_css_provider_get_default() {
	C.gtk_css_provider_get_default()
}

func Fn_gtk_css_provider_get_named(param0 string, param1 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_css_provider_get_named(cValue0, cValue1)
}

func Fn_gtk_dialog_new() {
	C.gtk_dialog_new()
}

// UNSUPPORTED : new_with_buttons : has varargs

func Fn_gtk_dialog_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_dialog_add_action_widget(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_dialog_add_button(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_dialog_add_button(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : add_buttons : has varargs

func Fn_gtk_dialog_get_action_area(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	C.gtk_dialog_get_action_area(cValueInstance)
}

func Fn_gtk_dialog_get_content_area(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	C.gtk_dialog_get_content_area(cValueInstance)
}

func Fn_gtk_dialog_get_response_for_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_dialog_get_response_for_widget(cValueInstance, cValue0)
}

func Fn_gtk_dialog_response(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_dialog_response(cValueInstance, cValue0)
}

func Fn_gtk_dialog_run(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	C.gtk_dialog_run(cValueInstance)
}

// UNSUPPORTED : set_alternative_button_order : has varargs

func Fn_gtk_dialog_set_alternative_button_order_from_array(paramInstance unsafe.Pointer, param0 int, param1 []int) {
	// has non-string array param
}

func Fn_gtk_dialog_set_default_response(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_dialog_set_default_response(cValueInstance, cValue0)
}

func Fn_gtk_dialog_set_response_sensitive(paramInstance unsafe.Pointer, param0 int, param1 bool) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := toCBool(param1)

	C.gtk_dialog_set_response_sensitive(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_drawing_area_new() {
	C.gtk_drawing_area_new()
}

func Fn_gtk_entry_new() {
	C.gtk_entry_new()
}

func Fn_gtk_entry_new_with_buffer(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkEntryBuffer)(unsafe.Pointer(param0))

	C.gtk_entry_new_with_buffer(cValue0)
}

func Fn_gtk_entry_get_activates_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_activates_default(cValueInstance)
}

func Fn_gtk_entry_get_alignment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_alignment(cValueInstance)
}

func Fn_gtk_entry_get_buffer(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_buffer(cValueInstance)
}

func Fn_gtk_entry_get_completion(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_completion(cValueInstance)
}

func Fn_gtk_entry_get_current_icon_drag_source(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_current_icon_drag_source(cValueInstance)
}

func Fn_gtk_entry_get_cursor_hadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_cursor_hadjustment(cValueInstance)
}

func Fn_gtk_entry_get_has_frame(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_has_frame(cValueInstance)
}

func Fn_gtk_entry_get_icon_activatable(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	C.gtk_entry_get_icon_activatable(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_icon_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_entry_get_icon_at_pos(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_get_icon_gicon(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	C.gtk_entry_get_icon_gicon(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_icon_name(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	C.gtk_entry_get_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_icon_pixbuf(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	C.gtk_entry_get_icon_pixbuf(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_icon_sensitive(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	C.gtk_entry_get_icon_sensitive(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_icon_stock(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	C.gtk_entry_get_icon_stock(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_icon_storage_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	C.gtk_entry_get_icon_storage_type(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_icon_tooltip_markup(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	C.gtk_entry_get_icon_tooltip_markup(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_icon_tooltip_text(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	C.gtk_entry_get_icon_tooltip_text(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_inner_border(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_inner_border(cValueInstance)
}

func Fn_gtk_entry_get_invisible_char(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_invisible_char(cValueInstance)
}

func Fn_gtk_entry_get_layout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_layout(cValueInstance)
}

func Fn_gtk_entry_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_entry_get_layout_offsets(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_get_max_length(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_max_length(cValueInstance)
}

func Fn_gtk_entry_get_overwrite_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_overwrite_mode(cValueInstance)
}

func Fn_gtk_entry_get_progress_fraction(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_progress_fraction(cValueInstance)
}

func Fn_gtk_entry_get_progress_pulse_step(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_progress_pulse_step(cValueInstance)
}

func Fn_gtk_entry_get_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_text(cValueInstance)
}

func Fn_gtk_entry_get_text_length(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_text_length(cValueInstance)
}

func Fn_gtk_entry_get_visibility(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_visibility(cValueInstance)
}

func Fn_gtk_entry_get_width_chars(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_get_width_chars(cValueInstance)
}

func Fn_gtk_entry_layout_index_to_text_index(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_layout_index_to_text_index(cValueInstance, cValue0)
}

func Fn_gtk_entry_progress_pulse(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_progress_pulse(cValueInstance)
}

func Fn_gtk_entry_set_activates_default(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_set_activates_default(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_alignment(paramInstance unsafe.Pointer, param0 float32) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	C.gtk_entry_set_alignment(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_buffer(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkEntryBuffer)(unsafe.Pointer(param0))

	C.gtk_entry_set_buffer(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_completion(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkEntryCompletion)(unsafe.Pointer(param0))

	C.gtk_entry_set_completion(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_cursor_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_entry_set_cursor_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_has_frame(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_set_has_frame(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_icon_activatable(paramInstance unsafe.Pointer, param0 int, param1 bool) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := toCBool(param1)

	C.gtk_entry_set_icon_activatable(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_drag_source(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, param2 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.GtkTargetList)(unsafe.Pointer(param1))

	cValue2 := (C.GdkDragAction)(param2)

	C.gtk_entry_set_icon_drag_source(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_entry_set_icon_from_gicon(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.GIcon)(unsafe.Pointer(param1))

	C.gtk_entry_set_icon_from_gicon(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_from_icon_name(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_set_icon_from_icon_name(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_from_pixbuf(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	C.gtk_entry_set_icon_from_pixbuf(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_from_stock(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_set_icon_from_stock(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_sensitive(paramInstance unsafe.Pointer, param0 int, param1 bool) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := toCBool(param1)

	C.gtk_entry_set_icon_sensitive(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_tooltip_markup(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_set_icon_tooltip_markup(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_icon_tooltip_text(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_set_icon_tooltip_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_set_inner_border(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBorder)(unsafe.Pointer(param0))

	C.gtk_entry_set_inner_border(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_invisible_char(paramInstance unsafe.Pointer, param0 rune) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gunichar)(param0)

	C.gtk_entry_set_invisible_char(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_max_length(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_set_max_length(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_overwrite_mode(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_set_overwrite_mode(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_progress_fraction(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_entry_set_progress_fraction(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_progress_pulse_step(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_entry_set_progress_pulse_step(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_entry_set_text(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_visibility(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_set_visibility(cValueInstance, cValue0)
}

func Fn_gtk_entry_set_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_set_width_chars(cValueInstance, cValue0)
}

func Fn_gtk_entry_text_index_to_layout_index(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_text_index_to_layout_index(cValueInstance, cValue0)
}

func Fn_gtk_entry_unset_invisible_char(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_unset_invisible_char(cValueInstance)
}

func Fn_gtk_entry_buffer_new(param0 string, param1 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_entry_buffer_new(cValue0, cValue1)
}

func Fn_gtk_entry_buffer_delete_text(paramInstance unsafe.Pointer, param0 uint, param1 int) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_entry_buffer_delete_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_buffer_emit_deleted_text(paramInstance unsafe.Pointer, param0 uint, param1 uint) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	C.gtk_entry_buffer_emit_deleted_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_buffer_emit_inserted_text(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 uint) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.guint)(param2)

	C.gtk_entry_buffer_emit_inserted_text(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_entry_buffer_get_bytes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_entry_buffer_get_bytes(cValueInstance)
}

func Fn_gtk_entry_buffer_get_length(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_entry_buffer_get_length(cValueInstance)
}

func Fn_gtk_entry_buffer_get_max_length(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_entry_buffer_get_max_length(cValueInstance)
}

func Fn_gtk_entry_buffer_get_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_entry_buffer_get_text(cValueInstance)
}

func Fn_gtk_entry_buffer_insert_text(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 int) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	C.gtk_entry_buffer_insert_text(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_entry_buffer_set_max_length(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_buffer_set_max_length(cValueInstance, cValue0)
}

func Fn_gtk_entry_buffer_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_entry_buffer_set_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_completion_new() {
	C.gtk_entry_completion_new()
}

func Fn_gtk_entry_completion_complete(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_complete(cValueInstance)
}

func Fn_gtk_entry_completion_delete_action(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_completion_delete_action(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_get_completion_prefix(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_completion_prefix(cValueInstance)
}

func Fn_gtk_entry_completion_get_entry(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_entry(cValueInstance)
}

func Fn_gtk_entry_completion_get_inline_completion(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_inline_completion(cValueInstance)
}

func Fn_gtk_entry_completion_get_inline_selection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_inline_selection(cValueInstance)
}

func Fn_gtk_entry_completion_get_minimum_key_length(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_minimum_key_length(cValueInstance)
}

func Fn_gtk_entry_completion_get_model(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_model(cValueInstance)
}

func Fn_gtk_entry_completion_get_popup_completion(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_popup_completion(cValueInstance)
}

func Fn_gtk_entry_completion_get_popup_set_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_popup_set_width(cValueInstance)
}

func Fn_gtk_entry_completion_get_popup_single_match(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_popup_single_match(cValueInstance)
}

func Fn_gtk_entry_completion_get_text_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_get_text_column(cValueInstance)
}

func Fn_gtk_entry_completion_insert_action_markup(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_completion_insert_action_markup(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_completion_insert_action_text(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_entry_completion_insert_action_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_completion_insert_prefix(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	C.gtk_entry_completion_insert_prefix(cValueInstance)
}

func Fn_gtk_entry_completion_set_inline_completion(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_inline_completion(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_set_inline_selection(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_inline_selection(cValueInstance, cValue0)
}

// UNSUPPORTED : set_match_func : has callback

func Fn_gtk_entry_completion_set_minimum_key_length(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_completion_set_minimum_key_length(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_entry_completion_set_model(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_set_popup_completion(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_popup_completion(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_set_popup_set_width(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_popup_set_width(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_set_popup_single_match(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_entry_completion_set_popup_single_match(cValueInstance, cValue0)
}

func Fn_gtk_entry_completion_set_text_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_entry_completion_set_text_column(cValueInstance, cValue0)
}

func Fn_gtk_event_box_new() {
	C.gtk_event_box_new()
}

func Fn_gtk_event_box_get_above_child(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	C.gtk_event_box_get_above_child(cValueInstance)
}

func Fn_gtk_event_box_get_visible_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	C.gtk_event_box_get_visible_window(cValueInstance)
}

func Fn_gtk_event_box_set_above_child(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_event_box_set_above_child(cValueInstance, cValue0)
}

func Fn_gtk_event_box_set_visible_window(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_event_box_set_visible_window(cValueInstance, cValue0)
}

// UNSUPPORTED : new : blacklisted
// UNSUPPORTED : forward : blacklisted
// UNSUPPORTED : get_group : blacklisted
// UNSUPPORTED : set_im_context : blacklisted
func Fn_gtk_expander_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_expander_new(cValue0)
}

func Fn_gtk_expander_new_with_mnemonic(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_expander_new_with_mnemonic(cValue0)
}

func Fn_gtk_expander_get_expanded(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	C.gtk_expander_get_expanded(cValueInstance)
}

func Fn_gtk_expander_get_label(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	C.gtk_expander_get_label(cValueInstance)
}

func Fn_gtk_expander_get_label_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	C.gtk_expander_get_label_widget(cValueInstance)
}

func Fn_gtk_expander_get_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	C.gtk_expander_get_spacing(cValueInstance)
}

func Fn_gtk_expander_get_use_markup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	C.gtk_expander_get_use_markup(cValueInstance)
}

func Fn_gtk_expander_get_use_underline(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	C.gtk_expander_get_use_underline(cValueInstance)
}

func Fn_gtk_expander_set_expanded(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_expander_set_expanded(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_expander_set_label(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_expander_set_label_widget(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_expander_set_spacing(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_use_markup(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_expander_set_use_markup(cValueInstance, cValue0)
}

func Fn_gtk_expander_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_expander_set_use_underline(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_button_new(param0 string, param1 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkFileChooserAction)(param1)

	C.gtk_file_chooser_button_new(cValue0, cValue1)
}

func Fn_gtk_file_chooser_button_new_with_dialog(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_file_chooser_button_new_with_dialog(cValue0)
}

func Fn_gtk_file_chooser_button_get_focus_on_click(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	C.gtk_file_chooser_button_get_focus_on_click(cValueInstance)
}

func Fn_gtk_file_chooser_button_get_title(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	C.gtk_file_chooser_button_get_title(cValueInstance)
}

func Fn_gtk_file_chooser_button_get_width_chars(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	C.gtk_file_chooser_button_get_width_chars(cValueInstance)
}

func Fn_gtk_file_chooser_button_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_button_set_focus_on_click(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_button_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_button_set_title(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_button_set_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_file_chooser_button_set_width_chars(cValueInstance, cValue0)
}

// UNSUPPORTED : new : has varargs

func Fn_gtk_file_chooser_widget_new(param0 int) {
	cValue0 := (C.GtkFileChooserAction)(param0)

	C.gtk_file_chooser_widget_new(cValue0)
}

func Fn_gtk_file_filter_new() {
	C.gtk_file_filter_new()
}

// UNSUPPORTED : add_custom : has callback

func Fn_gtk_file_filter_add_mime_type(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_filter_add_mime_type(cValueInstance, cValue0)
}

func Fn_gtk_file_filter_add_pattern(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_filter_add_pattern(cValueInstance, cValue0)
}

func Fn_gtk_file_filter_add_pixbuf_formats(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	C.gtk_file_filter_add_pixbuf_formats(cValueInstance)
}

func Fn_gtk_file_filter_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFileFilterInfo)(unsafe.Pointer(param0))

	C.gtk_file_filter_filter(cValueInstance, cValue0)
}

func Fn_gtk_file_filter_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	C.gtk_file_filter_get_name(cValueInstance)
}

func Fn_gtk_file_filter_get_needed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	C.gtk_file_filter_get_needed(cValueInstance)
}

func Fn_gtk_file_filter_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_filter_set_name(cValueInstance, cValue0)
}

func Fn_gtk_fixed_new() {
	C.gtk_fixed_new()
}

func Fn_gtk_fixed_move(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkFixed)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_fixed_move(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_fixed_put(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkFixed)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_fixed_put(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : bind_model : has callback

// UNSUPPORTED : selected_foreach : has callback

// UNSUPPORTED : set_filter_func : has callback

// UNSUPPORTED : set_sort_func : has callback

func Fn_gtk_font_button_new() {
	C.gtk_font_button_new()
}

func Fn_gtk_font_button_new_with_font(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_button_new_with_font(cValue0)
}

func Fn_gtk_font_button_get_font_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	C.gtk_font_button_get_font_name(cValueInstance)
}

func Fn_gtk_font_button_get_show_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	C.gtk_font_button_get_show_size(cValueInstance)
}

func Fn_gtk_font_button_get_show_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	C.gtk_font_button_get_show_style(cValueInstance)
}

func Fn_gtk_font_button_get_title(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	C.gtk_font_button_get_title(cValueInstance)
}

func Fn_gtk_font_button_get_use_font(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	C.gtk_font_button_get_use_font(cValueInstance)
}

func Fn_gtk_font_button_get_use_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	C.gtk_font_button_get_use_size(cValueInstance)
}

func Fn_gtk_font_button_set_font_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_button_set_font_name(cValueInstance, cValue0)
}

func Fn_gtk_font_button_set_show_size(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_font_button_set_show_size(cValueInstance, cValue0)
}

func Fn_gtk_font_button_set_show_style(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_font_button_set_show_style(cValueInstance, cValue0)
}

func Fn_gtk_font_button_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_button_set_title(cValueInstance, cValue0)
}

func Fn_gtk_font_button_set_use_font(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_font_button_set_use_font(cValueInstance, cValue0)
}

func Fn_gtk_font_button_set_use_size(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_font_button_set_use_size(cValueInstance, cValue0)
}

func Fn_gtk_font_selection_new() {
	C.gtk_font_selection_new()
}

func Fn_gtk_font_selection_get_face(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_face(cValueInstance)
}

func Fn_gtk_font_selection_get_face_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_face_list(cValueInstance)
}

func Fn_gtk_font_selection_get_family(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_family(cValueInstance)
}

func Fn_gtk_font_selection_get_family_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_family_list(cValueInstance)
}

func Fn_gtk_font_selection_get_font_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_font_name(cValueInstance)
}

func Fn_gtk_font_selection_get_preview_entry(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_preview_entry(cValueInstance)
}

func Fn_gtk_font_selection_get_preview_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_preview_text(cValueInstance)
}

func Fn_gtk_font_selection_get_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_size(cValueInstance)
}

func Fn_gtk_font_selection_get_size_entry(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_size_entry(cValueInstance)
}

func Fn_gtk_font_selection_get_size_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_get_size_list(cValueInstance)
}

func Fn_gtk_font_selection_set_font_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_selection_set_font_name(cValueInstance, cValue0)
}

func Fn_gtk_font_selection_set_preview_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_selection_set_preview_text(cValueInstance, cValue0)
}

func Fn_gtk_font_selection_dialog_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_selection_dialog_new(cValue0)
}

func Fn_gtk_font_selection_dialog_get_cancel_button(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_dialog_get_cancel_button(cValueInstance)
}

func Fn_gtk_font_selection_dialog_get_font_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_dialog_get_font_name(cValueInstance)
}

func Fn_gtk_font_selection_dialog_get_ok_button(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_dialog_get_ok_button(cValueInstance)
}

func Fn_gtk_font_selection_dialog_get_preview_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	C.gtk_font_selection_dialog_get_preview_text(cValueInstance)
}

func Fn_gtk_font_selection_dialog_set_font_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_selection_dialog_set_font_name(cValueInstance, cValue0)
}

func Fn_gtk_font_selection_dialog_set_preview_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_selection_dialog_set_preview_text(cValueInstance, cValue0)
}

func Fn_gtk_frame_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_frame_new(cValue0)
}

func Fn_gtk_frame_get_label(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	C.gtk_frame_get_label(cValueInstance)
}

func Fn_gtk_frame_get_label_align(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gfloat)(unsafe.Pointer(param0))

	cValue1 := (*C.gfloat)(unsafe.Pointer(param1))

	C.gtk_frame_get_label_align(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_frame_get_label_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	C.gtk_frame_get_label_widget(cValueInstance)
}

func Fn_gtk_frame_get_shadow_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	C.gtk_frame_get_shadow_type(cValueInstance)
}

func Fn_gtk_frame_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_frame_set_label(cValueInstance, cValue0)
}

func Fn_gtk_frame_set_label_align(paramInstance unsafe.Pointer, param0 float32, param1 float32) {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	C.gtk_frame_set_label_align(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_frame_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_frame_set_label_widget(cValueInstance, cValue0)
}

func Fn_gtk_frame_set_shadow_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkShadowType)(param0)

	C.gtk_frame_set_shadow_type(cValueInstance, cValue0)
}

func Fn_gtk_gesture_get_last_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventSequence)(unsafe.Pointer(param0))

	C.gtk_gesture_get_last_event(cValueInstance, cValue0)
}

func Fn_gtk_grid_new() {
	C.gtk_grid_new()
}

func Fn_gtk_grid_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	C.gtk_grid_attach(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_grid_attach_next_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.GtkPositionType)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	C.gtk_grid_attach_next_to(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_grid_get_column_homogeneous(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	C.gtk_grid_get_column_homogeneous(cValueInstance)
}

func Fn_gtk_grid_get_column_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	C.gtk_grid_get_column_spacing(cValueInstance)
}

func Fn_gtk_grid_get_row_homogeneous(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	C.gtk_grid_get_row_homogeneous(cValueInstance)
}

func Fn_gtk_grid_get_row_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	C.gtk_grid_get_row_spacing(cValueInstance)
}

func Fn_gtk_grid_set_column_homogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_grid_set_column_homogeneous(cValueInstance, cValue0)
}

func Fn_gtk_grid_set_column_spacing(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_grid_set_column_spacing(cValueInstance, cValue0)
}

func Fn_gtk_grid_set_row_homogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_grid_set_row_homogeneous(cValueInstance, cValue0)
}

func Fn_gtk_grid_set_row_spacing(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_grid_set_row_spacing(cValueInstance, cValue0)
}

func Fn_gtk_hbox_new(param0 bool, param1 int) {
	cValue0 := toCBool(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_hbox_new(cValue0, cValue1)
}

func Fn_gtk_hbutton_box_new() {
	C.gtk_hbutton_box_new()
}

func Fn_gtk_hpaned_new() {
	C.gtk_hpaned_new()
}

func Fn_gtk_hsv_new() {
	C.gtk_hsv_new()
}

func Fn_gtk_hsv_get_color(paramInstance unsafe.Pointer, param0 *float64, param1 *float64, param2 *float64) {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	C.gtk_hsv_get_color(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_hsv_get_metrics(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_hsv_get_metrics(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_hsv_is_adjusting(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	C.gtk_hsv_is_adjusting(cValueInstance)
}

func Fn_gtk_hsv_set_color(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 float64) {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	cValue0 := (C.double)(param0)

	cValue1 := (C.double)(param1)

	cValue2 := (C.double)(param2)

	C.gtk_hsv_set_color(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_hsv_set_metrics(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_hsv_set_metrics(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_hsv_to_rgb(param0 float64, param1 float64, param2 float64, param3 *float64, param4 *float64, param5 *float64) {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (*C.gdouble)(unsafe.Pointer(param3))

	cValue4 := (*C.gdouble)(unsafe.Pointer(param4))

	cValue5 := (*C.gdouble)(unsafe.Pointer(param5))

	C.gtk_hsv_to_rgb(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_hscale_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_hscale_new(cValue0)
}

func Fn_gtk_hscale_new_with_range(param0 float64, param1 float64, param2 float64) {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	C.gtk_hscale_new_with_range(cValue0, cValue1, cValue2)
}

func Fn_gtk_hscrollbar_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_hscrollbar_new(cValue0)
}

func Fn_gtk_hseparator_new() {
	C.gtk_hseparator_new()
}

func Fn_gtk_handle_box_new() {
	C.gtk_handle_box_new()
}

func Fn_gtk_handle_box_get_child_detached(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	C.gtk_handle_box_get_child_detached(cValueInstance)
}

func Fn_gtk_handle_box_get_handle_position(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	C.gtk_handle_box_get_handle_position(cValueInstance)
}

func Fn_gtk_handle_box_get_shadow_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	C.gtk_handle_box_get_shadow_type(cValueInstance)
}

func Fn_gtk_handle_box_get_snap_edge(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	C.gtk_handle_box_get_snap_edge(cValueInstance)
}

func Fn_gtk_handle_box_set_handle_position(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPositionType)(param0)

	C.gtk_handle_box_set_handle_position(cValueInstance, cValue0)
}

func Fn_gtk_handle_box_set_shadow_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkShadowType)(param0)

	C.gtk_handle_box_set_shadow_type(cValueInstance, cValue0)
}

func Fn_gtk_handle_box_set_snap_edge(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPositionType)(param0)

	C.gtk_handle_box_set_snap_edge(cValueInstance, cValue0)
}

func Fn_gtk_im_context_delete_surrounding(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_im_context_delete_surrounding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_im_context_filter_keypress(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	C.gtk_im_context_filter_keypress(cValueInstance, cValue0)
}

func Fn_gtk_im_context_focus_in(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	C.gtk_im_context_focus_in(cValueInstance)
}

func Fn_gtk_im_context_focus_out(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	C.gtk_im_context_focus_out(cValueInstance)
}

func Fn_gtk_im_context_get_preedit_string(paramInstance unsafe.Pointer, param0 *string, param1 *unsafe.Pointer, param2 *int) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (**C.PangoAttrList)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_im_context_get_preedit_string(cValueInstance, cValue0, cValue1, cValue2)

	param0String := C.GoString(cValue0String)
	*param0 = param0String
}

func Fn_gtk_im_context_get_surrounding(paramInstance unsafe.Pointer, param0 *string, param1 *int) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_im_context_get_surrounding(cValueInstance, cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String
}

func Fn_gtk_im_context_reset(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	C.gtk_im_context_reset(cValueInstance)
}

func Fn_gtk_im_context_set_client_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gtk_im_context_set_client_window(cValueInstance, cValue0)
}

func Fn_gtk_im_context_set_cursor_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_im_context_set_cursor_location(cValueInstance, cValue0)
}

func Fn_gtk_im_context_set_surrounding(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_im_context_set_surrounding(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_im_context_set_use_preedit(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_im_context_set_use_preedit(cValueInstance, cValue0)
}

func Fn_gtk_im_context_simple_new() {
	C.gtk_im_context_simple_new()
}

func Fn_gtk_im_context_simple_add_compose_file(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIMContextSimple)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_im_context_simple_add_compose_file(cValueInstance, cValue0)
}

func Fn_gtk_im_context_simple_add_table(paramInstance unsafe.Pointer, param0 []uint16, param1 int, param2 int) {
	// has non-string array param
}

func Fn_gtk_im_multicontext_new() {
	C.gtk_im_multicontext_new()
}

func Fn_gtk_im_multicontext_append_menuitems(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIMMulticontext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkMenuShell)(unsafe.Pointer(param0))

	C.gtk_im_multicontext_append_menuitems(cValueInstance, cValue0)
}

func Fn_gtk_im_multicontext_get_context_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIMMulticontext)(unsafe.Pointer(paramInstance))

	C.gtk_im_multicontext_get_context_id(cValueInstance)
}

func Fn_gtk_im_multicontext_set_context_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIMMulticontext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_im_multicontext_set_context_id(cValueInstance, cValue0)
}

func Fn_gtk_icon_factory_new() {
	C.gtk_icon_factory_new()
}

func Fn_gtk_icon_factory_add(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconFactory)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkIconSet)(unsafe.Pointer(param1))

	C.gtk_icon_factory_add(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_factory_add_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconFactory)(unsafe.Pointer(paramInstance))

	C.gtk_icon_factory_add_default(cValueInstance)
}

func Fn_gtk_icon_factory_lookup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconFactory)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_factory_lookup(cValueInstance, cValue0)
}

func Fn_gtk_icon_factory_remove_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconFactory)(unsafe.Pointer(paramInstance))

	C.gtk_icon_factory_remove_default(cValueInstance)
}

func Fn_gtk_icon_factory_lookup_default(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_factory_lookup_default(cValue0)
}

func Fn_gtk_icon_info_new_for_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GtkIconTheme)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	C.gtk_icon_info_new_for_pixbuf(cValue0, cValue1)
}

func Fn_gtk_icon_info_copy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	C.gtk_icon_info_copy(cValueInstance)
}

func Fn_gtk_icon_info_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	C.gtk_icon_info_free(cValueInstance)
}

func Fn_gtk_icon_info_get_attach_points(paramInstance unsafe.Pointer, param0 []unsafe.Pointer, param1 *int) {
	// has non-string array param
}

func Fn_gtk_icon_info_get_base_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	C.gtk_icon_info_get_base_size(cValueInstance)
}

func Fn_gtk_icon_info_get_builtin_pixbuf(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	C.gtk_icon_info_get_builtin_pixbuf(cValueInstance)
}

func Fn_gtk_icon_info_get_display_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	C.gtk_icon_info_get_display_name(cValueInstance)
}

func Fn_gtk_icon_info_get_embedded_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_icon_info_get_embedded_rect(cValueInstance, cValue0)
}

func Fn_gtk_icon_info_get_filename(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	C.gtk_icon_info_get_filename(cValueInstance)
}

func Fn_gtk_icon_info_load_icon(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	C.gtk_icon_info_load_icon(cValueInstance, cError)
}

// UNSUPPORTED : load_icon_async : has callback

// UNSUPPORTED : load_symbolic_async : has callback

// UNSUPPORTED : load_symbolic_for_context_async : has callback

func Fn_gtk_icon_info_set_raw_coordinates(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_icon_info_set_raw_coordinates(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_new() {
	C.gtk_icon_theme_new()
}

func Fn_gtk_icon_theme_append_search_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_append_search_path(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_choose_icon(paramInstance unsafe.Pointer, param0 []string, param1 int, param2 int) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	C.gtk_icon_theme_choose_icon(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_theme_get_example_icon_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	C.gtk_icon_theme_get_example_icon_name(cValueInstance)
}

func Fn_gtk_icon_theme_get_icon_sizes(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_get_icon_sizes(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_get_search_path(paramInstance unsafe.Pointer, param0 *[]string, param1 *int) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	var cValue0ArrayPointer **C.gchar
	cValue0 := &cValue0ArrayPointer

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_icon_theme_get_search_path(cValueInstance, cValue0, cValue1)

	param0OutLen := int(*cValue1)
	param0Out := make([]string, param0OutLen, param0OutLen)
	if param0OutLen > 0 {
		param0OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0ArrayPointer))[:param0OutLen:param0OutLen]
		for param0Outi, param0OutCString := range param0OutCSlice {
			param0Out[param0Outi] = C.GoString(param0OutCString)
		}
	}
	*param0 = param0Out
}

func Fn_gtk_icon_theme_has_icon(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_has_icon(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_list_contexts(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	C.gtk_icon_theme_list_contexts(cValueInstance)
}

func Fn_gtk_icon_theme_list_icons(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_list_icons(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_load_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, error unsafe.Pointer) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	cError := (**C.GError)(error)

	C.gtk_icon_theme_load_icon(cValueInstance, cValue0, cValue1, cValue2, cError)
}

func Fn_gtk_icon_theme_lookup_by_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	C.gtk_icon_theme_lookup_by_gicon(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_theme_lookup_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	C.gtk_icon_theme_lookup_icon(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_theme_prepend_search_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_prepend_search_path(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_rescan_if_needed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	C.gtk_icon_theme_rescan_if_needed(cValueInstance)
}

func Fn_gtk_icon_theme_set_custom_theme(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_set_custom_theme(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_icon_theme_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_set_search_path(paramInstance unsafe.Pointer, param0 []string, param1 int) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	cValue1 := (C.gint)(param1)

	C.gtk_icon_theme_set_search_path(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_theme_add_builtin_icon(param0 string, param1 int, param2 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.GdkPixbuf)(unsafe.Pointer(param2))

	C.gtk_icon_theme_add_builtin_icon(cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_theme_get_default() {
	C.gtk_icon_theme_get_default()
}

func Fn_gtk_icon_theme_get_for_screen(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_icon_theme_get_for_screen(cValue0)
}

func Fn_gtk_icon_view_new() {
	C.gtk_icon_view_new()
}

func Fn_gtk_icon_view_new_with_model(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_icon_view_new_with_model(cValue0)
}

func Fn_gtk_icon_view_convert_widget_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_icon_view_convert_widget_to_bin_window_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_icon_view_create_drag_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_icon_view_create_drag_icon(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_enable_model_drag_dest(paramInstance unsafe.Pointer, param0 []TargetEntry, param1 int, param2 int) {
	// has non-string array param
}

func Fn_gtk_icon_view_enable_model_drag_source(paramInstance unsafe.Pointer, param0 int, param1 []TargetEntry, param2 int, param3 int) {
	// has non-string array param
}

func Fn_gtk_icon_view_get_column_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_column_spacing(cValueInstance)
}

func Fn_gtk_icon_view_get_columns(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_columns(cValueInstance)
}

func Fn_gtk_icon_view_get_cursor(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkCellRenderer)(unsafe.Pointer(param1))

	C.gtk_icon_view_get_cursor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_view_get_dest_item_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkIconViewDropPosition)(unsafe.Pointer(param3))

	C.gtk_icon_view_get_dest_item_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_icon_view_get_drag_dest_item(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkIconViewDropPosition)(unsafe.Pointer(param1))

	C.gtk_icon_view_get_drag_dest_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_view_get_item_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (**C.GtkCellRenderer)(unsafe.Pointer(param3))

	C.gtk_icon_view_get_item_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_icon_view_get_item_orientation(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_item_orientation(cValueInstance)
}

func Fn_gtk_icon_view_get_item_padding(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_item_padding(cValueInstance)
}

func Fn_gtk_icon_view_get_item_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_item_width(cValueInstance)
}

func Fn_gtk_icon_view_get_margin(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_margin(cValueInstance)
}

func Fn_gtk_icon_view_get_markup_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_markup_column(cValueInstance)
}

func Fn_gtk_icon_view_get_model(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_model(cValueInstance)
}

func Fn_gtk_icon_view_get_path_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_icon_view_get_path_at_pos(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_view_get_pixbuf_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_pixbuf_column(cValueInstance)
}

func Fn_gtk_icon_view_get_reorderable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_reorderable(cValueInstance)
}

func Fn_gtk_icon_view_get_row_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_row_spacing(cValueInstance)
}

func Fn_gtk_icon_view_get_selected_items(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_selected_items(cValueInstance)
}

func Fn_gtk_icon_view_get_selection_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_selection_mode(cValueInstance)
}

func Fn_gtk_icon_view_get_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_spacing(cValueInstance)
}

func Fn_gtk_icon_view_get_text_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_text_column(cValueInstance)
}

func Fn_gtk_icon_view_get_tooltip_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_get_tooltip_column(cValueInstance)
}

func Fn_gtk_icon_view_get_tooltip_context(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := (**C.GtkTreeModel)(unsafe.Pointer(param3))

	cValue4 := (**C.GtkTreePath)(unsafe.Pointer(param4))

	cValue5 := (*C.GtkTreeIter)(unsafe.Pointer(param5))

	C.gtk_icon_view_get_tooltip_context(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_icon_view_get_visible_range(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_icon_view_get_visible_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_view_item_activated(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_icon_view_item_activated(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_path_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_icon_view_path_is_selected(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_scroll_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 float32, param3 float32) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := (C.gfloat)(param2)

	cValue3 := (C.gfloat)(param3)

	C.gtk_icon_view_scroll_to_path(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_icon_view_select_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_select_all(cValueInstance)
}

func Fn_gtk_icon_view_select_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_icon_view_select_path(cValueInstance, cValue0)
}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_icon_view_set_column_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_column_spacing(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_columns(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_columns(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_icon_view_set_cursor(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_view_set_drag_dest_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconViewDropPosition)(param1)

	C.gtk_icon_view_set_drag_dest_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_view_set_item_orientation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkOrientation)(param0)

	C.gtk_icon_view_set_item_orientation(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_item_padding(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_item_padding(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_item_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_item_width(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_margin(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_markup_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_markup_column(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_icon_view_set_model(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_pixbuf_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_pixbuf_column(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_reorderable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_icon_view_set_reorderable(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_row_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_row_spacing(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_selection_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSelectionMode)(param0)

	C.gtk_icon_view_set_selection_mode(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_spacing(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_text_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_text_column(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_tooltip_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTooltip)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkCellRenderer)(unsafe.Pointer(param2))

	C.gtk_icon_view_set_tooltip_cell(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_icon_view_set_tooltip_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_icon_view_set_tooltip_column(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_set_tooltip_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTooltip)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_icon_view_set_tooltip_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_view_unselect_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_unselect_all(cValueInstance)
}

func Fn_gtk_icon_view_unselect_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_icon_view_unselect_path(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_unset_model_drag_dest(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_unset_model_drag_dest(cValueInstance)
}

func Fn_gtk_icon_view_unset_model_drag_source(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	C.gtk_icon_view_unset_model_drag_source(cValueInstance)
}

func Fn_gtk_image_new() {
	C.gtk_image_new()
}

func Fn_gtk_image_new_from_animation(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkPixbufAnimation)(unsafe.Pointer(param0))

	C.gtk_image_new_from_animation(cValue0)
}

func Fn_gtk_image_new_from_file(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_image_new_from_file(cValue0)
}

func Fn_gtk_image_new_from_gicon(param0 unsafe.Pointer, param1 int) {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_new_from_gicon(cValue0, cValue1)
}

func Fn_gtk_image_new_from_icon_name(param0 string, param1 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_new_from_icon_name(cValue0, cValue1)
}

func Fn_gtk_image_new_from_icon_set(param0 unsafe.Pointer, param1 int) {
	cValue0 := (*C.GtkIconSet)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_new_from_icon_set(cValue0, cValue1)
}

func Fn_gtk_image_new_from_pixbuf(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_image_new_from_pixbuf(cValue0)
}

func Fn_gtk_image_new_from_stock(param0 string, param1 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_new_from_stock(cValue0, cValue1)
}

func Fn_gtk_image_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	C.gtk_image_clear(cValueInstance)
}

func Fn_gtk_image_get_animation(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	C.gtk_image_get_animation(cValueInstance)
}

func Fn_gtk_image_get_gicon(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkIconSize)(unsafe.Pointer(param1))

	C.gtk_image_get_gicon(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_image_get_icon_name(paramInstance unsafe.Pointer, param0 *string, param1 *int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (*C.GtkIconSize)(unsafe.Pointer(param1))

	C.gtk_image_get_icon_name(cValueInstance, cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String
}

func Fn_gtk_image_get_icon_set(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkIconSet)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkIconSize)(unsafe.Pointer(param1))

	C.gtk_image_get_icon_set(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_image_get_pixbuf(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	C.gtk_image_get_pixbuf(cValueInstance)
}

func Fn_gtk_image_get_pixel_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	C.gtk_image_get_pixel_size(cValueInstance)
}

func Fn_gtk_image_get_stock(paramInstance unsafe.Pointer, param0 *string, param1 *int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (*C.GtkIconSize)(unsafe.Pointer(param1))

	C.gtk_image_get_stock(cValueInstance, cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String
}

func Fn_gtk_image_get_storage_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	C.gtk_image_get_storage_type(cValueInstance)
}

func Fn_gtk_image_set_from_animation(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbufAnimation)(unsafe.Pointer(param0))

	C.gtk_image_set_from_animation(cValueInstance, cValue0)
}

func Fn_gtk_image_set_from_file(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_image_set_from_file(cValueInstance, cValue0)
}

func Fn_gtk_image_set_from_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_set_from_gicon(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_image_set_from_icon_name(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_set_from_icon_name(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_image_set_from_icon_set(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkIconSet)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_set_from_icon_set(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_image_set_from_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_image_set_from_pixbuf(cValueInstance, cValue0)
}

func Fn_gtk_image_set_from_resource(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_image_set_from_resource(cValueInstance, cValue0)
}

func Fn_gtk_image_set_from_stock(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_image_set_from_stock(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_image_set_pixel_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_image_set_pixel_size(cValueInstance, cValue0)
}

func Fn_gtk_image_menu_item_new() {
	C.gtk_image_menu_item_new()
}

func Fn_gtk_image_menu_item_new_from_stock(param0 string, param1 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkAccelGroup)(unsafe.Pointer(param1))

	C.gtk_image_menu_item_new_from_stock(cValue0, cValue1)
}

func Fn_gtk_image_menu_item_new_with_label(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_image_menu_item_new_with_label(cValue0)
}

func Fn_gtk_image_menu_item_new_with_mnemonic(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_image_menu_item_new_with_mnemonic(cValue0)
}

func Fn_gtk_image_menu_item_get_always_show_image(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_image_menu_item_get_always_show_image(cValueInstance)
}

func Fn_gtk_image_menu_item_get_image(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_image_menu_item_get_image(cValueInstance)
}

func Fn_gtk_image_menu_item_get_use_stock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_image_menu_item_get_use_stock(cValueInstance)
}

func Fn_gtk_image_menu_item_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	C.gtk_image_menu_item_set_accel_group(cValueInstance, cValue0)
}

func Fn_gtk_image_menu_item_set_always_show_image(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_image_menu_item_set_always_show_image(cValueInstance, cValue0)
}

func Fn_gtk_image_menu_item_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_image_menu_item_set_image(cValueInstance, cValue0)
}

func Fn_gtk_image_menu_item_set_use_stock(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_image_menu_item_set_use_stock(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_new() {
	C.gtk_info_bar_new()
}

// UNSUPPORTED : new_with_buttons : has varargs

func Fn_gtk_info_bar_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_info_bar_add_action_widget(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_info_bar_add_button(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_info_bar_add_button(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : add_buttons : has varargs

func Fn_gtk_info_bar_get_action_area(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	C.gtk_info_bar_get_action_area(cValueInstance)
}

func Fn_gtk_info_bar_get_content_area(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	C.gtk_info_bar_get_content_area(cValueInstance)
}

func Fn_gtk_info_bar_get_message_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	C.gtk_info_bar_get_message_type(cValueInstance)
}

func Fn_gtk_info_bar_response(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_info_bar_response(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_set_default_response(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_info_bar_set_default_response(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_set_message_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkMessageType)(param0)

	C.gtk_info_bar_set_message_type(cValueInstance, cValue0)
}

func Fn_gtk_info_bar_set_response_sensitive(paramInstance unsafe.Pointer, param0 int, param1 bool) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := toCBool(param1)

	C.gtk_info_bar_set_response_sensitive(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_invisible_new() {
	C.gtk_invisible_new()
}

func Fn_gtk_invisible_new_for_screen(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_invisible_new_for_screen(cValue0)
}

func Fn_gtk_invisible_get_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkInvisible)(unsafe.Pointer(paramInstance))

	C.gtk_invisible_get_screen(cValueInstance)
}

func Fn_gtk_invisible_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkInvisible)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_invisible_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_label_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_label_new(cValue0)
}

func Fn_gtk_label_new_with_mnemonic(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_label_new_with_mnemonic(cValue0)
}

func Fn_gtk_label_get_angle(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_angle(cValueInstance)
}

func Fn_gtk_label_get_attributes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_attributes(cValueInstance)
}

func Fn_gtk_label_get_current_uri(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_current_uri(cValueInstance)
}

func Fn_gtk_label_get_ellipsize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_ellipsize(cValueInstance)
}

func Fn_gtk_label_get_justify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_justify(cValueInstance)
}

func Fn_gtk_label_get_label(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_label(cValueInstance)
}

func Fn_gtk_label_get_layout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_layout(cValueInstance)
}

func Fn_gtk_label_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_label_get_layout_offsets(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_label_get_line_wrap(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_line_wrap(cValueInstance)
}

func Fn_gtk_label_get_line_wrap_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_line_wrap_mode(cValueInstance)
}

func Fn_gtk_label_get_max_width_chars(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_max_width_chars(cValueInstance)
}

func Fn_gtk_label_get_mnemonic_keyval(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_mnemonic_keyval(cValueInstance)
}

func Fn_gtk_label_get_mnemonic_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_mnemonic_widget(cValueInstance)
}

func Fn_gtk_label_get_selectable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_selectable(cValueInstance)
}

func Fn_gtk_label_get_selection_bounds(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_label_get_selection_bounds(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_label_get_single_line_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_single_line_mode(cValueInstance)
}

func Fn_gtk_label_get_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_text(cValueInstance)
}

func Fn_gtk_label_get_track_visited_links(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_track_visited_links(cValueInstance)
}

func Fn_gtk_label_get_use_markup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_use_markup(cValueInstance)
}

func Fn_gtk_label_get_use_underline(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_use_underline(cValueInstance)
}

func Fn_gtk_label_get_width_chars(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	C.gtk_label_get_width_chars(cValueInstance)
}

func Fn_gtk_label_select_region(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_label_select_region(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_label_set_angle(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_label_set_angle(cValueInstance, cValue0)
}

func Fn_gtk_label_set_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoAttrList)(unsafe.Pointer(param0))

	C.gtk_label_set_attributes(cValueInstance, cValue0)
}

func Fn_gtk_label_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoEllipsizeMode)(param0)

	C.gtk_label_set_ellipsize(cValueInstance, cValue0)
}

func Fn_gtk_label_set_justify(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkJustification)(param0)

	C.gtk_label_set_justify(cValueInstance, cValue0)
}

func Fn_gtk_label_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_label_set_label(cValueInstance, cValue0)
}

func Fn_gtk_label_set_line_wrap(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_label_set_line_wrap(cValueInstance, cValue0)
}

func Fn_gtk_label_set_line_wrap_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoWrapMode)(param0)

	C.gtk_label_set_line_wrap_mode(cValueInstance, cValue0)
}

func Fn_gtk_label_set_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_label_set_markup(cValueInstance, cValue0)
}

func Fn_gtk_label_set_markup_with_mnemonic(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_label_set_markup_with_mnemonic(cValueInstance, cValue0)
}

func Fn_gtk_label_set_max_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_label_set_max_width_chars(cValueInstance, cValue0)
}

func Fn_gtk_label_set_mnemonic_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_label_set_mnemonic_widget(cValueInstance, cValue0)
}

func Fn_gtk_label_set_pattern(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_label_set_pattern(cValueInstance, cValue0)
}

func Fn_gtk_label_set_selectable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_label_set_selectable(cValueInstance, cValue0)
}

func Fn_gtk_label_set_single_line_mode(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_label_set_single_line_mode(cValueInstance, cValue0)
}

func Fn_gtk_label_set_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_label_set_text(cValueInstance, cValue0)
}

func Fn_gtk_label_set_text_with_mnemonic(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_label_set_text_with_mnemonic(cValueInstance, cValue0)
}

func Fn_gtk_label_set_track_visited_links(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_label_set_track_visited_links(cValueInstance, cValue0)
}

func Fn_gtk_label_set_use_markup(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_label_set_use_markup(cValueInstance, cValue0)
}

func Fn_gtk_label_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_label_set_use_underline(cValueInstance, cValue0)
}

func Fn_gtk_label_set_width_chars(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_label_set_width_chars(cValueInstance, cValue0)
}

func Fn_gtk_layout_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	C.gtk_layout_new(cValue0, cValue1)
}

func Fn_gtk_layout_get_bin_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	C.gtk_layout_get_bin_window(cValueInstance)
}

func Fn_gtk_layout_get_hadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	C.gtk_layout_get_hadjustment(cValueInstance)
}

func Fn_gtk_layout_get_size(paramInstance unsafe.Pointer, param0 *uint, param1 *uint) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	C.gtk_layout_get_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_layout_get_vadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	C.gtk_layout_get_vadjustment(cValueInstance)
}

func Fn_gtk_layout_move(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_layout_move(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_layout_put(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_layout_put(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_layout_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_layout_set_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_layout_set_size(paramInstance unsafe.Pointer, param0 uint, param1 uint) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	C.gtk_layout_set_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_layout_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_layout_set_vadjustment(cValueInstance, cValue0)
}

func Fn_gtk_link_button_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_link_button_new(cValue0)
}

func Fn_gtk_link_button_new_with_label(param0 string, param1 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_link_button_new_with_label(cValue0, cValue1)
}

func Fn_gtk_link_button_get_uri(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	C.gtk_link_button_get_uri(cValueInstance)
}

func Fn_gtk_link_button_get_visited(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	C.gtk_link_button_get_visited(cValueInstance)
}

func Fn_gtk_link_button_set_uri(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_link_button_set_uri(cValueInstance, cValue0)
}

func Fn_gtk_link_button_set_visited(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_link_button_set_visited(cValueInstance, cValue0)
}

// UNSUPPORTED : bind_model : has callback

// UNSUPPORTED : selected_foreach : has callback

// UNSUPPORTED : set_filter_func : has callback

// UNSUPPORTED : set_header_func : has callback

// UNSUPPORTED : set_sort_func : has callback

// UNSUPPORTED : new : has varargs

func Fn_gtk_list_store_newv(param0 int, param1 []uint64) {
	// has non-string array param
}

func Fn_gtk_list_store_append(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_list_store_append(cValueInstance, cValue0)
}

func Fn_gtk_list_store_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	C.gtk_list_store_clear(cValueInstance)
}

func Fn_gtk_list_store_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_list_store_insert(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_list_store_insert_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_list_store_insert_after(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_list_store_insert_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_list_store_insert_before(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : insert_with_values : has varargs

func Fn_gtk_list_store_insert_with_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 []int, param3 []gobject.Value, param4 int) {
	// has non-string array param
}

func Fn_gtk_list_store_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_list_store_iter_is_valid(cValueInstance, cValue0)
}

func Fn_gtk_list_store_move_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_list_store_move_after(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_list_store_move_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_list_store_move_before(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_list_store_prepend(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_list_store_prepend(cValueInstance, cValue0)
}

func Fn_gtk_list_store_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_list_store_remove(cValueInstance, cValue0)
}

func Fn_gtk_list_store_reorder(paramInstance unsafe.Pointer, param0 []int) {
	// has non-string array param
}

// UNSUPPORTED : set : has varargs

func Fn_gtk_list_store_set_column_types(paramInstance unsafe.Pointer, param0 int, param1 []uint64) {
	// has non-string array param
}

// UNSUPPORTED : set_valist : has va_list

func Fn_gtk_list_store_set_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_list_store_set_value(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_list_store_set_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []int, param2 []gobject.Value, param3 int) {
	// has non-string array param
}

func Fn_gtk_list_store_swap(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_list_store_swap(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_menu_new() {
	C.gtk_menu_new()
}

func Fn_gtk_menu_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.guint)(param3)

	cValue4 := (C.guint)(param4)

	C.gtk_menu_attach(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

// UNSUPPORTED : attach_to_widget : has callback

func Fn_gtk_menu_detach(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_detach(cValueInstance)
}

func Fn_gtk_menu_get_accel_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_get_accel_group(cValueInstance)
}

func Fn_gtk_menu_get_accel_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_get_accel_path(cValueInstance)
}

func Fn_gtk_menu_get_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_get_active(cValueInstance)
}

func Fn_gtk_menu_get_attach_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_get_attach_widget(cValueInstance)
}

func Fn_gtk_menu_get_monitor(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_get_monitor(cValueInstance)
}

func Fn_gtk_menu_get_reserve_toggle_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_get_reserve_toggle_size(cValueInstance)
}

func Fn_gtk_menu_get_tearoff_state(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_get_tearoff_state(cValueInstance)
}

func Fn_gtk_menu_get_title(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_get_title(cValueInstance)
}

func Fn_gtk_menu_popdown(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_popdown(cValueInstance)
}

// UNSUPPORTED : popup : has callback

// UNSUPPORTED : popup_for_device : has callback

func Fn_gtk_menu_reorder_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_menu_reorder_child(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_menu_reposition(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_reposition(cValueInstance)
}

func Fn_gtk_menu_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	C.gtk_menu_set_accel_group(cValueInstance, cValue0)
}

func Fn_gtk_menu_set_accel_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_set_accel_path(cValueInstance, cValue0)
}

func Fn_gtk_menu_set_active(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_menu_set_active(cValueInstance, cValue0)
}

func Fn_gtk_menu_set_monitor(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_menu_set_monitor(cValueInstance, cValue0)
}

func Fn_gtk_menu_set_reserve_toggle_size(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_set_reserve_toggle_size(cValueInstance, cValue0)
}

func Fn_gtk_menu_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_menu_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_menu_set_tearoff_state(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_set_tearoff_state(cValueInstance, cValue0)
}

func Fn_gtk_menu_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_set_title(cValueInstance, cValue0)
}

func Fn_gtk_menu_get_for_attach_widget(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_get_for_attach_widget(cValue0)
}

func Fn_gtk_menu_bar_new() {
	C.gtk_menu_bar_new()
}

func Fn_gtk_menu_bar_get_child_pack_direction(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	C.gtk_menu_bar_get_child_pack_direction(cValueInstance)
}

func Fn_gtk_menu_bar_get_pack_direction(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	C.gtk_menu_bar_get_pack_direction(cValueInstance)
}

func Fn_gtk_menu_bar_set_child_pack_direction(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPackDirection)(param0)

	C.gtk_menu_bar_set_child_pack_direction(cValueInstance, cValue0)
}

func Fn_gtk_menu_bar_set_pack_direction(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPackDirection)(param0)

	C.gtk_menu_bar_set_pack_direction(cValueInstance, cValue0)
}

func Fn_gtk_menu_item_new() {
	C.gtk_menu_item_new()
}

func Fn_gtk_menu_item_new_with_label(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_item_new_with_label(cValue0)
}

func Fn_gtk_menu_item_new_with_mnemonic(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_item_new_with_mnemonic(cValue0)
}

func Fn_gtk_menu_item_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_activate(cValueInstance)
}

func Fn_gtk_menu_item_deselect(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_deselect(cValueInstance)
}

func Fn_gtk_menu_item_get_accel_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_get_accel_path(cValueInstance)
}

func Fn_gtk_menu_item_get_label(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_get_label(cValueInstance)
}

func Fn_gtk_menu_item_get_right_justified(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_get_right_justified(cValueInstance)
}

func Fn_gtk_menu_item_get_submenu(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_get_submenu(cValueInstance)
}

func Fn_gtk_menu_item_get_use_underline(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_get_use_underline(cValueInstance)
}

func Fn_gtk_menu_item_select(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_select(cValueInstance)
}

func Fn_gtk_menu_item_set_accel_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_item_set_accel_path(cValueInstance, cValue0)
}

func Fn_gtk_menu_item_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_item_set_label(cValueInstance, cValue0)
}

func Fn_gtk_menu_item_set_right_justified(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_item_set_right_justified(cValueInstance, cValue0)
}

func Fn_gtk_menu_item_set_submenu(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_item_set_submenu(cValueInstance, cValue0)
}

func Fn_gtk_menu_item_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_item_set_use_underline(cValueInstance, cValue0)
}

func Fn_gtk_menu_item_toggle_size_allocate(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_menu_item_toggle_size_allocate(cValueInstance, cValue0)
}

func Fn_gtk_menu_item_toggle_size_request(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	C.gtk_menu_item_toggle_size_request(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_activate_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_menu_shell_activate_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_menu_shell_append(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_shell_append(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	C.gtk_menu_shell_cancel(cValueInstance)
}

func Fn_gtk_menu_shell_deactivate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	C.gtk_menu_shell_deactivate(cValueInstance)
}

func Fn_gtk_menu_shell_deselect(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	C.gtk_menu_shell_deselect(cValueInstance)
}

func Fn_gtk_menu_shell_get_take_focus(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	C.gtk_menu_shell_get_take_focus(cValueInstance)
}

func Fn_gtk_menu_shell_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_menu_shell_insert(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_menu_shell_prepend(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_shell_prepend(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_select_first(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_shell_select_first(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_select_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_shell_select_item(cValueInstance, cValue0)
}

func Fn_gtk_menu_shell_set_take_focus(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_shell_set_take_focus(cValueInstance, cValue0)
}

func Fn_gtk_menu_tool_button_new(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_menu_tool_button_new(cValue0, cValue1)
}

func Fn_gtk_menu_tool_button_new_from_stock(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_tool_button_new_from_stock(cValue0)
}

func Fn_gtk_menu_tool_button_get_menu(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuToolButton)(unsafe.Pointer(paramInstance))

	C.gtk_menu_tool_button_get_menu(cValueInstance)
}

func Fn_gtk_menu_tool_button_set_arrow_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMenuToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_tool_button_set_arrow_tooltip_markup(cValueInstance, cValue0)
}

func Fn_gtk_menu_tool_button_set_arrow_tooltip_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMenuToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_menu_tool_button_set_arrow_tooltip_text(cValueInstance, cValue0)
}

func Fn_gtk_menu_tool_button_set_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_menu_tool_button_set_menu(cValueInstance, cValue0)
}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_with_markup : has varargs

// UNSUPPORTED : format_secondary_markup : has varargs

// UNSUPPORTED : format_secondary_text : has varargs

func Fn_gtk_message_dialog_get_image(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	C.gtk_message_dialog_get_image(cValueInstance)
}

func Fn_gtk_message_dialog_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_message_dialog_set_image(cValueInstance, cValue0)
}

func Fn_gtk_message_dialog_set_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_message_dialog_set_markup(cValueInstance, cValue0)
}

func Fn_gtk_misc_get_alignment(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {
	cValueInstance := (*C.GtkMisc)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gfloat)(unsafe.Pointer(param0))

	cValue1 := (*C.gfloat)(unsafe.Pointer(param1))

	C.gtk_misc_get_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_misc_get_padding(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkMisc)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_misc_get_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_misc_set_alignment(paramInstance unsafe.Pointer, param0 float32, param1 float32) {
	cValueInstance := (*C.GtkMisc)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	C.gtk_misc_set_alignment(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_misc_set_padding(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkMisc)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_misc_set_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_mount_operation_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_mount_operation_new(cValue0)
}

func Fn_gtk_mount_operation_get_parent(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	C.gtk_mount_operation_get_parent(cValueInstance)
}

func Fn_gtk_mount_operation_get_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	C.gtk_mount_operation_get_screen(cValueInstance)
}

func Fn_gtk_mount_operation_is_showing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	C.gtk_mount_operation_is_showing(cValueInstance)
}

func Fn_gtk_mount_operation_set_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_mount_operation_set_parent(cValueInstance, cValue0)
}

func Fn_gtk_mount_operation_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_mount_operation_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_notebook_new() {
	C.gtk_notebook_new()
}

func Fn_gtk_notebook_append_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	C.gtk_notebook_append_page(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_append_page_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkWidget)(unsafe.Pointer(param2))

	C.gtk_notebook_append_page_menu(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_notebook_get_current_page(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_get_current_page(cValueInstance)
}

func Fn_gtk_notebook_get_menu_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_notebook_get_menu_label(cValueInstance, cValue0)
}

func Fn_gtk_notebook_get_menu_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_notebook_get_menu_label_text(cValueInstance, cValue0)
}

func Fn_gtk_notebook_get_n_pages(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_get_n_pages(cValueInstance)
}

func Fn_gtk_notebook_get_nth_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_notebook_get_nth_page(cValueInstance, cValue0)
}

func Fn_gtk_notebook_get_scrollable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_get_scrollable(cValueInstance)
}

func Fn_gtk_notebook_get_show_border(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_get_show_border(cValueInstance)
}

func Fn_gtk_notebook_get_show_tabs(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_get_show_tabs(cValueInstance)
}

func Fn_gtk_notebook_get_tab_detachable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_notebook_get_tab_detachable(cValueInstance, cValue0)
}

func Fn_gtk_notebook_get_tab_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_notebook_get_tab_label(cValueInstance, cValue0)
}

func Fn_gtk_notebook_get_tab_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_notebook_get_tab_label_text(cValueInstance, cValue0)
}

func Fn_gtk_notebook_get_tab_pos(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_get_tab_pos(cValueInstance)
}

func Fn_gtk_notebook_get_tab_reorderable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_notebook_get_tab_reorderable(cValueInstance, cValue0)
}

func Fn_gtk_notebook_insert_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	C.gtk_notebook_insert_page(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_notebook_insert_page_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkWidget)(unsafe.Pointer(param2))

	cValue3 := (C.gint)(param3)

	C.gtk_notebook_insert_page_menu(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_notebook_next_page(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_next_page(cValueInstance)
}

func Fn_gtk_notebook_page_num(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_notebook_page_num(cValueInstance, cValue0)
}

func Fn_gtk_notebook_popup_disable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_popup_disable(cValueInstance)
}

func Fn_gtk_notebook_popup_enable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_popup_enable(cValueInstance)
}

func Fn_gtk_notebook_prepend_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	C.gtk_notebook_prepend_page(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_prepend_page_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkWidget)(unsafe.Pointer(param2))

	C.gtk_notebook_prepend_page_menu(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_notebook_prev_page(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_prev_page(cValueInstance)
}

func Fn_gtk_notebook_remove_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_notebook_remove_page(cValueInstance, cValue0)
}

func Fn_gtk_notebook_reorder_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_notebook_reorder_child(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_set_current_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_notebook_set_current_page(cValueInstance, cValue0)
}

func Fn_gtk_notebook_set_menu_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	C.gtk_notebook_set_menu_label(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_set_menu_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_notebook_set_menu_label_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_set_scrollable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_notebook_set_scrollable(cValueInstance, cValue0)
}

func Fn_gtk_notebook_set_show_border(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_notebook_set_show_border(cValueInstance, cValue0)
}

func Fn_gtk_notebook_set_show_tabs(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_notebook_set_show_tabs(cValueInstance, cValue0)
}

func Fn_gtk_notebook_set_tab_detachable(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_notebook_set_tab_detachable(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_set_tab_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	C.gtk_notebook_set_tab_label(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_set_tab_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_notebook_set_tab_label_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_set_tab_pos(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPositionType)(param0)

	C.gtk_notebook_set_tab_pos(cValueInstance, cValue0)
}

func Fn_gtk_notebook_set_tab_reorderable(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_notebook_set_tab_reorderable(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_page_accessible_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GtkNotebookAccessible)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	C.gtk_notebook_page_accessible_new(cValue0, cValue1)
}

func Fn_gtk_notebook_page_accessible_invalidate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebookPageAccessible)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_page_accessible_invalidate(cValueInstance)
}

func Fn_gtk_page_setup_new() {
	C.gtk_page_setup_new()
}

func Fn_gtk_page_setup_new_from_file(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_page_setup_new_from_file(cValue0, cError)
}

func Fn_gtk_page_setup_new_from_key_file(param0 unsafe.Pointer, param1 string, error unsafe.Pointer) {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	C.gtk_page_setup_new_from_key_file(cValue0, cValue1, cError)
}

func Fn_gtk_page_setup_copy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	C.gtk_page_setup_copy(cValueInstance)
}

func Fn_gtk_page_setup_get_bottom_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_page_setup_get_bottom_margin(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_get_left_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_page_setup_get_left_margin(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_get_orientation(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	C.gtk_page_setup_get_orientation(cValueInstance)
}

func Fn_gtk_page_setup_get_page_height(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_page_setup_get_page_height(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_get_page_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_page_setup_get_page_width(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_get_paper_height(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_page_setup_get_paper_height(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_get_paper_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	C.gtk_page_setup_get_paper_size(cValueInstance)
}

func Fn_gtk_page_setup_get_paper_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_page_setup_get_paper_width(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_get_right_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_page_setup_get_right_margin(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_get_top_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_page_setup_get_top_margin(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_load_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_page_setup_load_file(cValueInstance, cValue0, cError)
}

func Fn_gtk_page_setup_load_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	C.gtk_page_setup_load_key_file(cValueInstance, cValue0, cValue1, cError)
}

func Fn_gtk_page_setup_set_bottom_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_page_setup_set_bottom_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_page_setup_set_left_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_page_setup_set_left_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_page_setup_set_orientation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPageOrientation)(param0)

	C.gtk_page_setup_set_orientation(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_set_paper_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPaperSize)(unsafe.Pointer(param0))

	C.gtk_page_setup_set_paper_size(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_set_paper_size_and_default_margins(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPaperSize)(unsafe.Pointer(param0))

	C.gtk_page_setup_set_paper_size_and_default_margins(cValueInstance, cValue0)
}

func Fn_gtk_page_setup_set_right_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_page_setup_set_right_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_page_setup_set_top_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_page_setup_set_top_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_page_setup_to_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_page_setup_to_file(cValueInstance, cValue0, cError)
}

func Fn_gtk_page_setup_to_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_page_setup_to_key_file(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_paned_add1(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_paned_add1(cValueInstance, cValue0)
}

func Fn_gtk_paned_add2(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_paned_add2(cValueInstance, cValue0)
}

func Fn_gtk_paned_get_child1(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	C.gtk_paned_get_child1(cValueInstance)
}

func Fn_gtk_paned_get_child2(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	C.gtk_paned_get_child2(cValueInstance)
}

func Fn_gtk_paned_get_position(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	C.gtk_paned_get_position(cValueInstance)
}

func Fn_gtk_paned_pack1(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool) {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	C.gtk_paned_pack1(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_paned_pack2(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool) {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	C.gtk_paned_pack2(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_paned_set_position(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_paned_set_position(cValueInstance, cValue0)
}

func Fn_gtk_places_sidebar_get_show_connect_to_server(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	C.gtk_places_sidebar_get_show_connect_to_server(cValueInstance)
}

func Fn_gtk_plug_new(param0 uint64) {
	cValue0 := (C.Window)(param0)

	C.gtk_plug_new(cValue0)
}

func Fn_gtk_plug_new_for_display(param0 unsafe.Pointer, param1 uint64) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.Window)(param1)

	C.gtk_plug_new_for_display(cValue0, cValue1)
}

func Fn_gtk_plug_construct(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	cValue0 := (C.Window)(param0)

	C.gtk_plug_construct(cValueInstance, cValue0)
}

func Fn_gtk_plug_construct_for_display(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint64) {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.Window)(param1)

	C.gtk_plug_construct_for_display(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_plug_get_embedded(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	C.gtk_plug_get_embedded(cValueInstance)
}

func Fn_gtk_plug_get_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	C.gtk_plug_get_id(cValueInstance)
}

func Fn_gtk_plug_get_socket_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	C.gtk_plug_get_socket_window(cValueInstance)
}

func Fn_gtk_popover_get_pointing_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_popover_get_pointing_to(cValueInstance, cValue0)
}

func Fn_gtk_popover_get_position(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	C.gtk_popover_get_position(cValueInstance)
}

func Fn_gtk_print_context_create_pango_context(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	C.gtk_print_context_create_pango_context(cValueInstance)
}

func Fn_gtk_print_context_create_pango_layout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	C.gtk_print_context_create_pango_layout(cValueInstance)
}

func Fn_gtk_print_context_get_cairo_context(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	C.gtk_print_context_get_cairo_context(cValueInstance)
}

func Fn_gtk_print_context_get_dpi_x(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	C.gtk_print_context_get_dpi_x(cValueInstance)
}

func Fn_gtk_print_context_get_dpi_y(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	C.gtk_print_context_get_dpi_y(cValueInstance)
}

func Fn_gtk_print_context_get_height(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	C.gtk_print_context_get_height(cValueInstance)
}

func Fn_gtk_print_context_get_page_setup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	C.gtk_print_context_get_page_setup(cValueInstance)
}

func Fn_gtk_print_context_get_pango_fontmap(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	C.gtk_print_context_get_pango_fontmap(cValueInstance)
}

func Fn_gtk_print_context_get_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	C.gtk_print_context_get_width(cValueInstance)
}

func Fn_gtk_print_context_set_cairo_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 float64) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (C.double)(param1)

	cValue2 := (C.double)(param2)

	C.gtk_print_context_set_cairo_context(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_print_operation_new() {
	C.gtk_print_operation_new()
}

func Fn_gtk_print_operation_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_cancel(cValueInstance)
}

func Fn_gtk_print_operation_draw_page_finish(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_draw_page_finish(cValueInstance)
}

func Fn_gtk_print_operation_get_default_page_setup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_get_default_page_setup(cValueInstance)
}

func Fn_gtk_print_operation_get_embed_page_setup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_get_embed_page_setup(cValueInstance)
}

func Fn_gtk_print_operation_get_error(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	C.gtk_print_operation_get_error(cValueInstance, cError)
}

func Fn_gtk_print_operation_get_has_selection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_get_has_selection(cValueInstance)
}

func Fn_gtk_print_operation_get_n_pages_to_print(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_get_n_pages_to_print(cValueInstance)
}

func Fn_gtk_print_operation_get_print_settings(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_get_print_settings(cValueInstance)
}

func Fn_gtk_print_operation_get_status(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_get_status(cValueInstance)
}

func Fn_gtk_print_operation_get_status_string(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_get_status_string(cValueInstance)
}

func Fn_gtk_print_operation_get_support_selection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_get_support_selection(cValueInstance)
}

func Fn_gtk_print_operation_is_finished(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_is_finished(cValueInstance)
}

func Fn_gtk_print_operation_run(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPrintOperationAction)(param0)

	cValue1 := (*C.GtkWindow)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	C.gtk_print_operation_run(cValueInstance, cValue0, cValue1, cError)
}

func Fn_gtk_print_operation_set_allow_async(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_allow_async(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_current_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_operation_set_current_page(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_custom_tab_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_operation_set_custom_tab_label(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_default_page_setup(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPageSetup)(unsafe.Pointer(param0))

	C.gtk_print_operation_set_default_page_setup(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_defer_drawing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_set_defer_drawing(cValueInstance)
}

func Fn_gtk_print_operation_set_embed_page_setup(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_embed_page_setup(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_export_filename(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_operation_set_export_filename(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_has_selection(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_has_selection(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_job_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_operation_set_job_name(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_n_pages(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_operation_set_n_pages(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_print_settings(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPrintSettings)(unsafe.Pointer(param0))

	C.gtk_print_operation_set_print_settings(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_show_progress(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_show_progress(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_support_selection(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_support_selection(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_track_print_status(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_track_print_status(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_unit(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_print_operation_set_unit(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_set_use_full_page(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_operation_set_use_full_page(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_new() {
	C.gtk_print_settings_new()
}

func Fn_gtk_print_settings_new_from_file(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_print_settings_new_from_file(cValue0, cError)
}

func Fn_gtk_print_settings_new_from_key_file(param0 unsafe.Pointer, param1 string, error unsafe.Pointer) {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	C.gtk_print_settings_new_from_key_file(cValue0, cValue1, cError)
}

func Fn_gtk_print_settings_copy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_copy(cValueInstance)
}

// UNSUPPORTED : foreach : has callback

func Fn_gtk_print_settings_get(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_get(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_get_bool(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_get_bool(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_get_collate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_collate(cValueInstance)
}

func Fn_gtk_print_settings_get_default_source(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_default_source(cValueInstance)
}

func Fn_gtk_print_settings_get_dither(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_dither(cValueInstance)
}

func Fn_gtk_print_settings_get_double(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_get_double(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_get_double_with_default(paramInstance unsafe.Pointer, param0 string, param1 float64) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	C.gtk_print_settings_get_double_with_default(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_get_duplex(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_duplex(cValueInstance)
}

func Fn_gtk_print_settings_get_finishings(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_finishings(cValueInstance)
}

func Fn_gtk_print_settings_get_int(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_get_int(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_get_int_with_default(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_print_settings_get_int_with_default(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_get_length(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_print_settings_get_length(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_get_media_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_media_type(cValueInstance)
}

func Fn_gtk_print_settings_get_n_copies(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_n_copies(cValueInstance)
}

func Fn_gtk_print_settings_get_number_up(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_number_up(cValueInstance)
}

func Fn_gtk_print_settings_get_number_up_layout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_number_up_layout(cValueInstance)
}

func Fn_gtk_print_settings_get_orientation(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_orientation(cValueInstance)
}

func Fn_gtk_print_settings_get_output_bin(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_output_bin(cValueInstance)
}

func Fn_gtk_print_settings_get_page_ranges(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	C.gtk_print_settings_get_page_ranges(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_get_page_set(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_page_set(cValueInstance)
}

func Fn_gtk_print_settings_get_paper_height(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_print_settings_get_paper_height(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_get_paper_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_paper_size(cValueInstance)
}

func Fn_gtk_print_settings_get_paper_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	C.gtk_print_settings_get_paper_width(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_get_print_pages(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_print_pages(cValueInstance)
}

func Fn_gtk_print_settings_get_printer(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_printer(cValueInstance)
}

func Fn_gtk_print_settings_get_printer_lpi(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_printer_lpi(cValueInstance)
}

func Fn_gtk_print_settings_get_quality(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_quality(cValueInstance)
}

func Fn_gtk_print_settings_get_resolution(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_resolution(cValueInstance)
}

func Fn_gtk_print_settings_get_resolution_x(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_resolution_x(cValueInstance)
}

func Fn_gtk_print_settings_get_resolution_y(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_resolution_y(cValueInstance)
}

func Fn_gtk_print_settings_get_reverse(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_reverse(cValueInstance)
}

func Fn_gtk_print_settings_get_scale(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_scale(cValueInstance)
}

func Fn_gtk_print_settings_get_use_color(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	C.gtk_print_settings_get_use_color(cValueInstance)
}

func Fn_gtk_print_settings_has_key(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_has_key(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_load_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_print_settings_load_file(cValueInstance, cValue0, cError)
}

func Fn_gtk_print_settings_load_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	C.gtk_print_settings_load_key_file(cValueInstance, cValue0, cValue1, cError)
}

func Fn_gtk_print_settings_set(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_print_settings_set(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_bool(paramInstance unsafe.Pointer, param0 string, param1 bool) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	C.gtk_print_settings_set_bool(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_collate(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_settings_set_collate(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_default_source(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_default_source(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_dither(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_dither(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_double(paramInstance unsafe.Pointer, param0 string, param1 float64) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	C.gtk_print_settings_set_double(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_duplex(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPrintDuplex)(param0)

	C.gtk_print_settings_set_duplex(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_finishings(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_finishings(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_int(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_print_settings_set_int(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_length(paramInstance unsafe.Pointer, param0 string, param1 float64, param2 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.GtkUnit)(param2)

	C.gtk_print_settings_set_length(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_print_settings_set_media_type(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_media_type(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_n_copies(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_settings_set_n_copies(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_number_up(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_settings_set_number_up(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_number_up_layout(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkNumberUpLayout)(param0)

	C.gtk_print_settings_set_number_up_layout(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_orientation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPageOrientation)(param0)

	C.gtk_print_settings_set_orientation(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_output_bin(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_output_bin(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_page_ranges(paramInstance unsafe.Pointer, param0 []PageRange, param1 int) {
	// has non-string array param
}

func Fn_gtk_print_settings_set_page_set(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPageSet)(param0)

	C.gtk_print_settings_set_page_set(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_paper_height(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_print_settings_set_paper_height(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_paper_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPaperSize)(unsafe.Pointer(param0))

	C.gtk_print_settings_set_paper_size(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_paper_width(paramInstance unsafe.Pointer, param0 float64, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkUnit)(param1)

	C.gtk_print_settings_set_paper_width(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_print_pages(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPrintPages)(param0)

	C.gtk_print_settings_set_print_pages(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_printer(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_set_printer(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_printer_lpi(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_print_settings_set_printer_lpi(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_quality(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPrintQuality)(param0)

	C.gtk_print_settings_set_quality(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_resolution(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_settings_set_resolution(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_resolution_xy(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_print_settings_set_resolution_xy(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_set_reverse(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_settings_set_reverse(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_scale(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_print_settings_set_scale(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_set_use_color(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_print_settings_set_use_color(cValueInstance, cValue0)
}

func Fn_gtk_print_settings_to_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_print_settings_to_file(cValueInstance, cValue0, cError)
}

func Fn_gtk_print_settings_to_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_print_settings_to_key_file(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_print_settings_unset(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_print_settings_unset(cValueInstance, cValue0)
}

func Fn_gtk_progress_bar_new() {
	C.gtk_progress_bar_new()
}

func Fn_gtk_progress_bar_get_ellipsize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	C.gtk_progress_bar_get_ellipsize(cValueInstance)
}

func Fn_gtk_progress_bar_get_fraction(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	C.gtk_progress_bar_get_fraction(cValueInstance)
}

func Fn_gtk_progress_bar_get_inverted(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	C.gtk_progress_bar_get_inverted(cValueInstance)
}

func Fn_gtk_progress_bar_get_pulse_step(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	C.gtk_progress_bar_get_pulse_step(cValueInstance)
}

func Fn_gtk_progress_bar_get_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	C.gtk_progress_bar_get_text(cValueInstance)
}

func Fn_gtk_progress_bar_pulse(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	C.gtk_progress_bar_pulse(cValueInstance)
}

func Fn_gtk_progress_bar_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoEllipsizeMode)(param0)

	C.gtk_progress_bar_set_ellipsize(cValueInstance, cValue0)
}

func Fn_gtk_progress_bar_set_fraction(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_progress_bar_set_fraction(cValueInstance, cValue0)
}

func Fn_gtk_progress_bar_set_inverted(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_progress_bar_set_inverted(cValueInstance, cValue0)
}

func Fn_gtk_progress_bar_set_pulse_step(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_progress_bar_set_pulse_step(cValueInstance, cValue0)
}

func Fn_gtk_progress_bar_set_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_progress_bar_set_text(cValueInstance, cValue0)
}

func Fn_gtk_radio_action_new(param0 string, param1 string, param2 string, param3 string, param4 int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (C.gint)(param4)

	C.gtk_radio_action_new(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_radio_action_get_current_value(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	C.gtk_radio_action_get_current_value(cValueInstance)
}

func Fn_gtk_radio_action_get_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	C.gtk_radio_action_get_group(cValueInstance)
}

func Fn_gtk_radio_action_set_current_value(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_radio_action_set_current_value(cValueInstance, cValue0)
}

func Fn_gtk_radio_action_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_action_set_group(cValueInstance, cValue0)
}

func Fn_gtk_radio_button_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_button_new(cValue0)
}

func Fn_gtk_radio_button_new_from_widget(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkRadioButton)(unsafe.Pointer(param0))

	C.gtk_radio_button_new_from_widget(cValue0)
}

func Fn_gtk_radio_button_new_with_label(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_button_new_with_label(cValue0, cValue1)
}

func Fn_gtk_radio_button_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkRadioButton)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_button_new_with_label_from_widget(cValue0, cValue1)
}

func Fn_gtk_radio_button_new_with_mnemonic(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_button_new_with_mnemonic(cValue0, cValue1)
}

func Fn_gtk_radio_button_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkRadioButton)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_button_new_with_mnemonic_from_widget(cValue0, cValue1)
}

func Fn_gtk_radio_button_get_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioButton)(unsafe.Pointer(paramInstance))

	C.gtk_radio_button_get_group(cValueInstance)
}

func Fn_gtk_radio_button_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_button_set_group(cValueInstance, cValue0)
}

func Fn_gtk_radio_menu_item_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_menu_item_new(cValue0)
}

func Fn_gtk_radio_menu_item_new_from_widget(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	C.gtk_radio_menu_item_new_from_widget(cValue0)
}

func Fn_gtk_radio_menu_item_new_with_label(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_menu_item_new_with_label(cValue0, cValue1)
}

func Fn_gtk_radio_menu_item_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_menu_item_new_with_label_from_widget(cValue0, cValue1)
}

func Fn_gtk_radio_menu_item_new_with_mnemonic(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_menu_item_new_with_mnemonic(cValue0, cValue1)
}

func Fn_gtk_radio_menu_item_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_menu_item_new_with_mnemonic_from_widget(cValue0, cValue1)
}

func Fn_gtk_radio_menu_item_get_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_radio_menu_item_get_group(cValueInstance)
}

func Fn_gtk_radio_menu_item_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_menu_item_set_group(cValueInstance, cValue0)
}

func Fn_gtk_radio_tool_button_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_tool_button_new(cValue0)
}

func Fn_gtk_radio_tool_button_new_from_stock(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_tool_button_new_from_stock(cValue0, cValue1)
}

func Fn_gtk_radio_tool_button_new_from_widget(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkRadioToolButton)(unsafe.Pointer(param0))

	C.gtk_radio_tool_button_new_from_widget(cValue0)
}

func Fn_gtk_radio_tool_button_new_with_stock_from_widget(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkRadioToolButton)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_radio_tool_button_new_with_stock_from_widget(cValue0, cValue1)
}

func Fn_gtk_radio_tool_button_get_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioToolButton)(unsafe.Pointer(paramInstance))

	C.gtk_radio_tool_button_get_group(cValueInstance)
}

func Fn_gtk_radio_tool_button_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_tool_button_set_group(cValueInstance, cValue0)
}

func Fn_gtk_range_get_adjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	C.gtk_range_get_adjustment(cValueInstance)
}

func Fn_gtk_range_get_fill_level(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	C.gtk_range_get_fill_level(cValueInstance)
}

func Fn_gtk_range_get_flippable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	C.gtk_range_get_flippable(cValueInstance)
}

func Fn_gtk_range_get_inverted(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	C.gtk_range_get_inverted(cValueInstance)
}

func Fn_gtk_range_get_lower_stepper_sensitivity(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	C.gtk_range_get_lower_stepper_sensitivity(cValueInstance)
}

func Fn_gtk_range_get_restrict_to_fill_level(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	C.gtk_range_get_restrict_to_fill_level(cValueInstance)
}

func Fn_gtk_range_get_show_fill_level(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	C.gtk_range_get_show_fill_level(cValueInstance)
}

func Fn_gtk_range_get_upper_stepper_sensitivity(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	C.gtk_range_get_upper_stepper_sensitivity(cValueInstance)
}

func Fn_gtk_range_get_value(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	C.gtk_range_get_value(cValueInstance)
}

func Fn_gtk_range_set_adjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_range_set_adjustment(cValueInstance, cValue0)
}

func Fn_gtk_range_set_fill_level(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_range_set_fill_level(cValueInstance, cValue0)
}

func Fn_gtk_range_set_flippable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_flippable(cValueInstance, cValue0)
}

func Fn_gtk_range_set_increments(paramInstance unsafe.Pointer, param0 float64, param1 float64) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	C.gtk_range_set_increments(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_range_set_inverted(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_inverted(cValueInstance, cValue0)
}

func Fn_gtk_range_set_lower_stepper_sensitivity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSensitivityType)(param0)

	C.gtk_range_set_lower_stepper_sensitivity(cValueInstance, cValue0)
}

func Fn_gtk_range_set_range(paramInstance unsafe.Pointer, param0 float64, param1 float64) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	C.gtk_range_set_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_range_set_restrict_to_fill_level(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_restrict_to_fill_level(cValueInstance, cValue0)
}

func Fn_gtk_range_set_show_fill_level(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_show_fill_level(cValueInstance, cValue0)
}

func Fn_gtk_range_set_upper_stepper_sensitivity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSensitivityType)(param0)

	C.gtk_range_set_upper_stepper_sensitivity(cValueInstance, cValue0)
}

func Fn_gtk_range_set_value(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_range_set_value(cValueInstance, cValue0)
}

func Fn_gtk_rc_style_new() {
	C.gtk_rc_style_new()
}

func Fn_gtk_rc_style_copy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRcStyle)(unsafe.Pointer(paramInstance))

	C.gtk_rc_style_copy(cValueInstance)
}

func Fn_gtk_recent_action_new(param0 string, param1 string, param2 string, param3 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	C.gtk_recent_action_new(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_recent_action_new_for_manager(param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.GtkRecentManager)(unsafe.Pointer(param4))

	C.gtk_recent_action_new_for_manager(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_recent_action_get_show_numbers(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentAction)(unsafe.Pointer(paramInstance))

	C.gtk_recent_action_get_show_numbers(cValueInstance)
}

func Fn_gtk_recent_action_set_show_numbers(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_action_set_show_numbers(cValueInstance, cValue0)
}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_for_manager : has varargs

func Fn_gtk_recent_chooser_menu_new() {
	C.gtk_recent_chooser_menu_new()
}

func Fn_gtk_recent_chooser_menu_new_for_manager(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkRecentManager)(unsafe.Pointer(param0))

	C.gtk_recent_chooser_menu_new_for_manager(cValue0)
}

func Fn_gtk_recent_chooser_menu_get_show_numbers(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooserMenu)(unsafe.Pointer(paramInstance))

	C.gtk_recent_chooser_menu_get_show_numbers(cValueInstance)
}

func Fn_gtk_recent_chooser_menu_set_show_numbers(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooserMenu)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_menu_set_show_numbers(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_widget_new() {
	C.gtk_recent_chooser_widget_new()
}

func Fn_gtk_recent_chooser_widget_new_for_manager(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkRecentManager)(unsafe.Pointer(param0))

	C.gtk_recent_chooser_widget_new_for_manager(cValue0)
}

func Fn_gtk_recent_filter_new() {
	C.gtk_recent_filter_new()
}

func Fn_gtk_recent_filter_add_age(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_recent_filter_add_age(cValueInstance, cValue0)
}

func Fn_gtk_recent_filter_add_application(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_add_application(cValueInstance, cValue0)
}

// UNSUPPORTED : add_custom : has callback

func Fn_gtk_recent_filter_add_group(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_add_group(cValueInstance, cValue0)
}

func Fn_gtk_recent_filter_add_mime_type(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_add_mime_type(cValueInstance, cValue0)
}

func Fn_gtk_recent_filter_add_pattern(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_add_pattern(cValueInstance, cValue0)
}

func Fn_gtk_recent_filter_add_pixbuf_formats(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	C.gtk_recent_filter_add_pixbuf_formats(cValueInstance)
}

func Fn_gtk_recent_filter_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentFilterInfo)(unsafe.Pointer(param0))

	C.gtk_recent_filter_filter(cValueInstance, cValue0)
}

func Fn_gtk_recent_filter_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	C.gtk_recent_filter_get_name(cValueInstance)
}

func Fn_gtk_recent_filter_get_needed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	C.gtk_recent_filter_get_needed(cValueInstance)
}

func Fn_gtk_recent_filter_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_set_name(cValueInstance, cValue0)
}

func Fn_gtk_recent_manager_new() {
	C.gtk_recent_manager_new()
}

func Fn_gtk_recent_manager_add_full(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkRecentData)(unsafe.Pointer(param1))

	C.gtk_recent_manager_add_full(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_recent_manager_add_item(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_manager_add_item(cValueInstance, cValue0)
}

func Fn_gtk_recent_manager_get_items(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	C.gtk_recent_manager_get_items(cValueInstance)
}

func Fn_gtk_recent_manager_has_item(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_manager_has_item(cValueInstance, cValue0)
}

func Fn_gtk_recent_manager_lookup_item(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_recent_manager_lookup_item(cValueInstance, cValue0, cError)
}

func Fn_gtk_recent_manager_move_item(paramInstance unsafe.Pointer, param0 string, param1 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	C.gtk_recent_manager_move_item(cValueInstance, cValue0, cValue1, cError)
}

func Fn_gtk_recent_manager_purge_items(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	C.gtk_recent_manager_purge_items(cValueInstance, cError)
}

func Fn_gtk_recent_manager_remove_item(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_recent_manager_remove_item(cValueInstance, cValue0, cError)
}

func Fn_gtk_recent_manager_get_default() {
	C.gtk_recent_manager_get_default()
}

func Fn_gtk_renderer_cell_accessible_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_renderer_cell_accessible_new(cValue0)
}

func Fn_gtk_scale_add_mark(paramInstance unsafe.Pointer, param0 float64, param1 int, param2 string) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.GtkPositionType)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_scale_add_mark(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_scale_clear_marks(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	C.gtk_scale_clear_marks(cValueInstance)
}

func Fn_gtk_scale_get_digits(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	C.gtk_scale_get_digits(cValueInstance)
}

func Fn_gtk_scale_get_draw_value(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	C.gtk_scale_get_draw_value(cValueInstance)
}

func Fn_gtk_scale_get_layout(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	C.gtk_scale_get_layout(cValueInstance)
}

func Fn_gtk_scale_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_scale_get_layout_offsets(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_scale_get_value_pos(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	C.gtk_scale_get_value_pos(cValueInstance)
}

func Fn_gtk_scale_set_digits(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_scale_set_digits(cValueInstance, cValue0)
}

func Fn_gtk_scale_set_draw_value(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_scale_set_draw_value(cValueInstance, cValue0)
}

func Fn_gtk_scale_set_value_pos(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPositionType)(param0)

	C.gtk_scale_set_value_pos(cValueInstance, cValue0)
}

func Fn_gtk_scale_button_new(param0 int, param1 float64, param2 float64, param3 float64, param4 []string) {
	cValue0 := (C.GtkIconSize)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	param4Len := len(param4)
	cValue4Array := C.malloc((C.ulong)(param4Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue4Array))
	param4Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue4Array))[:param4Len:param4Len]
	for param4i, param4String := range param4 {
		param4Slice[param4i] = (*C.gchar)(C.CString(param4String))
		defer C.free(unsafe.Pointer(param4Slice[param4i]))
	}
	cValue4 := &param4Slice[0]

	C.gtk_scale_button_new(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_scale_button_get_adjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	C.gtk_scale_button_get_adjustment(cValueInstance)
}

func Fn_gtk_scale_button_get_minus_button(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	C.gtk_scale_button_get_minus_button(cValueInstance)
}

func Fn_gtk_scale_button_get_plus_button(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	C.gtk_scale_button_get_plus_button(cValueInstance)
}

func Fn_gtk_scale_button_get_popup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	C.gtk_scale_button_get_popup(cValueInstance)
}

func Fn_gtk_scale_button_get_value(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	C.gtk_scale_button_get_value(cValueInstance)
}

func Fn_gtk_scale_button_set_adjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_scale_button_set_adjustment(cValueInstance, cValue0)
}

func Fn_gtk_scale_button_set_icons(paramInstance unsafe.Pointer, param0 []string) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	C.gtk_scale_button_set_icons(cValueInstance, cValue0)
}

func Fn_gtk_scale_button_set_value(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_scale_button_set_value(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	C.gtk_scrolled_window_new(cValue0, cValue1)
}

func Fn_gtk_scrolled_window_add_with_viewport(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_scrolled_window_add_with_viewport(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_get_hadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	C.gtk_scrolled_window_get_hadjustment(cValueInstance)
}

func Fn_gtk_scrolled_window_get_hscrollbar(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	C.gtk_scrolled_window_get_hscrollbar(cValueInstance)
}

func Fn_gtk_scrolled_window_get_placement(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	C.gtk_scrolled_window_get_placement(cValueInstance)
}

func Fn_gtk_scrolled_window_get_policy(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPolicyType)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkPolicyType)(unsafe.Pointer(param1))

	C.gtk_scrolled_window_get_policy(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_scrolled_window_get_shadow_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	C.gtk_scrolled_window_get_shadow_type(cValueInstance)
}

func Fn_gtk_scrolled_window_get_vadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	C.gtk_scrolled_window_get_vadjustment(cValueInstance)
}

func Fn_gtk_scrolled_window_get_vscrollbar(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	C.gtk_scrolled_window_get_vscrollbar(cValueInstance)
}

func Fn_gtk_scrolled_window_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_scrolled_window_set_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_placement(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkCornerType)(param0)

	C.gtk_scrolled_window_set_placement(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_policy(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPolicyType)(param0)

	cValue1 := (C.GtkPolicyType)(param1)

	C.gtk_scrolled_window_set_policy(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_scrolled_window_set_shadow_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkShadowType)(param0)

	C.gtk_scrolled_window_set_shadow_type(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_scrolled_window_set_vadjustment(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_unset_placement(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	C.gtk_scrolled_window_unset_placement(cValueInstance)
}

func Fn_gtk_separator_menu_item_new() {
	C.gtk_separator_menu_item_new()
}

func Fn_gtk_separator_tool_item_new() {
	C.gtk_separator_tool_item_new()
}

func Fn_gtk_separator_tool_item_get_draw(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSeparatorToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_separator_tool_item_get_draw(cValueInstance)
}

func Fn_gtk_separator_tool_item_set_draw(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSeparatorToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_separator_tool_item_set_draw(cValueInstance, cValue0)
}

func Fn_gtk_settings_set_double_property(paramInstance unsafe.Pointer, param0 string, param1 float64, param2 string) {
	cValueInstance := (*C.GtkSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_settings_set_double_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_settings_set_long_property(paramInstance unsafe.Pointer, param0 string, param1 int64, param2 string) {
	cValueInstance := (*C.GtkSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.glong)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_settings_set_long_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_settings_set_property_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkSettingsValue)(unsafe.Pointer(param1))

	C.gtk_settings_set_property_value(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_settings_set_string_property(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string) {
	cValueInstance := (*C.GtkSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_settings_set_string_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_settings_get_default() {
	C.gtk_settings_get_default()
}

func Fn_gtk_settings_get_for_screen(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_settings_get_for_screen(cValue0)
}

func Fn_gtk_settings_install_property(param0 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	C.gtk_settings_install_property(cValue0)
}

// UNSUPPORTED : install_property_parser : has callback

func Fn_gtk_size_group_new(param0 int) {
	cValue0 := (C.GtkSizeGroupMode)(param0)

	C.gtk_size_group_new(cValue0)
}

func Fn_gtk_size_group_add_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_size_group_add_widget(cValueInstance, cValue0)
}

func Fn_gtk_size_group_get_ignore_hidden(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	C.gtk_size_group_get_ignore_hidden(cValueInstance)
}

func Fn_gtk_size_group_get_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	C.gtk_size_group_get_mode(cValueInstance)
}

func Fn_gtk_size_group_get_widgets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	C.gtk_size_group_get_widgets(cValueInstance)
}

func Fn_gtk_size_group_remove_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_size_group_remove_widget(cValueInstance, cValue0)
}

func Fn_gtk_size_group_set_ignore_hidden(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_size_group_set_ignore_hidden(cValueInstance, cValue0)
}

func Fn_gtk_size_group_set_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSizeGroupMode)(param0)

	C.gtk_size_group_set_mode(cValueInstance, cValue0)
}

func Fn_gtk_socket_new() {
	C.gtk_socket_new()
}

func Fn_gtk_socket_add_id(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GtkSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.Window)(param0)

	C.gtk_socket_add_id(cValueInstance, cValue0)
}

func Fn_gtk_socket_get_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSocket)(unsafe.Pointer(paramInstance))

	C.gtk_socket_get_id(cValueInstance)
}

func Fn_gtk_socket_get_plug_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSocket)(unsafe.Pointer(paramInstance))

	C.gtk_socket_get_plug_window(cValueInstance)
}

func Fn_gtk_spin_button_new(param0 unsafe.Pointer, param1 float64, param2 uint) {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.guint)(param2)

	C.gtk_spin_button_new(cValue0, cValue1, cValue2)
}

func Fn_gtk_spin_button_new_with_range(param0 float64, param1 float64, param2 float64) {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	C.gtk_spin_button_new_with_range(cValue0, cValue1, cValue2)
}

func Fn_gtk_spin_button_configure(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 uint) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.guint)(param2)

	C.gtk_spin_button_configure(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_spin_button_get_adjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	C.gtk_spin_button_get_adjustment(cValueInstance)
}

func Fn_gtk_spin_button_get_digits(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	C.gtk_spin_button_get_digits(cValueInstance)
}

func Fn_gtk_spin_button_get_increments(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	C.gtk_spin_button_get_increments(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_spin_button_get_numeric(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	C.gtk_spin_button_get_numeric(cValueInstance)
}

func Fn_gtk_spin_button_get_range(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	C.gtk_spin_button_get_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_spin_button_get_snap_to_ticks(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	C.gtk_spin_button_get_snap_to_ticks(cValueInstance)
}

func Fn_gtk_spin_button_get_update_policy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	C.gtk_spin_button_get_update_policy(cValueInstance)
}

func Fn_gtk_spin_button_get_value(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	C.gtk_spin_button_get_value(cValueInstance)
}

func Fn_gtk_spin_button_get_value_as_int(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	C.gtk_spin_button_get_value_as_int(cValueInstance)
}

func Fn_gtk_spin_button_get_wrap(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	C.gtk_spin_button_get_wrap(cValueInstance)
}

func Fn_gtk_spin_button_set_adjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_spin_button_set_adjustment(cValueInstance, cValue0)
}

func Fn_gtk_spin_button_set_digits(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_spin_button_set_digits(cValueInstance, cValue0)
}

func Fn_gtk_spin_button_set_increments(paramInstance unsafe.Pointer, param0 float64, param1 float64) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	C.gtk_spin_button_set_increments(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_spin_button_set_numeric(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_spin_button_set_numeric(cValueInstance, cValue0)
}

func Fn_gtk_spin_button_set_range(paramInstance unsafe.Pointer, param0 float64, param1 float64) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	C.gtk_spin_button_set_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_spin_button_set_snap_to_ticks(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_spin_button_set_snap_to_ticks(cValueInstance, cValue0)
}

func Fn_gtk_spin_button_set_update_policy(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSpinButtonUpdatePolicy)(param0)

	C.gtk_spin_button_set_update_policy(cValueInstance, cValue0)
}

func Fn_gtk_spin_button_set_value(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_spin_button_set_value(cValueInstance, cValue0)
}

func Fn_gtk_spin_button_set_wrap(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_spin_button_set_wrap(cValueInstance, cValue0)
}

func Fn_gtk_spin_button_spin(paramInstance unsafe.Pointer, param0 int, param1 float64) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSpinType)(param0)

	cValue1 := (C.gdouble)(param1)

	C.gtk_spin_button_spin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_spin_button_update(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	C.gtk_spin_button_update(cValueInstance)
}

func Fn_gtk_status_icon_new() {
	C.gtk_status_icon_new()
}

func Fn_gtk_status_icon_new_from_file(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_new_from_file(cValue0)
}

func Fn_gtk_status_icon_new_from_gicon(param0 unsafe.Pointer) {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.gtk_status_icon_new_from_gicon(cValue0)
}

func Fn_gtk_status_icon_new_from_icon_name(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_new_from_icon_name(cValue0)
}

func Fn_gtk_status_icon_new_from_pixbuf(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_status_icon_new_from_pixbuf(cValue0)
}

func Fn_gtk_status_icon_new_from_stock(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_new_from_stock(cValue0)
}

func Fn_gtk_status_icon_get_geometry(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, param2 *int) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkOrientation)(unsafe.Pointer(param2))

	C.gtk_status_icon_get_geometry(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_status_icon_get_gicon(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_gicon(cValueInstance)
}

func Fn_gtk_status_icon_get_has_tooltip(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_has_tooltip(cValueInstance)
}

func Fn_gtk_status_icon_get_icon_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_icon_name(cValueInstance)
}

func Fn_gtk_status_icon_get_pixbuf(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_pixbuf(cValueInstance)
}

func Fn_gtk_status_icon_get_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_screen(cValueInstance)
}

func Fn_gtk_status_icon_get_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_size(cValueInstance)
}

func Fn_gtk_status_icon_get_stock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_stock(cValueInstance)
}

func Fn_gtk_status_icon_get_storage_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_storage_type(cValueInstance)
}

func Fn_gtk_status_icon_get_title(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_title(cValueInstance)
}

func Fn_gtk_status_icon_get_tooltip_markup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_tooltip_markup(cValueInstance)
}

func Fn_gtk_status_icon_get_tooltip_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_tooltip_text(cValueInstance)
}

func Fn_gtk_status_icon_get_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_visible(cValueInstance)
}

func Fn_gtk_status_icon_get_x11_window_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_get_x11_window_id(cValueInstance)
}

func Fn_gtk_status_icon_is_embedded(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	C.gtk_status_icon_is_embedded(cValueInstance)
}

func Fn_gtk_status_icon_set_from_file(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_from_file(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_from_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.gtk_status_icon_set_from_gicon(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_from_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_from_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_from_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_status_icon_set_from_pixbuf(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_from_stock(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_from_stock(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_has_tooltip(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_status_icon_set_has_tooltip(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_status_icon_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_title(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_tooltip_markup(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_tooltip_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_tooltip_text(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_status_icon_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_status_icon_position_menu(param0 unsafe.Pointer, param1 *int, param2 *int, param3 *bool, param4 unsafe.Pointer) {
	cValue0 := (*C.GtkMenu)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gboolean)(unsafe.Pointer(param3))

	cValue4 := (C.gpointer)(param4)

	C.gtk_status_icon_position_menu(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_statusbar_new() {
	C.gtk_statusbar_new()
}

func Fn_gtk_statusbar_get_context_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_statusbar_get_context_id(cValueInstance, cValue0)
}

func Fn_gtk_statusbar_pop(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_statusbar_pop(cValueInstance, cValue0)
}

func Fn_gtk_statusbar_push(paramInstance unsafe.Pointer, param0 uint, param1 string) {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_statusbar_push(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_statusbar_remove(paramInstance unsafe.Pointer, param0 uint, param1 uint) {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	C.gtk_statusbar_remove(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_new() {
	C.gtk_style_new()
}

func Fn_gtk_style_apply_default_background(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int, param5 int, param6 int) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	cValue5 := (C.gint)(param5)

	cValue6 := (C.gint)(param6)

	C.gtk_style_apply_default_background(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gtk_style_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gtk_style_attach(cValueInstance, cValue0)
}

func Fn_gtk_style_copy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	C.gtk_style_copy(cValueInstance)
}

func Fn_gtk_style_detach(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	C.gtk_style_detach(cValueInstance)
}

// UNSUPPORTED : get : has varargs

func Fn_gtk_style_get_style_property(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_style_get_style_property(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_style_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gtk_style_lookup_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_lookup_icon_set(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_style_lookup_icon_set(cValueInstance, cValue0)
}

func Fn_gtk_style_render_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 unsafe.Pointer, param5 string) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkIconSource)(unsafe.Pointer(param0))

	cValue1 := (C.GtkTextDirection)(param1)

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkIconSize)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	C.gtk_style_render_icon(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_style_set_background(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GtkStateType)(param1)

	C.gtk_style_set_background(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_new() {
	C.gtk_style_context_new()
}

// UNSUPPORTED : get : has varargs

func Fn_gtk_style_context_get_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.gtk_style_context_get_screen(cValueInstance)
}

func Fn_gtk_style_context_get_section(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_style_context_get_section(cValueInstance, cValue0)
}

// UNSUPPORTED : get_style : has varargs

func Fn_gtk_style_context_get_style_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.gtk_style_context_get_style_property(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : get_style_valist : has va_list

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_style_context_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_style_context_lookup_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_lookup_icon_set(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_style_context_lookup_icon_set(cValueInstance, cValue0)
}

func Fn_gtk_style_properties_new() {
	C.gtk_style_properties_new()
}

func Fn_gtk_style_properties_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	C.gtk_style_properties_clear(cValueInstance)
}

// UNSUPPORTED : get : has varargs

// UNSUPPORTED : get_valist : has va_list

// UNSUPPORTED : set : has varargs

// UNSUPPORTED : set_valist : has va_list

// UNSUPPORTED : lookup_property : has callback

// UNSUPPORTED : register_property : has callback

func Fn_gtk_table_new(param0 uint, param1 uint, param2 bool) {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := toCBool(param2)

	C.gtk_table_new(cValue0, cValue1, cValue2)
}

func Fn_gtk_table_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint, param5 int, param6 int, param7 uint, param8 uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.guint)(param3)

	cValue4 := (C.guint)(param4)

	cValue5 := (C.GtkAttachOptions)(param5)

	cValue6 := (C.GtkAttachOptions)(param6)

	cValue7 := (C.guint)(param7)

	cValue8 := (C.guint)(param8)

	C.gtk_table_attach(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8)
}

func Fn_gtk_table_attach_defaults(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	cValue3 := (C.guint)(param3)

	cValue4 := (C.guint)(param4)

	C.gtk_table_attach_defaults(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_table_get_col_spacing(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_table_get_col_spacing(cValueInstance, cValue0)
}

func Fn_gtk_table_get_default_col_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	C.gtk_table_get_default_col_spacing(cValueInstance)
}

func Fn_gtk_table_get_default_row_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	C.gtk_table_get_default_row_spacing(cValueInstance)
}

func Fn_gtk_table_get_homogeneous(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	C.gtk_table_get_homogeneous(cValueInstance)
}

func Fn_gtk_table_get_row_spacing(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_table_get_row_spacing(cValueInstance, cValue0)
}

func Fn_gtk_table_resize(paramInstance unsafe.Pointer, param0 uint, param1 uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	C.gtk_table_resize(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_table_set_col_spacing(paramInstance unsafe.Pointer, param0 uint, param1 uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	C.gtk_table_set_col_spacing(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_table_set_col_spacings(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_table_set_col_spacings(cValueInstance, cValue0)
}

func Fn_gtk_table_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_table_set_homogeneous(cValueInstance, cValue0)
}

func Fn_gtk_table_set_row_spacing(paramInstance unsafe.Pointer, param0 uint, param1 uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	C.gtk_table_set_row_spacing(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_table_set_row_spacings(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_table_set_row_spacings(cValueInstance, cValue0)
}

func Fn_gtk_tearoff_menu_item_new() {
	C.gtk_tearoff_menu_item_new()
}

func Fn_gtk_text_buffer_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkTextTagTable)(unsafe.Pointer(param0))

	C.gtk_text_buffer_new(cValue0)
}

func Fn_gtk_text_buffer_add_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextMark)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_add_mark(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_add_selection_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkClipboard)(unsafe.Pointer(param0))

	C.gtk_text_buffer_add_selection_clipboard(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_apply_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	C.gtk_text_buffer_apply_tag(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_apply_tag_by_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	C.gtk_text_buffer_apply_tag_by_name(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_backspace(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	C.gtk_text_buffer_backspace(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_begin_user_action(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_begin_user_action(cValueInstance)
}

func Fn_gtk_text_buffer_copy_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkClipboard)(unsafe.Pointer(param0))

	C.gtk_text_buffer_copy_clipboard(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_create_child_anchor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_buffer_create_child_anchor(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_create_mark(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_text_buffer_create_mark(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : create_tag : has varargs

func Fn_gtk_text_buffer_cut_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkClipboard)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_text_buffer_cut_clipboard(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_delete(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_delete(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_delete_interactive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_text_buffer_delete_interactive(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_delete_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextMark)(unsafe.Pointer(param0))

	C.gtk_text_buffer_delete_mark(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_delete_mark_by_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_text_buffer_delete_mark_by_name(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_delete_selection(paramInstance unsafe.Pointer, param0 bool, param1 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	cValue1 := toCBool(param1)

	C.gtk_text_buffer_delete_selection(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_deserialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 gdk.Atom, param2 unsafe.Pointer, param3 []uint8, param4 uint64, error unsafe.Pointer) {
	// has non-string array param
}

func Fn_gtk_text_buffer_deserialize_get_can_create_tags(paramInstance unsafe.Pointer, param0 gdk.Atom) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	C.gtk_text_buffer_deserialize_get_can_create_tags(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_deserialize_set_can_create_tags(paramInstance unsafe.Pointer, param0 gdk.Atom, param1 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_text_buffer_deserialize_set_can_create_tags(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_end_user_action(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_end_user_action(cValueInstance)
}

func Fn_gtk_text_buffer_get_bounds(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_get_bounds(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_get_char_count(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_get_char_count(cValueInstance)
}

func Fn_gtk_text_buffer_get_copy_target_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_get_copy_target_list(cValueInstance)
}

func Fn_gtk_text_buffer_get_deserialize_formats(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	C.gtk_text_buffer_get_deserialize_formats(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_get_end_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_buffer_get_end_iter(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_get_has_selection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_get_has_selection(cValueInstance)
}

func Fn_gtk_text_buffer_get_insert(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_get_insert(cValueInstance)
}

func Fn_gtk_text_buffer_get_iter_at_child_anchor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextChildAnchor)(unsafe.Pointer(param1))

	C.gtk_text_buffer_get_iter_at_child_anchor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_get_iter_at_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_text_buffer_get_iter_at_line(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_get_iter_at_line_index(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_text_buffer_get_iter_at_line_index(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_get_iter_at_line_offset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_text_buffer_get_iter_at_line_offset(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_get_iter_at_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextMark)(unsafe.Pointer(param1))

	C.gtk_text_buffer_get_iter_at_mark(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_get_iter_at_offset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_text_buffer_get_iter_at_offset(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_get_line_count(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_get_line_count(cValueInstance)
}

func Fn_gtk_text_buffer_get_mark(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_text_buffer_get_mark(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_get_modified(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_get_modified(cValueInstance)
}

func Fn_gtk_text_buffer_get_paste_target_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_get_paste_target_list(cValueInstance)
}

func Fn_gtk_text_buffer_get_selection_bound(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_get_selection_bound(cValueInstance)
}

func Fn_gtk_text_buffer_get_selection_bounds(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_get_selection_bounds(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_get_serialize_formats(paramInstance unsafe.Pointer, param0 *int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	C.gtk_text_buffer_get_serialize_formats(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_get_slice(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_text_buffer_get_slice(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_get_start_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_buffer_get_start_iter(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_get_tag_table(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	C.gtk_text_buffer_get_tag_table(cValueInstance)
}

func Fn_gtk_text_buffer_get_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_text_buffer_get_text(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	C.gtk_text_buffer_insert(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_insert_at_cursor(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_text_buffer_insert_at_cursor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_insert_child_anchor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextChildAnchor)(unsafe.Pointer(param1))

	C.gtk_text_buffer_insert_child_anchor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_insert_interactive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int, param3 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	cValue3 := toCBool(param3)

	C.gtk_text_buffer_insert_interactive(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_text_buffer_insert_interactive_at_cursor(paramInstance unsafe.Pointer, param0 string, param1 int, param2 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := toCBool(param2)

	C.gtk_text_buffer_insert_interactive_at_cursor(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_insert_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	C.gtk_text_buffer_insert_pixbuf(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_insert_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	C.gtk_text_buffer_insert_range(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_insert_range_interactive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	cValue3 := toCBool(param3)

	C.gtk_text_buffer_insert_range_interactive(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : insert_with_tags : has varargs

// UNSUPPORTED : insert_with_tags_by_name : has varargs

func Fn_gtk_text_buffer_move_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextMark)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_move_mark(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_move_mark_by_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_move_mark_by_name(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_paste_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkClipboard)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_text_buffer_paste_clipboard(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_place_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_buffer_place_cursor(cValueInstance, cValue0)
}

// UNSUPPORTED : register_deserialize_format : has callback

func Fn_gtk_text_buffer_register_deserialize_tagset(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_text_buffer_register_deserialize_tagset(cValueInstance, cValue0)
}

// UNSUPPORTED : register_serialize_format : has callback

func Fn_gtk_text_buffer_register_serialize_tagset(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_text_buffer_register_serialize_tagset(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_remove_all_tags(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_remove_all_tags(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_remove_selection_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkClipboard)(unsafe.Pointer(param0))

	C.gtk_text_buffer_remove_selection_clipboard(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_remove_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	C.gtk_text_buffer_remove_tag(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_remove_tag_by_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	C.gtk_text_buffer_remove_tag_by_name(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_buffer_select_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	C.gtk_text_buffer_select_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_serialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 gdk.Atom, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 *uint64) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkTextIter)(unsafe.Pointer(param3))

	cValue4 := (*C.gsize)(unsafe.Pointer(param4))

	C.gtk_text_buffer_serialize(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_text_buffer_set_modified(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_buffer_set_modified(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	C.gtk_text_buffer_set_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_buffer_unregister_deserialize_format(paramInstance unsafe.Pointer, param0 gdk.Atom) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	C.gtk_text_buffer_unregister_deserialize_format(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_unregister_serialize_format(paramInstance unsafe.Pointer, param0 gdk.Atom) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	C.gtk_text_buffer_unregister_serialize_format(cValueInstance, cValue0)
}

func Fn_gtk_text_child_anchor_new() {
	C.gtk_text_child_anchor_new()
}

func Fn_gtk_text_child_anchor_get_deleted(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextChildAnchor)(unsafe.Pointer(paramInstance))

	C.gtk_text_child_anchor_get_deleted(cValueInstance)
}

func Fn_gtk_text_child_anchor_get_widgets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextChildAnchor)(unsafe.Pointer(paramInstance))

	C.gtk_text_child_anchor_get_widgets(cValueInstance)
}

func Fn_gtk_text_mark_new(param0 string, param1 bool) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	C.gtk_text_mark_new(cValue0, cValue1)
}

func Fn_gtk_text_mark_get_buffer(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	C.gtk_text_mark_get_buffer(cValueInstance)
}

func Fn_gtk_text_mark_get_deleted(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	C.gtk_text_mark_get_deleted(cValueInstance)
}

func Fn_gtk_text_mark_get_left_gravity(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	C.gtk_text_mark_get_left_gravity(cValueInstance)
}

func Fn_gtk_text_mark_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	C.gtk_text_mark_get_name(cValueInstance)
}

func Fn_gtk_text_mark_get_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	C.gtk_text_mark_get_visible(cValueInstance)
}

func Fn_gtk_text_mark_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_mark_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_text_tag_new(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_text_tag_new(cValue0)
}

func Fn_gtk_text_tag_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextTag)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkEvent)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	C.gtk_text_tag_event(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_tag_get_priority(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextTag)(unsafe.Pointer(paramInstance))

	C.gtk_text_tag_get_priority(cValueInstance)
}

func Fn_gtk_text_tag_set_priority(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextTag)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_tag_set_priority(cValueInstance, cValue0)
}

func Fn_gtk_text_tag_table_new() {
	C.gtk_text_tag_table_new()
}

func Fn_gtk_text_tag_table_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextTagTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	C.gtk_text_tag_table_add(cValueInstance, cValue0)
}

// UNSUPPORTED : foreach : has callback

func Fn_gtk_text_tag_table_get_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextTagTable)(unsafe.Pointer(paramInstance))

	C.gtk_text_tag_table_get_size(cValueInstance)
}

func Fn_gtk_text_tag_table_lookup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTextTagTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_text_tag_table_lookup(cValueInstance, cValue0)
}

func Fn_gtk_text_tag_table_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextTagTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	C.gtk_text_tag_table_remove(cValueInstance, cValue0)
}

func Fn_gtk_text_view_new() {
	C.gtk_text_view_new()
}

func Fn_gtk_text_view_new_with_buffer(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	C.gtk_text_view_new_with_buffer(cValue0)
}

func Fn_gtk_text_view_add_child_at_anchor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextChildAnchor)(unsafe.Pointer(param1))

	C.gtk_text_view_add_child_at_anchor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_view_add_child_in_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkTextWindowType)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_text_view_add_child_in_window(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_text_view_backward_display_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_view_backward_display_line(cValueInstance, cValue0)
}

func Fn_gtk_text_view_backward_display_line_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_view_backward_display_line_start(cValueInstance, cValue0)
}

func Fn_gtk_text_view_buffer_to_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 *int, param4 *int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextWindowType)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	C.gtk_text_view_buffer_to_window_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_text_view_forward_display_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_view_forward_display_line(cValueInstance, cValue0)
}

func Fn_gtk_text_view_forward_display_line_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_view_forward_display_line_end(cValueInstance, cValue0)
}

func Fn_gtk_text_view_get_accepts_tab(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_accepts_tab(cValueInstance)
}

func Fn_gtk_text_view_get_border_window_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextWindowType)(param0)

	C.gtk_text_view_get_border_window_size(cValueInstance, cValue0)
}

func Fn_gtk_text_view_get_buffer(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_buffer(cValueInstance)
}

func Fn_gtk_text_view_get_cursor_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_cursor_visible(cValueInstance)
}

func Fn_gtk_text_view_get_default_attributes(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_default_attributes(cValueInstance)
}

func Fn_gtk_text_view_get_editable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_editable(cValueInstance)
}

func Fn_gtk_text_view_get_indent(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_indent(cValueInstance)
}

func Fn_gtk_text_view_get_iter_at_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_text_view_get_iter_at_location(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_view_get_iter_at_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 int, param3 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_text_view_get_iter_at_position(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_text_view_get_iter_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gtk_text_view_get_iter_location(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_view_get_justification(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_justification(cValueInstance)
}

func Fn_gtk_text_view_get_left_margin(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_left_margin(cValueInstance)
}

func Fn_gtk_text_view_get_line_at_y(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_text_view_get_line_at_y(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_view_get_line_yrange(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_text_view_get_line_yrange(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_view_get_overwrite(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_overwrite(cValueInstance)
}

func Fn_gtk_text_view_get_pixels_above_lines(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_pixels_above_lines(cValueInstance)
}

func Fn_gtk_text_view_get_pixels_below_lines(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_pixels_below_lines(cValueInstance)
}

func Fn_gtk_text_view_get_pixels_inside_wrap(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_pixels_inside_wrap(cValueInstance)
}

func Fn_gtk_text_view_get_right_margin(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_right_margin(cValueInstance)
}

func Fn_gtk_text_view_get_tabs(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_tabs(cValueInstance)
}

func Fn_gtk_text_view_get_visible_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_text_view_get_visible_rect(cValueInstance, cValue0)
}

func Fn_gtk_text_view_get_window(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextWindowType)(param0)

	C.gtk_text_view_get_window(cValueInstance, cValue0)
}

func Fn_gtk_text_view_get_window_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gtk_text_view_get_window_type(cValueInstance, cValue0)
}

func Fn_gtk_text_view_get_wrap_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_get_wrap_mode(cValueInstance)
}

func Fn_gtk_text_view_move_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_text_view_move_child(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_view_move_mark_onscreen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextMark)(unsafe.Pointer(param0))

	C.gtk_text_view_move_mark_onscreen(cValueInstance, cValue0)
}

func Fn_gtk_text_view_move_visually(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_text_view_move_visually(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_view_place_cursor_onscreen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_place_cursor_onscreen(cValueInstance)
}

func Fn_gtk_text_view_scroll_mark_onscreen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextMark)(unsafe.Pointer(param0))

	C.gtk_text_view_scroll_mark_onscreen(cValueInstance, cValue0)
}

func Fn_gtk_text_view_scroll_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 bool, param3 float64, param4 float64) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	C.gtk_text_view_scroll_to_iter(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_text_view_scroll_to_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 bool, param3 float64, param4 float64) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextMark)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	C.gtk_text_view_scroll_to_mark(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_text_view_set_accepts_tab(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_view_set_accepts_tab(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_border_window_size(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextWindowType)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_text_view_set_border_window_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_view_set_buffer(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	C.gtk_text_view_set_buffer(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_cursor_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_view_set_cursor_visible(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_editable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_view_set_editable(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_indent(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_view_set_indent(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_justification(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkJustification)(param0)

	C.gtk_text_view_set_justification(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_left_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_view_set_left_margin(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_overwrite(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_view_set_overwrite(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_pixels_above_lines(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_view_set_pixels_above_lines(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_pixels_below_lines(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_view_set_pixels_below_lines(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_pixels_inside_wrap(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_view_set_pixels_inside_wrap(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_right_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_view_set_right_margin(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_tabs(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoTabArray)(unsafe.Pointer(param0))

	C.gtk_text_view_set_tabs(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_wrap_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkWrapMode)(param0)

	C.gtk_text_view_set_wrap_mode(cValueInstance, cValue0)
}

func Fn_gtk_text_view_starts_display_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_view_starts_display_line(cValueInstance, cValue0)
}

func Fn_gtk_text_view_window_to_buffer_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 *int, param4 *int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextWindowType)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	C.gtk_text_view_window_to_buffer_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

// UNSUPPORTED : get : has varargs

func Fn_gtk_theming_engine_get_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	C.gtk_theming_engine_get_screen(cValueInstance)
}

// UNSUPPORTED : get_style : has varargs

// UNSUPPORTED : get_style_valist : has va_list

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_theming_engine_load(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_theming_engine_load(cValue0)
}

// UNSUPPORTED : register_property : has callback

func Fn_gtk_toggle_action_new(param0 string, param1 string, param2 string, param3 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	C.gtk_toggle_action_new(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_toggle_action_get_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	C.gtk_toggle_action_get_active(cValueInstance)
}

func Fn_gtk_toggle_action_get_draw_as_radio(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	C.gtk_toggle_action_get_draw_as_radio(cValueInstance)
}

func Fn_gtk_toggle_action_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_action_set_active(cValueInstance, cValue0)
}

func Fn_gtk_toggle_action_set_draw_as_radio(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_action_set_draw_as_radio(cValueInstance, cValue0)
}

func Fn_gtk_toggle_action_toggled(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	C.gtk_toggle_action_toggled(cValueInstance)
}

func Fn_gtk_toggle_button_new() {
	C.gtk_toggle_button_new()
}

func Fn_gtk_toggle_button_new_with_label(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_toggle_button_new_with_label(cValue0)
}

func Fn_gtk_toggle_button_new_with_mnemonic(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_toggle_button_new_with_mnemonic(cValue0)
}

func Fn_gtk_toggle_button_get_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	C.gtk_toggle_button_get_active(cValueInstance)
}

func Fn_gtk_toggle_button_get_inconsistent(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	C.gtk_toggle_button_get_inconsistent(cValueInstance)
}

func Fn_gtk_toggle_button_get_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	C.gtk_toggle_button_get_mode(cValueInstance)
}

func Fn_gtk_toggle_button_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_button_set_active(cValueInstance, cValue0)
}

func Fn_gtk_toggle_button_set_inconsistent(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_button_set_inconsistent(cValueInstance, cValue0)
}

func Fn_gtk_toggle_button_set_mode(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_button_set_mode(cValueInstance, cValue0)
}

func Fn_gtk_toggle_button_toggled(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	C.gtk_toggle_button_toggled(cValueInstance)
}

func Fn_gtk_toggle_tool_button_new() {
	C.gtk_toggle_tool_button_new()
}

func Fn_gtk_toggle_tool_button_new_from_stock(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_toggle_tool_button_new_from_stock(cValue0)
}

func Fn_gtk_toggle_tool_button_get_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToggleToolButton)(unsafe.Pointer(paramInstance))

	C.gtk_toggle_tool_button_get_active(cValueInstance)
}

func Fn_gtk_toggle_tool_button_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_tool_button_set_active(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_new(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_tool_button_new(cValue0, cValue1)
}

func Fn_gtk_tool_button_new_from_stock(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_button_new_from_stock(cValue0)
}

func Fn_gtk_tool_button_get_icon_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	C.gtk_tool_button_get_icon_name(cValueInstance)
}

func Fn_gtk_tool_button_get_icon_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	C.gtk_tool_button_get_icon_widget(cValueInstance)
}

func Fn_gtk_tool_button_get_label(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	C.gtk_tool_button_get_label(cValueInstance)
}

func Fn_gtk_tool_button_get_label_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	C.gtk_tool_button_get_label_widget(cValueInstance)
}

func Fn_gtk_tool_button_get_stock_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	C.gtk_tool_button_get_stock_id(cValueInstance)
}

func Fn_gtk_tool_button_get_use_underline(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	C.gtk_tool_button_get_use_underline(cValueInstance)
}

func Fn_gtk_tool_button_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_button_set_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_set_icon_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_tool_button_set_icon_widget(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_button_set_label(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_tool_button_set_label_widget(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_set_stock_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_button_set_stock_id(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_button_set_use_underline(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_new() {
	C.gtk_tool_item_new()
}

func Fn_gtk_tool_item_get_expand(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_expand(cValueInstance)
}

func Fn_gtk_tool_item_get_homogeneous(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_homogeneous(cValueInstance)
}

func Fn_gtk_tool_item_get_icon_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_icon_size(cValueInstance)
}

func Fn_gtk_tool_item_get_is_important(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_is_important(cValueInstance)
}

func Fn_gtk_tool_item_get_orientation(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_orientation(cValueInstance)
}

func Fn_gtk_tool_item_get_proxy_menu_item(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_item_get_proxy_menu_item(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_get_relief_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_relief_style(cValueInstance)
}

func Fn_gtk_tool_item_get_toolbar_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_toolbar_style(cValueInstance)
}

func Fn_gtk_tool_item_get_use_drag_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_use_drag_window(cValueInstance)
}

func Fn_gtk_tool_item_get_visible_horizontal(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_visible_horizontal(cValueInstance)
}

func Fn_gtk_tool_item_get_visible_vertical(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_get_visible_vertical(cValueInstance)
}

func Fn_gtk_tool_item_rebuild_menu(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_rebuild_menu(cValueInstance)
}

func Fn_gtk_tool_item_retrieve_proxy_menu_item(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_retrieve_proxy_menu_item(cValueInstance)
}

func Fn_gtk_tool_item_set_expand(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_expand(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_homogeneous(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_is_important(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_is_important(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_proxy_menu_item(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	C.gtk_tool_item_set_proxy_menu_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_item_set_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_item_set_tooltip_markup(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_tooltip_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_item_set_tooltip_text(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_use_drag_window(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_use_drag_window(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_visible_horizontal(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_visible_horizontal(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_set_visible_vertical(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_set_visible_vertical(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_toolbar_reconfigured(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_toolbar_reconfigured(cValueInstance)
}

func Fn_gtk_toolbar_new() {
	C.gtk_toolbar_new()
}

func Fn_gtk_toolbar_get_drop_index(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_toolbar_get_drop_index(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_toolbar_get_icon_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	C.gtk_toolbar_get_icon_size(cValueInstance)
}

func Fn_gtk_toolbar_get_item_index(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	C.gtk_toolbar_get_item_index(cValueInstance, cValue0)
}

func Fn_gtk_toolbar_get_n_items(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	C.gtk_toolbar_get_n_items(cValueInstance)
}

func Fn_gtk_toolbar_get_nth_item(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_toolbar_get_nth_item(cValueInstance, cValue0)
}

func Fn_gtk_toolbar_get_relief_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	C.gtk_toolbar_get_relief_style(cValueInstance)
}

func Fn_gtk_toolbar_get_show_arrow(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	C.gtk_toolbar_get_show_arrow(cValueInstance)
}

func Fn_gtk_toolbar_get_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	C.gtk_toolbar_get_style(cValueInstance)
}

func Fn_gtk_toolbar_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_toolbar_insert(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_toolbar_set_drop_highlight_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_toolbar_set_drop_highlight_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_toolbar_set_icon_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkIconSize)(param0)

	C.gtk_toolbar_set_icon_size(cValueInstance, cValue0)
}

func Fn_gtk_toolbar_set_show_arrow(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toolbar_set_show_arrow(cValueInstance, cValue0)
}

func Fn_gtk_toolbar_set_style(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkToolbarStyle)(param0)

	C.gtk_toolbar_set_style(cValueInstance, cValue0)
}

func Fn_gtk_toolbar_unset_icon_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	C.gtk_toolbar_unset_icon_size(cValueInstance)
}

func Fn_gtk_toolbar_unset_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	C.gtk_toolbar_unset_style(cValueInstance)
}

func Fn_gtk_tooltip_set_custom(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_tooltip_set_custom(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_tooltip_set_icon(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_set_icon_from_icon_name(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_tooltip_set_icon_from_icon_name(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tooltip_set_icon_from_stock(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_tooltip_set_icon_from_stock(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tooltip_set_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tooltip_set_markup(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_set_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tooltip_set_text(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_set_tip_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_tooltip_set_tip_area(cValueInstance, cValue0)
}

func Fn_gtk_tooltip_trigger_tooltip_query(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	C.gtk_tooltip_trigger_tooltip_query(cValue0)
}

func Fn_gtk_toplevel_accessible_get_children(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToplevelAccessible)(unsafe.Pointer(paramInstance))

	C.gtk_toplevel_accessible_get_children(cValueInstance)
}

func Fn_gtk_tree_model_filter_clear_cache(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_filter_clear_cache(cValueInstance)
}

func Fn_gtk_tree_model_filter_convert_child_iter_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_filter_convert_child_iter_to_iter(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_model_filter_convert_child_path_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_model_filter_convert_child_path_to_path(cValueInstance, cValue0)
}

func Fn_gtk_tree_model_filter_convert_iter_to_child_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_filter_convert_iter_to_child_iter(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_model_filter_convert_path_to_child_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_model_filter_convert_path_to_child_path(cValueInstance, cValue0)
}

func Fn_gtk_tree_model_filter_get_model(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_filter_get_model(cValueInstance)
}

func Fn_gtk_tree_model_filter_refilter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_filter_refilter(cValueInstance)
}

// UNSUPPORTED : set_modify_func : has callback

func Fn_gtk_tree_model_filter_set_visible_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_model_filter_set_visible_column(cValueInstance, cValue0)
}

// UNSUPPORTED : set_visible_func : has callback

func Fn_gtk_tree_model_sort_clear_cache(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_sort_clear_cache(cValueInstance)
}

func Fn_gtk_tree_model_sort_convert_child_iter_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_sort_convert_child_iter_to_iter(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_model_sort_convert_child_path_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_model_sort_convert_child_path_to_path(cValueInstance, cValue0)
}

func Fn_gtk_tree_model_sort_convert_iter_to_child_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_sort_convert_iter_to_child_iter(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_model_sort_convert_path_to_child_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_model_sort_convert_path_to_child_path(cValueInstance, cValue0)
}

func Fn_gtk_tree_model_sort_get_model(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_sort_get_model(cValueInstance)
}

func Fn_gtk_tree_model_sort_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_tree_model_sort_iter_is_valid(cValueInstance, cValue0)
}

func Fn_gtk_tree_model_sort_reset_default_sort_func(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_sort_reset_default_sort_func(cValueInstance)
}

func Fn_gtk_tree_selection_count_selected_rows(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	C.gtk_tree_selection_count_selected_rows(cValueInstance)
}

func Fn_gtk_tree_selection_get_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	C.gtk_tree_selection_get_mode(cValueInstance)
}

func Fn_gtk_tree_selection_get_select_function(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	C.gtk_tree_selection_get_select_function(cValueInstance)
}

func Fn_gtk_tree_selection_get_selected(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreeModel)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_selection_get_selected(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_selection_get_selected_rows(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_tree_selection_get_selected_rows(cValueInstance, cValue0)
}

func Fn_gtk_tree_selection_get_tree_view(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	C.gtk_tree_selection_get_tree_view(cValueInstance)
}

func Fn_gtk_tree_selection_get_user_data(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	C.gtk_tree_selection_get_user_data(cValueInstance)
}

func Fn_gtk_tree_selection_iter_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_tree_selection_iter_is_selected(cValueInstance, cValue0)
}

func Fn_gtk_tree_selection_path_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_selection_path_is_selected(cValueInstance, cValue0)
}

func Fn_gtk_tree_selection_select_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	C.gtk_tree_selection_select_all(cValueInstance)
}

func Fn_gtk_tree_selection_select_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_tree_selection_select_iter(cValueInstance, cValue0)
}

func Fn_gtk_tree_selection_select_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_selection_select_path(cValueInstance, cValue0)
}

func Fn_gtk_tree_selection_select_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_tree_selection_select_range(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_tree_selection_set_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSelectionMode)(param0)

	C.gtk_tree_selection_set_mode(cValueInstance, cValue0)
}

// UNSUPPORTED : set_select_function : has callback

func Fn_gtk_tree_selection_unselect_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	C.gtk_tree_selection_unselect_all(cValueInstance)
}

func Fn_gtk_tree_selection_unselect_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_tree_selection_unselect_iter(cValueInstance, cValue0)
}

func Fn_gtk_tree_selection_unselect_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_selection_unselect_path(cValueInstance, cValue0)
}

func Fn_gtk_tree_selection_unselect_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_tree_selection_unselect_range(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : new : has varargs

func Fn_gtk_tree_store_newv(param0 int, param1 []uint64) {
	// has non-string array param
}

func Fn_gtk_tree_store_append(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_append(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_store_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	C.gtk_tree_store_clear(cValueInstance)
}

func Fn_gtk_tree_store_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	C.gtk_tree_store_insert(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_store_insert_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTreeIter)(unsafe.Pointer(param2))

	C.gtk_tree_store_insert_after(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_store_insert_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTreeIter)(unsafe.Pointer(param2))

	C.gtk_tree_store_insert_before(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : insert_with_values : has varargs

func Fn_gtk_tree_store_insert_with_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 []int, param4 []gobject.Value, param5 int) {
	// has non-string array param
}

func Fn_gtk_tree_store_is_ancestor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_is_ancestor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_store_iter_depth(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_tree_store_iter_depth(cValueInstance, cValue0)
}

func Fn_gtk_tree_store_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_tree_store_iter_is_valid(cValueInstance, cValue0)
}

func Fn_gtk_tree_store_move_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_move_after(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_store_move_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_move_before(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_store_prepend(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_prepend(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_store_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_tree_store_remove(cValueInstance, cValue0)
}

func Fn_gtk_tree_store_reorder(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []int) {
	// has non-string array param
}

// UNSUPPORTED : set : has varargs

func Fn_gtk_tree_store_set_column_types(paramInstance unsafe.Pointer, param0 int, param1 []uint64) {
	// has non-string array param
}

// UNSUPPORTED : set_valist : has va_list

func Fn_gtk_tree_store_set_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_tree_store_set_value(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_store_set_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []int, param2 []gobject.Value, param3 int) {
	// has non-string array param
}

func Fn_gtk_tree_store_swap(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_swap(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_new() {
	C.gtk_tree_view_new()
}

func Fn_gtk_tree_view_new_with_model(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_tree_view_new_with_model(cValue0)
}

func Fn_gtk_tree_view_append_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param0))

	C.gtk_tree_view_append_column(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_collapse_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_collapse_all(cValueInstance)
}

func Fn_gtk_tree_view_collapse_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_view_collapse_row(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_columns_autosize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_columns_autosize(cValueInstance)
}

func Fn_gtk_tree_view_convert_bin_window_to_tree_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_bin_window_to_tree_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_bin_window_to_widget_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_bin_window_to_widget_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_tree_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_tree_to_bin_window_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_tree_to_widget_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_tree_to_widget_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_widget_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_widget_to_bin_window_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_convert_widget_to_tree_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_tree_view_convert_widget_to_tree_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_create_row_drag_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_view_create_row_drag_icon(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_enable_model_drag_dest(paramInstance unsafe.Pointer, param0 []TargetEntry, param1 int, param2 int) {
	// has non-string array param
}

func Fn_gtk_tree_view_enable_model_drag_source(paramInstance unsafe.Pointer, param0 int, param1 []TargetEntry, param2 int, param3 int) {
	// has non-string array param
}

func Fn_gtk_tree_view_expand_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_expand_all(cValueInstance)
}

func Fn_gtk_tree_view_expand_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_tree_view_expand_row(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_expand_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_view_expand_to_path(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_get_background_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	C.gtk_tree_view_get_background_area(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_view_get_bin_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_bin_window(cValueInstance)
}

func Fn_gtk_tree_view_get_cell_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	C.gtk_tree_view_get_cell_area(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_view_get_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_get_column(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_get_columns(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_columns(cValueInstance)
}

func Fn_gtk_tree_view_get_cursor(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	C.gtk_tree_view_get_cursor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_get_dest_row_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkTreeViewDropPosition)(unsafe.Pointer(param3))

	C.gtk_tree_view_get_dest_row_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_get_drag_dest_row(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewDropPosition)(unsafe.Pointer(param1))

	C.gtk_tree_view_get_drag_dest_row(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_get_enable_search(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_enable_search(cValueInstance)
}

func Fn_gtk_tree_view_get_enable_tree_lines(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_enable_tree_lines(cValueInstance)
}

func Fn_gtk_tree_view_get_expander_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_expander_column(cValueInstance)
}

func Fn_gtk_tree_view_get_fixed_height_mode(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_fixed_height_mode(cValueInstance)
}

func Fn_gtk_tree_view_get_grid_lines(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_grid_lines(cValueInstance)
}

func Fn_gtk_tree_view_get_hadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_hadjustment(cValueInstance)
}

func Fn_gtk_tree_view_get_headers_clickable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_headers_clickable(cValueInstance)
}

func Fn_gtk_tree_view_get_headers_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_headers_visible(cValueInstance)
}

func Fn_gtk_tree_view_get_hover_expand(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_hover_expand(cValueInstance)
}

func Fn_gtk_tree_view_get_hover_selection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_hover_selection(cValueInstance)
}

func Fn_gtk_tree_view_get_level_indentation(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_level_indentation(cValueInstance)
}

func Fn_gtk_tree_view_get_model(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_model(cValueInstance)
}

func Fn_gtk_tree_view_get_path_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *int, param5 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (**C.GtkTreeViewColumn)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (*C.gint)(unsafe.Pointer(param5))

	C.gtk_tree_view_get_path_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_tree_view_get_reorderable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_reorderable(cValueInstance)
}

func Fn_gtk_tree_view_get_row_separator_func(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_row_separator_func(cValueInstance)
}

func Fn_gtk_tree_view_get_rubber_banding(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_rubber_banding(cValueInstance)
}

func Fn_gtk_tree_view_get_rules_hint(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_rules_hint(cValueInstance)
}

func Fn_gtk_tree_view_get_search_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_search_column(cValueInstance)
}

func Fn_gtk_tree_view_get_search_entry(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_search_entry(cValueInstance)
}

func Fn_gtk_tree_view_get_search_equal_func(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_search_equal_func(cValueInstance)
}

func Fn_gtk_tree_view_get_search_position_func(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_search_position_func(cValueInstance)
}

func Fn_gtk_tree_view_get_selection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_selection(cValueInstance)
}

func Fn_gtk_tree_view_get_show_expanders(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_show_expanders(cValueInstance)
}

func Fn_gtk_tree_view_get_tooltip_column(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_tooltip_column(cValueInstance)
}

func Fn_gtk_tree_view_get_tooltip_context(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := (**C.GtkTreeModel)(unsafe.Pointer(param3))

	cValue4 := (**C.GtkTreePath)(unsafe.Pointer(param4))

	cValue5 := (*C.GtkTreeIter)(unsafe.Pointer(param5))

	C.gtk_tree_view_get_tooltip_context(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_tree_view_get_vadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_get_vadjustment(cValueInstance)
}

func Fn_gtk_tree_view_get_visible_range(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_tree_view_get_visible_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_get_visible_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_tree_view_get_visible_rect(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_insert_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_tree_view_insert_column(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : insert_column_with_attributes : has varargs

// UNSUPPORTED : insert_column_with_data_func : has callback

func Fn_gtk_tree_view_is_rubber_banding_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_is_rubber_banding_active(cValueInstance)
}

// UNSUPPORTED : map_expanded_rows : has callback

func Fn_gtk_tree_view_move_column_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	C.gtk_tree_view_move_column_after(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_remove_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param0))

	C.gtk_tree_view_remove_column(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_row_activated(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	C.gtk_tree_view_row_activated(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_row_expanded(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_view_row_expanded(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_scroll_to_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 float32, param4 float32) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := (C.gfloat)(param3)

	cValue4 := (C.gfloat)(param4)

	C.gtk_tree_view_scroll_to_cell(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_tree_view_scroll_to_point(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_tree_view_scroll_to_point(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : set_column_drag_function : has callback

func Fn_gtk_tree_view_set_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_tree_view_set_cursor(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_view_set_cursor_on_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkCellRenderer)(unsafe.Pointer(param2))

	cValue3 := toCBool(param3)

	C.gtk_tree_view_set_cursor_on_cell(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

// UNSUPPORTED : set_destroy_count_func : has callback

func Fn_gtk_tree_view_set_drag_dest_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (C.GtkTreeViewDropPosition)(param1)

	C.gtk_tree_view_set_drag_dest_row(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_set_enable_search(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_enable_search(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_enable_tree_lines(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_enable_tree_lines(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_expander_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param0))

	C.gtk_tree_view_set_expander_column(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_fixed_height_mode(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_fixed_height_mode(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_grid_lines(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTreeViewGridLines)(param0)

	C.gtk_tree_view_set_grid_lines(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_tree_view_set_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_headers_clickable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_headers_clickable(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_headers_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_headers_visible(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_hover_expand(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_hover_expand(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_hover_selection(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_hover_selection(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_level_indentation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_set_level_indentation(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_tree_view_set_model(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_reorderable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_reorderable(cValueInstance, cValue0)
}

// UNSUPPORTED : set_row_separator_func : has callback

func Fn_gtk_tree_view_set_rubber_banding(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_rubber_banding(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_rules_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_rules_hint(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_search_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_set_search_column(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_search_entry(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkEntry)(unsafe.Pointer(param0))

	C.gtk_tree_view_set_search_entry(cValueInstance, cValue0)
}

// UNSUPPORTED : set_search_equal_func : has callback

// UNSUPPORTED : set_search_position_func : has callback

func Fn_gtk_tree_view_set_show_expanders(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_set_show_expanders(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_tooltip_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTooltip)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkCellRenderer)(unsafe.Pointer(param3))

	C.gtk_tree_view_set_tooltip_cell(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_set_tooltip_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_set_tooltip_column(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_set_tooltip_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTooltip)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	C.gtk_tree_view_set_tooltip_row(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_tree_view_set_vadjustment(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_unset_rows_drag_dest(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_unset_rows_drag_dest(cValueInstance)
}

func Fn_gtk_tree_view_unset_rows_drag_source(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_unset_rows_drag_source(cValueInstance)
}

func Fn_gtk_tree_view_column_new() {
	C.gtk_tree_view_column_new()
}

// UNSUPPORTED : new_with_attributes : has varargs

func Fn_gtk_tree_view_column_add_attribute(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	C.gtk_tree_view_column_add_attribute(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_view_column_cell_get_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_tree_view_column_cell_get_position(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_view_column_cell_get_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int, param3 *int, param4 *int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	C.gtk_tree_view_column_cell_get_size(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_tree_view_column_cell_is_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_cell_is_visible(cValueInstance)
}

func Fn_gtk_tree_view_column_cell_set_cell_data(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := toCBool(param3)

	C.gtk_tree_view_column_cell_set_cell_data(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tree_view_column_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_clear(cValueInstance)
}

func Fn_gtk_tree_view_column_clear_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_tree_view_column_clear_attributes(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_clicked(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_clicked(cValueInstance)
}

func Fn_gtk_tree_view_column_focus_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_tree_view_column_focus_cell(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_get_alignment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_alignment(cValueInstance)
}

func Fn_gtk_tree_view_column_get_clickable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_clickable(cValueInstance)
}

func Fn_gtk_tree_view_column_get_expand(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_expand(cValueInstance)
}

func Fn_gtk_tree_view_column_get_fixed_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_fixed_width(cValueInstance)
}

func Fn_gtk_tree_view_column_get_max_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_max_width(cValueInstance)
}

func Fn_gtk_tree_view_column_get_min_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_min_width(cValueInstance)
}

func Fn_gtk_tree_view_column_get_reorderable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_reorderable(cValueInstance)
}

func Fn_gtk_tree_view_column_get_resizable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_resizable(cValueInstance)
}

func Fn_gtk_tree_view_column_get_sizing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_sizing(cValueInstance)
}

func Fn_gtk_tree_view_column_get_sort_column_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_sort_column_id(cValueInstance)
}

func Fn_gtk_tree_view_column_get_sort_indicator(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_sort_indicator(cValueInstance)
}

func Fn_gtk_tree_view_column_get_sort_order(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_sort_order(cValueInstance)
}

func Fn_gtk_tree_view_column_get_spacing(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_spacing(cValueInstance)
}

func Fn_gtk_tree_view_column_get_title(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_title(cValueInstance)
}

func Fn_gtk_tree_view_column_get_tree_view(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_tree_view(cValueInstance)
}

func Fn_gtk_tree_view_column_get_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_visible(cValueInstance)
}

func Fn_gtk_tree_view_column_get_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_widget(cValueInstance)
}

func Fn_gtk_tree_view_column_get_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_get_width(cValueInstance)
}

func Fn_gtk_tree_view_column_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_tree_view_column_pack_end(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_column_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_tree_view_column_pack_start(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_column_queue_resize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_column_queue_resize(cValueInstance)
}

func Fn_gtk_tree_view_column_set_alignment(paramInstance unsafe.Pointer, param0 float32) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	C.gtk_tree_view_column_set_alignment(cValueInstance, cValue0)
}

// UNSUPPORTED : set_attributes : has varargs

// UNSUPPORTED : set_cell_data_func : has callback

func Fn_gtk_tree_view_column_set_clickable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_column_set_clickable(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_expand(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_column_set_expand(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_fixed_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_column_set_fixed_width(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_max_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_column_set_max_width(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_min_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_column_set_min_width(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_reorderable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_column_set_reorderable(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_resizable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_column_set_resizable(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_sizing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTreeViewColumnSizing)(param0)

	C.gtk_tree_view_column_set_sizing(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_sort_column_id(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_column_set_sort_column_id(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_sort_indicator(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_column_set_sort_indicator(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_sort_order(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSortType)(param0)

	C.gtk_tree_view_column_set_sort_order(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_view_column_set_spacing(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tree_view_column_set_title(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tree_view_column_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_column_set_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_tree_view_column_set_widget(cValueInstance, cValue0)
}

func Fn_gtk_ui_manager_new() {
	C.gtk_ui_manager_new()
}

func Fn_gtk_ui_manager_add_ui(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 string, param3 string, param4 int, param5 bool) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (C.GtkUIManagerItemType)(param4)

	cValue5 := toCBool(param5)

	C.gtk_ui_manager_add_ui(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_ui_manager_add_ui_from_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_ui_manager_add_ui_from_file(cValueInstance, cValue0, cError)
}

func Fn_gtk_ui_manager_add_ui_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64, error unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gssize)(param1)

	cError := (**C.GError)(error)

	C.gtk_ui_manager_add_ui_from_string(cValueInstance, cValue0, cValue1, cError)
}

func Fn_gtk_ui_manager_ensure_update(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	C.gtk_ui_manager_ensure_update(cValueInstance)
}

func Fn_gtk_ui_manager_get_accel_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	C.gtk_ui_manager_get_accel_group(cValueInstance)
}

func Fn_gtk_ui_manager_get_action(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_ui_manager_get_action(cValueInstance, cValue0)
}

func Fn_gtk_ui_manager_get_action_groups(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	C.gtk_ui_manager_get_action_groups(cValueInstance)
}

func Fn_gtk_ui_manager_get_add_tearoffs(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	C.gtk_ui_manager_get_add_tearoffs(cValueInstance)
}

func Fn_gtk_ui_manager_get_toplevels(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUIManagerItemType)(param0)

	C.gtk_ui_manager_get_toplevels(cValueInstance, cValue0)
}

func Fn_gtk_ui_manager_get_ui(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	C.gtk_ui_manager_get_ui(cValueInstance)
}

func Fn_gtk_ui_manager_get_widget(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_ui_manager_get_widget(cValueInstance, cValue0)
}

func Fn_gtk_ui_manager_insert_action_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkActionGroup)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_ui_manager_insert_action_group(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_ui_manager_new_merge_id(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	C.gtk_ui_manager_new_merge_id(cValueInstance)
}

func Fn_gtk_ui_manager_remove_action_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkActionGroup)(unsafe.Pointer(param0))

	C.gtk_ui_manager_remove_action_group(cValueInstance, cValue0)
}

func Fn_gtk_ui_manager_remove_ui(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_ui_manager_remove_ui(cValueInstance, cValue0)
}

func Fn_gtk_ui_manager_set_add_tearoffs(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_ui_manager_set_add_tearoffs(cValueInstance, cValue0)
}

func Fn_gtk_vbox_new(param0 bool, param1 int) {
	cValue0 := toCBool(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_vbox_new(cValue0, cValue1)
}

func Fn_gtk_vbutton_box_new() {
	C.gtk_vbutton_box_new()
}

func Fn_gtk_vpaned_new() {
	C.gtk_vpaned_new()
}

func Fn_gtk_vscale_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_vscale_new(cValue0)
}

func Fn_gtk_vscale_new_with_range(param0 float64, param1 float64, param2 float64) {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	C.gtk_vscale_new_with_range(cValue0, cValue1, cValue2)
}

func Fn_gtk_vscrollbar_new(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_vscrollbar_new(cValue0)
}

func Fn_gtk_vseparator_new() {
	C.gtk_vseparator_new()
}

func Fn_gtk_viewport_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	C.gtk_viewport_new(cValue0, cValue1)
}

func Fn_gtk_viewport_get_hadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	C.gtk_viewport_get_hadjustment(cValueInstance)
}

func Fn_gtk_viewport_get_shadow_type(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	C.gtk_viewport_get_shadow_type(cValueInstance)
}

func Fn_gtk_viewport_get_vadjustment(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	C.gtk_viewport_get_vadjustment(cValueInstance)
}

func Fn_gtk_viewport_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_viewport_set_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_viewport_set_shadow_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkShadowType)(param0)

	C.gtk_viewport_set_shadow_type(cValueInstance, cValue0)
}

func Fn_gtk_viewport_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_viewport_set_vadjustment(cValueInstance, cValue0)
}

func Fn_gtk_volume_button_new() {
	C.gtk_volume_button_new()
}

// UNSUPPORTED : new : has varargs

func Fn_gtk_widget_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_activate(cValueInstance)
}

func Fn_gtk_widget_add_accelerator(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 uint, param3 int, param4 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkAccelGroup)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (C.GdkModifierType)(param3)

	cValue4 := (C.GtkAccelFlags)(param4)

	C.gtk_widget_add_accelerator(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_widget_add_events(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_add_events(cValueInstance, cValue0)
}

func Fn_gtk_widget_add_mnemonic_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_widget_add_mnemonic_label(cValueInstance, cValue0)
}

// UNSUPPORTED : add_tick_callback : has callback

func Fn_gtk_widget_can_activate_accel(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_widget_can_activate_accel(cValueInstance, cValue0)
}

func Fn_gtk_widget_child_focus(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkDirectionType)(param0)

	C.gtk_widget_child_focus(cValueInstance, cValue0)
}

func Fn_gtk_widget_child_notify(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_child_notify(cValueInstance, cValue0)
}

func Fn_gtk_widget_class_path(paramInstance unsafe.Pointer, param0 *uint, param1 *string, param2 *string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	var cValue2String *C.gchar
	cValue2 := &cValue2String

	C.gtk_widget_class_path(cValueInstance, cValue0, cValue1, cValue2)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	param2String := C.GoString(cValue2String)
	*param2 = param2String
}

func Fn_gtk_widget_compute_expand(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkOrientation)(param0)

	C.gtk_widget_compute_expand(cValueInstance, cValue0)
}

func Fn_gtk_widget_create_pango_context(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_create_pango_context(cValueInstance)
}

func Fn_gtk_widget_create_pango_layout(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_create_pango_layout(cValueInstance, cValue0)
}

func Fn_gtk_widget_destroy(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_destroy(cValueInstance)
}

func Fn_gtk_widget_destroyed(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_widget_destroyed(cValueInstance, cValue0)
}

func Fn_gtk_drag_begin(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTargetList)(unsafe.Pointer(param0))

	cValue1 := (C.GdkDragAction)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.GdkEvent)(unsafe.Pointer(param3))

	C.gtk_drag_begin(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_drag_check_threshold(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_drag_check_threshold(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_drag_dest_add_image_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_dest_add_image_targets(cValueInstance)
}

func Fn_gtk_drag_dest_add_text_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_dest_add_text_targets(cValueInstance)
}

func Fn_gtk_drag_dest_add_uri_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_dest_add_uri_targets(cValueInstance)
}

func Fn_gtk_drag_dest_find_target(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTargetList)(unsafe.Pointer(param1))

	C.gtk_drag_dest_find_target(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_drag_dest_get_target_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_dest_get_target_list(cValueInstance)
}

func Fn_gtk_drag_dest_get_track_motion(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_dest_get_track_motion(cValueInstance)
}

func Fn_gtk_drag_dest_set(paramInstance unsafe.Pointer, param0 int, param1 []TargetEntry, param2 int, param3 int) {
	// has non-string array param
}

func Fn_gtk_drag_dest_set_proxy(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GdkDragProtocol)(param1)

	cValue2 := toCBool(param2)

	C.gtk_drag_dest_set_proxy(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_drag_dest_set_target_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTargetList)(unsafe.Pointer(param0))

	C.gtk_drag_dest_set_target_list(cValueInstance, cValue0)
}

func Fn_gtk_drag_dest_set_track_motion(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_drag_dest_set_track_motion(cValueInstance, cValue0)
}

func Fn_gtk_drag_dest_unset(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_dest_unset(cValueInstance)
}

func Fn_gtk_drag_get_data(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 gdk.Atom, param2 uint32) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	cValue2 := (C.guint32)(param2)

	C.gtk_drag_get_data(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_drag_highlight(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_highlight(cValueInstance)
}

func Fn_gtk_drag_source_add_image_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_source_add_image_targets(cValueInstance)
}

func Fn_gtk_drag_source_add_text_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_source_add_text_targets(cValueInstance)
}

func Fn_gtk_drag_source_add_uri_targets(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_source_add_uri_targets(cValueInstance)
}

func Fn_gtk_drag_source_get_target_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_source_get_target_list(cValueInstance)
}

func Fn_gtk_drag_source_set(paramInstance unsafe.Pointer, param0 int, param1 []TargetEntry, param2 int, param3 int) {
	// has non-string array param
}

func Fn_gtk_drag_source_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_drag_source_set_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_drag_source_set_icon_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_drag_source_set_icon_pixbuf(cValueInstance, cValue0)
}

func Fn_gtk_drag_source_set_icon_stock(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_drag_source_set_icon_stock(cValueInstance, cValue0)
}

func Fn_gtk_drag_source_set_target_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTargetList)(unsafe.Pointer(param0))

	C.gtk_drag_source_set_target_list(cValueInstance, cValue0)
}

func Fn_gtk_drag_source_unset(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_source_unset(cValueInstance)
}

func Fn_gtk_drag_unhighlight(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_drag_unhighlight(cValueInstance)
}

func Fn_gtk_widget_ensure_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_ensure_style(cValueInstance)
}

func Fn_gtk_widget_error_bell(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_error_bell(cValueInstance)
}

func Fn_gtk_widget_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	C.gtk_widget_event(cValueInstance, cValue0)
}

func Fn_gtk_widget_freeze_child_notify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_freeze_child_notify(cValueInstance)
}

func Fn_gtk_widget_get_accessible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_accessible(cValueInstance)
}

func Fn_gtk_widget_get_allocated_height(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_allocated_height(cValueInstance)
}

func Fn_gtk_widget_get_allocated_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_allocated_width(cValueInstance)
}

func Fn_gtk_widget_get_allocation(paramInstance unsafe.Pointer, param0 *gdk.Rectangle) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	C.gtk_widget_get_allocation(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_ancestor(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	C.gtk_widget_get_ancestor(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_app_paintable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_app_paintable(cValueInstance)
}

func Fn_gtk_widget_get_can_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_can_default(cValueInstance)
}

func Fn_gtk_widget_get_can_focus(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_can_focus(cValueInstance)
}

func Fn_gtk_widget_get_child_requisition(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRequisition)(unsafe.Pointer(param0))

	C.gtk_widget_get_child_requisition(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_child_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_child_visible(cValueInstance)
}

func Fn_gtk_widget_get_clipboard(paramInstance unsafe.Pointer, param0 gdk.Atom) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	C.gtk_widget_get_clipboard(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_composite_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_composite_name(cValueInstance)
}

func Fn_gtk_widget_get_direction(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_direction(cValueInstance)
}

func Fn_gtk_widget_get_display(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_display(cValueInstance)
}

func Fn_gtk_widget_get_double_buffered(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_double_buffered(cValueInstance)
}

func Fn_gtk_widget_get_events(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_events(cValueInstance)
}

func Fn_gtk_widget_get_halign(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_halign(cValueInstance)
}

func Fn_gtk_widget_get_has_tooltip(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_has_tooltip(cValueInstance)
}

func Fn_gtk_widget_get_has_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_has_window(cValueInstance)
}

func Fn_gtk_widget_get_hexpand(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_hexpand(cValueInstance)
}

func Fn_gtk_widget_get_hexpand_set(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_hexpand_set(cValueInstance)
}

func Fn_gtk_widget_get_modifier_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_modifier_style(cValueInstance)
}

func Fn_gtk_widget_get_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_name(cValueInstance)
}

func Fn_gtk_widget_get_no_show_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_no_show_all(cValueInstance)
}

func Fn_gtk_widget_get_pango_context(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_pango_context(cValueInstance)
}

func Fn_gtk_widget_get_parent(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_parent(cValueInstance)
}

func Fn_gtk_widget_get_parent_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_parent_window(cValueInstance)
}

func Fn_gtk_widget_get_path(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_path(cValueInstance)
}

func Fn_gtk_widget_get_pointer(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_widget_get_pointer(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_receives_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_receives_default(cValueInstance)
}

func Fn_gtk_widget_get_root_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_root_window(cValueInstance)
}

func Fn_gtk_widget_get_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_screen(cValueInstance)
}

func Fn_gtk_widget_get_sensitive(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_sensitive(cValueInstance)
}

func Fn_gtk_widget_get_settings(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_settings(cValueInstance)
}

func Fn_gtk_widget_get_size_request(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_widget_get_size_request(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_state(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_state(cValueInstance)
}

func Fn_gtk_widget_get_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_style(cValueInstance)
}

func Fn_gtk_widget_get_style_context(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_style_context(cValueInstance)
}

func Fn_gtk_widget_get_support_multidevice(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_support_multidevice(cValueInstance)
}

func Fn_gtk_widget_get_template_child(paramInstance unsafe.Pointer, param0 uint64, param1 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_get_template_child(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_tooltip_markup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_tooltip_markup(cValueInstance)
}

func Fn_gtk_widget_get_tooltip_text(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_tooltip_text(cValueInstance)
}

func Fn_gtk_widget_get_tooltip_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_tooltip_window(cValueInstance)
}

func Fn_gtk_widget_get_toplevel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_toplevel(cValueInstance)
}

func Fn_gtk_widget_get_valign(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_valign(cValueInstance)
}

func Fn_gtk_widget_get_vexpand(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_vexpand(cValueInstance)
}

func Fn_gtk_widget_get_vexpand_set(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_vexpand_set(cValueInstance)
}

func Fn_gtk_widget_get_visible(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_visible(cValueInstance)
}

func Fn_gtk_widget_get_visual(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_visual(cValueInstance)
}

func Fn_gtk_widget_get_window(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_get_window(cValueInstance)
}

func Fn_gtk_grab_add(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_grab_add(cValueInstance)
}

func Fn_gtk_widget_grab_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_grab_default(cValueInstance)
}

func Fn_gtk_widget_grab_focus(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_grab_focus(cValueInstance)
}

func Fn_gtk_grab_remove(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_grab_remove(cValueInstance)
}

func Fn_gtk_widget_has_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_has_default(cValueInstance)
}

func Fn_gtk_widget_has_focus(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_has_focus(cValueInstance)
}

func Fn_gtk_widget_has_grab(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_has_grab(cValueInstance)
}

func Fn_gtk_widget_has_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_has_screen(cValueInstance)
}

func Fn_gtk_widget_hide(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_hide(cValueInstance)
}

func Fn_gtk_widget_hide_on_delete(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_hide_on_delete(cValueInstance)
}

func Fn_gtk_widget_in_destruction(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_in_destruction(cValueInstance)
}

func Fn_gtk_widget_intersect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gtk_widget_intersect(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_is_ancestor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_widget_is_ancestor(cValueInstance, cValue0)
}

func Fn_gtk_widget_is_composited(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_is_composited(cValueInstance)
}

func Fn_gtk_widget_is_drawable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_is_drawable(cValueInstance)
}

func Fn_gtk_widget_is_focus(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_is_focus(cValueInstance)
}

func Fn_gtk_widget_is_sensitive(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_is_sensitive(cValueInstance)
}

func Fn_gtk_widget_is_toplevel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_is_toplevel(cValueInstance)
}

func Fn_gtk_widget_keynav_failed(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkDirectionType)(param0)

	C.gtk_widget_keynav_failed(cValueInstance, cValue0)
}

func Fn_gtk_widget_list_accel_closures(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_list_accel_closures(cValueInstance)
}

func Fn_gtk_widget_list_mnemonic_labels(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_list_mnemonic_labels(cValueInstance)
}

func Fn_gtk_widget_map(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_map(cValueInstance)
}

func Fn_gtk_widget_mnemonic_activate(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_mnemonic_activate(cValueInstance, cValue0)
}

func Fn_gtk_widget_modify_base(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gtk_widget_modify_base(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_modify_bg(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gtk_widget_modify_bg(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_modify_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gtk_widget_modify_cursor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_modify_fg(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gtk_widget_modify_fg(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_modify_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	C.gtk_widget_modify_font(cValueInstance, cValue0)
}

func Fn_gtk_widget_modify_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRcStyle)(unsafe.Pointer(param0))

	C.gtk_widget_modify_style(cValueInstance, cValue0)
}

func Fn_gtk_widget_modify_text(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	C.gtk_widget_modify_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path(paramInstance unsafe.Pointer, param0 *uint, param1 *string, param2 *string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	var cValue2String *C.gchar
	cValue2 := &cValue2String

	C.gtk_widget_path(cValueInstance, cValue0, cValue1, cValue2)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	param2String := C.GoString(cValue2String)
	*param2 = param2String
}

func Fn_gtk_widget_queue_compute_expand(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_queue_compute_expand(cValueInstance)
}

func Fn_gtk_widget_queue_draw(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_queue_draw(cValueInstance)
}

func Fn_gtk_widget_queue_draw_area(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	C.gtk_widget_queue_draw_area(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_widget_queue_resize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_queue_resize(cValueInstance)
}

func Fn_gtk_widget_queue_resize_no_redraw(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_queue_resize_no_redraw(cValueInstance)
}

func Fn_gtk_widget_realize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_realize(cValueInstance)
}

func Fn_gtk_widget_region_intersect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	C.gtk_widget_region_intersect(cValueInstance, cValue0)
}

func Fn_gtk_widget_remove_accelerator(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	C.gtk_widget_remove_accelerator(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_widget_remove_mnemonic_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_widget_remove_mnemonic_label(cValueInstance, cValue0)
}

func Fn_gtk_widget_render_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_widget_render_icon(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_widget_reparent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_widget_reparent(cValueInstance, cValue0)
}

func Fn_gtk_widget_reset_rc_styles(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_reset_rc_styles(cValueInstance)
}

func Fn_gtk_widget_send_expose(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	C.gtk_widget_send_expose(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_accel_path(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkAccelGroup)(unsafe.Pointer(param1))

	C.gtk_widget_set_accel_path(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_set_allocation(paramInstance unsafe.Pointer, param0 *gdk.Rectangle) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	C.gtk_widget_set_allocation(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_app_paintable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_app_paintable(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_can_default(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_can_default(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_can_focus(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_can_focus(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_child_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_child_visible(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_composite_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_set_composite_name(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_direction(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextDirection)(param0)

	C.gtk_widget_set_direction(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_double_buffered(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_double_buffered(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_events(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_events(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_halign(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkAlign)(param0)

	C.gtk_widget_set_halign(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_has_tooltip(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_has_tooltip(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_has_window(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_has_window(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_hexpand(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_hexpand(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_hexpand_set(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_hexpand_set(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_set_name(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_no_show_all(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_no_show_all(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_widget_set_parent(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_parent_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gtk_widget_set_parent_window(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_receives_default(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_receives_default(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_redraw_on_allocate(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_redraw_on_allocate(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_sensitive(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_size_request(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_widget_set_size_request(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_set_state(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	C.gtk_widget_set_state(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	C.gtk_widget_set_style(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_set_tooltip_markup(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_tooltip_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_widget_set_tooltip_text(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_tooltip_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_widget_set_tooltip_window(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_valign(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkAlign)(param0)

	C.gtk_widget_set_valign(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_vexpand(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_vexpand(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_vexpand_set(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_vexpand_set(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_visual(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkVisual)(unsafe.Pointer(param0))

	C.gtk_widget_set_visual(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gtk_widget_set_window(cValueInstance, cValue0)
}

func Fn_gtk_widget_show(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_show(cValueInstance)
}

func Fn_gtk_widget_show_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_show_all(cValueInstance)
}

func Fn_gtk_widget_show_now(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_show_now(cValueInstance)
}

func Fn_gtk_widget_size_allocate(paramInstance unsafe.Pointer, param0 *gdk.Rectangle) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	C.gtk_widget_size_allocate(cValueInstance, cValue0)
}

func Fn_gtk_widget_size_request(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRequisition)(unsafe.Pointer(param0))

	C.gtk_widget_size_request(cValueInstance, cValue0)
}

// UNSUPPORTED : style_get : has varargs

func Fn_gtk_widget_style_get_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.gtk_widget_style_get_property(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : style_get_valist : has va_list

func Fn_gtk_widget_thaw_child_notify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_thaw_child_notify(cValueInstance)
}

func Fn_gtk_widget_translate_coordinates(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 *int, param4 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	C.gtk_widget_translate_coordinates(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_widget_trigger_tooltip_query(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_trigger_tooltip_query(cValueInstance)
}

func Fn_gtk_widget_unmap(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_unmap(cValueInstance)
}

func Fn_gtk_widget_unparent(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_unparent(cValueInstance)
}

func Fn_gtk_widget_unrealize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_unrealize(cValueInstance)
}

func Fn_gtk_widget_get_default_direction() {
	C.gtk_widget_get_default_direction()
}

func Fn_gtk_widget_get_default_style() {
	C.gtk_widget_get_default_style()
}

func Fn_gtk_widget_pop_composite_child() {
	C.gtk_widget_pop_composite_child()
}

func Fn_gtk_widget_push_composite_child() {
	C.gtk_widget_push_composite_child()
}

func Fn_gtk_widget_set_default_direction(param0 int) {
	cValue0 := (C.GtkTextDirection)(param0)

	C.gtk_widget_set_default_direction(cValue0)
}

func Fn_gtk_window_new(param0 int) {
	cValue0 := (C.GtkWindowType)(param0)

	C.gtk_window_new(cValue0)
}

func Fn_gtk_window_activate_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_activate_default(cValueInstance)
}

func Fn_gtk_window_activate_focus(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_activate_focus(cValueInstance)
}

func Fn_gtk_window_activate_key(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	C.gtk_window_activate_key(cValueInstance, cValue0)
}

func Fn_gtk_window_add_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	C.gtk_window_add_accel_group(cValueInstance, cValue0)
}

func Fn_gtk_window_add_mnemonic(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	C.gtk_window_add_mnemonic(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_begin_move_drag(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 uint32) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.guint32)(param3)

	C.gtk_window_begin_move_drag(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_window_begin_resize_drag(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 uint32) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkWindowEdge)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	cValue4 := (C.guint32)(param4)

	C.gtk_window_begin_resize_drag(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_window_deiconify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_deiconify(cValueInstance)
}

func Fn_gtk_window_fullscreen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_fullscreen(cValueInstance)
}

func Fn_gtk_window_get_accept_focus(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_accept_focus(cValueInstance)
}

func Fn_gtk_window_get_decorated(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_decorated(cValueInstance)
}

func Fn_gtk_window_get_default_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_window_get_default_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_get_default_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_default_widget(cValueInstance)
}

func Fn_gtk_window_get_deletable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_deletable(cValueInstance)
}

func Fn_gtk_window_get_destroy_with_parent(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_destroy_with_parent(cValueInstance)
}

func Fn_gtk_window_get_focus(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_focus(cValueInstance)
}

func Fn_gtk_window_get_focus_on_map(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_focus_on_map(cValueInstance)
}

func Fn_gtk_window_get_gravity(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_gravity(cValueInstance)
}

func Fn_gtk_window_get_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_group(cValueInstance)
}

func Fn_gtk_window_get_icon(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_icon(cValueInstance)
}

func Fn_gtk_window_get_icon_list(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_icon_list(cValueInstance)
}

func Fn_gtk_window_get_icon_name(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_icon_name(cValueInstance)
}

func Fn_gtk_window_get_mnemonic_modifier(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_mnemonic_modifier(cValueInstance)
}

func Fn_gtk_window_get_modal(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_modal(cValueInstance)
}

func Fn_gtk_window_get_opacity(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_opacity(cValueInstance)
}

func Fn_gtk_window_get_position(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_window_get_position(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_get_resizable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_resizable(cValueInstance)
}

func Fn_gtk_window_get_role(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_role(cValueInstance)
}

func Fn_gtk_window_get_screen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_screen(cValueInstance)
}

func Fn_gtk_window_get_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_window_get_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_get_skip_pager_hint(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_skip_pager_hint(cValueInstance)
}

func Fn_gtk_window_get_skip_taskbar_hint(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_skip_taskbar_hint(cValueInstance)
}

func Fn_gtk_window_get_title(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_title(cValueInstance)
}

func Fn_gtk_window_get_transient_for(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_transient_for(cValueInstance)
}

func Fn_gtk_window_get_type_hint(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_type_hint(cValueInstance)
}

func Fn_gtk_window_get_urgency_hint(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_get_urgency_hint(cValueInstance)
}

func Fn_gtk_window_has_group(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_has_group(cValueInstance)
}

func Fn_gtk_window_has_toplevel_focus(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_has_toplevel_focus(cValueInstance)
}

func Fn_gtk_window_iconify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_iconify(cValueInstance)
}

func Fn_gtk_window_is_active(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_is_active(cValueInstance)
}

func Fn_gtk_window_maximize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_maximize(cValueInstance)
}

func Fn_gtk_window_mnemonic_activate(paramInstance unsafe.Pointer, param0 uint, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	C.gtk_window_mnemonic_activate(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_move(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_window_move(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_parse_geometry(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_window_parse_geometry(cValueInstance, cValue0)
}

func Fn_gtk_window_present(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_present(cValueInstance)
}

func Fn_gtk_window_present_with_time(paramInstance unsafe.Pointer, param0 uint32) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint32)(param0)

	C.gtk_window_present_with_time(cValueInstance, cValue0)
}

func Fn_gtk_window_propagate_key_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	C.gtk_window_propagate_key_event(cValueInstance, cValue0)
}

func Fn_gtk_window_remove_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	C.gtk_window_remove_accel_group(cValueInstance, cValue0)
}

func Fn_gtk_window_remove_mnemonic(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	C.gtk_window_remove_mnemonic(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_reshow_with_initial_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_reshow_with_initial_size(cValueInstance)
}

func Fn_gtk_window_resize(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_window_resize(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_set_accept_focus(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_accept_focus(cValueInstance, cValue0)
}

func Fn_gtk_window_set_decorated(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_decorated(cValueInstance, cValue0)
}

func Fn_gtk_window_set_default(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_window_set_default(cValueInstance, cValue0)
}

func Fn_gtk_window_set_default_size(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_window_set_default_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_set_deletable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_deletable(cValueInstance, cValue0)
}

func Fn_gtk_window_set_destroy_with_parent(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_destroy_with_parent(cValueInstance, cValue0)
}

func Fn_gtk_window_set_focus(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_window_set_focus(cValueInstance, cValue0)
}

func Fn_gtk_window_set_focus_on_map(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_focus_on_map(cValueInstance, cValue0)
}

func Fn_gtk_window_set_geometry_hints(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkGeometry)(unsafe.Pointer(param1))

	cValue2 := (C.GdkWindowHints)(param2)

	C.gtk_window_set_geometry_hints(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_window_set_gravity(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkGravity)(param0)

	C.gtk_window_set_gravity(cValueInstance, cValue0)
}

func Fn_gtk_window_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_window_set_icon(cValueInstance, cValue0)
}

func Fn_gtk_window_set_icon_from_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_window_set_icon_from_file(cValueInstance, cValue0, cError)
}

func Fn_gtk_window_set_icon_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.gtk_window_set_icon_list(cValueInstance, cValue0)
}

func Fn_gtk_window_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_window_set_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_window_set_keep_above(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_keep_above(cValueInstance, cValue0)
}

func Fn_gtk_window_set_keep_below(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_keep_below(cValueInstance, cValue0)
}

func Fn_gtk_window_set_mnemonic_modifier(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkModifierType)(param0)

	C.gtk_window_set_mnemonic_modifier(cValueInstance, cValue0)
}

func Fn_gtk_window_set_modal(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_modal(cValueInstance, cValue0)
}

func Fn_gtk_window_set_opacity(paramInstance unsafe.Pointer, param0 float64) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	C.gtk_window_set_opacity(cValueInstance, cValue0)
}

func Fn_gtk_window_set_position(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkWindowPosition)(param0)

	C.gtk_window_set_position(cValueInstance, cValue0)
}

func Fn_gtk_window_set_resizable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_resizable(cValueInstance, cValue0)
}

func Fn_gtk_window_set_role(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_window_set_role(cValueInstance, cValue0)
}

func Fn_gtk_window_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_window_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_window_set_skip_pager_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_skip_pager_hint(cValueInstance, cValue0)
}

func Fn_gtk_window_set_skip_taskbar_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_skip_taskbar_hint(cValueInstance, cValue0)
}

func Fn_gtk_window_set_startup_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_window_set_startup_id(cValueInstance, cValue0)
}

func Fn_gtk_window_set_title(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_window_set_title(cValueInstance, cValue0)
}

func Fn_gtk_window_set_transient_for(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_window_set_transient_for(cValueInstance, cValue0)
}

func Fn_gtk_window_set_type_hint(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkWindowTypeHint)(param0)

	C.gtk_window_set_type_hint(cValueInstance, cValue0)
}

func Fn_gtk_window_set_urgency_hint(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_urgency_hint(cValueInstance, cValue0)
}

func Fn_gtk_window_set_wmclass(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_window_set_wmclass(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_stick(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_stick(cValueInstance)
}

func Fn_gtk_window_unfullscreen(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_unfullscreen(cValueInstance)
}

func Fn_gtk_window_unmaximize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_unmaximize(cValueInstance)
}

func Fn_gtk_window_unstick(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_unstick(cValueInstance)
}

func Fn_gtk_window_get_default_icon_list() {
	C.gtk_window_get_default_icon_list()
}

func Fn_gtk_window_get_default_icon_name() {
	C.gtk_window_get_default_icon_name()
}

func Fn_gtk_window_list_toplevels() {
	C.gtk_window_list_toplevels()
}

func Fn_gtk_window_set_auto_startup_notification(param0 bool) {
	cValue0 := toCBool(param0)

	C.gtk_window_set_auto_startup_notification(cValue0)
}

func Fn_gtk_window_set_default_icon(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_window_set_default_icon(cValue0)
}

func Fn_gtk_window_set_default_icon_from_file(param0 string, error unsafe.Pointer) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	C.gtk_window_set_default_icon_from_file(cValue0, cError)
}

func Fn_gtk_window_set_default_icon_list(param0 unsafe.Pointer) {
	cValue0 := (*C.GList)(unsafe.Pointer(param0))

	C.gtk_window_set_default_icon_list(cValue0)
}

func Fn_gtk_window_set_default_icon_name(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_window_set_default_icon_name(cValue0)
}

func Fn_gtk_window_group_new() {
	C.gtk_window_group_new()
}

func Fn_gtk_window_group_add_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_window_group_add_window(cValueInstance, cValue0)
}

func Fn_gtk_window_group_list_windows(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	C.gtk_window_group_list_windows(cValueInstance)
}

func Fn_gtk_window_group_remove_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_window_group_remove_window(cValueInstance, cValue0)
}