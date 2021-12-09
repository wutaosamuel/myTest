// BLEMyClient.h

#ifndef BLEMYCLIENT_H
#define BLEMYCLIENT_H

#include <string>

#include <Arduino.h>
#include <BLEDevice.h>
//#include <BLEScan.h>
using namespace std;

 /**
   * Called for each advertising BLE server.
   */
class MyAdvertisedDeviceCallbacks: public BLEAdvertisedDeviceCallbacks 
{
public:
  boolean doConnect = false;
  boolean connected = false;
  boolean doScan = false;
  BLEAdvertisedDevice *pDevice;
  // setter
  void setDoConnect(boolean b) { doConnect = b; }
  void setConnected(boolean b) { connected = b; }
  void setDoScan(boolean b) { doScan = b; }
  // set
  void setServiceUuid(string s) { bleServiceUuid = BLEUUID(s); }
  // getter
  boolean getDoConnect() { return doConnect; }
  boolean getConnected() { return connected; }
  boolean getDoScan() { return doScan; }
  // get
  void getBLEDevice(BLEAdvertisedDevice* &d) { d = pDevice; }

  void onResult(BLEAdvertisedDevice advertisedDevice) {
    Serial.print("BLE Advertised Device found: ");
    Serial.println(advertisedDevice.toString().c_str());

    // We have found a device, let us now see if it contains the service we are looking for.
    if (advertisedDevice.haveServiceUUID() && advertisedDevice.isAdvertisingService(bleServiceUuid)) {
      Serial.println("start to connect");
      BLEDevice::getScan()->stop();
      pDevice = new BLEAdvertisedDevice(advertisedDevice);
      doConnect = true;
      doScan = true;
      Serial.println("vlaue ok");

    } // Found our server
  } // onResult

private:
  BLEUUID bleServiceUuid;
};

class MyClientCallback : public BLEClientCallbacks 
{
public:
  boolean connected = false;
  // setter
  void setConnected(boolean b) { connected = b; }
  // getter
  boolean getConnected() { return connected; }

  void onConnect(BLEClient* pclient) {
  }

  void onDisconnect(BLEClient* pclient) {
    connected = false;
    Serial.println("onDisconnect");
  }
};

class BLEMyClient
{
public:
  BLEMyClient();
  ~BLEMyClient();

  bool connectToServer();

  void begin();
  void init();
  int handleOp();
  static void notifyCallback(
          BLERemoteCharacteristic* pBLERemoteCharacteristic,
          uint8_t* pData,
          size_t length,
          bool isNotify);
  
  // setter 
  void setCharUuid(string s) { charUuid = s; }
  void setServiceUuid(string s) { serviceUuid = s; }
  void setCharacteristicString(string s) { characteristicString = s; }
  void setDoConnect(boolean b) { doConnect = b; }
  void setConnected(boolean b) { connected = b; }
  void setDoScan(boolean b) { doScan = b; }
  // set
  // getter
  string getCharUuid() { return charUuid; }
  string getServiceUuid() { return serviceUuid; }
  string getCharacteristicString() { return characteristicString; }
  boolean getDoConnect() { return doConnect; }
  boolean getConnected() { return connected; }
  boolean getDoScan() { return doScan; }
  // get

private:
  uint8_t *myData;
  size_t dataLength;
  boolean doConnect;
  boolean connected;
  boolean doScan;
  string characteristicString;
  string charUuid = "45678df2-1fd6-45b7-9376-a44282f3b2d5";
  string serviceUuid = "4fafc201-1fb5-459e-8fcc-c5c9c331914b";

  MyAdvertisedDeviceCallbacks* myAdvertisedDeviceCallbacks;
  MyClientCallback* myClientCallback;
  BLERemoteCharacteristic* myRemoteCharacteristic;
  BLERemoteCharacteristic* myBLERemoteCharacteristic;
  BLEAdvertisedDevice* myDevice;
  BLEScan *pBLEScan;
  BLEUUID bleCharUuid;
  BLEUUID bleServiceUuid;
};


#endif 
