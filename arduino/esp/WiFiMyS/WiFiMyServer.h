// WiFiMyServer.h

#ifndef WIFIMYSERVER_H
#define WIFIMYSERVER_H

#include <string>
#include <vector>

#include <WiFi.h>
#include <Arduino.h>
#include <WiFiServer.h>

#define SUCCESS 1
#define FAILURE !SUCCESS

using namespace std;

class WiFiMyServer
{
public:
  WiFiMyServer();
  WiFiMyServer(string, string, int);
  ~WiFiMyServer();
  void init(string, string, int);
  void begin();
  void clear();
  int sendData(string);
  int readData();

  // setter
  void setPort(int p) { port = p; }
  void setWifiSsid(string s) { wifiSsid = s; }
  void setPassword(string s) { password = s; }
  void setIpAddress(string s) { ipAddress = s; }
  // getter
  int getPort() { return port;}
  string getWifiSsid() { return wifiSsid; }
  string getPassword() { return password; }
  string getIpAddress() { return ipAddress; }
  // get
  vector <string> getData() { return data_; }
  string getMacAddress() { return macAddress; }
  IPAddress getGateway() { return gateway; }
  IPAddress getSubnet() { return subnet; }

private:
  int port;
  int lineCount;
  string wifiSsid;
  string password;
  string ipAddress;
  string macAddress;
  vector <string> data_;

  WiFiClient client;
  WiFiServer* server;
  IPAddress subnet;
  IPAddress gateway;

  int checkWiFi(IPAddress);
  string getESP_id();
  string getESP_ip();
};

#endif
