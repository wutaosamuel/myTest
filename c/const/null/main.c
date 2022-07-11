#include <stdio.h>
#include <stdlib.h>

struct Operation {
  void (*Print)(void* self);
};

struct Object {
  int ID;

  const struct Operation* op;
};

struct Object* NewObject() {
  struct Object* obj = (struct Object*)malloc(sizeof(struct Object));
  obj->ID = 5;
  obj->op = NULL;

  return obj;
}

void pt(void* self) {
  struct Object* _self = (struct Object*)self;
  printf("object id: %d\n", _self->ID);
}

static struct Operation op = {
    .Print = pt,
};

int main() {
  struct Object* obj = NewObject();
  obj->op = &op;

  obj->op->Print((void*)obj);

  return 0;
}