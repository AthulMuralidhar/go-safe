#include <gtkmm.h>

// run command: g++ main.cpp -o main `pkg-config --cflags --libs gtkmm-3.0` -std=c++17   && ./main
// main thing to note - this is only gtkmm3.0 compatible
// things change in gtk4.0 so watch out! 


class MyWindow : public Gtk::Window
{
public:
  MyWindow();
};

MyWindow::MyWindow()
{
  set_title("Basic application");
  set_default_size(200, 200);
}

int main(int argc, char *argv[])
{
  Glib::RefPtr<Gtk::Application> app = Gtk::Application::create(argc, argv, "org.gtkmm.examples.base");

  MyWindow window; // Create an instance of MyWindow

  // Shows the window and returns when it is closed.
  return app->run(window);
}
