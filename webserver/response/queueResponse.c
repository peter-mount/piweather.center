/*
 * The embedded webserver
 */

#include <microhttpd.h>
#include <stdlib.h>
#include <pthread.h>
#include "webserver/webserver.h"

/**
 * Queues a response to a connection. This does it atomically so that the response will be queued even if another
 * thread replaces the response, i.e. a new camera image.
 * @param connection
 * @param server
 * @param response
 * @return 
 */
int queueResponse(struct MHD_Connection * connection, struct MHD_Response **response) {
    int ret;
    
    pthread_mutex_lock(&webserver.mutex);
    
    if (*response) {
        ret = MHD_queue_response(connection, MHD_HTTP_OK, *response);
    } else {
        // Return error if there's no response available, usually before an image is available.
        ret = MHD_queue_response(connection, MHD_HTTP_NOT_FOUND, webserver.notFoundResponse);
    }
    
    pthread_mutex_unlock(&webserver.mutex);
    
    return ret;
}
