#include <iostream>
#include <string>

#include "interface.h"

using namespace std;

class User : public UserInterface {
 public:
  string Name;
  int Age;
  User(string name, int age) {
    this->Name = name;
    this->Age = age;
  };

  void Print() {
    cout << "user: " << this->Name << " age: " << this->Age << endl;
  };
};

class Root : public UserInterface {
 public:
  string Name;
  int ID;
  Root(string name, int id) {
    this->Name = name;
    this->ID = id;
  };

  void Print() {
    cout << "root: " << this->Name << " id: " << this->ID << endl;
  };
};

int main() {
	User* user = new User("user0", 18);
	Root* root = new Root("root", 0);

	Pt(*user);
	Pt(*root);

	return 0;
}