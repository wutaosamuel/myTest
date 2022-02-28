#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  char* Name;
  int NameSize;
  int ID;
} Object;

Object* NewObject();
void SetObject(Object*, const char*, int, int);
void init_object_list();
void init_list(Object**);

static Object** list1 = NULL;
Object** list2 = NULL;

int main() {
  printf("list1 size: %ld\n", sizeof(list1));
  printf("list2 size: %ld\n", sizeof(list2));
  printf("list1 pointer: %p\n", list1);
  printf("list2 pointer: %p\n", list2);
  init_object_list();
  printf("after initing, list1 size: %ld\n", sizeof(list1));
  printf("after initing, list2 size: %ld\n", sizeof(list2));
  printf("list1 pointer: %p\n", list1);
  printf("list2 pointer: %p\n", list2);
  return 0;
}

Object* NewObject() {
  Object* object = (Object*)malloc(sizeof(Object));
  SetObject(object, "object 0", 8, 0);

  return object;
}

void SetObject(Object* object, const char* name, int size, int id) {
  object->NameSize = size;
  object->Name = (char*)malloc(size * sizeof(char));
  for (int i = 0; i < size; i++) object->Name[i] = name[i];
  object->ID = id;
}

void init_object_list() {
  list1 = (Object**)malloc(8 * sizeof(Object*));
  // list2 = (Object**)malloc(8 * sizeof(Object*));
  init_list(list2);
}

void init_list(Object** list) { list = (Object**)malloc(8 * sizeof(Object*)); }