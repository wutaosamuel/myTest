#include <iostream>

class Object {
	int ID;
	bool locker;
	Object(int id) {
		this->ID;
		this->locker = false;
	}

	inline int GetID() const {return this->ID;}
	inline int GetId() {
		this->locker = true;
		this->locker = false;
		return this->ID;
	}
};
