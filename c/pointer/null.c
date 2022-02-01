#include <stdio.h>
#include <stdlib.h>

int main() {
  int *a = NULL;
  if (!a) printf("if a == NULL is true\n");

  printf("after setting the value\n");
  a = (int *)malloc(sizeof(int));
  *a = 2;
  if (a) printf("b: %d\n", *a);
  free(a);
}