#include "browser/browser_client.h"

#include "browser/browser_main_parts.h"

namespace brightray_example {

BrowserClient::BrowserClient() {
}

BrowserClient::~BrowserClient() {
}

brightray::BrowserMainParts* BrowserClient::OverrideCreateBrowserMainParts(const content::MainFunctionParams&) {
  return new BrowserMainParts;
}

}
