#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static inline void TestStrcat() {
  char val[] = "the string";
  size_t l = strlen(val);
  char* str = (char*)malloc(sizeof(char));
  // char* str = NULL;
  printf("before size: %ld\n", strlen(str));

  // str=NULL not work
  // strcat(str, val);
  // str=NULL not work
  // strncat(str, val, l);
  // strncat(str, val, l);
  strcat(str, val);

  printf("after size: %ld\n", strlen(str));
  printf("strcat: %s\n", str);
  free(str);
};

int main() {
  printf("strcat/strncat not work with NULL\n");
  printf("strcat/strncat is work with small length\n");
  TestStrcat();

  return 0;
}
