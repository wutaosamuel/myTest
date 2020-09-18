#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

int search(vector<int> &nums, int target)
{
  if (nums.size() == 0)
    return -1;
  if (nums[0] == target)
  {
    return 0;
  }
  else if (nums[0] > target)
  {
    for (int i = 1; i < nums.size(); i++)
    {
      if (nums[i] == target)
        return i;
    }
  }
  else
  {
    for (int j = nums.size() - 1; j > 0; j--)
    {
      if (nums[j] == target)
        return j;
    }
  }

  return -1;
}

int main()
{
  vector<int> vec = {1, 3};
  cout << search(vec, 3) << endl;
}