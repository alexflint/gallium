#include <vector>
#include <string>
#include <pthread.h>

#include "base/bind.h"

#include "api/gallium.h"
#include "common/main_delegate.h"
#include "browser/window.h"
#include "browser/browser_client.h"

#include "content/public/app/content_main.h"
#include "content/public/browser/browser_thread.h"

std::unique_ptr<gallium::MainDelegate> delegate;

int GalliumLoop(const char* argv0, struct gallium_error** err) {
  uint64_t tid;
  pthread_threadid_np(NULL, &tid);
  printf("in GalliumLoop, thread=%llu\n", tid);
  fflush(stdout);

  delegate.reset(new gallium::MainDelegate);
  content::ContentMainParams params(delegate.get());

  const char* argv[] = {argv0};
  params.argc = 1;
  params.argv = argv;

  return content::ContentMain(params);
}

// DoCreateWindow is the code that runs on the UI thread when GalliumCreateWindow is called
void DoCreateWindow(const std::string& url, const std::string& title) {
  printf("At DoCreateWindow\n");
  auto window = gallium::Window::Create(
    delegate->browser_client()->browser_context());
  window->SetInitURL(url);
  window->Show();
}

// GalliumCreateWindow creates a window pointed at the given url
struct gallium_window* GalliumCreateWindow(const char* url, const char* title, struct gallium_error** err) {
  printf("in GalliumCreateWindow\n");
  content::BrowserThread::PostTask(
    content::BrowserThread::UI,
    FROM_HERE,
    base::Bind(
      &DoCreateWindow,
      std::string(url),
      std::string(title)));
  return NULL;
}
