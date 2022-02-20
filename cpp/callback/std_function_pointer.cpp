#include <iostream>
#include <functional>

typedef std::function<int(int, int)> add_t;
int Add(int, int);
void printCallback(add_t);
void TestPointer(add_t addFunc);

int main() {
	printCallback(Add);
	TestPointer(Add);
}

int Add(int a, int b) {
	return a+b;
}

void printCallback(add_t addFunc) {
	std::cout << "printCallback --> pass" << std::endl;

	std::cout << "add callback result: " << addFunc(1, 2) << std::endl;
	std::cout << "" << std::endl;
}

void TestPointer(add_t addFunc) {
	std::cout << "TestPointer --> pass" << std::endl;

	add_t addF;
	if (!addF)
		std::cout<< "true, function is not callable yet" << std::endl;
	
	addF = addFunc;
	if (addF) {
		std::cout<< "true, function is called" << std::endl;
		std::cout << "add callback result: " << addFunc(1, 2) << std::endl;
	}
	std::cout << "" << std::endl;
}