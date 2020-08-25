#include <iostream>
#include <vector>

using namespace std;

int main() 
{
  vector<int> v;
  cout << "empty size: " << v.size() << endl;
  cout << "empty capacity: " << v.capacity() << endl;
  cout << "empty max_size: " << v.max_size() << endl;

  v.push_back(1);
  cout << "1 size: " << v.size() << endl;
  cout << "1 capacity: " << v.capacity() << endl;
  cout << "1 max_size: " << v.max_size() << endl;
}