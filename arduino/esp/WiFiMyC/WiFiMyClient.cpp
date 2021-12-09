// WiFiClient.cpp

#include "WiFiMyClient.h"

WiFiMyClient::WiFiMyClient()
{
  string ssid = "yourssid";
  string pswd = "yourpassword";
  int p = 5121;
  int localP = 5120;
  init(ssid, pswd, p, localP);
}

WiFiMyClient::WiFiMyClient(string ssid, string pswd, int p)
{
  init(ssid, pswd, p);
}
WiFiMyClient::WiFiMyClient(string ssid, string pswd, int p, int localP)
{
  init(ssid, pswd, p, localP);
}

void WiFiMyClient::init(string ssid, string pswd, int p)
{
  wifiSsid = ssid;
  password = pswd;
  port = p;
  localPort = 5120;
  clear();
}

void WiFiMyClient::init(string ssid, string pswd, int p, int localP)
{
  wifiSsid = ssid;
  password = pswd;
  port = p;
  localPort = localP;
  clear();
}

void WiFiMyClient::clear()
{
  data_.clear();
}

void WiFiMyClient::begin()
{
  WiFi.begin(wifiSsid.c_str(), password.c_str());
  while (WiFi.status() != WL_CONNECTED)
    delay(500);
  localIpAddress = WiFi.localIP();
  server = new WiFiServer(localPort, 2);
  server->begin();
}

int WiFiMyClient::sendData(string s)
{
  if (client.connect(ipAddress, port) == 0)
  {
    client.stop();
    return FAILURE;
  }
  if (client.connected() == false)
  {
    client.stop();
    return FAILURE;
  }
  // send data or send request
  client.println(s);
  //client.println("HW!");
  // read respond
  //String buf = client.readStringUntil('\r');
  client.stop();
  return SUCCESS;
}

int WiFiMyClient::readData()
{
  // listen for incoming clients
  if (client.connect(ipAddress, port))
  {
    if (client.connected() == false)
    {
      client.stop();
      return FAILURE;
    }

    string buf = "";
    if (client.available())
    {
      char c = client.read();
      // if it is a new line
      if (c == '\n')
      {
        data_.push_back(buf);
        //++lineCount;
        buf = "";
      }
      // if get anything else but a carriage return character
      else if (c != '\r')
      {
        buf += c;
      }
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
