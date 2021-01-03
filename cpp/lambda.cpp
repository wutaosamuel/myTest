#include <iostream>
#include <functional>

int main()
{
  auto printHello = [] { std::cout << "Hello, World!" << std::endl; };
  auto sum = [](int a, int b) { return a + b; };

  std::function<void()> printLambda = [] {
    std::cout << "Hello, Lambda!" << std::endl;
  };
  std::function<int(int, int)> multiply = [](int a, int b) { return a * b; };

  printHello();
  printLambda();
  std::cout << "Sum 1 + 2: " << sum(1, 2) << std::endl;
  std::cout << "Multiply 2 * 6: " << multiply(2, 6) << std::endl;

  return 1;
}