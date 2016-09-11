#include "renderer/renderer_client.h"

#include "renderer/render_view_observer.h"

namespace gallium {

RendererClient::RendererClient() {
}

RendererClient::~RendererClient() {
}

void RendererClient::RenderViewCreated(content::RenderView* render_view) {
  auto observer = new RenderViewObserver(render_view);
  // observer will be deleted automatically when render_view is destroyed.
  (void)observer;
}

}
