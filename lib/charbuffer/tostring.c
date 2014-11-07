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
 * This is identical to charbuffer_toarray except this ensures that
 * the string is null terminated
 * 
 * @param b charbuffer
 * @return array containing the buffer content, null on error
 */

char *charbuffer_tostring(struct charbuffer *b, int *len) {
    if (0 != pthread_mutex_lock(&b->mutex)) {
        return NULL;
    }

    *len = b->pos;

    char *ret = (char *) malloc(b->pos + 1);
    if (ret) {
        if (b->pos)
            memcpy(ret, b->buffer, b->pos);
        ret[*len] = '\0';
    }

    pthread_mutex_unlock(&b->mutex);
    return ret;
}
