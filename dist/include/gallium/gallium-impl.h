#ifndef GALLIUM_API_GALLIUM_IMPL_H_
#define GALLIUM_API_GALLIUM_IMPL_H_

#include "gallium/gallium.h"

#include "browser/default_web_contents_delegate.h"
#include "content/public/browser/web_contents.h"

typedef struct gallium_view {
  content::WebContents* web_contents;
  gallium::DefaultWebContentsDelegate* delegate;
} gallium_view_t;

#endif