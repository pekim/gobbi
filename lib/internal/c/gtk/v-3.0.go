// Code generated - DO NOT EDIT.
// +build gtk_3.0

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
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
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
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

// UNSUPPORTED : gtk_binding_entry_add_signal : has varargs

func Fn_gtk_binding_entry_add_signal_from_string(param0 unsafe.Pointer, param1 string) int {
	cValue0 := (*C.GtkBindingSet)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_binding_entry_add_signal_from_string(cValue0, cValue1)

	return (int)(ret)
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

func Fn_gtk_binding_set_activate(paramInstance unsafe.Pointer, param0 uint, param1 int, param2 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkBindingSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	cValue2 := (*C.GObject)(unsafe.Pointer(param2))

	ret := C.gtk_binding_set_activate(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_binding_set_add_path(paramInstance unsafe.Pointer, param0 int, param1 string, param2 int) {
	cValueInstance := (*C.GtkBindingSet)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPathType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GtkPathPriorityType)(param2)

	C.gtk_binding_set_add_path(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_binding_set_by_class(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.gpointer)(param0)

	ret := C.gtk_binding_set_by_class(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_binding_set_find(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_binding_set_find(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_binding_set_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_binding_set_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_border_new() unsafe.Pointer {
	ret := C.gtk_border_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_border_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkBorder)(unsafe.Pointer(paramInstance))

	ret := C.gtk_border_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_border_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkBorder)(unsafe.Pointer(paramInstance))

	C.gtk_border_free(cValueInstance)
}

func Fn_gtk_cell_area_class_find_cell_property(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkCellAreaClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_cell_area_class_find_cell_property(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_class_install_cell_property(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellAreaClass)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

	C.gtk_cell_area_class_install_cell_property(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_cell_area_class_list_cell_properties : has array return

func Fn_gtk_cell_renderer_class_set_accessible_type(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GtkCellRendererClass)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	C.gtk_cell_renderer_class_set_accessible_type(cValueInstance, cValue0)
}

func Fn_gtk_container_class_find_child_property(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GObjectClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_container_class_find_child_property(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_class_handle_border_width(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainerClass)(unsafe.Pointer(paramInstance))

	C.gtk_container_class_handle_border_width(cValueInstance)
}

// UNSUPPORTED : gtk_container_class_install_child_properties : has non-string array param pspecs

func Fn_gtk_container_class_install_child_property(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainerClass)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

	C.gtk_container_class_install_child_property(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_container_class_list_child_properties : has array return

func Fn_gtk_gradient_new_linear(param0 float64, param1 float64, param2 float64, param3 float64) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	ret := C.gtk_gradient_new_linear(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gradient_new_radial(param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	ret := C.gtk_gradient_new_radial(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gradient_add_color_stop(paramInstance unsafe.Pointer, param0 float64, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (*C.GtkSymbolicColor)(unsafe.Pointer(param1))

	C.gtk_gradient_add_color_stop(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_gradient_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gradient_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gradient_resolve(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProperties)(unsafe.Pointer(param0))

	cValue1 := (**C.cairo_pattern_t)(unsafe.Pointer(param1))

	ret := C.gtk_gradient_resolve(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_gradient_resolve_for_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	ret := C.gtk_gradient_resolve_for_context(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gradient_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	ret := C.gtk_gradient_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_gradient_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkGradient)(unsafe.Pointer(paramInstance))

	C.gtk_gradient_unref(cValueInstance)
}

func Fn_gtk_icon_set_new() unsafe.Pointer {
	ret := C.gtk_icon_set_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_new_from_pixbuf(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	ret := C.gtk_icon_set_new_from_pixbuf(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_add_source(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkIconSource)(unsafe.Pointer(param0))

	C.gtk_icon_set_add_source(cValueInstance, cValue0)
}

func Fn_gtk_icon_set_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconSet)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_set_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_set_get_sizes : has non-string array param sizes

func Fn_gtk_icon_set_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconSet)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_set_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_render_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 unsafe.Pointer, param5 string) unsafe.Pointer {
	cValueInstance := (*C.GtkIconSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (C.GtkTextDirection)(param1)

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkIconSize)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	ret := C.gtk_icon_set_render_icon(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_render_icon_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconSet)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_icon_set_render_icon_pixbuf(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconSet)(unsafe.Pointer(paramInstance))

	C.gtk_icon_set_unref(cValueInstance)
}

func Fn_gtk_icon_source_new() unsafe.Pointer {
	ret := C.gtk_icon_source_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_source_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_source_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	C.gtk_icon_source_free(cValueInstance)
}

func Fn_gtk_icon_source_get_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_get_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_source_get_direction_wildcarded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_get_direction_wildcarded(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_icon_source_get_filename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_get_filename(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_icon_source_get_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_get_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_icon_source_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_source_get_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_get_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_source_get_size_wildcarded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_get_size_wildcarded(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_icon_source_get_state(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_get_state(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_source_get_state_wildcarded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_source_get_state_wildcarded(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_icon_source_set_direction(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextDirection)(param0)

	C.gtk_icon_source_set_direction(cValueInstance, cValue0)
}

func Fn_gtk_icon_source_set_direction_wildcarded(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_icon_source_set_direction_wildcarded(cValueInstance, cValue0)
}

func Fn_gtk_icon_source_set_filename(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_source_set_filename(cValueInstance, cValue0)
}

func Fn_gtk_icon_source_set_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_source_set_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_icon_source_set_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_icon_source_set_pixbuf(cValueInstance, cValue0)
}

func Fn_gtk_icon_source_set_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkIconSize)(param0)

	C.gtk_icon_source_set_size(cValueInstance, cValue0)
}

func Fn_gtk_icon_source_set_size_wildcarded(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_icon_source_set_size_wildcarded(cValueInstance, cValue0)
}

func Fn_gtk_icon_source_set_state(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	C.gtk_icon_source_set_state(cValueInstance, cValue0)
}

func Fn_gtk_icon_source_set_state_wildcarded(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkIconSource)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_icon_source_set_state_wildcarded(cValueInstance, cValue0)
}

func Fn_gtk_paper_size_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_paper_size_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_new_custom(param0 string, param1 string, param2 float64, param3 float64, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.GtkUnit)(param4)

	ret := C.gtk_paper_size_new_custom(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_new_from_key_file(param0 unsafe.Pointer, param1 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_paper_size_new_from_key_file(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_new_from_ppd(param0 string, param1 string, param2 float64, param3 float64) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	ret := C.gtk_paper_size_new_from_ppd(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paper_size_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	C.gtk_paper_size_free(cValueInstance)
}

func Fn_gtk_paper_size_get_default_bottom_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_default_bottom_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_default_left_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_default_left_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_default_right_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_default_right_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_default_top_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_default_top_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_display_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_get_display_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_paper_size_get_height(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_height(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_paper_size_get_ppd_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_get_ppd_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_paper_size_get_width(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_paper_size_get_width(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_paper_size_is_custom(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_is_custom(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_paper_size_is_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPaperSize)(unsafe.Pointer(param0))

	ret := C.gtk_paper_size_is_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_paper_size_is_ipp(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paper_size_is_ipp(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_paper_size_set_size(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 int) {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.GtkUnit)(param2)

	C.gtk_paper_size_set_size(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_paper_size_to_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkPaperSize)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_paper_size_to_key_file(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_paper_size_get_default() string {
	ret := C.gtk_paper_size_get_default()

	return C.GoString(ret)
}

func Fn_gtk_paper_size_get_paper_sizes(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.gtk_paper_size_get_paper_sizes(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_rc_property_parse_border(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	ret := C.gtk_rc_property_parse_border(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_rc_property_parse_color(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	ret := C.gtk_rc_property_parse_color(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_rc_property_parse_enum(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	ret := C.gtk_rc_property_parse_enum(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_rc_property_parse_flags(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	ret := C.gtk_rc_property_parse_flags(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_rc_property_parse_requisition(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.GString)(unsafe.Pointer(param1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	ret := C.gtk_rc_property_parse_requisition(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_create_app_info(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_info_create_app_info(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_info_exists(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_exists(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_get_added(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_added(cValueInstance)

	return (int64)(ret)
}

func Fn_gtk_recent_info_get_age(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_age(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_recent_info_get_application_info(paramInstance unsafe.Pointer, param0 string, param1 *string, param2 *uint, param3 *int64) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	var cValue1String *C.gchar
	cValue1 := &cValue1String

	cValue2 := (*C.guint)(unsafe.Pointer(param2))

	cValue3 := (*C.time_t)(unsafe.Pointer(param3))

	ret := C.gtk_recent_info_get_application_info(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	param1String := C.GoString(cValue1String)
	*param1 = param1String

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_recent_info_get_applications : has array return

func Fn_gtk_recent_info_get_description(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_description(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_display_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_display_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_gicon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_gicon(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_recent_info_get_groups : has array return

func Fn_gtk_recent_info_get_icon(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_recent_info_get_icon(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_info_get_mime_type(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_mime_type(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_modified(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_modified(cValueInstance)

	return (int64)(ret)
}

func Fn_gtk_recent_info_get_private_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_private_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_get_short_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_short_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_uri_display(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_uri_display(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_get_visited(paramInstance unsafe.Pointer) int64 {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_get_visited(cValueInstance)

	return (int64)(ret)
}

func Fn_gtk_recent_info_has_application(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_recent_info_has_application(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_has_group(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_recent_info_has_group(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_is_local(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_is_local(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_last_application(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_last_application(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_info_match(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentInfo)(unsafe.Pointer(param0))

	ret := C.gtk_recent_info_match(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_info_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_info_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentInfo)(unsafe.Pointer(paramInstance))

	C.gtk_recent_info_unref(cValueInstance)
}

func Fn_gtk_requisition_new() unsafe.Pointer {
	ret := C.gtk_requisition_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_requisition_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRequisition)(unsafe.Pointer(paramInstance))

	ret := C.gtk_requisition_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_requisition_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRequisition)(unsafe.Pointer(paramInstance))

	C.gtk_requisition_free(cValueInstance)
}

func Fn_gtk_selection_data_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_selection_data_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	C.gtk_selection_data_free(cValueInstance)
}

// UNSUPPORTED : gtk_selection_data_get_data : has array return

func Fn_gtk_selection_data_get_data_type(paramInstance unsafe.Pointer) gdk.Atom {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_data_type(cValueInstance)

	return gdk.Atom(unsafe.Pointer(ret))
}

// UNSUPPORTED : gtk_selection_data_get_data_with_length : has array return

func Fn_gtk_selection_data_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_selection_data_get_format(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_format(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_selection_data_get_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_length(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_selection_data_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_selection_data_get_selection(paramInstance unsafe.Pointer) gdk.Atom {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_selection(cValueInstance)

	return gdk.Atom(unsafe.Pointer(ret))
}

func Fn_gtk_selection_data_get_target(paramInstance unsafe.Pointer) gdk.Atom {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_target(cValueInstance)

	return gdk.Atom(unsafe.Pointer(ret))
}

// UNSUPPORTED : gtk_selection_data_get_targets : has non-string array param targets

func Fn_gtk_selection_data_get_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_get_text(cValueInstance)

	return C.GoString((*C.char)(unsafe.Pointer(ret)))
}

// UNSUPPORTED : gtk_selection_data_get_uris : has array return

// UNSUPPORTED : gtk_selection_data_set : has non-string array param data

func Fn_gtk_selection_data_set_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	ret := C.gtk_selection_data_set_pixbuf(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_selection_data_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_selection_data_set_text(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_selection_data_set_uris(paramInstance unsafe.Pointer, param0 []string) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	param0Len := len(param0)
	cValue0Array := C.malloc((C.ulong)(param0Len) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(cValue0Array))
	param0Slice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue0Array))[:param0Len:param0Len]
	for param0i, param0String := range param0 {
		param0Slice[param0i] = (*C.gchar)(C.CString(param0String))
		defer C.free(unsafe.Pointer(param0Slice[param0i]))
	}
	cValue0 := &param0Slice[0]

	ret := C.gtk_selection_data_set_uris(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_selection_data_targets_include_image(paramInstance unsafe.Pointer, param0 bool) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	ret := C.gtk_selection_data_targets_include_image(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_selection_data_targets_include_rich_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	ret := C.gtk_selection_data_targets_include_rich_text(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_selection_data_targets_include_text(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_targets_include_text(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_selection_data_targets_include_uri(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSelectionData)(unsafe.Pointer(paramInstance))

	ret := C.gtk_selection_data_targets_include_uri(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_stock_item_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStockItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stock_item_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_stock_item_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStockItem)(unsafe.Pointer(paramInstance))

	C.gtk_stock_item_free(cValueInstance)
}

func Fn_gtk_symbolic_color_new_alpha(param0 unsafe.Pointer, param1 float64) unsafe.Pointer {
	cValue0 := (*C.GtkSymbolicColor)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	ret := C.gtk_symbolic_color_new_alpha(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_new_literal(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	ret := C.gtk_symbolic_color_new_literal(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_new_mix(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64) unsafe.Pointer {
	cValue0 := (*C.GtkSymbolicColor)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkSymbolicColor)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	ret := C.gtk_symbolic_color_new_mix(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_new_name(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_symbolic_color_new_name(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_new_shade(param0 unsafe.Pointer, param1 float64) unsafe.Pointer {
	cValue0 := (*C.GtkSymbolicColor)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	ret := C.gtk_symbolic_color_new_shade(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSymbolicColor)(unsafe.Pointer(paramInstance))

	ret := C.gtk_symbolic_color_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_symbolic_color_resolve(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSymbolicColor)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProperties)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	ret := C.gtk_symbolic_color_resolve(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_symbolic_color_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkSymbolicColor)(unsafe.Pointer(paramInstance))

	ret := C.gtk_symbolic_color_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_symbolic_color_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSymbolicColor)(unsafe.Pointer(paramInstance))

	C.gtk_symbolic_color_unref(cValueInstance)
}

func Fn_gtk_target_entry_new(param0 string, param1 uint, param2 uint) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	ret := C.gtk_target_entry_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_target_entry_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTargetEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_target_entry_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_target_entry_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTargetEntry)(unsafe.Pointer(paramInstance))

	C.gtk_target_entry_free(cValueInstance)
}

// UNSUPPORTED : gtk_target_list_new : has non-string array param targets

func Fn_gtk_target_list_add(paramInstance unsafe.Pointer, param0 gdk.Atom, param1 uint, param2 uint) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	C.gtk_target_list_add(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_target_list_add_image_targets(paramInstance unsafe.Pointer, param0 uint, param1 bool) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := toCBool(param1)

	C.gtk_target_list_add_image_targets(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_target_list_add_rich_text_targets(paramInstance unsafe.Pointer, param0 uint, param1 bool, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := toCBool(param1)

	cValue2 := (*C.GtkTextBuffer)(unsafe.Pointer(param2))

	C.gtk_target_list_add_rich_text_targets(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_target_list_add_table : has non-string array param targets

func Fn_gtk_target_list_add_text_targets(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_target_list_add_text_targets(cValueInstance, cValue0)
}

func Fn_gtk_target_list_add_uri_targets(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_target_list_add_uri_targets(cValueInstance, cValue0)
}

func Fn_gtk_target_list_find(paramInstance unsafe.Pointer, param0 gdk.Atom, param1 *uint) bool {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	ret := C.gtk_target_list_find(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_target_list_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	ret := C.gtk_target_list_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_target_list_remove(paramInstance unsafe.Pointer, param0 gdk.Atom) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	C.gtk_target_list_remove(cValueInstance, cValue0)
}

func Fn_gtk_target_list_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTargetList)(unsafe.Pointer(paramInstance))

	C.gtk_target_list_unref(cValueInstance)
}

func Fn_gtk_text_attributes_new() unsafe.Pointer {
	ret := C.gtk_text_attributes_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_attributes_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextAttributes)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_attributes_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_attributes_copy_values(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextAttributes)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextAttributes)(unsafe.Pointer(param0))

	C.gtk_text_attributes_copy_values(cValueInstance, cValue0)
}

func Fn_gtk_text_attributes_ref(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextAttributes)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_attributes_ref(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_attributes_unref(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextAttributes)(unsafe.Pointer(paramInstance))

	C.gtk_text_attributes_unref(cValueInstance)
}

func Fn_gtk_text_iter_backward_char(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_char(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_chars(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_chars(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_cursor_position(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_cursor_position(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_cursor_positions(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_cursor_positions(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_iter_backward_find_char : has callback

func Fn_gtk_text_iter_backward_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_lines(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_lines(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_search(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkTextSearchFlags)(param1)

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkTextIter)(unsafe.Pointer(param3))

	cValue4 := (*C.GtkTextIter)(unsafe.Pointer(param4))

	ret := C.gtk_text_iter_backward_search(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_sentence_start(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_sentence_start(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_sentence_starts(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_sentence_starts(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_to_tag_toggle(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_backward_to_tag_toggle(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_cursor_position(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_visible_cursor_position(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_cursor_positions(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_visible_cursor_positions(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_visible_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_lines(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_visible_lines(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_word_start(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_visible_word_start(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_visible_word_starts(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_visible_word_starts(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_word_start(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_backward_word_start(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_word_starts(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_backward_word_starts(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_begins_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_begins_tag(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_can_insert(paramInstance unsafe.Pointer, param0 bool) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	ret := C.gtk_text_iter_can_insert(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_compare(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_compare(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_text_iter_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_editable(paramInstance unsafe.Pointer, param0 bool) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	ret := C.gtk_text_iter_editable(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_ends_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_ends_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_ends_sentence(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_ends_sentence(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_ends_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_ends_tag(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_ends_word(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_ends_word(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_equal(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_equal(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_char(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_char(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_chars(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_chars(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_cursor_position(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_cursor_position(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_cursor_positions(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_cursor_positions(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_iter_forward_find_char : has callback

func Fn_gtk_text_iter_forward_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_lines(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_lines(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_search(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkTextSearchFlags)(param1)

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkTextIter)(unsafe.Pointer(param3))

	cValue4 := (*C.GtkTextIter)(unsafe.Pointer(param4))

	ret := C.gtk_text_iter_forward_search(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_sentence_end(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_sentence_end(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_sentence_ends(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_sentence_ends(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_to_end(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	C.gtk_text_iter_forward_to_end(cValueInstance)
}

func Fn_gtk_text_iter_forward_to_line_end(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_to_line_end(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_to_tag_toggle(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_forward_to_tag_toggle(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_cursor_position(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_visible_cursor_position(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_cursor_positions(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_visible_cursor_positions(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_visible_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_lines(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_visible_lines(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_word_end(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_visible_word_end(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_visible_word_ends(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_visible_word_ends(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_word_end(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_forward_word_end(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_word_ends(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_text_iter_forward_word_ends(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	C.gtk_text_iter_free(cValueInstance)
}

func Fn_gtk_text_iter_get_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextAttributes)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_get_attributes(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_get_buffer(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_buffer(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_bytes_in_line(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_bytes_in_line(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_char(paramInstance unsafe.Pointer) rune {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_char(cValueInstance)

	return (rune)(ret)
}

func Fn_gtk_text_iter_get_chars_in_line(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_chars_in_line(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_child_anchor(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_child_anchor(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_language(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_language(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_line(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_line(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_line_index(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_line_index(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_line_offset(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_line_offset(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_marks(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_marks(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_offset(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_offset(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_slice(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_get_slice(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_text_iter_get_tags(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_tags(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_get_text(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_text_iter_get_toggled_tags(paramInstance unsafe.Pointer, param0 bool) unsafe.Pointer {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	ret := C.gtk_text_iter_get_toggled_tags(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_visible_line_index(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_visible_line_index(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_visible_line_offset(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_get_visible_line_offset(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_visible_slice(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_get_visible_slice(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_text_iter_get_visible_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_get_visible_text(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_text_iter_has_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_has_tag(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_in_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	ret := C.gtk_text_iter_in_range(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_inside_sentence(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_inside_sentence(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_inside_word(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_inside_word(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_is_cursor_position(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_is_cursor_position(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_is_end(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_is_end(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_is_start(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_is_start(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_order(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_iter_order(cValueInstance, cValue0)
}

func Fn_gtk_text_iter_set_line(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_iter_set_line(cValueInstance, cValue0)
}

func Fn_gtk_text_iter_set_line_index(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_iter_set_line_index(cValueInstance, cValue0)
}

func Fn_gtk_text_iter_set_line_offset(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_iter_set_line_offset(cValueInstance, cValue0)
}

func Fn_gtk_text_iter_set_offset(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_iter_set_offset(cValueInstance, cValue0)
}

func Fn_gtk_text_iter_set_visible_line_index(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_iter_set_visible_line_index(cValueInstance, cValue0)
}

func Fn_gtk_text_iter_set_visible_line_offset(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_iter_set_visible_line_offset(cValueInstance, cValue0)
}

func Fn_gtk_text_iter_starts_line(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_starts_line(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_starts_sentence(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_starts_sentence(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_starts_word(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_iter_starts_word(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_toggles_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextIter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	ret := C.gtk_text_iter_toggles_tag(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_iter_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeIter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_iter_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_iter_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeIter)(unsafe.Pointer(paramInstance))

	C.gtk_tree_iter_free(cValueInstance)
}

func Fn_gtk_tree_path_new() unsafe.Pointer {
	ret := C.gtk_tree_path_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_path_new_first() unsafe.Pointer {
	ret := C.gtk_tree_path_new_first()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_path_new_from_indices : has varargs

// UNSUPPORTED : gtk_tree_path_new_from_indicesv : has non-string array param indices

func Fn_gtk_tree_path_new_from_string(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_tree_path_new_from_string(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_path_append_index(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_path_append_index(cValueInstance, cValue0)
}

func Fn_gtk_tree_path_compare(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_path_compare(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_tree_path_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_path_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_path_down(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	C.gtk_tree_path_down(cValueInstance)
}

func Fn_gtk_tree_path_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	C.gtk_tree_path_free(cValueInstance)
}

func Fn_gtk_tree_path_get_depth(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_path_get_depth(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_tree_path_get_indices_with_depth : has array return

func Fn_gtk_tree_path_is_ancestor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_path_is_ancestor(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_path_is_descendant(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_path_is_descendant(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_path_next(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	C.gtk_tree_path_next(cValueInstance)
}

func Fn_gtk_tree_path_prepend_index(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_path_prepend_index(cValueInstance, cValue0)
}

func Fn_gtk_tree_path_prev(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_path_prev(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_path_to_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_path_to_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tree_path_up(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreePath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_path_up(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_row_reference_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	ret := C.gtk_tree_row_reference_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_row_reference_new_proxy(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeModel)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTreePath)(unsafe.Pointer(param2))

	ret := C.gtk_tree_row_reference_new_proxy(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_row_reference_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeRowReference)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_row_reference_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_row_reference_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeRowReference)(unsafe.Pointer(paramInstance))

	C.gtk_tree_row_reference_free(cValueInstance)
}

func Fn_gtk_tree_row_reference_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeRowReference)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_row_reference_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_row_reference_get_path(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeRowReference)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_row_reference_get_path(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_row_reference_valid(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeRowReference)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_row_reference_valid(cValueInstance)

	return toGoBool(ret)
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

// UNSUPPORTED : gtk_tree_row_reference_reordered : has non-string array param new_order

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : has callback

func Fn_gtk_widget_class_find_style_property(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_widget_class_find_style_property(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_class_install_style_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidgetClass)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	C.gtk_widget_class_install_style_property(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : has callback

// UNSUPPORTED : gtk_widget_class_list_style_properties : has array return

// UNSUPPORTED : gtk_widget_class_set_connect_func : has callback

func Fn_gtk_widget_path_new() unsafe.Pointer {
	ret := C.gtk_widget_path_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_append_type(paramInstance unsafe.Pointer, param0 uint64) int {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.gtk_widget_path_append_type(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_widget_path_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_path_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	C.gtk_widget_path_free(cValueInstance)
}

func Fn_gtk_widget_path_get_object_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_path_get_object_type(cValueInstance)

	return (uint64)(ret)
}

func Fn_gtk_widget_path_has_parent(paramInstance unsafe.Pointer, param0 uint64) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.gtk_widget_path_has_parent(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_is_type(paramInstance unsafe.Pointer, param0 uint64) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.gtk_widget_path_is_type(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_add_class(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_path_iter_add_class(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_iter_add_region(paramInstance unsafe.Pointer, param0 int, param1 string, param2 int) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GtkRegionFlags)(param2)

	C.gtk_widget_path_iter_add_region(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_widget_path_iter_clear_classes(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_path_iter_clear_classes(cValueInstance, cValue0)
}

func Fn_gtk_widget_path_iter_clear_regions(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_path_iter_clear_regions(cValueInstance, cValue0)
}

func Fn_gtk_widget_path_iter_get_name(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_get_name(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_widget_path_iter_get_object_type(paramInstance unsafe.Pointer, param0 int) uint64 {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_get_object_type(cValueInstance, cValue0)

	return (uint64)(ret)
}

func Fn_gtk_widget_path_iter_get_sibling_index(paramInstance unsafe.Pointer, param0 int) uint {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_get_sibling_index(cValueInstance, cValue0)

	return (uint)(ret)
}

func Fn_gtk_widget_path_iter_get_siblings(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_get_siblings(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_iter_has_class(paramInstance unsafe.Pointer, param0 int, param1 string) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_widget_path_iter_has_class(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_name(paramInstance unsafe.Pointer, param0 int, param1 string) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_widget_path_iter_has_name(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_qclass(paramInstance unsafe.Pointer, param0 int, param1 uint32) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GQuark)(param1)

	ret := C.gtk_widget_path_iter_has_qclass(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_qname(paramInstance unsafe.Pointer, param0 int, param1 uint32) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GQuark)(param1)

	ret := C.gtk_widget_path_iter_has_qname(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_qregion(paramInstance unsafe.Pointer, param0 int, param1 uint32, param2 *int) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GQuark)(param1)

	cValue2 := (*C.GtkRegionFlags)(unsafe.Pointer(param2))

	ret := C.gtk_widget_path_iter_has_qregion(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_has_region(paramInstance unsafe.Pointer, param0 int, param1 string, param2 *int) bool {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GtkRegionFlags)(unsafe.Pointer(param2))

	ret := C.gtk_widget_path_iter_has_region(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_widget_path_iter_list_classes(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_list_classes(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_iter_list_regions(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_widget_path_iter_list_regions(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_path_iter_remove_class(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_path_iter_remove_class(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_iter_remove_region(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_path_iter_remove_region(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_iter_set_name(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_widget_path_iter_set_name(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_iter_set_object_type(paramInstance unsafe.Pointer, param0 int, param1 uint64) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GType)(param1)

	C.gtk_widget_path_iter_set_object_type(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_path_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_path_length(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_path_prepend_type(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GtkWidgetPath)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	C.gtk_widget_path_prepend_type(cValueInstance, cValue0)
}

func Fn_gtk_accel_groups_activate(param0 unsafe.Pointer, param1 uint, param2 int) bool {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	ret := C.gtk_accel_groups_activate(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_accel_groups_from_object(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	ret := C.gtk_accel_groups_from_object(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accelerator_get_default_mod_mask() int {
	ret := C.gtk_accelerator_get_default_mod_mask()

	return (int)(ret)
}

func Fn_gtk_accelerator_get_label(param0 uint, param1 int) string {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	ret := C.gtk_accelerator_get_label(cValue0, cValue1)

	return C.GoString(ret)
}

func Fn_gtk_accelerator_name(param0 uint, param1 int) string {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	ret := C.gtk_accelerator_name(cValue0, cValue1)

	return C.GoString(ret)
}

func Fn_gtk_accelerator_parse(param0 string, param1 *uint, param2 *int) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkModifierType)(unsafe.Pointer(param2))

	C.gtk_accelerator_parse(cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : has non-string array param accelerator_codes

func Fn_gtk_accelerator_set_default_mod_mask(param0 int) {
	cValue0 := (C.GdkModifierType)(param0)

	C.gtk_accelerator_set_default_mod_mask(cValue0)
}

func Fn_gtk_accelerator_valid(param0 uint, param1 int) bool {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	ret := C.gtk_accelerator_valid(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_alternative_dialog_button_order(param0 unsafe.Pointer) bool {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	ret := C.gtk_alternative_dialog_button_order(cValue0)

	return toGoBool(ret)
}

func Fn_gtk_bindings_activate(param0 unsafe.Pointer, param1 uint, param2 int) bool {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	ret := C.gtk_bindings_activate(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_bindings_activate_event(param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkEventKey)(unsafe.Pointer(param1))

	ret := C.gtk_bindings_activate_event(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_cairo_should_draw_window(param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkWindow)(unsafe.Pointer(param1))

	ret := C.gtk_cairo_should_draw_window(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_cairo_transform_to_window(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkWindow)(unsafe.Pointer(param2))

	C.gtk_cairo_transform_to_window(cValue0, cValue1, cValue2)
}

func Fn_gtk_check_version(param0 uint, param1 uint, param2 uint) string {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (C.guint)(param2)

	ret := C.gtk_check_version(cValue0, cValue1, cValue2)

	return C.GoString(ret)
}

func Fn_gtk_device_grab_add(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkDevice)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	C.gtk_device_grab_add(cValue0, cValue1, cValue2)
}

func Fn_gtk_device_grab_remove(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkDevice)(unsafe.Pointer(param1))

	C.gtk_device_grab_remove(cValue0, cValue1)
}

func Fn_gtk_disable_setlocale() {
	C.gtk_disable_setlocale()
}

func Fn_gtk_distribute_natural_allocation(param0 int, param1 uint, param2 unsafe.Pointer) int {
	cValue0 := (C.gint)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := (*C.GtkRequestedSize)(unsafe.Pointer(param2))

	ret := C.gtk_distribute_natural_allocation(cValue0, cValue1, cValue2)

	return (int)(ret)
}

func Fn_gtk_drag_finish(param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint32) {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.guint32)(param3)

	C.gtk_drag_finish(cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_drag_get_source_widget(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	ret := C.gtk_drag_get_source_widget(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_gtk_draw_insertion_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool, param4 int, param5 bool) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	cValue3 := toCBool(param3)

	cValue4 := (C.GtkTextDirection)(param4)

	cValue5 := toCBool(param5)

	C.gtk_draw_insertion_cursor(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_events_pending() bool {
	ret := C.gtk_events_pending()

	return toGoBool(ret)
}

func Fn_gtk_false() bool {
	ret := C.gtk_false()

	return toGoBool(ret)
}

func Fn_gtk_get_binary_age() uint {
	ret := C.gtk_get_binary_age()

	return (uint)(ret)
}

func Fn_gtk_get_current_event() unsafe.Pointer {
	ret := C.gtk_get_current_event()

	return unsafe.Pointer(ret)
}

func Fn_gtk_get_current_event_device() unsafe.Pointer {
	ret := C.gtk_get_current_event_device()

	return unsafe.Pointer(ret)
}

func Fn_gtk_get_current_event_state(param0 *int) bool {
	cValue0 := (*C.GdkModifierType)(unsafe.Pointer(param0))

	ret := C.gtk_get_current_event_state(cValue0)

	return toGoBool(ret)
}

func Fn_gtk_get_current_event_time() uint32 {
	ret := C.gtk_get_current_event_time()

	return (uint32)(ret)
}

func Fn_gtk_get_debug_flags() uint {
	ret := C.gtk_get_debug_flags()

	return (uint)(ret)
}

func Fn_gtk_get_default_language() unsafe.Pointer {
	ret := C.gtk_get_default_language()

	return unsafe.Pointer(ret)
}

func Fn_gtk_get_event_widget(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	ret := C.gtk_get_event_widget(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_get_interface_age() uint {
	ret := C.gtk_get_interface_age()

	return (uint)(ret)
}

func Fn_gtk_get_major_version() uint {
	ret := C.gtk_get_major_version()

	return (uint)(ret)
}

func Fn_gtk_get_micro_version() uint {
	ret := C.gtk_get_micro_version()

	return (uint)(ret)
}

func Fn_gtk_get_minor_version() uint {
	ret := C.gtk_get_minor_version()

	return (uint)(ret)
}

func Fn_gtk_get_option_group(param0 bool) unsafe.Pointer {
	cValue0 := toCBool(param0)

	ret := C.gtk_get_option_group(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_grab_get_current() unsafe.Pointer {
	ret := C.gtk_grab_get_current()

	return unsafe.Pointer(ret)
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

func Fn_gtk_init_check(param0 *int, param1 *[]string) bool {
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

	ret := C.gtk_init_check(cValue0, cValue1)

	param1OutLen := int(*cValue0)
	param1Out := make([]string, param1OutLen, param1OutLen)
	if param1OutLen > 0 {
		param1OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1ArrayPointer))[:param1OutLen:param1OutLen]
		for param1Outi, param1OutCString := range param1OutCSlice {
			param1Out[param1Outi] = C.GoString(param1OutCString)
		}
	}
	*param1 = param1Out

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_init_with_args : has non-string array param entries

// UNSUPPORTED : gtk_key_snooper_install : has callback

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

func Fn_gtk_main_iteration() bool {
	ret := C.gtk_main_iteration()

	return toGoBool(ret)
}

func Fn_gtk_main_iteration_do(param0 bool) bool {
	cValue0 := toCBool(param0)

	ret := C.gtk_main_iteration_do(cValue0)

	return toGoBool(ret)
}

func Fn_gtk_main_level() uint {
	ret := C.gtk_main_level()

	return (uint)(ret)
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

func Fn_gtk_parse_args(param0 *int, param1 *[]string) bool {
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

	ret := C.gtk_parse_args(cValue0, cValue1)

	param1OutLen := int(*cValue0)
	param1Out := make([]string, param1OutLen, param1OutLen)
	if param1OutLen > 0 {
		param1OutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(cValue1ArrayPointer))[:param1OutLen:param1OutLen]
		for param1Outi, param1OutCString := range param1OutCSlice {
			param1Out[param1Outi] = C.GoString(param1OutCString)
		}
	}
	*param1 = param1Out

	return toGoBool(ret)
}

func Fn_gtk_print_run_page_setup_dialog(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkPageSetup)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkPrintSettings)(unsafe.Pointer(param2))

	ret := C.gtk_print_run_page_setup_dialog(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : has callback

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

func Fn_gtk_rc_find_module_in_path(param0 string) string {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_rc_find_module_in_path(cValue0)

	return C.GoString(ret)
}

func Fn_gtk_rc_find_pixmap_in_path(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) string {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	cValue1 := (*C.GScanner)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.gtk_rc_find_pixmap_in_path(cValue0, cValue1, cValue2)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_rc_get_default_files : has array return

func Fn_gtk_rc_get_im_module_file() string {
	ret := C.gtk_rc_get_im_module_file()

	return C.GoString(ret)
}

func Fn_gtk_rc_get_im_module_path() string {
	ret := C.gtk_rc_get_im_module_path()

	return C.GoString(ret)
}

func Fn_gtk_rc_get_module_dir() string {
	ret := C.gtk_rc_get_module_dir()

	return C.GoString(ret)
}

func Fn_gtk_rc_get_style(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_rc_get_style(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_rc_get_style_by_paths(param0 unsafe.Pointer, param1 string, param2 string, param3 uint64) unsafe.Pointer {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	cValue1 := (*C.char)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.char)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.GType)(param3)

	ret := C.gtk_rc_get_style_by_paths(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_rc_get_theme_dir() string {
	ret := C.gtk_rc_get_theme_dir()

	return C.GoString(ret)
}

func Fn_gtk_rc_parse(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_rc_parse(cValue0)
}

func Fn_gtk_rc_parse_color(param0 unsafe.Pointer, param1 unsafe.Pointer) uint {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	ret := C.gtk_rc_parse_color(cValue0, cValue1)

	return (uint)(ret)
}

func Fn_gtk_rc_parse_color_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) uint {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRcStyle)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkColor)(unsafe.Pointer(param2))

	ret := C.gtk_rc_parse_color_full(cValue0, cValue1, cValue2)

	return (uint)(ret)
}

func Fn_gtk_rc_parse_priority(param0 unsafe.Pointer, param1 *int) uint {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkPathPriorityType)(unsafe.Pointer(param1))

	ret := C.gtk_rc_parse_priority(cValue0, cValue1)

	return (uint)(ret)
}

func Fn_gtk_rc_parse_state(param0 unsafe.Pointer, param1 *int) uint {
	cValue0 := (*C.GScanner)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkStateType)(unsafe.Pointer(param1))

	ret := C.gtk_rc_parse_state(cValue0, cValue1)

	return (uint)(ret)
}

func Fn_gtk_rc_parse_string(param0 string) {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_rc_parse_string(cValue0)
}

func Fn_gtk_rc_reparse_all() bool {
	ret := C.gtk_rc_reparse_all()

	return toGoBool(ret)
}

func Fn_gtk_rc_reparse_all_for_settings(param0 unsafe.Pointer, param1 bool) bool {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	ret := C.gtk_rc_reparse_all_for_settings(cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_rc_reset_styles(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkSettings)(unsafe.Pointer(param0))

	C.gtk_rc_reset_styles(cValue0)
}

func Fn_gtk_rc_scanner_new() unsafe.Pointer {
	ret := C.gtk_rc_scanner_new()

	return unsafe.Pointer(ret)
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

func Fn_gtk_render_activity(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_activity(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_arrow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_arrow(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_background(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_background(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_check(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_check(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_expander(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_expander(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_extension(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	cValue6 := (C.GtkPositionType)(param6)

	C.gtk_render_extension(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gtk_render_focus(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_focus(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_frame(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_frame(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_frame_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int, param7 float64, param8 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	cValue6 := (C.GtkPositionType)(param6)

	cValue7 := (C.gdouble)(param7)

	cValue8 := (C.gdouble)(param8)

	C.gtk_render_frame_gap(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6, cValue7, cValue8)
}

func Fn_gtk_render_handle(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_handle(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_icon_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) unsafe.Pointer {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkIconSource)(unsafe.Pointer(param1))

	cValue2 := (C.GtkIconSize)(param2)

	ret := C.gtk_render_icon_pixbuf(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_render_layout(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 unsafe.Pointer) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (*C.PangoLayout)(unsafe.Pointer(param4))

	C.gtk_render_layout(cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_render_line(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_line(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_option(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	C.gtk_render_option(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_render_slider(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int) {
	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.cairo_t)(unsafe.Pointer(param1))

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	cValue6 := (C.GtkOrientation)(param6)

	C.gtk_render_slider(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
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

// UNSUPPORTED : gtk_selection_add_targets : has non-string array param targets

func Fn_gtk_selection_clear_targets(param0 unsafe.Pointer, param1 gdk.Atom) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	C.gtk_selection_clear_targets(cValue0, cValue1)
}

func Fn_gtk_selection_convert(param0 unsafe.Pointer, param1 gdk.Atom, param2 gdk.Atom, param3 uint32) bool {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	cValue2 := (C.GdkAtom)(unsafe.Pointer(param2))

	cValue3 := (C.guint32)(param3)

	ret := C.gtk_selection_convert(cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_selection_owner_set(param0 unsafe.Pointer, param1 gdk.Atom, param2 uint32) bool {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	cValue2 := (C.guint32)(param2)

	ret := C.gtk_selection_owner_set(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_selection_owner_set_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 gdk.Atom, param3 uint32) bool {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.GdkAtom)(unsafe.Pointer(param2))

	cValue3 := (C.guint32)(param3)

	ret := C.gtk_selection_owner_set_for_display(cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_selection_remove_all(param0 unsafe.Pointer) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_selection_remove_all(cValue0)
}

func Fn_gtk_set_debug_flags(param0 uint) {
	cValue0 := (C.guint)(param0)

	C.gtk_set_debug_flags(cValue0)
}

// UNSUPPORTED : gtk_show_about_dialog : has varargs

func Fn_gtk_show_uri(param0 unsafe.Pointer, param1 string, param2 uint32, error unsafe.Pointer) bool {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.guint32)(param2)

	cError := (**C.GError)(error)

	ret := C.gtk_show_uri(cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_stock_add : has non-string array param items

// UNSUPPORTED : gtk_stock_add_static : has non-string array param items

func Fn_gtk_stock_list_ids() unsafe.Pointer {
	ret := C.gtk_stock_list_ids()

	return unsafe.Pointer(ret)
}

func Fn_gtk_stock_lookup(param0 string, param1 unsafe.Pointer) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkStockItem)(unsafe.Pointer(param1))

	ret := C.gtk_stock_lookup(cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_stock_set_translate_func : has callback

// UNSUPPORTED : gtk_target_table_free : has non-string array param targets

// UNSUPPORTED : gtk_target_table_new_from_list : has array return

// UNSUPPORTED : gtk_targets_include_image : has non-string array param targets

// UNSUPPORTED : gtk_targets_include_rich_text : has non-string array param targets

// UNSUPPORTED : gtk_targets_include_text : has non-string array param targets

// UNSUPPORTED : gtk_targets_include_uri : has non-string array param targets

func Fn_gtk_test_create_simple_window(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_test_create_simple_window(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_test_create_widget : has varargs

// UNSUPPORTED : gtk_test_display_button_window : has varargs

func Fn_gtk_test_find_label(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_test_find_label(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_test_find_sibling(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	ret := C.gtk_test_find_sibling(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_test_find_widget(param0 unsafe.Pointer, param1 string, param2 uint64) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.GType)(param2)

	ret := C.gtk_test_find_widget(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_test_init : has varargs

// UNSUPPORTED : gtk_test_list_all_types : has array return

func Fn_gtk_test_register_all_types() {
	C.gtk_test_register_all_types()
}

func Fn_gtk_test_slider_get_value(param0 unsafe.Pointer) float64 {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_test_slider_get_value(cValue0)

	return (float64)(ret)
}

func Fn_gtk_test_slider_set_perc(param0 unsafe.Pointer, param1 float64) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.double)(param1)

	C.gtk_test_slider_set_perc(cValue0, cValue1)
}

func Fn_gtk_test_spin_button_click(param0 unsafe.Pointer, param1 uint, param2 bool) bool {
	cValue0 := (*C.GtkSpinButton)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := toCBool(param2)

	ret := C.gtk_test_spin_button_click(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_test_text_get(param0 unsafe.Pointer) string {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_test_text_get(cValue0)

	return C.GoString(ret)
}

func Fn_gtk_test_text_set(param0 unsafe.Pointer, param1 string) {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_test_text_set(cValue0, cValue1)
}

func Fn_gtk_test_widget_click(param0 unsafe.Pointer, param1 uint, param2 int) bool {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	ret := C.gtk_test_widget_click(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_test_widget_send_key(param0 unsafe.Pointer, param1 uint, param2 int) bool {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	ret := C.gtk_test_widget_send_key(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_tree_get_row_drag_data(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer) bool {
	cValue0 := (*C.GtkSelectionData)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreeModel)(unsafe.Pointer(param1))

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	ret := C.gtk_tree_get_row_drag_data(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_row_reference_reordered : has non-string array param new_order

func Fn_gtk_tree_set_row_drag_data(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValue0 := (*C.GtkSelectionData)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeModel)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTreePath)(unsafe.Pointer(param2))

	ret := C.gtk_tree_set_row_drag_data(cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_true() bool {
	ret := C.gtk_true()

	return toGoBool(ret)
}

func Fn_gtk_about_dialog_new() unsafe.Pointer {
	ret := C.gtk_about_dialog_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_about_dialog_get_artists : has array return

// UNSUPPORTED : gtk_about_dialog_get_authors : has array return

func Fn_gtk_about_dialog_get_comments(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_comments(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_copyright(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_copyright(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_about_dialog_get_documenters : has array return

func Fn_gtk_about_dialog_get_license(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_license(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_license_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_license_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_about_dialog_get_logo(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_logo(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_about_dialog_get_logo_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_logo_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_program_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_program_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_translator_credits(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_translator_credits(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_version(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_version(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_website(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_website(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_website_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_website_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_about_dialog_get_wrap_license(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_about_dialog_get_wrap_license(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_about_dialog_set_license_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkAboutDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkLicense)(param0)

	C.gtk_about_dialog_set_license_type(cValueInstance, cValue0)
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

func Fn_gtk_accel_group_new() unsafe.Pointer {
	ret := C.gtk_accel_group_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_group_activate(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer, param2 uint, param3 int) bool {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GQuark)(param0)

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	cValue3 := (C.GdkModifierType)(param3)

	ret := C.gtk_accel_group_activate(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
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

func Fn_gtk_accel_group_disconnect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	ret := C.gtk_accel_group_disconnect(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_accel_group_disconnect_key(paramInstance unsafe.Pointer, param0 uint, param1 int) bool {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	ret := C.gtk_accel_group_disconnect_key(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_accel_group_find : has callback

func Fn_gtk_accel_group_get_is_locked(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_accel_group_get_is_locked(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_accel_group_get_modifier_mask(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_accel_group_get_modifier_mask(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_accel_group_lock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	C.gtk_accel_group_lock(cValueInstance)
}

// UNSUPPORTED : gtk_accel_group_query : has array return

func Fn_gtk_accel_group_unlock(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAccelGroup)(unsafe.Pointer(paramInstance))

	C.gtk_accel_group_unlock(cValueInstance)
}

func Fn_gtk_accel_group_from_accel_closure(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GClosure)(unsafe.Pointer(param0))

	ret := C.gtk_accel_group_from_accel_closure(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_label_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_accel_label_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_label_get_accel_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_accel_label_get_accel_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_label_get_accel_width(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_accel_label_get_accel_width(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_accel_label_refetch(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAccelLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_accel_label_refetch(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_accel_map_change_entry(param0 string, param1 uint, param2 int, param3 bool) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	cValue3 := toCBool(param3)

	ret := C.gtk_accel_map_change_entry(cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_accel_map_foreach : has callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : has callback

func Fn_gtk_accel_map_get() unsafe.Pointer {
	ret := C.gtk_accel_map_get()

	return unsafe.Pointer(ret)
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

func Fn_gtk_accel_map_lookup_entry(param0 string, param1 unsafe.Pointer) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkAccelKey)(unsafe.Pointer(param1))

	ret := C.gtk_accel_map_lookup_entry(cValue0, cValue1)

	return toGoBool(ret)
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

func Fn_gtk_accessible_get_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAccessible)(unsafe.Pointer(paramInstance))

	ret := C.gtk_accessible_get_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accessible_set_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAccessible)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_accessible_set_widget(cValueInstance, cValue0)
}

func Fn_gtk_action_new(param0 string, param1 string, param2 string, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.gtk_action_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
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

func Fn_gtk_action_create_icon(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkIconSize)(param0)

	ret := C.gtk_action_create_icon(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_create_menu(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_create_menu(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_create_menu_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_create_menu_item(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_create_tool_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_create_tool_item(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_disconnect_accelerator(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	C.gtk_action_disconnect_accelerator(cValueInstance)
}

func Fn_gtk_action_get_accel_closure(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_accel_closure(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_get_accel_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_accel_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_always_show_image(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_always_show_image(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_get_gicon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_gicon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_get_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_is_important(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_is_important(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_proxies(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_proxies(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_get_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_get_short_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_short_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_stock_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_stock_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_tooltip(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_tooltip(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_get_visible_horizontal(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_visible_horizontal(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_get_visible_vertical(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_get_visible_vertical(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_is_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_is_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_is_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_is_visible(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_action_set_always_show_image(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_action_set_always_show_image(cValueInstance, cValue0)
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

func Fn_gtk_action_group_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_action_group_new(cValue0)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : gtk_action_group_add_actions : has non-string array param entries

// UNSUPPORTED : gtk_action_group_add_actions_full : has callback

// UNSUPPORTED : gtk_action_group_add_radio_actions : has callback

// UNSUPPORTED : gtk_action_group_add_radio_actions_full : has callback

// UNSUPPORTED : gtk_action_group_add_toggle_actions : has non-string array param entries

// UNSUPPORTED : gtk_action_group_add_toggle_actions_full : has callback

func Fn_gtk_action_group_get_action(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_action_group_get_action(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_action_group_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_group_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_action_group_get_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_group_get_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_group_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_group_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_action_group_list_actions(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_action_group_list_actions(cValueInstance)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : gtk_action_group_set_translate_func : has callback

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

func Fn_gtk_action_group_translate_string(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GtkActionGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_action_group_translate_string(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_adjustment_new(param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	cValue5 := (C.gdouble)(param5)

	ret := C.gtk_adjustment_new(cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
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

func Fn_gtk_adjustment_get_lower(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_lower(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_get_page_increment(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_page_increment(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_get_page_size(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_page_size(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_get_step_increment(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_step_increment(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_get_upper(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_upper(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_adjustment_get_value(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkAdjustment)(unsafe.Pointer(paramInstance))

	ret := C.gtk_adjustment_get_value(cValueInstance)

	return (float64)(ret)
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

func Fn_gtk_alignment_new(param0 float32, param1 float32, param2 float32, param3 float32) unsafe.Pointer {
	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	cValue2 := (C.gfloat)(param2)

	cValue3 := (C.gfloat)(param3)

	ret := C.gtk_alignment_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
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

func Fn_gtk_app_chooser_button_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_app_chooser_button_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_button_append_custom_item(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GIcon)(unsafe.Pointer(param2))

	C.gtk_app_chooser_button_append_custom_item(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_app_chooser_button_append_separator(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	C.gtk_app_chooser_button_append_separator(cValueInstance)
}

func Fn_gtk_app_chooser_button_get_heading(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_button_get_heading(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_app_chooser_button_get_show_dialog_item(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_button_get_show_dialog_item(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_button_set_active_custom_item(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_app_chooser_button_set_active_custom_item(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_button_set_heading(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_app_chooser_button_set_heading(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_button_set_show_dialog_item(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_button_set_show_dialog_item(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_dialog_new(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GtkDialogFlags)(param1)

	cValue2 := (*C.GFile)(unsafe.Pointer(param2))

	ret := C.gtk_app_chooser_dialog_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_dialog_new_for_content_type(param0 unsafe.Pointer, param1 int, param2 string) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GtkDialogFlags)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.gtk_app_chooser_dialog_new_for_content_type(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_dialog_get_heading(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAppChooserDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_dialog_get_heading(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_app_chooser_dialog_get_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAppChooserDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_dialog_get_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_dialog_set_heading(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAppChooserDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_app_chooser_dialog_set_heading(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_app_chooser_widget_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_widget_get_default_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_default_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_app_chooser_widget_get_show_all(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_all(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_get_show_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_get_show_fallback(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_fallback(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_get_show_other(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_other(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_get_show_recommended(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_widget_get_show_recommended(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_app_chooser_widget_set_default_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_app_chooser_widget_set_default_text(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_show_all(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_all(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_show_default(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_default(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_show_fallback(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_fallback(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_show_other(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_other(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_widget_set_show_recommended(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkAppChooserWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_app_chooser_widget_set_show_recommended(cValueInstance, cValue0)
}

func Fn_gtk_application_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GApplicationFlags)(param1)

	ret := C.gtk_application_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_application_add_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_application_add_window(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_application_get_accels_for_action : has array return

// UNSUPPORTED : gtk_application_get_actions_for_accel : has array return

func Fn_gtk_application_get_windows(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	ret := C.gtk_application_get_windows(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_application_list_action_descriptions : has array return

func Fn_gtk_application_remove_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkApplication)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_application_remove_window(cValueInstance, cValue0)
}

func Fn_gtk_arrow_new(param0 int, param1 int) unsafe.Pointer {
	cValue0 := (C.GtkArrowType)(param0)

	cValue1 := (C.GtkShadowType)(param1)

	ret := C.gtk_arrow_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_arrow_set(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkArrow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkArrowType)(param0)

	cValue1 := (C.GtkShadowType)(param1)

	C.gtk_arrow_set(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_aspect_frame_new(param0 string, param1 float32, param2 float32, param3 float32, param4 bool) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gfloat)(param1)

	cValue2 := (C.gfloat)(param2)

	cValue3 := (C.gfloat)(param3)

	cValue4 := toCBool(param4)

	ret := C.gtk_aspect_frame_new(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_aspect_frame_set(paramInstance unsafe.Pointer, param0 float32, param1 float32, param2 float32, param3 bool) {
	cValueInstance := (*C.GtkAspectFrame)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gfloat)(param0)

	cValue1 := (C.gfloat)(param1)

	cValue2 := (C.gfloat)(param2)

	cValue3 := toCBool(param3)

	C.gtk_aspect_frame_set(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_assistant_new() unsafe.Pointer {
	ret := C.gtk_assistant_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_assistant_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_assistant_add_action_widget(cValueInstance, cValue0)
}

func Fn_gtk_assistant_append_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_append_page(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_assistant_commit(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	C.gtk_assistant_commit(cValueInstance)
}

func Fn_gtk_assistant_get_current_page(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	ret := C.gtk_assistant_get_current_page(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_assistant_get_n_pages(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	ret := C.gtk_assistant_get_n_pages(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_assistant_get_nth_page(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_assistant_get_nth_page(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_assistant_get_page_complete(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_complete(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_assistant_get_page_header_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_header_image(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_assistant_get_page_side_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_side_image(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_assistant_get_page_title(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_title(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_assistant_get_page_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_type(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_assistant_insert_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_assistant_insert_page(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_assistant_next_page(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	C.gtk_assistant_next_page(cValueInstance)
}

func Fn_gtk_assistant_prepend_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_prepend_page(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_assistant_previous_page(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	C.gtk_assistant_previous_page(cValueInstance)
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

// UNSUPPORTED : gtk_assistant_set_forward_page_func : has callback

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

func Fn_gtk_bin_get_child(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkBin)(unsafe.Pointer(paramInstance))

	ret := C.gtk_bin_get_child(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_box_new(param0 int, param1 int) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_box_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_box_get_homogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_box_get_homogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_box_get_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_box_get_spacing(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_builder_new() unsafe.Pointer {
	ret := C.gtk_builder_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_builder_add_callback_symbol : has callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : has varargs

func Fn_gtk_builder_add_from_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_builder_add_from_file(cValueInstance, cValue0, cError)

	return (uint)(ret)
}

func Fn_gtk_builder_add_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gsize)(param1)

	cError := (**C.GError)(error)

	ret := C.gtk_builder_add_from_string(cValueInstance, cValue0, cValue1, cError)

	return (uint)(ret)
}

func Fn_gtk_builder_add_objects_from_file(paramInstance unsafe.Pointer, param0 string, param1 []string, error unsafe.Pointer) uint {
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

	ret := C.gtk_builder_add_objects_from_file(cValueInstance, cValue0, cValue1, cError)

	return (uint)(ret)
}

func Fn_gtk_builder_add_objects_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64, param2 []string, error unsafe.Pointer) uint {
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

	ret := C.gtk_builder_add_objects_from_string(cValueInstance, cValue0, cValue1, cValue2, cError)

	return (uint)(ret)
}

func Fn_gtk_builder_connect_signals(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.gtk_builder_connect_signals(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_builder_connect_signals_full : has callback

func Fn_gtk_builder_extend_with_template(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint64, param2 string, param3 uint64, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GType)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gsize)(param3)

	cError := (**C.GError)(error)

	ret := C.gtk_builder_extend_with_template(cValueInstance, cValue0, cValue1, cValue2, cValue3, cError)

	return (uint)(ret)
}

func Fn_gtk_builder_get_object(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_builder_get_object(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_builder_get_objects(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	ret := C.gtk_builder_get_objects(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_builder_get_translation_domain(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	ret := C.gtk_builder_get_translation_domain(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_builder_get_type_from_name(paramInstance unsafe.Pointer, param0 string) uint64 {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_builder_get_type_from_name(cValueInstance, cValue0)

	return (uint64)(ret)
}

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

func Fn_gtk_builder_set_translation_domain(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_builder_set_translation_domain(cValueInstance, cValue0)
}

func Fn_gtk_builder_value_from_string(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.gtk_builder_value_from_string(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_gtk_builder_value_from_string_type(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkBuilder)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.gtk_builder_value_from_string_type(cValueInstance, cValue0, cValue1, cValue2, cError)

	return toGoBool(ret)
}

func Fn_gtk_button_new() unsafe.Pointer {
	ret := C.gtk_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_new_from_stock(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_button_new_from_stock(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_new_with_label(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_button_new_with_label(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_new_with_mnemonic(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_button_new_with_mnemonic(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_gtk_button_get_event_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_event_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_get_focus_on_click(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_focus_on_click(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_button_get_image(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_image(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_get_image_position(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_image_position(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_button_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_button_get_relief(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_relief(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_button_get_use_stock(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_use_stock(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_button_get_use_underline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_get_use_underline(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_button_box_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	ret := C.gtk_button_box_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_box_get_child_secondary(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkButtonBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_button_box_get_child_secondary(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_button_box_get_layout(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkButtonBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_button_box_get_layout(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_calendar_new() unsafe.Pointer {
	ret := C.gtk_calendar_new()

	return unsafe.Pointer(ret)
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

func Fn_gtk_calendar_get_day_is_marked(paramInstance unsafe.Pointer, param0 uint) bool {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gtk_calendar_get_day_is_marked(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_calendar_get_detail_height_rows(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_calendar_get_detail_height_rows(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_calendar_get_detail_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_calendar_get_detail_width_chars(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_calendar_get_display_options(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCalendar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_calendar_get_display_options(cValueInstance)

	return (int)(ret)
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

// UNSUPPORTED : gtk_calendar_set_detail_func : has callback

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

func Fn_gtk_cell_area_activate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 bool) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	cValue3 := (C.GtkCellRendererState)(param3)

	cValue4 := toCBool(param4)

	ret := C.gtk_cell_area_activate(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_activate_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkEvent)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (C.GtkCellRendererState)(param4)

	ret := C.gtk_cell_area_activate_cell(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_cell_area_add(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_add_focus_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	C.gtk_cell_area_add_focus_sibling(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_cell_area_add_with_properties : has varargs

func Fn_gtk_cell_area_apply_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 bool) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := toCBool(param3)

	C.gtk_cell_area_apply_attributes(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_attribute_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	C.gtk_cell_area_attribute_connect(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_attribute_disconnect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_cell_area_attribute_disconnect(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_cell_area_cell_get : has varargs

func Fn_gtk_cell_area_cell_get_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_cell_area_cell_get_property(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_cell_area_cell_get_valist : has va_list

// UNSUPPORTED : gtk_cell_area_cell_set : has varargs

func Fn_gtk_cell_area_cell_set_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_cell_area_cell_set_property(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_cell_area_cell_set_valist : has va_list

func Fn_gtk_cell_area_copy_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	ret := C.gtk_cell_area_copy_context(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_create_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_create_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) int {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkEvent)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (C.GtkCellRendererState)(param4)

	ret := C.gtk_cell_area_event(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return (int)(ret)
}

func Fn_gtk_cell_area_focus(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkDirectionType)(param0)

	ret := C.gtk_cell_area_focus(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_cell_area_foreach : has callback

// UNSUPPORTED : gtk_cell_area_foreach_alloc : has callback

func Fn_gtk_cell_area_get_cell_allocation(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkCellRenderer)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (*C.GdkRectangle)(unsafe.Pointer(param4))

	C.gtk_cell_area_get_cell_allocation(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_cell_area_get_cell_at_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 int, param5 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	cValue3 := (C.gint)(param3)

	cValue4 := (C.gint)(param4)

	cValue5 := (*C.GdkRectangle)(unsafe.Pointer(param5))

	ret := C.gtk_cell_area_get_cell_at_position(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_current_path_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_current_path_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_cell_area_get_edit_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_edit_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_edited_cell(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_edited_cell(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_focus_cell(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_focus_cell(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_focus_from_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	ret := C.gtk_cell_area_get_focus_from_sibling(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_focus_siblings(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	ret := C.gtk_cell_area_get_focus_siblings(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_get_preferred_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_cell_area_get_preferred_height(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 *int, param4 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	C.gtk_cell_area_get_preferred_height_for_width(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_cell_area_get_preferred_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_cell_area_get_preferred_width(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 *int, param4 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	C.gtk_cell_area_get_preferred_width_for_height(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)
}

func Fn_gtk_cell_area_get_request_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_get_request_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_cell_area_has_renderer(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	ret := C.gtk_cell_area_has_renderer(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_inner_cell_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	C.gtk_cell_area_inner_cell_area(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_is_activatable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_is_activatable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_is_focus_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	ret := C.gtk_cell_area_is_focus_sibling(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_cell_area_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_cell_area_remove(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_remove_focus_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellRenderer)(unsafe.Pointer(param1))

	C.gtk_cell_area_remove_focus_sibling(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_render(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int, param6 bool) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAreaContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.cairo_t)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (*C.GdkRectangle)(unsafe.Pointer(param4))

	cValue5 := (C.GtkCellRendererState)(param5)

	cValue6 := toCBool(param6)

	C.gtk_cell_area_render(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5, cValue6)
}

func Fn_gtk_cell_area_request_renderer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 int, param4 *int, param5 *int) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (C.GtkOrientation)(param1)

	cValue2 := (*C.GtkWidget)(unsafe.Pointer(param2))

	cValue3 := (C.gint)(param3)

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (*C.gint)(unsafe.Pointer(param5))

	C.gtk_cell_area_request_renderer(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

func Fn_gtk_cell_area_set_focus_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_cell_area_set_focus_cell(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_stop_editing(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellArea)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_area_stop_editing(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_box_new() unsafe.Pointer {
	ret := C.gtk_cell_area_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_box_get_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCellAreaBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_box_get_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_cell_area_box_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 bool) {
	cValueInstance := (*C.GtkCellAreaBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	cValue3 := toCBool(param3)

	C.gtk_cell_area_box_pack_end(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_box_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 bool) {
	cValueInstance := (*C.GtkCellAreaBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	cValue3 := toCBool(param3)

	C.gtk_cell_area_box_pack_start(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_area_box_set_spacing(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCellAreaBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_cell_area_box_set_spacing(cValueInstance, cValue0)
}

func Fn_gtk_cell_area_context_allocate(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_cell_area_context_allocate(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_get_allocation(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_cell_area_context_get_allocation(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_get_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_area_context_get_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_area_context_get_preferred_height(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_cell_area_context_get_preferred_height(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_cell_area_context_get_preferred_height_for_width(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_context_get_preferred_width(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_cell_area_context_get_preferred_width(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_cell_area_context_get_preferred_width_for_height(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_area_context_push_preferred_height(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_cell_area_context_push_preferred_height(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_push_preferred_width(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_cell_area_context_push_preferred_width(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_area_context_reset(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellAreaContext)(unsafe.Pointer(paramInstance))

	C.gtk_cell_area_context_reset(cValueInstance)
}

func Fn_gtk_cell_renderer_activate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) bool {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (*C.GdkRectangle)(unsafe.Pointer(param4))

	cValue5 := (C.GtkCellRendererState)(param5)

	ret := C.gtk_cell_renderer_activate(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_get_aligned_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkCellRendererState)(param1)

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	C.gtk_cell_renderer_get_aligned_area(cValueInstance, cValue0, cValue1, cValue2, cValue3)
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

func Fn_gtk_cell_renderer_get_preferred_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_cell_renderer_get_preferred_height(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_renderer_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_cell_renderer_get_preferred_height_for_width(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_renderer_get_preferred_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRequisition)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkRequisition)(unsafe.Pointer(param2))

	C.gtk_cell_renderer_get_preferred_size(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_renderer_get_preferred_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_cell_renderer_get_preferred_width(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_renderer_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_cell_renderer_get_preferred_width_for_height(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_cell_renderer_get_request_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_get_request_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_cell_renderer_get_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_get_sensitive(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_cell_renderer_get_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) int {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkCellRendererState)(param1)

	ret := C.gtk_cell_renderer_get_state(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_cell_renderer_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_is_activatable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_is_activatable(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_cell_renderer_start_editing(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) unsafe.Pointer {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GdkRectangle)(unsafe.Pointer(param3))

	cValue4 := (*C.GdkRectangle)(unsafe.Pointer(param4))

	cValue5 := (C.GtkCellRendererState)(param5)

	ret := C.gtk_cell_renderer_start_editing(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_stop_editing(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellRenderer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_renderer_stop_editing(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_accel_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_accel_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_combo_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_combo_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_pixbuf_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_pixbuf_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_progress_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_progress_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_spin_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_spin_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_spinner_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_spinner_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_text_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_text_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_text_set_fixed_height_from_font(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkCellRendererText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_cell_renderer_text_set_fixed_height_from_font(cValueInstance, cValue0)
}

func Fn_gtk_cell_renderer_toggle_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_toggle_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_toggle_get_activatable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_toggle_get_activatable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_toggle_get_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_toggle_get_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_toggle_get_radio(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellRendererToggle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_renderer_toggle_get_radio(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_cell_view_new() unsafe.Pointer {
	ret := C.gtk_cell_view_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_new_with_context(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkCellAreaContext)(unsafe.Pointer(param1))

	ret := C.gtk_cell_view_new_with_context(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_new_with_markup(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_cell_view_new_with_markup(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_new_with_pixbuf(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	ret := C.gtk_cell_view_new_with_pixbuf(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_new_with_text(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_cell_view_new_with_text(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_get_displayed_row(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_view_get_displayed_row(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_get_draw_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_view_get_draw_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_view_get_fit_model(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_view_get_fit_model(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_cell_view_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_view_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_view_get_size_of_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRequisition)(unsafe.Pointer(param1))

	ret := C.gtk_cell_view_get_size_of_row(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_cell_view_set_background_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_background_color(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_background_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_background_rgba(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_displayed_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_displayed_row(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_draw_sensitive(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_view_set_draw_sensitive(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_fit_model(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_cell_view_set_fit_model(cValueInstance, cValue0)
}

func Fn_gtk_cell_view_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_cell_view_set_model(cValueInstance, cValue0)
}

func Fn_gtk_check_button_new() unsafe.Pointer {
	ret := C.gtk_check_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_button_new_with_label(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_check_button_new_with_label(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_button_new_with_mnemonic(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_check_button_new_with_mnemonic(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_menu_item_new() unsafe.Pointer {
	ret := C.gtk_check_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_menu_item_new_with_label(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_check_menu_item_new_with_label(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_menu_item_new_with_mnemonic(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_check_menu_item_new_with_mnemonic(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_menu_item_get_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_check_menu_item_get_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_check_menu_item_get_draw_as_radio(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_check_menu_item_get_draw_as_radio(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_check_menu_item_get_inconsistent(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCheckMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_check_menu_item_get_inconsistent(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_clipboard_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_clipboard_get_owner(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_get_owner(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_clipboard_request_contents : has callback

// UNSUPPORTED : gtk_clipboard_request_image : has callback

// UNSUPPORTED : gtk_clipboard_request_rich_text : has callback

// UNSUPPORTED : gtk_clipboard_request_targets : has callback

// UNSUPPORTED : gtk_clipboard_request_text : has callback

// UNSUPPORTED : gtk_clipboard_request_uris : has callback

// UNSUPPORTED : gtk_clipboard_set_can_store : has non-string array param targets

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

// UNSUPPORTED : gtk_clipboard_set_with_data : has callback

// UNSUPPORTED : gtk_clipboard_set_with_owner : has callback

func Fn_gtk_clipboard_store(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	C.gtk_clipboard_store(cValueInstance)
}

func Fn_gtk_clipboard_wait_for_contents(paramInstance unsafe.Pointer, param0 gdk.Atom) unsafe.Pointer {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	ret := C.gtk_clipboard_wait_for_contents(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_clipboard_wait_for_image(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_wait_for_image(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_clipboard_wait_for_rich_text : has array return

// UNSUPPORTED : gtk_clipboard_wait_for_targets : has non-string array param targets

func Fn_gtk_clipboard_wait_for_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_wait_for_text(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_clipboard_wait_for_uris : has array return

func Fn_gtk_clipboard_wait_is_image_available(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_wait_is_image_available(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_clipboard_wait_is_rich_text_available(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	ret := C.gtk_clipboard_wait_is_rich_text_available(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_clipboard_wait_is_target_available(paramInstance unsafe.Pointer, param0 gdk.Atom) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	ret := C.gtk_clipboard_wait_is_target_available(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_clipboard_wait_is_text_available(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_wait_is_text_available(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_clipboard_wait_is_uris_available(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkClipboard)(unsafe.Pointer(paramInstance))

	ret := C.gtk_clipboard_wait_is_uris_available(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_clipboard_get(param0 gdk.Atom) unsafe.Pointer {
	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	ret := C.gtk_clipboard_get(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_clipboard_get_for_display(param0 unsafe.Pointer, param1 gdk.Atom) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.GdkAtom)(unsafe.Pointer(param1))

	ret := C.gtk_clipboard_get_for_display(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_button_new() unsafe.Pointer {
	ret := C.gtk_color_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_button_new_with_color(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	ret := C.gtk_color_button_new_with_color(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_button_new_with_rgba(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	ret := C.gtk_color_button_new_with_rgba(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_button_get_alpha(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_button_get_alpha(cValueInstance)

	return (uint16)(ret)
}

func Fn_gtk_color_button_get_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_button_get_color(cValueInstance, cValue0)
}

func Fn_gtk_color_button_get_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_button_get_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_button_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_button_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_color_button_get_use_alpha(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_button_get_use_alpha(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_color_button_set_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_button_set_rgba(cValueInstance, cValue0)
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

func Fn_gtk_color_selection_new() unsafe.Pointer {
	ret := C.gtk_color_selection_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_selection_get_current_alpha(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_selection_get_current_alpha(cValueInstance)

	return (uint16)(ret)
}

func Fn_gtk_color_selection_get_current_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_selection_get_current_color(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_get_current_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_selection_get_current_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_get_has_opacity_control(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_selection_get_has_opacity_control(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_color_selection_get_has_palette(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_selection_get_has_palette(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_color_selection_get_previous_alpha(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_selection_get_previous_alpha(cValueInstance)

	return (uint16)(ret)
}

func Fn_gtk_color_selection_get_previous_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkColor)(unsafe.Pointer(param0))

	C.gtk_color_selection_get_previous_color(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_get_previous_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_selection_get_previous_rgba(cValueInstance, cValue0)
}

func Fn_gtk_color_selection_is_adjusting(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_selection_is_adjusting(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_color_selection_set_current_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_selection_set_current_rgba(cValueInstance, cValue0)
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

func Fn_gtk_color_selection_set_previous_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkColorSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	C.gtk_color_selection_set_previous_rgba(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_color_selection_palette_from_string : has non-string array param colors

// UNSUPPORTED : gtk_color_selection_palette_to_string : has non-string array param colors

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : has callback

func Fn_gtk_color_selection_dialog_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_color_selection_dialog_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_selection_dialog_get_color_selection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkColorSelectionDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_color_selection_dialog_get_color_selection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_new() unsafe.Pointer {
	ret := C.gtk_combo_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_new_with_area(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	ret := C.gtk_combo_box_new_with_area(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_new_with_area_and_entry(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	ret := C.gtk_combo_box_new_with_area_and_entry(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_new_with_entry() unsafe.Pointer {
	ret := C.gtk_combo_box_new_with_entry()

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_new_with_model(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	ret := C.gtk_combo_box_new_with_model(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_new_with_model_and_entry(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	ret := C.gtk_combo_box_new_with_model_and_entry(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_get_active(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_active(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_active_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_active_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_combo_box_get_active_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_combo_box_get_active_iter(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_combo_box_get_add_tearoffs(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_add_tearoffs(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_combo_box_get_button_sensitivity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_button_sensitivity(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_column_span_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_column_span_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_entry_text_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_entry_text_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_focus_on_click(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_focus_on_click(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_combo_box_get_has_entry(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_has_entry(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_combo_box_get_id_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_id_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_get_popup_accessible(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_popup_accessible(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_get_popup_fixed_width(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_popup_fixed_width(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

func Fn_gtk_combo_box_get_row_span_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_row_span_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_combo_box_get_wrap_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_get_wrap_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_combo_box_popdown(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_popdown(cValueInstance)
}

func Fn_gtk_combo_box_popup(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_popup(cValueInstance)
}

func Fn_gtk_combo_box_popup_for_device(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	C.gtk_combo_box_popup_for_device(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_active(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_active(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_active_id(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_combo_box_set_active_id(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_combo_box_set_entry_text_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_entry_text_column(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_combo_box_set_focus_on_click(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_id_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_set_id_column(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	C.gtk_combo_box_set_model(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_set_popup_fixed_width(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkComboBox)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_combo_box_set_popup_fixed_width(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : has callback

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

func Fn_gtk_combo_box_text_new() unsafe.Pointer {
	ret := C.gtk_combo_box_text_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_text_new_with_entry() unsafe.Pointer {
	ret := C.gtk_combo_box_text_new_with_entry()

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_text_append(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_combo_box_text_append(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_combo_box_text_append_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_combo_box_text_append_text(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_text_get_active_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	ret := C.gtk_combo_box_text_get_active_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_combo_box_text_insert(paramInstance unsafe.Pointer, param0 int, param1 string, param2 string) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_combo_box_text_insert(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_combo_box_text_insert_text(paramInstance unsafe.Pointer, param0 int, param1 string) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_combo_box_text_insert_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_combo_box_text_prepend(paramInstance unsafe.Pointer, param0 string, param1 string) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_combo_box_text_prepend(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_combo_box_text_prepend_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_combo_box_text_prepend_text(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_text_remove(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_combo_box_text_remove(cValueInstance, cValue0)
}

func Fn_gtk_combo_box_text_remove_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkComboBoxText)(unsafe.Pointer(paramInstance))

	C.gtk_combo_box_text_remove_all(cValueInstance)
}

func Fn_gtk_container_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_container_add(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_container_add_with_properties : has varargs

func Fn_gtk_container_check_resize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	C.gtk_container_check_resize(cValueInstance)
}

// UNSUPPORTED : gtk_container_child_get : has varargs

func Fn_gtk_container_child_get_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_container_child_get_property(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_container_child_get_valist : has va_list

// UNSUPPORTED : gtk_container_child_set : has varargs

func Fn_gtk_container_child_set_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_container_child_set_property(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_container_child_set_valist : has va_list

func Fn_gtk_container_child_type(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_container_child_type(cValueInstance)

	return (uint64)(ret)
}

// UNSUPPORTED : gtk_container_forall : has callback

// UNSUPPORTED : gtk_container_foreach : has callback

func Fn_gtk_container_get_border_width(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_container_get_border_width(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_container_get_children(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_container_get_children(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_get_focus_chain(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GList)(unsafe.Pointer(param0))

	ret := C.gtk_container_get_focus_chain(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_container_get_focus_child(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_container_get_focus_child(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_get_focus_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_container_get_focus_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_get_focus_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_container_get_focus_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_get_path_for_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_container_get_path_for_child(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_get_resize_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_container_get_resize_mode(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_container_cell_accessible_new() unsafe.Pointer {
	ret := C.gtk_container_cell_accessible_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_cell_accessible_add_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainerCellAccessible)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	C.gtk_container_cell_accessible_add_child(cValueInstance, cValue0)
}

func Fn_gtk_container_cell_accessible_get_children(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkContainerCellAccessible)(unsafe.Pointer(paramInstance))

	ret := C.gtk_container_cell_accessible_get_children(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_cell_accessible_remove_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainerCellAccessible)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	C.gtk_container_cell_accessible_remove_child(cValueInstance, cValue0)
}

func Fn_gtk_css_provider_new() unsafe.Pointer {
	ret := C.gtk_css_provider_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_css_provider_load_from_data : has non-string array param data

func Fn_gtk_css_provider_load_from_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCssProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.gtk_css_provider_load_from_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_css_provider_load_from_path(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCssProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_css_provider_load_from_path(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_css_provider_get_default() unsafe.Pointer {
	ret := C.gtk_css_provider_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gtk_css_provider_get_named(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_css_provider_get_named(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_new() unsafe.Pointer {
	ret := C.gtk_dialog_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_dialog_new_with_buttons : has varargs

func Fn_gtk_dialog_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_dialog_add_action_widget(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_dialog_add_button(paramInstance unsafe.Pointer, param0 string, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_dialog_add_button(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_dialog_add_buttons : has varargs

func Fn_gtk_dialog_get_action_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_dialog_get_action_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_get_content_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_dialog_get_content_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_get_response_for_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_dialog_get_response_for_widget(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_dialog_get_widget_for_response(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_dialog_get_widget_for_response(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_response(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_dialog_response(cValueInstance, cValue0)
}

func Fn_gtk_dialog_run(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_dialog_run(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_dialog_set_alternative_button_order : has varargs

// UNSUPPORTED : gtk_dialog_set_alternative_button_order_from_array : has non-string array param new_order

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

func Fn_gtk_drawing_area_new() unsafe.Pointer {
	ret := C.gtk_drawing_area_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_new() unsafe.Pointer {
	ret := C.gtk_entry_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_new_with_buffer(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkEntryBuffer)(unsafe.Pointer(param0))

	ret := C.gtk_entry_new_with_buffer(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_activates_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_activates_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_alignment(paramInstance unsafe.Pointer) float32 {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_alignment(cValueInstance)

	return (float32)(ret)
}

func Fn_gtk_entry_get_buffer(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_buffer(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_completion(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_completion(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_current_icon_drag_source(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_current_icon_drag_source(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_get_cursor_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_cursor_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_has_frame(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_has_frame(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_icon_activatable(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_activatable(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_icon_area(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gtk_entry_get_icon_area(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_get_icon_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_entry_get_icon_at_pos(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_entry_get_icon_gicon(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_gicon(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_icon_name(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_name(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_icon_pixbuf(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_pixbuf(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_icon_sensitive(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_sensitive(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_icon_stock(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_stock(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_icon_storage_type(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_storage_type(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_entry_get_icon_tooltip_markup(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_tooltip_markup(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_icon_tooltip_text(paramInstance unsafe.Pointer, param0 int) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkEntryIconPosition)(param0)

	ret := C.gtk_entry_get_icon_tooltip_text(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_inner_border(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_inner_border(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_invisible_char(paramInstance unsafe.Pointer) rune {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_invisible_char(cValueInstance)

	return (rune)(ret)
}

func Fn_gtk_entry_get_layout(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_layout(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_entry_get_layout_offsets(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_entry_get_max_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_max_length(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_get_overwrite_mode(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_overwrite_mode(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_progress_fraction(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_progress_fraction(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_entry_get_progress_pulse_step(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_progress_pulse_step(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_entry_get_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_text_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_entry_get_text_area(cValueInstance, cValue0)
}

func Fn_gtk_entry_get_text_length(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_text_length(cValueInstance)

	return (uint16)(ret)
}

func Fn_gtk_entry_get_visibility(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_visibility(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_get_width_chars(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_im_context_filter_keypress(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	ret := C.gtk_entry_im_context_filter_keypress(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_entry_layout_index_to_text_index(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_entry_layout_index_to_text_index(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_entry_progress_pulse(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_progress_pulse(cValueInstance)
}

func Fn_gtk_entry_reset_im_context(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_reset_im_context(cValueInstance)
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

func Fn_gtk_entry_text_index_to_layout_index(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_entry_text_index_to_layout_index(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_entry_unset_invisible_char(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEntry)(unsafe.Pointer(paramInstance))

	C.gtk_entry_unset_invisible_char(cValueInstance)
}

func Fn_gtk_entry_buffer_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_entry_buffer_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_buffer_delete_text(paramInstance unsafe.Pointer, param0 uint, param1 int) uint {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_entry_buffer_delete_text(cValueInstance, cValue0, cValue1)

	return (uint)(ret)
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

func Fn_gtk_entry_buffer_get_bytes(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_buffer_get_bytes(cValueInstance)

	return (uint64)(ret)
}

func Fn_gtk_entry_buffer_get_length(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_buffer_get_length(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_entry_buffer_get_max_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_buffer_get_max_length(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_buffer_get_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_buffer_get_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_entry_buffer_insert_text(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 int) uint {
	cValueInstance := (*C.GtkEntryBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	ret := C.gtk_entry_buffer_insert_text(cValueInstance, cValue0, cValue1, cValue2)

	return (uint)(ret)
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

func Fn_gtk_entry_completion_new() unsafe.Pointer {
	ret := C.gtk_entry_completion_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_completion_new_with_area(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	ret := C.gtk_entry_completion_new_with_area(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_gtk_entry_completion_get_completion_prefix(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_completion_prefix(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_entry_completion_get_entry(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_entry(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_completion_get_inline_completion(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_inline_completion(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_completion_get_inline_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_inline_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_completion_get_minimum_key_length(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_minimum_key_length(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_entry_completion_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_completion_get_popup_completion(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_popup_completion(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_completion_get_popup_set_width(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_popup_set_width(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_completion_get_popup_single_match(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_popup_single_match(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_entry_completion_get_text_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEntryCompletion)(unsafe.Pointer(paramInstance))

	ret := C.gtk_entry_completion_get_text_column(cValueInstance)

	return (int)(ret)
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

// UNSUPPORTED : gtk_entry_completion_set_match_func : has callback

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

func Fn_gtk_event_box_new() unsafe.Pointer {
	ret := C.gtk_event_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_event_box_get_above_child(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_event_box_get_above_child(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_event_box_get_visible_window(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEventBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_event_box_get_visible_window(cValueInstance)

	return toGoBool(ret)
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

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted
// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted
// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted
// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted
func Fn_gtk_expander_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_expander_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_expander_new_with_mnemonic(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_expander_new_with_mnemonic(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_expander_get_expanded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_expanded(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_expander_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_expander_get_label_fill(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_label_fill(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_expander_get_label_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_label_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_expander_get_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_expander_get_use_markup(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_use_markup(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_expander_get_use_underline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	ret := C.gtk_expander_get_use_underline(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_expander_set_label_fill(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkExpander)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_expander_set_label_fill(cValueInstance, cValue0)
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

func Fn_gtk_file_chooser_button_new(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkFileChooserAction)(param1)

	ret := C.gtk_file_chooser_button_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_button_new_with_dialog(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_file_chooser_button_new_with_dialog(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_button_get_focus_on_click(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_button_get_focus_on_click(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_button_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_button_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_button_get_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFileChooserButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_button_get_width_chars(cValueInstance)

	return (int)(ret)
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

// UNSUPPORTED : gtk_file_chooser_dialog_new : has varargs

func Fn_gtk_file_chooser_widget_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkFileChooserAction)(param0)

	ret := C.gtk_file_chooser_widget_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_filter_new() unsafe.Pointer {
	ret := C.gtk_file_filter_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_file_filter_add_custom : has callback

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

func Fn_gtk_file_filter_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFileFilterInfo)(unsafe.Pointer(param0))

	ret := C.gtk_file_filter_filter(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_filter_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_filter_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_filter_get_needed(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_filter_get_needed(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_file_filter_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_filter_set_name(cValueInstance, cValue0)
}

func Fn_gtk_fixed_new() unsafe.Pointer {
	ret := C.gtk_fixed_new()

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : gtk_flow_box_bind_model : has callback

// UNSUPPORTED : gtk_flow_box_selected_foreach : has callback

// UNSUPPORTED : gtk_flow_box_set_filter_func : has callback

// UNSUPPORTED : gtk_flow_box_set_sort_func : has callback

func Fn_gtk_font_button_new() unsafe.Pointer {
	ret := C.gtk_font_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_button_new_with_font(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_font_button_new_with_font(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_button_get_font_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_font_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_button_get_show_size(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_show_size(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_font_button_get_show_style(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_show_style(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_font_button_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_button_get_use_font(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_use_font(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_font_button_get_use_size(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_button_get_use_size(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_font_button_set_font_name(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFontButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_font_button_set_font_name(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_font_selection_new() unsafe.Pointer {
	ret := C.gtk_font_selection_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_face(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_face(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_face_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_face_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_family(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_family(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_family_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_family_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_font_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_font_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_selection_get_preview_entry(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_preview_entry(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_preview_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_preview_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_selection_get_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_font_selection_get_size_entry(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_size_entry(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_size_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_get_size_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_set_font_name(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_font_selection_set_font_name(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_font_selection_set_preview_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_selection_set_preview_text(cValueInstance, cValue0)
}

func Fn_gtk_font_selection_dialog_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_font_selection_dialog_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_dialog_get_cancel_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_dialog_get_cancel_button(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_dialog_get_font_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_dialog_get_font_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_selection_dialog_get_font_selection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_dialog_get_font_selection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_dialog_get_ok_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_dialog_get_ok_button(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_dialog_get_preview_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_selection_dialog_get_preview_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_font_selection_dialog_set_font_name(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_font_selection_dialog_set_font_name(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_font_selection_dialog_set_preview_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFontSelectionDialog)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_font_selection_dialog_set_preview_text(cValueInstance, cValue0)
}

func Fn_gtk_frame_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_frame_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_frame_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	ret := C.gtk_frame_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_frame_get_label_align(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gfloat)(unsafe.Pointer(param0))

	cValue1 := (*C.gfloat)(unsafe.Pointer(param1))

	C.gtk_frame_get_label_align(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_frame_get_label_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	ret := C.gtk_frame_get_label_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_frame_get_shadow_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFrame)(unsafe.Pointer(paramInstance))

	ret := C.gtk_frame_get_shadow_type(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_gesture_get_last_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkGesture)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventSequence)(unsafe.Pointer(param0))

	ret := C.gtk_gesture_get_last_event(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_gesture_stylus_get_axes : has non-string array param axes

func Fn_gtk_grid_new() unsafe.Pointer {
	ret := C.gtk_grid_new()

	return unsafe.Pointer(ret)
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

func Fn_gtk_grid_get_column_homogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	ret := C.gtk_grid_get_column_homogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_grid_get_column_spacing(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	ret := C.gtk_grid_get_column_spacing(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_grid_get_row_homogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	ret := C.gtk_grid_get_row_homogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_grid_get_row_spacing(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkGrid)(unsafe.Pointer(paramInstance))

	ret := C.gtk_grid_get_row_spacing(cValueInstance)

	return (uint)(ret)
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

func Fn_gtk_hbox_new(param0 bool, param1 int) unsafe.Pointer {
	cValue0 := toCBool(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_hbox_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_hbutton_box_new() unsafe.Pointer {
	ret := C.gtk_hbutton_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_hpaned_new() unsafe.Pointer {
	ret := C.gtk_hpaned_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_hsv_new() unsafe.Pointer {
	ret := C.gtk_hsv_new()

	return unsafe.Pointer(ret)
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

func Fn_gtk_hsv_is_adjusting(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkHSV)(unsafe.Pointer(paramInstance))

	ret := C.gtk_hsv_is_adjusting(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_hscale_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	ret := C.gtk_hscale_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_hscale_new_with_range(param0 float64, param1 float64, param2 float64) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	ret := C.gtk_hscale_new_with_range(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_hscrollbar_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	ret := C.gtk_hscrollbar_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_hseparator_new() unsafe.Pointer {
	ret := C.gtk_hseparator_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_handle_box_new() unsafe.Pointer {
	ret := C.gtk_handle_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_handle_box_get_child_detached(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_handle_box_get_child_detached(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_handle_box_get_handle_position(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_handle_box_get_handle_position(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_handle_box_get_shadow_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_handle_box_get_shadow_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_handle_box_get_snap_edge(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkHandleBox)(unsafe.Pointer(paramInstance))

	ret := C.gtk_handle_box_get_snap_edge(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_im_context_delete_surrounding(paramInstance unsafe.Pointer, param0 int, param1 int) bool {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_im_context_delete_surrounding(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_im_context_filter_keypress(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	ret := C.gtk_im_context_filter_keypress(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_im_context_get_surrounding(paramInstance unsafe.Pointer, param0 *string, param1 *int) bool {
	cValueInstance := (*C.GtkIMContext)(unsafe.Pointer(paramInstance))

	var cValue0String *C.gchar
	cValue0 := &cValue0String

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	ret := C.gtk_im_context_get_surrounding(cValueInstance, cValue0, cValue1)

	param0String := C.GoString(cValue0String)
	*param0 = param0String

	return toGoBool(ret)
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

func Fn_gtk_im_context_simple_new() unsafe.Pointer {
	ret := C.gtk_im_context_simple_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_im_context_simple_add_compose_file(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIMContextSimple)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_im_context_simple_add_compose_file(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_im_context_simple_add_table : has non-string array param data

func Fn_gtk_im_multicontext_new() unsafe.Pointer {
	ret := C.gtk_im_multicontext_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_im_multicontext_append_menuitems(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIMMulticontext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkMenuShell)(unsafe.Pointer(param0))

	C.gtk_im_multicontext_append_menuitems(cValueInstance, cValue0)
}

func Fn_gtk_im_multicontext_get_context_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIMMulticontext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_im_multicontext_get_context_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_im_multicontext_set_context_id(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIMMulticontext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_im_multicontext_set_context_id(cValueInstance, cValue0)
}

func Fn_gtk_icon_factory_new() unsafe.Pointer {
	ret := C.gtk_icon_factory_new()

	return unsafe.Pointer(ret)
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

func Fn_gtk_icon_factory_lookup(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkIconFactory)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_icon_factory_lookup(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_factory_remove_default(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconFactory)(unsafe.Pointer(paramInstance))

	C.gtk_icon_factory_remove_default(cValueInstance)
}

func Fn_gtk_icon_factory_lookup_default(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_icon_factory_lookup_default(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_info_new_for_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkIconTheme)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkPixbuf)(unsafe.Pointer(param1))

	ret := C.gtk_icon_info_new_for_pixbuf(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_info_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_info_free(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	C.gtk_icon_info_free(cValueInstance)
}

// UNSUPPORTED : gtk_icon_info_get_attach_points : has non-string array param points

func Fn_gtk_icon_info_get_base_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_get_base_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_info_get_builtin_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_get_builtin_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_info_get_display_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_get_display_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_icon_info_get_embedded_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	ret := C.gtk_icon_info_get_embedded_rect(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_icon_info_get_filename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_info_get_filename(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_icon_info_load_icon(paramInstance unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.gtk_icon_info_load_icon(cValueInstance, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_info_load_icon_async : has callback

func Fn_gtk_icon_info_load_symbolic(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 *bool, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRGBA)(unsafe.Pointer(param2))

	cValue3 := (*C.GdkRGBA)(unsafe.Pointer(param3))

	cValue4 := (*C.gboolean)(unsafe.Pointer(param4))

	cError := (**C.GError)(error)

	ret := C.gtk_icon_info_load_symbolic(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : has callback

func Fn_gtk_icon_info_load_symbolic_for_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *bool, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	cValue1 := (*C.gboolean)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.gtk_icon_info_load_symbolic_for_context(cValueInstance, cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : has callback

func Fn_gtk_icon_info_load_symbolic_for_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *bool, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	cValue1 := (C.GtkStateType)(param1)

	cValue2 := (*C.gboolean)(unsafe.Pointer(param2))

	cError := (**C.GError)(error)

	ret := C.gtk_icon_info_load_symbolic_for_style(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_info_set_raw_coordinates(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkIconInfo)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_icon_info_set_raw_coordinates(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_new() unsafe.Pointer {
	ret := C.gtk_icon_theme_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_append_search_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_append_search_path(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_choose_icon(paramInstance unsafe.Pointer, param0 []string, param1 int, param2 int) unsafe.Pointer {
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

	ret := C.gtk_icon_theme_choose_icon(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_get_example_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_theme_get_example_icon_name(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : has array return

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

func Fn_gtk_icon_theme_has_icon(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_icon_theme_has_icon(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_icon_theme_list_contexts(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_theme_list_contexts(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_list_icons(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_icon_theme_list_icons(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_load_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	cError := (**C.GError)(error)

	ret := C.gtk_icon_theme_load_icon(cValueInstance, cValue0, cValue1, cValue2, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_lookup_by_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	ret := C.gtk_icon_theme_lookup_by_gicon(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_lookup_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.GtkIconLookupFlags)(param2)

	ret := C.gtk_icon_theme_lookup_icon(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_prepend_search_path(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_icon_theme_prepend_search_path(cValueInstance, cValue0)
}

func Fn_gtk_icon_theme_rescan_if_needed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconTheme)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_theme_rescan_if_needed(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_icon_theme_get_default() unsafe.Pointer {
	ret := C.gtk_icon_theme_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_theme_get_for_screen(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	ret := C.gtk_icon_theme_get_for_screen(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_new() unsafe.Pointer {
	ret := C.gtk_icon_view_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_new_with_area(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_new_with_area(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_new_with_model(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_new_with_model(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_convert_widget_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	C.gtk_icon_view_convert_widget_to_bin_window_coords(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_icon_view_create_drag_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_create_drag_icon(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_view_enable_model_drag_dest : has non-string array param targets

// UNSUPPORTED : gtk_icon_view_enable_model_drag_source : has non-string array param targets

func Fn_gtk_icon_view_get_column_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_column_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_columns(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_columns(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_cursor(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkCellRenderer)(unsafe.Pointer(param1))

	ret := C.gtk_icon_view_get_cursor(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_get_dest_item_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *int) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkIconViewDropPosition)(unsafe.Pointer(param3))

	ret := C.gtk_icon_view_get_dest_item_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_get_drag_dest_item(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkIconViewDropPosition)(unsafe.Pointer(param1))

	C.gtk_icon_view_get_drag_dest_item(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_icon_view_get_item_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (**C.GtkCellRenderer)(unsafe.Pointer(param3))

	ret := C.gtk_icon_view_get_item_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_get_item_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_get_item_column(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_item_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_item_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_item_padding(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_item_padding(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_item_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_get_item_row(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_item_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_item_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_margin(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_margin(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_markup_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_markup_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_get_path_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_icon_view_get_path_at_pos(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_get_pixbuf_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_pixbuf_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_reorderable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_reorderable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_get_row_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_row_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_selected_items(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_selected_items(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_view_get_selection_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_selection_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_text_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_text_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_tooltip_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_icon_view_get_tooltip_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_icon_view_get_tooltip_context(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := (**C.GtkTreeModel)(unsafe.Pointer(param3))

	cValue4 := (**C.GtkTreePath)(unsafe.Pointer(param4))

	cValue5 := (*C.GtkTreeIter)(unsafe.Pointer(param5))

	ret := C.gtk_icon_view_get_tooltip_context(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_get_visible_range(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreePath)(unsafe.Pointer(param1))

	ret := C.gtk_icon_view_get_visible_range(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_icon_view_item_activated(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_icon_view_item_activated(cValueInstance, cValue0)
}

func Fn_gtk_icon_view_path_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkIconView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_icon_view_path_is_selected(cValueInstance, cValue0)

	return toGoBool(ret)
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

// UNSUPPORTED : gtk_icon_view_selected_foreach : has callback

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

func Fn_gtk_image_new() unsafe.Pointer {
	ret := C.gtk_image_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_animation(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkPixbufAnimation)(unsafe.Pointer(param0))

	ret := C.gtk_image_new_from_animation(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_file(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_image_new_from_file(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_gicon(param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_image_new_from_gicon(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_icon_name(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_image_new_from_icon_name(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_icon_set(param0 unsafe.Pointer, param1 int) unsafe.Pointer {
	cValue0 := (*C.GtkIconSet)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_image_new_from_icon_set(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_pixbuf(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	ret := C.gtk_image_new_from_pixbuf(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_stock(param0 string, param1 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_image_new_from_stock(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	C.gtk_image_clear(cValueInstance)
}

func Fn_gtk_image_get_animation(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_get_animation(cValueInstance)

	return unsafe.Pointer(ret)
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

func Fn_gtk_image_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_get_pixel_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_get_pixel_size(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_image_get_storage_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkImage)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_get_storage_type(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_image_menu_item_new() unsafe.Pointer {
	ret := C.gtk_image_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_new_from_stock(param0 string, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkAccelGroup)(unsafe.Pointer(param1))

	ret := C.gtk_image_menu_item_new_from_stock(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_new_with_label(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_image_menu_item_new_with_label(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_new_with_mnemonic(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_image_menu_item_new_with_mnemonic(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_get_always_show_image(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_menu_item_get_always_show_image(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_image_menu_item_get_image(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_menu_item_get_image(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_get_use_stock(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkImageMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_image_menu_item_get_use_stock(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_info_bar_new() unsafe.Pointer {
	ret := C.gtk_info_bar_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_info_bar_new_with_buttons : has varargs

func Fn_gtk_info_bar_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_info_bar_add_action_widget(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_info_bar_add_button(paramInstance unsafe.Pointer, param0 string, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_info_bar_add_button(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_info_bar_add_buttons : has varargs

func Fn_gtk_info_bar_get_action_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_info_bar_get_action_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_info_bar_get_content_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_info_bar_get_content_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_info_bar_get_message_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkInfoBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_info_bar_get_message_type(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_invisible_new() unsafe.Pointer {
	ret := C.gtk_invisible_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_invisible_new_for_screen(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	ret := C.gtk_invisible_new_for_screen(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_invisible_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkInvisible)(unsafe.Pointer(paramInstance))

	ret := C.gtk_invisible_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_invisible_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkInvisible)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_invisible_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_label_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_label_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_new_with_mnemonic(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_label_new_with_mnemonic(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_get_angle(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_angle(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_label_get_attributes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_attributes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_get_current_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_current_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_label_get_ellipsize(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_ellipsize(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_label_get_justify(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_justify(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_label_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_label_get_layout(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_layout(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_label_get_layout_offsets(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_label_get_line_wrap(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_line_wrap(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_label_get_line_wrap_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_line_wrap_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_label_get_max_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_max_width_chars(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_label_get_mnemonic_keyval(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_mnemonic_keyval(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_label_get_mnemonic_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_mnemonic_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_get_selectable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_selectable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_label_get_selection_bounds(paramInstance unsafe.Pointer, param0 *int, param1 *int) bool {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	ret := C.gtk_label_get_selection_bounds(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_label_get_single_line_mode(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_single_line_mode(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_label_get_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_label_get_track_visited_links(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_track_visited_links(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_label_get_use_markup(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_use_markup(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_label_get_use_underline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_use_underline(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_label_get_width_chars(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkLabel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_label_get_width_chars(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_layout_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	ret := C.gtk_layout_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_layout_get_bin_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	ret := C.gtk_layout_get_bin_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_layout_get_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	ret := C.gtk_layout_get_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_layout_get_size(paramInstance unsafe.Pointer, param0 *uint, param1 *uint) {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	C.gtk_layout_get_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_layout_get_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkLayout)(unsafe.Pointer(paramInstance))

	ret := C.gtk_layout_get_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
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

func Fn_gtk_link_button_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_link_button_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_link_button_new_with_label(param0 string, param1 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_link_button_new_with_label(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_link_button_get_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_link_button_get_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_link_button_get_visited(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkLinkButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_link_button_get_visited(cValueInstance)

	return toGoBool(ret)
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

// UNSUPPORTED : gtk_list_box_bind_model : has callback

// UNSUPPORTED : gtk_list_box_selected_foreach : has callback

// UNSUPPORTED : gtk_list_box_set_filter_func : has callback

// UNSUPPORTED : gtk_list_box_set_header_func : has callback

// UNSUPPORTED : gtk_list_box_set_sort_func : has callback

// UNSUPPORTED : gtk_list_store_new : has varargs

// UNSUPPORTED : gtk_list_store_newv : has non-string array param types

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

// UNSUPPORTED : gtk_list_store_insert_with_values : has varargs

// UNSUPPORTED : gtk_list_store_insert_with_valuesv : has non-string array param columns

func Fn_gtk_list_store_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_list_store_iter_is_valid(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_list_store_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_list_store_remove(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_list_store_reorder : has non-string array param new_order

// UNSUPPORTED : gtk_list_store_set : has varargs

// UNSUPPORTED : gtk_list_store_set_column_types : has non-string array param types

// UNSUPPORTED : gtk_list_store_set_valist : has va_list

func Fn_gtk_list_store_set_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_list_store_set_value(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_list_store_set_valuesv : has non-string array param columns

func Fn_gtk_list_store_swap(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkListStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_list_store_swap(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_menu_new() unsafe.Pointer {
	ret := C.gtk_menu_new()

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : gtk_menu_attach_to_widget : has callback

func Fn_gtk_menu_detach(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_detach(cValueInstance)
}

func Fn_gtk_menu_get_accel_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_accel_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_get_accel_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_accel_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_menu_get_active(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_active(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_get_attach_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_attach_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_get_monitor(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_monitor(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_menu_get_reserve_toggle_size(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_reserve_toggle_size(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_menu_get_tearoff_state(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_tearoff_state(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_menu_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_menu_popdown(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenu)(unsafe.Pointer(paramInstance))

	C.gtk_menu_popdown(cValueInstance)
}

// UNSUPPORTED : gtk_menu_popup : has callback

// UNSUPPORTED : gtk_menu_popup_for_device : has callback

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

func Fn_gtk_menu_get_for_attach_widget(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_menu_get_for_attach_widget(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_bar_new() unsafe.Pointer {
	ret := C.gtk_menu_bar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_bar_get_child_pack_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_bar_get_child_pack_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_menu_bar_get_pack_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkMenuBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_bar_get_pack_direction(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_menu_item_new() unsafe.Pointer {
	ret := C.gtk_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_item_new_with_label(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_menu_item_new_with_label(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_item_new_with_mnemonic(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_menu_item_new_with_mnemonic(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_item_activate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_activate(cValueInstance)
}

func Fn_gtk_menu_item_deselect(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	C.gtk_menu_item_deselect(cValueInstance)
}

func Fn_gtk_menu_item_get_accel_path(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_accel_path(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_menu_item_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_menu_item_get_reserve_indicator(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_reserve_indicator(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_menu_item_get_right_justified(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_right_justified(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_menu_item_get_submenu(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_submenu(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_item_get_use_underline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_item_get_use_underline(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_menu_item_set_reserve_indicator(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_menu_item_set_reserve_indicator(cValueInstance, cValue0)
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

func Fn_gtk_menu_shell_get_parent_shell(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_shell_get_parent_shell(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_shell_get_selected_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_shell_get_selected_item(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_shell_get_take_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMenuShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_shell_get_take_focus(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_menu_tool_button_new(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_menu_tool_button_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_tool_button_new_from_stock(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_menu_tool_button_new_from_stock(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_tool_button_get_menu(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMenuToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_menu_tool_button_get_menu(cValueInstance)

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : gtk_message_dialog_new : has varargs

// UNSUPPORTED : gtk_message_dialog_new_with_markup : has varargs

// UNSUPPORTED : gtk_message_dialog_format_secondary_markup : has varargs

// UNSUPPORTED : gtk_message_dialog_format_secondary_text : has varargs

func Fn_gtk_message_dialog_get_image(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_message_dialog_get_image(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_message_dialog_get_message_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMessageDialog)(unsafe.Pointer(paramInstance))

	ret := C.gtk_message_dialog_get_message_area(cValueInstance)

	return unsafe.Pointer(ret)
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

func Fn_gtk_mount_operation_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	ret := C.gtk_mount_operation_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_mount_operation_get_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_mount_operation_get_parent(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_mount_operation_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_mount_operation_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_mount_operation_is_showing(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkMountOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_mount_operation_is_showing(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_notebook_new() unsafe.Pointer {
	ret := C.gtk_notebook_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_append_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	ret := C.gtk_notebook_append_page(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_notebook_append_page_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkWidget)(unsafe.Pointer(param2))

	ret := C.gtk_notebook_append_page_menu(cValueInstance, cValue0, cValue1, cValue2)

	return (int)(ret)
}

func Fn_gtk_notebook_get_action_widget(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPackType)(param0)

	ret := C.gtk_notebook_get_action_widget(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_get_current_page(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_current_page(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_notebook_get_group_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_group_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_notebook_get_menu_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_notebook_get_menu_label(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_get_menu_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_notebook_get_menu_label_text(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_notebook_get_n_pages(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_n_pages(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_notebook_get_nth_page(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_notebook_get_nth_page(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_get_scrollable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_scrollable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_notebook_get_show_border(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_show_border(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_notebook_get_show_tabs(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_show_tabs(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_notebook_get_tab_detachable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_notebook_get_tab_detachable(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_notebook_get_tab_hborder(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_tab_hborder(cValueInstance)

	return (uint16)(ret)
}

func Fn_gtk_notebook_get_tab_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_notebook_get_tab_label(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_get_tab_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_notebook_get_tab_label_text(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_notebook_get_tab_pos(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_tab_pos(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_notebook_get_tab_reorderable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_notebook_get_tab_reorderable(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_notebook_get_tab_vborder(paramInstance unsafe.Pointer) uint16 {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	ret := C.gtk_notebook_get_tab_vborder(cValueInstance)

	return (uint16)(ret)
}

func Fn_gtk_notebook_insert_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	ret := C.gtk_notebook_insert_page(cValueInstance, cValue0, cValue1, cValue2)

	return (int)(ret)
}

func Fn_gtk_notebook_insert_page_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkWidget)(unsafe.Pointer(param2))

	cValue3 := (C.gint)(param3)

	ret := C.gtk_notebook_insert_page_menu(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return (int)(ret)
}

func Fn_gtk_notebook_next_page(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_next_page(cValueInstance)
}

func Fn_gtk_notebook_page_num(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_notebook_page_num(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_notebook_popup_disable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_popup_disable(cValueInstance)
}

func Fn_gtk_notebook_popup_enable(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_popup_enable(cValueInstance)
}

func Fn_gtk_notebook_prepend_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	ret := C.gtk_notebook_prepend_page(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_notebook_prepend_page_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) int {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkWidget)(unsafe.Pointer(param2))

	ret := C.gtk_notebook_prepend_page_menu(cValueInstance, cValue0, cValue1, cValue2)

	return (int)(ret)
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

func Fn_gtk_notebook_set_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkPackType)(param1)

	C.gtk_notebook_set_action_widget(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_notebook_set_current_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_notebook_set_current_page(cValueInstance, cValue0)
}

func Fn_gtk_notebook_set_group_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkNotebook)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_notebook_set_group_name(cValueInstance, cValue0)
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

func Fn_gtk_notebook_page_accessible_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkNotebookAccessible)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkWidget)(unsafe.Pointer(param1))

	ret := C.gtk_notebook_page_accessible_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_page_accessible_invalidate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkNotebookPageAccessible)(unsafe.Pointer(paramInstance))

	C.gtk_notebook_page_accessible_invalidate(cValueInstance)
}

func Fn_gtk_numerable_icon_get_background_gicon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_background_gicon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_numerable_icon_get_background_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_background_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_numerable_icon_get_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_count(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_numerable_icon_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_numerable_icon_get_style_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_numerable_icon_get_style_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_numerable_icon_set_background_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	C.gtk_numerable_icon_set_background_gicon(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_set_background_icon_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_numerable_icon_set_background_icon_name(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_set_count(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_numerable_icon_set_count(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_numerable_icon_set_label(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_set_style_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkNumerableIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleContext)(unsafe.Pointer(param0))

	C.gtk_numerable_icon_set_style_context(cValueInstance, cValue0)
}

func Fn_gtk_numerable_icon_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	ret := C.gtk_numerable_icon_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_numerable_icon_new_with_style_context(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkStyleContext)(unsafe.Pointer(param1))

	ret := C.gtk_numerable_icon_new_with_style_context(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_offscreen_window_new() unsafe.Pointer {
	ret := C.gtk_offscreen_window_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_offscreen_window_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkOffscreenWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_offscreen_window_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_offscreen_window_get_surface(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkOffscreenWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_offscreen_window_get_surface(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_pad_controller_set_action_entries : has non-string array param entries

func Fn_gtk_page_setup_new() unsafe.Pointer {
	ret := C.gtk_page_setup_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_new_from_file(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_new_from_file(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_new_from_key_file(param0 unsafe.Pointer, param1 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_new_from_key_file(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_page_setup_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_get_bottom_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_bottom_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_left_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_left_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_page_setup_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_page_setup_get_page_height(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_page_height(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_page_width(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_page_width(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_paper_height(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_paper_height(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_paper_size(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_page_setup_get_paper_size(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_page_setup_get_paper_width(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_paper_width(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_right_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_right_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_get_top_margin(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_page_setup_get_top_margin(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_page_setup_load_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_load_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_page_setup_load_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_load_key_file(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
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

func Fn_gtk_page_setup_to_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_page_setup_to_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_page_setup_to_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
	cValueInstance := (*C.GtkPageSetup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	C.gtk_page_setup_to_key_file(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_paned_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	ret := C.gtk_paned_new(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_gtk_paned_get_child1(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paned_get_child1(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paned_get_child2(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paned_get_child2(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paned_get_handle_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paned_get_handle_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_paned_get_position(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPaned)(unsafe.Pointer(paramInstance))

	ret := C.gtk_paned_get_position(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_places_sidebar_get_show_connect_to_server(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_get_show_connect_to_server(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_plug_new(param0 uint64) unsafe.Pointer {
	cValue0 := (C.Window)(param0)

	ret := C.gtk_plug_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_plug_new_for_display(param0 unsafe.Pointer, param1 uint64) unsafe.Pointer {
	cValue0 := (*C.GdkDisplay)(unsafe.Pointer(param0))

	cValue1 := (C.Window)(param1)

	ret := C.gtk_plug_new_for_display(cValue0, cValue1)

	return unsafe.Pointer(ret)
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

func Fn_gtk_plug_get_embedded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	ret := C.gtk_plug_get_embedded(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_plug_get_id(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	ret := C.gtk_plug_get_id(cValueInstance)

	return (uint64)(ret)
}

func Fn_gtk_plug_get_socket_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPlug)(unsafe.Pointer(paramInstance))

	ret := C.gtk_plug_get_socket_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_popover_get_pointing_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	ret := C.gtk_popover_get_pointing_to(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_popover_get_position(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	ret := C.gtk_popover_get_position(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_context_create_pango_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_create_pango_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_create_pango_layout(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_create_pango_layout(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_get_cairo_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_cairo_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_get_dpi_x(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_dpi_x(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_context_get_dpi_y(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_dpi_y(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_context_get_hard_margins(paramInstance unsafe.Pointer, param0 *float64, param1 *float64, param2 *float64, param3 *float64) bool {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	cValue2 := (*C.gdouble)(unsafe.Pointer(param2))

	cValue3 := (*C.gdouble)(unsafe.Pointer(param3))

	ret := C.gtk_print_context_get_hard_margins(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_print_context_get_height(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_height(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_context_get_page_setup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_page_setup(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_get_pango_fontmap(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_pango_fontmap(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_context_get_width(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_context_get_width(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_context_set_cairo_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 float64) {
	cValueInstance := (*C.GtkPrintContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	cValue1 := (C.double)(param1)

	cValue2 := (C.double)(param2)

	C.gtk_print_context_set_cairo_context(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_print_operation_new() unsafe.Pointer {
	ret := C.gtk_print_operation_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_operation_cancel(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_cancel(cValueInstance)
}

func Fn_gtk_print_operation_draw_page_finish(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_draw_page_finish(cValueInstance)
}

func Fn_gtk_print_operation_get_default_page_setup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_default_page_setup(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_operation_get_embed_page_setup(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_embed_page_setup(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_get_error(paramInstance unsafe.Pointer, error unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	C.gtk_print_operation_get_error(cValueInstance, cError)
}

func Fn_gtk_print_operation_get_has_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_has_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_get_n_pages_to_print(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_n_pages_to_print(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_operation_get_print_settings(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_print_settings(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_operation_get_status(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_status(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_operation_get_status_string(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_status_string(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_operation_get_support_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_get_support_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_is_finished(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_operation_is_finished(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_run(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintOperation)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkPrintOperationAction)(param0)

	cValue1 := (*C.GtkWindow)(unsafe.Pointer(param1))

	cError := (**C.GError)(error)

	ret := C.gtk_print_operation_run(cValueInstance, cValue0, cValue1, cError)

	return (int)(ret)
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

func Fn_gtk_print_settings_new() unsafe.Pointer {
	ret := C.gtk_print_settings_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_settings_new_from_file(param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_new_from_file(cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_settings_new_from_key_file(param0 unsafe.Pointer, param1 string, error unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_new_from_key_file(cValue0, cValue1, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_settings_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : has callback

func Fn_gtk_print_settings_get(paramInstance unsafe.Pointer, param0 string) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_get(cValueInstance, cValue0)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_bool(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_get_bool(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_get_collate(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_collate(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_get_default_source(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_default_source(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_dither(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_dither(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_double(paramInstance unsafe.Pointer, param0 string) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_get_double(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_double_with_default(paramInstance unsafe.Pointer, param0 string, param1 float64) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gdouble)(param1)

	ret := C.gtk_print_settings_get_double_with_default(cValueInstance, cValue0, cValue1)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_duplex(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_duplex(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_finishings(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_finishings(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_int(paramInstance unsafe.Pointer, param0 string) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_get_int(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_int_with_default(paramInstance unsafe.Pointer, param0 string, param1 int) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_print_settings_get_int_with_default(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_length(paramInstance unsafe.Pointer, param0 string, param1 int) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkUnit)(param1)

	ret := C.gtk_print_settings_get_length(cValueInstance, cValue0, cValue1)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_media_type(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_media_type(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_n_copies(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_n_copies(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_number_up(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_number_up(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_number_up_layout(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_number_up_layout(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_output_bin(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_output_bin(cValueInstance)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_print_settings_get_page_ranges : has array return

func Fn_gtk_print_settings_get_page_set(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_page_set(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_paper_height(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_print_settings_get_paper_height(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_paper_size(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_paper_size(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_print_settings_get_paper_width(paramInstance unsafe.Pointer, param0 int) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUnit)(param0)

	ret := C.gtk_print_settings_get_paper_width(cValueInstance, cValue0)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_print_pages(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_print_pages(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_printer(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_printer(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_print_settings_get_printer_lpi(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_printer_lpi(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_quality(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_quality(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_resolution(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_resolution(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_resolution_x(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_resolution_x(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_resolution_y(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_resolution_y(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_print_settings_get_reverse(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_reverse(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_get_scale(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_scale(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_print_settings_get_use_color(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	ret := C.gtk_print_settings_get_use_color(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_has_key(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_print_settings_has_key(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_load_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_load_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_print_settings_load_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GKeyFile)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_load_key_file(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
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

// UNSUPPORTED : gtk_print_settings_set_page_ranges : has non-string array param page_ranges

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

func Fn_gtk_print_settings_to_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPrintSettings)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_print_settings_to_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
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

func Fn_gtk_progress_bar_new() unsafe.Pointer {
	ret := C.gtk_progress_bar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_progress_bar_get_ellipsize(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_progress_bar_get_ellipsize(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_progress_bar_get_fraction(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_progress_bar_get_fraction(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_progress_bar_get_inverted(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_progress_bar_get_inverted(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_progress_bar_get_pulse_step(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_progress_bar_get_pulse_step(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_progress_bar_get_show_text(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_progress_bar_get_show_text(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_progress_bar_get_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_progress_bar_get_text(cValueInstance)

	return C.GoString(ret)
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

func Fn_gtk_progress_bar_set_show_text(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_progress_bar_set_show_text(cValueInstance, cValue0)
}

func Fn_gtk_progress_bar_set_text(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkProgressBar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_progress_bar_set_text(cValueInstance, cValue0)
}

func Fn_gtk_radio_action_new(param0 string, param1 string, param2 string, param3 string, param4 int) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (C.gint)(param4)

	ret := C.gtk_radio_action_new(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_action_get_current_value(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_radio_action_get_current_value(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_radio_action_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_radio_action_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_action_join_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioAction)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRadioAction)(unsafe.Pointer(param0))

	C.gtk_radio_action_join_group(cValueInstance, cValue0)
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

func Fn_gtk_radio_button_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	ret := C.gtk_radio_button_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_from_widget(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkRadioButton)(unsafe.Pointer(param0))

	ret := C.gtk_radio_button_new_from_widget(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_with_label(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_button_new_with_label(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkRadioButton)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_button_new_with_label_from_widget(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_with_mnemonic(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_button_new_with_mnemonic(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkRadioButton)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_button_new_with_mnemonic_from_widget(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRadioButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_radio_button_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_join_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRadioButton)(unsafe.Pointer(param0))

	C.gtk_radio_button_join_group(cValueInstance, cValue0)
}

func Fn_gtk_radio_button_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_button_set_group(cValueInstance, cValue0)
}

func Fn_gtk_radio_menu_item_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	ret := C.gtk_radio_menu_item_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_new_from_widget(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	ret := C.gtk_radio_menu_item_new_from_widget(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_new_with_label(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_menu_item_new_with_label(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_menu_item_new_with_label_from_widget(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_new_with_mnemonic(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_menu_item_new_with_mnemonic(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRadioMenuItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_radio_menu_item_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_menu_item_set_group(cValueInstance, cValue0)
}

func Fn_gtk_radio_tool_button_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	ret := C.gtk_radio_tool_button_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_new_from_stock(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_tool_button_new_from_stock(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_new_from_widget(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkRadioToolButton)(unsafe.Pointer(param0))

	ret := C.gtk_radio_tool_button_new_from_widget(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_new_with_stock_from_widget(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkRadioToolButton)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_radio_tool_button_new_with_stock_from_widget(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRadioToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_radio_tool_button_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_tool_button_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GSList)(unsafe.Pointer(param0))

	C.gtk_radio_tool_button_set_group(cValueInstance, cValue0)
}

func Fn_gtk_range_get_adjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_adjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_range_get_fill_level(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_fill_level(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_range_get_flippable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_flippable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_range_get_inverted(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_inverted(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_range_get_lower_stepper_sensitivity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_lower_stepper_sensitivity(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_range_get_min_slider_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_min_slider_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_range_get_range_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_range_get_range_rect(cValueInstance, cValue0)
}

func Fn_gtk_range_get_restrict_to_fill_level(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_restrict_to_fill_level(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_range_get_round_digits(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_round_digits(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_range_get_show_fill_level(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_show_fill_level(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_range_get_slider_range(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_range_get_slider_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_range_get_slider_size_fixed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_slider_size_fixed(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_range_get_upper_stepper_sensitivity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_upper_stepper_sensitivity(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_range_get_value(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	ret := C.gtk_range_get_value(cValueInstance)

	return (float64)(ret)
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

func Fn_gtk_range_set_min_slider_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_range_set_min_slider_size(cValueInstance, cValue0)
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

func Fn_gtk_range_set_round_digits(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_range_set_round_digits(cValueInstance, cValue0)
}

func Fn_gtk_range_set_show_fill_level(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_show_fill_level(cValueInstance, cValue0)
}

func Fn_gtk_range_set_slider_size_fixed(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRange)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_range_set_slider_size_fixed(cValueInstance, cValue0)
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

func Fn_gtk_rc_style_new() unsafe.Pointer {
	ret := C.gtk_rc_style_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_rc_style_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRcStyle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_rc_style_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_action_new(param0 string, param1 string, param2 string, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.gtk_recent_action_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_action_new_for_manager(param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	cValue4 := (*C.GtkRecentManager)(unsafe.Pointer(param4))

	ret := C.gtk_recent_action_new_for_manager(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_action_get_show_numbers(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_action_get_show_numbers(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_action_set_show_numbers(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentAction)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_action_set_show_numbers(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_chooser_dialog_new : has varargs

// UNSUPPORTED : gtk_recent_chooser_dialog_new_for_manager : has varargs

func Fn_gtk_recent_chooser_menu_new() unsafe.Pointer {
	ret := C.gtk_recent_chooser_menu_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_menu_new_for_manager(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkRecentManager)(unsafe.Pointer(param0))

	ret := C.gtk_recent_chooser_menu_new_for_manager(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_menu_get_show_numbers(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooserMenu)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_menu_get_show_numbers(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_menu_set_show_numbers(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooserMenu)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_menu_set_show_numbers(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_widget_new() unsafe.Pointer {
	ret := C.gtk_recent_chooser_widget_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_widget_new_for_manager(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkRecentManager)(unsafe.Pointer(param0))

	ret := C.gtk_recent_chooser_widget_new_for_manager(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_filter_new() unsafe.Pointer {
	ret := C.gtk_recent_filter_new()

	return unsafe.Pointer(ret)
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

// UNSUPPORTED : gtk_recent_filter_add_custom : has callback

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

func Fn_gtk_recent_filter_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentFilterInfo)(unsafe.Pointer(param0))

	ret := C.gtk_recent_filter_filter(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_filter_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_filter_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_filter_get_needed(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_filter_get_needed(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_recent_filter_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_filter_set_name(cValueInstance, cValue0)
}

func Fn_gtk_recent_manager_new() unsafe.Pointer {
	ret := C.gtk_recent_manager_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_manager_add_full(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkRecentData)(unsafe.Pointer(param1))

	ret := C.gtk_recent_manager_add_full(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_add_item(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_recent_manager_add_item(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_get_items(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_manager_get_items(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_manager_has_item(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_recent_manager_has_item(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_lookup_item(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_manager_lookup_item(cValueInstance, cValue0, cError)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_manager_move_item(paramInstance unsafe.Pointer, param0 string, param1 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_manager_move_item(cValueInstance, cValue0, cValue1, cError)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_purge_items(paramInstance unsafe.Pointer, error unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_manager_purge_items(cValueInstance, cError)

	return (int)(ret)
}

func Fn_gtk_recent_manager_remove_item(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_manager_remove_item(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_recent_manager_get_default() unsafe.Pointer {
	ret := C.gtk_recent_manager_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gtk_renderer_cell_accessible_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	ret := C.gtk_renderer_cell_accessible_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_new(param0 int, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	ret := C.gtk_scale_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_new_with_range(param0 int, param1 float64, param2 float64, param3 float64) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	cValue3 := (C.gdouble)(param3)

	ret := C.gtk_scale_new_with_range(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
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

func Fn_gtk_scale_get_digits(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_get_digits(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scale_get_draw_value(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_get_draw_value(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_scale_get_layout(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_get_layout(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_scale_get_layout_offsets(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_scale_get_value_pos(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScale)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_get_value_pos(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_scale_button_new(param0 int, param1 float64, param2 float64, param3 float64, param4 []string) unsafe.Pointer {
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

	ret := C.gtk_scale_button_new(cValue0, cValue1, cValue2, cValue3, cValue4)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_button_get_adjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_adjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_button_get_minus_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_minus_button(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_button_get_plus_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_plus_button(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_button_get_popup(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_popup(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_button_get_value(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkScaleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scale_button_get_value(cValueInstance)

	return (float64)(ret)
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

func Fn_gtk_scrollbar_new(param0 int, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	ret := C.gtk_scrollbar_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	ret := C.gtk_scrolled_window_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_add_with_viewport(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_scrolled_window_add_with_viewport(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_get_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_get_hscrollbar(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_hscrollbar(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_get_min_content_height(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_min_content_height(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_get_min_content_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_min_content_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_get_placement(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_placement(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_get_policy(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkPolicyType)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkPolicyType)(unsafe.Pointer(param1))

	C.gtk_scrolled_window_get_policy(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_scrolled_window_get_shadow_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_shadow_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_get_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_get_vscrollbar(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrolled_window_get_vscrollbar(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_scrolled_window_set_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_min_content_height(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_scrolled_window_set_min_content_height(cValueInstance, cValue0)
}

func Fn_gtk_scrolled_window_set_min_content_width(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrolledWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_scrolled_window_set_min_content_width(cValueInstance, cValue0)
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

func Fn_gtk_separator_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkOrientation)(param0)

	ret := C.gtk_separator_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_separator_menu_item_new() unsafe.Pointer {
	ret := C.gtk_separator_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_separator_tool_item_new() unsafe.Pointer {
	ret := C.gtk_separator_tool_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_separator_tool_item_get_draw(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSeparatorToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_separator_tool_item_get_draw(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_settings_get_default() unsafe.Pointer {
	ret := C.gtk_settings_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gtk_settings_get_for_screen(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	ret := C.gtk_settings_get_for_screen(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_settings_install_property(param0 unsafe.Pointer) {
	cValue0 := (*C.GParamSpec)(unsafe.Pointer(param0))

	C.gtk_settings_install_property(cValue0)
}

// UNSUPPORTED : gtk_settings_install_property_parser : has callback

func Fn_gtk_size_group_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkSizeGroupMode)(param0)

	ret := C.gtk_size_group_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_size_group_add_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_size_group_add_widget(cValueInstance, cValue0)
}

func Fn_gtk_size_group_get_ignore_hidden(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_size_group_get_ignore_hidden(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_size_group_get_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_size_group_get_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_size_group_get_widgets(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSizeGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_size_group_get_widgets(cValueInstance)

	return unsafe.Pointer(ret)
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

func Fn_gtk_socket_new() unsafe.Pointer {
	ret := C.gtk_socket_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_socket_add_id(paramInstance unsafe.Pointer, param0 uint64) {
	cValueInstance := (*C.GtkSocket)(unsafe.Pointer(paramInstance))

	cValue0 := (C.Window)(param0)

	C.gtk_socket_add_id(cValueInstance, cValue0)
}

func Fn_gtk_socket_get_id(paramInstance unsafe.Pointer) uint64 {
	cValueInstance := (*C.GtkSocket)(unsafe.Pointer(paramInstance))

	ret := C.gtk_socket_get_id(cValueInstance)

	return (uint64)(ret)
}

func Fn_gtk_socket_get_plug_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSocket)(unsafe.Pointer(paramInstance))

	ret := C.gtk_socket_get_plug_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_spin_button_new(param0 unsafe.Pointer, param1 float64, param2 uint) unsafe.Pointer {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.guint)(param2)

	ret := C.gtk_spin_button_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_spin_button_new_with_range(param0 float64, param1 float64, param2 float64) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	ret := C.gtk_spin_button_new_with_range(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_spin_button_configure(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 uint) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.guint)(param2)

	C.gtk_spin_button_configure(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_spin_button_get_adjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_spin_button_get_adjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_spin_button_get_digits(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_spin_button_get_digits(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_spin_button_get_increments(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	C.gtk_spin_button_get_increments(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_spin_button_get_numeric(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_spin_button_get_numeric(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_spin_button_get_range(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gdouble)(unsafe.Pointer(param0))

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	C.gtk_spin_button_get_range(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_spin_button_get_snap_to_ticks(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_spin_button_get_snap_to_ticks(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_spin_button_get_update_policy(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_spin_button_get_update_policy(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_spin_button_get_value(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_spin_button_get_value(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_spin_button_get_value_as_int(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_spin_button_get_value_as_int(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_spin_button_get_wrap(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSpinButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_spin_button_get_wrap(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_spinner_new() unsafe.Pointer {
	ret := C.gtk_spinner_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_spinner_start(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinner)(unsafe.Pointer(paramInstance))

	C.gtk_spinner_start(cValueInstance)
}

func Fn_gtk_spinner_stop(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkSpinner)(unsafe.Pointer(paramInstance))

	C.gtk_spinner_stop(cValueInstance)
}

func Fn_gtk_status_icon_new() unsafe.Pointer {
	ret := C.gtk_status_icon_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_file(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_status_icon_new_from_file(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_gicon(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	ret := C.gtk_status_icon_new_from_gicon(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_icon_name(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_status_icon_new_from_icon_name(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_pixbuf(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	ret := C.gtk_status_icon_new_from_pixbuf(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_new_from_stock(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_status_icon_new_from_stock(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_get_geometry(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, param2 *int) bool {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkOrientation)(unsafe.Pointer(param2))

	ret := C.gtk_status_icon_get_geometry(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_status_icon_get_gicon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_gicon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_get_has_tooltip(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_has_tooltip(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_status_icon_get_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_get_pixbuf(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_pixbuf(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_status_icon_get_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_status_icon_get_stock(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_stock(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_get_storage_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_storage_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_status_icon_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_get_tooltip_markup(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_tooltip_markup(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_get_tooltip_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_tooltip_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_status_icon_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_status_icon_get_x11_window_id(paramInstance unsafe.Pointer) uint32 {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_get_x11_window_id(cValueInstance)

	return (uint32)(ret)
}

func Fn_gtk_status_icon_is_embedded(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	ret := C.gtk_status_icon_is_embedded(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_status_icon_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStatusIcon)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_status_icon_set_name(cValueInstance, cValue0)
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

func Fn_gtk_statusbar_new() unsafe.Pointer {
	ret := C.gtk_statusbar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_statusbar_get_context_id(paramInstance unsafe.Pointer, param0 string) uint {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_statusbar_get_context_id(cValueInstance, cValue0)

	return (uint)(ret)
}

func Fn_gtk_statusbar_get_message_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_statusbar_get_message_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_statusbar_pop(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_statusbar_pop(cValueInstance, cValue0)
}

func Fn_gtk_statusbar_push(paramInstance unsafe.Pointer, param0 uint, param1 string) uint {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_statusbar_push(cValueInstance, cValue0, cValue1)

	return (uint)(ret)
}

func Fn_gtk_statusbar_remove(paramInstance unsafe.Pointer, param0 uint, param1 uint) {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	C.gtk_statusbar_remove(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_statusbar_remove_all(paramInstance unsafe.Pointer, param0 uint) {
	cValueInstance := (*C.GtkStatusbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	C.gtk_statusbar_remove_all(cValueInstance, cValue0)
}

func Fn_gtk_style_new() unsafe.Pointer {
	ret := C.gtk_style_new()

	return unsafe.Pointer(ret)
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

func Fn_gtk_style_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	ret := C.gtk_style_attach(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_copy(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_copy(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_detach(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	C.gtk_style_detach(cValueInstance)
}

// UNSUPPORTED : gtk_style_get : has varargs

func Fn_gtk_style_get_style_property(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_style_get_style_property(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_style_get_valist : has va_list

func Fn_gtk_style_has_context(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_has_context(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_style_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GdkColor)(unsafe.Pointer(param1))

	ret := C.gtk_style_lookup_color(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_style_lookup_icon_set(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_style_lookup_icon_set(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_render_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 unsafe.Pointer, param5 string) unsafe.Pointer {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkIconSource)(unsafe.Pointer(param0))

	cValue1 := (C.GtkTextDirection)(param1)

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := (C.GtkIconSize)(param3)

	cValue4 := (*C.GtkWidget)(unsafe.Pointer(param4))

	cValue5 := (*C.gchar)(C.CString(param5))
	defer C.free(unsafe.Pointer(cValue5))

	ret := C.gtk_style_render_icon(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_set_background(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkStyle)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.GtkStateType)(param1)

	C.gtk_style_set_background(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_new() unsafe.Pointer {
	ret := C.gtk_style_context_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_add_class(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_style_context_add_class(cValueInstance, cValue0)
}

func Fn_gtk_style_context_add_provider(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProvider)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	C.gtk_style_context_add_provider(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_add_region(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkRegionFlags)(param1)

	C.gtk_style_context_add_region(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_cancel_animations(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.gtk_style_context_cancel_animations(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_style_context_get : has varargs

func Fn_gtk_style_context_get_background_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_style_context_get_background_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_border(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_style_context_get_border(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_border_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_style_context_get_border_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_style_context_get_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_style_context_get_font(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	ret := C.gtk_style_context_get_font(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_get_junction_sides(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_junction_sides(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_style_context_get_margin(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_style_context_get_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_padding(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_style_context_get_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_context_get_path(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_path(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_get_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_style_context_get_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_style_context_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_get_section(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_style_context_get_section(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_get_state(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_get_state(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_style_context_get_style : has varargs

func Fn_gtk_style_context_get_style_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.gtk_style_context_get_style_property(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_style_context_get_style_valist : has va_list

// UNSUPPORTED : gtk_style_context_get_valist : has va_list

func Fn_gtk_style_context_has_class(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_style_context_has_class(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_style_context_has_region(paramInstance unsafe.Pointer, param0 string, param1 *int) bool {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkRegionFlags)(unsafe.Pointer(param1))

	ret := C.gtk_style_context_has_region(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_style_context_invalidate(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.gtk_style_context_invalidate(cValueInstance)
}

func Fn_gtk_style_context_list_classes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_list_classes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_list_regions(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	ret := C.gtk_style_context_list_regions(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	ret := C.gtk_style_context_lookup_color(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_style_context_lookup_icon_set(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_style_context_lookup_icon_set(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_notify_state_change(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 bool) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.gpointer)(param1)

	cValue2 := (C.GtkStateType)(param2)

	cValue3 := toCBool(param3)

	C.gtk_style_context_notify_state_change(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_style_context_pop_animatable_region(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.gtk_style_context_pop_animatable_region(cValueInstance)
}

func Fn_gtk_style_context_push_animatable_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gpointer)(param0)

	C.gtk_style_context_push_animatable_region(cValueInstance, cValue0)
}

func Fn_gtk_style_context_remove_class(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_style_context_remove_class(cValueInstance, cValue0)
}

func Fn_gtk_style_context_remove_provider(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProvider)(unsafe.Pointer(param0))

	C.gtk_style_context_remove_provider(cValueInstance, cValue0)
}

func Fn_gtk_style_context_remove_region(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_style_context_remove_region(cValueInstance, cValue0)
}

func Fn_gtk_style_context_restore(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.gtk_style_context_restore(cValueInstance)
}

func Fn_gtk_style_context_save(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	C.gtk_style_context_save(cValueInstance)
}

func Fn_gtk_style_context_scroll_animations(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_style_context_scroll_animations(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_style_context_set_background(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	C.gtk_style_context_set_background(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_direction(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextDirection)(param0)

	C.gtk_style_context_set_direction(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_junction_sides(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkJunctionSides)(param0)

	C.gtk_style_context_set_junction_sides(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidgetPath)(unsafe.Pointer(param0))

	C.gtk_style_context_set_path(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_style_context_set_screen(cValueInstance, cValue0)
}

func Fn_gtk_style_context_set_state(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.gtk_style_context_set_state(cValueInstance, cValue0)
}

func Fn_gtk_style_context_state_is_running(paramInstance unsafe.Pointer, param0 int, param1 *float64) bool {
	cValueInstance := (*C.GtkStyleContext)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gtk_style_context_state_is_running(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_style_context_add_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkStyleProvider)(unsafe.Pointer(param1))

	cValue2 := (C.guint)(param2)

	C.gtk_style_context_add_provider_for_screen(cValue0, cValue1, cValue2)
}

func Fn_gtk_style_context_remove_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkStyleProvider)(unsafe.Pointer(param1))

	C.gtk_style_context_remove_provider_for_screen(cValue0, cValue1)
}

func Fn_gtk_style_context_reset_widgets(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	C.gtk_style_context_reset_widgets(cValue0)
}

func Fn_gtk_style_properties_new() unsafe.Pointer {
	ret := C.gtk_style_properties_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_properties_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	C.gtk_style_properties_clear(cValueInstance)
}

// UNSUPPORTED : gtk_style_properties_get : has varargs

func Fn_gtk_style_properties_get_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	ret := C.gtk_style_properties_get_property(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_style_properties_get_valist : has va_list

func Fn_gtk_style_properties_lookup_color(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_style_properties_lookup_color(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_properties_map_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkSymbolicColor)(unsafe.Pointer(param1))

	C.gtk_style_properties_map_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_style_properties_merge(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyleProperties)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_style_properties_merge(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_style_properties_set : has varargs

func Fn_gtk_style_properties_set_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_style_properties_set_property(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_style_properties_set_valist : has va_list

func Fn_gtk_style_properties_unset_property(paramInstance unsafe.Pointer, param0 string, param1 int) {
	cValueInstance := (*C.GtkStyleProperties)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	C.gtk_style_properties_unset_property(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : has callback

// UNSUPPORTED : gtk_style_properties_register_property : has callback

func Fn_gtk_switch_new() unsafe.Pointer {
	ret := C.gtk_switch_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_switch_get_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkSwitch)(unsafe.Pointer(paramInstance))

	ret := C.gtk_switch_get_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_switch_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkSwitch)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_switch_set_active(cValueInstance, cValue0)
}

func Fn_gtk_table_new(param0 uint, param1 uint, param2 bool) unsafe.Pointer {
	cValue0 := (C.guint)(param0)

	cValue1 := (C.guint)(param1)

	cValue2 := toCBool(param2)

	ret := C.gtk_table_new(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
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

func Fn_gtk_table_get_col_spacing(paramInstance unsafe.Pointer, param0 uint) uint {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gtk_table_get_col_spacing(cValueInstance, cValue0)

	return (uint)(ret)
}

func Fn_gtk_table_get_default_col_spacing(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_table_get_default_col_spacing(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_table_get_default_row_spacing(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_table_get_default_row_spacing(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_table_get_homogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_table_get_homogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_table_get_row_spacing(paramInstance unsafe.Pointer, param0 uint) uint {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gtk_table_get_row_spacing(cValueInstance, cValue0)

	return (uint)(ret)
}

func Fn_gtk_table_get_size(paramInstance unsafe.Pointer, param0 *uint, param1 *uint) {
	cValueInstance := (*C.GtkTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.guint)(unsafe.Pointer(param0))

	cValue1 := (*C.guint)(unsafe.Pointer(param1))

	C.gtk_table_get_size(cValueInstance, cValue0, cValue1)
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

func Fn_gtk_tearoff_menu_item_new() unsafe.Pointer {
	ret := C.gtk_tearoff_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkTextTagTable)(unsafe.Pointer(param0))

	ret := C.gtk_text_buffer_new(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_gtk_text_buffer_backspace(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	cValue2 := toCBool(param2)

	ret := C.gtk_text_buffer_backspace(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
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

func Fn_gtk_text_buffer_create_child_anchor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_buffer_create_child_anchor(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_create_mark(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 bool) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	ret := C.gtk_text_buffer_create_mark(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_text_buffer_create_tag : has varargs

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

func Fn_gtk_text_buffer_delete_interactive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	ret := C.gtk_text_buffer_delete_interactive(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
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

func Fn_gtk_text_buffer_delete_selection(paramInstance unsafe.Pointer, param0 bool, param1 bool) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	cValue1 := toCBool(param1)

	ret := C.gtk_text_buffer_delete_selection(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_buffer_deserialize : has non-string array param data

func Fn_gtk_text_buffer_deserialize_get_can_create_tags(paramInstance unsafe.Pointer, param0 gdk.Atom) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	ret := C.gtk_text_buffer_deserialize_get_can_create_tags(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_text_buffer_get_char_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_char_count(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_buffer_get_copy_target_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_copy_target_list(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_text_buffer_get_deserialize_formats : has array return

func Fn_gtk_text_buffer_get_end_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_buffer_get_end_iter(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_get_has_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_has_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_get_insert(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_insert(cValueInstance)

	return unsafe.Pointer(ret)
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

func Fn_gtk_text_buffer_get_line_count(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_line_count(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_buffer_get_mark(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_text_buffer_get_mark(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_modified(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_modified(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_get_paste_target_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_paste_target_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_selection_bound(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_selection_bound(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_selection_bounds(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	ret := C.gtk_text_buffer_get_selection_bounds(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_buffer_get_serialize_formats : has array return

func Fn_gtk_text_buffer_get_slice(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) string {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	ret := C.gtk_text_buffer_get_slice(cValueInstance, cValue0, cValue1, cValue2)

	return C.GoString(ret)
}

func Fn_gtk_text_buffer_get_start_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	C.gtk_text_buffer_get_start_iter(cValueInstance, cValue0)
}

func Fn_gtk_text_buffer_get_tag_table(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_buffer_get_tag_table(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) string {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	ret := C.gtk_text_buffer_get_text(cValueInstance, cValue0, cValue1, cValue2)

	return C.GoString(ret)
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

func Fn_gtk_text_buffer_insert_interactive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int, param3 bool) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	cValue3 := toCBool(param3)

	ret := C.gtk_text_buffer_insert_interactive(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_insert_interactive_at_cursor(paramInstance unsafe.Pointer, param0 string, param1 int, param2 bool) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := toCBool(param2)

	ret := C.gtk_text_buffer_insert_interactive_at_cursor(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
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

func Fn_gtk_text_buffer_insert_range_interactive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) bool {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTextIter)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	cValue3 := toCBool(param3)

	ret := C.gtk_text_buffer_insert_range_interactive(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_buffer_insert_with_tags : has varargs

// UNSUPPORTED : gtk_text_buffer_insert_with_tags_by_name : has varargs

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

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : has callback

func Fn_gtk_text_buffer_register_deserialize_tagset(paramInstance unsafe.Pointer, param0 string) gdk.Atom {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_text_buffer_register_deserialize_tagset(cValueInstance, cValue0)

	return gdk.Atom(unsafe.Pointer(ret))
}

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : has callback

func Fn_gtk_text_buffer_register_serialize_tagset(paramInstance unsafe.Pointer, param0 string) gdk.Atom {
	cValueInstance := (*C.GtkTextBuffer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_text_buffer_register_serialize_tagset(cValueInstance, cValue0)

	return gdk.Atom(unsafe.Pointer(ret))
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

// UNSUPPORTED : gtk_text_buffer_serialize : has array return

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

func Fn_gtk_text_child_anchor_new() unsafe.Pointer {
	ret := C.gtk_text_child_anchor_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_child_anchor_get_deleted(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextChildAnchor)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_child_anchor_get_deleted(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_child_anchor_get_widgets(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextChildAnchor)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_child_anchor_get_widgets(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_mark_new(param0 string, param1 bool) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := toCBool(param1)

	ret := C.gtk_text_mark_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_mark_get_buffer(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_mark_get_buffer(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_mark_get_deleted(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_mark_get_deleted(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_mark_get_left_gravity(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_mark_get_left_gravity(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_mark_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_mark_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_text_mark_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_mark_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_mark_set_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkTextMark)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_text_mark_set_visible(cValueInstance, cValue0)
}

func Fn_gtk_text_tag_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_text_tag_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_tag_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextTag)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GObject)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkEvent)(unsafe.Pointer(param1))

	cValue2 := (*C.GtkTextIter)(unsafe.Pointer(param2))

	ret := C.gtk_text_tag_event(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_text_tag_get_priority(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextTag)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_tag_get_priority(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_tag_set_priority(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextTag)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_tag_set_priority(cValueInstance, cValue0)
}

func Fn_gtk_text_tag_table_new() unsafe.Pointer {
	ret := C.gtk_text_tag_table_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_tag_table_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextTagTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	ret := C.gtk_text_tag_table_add(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_tag_table_foreach : has callback

func Fn_gtk_text_tag_table_get_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextTagTable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_tag_table_get_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_tag_table_lookup(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkTextTagTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_text_tag_table_lookup(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_tag_table_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextTagTable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextTag)(unsafe.Pointer(param0))

	C.gtk_text_tag_table_remove(cValueInstance, cValue0)
}

func Fn_gtk_text_view_new() unsafe.Pointer {
	ret := C.gtk_text_view_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_new_with_buffer(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkTextBuffer)(unsafe.Pointer(param0))

	ret := C.gtk_text_view_new_with_buffer(cValue0)

	return unsafe.Pointer(ret)
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

func Fn_gtk_text_view_backward_display_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_view_backward_display_line(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_view_backward_display_line_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_view_backward_display_line_start(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_text_view_forward_display_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_view_forward_display_line(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_view_forward_display_line_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_view_forward_display_line_end(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_accepts_tab(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_accepts_tab(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_border_window_size(paramInstance unsafe.Pointer, param0 int) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextWindowType)(param0)

	ret := C.gtk_text_view_get_border_window_size(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_text_view_get_buffer(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_buffer(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_cursor_locations(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	C.gtk_text_view_get_cursor_locations(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_view_get_cursor_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_cursor_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_default_attributes(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_default_attributes(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_editable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_editable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_indent(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_indent(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_get_iter_at_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	ret := C.gtk_text_view_get_iter_at_location(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_iter_at_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 int, param3 int) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	ret := C.gtk_text_view_get_iter_at_position(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_iter_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gtk_text_view_get_iter_location(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_text_view_get_justification(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_justification(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_get_left_margin(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_left_margin(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_text_view_get_overwrite(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_overwrite(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_pixels_above_lines(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_pixels_above_lines(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_get_pixels_below_lines(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_pixels_below_lines(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_get_pixels_inside_wrap(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_pixels_inside_wrap(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_get_right_margin(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_right_margin(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_get_tabs(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_tabs(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_visible_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_text_view_get_visible_rect(cValueInstance, cValue0)
}

func Fn_gtk_text_view_get_window(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkTextWindowType)(param0)

	ret := C.gtk_text_view_get_window(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_window_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkWindow)(unsafe.Pointer(param0))

	ret := C.gtk_text_view_get_window_type(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_text_view_get_wrap_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_wrap_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_im_context_filter_keypress(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	ret := C.gtk_text_view_im_context_filter_keypress(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_view_move_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	C.gtk_text_view_move_child(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_text_view_move_mark_onscreen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextMark)(unsafe.Pointer(param0))

	ret := C.gtk_text_view_move_mark_onscreen(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_text_view_move_visually(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_text_view_move_visually(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_text_view_place_cursor_onscreen(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_place_cursor_onscreen(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_text_view_reset_im_context(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	C.gtk_text_view_reset_im_context(cValueInstance)
}

func Fn_gtk_text_view_scroll_mark_onscreen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextMark)(unsafe.Pointer(param0))

	C.gtk_text_view_scroll_mark_onscreen(cValueInstance, cValue0)
}

func Fn_gtk_text_view_scroll_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 bool, param3 float64, param4 float64) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	cValue1 := (C.gdouble)(param1)

	cValue2 := toCBool(param2)

	cValue3 := (C.gdouble)(param3)

	cValue4 := (C.gdouble)(param4)

	ret := C.gtk_text_view_scroll_to_iter(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
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

func Fn_gtk_text_view_starts_display_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTextIter)(unsafe.Pointer(param0))

	ret := C.gtk_text_view_starts_display_line(cValueInstance, cValue0)

	return toGoBool(ret)
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

// UNSUPPORTED : gtk_theming_engine_get : has varargs

func Fn_gtk_theming_engine_get_background_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_background_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_border(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_border(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_border_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_border_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	ret := C.gtk_theming_engine_get_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_theming_engine_get_font(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	ret := C.gtk_theming_engine_get_font(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_theming_engine_get_junction_sides(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	ret := C.gtk_theming_engine_get_junction_sides(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_theming_engine_get_margin(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_margin(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_padding(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GtkBorder)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_padding(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_theming_engine_get_path(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	ret := C.gtk_theming_engine_get_path(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_theming_engine_get_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_theming_engine_get_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_theming_engine_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	ret := C.gtk_theming_engine_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_theming_engine_get_state(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	ret := C.gtk_theming_engine_get_state(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_theming_engine_get_style : has varargs

func Fn_gtk_theming_engine_get_style_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.gtk_theming_engine_get_style_property(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_theming_engine_get_style_valist : has va_list

// UNSUPPORTED : gtk_theming_engine_get_valist : has va_list

func Fn_gtk_theming_engine_has_class(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_theming_engine_has_class(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_theming_engine_has_region(paramInstance unsafe.Pointer, param0 string, param1 *int) bool {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GtkRegionFlags)(unsafe.Pointer(param1))

	ret := C.gtk_theming_engine_has_region(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_theming_engine_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	ret := C.gtk_theming_engine_lookup_color(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_theming_engine_state_is_running(paramInstance unsafe.Pointer, param0 int, param1 *float64) bool {
	cValueInstance := (*C.GtkThemingEngine)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateType)(param0)

	cValue1 := (*C.gdouble)(unsafe.Pointer(param1))

	ret := C.gtk_theming_engine_state_is_running(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_theming_engine_load(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_theming_engine_load(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_theming_engine_register_property : has callback

func Fn_gtk_toggle_action_new(param0 string, param1 string, param2 string, param3 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gchar)(C.CString(param3))
	defer C.free(unsafe.Pointer(cValue3))

	ret := C.gtk_toggle_action_new(cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_action_get_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toggle_action_get_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_toggle_action_get_draw_as_radio(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToggleAction)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toggle_action_get_draw_as_radio(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_toggle_button_new() unsafe.Pointer {
	ret := C.gtk_toggle_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_button_new_with_label(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_toggle_button_new_with_label(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_button_new_with_mnemonic(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_toggle_button_new_with_mnemonic(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_button_get_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toggle_button_get_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_toggle_button_get_inconsistent(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toggle_button_get_inconsistent(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_toggle_button_get_mode(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToggleButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toggle_button_get_mode(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_toggle_tool_button_new() unsafe.Pointer {
	ret := C.gtk_toggle_tool_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_tool_button_new_from_stock(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_toggle_tool_button_new_from_stock(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_tool_button_get_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToggleToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toggle_tool_button_get_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_toggle_tool_button_set_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToggleToolButton)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_toggle_tool_button_set_active(cValueInstance, cValue0)
}

func Fn_gtk_tool_button_new(param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_tool_button_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_button_new_from_stock(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_tool_button_new_from_stock(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_button_get_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tool_button_get_icon_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_icon_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_button_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tool_button_get_label_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_label_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_button_get_stock_id(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_stock_id(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tool_button_get_use_underline(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolButton)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_button_get_use_underline(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_tool_item_new() unsafe.Pointer {
	ret := C.gtk_tool_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_get_ellipsize_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_ellipsize_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_expand(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_expand(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_homogeneous(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_homogeneous(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_icon_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_icon_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_is_important(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_is_important(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_proxy_menu_item(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_tool_item_get_proxy_menu_item(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_get_relief_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_relief_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_text_alignment(paramInstance unsafe.Pointer) float32 {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_text_alignment(cValueInstance)

	return (float32)(ret)
}

func Fn_gtk_tool_item_get_text_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_text_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_text_size_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_text_size_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_get_toolbar_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_toolbar_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_get_use_drag_window(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_use_drag_window(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_visible_horizontal(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_visible_horizontal(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_get_visible_vertical(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_get_visible_vertical(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_rebuild_menu(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	C.gtk_tool_item_rebuild_menu(cValueInstance)
}

func Fn_gtk_tool_item_retrieve_proxy_menu_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItem)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_retrieve_proxy_menu_item(cValueInstance)

	return unsafe.Pointer(ret)
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

func Fn_gtk_tool_item_group_new(param0 string) unsafe.Pointer {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_tool_item_group_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_group_get_collapsed(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_collapsed(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tool_item_group_get_drop_item(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_tool_item_group_get_drop_item(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_group_get_ellipsize(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_ellipsize(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_group_get_header_relief(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_header_relief(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_item_group_get_item_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	ret := C.gtk_tool_item_group_get_item_position(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_tool_item_group_get_label(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_label(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tool_item_group_get_label_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_label_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_group_get_n_items(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_item_group_get_n_items(cValueInstance)

	return (uint)(ret)
}

func Fn_gtk_tool_item_group_get_nth_item(paramInstance unsafe.Pointer, param0 uint) unsafe.Pointer {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gtk_tool_item_group_get_nth_item(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_item_group_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_tool_item_group_insert(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_item_group_set_collapsed(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_tool_item_group_set_collapsed(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_group_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.PangoEllipsizeMode)(param0)

	C.gtk_tool_item_group_set_ellipsize(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_group_set_header_relief(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkReliefStyle)(param0)

	C.gtk_tool_item_group_set_header_relief(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_group_set_item_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_tool_item_group_set_item_position(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_item_group_set_label(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_tool_item_group_set_label(cValueInstance, cValue0)
}

func Fn_gtk_tool_item_group_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkToolItemGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_tool_item_group_set_label_widget(cValueInstance, cValue0)
}

func Fn_gtk_tool_palette_new() unsafe.Pointer {
	ret := C.gtk_tool_palette_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_add_drag_dest(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.GtkDestDefaults)(param1)

	cValue2 := (C.GtkToolPaletteDragTargets)(param2)

	cValue3 := (C.GdkDragAction)(param3)

	C.gtk_tool_palette_add_drag_dest(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_tool_palette_get_drag_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkSelectionData)(unsafe.Pointer(param0))

	ret := C.gtk_tool_palette_get_drag_item(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_drop_group(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_tool_palette_get_drop_group(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_drop_item(paramInstance unsafe.Pointer, param0 int, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_tool_palette_get_drop_item(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_exclusive(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	ret := C.gtk_tool_palette_get_exclusive(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tool_palette_get_expand(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	ret := C.gtk_tool_palette_get_expand(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tool_palette_get_group_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	ret := C.gtk_tool_palette_get_group_position(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_tool_palette_get_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_palette_get_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_icon_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_palette_get_icon_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_palette_get_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_palette_get_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_palette_get_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_palette_get_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_set_drag_source(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkToolPaletteDragTargets)(param0)

	C.gtk_tool_palette_set_drag_source(cValueInstance, cValue0)
}

func Fn_gtk_tool_palette_set_exclusive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_tool_palette_set_exclusive(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_palette_set_expand(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_tool_palette_set_expand(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_palette_set_group_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItemGroup)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_tool_palette_set_group_position(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tool_palette_set_icon_size(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkIconSize)(param0)

	C.gtk_tool_palette_set_icon_size(cValueInstance, cValue0)
}

func Fn_gtk_tool_palette_set_style(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkToolbarStyle)(param0)

	C.gtk_tool_palette_set_style(cValueInstance, cValue0)
}

func Fn_gtk_tool_palette_unset_icon_size(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	C.gtk_tool_palette_unset_icon_size(cValueInstance)
}

func Fn_gtk_tool_palette_unset_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolPalette)(unsafe.Pointer(paramInstance))

	C.gtk_tool_palette_unset_style(cValueInstance)
}

func Fn_gtk_tool_palette_get_drag_target_group() unsafe.Pointer {
	ret := C.gtk_tool_palette_get_drag_target_group()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_palette_get_drag_target_item() unsafe.Pointer {
	ret := C.gtk_tool_palette_get_drag_target_item()

	return unsafe.Pointer(ret)
}

func Fn_gtk_toolbar_new() unsafe.Pointer {
	ret := C.gtk_toolbar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_toolbar_get_drop_index(paramInstance unsafe.Pointer, param0 int, param1 int) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_toolbar_get_drop_index(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_icon_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toolbar_get_icon_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_item_index(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkToolItem)(unsafe.Pointer(param0))

	ret := C.gtk_toolbar_get_item_index(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_n_items(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toolbar_get_n_items(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_nth_item(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_toolbar_get_nth_item(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toolbar_get_relief_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toolbar_get_relief_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_show_arrow(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toolbar_get_show_arrow(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_toolbar_get_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolbar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toolbar_get_style(cValueInstance)

	return (int)(ret)
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

func Fn_gtk_tooltip_set_icon_from_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkTooltip)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GIcon)(unsafe.Pointer(param0))

	cValue1 := (C.GtkIconSize)(param1)

	C.gtk_tooltip_set_icon_from_gicon(cValueInstance, cValue0, cValue1)
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

func Fn_gtk_toplevel_accessible_get_children(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToplevelAccessible)(unsafe.Pointer(paramInstance))

	ret := C.gtk_toplevel_accessible_get_children(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_filter_clear_cache(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_filter_clear_cache(cValueInstance)
}

func Fn_gtk_tree_model_filter_convert_child_iter_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	ret := C.gtk_tree_model_filter_convert_child_iter_to_iter(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_filter_convert_child_path_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_filter_convert_child_path_to_path(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_filter_convert_iter_to_child_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_filter_convert_iter_to_child_iter(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_model_filter_convert_path_to_child_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_filter_convert_path_to_child_path(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_filter_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_model_filter_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_filter_refilter(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_filter_refilter(cValueInstance)
}

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : has callback

func Fn_gtk_tree_model_filter_set_visible_column(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeModelFilter)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_tree_model_filter_set_visible_column(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : has callback

func Fn_gtk_tree_model_sort_clear_cache(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_sort_clear_cache(cValueInstance)
}

func Fn_gtk_tree_model_sort_convert_child_iter_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	ret := C.gtk_tree_model_sort_convert_child_iter_to_iter(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_sort_convert_child_path_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_sort_convert_child_path_to_path(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_sort_convert_iter_to_child_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_sort_convert_iter_to_child_iter(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_model_sort_convert_path_to_child_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_sort_convert_path_to_child_path(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_sort_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_model_sort_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_sort_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_sort_iter_is_valid(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_sort_reset_default_sort_func(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModelSort)(unsafe.Pointer(paramInstance))

	C.gtk_tree_model_sort_reset_default_sort_func(cValueInstance)
}

func Fn_gtk_tree_selection_count_selected_rows(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_selection_count_selected_rows(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_selection_get_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_selection_get_mode(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

func Fn_gtk_tree_selection_get_selected(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreeModel)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	ret := C.gtk_tree_selection_get_selected(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_selection_get_selected_rows(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreeModel)(unsafe.Pointer(param0))

	ret := C.gtk_tree_selection_get_selected_rows(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_selection_get_tree_view(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_selection_get_tree_view(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_selection_get_user_data(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_selection_get_user_data(cValueInstance)

	return (unsafe.Pointer)(ret)
}

func Fn_gtk_tree_selection_iter_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_selection_iter_is_selected(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_selection_path_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_selection_path_is_selected(cValueInstance, cValue0)

	return toGoBool(ret)
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

// UNSUPPORTED : gtk_tree_selection_selected_foreach : has callback

func Fn_gtk_tree_selection_set_mode(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTreeSelection)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkSelectionMode)(param0)

	C.gtk_tree_selection_set_mode(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_tree_selection_set_select_function : has callback

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

// UNSUPPORTED : gtk_tree_store_new : has varargs

// UNSUPPORTED : gtk_tree_store_newv : has non-string array param types

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

// UNSUPPORTED : gtk_tree_store_insert_with_values : has varargs

// UNSUPPORTED : gtk_tree_store_insert_with_valuesv : has non-string array param columns

func Fn_gtk_tree_store_is_ancestor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	ret := C.gtk_tree_store_is_ancestor(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_store_iter_depth(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_store_iter_depth(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_tree_store_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_store_iter_is_valid(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_tree_store_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_store_remove(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_store_reorder : has non-string array param new_order

// UNSUPPORTED : gtk_tree_store_set : has varargs

// UNSUPPORTED : gtk_tree_store_set_column_types : has non-string array param types

// UNSUPPORTED : gtk_tree_store_set_valist : has va_list

func Fn_gtk_tree_store_set_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_tree_store_set_value(cValueInstance, cValue0, cValue1, cValue2)
}

// UNSUPPORTED : gtk_tree_store_set_valuesv : has non-string array param columns

func Fn_gtk_tree_store_swap(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeStore)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_store_swap(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_new() unsafe.Pointer {
	ret := C.gtk_tree_view_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_new_with_model(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkTreeModel)(unsafe.Pointer(param0))

	ret := C.gtk_tree_view_new_with_model(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_append_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param0))

	ret := C.gtk_tree_view_append_column(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_tree_view_collapse_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_collapse_all(cValueInstance)
}

func Fn_gtk_tree_view_collapse_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_view_collapse_row(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_tree_view_create_row_drag_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_view_create_row_drag_icon(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_enable_model_drag_dest : has non-string array param targets

// UNSUPPORTED : gtk_tree_view_enable_model_drag_source : has non-string array param targets

func Fn_gtk_tree_view_expand_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	C.gtk_tree_view_expand_all(cValueInstance)
}

func Fn_gtk_tree_view_expand_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	ret := C.gtk_tree_view_expand_row(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
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

func Fn_gtk_tree_view_get_bin_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_bin_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_cell_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	cValue2 := (*C.GdkRectangle)(unsafe.Pointer(param2))

	C.gtk_tree_view_get_cell_area(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_view_get_column(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_tree_view_get_column(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_columns(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_columns(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_cursor(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	C.gtk_tree_view_get_cursor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_get_dest_row_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *int) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (*C.GtkTreeViewDropPosition)(unsafe.Pointer(param3))

	ret := C.gtk_tree_view_get_dest_row_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_drag_dest_row(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *int) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewDropPosition)(unsafe.Pointer(param1))

	C.gtk_tree_view_get_drag_dest_row(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_get_enable_search(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_enable_search(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_enable_tree_lines(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_enable_tree_lines(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_expander_column(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_expander_column(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_fixed_height_mode(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_fixed_height_mode(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_grid_lines(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_grid_lines(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_get_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_headers_clickable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_headers_clickable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_headers_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_headers_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_hover_expand(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_hover_expand(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_hover_selection(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_hover_selection(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_level_indentation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_level_indentation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_get_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_path_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *int, param5 *int) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (**C.GtkTreeViewColumn)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (*C.gint)(unsafe.Pointer(param5))

	ret := C.gtk_tree_view_get_path_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_reorderable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_reorderable(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

func Fn_gtk_tree_view_get_rubber_banding(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_rubber_banding(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_rules_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_rules_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_search_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_search_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_get_search_entry(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_search_entry(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

func Fn_gtk_tree_view_get_selection(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_selection(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_show_expanders(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_show_expanders(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_tooltip_column(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_tooltip_column(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_get_tooltip_context(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := toCBool(param2)

	cValue3 := (**C.GtkTreeModel)(unsafe.Pointer(param3))

	cValue4 := (**C.GtkTreePath)(unsafe.Pointer(param4))

	cValue5 := (*C.GtkTreeIter)(unsafe.Pointer(param5))

	ret := C.gtk_tree_view_get_tooltip_context(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_get_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_visible_range(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (**C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (**C.GtkTreePath)(unsafe.Pointer(param1))

	ret := C.gtk_tree_view_get_visible_range(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_visible_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	C.gtk_tree_view_get_visible_rect(cValueInstance, cValue0)
}

func Fn_gtk_tree_view_insert_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	ret := C.gtk_tree_view_insert_column(cValueInstance, cValue0, cValue1)

	return (int)(ret)
}

// UNSUPPORTED : gtk_tree_view_insert_column_with_attributes : has varargs

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : has callback

func Fn_gtk_tree_view_is_blank_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *int, param5 *int) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (**C.GtkTreePath)(unsafe.Pointer(param2))

	cValue3 := (**C.GtkTreeViewColumn)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (*C.gint)(unsafe.Pointer(param5))

	ret := C.gtk_tree_view_is_blank_at_pos(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_is_rubber_banding_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_is_rubber_banding_active(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : has callback

func Fn_gtk_tree_view_move_column_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	C.gtk_tree_view_move_column_after(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_remove_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param0))

	ret := C.gtk_tree_view_remove_column(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_tree_view_row_activated(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeViewColumn)(unsafe.Pointer(param1))

	C.gtk_tree_view_row_activated(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_view_row_expanded(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeView)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_view_row_expanded(cValueInstance, cValue0)

	return toGoBool(ret)
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

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : has callback

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

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : has callback

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

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : has callback

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

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : has callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : has callback

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

func Fn_gtk_tree_view_column_new() unsafe.Pointer {
	ret := C.gtk_tree_view_column_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_column_new_with_area(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkCellArea)(unsafe.Pointer(param0))

	ret := C.gtk_tree_view_column_new_with_area(cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_column_new_with_attributes : has varargs

func Fn_gtk_tree_view_column_add_attribute(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	C.gtk_tree_view_column_add_attribute(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_view_column_cell_get_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) bool {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	ret := C.gtk_tree_view_column_cell_get_position(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
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

func Fn_gtk_tree_view_column_cell_is_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_cell_is_visible(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_tree_view_column_get_alignment(paramInstance unsafe.Pointer) float32 {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_alignment(cValueInstance)

	return (float32)(ret)
}

func Fn_gtk_tree_view_column_get_button(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_button(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_column_get_clickable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_clickable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_expand(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_expand(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_fixed_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_fixed_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_max_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_max_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_min_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_min_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_reorderable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_reorderable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_resizable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_resizable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_sizing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_sizing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_sort_column_id(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_sort_column_id(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_sort_indicator(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_sort_indicator(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_sort_order(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_sort_order(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_spacing(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_spacing(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_tree_view_column_get_tree_view(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_tree_view(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_column_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_column_get_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeViewColumn)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_view_column_get_width(cValueInstance)

	return (int)(ret)
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

// UNSUPPORTED : gtk_tree_view_column_set_attributes : has varargs

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : has callback

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

func Fn_gtk_ui_manager_new() unsafe.Pointer {
	ret := C.gtk_ui_manager_new()

	return unsafe.Pointer(ret)
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

func Fn_gtk_ui_manager_add_ui_from_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_ui_manager_add_ui_from_file(cValueInstance, cValue0, cError)

	return (uint)(ret)
}

func Fn_gtk_ui_manager_add_ui_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64, error unsafe.Pointer) uint {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gssize)(param1)

	cError := (**C.GError)(error)

	ret := C.gtk_ui_manager_add_ui_from_string(cValueInstance, cValue0, cValue1, cError)

	return (uint)(ret)
}

func Fn_gtk_ui_manager_ensure_update(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	C.gtk_ui_manager_ensure_update(cValueInstance)
}

func Fn_gtk_ui_manager_get_accel_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_get_accel_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_get_action(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_ui_manager_get_action(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_get_action_groups(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_get_action_groups(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_get_add_tearoffs(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_get_add_tearoffs(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_ui_manager_get_toplevels(paramInstance unsafe.Pointer, param0 int) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkUIManagerItemType)(param0)

	ret := C.gtk_ui_manager_get_toplevels(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_get_ui(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_get_ui(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_ui_manager_get_widget(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_ui_manager_get_widget(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_ui_manager_insert_action_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkActionGroup)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_ui_manager_insert_action_group(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_ui_manager_new_merge_id(paramInstance unsafe.Pointer) uint {
	cValueInstance := (*C.GtkUIManager)(unsafe.Pointer(paramInstance))

	ret := C.gtk_ui_manager_new_merge_id(cValueInstance)

	return (uint)(ret)
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

func Fn_gtk_vbox_new(param0 bool, param1 int) unsafe.Pointer {
	cValue0 := toCBool(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_vbox_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_vbutton_box_new() unsafe.Pointer {
	ret := C.gtk_vbutton_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_vpaned_new() unsafe.Pointer {
	ret := C.gtk_vpaned_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_vscale_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	ret := C.gtk_vscale_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_vscale_new_with_range(param0 float64, param1 float64, param2 float64) unsafe.Pointer {
	cValue0 := (C.gdouble)(param0)

	cValue1 := (C.gdouble)(param1)

	cValue2 := (C.gdouble)(param2)

	ret := C.gtk_vscale_new_with_range(cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_vscrollbar_new(param0 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	ret := C.gtk_vscrollbar_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_vseparator_new() unsafe.Pointer {
	ret := C.gtk_vseparator_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_viewport_new(param0 unsafe.Pointer, param1 unsafe.Pointer) unsafe.Pointer {
	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkAdjustment)(unsafe.Pointer(param1))

	ret := C.gtk_viewport_new(cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_viewport_get_bin_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	ret := C.gtk_viewport_get_bin_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_viewport_get_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	ret := C.gtk_viewport_get_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_viewport_get_shadow_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	ret := C.gtk_viewport_get_shadow_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_viewport_get_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	ret := C.gtk_viewport_get_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_viewport_get_view_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkViewport)(unsafe.Pointer(paramInstance))

	ret := C.gtk_viewport_get_view_window(cValueInstance)

	return unsafe.Pointer(ret)
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

func Fn_gtk_volume_button_new() unsafe.Pointer {
	ret := C.gtk_volume_button_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_widget_new : has varargs

func Fn_gtk_widget_activate(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_activate(cValueInstance)

	return toGoBool(ret)
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

func Fn_gtk_widget_add_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := (C.GdkEventMask)(param1)

	C.gtk_widget_add_device_events(cValueInstance, cValue0, cValue1)
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

// UNSUPPORTED : gtk_widget_add_tick_callback : has callback

func Fn_gtk_widget_can_activate_accel(paramInstance unsafe.Pointer, param0 uint) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	ret := C.gtk_widget_can_activate_accel(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_child_focus(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkDirectionType)(param0)

	ret := C.gtk_widget_child_focus(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_widget_compute_expand(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkOrientation)(param0)

	ret := C.gtk_widget_compute_expand(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_create_pango_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_create_pango_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_create_pango_layout(paramInstance unsafe.Pointer, param0 string) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_widget_create_pango_layout(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
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

func Fn_gtk_widget_device_is_shadowed(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gtk_widget_device_is_shadowed(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_drag_begin(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTargetList)(unsafe.Pointer(param0))

	cValue1 := (C.GdkDragAction)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.GdkEvent)(unsafe.Pointer(param3))

	ret := C.gtk_drag_begin(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return unsafe.Pointer(ret)
}

func Fn_gtk_drag_check_threshold(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (C.gint)(param3)

	ret := C.gtk_drag_check_threshold(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
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

func Fn_gtk_drag_dest_find_target(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) gdk.Atom {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDragContext)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTargetList)(unsafe.Pointer(param1))

	ret := C.gtk_drag_dest_find_target(cValueInstance, cValue0, cValue1)

	return gdk.Atom(unsafe.Pointer(ret))
}

func Fn_gtk_drag_dest_get_target_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_drag_dest_get_target_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_drag_dest_get_track_motion(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_drag_dest_get_track_motion(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_drag_dest_set : has non-string array param targets

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

func Fn_gtk_drag_source_get_target_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_drag_source_get_target_list(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_drag_source_set : has non-string array param targets

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

func Fn_gtk_widget_draw(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_t)(unsafe.Pointer(param0))

	C.gtk_widget_draw(cValueInstance, cValue0)
}

func Fn_gtk_widget_ensure_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_ensure_style(cValueInstance)
}

func Fn_gtk_widget_error_bell(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_error_bell(cValueInstance)
}

func Fn_gtk_widget_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	ret := C.gtk_widget_event(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_freeze_child_notify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_freeze_child_notify(cValueInstance)
}

func Fn_gtk_widget_get_accessible(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_accessible(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_allocated_height(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_allocated_height(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_allocated_width(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_allocated_width(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_allocation(paramInstance unsafe.Pointer, param0 *gdk.Rectangle) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAllocation)(unsafe.Pointer(param0))

	C.gtk_widget_get_allocation(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_ancestor(paramInstance unsafe.Pointer, param0 uint64) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	ret := C.gtk_widget_get_ancestor(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_app_paintable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_app_paintable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_can_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_can_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_can_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_can_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_child_requisition(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRequisition)(unsafe.Pointer(param0))

	C.gtk_widget_get_child_requisition(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_child_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_child_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_clipboard(paramInstance unsafe.Pointer, param0 gdk.Atom) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GdkAtom)(unsafe.Pointer(param0))

	ret := C.gtk_widget_get_clipboard(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_composite_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_composite_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_widget_get_device_enabled(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gtk_widget_get_device_enabled(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gtk_widget_get_device_events(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_widget_get_direction(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_direction(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_display(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_display(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_double_buffered(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_double_buffered(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_events(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_events(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_halign(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_halign(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_has_tooltip(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_has_tooltip(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_has_window(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_has_window(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_hexpand(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_hexpand(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_hexpand_set(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_hexpand_set(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_mapped(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_mapped(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_margin_bottom(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_bottom(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_margin_left(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_left(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_margin_right(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_right(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_margin_top(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_margin_top(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_modifier_style(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_modifier_style(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_widget_get_no_show_all(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_no_show_all(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_pango_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_pango_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_parent(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_parent(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_parent_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_parent_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_path(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_path(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_pointer(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_widget_get_pointer(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_preferred_height(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_widget_get_preferred_height(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_widget_get_preferred_height_for_width(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_widget_get_preferred_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRequisition)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkRequisition)(unsafe.Pointer(param1))

	C.gtk_widget_get_preferred_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_preferred_width(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_widget_get_preferred_width(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_widget_get_preferred_width_for_height(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_widget_get_realized(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_realized(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_receives_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_receives_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_request_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_request_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_requisition(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRequisition)(unsafe.Pointer(param0))

	C.gtk_widget_get_requisition(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_root_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_root_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_settings(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_settings(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_size_request(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_widget_get_size_request(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_get_state(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_state(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_state_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_state_flags(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_style(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_style(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_style_context(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_style_context(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_support_multidevice(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_support_multidevice(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_template_child(paramInstance unsafe.Pointer, param0 uint64, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GType)(param0)

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_widget_get_template_child(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_tooltip_markup(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_tooltip_markup(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_widget_get_tooltip_text(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_tooltip_text(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_widget_get_tooltip_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_tooltip_window(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_toplevel(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_toplevel(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_valign(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_valign(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_widget_get_vexpand(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_vexpand(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_vexpand_set(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_vexpand_set(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_visual(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_visual(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_window(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_window(cValueInstance)

	return unsafe.Pointer(ret)
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

func Fn_gtk_widget_has_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_has_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_has_grab(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_grab(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_has_rc_style(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_rc_style(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_has_screen(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_has_screen(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_hide(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_hide(cValueInstance)
}

func Fn_gtk_widget_hide_on_delete(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_hide_on_delete(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_in_destruction(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_in_destruction(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_input_shape_combine_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	C.gtk_widget_input_shape_combine_region(cValueInstance, cValue0)
}

func Fn_gtk_widget_intersect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	ret := C.gtk_widget_intersect(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_ancestor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_widget_is_ancestor(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_composited(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_is_composited(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_drawable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_is_drawable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_is_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_sensitive(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_is_sensitive(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_toplevel(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_is_toplevel(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_widget_keynav_failed(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkDirectionType)(param0)

	ret := C.gtk_widget_keynav_failed(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_widget_list_accel_closures(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_list_accel_closures(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : has array return

func Fn_gtk_widget_list_mnemonic_labels(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_list_mnemonic_labels(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_map(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_map(cValueInstance)
}

func Fn_gtk_widget_mnemonic_activate(paramInstance unsafe.Pointer, param0 bool) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	ret := C.gtk_widget_mnemonic_activate(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_widget_override_background_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_widget_override_background_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_override_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_widget_override_color(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_override_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRGBA)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_widget_override_cursor(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_override_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontDescription)(unsafe.Pointer(param0))

	C.gtk_widget_override_font(cValueInstance, cValue0)
}

func Fn_gtk_widget_override_symbolic_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GdkRGBA)(unsafe.Pointer(param1))

	C.gtk_widget_override_symbolic_color(cValueInstance, cValue0, cValue1)
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

func Fn_gtk_widget_queue_draw_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	C.gtk_widget_queue_draw_region(cValueInstance, cValue0)
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

func Fn_gtk_widget_region_intersect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	ret := C.gtk_widget_region_intersect(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_remove_accelerator(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 int) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAccelGroup)(unsafe.Pointer(param0))

	cValue1 := (C.guint)(param1)

	cValue2 := (C.GdkModifierType)(param2)

	ret := C.gtk_widget_remove_accelerator(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_widget_remove_mnemonic_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_widget_remove_mnemonic_label(cValueInstance, cValue0)
}

func Fn_gtk_widget_render_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 string) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	ret := C.gtk_widget_render_icon(cValueInstance, cValue0, cValue1, cValue2)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_render_icon_pixbuf(paramInstance unsafe.Pointer, param0 string, param1 int) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.GtkIconSize)(param1)

	ret := C.gtk_widget_render_icon_pixbuf(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
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

func Fn_gtk_widget_reset_style(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_reset_style(cValueInstance)
}

func Fn_gtk_widget_send_expose(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	ret := C.gtk_widget_send_expose(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_widget_send_focus_change(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	ret := C.gtk_widget_send_focus_change(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_widget_set_device_enabled(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_widget_set_device_enabled(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_set_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	cValue1 := (C.GdkEventMask)(param1)

	C.gtk_widget_set_device_events(cValueInstance, cValue0, cValue1)
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

func Fn_gtk_widget_set_mapped(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_mapped(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_margin_bottom(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_bottom(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_margin_left(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_left(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_margin_right(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_right(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_margin_top(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_widget_set_margin_top(cValueInstance, cValue0)
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

func Fn_gtk_widget_set_realized(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_realized(cValueInstance, cValue0)
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

func Fn_gtk_widget_set_state_flags(paramInstance unsafe.Pointer, param0 int, param1 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	cValue1 := toCBool(param1)

	C.gtk_widget_set_state_flags(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_widget_set_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkStyle)(unsafe.Pointer(param0))

	C.gtk_widget_set_style(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_support_multidevice(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_widget_set_support_multidevice(cValueInstance, cValue0)
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

func Fn_gtk_widget_shape_combine_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_region_t)(unsafe.Pointer(param0))

	C.gtk_widget_shape_combine_region(cValueInstance, cValue0)
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

func Fn_gtk_widget_style_attach(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_style_attach(cValueInstance)
}

// UNSUPPORTED : gtk_widget_style_get : has varargs

func Fn_gtk_widget_style_get_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (*C.GValue)(unsafe.Pointer(param1))

	C.gtk_widget_style_get_property(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_widget_style_get_valist : has va_list

func Fn_gtk_widget_thaw_child_notify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	C.gtk_widget_thaw_child_notify(cValueInstance)
}

func Fn_gtk_widget_translate_coordinates(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 *int, param4 *int) bool {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (C.gint)(param2)

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	ret := C.gtk_widget_translate_coordinates(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
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

func Fn_gtk_widget_unset_state_flags(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkStateFlags)(param0)

	C.gtk_widget_unset_state_flags(cValueInstance, cValue0)
}

func Fn_gtk_widget_get_default_direction() int {
	ret := C.gtk_widget_get_default_direction()

	return (int)(ret)
}

func Fn_gtk_widget_get_default_style() unsafe.Pointer {
	ret := C.gtk_widget_get_default_style()

	return unsafe.Pointer(ret)
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

func Fn_gtk_window_new(param0 int) unsafe.Pointer {
	cValue0 := (C.GtkWindowType)(param0)

	ret := C.gtk_window_new(cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_activate_default(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_activate_default(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_activate_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_activate_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_activate_key(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	ret := C.gtk_window_activate_key(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_window_get_accept_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_accept_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_application(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_application(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_decorated(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_decorated(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_default_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_window_get_default_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_get_default_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_default_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_deletable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_deletable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_destroy_with_parent(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_destroy_with_parent(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_focus(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_focus(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_focus_on_map(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_focus_on_map(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_gravity(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_gravity(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_window_get_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_has_resize_grip(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_has_resize_grip(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_icon(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_icon(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_icon_list(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_icon_list(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_icon_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_icon_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_window_get_mnemonic_modifier(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_mnemonic_modifier(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_window_get_mnemonics_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_mnemonics_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_modal(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_modal(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_opacity(paramInstance unsafe.Pointer) float64 {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_opacity(cValueInstance)

	return (float64)(ret)
}

func Fn_gtk_window_get_position(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_window_get_position(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_get_resizable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_resizable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_resize_grip_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkRectangle)(unsafe.Pointer(param0))

	ret := C.gtk_window_get_resize_grip_area(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_window_get_role(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_role(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_window_get_screen(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_screen(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	C.gtk_window_get_size(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_get_skip_pager_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_skip_pager_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_skip_taskbar_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_skip_taskbar_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_title(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_title(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_window_get_transient_for(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_transient_for(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_type_hint(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_type_hint(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_window_get_urgency_hint(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_urgency_hint(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_get_window_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_get_window_type(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_window_has_group(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_has_group(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_has_toplevel_focus(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_has_toplevel_focus(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_iconify(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_iconify(cValueInstance)
}

func Fn_gtk_window_is_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_is_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_maximize(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	C.gtk_window_maximize(cValueInstance)
}

func Fn_gtk_window_mnemonic_activate(paramInstance unsafe.Pointer, param0 uint, param1 int) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (C.GdkModifierType)(param1)

	ret := C.gtk_window_mnemonic_activate(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_window_move(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_window_move(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_parse_geometry(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_window_parse_geometry(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_window_propagate_key_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEventKey)(unsafe.Pointer(param0))

	ret := C.gtk_window_propagate_key_event(cValueInstance, cValue0)

	return toGoBool(ret)
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

func Fn_gtk_window_resize_grip_is_visible(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_resize_grip_is_visible(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_window_resize_to_geometry(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_window_resize_to_geometry(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_window_set_accept_focus(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_accept_focus(cValueInstance, cValue0)
}

func Fn_gtk_window_set_application(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkApplication)(unsafe.Pointer(param0))

	C.gtk_window_set_application(cValueInstance, cValue0)
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

func Fn_gtk_window_set_default_geometry(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_window_set_default_geometry(cValueInstance, cValue0, cValue1)
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

func Fn_gtk_window_set_has_resize_grip(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_has_resize_grip(cValueInstance, cValue0)
}

func Fn_gtk_window_set_has_user_ref_count(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_has_user_ref_count(cValueInstance, cValue0)
}

func Fn_gtk_window_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_window_set_icon(cValueInstance, cValue0)
}

func Fn_gtk_window_set_icon_from_file(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_window_set_icon_from_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
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

func Fn_gtk_window_set_mnemonics_visible(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_window_set_mnemonics_visible(cValueInstance, cValue0)
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

func Fn_gtk_window_get_default_icon_list() unsafe.Pointer {
	ret := C.gtk_window_get_default_icon_list()

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_default_icon_name() string {
	ret := C.gtk_window_get_default_icon_name()

	return C.GoString(ret)
}

func Fn_gtk_window_list_toplevels() unsafe.Pointer {
	ret := C.gtk_window_list_toplevels()

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_set_auto_startup_notification(param0 bool) {
	cValue0 := toCBool(param0)

	C.gtk_window_set_auto_startup_notification(cValue0)
}

func Fn_gtk_window_set_default_icon(param0 unsafe.Pointer) {
	cValue0 := (*C.GdkPixbuf)(unsafe.Pointer(param0))

	C.gtk_window_set_default_icon(cValue0)
}

func Fn_gtk_window_set_default_icon_from_file(param0 string, error unsafe.Pointer) bool {
	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_window_set_default_icon_from_file(cValue0, cError)

	return toGoBool(ret)
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

func Fn_gtk_window_group_new() unsafe.Pointer {
	ret := C.gtk_window_group_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_group_add_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_window_group_add_window(cValueInstance, cValue0)
}

func Fn_gtk_window_group_get_current_device_grab(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkDevice)(unsafe.Pointer(param0))

	ret := C.gtk_window_group_get_current_device_grab(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_group_get_current_grab(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_group_get_current_grab(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_group_list_windows(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	ret := C.gtk_window_group_list_windows(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_group_remove_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWindowGroup)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWindow)(unsafe.Pointer(param0))

	C.gtk_window_group_remove_window(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_actionable_set_action_target : has varargs

func Fn_gtk_activatable_do_set_related_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_activatable_do_set_related_action(cValueInstance, cValue0)
}

func Fn_gtk_activatable_get_related_action(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_activatable_get_related_action(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_activatable_get_use_action_appearance(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_activatable_get_use_action_appearance(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_activatable_set_related_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_activatable_set_related_action(cValueInstance, cValue0)
}

func Fn_gtk_activatable_set_use_action_appearance(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_activatable_set_use_action_appearance(cValueInstance, cValue0)
}

func Fn_gtk_activatable_sync_action_properties(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkActivatable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAction)(unsafe.Pointer(param0))

	C.gtk_activatable_sync_action_properties(cValueInstance, cValue0)
}

func Fn_gtk_app_chooser_get_app_info(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkAppChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_get_app_info(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_app_chooser_get_content_type(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkAppChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_app_chooser_get_content_type(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_app_chooser_refresh(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkAppChooser)(unsafe.Pointer(paramInstance))

	C.gtk_app_chooser_refresh(cValueInstance)
}

func Fn_gtk_buildable_add_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	C.gtk_buildable_add_child(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_buildable_construct_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_buildable_construct_child(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_buildable_custom_finished(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (C.gpointer)(param3)

	C.gtk_buildable_custom_finished(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_buildable_custom_tag_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 *unsafe.Pointer) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.gpointer)(unsafe.Pointer(param3))

	C.gtk_buildable_custom_tag_end(cValueInstance, cValue0, cValue1, cValue2, cValue3)
}

func Fn_gtk_buildable_custom_tag_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 *unsafe.Pointer) bool {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.GObject)(unsafe.Pointer(param1))

	cValue2 := (*C.gchar)(C.CString(param2))
	defer C.free(unsafe.Pointer(cValue2))

	cValue3 := (*C.GMarkupParser)(unsafe.Pointer(param3))

	cValue4 := (*C.gpointer)(unsafe.Pointer(param4))

	ret := C.gtk_buildable_custom_tag_start(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4)

	return toGoBool(ret)
}

func Fn_gtk_buildable_get_internal_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) unsafe.Pointer {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_buildable_get_internal_child(cValueInstance, cValue0, cValue1)

	return unsafe.Pointer(ret)
}

func Fn_gtk_buildable_get_name(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_buildable_get_name(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_buildable_parser_finished(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	C.gtk_buildable_parser_finished(cValueInstance, cValue0)
}

func Fn_gtk_buildable_set_buildable_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkBuilder)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_buildable_set_buildable_property(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_buildable_set_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkBuildable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_buildable_set_name(cValueInstance, cValue0)
}

func Fn_gtk_cell_accessible_parent_activate(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellAccessibleParent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	C.gtk_cell_accessible_parent_activate(cValueInstance, cValue0)
}

func Fn_gtk_cell_accessible_parent_edit(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellAccessibleParent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	C.gtk_cell_accessible_parent_edit(cValueInstance, cValue0)
}

func Fn_gtk_cell_accessible_parent_expand_collapse(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellAccessibleParent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	C.gtk_cell_accessible_parent_expand_collapse(cValueInstance, cValue0)
}

func Fn_gtk_cell_accessible_parent_get_cell_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellAccessibleParent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	cValue1 := (*C.GdkRectangle)(unsafe.Pointer(param1))

	C.gtk_cell_accessible_parent_get_cell_area(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_accessible_parent_get_cell_extents(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int, param3 *int, param4 *int, param5 int) {
	cValueInstance := (*C.GtkCellAccessibleParent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	cValue3 := (*C.gint)(unsafe.Pointer(param3))

	cValue4 := (*C.gint)(unsafe.Pointer(param4))

	cValue5 := (C.AtkCoordType)(param5)

	C.gtk_cell_accessible_parent_get_cell_extents(cValueInstance, cValue0, cValue1, cValue2, cValue3, cValue4, cValue5)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted
func Fn_gtk_cell_accessible_parent_get_child_index(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkCellAccessibleParent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	ret := C.gtk_cell_accessible_parent_get_child_index(cValueInstance, cValue0)

	return (int)(ret)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : has array return

func Fn_gtk_cell_accessible_parent_get_renderer_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkCellAccessibleParent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	ret := C.gtk_cell_accessible_parent_get_renderer_state(cValueInstance, cValue0)

	return (int)(ret)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : has array return

func Fn_gtk_cell_accessible_parent_grab_focus(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkCellAccessibleParent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	ret := C.gtk_cell_accessible_parent_grab_focus(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_cell_accessible_parent_update_relationset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellAccessibleParent)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellAccessible)(unsafe.Pointer(param0))

	cValue1 := (*C.AtkRelationSet)(unsafe.Pointer(param1))

	C.gtk_cell_accessible_parent_update_relationset(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_editable_editing_done(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellEditable)(unsafe.Pointer(paramInstance))

	C.gtk_cell_editable_editing_done(cValueInstance)
}

func Fn_gtk_cell_editable_remove_widget(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellEditable)(unsafe.Pointer(paramInstance))

	C.gtk_cell_editable_remove_widget(cValueInstance)
}

func Fn_gtk_cell_editable_start_editing(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellEditable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkEvent)(unsafe.Pointer(param0))

	C.gtk_cell_editable_start_editing(cValueInstance, cValue0)
}

func Fn_gtk_cell_layout_add_attribute(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	cValue2 := (C.gint)(param2)

	C.gtk_cell_layout_add_attribute(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_cell_layout_clear(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	C.gtk_cell_layout_clear(cValueInstance)
}

func Fn_gtk_cell_layout_clear_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	C.gtk_cell_layout_clear_attributes(cValueInstance, cValue0)
}

func Fn_gtk_cell_layout_get_area(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_layout_get_area(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_layout_get_cells(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	ret := C.gtk_cell_layout_get_cells(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_layout_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_cell_layout_pack_end(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_layout_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_cell_layout_pack_start(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_cell_layout_reorder(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkCellLayout)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkCellRenderer)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_cell_layout_reorder(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_cell_layout_set_attributes : has varargs

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : has callback

// UNSUPPORTED : gtk_color_chooser_add_palette : has non-string array param colors

func Fn_gtk_editable_copy_clipboard(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	C.gtk_editable_copy_clipboard(cValueInstance)
}

func Fn_gtk_editable_cut_clipboard(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	C.gtk_editable_cut_clipboard(cValueInstance)
}

func Fn_gtk_editable_delete_selection(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	C.gtk_editable_delete_selection(cValueInstance)
}

func Fn_gtk_editable_delete_text(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_editable_delete_text(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_editable_get_chars(paramInstance unsafe.Pointer, param0 int, param1 int) string {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	ret := C.gtk_editable_get_chars(cValueInstance, cValue0, cValue1)

	return C.GoString(ret)
}

func Fn_gtk_editable_get_editable(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_editable_get_editable(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_editable_get_position(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_editable_get_position(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_editable_get_selection_bounds(paramInstance unsafe.Pointer, param0 *int, param1 *int) bool {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.gint)(unsafe.Pointer(param1))

	ret := C.gtk_editable_get_selection_bounds(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_editable_insert_text(paramInstance unsafe.Pointer, param0 string, param1 int, param2 *int) {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.gint)(unsafe.Pointer(param2))

	C.gtk_editable_insert_text(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_editable_paste_clipboard(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	C.gtk_editable_paste_clipboard(cValueInstance)
}

func Fn_gtk_editable_select_region(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.gint)(param1)

	C.gtk_editable_select_region(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_editable_set_editable(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_editable_set_editable(cValueInstance, cValue0)
}

func Fn_gtk_editable_set_position(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkEditable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_editable_set_position(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_add_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFileFilter)(unsafe.Pointer(param0))

	C.gtk_file_chooser_add_filter(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_add_shortcut_folder(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_add_shortcut_folder(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_add_shortcut_folder_uri(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_add_shortcut_folder_uri(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_action(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_action(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_file_chooser_get_create_folders(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_create_folders(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_current_folder(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_current_folder(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_current_folder_file(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_current_folder_file(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_current_folder_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_current_folder_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_do_overwrite_confirmation(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_do_overwrite_confirmation(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_extra_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_extra_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_file(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_file(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_filename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_filename(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_filenames(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_filenames(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_files(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_files(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_filter(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_filter(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_local_only(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_local_only(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_preview_file(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_file(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_preview_filename(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_filename(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_preview_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_preview_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_preview_widget_active(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_preview_widget_active(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_select_multiple(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_select_multiple(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_show_hidden(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_show_hidden(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_get_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_file_chooser_get_uris(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_uris(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_get_use_preview_label(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_get_use_preview_label(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_list_filters(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_list_filters(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_list_shortcut_folder_uris(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_list_shortcut_folder_uris(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_list_shortcut_folders(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_file_chooser_list_shortcut_folders(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_file_chooser_remove_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFileFilter)(unsafe.Pointer(param0))

	C.gtk_file_chooser_remove_filter(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_remove_shortcut_folder(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_remove_shortcut_folder(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_remove_shortcut_folder_uri(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_remove_shortcut_folder_uri(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_select_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	C.gtk_file_chooser_select_all(cValueInstance)
}

func Fn_gtk_file_chooser_select_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_select_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_select_filename(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_select_filename(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_select_uri(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_select_uri(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_action(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkFileChooserAction)(param0)

	C.gtk_file_chooser_set_action(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_create_folders(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_create_folders(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_current_folder(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_set_current_folder(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_current_folder_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_set_current_folder_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_current_folder_uri(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_set_current_folder_uri(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_current_name(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_set_current_name(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_do_overwrite_confirmation(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_do_overwrite_confirmation(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_extra_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_file_chooser_set_extra_widget(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	cError := (**C.GError)(error)

	ret := C.gtk_file_chooser_set_file(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_filename(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_set_filename(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkFileFilter)(unsafe.Pointer(param0))

	C.gtk_file_chooser_set_filter(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_local_only(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_local_only(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_preview_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_file_chooser_set_preview_widget(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_preview_widget_active(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_preview_widget_active(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_select_multiple(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_select_multiple(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_show_hidden(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_show_hidden(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_set_uri(paramInstance unsafe.Pointer, param0 string) bool {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	ret := C.gtk_file_chooser_set_uri(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_file_chooser_set_use_preview_label(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_file_chooser_set_use_preview_label(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_unselect_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	C.gtk_file_chooser_unselect_all(cValueInstance)
}

func Fn_gtk_file_chooser_unselect_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GFile)(unsafe.Pointer(param0))

	C.gtk_file_chooser_unselect_file(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_unselect_filename(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_unselect_filename(cValueInstance, cValue0)
}

func Fn_gtk_file_chooser_unselect_uri(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkFileChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.char)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_file_chooser_unselect_uri(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : has callback

func Fn_gtk_orientable_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkOrientable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_orientable_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_orientable_set_orientation(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkOrientable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkOrientation)(param0)

	C.gtk_orientable_set_orientation(cValueInstance, cValue0)
}

func Fn_gtk_print_operation_preview_end_preview(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkPrintOperationPreview)(unsafe.Pointer(paramInstance))

	C.gtk_print_operation_preview_end_preview(cValueInstance)
}

func Fn_gtk_print_operation_preview_is_selected(paramInstance unsafe.Pointer, param0 int) bool {
	cValueInstance := (*C.GtkPrintOperationPreview)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_print_operation_preview_is_selected(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_print_operation_preview_render_page(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkPrintOperationPreview)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_print_operation_preview_render_page(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_add_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentFilter)(unsafe.Pointer(param0))

	C.gtk_recent_chooser_add_filter(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_get_current_item(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_current_item(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_get_current_uri(paramInstance unsafe.Pointer) string {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_current_uri(cValueInstance)

	return C.GoString(ret)
}

func Fn_gtk_recent_chooser_get_filter(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_filter(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_get_items(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_items(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_get_limit(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_limit(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_recent_chooser_get_local_only(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_local_only(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_select_multiple(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_select_multiple(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_show_icons(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_show_icons(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_show_not_found(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_show_not_found(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_show_private(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_show_private(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_show_tips(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_show_tips(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_get_sort_type(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_get_sort_type(cValueInstance)

	return (int)(ret)
}

// UNSUPPORTED : gtk_recent_chooser_get_uris : has array return

func Fn_gtk_recent_chooser_list_filters(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_recent_chooser_list_filters(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_recent_chooser_remove_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentFilter)(unsafe.Pointer(param0))

	C.gtk_recent_chooser_remove_filter(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_select_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	C.gtk_recent_chooser_select_all(cValueInstance)
}

func Fn_gtk_recent_chooser_select_uri(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_chooser_select_uri(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_set_current_uri(paramInstance unsafe.Pointer, param0 string, error unsafe.Pointer) bool {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	cError := (**C.GError)(error)

	ret := C.gtk_recent_chooser_set_current_uri(cValueInstance, cValue0, cError)

	return toGoBool(ret)
}

func Fn_gtk_recent_chooser_set_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRecentFilter)(unsafe.Pointer(param0))

	C.gtk_recent_chooser_set_filter(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_limit(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_recent_chooser_set_limit(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_local_only(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_local_only(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_select_multiple(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_select_multiple(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_show_icons(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_show_icons(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_show_not_found(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_show_not_found(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_show_private(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_show_private(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_set_show_tips(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_recent_chooser_set_show_tips(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : has callback

func Fn_gtk_recent_chooser_set_sort_type(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkRecentSortType)(param0)

	C.gtk_recent_chooser_set_sort_type(cValueInstance, cValue0)
}

func Fn_gtk_recent_chooser_unselect_all(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	C.gtk_recent_chooser_unselect_all(cValueInstance)
}

func Fn_gtk_recent_chooser_unselect_uri(paramInstance unsafe.Pointer, param0 string) {
	cValueInstance := (*C.GtkRecentChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gchar)(C.CString(param0))
	defer C.free(unsafe.Pointer(cValue0))

	C.gtk_recent_chooser_unselect_uri(cValueInstance, cValue0)
}

func Fn_gtk_scrollable_get_hadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrollable_get_hadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrollable_get_hscroll_policy(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrollable_get_hscroll_policy(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrollable_get_vadjustment(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrollable_get_vadjustment(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrollable_get_vscroll_policy(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_scrollable_get_vscroll_policy(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_scrollable_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_scrollable_set_hadjustment(cValueInstance, cValue0)
}

func Fn_gtk_scrollable_set_hscroll_policy(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkScrollablePolicy)(param0)

	C.gtk_scrollable_set_hscroll_policy(cValueInstance, cValue0)
}

func Fn_gtk_scrollable_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkAdjustment)(unsafe.Pointer(param0))

	C.gtk_scrollable_set_vadjustment(cValueInstance, cValue0)
}

func Fn_gtk_scrollable_set_vscroll_policy(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkScrollable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.GtkScrollablePolicy)(param0)

	C.gtk_scrollable_set_vscroll_policy(cValueInstance, cValue0)
}

func Fn_gtk_style_provider_get_icon_factory(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidgetPath)(unsafe.Pointer(param0))

	ret := C.gtk_style_provider_get_icon_factory(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_provider_get_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkStyleProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidgetPath)(unsafe.Pointer(param0))

	ret := C.gtk_style_provider_get_style(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_provider_get_style_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStyleProvider)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidgetPath)(unsafe.Pointer(param0))

	cValue1 := (C.GtkStateFlags)(param1)

	cValue2 := (*C.GParamSpec)(unsafe.Pointer(param2))

	cValue3 := (*C.GValue)(unsafe.Pointer(param3))

	ret := C.gtk_style_provider_get_style_property(cValueInstance, cValue0, cValue1, cValue2, cValue3)

	return toGoBool(ret)
}

func Fn_gtk_tool_shell_get_ellipsize_mode(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_ellipsize_mode(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_icon_size(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_icon_size(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_relief_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_relief_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_style(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_style(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_text_alignment(paramInstance unsafe.Pointer) float32 {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_text_alignment(cValueInstance)

	return (float32)(ret)
}

func Fn_gtk_tool_shell_get_text_orientation(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_text_orientation(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tool_shell_get_text_size_group(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tool_shell_get_text_size_group(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tool_shell_rebuild_menu(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkToolShell)(unsafe.Pointer(paramInstance))

	C.gtk_tool_shell_rebuild_menu(cValueInstance)
}

func Fn_gtk_tree_drag_dest_drag_data_received(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeDragDest)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkSelectionData)(unsafe.Pointer(param1))

	ret := C.gtk_tree_drag_dest_drag_data_received(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_drag_dest_row_drop_possible(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeDragDest)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkSelectionData)(unsafe.Pointer(param1))

	ret := C.gtk_tree_drag_dest_row_drop_possible(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_drag_source_drag_data_delete(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeDragSource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_drag_source_drag_data_delete(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_drag_source_drag_data_get(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeDragSource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkSelectionData)(unsafe.Pointer(param1))

	ret := C.gtk_tree_drag_source_drag_data_get(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_drag_source_row_draggable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeDragSource)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_drag_source_row_draggable(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_filter_new(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_filter_new(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_model_foreach : has callback

// UNSUPPORTED : gtk_tree_model_get : has varargs

func Fn_gtk_tree_model_get_column_type(paramInstance unsafe.Pointer, param0 int) uint64 {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	ret := C.gtk_tree_model_get_column_type(cValueInstance, cValue0)

	return (uint64)(ret)
}

func Fn_gtk_tree_model_get_flags(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_model_get_flags(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_model_get_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreePath)(unsafe.Pointer(param1))

	ret := C.gtk_tree_model_get_iter(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_get_iter_first(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_get_iter_first(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_get_iter_from_string(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.gchar)(C.CString(param1))
	defer C.free(unsafe.Pointer(cValue1))

	ret := C.gtk_tree_model_get_iter_from_string(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_get_n_columns(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_model_get_n_columns(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_tree_model_get_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_get_path(cValueInstance, cValue0)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_get_string_from_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) string {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_get_string_from_iter(cValueInstance, cValue0)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_tree_model_get_valist : has va_list

func Fn_gtk_tree_model_get_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	cValue2 := (*C.GValue)(unsafe.Pointer(param2))

	C.gtk_tree_model_get_value(cValueInstance, cValue0, cValue1, cValue2)
}

func Fn_gtk_tree_model_iter_children(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	ret := C.gtk_tree_model_iter_children(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_iter_has_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_iter_has_child(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_iter_n_children(paramInstance unsafe.Pointer, param0 unsafe.Pointer) int {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_iter_n_children(cValueInstance, cValue0)

	return (int)(ret)
}

func Fn_gtk_tree_model_iter_next(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_iter_next(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_iter_nth_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	cValue2 := (C.gint)(param2)

	ret := C.gtk_tree_model_iter_nth_child(cValueInstance, cValue0, cValue1, cValue2)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_iter_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	ret := C.gtk_tree_model_iter_parent(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_iter_previous(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	ret := C.gtk_tree_model_iter_previous(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_ref_node(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_tree_model_ref_node(cValueInstance, cValue0)
}

func Fn_gtk_tree_model_row_changed(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_row_changed(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_model_row_deleted(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	C.gtk_tree_model_row_deleted(cValueInstance, cValue0)
}

func Fn_gtk_tree_model_row_has_child_toggled(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_row_has_child_toggled(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_tree_model_row_inserted(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreePath)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkTreeIter)(unsafe.Pointer(param1))

	C.gtk_tree_model_row_inserted(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_tree_model_rows_reordered_with_length : has non-string array param new_order

func Fn_gtk_tree_model_sort_new_with_model(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_model_sort_new_with_model(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_unref_node(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeModel)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkTreeIter)(unsafe.Pointer(param0))

	C.gtk_tree_model_unref_node(cValueInstance, cValue0)
}

func Fn_gtk_tree_sortable_get_sort_column_id(paramInstance unsafe.Pointer, param0 *int, param1 *int) bool {
	cValueInstance := (*C.GtkTreeSortable)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.gint)(unsafe.Pointer(param0))

	cValue1 := (*C.GtkSortType)(unsafe.Pointer(param1))

	ret := C.gtk_tree_sortable_get_sort_column_id(cValueInstance, cValue0, cValue1)

	return toGoBool(ret)
}

func Fn_gtk_tree_sortable_has_default_sort_func(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkTreeSortable)(unsafe.Pointer(paramInstance))

	ret := C.gtk_tree_sortable_has_default_sort_func(cValueInstance)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : has callback

func Fn_gtk_tree_sortable_set_sort_column_id(paramInstance unsafe.Pointer, param0 int, param1 int) {
	cValueInstance := (*C.GtkTreeSortable)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	cValue1 := (C.GtkSortType)(param1)

	C.gtk_tree_sortable_set_sort_column_id(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : has callback

func Fn_gtk_tree_sortable_sort_column_changed(paramInstance unsafe.Pointer) {
	cValueInstance := (*C.GtkTreeSortable)(unsafe.Pointer(paramInstance))

	C.gtk_tree_sortable_sort_column_changed(cValueInstance)
}
