/* Automatically generated nanopb constant definitions */
/* Generated by nanopb-0.3.9.3 at Mon Aug 12 22:32:44 2019. */

#include "BleMessage.pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 30
#error Regenerate this file with the current version of nanopb generator.
#endif



const pb_field_t BleMessage_fields[5] = {
    PB_FIELD(  1, INT32   , SINGULAR, STATIC  , FIRST, BleMessage, op, op, 0),
    PB_FIELD(  2, STRING  , SINGULAR, STATIC  , OTHER, BleMessage, mac_address, op, 0),
    PB_FIELD(  3, STRING  , SINGULAR, STATIC  , OTHER, BleMessage, wifiSsid, mac_address, 0),
    PB_FIELD(  4, STRING  , SINGULAR, STATIC  , OTHER, BleMessage, password, wifiSsid, 0),
    PB_LAST_FIELD
};


/* @@protoc_insertion_point(eof) */