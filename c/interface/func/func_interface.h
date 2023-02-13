#ifndef MY_FUNC_INTERFACE_H
#define MY_FUNC_INTERFACE_H

typedef void (*FuncPrintPtr)();

void SetFuncPrintPtr(FuncPrintPtr func);
FuncPrintPtr GetFuncPtr();

#endif