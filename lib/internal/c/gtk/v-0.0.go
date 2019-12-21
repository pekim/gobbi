// Code generated - DO NOT EDIT.
// +build !gtk_2.2,!gtk_2.4,!gtk_2.6,!gtk_2.8,!gtk_2.10,!gtk_2.12,!gtk_2.14,!gtk_2.16,!gtk_2.18,!gtk_2.20,!gtk_3.0,!gtk_3.2,!gtk_3.4,!gtk_3.6,!gtk_3.10,!gtk_3.12,!gtk_3.14,!gtk_3.16,!gtk_3.20,!gtk_3.24

package gtk

import (
	c "github.com/pekim/gobbi/lib/internal/c"
	cairo "github.com/pekim/gobbi/lib/internal/c/cairo"
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/internal/c/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	pango "github.com/pekim/gobbi/lib/internal/c/pango"
	"unsafe"
)

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// bitfields
type AccelFlags C.GtkAccelFlags
type AttachOptions C.GtkAttachOptions
type CalendarDisplayOptions C.GtkCalendarDisplayOptions
type CellRendererState C.GtkCellRendererState
type DebugFlag C.GtkDebugFlag
type DestDefaults C.GtkDestDefaults
type DialogFlags C.GtkDialogFlags
type FileFilterFlags C.GtkFileFilterFlags
type IconLookupFlags C.GtkIconLookupFlags
type JunctionSides C.GtkJunctionSides
type PlacesOpenFlags C.GtkPlacesOpenFlags
type RcFlags C.GtkRcFlags
type RecentFilterFlags C.GtkRecentFilterFlags
type RegionFlags C.GtkRegionFlags
type StateFlags C.GtkStateFlags
type StyleContextPrintFlags C.GtkStyleContextPrintFlags
type TargetFlags C.GtkTargetFlags
type TextSearchFlags C.GtkTextSearchFlags
type ToolPaletteDragTargets C.GtkToolPaletteDragTargets
type TreeModelFlags C.GtkTreeModelFlags
type UIManagerItemType C.GtkUIManagerItemType

// enumerations
type Align C.GtkAlign
type ArrowPlacement C.GtkArrowPlacement
type ArrowType C.GtkArrowType
type AssistantPageType C.GtkAssistantPageType
type BorderStyle C.GtkBorderStyle
type BuilderError C.GtkBuilderError
type ButtonBoxStyle C.GtkButtonBoxStyle
type ButtonRole C.GtkButtonRole
type ButtonsType C.GtkButtonsType
type CellRendererAccelMode C.GtkCellRendererAccelMode
type CellRendererMode C.GtkCellRendererMode
type CornerType C.GtkCornerType
type CssProviderError C.GtkCssProviderError
type DeleteType C.GtkDeleteType
type DirectionType C.GtkDirectionType
type DragResult C.GtkDragResult
type ExpanderStyle C.GtkExpanderStyle
type FileChooserAction C.GtkFileChooserAction
type FileChooserError C.GtkFileChooserError
type IMPreeditStyle C.GtkIMPreeditStyle
type IMStatusStyle C.GtkIMStatusStyle
type IconSize C.GtkIconSize
type IconThemeError C.GtkIconThemeError
type IconViewDropPosition C.GtkIconViewDropPosition
type ImageType C.GtkImageType
type Justification C.GtkJustification
type MenuDirectionType C.GtkMenuDirectionType
type MessageType C.GtkMessageType
type MovementStep C.GtkMovementStep
type NotebookTab C.GtkNotebookTab
type NumberUpLayout C.GtkNumberUpLayout
type Orientation C.GtkOrientation
type PackDirection C.GtkPackDirection
type PackType C.GtkPackType
type PadActionType C.GtkPadActionType
type PageOrientation C.GtkPageOrientation
type PageSet C.GtkPageSet
type PathPriorityType C.GtkPathPriorityType
type PathType C.GtkPathType
type PolicyType C.GtkPolicyType
type PositionType C.GtkPositionType
type PrintDuplex C.GtkPrintDuplex
type PrintError C.GtkPrintError
type PrintOperationAction C.GtkPrintOperationAction
type PrintOperationResult C.GtkPrintOperationResult
type PrintPages C.GtkPrintPages
type PrintQuality C.GtkPrintQuality
type PrintStatus C.GtkPrintStatus
type RcTokenType C.GtkRcTokenType
type ReliefStyle C.GtkReliefStyle
type ResizeMode C.GtkResizeMode
type ResponseType C.GtkResponseType
type RevealerTransitionType C.GtkRevealerTransitionType
type ScrollStep C.GtkScrollStep
type ScrollType C.GtkScrollType
type ScrollablePolicy C.GtkScrollablePolicy
type SelectionMode C.GtkSelectionMode
type SensitivityType C.GtkSensitivityType
type ShadowType C.GtkShadowType
type SizeGroupMode C.GtkSizeGroupMode
type SizeRequestMode C.GtkSizeRequestMode
type SortType C.GtkSortType
type SpinButtonUpdatePolicy C.GtkSpinButtonUpdatePolicy
type SpinType C.GtkSpinType
type StackTransitionType C.GtkStackTransitionType
type StateType C.GtkStateType
type TextBufferTargetInfo C.GtkTextBufferTargetInfo
type TextDirection C.GtkTextDirection
type TextViewLayer C.GtkTextViewLayer
type TextWindowType C.GtkTextWindowType
type ToolbarSpaceStyle C.GtkToolbarSpaceStyle
type ToolbarStyle C.GtkToolbarStyle
type TreeViewColumnSizing C.GtkTreeViewColumnSizing
type TreeViewDropPosition C.GtkTreeViewDropPosition
type TreeViewGridLines C.GtkTreeViewGridLines
type Unit C.GtkUnit
type WidgetHelpType C.GtkWidgetHelpType
type WindowPosition C.GtkWindowPosition
type WindowType C.GtkWindowType
type WrapMode C.GtkWrapMode

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

