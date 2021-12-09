// BLEMyServer.cpp

#include "BLEMyServer.h"

BleMyServer::BleMyServer() 
{
  init();
  wifissid = "NULL";
  password = "NULL";
}
BleMyServer::BleMyServer(string wifissid_, string password_)
{
  init();
  wifissid = wifissid_;
  password = password_;
}
BleMyServer::~BleMyServer() {}

void BleMyServer::init()
{
  op = 0x00;
  macAddress = getESP_id();
  bufferSize = 0;
  setMessage(op, macAddress, wifissid, password);
  buffer[256] = {0};
}

void BleMyServer::begin()
{
  BLEDevice::init("ESP_BLEMyServer");
  pServer = BLEDevice::createServer();
  pService = pServer->createService(SERVICE_UUID);
  remoteCharacteristic = pService->createCharacteristic(
                                remoteMacUuid,
                                BLECharacteristic::PROPERTY_READ |
                                BLECharacteristic::PROPERTY_WRITE
                              );
  remoteCharacteristic->setValue("Waiting for device!");
  pService->start();
  pAdvertising = BLEDevice::getAdvertising();
  pAdvertising->addServiceUUID(SERVICE_UUID);
  pAdvertising->setScanResponse(true);
  pAdvertising->setMinPreferred(0x06);  // functions that help with iPhone connections issue
  pAdvertising->setMinPreferred(0x12);
  pAdvertising->start();
}

int BleMyServer::setBuffer(const uint8_t *buffer_, size_t size_)
{
  if (size_ > (size_t)256)
    return FAILURE;
  bufferSize = size_;
  for (int i = 0; i < (int)size_; ++i)
    buffer[i] = *(buffer_ + i);

  return SUCCESS;
}

int BleMyServer::getBuffer(uint8_t *buffer_, size_t size_)
{
  if (size_ > (size_t)256)
    return FAILURE;
  for (int i = 0; i < (int)size_; ++i)
    *(buffer + i) = buffer_[i];
  
  return SUCCESS;
}

string BleMyServer::getESP_id()
{
  char chipId[12];
  uint16_t chip;
  uint64_t chipid;
  string buf;

  chipid=ESP.getEfuseMac();
  chip = (uint16_t)(chipid >> 32);
  snprintf(chipId, 12, "%04X%08X", chip, (uint32_t)chipid);
  buf = chipId;

  return buf;
}

int BleMyServer::encodeAct()
{
  pb_ostream_t osstream = pb_ostream_from_buffer(buffer, 256);
  if (!pb_encode(&osstream, BleMessage_fields, &bleMessage))
    return FAILURE;
  bufferSize = osstream.bytes_written;
  return SUCCESS;
}

int BleMyServer::decodeAct()
{
  bleMessage = {};
  pb_istream_t isstream = pb_istream_from_buffer(buffer, 256);
  if (!pb_decode(&isstream, BleMessage_fields, &bleMessage))
    return 0;
  op = bleMessage.op;
  macAddress = bleMessage.mac_address;
  wifissid = bleMessage.wifiSsid;
  password = bleMessage.password;
  return 1;
}

void BleMyServer::setMessage(int op_, string mac_, string wifissid_, string password_)
{
  bleMessage = {};
  bleMessage.op = op_;
  strcpy(bleMessage.mac_address, mac_.c_str());
  strcpy(bleMessage.wifiSsid, wifissid_.c_str());
  strcpy(bleMessage.password, password_.c_str());
}

int BleMyServer::handleOp()
{
  // check for value is empty
  if (remoteCharacteristic->getValue().compare("Waiting for device!") != 0)
    return 3;
  // start to decode 
  if (decodeAct() == FAILURE)
    return 4;
    // start to run op
  if (op == 0x00)
  {
    // do nothing
    init();
    setMessage(op, getESP_id(), wifissid, password);
    remoteCharacteristic->setValue("Waiting for device!");
    return 0;
  }
  else if (op == 0x01)
  {
    op = 0x11;
    if (encodeAct() == FAILURE)
      return 2;
    remoteCharacteristic->setValue(buffer, bufferSize);
  }
  else if (op == 0x02)
  {
    op = 0x10;
    setMessage(op, macAddress, "NULL", "NULL");
    if (encodeAct() == FAILURE)
      return 2;
    remoteCharacteristic->setValue(buffer, bufferSize);
  }
  else if (op == 0x10)
  {
    init();
    setMessage(op, getESP_id(), wifissid, password);
    remoteCharacteristic->setValue("Waiting for device!");
  }

  return SUCCESS;
}
