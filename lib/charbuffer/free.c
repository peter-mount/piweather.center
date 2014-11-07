/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

void charbuffer_free(struct charbuffer *b) {
    if (b->buffer)
        free(b->buffer);
    b->buffer = NULL;
    b->pos = b->size = 0;
}
