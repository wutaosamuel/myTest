#include <iostream>

int maxProfit(int* prices, int priceSize);

int main() {
  int input[] = {1, 2, 3, 0, 2};
  int output = 4;

	int o = maxProfit(input, 5);
	std::cout << o << std::endl;

  return 0;
}

int maxProfit(int* prices, int priceSize) {
  int result = 0;

  for (int i = 0; i < priceSize; i++) {
		bool end = false;
    for (int j = i; j < priceSize; j++) {
      if (*(prices + j) < *(prices + i)) {
        result += *(prices + j - 1) - *(prices + i);
        i = j;
				break;
      }
			end = true;
    }

		if (end) {
			result += *(prices + priceSize -1) - *(prices + i);
			return result;
		}
  }

  return result;
}
