// Code generated - DO NOT EDIT.
// +build gtk_2.14

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

// records
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

func Fn_accel_groups_activate(object unsafe.Pointer, accelKey uint, accelMods int) {}

func Fn_accel_groups_from_object(object unsafe.Pointer) {}

func Fn_accelerator_get_default_mod_mask() {
	C.gtk_accelerator_get_default_mod_mask()
}

func Fn_accelerator_get_label(acceleratorKey uint, acceleratorMods int) {}

func Fn_accelerator_name(acceleratorKey uint, acceleratorMods int) {}

func Fn_accelerator_parse(accelerator string, acceleratorKey *uint, acceleratorMods int) {}

func Fn_accelerator_set_default_mod_mask(defaultModMask int) {}

func Fn_accelerator_valid(keyval uint, modifiers int) {}

func Fn_alternative_dialog_button_order(screen unsafe.Pointer) {}

func Fn_binding_entry_add_signall(bindingSet unsafe.Pointer, keyval uint, modifiers int, signalName string, bindingArgs unsafe.Pointer) {
}

func Fn_binding_entry_remove(bindingSet unsafe.Pointer, keyval uint, modifiers int) {}

func Fn_binding_entry_skip(bindingSet unsafe.Pointer, keyval uint, modifiers int) {}

func Fn_binding_set_by_class(objectClass unsafe.Pointer) {}

func Fn_binding_set_find(setName string) {}

func Fn_binding_set_new(setName string) {}

func Fn_bindings_activate(object unsafe.Pointer, keyval uint, modifiers int) {}

func Fn_bindings_activate_event(object unsafe.Pointer, event unsafe.Pointer) {}

func Fn_builder_error_quark() {
	C.gtk_builder_error_quark()
}

func Fn_check_version(requiredMajor uint, requiredMinor uint, requiredMicro uint) {}

func Fn_css_provider_error_quark() {
	C.gtk_css_provider_error_quark()
}

func Fn_disable_setlocale() {
	C.gtk_disable_setlocale()
}

func Fn_distribute_natural_allocation(extraSpace int, nRequestedSizes uint, sizes unsafe.Pointer) {}

func Fn_drag_finish(context unsafe.Pointer, success bool, del bool, time uint32) {}

func Fn_drag_get_source_widget(context unsafe.Pointer) {}

func Fn_drag_set_icon_default(context unsafe.Pointer) {}

func Fn_drag_set_icon_name(context unsafe.Pointer, iconName string, hotX int, hotY int) {}

func Fn_drag_set_icon_pixbuf(context unsafe.Pointer, pixbuf unsafe.Pointer, hotX int, hotY int) {}

func Fn_drag_set_icon_stock(context unsafe.Pointer, stockId string, hotX int, hotY int) {}

func Fn_drag_set_icon_surface(context unsafe.Pointer, surface unsafe.Pointer) {}

func Fn_drag_set_icon_widget(context unsafe.Pointer, widget unsafe.Pointer, hotX int, hotY int) {}

func Fn_events_pending() {
	C.gtk_events_pending()
}

func Fn_false() {
	C.gtk_false()
}

func Fn_file_chooser_error_quark() {
	C.gtk_file_chooser_error_quark()
}

func Fn_get_current_event() {
	C.gtk_get_current_event()
}

func Fn_get_current_event_device() {
	C.gtk_get_current_event_device()
}

func Fn_get_current_event_state(state int) {}

func Fn_get_current_event_time() {
	C.gtk_get_current_event_time()
}

func Fn_get_debug_flags() {
	C.gtk_get_debug_flags()
}

func Fn_get_default_language() {
	C.gtk_get_default_language()
}

func Fn_get_event_widget(event unsafe.Pointer) {}

func Fn_get_option_group(openDefaultDisplay bool) {}

func Fn_grab_get_current() {
	C.gtk_grab_get_current()
}

func Fn_icon_size_from_name(name string) {}

func Fn_icon_size_get_name(size int) {}

func Fn_icon_size_lookup(size int, width *int, height *int) {}

func Fn_icon_size_lookup_for_settings(settings unsafe.Pointer, size int, width *int, height *int) {}

func Fn_icon_size_register(name string, width int, height int) {}

func Fn_icon_size_register_alias(alias string, target int) {}

func Fn_icon_theme_error_quark() {
	C.gtk_icon_theme_error_quark()
}

func Fn_init(argc *int, argv *[]string) {}

func Fn_init_check(argc *int, argv *[]string) {}

func Fn_init_with_args(argc *int, argv *[]string, parameterString string, entries []glib.OptionEntry, translationDomain string) {
}

