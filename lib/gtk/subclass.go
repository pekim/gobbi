package gtk

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>
#include <stdlib.h>

gboolean widget_vf_draw(GtkEntry *widget, cairo_t *cr) {
	printf("C draw\n");
	return TRUE;
}

void entry_vf_move_cursor(GtkEntry       *entry,
			       GtkMovementStep       step,
			       gint                  count,
			       gboolean              extend_selection) {
	printf("C move cursor\n");
}

void
entry_vf_realize        (GtkWidget        *widget)
{
	gtk_widget_set_realized(widget, TRUE);

  GdkWindowAttr attributes;
  gint attributes_mask;

  //GtkWidget *widget;
  //widget = (GtkWidget*) widget;

  attributes.window_type = GDK_WINDOW_CHILD;
  //attributes.x = widget->allocation.x;
  //attributes.y = widget->allocation.y;
  //attributes.width = widget->allocation.width;
  //attributes.height = widget->allocation.height;
  attributes.x = 20;
  attributes.y = 20;
  attributes.width = 50;
  attributes.height = 25;
  attributes.wclass = GDK_INPUT_OUTPUT;
  attributes.visual = gtk_widget_get_visual (widget);
  //attributes.colormap = gtk_widget_get_colormap (widget);
  attributes.event_mask = gtk_widget_get_events (widget) | GDK_EXPOSURE_MASK;

  attributes_mask = GDK_WA_X | GDK_WA_Y | GDK_WA_VISUAL;

  //widget->window = gdk_window_new (gtk_widget_get_parent_window (widget),
  //                                 &attributes, attributes_mask);
  gtk_widget_set_window(widget, gdk_window_new (gtk_widget_get_parent_window (widget),
                                   &attributes, attributes_mask));

  printf("realize\n");
  //gdk_window_set_user_data (gtk_widget_get_window(widget), widget);
}

void widget_class_init(GtkEntryClass *g_class, gpointer class_data) {
	//g_class.draw = widget_vf_draw;

	GtkWidgetClass *widget_class;
	widget_class = (GtkWidgetClass*) g_class;
	widget_class->realize = entry_vf_realize;

	g_class->move_cursor = entry_vf_move_cursor;

	printf("class init\n");
}

*/
import "C"

import (
	"fmt"
	"github.com/pekim/gobbi/lib/cairo"
	"unsafe"
)

type EntryVirtualDraw interface {
	Draw(cr *cairo.Context) bool
}

//type EntryVirtualFunctions struct {
//	draw EntryDrawFunc
//}

type EntryDerivedClass struct {
	name  string
	gtype C.GType
}

type EntryDerived struct {
	class  *EntryDerivedClass
	native *C.GtkEntry
}

func EntryDerive(name string) *EntryDerivedClass {
	var typeInfo C.GTypeInfo
	typeInfo.class_size = C.sizeof_GtkEntryClass
	typeInfo.instance_size = C.sizeof_GtkEntry
	typeInfo.class_init = C.GClassInitFunc(C.widget_class_init)

	cTypeName := C.CString(name)
	defer C.free(unsafe.Pointer(cTypeName))

	gtype := C.g_type_register_static(C.GTK_TYPE_ENTRY, cTypeName, &typeInfo, 0)

	class := &EntryDerivedClass{
		name:  name,
		gtype: gtype,
	}

	return class
}

func (c *EntryDerivedClass) New(virtualFunctions interface{}) *EntryDerived {
	f, ok := virtualFunctions.(EntryVirtualDraw)
	fmt.Println("draw func :", ok)
	f.Draw(nil)

	native := (*C.GtkEntry)(C.g_object_newv(c.gtype, 0, nil))

	instance := &EntryDerived{
		class:  c,
		native: native,
	}

	return instance
}

// Entry upcasts to *Entry
func (recv *EntryDerived) Entry() *Entry {
	return EntryNewFromC(unsafe.Pointer(recv.native))

}
