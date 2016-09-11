#include "browser/window.h"

#include "brightray/browser/browser_context.h"
#include "brightray/browser/default_web_contents_delegate.h"
#include "brightray/browser/inspectable_web_contents.h"

namespace brightray_example {

Window::Window(brightray::BrowserContext* browser_context)
  : browser_context_(browser_context),
    inspectable_web_contents_(brightray::InspectableWebContents::Create(content::WebContents::CreateParams(browser_context))),
    web_contents_delegate_(new brightray::DefaultWebContentsDelegate) {
  auto web_contents = inspectable_web_contents_->GetWebContents();
  web_contents->SetDelegate(web_contents_delegate_.get());
}

Window::~Window() {
}

void Window::WindowReady() {
  auto web_contents = inspectable_web_contents_->GetWebContents();
  web_contents->GetController().LoadURL(GURL("http://adam.roben.org/brightray_example/start.html"), content::Referrer(), ui::PAGE_TRANSITION_AUTO_TOPLEVEL, std::string());
  web_contents->SetInitialFocus();
}

}
