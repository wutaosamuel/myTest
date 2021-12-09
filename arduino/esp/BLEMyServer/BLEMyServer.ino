
#include "BLEMyServer.h"

BleMyServer *bleMyServer = new BleMyServer("noname", "wutaowutao");
void setup()
{
  Serial.begin(115200);
  bleMyServer->init();
  Serial.println("BLE server init ok");
}

void loop()
{
  if (bleMyServer->handleOp() == SUCCESS)
    Serial.println("running ok");
  else if (bleMyServer->handleOp() == 3)
    Serial.println("Wait for device!");
  else
    Serial.println("failure");
  delay(2000);
}
