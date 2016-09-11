#include "app/library_main.h"
#include "common/main_delegate.h"

#include "content/public/app/content_main.h"
#include "sandbox/win/src/sandbox_types.h"

#if defined(OS_WIN)

#include "content/public/app/startup_helper_win.h"
#include <Windows.h>

int WINAPI wWinMain(HINSTANCE instance, HINSTANCE, wchar_t*, int) {
  sandbox::SandboxInterfaceInfo sandbox_info = {0};
  content::InitializeSandboxInfo(&sandbox_info);
  brightray::MainDelegate delegate;

  content::ContentMainParams params(&delegate);
  params.instance = instance;
  params.sandbox_info = &sandbox_info;
  return content::ContentMain(params);
}

#elif defined(OS_MACOSX)

int main(int argc, const char* argv[]) {
  BrightrayExampleMain(argc, argv);
}

#else // OS_LINUX

int main(int argc, const char* argv[]) {
  brightray::MainDelegate delegate;
  content::ContentMainParams params(&delegate);
  params.argc = argc;
  params.argv = argv;
  return content::ContentMain(params);
}

#endif
