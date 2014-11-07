/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

void charbuffer_read(struct charbuffer *b, FILE *in) {
    char t[1024];
    int i = fread((void*) t, 1, 1024, in);
    while (i > 0) {
        charbuffer_put(b, (void *) t, i);
        i = fread((void*) t, 1, 1024, in);
    }
}
