// Code generated - DO NOT EDIT.

package gdk

type AxisUse uint32

const (
	AxisUse_Ignore   AxisUse = int64(0)
	AxisUse_X        AxisUse = int64(1)
	AxisUse_Y        AxisUse = int64(2)
	AxisUse_Pressure AxisUse = int64(3)
	AxisUse_Xtilt    AxisUse = int64(4)
	AxisUse_Ytilt    AxisUse = int64(5)
	AxisUse_Wheel    AxisUse = int64(6)
	AxisUse_Distance AxisUse = int64(7)
	AxisUse_Rotation AxisUse = int64(8)
	AxisUse_Slider   AxisUse = int64(9)
	AxisUse_Last     AxisUse = int64(10)
)

type ByteOrder uint32

const (
	ByteOrder_LsbFirst ByteOrder = int64(0)
	ByteOrder_MsbFirst ByteOrder = int64(1)
)

type CrossingMode uint32

const (
	CrossingMode_Normal       CrossingMode = int64(0)
	CrossingMode_Grab         CrossingMode = int64(1)
	CrossingMode_Ungrab       CrossingMode = int64(2)
	CrossingMode_GtkGrab      CrossingMode = int64(3)
	CrossingMode_GtkUngrab    CrossingMode = int64(4)
	CrossingMode_StateChanged CrossingMode = int64(5)
	CrossingMode_TouchBegin   CrossingMode = int64(6)
	CrossingMode_TouchEnd     CrossingMode = int64(7)
	CrossingMode_DeviceSwitch CrossingMode = int64(8)
)

type CursorType int32

const (
	CursorType_XCursor           CursorType = int64(0)
	CursorType_Arrow             CursorType = int64(2)
	CursorType_BasedArrowDown    CursorType = int64(4)
	CursorType_BasedArrowUp      CursorType = int64(6)
	CursorType_Boat              CursorType = int64(8)
	CursorType_Bogosity          CursorType = int64(10)
	CursorType_BottomLeftCorner  CursorType = int64(12)
	CursorType_BottomRightCorner CursorType = int64(14)
	CursorType_BottomSide        CursorType = int64(16)
	CursorType_BottomTee         CursorType = int64(18)
	CursorType_BoxSpiral         CursorType = int64(20)
	CursorType_CenterPtr         CursorType = int64(22)
	CursorType_Circle            CursorType = int64(24)
	CursorType_Clock             CursorType = int64(26)
	CursorType_CoffeeMug         CursorType = int64(28)
	CursorType_Cross             CursorType = int64(30)
	CursorType_CrossReverse      CursorType = int64(32)
	CursorType_Crosshair         CursorType = int64(34)
	CursorType_DiamondCross      CursorType = int64(36)
	CursorType_Dot               CursorType = int64(38)
	CursorType_Dotbox            CursorType = int64(40)
	CursorType_DoubleArrow       CursorType = int64(42)
	CursorType_DraftLarge        CursorType = int64(44)
	CursorType_DraftSmall        CursorType = int64(46)
	CursorType_DrapedBox         CursorType = int64(48)
	CursorType_Exchange          CursorType = int64(50)
	CursorType_Fleur             CursorType = int64(52)
	CursorType_Gobbler           CursorType = int64(54)
	CursorType_Gumby             CursorType = int64(56)
	CursorType_Hand1             CursorType = int64(58)
	CursorType_Hand2             CursorType = int64(60)
	CursorType_Heart             CursorType = int64(62)
	CursorType_Icon              CursorType = int64(64)
	CursorType_IronCross         CursorType = int64(66)
	CursorType_LeftPtr           CursorType = int64(68)
	CursorType_LeftSide          CursorType = int64(70)
	CursorType_LeftTee           CursorType = int64(72)
	CursorType_Leftbutton        CursorType = int64(74)
	CursorType_LlAngle           CursorType = int64(76)
	CursorType_LrAngle           CursorType = int64(78)
	CursorType_Man               CursorType = int64(80)
	CursorType_Middlebutton      CursorType = int64(82)
	CursorType_Mouse             CursorType = int64(84)
	CursorType_Pencil            CursorType = int64(86)
	CursorType_Pirate            CursorType = int64(88)
	CursorType_Plus              CursorType = int64(90)
	CursorType_QuestionArrow     CursorType = int64(92)
	CursorType_RightPtr          CursorType = int64(94)
	CursorType_RightSide         CursorType = int64(96)
	CursorType_RightTee          CursorType = int64(98)
	CursorType_Rightbutton       CursorType = int64(100)
	CursorType_RtlLogo           CursorType = int64(102)
	CursorType_Sailboat          CursorType = int64(104)
	CursorType_SbDownArrow       CursorType = int64(106)
	CursorType_SbHDoubleArrow    CursorType = int64(108)
	CursorType_SbLeftArrow       CursorType = int64(110)
	CursorType_SbRightArrow      CursorType = int64(112)
	CursorType_SbUpArrow         CursorType = int64(114)
	CursorType_SbVDoubleArrow    CursorType = int64(116)
	CursorType_Shuttle           CursorType = int64(118)
	CursorType_Sizing            CursorType = int64(120)
	CursorType_Spider            CursorType = int64(122)
	CursorType_Spraycan          CursorType = int64(124)
	CursorType_Star              CursorType = int64(126)
	CursorType_Target            CursorType = int64(128)
	CursorType_Tcross            CursorType = int64(130)
	CursorType_TopLeftArrow      CursorType = int64(132)
	CursorType_TopLeftCorner     CursorType = int64(134)
	CursorType_TopRightCorner    CursorType = int64(136)
	CursorType_TopSide           CursorType = int64(138)
	CursorType_TopTee            CursorType = int64(140)
	CursorType_Trek              CursorType = int64(142)
	CursorType_UlAngle           CursorType = int64(144)
	CursorType_Umbrella          CursorType = int64(146)
	CursorType_UrAngle           CursorType = int64(148)
	CursorType_Watch             CursorType = int64(150)
	CursorType_Xterm             CursorType = int64(152)
	CursorType_LastCursor        CursorType = int64(153)
	CursorType_BlankCursor       CursorType = int64(-2)
	CursorType_CursorIsPixmap    CursorType = int64(-1)
)

