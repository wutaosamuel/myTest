#include <iostream>
#include <vector>

using namespace std;

void print(vector<int> a, vector<int> b)
{
  int length = a[0];
  int times = b[0];
  vector<int> l;
  for (int i = 0; i < length; i++)
    l.push_back(0);
  for (int j = 1; j <= times; j++)
  {
    for (int k = a[j] - 1; k < b[j]; k++)
    {
      l[k]++;
    }
  }
  for (int z = 0; z < length; z++)
  {
    if (z == 0)
    {
      cout << l[0];
    }
    else if (z == length - 1)
    {
      cout << " " << l[z] << endl;
    }
    else
    {
      cout << " " << l[z];
    }
  }
}

int main()
{
  int num;
  int a_num, b_num;
  vector<int> n_num;
  vector<vector<int>> a, b;
  cin >> num;
  for (int i = 0; i < num; i++)
  {
    int count = 0;
    a.push_back(vector<int>{});
    b.push_back(vector<int>{});
    while (cin >> a_num >> b_num)
    {
      a[i].push_back(a_num);
      b[i].push_back(b_num);
      if (b[i][0] == count)
        break;
      count++;
    }
    print(a[i], b[i]);
  }
  return 0;
}