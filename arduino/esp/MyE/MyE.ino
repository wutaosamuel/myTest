// MyE.ino / cpp

#include <pb_encode.h>
#include <pb_decode.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "MyEEPROM.h"
#include "SaveOnServer.pb.h"
#include <ESP32Uuid.h>

int addr = 0;

bool encode_string(pb_ostream_t *stream, const pb_field_t* field, void * const *arg);
bool print_string(pb_istream_t *stream, const pb_field_t *field, void **arg);
SaveOnServer saveOnServer;
size_t mySize;
byte uuid_[16];

void printHex(byte number) {
  int topDigit = number >> 4;
  int bottomDigit = number & 0x0f;
  // Print high hex digit
  Serial.print( "0123456789ABCDEF"[topDigit] );
  // Low hex digit
  Serial.print( "0123456789ABCDEF"[bottomDigit] );
}

void printUuid_char(char* uuidNumber) {
  // Print a UUID in the form
  //   12345678-1234-1234-1234-123456789ABC
  int i;
  for (i=0; i<16; i++) {
    if (i==4) Serial.print("-");
    if (i==6) Serial.print("-");
    if (i==8) Serial.print("-");
    if (i==10) Serial.print("-");
    printHex(uuidNumber[i]);
  }
}

void printUuid(byte* uuidNumber) {
  int i;
  for (i=0; i<16; i++) {
    if (i==4) Serial.print("-");
    if (i==6) Serial.print("-");
    if (i==8) Serial.print("-");
    if (i==10) Serial.print("-");
    printHex(uuidNumber[i]);
  }
}

void setup()
{
  Serial.begin(115200);
  EEPROM.begin(100);
  Serial.println("The initial value in EEPROM is ");
  Serial.println(EEPROM.read(addr));
  Serial.println("Start to make protobuf");
  //uint8_t buffer[100];
  uint8_t *buffer;
  saveOnServer = SaveOnServer_init_zero;
  //pb_ostream_t osstream = pb_ostream_from_buffer(buffer, sizeof(buffer));
  pb_ostream_t osstream = pb_ostream_from_buffer(buffer, 100);
  saveOnServer.device_id_mac.funcs.encode = &encode_string;
  //saveOnServer.device_id_mac.arg = "Hello World";
  saveOnServer.group_id = 0;
  saveOnServer.device_type = 10;
  if (!pb_encode(&osstream, SaveOnServer_fields, &saveOnServer))
  {
    Serial.println("encode error");
    while (1);
  }
  mySize = osstream.bytes_written;
  Serial.printf("size : %d\n", mySize);
  Serial.println("Start to write to EEPROM");
  for (int i=0; i<mySize; ++i)
    EEPROM.write(addr+i, buffer[i]);
  Serial.println("write finish");
  //}
  Serial.println("encode finish");
  ESP32Uuid.uuid4(uuid_);
  Serial.println(ESP32Uuid.uuidToString(uuid_));
}

string bufferString="";
void loop()
{
  uint8_t buffer_[100];
  bufferString="";
  pb_istream_t isstream;
  SaveOnServer readServer = {{{NULL}}};
  Serial.println("Start to read EEPROM");
  for (int i=0; i<mySize; ++i)
    buffer_[i] = EEPROM.read(addr+i);
  isstream = pb_istream_from_buffer(buffer_,mySize);
  Serial.println("Read done");
  Serial.println("Start to decode");
  readServer.device_id_mac.funcs.decode = &print_string;
  //readServer.device_id_mac.arg = bufferChar;
  //Serial.println("add decode function down");
  if (!pb_decode(&isstream, SaveOnServer_fields, &readServer))
  {
    Serial.println("Decode error");
    while(1);
  }
  Serial.printf("Decode down\n");
  //Serial.printf("The device mac: %s\n", readServer.device_id_mac);
  Serial.println(bufferString.c_str());
  //Serial.printf("The device mac: %s\n", bufferChar);
  Serial.printf("The group: %d\n", readServer.group_id);
  Serial.printf("The device type: %d\n", readServer.device_type);

  delay(3000);
}

bool encode_string(pb_ostream_t *stream, const pb_field_t* field, void * const *arg)
{
    char *str = "Hello world!";
    if (!pb_encode_tag_for_field(stream, field))
        return false;
    
    return pb_encode_string(stream, (uint8_t *)str, strlen((str)));
    //return pb_encode_string(stream, *arg, strlen(*arg));
}

bool print_string(pb_istream_t *stream, const pb_field_t *field, void **arg)
{
    uint8_t buffer[100] = {0};
    
    /* We could read block-by-block to avoid the large buffer... */
    if (stream->bytes_left > sizeof(buffer) - 1)
        return false;
    //if (!pb_read(stream, buffer, stream->bytes_left))
        //return false;

    /* Print the string, in format comparable with protoc --decode.
     * Format comes from the arg defined in main().
     */
    while (stream->bytes_left)
    {
      uint64_t value;
      if (!pb_decode_varint(stream, &value))
        return false;
      bufferString += (char)value;
    }
    //strcpy((char*)*arg, bufferString.c_str());
    return true;
}
