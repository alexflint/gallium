#import "browser/mac/window_controller.h"

#import "browser/mac/window_mac.h"

#import "brightray/browser/inspectable_web_contents.h"
#import "brightray/browser/inspectable_web_contents_view.h"

@implementation WindowController

- (instancetype)initWithBrowserContext:(brightray::BrowserContext *)browserContext {
  self = [super initWithWindowNibName:@"WindowController"];
  if (!self)
    return nil;

  wrapper_window_.reset(new brightray_example::WindowMac(browserContext, self));

  return self;
}

- (void)windowDidLoad {
  [super windowDidLoad];

  auto contentsView = self.wrapperWindow->inspectable_web_contents()->GetView()->GetNativeView();

  contentsView.frame = [self.window.contentView bounds];
  contentsView.autoresizingMask = NSViewWidthSizable | NSViewHeightSizable;

  [self.window.contentView addSubview:contentsView];
  self.window.title = @"Window title!";

  self.wrapperWindow->WindowReady();
}

- (brightray_example::WindowMac*)wrapperWindow {
  return wrapper_window_.get();
}

- (void)windowWillClose:(NSNotification *)notification {
  [self performSelector:@selector(autorelease) withObject:nil afterDelay:0];
}

@end
