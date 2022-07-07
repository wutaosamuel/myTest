#include "struct.h"

void pt1(int id) { printf("the id is %d\n", id); }

static const struct Operation ObjectOp = {
    .print_id = pt1,
};

struct Object* NewObject() {
  struct Object* obj = malloc(sizeof(struct Object));
  obj->ID = 0;
  obj->op = &ObjectOp;
  return obj;
}
