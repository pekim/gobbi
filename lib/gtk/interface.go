// Code generated - DO NOT EDIT.

package gtk

import (
	callback "github.com/pekim/gobbi/internal/cgo/callback"
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	atk "github.com/pekim/gobbi/lib/atk"
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

var actionableInterface *gi.Interface
var actionableInterface_Once sync.Once

func actionableInterface_Set() error {
	var err error
	actionableInterface_Once.Do(func() {
		actionableInterface, err = gi.InterfaceNew("Gtk", "Actionable")
	})
	return err
}

type Actionable struct {
	native unsafe.Pointer
}

func ActionableNewFromNative(native unsafe.Pointer) *Actionable {
	instance := &Actionable{native: native}

	return instance
}

/*
CastToActionable down casts any arbitrary Object to Actionable.
Exercise care, as this is a potentially dangerous function
if the Object is not a Actionable.
*/
func CastToActionable(object *gobject.Object) *Actionable {
	return ActionableNewFromNative(object.Native())
}

// Equals compares this Actionable with another Actionable, and returns true if they represent the same Object.
func (recv *Actionable) Equals(other *Actionable) bool {
	return other.Native() == recv.Native()
}

func (recv *Actionable) Native() unsafe.Pointer {
	return recv.native
}

var actionableGetActionNameFunction *gi.Function
var actionableGetActionNameFunction_Once sync.Once

func actionableGetActionNameFunction_Set() error {
	var err error
	actionableGetActionNameFunction_Once.Do(func() {
		err = actionableInterface_Set()
		if err != nil {
			return
		}
		actionableGetActionNameFunction, err = actionableInterface.InvokerNew("get_action_name")
	})
	return err
}

// GetActionName is a representation of the C type gtk_actionable_get_action_name.
func (recv *Actionable) GetActionName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := actionableGetActionNameFunction_Set()
	if err == nil {
		ret = actionableGetActionNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var actionableGetActionTargetValueFunction *gi.Function
var actionableGetActionTargetValueFunction_Once sync.Once

func actionableGetActionTargetValueFunction_Set() error {
	var err error
	actionableGetActionTargetValueFunction_Once.Do(func() {
		err = actionableInterface_Set()
		if err != nil {
			return
		}
		actionableGetActionTargetValueFunction, err = actionableInterface.InvokerNew("get_action_target_value")
	})
	return err
}

// GetActionTargetValue is a representation of the C type gtk_actionable_get_action_target_value.
func (recv *Actionable) GetActionTargetValue() *glib.Variant {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := actionableGetActionTargetValueFunction_Set()
	if err == nil {
		ret = actionableGetActionTargetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.VariantNewFromNative(ret.Pointer())

	return retGo
}

var actionableSetActionNameFunction *gi.Function
var actionableSetActionNameFunction_Once sync.Once

func actionableSetActionNameFunction_Set() error {
	var err error
	actionableSetActionNameFunction_Once.Do(func() {
		err = actionableInterface_Set()
		if err != nil {
			return
		}
		actionableSetActionNameFunction, err = actionableInterface.InvokerNew("set_action_name")
	})
	return err
}

// SetActionName is a representation of the C type gtk_actionable_set_action_name.
func (recv *Actionable) SetActionName(actionName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(actionName)

	err := actionableSetActionNameFunction_Set()
	if err == nil {
		actionableSetActionNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_actionable_set_action_target' : parameter '...' of type 'nil' not supported

var actionableSetActionTargetValueFunction *gi.Function
var actionableSetActionTargetValueFunction_Once sync.Once

func actionableSetActionTargetValueFunction_Set() error {
	var err error
	actionableSetActionTargetValueFunction_Once.Do(func() {
		err = actionableInterface_Set()
		if err != nil {
			return
		}
		actionableSetActionTargetValueFunction, err = actionableInterface.InvokerNew("set_action_target_value")
	})
	return err
}

// SetActionTargetValue is a representation of the C type gtk_actionable_set_action_target_value.
func (recv *Actionable) SetActionTargetValue(targetValue *glib.Variant) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(targetValue.Native())

	err := actionableSetActionTargetValueFunction_Set()
	if err == nil {
		actionableSetActionTargetValueFunction.Invoke(inArgs[:], nil)
	}

	return
}

var actionableSetDetailedActionNameFunction *gi.Function
var actionableSetDetailedActionNameFunction_Once sync.Once

func actionableSetDetailedActionNameFunction_Set() error {
	var err error
	actionableSetDetailedActionNameFunction_Once.Do(func() {
		err = actionableInterface_Set()
		if err != nil {
			return
		}
		actionableSetDetailedActionNameFunction, err = actionableInterface.InvokerNew("set_detailed_action_name")
	})
	return err
}

// SetDetailedActionName is a representation of the C type gtk_actionable_set_detailed_action_name.
func (recv *Actionable) SetDetailedActionName(detailedActionName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(detailedActionName)

	err := actionableSetDetailedActionNameFunction_Set()
	if err == nil {
		actionableSetDetailedActionNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var activatableInterface *gi.Interface
var activatableInterface_Once sync.Once

func activatableInterface_Set() error {
	var err error
	activatableInterface_Once.Do(func() {
		activatableInterface, err = gi.InterfaceNew("Gtk", "Activatable")
	})
	return err
}

type Activatable struct {
	native unsafe.Pointer
}

func ActivatableNewFromNative(native unsafe.Pointer) *Activatable {
	instance := &Activatable{native: native}

	return instance
}

/*
CastToActivatable down casts any arbitrary Object to Activatable.
Exercise care, as this is a potentially dangerous function
if the Object is not a Activatable.
*/
func CastToActivatable(object *gobject.Object) *Activatable {
	return ActivatableNewFromNative(object.Native())
}

// Equals compares this Activatable with another Activatable, and returns true if they represent the same Object.
func (recv *Activatable) Equals(other *Activatable) bool {
	return other.Native() == recv.Native()
}

func (recv *Activatable) Native() unsafe.Pointer {
	return recv.native
}

var activatableDoSetRelatedActionFunction *gi.Function
var activatableDoSetRelatedActionFunction_Once sync.Once

func activatableDoSetRelatedActionFunction_Set() error {
	var err error
	activatableDoSetRelatedActionFunction_Once.Do(func() {
		err = activatableInterface_Set()
		if err != nil {
			return
		}
		activatableDoSetRelatedActionFunction, err = activatableInterface.InvokerNew("do_set_related_action")
	})
	return err
}

// DoSetRelatedAction is a representation of the C type gtk_activatable_do_set_related_action.
func (recv *Activatable) DoSetRelatedAction(action *Action) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(action.Native())

	err := activatableDoSetRelatedActionFunction_Set()
	if err == nil {
		activatableDoSetRelatedActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var activatableGetRelatedActionFunction *gi.Function
var activatableGetRelatedActionFunction_Once sync.Once

func activatableGetRelatedActionFunction_Set() error {
	var err error
	activatableGetRelatedActionFunction_Once.Do(func() {
		err = activatableInterface_Set()
		if err != nil {
			return
		}
		activatableGetRelatedActionFunction, err = activatableInterface.InvokerNew("get_related_action")
	})
	return err
}

// GetRelatedAction is a representation of the C type gtk_activatable_get_related_action.
func (recv *Activatable) GetRelatedAction() *Action {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := activatableGetRelatedActionFunction_Set()
	if err == nil {
		ret = activatableGetRelatedActionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ActionNewFromNative(ret.Pointer())

	return retGo
}

var activatableGetUseActionAppearanceFunction *gi.Function
var activatableGetUseActionAppearanceFunction_Once sync.Once

func activatableGetUseActionAppearanceFunction_Set() error {
	var err error
	activatableGetUseActionAppearanceFunction_Once.Do(func() {
		err = activatableInterface_Set()
		if err != nil {
			return
		}
		activatableGetUseActionAppearanceFunction, err = activatableInterface.InvokerNew("get_use_action_appearance")
	})
	return err
}

// GetUseActionAppearance is a representation of the C type gtk_activatable_get_use_action_appearance.
func (recv *Activatable) GetUseActionAppearance() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := activatableGetUseActionAppearanceFunction_Set()
	if err == nil {
		ret = activatableGetUseActionAppearanceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var activatableSetRelatedActionFunction *gi.Function
var activatableSetRelatedActionFunction_Once sync.Once

func activatableSetRelatedActionFunction_Set() error {
	var err error
	activatableSetRelatedActionFunction_Once.Do(func() {
		err = activatableInterface_Set()
		if err != nil {
			return
		}
		activatableSetRelatedActionFunction, err = activatableInterface.InvokerNew("set_related_action")
	})
	return err
}

// SetRelatedAction is a representation of the C type gtk_activatable_set_related_action.
func (recv *Activatable) SetRelatedAction(action *Action) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(action.Native())

	err := activatableSetRelatedActionFunction_Set()
	if err == nil {
		activatableSetRelatedActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var activatableSetUseActionAppearanceFunction *gi.Function
var activatableSetUseActionAppearanceFunction_Once sync.Once

func activatableSetUseActionAppearanceFunction_Set() error {
	var err error
	activatableSetUseActionAppearanceFunction_Once.Do(func() {
		err = activatableInterface_Set()
		if err != nil {
			return
		}
		activatableSetUseActionAppearanceFunction, err = activatableInterface.InvokerNew("set_use_action_appearance")
	})
	return err
}

// SetUseActionAppearance is a representation of the C type gtk_activatable_set_use_action_appearance.
func (recv *Activatable) SetUseActionAppearance(useAppearance bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(useAppearance)

	err := activatableSetUseActionAppearanceFunction_Set()
	if err == nil {
		activatableSetUseActionAppearanceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var activatableSyncActionPropertiesFunction *gi.Function
var activatableSyncActionPropertiesFunction_Once sync.Once

func activatableSyncActionPropertiesFunction_Set() error {
	var err error
	activatableSyncActionPropertiesFunction_Once.Do(func() {
		err = activatableInterface_Set()
		if err != nil {
			return
		}
		activatableSyncActionPropertiesFunction, err = activatableInterface.InvokerNew("sync_action_properties")
	})
	return err
}

// SyncActionProperties is a representation of the C type gtk_activatable_sync_action_properties.
func (recv *Activatable) SyncActionProperties(action *Action) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(action.Native())

	err := activatableSyncActionPropertiesFunction_Set()
	if err == nil {
		activatableSyncActionPropertiesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var appChooserInterface *gi.Interface
var appChooserInterface_Once sync.Once

func appChooserInterface_Set() error {
	var err error
	appChooserInterface_Once.Do(func() {
		appChooserInterface, err = gi.InterfaceNew("Gtk", "AppChooser")
	})
	return err
}

type AppChooser struct {
	native unsafe.Pointer
}

func AppChooserNewFromNative(native unsafe.Pointer) *AppChooser {
	instance := &AppChooser{native: native}

	return instance
}

/*
CastToAppChooser down casts any arbitrary Object to AppChooser.
Exercise care, as this is a potentially dangerous function
if the Object is not a AppChooser.
*/
func CastToAppChooser(object *gobject.Object) *AppChooser {
	return AppChooserNewFromNative(object.Native())
}

// Equals compares this AppChooser with another AppChooser, and returns true if they represent the same Object.
func (recv *AppChooser) Equals(other *AppChooser) bool {
	return other.Native() == recv.Native()
}

func (recv *AppChooser) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'gtk_app_chooser_get_app_info' : return type 'Gio.AppInfo' not supported

var appChooserGetContentTypeFunction *gi.Function
var appChooserGetContentTypeFunction_Once sync.Once

func appChooserGetContentTypeFunction_Set() error {
	var err error
	appChooserGetContentTypeFunction_Once.Do(func() {
		err = appChooserInterface_Set()
		if err != nil {
			return
		}
		appChooserGetContentTypeFunction, err = appChooserInterface.InvokerNew("get_content_type")
	})
	return err
}

// GetContentType is a representation of the C type gtk_app_chooser_get_content_type.
func (recv *AppChooser) GetContentType() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := appChooserGetContentTypeFunction_Set()
	if err == nil {
		ret = appChooserGetContentTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var appChooserRefreshFunction *gi.Function
var appChooserRefreshFunction_Once sync.Once

func appChooserRefreshFunction_Set() error {
	var err error
	appChooserRefreshFunction_Once.Do(func() {
		err = appChooserInterface_Set()
		if err != nil {
			return
		}
		appChooserRefreshFunction, err = appChooserInterface.InvokerNew("refresh")
	})
	return err
}

// Refresh is a representation of the C type gtk_app_chooser_refresh.
func (recv *AppChooser) Refresh() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := appChooserRefreshFunction_Set()
	if err == nil {
		appChooserRefreshFunction.Invoke(inArgs[:], nil)
	}

	return
}

var buildableInterface *gi.Interface
var buildableInterface_Once sync.Once

func buildableInterface_Set() error {
	var err error
	buildableInterface_Once.Do(func() {
		buildableInterface, err = gi.InterfaceNew("Gtk", "Buildable")
	})
	return err
}

type Buildable struct {
	native unsafe.Pointer
}

func BuildableNewFromNative(native unsafe.Pointer) *Buildable {
	instance := &Buildable{native: native}

	return instance
}

/*
CastToBuildable down casts any arbitrary Object to Buildable.
Exercise care, as this is a potentially dangerous function
if the Object is not a Buildable.
*/
func CastToBuildable(object *gobject.Object) *Buildable {
	return BuildableNewFromNative(object.Native())
}

// Equals compares this Buildable with another Buildable, and returns true if they represent the same Object.
func (recv *Buildable) Equals(other *Buildable) bool {
	return other.Native() == recv.Native()
}

func (recv *Buildable) Native() unsafe.Pointer {
	return recv.native
}

var buildableAddChildFunction *gi.Function
var buildableAddChildFunction_Once sync.Once

func buildableAddChildFunction_Set() error {
	var err error
	buildableAddChildFunction_Once.Do(func() {
		err = buildableInterface_Set()
		if err != nil {
			return
		}
		buildableAddChildFunction, err = buildableInterface.InvokerNew("add_child")
	})
	return err
}

// AddChild is a representation of the C type gtk_buildable_add_child.
func (recv *Buildable) AddChild(builder *Builder, child *gobject.Object, type_ string) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(builder.Native())
	inArgs[2].SetPointer(child.Native())
	inArgs[3].SetString(type_)

	err := buildableAddChildFunction_Set()
	if err == nil {
		buildableAddChildFunction.Invoke(inArgs[:], nil)
	}

	return
}

var buildableConstructChildFunction *gi.Function
var buildableConstructChildFunction_Once sync.Once

func buildableConstructChildFunction_Set() error {
	var err error
	buildableConstructChildFunction_Once.Do(func() {
		err = buildableInterface_Set()
		if err != nil {
			return
		}
		buildableConstructChildFunction, err = buildableInterface.InvokerNew("construct_child")
	})
	return err
}

// ConstructChild is a representation of the C type gtk_buildable_construct_child.
func (recv *Buildable) ConstructChild(builder *Builder, name string) *gobject.Object {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(builder.Native())
	inArgs[2].SetString(name)

	var ret gi.Argument

	err := buildableConstructChildFunction_Set()
	if err == nil {
		ret = buildableConstructChildFunction.Invoke(inArgs[:], nil)
	}

	retGo := gobject.ObjectNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_buildable_custom_finished' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'gtk_buildable_custom_tag_end' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'gtk_buildable_custom_tag_start' : parameter 'data' of type 'gpointer' not supported

var buildableGetInternalChildFunction *gi.Function
var buildableGetInternalChildFunction_Once sync.Once

func buildableGetInternalChildFunction_Set() error {
	var err error
	buildableGetInternalChildFunction_Once.Do(func() {
		err = buildableInterface_Set()
		if err != nil {
			return
		}
		buildableGetInternalChildFunction, err = buildableInterface.InvokerNew("get_internal_child")
	})
	return err
}

// GetInternalChild is a representation of the C type gtk_buildable_get_internal_child.
func (recv *Buildable) GetInternalChild(builder *Builder, childname string) *gobject.Object {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(builder.Native())
	inArgs[2].SetString(childname)

	var ret gi.Argument

	err := buildableGetInternalChildFunction_Set()
	if err == nil {
		ret = buildableGetInternalChildFunction.Invoke(inArgs[:], nil)
	}

	retGo := gobject.ObjectNewFromNative(ret.Pointer())

	return retGo
}

var buildableGetNameFunction *gi.Function
var buildableGetNameFunction_Once sync.Once

func buildableGetNameFunction_Set() error {
	var err error
	buildableGetNameFunction_Once.Do(func() {
		err = buildableInterface_Set()
		if err != nil {
			return
		}
		buildableGetNameFunction, err = buildableInterface.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gtk_buildable_get_name.
func (recv *Buildable) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := buildableGetNameFunction_Set()
	if err == nil {
		ret = buildableGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var buildableParserFinishedFunction *gi.Function
var buildableParserFinishedFunction_Once sync.Once

func buildableParserFinishedFunction_Set() error {
	var err error
	buildableParserFinishedFunction_Once.Do(func() {
		err = buildableInterface_Set()
		if err != nil {
			return
		}
		buildableParserFinishedFunction, err = buildableInterface.InvokerNew("parser_finished")
	})
	return err
}

// ParserFinished is a representation of the C type gtk_buildable_parser_finished.
func (recv *Buildable) ParserFinished(builder *Builder) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(builder.Native())

	err := buildableParserFinishedFunction_Set()
	if err == nil {
		buildableParserFinishedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var buildableSetBuildablePropertyFunction *gi.Function
var buildableSetBuildablePropertyFunction_Once sync.Once

func buildableSetBuildablePropertyFunction_Set() error {
	var err error
	buildableSetBuildablePropertyFunction_Once.Do(func() {
		err = buildableInterface_Set()
		if err != nil {
			return
		}
		buildableSetBuildablePropertyFunction, err = buildableInterface.InvokerNew("set_buildable_property")
	})
	return err
}

// SetBuildableProperty is a representation of the C type gtk_buildable_set_buildable_property.
func (recv *Buildable) SetBuildableProperty(builder *Builder, name string, value *gobject.Value) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(builder.Native())
	inArgs[2].SetString(name)
	inArgs[3].SetPointer(value.Native())

	err := buildableSetBuildablePropertyFunction_Set()
	if err == nil {
		buildableSetBuildablePropertyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var buildableSetNameFunction *gi.Function
var buildableSetNameFunction_Once sync.Once

func buildableSetNameFunction_Set() error {
	var err error
	buildableSetNameFunction_Once.Do(func() {
		err = buildableInterface_Set()
		if err != nil {
			return
		}
		buildableSetNameFunction, err = buildableInterface.InvokerNew("set_name")
	})
	return err
}

// SetName is a representation of the C type gtk_buildable_set_name.
func (recv *Buildable) SetName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	err := buildableSetNameFunction_Set()
	if err == nil {
		buildableSetNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellAccessibleParentInterface *gi.Interface
var cellAccessibleParentInterface_Once sync.Once

func cellAccessibleParentInterface_Set() error {
	var err error
	cellAccessibleParentInterface_Once.Do(func() {
		cellAccessibleParentInterface, err = gi.InterfaceNew("Gtk", "CellAccessibleParent")
	})
	return err
}

type CellAccessibleParent struct {
	native unsafe.Pointer
}

func CellAccessibleParentNewFromNative(native unsafe.Pointer) *CellAccessibleParent {
	instance := &CellAccessibleParent{native: native}

	return instance
}

/*
CastToCellAccessibleParent down casts any arbitrary Object to CellAccessibleParent.
Exercise care, as this is a potentially dangerous function
if the Object is not a CellAccessibleParent.
*/
func CastToCellAccessibleParent(object *gobject.Object) *CellAccessibleParent {
	return CellAccessibleParentNewFromNative(object.Native())
}

// Equals compares this CellAccessibleParent with another CellAccessibleParent, and returns true if they represent the same Object.
func (recv *CellAccessibleParent) Equals(other *CellAccessibleParent) bool {
	return other.Native() == recv.Native()
}

func (recv *CellAccessibleParent) Native() unsafe.Pointer {
	return recv.native
}

var cellAccessibleParentActivateFunction *gi.Function
var cellAccessibleParentActivateFunction_Once sync.Once

func cellAccessibleParentActivateFunction_Set() error {
	var err error
	cellAccessibleParentActivateFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentActivateFunction, err = cellAccessibleParentInterface.InvokerNew("activate")
	})
	return err
}

// Activate is a representation of the C type gtk_cell_accessible_parent_activate.
func (recv *CellAccessibleParent) Activate(cell *CellAccessible) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	err := cellAccessibleParentActivateFunction_Set()
	if err == nil {
		cellAccessibleParentActivateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellAccessibleParentEditFunction *gi.Function
var cellAccessibleParentEditFunction_Once sync.Once

func cellAccessibleParentEditFunction_Set() error {
	var err error
	cellAccessibleParentEditFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentEditFunction, err = cellAccessibleParentInterface.InvokerNew("edit")
	})
	return err
}

// Edit is a representation of the C type gtk_cell_accessible_parent_edit.
func (recv *CellAccessibleParent) Edit(cell *CellAccessible) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	err := cellAccessibleParentEditFunction_Set()
	if err == nil {
		cellAccessibleParentEditFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellAccessibleParentExpandCollapseFunction *gi.Function
var cellAccessibleParentExpandCollapseFunction_Once sync.Once

func cellAccessibleParentExpandCollapseFunction_Set() error {
	var err error
	cellAccessibleParentExpandCollapseFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentExpandCollapseFunction, err = cellAccessibleParentInterface.InvokerNew("expand_collapse")
	})
	return err
}

// ExpandCollapse is a representation of the C type gtk_cell_accessible_parent_expand_collapse.
func (recv *CellAccessibleParent) ExpandCollapse(cell *CellAccessible) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	err := cellAccessibleParentExpandCollapseFunction_Set()
	if err == nil {
		cellAccessibleParentExpandCollapseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellAccessibleParentGetCellAreaFunction *gi.Function
var cellAccessibleParentGetCellAreaFunction_Once sync.Once

func cellAccessibleParentGetCellAreaFunction_Set() error {
	var err error
	cellAccessibleParentGetCellAreaFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentGetCellAreaFunction, err = cellAccessibleParentInterface.InvokerNew("get_cell_area")
	})
	return err
}

// GetCellArea is a representation of the C type gtk_cell_accessible_parent_get_cell_area.
func (recv *CellAccessibleParent) GetCellArea(cell *CellAccessible) *gdk.Rectangle {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	var outArgs [1]gi.Argument

	err := cellAccessibleParentGetCellAreaFunction_Set()
	if err == nil {
		cellAccessibleParentGetCellAreaFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := gdk.RectangleNewFromNative(outArgs[0].Pointer())

	return out0
}

// UNSUPPORTED : C value 'gtk_cell_accessible_parent_get_cell_extents' : parameter 'coord_type' of type 'Atk.CoordType' not supported

var cellAccessibleParentGetCellPositionFunction *gi.Function
var cellAccessibleParentGetCellPositionFunction_Once sync.Once

func cellAccessibleParentGetCellPositionFunction_Set() error {
	var err error
	cellAccessibleParentGetCellPositionFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentGetCellPositionFunction, err = cellAccessibleParentInterface.InvokerNew("get_cell_position")
	})
	return err
}

// GetCellPosition is a representation of the C type gtk_cell_accessible_parent_get_cell_position.
func (recv *CellAccessibleParent) GetCellPosition(cell *CellAccessible) (int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	var outArgs [2]gi.Argument

	err := cellAccessibleParentGetCellPositionFunction_Set()
	if err == nil {
		cellAccessibleParentGetCellPositionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var cellAccessibleParentGetChildIndexFunction *gi.Function
var cellAccessibleParentGetChildIndexFunction_Once sync.Once

func cellAccessibleParentGetChildIndexFunction_Set() error {
	var err error
	cellAccessibleParentGetChildIndexFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentGetChildIndexFunction, err = cellAccessibleParentInterface.InvokerNew("get_child_index")
	})
	return err
}

// GetChildIndex is a representation of the C type gtk_cell_accessible_parent_get_child_index.
func (recv *CellAccessibleParent) GetChildIndex(cell *CellAccessible) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	var ret gi.Argument

	err := cellAccessibleParentGetChildIndexFunction_Set()
	if err == nil {
		ret = cellAccessibleParentGetChildIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var cellAccessibleParentGetColumnHeaderCellsFunction *gi.Function
var cellAccessibleParentGetColumnHeaderCellsFunction_Once sync.Once

func cellAccessibleParentGetColumnHeaderCellsFunction_Set() error {
	var err error
	cellAccessibleParentGetColumnHeaderCellsFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentGetColumnHeaderCellsFunction, err = cellAccessibleParentInterface.InvokerNew("get_column_header_cells")
	})
	return err
}

// GetColumnHeaderCells is a representation of the C type gtk_cell_accessible_parent_get_column_header_cells.
func (recv *CellAccessibleParent) GetColumnHeaderCells(cell *CellAccessible) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	err := cellAccessibleParentGetColumnHeaderCellsFunction_Set()
	if err == nil {
		cellAccessibleParentGetColumnHeaderCellsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellAccessibleParentGetRendererStateFunction *gi.Function
var cellAccessibleParentGetRendererStateFunction_Once sync.Once

func cellAccessibleParentGetRendererStateFunction_Set() error {
	var err error
	cellAccessibleParentGetRendererStateFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentGetRendererStateFunction, err = cellAccessibleParentInterface.InvokerNew("get_renderer_state")
	})
	return err
}

// GetRendererState is a representation of the C type gtk_cell_accessible_parent_get_renderer_state.
func (recv *CellAccessibleParent) GetRendererState(cell *CellAccessible) CellRendererState {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	var ret gi.Argument

	err := cellAccessibleParentGetRendererStateFunction_Set()
	if err == nil {
		ret = cellAccessibleParentGetRendererStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := CellRendererState(ret.Int32())

	return retGo
}

var cellAccessibleParentGetRowHeaderCellsFunction *gi.Function
var cellAccessibleParentGetRowHeaderCellsFunction_Once sync.Once

func cellAccessibleParentGetRowHeaderCellsFunction_Set() error {
	var err error
	cellAccessibleParentGetRowHeaderCellsFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentGetRowHeaderCellsFunction, err = cellAccessibleParentInterface.InvokerNew("get_row_header_cells")
	})
	return err
}

// GetRowHeaderCells is a representation of the C type gtk_cell_accessible_parent_get_row_header_cells.
func (recv *CellAccessibleParent) GetRowHeaderCells(cell *CellAccessible) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	err := cellAccessibleParentGetRowHeaderCellsFunction_Set()
	if err == nil {
		cellAccessibleParentGetRowHeaderCellsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellAccessibleParentGrabFocusFunction *gi.Function
var cellAccessibleParentGrabFocusFunction_Once sync.Once

func cellAccessibleParentGrabFocusFunction_Set() error {
	var err error
	cellAccessibleParentGrabFocusFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentGrabFocusFunction, err = cellAccessibleParentInterface.InvokerNew("grab_focus")
	})
	return err
}

// GrabFocus is a representation of the C type gtk_cell_accessible_parent_grab_focus.
func (recv *CellAccessibleParent) GrabFocus(cell *CellAccessible) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	var ret gi.Argument

	err := cellAccessibleParentGrabFocusFunction_Set()
	if err == nil {
		ret = cellAccessibleParentGrabFocusFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var cellAccessibleParentUpdateRelationsetFunction *gi.Function
var cellAccessibleParentUpdateRelationsetFunction_Once sync.Once

func cellAccessibleParentUpdateRelationsetFunction_Set() error {
	var err error
	cellAccessibleParentUpdateRelationsetFunction_Once.Do(func() {
		err = cellAccessibleParentInterface_Set()
		if err != nil {
			return
		}
		cellAccessibleParentUpdateRelationsetFunction, err = cellAccessibleParentInterface.InvokerNew("update_relationset")
	})
	return err
}

// UpdateRelationset is a representation of the C type gtk_cell_accessible_parent_update_relationset.
func (recv *CellAccessibleParent) UpdateRelationset(cell *CellAccessible, relationset *atk.RelationSet) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())
	inArgs[2].SetPointer(relationset.Native())

	err := cellAccessibleParentUpdateRelationsetFunction_Set()
	if err == nil {
		cellAccessibleParentUpdateRelationsetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellEditableInterface *gi.Interface
var cellEditableInterface_Once sync.Once

func cellEditableInterface_Set() error {
	var err error
	cellEditableInterface_Once.Do(func() {
		cellEditableInterface, err = gi.InterfaceNew("Gtk", "CellEditable")
	})
	return err
}

type CellEditable struct {
	native unsafe.Pointer
}

func CellEditableNewFromNative(native unsafe.Pointer) *CellEditable {
	instance := &CellEditable{native: native}

	return instance
}

/*
CastToCellEditable down casts any arbitrary Object to CellEditable.
Exercise care, as this is a potentially dangerous function
if the Object is not a CellEditable.
*/
func CastToCellEditable(object *gobject.Object) *CellEditable {
	return CellEditableNewFromNative(object.Native())
}

// Equals compares this CellEditable with another CellEditable, and returns true if they represent the same Object.
func (recv *CellEditable) Equals(other *CellEditable) bool {
	return other.Native() == recv.Native()
}

func (recv *CellEditable) Native() unsafe.Pointer {
	return recv.native
}

var cellEditableEditingDoneFunction *gi.Function
var cellEditableEditingDoneFunction_Once sync.Once

func cellEditableEditingDoneFunction_Set() error {
	var err error
	cellEditableEditingDoneFunction_Once.Do(func() {
		err = cellEditableInterface_Set()
		if err != nil {
			return
		}
		cellEditableEditingDoneFunction, err = cellEditableInterface.InvokerNew("editing_done")
	})
	return err
}

// EditingDone is a representation of the C type gtk_cell_editable_editing_done.
func (recv *CellEditable) EditingDone() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cellEditableEditingDoneFunction_Set()
	if err == nil {
		cellEditableEditingDoneFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellEditableRemoveWidgetFunction *gi.Function
var cellEditableRemoveWidgetFunction_Once sync.Once

func cellEditableRemoveWidgetFunction_Set() error {
	var err error
	cellEditableRemoveWidgetFunction_Once.Do(func() {
		err = cellEditableInterface_Set()
		if err != nil {
			return
		}
		cellEditableRemoveWidgetFunction, err = cellEditableInterface.InvokerNew("remove_widget")
	})
	return err
}

// RemoveWidget is a representation of the C type gtk_cell_editable_remove_widget.
func (recv *CellEditable) RemoveWidget() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cellEditableRemoveWidgetFunction_Set()
	if err == nil {
		cellEditableRemoveWidgetFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_cell_editable_start_editing' : parameter 'event' of type 'Gdk.Event' not supported

/*
ConnectEditingDone connects a callback to the 'editing-done' signal of the CellEditable.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *CellEditable) ConnectEditingDone(handler func(instance *CellEditable)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := CellEditableNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "editing-done", marshal)
}

/*
ConnectRemoveWidget connects a callback to the 'remove-widget' signal of the CellEditable.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *CellEditable) ConnectRemoveWidget(handler func(instance *CellEditable)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := CellEditableNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "remove-widget", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *CellEditable) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var cellLayoutInterface *gi.Interface
var cellLayoutInterface_Once sync.Once

func cellLayoutInterface_Set() error {
	var err error
	cellLayoutInterface_Once.Do(func() {
		cellLayoutInterface, err = gi.InterfaceNew("Gtk", "CellLayout")
	})
	return err
}

type CellLayout struct {
	native unsafe.Pointer
}

func CellLayoutNewFromNative(native unsafe.Pointer) *CellLayout {
	instance := &CellLayout{native: native}

	return instance
}

/*
CastToCellLayout down casts any arbitrary Object to CellLayout.
Exercise care, as this is a potentially dangerous function
if the Object is not a CellLayout.
*/
func CastToCellLayout(object *gobject.Object) *CellLayout {
	return CellLayoutNewFromNative(object.Native())
}

// Equals compares this CellLayout with another CellLayout, and returns true if they represent the same Object.
func (recv *CellLayout) Equals(other *CellLayout) bool {
	return other.Native() == recv.Native()
}

func (recv *CellLayout) Native() unsafe.Pointer {
	return recv.native
}

var cellLayoutAddAttributeFunction *gi.Function
var cellLayoutAddAttributeFunction_Once sync.Once

func cellLayoutAddAttributeFunction_Set() error {
	var err error
	cellLayoutAddAttributeFunction_Once.Do(func() {
		err = cellLayoutInterface_Set()
		if err != nil {
			return
		}
		cellLayoutAddAttributeFunction, err = cellLayoutInterface.InvokerNew("add_attribute")
	})
	return err
}

// AddAttribute is a representation of the C type gtk_cell_layout_add_attribute.
func (recv *CellLayout) AddAttribute(cell *CellRenderer, attribute string, column int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())
	inArgs[2].SetString(attribute)
	inArgs[3].SetInt32(column)

	err := cellLayoutAddAttributeFunction_Set()
	if err == nil {
		cellLayoutAddAttributeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellLayoutClearFunction *gi.Function
var cellLayoutClearFunction_Once sync.Once

func cellLayoutClearFunction_Set() error {
	var err error
	cellLayoutClearFunction_Once.Do(func() {
		err = cellLayoutInterface_Set()
		if err != nil {
			return
		}
		cellLayoutClearFunction, err = cellLayoutInterface.InvokerNew("clear")
	})
	return err
}

// Clear is a representation of the C type gtk_cell_layout_clear.
func (recv *CellLayout) Clear() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := cellLayoutClearFunction_Set()
	if err == nil {
		cellLayoutClearFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellLayoutClearAttributesFunction *gi.Function
var cellLayoutClearAttributesFunction_Once sync.Once

func cellLayoutClearAttributesFunction_Set() error {
	var err error
	cellLayoutClearAttributesFunction_Once.Do(func() {
		err = cellLayoutInterface_Set()
		if err != nil {
			return
		}
		cellLayoutClearAttributesFunction, err = cellLayoutInterface.InvokerNew("clear_attributes")
	})
	return err
}

// ClearAttributes is a representation of the C type gtk_cell_layout_clear_attributes.
func (recv *CellLayout) ClearAttributes(cell *CellRenderer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())

	err := cellLayoutClearAttributesFunction_Set()
	if err == nil {
		cellLayoutClearAttributesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellLayoutGetAreaFunction *gi.Function
var cellLayoutGetAreaFunction_Once sync.Once

func cellLayoutGetAreaFunction_Set() error {
	var err error
	cellLayoutGetAreaFunction_Once.Do(func() {
		err = cellLayoutInterface_Set()
		if err != nil {
			return
		}
		cellLayoutGetAreaFunction, err = cellLayoutInterface.InvokerNew("get_area")
	})
	return err
}

// GetArea is a representation of the C type gtk_cell_layout_get_area.
func (recv *CellLayout) GetArea() *CellArea {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cellLayoutGetAreaFunction_Set()
	if err == nil {
		ret = cellLayoutGetAreaFunction.Invoke(inArgs[:], nil)
	}

	retGo := CellAreaNewFromNative(ret.Pointer())

	return retGo
}

var cellLayoutGetCellsFunction *gi.Function
var cellLayoutGetCellsFunction_Once sync.Once

func cellLayoutGetCellsFunction_Set() error {
	var err error
	cellLayoutGetCellsFunction_Once.Do(func() {
		err = cellLayoutInterface_Set()
		if err != nil {
			return
		}
		cellLayoutGetCellsFunction, err = cellLayoutInterface.InvokerNew("get_cells")
	})
	return err
}

// GetCells is a representation of the C type gtk_cell_layout_get_cells.
func (recv *CellLayout) GetCells() *glib.List {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := cellLayoutGetCellsFunction_Set()
	if err == nil {
		ret = cellLayoutGetCellsFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var cellLayoutPackEndFunction *gi.Function
var cellLayoutPackEndFunction_Once sync.Once

func cellLayoutPackEndFunction_Set() error {
	var err error
	cellLayoutPackEndFunction_Once.Do(func() {
		err = cellLayoutInterface_Set()
		if err != nil {
			return
		}
		cellLayoutPackEndFunction, err = cellLayoutInterface.InvokerNew("pack_end")
	})
	return err
}

// PackEnd is a representation of the C type gtk_cell_layout_pack_end.
func (recv *CellLayout) PackEnd(cell *CellRenderer, expand bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())
	inArgs[2].SetBoolean(expand)

	err := cellLayoutPackEndFunction_Set()
	if err == nil {
		cellLayoutPackEndFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellLayoutPackStartFunction *gi.Function
var cellLayoutPackStartFunction_Once sync.Once

func cellLayoutPackStartFunction_Set() error {
	var err error
	cellLayoutPackStartFunction_Once.Do(func() {
		err = cellLayoutInterface_Set()
		if err != nil {
			return
		}
		cellLayoutPackStartFunction, err = cellLayoutInterface.InvokerNew("pack_start")
	})
	return err
}

// PackStart is a representation of the C type gtk_cell_layout_pack_start.
func (recv *CellLayout) PackStart(cell *CellRenderer, expand bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())
	inArgs[2].SetBoolean(expand)

	err := cellLayoutPackStartFunction_Set()
	if err == nil {
		cellLayoutPackStartFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cellLayoutReorderFunction *gi.Function
var cellLayoutReorderFunction_Once sync.Once

func cellLayoutReorderFunction_Set() error {
	var err error
	cellLayoutReorderFunction_Once.Do(func() {
		err = cellLayoutInterface_Set()
		if err != nil {
			return
		}
		cellLayoutReorderFunction, err = cellLayoutInterface.InvokerNew("reorder")
	})
	return err
}

// Reorder is a representation of the C type gtk_cell_layout_reorder.
func (recv *CellLayout) Reorder(cell *CellRenderer, position int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(cell.Native())
	inArgs[2].SetInt32(position)

	err := cellLayoutReorderFunction_Set()
	if err == nil {
		cellLayoutReorderFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_cell_layout_set_attributes' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_cell_layout_set_cell_data_func' : parameter 'func' of type 'CellLayoutDataFunc' not supported

var colorChooserInterface *gi.Interface
var colorChooserInterface_Once sync.Once

func colorChooserInterface_Set() error {
	var err error
	colorChooserInterface_Once.Do(func() {
		colorChooserInterface, err = gi.InterfaceNew("Gtk", "ColorChooser")
	})
	return err
}

type ColorChooser struct {
	native unsafe.Pointer
}

func ColorChooserNewFromNative(native unsafe.Pointer) *ColorChooser {
	instance := &ColorChooser{native: native}

	return instance
}

/*
CastToColorChooser down casts any arbitrary Object to ColorChooser.
Exercise care, as this is a potentially dangerous function
if the Object is not a ColorChooser.
*/
func CastToColorChooser(object *gobject.Object) *ColorChooser {
	return ColorChooserNewFromNative(object.Native())
}

// Equals compares this ColorChooser with another ColorChooser, and returns true if they represent the same Object.
func (recv *ColorChooser) Equals(other *ColorChooser) bool {
	return other.Native() == recv.Native()
}

func (recv *ColorChooser) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'gtk_color_chooser_add_palette' : parameter 'colors' of type 'nil' not supported

var colorChooserGetRgbaFunction *gi.Function
var colorChooserGetRgbaFunction_Once sync.Once

func colorChooserGetRgbaFunction_Set() error {
	var err error
	colorChooserGetRgbaFunction_Once.Do(func() {
		err = colorChooserInterface_Set()
		if err != nil {
			return
		}
		colorChooserGetRgbaFunction, err = colorChooserInterface.InvokerNew("get_rgba")
	})
	return err
}

// GetRgba is a representation of the C type gtk_color_chooser_get_rgba.
func (recv *ColorChooser) GetRgba() *gdk.RGBA {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := colorChooserGetRgbaFunction_Set()
	if err == nil {
		colorChooserGetRgbaFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := gdk.RGBANewFromNative(outArgs[0].Pointer())

	return out0
}

var colorChooserGetUseAlphaFunction *gi.Function
var colorChooserGetUseAlphaFunction_Once sync.Once

func colorChooserGetUseAlphaFunction_Set() error {
	var err error
	colorChooserGetUseAlphaFunction_Once.Do(func() {
		err = colorChooserInterface_Set()
		if err != nil {
			return
		}
		colorChooserGetUseAlphaFunction, err = colorChooserInterface.InvokerNew("get_use_alpha")
	})
	return err
}

// GetUseAlpha is a representation of the C type gtk_color_chooser_get_use_alpha.
func (recv *ColorChooser) GetUseAlpha() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := colorChooserGetUseAlphaFunction_Set()
	if err == nil {
		ret = colorChooserGetUseAlphaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var colorChooserSetRgbaFunction *gi.Function
var colorChooserSetRgbaFunction_Once sync.Once

func colorChooserSetRgbaFunction_Set() error {
	var err error
	colorChooserSetRgbaFunction_Once.Do(func() {
		err = colorChooserInterface_Set()
		if err != nil {
			return
		}
		colorChooserSetRgbaFunction, err = colorChooserInterface.InvokerNew("set_rgba")
	})
	return err
}

// SetRgba is a representation of the C type gtk_color_chooser_set_rgba.
func (recv *ColorChooser) SetRgba(color *gdk.RGBA) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(color.Native())

	err := colorChooserSetRgbaFunction_Set()
	if err == nil {
		colorChooserSetRgbaFunction.Invoke(inArgs[:], nil)
	}

	return
}

var colorChooserSetUseAlphaFunction *gi.Function
var colorChooserSetUseAlphaFunction_Once sync.Once

func colorChooserSetUseAlphaFunction_Set() error {
	var err error
	colorChooserSetUseAlphaFunction_Once.Do(func() {
		err = colorChooserInterface_Set()
		if err != nil {
			return
		}
		colorChooserSetUseAlphaFunction, err = colorChooserInterface.InvokerNew("set_use_alpha")
	})
	return err
}

// SetUseAlpha is a representation of the C type gtk_color_chooser_set_use_alpha.
func (recv *ColorChooser) SetUseAlpha(useAlpha bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(useAlpha)

	err := colorChooserSetUseAlphaFunction_Set()
	if err == nil {
		colorChooserSetUseAlphaFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectColorActivated connects a callback to the 'color-activated' signal of the ColorChooser.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *ColorChooser) ConnectColorActivated(handler func(instance *ColorChooser, color *gdk.RGBA)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := ColorChooserNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := gdk.RGBANewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "color-activated", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *ColorChooser) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var editableInterface *gi.Interface
var editableInterface_Once sync.Once

func editableInterface_Set() error {
	var err error
	editableInterface_Once.Do(func() {
		editableInterface, err = gi.InterfaceNew("Gtk", "Editable")
	})
	return err
}

type Editable struct {
	native unsafe.Pointer
}

func EditableNewFromNative(native unsafe.Pointer) *Editable {
	instance := &Editable{native: native}

	return instance
}

/*
CastToEditable down casts any arbitrary Object to Editable.
Exercise care, as this is a potentially dangerous function
if the Object is not a Editable.
*/
func CastToEditable(object *gobject.Object) *Editable {
	return EditableNewFromNative(object.Native())
}

// Equals compares this Editable with another Editable, and returns true if they represent the same Object.
func (recv *Editable) Equals(other *Editable) bool {
	return other.Native() == recv.Native()
}

func (recv *Editable) Native() unsafe.Pointer {
	return recv.native
}

var editableCopyClipboardFunction *gi.Function
var editableCopyClipboardFunction_Once sync.Once

func editableCopyClipboardFunction_Set() error {
	var err error
	editableCopyClipboardFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableCopyClipboardFunction, err = editableInterface.InvokerNew("copy_clipboard")
	})
	return err
}

// CopyClipboard is a representation of the C type gtk_editable_copy_clipboard.
func (recv *Editable) CopyClipboard() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := editableCopyClipboardFunction_Set()
	if err == nil {
		editableCopyClipboardFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableCutClipboardFunction *gi.Function
var editableCutClipboardFunction_Once sync.Once

func editableCutClipboardFunction_Set() error {
	var err error
	editableCutClipboardFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableCutClipboardFunction, err = editableInterface.InvokerNew("cut_clipboard")
	})
	return err
}

// CutClipboard is a representation of the C type gtk_editable_cut_clipboard.
func (recv *Editable) CutClipboard() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := editableCutClipboardFunction_Set()
	if err == nil {
		editableCutClipboardFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableDeleteSelectionFunction *gi.Function
var editableDeleteSelectionFunction_Once sync.Once

func editableDeleteSelectionFunction_Set() error {
	var err error
	editableDeleteSelectionFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableDeleteSelectionFunction, err = editableInterface.InvokerNew("delete_selection")
	})
	return err
}

// DeleteSelection is a representation of the C type gtk_editable_delete_selection.
func (recv *Editable) DeleteSelection() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := editableDeleteSelectionFunction_Set()
	if err == nil {
		editableDeleteSelectionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableDeleteTextFunction *gi.Function
var editableDeleteTextFunction_Once sync.Once

func editableDeleteTextFunction_Set() error {
	var err error
	editableDeleteTextFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableDeleteTextFunction, err = editableInterface.InvokerNew("delete_text")
	})
	return err
}

// DeleteText is a representation of the C type gtk_editable_delete_text.
func (recv *Editable) DeleteText(startPos int32, endPos int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startPos)
	inArgs[2].SetInt32(endPos)

	err := editableDeleteTextFunction_Set()
	if err == nil {
		editableDeleteTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableGetCharsFunction *gi.Function
var editableGetCharsFunction_Once sync.Once

func editableGetCharsFunction_Set() error {
	var err error
	editableGetCharsFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableGetCharsFunction, err = editableInterface.InvokerNew("get_chars")
	})
	return err
}

// GetChars is a representation of the C type gtk_editable_get_chars.
func (recv *Editable) GetChars(startPos int32, endPos int32) string {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startPos)
	inArgs[2].SetInt32(endPos)

	var ret gi.Argument

	err := editableGetCharsFunction_Set()
	if err == nil {
		ret = editableGetCharsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var editableGetEditableFunction *gi.Function
var editableGetEditableFunction_Once sync.Once

func editableGetEditableFunction_Set() error {
	var err error
	editableGetEditableFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableGetEditableFunction, err = editableInterface.InvokerNew("get_editable")
	})
	return err
}

// GetEditable is a representation of the C type gtk_editable_get_editable.
func (recv *Editable) GetEditable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := editableGetEditableFunction_Set()
	if err == nil {
		ret = editableGetEditableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var editableGetPositionFunction *gi.Function
var editableGetPositionFunction_Once sync.Once

func editableGetPositionFunction_Set() error {
	var err error
	editableGetPositionFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableGetPositionFunction, err = editableInterface.InvokerNew("get_position")
	})
	return err
}

// GetPosition is a representation of the C type gtk_editable_get_position.
func (recv *Editable) GetPosition() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := editableGetPositionFunction_Set()
	if err == nil {
		ret = editableGetPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var editableGetSelectionBoundsFunction *gi.Function
var editableGetSelectionBoundsFunction_Once sync.Once

func editableGetSelectionBoundsFunction_Set() error {
	var err error
	editableGetSelectionBoundsFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableGetSelectionBoundsFunction, err = editableInterface.InvokerNew("get_selection_bounds")
	})
	return err
}

// GetSelectionBounds is a representation of the C type gtk_editable_get_selection_bounds.
func (recv *Editable) GetSelectionBounds() (bool, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := editableGetSelectionBoundsFunction_Set()
	if err == nil {
		ret = editableGetSelectionBoundsFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var editableInsertTextFunction *gi.Function
var editableInsertTextFunction_Once sync.Once

func editableInsertTextFunction_Set() error {
	var err error
	editableInsertTextFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableInsertTextFunction, err = editableInterface.InvokerNew("insert_text")
	})
	return err
}

// InsertText is a representation of the C type gtk_editable_insert_text.
func (recv *Editable) InsertText(newText string, newTextLength int32, position int32) int32 {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(newText)
	inArgs[2].SetInt32(newTextLength)
	inArgs[3].SetInt32(position)

	var outArgs [1]gi.Argument

	err := editableInsertTextFunction_Set()
	if err == nil {
		editableInsertTextFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var editablePasteClipboardFunction *gi.Function
var editablePasteClipboardFunction_Once sync.Once

func editablePasteClipboardFunction_Set() error {
	var err error
	editablePasteClipboardFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editablePasteClipboardFunction, err = editableInterface.InvokerNew("paste_clipboard")
	})
	return err
}

// PasteClipboard is a representation of the C type gtk_editable_paste_clipboard.
func (recv *Editable) PasteClipboard() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := editablePasteClipboardFunction_Set()
	if err == nil {
		editablePasteClipboardFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableSelectRegionFunction *gi.Function
var editableSelectRegionFunction_Once sync.Once

func editableSelectRegionFunction_Set() error {
	var err error
	editableSelectRegionFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableSelectRegionFunction, err = editableInterface.InvokerNew("select_region")
	})
	return err
}

// SelectRegion is a representation of the C type gtk_editable_select_region.
func (recv *Editable) SelectRegion(startPos int32, endPos int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(startPos)
	inArgs[2].SetInt32(endPos)

	err := editableSelectRegionFunction_Set()
	if err == nil {
		editableSelectRegionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableSetEditableFunction *gi.Function
var editableSetEditableFunction_Once sync.Once

func editableSetEditableFunction_Set() error {
	var err error
	editableSetEditableFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableSetEditableFunction, err = editableInterface.InvokerNew("set_editable")
	})
	return err
}

// SetEditable is a representation of the C type gtk_editable_set_editable.
func (recv *Editable) SetEditable(isEditable bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(isEditable)

	err := editableSetEditableFunction_Set()
	if err == nil {
		editableSetEditableFunction.Invoke(inArgs[:], nil)
	}

	return
}

var editableSetPositionFunction *gi.Function
var editableSetPositionFunction_Once sync.Once

func editableSetPositionFunction_Set() error {
	var err error
	editableSetPositionFunction_Once.Do(func() {
		err = editableInterface_Set()
		if err != nil {
			return
		}
		editableSetPositionFunction, err = editableInterface.InvokerNew("set_position")
	})
	return err
}

// SetPosition is a representation of the C type gtk_editable_set_position.
func (recv *Editable) SetPosition(position int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(position)

	err := editableSetPositionFunction_Set()
	if err == nil {
		editableSetPositionFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectChanged connects a callback to the 'changed' signal of the Editable.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Editable) ConnectChanged(handler func(instance *Editable)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := EditableNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "changed", marshal)
}

/*
ConnectDeleteText connects a callback to the 'delete-text' signal of the Editable.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Editable) ConnectDeleteText(handler func(instance *Editable, startPos int32, endPos int32)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := EditableNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetInt()

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := object2.GetInt()

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "delete-text", marshal)
}

/*
ConnectInsertText connects a callback to the 'insert-text' signal of the Editable.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *Editable) ConnectInsertText(handler func(instance *Editable, newText string, newTextLength int32, position int32) int32) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		// Has out params
	}

	return callback.ConnectSignal(recv.Native(), "insert-text", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *Editable) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var fileChooserInterface *gi.Interface
var fileChooserInterface_Once sync.Once

func fileChooserInterface_Set() error {
	var err error
	fileChooserInterface_Once.Do(func() {
		fileChooserInterface, err = gi.InterfaceNew("Gtk", "FileChooser")
	})
	return err
}

type FileChooser struct {
	native unsafe.Pointer
}

func FileChooserNewFromNative(native unsafe.Pointer) *FileChooser {
	instance := &FileChooser{native: native}

	return instance
}

/*
CastToFileChooser down casts any arbitrary Object to FileChooser.
Exercise care, as this is a potentially dangerous function
if the Object is not a FileChooser.
*/
func CastToFileChooser(object *gobject.Object) *FileChooser {
	return FileChooserNewFromNative(object.Native())
}

// Equals compares this FileChooser with another FileChooser, and returns true if they represent the same Object.
func (recv *FileChooser) Equals(other *FileChooser) bool {
	return other.Native() == recv.Native()
}

func (recv *FileChooser) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'gtk_file_chooser_add_choice' : parameter 'options' of type 'nil' not supported

var fileChooserAddFilterFunction *gi.Function
var fileChooserAddFilterFunction_Once sync.Once

func fileChooserAddFilterFunction_Set() error {
	var err error
	fileChooserAddFilterFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserAddFilterFunction, err = fileChooserInterface.InvokerNew("add_filter")
	})
	return err
}

// AddFilter is a representation of the C type gtk_file_chooser_add_filter.
func (recv *FileChooser) AddFilter(filter *FileFilter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(filter.Native())

	err := fileChooserAddFilterFunction_Set()
	if err == nil {
		fileChooserAddFilterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserAddShortcutFolderFunction *gi.Function
var fileChooserAddShortcutFolderFunction_Once sync.Once

func fileChooserAddShortcutFolderFunction_Set() error {
	var err error
	fileChooserAddShortcutFolderFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserAddShortcutFolderFunction, err = fileChooserInterface.InvokerNew("add_shortcut_folder")
	})
	return err
}

// AddShortcutFolder is a representation of the C type gtk_file_chooser_add_shortcut_folder.
func (recv *FileChooser) AddShortcutFolder(folder string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(folder)

	var ret gi.Argument

	err := fileChooserAddShortcutFolderFunction_Set()
	if err == nil {
		ret = fileChooserAddShortcutFolderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserAddShortcutFolderUriFunction *gi.Function
var fileChooserAddShortcutFolderUriFunction_Once sync.Once

func fileChooserAddShortcutFolderUriFunction_Set() error {
	var err error
	fileChooserAddShortcutFolderUriFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserAddShortcutFolderUriFunction, err = fileChooserInterface.InvokerNew("add_shortcut_folder_uri")
	})
	return err
}

// AddShortcutFolderUri is a representation of the C type gtk_file_chooser_add_shortcut_folder_uri.
func (recv *FileChooser) AddShortcutFolderUri(uri string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := fileChooserAddShortcutFolderUriFunction_Set()
	if err == nil {
		ret = fileChooserAddShortcutFolderUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserGetActionFunction *gi.Function
var fileChooserGetActionFunction_Once sync.Once

func fileChooserGetActionFunction_Set() error {
	var err error
	fileChooserGetActionFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetActionFunction, err = fileChooserInterface.InvokerNew("get_action")
	})
	return err
}

// GetAction is a representation of the C type gtk_file_chooser_get_action.
func (recv *FileChooser) GetAction() FileChooserAction {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetActionFunction_Set()
	if err == nil {
		ret = fileChooserGetActionFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileChooserAction(ret.Int32())

	return retGo
}

var fileChooserGetChoiceFunction *gi.Function
var fileChooserGetChoiceFunction_Once sync.Once

func fileChooserGetChoiceFunction_Set() error {
	var err error
	fileChooserGetChoiceFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetChoiceFunction, err = fileChooserInterface.InvokerNew("get_choice")
	})
	return err
}

// GetChoice is a representation of the C type gtk_file_chooser_get_choice.
func (recv *FileChooser) GetChoice(id string) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(id)

	var ret gi.Argument

	err := fileChooserGetChoiceFunction_Set()
	if err == nil {
		ret = fileChooserGetChoiceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fileChooserGetCreateFoldersFunction *gi.Function
var fileChooserGetCreateFoldersFunction_Once sync.Once

func fileChooserGetCreateFoldersFunction_Set() error {
	var err error
	fileChooserGetCreateFoldersFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetCreateFoldersFunction, err = fileChooserInterface.InvokerNew("get_create_folders")
	})
	return err
}

// GetCreateFolders is a representation of the C type gtk_file_chooser_get_create_folders.
func (recv *FileChooser) GetCreateFolders() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetCreateFoldersFunction_Set()
	if err == nil {
		ret = fileChooserGetCreateFoldersFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserGetCurrentFolderFunction *gi.Function
var fileChooserGetCurrentFolderFunction_Once sync.Once

func fileChooserGetCurrentFolderFunction_Set() error {
	var err error
	fileChooserGetCurrentFolderFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetCurrentFolderFunction, err = fileChooserInterface.InvokerNew("get_current_folder")
	})
	return err
}

// GetCurrentFolder is a representation of the C type gtk_file_chooser_get_current_folder.
func (recv *FileChooser) GetCurrentFolder() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetCurrentFolderFunction_Set()
	if err == nil {
		ret = fileChooserGetCurrentFolderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_file_chooser_get_current_folder_file' : return type 'Gio.File' not supported

var fileChooserGetCurrentFolderUriFunction *gi.Function
var fileChooserGetCurrentFolderUriFunction_Once sync.Once

func fileChooserGetCurrentFolderUriFunction_Set() error {
	var err error
	fileChooserGetCurrentFolderUriFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetCurrentFolderUriFunction, err = fileChooserInterface.InvokerNew("get_current_folder_uri")
	})
	return err
}

// GetCurrentFolderUri is a representation of the C type gtk_file_chooser_get_current_folder_uri.
func (recv *FileChooser) GetCurrentFolderUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetCurrentFolderUriFunction_Set()
	if err == nil {
		ret = fileChooserGetCurrentFolderUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileChooserGetCurrentNameFunction *gi.Function
var fileChooserGetCurrentNameFunction_Once sync.Once

func fileChooserGetCurrentNameFunction_Set() error {
	var err error
	fileChooserGetCurrentNameFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetCurrentNameFunction, err = fileChooserInterface.InvokerNew("get_current_name")
	})
	return err
}

// GetCurrentName is a representation of the C type gtk_file_chooser_get_current_name.
func (recv *FileChooser) GetCurrentName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetCurrentNameFunction_Set()
	if err == nil {
		ret = fileChooserGetCurrentNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileChooserGetDoOverwriteConfirmationFunction *gi.Function
var fileChooserGetDoOverwriteConfirmationFunction_Once sync.Once

func fileChooserGetDoOverwriteConfirmationFunction_Set() error {
	var err error
	fileChooserGetDoOverwriteConfirmationFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetDoOverwriteConfirmationFunction, err = fileChooserInterface.InvokerNew("get_do_overwrite_confirmation")
	})
	return err
}

// GetDoOverwriteConfirmation is a representation of the C type gtk_file_chooser_get_do_overwrite_confirmation.
func (recv *FileChooser) GetDoOverwriteConfirmation() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetDoOverwriteConfirmationFunction_Set()
	if err == nil {
		ret = fileChooserGetDoOverwriteConfirmationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserGetExtraWidgetFunction *gi.Function
var fileChooserGetExtraWidgetFunction_Once sync.Once

func fileChooserGetExtraWidgetFunction_Set() error {
	var err error
	fileChooserGetExtraWidgetFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetExtraWidgetFunction, err = fileChooserInterface.InvokerNew("get_extra_widget")
	})
	return err
}

// GetExtraWidget is a representation of the C type gtk_file_chooser_get_extra_widget.
func (recv *FileChooser) GetExtraWidget() *Widget {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetExtraWidgetFunction_Set()
	if err == nil {
		ret = fileChooserGetExtraWidgetFunction.Invoke(inArgs[:], nil)
	}

	retGo := WidgetNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'gtk_file_chooser_get_file' : return type 'Gio.File' not supported

var fileChooserGetFilenameFunction *gi.Function
var fileChooserGetFilenameFunction_Once sync.Once

func fileChooserGetFilenameFunction_Set() error {
	var err error
	fileChooserGetFilenameFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetFilenameFunction, err = fileChooserInterface.InvokerNew("get_filename")
	})
	return err
}

// GetFilename is a representation of the C type gtk_file_chooser_get_filename.
func (recv *FileChooser) GetFilename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetFilenameFunction_Set()
	if err == nil {
		ret = fileChooserGetFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileChooserGetFilenamesFunction *gi.Function
var fileChooserGetFilenamesFunction_Once sync.Once

func fileChooserGetFilenamesFunction_Set() error {
	var err error
	fileChooserGetFilenamesFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetFilenamesFunction, err = fileChooserInterface.InvokerNew("get_filenames")
	})
	return err
}

// GetFilenames is a representation of the C type gtk_file_chooser_get_filenames.
func (recv *FileChooser) GetFilenames() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetFilenamesFunction_Set()
	if err == nil {
		ret = fileChooserGetFilenamesFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var fileChooserGetFilesFunction *gi.Function
var fileChooserGetFilesFunction_Once sync.Once

func fileChooserGetFilesFunction_Set() error {
	var err error
	fileChooserGetFilesFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetFilesFunction, err = fileChooserInterface.InvokerNew("get_files")
	})
	return err
}

// GetFiles is a representation of the C type gtk_file_chooser_get_files.
func (recv *FileChooser) GetFiles() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetFilesFunction_Set()
	if err == nil {
		ret = fileChooserGetFilesFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var fileChooserGetFilterFunction *gi.Function
var fileChooserGetFilterFunction_Once sync.Once

func fileChooserGetFilterFunction_Set() error {
	var err error
	fileChooserGetFilterFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetFilterFunction, err = fileChooserInterface.InvokerNew("get_filter")
	})
	return err
}

// GetFilter is a representation of the C type gtk_file_chooser_get_filter.
func (recv *FileChooser) GetFilter() *FileFilter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetFilterFunction_Set()
	if err == nil {
		ret = fileChooserGetFilterFunction.Invoke(inArgs[:], nil)
	}

	retGo := FileFilterNewFromNative(ret.Pointer())

	return retGo
}

