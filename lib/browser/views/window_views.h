#ifndef BRIGHTRAY_EXAMPLE_BROWSER_VIEWS_WINDOW_VIEWS_H_
#define BRIGHTRAY_EXAMPLE_BROWSER_VIEWS_WINDOW_VIEWS_H_

#include "browser/window.h"

namespace views {
class Widget;
}

namespace gallium {

class WindowViews : public Window {
 public:
  WindowViews(gallium::BrowserContext*);
  ~WindowViews();

  void Show() override;

 private:
  views::Widget* widget_;

  DISALLOW_COPY_AND_ASSIGN(WindowViews);
};

}

#endif
