#include <stdio.h>

void floatTest();
void doubleTest();

int main() {
  float f = 0.1;
  double d = 0.1;

  printf("float: %f, double: %lf\n", f, d);

  floatTest();
  doubleTest();
  return 0;
}

void floatTest() {
  int value = 65535;
  float v1 = (float)value * 196.8;
  float v2 = v1 * 27;
  printf("value: %d, v1: %f, v2: %f\n", value, v1, v2);
}

void doubleTest() {
  int value = 65535;
  double v1 = (double)value * 196.8;
  double v2 = v1 * 27;
  printf("value: %d, v1: %lf, v2: %lf\n", value, v1, v2);
}