var fileChooserGetLocalOnlyFunction *gi.Function
var fileChooserGetLocalOnlyFunction_Once sync.Once

func fileChooserGetLocalOnlyFunction_Set() error {
	var err error
	fileChooserGetLocalOnlyFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetLocalOnlyFunction, err = fileChooserInterface.InvokerNew("get_local_only")
	})
	return err
}

// GetLocalOnly is a representation of the C type gtk_file_chooser_get_local_only.
func (recv *FileChooser) GetLocalOnly() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetLocalOnlyFunction_Set()
	if err == nil {
		ret = fileChooserGetLocalOnlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_file_chooser_get_preview_file' : return type 'Gio.File' not supported

var fileChooserGetPreviewFilenameFunction *gi.Function
var fileChooserGetPreviewFilenameFunction_Once sync.Once

func fileChooserGetPreviewFilenameFunction_Set() error {
	var err error
	fileChooserGetPreviewFilenameFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetPreviewFilenameFunction, err = fileChooserInterface.InvokerNew("get_preview_filename")
	})
	return err
}

// GetPreviewFilename is a representation of the C type gtk_file_chooser_get_preview_filename.
func (recv *FileChooser) GetPreviewFilename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetPreviewFilenameFunction_Set()
	if err == nil {
		ret = fileChooserGetPreviewFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileChooserGetPreviewUriFunction *gi.Function
var fileChooserGetPreviewUriFunction_Once sync.Once

func fileChooserGetPreviewUriFunction_Set() error {
	var err error
	fileChooserGetPreviewUriFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetPreviewUriFunction, err = fileChooserInterface.InvokerNew("get_preview_uri")
	})
	return err
}

