#ifndef MY_INTERFACE_H
#define MY_INTERFACE_H

struct StructInterface {
	void (*Print)(void* self);
};

#endif