#include <stdio.h>
#include <stdlib.h>

int main() {
  int size = 8;
  int* arr = (int*)malloc(size * sizeof(int));
	printf("arr pointer: %p\n", arr);
	printf("arr+0 pointer: %p\n", (arr+0));
	printf("arr+1 pointer: %p\n", (arr+1));
  return 0;
}