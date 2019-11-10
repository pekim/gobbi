// Code generated - DO NOT EDIT.

package gdk

type AnchorHints uint32

const (
	AnchorHints_FlipX   AnchorHints = int64(1)
	AnchorHints_FlipY   AnchorHints = int64(2)
	AnchorHints_SlideX  AnchorHints = int64(4)
	AnchorHints_SlideY  AnchorHints = int64(8)
	AnchorHints_ResizeX AnchorHints = int64(16)
	AnchorHints_ResizeY AnchorHints = int64(32)
	AnchorHints_Flip    AnchorHints = int64(3)
	AnchorHints_Slide   AnchorHints = int64(12)
	AnchorHints_Resize  AnchorHints = int64(48)
)

type AxisFlags uint32

const (
	AxisFlags_X        AxisFlags = int64(2)
	AxisFlags_Y        AxisFlags = int64(4)
	AxisFlags_Pressure AxisFlags = int64(8)
	AxisFlags_Xtilt    AxisFlags = int64(16)
	AxisFlags_Ytilt    AxisFlags = int64(32)
	AxisFlags_Wheel    AxisFlags = int64(64)
	AxisFlags_Distance AxisFlags = int64(128)
	AxisFlags_Rotation AxisFlags = int64(256)
	AxisFlags_Slider   AxisFlags = int64(512)
)

type DragAction uint32

const (
	DragAction_Default DragAction = int64(1)
	DragAction_Copy    DragAction = int64(2)
	DragAction_Move    DragAction = int64(4)
	DragAction_Link    DragAction = int64(8)
	DragAction_Private DragAction = int64(16)
	DragAction_Ask     DragAction = int64(32)
)

type EventMask uint32

const (
	EventMask_ExposureMask          EventMask = int64(2)
	EventMask_PointerMotionMask     EventMask = int64(4)
	EventMask_PointerMotionHintMask EventMask = int64(8)
	EventMask_ButtonMotionMask      EventMask = int64(16)
	EventMask_Button1MotionMask     EventMask = int64(32)
	EventMask_Button2MotionMask     EventMask = int64(64)
	EventMask_Button3MotionMask     EventMask = int64(128)
	EventMask_ButtonPressMask       EventMask = int64(256)
	EventMask_ButtonReleaseMask     EventMask = int64(512)
	EventMask_KeyPressMask          EventMask = int64(1024)
	EventMask_KeyReleaseMask        EventMask = int64(2048)
	EventMask_EnterNotifyMask       EventMask = int64(4096)
	EventMask_LeaveNotifyMask       EventMask = int64(8192)
	EventMask_FocusChangeMask       EventMask = int64(16384)
	EventMask_StructureMask         EventMask = int64(32768)
	EventMask_PropertyChangeMask    EventMask = int64(65536)
	EventMask_VisibilityNotifyMask  EventMask = int64(131072)
	EventMask_ProximityInMask       EventMask = int64(262144)
	EventMask_ProximityOutMask      EventMask = int64(524288)
	EventMask_SubstructureMask      EventMask = int64(1048576)
	EventMask_ScrollMask            EventMask = int64(2097152)
	EventMask_TouchMask             EventMask = int64(4194304)
	EventMask_SmoothScrollMask      EventMask = int64(8388608)
	EventMask_TouchpadGestureMask   EventMask = int64(16777216)
	EventMask_TabletPadMask         EventMask = int64(33554432)
	EventMask_AllEventsMask         EventMask = int64(67108862)
)

type FrameClockPhase uint32

const (
	FrameClockPhase_None         FrameClockPhase = int64(0)
	FrameClockPhase_FlushEvents  FrameClockPhase = int64(1)
	FrameClockPhase_BeforePaint  FrameClockPhase = int64(2)
	FrameClockPhase_Update       FrameClockPhase = int64(4)
	FrameClockPhase_Layout       FrameClockPhase = int64(8)
	FrameClockPhase_Paint        FrameClockPhase = int64(16)
	FrameClockPhase_ResumeEvents FrameClockPhase = int64(32)
	FrameClockPhase_AfterPaint   FrameClockPhase = int64(64)
)

type ModifierType uint32

