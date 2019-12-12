package callback

// #cgo pkg-config: gobject-2.0
//
// #include <stdlib.h>
// #include "signal.h"
import "C"

import (
	"fmt"
	"github.com/pekim/gobbi/internal/cgo"
	"sync"
	"unsafe"
)

const PackageName = "github.com/pekim/gobbi/internal/cgo/callback"

type Value C.GValue

type Handler func(return_value *Value, paramValues []Value)

type signalConnection struct {
	cConnection C.GobbiSignalConnection
	signalName  string
	handler     Handler
}

var signalConnections = make(map[int]signalConnection)
var signalConnectionsId int
var signalConnectionsMutex sync.Mutex

func ConnectSignal(instance unsafe.Pointer, signalName string, handler Handler) int {
	signalConnectionsMutex.Lock()
	defer signalConnectionsMutex.Unlock()

	id := signalConnectionsId
	signalConnectionsId++

	cSignalName := C.CString(signalName)
	defer C.free(unsafe.Pointer(cSignalName))

	cConnection := C.gobbi_connect_signal(instance, cSignalName, unsafe.Pointer(uintptr(id)))

	connection := signalConnection{
		cConnection: cConnection,
		signalName:  signalName,
		handler:     handler,
	}

	signalConnections[id] = connection

	return id
}

func DisconnectSignal(id int) {
	signalConnectionsMutex.Lock()
	defer signalConnectionsMutex.Unlock()

	connection := signalConnections[id]
	C.gobbi_disconnect_signal(connection.cConnection)

	delete(signalConnections, id)
}

//export gobbiCSignalHandler
func gobbiCSignalHandler(cReturnValue *C.GValue, nParamValues C.guint, cParamValues *C.GValue, data unsafe.Pointer) {
	id := int(uintptr(data))
	connection := signalConnections[id]

	returnValue := (*Value)(cReturnValue)
	paramValues := (*[1 << 30]Value)(unsafe.Pointer(cParamValues))[:nParamValues:nParamValues]

	if cgo.Tracing() {
		traceCall(connection)
	}

	connection.handler(returnValue, paramValues)

	if cgo.Tracing() {
		traceReturn(returnValue, paramValues)
	}
}

func traceCall(connection signalConnection) {
	cgo.Trace(fmt.Sprintf("Signal : %s\n", connection.signalName))
}

// traceReturn includes param values that might be out parameters.
// So it must be called after the handler has returned.
func traceReturn(returnValue *Value, paramValues []Value) {
	ret := ""
	if returnValue != nil {
		ret = fmt.Sprintf("   ret %v\n", *returnValue)
	}

	args := ""
	if len(paramValues) > 0 {
		args = fmt.Sprintf("  args %v\n", paramValues)
	}

	cgo.Trace(fmt.Sprintf("Signal : %s%s\n", ret, args))
}
