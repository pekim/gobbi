// Code generated - DO NOT EDIT.
// +build gtk_3.12

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
import "C"

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
type CssSection C.GtkCssSection
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

func Fn_gtk_accel_groups_activate(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_accel_groups_from_object(param0 unsafe.Pointer) {}

func Fn_gtk_accelerator_get_default_mod_mask() {
	C.gtk_accelerator_get_default_mod_mask()
}

func Fn_gtk_accelerator_get_label(param0 uint, param1 int) {}

func Fn_gtk_accelerator_get_label_with_keycode(param0 unsafe.Pointer, param1 uint, param2 uint, param3 int) {
}

func Fn_gtk_accelerator_name(param0 uint, param1 int) {}

func Fn_gtk_accelerator_name_with_keycode(param0 unsafe.Pointer, param1 uint, param2 uint, param3 int) {
}

func Fn_gtk_accelerator_parse(param0 string, param1 *uint, param2 int) {}

func Fn_gtk_accelerator_parse_with_keycode(param0 string, param1 *uint, param2 []*uint, param3 int) {
}

func Fn_gtk_accelerator_set_default_mod_mask(param0 int) {}

func Fn_gtk_accelerator_valid(param0 uint, param1 int) {}

func Fn_gtk_alternative_dialog_button_order(param0 unsafe.Pointer) {}

func Fn_gtk_binding_entry_add_signal_from_string(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_binding_entry_add_signall(param0 unsafe.Pointer, param1 uint, param2 int, param3 string, param4 unsafe.Pointer) {
}

func Fn_gtk_binding_entry_remove(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_binding_entry_skip(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_binding_set_by_class(param0 *unsafe.Pointer) {}

func Fn_gtk_binding_set_find(param0 string) {}

func Fn_gtk_binding_set_new(param0 string) {}

func Fn_gtk_bindings_activate(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_bindings_activate_event(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_builder_error_quark() {
	C.gtk_builder_error_quark()
}

func Fn_gtk_cairo_should_draw_window(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_cairo_transform_to_window(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_check_version(param0 uint, param1 uint, param2 uint) {}

func Fn_gtk_css_provider_error_quark() {
	C.gtk_css_provider_error_quark()
}

func Fn_gtk_device_grab_add(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {}

func Fn_gtk_device_grab_remove(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_disable_setlocale() {
	C.gtk_disable_setlocale()
}

func Fn_gtk_distribute_natural_allocation(param0 int, param1 uint, param2 unsafe.Pointer) {}

func Fn_gtk_drag_finish(param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint32) {}

func Fn_gtk_drag_get_source_widget(param0 unsafe.Pointer) {}

func Fn_gtk_drag_set_icon_default(param0 unsafe.Pointer) {}

func Fn_gtk_drag_set_icon_gicon(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
}

func Fn_gtk_drag_set_icon_name(param0 unsafe.Pointer, param1 string, param2 int, param3 int) {}

func Fn_gtk_drag_set_icon_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
}

func Fn_gtk_drag_set_icon_stock(param0 unsafe.Pointer, param1 string, param2 int, param3 int) {}

func Fn_gtk_drag_set_icon_surface(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_drag_set_icon_widget(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
}

func Fn_gtk_draw_insertion_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool, param4 int, param5 bool) {
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

func Fn_gtk_get_binary_age() {
	C.gtk_get_binary_age()
}

func Fn_gtk_get_current_event() {
	C.gtk_get_current_event()
}

func Fn_gtk_get_current_event_device() {
	C.gtk_get_current_event_device()
}

func Fn_gtk_get_current_event_state(param0 int) {}

func Fn_gtk_get_current_event_time() {
	C.gtk_get_current_event_time()
}

func Fn_gtk_get_debug_flags() {
	C.gtk_get_debug_flags()
}

func Fn_gtk_get_default_language() {
	C.gtk_get_default_language()
}

func Fn_gtk_get_event_widget(param0 unsafe.Pointer) {}

func Fn_gtk_get_interface_age() {
	C.gtk_get_interface_age()
}

func Fn_gtk_get_locale_direction() {
	C.gtk_get_locale_direction()
}

func Fn_gtk_get_major_version() {
	C.gtk_get_major_version()
}

func Fn_gtk_get_micro_version() {
	C.gtk_get_micro_version()
}

func Fn_gtk_get_minor_version() {
	C.gtk_get_minor_version()
}

func Fn_gtk_get_option_group(param0 bool) {}

func Fn_gtk_grab_get_current() {
	C.gtk_grab_get_current()
}

func Fn_gtk_icon_size_from_name(param0 string) {}

func Fn_gtk_icon_size_get_name(param0 int) {}

func Fn_gtk_icon_size_lookup(param0 int, param1 *int, param2 *int) {}

func Fn_gtk_icon_size_lookup_for_settings(param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_icon_size_register(param0 string, param1 int, param2 int) {}

func Fn_gtk_icon_size_register_alias(param0 string, param1 int) {}

func Fn_gtk_icon_theme_error_quark() {
	C.gtk_icon_theme_error_quark()
}

func Fn_gtk_init(param0 *int, param1 *[]string) {}

func Fn_gtk_init_check(param0 *int, param1 *[]string) {}

func Fn_gtk_init_with_args(param0 *int, param1 *[]string, param2 string, param3 []glib.OptionEntry, param4 string) {
}

// UNSUPPORTED : key_snooper_install : has callback

func Fn_gtk_key_snooper_remove(param0 uint) {}

func Fn_gtk_main() {
	C.gtk_main()
}

func Fn_gtk_main_do_event(param0 unsafe.Pointer) {}

func Fn_gtk_main_iteration() {
	C.gtk_main_iteration()
}

func Fn_gtk_main_iteration_do(param0 bool) {}

func Fn_gtk_main_level() {
	C.gtk_main_level()
}

func Fn_gtk_main_quit() {
	C.gtk_main_quit()
}

func Fn_gtk_paint_arrow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 bool, param8 int, param9 int, param10 int, param11 int) {
}

func Fn_gtk_paint_box(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_box_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int, param11 int, param12 int) {
}

func Fn_gtk_paint_check(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_diamond(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_expander(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int) {
}

func Fn_gtk_paint_extension(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int) {
}

func Fn_gtk_paint_flat_box(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_focus(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int, param8 int) {
}

func Fn_gtk_paint_handle(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int) {
}

func Fn_gtk_paint_hline(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int) {
}

func Fn_gtk_paint_layout(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 bool, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 unsafe.Pointer) {
}

func Fn_gtk_paint_option(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_resize_grip(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_shadow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_shadow_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int, param11 int, param12 int) {
}

func Fn_gtk_paint_slider(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int) {
}

func Fn_gtk_paint_spinner(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 uint, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_tab(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_vline(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int) {
}

func Fn_gtk_paper_size_get_default() {
	C.gtk_paper_size_get_default()
}

func Fn_gtk_paper_size_get_paper_sizes(param0 bool) {}

func Fn_gtk_parse_args(param0 *int, param1 *[]string) {}

func Fn_gtk_print_error_quark() {
	C.gtk_print_error_quark()
}

func Fn_gtk_print_run_page_setup_dialog(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

// UNSUPPORTED : print_run_page_setup_dialog_async : has callback

func Fn_gtk_propagate_event(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_rc_add_default_file(param0 string) {}

func Fn_gtk_rc_find_module_in_path(param0 string) {}

func Fn_gtk_rc_find_pixmap_in_path(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) {}

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

func Fn_gtk_rc_get_style(param0 unsafe.Pointer) {}

func Fn_gtk_rc_get_style_by_paths(param0 unsafe.Pointer, param1 string, param2 string, param3 uint64) {
}

func Fn_gtk_rc_get_theme_dir() {
	C.gtk_rc_get_theme_dir()
}

func Fn_gtk_rc_parse(param0 string) {}

func Fn_gtk_rc_parse_color(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_rc_parse_color_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_parse_priority(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_rc_parse_state(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_rc_parse_string(param0 string) {}

func Fn_gtk_rc_property_parse_border(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_property_parse_color(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_property_parse_enum(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_property_parse_flags(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_property_parse_requisition(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_reparse_all() {
	C.gtk_rc_reparse_all()
}

func Fn_gtk_rc_reparse_all_for_settings(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_rc_reset_styles(param0 unsafe.Pointer) {}

func Fn_gtk_rc_scanner_new() {
	C.gtk_rc_scanner_new()
}

func Fn_gtk_rc_set_default_files(param0 []string) {}

func Fn_gtk_recent_chooser_error_quark() {
	C.gtk_recent_chooser_error_quark()
}

func Fn_gtk_recent_manager_error_quark() {
	C.gtk_recent_manager_error_quark()
}

func Fn_gtk_render_activity(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_arrow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_background(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_check(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_expander(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_extension(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int) {
}

func Fn_gtk_render_focus(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_frame(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_frame_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int, param7 float64, param8 float64) {
}

func Fn_gtk_render_handle(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_icon(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 float64, param4 float64) {
}

func Fn_gtk_render_icon_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {}

func Fn_gtk_render_icon_surface(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 float64, param4 float64) {
}

func Fn_gtk_render_insertion_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 unsafe.Pointer, param5 int, param6 int) {
}

func Fn_gtk_render_layout(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 unsafe.Pointer) {
}

func Fn_gtk_render_line(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_option(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_slider(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int) {
}

func Fn_gtk_rgb_to_hsv(param0 float64, param1 float64, param2 float64, param3 *float64, param4 *float64, param5 *float64) {
}

func Fn_gtk_selection_add_target(param0 unsafe.Pointer, param1 gdk.Atom, param2 gdk.Atom, param3 uint) {
}

func Fn_gtk_selection_add_targets(param0 unsafe.Pointer, param1 gdk.Atom, param2 []TargetEntry, param3 uint) {
}

func Fn_gtk_selection_clear_targets(param0 unsafe.Pointer, param1 gdk.Atom) {}

func Fn_gtk_selection_convert(param0 unsafe.Pointer, param1 gdk.Atom, param2 gdk.Atom, param3 uint32) {
}

func Fn_gtk_selection_owner_set(param0 unsafe.Pointer, param1 gdk.Atom, param2 uint32) {}

func Fn_gtk_selection_owner_set_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 gdk.Atom, param3 uint32) {
}

func Fn_gtk_selection_remove_all(param0 unsafe.Pointer) {}

func Fn_gtk_set_debug_flags(param0 uint) {}

// UNSUPPORTED : show_about_dialog : has varargs

func Fn_gtk_show_uri(param0 unsafe.Pointer, param1 string, param2 uint32) {}

func Fn_gtk_stock_add(param0 []StockItem, param1 uint) {}

func Fn_gtk_stock_add_static(param0 []StockItem, param1 uint) {}

func Fn_gtk_stock_list_ids() {
	C.gtk_stock_list_ids()
}

func Fn_gtk_stock_lookup(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : stock_set_translate_func : has callback

func Fn_gtk_target_table_free(param0 []TargetEntry, param1 int) {}

func Fn_gtk_target_table_new_from_list(param0 unsafe.Pointer, param1 *int) {}

func Fn_gtk_targets_include_image(param0 []gdk.Atom, param1 int, param2 bool) {}

func Fn_gtk_targets_include_rich_text(param0 []gdk.Atom, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_targets_include_text(param0 []gdk.Atom, param1 int) {}

func Fn_gtk_targets_include_uri(param0 []gdk.Atom, param1 int) {}

func Fn_gtk_test_create_simple_window(param0 string, param1 string) {}

// UNSUPPORTED : test_create_widget : has varargs

// UNSUPPORTED : test_display_button_window : has varargs

func Fn_gtk_test_find_label(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_test_find_sibling(param0 unsafe.Pointer, param1 uint64) {}

func Fn_gtk_test_find_widget(param0 unsafe.Pointer, param1 string, param2 uint64) {}

// UNSUPPORTED : test_init : has varargs

func Fn_gtk_test_list_all_types(param0 *uint) {}

func Fn_gtk_test_register_all_types() {
	C.gtk_test_register_all_types()
}

func Fn_gtk_test_slider_get_value(param0 unsafe.Pointer) {}

func Fn_gtk_test_slider_set_perc(param0 unsafe.Pointer, param1 float64) {}

func Fn_gtk_test_spin_button_click(param0 unsafe.Pointer, param1 uint, param2 bool) {}

func Fn_gtk_test_text_get(param0 unsafe.Pointer) {}

func Fn_gtk_test_text_set(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_test_widget_click(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_test_widget_send_key(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_test_widget_wait_for_draw(param0 unsafe.Pointer) {}

func Fn_gtk_tree_get_row_drag_data(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer) {
}

func Fn_gtk_tree_row_reference_deleted(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_row_reference_inserted(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_row_reference_reordered(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 []int) {
}

func Fn_gtk_tree_set_row_drag_data(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_true() {
	C.gtk_true()
}

func Fn_gtk_about_dialog_new() {
	C.gtk_about_dialog_new()
}

func Fn_gtk_about_dialog_add_credit_section(paramInstance unsafe.Pointer, param0 string, param1 []string) {
}

func Fn_gtk_about_dialog_get_artists(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_authors(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_comments(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_copyright(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_documenters(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_license(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_license_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_logo(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_logo_icon_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_program_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_translator_credits(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_version(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_website(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_website_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_get_wrap_license(paramInstance unsafe.Pointer) {}

func Fn_gtk_about_dialog_set_artists(paramInstance unsafe.Pointer, param0 []string) {}

func Fn_gtk_about_dialog_set_authors(paramInstance unsafe.Pointer, param0 []string) {}

func Fn_gtk_about_dialog_set_comments(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_about_dialog_set_copyright(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_about_dialog_set_documenters(paramInstance unsafe.Pointer, param0 []string) {}

func Fn_gtk_about_dialog_set_license(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_about_dialog_set_license_type(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_about_dialog_set_logo(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_about_dialog_set_logo_icon_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_about_dialog_set_program_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_about_dialog_set_translator_credits(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_about_dialog_set_version(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_about_dialog_set_website(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_about_dialog_set_website_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_about_dialog_set_wrap_license(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_accel_group_new() {
	C.gtk_accel_group_new()
}

func Fn_gtk_accel_group_activate(paramInstance unsafe.Pointer, param0 uint32, param1 unsafe.Pointer, param2 uint, param3 int) {
}

func Fn_gtk_accel_group_connect(paramInstance unsafe.Pointer, param0 uint, param1 int, param2 int, param3 unsafe.Pointer) {
}

func Fn_gtk_accel_group_connect_by_path(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_accel_group_disconnect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_accel_group_disconnect_key(paramInstance unsafe.Pointer, param0 uint, param1 int) {}

// UNSUPPORTED : find : has callback

func Fn_gtk_accel_group_get_is_locked(paramInstance unsafe.Pointer) {}

func Fn_gtk_accel_group_get_modifier_mask(paramInstance unsafe.Pointer) {}

func Fn_gtk_accel_group_lock(paramInstance unsafe.Pointer) {}

func Fn_gtk_accel_group_query(paramInstance unsafe.Pointer, param0 uint, param1 int, param2 *uint) {}

func Fn_gtk_accel_group_unlock(paramInstance unsafe.Pointer) {}

func Fn_gtk_accel_group_from_accel_closure(param0 unsafe.Pointer) {}

func Fn_gtk_accel_label_new(param0 string) {}

func Fn_gtk_accel_label_get_accel(paramInstance unsafe.Pointer, param0 *uint, param1 int) {}

func Fn_gtk_accel_label_get_accel_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_accel_label_get_accel_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_accel_label_refetch(paramInstance unsafe.Pointer) {}

func Fn_gtk_accel_label_set_accel(paramInstance unsafe.Pointer, param0 uint, param1 int) {}

func Fn_gtk_accel_label_set_accel_closure(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_accel_label_set_accel_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_accel_map_add_entry(param0 string, param1 uint, param2 int) {}

func Fn_gtk_accel_map_add_filter(param0 string) {}

func Fn_gtk_accel_map_change_entry(param0 string, param1 uint, param2 int, param3 bool) {}

// UNSUPPORTED : foreach : has callback

// UNSUPPORTED : foreach_unfiltered : has callback

func Fn_gtk_accel_map_get() {
	C.gtk_accel_map_get()
}

func Fn_gtk_accel_map_load(param0 string) {}

func Fn_gtk_accel_map_load_fd(param0 int) {}

func Fn_gtk_accel_map_load_scanner(param0 unsafe.Pointer) {}

func Fn_gtk_accel_map_lock_path(param0 string) {}

func Fn_gtk_accel_map_lookup_entry(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_accel_map_save(param0 string) {}

func Fn_gtk_accel_map_save_fd(param0 int) {}

func Fn_gtk_accel_map_unlock_path(param0 string) {}

func Fn_gtk_accessible_connect_widget_destroyed(paramInstance unsafe.Pointer) {}

func Fn_gtk_accessible_get_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_accessible_set_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_action_new(param0 string, param1 string, param2 string, param3 string) {}

func Fn_gtk_action_activate(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_block_activate(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_connect_accelerator(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_create_icon(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_action_create_menu(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_create_menu_item(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_create_tool_item(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_disconnect_accelerator(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_accel_closure(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_accel_path(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_always_show_image(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_gicon(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_icon_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_is_important(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_proxies(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_sensitive(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_short_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_stock_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_tooltip(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_visible_horizontal(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_get_visible_vertical(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_is_sensitive(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_is_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_action_set_accel_path(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_action_set_always_show_image(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_action_set_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_action_set_icon_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_action_set_is_important(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_action_set_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_action_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_action_set_short_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_action_set_stock_id(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_action_set_tooltip(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_action_set_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_action_set_visible_horizontal(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_action_set_visible_vertical(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_action_unblock_activate(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_bar_new() {
	C.gtk_action_bar_new()
}

func Fn_gtk_action_bar_get_center_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_bar_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_action_bar_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_action_bar_set_center_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_action_group_new(param0 string) {}

func Fn_gtk_action_group_add_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_action_group_add_action_with_accel(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

func Fn_gtk_action_group_add_actions(paramInstance unsafe.Pointer, param0 []ActionEntry, param1 uint, param2 *unsafe.Pointer) {
}

// UNSUPPORTED : add_actions_full : has callback

// UNSUPPORTED : add_radio_actions : has callback

// UNSUPPORTED : add_radio_actions_full : has callback

func Fn_gtk_action_group_add_toggle_actions(paramInstance unsafe.Pointer, param0 []ToggleActionEntry, param1 uint, param2 *unsafe.Pointer) {
}

// UNSUPPORTED : add_toggle_actions_full : has callback

func Fn_gtk_action_group_get_accel_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_group_get_action(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_action_group_get_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_group_get_sensitive(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_group_get_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_group_list_actions(paramInstance unsafe.Pointer) {}

func Fn_gtk_action_group_remove_action(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_action_group_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_action_group_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {}

// UNSUPPORTED : set_translate_func : has callback

func Fn_gtk_action_group_set_translation_domain(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_action_group_set_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_action_group_translate_string(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_adjustment_new(param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_adjustment_changed(paramInstance unsafe.Pointer) {}

func Fn_gtk_adjustment_clamp_page(paramInstance unsafe.Pointer, param0 float64, param1 float64) {}

func Fn_gtk_adjustment_configure(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_adjustment_get_lower(paramInstance unsafe.Pointer) {}

func Fn_gtk_adjustment_get_minimum_increment(paramInstance unsafe.Pointer) {}

func Fn_gtk_adjustment_get_page_increment(paramInstance unsafe.Pointer) {}

func Fn_gtk_adjustment_get_page_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_adjustment_get_step_increment(paramInstance unsafe.Pointer) {}

func Fn_gtk_adjustment_get_upper(paramInstance unsafe.Pointer) {}

func Fn_gtk_adjustment_get_value(paramInstance unsafe.Pointer) {}

func Fn_gtk_adjustment_set_lower(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_adjustment_set_page_increment(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_adjustment_set_page_size(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_adjustment_set_step_increment(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_adjustment_set_upper(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_adjustment_set_value(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_adjustment_value_changed(paramInstance unsafe.Pointer) {}

func Fn_gtk_alignment_new(param0 float32, param1 float32, param2 float32, param3 float32) {}

func Fn_gtk_alignment_get_padding(paramInstance unsafe.Pointer, param0 *uint, param1 *uint, param2 *uint, param3 *uint) {
}

func Fn_gtk_alignment_set(paramInstance unsafe.Pointer, param0 float32, param1 float32, param2 float32, param3 float32) {
}

func Fn_gtk_alignment_set_padding(paramInstance unsafe.Pointer, param0 uint, param1 uint, param2 uint, param3 uint) {
}

func Fn_gtk_app_chooser_button_new(param0 string) {}

func Fn_gtk_app_chooser_button_append_custom_item(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer) {
}

func Fn_gtk_app_chooser_button_append_separator(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_button_get_heading(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_button_get_show_default_item(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_button_get_show_dialog_item(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_button_set_active_custom_item(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_app_chooser_button_set_heading(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_app_chooser_button_set_show_default_item(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_app_chooser_button_set_show_dialog_item(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_app_chooser_dialog_new(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_app_chooser_dialog_new_for_content_type(param0 unsafe.Pointer, param1 int, param2 string) {
}

func Fn_gtk_app_chooser_dialog_get_heading(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_dialog_get_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_dialog_set_heading(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_app_chooser_widget_new(param0 string) {}

func Fn_gtk_app_chooser_widget_get_default_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_widget_get_show_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_widget_get_show_default(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_widget_get_show_fallback(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_widget_get_show_other(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_widget_get_show_recommended(paramInstance unsafe.Pointer) {}

func Fn_gtk_app_chooser_widget_set_default_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_app_chooser_widget_set_show_all(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_app_chooser_widget_set_show_default(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_app_chooser_widget_set_show_fallback(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_app_chooser_widget_set_show_other(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_app_chooser_widget_set_show_recommended(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_application_new(param0 string, param1 int) {}

func Fn_gtk_application_add_accelerator(paramInstance unsafe.Pointer, param0 string, param1 string, param2 unsafe.Pointer) {
}

func Fn_gtk_application_add_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_application_get_accels_for_action(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_application_get_active_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_application_get_app_menu(paramInstance unsafe.Pointer) {}

func Fn_gtk_application_get_menubar(paramInstance unsafe.Pointer) {}

func Fn_gtk_application_get_window_by_id(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_application_get_windows(paramInstance unsafe.Pointer) {}

func Fn_gtk_application_inhibit(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 string) {
}

func Fn_gtk_application_is_inhibited(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_application_list_action_descriptions(paramInstance unsafe.Pointer) {}

func Fn_gtk_application_remove_accelerator(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_application_remove_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_application_set_accels_for_action(paramInstance unsafe.Pointer, param0 string, param1 []string) {
}

func Fn_gtk_application_set_app_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_application_set_menubar(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_application_uninhibit(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_application_window_new(param0 unsafe.Pointer) {}

func Fn_gtk_application_window_get_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_application_window_get_show_menubar(paramInstance unsafe.Pointer) {}

func Fn_gtk_application_window_set_show_menubar(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_arrow_new(param0 int, param1 int) {}

func Fn_gtk_arrow_set(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_aspect_frame_new(param0 string, param1 float32, param2 float32, param3 float32, param4 bool) {
}

func Fn_gtk_aspect_frame_set(paramInstance unsafe.Pointer, param0 float32, param1 float32, param2 float32, param3 bool) {
}

func Fn_gtk_assistant_new() {
	C.gtk_assistant_new()
}

func Fn_gtk_assistant_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_assistant_append_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_assistant_commit(paramInstance unsafe.Pointer) {}

func Fn_gtk_assistant_get_current_page(paramInstance unsafe.Pointer) {}

func Fn_gtk_assistant_get_n_pages(paramInstance unsafe.Pointer) {}

func Fn_gtk_assistant_get_nth_page(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_assistant_get_page_complete(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_assistant_get_page_header_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_assistant_get_page_side_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_assistant_get_page_title(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_assistant_get_page_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_assistant_insert_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_assistant_next_page(paramInstance unsafe.Pointer) {}

func Fn_gtk_assistant_prepend_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_assistant_previous_page(paramInstance unsafe.Pointer) {}

func Fn_gtk_assistant_remove_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_assistant_remove_page(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_assistant_set_current_page(paramInstance unsafe.Pointer, param0 int) {}

// UNSUPPORTED : set_forward_page_func : has callback

func Fn_gtk_assistant_set_page_complete(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_assistant_set_page_header_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_assistant_set_page_side_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_assistant_set_page_title(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

func Fn_gtk_assistant_set_page_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_assistant_update_buttons_state(paramInstance unsafe.Pointer) {}

func Fn_gtk_bin_get_child(paramInstance unsafe.Pointer) {}

func Fn_gtk_box_new(param0 int, param1 int) {}

func Fn_gtk_box_get_baseline_position(paramInstance unsafe.Pointer) {}

func Fn_gtk_box_get_center_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_box_get_homogeneous(paramInstance unsafe.Pointer) {}

func Fn_gtk_box_get_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_box_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint) {
}

func Fn_gtk_box_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint) {
}

func Fn_gtk_box_query_child_packing(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *bool, param2 *bool, param3 *uint, param4 int) {
}

func Fn_gtk_box_reorder_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_box_set_baseline_position(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_box_set_center_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_box_set_child_packing(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint, param4 int) {
}

func Fn_gtk_box_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_box_set_spacing(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_builder_new() {
	C.gtk_builder_new()
}

func Fn_gtk_builder_new_from_file(param0 string) {}

func Fn_gtk_builder_new_from_resource(param0 string) {}

func Fn_gtk_builder_new_from_string(param0 string, param1 uint64) {}

// UNSUPPORTED : add_callback_symbol : has callback

// UNSUPPORTED : add_callback_symbols : has varargs

func Fn_gtk_builder_add_from_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_builder_add_from_resource(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_builder_add_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64) {}

func Fn_gtk_builder_add_objects_from_file(paramInstance unsafe.Pointer, param0 string, param1 []string) {
}

func Fn_gtk_builder_add_objects_from_resource(paramInstance unsafe.Pointer, param0 string, param1 []string) {
}

func Fn_gtk_builder_add_objects_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64, param2 []string) {
}

func Fn_gtk_builder_connect_signals(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {}

// UNSUPPORTED : connect_signals_full : has callback

func Fn_gtk_builder_expose_object(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_builder_extend_with_template(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint64, param2 string, param3 uint64) {
}

func Fn_gtk_builder_get_application(paramInstance unsafe.Pointer) {}

func Fn_gtk_builder_get_object(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_builder_get_objects(paramInstance unsafe.Pointer) {}

func Fn_gtk_builder_get_translation_domain(paramInstance unsafe.Pointer) {}

func Fn_gtk_builder_get_type_from_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_builder_lookup_callback_symbol(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_builder_set_application(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_builder_set_translation_domain(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_builder_value_from_string(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
}

func Fn_gtk_builder_value_from_string_type(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer) {
}

func Fn_gtk_button_new() {
	C.gtk_button_new()
}

func Fn_gtk_button_new_from_icon_name(param0 string, param1 int) {}

func Fn_gtk_button_new_from_stock(param0 string) {}

func Fn_gtk_button_new_with_label(param0 string) {}

func Fn_gtk_button_new_with_mnemonic(param0 string) {}

func Fn_gtk_button_clicked(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_enter(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_get_alignment(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {}

func Fn_gtk_button_get_always_show_image(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_get_event_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_get_focus_on_click(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_get_image(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_get_image_position(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_get_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_get_relief(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_get_use_stock(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_get_use_underline(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_leave(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_pressed(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_released(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_set_alignment(paramInstance unsafe.Pointer, param0 float32, param1 float32) {}

func Fn_gtk_button_set_always_show_image(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_button_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_button_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_button_set_image_position(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_button_set_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_button_set_relief(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_button_set_use_stock(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_button_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_button_box_new(param0 int) {}

func Fn_gtk_button_box_get_child_non_homogeneous(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_button_box_get_child_secondary(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_button_box_get_layout(paramInstance unsafe.Pointer) {}

func Fn_gtk_button_box_set_child_non_homogeneous(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_button_box_set_child_secondary(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_button_box_set_layout(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_calendar_new() {
	C.gtk_calendar_new()
}

func Fn_gtk_calendar_clear_marks(paramInstance unsafe.Pointer) {}

func Fn_gtk_calendar_get_date(paramInstance unsafe.Pointer, param0 *uint, param1 *uint, param2 *uint) {
}

func Fn_gtk_calendar_get_day_is_marked(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_calendar_get_detail_height_rows(paramInstance unsafe.Pointer) {}

func Fn_gtk_calendar_get_detail_width_chars(paramInstance unsafe.Pointer) {}

func Fn_gtk_calendar_get_display_options(paramInstance unsafe.Pointer) {}

func Fn_gtk_calendar_mark_day(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_calendar_select_day(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_calendar_select_month(paramInstance unsafe.Pointer, param0 uint, param1 uint) {}

// UNSUPPORTED : set_detail_func : has callback

func Fn_gtk_calendar_set_detail_height_rows(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_calendar_set_detail_width_chars(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_calendar_set_display_options(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_calendar_unmark_day(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_cell_area_activate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 bool) {
}

func Fn_gtk_cell_area_activate_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_gtk_cell_area_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_add_focus_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

// UNSUPPORTED : add_with_properties : has varargs

func Fn_gtk_cell_area_apply_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 bool) {
}

func Fn_gtk_cell_area_attribute_connect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
}

func Fn_gtk_cell_area_attribute_disconnect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

// UNSUPPORTED : cell_get : has varargs

func Fn_gtk_cell_area_cell_get_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
}

// UNSUPPORTED : cell_get_valist : has va_list

// UNSUPPORTED : cell_set : has varargs

func Fn_gtk_cell_area_cell_set_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
}

// UNSUPPORTED : cell_set_valist : has va_list

func Fn_gtk_cell_area_copy_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_create_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_area_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_gtk_cell_area_focus(paramInstance unsafe.Pointer, param0 int) {}

// UNSUPPORTED : foreach : has callback

// UNSUPPORTED : foreach_alloc : has callback

func Fn_gtk_cell_area_get_cell_allocation(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer) {
}

func Fn_gtk_cell_area_get_cell_at_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 int, param5 unsafe.Pointer) {
}

func Fn_gtk_cell_area_get_current_path_string(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_area_get_edit_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_area_get_edited_cell(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_area_get_focus_cell(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_area_get_focus_from_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_get_focus_siblings(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_get_preferred_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int) {
}

func Fn_gtk_cell_area_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 *int, param4 *int) {
}

func Fn_gtk_cell_area_get_preferred_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int) {
}

func Fn_gtk_cell_area_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 *int, param4 *int) {
}

func Fn_gtk_cell_area_get_request_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_area_has_renderer(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_inner_cell_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_cell_area_is_activatable(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_area_is_focus_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_cell_area_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_remove_focus_sibling(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_cell_area_render(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int, param6 bool) {
}

func Fn_gtk_cell_area_request_renderer(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 int, param4 *int, param5 *int) {
}

func Fn_gtk_cell_area_set_focus_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_stop_editing(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_cell_area_box_new() {
	C.gtk_cell_area_box_new()
}

func Fn_gtk_cell_area_box_get_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_area_box_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 bool) {
}

func Fn_gtk_cell_area_box_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool, param3 bool) {
}

func Fn_gtk_cell_area_box_set_spacing(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_cell_area_context_allocate(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_cell_area_context_get_allocation(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
}

func Fn_gtk_cell_area_context_get_area(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_area_context_get_preferred_height(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
}

func Fn_gtk_cell_area_context_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
}

func Fn_gtk_cell_area_context_get_preferred_width(paramInstance unsafe.Pointer, param0 *int, param1 *int) {
}

func Fn_gtk_cell_area_context_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
}

func Fn_gtk_cell_area_context_push_preferred_height(paramInstance unsafe.Pointer, param0 int, param1 int) {
}

func Fn_gtk_cell_area_context_push_preferred_width(paramInstance unsafe.Pointer, param0 int, param1 int) {
}

func Fn_gtk_cell_area_context_reset(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_renderer_activate(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) {
}

func Fn_gtk_cell_renderer_get_aligned_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {
}

func Fn_gtk_cell_renderer_get_alignment(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {
}

func Fn_gtk_cell_renderer_get_fixed_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_cell_renderer_get_padding(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_cell_renderer_get_preferred_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
}

func Fn_gtk_cell_renderer_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_cell_renderer_get_preferred_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_cell_renderer_get_preferred_width(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
}

func Fn_gtk_cell_renderer_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_cell_renderer_get_request_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_renderer_get_sensitive(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_renderer_get_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int, param4 *int, param5 *int) {
}

func Fn_gtk_cell_renderer_get_state(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_cell_renderer_get_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_renderer_is_activatable(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_renderer_render(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_gtk_cell_renderer_set_alignment(paramInstance unsafe.Pointer, param0 float32, param1 float32) {
}

func Fn_gtk_cell_renderer_set_fixed_size(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_cell_renderer_set_padding(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_cell_renderer_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_cell_renderer_set_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_cell_renderer_start_editing(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) {
}

func Fn_gtk_cell_renderer_stop_editing(paramInstance unsafe.Pointer, param0 bool) {}

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

func Fn_gtk_cell_renderer_spinner_new() {
	C.gtk_cell_renderer_spinner_new()
}

func Fn_gtk_cell_renderer_text_new() {
	C.gtk_cell_renderer_text_new()
}

func Fn_gtk_cell_renderer_text_set_fixed_height_from_font(paramInstance unsafe.Pointer, param0 int) {
}

func Fn_gtk_cell_renderer_toggle_new() {
	C.gtk_cell_renderer_toggle_new()
}

func Fn_gtk_cell_renderer_toggle_get_activatable(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_renderer_toggle_get_active(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_renderer_toggle_get_radio(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_renderer_toggle_set_activatable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_cell_renderer_toggle_set_active(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_cell_renderer_toggle_set_radio(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_cell_view_new() {
	C.gtk_cell_view_new()
}

func Fn_gtk_cell_view_new_with_context(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_cell_view_new_with_markup(param0 string) {}

func Fn_gtk_cell_view_new_with_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_cell_view_new_with_text(param0 string) {}

func Fn_gtk_cell_view_get_displayed_row(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_view_get_draw_sensitive(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_view_get_fit_model(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_view_get_model(paramInstance unsafe.Pointer) {}

func Fn_gtk_cell_view_get_size_of_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_cell_view_set_background_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_view_set_background_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_view_set_displayed_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_cell_view_set_draw_sensitive(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_cell_view_set_fit_model(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_cell_view_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_check_button_new() {
	C.gtk_check_button_new()
}

func Fn_gtk_check_button_new_with_label(param0 string) {}

func Fn_gtk_check_button_new_with_mnemonic(param0 string) {}

func Fn_gtk_check_menu_item_new() {
	C.gtk_check_menu_item_new()
}

func Fn_gtk_check_menu_item_new_with_label(param0 string) {}

func Fn_gtk_check_menu_item_new_with_mnemonic(param0 string) {}

func Fn_gtk_check_menu_item_get_active(paramInstance unsafe.Pointer) {}

func Fn_gtk_check_menu_item_get_draw_as_radio(paramInstance unsafe.Pointer) {}

func Fn_gtk_check_menu_item_get_inconsistent(paramInstance unsafe.Pointer) {}

func Fn_gtk_check_menu_item_set_active(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_check_menu_item_set_draw_as_radio(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_check_menu_item_set_inconsistent(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_check_menu_item_toggled(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_clear(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_get_display(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_get_owner(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : request_contents : has callback

// UNSUPPORTED : request_image : has callback

// UNSUPPORTED : request_rich_text : has callback

// UNSUPPORTED : request_targets : has callback

// UNSUPPORTED : request_text : has callback

// UNSUPPORTED : request_uris : has callback

func Fn_gtk_clipboard_set_can_store(paramInstance unsafe.Pointer, param0 []TargetEntry, param1 int) {
}

func Fn_gtk_clipboard_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_clipboard_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) {}

// UNSUPPORTED : set_with_data : has callback

// UNSUPPORTED : set_with_owner : has callback

func Fn_gtk_clipboard_store(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_wait_for_contents(paramInstance unsafe.Pointer, param0 gdk.Atom) {}

func Fn_gtk_clipboard_wait_for_image(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_wait_for_rich_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *uint64) {
}

func Fn_gtk_clipboard_wait_for_targets(paramInstance unsafe.Pointer, param0 []unsafe.Pointer, param1 *int) {
}

func Fn_gtk_clipboard_wait_for_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_wait_for_uris(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_wait_is_image_available(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_wait_is_rich_text_available(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_clipboard_wait_is_target_available(paramInstance unsafe.Pointer, param0 gdk.Atom) {}

func Fn_gtk_clipboard_wait_is_text_available(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_wait_is_uris_available(paramInstance unsafe.Pointer) {}

func Fn_gtk_clipboard_get(param0 gdk.Atom) {}

func Fn_gtk_clipboard_get_for_display(param0 unsafe.Pointer, param1 gdk.Atom) {}

func Fn_gtk_color_button_new() {
	C.gtk_color_button_new()
}

func Fn_gtk_color_button_new_with_color(param0 unsafe.Pointer) {}

func Fn_gtk_color_button_new_with_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_color_button_get_alpha(paramInstance unsafe.Pointer) {}

func Fn_gtk_color_button_get_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_button_get_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_button_get_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_color_button_get_use_alpha(paramInstance unsafe.Pointer) {}

func Fn_gtk_color_button_set_alpha(paramInstance unsafe.Pointer, param0 uint16) {}

func Fn_gtk_color_button_set_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_button_set_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_button_set_title(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_color_button_set_use_alpha(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_color_chooser_dialog_new(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_color_chooser_widget_new() {
	C.gtk_color_chooser_widget_new()
}

func Fn_gtk_color_selection_new() {
	C.gtk_color_selection_new()
}

func Fn_gtk_color_selection_get_current_alpha(paramInstance unsafe.Pointer) {}

func Fn_gtk_color_selection_get_current_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_get_current_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_get_has_opacity_control(paramInstance unsafe.Pointer) {}

func Fn_gtk_color_selection_get_has_palette(paramInstance unsafe.Pointer) {}

func Fn_gtk_color_selection_get_previous_alpha(paramInstance unsafe.Pointer) {}

func Fn_gtk_color_selection_get_previous_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_color_selection_get_previous_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_is_adjusting(paramInstance unsafe.Pointer) {}

func Fn_gtk_color_selection_set_current_alpha(paramInstance unsafe.Pointer, param0 uint16) {}

func Fn_gtk_color_selection_set_current_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_set_current_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_set_has_opacity_control(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_color_selection_set_has_palette(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_color_selection_set_previous_alpha(paramInstance unsafe.Pointer, param0 uint16) {}

func Fn_gtk_color_selection_set_previous_color(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_color_selection_set_previous_rgba(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_palette_from_string(param0 string, param1 []unsafe.Pointer, param2 *int) {
}

func Fn_gtk_color_selection_palette_to_string(param0 []gdk.Color, param1 int) {}

// UNSUPPORTED : set_change_palette_with_screen_hook : has callback

func Fn_gtk_color_selection_dialog_new(param0 string) {}

func Fn_gtk_color_selection_dialog_get_color_selection(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_new() {
	C.gtk_combo_box_new()
}

func Fn_gtk_combo_box_new_with_area(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_new_with_area_and_entry(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_new_with_entry() {
	C.gtk_combo_box_new_with_entry()
}

func Fn_gtk_combo_box_new_with_model(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_new_with_model_and_entry(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_get_active(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_active_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_active_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_get_add_tearoffs(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_button_sensitivity(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_column_span_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_entry_text_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_focus_on_click(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_has_entry(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_id_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_model(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_popup_accessible(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_popup_fixed_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_row_separator_func(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_row_span_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_get_wrap_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_popdown(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_popup(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_popup_for_device(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_set_active(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_combo_box_set_active_id(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_combo_box_set_active_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_set_add_tearoffs(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_combo_box_set_button_sensitivity(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_combo_box_set_column_span_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_combo_box_set_entry_text_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_combo_box_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_combo_box_set_id_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_combo_box_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_set_popup_fixed_width(paramInstance unsafe.Pointer, param0 bool) {}

// UNSUPPORTED : set_row_separator_func : has callback

func Fn_gtk_combo_box_set_row_span_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_combo_box_set_title(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_combo_box_set_wrap_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_combo_box_text_new() {
	C.gtk_combo_box_text_new()
}

func Fn_gtk_combo_box_text_new_with_entry() {
	C.gtk_combo_box_text_new_with_entry()
}

func Fn_gtk_combo_box_text_append(paramInstance unsafe.Pointer, param0 string, param1 string) {}

func Fn_gtk_combo_box_text_append_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_combo_box_text_get_active_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_combo_box_text_insert(paramInstance unsafe.Pointer, param0 int, param1 string, param2 string) {
}

func Fn_gtk_combo_box_text_insert_text(paramInstance unsafe.Pointer, param0 int, param1 string) {}

func Fn_gtk_combo_box_text_prepend(paramInstance unsafe.Pointer, param0 string, param1 string) {}

func Fn_gtk_combo_box_text_prepend_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_combo_box_text_remove(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_combo_box_text_remove_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : add_with_properties : has varargs

func Fn_gtk_container_check_resize(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : child_get : has varargs

func Fn_gtk_container_child_get_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
}

// UNSUPPORTED : child_get_valist : has va_list

func Fn_gtk_container_child_notify(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

// UNSUPPORTED : child_set : has varargs

func Fn_gtk_container_child_set_property(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
}

// UNSUPPORTED : child_set_valist : has va_list

func Fn_gtk_container_child_type(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : forall : has callback

// UNSUPPORTED : foreach : has callback

func Fn_gtk_container_get_border_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_get_children(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_get_focus_chain(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {}

func Fn_gtk_container_get_focus_child(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_get_focus_hadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_get_focus_vadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_get_path_for_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_container_get_resize_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_propagate_draw(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_container_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_container_resize_children(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_set_border_width(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_container_set_focus_chain(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_container_set_focus_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_container_set_focus_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_container_set_focus_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_container_set_reallocate_redraws(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_container_set_resize_mode(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_container_unset_focus_chain(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_cell_accessible_new() {
	C.gtk_container_cell_accessible_new()
}

func Fn_gtk_container_cell_accessible_add_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_container_cell_accessible_get_children(paramInstance unsafe.Pointer) {}

func Fn_gtk_container_cell_accessible_remove_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_css_provider_new() {
	C.gtk_css_provider_new()
}

func Fn_gtk_css_provider_load_from_data(paramInstance unsafe.Pointer, param0 []uint8, param1 uint64) {
}

func Fn_gtk_css_provider_load_from_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_css_provider_load_from_path(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_css_provider_to_string(paramInstance unsafe.Pointer) {}

func Fn_gtk_css_provider_get_default() {
	C.gtk_css_provider_get_default()
}

func Fn_gtk_css_provider_get_named(param0 string, param1 string) {}

func Fn_gtk_dialog_new() {
	C.gtk_dialog_new()
}

// UNSUPPORTED : new_with_buttons : has varargs

func Fn_gtk_dialog_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_dialog_add_button(paramInstance unsafe.Pointer, param0 string, param1 int) {}

// UNSUPPORTED : add_buttons : has varargs

func Fn_gtk_dialog_get_action_area(paramInstance unsafe.Pointer) {}

func Fn_gtk_dialog_get_content_area(paramInstance unsafe.Pointer) {}

func Fn_gtk_dialog_get_header_bar(paramInstance unsafe.Pointer) {}

func Fn_gtk_dialog_get_response_for_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_dialog_get_widget_for_response(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_dialog_response(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_dialog_run(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : set_alternative_button_order : has varargs

func Fn_gtk_dialog_set_alternative_button_order_from_array(paramInstance unsafe.Pointer, param0 int, param1 []int) {
}

func Fn_gtk_dialog_set_default_response(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_dialog_set_response_sensitive(paramInstance unsafe.Pointer, param0 int, param1 bool) {}

func Fn_gtk_drawing_area_new() {
	C.gtk_drawing_area_new()
}

func Fn_gtk_entry_new() {
	C.gtk_entry_new()
}

func Fn_gtk_entry_new_with_buffer(param0 unsafe.Pointer) {}

func Fn_gtk_entry_get_activates_default(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_alignment(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_attributes(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_buffer(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_completion(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_current_icon_drag_source(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_cursor_hadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_has_frame(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_icon_activatable(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_get_icon_area(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_entry_get_icon_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_entry_get_icon_gicon(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_get_icon_name(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_get_icon_pixbuf(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_get_icon_sensitive(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_get_icon_stock(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_get_icon_storage_type(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_get_icon_tooltip_markup(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_get_icon_tooltip_text(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_get_inner_border(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_input_hints(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_input_purpose(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_invisible_char(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_layout(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_entry_get_max_length(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_max_width_chars(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_overwrite_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_placeholder_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_progress_fraction(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_progress_pulse_step(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_tabs(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_text_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_entry_get_text_length(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_visibility(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_get_width_chars(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_im_context_filter_keypress(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_entry_layout_index_to_text_index(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_progress_pulse(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_reset_im_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_set_activates_default(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_entry_set_alignment(paramInstance unsafe.Pointer, param0 float32) {}

func Fn_gtk_entry_set_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_buffer(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_completion(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_cursor_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_has_frame(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_entry_set_icon_activatable(paramInstance unsafe.Pointer, param0 int, param1 bool) {}

func Fn_gtk_entry_set_icon_drag_source(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer, param2 int) {
}

func Fn_gtk_entry_set_icon_from_gicon(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_entry_set_icon_from_icon_name(paramInstance unsafe.Pointer, param0 int, param1 string) {}

func Fn_gtk_entry_set_icon_from_pixbuf(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_entry_set_icon_from_stock(paramInstance unsafe.Pointer, param0 int, param1 string) {}

func Fn_gtk_entry_set_icon_sensitive(paramInstance unsafe.Pointer, param0 int, param1 bool) {}

func Fn_gtk_entry_set_icon_tooltip_markup(paramInstance unsafe.Pointer, param0 int, param1 string) {}

func Fn_gtk_entry_set_icon_tooltip_text(paramInstance unsafe.Pointer, param0 int, param1 string) {}

func Fn_gtk_entry_set_inner_border(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_input_hints(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_set_input_purpose(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_set_invisible_char(paramInstance unsafe.Pointer, param0 rune) {}

func Fn_gtk_entry_set_max_length(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_set_max_width_chars(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_set_overwrite_mode(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_entry_set_placeholder_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_entry_set_progress_fraction(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_entry_set_progress_pulse_step(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_entry_set_tabs(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_entry_set_visibility(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_entry_set_width_chars(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_text_index_to_layout_index(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_unset_invisible_char(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_buffer_new(param0 string, param1 int) {}

func Fn_gtk_entry_buffer_delete_text(paramInstance unsafe.Pointer, param0 uint, param1 int) {}

func Fn_gtk_entry_buffer_emit_deleted_text(paramInstance unsafe.Pointer, param0 uint, param1 uint) {}

func Fn_gtk_entry_buffer_emit_inserted_text(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 uint) {
}

func Fn_gtk_entry_buffer_get_bytes(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_buffer_get_length(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_buffer_get_max_length(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_buffer_get_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_buffer_insert_text(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 int) {
}

func Fn_gtk_entry_buffer_set_max_length(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_buffer_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_entry_completion_new() {
	C.gtk_entry_completion_new()
}

func Fn_gtk_entry_completion_new_with_area(param0 unsafe.Pointer) {}

func Fn_gtk_entry_completion_complete(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_compute_prefix(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_entry_completion_delete_action(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_completion_get_completion_prefix(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_get_entry(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_get_inline_completion(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_get_inline_selection(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_get_minimum_key_length(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_get_model(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_get_popup_completion(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_get_popup_set_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_get_popup_single_match(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_get_text_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_insert_action_markup(paramInstance unsafe.Pointer, param0 int, param1 string) {
}

func Fn_gtk_entry_completion_insert_action_text(paramInstance unsafe.Pointer, param0 int, param1 string) {
}

func Fn_gtk_entry_completion_insert_prefix(paramInstance unsafe.Pointer) {}

func Fn_gtk_entry_completion_set_inline_completion(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_entry_completion_set_inline_selection(paramInstance unsafe.Pointer, param0 bool) {}

// UNSUPPORTED : set_match_func : has callback

func Fn_gtk_entry_completion_set_minimum_key_length(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_entry_completion_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_entry_completion_set_popup_completion(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_entry_completion_set_popup_set_width(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_entry_completion_set_popup_single_match(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_entry_completion_set_text_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_event_box_new() {
	C.gtk_event_box_new()
}

func Fn_gtk_event_box_get_above_child(paramInstance unsafe.Pointer) {}

func Fn_gtk_event_box_get_visible_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_event_box_set_above_child(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_event_box_set_visible_window(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_event_controller_key_new(param0 unsafe.Pointer) {}

func Fn_gtk_event_controller_key_forward(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_event_controller_key_get_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_event_controller_key_set_im_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_expander_new(param0 string) {}

func Fn_gtk_expander_new_with_mnemonic(param0 string) {}

func Fn_gtk_expander_get_expanded(paramInstance unsafe.Pointer) {}

func Fn_gtk_expander_get_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_expander_get_label_fill(paramInstance unsafe.Pointer) {}

func Fn_gtk_expander_get_label_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_expander_get_resize_toplevel(paramInstance unsafe.Pointer) {}

func Fn_gtk_expander_get_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_expander_get_use_markup(paramInstance unsafe.Pointer) {}

func Fn_gtk_expander_get_use_underline(paramInstance unsafe.Pointer) {}

func Fn_gtk_expander_set_expanded(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_expander_set_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_expander_set_label_fill(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_expander_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_expander_set_resize_toplevel(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_expander_set_spacing(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_expander_set_use_markup(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_expander_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_file_chooser_button_new(param0 string, param1 int) {}

func Fn_gtk_file_chooser_button_new_with_dialog(param0 unsafe.Pointer) {}

func Fn_gtk_file_chooser_button_get_focus_on_click(paramInstance unsafe.Pointer) {}

func Fn_gtk_file_chooser_button_get_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_file_chooser_button_get_width_chars(paramInstance unsafe.Pointer) {}

func Fn_gtk_file_chooser_button_set_focus_on_click(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_file_chooser_button_set_title(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_file_chooser_button_set_width_chars(paramInstance unsafe.Pointer, param0 int) {}

// UNSUPPORTED : new : has varargs

func Fn_gtk_file_chooser_widget_new(param0 int) {}

func Fn_gtk_file_filter_new() {
	C.gtk_file_filter_new()
}

// UNSUPPORTED : add_custom : has callback

func Fn_gtk_file_filter_add_mime_type(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_file_filter_add_pattern(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_file_filter_add_pixbuf_formats(paramInstance unsafe.Pointer) {}

func Fn_gtk_file_filter_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_file_filter_get_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_file_filter_get_needed(paramInstance unsafe.Pointer) {}

func Fn_gtk_file_filter_set_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_fixed_new() {
	C.gtk_fixed_new()
}

func Fn_gtk_fixed_move(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
}

func Fn_gtk_fixed_put(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_flow_box_new() {
	C.gtk_flow_box_new()
}

// UNSUPPORTED : bind_model : has callback

func Fn_gtk_flow_box_get_activate_on_single_click(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_get_child_at_index(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_flow_box_get_column_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_get_homogeneous(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_get_max_children_per_line(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_get_min_children_per_line(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_get_row_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_get_selected_children(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_get_selection_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_flow_box_invalidate_filter(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_invalidate_sort(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_select_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_select_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_flow_box_set_activate_on_single_click(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_flow_box_set_column_spacing(paramInstance unsafe.Pointer, param0 uint) {}

// UNSUPPORTED : set_filter_func : has callback

func Fn_gtk_flow_box_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_flow_box_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_flow_box_set_max_children_per_line(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_flow_box_set_min_children_per_line(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_flow_box_set_row_spacing(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_flow_box_set_selection_mode(paramInstance unsafe.Pointer, param0 int) {}

// UNSUPPORTED : set_sort_func : has callback

func Fn_gtk_flow_box_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_flow_box_unselect_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_unselect_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_flow_box_child_new() {
	C.gtk_flow_box_child_new()
}

func Fn_gtk_flow_box_child_changed(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_child_get_index(paramInstance unsafe.Pointer) {}

func Fn_gtk_flow_box_child_is_selected(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_button_new() {
	C.gtk_font_button_new()
}

func Fn_gtk_font_button_new_with_font(param0 string) {}

func Fn_gtk_font_button_get_font_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_button_get_show_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_button_get_show_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_button_get_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_button_get_use_font(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_button_get_use_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_button_set_font_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_font_button_set_show_size(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_font_button_set_show_style(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_font_button_set_title(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_font_button_set_use_font(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_font_button_set_use_size(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_font_chooser_dialog_new(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_font_chooser_widget_new() {
	C.gtk_font_chooser_widget_new()
}

func Fn_gtk_font_selection_new() {
	C.gtk_font_selection_new()
}

func Fn_gtk_font_selection_get_face(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_get_face_list(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_get_family(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_get_family_list(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_get_font_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_get_preview_entry(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_get_preview_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_get_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_get_size_entry(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_get_size_list(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_set_font_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_font_selection_set_preview_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_font_selection_dialog_new(param0 string) {}

func Fn_gtk_font_selection_dialog_get_cancel_button(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_dialog_get_font_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_dialog_get_font_selection(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_dialog_get_ok_button(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_dialog_get_preview_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_font_selection_dialog_set_font_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_font_selection_dialog_set_preview_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_frame_new(param0 string) {}

func Fn_gtk_frame_get_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_frame_get_label_align(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {}

func Fn_gtk_frame_get_label_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_frame_get_shadow_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_frame_set_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_frame_set_label_align(paramInstance unsafe.Pointer, param0 float32, param1 float32) {}

func Fn_gtk_frame_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_frame_set_shadow_type(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_gesture_get_last_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_grid_new() {
	C.gtk_grid_new()
}

func Fn_gtk_grid_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int) {
}

func Fn_gtk_grid_attach_next_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int) {
}

func Fn_gtk_grid_get_baseline_row(paramInstance unsafe.Pointer) {}

func Fn_gtk_grid_get_child_at(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_grid_get_column_homogeneous(paramInstance unsafe.Pointer) {}

func Fn_gtk_grid_get_column_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_grid_get_row_baseline_position(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_grid_get_row_homogeneous(paramInstance unsafe.Pointer) {}

func Fn_gtk_grid_get_row_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_grid_insert_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_grid_insert_next_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_grid_insert_row(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_grid_remove_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_grid_remove_row(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_grid_set_baseline_row(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_grid_set_column_homogeneous(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_grid_set_column_spacing(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_grid_set_row_baseline_position(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_grid_set_row_homogeneous(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_grid_set_row_spacing(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_hbox_new(param0 bool, param1 int) {}

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
}

func Fn_gtk_hsv_get_metrics(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_hsv_is_adjusting(paramInstance unsafe.Pointer) {}

func Fn_gtk_hsv_set_color(paramInstance unsafe.Pointer, param0 float64, param1 float64, param2 float64) {
}

func Fn_gtk_hsv_set_metrics(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_hsv_to_rgb(param0 float64, param1 float64, param2 float64, param3 *float64, param4 *float64, param5 *float64) {
}

func Fn_gtk_hscale_new(param0 unsafe.Pointer) {}

func Fn_gtk_hscale_new_with_range(param0 float64, param1 float64, param2 float64) {}

func Fn_gtk_hscrollbar_new(param0 unsafe.Pointer) {}

func Fn_gtk_hseparator_new() {
	C.gtk_hseparator_new()
}

func Fn_gtk_handle_box_new() {
	C.gtk_handle_box_new()
}

func Fn_gtk_handle_box_get_child_detached(paramInstance unsafe.Pointer) {}

func Fn_gtk_handle_box_get_handle_position(paramInstance unsafe.Pointer) {}

func Fn_gtk_handle_box_get_shadow_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_handle_box_get_snap_edge(paramInstance unsafe.Pointer) {}

func Fn_gtk_handle_box_set_handle_position(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_handle_box_set_shadow_type(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_handle_box_set_snap_edge(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_header_bar_new() {
	C.gtk_header_bar_new()
}

func Fn_gtk_header_bar_get_custom_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_header_bar_get_decoration_layout(paramInstance unsafe.Pointer) {}

func Fn_gtk_header_bar_get_has_subtitle(paramInstance unsafe.Pointer) {}

func Fn_gtk_header_bar_get_show_close_button(paramInstance unsafe.Pointer) {}

func Fn_gtk_header_bar_get_subtitle(paramInstance unsafe.Pointer) {}

func Fn_gtk_header_bar_get_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_header_bar_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_header_bar_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_header_bar_set_custom_title(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_header_bar_set_decoration_layout(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_header_bar_set_has_subtitle(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_header_bar_set_show_close_button(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_header_bar_set_subtitle(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_header_bar_set_title(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_im_context_delete_surrounding(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_im_context_filter_keypress(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_im_context_focus_in(paramInstance unsafe.Pointer) {}

func Fn_gtk_im_context_focus_out(paramInstance unsafe.Pointer) {}

func Fn_gtk_im_context_get_preedit_string(paramInstance unsafe.Pointer, param0 string, param1 *unsafe.Pointer, param2 *int) {
}

func Fn_gtk_im_context_get_surrounding(paramInstance unsafe.Pointer, param0 string, param1 *int) {}

func Fn_gtk_im_context_reset(paramInstance unsafe.Pointer) {}

func Fn_gtk_im_context_set_client_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_im_context_set_cursor_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_im_context_set_surrounding(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int) {
}

func Fn_gtk_im_context_set_use_preedit(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_im_context_simple_new() {
	C.gtk_im_context_simple_new()
}

func Fn_gtk_im_context_simple_add_compose_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_im_context_simple_add_table(paramInstance unsafe.Pointer, param0 []uint16, param1 int, param2 int) {
}

func Fn_gtk_im_multicontext_new() {
	C.gtk_im_multicontext_new()
}

func Fn_gtk_im_multicontext_append_menuitems(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_im_multicontext_get_context_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_im_multicontext_set_context_id(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_icon_factory_new() {
	C.gtk_icon_factory_new()
}

func Fn_gtk_icon_factory_add(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_icon_factory_add_default(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_factory_lookup(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_icon_factory_remove_default(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_factory_lookup_default(param0 string) {}

func Fn_gtk_icon_info_new_for_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_icon_info_copy(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_info_free(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_info_get_attach_points(paramInstance unsafe.Pointer, param0 []unsafe.Pointer, param1 *int) {
}

func Fn_gtk_icon_info_get_base_scale(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_info_get_base_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_info_get_builtin_pixbuf(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_info_get_display_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_info_get_embedded_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_info_get_filename(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_info_is_symbolic(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_info_load_icon(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : load_icon_async : has callback

func Fn_gtk_icon_info_load_icon_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_info_load_surface(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_info_load_symbolic(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 *bool) {
}

// UNSUPPORTED : load_symbolic_async : has callback

func Fn_gtk_icon_info_load_symbolic_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *bool) {
}

func Fn_gtk_icon_info_load_symbolic_for_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *bool) {
}

// UNSUPPORTED : load_symbolic_for_context_async : has callback

func Fn_gtk_icon_info_load_symbolic_for_context_finish(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *bool) {
}

func Fn_gtk_icon_info_load_symbolic_for_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *bool) {
}

func Fn_gtk_icon_info_set_raw_coordinates(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_icon_theme_new() {
	C.gtk_icon_theme_new()
}

func Fn_gtk_icon_theme_append_search_path(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_icon_theme_choose_icon(paramInstance unsafe.Pointer, param0 []string, param1 int, param2 int) {
}

func Fn_gtk_icon_theme_choose_icon_for_scale(paramInstance unsafe.Pointer, param0 []string, param1 int, param2 int, param3 int) {
}

func Fn_gtk_icon_theme_get_example_icon_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_theme_get_icon_sizes(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_icon_theme_get_search_path(paramInstance unsafe.Pointer, param0 *[]string, param1 *int) {
}

func Fn_gtk_icon_theme_has_icon(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_icon_theme_list_contexts(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_theme_list_icons(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_icon_theme_load_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int) {
}

func Fn_gtk_icon_theme_load_icon_for_scale(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, param3 int) {
}

func Fn_gtk_icon_theme_load_surface(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, param3 unsafe.Pointer, param4 int) {
}

func Fn_gtk_icon_theme_lookup_by_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
}

func Fn_gtk_icon_theme_lookup_by_gicon_for_scale(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int) {
}

func Fn_gtk_icon_theme_lookup_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int) {
}

func Fn_gtk_icon_theme_lookup_icon_for_scale(paramInstance unsafe.Pointer, param0 string, param1 int, param2 int, param3 int) {
}

func Fn_gtk_icon_theme_prepend_search_path(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_icon_theme_rescan_if_needed(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_theme_set_custom_theme(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_icon_theme_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_theme_set_search_path(paramInstance unsafe.Pointer, param0 []string, param1 int) {}

func Fn_gtk_icon_theme_add_builtin_icon(param0 string, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_icon_theme_get_default() {
	C.gtk_icon_theme_get_default()
}

func Fn_gtk_icon_theme_get_for_screen(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_new() {
	C.gtk_icon_view_new()
}

func Fn_gtk_icon_view_new_with_area(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_new_with_model(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_convert_widget_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_icon_view_create_drag_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_enable_model_drag_dest(paramInstance unsafe.Pointer, param0 []TargetEntry, param1 int, param2 int) {
}

func Fn_gtk_icon_view_enable_model_drag_source(paramInstance unsafe.Pointer, param0 int, param1 []TargetEntry, param2 int, param3 int) {
}

func Fn_gtk_icon_view_get_activate_on_single_click(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_cell_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_icon_view_get_column_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_columns(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_cursor(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) {
}

func Fn_gtk_icon_view_get_dest_item_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 int) {
}

func Fn_gtk_icon_view_get_drag_dest_item(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 int) {
}

func Fn_gtk_icon_view_get_item_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer) {
}

func Fn_gtk_icon_view_get_item_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_get_item_orientation(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_item_padding(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_item_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_get_item_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_margin(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_markup_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_model(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_path_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_icon_view_get_pixbuf_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_reorderable(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_row_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_selected_items(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_selection_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_text_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_tooltip_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_get_tooltip_context(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_gtk_icon_view_get_visible_range(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) {
}

func Fn_gtk_icon_view_item_activated(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_path_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_scroll_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 float32, param3 float32) {
}

func Fn_gtk_icon_view_select_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_select_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_icon_view_set_activate_on_single_click(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_icon_view_set_column_spacing(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_columns(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
}

func Fn_gtk_icon_view_set_drag_dest_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_icon_view_set_item_orientation(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_item_padding(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_item_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_margin(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_markup_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_set_pixbuf_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_reorderable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_icon_view_set_row_spacing(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_selection_mode(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_spacing(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_text_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_tooltip_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_icon_view_set_tooltip_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_icon_view_set_tooltip_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_icon_view_unselect_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_unselect_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_unset_model_drag_dest(paramInstance unsafe.Pointer) {}

func Fn_gtk_icon_view_unset_model_drag_source(paramInstance unsafe.Pointer) {}

func Fn_gtk_image_new() {
	C.gtk_image_new()
}

func Fn_gtk_image_new_from_animation(param0 unsafe.Pointer) {}

func Fn_gtk_image_new_from_file(param0 string) {}

func Fn_gtk_image_new_from_gicon(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_image_new_from_icon_name(param0 string, param1 int) {}

func Fn_gtk_image_new_from_icon_set(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_image_new_from_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_image_new_from_resource(param0 string) {}

func Fn_gtk_image_new_from_stock(param0 string, param1 int) {}

func Fn_gtk_image_new_from_surface(param0 unsafe.Pointer) {}

func Fn_gtk_image_clear(paramInstance unsafe.Pointer) {}

func Fn_gtk_image_get_animation(paramInstance unsafe.Pointer) {}

func Fn_gtk_image_get_gicon(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 int) {}

func Fn_gtk_image_get_icon_name(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_image_get_icon_set(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 int) {}

func Fn_gtk_image_get_pixbuf(paramInstance unsafe.Pointer) {}

func Fn_gtk_image_get_pixel_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_image_get_stock(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_image_get_storage_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_image_set_from_animation(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_image_set_from_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_image_set_from_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_image_set_from_icon_name(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_image_set_from_icon_set(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_image_set_from_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_image_set_from_resource(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_image_set_from_stock(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_image_set_from_surface(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_image_set_pixel_size(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_image_menu_item_new() {
	C.gtk_image_menu_item_new()
}

func Fn_gtk_image_menu_item_new_from_stock(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_image_menu_item_new_with_label(param0 string) {}

func Fn_gtk_image_menu_item_new_with_mnemonic(param0 string) {}

func Fn_gtk_image_menu_item_get_always_show_image(paramInstance unsafe.Pointer) {}

func Fn_gtk_image_menu_item_get_image(paramInstance unsafe.Pointer) {}

func Fn_gtk_image_menu_item_get_use_stock(paramInstance unsafe.Pointer) {}

func Fn_gtk_image_menu_item_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_image_menu_item_set_always_show_image(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_image_menu_item_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_image_menu_item_set_use_stock(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_info_bar_new() {
	C.gtk_info_bar_new()
}

// UNSUPPORTED : new_with_buttons : has varargs

func Fn_gtk_info_bar_add_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_info_bar_add_button(paramInstance unsafe.Pointer, param0 string, param1 int) {}

// UNSUPPORTED : add_buttons : has varargs

func Fn_gtk_info_bar_get_action_area(paramInstance unsafe.Pointer) {}

func Fn_gtk_info_bar_get_content_area(paramInstance unsafe.Pointer) {}

func Fn_gtk_info_bar_get_message_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_info_bar_get_show_close_button(paramInstance unsafe.Pointer) {}

func Fn_gtk_info_bar_response(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_info_bar_set_default_response(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_info_bar_set_message_type(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_info_bar_set_response_sensitive(paramInstance unsafe.Pointer, param0 int, param1 bool) {}

func Fn_gtk_info_bar_set_show_close_button(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_invisible_new() {
	C.gtk_invisible_new()
}

func Fn_gtk_invisible_new_for_screen(param0 unsafe.Pointer) {}

func Fn_gtk_invisible_get_screen(paramInstance unsafe.Pointer) {}

func Fn_gtk_invisible_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_label_new(param0 string) {}

func Fn_gtk_label_new_with_mnemonic(param0 string) {}

func Fn_gtk_label_get_angle(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_attributes(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_current_uri(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_ellipsize(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_justify(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_layout(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_label_get_line_wrap(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_line_wrap_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_lines(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_max_width_chars(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_mnemonic_keyval(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_mnemonic_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_selectable(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_selection_bounds(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_label_get_single_line_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_track_visited_links(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_use_markup(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_use_underline(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_get_width_chars(paramInstance unsafe.Pointer) {}

func Fn_gtk_label_select_region(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_label_set_angle(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_label_set_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_label_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_label_set_justify(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_label_set_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_label_set_line_wrap(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_label_set_line_wrap_mode(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_label_set_lines(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_label_set_markup(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_label_set_markup_with_mnemonic(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_label_set_max_width_chars(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_label_set_mnemonic_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_label_set_pattern(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_label_set_selectable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_label_set_single_line_mode(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_label_set_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_label_set_text_with_mnemonic(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_label_set_track_visited_links(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_label_set_use_markup(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_label_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_label_set_width_chars(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_layout_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_layout_get_bin_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_layout_get_hadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_layout_get_size(paramInstance unsafe.Pointer, param0 *uint, param1 *uint) {}

func Fn_gtk_layout_get_vadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_layout_move(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
}

func Fn_gtk_layout_put(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
}

func Fn_gtk_layout_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_layout_set_size(paramInstance unsafe.Pointer, param0 uint, param1 uint) {}

func Fn_gtk_layout_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_level_bar_new() {
	C.gtk_level_bar_new()
}

func Fn_gtk_level_bar_new_for_interval(param0 float64, param1 float64) {}

func Fn_gtk_level_bar_add_offset_value(paramInstance unsafe.Pointer, param0 string, param1 float64) {}

func Fn_gtk_level_bar_get_inverted(paramInstance unsafe.Pointer) {}

func Fn_gtk_level_bar_get_max_value(paramInstance unsafe.Pointer) {}

func Fn_gtk_level_bar_get_min_value(paramInstance unsafe.Pointer) {}

func Fn_gtk_level_bar_get_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_level_bar_get_offset_value(paramInstance unsafe.Pointer, param0 string, param1 *float64) {
}

func Fn_gtk_level_bar_get_value(paramInstance unsafe.Pointer) {}

func Fn_gtk_level_bar_remove_offset_value(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_level_bar_set_inverted(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_level_bar_set_max_value(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_level_bar_set_min_value(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_level_bar_set_mode(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_level_bar_set_value(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_link_button_new(param0 string) {}

func Fn_gtk_link_button_new_with_label(param0 string, param1 string) {}

func Fn_gtk_link_button_get_uri(paramInstance unsafe.Pointer) {}

func Fn_gtk_link_button_get_visited(paramInstance unsafe.Pointer) {}

func Fn_gtk_link_button_set_uri(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_link_button_set_visited(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_list_box_new() {
	C.gtk_list_box_new()
}

// UNSUPPORTED : bind_model : has callback

func Fn_gtk_list_box_drag_highlight_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_list_box_drag_unhighlight_row(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_get_activate_on_single_click(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_get_adjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_get_row_at_index(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_list_box_get_row_at_y(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_list_box_get_selected_row(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_get_selection_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_list_box_invalidate_filter(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_invalidate_headers(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_invalidate_sort(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_prepend(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_list_box_select_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_list_box_set_activate_on_single_click(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_list_box_set_adjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : set_filter_func : has callback

// UNSUPPORTED : set_header_func : has callback

func Fn_gtk_list_box_set_placeholder(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_list_box_set_selection_mode(paramInstance unsafe.Pointer, param0 int) {}

// UNSUPPORTED : set_sort_func : has callback

func Fn_gtk_list_box_row_new() {
	C.gtk_list_box_row_new()
}

func Fn_gtk_list_box_row_changed(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_row_get_header(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_row_get_index(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_box_row_set_header(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : new : has varargs

func Fn_gtk_list_store_newv(param0 int, param1 []uint64) {}

func Fn_gtk_list_store_append(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_list_store_clear(paramInstance unsafe.Pointer) {}

func Fn_gtk_list_store_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_list_store_insert_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_list_store_insert_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

// UNSUPPORTED : insert_with_values : has varargs

func Fn_gtk_list_store_insert_with_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 []int, param3 []gobject.Value, param4 int) {
}

func Fn_gtk_list_store_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_list_store_move_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_list_store_move_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_list_store_prepend(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_list_store_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_list_store_reorder(paramInstance unsafe.Pointer, param0 []int) {}

// UNSUPPORTED : set : has varargs

func Fn_gtk_list_store_set_column_types(paramInstance unsafe.Pointer, param0 int, param1 []uint64) {
}

// UNSUPPORTED : set_valist : has va_list

func Fn_gtk_list_store_set_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {
}

func Fn_gtk_list_store_set_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []int, param2 []gobject.Value, param3 int) {
}

func Fn_gtk_list_store_swap(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_lock_button_new(param0 unsafe.Pointer) {}

func Fn_gtk_lock_button_get_permission(paramInstance unsafe.Pointer) {}

func Fn_gtk_lock_button_set_permission(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_new() {
	C.gtk_menu_new()
}

func Fn_gtk_menu_new_from_model(param0 unsafe.Pointer) {}

func Fn_gtk_menu_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint) {
}

// UNSUPPORTED : attach_to_widget : has callback

func Fn_gtk_menu_detach(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_get_accel_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_get_accel_path(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_get_active(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_get_attach_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_get_monitor(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_get_reserve_toggle_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_get_tearoff_state(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_get_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_popdown(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : popup : has callback

// UNSUPPORTED : popup_for_device : has callback

func Fn_gtk_menu_reorder_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_menu_reposition(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_set_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_set_accel_path(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_menu_set_active(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_menu_set_monitor(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_menu_set_reserve_toggle_size(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_menu_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_set_tearoff_state(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_menu_set_title(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_menu_get_for_attach_widget(param0 unsafe.Pointer) {}

func Fn_gtk_menu_bar_new() {
	C.gtk_menu_bar_new()
}

func Fn_gtk_menu_bar_new_from_model(param0 unsafe.Pointer) {}

func Fn_gtk_menu_bar_get_child_pack_direction(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_bar_get_pack_direction(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_bar_set_child_pack_direction(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_menu_bar_set_pack_direction(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_menu_button_new() {
	C.gtk_menu_button_new()
}

func Fn_gtk_menu_button_get_align_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_button_get_direction(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_button_get_menu_model(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_button_get_popover(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_button_get_popup(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_button_get_use_popover(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_button_set_align_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_button_set_direction(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_menu_button_set_menu_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_button_set_popover(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_button_set_popup(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_button_set_use_popover(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_menu_item_new() {
	C.gtk_menu_item_new()
}

func Fn_gtk_menu_item_new_with_label(param0 string) {}

func Fn_gtk_menu_item_new_with_mnemonic(param0 string) {}

func Fn_gtk_menu_item_activate(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_item_deselect(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_item_get_accel_path(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_item_get_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_item_get_reserve_indicator(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_item_get_right_justified(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_item_get_submenu(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_item_get_use_underline(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_item_select(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_item_set_accel_path(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_menu_item_set_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_menu_item_set_reserve_indicator(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_menu_item_set_right_justified(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_menu_item_set_submenu(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_item_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_menu_item_toggle_size_allocate(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_menu_item_toggle_size_request(paramInstance unsafe.Pointer, param0 *int) {}

func Fn_gtk_menu_shell_activate_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_menu_shell_append(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_shell_bind_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 bool) {
}

func Fn_gtk_menu_shell_cancel(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_shell_deactivate(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_shell_deselect(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_shell_get_parent_shell(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_shell_get_selected_item(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_shell_get_take_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_shell_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_menu_shell_prepend(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_shell_select_first(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_menu_shell_select_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_menu_shell_set_take_focus(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_menu_tool_button_new(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_menu_tool_button_new_from_stock(param0 string) {}

func Fn_gtk_menu_tool_button_get_menu(paramInstance unsafe.Pointer) {}

func Fn_gtk_menu_tool_button_set_arrow_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_menu_tool_button_set_arrow_tooltip_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_menu_tool_button_set_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_with_markup : has varargs

// UNSUPPORTED : format_secondary_markup : has varargs

// UNSUPPORTED : format_secondary_text : has varargs

func Fn_gtk_message_dialog_get_image(paramInstance unsafe.Pointer) {}

func Fn_gtk_message_dialog_get_message_area(paramInstance unsafe.Pointer) {}

func Fn_gtk_message_dialog_set_image(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_message_dialog_set_markup(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_misc_get_alignment(paramInstance unsafe.Pointer, param0 *float32, param1 *float32) {}

func Fn_gtk_misc_get_padding(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_misc_set_alignment(paramInstance unsafe.Pointer, param0 float32, param1 float32) {}

func Fn_gtk_misc_set_padding(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_mount_operation_new(param0 unsafe.Pointer) {}

func Fn_gtk_mount_operation_get_parent(paramInstance unsafe.Pointer) {}

func Fn_gtk_mount_operation_get_screen(paramInstance unsafe.Pointer) {}

func Fn_gtk_mount_operation_is_showing(paramInstance unsafe.Pointer) {}

func Fn_gtk_mount_operation_set_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_mount_operation_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_notebook_new() {
	C.gtk_notebook_new()
}

func Fn_gtk_notebook_append_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_notebook_append_page_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_notebook_get_action_widget(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_notebook_get_current_page(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_get_group_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_get_menu_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_menu_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_n_pages(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_get_nth_page(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_notebook_get_scrollable(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_get_show_border(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_get_show_tabs(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_detachable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_hborder(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_pos(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_reorderable(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_vborder(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_insert_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
}

func Fn_gtk_notebook_insert_page_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int) {
}

func Fn_gtk_notebook_next_page(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_page_num(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_notebook_popup_disable(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_popup_enable(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_prepend_page(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_notebook_prepend_page_menu(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_notebook_prev_page(paramInstance unsafe.Pointer) {}

func Fn_gtk_notebook_remove_page(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_notebook_reorder_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_notebook_set_action_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_notebook_set_current_page(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_notebook_set_group_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_notebook_set_menu_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_notebook_set_menu_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

func Fn_gtk_notebook_set_scrollable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_notebook_set_show_border(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_notebook_set_show_tabs(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_notebook_set_tab_detachable(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_notebook_set_tab_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_notebook_set_tab_label_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

func Fn_gtk_notebook_set_tab_pos(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_notebook_set_tab_reorderable(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_notebook_page_accessible_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_notebook_page_accessible_invalidate(paramInstance unsafe.Pointer) {}

func Fn_gtk_numerable_icon_get_background_gicon(paramInstance unsafe.Pointer) {}

func Fn_gtk_numerable_icon_get_background_icon_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_numerable_icon_get_count(paramInstance unsafe.Pointer) {}

func Fn_gtk_numerable_icon_get_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_numerable_icon_get_style_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_numerable_icon_set_background_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_numerable_icon_set_background_icon_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_numerable_icon_set_count(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_numerable_icon_set_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_numerable_icon_set_style_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_numerable_icon_new(param0 unsafe.Pointer) {}

func Fn_gtk_numerable_icon_new_with_style_context(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_offscreen_window_new() {
	C.gtk_offscreen_window_new()
}

func Fn_gtk_offscreen_window_get_pixbuf(paramInstance unsafe.Pointer) {}

func Fn_gtk_offscreen_window_get_surface(paramInstance unsafe.Pointer) {}

func Fn_gtk_overlay_new() {
	C.gtk_overlay_new()
}

func Fn_gtk_overlay_add_overlay(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_page_setup_new() {
	C.gtk_page_setup_new()
}

func Fn_gtk_page_setup_new_from_file(param0 string) {}

func Fn_gtk_page_setup_new_from_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_page_setup_copy(paramInstance unsafe.Pointer) {}

func Fn_gtk_page_setup_get_bottom_margin(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_page_setup_get_left_margin(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_page_setup_get_orientation(paramInstance unsafe.Pointer) {}

func Fn_gtk_page_setup_get_page_height(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_page_setup_get_page_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_page_setup_get_paper_height(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_page_setup_get_paper_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_page_setup_get_paper_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_page_setup_get_right_margin(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_page_setup_get_top_margin(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_page_setup_load_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_page_setup_load_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

func Fn_gtk_page_setup_set_bottom_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {}

func Fn_gtk_page_setup_set_left_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {}

func Fn_gtk_page_setup_set_orientation(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_page_setup_set_paper_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_page_setup_set_paper_size_and_default_margins(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_page_setup_set_right_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {}

func Fn_gtk_page_setup_set_top_margin(paramInstance unsafe.Pointer, param0 float64, param1 int) {}

func Fn_gtk_page_setup_to_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_page_setup_to_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

func Fn_gtk_paned_new(param0 int) {}

func Fn_gtk_paned_add1(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_paned_add2(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_paned_get_child1(paramInstance unsafe.Pointer) {}

func Fn_gtk_paned_get_child2(paramInstance unsafe.Pointer) {}

func Fn_gtk_paned_get_handle_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_paned_get_position(paramInstance unsafe.Pointer) {}

func Fn_gtk_paned_pack1(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool) {
}

func Fn_gtk_paned_pack2(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool) {
}

func Fn_gtk_paned_set_position(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_places_sidebar_new() {
	C.gtk_places_sidebar_new()
}

func Fn_gtk_places_sidebar_add_shortcut(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_places_sidebar_get_local_only(paramInstance unsafe.Pointer) {}

func Fn_gtk_places_sidebar_get_location(paramInstance unsafe.Pointer) {}

func Fn_gtk_places_sidebar_get_nth_bookmark(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_places_sidebar_get_open_flags(paramInstance unsafe.Pointer) {}

func Fn_gtk_places_sidebar_get_show_connect_to_server(paramInstance unsafe.Pointer) {}

func Fn_gtk_places_sidebar_get_show_desktop(paramInstance unsafe.Pointer) {}

func Fn_gtk_places_sidebar_list_shortcuts(paramInstance unsafe.Pointer) {}

func Fn_gtk_places_sidebar_remove_shortcut(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_places_sidebar_set_local_only(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_places_sidebar_set_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_places_sidebar_set_open_flags(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_places_sidebar_set_show_connect_to_server(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_places_sidebar_set_show_desktop(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_plug_new(param0 uint64) {}

func Fn_gtk_plug_new_for_display(param0 unsafe.Pointer, param1 uint64) {}

func Fn_gtk_plug_construct(paramInstance unsafe.Pointer, param0 uint64) {}

func Fn_gtk_plug_construct_for_display(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint64) {
}

func Fn_gtk_plug_get_embedded(paramInstance unsafe.Pointer) {}

func Fn_gtk_plug_get_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_plug_get_socket_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_popover_new(param0 unsafe.Pointer) {}

func Fn_gtk_popover_new_from_model(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_popover_bind_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_popover_get_modal(paramInstance unsafe.Pointer) {}

func Fn_gtk_popover_get_pointing_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_popover_get_position(paramInstance unsafe.Pointer) {}

func Fn_gtk_popover_get_relative_to(paramInstance unsafe.Pointer) {}

func Fn_gtk_popover_set_modal(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_popover_set_pointing_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_popover_set_position(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_popover_set_relative_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_print_context_create_pango_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_context_create_pango_layout(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_context_get_cairo_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_context_get_dpi_x(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_context_get_dpi_y(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_context_get_hard_margins(paramInstance unsafe.Pointer, param0 *float64, param1 *float64, param2 *float64, param3 *float64) {
}

func Fn_gtk_print_context_get_height(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_context_get_page_setup(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_context_get_pango_fontmap(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_context_get_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_context_set_cairo_context(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 float64) {
}

func Fn_gtk_print_operation_new() {
	C.gtk_print_operation_new()
}

func Fn_gtk_print_operation_cancel(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_draw_page_finish(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_get_default_page_setup(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_get_embed_page_setup(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_get_error(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_get_has_selection(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_get_n_pages_to_print(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_get_print_settings(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_get_status(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_get_status_string(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_get_support_selection(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_is_finished(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_run(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_print_operation_set_allow_async(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_operation_set_current_page(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_operation_set_custom_tab_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_operation_set_default_page_setup(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_print_operation_set_defer_drawing(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_operation_set_embed_page_setup(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_operation_set_export_filename(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_operation_set_has_selection(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_operation_set_job_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_operation_set_n_pages(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_operation_set_print_settings(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_print_operation_set_show_progress(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_operation_set_support_selection(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_operation_set_track_print_status(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_operation_set_unit(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_operation_set_use_full_page(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_settings_new() {
	C.gtk_print_settings_new()
}

func Fn_gtk_print_settings_new_from_file(param0 string) {}

func Fn_gtk_print_settings_new_from_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_print_settings_copy(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : foreach : has callback

func Fn_gtk_print_settings_get(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_get_bool(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_get_collate(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_default_source(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_dither(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_double(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_get_double_with_default(paramInstance unsafe.Pointer, param0 string, param1 float64) {
}

func Fn_gtk_print_settings_get_duplex(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_finishings(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_int(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_get_int_with_default(paramInstance unsafe.Pointer, param0 string, param1 int) {
}

func Fn_gtk_print_settings_get_length(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_print_settings_get_media_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_n_copies(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_number_up(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_number_up_layout(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_orientation(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_output_bin(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_page_ranges(paramInstance unsafe.Pointer, param0 *int) {}

func Fn_gtk_print_settings_get_page_set(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_paper_height(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_get_paper_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_paper_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_get_print_pages(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_printer(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_printer_lpi(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_quality(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_resolution(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_resolution_x(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_resolution_y(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_reverse(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_scale(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_get_use_color(paramInstance unsafe.Pointer) {}

func Fn_gtk_print_settings_has_key(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_load_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_load_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

func Fn_gtk_print_settings_set(paramInstance unsafe.Pointer, param0 string, param1 string) {}

func Fn_gtk_print_settings_set_bool(paramInstance unsafe.Pointer, param0 string, param1 bool) {}

func Fn_gtk_print_settings_set_collate(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_settings_set_default_source(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_set_dither(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_set_double(paramInstance unsafe.Pointer, param0 string, param1 float64) {}

func Fn_gtk_print_settings_set_duplex(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_set_finishings(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_set_int(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_print_settings_set_length(paramInstance unsafe.Pointer, param0 string, param1 float64, param2 int) {
}

func Fn_gtk_print_settings_set_media_type(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_set_n_copies(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_set_number_up(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_set_number_up_layout(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_set_orientation(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_set_output_bin(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_set_page_ranges(paramInstance unsafe.Pointer, param0 []PageRange, param1 int) {
}

func Fn_gtk_print_settings_set_page_set(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_set_paper_height(paramInstance unsafe.Pointer, param0 float64, param1 int) {
}

func Fn_gtk_print_settings_set_paper_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_print_settings_set_paper_width(paramInstance unsafe.Pointer, param0 float64, param1 int) {
}

func Fn_gtk_print_settings_set_print_pages(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_set_printer(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_set_printer_lpi(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_print_settings_set_quality(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_set_resolution(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_print_settings_set_resolution_xy(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_print_settings_set_reverse(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_settings_set_scale(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_print_settings_set_use_color(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_print_settings_to_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_print_settings_to_key_file(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {
}

func Fn_gtk_print_settings_unset(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_progress_bar_new() {
	C.gtk_progress_bar_new()
}

func Fn_gtk_progress_bar_get_ellipsize(paramInstance unsafe.Pointer) {}

func Fn_gtk_progress_bar_get_fraction(paramInstance unsafe.Pointer) {}

func Fn_gtk_progress_bar_get_inverted(paramInstance unsafe.Pointer) {}

func Fn_gtk_progress_bar_get_pulse_step(paramInstance unsafe.Pointer) {}

func Fn_gtk_progress_bar_get_show_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_progress_bar_get_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_progress_bar_pulse(paramInstance unsafe.Pointer) {}

func Fn_gtk_progress_bar_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_progress_bar_set_fraction(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_progress_bar_set_inverted(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_progress_bar_set_pulse_step(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_progress_bar_set_show_text(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_progress_bar_set_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_radio_action_new(param0 string, param1 string, param2 string, param3 string, param4 int) {}

func Fn_gtk_radio_action_get_current_value(paramInstance unsafe.Pointer) {}

func Fn_gtk_radio_action_get_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_radio_action_join_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_radio_action_set_current_value(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_radio_action_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_new(param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_new_from_widget(param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_new_with_label(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_new_with_mnemonic(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_get_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_radio_button_join_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_radio_menu_item_new(param0 unsafe.Pointer) {}

func Fn_gtk_radio_menu_item_new_from_widget(param0 unsafe.Pointer) {}

func Fn_gtk_radio_menu_item_new_with_label(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new_with_mnemonic(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_get_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_radio_menu_item_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_radio_tool_button_new(param0 unsafe.Pointer) {}

func Fn_gtk_radio_tool_button_new_from_stock(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_tool_button_new_from_widget(param0 unsafe.Pointer) {}

func Fn_gtk_radio_tool_button_new_with_stock_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_tool_button_get_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_radio_tool_button_set_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_range_get_adjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_fill_level(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_flippable(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_inverted(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_lower_stepper_sensitivity(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_min_slider_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_range_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_range_get_restrict_to_fill_level(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_round_digits(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_show_fill_level(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_slider_range(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_range_get_slider_size_fixed(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_upper_stepper_sensitivity(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_get_value(paramInstance unsafe.Pointer) {}

func Fn_gtk_range_set_adjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_range_set_fill_level(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_range_set_flippable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_range_set_increments(paramInstance unsafe.Pointer, param0 float64, param1 float64) {}

func Fn_gtk_range_set_inverted(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_range_set_lower_stepper_sensitivity(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_range_set_min_slider_size(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_range_set_range(paramInstance unsafe.Pointer, param0 float64, param1 float64) {}

func Fn_gtk_range_set_restrict_to_fill_level(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_range_set_round_digits(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_range_set_show_fill_level(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_range_set_slider_size_fixed(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_range_set_upper_stepper_sensitivity(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_range_set_value(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_rc_style_new() {
	C.gtk_rc_style_new()
}

func Fn_gtk_rc_style_copy(paramInstance unsafe.Pointer) {}

func Fn_gtk_recent_action_new(param0 string, param1 string, param2 string, param3 string) {}

func Fn_gtk_recent_action_new_for_manager(param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer) {
}

func Fn_gtk_recent_action_get_show_numbers(paramInstance unsafe.Pointer) {}

func Fn_gtk_recent_action_set_show_numbers(paramInstance unsafe.Pointer, param0 bool) {}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_for_manager : has varargs

func Fn_gtk_recent_chooser_menu_new() {
	C.gtk_recent_chooser_menu_new()
}

func Fn_gtk_recent_chooser_menu_new_for_manager(param0 unsafe.Pointer) {}

func Fn_gtk_recent_chooser_menu_get_show_numbers(paramInstance unsafe.Pointer) {}

func Fn_gtk_recent_chooser_menu_set_show_numbers(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_recent_chooser_widget_new() {
	C.gtk_recent_chooser_widget_new()
}

func Fn_gtk_recent_chooser_widget_new_for_manager(param0 unsafe.Pointer) {}

func Fn_gtk_recent_filter_new() {
	C.gtk_recent_filter_new()
}

func Fn_gtk_recent_filter_add_age(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_recent_filter_add_application(paramInstance unsafe.Pointer, param0 string) {}

// UNSUPPORTED : add_custom : has callback

func Fn_gtk_recent_filter_add_group(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_recent_filter_add_mime_type(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_recent_filter_add_pattern(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_recent_filter_add_pixbuf_formats(paramInstance unsafe.Pointer) {}

func Fn_gtk_recent_filter_filter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_recent_filter_get_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_recent_filter_get_needed(paramInstance unsafe.Pointer) {}

func Fn_gtk_recent_filter_set_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_recent_manager_new() {
	C.gtk_recent_manager_new()
}

func Fn_gtk_recent_manager_add_full(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_recent_manager_add_item(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_recent_manager_get_items(paramInstance unsafe.Pointer) {}

func Fn_gtk_recent_manager_has_item(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_recent_manager_lookup_item(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_recent_manager_move_item(paramInstance unsafe.Pointer, param0 string, param1 string) {}

func Fn_gtk_recent_manager_purge_items(paramInstance unsafe.Pointer) {}

func Fn_gtk_recent_manager_remove_item(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_recent_manager_get_default() {
	C.gtk_recent_manager_get_default()
}

func Fn_gtk_renderer_cell_accessible_new(param0 unsafe.Pointer) {}

func Fn_gtk_revealer_new() {
	C.gtk_revealer_new()
}

func Fn_gtk_revealer_get_child_revealed(paramInstance unsafe.Pointer) {}

func Fn_gtk_revealer_get_reveal_child(paramInstance unsafe.Pointer) {}

func Fn_gtk_revealer_get_transition_duration(paramInstance unsafe.Pointer) {}

func Fn_gtk_revealer_get_transition_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_revealer_set_reveal_child(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_revealer_set_transition_duration(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_revealer_set_transition_type(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_scale_new(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_scale_new_with_range(param0 int, param1 float64, param2 float64, param3 float64) {}

func Fn_gtk_scale_add_mark(paramInstance unsafe.Pointer, param0 float64, param1 int, param2 string) {}

func Fn_gtk_scale_clear_marks(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_get_digits(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_get_draw_value(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_get_has_origin(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_get_layout(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_get_layout_offsets(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_scale_get_value_pos(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_set_digits(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_scale_set_draw_value(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_scale_set_has_origin(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_scale_set_value_pos(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_scale_button_new(param0 int, param1 float64, param2 float64, param3 float64, param4 []string) {
}

func Fn_gtk_scale_button_get_adjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_button_get_minus_button(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_button_get_plus_button(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_button_get_popup(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_button_get_value(paramInstance unsafe.Pointer) {}

func Fn_gtk_scale_button_set_adjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_scale_button_set_icons(paramInstance unsafe.Pointer, param0 []string) {}

func Fn_gtk_scale_button_set_value(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_scrollbar_new(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_add_with_viewport(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_capture_button_press(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_hadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_hscrollbar(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_kinetic_scrolling(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_min_content_height(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_min_content_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_placement(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_policy(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_scrolled_window_get_shadow_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_vadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_vscrollbar(paramInstance unsafe.Pointer) {}

func Fn_gtk_scrolled_window_set_capture_button_press(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_scrolled_window_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_set_kinetic_scrolling(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_scrolled_window_set_min_content_height(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_scrolled_window_set_min_content_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_scrolled_window_set_placement(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_scrolled_window_set_policy(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_scrolled_window_set_shadow_type(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_scrolled_window_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_unset_placement(paramInstance unsafe.Pointer) {}

func Fn_gtk_search_bar_new() {
	C.gtk_search_bar_new()
}

func Fn_gtk_search_bar_connect_entry(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_search_bar_get_search_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_search_bar_get_show_close_button(paramInstance unsafe.Pointer) {}

func Fn_gtk_search_bar_handle_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_search_bar_set_search_mode(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_search_bar_set_show_close_button(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_search_entry_new() {
	C.gtk_search_entry_new()
}

func Fn_gtk_separator_new(param0 int) {}

func Fn_gtk_separator_menu_item_new() {
	C.gtk_separator_menu_item_new()
}

func Fn_gtk_separator_tool_item_new() {
	C.gtk_separator_tool_item_new()
}

func Fn_gtk_separator_tool_item_get_draw(paramInstance unsafe.Pointer) {}

func Fn_gtk_separator_tool_item_set_draw(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_settings_set_double_property(paramInstance unsafe.Pointer, param0 string, param1 float64, param2 string) {
}

func Fn_gtk_settings_set_long_property(paramInstance unsafe.Pointer, param0 string, param1 int64, param2 string) {
}

func Fn_gtk_settings_set_property_value(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_settings_set_string_property(paramInstance unsafe.Pointer, param0 string, param1 string, param2 string) {
}

func Fn_gtk_settings_get_default() {
	C.gtk_settings_get_default()
}

func Fn_gtk_settings_get_for_screen(param0 unsafe.Pointer) {}

func Fn_gtk_settings_install_property(param0 unsafe.Pointer) {}

// UNSUPPORTED : install_property_parser : has callback

func Fn_gtk_size_group_new(param0 int) {}

func Fn_gtk_size_group_add_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_size_group_get_ignore_hidden(paramInstance unsafe.Pointer) {}

func Fn_gtk_size_group_get_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_size_group_get_widgets(paramInstance unsafe.Pointer) {}

func Fn_gtk_size_group_remove_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_size_group_set_ignore_hidden(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_size_group_set_mode(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_socket_new() {
	C.gtk_socket_new()
}

func Fn_gtk_socket_add_id(paramInstance unsafe.Pointer, param0 uint64) {}

func Fn_gtk_socket_get_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_socket_get_plug_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_spin_button_new(param0 unsafe.Pointer, param1 float64, param2 uint) {}

func Fn_gtk_spin_button_new_with_range(param0 float64, param1 float64, param2 float64) {}

func Fn_gtk_spin_button_configure(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 uint) {
}

func Fn_gtk_spin_button_get_adjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_spin_button_get_digits(paramInstance unsafe.Pointer) {}

func Fn_gtk_spin_button_get_increments(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) {
}

func Fn_gtk_spin_button_get_numeric(paramInstance unsafe.Pointer) {}

func Fn_gtk_spin_button_get_range(paramInstance unsafe.Pointer, param0 *float64, param1 *float64) {}

func Fn_gtk_spin_button_get_snap_to_ticks(paramInstance unsafe.Pointer) {}

func Fn_gtk_spin_button_get_update_policy(paramInstance unsafe.Pointer) {}

func Fn_gtk_spin_button_get_value(paramInstance unsafe.Pointer) {}

func Fn_gtk_spin_button_get_value_as_int(paramInstance unsafe.Pointer) {}

func Fn_gtk_spin_button_get_wrap(paramInstance unsafe.Pointer) {}

func Fn_gtk_spin_button_set_adjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_spin_button_set_digits(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_spin_button_set_increments(paramInstance unsafe.Pointer, param0 float64, param1 float64) {
}

func Fn_gtk_spin_button_set_numeric(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_spin_button_set_range(paramInstance unsafe.Pointer, param0 float64, param1 float64) {}

func Fn_gtk_spin_button_set_snap_to_ticks(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_spin_button_set_update_policy(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_spin_button_set_value(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_spin_button_set_wrap(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_spin_button_spin(paramInstance unsafe.Pointer, param0 int, param1 float64) {}

func Fn_gtk_spin_button_update(paramInstance unsafe.Pointer) {}

func Fn_gtk_spinner_new() {
	C.gtk_spinner_new()
}

func Fn_gtk_spinner_start(paramInstance unsafe.Pointer) {}

func Fn_gtk_spinner_stop(paramInstance unsafe.Pointer) {}

func Fn_gtk_stack_new() {
	C.gtk_stack_new()
}

func Fn_gtk_stack_add_named(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_stack_add_titled(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 string) {
}

func Fn_gtk_stack_get_child_by_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_stack_get_homogeneous(paramInstance unsafe.Pointer) {}

func Fn_gtk_stack_get_transition_duration(paramInstance unsafe.Pointer) {}

func Fn_gtk_stack_get_transition_running(paramInstance unsafe.Pointer) {}

func Fn_gtk_stack_get_transition_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_stack_get_visible_child(paramInstance unsafe.Pointer) {}

func Fn_gtk_stack_get_visible_child_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_stack_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_stack_set_transition_duration(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_stack_set_transition_type(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_stack_set_visible_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_stack_set_visible_child_full(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_stack_set_visible_child_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_stack_switcher_new() {
	C.gtk_stack_switcher_new()
}

func Fn_gtk_stack_switcher_get_stack(paramInstance unsafe.Pointer) {}

func Fn_gtk_stack_switcher_set_stack(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_new() {
	C.gtk_status_icon_new()
}

func Fn_gtk_status_icon_new_from_file(param0 string) {}

func Fn_gtk_status_icon_new_from_gicon(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_new_from_icon_name(param0 string) {}

func Fn_gtk_status_icon_new_from_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_new_from_stock(param0 string) {}

func Fn_gtk_status_icon_get_geometry(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
}

func Fn_gtk_status_icon_get_gicon(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_has_tooltip(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_icon_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_pixbuf(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_screen(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_stock(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_storage_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_tooltip_markup(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_tooltip_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_get_x11_window_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_is_embedded(paramInstance unsafe.Pointer) {}

func Fn_gtk_status_icon_set_from_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_status_icon_set_from_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_set_from_icon_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_status_icon_set_from_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_set_from_stock(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_status_icon_set_has_tooltip(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_status_icon_set_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_status_icon_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_set_title(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_status_icon_set_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_status_icon_set_tooltip_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_status_icon_set_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_status_icon_position_menu(param0 unsafe.Pointer, param1 *int, param2 *int, param3 *bool, param4 *unsafe.Pointer) {
}

func Fn_gtk_statusbar_new() {
	C.gtk_statusbar_new()
}

func Fn_gtk_statusbar_get_context_id(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_statusbar_get_message_area(paramInstance unsafe.Pointer) {}

func Fn_gtk_statusbar_pop(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_statusbar_push(paramInstance unsafe.Pointer, param0 uint, param1 string) {}

func Fn_gtk_statusbar_remove(paramInstance unsafe.Pointer, param0 uint, param1 uint) {}

func Fn_gtk_statusbar_remove_all(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_style_new() {
	C.gtk_style_new()
}

func Fn_gtk_style_apply_default_background(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int, param5 int, param6 int) {
}

func Fn_gtk_style_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_style_copy(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_detach(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : get : has varargs

func Fn_gtk_style_get_style_property(paramInstance unsafe.Pointer, param0 uint64, param1 string, param2 unsafe.Pointer) {
}

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_style_has_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_style_lookup_icon_set(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_style_render_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 unsafe.Pointer, param5 string) {
}

func Fn_gtk_style_set_background(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_style_context_new() {
	C.gtk_style_context_new()
}

func Fn_gtk_style_context_add_class(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_style_context_add_provider(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint) {
}

func Fn_gtk_style_context_add_region(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_style_context_cancel_animations(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {}

// UNSUPPORTED : get : has varargs

func Fn_gtk_style_context_get_background_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_style_context_get_border(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_style_context_get_border_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_style_context_get_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_style_context_get_direction(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_get_font(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_style_context_get_frame_clock(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_get_junction_sides(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_get_margin(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_style_context_get_padding(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_style_context_get_parent(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_get_path(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_get_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
}

func Fn_gtk_style_context_get_scale(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_get_screen(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_get_section(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_style_context_get_state(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : get_style : has varargs

func Fn_gtk_style_context_get_style_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

// UNSUPPORTED : get_style_valist : has va_list

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_style_context_has_class(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_style_context_has_region(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_style_context_invalidate(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_list_classes(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_list_regions(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_style_context_lookup_icon_set(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_style_context_notify_state_change(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 int, param3 bool) {
}

func Fn_gtk_style_context_pop_animatable_region(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_push_animatable_region(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {
}

func Fn_gtk_style_context_remove_class(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_style_context_remove_provider(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_style_context_remove_region(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_style_context_restore(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_save(paramInstance unsafe.Pointer) {}

func Fn_gtk_style_context_scroll_animations(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
}

func Fn_gtk_style_context_set_background(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_direction(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_style_context_set_frame_clock(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_junction_sides(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_style_context_set_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_scale(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_style_context_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_state(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_style_context_state_is_running(paramInstance unsafe.Pointer, param0 int, param1 *float64) {
}

func Fn_gtk_style_context_add_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint) {
}

func Fn_gtk_style_context_remove_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_reset_widgets(param0 unsafe.Pointer) {}

func Fn_gtk_style_properties_new() {
	C.gtk_style_properties_new()
}

func Fn_gtk_style_properties_clear(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : get : has varargs

func Fn_gtk_style_properties_get_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
}

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_style_properties_lookup_color(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_style_properties_map_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_style_properties_merge(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

// UNSUPPORTED : set : has varargs

func Fn_gtk_style_properties_set_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
}

// UNSUPPORTED : set_valist : has va_list

func Fn_gtk_style_properties_unset_property(paramInstance unsafe.Pointer, param0 string, param1 int) {}

// UNSUPPORTED : lookup_property : has callback

// UNSUPPORTED : register_property : has callback

func Fn_gtk_switch_new() {
	C.gtk_switch_new()
}

func Fn_gtk_switch_get_active(paramInstance unsafe.Pointer) {}

func Fn_gtk_switch_set_active(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_table_new(param0 uint, param1 uint, param2 bool) {}

func Fn_gtk_table_attach(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint, param5 int, param6 int, param7 uint, param8 uint) {
}

func Fn_gtk_table_attach_defaults(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint) {
}

func Fn_gtk_table_get_col_spacing(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_table_get_default_col_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_table_get_default_row_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_table_get_homogeneous(paramInstance unsafe.Pointer) {}

func Fn_gtk_table_get_row_spacing(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_table_get_size(paramInstance unsafe.Pointer, param0 *uint, param1 *uint) {}

func Fn_gtk_table_resize(paramInstance unsafe.Pointer, param0 uint, param1 uint) {}

func Fn_gtk_table_set_col_spacing(paramInstance unsafe.Pointer, param0 uint, param1 uint) {}

func Fn_gtk_table_set_col_spacings(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_table_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_table_set_row_spacing(paramInstance unsafe.Pointer, param0 uint, param1 uint) {}

func Fn_gtk_table_set_row_spacings(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_tearoff_menu_item_new() {
	C.gtk_tearoff_menu_item_new()
}

func Fn_gtk_text_buffer_new(param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_add_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_add_selection_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_apply_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_apply_tag_by_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_backspace(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool, param2 bool) {
}

func Fn_gtk_text_buffer_begin_user_action(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_copy_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_create_child_anchor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_create_mark(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 bool) {
}

// UNSUPPORTED : create_tag : has varargs

func Fn_gtk_text_buffer_cut_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_text_buffer_delete(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_delete_interactive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
}

func Fn_gtk_text_buffer_delete_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_delete_mark_by_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_text_buffer_delete_selection(paramInstance unsafe.Pointer, param0 bool, param1 bool) {}

func Fn_gtk_text_buffer_deserialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 gdk.Atom, param2 unsafe.Pointer, param3 []uint8, param4 uint64) {
}

func Fn_gtk_text_buffer_deserialize_get_can_create_tags(paramInstance unsafe.Pointer, param0 gdk.Atom) {
}

func Fn_gtk_text_buffer_deserialize_set_can_create_tags(paramInstance unsafe.Pointer, param0 gdk.Atom, param1 bool) {
}

func Fn_gtk_text_buffer_end_user_action(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_bounds(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_get_char_count(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_copy_target_list(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_deserialize_formats(paramInstance unsafe.Pointer, param0 *int) {}

func Fn_gtk_text_buffer_get_end_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_has_selection(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_insert(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_iter_at_child_anchor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_get_iter_at_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_text_buffer_get_iter_at_line_index(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
}

func Fn_gtk_text_buffer_get_iter_at_line_offset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
}

func Fn_gtk_text_buffer_get_iter_at_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_get_iter_at_offset(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_text_buffer_get_line_count(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_mark(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_text_buffer_get_modified(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_paste_target_list(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_selection_bound(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_selection_bounds(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_get_serialize_formats(paramInstance unsafe.Pointer, param0 *int) {}

func Fn_gtk_text_buffer_get_slice(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
}

func Fn_gtk_text_buffer_get_start_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_tag_table(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_text(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
}

func Fn_gtk_text_buffer_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
}

func Fn_gtk_text_buffer_insert_at_cursor(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_text_buffer_insert_child_anchor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_insert_interactive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int, param3 bool) {
}

func Fn_gtk_text_buffer_insert_interactive_at_cursor(paramInstance unsafe.Pointer, param0 string, param1 int, param2 bool) {
}

func Fn_gtk_text_buffer_insert_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_insert_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_insert_range_interactive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) {
}

// UNSUPPORTED : insert_with_tags : has varargs

// UNSUPPORTED : insert_with_tags_by_name : has varargs

func Fn_gtk_text_buffer_move_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_move_mark_by_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_paste_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
}

func Fn_gtk_text_buffer_place_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : register_deserialize_format : has callback

func Fn_gtk_text_buffer_register_deserialize_tagset(paramInstance unsafe.Pointer, param0 string) {}

// UNSUPPORTED : register_serialize_format : has callback

func Fn_gtk_text_buffer_register_serialize_tagset(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_text_buffer_remove_all_tags(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_remove_selection_clipboard(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_remove_tag(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_remove_tag_by_name(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_select_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_serialize(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 gdk.Atom, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 *uint64) {
}

func Fn_gtk_text_buffer_set_modified(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_text_buffer_set_text(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_text_buffer_unregister_deserialize_format(paramInstance unsafe.Pointer, param0 gdk.Atom) {
}

func Fn_gtk_text_buffer_unregister_serialize_format(paramInstance unsafe.Pointer, param0 gdk.Atom) {}

func Fn_gtk_text_child_anchor_new() {
	C.gtk_text_child_anchor_new()
}

func Fn_gtk_text_child_anchor_get_deleted(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_child_anchor_get_widgets(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_mark_new(param0 string, param1 bool) {}

func Fn_gtk_text_mark_get_buffer(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_mark_get_deleted(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_mark_get_left_gravity(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_mark_get_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_mark_get_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_mark_set_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_text_tag_new(param0 string) {}

func Fn_gtk_text_tag_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_tag_get_priority(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_tag_set_priority(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_tag_table_new() {
	C.gtk_text_tag_table_new()
}

func Fn_gtk_text_tag_table_add(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : foreach : has callback

func Fn_gtk_text_tag_table_get_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_tag_table_lookup(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_text_tag_table_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_new() {
	C.gtk_text_view_new()
}

func Fn_gtk_text_view_new_with_buffer(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_add_child_at_anchor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_view_add_child_in_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int) {
}

func Fn_gtk_text_view_backward_display_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_backward_display_line_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_text_view_buffer_to_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 *int, param4 *int) {
}

func Fn_gtk_text_view_forward_display_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_forward_display_line_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_text_view_get_accepts_tab(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_border_window_size(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_get_buffer(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_cursor_locations(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_view_get_cursor_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_default_attributes(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_editable(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_hadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_indent(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_input_hints(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_input_purpose(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_iter_at_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
}

func Fn_gtk_text_view_get_iter_at_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 int, param3 int) {
}

func Fn_gtk_text_view_get_iter_location(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_text_view_get_justification(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_left_margin(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_line_at_y(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 *int) {
}

func Fn_gtk_text_view_get_line_yrange(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
}

func Fn_gtk_text_view_get_overwrite(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_pixels_above_lines(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_pixels_below_lines(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_pixels_inside_wrap(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_right_margin(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_tabs(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_vadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_get_visible_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_get_window(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_get_window_type(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_get_wrap_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_im_context_filter_keypress(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_text_view_move_child(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int) {
}

func Fn_gtk_text_view_move_mark_onscreen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_move_visually(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_text_view_place_cursor_onscreen(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_reset_im_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_text_view_scroll_mark_onscreen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_scroll_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 bool, param3 float64, param4 float64) {
}

func Fn_gtk_text_view_scroll_to_mark(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 float64, param2 bool, param3 float64, param4 float64) {
}

func Fn_gtk_text_view_set_accepts_tab(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_text_view_set_border_window_size(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_text_view_set_buffer(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_set_cursor_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_text_view_set_editable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_text_view_set_indent(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_set_input_hints(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_set_input_purpose(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_set_justification(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_set_left_margin(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_set_overwrite(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_text_view_set_pixels_above_lines(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_set_pixels_below_lines(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_set_pixels_inside_wrap(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_set_right_margin(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_set_tabs(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_set_wrap_mode(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_text_view_starts_display_line(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_text_view_window_to_buffer_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 *int, param4 *int) {
}

// UNSUPPORTED : get : has varargs

func Fn_gtk_theming_engine_get_background_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_theming_engine_get_border(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_theming_engine_get_border_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_theming_engine_get_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_theming_engine_get_direction(paramInstance unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_font(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_theming_engine_get_junction_sides(paramInstance unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_margin(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_theming_engine_get_padding(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_theming_engine_get_path(paramInstance unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_property(paramInstance unsafe.Pointer, param0 string, param1 int, param2 unsafe.Pointer) {
}

func Fn_gtk_theming_engine_get_screen(paramInstance unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_state(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : get_style : has varargs

func Fn_gtk_theming_engine_get_style_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

// UNSUPPORTED : get_style_valist : has va_list

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_theming_engine_has_class(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_theming_engine_has_region(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_theming_engine_lookup_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_theming_engine_state_is_running(paramInstance unsafe.Pointer, param0 int, param1 *float64) {
}

func Fn_gtk_theming_engine_load(param0 string) {}

// UNSUPPORTED : register_property : has callback

func Fn_gtk_toggle_action_new(param0 string, param1 string, param2 string, param3 string) {}

func Fn_gtk_toggle_action_get_active(paramInstance unsafe.Pointer) {}

func Fn_gtk_toggle_action_get_draw_as_radio(paramInstance unsafe.Pointer) {}

func Fn_gtk_toggle_action_set_active(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_toggle_action_set_draw_as_radio(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_toggle_action_toggled(paramInstance unsafe.Pointer) {}

func Fn_gtk_toggle_button_new() {
	C.gtk_toggle_button_new()
}

func Fn_gtk_toggle_button_new_with_label(param0 string) {}

func Fn_gtk_toggle_button_new_with_mnemonic(param0 string) {}

func Fn_gtk_toggle_button_get_active(paramInstance unsafe.Pointer) {}

func Fn_gtk_toggle_button_get_inconsistent(paramInstance unsafe.Pointer) {}

func Fn_gtk_toggle_button_get_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_toggle_button_set_active(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_toggle_button_set_inconsistent(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_toggle_button_set_mode(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_toggle_button_toggled(paramInstance unsafe.Pointer) {}

func Fn_gtk_toggle_tool_button_new() {
	C.gtk_toggle_tool_button_new()
}

func Fn_gtk_toggle_tool_button_new_from_stock(param0 string) {}

func Fn_gtk_toggle_tool_button_get_active(paramInstance unsafe.Pointer) {}

func Fn_gtk_toggle_tool_button_set_active(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tool_button_new(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_tool_button_new_from_stock(param0 string) {}

func Fn_gtk_tool_button_get_icon_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_button_get_icon_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_button_get_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_button_get_label_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_button_get_stock_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_button_get_use_underline(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_button_set_icon_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tool_button_set_icon_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tool_button_set_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tool_button_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tool_button_set_stock_id(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tool_button_set_use_underline(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tool_item_new() {
	C.gtk_tool_item_new()
}

func Fn_gtk_tool_item_get_ellipsize_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_expand(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_homogeneous(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_icon_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_is_important(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_orientation(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_proxy_menu_item(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tool_item_get_relief_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_text_alignment(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_text_orientation(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_text_size_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_toolbar_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_use_drag_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_visible_horizontal(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_get_visible_vertical(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_rebuild_menu(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_retrieve_proxy_menu_item(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_set_expand(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tool_item_set_homogeneous(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tool_item_set_is_important(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tool_item_set_proxy_menu_item(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_tool_item_set_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tool_item_set_tooltip_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tool_item_set_use_drag_window(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tool_item_set_visible_horizontal(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tool_item_set_visible_vertical(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tool_item_toolbar_reconfigured(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_group_new(param0 string) {}

func Fn_gtk_tool_item_group_get_collapsed(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_group_get_drop_item(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_tool_item_group_get_ellipsize(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_group_get_header_relief(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_group_get_item_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tool_item_group_get_label(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_group_get_label_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_group_get_n_items(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_item_group_get_nth_item(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_tool_item_group_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_tool_item_group_set_collapsed(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tool_item_group_set_ellipsize(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tool_item_group_set_header_relief(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tool_item_group_set_item_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_tool_item_group_set_label(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tool_item_group_set_label_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_new() {
	C.gtk_tool_palette_new()
}

func Fn_gtk_tool_palette_add_drag_dest(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 int) {
}

func Fn_gtk_tool_palette_get_drag_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_drop_group(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_tool_palette_get_drop_item(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_tool_palette_get_exclusive(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_expand(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_group_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_hadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_icon_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_vadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_palette_set_drag_source(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tool_palette_set_exclusive(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_tool_palette_set_expand(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_tool_palette_set_group_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_tool_palette_set_icon_size(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tool_palette_set_style(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tool_palette_unset_icon_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_palette_unset_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_drag_target_group() {
	C.gtk_tool_palette_get_drag_target_group()
}

func Fn_gtk_tool_palette_get_drag_target_item() {
	C.gtk_tool_palette_get_drag_target_item()
}

func Fn_gtk_toolbar_new() {
	C.gtk_toolbar_new()
}

func Fn_gtk_toolbar_get_drop_index(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_toolbar_get_icon_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_toolbar_get_item_index(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_toolbar_get_n_items(paramInstance unsafe.Pointer) {}

func Fn_gtk_toolbar_get_nth_item(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_toolbar_get_relief_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_toolbar_get_show_arrow(paramInstance unsafe.Pointer) {}

func Fn_gtk_toolbar_get_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_toolbar_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_toolbar_set_drop_highlight_item(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_toolbar_set_icon_size(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_toolbar_set_show_arrow(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_toolbar_set_style(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_toolbar_unset_icon_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_toolbar_unset_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_tooltip_set_custom(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tooltip_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tooltip_set_icon_from_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_tooltip_set_icon_from_icon_name(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_tooltip_set_icon_from_stock(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_tooltip_set_markup(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tooltip_set_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tooltip_set_tip_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tooltip_trigger_tooltip_query(param0 unsafe.Pointer) {}

func Fn_gtk_toplevel_accessible_get_children(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_model_filter_clear_cache(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_model_filter_convert_child_iter_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_model_filter_convert_child_path_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_tree_model_filter_convert_iter_to_child_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_model_filter_convert_path_to_child_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_tree_model_filter_get_model(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_model_filter_refilter(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : set_modify_func : has callback

func Fn_gtk_tree_model_filter_set_visible_column(paramInstance unsafe.Pointer, param0 int) {}

// UNSUPPORTED : set_visible_func : has callback

func Fn_gtk_tree_model_sort_clear_cache(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_model_sort_convert_child_iter_to_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_model_sort_convert_child_path_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_tree_model_sort_convert_iter_to_child_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_model_sort_convert_path_to_child_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_tree_model_sort_get_model(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_model_sort_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_model_sort_reset_default_sort_func(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_selection_count_selected_rows(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_selection_get_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_selection_get_select_function(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_selection_get_selected(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_selection_get_selected_rows(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {}

func Fn_gtk_tree_selection_get_tree_view(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_selection_get_user_data(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_selection_iter_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_path_is_selected(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_select_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_selection_select_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_select_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_select_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_tree_selection_set_mode(paramInstance unsafe.Pointer, param0 int) {}

// UNSUPPORTED : set_select_function : has callback

func Fn_gtk_tree_selection_unselect_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_selection_unselect_iter(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_unselect_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_unselect_range(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

// UNSUPPORTED : new : has varargs

func Fn_gtk_tree_store_newv(param0 int, param1 []uint64) {}

func Fn_gtk_tree_store_append(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_store_clear(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_store_insert(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
}

func Fn_gtk_tree_store_insert_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_tree_store_insert_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

// UNSUPPORTED : insert_with_values : has varargs

func Fn_gtk_tree_store_insert_with_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 []int, param4 []gobject.Value, param5 int) {
}

func Fn_gtk_tree_store_is_ancestor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_store_iter_depth(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_store_iter_is_valid(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_store_move_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_store_move_before(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_store_prepend(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_store_remove(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_store_reorder(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []int) {
}

// UNSUPPORTED : set : has varargs

func Fn_gtk_tree_store_set_column_types(paramInstance unsafe.Pointer, param0 int, param1 []uint64) {
}

// UNSUPPORTED : set_valist : has va_list

func Fn_gtk_tree_store_set_value(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {
}

func Fn_gtk_tree_store_set_valuesv(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 []int, param2 []gobject.Value, param3 int) {
}

func Fn_gtk_tree_store_swap(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_view_new() {
	C.gtk_tree_view_new()
}

func Fn_gtk_tree_view_new_with_model(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_append_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_collapse_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_collapse_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_columns_autosize(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_convert_bin_window_to_tree_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_bin_window_to_widget_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_tree_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_tree_to_widget_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_widget_to_bin_window_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_widget_to_tree_coords(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_create_row_drag_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_enable_model_drag_dest(paramInstance unsafe.Pointer, param0 []TargetEntry, param1 int, param2 int) {
}

func Fn_gtk_tree_view_enable_model_drag_source(paramInstance unsafe.Pointer, param0 int, param1 []TargetEntry, param2 int, param3 int) {
}

func Fn_gtk_tree_view_expand_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_expand_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_tree_view_expand_to_path(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_get_activate_on_single_click(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_background_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_tree_view_get_bin_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_cell_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_tree_view_get_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_get_columns(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_cursor(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) {
}

func Fn_gtk_tree_view_get_dest_row_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 int) {
}

func Fn_gtk_tree_view_get_drag_dest_row(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 int) {
}

func Fn_gtk_tree_view_get_enable_search(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_enable_tree_lines(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_expander_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_fixed_height_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_grid_lines(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_hadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_headers_clickable(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_headers_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_hover_expand(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_hover_selection(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_level_indentation(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_model(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_n_columns(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_path_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *int, param5 *int) {
}

func Fn_gtk_tree_view_get_reorderable(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_row_separator_func(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_rubber_banding(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_rules_hint(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_search_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_search_entry(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_search_equal_func(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_search_position_func(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_selection(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_show_expanders(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_tooltip_column(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_tooltip_context(paramInstance unsafe.Pointer, param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_gtk_tree_view_get_vadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_get_visible_range(paramInstance unsafe.Pointer, param0 *unsafe.Pointer, param1 *unsafe.Pointer) {
}

func Fn_gtk_tree_view_get_visible_rect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_insert_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

// UNSUPPORTED : insert_column_with_attributes : has varargs

// UNSUPPORTED : insert_column_with_data_func : has callback

func Fn_gtk_tree_view_is_blank_at_pos(paramInstance unsafe.Pointer, param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *int, param5 *int) {
}

func Fn_gtk_tree_view_is_rubber_banding_active(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : map_expanded_rows : has callback

func Fn_gtk_tree_view_move_column_after(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_view_remove_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_row_activated(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_view_row_expanded(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_scroll_to_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 float32, param4 float32) {
}

func Fn_gtk_tree_view_scroll_to_point(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_tree_view_set_activate_on_single_click(paramInstance unsafe.Pointer, param0 bool) {}

// UNSUPPORTED : set_column_drag_function : has callback

func Fn_gtk_tree_view_set_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
}

func Fn_gtk_tree_view_set_cursor_on_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) {
}

// UNSUPPORTED : set_destroy_count_func : has callback

func Fn_gtk_tree_view_set_drag_dest_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_tree_view_set_enable_search(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_enable_tree_lines(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_expander_column(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_set_fixed_height_mode(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_grid_lines(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_set_headers_clickable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_headers_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_hover_expand(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_hover_selection(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_level_indentation(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_set_model(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_set_reorderable(paramInstance unsafe.Pointer, param0 bool) {}

// UNSUPPORTED : set_row_separator_func : has callback

func Fn_gtk_tree_view_set_rubber_banding(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_rules_hint(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_search_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_set_search_entry(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : set_search_equal_func : has callback

// UNSUPPORTED : set_search_position_func : has callback

func Fn_gtk_tree_view_set_show_expanders(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_set_tooltip_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer) {
}

func Fn_gtk_tree_view_set_tooltip_column(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_set_tooltip_row(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_view_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_unset_rows_drag_dest(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_unset_rows_drag_source(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_new() {
	C.gtk_tree_view_column_new()
}

func Fn_gtk_tree_view_column_new_with_area(param0 unsafe.Pointer) {}

// UNSUPPORTED : new_with_attributes : has varargs

func Fn_gtk_tree_view_column_add_attribute(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 string, param2 int) {
}

func Fn_gtk_tree_view_column_cell_get_position(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int) {
}

func Fn_gtk_tree_view_column_cell_get_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 *int, param2 *int, param3 *int, param4 *int) {
}

func Fn_gtk_tree_view_column_cell_is_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_cell_set_cell_data(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 bool) {
}

func Fn_gtk_tree_view_column_clear(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_clear_attributes(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_column_clicked(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_focus_cell(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_alignment(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_button(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_clickable(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_expand(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_fixed_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_max_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_min_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_reorderable(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_resizable(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_sizing(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_sort_column_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_sort_indicator(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_sort_order(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_spacing(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_tree_view(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_x_offset(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_pack_end(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_tree_view_column_pack_start(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_tree_view_column_queue_resize(paramInstance unsafe.Pointer) {}

func Fn_gtk_tree_view_column_set_alignment(paramInstance unsafe.Pointer, param0 float32) {}

// UNSUPPORTED : set_attributes : has varargs

// UNSUPPORTED : set_cell_data_func : has callback

func Fn_gtk_tree_view_column_set_clickable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_column_set_expand(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_column_set_fixed_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_column_set_max_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_column_set_min_width(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_column_set_reorderable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_column_set_resizable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_column_set_sizing(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_column_set_sort_column_id(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_column_set_sort_indicator(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_column_set_sort_order(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_column_set_spacing(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_tree_view_column_set_title(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_tree_view_column_set_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_tree_view_column_set_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_ui_manager_new() {
	C.gtk_ui_manager_new()
}

func Fn_gtk_ui_manager_add_ui(paramInstance unsafe.Pointer, param0 uint, param1 string, param2 string, param3 string, param4 int, param5 bool) {
}

func Fn_gtk_ui_manager_add_ui_from_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_ui_manager_add_ui_from_resource(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_ui_manager_add_ui_from_string(paramInstance unsafe.Pointer, param0 string, param1 uint64) {
}

func Fn_gtk_ui_manager_ensure_update(paramInstance unsafe.Pointer) {}

func Fn_gtk_ui_manager_get_accel_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_ui_manager_get_action(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_ui_manager_get_action_groups(paramInstance unsafe.Pointer) {}

func Fn_gtk_ui_manager_get_add_tearoffs(paramInstance unsafe.Pointer) {}

func Fn_gtk_ui_manager_get_toplevels(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_ui_manager_get_ui(paramInstance unsafe.Pointer) {}

func Fn_gtk_ui_manager_get_widget(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_ui_manager_insert_action_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_ui_manager_new_merge_id(paramInstance unsafe.Pointer) {}

func Fn_gtk_ui_manager_remove_action_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_ui_manager_remove_ui(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_ui_manager_set_add_tearoffs(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_vbox_new(param0 bool, param1 int) {}

func Fn_gtk_vbutton_box_new() {
	C.gtk_vbutton_box_new()
}

func Fn_gtk_vpaned_new() {
	C.gtk_vpaned_new()
}

func Fn_gtk_vscale_new(param0 unsafe.Pointer) {}

func Fn_gtk_vscale_new_with_range(param0 float64, param1 float64, param2 float64) {}

func Fn_gtk_vscrollbar_new(param0 unsafe.Pointer) {}

func Fn_gtk_vseparator_new() {
	C.gtk_vseparator_new()
}

func Fn_gtk_viewport_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_viewport_get_bin_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_viewport_get_hadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_viewport_get_shadow_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_viewport_get_vadjustment(paramInstance unsafe.Pointer) {}

func Fn_gtk_viewport_get_view_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_viewport_set_hadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_viewport_set_shadow_type(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_viewport_set_vadjustment(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_volume_button_new() {
	C.gtk_volume_button_new()
}

// UNSUPPORTED : new : has varargs

func Fn_gtk_widget_activate(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_add_accelerator(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer, param2 uint, param3 int, param4 int) {
}

func Fn_gtk_widget_add_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_widget_add_events(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_add_mnemonic_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

// UNSUPPORTED : add_tick_callback : has callback

func Fn_gtk_widget_can_activate_accel(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_widget_child_focus(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_child_notify(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_widget_class_path(paramInstance unsafe.Pointer, param0 *uint, param1 string, param2 string) {
}

func Fn_gtk_widget_compute_expand(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_create_pango_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_create_pango_layout(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_widget_destroy(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_destroyed(paramInstance unsafe.Pointer, param0 *unsafe.Pointer) {}

func Fn_gtk_widget_device_is_shadowed(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_drag_begin(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer) {
}

func Fn_gtk_drag_begin_with_coordinates(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer, param4 int, param5 int) {
}

func Fn_gtk_drag_check_threshold(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) {
}

func Fn_gtk_drag_dest_add_image_targets(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_dest_add_text_targets(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_dest_add_uri_targets(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_dest_find_target(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_drag_dest_get_target_list(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_dest_get_track_motion(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_dest_set(paramInstance unsafe.Pointer, param0 int, param1 []TargetEntry, param2 int, param3 int) {
}

func Fn_gtk_drag_dest_set_proxy(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 bool) {
}

func Fn_gtk_drag_dest_set_target_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_drag_dest_set_track_motion(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_drag_dest_unset(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_get_data(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 gdk.Atom, param2 uint32) {
}

func Fn_gtk_drag_highlight(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_source_add_image_targets(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_source_add_text_targets(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_source_add_uri_targets(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_source_get_target_list(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_source_set(paramInstance unsafe.Pointer, param0 int, param1 []TargetEntry, param2 int, param3 int) {
}

func Fn_gtk_drag_source_set_icon_gicon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_drag_source_set_icon_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_drag_source_set_icon_pixbuf(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_drag_source_set_icon_stock(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_drag_source_set_target_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_drag_source_unset(paramInstance unsafe.Pointer) {}

func Fn_gtk_drag_unhighlight(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_draw(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_ensure_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_error_bell(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_freeze_child_notify(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_accessible(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_allocated_baseline(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_allocated_height(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_allocated_width(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_allocation(paramInstance unsafe.Pointer, param0 gdk.Rectangle) {}

func Fn_gtk_widget_get_ancestor(paramInstance unsafe.Pointer, param0 uint64) {}

func Fn_gtk_widget_get_app_paintable(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_can_default(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_can_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_child_requisition(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_get_child_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_clipboard(paramInstance unsafe.Pointer, param0 gdk.Atom) {}

func Fn_gtk_widget_get_composite_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_device_enabled(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_get_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_get_direction(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_display(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_double_buffered(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_events(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_frame_clock(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_halign(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_has_tooltip(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_has_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_hexpand(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_hexpand_set(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_mapped(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_margin_bottom(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_margin_end(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_margin_left(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_margin_right(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_margin_start(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_margin_top(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_modifier_mask(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_get_modifier_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_no_show_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_opacity(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_pango_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_parent(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_parent_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_path(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_pointer(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_widget_get_preferred_height(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_widget_get_preferred_height_and_baseline_for_width(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int, param3 *int, param4 *int) {
}

func Fn_gtk_widget_get_preferred_height_for_width(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
}

func Fn_gtk_widget_get_preferred_size(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_widget_get_preferred_width(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_widget_get_preferred_width_for_height(paramInstance unsafe.Pointer, param0 int, param1 *int, param2 *int) {
}

func Fn_gtk_widget_get_realized(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_receives_default(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_request_mode(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_requisition(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_get_root_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_scale_factor(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_screen(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_sensitive(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_settings(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_size_request(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_widget_get_state(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_state_flags(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_style_context(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_support_multidevice(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_template_child(paramInstance unsafe.Pointer, param0 uint64, param1 string) {}

func Fn_gtk_widget_get_tooltip_markup(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_tooltip_text(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_tooltip_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_toplevel(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_valign(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_valign_with_baseline(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_vexpand(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_vexpand_set(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_visual(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_get_window(paramInstance unsafe.Pointer) {}

func Fn_gtk_grab_add(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_grab_default(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_grab_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_grab_remove(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_has_default(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_has_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_has_grab(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_has_rc_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_has_screen(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_has_visible_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_hide(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_hide_on_delete(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_in_destruction(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_init_template(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_input_shape_combine_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_insert_action_group(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_widget_intersect(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_widget_is_ancestor(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_is_composited(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_is_drawable(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_is_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_is_sensitive(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_is_toplevel(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_is_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_keynav_failed(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_list_accel_closures(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_list_mnemonic_labels(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_map(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_mnemonic_activate(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_modify_base(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_modify_bg(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_modify_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_widget_modify_fg(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_modify_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_modify_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_modify_text(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_override_background_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {
}

func Fn_gtk_widget_override_color(paramInstance unsafe.Pointer, param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_override_cursor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_widget_override_font(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_override_symbolic_color(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_widget_path(paramInstance unsafe.Pointer, param0 *uint, param1 string, param2 string) {}

func Fn_gtk_widget_queue_compute_expand(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_queue_draw(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_queue_draw_area(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int) {
}

func Fn_gtk_widget_queue_draw_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_queue_resize(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_queue_resize_no_redraw(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_realize(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_region_intersect(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_register_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_remove_accelerator(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 uint, param2 int) {
}

func Fn_gtk_widget_remove_mnemonic_label(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_remove_tick_callback(paramInstance unsafe.Pointer, param0 uint) {}

func Fn_gtk_widget_render_icon(paramInstance unsafe.Pointer, param0 string, param1 int, param2 string) {
}

func Fn_gtk_widget_render_icon_pixbuf(paramInstance unsafe.Pointer, param0 string, param1 int) {}

func Fn_gtk_widget_reparent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_reset_rc_styles(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_reset_style(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_send_expose(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_send_focus_change(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_accel_path(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

func Fn_gtk_widget_set_allocation(paramInstance unsafe.Pointer, param0 gdk.Rectangle) {}

func Fn_gtk_widget_set_app_paintable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_can_default(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_can_focus(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_child_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_composite_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_widget_set_device_enabled(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
}

func Fn_gtk_widget_set_device_events(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
}

func Fn_gtk_widget_set_direction(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_double_buffered(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_events(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_halign(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_has_tooltip(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_has_window(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_hexpand(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_hexpand_set(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_mapped(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_margin_bottom(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_margin_end(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_margin_left(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_margin_right(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_margin_start(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_margin_top(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_widget_set_no_show_all(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_opacity(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_widget_set_parent(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_parent_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_realized(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_receives_default(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_redraw_on_allocate(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_sensitive(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_size_request(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_widget_set_state(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_state_flags(paramInstance unsafe.Pointer, param0 int, param1 bool) {}

func Fn_gtk_widget_set_style(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_support_multidevice(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_tooltip_markup(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_widget_set_tooltip_text(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_widget_set_tooltip_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_valign(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_widget_set_vexpand(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_vexpand_set(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_widget_set_visual(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_shape_combine_region(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_show(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_show_all(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_show_now(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_size_allocate(paramInstance unsafe.Pointer, param0 gdk.Rectangle) {}

func Fn_gtk_widget_size_allocate_with_baseline(paramInstance unsafe.Pointer, param0 gdk.Rectangle, param1 int) {
}

func Fn_gtk_widget_size_request(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_style_attach(paramInstance unsafe.Pointer) {}

// UNSUPPORTED : style_get : has varargs

func Fn_gtk_widget_style_get_property(paramInstance unsafe.Pointer, param0 string, param1 unsafe.Pointer) {
}

// UNSUPPORTED : style_get_valist : has va_list

func Fn_gtk_widget_thaw_child_notify(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_translate_coordinates(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int, param2 int, param3 *int, param4 *int) {
}

func Fn_gtk_widget_trigger_tooltip_query(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_unmap(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_unparent(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_unrealize(paramInstance unsafe.Pointer) {}

func Fn_gtk_widget_unregister_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_widget_unset_state_flags(paramInstance unsafe.Pointer, param0 int) {}

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

func Fn_gtk_widget_set_default_direction(param0 int) {}

func Fn_gtk_window_new(param0 int) {}

func Fn_gtk_window_activate_default(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_activate_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_activate_key(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_add_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_add_mnemonic(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {}

func Fn_gtk_window_begin_move_drag(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 uint32) {
}

func Fn_gtk_window_begin_resize_drag(paramInstance unsafe.Pointer, param0 int, param1 int, param2 int, param3 int, param4 uint32) {
}

func Fn_gtk_window_close(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_deiconify(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_fullscreen(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_accept_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_application(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_attached_to(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_decorated(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_default_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_window_get_default_widget(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_deletable(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_destroy_with_parent(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_focus_on_map(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_focus_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_gravity(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_has_resize_grip(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_hide_titlebar_when_maximized(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_icon(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_icon_list(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_icon_name(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_mnemonic_modifier(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_mnemonics_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_modal(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_opacity(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_position(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_window_get_resizable(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_resize_grip_area(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_get_role(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_screen(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_size(paramInstance unsafe.Pointer, param0 *int, param1 *int) {}

func Fn_gtk_window_get_skip_pager_hint(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_skip_taskbar_hint(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_title(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_transient_for(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_type_hint(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_urgency_hint(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_window_type(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_has_group(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_has_toplevel_focus(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_iconify(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_is_active(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_is_maximized(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_maximize(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_mnemonic_activate(paramInstance unsafe.Pointer, param0 uint, param1 int) {}

func Fn_gtk_window_move(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_window_parse_geometry(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_window_present(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_present_with_time(paramInstance unsafe.Pointer, param0 uint32) {}

func Fn_gtk_window_propagate_key_event(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_remove_accel_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_remove_mnemonic(paramInstance unsafe.Pointer, param0 uint, param1 unsafe.Pointer) {
}

func Fn_gtk_window_reshow_with_initial_size(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_resize(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_window_resize_grip_is_visible(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_resize_to_geometry(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_window_set_accept_focus(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_application(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_set_attached_to(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_set_decorated(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_default(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_set_default_geometry(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_window_set_default_size(paramInstance unsafe.Pointer, param0 int, param1 int) {}

func Fn_gtk_window_set_deletable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_destroy_with_parent(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_focus(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_set_focus_on_map(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_focus_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_geometry_hints(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {
}

func Fn_gtk_window_set_gravity(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_window_set_has_resize_grip(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_has_user_ref_count(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_hide_titlebar_when_maximized(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_icon(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_set_icon_from_file(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_window_set_icon_list(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_set_icon_name(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_window_set_keep_above(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_keep_below(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_mnemonic_modifier(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_window_set_mnemonics_visible(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_modal(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_opacity(paramInstance unsafe.Pointer, param0 float64) {}

func Fn_gtk_window_set_position(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_window_set_resizable(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_role(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_window_set_screen(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_set_skip_pager_hint(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_skip_taskbar_hint(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_startup_id(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_window_set_title(paramInstance unsafe.Pointer, param0 string) {}

func Fn_gtk_window_set_titlebar(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_set_transient_for(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_set_type_hint(paramInstance unsafe.Pointer, param0 int) {}

func Fn_gtk_window_set_urgency_hint(paramInstance unsafe.Pointer, param0 bool) {}

func Fn_gtk_window_set_wmclass(paramInstance unsafe.Pointer, param0 string, param1 string) {}

func Fn_gtk_window_stick(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_unfullscreen(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_unmaximize(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_unstick(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_get_default_icon_list() {
	C.gtk_window_get_default_icon_list()
}

func Fn_gtk_window_get_default_icon_name() {
	C.gtk_window_get_default_icon_name()
}

func Fn_gtk_window_list_toplevels() {
	C.gtk_window_list_toplevels()
}

func Fn_gtk_window_set_auto_startup_notification(param0 bool) {}

func Fn_gtk_window_set_default_icon(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_default_icon_from_file(param0 string) {}

func Fn_gtk_window_set_default_icon_list(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_default_icon_name(param0 string) {}

func Fn_gtk_window_group_new() {
	C.gtk_window_group_new()
}

func Fn_gtk_window_group_add_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}

func Fn_gtk_window_group_get_current_device_grab(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
}

func Fn_gtk_window_group_get_current_grab(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_group_list_windows(paramInstance unsafe.Pointer) {}

func Fn_gtk_window_group_remove_window(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {}
