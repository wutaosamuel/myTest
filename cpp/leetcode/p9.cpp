#include <iostream>

using namespace std;

bool isPalindrome(int x)
{
  int input = x;
  int result = 0;
  if (x < 0)
    return false;
  for (;;)
  {
    if (input < 10)
    {
    result = result * 10 + input;
      break;
    }
    result = result * 10 + input % 10;
    input /= 10;
    cout << result << endl;
  }
  cout << result << endl;
  if (x != result)
    return false;
  return true;
}

int main()
{
  int a = 123454321;
  cout << isPalindrome(a) << endl;
}