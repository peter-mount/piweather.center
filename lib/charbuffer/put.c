/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

static int ensure_capacity(struct charbuffer *b, int size) {
    if (size > b->size) {
        char *newbuffer = (char *) malloc(size);

        if (!newbuffer)
            return CHARBUFFER_ERROR;

        if (b->buffer) {
            memcpy(newbuffer, b->buffer, b->size);
            free(b->buffer);
        }

        b->size = size;
        b->buffer = newbuffer;
    }
    return CHARBUFFER_OK;
}

/**
 * append some data to a charbuffer. If there's not enough room in the buffer then the buffer will be extended to
 * accomodate the data
 * 
 * @param b     charbuffer
 * @param src   source
 * @param len   length
 * @return 0 if added, 1 if error
 */
int charbuffer_put(struct charbuffer *b, char *src, int len) {
    if (0 != pthread_mutex_lock(&b->mutex)) {
        return CHARBUFFER_ERROR;
    }

    if (ensure_capacity(b, b->pos + len + 64) == CHARBUFFER_ERROR)
        return CHARBUFFER_ERROR;

    memcpy(b->buffer + b->pos, src, len);
    b->pos += len;

    // Ensure we are terminated
    b->buffer[b->pos] = '\0';

    pthread_mutex_unlock(&b->mutex);
    return CHARBUFFER_OK;
}

/**
 * append some data to a charbuffer. If there's not enough room in the buffer then the buffer will be extended to
 * accomodate the data
 * 
 * @param b     charbuffer
 * @param src   source
 * @param len   length
 * @return 0 if added, 1 if error
 */
int charbuffer_add(struct charbuffer *b, char c) {
    if (0 != pthread_mutex_lock(&b->mutex)) {
        return CHARBUFFER_ERROR;
    }

    // Ensure we have room, but do so in a way we don't keep growing just for 1 character
    ensure_capacity(b, b->pos + (b->pos % 64) + 1);
    b->buffer[b->pos++] = c;

    // Ensure we are terminated
    b->buffer[b->pos] = '\0';

    pthread_mutex_unlock(&b->mutex);
    return CHARBUFFER_OK;
}
