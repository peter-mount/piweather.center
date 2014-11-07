/**
 * Handles an extensible reusable byte buffer
 */

#include <stdio.h>
#include <stdlib.h>
#include <memory.h>
#include "lib/bytebuffer.h"

void bytebuffer_free(struct bytebuffer *b) {
    if (b->buffer)
        free(b->buffer);
    b->buffer = NULL;
    b->pos = b->size = 0;
}
