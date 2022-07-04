#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define STR "Hello World!"
#define STR_LEN 13

char** NewCharPP() {
  char** c = (char**)malloc(sizeof(char*));
  *c = (char*)malloc(STR_LEN * sizeof(char));
  strcpy(*c, STR);

  return c;
}

void FreeCharPP(char*** c) {
  if (!c) return;
  if (!*c) return;
  if (!**c) return;
  printf("free now with string: %s\n", **c);
  printf("before free **c: %p\n", **c);
  free(**c);
  **c = NULL;
  printf("after free **c: %p\n", **c);
  printf("before free *c: %p\n", *c);
  free(*c);
  *c = NULL;
  printf("after free *c: %p\n", *c);
}

int main() {
  char** str = NewCharPP();
  printf("before str: %p\n", str);
  FreeCharPP(&str);
  printf("after str: %p\n", str);
  return 0;
}