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

func TestBuildConnectSignals_SignalsOfClass(t *testing.T) {
	gtk.Init(os.Args)
	builder := gtk.BuilderNewFromString(`
		<interface>
		  <object class="GtkButton" id="ok_button">
			<signal name="clicked" handler="ok_button_clicked"/>
			<signal name="enter" handler="ok_button_enter"/>
		  </object>
		</interface>
	`)

	err := gtkx.BuilderConnectSignals(builder, map[string]interface{}{
		"ok_button_clicked": func() {},
		"ok_button_enter":   func() {},
	})

	assert.Nil(t, err)
}

func TestBuildConnectSignals_SignalOfAncestor(t *testing.T) {
	gtk.Init(os.Args)
	builder := gtk.BuilderNewFromString(`
		<interface>
		  <object class="GtkButton" id="ok_button">
			<signal name="draw" handler="draw"/>
		  </object>
		</interface>
	`)

	err := gtkx.BuilderConnectSignals(builder, map[string]interface{}{
		"draw": func(cr *cairo.Context) bool { return false },
	})

	assert.Nil(t, err)
}

func TestBuildConnectSignals_SignalOfAncestorInterface(t *testing.T) {
	gtk.Init(os.Args)
	builder := gtk.BuilderNewFromString(`
		<interface>
		  <object class="GtkSpinButton" id="spin_button">
			<signal name="changed" handler="spinbutton_changed"/>
		  </object>
		</interface>
	`)

	err := gtkx.BuilderConnectSignals(builder, map[string]interface{}{
		"spinbutton_changed": func() {},
	})

	assert.Nil(t, err)
}

func TestBuildConnectSignals_BadHandlerName(t *testing.T) {
	gtk.Init(os.Args)
	builder := gtk.BuilderNewFromString(`
		<interface>
		  <object class="GtkButton" id="ok_button">
			<signal name="clicked" handler="ok_button_clicked"/>
		  </object>
		</interface>
	`)

	err := gtkx.BuilderConnectSignals(builder, map[string]interface{}{
		"bad_handler_name": func() {},
	})

	assert.NotNil(t, err)
}

func TestBuildConnectSignals_BadHandlerSignature(t *testing.T) {
	gtk.Init(os.Args)
	builder := gtk.BuilderNewFromString(`
		<interface>
		  <object class="GtkButton" id="ok_button">
			<signal name="clicked" handler="ok_button_clicked"/>
		  </object>
		</interface>
	`)

	err := gtkx.BuilderConnectSignals(builder, map[string]interface{}{
		"ok_button_clicked": func(bad string) {},
	})

	assert.NotNil(t, err)
}

func TestBuildConnectSignals_NotifyProperty(t *testing.T) {
	signalHandled := false

	gtk.Init(os.Args)
	builder := gtk.BuilderNewFromString(`
		<interface>
		  <object class="GtkLabel" id="label">
			<property name="label">one</property>
			<signal name="notify::label" handler="label_changed"/>
		  </object>
		</interface>
	`)
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

func TestBuildConnectSignals_NotifyPropertyBadHandlerSignature(t *testing.T) {
	gtk.Init(os.Args)
	builder := gtk.BuilderNewFromString(`
		<interface>
		  <object class="GtkLabel" id="label">
			<property name="label">one</property>
			<signal name="notify::label" handler="label_changed"/>
		  </object>
		</interface>
	`)

	err := gtkx.BuilderConnectSignals(builder, map[string]interface{}{
		"label_changed": func(bad string) {},
	})
	assert.NotNil(t, err)
}
