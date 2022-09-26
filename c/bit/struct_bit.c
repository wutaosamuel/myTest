#include <stdio.h>

struct StructBit {
	unsigned First : 1;
	unsigned Second : 1;
	unsigned Third : 32;
};

struct CharBit {
	char char0;
	char char1;
	char char2;
};

int main() {
	char Char;
	unsigned int ui;
	struct StructBit sb;
	struct CharBit cb;

	printf("Char size: %ld\n", sizeof(Char));
	printf("uint size: %ld\n", sizeof(ui));
	printf("bits size: %ld\n", sizeof(sb));
	printf("char size: %ld\n", sizeof(cb));

	return 0;
}