/* 
 * File:   rest.h
 * Author: Peter T Mount
 *
 * Created on April 6, 2014, 9:24 AM
 */

#ifndef REST_H
#define	REST_H

#include "lib/charbuffer.h"
#include "lib/hashmap.h"
#include "logger/logger.h"
#include "webserver/webserver.h"

/*
 * Definition of a read-only rest endpoint
 */
struct rest_service {
    // =============================================
    // The response holding the current output
    // =============================================
    // JSON 
    const char *json;
    // XML
    const char *xml;
    // Text
    const char *text;
    // =============================================
    // JSON list containing last 24 hours of data
//    struct History history;
//    struct MHD_Response *histJson;
};

struct rest_logger {
    struct logger logger;
    // Buffer used in generating the output
    struct charbuffer buffer;
    // Hashmap of sensor names -> rest_service used by json
    Hashmap *services;
};

#endif	/* REST_H */

