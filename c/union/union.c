#include <stdio.h>

union DaTa {
  int n;
  char c;
};

int main() {
  union DaTa a;

	a.n = -1;
	a.c = 'a';
  printf("a: %d\n", a.n);
  printf("b: %c\n", a.c);
  printf("\n");

  for (int i = 0; i < 3; i++) {
    a.n = i;
    printf("a: %d\n", a.n);
    printf("b: %c\n", a.c);
    printf("\n");
  }
}