#include <iostream>
#include <vector>

using namespace std;

int main() {
  vector<int> vec = {0, 1, 2, 3, 4, 5, 6};
  cout << "empty size: " << vec.size() << endl;
  cout << "empty capacity: " << vec.capacity() << endl;

  int i = 0;
  for (vector<int>::iterator it = vec.begin(); it != vec.end(); /* ++it */) {
    cout << " " << endl;
    cout << "index: " << i << endl;
    cout << "value: " << *it << endl;
    if (i & 1 == 1) {
      vec.erase(it);
      cout << "after erase size: " << vec.size() << endl;
    } else {
      ++it;
    }
    i++;
  }
}