// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE-CHROMIUM file.

#ifndef BRIGHTRAY_COMMON_MAIN_DELEGATE_H_
#define BRIGHTRAY_COMMON_MAIN_DELEGATE_H_

#include "base/compiler_specific.h"
#include "base/memory/scoped_ptr.h"
#include "content/public/app/content_main_delegate.h"

namespace base {
class FilePath;
}

namespace ui {
class ResourceBundle;
}

namespace gallium {

class BrowserClient;
class ContentClient;
class RendererClient;

class MainDelegate : public content::ContentMainDelegate {
 public:
  MainDelegate();
  ~MainDelegate();

  BrowserClient* browser_client();

 protected:
  // Subclasses can override this to provide additional .pak files to be
  // included in the ui::ResourceBundle.
  virtual void AddDataPackFromPath(
      ui::ResourceBundle* bundle, const base::FilePath& pak_dir) {}

#if defined(OS_MACOSX)
  // Subclasses can override this to custom the paths of child process and
  // framework bundle.
  virtual base::FilePath GetResourcesPakFilePath();
  virtual void OverrideChildProcessPath();
  virtual void OverrideFrameworkBundlePath();
#endif

  virtual bool BasicStartupComplete(int* exit_code) override;
  virtual void PreSandboxStartup() override;

 private:
  virtual content::ContentBrowserClient* CreateContentBrowserClient() override;
  virtual content::ContentRendererClient* CreateContentRendererClient() override;

  void InitializeResourceBundle();

  scoped_ptr<ContentClient> content_client_;
  scoped_ptr<BrowserClient> browser_client_;
  scoped_ptr<RendererClient> renderer_client_;

  DISALLOW_COPY_AND_ASSIGN(MainDelegate);
};

}  // namespace gallium
#endif
