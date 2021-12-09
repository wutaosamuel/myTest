// WiFiMyServer.ino/cpp

#include "WiFiMyServer.h"

WiFiMyServer *wifiMyServer;

void setup()
{
  wifiMyServer = new WiFiMyServer();
  wifiMyServer->setWifiSsid("noname");
  wifiMyServer->setPassword("wutaowutao");
  
  Serial.begin(115200);
  delay(10);
  Serial.println();
  Serial.println();
  wifiMyServer->begin();
  Serial.print("Connecting to ");
  Serial.println(wifiMyServer->getWifiSsid().c_str());
  Serial.print("The ip address: ");
  Serial.println(wifiMyServer->getIpAddress());
  Serial.print("The macAddress: ");
  Serial.println(wifiMyServer->getMacAddress().c_str());
  Serial.print("The gateway: ");
  Serial.println(wifiMyServer->getGateway());
  Serial.print("The subnet mask: ");
  Serial.println(wifiMyServer->getSubnet());
  Serial.println("Waiting for incoming data");
}

void loop()
{
  if (wifiMyServer->readData() == SUCCESS)
  {
    vector <string> buf = wifiMyServer->getData();
    Serial.print("Start read incoming data\n");
    Serial.printf("The size=%d \n", buf.size());
    for (int i=0; i < buf.size(); ++i)
    {
      Serial.printf("%s\n",buf[i].c_str());
    }
    Serial.println("Done");
  }
  else
    wifiMyServer->clear();
}
