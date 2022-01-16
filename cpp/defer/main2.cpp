#include <iostream>

using namespace std;

struct defer_dummy {};
template <class F>
struct deferrer {
  F f;
};
template <class F>
deferrer<F> operator*(defer_dummy, F f) {
  return {f};
}

int main() { return 0; }
