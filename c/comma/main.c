#include <stdio.h>

#define NUM 1,000

int main() {
	int a = 0;

	for(int i=0; i<NUM; i++) {
		a++;
		printf("value %d", a);
	} 

	return 0;
}