// ble server

#ifndef BLEMYSERVER_H
#define BLEMYSERVER_H

#include <string>
#include <ESP.h>
#include <Arduino.h>
#include <BLEDevice.h>
#include <BLEUtils.h>
#include <BLEServer.h>
#include <pb.h>
#include <pb_common.h>
#include <pb_decode.h>
#include <pb_encode.h>
#include "BleMessage.pb.h"

#define SUCCESS 1
#define FAILURE !SUCCESS

using namespace std;

class BleMyServer
{
  //TODO: set characteristic uuid by authoritative user
public:
  BleMyServer();
  BleMyServer(string, string);
  ~BleMyServer();
  void init();
  void begin();
  int handleOp();

  // setter
  void setWifiSsid(string s) { wifissid = s; }
  void setPassword(string s) { password = s; }
  // set
  int setBuffer(const uint8_t*, size_t);
  // getter
  string getWifiSsid() { return wifissid; }
  string getPassword() { return password; }
  // get
  int getBuffer(uint8_t*, size_t);

private:
  // config
  int op;
  string wifissid;
  string password;
  size_t bufferSize;
  string macAddress;
  uint8_t buffer[256];
  char *SERVICE_UUID = "4fafc201-1fb5-459e-8fcc-c5c9c331914b";
  char *remoteMacUuid = "45678df2-1fd6-45b7-9376-a44282f3b2d5";

  BLEServer *pServer;
  BLEService *pService;
  BleMessage bleMessage;
  BLEAdvertising *pAdvertising;
  BLECharacteristic *remoteCharacteristic;

  int encodeAct();
  int decodeAct();
  string getESP_id();
  void setMessage(int, string, string, string);
};

#endif
