// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

var devicePadInterface *gi.Interface
var devicePadInterface_Once sync.Once

func devicePadInterface_Set() error {
	var err error
	devicePadInterface_Once.Do(func() {
		devicePadInterface, err = gi.InterfaceNew("Gdk", "DevicePad")
	})
	return err
}

type DevicePad struct {
	native unsafe.Pointer
}

func DevicePadNewFromNative(native unsafe.Pointer) *DevicePad {
	instance := &DevicePad{native: native}

	return instance
}

/*
CastToDevicePad down casts any arbitrary Object to DevicePad.
Exercise care, as this is a potentially dangerous function
if the Object is not a DevicePad.
*/
func CastToDevicePad(object *gobject.Object) *DevicePad {
	return DevicePadNewFromNative(object.Native())
}

// Equals compares this DevicePad with another DevicePad, and returns true if they represent the same GObject.
func (recv *DevicePad) Equals(other *DevicePad) bool {
	return other.Native() == recv.Native()
}

func (recv *DevicePad) Native() unsafe.Pointer {
	return recv.native
}

var devicePadGetFeatureGroupFunction *gi.Function
var devicePadGetFeatureGroupFunction_Once sync.Once

func devicePadGetFeatureGroupFunction_Set() error {
	var err error
	devicePadGetFeatureGroupFunction_Once.Do(func() {
		err = devicePadInterface_Set()
		if err != nil {
			return
		}
		devicePadGetFeatureGroupFunction, err = devicePadInterface.InvokerNew("get_feature_group")
	})
	return err
}

// GetFeatureGroup is a representation of the C type gdk_device_pad_get_feature_group.
func (recv *DevicePad) GetFeatureGroup(feature DevicePadFeature, featureIdx int32) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(feature))
	inArgs[2].SetInt32(featureIdx)

	var ret gi.Argument

	err := devicePadGetFeatureGroupFunction_Set()
	if err == nil {
		ret = devicePadGetFeatureGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var devicePadGetGroupNModesFunction *gi.Function
var devicePadGetGroupNModesFunction_Once sync.Once

func devicePadGetGroupNModesFunction_Set() error {
	var err error
	devicePadGetGroupNModesFunction_Once.Do(func() {
		err = devicePadInterface_Set()
		if err != nil {
			return
		}
		devicePadGetGroupNModesFunction, err = devicePadInterface.InvokerNew("get_group_n_modes")
	})
	return err
}

// GetGroupNModes is a representation of the C type gdk_device_pad_get_group_n_modes.
func (recv *DevicePad) GetGroupNModes(groupIdx int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(groupIdx)

	var ret gi.Argument

	err := devicePadGetGroupNModesFunction_Set()
	if err == nil {
		ret = devicePadGetGroupNModesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var devicePadGetNFeaturesFunction *gi.Function
var devicePadGetNFeaturesFunction_Once sync.Once

func devicePadGetNFeaturesFunction_Set() error {
	var err error
	devicePadGetNFeaturesFunction_Once.Do(func() {
		err = devicePadInterface_Set()
		if err != nil {
			return
		}
		devicePadGetNFeaturesFunction, err = devicePadInterface.InvokerNew("get_n_features")
	})
	return err
}

// GetNFeatures is a representation of the C type gdk_device_pad_get_n_features.
func (recv *DevicePad) GetNFeatures(feature DevicePadFeature) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native())
	inArgs[1].SetInt32(int32(feature))

	var ret gi.Argument

	err := devicePadGetNFeaturesFunction_Set()
	if err == nil {
		ret = devicePadGetNFeaturesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var devicePadGetNGroupsFunction *gi.Function
var devicePadGetNGroupsFunction_Once sync.Once

func devicePadGetNGroupsFunction_Set() error {
	var err error
	devicePadGetNGroupsFunction_Once.Do(func() {
		err = devicePadInterface_Set()
		if err != nil {
			return
		}
		devicePadGetNGroupsFunction, err = devicePadInterface.InvokerNew("get_n_groups")
	})
	return err
}

// GetNGroups is a representation of the C type gdk_device_pad_get_n_groups.
func (recv *DevicePad) GetNGroups() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	var ret gi.Argument

	err := devicePadGetNGroupsFunction_Set()
	if err == nil {
		ret = devicePadGetNGroupsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}
