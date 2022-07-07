#include <stdio.h>
#include <stdlib.h>

struct Operation {
  void (*print_id)(int);
};

struct Object {
  int ID;

  const struct Operation *op;
};

struct Object* NewObject();