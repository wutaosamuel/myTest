#ifndef MY_PEOPLE_H
#define MY_PEOPLE_H

#include <stdio.h>
#include <stdlib.h>

#include "interface.h"

struct People {
  int Age;

  const struct StructInterface* interface;
};

static inline void printPeople(void* people) {
  struct People* p = (struct People*)people;
  printf("people: \n");
  printf("the age is %d\n", p->Age);
}

static const struct StructInterface peopleInterface = {
    .Print = printPeople,
};

struct People* NewPeople() {
  struct People* people = (struct People*)malloc(sizeof(struct People));
  people->Age = 18;
  people->interface = &peopleInterface;

  return people;
}

#endif