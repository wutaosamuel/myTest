#include "extern.h"

extern int the_var;

void assign_var(int);
void assign_extern_var(int);

int main() {
	printf("initial value\n");
	print_var();
	printf("\n");

	printf("assign value in main function\n");
	the_var = 1;
	print_var();
	printf("\n");

	printf("assign value in assign_var function\n");
	assign_var(2);
	printf("\n");

	printf("initial value\n");
	print_extern_var();	
	printf("\n");

	printf("assign value in main function\n");
	extern_var = 1;
	print_extern_var();
	printf("\n");

	printf("assign value in assign_extern_var function\n");
	assign_extern_var(2);
	printf("\n");

	return 0;
}

void assign_var(int value) {
	the_var = value;
	print_var();
}

void assign_extern_var(int value) {
	extern_var = value;
	print_extern_var();
}
