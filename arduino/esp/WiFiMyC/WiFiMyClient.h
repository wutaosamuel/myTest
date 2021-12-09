// WiFiMyClient.h

#ifndef WIFIMYCLIENT_H
#define WIFIMYCLIENT_H

#include <WiFi.h>
#include <string>
#include <vector>

#define SUCCESS 1
#define FAILURE !SUCCESS

using namespace std;

class WiFiMyClient
{
public:
  WiFiMyClient();
  WiFiMyClient(string, string, int);
  WiFiMyClient(string, string, int, int);
  ~WiFiMyClient();
  void init(string, string, int);
  void init(string, string, int, int);
  void begin();
  void clear();
  int sendData(string);
  int readData();

  // setter
  void setPort(int p) { port = p; }
  void setLocalPort(int p) { localPort = p; }
  void setWifiSsid(string s) { wifiSsid = s; }
  void setPassword(string s) { password = s; }
  void setIpAddress(IPAddress s) { ipAddress = s; }
  void setLocalIpAddress(IPAddress s) { localIpAddress = s; }
  // getter
  int getPort() { return port;}
  int getLocalPort() { return localPort;}
  string getWifiSsid() { return wifiSsid; }
  string getPassword() { return password; }
  IPAddress getIpAddress() { return ipAddress; }
  IPAddress getLocalIpAddress() { return localIpAddress; }

private:
  int port;
  int localPort;
  string wifiSsid;
  string password;
  vector <string> data_;

  IPAddress ipAddress;
  IPAddress localIpAddress;
  WiFiClient client;
  WiFiServer* server;
};

#endif