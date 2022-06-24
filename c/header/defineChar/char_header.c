#include <stdio.h>

#define CHAR_STRING \
  "hello "           \
  "world"           \
  "!"

int main() {
  char str[] = {CHAR_STRING};
  int size = sizeof(str);

  printf("there no unsigned int8_t\n");
  printf("str: %s\nsize: %d\n", str, size);
}