type DevicePadFeature uint32

const (
	DevicePadFeature_Button DevicePadFeature = int64(0)
	DevicePadFeature_Ring   DevicePadFeature = int64(1)
	DevicePadFeature_Strip  DevicePadFeature = int64(2)
)

type DeviceToolType uint32

const (
	DeviceToolType_Unknown  DeviceToolType = int64(0)
	DeviceToolType_Pen      DeviceToolType = int64(1)
	DeviceToolType_Eraser   DeviceToolType = int64(2)
	DeviceToolType_Brush    DeviceToolType = int64(3)
	DeviceToolType_Pencil   DeviceToolType = int64(4)
	DeviceToolType_Airbrush DeviceToolType = int64(5)
	DeviceToolType_Mouse    DeviceToolType = int64(6)
	DeviceToolType_Lens     DeviceToolType = int64(7)
)

type DeviceType uint32

const (
	DeviceType_Master   DeviceType = int64(0)
	DeviceType_Slave    DeviceType = int64(1)
	DeviceType_Floating DeviceType = int64(2)
)

type DragCancelReason uint32

const (
	DragCancelReason_NoTarget      DragCancelReason = int64(0)
	DragCancelReason_UserCancelled DragCancelReason = int64(1)
	DragCancelReason_Error         DragCancelReason = int64(2)
)

type DragProtocol uint32

const (
	DragProtocol_None           DragProtocol = int64(0)
	DragProtocol_Motif          DragProtocol = int64(1)
	DragProtocol_Xdnd           DragProtocol = int64(2)
	DragProtocol_Rootwin        DragProtocol = int64(3)
	DragProtocol_Win32Dropfiles DragProtocol = int64(4)
	DragProtocol_Ole2           DragProtocol = int64(5)
	DragProtocol_Local          DragProtocol = int64(6)
	DragProtocol_Wayland        DragProtocol = int64(7)
)

