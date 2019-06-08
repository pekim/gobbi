package gtk

/*
#include <gtk/gtk.h>
#include <stdlib.h>

gboolean drawing_area_vf_draw(GtkDrawingArea *widget, cairo_t *cr);

gboolean drawing_area_vf_draw(GtkDrawingArea *widget, cairo_t *cr) {
	printf("C draw\n");
	DrawingAreaDraw(widget, cr);
	return FALSE;
}

void drawing_area_class_init(GtkDrawingAreaClass *g_class, gpointer class_data) {
	//g_class.draw = widget_vf_draw;

	GtkWidgetClass *widget_class;
	widget_class = (GtkWidgetClass*) g_class;
	//widget_class->realize = entry_vf_realize;
	//g_class->move_cursor = entry_vf_move_cursor;
	widget_class->draw = drawing_area_vf_draw;

	printf("class init\n");
}

*/
import "C"
