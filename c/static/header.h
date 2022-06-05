#ifndef STATIC_HEADER_H
#define STATIC_HEADER_H

#include <stdio.h>

typedef struct {
  int Type;
  int ID;
} Object;

//void s_Print(Object* obj);

// Not work in this way
inline void i_Print(Object* obj);
inline void i_Print(Object* obj) {
  printf("Object in inline type: %d, ID: %d\n", obj->Type, obj->ID);
};

extern inline void e_Print(Object* obj);
inline void e_Print(Object* obj) {
  printf("Object in a extern inline type: %d, ID: %d\n", obj->Type, obj->ID);
};

#endif