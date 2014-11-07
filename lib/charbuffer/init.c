/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

/**
 * Initialise a charbuffer.
 * 
 * WARNING: calling this on an already initialised charbuffer will cause a memory leak!
 * 
 * @param b charbuffer
 * @return 0 if initialised, 1 if no memory
 */
int charbuffer_init(struct charbuffer *b) {
    b->pos = 0;
    b->size = 0;
    b->buffer = malloc(CHARBUFFER_INITIAL_SIZE);
    pthread_mutex_init(&b->mutex, NULL);
    return b->buffer ? CHARBUFFER_OK : CHARBUFFER_ERROR;
}
