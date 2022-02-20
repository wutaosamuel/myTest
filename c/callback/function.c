#include <stdio.h>

typedef int (*AddFunc)(int, int);
int Add(int, int);
void printCallback(AddFunc, int, int);
void checkNULL(AddFunc, int, int);

int main() {
	printCallback(Add, 1, 1);
	checkNULL(NULL, 1, 1);
}

int Add(int a, int b) { return a + b; }

void printCallback(AddFunc add, int a, int b) {
  printf("add callback result is %d\n", (*add)(a, b));
}

void checkNULL(AddFunc add, int a, int b) {
	if (!add) {
		printf("func pointer is NULL\n");
		return;
	}
	printf("check NULL AddFunc Value %d\n", (*add)(a,b ));
}