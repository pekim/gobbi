// Code generated - DO NOT EDIT.

package webkit2

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

var permissionRequestInterface *gi.Interface
var permissionRequestInterface_Once sync.Once

func permissionRequestInterface_Set() error {
	var err error
	permissionRequestInterface_Once.Do(func() {
		permissionRequestInterface, err = gi.InterfaceNew("WebKit2", "PermissionRequest")
	})
	return err
}

type PermissionRequest struct {
	native unsafe.Pointer
}

func PermissionRequestNewFromNative(native unsafe.Pointer) *PermissionRequest {
	instance := &PermissionRequest{native: native}

	return instance
}

/*
CastToPermissionRequest down casts any arbitrary Object to PermissionRequest.
Exercise care, as this is a potentially dangerous function
if the Object is not a PermissionRequest.
*/
func CastToPermissionRequest(object *gobject.Object) *PermissionRequest {
	return PermissionRequestNewFromNative(object.Native())
}

// Equals compares this PermissionRequest with another PermissionRequest, and returns true if they represent the same Object.
func (recv *PermissionRequest) Equals(other *PermissionRequest) bool {
	return other.Native() == recv.Native()
}

func (recv *PermissionRequest) Native() unsafe.Pointer {
	return recv.native
}

var permissionRequestAllowFunction *gi.Function
var permissionRequestAllowFunction_Once sync.Once

func permissionRequestAllowFunction_Set() error {
	var err error
	permissionRequestAllowFunction_Once.Do(func() {
		err = permissionRequestInterface_Set()
		if err != nil {
			return
		}
		permissionRequestAllowFunction, err = permissionRequestInterface.InvokerNew("allow")
	})
	return err
}

// Allow is a representation of the C type webkit_permission_request_allow.
func (recv *PermissionRequest) Allow() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := permissionRequestAllowFunction_Set()
	if err == nil {
		permissionRequestAllowFunction.Invoke(inArgs[:], nil)
	}

	return
}

var permissionRequestDenyFunction *gi.Function
var permissionRequestDenyFunction_Once sync.Once

func permissionRequestDenyFunction_Set() error {
	var err error
	permissionRequestDenyFunction_Once.Do(func() {
		err = permissionRequestInterface_Set()
		if err != nil {
			return
		}
		permissionRequestDenyFunction, err = permissionRequestInterface.InvokerNew("deny")
	})
	return err
}

// Deny is a representation of the C type webkit_permission_request_deny.
func (recv *PermissionRequest) Deny() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native())

	err := permissionRequestDenyFunction_Set()
	if err == nil {
		permissionRequestDenyFunction.Invoke(inArgs[:], nil)
	}

	return
}
