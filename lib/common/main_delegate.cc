#include "common/main_delegate.h"

#include "browser/browser_client.h"
#include "renderer/renderer_client.h"

namespace brightray_example {

MainDelegate::MainDelegate() {
}

MainDelegate::~MainDelegate() {
}


brightray::BrowserClient* MainDelegate::get_browser_client() {
  return browser_client_.get();
}

content::ContentBrowserClient* MainDelegate::CreateContentBrowserClient() {
  browser_client_.reset(new BrowserClient);
  return browser_client_.get();
}

content::ContentRendererClient* MainDelegate::CreateContentRendererClient() {
  renderer_client_.reset(new RendererClient);
  return renderer_client_.get();
}

}
