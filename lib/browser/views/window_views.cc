#include "browser/views/window_views.h"

#include "brightray/browser/inspectable_web_contents.h"
#include "brightray/browser/inspectable_web_contents_view.h"

#include "ui/views/layout/fill_layout.h"
#include "ui/views/widget/widget.h"
#include "ui/views/widget/widget_delegate.h"

#include "base/strings/utf_string_conversions.h"

namespace brightray_example {

namespace {

class WidgetDelegateView : public views::WidgetDelegateView {
 public:
  WidgetDelegateView(scoped_ptr<WindowViews> window)
      : window_(window.Pass()) {
    SetLayoutManager(new views::FillLayout);
  }
  ~WidgetDelegateView() {
  }

  void DeleteDelegate() override { delete this; }
  views::View* GetContentsView() override { return this; }
  bool CanResize() const override { return true; }
  bool CanMaximize() const override { return true; }
  base::string16 GetWindowTitle() const override {
    return base::ASCIIToUTF16("Brightray Example");
  }
  gfx::Size GetPreferredSize() const override { return gfx::Size(800, 600); }
  gfx::Size GetMinimumSize() const override { return gfx::Size(100, 100); }

 private:
  scoped_ptr<WindowViews> window_;

  DISALLOW_COPY_AND_ASSIGN(WidgetDelegateView);
};

}

Window* Window::Create(brightray::BrowserContext* browser_context) {
  return new WindowViews(browser_context);
}

WindowViews::WindowViews(brightray::BrowserContext* browser_context)
    : Window(browser_context),
      widget_(new views::Widget) {
  auto delegate_view = new WidgetDelegateView(make_scoped_ptr(this).Pass());

  views::Widget::InitParams params;
  params.delegate = delegate_view;
  widget_->Init(params);
  delegate_view->AddChildView(inspectable_web_contents()->GetView()->GetView());
  delegate_view->Layout();
  WindowReady();
}

WindowViews::~WindowViews() {
}

void WindowViews::Show() {
  widget_->Show();
}

}
