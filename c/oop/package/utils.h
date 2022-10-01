#ifndef MY_UTILS_H
#define MY_UTILS_H
#include <stdio.h>

// not work with struct UtilsInterface, it requires typedef
struct UtilsInterface {
  int (*Printf)(const char*);
};

extern const struct UtilsInterface utils;

#endif