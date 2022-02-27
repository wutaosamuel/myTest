#include <iostream>
#include <string>

using namespace std;

void toLower(string& s) {
  for (int i = 0; i < s.length(); i++) s[i] = static_cast<char>(tolower(s[i]));
}

void toUpper(string& s) {
  for (int i = 0; i < s.length(); i++) s[i] = static_cast<char>(toupper(s[i]));
}

int main() {
  string s = "This is test String to Lowercase";
  cout << s << " toLower -> pass" << endl;

  toLower(s);
  cout << s << endl;

  cout << "toUpper -> pass" << endl;
  toUpper(s);
  cout << s << endl;

  return 0;
}
