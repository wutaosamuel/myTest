#include <stdio.h>
#include <stdlib.h>

int main() {
  int *a = NULL;
  a = (int *)malloc(sizeof(int));
  *a = 2;

  printf("after transfer void* the value\n");
	void *v = NULL;
	v = a;
	// *v = a; // complier fault
  printf("address value a=%p, v=%p\n", a, v);

  free(a);
}