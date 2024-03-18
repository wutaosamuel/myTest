#include <stdio.h>
#include <stdlib.h>

struct List {
  int Data;
  struct List* next;
};

static void PrintList(struct List*);

int main() {
  struct List head;
  struct List *list1, *list2;
  head.Data = 0;
  head.next = NULL;

  list1 = (struct List*)malloc(sizeof(struct List));
  list2 = (struct List*)malloc(sizeof(struct List));
  list1->Data = 1;
  list1->next = NULL;
  list2->Data = 2;
  list2->next = NULL;

  // list connect
  head.next = list1;
  list1->next = list2;
  list2->next = NULL;

  PrintList(&head);
}

static void PrintList(struct List* list) {
  int count = 0;

  while (list != NULL) {
    printf("%dth element: %d\n", count + 1, list->Data);

    list = list->next;
    count++;
  }
}