#include <iostream>
#include <functional>
#include <atomic>

#include <libgo/coroutine.h>

void foo()
{
  std::cout << "Foo function" << std::endl;
}

int main()
{
  std::cout << "Execute foo func" << std::endl;
  go foo;
  
  std::function<void()> helloLambda = [] {
    std::cout << "Hello, Lambda!" << std::endl;
  };
  go helloLambda;
  // not work for ```go helloLambda();```
  go [] {
    std::cout << "Hello, World!" << std::endl;
  };

  // run
  std::atomic<int> done{0};
  int task_num = 0;

  auto compute = [](std::atomic<int&> done) {
    int v = 0;
    for (int i=0; i < 20000000; i++) {
      v += i;
    }

    done++;
  };

  // run 10 tasks

  return 0;
}