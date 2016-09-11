#ifndef BRIGHTRAY_EXAMPLE_BROWSER_WINDOW_H_
#define BRIGHTRAY_EXAMPLE_BROWSER_WINDOW_H_

#include <string>

#include "base/basictypes.h"
#include "base/memory/scoped_ptr.h"

namespace gallium {
class BrowserContext;
class DefaultWebContentsDelegate;
class InspectableWebContents;

class Window {
 public:
  static Window* Create(gallium::BrowserContext* ctx);

  gallium::BrowserContext* browser_context() { return browser_context_; }
  gallium::InspectableWebContents* inspectable_web_contents() { return inspectable_web_contents_.get(); }
  
  void SetInitURL(const std::string& url);

  void WindowReady();
  virtual void Show() = 0;

 protected:
  Window(gallium::BrowserContext*);
  virtual ~Window();

 private:
  std::string init_url_;
  gallium::BrowserContext* browser_context_;
  scoped_ptr<gallium::InspectableWebContents> inspectable_web_contents_;
  scoped_ptr<gallium::DefaultWebContentsDelegate> web_contents_delegate_;

  DISALLOW_COPY_AND_ASSIGN(Window);
};

}

#endif
