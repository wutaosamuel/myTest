#include "static_inline.h"

int main() {
  Object o = {
      .Type = 0,
      .ID = -1,
  };

  printf("\n");
  printf("test static inline now --> works\n");
	s_Print(&o);

  return 0;
}