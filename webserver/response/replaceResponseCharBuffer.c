/*
 * The embedded webserver
 */

#include <microhttpd.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdarg.h>
#include "lib/config.h"
#include "webserver/webserver.h"

void replaceResponseCharBuffer(const char *url, struct charbuffer *b, const char *contentType) {
    int len;
    void *data = charbuffer_toarray(b, &len);
    replaceResponseArray(url, data, len, contentType);
}
