// Code generated - DO NOT EDIT.

package gdk

// Axisuse is a representation of the C type AxisUse.
type Axisuse int

const (
	Axisuse_Ignore   Axisuse = 0
	Axisuse_X        Axisuse = 1
	Axisuse_Y        Axisuse = 2
	Axisuse_Pressure Axisuse = 3
	Axisuse_Xtilt    Axisuse = 4
	Axisuse_Ytilt    Axisuse = 5
	Axisuse_Wheel    Axisuse = 6
	Axisuse_Distance Axisuse = 7
	Axisuse_Rotation Axisuse = 8
	Axisuse_Slider   Axisuse = 9
	Axisuse_Last     Axisuse = 10
)

// Byteorder is a representation of the C type ByteOrder.
type Byteorder int

const (
	Byteorder_LsbFirst Byteorder = 0
	Byteorder_MsbFirst Byteorder = 1
)

// Crossingmode is a representation of the C type CrossingMode.
type Crossingmode int

const (
	Crossingmode_Normal       Crossingmode = 0
	Crossingmode_Grab         Crossingmode = 1
	Crossingmode_Ungrab       Crossingmode = 2
	Crossingmode_GtkGrab      Crossingmode = 3
	Crossingmode_GtkUngrab    Crossingmode = 4
	Crossingmode_StateChanged Crossingmode = 5
	Crossingmode_TouchBegin   Crossingmode = 6
	Crossingmode_TouchEnd     Crossingmode = 7
	Crossingmode_DeviceSwitch Crossingmode = 8
)

// Cursortype is a representation of the C type CursorType.
type Cursortype int

const (
	Cursortype_XCursor           Cursortype = 0
	Cursortype_Arrow             Cursortype = 2
	Cursortype_BasedArrowDown    Cursortype = 4
	Cursortype_BasedArrowUp      Cursortype = 6
	Cursortype_Boat              Cursortype = 8
	Cursortype_Bogosity          Cursortype = 10
	Cursortype_BottomLeftCorner  Cursortype = 12
	Cursortype_BottomRightCorner Cursortype = 14
	Cursortype_BottomSide        Cursortype = 16
	Cursortype_BottomTee         Cursortype = 18
	Cursortype_BoxSpiral         Cursortype = 20
	Cursortype_CenterPtr         Cursortype = 22
	Cursortype_Circle            Cursortype = 24
	Cursortype_Clock             Cursortype = 26
	Cursortype_CoffeeMug         Cursortype = 28
	Cursortype_Cross             Cursortype = 30
	Cursortype_CrossReverse      Cursortype = 32
	Cursortype_Crosshair         Cursortype = 34
	Cursortype_DiamondCross      Cursortype = 36
	Cursortype_Dot               Cursortype = 38
	Cursortype_Dotbox            Cursortype = 40
	Cursortype_DoubleArrow       Cursortype = 42
	Cursortype_DraftLarge        Cursortype = 44
	Cursortype_DraftSmall        Cursortype = 46
	Cursortype_DrapedBox         Cursortype = 48
	Cursortype_Exchange          Cursortype = 50
	Cursortype_Fleur             Cursortype = 52
	Cursortype_Gobbler           Cursortype = 54
	Cursortype_Gumby             Cursortype = 56
	Cursortype_Hand1             Cursortype = 58
	Cursortype_Hand2             Cursortype = 60
	Cursortype_Heart             Cursortype = 62
	Cursortype_Icon              Cursortype = 64
	Cursortype_IronCross         Cursortype = 66
	Cursortype_LeftPtr           Cursortype = 68
	Cursortype_LeftSide          Cursortype = 70
	Cursortype_LeftTee           Cursortype = 72
	Cursortype_Leftbutton        Cursortype = 74
	Cursortype_LlAngle           Cursortype = 76
	Cursortype_LrAngle           Cursortype = 78
	Cursortype_Man               Cursortype = 80
	Cursortype_Middlebutton      Cursortype = 82
	Cursortype_Mouse             Cursortype = 84
	Cursortype_Pencil            Cursortype = 86
	Cursortype_Pirate            Cursortype = 88
	Cursortype_Plus              Cursortype = 90
	Cursortype_QuestionArrow     Cursortype = 92
	Cursortype_RightPtr          Cursortype = 94
	Cursortype_RightSide         Cursortype = 96
	Cursortype_RightTee          Cursortype = 98
	Cursortype_Rightbutton       Cursortype = 100
	Cursortype_RtlLogo           Cursortype = 102
	Cursortype_Sailboat          Cursortype = 104
	Cursortype_SbDownArrow       Cursortype = 106
	Cursortype_SbHDoubleArrow    Cursortype = 108
	Cursortype_SbLeftArrow       Cursortype = 110
	Cursortype_SbRightArrow      Cursortype = 112
	Cursortype_SbUpArrow         Cursortype = 114
	Cursortype_SbVDoubleArrow    Cursortype = 116
	Cursortype_Shuttle           Cursortype = 118
	Cursortype_Sizing            Cursortype = 120
	Cursortype_Spider            Cursortype = 122
	Cursortype_Spraycan          Cursortype = 124
	Cursortype_Star              Cursortype = 126
	Cursortype_Target            Cursortype = 128
	Cursortype_Tcross            Cursortype = 130
	Cursortype_TopLeftArrow      Cursortype = 132
	Cursortype_TopLeftCorner     Cursortype = 134
	Cursortype_TopRightCorner    Cursortype = 136
	Cursortype_TopSide           Cursortype = 138
	Cursortype_TopTee            Cursortype = 140
	Cursortype_Trek              Cursortype = 142
	Cursortype_UlAngle           Cursortype = 144
	Cursortype_Umbrella          Cursortype = 146
	Cursortype_UrAngle           Cursortype = 148
	Cursortype_Watch             Cursortype = 150
	Cursortype_Xterm             Cursortype = 152
	Cursortype_LastCursor        Cursortype = 153
	Cursortype_BlankCursor       Cursortype = -2
	Cursortype_CursorIsPixmap    Cursortype = -1
)

