/**
 * Handles an extensible reusable char buffer
 */

#include <stdlib.h>
#include <memory.h>
#include <stdio.h>
#include "lib/charbuffer.h"

int charbuffer_append_center(struct charbuffer *b, char *src, int width) {
    int l, r0, r1, r2, i;

    l = strlen(src);
    if (l > width)
        l = width;

    r0 = width - l;
    r1 = r0 >> 1;
    r2 = r0 - r1;

    for (i = 0; i < r1; i++)
        charbuffer_add(b, ' ');

    int r = charbuffer_put(b, src, l);

    for (
            i = 0; i < r2; i++)
        charbuffer_add(b, ' ');

    return r;
}

