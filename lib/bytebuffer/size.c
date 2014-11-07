/**
 * Handles an extensible reusable byte buffer
 */

#include <stdio.h>
#include <stdlib.h>
#include <memory.h>
#include "lib/bytebuffer.h"

/**
 * Returns the size of data currently in the bytebuffer
 * 
 * @param b bytebuffer
 * @return the size, -1 if the lock cannot be obtained
 */
int bytebuffer_size(struct bytebuffer *b) {
    int ret = -1;
    if (0 == pthread_mutex_lock(&b->mutex)) {
        // pos is the append point so to the client the size of data present!
        ret = b->pos;
        pthread_mutex_unlock(&b->mutex);
    }
    return ret;
}