type EventType int32

const (
	EventType_Nothing           EventType = int64(-1)
	EventType_Delete            EventType = int64(0)
	EventType_Destroy           EventType = int64(1)
	EventType_Expose            EventType = int64(2)
	EventType_MotionNotify      EventType = int64(3)
	EventType_ButtonPress       EventType = int64(4)
	EventType_2buttonPress      EventType = int64(5)
	EventType_DoubleButtonPress EventType = int64(5)
	EventType_3buttonPress      EventType = int64(6)
	EventType_TripleButtonPress EventType = int64(6)
	EventType_ButtonRelease     EventType = int64(7)
	EventType_KeyPress          EventType = int64(8)
	EventType_KeyRelease        EventType = int64(9)
	EventType_EnterNotify       EventType = int64(10)
	EventType_LeaveNotify       EventType = int64(11)
	EventType_FocusChange       EventType = int64(12)
	EventType_Configure         EventType = int64(13)
	EventType_Map               EventType = int64(14)
	EventType_Unmap             EventType = int64(15)
	EventType_PropertyNotify    EventType = int64(16)
	EventType_SelectionClear    EventType = int64(17)
	EventType_SelectionRequest  EventType = int64(18)
	EventType_SelectionNotify   EventType = int64(19)
	EventType_ProximityIn       EventType = int64(20)
	EventType_ProximityOut      EventType = int64(21)
	EventType_DragEnter         EventType = int64(22)
	EventType_DragLeave         EventType = int64(23)
	EventType_DragMotion        EventType = int64(24)
	EventType_DragStatus        EventType = int64(25)
	EventType_DropStart         EventType = int64(26)
	EventType_DropFinished      EventType = int64(27)
	EventType_ClientEvent       EventType = int64(28)
	EventType_VisibilityNotify  EventType = int64(29)
	EventType_Scroll            EventType = int64(31)
	EventType_WindowState       EventType = int64(32)
	EventType_Setting           EventType = int64(33)
	EventType_OwnerChange       EventType = int64(34)
	EventType_GrabBroken        EventType = int64(35)
	EventType_Damage            EventType = int64(36)
	EventType_TouchBegin        EventType = int64(37)
	EventType_TouchUpdate       EventType = int64(38)
	EventType_TouchEnd          EventType = int64(39)
	EventType_TouchCancel       EventType = int64(40)
	EventType_TouchpadSwipe     EventType = int64(41)
	EventType_TouchpadPinch     EventType = int64(42)
	EventType_PadButtonPress    EventType = int64(43)
	EventType_PadButtonRelease  EventType = int64(44)
	EventType_PadRing           EventType = int64(45)
	EventType_PadStrip          EventType = int64(46)
	EventType_PadGroupMode      EventType = int64(47)
	EventType_EventLast         EventType = int64(48)
)

type FilterReturn uint32

const (
	FilterReturn_Continue  FilterReturn = int64(0)
	FilterReturn_Translate FilterReturn = int64(1)
	FilterReturn_Remove    FilterReturn = int64(2)
)

type FullscreenMode uint32

const (
	FullscreenMode_CurrentMonitor FullscreenMode = int64(0)
	FullscreenMode_AllMonitors    FullscreenMode = int64(1)
)

type GLError uint32

const (
	GLError_NotAvailable       GLError = int64(0)
	GLError_UnsupportedFormat  GLError = int64(1)
	GLError_UnsupportedProfile GLError = int64(2)
)

type GrabOwnership uint32

const (
	GrabOwnership_None        GrabOwnership = int64(0)
	GrabOwnership_Window      GrabOwnership = int64(1)
	GrabOwnership_Application GrabOwnership = int64(2)
)

type GrabStatus uint32

