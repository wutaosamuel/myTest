#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef char* str;

str* NewString(const char* string) {
  size_t size = strlen(string);
  printf("string size: %ld\n", size);
  printf("str sizeof: %ld\n", sizeof(str));
  str* _string = (str*)malloc(size * sizeof(str));
  printf("_string size: %ld\n", sizeof(_string));
	memcpy(_string, string, size+1);

  return _string;
}

void Pt(str* s) {
	printf("\n");
}

int main() {
  printf("char sizeof: %ld\n", sizeof(char));
  printf("char* sizeof: %ld\n", sizeof(char*));
  printf("test char** --> pass\n");
  str* s = NewString("This is a string");
	printf("\n");

  return 0;
}