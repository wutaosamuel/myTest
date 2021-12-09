// WiFiMyClient.ino/cpp

#include "WiFiMyClient.h"

WiFiMyClient *wifiMyClient;

void setup()
{
  wifiMyClient = new WiFiMyClient();
  wifiMyClient->setWifiSsid("noname");
  wifiMyClient->setPassword("wutaowutao");
  wifiMyClient->setIpAddress(IPAddress(192,168,137,219));

  Serial.begin(115200);
  pinMode(5, OUTPUT);
  wifiMyClient->begin();
  Serial.print("The ip address: ");
  Serial.println(wifiMyClient->getLocalIpAddress());
  Serial.println("client start");
}

void loop()
{
  Serial.println("start to send hello world");

  if(wifiMyClient->sendData("Hello World!") == SUCCESS)
    Serial.println("sent");
  else
    Serial.println("send failure");
  delay(500);
}
