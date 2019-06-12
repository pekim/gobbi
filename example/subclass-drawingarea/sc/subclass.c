#include <gtk/gtk.h>
#include "_cgo_export.h"

// virtual function implementation
gboolean drawing_area_vf_draw(GtkWidget *widget, cairo_t *cr) {
    // call Go function
	return DrawingAreaDraw(widget, cr);
}

void drawing_area_class_init(GtkDrawingAreaClass *g_class, gpointer class_data) {
	GtkWidgetClass *widget_class;
	widget_class = (GtkWidgetClass*) g_class;

    // override virtual functions
	widget_class->draw = drawing_area_vf_draw;
}
