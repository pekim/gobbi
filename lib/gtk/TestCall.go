package gtk

import (
	"fmt"
	"github.com/pekim/gobbi/lib/pango"
)

func TestCall() {
	//d := call.Data{
	//	ReturnType: call.RT_INT,
	//}
	//call.Function(2460, &d)

	fmt.Println(pango.Version())
}
