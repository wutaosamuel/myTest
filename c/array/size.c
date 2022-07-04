#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define STR "Hello world!"

int main() {
	size_t len = strlen(STR);
	printf("Not working\n");
	char* arr = (char*)malloc((len + 10) * sizeof(char));
	printf("arr size: %ld, STR len: %ld\n", strlen(arr), len);

	free(arr);
	return 0;
}