const (
	ModifierType_ShiftMask              ModifierType = int64(1)
	ModifierType_LockMask               ModifierType = int64(2)
	ModifierType_ControlMask            ModifierType = int64(4)
	ModifierType_Mod1Mask               ModifierType = int64(8)
	ModifierType_Mod2Mask               ModifierType = int64(16)
	ModifierType_Mod3Mask               ModifierType = int64(32)
	ModifierType_Mod4Mask               ModifierType = int64(64)
	ModifierType_Mod5Mask               ModifierType = int64(128)
	ModifierType_Button1Mask            ModifierType = int64(256)
	ModifierType_Button2Mask            ModifierType = int64(512)
	ModifierType_Button3Mask            ModifierType = int64(1024)
	ModifierType_Button4Mask            ModifierType = int64(2048)
	ModifierType_Button5Mask            ModifierType = int64(4096)
	ModifierType_ModifierReserved13Mask ModifierType = int64(8192)
	ModifierType_ModifierReserved14Mask ModifierType = int64(16384)
	ModifierType_ModifierReserved15Mask ModifierType = int64(32768)
	ModifierType_ModifierReserved16Mask ModifierType = int64(65536)
	ModifierType_ModifierReserved17Mask ModifierType = int64(131072)
	ModifierType_ModifierReserved18Mask ModifierType = int64(262144)
	ModifierType_ModifierReserved19Mask ModifierType = int64(524288)
	ModifierType_ModifierReserved20Mask ModifierType = int64(1048576)
	ModifierType_ModifierReserved21Mask ModifierType = int64(2097152)
	ModifierType_ModifierReserved22Mask ModifierType = int64(4194304)
	ModifierType_ModifierReserved23Mask ModifierType = int64(8388608)
	ModifierType_ModifierReserved24Mask ModifierType = int64(16777216)
	ModifierType_ModifierReserved25Mask ModifierType = int64(33554432)
	ModifierType_SuperMask              ModifierType = int64(67108864)
	ModifierType_HyperMask              ModifierType = int64(134217728)
	ModifierType_MetaMask               ModifierType = int64(268435456)
	ModifierType_ModifierReserved29Mask ModifierType = int64(536870912)
	ModifierType_ReleaseMask            ModifierType = int64(1073741824)
	ModifierType_ModifierMask           ModifierType = int64(1543512063)
)

type SeatCapabilities uint32

const (
	SeatCapabilities_None         SeatCapabilities = int64(0)
	SeatCapabilities_Pointer      SeatCapabilities = int64(1)
	SeatCapabilities_Touch        SeatCapabilities = int64(2)
	SeatCapabilities_TabletStylus SeatCapabilities = int64(4)
	SeatCapabilities_Keyboard     SeatCapabilities = int64(8)
	SeatCapabilities_AllPointing  SeatCapabilities = int64(7)
	SeatCapabilities_All          SeatCapabilities = int64(15)
)

type WMDecoration uint32

const (
	WMDecoration_All      WMDecoration = int64(1)
	WMDecoration_Border   WMDecoration = int64(2)
	WMDecoration_Resizeh  WMDecoration = int64(4)
	WMDecoration_Title    WMDecoration = int64(8)
	WMDecoration_Menu     WMDecoration = int64(16)
	WMDecoration_Minimize WMDecoration = int64(32)
	WMDecoration_Maximize WMDecoration = int64(64)
)

type WMFunction uint32

const (
	WMFunction_All      WMFunction = int64(1)
	WMFunction_Resize   WMFunction = int64(2)
	WMFunction_Move     WMFunction = int64(4)
	WMFunction_Minimize WMFunction = int64(8)
	WMFunction_Maximize WMFunction = int64(16)
	WMFunction_Close    WMFunction = int64(32)
)

type WindowAttributesType uint32

const (
	WindowAttributesType_Title    WindowAttributesType = int64(2)
	WindowAttributesType_X        WindowAttributesType = int64(4)
	WindowAttributesType_Y        WindowAttributesType = int64(8)
	WindowAttributesType_Cursor   WindowAttributesType = int64(16)
	WindowAttributesType_Visual   WindowAttributesType = int64(32)
	WindowAttributesType_Wmclass  WindowAttributesType = int64(64)
	WindowAttributesType_Noredir  WindowAttributesType = int64(128)
	WindowAttributesType_TypeHint WindowAttributesType = int64(256)
)

type WindowHints uint32

const (
	WindowHints_Pos        WindowHints = int64(1)
	WindowHints_MinSize    WindowHints = int64(2)
	WindowHints_MaxSize    WindowHints = int64(4)
	WindowHints_BaseSize   WindowHints = int64(8)
	WindowHints_Aspect     WindowHints = int64(16)
	WindowHints_ResizeInc  WindowHints = int64(32)
	WindowHints_WinGravity WindowHints = int64(64)
	WindowHints_UserPos    WindowHints = int64(128)
	WindowHints_UserSize   WindowHints = int64(256)
)

type WindowState uint32

const (
	WindowState_Withdrawn       WindowState = int64(1)
	WindowState_Iconified       WindowState = int64(2)
	WindowState_Maximized       WindowState = int64(4)
	WindowState_Sticky          WindowState = int64(8)
	WindowState_Fullscreen      WindowState = int64(16)
	WindowState_Above           WindowState = int64(32)
	WindowState_Below           WindowState = int64(64)
	WindowState_Focused         WindowState = int64(128)
	WindowState_Tiled           WindowState = int64(256)
	WindowState_TopTiled        WindowState = int64(512)
	WindowState_TopResizable    WindowState = int64(1024)
	WindowState_RightTiled      WindowState = int64(2048)
	WindowState_RightResizable  WindowState = int64(4096)
	WindowState_BottomTiled     WindowState = int64(8192)
	WindowState_BottomResizable WindowState = int64(16384)
	WindowState_LeftTiled       WindowState = int64(32768)
	WindowState_LeftResizable   WindowState = int64(65536)
)
