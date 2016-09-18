#include "menu.h"
#import <Cocoa/Cocoa.h>

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

typedef struct gallium_nsapplication {
	NSApplication* impl;
} gallium_nsapplication_t;

NSString* str(const char* cstr) {
	return [NSString stringWithUTF8String:cstr];
}

gallium_nsmenu_t* NSMenu_New(const char* title) {
	gallium_nsmenu_t* st = (gallium_nsmenu_t*)malloc(sizeof(gallium_nsmenu_t));
	st->impl = [[NSMenu alloc] initWithTitle:str(title)];
	return st;
}

gallium_nsmenuitem_t* NSMenu_AddMenuItem(
	gallium_nsmenu_t* menu,
	const char* title,
	const char* keyEquivalent,
	gallium_callback_t callback,
	void* callbackArg) {

	gallium_nsmenuitem_t* st = (gallium_nsmenuitem_t*)malloc(sizeof(gallium_nsmenuitem_t));
	st->impl = [[NSMenuItem alloc] initWithTitle:str(title) action:@selector(call:) keyEquivalent:str(keyEquivalent)];
	[st->impl setTarget:[[CCallback alloc] initWithFunc:callback arg:callbackArg]];
	[menu->impl addItem:st->impl];
	return st;
}

void NSMenuItem_SetSubmenu(
	gallium_nsmenuitem_t* menuitem,
	gallium_nsmenu_t* submenu) {

	menuitem->impl.submenu = submenu->impl;
}

gallium_nsapplication_t* NSApplication_SharedApplication() {
	gallium_nsapplication_t* st = (gallium_nsapplication_t*)malloc(sizeof(gallium_nsapplication_t));
	st->impl = [NSApplication sharedApplication];
	return st;
}

void NSApplication_SetMainMenu(
	gallium_nsapplication_t* app,
	gallium_nsmenu_t* menu) {

	app->impl.mainMenu = menu->impl;
}

void NSApplication_Run(gallium_nsapplication_t* app) {
	[app->impl run];
}
