#ifndef LIB_MY_STRING_H
#define LIB_MY_STRING_H

#include <stdlib.h>
#include <string.h>

inline int utils_SetString(char* dest, const char* str) {
  size_t len = strlen(str);
  // is NULL
  if (!dest) {
    return 0;
  }

  if (strlen(dest) != len) {
    free(dest);
    dest = (char*)malloc(len * sizeof(char));
  }

  strcpy(dest, str);
  return 1;
};

inline int utils_AppendString(char* dest, const char* str) {
  size_t len = strlen(str);
  // is NULL
  if (!dest) {
    return 0;
  }

  strncat(dest, str, len);
  return 1;
};

#endif