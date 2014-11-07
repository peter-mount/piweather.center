/**
 * Handles an extensible reusable byte buffer
 */

#include <stdio.h>
#include <stdlib.h>
#include <memory.h>
#include "lib/bytebuffer.h"

void bytebuffer_write(struct bytebuffer *b, FILE *out) {
    if (0 == pthread_mutex_lock(&b->mutex)) {
        fwrite(b->buffer, 1, b->pos, out);
        pthread_mutex_unlock(&b->mutex);
    }
}
