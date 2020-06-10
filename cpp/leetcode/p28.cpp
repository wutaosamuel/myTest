#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

int strStr(string haystack, string needle)
{
  if (needle == "")
    return 0;
  if (haystack == "")
    return -1;
  int pos = -1;
  for (int i = 0; i < haystack.size() - needle.size() + 1; i++)
  {
    cout << haystack.substr(i, needle.size()) << endl;
    cout << needle << endl;
    if (haystack.substr(i, needle.size()) == needle)
    {
      pos = i;
      break;
    }
  }
  return pos;
}

int main()
{
  string h = "hello";
  string n = "ello";
  cout << strStr(h, n) << endl;
}