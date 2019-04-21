package gtk

import (
	"github.com/pekim/gobbi/lib/internal/call"
)

func TestCall() {
	d := call.Data{}
	call.Function(2460, d)
}
