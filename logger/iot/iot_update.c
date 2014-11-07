#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <curl/curl.h>
#include "lib/charbuffer.h"
#include "logger/iot/iot.h"

extern verbose;

void iot_update(struct iot *iot, struct charbuffer *b, int sensorId, char *text, char *value, time_t timestamp) {
    CURL *curl = curl_easy_init();
    if (curl) {
        // Form our post data
        charbuffer_reset(b);
        charbuffer_append_form_field(b, "key", iot->apikey);
        charbuffer_append_form_field(b, "text", text);
        charbuffer_append_form_fieldf(b, "value", "%d", value);
        // time is in seconds but 000 in format gives us ms for Java & PostgreSQL
        charbuffer_append_form_fieldf(b, "timestamp", "%ld000", (long int)timestamp);

        int postlen = 0;
        char *postdata = charbuffer_tostring(b, &postlen);
        if (!postdata) {
            curl_easy_cleanup(curl);
            return;
        }

        char *url;
        if (!asprintf(&url, "%s/api/sensor/update/%d/%d/%d", iot->endpoint, iot->userId, iot->deviceId, sensorId)) {
            curl_easy_cleanup(curl);
            free(postdata);
            return;
        }

        if (verbose)
            fprintf(stderr, "iot.onl update: %s %s", url, postdata);

        curl_easy_setopt(curl, CURLOPT_URL, url);
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, postdata);

        if (verbose > 1)
            curl_easy_setopt(curl, CURLOPT_VERBOSE, 1L);

        // Sink the response - otherwise libcurl will write it to the console
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, charbuffer_curl_write_sink);
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

        free(url);
        free(postdata);
    }

}