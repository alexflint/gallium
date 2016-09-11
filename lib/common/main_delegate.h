#ifndef BRIGHTRAY_EXAMPLE_COMMON_MAIN_DELEGATE_H_
#define BRIGHTRAY_EXAMPLE_COMMON_MAIN_DELEGATE_H_

#include "browser/browser_client.h"
#include "renderer/renderer_client.h"

#include "brightray/common/main_delegate.h"

namespace brightray_example {

class MainDelegate : public brightray::MainDelegate {
public:
  MainDelegate();
  ~MainDelegate();

  brightray::BrowserClient* get_browser_client();

private:
  content::ContentBrowserClient* CreateContentBrowserClient() override;
  content::ContentRendererClient* CreateContentRendererClient() override;

  scoped_ptr<BrowserClient> browser_client_;
  scoped_ptr<RendererClient> renderer_client_;

  DISALLOW_COPY_AND_ASSIGN(MainDelegate);
};

}

#endif
