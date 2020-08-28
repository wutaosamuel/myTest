#include <iostream>
#include <vector>

using namespace std;

int main()
{
  vector<int> v{1, 2, 3, 4, 5, 6, 7, 8, 9};
  cout << "iterator start" << endl;
  for(auto i = v.begin(); i != v.end(); i++)
    cout << "iterator: " << *i << endl;
  cout << "number start" << endl;
  for(int i=0; i < v.size(); i++)
    cout << "number: " << v[i] << endl;
}