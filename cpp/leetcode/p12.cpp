#include <iostream>
#include <string>
#include <map>

using namespace std;

string intToRoman(int num)
{
  if (num < 1 || num > 3999)
    return "";
  map<int, string> t = {
      {1, "I"}, {4, "IV"}, {5, "V"}, {9, "IX"}, {10, "X"}, {40, "XL"}, {50, "L"}, {90, "XC"}, {100, "C"}, {400, "CD"}, {500, "D"}, {900, "CM"}, {1000, "M"}};
  string result = "";
  int devider = 10;
  for (;;)
  {
    int devider10 = devider / 10;
    int remainder = num % devider;
    num -= remainder;
    for (;;)
    {
      //cout << remainder << endl;
      if (remainder <= 0)
        break;
      if (remainder == 4 * devider10)
      {
        result = t[4 * devider10] + result;
        remainder -= 4 * devider10;
      }
      else if (remainder == 5 * devider10)
      {
        result = t[5 * devider10] + result;
        remainder -= 5 * devider10;
      }
      else if (remainder == 9 * devider10)
      {
        result = t[9 * devider10] + result;
        remainder -= 9 * devider10;
      }
      else
      {
        result = t[devider10] + result;
        remainder -= devider10;
      }
      cout << remainder << endl;
    }
    if (num <= 0)
      break;
    devider *= 10;
  }

  return result;
}

int main() 
{
  int a = 10;
  //cout << intToRoman(a) << endl;
  map<string, int> t = {{"a", 1},{"b", 2}};
  cout << t["c"] << endl;
}