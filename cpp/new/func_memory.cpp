#include <iostream>
#include <string>

using namespace std;

class Object {
 public:
  string Name;
  Object(string name) { this->Name = name; };
};

Object* New() { return new Object("name"); }

void NewObject(Object* obj) { obj = new Object("new"); };

int main() {
  Object* o1 = nullptr;
  cout << "return by function --> works" << endl;
  printf("o1 before pointer %p\n", o1);
  o1 = New();
  printf("o1 after pointer %p\n", o1);

  cout << endl;
  cout << "return by function --> not work" << endl;
  Object* o2 = nullptr;
  printf("o2 before pointer %p\n", o2);
  NewObject(o2);
  printf("o2 after pointer %p\n", o2);

  return 0;
}