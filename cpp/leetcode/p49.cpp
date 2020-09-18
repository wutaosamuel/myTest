#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

vector<vector<string>> groupAnagrams(vector<string> &strs)
{
  if (strs.size() == 0)
    return vector<vector<string>>{{}};
  vector<vector<string>> result = {vector<string>{strs[0]}};
  for (int i = 1; i < strs.size(); i++)
  {
    bool canRow = true;
    for (int j = 0; j < result.size(); j++)
    {
      string rowfirst = result[j][0];
      string tmp = "";
      for (int k = 0; k < rowfirst.size(); k++)
      {
        if (strs[i].find(rowfirst.substr(k, 1)) != string::npos)
        {
          tmp += rowfirst.substr(k, 1);
        }
        else
        {
          tmp = "";
          canRow = true;
          break;
        }
      }
      cout << "rowfirst: " << rowfirst << endl;
      cout << "tmp: " << tmp << endl;
      cout << i << ": " << strs[i] << endl;
      if (rowfirst == tmp && rowfirst != "")
      {
        result[j].push_back(strs[i]);
        canRow = false;
        break;
      }
      if (rowfirst == tmp && strs[i] != "")
      {
        canRow = true;
        break;
      }
      if (rowfirst == tmp && strs[i] == "")
      {
        result[j].push_back(strs[i]);
        canRow = false;
        break;
      }
    }
    if (canRow)
    {
      vector<string> res = {strs[i]};
      result.push_back(res);
    }
  }
  return result;
}

int main()
{
  //vector<string> vec = {"","dans"};
  vector<string> vec = {"and","dans"};
  int i = 0;
  for (auto a : groupAnagrams(vec))
  {
    cout << "array " << i << ": " << endl;
    i++;
    for (string s : a)
      cout << s << endl;
  }
}