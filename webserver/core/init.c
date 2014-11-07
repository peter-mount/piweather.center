/*
 * The embedded webserver
 */

#include <microhttpd.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdarg.h>
#include "lib/config.h"
#include "lib/list.h"
#include "webserver/webserver.h"

WEBSERVER webserver;

static const char *errorPage = "<html><head><title>Internal Server Error</title></head><body><h1>Internal Server Error</h1></body></html>";
static const char *notFoundPage = "<html><head><title>Page not found</title></head><body><h1>Page Not Found</h1></body></html>";

void webserver_initialise() {

    memset(&webserver, 0, sizeof (webserver));

    pthread_mutex_init(&webserver.mutex, NULL);

    webserver.port = 0;
    webserver.stack = 0;
    webserver.contextPath = NULL;
    webserver.daemon4 = NULL;
    webserver.daemon6 = NULL;

    webserver.responseHandlers = hashmapCreate(10, hashmapStringHash, hashmapStringEquals);
    list_init(&webserver.handlers);

    // The static responses - these never change
    webserver.errorResponse = MHD_create_response_from_buffer(strlen(errorPage), (void*) errorPage, MHD_RESPMEM_PERSISTENT);
    webserver.notFoundResponse = MHD_create_response_from_buffer(strlen(notFoundPage), (void*) notFoundPage, MHD_RESPMEM_PERSISTENT);

    // Read the config
    CONFIG_SECTION *sect = config_getSection("webserver");
    if (sect) {
        config_getIntParameter(sect, "port", &webserver.port);
        config_getCharParameter(sect, "context", &webserver.contextPath);

        int v = 0;
        config_getBooleanParameter(sect, "ipv4", &v);
        if (v) webserver.stack |= STACK_IPv4;

        v = 0;
        config_getBooleanParameter(sect, "ipv6", &v);
        if (v) webserver.stack |= STACK_IPv6;
    }
}
