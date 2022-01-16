#include <iostream>

using namespace std;

int main() {
  int a = 1;
  int& b = a;
  int c = b;
  int& d = b;

  cout << "a: " << a << " b: " << b << " c: " << c << " d: " << d << endl;
  cout << "a addr: " << &a << " b addr: " << &b << " c addr: " << &c
       << " d addr: " << &d << endl;

  return 0;
}