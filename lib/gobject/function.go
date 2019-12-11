// Code generated - DO NOT EDIT.

package gobject

import (
	gi "github.com/pekim/gobbi/internal/cgo/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

var boxedCopyFunction *gi.Function
var boxedCopyFunction_Once sync.Once

func boxedCopyFunction_Set() error {
	var err error
	boxedCopyFunction_Once.Do(func() {
		boxedCopyFunction, err = gi.FunctionInvokerNew("GObject", "boxed_copy")
	})
	return err
}

// BoxedCopy is a representation of the C type g_boxed_copy.
func BoxedCopy(boxedType int64, srcBoxed unsafe.Pointer) unsafe.Pointer {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(boxedType)
	inArgs[1].SetPointer(srcBoxed)

	var ret gi.Argument

	err := boxedCopyFunction_Set()
	if err == nil {
		ret = boxedCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Pointer()

	return retGo
}

var boxedFreeFunction *gi.Function
var boxedFreeFunction_Once sync.Once

func boxedFreeFunction_Set() error {
	var err error
	boxedFreeFunction_Once.Do(func() {
		boxedFreeFunction, err = gi.FunctionInvokerNew("GObject", "boxed_free")
	})
	return err
}

// BoxedFree is a representation of the C type g_boxed_free.
func BoxedFree(boxedType int64, boxed unsafe.Pointer) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(boxedType)
	inArgs[1].SetPointer(boxed)

	err := boxedFreeFunction_Set()
	if err == nil {
		boxedFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_boxed_type_register_static' : parameter 'boxed_copy' of type 'BoxedCopyFunc' not supported

var cclosureMarshalBOOLEANBOXEDBOXEDFunction *gi.Function
var cclosureMarshalBOOLEANBOXEDBOXEDFunction_Once sync.Once

func cclosureMarshalBOOLEANBOXEDBOXEDFunction_Set() error {
	var err error
	cclosureMarshalBOOLEANBOXEDBOXEDFunction_Once.Do(func() {
		cclosureMarshalBOOLEANBOXEDBOXEDFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_BOOLEAN__BOXED_BOXED")
	})
	return err
}

// CclosureMarshalBOOLEANBOXEDBOXED is a representation of the C type g_cclosure_marshal_BOOLEAN__BOXED_BOXED.
func CclosureMarshalBOOLEANBOXEDBOXED(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalBOOLEANBOXEDBOXEDFunction_Set()
	if err == nil {
		cclosureMarshalBOOLEANBOXEDBOXEDFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalBOOLEANFLAGSFunction *gi.Function
var cclosureMarshalBOOLEANFLAGSFunction_Once sync.Once

func cclosureMarshalBOOLEANFLAGSFunction_Set() error {
	var err error
	cclosureMarshalBOOLEANFLAGSFunction_Once.Do(func() {
		cclosureMarshalBOOLEANFLAGSFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_BOOLEAN__FLAGS")
	})
	return err
}

// CclosureMarshalBOOLEANFLAGS is a representation of the C type g_cclosure_marshal_BOOLEAN__FLAGS.
func CclosureMarshalBOOLEANFLAGS(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalBOOLEANFLAGSFunction_Set()
	if err == nil {
		cclosureMarshalBOOLEANFLAGSFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalSTRINGOBJECTPOINTERFunction *gi.Function
var cclosureMarshalSTRINGOBJECTPOINTERFunction_Once sync.Once

func cclosureMarshalSTRINGOBJECTPOINTERFunction_Set() error {
	var err error
	cclosureMarshalSTRINGOBJECTPOINTERFunction_Once.Do(func() {
		cclosureMarshalSTRINGOBJECTPOINTERFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_STRING__OBJECT_POINTER")
	})
	return err
}

// CclosureMarshalSTRINGOBJECTPOINTER is a representation of the C type g_cclosure_marshal_STRING__OBJECT_POINTER.
func CclosureMarshalSTRINGOBJECTPOINTER(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalSTRINGOBJECTPOINTERFunction_Set()
	if err == nil {
		cclosureMarshalSTRINGOBJECTPOINTERFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDBOOLEANFunction *gi.Function
var cclosureMarshalVOIDBOOLEANFunction_Once sync.Once

func cclosureMarshalVOIDBOOLEANFunction_Set() error {
	var err error
	cclosureMarshalVOIDBOOLEANFunction_Once.Do(func() {
		cclosureMarshalVOIDBOOLEANFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__BOOLEAN")
	})
	return err
}

// CclosureMarshalVOIDBOOLEAN is a representation of the C type g_cclosure_marshal_VOID__BOOLEAN.
func CclosureMarshalVOIDBOOLEAN(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDBOOLEANFunction_Set()
	if err == nil {
		cclosureMarshalVOIDBOOLEANFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDBOXEDFunction *gi.Function
var cclosureMarshalVOIDBOXEDFunction_Once sync.Once

func cclosureMarshalVOIDBOXEDFunction_Set() error {
	var err error
	cclosureMarshalVOIDBOXEDFunction_Once.Do(func() {
		cclosureMarshalVOIDBOXEDFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__BOXED")
	})
	return err
}

// CclosureMarshalVOIDBOXED is a representation of the C type g_cclosure_marshal_VOID__BOXED.
func CclosureMarshalVOIDBOXED(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDBOXEDFunction_Set()
	if err == nil {
		cclosureMarshalVOIDBOXEDFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDCHARFunction *gi.Function
var cclosureMarshalVOIDCHARFunction_Once sync.Once

func cclosureMarshalVOIDCHARFunction_Set() error {
	var err error
	cclosureMarshalVOIDCHARFunction_Once.Do(func() {
		cclosureMarshalVOIDCHARFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__CHAR")
	})
	return err
}

// CclosureMarshalVOIDCHAR is a representation of the C type g_cclosure_marshal_VOID__CHAR.
func CclosureMarshalVOIDCHAR(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDCHARFunction_Set()
	if err == nil {
		cclosureMarshalVOIDCHARFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDDOUBLEFunction *gi.Function
var cclosureMarshalVOIDDOUBLEFunction_Once sync.Once

func cclosureMarshalVOIDDOUBLEFunction_Set() error {
	var err error
	cclosureMarshalVOIDDOUBLEFunction_Once.Do(func() {
		cclosureMarshalVOIDDOUBLEFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__DOUBLE")
	})
	return err
}

// CclosureMarshalVOIDDOUBLE is a representation of the C type g_cclosure_marshal_VOID__DOUBLE.
func CclosureMarshalVOIDDOUBLE(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDDOUBLEFunction_Set()
	if err == nil {
		cclosureMarshalVOIDDOUBLEFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDENUMFunction *gi.Function
var cclosureMarshalVOIDENUMFunction_Once sync.Once

func cclosureMarshalVOIDENUMFunction_Set() error {
	var err error
	cclosureMarshalVOIDENUMFunction_Once.Do(func() {
		cclosureMarshalVOIDENUMFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__ENUM")
	})
	return err
}

// CclosureMarshalVOIDENUM is a representation of the C type g_cclosure_marshal_VOID__ENUM.
func CclosureMarshalVOIDENUM(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDENUMFunction_Set()
	if err == nil {
		cclosureMarshalVOIDENUMFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDFLAGSFunction *gi.Function
var cclosureMarshalVOIDFLAGSFunction_Once sync.Once

func cclosureMarshalVOIDFLAGSFunction_Set() error {
	var err error
	cclosureMarshalVOIDFLAGSFunction_Once.Do(func() {
		cclosureMarshalVOIDFLAGSFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__FLAGS")
	})
	return err
}

// CclosureMarshalVOIDFLAGS is a representation of the C type g_cclosure_marshal_VOID__FLAGS.
func CclosureMarshalVOIDFLAGS(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDFLAGSFunction_Set()
	if err == nil {
		cclosureMarshalVOIDFLAGSFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDFLOATFunction *gi.Function
var cclosureMarshalVOIDFLOATFunction_Once sync.Once

func cclosureMarshalVOIDFLOATFunction_Set() error {
	var err error
	cclosureMarshalVOIDFLOATFunction_Once.Do(func() {
		cclosureMarshalVOIDFLOATFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__FLOAT")
	})
	return err
}

// CclosureMarshalVOIDFLOAT is a representation of the C type g_cclosure_marshal_VOID__FLOAT.
func CclosureMarshalVOIDFLOAT(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDFLOATFunction_Set()
	if err == nil {
		cclosureMarshalVOIDFLOATFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDINTFunction *gi.Function
var cclosureMarshalVOIDINTFunction_Once sync.Once

func cclosureMarshalVOIDINTFunction_Set() error {
	var err error
	cclosureMarshalVOIDINTFunction_Once.Do(func() {
		cclosureMarshalVOIDINTFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__INT")
	})
	return err
}

// CclosureMarshalVOIDINT is a representation of the C type g_cclosure_marshal_VOID__INT.
func CclosureMarshalVOIDINT(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDINTFunction_Set()
	if err == nil {
		cclosureMarshalVOIDINTFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDLONGFunction *gi.Function
var cclosureMarshalVOIDLONGFunction_Once sync.Once

func cclosureMarshalVOIDLONGFunction_Set() error {
	var err error
	cclosureMarshalVOIDLONGFunction_Once.Do(func() {
		cclosureMarshalVOIDLONGFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__LONG")
	})
	return err
}

// CclosureMarshalVOIDLONG is a representation of the C type g_cclosure_marshal_VOID__LONG.
func CclosureMarshalVOIDLONG(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDLONGFunction_Set()
	if err == nil {
		cclosureMarshalVOIDLONGFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDOBJECTFunction *gi.Function
var cclosureMarshalVOIDOBJECTFunction_Once sync.Once

func cclosureMarshalVOIDOBJECTFunction_Set() error {
	var err error
	cclosureMarshalVOIDOBJECTFunction_Once.Do(func() {
		cclosureMarshalVOIDOBJECTFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__OBJECT")
	})
	return err
}

// CclosureMarshalVOIDOBJECT is a representation of the C type g_cclosure_marshal_VOID__OBJECT.
func CclosureMarshalVOIDOBJECT(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDOBJECTFunction_Set()
	if err == nil {
		cclosureMarshalVOIDOBJECTFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDPARAMFunction *gi.Function
var cclosureMarshalVOIDPARAMFunction_Once sync.Once

func cclosureMarshalVOIDPARAMFunction_Set() error {
	var err error
	cclosureMarshalVOIDPARAMFunction_Once.Do(func() {
		cclosureMarshalVOIDPARAMFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__PARAM")
	})
	return err
}

// CclosureMarshalVOIDPARAM is a representation of the C type g_cclosure_marshal_VOID__PARAM.
func CclosureMarshalVOIDPARAM(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDPARAMFunction_Set()
	if err == nil {
		cclosureMarshalVOIDPARAMFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDPOINTERFunction *gi.Function
var cclosureMarshalVOIDPOINTERFunction_Once sync.Once

func cclosureMarshalVOIDPOINTERFunction_Set() error {
	var err error
	cclosureMarshalVOIDPOINTERFunction_Once.Do(func() {
		cclosureMarshalVOIDPOINTERFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__POINTER")
	})
	return err
}

// CclosureMarshalVOIDPOINTER is a representation of the C type g_cclosure_marshal_VOID__POINTER.
func CclosureMarshalVOIDPOINTER(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDPOINTERFunction_Set()
	if err == nil {
		cclosureMarshalVOIDPOINTERFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDSTRINGFunction *gi.Function
var cclosureMarshalVOIDSTRINGFunction_Once sync.Once

func cclosureMarshalVOIDSTRINGFunction_Set() error {
	var err error
	cclosureMarshalVOIDSTRINGFunction_Once.Do(func() {
		cclosureMarshalVOIDSTRINGFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__STRING")
	})
	return err
}

// CclosureMarshalVOIDSTRING is a representation of the C type g_cclosure_marshal_VOID__STRING.
func CclosureMarshalVOIDSTRING(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDSTRINGFunction_Set()
	if err == nil {
		cclosureMarshalVOIDSTRINGFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDUCHARFunction *gi.Function
var cclosureMarshalVOIDUCHARFunction_Once sync.Once

func cclosureMarshalVOIDUCHARFunction_Set() error {
	var err error
	cclosureMarshalVOIDUCHARFunction_Once.Do(func() {
		cclosureMarshalVOIDUCHARFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__UCHAR")
	})
	return err
}

// CclosureMarshalVOIDUCHAR is a representation of the C type g_cclosure_marshal_VOID__UCHAR.
func CclosureMarshalVOIDUCHAR(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDUCHARFunction_Set()
	if err == nil {
		cclosureMarshalVOIDUCHARFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDUINTFunction *gi.Function
var cclosureMarshalVOIDUINTFunction_Once sync.Once

func cclosureMarshalVOIDUINTFunction_Set() error {
	var err error
	cclosureMarshalVOIDUINTFunction_Once.Do(func() {
		cclosureMarshalVOIDUINTFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__UINT")
	})
	return err
}

// CclosureMarshalVOIDUINT is a representation of the C type g_cclosure_marshal_VOID__UINT.
func CclosureMarshalVOIDUINT(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDUINTFunction_Set()
	if err == nil {
		cclosureMarshalVOIDUINTFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDUINTPOINTERFunction *gi.Function
var cclosureMarshalVOIDUINTPOINTERFunction_Once sync.Once

func cclosureMarshalVOIDUINTPOINTERFunction_Set() error {
	var err error
	cclosureMarshalVOIDUINTPOINTERFunction_Once.Do(func() {
		cclosureMarshalVOIDUINTPOINTERFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__UINT_POINTER")
	})
	return err
}

// CclosureMarshalVOIDUINTPOINTER is a representation of the C type g_cclosure_marshal_VOID__UINT_POINTER.
func CclosureMarshalVOIDUINTPOINTER(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDUINTPOINTERFunction_Set()
	if err == nil {
		cclosureMarshalVOIDUINTPOINTERFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDULONGFunction *gi.Function
var cclosureMarshalVOIDULONGFunction_Once sync.Once

func cclosureMarshalVOIDULONGFunction_Set() error {
	var err error
	cclosureMarshalVOIDULONGFunction_Once.Do(func() {
		cclosureMarshalVOIDULONGFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__ULONG")
	})
	return err
}

// CclosureMarshalVOIDULONG is a representation of the C type g_cclosure_marshal_VOID__ULONG.
func CclosureMarshalVOIDULONG(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDULONGFunction_Set()
	if err == nil {
		cclosureMarshalVOIDULONGFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDVARIANTFunction *gi.Function
var cclosureMarshalVOIDVARIANTFunction_Once sync.Once

func cclosureMarshalVOIDVARIANTFunction_Set() error {
	var err error
	cclosureMarshalVOIDVARIANTFunction_Once.Do(func() {
		cclosureMarshalVOIDVARIANTFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__VARIANT")
	})
	return err
}

// CclosureMarshalVOIDVARIANT is a representation of the C type g_cclosure_marshal_VOID__VARIANT.
func CclosureMarshalVOIDVARIANT(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDVARIANTFunction_Set()
	if err == nil {
		cclosureMarshalVOIDVARIANTFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalVOIDVOIDFunction *gi.Function
var cclosureMarshalVOIDVOIDFunction_Once sync.Once

func cclosureMarshalVOIDVOIDFunction_Set() error {
	var err error
	cclosureMarshalVOIDVOIDFunction_Once.Do(func() {
		cclosureMarshalVOIDVOIDFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_VOID__VOID")
	})
	return err
}

// CclosureMarshalVOIDVOID is a representation of the C type g_cclosure_marshal_VOID__VOID.
func CclosureMarshalVOIDVOID(closure *Closure, returnValue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnValue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalVOIDVOIDFunction_Set()
	if err == nil {
		cclosureMarshalVOIDVOIDFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cclosureMarshalGenericFunction *gi.Function
var cclosureMarshalGenericFunction_Once sync.Once

func cclosureMarshalGenericFunction_Set() error {
	var err error
	cclosureMarshalGenericFunction_Once.Do(func() {
		cclosureMarshalGenericFunction, err = gi.FunctionInvokerNew("GObject", "cclosure_marshal_generic")
	})
	return err
}

// CclosureMarshalGeneric is a representation of the C type g_cclosure_marshal_generic.
func CclosureMarshalGeneric(closure *Closure, returnGvalue *Value, nParamValues uint32, paramValues *Value, invocationHint unsafe.Pointer, marshalData unsafe.Pointer) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(closure.Native())
	inArgs[1].SetPointer(returnGvalue.Native())
	inArgs[2].SetUint32(nParamValues)
	inArgs[3].SetPointer(paramValues.Native())
	inArgs[4].SetPointer(invocationHint)
	inArgs[5].SetPointer(marshalData)

	err := cclosureMarshalGenericFunction_Set()
	if err == nil {
		cclosureMarshalGenericFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_cclosure_new' : parameter 'callback_func' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_cclosure_new_object' : parameter 'callback_func' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_cclosure_new_object_swap' : parameter 'callback_func' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_cclosure_new_swap' : parameter 'callback_func' of type 'Callback' not supported

var clearObjectFunction *gi.Function
var clearObjectFunction_Once sync.Once

func clearObjectFunction_Set() error {
	var err error
	clearObjectFunction_Once.Do(func() {
		clearObjectFunction, err = gi.FunctionInvokerNew("GObject", "clear_object")
	})
	return err
}

// ClearObject is a representation of the C type g_clear_object.
func ClearObject(objectPtr *Object) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(objectPtr.Native())

	err := clearObjectFunction_Set()
	if err == nil {
		clearObjectFunction.Invoke(inArgs[:], nil)
	}

	return
}

var clearSignalHandlerFunction *gi.Function
var clearSignalHandlerFunction_Once sync.Once

func clearSignalHandlerFunction_Set() error {
	var err error
	clearSignalHandlerFunction_Once.Do(func() {
		clearSignalHandlerFunction, err = gi.FunctionInvokerNew("GObject", "clear_signal_handler")
	})
	return err
}

// ClearSignalHandler is a representation of the C type g_clear_signal_handler.
func ClearSignalHandler(handlerIdPtr uint64, instance *Object) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint64(handlerIdPtr)
	inArgs[1].SetPointer(instance.Native())

	err := clearSignalHandlerFunction_Set()
	if err == nil {
		clearSignalHandlerFunction.Invoke(inArgs[:], nil)
	}

	return
}

var enumCompleteTypeInfoFunction *gi.Function
var enumCompleteTypeInfoFunction_Once sync.Once

func enumCompleteTypeInfoFunction_Set() error {
	var err error
	enumCompleteTypeInfoFunction_Once.Do(func() {
		enumCompleteTypeInfoFunction, err = gi.FunctionInvokerNew("GObject", "enum_complete_type_info")
	})
	return err
}

// EnumCompleteTypeInfo is a representation of the C type g_enum_complete_type_info.
func EnumCompleteTypeInfo(gEnumType int64, constValues *EnumValue) *TypeInfo {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(gEnumType)
	inArgs[1].SetPointer(constValues.Native())

	var outArgs [1]gi.Argument

	err := enumCompleteTypeInfoFunction_Set()
	if err == nil {
		enumCompleteTypeInfoFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := TypeInfoNewFromNative(outArgs[0].Pointer())

	return out0
}

var enumGetValueFunction *gi.Function
var enumGetValueFunction_Once sync.Once

func enumGetValueFunction_Set() error {
	var err error
	enumGetValueFunction_Once.Do(func() {
		enumGetValueFunction, err = gi.FunctionInvokerNew("GObject", "enum_get_value")
	})
	return err
}

// EnumGetValue is a representation of the C type g_enum_get_value.
func EnumGetValue(enumClass *EnumClass, value int32) *EnumValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(enumClass.Native())
	inArgs[1].SetInt32(value)

	var ret gi.Argument

	err := enumGetValueFunction_Set()
	if err == nil {
		ret = enumGetValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := EnumValueNewFromNative(ret.Pointer())

	return retGo
}

var enumGetValueByNameFunction *gi.Function
var enumGetValueByNameFunction_Once sync.Once

func enumGetValueByNameFunction_Set() error {
	var err error
	enumGetValueByNameFunction_Once.Do(func() {
		enumGetValueByNameFunction, err = gi.FunctionInvokerNew("GObject", "enum_get_value_by_name")
	})
	return err
}

// EnumGetValueByName is a representation of the C type g_enum_get_value_by_name.
func EnumGetValueByName(enumClass *EnumClass, name string) *EnumValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(enumClass.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := enumGetValueByNameFunction_Set()
	if err == nil {
		ret = enumGetValueByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := EnumValueNewFromNative(ret.Pointer())

	return retGo
}

var enumGetValueByNickFunction *gi.Function
var enumGetValueByNickFunction_Once sync.Once

func enumGetValueByNickFunction_Set() error {
	var err error
	enumGetValueByNickFunction_Once.Do(func() {
		enumGetValueByNickFunction, err = gi.FunctionInvokerNew("GObject", "enum_get_value_by_nick")
	})
	return err
}

// EnumGetValueByNick is a representation of the C type g_enum_get_value_by_nick.
func EnumGetValueByNick(enumClass *EnumClass, nick string) *EnumValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(enumClass.Native())
	inArgs[1].SetString(nick)

	var ret gi.Argument

	err := enumGetValueByNickFunction_Set()
	if err == nil {
		ret = enumGetValueByNickFunction.Invoke(inArgs[:], nil)
	}

	retGo := EnumValueNewFromNative(ret.Pointer())

	return retGo
}

var enumRegisterStaticFunction *gi.Function
var enumRegisterStaticFunction_Once sync.Once

func enumRegisterStaticFunction_Set() error {
	var err error
	enumRegisterStaticFunction_Once.Do(func() {
		enumRegisterStaticFunction, err = gi.FunctionInvokerNew("GObject", "enum_register_static")
	})
	return err
}

// EnumRegisterStatic is a representation of the C type g_enum_register_static.
func EnumRegisterStatic(name string, constStaticValues *EnumValue) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetPointer(constStaticValues.Native())

	var ret gi.Argument

	err := enumRegisterStaticFunction_Set()
	if err == nil {
		ret = enumRegisterStaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var enumToStringFunction *gi.Function
var enumToStringFunction_Once sync.Once

func enumToStringFunction_Set() error {
	var err error
	enumToStringFunction_Once.Do(func() {
		enumToStringFunction, err = gi.FunctionInvokerNew("GObject", "enum_to_string")
	})
	return err
}

// EnumToString is a representation of the C type g_enum_to_string.
func EnumToString(gEnumType int64, value int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(gEnumType)
	inArgs[1].SetInt32(value)

	var ret gi.Argument

	err := enumToStringFunction_Set()
	if err == nil {
		ret = enumToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var flagsCompleteTypeInfoFunction *gi.Function
var flagsCompleteTypeInfoFunction_Once sync.Once

func flagsCompleteTypeInfoFunction_Set() error {
	var err error
	flagsCompleteTypeInfoFunction_Once.Do(func() {
		flagsCompleteTypeInfoFunction, err = gi.FunctionInvokerNew("GObject", "flags_complete_type_info")
	})
	return err
}

// FlagsCompleteTypeInfo is a representation of the C type g_flags_complete_type_info.
func FlagsCompleteTypeInfo(gFlagsType int64, constValues *FlagsValue) *TypeInfo {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(gFlagsType)
	inArgs[1].SetPointer(constValues.Native())

	var outArgs [1]gi.Argument

	err := flagsCompleteTypeInfoFunction_Set()
	if err == nil {
		flagsCompleteTypeInfoFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := TypeInfoNewFromNative(outArgs[0].Pointer())

	return out0
}

var flagsGetFirstValueFunction *gi.Function
var flagsGetFirstValueFunction_Once sync.Once

func flagsGetFirstValueFunction_Set() error {
	var err error
	flagsGetFirstValueFunction_Once.Do(func() {
		flagsGetFirstValueFunction, err = gi.FunctionInvokerNew("GObject", "flags_get_first_value")
	})
	return err
}

// FlagsGetFirstValue is a representation of the C type g_flags_get_first_value.
func FlagsGetFirstValue(flagsClass *FlagsClass, value uint32) *FlagsValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(flagsClass.Native())
	inArgs[1].SetUint32(value)

	var ret gi.Argument

	err := flagsGetFirstValueFunction_Set()
	if err == nil {
		ret = flagsGetFirstValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := FlagsValueNewFromNative(ret.Pointer())

	return retGo
}

var flagsGetValueByNameFunction *gi.Function
var flagsGetValueByNameFunction_Once sync.Once

func flagsGetValueByNameFunction_Set() error {
	var err error
	flagsGetValueByNameFunction_Once.Do(func() {
		flagsGetValueByNameFunction, err = gi.FunctionInvokerNew("GObject", "flags_get_value_by_name")
	})
	return err
}

// FlagsGetValueByName is a representation of the C type g_flags_get_value_by_name.
func FlagsGetValueByName(flagsClass *FlagsClass, name string) *FlagsValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(flagsClass.Native())
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := flagsGetValueByNameFunction_Set()
	if err == nil {
		ret = flagsGetValueByNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := FlagsValueNewFromNative(ret.Pointer())

	return retGo
}

var flagsGetValueByNickFunction *gi.Function
var flagsGetValueByNickFunction_Once sync.Once

func flagsGetValueByNickFunction_Set() error {
	var err error
	flagsGetValueByNickFunction_Once.Do(func() {
		flagsGetValueByNickFunction, err = gi.FunctionInvokerNew("GObject", "flags_get_value_by_nick")
	})
	return err
}

// FlagsGetValueByNick is a representation of the C type g_flags_get_value_by_nick.
func FlagsGetValueByNick(flagsClass *FlagsClass, nick string) *FlagsValue {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(flagsClass.Native())
	inArgs[1].SetString(nick)

	var ret gi.Argument

	err := flagsGetValueByNickFunction_Set()
	if err == nil {
		ret = flagsGetValueByNickFunction.Invoke(inArgs[:], nil)
	}

	retGo := FlagsValueNewFromNative(ret.Pointer())

	return retGo
}

var flagsRegisterStaticFunction *gi.Function
var flagsRegisterStaticFunction_Once sync.Once

func flagsRegisterStaticFunction_Set() error {
	var err error
	flagsRegisterStaticFunction_Once.Do(func() {
		flagsRegisterStaticFunction, err = gi.FunctionInvokerNew("GObject", "flags_register_static")
	})
	return err
}

// FlagsRegisterStatic is a representation of the C type g_flags_register_static.
func FlagsRegisterStatic(name string, constStaticValues *FlagsValue) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetPointer(constStaticValues.Native())

	var ret gi.Argument

	err := flagsRegisterStaticFunction_Set()
	if err == nil {
		ret = flagsRegisterStaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var flagsToStringFunction *gi.Function
var flagsToStringFunction_Once sync.Once

func flagsToStringFunction_Set() error {
	var err error
	flagsToStringFunction_Once.Do(func() {
		flagsToStringFunction, err = gi.FunctionInvokerNew("GObject", "flags_to_string")
	})
	return err
}

// FlagsToString is a representation of the C type g_flags_to_string.
func FlagsToString(flagsType int64, value uint32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(flagsType)
	inArgs[1].SetUint32(value)

	var ret gi.Argument

	err := flagsToStringFunction_Set()
	if err == nil {
		ret = flagsToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var gtypeGetTypeFunction *gi.Function
var gtypeGetTypeFunction_Once sync.Once

func gtypeGetTypeFunction_Set() error {
	var err error
	gtypeGetTypeFunction_Once.Do(func() {
		gtypeGetTypeFunction, err = gi.FunctionInvokerNew("GObject", "gtype_get_type")
	})
	return err
}

// GtypeGetType is a representation of the C type g_gtype_get_type.
func GtypeGetType() int64 {

	var ret gi.Argument

	err := gtypeGetTypeFunction_Set()
	if err == nil {
		ret = gtypeGetTypeFunction.Invoke(nil, nil)
	}

	retGo := ret.Int64()

	return retGo
}

var paramSpecBooleanFunction *gi.Function
var paramSpecBooleanFunction_Once sync.Once

func paramSpecBooleanFunction_Set() error {
	var err error
	paramSpecBooleanFunction_Once.Do(func() {
		paramSpecBooleanFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_boolean")
	})
	return err
}

// ParamSpecBoolean_ is a representation of the C type g_param_spec_boolean.
func ParamSpecBoolean_(name string, nick string, blurb string, defaultValue bool, flags ParamFlags) *ParamSpec {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetBoolean(defaultValue)
	inArgs[4].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecBooleanFunction_Set()
	if err == nil {
		ret = paramSpecBooleanFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecBoxedFunction *gi.Function
var paramSpecBoxedFunction_Once sync.Once

func paramSpecBoxedFunction_Set() error {
	var err error
	paramSpecBoxedFunction_Once.Do(func() {
		paramSpecBoxedFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_boxed")
	})
	return err
}

// ParamSpecBoxed_ is a representation of the C type g_param_spec_boxed.
func ParamSpecBoxed_(name string, nick string, blurb string, boxedType int64, flags ParamFlags) *ParamSpec {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt64(boxedType)
	inArgs[4].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecBoxedFunction_Set()
	if err == nil {
		ret = paramSpecBoxedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecCharFunction *gi.Function
var paramSpecCharFunction_Once sync.Once

func paramSpecCharFunction_Set() error {
	var err error
	paramSpecCharFunction_Once.Do(func() {
		paramSpecCharFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_char")
	})
	return err
}

// ParamSpecChar_ is a representation of the C type g_param_spec_char.
func ParamSpecChar_(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt8(minimum)
	inArgs[4].SetInt8(maximum)
	inArgs[5].SetInt8(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecCharFunction_Set()
	if err == nil {
		ret = paramSpecCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecDoubleFunction *gi.Function
var paramSpecDoubleFunction_Once sync.Once

func paramSpecDoubleFunction_Set() error {
	var err error
	paramSpecDoubleFunction_Once.Do(func() {
		paramSpecDoubleFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_double")
	})
	return err
}

// ParamSpecDouble_ is a representation of the C type g_param_spec_double.
func ParamSpecDouble_(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetFloat64(minimum)
	inArgs[4].SetFloat64(maximum)
	inArgs[5].SetFloat64(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecDoubleFunction_Set()
	if err == nil {
		ret = paramSpecDoubleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecEnumFunction *gi.Function
var paramSpecEnumFunction_Once sync.Once

func paramSpecEnumFunction_Set() error {
	var err error
	paramSpecEnumFunction_Once.Do(func() {
		paramSpecEnumFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_enum")
	})
	return err
}

// ParamSpecEnum_ is a representation of the C type g_param_spec_enum.
func ParamSpecEnum_(name string, nick string, blurb string, enumType int64, defaultValue int32, flags ParamFlags) *ParamSpec {
	var inArgs [6]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt64(enumType)
	inArgs[4].SetInt32(defaultValue)
	inArgs[5].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecEnumFunction_Set()
	if err == nil {
		ret = paramSpecEnumFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecFlagsFunction *gi.Function
var paramSpecFlagsFunction_Once sync.Once

func paramSpecFlagsFunction_Set() error {
	var err error
	paramSpecFlagsFunction_Once.Do(func() {
		paramSpecFlagsFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_flags")
	})
	return err
}

// ParamSpecFlags_ is a representation of the C type g_param_spec_flags.
func ParamSpecFlags_(name string, nick string, blurb string, flagsType int64, defaultValue uint32, flags ParamFlags) *ParamSpec {
	var inArgs [6]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt64(flagsType)
	inArgs[4].SetUint32(defaultValue)
	inArgs[5].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecFlagsFunction_Set()
	if err == nil {
		ret = paramSpecFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecFloatFunction *gi.Function
var paramSpecFloatFunction_Once sync.Once

func paramSpecFloatFunction_Set() error {
	var err error
	paramSpecFloatFunction_Once.Do(func() {
		paramSpecFloatFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_float")
	})
	return err
}

// ParamSpecFloat_ is a representation of the C type g_param_spec_float.
func ParamSpecFloat_(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetFloat32(minimum)
	inArgs[4].SetFloat32(maximum)
	inArgs[5].SetFloat32(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecFloatFunction_Set()
	if err == nil {
		ret = paramSpecFloatFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecGtypeFunction *gi.Function
var paramSpecGtypeFunction_Once sync.Once

func paramSpecGtypeFunction_Set() error {
	var err error
	paramSpecGtypeFunction_Once.Do(func() {
		paramSpecGtypeFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_gtype")
	})
	return err
}

// ParamSpecGtype is a representation of the C type g_param_spec_gtype.
func ParamSpecGtype(name string, nick string, blurb string, isAType int64, flags ParamFlags) *ParamSpec {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt64(isAType)
	inArgs[4].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecGtypeFunction_Set()
	if err == nil {
		ret = paramSpecGtypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecIntFunction *gi.Function
var paramSpecIntFunction_Once sync.Once

func paramSpecIntFunction_Set() error {
	var err error
	paramSpecIntFunction_Once.Do(func() {
		paramSpecIntFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_int")
	})
	return err
}

// ParamSpecInt_ is a representation of the C type g_param_spec_int.
func ParamSpecInt_(name string, nick string, blurb string, minimum int32, maximum int32, defaultValue int32, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt32(minimum)
	inArgs[4].SetInt32(maximum)
	inArgs[5].SetInt32(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecIntFunction_Set()
	if err == nil {
		ret = paramSpecIntFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecInt64Function *gi.Function
var paramSpecInt64Function_Once sync.Once

func paramSpecInt64Function_Set() error {
	var err error
	paramSpecInt64Function_Once.Do(func() {
		paramSpecInt64Function, err = gi.FunctionInvokerNew("GObject", "param_spec_int64")
	})
	return err
}

// ParamSpecInt64_ is a representation of the C type g_param_spec_int64.
func ParamSpecInt64_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt64(minimum)
	inArgs[4].SetInt64(maximum)
	inArgs[5].SetInt64(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecInt64Function_Set()
	if err == nil {
		ret = paramSpecInt64Function.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecLongFunction *gi.Function
var paramSpecLongFunction_Once sync.Once

func paramSpecLongFunction_Set() error {
	var err error
	paramSpecLongFunction_Once.Do(func() {
		paramSpecLongFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_long")
	})
	return err
}

// ParamSpecLong_ is a representation of the C type g_param_spec_long.
func ParamSpecLong_(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt64(minimum)
	inArgs[4].SetInt64(maximum)
	inArgs[5].SetInt64(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecLongFunction_Set()
	if err == nil {
		ret = paramSpecLongFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecObjectFunction *gi.Function
var paramSpecObjectFunction_Once sync.Once

func paramSpecObjectFunction_Set() error {
	var err error
	paramSpecObjectFunction_Once.Do(func() {
		paramSpecObjectFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_object")
	})
	return err
}

// ParamSpecObject_ is a representation of the C type g_param_spec_object.
func ParamSpecObject_(name string, nick string, blurb string, objectType int64, flags ParamFlags) *ParamSpec {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt64(objectType)
	inArgs[4].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecObjectFunction_Set()
	if err == nil {
		ret = paramSpecObjectFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecOverrideFunction *gi.Function
var paramSpecOverrideFunction_Once sync.Once

func paramSpecOverrideFunction_Set() error {
	var err error
	paramSpecOverrideFunction_Once.Do(func() {
		paramSpecOverrideFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_override")
	})
	return err
}

// ParamSpecOverride_ is a representation of the C type g_param_spec_override.
func ParamSpecOverride_(name string, overridden *ParamSpec) *ParamSpec {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetPointer(overridden.Native())

	var ret gi.Argument

	err := paramSpecOverrideFunction_Set()
	if err == nil {
		ret = paramSpecOverrideFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecParamFunction *gi.Function
var paramSpecParamFunction_Once sync.Once

func paramSpecParamFunction_Set() error {
	var err error
	paramSpecParamFunction_Once.Do(func() {
		paramSpecParamFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_param")
	})
	return err
}

// ParamSpecParam_ is a representation of the C type g_param_spec_param.
func ParamSpecParam_(name string, nick string, blurb string, paramType int64, flags ParamFlags) *ParamSpec {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt64(paramType)
	inArgs[4].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecParamFunction_Set()
	if err == nil {
		ret = paramSpecParamFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecPointerFunction *gi.Function
var paramSpecPointerFunction_Once sync.Once

func paramSpecPointerFunction_Set() error {
	var err error
	paramSpecPointerFunction_Once.Do(func() {
		paramSpecPointerFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_pointer")
	})
	return err
}

// ParamSpecPointer_ is a representation of the C type g_param_spec_pointer.
func ParamSpecPointer_(name string, nick string, blurb string, flags ParamFlags) *ParamSpec {
	var inArgs [4]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecPointerFunction_Set()
	if err == nil {
		ret = paramSpecPointerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecPoolNewFunction *gi.Function
var paramSpecPoolNewFunction_Once sync.Once

func paramSpecPoolNewFunction_Set() error {
	var err error
	paramSpecPoolNewFunction_Once.Do(func() {
		paramSpecPoolNewFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_pool_new")
	})
	return err
}

// ParamSpecPoolNew is a representation of the C type g_param_spec_pool_new.
func ParamSpecPoolNew(typePrefixing bool) *ParamSpecPool {
	var inArgs [1]gi.Argument
	inArgs[0].SetBoolean(typePrefixing)

	var ret gi.Argument

	err := paramSpecPoolNewFunction_Set()
	if err == nil {
		ret = paramSpecPoolNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecPoolNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecStringFunction *gi.Function
var paramSpecStringFunction_Once sync.Once

func paramSpecStringFunction_Set() error {
	var err error
	paramSpecStringFunction_Once.Do(func() {
		paramSpecStringFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_string")
	})
	return err
}

// ParamSpecString_ is a representation of the C type g_param_spec_string.
func ParamSpecString_(name string, nick string, blurb string, defaultValue string, flags ParamFlags) *ParamSpec {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetString(defaultValue)
	inArgs[4].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecStringFunction_Set()
	if err == nil {
		ret = paramSpecStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecUcharFunction *gi.Function
var paramSpecUcharFunction_Once sync.Once

func paramSpecUcharFunction_Set() error {
	var err error
	paramSpecUcharFunction_Once.Do(func() {
		paramSpecUcharFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_uchar")
	})
	return err
}

// ParamSpecUchar is a representation of the C type g_param_spec_uchar.
func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, defaultValue uint8, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetUint8(minimum)
	inArgs[4].SetUint8(maximum)
	inArgs[5].SetUint8(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecUcharFunction_Set()
	if err == nil {
		ret = paramSpecUcharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecUintFunction *gi.Function
var paramSpecUintFunction_Once sync.Once

func paramSpecUintFunction_Set() error {
	var err error
	paramSpecUintFunction_Once.Do(func() {
		paramSpecUintFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_uint")
	})
	return err
}

// ParamSpecUint is a representation of the C type g_param_spec_uint.
func ParamSpecUint(name string, nick string, blurb string, minimum uint32, maximum uint32, defaultValue uint32, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetUint32(minimum)
	inArgs[4].SetUint32(maximum)
	inArgs[5].SetUint32(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecUintFunction_Set()
	if err == nil {
		ret = paramSpecUintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecUint64Function *gi.Function
var paramSpecUint64Function_Once sync.Once

func paramSpecUint64Function_Set() error {
	var err error
	paramSpecUint64Function_Once.Do(func() {
		paramSpecUint64Function, err = gi.FunctionInvokerNew("GObject", "param_spec_uint64")
	})
	return err
}

// ParamSpecUint64 is a representation of the C type g_param_spec_uint64.
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetUint64(minimum)
	inArgs[4].SetUint64(maximum)
	inArgs[5].SetUint64(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecUint64Function_Set()
	if err == nil {
		ret = paramSpecUint64Function.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecUlongFunction *gi.Function
var paramSpecUlongFunction_Once sync.Once

func paramSpecUlongFunction_Set() error {
	var err error
	paramSpecUlongFunction_Once.Do(func() {
		paramSpecUlongFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_ulong")
	})
	return err
}

// ParamSpecUlong is a representation of the C type g_param_spec_ulong.
func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) *ParamSpec {
	var inArgs [7]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetUint64(minimum)
	inArgs[4].SetUint64(maximum)
	inArgs[5].SetUint64(defaultValue)
	inArgs[6].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecUlongFunction_Set()
	if err == nil {
		ret = paramSpecUlongFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_param_spec_unichar' : parameter 'default_value' of type 'gunichar' not supported

var paramSpecValueArrayFunction *gi.Function
var paramSpecValueArrayFunction_Once sync.Once

func paramSpecValueArrayFunction_Set() error {
	var err error
	paramSpecValueArrayFunction_Once.Do(func() {
		paramSpecValueArrayFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_value_array")
	})
	return err
}

// ParamSpecValueArray_ is a representation of the C type g_param_spec_value_array.
func ParamSpecValueArray_(name string, nick string, blurb string, elementSpec *ParamSpec, flags ParamFlags) *ParamSpec {
	var inArgs [5]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetPointer(elementSpec.Native())
	inArgs[4].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecValueArrayFunction_Set()
	if err == nil {
		ret = paramSpecValueArrayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramSpecVariantFunction *gi.Function
var paramSpecVariantFunction_Once sync.Once

func paramSpecVariantFunction_Set() error {
	var err error
	paramSpecVariantFunction_Once.Do(func() {
		paramSpecVariantFunction, err = gi.FunctionInvokerNew("GObject", "param_spec_variant")
	})
	return err
}

// ParamSpecVariant_ is a representation of the C type g_param_spec_variant.
func ParamSpecVariant_(name string, nick string, blurb string, type_ *glib.VariantType, defaultValue *glib.Variant, flags ParamFlags) *ParamSpec {
	var inArgs [6]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetString(nick)
	inArgs[2].SetString(blurb)
	inArgs[3].SetPointer(type_.Native())
	inArgs[4].SetPointer(defaultValue.Native())
	inArgs[5].SetInt32(int32(flags))

	var ret gi.Argument

	err := paramSpecVariantFunction_Set()
	if err == nil {
		ret = paramSpecVariantFunction.Invoke(inArgs[:], nil)
	}

	retGo := ParamSpecNewFromNative(ret.Pointer())

	return retGo
}

var paramTypeRegisterStaticFunction *gi.Function
var paramTypeRegisterStaticFunction_Once sync.Once

func paramTypeRegisterStaticFunction_Set() error {
	var err error
	paramTypeRegisterStaticFunction_Once.Do(func() {
		paramTypeRegisterStaticFunction, err = gi.FunctionInvokerNew("GObject", "param_type_register_static")
	})
	return err
}

// ParamTypeRegisterStatic is a representation of the C type g_param_type_register_static.
func ParamTypeRegisterStatic(name string, pspecInfo *ParamSpecTypeInfo) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetPointer(pspecInfo.Native())

	var ret gi.Argument

	err := paramTypeRegisterStaticFunction_Set()
	if err == nil {
		ret = paramTypeRegisterStaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var paramValueConvertFunction *gi.Function
var paramValueConvertFunction_Once sync.Once

func paramValueConvertFunction_Set() error {
	var err error
	paramValueConvertFunction_Once.Do(func() {
		paramValueConvertFunction, err = gi.FunctionInvokerNew("GObject", "param_value_convert")
	})
	return err
}

// ParamValueConvert is a representation of the C type g_param_value_convert.
func ParamValueConvert(pspec *ParamSpec, srcValue *Value, destValue *Value, strictValidation bool) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(pspec.Native())
	inArgs[1].SetPointer(srcValue.Native())
	inArgs[2].SetPointer(destValue.Native())
	inArgs[3].SetBoolean(strictValidation)

	var ret gi.Argument

	err := paramValueConvertFunction_Set()
	if err == nil {
		ret = paramValueConvertFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var paramValueDefaultsFunction *gi.Function
var paramValueDefaultsFunction_Once sync.Once

func paramValueDefaultsFunction_Set() error {
	var err error
	paramValueDefaultsFunction_Once.Do(func() {
		paramValueDefaultsFunction, err = gi.FunctionInvokerNew("GObject", "param_value_defaults")
	})
	return err
}

// ParamValueDefaults is a representation of the C type g_param_value_defaults.
func ParamValueDefaults(pspec *ParamSpec, value *Value) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(pspec.Native())
	inArgs[1].SetPointer(value.Native())

	var ret gi.Argument

	err := paramValueDefaultsFunction_Set()
	if err == nil {
		ret = paramValueDefaultsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var paramValueSetDefaultFunction *gi.Function
var paramValueSetDefaultFunction_Once sync.Once

func paramValueSetDefaultFunction_Set() error {
	var err error
	paramValueSetDefaultFunction_Once.Do(func() {
		paramValueSetDefaultFunction, err = gi.FunctionInvokerNew("GObject", "param_value_set_default")
	})
	return err
}

// ParamValueSetDefault is a representation of the C type g_param_value_set_default.
func ParamValueSetDefault(pspec *ParamSpec, value *Value) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(pspec.Native())
	inArgs[1].SetPointer(value.Native())

	err := paramValueSetDefaultFunction_Set()
	if err == nil {
		paramValueSetDefaultFunction.Invoke(inArgs[:], nil)
	}

	return
}

var paramValueValidateFunction *gi.Function
var paramValueValidateFunction_Once sync.Once

func paramValueValidateFunction_Set() error {
	var err error
	paramValueValidateFunction_Once.Do(func() {
		paramValueValidateFunction, err = gi.FunctionInvokerNew("GObject", "param_value_validate")
	})
	return err
}

// ParamValueValidate is a representation of the C type g_param_value_validate.
func ParamValueValidate(pspec *ParamSpec, value *Value) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(pspec.Native())
	inArgs[1].SetPointer(value.Native())

	var ret gi.Argument

	err := paramValueValidateFunction_Set()
	if err == nil {
		ret = paramValueValidateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var paramValuesCmpFunction *gi.Function
var paramValuesCmpFunction_Once sync.Once

func paramValuesCmpFunction_Set() error {
	var err error
	paramValuesCmpFunction_Once.Do(func() {
		paramValuesCmpFunction, err = gi.FunctionInvokerNew("GObject", "param_values_cmp")
	})
	return err
}

// ParamValuesCmp is a representation of the C type g_param_values_cmp.
func ParamValuesCmp(pspec *ParamSpec, value1 *Value, value2 *Value) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(pspec.Native())
	inArgs[1].SetPointer(value1.Native())
	inArgs[2].SetPointer(value2.Native())

	var ret gi.Argument

	err := paramValuesCmpFunction_Set()
	if err == nil {
		ret = paramValuesCmpFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var pointerTypeRegisterStaticFunction *gi.Function
var pointerTypeRegisterStaticFunction_Once sync.Once

func pointerTypeRegisterStaticFunction_Set() error {
	var err error
	pointerTypeRegisterStaticFunction_Once.Do(func() {
		pointerTypeRegisterStaticFunction, err = gi.FunctionInvokerNew("GObject", "pointer_type_register_static")
	})
	return err
}

// PointerTypeRegisterStatic is a representation of the C type g_pointer_type_register_static.
func PointerTypeRegisterStatic(name string) int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := pointerTypeRegisterStaticFunction_Set()
	if err == nil {
		ret = pointerTypeRegisterStaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var signalAccumulatorFirstWinsFunction *gi.Function
var signalAccumulatorFirstWinsFunction_Once sync.Once

func signalAccumulatorFirstWinsFunction_Set() error {
	var err error
	signalAccumulatorFirstWinsFunction_Once.Do(func() {
		signalAccumulatorFirstWinsFunction, err = gi.FunctionInvokerNew("GObject", "signal_accumulator_first_wins")
	})
	return err
}

// SignalAccumulatorFirstWins is a representation of the C type g_signal_accumulator_first_wins.
func SignalAccumulatorFirstWins(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(ihint.Native())
	inArgs[1].SetPointer(returnAccu.Native())
	inArgs[2].SetPointer(handlerReturn.Native())
	inArgs[3].SetPointer(dummy)

	var ret gi.Argument

	err := signalAccumulatorFirstWinsFunction_Set()
	if err == nil {
		ret = signalAccumulatorFirstWinsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var signalAccumulatorTrueHandledFunction *gi.Function
var signalAccumulatorTrueHandledFunction_Once sync.Once

func signalAccumulatorTrueHandledFunction_Set() error {
	var err error
	signalAccumulatorTrueHandledFunction_Once.Do(func() {
		signalAccumulatorTrueHandledFunction, err = gi.FunctionInvokerNew("GObject", "signal_accumulator_true_handled")
	})
	return err
}

// SignalAccumulatorTrueHandled is a representation of the C type g_signal_accumulator_true_handled.
func SignalAccumulatorTrueHandled(ihint *SignalInvocationHint, returnAccu *Value, handlerReturn *Value, dummy unsafe.Pointer) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(ihint.Native())
	inArgs[1].SetPointer(returnAccu.Native())
	inArgs[2].SetPointer(handlerReturn.Native())
	inArgs[3].SetPointer(dummy)

	var ret gi.Argument

	err := signalAccumulatorTrueHandledFunction_Set()
	if err == nil {
		ret = signalAccumulatorTrueHandledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'g_signal_add_emission_hook' : parameter 'hook_func' of type 'SignalEmissionHook' not supported

// UNSUPPORTED : C value 'g_signal_chain_from_overridden' : parameter 'instance_and_params' of type 'nil' not supported

// UNSUPPORTED : C value 'g_signal_chain_from_overridden_handler' : parameter '...' of type 'nil' not supported

var signalConnectClosureFunction *gi.Function
var signalConnectClosureFunction_Once sync.Once

func signalConnectClosureFunction_Set() error {
	var err error
	signalConnectClosureFunction_Once.Do(func() {
		signalConnectClosureFunction, err = gi.FunctionInvokerNew("GObject", "signal_connect_closure")
	})
	return err
}

// SignalConnectClosure is a representation of the C type g_signal_connect_closure.
func SignalConnectClosure(instance *Object, detailedSignal string, closure *Closure, after bool) uint64 {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetString(detailedSignal)
	inArgs[2].SetPointer(closure.Native())
	inArgs[3].SetBoolean(after)

	var ret gi.Argument

	err := signalConnectClosureFunction_Set()
	if err == nil {
		ret = signalConnectClosureFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var signalConnectClosureByIdFunction *gi.Function
var signalConnectClosureByIdFunction_Once sync.Once

func signalConnectClosureByIdFunction_Set() error {
	var err error
	signalConnectClosureByIdFunction_Once.Do(func() {
		signalConnectClosureByIdFunction, err = gi.FunctionInvokerNew("GObject", "signal_connect_closure_by_id")
	})
	return err
}

// SignalConnectClosureById is a representation of the C type g_signal_connect_closure_by_id.
func SignalConnectClosureById(instance *Object, signalId uint32, detail glib.Quark, closure *Closure, after bool) uint64 {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetUint32(signalId)
	inArgs[2].SetUint32(uint32(detail))
	inArgs[3].SetPointer(closure.Native())
	inArgs[4].SetBoolean(after)

	var ret gi.Argument

	err := signalConnectClosureByIdFunction_Set()
	if err == nil {
		ret = signalConnectClosureByIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'g_signal_connect_data' : parameter 'c_handler' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_signal_connect_object' : parameter 'c_handler' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_signal_emit' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_signal_emit_by_name' : parameter '...' of type 'nil' not supported

// UNSUPPORTED : C value 'g_signal_emit_valist' : parameter 'var_args' of type 'va_list' not supported

// UNSUPPORTED : C value 'g_signal_emitv' : parameter 'instance_and_params' of type 'nil' not supported

var signalGetInvocationHintFunction *gi.Function
var signalGetInvocationHintFunction_Once sync.Once

func signalGetInvocationHintFunction_Set() error {
	var err error
	signalGetInvocationHintFunction_Once.Do(func() {
		signalGetInvocationHintFunction, err = gi.FunctionInvokerNew("GObject", "signal_get_invocation_hint")
	})
	return err
}

// SignalGetInvocationHint is a representation of the C type g_signal_get_invocation_hint.
func SignalGetInvocationHint(instance *Object) *SignalInvocationHint {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(instance.Native())

	var ret gi.Argument

	err := signalGetInvocationHintFunction_Set()
	if err == nil {
		ret = signalGetInvocationHintFunction.Invoke(inArgs[:], nil)
	}

	retGo := SignalInvocationHintNewFromNative(ret.Pointer())

	return retGo
}

var signalHandlerBlockFunction *gi.Function
var signalHandlerBlockFunction_Once sync.Once

func signalHandlerBlockFunction_Set() error {
	var err error
	signalHandlerBlockFunction_Once.Do(func() {
		signalHandlerBlockFunction, err = gi.FunctionInvokerNew("GObject", "signal_handler_block")
	})
	return err
}

// SignalHandlerBlock is a representation of the C type g_signal_handler_block.
func SignalHandlerBlock(instance *Object, handlerId uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetUint64(handlerId)

	err := signalHandlerBlockFunction_Set()
	if err == nil {
		signalHandlerBlockFunction.Invoke(inArgs[:], nil)
	}

	return
}

var signalHandlerDisconnectFunction *gi.Function
var signalHandlerDisconnectFunction_Once sync.Once

func signalHandlerDisconnectFunction_Set() error {
	var err error
	signalHandlerDisconnectFunction_Once.Do(func() {
		signalHandlerDisconnectFunction, err = gi.FunctionInvokerNew("GObject", "signal_handler_disconnect")
	})
	return err
}

// SignalHandlerDisconnect is a representation of the C type g_signal_handler_disconnect.
func SignalHandlerDisconnect(instance *Object, handlerId uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetUint64(handlerId)

	err := signalHandlerDisconnectFunction_Set()
	if err == nil {
		signalHandlerDisconnectFunction.Invoke(inArgs[:], nil)
	}

	return
}

var signalHandlerFindFunction *gi.Function
var signalHandlerFindFunction_Once sync.Once

func signalHandlerFindFunction_Set() error {
	var err error
	signalHandlerFindFunction_Once.Do(func() {
		signalHandlerFindFunction, err = gi.FunctionInvokerNew("GObject", "signal_handler_find")
	})
	return err
}

// SignalHandlerFind is a representation of the C type g_signal_handler_find.
func SignalHandlerFind(instance *Object, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint64 {
	var inArgs [7]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetInt32(int32(mask))
	inArgs[2].SetUint32(signalId)
	inArgs[3].SetUint32(uint32(detail))
	inArgs[4].SetPointer(closure.Native())
	inArgs[5].SetPointer(func_)
	inArgs[6].SetPointer(data)

	var ret gi.Argument

	err := signalHandlerFindFunction_Set()
	if err == nil {
		ret = signalHandlerFindFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var signalHandlerIsConnectedFunction *gi.Function
var signalHandlerIsConnectedFunction_Once sync.Once

func signalHandlerIsConnectedFunction_Set() error {
	var err error
	signalHandlerIsConnectedFunction_Once.Do(func() {
		signalHandlerIsConnectedFunction, err = gi.FunctionInvokerNew("GObject", "signal_handler_is_connected")
	})
	return err
}

// SignalHandlerIsConnected is a representation of the C type g_signal_handler_is_connected.
func SignalHandlerIsConnected(instance *Object, handlerId uint64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetUint64(handlerId)

	var ret gi.Argument

	err := signalHandlerIsConnectedFunction_Set()
	if err == nil {
		ret = signalHandlerIsConnectedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var signalHandlerUnblockFunction *gi.Function
var signalHandlerUnblockFunction_Once sync.Once

func signalHandlerUnblockFunction_Set() error {
	var err error
	signalHandlerUnblockFunction_Once.Do(func() {
		signalHandlerUnblockFunction, err = gi.FunctionInvokerNew("GObject", "signal_handler_unblock")
	})
	return err
}

// SignalHandlerUnblock is a representation of the C type g_signal_handler_unblock.
func SignalHandlerUnblock(instance *Object, handlerId uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetUint64(handlerId)

	err := signalHandlerUnblockFunction_Set()
	if err == nil {
		signalHandlerUnblockFunction.Invoke(inArgs[:], nil)
	}

	return
}

var signalHandlersBlockMatchedFunction *gi.Function
var signalHandlersBlockMatchedFunction_Once sync.Once

func signalHandlersBlockMatchedFunction_Set() error {
	var err error
	signalHandlersBlockMatchedFunction_Once.Do(func() {
		signalHandlersBlockMatchedFunction, err = gi.FunctionInvokerNew("GObject", "signal_handlers_block_matched")
	})
	return err
}

// SignalHandlersBlockMatched is a representation of the C type g_signal_handlers_block_matched.
func SignalHandlersBlockMatched(instance *Object, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint32 {
	var inArgs [7]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetInt32(int32(mask))
	inArgs[2].SetUint32(signalId)
	inArgs[3].SetUint32(uint32(detail))
	inArgs[4].SetPointer(closure.Native())
	inArgs[5].SetPointer(func_)
	inArgs[6].SetPointer(data)

	var ret gi.Argument

	err := signalHandlersBlockMatchedFunction_Set()
	if err == nil {
		ret = signalHandlersBlockMatchedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var signalHandlersDestroyFunction *gi.Function
var signalHandlersDestroyFunction_Once sync.Once

func signalHandlersDestroyFunction_Set() error {
	var err error
	signalHandlersDestroyFunction_Once.Do(func() {
		signalHandlersDestroyFunction, err = gi.FunctionInvokerNew("GObject", "signal_handlers_destroy")
	})
	return err
}

// SignalHandlersDestroy is a representation of the C type g_signal_handlers_destroy.
func SignalHandlersDestroy(instance *Object) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(instance.Native())

	err := signalHandlersDestroyFunction_Set()
	if err == nil {
		signalHandlersDestroyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var signalHandlersDisconnectMatchedFunction *gi.Function
var signalHandlersDisconnectMatchedFunction_Once sync.Once

func signalHandlersDisconnectMatchedFunction_Set() error {
	var err error
	signalHandlersDisconnectMatchedFunction_Once.Do(func() {
		signalHandlersDisconnectMatchedFunction, err = gi.FunctionInvokerNew("GObject", "signal_handlers_disconnect_matched")
	})
	return err
}

// SignalHandlersDisconnectMatched is a representation of the C type g_signal_handlers_disconnect_matched.
func SignalHandlersDisconnectMatched(instance *Object, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint32 {
	var inArgs [7]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetInt32(int32(mask))
	inArgs[2].SetUint32(signalId)
	inArgs[3].SetUint32(uint32(detail))
	inArgs[4].SetPointer(closure.Native())
	inArgs[5].SetPointer(func_)
	inArgs[6].SetPointer(data)

	var ret gi.Argument

	err := signalHandlersDisconnectMatchedFunction_Set()
	if err == nil {
		ret = signalHandlersDisconnectMatchedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var signalHandlersUnblockMatchedFunction *gi.Function
var signalHandlersUnblockMatchedFunction_Once sync.Once

func signalHandlersUnblockMatchedFunction_Set() error {
	var err error
	signalHandlersUnblockMatchedFunction_Once.Do(func() {
		signalHandlersUnblockMatchedFunction, err = gi.FunctionInvokerNew("GObject", "signal_handlers_unblock_matched")
	})
	return err
}

// SignalHandlersUnblockMatched is a representation of the C type g_signal_handlers_unblock_matched.
func SignalHandlersUnblockMatched(instance *Object, mask SignalMatchType, signalId uint32, detail glib.Quark, closure *Closure, func_ unsafe.Pointer, data unsafe.Pointer) uint32 {
	var inArgs [7]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetInt32(int32(mask))
	inArgs[2].SetUint32(signalId)
	inArgs[3].SetUint32(uint32(detail))
	inArgs[4].SetPointer(closure.Native())
	inArgs[5].SetPointer(func_)
	inArgs[6].SetPointer(data)

	var ret gi.Argument

	err := signalHandlersUnblockMatchedFunction_Set()
	if err == nil {
		ret = signalHandlersUnblockMatchedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var signalHasHandlerPendingFunction *gi.Function
var signalHasHandlerPendingFunction_Once sync.Once

func signalHasHandlerPendingFunction_Set() error {
	var err error
	signalHasHandlerPendingFunction_Once.Do(func() {
		signalHasHandlerPendingFunction, err = gi.FunctionInvokerNew("GObject", "signal_has_handler_pending")
	})
	return err
}

// SignalHasHandlerPending is a representation of the C type g_signal_has_handler_pending.
func SignalHasHandlerPending(instance *Object, signalId uint32, detail glib.Quark, mayBeBlocked bool) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetUint32(signalId)
	inArgs[2].SetUint32(uint32(detail))
	inArgs[3].SetBoolean(mayBeBlocked)

	var ret gi.Argument

	err := signalHasHandlerPendingFunction_Set()
	if err == nil {
		ret = signalHasHandlerPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var signalListIdsFunction *gi.Function
var signalListIdsFunction_Once sync.Once

func signalListIdsFunction_Set() error {
	var err error
	signalListIdsFunction_Once.Do(func() {
		signalListIdsFunction, err = gi.FunctionInvokerNew("GObject", "signal_list_ids")
	})
	return err
}

// SignalListIds is a representation of the C type g_signal_list_ids.
func SignalListIds(itype int64) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(itype)

	var outArgs [1]gi.Argument

	err := signalListIdsFunction_Set()
	if err == nil {
		signalListIdsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
}

var signalLookupFunction *gi.Function
var signalLookupFunction_Once sync.Once

func signalLookupFunction_Set() error {
	var err error
	signalLookupFunction_Once.Do(func() {
		signalLookupFunction, err = gi.FunctionInvokerNew("GObject", "signal_lookup")
	})
	return err
}

// SignalLookup is a representation of the C type g_signal_lookup.
func SignalLookup(name string, itype int64) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(name)
	inArgs[1].SetInt64(itype)

	var ret gi.Argument

	err := signalLookupFunction_Set()
	if err == nil {
		ret = signalLookupFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var signalNameFunction *gi.Function
var signalNameFunction_Once sync.Once

func signalNameFunction_Set() error {
	var err error
	signalNameFunction_Once.Do(func() {
		signalNameFunction, err = gi.FunctionInvokerNew("GObject", "signal_name")
	})
	return err
}

// SignalName is a representation of the C type g_signal_name.
func SignalName(signalId uint32) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(signalId)

	var ret gi.Argument

	err := signalNameFunction_Set()
	if err == nil {
		ret = signalNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'g_signal_new' : parameter 'accumulator' of type 'SignalAccumulator' not supported

// UNSUPPORTED : C value 'g_signal_new_class_handler' : parameter 'class_handler' of type 'Callback' not supported

// UNSUPPORTED : C value 'g_signal_new_valist' : parameter 'accumulator' of type 'SignalAccumulator' not supported

// UNSUPPORTED : C value 'g_signal_newv' : parameter 'accumulator' of type 'SignalAccumulator' not supported

var signalOverrideClassClosureFunction *gi.Function
var signalOverrideClassClosureFunction_Once sync.Once

func signalOverrideClassClosureFunction_Set() error {
	var err error
	signalOverrideClassClosureFunction_Once.Do(func() {
		signalOverrideClassClosureFunction, err = gi.FunctionInvokerNew("GObject", "signal_override_class_closure")
	})
	return err
}

// SignalOverrideClassClosure is a representation of the C type g_signal_override_class_closure.
func SignalOverrideClassClosure(signalId uint32, instanceType int64, classClosure *Closure) {
	var inArgs [3]gi.Argument
	inArgs[0].SetUint32(signalId)
	inArgs[1].SetInt64(instanceType)
	inArgs[2].SetPointer(classClosure.Native())

	err := signalOverrideClassClosureFunction_Set()
	if err == nil {
		signalOverrideClassClosureFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_signal_override_class_handler' : parameter 'class_handler' of type 'Callback' not supported

var signalParseNameFunction *gi.Function
var signalParseNameFunction_Once sync.Once

func signalParseNameFunction_Set() error {
	var err error
	signalParseNameFunction_Once.Do(func() {
		signalParseNameFunction, err = gi.FunctionInvokerNew("GObject", "signal_parse_name")
	})
	return err
}

// SignalParseName is a representation of the C type g_signal_parse_name.
func SignalParseName(detailedSignal string, itype int64, forceDetailQuark bool) (bool, uint32, glib.Quark) {
	var inArgs [3]gi.Argument
	inArgs[0].SetString(detailedSignal)
	inArgs[1].SetInt64(itype)
	inArgs[2].SetBoolean(forceDetailQuark)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := signalParseNameFunction_Set()
	if err == nil {
		ret = signalParseNameFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Uint32()
	out1 := glib.Quark(outArgs[1].Uint32())

	return retGo, out0, out1
}

var signalQueryFunction *gi.Function
var signalQueryFunction_Once sync.Once

func signalQueryFunction_Set() error {
	var err error
	signalQueryFunction_Once.Do(func() {
		signalQueryFunction, err = gi.FunctionInvokerNew("GObject", "signal_query")
	})
	return err
}

// SignalQuery__ is a representation of the C type g_signal_query.
func SignalQuery__(signalId uint32) *SignalQuery_ {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(signalId)

	var outArgs [1]gi.Argument

	err := signalQueryFunction_Set()
	if err == nil {
		signalQueryFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := SignalQuery_NewFromNative(outArgs[0].Pointer())

	return out0
}

var signalRemoveEmissionHookFunction *gi.Function
var signalRemoveEmissionHookFunction_Once sync.Once

func signalRemoveEmissionHookFunction_Set() error {
	var err error
	signalRemoveEmissionHookFunction_Once.Do(func() {
		signalRemoveEmissionHookFunction, err = gi.FunctionInvokerNew("GObject", "signal_remove_emission_hook")
	})
	return err
}

// SignalRemoveEmissionHook is a representation of the C type g_signal_remove_emission_hook.
func SignalRemoveEmissionHook(signalId uint32, hookId uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetUint32(signalId)
	inArgs[1].SetUint64(hookId)

	err := signalRemoveEmissionHookFunction_Set()
	if err == nil {
		signalRemoveEmissionHookFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_signal_set_va_marshaller' : parameter 'va_marshaller' of type 'SignalCVaMarshaller' not supported

var signalStopEmissionFunction *gi.Function
var signalStopEmissionFunction_Once sync.Once

func signalStopEmissionFunction_Set() error {
	var err error
	signalStopEmissionFunction_Once.Do(func() {
		signalStopEmissionFunction, err = gi.FunctionInvokerNew("GObject", "signal_stop_emission")
	})
	return err
}

// SignalStopEmission is a representation of the C type g_signal_stop_emission.
func SignalStopEmission(instance *Object, signalId uint32, detail glib.Quark) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetUint32(signalId)
	inArgs[2].SetUint32(uint32(detail))

	err := signalStopEmissionFunction_Set()
	if err == nil {
		signalStopEmissionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var signalStopEmissionByNameFunction *gi.Function
var signalStopEmissionByNameFunction_Once sync.Once

func signalStopEmissionByNameFunction_Set() error {
	var err error
	signalStopEmissionByNameFunction_Once.Do(func() {
		signalStopEmissionByNameFunction, err = gi.FunctionInvokerNew("GObject", "signal_stop_emission_by_name")
	})
	return err
}

// SignalStopEmissionByName is a representation of the C type g_signal_stop_emission_by_name.
func SignalStopEmissionByName(instance *Object, detailedSignal string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetString(detailedSignal)

	err := signalStopEmissionByNameFunction_Set()
	if err == nil {
		signalStopEmissionByNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var signalTypeCclosureNewFunction *gi.Function
var signalTypeCclosureNewFunction_Once sync.Once

func signalTypeCclosureNewFunction_Set() error {
	var err error
	signalTypeCclosureNewFunction_Once.Do(func() {
		signalTypeCclosureNewFunction, err = gi.FunctionInvokerNew("GObject", "signal_type_cclosure_new")
	})
	return err
}

// SignalTypeCclosureNew is a representation of the C type g_signal_type_cclosure_new.
func SignalTypeCclosureNew(itype int64, structOffset uint32) *Closure {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(itype)
	inArgs[1].SetUint32(structOffset)

	var ret gi.Argument

	err := signalTypeCclosureNewFunction_Set()
	if err == nil {
		ret = signalTypeCclosureNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := ClosureNewFromNative(ret.Pointer())

	return retGo
}

var sourceSetClosureFunction *gi.Function
var sourceSetClosureFunction_Once sync.Once

func sourceSetClosureFunction_Set() error {
	var err error
	sourceSetClosureFunction_Once.Do(func() {
		sourceSetClosureFunction, err = gi.FunctionInvokerNew("GObject", "source_set_closure")
	})
	return err
}

// SourceSetClosure is a representation of the C type g_source_set_closure.
func SourceSetClosure(source *glib.Source, closure *Closure) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(source.Native())
	inArgs[1].SetPointer(closure.Native())

	err := sourceSetClosureFunction_Set()
	if err == nil {
		sourceSetClosureFunction.Invoke(inArgs[:], nil)
	}

	return
}

var sourceSetDummyCallbackFunction *gi.Function
var sourceSetDummyCallbackFunction_Once sync.Once

func sourceSetDummyCallbackFunction_Set() error {
	var err error
	sourceSetDummyCallbackFunction_Once.Do(func() {
		sourceSetDummyCallbackFunction, err = gi.FunctionInvokerNew("GObject", "source_set_dummy_callback")
	})
	return err
}

// SourceSetDummyCallback is a representation of the C type g_source_set_dummy_callback.
func SourceSetDummyCallback(source *glib.Source) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(source.Native())

	err := sourceSetDummyCallbackFunction_Set()
	if err == nil {
		sourceSetDummyCallbackFunction.Invoke(inArgs[:], nil)
	}

	return
}

var strdupValueContentsFunction *gi.Function
var strdupValueContentsFunction_Once sync.Once

func strdupValueContentsFunction_Set() error {
	var err error
	strdupValueContentsFunction_Once.Do(func() {
		strdupValueContentsFunction, err = gi.FunctionInvokerNew("GObject", "strdup_value_contents")
	})
	return err
}

// StrdupValueContents is a representation of the C type g_strdup_value_contents.
func StrdupValueContents(value *Value) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(value.Native())

	var ret gi.Argument

	err := strdupValueContentsFunction_Set()
	if err == nil {
		ret = strdupValueContentsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'g_type_add_class_cache_func' : parameter 'cache_func' of type 'TypeClassCacheFunc' not supported

var typeAddClassPrivateFunction *gi.Function
var typeAddClassPrivateFunction_Once sync.Once

func typeAddClassPrivateFunction_Set() error {
	var err error
	typeAddClassPrivateFunction_Once.Do(func() {
		typeAddClassPrivateFunction, err = gi.FunctionInvokerNew("GObject", "type_add_class_private")
	})
	return err
}

// TypeAddClassPrivate is a representation of the C type g_type_add_class_private.
func TypeAddClassPrivate(classType int64, privateSize uint64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(classType)
	inArgs[1].SetUint64(privateSize)

	err := typeAddClassPrivateFunction_Set()
	if err == nil {
		typeAddClassPrivateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeAddInstancePrivateFunction *gi.Function
var typeAddInstancePrivateFunction_Once sync.Once

func typeAddInstancePrivateFunction_Set() error {
	var err error
	typeAddInstancePrivateFunction_Once.Do(func() {
		typeAddInstancePrivateFunction, err = gi.FunctionInvokerNew("GObject", "type_add_instance_private")
	})
	return err
}

// TypeAddInstancePrivate is a representation of the C type g_type_add_instance_private.
func TypeAddInstancePrivate(classType int64, privateSize uint64) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(classType)
	inArgs[1].SetUint64(privateSize)

	var ret gi.Argument

	err := typeAddInstancePrivateFunction_Set()
	if err == nil {
		ret = typeAddInstancePrivateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_type_add_interface_check' : parameter 'check_func' of type 'TypeInterfaceCheckFunc' not supported

// UNSUPPORTED : C value 'g_type_add_interface_dynamic' : parameter 'plugin' of type 'TypePlugin' not supported

var typeAddInterfaceStaticFunction *gi.Function
var typeAddInterfaceStaticFunction_Once sync.Once

func typeAddInterfaceStaticFunction_Set() error {
	var err error
	typeAddInterfaceStaticFunction_Once.Do(func() {
		typeAddInterfaceStaticFunction, err = gi.FunctionInvokerNew("GObject", "type_add_interface_static")
	})
	return err
}

// TypeAddInterfaceStatic is a representation of the C type g_type_add_interface_static.
func TypeAddInterfaceStatic(instanceType int64, interfaceType int64, info *InterfaceInfo) {
	var inArgs [3]gi.Argument
	inArgs[0].SetInt64(instanceType)
	inArgs[1].SetInt64(interfaceType)
	inArgs[2].SetPointer(info.Native())

	err := typeAddInterfaceStaticFunction_Set()
	if err == nil {
		typeAddInterfaceStaticFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeCheckClassCastFunction *gi.Function
var typeCheckClassCastFunction_Once sync.Once

func typeCheckClassCastFunction_Set() error {
	var err error
	typeCheckClassCastFunction_Once.Do(func() {
		typeCheckClassCastFunction, err = gi.FunctionInvokerNew("GObject", "type_check_class_cast")
	})
	return err
}

// TypeCheckClassCast is a representation of the C type g_type_check_class_cast.
func TypeCheckClassCast(gClass *TypeClass, isAType int64) *TypeClass {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(gClass.Native())
	inArgs[1].SetInt64(isAType)

	var ret gi.Argument

	err := typeCheckClassCastFunction_Set()
	if err == nil {
		ret = typeCheckClassCastFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeClassNewFromNative(ret.Pointer())

	return retGo
}

var typeCheckClassIsAFunction *gi.Function
var typeCheckClassIsAFunction_Once sync.Once

func typeCheckClassIsAFunction_Set() error {
	var err error
	typeCheckClassIsAFunction_Once.Do(func() {
		typeCheckClassIsAFunction, err = gi.FunctionInvokerNew("GObject", "type_check_class_is_a")
	})
	return err
}

// TypeCheckClassIsA is a representation of the C type g_type_check_class_is_a.
func TypeCheckClassIsA(gClass *TypeClass, isAType int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(gClass.Native())
	inArgs[1].SetInt64(isAType)

	var ret gi.Argument

	err := typeCheckClassIsAFunction_Set()
	if err == nil {
		ret = typeCheckClassIsAFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var typeCheckInstanceFunction *gi.Function
var typeCheckInstanceFunction_Once sync.Once

func typeCheckInstanceFunction_Set() error {
	var err error
	typeCheckInstanceFunction_Once.Do(func() {
		typeCheckInstanceFunction, err = gi.FunctionInvokerNew("GObject", "type_check_instance")
	})
	return err
}

// TypeCheckInstance is a representation of the C type g_type_check_instance.
func TypeCheckInstance(instance *TypeInstance) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(instance.Native())

	var ret gi.Argument

	err := typeCheckInstanceFunction_Set()
	if err == nil {
		ret = typeCheckInstanceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var typeCheckInstanceCastFunction *gi.Function
var typeCheckInstanceCastFunction_Once sync.Once

func typeCheckInstanceCastFunction_Set() error {
	var err error
	typeCheckInstanceCastFunction_Once.Do(func() {
		typeCheckInstanceCastFunction, err = gi.FunctionInvokerNew("GObject", "type_check_instance_cast")
	})
	return err
}

// TypeCheckInstanceCast is a representation of the C type g_type_check_instance_cast.
func TypeCheckInstanceCast(instance *TypeInstance, ifaceType int64) *TypeInstance {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetInt64(ifaceType)

	var ret gi.Argument

	err := typeCheckInstanceCastFunction_Set()
	if err == nil {
		ret = typeCheckInstanceCastFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeInstanceNewFromNative(ret.Pointer())

	return retGo
}

var typeCheckInstanceIsAFunction *gi.Function
var typeCheckInstanceIsAFunction_Once sync.Once

func typeCheckInstanceIsAFunction_Set() error {
	var err error
	typeCheckInstanceIsAFunction_Once.Do(func() {
		typeCheckInstanceIsAFunction, err = gi.FunctionInvokerNew("GObject", "type_check_instance_is_a")
	})
	return err
}

// TypeCheckInstanceIsA is a representation of the C type g_type_check_instance_is_a.
func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetInt64(ifaceType)

	var ret gi.Argument

	err := typeCheckInstanceIsAFunction_Set()
	if err == nil {
		ret = typeCheckInstanceIsAFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var typeCheckInstanceIsFundamentallyAFunction *gi.Function
var typeCheckInstanceIsFundamentallyAFunction_Once sync.Once

func typeCheckInstanceIsFundamentallyAFunction_Set() error {
	var err error
	typeCheckInstanceIsFundamentallyAFunction_Once.Do(func() {
		typeCheckInstanceIsFundamentallyAFunction, err = gi.FunctionInvokerNew("GObject", "type_check_instance_is_fundamentally_a")
	})
	return err
}

// TypeCheckInstanceIsFundamentallyA is a representation of the C type g_type_check_instance_is_fundamentally_a.
func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(instance.Native())
	inArgs[1].SetInt64(fundamentalType)

	var ret gi.Argument

	err := typeCheckInstanceIsFundamentallyAFunction_Set()
	if err == nil {
		ret = typeCheckInstanceIsFundamentallyAFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var typeCheckIsValueTypeFunction *gi.Function
var typeCheckIsValueTypeFunction_Once sync.Once

func typeCheckIsValueTypeFunction_Set() error {
	var err error
	typeCheckIsValueTypeFunction_Once.Do(func() {
		typeCheckIsValueTypeFunction, err = gi.FunctionInvokerNew("GObject", "type_check_is_value_type")
	})
	return err
}

// TypeCheckIsValueType is a representation of the C type g_type_check_is_value_type.
func TypeCheckIsValueType(type_ int64) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeCheckIsValueTypeFunction_Set()
	if err == nil {
		ret = typeCheckIsValueTypeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var typeCheckValueFunction *gi.Function
var typeCheckValueFunction_Once sync.Once

func typeCheckValueFunction_Set() error {
	var err error
	typeCheckValueFunction_Once.Do(func() {
		typeCheckValueFunction, err = gi.FunctionInvokerNew("GObject", "type_check_value")
	})
	return err
}

// TypeCheckValue is a representation of the C type g_type_check_value.
func TypeCheckValue(value *Value) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(value.Native())

	var ret gi.Argument

	err := typeCheckValueFunction_Set()
	if err == nil {
		ret = typeCheckValueFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var typeCheckValueHoldsFunction *gi.Function
var typeCheckValueHoldsFunction_Once sync.Once

func typeCheckValueHoldsFunction_Set() error {
	var err error
	typeCheckValueHoldsFunction_Once.Do(func() {
		typeCheckValueHoldsFunction, err = gi.FunctionInvokerNew("GObject", "type_check_value_holds")
	})
	return err
}

// TypeCheckValueHolds is a representation of the C type g_type_check_value_holds.
func TypeCheckValueHolds(value *Value, type_ int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(value.Native())
	inArgs[1].SetInt64(type_)

	var ret gi.Argument

	err := typeCheckValueHoldsFunction_Set()
	if err == nil {
		ret = typeCheckValueHoldsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var typeChildrenFunction *gi.Function
var typeChildrenFunction_Once sync.Once

func typeChildrenFunction_Set() error {
	var err error
	typeChildrenFunction_Once.Do(func() {
		typeChildrenFunction, err = gi.FunctionInvokerNew("GObject", "type_children")
	})
	return err
}

// TypeChildren is a representation of the C type g_type_children.
func TypeChildren(type_ int64) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var outArgs [1]gi.Argument

	err := typeChildrenFunction_Set()
	if err == nil {
		typeChildrenFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
}

var typeClassAdjustPrivateOffsetFunction *gi.Function
var typeClassAdjustPrivateOffsetFunction_Once sync.Once

func typeClassAdjustPrivateOffsetFunction_Set() error {
	var err error
	typeClassAdjustPrivateOffsetFunction_Once.Do(func() {
		typeClassAdjustPrivateOffsetFunction, err = gi.FunctionInvokerNew("GObject", "type_class_adjust_private_offset")
	})
	return err
}

// TypeClassAdjustPrivateOffset is a representation of the C type g_type_class_adjust_private_offset.
func TypeClassAdjustPrivateOffset(gClass unsafe.Pointer, privateSizeOrOffset int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(gClass)
	inArgs[1].SetInt32(privateSizeOrOffset)

	err := typeClassAdjustPrivateOffsetFunction_Set()
	if err == nil {
		typeClassAdjustPrivateOffsetFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeClassPeekFunction *gi.Function
var typeClassPeekFunction_Once sync.Once

func typeClassPeekFunction_Set() error {
	var err error
	typeClassPeekFunction_Once.Do(func() {
		typeClassPeekFunction, err = gi.FunctionInvokerNew("GObject", "type_class_peek")
	})
	return err
}

// TypeClassPeek is a representation of the C type g_type_class_peek.
func TypeClassPeek(type_ int64) *TypeClass {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeClassPeekFunction_Set()
	if err == nil {
		ret = typeClassPeekFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeClassNewFromNative(ret.Pointer())

	return retGo
}

var typeClassPeekStaticFunction *gi.Function
var typeClassPeekStaticFunction_Once sync.Once

func typeClassPeekStaticFunction_Set() error {
	var err error
	typeClassPeekStaticFunction_Once.Do(func() {
		typeClassPeekStaticFunction, err = gi.FunctionInvokerNew("GObject", "type_class_peek_static")
	})
	return err
}

// TypeClassPeekStatic is a representation of the C type g_type_class_peek_static.
func TypeClassPeekStatic(type_ int64) *TypeClass {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeClassPeekStaticFunction_Set()
	if err == nil {
		ret = typeClassPeekStaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeClassNewFromNative(ret.Pointer())

	return retGo
}

var typeClassRefFunction *gi.Function
var typeClassRefFunction_Once sync.Once

func typeClassRefFunction_Set() error {
	var err error
	typeClassRefFunction_Once.Do(func() {
		typeClassRefFunction, err = gi.FunctionInvokerNew("GObject", "type_class_ref")
	})
	return err
}

// TypeClassRef is a representation of the C type g_type_class_ref.
func TypeClassRef(type_ int64) *TypeClass {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeClassRefFunction_Set()
	if err == nil {
		ret = typeClassRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeClassNewFromNative(ret.Pointer())

	return retGo
}

var typeCreateInstanceFunction *gi.Function
var typeCreateInstanceFunction_Once sync.Once

func typeCreateInstanceFunction_Set() error {
	var err error
	typeCreateInstanceFunction_Once.Do(func() {
		typeCreateInstanceFunction, err = gi.FunctionInvokerNew("GObject", "type_create_instance")
	})
	return err
}

// TypeCreateInstance is a representation of the C type g_type_create_instance.
func TypeCreateInstance(type_ int64) *TypeInstance {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeCreateInstanceFunction_Set()
	if err == nil {
		ret = typeCreateInstanceFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeInstanceNewFromNative(ret.Pointer())

	return retGo
}

var typeDefaultInterfacePeekFunction *gi.Function
var typeDefaultInterfacePeekFunction_Once sync.Once

func typeDefaultInterfacePeekFunction_Set() error {
	var err error
	typeDefaultInterfacePeekFunction_Once.Do(func() {
		typeDefaultInterfacePeekFunction, err = gi.FunctionInvokerNew("GObject", "type_default_interface_peek")
	})
	return err
}

// TypeDefaultInterfacePeek is a representation of the C type g_type_default_interface_peek.
func TypeDefaultInterfacePeek(gType int64) *TypeInterface {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(gType)

	var ret gi.Argument

	err := typeDefaultInterfacePeekFunction_Set()
	if err == nil {
		ret = typeDefaultInterfacePeekFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeInterfaceNewFromNative(ret.Pointer())

	return retGo
}

var typeDefaultInterfaceRefFunction *gi.Function
var typeDefaultInterfaceRefFunction_Once sync.Once

func typeDefaultInterfaceRefFunction_Set() error {
	var err error
	typeDefaultInterfaceRefFunction_Once.Do(func() {
		typeDefaultInterfaceRefFunction, err = gi.FunctionInvokerNew("GObject", "type_default_interface_ref")
	})
	return err
}

// TypeDefaultInterfaceRef is a representation of the C type g_type_default_interface_ref.
func TypeDefaultInterfaceRef(gType int64) *TypeInterface {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(gType)

	var ret gi.Argument

	err := typeDefaultInterfaceRefFunction_Set()
	if err == nil {
		ret = typeDefaultInterfaceRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeInterfaceNewFromNative(ret.Pointer())

	return retGo
}

var typeDefaultInterfaceUnrefFunction *gi.Function
var typeDefaultInterfaceUnrefFunction_Once sync.Once

func typeDefaultInterfaceUnrefFunction_Set() error {
	var err error
	typeDefaultInterfaceUnrefFunction_Once.Do(func() {
		typeDefaultInterfaceUnrefFunction, err = gi.FunctionInvokerNew("GObject", "type_default_interface_unref")
	})
	return err
}

// TypeDefaultInterfaceUnref is a representation of the C type g_type_default_interface_unref.
func TypeDefaultInterfaceUnref(gIface *TypeInterface) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(gIface.Native())

	err := typeDefaultInterfaceUnrefFunction_Set()
	if err == nil {
		typeDefaultInterfaceUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeDepthFunction *gi.Function
var typeDepthFunction_Once sync.Once

func typeDepthFunction_Set() error {
	var err error
	typeDepthFunction_Once.Do(func() {
		typeDepthFunction, err = gi.FunctionInvokerNew("GObject", "type_depth")
	})
	return err
}

// TypeDepth is a representation of the C type g_type_depth.
func TypeDepth(type_ int64) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeDepthFunction_Set()
	if err == nil {
		ret = typeDepthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var typeEnsureFunction *gi.Function
var typeEnsureFunction_Once sync.Once

func typeEnsureFunction_Set() error {
	var err error
	typeEnsureFunction_Once.Do(func() {
		typeEnsureFunction, err = gi.FunctionInvokerNew("GObject", "type_ensure")
	})
	return err
}

// TypeEnsure is a representation of the C type g_type_ensure.
func TypeEnsure(type_ int64) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	err := typeEnsureFunction_Set()
	if err == nil {
		typeEnsureFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeFreeInstanceFunction *gi.Function
var typeFreeInstanceFunction_Once sync.Once

func typeFreeInstanceFunction_Set() error {
	var err error
	typeFreeInstanceFunction_Once.Do(func() {
		typeFreeInstanceFunction, err = gi.FunctionInvokerNew("GObject", "type_free_instance")
	})
	return err
}

// TypeFreeInstance is a representation of the C type g_type_free_instance.
func TypeFreeInstance(instance *TypeInstance) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(instance.Native())

	err := typeFreeInstanceFunction_Set()
	if err == nil {
		typeFreeInstanceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeFromNameFunction *gi.Function
var typeFromNameFunction_Once sync.Once

func typeFromNameFunction_Set() error {
	var err error
	typeFromNameFunction_Once.Do(func() {
		typeFromNameFunction, err = gi.FunctionInvokerNew("GObject", "type_from_name")
	})
	return err
}

// TypeFromName is a representation of the C type g_type_from_name.
func TypeFromName(name string) int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(name)

	var ret gi.Argument

	err := typeFromNameFunction_Set()
	if err == nil {
		ret = typeFromNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var typeFundamentalFunction *gi.Function
var typeFundamentalFunction_Once sync.Once

func typeFundamentalFunction_Set() error {
	var err error
	typeFundamentalFunction_Once.Do(func() {
		typeFundamentalFunction, err = gi.FunctionInvokerNew("GObject", "type_fundamental")
	})
	return err
}

// TypeFundamental is a representation of the C type g_type_fundamental.
func TypeFundamental(typeId int64) int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(typeId)

	var ret gi.Argument

	err := typeFundamentalFunction_Set()
	if err == nil {
		ret = typeFundamentalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var typeFundamentalNextFunction *gi.Function
var typeFundamentalNextFunction_Once sync.Once

func typeFundamentalNextFunction_Set() error {
	var err error
	typeFundamentalNextFunction_Once.Do(func() {
		typeFundamentalNextFunction, err = gi.FunctionInvokerNew("GObject", "type_fundamental_next")
	})
	return err
}

// TypeFundamentalNext is a representation of the C type g_type_fundamental_next.
func TypeFundamentalNext() int64 {

	var ret gi.Argument

	err := typeFundamentalNextFunction_Set()
	if err == nil {
		ret = typeFundamentalNextFunction.Invoke(nil, nil)
	}

	retGo := ret.Int64()

	return retGo
}

var typeGetInstanceCountFunction *gi.Function
var typeGetInstanceCountFunction_Once sync.Once

func typeGetInstanceCountFunction_Set() error {
	var err error
	typeGetInstanceCountFunction_Once.Do(func() {
		typeGetInstanceCountFunction, err = gi.FunctionInvokerNew("GObject", "type_get_instance_count")
	})
	return err
}

// TypeGetInstanceCount is a representation of the C type g_type_get_instance_count.
func TypeGetInstanceCount(type_ int64) int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeGetInstanceCountFunction_Set()
	if err == nil {
		ret = typeGetInstanceCountFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'g_type_get_plugin' : return type 'TypePlugin' not supported

var typeGetQdataFunction *gi.Function
var typeGetQdataFunction_Once sync.Once

func typeGetQdataFunction_Set() error {
	var err error
	typeGetQdataFunction_Once.Do(func() {
		typeGetQdataFunction, err = gi.FunctionInvokerNew("GObject", "type_get_qdata")
	})
	return err
}

// TypeGetQdata is a representation of the C type g_type_get_qdata.
func TypeGetQdata(type_ int64, quark glib.Quark) unsafe.Pointer {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(type_)
	inArgs[1].SetUint32(uint32(quark))

	var ret gi.Argument

	err := typeGetQdataFunction_Set()
	if err == nil {
		ret = typeGetQdataFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Pointer()

	return retGo
}

var typeGetTypeRegistrationSerialFunction *gi.Function
var typeGetTypeRegistrationSerialFunction_Once sync.Once

func typeGetTypeRegistrationSerialFunction_Set() error {
	var err error
	typeGetTypeRegistrationSerialFunction_Once.Do(func() {
		typeGetTypeRegistrationSerialFunction, err = gi.FunctionInvokerNew("GObject", "type_get_type_registration_serial")
	})
	return err
}

// TypeGetTypeRegistrationSerial is a representation of the C type g_type_get_type_registration_serial.
func TypeGetTypeRegistrationSerial() uint32 {

	var ret gi.Argument

	err := typeGetTypeRegistrationSerialFunction_Set()
	if err == nil {
		ret = typeGetTypeRegistrationSerialFunction.Invoke(nil, nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var typeInitFunction *gi.Function
var typeInitFunction_Once sync.Once

func typeInitFunction_Set() error {
	var err error
	typeInitFunction_Once.Do(func() {
		typeInitFunction, err = gi.FunctionInvokerNew("GObject", "type_init")
	})
	return err
}

// TypeInit is a representation of the C type g_type_init.
func TypeInit() {

	err := typeInitFunction_Set()
	if err == nil {
		typeInitFunction.Invoke(nil, nil)
	}

	return
}

var typeInitWithDebugFlagsFunction *gi.Function
var typeInitWithDebugFlagsFunction_Once sync.Once

func typeInitWithDebugFlagsFunction_Set() error {
	var err error
	typeInitWithDebugFlagsFunction_Once.Do(func() {
		typeInitWithDebugFlagsFunction, err = gi.FunctionInvokerNew("GObject", "type_init_with_debug_flags")
	})
	return err
}

// TypeInitWithDebugFlags is a representation of the C type g_type_init_with_debug_flags.
func TypeInitWithDebugFlags(debugFlags TypeDebugFlags) {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt32(int32(debugFlags))

	err := typeInitWithDebugFlagsFunction_Set()
	if err == nil {
		typeInitWithDebugFlagsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeInterfaceAddPrerequisiteFunction *gi.Function
var typeInterfaceAddPrerequisiteFunction_Once sync.Once

func typeInterfaceAddPrerequisiteFunction_Set() error {
	var err error
	typeInterfaceAddPrerequisiteFunction_Once.Do(func() {
		typeInterfaceAddPrerequisiteFunction, err = gi.FunctionInvokerNew("GObject", "type_interface_add_prerequisite")
	})
	return err
}

// TypeInterfaceAddPrerequisite is a representation of the C type g_type_interface_add_prerequisite.
func TypeInterfaceAddPrerequisite(interfaceType int64, prerequisiteType int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(interfaceType)
	inArgs[1].SetInt64(prerequisiteType)

	err := typeInterfaceAddPrerequisiteFunction_Set()
	if err == nil {
		typeInterfaceAddPrerequisiteFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'g_type_interface_get_plugin' : return type 'TypePlugin' not supported

var typeInterfacePeekFunction *gi.Function
var typeInterfacePeekFunction_Once sync.Once

func typeInterfacePeekFunction_Set() error {
	var err error
	typeInterfacePeekFunction_Once.Do(func() {
		typeInterfacePeekFunction, err = gi.FunctionInvokerNew("GObject", "type_interface_peek")
	})
	return err
}

// TypeInterfacePeek is a representation of the C type g_type_interface_peek.
func TypeInterfacePeek(instanceClass *TypeClass, ifaceType int64) *TypeInterface {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(instanceClass.Native())
	inArgs[1].SetInt64(ifaceType)

	var ret gi.Argument

	err := typeInterfacePeekFunction_Set()
	if err == nil {
		ret = typeInterfacePeekFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeInterfaceNewFromNative(ret.Pointer())

	return retGo
}

var typeInterfacePrerequisitesFunction *gi.Function
var typeInterfacePrerequisitesFunction_Once sync.Once

func typeInterfacePrerequisitesFunction_Set() error {
	var err error
	typeInterfacePrerequisitesFunction_Once.Do(func() {
		typeInterfacePrerequisitesFunction, err = gi.FunctionInvokerNew("GObject", "type_interface_prerequisites")
	})
	return err
}

// TypeInterfacePrerequisites is a representation of the C type g_type_interface_prerequisites.
func TypeInterfacePrerequisites(interfaceType int64) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(interfaceType)

	var outArgs [1]gi.Argument

	err := typeInterfacePrerequisitesFunction_Set()
	if err == nil {
		typeInterfacePrerequisitesFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
}

var typeInterfacesFunction *gi.Function
var typeInterfacesFunction_Once sync.Once

func typeInterfacesFunction_Set() error {
	var err error
	typeInterfacesFunction_Once.Do(func() {
		typeInterfacesFunction, err = gi.FunctionInvokerNew("GObject", "type_interfaces")
	})
	return err
}

// TypeInterfaces is a representation of the C type g_type_interfaces.
func TypeInterfaces(type_ int64) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var outArgs [1]gi.Argument

	err := typeInterfacesFunction_Set()
	if err == nil {
		typeInterfacesFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()

	return out0
}

var typeIsAFunction *gi.Function
var typeIsAFunction_Once sync.Once

func typeIsAFunction_Set() error {
	var err error
	typeIsAFunction_Once.Do(func() {
		typeIsAFunction, err = gi.FunctionInvokerNew("GObject", "type_is_a")
	})
	return err
}

// TypeIsA is a representation of the C type g_type_is_a.
func TypeIsA(type_ int64, isAType int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(type_)
	inArgs[1].SetInt64(isAType)

	var ret gi.Argument

	err := typeIsAFunction_Set()
	if err == nil {
		ret = typeIsAFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var typeNameFunction *gi.Function
var typeNameFunction_Once sync.Once

func typeNameFunction_Set() error {
	var err error
	typeNameFunction_Once.Do(func() {
		typeNameFunction, err = gi.FunctionInvokerNew("GObject", "type_name")
	})
	return err
}

// TypeName is a representation of the C type g_type_name.
func TypeName(type_ int64) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeNameFunction_Set()
	if err == nil {
		ret = typeNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var typeNameFromClassFunction *gi.Function
var typeNameFromClassFunction_Once sync.Once

func typeNameFromClassFunction_Set() error {
	var err error
	typeNameFromClassFunction_Once.Do(func() {
		typeNameFromClassFunction, err = gi.FunctionInvokerNew("GObject", "type_name_from_class")
	})
	return err
}

// TypeNameFromClass is a representation of the C type g_type_name_from_class.
func TypeNameFromClass(gClass *TypeClass) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(gClass.Native())

	var ret gi.Argument

	err := typeNameFromClassFunction_Set()
	if err == nil {
		ret = typeNameFromClassFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var typeNameFromInstanceFunction *gi.Function
var typeNameFromInstanceFunction_Once sync.Once

func typeNameFromInstanceFunction_Set() error {
	var err error
	typeNameFromInstanceFunction_Once.Do(func() {
		typeNameFromInstanceFunction, err = gi.FunctionInvokerNew("GObject", "type_name_from_instance")
	})
	return err
}

// TypeNameFromInstance is a representation of the C type g_type_name_from_instance.
func TypeNameFromInstance(instance *TypeInstance) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(instance.Native())

	var ret gi.Argument

	err := typeNameFromInstanceFunction_Set()
	if err == nil {
		ret = typeNameFromInstanceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var typeNextBaseFunction *gi.Function
var typeNextBaseFunction_Once sync.Once

func typeNextBaseFunction_Set() error {
	var err error
	typeNextBaseFunction_Once.Do(func() {
		typeNextBaseFunction, err = gi.FunctionInvokerNew("GObject", "type_next_base")
	})
	return err
}

// TypeNextBase is a representation of the C type g_type_next_base.
func TypeNextBase(leafType int64, rootType int64) int64 {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(leafType)
	inArgs[1].SetInt64(rootType)

	var ret gi.Argument

	err := typeNextBaseFunction_Set()
	if err == nil {
		ret = typeNextBaseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var typeParentFunction *gi.Function
var typeParentFunction_Once sync.Once

func typeParentFunction_Set() error {
	var err error
	typeParentFunction_Once.Do(func() {
		typeParentFunction, err = gi.FunctionInvokerNew("GObject", "type_parent")
	})
	return err
}

// TypeParent is a representation of the C type g_type_parent.
func TypeParent(type_ int64) int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeParentFunction_Set()
	if err == nil {
		ret = typeParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var typeQnameFunction *gi.Function
var typeQnameFunction_Once sync.Once

func typeQnameFunction_Set() error {
	var err error
	typeQnameFunction_Once.Do(func() {
		typeQnameFunction, err = gi.FunctionInvokerNew("GObject", "type_qname")
	})
	return err
}

// TypeQname is a representation of the C type g_type_qname.
func TypeQname(type_ int64) glib.Quark {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeQnameFunction_Set()
	if err == nil {
		ret = typeQnameFunction.Invoke(inArgs[:], nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

var typeQueryFunction *gi.Function
var typeQueryFunction_Once sync.Once

func typeQueryFunction_Set() error {
	var err error
	typeQueryFunction_Once.Do(func() {
		typeQueryFunction, err = gi.FunctionInvokerNew("GObject", "type_query")
	})
	return err
}

// TypeQuery__ is a representation of the C type g_type_query.
func TypeQuery__(type_ int64) *TypeQuery {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var outArgs [1]gi.Argument

	err := typeQueryFunction_Set()
	if err == nil {
		typeQueryFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := TypeQueryNewFromNative(outArgs[0].Pointer())

	return out0
}

// UNSUPPORTED : C value 'g_type_register_dynamic' : parameter 'plugin' of type 'TypePlugin' not supported

var typeRegisterFundamentalFunction *gi.Function
var typeRegisterFundamentalFunction_Once sync.Once

func typeRegisterFundamentalFunction_Set() error {
	var err error
	typeRegisterFundamentalFunction_Once.Do(func() {
		typeRegisterFundamentalFunction, err = gi.FunctionInvokerNew("GObject", "type_register_fundamental")
	})
	return err
}

// TypeRegisterFundamental is a representation of the C type g_type_register_fundamental.
func TypeRegisterFundamental(typeId int64, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags TypeFlags) int64 {
	var inArgs [5]gi.Argument
	inArgs[0].SetInt64(typeId)
	inArgs[1].SetString(typeName)
	inArgs[2].SetPointer(info.Native())
	inArgs[3].SetPointer(finfo.Native())
	inArgs[4].SetInt32(int32(flags))

	var ret gi.Argument

	err := typeRegisterFundamentalFunction_Set()
	if err == nil {
		ret = typeRegisterFundamentalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var typeRegisterStaticFunction *gi.Function
var typeRegisterStaticFunction_Once sync.Once

func typeRegisterStaticFunction_Set() error {
	var err error
	typeRegisterStaticFunction_Once.Do(func() {
		typeRegisterStaticFunction, err = gi.FunctionInvokerNew("GObject", "type_register_static")
	})
	return err
}

// TypeRegisterStatic is a representation of the C type g_type_register_static.
func TypeRegisterStatic(parentType int64, typeName string, info *TypeInfo, flags TypeFlags) int64 {
	var inArgs [4]gi.Argument
	inArgs[0].SetInt64(parentType)
	inArgs[1].SetString(typeName)
	inArgs[2].SetPointer(info.Native())
	inArgs[3].SetInt32(int32(flags))

	var ret gi.Argument

	err := typeRegisterStaticFunction_Set()
	if err == nil {
		ret = typeRegisterStaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

// UNSUPPORTED : C value 'g_type_register_static_simple' : parameter 'class_init' of type 'ClassInitFunc' not supported

// UNSUPPORTED : C value 'g_type_remove_class_cache_func' : parameter 'cache_func' of type 'TypeClassCacheFunc' not supported

// UNSUPPORTED : C value 'g_type_remove_interface_check' : parameter 'check_func' of type 'TypeInterfaceCheckFunc' not supported

var typeSetQdataFunction *gi.Function
var typeSetQdataFunction_Once sync.Once

func typeSetQdataFunction_Set() error {
	var err error
	typeSetQdataFunction_Once.Do(func() {
		typeSetQdataFunction, err = gi.FunctionInvokerNew("GObject", "type_set_qdata")
	})
	return err
}

// TypeSetQdata is a representation of the C type g_type_set_qdata.
func TypeSetQdata(type_ int64, quark glib.Quark, data unsafe.Pointer) {
	var inArgs [3]gi.Argument
	inArgs[0].SetInt64(type_)
	inArgs[1].SetUint32(uint32(quark))
	inArgs[2].SetPointer(data)

	err := typeSetQdataFunction_Set()
	if err == nil {
		typeSetQdataFunction.Invoke(inArgs[:], nil)
	}

	return
}

var typeTestFlagsFunction *gi.Function
var typeTestFlagsFunction_Once sync.Once

func typeTestFlagsFunction_Set() error {
	var err error
	typeTestFlagsFunction_Once.Do(func() {
		typeTestFlagsFunction, err = gi.FunctionInvokerNew("GObject", "type_test_flags")
	})
	return err
}

// TypeTestFlags is a representation of the C type g_type_test_flags.
func TypeTestFlags(type_ int64, flags uint32) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(type_)
	inArgs[1].SetUint32(flags)

	var ret gi.Argument

	err := typeTestFlagsFunction_Set()
	if err == nil {
		ret = typeTestFlagsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var typeValueTablePeekFunction *gi.Function
var typeValueTablePeekFunction_Once sync.Once

func typeValueTablePeekFunction_Set() error {
	var err error
	typeValueTablePeekFunction_Once.Do(func() {
		typeValueTablePeekFunction, err = gi.FunctionInvokerNew("GObject", "type_value_table_peek")
	})
	return err
}

// TypeValueTablePeek is a representation of the C type g_type_value_table_peek.
func TypeValueTablePeek(type_ int64) *TypeValueTable {
	var inArgs [1]gi.Argument
	inArgs[0].SetInt64(type_)

	var ret gi.Argument

	err := typeValueTablePeekFunction_Set()
	if err == nil {
		ret = typeValueTablePeekFunction.Invoke(inArgs[:], nil)
	}

	retGo := TypeValueTableNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'g_value_register_transform_func' : parameter 'transform_func' of type 'ValueTransform' not supported

var valueTypeCompatibleFunction *gi.Function
var valueTypeCompatibleFunction_Once sync.Once

func valueTypeCompatibleFunction_Set() error {
	var err error
	valueTypeCompatibleFunction_Once.Do(func() {
		valueTypeCompatibleFunction, err = gi.FunctionInvokerNew("GObject", "value_type_compatible")
	})
	return err
}

// ValueTypeCompatible is a representation of the C type g_value_type_compatible.
func ValueTypeCompatible(srcType int64, destType int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(srcType)
	inArgs[1].SetInt64(destType)

	var ret gi.Argument

	err := valueTypeCompatibleFunction_Set()
	if err == nil {
		ret = valueTypeCompatibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var valueTypeTransformableFunction *gi.Function
var valueTypeTransformableFunction_Once sync.Once

func valueTypeTransformableFunction_Set() error {
	var err error
	valueTypeTransformableFunction_Once.Do(func() {
		valueTypeTransformableFunction, err = gi.FunctionInvokerNew("GObject", "value_type_transformable")
	})
	return err
}

// ValueTypeTransformable is a representation of the C type g_value_type_transformable.
func ValueTypeTransformable(srcType int64, destType int64) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt64(srcType)
	inArgs[1].SetInt64(destType)

	var ret gi.Argument

	err := valueTypeTransformableFunction_Set()
	if err == nil {
		ret = valueTypeTransformableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var badFunction *gi.Function
var badFunction_Once sync.Once

func badFunction_Set() error {
	var err error
	badFunction_Once.Do(func() {
		badFunction, err = gi.FunctionInvokerNew("GObject", "bad")
	})
	return err
}

// Bad is a representation of the C type g_bad.
func Bad() {

	err := badFunction_Set()
	if err == nil {
		badFunction.Invoke(nil, nil)
	}

	return
}
