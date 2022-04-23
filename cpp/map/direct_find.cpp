#include <iostream>
#include <map>

using namespace std;

int main() {
  std::map<int, int> m;
  int array[] = {2, 4, 1, 8, 5};

  for (int i = 0; i < 5; i++) {
    if (m.find(array[i]) != m.end()) {
      std::cout << "already exist" << std::endl;
      return 0;
    }
    m[array[i]] = i;
  }

  std::cout << "m[0] = " << m[0] << std::endl;
  std::cout << "m[10] = " << m[10] << std::endl;
  std::cout << "m[10] will add 0 into m[10], if m[10] not exist " << m[10]
            << std::endl;

  return 0;
}