#include <vector>
#include <string>
#include <pthread.h>

#include "base/bind.h"
#include "common/gallium.h"
#include "common/main_delegate.h"
#include "browser/window.h"
#include "browser/browser_client.h"

#include "content/public/app/content_main.h"
#include "content/public/browser/browser_thread.h"

std::vector<std::string> args;

void AddArg(const char* arg) {
  args.push_back(strdup(arg));
}

int RunGallium() {
  const char** argv = new const char*[args.size()];
  for (size_t i = 0; i < args.size(); i++) {
    argv[i] = args[i].c_str();
  }
  brightray_example::MainDelegate delegate;
  content::ContentMainParams params(&delegate);
  params.argc = args.size();
  params.argv = argv;
  return content::ContentMain(params);
}

std::unique_ptr<brightray_example::MainDelegate> delegate;

int GalliumLoop(const char* argv0, struct gallium_error** err) {
  uint64_t tid;
  pthread_threadid_np(NULL, &tid);
  printf("in GalliumLoop, thread=%llu\n", tid);
  fflush(stdout);

  delegate.reset(new brightray_example::MainDelegate);
  content::ContentMainParams params(delegate.get());

  const char* argv[] = {argv0};
  params.argc = 1;
  params.argv = argv;

  return content::ContentMain(params);
}

void DoCreateWindow(const char* title) {
  printf("At DoCreateWindow\n");
  auto window = brightray_example::Window::Create(delegate->get_browser_client()->browser_context());
  window->Show();
}

// GalliumCreateWindow creates a window pointed at the given url
struct gallium_window* GalliumCreateWindow(const char* title, struct gallium_error** err) {
  printf("in GalliumCreateWindow\n");
  content::BrowserThread::PostTask(content::BrowserThread::UI, FROM_HERE, base::Bind(&DoCreateWindow, title));
  return NULL;
}
