#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
  vector<int> v = {1, 3, 2,5,7, 3};
  sort(v.begin(), v.end());
  v.erase(unique(v.begin(), v.end()), v.end());
  for(int s: v) {
    cout << s << endl;
  }
}