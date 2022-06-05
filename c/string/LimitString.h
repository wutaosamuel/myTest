#ifndef LIMIT_STRING_H
#define LIMIT_STRING_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef char* utilsLS_String;

inline utilsLS_String* utilsLS_NewStr();
utilsLS_String* utilsLS_NewString(const char* string);
inline void utilsLS_FreeString(utilsLS_String* string);

inline size_t utilsLS_Size(utilsLS_String* string);

/*========= utilsLS ========*/
inline utilsLS_String* utilsLS_NewStr() {
  utilsLS_String* string = (utilsLS_String*)malloc(sizeof(utilsLS_String));
  string[0] = malloc(sizeof(char));
  string[0] = '\0';

  return string;
}

void utilsLS_FreeString(utilsLS_String* string) {
  free(string->Chars);
  free(string);
  string = NULL;
}

size_t utilsLS_Size(utilsLS_String* string) { return sizeof(string); }

#endif