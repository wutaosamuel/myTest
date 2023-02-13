#include "func_interface.h"

static FuncPrintPtr publishFunc;

void SetFuncPrintPtr(FuncPrintPtr func) {
	publishFunc = func;
}

FuncPrintPtr GetFuncPtr() {
	return publishFunc;
}