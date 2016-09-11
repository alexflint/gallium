#import "browser/default_web_contents_delegate.h"

#import "content/public/browser/native_web_keyboard_event.h"
#import <AppKit/AppKit.h>

namespace gallium {

void DefaultWebContentsDelegate::HandleKeyboardEvent(content::WebContents*, const content::NativeWebKeyboardEvent& event) {
  if (event.skip_in_browser)
    return;

  [[NSApp mainMenu] performKeyEquivalent:event.os_event];
}

}
