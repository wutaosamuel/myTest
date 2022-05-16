#include <functional>
#include <iostream>
#include <map>
#include <string>

using namespace std;

typedef function<void()> FuncType;

void pt() { cout << "this pt output" << endl; }

int main() {
  map<int, FuncType> m;

  m[2] = [&]() { pt(); };

  auto it = m.begin();

	// not work with it->second;
  it->second();

  return 0;
}
