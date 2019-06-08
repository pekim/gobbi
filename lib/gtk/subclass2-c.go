package gtk

/*
#include <gtk/gtk.h>

gboolean DrawingAreaDraw(GtkDrawingArea *widget, cairo_t *cr);

gboolean drawing_area_vf_draw(GtkDrawingArea *widget, cairo_t *cr) {
	return DrawingAreaDraw(widget, cr);
}

void drawing_area_class_init(GtkDrawingAreaClass *g_class, gpointer class_data) {

	GtkWidgetClass *widget_class;
	widget_class = (GtkWidgetClass*) g_class;

	widget_class->draw = drawing_area_vf_draw;
}

*/
import "C"
