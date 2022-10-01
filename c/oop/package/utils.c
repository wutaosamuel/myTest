#include "utils.h"

#include <stdarg.h>
#include <string.h>

static inline int utils_Printf(const char*);

const struct UtilsInterface utils = {
    .Printf = utils_Printf,
};

static inline int utils_Printf(const char* str) { return printf("%s", str); }