// BLEMyClient.ino

#include "BLEMyClient.h"

BLEMyClient *bleMyClient = new BLEMyClient();
void setup()
{
  Serial.begin(115200);
  Serial.println("Start");
  bleMyClient = new BLEMyClient();
  bleMyClient->begin();
}

void loop() 
{
    // If the flag "doConnect" is true then we have scanned for and found the desired
  // BLE Server with which we wish to connect.  Now we connect to it.  Once we are 
  // connected we set the connected flag to be true.
  Serial.println(bleMyClient->getDoConnect());
  if (bleMyClient->getDoConnect() == true) {
    Serial.println("connect client ok");
    if (bleMyClient->connectToServer()) {
      Serial.println("We are now connected to the BLE Server.");
      Serial.println(bleMyClient->getCharacteristicString().c_str());
    } else {
      Serial.println("We have failed to connect to the server; there is nothin more we will do.");
    }
    bleMyClient->setDoConnect(false);
  }

  // If we are connected to a peer BLE Server, update the characteristic each time we are reached
  // with the current time since boot.
  Serial.println(bleMyClient->getConnected());
  if (bleMyClient->getConnected()) {
    String newValue = "Time since boot: " + String(millis()/1000);
    Serial.println("Setting new characteristic value to \"" + newValue + "\"");
    
    // Set the characteristic's value to be the array of bytes that is actually a string.
    //****pRemoteCharacteristic->writeValue(newValue.c_str(), newValue.length());
  }else if(bleMyClient->getDoScan()){
    BLEDevice::getScan()->start(0);  // this is just eample to start scan after disconnect, most likely there is better way to do it in arduino
  }
  
  delay(1000); // Delay a second between loops.
}
