#include "LimitString.h"

void TestCreateString() {
  utilsLS_String* empty = utilsLS_NewStr();
  printf("Test NewStr() --> pass\n");
  printf("empty pointer: %p\n", empty);
  printf("empty size: %ld\n", utilsLS_Size(empty));
  printf("empty size: %ld\n", sizeof(*empty));

  utilsLS_String* string = utilsLS_NewString("This is string");
  printf("\n");
  printf("Test NewString() --> pass\n");
  printf("string pointer: %p\n", string);
  printf("string size: %ld\n", utilsLS_Size(string));

  utilsLS_FreeString(string);
  utilsLS_FreeString(empty);
  empty = NULL;
  printf("\n");
  printf("Test FreeString() --> pass\n");
  printf("empty pointer: %p\n", empty);
  printf("empty size: %ld\n", sizeof(*empty));
  printf("string pointer: %p\n", string);
}

int main() {
  TestCreateString();
  return 0;
}