// UNSUPPORTED : key_snooper_install : has callback

func Fn_key_snooper_remove(snooperHandlerId uint) {}

func Fn_main() {
	C.gtk_main()
}

func Fn_main_do_event(event unsafe.Pointer) {}

func Fn_main_iteration() {
	C.gtk_main_iteration()
}

func Fn_main_iteration_do(blocking bool) {}

func Fn_main_level() {
	C.gtk_main_level()
}

func Fn_main_quit() {
	C.gtk_main_quit()
}

func Fn_paint_arrow(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, arrowType int, fill bool, x int, y int, width int, height int) {
}

func Fn_paint_box(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int) {
}

func Fn_paint_box_gap(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int, gapSide int, gapX int, gapWidth int) {
}

func Fn_paint_check(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int) {
}

func Fn_paint_diamond(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int) {
}

func Fn_paint_expander(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail string, x int, y int, expanderStyle int) {
}

func Fn_paint_extension(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int, gapSide int) {
}

func Fn_paint_flat_box(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int) {
}

func Fn_paint_focus(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int) {
}

func Fn_paint_handle(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int, orientation int) {
}

func Fn_paint_hline(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail string, x1 int, x2 int, y int) {
}

func Fn_paint_layout(style unsafe.Pointer, cr unsafe.Pointer, stateType int, useText bool, widget unsafe.Pointer, detail string, x int, y int, layout unsafe.Pointer) {
}

func Fn_paint_option(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int) {
}

func Fn_paint_resize_grip(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail string, edge int, x int, y int, width int, height int) {
}

func Fn_paint_shadow(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int) {
}

func Fn_paint_shadow_gap(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int, gapSide int, gapX int, gapWidth int) {
}

func Fn_paint_slider(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int, orientation int) {
}

func Fn_paint_spinner(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail string, step uint, x int, y int, width int, height int) {
}

func Fn_paint_tab(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail string, x int, y int, width int, height int) {
}

func Fn_paint_vline(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail string, y1 int, y2 int, x int) {
}

func Fn_paper_size_get_default() {
	C.gtk_paper_size_get_default()
}

func Fn_paper_size_get_paper_sizes(includeCustom bool) {}

func Fn_parse_args(argc *int, argv *[]string) {}

func Fn_print_error_quark() {
	C.gtk_print_error_quark()
}

func Fn_print_run_page_setup_dialog(parent unsafe.Pointer, pageSetup unsafe.Pointer, settings unsafe.Pointer) {
}

// UNSUPPORTED : print_run_page_setup_dialog_async : has callback

func Fn_propagate_event(widget unsafe.Pointer, event unsafe.Pointer) {}

func Fn_rc_add_default_file(filename string) {}

func Fn_rc_find_module_in_path(moduleFile string) {}

func Fn_rc_find_pixmap_in_path(settings unsafe.Pointer, scanner unsafe.Pointer, pixmapFile string) {}

func Fn_rc_get_default_files() {
	C.gtk_rc_get_default_files()
}

func Fn_rc_get_im_module_file() {
	C.gtk_rc_get_im_module_file()
}

func Fn_rc_get_im_module_path() {
	C.gtk_rc_get_im_module_path()
}

func Fn_rc_get_module_dir() {
	C.gtk_rc_get_module_dir()
}

func Fn_rc_get_style(widget unsafe.Pointer) {}

func Fn_rc_get_style_by_paths(settings unsafe.Pointer, widgetPath string, classPath string, type_ uint64) {
}

func Fn_rc_get_theme_dir() {
	C.gtk_rc_get_theme_dir()
}

func Fn_rc_parse(filename string) {}

func Fn_rc_parse_color(scanner unsafe.Pointer, color unsafe.Pointer) {}

func Fn_rc_parse_color_full(scanner unsafe.Pointer, style unsafe.Pointer, color unsafe.Pointer) {}

func Fn_rc_parse_priority(scanner unsafe.Pointer, priority int) {}

func Fn_rc_parse_state(scanner unsafe.Pointer, state int) {}

func Fn_rc_parse_string(rcString string) {}

func Fn_rc_property_parse_border(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) {
}

func Fn_rc_property_parse_color(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) {
}

func Fn_rc_property_parse_enum(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) {
}

func Fn_rc_property_parse_flags(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) {
}

func Fn_rc_property_parse_requisition(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) {
}

func Fn_rc_reparse_all() {
	C.gtk_rc_reparse_all()
}

