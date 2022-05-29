#include <iostream>
#include <string>
#include <vector>

using namespace std;

class UserInterface {
 public:
  virtual void Print() = 0;
  virtual UserInterface* GetUser() = 0;
};

class User : public UserInterface {
 public:
  string Name;
  int Age;

  User(string name, int age) {
    this->Name = name;
    this->Age = age;
  };
  inline void Print() {
    cout << "User name: " << this->Name << ", Age: " << this->Age << endl;
  };
  inline UserInterface* GetUser() {
    return new User(this->Name, this->Age);
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
  inline void Print() {
    cout << "Root name: " << this->Name << ", id: " << this->ID << endl;
  };
  inline UserInterface* GetUser() {
    return new Root(this->Name, this->ID);
  };
};

class Manager {
 public:
  vector<UserInterface*> vec;

  inline void Push(UserInterface& user) {
    this->vec.push_back(&user);
  };
  void Print() {
    for(auto it = this->vec.begin(); it != vec.end(); it++)
      (*it)->Print();
  };
};

int main() {
  User* user0 = new User("user0", 20);
  User* user1 = new User("userAdd", 20);
  Root* root0 = new Root("root0", 0);
  Root* root1 = new Root("rootAdd", 1);

  Manager* m = new Manager;
  m->Push(*user0);
  m->Push(*root0);
  cout << "just push back --> work" << endl;
  m->Print();

  cout << endl;
  cout << "return with function --> work" << endl;
  m->Push(*user1->GetUser());
  m->Push(*root1->GetUser());
  m->Print();

  return 0;
}
