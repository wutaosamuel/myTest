#include "struct.h"

int main() {
	struct Object* obj = NewObject();
	obj->op->print_id(obj->ID);

	free(obj);
}