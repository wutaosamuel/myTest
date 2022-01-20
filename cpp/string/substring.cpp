#include <iostream>

#include <string>

using namespace std;

int main() {
	string s = "This is a9 test string";
	cout << s << endl;

	string sub = s.substr(9, 0);
	cout << "test 0 length of string --> pass" << endl;
	cout << sub << endl;
	cout << "0 ok" << endl;

	return 0;
}