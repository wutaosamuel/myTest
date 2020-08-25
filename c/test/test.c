#include "stdio.h"

int main() {

  float x = fc(10/5);
  printf("%f", x);
}

int fc(int a, int b) {
  return a/b;
}