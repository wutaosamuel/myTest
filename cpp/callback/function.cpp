#include <iostream>
#include <functional>

typedef std::function<int(int, int)> add_t;
int Add(int, int);
void printCallback(add_t);

int main() {
	printCallback(Add);
}

int Add(int a, int b) {
	return a+b;
}

void printCallback(add_t addFunc) {
	std::cout << "add callback result: " << addFunc(1, 2) << std::endl;
}