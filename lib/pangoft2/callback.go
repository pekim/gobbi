// This is a generated file - DO NOT EDIT

package pangoft2

import (
	fontconfig "github.com/pekim/gobbi/lib/fontconfig"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <pango/pangoft2.h>
// #include <stdlib.h>
/*

	void callback_substitutefuncHandler(GObject *, FcPattern*, gpointer);

*/
import "C"

var callbackSubstitutefuncId int
var callbackSubstitutefuncMap = make(map[int]SubstitutefuncCallback)
var callbackSubstitutefuncLock sync.RWMutex

// SubstitutefuncCallback is a callback function for a 'SubstituteFunc' callback.
type SubstitutefuncCallback func(pattern *fontconfig.Pattern)

//export callback_substitutefuncHandler
func callback_substitutefuncHandler(_ *C.GObject, c_pattern *C.FcPattern, c_data C.gpointer) {
	callbackSubstitutefuncLock.RLock()
	defer callbackSubstitutefuncLock.RUnlock()

	pattern := fontconfig.PatternNewFromC(unsafe.Pointer(c_pattern))

	index := int(uintptr(c_data))
	callback := callbackSubstitutefuncMap[index]
	callback(pattern)
}
