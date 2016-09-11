#ifndef BRIGHTRAY_EXAMPLE_BROWSER_MAC_WINDOW_CONTROLLER_H_
#define BRIGHTRAY_EXAMPLE_BROWSER_MAC_WINDOW_CONTROLLER_H_

#import "base/memory/scoped_ptr.h"
#import <Cocoa/Cocoa.h>

namespace brightray {
class BrowserContext;
}

namespace brightray_example {
class WindowMac;
}

@interface WindowController : NSWindowController {
 @private
  scoped_ptr<brightray_example::WindowMac> wrapper_window_;
}

@property (nonatomic, readonly, assign) brightray_example::WindowMac* wrapperWindow;

- (instancetype)initWithBrowserContext:(brightray::BrowserContext*)browserContext;

@end

#endif
