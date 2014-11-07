/* 
 * File:   webserver.h
 * Author: peter
 *
 * Created on February 5, 2014, 8:40 PM
 */

#ifndef WEBSERVER_H
#define	WEBSERVER_H

#include <pthread.h>
#include <microhttpd.h>
#include "lib/bytebuffer.h"
#include "lib/charbuffer.h"
#include "lib/config.h"
#include "lib/list.h"
#include "lib/hashmap.h"

#define STACK_IPv4 1
#define STACK_IPv6 2

typedef struct webserverHandler WEBSERVER_HANDLER;

typedef struct {
    // The port we will listen on
    int port;
    // 1 IPv4 support, 2 IPv6, 4=both
    int stack;
    // Our context path
    char *contextPath;
    // Our static response handlers
    Hashmap *responseHandlers;
    //
    struct MHD_Response *homePageResponse;
    // Not found Response
    struct MHD_Response *notFoundResponse;
    // redirection page to the index
    struct MHD_Response *redirectResponse;
    // Internal server error response
    struct MHD_Response *errorResponse;
    // mutex for updating any responses
    pthread_mutex_t mutex;
    // The webserver, one daemon per stack as a single dual stack one listens on ipv6 only
    struct MHD_Daemon *daemon4;
    struct MHD_Daemon *daemon6;
    // Our handlers
    struct List handlers;
} WEBSERVER;

extern WEBSERVER webserver;

/**
 * A handler that will respond to a specific URI on the website
 */
struct webserverHandler {
    struct Node node;
    // The function to handle this request
    int (*handler)(struct MHD_Connection *connection, WEBSERVER_HANDLER *handler);
    // optional userdata
    void *userdata;
};

/**
 * Atomically replaces destResponse with newResponse, disposing the original
 * @param server WEBSERVER
 * @param destResponse Pointer to the pointer to update
 * @param newResponse New response
 */
extern void replaceResponse(const char *url, struct MHD_Response *newResponse);
extern void replaceResponseArray(const char *url, void *data, int len, const char *contentType);
extern void replaceResponseByteBuffer(const char *url, struct bytebuffer *b, const char *contentType);
extern void replaceResponseCharBuffer(const char *url, struct charbuffer *b, const char *contentType);

/**
 * Returns the current value of reponse atomically
 * @param server WEBSERVER
 * @param response Pointer to the pointer to get
 * @return The value at response at the time of the call, NULL if not able to get it
 */
extern struct MHD_Response *getResponse(const char *url);

/**
 * Queues a response to a connection. This does it atomically so that the response will be queued even if another
 * thread replaces the response, i.e. a new camera image.
 * @param connection
 * @param server
 * @param response
 * @return 
 */
extern int queueResponse(struct MHD_Connection * connection, struct MHD_Response **response);

extern char* genurl(const char *contextPath, const char *suffix);

extern int sendResponse(struct MHD_Connection *connection, int status, struct MHD_Response *response);

extern WEBSERVER_HANDLER * webserver_add_handler(const char *url, int (*handler)(struct MHD_Connection *connection, WEBSERVER_HANDLER *handler));
extern void webserver_add_response_handler(const char *url);
extern void webserver_initialise();
extern void webserver_set_defaults();
extern void webserver_start();
extern void webserver_stop();
extern void webserver_update_index();

extern int notFoundHandler(struct MHD_Connection * connection, WEBSERVER_HANDLER *handler);
extern int staticHandler(struct MHD_Connection * connection, const char *url);

#endif	/* WEBSERVER_H */

