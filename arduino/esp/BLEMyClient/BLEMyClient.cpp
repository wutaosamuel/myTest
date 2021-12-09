// BLEMyClient.cpp

#include "BLEMyClient.h"

BLEMyClient::BLEMyClient() 
{
  init();
}
BLEMyClient::~BLEMyClient() {}

void BLEMyClient::begin()
{
  BLEDevice::init("");
  characteristicString = "";

  // Retrieve a Scanner and set the callback we want to use to be informed when we
  // have detected a new device.  Specify that we want active scanning and start the
  // scan to run for 5 seconds.
  pBLEScan = BLEDevice::getScan();
  Serial.println("scan ok");
  myAdvertisedDeviceCallbacks = new MyAdvertisedDeviceCallbacks();
  Serial.println("New ADcallbacks");
  myAdvertisedDeviceCallbacks->setServiceUuid(serviceUuid);
  myAdvertisedDeviceCallbacks->setDoConnect(false);
  myAdvertisedDeviceCallbacks->setConnected(false);
  myAdvertisedDeviceCallbacks->setDoScan(false);
  pBLEScan->setAdvertisedDeviceCallbacks(myAdvertisedDeviceCallbacks);
  pBLEScan->setInterval(1349);
  pBLEScan->setWindow(449);
  pBLEScan->setActiveScan(true);
  pBLEScan->start(5, false);
  Serial.println("ble myclient setup ok");

  doConnect = myAdvertisedDeviceCallbacks->getDoConnect();
  connected = myAdvertisedDeviceCallbacks->getConnected();
  doScan = myAdvertisedDeviceCallbacks->getDoScan();
  myAdvertisedDeviceCallbacks->getBLEDevice(myDevice);
}

void BLEMyClient::init()
{
  doConnect = false;
  connected = false;
  doScan = false;
  bleCharUuid = BLEUUID(charUuid);
  bleServiceUuid = BLEUUID(serviceUuid);
}

void BLEMyClient::notifyCallback(
        BLERemoteCharacteristic* pBLERemoteCharacteristic,
        uint8_t* pData,
        size_t length,
        bool isNotify)
{
    Serial.print("Notify callback for characteristic ");
    Serial.print(pBLERemoteCharacteristic->getUUID().toString().c_str());
    //myBLERemoteCharacteristic = pBLERemoteCharacteristic;
    Serial.print(" of data length ");
    //dataLength = length;
    Serial.println(length);
    Serial.print("data: ");
    //myData = pData;
    Serial.println((char*)pData);
}

bool BLEMyClient::connectToServer()
{
  Serial.print("Forming a connection to ");
  Serial.println(myDevice->getAddress().toString().c_str());
  
  BLEClient *pClient  = BLEDevice::createClient();
  Serial.println(" - Created client");

  myClientCallback = new MyClientCallback();
  myClientCallback->setConnected(connected);
  pClient->setClientCallbacks(myClientCallback);
  connected = myClientCallback->getConnected();
  Serial.println("clientcallback ok");

  // Connect to the remove BLE Server.
  pClient->connect(myDevice);  // if you pass BLEAdvertisedDevice instead of address, it will be recognized type of peer device address (public or private) //Serial.println(" - Connected to server");
  // Obtain a reference to the service we are after in the remote BLE server.
    BLERemoteService* pRemoteService = pClient->getService(bleServiceUuid);
  if (pRemoteService == nullptr) {
    Serial.print("Failed to find our service UUID: ");
    Serial.println(bleServiceUuid.toString().c_str());
    pClient->disconnect();
    return false;
  }
  Serial.println(" - Found our service");

  // Obtain a reference to the characteristic in the service of the remote BLE server.
  myRemoteCharacteristic = pRemoteService->getCharacteristic(bleCharUuid);
  if (myRemoteCharacteristic == nullptr) {
    Serial.print("Failed to find our characteristic UUID: ");
    Serial.println(bleCharUuid.toString().c_str());
    pClient->disconnect();
    return false;
  }
  Serial.println(" - Found our characteristic");

  // Read the value of the characteristic.
  if(myRemoteCharacteristic->canRead()) {
    characteristicString = myRemoteCharacteristic->readValue();
    Serial.print("The characteristic value was: ");
    Serial.println(characteristicString.c_str());
  }

  if(myRemoteCharacteristic->canNotify())
    myRemoteCharacteristic->registerForNotify(notifyCallback);

  connected = true;
}
