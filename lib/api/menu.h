#ifndef GALLIUM_API_MENU_H_
#define GALLIUM_API_MENU_H_

#include "gallium.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef void(*gallium_callback_t)(void*);

typedef struct GALLIUM_EXPORT gallium_nsmenu gallium_nsmenu_t;
typedef struct GALLIUM_EXPORT gallium_nsmenuitem gallium_nsmenuitem_t;

typedef enum gallium_modifier {
	GalliumCmdModifier = 1 << 0,
	GalliumCtrlModifier = 1 << 1,
	GalliumCmdOrCtrlModifier = 1 << 2,
	GalliumAltOrOptionModifier = 1 << 3,
	GalliumFunctionModifier = 1 << 4,
	GalliumShiftModifier = 1 << 5,
} gallium_modifier_t;

GALLIUM_EXPORT gallium_nsmenu_t* NSMenu_New(const char* title);

GALLIUM_EXPORT gallium_nsmenuitem_t* NSMenu_AddMenuItem(
	gallium_nsmenu_t* menu,
	const char* title,
	const char* shortcutkey,
	const gallium_modifier_t shortcutModifier,
	gallium_callback_t callback,
	void* callbackArg);

GALLIUM_EXPORT void NSMenuItem_SetSubmenu(
	gallium_nsmenuitem_t* menuitem,
	gallium_nsmenu_t* submenu);

GALLIUM_EXPORT void NSApplication_SetMainMenu(
	gallium_nsmenu_t* submenu);

GALLIUM_EXPORT void NSApplication_Run();

// Tells OSX that this is a UI application
GALLIUM_EXPORT void SetUIApplication();


#ifdef __cplusplus
}
#endif

#endif
