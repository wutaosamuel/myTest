#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Object {
 public:
  string Name;
  int ID;

  Object(string name, int id) {
    this->Name = name;
    this->ID = id;
  };

  string ToString() { return this->Name + ": " + to_string(this->ID); };
};

int main() {
  vector<int> v{1, 2, 3, 4, 5, 6, 7, 8, 9};

  cout << "iterator start" << endl;
  for (auto i = v.begin(); i != v.end(); i++)
    cout << "iterator: " << *i << endl;

  cout << "number start" << endl;
  for (int i = 0; i < v.size(); i++) cout << "number: " << v[i] << endl;

  vector<Object *> vec;
  vec.push_back(new Object("object1", 1));
  vec.push_back(new Object("object2", 2));
  vec.push_back(new Object("object3", 3));
  vec.push_back(new Object("object4", 4));
  cout << "object iterator start" << endl;
  for (auto it = vec.begin(); it != vec.end(); it++) {
    cout << "" << endl;
    printf("pointer -> %p\n", *it);
    cout << *it << endl;
    cout << (*it)->ToString() << endl;
  }

  return 0;
}