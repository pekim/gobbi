package sc

/*
#include <gtk/gtk.h>

gboolean DrawingAreaDraw2(GtkDrawingArea *widget, cairo_t *cr);

gboolean drawing_area_vf_draw2(GtkDrawingArea *widget, cairo_t *cr) {
	return DrawingAreaDraw2(widget, cr);
}

void drawing_area_class_init2(GtkDrawingAreaClass *g_class, gpointer class_data) {
	GtkWidgetClass *widget_class;
	widget_class = (GtkWidgetClass*) g_class;

	widget_class->draw = drawing_area_vf_draw2;
}

*/
import "C"