// GetPreviewUri is a representation of the C type gtk_file_chooser_get_preview_uri.
func (recv *FileChooser) GetPreviewUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetPreviewUriFunction_Set()
	if err == nil {
		ret = fileChooserGetPreviewUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileChooserGetPreviewWidgetFunction *gi.Function
var fileChooserGetPreviewWidgetFunction_Once sync.Once

func fileChooserGetPreviewWidgetFunction_Set() error {
	var err error
	fileChooserGetPreviewWidgetFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetPreviewWidgetFunction, err = fileChooserInterface.InvokerNew("get_preview_widget")
	})
	return err
}

// GetPreviewWidget is a representation of the C type gtk_file_chooser_get_preview_widget.
func (recv *FileChooser) GetPreviewWidget() *Widget {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetPreviewWidgetFunction_Set()
	if err == nil {
		ret = fileChooserGetPreviewWidgetFunction.Invoke(inArgs[:], nil)
	}

	retGo := WidgetNewFromNative(ret.Pointer())

	return retGo
}

var fileChooserGetPreviewWidgetActiveFunction *gi.Function
var fileChooserGetPreviewWidgetActiveFunction_Once sync.Once

func fileChooserGetPreviewWidgetActiveFunction_Set() error {
	var err error
	fileChooserGetPreviewWidgetActiveFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetPreviewWidgetActiveFunction, err = fileChooserInterface.InvokerNew("get_preview_widget_active")
	})
	return err
}

// GetPreviewWidgetActive is a representation of the C type gtk_file_chooser_get_preview_widget_active.
func (recv *FileChooser) GetPreviewWidgetActive() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetPreviewWidgetActiveFunction_Set()
	if err == nil {
		ret = fileChooserGetPreviewWidgetActiveFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserGetSelectMultipleFunction *gi.Function
var fileChooserGetSelectMultipleFunction_Once sync.Once

func fileChooserGetSelectMultipleFunction_Set() error {
	var err error
	fileChooserGetSelectMultipleFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetSelectMultipleFunction, err = fileChooserInterface.InvokerNew("get_select_multiple")
	})
	return err
}