// Devicepadfeature is a representation of the C type DevicePadFeature.
type Devicepadfeature int

const (
	Devicepadfeature_Button Devicepadfeature = 0
	Devicepadfeature_Ring   Devicepadfeature = 1
	Devicepadfeature_Strip  Devicepadfeature = 2
)

// Devicetooltype is a representation of the C type DeviceToolType.
//
// since 3.22
type Devicetooltype int

const (
	Devicetooltype_Unknown  Devicetooltype = 0
	Devicetooltype_Pen      Devicetooltype = 1
	Devicetooltype_Eraser   Devicetooltype = 2
	Devicetooltype_Brush    Devicetooltype = 3
	Devicetooltype_Pencil   Devicetooltype = 4
	Devicetooltype_Airbrush Devicetooltype = 5
	Devicetooltype_Mouse    Devicetooltype = 6
	Devicetooltype_Lens     Devicetooltype = 7
)

// Devicetype is a representation of the C type DeviceType.
type Devicetype int

const (
	Devicetype_Master   Devicetype = 0
	Devicetype_Slave    Devicetype = 1
	Devicetype_Floating Devicetype = 2
)

// Dragcancelreason is a representation of the C type DragCancelReason.
//
// since 3.20
type Dragcancelreason int

const (
	Dragcancelreason_NoTarget      Dragcancelreason = 0
	Dragcancelreason_UserCancelled Dragcancelreason = 1
	Dragcancelreason_Error         Dragcancelreason = 2
)

// Dragprotocol is a representation of the C type DragProtocol.
type Dragprotocol int

const (
	Dragprotocol_None           Dragprotocol = 0
	Dragprotocol_Motif          Dragprotocol = 1
	Dragprotocol_Xdnd           Dragprotocol = 2
	Dragprotocol_Rootwin        Dragprotocol = 3
	Dragprotocol_Win32Dropfiles Dragprotocol = 4
	Dragprotocol_Ole2           Dragprotocol = 5
	Dragprotocol_Local          Dragprotocol = 6
	Dragprotocol_Wayland        Dragprotocol = 7
)

