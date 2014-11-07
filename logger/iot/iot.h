/* 
 * File:   iot.h
 * Author: Peter T Mount
 *
 * Created on June 10, 2014, 8:55 PM
 */

#ifndef IOT_H
#define	IOT_H

#include "lib/charbuffer.h"
#include "lib/config.h"
#include "lib/hashmap.h"

struct iot {
    // The http endpoint
    char *endpoint;
    // The userId
    int userId;
    // The deviceID
    int deviceId;
    // The devices api key
    char *apikey;
    // Hashmap used to map sensor name to iot.onl sensorId's
    Hashmap *sensorNameMap;
};

extern int iot_configure(CONFIG_SECTION *sect, struct iot *iot);
extern int iot_lookup_sensorId(struct iot *iot, struct charbuffer *b, char *name);

#endif	/* IOT_H */

