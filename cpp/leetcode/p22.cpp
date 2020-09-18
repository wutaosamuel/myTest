#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

vector<string> generateParenthesis(int n)
{
  vector<int> left = {0};
  vector<int> right = {0};
  vector<string> result = {""};
  for (int i = 0; i < n; i++)
  {
    for (int j = 0; j < result.size(); j++)
    {
      int l = left[j];
      int r = right[j];
      string s = result[j];
      for (int k = 0; k < 2; k++)
      {
        if (k % 2 == 0)
        {
          s += "(";
          l++;
          if (l > 3)
            continue;
          left[j] = l;
          result[j] = s;
          continue;
        }
        s += ")";
        r++;
        if (r > 3)
          continue;
        if (l - r < 0)
          continue;
        right.push_back(r);
        result.push_back(s);
        left.push_back(l);
      }
    }
  }
  return result;
}

int main()
{
  for (string s : generateParenthesis(3))
  {
    cout << s << endl;
  }
}