const (
	GrabStatus_Success        GrabStatus = int64(0)
	GrabStatus_AlreadyGrabbed GrabStatus = int64(1)
	GrabStatus_InvalidTime    GrabStatus = int64(2)
	GrabStatus_NotViewable    GrabStatus = int64(3)
	GrabStatus_Frozen         GrabStatus = int64(4)
	GrabStatus_Failed         GrabStatus = int64(5)
)

type Gravity uint32

const (
	Gravity_NorthWest Gravity = int64(1)
	Gravity_North     Gravity = int64(2)
	Gravity_NorthEast Gravity = int64(3)
	Gravity_West      Gravity = int64(4)
	Gravity_Center    Gravity = int64(5)
	Gravity_East      Gravity = int64(6)
	Gravity_SouthWest Gravity = int64(7)
	Gravity_South     Gravity = int64(8)
	Gravity_SouthEast Gravity = int64(9)
	Gravity_Static    Gravity = int64(10)
)

type InputMode uint32

const (
	InputMode_Disabled InputMode = int64(0)
	InputMode_Screen   InputMode = int64(1)
	InputMode_Window   InputMode = int64(2)
)

type InputSource uint32

const (
	InputSource_Mouse       InputSource = int64(0)
	InputSource_Pen         InputSource = int64(1)
	InputSource_Eraser      InputSource = int64(2)
	InputSource_Cursor      InputSource = int64(3)
	InputSource_Keyboard    InputSource = int64(4)
	InputSource_Touchscreen InputSource = int64(5)
	InputSource_Touchpad    InputSource = int64(6)
	InputSource_Trackpoint  InputSource = int64(7)
	InputSource_TabletPad   InputSource = int64(8)
)

type ModifierIntent uint32

const (
	ModifierIntent_PrimaryAccelerator ModifierIntent = int64(0)
	ModifierIntent_ContextMenu        ModifierIntent = int64(1)
	ModifierIntent_ExtendSelection    ModifierIntent = int64(2)
	ModifierIntent_ModifySelection    ModifierIntent = int64(3)
	ModifierIntent_NoTextInput        ModifierIntent = int64(4)
	ModifierIntent_ShiftGroup         ModifierIntent = int64(5)
	ModifierIntent_DefaultModMask     ModifierIntent = int64(6)
)

type NotifyType uint32

const (
	NotifyType_Ancestor         NotifyType = int64(0)
	NotifyType_Virtual          NotifyType = int64(1)
	NotifyType_Inferior         NotifyType = int64(2)
	NotifyType_Nonlinear        NotifyType = int64(3)
	NotifyType_NonlinearVirtual NotifyType = int64(4)
	NotifyType_Unknown          NotifyType = int64(5)
)

type OwnerChange uint32

const (
	OwnerChange_NewOwner OwnerChange = int64(0)
	OwnerChange_Destroy  OwnerChange = int64(1)
	OwnerChange_Close    OwnerChange = int64(2)
)

type PropMode uint32

const (
	PropMode_Replace PropMode = int64(0)
	PropMode_Prepend PropMode = int64(1)
	PropMode_Append  PropMode = int64(2)
)

type PropertyState uint32

const (
	PropertyState_NewValue PropertyState = int64(0)
	PropertyState_Delete   PropertyState = int64(1)
)

type ScrollDirection uint32

const (
	ScrollDirection_Up     ScrollDirection = int64(0)
	ScrollDirection_Down   ScrollDirection = int64(1)
	ScrollDirection_Left   ScrollDirection = int64(2)
	ScrollDirection_Right  ScrollDirection = int64(3)
	ScrollDirection_Smooth ScrollDirection = int64(4)
)

type SettingAction uint32

const (
	SettingAction_New     SettingAction = int64(0)
	SettingAction_Changed SettingAction = int64(1)
	SettingAction_Deleted SettingAction = int64(2)
)

type Status int32

