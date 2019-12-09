// Code generated - DO NOT EDIT.

package atk

import (
	callback "github.com/pekim/gobbi/internal/cgo/callback"
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

var actionInterface *gi.Interface
var actionInterface_Once sync.Once

func actionInterface_Set() error {
	var err error
	actionInterface_Once.Do(func() {
		actionInterface, err = gi.InterfaceNew("Atk", "Action")
	})
	return err
}

type Action struct {
	native unsafe.Pointer
}

func ActionNewFromNative(native unsafe.Pointer) *Action {
	instance := &Action{native: native}

	return instance
}

/*
CastToAction down casts any arbitrary Object to Action.
Exercise care, as this is a potentially dangerous function
if the Object is not a Action.
*/
func CastToAction(object *gobject.Object) *Action {
	return ActionNewFromNative(object.Native())
}

// Equals compares this Action with another Action, and returns true if they represent the same Object.
func (recv *Action) Equals(other *Action) bool {
	return other.Native() == recv.Native()
}

func (recv *Action) Native() unsafe.Pointer {
	return recv.native
}

var actionDoActionFunction *gi.Function
var actionDoActionFunction_Once sync.Once

func actionDoActionFunction_Set() error {
	var err error
	actionDoActionFunction_Once.Do(func() {
		err = actionInterface_Set()
		if err != nil {
			return
		}
		actionDoActionFunction, err = actionInterface.InvokerNew("do_action")
	})
	return err
}

// DoAction is a representation of the C type atk_action_do_action.
func (recv *Action) DoAction(i int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := actionDoActionFunction_Set()
	if err == nil {
		ret = actionDoActionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var actionGetDescriptionFunction *gi.Function
var actionGetDescriptionFunction_Once sync.Once

func actionGetDescriptionFunction_Set() error {
	var err error
	actionGetDescriptionFunction_Once.Do(func() {
		err = actionInterface_Set()
		if err != nil {
			return
		}
		actionGetDescriptionFunction, err = actionInterface.InvokerNew("get_description")
	})
	return err
}

// GetDescription is a representation of the C type atk_action_get_description.
func (recv *Action) GetDescription(i int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := actionGetDescriptionFunction_Set()
	if err == nil {
		ret = actionGetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var actionGetKeybindingFunction *gi.Function
var actionGetKeybindingFunction_Once sync.Once

func actionGetKeybindingFunction_Set() error {
	var err error
	actionGetKeybindingFunction_Once.Do(func() {
		err = actionInterface_Set()
		if err != nil {
			return
		}
		actionGetKeybindingFunction, err = actionInterface.InvokerNew("get_keybinding")
	})
	return err
}

// GetKeybinding is a representation of the C type atk_action_get_keybinding.
func (recv *Action) GetKeybinding(i int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := actionGetKeybindingFunction_Set()
	if err == nil {
		ret = actionGetKeybindingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var actionGetLocalizedNameFunction *gi.Function
var actionGetLocalizedNameFunction_Once sync.Once

func actionGetLocalizedNameFunction_Set() error {
	var err error
	actionGetLocalizedNameFunction_Once.Do(func() {
		err = actionInterface_Set()
		if err != nil {
			return
		}
		actionGetLocalizedNameFunction, err = actionInterface.InvokerNew("get_localized_name")
	})
	return err
}

// GetLocalizedName is a representation of the C type atk_action_get_localized_name.
func (recv *Action) GetLocalizedName(i int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := actionGetLocalizedNameFunction_Set()
	if err == nil {
		ret = actionGetLocalizedNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var actionGetNActionsFunction *gi.Function
var actionGetNActionsFunction_Once sync.Once

func actionGetNActionsFunction_Set() error {
	var err error
	actionGetNActionsFunction_Once.Do(func() {
		err = actionInterface_Set()
		if err != nil {
			return
		}
		actionGetNActionsFunction, err = actionInterface.InvokerNew("get_n_actions")
	})
	return err
}

// GetNActions is a representation of the C type atk_action_get_n_actions.
func (recv *Action) GetNActions() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := actionGetNActionsFunction_Set()
	if err == nil {
		ret = actionGetNActionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var actionGetNameFunction *gi.Function
var actionGetNameFunction_Once sync.Once

func actionGetNameFunction_Set() error {
	var err error
	actionGetNameFunction_Once.Do(func() {
		err = actionInterface_Set()
		if err != nil {
			return
		}
		actionGetNameFunction, err = actionInterface.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type atk_action_get_name.
func (recv *Action) GetName(i int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := actionGetNameFunction_Set()
	if err == nil {
		ret = actionGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var actionSetDescriptionFunction *gi.Function
var actionSetDescriptionFunction_Once sync.Once

func actionSetDescriptionFunction_Set() error {
	var err error
	actionSetDescriptionFunction_Once.Do(func() {
		err = actionInterface_Set()
		if err != nil {
			return
		}
		actionSetDescriptionFunction, err = actionInterface.InvokerNew("set_description")
	})
	return err
}

// SetDescription is a representation of the C type atk_action_set_description.
func (recv *Action) SetDescription(i int32, desc string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)
	inArgs[2].SetString(desc)

	var ret gi.Argument

	err := actionSetDescriptionFunction_Set()
	if err == nil {
		ret = actionSetDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var componentInterface *gi.Interface
var componentInterface_Once sync.Once

func componentInterface_Set() error {
	var err error
	componentInterface_Once.Do(func() {
		componentInterface, err = gi.InterfaceNew("Atk", "Component")
	})
	return err
}

type Component struct {
	native unsafe.Pointer
}

func ComponentNewFromNative(native unsafe.Pointer) *Component {
	instance := &Component{native: native}

	return instance
}

/*
CastToComponent down casts any arbitrary Object to Component.
Exercise care, as this is a potentially dangerous function
if the Object is not a Component.
*/
func CastToComponent(object *gobject.Object) *Component {
	return ComponentNewFromNative(object.Native())
}

// Equals compares this Component with another Component, and returns true if they represent the same Object.
func (recv *Component) Equals(other *Component) bool {
	return other.Native() == recv.Native()
}

func (recv *Component) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'atk_component_add_focus_handler' : parameter 'handler' of type 'FocusHandler' not supported

var componentContainsFunction *gi.Function
var componentContainsFunction_Once sync.Once

func componentContainsFunction_Set() error {
	var err error
	componentContainsFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentContainsFunction, err = componentInterface.InvokerNew("contains")
	})
	return err
}

// Contains is a representation of the C type atk_component_contains.
func (recv *Component) Contains(x int32, y int32, coordType CoordType) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)
	inArgs[3].SetInt32(int32(coordType))

	var ret gi.Argument

	err := componentContainsFunction_Set()
	if err == nil {
		ret = componentContainsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var componentGetAlphaFunction *gi.Function
var componentGetAlphaFunction_Once sync.Once

func componentGetAlphaFunction_Set() error {
	var err error
	componentGetAlphaFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentGetAlphaFunction, err = componentInterface.InvokerNew("get_alpha")
	})
	return err
}

// GetAlpha is a representation of the C type atk_component_get_alpha.
func (recv *Component) GetAlpha() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := componentGetAlphaFunction_Set()
	if err == nil {
		ret = componentGetAlphaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var componentGetExtentsFunction *gi.Function
var componentGetExtentsFunction_Once sync.Once

func componentGetExtentsFunction_Set() error {
	var err error
	componentGetExtentsFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentGetExtentsFunction, err = componentInterface.InvokerNew("get_extents")
	})
	return err
}

// GetExtents is a representation of the C type atk_component_get_extents.
func (recv *Component) GetExtents(coordType CoordType) (int32, int32, int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(coordType))

	var outArgs [4]gi.Argument

	err := componentGetExtentsFunction_Set()
	if err == nil {
		componentGetExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()
	out3 := outArgs[3].Int32()

	return out0, out1, out2, out3
}

var componentGetLayerFunction *gi.Function
var componentGetLayerFunction_Once sync.Once

func componentGetLayerFunction_Set() error {
	var err error
	componentGetLayerFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentGetLayerFunction, err = componentInterface.InvokerNew("get_layer")
	})
	return err
}

// GetLayer is a representation of the C type atk_component_get_layer.
func (recv *Component) GetLayer() Layer {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := componentGetLayerFunction_Set()
	if err == nil {
		ret = componentGetLayerFunction.Invoke(inArgs[:], nil)
	}

	retGo := Layer(ret.Int32())

	return retGo
}

var componentGetMdiZorderFunction *gi.Function
var componentGetMdiZorderFunction_Once sync.Once

func componentGetMdiZorderFunction_Set() error {
	var err error
	componentGetMdiZorderFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentGetMdiZorderFunction, err = componentInterface.InvokerNew("get_mdi_zorder")
	})
	return err
}

// GetMdiZorder is a representation of the C type atk_component_get_mdi_zorder.
func (recv *Component) GetMdiZorder() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := componentGetMdiZorderFunction_Set()
	if err == nil {
		ret = componentGetMdiZorderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var componentGetPositionFunction *gi.Function
var componentGetPositionFunction_Once sync.Once

func componentGetPositionFunction_Set() error {
	var err error
	componentGetPositionFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentGetPositionFunction, err = componentInterface.InvokerNew("get_position")
	})
	return err
}

// GetPosition is a representation of the C type atk_component_get_position.
func (recv *Component) GetPosition(coordType CoordType) (int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(coordType))

	var outArgs [2]gi.Argument

	err := componentGetPositionFunction_Set()
	if err == nil {
		componentGetPositionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var componentGetSizeFunction *gi.Function
var componentGetSizeFunction_Once sync.Once

func componentGetSizeFunction_Set() error {
	var err error
	componentGetSizeFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentGetSizeFunction, err = componentInterface.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type atk_component_get_size.
func (recv *Component) GetSize() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument

	err := componentGetSizeFunction_Set()
	if err == nil {
		componentGetSizeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var componentGrabFocusFunction *gi.Function
var componentGrabFocusFunction_Once sync.Once

func componentGrabFocusFunction_Set() error {
	var err error
	componentGrabFocusFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentGrabFocusFunction, err = componentInterface.InvokerNew("grab_focus")
	})
	return err
}

// GrabFocus is a representation of the C type atk_component_grab_focus.
func (recv *Component) GrabFocus() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := componentGrabFocusFunction_Set()
	if err == nil {
		ret = componentGrabFocusFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var componentRefAccessibleAtPointFunction *gi.Function
var componentRefAccessibleAtPointFunction_Once sync.Once

func componentRefAccessibleAtPointFunction_Set() error {
	var err error
	componentRefAccessibleAtPointFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentRefAccessibleAtPointFunction, err = componentInterface.InvokerNew("ref_accessible_at_point")
	})
	return err
}

// RefAccessibleAtPoint is a representation of the C type atk_component_ref_accessible_at_point.
func (recv *Component) RefAccessibleAtPoint(x int32, y int32, coordType CoordType) *Object {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)
	inArgs[3].SetInt32(int32(coordType))

	var ret gi.Argument

	err := componentRefAccessibleAtPointFunction_Set()
	if err == nil {
		ret = componentRefAccessibleAtPointFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var componentRemoveFocusHandlerFunction *gi.Function
var componentRemoveFocusHandlerFunction_Once sync.Once

func componentRemoveFocusHandlerFunction_Set() error {
	var err error
	componentRemoveFocusHandlerFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentRemoveFocusHandlerFunction, err = componentInterface.InvokerNew("remove_focus_handler")
	})
	return err
}

// RemoveFocusHandler is a representation of the C type atk_component_remove_focus_handler.
func (recv *Component) RemoveFocusHandler(handlerId uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetUint32(handlerId)

	err := componentRemoveFocusHandlerFunction_Set()
	if err == nil {
		componentRemoveFocusHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var componentScrollToFunction *gi.Function
var componentScrollToFunction_Once sync.Once

func componentScrollToFunction_Set() error {
	var err error
	componentScrollToFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentScrollToFunction, err = componentInterface.InvokerNew("scroll_to")
	})
	return err
}

// ScrollTo is a representation of the C type atk_component_scroll_to.
func (recv *Component) ScrollTo(type_ ScrollType) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(type_))

	var ret gi.Argument

	err := componentScrollToFunction_Set()
	if err == nil {
		ret = componentScrollToFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var componentScrollToPointFunction *gi.Function
var componentScrollToPointFunction_Once sync.Once

func componentScrollToPointFunction_Set() error {
	var err error
	componentScrollToPointFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentScrollToPointFunction, err = componentInterface.InvokerNew("scroll_to_point")
	})
	return err
}

// ScrollToPoint is a representation of the C type atk_component_scroll_to_point.
func (recv *Component) ScrollToPoint(coords CoordType, x int32, y int32) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(coords))
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	var ret gi.Argument

	err := componentScrollToPointFunction_Set()
	if err == nil {
		ret = componentScrollToPointFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var componentSetExtentsFunction *gi.Function
var componentSetExtentsFunction_Once sync.Once

func componentSetExtentsFunction_Set() error {
	var err error
	componentSetExtentsFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentSetExtentsFunction, err = componentInterface.InvokerNew("set_extents")
	})
	return err
}

// SetExtents is a representation of the C type atk_component_set_extents.
func (recv *Component) SetExtents(x int32, y int32, width int32, height int32, coordType CoordType) bool {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)
	inArgs[3].SetInt32(width)
	inArgs[4].SetInt32(height)
	inArgs[5].SetInt32(int32(coordType))

	var ret gi.Argument

	err := componentSetExtentsFunction_Set()
	if err == nil {
		ret = componentSetExtentsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var componentSetPositionFunction *gi.Function
var componentSetPositionFunction_Once sync.Once

func componentSetPositionFunction_Set() error {
	var err error
	componentSetPositionFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentSetPositionFunction, err = componentInterface.InvokerNew("set_position")
	})
	return err
}

