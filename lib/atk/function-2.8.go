// This is a generated file - DO NOT EDIT
// +build atk_2.8 atk_2.12

package atk

import call "github.com/pekim/gobbi/lib/internal/call"

// GetBinaryAge is a wrapper around the C function atk_get_binary_age.
func GetBinaryAge() uint32 {
	data := call.Data{Return: call.Value{Type: call.TYPE_UINT}}
	call.Function(39, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetInterfaceAge is a wrapper around the C function atk_get_interface_age.
func GetInterfaceAge() uint32 {
	data := call.Data{Return: call.Value{Type: call.TYPE_UINT}}
	call.Function(42, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetMajorVersion is a wrapper around the C function atk_get_major_version.
func GetMajorVersion() uint32 {
	data := call.Data{Return: call.Value{Type: call.TYPE_UINT}}
	call.Function(43, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetMicroVersion is a wrapper around the C function atk_get_micro_version.
func GetMicroVersion() uint32 {
	data := call.Data{Return: call.Value{Type: call.TYPE_UINT}}
	call.Function(44, &data)
	ret := data.Return.Uint32()

	return ret
}

// GetMinorVersion is a wrapper around the C function atk_get_minor_version.
func GetMinorVersion() uint32 {
	data := call.Data{Return: call.Value{Type: call.TYPE_UINT}}
	call.Function(45, &data)
	ret := data.Return.Uint32()

	return ret
}
