#include "obj-sc.h"

struct _MyAppWindow
{
  GtkApplicationWindow parent_instance;

  // instance variables for subclass go here
};

G_DEFINE_TYPE(MyAppWindow, my_app_window, GTK_TYPE_APPLICATION_WINDOW)

static void
my_app_window_init (MyAppWindow *window)
{
  // initialisation goes here
}

static void
my_app_window_class_init (MyAppWindowClass *class)
{
  // virtual function overrides go here
  // property and signal definitions go here
}

MyAppWindow *
my_app_window_new (void)
{
  return g_object_new (MY_APP_TYPE_WINDOW, NULL);
}