// SetPosition is a representation of the C type atk_component_set_position.
func (recv *Component) SetPosition(x int32, y int32, coordType CoordType) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)
	inArgs[3].SetInt32(int32(coordType))

	var ret gi.Argument

	err := componentSetPositionFunction_Set()
	if err == nil {
		ret = componentSetPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var componentSetSizeFunction *gi.Function
var componentSetSizeFunction_Once sync.Once

func componentSetSizeFunction_Set() error {
	var err error
	componentSetSizeFunction_Once.Do(func() {
		err = componentInterface_Set()
		if err != nil {
			return
		}
		componentSetSizeFunction, err = componentInterface.InvokerNew("set_size")
	})
	return err
}

// SetSize is a representation of the C type atk_component_set_size.
func (recv *Component) SetSize(width int32, height int32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(width)
	inArgs[2].SetInt32(height)

	var ret gi.Argument

	err := componentSetSizeFunction_Set()
	if err == nil {
		ret = componentSetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

/*
ConnectBoundsChanged connects a callback to the 'bounds-changed' signal of the Component.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Component) ConnectBoundsChanged(handler func(instance *Component, arg1 *Rectangle)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "bounds-changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Component) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var documentInterface *gi.Interface
var documentInterface_Once sync.Once

func documentInterface_Set() error {
	var err error
	documentInterface_Once.Do(func() {
		documentInterface, err = gi.InterfaceNew("Atk", "Document")
	})
	return err
}

type Document struct {
	native unsafe.Pointer
}

func DocumentNewFromNative(native unsafe.Pointer) *Document {
	instance := &Document{native: native}

	return instance
}

/*
CastToDocument down casts any arbitrary Object to Document.
Exercise care, as this is a potentially dangerous function
if the Object is not a Document.
*/
func CastToDocument(object *gobject.Object) *Document {
	return DocumentNewFromNative(object.Native())
}

// Equals compares this Document with another Document, and returns true if they represent the same Object.
func (recv *Document) Equals(other *Document) bool {
	return other.Native() == recv.Native()
}

func (recv *Document) Native() unsafe.Pointer {
	return recv.native
}

var documentGetAttributeValueFunction *gi.Function
var documentGetAttributeValueFunction_Once sync.Once

func documentGetAttributeValueFunction_Set() error {
	var err error
	documentGetAttributeValueFunction_Once.Do(func() {
		err = documentInterface_Set()
		if err != nil {
			return
		}
		documentGetAttributeValueFunction, err = documentInterface.InvokerNew("get_attribute_value")
	})
	return err
}

// GetAttributeValue is a representation of the C type atk_document_get_attribute_value.
func (recv *Document) GetAttributeValue(attributeName string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attributeName)

	var ret gi.Argument

	err := documentGetAttributeValueFunction_Set()
	if err == nil {
		ret = documentGetAttributeValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'atk_document_get_attributes' : return type 'AttributeSet' not supported

var documentGetCurrentPageNumberFunction *gi.Function
var documentGetCurrentPageNumberFunction_Once sync.Once

func documentGetCurrentPageNumberFunction_Set() error {
	var err error
	documentGetCurrentPageNumberFunction_Once.Do(func() {
		err = documentInterface_Set()
		if err != nil {
			return
		}
		documentGetCurrentPageNumberFunction, err = documentInterface.InvokerNew("get_current_page_number")
	})
	return err
}

// GetCurrentPageNumber is a representation of the C type atk_document_get_current_page_number.
func (recv *Document) GetCurrentPageNumber() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := documentGetCurrentPageNumberFunction_Set()
	if err == nil {
		ret = documentGetCurrentPageNumberFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'atk_document_get_document' : return type 'gpointer' not supported

var documentGetDocumentTypeFunction *gi.Function
var documentGetDocumentTypeFunction_Once sync.Once

func documentGetDocumentTypeFunction_Set() error {
	var err error
	documentGetDocumentTypeFunction_Once.Do(func() {
		err = documentInterface_Set()
		if err != nil {
			return
		}
		documentGetDocumentTypeFunction, err = documentInterface.InvokerNew("get_document_type")
	})
	return err
}

// GetDocumentType is a representation of the C type atk_document_get_document_type.
func (recv *Document) GetDocumentType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := documentGetDocumentTypeFunction_Set()
	if err == nil {
		ret = documentGetDocumentTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var documentGetLocaleFunction *gi.Function
var documentGetLocaleFunction_Once sync.Once

func documentGetLocaleFunction_Set() error {
	var err error
	documentGetLocaleFunction_Once.Do(func() {
		err = documentInterface_Set()
		if err != nil {
			return
		}
		documentGetLocaleFunction, err = documentInterface.InvokerNew("get_locale")
	})
	return err
}

// GetLocale is a representation of the C type atk_document_get_locale.
func (recv *Document) GetLocale() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := documentGetLocaleFunction_Set()
	if err == nil {
		ret = documentGetLocaleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var documentGetPageCountFunction *gi.Function
var documentGetPageCountFunction_Once sync.Once

func documentGetPageCountFunction_Set() error {
	var err error
	documentGetPageCountFunction_Once.Do(func() {
		err = documentInterface_Set()
		if err != nil {
			return
		}
		documentGetPageCountFunction, err = documentInterface.InvokerNew("get_page_count")
	})
	return err
}

// GetPageCount is a representation of the C type atk_document_get_page_count.
func (recv *Document) GetPageCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := documentGetPageCountFunction_Set()
	if err == nil {
		ret = documentGetPageCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var documentSetAttributeValueFunction *gi.Function
var documentSetAttributeValueFunction_Once sync.Once

func documentSetAttributeValueFunction_Set() error {
	var err error
	documentSetAttributeValueFunction_Once.Do(func() {
		err = documentInterface_Set()
		if err != nil {
			return
		}
		documentSetAttributeValueFunction, err = documentInterface.InvokerNew("set_attribute_value")
	})
	return err
}

// SetAttributeValue is a representation of the C type atk_document_set_attribute_value.
func (recv *Document) SetAttributeValue(attributeName string, attributeValue string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(attributeName)
	inArgs[2].SetString(attributeValue)

	var ret gi.Argument

	err := documentSetAttributeValueFunction_Set()
	if err == nil {
		ret = documentSetAttributeValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

/*
ConnectLoadComplete connects a callback to the 'load-complete' signal of the Document.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Document) ConnectLoadComplete(handler func(instance *Document)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "load-complete", marshal)
}

/*
ConnectLoadStopped connects a callback to the 'load-stopped' signal of the Document.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Document) ConnectLoadStopped(handler func(instance *Document)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "load-stopped", marshal)
}

/*
ConnectPageChanged connects a callback to the 'page-changed' signal of the Document.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Document) ConnectPageChanged(handler func(instance *Document, pageNumber int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "page-changed", marshal)
}

/*
ConnectReload connects a callback to the 'reload' signal of the Document.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Document) ConnectReload(handler func(instance *Document)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "reload", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Document) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var editableTextInterface *gi.Interface
var editableTextInterface_Once sync.Once

func editableTextInterface_Set() error {
	var err error
	editableTextInterface_Once.Do(func() {
		editableTextInterface, err = gi.InterfaceNew("Atk", "EditableText")
	})
	return err
}

type EditableText struct {
	native unsafe.Pointer
}

func EditableTextNewFromNative(native unsafe.Pointer) *EditableText {
	instance := &EditableText{native: native}

	return instance
}

/*
CastToEditableText down casts any arbitrary Object to EditableText.
Exercise care, as this is a potentially dangerous function
if the Object is not a EditableText.
*/
func CastToEditableText(object *gobject.Object) *EditableText {
	return EditableTextNewFromNative(object.Native())
}

// Equals compares this EditableText with another EditableText, and returns true if they represent the same Object.
func (recv *EditableText) Equals(other *EditableText) bool {
	return other.Native() == recv.Native()
}

func (recv *EditableText) Native() unsafe.Pointer {
	return recv.native
}

var editableTextCopyTextFunction *gi.Function
var editableTextCopyTextFunction_Once sync.Once

func editableTextCopyTextFunction_Set() error {
	var err error
	editableTextCopyTextFunction_Once.Do(func() {
		err = editableTextInterface_Set()
		if err != nil {
			return
		}
		editableTextCopyTextFunction, err = editableTextInterface.InvokerNew("copy_text")
	})
	return err
}

// CopyText is a representation of the C type atk_editable_text_copy_text.
func (recv *EditableText) CopyText(startPos int32, endPos int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startPos)
	inArgs[2].SetInt32(endPos)

	err := editableTextCopyTextFunction_Set()
	if err == nil {
		editableTextCopyTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableTextCutTextFunction *gi.Function
var editableTextCutTextFunction_Once sync.Once

func editableTextCutTextFunction_Set() error {
	var err error
	editableTextCutTextFunction_Once.Do(func() {
		err = editableTextInterface_Set()
		if err != nil {
			return
		}
		editableTextCutTextFunction, err = editableTextInterface.InvokerNew("cut_text")
	})
	return err
}

// CutText is a representation of the C type atk_editable_text_cut_text.
func (recv *EditableText) CutText(startPos int32, endPos int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startPos)
	inArgs[2].SetInt32(endPos)

	err := editableTextCutTextFunction_Set()
	if err == nil {
		editableTextCutTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableTextDeleteTextFunction *gi.Function
var editableTextDeleteTextFunction_Once sync.Once

func editableTextDeleteTextFunction_Set() error {
	var err error
	editableTextDeleteTextFunction_Once.Do(func() {
		err = editableTextInterface_Set()
		if err != nil {
			return
		}
		editableTextDeleteTextFunction, err = editableTextInterface.InvokerNew("delete_text")
	})
	return err
}

// DeleteText is a representation of the C type atk_editable_text_delete_text.
func (recv *EditableText) DeleteText(startPos int32, endPos int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startPos)
	inArgs[2].SetInt32(endPos)

	err := editableTextDeleteTextFunction_Set()
	if err == nil {
		editableTextDeleteTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableTextInsertTextFunction *gi.Function
var editableTextInsertTextFunction_Once sync.Once

func editableTextInsertTextFunction_Set() error {
	var err error
	editableTextInsertTextFunction_Once.Do(func() {
		err = editableTextInterface_Set()
		if err != nil {
			return
		}
		editableTextInsertTextFunction, err = editableTextInterface.InvokerNew("insert_text")
	})
	return err
}

// InsertText is a representation of the C type atk_editable_text_insert_text.
func (recv *EditableText) InsertText(string_ string, length int32, position int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(string_)
	inArgs[2].SetInt32(length)
	inArgs[3].SetInt32(position)

	err := editableTextInsertTextFunction_Set()
	if err == nil {
		editableTextInsertTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableTextPasteTextFunction *gi.Function
var editableTextPasteTextFunction_Once sync.Once

func editableTextPasteTextFunction_Set() error {
	var err error
	editableTextPasteTextFunction_Once.Do(func() {
		err = editableTextInterface_Set()
		if err != nil {
			return
		}
		editableTextPasteTextFunction, err = editableTextInterface.InvokerNew("paste_text")
	})
	return err
}

// PasteText is a representation of the C type atk_editable_text_paste_text.
func (recv *EditableText) PasteText(position int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(position)

	err := editableTextPasteTextFunction_Set()
	if err == nil {
		editableTextPasteTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'atk_editable_text_set_run_attributes' : parameter 'attrib_set' of type 'AttributeSet' not supported

var editableTextSetTextContentsFunction *gi.Function
var editableTextSetTextContentsFunction_Once sync.Once

func editableTextSetTextContentsFunction_Set() error {
	var err error
	editableTextSetTextContentsFunction_Once.Do(func() {
		err = editableTextInterface_Set()
		if err != nil {
			return
		}
		editableTextSetTextContentsFunction, err = editableTextInterface.InvokerNew("set_text_contents")
	})
	return err
}

// SetTextContents is a representation of the C type atk_editable_text_set_text_contents.
func (recv *EditableText) SetTextContents(string_ string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(string_)

	err := editableTextSetTextContentsFunction_Set()
	if err == nil {
		editableTextSetTextContentsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var hyperlinkImplInterface *gi.Interface
var hyperlinkImplInterface_Once sync.Once

func hyperlinkImplInterface_Set() error {
	var err error
	hyperlinkImplInterface_Once.Do(func() {
		hyperlinkImplInterface, err = gi.InterfaceNew("Atk", "HyperlinkImpl")
	})
	return err
}

type HyperlinkImpl struct {
	native unsafe.Pointer
}

func HyperlinkImplNewFromNative(native unsafe.Pointer) *HyperlinkImpl {
	instance := &HyperlinkImpl{native: native}

	return instance
}

/*
CastToHyperlinkImpl down casts any arbitrary Object to HyperlinkImpl.
Exercise care, as this is a potentially dangerous function
if the Object is not a HyperlinkImpl.
*/
func CastToHyperlinkImpl(object *gobject.Object) *HyperlinkImpl {
	return HyperlinkImplNewFromNative(object.Native())
}

// Equals compares this HyperlinkImpl with another HyperlinkImpl, and returns true if they represent the same Object.
func (recv *HyperlinkImpl) Equals(other *HyperlinkImpl) bool {
	return other.Native() == recv.Native()
}

func (recv *HyperlinkImpl) Native() unsafe.Pointer {
	return recv.native
}

var hyperlinkImplGetHyperlinkFunction *gi.Function
var hyperlinkImplGetHyperlinkFunction_Once sync.Once

func hyperlinkImplGetHyperlinkFunction_Set() error {
	var err error
	hyperlinkImplGetHyperlinkFunction_Once.Do(func() {
		err = hyperlinkImplInterface_Set()
		if err != nil {
			return
		}
		hyperlinkImplGetHyperlinkFunction, err = hyperlinkImplInterface.InvokerNew("get_hyperlink")
	})
	return err
}

// GetHyperlink is a representation of the C type atk_hyperlink_impl_get_hyperlink.
func (recv *HyperlinkImpl) GetHyperlink() *Hyperlink {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hyperlinkImplGetHyperlinkFunction_Set()
	if err == nil {
		ret = hyperlinkImplGetHyperlinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := HyperlinkNewFromNative(ret.Pointer())

	return retGo
}

var hypertextInterface *gi.Interface
var hypertextInterface_Once sync.Once

func hypertextInterface_Set() error {
	var err error
	hypertextInterface_Once.Do(func() {
		hypertextInterface, err = gi.InterfaceNew("Atk", "Hypertext")
	})
	return err
}

type Hypertext struct {
	native unsafe.Pointer
}

func HypertextNewFromNative(native unsafe.Pointer) *Hypertext {
	instance := &Hypertext{native: native}

	return instance
}

/*
CastToHypertext down casts any arbitrary Object to Hypertext.
Exercise care, as this is a potentially dangerous function
if the Object is not a Hypertext.
*/
func CastToHypertext(object *gobject.Object) *Hypertext {
	return HypertextNewFromNative(object.Native())
}

// Equals compares this Hypertext with another Hypertext, and returns true if they represent the same Object.
func (recv *Hypertext) Equals(other *Hypertext) bool {
	return other.Native() == recv.Native()
}

func (recv *Hypertext) Native() unsafe.Pointer {
	return recv.native
}

var hypertextGetLinkFunction *gi.Function
var hypertextGetLinkFunction_Once sync.Once

func hypertextGetLinkFunction_Set() error {
	var err error
	hypertextGetLinkFunction_Once.Do(func() {
		err = hypertextInterface_Set()
		if err != nil {
			return
		}
		hypertextGetLinkFunction, err = hypertextInterface.InvokerNew("get_link")
	})
	return err
}

// GetLink is a representation of the C type atk_hypertext_get_link.
func (recv *Hypertext) GetLink(linkIndex int32) *Hyperlink {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(linkIndex)

	var ret gi.Argument

	err := hypertextGetLinkFunction_Set()
	if err == nil {
		ret = hypertextGetLinkFunction.Invoke(inArgs[:], nil)
	}

	retGo := HyperlinkNewFromNative(ret.Pointer())

	return retGo
}

var hypertextGetLinkIndexFunction *gi.Function
var hypertextGetLinkIndexFunction_Once sync.Once

func hypertextGetLinkIndexFunction_Set() error {
	var err error
	hypertextGetLinkIndexFunction_Once.Do(func() {
		err = hypertextInterface_Set()
		if err != nil {
			return
		}
		hypertextGetLinkIndexFunction, err = hypertextInterface.InvokerNew("get_link_index")
	})
	return err
}

// GetLinkIndex is a representation of the C type atk_hypertext_get_link_index.
func (recv *Hypertext) GetLinkIndex(charIndex int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(charIndex)

	var ret gi.Argument

	err := hypertextGetLinkIndexFunction_Set()
	if err == nil {
		ret = hypertextGetLinkIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var hypertextGetNLinksFunction *gi.Function
var hypertextGetNLinksFunction_Once sync.Once

func hypertextGetNLinksFunction_Set() error {
	var err error
	hypertextGetNLinksFunction_Once.Do(func() {
		err = hypertextInterface_Set()
		if err != nil {
			return
		}
		hypertextGetNLinksFunction, err = hypertextInterface.InvokerNew("get_n_links")
	})
	return err
}

// GetNLinks is a representation of the C type atk_hypertext_get_n_links.
func (recv *Hypertext) GetNLinks() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := hypertextGetNLinksFunction_Set()
	if err == nil {
		ret = hypertextGetNLinksFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

/*
ConnectLinkSelected connects a callback to the 'link-selected' signal of the Hypertext.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Hypertext) ConnectLinkSelected(handler func(instance *Hypertext, arg1 int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "link-selected", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Hypertext) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var imageInterface *gi.Interface
var imageInterface_Once sync.Once

func imageInterface_Set() error {
	var err error
	imageInterface_Once.Do(func() {
		imageInterface, err = gi.InterfaceNew("Atk", "Image")
	})
	return err
}

type Image struct {
	native unsafe.Pointer
}

func ImageNewFromNative(native unsafe.Pointer) *Image {
	instance := &Image{native: native}

	return instance
}

/*
CastToImage down casts any arbitrary Object to Image.
Exercise care, as this is a potentially dangerous function
if the Object is not a Image.
*/
func CastToImage(object *gobject.Object) *Image {
	return ImageNewFromNative(object.Native())
}

// Equals compares this Image with another Image, and returns true if they represent the same Object.
func (recv *Image) Equals(other *Image) bool {
	return other.Native() == recv.Native()
}

func (recv *Image) Native() unsafe.Pointer {
	return recv.native
}

var imageGetImageDescriptionFunction *gi.Function
var imageGetImageDescriptionFunction_Once sync.Once

func imageGetImageDescriptionFunction_Set() error {
	var err error
	imageGetImageDescriptionFunction_Once.Do(func() {
		err = imageInterface_Set()
		if err != nil {
			return
		}
		imageGetImageDescriptionFunction, err = imageInterface.InvokerNew("get_image_description")
	})
	return err
}

// GetImageDescription is a representation of the C type atk_image_get_image_description.
func (recv *Image) GetImageDescription() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := imageGetImageDescriptionFunction_Set()
	if err == nil {
		ret = imageGetImageDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var imageGetImageLocaleFunction *gi.Function
var imageGetImageLocaleFunction_Once sync.Once

func imageGetImageLocaleFunction_Set() error {
	var err error
	imageGetImageLocaleFunction_Once.Do(func() {
		err = imageInterface_Set()
		if err != nil {
			return
		}
		imageGetImageLocaleFunction, err = imageInterface.InvokerNew("get_image_locale")
	})
	return err
}

// GetImageLocale is a representation of the C type atk_image_get_image_locale.
func (recv *Image) GetImageLocale() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := imageGetImageLocaleFunction_Set()
	if err == nil {
		ret = imageGetImageLocaleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var imageGetImagePositionFunction *gi.Function
var imageGetImagePositionFunction_Once sync.Once

func imageGetImagePositionFunction_Set() error {
	var err error
	imageGetImagePositionFunction_Once.Do(func() {
		err = imageInterface_Set()
		if err != nil {
			return
		}
		imageGetImagePositionFunction, err = imageInterface.InvokerNew("get_image_position")
	})
	return err
}

// GetImagePosition is a representation of the C type atk_image_get_image_position.
func (recv *Image) GetImagePosition(coordType CoordType) (int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(coordType))

	var outArgs [2]gi.Argument

	err := imageGetImagePositionFunction_Set()
	if err == nil {
		imageGetImagePositionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var imageGetImageSizeFunction *gi.Function
var imageGetImageSizeFunction_Once sync.Once

func imageGetImageSizeFunction_Set() error {
	var err error
	imageGetImageSizeFunction_Once.Do(func() {
		err = imageInterface_Set()
		if err != nil {
			return
		}
		imageGetImageSizeFunction, err = imageInterface.InvokerNew("get_image_size")
	})
	return err
}

// GetImageSize is a representation of the C type atk_image_get_image_size.
func (recv *Image) GetImageSize() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument

	err := imageGetImageSizeFunction_Set()
	if err == nil {
		imageGetImageSizeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var imageSetImageDescriptionFunction *gi.Function
var imageSetImageDescriptionFunction_Once sync.Once

func imageSetImageDescriptionFunction_Set() error {
	var err error
	imageSetImageDescriptionFunction_Once.Do(func() {
		err = imageInterface_Set()
		if err != nil {
			return
		}
		imageSetImageDescriptionFunction, err = imageInterface.InvokerNew("set_image_description")
	})
	return err
}

// SetImageDescription is a representation of the C type atk_image_set_image_description.
func (recv *Image) SetImageDescription(description string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(description)

	var ret gi.Argument

	err := imageSetImageDescriptionFunction_Set()
	if err == nil {
		ret = imageSetImageDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var implementorIfaceInterface *gi.Interface
var implementorIfaceInterface_Once sync.Once

func implementorIfaceInterface_Set() error {
	var err error
	implementorIfaceInterface_Once.Do(func() {
		implementorIfaceInterface, err = gi.InterfaceNew("Atk", "ImplementorIface")
	})
	return err
}

type ImplementorIface struct {
	native unsafe.Pointer
}

func ImplementorIfaceNewFromNative(native unsafe.Pointer) *ImplementorIface {
	instance := &ImplementorIface{native: native}

	return instance
}

/*
CastToImplementorIface down casts any arbitrary Object to ImplementorIface.
Exercise care, as this is a potentially dangerous function
if the Object is not a ImplementorIface.
*/
func CastToImplementorIface(object *gobject.Object) *ImplementorIface {
	return ImplementorIfaceNewFromNative(object.Native())
}

// Equals compares this ImplementorIface with another ImplementorIface, and returns true if they represent the same Object.
func (recv *ImplementorIface) Equals(other *ImplementorIface) bool {
	return other.Native() == recv.Native()
}

func (recv *ImplementorIface) Native() unsafe.Pointer {
	return recv.native
}

var selectionInterface *gi.Interface
var selectionInterface_Once sync.Once

func selectionInterface_Set() error {
	var err error
	selectionInterface_Once.Do(func() {
		selectionInterface, err = gi.InterfaceNew("Atk", "Selection")
	})
	return err
}

type Selection struct {
	native unsafe.Pointer
}

func SelectionNewFromNative(native unsafe.Pointer) *Selection {
	instance := &Selection{native: native}

	return instance
}

/*
CastToSelection down casts any arbitrary Object to Selection.
Exercise care, as this is a potentially dangerous function
if the Object is not a Selection.
*/
func CastToSelection(object *gobject.Object) *Selection {
	return SelectionNewFromNative(object.Native())
}

// Equals compares this Selection with another Selection, and returns true if they represent the same Object.
func (recv *Selection) Equals(other *Selection) bool {
	return other.Native() == recv.Native()
}

func (recv *Selection) Native() unsafe.Pointer {
	return recv.native
}

var selectionAddSelectionFunction *gi.Function
var selectionAddSelectionFunction_Once sync.Once

func selectionAddSelectionFunction_Set() error {
	var err error
	selectionAddSelectionFunction_Once.Do(func() {
		err = selectionInterface_Set()
		if err != nil {
			return
		}
		selectionAddSelectionFunction, err = selectionInterface.InvokerNew("add_selection")
	})
	return err
}

// AddSelection is a representation of the C type atk_selection_add_selection.
func (recv *Selection) AddSelection(i int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := selectionAddSelectionFunction_Set()
	if err == nil {
		ret = selectionAddSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var selectionClearSelectionFunction *gi.Function
var selectionClearSelectionFunction_Once sync.Once

func selectionClearSelectionFunction_Set() error {
	var err error
	selectionClearSelectionFunction_Once.Do(func() {
		err = selectionInterface_Set()
		if err != nil {
			return
		}
		selectionClearSelectionFunction, err = selectionInterface.InvokerNew("clear_selection")
	})
	return err
}

// ClearSelection is a representation of the C type atk_selection_clear_selection.
func (recv *Selection) ClearSelection() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := selectionClearSelectionFunction_Set()
	if err == nil {
		ret = selectionClearSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var selectionGetSelectionCountFunction *gi.Function
var selectionGetSelectionCountFunction_Once sync.Once

func selectionGetSelectionCountFunction_Set() error {
	var err error
	selectionGetSelectionCountFunction_Once.Do(func() {
		err = selectionInterface_Set()
		if err != nil {
			return
		}
		selectionGetSelectionCountFunction, err = selectionInterface.InvokerNew("get_selection_count")
	})
	return err
}

// GetSelectionCount is a representation of the C type atk_selection_get_selection_count.
func (recv *Selection) GetSelectionCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := selectionGetSelectionCountFunction_Set()
	if err == nil {
		ret = selectionGetSelectionCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var selectionIsChildSelectedFunction *gi.Function
var selectionIsChildSelectedFunction_Once sync.Once

func selectionIsChildSelectedFunction_Set() error {
	var err error
	selectionIsChildSelectedFunction_Once.Do(func() {
		err = selectionInterface_Set()
		if err != nil {
			return
		}
		selectionIsChildSelectedFunction, err = selectionInterface.InvokerNew("is_child_selected")
	})
	return err
}

// IsChildSelected is a representation of the C type atk_selection_is_child_selected.
func (recv *Selection) IsChildSelected(i int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := selectionIsChildSelectedFunction_Set()
	if err == nil {
		ret = selectionIsChildSelectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var selectionRefSelectionFunction *gi.Function
var selectionRefSelectionFunction_Once sync.Once

func selectionRefSelectionFunction_Set() error {
	var err error
	selectionRefSelectionFunction_Once.Do(func() {
		err = selectionInterface_Set()
		if err != nil {
			return
		}
		selectionRefSelectionFunction, err = selectionInterface.InvokerNew("ref_selection")
	})
	return err
}

// RefSelection is a representation of the C type atk_selection_ref_selection.
func (recv *Selection) RefSelection(i int32) *Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := selectionRefSelectionFunction_Set()
	if err == nil {
		ret = selectionRefSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var selectionRemoveSelectionFunction *gi.Function
var selectionRemoveSelectionFunction_Once sync.Once

func selectionRemoveSelectionFunction_Set() error {
	var err error
	selectionRemoveSelectionFunction_Once.Do(func() {
		err = selectionInterface_Set()
		if err != nil {
			return
		}
		selectionRemoveSelectionFunction, err = selectionInterface.InvokerNew("remove_selection")
	})
	return err
}

// RemoveSelection is a representation of the C type atk_selection_remove_selection.
func (recv *Selection) RemoveSelection(i int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := selectionRemoveSelectionFunction_Set()
	if err == nil {
		ret = selectionRemoveSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var selectionSelectAllSelectionFunction *gi.Function
var selectionSelectAllSelectionFunction_Once sync.Once

func selectionSelectAllSelectionFunction_Set() error {
	var err error
	selectionSelectAllSelectionFunction_Once.Do(func() {
		err = selectionInterface_Set()
		if err != nil {
			return
		}
		selectionSelectAllSelectionFunction, err = selectionInterface.InvokerNew("select_all_selection")
	})
	return err
}

// SelectAllSelection is a representation of the C type atk_selection_select_all_selection.
func (recv *Selection) SelectAllSelection() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := selectionSelectAllSelectionFunction_Set()
	if err == nil {
		ret = selectionSelectAllSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

/*
ConnectSelectionChanged connects a callback to the 'selection-changed' signal of the Selection.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Selection) ConnectSelectionChanged(handler func(instance *Selection)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "selection-changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Selection) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var streamableContentInterface *gi.Interface
var streamableContentInterface_Once sync.Once

func streamableContentInterface_Set() error {
	var err error
	streamableContentInterface_Once.Do(func() {
		streamableContentInterface, err = gi.InterfaceNew("Atk", "StreamableContent")
	})
	return err
}

type StreamableContent struct {
	native unsafe.Pointer
}

func StreamableContentNewFromNative(native unsafe.Pointer) *StreamableContent {
	instance := &StreamableContent{native: native}

	return instance
}

/*
CastToStreamableContent down casts any arbitrary Object to StreamableContent.
Exercise care, as this is a potentially dangerous function
if the Object is not a StreamableContent.
*/
func CastToStreamableContent(object *gobject.Object) *StreamableContent {
	return StreamableContentNewFromNative(object.Native())
}

// Equals compares this StreamableContent with another StreamableContent, and returns true if they represent the same Object.
func (recv *StreamableContent) Equals(other *StreamableContent) bool {
	return other.Native() == recv.Native()
}

func (recv *StreamableContent) Native() unsafe.Pointer {
	return recv.native
}

var streamableContentGetMimeTypeFunction *gi.Function
var streamableContentGetMimeTypeFunction_Once sync.Once

func streamableContentGetMimeTypeFunction_Set() error {
	var err error
	streamableContentGetMimeTypeFunction_Once.Do(func() {
		err = streamableContentInterface_Set()
		if err != nil {
			return
		}
		streamableContentGetMimeTypeFunction, err = streamableContentInterface.InvokerNew("get_mime_type")
	})
	return err
}

// GetMimeType is a representation of the C type atk_streamable_content_get_mime_type.
func (recv *StreamableContent) GetMimeType(i int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(i)

	var ret gi.Argument

	err := streamableContentGetMimeTypeFunction_Set()
	if err == nil {
		ret = streamableContentGetMimeTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var streamableContentGetNMimeTypesFunction *gi.Function
var streamableContentGetNMimeTypesFunction_Once sync.Once

func streamableContentGetNMimeTypesFunction_Set() error {
	var err error
	streamableContentGetNMimeTypesFunction_Once.Do(func() {
		err = streamableContentInterface_Set()
		if err != nil {
			return
		}
		streamableContentGetNMimeTypesFunction, err = streamableContentInterface.InvokerNew("get_n_mime_types")
	})
	return err
}

// GetNMimeTypes is a representation of the C type atk_streamable_content_get_n_mime_types.
func (recv *StreamableContent) GetNMimeTypes() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := streamableContentGetNMimeTypesFunction_Set()
	if err == nil {
		ret = streamableContentGetNMimeTypesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var streamableContentGetStreamFunction *gi.Function
var streamableContentGetStreamFunction_Once sync.Once

func streamableContentGetStreamFunction_Set() error {
	var err error
	streamableContentGetStreamFunction_Once.Do(func() {
		err = streamableContentInterface_Set()
		if err != nil {
			return
		}
		streamableContentGetStreamFunction, err = streamableContentInterface.InvokerNew("get_stream")
	})
	return err
}

// GetStream is a representation of the C type atk_streamable_content_get_stream.
func (recv *StreamableContent) GetStream(mimeType string) *glib.IOChannel {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(mimeType)

	var ret gi.Argument

	err := streamableContentGetStreamFunction_Set()
	if err == nil {
		ret = streamableContentGetStreamFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.IOChannelNewFromNative(ret.Pointer())

	return retGo
}

var streamableContentGetUriFunction *gi.Function
var streamableContentGetUriFunction_Once sync.Once

func streamableContentGetUriFunction_Set() error {
	var err error
	streamableContentGetUriFunction_Once.Do(func() {
		err = streamableContentInterface_Set()
		if err != nil {
			return
		}
		streamableContentGetUriFunction, err = streamableContentInterface.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type atk_streamable_content_get_uri.
func (recv *StreamableContent) GetUri(mimeType string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(mimeType)

	var ret gi.Argument

	err := streamableContentGetUriFunction_Set()
	if err == nil {
		ret = streamableContentGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var tableInterface *gi.Interface
var tableInterface_Once sync.Once

func tableInterface_Set() error {
	var err error
	tableInterface_Once.Do(func() {
		tableInterface, err = gi.InterfaceNew("Atk", "Table")
	})
	return err
}

type Table struct {
	native unsafe.Pointer
}

func TableNewFromNative(native unsafe.Pointer) *Table {
	instance := &Table{native: native}

	return instance
}

/*
CastToTable down casts any arbitrary Object to Table.
Exercise care, as this is a potentially dangerous function
if the Object is not a Table.
*/
func CastToTable(object *gobject.Object) *Table {
	return TableNewFromNative(object.Native())
}

// Equals compares this Table with another Table, and returns true if they represent the same Object.
func (recv *Table) Equals(other *Table) bool {
	return other.Native() == recv.Native()
}

func (recv *Table) Native() unsafe.Pointer {
	return recv.native
}

var tableAddColumnSelectionFunction *gi.Function
var tableAddColumnSelectionFunction_Once sync.Once

func tableAddColumnSelectionFunction_Set() error {
	var err error
	tableAddColumnSelectionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableAddColumnSelectionFunction, err = tableInterface.InvokerNew("add_column_selection")
	})
	return err
}

// AddColumnSelection is a representation of the C type atk_table_add_column_selection.
func (recv *Table) AddColumnSelection(column int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(column)

	var ret gi.Argument

	err := tableAddColumnSelectionFunction_Set()
	if err == nil {
		ret = tableAddColumnSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tableAddRowSelectionFunction *gi.Function
var tableAddRowSelectionFunction_Once sync.Once

func tableAddRowSelectionFunction_Set() error {
	var err error
	tableAddRowSelectionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableAddRowSelectionFunction, err = tableInterface.InvokerNew("add_row_selection")
	})
	return err
}

// AddRowSelection is a representation of the C type atk_table_add_row_selection.
func (recv *Table) AddRowSelection(row int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)

	var ret gi.Argument

	err := tableAddRowSelectionFunction_Set()
	if err == nil {
		ret = tableAddRowSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tableGetCaptionFunction *gi.Function
var tableGetCaptionFunction_Once sync.Once

func tableGetCaptionFunction_Set() error {
	var err error
	tableGetCaptionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetCaptionFunction, err = tableInterface.InvokerNew("get_caption")
	})
	return err
}

// GetCaption is a representation of the C type atk_table_get_caption.
func (recv *Table) GetCaption() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tableGetCaptionFunction_Set()
	if err == nil {
		ret = tableGetCaptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var tableGetColumnAtIndexFunction *gi.Function
var tableGetColumnAtIndexFunction_Once sync.Once

func tableGetColumnAtIndexFunction_Set() error {
	var err error
	tableGetColumnAtIndexFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetColumnAtIndexFunction, err = tableInterface.InvokerNew("get_column_at_index")
	})
	return err
}

// GetColumnAtIndex is a representation of the C type atk_table_get_column_at_index.
func (recv *Table) GetColumnAtIndex(index int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(index)

	var ret gi.Argument

	err := tableGetColumnAtIndexFunction_Set()
	if err == nil {
		ret = tableGetColumnAtIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableGetColumnDescriptionFunction *gi.Function
var tableGetColumnDescriptionFunction_Once sync.Once

func tableGetColumnDescriptionFunction_Set() error {
	var err error
	tableGetColumnDescriptionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetColumnDescriptionFunction, err = tableInterface.InvokerNew("get_column_description")
	})
	return err
}

// GetColumnDescription is a representation of the C type atk_table_get_column_description.
func (recv *Table) GetColumnDescription(column int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(column)

	var ret gi.Argument

	err := tableGetColumnDescriptionFunction_Set()
	if err == nil {
		ret = tableGetColumnDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var tableGetColumnExtentAtFunction *gi.Function
var tableGetColumnExtentAtFunction_Once sync.Once

func tableGetColumnExtentAtFunction_Set() error {
	var err error
	tableGetColumnExtentAtFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetColumnExtentAtFunction, err = tableInterface.InvokerNew("get_column_extent_at")
	})
	return err
}

// GetColumnExtentAt is a representation of the C type atk_table_get_column_extent_at.
func (recv *Table) GetColumnExtentAt(row int32, column int32) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)
	inArgs[2].SetInt32(column)

	var ret gi.Argument

	err := tableGetColumnExtentAtFunction_Set()
	if err == nil {
		ret = tableGetColumnExtentAtFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableGetColumnHeaderFunction *gi.Function
var tableGetColumnHeaderFunction_Once sync.Once

func tableGetColumnHeaderFunction_Set() error {
	var err error
	tableGetColumnHeaderFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetColumnHeaderFunction, err = tableInterface.InvokerNew("get_column_header")
	})
	return err
}

// GetColumnHeader is a representation of the C type atk_table_get_column_header.
func (recv *Table) GetColumnHeader(column int32) *Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(column)

	var ret gi.Argument

	err := tableGetColumnHeaderFunction_Set()
	if err == nil {
		ret = tableGetColumnHeaderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var tableGetIndexAtFunction *gi.Function
var tableGetIndexAtFunction_Once sync.Once

func tableGetIndexAtFunction_Set() error {
	var err error
	tableGetIndexAtFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetIndexAtFunction, err = tableInterface.InvokerNew("get_index_at")
	})
	return err
}

// GetIndexAt is a representation of the C type atk_table_get_index_at.
func (recv *Table) GetIndexAt(row int32, column int32) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)
	inArgs[2].SetInt32(column)

	var ret gi.Argument

	err := tableGetIndexAtFunction_Set()
	if err == nil {
		ret = tableGetIndexAtFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableGetNColumnsFunction *gi.Function
var tableGetNColumnsFunction_Once sync.Once

func tableGetNColumnsFunction_Set() error {
	var err error
	tableGetNColumnsFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetNColumnsFunction, err = tableInterface.InvokerNew("get_n_columns")
	})
	return err
}

// GetNColumns is a representation of the C type atk_table_get_n_columns.
func (recv *Table) GetNColumns() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tableGetNColumnsFunction_Set()
	if err == nil {
		ret = tableGetNColumnsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableGetNRowsFunction *gi.Function
var tableGetNRowsFunction_Once sync.Once

func tableGetNRowsFunction_Set() error {
	var err error
	tableGetNRowsFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetNRowsFunction, err = tableInterface.InvokerNew("get_n_rows")
	})
	return err
}

// GetNRows is a representation of the C type atk_table_get_n_rows.
func (recv *Table) GetNRows() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tableGetNRowsFunction_Set()
	if err == nil {
		ret = tableGetNRowsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableGetRowAtIndexFunction *gi.Function
var tableGetRowAtIndexFunction_Once sync.Once

func tableGetRowAtIndexFunction_Set() error {
	var err error
	tableGetRowAtIndexFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetRowAtIndexFunction, err = tableInterface.InvokerNew("get_row_at_index")
	})
	return err
}

// GetRowAtIndex is a representation of the C type atk_table_get_row_at_index.
func (recv *Table) GetRowAtIndex(index int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(index)

	var ret gi.Argument

	err := tableGetRowAtIndexFunction_Set()
	if err == nil {
		ret = tableGetRowAtIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableGetRowDescriptionFunction *gi.Function
var tableGetRowDescriptionFunction_Once sync.Once

func tableGetRowDescriptionFunction_Set() error {
	var err error
	tableGetRowDescriptionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetRowDescriptionFunction, err = tableInterface.InvokerNew("get_row_description")
	})
	return err
}

// GetRowDescription is a representation of the C type atk_table_get_row_description.
func (recv *Table) GetRowDescription(row int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)

	var ret gi.Argument

	err := tableGetRowDescriptionFunction_Set()
	if err == nil {
		ret = tableGetRowDescriptionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var tableGetRowExtentAtFunction *gi.Function
var tableGetRowExtentAtFunction_Once sync.Once

func tableGetRowExtentAtFunction_Set() error {
	var err error
	tableGetRowExtentAtFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetRowExtentAtFunction, err = tableInterface.InvokerNew("get_row_extent_at")
	})
	return err
}

// GetRowExtentAt is a representation of the C type atk_table_get_row_extent_at.
func (recv *Table) GetRowExtentAt(row int32, column int32) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)
	inArgs[2].SetInt32(column)

	var ret gi.Argument

	err := tableGetRowExtentAtFunction_Set()
	if err == nil {
		ret = tableGetRowExtentAtFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableGetRowHeaderFunction *gi.Function
var tableGetRowHeaderFunction_Once sync.Once

func tableGetRowHeaderFunction_Set() error {
	var err error
	tableGetRowHeaderFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetRowHeaderFunction, err = tableInterface.InvokerNew("get_row_header")
	})
	return err
}

// GetRowHeader is a representation of the C type atk_table_get_row_header.
func (recv *Table) GetRowHeader(row int32) *Object {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)

	var ret gi.Argument

	err := tableGetRowHeaderFunction_Set()
	if err == nil {
		ret = tableGetRowHeaderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var tableGetSelectedColumnsFunction *gi.Function
var tableGetSelectedColumnsFunction_Once sync.Once

func tableGetSelectedColumnsFunction_Set() error {
	var err error
	tableGetSelectedColumnsFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetSelectedColumnsFunction, err = tableInterface.InvokerNew("get_selected_columns")
	})
	return err
}

// GetSelectedColumns is a representation of the C type atk_table_get_selected_columns.
func (recv *Table) GetSelectedColumns(selected int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(selected)

	var ret gi.Argument

	err := tableGetSelectedColumnsFunction_Set()
	if err == nil {
		ret = tableGetSelectedColumnsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableGetSelectedRowsFunction *gi.Function
var tableGetSelectedRowsFunction_Once sync.Once

func tableGetSelectedRowsFunction_Set() error {
	var err error
	tableGetSelectedRowsFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetSelectedRowsFunction, err = tableInterface.InvokerNew("get_selected_rows")
	})
	return err
}

// GetSelectedRows is a representation of the C type atk_table_get_selected_rows.
func (recv *Table) GetSelectedRows(selected int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(selected)

	var ret gi.Argument

	err := tableGetSelectedRowsFunction_Set()
	if err == nil {
		ret = tableGetSelectedRowsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableGetSummaryFunction *gi.Function
var tableGetSummaryFunction_Once sync.Once

func tableGetSummaryFunction_Set() error {
	var err error
	tableGetSummaryFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableGetSummaryFunction, err = tableInterface.InvokerNew("get_summary")
	})
	return err
}

// GetSummary is a representation of the C type atk_table_get_summary.
func (recv *Table) GetSummary() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tableGetSummaryFunction_Set()
	if err == nil {
		ret = tableGetSummaryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var tableIsColumnSelectedFunction *gi.Function
var tableIsColumnSelectedFunction_Once sync.Once

func tableIsColumnSelectedFunction_Set() error {
	var err error
	tableIsColumnSelectedFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableIsColumnSelectedFunction, err = tableInterface.InvokerNew("is_column_selected")
	})
	return err
}

// IsColumnSelected is a representation of the C type atk_table_is_column_selected.
func (recv *Table) IsColumnSelected(column int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(column)

	var ret gi.Argument

	err := tableIsColumnSelectedFunction_Set()
	if err == nil {
		ret = tableIsColumnSelectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tableIsRowSelectedFunction *gi.Function
var tableIsRowSelectedFunction_Once sync.Once

func tableIsRowSelectedFunction_Set() error {
	var err error
	tableIsRowSelectedFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableIsRowSelectedFunction, err = tableInterface.InvokerNew("is_row_selected")
	})
	return err
}

// IsRowSelected is a representation of the C type atk_table_is_row_selected.
func (recv *Table) IsRowSelected(row int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)

	var ret gi.Argument

	err := tableIsRowSelectedFunction_Set()
	if err == nil {
		ret = tableIsRowSelectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tableIsSelectedFunction *gi.Function
var tableIsSelectedFunction_Once sync.Once

func tableIsSelectedFunction_Set() error {
	var err error
	tableIsSelectedFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableIsSelectedFunction, err = tableInterface.InvokerNew("is_selected")
	})
	return err
}

// IsSelected is a representation of the C type atk_table_is_selected.
func (recv *Table) IsSelected(row int32, column int32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)
	inArgs[2].SetInt32(column)

	var ret gi.Argument

	err := tableIsSelectedFunction_Set()
	if err == nil {
		ret = tableIsSelectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tableRefAtFunction *gi.Function
var tableRefAtFunction_Once sync.Once

func tableRefAtFunction_Set() error {
	var err error
	tableRefAtFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableRefAtFunction, err = tableInterface.InvokerNew("ref_at")
	})
	return err
}

// RefAt is a representation of the C type atk_table_ref_at.
func (recv *Table) RefAt(row int32, column int32) *Object {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)
	inArgs[2].SetInt32(column)

	var ret gi.Argument

	err := tableRefAtFunction_Set()
	if err == nil {
		ret = tableRefAtFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var tableRemoveColumnSelectionFunction *gi.Function
var tableRemoveColumnSelectionFunction_Once sync.Once

func tableRemoveColumnSelectionFunction_Set() error {
	var err error
	tableRemoveColumnSelectionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableRemoveColumnSelectionFunction, err = tableInterface.InvokerNew("remove_column_selection")
	})
	return err
}

// RemoveColumnSelection is a representation of the C type atk_table_remove_column_selection.
func (recv *Table) RemoveColumnSelection(column int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(column)

	var ret gi.Argument

	err := tableRemoveColumnSelectionFunction_Set()
	if err == nil {
		ret = tableRemoveColumnSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tableRemoveRowSelectionFunction *gi.Function
var tableRemoveRowSelectionFunction_Once sync.Once

func tableRemoveRowSelectionFunction_Set() error {
	var err error
	tableRemoveRowSelectionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableRemoveRowSelectionFunction, err = tableInterface.InvokerNew("remove_row_selection")
	})
	return err
}

// RemoveRowSelection is a representation of the C type atk_table_remove_row_selection.
func (recv *Table) RemoveRowSelection(row int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)

	var ret gi.Argument

	err := tableRemoveRowSelectionFunction_Set()
	if err == nil {
		ret = tableRemoveRowSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tableSetCaptionFunction *gi.Function
var tableSetCaptionFunction_Once sync.Once

func tableSetCaptionFunction_Set() error {
	var err error
	tableSetCaptionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableSetCaptionFunction, err = tableInterface.InvokerNew("set_caption")
	})
	return err
}

// SetCaption is a representation of the C type atk_table_set_caption.
func (recv *Table) SetCaption(caption *Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(caption.Native())

	err := tableSetCaptionFunction_Set()
	if err == nil {
		tableSetCaptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tableSetColumnDescriptionFunction *gi.Function
var tableSetColumnDescriptionFunction_Once sync.Once

func tableSetColumnDescriptionFunction_Set() error {
	var err error
	tableSetColumnDescriptionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableSetColumnDescriptionFunction, err = tableInterface.InvokerNew("set_column_description")
	})
	return err
}

// SetColumnDescription is a representation of the C type atk_table_set_column_description.
func (recv *Table) SetColumnDescription(column int32, description string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(column)
	inArgs[2].SetString(description)

	err := tableSetColumnDescriptionFunction_Set()
	if err == nil {
		tableSetColumnDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tableSetColumnHeaderFunction *gi.Function
var tableSetColumnHeaderFunction_Once sync.Once

func tableSetColumnHeaderFunction_Set() error {
	var err error
	tableSetColumnHeaderFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableSetColumnHeaderFunction, err = tableInterface.InvokerNew("set_column_header")
	})
	return err
}

// SetColumnHeader is a representation of the C type atk_table_set_column_header.
func (recv *Table) SetColumnHeader(column int32, header *Object) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(column)
	inArgs[2].SetPointer(header.Native())

	err := tableSetColumnHeaderFunction_Set()
	if err == nil {
		tableSetColumnHeaderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tableSetRowDescriptionFunction *gi.Function
var tableSetRowDescriptionFunction_Once sync.Once

func tableSetRowDescriptionFunction_Set() error {
	var err error
	tableSetRowDescriptionFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableSetRowDescriptionFunction, err = tableInterface.InvokerNew("set_row_description")
	})
	return err
}

// SetRowDescription is a representation of the C type atk_table_set_row_description.
func (recv *Table) SetRowDescription(row int32, description string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)
	inArgs[2].SetString(description)

	err := tableSetRowDescriptionFunction_Set()
	if err == nil {
		tableSetRowDescriptionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tableSetRowHeaderFunction *gi.Function
var tableSetRowHeaderFunction_Once sync.Once

func tableSetRowHeaderFunction_Set() error {
	var err error
	tableSetRowHeaderFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableSetRowHeaderFunction, err = tableInterface.InvokerNew("set_row_header")
	})
	return err
}

// SetRowHeader is a representation of the C type atk_table_set_row_header.
func (recv *Table) SetRowHeader(row int32, header *Object) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(row)
	inArgs[2].SetPointer(header.Native())

	err := tableSetRowHeaderFunction_Set()
	if err == nil {
		tableSetRowHeaderFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tableSetSummaryFunction *gi.Function
var tableSetSummaryFunction_Once sync.Once

func tableSetSummaryFunction_Set() error {
	var err error
	tableSetSummaryFunction_Once.Do(func() {
		err = tableInterface_Set()
		if err != nil {
			return
		}
		tableSetSummaryFunction, err = tableInterface.InvokerNew("set_summary")
	})
	return err
}

// SetSummary is a representation of the C type atk_table_set_summary.
func (recv *Table) SetSummary(accessible *Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(accessible.Native())

	err := tableSetSummaryFunction_Set()
	if err == nil {
		tableSetSummaryFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectColumnDeleted connects a callback to the 'column-deleted' signal of the Table.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Table) ConnectColumnDeleted(handler func(instance *Table, arg1 int32, arg2 int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "column-deleted", marshal)
}

/*
ConnectColumnInserted connects a callback to the 'column-inserted' signal of the Table.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Table) ConnectColumnInserted(handler func(instance *Table, arg1 int32, arg2 int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "column-inserted", marshal)
}

/*
ConnectColumnReordered connects a callback to the 'column-reordered' signal of the Table.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Table) ConnectColumnReordered(handler func(instance *Table)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "column-reordered", marshal)
}

/*
ConnectModelChanged connects a callback to the 'model-changed' signal of the Table.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Table) ConnectModelChanged(handler func(instance *Table)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "model-changed", marshal)
}

/*
ConnectRowDeleted connects a callback to the 'row-deleted' signal of the Table.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Table) ConnectRowDeleted(handler func(instance *Table, arg1 int32, arg2 int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "row-deleted", marshal)
}

/*
ConnectRowInserted connects a callback to the 'row-inserted' signal of the Table.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Table) ConnectRowInserted(handler func(instance *Table, arg1 int32, arg2 int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "row-inserted", marshal)
}

/*
ConnectRowReordered connects a callback to the 'row-reordered' signal of the Table.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Table) ConnectRowReordered(handler func(instance *Table)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "row-reordered", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Table) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var tableCellInterface *gi.Interface
var tableCellInterface_Once sync.Once

func tableCellInterface_Set() error {
	var err error
	tableCellInterface_Once.Do(func() {
		tableCellInterface, err = gi.InterfaceNew("Atk", "TableCell")
	})
	return err
}

type TableCell struct {
	native unsafe.Pointer
}

func TableCellNewFromNative(native unsafe.Pointer) *TableCell {
	instance := &TableCell{native: native}

	return instance
}

/*
CastToTableCell down casts any arbitrary Object to TableCell.
Exercise care, as this is a potentially dangerous function
if the Object is not a TableCell.
*/
func CastToTableCell(object *gobject.Object) *TableCell {
	return TableCellNewFromNative(object.Native())
}

// Equals compares this TableCell with another TableCell, and returns true if they represent the same Object.
func (recv *TableCell) Equals(other *TableCell) bool {
	return other.Native() == recv.Native()
}

func (recv *TableCell) Native() unsafe.Pointer {
	return recv.native
}

var tableCellGetColumnHeaderCellsFunction *gi.Function
var tableCellGetColumnHeaderCellsFunction_Once sync.Once

func tableCellGetColumnHeaderCellsFunction_Set() error {
	var err error
	tableCellGetColumnHeaderCellsFunction_Once.Do(func() {
		err = tableCellInterface_Set()
		if err != nil {
			return
		}
		tableCellGetColumnHeaderCellsFunction, err = tableCellInterface.InvokerNew("get_column_header_cells")
	})
	return err
}

// GetColumnHeaderCells is a representation of the C type atk_table_cell_get_column_header_cells.
func (recv *TableCell) GetColumnHeaderCells() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := tableCellGetColumnHeaderCellsFunction_Set()
	if err == nil {
		tableCellGetColumnHeaderCellsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tableCellGetColumnSpanFunction *gi.Function
var tableCellGetColumnSpanFunction_Once sync.Once

func tableCellGetColumnSpanFunction_Set() error {
	var err error
	tableCellGetColumnSpanFunction_Once.Do(func() {
		err = tableCellInterface_Set()
		if err != nil {
			return
		}
		tableCellGetColumnSpanFunction, err = tableCellInterface.InvokerNew("get_column_span")
	})
	return err
}

// GetColumnSpan is a representation of the C type atk_table_cell_get_column_span.
func (recv *TableCell) GetColumnSpan() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tableCellGetColumnSpanFunction_Set()
	if err == nil {
		ret = tableCellGetColumnSpanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableCellGetPositionFunction *gi.Function
var tableCellGetPositionFunction_Once sync.Once

func tableCellGetPositionFunction_Set() error {
	var err error
	tableCellGetPositionFunction_Once.Do(func() {
		err = tableCellInterface_Set()
		if err != nil {
			return
		}
		tableCellGetPositionFunction, err = tableCellInterface.InvokerNew("get_position")
	})
	return err
}

// GetPosition is a representation of the C type atk_table_cell_get_position.
func (recv *TableCell) GetPosition() (bool, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := tableCellGetPositionFunction_Set()
	if err == nil {
		ret = tableCellGetPositionFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var tableCellGetRowColumnSpanFunction *gi.Function
var tableCellGetRowColumnSpanFunction_Once sync.Once

func tableCellGetRowColumnSpanFunction_Set() error {
	var err error
	tableCellGetRowColumnSpanFunction_Once.Do(func() {
		err = tableCellInterface_Set()
		if err != nil {
			return
		}
		tableCellGetRowColumnSpanFunction, err = tableCellInterface.InvokerNew("get_row_column_span")
	})
	return err
}

// GetRowColumnSpan is a representation of the C type atk_table_cell_get_row_column_span.
func (recv *TableCell) GetRowColumnSpan() (bool, int32, int32, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [4]gi.Argument
	var ret gi.Argument

	err := tableCellGetRowColumnSpanFunction_Set()
	if err == nil {
		ret = tableCellGetRowColumnSpanFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()
	out3 := outArgs[3].Int32()

	return retGo, out0, out1, out2, out3
}

var tableCellGetRowHeaderCellsFunction *gi.Function
var tableCellGetRowHeaderCellsFunction_Once sync.Once

func tableCellGetRowHeaderCellsFunction_Set() error {
	var err error
	tableCellGetRowHeaderCellsFunction_Once.Do(func() {
		err = tableCellInterface_Set()
		if err != nil {
			return
		}
		tableCellGetRowHeaderCellsFunction, err = tableCellInterface.InvokerNew("get_row_header_cells")
	})
	return err
}

// GetRowHeaderCells is a representation of the C type atk_table_cell_get_row_header_cells.
func (recv *TableCell) GetRowHeaderCells() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := tableCellGetRowHeaderCellsFunction_Set()
	if err == nil {
		tableCellGetRowHeaderCellsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tableCellGetRowSpanFunction *gi.Function
var tableCellGetRowSpanFunction_Once sync.Once

func tableCellGetRowSpanFunction_Set() error {
	var err error
	tableCellGetRowSpanFunction_Once.Do(func() {
		err = tableCellInterface_Set()
		if err != nil {
			return
		}
		tableCellGetRowSpanFunction, err = tableCellInterface.InvokerNew("get_row_span")
	})
	return err
}

// GetRowSpan is a representation of the C type atk_table_cell_get_row_span.
func (recv *TableCell) GetRowSpan() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tableCellGetRowSpanFunction_Set()
	if err == nil {
		ret = tableCellGetRowSpanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tableCellGetTableFunction *gi.Function
var tableCellGetTableFunction_Once sync.Once

func tableCellGetTableFunction_Set() error {
	var err error
	tableCellGetTableFunction_Once.Do(func() {
		err = tableCellInterface_Set()
		if err != nil {
			return
		}
		tableCellGetTableFunction, err = tableCellInterface.InvokerNew("get_table")
	})
	return err
}

// GetTable is a representation of the C type atk_table_cell_get_table.
func (recv *TableCell) GetTable() *Object {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := tableCellGetTableFunction_Set()
	if err == nil {
		ret = tableCellGetTableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ObjectNewFromNative(ret.Pointer())

	return retGo
}

var textInterface *gi.Interface
var textInterface_Once sync.Once

func textInterface_Set() error {
	var err error
	textInterface_Once.Do(func() {
		textInterface, err = gi.InterfaceNew("Atk", "Text")
	})
	return err
}

type Text struct {
	native unsafe.Pointer
}

func TextNewFromNative(native unsafe.Pointer) *Text {
	instance := &Text{native: native}

	return instance
}

/*
CastToText down casts any arbitrary Object to Text.
Exercise care, as this is a potentially dangerous function
if the Object is not a Text.
*/
func CastToText(object *gobject.Object) *Text {
	return TextNewFromNative(object.Native())
}

// Equals compares this Text with another Text, and returns true if they represent the same Object.
func (recv *Text) Equals(other *Text) bool {
	return other.Native() == recv.Native()
}

func (recv *Text) Native() unsafe.Pointer {
	return recv.native
}

var textAddSelectionFunction *gi.Function
var textAddSelectionFunction_Once sync.Once

func textAddSelectionFunction_Set() error {
	var err error
	textAddSelectionFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textAddSelectionFunction, err = textInterface.InvokerNew("add_selection")
	})
	return err
}

// AddSelection is a representation of the C type atk_text_add_selection.
func (recv *Text) AddSelection(startOffset int32, endOffset int32) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startOffset)
	inArgs[2].SetInt32(endOffset)

	var ret gi.Argument

	err := textAddSelectionFunction_Set()
	if err == nil {
		ret = textAddSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textGetBoundedRangesFunction *gi.Function
var textGetBoundedRangesFunction_Once sync.Once

func textGetBoundedRangesFunction_Set() error {
	var err error
	textGetBoundedRangesFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetBoundedRangesFunction, err = textInterface.InvokerNew("get_bounded_ranges")
	})
	return err
}

// GetBoundedRanges is a representation of the C type atk_text_get_bounded_ranges.
func (recv *Text) GetBoundedRanges(rect *TextRectangle, coordType CoordType, xClipType TextClipType, yClipType TextClipType) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(rect.Native())
	inArgs[2].SetInt32(int32(coordType))
	inArgs[3].SetInt32(int32(xClipType))
	inArgs[4].SetInt32(int32(yClipType))

	err := textGetBoundedRangesFunction_Set()
	if err == nil {
		textGetBoundedRangesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var textGetCaretOffsetFunction *gi.Function
var textGetCaretOffsetFunction_Once sync.Once

func textGetCaretOffsetFunction_Set() error {
	var err error
	textGetCaretOffsetFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetCaretOffsetFunction, err = textInterface.InvokerNew("get_caret_offset")
	})
	return err
}

// GetCaretOffset is a representation of the C type atk_text_get_caret_offset.
func (recv *Text) GetCaretOffset() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := textGetCaretOffsetFunction_Set()
	if err == nil {
		ret = textGetCaretOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'atk_text_get_character_at_offset' : return type 'gunichar' not supported

var textGetCharacterCountFunction *gi.Function
var textGetCharacterCountFunction_Once sync.Once

func textGetCharacterCountFunction_Set() error {
	var err error
	textGetCharacterCountFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetCharacterCountFunction, err = textInterface.InvokerNew("get_character_count")
	})
	return err
}

// GetCharacterCount is a representation of the C type atk_text_get_character_count.
func (recv *Text) GetCharacterCount() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := textGetCharacterCountFunction_Set()
	if err == nil {
		ret = textGetCharacterCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var textGetCharacterExtentsFunction *gi.Function
var textGetCharacterExtentsFunction_Once sync.Once

func textGetCharacterExtentsFunction_Set() error {
	var err error
	textGetCharacterExtentsFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetCharacterExtentsFunction, err = textInterface.InvokerNew("get_character_extents")
	})
	return err
}

// GetCharacterExtents is a representation of the C type atk_text_get_character_extents.
func (recv *Text) GetCharacterExtents(offset int32, coords CoordType) (int32, int32, int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(offset)
	inArgs[2].SetInt32(int32(coords))

	var outArgs [4]gi.Argument

	err := textGetCharacterExtentsFunction_Set()
	if err == nil {
		textGetCharacterExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()
	out3 := outArgs[3].Int32()

	return out0, out1, out2, out3
}

// UNSUPPORTED : C value 'atk_text_get_default_attributes' : return type 'AttributeSet' not supported

var textGetNSelectionsFunction *gi.Function
var textGetNSelectionsFunction_Once sync.Once

func textGetNSelectionsFunction_Set() error {
	var err error
	textGetNSelectionsFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetNSelectionsFunction, err = textInterface.InvokerNew("get_n_selections")
	})
	return err
}

// GetNSelections is a representation of the C type atk_text_get_n_selections.
func (recv *Text) GetNSelections() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := textGetNSelectionsFunction_Set()
	if err == nil {
		ret = textGetNSelectionsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var textGetOffsetAtPointFunction *gi.Function
var textGetOffsetAtPointFunction_Once sync.Once

func textGetOffsetAtPointFunction_Set() error {
	var err error
	textGetOffsetAtPointFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetOffsetAtPointFunction, err = textInterface.InvokerNew("get_offset_at_point")
	})
	return err
}

// GetOffsetAtPoint is a representation of the C type atk_text_get_offset_at_point.
func (recv *Text) GetOffsetAtPoint(x int32, y int32, coords CoordType) int32 {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)
	inArgs[3].SetInt32(int32(coords))

	var ret gi.Argument

	err := textGetOffsetAtPointFunction_Set()
	if err == nil {
		ret = textGetOffsetAtPointFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var textGetRangeExtentsFunction *gi.Function
var textGetRangeExtentsFunction_Once sync.Once

func textGetRangeExtentsFunction_Set() error {
	var err error
	textGetRangeExtentsFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetRangeExtentsFunction, err = textInterface.InvokerNew("get_range_extents")
	})
	return err
}

// GetRangeExtents is a representation of the C type atk_text_get_range_extents.
func (recv *Text) GetRangeExtents(startOffset int32, endOffset int32, coordType CoordType) *TextRectangle {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startOffset)
	inArgs[2].SetInt32(endOffset)
	inArgs[3].SetInt32(int32(coordType))

	var outArgs [1]gi.Argument

	err := textGetRangeExtentsFunction_Set()
	if err == nil {
		textGetRangeExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := TextRectangleNewFromNative(outArgs[0].Pointer())

	return out0
}

// UNSUPPORTED : C value 'atk_text_get_run_attributes' : return type 'AttributeSet' not supported

var textGetSelectionFunction *gi.Function
var textGetSelectionFunction_Once sync.Once

func textGetSelectionFunction_Set() error {
	var err error
	textGetSelectionFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetSelectionFunction, err = textInterface.InvokerNew("get_selection")
	})
	return err
}

// GetSelection is a representation of the C type atk_text_get_selection.
func (recv *Text) GetSelection(selectionNum int32) (string, int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(selectionNum)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := textGetSelectionFunction_Set()
	if err == nil {
		ret = textGetSelectionFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var textGetStringAtOffsetFunction *gi.Function
var textGetStringAtOffsetFunction_Once sync.Once

func textGetStringAtOffsetFunction_Set() error {
	var err error
	textGetStringAtOffsetFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetStringAtOffsetFunction, err = textInterface.InvokerNew("get_string_at_offset")
	})
	return err
}

// GetStringAtOffset is a representation of the C type atk_text_get_string_at_offset.
func (recv *Text) GetStringAtOffset(offset int32, granularity TextGranularity) (string, int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(offset)
	inArgs[2].SetInt32(int32(granularity))

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := textGetStringAtOffsetFunction_Set()
	if err == nil {
		ret = textGetStringAtOffsetFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var textGetTextFunction *gi.Function
var textGetTextFunction_Once sync.Once

func textGetTextFunction_Set() error {
	var err error
	textGetTextFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetTextFunction, err = textInterface.InvokerNew("get_text")
	})
	return err
}

// GetText is a representation of the C type atk_text_get_text.
func (recv *Text) GetText(startOffset int32, endOffset int32) string {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startOffset)
	inArgs[2].SetInt32(endOffset)

	var ret gi.Argument

	err := textGetTextFunction_Set()
	if err == nil {
		ret = textGetTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var textGetTextAfterOffsetFunction *gi.Function
var textGetTextAfterOffsetFunction_Once sync.Once

func textGetTextAfterOffsetFunction_Set() error {
	var err error
	textGetTextAfterOffsetFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetTextAfterOffsetFunction, err = textInterface.InvokerNew("get_text_after_offset")
	})
	return err
}

// GetTextAfterOffset is a representation of the C type atk_text_get_text_after_offset.
func (recv *Text) GetTextAfterOffset(offset int32, boundaryType TextBoundary) (string, int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(offset)
	inArgs[2].SetInt32(int32(boundaryType))

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := textGetTextAfterOffsetFunction_Set()
	if err == nil {
		ret = textGetTextAfterOffsetFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var textGetTextAtOffsetFunction *gi.Function
var textGetTextAtOffsetFunction_Once sync.Once

func textGetTextAtOffsetFunction_Set() error {
	var err error
	textGetTextAtOffsetFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetTextAtOffsetFunction, err = textInterface.InvokerNew("get_text_at_offset")
	})
	return err
}

// GetTextAtOffset is a representation of the C type atk_text_get_text_at_offset.
func (recv *Text) GetTextAtOffset(offset int32, boundaryType TextBoundary) (string, int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(offset)
	inArgs[2].SetInt32(int32(boundaryType))

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := textGetTextAtOffsetFunction_Set()
	if err == nil {
		ret = textGetTextAtOffsetFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var textGetTextBeforeOffsetFunction *gi.Function
var textGetTextBeforeOffsetFunction_Once sync.Once

func textGetTextBeforeOffsetFunction_Set() error {
	var err error
	textGetTextBeforeOffsetFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textGetTextBeforeOffsetFunction, err = textInterface.InvokerNew("get_text_before_offset")
	})
	return err
}

// GetTextBeforeOffset is a representation of the C type atk_text_get_text_before_offset.
func (recv *Text) GetTextBeforeOffset(offset int32, boundaryType TextBoundary) (string, int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(offset)
	inArgs[2].SetInt32(int32(boundaryType))

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := textGetTextBeforeOffsetFunction_Set()
	if err == nil {
		ret = textGetTextBeforeOffsetFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.String(true)
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var textRemoveSelectionFunction *gi.Function
var textRemoveSelectionFunction_Once sync.Once

func textRemoveSelectionFunction_Set() error {
	var err error
	textRemoveSelectionFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textRemoveSelectionFunction, err = textInterface.InvokerNew("remove_selection")
	})
	return err
}

// RemoveSelection is a representation of the C type atk_text_remove_selection.
func (recv *Text) RemoveSelection(selectionNum int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(selectionNum)

	var ret gi.Argument

	err := textRemoveSelectionFunction_Set()
	if err == nil {
		ret = textRemoveSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textScrollSubstringToFunction *gi.Function
var textScrollSubstringToFunction_Once sync.Once

func textScrollSubstringToFunction_Set() error {
	var err error
	textScrollSubstringToFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textScrollSubstringToFunction, err = textInterface.InvokerNew("scroll_substring_to")
	})
	return err
}

// ScrollSubstringTo is a representation of the C type atk_text_scroll_substring_to.
func (recv *Text) ScrollSubstringTo(startOffset int32, endOffset int32, type_ ScrollType) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startOffset)
	inArgs[2].SetInt32(endOffset)
	inArgs[3].SetInt32(int32(type_))

	var ret gi.Argument

	err := textScrollSubstringToFunction_Set()
	if err == nil {
		ret = textScrollSubstringToFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textScrollSubstringToPointFunction *gi.Function
var textScrollSubstringToPointFunction_Once sync.Once

func textScrollSubstringToPointFunction_Set() error {
	var err error
	textScrollSubstringToPointFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textScrollSubstringToPointFunction, err = textInterface.InvokerNew("scroll_substring_to_point")
	})
	return err
}

// ScrollSubstringToPoint is a representation of the C type atk_text_scroll_substring_to_point.
func (recv *Text) ScrollSubstringToPoint(startOffset int32, endOffset int32, coords CoordType, x int32, y int32) bool {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startOffset)
	inArgs[2].SetInt32(endOffset)
	inArgs[3].SetInt32(int32(coords))
	inArgs[4].SetInt32(x)
	inArgs[5].SetInt32(y)

	var ret gi.Argument

	err := textScrollSubstringToPointFunction_Set()
	if err == nil {
		ret = textScrollSubstringToPointFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textSetCaretOffsetFunction *gi.Function
var textSetCaretOffsetFunction_Once sync.Once

func textSetCaretOffsetFunction_Set() error {
	var err error
	textSetCaretOffsetFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textSetCaretOffsetFunction, err = textInterface.InvokerNew("set_caret_offset")
	})
	return err
}

// SetCaretOffset is a representation of the C type atk_text_set_caret_offset.
func (recv *Text) SetCaretOffset(offset int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(offset)

	var ret gi.Argument

	err := textSetCaretOffsetFunction_Set()
	if err == nil {
		ret = textSetCaretOffsetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var textSetSelectionFunction *gi.Function
var textSetSelectionFunction_Once sync.Once

func textSetSelectionFunction_Set() error {
	var err error
	textSetSelectionFunction_Once.Do(func() {
		err = textInterface_Set()
		if err != nil {
			return
		}
		textSetSelectionFunction, err = textInterface.InvokerNew("set_selection")
	})
	return err
}

// SetSelection is a representation of the C type atk_text_set_selection.
func (recv *Text) SetSelection(selectionNum int32, startOffset int32, endOffset int32) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(selectionNum)
	inArgs[2].SetInt32(startOffset)
	inArgs[3].SetInt32(endOffset)

	var ret gi.Argument

	err := textSetSelectionFunction_Set()
	if err == nil {
		ret = textSetSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

/*
ConnectTextAttributesChanged connects a callback to the 'text-attributes-changed' signal of the Text.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Text) ConnectTextAttributesChanged(handler func(instance *Text)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "text-attributes-changed", marshal)
}

/*
ConnectTextCaretMoved connects a callback to the 'text-caret-moved' signal of the Text.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Text) ConnectTextCaretMoved(handler func(instance *Text, arg1 int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "text-caret-moved", marshal)
}

/*
ConnectTextChanged connects a callback to the 'text-changed' signal of the Text.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Text) ConnectTextChanged(handler func(instance *Text, arg1 int32, arg2 int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "text-changed", marshal)
}

/*
ConnectTextInsert connects a callback to the 'text-insert' signal of the Text.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Text) ConnectTextInsert(handler func(instance *Text, arg1 int32, arg2 int32, arg3 string)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "text-insert", marshal)
}

/*
ConnectTextRemove connects a callback to the 'text-remove' signal of the Text.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Text) ConnectTextRemove(handler func(instance *Text, arg1 int32, arg2 int32, arg3 string)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "text-remove", marshal)
}

/*
ConnectTextSelectionChanged connects a callback to the 'text-selection-changed' signal of the Text.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Text) ConnectTextSelectionChanged(handler func(instance *Text)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "text-selection-changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Text) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var valueInterface *gi.Interface
var valueInterface_Once sync.Once

func valueInterface_Set() error {
	var err error
	valueInterface_Once.Do(func() {
		valueInterface, err = gi.InterfaceNew("Atk", "Value")
	})
	return err
}

type Value struct {
	native unsafe.Pointer
}

func ValueNewFromNative(native unsafe.Pointer) *Value {
	instance := &Value{native: native}

	return instance
}

/*
CastToValue down casts any arbitrary Object to Value.
Exercise care, as this is a potentially dangerous function
if the Object is not a Value.
*/
func CastToValue(object *gobject.Object) *Value {
	return ValueNewFromNative(object.Native())
}

// Equals compares this Value with another Value, and returns true if they represent the same Object.
func (recv *Value) Equals(other *Value) bool {
	return other.Native() == recv.Native()
}

func (recv *Value) Native() unsafe.Pointer {
	return recv.native
}

var valueGetCurrentValueFunction *gi.Function
var valueGetCurrentValueFunction_Once sync.Once

func valueGetCurrentValueFunction_Set() error {
	var err error
	valueGetCurrentValueFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueGetCurrentValueFunction, err = valueInterface.InvokerNew("get_current_value")
	})
	return err
}

// GetCurrentValue is a representation of the C type atk_value_get_current_value.
func (recv *Value) GetCurrentValue() *gobject.Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := valueGetCurrentValueFunction_Set()
	if err == nil {
		valueGetCurrentValueFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := gobject.ValueNewFromNative(outArgs[0].Pointer())

	return out0
}

var valueGetIncrementFunction *gi.Function
var valueGetIncrementFunction_Once sync.Once

func valueGetIncrementFunction_Set() error {
	var err error
	valueGetIncrementFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueGetIncrementFunction, err = valueInterface.InvokerNew("get_increment")
	})
	return err
}

// GetIncrement is a representation of the C type atk_value_get_increment.
func (recv *Value) GetIncrement() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueGetIncrementFunction_Set()
	if err == nil {
		ret = valueGetIncrementFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var valueGetMaximumValueFunction *gi.Function
var valueGetMaximumValueFunction_Once sync.Once

func valueGetMaximumValueFunction_Set() error {
	var err error
	valueGetMaximumValueFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueGetMaximumValueFunction, err = valueInterface.InvokerNew("get_maximum_value")
	})
	return err
}

// GetMaximumValue is a representation of the C type atk_value_get_maximum_value.
func (recv *Value) GetMaximumValue() *gobject.Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := valueGetMaximumValueFunction_Set()
	if err == nil {
		valueGetMaximumValueFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := gobject.ValueNewFromNative(outArgs[0].Pointer())

	return out0
}

var valueGetMinimumIncrementFunction *gi.Function
var valueGetMinimumIncrementFunction_Once sync.Once

func valueGetMinimumIncrementFunction_Set() error {
	var err error
	valueGetMinimumIncrementFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueGetMinimumIncrementFunction, err = valueInterface.InvokerNew("get_minimum_increment")
	})
	return err
}

// GetMinimumIncrement is a representation of the C type atk_value_get_minimum_increment.
func (recv *Value) GetMinimumIncrement() *gobject.Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := valueGetMinimumIncrementFunction_Set()
	if err == nil {
		valueGetMinimumIncrementFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := gobject.ValueNewFromNative(outArgs[0].Pointer())

	return out0
}

var valueGetMinimumValueFunction *gi.Function
var valueGetMinimumValueFunction_Once sync.Once

func valueGetMinimumValueFunction_Set() error {
	var err error
	valueGetMinimumValueFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueGetMinimumValueFunction, err = valueInterface.InvokerNew("get_minimum_value")
	})
	return err
}

// GetMinimumValue is a representation of the C type atk_value_get_minimum_value.
func (recv *Value) GetMinimumValue() *gobject.Value {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := valueGetMinimumValueFunction_Set()
	if err == nil {
		valueGetMinimumValueFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := gobject.ValueNewFromNative(outArgs[0].Pointer())

	return out0
}

var valueGetRangeFunction *gi.Function
var valueGetRangeFunction_Once sync.Once

func valueGetRangeFunction_Set() error {
	var err error
	valueGetRangeFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueGetRangeFunction, err = valueInterface.InvokerNew("get_range")
	})
	return err
}

// GetRange is a representation of the C type atk_value_get_range.
func (recv *Value) GetRange() *Range {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueGetRangeFunction_Set()
	if err == nil {
		ret = valueGetRangeFunction.Invoke(inArgs[:], nil)
	}

	retGo := RangeNewFromNative(ret.Pointer())

	return retGo
}

var valueGetSubRangesFunction *gi.Function
var valueGetSubRangesFunction_Once sync.Once

func valueGetSubRangesFunction_Set() error {
	var err error
	valueGetSubRangesFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueGetSubRangesFunction, err = valueInterface.InvokerNew("get_sub_ranges")
	})
	return err
}

// GetSubRanges is a representation of the C type atk_value_get_sub_ranges.
func (recv *Value) GetSubRanges() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := valueGetSubRangesFunction_Set()
	if err == nil {
		ret = valueGetSubRangesFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var valueGetValueAndTextFunction *gi.Function
var valueGetValueAndTextFunction_Once sync.Once

func valueGetValueAndTextFunction_Set() error {
	var err error
	valueGetValueAndTextFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueGetValueAndTextFunction, err = valueInterface.InvokerNew("get_value_and_text")
	})
	return err
}

// GetValueAndText is a representation of the C type atk_value_get_value_and_text.
func (recv *Value) GetValueAndText() (float64, string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument

	err := valueGetValueAndTextFunction_Set()
	if err == nil {
		valueGetValueAndTextFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].String(true)

	return out0, out1
}

var valueSetCurrentValueFunction *gi.Function
var valueSetCurrentValueFunction_Once sync.Once

func valueSetCurrentValueFunction_Set() error {
	var err error
	valueSetCurrentValueFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueSetCurrentValueFunction, err = valueInterface.InvokerNew("set_current_value")
	})
	return err
}

// SetCurrentValue is a representation of the C type atk_value_set_current_value.
func (recv *Value) SetCurrentValue(value *gobject.Value) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(value.Native())

	var ret gi.Argument

	err := valueSetCurrentValueFunction_Set()
	if err == nil {
		ret = valueSetCurrentValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueSetValueFunction *gi.Function
var valueSetValueFunction_Once sync.Once

func valueSetValueFunction_Set() error {
	var err error
	valueSetValueFunction_Once.Do(func() {
		err = valueInterface_Set()
		if err != nil {
			return
		}
		valueSetValueFunction, err = valueInterface.InvokerNew("set_value")
	})
	return err
}

// SetValue is a representation of the C type atk_value_set_value.
func (recv *Value) SetValue(newValue float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetFloat64(newValue)

	err := valueSetValueFunction_Set()
	if err == nil {
		valueSetValueFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectValueChanged connects a callback to the 'value-changed' signal of the Value.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Value) ConnectValueChanged(handler func(instance *Value, value float64, text string)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "value-changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Value) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var windowInterface *gi.Interface
var windowInterface_Once sync.Once

func windowInterface_Set() error {
	var err error
	windowInterface_Once.Do(func() {
		windowInterface, err = gi.InterfaceNew("Atk", "Window")
	})
	return err
}

type Window struct {
	native unsafe.Pointer
}

func WindowNewFromNative(native unsafe.Pointer) *Window {
	instance := &Window{native: native}

	return instance
}

/*
CastToWindow down casts any arbitrary Object to Window.
Exercise care, as this is a potentially dangerous function
if the Object is not a Window.
*/
func CastToWindow(object *gobject.Object) *Window {
	return WindowNewFromNative(object.Native())
}

// Equals compares this Window with another Window, and returns true if they represent the same Object.
func (recv *Window) Equals(other *Window) bool {
	return other.Native() == recv.Native()
}

func (recv *Window) Native() unsafe.Pointer {
	return recv.native
}

/*
ConnectActivate connects a callback to the 'activate' signal of the Window.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Window) ConnectActivate(handler func(instance *Window)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "activate", marshal)
}

/*
ConnectCreate connects a callback to the 'create' signal of the Window.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Window) ConnectCreate(handler func(instance *Window)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "create", marshal)
}

/*
ConnectDeactivate connects a callback to the 'deactivate' signal of the Window.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Window) ConnectDeactivate(handler func(instance *Window)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "deactivate", marshal)
}

/*
ConnectDestroy connects a callback to the 'destroy' signal of the Window.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Window) ConnectDestroy(handler func(instance *Window)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "destroy", marshal)
}

/*
ConnectMaximize connects a callback to the 'maximize' signal of the Window.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Window) ConnectMaximize(handler func(instance *Window)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "maximize", marshal)
}

/*
ConnectMinimize connects a callback to the 'minimize' signal of the Window.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Window) ConnectMinimize(handler func(instance *Window)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "minimize", marshal)
}

/*
ConnectMove connects a callback to the 'move' signal of the Window.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Window) ConnectMove(handler func(instance *Window)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "move", marshal)
}

/*
ConnectResize connects a callback to the 'resize' signal of the Window.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Window) ConnectResize(handler func(instance *Window)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "resize", marshal)
}

/*
ConnectRestore connects a callback to the 'restore' signal of the Window.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Window) ConnectRestore(handler func(instance *Window)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {}

	return callback.ConnectSignal(recv.Native(), "restore", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Window) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}
