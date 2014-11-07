/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

/**
 * Returns a new array containing the buffer's content.
 * 
 * It is up to the application to free the returned array
 * 
 * @param b charbuffer
 * @return array containing the buffer content, null on error
 */

void *charbuffer_toarray(struct charbuffer *b, int *len) {
    if (0 != pthread_mutex_lock(&b->mutex)) {
        return NULL;
    }

    *len = b->pos;

    void *ret = malloc(b->pos);
    if (ret && b->pos)
        memcpy(ret, b->buffer, b->pos);

    pthread_mutex_unlock(&b->mutex);
    return ret;
}
