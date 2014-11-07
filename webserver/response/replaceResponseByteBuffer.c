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

void replaceResponseByteBuffer(const char *url, struct bytebuffer *b, const char *contentType) {
    int len;
    void *data = bytebuffer_toarray(b, &len);
    replaceResponseArray(url, data, len, contentType);
}
