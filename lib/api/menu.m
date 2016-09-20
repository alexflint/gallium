#include "menu.h"
#import <Cocoa/Cocoa.h>
#import <AppKit/AppKit.h>

// CCallback is an objective C wrapper around a C function pointer. It 
// has a single method "call", which invokes the underlying C function
// pointer with the void* data pointer provided to the constructor. This
// makes it possible to pass a C function pointer to APIs that follow the
// target/selector pattern.
@interface CCallback : NSObject {
	gallium_callback_t _cfunc;
	void* _arg;
}
- (CCallback*)initWithFunc:(gallium_callback_t)cfunc arg:(void*)arg;
- (IBAction)call:(id)sender;
@end

@implementation CCallback
- (CCallback*)initWithFunc:(gallium_callback_t)cfunc arg:(void*)arg {
	_cfunc = cfunc;
	_arg = arg;
	return self;
}

- (IBAction)call:(id)sender {
	NSLog(@"in CCallback:call");
	if (_cfunc == nil) {
		NSLog(@"got CCallback:call but function pointer was nil");
		return;
	}
	_cfunc(_arg);
}
@end


typedef struct gallium_nsmenu {
	NSMenu* impl;
} gallium_nsmenu_t;

typedef struct gallium_nsmenuitem {
	NSMenuItem* impl;
} gallium_nsmenuitem_t;

NSString* str(const char* cstr) {
	return [NSString stringWithUTF8String:cstr];
}

gallium_nsmenu_t* NSMenu_New(const char* title) {
	gallium_nsmenu_t* st = (gallium_nsmenu_t*)malloc(sizeof(gallium_nsmenu_t));
	st->impl = [[NSMenu alloc] initWithTitle:str(title)];
	return st;
}

NSEventModifierFlags makeModifierMask(gallium_modifier_t mask) {
	// The OSX docs say that constants like "NSEventModifierFlagControl"
	// are deprecated in favor of constants like "NSControlKeyMask" but 
	// I cannot find the definitions of the latter symbols. Perhaps they
	// are only defined in the swift versions of these APIs.
	NSEventModifierFlags out = 0;
	if (mask & GalliumCmdModifier) {
		out |= NSCommandKeyMask;
	}
	if (mask & GalliumCtrlModifier) {
		out |= NSControlKeyMask;
	}
	if (mask & GalliumCmdOrCtrlModifier) {
		out |= NSCommandKeyMask;
	}
	if (mask & GalliumAltOrOptionModifier) {
		out |= NSAlternateKeyMask;
	}
	if (mask & GalliumFunctionModifier) {
		out |= NSFunctionKeyMask;
	}
	if (mask & GalliumShiftModifier) {
		out |= NSShiftKeyMask;
	}
	return out;
}

gallium_nsmenuitem_t* NSMenu_AddMenuItem(
	gallium_nsmenu_t* menu,
	const char* title,
	const char* shortcutkey,
	const gallium_modifier_t shortcutModifier,
	gallium_callback_t callback,
	void* callbackArg) {

	gallium_nsmenuitem_t* st = (gallium_nsmenuitem_t*)malloc(sizeof(gallium_nsmenuitem_t));
	st->impl = [[NSMenuItem alloc] initWithTitle:str(title) action:@selector(call:) keyEquivalent:@""];
	if (shortcutkey != nil) {
		[st->impl setKeyEquivalent:str(shortcutkey)];
		[st->impl setKeyEquivalentModifierMask:makeModifierMask(shortcutModifier)];
	}
	[st->impl setTarget:[[CCallback alloc] initWithFunc:callback arg:callbackArg]];
	[menu->impl addItem:st->impl];
	return st;
}

void NSMenuItem_SetSubmenu(
	gallium_nsmenuitem_t* menuitem,
	gallium_nsmenu_t* submenu) {

	menuitem->impl.submenu = submenu->impl;
}

void NSApplication_SetMainMenu(gallium_nsmenu_t* menu) {
	[[NSApplication sharedApplication] setMainMenu:menu->impl];
}

void NSApplication_Run() {
	[[NSApplication sharedApplication] run];
}

void NSStatusBar_AddItem(
	int width,
	const char* title,
	bool highlightMode,
	gallium_nsmenu_t* menu) {

	NSStatusBar* bar = [NSStatusBar systemStatusBar];
	NSStatusItem* item = [bar statusItemWithLength:NSVariableStatusItemLength];
	[item setTitle: str(title)];
	[item setHighlightMode:highlightMode];
	if (menu != nil) {
		[item setMenu:menu->impl];
	}
}

void SetUIApplication() {
	NSLog(@"in SetUIApplication");
	[[NSApplication sharedApplication] setActivationPolicy:NSApplicationActivationPolicyRegular];
	[[NSApplication sharedApplication] setPresentationOptions:NSApplicationPresentationDefault]; // probably not necessary since it's the default
	[[NSApplication sharedApplication] activateIgnoringOtherApps:NO];
	[NSMenu setMenuBarVisible:NO]; // these two lines may not be necessary, either; using -setActivationPolicy: instead of TransformProcessType() may be enough
	[NSMenu setMenuBarVisible:YES];
}
