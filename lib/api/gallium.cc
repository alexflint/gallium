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
  printf("in GalliumLoop\n"); fflush(stdout);

  delegate.reset(new gallium::MainDelegate);
  content::ContentMainParams params(delegate.get());

  const char* argv[] = {
    argv0,
    "--single-process",
    "--enable-logging=stderr",
    "--v=1",
  };
  params.argc = 4;
  params.argv = argv;

  printf("calling content::ContentMain...\n"); fflush(stdout);
  int ret = content::ContentMain(params);
  printf("content::ContentMain returned %d\n", ret); fflush(stdout);
  return ret;
}

// DoCreateWindow is the code that runs on the UI thread when GalliumCreateWindow is called
void DoCreateWindow(const std::string& url, const std::string& title) {
  printf("at DoCreateWindow\n");
  auto window = gallium::Window::Create(
    delegate->browser_client()->browser_context());
  window->SetInitURL(url);
  window->Show();
}

// GalliumCreateWindow creates a window pointed at the given url
struct gallium_window* GalliumCreateWindow(const char* url, const char* title, struct gallium_error** err) {
  printf("at GalliumCreateWindow\n");
  content::BrowserThread::PostTask(
    content::BrowserThread::UI,
    FROM_HERE,
    base::Bind(
      &DoCreateWindow,
      std::string(url),
      std::string(title)));
  return NULL;
}
