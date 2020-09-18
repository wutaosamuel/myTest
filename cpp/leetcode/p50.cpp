#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

double myPow(double x, int n)
{
  bool positive;
  int positiven;
  double result;

  if (n == 0)
    return 1;
  if (x == 0)
    return 0;
  if (n == 1)
    return x;

  positive = (n >= 0) ? true : false;
  positiven = (positive) ? n : -n;
  result = 1;

  for (int i = 1; i <= positiven; i <<= 1)
  {
    cout << i << endl;
    if ((i & positiven) == i)
      result *= x;
    x *= x;
  }

  return (positive) ? result : double(1) / result;
}

int main()
{
  cout << myPow(2.00, -2) << endl;
}