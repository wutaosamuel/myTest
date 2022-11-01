#include <stdio.h>
#include <string.h>

void case1(const char* str);

int main() {
	char str[] = "012345678901234567890123456789";

	printf("Start case 1\n");
	case1(str);
	return 0;
}

void case1(const char* str) {
	char s[10];
	strcpy(s, "01234\0");
	int left = 10 - strlen(s);
	printf("original string: %s, length: %lu, space left: %d\n", s, strlen(s), left);
	strncat(s, str, left-1);
	printf("after concatenation, string: %s, length %lu\n", s, strlen(s));
}