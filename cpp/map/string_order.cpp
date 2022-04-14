#include <iostream>
#include <string>
#include <map>

using namespace std;

int main() {
	std::map<std::string, int> m;
	std::string array[] = {"c", "1", "bac", "2", "a"};

	for(int i=0; i<5; i++) {
		if (m.find(array[i]) != m.end()) {
			std::cout << "already exist" << std::endl;
			return 0;
		}
		m[array[i]] = i;
	}

	std::cout << "ordered:" << std::endl;
	for(auto it = m.begin(); it != m.end(); it++) {
		std::cout << it->first << " -> " << it->second << std::endl;
	}

	return 0;
}