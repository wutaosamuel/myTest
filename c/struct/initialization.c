#include <stdio.h>

struct Object {
  int ID;
  char* Name;
  int Data[8];
  struct Object* next;
};

void printObject(struct Object);

int main() {
  struct Object Obj = {
      .ID = 0,
      .Name = "name",
      .Data = {1, [1 ... 2] = 10, 4},
      .next = NULL,
  };

  printObject(Obj);

  return 0;
}

void printObject(struct Object obj) {
  printf("object name: %s, id: %d, data:", obj.Name, obj.ID);
  for (int i = 0; i < 8; i++) printf(" %d", obj.Data[i]);
  printf(" next addr: %p\n", obj.next);
}