// GetSelectMultiple is a representation of the C type gtk_file_chooser_get_select_multiple.
func (recv *FileChooser) GetSelectMultiple() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetSelectMultipleFunction_Set()
	if err == nil {
		ret = fileChooserGetSelectMultipleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserGetShowHiddenFunction *gi.Function
var fileChooserGetShowHiddenFunction_Once sync.Once

func fileChooserGetShowHiddenFunction_Set() error {
	var err error
	fileChooserGetShowHiddenFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetShowHiddenFunction, err = fileChooserInterface.InvokerNew("get_show_hidden")
	})
	return err
}

// GetShowHidden is a representation of the C type gtk_file_chooser_get_show_hidden.
func (recv *FileChooser) GetShowHidden() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetShowHiddenFunction_Set()
	if err == nil {
		ret = fileChooserGetShowHiddenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserGetUriFunction *gi.Function
var fileChooserGetUriFunction_Once sync.Once

func fileChooserGetUriFunction_Set() error {
	var err error
	fileChooserGetUriFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetUriFunction, err = fileChooserInterface.InvokerNew("get_uri")
	})
	return err
}

// GetUri is a representation of the C type gtk_file_chooser_get_uri.
func (recv *FileChooser) GetUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetUriFunction_Set()
	if err == nil {
		ret = fileChooserGetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fileChooserGetUrisFunction *gi.Function
var fileChooserGetUrisFunction_Once sync.Once

func fileChooserGetUrisFunction_Set() error {
	var err error
	fileChooserGetUrisFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetUrisFunction, err = fileChooserInterface.InvokerNew("get_uris")
	})
	return err
}

// GetUris is a representation of the C type gtk_file_chooser_get_uris.
func (recv *FileChooser) GetUris() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetUrisFunction_Set()
	if err == nil {
		ret = fileChooserGetUrisFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var fileChooserGetUsePreviewLabelFunction *gi.Function
var fileChooserGetUsePreviewLabelFunction_Once sync.Once

func fileChooserGetUsePreviewLabelFunction_Set() error {
	var err error
	fileChooserGetUsePreviewLabelFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserGetUsePreviewLabelFunction, err = fileChooserInterface.InvokerNew("get_use_preview_label")
	})
	return err
}

// GetUsePreviewLabel is a representation of the C type gtk_file_chooser_get_use_preview_label.
func (recv *FileChooser) GetUsePreviewLabel() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserGetUsePreviewLabelFunction_Set()
	if err == nil {
		ret = fileChooserGetUsePreviewLabelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserListFiltersFunction *gi.Function
var fileChooserListFiltersFunction_Once sync.Once

func fileChooserListFiltersFunction_Set() error {
	var err error
	fileChooserListFiltersFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserListFiltersFunction, err = fileChooserInterface.InvokerNew("list_filters")
	})
	return err
}

// ListFilters is a representation of the C type gtk_file_chooser_list_filters.
func (recv *FileChooser) ListFilters() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserListFiltersFunction_Set()
	if err == nil {
		ret = fileChooserListFiltersFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var fileChooserListShortcutFolderUrisFunction *gi.Function
var fileChooserListShortcutFolderUrisFunction_Once sync.Once

func fileChooserListShortcutFolderUrisFunction_Set() error {
	var err error
	fileChooserListShortcutFolderUrisFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserListShortcutFolderUrisFunction, err = fileChooserInterface.InvokerNew("list_shortcut_folder_uris")
	})
	return err
}

// ListShortcutFolderUris is a representation of the C type gtk_file_chooser_list_shortcut_folder_uris.
func (recv *FileChooser) ListShortcutFolderUris() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserListShortcutFolderUrisFunction_Set()
	if err == nil {
		ret = fileChooserListShortcutFolderUrisFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var fileChooserListShortcutFoldersFunction *gi.Function
var fileChooserListShortcutFoldersFunction_Once sync.Once

func fileChooserListShortcutFoldersFunction_Set() error {
	var err error
	fileChooserListShortcutFoldersFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserListShortcutFoldersFunction, err = fileChooserInterface.InvokerNew("list_shortcut_folders")
	})
	return err
}

// ListShortcutFolders is a representation of the C type gtk_file_chooser_list_shortcut_folders.
func (recv *FileChooser) ListShortcutFolders() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fileChooserListShortcutFoldersFunction_Set()
	if err == nil {
		ret = fileChooserListShortcutFoldersFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var fileChooserRemoveChoiceFunction *gi.Function
var fileChooserRemoveChoiceFunction_Once sync.Once

func fileChooserRemoveChoiceFunction_Set() error {
	var err error
	fileChooserRemoveChoiceFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserRemoveChoiceFunction, err = fileChooserInterface.InvokerNew("remove_choice")
	})
	return err
}

// RemoveChoice is a representation of the C type gtk_file_chooser_remove_choice.
func (recv *FileChooser) RemoveChoice(id string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(id)

	err := fileChooserRemoveChoiceFunction_Set()
	if err == nil {
		fileChooserRemoveChoiceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserRemoveFilterFunction *gi.Function
var fileChooserRemoveFilterFunction_Once sync.Once

func fileChooserRemoveFilterFunction_Set() error {
	var err error
	fileChooserRemoveFilterFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserRemoveFilterFunction, err = fileChooserInterface.InvokerNew("remove_filter")
	})
	return err
}

// RemoveFilter is a representation of the C type gtk_file_chooser_remove_filter.
func (recv *FileChooser) RemoveFilter(filter *FileFilter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(filter.Native())

	err := fileChooserRemoveFilterFunction_Set()
	if err == nil {
		fileChooserRemoveFilterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserRemoveShortcutFolderFunction *gi.Function
var fileChooserRemoveShortcutFolderFunction_Once sync.Once

func fileChooserRemoveShortcutFolderFunction_Set() error {
	var err error
	fileChooserRemoveShortcutFolderFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserRemoveShortcutFolderFunction, err = fileChooserInterface.InvokerNew("remove_shortcut_folder")
	})
	return err
}

// RemoveShortcutFolder is a representation of the C type gtk_file_chooser_remove_shortcut_folder.
func (recv *FileChooser) RemoveShortcutFolder(folder string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(folder)

	var ret gi.Argument

	err := fileChooserRemoveShortcutFolderFunction_Set()
	if err == nil {
		ret = fileChooserRemoveShortcutFolderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserRemoveShortcutFolderUriFunction *gi.Function
var fileChooserRemoveShortcutFolderUriFunction_Once sync.Once

func fileChooserRemoveShortcutFolderUriFunction_Set() error {
	var err error
	fileChooserRemoveShortcutFolderUriFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserRemoveShortcutFolderUriFunction, err = fileChooserInterface.InvokerNew("remove_shortcut_folder_uri")
	})
	return err
}

// RemoveShortcutFolderUri is a representation of the C type gtk_file_chooser_remove_shortcut_folder_uri.
func (recv *FileChooser) RemoveShortcutFolderUri(uri string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := fileChooserRemoveShortcutFolderUriFunction_Set()
	if err == nil {
		ret = fileChooserRemoveShortcutFolderUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserSelectAllFunction *gi.Function
var fileChooserSelectAllFunction_Once sync.Once

func fileChooserSelectAllFunction_Set() error {
	var err error
	fileChooserSelectAllFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSelectAllFunction, err = fileChooserInterface.InvokerNew("select_all")
	})
	return err
}

// SelectAll is a representation of the C type gtk_file_chooser_select_all.
func (recv *FileChooser) SelectAll() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := fileChooserSelectAllFunction_Set()
	if err == nil {
		fileChooserSelectAllFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_file_chooser_select_file' : parameter 'file' of type 'Gio.File' not supported

var fileChooserSelectFilenameFunction *gi.Function
var fileChooserSelectFilenameFunction_Once sync.Once

func fileChooserSelectFilenameFunction_Set() error {
	var err error
	fileChooserSelectFilenameFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSelectFilenameFunction, err = fileChooserInterface.InvokerNew("select_filename")
	})
	return err
}

// SelectFilename is a representation of the C type gtk_file_chooser_select_filename.
func (recv *FileChooser) SelectFilename(filename string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(filename)

	var ret gi.Argument

	err := fileChooserSelectFilenameFunction_Set()
	if err == nil {
		ret = fileChooserSelectFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserSelectUriFunction *gi.Function
var fileChooserSelectUriFunction_Once sync.Once

func fileChooserSelectUriFunction_Set() error {
	var err error
	fileChooserSelectUriFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSelectUriFunction, err = fileChooserInterface.InvokerNew("select_uri")
	})
	return err
}

// SelectUri is a representation of the C type gtk_file_chooser_select_uri.
func (recv *FileChooser) SelectUri(uri string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := fileChooserSelectUriFunction_Set()
	if err == nil {
		ret = fileChooserSelectUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserSetActionFunction *gi.Function
var fileChooserSetActionFunction_Once sync.Once

func fileChooserSetActionFunction_Set() error {
	var err error
	fileChooserSetActionFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetActionFunction, err = fileChooserInterface.InvokerNew("set_action")
	})
	return err
}

// SetAction is a representation of the C type gtk_file_chooser_set_action.
func (recv *FileChooser) SetAction(action FileChooserAction) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(action))

	err := fileChooserSetActionFunction_Set()
	if err == nil {
		fileChooserSetActionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetChoiceFunction *gi.Function
var fileChooserSetChoiceFunction_Once sync.Once

func fileChooserSetChoiceFunction_Set() error {
	var err error
	fileChooserSetChoiceFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetChoiceFunction, err = fileChooserInterface.InvokerNew("set_choice")
	})
	return err
}

// SetChoice is a representation of the C type gtk_file_chooser_set_choice.
func (recv *FileChooser) SetChoice(id string, option string) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(id)
	inArgs[2].SetString(option)

	err := fileChooserSetChoiceFunction_Set()
	if err == nil {
		fileChooserSetChoiceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetCreateFoldersFunction *gi.Function
var fileChooserSetCreateFoldersFunction_Once sync.Once

func fileChooserSetCreateFoldersFunction_Set() error {
	var err error
	fileChooserSetCreateFoldersFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetCreateFoldersFunction, err = fileChooserInterface.InvokerNew("set_create_folders")
	})
	return err
}

// SetCreateFolders is a representation of the C type gtk_file_chooser_set_create_folders.
func (recv *FileChooser) SetCreateFolders(createFolders bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(createFolders)

	err := fileChooserSetCreateFoldersFunction_Set()
	if err == nil {
		fileChooserSetCreateFoldersFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetCurrentFolderFunction *gi.Function
var fileChooserSetCurrentFolderFunction_Once sync.Once

func fileChooserSetCurrentFolderFunction_Set() error {
	var err error
	fileChooserSetCurrentFolderFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetCurrentFolderFunction, err = fileChooserInterface.InvokerNew("set_current_folder")
	})
	return err
}

// SetCurrentFolder is a representation of the C type gtk_file_chooser_set_current_folder.
func (recv *FileChooser) SetCurrentFolder(filename string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(filename)

	var ret gi.Argument

	err := fileChooserSetCurrentFolderFunction_Set()
	if err == nil {
		ret = fileChooserSetCurrentFolderFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_file_chooser_set_current_folder_file' : parameter 'file' of type 'Gio.File' not supported

var fileChooserSetCurrentFolderUriFunction *gi.Function
var fileChooserSetCurrentFolderUriFunction_Once sync.Once

func fileChooserSetCurrentFolderUriFunction_Set() error {
	var err error
	fileChooserSetCurrentFolderUriFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetCurrentFolderUriFunction, err = fileChooserInterface.InvokerNew("set_current_folder_uri")
	})
	return err
}

// SetCurrentFolderUri is a representation of the C type gtk_file_chooser_set_current_folder_uri.
func (recv *FileChooser) SetCurrentFolderUri(uri string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := fileChooserSetCurrentFolderUriFunction_Set()
	if err == nil {
		ret = fileChooserSetCurrentFolderUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserSetCurrentNameFunction *gi.Function
var fileChooserSetCurrentNameFunction_Once sync.Once

func fileChooserSetCurrentNameFunction_Set() error {
	var err error
	fileChooserSetCurrentNameFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetCurrentNameFunction, err = fileChooserInterface.InvokerNew("set_current_name")
	})
	return err
}

// SetCurrentName is a representation of the C type gtk_file_chooser_set_current_name.
func (recv *FileChooser) SetCurrentName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(name)

	err := fileChooserSetCurrentNameFunction_Set()
	if err == nil {
		fileChooserSetCurrentNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetDoOverwriteConfirmationFunction *gi.Function
var fileChooserSetDoOverwriteConfirmationFunction_Once sync.Once

func fileChooserSetDoOverwriteConfirmationFunction_Set() error {
	var err error
	fileChooserSetDoOverwriteConfirmationFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetDoOverwriteConfirmationFunction, err = fileChooserInterface.InvokerNew("set_do_overwrite_confirmation")
	})
	return err
}

// SetDoOverwriteConfirmation is a representation of the C type gtk_file_chooser_set_do_overwrite_confirmation.
func (recv *FileChooser) SetDoOverwriteConfirmation(doOverwriteConfirmation bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(doOverwriteConfirmation)

	err := fileChooserSetDoOverwriteConfirmationFunction_Set()
	if err == nil {
		fileChooserSetDoOverwriteConfirmationFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetExtraWidgetFunction *gi.Function
var fileChooserSetExtraWidgetFunction_Once sync.Once

func fileChooserSetExtraWidgetFunction_Set() error {
	var err error
	fileChooserSetExtraWidgetFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetExtraWidgetFunction, err = fileChooserInterface.InvokerNew("set_extra_widget")
	})
	return err
}

// SetExtraWidget is a representation of the C type gtk_file_chooser_set_extra_widget.
func (recv *FileChooser) SetExtraWidget(extraWidget *Widget) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(extraWidget.Native())

	err := fileChooserSetExtraWidgetFunction_Set()
	if err == nil {
		fileChooserSetExtraWidgetFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_file_chooser_set_file' : parameter 'file' of type 'Gio.File' not supported

var fileChooserSetFilenameFunction *gi.Function
var fileChooserSetFilenameFunction_Once sync.Once

func fileChooserSetFilenameFunction_Set() error {
	var err error
	fileChooserSetFilenameFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetFilenameFunction, err = fileChooserInterface.InvokerNew("set_filename")
	})
	return err
}

// SetFilename is a representation of the C type gtk_file_chooser_set_filename.
func (recv *FileChooser) SetFilename(filename string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(filename)

	var ret gi.Argument

	err := fileChooserSetFilenameFunction_Set()
	if err == nil {
		ret = fileChooserSetFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserSetFilterFunction *gi.Function
var fileChooserSetFilterFunction_Once sync.Once

func fileChooserSetFilterFunction_Set() error {
	var err error
	fileChooserSetFilterFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetFilterFunction, err = fileChooserInterface.InvokerNew("set_filter")
	})
	return err
}

// SetFilter is a representation of the C type gtk_file_chooser_set_filter.
func (recv *FileChooser) SetFilter(filter *FileFilter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(filter.Native())

	err := fileChooserSetFilterFunction_Set()
	if err == nil {
		fileChooserSetFilterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetLocalOnlyFunction *gi.Function
var fileChooserSetLocalOnlyFunction_Once sync.Once

func fileChooserSetLocalOnlyFunction_Set() error {
	var err error
	fileChooserSetLocalOnlyFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetLocalOnlyFunction, err = fileChooserInterface.InvokerNew("set_local_only")
	})
	return err
}

// SetLocalOnly is a representation of the C type gtk_file_chooser_set_local_only.
func (recv *FileChooser) SetLocalOnly(localOnly bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(localOnly)

	err := fileChooserSetLocalOnlyFunction_Set()
	if err == nil {
		fileChooserSetLocalOnlyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetPreviewWidgetFunction *gi.Function
var fileChooserSetPreviewWidgetFunction_Once sync.Once

func fileChooserSetPreviewWidgetFunction_Set() error {
	var err error
	fileChooserSetPreviewWidgetFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetPreviewWidgetFunction, err = fileChooserInterface.InvokerNew("set_preview_widget")
	})
	return err
}

// SetPreviewWidget is a representation of the C type gtk_file_chooser_set_preview_widget.
func (recv *FileChooser) SetPreviewWidget(previewWidget *Widget) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(previewWidget.Native())

	err := fileChooserSetPreviewWidgetFunction_Set()
	if err == nil {
		fileChooserSetPreviewWidgetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetPreviewWidgetActiveFunction *gi.Function
var fileChooserSetPreviewWidgetActiveFunction_Once sync.Once

func fileChooserSetPreviewWidgetActiveFunction_Set() error {
	var err error
	fileChooserSetPreviewWidgetActiveFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetPreviewWidgetActiveFunction, err = fileChooserInterface.InvokerNew("set_preview_widget_active")
	})
	return err
}

// SetPreviewWidgetActive is a representation of the C type gtk_file_chooser_set_preview_widget_active.
func (recv *FileChooser) SetPreviewWidgetActive(active bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(active)

	err := fileChooserSetPreviewWidgetActiveFunction_Set()
	if err == nil {
		fileChooserSetPreviewWidgetActiveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetSelectMultipleFunction *gi.Function
var fileChooserSetSelectMultipleFunction_Once sync.Once

func fileChooserSetSelectMultipleFunction_Set() error {
	var err error
	fileChooserSetSelectMultipleFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetSelectMultipleFunction, err = fileChooserInterface.InvokerNew("set_select_multiple")
	})
	return err
}

// SetSelectMultiple is a representation of the C type gtk_file_chooser_set_select_multiple.
func (recv *FileChooser) SetSelectMultiple(selectMultiple bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(selectMultiple)

	err := fileChooserSetSelectMultipleFunction_Set()
	if err == nil {
		fileChooserSetSelectMultipleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetShowHiddenFunction *gi.Function
var fileChooserSetShowHiddenFunction_Once sync.Once

func fileChooserSetShowHiddenFunction_Set() error {
	var err error
	fileChooserSetShowHiddenFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetShowHiddenFunction, err = fileChooserInterface.InvokerNew("set_show_hidden")
	})
	return err
}

// SetShowHidden is a representation of the C type gtk_file_chooser_set_show_hidden.
func (recv *FileChooser) SetShowHidden(showHidden bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(showHidden)

	err := fileChooserSetShowHiddenFunction_Set()
	if err == nil {
		fileChooserSetShowHiddenFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserSetUriFunction *gi.Function
var fileChooserSetUriFunction_Once sync.Once

func fileChooserSetUriFunction_Set() error {
	var err error
	fileChooserSetUriFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetUriFunction, err = fileChooserInterface.InvokerNew("set_uri")
	})
	return err
}

// SetUri is a representation of the C type gtk_file_chooser_set_uri.
func (recv *FileChooser) SetUri(uri string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := fileChooserSetUriFunction_Set()
	if err == nil {
		ret = fileChooserSetUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fileChooserSetUsePreviewLabelFunction *gi.Function
var fileChooserSetUsePreviewLabelFunction_Once sync.Once

func fileChooserSetUsePreviewLabelFunction_Set() error {
	var err error
	fileChooserSetUsePreviewLabelFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserSetUsePreviewLabelFunction, err = fileChooserInterface.InvokerNew("set_use_preview_label")
	})
	return err
}

// SetUsePreviewLabel is a representation of the C type gtk_file_chooser_set_use_preview_label.
func (recv *FileChooser) SetUsePreviewLabel(useLabel bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(useLabel)

	err := fileChooserSetUsePreviewLabelFunction_Set()
	if err == nil {
		fileChooserSetUsePreviewLabelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserUnselectAllFunction *gi.Function
var fileChooserUnselectAllFunction_Once sync.Once

func fileChooserUnselectAllFunction_Set() error {
	var err error
	fileChooserUnselectAllFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserUnselectAllFunction, err = fileChooserInterface.InvokerNew("unselect_all")
	})
	return err
}

// UnselectAll is a representation of the C type gtk_file_chooser_unselect_all.
func (recv *FileChooser) UnselectAll() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := fileChooserUnselectAllFunction_Set()
	if err == nil {
		fileChooserUnselectAllFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_file_chooser_unselect_file' : parameter 'file' of type 'Gio.File' not supported

var fileChooserUnselectFilenameFunction *gi.Function
var fileChooserUnselectFilenameFunction_Once sync.Once

func fileChooserUnselectFilenameFunction_Set() error {
	var err error
	fileChooserUnselectFilenameFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserUnselectFilenameFunction, err = fileChooserInterface.InvokerNew("unselect_filename")
	})
	return err
}

// UnselectFilename is a representation of the C type gtk_file_chooser_unselect_filename.
func (recv *FileChooser) UnselectFilename(filename string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(filename)

	err := fileChooserUnselectFilenameFunction_Set()
	if err == nil {
		fileChooserUnselectFilenameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fileChooserUnselectUriFunction *gi.Function
var fileChooserUnselectUriFunction_Once sync.Once

func fileChooserUnselectUriFunction_Set() error {
	var err error
	fileChooserUnselectUriFunction_Once.Do(func() {
		err = fileChooserInterface_Set()
		if err != nil {
			return
		}
		fileChooserUnselectUriFunction, err = fileChooserInterface.InvokerNew("unselect_uri")
	})
	return err
}

// UnselectUri is a representation of the C type gtk_file_chooser_unselect_uri.
func (recv *FileChooser) UnselectUri(uri string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)

	err := fileChooserUnselectUriFunction_Set()
	if err == nil {
		fileChooserUnselectUriFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectConfirmOverwrite connects a callback to the 'confirm-overwrite' signal of the FileChooser.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *FileChooser) ConnectConfirmOverwrite(handler func(instance *FileChooser) FileChooserConfirmation) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := FileChooserNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "confirm-overwrite", marshal)
}

