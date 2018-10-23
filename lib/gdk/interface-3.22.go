// This is a generated file - DO NOT EDIT
// +build gdk_3.22

package gdk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <stdlib.h>
import "C"

// GetFeatureGroup is a wrapper around the C function gdk_device_pad_get_feature_group.
func (recv *DevicePad) GetFeatureGroup(feature DevicePadFeature, featureIdx int32) int32 {
	c_feature := (C.GdkDevicePadFeature)(feature)

	c_feature_idx := (C.gint)(featureIdx)

	retC := C.gdk_device_pad_get_feature_group((*C.GdkDevicePad)(recv.native), c_feature, c_feature_idx)
	retGo := (int32)(retC)

	return retGo
}

// GetGroupNModes is a wrapper around the C function gdk_device_pad_get_group_n_modes.
func (recv *DevicePad) GetGroupNModes(groupIdx int32) int32 {
	c_group_idx := (C.gint)(groupIdx)

	retC := C.gdk_device_pad_get_group_n_modes((*C.GdkDevicePad)(recv.native), c_group_idx)
	retGo := (int32)(retC)

	return retGo
}

// GetNFeatures is a wrapper around the C function gdk_device_pad_get_n_features.
func (recv *DevicePad) GetNFeatures(feature DevicePadFeature) int32 {
	c_feature := (C.GdkDevicePadFeature)(feature)

	retC := C.gdk_device_pad_get_n_features((*C.GdkDevicePad)(recv.native), c_feature)
	retGo := (int32)(retC)

	return retGo
}

// GetNGroups is a wrapper around the C function gdk_device_pad_get_n_groups.
func (recv *DevicePad) GetNGroups() int32 {
	retC := C.gdk_device_pad_get_n_groups((*C.GdkDevicePad)(recv.native))
	retGo := (int32)(retC)

	return retGo
}
