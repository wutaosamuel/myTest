#include <iostream>
#include <string>

#include "wrapper.h"

using namespace std;

class User : public UserInterface {
 public:
  string Name;
  int Age;
  User(string name, int age) {
    this->Name = name;
    this->Age = age;
  };

  inline void Print() {
    cout << "user: " << this->Name << " age: " << this->Age << endl;
  };
};

class UserWrapper: public UserInterface, TaskInterface {};

class Root: public UserWrapper {
  string Name;
  int ID;
};

int main() {
  User* user = new User("user0", 18);

  Pt(*user);

  return 0;
}