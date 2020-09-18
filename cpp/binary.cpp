#include <iostream>

using namespace std;

int main()
{
  int n = 10;
  int b = 1;

  for (;;)
  {
    int s = n & b;
    b = b << 1;
    cout << "n & b: " << s << endl;
    cout << "b: " << b << endl;
    if (b > n)
      break;
  }
}