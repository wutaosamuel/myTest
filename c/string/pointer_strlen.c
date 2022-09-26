#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SizeAnswer 13
#define Answer "Hello World!\n"

int main() {
  char* str = (char*)malloc(15 * sizeof(char));
  *(str + 0) = 'H';
  *(str + 1) = 'e';
  *(str + 2) = 'l';
  *(str + 3) = 'l';
  *(str + 4) = 'o';
  *(str + 5) = ' ';
  *(str + 6) = 'W';
  *(str + 7) = 'o';
  *(str + 8) = 'r';
  *(str + 9) = 'l';
  *(str + 10) = 'd';
  *(str + 11) = '!';
  *(str + 12) = '\n';
  *(str + 13) = '\0';

  printf("answer size: %d\n", (int)strlen(Answer));
  printf("str size: %d\n", (int)strlen(str));
}