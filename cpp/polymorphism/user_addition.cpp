#include <iostream>
#include <string>

#include "interface.h"

using namespace std;

class Addition {
 public:
  virtual void APrint() = 0;
};

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

class Man : public UserInterface, public Addition {
 public:
  string Name;
  int Age;
  Man(string name, int age) {
    this->Name = name;
    this->Age = age;
  }

  void Print() {
    cout << "man name: " << this->Name << " age: " << this->Age << endl;
  }
  void APrint() { cout << "addition printing" << endl; }
};

void APrint(Addition& add) { add.APrint(); }

int main() {
  User* user = new User("user0", 18);
  Root* root = new Root("root", 0);

  Pt(*user);
  Pt(*root);

  Man* man = new Man("man", 18);
  Pt(*man);
  APrint(*man);

  return 0;
}