func Fn_rc_reparse_all_for_settings(settings unsafe.Pointer, forceLoad bool) {}

func Fn_rc_reset_styles(settings unsafe.Pointer) {}

func Fn_rc_scanner_new() {
	C.gtk_rc_scanner_new()
}

func Fn_rc_set_default_files(filenames []string) {}

func Fn_recent_chooser_error_quark() {
	C.gtk_recent_chooser_error_quark()
}

func Fn_recent_manager_error_quark() {
	C.gtk_recent_manager_error_quark()
}

func Fn_rgb_to_hsv(r float64, g float64, b float64, h *float64, s *float64, v *float64) {}

func Fn_selection_add_target(widget unsafe.Pointer, selection gdk.Atom, target gdk.Atom, info uint) {}

func Fn_selection_add_targets(widget unsafe.Pointer, selection gdk.Atom, targets []TargetEntry, ntargets uint) {
}

func Fn_selection_clear_targets(widget unsafe.Pointer, selection gdk.Atom) {}

func Fn_selection_convert(widget unsafe.Pointer, selection gdk.Atom, target gdk.Atom, time uint32) {}

func Fn_selection_owner_set(widget unsafe.Pointer, selection gdk.Atom, time uint32) {}

func Fn_selection_owner_set_for_display(display unsafe.Pointer, widget unsafe.Pointer, selection gdk.Atom, time uint32) {
}

func Fn_selection_remove_all(widget unsafe.Pointer) {}

func Fn_set_debug_flags(flags uint) {}

// UNSUPPORTED : show_about_dialog : has varargs

func Fn_show_uri(screen unsafe.Pointer, uri string, timestamp uint32) {}

func Fn_stock_add(items []StockItem, nItems uint) {}

func Fn_stock_add_static(items []StockItem, nItems uint) {}

func Fn_stock_list_ids() {
	C.gtk_stock_list_ids()
}

func Fn_stock_lookup(stockId string, item unsafe.Pointer) {}

// UNSUPPORTED : stock_set_translate_func : has callback

func Fn_target_table_free(targets []TargetEntry, nTargets int) {}

func Fn_target_table_new_from_list(list unsafe.Pointer, nTargets *int) {}

func Fn_targets_include_image(targets []gdk.Atom, nTargets int, writable bool) {}

func Fn_targets_include_rich_text(targets []gdk.Atom, nTargets int, buffer unsafe.Pointer) {}

func Fn_targets_include_text(targets []gdk.Atom, nTargets int) {}

func Fn_targets_include_uri(targets []gdk.Atom, nTargets int) {}

func Fn_test_create_simple_window(windowTitle string, dialogText string) {}

// UNSUPPORTED : test_create_widget : has varargs

// UNSUPPORTED : test_display_button_window : has varargs

func Fn_test_find_label(widget unsafe.Pointer, labelPattern string) {}

func Fn_test_find_sibling(baseWidget unsafe.Pointer, widgetType uint64) {}

func Fn_test_find_widget(widget unsafe.Pointer, labelPattern string, widgetType uint64) {}

// UNSUPPORTED : test_init : has varargs

func Fn_test_list_all_types(nTypes *uint) {}

func Fn_test_register_all_types() {
	C.gtk_test_register_all_types()
}

func Fn_test_slider_get_value(widget unsafe.Pointer) {}

func Fn_test_slider_set_perc(widget unsafe.Pointer, percentage float64) {}

func Fn_test_spin_button_click(spinner unsafe.Pointer, button uint, upwards bool) {}

func Fn_test_text_get(widget unsafe.Pointer) {}

func Fn_test_text_set(widget unsafe.Pointer, string_ string) {}

func Fn_test_widget_click(widget unsafe.Pointer, button uint, modifiers int) {}

func Fn_test_widget_send_key(widget unsafe.Pointer, keyval uint, modifiers int) {}

func Fn_tree_get_row_drag_data(selectionData unsafe.Pointer, treeModel *unsafe.Pointer, path *unsafe.Pointer) {
}

func Fn_tree_row_reference_deleted(proxy unsafe.Pointer, path unsafe.Pointer) {}

func Fn_tree_row_reference_inserted(proxy unsafe.Pointer, path unsafe.Pointer) {}

func Fn_tree_row_reference_reordered(proxy unsafe.Pointer, path unsafe.Pointer, iter unsafe.Pointer, newOrder []int) {
}

func Fn_tree_set_row_drag_data(selectionData unsafe.Pointer, treeModel unsafe.Pointer, path unsafe.Pointer) {
}

func Fn_true() {
	C.gtk_true()
}
