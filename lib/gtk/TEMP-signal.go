package gtk

import (
	"github.com/pekim/gobbi/internal/cgo/callback"
	"github.com/pekim/gobbi/lib/gobject"
	"unsafe"
)

func (recv *Widget) ConnectShow(handler func(widget *Widget)) int {
	return callback.ConnectSignal(recv.Native(), "show",
		func(return_value *callback.Value, paramValues []callback.Value) {
			value0 := gobject.ValueNewFromNative(unsafe.Pointer(&paramValues[0]))
			widget := WidgetNewFromNative(value0.GetObject().Native())

			handler(widget)
		},
	)
}

func (recv *Widget) DisconnectSignal(id int) {
	callback.DisconnectSignal(id)
}
