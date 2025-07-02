#include "extern.h"

int the_var = 0;
int extern_var = 0;

void print_var(){
	printf("the var: %d\n", the_var);
}

void print_extern_var() {
	printf("extern var: %d\n", extern_var);
}

