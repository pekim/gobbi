// Code generated - DO NOT EDIT.

package gdk

// Anchorhints is a representation of the C type AnchorHints.
//
// since 3.22
type Anchorhints int

const (
	Anchorhints_FlipX   Anchorhints = 1
	Anchorhints_FlipY   Anchorhints = 2
	Anchorhints_SlideX  Anchorhints = 4
	Anchorhints_SlideY  Anchorhints = 8
	Anchorhints_ResizeX Anchorhints = 16
	Anchorhints_ResizeY Anchorhints = 32
	Anchorhints_Flip    Anchorhints = 3
	Anchorhints_Slide   Anchorhints = 12
	Anchorhints_Resize  Anchorhints = 48
)

// Axisflags is a representation of the C type AxisFlags.
//
// since 3.22
type Axisflags int

const (
	Axisflags_X        Axisflags = 2
	Axisflags_Y        Axisflags = 4
	Axisflags_Pressure Axisflags = 8
	Axisflags_Xtilt    Axisflags = 16
	Axisflags_Ytilt    Axisflags = 32
	Axisflags_Wheel    Axisflags = 64
	Axisflags_Distance Axisflags = 128
	Axisflags_Rotation Axisflags = 256
	Axisflags_Slider   Axisflags = 512
)

// Dragaction is a representation of the C type DragAction.
type Dragaction int

const (
	Dragaction_Default Dragaction = 1
	Dragaction_Copy    Dragaction = 2
	Dragaction_Move    Dragaction = 4
	Dragaction_Link    Dragaction = 8
	Dragaction_Private Dragaction = 16
	Dragaction_Ask     Dragaction = 32
)

// Eventmask is a representation of the C type EventMask.
type Eventmask int

const (
	Eventmask_ExposureMask          Eventmask = 2
	Eventmask_PointerMotionMask     Eventmask = 4
	Eventmask_PointerMotionHintMask Eventmask = 8
	Eventmask_ButtonMotionMask      Eventmask = 16
	Eventmask_Button1MotionMask     Eventmask = 32
	Eventmask_Button2MotionMask     Eventmask = 64
	Eventmask_Button3MotionMask     Eventmask = 128
	Eventmask_ButtonPressMask       Eventmask = 256
	Eventmask_ButtonReleaseMask     Eventmask = 512
	Eventmask_KeyPressMask          Eventmask = 1024
	Eventmask_KeyReleaseMask        Eventmask = 2048
	Eventmask_EnterNotifyMask       Eventmask = 4096
	Eventmask_LeaveNotifyMask       Eventmask = 8192
	Eventmask_FocusChangeMask       Eventmask = 16384
	Eventmask_StructureMask         Eventmask = 32768
	Eventmask_PropertyChangeMask    Eventmask = 65536
	Eventmask_VisibilityNotifyMask  Eventmask = 131072
	Eventmask_ProximityInMask       Eventmask = 262144
	Eventmask_ProximityOutMask      Eventmask = 524288
	Eventmask_SubstructureMask      Eventmask = 1048576
	Eventmask_ScrollMask            Eventmask = 2097152
	Eventmask_TouchMask             Eventmask = 4194304
	Eventmask_SmoothScrollMask      Eventmask = 8388608
	Eventmask_TouchpadGestureMask   Eventmask = 16777216
	Eventmask_TabletPadMask         Eventmask = 33554432
	Eventmask_AllEventsMask         Eventmask = 67108862
)

// Frameclockphase is a representation of the C type FrameClockPhase.
//
// since 3.8
type Frameclockphase int

const (
	Frameclockphase_None         Frameclockphase = 0
	Frameclockphase_FlushEvents  Frameclockphase = 1
	Frameclockphase_BeforePaint  Frameclockphase = 2
	Frameclockphase_Update       Frameclockphase = 4
	Frameclockphase_Layout       Frameclockphase = 8
	Frameclockphase_Paint        Frameclockphase = 16
	Frameclockphase_ResumeEvents Frameclockphase = 32
	Frameclockphase_AfterPaint   Frameclockphase = 64
)

// Modifiertype is a representation of the C type ModifierType.
type Modifiertype int

