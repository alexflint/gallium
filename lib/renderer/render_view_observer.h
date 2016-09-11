#ifndef BRIGHTRAY_EXAMPLE_RENDERER_RENDER_VIEW_OBSERVER_H_
#define BRIGHTRAY_EXAMPLE_RENDERER_RENDER_VIEW_OBSERVER_H_

#include "content/public/renderer/render_view_observer.h"

namespace brightray {

class RenderViewObserver : content::RenderViewObserver {
public:
  explicit RenderViewObserver(content::RenderView*);
private:
  ~RenderViewObserver();

  void DidClearWindowObject(blink::WebLocalFrame* frame) override;
};

}

#endif
