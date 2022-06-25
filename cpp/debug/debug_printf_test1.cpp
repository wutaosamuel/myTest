#include <iostream>

#include "debug_printf1.h"


void DBG_PRINTF_test();

int main() {
  DBG_PRINTF_test();
  return 0;
}

void DBG_PRINTF_test() {
  DBG_PRINTF(DBG_ERROR ,"This is a error test func %d\n", 11);
  DBG_PRINTF(DBG_NONE ,"This is a none test func %d\n", 11);
  printf("There is a none test\n");
  DBG_PRINTF(DBG_WARNING ,"This is a warning test func %d\n", 11);
  return;
}