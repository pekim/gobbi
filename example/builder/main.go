package main

import (
	"github.com/pekim/gobbi/lib/gtk"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

const ui = `
<interface>
  <object class="GtkBox" id="box1">
    <property name="orientation">vertical</property>
    <property name="spacing">20</property>
    <child>
      <object class="GtkLabel" id="label1">
        <property name="label">label 1</property>
      </object>
    </child>
    <child>
      <object class="GtkLabel" id="label2">
        <property name="label">label 2</property>
      </object>
    </child>
    <child>
      <object class="GtkLabel" id="label3">
        <property name="label">label 3</property>
      </object>
    </child>
  </object>
</interface>
`

func main() {
	gtk.Init(os.Args)

	window := gtk.WindowNew(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("builder")
	window.SetDefaultSize(200, 150)

	builder := gtk.BuilderNewFromString(ui, -1)

	label3 := gtk.CastToLabel(builder.GetObject("label3"))
	label3.SetText(label3.GetText() + " - AMENDED")

	vbox := gtk.CastToVBox(builder.GetObject("box1"))
	window.Container().Add(vbox.Widget())

	window.Widget().ConnectDestroy(gtk.MainQuit)
	window.Widget().ShowAll()

	gtk.Main()
}
