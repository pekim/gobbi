package gtk

import (
	"github.com/pekim/gobbi/internal/gi"
	"unsafe"
)

// TODO remove when Init is generated.

func Init() {
	fn, _ := gi.FunctionInvokerNew("Gtk", "init")

	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(unsafe.Pointer(uintptr(0)))
	inArgs[1].SetPointer(unsafe.Pointer(uintptr(0)))

	var outArgs [2]gi.Argument

	fn.Invoke(inArgs[:], outArgs[:])
}
