{
  'variables': {
    'project_name': 'gallium',
    'product_name': 'Gallium',
    'source_root': '<!(["python", "tools/source_root.py"])',
    'app_sources': [
      'app/win/brightray_example.rc',
      'app/win/resource.h',
      'app/main.cc',
    ],
    'lib_sources': [
      'api/gallium.cc',
      'api/gallium.h',

      'browser/mac/window_controller.mm',
      'browser/mac/window_controller.h',
      'browser/mac/window_mac.h',
      'browser/mac/window_mac.mm',
      'browser/views/window_views.cc',
      'browser/views/window_views.h',
      'browser/linux/application_info_linux.cc',
      'browser/window.cc',
      'browser/window.h',
      'renderer/render_view_observer.cc',
      'renderer/render_view_observer.h',
      'renderer/renderer_client.cc',
      'renderer/renderer_client.h',

      'browser/paths.h',
      'browser/browser_client.cc',
      'browser/browser_client.h',
      'browser/browser_context.cc',
      'browser/browser_context.h',
      'browser/browser_main_parts.cc',
      'browser/browser_main_parts.h',
      'browser/browser_main_parts_mac.mm',
      'browser/default_web_contents_delegate.cc',
      'browser/default_web_contents_delegate.h',
      'browser/default_web_contents_delegate_mac.mm',
      'browser/devtools_contents_resizing_strategy.cc',
      'browser/devtools_contents_resizing_strategy.h',
      'browser/devtools_embedder_message_dispatcher.cc',
      'browser/devtools_embedder_message_dispatcher.h',
      'browser/devtools_manager_delegate.cc',
      'browser/devtools_manager_delegate.h',
      'browser/devtools_ui.cc',
      'browser/devtools_ui.h',
      'browser/inspectable_web_contents.cc',
      'browser/inspectable_web_contents.h',
      'browser/inspectable_web_contents_delegate.cc',
      'browser/inspectable_web_contents_delegate.h',
      'browser/inspectable_web_contents_impl.cc',
      'browser/inspectable_web_contents_impl.h',
      'browser/inspectable_web_contents_view.h',
      'browser/inspectable_web_contents_view_mac.h',
      'browser/inspectable_web_contents_view_mac.mm',
      'browser/mac/bry_application.h',
      'browser/mac/bry_application.mm',
      'browser/mac/bry_inspectable_web_contents_view.h',
      'browser/mac/bry_inspectable_web_contents_view.mm',
      'browser/media/media_capture_devices_dispatcher.cc',
      'browser/media/media_capture_devices_dispatcher.h',
      'browser/media/media_stream_devices_controller.cc',
      'browser/media/media_stream_devices_controller.h',
      'browser/network_delegate.cc',
      'browser/network_delegate.h',
      'browser/notification_presenter.h',
      'browser/notification_presenter_mac.h',
      'browser/notification_presenter_mac.mm',
      'browser/linux/notification_presenter_linux.h',
      'browser/linux/notification_presenter_linux.cc',
      'browser/remote_debugging_server.cc',
      'browser/remote_debugging_server.h',
      'browser/url_request_context_getter.cc',
      'browser/url_request_context_getter.h',
      'browser/views/inspectable_web_contents_view_views.h',
      'browser/views/inspectable_web_contents_view_views.cc',
      'browser/views/views_delegate.cc',
      'browser/views/views_delegate.h',
      'browser/web_ui_controller_factory.cc',
      'browser/web_ui_controller_factory.h',
      'common/application_info.h',
      'common/application_info_mac.mm',
      'common/application_info_win.cc',
      'common/content_client.cc',
      'common/content_client.h',
      'common/mac/foundation_util.h',
      'common/mac/main_application_bundle.h',
      'common/mac/main_application_bundle.mm',
      'common/main_delegate.cc',
      'common/main_delegate.h',
      'common/main_delegate_mac.mm',
    ],
    'framework_sources': [
      'app/library_main.cc',
      'app/library_main.h',
      'api/gallium.cc',
      'api/gallium.h',
    ],
    'conditions': [
      ['OS=="win"', {
        'app_sources': [
          '<(libchromiumcontent_src_dir)/content/app/startup_helper_win.cc',
        ],
      }],
    ],
  },
  'includes': [
    'gallium.gypi',
  ],
  'targets': [
    {
      'target_name': '<(project_name)',
      'type': 'executable',
      'dependencies': [
        '<(project_name)_lib',
      ],
      'sources': [
        '<@(app_sources)',
      ],
      'conditions': [
        ['OS=="mac"', {
          'product_name': '<(product_name)',
          'mac_bundle': 1,
          'dependencies!': [
            '<(project_name)_lib',
          ],
          'dependencies': [
            '<(project_name)_framework',
            '<(project_name)_helper',
          ],
          'xcode_settings': {
            'INFOPLIST_FILE': 'browser/mac/Info.plist',
            'LD_RUNPATH_SEARCH_PATHS': '@executable_path/../Frameworks',
          },
          'copies': [
            {
              'destination': '<(PRODUCT_DIR)/<(product_name).app/Contents/Frameworks',
              'files': [
                '<(PRODUCT_DIR)/<(product_name) Helper.app',
                '<(PRODUCT_DIR)/<(product_name).framework',
              ],
            },
          ],
          # 'postbuilds': [
          #   {
          #     # This postbuid step is responsible for creating the following
          #     # helpers:
          #     #
          #     # <(product_name) EH.app and <(product_name) NP.app are created
          #     # from <(product_name).app.
          #     #
          #     # The EH helper is marked for an executable heap. The NP helper
          #     # is marked for no PIE (ASLR).
          #     'postbuild_name': 'Make More Helpers',
          #     'action': [
          #       'vendor/brightray/tools/mac/make_more_helpers.sh',
          #       'Frameworks',
          #       '<(product_name)',
          #     ],
          #   },
          # ]
        }],
        ['OS=="win"', {
          'copies': [
            {
              'destination': '<(PRODUCT_DIR)',
              'files': [
                '<(libchromiumcontent_library_dir)/chromiumcontent.dll',
                '<(libchromiumcontent_library_dir)/content_shell.pak',
                '<(libchromiumcontent_library_dir)/icudtl.dat',
                '<(libchromiumcontent_library_dir)/libGLESv2.dll',
              ],
            },
          ],
        }],
      ],
    },
    {
      'target_name': '<(project_name)_lib',
      'type': 'static_library',
      'sources': [
        '<@(lib_sources)',
      ],
      'include_dirs': [
        '.',
        '<(libchromiumcontent_include_dir)',
        '<(libchromiumcontent_include_dir)/skia/config',
        '<(libchromiumcontent_include_dir)/third_party/skia/include/core',
        '<(libchromiumcontent_include_dir)/third_party/WebKit',
        '<(libchromiumcontent_library_dir)/gen',
      ],
      'direct_dependent_settings': {
        'include_dirs': [
          '.',
          '..',
          '<(libchromiumcontent_include_dir)',
          '<(libchromiumcontent_include_dir)/skia/config',
          '<(libchromiumcontent_include_dir)/third_party/skia/include/core',
          '<(libchromiumcontent_include_dir)/third_party/icu/source/common',
          '<(libchromiumcontent_include_dir)/third_party/WebKit',
          '<(libchromiumcontent_library_dir)/gen',
        ],
      },
      'conditions': [
        ['OS!="linux"', {
          'sources/': [
            ['exclude', '/linux/'],
            ['exclude', '_linux\.(cc|h)$'],
          ],
        }],
        ['OS!="mac"', {
          'sources/': [
            ['exclude', '/mac/'],
            ['exclude', '_mac\.(mm|h)$'],
          ],
        },{
          'sources/': [
            ['exclude', '/views/'],
            ['exclude', '_views\.(cc|h)$'],
          ],
        }],
        ['OS!="win"', {
          'sources/': [
            ['exclude', '/win/'],
            ['exclude', '_win\.(cc|h)$'],
          ],
        }],

        ['OS=="linux"', {
          'cflags_cc': [
            '-Wno-deprecated-register',
            '-fno-rtti',
          ],
          'link_settings': {
            'ldflags': [
              '<!@(pkg-config --libs-only-L --libs-only-other gtk+-2.0 libnotify dbus-1 x11 xrandr xext gconf-2.0)',
            ],
            'libraries': [
              '<(source_root)/<(libchromiumcontent_library_dir)/libchromiumcontent.so',
              '<(source_root)/<(libchromiumcontent_library_dir)/libchromiumviews.a',
              '-lpthread',
              '<!@(pkg-config --libs-only-l gtk+-2.0 libnotify dbus-1 x11 xrandr xext gconf-2.0)',
            ],
          },
        }],
        ['OS=="mac"', {
          'link_settings': {
            'libraries': [
              '<(source_root)/<(libchromiumcontent_library_dir)/libchromiumcontent.dylib',
              '$(SDKROOT)/System/Library/Frameworks/AppKit.framework',
            ],
          },
        }],
        ['OS=="win"', {
          'link_settings': {
            'libraries': [
              '<(source_root)/<(libchromiumcontent_library_dir)/base_static.lib',
              '<(source_root)/<(libchromiumcontent_library_dir)/chromiumcontent.dll.lib',
              '<(source_root)/<(libchromiumcontent_library_dir)/chromiumviews.lib',
              '<(source_root)/<(libchromiumcontent_library_dir)/sandbox_static.lib',
            ],
          },
        }],
      ],
    },
  ],
  'conditions': [
    ['OS=="mac"', {
      'targets': [
        {
          'target_name': '<(project_name)_framework',
          'product_name': '<(product_name)',
          'type': 'shared_library',
          'dependencies': [
            '<(project_name)_lib',
          ],
          'sources': [
            '<@(framework_sources)',
          ],
          'mac_bundle': 1,
          'mac_bundle_resources': [
            'browser/mac/MainMenu.xib',
            'browser/mac/WindowController.xib',
            '<(libchromiumcontent_resources_dir)/content_shell.pak',
            '<(libchromiumcontent_resources_dir)/icudtl.dat',
          ],
          'xcode_settings': {
            'LIBRARY_SEARCH_PATHS': '<(libchromiumcontent_library_dir)',
            'LD_DYLIB_INSTALL_NAME': '@rpath/<(product_name).framework/<(product_name)',
            'LD_RUNPATH_SEARCH_PATHS': '@loader_path/Libraries',
            'OTHER_LDFLAGS': [
              '-ObjC',
            ],
          },
          'copies': [
            {
              'destination': '<(PRODUCT_DIR)/<(product_name).framework/Versions/A/Libraries',
              'files': [
                '<(libchromiumcontent_library_dir)/ffmpegsumo.so',
                '<(libchromiumcontent_library_dir)/libchromiumcontent.dylib',
              ],
            },
          ],
          'postbuilds': [
            {
              'postbuild_name': 'Add symlinks for framework subdirectories',
              'action': [
                'tools/mac/create-framework-subdir-symlinks.sh',
                '<(product_name)',
                'Libraries',
                'Frameworks',
              ],
            },
          ],
          'export_dependent_settings': [
            '<(project_name)_lib',
          ],
        },
        {
          'target_name': '<(project_name)_helper',
          'product_name': '<(product_name) Helper',
          'type': 'executable',
          'dependencies': [
            '<(project_name)_framework',
          ],
          'sources': [
            '<@(app_sources)',
          ],
          'mac_bundle': 1,
          'xcode_settings': {
            'INFOPLIST_FILE': 'renderer/mac/Info.plist',
            'LD_RUNPATH_SEARCH_PATHS': '@executable_path/../../../../Frameworks',
          },
        },
      ],
    }],
  ],
}