const (
	Status_Ok         Status = int64(0)
	Status_Error      Status = int64(-1)
	Status_ErrorParam Status = int64(-2)
	Status_ErrorFile  Status = int64(-3)
	Status_ErrorMem   Status = int64(-4)
)

type SubpixelLayout uint32

const (
	SubpixelLayout_Unknown       SubpixelLayout = int64(0)
	SubpixelLayout_None          SubpixelLayout = int64(1)
	SubpixelLayout_HorizontalRgb SubpixelLayout = int64(2)
	SubpixelLayout_HorizontalBgr SubpixelLayout = int64(3)
	SubpixelLayout_VerticalRgb   SubpixelLayout = int64(4)
	SubpixelLayout_VerticalBgr   SubpixelLayout = int64(5)
)

type TouchpadGesturePhase uint32

const (
	TouchpadGesturePhase_Begin  TouchpadGesturePhase = int64(0)
	TouchpadGesturePhase_Update TouchpadGesturePhase = int64(1)
	TouchpadGesturePhase_End    TouchpadGesturePhase = int64(2)
	TouchpadGesturePhase_Cancel TouchpadGesturePhase = int64(3)
)

type VisibilityState uint32

const (
	VisibilityState_Unobscured    VisibilityState = int64(0)
	VisibilityState_Partial       VisibilityState = int64(1)
	VisibilityState_FullyObscured VisibilityState = int64(2)
)

type VisualType uint32

const (
	VisualType_StaticGray  VisualType = int64(0)
	VisualType_Grayscale   VisualType = int64(1)
	VisualType_StaticColor VisualType = int64(2)
	VisualType_PseudoColor VisualType = int64(3)
	VisualType_TrueColor   VisualType = int64(4)
	VisualType_DirectColor VisualType = int64(5)
)

type WindowEdge uint32

const (
	WindowEdge_NorthWest WindowEdge = int64(0)
	WindowEdge_North     WindowEdge = int64(1)
	WindowEdge_NorthEast WindowEdge = int64(2)
	WindowEdge_West      WindowEdge = int64(3)
	WindowEdge_East      WindowEdge = int64(4)
	WindowEdge_SouthWest WindowEdge = int64(5)
	WindowEdge_South     WindowEdge = int64(6)
	WindowEdge_SouthEast WindowEdge = int64(7)
)

type WindowType uint32

const (
	WindowType_Root       WindowType = int64(0)
	WindowType_Toplevel   WindowType = int64(1)
	WindowType_Child      WindowType = int64(2)
	WindowType_Temp       WindowType = int64(3)
	WindowType_Foreign    WindowType = int64(4)
	WindowType_Offscreen  WindowType = int64(5)
	WindowType_Subsurface WindowType = int64(6)
)

type WindowTypeHint uint32

const (
	WindowTypeHint_Normal       WindowTypeHint = int64(0)
	WindowTypeHint_Dialog       WindowTypeHint = int64(1)
	WindowTypeHint_Menu         WindowTypeHint = int64(2)
	WindowTypeHint_Toolbar      WindowTypeHint = int64(3)
	WindowTypeHint_Splashscreen WindowTypeHint = int64(4)
	WindowTypeHint_Utility      WindowTypeHint = int64(5)
	WindowTypeHint_Dock         WindowTypeHint = int64(6)
	WindowTypeHint_Desktop      WindowTypeHint = int64(7)
	WindowTypeHint_DropdownMenu WindowTypeHint = int64(8)
	WindowTypeHint_PopupMenu    WindowTypeHint = int64(9)
	WindowTypeHint_Tooltip      WindowTypeHint = int64(10)
	WindowTypeHint_Notification WindowTypeHint = int64(11)
	WindowTypeHint_Combo        WindowTypeHint = int64(12)
	WindowTypeHint_Dnd          WindowTypeHint = int64(13)
)

type WindowWindowClass uint32

const (
	WindowWindowClass_InputOutput WindowWindowClass = int64(0)
	WindowWindowClass_InputOnly   WindowWindowClass = int64(1)
)
