#ifndef _obj_sc_h_
#define _obj_sc_h_

#include <glib-object.h>
#include <gtk/gtk.h>

G_BEGIN_DECLS

#define MY_APP_TYPE_WINDOW (my_app_window_get_type ())
G_DECLARE_FINAL_TYPE(MyAppWindow, my_app_window, MY_APP, WINDOW, GtkApplicationWindow)

MyAppWindow *          my_app_window_new            (void);

G_END_DECLS

#endif /* _obj_sch_ */
