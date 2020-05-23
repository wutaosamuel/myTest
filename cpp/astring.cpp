#include<iostream>
#include<string>

using namespace std;

string convert(string s, int numRows) {
  string result = "";
  int index = 0;
  int stringSize = s.size();
  int turnSize = 2 * numRows - 2;
  int remainder = stringSize % turnSize;
  int turns = (remainder == 0) ? stringSize / turnSize : stringSize / turnSize + 1;
  for (int i=0; i < numRows; i++)
  {
    for (int j = 0; j < turns; j++)
    {
      index = j * turnSize + i;
      if (index > stringSize)
        continue;
      result += s[index];
      index = j * turnSize + 2 * numRows - 2 - i;
      if (index > stringSize)
        continue;
      if (i == 0 || i == numRows -1)
        continue;
      result += s[index];
    }
  }
  return result;
}

int main()
{
  string input = "abcdef";
  cout << input.size() << endl;
  cout << input[1] << endl;
  //cout << convert("LEETCODEISHIRING", 4) << endl;
  cout << input.substr(5, 2) << endl;
}