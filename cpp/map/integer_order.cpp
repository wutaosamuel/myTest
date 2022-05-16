#include <iostream>
#include <map>

using namespace std;

int main() {
	std::map<int, int> m;
	int array[] = {2, 4, 1, 8, 5, -1, -4, -9};

	for(int i=0; i<8; i++) {
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