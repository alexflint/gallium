#include "menu.h"
#import <Cocoa/Cocoa.h>

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
	const char* keyEquivalent) {

	gallium_nsmenuitem_t* st = (gallium_nsmenuitem_t*)malloc(sizeof(gallium_nsmenuitem_t));
	st->impl = [menu->impl addItemWithTitle:str(title) action:nil keyEquivalent:str(keyEquivalent)];
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

void SetMenu(gallium_menu_t* menu, gallium_error_t** err) {
	NSMenu* root = [[NSMenu alloc] initWithTitle:@"root"];
	
	NSMenuItem* gallium = [root addItemWithTitle:@"Gallium" action:nil keyEquivalent:@""];
	gallium.submenu = [[NSMenu alloc] initWithTitle:@"Gallium"];
	[gallium.submenu addItemWithTitle:@"item" action:nil keyEquivalent:@""];

	NSMenuItem* view = [root addItemWithTitle:@"View" action:nil keyEquivalent:@""];
	view.submenu = [[NSMenu alloc] initWithTitle:@"View"];
	[view.submenu addItemWithTitle:@"Show" action:nil keyEquivalent:@""];

	NSApplication* app = [NSApplication sharedApplication];
	app.mainMenu = root;
	[app run];
}
