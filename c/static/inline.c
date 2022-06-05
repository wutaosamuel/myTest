#include "header.h"

static inline void s_print(Object* obj) {
  printf("Object type: %d, ID: %d\n", obj->Type, obj->ID);
}

int main() {
  Object o = {
      .Type = 0,
      .ID = -1,
  };

  printf("test inline now --> not works\n"); 
  // i_Print(&o);

  printf("\n");
  printf("test extern inline now --> works\n"); 
  e_Print(&o);

  printf("\n");
  printf("test static inline now --> works\n");
  s_Print(&o);

  return 0;
}

void s_Print(Object* obj) { s_print(obj); }