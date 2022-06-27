#ifndef STATIC_HEADER_H
#define STATIC_HEADER_H

#include <stdio.h>

typedef struct {
  int Type;
  int ID;
} Object;

static inline void s_Print(Object* obj) {
  printf("Object in inline type: %d, ID: %d\n", obj->Type, obj->ID);
};

#endif