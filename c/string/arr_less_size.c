#include <stdio.h>
#include <string.h>

int main() {
	char str[300] = "Hello World!\n";
	int a[300];

	int realSize = strlen(str);
	printf("real size is %d\n", realSize);
}