// classes
type AboutDialog C.GtkAboutDialog
type AccelGroup C.GtkAccelGroup
type AccelLabel C.GtkAccelLabel
type AccelMap C.GtkAccelMap
type Accessible C.GtkAccessible
type Action C.GtkAction
type ActionBar C.GtkActionBar
type ActionGroup C.GtkActionGroup
type Adjustment C.GtkAdjustment
type Alignment C.GtkAlignment
type AppChooserButton C.GtkAppChooserButton
type AppChooserDialog C.GtkAppChooserDialog
type AppChooserWidget C.GtkAppChooserWidget
type Application C.GtkApplication
type ApplicationWindow C.GtkApplicationWindow
type Arrow C.GtkArrow
type ArrowAccessible C.GtkArrowAccessible
type AspectFrame C.GtkAspectFrame
type Assistant C.GtkAssistant
type Bin C.GtkBin
type BooleanCellAccessible C.GtkBooleanCellAccessible
type Box C.GtkBox
type Builder C.GtkBuilder
type Button C.GtkButton
type ButtonAccessible C.GtkButtonAccessible
type ButtonBox C.GtkButtonBox
type Calendar C.GtkCalendar
type CellAccessible C.GtkCellAccessible
type CellArea C.GtkCellArea
type CellAreaBox C.GtkCellAreaBox
type CellAreaContext C.GtkCellAreaContext
type CellRenderer C.GtkCellRenderer
type CellRendererAccel C.GtkCellRendererAccel
type CellRendererCombo C.GtkCellRendererCombo
type CellRendererPixbuf C.GtkCellRendererPixbuf
type CellRendererProgress C.GtkCellRendererProgress
type CellRendererSpin C.GtkCellRendererSpin
type CellRendererSpinner C.GtkCellRendererSpinner
type CellRendererText C.GtkCellRendererText
type CellRendererToggle C.GtkCellRendererToggle
type CellView C.GtkCellView
type CheckButton C.GtkCheckButton
type CheckMenuItem C.GtkCheckMenuItem
type CheckMenuItemAccessible C.GtkCheckMenuItemAccessible
type Clipboard C.GtkClipboard
type ColorButton C.GtkColorButton
type ColorSelection C.GtkColorSelection
type ColorSelectionDialog C.GtkColorSelectionDialog
type ComboBox C.GtkComboBox
type ComboBoxAccessible C.GtkComboBoxAccessible
type ComboBoxText C.GtkComboBoxText
type Container C.GtkContainer
type ContainerAccessible C.GtkContainerAccessible
type ContainerCellAccessible C.GtkContainerCellAccessible
type CssProvider C.GtkCssProvider
type Dialog C.GtkDialog
type DrawingArea C.GtkDrawingArea
type Entry C.GtkEntry
type EntryAccessible C.GtkEntryAccessible
type EntryCompletion C.GtkEntryCompletion

