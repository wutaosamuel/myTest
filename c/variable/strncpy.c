#include <stdio.h>
#include <string.h>

void fromVar();
void fromStack();

int main() {
  printf("from var: \n");
  fromVar();

  printf("\n");

  printf("from string: \n");
  fromStack();
  return 0;
}

void fromVar() {
  char source[] = "Hello world";
  char str[100] = "";

  strncpy(str, source, 20);
  printf("From other variable: %s\n", str);
}

void fromStack() {
  char str[100] = "";

  strncpy(str, "Hello world!", 50);
  printf("From a string: %s\n", str);
}