/*
ConnectCurrentFolderChanged connects a callback to the 'current-folder-changed' signal of the FileChooser.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *FileChooser) ConnectCurrentFolderChanged(handler func(instance *FileChooser)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := FileChooserNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "current-folder-changed", marshal)
}

/*
ConnectFileActivated connects a callback to the 'file-activated' signal of the FileChooser.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *FileChooser) ConnectFileActivated(handler func(instance *FileChooser)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := FileChooserNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "file-activated", marshal)
}

/*
ConnectSelectionChanged connects a callback to the 'selection-changed' signal of the FileChooser.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *FileChooser) ConnectSelectionChanged(handler func(instance *FileChooser)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := FileChooserNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "selection-changed", marshal)
}

/*
ConnectUpdatePreview connects a callback to the 'update-preview' signal of the FileChooser.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *FileChooser) ConnectUpdatePreview(handler func(instance *FileChooser)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := FileChooserNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "update-preview", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *FileChooser) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var fontChooserInterface *gi.Interface
var fontChooserInterface_Once sync.Once

func fontChooserInterface_Set() error {
	var err error
	fontChooserInterface_Once.Do(func() {
		fontChooserInterface, err = gi.InterfaceNew("Gtk", "FontChooser")
	})
	return err
}

type FontChooser struct {
	native unsafe.Pointer
}

func FontChooserNewFromNative(native unsafe.Pointer) *FontChooser {
	instance := &FontChooser{native: native}

	return instance
}

/*
CastToFontChooser down casts any arbitrary Object to FontChooser.
Exercise care, as this is a potentially dangerous function
if the Object is not a FontChooser.
*/
func CastToFontChooser(object *gobject.Object) *FontChooser {
	return FontChooserNewFromNative(object.Native())
}

// Equals compares this FontChooser with another FontChooser, and returns true if they represent the same Object.
func (recv *FontChooser) Equals(other *FontChooser) bool {
	return other.Native() == recv.Native()
}

func (recv *FontChooser) Native() unsafe.Pointer {
	return recv.native
}

var fontChooserGetFontFunction *gi.Function
var fontChooserGetFontFunction_Once sync.Once

func fontChooserGetFontFunction_Set() error {
	var err error
	fontChooserGetFontFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetFontFunction, err = fontChooserInterface.InvokerNew("get_font")
	})
	return err
}

// GetFont is a representation of the C type gtk_font_chooser_get_font.
func (recv *FontChooser) GetFont() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetFontFunction_Set()
	if err == nil {
		ret = fontChooserGetFontFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fontChooserGetFontDescFunction *gi.Function
var fontChooserGetFontDescFunction_Once sync.Once

func fontChooserGetFontDescFunction_Set() error {
	var err error
	fontChooserGetFontDescFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetFontDescFunction, err = fontChooserInterface.InvokerNew("get_font_desc")
	})
	return err
}

// GetFontDesc is a representation of the C type gtk_font_chooser_get_font_desc.
func (recv *FontChooser) GetFontDesc() *pango.FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetFontDescFunction_Set()
	if err == nil {
		ret = fontChooserGetFontDescFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.FontDescriptionNewFromNative(ret.Pointer())

	return retGo
}

var fontChooserGetFontFaceFunction *gi.Function
var fontChooserGetFontFaceFunction_Once sync.Once

func fontChooserGetFontFaceFunction_Set() error {
	var err error
	fontChooserGetFontFaceFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetFontFaceFunction, err = fontChooserInterface.InvokerNew("get_font_face")
	})
	return err
}

// GetFontFace is a representation of the C type gtk_font_chooser_get_font_face.
func (recv *FontChooser) GetFontFace() *pango.FontFace {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetFontFaceFunction_Set()
	if err == nil {
		ret = fontChooserGetFontFaceFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.FontFaceNewFromNative(ret.Pointer())

	return retGo
}

var fontChooserGetFontFamilyFunction *gi.Function
var fontChooserGetFontFamilyFunction_Once sync.Once

func fontChooserGetFontFamilyFunction_Set() error {
	var err error
	fontChooserGetFontFamilyFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetFontFamilyFunction, err = fontChooserInterface.InvokerNew("get_font_family")
	})
	return err
}

// GetFontFamily is a representation of the C type gtk_font_chooser_get_font_family.
func (recv *FontChooser) GetFontFamily() *pango.FontFamily {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetFontFamilyFunction_Set()
	if err == nil {
		ret = fontChooserGetFontFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.FontFamilyNewFromNative(ret.Pointer())

	return retGo
}

var fontChooserGetFontFeaturesFunction *gi.Function
var fontChooserGetFontFeaturesFunction_Once sync.Once

func fontChooserGetFontFeaturesFunction_Set() error {
	var err error
	fontChooserGetFontFeaturesFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetFontFeaturesFunction, err = fontChooserInterface.InvokerNew("get_font_features")
	})
	return err
}

// GetFontFeatures is a representation of the C type gtk_font_chooser_get_font_features.
func (recv *FontChooser) GetFontFeatures() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetFontFeaturesFunction_Set()
	if err == nil {
		ret = fontChooserGetFontFeaturesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fontChooserGetFontMapFunction *gi.Function
var fontChooserGetFontMapFunction_Once sync.Once

func fontChooserGetFontMapFunction_Set() error {
	var err error
	fontChooserGetFontMapFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetFontMapFunction, err = fontChooserInterface.InvokerNew("get_font_map")
	})
	return err
}

// GetFontMap is a representation of the C type gtk_font_chooser_get_font_map.
func (recv *FontChooser) GetFontMap() *pango.FontMap {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetFontMapFunction_Set()
	if err == nil {
		ret = fontChooserGetFontMapFunction.Invoke(inArgs[:], nil)
	}

	retGo := pango.FontMapNewFromNative(ret.Pointer())

	return retGo
}

var fontChooserGetFontSizeFunction *gi.Function
var fontChooserGetFontSizeFunction_Once sync.Once

func fontChooserGetFontSizeFunction_Set() error {
	var err error
	fontChooserGetFontSizeFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetFontSizeFunction, err = fontChooserInterface.InvokerNew("get_font_size")
	})
	return err
}

// GetFontSize is a representation of the C type gtk_font_chooser_get_font_size.
func (recv *FontChooser) GetFontSize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetFontSizeFunction_Set()
	if err == nil {
		ret = fontChooserGetFontSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontChooserGetLanguageFunction *gi.Function
var fontChooserGetLanguageFunction_Once sync.Once

func fontChooserGetLanguageFunction_Set() error {
	var err error
	fontChooserGetLanguageFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetLanguageFunction, err = fontChooserInterface.InvokerNew("get_language")
	})
	return err
}

// GetLanguage is a representation of the C type gtk_font_chooser_get_language.
func (recv *FontChooser) GetLanguage() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetLanguageFunction_Set()
	if err == nil {
		ret = fontChooserGetLanguageFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fontChooserGetLevelFunction *gi.Function
var fontChooserGetLevelFunction_Once sync.Once

func fontChooserGetLevelFunction_Set() error {
	var err error
	fontChooserGetLevelFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetLevelFunction, err = fontChooserInterface.InvokerNew("get_level")
	})
	return err
}

// GetLevel is a representation of the C type gtk_font_chooser_get_level.
func (recv *FontChooser) GetLevel() FontChooserLevel {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetLevelFunction_Set()
	if err == nil {
		ret = fontChooserGetLevelFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontChooserLevel(ret.Int32())

	return retGo
}

var fontChooserGetPreviewTextFunction *gi.Function
var fontChooserGetPreviewTextFunction_Once sync.Once

func fontChooserGetPreviewTextFunction_Set() error {
	var err error
	fontChooserGetPreviewTextFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetPreviewTextFunction, err = fontChooserInterface.InvokerNew("get_preview_text")
	})
	return err
}

// GetPreviewText is a representation of the C type gtk_font_chooser_get_preview_text.
func (recv *FontChooser) GetPreviewText() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetPreviewTextFunction_Set()
	if err == nil {
		ret = fontChooserGetPreviewTextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fontChooserGetShowPreviewEntryFunction *gi.Function
var fontChooserGetShowPreviewEntryFunction_Once sync.Once

func fontChooserGetShowPreviewEntryFunction_Set() error {
	var err error
	fontChooserGetShowPreviewEntryFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserGetShowPreviewEntryFunction, err = fontChooserInterface.InvokerNew("get_show_preview_entry")
	})
	return err
}

// GetShowPreviewEntry is a representation of the C type gtk_font_chooser_get_show_preview_entry.
func (recv *FontChooser) GetShowPreviewEntry() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := fontChooserGetShowPreviewEntryFunction_Set()
	if err == nil {
		ret = fontChooserGetShowPreviewEntryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_font_chooser_set_filter_func' : parameter 'filter' of type 'FontFilterFunc' not supported

var fontChooserSetFontFunction *gi.Function
var fontChooserSetFontFunction_Once sync.Once

func fontChooserSetFontFunction_Set() error {
	var err error
	fontChooserSetFontFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserSetFontFunction, err = fontChooserInterface.InvokerNew("set_font")
	})
	return err
}

// SetFont is a representation of the C type gtk_font_chooser_set_font.
func (recv *FontChooser) SetFont(fontname string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(fontname)

	err := fontChooserSetFontFunction_Set()
	if err == nil {
		fontChooserSetFontFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontChooserSetFontDescFunction *gi.Function
var fontChooserSetFontDescFunction_Once sync.Once

func fontChooserSetFontDescFunction_Set() error {
	var err error
	fontChooserSetFontDescFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserSetFontDescFunction, err = fontChooserInterface.InvokerNew("set_font_desc")
	})
	return err
}

// SetFontDesc is a representation of the C type gtk_font_chooser_set_font_desc.
func (recv *FontChooser) SetFontDesc(fontDesc *pango.FontDescription) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(fontDesc.Native())

	err := fontChooserSetFontDescFunction_Set()
	if err == nil {
		fontChooserSetFontDescFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontChooserSetFontMapFunction *gi.Function
var fontChooserSetFontMapFunction_Once sync.Once

func fontChooserSetFontMapFunction_Set() error {
	var err error
	fontChooserSetFontMapFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserSetFontMapFunction, err = fontChooserInterface.InvokerNew("set_font_map")
	})
	return err
}

// SetFontMap is a representation of the C type gtk_font_chooser_set_font_map.
func (recv *FontChooser) SetFontMap(fontmap *pango.FontMap) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(fontmap.Native())

	err := fontChooserSetFontMapFunction_Set()
	if err == nil {
		fontChooserSetFontMapFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontChooserSetLanguageFunction *gi.Function
var fontChooserSetLanguageFunction_Once sync.Once

func fontChooserSetLanguageFunction_Set() error {
	var err error
	fontChooserSetLanguageFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserSetLanguageFunction, err = fontChooserInterface.InvokerNew("set_language")
	})
	return err
}

// SetLanguage is a representation of the C type gtk_font_chooser_set_language.
func (recv *FontChooser) SetLanguage(language string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(language)

	err := fontChooserSetLanguageFunction_Set()
	if err == nil {
		fontChooserSetLanguageFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontChooserSetLevelFunction *gi.Function
var fontChooserSetLevelFunction_Once sync.Once

func fontChooserSetLevelFunction_Set() error {
	var err error
	fontChooserSetLevelFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserSetLevelFunction, err = fontChooserInterface.InvokerNew("set_level")
	})
	return err
}

// SetLevel is a representation of the C type gtk_font_chooser_set_level.
func (recv *FontChooser) SetLevel(level FontChooserLevel) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(level))

	err := fontChooserSetLevelFunction_Set()
	if err == nil {
		fontChooserSetLevelFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontChooserSetPreviewTextFunction *gi.Function
var fontChooserSetPreviewTextFunction_Once sync.Once

func fontChooserSetPreviewTextFunction_Set() error {
	var err error
	fontChooserSetPreviewTextFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserSetPreviewTextFunction, err = fontChooserInterface.InvokerNew("set_preview_text")
	})
	return err
}

// SetPreviewText is a representation of the C type gtk_font_chooser_set_preview_text.
func (recv *FontChooser) SetPreviewText(text string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(text)

	err := fontChooserSetPreviewTextFunction_Set()
	if err == nil {
		fontChooserSetPreviewTextFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontChooserSetShowPreviewEntryFunction *gi.Function
var fontChooserSetShowPreviewEntryFunction_Once sync.Once

func fontChooserSetShowPreviewEntryFunction_Set() error {
	var err error
	fontChooserSetShowPreviewEntryFunction_Once.Do(func() {
		err = fontChooserInterface_Set()
		if err != nil {
			return
		}
		fontChooserSetShowPreviewEntryFunction, err = fontChooserInterface.InvokerNew("set_show_preview_entry")
	})
	return err
}

// SetShowPreviewEntry is a representation of the C type gtk_font_chooser_set_show_preview_entry.
func (recv *FontChooser) SetShowPreviewEntry(showPreviewEntry bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(showPreviewEntry)

	err := fontChooserSetShowPreviewEntryFunction_Set()
	if err == nil {
		fontChooserSetShowPreviewEntryFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectFontActivated connects a callback to the 'font-activated' signal of the FontChooser.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *FontChooser) ConnectFontActivated(handler func(instance *FontChooser, fontname string)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := FontChooserNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := object1.GetString()

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "font-activated", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *FontChooser) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var orientableInterface *gi.Interface
var orientableInterface_Once sync.Once

func orientableInterface_Set() error {
	var err error
	orientableInterface_Once.Do(func() {
		orientableInterface, err = gi.InterfaceNew("Gtk", "Orientable")
	})
	return err
}

type Orientable struct {
	native unsafe.Pointer
}

func OrientableNewFromNative(native unsafe.Pointer) *Orientable {
	instance := &Orientable{native: native}

	return instance
}

/*
CastToOrientable down casts any arbitrary Object to Orientable.
Exercise care, as this is a potentially dangerous function
if the Object is not a Orientable.
*/
func CastToOrientable(object *gobject.Object) *Orientable {
	return OrientableNewFromNative(object.Native())
}

// Equals compares this Orientable with another Orientable, and returns true if they represent the same Object.
func (recv *Orientable) Equals(other *Orientable) bool {
	return other.Native() == recv.Native()
}

func (recv *Orientable) Native() unsafe.Pointer {
	return recv.native
}

var orientableGetOrientationFunction *gi.Function
var orientableGetOrientationFunction_Once sync.Once

func orientableGetOrientationFunction_Set() error {
	var err error
	orientableGetOrientationFunction_Once.Do(func() {
		err = orientableInterface_Set()
		if err != nil {
			return
		}
		orientableGetOrientationFunction, err = orientableInterface.InvokerNew("get_orientation")
	})
	return err
}

// GetOrientation is a representation of the C type gtk_orientable_get_orientation.
func (recv *Orientable) GetOrientation() Orientation {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := orientableGetOrientationFunction_Set()
	if err == nil {
		ret = orientableGetOrientationFunction.Invoke(inArgs[:], nil)
	}

	retGo := Orientation(ret.Int32())

	return retGo
}

var orientableSetOrientationFunction *gi.Function
var orientableSetOrientationFunction_Once sync.Once

func orientableSetOrientationFunction_Set() error {
	var err error
	orientableSetOrientationFunction_Once.Do(func() {
		err = orientableInterface_Set()
		if err != nil {
			return
		}
		orientableSetOrientationFunction, err = orientableInterface.InvokerNew("set_orientation")
	})
	return err
}

// SetOrientation is a representation of the C type gtk_orientable_set_orientation.
func (recv *Orientable) SetOrientation(orientation Orientation) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(orientation))

	err := orientableSetOrientationFunction_Set()
	if err == nil {
		orientableSetOrientationFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printOperationPreviewInterface *gi.Interface
var printOperationPreviewInterface_Once sync.Once

func printOperationPreviewInterface_Set() error {
	var err error
	printOperationPreviewInterface_Once.Do(func() {
		printOperationPreviewInterface, err = gi.InterfaceNew("Gtk", "PrintOperationPreview")
	})
	return err
}

type PrintOperationPreview struct {
	native unsafe.Pointer
}

func PrintOperationPreviewNewFromNative(native unsafe.Pointer) *PrintOperationPreview {
	instance := &PrintOperationPreview{native: native}

	return instance
}

/*
CastToPrintOperationPreview down casts any arbitrary Object to PrintOperationPreview.
Exercise care, as this is a potentially dangerous function
if the Object is not a PrintOperationPreview.
*/
func CastToPrintOperationPreview(object *gobject.Object) *PrintOperationPreview {
	return PrintOperationPreviewNewFromNative(object.Native())
}

// Equals compares this PrintOperationPreview with another PrintOperationPreview, and returns true if they represent the same Object.
func (recv *PrintOperationPreview) Equals(other *PrintOperationPreview) bool {
	return other.Native() == recv.Native()
}

func (recv *PrintOperationPreview) Native() unsafe.Pointer {
	return recv.native
}

var printOperationPreviewEndPreviewFunction *gi.Function
var printOperationPreviewEndPreviewFunction_Once sync.Once

func printOperationPreviewEndPreviewFunction_Set() error {
	var err error
	printOperationPreviewEndPreviewFunction_Once.Do(func() {
		err = printOperationPreviewInterface_Set()
		if err != nil {
			return
		}
		printOperationPreviewEndPreviewFunction, err = printOperationPreviewInterface.InvokerNew("end_preview")
	})
	return err
}

// EndPreview is a representation of the C type gtk_print_operation_preview_end_preview.
func (recv *PrintOperationPreview) EndPreview() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := printOperationPreviewEndPreviewFunction_Set()
	if err == nil {
		printOperationPreviewEndPreviewFunction.Invoke(inArgs[:], nil)
	}

	return
}

var printOperationPreviewIsSelectedFunction *gi.Function
var printOperationPreviewIsSelectedFunction_Once sync.Once

func printOperationPreviewIsSelectedFunction_Set() error {
	var err error
	printOperationPreviewIsSelectedFunction_Once.Do(func() {
		err = printOperationPreviewInterface_Set()
		if err != nil {
			return
		}
		printOperationPreviewIsSelectedFunction, err = printOperationPreviewInterface.InvokerNew("is_selected")
	})
	return err
}

// IsSelected is a representation of the C type gtk_print_operation_preview_is_selected.
func (recv *PrintOperationPreview) IsSelected(pageNr int32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(pageNr)

	var ret gi.Argument

	err := printOperationPreviewIsSelectedFunction_Set()
	if err == nil {
		ret = printOperationPreviewIsSelectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var printOperationPreviewRenderPageFunction *gi.Function
var printOperationPreviewRenderPageFunction_Once sync.Once

func printOperationPreviewRenderPageFunction_Set() error {
	var err error
	printOperationPreviewRenderPageFunction_Once.Do(func() {
		err = printOperationPreviewInterface_Set()
		if err != nil {
			return
		}
		printOperationPreviewRenderPageFunction, err = printOperationPreviewInterface.InvokerNew("render_page")
	})
	return err
}

// RenderPage is a representation of the C type gtk_print_operation_preview_render_page.
func (recv *PrintOperationPreview) RenderPage(pageNr int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(pageNr)

	err := printOperationPreviewRenderPageFunction_Set()
	if err == nil {
		printOperationPreviewRenderPageFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectGotPageSize connects a callback to the 'got-page-size' signal of the PrintOperationPreview.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *PrintOperationPreview) ConnectGotPageSize(handler func(instance *PrintOperationPreview, context *PrintContext, pageSetup *PageSetup)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := PrintOperationPreviewNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := PrintContextNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := PageSetupNewFromNative(object2.GetObject().Native())

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "got-page-size", marshal)
}

/*
ConnectReady connects a callback to the 'ready' signal of the PrintOperationPreview.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *PrintOperationPreview) ConnectReady(handler func(instance *PrintOperationPreview, context *PrintContext)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := PrintOperationPreviewNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := PrintContextNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "ready", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *PrintOperationPreview) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var recentChooserInterface *gi.Interface
var recentChooserInterface_Once sync.Once

func recentChooserInterface_Set() error {
	var err error
	recentChooserInterface_Once.Do(func() {
		recentChooserInterface, err = gi.InterfaceNew("Gtk", "RecentChooser")
	})
	return err
}

type RecentChooser struct {
	native unsafe.Pointer
}

func RecentChooserNewFromNative(native unsafe.Pointer) *RecentChooser {
	instance := &RecentChooser{native: native}

	return instance
}

/*
CastToRecentChooser down casts any arbitrary Object to RecentChooser.
Exercise care, as this is a potentially dangerous function
if the Object is not a RecentChooser.
*/
func CastToRecentChooser(object *gobject.Object) *RecentChooser {
	return RecentChooserNewFromNative(object.Native())
}

// Equals compares this RecentChooser with another RecentChooser, and returns true if they represent the same Object.
func (recv *RecentChooser) Equals(other *RecentChooser) bool {
	return other.Native() == recv.Native()
}

func (recv *RecentChooser) Native() unsafe.Pointer {
	return recv.native
}

var recentChooserAddFilterFunction *gi.Function
var recentChooserAddFilterFunction_Once sync.Once

func recentChooserAddFilterFunction_Set() error {
	var err error
	recentChooserAddFilterFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserAddFilterFunction, err = recentChooserInterface.InvokerNew("add_filter")
	})
	return err
}

// AddFilter is a representation of the C type gtk_recent_chooser_add_filter.
func (recv *RecentChooser) AddFilter(filter *RecentFilter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(filter.Native())

	err := recentChooserAddFilterFunction_Set()
	if err == nil {
		recentChooserAddFilterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserGetCurrentItemFunction *gi.Function
var recentChooserGetCurrentItemFunction_Once sync.Once

func recentChooserGetCurrentItemFunction_Set() error {
	var err error
	recentChooserGetCurrentItemFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetCurrentItemFunction, err = recentChooserInterface.InvokerNew("get_current_item")
	})
	return err
}

// GetCurrentItem is a representation of the C type gtk_recent_chooser_get_current_item.
func (recv *RecentChooser) GetCurrentItem() *RecentInfo {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetCurrentItemFunction_Set()
	if err == nil {
		ret = recentChooserGetCurrentItemFunction.Invoke(inArgs[:], nil)
	}

	retGo := RecentInfoNewFromNative(ret.Pointer())

	return retGo
}

var recentChooserGetCurrentUriFunction *gi.Function
var recentChooserGetCurrentUriFunction_Once sync.Once

func recentChooserGetCurrentUriFunction_Set() error {
	var err error
	recentChooserGetCurrentUriFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetCurrentUriFunction, err = recentChooserInterface.InvokerNew("get_current_uri")
	})
	return err
}

