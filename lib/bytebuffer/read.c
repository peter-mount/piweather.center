/**
 * Handles an extensible reusable byte buffer
 */

#include <stdio.h>
#include <stdlib.h>
#include <memory.h>
#include "lib/bytebuffer.h"

void bytebuffer_read(struct bytebuffer *b, FILE *in) {
    char t[1024];
    int i = fread((void*) t, 1, 1024, in);
    while (i > 0 && bytebuffer_put(b, (void *) t, i) == BYTEBUFFER_OK) {
        i = fread((void*) t, 1, 1024, in);
    }
}
