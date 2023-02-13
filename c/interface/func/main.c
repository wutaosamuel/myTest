#include <stdio.h>
#include "func_interface.h"

void printFunc1() {
	printf("Hello world from printFunc1\n");
}

int main() {
	printf("Start\n");

	FuncPrintPtr pFunc = GetFuncPtr();
	printf("pFunc %p\n", pFunc);
	if (!pFunc) 
		SetFuncPrintPtr(printFunc1);
	printf("set pFunc %p\n", pFunc);
	pFunc = GetFuncPtr();
	printf("get pFunc %p\n", pFunc);
	pFunc();
  
	return 0;
}