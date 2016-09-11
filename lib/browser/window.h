#ifndef BRIGHTRAY_EXAMPLE_BROWSER_WINDOW_H_
#define BRIGHTRAY_EXAMPLE_BROWSER_WINDOW_H_

#include "base/basictypes.h"
#include "base/memory/scoped_ptr.h"

namespace brightray {
class BrowserContext;
class DefaultWebContentsDelegate;
class InspectableWebContents;
}

namespace brightray_example {

class Window {
 public:
  static Window* Create(brightray::BrowserContext*);

  brightray::BrowserContext* browser_context() { return browser_context_; }
  brightray::InspectableWebContents* inspectable_web_contents() { return inspectable_web_contents_.get(); }
  
  void WindowReady();
  virtual void Show() = 0;

 protected:
  Window(brightray::BrowserContext*);
  virtual ~Window();

 private:
  brightray::BrowserContext* browser_context_;
  scoped_ptr<brightray::InspectableWebContents> inspectable_web_contents_;
  scoped_ptr<brightray::DefaultWebContentsDelegate> web_contents_delegate_;

  DISALLOW_COPY_AND_ASSIGN(Window);
};

}

#endif
