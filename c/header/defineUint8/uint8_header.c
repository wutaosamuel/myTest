#include <stdio.h>

#define UINT8_HEADER_0 0, 2, 3
#define UINT8_STRING "hello world!"

int main() {
  char str[] = {UINT8_STRING};
  int size = sizeof(str);

  printf("there no unsigned int8_t\n");
  printf("str: %s\nsize: %d\n", str, size);
}