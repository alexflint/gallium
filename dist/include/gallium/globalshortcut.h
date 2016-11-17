#ifndef GALLIUM_GLOBALSHORTCUT_H_
#define GALLIUM_GLOBALSHORTCUT_H_

#include "gallium/core.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef void(*global_shortcut_handler_t)(int);

void GALLIUM_EXPORT GalliumAddGlobalShortcut(int ID,
                                             const char* key,
                                             gallium_modifier_t mask,
                                             global_shortcut_handler_t handler);

#ifdef __cplusplus
}
#endif

#endif // ifndef GALLIUM_GLOBALSHORTCUT_H_

