#include <inttypes.h>
#include <stdint.h>
#include <stdio.h>

struct StructBit {
  unsigned First : 4;
  unsigned Second : 8;
  unsigned Third : 4;
};

void StructBitPrint(struct StructBit* bit, size_t size) {
  printf("First: %u, Second: %u, Third: %u\n", bit->First, bit->Second,
         bit->Third);
  printf("size of struct: %zu bytes\n", sizeof(*bit));
  for (size_t i = 0; i < size; i++) {
    printf("bytes %zu: 0x%" PRIXPTR "\n", i, bit + i);
  }
  printf("cannot print address of bit field in struct, such as &(bit->First).");
  printf("because bit is small.\n");
}

void ArraySequencePrint(int* array, size_t size) {
  printf("size %zu\n", sizeof(array));
  for (size_t i = 0; i < size; i++) {
    printf("array bytes %zu: 0x%" PRIXPTR " %d\n", i, array + i, *(array + i));
  }
}

int main() {
  printf("print bytes address sequence\n");
  int array[9] = {9, 9, 9, 9, 9, 9};
  ArraySequencePrint(array, 9);
  printf("\n");

  printf("struct with bit address sequence\n");
  struct StructBit bit = {
      .First = 0b0001, .Second = 0b10000001, .Third = 0b1000};
  StructBitPrint(&bit, 3);

  return 0;
}