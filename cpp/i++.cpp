#include <iostream>

using namespace std;

int main()
{
  int a = 1;
  int b = 1;
  int c = 1;
  int d = 1;
  int e1 = 1;
  int e2 = 1;
  int f1 = 1;
  int f2 = 1;
  int g = 1;
  int h = 3;

  cout << "a++: " << a++ << endl;
  cout << "++b: " << ++b << endl;
  cout << "(++c) + (++c): " << (++c) + (++c) << endl;
  cout << "(d++) + (d++): " << (d++) + (d++) << endl;
  cout << "(++e1) + (++e2): " << (++e1) + (++e2) << endl;
  cout << "(f1++) + (f2++): " << (f1++) + (f2++) << endl;
  cout << "++g 1: " << ++g << endl;
  cout << "++g 2: " << ++g << endl;
  cout << "++h: " << ++h << endl;
  h = 3;
  cout << "++h + ++h: " << (++h) + (++h) << endl;
  h = 3;
  cout << "++h + ++h + ++h: " << (++h) + (++h) + (++h) << endl;
  h = 3;
  cout << "++h + ++h + ++h: " << (++h) + (++h) + (++h) + (++h)<< endl;

  return 0;
}