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

int main() {
  Object* obj = NULL;
  obj = NewObject();
  printf("value: %p\n", obj);
  printf("value->name: %p\n", obj->Name);
  printf("value->name: %s\n", obj->Name);
  printf("value->name size: %d\n", obj->NameSize);
  printf("value->id: %d\n", obj->ID);
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