const (
	Modifiertype_ShiftMask              Modifiertype = 1
	Modifiertype_LockMask               Modifiertype = 2
	Modifiertype_ControlMask            Modifiertype = 4
	Modifiertype_Mod1Mask               Modifiertype = 8
	Modifiertype_Mod2Mask               Modifiertype = 16
	Modifiertype_Mod3Mask               Modifiertype = 32
	Modifiertype_Mod4Mask               Modifiertype = 64
	Modifiertype_Mod5Mask               Modifiertype = 128
	Modifiertype_Button1Mask            Modifiertype = 256
	Modifiertype_Button2Mask            Modifiertype = 512
	Modifiertype_Button3Mask            Modifiertype = 1024
	Modifiertype_Button4Mask            Modifiertype = 2048
	Modifiertype_Button5Mask            Modifiertype = 4096
	Modifiertype_ModifierReserved13Mask Modifiertype = 8192
	Modifiertype_ModifierReserved14Mask Modifiertype = 16384
	Modifiertype_ModifierReserved15Mask Modifiertype = 32768
	Modifiertype_ModifierReserved16Mask Modifiertype = 65536
	Modifiertype_ModifierReserved17Mask Modifiertype = 131072
	Modifiertype_ModifierReserved18Mask Modifiertype = 262144
	Modifiertype_ModifierReserved19Mask Modifiertype = 524288
	Modifiertype_ModifierReserved20Mask Modifiertype = 1048576
	Modifiertype_ModifierReserved21Mask Modifiertype = 2097152
	Modifiertype_ModifierReserved22Mask Modifiertype = 4194304
	Modifiertype_ModifierReserved23Mask Modifiertype = 8388608
	Modifiertype_ModifierReserved24Mask Modifiertype = 16777216
	Modifiertype_ModifierReserved25Mask Modifiertype = 33554432
	Modifiertype_SuperMask              Modifiertype = 67108864
	Modifiertype_HyperMask              Modifiertype = 134217728
	Modifiertype_MetaMask               Modifiertype = 268435456
	Modifiertype_ModifierReserved29Mask Modifiertype = 536870912
	Modifiertype_ReleaseMask            Modifiertype = 1073741824
	Modifiertype_ModifierMask           Modifiertype = 1543512063
)

// Seatcapabilities is a representation of the C type SeatCapabilities.
//
// since 3.20
type Seatcapabilities int

const (
	Seatcapabilities_None         Seatcapabilities = 0
	Seatcapabilities_Pointer      Seatcapabilities = 1
	Seatcapabilities_Touch        Seatcapabilities = 2
	Seatcapabilities_TabletStylus Seatcapabilities = 4
	Seatcapabilities_Keyboard     Seatcapabilities = 8
	Seatcapabilities_AllPointing  Seatcapabilities = 7
	Seatcapabilities_All          Seatcapabilities = 15
)

// Wmdecoration is a representation of the C type WMDecoration.
type Wmdecoration int

const (
	Wmdecoration_All      Wmdecoration = 1
	Wmdecoration_Border   Wmdecoration = 2
	Wmdecoration_Resizeh  Wmdecoration = 4
	Wmdecoration_Title    Wmdecoration = 8
	Wmdecoration_Menu     Wmdecoration = 16
	Wmdecoration_Minimize Wmdecoration = 32
	Wmdecoration_Maximize Wmdecoration = 64
)

// Wmfunction is a representation of the C type WMFunction.
type Wmfunction int

const (
	Wmfunction_All      Wmfunction = 1
	Wmfunction_Resize   Wmfunction = 2
	Wmfunction_Move     Wmfunction = 4
	Wmfunction_Minimize Wmfunction = 8
	Wmfunction_Maximize Wmfunction = 16
	Wmfunction_Close    Wmfunction = 32
)

// Windowattributestype is a representation of the C type WindowAttributesType.
type Windowattributestype int

const (
	Windowattributestype_Title    Windowattributestype = 2
	Windowattributestype_X        Windowattributestype = 4
	Windowattributestype_Y        Windowattributestype = 8
	Windowattributestype_Cursor   Windowattributestype = 16
	Windowattributestype_Visual   Windowattributestype = 32
	Windowattributestype_Wmclass  Windowattributestype = 64
	Windowattributestype_Noredir  Windowattributestype = 128
	Windowattributestype_TypeHint Windowattributestype = 256
)

// Windowhints is a representation of the C type WindowHints.
type Windowhints int

const (
	Windowhints_Pos        Windowhints = 1
	Windowhints_MinSize    Windowhints = 2
	Windowhints_MaxSize    Windowhints = 4
	Windowhints_BaseSize   Windowhints = 8
	Windowhints_Aspect     Windowhints = 16
	Windowhints_ResizeInc  Windowhints = 32
	Windowhints_WinGravity Windowhints = 64
	Windowhints_UserPos    Windowhints = 128
	Windowhints_UserSize   Windowhints = 256
)

// Windowstate is a representation of the C type WindowState.
type Windowstate int

const (
	Windowstate_Withdrawn       Windowstate = 1
	Windowstate_Iconified       Windowstate = 2
	Windowstate_Maximized       Windowstate = 4
	Windowstate_Sticky          Windowstate = 8
	Windowstate_Fullscreen      Windowstate = 16
	Windowstate_Above           Windowstate = 32
	Windowstate_Below           Windowstate = 64
	Windowstate_Focused         Windowstate = 128
	Windowstate_Tiled           Windowstate = 256
	Windowstate_TopTiled        Windowstate = 512
	Windowstate_TopResizable    Windowstate = 1024
	Windowstate_RightTiled      Windowstate = 2048
	Windowstate_RightResizable  Windowstate = 4096
	Windowstate_BottomTiled     Windowstate = 8192
	Windowstate_BottomResizable Windowstate = 16384
	Windowstate_LeftTiled       Windowstate = 32768
	Windowstate_LeftResizable   Windowstate = 65536
)
