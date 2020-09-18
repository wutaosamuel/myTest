#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

string leftToRight(string l)
{
  if (l == "(")
    return ")";
  if (l == "[")
    return "]";
  if (l == "{")
    return "}";
  return "";
}

bool isValid(string s)
{
  if (s == "")
    return true;
  if (leftToRight(s.substr(0, 1)) == "")
    return false;
  int count = 0;
  string stt = "";
  string end = "";
  for (int i = 0; i < s.size(); i++)
  {
    string l = leftToRight(s.substr(i, 1));
    cout << i << endl;
    cout << l << endl;
    if (l != "")
    {
      stt += s.substr(i, 1);
      end = l + end;
      count++;
      if (i == s.size() -1)
        return false;
      cout << stt+end << endl;
    }
    else
    {
      if (count == 0)
        return false;
      //cout << s.substr(i - count, count + count) << endl;
      if (s.substr(i - count, count + count) == stt + end)
      {
        i = i + count - 1;
        count = 0;
        stt = "";
        end = "";
      }
      else
      {
        return false;
      }
    }
  }
  return true;
}

int main()
{
  cout << isValid("(([]){})") << endl;
}