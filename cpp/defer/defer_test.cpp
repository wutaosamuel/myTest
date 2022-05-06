#include "defer.h"

#include <iostream>
#include <string>

using namespace std;

inline void ReturnDeferTest();

inline void returnDefer();

int main() {
	cout << "Start to test" << endl;
	ReturnDeferTest();

	return 0;
}

inline void ReturnDeferTest() {
  defer { cout << "there is defer test" << endl; };
	return returnDefer();
}

inline void returnDefer() {
  defer { cout << "there is return returnDerfer func" << endl; };
  return;
}