// Eventtype is a representation of the C type EventType.
type Eventtype int

const (
	Eventtype_Nothing           Eventtype = -1
	Eventtype_Delete            Eventtype = 0
	Eventtype_Destroy           Eventtype = 1
	Eventtype_Expose            Eventtype = 2
	Eventtype_MotionNotify      Eventtype = 3
	Eventtype_ButtonPress       Eventtype = 4
	Eventtype_2buttonPress      Eventtype = 5
	Eventtype_DoubleButtonPress Eventtype = 5
	Eventtype_3buttonPress      Eventtype = 6
	Eventtype_TripleButtonPress Eventtype = 6
	Eventtype_ButtonRelease     Eventtype = 7
	Eventtype_KeyPress          Eventtype = 8
	Eventtype_KeyRelease        Eventtype = 9
	Eventtype_EnterNotify       Eventtype = 10
	Eventtype_LeaveNotify       Eventtype = 11
	Eventtype_FocusChange       Eventtype = 12
	Eventtype_Configure         Eventtype = 13
	Eventtype_Map               Eventtype = 14
	Eventtype_Unmap             Eventtype = 15
	Eventtype_PropertyNotify    Eventtype = 16
	Eventtype_SelectionClear    Eventtype = 17
	Eventtype_SelectionRequest  Eventtype = 18
	Eventtype_SelectionNotify   Eventtype = 19
	Eventtype_ProximityIn       Eventtype = 20
	Eventtype_ProximityOut      Eventtype = 21
	Eventtype_DragEnter         Eventtype = 22
	Eventtype_DragLeave         Eventtype = 23
	Eventtype_DragMotion        Eventtype = 24
	Eventtype_DragStatus        Eventtype = 25
	Eventtype_DropStart         Eventtype = 26
	Eventtype_DropFinished      Eventtype = 27
	Eventtype_ClientEvent       Eventtype = 28
	Eventtype_VisibilityNotify  Eventtype = 29
	Eventtype_Scroll            Eventtype = 31
	Eventtype_WindowState       Eventtype = 32
	Eventtype_Setting           Eventtype = 33
	Eventtype_OwnerChange       Eventtype = 34
	Eventtype_GrabBroken        Eventtype = 35
	Eventtype_Damage            Eventtype = 36
	Eventtype_TouchBegin        Eventtype = 37
	Eventtype_TouchUpdate       Eventtype = 38
	Eventtype_TouchEnd          Eventtype = 39
	Eventtype_TouchCancel       Eventtype = 40
	Eventtype_TouchpadSwipe     Eventtype = 41
	Eventtype_TouchpadPinch     Eventtype = 42
	Eventtype_PadButtonPress    Eventtype = 43
	Eventtype_PadButtonRelease  Eventtype = 44
	Eventtype_PadRing           Eventtype = 45
	Eventtype_PadStrip          Eventtype = 46
	Eventtype_PadGroupMode      Eventtype = 47
	Eventtype_EventLast         Eventtype = 48
)

// Filterreturn is a representation of the C type FilterReturn.
type Filterreturn int

const (
	Filterreturn_Continue  Filterreturn = 0
	Filterreturn_Translate Filterreturn = 1
	Filterreturn_Remove    Filterreturn = 2
)

// Fullscreenmode is a representation of the C type FullscreenMode.
//
// since 3.8
type Fullscreenmode int

const (
	Fullscreenmode_CurrentMonitor Fullscreenmode = 0
	Fullscreenmode_AllMonitors    Fullscreenmode = 1
)

// Glerror is a representation of the C type GLError.
//
// since 3.16
type Glerror int

const (
	Glerror_NotAvailable       Glerror = 0
	Glerror_UnsupportedFormat  Glerror = 1
	Glerror_UnsupportedProfile Glerror = 2
)

// Grabownership is a representation of the C type GrabOwnership.
type Grabownership int

const (
	Grabownership_None        Grabownership = 0
	Grabownership_Window      Grabownership = 1
	Grabownership_Application Grabownership = 2
)

