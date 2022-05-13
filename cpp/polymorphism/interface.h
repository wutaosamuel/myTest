#ifndef INTERFACE_H
#define INTERFACE_H

class UserInterface {
 public:
  virtual void Print() = 0;
};

inline void Pt(UserInterface& user) { user.Print(); };

#endif