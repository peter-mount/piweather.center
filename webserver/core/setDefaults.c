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

static const char *redir1 = "<html><head><title>Moved</title></head><body><h1>Moved</h1><p>This page has moved to <a href=\"";
static const char *redir2 = "/index.html\">";
static const char *redir3 = "</a></p></body></html>";

void webserver_set_defaults() {

    // Default to 8080 if not set on the command line
    if (!webserver.port) {
        webserver.port = 8080;
    }

    // If neither IPv4 or IPv6 selected on command line default to IPv4
    if (webserver.stack == 0) {
        webserver.stack = STACK_IPv4;
    }

    // Default context path or do we need redirects to the path?
    int addRedirects = webserver.contextPath != NULL;
    if (!addRedirects) {
        // The default contextPath
        webserver.contextPath = "/";
    }

    // The redirection url
    int len = strlen(webserver.contextPath);
    len = len + len + strlen(redir1) + strlen(redir2) + strlen(redir3);
    char *redirPage = (char*) malloc(len + 1);
    strcpy(redirPage, redir1);
    strcat(redirPage, webserver.contextPath);
    strcat(redirPage, redir2);
    strcat(redirPage, webserver.contextPath);
    strcat(redirPage, redir3);
    webserver.redirectResponse = MHD_create_response_from_buffer(len, (void*) redirPage, MHD_RESPMEM_PERSISTENT);
    MHD_add_response_header(webserver.redirectResponse, "Location", webserver.contextPath);

    // Redirect contextPath (either / or /context/) to index.html
    webserver_add_response_handler(webserver.contextPath);
    replaceResponse(webserver.contextPath, webserver.redirectResponse);

    // If we have a contextPath defined then redirect / to that path
    if (addRedirects) {
        webserver_add_response_handler("/");
        replaceResponse("/", webserver.redirectResponse);

        // Redirects on contextPath minus the trailing /
        char *s = strdup(webserver.contextPath);
        char *p = strchr(s, '\0'); // see GNU libc manual on why this is optimal
        *--p = '\0';
        webserver_add_response_handler(s);
        replaceResponse(s, webserver.redirectResponse);
    }
}
