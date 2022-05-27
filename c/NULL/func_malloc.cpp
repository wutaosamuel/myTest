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
void NewObjectFunc(Object*);

int main() {
  printf("return by func --> works \n");
  Object* obj = NULL;
  obj = NewObject();
  printf("value: %p\n", obj);
  printf("value->name: %p\n", obj->Name);
  printf("value->name: %s\n", obj->Name);
  printf("value->name size: %d\n", obj->NameSize);
  printf("value->id: %d\n", obj->ID);

  printf("\n");
  printf("return by func --> not works \n");
  Object* o = NULL;
  printf("before object point: %p\n", o);
  NewObjectFunc(o);
  printf("after object point: %p\n", o);

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

void NewObjectFunc(Object* object) { object = (Object*)malloc(sizeof(Object)); }