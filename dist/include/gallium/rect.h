#ifndef GALLIUM_RECT_H_
#define GALLIUM_RECT_H_

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct GALLIUM_EXPORT {
	int width;
	int height;
	int left;
	int bottom;
} gallium_rect_t;

#ifdef __cplusplus
}
#endif

#endif // ifndef GALLIUM_RECT_H_
