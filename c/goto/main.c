#include <stdio.h>

void a() {
	printf("this is 1 func run\n");
}

void b() {
	printf("this is 2 func run\n");
}

void c() {
	printf("this is 3 func run\n");
}

int pt(int i) {
	if (i == 1) {
		goto aTag;
	}
	if (i == 2) {
		goto bTag;
	}
	if (i == 3) {
		goto cTag;
	}

	return -1;

aTag:
	a();
bTag:
  b();
cTag:
	c();

	return 1;
}

int main() {
	printf("start to test goto keyword\n");
	printf("start to test 1\n");
	int result = pt(1);
	printf("result of testing 1 is %d\n", result);
	return 0;
}
