// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE-CHROMIUM file.

#include "common/main_delegate.h"

#include "browser/browser_client.h"
#include "renderer/renderer_client.h"
#include "common/content_client.h"

#include "base/command_line.h"
#include "base/path_service.h"
#include "content/public/common/content_switches.h"
#include "ui/base/resource/resource_bundle.h"

namespace gallium {

MainDelegate::MainDelegate() {
}

MainDelegate::~MainDelegate() {
}

BrowserClient* MainDelegate::browser_client() {
  return browser_client_.get();
}

bool MainDelegate::BasicStartupComplete(int* exit_code) {
  printf("MainDelegate::BasicStartupComplete...\n"); fflush(stdout);
  content_client_.reset(new ContentClient);
  SetContentClient(content_client_.get());
#if defined(OS_MACOSX)
  printf("going to OverrideFrameworkBundlePath...\n"); fflush(stdout);
  OverrideFrameworkBundlePath();
#endif
  return false;
}

void MainDelegate::PreSandboxStartup() {
  printf("MainDelegate::PreSandboxStartup\n"); fflush(stdout);
  InitializeResourceBundle();
}

void MainDelegate::InitializeResourceBundle() {
  base::FilePath path;
#if defined(OS_MACOSX)
  path = GetResourcesPakFilePath();
#else
  base::FilePath pak_dir;
  PathService::Get(base::DIR_MODULE, &pak_dir);
  path = pak_dir.Append(FILE_PATH_LITERAL("content_shell.pak"));
#endif

  ui::ResourceBundle::InitSharedInstanceWithPakPath(path);
  AddDataPackFromPath(&ui::ResourceBundle::GetSharedInstance(), path.DirName());
}

content::ContentBrowserClient* MainDelegate::CreateContentBrowserClient() {
  browser_client_.reset(new BrowserClient);
  return browser_client_.get();
}

content::ContentRendererClient* MainDelegate::CreateContentRendererClient() {
  renderer_client_.reset(new RendererClient);
  return renderer_client_.get();
}

}  // namespace gallium
