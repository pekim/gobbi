package cpuwidget

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
import "C"

import "fmt"

func CpuWidgetTest() {
	fmt.Println("cwt")
}
