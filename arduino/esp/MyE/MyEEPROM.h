// MyEEPROM.h

#ifndef MYEEPROM_H
#define MYEEPROM_H

#include <stdio.h>
#include <cstring>
#include <string>
#include <EEPROM.h>

#include <pb.h>
#include <pb_common.h>
#include <pb_encode.h>
#include <pb_decode.h>
#include "SaveOnServer.pb.h"

using namespace std;

class MyEEPROM
{
public:
  MyEEPROM();
  ~MyEEPROM() {}

  // settter
  void setAddr(int i) { addr = i; }
  // getter
  int getAddr() { return addr; }

private:
  int addr; // start address
};

#endif