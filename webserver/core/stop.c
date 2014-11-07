/*
 * The embedded webserver
 */

#include <microhttpd.h>
#include <stdlib.h>
#include <stdio.h>
#include "webserver/webserver.h"

extern int verbose;

void webserver_stop() {
    if (webserver.daemon6) {
        if (verbose > 1)
            fprintf(stderr, "Stopping IPv6 webserver\n");

        MHD_stop_daemon(webserver.daemon6);
        webserver.daemon6 = NULL;
    }
    if (webserver.daemon4) {
        if (verbose > 1)
            fprintf(stderr, "Stopping IPv4 webserver\n");

        MHD_stop_daemon(webserver.daemon4);
        webserver.daemon4 = NULL;
    }
}
