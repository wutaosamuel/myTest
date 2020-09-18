#include <iostream>
#include <math.h>

using namespace std;

const unsigned int INI = (int)pow(2, 31);

int AtoiSingle(string s)
{
  if (s == "0")
    return 0;
  else if (s == "1")
    return 1;
  else if (s == "2")
    return 2;
  else if (s == "3")
    return 3;
  else if (s == "4")
    return 4;
  else if (s == "5")
    return 5;
  else if (s == "6")
    return 6;
  else if (s == "7")
    return 7;
  else if (s == "8")
    return 8;
  else if (s == "9")
    return 9;
  else if (s == " ")
    return 10;
  else if (s == "-")
    return 11;
  else
    return -1;
}
int myAtoi(string str)
{
  unsigned int result = 0;
  bool accumulation = false;
  bool positive = true;
  for (int i = 0; i < str.size(); i++)
  {
    int tmp = AtoiSingle(str.substr(i, 1));

    if (tmp == 11 && !positive)
      break;
    if (!accumulation && result == 0)
    {
      if (tmp == 10)
        continue;
      if (tmp == -1)
        return 0;
    }
    if (accumulation)
    {
      if (tmp == 10 || tmp == 11 || tmp == -1)
        break;
    }

    if (tmp == 11)
    {
      positive = false;
      tmp = 0;
    }
    accumulation = true;
    result = result * 10 + tmp;
    cout << "result: " << endl;
    cout << result << endl;
    cout << tmp << endl;
    if (positive)
    {
      if (result > INI - 1)
        return (int)(INI - 1);
    }
    else
    {
      if (result > INI)
        return -1 * (int)INI;
    }
  }
  return positive ? (int)result : -1 * (int)result;
}

int main()
{
  string s = "-91283472332";
  cout << myAtoi(s) << endl;
}