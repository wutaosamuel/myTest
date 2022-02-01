#include <stdio.h>

typedef int (*AddFunc)(int, int);
int Add(int, int);
void printCallback(AddFunc, int, int);

int main() {
	printCallback(Add, 1, 1);
}

int Add(int a, int b) { return a + b; }

void printCallback(AddFunc add, int a, int b) {
  printf("add callback result is %d\n", (*add)(a, b));
}