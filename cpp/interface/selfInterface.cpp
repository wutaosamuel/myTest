#include <iostream>
#include <string>

using namespace std;

class UserInterface {
 public:
  virtual void Print() = 0;
  virtual void Equal(UserInterface& user) = 0;
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
  inline void Equal(User& user) {
    this->Name = user.Name;
    this->Age = user.Age;
  };
};

int main() {
  // User* user = new User("user0", 20);
  cout << "self interface --> Not work" << endl;
  return 0;
}
