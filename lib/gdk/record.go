// Code generated - DO NOT EDIT.

package gdk

type Atom struct {
	native uintptr
}

type Color struct {
	native uintptr
	Pixel  uint32
	Red    uint16
	Green  uint16
	Blue   uint16
}

type DevicePadInterface struct {
	native uintptr
}

type DrawingContextClass struct {
	native uintptr
}

type EventAny struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
}

type EventButton struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'axes' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'

	Button uint32
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'

	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'

}

type EventConfigure struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	X         int32
	Y         int32
	Width     int32
	Height    int32
}

type EventCrossing struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	// UNSUPPORTED : C value 'subwindow' : no Go type for 'Window'

	Time uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'mode' : no Go type for 'CrossingMode'

	// UNSUPPORTED : C value 'detail' : no Go type for 'NotifyType'

	// UNSUPPORTED : C value 'focus' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'

}

type EventDND struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	// UNSUPPORTED : C value 'context' : no Go type for 'DragContext'

	Time  uint32
	XRoot int16
	YRoot int16
}

type EventExpose struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Area      *Rectangle
	// UNSUPPORTED : C value 'region' : no Go type for 'cairo.Region'

	Count int32
}

type EventFocus struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	In        int16
}

type EventGrabBroken struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	// UNSUPPORTED : C value 'keyboard' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'implicit' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'grab_window' : no Go type for 'Window'

}

type EventKey struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'

	Keyval          uint32
	Length          int32
	String          string
	HardwareKeycode uint16
	Group           uint8
	IsModifier      uint32
}

type EventMotion struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'axes' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'

	IsHint int16
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'

	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'

}

type EventOwnerChange struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	// UNSUPPORTED : C value 'owner' : no Go type for 'Window'

	// UNSUPPORTED : C value 'reason' : no Go type for 'OwnerChange'

	Selection     *Atom
	Time          uint32
	SelectionTime uint32
}

type EventPadAxis struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Time      uint32
	Group     uint32
	Index     uint32
	Mode      uint32
	// UNSUPPORTED : C value 'value' : no Go type for 'gdouble'

}

type EventPadButton struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Time      uint32
	Group     uint32
	Button    uint32
	Mode      uint32
}

type EventPadGroupMode struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Time      uint32
	Group     uint32
	Mode      uint32
}

type EventProperty struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Atom      *Atom
	Time      uint32
	// UNSUPPORTED : C value 'state' : no Go type for 'PropertyState'

}

type EventProximity struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'device' : no Go type for 'Device'

}

type EventScroll struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'

	// UNSUPPORTED : C value 'direction' : no Go type for 'ScrollDirection'

	// UNSUPPORTED : C value 'device' : no Go type for 'Device'

	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'delta_x' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'delta_y' : no Go type for 'gdouble'

	IsStop uint32
}

type EventSelection struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Selection *Atom
	Target    *Atom
	Property  *Atom
	Time      uint32
	// UNSUPPORTED : C value 'requestor' : no Go type for 'Window'

}

type EventSequence struct {
	native uintptr
}

type EventSetting struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	// UNSUPPORTED : C value 'action' : no Go type for 'SettingAction'

	Name string
}

type EventTouch struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'axes' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'

	Sequence *EventSequence
	// UNSUPPORTED : C value 'emulating_pointer' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'device' : no Go type for 'Device'

	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'

}

type EventTouchpadPinch struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Phase     int8
	NFingers  int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'dx' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'dy' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'angle_delta' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'scale' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'

}

type EventTouchpadSwipe struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	Phase     int8
	NFingers  int8
	Time      uint32
	// UNSUPPORTED : C value 'x' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'dx' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'dy' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'x_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'y_root' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'state' : no Go type for 'ModifierType'

}

type EventVisibility struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	// UNSUPPORTED : C value 'state' : no Go type for 'VisibilityState'

}

type EventWindowState struct {
	native uintptr
	// UNSUPPORTED : C value 'type' : no Go type for 'EventType'

	// UNSUPPORTED : C value 'window' : no Go type for 'Window'

	SendEvent int8
	// UNSUPPORTED : C value 'changed_mask' : no Go type for 'WindowState'

	// UNSUPPORTED : C value 'new_window_state' : no Go type for 'WindowState'

}

type FrameClockClass struct {
	native uintptr
}

type FrameClockPrivate struct {
	native uintptr
}

type FrameTimings struct {
	native uintptr
}

type Geometry struct {
	native     uintptr
	MinWidth   int32
	MinHeight  int32
	MaxWidth   int32
	MaxHeight  int32
	BaseWidth  int32
	BaseHeight int32
	WidthInc   int32
	HeightInc  int32
	// UNSUPPORTED : C value 'min_aspect' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'max_aspect' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'win_gravity' : no Go type for 'Gravity'

}

type KeymapKey struct {
	native  uintptr
	Keycode uint32
	Group   int32
	Level   int32
}

type MonitorClass struct {
	native uintptr
}

type Point struct {
	native uintptr
	X      int32
	Y      int32
}

type RGBA struct {
	native uintptr
	// UNSUPPORTED : C value 'red' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'green' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'blue' : no Go type for 'gdouble'

	// UNSUPPORTED : C value 'alpha' : no Go type for 'gdouble'

}

type Rectangle struct {
	native uintptr
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type TimeCoord struct {
	native uintptr
	Time   uint32
	// UNSUPPORTED : C value 'axes' : missing Type

}

type WindowAttr struct {
	native    uintptr
	Title     string
	EventMask int32
	X         int32
	Y         int32
	Width     int32
	Height    int32
	// UNSUPPORTED : C value 'wclass' : no Go type for 'WindowWindowClass'

	// UNSUPPORTED : C value 'visual' : no Go type for 'Visual'

	// UNSUPPORTED : C value 'window_type' : no Go type for 'WindowType'

	// UNSUPPORTED : C value 'cursor' : no Go type for 'Cursor'

	WmclassName  string
	WmclassClass string
	// UNSUPPORTED : C value 'override_redirect' : no Go type for 'gboolean'

	// UNSUPPORTED : C value 'type_hint' : no Go type for 'WindowTypeHint'

}

type WindowClass struct {
	native uintptr
	// UNSUPPORTED : C value 'parent_class' : no Go type for 'GObject.ObjectClass'

	// UNSUPPORTED : C value 'pick_embedded_child' : missing Type

	// UNSUPPORTED : C value 'to_embedder' : missing Type

	// UNSUPPORTED : C value 'from_embedder' : missing Type

	// UNSUPPORTED : C value 'create_surface' : missing Type

	// UNSUPPORTED : C value '_gdk_reserved1' : missing Type

	// UNSUPPORTED : C value '_gdk_reserved2' : missing Type

	// UNSUPPORTED : C value '_gdk_reserved3' : missing Type

	// UNSUPPORTED : C value '_gdk_reserved4' : missing Type

	// UNSUPPORTED : C value '_gdk_reserved5' : missing Type

	// UNSUPPORTED : C value '_gdk_reserved6' : missing Type

	// UNSUPPORTED : C value '_gdk_reserved7' : missing Type

	// UNSUPPORTED : C value '_gdk_reserved8' : missing Type

}

type WindowRedirect struct {
	native uintptr
}
