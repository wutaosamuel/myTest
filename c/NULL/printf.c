#include <stdio.h>

typedef void* pvoid;

int main() {
	int* i = NULL;
	int j = 0;
	pvoid k = NULL;

	printf("NULL value: %p\n", i);
	printf("int value: %p\n", &j);
	printf("pvoid(void*) value: %p\n", k);
	return 0;
}