#include <stdio.h>

void printLocal() {
  static int local = 10;

  printf("static local value = %d\n", local);
}

void printLocalAgain() {
  int local;
  printf("local value = %d\n", local);
}

void printLocalAdd() {
  static int local = 10;

  printf("static local value = %d\n", local++);
}

void printAdd() {
  static int add;
  add = 10;

  printf("static local value = %d\n", add++);
}

void printAddMore() {
  static int add = 5;
  add = 10;

  printf("static local value = %d\n", add++);
}

int main() {

  printf("\n");
  printf("test static local variable now --> works\n");

  printLocal();

  printLocalAgain();

  for (int i=0; i < 5; i++)
    printLocalAdd();

  for (int i=0; i < 5; i++)
    printAdd();

  for (int i=0; i < 5; i++)
    printAddMore();

  return 0;
}