#include <stdio.h>

struct FuncStruct {
	int (*numPtr)();
	void* Func;
	int id;
	struct FuncStruct* next;
};

static int flag = 0;

static int appNum();
struct FuncStruct appFunc = {appNum, NULL};

static int appNum() {
	flag = 1;
	printf("app num func print\n");
	return 100;
}

int main() {
  printf("flag value %d\n", flag);
	if (appFunc.numPtr) appFunc.numPtr();
	printf("appFunc numPtr = %d\n", appFunc.numPtr);
	printf("appFunc Func = %s\n", appFunc.Func);
	printf("appFunc id = %d\n", appFunc.id);
	printf("appFunc next = %s\n", appFunc.next);

  printf("flag value %d\n", flag);
}