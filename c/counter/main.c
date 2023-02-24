#include <stdio.h>

#define NUM 10

int main() {
	int counter = 0;

	if (counter++ >= NUM) {
		printf("counter if %d\n", counter);
	}

	printf("counter %d\n", counter);

	return 0;
}