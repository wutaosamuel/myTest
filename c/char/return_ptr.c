#include <stdio.h>
#include <string.h>

#define STRING "This is a test string"

const char* StringFunc1();

int main() {
  printf("%s\n", StringFunc1());

  return 0;
}

const char* StringFunc1() { return STRING; }
