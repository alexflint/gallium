#include "renderer/render_view_observer.h"

#include "base/strings/stringprintf.h"
#include "base/strings/string_util.h"
#include "third_party/WebKit/public/web/WebDocument.h"
#include "third_party/WebKit/public/web/WebLocalFrame.h"
#include "v8/include/v8.h"

namespace gallium {

namespace {

void HelloWorld(v8::Local<v8::String> property, const v8::PropertyCallbackInfo<v8::Value>& info) {
  info.GetReturnValue().Set(v8::String::NewFromUtf8(v8::Isolate::GetCurrent(), base::StringPrintf("Hello, World from %s:%d!", __FILE__, __LINE__).c_str()));
}

v8::Local<v8::ObjectTemplate> CreateConstructorTemplate() {
  auto constructor_template = v8::ObjectTemplate::New();
  constructor_template->SetAccessor(v8::String::NewFromUtf8(v8::Isolate::GetCurrent(), "helloWorld"), HelloWorld);
  return constructor_template;
}


v8::Local<v8::ObjectTemplate> ConstructorTemplate() {
  auto isolate = v8::Isolate::GetCurrent();
  static v8::Persistent<v8::ObjectTemplate> constructor_template(isolate, CreateConstructorTemplate());
  return v8::Local<v8::ObjectTemplate>::New(isolate, constructor_template);
}

}

RenderViewObserver::RenderViewObserver(content::RenderView *render_view)
    : content::RenderViewObserver(render_view) {
}

RenderViewObserver::~RenderViewObserver() {
}

void RenderViewObserver::DidClearWindowObject(blink::WebLocalFrame* frame) {
  GURL url = frame->document().url();
  if (!url.SchemeIs("http") ||
      !url.DomainIs("adam.roben.org") ||
      !StartsWithASCII(url.path(), "/brightray_example/", true))
    return;

  v8::HandleScope scope(v8::Isolate::GetCurrent());

  auto context = frame->mainWorldScriptContext();
  v8::Context::Scope contextScope(context);

  context->Global()->Set(v8::String::NewFromUtf8(v8::Isolate::GetCurrent(), "brightray_example"), ConstructorTemplate()->NewInstance());
}

}
