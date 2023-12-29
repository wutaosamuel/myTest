#include <stdio.h>
#include <stdlib.h>

int* f_add(int, int);

int main() {
	int result = 0;
  printf("init result = %d\n", result);

	result = *f_add(1, 1);
  printf("1 + 1 = %d\n", result);

	return 0;
}

int* f_add(int a, int b) {
	// int res = a + b; // complier warning
	int *res = (int*)malloc(sizeof(int));
	*res = a + b;

	return res;
}