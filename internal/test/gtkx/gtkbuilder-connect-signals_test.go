package gtkx

import (
	"github.com/pekim/gobbi/lib/gtk"
	"github.com/pekim/gobbi/lib/gtkx"
	"os"
	"testing"
)

const ui = `
<interface>
  <object class="GtkBox" id="box1">
    <property name="orientation">vertical</property>
    <property name="spacing">20</property>
	<child>
	  <object class="GtkButton" id="ok_button">
		<property name="label">gtk-ok</property>
		<property name="use-stock">TRUE</property>
		<signal name="clicked" handler="ok_button_clicked"/>
	  </object>
	</child>
	<child>
	  <object class="GtkButton" id="cancel_button">
		<property name="label">gtk-cancel</property>
		<property name="use-stock">TRUE</property>
		<signal name="clicked" handler="cancel_button_clicked"/>
	  </object>
	</child>
  </object>
</interface>
`

func TestBuildConnectSignals(t *testing.T) {
	gtk.Init(os.Args)
	builder := gtk.BuilderNewFromString(ui)

	gtkx.BuilderConnectSignals(builder, map[string]interface{}{
		"ok_button_clicked": func() {},
		"bad_name":          func(i int) {},
	})
}
