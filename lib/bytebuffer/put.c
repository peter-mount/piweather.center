/**
 * Handles an extensible reusable byte buffer
 */

#include <stdio.h>
#include <stdlib.h>
#include <memory.h>
#include "lib/bytebuffer.h"

static int ensure_capacity(struct bytebuffer *b, int size) {
    if (size > b->size) {
        void *newbuffer = malloc(size);

        if (!newbuffer)
            return BYTEBUFFER_ERROR;

        if (b->buffer) {
            memcpy(newbuffer, b->buffer, b->size);
            free(b->buffer);
        }

        b->size = size;
        b->buffer = newbuffer;
    }
    return BYTEBUFFER_OK;
}

/**
 * append some data to a bytebuffer. If there's not enough room in the buffer then the buffer will be extended to
 * accomodate the data
 * 
 * @param b     bytebuffer
 * @param src   source
 * @param len   length
 * @return 0 if added, 1 if error
 */
int bytebuffer_put(struct bytebuffer *b, void *src, int len) {
    if (0 != pthread_mutex_lock(&b->mutex)) {
        return BYTEBUFFER_ERROR;
    }

    if(ensure_capacity(b, b->pos + len)==BYTEBUFFER_ERROR)
        return BYTEBUFFER_ERROR;
    
    memcpy(b->buffer + b->pos, src, len);
    b->pos += len;

    pthread_mutex_unlock(&b->mutex);
    return BYTEBUFFER_OK;
}
