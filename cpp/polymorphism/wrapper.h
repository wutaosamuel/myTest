#ifndef INTERFACE_H
#define INTERFACE_H

class UserInterface {
 public:
  virtual void Print() = 0;
};

class TaskInterface {
 public:
  virtual void GetTask() = 0;
};

class TaskWrapper : public UserInterface, TaskInterface {};

inline void Pt(UserInterface& user) { user.Print(); };
inline void Pt(TaskInterface& task) { task.GetTask(); };

#endif