#ifndef BRIGHTRAY_EXAMPLE_BROWSER_MAC_WINDOW_CONTROLLER_H_
#define BRIGHTRAY_EXAMPLE_BROWSER_MAC_WINDOW_CONTROLLER_H_

#import "base/memory/scoped_ptr.h"
#import <Cocoa/Cocoa.h>

namespace gallium {
class BrowserContext;
class WindowMac;
}

@interface WindowController : NSWindowController {
 @private
  scoped_ptr<gallium::WindowMac> wrapper_window_;
}

@property (nonatomic, readonly, assign) gallium::WindowMac* wrapperWindow;

- (instancetype)initWithBrowserContext:(gallium::BrowserContext*)browserContext;

@end

#endif
