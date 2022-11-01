#include <stdio.h>
#include <string.h>

void pt(const char*);

int main() {
  char str[] = "this is test string";

	printf("Start to test const char\n");
	pt(str);
  return 0;
}

void pt(const char* str) {
  int len = strlen(str);

  if (len < 4) {
    printf("len: %d and str: %s\n", len, str);
    return;
  }

  printf("len: %d and str: ", len);
  for (int i = 0; i < 4; i++) {
    printf("%c", *(str + i));
  }
  printf("\n");

	pt(str+4);
}
