/**
 * Handles an extensible reusable byte buffer
 */

#include <stdio.h>
#include <stdlib.h>
#include <memory.h>
#include "lib/bytebuffer.h"

/**
 * Resets a buffer so new data can be appended to it using the existing buffer.
 * @param b bytebuffer
 * @return 0 if reset, 1 if error
 */
void bytebuffer_reset(struct bytebuffer *b) {
    pthread_mutex_lock(&b->mutex);
    b->pos = 0;
    pthread_mutex_unlock(&b->mutex);
}
