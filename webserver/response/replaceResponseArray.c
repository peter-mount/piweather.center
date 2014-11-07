/*
 * The embedded webserver
 */

#include <microhttpd.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdarg.h>
#include "webserver/webserver.h"

void replaceResponseArray(const char *url, void *data, int len, const char *contentType) {
    struct MHD_Response *newResponse = MHD_create_response_from_buffer(len, data, MHD_RESPMEM_MUST_FREE);

    if (newResponse) {
        if (contentType)
            MHD_add_response_header(newResponse, "Content-Type", contentType);

        replaceResponse(url, newResponse);
    } else {
        // We must free data as there's no response to manage it
        free(data);
    }
}
