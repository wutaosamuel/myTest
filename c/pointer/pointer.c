#include <stdio.h>

int main()
{
  int a, *b, c;
  c = 1;
  b = &c;
  a = *b;
  printf("a: %d\n", a);
  printf("b: %d\n", *b);
  a = 2;
  *b = 5;
  printf("a: %d\n", a);
  printf("b: %d\n", *b);
}