// Grabstatus is a representation of the C type GrabStatus.
type Grabstatus int

const (
	Grabstatus_Success        Grabstatus = 0
	Grabstatus_AlreadyGrabbed Grabstatus = 1
	Grabstatus_InvalidTime    Grabstatus = 2
	Grabstatus_NotViewable    Grabstatus = 3
	Grabstatus_Frozen         Grabstatus = 4
	Grabstatus_Failed         Grabstatus = 5
)

// Gravity is a representation of the C type Gravity.
type Gravity int

const (
	Gravity_NorthWest Gravity = 1
	Gravity_North     Gravity = 2
	Gravity_NorthEast Gravity = 3
	Gravity_West      Gravity = 4
	Gravity_Center    Gravity = 5
	Gravity_East      Gravity = 6
	Gravity_SouthWest Gravity = 7
	Gravity_South     Gravity = 8
	Gravity_SouthEast Gravity = 9
	Gravity_Static    Gravity = 10
)

// Inputmode is a representation of the C type InputMode.
type Inputmode int

const (
	Inputmode_Disabled Inputmode = 0
	Inputmode_Screen   Inputmode = 1
	Inputmode_Window   Inputmode = 2
)

// Inputsource is a representation of the C type InputSource.
type Inputsource int

const (
	Inputsource_Mouse       Inputsource = 0
	Inputsource_Pen         Inputsource = 1
	Inputsource_Eraser      Inputsource = 2
	Inputsource_Cursor      Inputsource = 3
	Inputsource_Keyboard    Inputsource = 4
	Inputsource_Touchscreen Inputsource = 5
	Inputsource_Touchpad    Inputsource = 6
	Inputsource_Trackpoint  Inputsource = 7
	Inputsource_TabletPad   Inputsource = 8
)

// Modifierintent is a representation of the C type ModifierIntent.
//
// since 3.4
type Modifierintent int

const (
	Modifierintent_PrimaryAccelerator Modifierintent = 0
	Modifierintent_ContextMenu        Modifierintent = 1
	Modifierintent_ExtendSelection    Modifierintent = 2
	Modifierintent_ModifySelection    Modifierintent = 3
	Modifierintent_NoTextInput        Modifierintent = 4
	Modifierintent_ShiftGroup         Modifierintent = 5
	Modifierintent_DefaultModMask     Modifierintent = 6
)

// Notifytype is a representation of the C type NotifyType.
type Notifytype int

const (
	Notifytype_Ancestor         Notifytype = 0
	Notifytype_Virtual          Notifytype = 1
	Notifytype_Inferior         Notifytype = 2
	Notifytype_Nonlinear        Notifytype = 3
	Notifytype_NonlinearVirtual Notifytype = 4
	Notifytype_Unknown          Notifytype = 5
)

// Ownerchange is a representation of the C type OwnerChange.
type Ownerchange int

const (
	Ownerchange_NewOwner Ownerchange = 0
	Ownerchange_Destroy  Ownerchange = 1
	Ownerchange_Close    Ownerchange = 2
)

// Propmode is a representation of the C type PropMode.
type Propmode int

const (
	Propmode_Replace Propmode = 0
	Propmode_Prepend Propmode = 1
	Propmode_Append  Propmode = 2
)

// Propertystate is a representation of the C type PropertyState.
type Propertystate int

const (
	Propertystate_NewValue Propertystate = 0
	Propertystate_Delete   Propertystate = 1
)

// Scrolldirection is a representation of the C type ScrollDirection.
type Scrolldirection int

const (
	Scrolldirection_Up     Scrolldirection = 0
	Scrolldirection_Down   Scrolldirection = 1
	Scrolldirection_Left   Scrolldirection = 2
	Scrolldirection_Right  Scrolldirection = 3
	Scrolldirection_Smooth Scrolldirection = 4
)

// Settingaction is a representation of the C type SettingAction.
type Settingaction int

const (
	Settingaction_New     Settingaction = 0
	Settingaction_Changed Settingaction = 1
	Settingaction_Deleted Settingaction = 2
)

