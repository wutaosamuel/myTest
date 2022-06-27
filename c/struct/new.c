#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Object {
  char* Name;
  int ID;
};

static inline struct Object* NewObject() {
  char* name = (char*)malloc(sizeof(char));
  *name = '\0';

  struct Object* obj = (struct Object*)malloc(sizeof(struct Object));
  obj->Name = name;
  obj->ID = 0;

  return obj;
};

static inline void Free(struct Object* obj) {
  if (obj) return;
  if (obj->Name) free(obj->Name);
  free(obj);
}

static inline void Print(struct Object* obj) {
  printf("object name: %s, id: %d\n", obj->Name, obj->ID);
  printf("object ptr: %p\n", obj);
  printf("object name size: %ld\n", strlen(obj->Name));
};

static inline void SetName(struct Object* obj, const char* str) {
  size_t len = strlen(str);
  if (obj->Name == NULL) {
    obj->Name = (char*)malloc(len * sizeof(char));
    strcpy(obj->Name, str);
    return;
  }
  if (strlen(obj->Name) != len) {
    free(obj->Name);
    obj->Name = (char*)malloc(len * sizeof(char));
    printf("not set obj->Name value yet size: %ld\n", strlen(obj->Name));
  }

  strcpy(obj->Name, str);
}

int main() {
  struct Object* obj = NewObject();
  Print(obj);

  printf("\n");
  printf("set new name\n");
	SetName(obj, "NewName");
  Print(obj);

  free(obj);
  return 0;
}