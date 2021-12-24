#include <iostream>
#include <vector>

using namespace std;

class Object {
 public:
  string Name;
 public:
  Object(string name) {this->Name=name;};
};

int main() {
  vector<Object*> vec;

  cout << "size: " << vec.size() << endl;
  cout << "empty: " << vec.empty() << endl;
	cout << "" << endl;

	vec.push_back(new Object("o1"));
  cout << "size: " << vec.size() << endl;

	cout<< "" <<endl;
	cout<< "Add a nullptr -> can be added" <<endl;
	vec.push_back(nullptr);
  cout << "size: " << vec.size() << endl;
}