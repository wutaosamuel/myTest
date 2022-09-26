#include <stdio.h>

struct MixArray {
	char Char[256];
	int Int[256];
};

int main() {
	// TODO: memset in mArr.Int
	struct MixArray mArr = {
		.Char = "",
		.Int = {0},
	};
}