// UNSUPPORTED : EntryIconAccessible : blacklisted
type EventBox C.GtkEventBox
type EventController C.GtkEventController

// UNSUPPORTED : EventControllerKey : blacklisted
// UNSUPPORTED : EventControllerMotion : blacklisted
// UNSUPPORTED : EventControllerScroll : blacklisted
type Expander C.GtkExpander
type ExpanderAccessible C.GtkExpanderAccessible
type FileChooserButton C.GtkFileChooserButton
type FileChooserDialog C.GtkFileChooserDialog
type FileChooserNative C.GtkFileChooserNative
type FileChooserWidget C.GtkFileChooserWidget
type FileFilter C.GtkFileFilter
type Fixed C.GtkFixed
type FlowBox C.GtkFlowBox
type FlowBoxAccessible C.GtkFlowBoxAccessible
type FlowBoxChild C.GtkFlowBoxChild
type FlowBoxChildAccessible C.GtkFlowBoxChildAccessible
type FontButton C.GtkFontButton
type FontSelection C.GtkFontSelection
type FontSelectionDialog C.GtkFontSelectionDialog
type Frame C.GtkFrame
type FrameAccessible C.GtkFrameAccessible
type Gesture C.GtkGesture
type GestureDrag C.GtkGestureDrag
type GestureLongPress C.GtkGestureLongPress
type GestureMultiPress C.GtkGestureMultiPress
type GesturePan C.GtkGesturePan
type GestureRotate C.GtkGestureRotate
type GestureSingle C.GtkGestureSingle

// UNSUPPORTED : GestureStylus : blacklisted
type GestureSwipe C.GtkGestureSwipe
type GestureZoom C.GtkGestureZoom
type Grid C.GtkGrid
type HBox C.GtkHBox
type HButtonBox C.GtkHButtonBox
type HPaned C.GtkHPaned
type HSV C.GtkHSV
type HScale C.GtkHScale
type HScrollbar C.GtkHScrollbar
type HSeparator C.GtkHSeparator
type HandleBox C.GtkHandleBox
type HeaderBar C.GtkHeaderBar