// Status is a representation of the C type Status.
type Status int

const (
	Status_Ok         Status = 0
	Status_Error      Status = -1
	Status_ErrorParam Status = -2
	Status_ErrorFile  Status = -3
	Status_ErrorMem   Status = -4
)

// Subpixellayout is a representation of the C type SubpixelLayout.
//
// since 3.22
type Subpixellayout int

const (
	Subpixellayout_Unknown       Subpixellayout = 0
	Subpixellayout_None          Subpixellayout = 1
	Subpixellayout_HorizontalRgb Subpixellayout = 2
	Subpixellayout_HorizontalBgr Subpixellayout = 3
	Subpixellayout_VerticalRgb   Subpixellayout = 4
	Subpixellayout_VerticalBgr   Subpixellayout = 5
)

// Touchpadgesturephase is a representation of the C type TouchpadGesturePhase.
type Touchpadgesturephase int

const (
	Touchpadgesturephase_Begin  Touchpadgesturephase = 0
	Touchpadgesturephase_Update Touchpadgesturephase = 1
	Touchpadgesturephase_End    Touchpadgesturephase = 2
	Touchpadgesturephase_Cancel Touchpadgesturephase = 3
)

// Visibilitystate is a representation of the C type VisibilityState.
type Visibilitystate int

const (
	Visibilitystate_Unobscured    Visibilitystate = 0
	Visibilitystate_Partial       Visibilitystate = 1
	Visibilitystate_FullyObscured Visibilitystate = 2
)

// Visualtype is a representation of the C type VisualType.
type Visualtype int

const (
	Visualtype_StaticGray  Visualtype = 0
	Visualtype_Grayscale   Visualtype = 1
	Visualtype_StaticColor Visualtype = 2
	Visualtype_PseudoColor Visualtype = 3
	Visualtype_TrueColor   Visualtype = 4
	Visualtype_DirectColor Visualtype = 5
)

// Windowedge is a representation of the C type WindowEdge.
type Windowedge int

const (
	Windowedge_NorthWest Windowedge = 0
	Windowedge_North     Windowedge = 1
	Windowedge_NorthEast Windowedge = 2
	Windowedge_West      Windowedge = 3
	Windowedge_East      Windowedge = 4
	Windowedge_SouthWest Windowedge = 5
	Windowedge_South     Windowedge = 6
	Windowedge_SouthEast Windowedge = 7
)

// Windowtype is a representation of the C type WindowType.
type Windowtype int

const (
	Windowtype_Root       Windowtype = 0
	Windowtype_Toplevel   Windowtype = 1
	Windowtype_Child      Windowtype = 2
	Windowtype_Temp       Windowtype = 3
	Windowtype_Foreign    Windowtype = 4
	Windowtype_Offscreen  Windowtype = 5
	Windowtype_Subsurface Windowtype = 6
)

// Windowtypehint is a representation of the C type WindowTypeHint.
type Windowtypehint int

const (
	Windowtypehint_Normal       Windowtypehint = 0
	Windowtypehint_Dialog       Windowtypehint = 1
	Windowtypehint_Menu         Windowtypehint = 2
	Windowtypehint_Toolbar      Windowtypehint = 3
	Windowtypehint_Splashscreen Windowtypehint = 4
	Windowtypehint_Utility      Windowtypehint = 5
	Windowtypehint_Dock         Windowtypehint = 6
	Windowtypehint_Desktop      Windowtypehint = 7
	Windowtypehint_DropdownMenu Windowtypehint = 8
	Windowtypehint_PopupMenu    Windowtypehint = 9
	Windowtypehint_Tooltip      Windowtypehint = 10
	Windowtypehint_Notification Windowtypehint = 11
	Windowtypehint_Combo        Windowtypehint = 12
	Windowtypehint_Dnd          Windowtypehint = 13
)

// Windowwindowclass is a representation of the C type WindowWindowClass.
type Windowwindowclass int

const (
	Windowwindowclass_InputOutput Windowwindowclass = 0
	Windowwindowclass_InputOnly   Windowwindowclass = 1
)
