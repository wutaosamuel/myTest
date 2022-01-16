#include <iostream>

using namespace std;

// struct defer_dummy {};
// template <class F>
// struct deferrer {
// F f;
// };
// template <class F>
// deferrer<F> operator*(defer_dummy, F f) {
// cout << f << endl;
// return f;
// }

#include "defer.h"

void function1();
void function2();

int main() {
  defer { cout << "function1" << endl; };
	function1();
  defer { cout << "between function1 and function2" << endl; };
	function2();
  defer { cout << "function2" << endl; };

	auto dd = defer_dummy{} *[]{cout << "new one" << endl;};

  return 0;
}

void function1() {
  defer { cout << "function1 1" << endl; };
  defer { cout << "function1 2" << endl; };
}

void function2() {
  defer { cout << "function2 1" << endl; };
}
