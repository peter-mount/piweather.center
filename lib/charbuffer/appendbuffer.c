/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

void charbuffer_appendbuffer(struct charbuffer *dest, struct charbuffer *src) {
    if (0 == pthread_mutex_lock(&src->mutex)) {
        charbuffer_put(dest, src->buffer, src->pos);
        pthread_mutex_unlock(&src->mutex);
    }
}
