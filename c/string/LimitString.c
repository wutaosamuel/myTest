#include "LimitString.h"

/*========= utilsLS ========*/
utilsLS_String* utilsLS_NewString(const char* string) {
  size_t size = strlen(string);
  utilsLS_String* _string =
      (utilsLS_String*)malloc(size * sizeof(utilsLS_String));
  for (int i = 0; i < size; i++) {
    _string[i] = malloc(sizeof(char));
    _string[i] = *(string + i);
  }

  return _string;
}
