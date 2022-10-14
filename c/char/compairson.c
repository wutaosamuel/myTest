#include <stdio.h>
#include <string.h>

int main() {
  char a = 'a';
  char b = 'b';

  int r1 = -1;
  if (a > b)
    r1 = 1;
  else
    r1 = 0;
  printf("Test a > b: %d --> correct, false\n", r1);

  int r2 = -1;
  if (a < b)
    r2 = 1;
  else
    r2 = 0;
  printf("Test a < b: %d --> correct, true\n", r2);

  int r3 = -1;
  if (a == 'a')
    r3 = 1;
  else
    r3 = 0;
  printf("Test a == 'a': %d --> correct, true\n", r3);

  return 0;
}