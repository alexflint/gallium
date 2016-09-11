#include "browser/toplevel_browser_main_parts.h"

#include "browser/window.h"

namespace brightray_example {

BrowserMainParts::BrowserMainParts() {
}

BrowserMainParts::~BrowserMainParts() {
}

void BrowserMainParts::PreMainMessageLoopRun() {
  brightray::BrowserMainParts::PreMainMessageLoopRun();

  //auto window = Window::Create(browser_context());
  //window->Show();
}

}
