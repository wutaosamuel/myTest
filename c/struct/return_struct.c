#include <stdio.h>
#include <string.h>

struct CharList {
  char Value[252];

  unsigned int offset;
  struct CharList* next;
};

struct CharList NewCharList();
void CharListPrint(struct CharList*);

int main() {
  printf("Start\n");

	struct CharList list = NewCharList();
	CharListPrint(&list);

  return 0;
}

struct CharList NewCharList() {
	struct CharList charList = {
		.Value = "This is the char list",
		.offset = 0,
		.next = NULL,
	};
	charList.offset = strlen(charList.Value);

	return charList;
}

void CharListPrint(struct CharList* list) {
	printf("char list value: %s\n", list->Value);
	printf("char list offset: %d\n", list->offset);
}