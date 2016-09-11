#ifndef BRIGHTRAY_EXAMPLE_BROWSER_MAC_WINDOW_MAC_H_
#define BRIGHTRAY_EXAMPLE_BROWSER_MAC_WINDOW_MAC_H_

#include "browser/window.h"

@class WindowController;

namespace gallium {
  
class WindowMac : public Window {
 public:
  WindowMac(gallium::BrowserContext*, WindowController*);
  ~WindowMac();

  void Show() override;

 private:
  // Owns us.
  WindowController* controller_;

  DISALLOW_COPY_AND_ASSIGN(WindowMac);
};

}

#endif
