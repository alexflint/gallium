#ifndef GALLIUM_API_MENU_H_
#define GALLIUM_API_MENU_H_

#include "gallium.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct gallium_menu_item gallium_menu_item_t;

// gallium_error represents an error
typedef struct GALLIUM_EXPORT gallium_menu {
	const char* title;
	gallium_menu_item_t* items;
	int item_count;
} gallium_menu_t;

// gallium_error represents an error
typedef struct GALLIUM_EXPORT gallium_menu_item {
	const char* title;
	const char* shortcut;
	gallium_menu_t* submenu;
} gallium_menu_item_t;

// GalliumCreateWindow creates a window pointed at the given url
GALLIUM_EXPORT void SetMenu(gallium_menu_t* menu, gallium_error_t** err);

typedef struct GALLIUM_EXPORT gallium_nsmenu gallium_nsmenu_t;
typedef struct GALLIUM_EXPORT gallium_nsmenuitem gallium_nsmenuitem_t;
typedef struct GALLIUM_EXPORT gallium_nsapplication gallium_nsapplication_t;

GALLIUM_EXPORT gallium_nsmenu_t* NSMenu_New(const char* title);

GALLIUM_EXPORT gallium_nsmenuitem_t* NSMenu_AddMenuItem(
	gallium_nsmenu_t* menu,
	const char* title,
	const char* keyEquivalent);

GALLIUM_EXPORT void NSMenuItem_SetSubmenu(
	gallium_nsmenuitem_t* menuitem,
	gallium_nsmenu_t* submenu);

GALLIUM_EXPORT gallium_nsapplication_t* NSApplication_SharedApplication();

GALLIUM_EXPORT void NSApplication_SetMainMenu(
	gallium_nsapplication_t* app,
	gallium_nsmenu_t* submenu);

GALLIUM_EXPORT void NSApplication_Run(
	gallium_nsapplication_t* app);

#ifdef __cplusplus
}
#endif

#endif