// GetCurrentUri is a representation of the C type gtk_recent_chooser_get_current_uri.
func (recv *RecentChooser) GetCurrentUri() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetCurrentUriFunction_Set()
	if err == nil {
		ret = recentChooserGetCurrentUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var recentChooserGetFilterFunction *gi.Function
var recentChooserGetFilterFunction_Once sync.Once

func recentChooserGetFilterFunction_Set() error {
	var err error
	recentChooserGetFilterFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetFilterFunction, err = recentChooserInterface.InvokerNew("get_filter")
	})
	return err
}

// GetFilter is a representation of the C type gtk_recent_chooser_get_filter.
func (recv *RecentChooser) GetFilter() *RecentFilter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetFilterFunction_Set()
	if err == nil {
		ret = recentChooserGetFilterFunction.Invoke(inArgs[:], nil)
	}

	retGo := RecentFilterNewFromNative(ret.Pointer())

	return retGo
}

var recentChooserGetItemsFunction *gi.Function
var recentChooserGetItemsFunction_Once sync.Once

func recentChooserGetItemsFunction_Set() error {
	var err error
	recentChooserGetItemsFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetItemsFunction, err = recentChooserInterface.InvokerNew("get_items")
	})
	return err
}

// GetItems is a representation of the C type gtk_recent_chooser_get_items.
func (recv *RecentChooser) GetItems() *glib.List {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetItemsFunction_Set()
	if err == nil {
		ret = recentChooserGetItemsFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.ListNewFromNative(ret.Pointer())

	return retGo
}

var recentChooserGetLimitFunction *gi.Function
var recentChooserGetLimitFunction_Once sync.Once

func recentChooserGetLimitFunction_Set() error {
	var err error
	recentChooserGetLimitFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetLimitFunction, err = recentChooserInterface.InvokerNew("get_limit")
	})
	return err
}

// GetLimit is a representation of the C type gtk_recent_chooser_get_limit.
func (recv *RecentChooser) GetLimit() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetLimitFunction_Set()
	if err == nil {
		ret = recentChooserGetLimitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var recentChooserGetLocalOnlyFunction *gi.Function
var recentChooserGetLocalOnlyFunction_Once sync.Once

func recentChooserGetLocalOnlyFunction_Set() error {
	var err error
	recentChooserGetLocalOnlyFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetLocalOnlyFunction, err = recentChooserInterface.InvokerNew("get_local_only")
	})
	return err
}

// GetLocalOnly is a representation of the C type gtk_recent_chooser_get_local_only.
func (recv *RecentChooser) GetLocalOnly() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetLocalOnlyFunction_Set()
	if err == nil {
		ret = recentChooserGetLocalOnlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentChooserGetSelectMultipleFunction *gi.Function
var recentChooserGetSelectMultipleFunction_Once sync.Once

func recentChooserGetSelectMultipleFunction_Set() error {
	var err error
	recentChooserGetSelectMultipleFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetSelectMultipleFunction, err = recentChooserInterface.InvokerNew("get_select_multiple")
	})
	return err
}

// GetSelectMultiple is a representation of the C type gtk_recent_chooser_get_select_multiple.
func (recv *RecentChooser) GetSelectMultiple() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetSelectMultipleFunction_Set()
	if err == nil {
		ret = recentChooserGetSelectMultipleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentChooserGetShowIconsFunction *gi.Function
var recentChooserGetShowIconsFunction_Once sync.Once

func recentChooserGetShowIconsFunction_Set() error {
	var err error
	recentChooserGetShowIconsFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetShowIconsFunction, err = recentChooserInterface.InvokerNew("get_show_icons")
	})
	return err
}

// GetShowIcons is a representation of the C type gtk_recent_chooser_get_show_icons.
func (recv *RecentChooser) GetShowIcons() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetShowIconsFunction_Set()
	if err == nil {
		ret = recentChooserGetShowIconsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentChooserGetShowNotFoundFunction *gi.Function
var recentChooserGetShowNotFoundFunction_Once sync.Once

func recentChooserGetShowNotFoundFunction_Set() error {
	var err error
	recentChooserGetShowNotFoundFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetShowNotFoundFunction, err = recentChooserInterface.InvokerNew("get_show_not_found")
	})
	return err
}

// GetShowNotFound is a representation of the C type gtk_recent_chooser_get_show_not_found.
func (recv *RecentChooser) GetShowNotFound() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetShowNotFoundFunction_Set()
	if err == nil {
		ret = recentChooserGetShowNotFoundFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentChooserGetShowPrivateFunction *gi.Function
var recentChooserGetShowPrivateFunction_Once sync.Once

func recentChooserGetShowPrivateFunction_Set() error {
	var err error
	recentChooserGetShowPrivateFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetShowPrivateFunction, err = recentChooserInterface.InvokerNew("get_show_private")
	})
	return err
}

// GetShowPrivate is a representation of the C type gtk_recent_chooser_get_show_private.
func (recv *RecentChooser) GetShowPrivate() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetShowPrivateFunction_Set()
	if err == nil {
		ret = recentChooserGetShowPrivateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentChooserGetShowTipsFunction *gi.Function
var recentChooserGetShowTipsFunction_Once sync.Once

func recentChooserGetShowTipsFunction_Set() error {
	var err error
	recentChooserGetShowTipsFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetShowTipsFunction, err = recentChooserInterface.InvokerNew("get_show_tips")
	})
	return err
}

// GetShowTips is a representation of the C type gtk_recent_chooser_get_show_tips.
func (recv *RecentChooser) GetShowTips() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetShowTipsFunction_Set()
	if err == nil {
		ret = recentChooserGetShowTipsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentChooserGetSortTypeFunction *gi.Function
var recentChooserGetSortTypeFunction_Once sync.Once

func recentChooserGetSortTypeFunction_Set() error {
	var err error
	recentChooserGetSortTypeFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetSortTypeFunction, err = recentChooserInterface.InvokerNew("get_sort_type")
	})
	return err
}

// GetSortType is a representation of the C type gtk_recent_chooser_get_sort_type.
func (recv *RecentChooser) GetSortType() RecentSortType {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserGetSortTypeFunction_Set()
	if err == nil {
		ret = recentChooserGetSortTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := RecentSortType(ret.Int32())

	return retGo
}

var recentChooserGetUrisFunction *gi.Function
var recentChooserGetUrisFunction_Once sync.Once

func recentChooserGetUrisFunction_Set() error {
	var err error
	recentChooserGetUrisFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserGetUrisFunction, err = recentChooserInterface.InvokerNew("get_uris")
	})
	return err
}

// GetUris is a representation of the C type gtk_recent_chooser_get_uris.
func (recv *RecentChooser) GetUris() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument

	err := recentChooserGetUrisFunction_Set()
	if err == nil {
		recentChooserGetUrisFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint64()

	return out0
}

var recentChooserListFiltersFunction *gi.Function
var recentChooserListFiltersFunction_Once sync.Once

func recentChooserListFiltersFunction_Set() error {
	var err error
	recentChooserListFiltersFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserListFiltersFunction, err = recentChooserInterface.InvokerNew("list_filters")
	})
	return err
}

// ListFilters is a representation of the C type gtk_recent_chooser_list_filters.
func (recv *RecentChooser) ListFilters() *glib.SList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := recentChooserListFiltersFunction_Set()
	if err == nil {
		ret = recentChooserListFiltersFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.SListNewFromNative(ret.Pointer())

	return retGo
}

var recentChooserRemoveFilterFunction *gi.Function
var recentChooserRemoveFilterFunction_Once sync.Once

func recentChooserRemoveFilterFunction_Set() error {
	var err error
	recentChooserRemoveFilterFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserRemoveFilterFunction, err = recentChooserInterface.InvokerNew("remove_filter")
	})
	return err
}

// RemoveFilter is a representation of the C type gtk_recent_chooser_remove_filter.
func (recv *RecentChooser) RemoveFilter(filter *RecentFilter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(filter.Native())

	err := recentChooserRemoveFilterFunction_Set()
	if err == nil {
		recentChooserRemoveFilterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserSelectAllFunction *gi.Function
var recentChooserSelectAllFunction_Once sync.Once

func recentChooserSelectAllFunction_Set() error {
	var err error
	recentChooserSelectAllFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSelectAllFunction, err = recentChooserInterface.InvokerNew("select_all")
	})
	return err
}

// SelectAll is a representation of the C type gtk_recent_chooser_select_all.
func (recv *RecentChooser) SelectAll() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := recentChooserSelectAllFunction_Set()
	if err == nil {
		recentChooserSelectAllFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserSelectUriFunction *gi.Function
var recentChooserSelectUriFunction_Once sync.Once

func recentChooserSelectUriFunction_Set() error {
	var err error
	recentChooserSelectUriFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSelectUriFunction, err = recentChooserInterface.InvokerNew("select_uri")
	})
	return err
}

// SelectUri is a representation of the C type gtk_recent_chooser_select_uri.
func (recv *RecentChooser) SelectUri(uri string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := recentChooserSelectUriFunction_Set()
	if err == nil {
		ret = recentChooserSelectUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentChooserSetCurrentUriFunction *gi.Function
var recentChooserSetCurrentUriFunction_Once sync.Once

func recentChooserSetCurrentUriFunction_Set() error {
	var err error
	recentChooserSetCurrentUriFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetCurrentUriFunction, err = recentChooserInterface.InvokerNew("set_current_uri")
	})
	return err
}

// SetCurrentUri is a representation of the C type gtk_recent_chooser_set_current_uri.
func (recv *RecentChooser) SetCurrentUri(uri string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)

	var ret gi.Argument

	err := recentChooserSetCurrentUriFunction_Set()
	if err == nil {
		ret = recentChooserSetCurrentUriFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var recentChooserSetFilterFunction *gi.Function
var recentChooserSetFilterFunction_Once sync.Once

func recentChooserSetFilterFunction_Set() error {
	var err error
	recentChooserSetFilterFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetFilterFunction, err = recentChooserInterface.InvokerNew("set_filter")
	})
	return err
}

// SetFilter is a representation of the C type gtk_recent_chooser_set_filter.
func (recv *RecentChooser) SetFilter(filter *RecentFilter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(filter.Native())

	err := recentChooserSetFilterFunction_Set()
	if err == nil {
		recentChooserSetFilterFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserSetLimitFunction *gi.Function
var recentChooserSetLimitFunction_Once sync.Once

func recentChooserSetLimitFunction_Set() error {
	var err error
	recentChooserSetLimitFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetLimitFunction, err = recentChooserInterface.InvokerNew("set_limit")
	})
	return err
}

// SetLimit is a representation of the C type gtk_recent_chooser_set_limit.
func (recv *RecentChooser) SetLimit(limit int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(limit)

	err := recentChooserSetLimitFunction_Set()
	if err == nil {
		recentChooserSetLimitFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserSetLocalOnlyFunction *gi.Function
var recentChooserSetLocalOnlyFunction_Once sync.Once

func recentChooserSetLocalOnlyFunction_Set() error {
	var err error
	recentChooserSetLocalOnlyFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetLocalOnlyFunction, err = recentChooserInterface.InvokerNew("set_local_only")
	})
	return err
}

// SetLocalOnly is a representation of the C type gtk_recent_chooser_set_local_only.
func (recv *RecentChooser) SetLocalOnly(localOnly bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(localOnly)

	err := recentChooserSetLocalOnlyFunction_Set()
	if err == nil {
		recentChooserSetLocalOnlyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserSetSelectMultipleFunction *gi.Function
var recentChooserSetSelectMultipleFunction_Once sync.Once

func recentChooserSetSelectMultipleFunction_Set() error {
	var err error
	recentChooserSetSelectMultipleFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetSelectMultipleFunction, err = recentChooserInterface.InvokerNew("set_select_multiple")
	})
	return err
}

// SetSelectMultiple is a representation of the C type gtk_recent_chooser_set_select_multiple.
func (recv *RecentChooser) SetSelectMultiple(selectMultiple bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(selectMultiple)

	err := recentChooserSetSelectMultipleFunction_Set()
	if err == nil {
		recentChooserSetSelectMultipleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserSetShowIconsFunction *gi.Function
var recentChooserSetShowIconsFunction_Once sync.Once

func recentChooserSetShowIconsFunction_Set() error {
	var err error
	recentChooserSetShowIconsFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetShowIconsFunction, err = recentChooserInterface.InvokerNew("set_show_icons")
	})
	return err
}

// SetShowIcons is a representation of the C type gtk_recent_chooser_set_show_icons.
func (recv *RecentChooser) SetShowIcons(showIcons bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(showIcons)

	err := recentChooserSetShowIconsFunction_Set()
	if err == nil {
		recentChooserSetShowIconsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserSetShowNotFoundFunction *gi.Function
var recentChooserSetShowNotFoundFunction_Once sync.Once

func recentChooserSetShowNotFoundFunction_Set() error {
	var err error
	recentChooserSetShowNotFoundFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetShowNotFoundFunction, err = recentChooserInterface.InvokerNew("set_show_not_found")
	})
	return err
}

// SetShowNotFound is a representation of the C type gtk_recent_chooser_set_show_not_found.
func (recv *RecentChooser) SetShowNotFound(showNotFound bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(showNotFound)

	err := recentChooserSetShowNotFoundFunction_Set()
	if err == nil {
		recentChooserSetShowNotFoundFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserSetShowPrivateFunction *gi.Function
var recentChooserSetShowPrivateFunction_Once sync.Once

func recentChooserSetShowPrivateFunction_Set() error {
	var err error
	recentChooserSetShowPrivateFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetShowPrivateFunction, err = recentChooserInterface.InvokerNew("set_show_private")
	})
	return err
}

// SetShowPrivate is a representation of the C type gtk_recent_chooser_set_show_private.
func (recv *RecentChooser) SetShowPrivate(showPrivate bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(showPrivate)

	err := recentChooserSetShowPrivateFunction_Set()
	if err == nil {
		recentChooserSetShowPrivateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserSetShowTipsFunction *gi.Function
var recentChooserSetShowTipsFunction_Once sync.Once

func recentChooserSetShowTipsFunction_Set() error {
	var err error
	recentChooserSetShowTipsFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetShowTipsFunction, err = recentChooserInterface.InvokerNew("set_show_tips")
	})
	return err
}

// SetShowTips is a representation of the C type gtk_recent_chooser_set_show_tips.
func (recv *RecentChooser) SetShowTips(showTips bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetBoolean(showTips)

	err := recentChooserSetShowTipsFunction_Set()
	if err == nil {
		recentChooserSetShowTipsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_recent_chooser_set_sort_func' : parameter 'sort_func' of type 'RecentSortFunc' not supported

var recentChooserSetSortTypeFunction *gi.Function
var recentChooserSetSortTypeFunction_Once sync.Once

func recentChooserSetSortTypeFunction_Set() error {
	var err error
	recentChooserSetSortTypeFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserSetSortTypeFunction, err = recentChooserInterface.InvokerNew("set_sort_type")
	})
	return err
}

// SetSortType is a representation of the C type gtk_recent_chooser_set_sort_type.
func (recv *RecentChooser) SetSortType(sortType RecentSortType) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(sortType))

	err := recentChooserSetSortTypeFunction_Set()
	if err == nil {
		recentChooserSetSortTypeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserUnselectAllFunction *gi.Function
var recentChooserUnselectAllFunction_Once sync.Once

func recentChooserUnselectAllFunction_Set() error {
	var err error
	recentChooserUnselectAllFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserUnselectAllFunction, err = recentChooserInterface.InvokerNew("unselect_all")
	})
	return err
}

// UnselectAll is a representation of the C type gtk_recent_chooser_unselect_all.
func (recv *RecentChooser) UnselectAll() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := recentChooserUnselectAllFunction_Set()
	if err == nil {
		recentChooserUnselectAllFunction.Invoke(inArgs[:], nil)
	}

	return
}

var recentChooserUnselectUriFunction *gi.Function
var recentChooserUnselectUriFunction_Once sync.Once

func recentChooserUnselectUriFunction_Set() error {
	var err error
	recentChooserUnselectUriFunction_Once.Do(func() {
		err = recentChooserInterface_Set()
		if err != nil {
			return
		}
		recentChooserUnselectUriFunction, err = recentChooserInterface.InvokerNew("unselect_uri")
	})
	return err
}

// UnselectUri is a representation of the C type gtk_recent_chooser_unselect_uri.
func (recv *RecentChooser) UnselectUri(uri string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(uri)

	err := recentChooserUnselectUriFunction_Set()
	if err == nil {
		recentChooserUnselectUriFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectItemActivated connects a callback to the 'item-activated' signal of the RecentChooser.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *RecentChooser) ConnectItemActivated(handler func(instance *RecentChooser)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := RecentChooserNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "item-activated", marshal)
}

/*
ConnectSelectionChanged connects a callback to the 'selection-changed' signal of the RecentChooser.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *RecentChooser) ConnectSelectionChanged(handler func(instance *RecentChooser)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := RecentChooserNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "selection-changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *RecentChooser) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var scrollableInterface *gi.Interface
var scrollableInterface_Once sync.Once

func scrollableInterface_Set() error {
	var err error
	scrollableInterface_Once.Do(func() {
		scrollableInterface, err = gi.InterfaceNew("Gtk", "Scrollable")
	})
	return err
}

type Scrollable struct {
	native unsafe.Pointer
}

func ScrollableNewFromNative(native unsafe.Pointer) *Scrollable {
	instance := &Scrollable{native: native}

	return instance
}

/*
CastToScrollable down casts any arbitrary Object to Scrollable.
Exercise care, as this is a potentially dangerous function
if the Object is not a Scrollable.
*/
func CastToScrollable(object *gobject.Object) *Scrollable {
	return ScrollableNewFromNative(object.Native())
}

// Equals compares this Scrollable with another Scrollable, and returns true if they represent the same Object.
func (recv *Scrollable) Equals(other *Scrollable) bool {
	return other.Native() == recv.Native()
}

func (recv *Scrollable) Native() unsafe.Pointer {
	return recv.native
}

var scrollableGetBorderFunction *gi.Function
var scrollableGetBorderFunction_Once sync.Once

func scrollableGetBorderFunction_Set() error {
	var err error
	scrollableGetBorderFunction_Once.Do(func() {
		err = scrollableInterface_Set()
		if err != nil {
			return
		}
		scrollableGetBorderFunction, err = scrollableInterface.InvokerNew("get_border")
	})
	return err
}

