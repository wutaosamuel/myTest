#include "debug_printf.h"

#include <iostream>

void DBG_PRINTF_test();

int main() {
  DBG_PRINTF_test();
  return 0;
}

void DBG_PRINTF_test() {
  DBG_PRINTF("This is a test func %d\n", 11);
  return;
}
