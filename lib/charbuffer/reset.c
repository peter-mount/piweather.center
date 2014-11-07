/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

/**
 * Resets a buffer so new data can be appended to it using the existing buffer.
 * @param b charbuffer
 * @return 0 if reset, 1 if error
 */
void charbuffer_reset(struct charbuffer *b) {
    pthread_mutex_lock(&b->mutex);
    b->pos = 0;
    pthread_mutex_unlock(&b->mutex);
}
