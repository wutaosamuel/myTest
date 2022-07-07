#include "interface.h"
#include "object.h"
#include "people.h"

int main() {
	struct Object* obj = NewObject();
	struct People* peo = NewPeople();

	obj->interface->Print((void*)obj);
	printf("\n");
	peo->interface->Print((void*)peo);

	free(obj);
	free(peo);
}