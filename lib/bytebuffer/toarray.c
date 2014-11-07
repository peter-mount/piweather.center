/**
 * Handles an extensible reusable byte buffer
 */

#include <stdio.h>
#include <stdlib.h>
#include <memory.h>
#include "lib/bytebuffer.h"

/**
 * Returns a new array containing the buffer's content.
 * 
 * It is up to the application to free the returned array
 * 
 * @param b bytebuffer
 * @return array containing the buffer content, null on error
 */

void *bytebuffer_toarray(struct bytebuffer *b, int *len) {
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
