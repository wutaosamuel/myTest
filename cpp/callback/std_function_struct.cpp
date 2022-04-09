#include <functional>
#include <iostream>

typedef std::function<int(int, int)> add_t;
int Add(int, int);

typedef struct {
  std::string name;
  add_t addFunc;
} object_t;
object_t* newObject();
void TestObject();

int main() { TestObject(); }

int Add(int a, int b) { return a + b; }

object_t* newObject() { return (object_t*)malloc(sizeof(object_t)); }

void TestObject() {
  std::cout << "TestStruct --> " << std::endl;
  object_t* object = newObject();
  object->name = "object 0";
  object->addFunc = Add;
  std::cout << "name: " << object->name << std::endl;
  std::cout << "add result: " << object->addFunc(1, 2) << std::endl;
  std::cout << "" << std::endl;
}