// UNSUPPORTED : HeaderBarAccessible : blacklisted
type IMContext C.GtkIMContext
type IMContextSimple C.GtkIMContextSimple
type IMMulticontext C.GtkIMMulticontext
type IconFactory C.GtkIconFactory
type IconInfo C.GtkIconInfo
type IconTheme C.GtkIconTheme
type IconView C.GtkIconView
type IconViewAccessible C.GtkIconViewAccessible
type Image C.GtkImage
type ImageAccessible C.GtkImageAccessible
type ImageCellAccessible C.GtkImageCellAccessible
type ImageMenuItem C.GtkImageMenuItem
type InfoBar C.GtkInfoBar
type Invisible C.GtkInvisible
type Label C.GtkLabel
type LabelAccessible C.GtkLabelAccessible
type Layout C.GtkLayout
type LevelBar C.GtkLevelBar
type LevelBarAccessible C.GtkLevelBarAccessible
type LinkButton C.GtkLinkButton
type LinkButtonAccessible C.GtkLinkButtonAccessible
type ListBox C.GtkListBox
type ListBoxAccessible C.GtkListBoxAccessible
type ListBoxRow C.GtkListBoxRow
type ListBoxRowAccessible C.GtkListBoxRowAccessible
type ListStore C.GtkListStore
type LockButton C.GtkLockButton
type LockButtonAccessible C.GtkLockButtonAccessible
type Menu C.GtkMenu
type MenuAccessible C.GtkMenuAccessible
type MenuBar C.GtkMenuBar
type MenuButton C.GtkMenuButton
type MenuButtonAccessible C.GtkMenuButtonAccessible
type MenuItem C.GtkMenuItem
type MenuItemAccessible C.GtkMenuItemAccessible
type MenuShell C.GtkMenuShell
type MenuShellAccessible C.GtkMenuShellAccessible
type MenuToolButton C.GtkMenuToolButton
type MessageDialog C.GtkMessageDialog
type Misc C.GtkMisc
type ModelButton C.GtkModelButton
type MountOperation C.GtkMountOperation
type NativeDialog C.GtkNativeDialog
type Notebook C.GtkNotebook
type NotebookAccessible C.GtkNotebookAccessible
type NotebookPageAccessible C.GtkNotebookPageAccessible
type NumerableIcon C.GtkNumerableIcon
type OffscreenWindow C.GtkOffscreenWindow
type Overlay C.GtkOverlay
type PadController C.GtkPadController
type PageSetup C.GtkPageSetup
type Paned C.GtkPaned
type PanedAccessible C.GtkPanedAccessible
type PlacesSidebar C.GtkPlacesSidebar
type Plug C.GtkPlug
type PopoverAccessible C.GtkPopoverAccessible
type PopoverMenu C.GtkPopoverMenu
type PrintContext C.GtkPrintContext
type PrintOperation C.GtkPrintOperation
type PrintSettings C.GtkPrintSettings
type ProgressBar C.GtkProgressBar
type ProgressBarAccessible C.GtkProgressBarAccessible
type RadioAction C.GtkRadioAction
type RadioButton C.GtkRadioButton
type RadioButtonAccessible C.GtkRadioButtonAccessible
type RadioMenuItem C.GtkRadioMenuItem
type RadioMenuItemAccessible C.GtkRadioMenuItemAccessible
type RadioToolButton C.GtkRadioToolButton
type Range C.GtkRange
type RangeAccessible C.GtkRangeAccessible
type RcStyle C.GtkRcStyle
type RecentAction C.GtkRecentAction
type RecentChooserDialog C.GtkRecentChooserDialog
type RecentChooserMenu C.GtkRecentChooserMenu
type RecentChooserWidget C.GtkRecentChooserWidget
type RecentFilter C.GtkRecentFilter
type RendererCellAccessible C.GtkRendererCellAccessible
type Revealer C.GtkRevealer
type Scale C.GtkScale
type ScaleAccessible C.GtkScaleAccessible
type ScaleButton C.GtkScaleButton
type ScaleButtonAccessible C.GtkScaleButtonAccessible
type Scrollbar C.GtkScrollbar
type ScrolledWindow C.GtkScrolledWindow
type ScrolledWindowAccessible C.GtkScrolledWindowAccessible
type Separator C.GtkSeparator
type SeparatorMenuItem C.GtkSeparatorMenuItem
type SeparatorToolItem C.GtkSeparatorToolItem
type Settings C.GtkSettings
type ShortcutLabel C.GtkShortcutLabel
type ShortcutsGroup C.GtkShortcutsGroup
type ShortcutsSection C.GtkShortcutsSection
type ShortcutsShortcut C.GtkShortcutsShortcut
type ShortcutsWindow C.GtkShortcutsWindow
type SizeGroup C.GtkSizeGroup
type Socket C.GtkSocket
type SpinButton C.GtkSpinButton
type SpinButtonAccessible C.GtkSpinButtonAccessible
type Spinner C.GtkSpinner
type SpinnerAccessible C.GtkSpinnerAccessible
type Stack C.GtkStack

