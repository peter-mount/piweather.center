/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

/**
 * Returns the size of data currently in the charbuffer
 * 
 * @param b charbuffer
 * @return the size, -1 if the lock cannot be obtained
 */
int charbuffer_size(struct charbuffer *b) {
    int ret = -1;
    if (0 == pthread_mutex_lock(&b->mutex)) {
        // pos is the append point so to the client the size of data present!
        ret = b->pos;
        pthread_mutex_unlock(&b->mutex);
    }
    return ret;
}
