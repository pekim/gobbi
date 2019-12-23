// Code generated - DO NOT EDIT.
// +build gtk_3.14

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
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

func Fn_gtk_binding_set_by_class(param0 unsafe.Pointer) {}

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

func Fn_gtk_accel_group_new() {
	C.gtk_accel_group_new()
}

func Fn_gtk_accel_group_from_accel_closure(param0 unsafe.Pointer) {}

func Fn_gtk_accel_label_new(param0 string) {}

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

func Fn_gtk_action_new(param0 string, param1 string, param2 string, param3 string) {}

func Fn_gtk_action_bar_new() {
	C.gtk_action_bar_new()
}

func Fn_gtk_action_group_new(param0 string) {}

func Fn_gtk_adjustment_new(param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_alignment_new(param0 float32, param1 float32, param2 float32, param3 float32) {}

func Fn_gtk_app_chooser_button_new(param0 string) {}

func Fn_gtk_app_chooser_dialog_new(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_app_chooser_dialog_new_for_content_type(param0 unsafe.Pointer, param1 int, param2 string) {
}

func Fn_gtk_app_chooser_widget_new(param0 string) {}

func Fn_gtk_application_new(param0 string, param1 int) {}

func Fn_gtk_application_window_new(param0 unsafe.Pointer) {}

func Fn_gtk_arrow_new(param0 int, param1 int) {}

func Fn_gtk_aspect_frame_new(param0 string, param1 float32, param2 float32, param3 float32, param4 bool) {
}

func Fn_gtk_assistant_new() {
	C.gtk_assistant_new()
}

func Fn_gtk_box_new(param0 int, param1 int) {}

func Fn_gtk_builder_new() {
	C.gtk_builder_new()
}

func Fn_gtk_builder_new_from_file(param0 string) {}

func Fn_gtk_builder_new_from_resource(param0 string) {}

func Fn_gtk_builder_new_from_string(param0 string, param1 uint64) {}

func Fn_gtk_button_new() {
	C.gtk_button_new()
}

func Fn_gtk_button_new_from_icon_name(param0 string, param1 int) {}

func Fn_gtk_button_new_from_stock(param0 string) {}

func Fn_gtk_button_new_with_label(param0 string) {}

func Fn_gtk_button_new_with_mnemonic(param0 string) {}

func Fn_gtk_button_box_new(param0 int) {}

func Fn_gtk_calendar_new() {
	C.gtk_calendar_new()
}

func Fn_gtk_cell_area_box_new() {
	C.gtk_cell_area_box_new()
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

func Fn_gtk_cell_renderer_spinner_new() {
	C.gtk_cell_renderer_spinner_new()
}

func Fn_gtk_cell_renderer_text_new() {
	C.gtk_cell_renderer_text_new()
}

func Fn_gtk_cell_renderer_toggle_new() {
	C.gtk_cell_renderer_toggle_new()
}

func Fn_gtk_cell_view_new() {
	C.gtk_cell_view_new()
}

func Fn_gtk_cell_view_new_with_context(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_cell_view_new_with_markup(param0 string) {}

func Fn_gtk_cell_view_new_with_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_cell_view_new_with_text(param0 string) {}

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

func Fn_gtk_clipboard_get(param0 gdk.Atom) {}

func Fn_gtk_clipboard_get_for_display(param0 unsafe.Pointer, param1 gdk.Atom) {}

func Fn_gtk_color_button_new() {
	C.gtk_color_button_new()
}

func Fn_gtk_color_button_new_with_color(param0 unsafe.Pointer) {}

func Fn_gtk_color_button_new_with_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_color_chooser_dialog_new(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_color_chooser_widget_new() {
	C.gtk_color_chooser_widget_new()
}

func Fn_gtk_color_selection_new() {
	C.gtk_color_selection_new()
}

func Fn_gtk_color_selection_palette_from_string(param0 string, param1 []unsafe.Pointer, param2 *int) {
}

func Fn_gtk_color_selection_palette_to_string(param0 []gdk.Color, param1 int) {}

// UNSUPPORTED : set_change_palette_with_screen_hook : has callback

func Fn_gtk_color_selection_dialog_new(param0 string) {}

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

func Fn_gtk_combo_box_text_new() {
	C.gtk_combo_box_text_new()
}

func Fn_gtk_combo_box_text_new_with_entry() {
	C.gtk_combo_box_text_new_with_entry()
}

func Fn_gtk_container_cell_accessible_new() {
	C.gtk_container_cell_accessible_new()
}

func Fn_gtk_css_provider_new() {
	C.gtk_css_provider_new()
}

func Fn_gtk_css_provider_get_default() {
	C.gtk_css_provider_get_default()
}

func Fn_gtk_css_provider_get_named(param0 string, param1 string) {}

func Fn_gtk_dialog_new() {
	C.gtk_dialog_new()
}

// UNSUPPORTED : new_with_buttons : has varargs

func Fn_gtk_drawing_area_new() {
	C.gtk_drawing_area_new()
}

func Fn_gtk_entry_new() {
	C.gtk_entry_new()
}

func Fn_gtk_entry_new_with_buffer(param0 unsafe.Pointer) {}

func Fn_gtk_entry_buffer_new(param0 string, param1 int) {}

func Fn_gtk_entry_completion_new() {
	C.gtk_entry_completion_new()
}

func Fn_gtk_entry_completion_new_with_area(param0 unsafe.Pointer) {}

func Fn_gtk_event_box_new() {
	C.gtk_event_box_new()
}

func Fn_gtk_event_controller_key_new(param0 unsafe.Pointer) {}

func Fn_gtk_expander_new(param0 string) {}

func Fn_gtk_expander_new_with_mnemonic(param0 string) {}

func Fn_gtk_file_chooser_button_new(param0 string, param1 int) {}

func Fn_gtk_file_chooser_button_new_with_dialog(param0 unsafe.Pointer) {}

// UNSUPPORTED : new : has varargs

func Fn_gtk_file_chooser_widget_new(param0 int) {}

func Fn_gtk_file_filter_new() {
	C.gtk_file_filter_new()
}

func Fn_gtk_fixed_new() {
	C.gtk_fixed_new()
}

func Fn_gtk_flow_box_new() {
	C.gtk_flow_box_new()
}

func Fn_gtk_flow_box_child_new() {
	C.gtk_flow_box_child_new()
}

func Fn_gtk_font_button_new() {
	C.gtk_font_button_new()
}

func Fn_gtk_font_button_new_with_font(param0 string) {}

func Fn_gtk_font_chooser_dialog_new(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_font_chooser_widget_new() {
	C.gtk_font_chooser_widget_new()
}

func Fn_gtk_font_selection_new() {
	C.gtk_font_selection_new()
}

func Fn_gtk_font_selection_dialog_new(param0 string) {}

func Fn_gtk_frame_new(param0 string) {}

func Fn_gtk_gesture_drag_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_long_press_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_multi_press_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_pan_new(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_gesture_rotate_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_swipe_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_zoom_new(param0 unsafe.Pointer) {}

func Fn_gtk_grid_new() {
	C.gtk_grid_new()
}

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

func Fn_gtk_header_bar_new() {
	C.gtk_header_bar_new()
}

func Fn_gtk_im_context_simple_new() {
	C.gtk_im_context_simple_new()
}

func Fn_gtk_im_multicontext_new() {
	C.gtk_im_multicontext_new()
}

func Fn_gtk_icon_factory_new() {
	C.gtk_icon_factory_new()
}

func Fn_gtk_icon_factory_lookup_default(param0 string) {}

func Fn_gtk_icon_info_new_for_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_icon_theme_new() {
	C.gtk_icon_theme_new()
}

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

func Fn_gtk_image_menu_item_new() {
	C.gtk_image_menu_item_new()
}

func Fn_gtk_image_menu_item_new_from_stock(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_image_menu_item_new_with_label(param0 string) {}

func Fn_gtk_image_menu_item_new_with_mnemonic(param0 string) {}

func Fn_gtk_info_bar_new() {
	C.gtk_info_bar_new()
}

// UNSUPPORTED : new_with_buttons : has varargs

func Fn_gtk_invisible_new() {
	C.gtk_invisible_new()
}

func Fn_gtk_invisible_new_for_screen(param0 unsafe.Pointer) {}

func Fn_gtk_label_new(param0 string) {}

func Fn_gtk_label_new_with_mnemonic(param0 string) {}

func Fn_gtk_layout_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_level_bar_new() {
	C.gtk_level_bar_new()
}

func Fn_gtk_level_bar_new_for_interval(param0 float64, param1 float64) {}

func Fn_gtk_link_button_new(param0 string) {}

func Fn_gtk_link_button_new_with_label(param0 string, param1 string) {}

func Fn_gtk_list_box_new() {
	C.gtk_list_box_new()
}

func Fn_gtk_list_box_row_new() {
	C.gtk_list_box_row_new()
}

// UNSUPPORTED : new : has varargs

func Fn_gtk_list_store_newv(param0 int, param1 []uint64) {}

func Fn_gtk_lock_button_new(param0 unsafe.Pointer) {}

func Fn_gtk_menu_new() {
	C.gtk_menu_new()
}

func Fn_gtk_menu_new_from_model(param0 unsafe.Pointer) {}

func Fn_gtk_menu_get_for_attach_widget(param0 unsafe.Pointer) {}

func Fn_gtk_menu_bar_new() {
	C.gtk_menu_bar_new()
}

func Fn_gtk_menu_bar_new_from_model(param0 unsafe.Pointer) {}

func Fn_gtk_menu_button_new() {
	C.gtk_menu_button_new()
}

func Fn_gtk_menu_item_new() {
	C.gtk_menu_item_new()
}

func Fn_gtk_menu_item_new_with_label(param0 string) {}

func Fn_gtk_menu_item_new_with_mnemonic(param0 string) {}

func Fn_gtk_menu_tool_button_new(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_menu_tool_button_new_from_stock(param0 string) {}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_with_markup : has varargs

func Fn_gtk_mount_operation_new(param0 unsafe.Pointer) {}

func Fn_gtk_notebook_new() {
	C.gtk_notebook_new()
}

func Fn_gtk_notebook_page_accessible_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_numerable_icon_new(param0 unsafe.Pointer) {}

func Fn_gtk_numerable_icon_new_with_style_context(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_offscreen_window_new() {
	C.gtk_offscreen_window_new()
}

func Fn_gtk_overlay_new() {
	C.gtk_overlay_new()
}

func Fn_gtk_page_setup_new() {
	C.gtk_page_setup_new()
}

func Fn_gtk_page_setup_new_from_file(param0 string) {}

func Fn_gtk_page_setup_new_from_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_paned_new(param0 int) {}

func Fn_gtk_places_sidebar_new() {
	C.gtk_places_sidebar_new()
}

func Fn_gtk_plug_new(param0 uint64) {}

func Fn_gtk_plug_new_for_display(param0 unsafe.Pointer, param1 uint64) {}

func Fn_gtk_popover_new(param0 unsafe.Pointer) {}

func Fn_gtk_popover_new_from_model(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_print_operation_new() {
	C.gtk_print_operation_new()
}

func Fn_gtk_print_settings_new() {
	C.gtk_print_settings_new()
}

func Fn_gtk_print_settings_new_from_file(param0 string) {}

func Fn_gtk_print_settings_new_from_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_progress_bar_new() {
	C.gtk_progress_bar_new()
}

func Fn_gtk_radio_action_new(param0 string, param1 string, param2 string, param3 string, param4 int) {}

func Fn_gtk_radio_button_new(param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_new_from_widget(param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_new_with_label(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_new_with_mnemonic(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new(param0 unsafe.Pointer) {}

func Fn_gtk_radio_menu_item_new_from_widget(param0 unsafe.Pointer) {}

func Fn_gtk_radio_menu_item_new_with_label(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new_with_mnemonic(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_tool_button_new(param0 unsafe.Pointer) {}

func Fn_gtk_radio_tool_button_new_from_stock(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_tool_button_new_from_widget(param0 unsafe.Pointer) {}

func Fn_gtk_radio_tool_button_new_with_stock_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_rc_style_new() {
	C.gtk_rc_style_new()
}

func Fn_gtk_recent_action_new(param0 string, param1 string, param2 string, param3 string) {}

func Fn_gtk_recent_action_new_for_manager(param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer) {
}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_for_manager : has varargs

func Fn_gtk_recent_chooser_menu_new() {
	C.gtk_recent_chooser_menu_new()
}

func Fn_gtk_recent_chooser_menu_new_for_manager(param0 unsafe.Pointer) {}

func Fn_gtk_recent_chooser_widget_new() {
	C.gtk_recent_chooser_widget_new()
}

func Fn_gtk_recent_chooser_widget_new_for_manager(param0 unsafe.Pointer) {}

func Fn_gtk_recent_filter_new() {
	C.gtk_recent_filter_new()
}

func Fn_gtk_recent_manager_new() {
	C.gtk_recent_manager_new()
}

func Fn_gtk_recent_manager_get_default() {
	C.gtk_recent_manager_get_default()
}

func Fn_gtk_renderer_cell_accessible_new(param0 unsafe.Pointer) {}

func Fn_gtk_revealer_new() {
	C.gtk_revealer_new()
}

func Fn_gtk_scale_new(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_scale_new_with_range(param0 int, param1 float64, param2 float64, param3 float64) {}

func Fn_gtk_scale_button_new(param0 int, param1 float64, param2 float64, param3 float64, param4 []string) {
}

func Fn_gtk_scrollbar_new(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_search_bar_new() {
	C.gtk_search_bar_new()
}

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

func Fn_gtk_settings_get_default() {
	C.gtk_settings_get_default()
}

func Fn_gtk_settings_get_for_screen(param0 unsafe.Pointer) {}

func Fn_gtk_settings_install_property(param0 unsafe.Pointer) {}

// UNSUPPORTED : install_property_parser : has callback

func Fn_gtk_size_group_new(param0 int) {}

func Fn_gtk_socket_new() {
	C.gtk_socket_new()
}

func Fn_gtk_spin_button_new(param0 unsafe.Pointer, param1 float64, param2 uint) {}

func Fn_gtk_spin_button_new_with_range(param0 float64, param1 float64, param2 float64) {}

func Fn_gtk_spinner_new() {
	C.gtk_spinner_new()
}

func Fn_gtk_stack_new() {
	C.gtk_stack_new()
}

func Fn_gtk_stack_switcher_new() {
	C.gtk_stack_switcher_new()
}

func Fn_gtk_status_icon_new() {
	C.gtk_status_icon_new()
}

func Fn_gtk_status_icon_new_from_file(param0 string) {}

func Fn_gtk_status_icon_new_from_gicon(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_new_from_icon_name(param0 string) {}

func Fn_gtk_status_icon_new_from_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_new_from_stock(param0 string) {}

func Fn_gtk_status_icon_position_menu(param0 unsafe.Pointer, param1 *int, param2 *int, param3 *bool, param4 unsafe.Pointer) {
}

func Fn_gtk_statusbar_new() {
	C.gtk_statusbar_new()
}

func Fn_gtk_style_new() {
	C.gtk_style_new()
}

func Fn_gtk_style_context_new() {
	C.gtk_style_context_new()
}

func Fn_gtk_style_context_add_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint) {
}

func Fn_gtk_style_context_remove_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_reset_widgets(param0 unsafe.Pointer) {}

func Fn_gtk_style_properties_new() {
	C.gtk_style_properties_new()
}

// UNSUPPORTED : lookup_property : has callback

// UNSUPPORTED : register_property : has callback

func Fn_gtk_switch_new() {
	C.gtk_switch_new()
}

func Fn_gtk_table_new(param0 uint, param1 uint, param2 bool) {}

func Fn_gtk_tearoff_menu_item_new() {
	C.gtk_tearoff_menu_item_new()
}

func Fn_gtk_text_buffer_new(param0 unsafe.Pointer) {}

func Fn_gtk_text_child_anchor_new() {
	C.gtk_text_child_anchor_new()
}

func Fn_gtk_text_mark_new(param0 string, param1 bool) {}

func Fn_gtk_text_tag_new(param0 string) {}

func Fn_gtk_text_tag_table_new() {
	C.gtk_text_tag_table_new()
}

func Fn_gtk_text_view_new() {
	C.gtk_text_view_new()
}

func Fn_gtk_text_view_new_with_buffer(param0 unsafe.Pointer) {}

func Fn_gtk_theming_engine_load(param0 string) {}

// UNSUPPORTED : register_property : has callback

func Fn_gtk_toggle_action_new(param0 string, param1 string, param2 string, param3 string) {}

func Fn_gtk_toggle_button_new() {
	C.gtk_toggle_button_new()
}

func Fn_gtk_toggle_button_new_with_label(param0 string) {}

func Fn_gtk_toggle_button_new_with_mnemonic(param0 string) {}

func Fn_gtk_toggle_tool_button_new() {
	C.gtk_toggle_tool_button_new()
}

func Fn_gtk_toggle_tool_button_new_from_stock(param0 string) {}

func Fn_gtk_tool_button_new(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_tool_button_new_from_stock(param0 string) {}

func Fn_gtk_tool_item_new() {
	C.gtk_tool_item_new()
}

func Fn_gtk_tool_item_group_new(param0 string) {}

func Fn_gtk_tool_palette_new() {
	C.gtk_tool_palette_new()
}

func Fn_gtk_tool_palette_get_drag_target_group() {
	C.gtk_tool_palette_get_drag_target_group()
}

func Fn_gtk_tool_palette_get_drag_target_item() {
	C.gtk_tool_palette_get_drag_target_item()
}

func Fn_gtk_toolbar_new() {
	C.gtk_toolbar_new()
}

func Fn_gtk_tooltip_trigger_tooltip_query(param0 unsafe.Pointer) {}

// UNSUPPORTED : new : has varargs

func Fn_gtk_tree_store_newv(param0 int, param1 []uint64) {}

func Fn_gtk_tree_view_new() {
	C.gtk_tree_view_new()
}

func Fn_gtk_tree_view_new_with_model(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_column_new() {
	C.gtk_tree_view_column_new()
}

func Fn_gtk_tree_view_column_new_with_area(param0 unsafe.Pointer) {}

// UNSUPPORTED : new_with_attributes : has varargs

func Fn_gtk_ui_manager_new() {
	C.gtk_ui_manager_new()
}

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

func Fn_gtk_volume_button_new() {
	C.gtk_volume_button_new()
}

// UNSUPPORTED : new : has varargs

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

func Fn_gtk_window_set_interactive_debugging(param0 bool) {}

func Fn_gtk_window_group_new() {
	C.gtk_window_group_new()
}
