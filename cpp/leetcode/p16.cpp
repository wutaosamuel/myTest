#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

int threeSumClosest(vector<int> &nums, int target)
{
  vector<int> v = nums;
  sort(v.begin(), v.end());
  int min = v[0] + v[1] + v[2];
  int det = abs(min - target);
  for (int i = 0; i < v.size() - 2; i++)
  {
    //cout << i << endl;
    for (int j = i + 1; j < v.size() - 1; j++)
    {
      for (int k = j + 1; k < v.size(); k++)
      {
        int sum = v[i] + v[j] + v[k];
        cout << sum << endl;
        if (det > abs(sum - target))
        {
          min = sum;
          det = abs(sum - target);
        }
      }
    }
  }
  return min;
}

int main()
{
  vector<int> v = {0,0,0,0};
  cout << threeSumClosest(v, 1) << endl;
}