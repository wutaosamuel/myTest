#include <iostream>
#include <vector>

using namespace std;

int main() {
  vector<int> vec = {0, 1, 2, 3, 4, 5, 6};
  cout << "empty size: " << vec.size() << endl;
  cout << "empty capacity: " << vec.capacity() << endl;
  
  for (int i=0; i <vec.size(); ++i) {
    cout << " " << endl;
    cout << "index: " << i << endl;
    cout << "value: " << vec[i] << endl;
    if (i & 1 == 1) {
      vec.erase(vec.begin()+i);
      cout << "after erase size: " << vec.size() << endl;
    }
  }
}