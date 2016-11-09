#ifndef GALLIUM_API_COCOA_H_
#define GALLIUM_API_COCOA_H_

#include <stdbool.h>
#include <stdint.h>

#include "gallium/browser.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef void(*gallium_callback_t)(void*);

typedef struct GALLIUM_EXPORT gallium_nsmenu gallium_nsmenu_t;
typedef struct GALLIUM_EXPORT gallium_nsmenuitem gallium_nsmenuitem_t;
typedef struct GALLIUM_EXPORT gallium_nsusernotification gallium_nsusernotification_t;
typedef struct GALLIUM_EXPORT gallium_nsimage gallium_nsimage_t;

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

GALLIUM_EXPORT void NSMenu_AddSeparator(
	gallium_nsmenu_t* menu);

GALLIUM_EXPORT void NSMenuItem_SetSubmenu(
	gallium_nsmenuitem_t* menuitem,
	gallium_nsmenu_t* submenu);

GALLIUM_EXPORT void NSStatusBar_AddItem(
	gallium_nsimage_t* image,
	float width,
	bool highlightMode,
	gallium_nsmenu_t* menu);

  
GALLIUM_EXPORT gallium_nsusernotification_t* NSUserNotification_New(
	const char* title,
	const char* subtitle,
	const char* informativeText,
	gallium_nsimage_t* contentImage,
	const char* identifier,
	bool hasActionButton,
	bool hasReplyButton,
	const char* actionButtonTitle,
	const char* otherButtonTitle);

GALLIUM_EXPORT void NSUserNotificationCenter_DeliverNotification(
	gallium_nsusernotification_t* n);

GALLIUM_EXPORT gallium_nsimage_t* NSImage_NewFromPNG(
	const void* buf,
	int size);

GALLIUM_EXPORT void NSImage_WriteToFile(gallium_nsimage_t* img, const char* path);

GALLIUM_EXPORT void NSApplication_SetMainMenu(
	gallium_nsmenu_t* submenu);

GALLIUM_EXPORT void NSApplication_Run();

GALLIUM_EXPORT char* MainBundle_ObjectForKey(const char* key);

#ifdef __cplusplus
}
#endif

#endif
