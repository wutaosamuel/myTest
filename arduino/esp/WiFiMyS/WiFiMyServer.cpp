// WiFiMyServer.cpp

#include "WiFiMyServer.h"

WiFiMyServer::WiFiMyServer()
{
  string ssid = "yourssid";
  string pswd = "yourpassword";
  int p = 5121;
  init(ssid, pswd, p);
}
WiFiMyServer::WiFiMyServer(string ssid, string pswd, int p)
{
  init(ssid, pswd, p);
}

void WiFiMyServer::init(string ssid, string pswd, int p)
{
  wifiSsid = ssid;
  password = pswd;
  port = p;
  clear();
}

void WiFiMyServer::clear()
{
  lineCount = 0;
  data_.clear();
}

void WiFiMyServer::begin()
{
  WiFi.begin(wifiSsid.c_str(), password.c_str());
  while (WiFi.status() != WL_CONNECTED)
    delay(500);
  ipAddress = getESP_ip();
  macAddress = getESP_id();
  subnet = WiFi.subnetMask();
  gateway = WiFi.gatewayIP();
  while (checkWiFi(WiFi.localIP()) == FAILURE)
    delay(500);
  server = new WiFiServer(port, 20);
  server->begin();
}

// TODO: send data should depend on ipAddress and port
int WiFiMyServer::sendData(string s)
{
  // listen for incoming client.
  client= server->available();
  if (!client)
  {
    client.stop();
    return FAILURE;
  }

  if (client && client.connected() && client.available())
    client.print(s.c_str());

  // close connecting
  client.stop();
  return SUCCESS;
}

int WiFiMyServer::readData()
{
  // listen for incoming client.
  client = server->available();
  // detect client not connect
  if (client.connected() == false)
  {
    client.stop();
    return FAILURE;
  }
  // detect client cannot send data
  if (client.available() == false)
  {
    client.stop();
    return FAILURE;
  }

  if (client)
  {
    string buf = "";
    while (client.connected())
    {
      if (client.available())
      {
        char c = client.read();
        // if it is a new line
        if (c == '\n')
        {
          data_.push_back(buf);
          ++lineCount;
          buf = "";
        }
        // if get anything else but a carriage return character
        else if (c != '\r')
        {
          buf += c;
        }
        // if get end of line, push to data storage
        else if (c == '\0')
        {
          data_.push_back(buf);
          ++lineCount;
          buf = "";
        }
      }
    }
    // if client dont send anything return false
    if (data_.size() == 0)
    {
      client.stop();
      return FAILURE;
    }
  }
  else 
  {
    client.stop();
    return FAILURE;
  }

  client.stop();
  return SUCCESS;
}


/*
 * check wifi between x.x.x.80 to x.x.x.180
 * if yes, do nothing 
 * if no, connect with static ip betwen x.x.x.80 to x.x.x.180
 * 
 * @param address - the Ip Address of detecting 
 * @returns 
 */
int WiFiMyServer::checkWiFi(IPAddress address)
{
  // check wifi whether in 80 to 180
  for (uint8_t i=80; i < 181; ++i)
    if (address[3] == i)
      return SUCCESS;
  
  // if it not in 80 to 180
  // connect with static ip between 80 to 180
  for (uint8_t i=80; i < 181; ++i)
  {
    IPAddress bufAddress_ = IPAddress(address[0], address[1], address[2], i);
    WiFi.config(bufAddress_, gateway, subnet);
    if (WiFi.begin(wifiSsid.c_str(), password.c_str()) == WL_CONNECTED)
    {
      ipAddress = getESP_ip();
      subnet = WiFi.subnetMask();
      gateway = WiFi.gatewayIP();
      return SUCCESS;
    }
  }

  return FAILURE;
}

string WiFiMyServer::getESP_id()
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

string WiFiMyServer::getESP_ip()
{
  string buf;
  char bufChar[20];
  IPAddress ip_;

  ip_ = WiFi.localIP();
  sprintf(bufChar, "%d.%d.%d.%d", ip_[0], ip_[1], ip_[2], ip_[3]);
  buf = bufChar;

  return buf;
}