// GetBorder is a representation of the C type gtk_scrollable_get_border.
func (recv *Scrollable) GetBorder() (bool, *Border) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := scrollableGetBorderFunction_Set()
	if err == nil {
		ret = scrollableGetBorderFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := BorderNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var scrollableGetHadjustmentFunction *gi.Function
var scrollableGetHadjustmentFunction_Once sync.Once

func scrollableGetHadjustmentFunction_Set() error {
	var err error
	scrollableGetHadjustmentFunction_Once.Do(func() {
		err = scrollableInterface_Set()
		if err != nil {
			return
		}
		scrollableGetHadjustmentFunction, err = scrollableInterface.InvokerNew("get_hadjustment")
	})
	return err
}

// GetHadjustment is a representation of the C type gtk_scrollable_get_hadjustment.
func (recv *Scrollable) GetHadjustment() *Adjustment {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := scrollableGetHadjustmentFunction_Set()
	if err == nil {
		ret = scrollableGetHadjustmentFunction.Invoke(inArgs[:], nil)
	}

	retGo := AdjustmentNewFromNative(ret.Pointer())

	return retGo
}

var scrollableGetHscrollPolicyFunction *gi.Function
var scrollableGetHscrollPolicyFunction_Once sync.Once

func scrollableGetHscrollPolicyFunction_Set() error {
	var err error
	scrollableGetHscrollPolicyFunction_Once.Do(func() {
		err = scrollableInterface_Set()
		if err != nil {
			return
		}
		scrollableGetHscrollPolicyFunction, err = scrollableInterface.InvokerNew("get_hscroll_policy")
	})
	return err
}

// GetHscrollPolicy is a representation of the C type gtk_scrollable_get_hscroll_policy.
func (recv *Scrollable) GetHscrollPolicy() ScrollablePolicy {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := scrollableGetHscrollPolicyFunction_Set()
	if err == nil {
		ret = scrollableGetHscrollPolicyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ScrollablePolicy(ret.Int32())

	return retGo
}

var scrollableGetVadjustmentFunction *gi.Function
var scrollableGetVadjustmentFunction_Once sync.Once

func scrollableGetVadjustmentFunction_Set() error {
	var err error
	scrollableGetVadjustmentFunction_Once.Do(func() {
		err = scrollableInterface_Set()
		if err != nil {
			return
		}
		scrollableGetVadjustmentFunction, err = scrollableInterface.InvokerNew("get_vadjustment")
	})
	return err
}

// GetVadjustment is a representation of the C type gtk_scrollable_get_vadjustment.
func (recv *Scrollable) GetVadjustment() *Adjustment {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := scrollableGetVadjustmentFunction_Set()
	if err == nil {
		ret = scrollableGetVadjustmentFunction.Invoke(inArgs[:], nil)
	}

	retGo := AdjustmentNewFromNative(ret.Pointer())

	return retGo
}

var scrollableGetVscrollPolicyFunction *gi.Function
var scrollableGetVscrollPolicyFunction_Once sync.Once

func scrollableGetVscrollPolicyFunction_Set() error {
	var err error
	scrollableGetVscrollPolicyFunction_Once.Do(func() {
		err = scrollableInterface_Set()
		if err != nil {
			return
		}
		scrollableGetVscrollPolicyFunction, err = scrollableInterface.InvokerNew("get_vscroll_policy")
	})
	return err
}

// GetVscrollPolicy is a representation of the C type gtk_scrollable_get_vscroll_policy.
func (recv *Scrollable) GetVscrollPolicy() ScrollablePolicy {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := scrollableGetVscrollPolicyFunction_Set()
	if err == nil {
		ret = scrollableGetVscrollPolicyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ScrollablePolicy(ret.Int32())

	return retGo
}

var scrollableSetHadjustmentFunction *gi.Function
var scrollableSetHadjustmentFunction_Once sync.Once

func scrollableSetHadjustmentFunction_Set() error {
	var err error
	scrollableSetHadjustmentFunction_Once.Do(func() {
		err = scrollableInterface_Set()
		if err != nil {
			return
		}
		scrollableSetHadjustmentFunction, err = scrollableInterface.InvokerNew("set_hadjustment")
	})
	return err
}

// SetHadjustment is a representation of the C type gtk_scrollable_set_hadjustment.
func (recv *Scrollable) SetHadjustment(hadjustment *Adjustment) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(hadjustment.Native())

	err := scrollableSetHadjustmentFunction_Set()
	if err == nil {
		scrollableSetHadjustmentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var scrollableSetHscrollPolicyFunction *gi.Function
var scrollableSetHscrollPolicyFunction_Once sync.Once

func scrollableSetHscrollPolicyFunction_Set() error {
	var err error
	scrollableSetHscrollPolicyFunction_Once.Do(func() {
		err = scrollableInterface_Set()
		if err != nil {
			return
		}
		scrollableSetHscrollPolicyFunction, err = scrollableInterface.InvokerNew("set_hscroll_policy")
	})
	return err
}

// SetHscrollPolicy is a representation of the C type gtk_scrollable_set_hscroll_policy.
func (recv *Scrollable) SetHscrollPolicy(policy ScrollablePolicy) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(policy))

	err := scrollableSetHscrollPolicyFunction_Set()
	if err == nil {
		scrollableSetHscrollPolicyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var scrollableSetVadjustmentFunction *gi.Function
var scrollableSetVadjustmentFunction_Once sync.Once

func scrollableSetVadjustmentFunction_Set() error {
	var err error
	scrollableSetVadjustmentFunction_Once.Do(func() {
		err = scrollableInterface_Set()
		if err != nil {
			return
		}
		scrollableSetVadjustmentFunction, err = scrollableInterface.InvokerNew("set_vadjustment")
	})
	return err
}

// SetVadjustment is a representation of the C type gtk_scrollable_set_vadjustment.
func (recv *Scrollable) SetVadjustment(vadjustment *Adjustment) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(vadjustment.Native())

	err := scrollableSetVadjustmentFunction_Set()
	if err == nil {
		scrollableSetVadjustmentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var scrollableSetVscrollPolicyFunction *gi.Function
var scrollableSetVscrollPolicyFunction_Once sync.Once

func scrollableSetVscrollPolicyFunction_Set() error {
	var err error
	scrollableSetVscrollPolicyFunction_Once.Do(func() {
		err = scrollableInterface_Set()
		if err != nil {
			return
		}
		scrollableSetVscrollPolicyFunction, err = scrollableInterface.InvokerNew("set_vscroll_policy")
	})
	return err
}

// SetVscrollPolicy is a representation of the C type gtk_scrollable_set_vscroll_policy.
func (recv *Scrollable) SetVscrollPolicy(policy ScrollablePolicy) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(policy))

	err := scrollableSetVscrollPolicyFunction_Set()
	if err == nil {
		scrollableSetVscrollPolicyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var styleProviderInterface *gi.Interface
var styleProviderInterface_Once sync.Once

func styleProviderInterface_Set() error {
	var err error
	styleProviderInterface_Once.Do(func() {
		styleProviderInterface, err = gi.InterfaceNew("Gtk", "StyleProvider")
	})
	return err
}

type StyleProvider struct {
	native unsafe.Pointer
}

func StyleProviderNewFromNative(native unsafe.Pointer) *StyleProvider {
	instance := &StyleProvider{native: native}

	return instance
}

/*
CastToStyleProvider down casts any arbitrary Object to StyleProvider.
Exercise care, as this is a potentially dangerous function
if the Object is not a StyleProvider.
*/
func CastToStyleProvider(object *gobject.Object) *StyleProvider {
	return StyleProviderNewFromNative(object.Native())
}

// Equals compares this StyleProvider with another StyleProvider, and returns true if they represent the same Object.
func (recv *StyleProvider) Equals(other *StyleProvider) bool {
	return other.Native() == recv.Native()
}

func (recv *StyleProvider) Native() unsafe.Pointer {
	return recv.native
}

var styleProviderGetIconFactoryFunction *gi.Function
var styleProviderGetIconFactoryFunction_Once sync.Once

func styleProviderGetIconFactoryFunction_Set() error {
	var err error
	styleProviderGetIconFactoryFunction_Once.Do(func() {
		err = styleProviderInterface_Set()
		if err != nil {
			return
		}
		styleProviderGetIconFactoryFunction, err = styleProviderInterface.InvokerNew("get_icon_factory")
	})
	return err
}

// GetIconFactory is a representation of the C type gtk_style_provider_get_icon_factory.
func (recv *StyleProvider) GetIconFactory(path *WidgetPath) *IconFactory {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())

	var ret gi.Argument

	err := styleProviderGetIconFactoryFunction_Set()
	if err == nil {
		ret = styleProviderGetIconFactoryFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconFactoryNewFromNative(ret.Pointer())

	return retGo
}

var styleProviderGetStyleFunction *gi.Function
var styleProviderGetStyleFunction_Once sync.Once

func styleProviderGetStyleFunction_Set() error {
	var err error
	styleProviderGetStyleFunction_Once.Do(func() {
		err = styleProviderInterface_Set()
		if err != nil {
			return
		}
		styleProviderGetStyleFunction, err = styleProviderInterface.InvokerNew("get_style")
	})
	return err
}

// GetStyle is a representation of the C type gtk_style_provider_get_style.
func (recv *StyleProvider) GetStyle(path *WidgetPath) *StyleProperties {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())

	var ret gi.Argument

	err := styleProviderGetStyleFunction_Set()
	if err == nil {
		ret = styleProviderGetStyleFunction.Invoke(inArgs[:], nil)
	}

	retGo := StylePropertiesNewFromNative(ret.Pointer())

	return retGo
}

var styleProviderGetStylePropertyFunction *gi.Function
var styleProviderGetStylePropertyFunction_Once sync.Once

func styleProviderGetStylePropertyFunction_Set() error {
	var err error
	styleProviderGetStylePropertyFunction_Once.Do(func() {
		err = styleProviderInterface_Set()
		if err != nil {
			return
		}
		styleProviderGetStylePropertyFunction, err = styleProviderInterface.InvokerNew("get_style_property")
	})
	return err
}

// GetStyleProperty is a representation of the C type gtk_style_provider_get_style_property.
func (recv *StyleProvider) GetStyleProperty(path *WidgetPath, state StateFlags, pspec *gobject.ParamSpec) (bool, *gobject.Value) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())
	inArgs[2].SetInt32(int32(state))
	inArgs[3].SetPointer(pspec.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := styleProviderGetStylePropertyFunction_Set()
	if err == nil {
		ret = styleProviderGetStylePropertyFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := gobject.ValueNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var toolShellInterface *gi.Interface
var toolShellInterface_Once sync.Once

func toolShellInterface_Set() error {
	var err error
	toolShellInterface_Once.Do(func() {
		toolShellInterface, err = gi.InterfaceNew("Gtk", "ToolShell")
	})
	return err
}

type ToolShell struct {
	native unsafe.Pointer
}

func ToolShellNewFromNative(native unsafe.Pointer) *ToolShell {
	instance := &ToolShell{native: native}

	return instance
}

/*
CastToToolShell down casts any arbitrary Object to ToolShell.
Exercise care, as this is a potentially dangerous function
if the Object is not a ToolShell.
*/
func CastToToolShell(object *gobject.Object) *ToolShell {
	return ToolShellNewFromNative(object.Native())
}

// Equals compares this ToolShell with another ToolShell, and returns true if they represent the same Object.
func (recv *ToolShell) Equals(other *ToolShell) bool {
	return other.Native() == recv.Native()
}

func (recv *ToolShell) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'gtk_tool_shell_get_ellipsize_mode' : return type 'Pango.EllipsizeMode' not supported

var toolShellGetIconSizeFunction *gi.Function
var toolShellGetIconSizeFunction_Once sync.Once

func toolShellGetIconSizeFunction_Set() error {
	var err error
	toolShellGetIconSizeFunction_Once.Do(func() {
		err = toolShellInterface_Set()
		if err != nil {
			return
		}
		toolShellGetIconSizeFunction, err = toolShellInterface.InvokerNew("get_icon_size")
	})
	return err
}

// GetIconSize is a representation of the C type gtk_tool_shell_get_icon_size.
func (recv *ToolShell) GetIconSize() IconSize {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := toolShellGetIconSizeFunction_Set()
	if err == nil {
		ret = toolShellGetIconSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := IconSize(ret.Int32())

	return retGo
}

var toolShellGetOrientationFunction *gi.Function
var toolShellGetOrientationFunction_Once sync.Once

func toolShellGetOrientationFunction_Set() error {
	var err error
	toolShellGetOrientationFunction_Once.Do(func() {
		err = toolShellInterface_Set()
		if err != nil {
			return
		}
		toolShellGetOrientationFunction, err = toolShellInterface.InvokerNew("get_orientation")
	})
	return err
}

// GetOrientation is a representation of the C type gtk_tool_shell_get_orientation.
func (recv *ToolShell) GetOrientation() Orientation {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := toolShellGetOrientationFunction_Set()
	if err == nil {
		ret = toolShellGetOrientationFunction.Invoke(inArgs[:], nil)
	}

	retGo := Orientation(ret.Int32())

	return retGo
}

var toolShellGetReliefStyleFunction *gi.Function
var toolShellGetReliefStyleFunction_Once sync.Once

func toolShellGetReliefStyleFunction_Set() error {
	var err error
	toolShellGetReliefStyleFunction_Once.Do(func() {
		err = toolShellInterface_Set()
		if err != nil {
			return
		}
		toolShellGetReliefStyleFunction, err = toolShellInterface.InvokerNew("get_relief_style")
	})
	return err
}

// GetReliefStyle is a representation of the C type gtk_tool_shell_get_relief_style.
func (recv *ToolShell) GetReliefStyle() ReliefStyle {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := toolShellGetReliefStyleFunction_Set()
	if err == nil {
		ret = toolShellGetReliefStyleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ReliefStyle(ret.Int32())

	return retGo
}

var toolShellGetStyleFunction *gi.Function
var toolShellGetStyleFunction_Once sync.Once

func toolShellGetStyleFunction_Set() error {
	var err error
	toolShellGetStyleFunction_Once.Do(func() {
		err = toolShellInterface_Set()
		if err != nil {
			return
		}
		toolShellGetStyleFunction, err = toolShellInterface.InvokerNew("get_style")
	})
	return err
}

// GetStyle is a representation of the C type gtk_tool_shell_get_style.
func (recv *ToolShell) GetStyle() ToolbarStyle {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := toolShellGetStyleFunction_Set()
	if err == nil {
		ret = toolShellGetStyleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ToolbarStyle(ret.Int32())

	return retGo
}

var toolShellGetTextAlignmentFunction *gi.Function
var toolShellGetTextAlignmentFunction_Once sync.Once

func toolShellGetTextAlignmentFunction_Set() error {
	var err error
	toolShellGetTextAlignmentFunction_Once.Do(func() {
		err = toolShellInterface_Set()
		if err != nil {
			return
		}
		toolShellGetTextAlignmentFunction, err = toolShellInterface.InvokerNew("get_text_alignment")
	})
	return err
}

// GetTextAlignment is a representation of the C type gtk_tool_shell_get_text_alignment.
func (recv *ToolShell) GetTextAlignment() float32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := toolShellGetTextAlignmentFunction_Set()
	if err == nil {
		ret = toolShellGetTextAlignmentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float32()

	return retGo
}

var toolShellGetTextOrientationFunction *gi.Function
var toolShellGetTextOrientationFunction_Once sync.Once

func toolShellGetTextOrientationFunction_Set() error {
	var err error
	toolShellGetTextOrientationFunction_Once.Do(func() {
		err = toolShellInterface_Set()
		if err != nil {
			return
		}
		toolShellGetTextOrientationFunction, err = toolShellInterface.InvokerNew("get_text_orientation")
	})
	return err
}

// GetTextOrientation is a representation of the C type gtk_tool_shell_get_text_orientation.
func (recv *ToolShell) GetTextOrientation() Orientation {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := toolShellGetTextOrientationFunction_Set()
	if err == nil {
		ret = toolShellGetTextOrientationFunction.Invoke(inArgs[:], nil)
	}

	retGo := Orientation(ret.Int32())

	return retGo
}

var toolShellGetTextSizeGroupFunction *gi.Function
var toolShellGetTextSizeGroupFunction_Once sync.Once

func toolShellGetTextSizeGroupFunction_Set() error {
	var err error
	toolShellGetTextSizeGroupFunction_Once.Do(func() {
		err = toolShellInterface_Set()
		if err != nil {
			return
		}
		toolShellGetTextSizeGroupFunction, err = toolShellInterface.InvokerNew("get_text_size_group")
	})
	return err
}

// GetTextSizeGroup is a representation of the C type gtk_tool_shell_get_text_size_group.
func (recv *ToolShell) GetTextSizeGroup() *SizeGroup {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := toolShellGetTextSizeGroupFunction_Set()
	if err == nil {
		ret = toolShellGetTextSizeGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := SizeGroupNewFromNative(ret.Pointer())

	return retGo
}

var toolShellRebuildMenuFunction *gi.Function
var toolShellRebuildMenuFunction_Once sync.Once

func toolShellRebuildMenuFunction_Set() error {
	var err error
	toolShellRebuildMenuFunction_Once.Do(func() {
		err = toolShellInterface_Set()
		if err != nil {
			return
		}
		toolShellRebuildMenuFunction, err = toolShellInterface.InvokerNew("rebuild_menu")
	})
	return err
}

// RebuildMenu is a representation of the C type gtk_tool_shell_rebuild_menu.
func (recv *ToolShell) RebuildMenu() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := toolShellRebuildMenuFunction_Set()
	if err == nil {
		toolShellRebuildMenuFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treeDragDestInterface *gi.Interface
var treeDragDestInterface_Once sync.Once

func treeDragDestInterface_Set() error {
	var err error
	treeDragDestInterface_Once.Do(func() {
		treeDragDestInterface, err = gi.InterfaceNew("Gtk", "TreeDragDest")
	})
	return err
}

type TreeDragDest struct {
	native unsafe.Pointer
}

func TreeDragDestNewFromNative(native unsafe.Pointer) *TreeDragDest {
	instance := &TreeDragDest{native: native}

	return instance
}

/*
CastToTreeDragDest down casts any arbitrary Object to TreeDragDest.
Exercise care, as this is a potentially dangerous function
if the Object is not a TreeDragDest.
*/
func CastToTreeDragDest(object *gobject.Object) *TreeDragDest {
	return TreeDragDestNewFromNative(object.Native())
}

// Equals compares this TreeDragDest with another TreeDragDest, and returns true if they represent the same Object.
func (recv *TreeDragDest) Equals(other *TreeDragDest) bool {
	return other.Native() == recv.Native()
}

func (recv *TreeDragDest) Native() unsafe.Pointer {
	return recv.native
}

var treeDragDestDragDataReceivedFunction *gi.Function
var treeDragDestDragDataReceivedFunction_Once sync.Once

func treeDragDestDragDataReceivedFunction_Set() error {
	var err error
	treeDragDestDragDataReceivedFunction_Once.Do(func() {
		err = treeDragDestInterface_Set()
		if err != nil {
			return
		}
		treeDragDestDragDataReceivedFunction, err = treeDragDestInterface.InvokerNew("drag_data_received")
	})
	return err
}

// DragDataReceived is a representation of the C type gtk_tree_drag_dest_drag_data_received.
func (recv *TreeDragDest) DragDataReceived(dest *TreePath, selectionData *SelectionData) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(dest.Native())
	inArgs[2].SetPointer(selectionData.Native())

	var ret gi.Argument

	err := treeDragDestDragDataReceivedFunction_Set()
	if err == nil {
		ret = treeDragDestDragDataReceivedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeDragDestRowDropPossibleFunction *gi.Function
var treeDragDestRowDropPossibleFunction_Once sync.Once

func treeDragDestRowDropPossibleFunction_Set() error {
	var err error
	treeDragDestRowDropPossibleFunction_Once.Do(func() {
		err = treeDragDestInterface_Set()
		if err != nil {
			return
		}
		treeDragDestRowDropPossibleFunction, err = treeDragDestInterface.InvokerNew("row_drop_possible")
	})
	return err
}

// RowDropPossible is a representation of the C type gtk_tree_drag_dest_row_drop_possible.
func (recv *TreeDragDest) RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(destPath.Native())
	inArgs[2].SetPointer(selectionData.Native())

	var ret gi.Argument

	err := treeDragDestRowDropPossibleFunction_Set()
	if err == nil {
		ret = treeDragDestRowDropPossibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeDragSourceInterface *gi.Interface
var treeDragSourceInterface_Once sync.Once

func treeDragSourceInterface_Set() error {
	var err error
	treeDragSourceInterface_Once.Do(func() {
		treeDragSourceInterface, err = gi.InterfaceNew("Gtk", "TreeDragSource")
	})
	return err
}

type TreeDragSource struct {
	native unsafe.Pointer
}

func TreeDragSourceNewFromNative(native unsafe.Pointer) *TreeDragSource {
	instance := &TreeDragSource{native: native}

	return instance
}

/*
CastToTreeDragSource down casts any arbitrary Object to TreeDragSource.
Exercise care, as this is a potentially dangerous function
if the Object is not a TreeDragSource.
*/
func CastToTreeDragSource(object *gobject.Object) *TreeDragSource {
	return TreeDragSourceNewFromNative(object.Native())
}

// Equals compares this TreeDragSource with another TreeDragSource, and returns true if they represent the same Object.
func (recv *TreeDragSource) Equals(other *TreeDragSource) bool {
	return other.Native() == recv.Native()
}

func (recv *TreeDragSource) Native() unsafe.Pointer {
	return recv.native
}

var treeDragSourceDragDataDeleteFunction *gi.Function
var treeDragSourceDragDataDeleteFunction_Once sync.Once

func treeDragSourceDragDataDeleteFunction_Set() error {
	var err error
	treeDragSourceDragDataDeleteFunction_Once.Do(func() {
		err = treeDragSourceInterface_Set()
		if err != nil {
			return
		}
		treeDragSourceDragDataDeleteFunction, err = treeDragSourceInterface.InvokerNew("drag_data_delete")
	})
	return err
}

// DragDataDelete is a representation of the C type gtk_tree_drag_source_drag_data_delete.
func (recv *TreeDragSource) DragDataDelete(path *TreePath) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())

	var ret gi.Argument

	err := treeDragSourceDragDataDeleteFunction_Set()
	if err == nil {
		ret = treeDragSourceDragDataDeleteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeDragSourceDragDataGetFunction *gi.Function
var treeDragSourceDragDataGetFunction_Once sync.Once

func treeDragSourceDragDataGetFunction_Set() error {
	var err error
	treeDragSourceDragDataGetFunction_Once.Do(func() {
		err = treeDragSourceInterface_Set()
		if err != nil {
			return
		}
		treeDragSourceDragDataGetFunction, err = treeDragSourceInterface.InvokerNew("drag_data_get")
	})
	return err
}

// DragDataGet is a representation of the C type gtk_tree_drag_source_drag_data_get.
func (recv *TreeDragSource) DragDataGet(path *TreePath, selectionData *SelectionData) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())
	inArgs[2].SetPointer(selectionData.Native())

	var ret gi.Argument

	err := treeDragSourceDragDataGetFunction_Set()
	if err == nil {
		ret = treeDragSourceDragDataGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeDragSourceRowDraggableFunction *gi.Function
var treeDragSourceRowDraggableFunction_Once sync.Once

func treeDragSourceRowDraggableFunction_Set() error {
	var err error
	treeDragSourceRowDraggableFunction_Once.Do(func() {
		err = treeDragSourceInterface_Set()
		if err != nil {
			return
		}
		treeDragSourceRowDraggableFunction, err = treeDragSourceInterface.InvokerNew("row_draggable")
	})
	return err
}

// RowDraggable is a representation of the C type gtk_tree_drag_source_row_draggable.
func (recv *TreeDragSource) RowDraggable(path *TreePath) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())

	var ret gi.Argument

	err := treeDragSourceRowDraggableFunction_Set()
	if err == nil {
		ret = treeDragSourceRowDraggableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeModelInterface *gi.Interface
var treeModelInterface_Once sync.Once

func treeModelInterface_Set() error {
	var err error
	treeModelInterface_Once.Do(func() {
		treeModelInterface, err = gi.InterfaceNew("Gtk", "TreeModel")
	})
	return err
}

type TreeModel struct {
	native unsafe.Pointer
}

func TreeModelNewFromNative(native unsafe.Pointer) *TreeModel {
	instance := &TreeModel{native: native}

	return instance
}

/*
CastToTreeModel down casts any arbitrary Object to TreeModel.
Exercise care, as this is a potentially dangerous function
if the Object is not a TreeModel.
*/
func CastToTreeModel(object *gobject.Object) *TreeModel {
	return TreeModelNewFromNative(object.Native())
}

// Equals compares this TreeModel with another TreeModel, and returns true if they represent the same Object.
func (recv *TreeModel) Equals(other *TreeModel) bool {
	return other.Native() == recv.Native()
}

func (recv *TreeModel) Native() unsafe.Pointer {
	return recv.native
}

// UNSUPPORTED : C value 'gtk_tree_model_filter_new' : return type 'TreeModel' not supported

// UNSUPPORTED : C value 'gtk_tree_model_foreach' : parameter 'func' of type 'TreeModelForeachFunc' not supported

// UNSUPPORTED : C value 'gtk_tree_model_get' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_tree_model_get_column_type' : return type 'GType' not supported

var treeModelGetFlagsFunction *gi.Function
var treeModelGetFlagsFunction_Once sync.Once

func treeModelGetFlagsFunction_Set() error {
	var err error
	treeModelGetFlagsFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelGetFlagsFunction, err = treeModelInterface.InvokerNew("get_flags")
	})
	return err
}

// GetFlags is a representation of the C type gtk_tree_model_get_flags.
func (recv *TreeModel) GetFlags() TreeModelFlags {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := treeModelGetFlagsFunction_Set()
	if err == nil {
		ret = treeModelGetFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := TreeModelFlags(ret.Int32())

	return retGo
}

var treeModelGetIterFunction *gi.Function
var treeModelGetIterFunction_Once sync.Once

func treeModelGetIterFunction_Set() error {
	var err error
	treeModelGetIterFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelGetIterFunction, err = treeModelInterface.InvokerNew("get_iter")
	})
	return err
}

// GetIter is a representation of the C type gtk_tree_model_get_iter.
func (recv *TreeModel) GetIter(path *TreePath) (bool, *TreeIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := treeModelGetIterFunction_Set()
	if err == nil {
		ret = treeModelGetIterFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := TreeIterNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var treeModelGetIterFirstFunction *gi.Function
var treeModelGetIterFirstFunction_Once sync.Once

func treeModelGetIterFirstFunction_Set() error {
	var err error
	treeModelGetIterFirstFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelGetIterFirstFunction, err = treeModelInterface.InvokerNew("get_iter_first")
	})
	return err
}

// GetIterFirst is a representation of the C type gtk_tree_model_get_iter_first.
func (recv *TreeModel) GetIterFirst() (bool, *TreeIter) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := treeModelGetIterFirstFunction_Set()
	if err == nil {
		ret = treeModelGetIterFirstFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := TreeIterNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var treeModelGetIterFromStringFunction *gi.Function
var treeModelGetIterFromStringFunction_Once sync.Once

func treeModelGetIterFromStringFunction_Set() error {
	var err error
	treeModelGetIterFromStringFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelGetIterFromStringFunction, err = treeModelInterface.InvokerNew("get_iter_from_string")
	})
	return err
}

// GetIterFromString is a representation of the C type gtk_tree_model_get_iter_from_string.
func (recv *TreeModel) GetIterFromString(pathString string) (bool, *TreeIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetString(pathString)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := treeModelGetIterFromStringFunction_Set()
	if err == nil {
		ret = treeModelGetIterFromStringFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := TreeIterNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var treeModelGetNColumnsFunction *gi.Function
var treeModelGetNColumnsFunction_Once sync.Once

func treeModelGetNColumnsFunction_Set() error {
	var err error
	treeModelGetNColumnsFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelGetNColumnsFunction, err = treeModelInterface.InvokerNew("get_n_columns")
	})
	return err
}

// GetNColumns is a representation of the C type gtk_tree_model_get_n_columns.
func (recv *TreeModel) GetNColumns() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := treeModelGetNColumnsFunction_Set()
	if err == nil {
		ret = treeModelGetNColumnsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var treeModelGetPathFunction *gi.Function
var treeModelGetPathFunction_Once sync.Once

func treeModelGetPathFunction_Set() error {
	var err error
	treeModelGetPathFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelGetPathFunction, err = treeModelInterface.InvokerNew("get_path")
	})
	return err
}

// GetPath is a representation of the C type gtk_tree_model_get_path.
func (recv *TreeModel) GetPath(iter *TreeIter) *TreePath {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var ret gi.Argument

	err := treeModelGetPathFunction_Set()
	if err == nil {
		ret = treeModelGetPathFunction.Invoke(inArgs[:], nil)
	}

	retGo := TreePathNewFromNative(ret.Pointer())

	return retGo
}

var treeModelGetStringFromIterFunction *gi.Function
var treeModelGetStringFromIterFunction_Once sync.Once

func treeModelGetStringFromIterFunction_Set() error {
	var err error
	treeModelGetStringFromIterFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelGetStringFromIterFunction, err = treeModelInterface.InvokerNew("get_string_from_iter")
	})
	return err
}

// GetStringFromIter is a representation of the C type gtk_tree_model_get_string_from_iter.
func (recv *TreeModel) GetStringFromIter(iter *TreeIter) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var ret gi.Argument

	err := treeModelGetStringFromIterFunction_Set()
	if err == nil {
		ret = treeModelGetStringFromIterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gtk_tree_model_get_valist' : parameter 'var_args' of type 'va_list' not supported

var treeModelGetValueFunction *gi.Function
var treeModelGetValueFunction_Once sync.Once

func treeModelGetValueFunction_Set() error {
	var err error
	treeModelGetValueFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelGetValueFunction, err = treeModelInterface.InvokerNew("get_value")
	})
	return err
}

// GetValue is a representation of the C type gtk_tree_model_get_value.
func (recv *TreeModel) GetValue(iter *TreeIter, column int32) *gobject.Value {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())
	inArgs[2].SetInt32(column)

	var outArgs [1]gi.Argument

	err := treeModelGetValueFunction_Set()
	if err == nil {
		treeModelGetValueFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := gobject.ValueNewFromNative(outArgs[0].Pointer())

	return out0
}

var treeModelIterChildrenFunction *gi.Function
var treeModelIterChildrenFunction_Once sync.Once

func treeModelIterChildrenFunction_Set() error {
	var err error
	treeModelIterChildrenFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelIterChildrenFunction, err = treeModelInterface.InvokerNew("iter_children")
	})
	return err
}

// IterChildren is a representation of the C type gtk_tree_model_iter_children.
func (recv *TreeModel) IterChildren(parent *TreeIter) (bool, *TreeIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(parent.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := treeModelIterChildrenFunction_Set()
	if err == nil {
		ret = treeModelIterChildrenFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := TreeIterNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var treeModelIterHasChildFunction *gi.Function
var treeModelIterHasChildFunction_Once sync.Once

func treeModelIterHasChildFunction_Set() error {
	var err error
	treeModelIterHasChildFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelIterHasChildFunction, err = treeModelInterface.InvokerNew("iter_has_child")
	})
	return err
}

// IterHasChild is a representation of the C type gtk_tree_model_iter_has_child.
func (recv *TreeModel) IterHasChild(iter *TreeIter) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var ret gi.Argument

	err := treeModelIterHasChildFunction_Set()
	if err == nil {
		ret = treeModelIterHasChildFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeModelIterNChildrenFunction *gi.Function
var treeModelIterNChildrenFunction_Once sync.Once

func treeModelIterNChildrenFunction_Set() error {
	var err error
	treeModelIterNChildrenFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelIterNChildrenFunction, err = treeModelInterface.InvokerNew("iter_n_children")
	})
	return err
}

// IterNChildren is a representation of the C type gtk_tree_model_iter_n_children.
func (recv *TreeModel) IterNChildren(iter *TreeIter) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var ret gi.Argument

	err := treeModelIterNChildrenFunction_Set()
	if err == nil {
		ret = treeModelIterNChildrenFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var treeModelIterNextFunction *gi.Function
var treeModelIterNextFunction_Once sync.Once

func treeModelIterNextFunction_Set() error {
	var err error
	treeModelIterNextFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelIterNextFunction, err = treeModelInterface.InvokerNew("iter_next")
	})
	return err
}

// IterNext is a representation of the C type gtk_tree_model_iter_next.
func (recv *TreeModel) IterNext(iter *TreeIter) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var ret gi.Argument

	err := treeModelIterNextFunction_Set()
	if err == nil {
		ret = treeModelIterNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeModelIterNthChildFunction *gi.Function
var treeModelIterNthChildFunction_Once sync.Once

func treeModelIterNthChildFunction_Set() error {
	var err error
	treeModelIterNthChildFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelIterNthChildFunction, err = treeModelInterface.InvokerNew("iter_nth_child")
	})
	return err
}

