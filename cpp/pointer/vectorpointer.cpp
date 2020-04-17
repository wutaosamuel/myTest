#include <vector>
#include <iostream>

using namespace std;

void printvec(int &p)
{
  cout << "value: " << p << endl;
  p = 10;
}

int main()
{
  vector<int> vec;
  //int *p = new int;
  vec.push_back(1);
  vec.push_back(2);
  for (int i = 0; i < vec.capacity(); ++i)
  {
    //p = &vec[i];
    printvec(vec[i]);
  }
  for (int i = 0; i < vec.capacity(); ++i)
  {
    //p = &vec[i];
    printvec(vec[i]);
  }
  return 0;
}
