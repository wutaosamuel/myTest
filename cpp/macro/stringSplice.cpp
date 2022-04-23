#include <iostream>
#include <string>

#define STR_1 "string1"
#define STR_2 "string2"
#define STRING_1 STR_1 " " STR_2
#define STRING_2(s1, s2) s1#s2

int main() { 
  std::cout << "String 1 -> " << STRING_1 << std::endl;
  std::cout << "String 2 -> " << STRING_2(STR_1, STR_2) << std::endl;
  return 0; 
}
