#include <stdio.h>
#include <string.h>

void fromVar();
void withAddition();
void withFormat();

int main() {
  printf("from var: ---> fail\n");
  fromVar();

  printf("\n");

  printf("with addition: ---> fail\n");
  withAddition();

  printf("\n");
  
  printf("with format: \n");
  withFormat();
  return 0;
}

void fromVar() {
  char source[] = "Hello world";
  char str[100] = "";

  sprintf(str, source);
  printf("From other variable: %s\n", str);
}

void withAddition() {
  char str[100] = "";

  sprintf(str, "Hello world!", 50);
  printf("with addition: %s\n", str);
}

void withFormat() {
  char str[100] = "";
  char s[] = "World";

  sprintf(str, "%s %s!!!", "Hello", s);
  printf("with formate: %s\n", str);
}