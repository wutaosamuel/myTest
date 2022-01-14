#include <iostream>
#include <string>

using namespace std;

void FindPair(const string str, size_t pos, const string left,
              const string right, size_t* result_pos, size_t* result_length) {
  if (result_pos == nullptr || result_length == nullptr) return;
  *result_pos = string::npos;
  *result_length = string::npos;

  size_t found = str.find(left, pos);
  if (found == string::npos) return;

  size_t nextFound = str.find(left, found + left.length());
  size_t pairFound = str.find(right, found + left.length());
  if (pairFound == string::npos) return;
  if (nextFound != string::npos) {
    if (pairFound > nextFound) {
      FindPair(str, found + left.length(), left, right, result_pos,
               result_length);
      return;
    }  // do again
    // if (pairFound <= nextFound) {}// set result
  }
  // if (nextFound == string::npos) {} // set result
  *result_pos = found;
  *result_length = pairFound + right.length() - found;
}

int main() {
  string s = "this is {{ }} test {{ string }}";
  string left = "{{";
  string right = "}}";
  cout << "The string is " << s << endl;

  size_t result_pos, result_length = string::npos;
  FindPair(s, 0, left, right, &result_pos, &result_length);
  cout << "find first {{ }}. pos=" << result_pos << ", length=" << result_length
       << endl;
  if (result_pos == string::npos) return -1;
  FindPair(s, result_pos + left.length(), left, right, &result_pos,
           &result_length);
  cout << "find second {{ }}. pos=" << result_pos
       << ", length=" << result_length << endl;

  return 0;
}