// UNSUPPORTED : StackAccessible : blacklisted
type StackSwitcher C.GtkStackSwitcher
type StatusIcon C.GtkStatusIcon
type Statusbar C.GtkStatusbar
type StatusbarAccessible C.GtkStatusbarAccessible
type Style C.GtkStyle
type StyleContext C.GtkStyleContext
type StyleProperties C.GtkStyleProperties
type Switch C.GtkSwitch
type SwitchAccessible C.GtkSwitchAccessible
type Table C.GtkTable
type TearoffMenuItem C.GtkTearoffMenuItem
type TextBuffer C.GtkTextBuffer
type TextCellAccessible C.GtkTextCellAccessible
type TextChildAnchor C.GtkTextChildAnchor
type TextMark C.GtkTextMark
type TextTag C.GtkTextTag
type TextTagTable C.GtkTextTagTable
type TextView C.GtkTextView
type TextViewAccessible C.GtkTextViewAccessible
type ThemingEngine C.GtkThemingEngine
type ToggleAction C.GtkToggleAction
type ToggleButton C.GtkToggleButton
type ToggleButtonAccessible C.GtkToggleButtonAccessible
type ToggleToolButton C.GtkToggleToolButton
type ToolButton C.GtkToolButton
type ToolItem C.GtkToolItem
type Toolbar C.GtkToolbar
type Tooltip C.GtkTooltip
type ToplevelAccessible C.GtkToplevelAccessible
type TreeModelFilter C.GtkTreeModelFilter
type TreeModelSort C.GtkTreeModelSort
type TreeSelection C.GtkTreeSelection
type TreeStore C.GtkTreeStore
type TreeView C.GtkTreeView
type TreeViewAccessible C.GtkTreeViewAccessible
type TreeViewColumn C.GtkTreeViewColumn
type UIManager C.GtkUIManager
type VBox C.GtkVBox
type VButtonBox C.GtkVButtonBox
type VPaned C.GtkVPaned
type VScale C.GtkVScale
type VScrollbar C.GtkVScrollbar
type VSeparator C.GtkVSeparator
type Viewport C.GtkViewport
type VolumeButton C.GtkVolumeButton
type Widget C.GtkWidget
type WidgetAccessible C.GtkWidgetAccessible
type Window C.GtkWindow
type WindowAccessible C.GtkWindowAccessible
type WindowGroup C.GtkWindowGroup

// interfaces
type Activatable C.GtkActivatable
type AppChooser C.GtkAppChooser
type Buildable C.GtkBuildable
type CellAccessibleParent C.GtkCellAccessibleParent
type CellEditable C.GtkCellEditable
type CellLayout C.GtkCellLayout
type Editable C.GtkEditable
type FileChooser C.GtkFileChooser
type FontChooser C.GtkFontChooser
type Orientable C.GtkOrientable
type PrintOperationPreview C.GtkPrintOperationPreview
type RecentChooser C.GtkRecentChooser
type Scrollable C.GtkScrollable
type StyleProvider C.GtkStyleProvider
type ToolShell C.GtkToolShell
type TreeDragDest C.GtkTreeDragDest
type TreeDragSource C.GtkTreeDragSource
type TreeModel C.GtkTreeModel
type TreeSortable C.GtkTreeSortable

func Fn_accel_groups_activate(object *gobject.Object, accelKey uint, accelMods c.UndefinedParamType) {
}

func Fn_accel_groups_from_object(object *gobject.Object) {}

func Fn_accelerator_get_default_mod_mask() {}

func Fn_accelerator_name(acceleratorKey uint, acceleratorMods c.UndefinedParamType) {}

func Fn_accelerator_parse(accelerator string) {}

func Fn_accelerator_set_default_mod_mask(defaultModMask c.UndefinedParamType) {}

func Fn_accelerator_valid(keyval uint, modifiers c.UndefinedParamType) {}

func Fn_binding_entry_add_signall(bindingSet *BindingSet, keyval uint, modifiers c.UndefinedParamType, signalName string, bindingArgs *glib.SList) {
}

func Fn_binding_entry_remove(bindingSet *BindingSet, keyval uint, modifiers c.UndefinedParamType) {}

func Fn_binding_set_by_class(objectClass unsafe.Pointer) {}

func Fn_binding_set_find(setName string) {}

