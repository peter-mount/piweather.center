
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <curl/curl.h>
#include "lib/charbuffer.h"
#include "lib/config.h"
#include "iot.h"

extern verbose;

int iot_configure(CONFIG_SECTION *sect, struct iot *iot) {
    config_getCharParameter(sect, "iot.api.url", &iot->endpoint);
    config_getCharParameter(sect, "iot.api.key", &iot->apikey);
    config_getIntParameter(sect, "iot.userId", &iot->userId);
    config_getIntParameter(sect, "iot.deviceId", &iot->deviceId);

    iot->sensorNameMap = hashmapCreate(10, hashmapStringHash, hashmapStringEquals);

    return iot->endpoint && iot->apikey && iot->userId && iot->deviceId;
}
