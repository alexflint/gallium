// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Copyright (c) 2013 Adam Roben <adam@roben.org>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE-CHROMIUM file.

#import "common/mac/main_application_bundle.h"

#import "common/mac/foundation_util.h"

#import "base/files/file_path.h"
#import "base/path_service.h"

// This class exists only to help us find the bundle corresponding to Gallium.framework
@interface BundleHelper : NSObject
@end

@implementation BundleHelper
@end

namespace gallium {

NSBundle* GalliumFrameworkBundle() {
  return [NSBundle bundleForClass:[BundleHelper class]];
}

base::FilePath MainApplicationBundlePath() {
  NSLog(@"in MainApplicationBundlePath");
  // Start out with the path to the running executable.
  base::FilePath path;
  PathService::Get(base::FILE_EXE, &path);

  // One step up to MacOS, another to Contents.
  path = path.DirName().DirName();
  DCHECK_EQ(path.BaseName().value(), "Contents");

  // Up one more level to the .app.
  path = path.DirName();
  DCHECK_EQ(path.BaseName().Extension(), ".app");

  return path;
}

NSBundle* MainApplicationBundle() {
  return [NSBundle bundleWithPath:base::mac::FilePathToNSString(MainApplicationBundlePath())];
}

}