func Fn_binding_set_new(setName string) {}

func Fn_bindings_activate(object *gobject.Object, keyval uint, modifiers c.UndefinedParamType) {}

func Fn_builder_error_quark() {}

func Fn_check_version(requiredMajor uint, requiredMinor uint, requiredMicro uint) {}

func Fn_css_provider_error_quark() {}

func Fn_disable_setlocale() {}

func Fn_distribute_natural_allocation(extraSpace int, nRequestedSizes uint, sizes *RequestedSize) {}

func Fn_drag_finish(context *gdk.DragContext, success bool, del bool, time uint32) {}

func Fn_drag_get_source_widget(context *gdk.DragContext) {}

func Fn_drag_set_icon_default(context *gdk.DragContext) {}

func Fn_drag_set_icon_pixbuf(context *gdk.DragContext, pixbuf *gdkpixbuf.Pixbuf, hotX int, hotY int) {
}

func Fn_drag_set_icon_stock(context *gdk.DragContext, stockId string, hotX int, hotY int) {}

func Fn_drag_set_icon_surface(context *gdk.DragContext, surface cairo.Surface) {}

func Fn_drag_set_icon_widget(context *gdk.DragContext, widget *Widget, hotX int, hotY int) {}

func Fn_events_pending() {}

func Fn_false() {}

func Fn_get_current_event() {}

func Fn_get_current_event_device() {}

func Fn_get_current_event_state() {}

func Fn_get_current_event_time() {}

func Fn_get_debug_flags() {}

func Fn_get_default_language() {}

func Fn_get_event_widget(event c.UndefinedParamType) {}

func Fn_grab_get_current() {}

func Fn_icon_size_from_name(name string) {}

func Fn_icon_size_get_name(size IconSize) {}

func Fn_icon_size_lookup(size IconSize) {}

func Fn_icon_size_register(name string, width int, height int) {}

func Fn_icon_size_register_alias(alias string, target IconSize) {}

func Fn_icon_theme_error_quark() {}

func Fn_init(argc *int, argv c.UndefinedParamType) {}

func Fn_init_check(argc *int, argv c.UndefinedParamType) {}

func Fn_key_snooper_install(snooper c.UndefinedParamType, funcData unsafe.Pointer) {}

func Fn_key_snooper_remove(snooperHandlerId uint) {}

func Fn_main() {}

func Fn_main_do_event(event c.UndefinedParamType) {}

func Fn_main_iteration() {}

func Fn_main_iteration_do(blocking bool) {}

func Fn_main_level() {}

func Fn_main_quit() {}

func Fn_paint_arrow(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, arrowType ArrowType, fill bool, x int, y int, width int, height int) {
}

func Fn_paint_box(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
}

func Fn_paint_box_gap(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int) {
}

func Fn_paint_check(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
}

func Fn_paint_diamond(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
}

func Fn_paint_expander(style *Style, cr cairo.Context, stateType StateType, widget *Widget, detail string, x int, y int, expanderStyle ExpanderStyle) {
}

func Fn_paint_extension(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, gapSide PositionType) {
}

func Fn_paint_flat_box(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
}

func Fn_paint_focus(style *Style, cr cairo.Context, stateType StateType, widget *Widget, detail string, x int, y int, width int, height int) {
}

func Fn_paint_handle(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, orientation Orientation) {
}

func Fn_paint_hline(style *Style, cr cairo.Context, stateType StateType, widget *Widget, detail string, x1 int, x2 int, y int) {
}

func Fn_paint_layout(style *Style, cr cairo.Context, stateType StateType, useText bool, widget *Widget, detail string, x int, y int, layout *pango.Layout) {
}

func Fn_paint_option(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
}

func Fn_paint_resize_grip(style *Style, cr cairo.Context, stateType StateType, widget *Widget, detail string, edge c.UndefinedParamType, x int, y int, width int, height int) {
}

func Fn_paint_shadow(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
}

func Fn_paint_shadow_gap(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, gapSide PositionType, gapX int, gapWidth int) {
}

func Fn_paint_slider(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int, orientation Orientation) {
}