// IterNthChild is a representation of the C type gtk_tree_model_iter_nth_child.
func (recv *TreeModel) IterNthChild(parent *TreeIter, n int32) (bool, *TreeIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(parent.Native())
	inArgs[2].SetInt32(n)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := treeModelIterNthChildFunction_Set()
	if err == nil {
		ret = treeModelIterNthChildFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := TreeIterNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var treeModelIterParentFunction *gi.Function
var treeModelIterParentFunction_Once sync.Once

func treeModelIterParentFunction_Set() error {
	var err error
	treeModelIterParentFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelIterParentFunction, err = treeModelInterface.InvokerNew("iter_parent")
	})
	return err
}

// IterParent is a representation of the C type gtk_tree_model_iter_parent.
func (recv *TreeModel) IterParent(child *TreeIter) (bool, *TreeIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(child.Native())

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := treeModelIterParentFunction_Set()
	if err == nil {
		ret = treeModelIterParentFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := TreeIterNewFromNative(outArgs[0].Pointer())

	return retGo, out0
}

var treeModelIterPreviousFunction *gi.Function
var treeModelIterPreviousFunction_Once sync.Once

func treeModelIterPreviousFunction_Set() error {
	var err error
	treeModelIterPreviousFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelIterPreviousFunction, err = treeModelInterface.InvokerNew("iter_previous")
	})
	return err
}

// IterPrevious is a representation of the C type gtk_tree_model_iter_previous.
func (recv *TreeModel) IterPrevious(iter *TreeIter) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	var ret gi.Argument

	err := treeModelIterPreviousFunction_Set()
	if err == nil {
		ret = treeModelIterPreviousFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var treeModelRefNodeFunction *gi.Function
var treeModelRefNodeFunction_Once sync.Once

func treeModelRefNodeFunction_Set() error {
	var err error
	treeModelRefNodeFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelRefNodeFunction, err = treeModelInterface.InvokerNew("ref_node")
	})
	return err
}

// RefNode is a representation of the C type gtk_tree_model_ref_node.
func (recv *TreeModel) RefNode(iter *TreeIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	err := treeModelRefNodeFunction_Set()
	if err == nil {
		treeModelRefNodeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treeModelRowChangedFunction *gi.Function
var treeModelRowChangedFunction_Once sync.Once

func treeModelRowChangedFunction_Set() error {
	var err error
	treeModelRowChangedFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelRowChangedFunction, err = treeModelInterface.InvokerNew("row_changed")
	})
	return err
}

// RowChanged is a representation of the C type gtk_tree_model_row_changed.
func (recv *TreeModel) RowChanged(path *TreePath, iter *TreeIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())
	inArgs[2].SetPointer(iter.Native())

	err := treeModelRowChangedFunction_Set()
	if err == nil {
		treeModelRowChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treeModelRowDeletedFunction *gi.Function
var treeModelRowDeletedFunction_Once sync.Once

func treeModelRowDeletedFunction_Set() error {
	var err error
	treeModelRowDeletedFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelRowDeletedFunction, err = treeModelInterface.InvokerNew("row_deleted")
	})
	return err
}

// RowDeleted is a representation of the C type gtk_tree_model_row_deleted.
func (recv *TreeModel) RowDeleted(path *TreePath) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())

	err := treeModelRowDeletedFunction_Set()
	if err == nil {
		treeModelRowDeletedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treeModelRowHasChildToggledFunction *gi.Function
var treeModelRowHasChildToggledFunction_Once sync.Once

func treeModelRowHasChildToggledFunction_Set() error {
	var err error
	treeModelRowHasChildToggledFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelRowHasChildToggledFunction, err = treeModelInterface.InvokerNew("row_has_child_toggled")
	})
	return err
}

// RowHasChildToggled is a representation of the C type gtk_tree_model_row_has_child_toggled.
func (recv *TreeModel) RowHasChildToggled(path *TreePath, iter *TreeIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())
	inArgs[2].SetPointer(iter.Native())

	err := treeModelRowHasChildToggledFunction_Set()
	if err == nil {
		treeModelRowHasChildToggledFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treeModelRowInsertedFunction *gi.Function
var treeModelRowInsertedFunction_Once sync.Once

func treeModelRowInsertedFunction_Set() error {
	var err error
	treeModelRowInsertedFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelRowInsertedFunction, err = treeModelInterface.InvokerNew("row_inserted")
	})
	return err
}

// RowInserted is a representation of the C type gtk_tree_model_row_inserted.
func (recv *TreeModel) RowInserted(path *TreePath, iter *TreeIter) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())
	inArgs[2].SetPointer(iter.Native())

	err := treeModelRowInsertedFunction_Set()
	if err == nil {
		treeModelRowInsertedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var treeModelRowsReorderedFunction *gi.Function
var treeModelRowsReorderedFunction_Once sync.Once

func treeModelRowsReorderedFunction_Set() error {
	var err error
	treeModelRowsReorderedFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelRowsReorderedFunction, err = treeModelInterface.InvokerNew("rows_reordered")
	})
	return err
}

// RowsReordered is a representation of the C type gtk_tree_model_rows_reordered.
func (recv *TreeModel) RowsReordered(path *TreePath, iter *TreeIter, newOrder int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(path.Native())
	inArgs[2].SetPointer(iter.Native())
	inArgs[3].SetInt32(newOrder)

	err := treeModelRowsReorderedFunction_Set()
	if err == nil {
		treeModelRowsReorderedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_tree_model_rows_reordered_with_length' : parameter 'new_order' of type 'nil' not supported

// UNSUPPORTED : C value 'gtk_tree_model_sort_new_with_model' : return type 'TreeModel' not supported

var treeModelUnrefNodeFunction *gi.Function
var treeModelUnrefNodeFunction_Once sync.Once

func treeModelUnrefNodeFunction_Set() error {
	var err error
	treeModelUnrefNodeFunction_Once.Do(func() {
		err = treeModelInterface_Set()
		if err != nil {
			return
		}
		treeModelUnrefNodeFunction, err = treeModelInterface.InvokerNew("unref_node")
	})
	return err
}

// UnrefNode is a representation of the C type gtk_tree_model_unref_node.
func (recv *TreeModel) UnrefNode(iter *TreeIter) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetPointer(iter.Native())

	err := treeModelUnrefNodeFunction_Set()
	if err == nil {
		treeModelUnrefNodeFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectRowChanged connects a callback to the 'row-changed' signal of the TreeModel.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *TreeModel) ConnectRowChanged(handler func(instance *TreeModel, path *TreePath, iter *TreeIter)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := TreeModelNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := TreePathNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := TreeIterNewFromNative(object2.GetObject().Native())

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "row-changed", marshal)
}

/*
ConnectRowDeleted connects a callback to the 'row-deleted' signal of the TreeModel.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *TreeModel) ConnectRowDeleted(handler func(instance *TreeModel, path *TreePath)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := TreeModelNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := TreePathNewFromNative(object1.GetObject().Native())

		handler(argInstance, arg1)
	}

	return callback.ConnectSignal(recv.Native(), "row-deleted", marshal)
}

/*
ConnectRowHasChildToggled connects a callback to the 'row-has-child-toggled' signal of the TreeModel.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *TreeModel) ConnectRowHasChildToggled(handler func(instance *TreeModel, path *TreePath, iter *TreeIter)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := TreeModelNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := TreePathNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := TreeIterNewFromNative(object2.GetObject().Native())

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "row-has-child-toggled", marshal)
}

/*
ConnectRowInserted connects a callback to the 'row-inserted' signal of the TreeModel.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *TreeModel) ConnectRowInserted(handler func(instance *TreeModel, path *TreePath, iter *TreeIter)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := TreeModelNewFromNative(objectInstance.GetObject().Native())

		object1 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[1]))
		arg1 := TreePathNewFromNative(object1.GetObject().Native())

		object2 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[2]))
		arg2 := TreeIterNewFromNative(object2.GetObject().Native())

		handler(argInstance, arg1, arg2)
	}

	return callback.ConnectSignal(recv.Native(), "row-inserted", marshal)
}

// UNSUPPORTED : C value 'rows-reordered' : parameter 'new_order' of type 'gpointer' not supported

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *TreeModel) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}

var treeSortableInterface *gi.Interface
var treeSortableInterface_Once sync.Once

func treeSortableInterface_Set() error {
	var err error
	treeSortableInterface_Once.Do(func() {
		treeSortableInterface, err = gi.InterfaceNew("Gtk", "TreeSortable")
	})
	return err
}

type TreeSortable struct {
	native unsafe.Pointer
}

func TreeSortableNewFromNative(native unsafe.Pointer) *TreeSortable {
	instance := &TreeSortable{native: native}

	return instance
}

/*
CastToTreeSortable down casts any arbitrary Object to TreeSortable.
Exercise care, as this is a potentially dangerous function
if the Object is not a TreeSortable.
*/
func CastToTreeSortable(object *gobject.Object) *TreeSortable {
	return TreeSortableNewFromNative(object.Native())
}

// Equals compares this TreeSortable with another TreeSortable, and returns true if they represent the same Object.
func (recv *TreeSortable) Equals(other *TreeSortable) bool {
	return other.Native() == recv.Native()
}

func (recv *TreeSortable) Native() unsafe.Pointer {
	return recv.native
}

var treeSortableGetSortColumnIdFunction *gi.Function
var treeSortableGetSortColumnIdFunction_Once sync.Once

func treeSortableGetSortColumnIdFunction_Set() error {
	var err error
	treeSortableGetSortColumnIdFunction_Once.Do(func() {
		err = treeSortableInterface_Set()
		if err != nil {
			return
		}
		treeSortableGetSortColumnIdFunction, err = treeSortableInterface.InvokerNew("get_sort_column_id")
	})
	return err
}

// GetSortColumnId is a representation of the C type gtk_tree_sortable_get_sort_column_id.
func (recv *TreeSortable) GetSortColumnId() (bool, int32, SortType) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := treeSortableGetSortColumnIdFunction_Set()
	if err == nil {
		ret = treeSortableGetSortColumnIdFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := SortType(outArgs[1].Int32())

	return retGo, out0, out1
}

var treeSortableHasDefaultSortFuncFunction *gi.Function
var treeSortableHasDefaultSortFuncFunction_Once sync.Once

func treeSortableHasDefaultSortFuncFunction_Set() error {
	var err error
	treeSortableHasDefaultSortFuncFunction_Once.Do(func() {
		err = treeSortableInterface_Set()
		if err != nil {
			return
		}
		treeSortableHasDefaultSortFuncFunction, err = treeSortableInterface.InvokerNew("has_default_sort_func")
	})
	return err
}

// HasDefaultSortFunc is a representation of the C type gtk_tree_sortable_has_default_sort_func.
func (recv *TreeSortable) HasDefaultSortFunc() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := treeSortableHasDefaultSortFuncFunction_Set()
	if err == nil {
		ret = treeSortableHasDefaultSortFuncFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gtk_tree_sortable_set_default_sort_func' : parameter 'sort_func' of type 'TreeIterCompareFunc' not supported

var treeSortableSetSortColumnIdFunction *gi.Function
var treeSortableSetSortColumnIdFunction_Once sync.Once

func treeSortableSetSortColumnIdFunction_Set() error {
	var err error
	treeSortableSetSortColumnIdFunction_Once.Do(func() {
		err = treeSortableInterface_Set()
		if err != nil {
			return
		}
		treeSortableSetSortColumnIdFunction, err = treeSortableInterface.InvokerNew("set_sort_column_id")
	})
	return err
}

// SetSortColumnId is a representation of the C type gtk_tree_sortable_set_sort_column_id.
func (recv *TreeSortable) SetSortColumnId(sortColumnId int32, order SortType) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(sortColumnId)
	inArgs[2].SetInt32(int32(order))

	err := treeSortableSetSortColumnIdFunction_Set()
	if err == nil {
		treeSortableSetSortColumnIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gtk_tree_sortable_set_sort_func' : parameter 'sort_func' of type 'TreeIterCompareFunc' not supported

var treeSortableSortColumnChangedFunction *gi.Function
var treeSortableSortColumnChangedFunction_Once sync.Once

func treeSortableSortColumnChangedFunction_Set() error {
	var err error
	treeSortableSortColumnChangedFunction_Once.Do(func() {
		err = treeSortableInterface_Set()
		if err != nil {
			return
		}
		treeSortableSortColumnChangedFunction, err = treeSortableInterface.InvokerNew("sort_column_changed")
	})
	return err
}

// SortColumnChanged is a representation of the C type gtk_tree_sortable_sort_column_changed.
func (recv *TreeSortable) SortColumnChanged() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := treeSortableSortColumnChangedFunction_Set()
	if err == nil {
		treeSortableSortColumnChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

/*
ConnectSortColumnChanged connects a callback to the 'sort-column-changed' signal of the TreeSortable.

The returned value represents the connection, and may be passed to the Disconnect method to remove it.
*/
func (recv *TreeSortable) ConnectSortColumnChanged(handler func(instance *TreeSortable)) int {
	marshal := func(returnValue *callback.Value, paramValues []callback.Value) {
		objectInstance := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
		argInstance := TreeSortableNewFromNative(objectInstance.GetObject().Native())

		handler(argInstance)
	}

	return callback.ConnectSignal(recv.Native(), "sort-column-changed", marshal)
}

/*
Disconnect disconnects a callback previously registered with a Connect...() method.

The connectionID should be a value returned from a call to a Connect...() method.
*/
func (recv *TreeSortable) DisconnectSignal(connectionID int) {
	callback.DisconnectSignal(connectionID)
}
