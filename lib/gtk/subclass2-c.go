package gtk

/*
#include <gtk/gtk.h>

gboolean DrawingAreaDraw(GtkDrawingArea *widget, cairo_t *cr);

gboolean drawing_area_vf_draw(GtkDrawingArea *widget, cairo_t *cr) {
	return DrawingAreaDraw(widget, cr);
}

void DrawingAreaClassInit(GtkDrawingAreaClass *g_class, gpointer class_data);

void drawing_area_class_init(GtkDrawingAreaClass *g_class, gpointer class_data) {
	DrawingAreaClassInit(g_class, class_data);

	//GtkWidgetClass *widget_class;
	//widget_class = (GtkWidgetClass*) g_class;
	//
	//widget_class->draw = drawing_area_vf_draw;
}

*/
import "C"
