#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* NewString(const char* string);
char* NewChar(char c);

int main() {
  char* str = NewString("this is a string");
  int size = strlen(str);
  printf("size: %d\n", size);
  for (int i = 0; i < size; i++) {
    printf("%c", str[i]);
  }
  printf("\n");

  return 0;
}

char* NewString(const char* string) {
  int size = strlen(string);
  printf("new size: %d\n", size);
  char* chars = (char*)malloc(size * sizeof(char));
  memcpy(chars, string, size);

  return chars;
}

char* NewChar(char c) {
  char* ch = (char*)malloc(sizeof(char));
  memset(ch, (int)c, 1);

  return ch;
}
