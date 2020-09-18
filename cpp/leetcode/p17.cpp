#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

string letters[10] = {
    "",
    "",
    "abc",
    "def",
    "ghi",
    "jkl",
    "mno",
    "pqrs",
    "tuv",
    "wxyz"};

int stringToInt(string s)
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
  else
    return -1;
}

vector<string> combination2(vector<string> &vec, int num)
{
  string letter = letters[num];
  vector<string> result;
  for (int i = 0; i < vec.size(); i++)
  {
    cout << i << endl;
    for (int j = 0; j < letter.size(); j++)
    {
      string tmp = vec[i];
      cout << tmp << endl;
      result.push_back(vec[i] + letter.substr(j, 1));
    }
  }
  return result;
};

vector<string> letterCombinations(string digits)
{
  vector<string> result;
  if (stringToInt(digits.substr(0, 1)) == -1)
    return result;
  string letter0 = letters[stringToInt(digits.substr(0, 1))];
  for (int j = 0; j < letter0.size(); j++)
  {
    result.push_back(letter0.substr(j, 1));
  }
  for (int i = 1; i < digits.size(); i++)
  {
    result = combination2(result, stringToInt(digits.substr(i, 1)));
  }
  return result;
}

int main()
{
  vector<string> vec = letterCombinations("23");
  for (string s : vec)
    cout << s << endl;
}