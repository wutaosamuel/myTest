#ifndef MY_OBJECT_H
#define MY_OBJECT_H

#include <stdio.h>
#include <stdlib.h>

#include "interface.h"

struct Object {
  int ID;

  const struct StructInterface* interface;
};

static inline void printObject(void* object) {
  struct Object* obj = (struct Object*)object;

  printf("object: \n");
  printf("the id is %d\n", obj->ID);
};

static const struct StructInterface objectInterface = {
    .Print = printObject,
};

struct Object* NewObject() {
  struct Object* obj = (struct Object*)malloc(sizeof(struct Object));
  obj->ID = 0;
  obj->interface = &objectInterface;
  return obj;
}

#endif