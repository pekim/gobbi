// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// DevicePad is a wrapper around the C record GdkDevicePad.
type DevicePad struct {
	native *C.GdkDevicePad
}

func DevicePadNewFromC(u unsafe.Pointer) *DevicePad {
	c := (*C.GdkDevicePad)(u)
	if c == nil {
		return nil
	}

	g := &DevicePad{native: c}

	return g
}

func (recv *DevicePad) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Returns the group the given @feature and @idx belong to,
// or -1 if feature/index do not exist in @pad.
/*

C function : gdk_device_pad_get_feature_group
*/
func (recv *DevicePad) GetFeatureGroup(feature DevicePadFeature, featureIdx int32) int32 {
	c_feature := (C.GdkDevicePadFeature)(feature)

	c_feature_idx := (C.gint)(featureIdx)

	retC := C.gdk_device_pad_get_feature_group((*C.GdkDevicePad)(recv.native), c_feature, c_feature_idx)
	retGo := (int32)(retC)

	return retGo
}

// Returns the number of modes that @group may have.
/*

C function : gdk_device_pad_get_group_n_modes
*/
func (recv *DevicePad) GetGroupNModes(groupIdx int32) int32 {
	c_group_idx := (C.gint)(groupIdx)

	retC := C.gdk_device_pad_get_group_n_modes((*C.GdkDevicePad)(recv.native), c_group_idx)
	retGo := (int32)(retC)

	return retGo
}

// Returns the number of features a tablet pad has.
/*

C function : gdk_device_pad_get_n_features
*/
func (recv *DevicePad) GetNFeatures(feature DevicePadFeature) int32 {
	c_feature := (C.GdkDevicePadFeature)(feature)

	retC := C.gdk_device_pad_get_n_features((*C.GdkDevicePad)(recv.native), c_feature)
	retGo := (int32)(retC)

	return retGo
}

// Returns the number of groups this pad device has. Pads have
// at least one group. A pad group is a subcollection of
// buttons/strip/rings that is affected collectively by a same
// current mode.
/*

C function : gdk_device_pad_get_n_groups
*/
func (recv *DevicePad) GetNGroups() int32 {
	retC := C.gdk_device_pad_get_n_groups((*C.GdkDevicePad)(recv.native))
	retGo := (int32)(retC)

	return retGo
}
