#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

vector<int> searchRange(vector<int> &nums, int target)
{
  vector<int> result = {-1, -1};
  if (nums.size() == 0)
    return result;
  if (nums.size() == 1)
  {
    if (nums[0] == target)
    {
      result[0] = 0;
      result[1] = 0;
    }
    return result;
  }
  if (nums[0] > target)
    return result;
  if (nums.back() < target)
    return result;

  int size = nums.size();
  int mid = size / 2;
  int leftmid = mid - 1;
  int rightmid = (size == 2) ? mid : mid + 1;
  result[0] = result[1] = mid;
  if (nums[0] == target) {
    mid = leftmid = 0;
    rightmid = 1;
    result[0] = result[1] = mid;
  }
  if (nums[size - 1] == target) {
    mid = rightmid = size - 1;
    leftmid = mid - 1;
    result[0] = result[1] = mid;
  }
  for (;;)
  {
    size /= 2;
    cout << ":" << endl;
    cout << size << endl;
    cout << leftmid << endl;
    cout << rightmid << endl;
    if (nums[mid] == target)
    {
      if (nums[leftmid] == target)
      {
        result[0] = leftmid;
        leftmid -= size / 2;
        cout << "left: " << leftmid << endl;
      }
      else
      {
        leftmid += size / 2;
      }

      if (nums[rightmid] == target)
      {
        result[1] = rightmid;
        rightmid += size / 2;
        cout << "right: " << rightmid << endl;
      }
      else
      {
        rightmid -= size / 2;
      }
      if (size == 0)
        return result;
      continue;
    }

    cout << "size: " << size << endl;
    cout << "mid: " << mid << endl;
    if (nums[mid] < target)
    {
      mid += size / 2;
    }
    else 
    {
      mid -= size / 2;
    }
    cout << "size: " << size << endl;
    cout << "mid: " << mid << endl;
    leftmid = (mid - 1 < 0) ? mid : mid - 1;
    rightmid = (mid + 1 > nums.size()) ? mid : mid + 1;
    result[0] = result[1] = mid;
    if (size == 0)
      break;
  }
  return vector<int>{-1, -1};
}

int main()
{
  vector<int> vec = {3, 3, 3, 3};
  for (int a : searchRange(vec, 3))
    cout << a << endl;
}