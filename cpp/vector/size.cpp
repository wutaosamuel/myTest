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

  cout << "" << endl;

  v.push_back(2);
  cout << "push back size: " << v.size() << endl;
  cout << "push back cap: " << v.capacity() << endl;
  cout << "" << endl;
  v.pop_back();
  cout << "pop back size: " << v.size() << endl;
  cout << "pop back cap: " << v.capacity() << endl;

  cout << "" << endl;

  vector<int> *p;
  p = new vector<int>{};
  cout << "point empty size: " << p->size() << endl;
  cout << "point empty capacity: " << p->capacity() << endl;
  cout << "point empty max_size: " << p->max_size() << endl;

  p->push_back(1);
  cout << "point 1 size: " << p->size() << endl;
  cout << "point 1 capacity: " << p->capacity() << endl;
  cout << "point 1 max_size: " << p->max_size() << endl;
}