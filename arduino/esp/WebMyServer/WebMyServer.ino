#include <SPIFFS.h>
#include <AsyncTCP.h>
#include <ESPAsyncWebServer.h>

AsyncWebServer server(80);

const char *ssid = "noname";
const char *password = "wutaowutao";

const char *PARAM_MESSAGE = "message";

void notFound(AsyncWebServerRequest *request)
{
    request->send(404, "text/plain", "Not found");
}

void setup()
{

    Serial.begin(115200);
    WiFi.mode(WIFI_STA);
    WiFi.begin(ssid, password);
    if (WiFi.waitForConnectResult() != WL_CONNECTED)
    {
        Serial.printf("WiFi Failed!\n");
        return;
    }

    Serial.print("IP Address: ");
    Serial.println(WiFi.localIP());

    server.serveStatic("/", SPIFFS, "/html/W3CSS-Nature-Portfolio-master/").setDefaultFile("index.html");
    server.begin();
    //server.on("/", HTTP_GET, [](AsyncWebServerRequest *request){
    //request->send(200, "text/plain", "Hello, world");
    //});

    // Send a GET request to <IP>/get?message=<message>
    //server.on("/get", HTTP_GET, [](AsyncWebServerRequest *request) {
        //String message;
        //if (request->hasParam(PARAM_MESSAGE))
        //{
            //message = request->getParam(PARAM_MESSAGE)->value();
        //}
        //else
        //{
            //message = "No message sent";
        //}
        //request->send(200, "text/plain", "Hello, GET: " + message);
    //});

    //// Send a POST request to <IP>/post with a form field message set to <message>
    //server.on("/post", HTTP_POST, [](AsyncWebServerRequest *request) {
        //String message;
        //if (request->hasParam(PARAM_MESSAGE, true))
        //{
            //message = request->getParam(PARAM_MESSAGE, true)->value();
        //}
        //else
        //{
            //message = "No message sent";
        //}
        //request->send(200, "text/plain", "Hello, POST: " + message);
    //});

    //server.onNotFound(notFound);

    //server.begin();
}

void loop()
{
}
