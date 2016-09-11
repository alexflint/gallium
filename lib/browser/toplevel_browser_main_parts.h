#ifndef BRIGHTRAY_EXAMPLE_BROWSER_BROWSER_MAIN_PARTS_H_
#define BRIGHTRAY_EXAMPLE_BROWSER_BROWSER_MAIN_PARTS_H_

#include "browser/browser_main_parts.h"

namespace brightray_example {

class BrowserMainParts : public brightray::BrowserMainParts {
public:
  BrowserMainParts();
  ~BrowserMainParts();

protected:
  void PreMainMessageLoopRun() override;

  DISALLOW_COPY_AND_ASSIGN(BrowserMainParts);
};

}

#endif
