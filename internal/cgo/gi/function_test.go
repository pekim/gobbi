package gi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionIntegerReturn(t *testing.T) {
	Require("GLib", "2.0")
	fn, _ := Function2InvokerNew("GLib", "get_num_processors")

	ret := fn.Invoke([]Arg{}, 0, 0, Arg{typ: ArgType_uint})
	assert.True(t, ret.value.(uint) > 0)
}
