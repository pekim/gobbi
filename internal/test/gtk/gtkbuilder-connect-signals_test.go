package gtk

import (
	"github.com/pekim/gobbi/lib/cairo"
	"github.com/pekim/gobbi/lib/gobject"
	"github.com/pekim/gobbi/lib/gtk"
	"github.com/pekim/gobbi/lib/gtkx"
	"github.com/stretchr/testify/assert"
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
		<signal name="draw" handler="draw"/>
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

	err := gtkx.BuilderConnectSignals(builder, map[string]interface{}{
		"ok_button_clicked":     func(i int) {},
		"cancel_button_clicked": func(s string) {},
		"bad_name":              func() {},
		"draw":                  func(cr *cairo.Context) bool { return false },
	})

	assert.NotNil(t, err)
}

func TestBuildConnectNotifySignal(t *testing.T) {
	gtk.Init(os.Args)

	const ui = `
<interface>
  <object class="GtkLabel" id="label">
	<property name="label">one</property>
	<signal name="notify::label" handler="label_changed"/>
  </object>
</interface>
`
	signalHandled := false

	builder := gtk.BuilderNewFromString(ui)
	label := gtk.CastToLabel(builder.GetObject("label"))

	err := gtkx.BuilderConnectSignals(builder, map[string]interface{}{
		"label_changed": func(pspec *gobject.ParamSpec) {
			signalHandled = true
			assert.Equal(t, label.GetLabel(), "two")
		},
	})
	assert.Nil(t, err)

	label.SetText("two")

	assert.True(t, signalHandled)
}
