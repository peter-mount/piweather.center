/**
 * Handles an extensible reusable byte buffer
 */

#include <stdio.h>
#include <stdlib.h>
#include <memory.h>
#include "lib/bytebuffer.h"

/**
 * Initialise a bytebuffer.
 * 
 * WARNING: calling this on an already initialised bytebuffer will cause a memory leak!
 * 
 * @param b bytebuffer
 * @return 0 if initialised, 1 if no memory
 */
int bytebuffer_init(struct bytebuffer *b) {
    b->pos = 0;
    b->size = 0;
    b->buffer = malloc(BYTEBUFFER_INITIAL_SIZE);
    pthread_mutex_init(&b->mutex, NULL);
    return b->buffer ? BYTEBUFFER_OK : BYTEBUFFER_ERROR;
}
