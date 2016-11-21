#ifndef GALLIUM_EXPORT_H_
#define GALLIUM_EXPORT_H_

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#define GALLIUM_EXPORT __attribute__ ((visibility ("default")))

// gallium_error represents an error
typedef struct GALLIUM_EXPORT gallium_error {
	const char* msg;
} gallium_error_t;

#ifdef __cplusplus
}
#endif

#endif // ifndef GALLIUM_EXPORT_H_

