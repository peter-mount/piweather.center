
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <curl/curl.h>
#include "lib/charbuffer.h"
#include "iot.h"

extern verbose;

static int invoke(struct iot *iot, struct charbuffer *b, char *name) {
    int sensorId = 0;
    char *url = NULL;

    if (verbose)
        fprintf(stderr, "iot.onl: Looking up sensorId for %s\n", name);

    if (asprintf(&url, "%s/api/sensor/lookup/%d/%d?name=%s", iot->endpoint, iot->userId, iot->deviceId, name)) {
        CURL *curl = curl_easy_init();
        if (curl) {
            curl_easy_setopt(curl, CURLOPT_URL, url);

            if (verbose > 1)
                curl_easy_setopt(curl, CURLOPT_VERBOSE, 1L);

            curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, charbuffer_curl_write);
            curl_easy_setopt(curl, CURLOPT_WRITEDATA, b);
            charbuffer_reset(b);

            CURLcode res = curl_easy_perform(curl);

            long http_code = 0;
            curl_easy_getinfo(curl, CURLINFO_RESPONSE_CODE, &http_code);

            curl_easy_cleanup(curl);

            if (res != CURLE_OK)
                fprintf(stderr, "Failed to submit %s %s\n", url, curl_easy_strerror(res));
            else if (http_code != 200)
                fprintf(stderr, "Failed to submit %s %d\n", url, http_code);
            else {
                int len;
                char *r = charbuffer_tostring(b, &len);
                sensorId = atoi(r);
                free(r);
            }
        }

        free(url);
    }

    return sensorId;
}

/**
 * Translates a sensor name to it's sensorId on iot.onl
 * 
 * This uses the /api/sensor/lookup rest service
 * 
 * @param userId userId
 * @param deviceId deviceId
 * @param b charbuffer
 * @param name name to translate
 * @return sensorId, valid if result is >0
 */
int iot_lookup_sensorId(struct iot *iot, struct charbuffer *b, char *name) {
    int *sensorId = (int *) hashmapGet(iot->sensorNameMap, name);
    
    // No entry or it's 0 then do a remote lookup - allows for service to be offline
    if (!sensorId || sensorId==0) {
        int id = invoke(iot, b, name);

        if (sensorId)
            *sensorId = id;
        else {
            sensorId = (int *) malloc(sizeof (int));
            *sensorId = id;
            hashmapPut(iot->sensorNameMap, strdup(name), (void *) sensorId);
        }
    }
    return *sensorId;
}