func Fn_paint_spinner(style *Style, cr cairo.Context, stateType StateType, widget *Widget, detail string, step uint, x int, y int, width int, height int) {
}

func Fn_paint_tab(style *Style, cr cairo.Context, stateType StateType, shadowType ShadowType, widget *Widget, detail string, x int, y int, width int, height int) {
}

func Fn_paint_vline(style *Style, cr cairo.Context, stateType StateType, widget *Widget, detail string, y1 int, y2 int, x int) {
}

func Fn_parse_args(argc *int, argv c.UndefinedParamType) {}

func Fn_propagate_event(widget *Widget, event c.UndefinedParamType) {}

func Fn_rc_add_default_file(filename string) {}

func Fn_rc_find_module_in_path(moduleFile string) {}

func Fn_rc_find_pixmap_in_path(settings *Settings, scanner *glib.Scanner, pixmapFile string) {}

func Fn_rc_get_default_files() {}

func Fn_rc_get_im_module_file() {}

func Fn_rc_get_im_module_path() {}

func Fn_rc_get_module_dir() {}

func Fn_rc_get_style(widget *Widget) {}

func Fn_rc_get_style_by_paths(settings *Settings, widgetPath string, classPath string, type_ c.UndefinedParamType) {
}

func Fn_rc_get_theme_dir() {}

func Fn_rc_parse(filename string) {}

func Fn_rc_parse_color(scanner *glib.Scanner) {}

func Fn_rc_parse_priority(scanner *glib.Scanner, priority PathPriorityType) {}

func Fn_rc_parse_state(scanner *glib.Scanner) {}

func Fn_rc_parse_string(rcString string) {}

func Fn_rc_property_parse_border(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) {
}

func Fn_rc_property_parse_color(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) {
}

func Fn_rc_property_parse_enum(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) {
}

func Fn_rc_property_parse_flags(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) {
}

func Fn_rc_property_parse_requisition(pspec *gobject.ParamSpec, gstring *glib.String, propertyValue *gobject.Value) {
}

func Fn_rc_reparse_all() {}

func Fn_rc_reparse_all_for_settings(settings *Settings, forceLoad bool) {}

func Fn_rc_scanner_new() {}

func Fn_rc_set_default_files(filenames c.UndefinedParamType) {}

func Fn_recent_chooser_error_quark() {}

func Fn_recent_manager_error_quark() {}

func Fn_selection_add_target(widget *Widget, selection gdk.Atom, target gdk.Atom, info uint) {}

func Fn_selection_add_targets(widget *Widget, selection gdk.Atom, targets c.UndefinedParamType, ntargets uint) {
}

func Fn_selection_clear_targets(widget *Widget, selection gdk.Atom) {}

func Fn_selection_convert(widget *Widget, selection gdk.Atom, target gdk.Atom, time uint32) {}

func Fn_selection_owner_set(widget *Widget, selection gdk.Atom, time uint32) {}

func Fn_selection_remove_all(widget *Widget) {}

func Fn_set_debug_flags(flags uint) {}

// UNSUPPORTED : show_about_dialog : has varargs

func Fn_stock_add(items c.UndefinedParamType, nItems uint) {}

func Fn_stock_add_static(items c.UndefinedParamType, nItems uint) {}

func Fn_stock_list_ids() {}

func Fn_stock_lookup(stockId string) {}

// UNSUPPORTED : test_create_widget : has varargs

// UNSUPPORTED : test_display_button_window : has varargs

// UNSUPPORTED : test_init : has varargs

func Fn_tree_get_row_drag_data(selectionData *SelectionData) {}

func Fn_tree_row_reference_deleted(proxy *gobject.Object, path *TreePath) {}

func Fn_tree_row_reference_inserted(proxy *gobject.Object, path *TreePath) {}

func Fn_tree_row_reference_reordered(proxy *gobject.Object, path *TreePath, iter *TreeIter, newOrder c.UndefinedParamType) {
}

func Fn_tree_set_row_drag_data(selectionData *SelectionData, treeModel *TreeModel, path *TreePath) {}

func Fn_true() {}
